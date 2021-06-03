package chia

import "github.com/Pow-Duck/GetChiaKey/internal/models"

func GetKey() (*models.ChiaKey, error) {
	return &models.ChiaKey{
		Fingerprint:        "Fingerprint",
		MasterPublicKey:    "MasterPublicKey",
		FarmerPublicKey:    "FarmerPublicKey",
		PoolPublicKey:      "PoolPublicKey",
		FirstWalletAddress: "FirstWalletAddress",
	}, nil
}
