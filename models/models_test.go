package models_test

import (
	. "github.com/gravida/work/models"
	"github.com/gravida/work/pkg/settings"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestWork(t *testing.T) {
	settings.DatabaseCfg.Type = "sqlite3"
	settings.DatabaseCfg.Path = "data.db"
	Setup()
	defer func() {
		os.Remove("./data.db")
	}()
	Convey("添加work", t, func() {
		w := Work{Uid: 1, Name: "测试"}
		err := AddWork(&w)
		So(err, ShouldBeNil)

		w = Work{Uid: 2, Name: "测试"}
		err = AddWork(&w)
		So(err, ShouldBeNil)
	})

	Convey("判断work名称", t, func() {
		exist, err := ExistWorkByName(1, "测试")
		So(err, ShouldBeNil)
		So(exist, ShouldEqual, true)

		exist, err = ExistWorkByName(1, "测试1")
		So(err, ShouldBeNil)
		So(exist, ShouldEqual, false)

		exist, err = ExistWorkByName(3, "测试")
		So(err, ShouldBeNil)
		So(exist, ShouldEqual, false)
	})
	Convey("获取works", t, func() {
		works, err := QueryAllWorks(1, 20)

		for i := 0; i < len(works); i++ {
			t.Log(works[i])
		}
		So(err, ShouldBeNil)
	})
}
