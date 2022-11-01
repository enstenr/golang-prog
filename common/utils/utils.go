package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func GetSHA1(inputString string) (string) {

	 
	h := sha1.New()
	h.Write([]byte(inputString))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}
func cleanUpMessage(message string,treeName string) ( string) {
	message = strings.Replace(message,",","",-1)
	message=strings.Replace(message,"Metadata tree contains duplicate SKU for ","",-1)
	message=strings.Replace(message,"in GCS. Duplicate SKUs are","",-1)
	message=strings.Replace(message,"in GCS","",-1)
	message=strings.Replace(message,treeName,"",-1)
	message=strings.Replace(message,"'","",-1)
	message=strings.Replace(message,"  "," ",-1)
	message=strings.Trim(message," ")
	return message;
}