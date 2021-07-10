package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
	"time"
)

const (
	TraceKey = ""
)

func TraceIdMW() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		u := uuid.Must(uuid.NewUUID())
		ctx.Set(TraceKey, u.String())
		ctx.Next()
	}
}

type LoggerFormat struct {
	ServiceName string `json:"service_name,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
	Duration    string `json:"duration,omitempty"`
	Host        string `json:"host,omitempty"`
	Method      string `json:"method,omitempty"`
	Path        string `json:"path,omitempty"`
	Error       string `json:"error,omitempty"`
	TraceId     string `json:"trace_id,omitempty"`
}

func LoggerMW() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			if param.Latency > time.Minute {
				// Truncate in a golang < 1.8 safe way
				param.Latency = param.Latency - param.Latency%time.Second
			}
			l := LoggerFormat{
				ServiceName: "auth_service",
				CreateTime:  param.TimeStamp.Format("2006-01-02 15:04:05"),
				Duration:    param.Latency.String(),
				Host:        param.ClientIP,
				Method:      param.Method,
				Path:        param.Path,
				Error:       param.ErrorMessage,
				TraceId:     param.Keys[TraceKey].(string),
			}
			lStr, _ := json.Marshal(l)
			return string(lStr) + "\n"
		},
		Output:    os.Stdout,
		SkipPaths: nil,
	})
}
