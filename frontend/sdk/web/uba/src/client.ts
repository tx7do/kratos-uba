/**
 * UbaClient —— SDK 核心类
 *
 * 职责：
 * - 持有鉴权凭证（appId/appSecret）
 * - 提供 track / trackBehavior / trackRisk / identify / setSuperProperties / flush 等高层 API
 * - 自动补全公共字段（deviceId/sessionId/eventTime/platform/clientInfo/superProperties）
 * - 通过 Batcher 实现缓冲 + 批量上报
 */

import { Batcher } from './batcher';
import {
  getDeviceId,
  getSessionId,
  detectPlatform,
  getClientInfo,
  getPageUrl,
  getViewportWidth,
  buildElementXPath,
  getClickCoords,
  isBrowser,
} from './context';
import { watchRouteChanges, type RouteWatcherHandle } from './spa';
import { uuid, toRFC3339, merge } from './utils';
import {
  EventType,
  type UbaConfig,
  type ReportEvent,
  type TrackOptions,
  type ClientInfo,
  type PostReportResponse,
  type RiskEvent,
} from './types';

const DEFAULT_PATH = '/uba/v1/report';
const DEFAULT_BATCH_SIZE = 20;
const DEFAULT_FLUSH_INTERVAL = 5000;
const DEFAULT_MAX_RETRIES = 3;
const DEFAULT_TIMEOUT = 8000;
const DEFAULT_BASE_DELAY = 1000;
const DEFAULT_SAMPLE_RATE = 1;
const DEFAULT_PAGE_VIEW_EVENT = 'page_view';
const DEFAULT_MAX_PERSISTED_EVENTS = 500;
/** 页面停留小于该阈值（毫秒）不补发 page_dwell，避免噪声 */
const MIN_DWELL_TO_REPORT = 1000;

export class UbaClient {
  private config: Required<Pick<UbaConfig, 'appId' | 'appSecret' | 'endpoint' | 'path' | 'batchSize' | 'flushInterval' | 'maxRetries' | 'timeout' | 'retryBaseDelay' | 'debug' | 'enableBeacon' | 'autoTrack' | 'autoTrackPv' | 'pageViewEvent' | 'trackDwell' | 'sampleRate' | 'requireConsent' | 'respectDoNotTrack' | 'persistQueue' | 'maxPersistedEvents'>>;
  private batcher: Batcher;

  /** 当前登录用户 ID（identify 设置后注入每条事件） */
  private userId?: number;
  /** 公共属性，自动注入每条事件 */
  private superProperties: Record<string, string> = {};
  /** 平台缓存 */
  private platform: string;
  /** autotrack click 监听器引用（destroy 时解绑） */
  private clickHandler?: (e: MouseEvent) => void;
  /** SPA 路由监听句柄（destroy 时解绑） */
  private routeWatcher?: RouteWatcherHandle;
  /** 隐私授权状态（requireConsent=true 时未授权则不采集） */
  private consentGranted = true;
  /** 会话级采样判定结果（缓存，undefined=未判定） */
  private sampledIn?: boolean;
  /** 整体禁用（DNT）标记 */
  private disabled = false;
  /** 当前 PV 开始时间（停留计算基准） */
  private lastPvTime = 0;

