package otlp

import (
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/golang/protobuf/proto"
	otlpcommon "go.opentelemetry.io/proto/otlp/common/v1"
	otlplogs "go.opentelemetry.io/proto/otlp/logs/v1"
	"sync"
)

type Codec struct {
	sync.Mutex
	OtlpResourceLogs *otlplogs.ResourceLogs
}

type AttributeValue struct {
	orig *otlpcommon.AnyValue
}

type LogRecord struct {
	orig *otlplogs.LogRecord
}


func (c Codec) NewCodec() Codec {
	return Codec{OtlpResourceLogs: &otlplogs.ResourceLogs{}}
}

func (c *Codec) AddLogs(event *beat.Event) ([]byte, error) {
	//c.OtlpResourceLogs.Reset()
	c = AddEmptyInstrumentationLibraryLogs(c)
	c = AddEmptyLogRecord(c)
	c = AddTimeStamp(event, c)
	c = PopulateLogRecord(event, c)



	//c.OtlpResourceLogs, er := addTimestamp();
	//buf, err := proto.Marshal(c.OtlpResourceLogs)
	buf, err := Marshaler(c)
	return buf, err
}

func (c *Codec) Unmarshal(buf []byte) (*otlplogs.ResourceLogs, error) {
	ld := &otlplogs.ResourceLogs{}
	err := proto.Unmarshal(buf, ld)
	return ld, err
}

func AddEmptyInstrumentationLibraryLogs(c *Codec) *Codec {
	c.OtlpResourceLogs.InstrumentationLibraryLogs = append(c.OtlpResourceLogs.InstrumentationLibraryLogs, &otlplogs.InstrumentationLibraryLogs{})
	return c
}

func AddEmptyLogRecord(c *Codec) *Codec {
	c.OtlpResourceLogs.InstrumentationLibraryLogs[0].Logs = append(c.OtlpResourceLogs.InstrumentationLibraryLogs[0].Logs, &otlplogs.LogRecord{})
	return c
}

func AddTimeStamp(event *beat.Event, c *Codec) *Codec {
	c.OtlpResourceLogs.InstrumentationLibraryLogs[0].Logs[0].TimeUnixNano = uint64(event.Timestamp.UnixNano())
	return c
}

func PopulateLogRecord(event *beat.Event, c *Codec) *Codec {
	c.OtlpResourceLogs.InstrumentationLibraryLogs[0].Logs[0].Body = new(otlpcommon.AnyValue)
	c.OtlpResourceLogs.InstrumentationLibraryLogs[0].Logs[0].Body.Value = &otlpcommon.AnyValue_StringValue{StringValue: event.Fields["message"].(string) }
	return c
}

func Marshaler(c *Codec) ([]byte, error){
	buf, err := proto.Marshal(c.OtlpResourceLogs)
	return buf,err
}

//func Body() AttributeValue {
//	return newAttributeValue(&(*otlplogs.LogRecord))
//}

func newAttributeValue(orig *otlpcommon.AnyValue) *otlpcommon.AnyValue {
	return orig
}

//
//func addTimestamp(c *Codec, event *beat.Event){
//
//}


//
//func createNewOtlpResourceLogs(){
//
//}
