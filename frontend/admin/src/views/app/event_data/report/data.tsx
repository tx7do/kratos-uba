import { BasicColumn } from '/@/components/Table';

export const reportColumns: BasicColumn[] = [
  {
    title: '数据名称',
    dataIndex: 'name',
    sorter: false,
  },
  {
    title: '显示名',
    dataIndex: 'show_name',
    sorter: false,
  },
  {
    title: '已接收',
    dataIndex: 'recv_count',
    sorter: false,
  },
  {
    title: '已入库',
    dataIndex: 'save_count',
    sorter: false,
  },
  {
    title: '异常入库',
    dataIndex: 'failed_count',
    sorter: false,
  },
  {
    title: '无法入库',
    dataIndex: 'cant_count',
    sorter: false,
  },
];
