package main

import (
	"github.com/dgrijalva/jwt-go"
)

type  MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//func main(){
//	my:=&MyClaims{"ycy",
//	jwt.StandardClaims{
//		NotBefore: time.Now().Unix(),
//		ExpiresAt: time.Now().Unix()+ 60*60,
//		Issuer: "ycy",
//	}}
//	tocken:=jwt.NewWithClaims(jwt.SigningMethodHS256,my)
//	strdata,err:=tocken.SignedString([]byte("the secert"))
//	if err!=nil{
//		log.Println(err)
//	}
//	fmt.Println(strdata)
//	prasetocker,err:=jwt.ParseWithClaims(strdata,my,func(token *jwt.Token) (interface{}, error) {
//		return []byte("the secert"), nil
//	})
//	if err!=nil{
//		log.Println(err)
//	}
//	if p,ok:=prasetocker.Claims.(*MyClaims);ok{
//		fmt.Println(p.Username )
//		exptime:=time.Unix(p.ExpiresAt,0)
//		fmt.Println(exptime.Format(time.RFC3339))
//	}
//
//
//
//
//
//}