package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newPrometheus,
	).Build().Apply()
}
