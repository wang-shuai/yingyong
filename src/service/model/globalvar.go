package model

import (
	"os"
	"github.com/olebedev/config"
	"../tool"
	"github.com/garyburd/redigo/redis"
)

var (
	Basepath string
	Cfg *config.Config
	Pool *redis.Pool
)

func init(){
	Basepath,_ = os.Getwd()
	Cfg,_ = config.ParseJsonFile(Basepath + "/conf/app.json")

	address ,_:= Cfg.String("redis.address")
	idle ,_:= Cfg.Int("redis.maxidle")
	timeout ,_:= Cfg.Int("redis.idletimeout")
	Pool = tool.RedisInit(address,idle,timeout)
}