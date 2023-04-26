package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

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

func main() {
	del()
}