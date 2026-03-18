package formatter

import (
	"encoding/base64"
	"fmt"
)

func StyleC(shellcode []byte) {
	fmt.Printf("unsigned char shellcode[] = \n\"")
	for i := 0; i < len(shellcode); i++ {
		if i%16 == 0 && i != 0 {
			fmt.Print("\"\n")
			fmt.Printf("\"")
		}
		if i == len(shellcode)-1 {
			fmt.Printf("\\x%02x\";", shellcode[i])
		} else {
			fmt.Printf("\\x%02x", shellcode[i])
		}
	}
	fmt.Printf("\n")
}

func StyleGO(shellcode []byte) {
	fmt.Printf("shellcode := []byte{")
	for i := 0; i < len(shellcode); i++ {
		if i%16 == 0 && i != 0 {
			fmt.Printf("\n")
		}
		if i == len(shellcode)-1 {
			fmt.Printf("0x%02x}", shellcode[i])
		} else {
			fmt.Printf("0x%02x,", shellcode[i])
		}
	}
	fmt.Printf("\n")
}

func Base64(shellcode []byte) {
	var s = base64.StdEncoding.EncodeToString(shellcode)
	fmt.Printf("%s\n", s)
}
