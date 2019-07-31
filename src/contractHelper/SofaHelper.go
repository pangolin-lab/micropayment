package contractHelper

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"log"
	"math/big"
	"micropayment/contracts/PayToCheck"
	"micropayment/contracts/SofaCrowdSale"
	"micropayment/contracts/SofaToken"
	"micropayment/src/ethutils"
	"micropayment/src/resource"
	"micropayment/src/utils"
)

func SofaBanlanceOf(token *SofaToken.SofaToken,address common.Address) {
	balance, err := token.BalanceOf(nil, address)
	if err != nil {
		log.Fatalf("query balance error:%v", err)
	}
	fmt.Printf(address.String()+"'s balance is %s\n", balance)
}

func SofaTransfer(token *SofaToken.SofaToken,fromAuth *bind.TransactOpts, to common.Address, amount *big.Int){
	conn:=ethutils.Conn(resource.Rawurl)
	ctx := context.Background()
	tx,err:=token.Transfer(&bind.TransactOpts{
		From:fromAuth.From,
		Signer:fromAuth.Signer,
		Value: big.NewInt(0)},
		to,amount)
	if err != nil {
		log.Fatalf("test testTransfer error:%v", err)
	}
	receipt, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		log.Fatalf("query balance error:%v", err)
	}else{
		fmt.Println("transaction \nfrom: ",fromAuth.From.String(),
			"\nto: ",to.String(),
			"\nwith SOFA amount: ", amount.String(),
			"\ncost gas: ", receipt.GasUsed)
	}
}

func SofaBuyToken(crowd *SofaCrowdSale.SofaCrowdSale,buyer *bind.TransactOpts, amount *big.Int){
	conn:=ethutils.Conn(resource.Rawurl)
	ctx := context.Background()
	tp,_:=crowd.TokenPrice(nil)
	var product = new(big.Int)
	wei:= product.Mul(tp ,amount)
	tx,err:=crowd.BuyToken(&bind.TransactOpts{
		From:buyer.From,
		Signer:buyer.Signer,
		Value:wei,
		GasLimit:300000},amount)
	if err != nil {
		log.Fatalf("testCrowdSale error:%v", err)
	}
	receipt, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		log.Fatalf("testCrowdSale mine error:%v", err)
	}else{
		fmt.Println("cost gas: ",receipt.GasUsed)
	}

}

func SofaOpenMarket(crowd *SofaCrowdSale.SofaCrowdSale,admin *bind.TransactOpts, tokenPrice *big.Int){
	conn:=ethutils.Conn(resource.Rawurl)
	ctx := context.Background()
	//change tokenprice
	tx,err:=crowd.ResetTokenPrice(admin,tokenPrice)
	if err != nil {
		log.Fatalf("testOpenMarket error:%v", err)
	}
	receipt, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		log.Fatalf("reset token price error:%v", err)
	}else{
		tp,_:=crowd.TokenPrice(nil)
		fmt.Println("change token price to",tp.String(),"cost gas: ",receipt.GasUsed)
	}
	//change market open state
	tx2,err:=crowd.OpenForSale(admin)
	if err != nil {
		log.Println("change open state tx error:%v", err)
	}else{
		receipt2, err := bind.WaitMined(ctx, conn, tx2)
		if err != nil {
			log.Println("change open state error:%v", err)
		}
		state,_:= crowd.SaleOpen(nil)
		fmt.Println("open state",state,"\ncost gas: ",receipt2.GasUsed)
	}
}

func SofaCreateCheckSK(sk string,to common.Address, amount *big.Int,nonce *big.Int, payContractAddress common.Address)(v uint8,r32 [32]byte,s32 [32]byte){
	hash := solsha3.SoliditySHA3(
		solsha3.Address(to),
		solsha3.Uint256(amount),
		solsha3.Uint256(nonce),
		solsha3.Address(payContractAddress))
	pkBytes, _ := hex.DecodeString(sk)
	sig,_ := secp256k1.Sign(hash,pkBytes)
	fmt.Println("hash:\n",hex.EncodeToString(hash))
	fmt.Println("sig:\n",hex.EncodeToString(sig))
	r := sig[0:32]
	copy(r32[:],r)
	s := sig[32:64]
	copy(s32[:],s)
	v =uint8(sig[64])+27
	fmt.Println("to address:\n",to.String())
	fmt.Println("contract address:\n",payContractAddress.String())
	fmt.Println("amount:\n",amount)
	fmt.Println("nonce:\n",nonce)
	fmt.Println("r:\n",hex.EncodeToString(r))
	fmt.Println("s:\n",hex.EncodeToString(s))
	fmt.Println("v:\n",v)
	return
}


func RecoverSofaTokenContract()(*SofaToken.SofaToken,common.Address){
	conn:=ethutils.Conn(resource.Rawurl)
	jsonStr:=[]byte(utils.ReadTextAll(resource.Contractaddress))
	var dat map[string]interface{}
	if err := json.Unmarshal(jsonStr, &dat); err != nil {
		panic(err)
	}
	address:=common.HexToAddress(dat["SofaToken"].(string))
	token, err := SofaToken.NewSofaToken(address, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}else{
		fmt.Println("recover sofa token contract of address:\n",
			dat["SofaToken"].(string))
	}
	return token,address
}

func RecoverCrowdSaleContract()(*SofaCrowdSale.SofaCrowdSale,common.Address){
	conn:=ethutils.Conn(resource.Rawurl)
	jsonStr:=[]byte(utils.ReadTextAll(resource.Contractaddress))
	var dat map[string]interface{}
	if err := json.Unmarshal(jsonStr, &dat); err != nil {
		panic(err)
	}
	address:= common.HexToAddress(dat["SofaCrowdSale"].(string))
	crowd, err := SofaCrowdSale.NewSofaCrowdSale(address, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a crowdsale contract: %v", err)
	}else{
		fmt.Println("recover sofa crowd sale contract of address:\n",
			dat["SofaCrowdSale"].(string))
	}
	return crowd,address
}

func RecoverPayToCheckContract()(*PayToCheck.PayToCheck,common.Address)  {
	conn:=ethutils.Conn(resource.Rawurl)
	jsonStr:=[]byte(utils.ReadTextAll(resource.Contractaddress))
	var dat map[string]interface{}
	if err := json.Unmarshal(jsonStr, &dat); err != nil {
		panic(err)
	}
	address:=common.HexToAddress(dat["PayToCheck"].(string))
	payToCheck,err:= PayToCheck.NewPayToCheck(address, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a payToCheck contract: %v", err)
	}else{
		fmt.Println("recover sofa payToCheck contract of address:\n",
			dat["PayToCheck"].(string))
	}
	return payToCheck,address
}

func SofaClaimCheck(sender *bind.TransactOpts,amount *big.Int,nonce *big.Int,v uint8, r [32]byte,s [32]byte, payToCheck *PayToCheck.PayToCheck){
	ctx := context.Background()
	conn:=ethutils.Conn(resource.Rawurl)
	tx,err:=payToCheck.ClaimPayment(sender,amount,nonce,v,r,s,)
	if err != nil {
		log.Fatalf("test claim :%v", err)
	}
	receipt, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		log.Fatalf("test claim error:%v", err)
	}else{
		fmt.Println("cost gas: ",receipt.GasUsed)
	}
}