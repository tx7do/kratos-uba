<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { useCreateDashboard, useUpdateDashboard } from '#/api';

const { mutateAsync: createDashboard } = useCreateDashboard();
const { mutateAsync: updateDashboard } = useUpdateDashboard();

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('page.dashboard.button.create')
    : $t('page.dashboard.button.update'),
);

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.dashboard.name'),
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'description',
      label: $t('page.dashboard.description'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'appId',
      label: $t('page.dashboard.appId'),
      componentProps: {
        placeholder: $t('page.dashboard.appIdPlaceholder'),
        class: 'w-full',
        min: 0,
        precision: 0,
      },
    },
    {
      component: 'Switch',
      fieldName: 'enabled',
      label: $t('page.dashboard.enabled'),
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
    // 新建看板不带组件，组件在看板视图页内配置维护。
    const payload = { ...values, widgets: data.value?.row?.widgets || [] };

    try {
      await (data.value?.create
        ? createDashboard(payload)
        : updateDashboard({ id: data.value.row.id, values: payload }));

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
  <Drawer :title="getTitle">
    <BaseForm class="mx-4" />
  </Drawer>
</template>
