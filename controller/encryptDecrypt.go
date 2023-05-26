package controller

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"go-markdown/serivce/response"
	"log"
	"net/http"
	"os"

	"github.com/timemore/foundation/keystore/crypto"
)

func GenerateKey(w http.ResponseWriter, r *http.Request) {
	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}
	publickey := &privatekey.PublicKey

	// dump private key to file
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem, err := os.Create("private.pem")
	if err != nil {
		fmt.Printf("error when create private.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		fmt.Printf("error when encode private pem: %s \n", err)
		os.Exit(1)
	}

	// dump public key to file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		fmt.Printf("error when dumping publickey: %s \n", err)
		os.Exit(1)
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create("public.pem")
	if err != nil {
		fmt.Printf("error when create public.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		fmt.Printf("error when encode public pem: %s \n", err)
		os.Exit(1)
	}
}

func EncryptOld(w http.ResponseWriter, r *http.Request) {

	// privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// if err != nil {
	// 	panic(err)
	// }

	// publicKey := privateKey.PublicKey

	pkcs8DerKey, _ := base64.StdEncoding.DecodeString("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArpSRT+wtFmc9IFnPhYSAXxqfwd6jjXiwAue3d5pCnTB4TWtkZRU5OYM30j8pK94zOsO1ZGydoJm2YW7dz1Kffc0uzLPykBAgSADXN9PehzyhiLJHz9+HcSVT/Y6FSucRVK2BGHzTmspSM5tVaRA3p2IXuSyzpzFm1EsLYlgR40jr6puzKN3Z+02oTyP1nBaJEkD8OA6laWwTkjjnqyZD4Z0F5/sZTM/IsNykxlkdKiJzca7GOGqS56AMPoxgxDorZk9CQLthjSvGun47EyQXkk0wxwFGu5CvnybGu9Xq/D/skvC4rLsnoETCkEDdsdu4T8TaLD+CaGkZEC4fngxyLwIDAQAB")

	key, _ := x509.ParsePKIXPublicKey(pkcs8DerKey)

	publicKey, _ := key.(*rsa.PublicKey)
	rng := rand.Reader

	encryptedBytes, err := rsa.EncryptPKCS1v15(rng, publicKey, []byte(`{"password":"pdyYwEPtTs9q4pyh$y4e","old_password":"pdyYwEPtTs9q4pyh$y4e"}`))
	if err != nil {
		fmt.Println("error when create public.pem: ", err)
	}

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"encypt_data": base64.StdEncoding.EncodeToString(encryptedBytes),
	}).ParseResponse(w, r)
}

