import type { DrawerApiOptions, DrawerState } from './drawer';

import { Store } from '@vben-core/shared/store';
import { bindMethods, isFunction } from '@vben-core/shared/utils';

export class DrawerApi {
  private api: Pick<
    DrawerApiOptions,
    'onBeforeClose' | 'onCancel' | 'onConfirm' | 'onOpenChange'
  >;
  // private prevState!: DrawerState;
  private state!: DrawerState;

  // 共享数据
  public sharedData: Record<'payload', any> = {
    payload: {},
  };

  public store: Store<DrawerState>;

  constructor(options: DrawerApiOptions = {}) {
    const {
      connectedComponent: _,
      onBeforeClose,
      onCancel,
      onConfirm,
      onOpenChange,
      ...storeState
    } = options;

    const defaultState: DrawerState = {
      class: '',
      closable: true,
      closeOnClickModal: true,
      closeOnPressEscape: true,
      confirmLoading: false,
      contentClass: '',
      footer: true,
      header: true,
      isOpen: false,
      loading: false,
      modal: true,
      openAutoFocus: false,
      placement: 'right',
      showCancelButton: true,
      showConfirmButton: true,
      title: '',
    };

    this.store = new Store<DrawerState>(
      {
        ...defaultState,
        ...storeState,
      },
      {
        onUpdate: () => {
          const state = this.store.state;
          if (state?.isOpen === this.state?.isOpen) {
            this.state = state;
          } else {
            this.state = state;
            this.api.onOpenChange?.(!!state?.isOpen);
          }
        },
      },
    );
    this.state = this.store.state;
    this.api = {
      onBeforeClose,
      onCancel,
      onConfirm,
      onOpenChange,
    };
    bindMethods(this);
  }

  // 如果需要多次更新状态，可以使用 batch 方法
  batchStore(cb: () => void) {
    this.store.batch(cb);
  }

  /**
   * 关闭弹窗
   */
  close() {
    // 通过 onBeforeClose 钩子函数来判断是否允许关闭弹窗
    // 同步返回 false 不关闭；返回 Promise 时以 resolve 值决定（异步守卫，
    // 如表单抽屉的脏数据确认，见 docs/form_drawer_close_guard.md）
    const allowClose = this.api.onBeforeClose?.() ?? true;
    if (allowClose instanceof Promise) {
      void allowClose.then((ok) => {
        if (ok) {
          this.store.setState((prev) => ({ ...prev, isOpen: false }));
        }
      });
      return;
    }
    if (allowClose !== false) {
      this.store.setState((prev) => ({ ...prev, isOpen: false }));
    }
  }

  getData<T extends object = Record<string, any>>() {
    return (this.sharedData?.payload ?? {}) as T;
  }

  /**
   * 取消操作
   */
  onCancel() {
    if (this.api.onCancel) {
      this.api.onCancel?.();
    } else {
      this.close();
    }
  }

  /**
   * 确认操作
   */
  onConfirm() {
    this.api.onConfirm?.();
  }

  open() {
    this.store.setState((prev) => ({ ...prev, isOpen: true }));
  }

  setData<T>(payload: T) {
    this.sharedData.payload = payload;
  }

  setState(
    stateOrFn:
      | ((prev: DrawerState) => Partial<DrawerState>)
      | Partial<DrawerState>,
  ) {
    if (isFunction(stateOrFn)) {
      this.store.setState(stateOrFn);
    } else {
      this.store.setState((prev) => ({ ...prev, ...stateOrFn }));
    }
  }
}
