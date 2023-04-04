package smTest

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

var private_pem_pwd = []byte(">WY_.4Q236_Isa*}")

func Sm2WriteKeyFile() {
	priv, _ := sm2.GenerateKey(rand.Reader)
	pub := &priv.PublicKey

	privPem, _ := x509.WritePrivateKeyToPem(priv, private_pem_pwd)
	ioutil.WriteFile("configproxy.pr", privPem, os.FileMode(0600))

	pubPem, _ := x509.WritePublicKeyToPem(pub)
	ioutil.WriteFile("configproxy.pu", pubPem, os.FileMode(0600))
}

func Sm2Encrypt() {
	pr, _ := ioutil.ReadFile("configproxy.pr")
	pri, _ := x509.ReadPrivateKeyFromPem(pr, private_pem_pwd)
	puf, _ := ioutil.ReadFile("configproxy.pu")
	pub, _ := x509.ReadPublicKeyFromPem(puf)

	origin := "{\"mastername\":\"main\",\"pass\":\"6iDrKRF1OW5sKIvj\",\"sentinels\":\"10.0.19.102:26379\"}"
	enc, _ := pub.EncryptAsn1([]byte(origin), rand.Reader)
	fmt.Printf("encrypt = %s\n", base64.RawStdEncoding.EncodeToString(enc))
	ori, _ := pri.DecryptAsn1(enc)
	fmt.Printf("origin = %s\n", string(ori))
}
