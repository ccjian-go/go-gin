package apis

import (
	. "my/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	person := Person{
		FirstName:firstName,
		LastName:lastName,
	}

	id, err := person.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("insert person Id {}", id)
	msg := fmt.Sprintf("insert successful %d", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context) {
	person := Person{
	}

	persons := person.GetPersons()

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}