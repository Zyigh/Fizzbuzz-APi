package main

import (
	"fizzbuzz-api/conf"
	"fizzbuzz-api/controllers"
	"fizzbuzz-api/entities"
	"fizzbuzz-api/facades"
	"fizzbuzz-api/middlewares"
	"fmt"
	"log"
	"net/http"
	"time"
)

// waitForCouch allows the program to wait for couchdb creation before making sure database exists
// Useful in case of using docker compose to run the app
func waitForCouch(maxAttempts int) error {
	dburl := fmt.Sprintf(
		"http://%s:%s@%s:%d",
		conf.DBUserName,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
	)
	req, err := http.NewRequest("GET", dburl, nil)

	if err != nil {
		return err
	}
	client := &http.Client{}

	for i := 0; i < maxAttempts; i++ {
		log.Println("Attempt", i + 1, "/", maxAttempts, "to connect to database...")
		res, err := client.Do(req)

		if err != nil {
			time.Sleep(time.Second * 5)
			continue
		}

		if res.StatusCode == http.StatusOK {
			return nil
		}

		time.Sleep(time.Second * 5)
	}

	return fmt.Errorf("couldn't connect to dabase in %d attempts", maxAttempts)
}

// main waits for couchdb response before launching server
// creates couchdb database if it doesn't exist
// creates routes with middleware
// starts server
func main() {
	mux := http.NewServeMux()

	if err := waitForCouch(conf.MaxAttempts); err != nil {
		log.Fatal(err)
	}

	if err := (entities.Request{}).CreateDatabaseIfNotExists(); err != nil {
		log.Fatal(err)
	}

	fizzbuzzCtrl := controllers.GetFizzbuzzCtrl()
	mux.Handle(
		"/",
		middlewares.JSONHeaders(
		middlewares.GetParamConverter(
		&facades.GetFizzbuzzFacade{},
		fizzbuzzCtrl.GetFizzbuzz(),
	)))

	statsController := controllers.GetStatsController()
	mux.Handle("/stats/most-frequent-request", middlewares.JSONHeaders(statsController.GetMostAskedRequest()))

	appAddr := fmt.Sprintf("%s:%d", conf.AppHost, conf.AppHTTPSPort)
	if err := http.ListenAndServeTLS(appAddr, conf.AppHTTPSCertFile, conf.AppHTTPSKeyFile, mux); err != nil {
		log.Fatal(err)
	}
}
