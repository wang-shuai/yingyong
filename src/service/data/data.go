package data

import (
	"../model"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var (
	newcar_engine *xorm.Engine //新车数据库操作
	ucar_engine   *xorm.Engine // ucar数据库
	uleads_engine *xorm.Engine // ucarleads数据库
)

func init() {
	defer func() {
		fmt.Println("recover")
		recover()
	}()

	schema_n, _ := model.Cfg.String("newcardb.schema")
	connStr_n, _ := model.Cfg.String("newcardb.connectionString")
	schema_u, _ := model.Cfg.String("ucardb.schema")
	connStr_u, _ := model.Cfg.String("ucardb.connectionString")
	schema_l, _ := model.Cfg.String("ucarleadsdb.schema")
	connStr_l, _ := model.Cfg.String("ucarleadsdb.connectionString")

	if Eg, err := xorm.NewEngine(schema_n, connStr_n); err != nil {
		fmt.Println(err)
		panic("newcar数据库链接失败")
	} else {
		newcar_engine = Eg
	}
	newcar_engine.SetMapper(core.SameMapper{}) //与字段、表名一致  不区分大小写
	newcar_engine.ShowSQL(true)                //展示每次执行的sql

	if Eg, err := xorm.NewEngine(schema_u, connStr_u); err != nil {
		fmt.Println(err)
		panic("ucar数据库链接失败")
	} else {
		ucar_engine = Eg
	}
	ucar_engine.SetMapper(core.SameMapper{}) //与字段、表名一致  不区分大小写
	ucar_engine.ShowSQL(true)                //展示每次执行的sql

	if Eg, err := xorm.NewEngine(schema_l, connStr_l); err != nil {
		fmt.Println(err)
		panic("ucarleads数据库链接失败")
	} else {
		uleads_engine = Eg
	}
	uleads_engine.SetMapper(core.SameMapper{}) //与字段、表名一致  不区分大小写
	uleads_engine.ShowSQL(true)                //展示每次执行的sql
}

