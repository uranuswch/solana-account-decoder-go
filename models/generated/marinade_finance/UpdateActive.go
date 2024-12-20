// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateActive is the `updateActive` instruction.
type UpdateActive struct {
	StakeIndex     *uint32
	ValidatorIndex *uint32

	// ····· common: [0] = [WRITE] state
	//
	// ············· [1] = [WRITE] stakeList
	//
	// ············· [2] = [WRITE] stakeAccount
	//
	// ············· [3] = [] stakeWithdrawAuthority
	//
	// ············· [4] = [WRITE] reservePda
	//
	// ············· [5] = [WRITE] msolMint
	//
	// ············· [6] = [] msolMintAuthority
	//
	// ············· [7] = [WRITE] treasuryMsolAccount
	//
	// ············· [8] = [] clock
	//
	// ············· [9] = [] stakeHistory
	//
	// ············· [10] = [] stakeProgram
	//
	// ············· [11] = [] tokenProgram
	//
	// [12] = [WRITE] validatorList
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateActiveInstructionBuilder creates a new `UpdateActive` instruction builder.
func NewUpdateActiveInstructionBuilder() *UpdateActive {
	nd := &UpdateActive{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetStakeIndex sets the "stakeIndex" parameter.
func (inst *UpdateActive) SetStakeIndex(stakeIndex uint32) *UpdateActive {
	inst.StakeIndex = &stakeIndex
	return inst
}

// SetValidatorIndex sets the "validatorIndex" parameter.
func (inst *UpdateActive) SetValidatorIndex(validatorIndex uint32) *UpdateActive {
	inst.ValidatorIndex = &validatorIndex
	return inst
}

type UpdateActiveCommonAccountsBuilder struct {
	ag_solanago.AccountMetaSlice `bin:"-"`
}

func NewUpdateActiveCommonAccountsBuilder() *UpdateActiveCommonAccountsBuilder {
	return &UpdateActiveCommonAccountsBuilder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
}

func (inst *UpdateActive) SetCommonAccountsFromBuilder(updateActiveCommonAccountsBuilder *UpdateActiveCommonAccountsBuilder) *UpdateActive {
	inst.AccountMetaSlice[0] = updateActiveCommonAccountsBuilder.GetStateAccount()
	inst.AccountMetaSlice[1] = updateActiveCommonAccountsBuilder.GetStakeListAccount()
	inst.AccountMetaSlice[2] = updateActiveCommonAccountsBuilder.GetStakeAccountAccount()
	inst.AccountMetaSlice[3] = updateActiveCommonAccountsBuilder.GetStakeWithdrawAuthorityAccount()
	inst.AccountMetaSlice[4] = updateActiveCommonAccountsBuilder.GetReservePdaAccount()
	inst.AccountMetaSlice[5] = updateActiveCommonAccountsBuilder.GetMsolMintAccount()
	inst.AccountMetaSlice[6] = updateActiveCommonAccountsBuilder.GetMsolMintAuthorityAccount()
	inst.AccountMetaSlice[7] = updateActiveCommonAccountsBuilder.GetTreasuryMsolAccountAccount()
	inst.AccountMetaSlice[8] = updateActiveCommonAccountsBuilder.GetClockAccount()
	inst.AccountMetaSlice[9] = updateActiveCommonAccountsBuilder.GetStakeHistoryAccount()
	inst.AccountMetaSlice[10] = updateActiveCommonAccountsBuilder.GetStakeProgramAccount()
	inst.AccountMetaSlice[11] = updateActiveCommonAccountsBuilder.GetTokenProgramAccount()
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetStateAccount(state ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStakeListAccount sets the "stakeList" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetStakeListAccount(stakeList ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(stakeList).WRITE()
	return inst
}

// GetStakeListAccount gets the "stakeList" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetStakeListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStakeAccountAccount sets the "stakeAccount" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetStakeAccountAccount(stakeAccount ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(stakeAccount).WRITE()
	return inst
}

// GetStakeAccountAccount gets the "stakeAccount" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetStakeWithdrawAuthorityAccount sets the "stakeWithdrawAuthority" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetStakeWithdrawAuthorityAccount(stakeWithdrawAuthority ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(stakeWithdrawAuthority)
	return inst
}

// GetStakeWithdrawAuthorityAccount gets the "stakeWithdrawAuthority" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetStakeWithdrawAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReservePdaAccount sets the "reservePda" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetReservePdaAccount(reservePda ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(reservePda).WRITE()
	return inst
}

// GetReservePdaAccount gets the "reservePda" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetReservePdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMsolMintAccount sets the "msolMint" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetMsolMintAccount(msolMint ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(msolMint).WRITE()
	return inst
}

// GetMsolMintAccount gets the "msolMint" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetMsolMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMsolMintAuthorityAccount sets the "msolMintAuthority" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetMsolMintAuthorityAccount(msolMintAuthority ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(msolMintAuthority)
	return inst
}

