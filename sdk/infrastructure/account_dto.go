package infrastructure

type AccountDto struct {
	Address string `json:"address"`

	AddressHeight *UInt64Dto `json:"addressHeight"`

	PublicKey string `json:"publicKey"`

	PublicKeyHeight *UInt64Dto `json:"publicKeyHeight"`

	Mosaics []MosaicDto `json:"mosaics"`

	Importance *UInt64Dto `json:"importance"`

	ImportanceHeight *UInt64Dto `json:"importanceHeight"`
}
