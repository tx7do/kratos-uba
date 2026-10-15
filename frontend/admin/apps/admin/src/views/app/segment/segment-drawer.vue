<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { useCreateSegment, useUpdateSegment } from '#/api';

const { mutateAsync: createSegment } = useCreateSegment();
const { mutateAsync: updateSegment } = useUpdateSegment();

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('page.segment.button.create')
    : $t('page.segment.button.update'),
);

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.segment.name'),
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'description',
      label: $t('page.segment.description'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'appId',
      label: $t('page.segment.appId'),
      componentProps: {
        placeholder: $t('page.segment.appIdPlaceholder'),
        class: 'w-full',
        min: 0,
        precision: 0,
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'refreshIntervalMinutes',
      label: $t('page.segment.refreshIntervalMinutes'),
      defaultValue: 0,
      help: $t('page.segment.refreshIntervalHelp'),
      componentProps: {
        class: 'w-full',
        min: 0,
        precision: 0,
        addonAfter: 'min',
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'maxUsers',
      label: $t('page.segment.maxUsers'),
      defaultValue: 50_000,
      componentProps: {
        class: 'w-full',
        min: 1,
        precision: 0,
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'lookbackDays',
      label: $t('page.segment.lookbackDays'),
      defaultValue: 30,
      help: $t('page.segment.lookbackHelp'),
      componentProps: {
        class: 'w-full',
        min: 1,
        precision: 0,
        addonAfter: 'd',
      },
    },
    {
      component: 'Select',
      fieldName: 'includeEvent',
      label: $t('page.segment.includeEvent'),
      rules: 'required',
      componentProps: {
        class: 'w-full',
        placeholder: $t('ui.placeholder.select'),
        allowClear: true,
        showSearch: true,
        mode: 'tags',
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'includeMinTimes',
      label: $t('page.segment.includeMinTimes'),
      defaultValue: 1,
      componentProps: {
        class: 'w-full',
        min: 1,
        precision: 0,
      },
    },
    {
      component: 'Select',
      fieldName: 'excludeEvent',
      label: $t('page.segment.excludeEvent'),
      componentProps: {
        class: 'w-full',
        placeholder: $t('ui.placeholder.select'),
        allowClear: true,
        showSearch: true,
        mode: 'tags',
      },
    },
    {
      component: 'Switch',
      fieldName: 'enabled',
      label: $t('page.segment.enabled'),
      defaultValue: true,
    },
  ],
});

const [Drawer, drawerApi] = useVbenDrawer({
  onCancel() {
    drawerApi.close();
  },

  async onConfirm() {
    const validate = await baseFormApi.validate();
    if (!validate.valid) {
      return;
    }

    setLoading(true);

    const values = await baseFormApi.getValues();
    const payload = toDefinitionPayload(values);

    try {
      await (data.value?.create
        ? createSegment({ ...payload, refreshOnCreate: true })
        : updateSegment({ id: data.value.row.id, values: payload }));

      notification.success({
        message: data.value?.create
          ? $t('ui.notification.create_success')
          : $t('ui.notification.update_success'),
      });
    } catch {
      notification.error({
        message: data.value?.create
          ? $t('ui.notification.create_failed')
          : $t('ui.notification.update_failed'),
      });
    } finally {
      drawerApi.close();
      setLoading(false);
    }
  },

  onOpenChange(isOpen) {
    if (isOpen) {
      data.value = drawerApi.getData<Record<string, any>>();
      baseFormApi.setValues(
        (fromDefinitionRow(data.value?.row) || {}) as Record<string, any>,
      );
      setLoading(false);
    }
  },
});

// 把表单值折叠成 Segment 契约（definition 为 SegmentationDefinition 快照）。
function toDefinitionPayload(values: Record<string, any>) {
  const include = (values.includeEvent || []).map((eventName: string) => ({
    eventName,
    minTimes: values.includeMinTimes || 1,
  }));
  const exclude = (values.excludeEvent || []).map((eventName: string) => ({
    eventName,
    minTimes: 1,
  }));

  const endMs = Date.now();
  const startMs =
    endMs - Math.max(1, values.lookbackDays || 30) * 24 * 60 * 60 * 1000;

  return {
    name: values.name,
    description: values.description,
    appId: values.appId,
    enabled: values.enabled,
    refreshIntervalMinutes: values.refreshIntervalMinutes,
    maxUsers: values.maxUsers,
    definition: {
      timeRange: { startMs, endMs },
      include,
      exclude,
      limit: values.maxUsers,
    },
  };
}

// 把行数据展开回表单值。
function fromDefinitionRow(row: Record<string, any> | undefined) {
  if (!row) {
    return row;
  }

  const def = row.definition || {};
  const lookbackDays = def.timeRange?.endMs
    ? Math.max(
        1,
        Math.round(
          (def.timeRange.endMs - (def.timeRange.startMs || 0)) / 86_400_000,
        ),
      )
    : 30;

  return {
    name: row.name,
    description: row.description,
    appId: row.appId,
    enabled: row.enabled,
    refreshIntervalMinutes: row.refreshIntervalMinutes,
    maxUsers: row.maxUsers,
    lookbackDays,
    includeEvent: (def.include || []).map((c: any) => c.eventName),
    includeMinTimes: def.include?.[0]?.minTimes || 1,
    excludeEvent: (def.exclude || []).map((c: any) => c.eventName),
  };
}

function setLoading(loading: boolean) {
  drawerApi.setState({ loading });
}
</script>

<template>
  <Drawer :title="getTitle" class="w-[560px]">
    <BaseForm class="mx-4" />
  </Drawer>
</template>
