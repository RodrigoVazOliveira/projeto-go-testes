package database

import (
	"log"
        "os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=" + os.Getenv("HOST_DATABASE")  + " user="+ os.Getenv("USER_DATABASE") + 
        " password=" + os.Getenv("PASS_DATABASE")  +  " dbname=" + os.Getenv("NAME_DATABASE")  + " port=" + os.Getenv("PORT_DATABASE") + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
