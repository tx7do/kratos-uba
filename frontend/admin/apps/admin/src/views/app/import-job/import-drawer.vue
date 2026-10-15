<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { useCreateImportJob } from '#/api';

const { mutateAsync: createJob } = useCreateImportJob();

const data = ref();
const fileContent = ref('');

const getTitle = computed(() => $t('page.importJob.button.create'));

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.importJob.name'),
      rules: 'required',
      componentProps: {
        placeholder: 'import_20261015',
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'appId',
      label: $t('page.importJob.appId'),
      componentProps: {
        placeholder: 'demo_app_001',
        allowClear: true,
      },
    },
  ],
});

function handleFileChange(event: Event) {
  const input = event.target as HTMLInputElement;
  const file = input.files?.[0];
  if (!file) return;

  const reader = new FileReader();
  reader.onload = () => {
    fileContent.value = reader.result as string;
  };
  reader.readAsText(file);
}

const [Drawer, drawerApi] = useVbenDrawer({
  onCancel() {
    drawerApi.close();
  },

  async onConfirm() {
    const validate = await baseFormApi.validate();
    if (!validate.valid) {
      return;
    }

    const values = await baseFormApi.getValues();
    const eventsJson = fileContent.value || (values.eventsJson as string) || '';
    if (!eventsJson.trim()) {
      notification.error({ message: $t('page.importJob.emptyEvents') });
      return;
    }

    // 校验 JSON 格式
    try {
      const parsed = JSON.parse(eventsJson);
      if (!Array.isArray(parsed)) {
        notification.error({ message: $t('page.importJob.notArray') });
        return;
      }
    } catch {
      notification.error({ message: $t('page.importJob.invalidJson') });
      return;
    }

    drawerApi.setState({ loading: true });

    try {
      await createJob({
        name: values.name,
        appId: values.appId,
        eventsJson,
      });
      notification.success({ message: $t('ui.notification.create_success') });
    } catch {
      notification.error({ message: $t('ui.notification.create_failed') });
    } finally {
      drawerApi.close();
      drawerApi.setState({ loading: false });
    }
  },

  onOpenChange(isOpen) {
    if (isOpen) {
      data.value = drawerApi.getData<Record<string, any>>();
      baseFormApi.setValues({});
      fileContent.value = '';
      drawerApi.setState({ loading: false });
    }
  },
});
</script>

<template>
  <Drawer :title="getTitle" class="w-[560px]">
    <BaseForm class="mx-4" />
    <div class="mx-4 mt-4">
      <label class="mb-2 block text-sm font-medium">
        {{ $t('page.importJob.eventsFile') }}
      </label>
      <input
        type="file"
        accept=".json"
        class="block w-full text-sm file:mr-2 file:rounded file:border-0 file:px-3 file:py-1 file:text-sm file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100"
        @change="handleFileChange"
      />
      <p class="text-muted-foreground mt-2 text-xs">
        {{ $t('page.importJob.fileHint') }}
      </p>
    </div>
  </Drawer>
</template>
