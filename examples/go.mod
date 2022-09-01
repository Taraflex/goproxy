module github.com/taraflex/goproxy/examples/goproxy-transparent

go 1.18

require (
	github.com/gorilla/websocket v1.5.0
	github.com/inconshreveable/go-vhost v0.0.0-20160627193104-06d84117953b
	github.com/taraflex/goproxy v0.0.0-20181111060418-2ce16c963a8a
	github.com/taraflex/goproxy/ext v0.0.0-20220901020949-51bbc4890db1
)

require github.com/rogpeppe/go-charset v0.0.0-20180617210344-2471d30d28b4 // indirect

replace github.com/taraflex/goproxy => ../
