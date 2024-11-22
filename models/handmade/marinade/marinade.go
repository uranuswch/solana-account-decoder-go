package marinade

import (
	ag_solanago "github.com/gagliardetto/solana-go"
)

type List struct {
	Account     ag_solanago.PublicKey
	ItemSize    uint32
	Count       uint32
	NewAccount  ag_solanago.PublicKey
	CopiedCount uint32
}

type StakeSystem struct {
	StakeList                 List
	DelayedUnstakeCoolingDown uint64
	StakeDepositBumpSeed      uint8
	StakeWithdrawBumpSeed     uint8
	SlotsForStakeDelta        uint64
	LastStakeDeltaEpoch       uint64
	MinStake                  uint64
	ExtraStakeDeltaRuns       uint32
}

type ValidatorSystem struct {
	ValidatorList           List
	ManagerAuthority        ag_solanago.PublicKey
	TotalValidatorScore     uint32
	TotalActiveBalance      uint64
	AutoAddValidatorEnabled uint8
}

type LiqPool struct {
	LpMint                   ag_solanago.PublicKey
	LpMintAuthorityBumpSeed  uint8
	SolLegBumpSeed           uint8
	MsolLegAuthorityBumpSeed uint8
	MsolLeg                  ag_solanago.PublicKey
	LpLiquidityTarget        uint64
	LpMaxFee                 Fee
	LpMinFee                 Fee
	TreasuryCut              Fee
	LpSupply                 uint64
	LentFromSolLeg           uint64
	LiqSolCap                uint64
}

type Fee struct {
	BasisPoints uint32
}

type FeeCents struct {
	BpCents uint32
}

type State struct {
	Reserve                   [8]byte
	MsolMint                  ag_solanago.PublicKey
	AdminAuthority            ag_solanago.PublicKey
	OperationalSolAccount     ag_solanago.PublicKey
	TreasuryMsolAccount       ag_solanago.PublicKey
	ReserveBumpSeed           uint8
	MsolMintAuthorityBumpSeed uint8
	RentExemptForTokenAcc     uint64
	RewardFee                 Fee
	StakeSystem               StakeSystem
	ValidatorSystem           ValidatorSystem
	LiqPool                   LiqPool
	AvailableReserveBalance   uint64
	MsolSupply                uint64
	MsolPrice                 uint64

	CirculatingTicketCount uint64

	CirculatingTicketBalance uint64
	LentFromReserve          uint64
	MinDeposit               uint64
	MinWithdraw              uint64
	StakingSolCap            uint64
	EmergencyCoolingDown     uint64

	PauseAuthority              ag_solanago.PublicKey
	Paused                      bool
	DelayedUnstakeFee           FeeCents
	WithdrawStakeAccountFee     FeeCents
	WithdrawStakeAccountEnabled bool
	LastStakeMoveEpoch          uint64
	StakeMoved                  uint64
	MaxStakeMovedPerEpoch       Fee
}
