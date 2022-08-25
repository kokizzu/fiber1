package business

type CommonIn struct {
	Cookie string
}

type CommonOut struct {
	StatusCode int    `json:"code,omitempty"`
	ErrorMsg   string `json:"error,omitempty"`
	SetCookie  string `json:"token,omitempty"`
}

func (o *CommonOut) SetError(statusCode int, errMsg string) {
	o.StatusCode = statusCode
	o.ErrorMsg = errMsg
}

type CommonOutput interface {
	Common() *CommonOut
}

func (o CommonOut) Common() *CommonOut {
	return &o
}
