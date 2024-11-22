// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type SplitStakeAccountInfo struct {
	Account ag_solanago.PublicKey
	Index   uint32
}

func (obj SplitStakeAccountInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Account` param:
	err = encoder.Encode(obj.Account)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SplitStakeAccountInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Account`:
	err = decoder.Decode(&obj.Account)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	return nil
}

type U64ValueChange struct {
	Old uint64
	New uint64
}

func (obj U64ValueChange) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Old` param:
	err = encoder.Encode(obj.Old)
	if err != nil {
		return err
	}
	// Serialize `New` param:
	err = encoder.Encode(obj.New)
	if err != nil {
		return err
	}
	return nil
}

func (obj *U64ValueChange) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Old`:
	err = decoder.Decode(&obj.Old)
	if err != nil {
		return err
	}
	// Deserialize `New`:
	err = decoder.Decode(&obj.New)
	if err != nil {
		return err
	}
	return nil
}

type U32ValueChange struct {
	Old uint32
	New uint32
}

func (obj U32ValueChange) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Old` param:
	err = encoder.Encode(obj.Old)
	if err != nil {
		return err
	}
	// Serialize `New` param:
	err = encoder.Encode(obj.New)
	if err != nil {
		return err
	}
	return nil
}

func (obj *U32ValueChange) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Old`:
	err = decoder.Decode(&obj.Old)
	if err != nil {
		return err
	}
	// Deserialize `New`:
	err = decoder.Decode(&obj.New)
	if err != nil {
		return err
	}
	return nil
}

type FeeValueChange struct {
	Old Fee
	New Fee
}

func (obj FeeValueChange) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Old` param:
	err = encoder.Encode(obj.Old)
	if err != nil {
		return err
	}
	// Serialize `New` param:
	err = encoder.Encode(obj.New)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FeeValueChange) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Old`:
	err = decoder.Decode(&obj.Old)
	if err != nil {
		return err
	}
	// Deserialize `New`:
	err = decoder.Decode(&obj.New)
	if err != nil {
		return err
	}
	return nil
}

type FeeCentsValueChange struct {
	Old FeeCents
	New FeeCents
}

func (obj FeeCentsValueChange) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Old` param:
	err = encoder.Encode(obj.Old)
	if err != nil {
		return err
	}
	// Serialize `New` param:
	err = encoder.Encode(obj.New)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FeeCentsValueChange) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Old`:
	err = decoder.Decode(&obj.Old)
	if err != nil {
		return err
	}
	// Deserialize `New`:
	err = decoder.Decode(&obj.New)
	if err != nil {
		return err
	}
	return nil
}

type PubkeyValueChange struct {
	Old ag_solanago.PublicKey
	New ag_solanago.PublicKey
}

func (obj PubkeyValueChange) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Old` param:
	err = encoder.Encode(obj.Old)
	if err != nil {
		return err
	}
	// Serialize `New` param:
	err = encoder.Encode(obj.New)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PubkeyValueChange) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Old`:
	err = decoder.Decode(&obj.Old)
	if err != nil {
		return err
	}
	// Deserialize `New`:
	err = decoder.Decode(&obj.New)
	if err != nil {
		return err
	}
	return nil
}

type BoolValueChange struct {
	Old bool
	New bool
}

func (obj BoolValueChange) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Old` param:
	err = encoder.Encode(obj.Old)
	if err != nil {
		return err
	}
	// Serialize `New` param:
	err = encoder.Encode(obj.New)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BoolValueChange) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Old`:
	err = decoder.Decode(&obj.Old)
	if err != nil {
		return err
	}
	// Deserialize `New`:
	err = decoder.Decode(&obj.New)
	if err != nil {
		return err
	}
	return nil
}

type ChangeAuthorityData struct {
	Admin                 *ag_solanago.PublicKey `bin:"optional"`
	ValidatorManager      *ag_solanago.PublicKey `bin:"optional"`
	OperationalSolAccount *ag_solanago.PublicKey `bin:"optional"`
	TreasuryMsolAccount   *ag_solanago.PublicKey `bin:"optional"`
	PauseAuthority        *ag_solanago.PublicKey `bin:"optional"`
}

