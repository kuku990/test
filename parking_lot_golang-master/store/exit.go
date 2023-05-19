package store

import (
	"fmt"
	"parking_lot/errors"
	"parking_lot/schema"
	"strconv"
)

type leaveStore struct {
	*store
}

// NewParkStore returns new store object
func NewleaveStore(st *store) *leaveStore {
	pl := &leaveStore{st}
	return pl
}

//exit function

func (pl *leaveStore) Execute(cmd *schema.Command) (string, error) {
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}
	// if err := validateParkReq(cmd.Arguments); err != nil {
	// 	return "", err
	// }
	// slotInfo, err := ParkingLot.FindSlotByRegNumber(cmd.Arguments[0])

	slotInfoList := ParkingLot.Slots

	slotId, err := strconv.Atoi(cmd.Arguments[0])
	if err != nil {
		return "", err
	}
	//********slotnumber more then max value

	// slotInfo.ID = uint(slotId)
	slotInfo := slotInfoList[slotId-1]
	exitRequestErr := slotInfo.ExitPark()
	if exitRequestErr != nil {
		return "", errors.ErrCarNotFound
	}
	// sucessMsg  = fmt.Sprintf("Slot number '%s' is free")
	return fmt.Sprintf(SlotIsFreeInfo, slotInfo.ID), nil

}
