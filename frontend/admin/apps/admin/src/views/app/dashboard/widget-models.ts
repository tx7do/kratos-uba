/**
 * 自定义看板支持的分析模型注册表。
 *
 * widget.paramsJson 与对应分析接口的请求体一致，但时间窗支持 lastDays
 * 相对天数约定（渲染时由 dashboard-widget 换算为绝对 timeRange），
 * 避免把快照时间写死后看板"过期"。
 */

export interface WidgetModelOption {
  /** i18n 文案 key */
  labelKey: string;
  /** 分析模型标识（dashboard-widget.vue 的 render switch 分支） */
  value: string;
}

export const WIDGET_MODELS: WidgetModelOption[] = [
  {
    labelKey: 'page.dashboard.widget.modelActiveUsers',
    value: 'active-users',
  },
  { labelKey: 'page.dashboard.widget.modelEventTrend', value: 'event-trend' },
  { labelKey: 'page.dashboard.widget.modelGroupBy', value: 'group-by' },
  { labelKey: 'page.dashboard.widget.modelFunnel', value: 'funnel' },
  { labelKey: 'page.dashboard.widget.modelRetention', value: 'retention' },
];

/** 各模型的默认参数模板（新组件 / 切换模型时预填） */
const MODEL_DEFAULT_PARAMS: Record<string, Record<string, any>> = {
  'active-users': { granularity: 'DAY', lastDays: 7 },
  'event-trend': {
    eventName: '',
    granularity: 'ANALYTICS_GRANULARITY_UNSPECIFIED',
    lastDays: 7,
  },
  'group-by': { dimension: 'platform', lastDays: 7, metric: 'COUNT', topN: 10 },
  funnel: { lastDays: 7, steps: ['login', 'purchase'], windowMs: 1_800_000 },
  retention: { lastDays: 7, maxOffsetDays: 7, retentionType: 'ACTIVE' },
};

/** 生成某模型的默认参数 JSON（格式化缩进，便于编辑） */
export function defaultParamsJson(model: string): string {
  const params = MODEL_DEFAULT_PARAMS[model] || { lastDays: 7 };
  return JSON.stringify(params, null, 2);
}

/** 是否为已知模型 */
export function isKnownModel(model?: string): boolean {
  return !!model && WIDGET_MODELS.some((m) => m.value === model);
}

/** 模型标识 → i18n 文案 key（仅对已知模型调用） */
export function modelLabelKey(model?: string): string {
  return (
    WIDGET_MODELS.find((m) => m.value === model)?.labelKey ||
    'page.dashboard.widget.modelEventTrend'
  );
}
