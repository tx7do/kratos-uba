import type { ExtendedFormApi } from '@vben/common-ui';

import { Modal } from 'ant-design-vue';
import { $t } from '@vben/locales';

/**
 * 表单抽屉关闭守卫（规范见 docs/form_drawer_close_guard.md）
 *
 * 遮罩 / ESC / 标题栏 X / 底部取消四条关闭路径在 DrawerApi 中统一汇入 close()，
 * 把 onBeforeClose 接到 useVbenDrawer options 即可全路径生效。
 * 脏判定：打开填充完成后 syncBaseline() 记 getValues() JSON 快照，关闭尝试时对比
 * 当前值（程序化回填不算脏；vben 表单的 handleValuesChange 连 setValues 都会触发，
 * 因此不用事件置脏）。提交成功后先 markClean() 再 close()。
 *
 * @example
 * ```ts
 * const closeGuard = useFormCloseGuard(baseFormApi);
 * const [Drawer, drawerApi] = useVbenDrawer({
 *   onBeforeClose: () => closeGuard.onBeforeClose(),
 *   async onConfirm() {
 *     // ...提交成功后：
 *     closeGuard.markClean();
 *     drawerApi.close();
 *   },
 *   async onOpenChange(isOpen) {
 *     if (isOpen) {
 *       await baseFormApi.setValues(row);
 *       await closeGuard.syncBaseline();
 *     }
 *   },
 * });
 * ```
 */
export function useFormCloseGuard(formApi: ExtendedFormApi, getExtra?: () => unknown) {
  let baseline: string | null = null;

  /** 打开填充完成后调用：记录基线快照（回填不算脏） */
  async function syncBaseline() {
    baseline = JSON.stringify([await formApi.getValues(), getExtra?.()]);
  }

  /** 保存成功后手动清脏（后续关闭不再确认） */
  function markClean() {
    baseline = null;
  }

  /** 接 useVbenDrawer 的 onBeforeClose：脏则弹确认，确定放弃才关闭 */
  async function onBeforeClose(): Promise<boolean> {
    if (baseline === null) return true;
    const current = JSON.stringify([await formApi.getValues(), getExtra?.()]);
    if (current === baseline) return true;
    return new Promise<boolean>((resolve) => {
      Modal.confirm({
        title: $t('ui.drawer.unsavedMessage'),
        okText: $t('ui.drawer.discardButton'),
        cancelText: $t('common.cancel'),
        onOk: () => {
          baseline = null;
          resolve(true);
        },
        onCancel: () => resolve(false),
      });
    });
  }

  return { markClean, onBeforeClose, syncBaseline };
}
