package connection
import(
"crypto/sha1"
	"encoding/hex"
)

func GetSHA1(inputString string) (string) {

	 
	h := sha1.New()
	h.Write([]byte(inputString))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}
