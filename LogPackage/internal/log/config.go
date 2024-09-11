package log

type Config struct {
	Segment struct {
		MaxStoreBytes uint64 `json:"maxStoreBytes"`
		MaxIndexBytes uint64 `json:"maxIndexBytes"`
		InitialOffset uint64 `json:"initialOffset"`
	}
}
