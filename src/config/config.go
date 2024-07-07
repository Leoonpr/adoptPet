package config
import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	StringConexaoBanco = os.Getenv("DB_USER") + ":" + os.Getenv("DB_SENHA") + "@/" + os.Getenv("DB_NOME") + "?charset=utf8&parseTime=True&loc=Local"
}
