import { rental } from "../../gen/ts/auth/rental_pb"
import { TripService } from "../../service/trip"
import { formatDuration, formatFare } from "../../utils/format"
import { routing } from "../../utils/routing"

const updateIntervalSec = 5
const initialLat = 30
const initialLng = 120

// pages/driving/driving.ts
Page({

  /**
   * 页面的初始数据
   */
  timer: undefined as number | undefined,
  tripID: "",
  data: {
    location: {
      latitude: initialLat,
      longitude: initialLng,
    },
    scale: 12,
    elapsed: '00:00:00',
    fee: '0.00',
    markers: [
      {
        iconPath: "/resources/car.png",
        id: 0,
        latitude: initialLat,
        longitude: initialLng,
        width: 20,
        height: 20,
      },
    ],
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
  // 获取最新费用
  async setupTimer(tripID: string) {
    const trip = await TripService.GetTrip(tripID)
    if (trip.status !== rental.v1.TripStatus.IN_PROGRESS) {
      console.error('trip not in progress')
      return
    }
    let secSinceLastUpdate = 0
    let lastUpdateDurationSec = (trip.current!.timestampSec! as number) - (trip.start!.timestampSec! as number)
    const toLocation = (trip: rental.v1.ITrip) => ({
      latitude: trip.current?.location?.latitude || initialLat,
      longitude: trip.current?.location?.longitude || initialLng,
    })
    const location = toLocation(trip)
    this.data.markers[0].latitude = location.latitude
    this.data.markers[0].longitude = location.longitude
    this.setData({
      elapsed: formatDuration(lastUpdateDurationSec),
      fee: formatFare(trip.current!.feeCent!),
      location,
      markers: this.data.markers,
    })

    this.timer = setInterval(() => {
      secSinceLastUpdate++
      if (secSinceLastUpdate % updateIntervalSec === 0) {
        TripService.GetTrip(tripID).then(trip => {
          lastUpdateDurationSec = (trip.current!.timestampSec! as number) - (trip.start!.timestampSec! as number)
          secSinceLastUpdate = 0
          const location = toLocation(trip)
          this.data.markers[0].latitude = location.latitude
          this.data.markers[0].longitude = location.longitude
          console.log(formatFare(trip.current!.feeCent!))
          console.log(trip.current!.feeCent!)
          this.setData({
            fee: formatFare(trip.current!.feeCent!),
            location,
            markers: this.data.markers,
          })
        }).catch(console.error)
      }
      this.setData({
        elapsed: formatDuration(lastUpdateDurationSec + secSinceLastUpdate),
      })
    }, 1000)
  },

  onEndTripTap() {
    TripService.FinishTrip(this.tripID).then(() => {
      wx.redirectTo({
        url: routing.mytrips(),
      })
    }).catch(err => {
      console.error("end trip failed", err)
      wx.showToast({
        title: "结束行程失败",
        icon: "none",
      })
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt: Record<'trip_id', string>) {
    const o: routing.DrivingOpts = opt
    this.tripID = o.trip_id
    console.log('current trip', this.tripID)

    this.setupLocationUpdator()
    this.setupTimer(this.tripID)
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