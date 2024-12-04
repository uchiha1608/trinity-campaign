package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"trinity-campaign/apis"
	"trinity-campaign/database"
	"trinity-campaign/utils"
)

func main() {
	//Init config
	config, err := utils.LoadConfig("")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	utils.InitConfig(config)

	//Init database connection
	database.InitDB()

	router := mux.NewRouter()

	//campaign route
	router.HandleFunc("/campaign/{id}", apis.GetCampaignStatus).Methods("GET")
	router.HandleFunc("/campaign/", apis.CreateCampaign).Methods("POST")

	//voucher route
	router.HandleFunc("/vouchers/generate", apis.GenerateVoucher).Methods("POST")
	router.HandleFunc("/vouchers/validate", apis.ValidateVoucher).Methods("POST")

	//purchase route
	router.HandleFunc("/purchases", apis.ApplyVoucherToSubscription).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   utils.AppConfig.HttpsConfig.AllowedOrigins,   // You can specify allowed origins here
		AllowedMethods:   utils.AppConfig.HttpsConfig.AllowedMethods,   // Allowed HTTP methods
		AllowedHeaders:   utils.AppConfig.HttpsConfig.AllowedHeaders,   // Allowed headers
		AllowCredentials: utils.AppConfig.HttpsConfig.AllowCredentials, // Allow credentials such as cookies
		MaxAge:           utils.AppConfig.HttpsConfig.MaxAge,           // Cache preflight response for 12 hours
	})

	handler := c.Handler(router)

	fmt.Println("Server started on port: ")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
