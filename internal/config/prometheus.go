package config

import (
	"github.com/harluo/boot"
)

type Prometheus struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled"`
	// 路径
	Path string `default:"/metrics" json:"path" yaml:"path" xml:"path" toml:"path" valprometheusate:"required"`
	// 标签
	Labels map[string]string `json:"labels" yaml:"labels" xml:"labels" toml:"labels"`
}

func newPrometheus(config *boot.Config) (prometheus *Prometheus, err error) {
	prometheus = new(Prometheus)
	err = config.Build().Get(&struct {
		Prometheus *Prometheus `json:"prometheus,omitempty" valprometheusate:"required"`
	}{
		Prometheus: prometheus,
	})

	return
}
