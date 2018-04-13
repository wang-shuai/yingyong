package index

import (
	"golang-services/jingyong/model"
)

func (this *Index) BuildDealerUserIdx(filelist map[string]int64) {
	doc, dt := getTemplateIdx(model.DealerUserCode)

	buildFileRef(dt, model.DealerUserCode, filelist)

	jsonpath := model.Basepath + "/conf/dealer_user_idx.json"
	buildBcpDataStructure(dt, model.DealerUserCode, jsonpath)

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	dir := model.Basepath + model.DealerUserDir
	xmlpath := dir + model.XmlIndexName
	doc.WriteToFile(xmlpath)
}
