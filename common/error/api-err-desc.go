package api_err

const (
	Success      = 0
	UnknownError = 1
	ErrorUrl     = 2

	ErrorLackParams = 100
	ErrorSignature  = 101
	ErrorParams     = 102
	ErrorDbOperate  = 103
	ErrorApiParams  = 104

	ErrorInternal = 110

	ErrorNoGoodsInfo = 1000
)

var errDesc = map[int]string{
	Success:      "Success",
	UnknownError: "未知错误",
	ErrorUrl:     "访问的接口不存在",

	ErrorLackParams: "缺少必要参数",
	ErrorSignature:  "参数签名错误",
	ErrorParams:     "参数错误",
	ErrorDbOperate:  "数据库操作失败",
	ErrorApiParams:  "接口参数校验失败",

	ErrorInternal: "内部错误",

	ErrorNoGoodsInfo: "暂无该商品",
}

func Desc(code int) string {
	return errDesc[code]
}
