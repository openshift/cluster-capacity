package app

import (
	"k8s.io/kubernetes/pkg/scheduler/apis/config"
)

func (o *Options) GetConfig() *config.KubeSchedulerConfiguration {
	return o.config
}

func (o *Options) ReallyApplyDefaults() (err error) {
	o.config, err = o.ApplyDefaults(o.config)
	return err
}
