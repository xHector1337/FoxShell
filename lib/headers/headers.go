package headers

import "fmt"

// Operating Systems

const Linux = "Linux"
const Windows = "Windows"
const Multi = "Multi"

// Architectures

const X64 = 0x64
const X86 = 0x86

type Module struct {
	Name            string
	Description     string
	Arch            byte
	OperatingSystem string
	Size            int
	Author          []string

	Notes []string
}

func (module *Module) PrintModuleInfo() bool {
	if len(module.Name) == 0 {
		fmt.Printf("[-] Missing  module name!\n")
		return false
	}
	fmt.Printf("Name: %s\n", module.Name)
	if len(module.Description) == 0 {
		fmt.Printf("[-] Missing module description!\n")
		return false
	}
	fmt.Printf("Description: %s\n", module.Description)
	if module.Arch != X64 && module.Arch != X86 {
		fmt.Printf("[-] Missing module architecture!\n")
		return false
	}
	fmt.Printf("Architecture: x%x\n", module.Arch)
	if module.OperatingSystem != Windows && module.OperatingSystem != Linux && module.OperatingSystem != Multi {
		fmt.Printf("[-] Missing or unsupported module operating system!\n")
		return false
	}
	fmt.Printf("Operating System: %s\n", module.OperatingSystem)
	if len(module.Author) == 0 {
		fmt.Printf("[-] Missing module Author!\n")
		return false
	}
	if len(module.Author) > 1 {
		fmt.Printf("Authors: ")
		for i := 0; i < len(module.Author); i++ {
			if i == len(module.Author)-1 {
				fmt.Printf("[%s]\n", module.Author[i])
			} else {
				fmt.Printf("[%s], ", module.Author[i])
			}
		}
	} else {
		fmt.Printf("Author: %s\n", module.Author[0])
	}
	if module.Size != 0 {
		fmt.Printf("Payload size: %d\n", module.Size)
	}
	if len(module.Notes) != 0 {
		fmt.Printf("Notes:\n")
		for i := 0; i < len(module.Notes); i++ {
			fmt.Printf("\t%s\n", module.Notes[i])
		}
	}
	return true
}
