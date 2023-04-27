//响应体
func responseBody(r *http.Response)  {
	contents,err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n",contents)
}
//状态吗
func status(r *http.Response)  {
	fmt.Println(r.StatusCode) // 200
	fmt.Println(r.Status)   // 200 ok
}
//响应头
func header(r *http.Response)  {
	fmt.Println(r.Header)
}
//编码
func encoding(r *http.Response)  {
	//1.content-type中提供编码信息、Content-Type:"application/json;charset=utf-8"
	//2.html head meta中获取编码、<meta http-equiv=content-type content-type="application/json;charset=utf-8" ...
	//3.根据网页的头部猜测网页编码信息
	// go get golang.org/x/net/html
	// 获取网页前1024字节来猜测，多余的丢掉
	bufReader := bufio.NewReader(r.Body)
	bytes,err := bufReader.Peek(1024)   //不会移动reader的读取位置，前1024还能读到
	if err != nil {
		log.Fatal(err)
	}
	e,_,_ := charset.DetermineEncoding(bytes,r.Header.Get("content-type"))
	fmt.Println(e)
	// 转码 go get golang.org/x/text/transform
	// 旧reader转码为e编码格式的newReader
	// 然后读取就可以了
	newReader := transform.NewReader(bufReader,e.NewDecoder())
	contents,err := io.ReadAll(newReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n",contents)
}
