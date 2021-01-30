package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//定义结构体

//Mycfg mysql struct
type Mycfg struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}

//Redcfg redis struct
type Redcfg struct {
	Host       string `ini:"host"`
	Port       int    `ini:"port"`
	Password   string `ini:"password"`
	Appendonly bool   `ini:"appendonly"`
}

//Cfg config struct
type Cfg struct {
	Mysqlcfg Mycfg  `ini:"mysql"`
	Rediscfg Redcfg `ini:"redis"`
}

//Setcfg set value to struct
func Setcfg(ptr interface{}, path string) (err error) {
	var fileobj []byte
	//根据路径将文件内容读进变量fileobj
	fileobj, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	//将文件内容按照行分割，进入切片变量
	arrayfile := strings.Split(string(fileobj), "\r\n")
	var structname string
	var structtag string
	vStruct := reflect.ValueOf(ptr).Elem()
	//tStruct := reflect.TypeOf(ptr).Elem()
	//遍历每一行
	for i := 0; i < len(arrayfile); i++ {
		//设置行数变量
		linenum := i + 1
		//删除行内容开头和结尾的空格
		linecontent := strings.TrimSpace(arrayfile[i])
		//；和#开头的行数跳过
		if strings.HasPrefix(linecontent, ";") || strings.HasPrefix(linecontent, "#") {
			continue
		}
		//寻找有[ 或]的行
		if strings.Index(linecontent, "[") != -1 || strings.Index(linecontent, "]") != -1 {
			//开头为[结尾为]，格式正确，否则返回格式错误
			if strings.HasPrefix(linecontent, "[") && strings.HasSuffix(linecontent, "]") {
				//判断[]中长度不为0格式正确，否则返回错误
				initagname := linecontent[1 : len(linecontent)-1]
				if len(initagname) != 0 {
					//tStruct := reflect.TypeOf(ptr).Elem()
					//vStruct := reflect.ValueOf(ptr).Elem()
					//遍历一层结构体
					for y := 0; y < vStruct.NumField(); y++ {
						//判断一层结构体中的字段类型是否为struct
						if vStruct.Field(y).Type().Kind() == reflect.Struct {
							//判断一层结构体中字段的tag名是否和ini中[]内容一致，一致则取得字段名
							structtag = strings.TrimSpace(vStruct.Type().Field(y).Tag.Get("ini"))
							if structtag == initagname {
								structname = vStruct.Type().Field(y).Name
							}
						} else {
							err = fmt.Errorf("it is not reflect.Struct type in your type")
							return
						}
					}
				} else {
					err = fmt.Errorf("there is nothing in []!at line:%d", linenum)
				}
			} else {
				err = fmt.Errorf("format wrong at line:%d", linenum)
				return
			}
			//遍历ini文件中[]下面的行
		} else {
			//空行跳过
			if len(linecontent) == 0 {
				continue
			}
			//不包含=的行返回错误
			if strings.Index(linecontent, "=") == -1 {
				err = fmt.Errorf("format wrong need =,at line:%d", linenum)
				return
			} else {
				tmpindex := strings.Index(linecontent, "=")
				key := strings.TrimSpace(linecontent[:tmpindex])
				value := strings.TrimSpace(linecontent[tmpindex+1:])
				//fmt.Println(structname, key, value)
				//遍历二层结构体
				for x := 0; x < vStruct.FieldByName(structname).NumField(); x++ {
					vvtagname := strings.TrimSpace(vStruct.FieldByName(structname).Type().Field(x).Tag.Get("ini"))
					//判断二层结构体中的tag是否和行中的key相等
					if vvtagname == key {
						vvtype := vStruct.FieldByName(structname).Type().Field(x).Type.Kind()
						//判断二层结构体中字段类型，并根据类型对二层结构体字段赋值
						if vvtype == reflect.String {
							vStruct.FieldByName(structname).Field(x).SetString(value)
						}
						if vvtype == reflect.Int || vvtype == reflect.Int32 || vvtype == reflect.Int64 || vvtype == reflect.Int8 {
							var intvalue int64
							intvalue, err = strconv.ParseInt(value, 10, 64)
							if err != nil {
								err = fmt.Errorf("turn string to int err,at line:%d", linenum)
								return
							}
							vStruct.FieldByName(structname).Field(x).SetInt(intvalue)
						}
						if vvtype == reflect.Bool {
							var boolvalue bool
							boolvalue, err = strconv.ParseBool(value)
							if err != nil {
								err = fmt.Errorf("turn string to bool err,at line:%d", linenum)
								return
							}
							vStruct.FieldByName(structname).Field(x).SetBool(boolvalue)
						}
					}
				}
			}
		}

	}
	return
}

func main() {
	var testcfg Cfg
	var path string
	path = "./test.ini"
	err := Setcfg(&testcfg, path)
	if err != nil {
		fmt.Println("there is err:", err)
	}
	fmt.Printf("testcfg = %#v\n", testcfg)
	fmt.Println(testcfg.Mysqlcfg.Password)
}
