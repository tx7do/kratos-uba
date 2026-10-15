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

interface VariantRow {
  key: string;
  name: string;
  weight: number;
}

/** 变体在表单外独立维护（行内编辑），提交时并入 payload */
const variants = ref<VariantRow[]>([]);

const weightTotal = computed(() =>
  variants.value.reduce((sum, v) => sum + (v.weight || 0), 0),
);

const getTitle = computed(() =>
  data.value?.create
    ? $t('page.featureFlag.button.create')
    : $t('page.featureFlag.button.update'),
);

function defaultVariants(): VariantRow[] {
  return [
    { key: 'control', name: $t('page.featureFlag.controlGroup'), weight: 500 },
    { key: 'treatment', name: $t('page.featureFlag.treatmentGroup'), weight: 500 },
  ];
}

function addVariant() {
  variants.value.push({ key: '', name: '', weight: 0 });
}

function removeVariant(index: number) {
  variants.value.splice(index, 1);
}

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

    // 变体校验：key 非空且唯一、权重非负且总和大于 0
    const seen = new Set<string>();
    for (const variant of variants.value) {
      if (!variant.key.trim()) {
        notification.error({ message: $t('page.featureFlag.variantKeyRequired') });
        return;
      }
      if (seen.has(variant.key)) {
        notification.error({
          message: $t('page.featureFlag.variantKeyDuplicate', { key: variant.key }),
        });
        return;
      }
      seen.add(variant.key);
      if ((variant.weight || 0) < 0) {
        notification.error({ message: $t('page.featureFlag.variantWeightInvalid') });
        return;
      }
    }
    if (weightTotal.value <= 0) {
      notification.error({ message: $t('page.featureFlag.variantWeightSumZero') });
      return;
    }

    setLoading(true);

    const values = await baseFormApi.getValues();
    const payload = {
      ...values,
      variants: variants.value.map((v) => ({
        key: v.key.trim(),
        name: v.name?.trim(),
        weight: Math.round(v.weight || 0),
      })),
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

      const existing = data.value?.row?.variants;
      variants.value =
        existing && existing.length > 0
          ? existing.map((v: any) => ({
              key: v.key || '',
              name: v.name || '',
              weight: Number(v.weight) || 0,
            }))
          : defaultVariants();

      setLoading(false);
    }
  },
});

function setLoading(loading: boolean) {
  drawerApi.setState({ loading });
}
</script>

<template>
  <Drawer :title="getTitle" class="w-[640px]">
    <BaseForm class="mx-4" />

    <div class="mx-4 mt-4">
      <div class="mb-2 flex items-center justify-between">
        <span class="text-sm font-medium">
          {{ $t('page.featureFlag.variants') }}
        </span>
        <span class="text-muted-foreground text-xs">
          {{ $t('page.featureFlag.weightTotal') }}:
          {{ Math.round((weightTotal / 10) * 10) / 10 }}‰
        </span>
      </div>

      <div
        v-for="(variant, index) in variants"
        :key="index"
        class="mb-2 flex items-center gap-2"
      >
        <a-input
          v-model:value="variant.key"
          :placeholder="$t('page.featureFlag.variantKey')"
          class="w-36"
        />
        <a-input
          v-model:value="variant.name"
          :placeholder="$t('page.featureFlag.variantName')"
          class="w-32"
        />
        <a-input-number
          v-model:value="variant.weight"
          :min="0"
          :max="1000"
          :precision="0"
          :addon-after="'‰'"
          class="w-32"
        />
        <a-button
          danger
          size="small"
          type="link"
          :disabled="variants.length <= 1"
          @click="removeVariant(index)"
        >
          {{ $t('ui.button.delete') }}
        </a-button>
      </div>

      <a-button size="small" type="dashed" class="w-full" @click="addVariant">
        + {{ $t('page.featureFlag.addVariant') }}
      </a-button>
      <p class="text-muted-foreground mt-2 text-xs">
        {{ $t('page.featureFlag.weightHint') }}
      </p>
    </div>
  </Drawer>
</template>
