package log

type Layer uint8

const (
	UnsetLayer Layer = iota
	StorageLayer
	ServiceLayer
	TransportLayer
)

func (l Layer) String() string {
	switch l {
	case UnsetLayer:
		return "UNSET"
	case StorageLayer:
		return "STORAGE"
	case ServiceLayer:
		return "SERVICE"
	case TransportLayer:
		return "TRANSPORT"
	default:
		return ""
	}
}
