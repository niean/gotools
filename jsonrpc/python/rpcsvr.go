package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// RPC
type MetricValue struct {
	Endpoint  string      `json:"endpoint"`
	Metric    string      `json:"metric"`
	Value     interface{} `json:"value"`
	Step      int64       `json:"step"`
	Type      string      `json:"counterType"`
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
}

type TransferResponse struct {
	Message string
	Total   int
	Invalid int
	Latency int64
}

func (this *TransferResponse) String() string {
	return fmt.Sprintf(
		"<Total=%v, Invalid:%v, Latency=%vms, Message:%s>",
		this.Total,
		this.Invalid,
		this.Latency,
		this.Message,
	)
}

type Transfer int

func (t *Transfer) Update(args []*MetricValue, reply *TransferResponse) error {
	reply.Invalid = 0
	reply.Latency = 0
	reply.Total = len(args)
	reply.Message = "ok"

	for _, mv := range args {
		log.Printf("%+v\n", *mv)
	}

	return nil
}

// MAIN
func main() {
	l, err := net.Listen("tcp", "127.0.0.1:5090")
	defer l.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("listening:", l.Addr())

	trans := new(Transfer)
	rpc.Register(trans)

	for {
		log.Print("waiting for connections...")
		c, err := l.Accept()

		if err != nil {
			log.Printf("accept error: %s", c)
			continue
		}

		log.Printf("connection started: %v", c.RemoteAddr())
		go jsonrpc.ServeConn(c)
	}
}
