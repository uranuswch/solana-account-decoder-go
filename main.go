package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	ag_solanago "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/near/borsh-go"
)

type AccountType borsh.Enum

type Fee struct {
	Denominator uint64
	Numerator   uint64
}

type Lockup struct {
	UnixTimestamp uint64
	Epoch         uint64
	Custodian     ag_solanago.PublicKey
}

type StakePool struct {
	AccountType                           ag_binary.BorshEnum
	Manager                               ag_solanago.PublicKey
	Staker                                ag_solanago.PublicKey
	StakeDepositAuthority                 ag_solanago.PublicKey
	StakeWithdrawBumpSeed                 uint8
	ValidatorList                         ag_solanago.PublicKey
	ReserveStake                          ag_solanago.PublicKey
	PoolMint                              ag_solanago.PublicKey
	ManagerFeeAccount                     ag_solanago.PublicKey
	TokenProgramId                        ag_solanago.PublicKey
	TotalLamports                         uint64
	PoolTokenSupply                       uint64
	LastUpdateEpoch                       uint64
	Lockup                                Lockup
	EpochFee                              Fee
	NextEpochFee                          *Fee
	PreferredDepositValidatorVoteAddress  *ag_solanago.PublicKey
	PreferredWithdrawValidatorVoteAddress *ag_solanago.PublicKey
	StakeDepositFee                       Fee
	StakeWithdrawalFee                    Fee
	NextStakeWithdrawalFee                *Fee
	StakeReferralFee                      uint8
	SolDepositAuthority                   *ag_solanago.PublicKey
	SolDepositFee                         Fee
	SolReferralFee                        uint8
	SolWithdrawAuthority                  *ag_solanago.PublicKey
	SolWithdrawalFee                      Fee
	NextSolWithdrawalFee                  *Fee
	LastEpochPoolTokenSupply              uint64
	LastEpochTotalLamports                uint64
}

type MarinadeList struct {
	Account     ag_solanago.PublicKey
	ItemSize    uint32
	Count       uint32
	NewAccount  ag_solanago.PublicKey
	CopiedCount uint32
}

type MarinadeStakeSystem struct {
	StakeList                 MarinadeList
	DelayedUnstakeCoolingDown uint64
	StakeDepositBumpSeed      uint8
	StakeWithdrawBumpSeed     uint8
	SlotsForStakeDelta        uint64
	LastStakeDeltaEpoch       uint64
	MinStake                  uint64
	ExtraStakeDeltaRuns       uint32
}

type MarinadeValidatorSystem struct {
	ValidatorList           MarinadeList
	ManagerAuthority        ag_solanago.PublicKey
	TotalValidatorScore     uint32
	TotalActiveBalance      uint64
	AutoAddValidatorEnabled uint8
}

type MarinadeLiqPool struct {
	LpMint                   ag_solanago.PublicKey
	LpMintAuthorityBumpSeed  uint8
	SolLegBumpSeed           uint8
	MsolLegAuthorityBumpSeed uint8
	MsolLeg                  ag_solanago.PublicKey
	LpLiquidityTarget        uint64
	LpMaxFee                 MarinadeFee
	LpMinFee                 MarinadeFee
	TreasuryCut              MarinadeFee
	LpSupply                 uint64
	LentFromSolLeg           uint64
	LiqSolCap                uint64
}

type MarinadeFee struct {
	BasisPoints uint32
}

type MarinadeFeeCents struct {
	BpCents uint32
}

type MarinadeState struct {
	MsolMint                  ag_solanago.PublicKey
	AdminAuthority            ag_solanago.PublicKey
	OperationalSolAccount     ag_solanago.PublicKey
	TreasuryMsolAccount       ag_solanago.PublicKey
	ReserveBumpSeed           uint8
	MsolMintAuthorityBumpSeed uint8
	RentExemptForTokenAcc     uint64
	RewardFee                 MarinadeFee
	StakeSystem               MarinadeStakeSystem
	ValidatorSystem           MarinadeValidatorSystem
	LiqPool                   MarinadeLiqPool
	AvailableReserveBalance   uint64
	MsolSupply                uint64
	MsolPrice                 uint64
	CirculatingTicketCount    uint64
	CirculatingTicketBalance  uint64
	LentFromReserve           uint64
	MinDeposit                uint64
	MinWithdraw               uint64
	StakingSolCap             uint64
	EmergencyCoolingDown      uint64
	// PauseAuthority              ag_solanago.PublicKey
	// Paused                      bool
	// DelayedUnstakeFee           MarinadeFeeCents
	// WithdrawStakeAccountFee     MarinadeFeeCents
	// WithdrawStakeAccountEnabled bool
	// LastStakeMoveEpoch          uint64
	// StakeMoved                  uint64
	// MaxStakeMovedPerEpoch       MarinadeFee
}

var stakePools = []string{
	"Jito4APyf642JPZPx3hGc6WWJ8zPKtRbRs4P815Awbb",
	"stk9ApL5HeVAwPLr3TLhDXdZS8ptVu7zp6ov8HFDuMi",
	"9mhGNSPArRMHpLDMSmxAvuoizBqtBGqYdT8WGuqgxNdn",
	"3wK2g8ZdzAH8FJ7PKr2RcvGh7V9VYson5hrVsJM5Lmws",
	"DqhH94PjkZsjAqEze2BEkWhFQJ6EyU6MdtMphMgnXqeK",
	"Hr9pzexrBge3vgmBNRR8u42CNQgBXdHm4UkUN2DH4a7r",
	"2aMLkB5p5gVvCwKkdSo5eZAL1WwhZbxezQr1wxiynRhq",
	"po1osKDWYF9oiVEGmzKA4eTs8eMveFRMox3bUKazGN2",
}

var marinadePools = []string{
	"8szGkuLTAux9XMgZ2vtY39jVSowEcpBfFfD8hXSEqdGC",
}

var unknownPools = []string{
	"EMjuABxELpYWYEwjkKmQKBNCwdaFAy4QYAs6W9bDQDNw",
	"GdNXJobf8fbTR5JSE7adxa6niaygjx4EEbnnRaDCHMMW",
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
	call[StakePool](client, stakePools)
	call[MarinadeState](client, marinadePools)
	call[StakePool](client, unknownPools)
}
