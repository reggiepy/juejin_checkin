package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://api.juejin.cn/growth_api/v1/check_in"
	method := "POST"

	payload := strings.NewReader(`{`+""+`"aid": "2608",`+""+`"uuid": "6986922219256989198"`+""+`}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("accept-encoding", "gzip, deflate, br")
  req.Header.Add("accept-language", "zh,en;q=0.9,zh-CN;q=0.8")
  req.Header.Add("content-type", " application/json;charset=utf-8;")
  req.Header.Add("cookie", " MONITOR_WEB_ID=5516ef3f-1168-49b5-998f-d8655fb700df; _ga=GA1.2.848999477.1631242708; n_mh=Tl-vABvTTpH453pdpPTMTdwPHPUdQ5Duws89Gdmd58I; passport_csrf_token_default=909dd26560154ed7bbb002267e7084dd; passport_csrf_token=909dd26560154ed7bbb002267e7084dd; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227006141156514448927%2522%252C%2522ssid%2522%253A%25226e26364e-03ad-4601-bf59-63f9b653c474%2522%252C%2522user_unique_id%2522%253A%25227006141156514448927%2522%252C%2522timestamp%2522%253A1631244356882%257D; _tea_utm_cache_2608={%22utm_source%22:%22timeline%22%2C%22utm_medium%22:%22banner%22%2C%22utm_campaign%22:%22xiaoce_Linda_20220119%22}; _tea_utm_cache_2018=undefined; sid_guard=4cb722b2adfbbc88d4f5745c6a563c96%7C1644977333%7C5184000%7CSun%2C+17-Apr-2022+02%3A08%3A53+GMT; uid_tt=41eff0f895ee8d6429ea87559da47b5c; uid_tt_ss=41eff0f895ee8d6429ea87559da47b5c; sid_tt=4cb722b2adfbbc88d4f5745c6a563c96; sessionid=4cb722b2adfbbc88d4f5745c6a563c96; sessionid_ss=4cb722b2adfbbc88d4f5745c6a563c96; sid_ucp_v1=1.0.0-KGJkOTNjN2UwMTkxNjA3OThmMjE2MGQ4MDVkZmRiODljODA0Nzg0NjMKFwjIi5DA_fXxBhC1ubGQBhiwFDgCQO8HGgJsZiIgNGNiNzIyYjJhZGZiYmM4OGQ0ZjU3NDVjNmE1NjNjOTY; ssid_ucp_v1=1.0.0-KGJkOTNjN2UwMTkxNjA3OThmMjE2MGQ4MDVkZmRiODljODA0Nzg0NjMKFwjIi5DA_fXxBhC1ubGQBhiwFDgCQO8HGgJsZiIgNGNiNzIyYjJhZGZiYmM4OGQ0ZjU3NDVjNmE1NjNjOTY")
  req.Header.Add("origin", " https://juejin.cn")
  req.Header.Add("referer", " https://juejin.cn/")
  req.Header.Add("sec-ch-ua", " \" Not A;Brand\";v=\"99\", \"Chromium\";v=\"98\", \"Google Chrome\";v=\"98\"")
  req.Header.Add("sec-ch-ua-mobile", " ?0")
  req.Header.Add("sec-ch-ua-platform", " \"Windows\"")
  req.Header.Add("sec-fetch-dest", " empty")
  req.Header.Add("sec-fetch-mode", " cors")
  req.Header.Add("sec-fetch-site", " same-site")
  req.Header.Add("user-agent", " Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}