package logger

const (
	LogErrKey  = "err"
	LogRespKey = "response"
)

const (
	DomainJSONKey = "domain"
	LayerJSONKey  = "layer"
	MethodJSONKey = "method"
	TraceJSONKey  = "trace"
	LevelJSONKey  = "level"
)

type Args map[string]any

type Logger interface {
	PanicHandler()
	Info(domain string, layer Layer, method string, args Args)
	Error(domain string, layer Layer, method string, args Args)
	Panic(domain string, layer Layer, method string, args Args)
}
