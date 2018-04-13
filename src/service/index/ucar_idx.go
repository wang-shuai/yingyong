package index

import (
	"golang-services/jingyong/model"
)

func (this *Index) BuildUcarIdx(filelist map[string]int64) {
	doc, dt := getTemplateIdx(model.UcarCode)

	buildFileRef(dt, model.UcarCode, filelist)

	jsonpath := model.Basepath + "/conf/ucar_idx.json"
	buildBcpDataStructure(dt, model.UcarCode, jsonpath)

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	dir := model.Basepath + model.UcarDir
	xmlpath := dir + model.XmlIndexName
	doc.WriteToFile(xmlpath)
}
