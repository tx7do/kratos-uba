<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  fetchListSegments,
  PaginationQuery,
  useCreateCampaign,
  useUpdateCampaign,
} from '#/api';

const { mutateAsync: createCampaign } = useCreateCampaign();
const { mutateAsync: updateCampaign } = useUpdateCampaign();

const data = ref();
const segmentOptions = ref<{ label: string; value: number }[]>([]);

const getTitle = computed(() =>
  data.value?.create
    ? $t('page.campaign.button.create')
    : $t('page.campaign.button.update'),
);

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.campaign.name'),
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'description',
      label: $t('page.campaign.description'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'appId',
      label: $t('page.campaign.appId'),
      componentProps: {
        placeholder: $t('page.campaign.appIdPlaceholder'),
        class: 'w-full',
        min: 0,
        precision: 0,
      },
    },
    {
      component: 'Select',
      fieldName: 'segmentId',
      label: $t('page.campaign.segmentId'),
      rules: 'required',
      componentProps: {
        class: 'w-full',
        options: segmentOptions,
        placeholder: $t('ui.placeholder.select'),
        showSearch: true,
        optionFilterProp: 'label',
      },
    },
    {
      component: 'Input',
      fieldName: 'title',
      label: $t('page.campaign.title'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'content',
      label: $t('page.campaign.content'),
      componentProps: {
        rows: 5,
        placeholder: $t('page.campaign.contentPlaceholder'),
      },
    },
    {
      component: 'Input',
      fieldName: 'templateCode',
      label: $t('page.campaign.templateCode'),
      help: $t('page.campaign.templateHelp'),
      componentProps: {
        placeholder: 'campaign_promo',
        allowClear: true,
      },
    },
    {
      component: 'Select',
      fieldName: 'sendMode',
      label: $t('page.campaign.sendMode'),
      defaultValue: 'CAMPAIGN_SEND_MODE_MANUAL',
      componentProps: {
        class: 'w-full',
        options: [
          { label: $t('page.campaign.mode.manual'), value: 'CAMPAIGN_SEND_MODE_MANUAL' },
          { label: $t('page.campaign.mode.scheduled'), value: 'CAMPAIGN_SEND_MODE_SCHEDULED' },
        ],
      },
    },
    {
      component: 'DatePicker',
      fieldName: 'scheduledAt',
      label: $t('page.campaign.scheduledAt'),
      dependencies: {
        triggerFields: ['sendMode'],
        if: (values) => values.sendMode === 'CAMPAIGN_SEND_MODE_SCHEDULED',
      },
      componentProps: {
        class: 'w-full',
        showTime: true,
        format: 'YYYY-MM-DD HH:mm:ss',
        placeholder: $t('ui.placeholder.select'),
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

    // scheduledAt: DatePicker dayjs → RFC3339
    const payload: Record<string, any> = { ...values };
    if (values.scheduledAt?.toISOString) {
      payload.scheduledAt = values.scheduledAt.toISOString();
    }
    if (values.sendMode !== 'CAMPAIGN_SEND_MODE_SCHEDULED') {
      delete payload.scheduledAt;
    }
    payload.status = data.value?.create ? 'DRAFT' : values.status;

    try {
      await (data.value?.create
        ? createCampaign(payload)
        : updateCampaign({ id: data.value.row.id, values: payload }));

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
      const row = { ...(data.value?.row || {}) };
      // scheduledAt: RFC3339 字符串 → dayjs 由 DatePicker 接收
      baseFormApi.setValues(row);
      setLoading(false);

      // 拉分群选项（不分页全量，数量有限）
      fetchListSegments(
        new PaginationQuery({
          paging: { page: 1, pageSize: 200 },
        }),
      )
        .then((resp) => {
          segmentOptions.value = (resp.items || []).map((seg: any) => ({
            label: seg.name,
            value: seg.id,
          }));
        })
        .catch(() => {
          segmentOptions.value = [];
        });
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
