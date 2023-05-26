package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"go-markdown/serivce/response"
	"io"
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
	passphrase := "5777a5d6b4c4fc15f9d8ee5691012e3ee8d0f6a83f61b4084685d55bc0811fbe"

	byteArr := make([]byte, 32)
	AESKey := GenerateAESKey(passphrase, byteArr)
	fmt.Println(AESKey)
	decodeAESKey, err := base64.StdEncoding.DecodeString(AESKey)
	fmt.Println(decodeAESKey)
	if err != nil {
		fmt.Println("Endocde String: ", err)
	}

	encryptedData := "TUNXYW03MUdNQXFwVkJDbTlGcm95QzJpZm1LQndJQ2hQNVhHQWZNdm1rSktBeDhnWGUwY3dGZGF4cTlzZDViRkNnempXM1hKOWRJSDBHZzJ5aXNkbjg0dXZuZXFvcGdOUkRwYU53bzJjS1ZKcmlzZkQxTG5sbGFNZGFVMmlmYlBFQzhvcWZuQVpJbUhYYWl3bkc5eHlQU2t3ZHV3N0F6cGczOXJZMWVxZWJ3UHdrZ1ZURWlPL2hnOEhpdTNqdVA3aDY0NGg0N3BCYnVmSTl1V2RSc0dJQ0dqQUk1dTBsWnhHYjBUMUs5QXF0dnBMMTFNcFRzcHZ1QWlnSHRDdXRkUHVtZWhuZHNvNlluci9nV0pob0lQUVFPbVhHZTZlbTZRVDNSVlNRakdIV3NBV3E1OXlheVpkbTJaUmI2aXY4VzNGaWlyNkx2YXQ4WTBPYkRYZGJtMTIrTi9SeXVNZUhueHNxT0pPZzljbnFGQkQ2V3dKSDhuVjhMdFBSSFFkRk9rbmpreU9HZVRiOC8rb2l5YmVaNnBuc0JZS1pXMkVsYjM4MFBkdFVKT3l1eUtpcExJbnRiVFhYRy9MQ3k0Mml3b0VSWDJzYkdROHk0RGI2Sy9JRWhlWkNwbW96ZnZiNHR6aWRSS2ozSGthYU8zaFVBZ0FuWGVTTE1SbjcxMGhkOVM0TVErOUZ6b1l0SHZPeG42TS9OeTJ6c2dwYWN4Tm16SGVrQ2EzcWlNTDBuUFd3c1JpSCtSSDFhcjRKQTVTWUZoZTZNSkxDblIxWS9kNitnV3ZtTXArNXhuNTIwMHBEUUxPbTF3NUl5OTFCUGxYS1AxQ3dSOUZQbE1RNzIyNVA0RkFPTS94ZXBoOHE4VDFiTHVRQmpqb0c3c3VrZWx4MWxONGNKU3NBT2Fvc2hWZmphdVJlZEp5RkNPVFJmUENoSzZZekU1MEovN0tGTjRicXRNeHFCWmE2SGhwVWtzSS9rTjhSSzVSWkV5eldLNDlqUGhPOGg3czBZMkRkNmVRMmtQOStxaC9aTkNsczdaUUFuRFZkZTlnSm9EbGorenFENHZDcGpQdlRuYTN4aGZCcklBN1ViTkJWYityNDBQZ214TVdIakJNODBCVWpBV2xjWm9FRC9NdVFqZDAwclZ3ZUl6ZlNJR2QzRHo2Ym1Qajk2cVAzeXVDN1ZOdGVCRDUvSTBzV3JZbXZtUTZpWUhyWTJtTHVpd3EvOEZKVStrQzVUNy9aVW5jcEVkb0VDZVpOQlltREJ4Sk1ySUYwc1dGLzgxbVQxRFgybkhmV0M0VlVob2grZW9Hdko4TWJOcjc4bjF0NzdrZGhPMEtwdVlwaE9vSWpLRW0wQnoxa1VtbThOdWszcnBINlZDZFJTOURYM0RueXJqTTV1NFNWOFhpb0d6VlRwbEdlekcwbE5BRGNjMk1rVnRscEFvRW0xWlN4Y0c1Qkl4SXpDZVMrbzNETHJ6WVU0SzhRYzNSUlBVeitnL3B0WU1mSzVCUnc4UFdNOWV5VVhFazhLS0g3YjhqRktYcEYxYjdsTVZFL3ZtZjg3UGIxZENUTGJHdCt6R2VscDUrR0JTUXBrek93c1BHall3RVljbkNpV2hJb3BNNjgra1VkazlyK2I2bitsOGhZWnQ2Yk5yWk9RY2VibmZSTzk0dFJZWWJ4OFBHNUVMVm9KSWZNZE1nZWFWUWFPc0cyZlIreXdiL1RWZXE2TjJFYVdLKy9mR3RKa05jMFdUc2dYNjZVbVJIVkpLMVc3djAzK0NYTFF0blc0WDh3MThJWnVYRmRMZ1kvK3pPSDhJanRRenpsYmE3ZnhsbTNzMzBvMFNQeURZY2RNdWxmUkUyYnQ2VlFPYTQ0WEd1RFBHR3ZiMnM5bXplc3phSVRMdk05aXUvSGVrdlhTLzlVMm12b2xhWERhN3RsWG1mR0FMVEVLS3BvTHZlQ3N2N1VZRzhJRzNOemJmNnREK00xenpjcHNncGFxSHo0TWV6MU8xdkUrWWpoTko2OFErWWZoa1d1TTJzSGhOU1ZXU2RpVmFzQnQrYVorYmZ3S3YxbTBBQ29Ta3BPTjE2THRiVTd5bjBhcUpVbzAzMGd4QTFDWWRuK1dZSDNiUDNYU01UbHNhYnBDQWVFMk1QdTV3K3pzeTh6bzI5RGlQNFE3cUUrL3pQQUJEUjh0K3VyZTBremNjM1VZR09zRGFwemJ5akR2NUg3cHFpNFVuNDJSNlladm9mN1h1RGpQVTN3dCtibDcvQkd6SlVtMUk4NSs5TlJYZmtabkltQTFlc1VwaEVNbjJqM1loNVU0MWFsNm50ZTRtb3dLTjdHenIvdEpMWFJsUGRyRXVFcDRCNTZPY0pZalJWTWxWWlJUcmxwNElqU3RtUXZzM3ZZcWM3anNNMU0yWDNlYlZqOWFYRWU5dzY1M242Snl3UkFnT0lTUVhQVXNyNDZYNGdSRnFhZFZDNGo2Mm04NDVZZWR3WmxBTU5VV2R5dFlldll0QVF2cGlxZTlSbnBRTmxYWXV4THdzUVB0RDhIYzBNckRsMEJIMkJQOHNFcEgraTdmK2xYYkhnSXJWZVpxclJ2K2p6TFhZM29uMzl4Q0hmU1gxTFR5R05RR0VnOG1wYUVUY3hqV1NMK3hhSlNweGFVY00vYVk5MzNVSnJIUHN2UHVYamVFQUNhamtnQTVuVU02YlNzNVFvdDdoVmhNMGdyejRjWmtwSmphbFh6VVkzb2hhVnZ5NjV2cUw1MFpsWExnMTVlckcyU3RaY0x2YTYvdz0="
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

func TestEncryptAes(w http.ResponseWriter, r *http.Request) {
	key := []byte("ck3SXyuVFxBW5enWyq7qKH4WSbz7RVfB")

	// Copy the input data to the byte array
	fmt.Println("panjnag string key  ==== ", len(key))

	plaintext := []byte("HALOOO BITCH BABII HEHEHE")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	fmt.Printf("%x", ciphertext)

	output := base64.StdEncoding.EncodeToString(ciphertext)

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"output": output,
	}).ParseResponse(w, r)
}

func TestingDecryptAes(w http.ResponseWriter, r *http.Request) {
	strings := hex.EncodeToString([]byte("bimapw04"))
	fmt.Println("sasasa ====", strings)

	key := []byte(strings)

	enc, err := base64.StdEncoding.DecodeString("2Mwb+geUZi7sVhlhzVJfv7uRcKRq59fGND/N+SmQocIZyZ4SNpCP9rhFnCB39odMn5F95Wk=")
	if err != nil {
		fmt.Println("masuk ini ", err)
	}
	fmt.Println("sasa ==== ", enc)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s", plaintext)

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"output": string(plaintext),
	}).ParseResponse(w, r)
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
