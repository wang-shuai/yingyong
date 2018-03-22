package model

import (
	"os"
	"github.com/olebedev/config"
)

var (
	Basepath string
	Cfg *config.Config
)

func init(){
	Basepath,_ = os.Getwd()
	Cfg,_ = config.ParseJsonFile(Basepath + "/conf/app.json")
}