package nhle

import "time"

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(data []byte) error {
	date, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}
	t.Time = date
	return nil
}
