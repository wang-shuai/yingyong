package index

import (
	"golang-services/jingyong/model"
)

func (this *Index) BuildEvaluateIdx(filelist map[string]int64) {
	doc, dt := getTemplateIdx(model.EvaluateCode)

	buildFileRef(dt, model.EvaluateCode, filelist)

	jsonpath := model.Basepath + "/conf/evaluate_idx.json"
	buildBcpDataStructure(dt, model.EvaluateCode, jsonpath)

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	dir := model.Basepath + model.EvaluateDir
	xmlpath := dir + model.XmlIndexName
	doc.WriteToFile(xmlpath)
}
