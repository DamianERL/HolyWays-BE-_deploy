package funddto

type FundRequest struct {
	Name  string `json:"name" form:"name"`
	Image string `json:"image" form:"image"`
	Desc  string `json:"desc" form:"desc"`
	Goals int `json:"goals" form:"goals"`
}
