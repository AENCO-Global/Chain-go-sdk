package coders

import (
	"fmt"
	"strings"
	"encoding/binary"
)

var namespace_base_id = []uint64{uint64(0), uint64(0)}

const (
	namespace_max_depth = 3
	name_pattern = "/^[a-z0-9][a-z0-9-_]*$/"
)

func generateId(parentId int, name string)  {
	parentIdBytes := make([]byte, 8)
	binary.LittleEndian.PutUint32(parentIdBytes, uint32(parentId))
	//hash := sha3.New256()
}

// Generates a mosaic id given a unified mosaic name.
// param name The unified mosaic name.
// returns The mosaic id.
func GenerateMosaicId(name string)  {
	if 0 >= len(name) {
		fmt.Errorf("having zero length %v", name)
	}
	//mosaicSeparatorIndex := findMosaicSeparatorIndex(name)
	//namespacePath := mosaicSeparatorIndex[0]

}
func findMosaicSeparatorIndex(name string) []string {

	var mosaicSeparatorIndex = strings.Split(name,":")
	if 0 > len(mosaicSeparatorIndex) {
		fmt.Errorf("missing mosaic %v", name)
	}

	if 0 == len(mosaicSeparatorIndex) {
		fmt.Errorf("empty part %v", name)

	}
	return mosaicSeparatorIndex
}

// Parses a unified namespace name into a path.
// param {string} name The unified namespace name.
// returns The namespace path.
//func generateNamespacePath(name string) []uint64 {
//	if 0 >= len(name) {
//		fmt.Errorf("having zero length")
//	}
//	namespaceId := namespace_base_id
//	path := make(map[string]string)
//	start :=
//}