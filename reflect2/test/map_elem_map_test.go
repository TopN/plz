package test

import (
	"testing"
	"github.com/v2pro/plz/reflect2"
)

func Test_map_elem_map(t *testing.T) {
	var pInt = func(val int) *int {
		return &val
	}
	var pMap = func(val map[int]int) *map[int]int {
		return &val
	}
	t.Run("Set", testOp(func(api reflect2.API) interface{} {
		obj := map[int]map[int]int{}
		valType := api.TypeOf(obj).(reflect2.MapType)
		valType.Set(&obj, pInt(2), pMap(map[int]int{4:4}))
		valType.Set(&obj, pInt(3), pMap(map[int]int{9:9}))
		valType.Set(&obj, pInt(3), pMap(nil))
		return obj
	}))
}