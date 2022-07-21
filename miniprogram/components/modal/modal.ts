import { ModalResult } from "../types"

// components/modal/modal.ts
Component({
  /**
   * 组件的属性列表
   */
  properties: {
    // 属性列表可以在wxml中直接使用
    showModal: Boolean,
    showCancel: Boolean,
    title: String,
    contents: String,
  },

  options: {
    // 让该自定义组件也使用全局样式
    addGlobalClass: true,
  },

  /**
   * 组件的初始数据
   */
  data: {
    resolve: undefined as ((r: ModalResult) => void) | undefined
  },

  /**
   * 组件的方法列表
   */
  methods: {
    onCancel() {
      this.hideModal('cancel')
    },

    onOK() {
      this.hideModal('ok')
    },

    hideModal(res: ModalResult) {
      this.setData({
        showModal: false,
      })
      // 产生事件并通知外界
      this.triggerEvent(res)
      if (this.data.resolve) {
        this.data.resolve(res)
      }
    },

    showModal(): Promise<ModalResult> {
      this.setData({
        showModal: true,
      })
      return new Promise((resolve) => {
        this.data.resolve = resolve
      })
    },
  }
})
