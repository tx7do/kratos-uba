<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { useCreateVisualEvent, useUpdateVisualEvent } from '#/api';

const { mutateAsync: createEvent } = useCreateVisualEvent();
const { mutateAsync: updateEvent } = useUpdateVisualEvent();

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('page.visualEvent.button.create')
    : $t('page.visualEvent.button.update'),
);

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'eventName',
      label: $t('page.visualEvent.eventName'),
      rules: 'required',
      componentProps: {
        placeholder: 'buy_button_click',
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'displayName',
      label: $t('page.visualEvent.displayName'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'appId',
      label: $t('page.visualEvent.appId'),
      componentProps: {
        placeholder: $t('page.visualEvent.appIdPlaceholder'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'pageUrl',
      label: $t('page.visualEvent.pageUrl'),
      help: $t('page.visualEvent.pageUrlHelp'),
      componentProps: {
        placeholder: 'https://example.com/product/',
        allowClear: true,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'elementXpath',
      label: $t('page.visualEvent.elementXpath'),
      rules: 'required',
      help: $t('page.visualEvent.elementXpathHelp'),
      componentProps: {
        rows: 3,
        class: 'font-mono',
        placeholder: '/html[1]/body[1]/div[1]/button[1]',
      },
    },
    {
      component: 'Select',
      fieldName: 'matchType',
      label: $t('page.visualEvent.matchType'),
      defaultValue: 'XPATH_EXACT',
      componentProps: {
        class: 'w-full',
        options: [
          { label: $t('page.visualEvent.match.exact'), value: 'XPATH_EXACT' },
          { label: $t('page.visualEvent.match.prefix'), value: 'XPATH_PREFIX' },
        ],
      },
    },
    {
      component: 'Switch',
      fieldName: 'enabled',
      label: $t('page.visualEvent.enabled'),
      defaultValue: true,
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      label: $t('page.visualEvent.remark'),
      componentProps: {
        rows: 2,
        placeholder: $t('ui.placeholder.input'),
      },
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
        ? createEvent(values)
        : updateEvent({ id: data.value.row.id, values }));

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
      baseFormApi.setValues({
        ...data.value?.row,
        matchType: data.value?.row?.matchType || 'XPATH_EXACT',
      });
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
