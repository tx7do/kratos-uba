<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { useCreateAiProvider, useUpdateAiProvider } from '#/api';
import { makeUpdateMask } from '#/transport/rest';

const { mutateAsync: createProvider } = useCreateAiProvider();
const { mutateAsync: updateProvider } = useUpdateAiProvider();

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('page.ai.provider.button.create')
    : $t('page.ai.provider.button.update'),
);

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
  },
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.ai.provider.name'),
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'RadioGroup',
      fieldName: 'modelType',
      defaultValue: 2,
      label: $t('page.ai.provider.modelType'),
      rules: 'required',
      componentProps: {
        options: [
          { label: $t('page.ai.provider.modelTypeCloud'), value: 2 },
          { label: $t('page.ai.provider.modelTypeLocal'), value: 1 },
        ],
      },
    },
    {
      component: 'Input',
      fieldName: 'modelName',
      label: $t('page.ai.provider.modelName'),
      rules: 'required',
      componentProps: {
        placeholder: $t('page.ai.provider.modelNamePlaceholder'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'baseUrl',
      label: $t('page.ai.provider.baseUrl'),
      dependencies: {
        triggerFields: ['modelType'],
        if: () => true,
      },
      componentProps: {
        placeholder: 'https://api.deepseek.com/v1',
        allowClear: true,
      },
      help: $t('page.ai.provider.baseUrlHelp'),
    },
    {
      component: 'Input',
      fieldName: 'organization',
      label: $t('page.ai.provider.organization'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'InputPassword',
      fieldName: 'apiKey',
      label: $t('page.ai.provider.apiKey'),
      componentProps: {
        placeholder: $t('page.ai.provider.apiKeyPlaceholder'),
        allowClear: true,
        autocomplete: 'new-password',
      },
      help: $t('page.ai.provider.apiKeyHelp'),
    },
    {
      component: 'Input',
      fieldName: 'localHost',
      label: $t('page.ai.provider.localHost'),
      componentProps: {
        placeholder: 'localhost',
        allowClear: true,
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'localPort',
      label: $t('page.ai.provider.localPort'),
      componentProps: {
        placeholder: '11434',
        min: 1,
        max: 65_535,
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'timeoutSeconds',
      label: $t('page.ai.provider.timeoutSeconds'),
      defaultValue: 30,
      componentProps: {
        min: 1,
        max: 600,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'systemPrompt',
      label: $t('page.ai.provider.systemPrompt'),
      componentProps: {
        placeholder: $t('page.ai.provider.systemPromptPlaceholder'),
        rows: 3,
      },
    },
    {
      component: 'Switch',
      fieldName: 'isDefault',
      defaultValue: false,
      label: $t('page.ai.provider.isDefault'),
      componentProps: {
        class: 'w-auto',
      },
    },
    {
      component: 'Switch',
      fieldName: 'isEnabled',
      defaultValue: true,
      label: $t('page.ai.provider.isEnabled'),
      componentProps: {
        class: 'w-auto',
      },
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      label: $t('ui.table.remark'),
      componentProps: {
        rows: 2,
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
        ? createProvider({ data: values as never })
        : updateProvider({
            id: data.value.row.id,
            data: values as never,
            // apiKey 仅在用户填写时才参与更新（空值 = 不改密钥）
            updateMask: makeUpdateMask([
              'name',
              'modelType',
              'modelName',
              'baseUrl',
              'organization',
              'localHost',
              'localPort',
              'timeoutSeconds',
              'systemPrompt',
              'isDefault',
              'isEnabled',
              'remark',
              ...(values.apiKey ? ['apiKey'] : []),
            ]),
          }));

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

  onOpenChange(isOpen: boolean) {
    if (isOpen) {
      data.value = drawerApi.getData<Record<string, any>>();
      if (!data.value?.create && data.value?.row) {
        // 密钥不回显（后端从不返回明文），编辑时空值表示不修改
        const { apiKey: _apiKey, ...rest } = data.value.row;
        baseFormApi.setValues(rest);
      }
    }
  },
});

function setLoading(loading: boolean) {
  drawerApi.setState({ confirmLoading: loading });
}
</script>

<template>
  <Drawer class="w-[560px]" :title="getTitle">
    <BaseForm />
  </Drawer>
</template>
