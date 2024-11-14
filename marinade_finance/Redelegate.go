// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Redelegate is the `redelegate` instruction.
type Redelegate struct {
	StakeIndex           *uint32
	SourceValidatorIndex *uint32
	DestValidatorIndex   *uint32

	// [0] = [WRITE] state
	//
	// [1] = [WRITE] validatorList
	//
	// [2] = [WRITE] stakeList
	//
	// [3] = [WRITE] stakeAccount
	//
	// [4] = [] stakeDepositAuthority
	//
	// [5] = [] reservePda
	//
	// [6] = [WRITE, SIGNER] splitStakeAccount
	//
	// [7] = [WRITE, SIGNER] splitStakeRentPayer
	//
	// [8] = [] destValidatorAccount
	//
	// [9] = [WRITE, SIGNER] redelegateStakeAccount
	//
	// [10] = [] clock
	//
	// [11] = [] stakeHistory
	//
	// [12] = [] stakeConfig
	//
	// [13] = [] systemProgram
	//
	// [14] = [] stakeProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRedelegateInstructionBuilder creates a new `Redelegate` instruction builder.
func NewRedelegateInstructionBuilder() *Redelegate {
	nd := &Redelegate{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 15),
	}
	return nd
}

// SetStakeIndex sets the "stakeIndex" parameter.
func (inst *Redelegate) SetStakeIndex(stakeIndex uint32) *Redelegate {
	inst.StakeIndex = &stakeIndex
	return inst
}

// SetSourceValidatorIndex sets the "sourceValidatorIndex" parameter.
func (inst *Redelegate) SetSourceValidatorIndex(sourceValidatorIndex uint32) *Redelegate {
	inst.SourceValidatorIndex = &sourceValidatorIndex
	return inst
}

// SetDestValidatorIndex sets the "destValidatorIndex" parameter.
func (inst *Redelegate) SetDestValidatorIndex(destValidatorIndex uint32) *Redelegate {
	inst.DestValidatorIndex = &destValidatorIndex
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *Redelegate) SetStateAccount(state ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *Redelegate) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetValidatorListAccount sets the "validatorList" account.
func (inst *Redelegate) SetValidatorListAccount(validatorList ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

// GetValidatorListAccount gets the "validatorList" account.
func (inst *Redelegate) GetValidatorListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStakeListAccount sets the "stakeList" account.
func (inst *Redelegate) SetStakeListAccount(stakeList ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(stakeList).WRITE()
	return inst
}

// GetStakeListAccount gets the "stakeList" account.
func (inst *Redelegate) GetStakeListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetStakeAccountAccount sets the "stakeAccount" account.
func (inst *Redelegate) SetStakeAccountAccount(stakeAccount ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(stakeAccount).WRITE()
	return inst
}

// GetStakeAccountAccount gets the "stakeAccount" account.
func (inst *Redelegate) GetStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetStakeDepositAuthorityAccount sets the "stakeDepositAuthority" account.
func (inst *Redelegate) SetStakeDepositAuthorityAccount(stakeDepositAuthority ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(stakeDepositAuthority)
	return inst
}

// GetStakeDepositAuthorityAccount gets the "stakeDepositAuthority" account.
func (inst *Redelegate) GetStakeDepositAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReservePdaAccount sets the "reservePda" account.
func (inst *Redelegate) SetReservePdaAccount(reservePda ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(reservePda)
	return inst
}

// GetReservePdaAccount gets the "reservePda" account.
func (inst *Redelegate) GetReservePdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSplitStakeAccountAccount sets the "splitStakeAccount" account.
func (inst *Redelegate) SetSplitStakeAccountAccount(splitStakeAccount ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(splitStakeAccount).WRITE().SIGNER()
	return inst
}

// GetSplitStakeAccountAccount gets the "splitStakeAccount" account.
func (inst *Redelegate) GetSplitStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSplitStakeRentPayerAccount sets the "splitStakeRentPayer" account.
func (inst *Redelegate) SetSplitStakeRentPayerAccount(splitStakeRentPayer ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(splitStakeRentPayer).WRITE().SIGNER()
	return inst
}

// GetSplitStakeRentPayerAccount gets the "splitStakeRentPayer" account.
func (inst *Redelegate) GetSplitStakeRentPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetDestValidatorAccountAccount sets the "destValidatorAccount" account.
func (inst *Redelegate) SetDestValidatorAccountAccount(destValidatorAccount ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(destValidatorAccount)
	return inst
}

// GetDestValidatorAccountAccount gets the "destValidatorAccount" account.
func (inst *Redelegate) GetDestValidatorAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetRedelegateStakeAccountAccount sets the "redelegateStakeAccount" account.
func (inst *Redelegate) SetRedelegateStakeAccountAccount(redelegateStakeAccount ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(redelegateStakeAccount).WRITE().SIGNER()
	return inst
}

// GetRedelegateStakeAccountAccount gets the "redelegateStakeAccount" account.
func (inst *Redelegate) GetRedelegateStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetClockAccount sets the "clock" account.
func (inst *Redelegate) SetClockAccount(clock ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *Redelegate) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetStakeHistoryAccount sets the "stakeHistory" account.
func (inst *Redelegate) SetStakeHistoryAccount(stakeHistory ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(stakeHistory)
	return inst
}

// GetStakeHistoryAccount gets the "stakeHistory" account.
func (inst *Redelegate) GetStakeHistoryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetStakeConfigAccount sets the "stakeConfig" account.
func (inst *Redelegate) SetStakeConfigAccount(stakeConfig ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(stakeConfig)
	return inst
}

// GetStakeConfigAccount gets the "stakeConfig" account.
func (inst *Redelegate) GetStakeConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Redelegate) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Redelegate) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetStakeProgramAccount sets the "stakeProgram" account.
func (inst *Redelegate) SetStakeProgramAccount(stakeProgram ag_solanago.PublicKey) *Redelegate {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(stakeProgram)
	return inst
}

