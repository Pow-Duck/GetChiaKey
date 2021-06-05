package chia

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	var r = `Showing all public keys derived from your private keys:

Fingerprint: 1904308493
Master public key (m): 91aa9e00e4ab2a0cb591100304393230f3e916eff9b987f4eb88cb70f70d9008df981e11b8e8d38492081b5e6c411c76
Farmer public key (m/12381/8444/0/0): a21841c2ab0b9bc5489448aac27404d1abfe2f4c8ca79e76e370369081be610d22971f483ebc7aef7cf932ed1164500d
Pool public key (m/12381/8444/1/0): b081e71b6345728a08d5e11ab422386729b5ad72f54c1f431a79b50840e781c66a8e4ba86691799dc112917780e3b91b
First wallet address: xch1t6th6g34hgdhe4g7rck4df2uhl3kywhunyewxvne0d7pggxt22lsqlm0t5`

	split := strings.Split(r, ":")
	for k, v := range split {
		index := strings.Index(v, "\n")
		if index != -1 {
			fmt.Printf("key: %d val: %s\n", k, v[:index])
		}
	}
}
