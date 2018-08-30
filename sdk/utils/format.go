package utils

import "strings"

func AddressFormat(adddress string) string  {
	ad := strings.TrimSpace(adddress)
	ad = strings.Replace(ad, "-", "", -1)
	ad = strings.ToUpper(ad)
	return ad
}
