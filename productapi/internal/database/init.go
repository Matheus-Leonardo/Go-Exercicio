package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB tenta abrir a conexão com retry robusto
func InitDB(dsn string, retries int, interval time.Duration) *gorm.DB {
	var db *gorm.DB
	var err error

	for i := 0; i < retries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Erro ao abrir conexão (tentativa %d/%d): %v", i+1, retries, err)
			time.Sleep(interval)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Erro ao obter sql.DB (tentativa %d/%d): %v", i+1, retries, err)
			time.Sleep(interval)
			continue
		}

		if err = sqlDB.Ping(); err != nil {
			log.Printf("Banco ainda não pronto (tentativa %d/%d): %v", i+1, retries, err)
			time.Sleep(interval)
			continue
		}

		log.Println("Conexão com o banco estabelecida!")
		return db
	}

	log.Fatalf("Não foi possível conectar ao banco após %d tentativas: %v", retries, err)
	return nil
}

// OpenSimpleDB abre conexão sem retry
func OpenSimpleDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao abrir conexão simples: %v", err)
	}
	log.Println("Conexão simples com o banco estabelecida!")
	return db
}

// Close encerra a conexão com o banco
func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
	}
}
