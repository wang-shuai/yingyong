package index

import (
	"../model"
)

func (this *Index) BuildLoanOrderIdx(filelist map[string]int64) {
	doc, dt := getTemplateIdx(model.LoanOrderCode)

	buildFileRef(dt, model.LoanOrderCode, filelist)

	jsonpath := model.Basepath + "/conf/loanorder_idx.json"
	buildBcpDataStructure(dt, model.LoanOrderCode, jsonpath)

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	dir := model.Basepath + model.LoanOrderDir
	xmlpath := dir + model.XmlIndexName
	doc.WriteToFile(xmlpath)
}
