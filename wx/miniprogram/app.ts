import camelcaseKeys from "camelcase-keys"
import { IAppOption } from "./appoption"
import { auth } from "./gen/ts/auth/auth_pb"
import { rental } from "./gen/ts/auth/rental_pb"

// app.ts
App<IAppOption>({
  globalData: {},
  onLaunch() {
    // 登录
    wx.login({
      success: res => {
        console.log(res.code)
        wx.request({
          url: "http://localhost:8080/v1/auth/login",
          method: "POST",
          data: {
            code: res.code,
          } as auth.v1.ILoginRequest,
          success: res => {
            const loginResp = auth.v1.LoginResponse.fromObject(
              camelcaseKeys(res.data as object)
            )
            console.log(loginResp)
            wx.request({
              url: "http://localhost:8080/v1/trip",
              method: "POST",
              data: {
                start: "abc"
              } as rental.v1.ICreateTripRequest,
              header: {
                authorization: "Bearer " + loginResp.accessToken,
              },
            })
          },
          // success: console.log,
          fail: console.error,
        })
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      },
    })
  },
})