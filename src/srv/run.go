package srv

import (
	"context"
	"github.com/kataras/golog"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-plugins/registry/consul/v2"
	"hrefs.cn/src/utils"
	"time"
)

type Hrefs struct{}

func Start() {
	urls := utils.GetConsulUrls()
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = urls
	})

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

		golog.Infof("%s %s", time.Since(start), req.Endpoint())
		return err
	}
}
