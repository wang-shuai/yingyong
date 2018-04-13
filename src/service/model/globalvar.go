package model

import (
	"os"
	"github.com/olebedev/config"
	"golang-services/jingyong/tool"
	"github.com/garyburd/redigo/redis"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
)

var (
	Basepath string
	Cfg      *config.Config
	Pool     *redis.Pool
)

func init() {
	Basepath, _ = os.Getwd()
	Cfg, _ = config.ParseJsonFile(Basepath + "/conf/app.json")

	address, _ := Cfg.String("redis.address")
	idle, _ := Cfg.Int("redis.maxidle")
	timeout, _ := Cfg.Int("redis.idletimeout")
	Pool = tool.RedisInit(address, idle, timeout)

	fname, _ := Cfg.String("log.filename")
	fpath, _ := Cfg.String("log.path")
	flevel, _ := Cfg.Int("log.loglevel")
	flog.Init(fname, fpath, flevel)

	if _, e := os.Open(fpath); e != nil {
		os.MkdirAll(fpath, os.ModePerm)
	}
}
