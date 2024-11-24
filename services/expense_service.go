package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/giuszeppe/expense-tracker-go-cli/single"
)

func ListExpenses() {
	db := single.GetInstance()

	out, err := db.Database.Query("Select id,amount,description,collected_date from expenses")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# %-5s %-10s %-20s %-10s\n", "ID", "AMOUNT", "DESCRIPTION", "DATE")

	for out.Next() {
		var id, amount, description, date string
		out.Scan(&id, &amount, &description, &date)
		fmt.Printf("# %-5s %-10s %-20s %-10s\n", id, amount, description, date)
	}
}

func AddExpense(amount float64, description string) {
	db := single.GetInstance()

	stmt, err := db.Database.Prepare("INSERT INTO EXPENSES (amount, description) VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	out, err := stmt.Exec(amount, description)
	if err != nil {
		log.Fatal(err)
	}
	id, err := out.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("# Expenses correctly added with ID:", id)
}
func RemoveExpense(id int) {
	db := single.GetInstance()

	stmt, err := db.Database.Prepare("DELETE FROM expenses WHERE id=? ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

func Summary(month int) {
	db := single.GetInstance()
	var out *sql.Rows

	if month == 0 {
		query := "SELECT SUM(amount) FROM expenses"
		stmt, err := db.Database.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		out, err = stmt.Query()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		stmt, err := db.Database.Prepare("SELECT SUM(amount) FROM expenses WHERE substr(collected_date,3,2) = ?")
		if err != nil {
			log.Fatal(err)
		}
		out, err = stmt.Query(month)
		if err != nil {
			log.Fatal(err)
		}
	}

	for out.Next() {
		var amount float64
		out.Scan(&amount)
		if month == 0 {
			fmt.Println("# Amount is ", amount)
		} else {
			fmt.Printf("# Amount for month %v is %v", month, amount)
		}
	}
}
