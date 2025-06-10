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
		{"John", "23", "2343343", "Google", user.Address{"New York", "NY", "USA", "10001"}},
		{"Alice", "30", "1234567", "Microsoft", user.Address{"Redmond", "WA", "USA", "98052"}},
		{"Bob", "28", "9876543", "Amazon", user.Address{"Seattle", "WA", "USA", "98101"}},
		{"Charlie", "35", "4567890", "Apple", user.Address{"Cupertino", "CA", "USA", "95014"}},
		{"Diana", "29", "3456789", "Facebook", user.Address{"Menlo Park", "CA", "USA", "94025"}},
		{"Eve", "32", "7890123", "Tesla", user.Address{"Palo Alto", "CA", "USA", "94301"}},
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
		Address: user.Address{"Redmond", "WA", "USA", "98052"},
	}); err != nil {
		fmt.Println("Error writing user:", err)
	} else {
		fmt.Println("User written successfully")
	}

}
