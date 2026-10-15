import type { App } from 'vue';

import {
  Alert,
  Badge,
  Button,
  Card,
  Col,
  Descriptions,
  Divider,
  Dropdown,
  Empty,
  Form,
  Input,
  InputNumber,
  Layout,
  Menu,
  Modal,
  Popconfirm,
  Progress,
  RadioButton,
  RadioGroup,
  RangePicker,
  Row,
  Select,
  Space,
  Spin,
  Statistic,
  Switch,
  Table,
  Tabs,
  Tag,
  Timeline,
  Tooltip,
  Tree,
} from 'ant-design-vue';

/**
 * 注册全局组件
 * @param app
 */
export function registerGlobComp(app: App) {
  app
    .use(Input)
    .use(Button)
    .use(Layout)
    .use(Space)
    .use(Card)
    .use(Switch)
    .use(Popconfirm)
    .use(Dropdown)
    .use(Tag)
    .use(Tabs)
    .use(Divider)
    .use(Menu)
    .use(Select)
    .use(Tree)
    .use(Table)
    .use(Statistic)
    .use(InputNumber)
    .use(Timeline)
    .use(Empty)
    .use(Badge)
    .use(Spin)
    .use(Alert)
    // Form/Descriptions 的 install 会连带注册 a-form-item / a-descriptions-item，
    // Table/Timeline/Select 的 install 同理会带上 a-table-column / a-timeline-item /
    // a-select-option，不要在下方重复 component 注册。
    .use(Form)
    .use(Descriptions)
    .use(Modal)
    .use(Progress)
    .use(Tooltip);

  // 以下组件是普通组件（非 Plugin，无 install 方法），需用 app.component 单独注册。
  // 注册名必须是组件自身 .name（带 A 前缀），模板里的 <a-range-picker> 等才能匹配。
  app
    .component('ARangePicker', RangePicker)
    .component('ARadioGroup', RadioGroup)
    .component('ARadioButton', RadioButton)
    .component('ARow', Row)
    .component('ACol', Col);
}
