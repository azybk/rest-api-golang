package connection

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api-golang/internal/config"

	_ "github.com/lib/pq"
)

func GetDatabase(cnf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable Timezone=%s",
		cnf.Host, cnf.Port, cnf.User, cnf.Pass, cnf.Name, cnf.Tz)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("error connect db: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error ping to db: ", err.Error())
	}

	return db
}
