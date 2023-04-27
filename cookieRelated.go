func rrCookie()  {
	// 模拟完成一个登陆
	// 请求一个页面，传递基本的额登陆信息，将响应的cookie设置到下一次的请求中重新请求
	// 请求http://httpbin.org/cookies/set?name=cyd&password=123456
	// 再一次请求时携带cookie (默认的client在重定向请求时不会携带cookie)
	// 会重定向到http://httpbin.org/cookies 并返回cookie
  // 1.自定义client，设置禁止重定向
	// 2.完成第一次请求并保存返回的cookie
	// 3.将第一次返回的cookie添加的第二次请求的cookie中
	// 4.完成第二次请求，打印cookie
	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	firstRequest,_ := http.NewRequest(http.MethodGet,"http://httpbin.org/cookies/set?name=cyd&password=123456",nil)
	firstResp,_ := client.Do(firstRequest)
	defer firstResp.Body.Close()
	secondRequest,_ := http.NewRequest(http.MethodGet,"http://httpbin.org/cookies",nil)
	for _,cookie := range firstResp .Cookies() {
		secondRequest.AddCookie(cookie)
	}
	secondResp,_ := client.Do(secondRequest)
	defer secondResp.Body.Close()
	contents,_ := io.ReadAll(secondResp.Body)
	fmt.Printf("%s",contents)
}

// cookie 自动完成cookie的携带
func jarCookie() {
	// cookie类型：会话型（默认）、持久型
	//go get github.com/juju/persistent-cookiejar

	// jar,_ := cookiejar.New(nil)  // 会话型
	 jar,_ := cookiejar2.New(nil) // 持久型
	 client := http.Client{Jar:jar}
	 request,_ := http.NewRequest(http.MethodGet,"http://httpbin.org/cookies/set?name=cyd&password=123456",nil)
	 resp,_ := client.Do(request)
	 jar.Save() //持久型需要save保存(~/.go-cookies)，会话型不需要
	 defer resp.Body.Close()
	 io.Copy(os.Stdout,resp.Body)
}