func (obj ChangeAuthorityData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Admin` param (optional):
	{
		if obj.Admin == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Admin)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `ValidatorManager` param (optional):
	{
		if obj.ValidatorManager == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.ValidatorManager)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `OperationalSolAccount` param (optional):
	{
		if obj.OperationalSolAccount == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.OperationalSolAccount)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `TreasuryMsolAccount` param (optional):
	{
		if obj.TreasuryMsolAccount == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TreasuryMsolAccount)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `PauseAuthority` param (optional):
	{
		if obj.PauseAuthority == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.PauseAuthority)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *ChangeAuthorityData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Admin` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Admin)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `ValidatorManager` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.ValidatorManager)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `OperationalSolAccount` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.OperationalSolAccount)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `TreasuryMsolAccount` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TreasuryMsolAccount)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `PauseAuthority` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.PauseAuthority)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type ConfigLpParams struct {
	MinFee          *Fee    `bin:"optional"`
	MaxFee          *Fee    `bin:"optional"`
	LiquidityTarget *uint64 `bin:"optional"`
	TreasuryCut     *Fee    `bin:"optional"`
}

func (obj ConfigLpParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MinFee` param (optional):
	{
		if obj.MinFee == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MinFee)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `MaxFee` param (optional):
	{
		if obj.MaxFee == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MaxFee)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `LiquidityTarget` param (optional):
	{
		if obj.LiquidityTarget == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.LiquidityTarget)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `TreasuryCut` param (optional):
	{
		if obj.TreasuryCut == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TreasuryCut)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *ConfigLpParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MinFee` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MinFee)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `MaxFee` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MaxFee)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `LiquidityTarget` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.LiquidityTarget)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `TreasuryCut` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TreasuryCut)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type ConfigMarinadeParams struct {
	RewardsFee                  *Fee      `bin:"optional"`
	SlotsForStakeDelta          *uint64   `bin:"optional"`
	MinStake                    *uint64   `bin:"optional"`
	MinDeposit                  *uint64   `bin:"optional"`
	MinWithdraw                 *uint64   `bin:"optional"`
	StakingSolCap               *uint64   `bin:"optional"`
	LiquiditySolCap             *uint64   `bin:"optional"`
	WithdrawStakeAccountEnabled *bool     `bin:"optional"`
	DelayedUnstakeFee           *FeeCents `bin:"optional"`
	WithdrawStakeAccountFee     *FeeCents `bin:"optional"`
	MaxStakeMovedPerEpoch       *Fee      `bin:"optional"`
}

func (obj ConfigMarinadeParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RewardsFee` param (optional):
	{
		if obj.RewardsFee == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.RewardsFee)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `SlotsForStakeDelta` param (optional):
	{
		if obj.SlotsForStakeDelta == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SlotsForStakeDelta)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `MinStake` param (optional):
	{
		if obj.MinStake == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MinStake)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `MinDeposit` param (optional):
	{
		if obj.MinDeposit == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MinDeposit)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `MinWithdraw` param (optional):
	{
		if obj.MinWithdraw == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MinWithdraw)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `StakingSolCap` param (optional):
	{
		if obj.StakingSolCap == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.StakingSolCap)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `LiquiditySolCap` param (optional):
	{
		if obj.LiquiditySolCap == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.LiquiditySolCap)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `WithdrawStakeAccountEnabled` param (optional):
	{
		if obj.WithdrawStakeAccountEnabled == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.WithdrawStakeAccountEnabled)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `DelayedUnstakeFee` param (optional):
	{
		if obj.DelayedUnstakeFee == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.DelayedUnstakeFee)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `WithdrawStakeAccountFee` param (optional):
	{
		if obj.WithdrawStakeAccountFee == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.WithdrawStakeAccountFee)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `MaxStakeMovedPerEpoch` param (optional):
	{
		if obj.MaxStakeMovedPerEpoch == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MaxStakeMovedPerEpoch)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *ConfigMarinadeParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RewardsFee` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.RewardsFee)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `SlotsForStakeDelta` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SlotsForStakeDelta)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `MinStake` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MinStake)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `MinDeposit` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MinDeposit)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `MinWithdraw` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MinWithdraw)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `StakingSolCap` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.StakingSolCap)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `LiquiditySolCap` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.LiquiditySolCap)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `WithdrawStakeAccountEnabled` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.WithdrawStakeAccountEnabled)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `DelayedUnstakeFee` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.DelayedUnstakeFee)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `WithdrawStakeAccountFee` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.WithdrawStakeAccountFee)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `MaxStakeMovedPerEpoch` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MaxStakeMovedPerEpoch)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type InitializeData struct {
	AdminAuthority                 ag_solanago.PublicKey
	ValidatorManagerAuthority      ag_solanago.PublicKey
	MinStake                       uint64
	RewardsFee                     Fee
	LiqPool                        LiqPoolInitializeData
	AdditionalStakeRecordSpace     uint32
	AdditionalValidatorRecordSpace uint32
	SlotsForStakeDelta             uint64
	PauseAuthority                 ag_solanago.PublicKey
}

