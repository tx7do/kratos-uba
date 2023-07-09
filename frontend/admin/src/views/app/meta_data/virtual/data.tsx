import { BasicColumn } from '/@/components/Table';

export const eventColumns: BasicColumn[] = [
  {
    title: '属性名',
    dataIndex: 'name',
    sorter: false,
  },
  {
    title: '属性显示名',
    dataIndex: 'show_name',
    sorter: false,
  },
  {
    title: '数据类型',
    dataIndex: 'dataType',
    sorter: false,
  },
  {
    title: '上报数据',
    dataIndex: 'has_data',
    sorter: false,
  },
];
