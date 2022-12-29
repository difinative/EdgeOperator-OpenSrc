package v1

import (
	"context"

	v1 "github.com/difinative/Edge-Operator/api/v1"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type ScEdgeInterface interface {
	List(opt metav1.ListOptions) (*v1.ScEdgeList, error)
	Get(name string, opt metav1.GetOptions) (*v1.ScEdge, error)
	Create(stEdge *v1.ScEdge) (*v1.ScEdge, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type ScEdgeClient struct {
	restClient rest.Interface
	namespace  string
}

func (c *ScEdgeClient) List(opts metav1.ListOptions) (*v1.ScEdgeList, error) {
	result := v1.ScEdgeList{}

	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("scedges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *ScEdgeClient) Get(name string, opts metav1.GetOptions) (*v1.ScEdge, error) {
	result := v1.ScEdge{}

	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("scedges").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *ScEdgeClient) Create(scEdge *v1.ScEdge) (*v1.ScEdge, error) {
	result := v1.ScEdge{}

	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("scedges").
		Body(scEdge).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *ScEdgeClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("scedges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.TODO())
}
