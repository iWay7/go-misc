package misc

import "strconv"

func AbsInt(a, b int) int {
	r := a - b
	if r < 0 {
		r = -r
	}
	return r
}

func AbsInt64(a, b int64) int64 {
	r := a - b
	if r < 0 {
		r = -r
	}
	return r
}

func AbsInt32(a, b int32) int32 {
	r := a - b
	if r < 0 {
		r = -r
	}
	return r
}

func TryDivide(a int64, b int64) float64 {
	if b == 0 {
		return 0.0
	}
	return float64(a) / float64(b)
}

func TryDivideToPercent(a int64, b int64) float64 {
	return TryDivide(a, b) * 100
}

func CalcStart(pageSize, pageIndex int) int {
	return pageSize * (pageIndex - 1)
}

func CalcLimit(pageSize, pageIndex int) int {
	return pageSize
}

func CalcPageCount(pageSize int, totalCount int64) int64 {
	if totalCount%int64(pageSize) == 0 {
		return totalCount / int64(pageSize)
	}
	return totalCount/int64(pageSize) + 1
}

func TryParseInt(s string, def int) int {
	i, err := strconv.Atoi(s)
	if err == nil {
		return i
	}
	return def
}

func TryParseInt64(s string, def int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return i
	}
	return def
}

func TryParseFloat32(s string, def float32) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err == nil {
		return float32(f)
	}
	return def
}

func TryParseFloat64(s string, def float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return f
	}
	return def
}

func TryParseBool(s string, def bool) bool {
	f, err := strconv.ParseBool(s)
	if err == nil {
		return f
	}
	return def
}

func TryParseIntSlice(ss []string, def int) []int {
	if ss == nil {
		return nil
	}
	ii := []int{}
	for _, s := range ss {
		ii = append(ii, TryParseInt(s, def))
	}
	return ii
}

func TryParseInt64Slice(ss []string, def int64) []int64 {
	if ss == nil {
		return nil
	}
	ii := []int64{}
	for _, s := range ss {
		ii = append(ii, TryParseInt64(s, def))
	}
	return ii
}

func TryParseFloat32Slice(ss []string, def float32) []float32 {
	if ss == nil {
		return nil
	}
	ff := []float32{}
	for _, s := range ss {
		ff = append(ff, TryParseFloat32(s, def))
	}
	return ff
}

func TryParseFloat64Slice(ss []string, def float64) []float64 {
	if ss == nil {
		return nil
	}
	ff := []float64{}
	for _, s := range ss {
		ff = append(ff, TryParseFloat64(s, def))
	}
	return ff
}

func TryParseBoolSlice(ss []string, def bool) []bool {
	if ss == nil {
		return nil
	}
	bb := []bool{}
	for _, s := range ss {
		bb = append(bb, TryParseBool(s, def))
	}
	return bb
}
