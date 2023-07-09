import { BasicColumn } from '/@/components/Table';

export const metaColumns: BasicColumn[] = [
  {
    title: '事件名',
    dataIndex: 'name',
    sorter: false,
  },
  {
    title: '显示状态',
    dataIndex: 'status',
    sorter: false,
  },
  {
    title: '是否接收',
    dataIndex: 'enable_recv',
    sorter: false,
  },
  {
    title: '过去30天入库',
    dataIndex: 'recv_count',
    sorter: false,
  },
];
