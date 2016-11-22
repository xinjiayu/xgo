package model

type GpsData struct {
	Msg_id     string  `json:"msg_id"`
	Device_num string  `json:"device_num"`
	Auth_code  string  `json:"auth_code"`
	Content    Content `json:"content"`
}

type Content struct {
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Speed     string `json:"speed"`
	Timestamp int64  `json:"timestamp"`
	//海拔高度
	Alt string `json:"alt"`
	//	是否纠偏，0-未纠偏，1-纠偏
	Rectify string `json:"rectify"`
	//	方向
	Direction string `json:"direction"`
	Lat_type  string `json:"lat_type"`
	Lon_type  string `json:"lon_type"`
}
