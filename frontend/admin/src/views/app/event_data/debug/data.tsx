import { BasicColumn } from '/@/components/Table';

export const debugColumns: BasicColumn[] = [
  {
    title: '上报时间',
    dataIndex: 'createTime',
    sorter: false,
  },
  {
    title: '数据名称',
    dataIndex: 'name',
    sorter: false,
  },
  {
    title: '数据判断',
    dataIndex: 'status',
    sorter: false,
  },
  {
    title: '错误原因',
    dataIndex: 'reason',
    sorter: false,
  },
  {
    title: '数据明细',
    dataIndex: 'detail',
    sorter: false,
  },
];
