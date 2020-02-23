package hasauth

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
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
func RequestTgt(hasConfig *HasConfig, tokenString string) *Response {

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
func EncodeTokenWithBase64(hasConfig *HasConfig) string {
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
func GetKrbMessageWithBase64(resp *Response, rtype string) ([]byte, error) {
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



func Base64DecodeStripped(s string) ([]byte, error) {
	if i := len(s) % 4; i != 0 {
		s += strings.Repeat("=", 4-i)
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	return decoded, err
}
