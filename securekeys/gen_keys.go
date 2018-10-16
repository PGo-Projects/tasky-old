package main

import (
	"encoding/base64"
	"fmt"

	"github.com/gorilla/securecookie"
)

func main() {
	key := securecookie.GenerateRandomKey(64)
	if key == nil {
		panic("Key not generated sucessfully")
	}
	fmt.Println(base64.StdEncoding.EncodeToString(key))
	key = securecookie.GenerateRandomKey(64)
	if key == nil {
		panic("Key not generated sucessfully")
	}
	fmt.Println(base64.StdEncoding.EncodeToString(key))
	key = securecookie.GenerateRandomKey(32)
	if key == nil {
		panic("Key not generated sucessfully")
	}
	fmt.Println(base64.StdEncoding.EncodeToString(key))
}
