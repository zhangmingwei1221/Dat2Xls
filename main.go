package main

import (
	"fmt"
	//_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	LogFile *os.File
)

func main() {
	exeloc := vcl.Application.Location()
	exefilename := vcl.Application.ExeName()
	pathsep := string(os.PathSeparator)
	logpath := exeloc + "logs"
	err := os.MkdirAll(logpath, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	exename := strings.TrimSuffix(exefilename, ".exe")
	exename = strings.TrimPrefix(exename, exeloc)
	logfilename := logpath + pathsep + exename + ".log"

	//logfilename := rtl.SysOpen() + ".log"
	//logfilename := ""
	LogFile, err := os.OpenFile(logfilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
	}
	log.SetOutput(LogFile)
	//log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	log.Printf("日志文件：[%s]\n", logfilename)
	res, err := strconv.Atoi("0x03")
	if err != nil {
		log.Println(res)
	}
	log.Println("..........")
	vcl.RunApp(&mainForm)
}
