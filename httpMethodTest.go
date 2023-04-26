// method get
func get() {
	r,err := http.Get("http://httpbin.org/get")
	defer r.Body.Close()
	if err!=nil{
		log.Fatal(err)
	}
	contents,err := io.ReadAll(r.Body)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s\n",contents)
	fmt.Println("STATUS:",r.StatusCode)
}

// method post
func post() {
	r,err := http.Post("http://httpbin.org/post","",nil)
	defer r.Body.Close()
	if err!=nil{
		log.Fatal(err)
	}
	contents,err := io.ReadAll(r.Body)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s\n",contents)
	fmt.Println("STATUS:",r.StatusCode)
}

// method put
func put() {
	request,err := http.NewRequest(http.MethodPut ,"http://httpbin.org/put",nil)
	if err != nil {
		log.Fatal(err)
	}
	r,err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	contents,err := io.ReadAll(r.Body)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s\n",contents)
	fmt.Println("STATUS:",r.StatusCode)
}

//method del
func del() {
	request,err := http.NewRequest(http.MethodDelete ,"http://httpbin.org/delete",nil)
	if err != nil {
		log.Fatal(err)
	}
	r,err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	contents,err := io.ReadAll(r.Body)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s\n",contents)
	fmt.Println("STATUS:",r.StatusCode)
}

//methed get with params
func getRequestByParams() {
	request,err := http.NewRequest(http.MethodGet,"http://httpbin.org/get",nil)
	if err != nil {
		log.Fatal(err)
	}
	params := make(url.Values)
	params.Add("name","cyd")
	params.Add("age","18")
	// 添加参数到url
	// fmt.Println(params.Encode())  输出：age=18&name=cyd
	request.URL.RawQuery = params.Encode()
	// fmt.Println(request.URL)  输出：http://httpbin.org/get?age=18&name=cyd
	r,err :=  http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	body,err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n",body)
}

// method get with header
func getRequestByHeader() {
	request,err := http.NewRequest(http.MethodGet,"http://httpbin.org/get",nil)
	if err != nil {
		log.Fatal()
	}
	request.Header.Add("User-Agent","chrome")
	r,err :=  http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	body,err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n",body)
}
