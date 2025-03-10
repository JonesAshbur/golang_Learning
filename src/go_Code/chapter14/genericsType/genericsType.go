package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

type user struct {
	ID   int64
	Name string
	Age  uint8
}
type address struct {
	ID       int
	Province string
	City     string
}

// 集合转列表
func mapToList[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, v := range mp {
		list[i] = v
		i++
	}
	return list
}
func myPrint[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TtypeCase() {
	userMap := make(map[int64]user, 0)
	userMap[1] = user{ID: 1, Name: "jones", Age: 18}
	userMap[2] = user{ID: 2, Name: "ash", Age: 28}
	userList := mapToList[int64, user](userMap)
	ch1 := make(chan user)
	go myPrint(ch1)
	for _, userval := range userList {
		ch1 <- userval
	}

	addressMap := make(map[int64]address, 0)
	addressMap[1] = address{ID: 1, Province: "直辖市", City: "北京"}
	addressMap[2] = address{ID: 2, Province: "直辖市", City: "天津"}
	addressList := mapToList[int64, address](addressMap)
	ch2 := make(chan address)
	go myPrint(ch2)
	for _, addressval := range addressList {
		ch2 <- addressval
	}
}

// 泛型切片
type List[T any] []T

// 泛型集合
type MapT[k comparable, V any] map[k]V

// 泛型channel
type Chan[T any] chan T

func TtypeCase1() {
	userMap := make(MapT[int64, user], 0)
	userMap[1] = user{ID: 1, Name: "jones", Age: 18}
	userMap[2] = user{ID: 2, Name: "ash", Age: 28}
	var userList List[user]
	userList = mapToList[int64, user](userMap)
	ch1 := make(Chan[user])
	go myPrint(ch1)
	for _, userval := range userList {
		ch1 <- userval
	}

	addressMap := make(MapT[int, address], 0)
	addressMap[1] = address{ID: 1, Province: "直辖市", City: "北京"}
	addressMap[2] = address{ID: 2, Province: "直辖市", City: "天津"}
	var addressList List[address]
	addressList = mapToList[int, address](addressMap)
	ch2 := make(Chan[address])
	go myPrint(ch2)
	for _, addressval := range addressList {
		ch2 <- addressval
	}
}
func main() {
	TtypeCase()
	//泛型集合
	TtypeCase1()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
