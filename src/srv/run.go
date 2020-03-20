package srv

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"hrefs.cn/src/utils"
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
	)

	service.Init()
	micro.RegisterHandler(service.Server(), new(Hrefs))
	service.Run()
}
