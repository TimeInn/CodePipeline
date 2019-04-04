package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

const tokenFile = ".token"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetToken() string {
	b, err := ioutil.ReadFile(".token")
	if err != nil {
		fmt.Println(".token file not found")
		return ""
	}

	s := string(b)

	return strings.Replace(s, "\n", "", -1)
}

func SaveToken() string {
	token := RandStringByte(64)
	err := ioutil.WriteFile(tokenFile, token, 0666)
	check(err)

	return string(token)
}

/**
 * Check file is exist
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**
 * rand token
 */
func RandStringBytesMaskImpr(n int64) string {
	return string(RandStringByte(n))
}

func RandStringByte(length int64) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := int64(0); i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}
