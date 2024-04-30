package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

func PrintJsonStruct(v any) string {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(jsonBytes)
}

func FormatHex(data []byte) string {
	var builder strings.Builder
	builder.Grow(len(data) * 4)
	builder.WriteString(fmt.Sprintf("\r\n"))
	count := 0
	pass := 1
	for _, b := range data {
		if count == 0 {
			builder.WriteString(fmt.Sprintf("%-6s\t", "["+fmt.Sprint((pass-1)*16)+"]"))
		}

		count++
		builder.WriteString(fmt.Sprintf("%02X", b))
		if count == 4 || count == 8 || count == 12 {
			builder.WriteString(" ")
		}
		if count == 16 {
			builder.WriteString("\t")
			for i := (pass * count) - 16; i < (pass * count); i++ {
				c := rune(data[i])
				if c > 0x1f && c < 0x80 {
					builder.WriteRune(c)
				} else {
					builder.WriteString(".")
				}
			}
			builder.WriteString("\r\n")
			count = 0
			pass++
		}
	}
	builder.WriteString("\r\n")
	builder.WriteString("\r\n")
	return builder.String()
}

func ParseIPFromInt32(ipInteger uint32) string {
	ip := make(net.IP, net.IPv4len)
	binary.LittleEndian.PutUint32(ip, ipInteger)
	return ip.String()
}
