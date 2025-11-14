package partyrobot

import(
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	sentence := fmt.Sprintf("Happy birthday %v! You are now %v years old", name, age)
	return sentence
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	newTable := fmt.Sprintf("%03d", table)
	sentence := fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %v. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, newTable, direction, distance, neighbor)
	return sentence
}