  private constructor(config: UbaConfig) {
    this.config = {
      appId: config.appId,
      appSecret: config.appSecret,
      endpoint: config.endpoint.replace(/\/+$/, ''), // 去尾部斜杠
      path: config.path || DEFAULT_PATH,
      batchSize: config.batchSize ?? DEFAULT_BATCH_SIZE,
      flushInterval: config.flushInterval ?? DEFAULT_FLUSH_INTERVAL,
      maxRetries: config.maxRetries ?? DEFAULT_MAX_RETRIES,
      timeout: config.timeout ?? DEFAULT_TIMEOUT,
      retryBaseDelay: config.retryBaseDelay ?? DEFAULT_BASE_DELAY,
      debug: config.debug ?? false,
      enableBeacon: config.enableBeacon ?? true,
      autoTrack: config.autoTrack ?? true,
      autoTrackPv: config.autoTrackPv ?? true,
      pageViewEvent: config.pageViewEvent || DEFAULT_PAGE_VIEW_EVENT,
      trackDwell: config.trackDwell ?? true,
      sampleRate: config.sampleRate ?? DEFAULT_SAMPLE_RATE,
      requireConsent: config.requireConsent ?? false,
      respectDoNotTrack: config.respectDoNotTrack ?? false,
      persistQueue: config.persistQueue ?? true,
      maxPersistedEvents: config.maxPersistedEvents ?? DEFAULT_MAX_PERSISTED_EVENTS,
    };

    // Do Not Track：用户明确拒绝跟踪时整体禁用采集
    if (isBrowser() && this.config.respectDoNotTrack && (navigator as any).doNotTrack === '1') {
      this.disabled = true;
    }
    // 隐私授权：要求授权时默认视为未授权，等 setConsent(true)
    if (this.config.requireConsent) {
      this.consentGranted = false;
    }

    this.platform = detectPlatform();
    const url = this.config.endpoint + this.config.path;

    this.batcher = new Batcher({
      appId: this.config.appId,
      appSecret: this.config.appSecret,
      url,
      batchSize: this.config.batchSize,
      flushInterval: this.config.flushInterval,
      maxRetries: this.config.maxRetries,
      timeout: this.config.timeout,
      baseDelay: this.config.retryBaseDelay,
      enableBeacon: this.config.enableBeacon,
      persistQueue: this.config.persistQueue,
      maxPersistedEvents: this.config.maxPersistedEvents,
      getClientInfo: () => this.resolveClientInfo(),
      log: (level, msg) => this.log(level, msg),
    });

    // 绑定页面卸载兜底
    if (isBrowser() && this.config.enableBeacon) {
      this.bindUnload();
    }
    // 自动埋点：监听 click 上报点击事件
    if (isBrowser() && this.config.autoTrack) {
      this.enableAutoTrack();
    }
    // 自动 PV：首屏 + SPA 路由变化 + 页面停留
    if (isBrowser() && this.config.autoTrackPv && !this.disabled) {
      this.enableAutoPv();
    }
  }

  /**
   * 初始化 SDK（单例）。
   * 多次调用以最后一次为准（重新创建 batcher）。
   */
  static init(config: UbaConfig): UbaClient {
    if (!config.appId || !config.appSecret) {
      throw new Error('UbaClient.init: appId and appSecret are required');
    }
    if (!config.endpoint) {
      throw new Error('UbaClient.init: endpoint is required');
    }
    // 单例：复用全局实例
    const g = globalThis as any;
    if (g.__uba_client__ instanceof UbaClient) {
      g.__uba_client__.destroy();
    }
    const client = new UbaClient(config);
    g.__uba_client__ = client;
    return client;
  }

  /** 获取单例（init 之后可用） */
  static getInstance(): UbaClient | undefined {
    return (globalThis as any).__uba_client__ as UbaClient | undefined;
  }

  /**
   * 通用埋点：上报一个行为事件。
   * @param eventName 事件名称（必填）
   * @param properties 自定义属性，并入 properties 字段
   * @param options 其它字段（userId/objectType/amount/metrics 等）
   */
  track(eventName: string, properties?: Record<string, string>, options?: TrackOptions): void {
    this.trackBehavior(eventName, properties, options);
  }

  /** 行为事件埋点（显式语义） */
  trackBehavior(eventName: string, properties?: Record<string, string>, options?: TrackOptions): void {
    const event = this.buildEvent(EventType.Behavior, eventName, properties, options);
    // 构造 behavior oneof payload（与顶层业务字段对齐，便于服务端落库）
    event.behavior = {
      eventAction: options?.eventAction,
      objectType: options?.objectType,
      objectId: options?.objectId,
      objectName: options?.objectName,
      durationMs: options?.durationMs,
      amount: options?.amount,
      quantity: options?.quantity,
      score: options?.score,
      metrics: options?.metrics,
      serverId: options?.serverId,
      level: options?.level,
      clickX: options?.clickX,
      clickY: options?.clickY,
      elementXpath: options?.elementXpath,
      pageUrl: event.pageUrl,
      viewportWidth: event.viewportWidth,
    };
    this.enqueue(event);
  }

  /** 风险事件埋点 */
  trackRisk(eventName: string, risk: Partial<RiskEvent>, options?: TrackOptions): void {
    const event = this.buildEvent(EventType.Risk, eventName, options?.properties, options);
    // 构造 risk oneof payload
    event.risk = {
      ...risk,
    };
    this.enqueue(event);
  }

  /** 设置当前登录用户 ID，后续事件自动带上 */
  identify(userId: number): void {
    this.userId = userId;
  }

