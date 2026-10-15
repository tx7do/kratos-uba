<script lang="ts" setup>
import type {
  ubaservicev1_DashboardWidget,
  ubaservicev1_EventTrendResponse,
  ubaservicev1_FunnelResponse,
  ubaservicev1_GroupByResponse,
  ubaservicev1_RetentionResponse,
} from '#/generated/api/admin/service/v1';

import { onMounted, ref } from 'vue';

import {
  EchartsUI,
  type EchartsUIType,
  useEcharts,
} from '@vben/plugins/echarts';

import { Card } from 'ant-design-vue';

import {
  fetchActiveUsers,
  fetchEventTrend,
  fetchFunnel,
  fetchGroupBy,
  fetchRetention,
  lastDaysRange,
} from '#/api';

interface Props {
  widget: ubaservicev1_DashboardWidget;
  /** 所属看板的应用 ID，组件参数未指定 appId 时继承 */
  appId?: number;
}

const props = defineProps<Props>();

const chartRef = ref<EchartsUIType>();
const { renderEcharts } = useEcharts(chartRef);
const loaded = ref(false);
const errorMsg = ref('');

function parseParams(): Record<string, any> {
  let params: Record<string, any>;
  try {
    params = JSON.parse(props.widget.paramsJson || '{}');
  } catch {
    return {};
  }
  if (typeof params !== 'object' || params === null || Array.isArray(params)) {
    return {};
  }
  // lastDays 相对天数约定：渲染时才换算绝对时间窗，看板不会"过期"
  const lastDays = Number(params.lastDays);
  if (Number.isFinite(lastDays) && lastDays > 0) {
    params.timeRange = lastDaysRange(lastDays);
    delete params.lastDays;
  }
  // 组件未指定应用时继承看板的应用（0/undefined 表示全部应用）
  if (params.appId === undefined && props.appId) {
    params.appId = props.appId;
  }
  return params;
}

// 时间桶格式化为坐标轴标签（protojson 会把 int64 序列化为字符串，需转数值）
function timeLabel(ts?: number | string) {
  const n = Number(ts);
  if (!ts || !Number.isFinite(n)) return '';
  const d = new Date(n);
  return `${d.getMonth() + 1}/${d.getDate()} ${String(d.getHours()).padStart(2, '0')}:00`;
}

// 日期（毫秒，可能为字符串）格式化为 MM/DD
function dayLabel(ms?: number | string) {
  const n = Number(ms);
  if (!ms || !Number.isFinite(n)) return '-';
  const d = new Date(n);
  return `${d.getMonth() + 1}/${d.getDate()}`;
}

function renderTrend(title: string, resp: ubaservicev1_EventTrendResponse) {
  const series = (
    resp.series && resp.series.length > 0
      ? resp.series
      : [{ name: title, points: resp.points }]
  ).map((s) => ({
    name: s.name || title,
    data: (s.points || []).map((p) => [timeLabel(p.timestamp), p.value ?? 0]),
    showSymbol: false,
    type: 'line' as const,
    smooth: true,
  }));

  renderEcharts({
    legend: { top: '2%' },
    grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
    series,
    tooltip: { trigger: 'axis' },
    xAxis: { type: 'category' },
    yAxis: { type: 'value' },
  });
}

