package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"hasparse/assign"
	"hasparse/unmarshal"
	"log"
)

func main(){
	msgString := "a4IDzzCCA8ugAwIBBaEDAgELoxUbE1NVTlNISU5FLkhBRE9PUC5DT02kFDASoAMCAQGhCzAJGwd3ZWJfbmxwpYICl2GCApMwggKPoAMCAQWhFRsTU1VOU0hJTkUuSEFET09QLkNPTaIoMCagAwIBAqEfMB0bBmtyYnRndBsTU1VOU0hJTkUuSEFET09QLkNPTaOCAkUwggJBoAMCARGhAwIBAaKCAjMEggIvALBKBxiD7pIHQxP5UZx4OizSpK7k1l2O6Hk81JausKyRmh43OIunqXL4RTfSZ46i4XxNztKKh/UAywrUIFWt8FlSOCQWtx74hRFfxdF2qcjma+rvfitrw3xqCBJ8oOQKeLk4kV9AwOCHb0QvKdw5kBIJ+68oxqHqkFImcenLDG/zLksm3gkv8ZtR3N57RtLd2IG+MwslKsxtYGGQO8npo4imM12A9zdljNeFn/Q1HH+vTMXvXScKmUAm/h7ujfUgxNK8x7eYM0XPnQUmzFGsc08Wb8S73Qbs5j8Yyghj6Dh6WrD8fvYy131pY2YyXK/S+w3vLA08+YW31M9R2QCFa8Y7/wKcZdQ+3iCarH3IkNK34hsX5GkM4s+sD7JNQC+6fNBr4MHwvS5Yg8VS92TUubLuofXyN0JmqAFVME0JyAdNf0NbgW77JiL0O1LQHkh81KdQKCTfy0A7KAccucYsaqXasFvt+S1j2y9NMWFsysLEJRcb4EWvov/HeNf84+djLRO0zjU3rcgme/nQCUnrdXWpMvea4JQq8JyC0q4JW+uAXKPDFo9kuNdOJCsm+9XKmbifLXd33gx1pGUL1L6xY4i58uyVV40/dWEffyvPXKbMKwFowLpIH59SRTwiOEY+G1fDFkG0feHid66t1xTa3nP8XcZa8CZejmU4SyWVdoA1Z7fYAjgZEpHIz7jH/K3VufPRESqAKXPLSDwixB0zXEGjHGJUca/ZcKNGyM2L0KaB9jCB86ADAgERooHrBIHoKfphOGtCpfNPstWObMZlZZ+qMGX9nu+NNuge7jp+/dy8qQRMBg+bxrTUTiBIzflw8huueyexSaE1dxjXzXqC3kk+ZGCAmm+Q3fbkg7frYOBHAbTJJrPU36/9Ghoz5agNldXDMvOgHgwDm2afUKNR+Gdsi/AuvpIzx//jn8/9QyO9Ezr3Z51Y0N/bMCMKDGa0dkySN1rPsIfipSpJj3WT+epwdhr1QYRIX7VSwO/WhYP6EEkp1vg/pj1w/K8z+NAi5SQ2BwG/r4r7rL9y8m5xAHEyd2lDzYfUs4gt40ai+ZedLFc+ySEt3A=="
	krbMessage, err := base64.StdEncoding.DecodeString(msgString)
	fmt.Println(krbMessage)
	if err != nil {
		log.Fatal(errors.New("Response->krbMessage decode fail by base64"))
	}
	buf := bytes.NewBuffer(krbMessage)
	s1 := unmarshal.Asn1ParserBuffer(*buf)
	asRep := &assign.KdcRep{}
	asRep.Init()
	asRep.Decode(s1)
	asRep.Display()
}
