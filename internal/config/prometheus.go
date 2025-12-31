package config

import (
	"github.com/harluo/config"
)

type Prometheus struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled,omitempty"`
	// 路径
	Path string `default:"/metrics" json:"path,omitempty" validate:"required"`
	// 标签
	Labels map[string]string `json:"labels,omitempty"`
}

func newPrometheus(getter config.Getter) (prometheus *Prometheus, err error) {
	prometheus = new(Prometheus)
	err = getter.Get(&struct {
		Prometheus *Prometheus `json:"prometheus,omitempty" validate:"required"`
	}{
		Prometheus: prometheus,
	})

	return
}
