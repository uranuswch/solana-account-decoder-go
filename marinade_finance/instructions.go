// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "MarinadeFinance"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	Instruction_ChangeAuthority = ag_binary.TypeID([8]byte{50, 106, 66, 104, 99, 118, 145, 88})

	Instruction_AddValidator = ag_binary.TypeID([8]byte{250, 113, 53, 54, 141, 117, 215, 185})

	Instruction_RemoveValidator = ag_binary.TypeID([8]byte{25, 96, 211, 155, 161, 14, 168, 188})

	Instruction_SetValidatorScore = ag_binary.TypeID([8]byte{101, 41, 206, 33, 216, 111, 25, 78})

	Instruction_ConfigValidatorSystem = ag_binary.TypeID([8]byte{27, 90, 97, 209, 17, 115, 7, 40})

	Instruction_Deposit = ag_binary.TypeID([8]byte{242, 35, 198, 137, 82, 225, 242, 182})

	Instruction_DepositStakeAccount = ag_binary.TypeID([8]byte{110, 130, 115, 41, 164, 102, 2, 59})

	Instruction_LiquidUnstake = ag_binary.TypeID([8]byte{30, 30, 119, 240, 191, 227, 12, 16})

	Instruction_AddLiquidity = ag_binary.TypeID([8]byte{181, 157, 89, 67, 143, 182, 52, 72})

	Instruction_RemoveLiquidity = ag_binary.TypeID([8]byte{80, 85, 209, 72, 24, 206, 177, 108})

	Instruction_ConfigLp = ag_binary.TypeID([8]byte{10, 24, 168, 119, 86, 48, 225, 17})

	Instruction_ConfigMarinade = ag_binary.TypeID([8]byte{67, 3, 34, 114, 190, 185, 17, 62})

	Instruction_OrderUnstake = ag_binary.TypeID([8]byte{97, 167, 144, 107, 117, 190, 128, 36})

	Instruction_Claim = ag_binary.TypeID([8]byte{62, 198, 214, 193, 213, 159, 108, 210})

	Instruction_StakeReserve = ag_binary.TypeID([8]byte{87, 217, 23, 179, 205, 25, 113, 129})

	Instruction_UpdateActive = ag_binary.TypeID([8]byte{4, 67, 81, 64, 136, 245, 93, 152})

	Instruction_UpdateDeactivated = ag_binary.TypeID([8]byte{16, 232, 131, 115, 156, 100, 239, 50})

	Instruction_DeactivateStake = ag_binary.TypeID([8]byte{165, 158, 229, 97, 168, 220, 187, 225})

	Instruction_EmergencyUnstake = ag_binary.TypeID([8]byte{123, 69, 168, 195, 183, 213, 199, 214})

	Instruction_PartialUnstake = ag_binary.TypeID([8]byte{55, 241, 205, 221, 45, 114, 205, 163})

	Instruction_MergeStakes = ag_binary.TypeID([8]byte{216, 36, 141, 225, 243, 78, 125, 237})

	Instruction_Redelegate = ag_binary.TypeID([8]byte{212, 82, 51, 160, 228, 80, 116, 35})

	Instruction_Pause = ag_binary.TypeID([8]byte{211, 22, 221, 251, 74, 121, 193, 47})

	Instruction_Resume = ag_binary.TypeID([8]byte{1, 166, 51, 170, 127, 32, 141, 206})

	Instruction_WithdrawStakeAccount = ag_binary.TypeID([8]byte{211, 85, 184, 65, 183, 177, 233, 217})

	Instruction_ReallocValidatorList = ag_binary.TypeID([8]byte{215, 59, 218, 133, 93, 138, 60, 123})

	Instruction_ReallocStakeList = ag_binary.TypeID([8]byte{12, 36, 124, 27, 128, 96, 85, 199})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_ChangeAuthority:
		return "ChangeAuthority"
	case Instruction_AddValidator:
		return "AddValidator"
	case Instruction_RemoveValidator:
		return "RemoveValidator"
	case Instruction_SetValidatorScore:
		return "SetValidatorScore"
	case Instruction_ConfigValidatorSystem:
		return "ConfigValidatorSystem"
	case Instruction_Deposit:
		return "Deposit"
	case Instruction_DepositStakeAccount:
		return "DepositStakeAccount"
	case Instruction_LiquidUnstake:
		return "LiquidUnstake"
	case Instruction_AddLiquidity:
		return "AddLiquidity"
	case Instruction_RemoveLiquidity:
		return "RemoveLiquidity"
	case Instruction_ConfigLp:
		return "ConfigLp"
	case Instruction_ConfigMarinade:
		return "ConfigMarinade"
	case Instruction_OrderUnstake:
		return "OrderUnstake"
	case Instruction_Claim:
		return "Claim"
	case Instruction_StakeReserve:
		return "StakeReserve"
	case Instruction_UpdateActive:
		return "UpdateActive"
	case Instruction_UpdateDeactivated:
		return "UpdateDeactivated"
	case Instruction_DeactivateStake:
		return "DeactivateStake"
	case Instruction_EmergencyUnstake:
		return "EmergencyUnstake"
	case Instruction_PartialUnstake:
		return "PartialUnstake"
	case Instruction_MergeStakes:
		return "MergeStakes"
	case Instruction_Redelegate:
		return "Redelegate"
	case Instruction_Pause:
		return "Pause"
	case Instruction_Resume:
		return "Resume"
	case Instruction_WithdrawStakeAccount:
		return "WithdrawStakeAccount"
	case Instruction_ReallocValidatorList:
		return "ReallocValidatorList"
	case Instruction_ReallocStakeList:
		return "ReallocStakeList"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"initialize", (*Initialize)(nil),
		},
		{
			"change_authority", (*ChangeAuthority)(nil),
		},
		{
			"add_validator", (*AddValidator)(nil),
		},
		{
			"remove_validator", (*RemoveValidator)(nil),
		},
		{
			"set_validator_score", (*SetValidatorScore)(nil),
		},
		{
			"config_validator_system", (*ConfigValidatorSystem)(nil),
		},
		{
			"deposit", (*Deposit)(nil),
		},
		{
			"deposit_stake_account", (*DepositStakeAccount)(nil),
		},
		{
			"liquid_unstake", (*LiquidUnstake)(nil),
		},
		{
			"add_liquidity", (*AddLiquidity)(nil),
		},
		{
			"remove_liquidity", (*RemoveLiquidity)(nil),
		},
		{
			"config_lp", (*ConfigLp)(nil),
		},
		{
			"config_marinade", (*ConfigMarinade)(nil),
		},
		{
			"order_unstake", (*OrderUnstake)(nil),
		},
		{
			"claim", (*Claim)(nil),
		},
		{
			"stake_reserve", (*StakeReserve)(nil),
		},
		{
			"update_active", (*UpdateActive)(nil),
		},
		{
			"update_deactivated", (*UpdateDeactivated)(nil),
		},
		{
			"deactivate_stake", (*DeactivateStake)(nil),
		},
		{
			"emergency_unstake", (*EmergencyUnstake)(nil),
		},
		{
			"partial_unstake", (*PartialUnstake)(nil),
		},
		{
			"merge_stakes", (*MergeStakes)(nil),
		},
		{
			"redelegate", (*Redelegate)(nil),
		},
		{
			"pause", (*Pause)(nil),
		},
		{
			"resume", (*Resume)(nil),
		},
		{
			"withdraw_stake_account", (*WithdrawStakeAccount)(nil),
		},
		{
			"realloc_validator_list", (*ReallocValidatorList)(nil),
		},
		{
			"realloc_stake_list", (*ReallocStakeList)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
