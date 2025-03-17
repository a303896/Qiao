package excel

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"strconv"
	"strings"
)

func excel2string() {
	excelFileName := "../heronickname.xlsx" // Excel文件名

	// 打开Excel文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("无法打开Excel文件:", err)
		return
	}

	// 遍历所有的Sheet
	for _, sheet := range xlFile.Sheets {
		fmt.Println("Sheet名称:", sheet.Name)

		// 遍历Sheet中的行
		for _, row := range sheet.Rows {
			//if index == 0 {
			//	continue
			//}
			// 遍历行中的单元格
			text := ""
			for i, cell := range row.Cells {
				tmp := strings.TrimSpace(cell.String())
				if i == 0 {
					text += tmp + "=>\""
				} else {
					if tmp == "" {
						continue
					}
					text += tmp + "|"
				}
			}
			text += "\","
			fmt.Println(text) // 换行
		}
	}
}

func excel2string3() {
	rowNum := 3
	excelFileName := "../tiku2.xlsx" // Excel文件名

	// 打开Excel文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("无法打开Excel文件:", err)
		return
	}
	qType := map[string]int{"英雄漫画问题": 0, "台词互动": 1, "小队故事": 2, "桑启问题": 3, "区域问题": 4, "其他趣味信息（i.e. 身高)": 5}
	data := make(map[string][]map[string]string)
	languages := make([]string, 0)
	// 遍历所有的Sheet
	for _, sheet := range xlFile.Sheets {
		//if idx > 0 {
		//	break
		//}
		fmt.Println("Sheet名称:", sheet.Name)
		sheetData := make([]map[string]string, 0)
		languages = append(languages, sheet.Name)
		// 遍历Sheet中的行
		if sheet.Name == "zh" {
			for index, row := range sheet.Rows {
				if index > rowNum {
					break
				}
				if index == 0 {
					continue
				}
				rawData := make(map[string]string)
				tmpJsonData := map[string]string{}
				for i, cell := range row.Cells {
					tmp := strings.TrimSpace(cell.String())
					if i == 0 {
						rawData["type"] = strconv.Itoa(qType["英雄漫画问题"])
					} else if i == 1 {
						rawData["id"] = tmp
					} else if i == 2 {
						rawData["content"] = base64.StdEncoding.EncodeToString([]byte(tmp))
					} else if i == 3 {
						tmpJsonData["A"] = base64.StdEncoding.EncodeToString([]byte(tmp))
					} else if i == 4 {
						tmpJsonData["B"] = base64.StdEncoding.EncodeToString([]byte(tmp))
					} else if i == 5 {
						tmpJsonData["C"] = base64.StdEncoding.EncodeToString([]byte(tmp))
					} else if i == 6 {
						tmpJsonData["D"] = base64.StdEncoding.EncodeToString([]byte(tmp))
					} else if i == 7 {
						jsonByte, _ := json.Marshal(tmpJsonData)
						rawData["options"] = string(jsonByte)
						rawData["answer"] = tmp
						break
					}
				}
				sheetData = append(sheetData, rawData)
			}
		} else {
			for index, row := range sheet.Rows {
				if index > rowNum {
					break
				}
				if index == 0 {
					continue
				}
				rawData := make(map[string]string)
				tmpJsonData := map[string]string{}
				for i, cell := range row.Cells {
					tmp := strings.TrimSpace(cell.String())
					if i == 0 {
						rawData["id"] = tmp
					} else if i == 1 {
						rawData["content"] = base64.StdEncoding.EncodeToString([]byte(tmp))
					} else if i == 2 {
						tmpJsonData["A"] = base64.StdEncoding.EncodeToString([]byte(tmp))
					} else if i == 3 {
						tmpJsonData["B"] = base64.StdEncoding.EncodeToString([]byte(tmp))
					} else if i == 4 {
						tmpJsonData["C"] = base64.StdEncoding.EncodeToString([]byte(tmp))
					} else if i == 5 {
						tmpJsonData["D"] = base64.StdEncoding.EncodeToString([]byte(tmp))
						jsonByte, _ := json.Marshal(tmpJsonData)
						rawData["options"] = string(jsonByte)
						break
					}
				}
				sheetData = append(sheetData, rawData)
			}
		}
		data[sheet.Name] = sheetData
	}
	result := make([]map[string]string, 0)
	for key, val := range data["zh"] {
		raw := make(map[string]string)
		contentJsonData := map[string]string{}
		optionsJsonData := map[string]map[string]string{}
		raw["id"] = val["id"]
		raw["type"] = val["type"]
		raw["answer"] = val["answer"]
		raw["from"] = "1"
		for _, lang := range languages {
			tmp := make(map[string]string)
			contentJsonData[lang] = data[lang][key]["content"]
			json.Unmarshal([]byte(data[lang][key]["options"]), &tmp)
			optionsJsonData[lang] = tmp
		}
		contentByte, _ := json.Marshal(contentJsonData)
		raw["content"] = string(contentByte)
		optionsByte, _ := json.Marshal(optionsJsonData)
		raw["options"] = string(optionsByte)
		result = append(result, raw)
	}
	//fmt.Printf("%+v\n", result)
	strJson, _ := json.Marshal(result)
	fmt.Printf("%+s\n", strJson)
}

