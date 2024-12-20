// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ConfigValidatorSystem is the `configValidatorSystem` instruction.
type ConfigValidatorSystem struct {
	ExtraRuns *uint32

	// [0] = [WRITE] state
	//
	// [1] = [SIGNER] managerAuthority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewConfigValidatorSystemInstructionBuilder creates a new `ConfigValidatorSystem` instruction builder.
func NewConfigValidatorSystemInstructionBuilder() *ConfigValidatorSystem {
	nd := &ConfigValidatorSystem{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetExtraRuns sets the "extraRuns" parameter.
func (inst *ConfigValidatorSystem) SetExtraRuns(extraRuns uint32) *ConfigValidatorSystem {
	inst.ExtraRuns = &extraRuns
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *ConfigValidatorSystem) SetStateAccount(state ag_solanago.PublicKey) *ConfigValidatorSystem {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *ConfigValidatorSystem) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetManagerAuthorityAccount sets the "managerAuthority" account.
func (inst *ConfigValidatorSystem) SetManagerAuthorityAccount(managerAuthority ag_solanago.PublicKey) *ConfigValidatorSystem {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(managerAuthority).SIGNER()
	return inst
}

// GetManagerAuthorityAccount gets the "managerAuthority" account.
func (inst *ConfigValidatorSystem) GetManagerAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst ConfigValidatorSystem) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ConfigValidatorSystem,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ConfigValidatorSystem) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ConfigValidatorSystem) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ExtraRuns == nil {
			return errors.New("ExtraRuns parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ManagerAuthority is not set")
		}
	}
	return nil
}

func (inst *ConfigValidatorSystem) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ConfigValidatorSystem")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ExtraRuns", *inst.ExtraRuns))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("managerAuthority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj ConfigValidatorSystem) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ExtraRuns` param:
	err = encoder.Encode(obj.ExtraRuns)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ConfigValidatorSystem) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ExtraRuns`:
	err = decoder.Decode(&obj.ExtraRuns)
	if err != nil {
		return err
	}
	return nil
}

// NewConfigValidatorSystemInstruction declares a new ConfigValidatorSystem instruction with the provided parameters and accounts.
func NewConfigValidatorSystemInstruction(
	// Parameters:
	extraRuns uint32,
	// Accounts:
	state ag_solanago.PublicKey,
	managerAuthority ag_solanago.PublicKey) *ConfigValidatorSystem {
	return NewConfigValidatorSystemInstructionBuilder().
		SetExtraRuns(extraRuns).
		SetStateAccount(state).
		SetManagerAuthorityAccount(managerAuthority)
}
