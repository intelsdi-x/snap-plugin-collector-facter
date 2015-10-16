//
// +build unit

package facter

import (
	"testing"

	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/intelsdi-x/pulse/control/plugin/cpolicy"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFacterPlugin(t *testing.T) {
	Convey("Meta should return metadata for the plugin", t, func() {
		meta := Meta()
		So(meta.Name, ShouldResemble, name)
		So(meta.Version, ShouldResemble, version)
		So(meta.Type, ShouldResemble, plugin.CollectorPluginType)
	})

	Convey("Create Facter Collector", t, func() {
		fCol := NewFacterCollector()
		Convey("So fCol should not be nil", func() {
			So(fCol, ShouldNotBeNil)
		})
		Convey("So fCol should be of Facter type", func() {
			So(fCol, ShouldHaveSameTypeAs, &Facter{})
		})
		configPolicy, err := fCol.GetConfigPolicy()
		Convey("fCol.GetConfigPolicy() should return a config policy", func() {
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So getting config policy should not return an error", func() {
				So(err, ShouldBeNil)
			})
			Convey("So config policy should be a cpolicy.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, &cpolicy.ConfigPolicy{})
			})
		})
	})
}
