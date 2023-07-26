package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func createUser1() *peopleAdrresses {
	Test1 := new(peopleAdrresses)
	Test1.phoneNumber = ""
	Test1.firstName = "Courtney"
	Test1.lastName = "Cox"
	Test1.streetNo = "1289"
	Test1.streetName = "Birmingham"
	Test1.zipCode = "1214"
	Test1.city = "Alabama"
	Test1.state = "USA"
	return Test1
}
func createUser2() *peopleAdrresses {
	Test2 := new(peopleAdrresses)
	Test2.phoneNumber = "795-262-13797"
	Test2.firstName = ""
	Test2.lastName = "Cox"
	Test2.streetNo = ""
	Test2.streetName = "Birmingham"
	Test2.zipCode = "1214"
	Test2.city = "Alabama"
	Test2.state = "USA"
	return Test2
}
func createUser3() *peopleAdrresses {
	Test3 := new(peopleAdrresses)
	Test3.phoneNumber = "795-262-13797"
	Test3.firstName = "Courtney"
	Test3.lastName = ""
	Test3.streetNo = "1289"
	Test3.streetName = "Birmingham"
	Test3.zipCode = "1214"
	Test3.city = "Alabama"
	Test3.state = "USA"
	return Test3
}
func createUser4() *peopleAdrresses {
	Test4 := new(peopleAdrresses)
	Test4.phoneNumber = "795-262-13797"
	Test4.firstName = "Courtney"
	Test4.lastName = "Cox"
	Test4.streetNo = "1279"
	Test4.streetName = ""
	Test4.zipCode = "1214"
	Test4.city = "Alabama"
	Test4.state = "USA"
	return Test4
}
func createUser5() *peopleAdrresses {
	Test5 := new(peopleAdrresses)
	Test5.phoneNumber = "795-262-13797"
	Test5.firstName = "Courtney"
	Test5.lastName = "Cox"
	Test5.streetNo = "1289"
	Test5.streetName = "Birmingham"
	Test5.zipCode = ""
	Test5.city = "Alabama"
	Test5.state = "USA"
	return Test5
}
func createUser6() *peopleAdrresses {
	Test6 := new(peopleAdrresses)
	Test6.phoneNumber = "795-262-13797"
	Test6.firstName = "Courtney"
	Test6.lastName = "Cox"
	Test6.streetNo = "1289"
	Test6.streetName = "Birmingham"
	Test6.zipCode = "1214"
	Test6.city = ""
	Test6.state = "USA"
	return Test6
}
func createUser7() *peopleAdrresses {
	Test7 := new(peopleAdrresses)
	Test7.phoneNumber = "795-262-13797"
	Test7.firstName = "Courtney"
	Test7.lastName = "Cox"
	Test7.streetNo = "1289"
	Test7.streetName = "Birmingham"
	Test7.zipCode = "1214"
	Test7.city = "Alabama"
	Test7.state = ""
	return Test7
}
func createUser8() *peopleAdrresses {
	Test8 := new(peopleAdrresses)
	Test8.phoneNumber = "795-262-13797"
	Test8.firstName = "Courtney"
	Test8.lastName = "Cox"
	Test8.streetNo = "1289"
	Test8.streetName = " Birmingham"
	Test8.zipCode = "1214"
	Test8.city = "Alabama"
	Test8.state = "USA"
	return Test8
}

func retrieveByName() ([]peopleAdrresses, error) {
	db, err := sql.Open("mysql", "Sarah:S04770371~@tcp(127.0.0.1:3306)/addressbook")
	if err != nil {
		panic(err.Error())
	}

	//	rows, err := db.Query("select phoneNumber, firstName from peopleAdrresses where firstName=?", firstName)
	rows, err := db.Query("SELECT * FROM peopleaddresses WHERE firstName = 'Jennifer'")

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

func TestInsertDB1(t *testing.T) { //testing.T is used to report the success or failure

	insertResult, err := InsertDB(*createUser1())
	if err != nil {
		fmt.Print("Error during the data insertion into the database \n")
	} else {
		t.Log("Done with the database insertion :) \n", insertResult)
	}
}

func TestInsertDB2(t *testing.T) { //testing.T is used to report the success or failure

	insertResult, err := InsertDB(*createUser2())
	if err != nil {
		fmt.Print("Error during the data insertion into the database \n")
	} else {
		t.Log("Done with the database insertion :) \n", insertResult)
	}
}

func TestInsertDB3(t *testing.T) { //testing.T is used to report the success or failure

	insertResult, err := InsertDB(*createUser3())
	if err != nil {
		fmt.Print("Error during the data insertion into the database \n")
	} else {
		t.Log("Done with the database insertion :) \n", insertResult)
	}
}
func TestInsertDB4(t *testing.T) { //testing.T is used to report the success or failure
	insertResult, err := InsertDB(*createUser4())
	if err != nil {
		fmt.Print("Error during the data insertion into the database \n")
	} else {
		t.Log("Done with the database insertion :) \n", insertResult)
	}
}
func TestInsertDB5(t *testing.T) { //testing.T is used to report the success or failure
	insertResult, err := InsertDB(*createUser5())
	if err != nil {
		fmt.Print("Error during the data insertion into the database \n")
	} else {
		t.Log("Done with the database insertion :) \n", insertResult)
	}
}
func TestInsertDB6(t *testing.T) { //testing.T is used to report the success or failure
	insertResult, err := InsertDB(*createUser6())
	if err != nil {
		fmt.Print("Error during the data insertion into the database \n")
	} else {
		t.Log("Done with the database insertion :) \n", insertResult)
	}
}
func TestInsertDB7(t *testing.T) { //testing.T is used to report the success or failure
	insertResult, err := InsertDB(*createUser7())
	if err != nil {
		fmt.Print("Error during the data insertion into the database \n")
	} else {
		t.Log("Done with the database insertion :) \n", insertResult)
	}
}
func TestInsertDB8(t *testing.T) { //testing.T is used to report the success or failure
	insertResult, err := InsertDB(*createUser8())
	if err != nil {
		fmt.Print("Error during the data insertion into the database \n")
	} else {
		t.Log("Done with the database insertion :) \n", insertResult)
	}
}

func TestRecordByName(t *testing.T) {
	var retrieve1 []peopleAdrresses
	var retreive2 []peopleAdrresses
	retrieve1, err1 := recordByName("Jennifer")

	if err1 != nil {
		print(err1.Error())
	}
	retreive2, err2 := retrieveByName()

	if err2 != nil {
		print(err2.Error())
	}
	if len(retrieve1) != len(retreive2) {
		t.Error("Unmached size \n")
	}
	for i, rec := range retrieve1 {
		if rec != retreive2[i] {
			t.Error("Records are not matched \n")
		} else {
			t.Log("Matched recoreds \n")
		}

	}

}
