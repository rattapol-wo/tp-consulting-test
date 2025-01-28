package main

import (
	"log"
	"tpconsulting/src/client"
	accountHandler "tpconsulting/src/handler/account"
	campaignHandler "tpconsulting/src/handler/campaign"
	"tpconsulting/src/repositories"
	"tpconsulting/src/usecase/account"
	"tpconsulting/src/usecase/campaign"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

func main()  {
	db, err := client.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	accountRepo := repositories.NewAccountRepo(db)
	pointRepo := repositories.NewPointRepo(db)
	campaignRepo := repositories.NewCampaignRepo(db)
	
	accountUseCase := account.NewAccountUseCase(db, accountRepo, pointRepo)
	campaignUseCase := campaign.NewCampaignUseCase(accountRepo, pointRepo, campaignRepo)

	campaignHandler := campaignHandler.NewCampaignUseCase(campaignUseCase)
	accountHandler := accountHandler.NewAccountUseCase(accountUseCase)

	e := echo.New()

	// campaign
	e.POST("/campaign/create", campaignHandler.CreateCampaign)
	e.PUT("/campaign/add-point", campaignHandler.CampaignAddPoint)

	// account
	e.POST("/account/create", accountHandler.CreateAccount)
	e.GET("/account/point", accountHandler.GetPointByMobileNumber)

	e.Logger.Fatal(e.Start(":8000"))
}

