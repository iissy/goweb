package srv

import (
	"context"
	"time"

	gCli "github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	gSrv "github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/config"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/iissy/goweb/src/utils"
	"github.com/sirupsen/logrus"
)

type Hrefs struct{}

func init() {
	urls := utils.GetConsulUrls()
	reg := consul.NewRegistry(registry.Addrs(urls...))

	server.DefaultServer = gSrv.NewServer(
		server.Registry(reg),
		server.Name(config.Get("srv").String("micro.hrefs.srv")),
	)

	client.DefaultClient = gCli.NewClient(
		client.Registry(reg),
	)
}

func Start() {
	service := micro.NewService(
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
