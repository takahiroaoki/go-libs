package string

// PadStart inserts padStr before orgStr until its length reaches to untilLen.
func PadStart(orgStr, padStr string, untilLen int) string {
	return calcPadding(orgStr, padStr, untilLen) + orgStr
}

// PadStart inserts padStr after orgStr until its length reaches to untilLen.
func PadEnd(orgStr, padStr string, untilLen int) string {
	return orgStr + calcPadding(orgStr, padStr, untilLen)
}

func calcPadding(orgStr, padStr string, untilLen int) string {
	padding := ""
	for len(padding)+len(orgStr) <= untilLen-len(padStr) {
		padding += padStr
	}
	padding += padStr[:untilLen-len(padding)-len(orgStr)]
	return padding
}
