package conf

import (
	"OnlinePhotoAlbum/models"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MongodbUri    string
	MongodbName string
	JwtRealm      string
	JwtSignMethod string
	JwtKey        string
	Port string
	GinMode string
)

func init() {
	// read .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongodbUri = os.Getenv("MONGODB_URI")
	MongodbName = os.Getenv("MONGODB_DBNAME")
	JwtRealm = os.Getenv("JWT_REALM")
	JwtSignMethod = os.Getenv("JWT_SIGN_METHOD")
	JwtKey = os.Getenv("JWT_KEY")
	Port = os.Getenv("PORT")
	GinMode = os.Getenv("GIN_MODE")

	// init mongodb
	models.MongodbInit(MongodbUri,MongodbName)
}
