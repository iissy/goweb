package srv

import (
	"context"
	"github.com/iissy/goweb/src/utils"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/config"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/sirupsen/logrus"
	"time"
)

type Hrefs struct{}

func Start() {
	urls := utils.GetConsulUrls()
	reg := consul.NewRegistry(registry.Addrs(urls...))

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name(config.Get("srv").String("micro.hrefs.srv")),
		micro.WrapHandler(logWrapper),
	)
	server.Init()
	service.Server().Init(server.Wait(nil))
	micro.RegisterHandler(service.Server(), new(Hrefs))
	service.Run()
}

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	start := time.Now()
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := fn(ctx, req, rsp)
		utils.WriteErrorLog(req.Endpoint(), err)

		logrus.Infof("%s %s", time.Since(start), req.Endpoint())
		return err
	}
}
