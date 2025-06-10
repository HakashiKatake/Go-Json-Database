package main

import (
	"encoding/json"
	"fmt"
)

const Version = "1.0.1"

func main() {
	dir := "./"
	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error creating database:", err)
	}

	employees := []User{
		{"John", "23", "2343343", "Google", Address{"New York", "NY", "USA", "10001"}},
		{"Alice", "30", "1234567", "Microsoft", Address{"Redmond", "WA", "USA", "98052"}},
		{"Bob", "28", "9876543", "Amazon", Address{"Seattle", "WA", "USA", "98101"}},
		{"Charlie", "35", "4567890", "Apple", Address{"Cupertino", "CA", "USA", "95014"}},
		{"Diana", "29", "3456789", "Facebook", Address{"Menlo Park", "CA", "USA", "94025"}},
		{"Eve", "32", "7890123", "Tesla", Address{"Palo Alto", "CA", "USA", "94301"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})

	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error unmarshalling record:", err)
			continue
		}

		allusers = append(allusers, employeeFound)

	}

	fmt.Println("All Users:", allusers)

	if err := db.Delete("users", "John"); err != nil {
		fmt.Println("Error deleting user:", err)
	} else {
		fmt.Println("User deleted successfully")
	}

}
