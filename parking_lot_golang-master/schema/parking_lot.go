package schema

import (
	"fmt"
	"parking_lot/errors"
	"time"
)

// ParkingLot struct holds all the parking lot information and parking history
type ParkingLot struct {
	Name        string         `json:"name"`
	Floor       string         `json:"floor"`
	TotalBlocks int            `json:"total_blocks"`
	BlockHeight int            `json:"block_height"`
	TotalSlots  int            `json:"total_slots"`
	Address     string         `json:"address"`
	Pincode     string         `json:"pincode"`
	Slots       []*Slot        `json:"slots"`
	ParkHistory []*ParkHistory `json:"park_history"`
}

// ParkHistory holds the parking information
type ParkHistory struct {
	SlotID             uint
	RegistrationNumber string
	Colour             string
	CreatedAt          time.Time
}

// FirstAvailableSlot returns the first available slot to park Vehicle
func (pl *ParkingLot) FirstAvailableSlot() (*Slot, error) {
	for _, slot := range pl.Slots {
		if slot.IsSlotAvailable() {
			return slot, nil
		}
	}

	return nil, errors.ErrParkingSlotsFull
}

//**********************************************
//find vehicle by registration number
func (pl *ParkingLot) FindSlotByRegNumber(regNo string) (*Slot, error) {
	for _, slot := range pl.Slots {
		if slot.IsSlotOccupied() {
			currentVehicle := slot.Vehicle

			if currentVehicle.IsVehicleRegNoMatched(regNo) {
				return slot, nil
			}
		}
	}
	return nil, errors.ErrNoVechileFoundByRegNo
}

//find vehicle by registration number
func (pl *ParkingLot) FindSlotByColor(colour string) ([]*Slot, error) {
	slotlist := []*Slot{}
	for _, slot := range pl.Slots {
		if slot.IsSlotOccupied() {
			currentVehicle := slot.Vehicle
			if currentVehicle.IsVehicleColurMatched(colour) {
				slotlist = append(slotlist, slot)
				fmt.Println(slotlist)
			}
		}
	}
	if len(slotlist) > 0 {
		return slotlist, nil
	}
	return slotlist, errors.ErrNoCarFoundByColour(colour)
}
