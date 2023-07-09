<template>
  <PageWrapper dense contentFullHeight>
    <BasicTable @register="registerTable">
      <template #bodyCell="{ column, record }">
        <TableAction
          v-if="(column as BasicColumn).dataIndex === 'action'"
          :actions="createActions(record)"
        />
      </template>
    </BasicTable>
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

  import { metaColumns } from './data';
  import { ListApplication } from '/@/api/app/app';

  const [registerTable] = useTable({
    api: ListApplication,
    columns: metaColumns,
    rowKey: 'id',
    showIndexColumn: false,
    canResize: false,
    useSearchForm: true,
    showTableSetting: false,
    bordered: true,
    actionColumn: {
      width: 100,
      title: '操作',
      dataIndex: 'action',
    },
  });

  function createActions(record: EditRecordRow): ActionItem[] {
    return [];
  }
</script>

<style lang="less" scoped></style>
