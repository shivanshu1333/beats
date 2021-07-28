package otlp

import (
	"bytes"
	"github.com/elastic/beats/v7/libbeat/logp"
	//otlplogs "go.opentelemetry.io/proto/otlp/logs/v1"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/outputs/codec"
	"github.com/elastic/go-structform/gotype"
)

// Encoder for serializing a beat.Event to json.
type Encoder struct {
	buf    bytes.Buffer
	folder *gotype.Iterator

	version string
	//OtlpResourceLogs *otlplogs.ResourceLogs
	OtlpEncoder *Codec
	config Config
	logger *logp.Logger
}

// Config is used to pass encoding parameters to New.
type Config struct {
}

var defaultConfig = Config{
}

func init() {
	codec.RegisterType("otlp", func(info beat.Info, cfg *common.Config) (codec.Codec, error) {
		config := defaultConfig
		if cfg != nil {
			if err := cfg.Unpack(&config); err != nil {
				return nil, err
			}
		}

		return New(info.Version, config), nil
	})
}

// New creates a new json Encoder.
func New(version string, config Config) *Encoder {
	e := &Encoder{version: version, config: config}
	e.reset()
	return e
}

func (e *Encoder) reset() {
	//var err error
	//
	//if err != nil {
	//	panic(err)
	//}
}

// Encode serializes a beat event to OTLP. It adds additional metadata in the
// `@metadata` namespace.
func (e *Encoder) Encode(index string, event *beat.Event) ([]byte, error) {
	e.buf.Reset()
	err := e.folder.Fold(makeEvent(index, e.version, event))
	if err != nil {
		e.reset()
		return nil, err
	}
	buf, er := e.OtlpEncoder.NewCodec(event)
	if er != nil {
		e.logger.Warn("Error ", er, "on creating new otlp codec")
		e.logger.Warn("Beat event is ", event)
		return nil, er
	}

	//var buf bytes.Buffer
	//if err = stdjson.Indent(&buf, json, "", "  "); err != nil {
	//	return nil, err
	//}

	return buf, nil
}
