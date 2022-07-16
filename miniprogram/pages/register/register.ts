// pages/register/register.ts
Page({

  /**
   * 页面的初始数据
   */
  data: {
    licNo: '',
    name: '',
    genderIndex: 0,
    genders: ['未知', '男', '女', '其他'],
    licImgURL: undefined as string | undefined,
    birthday: '1999-01-01',
  },

  // 上传驾驶证实现
  onUploadLic() {
    wx.chooseImage({
      success: (res) => {
        if (res.tempFilePaths.length > 0) {
          this.setData({
            licImgURL: res.tempFilePaths[0]
          })
        }
        // TODO: upload image
        setTimeout(() => {
          this.setData({
            licNo: '123456',
            name: '张三',
            genderIndex: 1,
            birthday: '1989-01-01',
          })
        }, 1000)
      }
    })
  },

  // 驾照认证界面-表单事件-性别表单改变实现
  // e是点击事件发生产生的数据
  onGenderChange(e: any) {
    // 这个只能打印出来看那些是我们想要的数据，再选择
    // console.log(e)
    this.setData({
      genderIndex: e.detail.value,
    })
  },

  // 驾照认证界面-表单事件-出生日期改变实现
  onBirthdayChange(e: any) {
    this.setData({
      birthday: e.detail.value,
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {

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