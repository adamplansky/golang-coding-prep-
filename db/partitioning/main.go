package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	randStr "showmax/rand"
	"time"
)

type Product struct {
	gorm.Model //ID, CreatedAt, UpdatedAt, DeletedAt
	Code       string
	Mw         uint
}

func getRandomProduct() Product {
	randomYear := rand.Int31n(50) + 2000
	//randomYear := 2010
	randoMonth := rand.Int31n(12) + 1
	randomDay := rand.Int31n(28) + 1
	randomTime := time.Date(int(randomYear), time.Month(randoMonth), int(randomDay), 0, 0, 0, 0, time.UTC)
	randomMw := uint(rand.Int31n(10000) + 1)
	randomCode := randStr.String(10)

	return Product{
		Code: randomCode,
		Mw:   randomMw,
		Model: gorm.Model{
			CreatedAt: randomTime,
			UpdatedAt: randomTime,
		},
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var err error
	dsn := "user=postgres password=postgres dbname=partitioning host=localhost port=15432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to open connection: %v", err)
	}
	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("unable to automigrate product: %v", err)
	}

	var product Product
	for i := 0; i < 10000; i++ {
		product = getRandomProduct()
		err = db.Create(&product).Error
		if err != nil {
			log.Fatalf("error inserting product to db %v", err)
		}
		//log.Printf("product inserted to. product: %+v\n", product)
		if i % 1000 == 0 {
			log.Print(i)
		}

	}

	//var products []Product
	//err = db.Find(&products).Error
	//if err != nil {
	//	log.Fatalf("unable to get all products %v", err)
	//}
	//for i, product := range products {
	//	log.Printf("%d: %+v\n", i, product)
	//}

}
