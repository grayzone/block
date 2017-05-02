package util

import "testing"
import "github.com/astaxie/beego"

func TestGetSeedData(t *testing.T) {
	v := GetSeedData()
	t.Logf("seed:", v)
	beego.Debug(v)
}