// GetMsolMintAuthorityAccount gets the "msolMintAuthority" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetMsolMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTreasuryMsolAccountAccount sets the "treasuryMsolAccount" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetTreasuryMsolAccountAccount(treasuryMsolAccount ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(treasuryMsolAccount).WRITE()
	return inst
}

// GetTreasuryMsolAccountAccount gets the "treasuryMsolAccount" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetTreasuryMsolAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetClockAccount sets the "clock" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetClockAccount(clock ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetStakeHistoryAccount sets the "stakeHistory" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetStakeHistoryAccount(stakeHistory ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(stakeHistory)
	return inst
}

// GetStakeHistoryAccount gets the "stakeHistory" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetStakeHistoryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetStakeProgramAccount sets the "stakeProgram" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetStakeProgramAccount(stakeProgram ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(stakeProgram)
	return inst
}

// GetStakeProgramAccount gets the "stakeProgram" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetStakeProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *UpdateActiveCommonAccountsBuilder) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *UpdateActiveCommonAccountsBuilder {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *UpdateActiveCommonAccountsBuilder) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetValidatorListAccount sets the "validatorList" account.
func (inst *UpdateActive) SetValidatorListAccount(validatorList ag_solanago.PublicKey) *UpdateActive {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

// GetValidatorListAccount gets the "validatorList" account.
func (inst *UpdateActive) GetValidatorListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst UpdateActive) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateActive,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateActive) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateActive) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.StakeIndex == nil {
			return errors.New("StakeIndex parameter is not set")
		}
		if inst.ValidatorIndex == nil {
			return errors.New("ValidatorIndex parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CommonState is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CommonStakeList is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.CommonStakeAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CommonStakeWithdrawAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CommonReservePda is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.CommonMsolMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.CommonMsolMintAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CommonTreasuryMsolAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.CommonClock is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.CommonStakeHistory is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.CommonStakeProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.CommonTokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.ValidatorList is not set")
		}
	}
	return nil
}

func (inst *UpdateActive) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateActive")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    StakeIndex", *inst.StakeIndex))
						paramsBranch.Child(ag_format.Param("ValidatorIndex", *inst.ValidatorIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 common/state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             common/stakeList", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                 common/stake", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("common/stakeWithdrawAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            common/reservePda", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("              common/msolMint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     common/msolMintAuthority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("          common/treasuryMsol", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                 common/clock", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("          common/stakeHistory", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("          common/stakeProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("          common/tokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("                validatorList", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj UpdateActive) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `StakeIndex` param:
	err = encoder.Encode(obj.StakeIndex)
	if err != nil {
		return err
	}
	// Serialize `ValidatorIndex` param:
	err = encoder.Encode(obj.ValidatorIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateActive) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `StakeIndex`:
	err = decoder.Decode(&obj.StakeIndex)
	if err != nil {
		return err
	}
	// Deserialize `ValidatorIndex`:
	err = decoder.Decode(&obj.ValidatorIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateActiveInstruction declares a new UpdateActive instruction with the provided parameters and accounts.
func NewUpdateActiveInstruction(
	// Parameters:
	stakeIndex uint32,
	validatorIndex uint32,
	// Accounts:
	commonState ag_solanago.PublicKey,
	commonStakeList ag_solanago.PublicKey,
	commonStakeAccount ag_solanago.PublicKey,
	commonStakeWithdrawAuthority ag_solanago.PublicKey,
	commonReservePda ag_solanago.PublicKey,
	commonMsolMint ag_solanago.PublicKey,
	commonMsolMintAuthority ag_solanago.PublicKey,
	commonTreasuryMsolAccount ag_solanago.PublicKey,
	commonClock ag_solanago.PublicKey,
	commonStakeHistory ag_solanago.PublicKey,
	commonStakeProgram ag_solanago.PublicKey,
	commonTokenProgram ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey) *UpdateActive {
	return NewUpdateActiveInstructionBuilder().
		SetStakeIndex(stakeIndex).
		SetValidatorIndex(validatorIndex).
		SetCommonAccountsFromBuilder(
			NewUpdateActiveCommonAccountsBuilder().
				SetStateAccount(commonState).
				SetStakeListAccount(commonStakeList).
				SetStakeAccountAccount(commonStakeAccount).
				SetStakeWithdrawAuthorityAccount(commonStakeWithdrawAuthority).
				SetReservePdaAccount(commonReservePda).
				SetMsolMintAccount(commonMsolMint).
				SetMsolMintAuthorityAccount(commonMsolMintAuthority).
				SetTreasuryMsolAccountAccount(commonTreasuryMsolAccount).
				SetClockAccount(commonClock).
				SetStakeHistoryAccount(commonStakeHistory).
				SetStakeProgramAccount(commonStakeProgram).
				SetTokenProgramAccount(commonTokenProgram),
		).
		SetValidatorListAccount(validatorList)
}
