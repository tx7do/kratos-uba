<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  fetchListAiProviders,
  useCreateAiKnowledgeBase,
  useUpdateAiKnowledgeBase,
} from '#/api';
import { PaginationQuery } from '#/transport/rest';

const { mutateAsync: createBase } = useCreateAiKnowledgeBase();
const { mutateAsync: updateBase } = useUpdateAiKnowledgeBase();

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('page.ai.knowledge.button.create')
    : $t('page.ai.knowledge.button.update'),
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
      label: $t('page.ai.knowledge.name'),
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'description',
      label: $t('page.ai.knowledge.description'),
      componentProps: {
        rows: 2,
      },
    },
    {
      component: 'Select',
      fieldName: 'providerId',
      label: $t('page.ai.knowledge.providerId'),
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        options: [] as Array<{ label: string; value: number }>,
      },
      help: $t('page.ai.knowledge.providerIdHelp'),
    },
    {
      component: 'Input',
      fieldName: 'embeddingModel',
      label: $t('page.ai.knowledge.embeddingModel'),
      rules: 'required',
      componentProps: {
        placeholder: 'text-embedding-3-small',
        allowClear: true,
      },
      help: $t('page.ai.knowledge.embeddingModelHelp'),
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
        ? createBase({ data: values as never })
        : updateBase({ id: data.value.row.id, data: values as never }));

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
        baseFormApi.setValues(data.value.row);
      }
    }
  },
});

function setLoading(loading: boolean) {
  drawerApi.setState({ confirmLoading: loading });
}

onMounted(async () => {
  // 提供商下拉：全量拉取启用的提供商（数量有限）
  try {
    const resp = await fetchListAiProviders(
      new PaginationQuery({ paging: { pageSize: 100 } }),
    );
    const options = (resp.items ?? []).map((p) => ({
      label: `${p.name ?? ''}（${p.modelName ?? ''}）`,
      value: p.id as number,
    }));
    baseFormApi.updateSchema([
      {
        componentProps: {
          placeholder: $t('ui.placeholder.select'),
          options,
        },
        fieldName: 'providerId',
      },
    ]);
  } catch {
    // 提供商加载失败不阻断表单：提交时后端仍会校验 provider 存在
  }
});
</script>

<template>
  <Drawer class="w-[520px]" :title="getTitle">
    <BaseForm />
  </Drawer>
</template>
