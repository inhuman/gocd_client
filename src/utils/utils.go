package utils

import (
	"net/http"
	"github.com/hokaccha/go-prettyjson"
	"log"
	"errors"
	"net"
	"fmt"
	"strings"
	"encoding/json"
	"regexp"
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

func CheckErrorHttp(err error, w http.ResponseWriter, code int) bool {

	if err != nil {
		var er HttpErrorMessage
		er.Error = err.Error()

		jsn, errr := json.Marshal(er)
		CheckError(errr)

		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(code)
		w.Write([]byte(jsn))
		return true
	}
	return false
}
type HttpErrorMessage struct {
	Error string `json:"error"`
}


func CheckError(err error) {
	if err != nil {
		log.Println("error:", err)
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

	stringArr  := strings.Split(serviceName, "[")
	stringArr2 := strings.Split(stringArr[1], "]")

	return stringArr2[0], stringArr2[1]
}

func ExternalIP() (string, error) {
	ifcs, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifcs {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
