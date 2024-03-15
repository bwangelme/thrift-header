package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"thrift_header/gen-go/tutorial"
)

var defaultCtx = context.Background()

func handleClient(client *tutorial.CalculatorClient) (err error) {
	sum, _ := client.Add(defaultCtx, 1, 1)
	fmt.Print("1+1=", sum, "\n")
	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration) error {
	var transport thrift.TTransport
	if secure {
		transport = thrift.NewTSSLSocketConf(addr, cfg)
	} else {
		transport = thrift.NewTSocketConf(addr, cfg)
	}
	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return handleClient(tutorial.NewCalculatorClient(thrift.NewTStandardClient(iprot, oprot)))
}
