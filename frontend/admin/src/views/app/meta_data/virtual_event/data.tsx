import { BasicColumn } from '/@/components/Table';

export const eventColumns: BasicColumn[] = [
  {
    title: '虚拟事件名',
    dataIndex: 'name',
    sorter: false,
  },
  {
    title: '虚拟事件显示名',
    dataIndex: 'show_name',
    sorter: false,
  },
  {
    title: '标签',
    dataIndex: 'tag',
    sorter: false,
  },
  {
    title: '事件截图',
    dataIndex: 'snapshot',
    sorter: false,
  },
  {
    title: '包含事件',
    dataIndex: 'include_event',
    sorter: false,
  },
];
