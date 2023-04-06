package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"go-markdown/serivce/response"
	"net/http"

	"golang.org/x/crypto/pbkdf2"
)

func EncryptAES(w http.ResponseWriter, r *http.Request) {

	// key := []byte("")
	// c, err := aes.NewCipher(key)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// out := make([]byte, len(plaintext))

	// c.Encrypt(out, []byte(plaintext))

	// return hex.EncodeToString(out)
}

// "337e06ed25a39f3bb05215b2de884c8212fb29bbae33d9706b8bf3235c6c8d28"
// "yjYyox6/GFFD5rEu14pRlyRLkYUcC69dXJD4C+6KPcuATIgKigZilJsXW0LxDiKckESp6pbUavd1jgCGbvDa/0rvNRw06RCIE1A78CtCBYjDuT9SykJaM/m2kJ2AzfUiyocO/179cIWBJiqEqURqoXVX7f4Z33jshFy08j+X38Uue1x2CndYJ+rjPXhIkrxEdK0o839laphe/U62Us4gVs6yWK0Y/MEinYTKIdMpXL8XWaHWA4yIQSxXfGz+Lp7znoA1tlUy5CNr1ndIbt8blr5LwcJ8329zwZTLXvsjVL8YF7387Mt4D8fTPomL5JGxVncFu06BJD5v2ha4WCYYJ5rtGftHXiDSLm11qSzPm0H7jkQVdtDYalhFnlnJn3I+VezEG/5ZOqd8p0NxqnW2YUJhcV3cDdCr9HAFhrDCOhTlIA=="

