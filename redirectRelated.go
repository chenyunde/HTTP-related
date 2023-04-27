// 限制重定向次数
func redirectLimitTimes()  {
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) > 10 {
				return errors.New("redirect too times")
			}
			return nil
		},
	}
	request,err := http.NewRequest(http.MethodGet,"http://httpbin.org/redirect/20",nil)
	if err != nil {
		log.Fatal(err)
	}
	_,err = client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
}
// 禁止重定向
func redirectForbidden() {
	//防止重定向到首页
	//定义client
	/*client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	 */
	request,err := http.NewRequest(http.MethodGet,"http://httpbin.org/cookies/set?name=cyd",nil)
	if err != nil {
		log.Fatal(err)
	}
	resp,err := http.DefaultClient.Do(request)
	//resp,err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Request.URL)
}
