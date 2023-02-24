package logger

const (
	LogErrKey  = "err"
	LogRespKey = "response"
	LogObjKey  = "obj"
)

const (
	DomainJSONKey = "domain"
	LayerJSONKey  = "layer"
	MethodJSONKey = "method"
	TraceJSONKey  = "trace"
	LevelJSONKey  = "level"
	FileJSONKey   = "file"
	LineJSONKey   = "line"
	CallerJSONKey = "caller"
)

type Args map[string]any

type Logger interface {
	PanicHandler()
	Info(Domain, Layer, Args)
	Error(Domain, Layer, error, Args)
	Panic(Domain, Layer, Args)
}
