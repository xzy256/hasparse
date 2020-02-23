# hasparse

用于解析 以Hasclient-Hasserver方式进行 Kerberos 认证中编码为 BER 情况的工具类，只实现了 EncryptedData.Cipher 以 Aes128 加密情况。
涉及票据的获取、BER 解码、Aes128解码等

入口类:  
`func HasKdc(userName, password, authSeverAddr, port string) *assign.KdcRep`  

KdcRep 结构如下, EncPart 是解密后最终的结果
```
type KdcRep struct {
	Pvno        int
	MsgType     int
	Padata      *PAData //optional
	Crealm      string
	Cname       *PrincipalName
	Ticket      *KdcTicket
	EncData     *EncryptedData
	fieldInfos  []int
	position    []byte 
	taggingList *list.List
	EncPart     *EncKdcRepPart
}
```

# TODO  
- kdcRep解析改成EncPart方式
- 部分条件例外分支尚未实现
- Aes128解码中未进行 checksum 校验

# concat
申明：个人能力有限，时间有限，通用性上不够完善  
意见或建议，email: 644351021@qq.com
