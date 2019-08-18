package main

import (
	"crypto/rand"
	"fmt"
	helper "github.com/proton-lab/micropayment/src/contractHelper"
	"github.com/proton-lab/micropayment/src/ethutils"
	"github.com/proton-lab/micropayment/src/resource"
	"log"
	"math/big"
)




func main(){
	auth:= ethutils.Auth(resource.Keypath, resource.Passpath)
	token,_ := helper.RecoverTokenContract()
	crowd,crowdAddress := helper.RecoverCrowdSaleContract()
	payToCheck,payToCheckAddress := helper.RecoverPayToCheckContract()
	if(auth==nil || token==nil || crowd == nil || payToCheck==nil){
		log.Fatal("not good")
	}
	fmt.Println(crowd,payToCheck)
	fmt.Println("========== before transaction =============")
	helper.BanlanceOf(token,auth[0].From)
	helper.BanlanceOf(token,auth[1].From)
	helper.Transfer(token,auth[0],auth[1].From,big.NewInt(1000))
	fmt.Println("========== after transaction ==============")
	helper.BanlanceOf(token,auth[0].From)
	helper.BanlanceOf(token,auth[1].From)
	fmt.Println("========== change token price and open for crowd sale ==============")
	helper.OpenMarket(crowd,auth[0],big.NewInt(1000000000000000))
	fmt.Println("=========== buy toke from crowd sale contract ===================")
	//transfer token to crowd sale contracr
	helper.Transfer(token,auth[0],crowdAddress,big.NewInt(10000))
	helper.BanlanceOf(token,crowdAddress)
	helper.BuyToken(crowd,auth[2],big.NewInt(100))
	helper.BanlanceOf(token,auth[2].From)
	fmt.Println("=========== auth[0] deposit token to payToCheckContract ============")
	helper.Transfer(token,auth[0],payToCheckAddress,big.NewInt(10000))
	fmt.Println("=========== auth[0] sign check to auth[3] with private key===========")
	keys:= ethutils.PrivateKeyRecover(resource.Keypath, resource.Passpath)
	sk := fmt.Sprintf("%x", keys[0].PrivateKey.D)
	amount :=big.NewInt(100)
	bg := big.NewInt(99999999999999999)
	nonce, _ := rand.Int(rand.Reader, bg)
	sig:=helper.CreateCheckSKAsm(sk,auth[3].From,amount,nonce,payToCheckAddress)
	fmt.Println("============ before auth[3] claim check ================")
	helper.BanlanceOf(token,payToCheckAddress)
	helper.BanlanceOf(token,auth[3].From)
	helper.ClaimCheckAsm(auth[3],amount,nonce,sig,payToCheck)
	fmt.Println("============ after auth[3] claim check ================")
	helper.BanlanceOf(token,payToCheckAddress)
	helper.BanlanceOf(token,auth[3].From)
}