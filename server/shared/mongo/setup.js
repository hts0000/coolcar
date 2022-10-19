db.account.createIndex({
    open_id: 1,
}, {
    unique: true,
})

db.trip.createIndex({
    "trip.accountid": 1,
    "trip.status": 1,
}, {
    unique: true,
    // 保证行程状态只有一个为 进行中，trip.status = 1 表示行程进行中
    partialFilterExpression: {
        "trip.status": 1,
    }
})

db.profile.createIndex({
    "accountid": 1,
}, {
    unique: true,
})