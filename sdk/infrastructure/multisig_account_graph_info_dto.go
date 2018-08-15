package infrastructure

type MultisigAccountGraphInfoDto struct {

	Level int32 `json:"level"`

	MultisigEntries []MultisigAccountInfoDto `json:"multisigEntries"`
}
