package nasmshell

import (
	"FoxShell/lib/assembler"
	"FoxShell/lib/headers"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var Module headers.Module

func Run() bool {
	Module.Name = "NasmShell"
	Module.Description = "The updated version of metasploit's nasm-shell."
	Module.Arch = headers.X64
	Module.Author = []string{"xHector1337"}
	Module.OperatingSystem = headers.Multi
	Module.PrintModuleInfo()

	var input string
	var Arch = headers.X64
	var OperatingSystem string
	if runtime.GOOS == "windows" {
		OperatingSystem = headers.Windows
	} else {
		OperatingSystem = headers.Linux
	}
	for {
		var output []byte
		fmt.Printf("nasmshell> ")
		var scanner = bufio.NewReader(os.Stdin)
		input, _ = scanner.ReadString('\n')
		if input == "exit\n" || input == "exit\r\n" {
			break
		}
		if input == "arch64\n" || input == "arch64\r\n" {
			Arch = headers.X64
			fmt.Printf("[+] Setting architecture to x64...\n")
		} else if input == "arch32\n" || input == "arch32\r\n" {
			Arch = headers.X86
			fmt.Printf("[+] Setting architecture to x86...\n")
		}
		if input == "windows\n" || input == "windows\r\n" {
			fmt.Printf("[+] Setting operating system to Windows...\n")
			OperatingSystem = headers.Windows
		} else if input == "linux\n" || input == "linux\r\n" {
			fmt.Printf("[+] Setting operating system to Linux...\n")
			OperatingSystem = headers.Linux
		}
		if strings.HasPrefix(input, "{\n") || strings.HasPrefix(input, "{\r\n") { // multi line support
			for !strings.HasSuffix(input, "}\n") && !strings.HasSuffix(input, "}\r\n") {
				var tempInput, _ = scanner.ReadString('\n') // fuck, go didn't let me do input += scanner.ReadString('\n')
				input += tempInput
			}
			if runtime.GOOS == "windows" {
				input = strings.ReplaceAll(input, "{\r\n", "")
				input = strings.ReplaceAll(input, "}\r\n", "")
			} else {
				input = strings.ReplaceAll(input, "{\n", "")
				input = strings.ReplaceAll(input, "}\n", "")
			}
		}
		assembler.Assemble(OperatingSystem, Arch, input, &output)
		for i := 0; i < len(output); i++ {
			fmt.Printf("%02x ", output[i])
		}
		fmt.Printf("\n")
	}
	return true

}