  /** 清除登录用户（登出） */
  resetUser(): void {
    this.userId = undefined;
  }

  /** 设置公共属性，注入后续每条事件的 properties */
  setSuperProperties(props: Record<string, string>): void {
    this.superProperties = merge(this.superProperties, props);
  }

  /** 清除公共属性 */
  clearSuperProperties(): void {
    this.superProperties = {};
  }

  /**
   * 隐私授权开关（GDPR/个保法场景）。
   * requireConsent=true 时，未授权前不采集不发送；授予后自动恢复（含补发离线队列）。
   */
  setConsent(granted: boolean): void {
    this.consentGranted = granted;
  }

  /** 手动上报一次页面浏览（autoTrackPv 关闭时可自行控制 PV 时机） */
  trackPageView(): void {
    this.sendPageView(0);
  }

  /** 当前是否处于采集状态（DNT/授权/采样综合判定，调试用） */
  isCollecting(): boolean {
    return !this.disabled && this.checkConsent() && this.checkSampled();
  }

  /** 手动触发批量上报 */
  async flush(): Promise<void> {
    await this.batcher.flush();
  }

  /** 销毁 SDK：停止定时器。通常无需手动调用。 */
  destroy(): void {
    this.batcher.destroy();
    // 解绑 autotrack click 监听
    if (this.clickHandler && isBrowser()) {
      document.removeEventListener('click', this.clickHandler, true);
      this.clickHandler = undefined;
    }
    // 解绑 SPA 路由监听与 pagehide
    if (this.routeWatcher) {
      this.routeWatcher.destroy();
      this.routeWatcher = undefined;
    }
    if (isBrowser()) {
      window.removeEventListener('pagehide', this.pageHideHandler);
    }
    const g = globalThis as any;
    if (g.__uba_client__ === this) {
      g.__uba_client__ = undefined;
    }
  }

  /**
   * 自动埋点：监听 document click，捕获时计算坐标+元素 xpath，
   * 自动上报 'click' 事件。初始化时按 config.autoTrack 默认开启。
   */
  enableAutoTrack(): void {
    if (!isBrowser() || this.clickHandler) {
      return;
    }
    this.clickHandler = (e: MouseEvent) => {
      try {
        const target = e.target;
        const coords = getClickCoords(e);
        const xpath = buildElementXPath(target);
        const pageUrl = getPageUrl();
        const viewportWidth = getViewportWidth();
        this.track('click', undefined, {
          clickX: coords.x,
          clickY: coords.y,
          elementXpath: xpath,
          pageUrl,
          viewportWidth,
        });
      } catch (err) {
        this.log('warn', `autotrack click failed: ${err}`);
      }
    };
    // 捕获阶段绑定，确保先于业务 click 逻辑采集
    document.addEventListener('click', this.clickHandler, true);
  }

  /** 当前队列中待上报的事件数 */
  pendingCount(): number {
    return this.batcher.size();
  }

  // ──────────── 自动 PV / 页面停留 ────────────

  /**
   * 自动 PV：首屏立即上报一次 page_view，之后监听 SPA 路由变化，
   * 每次变化上报新页面的 page_view，durationMs 带上一页停留时长。
   */
  enableAutoPv(): void {
    if (!isBrowser() || this.routeWatcher) {
      return;
    }

    this.lastPvTime = Date.now();
    this.sendPageView(0);

    this.routeWatcher = watchRouteChanges(() => {
      const dwell = Date.now() - this.lastPvTime;
      this.lastPvTime = Date.now();
      this.sendPageView(dwell);
    });

    // 页面卸载前补发最后一页的停留，再走 beacon 兜底
    if (this.config.trackDwell) {
      window.addEventListener('pagehide', this.pageHideHandler);
    }
  }

  private currentRoute(): string {
    return window.location.pathname + window.location.search + window.location.hash;
  }

  private sendPageView(dwellMs: number): void {
    if (this.config.debug) {
      this.log('info', `auto pv: ${this.currentRoute()} (prev dwell=${dwellMs}ms)`);
    }
    let title: string | undefined;
    try {
      title = document.title;
    } catch {
      // 忽略
    }
    this.trackBehavior(this.config.pageViewEvent, title ? { title } : undefined, {
      durationMs: dwellMs > 0 ? dwellMs : undefined,
      pageUrl: getPageUrl(),
      viewportWidth: getViewportWidth(),
      eventCategory: 'navigation',
    });
  }

