package infrastructure

type Multisig struct {

	Account string `json:"account"`

	AccountAddress string `json:"accountAddress,omitempty"`

	MinApproval int32 `json:"minApproval"`

	MinRemoval int32 `json:"minRemoval"`

	Cosignatories []string `json:"cosignatories"`

	MultisigAccounts []string `json:"multisigAccounts"`
}
