# FoxShell - My Own NasmShell

## Introduction

It's an old and abandoned project of mine which still does a good job. I wanted to have my own msf-nasm_shell that supports both x64 and x86 architectures on both Windows and Linux.

## Requirements

+ NASM installed on PATH
+ Go to build the project.

## Keywords

+ `arch64` to change the architecture to x64.
+ `arch32` to change the architecture to x86.
+ `windows` to change the operating system to Windows. (nasm -f winX
+ `linux` to change the operating system to Linux. (nasm -f elfX)
+ `{` & `}` for multi line support. For example:
```
nasmshell> {
push rax
xor eax,eax
pop rbx
}
50 31 c0 5b
```
