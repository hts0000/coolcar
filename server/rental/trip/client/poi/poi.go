package poi

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"hash/fnv"

	"google.golang.org/protobuf/proto"
)

var poi = []string{
	"中关村", "天安门", "陆家嘴", "平安大厦", "世界之窗", "迪士尼", "天河体育中心", "广州塔",
}

type Manager struct{}

func (m *Manager) Resolve(c context.Context, loc *rentalpb.Location) (string, error) {
	// 把 loc 转换成字节流
	b, err := proto.Marshal(loc)
	if err != nil {
		return "", nil
	}
	// 根据 loc 字节流计算 hash 值
	// 这样的好处是，当测试时，因为 loc 不变，
	// 则 hash 结果是固定的，因此返回的 poi 也是固定的
	h := fnv.New32()
	h.Write(b)
	return poi[int(h.Sum32())%len(poi)], nil
}