async function render() {
  const params = parseParams();
  const model = props.widget.model;

  errorMsg.value = '';
  try {
    switch (model) {
      case 'active-users': {
        const resp = await fetchActiveUsers(params as any);
        const points = resp.points || [];
        renderEcharts({
          legend: { top: '2%' },
          series: [
            {
              data: points.map((p) => [timeLabel(p.timestamp), p.dau ?? 0]),
              name: 'DAU',
              showSymbol: false,
              smooth: true,
              type: 'line',
            },
            {
              data: points.map((p) => [timeLabel(p.timestamp), p.wau ?? 0]),
              name: 'WAU',
              showSymbol: false,
              smooth: true,
              type: 'line',
            },
          ],
          tooltip: { trigger: 'axis' },
          xAxis: { type: 'category' },
          yAxis: { type: 'value' },
        });

        break;
      }
      case 'event-trend': {
        renderTrend(
          props.widget.title || 'event-trend',
          await fetchEventTrend(params as any),
        );

        break;
      }
      case 'funnel': {
        const resp: ubaservicev1_FunnelResponse = await fetchFunnel(
          params as any,
        );
        const data = (resp.steps || []).map((step) => ({
          name: step.eventName || `step${step.stepIndex}`,
          value: Number(step.count ?? 0),
        }));
        renderEcharts({
          series: [
            {
              data: data.length > 0 ? data : [{ name: '-', value: 1 }],
              funnelAlign: 'center',
              gap: 2,
              label: { position: 'inside', formatter: '{b}: {c}' },
              sort: 'descending',
              type: 'funnel',
            },
          ],
          tooltip: { trigger: 'item' },
        });

        break;
      }
      case 'group-by': {
        const resp: ubaservicev1_GroupByResponse = await fetchGroupBy(
          params as any,
        );
        const data = (resp.buckets || []).map((b) => ({
          name: b.label || 'unknown',
          value: Number(b.value ?? 0),
        }));
        renderEcharts({
          legend: { bottom: '2%' },
          series: [
            {
              avoidLabelOverlap: false,
              data: data.length > 0 ? data : [{ name: '-', value: 1 }],
              label: { position: 'center', show: false },
              radius: ['35%', '62%'],
              type: 'pie',
            },
          ],
          tooltip: { trigger: 'item' },
        });

        break;
      }
      case 'retention': {
        const resp: ubaservicev1_RetentionResponse = await fetchRetention(
          params as any,
        );
        // 热力图：行=同期群日期，列=偏移天，值=留存率(%)
        const cohorts = resp.cohorts || [];
        const offsets = resp.offsetDays || [];
        const xLabels = offsets.map((d) => `+${d}d`);
        const yLabels = cohorts.map((c) => dayLabel(c.cohortDate));
        const heatData: [number, number, number][] = [];
        cohorts.forEach((c, y) => {
          const byOffset = new Map(
            (c.cells || []).map((cell) => [
              cell.offsetDays ?? 0,
              cell.rate ?? 0,
            ]),
          );
          offsets.forEach((offset, x) => {
            const rate = byOffset.get(offset);
            if (rate !== undefined) {
              heatData.push([x, y, Math.round(rate * 1000) / 10]);
            }
          });
        });
        renderEcharts({
          grid: { left: '3%', right: '8%', bottom: '12%', containLabel: true },
          series: [
            {
              data: heatData,
              label: {
                formatter: (p: any) =>
                  p.value[2] === null || p.value[2] === undefined
                    ? ''
                    : `${p.value[2]}%`,
                show: true,
              },
              type: 'heatmap',
            },
          ],
          tooltip: {
            formatter: (p: any) =>
              `${yLabels[p.value[1]]} · ${xLabels[p.value[0]]}：${p.value[2]}%`,
            trigger: 'item',
          },
          visualMap: {
            bottom: '0',
            calculable: true,
            inRange: { color: ['#5073a5', '#f5f5f5', '#e2504c'] },
            left: 'center',
            max: 100,
            min: 0,
            orient: 'horizontal',
          },
          xAxis: { type: 'category', data: xLabels },
          yAxis: { type: 'category', data: yLabels },
        });

        break;
      }
      // No default
    }
    loaded.value = true;
  } catch (error) {
    loaded.value = false;
    const message = error instanceof Error ? error.message : String(error);
    errorMsg.value =
      message.length > 120 ? `${message.slice(0, 117)}...` : message;
  }
}

onMounted(render);
</script>

<template>
  <Card :title="widget.title" size="small">
    <template #extra>
      <a-tag>{{ widget.model }}</a-tag>
    </template>
    <EchartsUI v-if="loaded" ref="chartRef" style="height: 260px" />
    <a-empty
      v-else
      :description="errorMsg || $t('page.dashboard.widget.loadFailed')"
      :image-style="{ height: '80px' }"
      style="padding: 40px 0"
    />
  </Card>
</template>
