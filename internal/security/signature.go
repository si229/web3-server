package signature

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func PersonalSign(privateKeyHex string, message string) (string, string) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", ""
	}
	signatureBytes, err := crypto.Sign(eip191Hash(message).Bytes(), privateKey)

	if err != nil {
		return "", ""
	}
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKey).Hex()
	signatureBytes[64] += 27
	return address, hexutil.Encode(signatureBytes)
}

func eip191Hash(message string) common.Hash {
	data := []byte(message)
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256Hash([]byte(msg))
}

func Verify(address string, message string, signatureHex string) bool {

	sigBytes, err := hexutil.Decode(signatureHex)
	if err != nil {
		return false
	}

	sigBytes[64] %= 27
	if sigBytes[64] != 0 && sigBytes[64] != 1 {
		return false
	}

	pkey, err := crypto.SigToPub(eip191Hash(message).Bytes(), sigBytes)
	if err != nil {
		return false
	}

	paddress := crypto.PubkeyToAddress(*pkey).Hex()

	return strings.EqualFold(strings.ToLower(address), strings.ToLower(paddress))

}
