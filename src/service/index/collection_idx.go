package index

import (
	"golang-services/jingyong/model"
)

func (this *Index) BuildCollectionIdx(filelist map[string]int64) {
	doc, dt := getTemplateIdx(model.CollectionCode)

	buildFileRef(dt, model.CollectionCode, filelist)

	jsonpath := model.Basepath + "/conf/collection_idx.json"
	buildBcpDataStructure(dt, model.CollectionCode, jsonpath)

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	dir := model.Basepath + model.CollectionDir
	xmlpath := dir + model.XmlIndexName
	doc.WriteToFile(xmlpath)
}
