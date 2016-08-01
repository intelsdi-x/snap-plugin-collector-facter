// +build linux,integration

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
# testing
go test -v github.com/intelsdi-x/snap-plugin-collector-facter/facter
*/
package facter

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	. "github.com/smartystreets/goconvey/convey"
)

// fact expected to be available on every system
// can be allways received from Facter for test purposes
const someFact = "kernel"
const someValue = "linux 1234"

var existingNamespace = core.NewNamespace(vendor, prefix, someFact)

func TestFacterCollectMetrics(t *testing.T) {
	Convey("TestFacterCollect tests", t, func() {

		f := NewFacterCollector()
		// always return at least one metric
		f.getFacts = func(_ []string, _ time.Duration, _ *cmdConfig) (facts, error) {
			return facts{someFact: someValue}, nil
		}

		Convey("asked for nothing returns nothing", func() {
			metricTypes := []plugin.MetricType{}
			metrics, err := f.CollectMetrics(metricTypes)
			So(err, ShouldBeNil)
			So(metrics, ShouldBeEmpty)
		})

		Convey("asked for something returns something", func() {
			metricTypes := []plugin.MetricType{
				plugin.MetricType{
					Namespace_: existingNamespace,
				},
			}
			metrics, err := f.CollectMetrics(metricTypes)
			So(err, ShouldBeNil)
			So(metrics, ShouldNotBeEmpty)
			So(len(metrics), ShouldEqual, 1)

			// check just one metric
			metric := metrics[0]
			So(metric.Namespace()[2].Value, ShouldResemble, someFact)
			So(metric.Data().(string), ShouldEqual, someValue)
		})

		Convey("ask for inappriopriate metrics", func() {
			Convey("wrong number of parts", func() {
				_, err := f.CollectMetrics(
					[]plugin.MetricType{
						plugin.MetricType{
							Namespace_: core.NewNamespace("where are my other parts"),
						},
					},
				)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "segments")
			})

			Convey("wrong vendor", func() {
				_, err := f.CollectMetrics(
					[]plugin.MetricType{
						plugin.MetricType{
							Namespace_: core.NewNamespace("nonintelvendor", prefix, someFact),
						},
					},
				)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "expected vendor")
			})

			Convey("wrong prefix", func() {
				_, err := f.CollectMetrics(
					[]plugin.MetricType{
						plugin.MetricType{
							Namespace_: core.NewNamespace(vendor, "this is wrong prefix", someFact),
						},
					},
				)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "expected prefix")
			})

		})
	})
}

func TestFacterInvalidBehavior(t *testing.T) {

	Convey("returns errors as expected when cmd isn't working", t, func() {
		f := NewFacterCollector()
		// mock that getFacts returns error every time
		f.getFacts = func(_ []string, _ time.Duration, _ *cmdConfig) (facts, error) {
			return nil, errors.New("dummy error")
		}

		_, err := f.CollectMetrics([]plugin.MetricType{
			plugin.MetricType{
				Namespace_: existingNamespace,
			},
		},
		)
		So(err, ShouldNotBeNil)

		_, err = f.GetMetricTypes(plugin.ConfigType{})
		So(err, ShouldNotBeNil)
	})
	Convey("returns not as much values as asked", t, func() {

		f := NewFacterCollector()
		// mock that getFacts returns error every time
		//returns zero elements even when asked for one
		f.getFacts = func(_ []string, _ time.Duration, _ *cmdConfig) (facts, error) {
			return nil, nil
		}

		_, err := f.CollectMetrics([]plugin.MetricType{
			plugin.MetricType{
				Namespace_: existingNamespace,
			},
		},
		)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "more/less")
	})

}

func TestFacterGetMetricsTypes(t *testing.T) {

	Convey("GetMetricTypes functionallity", t, func() {

		f := NewFacterCollector()
		So(f, ShouldNotBeNil)

		Convey("GetMetricsTypes returns no error", func() {
			// executes without error
			metricTypes, err := f.GetMetricTypes(plugin.ConfigType{})
			So(err, ShouldBeNil)
			Convey("metricTypesReply should contain more than zero metrics", func() {
				So(metricTypes, ShouldNotBeEmpty)
			})

			Convey("at least one metric contains metric namespace \"intel/facter/kernel\"", func() {

				expectedNamespaceStr := existingNamespace.String()

				found := false
				for _, metricType := range metricTypes {
					if metricType.Namespace().String() == expectedNamespaceStr {
						found = true
						break
					}
				}
				if !found {
					t.Error("It was expected to find at least on intel/facter/kernel metricType (but it wasn't there)")
				}
			})

			Convey("Then list of metrics is returned", func() {
				So(len(metricTypes), ShouldNotBeNil)
				namespaces := []string{}
				for _, m := range metricTypes {
					namespaces = append(namespaces, strings.Join(m.Namespace().Strings(), "/"))
				}

				So(namespaces, ShouldContain, "intel/facter/timezone")
				So(namespaces, ShouldContain, "intel/facter/operatingsystem")
				So(namespaces, ShouldContain, "intel/facter/blockdevices")
				So(namespaces, ShouldContain, "intel/facter/uptime_seconds")
				So(namespaces, ShouldContain, "intel/facter/id")
				So(namespaces, ShouldContain, "intel/facter/osfamily")
				So(namespaces, ShouldContain, "intel/facter/kernelrelease")
				So(namespaces, ShouldContain, "intel/facter/facterversion")
				So(namespaces, ShouldContain, "intel/facter/hardwaremodel")
				So(namespaces, ShouldContain, "intel/facter/hostname")
				So(namespaces, ShouldContain, "intel/facter/memorysize")
				So(namespaces, ShouldContain, "intel/facter/interfaces")
				So(namespaces, ShouldContain, "intel/facter/ipaddress")
				So(namespaces, ShouldContain, "intel/facter/kernel")
				So(namespaces, ShouldContain, "intel/facter/os/name")
				So(namespaces, ShouldContain, "intel/facter/physicalprocessorcount")
				So(namespaces, ShouldContain, "intel/facter/memorysize_mb")
				So(namespaces, ShouldContain, "intel/facter/kernelmajversion")
				So(namespaces, ShouldContain, "intel/facter/kernelrelease")
				So(namespaces, ShouldContain, "intel/facter/processors/models/0")
			})
		})

	})
}
