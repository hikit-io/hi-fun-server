package logs

import (
	"auth-service/api/http/middlewares"
	"context"
	"log"
)

func init() {

}

func TraceLog(ctx context.Context, msg string) {
	traceId := ctx.Value(middlewares.TraceKey)
	log.Default().Println(traceId)
}
