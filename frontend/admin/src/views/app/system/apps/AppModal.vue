<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="getTitle" @ok="handleSubmit">
    <BasicForm @register="registerForm" />
  </BasicModal>
</template>

<script setup lang="ts">
  import { ref, computed, unref } from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { accountFormSchema } from './data';
  import { CreateApplication, UpdateApplication } from '/@/api/app/app';

  const isUpdate = ref(true);
  const rowId = ref('');

  const emits = defineEmits(['success', 'register']);

  const [registerForm, { setFieldsValue, updateSchema, resetFields, validate }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 24 },
    schemas: accountFormSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 23,
    },
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    await resetFields();

    setModalProps({ confirmLoading: false });
    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      rowId.value = data.record.id;
      await setFieldsValue({
        ...data.record,
      });
    }

    await updateSchema([
      {
        field: 'password',
        required: !unref(isUpdate),
      },
    ]);
  });

  const getTitle = computed(() => (!unref(isUpdate) ? '新增应用' : '编辑应用'));

  async function handleSubmit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });

      const _isUpdate = unref(isUpdate);
      const _rowId = unref(rowId);

      // API提交更改
      if (_isUpdate) {
        if (values.password == '') {
          values.password = undefined;
        }
        await UpdateApplication({ id: Number(_rowId), app: values });
      } else {
        await CreateApplication({ app: values });
      }

      closeModal();
      emits('success', { isUpdate: _isUpdate, values: { ...values, id: _rowId } });
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
