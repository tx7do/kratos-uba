<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { useFormCloseGuard } from '#/composables/use-form-close-guard';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { type permissionservicev1_PermissionGroup as PermissionGroup } from '#/generated/api/admin/service/v1';
import { buildPermissionTree, fetchListPermissionGroups, fetchListPermissions, PaginationQuery, statusList, useCreateRole, useUpdateRole } from '#/api';
import { filterNumbers } from '#/utils';

const { mutateAsync: createRole } = useCreateRole();
const { mutateAsync: updateRole } = useUpdateRole();

const data = ref();
const groups = ref<PermissionGroup[]>([]);

const getTitle = computed(() =>
  data.value?.create
    ? $t('ui.modal.create', { moduleName: $t('page.role.moduleName') })
    : $t('ui.modal.update', { moduleName: $t('page.role.moduleName') }),
);
// const isCreate = computed(() => data.value?.create);

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  // 所有表单项共用，可单独在表单内覆盖
  commonConfig: {
    // react 基准抽屉标签列为 120px，这里放宽到 130 容纳「字段权限（User）」整行不折行
    labelWidth: 130,
    // 所有表单项
    componentProps: {
      class: 'w-full',
    },
  },
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.role.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'code',
      label: $t('page.role.code'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'InputNumber',
      fieldName: 'sortOrder',
      defaultValue: 1,
      label: $t('ui.table.sortOrder'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
        min: 0,
        precision: 0,
      },
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      fieldName: 'status',
      label: $t('ui.table.status'),
      defaultValue: 'ON',
      rules: 'selectRequired',
      componentProps: {
        optionType: 'button',
        buttonStyle: 'solid',
        class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
        options: statusList,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'description',
      label: $t('ui.table.description'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'ApiTree',
      fieldName: 'permissions',
      componentProps: {
        title: $t('page.role.permissions'),
        showSearch: true,
        // 与 react 端权限树一致：数据加载完成后默认全展开
        treeDefaultExpandAll: true,
        loadingSlot: 'suffixIcon',
        childrenField: 'children',
        labelField: 'title',
        valueField: 'key',
        resultField: 'items',
        api: async () => {
          const groupData = await fetchListPermissionGroups(
            new PaginationQuery({ formValues: { status: 'ON' } }),
          );
          groups.value = groupData.items ?? [];

          return await fetchListPermissions(
            new PaginationQuery({ formValues: { status: 'ON' } }),
          );
        },
        afterFetch: (data: any) => {
          return buildPermissionTree(groups.value, data.items);
        },
      },
    },
    {
      // 字段权限：勾选 = 对该角色用户隐藏（黑名单语义）。
      // value 与后端各资源 proto 字段 json_name 逐字一致。
      // 多资源分组与 react 端 ROLE_FIELD_RESOURCES 登记表对齐（当前 User/Role）。
      component: 'CheckboxGroup',
      fieldName: 'userHiddenFields',
      defaultValue: [],
      label: $t('page.role.fieldPerm.userTitle'),
      help: $t('page.role.fieldPerm.hint'),
      componentProps: {
        options: [
          { label: $t('page.role.fieldPerm.field.email'), value: 'email' },
          { label: $t('page.role.fieldPerm.field.mobile'), value: 'mobile' },
          { label: $t('page.role.fieldPerm.field.telephone'), value: 'telephone' },
          { label: $t('page.role.fieldPerm.field.address'), value: 'address' },
          { label: $t('page.role.fieldPerm.field.region'), value: 'region' },
          { label: $t('page.role.fieldPerm.field.lastLoginAt'), value: 'lastLoginAt' },
          { label: $t('page.role.fieldPerm.field.lastLoginIp'), value: 'lastLoginIp' },
        ],
      },
    },
    {
      component: 'CheckboxGroup',
      fieldName: 'roleHiddenFields',
      defaultValue: [],
      label: $t('page.role.fieldPerm.roleTitle'),
      componentProps: {
        options: [
          {
            label: $t('page.role.fieldPerm.field.permissions'),
            value: 'permissions',
          },
        ],
      },
    },
  ],
});

