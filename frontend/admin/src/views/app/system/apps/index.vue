<template>
  <PageWrapper dense contentFullHeight>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button preIcon="ant-design:plus" type="primary" @click="handleCreate"> 添加 </a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="(column as BasicColumn).dataIndex === 'action'">
          <TableAction :actions="createActions(record)" />
        </template>
      </template>
    </BasicTable>
    <AppModal @register="registerModal" @success="handleSuccess" />
  </PageWrapper>
</template>

<script lang="ts" setup>
  import {
    BasicTable,
    useTable,
    TableAction,
    ActionItem,
    EditRecordRow,
    BasicColumn,
  } from '/@/components/Table';
  import { PageWrapper } from '/@/components/Page';

  import { columns, searchFormSchema } from './data';
  import { DeleteApplication, ListApplication } from '/@/api/app/app';
  import AppModal from './AppModal.vue';
  import { useModal } from '/@/components/Modal';
  import { notification } from 'ant-design-vue';

  const [registerTable, { reload, updateTableDataRecord }] = useTable({
    title: '应用列表',
    api: ListApplication,
    columns,
    rowKey: 'id',
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
      autoSubmitOnEnter: true,
    },
    showIndexColumn: true,
    canResize: false,
    useSearchForm: true,
    showTableSetting: true,
    bordered: true,
    actionColumn: {
      width: 180,
      title: '操作',
      dataIndex: 'action',
    },
  });

  const [registerModal, { openModal }] = useModal();

  function createActions(record: EditRecordRow): ActionItem[] {
    return [
      {
        label: '编辑',
        onClick: handleEdit.bind(null, record),
      },
      {
        label: '删除',
        color: 'error',
        onClick: handleDelete.bind(null, record),
      },
    ];
  }

  function handleCreate() {
    openModal(true, {
      isUpdate: false,
    });
  }

  function handleEdit(record: Recordable) {
    console.log(record);
    openModal(true, {
      record,
      isUpdate: true,
    });
  }

  function handleDelete(record: Recordable) {
    console.log(record);
    const { id = 0 } = record;
    DeleteApplication({ id }).then(() => {
      notification.success({
        message: '删除成功',
      });
      reload();
    });
  }

  function handleSuccess({ isUpdate, values }) {
    if (isUpdate) {
      // 演示不刷新表格直接更新内部数据。
      // 注意：updateTableDataRecord要求表格的rowKey属性为string并且存在于每一行的record的keys中
      const result = updateTableDataRecord(values.id, values);
      console.log(result);
    } else {
      reload();
    }
  }
</script>
