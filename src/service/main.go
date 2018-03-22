package main

import (
	//"./bcp"
	"fmt"
	"./data"
)

func main(){
	//op := new(bcp.BcpOperation)
	//op.ZipUserInfo()

	i ,_ :=	data.CountUCarDealer()
	ulist,_:= data.GetUCarDealers(20,60)

	fmt.Println(i,"\n",ulist)
}