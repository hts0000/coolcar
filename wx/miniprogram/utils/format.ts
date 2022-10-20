// 将时间计数转换成00:00:00的格式
export function formatDuration(sec: number): string {
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
export function formatFare(cents: number): string {
    return (cents / 100).toFixed(2)
}

export function myFormat(s: string): string {
    const date = s.split(" ")[0].split("/").join("-")
    return date
}