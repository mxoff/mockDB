package main

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// https://regex-escape.com/regex-escaper-online.php

func Test_getID(t *testing.T) {
	fdb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer fdb.Close()

	o := opt{db: fdb}

	type args struct {
		id      int
		model   string
		company string
		price   int
	}

	tests := []struct {
		name string
		args args
	}{

		{
			name: "get one",
			args: args{id: 1, model: "mmm", company: "ccc", price: 555},
		},
		{
			name: "get two",
			args: args{id: 2, model: "eee", company: "rrr", price: 444},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := `select \* from Products where id \= \$1`
			rows := sqlmock.NewRows([]string{"id", "model", "company", "price"}).
				AddRow(tt.args.id, tt.args.model, tt.args.model, tt.args.price)

			mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

			data, err := o.getID(1)

			assert.NoError(t, err)
			assert.NotEmpty(t, data)
			assert.Equal(t, tt.args.model, data.model)
		})
	}
}

func Test_add(t *testing.T) {

	fdb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer fdb.Close()

	o := opt{db: fdb}

	type args struct {
		model   string
		company string
		price   int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "add",
			args: args{company: "ccc", model: "mmm", price: 123},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectExec(`INSERT INTO Products \(model, company, price\) values \('\$1', \$2, \$3\)`).
				WithArgs(tt.args.model, tt.args.company, tt.args.price).
				WillReturnResult(sqlmock.NewResult(1, 1))

			c, err := o.add(tt.args.model, tt.args.company, tt.args.price)
			assert.NoError(t, err)
			assert.Equal(t, 1, c)
		})
	}
}

func Test_getAll(t *testing.T) {

	fdb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer fdb.Close()

	o := opt{db: fdb}

	tests := []struct {
		name string
	}{
		{name: "all"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rows := sqlmock.NewRows([]string{"id", "model", "company", "price"}).
				AddRow(0, "one", "zzz", 77).
				AddRow(0, "two", "xxx", 55)

			mock.ExpectQuery(`select \* from Products`).WillReturnRows(rows)

			data, err := o.getAll()
			assert.NoError(t, err)
			//assert.Equal(t, 1, )
			fmt.Println(data)
		})
	}
}
