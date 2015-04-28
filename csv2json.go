package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	DATA_NUMBER       = "number"
	DATA_ARRAY_NUMBER = "array-number"
	DATA_ARRAY_STRING = "array-string"
	DATA_STRING       = "string"
)

var (
	FILE_NAME = ""
	SAVE_NAME = ""
)

func main() {
	if len(os.Args) > 2 {
		FILE_NAME = os.Args[1]
		SAVE_NAME = os.Args[2]
	} else {
		fmt.Println("参数不足。")
		fmt.Println("请输入 [文件名] [保存的文件名]")
		return
	}

	lst := readFile()
	saveFile(lst)
}

func readFile() []map[string]interface{} {
	cntb, _ := ioutil.ReadFile(FILE_NAME)
	enc := mahonia.NewDecoder("gbk")
	src := enc.ConvertString(string(cntb))

	r2 := csv.NewReader(strings.NewReader(src))
	data, _ := r2.ReadAll()
	lst := []map[string]interface{}{}
	var keyLst []string //关键字数组
	var tLst []string   //类型数组
	var maxLen int
	for index, v := range data {
		if index == 0 {
			continue
		}

		if index == 1 { //创建关键字数组
			keyLst, maxLen = createKeyLst(v)
			continue
		}

		if index == 2 { //创建类型数组
			tLst = createTLst(v, maxLen)
			continue
		}

		fr := make(map[string]interface{})
		count := len(v)
		for j := 0; j < maxLen; j++ {

			if j < count {
				switch tLst[j] {
				case DATA_NUMBER:
					fr[keyLst[j]] = createNumber(v[j])
					break
				case DATA_ARRAY_NUMBER:
					fr[keyLst[j]] = createArrayNumber(v[j])
					break
				case DATA_ARRAY_STRING:
					fr[keyLst[j]] = createArrayString(v[j])
					break
				default:
					fr[keyLst[j]] = v[j]
					break
				}
			}
		}

		lst = append(lst, fr)
	}

	return lst
}

func createArrayString(value string) []string {
	arr := strings.Split(value, "|")
	return arr
}

func createArrayNumber(value string) []int {
	arr := strings.Split(value, "|")
	var nArr []int
	count := len(arr)
	for i := 0; i < count; i++ {
		num := createNumber(arr[i])
		nArr = append(nArr, num)
	}
	return nArr
}

func createNumber(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		num = 0
	}
	return num
}

//返回数组和长度
func createKeyLst(lst []string) ([]string, int) {
	var keyLst []string
	maxLen := len(lst)
	for i := 0; i < maxLen; i++ {
		keyLst = append(keyLst, lst[i])
	}
	return keyLst, maxLen
}

//创建类型数组
func createTLst(lst []string, maxLen int) []string {
	var tLst []string
	count := len(lst)

	for i := 0; i < maxLen; i++ {
		if i < count {
			if lst[i] == DATA_NUMBER || lst[i] == DATA_ARRAY_NUMBER || lst[i] == DATA_ARRAY_STRING || lst[i] == DATA_STRING {
				tLst = append(tLst, lst[i])
			} else {
				tLst = append(tLst, DATA_STRING)
			}
		} else {
			tLst = append(tLst, DATA_STRING)
		}
	}

	return tLst
}

func saveFile(lst []map[string]interface{}) {
	file, err := os.Create(SAVE_NAME)
	if err != nil {
		fmt.Println("create file error:", err.Error())
		fmt.Println("filename:", SAVE_NAME)
		return
	}

	defer file.Close()

	b, _ := json.Marshal(lst)
	fmt.Println("done")
	file.Write(b)
}
