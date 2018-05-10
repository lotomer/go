package common

// NoneWriter 空
type NoneWriter struct {
}

func (w *NoneWriter) Write(p []byte) (int, error) {
	return 0, nil
}
