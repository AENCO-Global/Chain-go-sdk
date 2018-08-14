package transaction

// Struct containing transaction type map.
var TransactionType  = map[string]int{
	// Transfer Transaction transaction type.
	"TRANSFER": 0x4154,
	// Register namespace transaction type.
	"REGISTER_NAMESPACE": 0x414e,
	// Mosaic definition transaction type.
	"MOSAIC_DEFINITION": 0x414d,
	// Mosaic supply change transaction.
	"MOSAIC_SUPPLY_CHANGE": 0x424d,
	// Modify multisig account transaction type.
	"MODIFY_MULTISIG_ACCOUNT": 0x4155,
	// Aggregate complete transaction type.
	"AGGREGATE_COMPLETE": 0x4141,
	// Aggregate bonded transaction type.
	"AGGREGATE_BONDED": 0x4241,
	// Lock transaction type.
	"LOCK": 0x414C,
	// Secret Lock Transaction type.
	"SECRET_LOCK": 0x424C,
	// Secret Proof transaction type.
	"SECRET_PROOF": 0x434C,
}