func DecryptAES(w http.ResponseWriter, r *http.Request) {
	passphrase := "b215190d10f7dc0bf7fe5fd78f568ac7e3c4845c9df9943092aa46fe6f621a40"

	byteArr := make([]byte, 32)
	AESKey := GenerateAESKey(passphrase, byteArr)
	fmt.Println(AESKey)
	decodeAESKey, err := base64.StdEncoding.DecodeString(AESKey)
	fmt.Println(decodeAESKey)
	if err != nil {
		fmt.Println("Endocde String: ", err)
	}

	encryptedData := "cmFxcEk2QVpGTzdtQkVPUUNmTDNCR3Fncjl0OWZEWGZBOGJ6Ry9zUDhFQ3JSenYrYm9rTzFkMUlHRVdrV1lYNlluU3ZlMmtoQ2lIOGJEL2FhUFk5VEJLM2h1eGVqSE9TZkUxV2lsUi9qdWRXQzh6R0VCOVBmUmV2SkxVNjJxUTVzUjhpYjlnVnRaRWVpbXp6eFlaVW04enZ0NVlxQS9YL0hXMlVyZXAyRklzU1dOTHpzY2ZLUzNSQTYyRWJWNlhFZFk2THZjSlFLOUhZMm9VMjdmUmc2UnBWNXkvb20zcyt1TDhOTlJUYUo3U2RHWC9GZE02eUZGTm1lUG9UM21meGxGQ25iWjhsc0VUeGg5RVIyblBaWCtEM0o1SkR0YlZEWjR0WmZqeFRhb2poaWFsQVVCcXZZelpjeUhOYWNBZExsMU9WdEw5a0szdXRoMkh0SEE5RU1JeUNnZW1RYVZCV3NyTU1QTFFqU1hoU2h2b1B6alI5c2NqZWV3NnlUeVRTcG4raWsxQmFaYkU4YVRDMXpsYXdqMytFejZEd1RTdkgvSlRhck42dE0zRkNyT3h5bmJyVHNLWE5MOHZSU01wVFViaC9wZEk0bCsycTNDbkptWmVXWElkVS9EVnNid2FRcGlZYVFtMHBqSnBQeXhWME8vd1NzOGpQQnA5T3k3T2NNNHFUOWZhcjZ2bEdySnNtMjh5TE8wSXBtWlg4dFBIMVE3aUtFTWNQU3NyOG1BTkdnaGY4YS9NMGl0SzA1SDc4T2tFVkNPbXgwd0Fxb0dhbmFkb1kvNGpqU29oajA0UFUrb09TckF5R2ZseHZ2aXVlRklkYXVhUWVmbGhXOElCT1I3SEp3WHV6aDd5Q0VoSjN4TGpJeFZYMjlrZU10MTFwVlg5MVlHUkJnWHJBY0g2OUF4SFBhemV2QmwvVEVLaXUwMml4KzMwblo0NlhVaGQvVlhrUGtKV3lBQTYvMXdzT2t0VHpXdWowNUk4OTJTZVRSalFFR3FEWmRNWUxOZ09hbkpHUmErbFQvRDhwakdNdVBxamhEYnhBNHhWeDNvb09yMGlPM29rOVZCcTdtbFZlUm9kUkdidGk4QTU1UUkxb3dIU0RGVTlMWXlxelI1TEk5SXRHK2liNDFZSWlVL3VwbzcrSjhkWEF0MGxSZVNYQzFCMHhOd2d1OHpwVTdhaGpZQ3NpTFJqVm1DeDlJM1VLNWV1TzM5WWFxVitDWmNHaFBVR2VMTHpMblA3dXR3ZmluWjRiSGt6V2pZdUlXTkNrOHhJbC9TTDNRd0FzQi9xVWJ6YitjMEpnbjB3aUY4dk0xQStZR2NHZW5wYjdLZXpTcnVMd0FIMEZacnBBZXVIaG52ZDgxNTNLNnhma3dqTWNZZmJiOXFkbUVJdG9nT2hQdXZPblVFamNzcGlaTnFJSnpJWW1Ba0haSVhsVXJ6YU5QcFl0WVBpVThvZHQxT3RvWm1YcHlnakJFQWRKMzh6S1FZeUwvUllxV1FnMXVzWjViZVAwa2xBejNERldaY0g2V2JTY1pwMVpETnYrendhWlY1cE1jT2tvTkp4Lzh2YnNpcWlNa1JTWlQyZGFSU1NmTy80WjYvRGpyV2hOMzdwVDlXYTJCdTluZDVYZGFnbU1QcXhPc0RMcnNMcktzcE5jaFlTTWhxWWsxMWE5SXo0U0t5QUV5Rk8vZFhvTElCUS9EaTNZV3RzOTdUb3ZRUWYyNndTcEJzbGJmN0hIZG5Bc0g3M3FNckpwYUhaWlFaRXgrRkZHcDRmbFVsRXptMUpEMmRERjltaHpHbThBeUtlYnMzRTdZWExRSmhvczJVMDAwNVZJbGd6K0JUZHhHSTlxZnFvRVVnNEJxWEU4VUlmT1RvcTNUeU03eGFKajM5ZUZFWmhoRGUvd2czN1JXcnNyUk1GNWNVUURvOERSbmtXdERKbnczQXVzL1ZvTDFjbTR1OGRFRmt6dTk4akcza0ovblhBcWtBUmdiY295OE5xdXBKQ1lLdHlRam5JWkpYOEg4eFpaMDVSdGZZOXUzWm9UNFNicGJSb3UwK1BWdVpaRjhzUTRMWVp0Z1luUEJHUGV6amM3dGFJa2tyZ1lVQ1hhZ09RMXpFcklCaXI5S0tMZ1JmUWZBTEl4eE4xczdIdm96NlFMS0Y2SE1LVk9KdDhPZ2lsUDFJd3M5MXl1K0NlNjJvWjRvcWtNT3hmQ21vdXlmOUk2RUtuRlZyRnJMbHlqSmFtQTZTRnVxM3BsY0lEbUdlc2hCTUNQektiQU5HWUwzQVFlNktFMVlma2EyaUZCdm1ORFVoRHQ0RTlQdTVNZjNXcDNEU0VQbldKeVFLTHdaRXJOdENnQlU4SXMxTE55ajZyem9qQWc2ZmtKY040WEdGUzg5L1VqWnJJV2FKVzdxanZHc3ZkQXA2NFY1Unl2V3dJeUtkT1AwMXlWOEtzTEJaSnNCVkdMVW1PVlRNMUVRcXJXbFoxRC9zWUNVK0hrcnVaakdzNWE4cmxtVXMyS2s0YzVpZUc5aWZqbkRWaUFla0FGQ1hjSVYyanY3NURoVVhTQVhQZDVEY2d0UmI1WW9mVXNrZE85djBvck1vamorN1dEM0NnNU9uMXpFcTd5ZUlkTnhoWT0="
	encryptedDataWithAeskeyBase64, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		fmt.Println("base Data: ", err)
	}
	decryptedData, err := DecryptDataWithAES(decodeAESKey, string(encryptedDataWithAeskeyBase64))
	if err != nil {
		fmt.Println("Decrypted Data: ", err)
	}

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"encypt_data": string(decryptedData),
	}).ParseResponse(w, r)
}

