package main

import (
	"./bcp"
)

func main() {
	bcp.ZipUserInfo()
	bcp.ZipDealerUserInfo()
	bcp.ZipCollectionInfo()
	bcp.ZipSubscribeInfo()
	bcp.ZipUcar()
	bcp.ZipEvaluate()
	bcp.ZipLoanOrder()

	//bcp.ZipBooks() // todo 后续的数据组合操作未完成 redis没有数据 最后再处理

	bcp.ZipDealers()
}
