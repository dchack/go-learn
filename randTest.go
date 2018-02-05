package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(20) // Try changing this number!
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	// 新设置Seed 才可以得到新的随机数
	rand.Seed(21)
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])

	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")

	s2 := rand.NewSource(42) //同前面一样的种子
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")

	// 随机产生字符串
	Init()
	fmt.Println("Magic 8-Ball says1:", RandStringRunes(2))

}


func Init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}