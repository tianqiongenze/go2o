package msq

import (
	"errors"
	"fmt"
)

var (
	producer Producer
)

const KAFKA = 1
const REDIS = 2

type Producer interface {
	// 推送消息
	Push(topic string, key string, message string) error
	// 关闭生产者
	Close()
}

// 设置
func Configure(mqType int, address []string) error {
	if mqType == KAFKA {
		producer = newKafkaProducer(address)
		return nil
	}
	return errors.New(fmt.Sprintf("not implement mq type %d", mqType))
}

// 推送消息
func Push(topic string, key string, message string) error {
	return producer.Push(topic, key, message)
}

// 关闭生产者
func Close() {
	producer.Close()
}