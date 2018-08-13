package blockchain

// Static var containing network type constants.
var NetworkType = struct {
	// Main net network
	MAIN_NET int
	// Test net network
	TEST_NET int
	// Mijin net network
	MIJIN int
	// Mijin test net network
	MIJIN_TEST int
}{
	MAIN_NET:   104,
	TEST_NET:   152,
	MIJIN:      96,
	MIJIN_TEST: 144,
}
