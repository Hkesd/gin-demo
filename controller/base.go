package controller

import (
	"gin-demo/common/error"
	"net/http"
)

type ResponseJson struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ReturnSuccessData 用于返回业务正常的数据
func ReturnSuccessData(info map[string]interface{}) (int, ResponseJson) {
	return http.StatusOK, ResponseJson{api_err.Success, api_err.Desc(api_err.Success), info}
}

// ReturnData 用于返回业务失败的数据
func ReturnData(code int, info map[string]interface{}) (int, ResponseJson) {
	return http.StatusOK, ResponseJson{code, api_err.Desc(code), info}
}

// ReturnSuccessStruct 返回结构体
func ReturnSuccessStruct(info interface{}) (int, ResponseJson) {
	return http.StatusOK, ResponseJson{api_err.Success, api_err.Desc(api_err.Success), info}
}
