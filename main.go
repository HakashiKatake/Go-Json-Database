package main

import (
	"encoding/json"
	"fmt"

	"github.com/HakashiKatake/Go-Json-Database/db"
	user "github.com/HakashiKatake/Go-Json-Database/types"
)

const Version = "1.0.1"

func main() {
	dir := "./"
	database, err := db.New(dir, nil)
	if err != nil {
		fmt.Println("Error creating database:", err)
	}

	employees := []user.User{
		{Name: "John", Age: "23", Contact: "2343343", Company: "Google", Address: user.Address{City: "New York", State: "NY", Country: "USA", Pincode: "10001"}},
		{Name: "Alice", Age: "30", Contact: "1234567", Company: "Microsoft", Address: user.Address{City: "Redmond", State: "WA", Country: "USA", Pincode: "98052"}},
		{Name: "Bob", Age: "28", Contact: "9876543", Company: "Amazon", Address: user.Address{City: "Seattle", State: "WA", Country: "USA", Pincode: "98101"}},
		{Name: "Charlie", Age: "35", Contact: "4567890", Company: "Apple", Address: user.Address{City: "Cupertino", State: "CA", Country: "USA", Pincode: "95014"}},
		{Name: "Diana", Age: "29", Contact: "3456789", Company: "Facebook", Address: user.Address{City: "Menlo Park", State: "CA", Country: "USA", Pincode: "94025"}},
		{Name: "Eve", Age: "32", Contact: "7890123", Company: "Tesla", Address: user.Address{City: "Palo Alto", State: "CA", Country: "USA", Pincode: "94301"}},
	}

	for _, value := range employees {
		database.Write("users", value.Name, user.User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})

	}

	records, err := database.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(records)

	allusers := []user.User{}

	for _, f := range records {
		employeeFound := user.User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error unmarshalling record:", err)
			continue
		}

		allusers = append(allusers, employeeFound)

	}

	fmt.Println("All Users:", allusers)

	if err := database.Delete("users", "Eve"); err != nil {
		fmt.Println("Error deleting user:", err)
	} else {
		fmt.Println("User deleted successfully")
	}

	if err := database.Read("users", "Alice", &user.User{}); err != nil {
		fmt.Println("Error reading user:", err)
	} else {
		fmt.Println("User read successfully")
	}

	if err := database.Write("users", "Alice", user.User{
		Name:    "Alice",
		Age:     "31",
		Contact: "1234567",
		Company: "Microsoft",
		Address: user.Address{City: "Redmond", State: "WA", Country: "USA", Pincode: "98052"},
	}); err != nil {
		fmt.Println("Error writing user:", err)
	} else {
		fmt.Println("User written successfully")
	}

}
