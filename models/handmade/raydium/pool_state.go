package raydium

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type RewardInfo struct {
	RewardState           uint8
	OpenTime              uint64
	EndTime               uint64
	LastUpdateTime        uint64
	EmissionsPerSecondX64 ag_binary.Uint128
	RewardTotalEmissioned uint64
	RewardClaimed         uint64
	TokenMint             ag_solanago.PublicKey
	TokenVault            ag_solanago.PublicKey
	Authority             ag_solanago.PublicKey
	RewardGrowthGlobalX64 ag_binary.Uint128
}

type PoolState struct {
	Reserve             [8]byte
	Bump                [1]uint8
	AmmConfig           ag_solanago.PublicKey
	Owner               ag_solanago.PublicKey
	TokenMint0          ag_solanago.PublicKey
	TokenMint1          ag_solanago.PublicKey
	TokenVault0         ag_solanago.PublicKey
	TokenVault1         ag_solanago.PublicKey
	ObservationKey      ag_solanago.PublicKey
	MintDecimals0       uint8
	MintDecimals1       uint8
	Tickspacing         uint16
	Liquidity           ag_binary.Uint128
	PricePriceX64       ag_binary.Uint128
	TicksCurrent        int32
	Padding3            uint16
	Padding4            uint16
	FeeGrowthGlobal0X64 ag_binary.Uint128
	FeeGrowthGlobal1X64 ag_binary.Uint128
	ProtocolFeesToken0  uint64
	ProtocolFeesToken1  uint64
	SwapInAmountToken0  ag_binary.Uint128
	SwapOutAmountToken1 ag_binary.Uint128
	SwapInAmountToken1  ag_binary.Uint128
	SwapOutAmountToken0 ag_binary.Uint128
	Status              uint8
	Padding             [7]uint8
	RewardInfos         [3]RewardInfo
	TickArrayMap        [16]uint64
	TotalFeesToken0     uint64
	TotalFeesToken1     uint64
	FundFeesToken0      uint64
	FundFeesToken1      uint64
	OpenTime            uint64
	RecentEpoch         uint64
	Padding1            [24]uint64
	Padding2            [32]uint64
}
