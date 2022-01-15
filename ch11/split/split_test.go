package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("ABABA", "B")
	want := []string{"A", "A", "A"}
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("expected:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}
