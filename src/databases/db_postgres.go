package databases

import (
	conf "app/src/configs"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DBPOSTGRES() *gorm.DB {
	db_data := conf.VarConfig().DB["DB_POSTGRES"]
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		db_data.User, db_data.Password, db_data.Host, db_data.Port, db_data.Database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Error conectando a la BD POSTGRES:", err)
	}

	return db
}
