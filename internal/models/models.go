package models

type ChiaKey struct {
	Fingerprint        string `json:"fingerprint"`
	MasterPublicKey    string `json:"master_public_key"`
	FarmerPublicKey    string `json:"farmer_public_key"`
	PoolPublicKey      string `json:"pool_public_key"`
	FirstWalletAddress string `json:"first_wallet_address"`
}
