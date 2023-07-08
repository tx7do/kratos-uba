import { BasicColumn, FormSchema } from '/@/components/Table';
import { SwitchStatusEnum } from '/@/enums/httpEnum';
import { Tag } from 'ant-design-vue';
import { h } from 'vue';

export const columns: BasicColumn[] = [
  {
    title: '应用名',
    dataIndex: 'name',
    sorter: false,
  },
  {
    title: 'AppId',
    dataIndex: 'appId',
    sorter: false,
  },
  {
    title: 'AppKey',
    dataIndex: 'appKey',
    sorter: false,
  },
  {
    title: '描述',
    dataIndex: 'remark',
    sorter: false,
  },
  {
    title: '状态',
    dataIndex: 'status',
    sorter: false,
    width: 80,
    customRender: ({ record }) => {
      const { status } = record as Application;
      const enable = status === SwitchStatusEnum.ON;
      const color = enable ? '#108ee9' : '#FF0000';
      const text = enable ? '启用' : '停用';
      return h(Tag, { color: color }, () => text);
    },
  },
  {
    title: '数据保留月数',
    dataIndex: 'keepMonth',
    sorter: false,
  },
  {
    title: '创建时间',
    dataIndex: 'createTime',
    sorter: false,
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'name',
    label: '应用名',
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'appId',
    label: 'AppId',
    colProps: { span: 8 },
    component: 'Input',
  },
];

export const accountFormSchema: FormSchema[] = [
  {
    field: 'name',
    label: '应用名',
    component: 'Input',
    rules: [
      {
        required: true,
        message: '请输入应用名',
      },
    ],
  },
  {
    field: 'remark',
    label: '应用描述',
    component: 'Input',
  },
  {
    field: 'keepMonth',
    label: '数据保留月数',
    component: 'InputNumber',
    required: true,
  },
];
