package initiator

import (
	"log"

	"github.com/JamesMorales04/personalPortfolioServer.git/internal/database"
)

func DatabaseInit() {
	if !database.CheckConnection() {
		log.Fatal("No database conection")
		return
	}
}
