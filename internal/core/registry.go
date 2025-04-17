package core

import (
	"net/http"
	"os"
	"strings"

	"github.com/goexl/log"
	"github.com/goexl/promethy"
	"github.com/harluo/promethy/internal/config"
	"github.com/harluo/promethy/internal/constant"
	"github.com/prometheus/client_golang/prometheus"
)

type Registry = prometheus.Registry

func newRegistry(
	prometheus *config.Prometheus,
	mux *http.ServeMux, logger log.Logger,
) (registry *promethy.Registry, err error) {
	builder := promethy.New()
	builder.Logger(logger)
	for key, value := range prometheus.Labels {
		builder.Label(key, value)
	}
	// 加载特殊的环境变量
	for _, environment := range os.Environ() {
		if key, value, checked := checkEvn(environment); checked {
			builder.Label(key, value)
		}
	}

	prom := builder.Build()
	registry = prom.Register()
	if handler, he := prom.Handler().Handle(); nil != he {
		err = he
	} else {
		mux.Handle(prometheus.Path, handler)
	}

	return
}

func checkEvn(env string) (key string, value string, checked bool) {
	values := strings.Split(env, constant.Equal)
	if strings.HasPrefix(values[0], constant.LabelPrometheusKey) {
		key = values[1]
		value = os.Getenv(strings.ReplaceAll(values[0], constant.LabelPrometheusKey, constant.LabelPrometheusValue))
		if "" != value {
			checked = true
		}
	}

	return
}
