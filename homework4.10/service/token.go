package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type Payload struct {
	Iss string `json:"iss"`
	Exp string `json:"exp"`
	Iat string `json:"iat"`
	Id  uint   `json:"id"`
	Username  string   `json:"username"`
}
type JWT struct {
	Header    Header
	Payload   Payload
	Signature string
	Token     string
}

func NewJWT(name string) JWT {
	var jwt JWT
	jwt.Header = Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	jwt.Payload = Payload{
		Iss: "chengyongle",
		Exp: strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),
		Iat: strconv.FormatInt(time.Now().Unix(), 10),
		Username:  name,
	}
	h, _ := json.Marshal(jwt.Header)
	p, _ := json.Marshal(jwt.Payload)
	baseh := base64.StdEncoding.EncodeToString(h)
	basep := base64.StdEncoding.EncodeToString(p)
	secret := baseh + "." + basep
	key := "jinwandalaohu"
	mac := hmac.New(sha256.New,[]byte(key))
	mac.Write([]byte(secret))
	s := mac.Sum(nil)
	jwt.Signature = base64.StdEncoding.EncodeToString(s)
	jwt.Token = secret + "." + jwt.Signature
	return jwt
}
func Check(token string) (jwt JWT, err error) {
	err = errors.New("token error")
	arr := strings.Split(token, ".")
	if len(arr) < 3 {
		fmt.Println("59------", err)
		return
	}
	//header里前七个字符为"bearer ",有一个空格，base64识别不了，直接把前七个字符截掉
	baseh := arr[0][7:]
	h, err := base64.StdEncoding.DecodeString(baseh)
	if err != nil {
		fmt.Println("decode header", err)
		return
	}
	err = json.Unmarshal(h, &jwt.Header)
	if err != nil {
		fmt.Println("unmarshal header", err)
		return
	}
	basep := arr[1]
	p, err := base64.StdEncoding.DecodeString(basep)
	if err != nil {
		fmt.Println("decode payload", err)
		return
	}
	err = json.Unmarshal(p, &jwt.Payload)
	if err != nil {
		fmt.Println("unmarshal payload", err)
		return
	}
	bases := arr[2]
	s1, err := base64.StdEncoding.DecodeString(bases)
	if err != nil {
		fmt.Println("decode bases", err)
		return
	}
	se := baseh + "." + basep
	w := []byte("jinwandalaohu")
	mac := hmac.New(sha256.New, w)
	mac.Write([]byte(se))
	s2 := mac.Sum(nil)
	if string(s1) != string(s2) {
		return
	} else {
		jwt.Signature = arr[2]
		jwt.Token = token
	}
	return jwt, nil
}