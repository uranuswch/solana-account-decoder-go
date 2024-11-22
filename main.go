package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"solana-decode/models/handmade/marinade"
	"solana-decode/models/handmade/raydium"
	"solana-decode/models/handmade/stakepool"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type ProgramEnum int

const (
	StakePool ProgramEnum = iota
	Marinade
	RaydiumCPMM
)

var accounts = map[ProgramEnum][]string{
	StakePool: {
		"Jito4APyf642JPZPx3hGc6WWJ8zPKtRbRs4P815Awbb",
		"stk9ApL5HeVAwPLr3TLhDXdZS8ptVu7zp6ov8HFDuMi",
		"9mhGNSPArRMHpLDMSmxAvuoizBqtBGqYdT8WGuqgxNdn",
		"3wK2g8ZdzAH8FJ7PKr2RcvGh7V9VYson5hrVsJM5Lmws",
		"DqhH94PjkZsjAqEze2BEkWhFQJ6EyU6MdtMphMgnXqeK",
		"Hr9pzexrBge3vgmBNRR8u42CNQgBXdHm4UkUN2DH4a7r",
		"2aMLkB5p5gVvCwKkdSo5eZAL1WwhZbxezQr1wxiynRhq",
		"po1osKDWYF9oiVEGmzKA4eTs8eMveFRMox3bUKazGN2",
	},
	Marinade: {
		"8szGkuLTAux9XMgZ2vtY39jVSowEcpBfFfD8hXSEqdGC",
	},
	RaydiumCPMM: {
		"3nMFwZXwY1s1M5s8vYAHqd4wGs4iSxXE4LRoUMMYqEgF",
		"21uKEoCGxzx69hBVMGG9WVBzCgHabc1VSqeA2q2KpegL",
	},
}

func call[T any](client *rpc.Client, pools []string) {
	var pubKeys []solana.PublicKey
	for _, pool := range pools {
		pubKeys = append(pubKeys, solana.MustPublicKeyFromBase58(pool))
	}
	var resp *rpc.GetMultipleAccountsResult
	resp, err := client.GetMultipleAccounts(context.Background(), pubKeys...)
	if err != nil {
		panic(err)
	}
	for idx, account := range resp.Value {
		if account == nil || account.Data == nil {
			fmt.Printf("%s, data is nil\n", pools[idx])
		} else {
			borshDecoder := bin.NewBorshDecoder([]byte(account.Data.GetBinary()))
			var result T
			err := borshDecoder.Decode(&result)
			if err != nil {
				fmt.Println(err)
			} else {
				debugStr, err := json.Marshal(result)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(pools[idx])
					fmt.Println(string(debugStr))
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	flag.Parse()
	client := rpc.New(flag.Arg(0))
	for program, pools := range accounts {
		switch program {
		case StakePool:
			call[stakepool.Pool](client, pools)
		case Marinade:
			call[marinade.State](client, pools)
		case RaydiumCPMM:
			call[raydium.PoolState](client, pools)
		}
	}
}