func (obj InitializeData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AdminAuthority` param:
	err = encoder.Encode(obj.AdminAuthority)
	if err != nil {
		return err
	}
	// Serialize `ValidatorManagerAuthority` param:
	err = encoder.Encode(obj.ValidatorManagerAuthority)
	if err != nil {
		return err
	}
	// Serialize `MinStake` param:
	err = encoder.Encode(obj.MinStake)
	if err != nil {
		return err
	}
	// Serialize `RewardsFee` param:
	err = encoder.Encode(obj.RewardsFee)
	if err != nil {
		return err
	}
	// Serialize `LiqPool` param:
	err = encoder.Encode(obj.LiqPool)
	if err != nil {
		return err
	}
	// Serialize `AdditionalStakeRecordSpace` param:
	err = encoder.Encode(obj.AdditionalStakeRecordSpace)
	if err != nil {
		return err
	}
	// Serialize `AdditionalValidatorRecordSpace` param:
	err = encoder.Encode(obj.AdditionalValidatorRecordSpace)
	if err != nil {
		return err
	}
	// Serialize `SlotsForStakeDelta` param:
	err = encoder.Encode(obj.SlotsForStakeDelta)
	if err != nil {
		return err
	}
	// Serialize `PauseAuthority` param:
	err = encoder.Encode(obj.PauseAuthority)
	if err != nil {
		return err
	}
	return nil
}

func (obj *InitializeData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AdminAuthority`:
	err = decoder.Decode(&obj.AdminAuthority)
	if err != nil {
		return err
	}
	// Deserialize `ValidatorManagerAuthority`:
	err = decoder.Decode(&obj.ValidatorManagerAuthority)
	if err != nil {
		return err
	}
	// Deserialize `MinStake`:
	err = decoder.Decode(&obj.MinStake)
	if err != nil {
		return err
	}
	// Deserialize `RewardsFee`:
	err = decoder.Decode(&obj.RewardsFee)
	if err != nil {
		return err
	}
	// Deserialize `LiqPool`:
	err = decoder.Decode(&obj.LiqPool)
	if err != nil {
		return err
	}
	// Deserialize `AdditionalStakeRecordSpace`:
	err = decoder.Decode(&obj.AdditionalStakeRecordSpace)
	if err != nil {
		return err
	}
	// Deserialize `AdditionalValidatorRecordSpace`:
	err = decoder.Decode(&obj.AdditionalValidatorRecordSpace)
	if err != nil {
		return err
	}
	// Deserialize `SlotsForStakeDelta`:
	err = decoder.Decode(&obj.SlotsForStakeDelta)
	if err != nil {
		return err
	}
	// Deserialize `PauseAuthority`:
	err = decoder.Decode(&obj.PauseAuthority)
	if err != nil {
		return err
	}
	return nil
}

type LiqPoolInitializeData struct {
	LpLiquidityTarget uint64
	LpMaxFee          Fee
	LpMinFee          Fee
	LpTreasuryCut     Fee
}

