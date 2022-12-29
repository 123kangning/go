package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc/hello-server/proto"
	"log"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "kuangshen",
		"appKey": "123123",
	}, nil
}
func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return true
}
func main() {
	credentials, _ := credentials.NewClientTLSFromFile("/home/kangning/go/src/grpc/key/test.pem",
		"*.kuangstudy.com")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(credentials))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(credentials))
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatalf("didn't connect:%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("close error")
		}
	}(conn)

	//建立连接
	client := pb.NewSayHelloClient(conn)
	//执行rpc调用（这个方法在服务端实现并返回结果）
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "kangning"})
	fmt.Println(resp.GetResponseMsg())

}