func DecryptOld(w http.ResponseWriter, r *http.Request) {
	ciphertext, _ := base64.StdEncoding.DecodeString("K2uUBKnocEkjDEOxXEDd9TaXygMyHUO7Dxm9Ylb1C2kTyIv/Fh2hYFOKDCV40hsd7TSoEhlKvOtL5M8ro/KA1friO/3CST80YEC5qYvfnMkROqZ7gC8iAVBiffMzGu7EWp8IckzELxDWfxCQCUpxyQMsn+cF8CehzimTQ3lgMI4fZesafewVKgVlRyy4GwhOTx3RM4adGV4uEe8wwELPbRg1sBaz6RsczxiSfRorZBWbRun2YOZxs+0mx+s+x53UNwOqvofTQ+XQz1gNXW69UprXoHGCMNV71eO1CiEDL8xO4/g4jYEonUgdezg8Xqzdz0mEd7yRWcx/BiCn0wdNWg==")

	// Import PKCS#8 key
	pkcs8DerKey, _ := base64.StdEncoding.DecodeString("MIIEpQIBAAKCAQEArpSRT+wtFmc9IFnPhYSAXxqfwd6jjXiwAue3d5pCnTB4TWtkZRU5OYM30j8pK94zOsO1ZGydoJm2YW7dz1Kffc0uzLPykBAgSADXN9PehzyhiLJHz9+HcSVT/Y6FSucRVK2BGHzTmspSM5tVaRA3p2IXuSyzpzFm1EsLYlgR40jr6puzKN3Z+02oTyP1nBaJEkD8OA6laWwTkjjnqyZD4Z0F5/sZTM/IsNykxlkdKiJzca7GOGqS56AMPoxgxDorZk9CQLthjSvGun47EyQXkk0wxwFGu5CvnybGu9Xq/D/skvC4rLsnoETCkEDdsdu4T8TaLD+CaGkZEC4fngxyLwIDAQABAoIBAQCi5Y+dUpNThys7J3ICdHtPwkj1HefgFo+tdkbzu6WShGqMWruHskD+EBki/bpN63xDxd9YgByogSTzTv569+5H9Q9JBFM95z4n0VkAB5po98MGCMwHA9yeT6VKPxTSqj66k1aU4cyUb1o6OyN78Bzt+xKFv1c2l1sKladesSyntJTgNzyCIcTYVhqpY8brQlEPZi0SRun9BR25FhBdKVJl+1NaRQLe0Mh01GbfTMxk1GC8ZlzftJvw3iPlbc9G5Ot0QqNsGTy0yxJ0FT1Fbwk4LBfP2rPNryEGbI3eyze3Ebr5zUiqYNDKIUj6zoZ8HDPOYVoNkAbDnzKOlM5zgXfhAoGBAODR6SScRrcxfYzwFmq7bVJo4xZTJhhhQ/qZbfQV7JFsvTRO+HG0mPJ/DhSbyfDlYr3lhPhRPrDBStxZ/k+D90+JXxd9lnRhngErgNGIMhRuvi6eWz3FzjwKjDeY3FefsHuWiql7oTKerE84Mb6YcCkzCIvuZXT8NE02rgIRBvkfAoGBAMbK7mTDexIVlhJJ/U+JKkTs+svsnqb2YpYTfXXKT0m5ouoJAJcXulqsadHWwhSQEQw5Cd6ff7qw+erGDNkp1pTXOS/UMUkPgZFEH1H+8I+gxoYXkw9GjWla5pctWj6w9rxBIiaFBOUQlzpJp+fVIA3hkJ6By8ixSRY5sn+EypTxAoGBAJzzztkhRPk10tnm7y96Q/sJgKggSnMwzF2SacHC4JIyIPD4xNfU5gY9j13x2QyVh9bc+nTFe4e4mgO1zCZFEg3z8HAc3EHJuJ5GebtOYvAC1EEFittYUf92uadCz2lY8cOGOK3TwpjtT4xKxeey0nLgzvGp8Ci4vea96sNEkeKvAoGAL3Jb99zYUPap/O9/8C3S0uSk72soar0/xoYcWbLRvcz631UKuQVGy0F/eEADzpicGQi8HFCBUfPPkoN2qZZcyuWblpjUspVq3VsfBpkMnZtIBtU1ml3CfrTdeJOwiGoAWepJ1lvFUU7maVkPJDwwDGAb/lkIZHw773wR6HGvUGECgYEAvIesmaaWtcGuvH2vLSa+QYSxx49NFM3VhFEZih8FfzmLmwRa/5oEtyUYxEJhfRBYTQSPQhPA/jNHwS0yhL6s0IoQL2WX7PZDWZ5akgwemzWlToo2DKny3PC6RjPlA2Oivg03nOyCadx2Ua6jW9TA/OZE5agtz+Jod3cSlf28Ddo=")
	privateKey, _ := x509.ParsePKCS1PrivateKey(pkcs8DerKey)

	rng := rand.Reader
	plaintext, _ := rsa.DecryptPKCS1v15(rng, privateKey, ciphertext)

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"encypt_data": string(plaintext),
	}).ParseResponse(w, r)
}

