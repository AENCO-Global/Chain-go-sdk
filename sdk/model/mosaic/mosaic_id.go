package mosaic

import (
	"github.com/slackve/nem2-sdk-go/core/coders"
)

//var Mosaicid = struct {
//	Create func(interface{}) MosaicId
//}{
//	Create: CreateMosaicId,
//}
// The mosaic id structure describes mosaic id
type MosaicId struct {
	// Mosaic id
	Id coders.UInt64 //Id
	// Mosaic full name
	FullName string `json:"fullName,omitempty"`
}

//func CreateMosaicId(id interface{}) MosaicId {
//	if s, ok := id.(string); ok {
//		limiterPosition := strings.Split(s, ":")
//		namespaceName := limiterPosition[0]
//		mosaicName := limiterPosition[1]
//
//		return MosaicId{
//			Id:       coders.UInt64{},
//			FullName: s,
//		}
//	}
//	return MosaicId{}
//}
