package taipower

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func taipower_parser() Taipower_j {
	res, err := http.Get("https://www.taipower.com.tw/d006/loadGraph/loadGraph/data/loadpara.json")
	if err != nil {
		log.Fatal(err)
	}

	taipower_jsonData, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	data := Taipower_j{}
	if err := json.Unmarshal(taipower_jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
	}

	return data
}

func Taipower_res(more_info bool) string {
	data := taipower_parser()
	var text = "今日電力資訊 "

	switch data.Records[1].ForePeakResvIndicator {
	case "G":
		text = text + " *🟢 供電充裕* - ( " + data.Records[1].PublishTime + " 更新 )\n\n"
	case "Y":
		text = text + " *🟡 供電吃緊* - ( " + data.Records[1].PublishTime + " 更新 )\n\n"
	case "O":
		text = text + " *🟠 供電警戒* - ( " + data.Records[1].PublishTime + " 更新 )\n\n"
	case "R":
		text = text + " *🔴 限電警戒* - ( " + data.Records[1].PublishTime + " 更新 )\n\n"
	case "B":
		text = text + " *⚫️ 限電準備* - ( " + data.Records[1].PublishTime + " 更新 )\n\n"
	default:
	}
	if more_info {
		fore_peak_dema_load, _ := strconv.ParseFloat(data.Records[1].ForePeakDemaLoad, 2)
		fore_maxi_sply_capacity, _ := strconv.ParseFloat(data.Records[1].ForeMaxiSplyCapacity, 2)
		text += "目前用電量： " + data.Records[0].CurrLoad + " 萬瓩\n" +
			"目前供電能力： " + data.Records[3].RealHrMaxiSplyCapacity + " 萬瓩\n" +
			"目前使用率： " + data.Records[0].CurrUtilRate + "%\n" +
			"尖峰使用率： " + strconv.Itoa(int(fore_peak_dema_load)*100/int(fore_maxi_sply_capacity)) + "%\n" +
			"預估最高用電： " + data.Records[1].ForePeakDemaLoad + " 萬瓩\n" +
			"預估最高用電時段：" + data.Records[1].ForePeakHourRange + "\n" +
			"預估最大供電能力： " + data.Records[1].ForeMaxiSplyCapacity + " 萬瓩\n" +
			"預估尖峰備轉容量率： " + data.Records[1].ForePeakResvRate + "%\n" +
			"預估尖峰備轉容量： " + data.Records[1].ForePeakResvCapacity + " 萬瓩\n"
	}

	return text
}
