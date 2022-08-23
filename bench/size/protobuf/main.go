package main

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	fmt.Println(proto.Marshal(timestamppb.New(time.Now())))
}
