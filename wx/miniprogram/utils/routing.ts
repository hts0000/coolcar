export namespace routing {
    export interface DrivingOpts {
        tripID: string
    }
    export function driving(o: DrivingOpts) {
        return `/pages/driving/driving?trip_id=${o.tripID}`
    }
    export interface LockOpts {
        carID: string
    }
    export function lock(o: LockOpts) {
        return `/pages/lock/lock?car_id=${o.carID}`
    }
    export interface RegisterOpts {
        redirectURL?: string
    }
    export interface RegisterParams {
        redirectURL: string
    }
    export function register(p?: RegisterParams) {
        const page = '/pages/register/register'
        if (!p) {
            return page
        }
        return `${page}?redirectURL=${encodeURIComponent(p.redirectURL)}`
    }
    export function mytrips() {
        return '/pages/mytrips/mytrips'
    }
}