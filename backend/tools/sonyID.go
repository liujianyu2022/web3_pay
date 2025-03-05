package tools

import "github.com/sony/sonyflake"

type SonyID struct {
	sonyflake *sonyflake.Sonyflake
}

func NewSonyID() *SonyID {
	sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sonyflake == nil {
		panic("sonyflake not created")
	}
	return &SonyID{
		sonyflake: sonyflake,
	}
}

func (sonyId SonyID) GenerateString() (string, error) {
	id, err := sonyId.sonyflake.NextID()
	if err != nil {
		return "", err
	}
	return UintToBase62(id), nil
}