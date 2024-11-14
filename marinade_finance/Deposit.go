// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Deposit is the `deposit` instruction.
type Deposit struct {
	Lamports *uint64

	// [0] = [WRITE] state
	//
	// [1] = [WRITE] msolMint
	//
	// [2] = [WRITE] liqPoolSolLegPda
	//
	// [3] = [WRITE] liqPoolMsolLeg
	//
	// [4] = [] liqPoolMsolLegAuthority
	//
	// [5] = [WRITE] reservePda
	//
	// [6] = [WRITE, SIGNER] transferFrom
	//
	// [7] = [WRITE] mintTo
	// ··········· user mSOL Token account to send the mSOL
	//
	// [8] = [] msolMintAuthority
	//
	// [9] = [] systemProgram
	//
	// [10] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDepositInstructionBuilder creates a new `Deposit` instruction builder.
func NewDepositInstructionBuilder() *Deposit {
	nd := &Deposit{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetLamports sets the "lamports" parameter.
func (inst *Deposit) SetLamports(lamports uint64) *Deposit {
	inst.Lamports = &lamports
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *Deposit) SetStateAccount(state ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *Deposit) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMsolMintAccount sets the "msolMint" account.
func (inst *Deposit) SetMsolMintAccount(msolMint ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(msolMint).WRITE()
	return inst
}

// GetMsolMintAccount gets the "msolMint" account.
func (inst *Deposit) GetMsolMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetLiqPoolSolLegPdaAccount sets the "liqPoolSolLegPda" account.
func (inst *Deposit) SetLiqPoolSolLegPdaAccount(liqPoolSolLegPda ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(liqPoolSolLegPda).WRITE()
	return inst
}

// GetLiqPoolSolLegPdaAccount gets the "liqPoolSolLegPda" account.
func (inst *Deposit) GetLiqPoolSolLegPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetLiqPoolMsolLegAccount sets the "liqPoolMsolLeg" account.
func (inst *Deposit) SetLiqPoolMsolLegAccount(liqPoolMsolLeg ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(liqPoolMsolLeg).WRITE()
	return inst
}

// GetLiqPoolMsolLegAccount gets the "liqPoolMsolLeg" account.
func (inst *Deposit) GetLiqPoolMsolLegAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetLiqPoolMsolLegAuthorityAccount sets the "liqPoolMsolLegAuthority" account.
func (inst *Deposit) SetLiqPoolMsolLegAuthorityAccount(liqPoolMsolLegAuthority ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(liqPoolMsolLegAuthority)
	return inst
}

// GetLiqPoolMsolLegAuthorityAccount gets the "liqPoolMsolLegAuthority" account.
func (inst *Deposit) GetLiqPoolMsolLegAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReservePdaAccount sets the "reservePda" account.
func (inst *Deposit) SetReservePdaAccount(reservePda ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(reservePda).WRITE()
	return inst
}

// GetReservePdaAccount gets the "reservePda" account.
func (inst *Deposit) GetReservePdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTransferFromAccount sets the "transferFrom" account.
func (inst *Deposit) SetTransferFromAccount(transferFrom ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(transferFrom).WRITE().SIGNER()
	return inst
}

// GetTransferFromAccount gets the "transferFrom" account.
func (inst *Deposit) GetTransferFromAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMintToAccount sets the "mintTo" account.
// user mSOL Token account to send the mSOL
func (inst *Deposit) SetMintToAccount(mintTo ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(mintTo).WRITE()
	return inst
}

// GetMintToAccount gets the "mintTo" account.
// user mSOL Token account to send the mSOL
func (inst *Deposit) GetMintToAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMsolMintAuthorityAccount sets the "msolMintAuthority" account.
func (inst *Deposit) SetMsolMintAuthorityAccount(msolMintAuthority ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(msolMintAuthority)
	return inst
}

// GetMsolMintAuthorityAccount gets the "msolMintAuthority" account.
func (inst *Deposit) GetMsolMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Deposit) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Deposit) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Deposit) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Deposit) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst Deposit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Deposit,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Deposit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Deposit) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Lamports == nil {
			return errors.New("Lamports parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MsolMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.LiqPoolSolLegPda is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.LiqPoolMsolLeg is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.LiqPoolMsolLegAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ReservePda is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TransferFrom is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.MintTo is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MsolMintAuthority is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *Deposit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Deposit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Lamports", *inst.Lamports))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("               msolMint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       liqPoolSolLegPda", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         liqPoolMsolLeg", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("liqPoolMsolLegAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("             reservePda", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           transferFrom", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                 mintTo", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("      msolMintAuthority", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("          systemProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("           tokenProgram", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj Deposit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Lamports` param:
	err = encoder.Encode(obj.Lamports)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Deposit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Lamports`:
	err = decoder.Decode(&obj.Lamports)
	if err != nil {
		return err
	}
	return nil
}

// NewDepositInstruction declares a new Deposit instruction with the provided parameters and accounts.
func NewDepositInstruction(
	// Parameters:
	lamports uint64,
	// Accounts:
	state ag_solanago.PublicKey,
	msolMint ag_solanago.PublicKey,
	liqPoolSolLegPda ag_solanago.PublicKey,
	liqPoolMsolLeg ag_solanago.PublicKey,
	liqPoolMsolLegAuthority ag_solanago.PublicKey,
	reservePda ag_solanago.PublicKey,
	transferFrom ag_solanago.PublicKey,
	mintTo ag_solanago.PublicKey,
	msolMintAuthority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *Deposit {
	return NewDepositInstructionBuilder().
		SetLamports(lamports).
		SetStateAccount(state).
		SetMsolMintAccount(msolMint).
		SetLiqPoolSolLegPdaAccount(liqPoolSolLegPda).
		SetLiqPoolMsolLegAccount(liqPoolMsolLeg).
		SetLiqPoolMsolLegAuthorityAccount(liqPoolMsolLegAuthority).
		SetReservePdaAccount(reservePda).
		SetTransferFromAccount(transferFrom).
		SetMintToAccount(mintTo).
		SetMsolMintAuthorityAccount(msolMintAuthority).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram)
}
