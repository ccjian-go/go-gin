package models

import (
	"my/database"
	"log"
)

var db = database.GetDB()

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson()(id int64,err error){
	rs,err := db.Exec("INSERT INTO person(first_name,last_name) VALUES (?,?)", p.FirstName, p.LastName)
	if err != nil{
		log.Fatalln(err)
		return
	}
	id,err = rs.LastInsertId()
	return
}

func (p *Person) GetPersons()(persons []Person){
	persons = make([]Person, 0)

	rows, err := db.Query("SELECT id, first_name, last_name FROM person")
	if err != nil {
		log.Fatalln(err)
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id,&person.FirstName,&person.LastName)
		persons = append(persons, person)
	}
	if err := rows.Err();err != nil{
		log.Fatalln(err)
	}
	return
}