func Encrypt(w http.ResponseWriter, r *http.Request) {

	// privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// if err != nil {
	// 	panic(err)
	// }

	// publicKey := privateKey.PublicKey

	pkcs8DerKey, err := base64.StdEncoding.DecodeString(`TFMwdExTMUNSVWRKVGlCU1UwRWdVRlZDVEVsRElFdEZXUzB0TFMwdENrMUpTVUpKYWtGT1FtZHJjV2hyYVVjNWR6QkNRVkZGUmtGQlQwTkJVVGhCVFVsSlFrTm5TME5CVVVWQk4yMDJaUzlrWVdWS1pVTklUR0UzZW1kVVYza0tlRFJZU0U5TVlVSkNSbmxKZEU1SWIzVm1iVXhhY2pBNWRVVjFhbmhLZG5kMk5UZHNla1YwWlNzNE1XeFhkSEpEYjFGaE1IaGtZbGNyWjI1TlMwOXpZd3BwV0dWR1FtdG1PWFIyTTA1NlIxbFhOV3BEVXl0V2JtSjFNa3h3Ynl0MWEyYzFOalFyVTFkV1NUaEdhV3QyYkU5d04xWlNZa1I2Y1Rjd2VGSkRPRmxTQ2xaWmIwaDZWaTgwUmxNd1YwaHNPVkUyYlc0eFNsQndSR3RQU2tSbU1HUXlSM2RIZEdoak5sUnJORWhtTW13d2VEbHFkRFpuZUVkdk9YZEZiRmRNYlVNS1ZsVjRXWFJrTlVWcVNVbE5WRlZwZEdOSVJYcFdUVko1V0hCRlprdG9hRXMzUWtkdGMzZHNTWFl4TVdoUUt6ZzRPVU41ZG1kMllsVm9WRzg0UlV0d1l3cE1NbFpDUjNRMFpHWnNka1pHYkhWa1EwSmpSbkZSY0VGWVIyVXZWVmw1WVc5Nk4xVlFaVk4zVVV4R1kyVXdhRkoxYml0UmJHNWhObVZzV2t0UVYzcHpDbVYzU1VSQlVVRkNDaTB0TFMwdFJVNUVJRkpUUVNCUVZVSk1TVU1nUzBWWkxTMHRMUzBL`)
	if err != nil {
		log.Println("error disini === ", err)
	}
	var decodedPublicRSAKeyString = string(pkcs8DerKey)
	keyP, err := base64.StdEncoding.DecodeString(decodedPublicRSAKeyString)
	if err != nil {
		fmt.Println("error disini === ", err)
	}
	key, err := crypto.ParseRSAPublicKeyFromPEM(keyP)
	if err != nil {
		fmt.Println("error disini === ", err)
	}
	// key, _ := x509.ParsePKIXPublicKey(pkcs8DerKey)

	// publicKey, _ := key.(*rsa.PublicKey)
	// rng := rand.Reader

	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(`{"privy_id":"UAT001","password":"Akuntes1"}`))
	if err != nil {
		fmt.Println("error when create public.pem: ", err)
	}

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"encypt_data": base64.StdEncoding.EncodeToString(encryptedBytes),
	}).ParseResponse(w, r)
}

