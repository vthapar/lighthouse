package controller

import (
	"github.com/submariner-io/admiral/pkg/syncer/broker"
	"k8s.io/client-go/rest"
)

type Controller struct {
	clusterID         string
	restConfig        *rest.Config
	brokerNS          string
	excludeNamespaces map[string]bool
	svcSyncer         *broker.Syncer
}

type AgentSpecification struct {
	ClusterID string
	Namespace string
	ExcludeNS []string
}

type BrokerSpecification struct {
	APIServer       string
	APIServerToken  string
	RemoteNamespace string
	Insecure        bool `default:"false"`
	Ca              string
}
