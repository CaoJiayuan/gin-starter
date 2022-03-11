package logging

import (
	"gin-demo/pkg/config"
	"gin-demo/pkg/server"

	"github.com/enorith/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggingConfig struct {
	Default  string `yaml:"default"`
	Channels map[string]LogChannelConfig
}

type LogChannelConfig struct {
	Outputs    []string `yaml:"outputs"`
	Errouts    []string `yaml:"errouts"`
	Encoding   string   `yaml:"encoding"`
	TimeFormat string   `yaml:"time_format"`
}

func Regiter(logDir string) server.Register {
	return func(_ *gin.Engine) error {
		var conf LoggingConfig
		config.Unmarshal("logging", &conf)

		logging.WithDefaults(logging.Config{
			BaseDir: logDir,
		})
		for ch, cc := range conf.Channels {
			cr := cc
			logging.DefaultManager.Resolve(ch, func(conf zap.Config) (*zap.Logger, error) {
				conf.OutputPaths = cr.Outputs
				conf.ErrorOutputPaths = cr.Errouts
				if cr.TimeFormat == "" {
					cr.TimeFormat = "2006-01-02T15:04:05.999"
				}
				if cr.Encoding == "" {
					cr.Encoding = "json"
				}
				conf.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(cr.TimeFormat)
				conf.EncoderConfig.StacktraceKey = "trace"
				conf.Encoding = cr.Encoding

				return conf.Build()
			})
		}
		return nil
	}
}
