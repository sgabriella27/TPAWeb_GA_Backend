package database

import (
	"github.com/sgabriella27/TPAWebGA_Back/graph/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var inst *gorm.DB

func init() {
	dsn := "host=localhost user=postgres password= dbname=tpaweb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	inst = db

	migration()
	//seed()
}

func GetDatabase() *gorm.DB {
	return inst
}

func migration() {
	//if err := inst.Migrator().DropTable(&model.User{}, &model.Game{}, &model.GameMedia{}, &model.GameSlideshow{}); err != nil {
	//	log.Fatal(err)
	//}

	if err := inst.AutoMigrate(&model.Promo{}, &model.User{}, &model.Game{}, &model.GameMedia{}, &model.GameSlideshow{}); err != nil {
		log.Fatal(err)
	}
}

func seed() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	inst.Create(&model.User{
		AccountName: "admin",
		Password:    string(hash),
	})
}
