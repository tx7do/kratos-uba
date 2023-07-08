import { BasicColumn, FormSchema } from '/@/components/Table';

export const columns: BasicColumn[] = [
  {
    title: '账户名',
    dataIndex: 'userName',
    sorter: false,
  },
  {
    title: '姓名',
    dataIndex: 'realName',
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
    field: 'realName',
    label: '姓名',
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'userName',
    label: '账户名',
    colProps: { span: 8 },
    component: 'Input',
  },
];

export const accountFormSchema: FormSchema[] = [
  {
    field: 'userName',
    label: '用户名',
    component: 'Input',
    rules: [
      {
        required: true,
        message: '请输入用户名',
      },
    ],
  },
  {
    field: 'password',
    label: '密码',
    component: 'InputPassword',
    required: false,
  },
  {
    field: 'realName',
    label: '姓名',
    component: 'Input',
    required: true,
  },
];
