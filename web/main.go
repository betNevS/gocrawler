package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/betNevS/gocrawler/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ShowImage(images string) template.HTML {
	var imgs []string
	json.Unmarshal([]byte(images), &imgs)
	res := ``
	for _, m := range imgs {
		res = res + `<img height="200" width="200" src="` + m + `" alt="">`
	}
	return template.HTML(res)
}

func main() {
	Db, _ := gorm.Open("mysql", "root:666666@tcp(127.0.0.1:3306)/crawler?charset=utf8mb4&parseTime=True&loc=Local")
	Db.SingularTable(true)
	//Db.LogMode(true)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// Funcs(template.FuncMap{"ShowImage": ShowImage}
		t := template.New("index.html").Funcs(template.FuncMap{"ShowImage": ShowImage})
		temp, err := t.ParseFiles("index.html")
		if err != nil {
			fmt.Println(err)
		}
		// temp.Funcs(template.FuncMap{"ShowImage": ShowImage})
		// 从数据库里面取数据
		var houses []models.HouseItem
		Db.Limit(10).Find(&houses)
		//fmt.Println(houses)
		err = temp.Execute(writer, houses)
		fmt.Println(err)
	})
	http.ListenAndServe(":8080", nil)
}
