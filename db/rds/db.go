package rds

import (
	"os"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

func New() *redis.Client {
	rdshost := os.Getenv("RDSHOST")
	rdsport := os.Getenv("RDSPORT")
	rdsname := os.Getenv("RDSNAME")

	dbname, e := strconv.Atoi(rdsname)
	if e != nil {
		panic(e)
	}

	connection := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rdshost, rdsport),
		Password: "", // no password set
		DB:       dbname,
	})

	_, err := connection.Ping().Result()
	if err != nil {
		panic(err)
	}

	//defer connection.Close()

	return connection
}