  private pageHideHandler = () => {
    try {
      const dwell = Date.now() - this.lastPvTime;
      if (dwell >= MIN_DWELL_TO_REPORT) {
        this.trackBehavior('page_dwell', undefined, {
          durationMs: dwell,
          pageUrl: getPageUrl(),
          eventCategory: 'navigation',
        });
      }
    } catch {
      // 忽略
    }
    this.batcher.flushBeacon();
  };

  // ──────────── 内部方法 ────────────

  /** 构造一个 ReportEvent，自动补全公共字段 */
  private buildEvent(
    eventType: EventType,
    eventName: string,
    properties?: Record<string, string>,
    options?: TrackOptions,
  ): ReportEvent {
    const eventTime = toRFC3339();
    // 合并 properties：superProperties（基础）+ 本次传入
    const mergedProps = merge(this.superProperties, properties);
    // 注入 pageUrl 便于路径分析（若在浏览器且有值）
    const pageUrl = getPageUrl();
    if (pageUrl && !mergedProps.pageUrl) {
      mergedProps.pageUrl = pageUrl;
    }

    return {
      eventType,
      eventId: uuid(),
      eventName,
      eventTime,
      userId: options?.userId ?? this.userId,
      deviceId: options?.deviceId ?? getDeviceId(),
      sessionId: options?.sessionId ?? getSessionId(),
      platform: options?.platform ?? this.platform,
      eventCategory: options?.eventCategory,
      eventAction: options?.eventAction,
      objectType: options?.objectType,
      objectId: options?.objectId,
      objectName: options?.objectName,
      durationMs: options?.durationMs,
      amount: options?.amount,
      quantity: options?.quantity,
      score: options?.score,
      metrics: options?.metrics,
      // 游戏专属维度（游戏方在 track 时传入）
      serverId: options?.serverId,
      level: options?.level,
      // 页面级公共字段（点击热力图/路径分析用）
      clickX: options?.clickX,
      clickY: options?.clickY,
      elementXpath: options?.elementXpath,
      pageUrl: options?.pageUrl ?? pageUrl,
      viewportWidth: options?.viewportWidth ?? (isBrowser() ? getViewportWidth() : undefined),
      properties: Object.keys(mergedProps).length > 0 ? mergedProps : undefined,
    };
  }

  private enqueue(event: ReportEvent): void {
    if (this.disabled || !this.checkConsent() || !this.checkSampled()) {
      if (this.config.debug) {
        this.log('info', `skip event (disabled/consent/sample): ${event.eventName}`);
      }
      return;
    }
    this.batcher.enqueue(event);
    if (this.config.debug) {
      this.log('info', `enqueued event: ${event.eventName} (queue=${this.batcher.size()})`);
    }
  }

  private checkConsent(): boolean {
    return !this.config.requireConsent || this.consentGranted;
  }

  /** 会话级确定性采样：按 sessionId 判定一次并缓存，同一会话全采或全不采 */
  private checkSampled(): boolean {
    if (this.config.sampleRate >= 1) {
      return true;
    }
    if (this.config.sampleRate <= 0) {
      return false;
    }
    if (this.sampledIn === undefined) {
      const sessionId = getSessionId();
      let hash = 0;
      for (let i = 0; i < sessionId.length; i++) {
        hash = (hash * 31 + sessionId.charCodeAt(i)) >>> 0;
      }
      this.sampledIn = hash % 1000 < this.config.sampleRate * 1000;
      if (this.config.debug) {
        this.log('info', `sampling decision for session ${sessionId}: ${this.sampledIn}`);
      }
    }
    return this.sampledIn;
  }

  private resolveClientInfo(): ClientInfo | undefined {
    return getClientInfo();
  }

  private bindUnload(): void {
    const handler = () => this.batcher.flushBeacon();
    // visibilitychange（移动端/现代浏览器）+ beforeunload 双保险
    window.addEventListener('pagehide', handler);
    window.addEventListener('beforeunload', handler);
  }

  private log(level: 'warn' | 'error' | 'info', msg: string): void {
    if (!this.config.debug && level !== 'error') {
      return;
    }
    const prefix = '[UBA SDK]';
    // eslint-disable-next-line no-console
    const fn = (console as any)[level === 'info' ? 'log' : level] || console.log;
    try {
      fn.call(console, prefix, msg);
    } catch {
      // console 不可用时忽略
    }
  }
}

/** 重新导出响应类型，便于调用方处理 flush 结果 */
export type { PostReportResponse };
