package main

type KrbMessageType int

const (
	NONE      KrbMessageType = -1
	AS_REQ    KrbMessageType = 10
	AS_REP    KrbMessageType = 11
	TGS_REQ   KrbMessageType = 12
	TGS_REP   KrbMessageType = 13
	AP_REQ    KrbMessageType = 14
	AP_REP    KrbMessageType = 15
	KRB_SAFE  KrbMessageType = 20
	KRB_PRIV  KrbMessageType = 21
	KRB_CRED  KrbMessageType = 22
	KRB_ERROR KrbMessageType = 30
)

func KrbMessageTypeFromValue(value int) int {
	for i := AS_REQ; i <= KRB_ERROR; i++ {
		if int(i) == value {
			return int(i)
		}
	}

	return int(NONE)
}
