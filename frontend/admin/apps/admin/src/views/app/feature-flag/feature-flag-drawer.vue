<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { useCreateFeatureFlag, useUpdateFeatureFlag } from '#/api';

const { mutateAsync: createFlag } = useCreateFeatureFlag();
const { mutateAsync: updateFlag } = useUpdateFeatureFlag();

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('page.featureFlag.button.create')
    : $t('page.featureFlag.button.update'),
);

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'key',
      label: $t('page.featureFlag.key'),
      rules: 'required',
      help: $t('page.featureFlag.keyHelp'),
      componentProps: {
        placeholder: 'new-checkout-flow',
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.featureFlag.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'description',
      label: $t('page.featureFlag.description'),
      componentProps: {
        rows: 2,
        placeholder: $t('page.featureFlag.descriptionPlaceholder'),
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'appId',
      label: $t('page.featureFlag.appId'),
      componentProps: {
        placeholder: $t('page.featureFlag.appIdPlaceholder'),
        class: 'w-full',
        min: 0,
        precision: 0,
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'rolloutPercentage',
      label: $t('page.featureFlag.rolloutPercentage'),
      defaultValue: 100,
      rules: 'required',
      componentProps: {
        class: 'w-full',
        min: 0,
        max: 100,
        precision: 0,
        addonAfter: '%',
      },
    },
    {
      component: 'Input',
      fieldName: 'goalEventName',
      label: $t('page.featureFlag.goalEventName'),
      help: $t('page.featureFlag.goalEventHelp'),
      componentProps: {
        placeholder: 'pay_success',
        allowClear: true,
      },
    },
    {
      component: 'Switch',
      fieldName: 'enabled',
      label: $t('page.featureFlag.enabled'),
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

    // 变体在表格行内编辑（简易 JSON），这里从行数据透传
    const values = await baseFormApi.getValues();
    const payload = {
      ...values,
      variants: data.value?.row?.variants || [
        { key: 'control', name: '对照组', weight: 500 },
        { key: 'treatment', name: '实验组', weight: 500 },
      ],
    };

    try {
      await (data.value?.create
        ? createFlag(payload)
        : updateFlag({ id: data.value.row.id, values: payload }));

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
