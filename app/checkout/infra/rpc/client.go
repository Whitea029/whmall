package rpc

import (
	"sync"

	"github.com/Whitea029/whmall/app/checkout/conf"
	checkoutUtils "github.com/Whitea029/whmall/app/checkout/utils"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	CartCLient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoutUtils.MustHandleError(err)
}

func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoutUtils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	CartCLient, err = cartservice.NewClient("cart", opts...)
	checkoutUtils.MustHandleError(err)
}