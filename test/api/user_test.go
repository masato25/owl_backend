package test

import (
	"fmt"
	"net/http"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/elgs/gojq"
	"github.com/franela/goreq"
	"github.com/go-resty/resty"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	host := "http://localhost:3000"
	Convey("Get User Login Failed", t, func() {
		resp, _ := resty.R().SetQueryParam("name", "gg123").
			SetQueryParam("name", "root").
			SetQueryParam("password", "willnotpass").
			Post(fmt.Sprintf("%s/user/login", host))
		So(resp.StatusCode(), ShouldEqual, 400)
	})
	var sig interface{}
	Convey("Get User Login Success", t, func() {
		resp, _ := resty.R().SetQueryParam("name", "test2").
			SetQueryParam("password", "test2").
			Post(fmt.Sprintf("%s/user/login", host))
		log.Info("result: ", resp.String())
		jss, err := gojq.NewStringQuery(resp.String())
		if err != nil {
			log.Info(err.Error(), resp.String())
		}
		sig, err = jss.Query("sig")
		if err != nil {
			log.Info(err.Error(), resp.String())
		}
		log.Info("sig: ", sig)
		So(resp.StatusCode(), ShouldEqual, 200)
	})
	Convey("Test Logout Session", t, func() {
		resp, err := goreq.Request{
			Uri: fmt.Sprintf("%s/user/logout", host),
		}.
			WithCookie(&http.Cookie{Name: "name", Value: "test2"}).
			WithCookie(&http.Cookie{Name: "sig", Value: sig.(string)}).
			Do()
		if err != nil {
			log.Error(err.Error())
		}
		// cookie1 := &http.Cookie{
		// 	Name:     "name",
		// 	Value:    "test2",
		// 	Path:     "/",
		// 	Domain:   "localhost",
		// 	HttpOnly: true,
		// 	Secure:   false,
		// }
		// cookie2 := &http.Cookie{
		// 	Name:     "sig",
		// 	Value:    sig.(string),
		// 	Path:     "/",
		// 	Domain:   "localhost",
		// 	HttpOnly: true,
		// 	Secure:   false,
		// }
		//
		// resty.SetCookies([]*http.Cookie{cookie2, cookie1})
		// resp, err := resty.R().Get(fmt.Sprintf("%s/user/logout", host))
		// if err != nil {
		// 	log.Info(err.Error())
		// }
		log.Info(resp.Body.ToString())
		So(resp.StatusCode, ShouldEqual, 200)
	})
}
