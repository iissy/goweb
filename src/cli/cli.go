package cli

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config"
)

var cli client.Client
var name string

func init() {
	service := micro.NewService()
	service.Init()
	cli = service.Client()
	name = config.Get("srv").String("micro.hrefs.srv")
}

func Call(method string, req interface{}, rsp interface{}) error {
	request := cli.NewRequest(name, fmt.Sprintf("Hrefs.%s", method), req, client.WithContentType("application/json"))

	if err := cli.Call(context.TODO(), request, &rsp); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
