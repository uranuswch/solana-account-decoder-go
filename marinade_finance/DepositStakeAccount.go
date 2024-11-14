// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DepositStakeAccount is the `depositStakeAccount` instruction.
type DepositStakeAccount struct {
	ValidatorIndex *uint32

	// [0] = [WRITE] state
	//
	// [1] = [WRITE] validatorList
	//
	// [2] = [WRITE] stakeList
	//
	// [3] = [WRITE] stakeAccount
	//
	// [4] = [SIGNER] stakeAuthority
	//
	// [5] = [WRITE] duplicationFlag
	//
	// [6] = [WRITE, SIGNER] rentPayer
	//
	// [7] = [WRITE] msolMint
	//
	// [8] = [WRITE] mintTo
	// ··········· user mSOL Token account to send the mSOL
	//
	// [9] = [] msolMintAuthority
	//
	// [10] = [] clock
	//
	// [11] = [] rent
	//
	// [12] = [] systemProgram
	//
	// [13] = [] tokenProgram
	//
	// [14] = [] stakeProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDepositStakeAccountInstructionBuilder creates a new `DepositStakeAccount` instruction builder.
func NewDepositStakeAccountInstructionBuilder() *DepositStakeAccount {
	nd := &DepositStakeAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 15),
	}
	return nd
}

// SetValidatorIndex sets the "validatorIndex" parameter.
func (inst *DepositStakeAccount) SetValidatorIndex(validatorIndex uint32) *DepositStakeAccount {
	inst.ValidatorIndex = &validatorIndex
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *DepositStakeAccount) SetStateAccount(state ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *DepositStakeAccount) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetValidatorListAccount sets the "validatorList" account.
func (inst *DepositStakeAccount) SetValidatorListAccount(validatorList ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

// GetValidatorListAccount gets the "validatorList" account.
func (inst *DepositStakeAccount) GetValidatorListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStakeListAccount sets the "stakeList" account.
func (inst *DepositStakeAccount) SetStakeListAccount(stakeList ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(stakeList).WRITE()
	return inst
}

// GetStakeListAccount gets the "stakeList" account.
func (inst *DepositStakeAccount) GetStakeListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetStakeAccountAccount sets the "stakeAccount" account.
func (inst *DepositStakeAccount) SetStakeAccountAccount(stakeAccount ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(stakeAccount).WRITE()
	return inst
}

// GetStakeAccountAccount gets the "stakeAccount" account.
func (inst *DepositStakeAccount) GetStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetStakeAuthorityAccount sets the "stakeAuthority" account.
func (inst *DepositStakeAccount) SetStakeAuthorityAccount(stakeAuthority ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(stakeAuthority).SIGNER()
	return inst
}

// GetStakeAuthorityAccount gets the "stakeAuthority" account.
func (inst *DepositStakeAccount) GetStakeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetDuplicationFlagAccount sets the "duplicationFlag" account.
func (inst *DepositStakeAccount) SetDuplicationFlagAccount(duplicationFlag ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(duplicationFlag).WRITE()
	return inst
}

// GetDuplicationFlagAccount gets the "duplicationFlag" account.
func (inst *DepositStakeAccount) GetDuplicationFlagAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetRentPayerAccount sets the "rentPayer" account.
func (inst *DepositStakeAccount) SetRentPayerAccount(rentPayer ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rentPayer).WRITE().SIGNER()
	return inst
}

// GetRentPayerAccount gets the "rentPayer" account.
func (inst *DepositStakeAccount) GetRentPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMsolMintAccount sets the "msolMint" account.
func (inst *DepositStakeAccount) SetMsolMintAccount(msolMint ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(msolMint).WRITE()
	return inst
}

// GetMsolMintAccount gets the "msolMint" account.
func (inst *DepositStakeAccount) GetMsolMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMintToAccount sets the "mintTo" account.
// user mSOL Token account to send the mSOL
func (inst *DepositStakeAccount) SetMintToAccount(mintTo ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(mintTo).WRITE()
	return inst
}

