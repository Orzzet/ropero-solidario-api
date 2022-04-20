package main

import (
	"crypto/rand"
	"fmt"
	"github.com/orzzet/ropero-solidario-api/internal/database"
	transportHTTP "github.com/orzzet/ropero-solidario-api/internal/transport/http"
	"math/big"
	"net/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	secret, err := GenerateRandomString(10)

	if err != nil {
		fmt.Println("Error generating secret")
		fmt.Println(err)
		return err
	}

	handler := transportHTTP.NewHandler(db, secret)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8850", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	fmt.Println("Running server on port 8850")
	return nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
// https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
