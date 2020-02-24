package assign

import (
	"bytes"
	"hasparse/hasauth"
	"hasparse/unmarshal"
	"log"
)

func HasKdc(userName, password, authSeverAddr, port string) *KdcRep{

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
		log.Println("Response->krbMessage decode fail by base64")
	}

	buf := bytes.NewBuffer(krbMessage)
	s1 := unmarshal.Asn1ParserBuffer(*buf)
	asRep := &KdcRep{}
	asRep.Init()
	asRep.Decode(s1)
	HandleKdcRep(asRep, userName+password) //"guestguestpassword0"
	asRep.Display()

	return asRep
}
