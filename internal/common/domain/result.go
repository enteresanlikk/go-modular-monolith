package common_domain

type result struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func _result(isSuccess bool, message string, data interface{}) result {
	return result{
		IsSuccess: isSuccess,
		Message:   message,
		Data:      data,
	}
}

func SuccessResult(message string) result {
	return _result(true, message, nil)
}

func ErrorResult(message string) result {
	return _result(false, message, nil)
}

func SuccessDataResult(message string, data interface{}) result {
	return _result(true, message, data)
}

func ErrorDataResult(message string, data interface{}) result {
	return _result(false, message, data)
}
