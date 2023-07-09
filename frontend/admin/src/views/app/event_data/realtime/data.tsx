import RealtimeDataPage from './realtime-data-page.vue';
import ErrorDataPage from './error-data-page.vue';
import { BasicColumn } from '/@/components/Table';

export interface TabItem {
  key: string;
  name: string;
  component: any;
}

export const tabList: TabItem[] = [
  {
    key: '1',
    name: '实时入库',
    component: RealtimeDataPage,
  },
  {
    key: '2',
    name: '错误数据',
    component: ErrorDataPage,
  },
];

export const realtimeColumns: BasicColumn[] = [
  {
    title: '上报时间',
    dataIndex: 'createTime',
    sorter: false,
  },
  {
    title: '数据名称',
    dataIndex: 'content',
    sorter: false,
  },
  {
    title: '抽样事例',
    dataIndex: 'status',
    sorter: false,
  },
];

export const errorColumns: BasicColumn[] = [
  {
    title: '采样日期',
    dataIndex: 'createTime',
    sorter: false,
  },
  {
    title: '采用时间',
    dataIndex: 'content',
    sorter: false,
  },
  {
    title: '错误条数',
    dataIndex: 'status',
    sorter: false,
  },
  {
    title: '错误处理',
    dataIndex: 'status',
    sorter: false,
  },
  {
    title: '错误原因',
    dataIndex: 'status',
    sorter: false,
  },
  {
    title: '抽样事例',
    dataIndex: 'status',
    sorter: false,
  },
];
