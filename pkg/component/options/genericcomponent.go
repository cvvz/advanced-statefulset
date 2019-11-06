package options

import (
	"github.com/cofyc/advanced-statefulset/pkg/component/config"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	componentbaseconfig "k8s.io/component-base/config"
	"k8s.io/kubernetes/pkg/client/leaderelectionconfig"
)

// GenericComponentOptions holds the options which are generic.
type GenericComponentOptions struct {
	MinResyncPeriod         metav1.Duration
	ContentType             string
	KubeAPIQPS              float32
	KubeAPIBurst            int32
	ControllerStartInterval metav1.Duration
	LeaderElection          componentbaseconfig.LeaderElectionConfiguration
}

// NewGenericComponentOptions returns generic configuration default
// values.
func NewGenericComponentOptions(cfg config.GenericComponentConfiguration) *GenericComponentOptions {
	o := &GenericComponentOptions{
		MinResyncPeriod:         cfg.MinResyncPeriod,
		ContentType:             cfg.ContentType,
		KubeAPIQPS:              cfg.KubeAPIQPS,
		KubeAPIBurst:            cfg.KubeAPIBurst,
		ControllerStartInterval: cfg.ControllerStartInterval,
		LeaderElection:          cfg.LeaderElection,
	}
	return o
}

// AddFlags adds flags related to generic for controller manager to the specified FlagSet.
func (o *GenericComponentOptions) AddFlags(fs *pflag.FlagSet) {
	if o == nil {
		return
	}

	fs.DurationVar(&o.MinResyncPeriod.Duration, "min-resync-period", o.MinResyncPeriod.Duration, "The resync period in reflectors will be random between MinResyncPeriod and 2*MinResyncPeriod.")
	fs.StringVar(&o.ContentType, "kube-api-content-type", o.ContentType, "Content type of requests sent to apiserver.")
	fs.Float32Var(&o.KubeAPIQPS, "kube-api-qps", o.KubeAPIQPS, "QPS to use while talking with kubernetes apiserver.")
	fs.Int32Var(&o.KubeAPIBurst, "kube-api-burst", o.KubeAPIBurst, "Burst to use while talking with kubernetes apiserver.")
	fs.DurationVar(&o.ControllerStartInterval.Duration, "controller-start-interval", o.ControllerStartInterval.Duration, "Interval between starting controller managers.")

	leaderelectionconfig.BindFlags(&o.LeaderElection, fs)
}

// Validate checks validation of GenericComponentOptions.
func (o *GenericComponentOptions) Validate() []error {
	if o == nil {
		return nil
	}

	errs := []error{}
	return errs
}

// ApplyTo fills up generic config with options.
func (o *GenericComponentOptions) ApplyTo(cfg *config.GenericComponentConfiguration) error {
	if o == nil {
		return nil
	}

	cfg.MinResyncPeriod = o.MinResyncPeriod
	cfg.ContentType = o.ContentType
	cfg.KubeAPIQPS = o.KubeAPIQPS
	cfg.KubeAPIBurst = o.KubeAPIBurst
	cfg.ControllerStartInterval = o.ControllerStartInterval
	cfg.LeaderElection = o.LeaderElection

	return nil
}
