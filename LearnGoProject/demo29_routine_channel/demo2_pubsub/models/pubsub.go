package models

import (
	"fmt"
	"sync"
	"time"
)

/*
发布订阅模型 (注意操作的同步及发布的一致性)
1. 每次订阅添加一个chan作为key来存储订阅者的信息, value来储存筛选得func, 并返回只读得chan
2. 发布的时候, 向每个订阅者发送信息. (向chan发送数据)
*/

type (
	subcriber chan interface{}  //todo 订阅者为一个管道
	topicFunc func (v interface{}) bool //todo 过滤器
)

type Publisher struct {
	m sync.Mutex
	buffer int
	timeout time.Duration
	subcribers map[subcriber]topicFunc
}

func NewPublisher(timeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer: buffer,
		timeout: timeout,
		subcribers: make(map[subcriber]topicFunc),
	}
}

//todo 添加一个订阅
func (this *Publisher)Subscribe() <-chan interface{} {
	return this.SubscribeTopic(nil)
}

//todo 添加一个带条件的订阅
func (this *Publisher)SubscribeTopic(topic topicFunc) <-chan interface{} {
	ch := make(subcriber, this.buffer)
	this.m.Lock()
	this.subcribers[ch] = topic
	this.m.Unlock()
	return ch
}

//todo 发布一个内容
func (this *Publisher)Publish(v interface{}) {
	this.m.Lock()
	defer this.m.Unlock()

	var wg sync.WaitGroup
	for sub, topic := range this.subcribers {
		wg.Add(1)
		this.SendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

//todo 向指定的订阅者发布一个内容
func (this *Publisher)SendTopic(sub subcriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {

	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	select {
		case sub <- v:
			fmt.Println("向管道写入值:", v)
		case <-time.After(this.timeout):
			//todo: 如上面得管道一直堵塞, 那么执行超时
			fmt.Println("向管道写入超时:", v)
	}
}


//todo 关闭指定订阅
func (this *Publisher)Evict(sub subcriber) {
	this.m.Lock()
	defer this.m.Unlock()

	delete(this.subcribers, sub)
	close(sub)
}
//todo 关闭所有订阅
func (this *Publisher)Close() {
	this.m.Lock()
	defer this.m.Unlock()

	for sub := range this.subcribers {
		delete(this.subcribers, sub)
		close(sub)
	}
}



