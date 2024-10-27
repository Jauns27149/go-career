package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-interview/protobuf/pb"
)

func main() {
	// 创建 Person 消息实例
	person := &pb.Person{
		Name:   "Alice",
		Age:    30,
		Active: true,
	}

	// 序列化 Person 消息
	data, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("Serialized data:", data)

	// 反序列化 Person 消息
	var newPerson pb.Person
	if err := proto.Unmarshal(data, &newPerson); err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	// 输出反序列化后的消息
	fmt.Println("Deserialized person:")
	fmt.Printf("%+v\n", newPerson)
}
