package filter

import (
	"fmt"
	"testing"
)

func TestLoadWordDict(t *testing.T) {
	filter, err := LoadWordDict("sensitiveDict.txt")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("start to place the string")
	str := []byte("我空ss子sss我是霸**王*龙,我是我我是个(S)(B)真的,TMD，他妈的")
	fmt.Println(string(filter.Replace(str, '*')))
}
