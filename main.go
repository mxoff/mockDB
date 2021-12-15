package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type opt struct {
	db *sql.DB
}

func main() {
	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	o := opt{
		db: db,
	}

	_, err = o.add("Apple", "iphone", 70000)
	dataID, err := o.getID(1)
	dataALL, err := o.getAll()

	fmt.Println(err, dataID)
	fmt.Println(err, dataALL)

}

//
//
//
func (o opt) add(model, comp string, price int) (int, error) {
	result, err := o.db.Exec("INSERT INTO Products (model, company, price) values ('$1', $2, $3)",
		model, comp, price)
	if err != nil {
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	// количество добавленных строк
	//fmt.Println("insert = ", count)
	return int(count), nil
}

type product struct {
	id      int
	model   string
	company string
	price   int
}

//
//
func (o opt) getAll() ([]product, error) {
	rows, err := o.db.Query("select * from Products")
	defer rows.Close()

	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	//  for _, p := range products {
	//  	fmt.Println(p.id, p.model, p.company, p.price)
	//  }

	return products, err
}

func (o opt) getID(id int) (product, error) {

	row := o.db.QueryRow("select * from Products where id = $1", id)
	prod := product{}
	err := row.Scan(&prod.id, &prod.model, &prod.company, &prod.price)

	return prod, err
}
