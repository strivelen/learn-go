// 生产者消费者模式，需求分析：
// 生产者每秒生产一个商品，并通知物流公司取货
// 物流公司将商品运输到商铺
// 消费者阻塞等待商铺到货，需要消费10次商品

package main

import (
	"fmt"
	"strconv"
	_ "time"
)

type Product struct {
	Name string
}

func Producer(storageChannel chan<- Product, count int) {
	for {
		producer := Product{"商品：" + strconv.Itoa(count)}
		storageChannel <- producer
		count--
		fmt.Println("生产者", producer.Name)
		if count == 0 {
			return
		}
	}
}

func Logistics(storageChannel <-chan Product, shopChannel chan<- Product) {
	for {
		Product := <-storageChannel
		shopChannel <- Product
		fmt.Println("物流公司", Product.Name)
	}
}

func Consumer(shopChannel <-chan Product, count int, exitChannel chan<- bool) {
	for {
		Product := <-shopChannel
		fmt.Println("消费者", Product.Name)
		count--
		if count == 0 {
			exitChannel <- true
			return
		}
	}
}

func main() {
	storageChannel := make(chan Product, 10)
	shopChannel := make(chan Product, 10)
	exitChannel := make(chan bool, 1)

	go Producer(storageChannel, 10)
	go Logistics(storageChannel, shopChannel)
	go Consumer(shopChannel, 10, exitChannel)

	if <-exitChannel {
		return
	}
}
