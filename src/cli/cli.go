package cli

import (
	"context"
	"fmt"

	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/config"
)

var name string

func init() {
	name = config.Get("srv").String("micro.hrefs.srv")
}

func Call(method string, req interface{}, rsp interface{}) error {
	request := client.NewRequest(name, fmt.Sprintf("Hrefs.%s", method), req, client.WithContentType("application/json"))

	if err := client.Call(context.TODO(), request, &rsp); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
