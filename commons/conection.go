package commons

import (
	"api-go-4softwaredeveloper/models"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConexion() *gorm.DB {

	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migrate() {
	db := GetConexion()
	defer db.Close()

	log.Println("Migrando...")

	db.AutoMigrate(&models.Persona{})
}
