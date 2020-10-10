package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"

	"github.com/Pallinder/go-randomdata"
	"github.com/tealeg/xlsx/v3"
)

type Person struct {
	CustomerID  string
	ContractID  string
	Username    string
	Gender      string
	Title       string
	Firstname   string
	Lastname    string
	Email       string
	City        string
	State       string
	Address     string
	PostalCode  string
	Birthdate   string
	PhoneNumber string
	ID          string
	Locale      string
	CreditScore string
}

func NewPerson() Person {
	p := Person{}
	p.CustomerID = fmt.Sprintf("%010d", randomdata.Number(1000000000))
	p.ContractID = randomdata.StringNumber(4, "-")

	gender := rand.Intn(2) // 0/1
	if gender == randomdata.Male {
		p.Gender = "male"
	} else {
		p.Gender = "female"
	}
	p.Username = strings.ToLower(randomdata.SillyName())
	p.Title = randomdata.Title(gender)
	p.Firstname = randomdata.FirstName(gender)
	p.Lastname = randomdata.LastName()
	p.Email = randomdata.Email()
	p.City = randomdata.City()
	p.State = randomdata.State(randomdata.Large)
	p.Address = randomdata.Address()
	p.PostalCode = randomdata.PostalCode("SE")
	p.Birthdate = randomdata.FullDate()
	p.PhoneNumber = randomdata.PhoneNumber()
	p.ID = fmt.Sprintf("%d-%d-%d",
		randomdata.Number(101, 999),
		randomdata.Number(01, 99),
		randomdata.Number(100, 9999),
	)
	p.Locale = randomdata.Locale()
	p.CreditScore = fmt.Sprintf("%d", randomdata.Number(101))
	return p
}

func main() {
	ePtr := flag.Int("c", 0, "Amount of people")
	sPtr := flag.String("p", "", "path to xlsx file")
	flag.Parse()

	entries := *ePtr
	savePath := *sPtr

	if entries == 0 || savePath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if entries >= 1048576-1 {
		log.Fatalln("Too many rows, excel only supports 1,048,576")
	}

	if !strings.HasSuffix(savePath, ".xlsx") {
		log.Fatalln("-p must have suffix .xlsx")
	}

	wb := xlsx.NewFile()
	sheet, err := wb.AddSheet("People")
	if err != nil {
		panic(err)
	}

	// first column
	row := sheet.AddRow()
	p := NewPerson()
	v := reflect.ValueOf(p)
	typeOfS := v.Type()
	for j := 0; j < v.NumField(); j++ {
		cell := row.AddCell()
		cell.Value = typeOfS.Field(j).Name
	}

	for i := 0; i < entries; i++ {
		row := sheet.AddRow()
		p := NewPerson()
		v := reflect.ValueOf(p) // Generics :)
		for j := 0; j < v.NumField(); j++ {
			cell := row.AddCell()
			cell.Value = v.Field(j).Interface().(string)
		}
	}

	wb.Save(savePath)
}
