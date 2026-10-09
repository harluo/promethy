package core

import (
	"net/http"
	"os"
	"strings"

	"github.com/goexl/log"
	"github.com/goexl/prometheus"
	"github.com/harluo/httpd"
	"github.com/harluo/prometheus/internal/config"
	"github.com/harluo/prometheus/internal/constant"
)

type Registry = prometheus.Registry

func newRegistry(
	conf *config.Prometheus,
	server *httpd.Server, logger log.Logger,
) (registry *Registry, err error) {
	builder := prometheus.New()
	builder.Logger(logger)
	for key, value := range conf.Labels {
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
	} else if conf.Port == 0 || conf.Port == server.Port() {
		server.Get(conf.Path, handler)
	} else {
		mux := http.NewServeMux()
		mux.Handle(conf.Path, handler)
		server.Http().Handler = mux
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
