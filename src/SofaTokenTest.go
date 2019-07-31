package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	sofa "micropayment/src/contractHelper"
	"micropayment/src/ethutils"
	"micropayment/src/resource"
)




func main(){
	auth:= ethutils.Auth(resource.Keypath, resource.Passpath)
	token,_ := sofa.RecoverSofaTokenContract()
	crowd,crowdAddress := sofa.RecoverCrowdSaleContract()
	payToCheck,payToCheckAddress := sofa.RecoverPayToCheckContract()
	if(auth==nil || token==nil || crowd == nil || payToCheck==nil){
		log.Fatal("not good")
	}
	fmt.Println(crowd,payToCheck)
	fmt.Println("========== before transaction =============")
	sofa.SofaBanlanceOf(token,auth[0].From)
	sofa.SofaBanlanceOf(token,auth[1].From)
	sofa.SofaTransfer(token,auth[0],auth[1].From,big.NewInt(1000))
	fmt.Println("========== after transaction ==============")
	sofa.SofaBanlanceOf(token,auth[0].From)
	sofa.SofaBanlanceOf(token,auth[1].From)
	fmt.Println("========== change token price and open for crowd sale ==============")
	sofa.SofaOpenMarket(crowd,auth[0],big.NewInt(1000000000000000))
	fmt.Println("=========== buy toke from crowd sale contract ===================")
	//transfer token to crowd sale contracr
	sofa.SofaTransfer(token,auth[0],crowdAddress,big.NewInt(10000))
	sofa.SofaBanlanceOf(token,crowdAddress)
	sofa.SofaBuyToken(crowd,auth[2],big.NewInt(100))
	sofa.SofaBanlanceOf(token,auth[2].From)
	fmt.Println("=========== auth[0] deposit token to payToCheckContract ============")
	sofa.SofaTransfer(token,auth[0],payToCheckAddress,big.NewInt(10000))
	fmt.Println("=========== auth[0] sign check to auth[3] with private key===========")
	keys:= ethutils.PrivateKeyRecover(resource.Keypath, resource.Passpath)
	sk := fmt.Sprintf("%x", keys[0].PrivateKey.D)
	amount :=big.NewInt(100)
	bg := big.NewInt(99999999999999999)
	nonce, _ := rand.Int(rand.Reader, bg)
	v,r,s:= sofa.SofaCreateCheckSK(sk,auth[3].From,amount,nonce,payToCheckAddress)
	fmt.Println("============ before auth[3] claim check ================")
	sofa.SofaBanlanceOf(token,payToCheckAddress)
	sofa.SofaBanlanceOf(token,auth[3].From)
	sofa.SofaClaimCheck(auth[3],amount,nonce,v,r,s,payToCheck)
	fmt.Println("============ after auth[3] claim check ================")
	sofa.SofaBanlanceOf(token,payToCheckAddress)
	sofa.SofaBanlanceOf(token,auth[3].From)
}