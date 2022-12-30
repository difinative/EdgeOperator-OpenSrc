package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type V1Client struct {
	restClient rest.Interface
}

func (c *V1Client) StandardEdges(namespace string) StandardEdgeInterface {
	return &StandardEdgeClient{
		restClient: c.restClient,
		namespace:  namespace,
	}
}

func (c *V1Client) GameEdges(namespace string) GameEdgeInterface {
	return &GameEdgeClient{
		restClient: c.restClient,
		namespace:  namespace,
	}
}

func (c *V1Client) ScEdges(namespace string) ScEdgeInterface {
	return &ScEdgeClient{
		restClient: c.restClient,
		namespace:  namespace,
	}
}

func NewConfig(c *rest.Config) (*V1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{
		Group:   "operator.difinative",
		Version: "v1",
	}

	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	clt, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &V1Client{restClient: clt}, nil
}
