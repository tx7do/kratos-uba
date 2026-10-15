<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { useCreateAlertRule, useUpdateAlertRule } from '#/api';

const { mutateAsync: createAlertRule } = useCreateAlertRule();
const { mutateAsync: updateAlertRule } = useUpdateAlertRule();

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('page.alertRule.button.create')
    : $t('page.alertRule.button.update'),
);

const aggregatorOptions = [
  { label: $t('page.alertRule.aggregator.count'), value: 'COUNT' },
  { label: $t('page.alertRule.aggregator.activeUsers'), value: 'ACTIVE_USERS' },
  { label: $t('page.alertRule.aggregator.sumAmount'), value: 'SUM_AMOUNT' },
  { label: $t('page.alertRule.aggregator.avgAmount'), value: 'AVG_AMOUNT' },
];

const comparisonOptions = [
  { label: '>', value: 'GT' },
  { label: '>=', value: 'GTE' },
  { label: '<', value: 'LT' },
  { label: '<=', value: 'LTE' },
];

const channelOptions = [
  {
    label: $t('page.alertRule.channel.internalMessage'),
    value: 'INTERNAL_MESSAGE',
  },
  { label: $t('page.alertRule.channel.webhook'), value: 'WEBHOOK' },
];

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.alertRule.name'),
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'description',
      label: $t('page.alertRule.description'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'appId',
      label: $t('page.alertRule.appId'),
      componentProps: {
        placeholder: $t('page.alertRule.appIdPlaceholder'),
        class: 'w-full',
        min: 0,
        precision: 0,
      },
    },
    {
      component: 'Input',
      fieldName: 'eventName',
      label: $t('page.alertRule.eventName'),
      componentProps: {
        placeholder: $t('page.alertRule.eventNamePlaceholder'),
        allowClear: true,
      },
    },
    {
      component: 'Select',
      fieldName: 'aggregator',
      label: $t('page.alertRule.aggregator.label'),
      defaultValue: 'COUNT',
      componentProps: {
        class: 'w-full',
        options: aggregatorOptions,
        placeholder: $t('ui.placeholder.select'),
      },
    },
    {
      component: 'Select',
      fieldName: 'comparison',
      label: $t('page.alertRule.comparison'),
      defaultValue: 'GT',
      componentProps: {
        class: 'w-full',
        options: comparisonOptions,
        placeholder: $t('ui.placeholder.select'),
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'threshold',
      label: $t('page.alertRule.threshold'),
      defaultValue: 0,
      rules: 'required',
      componentProps: {
        class: 'w-full',
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'windowMinutes',
      label: $t('page.alertRule.windowMinutes'),
      defaultValue: 60,
      componentProps: {
        class: 'w-full',
        min: 1,
        precision: 0,
        addonAfter: 'min',
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'silenceMinutes',
      label: $t('page.alertRule.silenceMinutes'),
      defaultValue: 60,
      componentProps: {
        class: 'w-full',
        min: 0,
        precision: 0,
        addonAfter: 'min',
      },
    },
    {
      component: 'Select',
      fieldName: 'channels',
      label: $t('page.alertRule.channels'),
      defaultValue: ['INTERNAL_MESSAGE'],
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
      label: $t('page.alertRule.enabled'),
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

    try {
      await (data.value?.create
        ? createAlertRule(values)
        : updateAlertRule({ id: data.value.row.id, values }));

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
