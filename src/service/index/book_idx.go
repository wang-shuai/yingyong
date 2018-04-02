package index

import (
	"../model"
)

func (this *Index) BuildBookIdx(filelist map[string]int64) {
	doc, dt := getTemplateIdx(model.BookCode)

	buildFileRef(dt, model.BookCode, filelist)

	jsonpath := model.Basepath + "/conf/book_idx.json"
	buildBcpDataStructure(dt, model.BookCode, jsonpath)

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	dir := model.Basepath + model.BookDir
	xmlpath := dir + model.XmlIndexName
	doc.WriteToFile(xmlpath)
}
