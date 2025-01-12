package types

import (
	// this line is used by starport scaffolding # genesis/types/import
	fmt "fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		VestingAccountList: []VestingAccount{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in vestingAccount
	vestingAccountIdMap := make(map[uint64]bool)
	vestingAccountCount := gs.GetVestingAccountCount()
	for _, elem := range gs.VestingAccountList {
		if _, ok := vestingAccountIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for vestingAccount")
		}
		if elem.Id >= vestingAccountCount {
			return fmt.Errorf("vestingAccount id should be lower or equal than the last id")
		}
		vestingAccountIdMap[elem.Id] = true
		err := elem.Validate()
		if err != nil {
			return err
		}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	err := gs.validateVestingTypes()
	if err != nil {
		return err
	}
	err = gs.validateAccountVestingPools()
	if err != nil {
		return err
	}
	return gs.Params.Validate()
}

func (gs GenesisState) validateVestingTypes() error {
	vts := gs.VestingTypes
	for _, vt := range vts {
		numOfNames := 0
		for _, vtCheck := range vts {
			if vt.Name == vtCheck.Name {
				numOfNames++
			}
			if numOfNames > 1 {
				return fmt.Errorf("vesting type with name: %s defined more than once", vt.Name)
			}
		}
		err := vt.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (gs GenesisState) validateAccountVestingPools() error {
	avts := gs.AccountVestingPools
	vts := gs.VestingTypes
	for _, avt := range avts {
		err := avt.Validate()
		if err != nil {
			return err
		}
		numOfAddress := 0

		for _, avtCheck := range avts {
			if avt.Address == avtCheck.Address {
				numOfAddress++
			}
			if numOfAddress > 1 {
				return fmt.Errorf("account vesting pools with address: %s defined more than once", avt.Address)
			}
		}
		err = avt.ValidateAgainstVestingTypes(vts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (gst GenesisVestingType) Validate() error {
	if len(gst.Name) == 0 {
		return fmt.Errorf("vesting type has no name")
	}

	duration, err := DurationFromUnits(PeriodUnit(gst.LockupPeriodUnit), gst.LockupPeriod)
	if err != nil {
		return fmt.Errorf("LockupPeriodUnit of veting type: %s error: %w", gst.Name, err)
	}
	if duration < 0 {
		return fmt.Errorf("LockupPeriod of veting type: %s less than 0", gst.Name)
	}
	duration, err = DurationFromUnits(PeriodUnit(gst.VestingPeriodUnit), gst.VestingPeriod)
	if err != nil {
		return fmt.Errorf("VestingPeriodUnit of veting type: %s error: %w", gst.Name, err)
	}
	if duration < 0 {
		return fmt.Errorf("VestingPeriod of veting type: %s less than 0", gst.Name)
	}
	return nil
}
