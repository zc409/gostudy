package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//直接将发送请求完整地址写进http.Get()中
// func main() {
// 	resp, err := http.Get("http://192.168.3.2:8080/client/?name=ha&age=18")
// 	if err != nil {
// 		fmt.Println("get url failed,err:", err)
// 		return
// 	}
// 	//从resp中把服务端返回的数据读出来
// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Printf("read resp.Body failed,err%v\n", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	fmt.Println(string(b))
// }

//定义全局的client请求，如果打开keepalived可以不停复用之前的连接
// var (
// 	client =http.Client{
// 		Transport: &http.Transport{
// 			DisableKeepAlives: true,
// 		},
// 	}
// )

//构造http请求
func main() {
	data := url.Values{} //url values
	data.Set("name", "小明")
	data.Set("age", "18")
	querystr := data.Encode() //  encode之后的参数
	urlobj, _ := url.Parse("http://192.168.3.2:8080/client/")
	urlobj.RawQuery = querystr                               //将请求参数添加到urlobj中
	req, err := http.NewRequest("GET", urlobj.String(), nil) //根据方法和urlobj来构造req请求结构体
	if err != nil {
		fmt.Printf("create request failed,err:%v\n", err)
		return
	}

	//使用默认的方式发送客户端连接请求
	//resp, err := http.DefaultClient.Do(req) //用req请求结构体对服务端发起请求，获取resp回应

	//自定义客户端请求方式，比如禁用长连接keepalived,预防连接过多导致无法连接
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("response failed,err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
