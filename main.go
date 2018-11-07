package main

import (
	"bufio"
	tk "github.com/eaciit/toolkit"
	"os"
	"strconv"
	"strings"
)

func main() {
	cars := []tk.M{}
	totalSlots := 0
	slotno := 0
	isLeaved := false
	slotsFree := ""
	for {
		reader := bufio.NewReader(os.Stdin)
		tk.Printf("$")
		text, _ := reader.ReadString('\n')
		reply := ""
		input := strings.Split(text, " ")
		shellCommand := strings.TrimSuffix(input[0], "\r\n")

		createSlots := "create_parking_lot"
		park := "park"
		status := "status"
		leave := "leave"
		findSlotsbyReg := "slot_number_for_registration_number"
		findRegbyColour := "registration_numbers_for_cars_with_colour"
		findSlotsbyColour := "slot_numbers_for_cars_with_colour"

		wrongMessage := "Command not completed"

		switch shellCommand {
		case createSlots:
			if len(input) == 2 {
				slotsValue := strings.TrimSuffix(input[1], "\r\n")

				jumlah, _ := strconv.Atoi(slotsValue)
				totalSlots = jumlah
				reply = "Created a parking lot with " + slotsValue + " slots"

			} else {
				reply = wrongMessage
			}
		case park:
			if len(input) == 3 {
				if totalSlots == 0 {
					reply = "Slots Not Yet Defined"
				} else if len(cars) < totalSlots {
					detail := tk.M{}
					if len(cars) > 0 && isLeaved == true {
						idx := 0
						sn := []string{}
						for _, cs := range cars {
							sn = append(sn, cs.GetString("slotno"))
						}
						for i := 0; i < totalSlots; i++ {
							idx = i + 1
							slotsFree = strconv.Itoa(idx)
							if !tk.HasMember(sn, slotsFree) {
								detail.Set("slotno", slotsFree)
								reply = "Allocated slot number: " + slotsFree
							}
						}
					} else {
						slotno = slotno + 1
						detail.Set("slotno", strconv.Itoa(slotno))
						reply = "Allocated slot number: " + strconv.Itoa(slotno)
					}

					regno := strings.TrimSuffix(input[1], "\r\n")
					colour := strings.TrimSuffix(input[2], "\r\n")

					detail.Set("regno", regno)
					detail.Set("colour", colour)

					cars = append(cars, detail)

				} else {
					reply = "Sorry, parking lot is full"
				}
			} else {
				reply = wrongMessage
			}
		case status:
			tk.Println("Slot No.", " ", "Registration No", " ", "Colour")
			for _, ca := range cars {
				tk.Println(ca.GetString("slotno"), "	  ", ca.GetString("regno"), "	", ca.GetString("colour"))
			}
		case findSlotsbyReg:
			if len(input) == 2 {
				regNo := strings.TrimSuffix(input[1], "\r\n")
				if totalSlots == 0 {
					reply = "Slots Not Yet Defined"
				} else {
					found := false
					for _, n := range cars {
						if n.GetString("regno") == regNo {
							found = true
							tk.Println(n.GetString("slotno"))
						}
					}

					if found == false {
						tk.Println("Not found")
					}
				}
			} else {
				reply = wrongMessage
			}
		case findRegbyColour:
			if len(input) == 2 {
				colour := strings.TrimSuffix(input[1], "\r\n")

				if totalSlots == 0 {
					reply = "Slots Not Yet Defined"
				} else {
					found := false
					result := ""
					for _, n := range cars {
						if n.GetString("colour") == colour {
							found = true
							result += n.GetString("regno") + ", "

						}
					}
					tk.Println(result)

					if found == false {
						tk.Println("Not found")
					}
				}
			} else {
				reply = wrongMessage
			}
		case findSlotsbyColour:
			if len(input) == 2 {
				colour := strings.TrimSuffix(input[1], "\r\n")

				if totalSlots == 0 {
					reply = "Slots Not Yet Defined"
				} else {
					found := false
					result := ""
					for _, n := range cars {
						if n.GetString("colour") == colour {
							found = true
							result += n.GetString("slotno") + ", "

						}
					}
					tk.Println(result)

					if found == false {
						tk.Println("Not found")
					}
				}
			} else {
				reply = wrongMessage
			}
		case leave:
			if len(input) == 2 {
				isLeaved = true
				slotNoStrings := strings.TrimSuffix(input[1], "\r\n")
				slotNo, _ := strconv.Atoi(slotNoStrings)
				idx := slotNo - 1

				cars = append(cars[:idx], cars[idx+1:]...)

				tk.Println("Slot number", slotNoStrings, "is free")
			} else {
				reply = wrongMessage
			}
		default:
			tk.Println(wrongMessage)
		}

		tk.Println(reply)
	}
}
