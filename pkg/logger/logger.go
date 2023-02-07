package log

type Args map[string]any

type Logger interface {
	PanicHandler()
	Info(domain string, layer Layer, method string, args Args)
	Error(domain string, layer Layer, method string, args Args)
	Panic(domain string, layer Layer, method string, args Args)
}
