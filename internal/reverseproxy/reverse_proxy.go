package reverseproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"github.com/ch-load-balancer/internal/config"
	"github.com/rs/zerolog"
)

type ReverseProxyConfig struct {
	Config  *config.Config
	proxies map[string]*httputil.ReverseProxy
	logger  *zerolog.Logger
}

func NewReverseConfig(config *config.Config, log *zerolog.Logger) *ReverseProxyConfig {
	return &ReverseProxyConfig{
		Config:  config,
		proxies: make(map[string]*httputil.ReverseProxy),
		logger:  log,
	}
}

func (r *ReverseProxyConfig) InitReverseProxies() error {
	sharedTransport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     90 * time.Second,
	}

	for _, backend := range r.Config.Backends {
		parseUrl := "http://" + backend.Host + ":" + strconv.Itoa(backend.Port)
		targetUrl, err := url.Parse(parseUrl)
		if err != nil {
			r.logger.Error().Msgf("failed to parse backend URL %s: %v", parseUrl, err)
			return err
		}
		target := targetUrl
		proxy := &httputil.ReverseProxy{
			Rewrite: func(pr *httputil.ProxyRequest) {
				pr.SetURL(target)
			},
			Transport: sharedTransport,
		}
		r.proxies[parseUrl] = proxy
	}
	return nil
}

func (r *ReverseProxyConfig) ForwardRequest(backendStr string, w http.ResponseWriter, req *http.Request) {
	proxy, ok := r.proxies[backendStr]
	if !ok {
		http.Error(w, "Target backend proxy not initialized", http.StatusBadGateway)
		return
	}
	proxy.ServeHTTP(w, req)
}
