package otlp

import (
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/golang/protobuf/proto"
	otlplogs "go.opentelemetry.io/proto/otlp/logs/v1"
)

type Codec struct {
	//sync.Mutex
	OtlpResourceLogs *otlplogs.ResourceLogs
}



func (c *Codec) NewCodec(event *beat.Event) ([]byte, error) {
	c.OtlpResourceLogs.Reset()
	c.OtlpResourceLogs.InstrumentationLibraryLogs[0].Logs[0].TimeUnixNano = uint64(event.Timestamp.UnixNano())
	//c.OtlpResourceLogs, er := addTimestamp();
	buf, err := proto.Marshal(c.OtlpResourceLogs)
	return buf, err
}
//
//func addTimestamp(){
//
//}
//
//func createNewOtlpResourceLogs(){
//
//}
