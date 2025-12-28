package initapp

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitApp() error {
	godotenv.Load()
	test := os.Getenv("TEST")
	is_test := test == "true"
	err := initTest(is_test)
	if err != nil {
		fmt.Println("error in usecase/initprocess/init_app.go:/InitApp/initTest")
		return err
	}
	err = initDir()
	if err != nil {
		fmt.Println("error in usecase/initprocess/init_app.go:/InitApp/initDir")
		return err
	}
	return err
}
