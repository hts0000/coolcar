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
    licImgURL: '',
    birthday: '1999-01-01',
    state: 'UNSUBMITTED' as 'UNSUBMITTED' | 'PENDING' | 'VERIFIED',
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

  // 上传驾驶证照片至服务器端
  onSubmit() {
    // TODE: 上传信息至服务端，等待后端返回数据
    // 未返回期间是PENDING状态，返回成功是VERIFIED状态
    this.setData({
      state: 'PENDING',
    })
    setTimeout(this.onLicVerified, 3000)
  },

  // 清掉之前表单的数据，让用户可以重新上传
  onReSubmit() {
    this.setData({
      state: 'UNSUBMITTED',
      licImgURL: '',
    })
  },

  // 修改驾驶证认证状态
  onLicVerified() {
    this.setData({
      state: 'VERIFIED',
    })
    // redirectTo跳转至新页面，不会保留当前页面，不可退回
    wx.redirectTo({
      url: '/pages/lock/lock',
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