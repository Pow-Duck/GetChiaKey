// +build windows

package chia

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/Pow-Duck/GetChiaKey/internal/models"
	"github.com/Pow-Duck/processes"
)

func GetKey() (*models.ChiaKey, error) {
	getenv := os.Getenv("USERPROFILE")
	basePath := path.Join(getenv, "AppData", "Local", "chia-blockchain")
	dir, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	var versions []string
	for _, v := range dir {
		if v.IsDir() {
			if strings.Contains(v.Name(), "app-") {
				versions = append(versions, v.Name())
			}
		}
	}

	if len(versions) == 0 {
		return nil, errors.New("Chia is not installed  Chia未安装")
	}
	vs := versions[len(versions)-1]
	fmt.Println(versions)
	chiaPath := path.Join(basePath, vs, "resources", "app.asar.unpacked", "daemon", "chia.exe")
	command, err := processes.RunCommand(fmt.Sprintf("%s keys show", chiaPath))
	if err != nil {
		return nil, err
	}

	split := strings.Split(command, ":")

	fmt.Println(command)
	return &models.ChiaKey{
		Fingerprint:        spc(split[2]),
		MasterPublicKey:    spc(split[3]),
		FarmerPublicKey:    spc(split[4]),
		PoolPublicKey:      spc(split[5]),
		FirstWalletAddress: spc(split[6]),
	}, nil
}

func spc(r string) string {
	index := strings.Index(r, "\n")
	if index != -1 {
		return strings.TrimSpace(r[:index])
	}

	return strings.TrimSpace(r)
}