// GetMintToAccount gets the "mintTo" account.
// user mSOL Token account to send the mSOL
func (inst *DepositStakeAccount) GetMintToAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetMsolMintAuthorityAccount sets the "msolMintAuthority" account.
func (inst *DepositStakeAccount) SetMsolMintAuthorityAccount(msolMintAuthority ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(msolMintAuthority)
	return inst
}

// GetMsolMintAuthorityAccount gets the "msolMintAuthority" account.
func (inst *DepositStakeAccount) GetMsolMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetClockAccount sets the "clock" account.
func (inst *DepositStakeAccount) SetClockAccount(clock ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *DepositStakeAccount) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetRentAccount sets the "rent" account.
func (inst *DepositStakeAccount) SetRentAccount(rent ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *DepositStakeAccount) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *DepositStakeAccount) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *DepositStakeAccount) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *DepositStakeAccount) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *DepositStakeAccount) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetStakeProgramAccount sets the "stakeProgram" account.
func (inst *DepositStakeAccount) SetStakeProgramAccount(stakeProgram ag_solanago.PublicKey) *DepositStakeAccount {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(stakeProgram)
	return inst
}

// GetStakeProgramAccount gets the "stakeProgram" account.
func (inst *DepositStakeAccount) GetStakeProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

func (inst DepositStakeAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DepositStakeAccount,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DepositStakeAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DepositStakeAccount) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ValidatorIndex == nil {
			return errors.New("ValidatorIndex parameter is not set")
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
			return errors.New("accounts.StakeAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.DuplicationFlag is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.RentPayer is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.MsolMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MintTo is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.MsolMintAuthority is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Clock is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.StakeProgram is not set")
		}
	}
	return nil
}

func (inst *DepositStakeAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DepositStakeAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ValidatorIndex", *inst.ValidatorIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=15]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("            state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    validatorList", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        stakeList", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            stake", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   stakeAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  duplicationFlag", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("        rentPayer", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         msolMint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("           mintTo", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("msolMintAuthority", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("            clock", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("             rent", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("    systemProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("     tokenProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("     stakeProgram", inst.AccountMetaSlice.Get(14)))
					})
				})
		})
}

func (obj DepositStakeAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ValidatorIndex` param:
	err = encoder.Encode(obj.ValidatorIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DepositStakeAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ValidatorIndex`:
	err = decoder.Decode(&obj.ValidatorIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewDepositStakeAccountInstruction declares a new DepositStakeAccount instruction with the provided parameters and accounts.
func NewDepositStakeAccountInstruction(
	// Parameters:
	validatorIndex uint32,
	// Accounts:
	state ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	stakeList ag_solanago.PublicKey,
	stakeAccount ag_solanago.PublicKey,
	stakeAuthority ag_solanago.PublicKey,
	duplicationFlag ag_solanago.PublicKey,
	rentPayer ag_solanago.PublicKey,
	msolMint ag_solanago.PublicKey,
	mintTo ag_solanago.PublicKey,
	msolMintAuthority ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey) *DepositStakeAccount {
	return NewDepositStakeAccountInstructionBuilder().
		SetValidatorIndex(validatorIndex).
		SetStateAccount(state).
		SetValidatorListAccount(validatorList).
		SetStakeListAccount(stakeList).
		SetStakeAccountAccount(stakeAccount).
		SetStakeAuthorityAccount(stakeAuthority).
		SetDuplicationFlagAccount(duplicationFlag).
		SetRentPayerAccount(rentPayer).
		SetMsolMintAccount(msolMint).
		SetMintToAccount(mintTo).
		SetMsolMintAuthorityAccount(msolMintAuthority).
		SetClockAccount(clock).
		SetRentAccount(rent).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram).
		SetStakeProgramAccount(stakeProgram)
}