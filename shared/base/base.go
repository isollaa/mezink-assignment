package base

type Base struct {
	HTTPCode int
	Code     int
	Msg      string
	Data     []BaseData
}

type BaseData struct {
	Field string
	Data  any
}

func Success(httpCode int, data ...BaseData) Base {
	return Base{
		HTTPCode: httpCode,
		Code:     0,
		Msg:      "success",
		Data:     data,
	}
}

func Failure(httpCode int, msg string, data ...BaseData) Base {
	return Base{
		HTTPCode: httpCode,
		Code:     httpCode,
		Msg:      msg,
		Data:     data,
	}
}
