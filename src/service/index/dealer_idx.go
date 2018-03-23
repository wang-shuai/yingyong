package index

import (
	"../model"
)

func (this *Index) BuildDealerIdx(filelist map[string]int64) {
	doc, dt := getTemplateIdx(model.DealerCode)

	buildFileRef(dt, model.DealerCode, filelist)

	jsonpath := model.Basepath + "/conf/dealer_idx.json"
	buildBcpDataStructure(dt, model.DealerCode, jsonpath)

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	dir := model.Basepath + model.DealerDir
	xmlpath := dir + model.XmlIndexName
	doc.WriteToFile(xmlpath)
}