func Decrypt(w http.ResponseWriter, r *http.Request) {
	ciphertext, _ := base64.StdEncoding.DecodeString("YWE2YTBmZTJlMjY2ZWE5Mjc3YWQ5ZmEyNDNkN2ViYmJkOWEwZDdhY2IxNDliNGFmMWJmNDc4NzI0MThhMDFjZTlkNmQxMGUzY2M4YTFhZTkxYjUxNGVmNDZmNWYyMjdhZTQ2YTNjZmUzY2RhZDYyMjBkNzU5NzcyNTQ0ZjlhNGJmNDVlNDNiOTkxYWFjNDkwMzM0Zjc2ZDhmMjcwM2RhOTI2MjBkYzZjNDQ2NzU1N2JjZWUwZGE0ZmRlMThiN2E3M2ZhZjQzY2MyY2ZjM2VkYjAwZDY2OGVjYmU4ZWIyN2NlNjNlNDZjMTM2MmNmM2IzZjQ1NWQ3NzUxMmIyNjIwNDdlMWQ0ZjA2NTFiMGQxNDhiOTg0MGJmODcxMjc4ZGFmNjdlYWRlMmNmMzg4ZWE5YzRiMjMwMDU2NWNkYTliMTUzZmI1MjI4ODQ3ZDU1OWMwN2VhN2IyYzU4NTUyODI1MDZjMTlhYWIyNTdmNmZjZDNiYzljOWQxOWM4YjljODcyZTI5ZjYwNzMyMjc4YTI3NmEwNjY2OWVlMzI4NWY4OWE2ZTQ0NTE2NTdhMTJmMmJkYzNkZDU1OTQ2YjI1OTE3NTUxMDcwMTViNGM5ZjlhMTQ3YWUyNjI5ZTdhMjI0NmQ4YjFkZjRiMzEwMDk3NmU1MGE2ZGUzOTIwZTZiM2EwMzM=")

	chip := string(ciphertext)
	stringToDecrypt, err := hex.DecodeString(chip)
	if err != nil {
		fmt.Println("1. error disini === ", err)
	}
	// Import PKCS#8 key
	pkcs8DerKey, _ := base64.StdEncoding.DecodeString("TFMwdExTMUNSVWRKVGlCU1UwRWdVRkpKVmtGVVJTQkxSVmt0TFMwdExRcE5TVWxGY0VGSlFrRkJTME5CVVVWQk4yMDJaUzlrWVdWS1pVTklUR0UzZW1kVVYzbDRORmhJVDB4aFFrSkdlVWwwVGtodmRXWnRURnB5TURsMVJYVnFDbmhLZG5kMk5UZHNla1YwWlNzNE1XeFhkSEpEYjFGaE1IaGtZbGNyWjI1TlMwOXpZMmxZWlVaQ2EyWTVkSFl6VG5wSFdWYzFha05USzFadVluVXlUSEFLYnl0MWEyYzFOalFyVTFkV1NUaEdhV3QyYkU5d04xWlNZa1I2Y1Rjd2VGSkRPRmxTVmxsdlNIcFdMelJHVXpCWFNHdzVVVFp0YmpGS1VIQkVhMDlLUkFwbU1HUXlSM2RIZEdoak5sUnJORWhtTW13d2VEbHFkRFpuZUVkdk9YZEZiRmRNYlVOV1ZYaFpkR1ExUldwSlNVMVVWV2wwWTBoRmVsWk5VbmxZY0VWbUNrdG9hRXMzUWtkdGMzZHNTWFl4TVdoUUt6ZzRPVU41ZG1kMllsVm9WRzg0UlV0d1kwd3lWa0pIZERSa1pteDJSa1pzZFdSRFFtTkdjVkZ3UVZoSFpTOEtWVmw1WVc5Nk4xVlFaVk4zVVV4R1kyVXdhRkoxYml0UmJHNWhObVZzV2t0UVYzcHpaWGRKUkVGUlFVSkJiMGxDUVVKVE5qVllVRzlpZDA5WlF5dHBOZ28xTUhvMmIwUnhZM1JRSzBoVWNIaFZaREJhVWxsamVWcHNOMk4wUTA5eE0wZE9kMFpQZERsUFRESnBaSFpRY0dwc01GcFpPSE5qVWxWRVNreFNXa2RxQ25aSE5YRlZjV0phUkRKWVIyRktZVFphYlZFMVJrSkRZbVJKZG1Sa2FUbExiMFZ0ZVU5VVIwRlpWa00wU21Oc1owUlFVQ3RZVVhCVWVESnhSSFZYUm04S00zcHpTVEEwVkN0YVZVNUVOM04yV1dsaE0wZFdPRU53ZVdKcFdFcHpTbWN6YjNSTk5rbDJSaXR4UlhSWmFqUkhZak16ZFN0bVNIUXdhMGd4UWpGNFlncFhNWHBTUlM5eVJWRmtRMmRxUTJkSGRtSm5SMEpVVmtvNE9FbDVkM0Z3UVVwaE1uVnRha1V3TVV3dmNGaFFibGd5WkUxT1luTjBiRVZIVEVsSFpIWnJDazFQTTNsd2NHUmxSbUkzTlRKRFYybFNVelJpVjFJelZuRkZNMFZQUjBkUVIyRXdjMGhRUmtad2QybzBNRWg0Um5CS1FXbFdRMDlhZWsxRkwweGhZMnNLVlVkaVF6bENSVU5uV1VWQkwyaE9lbFpSYVZKamFGZDNlbWxuTW1rdlFsQmlNRmd2WlRGakwxZFNiRVpOSzJOSFVsVXpXazQxTUU1eE4weFhRbFZpTXdwNFIwbEtTMjV1WkdzdlJGaDJjRTB5T0dWTGVuSXJlVmxxVVRKcGJtcDJhSGRITkhOaU4yOHhTVk55TUZwRlpUaDRZakIwUVU4MGMwcHZhVzl3UzFoSUNqUTVhV3BoS3pNemRHNTBNREl2VGxkblkwOHJRM1pNU1VkMlNqbHRUR2RWVFdWWE5HdFRiamxGZW0weWNXYzJZekJqYUhOQlRXTkRaMWxGUVRoRWVsZ0tOMjkyUkVSVU4wMHllR05zV210b1UyTXdTVE5qUzNka1VqSTJkVEZrTVhreVNIbGtUMVJTY1ZWa2RXOXpaV0kwYjNKeGJIcFhWV3BMYlZGVmRsVlpVQXBLYmxWa0t5OVpWa0ZMU0RkUU5HWmFVbVU0ZEZsTWRXbDJaVW93ZEU5UU9YWlhWbW93UTJkbVkyUXdWREJNSzNKT1RFNVNVVWR2VUU5QlZtczVRVXR1Q20xeFNIaHpTbkY0U0dSRVFtZDJVR3hrYjFoTU9GbHZTemw0U0N0T1JVcE9XazlxZVdGeE1FTm5XVVZCT1dFelJFdFBaRWhHVFZSTGJWSTBhbGRJT1M4S2JpODFWV2N5VTJwSFltMXZkeTgxVW1aMlVTczFWbkpXU0VWRGNWZExTRGs1VVZKRVJISkxWMUZ1V25oa05VMU5XSGx4ZVZNMWRVSkRNblY1Tkd0VVVBcEVNalkxVVhwUlVVdGpiRXRXYVd4VU4xQlVNMHhzTkdacVVXTTFOSHBPTW5KUFoxQmFZemRQV2poTVRFSmxRWFowZG5saGNVTlhTWFo0TTFwMWNraE5DblpWTnk5blJXSlZjM2xHYzJOMFdEaGFOVWhuWkZsTlEyZFpSVUV6Y210YU0wSXJXVGgyVDNoRGMwaFlkVFEzTjFkaU5HWkdaMlZNYTFscmN6aEVka0lLVjB4SFdFNWpla0pqVTJrM2NXdDFabGxaWVROYWMyZEpTbE51Y1VSR2QybGxXalJGT0ZjMlFYQkVhVWxpWVhKMGNHRmxURzh5TTNsQlEyTjBNbVpMZUFwalQwUktLemxpYlRaRVNqRlNOR2RpUjNRelVWTXlOMGx6TUVKV1EzVkNXRTFOV2twQk1Xc3pUVnBSZW1jNGRrbG9aWGxtVjJ0NVNIY3JNMkpvWlcxbENsSlNjWFZtU0ZWRFoxbEJaVlJuT1ZoQldIZ3lTMnBUT1M5bU5uSTRjRzlaTkVkRU9VTjBRVmg2U25wVVpETlNTMGRYYVhnM01UbExXbmRxZUU0dmJsY0tTRmx2WkhBMU5saEVVRU0zUW1kblNXTkJTekZxWlhsdVN5OVFkMkV6UVd3dlMwdFNTREpoWkhjNE5VUTNZbk56YlVjeldXOU1WRFZDWjBsSFoxRTJXZ292TVRkMlJTdDJVVFJ5Y25ZMVoxVlZjakYyVTFRdlpqVXpOV2t6T1dod2VFdG5OazFoUkZCM2JVdGtibTlxTm5KRlZuRjRjRkU5UFFvdExTMHRMVVZPUkNCU1UwRWdVRkpKVmtGVVJTQkxSVmt0TFMwdExRbz0=")
	var decodedPrivateRSAKeyString = string(pkcs8DerKey)
	keyP, err := base64.StdEncoding.DecodeString(decodedPrivateRSAKeyString)
	if err != nil {
		fmt.Println("2. error disini === ", err)
	}
	privateKey, err := crypto.ParseRSAPrivateKeyFromPEM(keyP)
	if err != nil {
		fmt.Println("3. error disini === ", err)
	}

	rng := rand.Reader
	plaintext, _ := rsa.DecryptPKCS1v15(rng, privateKey, stringToDecrypt)

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"encypt_data": string(plaintext),
	}).ParseResponse(w, r)
}
