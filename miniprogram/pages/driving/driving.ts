// 每秒钟0.7分钱
const centsPerSec = 0.7

// 将时间计数转换成00:00:00的格式
function formatDuration(sec: number): string {
  // 如果是个位数，就在前面补上'0'
  const padString = (n: number) =>
    n < 10 ? '0' + n.toFixed(0) : n.toFixed(0)
  const h = Math.floor(sec / 3600)
  sec -= h * 3600
  const m = Math.floor(sec / 60)
  sec -= m * 60
  const s = Math.floor(sec)
  return `${padString(h)}:${padString(m)}:${padString(s)}`
}

// 将分钱数转换成0.00的格式
// cents是多少分钱
function formatFare(cents: number): string {
  return (cents / 100).toFixed(2)
}

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

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {
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