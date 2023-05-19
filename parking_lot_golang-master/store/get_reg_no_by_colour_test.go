// package store

// import (
// 	"parking_lot/errors"
// 	"parking_lot/schema"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// var _ = Describe("Parking Lot Store Tests", func() {
// 	var (
// 		connection Store
// 	)
// 	connection = NewStore()
// 	It("Tear Down Store Data", func() {
// 		TearDown()
// 	})

// 	Context("Get Vehicles By Colour Store", func() {
// 		TearDown()

// 		// It("No parking lot available", func() {
// 		// 	cmd := &schema.Command{
// 		// 		Command:   "get_vehicles_by_colour",
// 		// 		Arguments: []string{"Red"},
// 		// 	}
// 		// 	res, err := connection.GetVehiclesByColour().Execute(cmd)
// 		// 	Expect(err).To(Equal(errors.ErrNoParkingLot))
// 		// 	Expect(res).To(Equal(""))
// 		// })
// 		It("No parking lot available", func() {
// 			cmd := &schema.Command{
// 				Command:   "park",
// 				Arguments: []string{"Red"},
// 			}
// 			res, err := connection.Park().Execute(cmd)

// 			Expect(err).To(Equal(errors.ErrNoParkingLot))
// 			Expect(res).To(Equal(""))
// 		})

// 		It("No vehicles found with the given colour", func() {
// 			cmd := &schema.Command{
// 				Command:   "get_vehicles_by_colour",
// 				Arguments: []string{"Red"},
// 			}
// 			createCmd := &schema.Command{
// 				Command:   "create_parking_lot",
// 				Arguments: []string{"2"},
// 			}
// 			connection.CreateParkingLot().Execute(createCmd)

// 			res, err := connection.GetVehiclesByColour().Execute(cmd)
// 			Expect(err).To(Equal(errors.ErrCarNotFound))
// 			Expect(res).To(Equal(""))
// 		})

// 		It("Get vehicles by colour", func() {
// 			cmd := &schema.Command{
// 				Command:   "get_vehicles_by_colour",
// 				Arguments: []string{"Red"},
// 			}
// 			createCmd := &schema.Command{
// 				Command:   "create_parking_lot",
// 				Arguments: []string{"2"},
// 			}
// 			parkCmd := &schema.Command{
// 				Command:   "park",
// 				Arguments: []string{"TN-24-AJ-8462", "Red"},
// 			}
// 			connection.CreateParkingLot().Execute(createCmd)
// 			connection.Park().Execute(parkCmd)

// 			res, err := connection.GetVehiclesByColour().Execute(cmd)
// 			Expect(err).NotTo(HaveOccurred())
// 			Expect(res).To(Equal("TN-24-AJ-8462"))
// 		})

// 		It("Get vehicles by colour - Multiple vehicles", func() {
// 			cmd := &schema.Command{
// 				Command:   "get_vehicles_by_colour",
// 				Arguments: []string{"Red"},
// 			}
// 			createCmd := &schema.Command{
// 				Command:   "create_parking_lot",
// 				Arguments: []string{"3"},
// 			}
// 			parkCmd1 := &schema.Command{
// 				Command:   "park",
// 				Arguments: []string{"TN-24-AJ-8462", "Red"},
// 			}
// 			parkCmd2 := &schema.Command{
// 				Command:   "park",
// 				Arguments: []string{"KA-03-AB-1234", "Red"},
// 			}
// 			connection.CreateParkingLot().Execute(createCmd)
// 			connection.Park().Execute(parkCmd1)
// 			connection.Park().Execute(parkCmd2)

// 			res, err := connection.GetVehiclesByColour().Execute(cmd)
// 			Expect(err).NotTo(HaveOccurred())
// 			Expect(res).To(Equal("TN-24-AJ-8462,KA-03-AB-1234"))
// 		})
// 	})
// })

package store

import (
	"fmt"

	"parking_lot/errors"
	"parking_lot/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parking lot store tests", func() {
	var (
		connection Store
	)
	connection = NewStore()
	It("Tear Down Store Data", func() {
		TearDown()
	})

	Context("parking_lot store execute", func() {
		TearDown()
		It("No parking lot available", func() {
			cmd := &schema.Command{
				Command:   "get_vehicles_by_colour",
				Arguments: []string{"red"},
			}
			res, err := connection.Status().Execute(cmd)
			Expect(err).To(Equal(errors.ErrNoParkingLot))
			Expect(res).To(Equal(""))
		})

		It("Create a parking lot with 2 slots", func() {
			cmd := &schema.Command{
				Command:   "create_parking_lot",
				Arguments: []string{"2"},
			}
			res, err := connection.CreateParkingLot().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(fmt.Sprintf(ParkinglotCreatedInfo, 2)))
		})

		It("park a vehicle", func() {
			cmd := &schema.Command{
				Command:   "park",
				Arguments: []string{"TN-24-AJ-8462", "Red"},
			}
			res, err := connection.Park().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal("Allocated slot number: 1"))
		})

		It("Get Status", func() {
			cmd := &schema.Command{
				Command:   "status",
				Arguments: []string{},
			}
			res, err := connection.Status().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(res))
		})
		It("Get Status", func() {
			cmd := &schema.Command{
				Command:   "status",
				Arguments: []string{},
			}
			res, err := connection.Status().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(res))
		})
	})
})
