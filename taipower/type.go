package taipower

type Taipower_j struct {
	Success string `json:"success"`
	Result  struct {
		ResourceID string `json:"resource_id"`
	} `json:"result"`
	Records []struct {
		CurrLoad               string `json:"curr_load,omitempty"`
		CurrUtilRate           string `json:"curr_util_rate,omitempty"`
		ForeMaxiSplyCapacity   string `json:"fore_maxi_sply_capacity,omitempty"`
		ForePeakDemaLoad       string `json:"fore_peak_dema_load,omitempty"`
		ForePeakResvCapacity   string `json:"fore_peak_resv_capacity,omitempty"`
		ForePeakResvRate       string `json:"fore_peak_resv_rate,omitempty"`
		ForePeakResvIndicator  string `json:"fore_peak_resv_indicator,omitempty"`
		ForePeakHourRange      string `json:"fore_peak_hour_range,omitempty"`
		PublishTime            string `json:"publish_time,omitempty"`
		YdayDate               string `json:"yday_date,omitempty"`
		YdayMaxiSplyCapacity   string `json:"yday_maxi_sply_capacity,omitempty"`
		YdayPeakDemaLoad       string `json:"yday_peak_dema_load,omitempty"`
		YdayPeakResvCapacity   string `json:"yday_peak_resv_capacity,omitempty"`
		YdayPeakResvRate       string `json:"yday_peak_resv_rate,omitempty"`
		YdayPeakResvIndicator  string `json:"yday_peak_resv_indicator,omitempty"`
		RealHrMaxiSplyCapacity string `json:"real_hr_maxi_sply_capacity,omitempty"`
		RealHrPeakTime         string `json:"real_hr_peak_time,omitempty"`
	} `json:"records"`
}
