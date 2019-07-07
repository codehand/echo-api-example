package types

// var global
var (
	OkStatus PayloadStatus = PayloadStatus{Code: "SUCCESS", Message: ""}
)

// PayloadStatus is class
type PayloadStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ParseStatus is func init status
func ParseStatus(code string, message string) PayloadStatus {
	return PayloadStatus{
		Code:    code,
		Message: message,
	}
}

// HasError is func test err
func (e PayloadStatus) HasError() bool {
	if e.Code == "SUCCESS" {
		return false
	}
	return e.Code != ""
}