func (obj LiqPoolInitializeData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LpLiquidityTarget` param:
	err = encoder.Encode(obj.LpLiquidityTarget)
	if err != nil {
		return err
	}
	// Serialize `LpMaxFee` param:
	err = encoder.Encode(obj.LpMaxFee)
	if err != nil {
		return err
	}
	// Serialize `LpMinFee` param:
	err = encoder.Encode(obj.LpMinFee)
	if err != nil {
		return err
	}
	// Serialize `LpTreasuryCut` param:
	err = encoder.Encode(obj.LpTreasuryCut)
	if err != nil {
		return err
	}
	return nil
}

func (obj *LiqPoolInitializeData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LpLiquidityTarget`:
	err = decoder.Decode(&obj.LpLiquidityTarget)
	if err != nil {
		return err
	}
	// Deserialize `LpMaxFee`:
	err = decoder.Decode(&obj.LpMaxFee)
	if err != nil {
		return err
	}
	// Deserialize `LpMinFee`:
	err = decoder.Decode(&obj.LpMinFee)
	if err != nil {
		return err
	}
	// Deserialize `LpTreasuryCut`:
	err = decoder.Decode(&obj.LpTreasuryCut)
	if err != nil {
		return err
	}
	return nil
}

type Fee struct {
	BasisPoints uint32
}

func (obj Fee) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `BasisPoints` param:
	err = encoder.Encode(obj.BasisPoints)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Fee) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `BasisPoints`:
	err = decoder.Decode(&obj.BasisPoints)
	if err != nil {
		return err
	}
	return nil
}

type FeeCents struct {
	BpCents uint32
}

func (obj FeeCents) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `BpCents` param:
	err = encoder.Encode(obj.BpCents)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FeeCents) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `BpCents`:
	err = decoder.Decode(&obj.BpCents)
	if err != nil {
		return err
	}
	return nil
}

type LiqPool struct {
	LpMint                   ag_solanago.PublicKey
	LpMintAuthorityBumpSeed  uint8
	SolLegBumpSeed           uint8
	MsolLegAuthorityBumpSeed uint8
	MsolLeg                  ag_solanago.PublicKey

	// Liquidity target. If the Liquidity reach this amount, the fee reaches lp_min_discount_fee
	LpLiquidityTarget uint64

	// Liquidity pool max fee
	LpMaxFee Fee

	// SOL/mSOL Liquidity pool min fee
	LpMinFee Fee

	// Treasury cut
	TreasuryCut     Fee
	LpSupply        uint64
	LentFromSolLeg  uint64
	LiquiditySolCap uint64
}

