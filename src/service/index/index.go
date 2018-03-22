package index

import (
	"strconv"
	"github.com/beevik/etree"
	"../model"
)

type Index struct {
}

// 生成idx模板，返回文档指针及要插入数据的节点指针
func getTemplateIdx(code string) (*etree.Document, *etree.Element) {

	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	// doc.CreateProcInst("xml-stylesheet", `type="text/xsl" href="style.xsl"`)

	msg := doc.CreateElement("MESSAGE")
	ds := msg.CreateElement("DATASET")
	ds.CreateAttr("name", "WA_COMMON_010000")
	ds.CreateAttr("rmk", "数据交互通用信息")

	dt := ds.CreateElement("DATA")
	item := dt.CreateElement("ITEM")
	item.CreateAttr("key", "CLUE_SRC_SYS")
	item.CreateAttr("val", "110100")
	item.CreateAttr("eng", "FROM")
	item.CreateAttr("chn", "发起节点的标识")
	item.CreateAttr("rmk", "北京市海淀区")

	item = dt.CreateElement("ITEM")
	item.CreateAttr("key", "CLUE_DST_SYS")
	item.CreateAttr("val", "110100")
	item.CreateAttr("eng", "TO")
	item.CreateAttr("chn", "目的节点的标识")
	item.CreateAttr("rmk", "北京市局")

	item = dt.CreateElement("ITEM")
	item.CreateAttr("key", "BUSINESS_SERVER_TYPE")
	item.CreateAttr("val", "03")
	item.CreateAttr("eng", "BUSINESS_SERVER_TYPE")
	item.CreateAttr("chn", "业务服务类型")
	item.CreateAttr("rmk", "主动上报类")

	ds = msg.CreateElement("DATASET")
	ds.CreateAttr("name", "WA_COMMON_010017")
	ds.CreateAttr("ver", "1.0")
	ds.CreateAttr("rmk", "数据文件索引信息")

	dt = ds.CreateElement("DATA")
	ds = dt.CreateElement("DATASET")
	ds.CreateAttr("name", "WA_COMMON_010013")
	ds.CreateAttr("rmk", "BCP文件描述信息")

	dt = ds.CreateElement("DATA")
	item = dt.CreateElement("ITEM")
	item.CreateAttr("key", "COLUMN_SPLIT")
	item.CreateAttr("val", "")
	item.CreateAttr("rmk", "列分隔符（缺少值时默认为制表符\\t）")

	item = dt.CreateElement("ITEM")
	item.CreateAttr("key", "ROW_SPLIT")
	item.CreateAttr("val", "")
	item.CreateAttr("rmk", "行分隔符（缺少值时默认为换行符\\n）")

	item = dt.CreateElement("ITEM")
	item.CreateAttr("key", "DATA_SOURCE")
	item.CreateAttr("val", "151")
	item.CreateAttr("rmk", "数据来源")

	item = dt.CreateElement("ITEM")
	item.CreateAttr("key", "SECURITY_SOFTWARE_ORGCODE")
	item.CreateAttr("val", strconv.Itoa(model.AppType))
	item.CreateAttr("rmk", "厂家组织机构代码")

	item = dt.CreateElement("ITEM")
	item.CreateAttr("key", "CALLBACK_SEQUENCE")
	item.CreateAttr("val", "1")
	item.CreateAttr("rmk", "数据起始行，可选项，不填写默认为第1行")

	item = dt.CreateElement("ITEM")
	item.CreateAttr("key", "CALLBACK_RECORDS")
	item.CreateAttr("val", "UTF-8")
	item.CreateAttr("rmk", "可选项，默认为UTF-8，BCP文件编码格式（采用不带格式的编码方式，如：UTF-8无BOM）")

	item = dt.CreateElement("ITEM")
	item.CreateAttr("key", "DATA_SET")
	item.CreateAttr("val", code)
	item.CreateAttr("rmk", "数据集代码")

	return doc, dt
}

func buildUserFileRef(dt *etree.Element, code string, filelist map[string]int64) {
	ds := dt.CreateElement("DATASET")
	ds.CreateAttr("name", code)
	ds.CreateAttr("rmk", "BCP数据文件信息")

	// 循环 添加文件目录

	data := ds.CreateElement("DATA")
	item := data.CreateElement("ITEM")
	item.CreateAttr("key", "TRANSFILE")
	item.CreateAttr("val", code)
	item.CreateAttr("rmk", "文件路径")

	for filename, cnt := range filelist {
		item = data.CreateElement("ITEM")
		item.CreateAttr("key", "FILE_NAME")
		item.CreateAttr("val", filename)
		item.CreateAttr("rmk", "文件名")

		item = data.CreateElement("ITEM")
		item.CreateAttr("key", "TEXT_FORMAT")
		item.CreateAttr("val", strconv.FormatInt(cnt, 10))
		item.CreateAttr("rmk", "记录行数")
	}

}
