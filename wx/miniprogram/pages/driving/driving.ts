import { TripService } from "../../service/trip"
import { formatDuration, formatFare } from "../../utils/format"
import { routing } from "../../utils/routing"

// 每秒钟0.7分钱
const centsPerSec = 0.7

// pages/driving/driving.ts
Page({

  /**
   * 页面的初始数据
   */
  timer: undefined as number | undefined,
  data: {
    elapsed: "00:00:00",
    fare: "0.00",
    location: {
      latitude: 23.099994,
      longitude: 113.324520,
    },
    scale: 10,
  },

  setupLocationUpdator() {
    wx.startLocationUpdate({
      fail: console.error
    })
    wx.onLocationChange(loc => {
      console.log("driving", loc)
      this.setData({
        location: {
          latitude: loc.latitude,
          longitude: loc.longitude,
        },
      })
    })
  },

  setupTimer() {
    let elapsedSec = 0
    let cents = 0
    this.timer = setInterval(() => {
      elapsedSec++
      cents += centsPerSec
      this.setData({
        elapsed: formatDuration(elapsedSec),
        fare: formatFare(cents),
      })
    }, 1000)
  },

  onEndTripTap() {
    wx.redirectTo({
      url: routing.mytrips(),
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt: Record<'trip_id', string>) {
    const o: routing.DrivingOpts = opt
    console.log('current trip', o.trip_id)
    // o.trip_id = "634961aaf7c609eb3461bc9e"
    // 模拟获取
    TripService.GetTrip(o.trip_id).then(console.log)
    this.setupLocationUpdator()
    this.setupTimer()
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
    wx.stopLocationUpdate()
    // 重置计时器
    if (this.timer) {
      clearInterval(this.timer)
    }
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