func excel2string2() {
	excelFileName := "../tiku.xlsx" // Excel文件名

	// 打开Excel文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("无法打开Excel文件:", err)
		return
	}
	qType := map[string]int{"英雄漫画问题": 0, "台词互动": 1, "小队故事": 2, "桑启问题": 3, "区域问题": 4, "其他趣味信息（i.e. 身高)": 5}
	rowData := make([][]string, 0)
	// 遍历所有的Sheet
	for idx, sheet := range xlFile.Sheets {
		if idx > 0 {
			break
		}
		fmt.Println("Sheet名称:", sheet.Name)

		// 遍历Sheet中的行
		for index, row := range sheet.Rows {
			if index > 166 {
				break
			}
			if index == 0 {
				continue
			}
			tmpData := make([]string, 5)
			tmpData[0] = strconv.Itoa(index + 0)
			tmpJsonData := map[string]string{}
			for i, cell := range row.Cells {
				tmp := strings.TrimSpace(cell.String())
				if i == 0 {
					tmpData[3] = strconv.Itoa(qType[tmp])
				} else if i == 1 {
					tmpData[1] = tmp
				} else if i == 2 {
					tmpJsonData["A"] = tmp
				} else if i == 3 {
					tmpJsonData["B"] = tmp
				} else if i == 4 {
					tmpJsonData["C"] = tmp
				} else if i == 5 {
					tmpJsonData["D"] = tmp
				} else if i == 6 {
					jsonByte, _ := json.Marshal(tmpJsonData)
					tmpData[2] = string(jsonByte)
					tmpData[4] = tmp
					break
				}
			}
			rowData = append(rowData, tmpData)
		}
	}
	fmt.Printf("%+v\n", rowData)
	fmt.Printf("%d\n", len(rowData))
	// 创建一个新的CSV文件
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatalf("无法创建文件: %s", err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF") //写入utf8 bom 防乱码
	// 创建一个新的csv.Writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//写入字段名
	writer.Write([]string{"id", "content", "options", "type", "answer"})
	// 将数据写入CSV文件
	for _, record := range rowData {
		if err := writer.Write(record); err != nil {
			log.Fatalf("无法写入记录到文件: %s", err)
		}
	}
}

func excel2string4() {
	excelFileName := "../tiku3.xlsx" // Excel文件名

	// 打开Excel文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("无法打开Excel文件:", err)
		return
	}
	//qType := map[string]int{"英雄漫画问题": 0, "台词互动": 1, "小队故事": 2, "桑启问题": 3, "区域问题": 4, "其他趣味信息（i.e. 身高)": 5}
	result := make([]map[string]string, 0)
	//languages := make([]string, 0)
	// 遍历所有的Sheet
	for _, sheet := range xlFile.Sheets {
		//if idx > 0 {
		//	continue
		//}
		fmt.Println("Sheet名称:", sheet.Name)
		// 遍历Sheet中的行
		for index, row := range sheet.Rows {
			if index > 2 {
				break
			}
			if index == 0 {
				continue
			}
			rawData := make(map[string]string)
			tmpJsonData := map[string]string{}
			content := ""
			//options := make(map[string]string)
			for i, cell := range row.Cells {
				tmp := strings.TrimSpace(cell.String())
				rawData["type"] = "0"
				rawData["from"] = "1"
				if i == 0 {
					rawData["id"] = tmp
				} else if i == 1 {
					content = base64.StdEncoding.EncodeToString([]byte(tmp))
				} else if i == 2 {
					tmpJsonData["A"] = base64.StdEncoding.EncodeToString([]byte(tmp))
				} else if i == 3 {
					tmpJsonData["B"] = base64.StdEncoding.EncodeToString([]byte(tmp))
				} else if i == 4 {
					tmpJsonData["C"] = base64.StdEncoding.EncodeToString([]byte(tmp))
				} else if i == 5 {
					tmpJsonData["D"] = base64.StdEncoding.EncodeToString([]byte(tmp))
				} else if i == 6 {
					rawData["answer"] = tmp
				} else if i == 7 {
					jsonByte, _ := json.Marshal(map[string]map[string]string{tmp: tmpJsonData})
					rawData["options"] = string(jsonByte)
					contentJsonByte, _ := json.Marshal(map[string]string{tmp: content})
					rawData["content"] = string(contentJsonByte)
					break
				}
			}
			result = append(result, rawData)
		}
	}
	strJson, _ := json.Marshal(result)
	fmt.Printf("%+s\n", strJson)
}