// GetStakeProgramAccount gets the "stakeProgram" account.
func (inst *Redelegate) GetStakeProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

func (inst Redelegate) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Redelegate,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Redelegate) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Redelegate) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.StakeIndex == nil {
			return errors.New("StakeIndex parameter is not set")
		}
		if inst.SourceValidatorIndex == nil {
			return errors.New("SourceValidatorIndex parameter is not set")
		}
		if inst.DestValidatorIndex == nil {
			return errors.New("DestValidatorIndex parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ValidatorList is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.StakeList is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.StakeAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.StakeDepositAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ReservePda is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SplitStakeAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SplitStakeRentPayer is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.DestValidatorAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.RedelegateStakeAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Clock is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.StakeHistory is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.StakeConfig is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.StakeProgram is not set")
		}
	}
	return nil
}

func (inst *Redelegate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Redelegate")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("          StakeIndex", *inst.StakeIndex))
						paramsBranch.Child(ag_format.Param("SourceValidatorIndex", *inst.SourceValidatorIndex))
						paramsBranch.Child(ag_format.Param("  DestValidatorIndex", *inst.DestValidatorIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=15]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        validatorList", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            stakeList", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                stake", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("stakeDepositAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           reservePda", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           splitStake", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("  splitStakeRentPayer", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("        destValidator", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("      redelegateStake", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                clock", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("         stakeHistory", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("          stakeConfig", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("        systemProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("         stakeProgram", inst.AccountMetaSlice.Get(14)))
					})
				})
		})
}

func (obj Redelegate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `StakeIndex` param:
	err = encoder.Encode(obj.StakeIndex)
	if err != nil {
		return err
	}
	// Serialize `SourceValidatorIndex` param:
	err = encoder.Encode(obj.SourceValidatorIndex)
	if err != nil {
		return err
	}
	// Serialize `DestValidatorIndex` param:
	err = encoder.Encode(obj.DestValidatorIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Redelegate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `StakeIndex`:
	err = decoder.Decode(&obj.StakeIndex)
	if err != nil {
		return err
	}
	// Deserialize `SourceValidatorIndex`:
	err = decoder.Decode(&obj.SourceValidatorIndex)
	if err != nil {
		return err
	}
	// Deserialize `DestValidatorIndex`:
	err = decoder.Decode(&obj.DestValidatorIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewRedelegateInstruction declares a new Redelegate instruction with the provided parameters and accounts.
func NewRedelegateInstruction(
	// Parameters:
	stakeIndex uint32,
	sourceValidatorIndex uint32,
	destValidatorIndex uint32,
	// Accounts:
	state ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	stakeList ag_solanago.PublicKey,
	stakeAccount ag_solanago.PublicKey,
	stakeDepositAuthority ag_solanago.PublicKey,
	reservePda ag_solanago.PublicKey,
	splitStakeAccount ag_solanago.PublicKey,
	splitStakeRentPayer ag_solanago.PublicKey,
	destValidatorAccount ag_solanago.PublicKey,
	redelegateStakeAccount ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	stakeHistory ag_solanago.PublicKey,
	stakeConfig ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey) *Redelegate {
	return NewRedelegateInstructionBuilder().
		SetStakeIndex(stakeIndex).
		SetSourceValidatorIndex(sourceValidatorIndex).
		SetDestValidatorIndex(destValidatorIndex).
		SetStateAccount(state).
		SetValidatorListAccount(validatorList).
		SetStakeListAccount(stakeList).
		SetStakeAccountAccount(stakeAccount).
		SetStakeDepositAuthorityAccount(stakeDepositAuthority).
		SetReservePdaAccount(reservePda).
		SetSplitStakeAccountAccount(splitStakeAccount).
		SetSplitStakeRentPayerAccount(splitStakeRentPayer).
		SetDestValidatorAccountAccount(destValidatorAccount).
		SetRedelegateStakeAccountAccount(redelegateStakeAccount).
		SetClockAccount(clock).
		SetStakeHistoryAccount(stakeHistory).
		SetStakeConfigAccount(stakeConfig).
		SetSystemProgramAccount(systemProgram).
		SetStakeProgramAccount(stakeProgram)
}
