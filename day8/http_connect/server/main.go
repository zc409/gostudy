package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	mes, err := ioutil.ReadFile("./test.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("read file failed,err:%v\n", err)))
	}
	w.Write([]byte(mes))
}

func f2(w http.ResponseWriter, r *http.Request) {
	mes, err := ioutil.ReadFile("./index.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("read file failed,err:%v\n", err)))
	}
	w.Write([]byte(mes))
}

func f3(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL)
	// fmt.Println(r.Method)
	// fmt.Println(ioutil.ReadAll(r.Body))
	//对于get请求，参数都放在url上（query parm），请求体中是没有数据的。
	queryParam := r.URL.Query() //自动帮我们识别url中query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	w.Write([]byte(fmt.Sprintln("name:", name)))
	w.Write([]byte(fmt.Sprintln("age:", age)))
	w.Write([]byte(fmt.Sprintln("url:", r.URL)))
	w.Write([]byte(fmt.Sprintln("method:", r.Method)))
	body, _ := ioutil.ReadAll(r.Body)
	w.Write([]byte("body:"))
	w.Write(body)
	w.Write([]byte("\n"))
	w.Write([]byte("ok\n"))
}

func main() {
	http.HandleFunc("/secret/", f1)
	http.HandleFunc("/", f2)
	http.HandleFunc("/client/", f3)
	http.ListenAndServe("192.168.3.2:8080", nil)
}
