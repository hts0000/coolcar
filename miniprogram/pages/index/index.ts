import { routing } from "../../utils/routing"

Page({
  // 当页面隐藏时，后台数据不再更新
  isPageShowing: false,
  data: {
    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true, // 展示当前位置
      showScale: true,  // 显示比例尺
      subKey: '',
      layerStyle: -1,
      enableZoom: true,
      enableScrool: true,
      enableRotate: false,
      showCompass: false,
      enable3D: false,
      enableOverlooking: false,
      enbaleSatellite: false,
      enableTraffic: false,
    },
    avatarURL: '' as string | undefined,
    location: {
      latitude: 23.099994,
      longitude: 113.324520,
    },
    scale: 10,  // 3~20，缩放比例，3最大缩放
    markers: [  // 叠在map上的元素
      {
        iconPath: "/resources/car.png",
        id: 0,
        latitude: 23.099994,
        longitude: 113.324520,
        width: 50,
        height: 50,
      },
      {
        iconPath: "/resources/car.png",
        id: 1,
        latitude: 23.09995,
        longitude: 113.324520,
        width: 50,
        height: 50,
      },
    ],
  },

  // 点击定位图标，将定位移动到当前位置
  onMyLocationTap() {
    // 获取当前位置的函数，传入是一个对象
    wx.getLocation({
      type: "gcj02",
      success: res => { // 函数执行成功的回调，res是成功回调后返回的数据
        this.setData({  // 成功回调后修改全局变量中的位置信息
          location: {
            latitude: res.latitude,
            longitude: res.longitude,
          },
        })
      },
      fail: () => { // 失败的回调
        wx.showToast({  // showToast会弹出一个窗口，显示内容提示用户
          title: '请在设置中打开地理位置授权',
          icon: 'none',
          duration: 2000
        })
      },
    })
  },

  // 移动车辆测试
  moveCars() {
    const dest = {
      latitude: 23.099994,
      longitude: 113.324520,
    }
    // 根据map元素id选中map，获取map对象进行操作
    const map = wx.createMapContext("mapId")
    const moveCar = () => {
      dest.latitude += 0.1
      dest.longitude += 0.1
      map.translateMarker({
        markerId: 0,
        // 移动到哪个位置上
        destination: {
          latitude: dest.latitude,
          longitude: dest.longitude,
        },
        autoRotate: false,
        rotate: 0,
        // 动画时间，用5秒的时间来移动，0的话就是瞬移了
        // 这个值不是定死的，尽量在指定时间内移动，可能距离太短一下就结束了
        duration: 5000,
        animationEnd: () => {
          // 页面未隐藏时，位置才更新
          if (this.isPageShowing) {
            moveCar()
          }
        }
      })
    }
    moveCar()
  },

  // 扫码租车按钮实现
  onScanTap() {
    wx.scanCode({
      success: async () => {
        // 展示一个自定义的对话框，当对话框关闭时，跳转到下一页面
        await this.selectComponent('#licModal').showModal()
        // TODO: 从二维码中获取car_id
        // 模拟已经获得car_id
        const carID = 'car_123'

        // 指示register页面接下来跳转到lock页面
        const redirectURL = routing.lock({
          carID: carID,
        })

        // navigateTo跳转至新页面，当前页面会保留，可退回
        // encodeURIComponent将url解析成合法形式（将/、空格之类的转义成%20这种形式）
        wx.navigateTo({
          url: routing.register({
            redirectURL: redirectURL,
          }),
        })

        // showModal弹出一个对话框
        // wx.showModal({
        //   title: '需要进行驾驶证审核',
        //   success: () => {
        //     console.log('微信内置对话框组件')
        //   },
        // })
      },
      fail: console.error
    })
  },

  // 点击头像前往个人页面
  onMyTripsTap() {
    wx.navigateTo({
      url: routing.mytrips(),
    })
  },

  onLoad() {
  },

  onShow() {
    this.isPageShowing = true
    // 第一次获取用户信息在lock页面
    console.log("onShow", getApp<IAppOption>().globalData.userInfo)
    if (!this.data.avatarURL) {
      const userInfo = getApp<IAppOption>().globalData.userInfo
      this.setData({
        avatarURL: userInfo?.avatarUrl || '',
      })
    }
  },

  onHide() {
    this.isPageShowing = false
  },
})
