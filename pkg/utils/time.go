package utils

import "time"

type JsonTime time.Time

func (jt *JsonTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1] // Remove the quotes around the string
	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	*jt = JsonTime(parsedTime)
	return nil
}

func (jt *JsonTime) ConvertToGoTime() time.Time {
	return time.Time(*jt)
}
