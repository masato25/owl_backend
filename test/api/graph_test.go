package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGraph(t *testing.T) {
	host := "http://localhost:3000"
	Convey("Get Endpoint Failed", t, func() {
		resp, _ := resty.R().Get(fmt.Sprintf("%s/endpoint", host))
		So(resp.StatusCode(), ShouldEqual, 400)
	})

	Convey("Get Endpoint List", t, func() {
		resp, _ := resty.R().SetQueryParam("q", "a.+").Get(fmt.Sprintf("%s/endpoint", host))
		So(resp.StatusCode(), ShouldEqual, 200)
	})

	Convey("Get Counter Failed", t, func() {
		resp, _ := resty.R().Get(fmt.Sprintf("%s/endpoint_counter", host))
		So(resp.StatusCode(), ShouldEqual, 400)
	})

	Convey("Get Counter List", t, func() {
		resp, _ := resty.R().SetQueryParam("eid", "6,7").Get(fmt.Sprintf("%s/endpoint_counter", host))
		So(resp.StatusCode(), ShouldEqual, 200)
	})

}
