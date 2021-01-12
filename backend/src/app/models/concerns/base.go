package concerns

import "time"

type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// TODO: Date型への変換
//type TimeOnlyDate struct {
//	time.Time
//}
//
//func (t TimeOnlyDate) MarshalJSON() ([]byte, error) {
//	return []byte(`"` + t.Time.Format("2006-01-02") + `"`), nil
//}
//
//func (t *TimeOnlyDate) UnmarshalJSON(b []byte) error {
//	tmpTime, err := time.Parse("2006-01-02", strings.Trim(string(b), `"`))
//	if err != nil {
//		return err
//	}
//	t.Time = tmpTime
//	return nil
//}
