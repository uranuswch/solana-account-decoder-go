// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Resume is the `resume` instruction.
type Resume struct {

	// [0] = [WRITE] state
	//
	// [1] = [SIGNER] pauseAuthority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewResumeInstructionBuilder creates a new `Resume` instruction builder.
func NewResumeInstructionBuilder() *Resume {
	nd := &Resume{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *Resume) SetStateAccount(state ag_solanago.PublicKey) *Resume {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *Resume) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPauseAuthorityAccount sets the "pauseAuthority" account.
func (inst *Resume) SetPauseAuthorityAccount(pauseAuthority ag_solanago.PublicKey) *Resume {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(pauseAuthority).SIGNER()
	return inst
}

// GetPauseAuthorityAccount gets the "pauseAuthority" account.
func (inst *Resume) GetPauseAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst Resume) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Resume,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Resume) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Resume) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PauseAuthority is not set")
		}
	}
	return nil
}

func (inst *Resume) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Resume")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("pauseAuthority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj Resume) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Resume) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewResumeInstruction declares a new Resume instruction with the provided parameters and accounts.
func NewResumeInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	pauseAuthority ag_solanago.PublicKey) *Resume {
	return NewResumeInstructionBuilder().
		SetStateAccount(state).
		SetPauseAuthorityAccount(pauseAuthority)
}
