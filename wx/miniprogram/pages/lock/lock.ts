import { IAppOption } from "../../appoption"
import { car } from "../../gen/ts/auth/car_pb"
import { rental } from "../../gen/ts/auth/rental_pb"
import { CarService } from "../../service/car"
import { TripService } from "../../service/trip"
import { routing } from "../../utils/routing"

const shareLocationKey = 'is_share_location'

// pages/lock/lock.ts
Page({
  /**
   * 页面的初始数据
   */
  car_id: "",
  carRefresher: 0,
  data: {
    isShareLocation: false,
    userInfo: {} as WechatMiniprogram.UserInfo,
    hasUserInfo: false,
    avatarURL: '',
  },

  // 获取用户信息的回调函数
  getUserProfile(e: any) {
    // console.log("eeeeeee", e)
    // 推荐使用 wx.getUserProfile 获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认
    // 开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    wx.getUserProfile({
      desc: '用于实时展示头像', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
        // console.log("res", res)
        // 要自己把userInfo存下来
        getApp<IAppOption>().globalData.userInfo = res.userInfo
        console.log(getApp<IAppOption>().globalData.userInfo)
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }
    })
  },

  // 记录用户是否展示行程
  onShareLocation(e: any) {
    this.data.isShareLocation = e.detail.value
    // setStorageSync会以键值对的方式存储在手机本地，重新打开小程序还可以获取到
    // 相当于一个键值对数据库
    wx.setStorageSync(shareLocationKey, true)
  },

  // 前往行程页面
  onUnlockTap() {
    // 获取位置信息权限，为后续行程页面做准备
    wx.getLocation({
      type: 'gcj02',
      success: async loc => {
        if (!this.car_id) {
          console.error("no car_id specified")
          return
        }

        // 尝试创建行程
        let trip: rental.v1.ITripEntity
        try {
          trip = await TripService.CreateTrip({
            start: {
              latitude: loc.latitude,
              longitude: loc.longitude,
            },
            carId: this.car_id,
            avatarUrl: this.data.isShareLocation ? this.data.avatarURL : '',
          })
          if (!trip.id) {
            console.error("no tripID in response", trip)
            return
          }
        } catch (error) {
          wx.showToast({
            title: "创建行程失败",
            icon: "none",
          })
          return
        }

        // 显示一个开锁中提示
        wx.showLoading({
          title: '开锁中',
          // 为页面覆盖一个透明的罩子，避免开锁中时点击到其他元素
          mask: true,
        })

        this.carRefresher = setInterval(async () => {
          const c = await CarService.getCar(this.car_id)
          if (c.status === car.v1.CarStatus.UNLOCKED) {
            this.clearCarRefresher()
            wx.redirectTo({
              url: routing.driving({
                trip_id: trip.id!,
              }),
              complete: () => {
                wx.hideLoading()
              },
            })
          }
        }, 2000)
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

  clearCarRefresher() {
    if (this.carRefresher) {
      clearInterval(this.carRefresher)
      this.carRefresher = 0
    }
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt: Record<'car_id', string>) {
    const o: routing.LockOpts = opt
    this.car_id = o.car_id
    console.log("car_id =", this.car_id)
    // 每次打开小程序时，就去获取是否分享行程这个值
    // 如果没有这个值，则默认设置为true
    // 有则取本地值
    this.setData({
      isShareLocation: wx.getStorageSync(shareLocationKey) || false,
      avatarURL: this.data.userInfo.avatarUrl,
    })
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {
    this.clearCarRefresher()
    wx.hideLoading()
  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  }
})