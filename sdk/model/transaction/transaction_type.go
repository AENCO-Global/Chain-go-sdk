package transaction

// Struct containing transaction type map.
var TransactionType = struct {
	// Transfer Transaction transaction type.
	TRANSFER int
	// Register namespace transaction type.
	REGISTER_NAMESPACE int
	// Mosaic supply change transaction.
	MOSAIC_DEFINITION int
	// Modify multisig account transaction type.
	MOSAIC_SUPPLY_CHANGE int
	// Modify multisig account transaction type.
	MODIFY_MULTISIG_ACCOUNT int
	// Aggregate complete transaction type.
	AGGREGATE_COMPLETE int
	// Aggregate bonded transaction type.
	AGGREGATE_BONDED int
	// Lock transaction type.
	LOCK int
	// Secret Lock Transaction type.
	SECRET_LOCK int
	// Secret Proof transaction type.
	SECRET_PROOF int
}{
	TRANSFER:                0x4154,
	REGISTER_NAMESPACE:      0x414e,
	MOSAIC_DEFINITION:       0x414d,
	MOSAIC_SUPPLY_CHANGE:    0x424d,
	MODIFY_MULTISIG_ACCOUNT: 0x4155,
	AGGREGATE_COMPLETE:      0x4141,
	AGGREGATE_BONDED:        0x4241,
	LOCK:                    0x414C,
	SECRET_LOCK:             0x424C,
	SECRET_PROOF:            0x434C,
}
