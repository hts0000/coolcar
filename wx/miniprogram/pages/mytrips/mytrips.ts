import { IAppOption } from "../../appoption"
import { rental } from "../../gen/ts/auth/rental_pb"
import { ProfileService } from "../../service/profile"
import { TripService } from "../../service/trip"
import { formatDuration, formatFare, myFormat } from "../../utils/format"
import { routing } from "../../utils/routing"

interface Trip {
  id: string
  start: string
  end: string
  duration: string
  fee: string
  distance: string
  status: string
  inProgress: boolean
}

const tripStatusMap = new Map([
  [rental.v1.TripStatus.IN_PROGRESS, "进行中"],
  [rental.v1.TripStatus.FINISHED, "已完成"],
])

const licStatusMap = new Map([
  [rental.v1.IdentityStatus.UNSUBMITTED, "未认证"],
  [rental.v1.IdentityStatus.PENDING, "未认证"],
  [rental.v1.IdentityStatus.VERIFIED, "已认证"],
])

// pages/mytrips/mytrips.ts
Page({

  /**
   * 页面的初始数据
   */
  data: {
    userInfo: {} as WechatMiniprogram.UserInfo,
    hasUserInfo: false,
    avatarURL: '' as string | undefined,
    indicatorDots: true,
    autoPlay: true,
    interval: 3000,
    duration: 500,
    circular: true,
    multiItemCount: 1,
    prevMargin: '',
    nextMargin: '',
    vertical: false,
    current: 0,
    licStatus: licStatusMap.get(rental.v1.IdentityStatus.UNSUBMITTED),
    promotionItems: [
      {
        img: 'https://img2.mukewang.com/62c64b510001a11117920764.jpg',
        promotionID: 1
      },
      {
        img: 'https://img3.mukewang.com/62d65b370001f6f417920764.jpg',
        promotionID: 2
      },
      {
        img: 'https://img2.mukewang.com/62d4c1950001da0d17920764.jpg',
        promotionID: 3
      },
      {
        img: 'https://img.mukewang.com/62d0d2de0001a9ad17920764.jpg',
        promotionID: 4
      },
    ],
    trips: [] as rental.v1.ITripEntity[],
    tripsHeight: 0,
  },

  onSwiperChange(e: any) {
    // e.detail.source监测轮播变换是自动触发还是用户滑动触发
    // 分为autoplay（自动轮播）、touch（用户滑动）、""（程序控制滑动）
    // console.log(e, e.detail.source)
  },

  // 监测轮播图片点击事件，通过data-可以返回指定数据
  onPromotionItemTap(e: any) {
    // e.currentTarget.dataset.promotionId获得点击的图片的ID
    // console.log(e)
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
          hasUserInfo: true,
          avatarURL: res.userInfo.avatarUrl,
        })
      }
    })
  },

  onRegisterTap() {
    wx.navigateTo({
      url: routing.register(),
    })
  },

  populateTrips(data: rental.v1.ITripEntity[]) {
    const trips: Trip[] = []
    for (let i = 0; i < data.length; i++) {
      let d = data[i].trip
      if (!d) {
        continue
      }
      const trip: Trip = {
        id: data[i].id!,
        start: d.start?.poiName || "未知",
        end: d.end?.poiName || "未知",
        duration: formatDuration((d.end?.timestampSec as number) - (d.start?.timestampSec as number)),
        fee: formatFare(d.end?.feeCent || 0),
        distance: d.end?.kmDriven?.toFixed(1) + "公里",
        status: tripStatusMap.get(d.status!) || "未知",
        inProgress: d.status! === rental.v1.TripStatus.IN_PROGRESS,
      }
      if (trip.status === "进行中") {
        trips.unshift(trip)
      } else {
        trips.push(trip)
      }
    }
    this.setData({
      trips: trips,
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  async onLoad() {
    const trs = await TripService.GetTrips()
    // 为垂直滚动栏加载模拟数据
    this.populateTrips(trs.trips)
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {
    // 根据实际机型长宽来设置垂直滚动栏的高度
    // createSelectorQuery: 类似于jQuery根据id或class获取页面元素
    //    然后对这个元素建立一个查询请求
    //    查询成功后执行回调（rect中包含了该元素的各种信息，包括该元素在页面中实际占据的长宽）
    // getSystemInfoSync: 获取当前机型的数据，如型号、系统版本、窗口高度等
    wx.createSelectorQuery().select('#heading').boundingClientRect(rect => {
      this.setData({
        // 当前机型窗口高度减去固定元素高度，剩下高度就是垂直滚动栏的高度
        tripsHeight: wx.getSystemInfoSync().windowHeight - rect.height
      })
    }).exec() // 实际执行boundingClientRect获取该元素信息
  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {
    // 第一次获取用户信息在lock页面
    if (!this.data.avatarURL) {
      const userInfo = getApp<IAppOption>().globalData.userInfo
      this.setData({
        avatarURL: userInfo?.avatarUrl || '',
      })
    }
    ProfileService.getProfile().then(p => {
      this.setData({
        licStatus: licStatusMap.get(p.identityStatus || 0),
      })
    })
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