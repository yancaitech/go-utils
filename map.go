package utils

// Int64RecordToMap func
func Int64RecordToMap(vals []byte, m map[int64]int) (err error) {
	for {
		if len(vals) < 8 {
			break
		}
		bs := vals[0:8]
		n := BytesToInt64(bs)
		m[n] = 1
		vals = vals[8:]
	}
	return err
}

// MapToInt64Record func
func MapToInt64Record(m map[int64]int) (r []byte, err error) {
	for k := range m {
		r = append(r, Int64ToBytes(k)...)
	}
	return r, err
}
