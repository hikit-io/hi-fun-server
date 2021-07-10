package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	_RedisCluster *redis.ClusterClient
)

func RedisCluster() *redis.ClusterClient {
	if _RedisCluster == nil {
		cli := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{
				"119.3.191.228:7000",
				"119.3.191.228:7001",
				"119.3.191.228:7002",
				"119.3.191.228:7003",
				"119.3.191.228:7004",
				"119.3.191.228:7005",
			},
			Password: "1234",
		})
		cmd := cli.Ping(context.Background())
		if cmd.Err() != nil {
			log.Panicln(cmd.Err())
		}
		log.Println(cmd.Result())
		_RedisCluster = cli
	}
	return _RedisCluster
}
