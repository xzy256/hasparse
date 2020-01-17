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
	msgString := "a4IDzzCCA8ugAwIBBaEDAgELoxUbE1NVTlNISU5FLkhBRE9PUC5DT02kFDASoAMCAQGhCzAJGwd3ZWJfbmxwpYICl2GCApMwggKPoAMCAQWhFRsTU1VOU0hJTkUuSEFET09QLkNPTaIoMCagAwIBAqEfMB0bBmtyYnRndBsTU1VOU0hJTkUuSEFET09QLkNPTaOCAkUwggJBoAMCARGhAwIBAaKCAjMEggIv5gCeUj8n2Grj0sRFQvpGn9rWpdOzF/s2Qc3Gj9kZs7WGUVP9n+sfAmjaBmL6SjyyG6KCsQrXsoArmZLqF4mt07KkL0xn/whj/jT0oSsDVgRfpjG47q1X8DyjMfIh4rANyMZwMf3EB/NifLjZfWQZRYB14UXqNMi+/azq47dMx9FTH8n92KE4TFWR3bko8z/BiyjRZoazSj9CJWOmj0Qd1sf3dKKKugPphvU/yemr6SifGeiqbJgMfvnu2Odc30ZLSCvZ9B8DgRzsIKPiCk395cRs+1xNI3GYIgn1r9mvSCMRllxGr0rqlgdF+xQ7o3hHuy/sxqgmRfXsCipGzQKNs8RVggcuXCdZ8c9+XlU7B4msm5kCGrS83EY/ut+vQnEkhRitj8NcAtBGA6AtKqqYyj/Dg6qnxVLLh+dGwxFuV6WdNOeBFDIzGK5R37ejYu7154CeeTvITWJfVtTUkpNIYaoM7jkFGjt05vu3bRREDgeQci6ivlXVTDLUseaB403VwzRHY8K+zKFZ38Bqrp8kDE/9FQUzk+/s/1obrIcfpGAkkSiwY0auerz/ref+ajSIOYXcggD7Ruae5LuOPPuV91flESvegtuQ3gxB3icc9SSU8/1XS/ERpz6Ej6NbFZB6DXwspc3vh83LuYbK8kY1t2EYfN0hEQa3+z4goTorvEAk/E7wYbty8/KLDGS6P53PP0HmUUdN2LXA6WxywqwvzPz4h58aiKYsw/dak4h30qaB9jCB86ADAgERooHrBIHoxKd5F1BSaVbP6SUGOQwVgfuLsyKNJYIZGNtEF9MjnobzgAUiVfkSZLKgPgQ/rbfWLElUUdqjmy8KmwlY/pJLd4yaYgoYEzdaUeVVlmyNkrHfz4f9D7E75q4gsGhvAbSjoGc6tTPLzKqZ20HhYHv8AAfrDRPGyiZYgmdfBNoDPRxhTcTcQcAh7mdZy3fD0k24rmrK7+vO6ruK2duU9WUjwFIViloTUzV54by+8NPB+5UBvOi0ntspMHYBwr4IJ2Lux1b7Y7NofG1CtZHkNe43I9z41Zlj74zkvLjeeue+tVdWHwNDtIrBkw=="
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
