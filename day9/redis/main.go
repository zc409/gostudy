package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//redis

var redisdb *redis.Client

//连接redis，并测试是否成功
func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.56.101:6379",
		Password: "123456", //no password set
		DB:       0,        //use default DB
	})

	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

//set/get示例
func redisexample() {
	err := redisdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed,err:%v\n", err)
		return
	}

	val, err := redisdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed,err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := redisdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed,err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

//zset示例
func redisExample2() {
	zsetkey := "rank"
	language := []*redis.Z{
		&redis.Z{Score: 90.0, Member: "Golang"},
		&redis.Z{Score: 98.0, Member: "Java"},
		&redis.Z{Score: 95.0, Member: "Python"},
		&redis.Z{Score: 97.0, Member: "Javascript"},
		&redis.Z{Score: 99.0, Member: "C/C++"},
	}
	//zadd
	num, err := redisdb.ZAdd(zsetkey, language...).Result()
	if err != nil {
		fmt.Printf("zadd failed,err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("init redisdb failed,err:%v\n", err)
		return
	} else {
		fmt.Printf("init redisdb success\n")
	}
	//redisexample()
	redisExample2()
}