func GenerateAESKey(passphrase string, salt []byte) string {
	if salt == nil {
		salt = make([]byte, 8)
		rand.Read(salt)
	}

	key := pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New)
	stringKey := base64.StdEncoding.EncodeToString(key)

	return stringKey
}

func DecryptDataWithAES(key []byte, chiperText string) (string, error) {
	cipherTextString, err := base64.StdEncoding.DecodeString(chiperText)
	if err != nil {
		fmt.Println("masuk sini 1", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("masuk sini 2", err)
	}

	if len(cipherTextString) < aes.BlockSize {
		fmt.Println("masuk sini 3", err)
	}

	iv := cipherTextString[:aes.BlockSize]
	cipherTextString = cipherTextString[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherTextString, cipherTextString)

	return string(cipherTextString), nil
}

func DecryptAES1(w http.ResponseWriter, r *http.Request) {
	key := []byte("74e61aca5e74917d516cef62c1262be22fbf559982782ff91221b9fd41cb757c")
	cipherTextString, _ := base64.StdEncoding.DecodeString("S3BNSEdzTU8weFBMa0crVDY0SUowejFZNit4dXpHMXFOMGk1ZzdCTS9JTHF5VzVrYzRLS0hkd2ZNYVQ0dUMyOEFUbDlpV21EcXp4MGdseHRjNlVxb24wek1zbUlIOVptRjZPenBINXJXcEVXeEZ0NDgzSlRxQlVSQkdVUG1XTW1nL2kwb0ovYnRucklWeWNUd2RhTTRMNDFDNVhCdGVRTnpXREMzb2tkS2Z6TkZXWDdXZ2dKS0FLWmV6NnVrZFd3ZGwyclZsdzFjeEJrSzgza1hvNVpDUVQvRDZUWlEvSEsxakVmYmdBbitLUkoxakl3L1kvdVBuaVJ6Z2hYdTdCd0pBVFRDTTRsNUtTTThpQ0NnSkQ3TCtBdlRId0dOczRKeGFjNHM2V2dBU1cyTlZQckRoNEtEdnJZWkhHZVI2cm5RRnBWOWg2ZU9EWS85YytLUVFHVWNUTXA1MkpOMldsdHJDZkZFY0pPaTJCTlhYSXlYblhUUkM2K1VNT0UyRTBiaVZDaStVOFhYMWlCYmpmY21IUHprc3pEdFl0eDZab2tKV0hoS3QyMk41UWpQN2JpOXdUcmpic2tIdU1xayt4dUdnMD0=")
	// if our program was unable to read the file
	// print out the reason why it can't
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	iv := cipherTextString[:aes.BlockSize]
	cipherTextString = cipherTextString[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherTextString, cipherTextString)

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"encypt_data": string(cipherTextString),
	}).ParseResponse(w, r)
}
