package timeservice

import (
	"encoding/binary"
	"time"
)

func decodeTime(buf []byte) time.Time {
	uTime := int64(binary.BigEndian.Uint32(buf))
	return time.Unix(uTime, 0)
}

func encodeTime(t time.Time) []byte {
	buf := make([]byte, 4)
	unixTime := uint32(t.Unix())
	binary.BigEndian.PutUint32(buf, unixTime)
	return buf
}
