package utils

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	DIRECTORY_REPORTS = "tmp/uploads/reports"
	DIRECTORY_UPLOADS = "tmp/uploads/id-cards"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	os.MkdirAll(DIRECTORY_REPORTS, os.ModePerm)
	os.MkdirAll(DIRECTORY_UPLOADS, os.ModePerm)
}

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder

	l := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(5)
}

func RandomMoney() int64 {
	return RandomInt(1000, 100000)
}

func RandomCurrency() string {
	var currencies []string = []string{"RUB", "USD", "CAD", "EUR", "IDR"}
	return currencies[rand.Intn(len(currencies))]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomFileName(file *multipart.FileHeader) string {
	fileName := filepath.Base(file.Filename)
	fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	return fmt.Sprintf("%s-%d-%s", fileNameWithoutExt, time.Now().UnixNano(), RandomString(8)+filepath.Ext(file.Filename))
}
