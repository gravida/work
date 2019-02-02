package utils

import (
	ers "errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
)

const (
	wxUrl = "https://api.weixin.qq.com/sns/jscode2session?grant_type=authorization_code"
)

type WeiXinData struct {
	SessionKey string
	Openid     string
}

// var wx_url = fmt.Sprintf("%s?appid=%s&secret=%s&grant_type=authorization_code%js_code=", wxUrl, wx_appid, wx_secret)
func WeiXinLogin(code, appid, secret string) (*WeiXinData, error) {
	s := fmt.Sprintf("%s&appid=%s&secret=%s&js_code=%s", wxUrl, appid, secret, code)
	resp, err := http.Get(s)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	js, err := simplejson.NewJson(body)

	if err != nil {
		return nil, err
	}

	fmt.Println(js)
	openid := js.Get("openid").MustString()
	if len(openid) == 0 {
		errmsg := js.Get("errmsg").MustString()
		if len(errmsg) == 0 {
			return nil, ers.New("unknow Error")
		} else {
			return nil, ers.New(errmsg)
		}
	}

	var data = &WeiXinData{}
	data.SessionKey = js.Get("session_key").MustString()
	data.Openid = openid
	return data, nil
}
