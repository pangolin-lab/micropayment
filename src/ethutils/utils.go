package ethutils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"micropayment/src/utils"
	"strings"
)

func Auth(keypath string, passpath string) []*bind.TransactOpts {
	keys:= utils.ReadTextByline(keypath)
	pass:=utils.ReadTextAll(passpath)
	var auth  []*bind.TransactOpts
	for i:=0;i<len(keys);i++{
		a, _ := bind.NewTransactor(strings.NewReader(keys[i]), pass)
		auth = append(auth, a)
	}
	return auth
}

func PrivateKeyRecover(keypath string, passpath string)[]*keystore.Key{
	keys:= utils.ReadTextByline(keypath)
	pass:=utils.ReadTextAll(passpath)
	var keyList  []*keystore.Key
	for i:=0;i<len(keys);i++{
		json, err := ioutil.ReadAll(strings.NewReader(keys[i]))
		if err != nil {
			return nil
		}
		key, err := keystore.DecryptKey(json, pass)
		keyList = append(keyList, key)
	}
	return keyList
}






func Conn(rawurl string) *ethclient.Client{
	conn, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	return conn
}

