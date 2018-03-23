package index

import (
	"../model"
)

func (this *Index) BuildUserIdx(filelist map[string]int64) {
	doc, dt := getTemplateIdx(model.UserCode)

	buildFileRef(dt, model.UserCode, filelist)

	jsonpath := model.Basepath + "/conf/user_idx.json"
	buildBcpDataStructure(dt, model.UserCode, jsonpath)

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	dir := model.Basepath + model.UserDir
	xmlpath := dir + model.XmlIndexName
	doc.WriteToFile(xmlpath)
}
