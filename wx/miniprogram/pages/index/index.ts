import { IAppOption } from "../../appoption"
import { rental } from "../../gen/ts/auth/rental_pb"
import { CarService } from "../../service/car"
import { ProfileService } from "../../service/profile"
import { TripService } from "../../service/trip"
import { routing } from "../../utils/routing"

interface Marker {
  iconPath: string,
  id: number,
  latitude: number,
  longitude: number,
  width: number,
  height: number,

}

// 默认头像
const defaultAvatar = "/resources/car.png"
// 初始位置
const initialLat = 30
const initialLng = 120

Page({
  // 当页面隐藏时，后台数据不再更新
  isPageShowing: false,
  socket: undefined as WechatMiniprogram.SocketTask | undefined,

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
      latitude: initialLat,
      longitude: initialLng,
    },
    // 3~20，缩放比例，3最大缩放
    scale: 10,
    // 叠在map上的元素
    markers: [] as Marker[],
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

  setupCarPosUpdater() {
    // 根据map元素id选中map，获取map对象进行操作
    const map = wx.createMapContext("mapId")
    const markersByCarID = new Map<string, Marker>()
    let translationInProgress = false
    const endTranslation = () => {
      translationInProgress = false
    }
    this.socket = CarService.subscribe(car => {
      if (!car.id || translationInProgress || !this.isPageShowing) {
        return
      }
      const marker = markersByCarID.get(car.id)
      if (!marker) {
        // Insert new car
        const newMarker: Marker = {
          id: this.data.markers.length,
          iconPath: car.car?.driver?.avatarUrl || defaultAvatar,
          latitude: car.car?.position?.latitude || initialLat,
          longitude: car.car?.position?.longitude || initialLng,
          height: 20,
          width: 20,
        }
        markersByCarID.set(car.id, newMarker)
        this.data.markers.push(newMarker)
        translationInProgress = true
        this.setData({
          markers: this.data.markers,
        }, endTranslation)
        return
      }

      const newAvatar = car.car?.driver?.avatarUrl || defaultAvatar
      const newLat = car.car?.position?.latitude || initialLat
      const newLng = car.car?.position?.longitude || initialLng
      if (marker.iconPath !== newAvatar) {
        // Change iconPath and possibly position
        marker.iconPath = newAvatar
        marker.latitude = newLat
        marker.longitude = newLng
        translationInProgress = true
        this.setData({
          markers: this.data.markers,
        }, endTranslation)
        return
      }

      if (marker.latitude !== newLat || marker.longitude !== newLng) {
        // Move marker
        translationInProgress = false
        map.translateMarker({
          markerId: marker.id,
          destination: {
            latitude: newLat,
            longitude: newLng,
          },
          autoRotate: false,
          rotate: 0,
          duration: 90,
          animationEnd: endTranslation,
        })
      }
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
  async onScanTap() {
    // 首先检查有无正在进行中的行程，有则跳转到行程页
    const trips = await TripService.GetTrips(rental.v1.TripStatus.IN_PROGRESS)
    if ((trips.trips?.length || 0) > 0) {
      await this.selectComponent("#tripModal").showModal()
      wx.navigateTo({
        url: routing.driving({
          trip_id: trips.trips![0].id!,
        }),
      })
      return
    }
    wx.scanCode({
      success: async () => {
        // TODO: 从二维码中获取car_id
        // 模拟已经获得car_id
        const car_id = '63933ba46ef7cc1ca1222d5e'

        // 指示register页面接下来跳转到lock页面
        const lockURL = routing.lock({
          car_id: car_id,
        })

        // 如果已经验证过驾照了，跳转到开锁页
        const prof = await ProfileService.getProfile()
        if (prof.identityStatus === rental.v1.IdentityStatus.VERIFIED) {
          wx.navigateTo({
            url: lockURL
          })
        } else {  // 没有验证过，跳转到验证页面
          // 展示一个自定义的对话框，当对话框关闭时，跳转到下一页面
          await this.selectComponent('#licModal').showModal()

          // navigateTo跳转至新页面，当前页面会保留，可退回
          // encodeURIComponent将url解析成合法形式（将/、空格之类的转义成%20这种形式）
          wx.navigateTo({
            url: routing.register({
              redirectURL: lockURL,
            }),
          })
        }
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
    if (!this.socket) {
      this.setData({
        markers: [],
      }, () => {
        this.setupCarPosUpdater()
      })
    }
  },

  onHide() {
    this.isPageShowing = false
    if (this.socket) {
      this.socket.close({
        success: () => {
          this.socket = undefined
        }
      })
    }
  },
})
