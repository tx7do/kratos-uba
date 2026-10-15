<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { useCreateReportSubscription, useUpdateReportSubscription } from '#/api';

const { mutateAsync: createSubscription } = useCreateReportSubscription();
const { mutateAsync: updateSubscription } = useUpdateReportSubscription();

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('page.report.button.create')
    : $t('page.report.button.update'),
);

const modelOptions = [
  { label: $t('page.report.model.eventTrend'), value: 'event-trend' },
  { label: $t('page.report.model.activeUsers'), value: 'active-users' },
  { label: $t('page.report.model.groupBy'), value: 'group-by' },
  { label: $t('page.report.model.funnel'), value: 'funnel' },
  { label: $t('page.report.model.retention'), value: 'retention' },
];

const channelOptions = [
  { label: $t('page.report.channel.internalMessage'), value: 'REPORT_CHANNEL_INTERNAL_MESSAGE' },
  { label: $t('page.report.channel.webhook'), value: 'REPORT_CHANNEL_WEBHOOK' },
];

const DEFAULT_PARAMS: Record<string, string> = {
  'event-trend': JSON.stringify({ granularity: 'DAY' }),
  'active-users': JSON.stringify({}),
  'group-by': JSON.stringify({ dimension: 'platform' }),
  'funnel': JSON.stringify({ steps: [{ eventName: 'view_home' }, { eventName: 'pay_success' }] }),
  'retention': JSON.stringify({}),
};

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.report.name'),
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'description',
      label: $t('page.report.description'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'appId',
      label: $t('page.report.appId'),
      componentProps: {
        placeholder: $t('page.report.appIdPlaceholder'),
        class: 'w-full',
        min: 0,
        precision: 0,
      },
    },
    {
      component: 'Select',
      fieldName: 'model',
      label: $t('page.report.model.label'),
      rules: 'required',
      componentProps: {
        class: 'w-full',
        options: modelOptions,
        placeholder: $t('ui.placeholder.select'),
        onChange: (model: string) => {
          // 切模型时若参数还是默认模板，则同步换成新模型的默认参数
          void baseFormApi.getValues().then((current) => {
            if (
              !current?.paramsJson ||
              Object.values(DEFAULT_PARAMS).includes(String(current.paramsJson))
            ) {
              baseFormApi.setValues({
                paramsJson: DEFAULT_PARAMS[model] || '{}',
              });
            }
          });
        },
      },
    },
    {
      component: 'Textarea',
      fieldName: 'paramsJson',
      label: $t('page.report.paramsJson'),
      defaultValue: '{}',
      help: $t('page.report.paramsHelp'),
      componentProps: {
        rows: 6,
        placeholder: '{}',
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'intervalMinutes',
      label: $t('page.report.intervalMinutes'),
      defaultValue: 1440,
      rules: 'required',
      componentProps: {
        class: 'w-full',
        min: 60,
        precision: 0,
        addonAfter: 'min',
      },
    },
    {
      component: 'Select',
      fieldName: 'channels',
      label: $t('page.report.channels'),
      defaultValue: ['REPORT_CHANNEL_INTERNAL_MESSAGE'],
      componentProps: {
        class: 'w-full',
        mode: 'multiple',
        options: channelOptions,
        placeholder: $t('ui.placeholder.select'),
      },
    },
    {
      component: 'Switch',
      fieldName: 'enabled',
      label: $t('page.report.enabled'),
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

    // 校验 paramsJson 合法性，避免存进去引擎跑不了
    try {
      JSON.parse(values.paramsJson || '{}');
    } catch {
      notification.error({ message: $t('page.report.paramsInvalid') });
      setLoading(false);
      return;
    }

    try {
      await (data.value?.create
        ? createSubscription(values)
        : updateSubscription({ id: data.value.row.id, values }));

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
      baseFormApi.setValues(data.value?.row);
      setLoading(false);
    }
  },
});

function setLoading(loading: boolean) {
  drawerApi.setState({ loading });
}
</script>

<template>
  <Drawer :title="getTitle" class="w-[560px]">
    <BaseForm class="mx-4" />
  </Drawer>
</template>
