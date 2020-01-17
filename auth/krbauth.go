package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	HTTPSSCHEMA = "https://"
	HTTPSCHEMA  = "http://"
)

type MapClaims map[string]interface{}

type Response struct {
	KrbMessage string `json:"krbMessage"`
	Success    string `json:"success"`
	Type       string `json:"type"`
}

// request tgt from KDC
func requestTgt(hasConfig *HasConfig, tokenString string) *Response {

	tr := &http.Transport{
		MaxIdleConns:       0,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig: &tls.Config{

		},
	}
	client := &http.Client{Transport: tr}

	response := &Response{}
	for _, kdcServerAddr := range strings.Split(hasConfig.HttpsHost, ",") {
		urlAddress := HTTPSCHEMA + kdcServerAddr + ":" + hasConfig.HttpsPort +
			"/has/v1?type=" + hasConfig.Type + "&authToken=" + tokenString
		u, err := url.Parse(urlAddress)

		if err != nil {
			continue
		}

		resp, err := client.Do(&http.Request{Method: "PUT", URL: u})
		if err != nil {
			continue
		}
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			continue
		}
		err = json.Unmarshal(bodyBytes, response)
		if err == nil {
			break
		}
		resp.Body.Close()
	}

	return response
}

// encode token by base64
func encodeTokenWithBase64(hasConfig *HasConfig) string {
	head := `{"alg":"none"}`
	head64 := strings.Trim(base64.StdEncoding.EncodeToString([]byte(head)), "=") // trim end equal flag

	uid, _ := uuid.GenerateUUID()
	now := time.Now()
	claims := MapClaims{
		"iss":        "has",
		"exp":        now.Add(time.Hour * time.Duration(1)).Unix(),
		"iat":        now.Unix(),
		"sub":        hasConfig.UserName,
		"secret":     hasConfig.Password, // verify password
		"user":       hasConfig.UserName, // verify userName
		"passPhrase": hasConfig.UserName + hasConfig.Password,
		"jti":        uid,
	}
	jsonStr, err := json.Marshal(claims)

	if err != nil {
		return ""
	}

	payload1Base64 := base64.StdEncoding.EncodeToString([]byte(string(jsonStr)))
	return head64 + "." + payload1Base64 + "."
}

// parse response
func getKrbMessageWithBase64(resp *Response, rtype string) ([]byte, error) {
	if resp.Success != "true" {
		return nil, errors.New("request TGT failed, response->success is failed")
	}
	if resp.Type == "" || resp.Type != rtype {
		return nil, errors.New("request TGT failed, response->type is empty or not matched with req")
	}
	msgBytes := []byte(resp.KrbMessage)
	krbMessage, err := base64.StdEncoding.DecodeString(string(msgBytes))
	if err != nil {
		return nil, errors.New("request TGT failed, response->krbMessage decode fail by base64")
	}

	return krbMessage, nil
}



func base64DecodeStripped(s string) ([]byte, error) {
	if i := len(s) % 4; i != 0 {
		s += strings.Repeat("=", 4-i)
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	return decoded, err
}

func Test() {

	hasConfig := &HasConfig{
		HttpsPort: "8091",
		HttpsHost: "rsync.master009.sunshine.hadoop.js.ted,rsync.master010.sunshine.hadoop.js.ted",
		UserName:  "web_nlp",
		Password:  "4e36cd9bb66c3039aca29db66b38fcae",
		Type:      "ConfigFile",
	}

	tokenString := encodeTokenWithBase64(hasConfig)
	fmt.Println("encode token:", tokenString)
	var response *Response
	response = requestTgt(hasConfig, tokenString)
/*	response = &Response{
		"a4IDzzCCA8ugAwIBBaEDAgELoxUbE1NVTlNISU5FLkhBRE9PUC5DT02kFDASoAMCAQGhCzAJGwd3ZWJfbmxwpYICl2GCApMwggKPoAMCAQWhFRsTU1VOU0hJTkUuSEFET09QLkNPTaIoMCagAwIBAqEfMB0bBmtyYnRndBsTU1VOU0hJTkUuSEFET09QLkNPTaOCAkUwggJBoAMCARGhAwIBAaKCAjMEggIv5gCeUj8n2Grj0sRFQvpGn9rWpdOzF/s2Qc3Gj9kZs7WGUVP9n+sfAmjaBmL6SjyyG6KCsQrXsoArmZLqF4mt07KkL0xn/whj/jT0oSsDVgRfpjG47q1X8DyjMfIh4rANyMZwMf3EB/NifLjZfWQZRYB14UXqNMi+/azq47dMx9FTH8n92KE4TFWR3bko8z/BiyjRZoazSj9CJWOmj0Qd1sf3dKKKugPphvU/yemr6SifGeiqbJgMfvnu2Odc30ZLSCvZ9B8DgRzsIKPiCk395cRs+1xNI3GYIgn1r9mvSCMRllxGr0rqlgdF+xQ7o3hHuy/sxqgmRfXsCipGzQKNs8RVggcuXCdZ8c9+XlU7B4msm5kCGrS83EY/ut+vQnEkhRitj8NcAtBGA6AtKqqYyj/Dg6qnxVLLh+dGwxFuV6WdNOeBFDIzGK5R37ejYu7154CeeTvITWJfVtTUkpNIYaoM7jkFGjt05vu3bRREDgeQci6ivlXVTDLUseaB403VwzRHY8K+zKFZ38Bqrp8kDE/9FQUzk+/s/1obrIcfpGAkkSiwY0auerz/ref+ajSIOYXcggD7Ruae5LuOPPuV91flESvegtuQ3gxB3icc9SSU8/1XS/ERpz6Ej6NbFZB6DXwspc3vh83LuYbK8kY1t2EYfN0hEQa3+z4goTorvEAk/E7wYbty8/KLDGS6P53PP0HmUUdN2LXA6WxywqwvzPz4h58aiKYsw/dak4h30qaB9jCB86ADAgERooHrBIHoxKd5F1BSaVbP6SUGOQwVgfuLsyKNJYIZGNtEF9MjnobzgAUiVfkSZLKgPgQ/rbfWLElUUdqjmy8KmwlY/pJLd4yaYgoYEzdaUeVVlmyNkrHfz4f9D7E75q4gsGhvAbSjoGc6tTPLzKqZ20HhYHv8AAfrDRPGyiZYgmdfBNoDPRxhTcTcQcAh7mdZy3fD0k24rmrK7+vO6ruK2duU9WUjwFIViloTUzV54by+8NPB+5UBvOi0ntspMHYBwr4IJ2Lux1b7Y7NofG1CtZHkNe43I9z41Zlj74zkvLjeeue+tVdWHwNDtIrBkw==",
		"true",
		"ConfigFile",
	}*/
	fmt.Println("response=", response)
	getKrbMessageWithBase64(response, "ConfigFile")


}
