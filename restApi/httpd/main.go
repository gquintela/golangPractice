package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Cd struct {
	id_cd     int
	cd_name   string
	band_name string
	year      int
}

func main() {
	r := gin.Default()
	db, err := sql.Open("mysql", "goRestApi:goRestApi2020.@tcp(127.0.0.1:3306)/cds")
	if err != nil {
		fmt.Println("Error while connecting to DB.")
		log.Fatal(err)
	}

	//get cd
	r.GET("/cd/:id_cd", func(c *gin.Context) {
		id := c.Param("id_cd")
		cd_query := "SELECT * FROM cds WHERE id_cd = " + id + ";"
		rows, err := db.Query(cd_query)
		if err != nil {
			println(err.Error())
			log.Fatal(err)
		}
		answer := Cd{}

		for rows.Next() {
			err := rows.Scan(&answer.id_cd, &answer.cd_name, &answer.band_name, &answer.year)
			if err != nil {
				log.Fatal(err)
			}
		}

		code:=200
		if(answer.id_cd == 0){
			code = 404
		}

		c.JSON(code, gin.H{
			"id_cd":     answer.id_cd,
			"cd_name":   answer.cd_name,
			"band_name": answer.band_name,
			"year":      answer.year,
		})
	})

	//get all cds
	r.GET("/cds", func(c *gin.Context) {
		cd_query := "SELECT * FROM cds;"
		rows, err := db.Query(cd_query)
		if err != nil {
			println(err.Error())
			log.Fatal(err)
		}
		var answer []Cd

		for rows.Next() {
			cd := Cd{}
			err := rows.Scan(&cd.id_cd, &cd.cd_name, &cd.band_name, &cd.year)
			answer = append(answer, cd)
			if err != nil {
				log.Fatal(err)
			}
		}

		for _, element := range answer {
			println(element.band_name, element.cd_name, element.id_cd, element.year)
		}

		c.JSON(200, gin.H{
			//todo: parse rows into Json
		})
	})

	//delete cd
	r.GET("/delete/:id_cd", func(c *gin.Context) {
		id := c.Param("id_cd")
		cd_query := "DELETE FROM cds WHERE id_cd = " + id + ";"
		_, err := db.Query(cd_query)
		if err != nil {
			println(err.Error())
			log.Fatal(err)
		}

		c.JSON(200, gin.H{
			"id_cd":     id,
			"message":   "deleted",
		})
	})

	r.Run()
	println("hello")
}
