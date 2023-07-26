package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type peopleAdrresses struct {
	phoneNumber string
	firstName   string
	lastName    string
	streetNo    string
	streetName  string
	zipCode     string
	city        string
	state       string
}

func InsertDB(userDetails peopleAdrresses) (bool, error) {
	if userDetails.phoneNumber == "" {
		fmt.Print("Empty Phone Number \n")
		os.Exit(1)
	}

	if userDetails.firstName == "" {
		fmt.Print("Empty First Name \n")
		os.Exit(1)
	}
	if userDetails.lastName == "" {
		fmt.Print("Empty Last Name \n")
		os.Exit(1)
	}
	if userDetails.streetNo == "" {
		fmt.Print("Empty Street Number \n")
		os.Exit(1)
	}
	if userDetails.streetName == "" {
		fmt.Print("Empty Street Name \n")
		os.Exit(1)
	}
	if userDetails.zipCode == "" {
		fmt.Print("Empty Zip Code \n")
		os.Exit(1)
	}
	if userDetails.city == "" {
		fmt.Print("Empty City Name \n")
		os.Exit(1)
	}
	if userDetails.state == "" {
		fmt.Print("Empty State  \n")
		os.Exit(1)
	}
	db, err := sql.Open("mysql", "Sarah:S04770371~@tcp(127.0.0.1:3306)/addressbook")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Printf("Success \n")
	stmt, err := db.Prepare("Insert into peopleaddresses set phoneNumber=?, firstName=?, lastName=?, streetNo=?, streetName=?, zipCode=?, city=?, state=? ")
	if err != nil {
		fmt.Print("helper_methods.go : 118")
		fmt.Println(err)
		return false, err
	}
	_, queryError := stmt.Exec(userDetails.phoneNumber, userDetails.firstName, userDetails.lastName, userDetails.streetNo, userDetails.streetName, userDetails.zipCode, userDetails.city, userDetails.state)

	if queryError != nil {
		fmt.Print("helper_methods.go : 125")
		fmt.Println(queryError)
		return false, err
	}

	return true, err
}

func recordByName(firstName string) ([]peopleAdrresses, error) {
	db, err := sql.Open("mysql", "Sarah:S04770371~@tcp(127.0.0.1:3306)/addressbook")
	if err != nil {
		panic(err.Error())
	}

	//	rows, err := db.Query("select phoneNumber, firstName from peopleAdrresses where firstName=?", firstName)
	rows, err := db.Query("SELECT * FROM peopleaddresses WHERE firstName = ?", firstName)

	if err != nil {
		print(err.Error())
		return []peopleAdrresses{}, err
	}
	defer rows.Close() // to release any resourses taken during the function call

	// A record slice to hold data from returned rows.
	var records []peopleAdrresses

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() { //iterates over the returned rows
		var record peopleAdrresses
		if err := rows.Scan(&record.phoneNumber, &record.firstName, &record.lastName, &record.streetNo, &record.streetName, &record.zipCode, &record.city, &record.state); err != nil {
			print(err.Error())
			return records, err
		}

		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		return records, err
	}

	return records, nil
}
func listReturn() ([]peopleAdrresses, error) {
	db, err := sql.Open("mysql", "Sarah:S04770371~@tcp(127.0.0.1:3306)/addressbook")
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM peopleaddresses ")

	if err != nil {
		print(err.Error())
		return []peopleAdrresses{}, err
	}
	defer rows.Close()

	// A record slice to hold data from returned rows.
	var records []peopleAdrresses

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() { //iterates over the returned rows
		var record peopleAdrresses
		if err := rows.Scan(&record.phoneNumber, &record.firstName, &record.lastName, &record.streetNo, &record.streetName, &record.zipCode, &record.city, &record.state); err != nil {
			print(err.Error())
			return records, err
		}

		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		return records, err
	}

	return records, nil
}
func UserInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return line
}
func main() {

	user2 := peopleAdrresses{}
	fmt.Print("Enter your Phone Number \n")
	fmt.Scanln(&user2.phoneNumber)
	fmt.Print("Enter your First Name \n")
	fmt.Scanln(&user2.firstName)
	fmt.Print("Enter your Last Name \n")
	fmt.Scanln(&user2.lastName)
	fmt.Print("Enter your Street Number \n")
	fmt.Scanln(&user2.streetNo)
	fmt.Print("Enter your Street Name \n")
	fmt.Scanln(&user2.streetName)
	fmt.Print("Enter your Zip Code \n")
	fmt.Scanln(&user2.zipCode)
	fmt.Print("Enter your City Name \n")
	fmt.Scanln(&user2.city)
	fmt.Print("Enter your State \n")
	fmt.Scanln(&user2.state)
	InsertDB(user2)

	fmt.Print("Enter any name you want :) \n")
	user2.firstName = UserInput()

	records, err := recordByName(user2.firstName)
	if err != nil {
		print("Error %s when selecting record ", err)
	}
	for _, record := range records {
		print("Users detaials \n", record.phoneNumber, record.firstName, record.lastName, record.streetNo, record.streetName, record.zipCode, record.city, record.state)
	}
	records2, err := listReturn()
	if err != nil {
		print("Error %s when selecting record ", err)
	}
	for _, record := range records2 {
		print("The database as a list \n", record.phoneNumber, record.firstName, record.lastName, record.streetNo, record.streetName, record.zipCode, record.city, record.state)
	}

}
