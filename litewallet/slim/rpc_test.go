package slim

import (
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	"github.com/tendermint/tendermint/rpc/client"
	"testing"
)

func TestQOSQueryAccountGet(t *testing.T) {
	SetBlockchainEntrance("192.168.1.23:1317", "forQmoonAddr")
	addr := "address1v26ael2jh0q7aetuk45yqf3jcyyywg2g6wq2tv"
	Aout := QOSQueryAccountGet(addr)
	t.Log(Aout)
}

func TestQSCQueryAccountGet(t *testing.T) {
	SetBlockchainEntrance("192.168.1.23:1317", "forQmoonAddr")
	addr := "address13l90zvt26szkrquutwpgj7kef58mgyntfs46l2"
	Aout := QSCQueryAccountGet(addr)
	t.Log(Aout)
}

func TestQSCtransferSendStr(t *testing.T) {
	SetBlockchainEntrance("192.168.1.23:1317", "forQmoonAddr")
	addrto := "address13l90zvt26szkrquutwpgj7kef58mgyntfs46l2"
	coinstr := "10000qos"
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	chainid := "qos-test"
	Tout := QSCtransferSendStr(addrto, coinstr, privkey, chainid)
	t.Log(Tout)

}

func TestTransferTxToQSC(t *testing.T) {
	t.Log("1111111111111111")

	SetBlockchainEntrance("47.105.52.237:26657", "forQmoonAddr")
	t.Log("2222222222222222")

	privkey := "wnEmxnWFgT93M5a9l7aPTdkxM8MLoenyMe60sD/8rqzslA7MvfoHydXqL4QGbplLhIlEbLAZ/0ue9G1rjBFMfQ=="
	chainid := "test-chain-xHEkEv"
	Tout := AdvertisersTrue(privkey, "ATOM", "100000", chainid)
	t.Log(Tout)


	result:=BroadcastTransferTxToQSC(Tout,"sync")
	t.Log(result)

}


func TestQueryAccount(t *testing.T) {
	t.Log("1111111111111111")

	SetBlockchainEntrance("47.105.52.237:26657", "forQmoonAddr")
	t.Log("2222222222222222")

	privkey := "j7wbT1lfG0ZRObMKwyYkVFpwRW0OlzFcgVL6ZwQJjeicPcStxppyWMpD1tfFg2YEcDu5JH+M6g+EhSBmjHq0bg=="


	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
	err1 := Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}

	addrben32,_ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	address, _ := getAddrFromBech32(addrben32)

	fmt.Println("address",string(address))

	account,err := RpcQueryAccount(address)

	if err!=nil{
		t.Error(err.Error())
	}else{
		t.Log("111",account.Nonce)

	}


}



func TestCommHandler(t *testing.T) {
	t.Log("1111111111111111")

	SetBlockchainEntrance("localhost:26657", "forQmoonAddr")
	t.Log("2222222222222222")

	privkey := "9QkouVPl29N2v1lBO1+azUDqm38fAgs6d3Xo8DcnCus7xjMqsavhc190xCGzZuXcjapUahi7Y7v2DD4hzVCAsQ=="
	chainid := "test-chain-dpYlL3"




    args:="[\"address1y9r4pjjnvkmpvw46de8tmwunw4nx4qnz2ax5ux\",\"0\",\"abcde\",\"20\",\"20\",\"10\",\"50\",\"20\",\"3\",\"ATOM\"]"

	Tout := CommHandler("ArticleTx",privkey, args,  chainid)
	t.Log(Tout)


	result:=BroadcastTransferTxToQSC(Tout,"sync")
	t.Log(result)

}

func TestLocalTxGen(t *testing.T) {
	fromStr := "address1vpszt2jp2j8m5l3mutvqserzuu9uylmzydqaj9"

	toStr := "address1eep59h9ez4thymept8nxl0padlrc6r78fsjmp3"

	coinstr := "2qos"
	//generate singed Tx
	chainid := "capricorn-2000"
	nonce := int64(1)
	//gas := NewBigInt(int64(0))

	//PrivKey output
	privkey := "sV5sRbwnR8DddL5e4UC1ntKPiOtGEaOFAqvePTfhJFI9GcC28zmPURSUI6C1oBlnk2ykBcAtIbYUazuCexWyqg=="

	jasonpayload := LocalTxGen(fromStr, toStr, coinstr, chainid, privkey, nonce)

	t.Log("msg\n", string(jasonpayload))

}

//HTTP POST to QOS chain
func TestHttpBrToChain(t *testing.T) {
	fromStr := "address1vpszt2jp2j8m5l3mutvqserzuu9uylmzydqaj9"
	toStr := "address1eep59h9ez4thymept8nxl0padlrc6r78fsjmp3"
	coinstr := "2qos"
	//generate singed Tx
	chainid := "capricorn-2000"
	nonce := int64(1)
	//gas := NewBigInt(int64(0))
	//PrivKey output
	privkey := "sV5sRbwnR8DddL5e4UC1ntKPiOtGEaOFAqvePTfhJFI9GcC28zmPURSUI6C1oBlnk2ykBcAtIbYUazuCexWyqg=="

	jasonpayload := LocalTxGen(fromStr, toStr, coinstr, chainid, privkey, nonce)

	//tbt := new(types.Tx)
	//err := Cdc.UnmarshalJSON(jasonpayload, tbt)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//txBytes, err := Cdc.MarshalBinaryBare(jasonpayload)
	//if err != nil {
	//	panic("use cdc encode object fail")
	//}

	client := client.NewHTTP("tcp://192.168.1.183:26657", "/websocket")
	result, err := client.BroadcastTxCommit(jasonpayload)
	if err != nil {
		fmt.Println(err)
	}

	t.Log(result)
}
