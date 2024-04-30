package encryption

import (
	"encoding/hex"
	"strings"

	"github.com/rxjh-emu/server/share/log"
)

type PasswordEncryption struct {
	cryptKey [12]byte
}

func NewEncryption() *PasswordEncryption {
	return &PasswordEncryption{
		cryptKey: [12]byte{170, 171, 172, 173, 174, 175, 186, 187, 188, 189, 190, 191},
	}
}

func (e *PasswordEncryption) DecryptPassword(encpwd string) string {
	log.Debugf("Decrypting password: %s", encpwd)
	// Extract the key index from the first two characters of encpwd
	keyHex := encpwd[0:2]
	keyIndex, _ := hex.DecodeString(keyHex)
	cryptKey := e.cryptKey[keyIndex[0]]

	// Convert the encrypted password hex string to bytes
	encBytes, _ := hex.DecodeString(encpwd)

	// Decrypt the password bytes using the cryptKey
	for i := 0; i < len(encBytes); i++ {
		encBytes[i] ^= cryptKey
	}

	// Convert the decrypted bytes back to hex string
	decPwd := hex.EncodeToString(encBytes)

	// Convert to lowercase and remove the first two and last two characters
	decPwd = strings.ToLower(decPwd[2 : len(decPwd)-2])

	return decPwd
}
