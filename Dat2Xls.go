package main

import "C"
import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/xuri/excelize/v2"
)

//var splitor = ""
//var Charset = "GBK"

type sflag struct {
	Filename   string `json:"filename"`
	Headfile   string `json:"headname"`
	HSplitor   string `json:"hsplitor"`
	Splitor    string `json:"splitor"`
	Charset    string `json:"charset"`
	Fldidxlist string `json:"fldidxlist"`
	Topn       string `json:"topn"`
	Reverse    bool   `json:"reverse"`
	CmdFlag    bool   `json:"cmdflag"`
}

var vsflag sflag
var vnflag = 0

//var vsflag = sflag{
//	Topn="-1",
//	HSplitor= "|",
//	splitor  =  "",
//	Charset    = "GBK",
//}

func ParseArgs(args []string) {
	//log.Printf("ParseArgs")

	flag.StringVar(&vsflag.Filename, "f", "111.dat", "--Filename 文件名【可以把文件拖放在此处】")
	flag.StringVar(&vsflag.Splitor, "sp", "\u0003", "--splitor [,  \u0008 ,\u0003, \u0009  ]  分割符，一般为\u0003或者逗号等可见字符单字符")
	flag.StringVar(&vsflag.Headfile, "hf", "111_head.txt", "--headfile 表头文件名【可以把文件拖放在此处,默认以|做分割符】，可以自定义")
	flag.StringVar(&vsflag.HSplitor, "hs", "|", "--HSplitor [,  \u0008 ,\u0003, \u0009  ]  分割符，一般为\u0003或者逗号等可见字符单字符")
	flag.StringVar(&vsflag.Charset, "c", "GBK", "--Charset [,GBK/UTF-8...  ]  字符集编码，一般为GBK")
	flag.StringVar(&vsflag.Fldidxlist, "fdlst", "1", "--fldidxlist 字段列表编号，用英文逗号分割，最后用双引号引用，如1，2，3，4，第一列编号为1")
	flag.StringVar(&vsflag.Topn, "t", "-1", "--Topn 取前多少个")
	flag.BoolVar(&vsflag.Reverse, "d", false, "反向转换")
	flag.BoolVar(&vsflag.CmdFlag, "x", false, "反向转换")

	flag.Usage = func() {
		//os.Stdout = LogFile
		fmt.Fprintln(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nUnrecognized flags will be ignored.\n")
	}
	//flag.CommandLine.Parse(args)
	flag.Parse()

	vnflag = flag.NFlag()
	//rune(splitor)
	fmt.Println("Filename:" + vsflag.Filename)
	//fmt.Println([]rune(splitor))
	//fmt.Printf("%s,%X\n", "splitor:"+string(vsflag.Splitor)+",hex:", vsflag.Splitor)
}
func init() {
	fmt.Println("................init............")
}

func splt(c rune) bool {
	//c = ''
	if c == rune(vsflag.Splitor[0]) {
		return true
	} else {
		return false
	}
}

func ToAlphaString(value int) string {
	if value < 0 {
		return ""
	}
	var ans string
	i := value + 1
	for i > 0 {
		ans = string((i-1)%26+65) + ans
		i = (i - 1) / 26
	}
	return ans
}

func XCode2GrapString(xcode string) string {
	startpos := 0
	if strings.HasPrefix(xcode, "0x") || strings.HasPrefix(xcode, "0X") {
		startpos = 2
	}
	res, err := strconv.ParseInt(xcode[startpos:], 16, 64)
	if err != nil {
		return xcode
	}
	//log.Println(string(res))
	//log.Println(strconv.QuoteToGraphic(xcode[startpos:]))
	return string(res)
	//
	//return *xcode
}

// 列表去重
func uniqvalue(vallist []string) (ret []string) {
	sort.Strings(vallist)
	vlen := len(vallist)
	for i := 0; i < vlen; i++ {
		if (i > 0 && vallist[i-1] == vallist[i]) || len(vallist[i]) == 0 {
			continue
		}
		ret = append(ret, vallist[i])
	}
	fmt.Println(ret)
	return ret
}

func list2sortdig(fldlistidx []string) []int {
	//flds:=strings.Split(fldlistidx,",")
	flds1 := make([]int, len(fldlistidx))
	for i, fld := range fldlistidx {
		if fv, err := strconv.ParseInt(fld, 10, 8); err == nil {
			flds1[i] = int(fv)
		}
	}
	sort.Ints(flds1)
	return flds1
}

// func oper(args []string) {
func oper(jsonData []byte) string {
	log.Println(".........in oper...........")
	var enc mahonia.Decoder
	//oldstdout := os.Stderr
	os.Stdout = LogFile
	//Filename := "C:\\Users\\11\\Downloads\\11.del"
	//var Filename string
	//var headfile string
	//var HSplitor = "|"
	//var fldidxlist string
	//var Topn = -1
	//var reverse = false
	var err error
	if jsonData == nil || len(jsonData) == 0 {

	} else {
		log.Println(".........." + string(jsonData))
		err := json.Unmarshal(jsonData, &vsflag) // 将JSON数据解析到Person变量中
		log.Println(vsflag)
		log.Println(vsflag.CmdFlag)
		if err != nil {
			log.Println("Error:", err)
			return err.Error()
		}
	}

	//log.Println(jsonData)

	log.Printf("%+v\n", vsflag)
	//argslen := 2
	if vsflag.CmdFlag {
		//os.Stdout = oldstdout
		if vnflag > 1 {
			//flag.ErrHelp = errors.New("命令行出错")
			//flag.StringVar(&vsflag.Filename, "f", "111.dat", "--Filename 文件名【可以把文件拖放在此处】")
			//flag.StringVar(&vsflag.Splitor, "sp", "\u0003", "--splitor [,  \u0008 ,\u0003, \u0009  ]  分割符，一般为\u0003或者逗号等可见字符单字符")
			//flag.StringVar(&vsflag.Headfile, "hf", "111_head.txt", "--headfile 表头文件名【可以把文件拖放在此处,默认以|做分割符】，可以自定义")
			//flag.StringVar(&vsflag.HSplitor, "hs", "|", "--HSplitor [,  \u0008 ,\u0003, \u0009  ]  分割符，一般为\u0003或者逗号等可见字符单字符")
			//flag.StringVar(&vsflag.Charset, "c", "GBK", "--Charset [,GBK/UTF-8...  ]  字符集编码，一般为GBK")
			//flag.StringVar(&vsflag.Fldidxlist, "fdlst", "1", "--fldidxlist 字段列表编号，用英文逗号分割，最后用双引号引用，如1，2，3，4，第一列编号为1")
			//flag.StringVar(&vsflag.Topn, "t", "-1", "--Topn 取前多少个")
			//flag.BoolVar(&vsflag.Reverse, "d", false, "反向转换")
			//flag.Parse()
			////rune(splitor)
			//fmt.Println("Filename:" + vsflag.Filename)
			////fmt.Println([]rune(splitor))
			//fmt.Printf("%s,%X\n", "splitor:"+string(vsflag.Splitor)+",hex:", vsflag.Splitor)
			////return
		} else {
			fmt.Println("请输入文件名，可以直接用鼠标点击文件拖入到此(如果有表头文件，请确保同目录下的文件名以_head.txt结束，如1.dat，头文件名为1_head.txt)：")
			fmt.Scan(&vsflag.Filename)
		}
	}
	enc = mahonia.NewDecoder(vsflag.Charset)

	if len(vsflag.Filename) < 5 {
		tmpstr := fmt.Sprintf("输入的文件名[%s]长度好像不够5位，是否有问题？", vsflag.Filename)
		return tmpstr
		//panic(any(tmpstr))
	}
	paths, basefile := filepath.Split(vsflag.Filename)
	filesuffix := path.Ext(basefile)
	newbasename := paths + basefile[0:len(basefile)-len(filesuffix)]
	headname := newbasename + "_head.txt"
	if len(vsflag.Headfile) > 0 {
		headname = vsflag.Headfile
	}
	xlsfilename := newbasename + ".xlsx"
	//res := 0
	vsflag.Splitor = XCode2GrapString(vsflag.Splitor)
	vsflag.HSplitor = XCode2GrapString(vsflag.HSplitor)
	//log.Println(vsflag.Splitor)
	//log.Println(vsflag.HSplitor)
	//log.Printf("Splitor:%v,%x,%c\n", []byte(vsflag.Splitor), res, res)
	//log.Printf("%s,%X\n", "splitor:"+string(vsflag.Splitor)+",hex:", vsflag.Splitor)
	//log.Printf("目标文件[%s]\n", xlsfilename)
	log.Printf("目标文件路径为：[%s]\n", xlsfilename)
	//logfilename := "\n" + xlsfilename
	//log.Printf("目标日志路径为：[%s]\n", logfilename)
	vtopn, err := strconv.Atoi(vsflag.Topn)
	if err != nil {
		vtopn = -1
	}
	funcName(err, vsflag.Filename, headname, vsflag.Splitor, vsflag.HSplitor, enc, xlsfilename, vsflag.Fldidxlist, vtopn)
	return "完成"
}

func funcName(err error, filename string, headname string, splitor string, hsplitor string, enc mahonia.Decoder, xlsfilename string, fldidxlist string, top int) string {
	fp1, err := os.Open(filename)
	if err != nil {
		tmpstr := fmt.Sprintf("%s文件打开错误！", filename)
		log.Println(tmpstr)
		//panic(any(err))
		return tmpstr
	}
	defer fp1.Close()

	row0 := make([]string, 10, 300)

	var headname0 = ""
	_, err = os.Stat(headname)
	if err == nil {
		headname0 = headname
		log.Printf("-----输入的头文件[%s]存在------，以此为解析！\n", headname)
	}

	if len(headname0) == 0 {
		log.Printf("-----输入的头文件不存在------！\n")
	}

	if hsplitor == "" {
		hsplitor = splitor
	}

	fp2, err := os.Open(headname0)
	if err != nil {
		tmpstr := fmt.Sprintf("%s文件打开错误！", headname0)
		log.Println(tmpstr)
	} else {
		log.Printf("-----读取%s------！\n", headname0)
		fs2 := bufio.NewScanner(fp2)
		fs2.Scan()
		splitstrLine := fs2.Text()
		log.Println("splitstrLine", splitstrLine)
		log.Println("hsplitor", hsplitor)
		row0 = strings.Split(splitstrLine, hsplitor)
		log.Println(len(row0))
		tmpstr := "{"
		//for i, rval := range row0 {
		//	tmpstr += strconv.Itoa(i+1) + ":" + enc.ConvertString(rval) + ","
		//}
		for i := 0; i < len(row0); i++ {
			tmpstr += strconv.Itoa(i+1) + ":" + enc.ConvertString(row0[i]) + ","
		}
		tmpstr += "}"
		log.Printf("-----读取字段%d个,第一个是[%s]------！\n", len(row0), tmpstr)
	}
	defer fp2.Close()

	fs := bufio.NewScanner(fp1)
	//fs1 := bufio.NewScanner(fp2)

	var rowno = 2
	if len(row0) > 1 {
		rowno = 1
	}
	log.Printf("-----rowno:%d------！\n", rowno)
	//xlsfilename := "1122221.xlsx"
	f := excelize.NewFile()
	f.SaveAs(xlsfilename)

	file1, err := excelize.OpenFile(xlsfilename)
	if err != nil {
		log.Println(err)
	}

	streamWriter, err := file1.NewStreamWriter(enc.ConvertString("Sheet1"))
	if err != nil {
		log.Println(err)
	}

	defer func() {
		if err := recover(); err != any(nil) {
			streamWriter.Flush()
			f.SaveAs(xlsfilename)
			file1.Save()
			log.Printf("recover:%v", err)
		}
	}()

	f.SetActiveSheet(1)
	//f.SetSheetName("Sheet1")
	//f.SetPanes("Sheet1",`{"freeze":true,"split":false,"x_split":0,"y_split":1,"top_left_cell":"A2","actve_pane":"topRight"}`)
	streamWriter.SetPanes(&excelize.Panes{
		Freeze:      true,
		Split:       false,
		XSplit:      0,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
		Selection: []excelize.Selection{
			{ActiveCell: "A2", Pane: "bottomLeft"},
		},
	})

	//defer f.SaveAs(xlsfilename)
	//index:=f.NewSheet("data")
	//f.SetActiveSheet(index)
	fldidxlist1 := list2sortdig(uniqvalue(strings.Split(fldidxlist, ",")))
	colcnt := 0
	for fs.Scan() {
		splitstrLine := fs.Text()
		//fildstr := strings.FieldsFunc(splitstrLine, splt)
		fildstr := strings.Split(splitstrLine, splitor)
		//log.Printf("fldidxlist1:[%d] fildstr:[%s]", len(fldidxlist1),fildstr)
		if len(fldidxlist1) > 0 {
			colcnt = len(fldidxlist1)
		} else {
			colcnt = len(fildstr)
		}
		//log.Printf("pos:%d-->colcnt:[%d]\n",rowno,colcnt)
		//row := make([]interface{}, colcnt)

		row := make([]interface{}, colcnt)
		if rowno == 1 {
			log.Printf("recover:%v", fldidxlist1)
			if len(fldidxlist1) >= 1 {
				colcnt = len(fldidxlist1)
				for i := 0; i < colcnt; i++ {
					row[i] = enc.ConvertString(row0[fldidxlist1[i]-1])
				}
			} else if len(row0) > 1 {
				//需要用此语法处理
				for i := 0; i < len(row0) && row0[i] != ""; i++ {
					row[i] = enc.ConvertString(row0[i])
				}
			}
			//log.Printf("-----[%d]第一个是[%s],rowno[%d]------！\n",colcnt,row[:colcnt],rowno)
			cell, _ := excelize.CoordinatesToCellName(1, rowno)
			//log.Printf("----1-")
			if err := streamWriter.SetRow(cell, row[:colcnt]); err != nil {
				log.Println(err)
			}
			//log.Printf("-----")
			row = make([]interface{}, len(row))
			//log.Printf("-----2")
			if top > 0 {
				top++
			}
			//log.Printf("Topn:[%d] ",Topn)
			rowno++
		}

		//log.Printf("colcnt:[%d] ",colcnt)
		//row = make([]interface{}, colcnt)
		for colno := 0; colno < colcnt; colno++ {
			tmpcolno := colno
			if len(fldidxlist1) >= 1 {
				tmpcolno = fldidxlist1[colno] - 1
			}
			//log.Printf("tmpcolno:[%d] ",tmpcolno)
			if fildstr[tmpcolno] != "" {
				row[colno] = enc.ConvertString(fildstr[tmpcolno])
			} else {
				row[colno] = ""
			}
		}
		//log.Println("for over ")
		cell, _ := excelize.CoordinatesToCellName(1, rowno)
		if err := streamWriter.SetRow(cell, row[:colcnt]); err != nil {
			log.Println(err)
		}
		if rowno%30000 == 0 {
			log.Printf("-----已写入[%d]行------\n", rowno)
		}

		if rowno > excelize.TotalRows {
			panic(any("rows number exceeds maximum limit"))
		}

		//if rowno >= 20000 {
		//	break
		//}

		if top > 0 && rowno >= top {
			log.Printf("-----已写入[%d]行------\n", rowno)
			break
		}
		rowno++
	}
	if len(row0) > 1 {
		rowno--
	}
	log.Printf("-----共写入数据[%d]行------\n", rowno)

	if err := streamWriter.Flush(); err != nil {
		log.Println(err)
		//return
	}

	if err := file1.Save(); err != nil {
		log.Println(err)
	}
	//println(fp1)
	log.Println("完成!")
	return "完成"
}

func main1() {
	//oper(os.Args)
	//oper(os.Args)
}
