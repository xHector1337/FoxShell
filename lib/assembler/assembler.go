package assembler

import (
	"FoxShell/lib/headers"
	"debug/elf"
	"debug/pe"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Assemble(OperatingSystem string, Architecture int, shellcode string, Output *[]byte) bool {
	var temp, e0 = os.Create("tempcode")
	if checker := errorCheck(e0); checker != true {
		return checker
	}
	var writtenBytes, e1 = temp.Write([]byte(shellcode))
	if checker := errorCheck(e1); checker != true && writtenBytes != len(shellcode) {
		return checker
	}
	if currentOS := runtime.GOOS; currentOS == "windows" {
		return windowsAssemble(&OperatingSystem, &Architecture, temp, Output)
	} else if currentOS == "linux" {
		return linuxAssemble(&OperatingSystem, &Architecture, temp, Output)
	}
	fmt.Printf("[-] FoxyShell doesn't work on your current operating system.\n")
	return false
}

func linuxAssemble(OS *string, Architecture *int, file *os.File, Output *[]byte) bool { // This function does assembling on Linux Systems

	if check := errorCheck(file.Close()); check != true {
		return check
	}

	var format string
	if *OS == headers.Windows {
		format = "win"
	} else if *OS == headers.Linux {
		format = "elf"
	} else {
		fmt.Printf("[-] Unknown operating system value given to Assembler!\nValue: %s\n", *OS)
		return false
	}
	if *Architecture == headers.X64 {
		format += "64"
	} else if *Architecture == headers.X86 {
		format += "32"
	} else {
		fmt.Printf("[-] Unknown Architecture value given to Assembler!\nValue: %d\n", *Architecture)
		return false
	}

	var cmd = exec.Command("nasm", "-f", format, "-o", "o.o", file.Name())
	var output, e0 = cmd.CombinedOutput()
	if check := errorCheck(e0); check != true {
		fmt.Printf("[-] Nasm output: %s\n", output)
		return check
	}
	if check := errorCheck(os.Remove(file.Name())); check != true {
		return check
	}
	if *OS == headers.Windows {
		var peFile, e1 = pe.Open("o.o")
		if check := errorCheck(e1); check != true {
			return check
		}
		var text = peFile.Section(".text")
		if text == nil {
			peFile.Close()
			os.Remove("o.o")
			return false
		}
		var textData, e2 = text.Data()
		if check := errorCheck(e2); check != true {
			return check
		}
		*Output = textData
		if check := errorCheck(peFile.Close()); check != true {
			return check
		}
		if check := errorCheck(os.Remove("o.o")); check != true {
			return check
		}
		return true
	}
	var elfFile, e1 = elf.Open("o.o")
	if check := errorCheck(e1); check != true {
		return check
	}
	var text = elfFile.Section(".text")
	if text == nil {
		elfFile.Close()
		os.Remove("o.o")
		return false
	}
	var textData, e2 = text.Data()
	if check := errorCheck(e2); check != true {
		return check
	}
	*Output = textData
	if check := errorCheck(elfFile.Close()); check != true {
		return check
	}
	if check := errorCheck(os.Remove("o.o")); check != true {
		return check
	}
	return true
}

func windowsAssemble(OS *string, Architecture *int, file *os.File, Output *[]byte) bool { // This function does assembling on Windows system.
	if check := errorCheck(file.Close()); check != true {
		return check
	}

	var format string
	if *OS == headers.Windows {
		format = "win"
	} else if *OS == headers.Linux {
		format = "elf"
	} else {
		fmt.Printf("[-] Unknown operating system value given to Assembler!\nValue: %s\n", *OS)
		return false
	}
	if *Architecture == headers.X86 {
		format += "32"
	} else if *Architecture == headers.X64 {
		format += "64"
	} else {
		fmt.Printf("[-] Unknown Architecture value given to Assembler!\nValue: %d\n", *Architecture)
		return false
	}
	var cmd = exec.Command("nasm", "-f", format, "-o", "o.o", file.Name())
	var output, e0 = cmd.CombinedOutput()
	if check := errorCheck(e0); check != true {
		fmt.Printf("%s\n", output)
		return check
	}
	if check := errorCheck(os.Remove(file.Name())); check != true {
		return check
	}
	if *OS == headers.Linux {
		var elfFile, e2 = elf.Open("o.o")
		if check := errorCheck(e2); check != true {
			return check
		}
		var text = elfFile.Section(".text")
		if text == nil {
			elfFile.Close()
			os.Remove("o.o")
			return false
		}
		var data, e3 = text.Data()
		if check := errorCheck(e3); check != true {
			return check
		}
		*Output = data
		if check := errorCheck(elfFile.Close()); check != true {
			return check
		}
		if check := errorCheck(os.Remove("o.o")); check != true {
			return check
		}
		return true
	}
	var peFile, e2 = pe.Open("o.o")
	if check := errorCheck(e2); check != true {
		return check
	}
	var text = peFile.Section(".text")
	if text == nil {
		peFile.Close()
		os.Remove("o.o")
		return false
	}
	var data, e3 = text.Data()
	if check := errorCheck(e3); check != true {
		return check
	}
	*Output = data
	if check := errorCheck(peFile.Close()); check != true {
		return check
	}
	if check := errorCheck(os.Remove("o.o")); check != true {
		return check
	}
	return true
}

func errorCheck(e error) bool {
	if e != nil {
		fmt.Printf("[-] Assembler returned false: %s\n", e)
		return false
	}
	return true
}
