/**
 * Author: K.o.s
 * Date: 2017-10-24
 * Email: longw@sctek.com
**/
// go_main
package main

import (
	"bufio"
	"bytes"
	"container/list"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	skey *string
)

func ConvertToSlice(listStr *list.List) []string {
	sli := []string{}
	for el := listStr.Front(); nil != el; el = el.Next() {
		sli = append(sli, el.Value.(string))
	}
	return sli
}

func GetFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

func PrintFilesName(path string) []string {
	fullPath := GetFullPath(path)
	listStr := list.New()
	filepath.Walk(fullPath, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		if strings.Contains(fi.Name(), "struct.go") {
			listStr.PushBack(fi.Name())
		}
		return nil
	})
	return ConvertToSlice(listStr)
}

//判断文件或文件夹是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func Line_First_Upper(sfile string) error {
	if Exist(sfile) {
		buf, err := ioutil.ReadFile(sfile)
		if err != nil {
			return err
		}
		br := bytes.NewReader(buf)
		rd := bufio.NewReader(br)
		f, _ := os.OpenFile(sfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		defer f.Close()
		for {
			line, err := rd.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			if strings.Contains(line, *skey) {
				s := strings.TrimLeft(line, " ")
				//gofmt 一般就是\x09 + first Letters
				f.Write([]byte(strings.ToUpper(s[:2]) + s[2:]))
			} else {
				f.Write([]byte(line))
			}
		}
	}
	return nil
}

func init() {
	skey = flag.String("keys", "json:", "line Contains keys!")
}

func main() {
	flag.Parse()
	spath := GetFullPath(".")
	flist := PrintFilesName(".")
	if len(flist) > 0 {
		for i := 0; i < len(flist); i++ {
			sfile := filepath.Join(spath, flist[i])
			Line_First_Upper(sfile)
		}
	} else {
		fmt.Println("no found struct define file!")
	}
}
