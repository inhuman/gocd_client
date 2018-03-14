package utils

import (
	"net/http"
	"github.com/hokaccha/go-prettyjson"
	"log"
	"fmt"
	"strings"
	"regexp"
	"os"
)

func IsHeaderExists(r *http.Request, header string) bool {

	for key := range r.Header {
		if key == header {
			return true
		}
	}

	return false
}

func PrettyPrintStruct(strct interface{}) {

	s, _ := prettyjson.Marshal(strct)
	fmt.Println(string(s))
}

func CheckError(err error) {
	if err != nil {
		log.Println("error:", err)
	}
}

func DebugMessage(string string) {
	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println(string)
	}
}

func InArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func InArrayRegexp(string string, regexpArr []string) bool {

	for _, regex := range regexpArr {

		var re = regexp.MustCompile(regex)

		if re.Match([]byte(string)) {
			return true
		}
	}

	return false
}

func ParseDcService(serviceName string) (string, string) {

	stringArr := strings.Split(serviceName, "[")
	stringArr2 := strings.Split(stringArr[1], "]")

	return stringArr2[0], stringArr2[1]
}
