package certmon

import (
	"context"
	"net"
	"time"

	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/agent/token"
	"github.com/hashicorp/consul/tlsutil"
	"github.com/hashicorp/go-hclog"
)

// FallbackFunc is used when the normal cache watch based Certificate
// updating fails to update the Certificate in time and a different
// method of updating the certificate is required.
type FallbackFunc func(context.Context) (*structs.SignedResponse, error)

type Config struct {
	Logger          hclog.Logger
	TLSConfigurator *tlsutil.Configurator
	Cache           Cache
	Tokens          *token.Store
	Fallback        FallbackFunc
	FallbackLeeway  time.Duration
	FallbackRetry   time.Duration
	InitialCerts    *structs.SignedResponse
	DNSSANs         []string
	IPSANs          []net.IP
	Datacenter      string
	NodeName        string
}

// WithCache will cause the created CertMonitor type to use the provided Cache
func (cfg *Config) WithCache(cache Cache) *Config {
	cfg.Cache = cache
	return cfg
}

// WithLogger will cause the created CertMonitor type to use the provided logger
func (cfg *Config) WithLogger(logger hclog.Logger) *Config {
	cfg.Logger = logger
	return cfg
}

// WithTLSConfigurator will cause the created CertMonitor type to use the provided configurator
func (cfg *Config) WithTLSConfigurator(tlsConfigurator *tlsutil.Configurator) *Config {
	cfg.TLSConfigurator = tlsConfigurator
	return cfg
}

// WithTokens will cause the created CertMonitor type to use the provided token store
func (cfg *Config) WithTokens(tokens *token.Store) *Config {
	cfg.Tokens = tokens
	return cfg
}

// WithFallback configures a fallback function to use if the normal update mechanisms
// fail to renew the certificate in time.
func (cfg *Config) WithFallback(fallback FallbackFunc) *Config {
	cfg.Fallback = fallback
	return cfg
}

// WithInitialCerts will cause the the initial TLS Client Certificate and CA certificates
// to be setup properly within the TLS Configurator prepopulated appropriately in the Cache.
func (cfg *Config) WithInitialCerts(info *structs.SignedResponse) *Config {
	cfg.InitialCerts = info
	return cfg
}

// WithDNSSANs configures the CertMonitor to request these DNS SANs when requesting a new
// certificate
func (cfg *Config) WithDNSSANs(sans []string) *Config {
	cfg.DNSSANs = sans
	return cfg
}

// WithIPSANs configures the CertMonitor to request these IP SANs when requesting a new
// certificate
func (cfg *Config) WithIPSANs(sans []net.IP) *Config {
	cfg.IPSANs = sans
	return cfg
}

// WithDatacenter configures the CertMonitor to request Certificates in this DC
func (cfg *Config) WithDatacenter(dc string) *Config {
	cfg.Datacenter = dc
	return cfg
}

// WithNodeName configures the CertMonitor to request Certificates with this agent name
func (cfg *Config) WithNodeName(name string) *Config {
	cfg.NodeName = name
	return cfg
}

// WithFallbackLeeway configures how long after a certificate expires before attempting to
// generarte a new certificate using the fallback mechanism. The default is 10s.
func (cfg *Config) WithFallbackLeeway(leeway time.Duration) *Config {
	cfg.FallbackLeeway = leeway
	return cfg
}

// WithFallbackRetry controls how quickly we will make subsequent invocations of
// the fallback func in the case of it erroring out.
func (cfg *Config) WithFallbackRetry(after time.Duration) *Config {
	cfg.FallbackRetry = after
	return cfg
}
