package config

import (
	"github.com/harluo/config"
)

type Prometheus struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled"`
	// 路径
	Path string `default:"/metrics" json:"path" yaml:"path" xml:"path" toml:"path" valprometheusate:"required"`
	// 标签
	Labels map[string]string `json:"labels" yaml:"labels" xml:"labels" toml:"labels"`
}

func newPrometheus(config *config.Getter) (prometheus *Prometheus, err error) {
	prometheus = new(Prometheus)
	err = config.Get(&struct {
		Prometheus *Prometheus `json:"prometheus,omitempty" valprometheusate:"required"`
	}{
		Prometheus: prometheus,
	})

	return
}
