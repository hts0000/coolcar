import { rental } from "../../gen/ts/auth/rental_pb"
import { ProfileService } from "../../service/profile"
import { routing } from "../../utils/routing"
import { formatTime } from "../../utils/wxapi";
import { myFormat } from "../../utils/format";

// pages/register/register.ts
Page({

  /**
   * 页面的初始数据
   */
  redirectURL: '',
  profileRefresher: 0,
  data: {
    licNo: '',
    name: '',
    genderIndex: 0,
    genders: ['未知', '男', '女'],
    licImgURL: '',
    birthday: '1999-01-01',
    state: rental.v1.IdentityStatus[rental.v1.IdentityStatus.UNSUBMITTED],
  },

  // 上传驾驶证实现
  onUploadLic() {
    wx.chooseImage({
      success: (res) => {
        if (res.tempFilePaths.length > 0) {
          this.setData({
            licImgURL: res.tempFilePaths[0]
          })

          const data = wx.getFileSystemManager().readFileSync(res.tempFilePaths[0])
          wx.request({
            method: "PUT",
            url: "https://coolcar-1300912551.cos.ap-guangzhou.myqcloud.com/account_1/63592214e24fb8ead7e52a01?q-sign-algorithm=sha1&q-ak=AKIDrdAUXKq69xVqwlV1HH0RguxlPpz50kHc&q-sign-time=1666785812%3B1666786812&q-key-time=1666785812%3B1666786812&q-header-list=host&q-url-param-list=&q-signature=a7c150eb5aa566ee7f4946ffe52de68a3a4d9b6a",
            data: data,
            success: console.log,
            fail: console.error,
          })
        }
      }
    })
  },

  // 驾照认证界面-表单事件-性别表单改变实现
  // e是点击事件发生产生的数据
  onGenderChange(e: any) {
    // 这个只能打印出来看那些是我们想要的数据，再选择
    // console.log(e)
    this.setData({
      genderIndex: parseInt(e.detail.value),
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
    ProfileService.submitProfile({
      licNumber: this.data.licNo,
      name: this.data.name,
      gender: this.data.genderIndex,
      birthDateMillis: Date.parse(this.data.birthday),
    }).then(p => {  // 提交审核之后，轮训等待后台审核通过
      this.renderProfile(p)
      this.scheduleProfileRefresher()
    })
  },

  scheduleProfileRefresher() {
    // 1s 做一次getProfile请求，检查是否通过审核
    this.profileRefresher = setInterval(() => {
      ProfileService.getProfile().then(p => {
        this.renderProfile(p)
        if (p.identityStatus !== rental.v1.IdentityStatus.PENDING) {
          this.clearProfileRefresher()
        }
        if (p.identityStatus === rental.v1.IdentityStatus.VERIFIED) {
          this.onLicVerified()
        }
      })
    }, 1000)
  },

  clearProfileRefresher() {
    if (this.profileRefresher) {
      clearInterval(this.profileRefresher)
      this.profileRefresher = 0
    }
  },

  // 清掉之前表单的数据，让用户可以重新上传
  onReSubmit() {
    ProfileService.clearProfile().then(p => this.renderProfile(p))
  },

  // 修改驾驶证认证状态
  onLicVerified() {
    // redirect会带上扫码的车辆信息，如果redirect为空，说明不是由租车扫码进入认证页面的
    // 因此留在当前页面即可。如果不为空，说明要租车，跳转至车辆解锁页面。
    if (this.redirectURL) {
      // redirectTo跳转至新页面，不会保留当前页面，不可退回
      wx.redirectTo({
        url: this.redirectURL,
      })
    }
  },

  renderProfile(p: rental.v1.IProfile) {
    this.setData({
      licNo: p.identity?.licNumber || "",
      name: p.identity?.name || "",
      genderIndex: p.identity?.gender || 0,
      birthday: myFormat(formatTime(new Date(p.identity?.birthDateMillis as number || 0))),
      state: rental.v1.IdentityStatus[p.identityStatus || 0],
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt: Record<'redirectURL', string>) {
    const o: routing.RegisterOpts = opt
    if (o.redirectURL) {
      this.redirectURL = decodeURIComponent(opt.redirectURL)
    }
    ProfileService.getProfile().then(p => this.renderProfile(p))
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
    this.clearProfileRefresher()
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