func (obj LiqPool) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LpMint` param:
	err = encoder.Encode(obj.LpMint)
	if err != nil {
		return err
	}
	// Serialize `LpMintAuthorityBumpSeed` param:
	err = encoder.Encode(obj.LpMintAuthorityBumpSeed)
	if err != nil {
		return err
	}
	// Serialize `SolLegBumpSeed` param:
	err = encoder.Encode(obj.SolLegBumpSeed)
	if err != nil {
		return err
	}
	// Serialize `MsolLegAuthorityBumpSeed` param:
	err = encoder.Encode(obj.MsolLegAuthorityBumpSeed)
	if err != nil {
		return err
	}
	// Serialize `MsolLeg` param:
	err = encoder.Encode(obj.MsolLeg)
	if err != nil {
		return err
	}
	// Serialize `LpLiquidityTarget` param:
	err = encoder.Encode(obj.LpLiquidityTarget)
	if err != nil {
		return err
	}
	// Serialize `LpMaxFee` param:
	err = encoder.Encode(obj.LpMaxFee)
	if err != nil {
		return err
	}
	// Serialize `LpMinFee` param:
	err = encoder.Encode(obj.LpMinFee)
	if err != nil {
		return err
	}
	// Serialize `TreasuryCut` param:
	err = encoder.Encode(obj.TreasuryCut)
	if err != nil {
		return err
	}
	// Serialize `LpSupply` param:
	err = encoder.Encode(obj.LpSupply)
	if err != nil {
		return err
	}
	// Serialize `LentFromSolLeg` param:
	err = encoder.Encode(obj.LentFromSolLeg)
	if err != nil {
		return err
	}
	// Serialize `LiquiditySolCap` param:
	err = encoder.Encode(obj.LiquiditySolCap)
	if err != nil {
		return err
	}
	return nil
}

func (obj *LiqPool) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LpMint`:
	err = decoder.Decode(&obj.LpMint)
	if err != nil {
		return err
	}
	// Deserialize `LpMintAuthorityBumpSeed`:
	err = decoder.Decode(&obj.LpMintAuthorityBumpSeed)
	if err != nil {
		return err
	}
	// Deserialize `SolLegBumpSeed`:
	err = decoder.Decode(&obj.SolLegBumpSeed)
	if err != nil {
		return err
	}
	// Deserialize `MsolLegAuthorityBumpSeed`:
	err = decoder.Decode(&obj.MsolLegAuthorityBumpSeed)
	if err != nil {
		return err
	}
	// Deserialize `MsolLeg`:
	err = decoder.Decode(&obj.MsolLeg)
	if err != nil {
		return err
	}
	// Deserialize `LpLiquidityTarget`:
	err = decoder.Decode(&obj.LpLiquidityTarget)
	if err != nil {
		return err
	}
	// Deserialize `LpMaxFee`:
	err = decoder.Decode(&obj.LpMaxFee)
	if err != nil {
		return err
	}
	// Deserialize `LpMinFee`:
	err = decoder.Decode(&obj.LpMinFee)
	if err != nil {
		return err
	}
	// Deserialize `TreasuryCut`:
	err = decoder.Decode(&obj.TreasuryCut)
	if err != nil {
		return err
	}
	// Deserialize `LpSupply`:
	err = decoder.Decode(&obj.LpSupply)
	if err != nil {
		return err
	}
	// Deserialize `LentFromSolLeg`:
	err = decoder.Decode(&obj.LentFromSolLeg)
	if err != nil {
		return err
	}
	// Deserialize `LiquiditySolCap`:
	err = decoder.Decode(&obj.LiquiditySolCap)
	if err != nil {
		return err
	}
	return nil
}

type List struct {
	Account   ag_solanago.PublicKey
	ItemSize  uint32
	Count     uint32
	Reserved1 ag_solanago.PublicKey
	Reserved2 uint32
}

func (obj List) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Account` param:
	err = encoder.Encode(obj.Account)
	if err != nil {
		return err
	}
	// Serialize `ItemSize` param:
	err = encoder.Encode(obj.ItemSize)
	if err != nil {
		return err
	}
	// Serialize `Count` param:
	err = encoder.Encode(obj.Count)
	if err != nil {
		return err
	}
	// Serialize `Reserved1` param:
	err = encoder.Encode(obj.Reserved1)
	if err != nil {
		return err
	}
	// Serialize `Reserved2` param:
	err = encoder.Encode(obj.Reserved2)
	if err != nil {
		return err
	}
	return nil
}

func (obj *List) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Account`:
	err = decoder.Decode(&obj.Account)
	if err != nil {
		return err
	}
	// Deserialize `ItemSize`:
	err = decoder.Decode(&obj.ItemSize)
	if err != nil {
		return err
	}
	// Deserialize `Count`:
	err = decoder.Decode(&obj.Count)
	if err != nil {
		return err
	}
	// Deserialize `Reserved1`:
	err = decoder.Decode(&obj.Reserved1)
	if err != nil {
		return err
	}
	// Deserialize `Reserved2`:
	err = decoder.Decode(&obj.Reserved2)
	if err != nil {
		return err
	}
	return nil
}

type StakeRecord struct {
	StakeAccount                ag_solanago.PublicKey
	LastUpdateDelegatedLamports uint64
	LastUpdateEpoch             uint64
	IsEmergencyUnstaking        uint8
}

func (obj StakeRecord) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `StakeAccount` param:
	err = encoder.Encode(obj.StakeAccount)
	if err != nil {
		return err
	}
	// Serialize `LastUpdateDelegatedLamports` param:
	err = encoder.Encode(obj.LastUpdateDelegatedLamports)
	if err != nil {
		return err
	}
	// Serialize `LastUpdateEpoch` param:
	err = encoder.Encode(obj.LastUpdateEpoch)
	if err != nil {
		return err
	}
	// Serialize `IsEmergencyUnstaking` param:
	err = encoder.Encode(obj.IsEmergencyUnstaking)
	if err != nil {
		return err
	}
	return nil
}

