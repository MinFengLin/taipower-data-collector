package taipower

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (tai *Taipower_data) taipower_parser() {
	res, err := http.Get("https://www.taipower.com.tw/d006/loadGraph/loadGraph/data/loadpara.json")
	if err != nil {
		log.Fatal(err)
	}

	taipower_jsonData, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(taipower_jsonData, &tai); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
	}
}

func (tai *Taipower_data) taipower_fore_peak_resv_indicator() string {
	var text = "今日電力資訊 "

	switch tai.Records[1].ForePeakResvIndicator {
	case "Y":
		text = text + " *🟡 供電吃緊* - ( " + tai.Records[1].PublishTime + " 更新 )\n\n"
	case "O":
		text = text + " *🟠 供電警戒* - ( " + tai.Records[1].PublishTime + " 更新 )\n\n"
	case "R":
		text = text + " *🔴 限電警戒* - ( " + tai.Records[1].PublishTime + " 更新 )\n\n"
	case "B":
		text = text + " *⚫️ 限電準備* - ( " + tai.Records[1].PublishTime + " 更新 )\n\n"
	default:
		text = text + " *🟢 供電充裕* - ( " + tai.Records[1].PublishTime + " 更新 )\n\n"
	}

	return text
}

func (tai *Taipower_data) taipower_more_info() string {
	fore_peak_dema_load, _ := strconv.ParseFloat(tai.Records[1].ForePeakDemaLoad, 2)
	fore_maxi_sply_capacity, _ := strconv.ParseFloat(tai.Records[1].ForeMaxiSplyCapacity, 2)
	text := "目前用電量： " + tai.Records[0].CurrLoad + " 萬瓩\n" +
		"目前供電能力： " + tai.Records[3].RealHrMaxiSplyCapacity + " 萬瓩\n" +
		"目前使用率： " + tai.Records[0].CurrUtilRate + "%\n" +
		"尖峰使用率： " + strconv.Itoa(int(fore_peak_dema_load)*100/int(fore_maxi_sply_capacity)) + "%\n" +
		"預估最高用電： " + tai.Records[1].ForePeakDemaLoad + " 萬瓩\n" +
		"預估最高用電時段：" + tai.Records[1].ForePeakHourRange + "\n" +
		"預估最大供電能力： " + tai.Records[1].ForeMaxiSplyCapacity + " 萬瓩\n" +
		"預估尖峰備轉容量率： " + tai.Records[1].ForePeakResvRate + "%\n" +
		"預估尖峰備轉容量： " + tai.Records[1].ForePeakResvCapacity + " 萬瓩\n"

	return text
}

func Parser_Taipower(more_info *bool) string {
	var taipower = &Taipower_data{}
	var res = ""
	taipower.taipower_parser()

	if *more_info {
		res = taipower.taipower_fore_peak_resv_indicator() + taipower.taipower_more_info()
	} else {
		res = taipower.taipower_fore_peak_resv_indicator()
	}

	return res
}
