package convert

import "fmt"

const (
	KBV  float64 = 1000.0
	KiBV float64 = 1024.0
)

func isValidUnit(unit string) bool {
	switch unit {
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
	KiB float64
	MiB float64
	GiB float64
	TiB float64
	PiB float64
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
	case "KB", "kb":
		b.KB = float64(value)
		b.KiB = (float64(value) * KBV) / KiBV
	case "MB", "mb":
		b.KB = float64(value) * KBV
		b.KiB = (float64(value) * KBV * KBV) / KiBV
	case "GB", "gb":
		b.KB = float64(value) * KBV * KBV
		b.KiB = (float64(value) * KBV * KBV * KBV) / KiBV
	case "TB", "tb":
		b.KB = float64(value) * KBV * KBV * KBV
		b.KiB = (float64(value) * KBV * KBV * KBV * KBV) / KiBV
	case "PB", "pb":
		b.KB = float64(value) * KBV * KBV * KBV * KBV
		b.KiB = (float64(value) * KBV * KBV * KBV * KBV * KBV) / KiBV
	case "KiB", "kib":
		b.KB = (float64(value) * KiBV) / KBV
		b.KiB = float64(value)
	case "MiB", "mib":
		b.KB = (float64(value) * KiBV * KiBV) / KBV
		b.KiB = float64(value) * KiBV
	case "GiB", "gib":
		b.KB = (float64(value) * KiBV * KiBV * KiBV) / KBV
		b.KiB = float64(value) * KiBV * KiBV
	case "TiB", "tib":
		b.KB = (float64(value) * KiBV * KiBV * KiBV * KiBV) / KBV
		b.KiB = float64(value) * KiBV * KiBV * KiBV
	case "PiB", "pib":
		b.KB = (float64(value) * KiBV * KiBV * KiBV * KiBV * KiBV) / KBV
		b.KiB = float64(value) * KiBV * KiBV * KiBV * KiBV
	}

	b.MB = b.KB / KBV
	b.GB = b.MB / KBV
	b.TB = b.GB / KBV
	b.PB = b.TB / KBV

	b.MiB = b.KiB / KiBV
	b.GiB = b.MiB / KiBV
	b.TiB = b.GiB / KiBV
	b.PiB = b.TiB / KiBV

	return b, nil
}
