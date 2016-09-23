package models

type McCategory struct {
	Id       int
	Pid      int
	Chi_name string `orm:"size(120)"`
	Eng_name string `orm:"size(120)"`
	Type     uint8
	Thumb    int
	Url      string
	Nav      uint8
}
