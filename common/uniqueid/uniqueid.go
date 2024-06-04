package uniqueid

import (
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
)

var flake *sonyflake.Sonyflake

func init() {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

// A Sonyflake ID is composed of
// 39 bits for time in units of 10 msec
// 8 bits for a sequence number
// 16 bits for a machine id

func GenId() int64 {
	// NextID can continue to generate IDs for about 174 years from StartTime.
	// But after the Sonyflake time is over the limit, NextID returns an error.
	id, err := flake.NextID()
	if err != nil {
		logx.Severef("flake NextID failed with %s \n", err)
		panic(err)
	}
	return int64(id)
}
