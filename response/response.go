package response

type ResBase struct {
	ResCode int `json:"resCode"`
	ResMsg string `json:"resMsg"`
}

type ResData struct {
	ResBase
	Data interface{} `json:"data"`
}

var (
	RESP_SUCC = Build(0,"请求成功")
)

func Build(resCode int,resMsg string) *ResBase{
	return &ResBase{
		ResCode: resCode,
		ResMsg:  resMsg,
	}
}

func NewResData(base *ResBase,data interface{})  *ResData{
	return &ResData{
		ResBase: *base,
		Data:    data,
	}
}

func BuildSuccessResp(data interface{})  *ResData{
	return NewResData(RESP_SUCC,data)
}


