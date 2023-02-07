package log

import (
	"fmt"
	"io"
	"runtime/debug"

	"github.com/rs/zerolog"
)

const (
	LogErrKey  = "err"
	LogRespKey = "resp"
)

const (
	domainJSONKey = "domain"
	layerJSONKey  = "layer"
	methodJSONKey = "method"
	traceJSONKey  = "trace"
	levelJSONKey  = "level"
)

type ZeroLog struct {
	logger zerolog.Logger
}

func New(w io.Writer) Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	l := ZeroLog{logger: zerolog.New(w).With().Timestamp().Logger()}

	return l
}

func (z ZeroLog) PanicHandler() {
	if r := recover(); r != nil {
		z.Panic("unknown", UnsetLayer, "unknown", nil)
	}
}

func (z ZeroLog) Info(domain string, layer Layer, method string, args Args) {
	e := z.logger.Info().
		Str(domainJSONKey, domain).
		Str(layerJSONKey, layer.String()).
		Str(methodJSONKey, method)

	for k, v := range args {
		e.Str(k, fmt.Sprintf("%v", v))
	}

	e.Msg("")
}

func (z ZeroLog) Error(domain string, layer Layer, method string, args Args) {
	e := z.logger.Error().
		Str(domainJSONKey, domain).
		Str(layerJSONKey, layer.String()).
		Str(methodJSONKey, method)

	for k, v := range args {
		e.Str(k, fmt.Sprintf("%v", v))
	}

	e.Msg("")
}

func (z ZeroLog) Panic(domain string, layer Layer, method string, args Args) {
	e := z.logger.Log().
		Str(levelJSONKey, "panic").
		Str(domainJSONKey, domain).
		Str(layerJSONKey, layer.String()).
		Str(methodJSONKey, method).
		Str(traceJSONKey, string(debug.Stack()))

	for k, v := range args {
		e.Str(k, fmt.Sprintf("%v", v))
	}

	e.Msg("")
}
