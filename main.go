package main

import (
	"fmt"

	webapp "github.com/chainHero/heroes-service/web"

	"github.com/chainHero/heroes-service/blockchain"
	"github.com/chainHero/heroes-service/web/controller"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.hf.chainhero.io",

		// Channel parameters
		ChannelID:     "chainhero",
		ChannelConfig: "/home/simpler/go/src/github.com/chainHero/heroes-service/fixtures/artifacts/chainhero.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "heroes-service",
		ChaincodeGoPath: "/home/simpler/go",
		ChaincodePath:   "github.com/chainHero/heroes-service/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}
	// payload, err1 := fSetup.RegisterSchool("cqu")
	// if err1 != nil {

	// }
	// fmt.Println(string(payload))
	// payload2, err2 := fSetup.RegisterCompany("baidu")
	// if err2 != nil {
	// 	fmt.Println(err2.Error())
	// }
	// fmt.Println(payload2)

	// Launch the web application listening
	app := &controller.Application{
		Fabric: &fSetup,
	}
	webapp.Webapp(app)
}
