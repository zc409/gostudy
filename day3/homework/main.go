package main

import "fmt"

/* 有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a.名字中每包含1个'e'或'E'分1枚金币
b.名字中每包含1个'i'或'I'分2枚金币
c.名字中每包含1个'o'或'O'分3枚金币
d.名字中每包含1个'u'或'U'分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？ */

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for k, v := range distribution {
		fmt.Println(k, v)
	}
}

func dispatchCoin() int {
	leftcount := coins
	for x := 0; x < len(users); x++ {
		counts := 0
		for y := 0; y < len(users[x]); y++ {
			// if users[x][y] == 'e' || users[x][y] == 'E' {
			// 	counts++
			// } else if users[x][y] == 'i' || users[x][y] == 'I' {
			// 	counts += 2
			// } else if users[x][y] == 'o' || users[x][y] == 'O' {
			// 	counts += 3
			// } else if users[x][y] == 'u' || users[x][y] == 'U' {
			// 	counts += 4
			// }
			switch users[x][y] {
			case 'e', 'E':
				counts++
			case 'i', 'I':
				counts += 2
			case 'o', 'O':
				counts += 3
			case 'u', 'U':
				counts += 4
			}
		}
		leftcount -= counts
		distribution[users[x]] = counts
	}

	return leftcount
}
