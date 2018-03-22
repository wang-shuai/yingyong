package index

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/olebedev/config"
	"../model"
)

func (this *Index) BuildUserIdx(filelist map[string]int64) {
	doc,dt := getTemplateIdx(model.UserCode)

	buildUserFileRef(dt,model.UserCode,filelist)
	buildUserBcpDataStructure(dt)

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	dir := model.Basepath + model.UserDir
	xmlpath := dir + model.XmlIndexName
	doc.WriteToFile(xmlpath)
}

func buildUserBcpDataStructure(dt *etree.Element){
	ds := dt.CreateElement("DATASET")
	ds.CreateAttr("name",model.UserCode)
	ds.CreateAttr("rmk","BCP文件数据结构")
	data := ds.CreateElement("DATA")

	conf, err := config.ParseJsonFile(model.Basepath+ "/conf/user_idx.json")
	if err != nil {
		fmt.Println("获取用户描述json异常：", err)
		return
	}

	usermap, err := conf.List("user")
	if err != nil {
		fmt.Println("获取用户描述字段错误：", err)
		return
	}

	var item *etree.Element
	for i:=0;i<len(usermap);i++{
		item = data.CreateElement("ITEM")
		name,_:=conf.String(fmt.Sprintf("user.%d.name",i))
		desc,_:=conf.String(fmt.Sprintf("user.%d.desc",i))
		item.CreateAttr("key",name)
		item.CreateAttr("eng",name)
		item.CreateAttr("rmk",desc)
	}
}
