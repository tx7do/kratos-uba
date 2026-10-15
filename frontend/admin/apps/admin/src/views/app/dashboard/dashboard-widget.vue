<script lang="ts" setup>
import type {
  ubaservicev1_DashboardWidget,
  ubaservicev1_EventTrendResponse,
  ubaservicev1_GroupByResponse,
} from '#/generated/api/admin/service/v1';

import { onMounted, ref } from 'vue';

import {
  EchartsUI,
  type EchartsUIType,
  useEcharts,
} from '@vben/plugins/echarts';

import { Card } from 'ant-design-vue';

import { fetchActiveUsers, fetchEventTrend, fetchGroupBy } from '#/api';

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
