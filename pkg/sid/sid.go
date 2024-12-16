package sid

import (
	"github.com/bwmarrin/snowflake"
	"github.com/sony/sonyflake"
)

type Sid struct {
	sf *sonyflake.Sonyflake
}

type SnowID struct {
	node *snowflake.Node
}

func NewSid() *Sid {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		panic("sonyflake not created")
	}
	return &Sid{sf}
}

func NewSnowID() *SnowID {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic("snowflake not created")
	}
	return &SnowID{node}
}

func (s Sid) GenString() (string, error) {
	id, err := s.sf.NextID()
	if err != nil {
		return "", err
	}
	return IntToBase62(int(id)), nil
}
func (s Sid) GenUint64() (uint64, error) {
	return s.sf.NextID()
}

func (s SnowID) GenInt64() (int64, error) {
	return s.GenInt64()
}
