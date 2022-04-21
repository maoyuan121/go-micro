package main

import (
	"fmt"
	"go-micro.dev/v4/cmd/protoc-gen-micro/plugin/micro"
	"os"

	"context"

	proto "github.com/asim/go-micro/examples/v4/service/proto"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
)

/*

Example usage of top level service initialisation

*/

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

// Setup and the client
func runClient(service micro.Service) {
	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Greeting)
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),

		// 设置些 flags。 --run_client 运行 client

		// 添加 runtime flags
		// We could do this below too
		micro.Flags(&cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	// Init 将解析命令行 flags。
	// 任何  flag 将覆盖上面的设置。
	// 这里定义的选项将覆盖命令行上设置的任何选项。
	service.Init(
		// Add runtime action
		// We could actually do this above
		micro.Action(func(c *cli.Context) error {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
			return nil
		}),
	)

	// By default we'll run the server unless the flags catch us

	// Setup the server

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
