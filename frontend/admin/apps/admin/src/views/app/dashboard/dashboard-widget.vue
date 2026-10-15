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
} from '#/api';

interface Props {
  widget: ubaservicev1_DashboardWidget;
}

const props = defineProps<Props>();

const chartRef = ref<EchartsUIType>();
const { renderEcharts } = useEcharts(chartRef);
const loaded = ref(false);

function parseParams(): Record<string, any> {
  try {
    return JSON.parse(props.widget.paramsJson || '{}');
  } catch {
    return {};
  }
}

// 时间桶格式化为坐标轴标签
function timeLabel(ts?: number) {
  if (!ts) return '';
  const d = new Date(ts);
  return `${d.getMonth() + 1}/${d.getDate()} ${String(d.getHours()).padStart(2, '0')}:00`;
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
      case 'retention': {
        const resp: ubaservicev1_RetentionResponse = await fetchRetention(
          params as any,
        );
        // 热力图：行=同期群日期，列=偏移天，值=留存率(%)
        const cohorts = resp.cohorts || [];
        const offsets = resp.offsetDays || [];
        const xLabels = offsets.map((d) => `+${d}d`);
        const yLabels = cohorts.map((c) => {
          const d = c.cohortDate ? new Date(c.cohortDate) : null;
          return d ? `${d.getMonth() + 1}/${d.getDate()}` : '-';
        });
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
                  p.value[2] == null ? '' : `${p.value[2]}%`,
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
  } catch {
    loaded.value = false;
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
    <a-empty v-else :image-style="{ height: '80px' }" style="padding: 40px 0" />
  </Card>
</template>
