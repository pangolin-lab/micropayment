
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/proton-lab/micropayment/contracts/PayToCheck"
	"github.com/proton-lab/micropayment/contracts/CrowdSale"
	"github.com/proton-lab/micropayment/contracts/Token"
	"github.com/proton-lab/micropayment/src/ethutils"
	"github.com/proton-lab/micropayment/src/resource"
	"github.com/proton-lab/micropayment/src/utils"
	"log"
	"math/big"
	"time"
)



func main(){
	tokenAddress := deployToken()
	crowdAddress := deployCrowdSale(tokenAddress)
	payAddress:=deployMicroPayment(tokenAddress)
	addressToSave := resource.ContractFamily{Token:tokenAddress.String(),CrowdSale:crowdAddress.String(),PayToCheck:payAddress.String()}
	b, err := json.Marshal(addressToSave)
	if err != nil {
		fmt.Println(err)
	}else{
		utils.WriteToFileOverride(resource.Contractaddress,string(b))
	}

}

func deployToken() common.Address{
	auth:= ethutils.Auth(resource.Keypath, resource.Passpath)
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn:=ethutils.Conn(resource.Rawurl)
	// Deploy a new awesome contract for the binding demo
	address, tx, token, err := Token.DeployToken(auth[0], conn, big.NewInt(1000000000000))
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Token Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	startTime := time.Now()
	fmt.Printf("TX start @:%s", time.Now())
	ctx := context.Background()
	addressAfterMined, err := bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
	}
	fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))
	if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
		log.Fatalf("mined address :%s,before mined address:%s", addressAfterMined, address)
	}
	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
	return addressAfterMined;
}

func deployCrowdSale(tokenAddress common.Address) common.Address{
	auth:= ethutils.Auth(resource.Keypath, resource.Passpath)
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn:=ethutils.Conn(resource.Rawurl)
	// Deploy a new awesome contract for the binding demo
	address, tx, crowd, err := CrowdSale.DeployCrowdSale(auth[0], conn, tokenAddress,big.NewInt(1000000000000000))
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("CrowdSale Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	startTime := time.Now()
	fmt.Printf("TX start @:%s", time.Now())
	ctx := context.Background()
	addressAfterMined, err := bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
	}
	fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))
	if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
		log.Fatalf("mined address :%s,before mined address:%s", addressAfterMined, address)
	}
	tc, err := crowd.TokenContract(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending contract:", tc.String())
	return addressAfterMined
}

func deployMicroPayment(tokenAddress common.Address) common.Address{
	auth:= ethutils.Auth(resource.Keypath, resource.Passpath)
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn:=ethutils.Conn(resource.Rawurl)
	// Deploy a new awesome contract for the binding demo
	address, tx, pay, err := PayToCheck.DeployPayToCheck(auth[0], conn, tokenAddress)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("PayToCheck Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	startTime := time.Now()
	fmt.Printf("TX start @:%s", time.Now())
	ctx := context.Background()
	addressAfterMined, err := bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
	}
	fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))
	if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
		log.Fatalf("mined address :%s,before mined address:%s", addressAfterMined, address)
	}
	tc, err := pay.TokenContract(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending contract:", tc.String())
	return addressAfterMined
}