func (obj *StakeRecord) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `StakeAccount`:
	err = decoder.Decode(&obj.StakeAccount)
	if err != nil {
		return err
	}
	// Deserialize `LastUpdateDelegatedLamports`:
	err = decoder.Decode(&obj.LastUpdateDelegatedLamports)
	if err != nil {
		return err
	}
	// Deserialize `LastUpdateEpoch`:
	err = decoder.Decode(&obj.LastUpdateEpoch)
	if err != nil {
		return err
	}
	// Deserialize `IsEmergencyUnstaking`:
	err = decoder.Decode(&obj.IsEmergencyUnstaking)
	if err != nil {
		return err
	}
	return nil
}

type StakeList struct{}

func (obj StakeList) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (obj *StakeList) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

type StakeSystem struct {
	StakeList                 List
	DelayedUnstakeCoolingDown uint64
	StakeDepositBumpSeed      uint8
	StakeWithdrawBumpSeed     uint8

	// set by admin, how much slots before the end of the epoch, stake-delta can start
	SlotsForStakeDelta uint64

	// Marks the start of stake-delta operations, meaning that if somebody starts a delayed-unstake ticket
	// after this var is set with epoch_num the ticket will have epoch_created = current_epoch+1
	// (the user must wait one more epoch, because their unstake-delta will be execute in this epoch)
	LastStakeDeltaEpoch uint64
	MinStake            uint64

	// can be set by validator-manager-auth to allow a second run of stake-delta to stake late stakers in the last minute of the epoch
	// so we maximize user's rewards
	ExtraStakeDeltaRuns uint32
}

func (obj StakeSystem) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `StakeList` param:
	err = encoder.Encode(obj.StakeList)
	if err != nil {
		return err
	}
	// Serialize `DelayedUnstakeCoolingDown` param:
	err = encoder.Encode(obj.DelayedUnstakeCoolingDown)
	if err != nil {
		return err
	}
	// Serialize `StakeDepositBumpSeed` param:
	err = encoder.Encode(obj.StakeDepositBumpSeed)
	if err != nil {
		return err
	}
	// Serialize `StakeWithdrawBumpSeed` param:
	err = encoder.Encode(obj.StakeWithdrawBumpSeed)
	if err != nil {
		return err
	}
	// Serialize `SlotsForStakeDelta` param:
	err = encoder.Encode(obj.SlotsForStakeDelta)
	if err != nil {
		return err
	}
	// Serialize `LastStakeDeltaEpoch` param:
	err = encoder.Encode(obj.LastStakeDeltaEpoch)
	if err != nil {
		return err
	}
	// Serialize `MinStake` param:
	err = encoder.Encode(obj.MinStake)
	if err != nil {
		return err
	}
	// Serialize `ExtraStakeDeltaRuns` param:
	err = encoder.Encode(obj.ExtraStakeDeltaRuns)
	if err != nil {
		return err
	}
	return nil
}

func (obj *StakeSystem) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `StakeList`:
	err = decoder.Decode(&obj.StakeList)
	if err != nil {
		return err
	}
	// Deserialize `DelayedUnstakeCoolingDown`:
	err = decoder.Decode(&obj.DelayedUnstakeCoolingDown)
	if err != nil {
		return err
	}
	// Deserialize `StakeDepositBumpSeed`:
	err = decoder.Decode(&obj.StakeDepositBumpSeed)
	if err != nil {
		return err
	}
	// Deserialize `StakeWithdrawBumpSeed`:
	err = decoder.Decode(&obj.StakeWithdrawBumpSeed)
	if err != nil {
		return err
	}
	// Deserialize `SlotsForStakeDelta`:
	err = decoder.Decode(&obj.SlotsForStakeDelta)
	if err != nil {
		return err
	}
	// Deserialize `LastStakeDeltaEpoch`:
	err = decoder.Decode(&obj.LastStakeDeltaEpoch)
	if err != nil {
		return err
	}
	// Deserialize `MinStake`:
	err = decoder.Decode(&obj.MinStake)
	if err != nil {
		return err
	}
	// Deserialize `ExtraStakeDeltaRuns`:
	err = decoder.Decode(&obj.ExtraStakeDeltaRuns)
	if err != nil {
		return err
	}
	return nil
}

