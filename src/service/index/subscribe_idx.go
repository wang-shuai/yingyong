package index

import (
	"golang-services/jingyong/model"
)

func (this *Index) BuildSubscribeIdx(filelist map[string]int64) {
	doc, dt := getTemplateIdx(model.SubscribeCode)

	buildFileRef(dt, model.SubscribeCode, filelist)

	jsonpath := model.Basepath + "/conf/subscribe_idx.json"
	buildBcpDataStructure(dt, model.SubscribeCode, jsonpath)

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	dir := model.Basepath + model.SubscribeDir
	xmlpath := dir + model.XmlIndexName
	doc.WriteToFile(xmlpath)
}
