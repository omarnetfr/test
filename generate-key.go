package main
import (
	"encoding/base64"
	"encoding/json"
	"log"
	"fmt"
	"github.com/hashicorp/vault/helper/dhutil"
)
const dhpath = "."

type KeyInfo struct {
        Curve25519PrivateKey string `json:"curve25519_private_key"`
        Curve25519PublicKey []byte `json:"curve25519_public_key"`
}
func main() {
	pub, pri, err := dhutil.GeneratePublicPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

    mKey, err := json.Marshal(&KeyInfo{
        Curve25519PrivateKey: base64.StdEncoding.EncodeToString(pri),
        Curve25519PublicKey: pub,
    })

    if err != nil {
            log.Fatal(err)
    }

    fmt.Println(string(mKey))
}