// 关闭守卫：遮罩/ESC/X/取消统一，脏数据先确认（docs/form_drawer_close_guard.md）
const closeGuard = useFormCloseGuard(baseFormApi);

const [Drawer, drawerApi] = useVbenDrawer({
  onBeforeClose: () => closeGuard.onBeforeClose(),
  onCancel() {
    drawerApi.close();
  },

  async onConfirm() {

    // 校验输入的数据
    const validate = await baseFormApi.validate();
    if (!validate.valid) {
      return;
    }

    setLoading(true);

    // 获取表单数据
    const values = await baseFormApi.getValues();
    // @ts-ignore JSON.stringify
    const finalValues = deepClone(values);

    if (
      finalValues.permissions !== null &&
      Array.isArray(finalValues.permissions) &&
      finalValues.permissions.length > 0
    ) {
      finalValues.permissions = filterNumbers(values.permissions);
    }

    // 仅 SELECTED_UNITS 档提交授权单元集（含清空场景）；
    // 其余档位不携带该字段，后端维持既有集不替换。
    if (finalValues.dataScope === 'SELECTED_UNITS') {
      finalValues.orgUnits = Array.isArray(values.orgUnits)
        ? filterNumbers(values.orgUnits)
        : [];
    } else {
      delete finalValues.orgUnits;
    }

    // 字段权限始终随表单提交（含清空场景）：抽屉所见即保存后的最终态。
    // userHiddenFields/roleHiddenFields 是表单专用中间字段，不是 Role proto
    // 字段，提交前必须折算成 fieldPermissions 多资源数组并移除。
    finalValues.fieldPermissions = (
      [
        ['User', finalValues.userHiddenFields],
        ['Role', finalValues.roleHiddenFields],
      ] as const
    )
      .filter(([, fields]) => Array.isArray(fields) && fields.length > 0)
      .map(([resource, fields]) => ({
        resource,
        hiddenFields: fields.filter((v: unknown) => typeof v === 'string'),
      }));
    delete finalValues.userHiddenFields;
    delete finalValues.roleHiddenFields;


    try {
      await (data.value?.create
        ? createRole(finalValues)
        : updateRole({ id: data.value.row.id, values: finalValues }));

      notification.success({
        message: data.value?.create
          ? $t('ui.notification.create_success')
          : $t('ui.notification.update_success'),
      });
      closeGuard.markClean();
    } catch {
      notification.error({
        message: data.value?.create
          ? $t('ui.notification.create_failed')
          : $t('ui.notification.update_failed'),
      });
    } finally {
      drawerApi.close();
      setLoading(false);
    }
  },

  async onOpenChange(isOpen) {
    if (isOpen) {
      // 获取传入的数据
      data.value = drawerApi.getData<Record<string, any>>();

      // 为表单赋值。注意：setValues 必须只调一次——连调两次时后一次会清掉
      // 前一次刚设置的值（vee-validate setValues 非合并语义），因此把字段权限
      // 的中间态 userHiddenFields/roleHiddenFields 合并在同一次调用里。
      const row = data.value?.row;
      const findHiddenFields = (resource: string) => {
        const entry = Array.isArray(row?.fieldPermissions)
          ? row.fieldPermissions.find((e: any) => e?.resource === resource)
          : undefined;
        return Array.isArray(entry?.hiddenFields)
          ? entry.hiddenFields.filter((v: any) => typeof v === 'string')
          : [];
      };
      await baseFormApi.setValues({
        ...(row ?? {}),
        userHiddenFields: findHiddenFields('User'),
        roleHiddenFields: findHiddenFields('Role'),
      });
      await closeGuard.syncBaseline();

      // setLoading(true);
      setLoading(false);
    }
  },
});

function setLoading(loading: boolean) {
  drawerApi.setState({ loading });
}
</script>

<template>
  <Drawer :title="getTitle">
    <BaseForm />
  </Drawer>
</template>
