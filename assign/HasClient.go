package assign

import (
	"bytes"
	"github.com/xzy256/hasparse/hasauth"
	"github.com/xzy256/hasparse/unmarshal"
	"log"
)

func HasKdc(userName, password, authSeverAddr, port string) (*KdcRep, error) {

	hasConfig := &hasauth.HasConfig{
		HttpsPort: port,
		HttpsHost: authSeverAddr,
		UserName:  userName,
		Password:  password,
		Type:      "ConfigFile",
	}

	tokenString := hasauth.EncodeTokenWithBase64(hasConfig)
	var response *hasauth.Response
	response = hasauth.RequestTgt(hasConfig, tokenString)

	krbMessage, err := hasauth.GetKrbMessageWithBase64(response, "ConfigFile")
	if err != nil {
		log.Printf("Response->krbMessage decode fail by base64, err: %v", err)
		return nil, err
	}

	buf := bytes.NewBuffer(krbMessage)
	s1 := unmarshal.Asn1ParserBuffer(*buf)
	asRep := &KdcRep{}
	asRep.Init()
	asRep.Decode(s1)
	HandleKdcRep(asRep, userName+password) //"guestguestpassword0"
	//asRep.Display()

	return asRep, nil
}