type ValidatorRecord struct {
	// Validator vote pubkey
	ValidatorAccount ag_solanago.PublicKey

	// Validator total balance in lamports
	ActiveBalance           uint64
	Score                   uint32
	LastStakeDeltaEpoch     uint64
	DuplicationFlagBumpSeed uint8
}

func (obj ValidatorRecord) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ValidatorAccount` param:
	err = encoder.Encode(obj.ValidatorAccount)
	if err != nil {
		return err
	}
	// Serialize `ActiveBalance` param:
	err = encoder.Encode(obj.ActiveBalance)
	if err != nil {
		return err
	}
	// Serialize `Score` param:
	err = encoder.Encode(obj.Score)
	if err != nil {
		return err
	}
	// Serialize `LastStakeDeltaEpoch` param:
	err = encoder.Encode(obj.LastStakeDeltaEpoch)
	if err != nil {
		return err
	}
	// Serialize `DuplicationFlagBumpSeed` param:
	err = encoder.Encode(obj.DuplicationFlagBumpSeed)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ValidatorRecord) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ValidatorAccount`:
	err = decoder.Decode(&obj.ValidatorAccount)
	if err != nil {
		return err
	}
	// Deserialize `ActiveBalance`:
	err = decoder.Decode(&obj.ActiveBalance)
	if err != nil {
		return err
	}
	// Deserialize `Score`:
	err = decoder.Decode(&obj.Score)
	if err != nil {
		return err
	}
	// Deserialize `LastStakeDeltaEpoch`:
	err = decoder.Decode(&obj.LastStakeDeltaEpoch)
	if err != nil {
		return err
	}
	// Deserialize `DuplicationFlagBumpSeed`:
	err = decoder.Decode(&obj.DuplicationFlagBumpSeed)
	if err != nil {
		return err
	}
	return nil
}

type ValidatorList struct{}

func (obj ValidatorList) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (obj *ValidatorList) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

type ValidatorSystem struct {
	ValidatorList       List
	ManagerAuthority    ag_solanago.PublicKey
	TotalValidatorScore uint32

	// sum of all active lamports staked
	TotalActiveBalance uint64

	// DEPRECATED, no longer used
	AutoAddValidatorEnabled uint8
}

func (obj ValidatorSystem) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ValidatorList` param:
	err = encoder.Encode(obj.ValidatorList)
	if err != nil {
		return err
	}
	// Serialize `ManagerAuthority` param:
	err = encoder.Encode(obj.ManagerAuthority)
	if err != nil {
		return err
	}
	// Serialize `TotalValidatorScore` param:
	err = encoder.Encode(obj.TotalValidatorScore)
	if err != nil {
		return err
	}
	// Serialize `TotalActiveBalance` param:
	err = encoder.Encode(obj.TotalActiveBalance)
	if err != nil {
		return err
	}
	// Serialize `AutoAddValidatorEnabled` param:
	err = encoder.Encode(obj.AutoAddValidatorEnabled)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ValidatorSystem) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ValidatorList`:
	err = decoder.Decode(&obj.ValidatorList)
	if err != nil {
		return err
	}
	// Deserialize `ManagerAuthority`:
	err = decoder.Decode(&obj.ManagerAuthority)
	if err != nil {
		return err
	}
	// Deserialize `TotalValidatorScore`:
	err = decoder.Decode(&obj.TotalValidatorScore)
	if err != nil {
		return err
	}
	// Deserialize `TotalActiveBalance`:
	err = decoder.Decode(&obj.TotalActiveBalance)
	if err != nil {
		return err
	}
	// Deserialize `AutoAddValidatorEnabled`:
	err = decoder.Decode(&obj.AutoAddValidatorEnabled)
	if err != nil {
		return err
	}
	return nil
}