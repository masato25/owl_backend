package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/masato25/resty"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGraph(t *testing.T) {
	host := "http://localhost:3000/api/v1/graph"
	cname := "test1"
	csig := "d4f71cba377911e699d60242ac110010"
	rt := resty.New()
	rt.SetCookies([]*http.Cookie{rt.MakeTestCookie("name", cname), rt.MakeTestCookie("sig", csig)})
	Convey("Get Endpoint Failed", t, func() {
		resp, _ := rt.R().Get(fmt.Sprintf("%s/endpoint", host))
		So(resp.StatusCode(), ShouldEqual, 400)
	})
	Convey("Get Endpoint without login session", t, func() {
		resp, _ := resty.R().Get(fmt.Sprintf("%s/endpoint", host))
		So(resp.StatusCode(), ShouldEqual, 401)
	})

	Convey("Get Endpoint List", t, func() {
		resp, _ := rt.R().SetQueryParam("q", "a.+").Get(fmt.Sprintf("%s/endpoint", host))
		So(resp.StatusCode(), ShouldEqual, 200)
	})

	Convey("Get Counter Failed", t, func() {
		resp, _ := rt.R().Get(fmt.Sprintf("%s/endpoint_counter", host))
		So(resp.StatusCode(), ShouldEqual, 400)
	})

	Convey("Get Counter List", t, func() {
		resp, _ := rt.R().SetQueryParam("eid", "6,7").Get(fmt.Sprintf("%s/endpoint_counter", host))
		So(resp.StatusCode(), ShouldEqual, 200)
	})

}
