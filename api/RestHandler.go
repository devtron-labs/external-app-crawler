package api

import (
	"encoding/json"
	"github.com/devtron-labs/external-app-crawler/common"
	"go.uber.org/zap"
	"net/http"
)

type RestHandler interface {
	TelemetryEventReceiver(w http.ResponseWriter, r *http.Request)
}

func NewRestHandlerImpl(logger *zap.SugaredLogger) *RestHandlerImpl {
	return &RestHandlerImpl{
		logger: logger,
	}
}

type RestHandlerImpl struct {
	logger *zap.SugaredLogger
}
type Response struct {
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Result interface{} `json:"result,omitempty"`
	Errors []*ApiError `json:"errors,omitempty"`
}
type ApiError struct {
	HttpStatusCode    int         `json:"-"`
	Code              string      `json:"code,omitempty"`
	InternalMessage   string      `json:"internalMessage,omitempty"`
	UserMessage       interface{} `json:"userMessage,omitempty"`
	UserDetailMessage string      `json:"userDetailMessage,omitempty"`
}

func (impl RestHandlerImpl) writeJsonResp(w http.ResponseWriter, err error, respBody interface{}, status int) {
	response := Response{}
	response.Code = status
	response.Status = http.StatusText(status)
	if err == nil {
		response.Result = respBody
	} else {
		apiErr := &ApiError{}
		apiErr.Code = "000" // 000=unknown
		apiErr.InternalMessage = err.Error()
		apiErr.UserMessage = respBody
		response.Errors = []*ApiError{apiErr}

	}
	b, err := json.Marshal(response)
	if err != nil {
		impl.logger.Error("error in marshaling err object", err)
		status = 500
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

type ResetRequest struct {
	AppId         int `json:"appId"`
	EnvironmentId int `json:"environmentId"`
}

func (impl *RestHandlerImpl) TelemetryEventReceiver(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var telemetryEvent common.TelemetryUserAnalyticsDto
	err := decoder.Decode(&telemetryEvent)
	if err != nil {
		impl.logger.Errorw("error in decode request", "error", err)
		writeJsonResp(w, err, nil, http.StatusBadRequest)
		return
	}

	impl.logger.Infow("req received for telemetry event", "req", telemetryEvent)

	impl.logger.Debugw("save", "status", nil)
	impl.writeJsonResp(w, err, "", 200)
}
