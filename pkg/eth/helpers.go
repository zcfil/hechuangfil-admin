package eth

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
)

// 将十六进制字符串值解析为int
func ParseInt(value string) (int, error) {
	i, err := strconv.ParseInt(strings.TrimPrefix(value, "0x"), 16, 64)
	if err != nil {
		log.Panicln("err:", err.Error())
		return 0, err
	}

	return int(i), nil
}

// 将十六进制字符串值解析为big.Int
func ParseBigInt(value string) (big.Int, error) {
	i := big.Int{}
	_, err := fmt.Sscan(value, &i)

	return i, err
}

// 将int转换为十六进制表示
func IntToHex(i int) string {
	return fmt.Sprintf("0x%x", i)
}

// BigToHex covert big.Int to hexadecimal representation
//到十六进制 长整型到十六进制表示
func BigToHex(bigInt big.Int) string {
	if bigInt.BitLen() == 0 {
		return "0x0"
	}

	return "0x" + strings.TrimPrefix(fmt.Sprintf("%x", bigInt.Bytes()), "0")
}
