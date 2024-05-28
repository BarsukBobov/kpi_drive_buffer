package kpidrive

type SaveFactForm struct {
	PeriodStart         string `form:"period_start" binding:"required,validDate"`
	PeriodEnd           string `form:"period_end" binding:"required,validDate"`
	PeriodKey           string `form:"period_key" binding:"required"`
	IndicatorToMoID     int    `form:"indicator_to_mo_id" binding:"required"`
	IndicatorToMoFactID int    `form:"indicator_to_mo_fact_id"`
	Value               int    `form:"value"`
	FactTime            string `form:"fact_time" binding:"required,validDate"`
	IsPlan              int    `form:"is_plan"`
	AuthUserID          int    `form:"auth_user_id" binding:"required"`
	Comment             string `form:"comment"`
}

type SaveFactResponse struct {
	MESSAGES struct {
		Error   []string `json:"error"`
		Warning []string `json:"warning"`
		Info    []string `json:"info"`
	} `json:"MESSAGES"`
	DATA struct {
		IndicatorToMoFactId int `json:"indicator_to_mo_fact_id"`
	} `json:"DATA"`
	STATUS string `json:"STATUS"`
}

type SaveFactResponseWithStatus struct {
	Body       *SaveFactResponse
	StatusCode int
}
