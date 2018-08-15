package main

import (
	"fmt"
	"github.com/slackve/nem2-sdk-go/sdk/infrastructure"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
)

const mosaic  = "d525ad41d95fcf29"
const mosaic2  = "d525ad41d95fcf28"

func main() {

	// Blockchain routes.
	const Server = "http://catapult.isarq.com:3000"

	client := infrastructure.NewAPIClient()
	client.ChangeServer(Server)

	// 01 - Get mosaic information.
	// Param mosaic - Mosaic identifier.
	a, err := client.MosaicRoutesApi.GetMosaic(mosaic)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetMosaic:\n%s", utils.Struc2Json(a))

	// 02 - Get information for a set of mosaics.
	// Param mosaics - Slice of mosaicIds..
	var mosaics infrastructure.MosaicIds
	// Append attachment into Mosaics
	mosaics.MosaicIds = append(mosaics.MosaicIds, mosaic)
	// Append attachment into Mosaics
	mosaics.MosaicIds = append(mosaics.MosaicIds, mosaic2)
	b, err := client.MosaicRoutesApi.GetMosaics(mosaics)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetMosaics:\n%s", utils.Struc2Json(b))

	// 03 - Get an array of MosaicInfo from mosaics created under provided namespace.
	// Param mosaic - Namespace identifier.
	// Param 0 - The number of mosaics to return.
	// Param "" - Identifier of the mosaic after which we want the transactions to be returned.
	c, err := client.MosaicRoutesApi.GetMosaicsFromNamespace("84b3552d375ffa4b", 0, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetMosaicsFromNamespace:\n%s", utils.Struc2Json(c))

	// 04 - Get readable names for a set of mosaics.
	// Param mosaic - Slice of mosaicIds.
	var msc infrastructure.MosaicIds
	//// Append attachment into Mosaic
	msc.MosaicIds = append(msc.MosaicIds, mosaic)
	d, err := client.MosaicRoutesApi.GetMosaicsName(msc)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetMosaicsName:\n%s", utils.Struc2Json(d))
}
