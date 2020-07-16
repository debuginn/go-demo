package splite

import (
	"reflect"
	_ "reflect"
	"testing"
)

func TestSplite(t *testing.T) {
	// 执行测试函数并声明期待结果
	result := Splite("a:b:c", ":")
	want := []string{"a", "b", "c"}

	// 利用 reflect 包中的 DeepEqual 函数判断结果与期待结果是否一致
	if !reflect.DeepEqual(result, want) {
		t.Errorf("excepted: %v, got:%v", result, want)
	}
}
