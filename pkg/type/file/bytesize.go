package file

import "fmt"

type ByteSize float64

const (
	_ ByteSize = 1 << (iota * 10)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
)

func (b ByteSize) String() string {
	switch {
	case b >= ZiB:
		return fmt.Sprintf("%.2fZiB", b/ZiB)
	case b >= EiB:
		return fmt.Sprintf("%.2fEiB", b/EiB)
	case b >= PiB:
		return fmt.Sprintf("%.2fPiB", b/PiB)
	case b >= TiB:
		return fmt.Sprintf("%.2fTiB", b/TiB)
	case b >= GiB:
		return fmt.Sprintf("%.2fGiB", b/GiB)
	case b >= MiB:
		return fmt.Sprintf("%.2fMiB", b/MiB)
	case b >= KiB:
		return fmt.Sprintf("%.2fKiB", b/KiB)
	}
	return fmt.Sprintf("%.2fB", b)
}
