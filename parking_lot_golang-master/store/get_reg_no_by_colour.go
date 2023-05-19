package store

import (
	"strings"

	"parking_lot/errors"
	"parking_lot/schema"
	"parking_lot/utils"
)

type getVehiclesByColourStore struct {
	*store
}

// NewgetVehiclesByColour returns a new store object
func NewgetVehiclesByColour(st *store) *getVehiclesByColourStore {
	pl := &getVehiclesByColourStore{st}
	return pl
}

func (pl *getVehiclesByColourStore) Execute(cmd *schema.Command) (string, error) {
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}
	if !utils.IsValidString(cmd.Arguments[0]) {
		return "", errors.ErrInvalidColour
	}

	colour := cmd.Arguments[0]
	slotsList, err := ParkingLot.FindSlotByColor(colour)
	if err != nil {
		return "", err
	}

	vehicleList := make([]string, len(slotsList))
	for i, slot := range slotsList {
		vehicleList[i] = slot.Vehicle.RegistrationNumber
	}

	return strings.Join(vehicleList, ","), nil
}
