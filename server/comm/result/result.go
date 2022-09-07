package result


type Result struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (errcode *Result) Error() string {
	return errcode.Msg
}

func NewResult(Code int32, Msg string, Data interface{}) *Result {
	return &Result{Code, Msg, Data}
}

func SuccessResult(Data interface{}) *Result {
	return &Result{200, "ok", Data}
}

func (r Result) SetMsg(msg string) *Result {
	r.Msg = msg
	return &r
}

func (r Result) SetData(data interface{}) *Result {
	r.Data = data
	return &r
}

type PageResult struct {
	Page  int         `json:"page,omitempty"`
	Size  int         `json:"size,omitempty"`
	Total int      `json:"total,omitempty"`
	Result *Result `json:"result,omitempty"`
}

func NewPageResult(page, size, total int, res *Result) *PageResult {
	return &PageResult{
		Page: page,
		Size: size,
		Total: total,
		Result: res,
	}
}


var (
	SUCCESS             = NewResult(200, "ok", nil)
	INVALID_PARAMS  	= NewResult(400, "unknown error", nil)
	UNKNOW_ERROR 		= NewResult(500, "unknown exception", nil)
	UNAUTHORIZED        = NewResult(401, "unauthorized", nil)
	FORBIDDEN        = NewResult(403, "forbidden", nil)


	FORMATE_ERROR       = NewResult(1001, "格式转换异常", nil)
	DATA_NOT_FOUND      = NewResult(1002, "不存在该数据", nil)
	DATA_NO_CHANGE      = NewResult(1003, "数据未发生改变", nil)
	PASSWORD_ERROR      = NewResult(1004, "密码错误", nil)
	PARAM_ERROR         = NewResult(1005, "参数错误", nil)
	USER_IS_EXISTED     = NewResult(1006, "用户已存在", nil)
	USER_IS_NOT_EXISTED = NewResult(1007, "用户不存在", nil)
	USER_PENDING_REVIEW = NewResult(1008, "用户待审核", nil)
	ROLE_IS_NOT_EXISTED = NewResult(1009, "角色不存在", nil)


)

