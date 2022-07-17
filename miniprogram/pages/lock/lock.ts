const shareLocationKey = 'is_share_location'

// pages/lock/lock.ts
Page({

  /**
   * 页面的初始数据
   */
  data: {
    isShareLocation: false,
    userInfo: {} as WechatMiniprogram.UserInfo,
    hasUserInfo: false,
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
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }
    })
  },

  // 记录用户是否展示行程
  onShareLocation(e: any) {
    const isShareLocation: Boolean = e.detail.value
    // setStorageSync会以键值对的方式存储在手机本地，重新打开小程序还可以获取到
    // 相当于一个键值对数据库
    wx.setStorageSync(shareLocationKey, isShareLocation)
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {
    // 每次打开小程序时，就去获取是否分享行程这个值
    // 如果没有这个值，则默认设置为false
    this.setData({
      isShareLocation: wx.getStorageSync(shareLocationKey) || false,
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