package convert

import "fmt"

const (
	KBV  float64 = 1000.0
	KiBV float64 = 1024.0
)

func isValidUnit(unit string) bool {
	switch unit {
	case "B", "b", "Bi", "BI", "bi":
		return true
	case "KB", "kb", "KiB", "kib":
		return true
	case "MB", "mb", "MiB", "mib":
		return true
	case "GB", "gb", "GiB", "gib":
		return true
	case "TB", "tb", "TiB", "tib":
		return true
	case "PB", "pb", "PiB", "pib":
		return true
	}

	return false
}

type ByteValues struct {
	Bi  float64
	KiB float64
	MiB float64
	GiB float64
	TiB float64
	PiB float64
	B   float64
	KB  float64
	MB  float64
	GB  float64
	TB  float64
	PB  float64
}

func Bytes(value uint64, unit string) (*ByteValues, error) {
	if !isValidUnit(unit) {
		return nil, fmt.Errorf("%s is not avalid unit", unit)
	}

	b := &ByteValues{}

	switch unit {
	case "B", "b", "BI", "Bi", "bi":
		b.B = float64(value)
		b.Bi = float64(value)
	case "KB", "kb":
		b.B = float64(value) * KBV
		b.Bi = (float64(value) * KiBV * KiBV) / KBV
	case "MB", "mb":
		b.B = float64(value) * KBV * KBV
		b.Bi = (float64(value) * KiBV * KiBV * KiBV) / KBV
	case "GB", "gb":
		b.B = float64(value) * KBV * KBV * KBV
		b.Bi = (float64(value) * KiBV * KiBV * KiBV * KiBV) / KBV
	case "TB", "tb":
		b.B = float64(value) * KBV * KBV * KBV * KBV
		b.Bi = (float64(value) * KiBV * KiBV * KiBV * KiBV * KiBV) / KBV
	case "PB", "pb":
		b.B = float64(value) * KBV * KBV * KBV * KBV * KBV
		b.Bi = (float64(value) * KiBV * KiBV * KiBV * KiBV * KiBV * KiBV) / KBV
	case "KiB", "kib":
		b.B = (float64(value) * KBV * KBV) / KiBV
		b.Bi = float64(value) * KiBV
	case "MiB", "mib":
		b.B = (float64(value) * KBV * KBV * KBV) / KiBV
		b.Bi = float64(value) * KiBV * KiBV
	case "GiB", "gib":
		b.B = (float64(value) * KBV * KBV * KBV * KBV) / KiBV
		b.Bi = float64(value) * KiBV * KiBV * KiBV
	case "TiB", "tib":
		b.B = (float64(value) * KBV * KBV * KBV * KBV * KBV) / KiBV
		b.Bi = float64(value) * KiBV * KiBV * KiBV * KiBV
	case "PiB", "pib":
		b.B = (float64(value) * KBV * KBV * KBV * KBV * KBV * KBV) / KiBV
		b.Bi = float64(value) * KiBV * KiBV * KiBV * KiBV * KiBV
	}

	b.KB = b.B / KBV
	b.MB = b.KB / KBV
	b.GB = b.MB / KBV
	b.TB = b.GB / KBV
	b.PB = b.TB / KBV

	b.KiB = b.Bi / KiBV
	b.MiB = b.KiB / KiBV
	b.GiB = b.MiB / KiBV
	b.TiB = b.GiB / KiBV
	b.PiB = b.TiB / KiBV

	return b, nil
}
