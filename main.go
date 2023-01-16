package main

// 导入模块
import (
	"fmt"
	"gin-study/db"
	"gin-study/global"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Person 自定义Person类
type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) getAll() (persons []Person, err error) {
	rows, err := global.DB.Query("select id,first_name,last_name from tb_person")
	fmt.Println(rows)
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	defer rows.Close()
	return
}

func main() {
	//创建路由引擎
	router := gin.Default()

	global.DB = db.InitDB()
	//查询,返回所有对象和对象个数
	router.GET("/persons", func(context *gin.Context) {
		p := Person{}
		persons, err := p.getAll()
		if err != nil {
			log.Fatalln(err)
		}
		context.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	})
	router.Run(":8080")
}
