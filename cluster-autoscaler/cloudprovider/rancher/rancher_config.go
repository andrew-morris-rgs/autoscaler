/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rancher

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type cloudConfig struct {
	URL                             string `yaml:"url"`
	Token                           string `datapolicy:"security-key" yaml:"token"`
	ClusterName                     string `yaml:"clusterName"`
	ClusterNamespace                string `yaml:"clusterNamespace"`
	ClusterAPIVersion               string `yaml:"clusterAPIVersion"`
	TLSInsecureSkipVerify           bool   `default:"false" yaml:"TLSInsecureSkipVerify"`
	// ServerName is passed to the server for SNI and is used in the client to check server
	// certificates against. If ServerName is empty, the hostname used to contact the
	// server is used.
	RancherTLSServerName            string `default:"" yaml:"RancherTLSServerName"`
	// Server requires TLS client certificate authentication
	ClusterAutoscaleServiceCertFile string `default:"" yaml:"ClusterAutoscaleServiceCertFile"`
	// Server requires TLS client certificate authentication
	ClusterAutoscaleServiceKeyFile  string `default:"" yaml:"ClusterAutoscaleServiceKeyFile"`
	// Trusted root certificates for Rancher server
	RancherServerCARootFile         string `default:"" yaml:"RancherServerCARootFile"`
	// CertData holds PEM-encoded bytes (typically read from a client certificate file).
	// CertData takes precedence over CertFile
	ClusterAutoscaleServiceCertData []byte `default:nil yaml="ClusterAutoscaleServiceCertData"`
	// KeyData holds PEM-encoded bytes (typically read from a client certificate key file).
	// KeyData takes precedence over KeyFile
	ClusterAutoscaleServiceKeyData  []byte `datapolicy:"security-key" default:nil yaml="ClusterAutoscaleServiceKeyData"`
	// CAData holds PEM-encoded bytes (typically read from a root certificates bundle).
	// CAData takes precedence over CAFile
	RancherServerCARoot             []byte `default:nil yaml:"RancherServerCARoot"`
	// NextProtos is a list of supported application level protocols, in order of preference.
	// Used to populate tls.Config.NextProtos.
	// To indicate to the server http/1.1 is preferred over http/2, set to ["http/1.1", "h2"] (though the server is free to ignore that preference).
	// To use only http/1.1, set to ["http/1.1"].
	NextProtos                      []string `default:nil yaml:NextProtos`
}


func newConfig(file string) (*cloudConfig, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read cloud config file: %w", err)
	}

	config := &cloudConfig{}
	if err := yaml.Unmarshal(b, config); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config file: %w", err)
	}

	return config, nil
}
