package models

//SvcConfig struct contains configuration file details
type SvcConfig struct {
	Service     Svc `json:"service"`
	DB          DB  `json:"db"`
	Application App `json:"application"`
}

//Svc contains service data
type Svc struct {
	Port    string `json:"port"`
	Company string `json:"Company"`
}

//DB contains db related details.
type DB struct {
	Uname string `json:"uname"`
	Pwd   string `json:"pwd"`
	IP    string `json:"ip"`
}

type App struct {
	Holidays    map[string]string `json:"holidays"`
	TotalLeaves int               `json:"totalLeaves"`
	Company     string            `json:"company"`
	YearStart   string            `json:"yearStart"`
}
