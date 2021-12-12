package aocg

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type AdventOfCode struct {
	Day  int
	Year int16
}

const year int16 = 2021

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CreateDay(day int) {

	//Get current folder
	path, err := os.Getwd()
	Check(err)

	daysFolderPath := fmt.Sprint(path, "/days")
	_, err = os.Stat(daysFolderPath)
	if os.IsNotExist(err) {
		os.Mkdir(daysFolderPath, 0766)
	}

	dayFolderPath := fmt.Sprint(path, "/days/day", day)
	_, err = os.Stat(dayFolderPath)
	if !os.IsNotExist(err) {
		log.Fatalf("Folder already exist")
	} else {
		os.Mkdir(dayFolderPath, 0766)
	}

	//Generate day golang file
	templatePath := fmt.Sprint(path, "/template.txt")
	dayPath := fmt.Sprint(dayFolderPath, "/day", day, ".go")

	d := AdventOfCode{day, year}
	tmpl, err := template.ParseFiles(templatePath)
	Check(err)

	f, err := os.Create(dayPath)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = tmpl.Execute(f, d)
	Check(err)

}
