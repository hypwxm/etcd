package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/hypwxm/rider"
	"google.golang.org/grpc"
)

var endpoints = []string{"https://11.11.11.111:2379", "https://11.11.11.112:2379", "https://11.11.11.113:2379"}

func main() {
	wd, _ := os.Getwd()
	tlsInfo := transport.TLSInfo{
		CertFile:      wd + "/tls/client.pem",
		KeyFile:       wd + "/tls/client-key.pem",
		TrustedCAFile: wd + "/tls/ca.pem",
	}
	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		log.Println(err)
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
		TLS:         tlsConfig,
	})
	if err != nil {
		// etcd clientv3 >= v3.2.10, grpc/grpc-go >= v1.7.3
		if err == context.DeadlineExceeded {
			// handle errors
		}

		// etcd clientv3 <= v3.2.9, grpc/grpc-go <= v1.2.1
		if err == grpc.ErrClientConnTimeout {
			// handle errors
		}
		log.Println(err)
	}
	defer cli.Close()

	app := rider.New()

	app.GET("/:key", func(c rider.Context) {
		key := c.Param("key")
		c.SendString(200, getValue(cli, key))
	})

	app.Graceful(":5001")

}

// 查询etcdkey
func getValue(cli *clientv3.Client, key string) string {
	kv := clientv3.NewKV(cli)
	resp, err := kv.Get(context.TODO(), key)
	if err != nil {
		return ""
	}
	kvs := resp.Kvs
	if len(kvs) == 0 {
		return ""
	}
	return string(kvs[0].Value)
}
