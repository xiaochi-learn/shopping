package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
)

func genMd5(code string) string {
	hasher := md5.New()
	_, _ = io.WriteString(hasher, code)
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	//dsn := "root:tianchi123@tcp(10.25.77.219:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//	logger.Config{
	//		SlowThreshold: time.Second,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	},
	//)
	//
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = db.AutoMigrate(&model.User{})
	//if err != nil {
	//	panic(err)
	//}

	// Using custom options
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	fmt.Println("salt1:", salt)
	fmt.Println("encodedPwd1:", encodedPwd)
	newPass := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println("len newPass:", len(newPass))
	fmt.Println("password:", newPass)
	passInfo := strings.Split(newPass, "$")
	check := password.Verify("generic password", passInfo[2], passInfo[3], options)
	fmt.Println("check1:", check) // true
}
