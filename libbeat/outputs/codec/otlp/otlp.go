//package otlp
//
//import (
//	"bytes"
//	stdjson "encoding/json"
//	logsV1 "go.opentelemetry.io/proto/otlp/logs/v1"
//
//	"github.com/elastic/beats/v7/libbeat/beat"
//	"github.com/elastic/beats/v7/libbeat/common"
//	"github.com/elastic/beats/v7/libbeat/outputs/codec"
//	"github.com/elastic/go-structform/gotype"
//	"github.com/elastic/go-structform/json"
//)
//
//// Encoder for serializing a beat.Event to json.
//type Encoder struct {
//	buf    bytes.Buffer
//	folder *gotype.Iterator
//
//	version string
//	OtlpResourceLogs []*logsV1.ResourceLogs
//	config  Config
//}
//
//// Config is used to pass encoding parameters to New.
//type Config struct {
//}
//
//var defaultConfig = Config{
//}
//
//func init() {
//	codec.RegisterType("otlp", func(info beat.Info, cfg *common.Config) (codec.Codec, error) {
//		config := defaultConfig
//		if cfg != nil {
//			if err := cfg.Unpack(&config); err != nil {
//				return nil, err
//			}
//		}
//
//		return New(info.Version, config), nil
//	})
//}
//
//// New creates a new json Encoder.
//func New(version string, config Config) *Encoder {
//	e := &Encoder{version: version, config: config}
//	e.reset()
//	return e
//}
//
//func (e *Encoder) reset() {
//	//var err error
//	//
//	//if err != nil {
//	//	panic(err)
//	//}
//}
//
//// Encode serializes a beat event to JSON. It adds additional metadata in the
//// `@metadata` namespace.
//func (e *Encoder) Encode(index string, event *beat.Event) ([]byte, error) {
//	e.buf.Reset()
//	err := e.folder.Fold(makeEvent(index, e.version, event))
//	if err != nil {
//		e.reset()
//		return nil, err
//	}
//
//	json := e.buf.Bytes()
//	if !e.config.Pretty {
//		return json, nil
//	}
//
//	var buf bytes.Buffer
//	if err = stdjson.Indent(&buf, json, "", "  "); err != nil {
//		return nil, err
//	}
//
//	return buf.Bytes(), nil
//}
