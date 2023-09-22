package otter

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"otter-dingtalk/internal/global"
	"strings"
	"time"
)

var cookies string
var client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
	Timeout: time.Second * 10,
}

func Login() {
	loginURL := global.OTTER_URL + "/login.htm"
	loginData := url.Values{
		"action":                {"user_action"},
		"event_submit_do_login": {"1"},
		"_fm.l._0.n":            {global.OTTER_USERNAME},
		"_fm.l._0.p":            {global.OTTER_PASSWORD},
	}
	resp, err := client.PostForm(loginURL, loginData)
	resp.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		global.GL_LOG.Error("登录请求发送失败:", err)
		os.Exit(2)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusFound {
		global.GL_LOG.Error("登录失败，状态码:", resp.StatusCode)
		os.Exit(2)
	}
	cook := []string{}
	for _, cookie := range resp.Cookies() {
		cook = append(cook, cookie.Name+"="+cookie.Value)
	}
	cookies = strings.Join(cook, ";")
}

func startchannel(channelId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	requrl := fmt.Sprintf("%s/?action=channelAction&channelId=%s&status=start&pageIndex=1&searchKey=&eventSubmitDoStatus=true", global.OTTER_URL, channelId)
	global.GL_LOG.Info(requrl)
	getReq, err := http.NewRequestWithContext(ctx, http.MethodGet, requrl, nil)
	if err != nil {
		global.GL_LOG.Error(err)
	}
	getReq.Header.Add("Cookie", cookies)
	getResp, err := client.Do(getReq)
	if err != nil {
		global.GL_LOG.Errorf("GET 请求发送失败:%s", err)
		return
	}
	if getResp.StatusCode == 302 {
		global.GL_LOG.Infof("channelId:%s 自动解挂任务请求成功", channelId)
	} else {
		global.GL_LOG.Errorf("channelId:%s 自动解挂任务请求失败", channelId)
	}
	defer getResp.Body.Close()
}
