package rest

import (
  "log"
  "net/url"
  "github.com/ddliu/go-httpclient"
  "github.com/gliderlabs/registrator/bridge"
)

type Factory struct{}

type RestAdapter struct {
  host string
}

func init() {
  f := new(Factory)
  bridge.Register(f, "rest")
}

func (f *Factory) New(uri *url.URL) bridge.RegistryAdapter {
  var host string = uri.Host
  if host == "" {
    host = "http://0.0.0.0:5000"
  } else {
    host = "http://" + uri.Host
  }

  return &RestAdapter{host: host}
}

func (r *RestAdapter) Ping() error {
  var endpoint string
  var err error
  endpoint = r.host + "/ping"
  response, err := httpclient.Get(endpoint)

log.Println("rest: on-ping:", response.StatusCode)

  if err != nil {
    log.Println("rest: Failed to ping")
    return err
  }

  return nil
}

func (r *RestAdapter) Register(service *bridge.Service) error {
  var endpoint string
  var err error

  endpoint = r.host + "/services"
  response, err := httpclient.PostJson(endpoint, service)

  log.Println("rest: on-register:", response.StatusCode)

  if err != nil {
    log.Println("rest: failed to register service:", err)
  }

  return err
}

func (r *RestAdapter) Deregister(service *bridge.Service) error {
  var endpoint string
  var err error

  endpoint = r.host + "/services/" + service.ID
  response, err := httpclient.Delete(endpoint)

  log.Println("rest: on-deregister:", response.StatusCode)

  if err != nil {
    log.Println("rest: failed to deregister service:", err)
  }

  return err
}

func (r *RestAdapter) Refresh(service *bridge.Service) error {
  var endpoint string
  var err error

  endpoint = r.host + "/services/" + service.ID
  response, err := httpclient.PutJson(endpoint, service)

  log.Println("rest: on-refresh:", response.StatusCode)

  if err != nil {
    log.Println("rest: failed to refresh the service", err)
  }

  return err
}

func (r *RestAdapter) Services() ([]*bridge.Service, error) {
  return []*bridge.Service{}, nil
}
