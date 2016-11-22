package model

type MapDevicegps struct {
	Id              int64  `xorm:"BIGINT(20)"`
	Msg             string `xorm:"VARCHAR(255)"`
	Num             string `xorm:"VARCHAR(255)"`
	Longitudetype   int    `xorm:"default 1 INT(11)"`
	Longitude       string `xorm:"VARCHAR(22)"`
	Latitudetype    int    `xorm:"default 1 INT(11)"`
	Latitude        string `xorm:"VARCHAR(22)"`
	Lng             string `xorm:"VARCHAR(22)"`
	Lat             string `xorm:"VARCHAR(22)"`
	OffsetLongitude string `xorm:"VARCHAR(22)"`
	OffsetLatitude  string `xorm:"VARCHAR(22)"`
	Altitude        string `xorm:"default '0.00' VARCHAR(20)"`
	Speed           string `xorm:"default '0.00' VARCHAR(20)"`
	Course          string `xorm:"default '0.00' VARCHAR(20)"`
	Signal          int    `xorm:"default 0 INT(11)"`
	Energy          int    `xorm:"default 0 INT(11)"`
	Gpstime         int64  `xorm:"BIGINT(13)"`
	Revicetime      int64  `xorm:"BIGINT(13)"`
	Status          int    `xorm:"default 1 INT(1)"`
}
