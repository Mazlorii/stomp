package main

import (
	"fmt"
	"server/pkg/packets"

	"google.golang.org/protobuf/proto"
)

func main() {
	/*
		packet := &packets.Packet{
			SenderId: 69,
			Msg:      packets.NewChat("Hello, world!"),
		}

		fmt.Println(packet)

		data, err := proto.Marshal(packet)

		if err != nil {
			panic(err)
		}

		fmt.Println(data)
	*/

	data := []byte{8, 69, 18, 15, 10, 13, 72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}

	packet := &packets.Packet{}
	err := proto.Unmarshal(data, packet)

	if err != nil {
		panic(err)
	}

	fmt.Println(packet)

}
