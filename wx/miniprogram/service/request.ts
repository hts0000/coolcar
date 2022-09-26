import camelcaseKeys from "camelcase-keys"
import { auth } from "../gen/ts/auth/auth_pb"

export namespace Coolcar {
    const serverAddr = "http://localhost:8080"
    const AUTH_ERR = "AUTH_ERR"

    const authData = {
        token: '',
        expirMs: 0,
    }

    export interface RequestOption<REQ, RES> {
        method: "GET" | "PUT" | "POST" | "DELETE"
        path: string
        data?: REQ
        respMarshaller: (r: object) => RES
    }

    export interface AuthOption {
        attachAuthHeader: boolean
        retryOnAuthError: boolean
    }

    export async function sendRequestWithAuthRetry<REQ extends string | AnyObject | ArrayBuffer | undefined,
        RES>(o: RequestOption<REQ, RES>, a?: AuthOption): Promise<RES> {
        const authOpt = a || {
            attachAuthHeader: true,
            retryOnAuthError: true,
        }
        try {
            await login()
            return sendRequest(o, authOpt)
        } catch (error) {
            if (error === AUTH_ERR && authOpt.retryOnAuthError) {
                authData.token = ""
                authData.expirMs = 0
                return sendRequestWithAuthRetry(o, {
                    attachAuthHeader: authOpt.attachAuthHeader,
                    retryOnAuthError: false,
                })
            } else {
                throw error
            }
        }
    }

    export async function login() {
        if (authData.token && authData.expirMs >= Date.now()) {
            return
        }
        const wxResp = await wxLogin()
        const reqTimeMs = Date.now()
        const resp = await sendRequest<auth.v1.ILoginRequest, auth.v1.ILoginResponse>({
            method: "POST",
            path: "/v1/auth/login",
            data: {
                code: wxResp.code,
            },
            respMarshaller: auth.v1.LoginResponse.fromObject,
        }, {
            attachAuthHeader: false,
            retryOnAuthError: false,
        })
        authData.token = resp.accessToken!
        authData.expirMs = reqTimeMs + resp.expirseIn! * 1000
    }

    function sendRequest<REQ extends string | AnyObject | ArrayBuffer | undefined,
        RES>(o: RequestOption<REQ, RES>, a: AuthOption): Promise<RES> {
        const authOpt = a || {
            attachAuthHeader: true,
        }
        return new Promise((resolve, reject) => {
            const header: Record<string, any> = {}
            if (authOpt.attachAuthHeader) {
                if (authData.token && authData.expirMs >= Date.now()) {
                    header.authorization = "Bearer " + authData.token
                } else {
                    reject(AUTH_ERR)
                    return
                }
            }
            wx.request({
                url: serverAddr + o.path,
                method: o.method,
                data: o.data,
                success: res => {
                    if (res.statusCode === 401) {
                        reject(AUTH_ERR)
                    } else if (res.statusCode >= 400) {
                        reject(res)
                    } else {
                        resolve(o.respMarshaller(camelcaseKeys(res.data as object, {
                            deep: true,
                        })))
                    }
                },
                fail: reject,
                header: header,

            })
        })
    }

    function wxLogin(): Promise<WechatMiniprogram.LoginSuccessCallbackResult> {
        return new Promise((resolve, reject) => {
            wx.login({
                success: resolve,
                fail: reject,
            })
        })
    }
}