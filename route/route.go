package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	. "go-gin/apis"
	"go-gin/database"
	_ "go-gin/logger"
	. "go-gin/models"
	"log"
	"net/http"
	"strconv"
)

var router = gin.Default()
var db = database.GetDB()

func InitRouter()  *gin.Engine {
	test()

	// 禁用控制台颜色
	gin.DisableConsoleColor()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works!")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 写入单个 POST http://localhost:8000/person?first_name=4&last_name=4
	router.POST("/person", AddPersonApi)

	// 获取列表 GET http://localhost:8000/persons
	router.GET("/persons", GetPersonsApi)

	// 获取单个 GET http://localhost:8000/person/1
	router.GET("/person/:id", func(c *gin.Context) {
		id := c.Param("id")
		var person Person
		err := db.QueryRow("SELECT id, first_name, last_name FROM person WHERE id=?", id).Scan(
			&person.Id, &person.FirstName, &person.LastName,
		)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"person": nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"person": person,
		})
	})

	// 修改单个 PUT http://localhost:8000/persons/6&first_name=666&last_name=666
	router.PUT("/persons/:id", func(c *gin.Context) {
		cid := c.Param("id")
		id, err := strconv.Atoi(cid)
		person := Person{Id: id}
		err = c.Bind(&person)
		if err != nil {
			log.Fatalln(err)
		}

		stmt, err := db.Prepare("UPDATE person SET first_name=?, last_name=? WHERE id=?")

		if err != nil {
			log.Fatalln(err)
		}
		defer stmt.Close()
		rs, err := stmt.Exec(person.FirstName, person.LastName, person.Id)
		if err != nil {
			log.Fatalln(err)
		}
		ra, err := rs.RowsAffected()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("Update person %d successful %d", person.Id, ra)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	// 删除单个 DELETE http://localhost:8000/person/6
	router.DELETE("/person/:id", func(c *gin.Context) {
		cid := c.Param("id")
		id, err := strconv.Atoi(cid)
		if err != nil {
			log.Fatalln(err)
		}
		rs, err := db.Exec("DELETE FROM person WHERE id=?", id)
		if err != nil {
			log.Fatalln(err)
		}
		ra, err := rs.RowsAffected()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("Delete person %d successful %d", id, ra)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	return router
}