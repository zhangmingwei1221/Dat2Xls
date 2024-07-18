package main

import (
	"encoding/json"
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"log"
	"reflect"
	"strings"
)

// ::private::
type TmainFormFields struct {
}

func (f *TmainForm) OnFormCreate(sender vcl.IObject) {
	//f.ClientWidth()
	f.SetAlign(types.AlTop)
	//fmt.Println(f.ClientWidth(), f.ClientHeight)
	//fmt.Println(f.Width(), f.Height())
	f.EnabledMaximize(false)
	f.SetCaption("Dat文件转Excel")
	f.SetWidth(f.Width() - 200)
	f.SetHeight(f.Height() - 80)
	//width, heigth := vcl.Screen.Width(), vcl.Screen.Height()
	//f.SetLeft((width - f.Width()) / 2)
	//f.SetTop((heigth - f.Height()) / 2)
	f.ScreenCenter()
	f.SetAllowDropFiles(true)
	//return
}

func (f *TmainForm) OnBitBtn_selectfileClick(sender vcl.IObject) {
	fmt.Println("OnBitBtn_selectfileClick")
}

func (f *TmainForm) OnEdit_filenameDragDrop(sender vcl.IObject, source vcl.IObject, x int32, y int32) {
	//fmt.Println("OnEdit_filenameDragOver")
	//fmt.Println(source.ClassType())
	//fmt.Println("OnEdit_filenameDragDrop")

	fmt.Println(sender.ToString())
	fmt.Println(source.ToString())
	fmt.Println(source)
	f.Edit_filename.SetDragMode(types.AlLeft)
	f.Edit_filename.SetOnEndDrag(func(sender, target vcl.IObject, x, y int32) {
		if target != nil {
		}
	})
}

func (f *TmainForm) OnEdit_filenameDblClick(sender vcl.IObject) {
	//fmt.Println("OnEdit_filenameDblClick")
	loc := ""
	if f.Edit_filename.GetTextLen() > 0 {
		loc = f.Edit_filename.Text()
	}
	f.Edit_filename.SetText(f.OnOpenDialogSelFilename(sender, loc))

}

func (f *TmainForm) OnEdit_headnameDblClick(sender vcl.IObject) {
	//fmt.Println("OnEdit_headnameDblClick")
	loc := ""
	if f.Edit_headname.GetTextLen() > 0 {
		loc = f.Edit_headname.Text()
	}
	f.Edit_headname.SetText(f.OnOpenDialogSelFilename(sender, loc))
	//f.OpenDialog1.FileName()
}

func (f *TmainForm) OnOpenDialogSelFilename(sender vcl.IObject, vloc string) (vfilename string) {
	vOpenDialogSelFilename := f.OpenDialogSelFilename
	vOpenDialogSelFilename.SetTitle("请选择文件：")
	if len(vloc) == 0 {
		vloc = vcl.Application.Location()
	}
	vOpenDialogSelFilename.SetInitialDir(vloc)
	vOpenDialogSelFilename.SetFilter("*")
	if vOpenDialogSelFilename.Execute() {
		vfilename = vOpenDialogSelFilename.FileName()
	}
	//vcl.ShowMessage("您选择的文件是" + vfilename)

	return vfilename
}

func (f *TmainForm) OnButton_exitClick(sender vcl.IObject) {
	canClose := vcl.MessageDlg("是否退出?", types.MtConfirmation, types.MbYes, types.MbNo) == types.MrYes
	if canClose {
		f.Close()
	}
}

func (f *TmainForm) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {

}

func (f *TmainForm) OnFormDropFiles(sender vcl.IObject, fileNames []string) {
	fmt.Println("OnFormDropFiles")
	fmt.Println(fileNames)

	//f.Edit_filename.SetText(fileNames[0])
}

func (f *TmainForm) OnEdit_filenameDragOver(sender vcl.IObject, source vcl.IObject, x int32, y int32, state types.TDragState, accept *bool) {
	vaccept := true
	accept = &vaccept
	//实现文件拖拽操作
	log.Println("OnEdit_filenameDragOver")
	log.Println(source)
	//print(source.ClassType())
	fmt.Println("OnEdit_filenameDragOver")
	fmt.Println(source.ClassType())

}
func (f *TmainForm) OnEdit_filenameDragDropEvent(sender vcl.IObject, source vcl.IObject, x int32, y int32) {
	//实现文件拖拽操作
	log.Println("OnEdit_filenameDragOver")
	log.Println(source)
	//print(source.ClassType())
	fmt.Println("OnEdit_filenameDragOver")
	fmt.Println(source.ClassType())
}

// extractStructFields 递归地提取结构体字段的值
func extractStructFields(v reflect.Value, parentFieldName string) {
	// 确保我们处理的是一个结构体
	if v.Kind() == reflect.Struct {
		// 遍历结构体的所有字段
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			typeField := v.Type().Field(i)
			fieldName := typeField.Name

			// 构建当前字段的完整名称（如果有父字段的话）
			var fullName string
			if parentFieldName != "" {
				fullName = parentFieldName + "." + fieldName
			} else {
				fullName = fieldName
			}

			// 根据字段类型进行不同的处理
			switch field.Kind() {
			case reflect.String:
				// 对于字符串类型，直接提取值
				fmt.Printf("%s: %s\n", fullName, field.String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				// 对于整型，直接提取值
				fmt.Printf("%s: %d\n", fullName, field.Int())
			case reflect.Struct:
				// 如果字段是另一个结构体，递归处理它
				extractStructFields(field, fullName)
			default:
				// 对于其他类型，可以根据需要添加处理逻辑
				fmt.Printf("Unsupported field type in %s: %s\n", fullName, field.Type())
			}
		}
	}
}

func (f *TmainForm) OnButton_convertClick(sender vcl.IObject) {
	fmt.Println("OnButton_convertClick")
	t := reflect.TypeOf(*mainForm)
	v := reflect.ValueOf(*mainForm)
	var myMap = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.String() == "*vcl.TEdit" {
			//fmt.Println("TEdit:Field:", v.Field(i).Type())
			vcontrolNames := strings.Split(t.Field(i).Name, "_")
			vcontrolName := vcontrolNames[0]
			if len(vcontrolNames) > 1 {
				vcontrolName = vcontrolNames[len(vcontrolNames)-1]
			}
			if structValue, ok := v.Field(i).Interface().(*vcl.TEdit); ok {
				fmt.Printf("Fieldname and value:[%s:%s:%s]\n", t.Field(i).Name, vcontrolName, structValue.Text())
				myMap[vcontrolName] = structValue.Text()
			} else {
				fmt.Println("Type assertion failed")
			}
		}
		//fmt.Println("Field:", t.Field(i).Name, "Value:", v.Field(i))

	}
	jsonData, _ := json.Marshal(myMap)
	fmt.Println(string(jsonData))
	vcl.ShowMessage(oper(jsonData))
}
