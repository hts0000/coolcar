// pages/driving/driving.ts
Page({

  /**
   * 页面的初始数据
   */
  data: {
    drivingTime: "01:23:45",
    fare: "12.34",
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
      this.setData({
        location: {
          latitude: loc.latitude,
          longitude: loc.longitude,
        },
      })
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {
    this.setupLocationUpdator()
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