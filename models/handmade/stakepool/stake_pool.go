package stakepool

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Fee struct {
	Denominator uint64
	Numerator   uint64
}

type Lockup struct {
	UnixTimestamp uint64
	Epoch         uint64
	Custodian     ag_solanago.PublicKey
}

type Pool struct {
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
