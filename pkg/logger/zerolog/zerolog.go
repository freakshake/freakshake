package zerolog

import (
	"fmt"
	"io"
	"runtime/debug"

	"github.com/rs/zerolog"

	"github.com/mehdieidi/storm/pkg/logger"
)

type zeroLog struct {
	logger zerolog.Logger
}

func New(w io.Writer) logger.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return zeroLog{logger: zerolog.New(w).With().Timestamp().Logger()}
}

func (z zeroLog) PanicHandler() {
	if r := recover(); r != nil {
		z.Panic("unknown", logger.UnsetLayer, "unknown", logger.Args{"err": r})
	}
}

func (z zeroLog) Info(domain string, layer logger.Layer, method string, args logger.Args) {
	e := z.logger.Info().
		Str(logger.DomainJSONKey, domain).
		Str(logger.LayerJSONKey, layer.String()).
		Str(logger.MethodJSONKey, method)

	for k, v := range args {
		e.Str(k, fmt.Sprintf("%+v", v))
	}

	e.Msg("")
}

func (z zeroLog) Error(domain string, layer logger.Layer, method string, args logger.Args) {
	e := z.logger.Error().
		Str(logger.DomainJSONKey, domain).
		Str(logger.LayerJSONKey, layer.String()).
		Str(logger.MethodJSONKey, method)

	for k, v := range args {
		e.Str(k, fmt.Sprintf("%+v", v))
	}

	e.Msg("")
}

func (z zeroLog) Panic(domain string, layer logger.Layer, method string, args logger.Args) {
	e := z.logger.Log().
		Str(logger.LevelJSONKey, "panic").
		Str(logger.DomainJSONKey, domain).
		Str(logger.LayerJSONKey, layer.String()).
		Str(logger.MethodJSONKey, method).
		Str(logger.TraceJSONKey, string(debug.Stack()))

	for k, v := range args {
		e.Str(k, fmt.Sprintf("%+v", v))
	}

	e.Msg("")
}
