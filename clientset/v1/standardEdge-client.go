package v1

import (
	"context"

	v1 "github.com/difinative/Edge-Operator/api/v1"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type StandardEdgeInterface interface {
	List(opt metav1.ListOptions) (*v1.StandardEdgeList, error)
	Get(name string, opt metav1.GetOptions) (*v1.StandardEdge, error)
	Create(stEdge *v1.StandardEdge) (*v1.StandardEdge, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Update(stEdge *v1.StandardEdge, opt metav1.UpdateOptions) (*v1.StandardEdge, error)
}

type StandardEdgeClient struct {
	restClient rest.Interface
	namespace  string
}

func (c *StandardEdgeClient) List(opts metav1.ListOptions) (*v1.StandardEdgeList, error) {
	result := v1.StandardEdgeList{}

	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("standardedges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *StandardEdgeClient) Get(name string, opts metav1.GetOptions) (*v1.StandardEdge, error) {
	result := v1.StandardEdge{}

	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("standardedges").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *StandardEdgeClient) Create(stEdge *v1.StandardEdge) (*v1.StandardEdge, error) {
	result := v1.StandardEdge{}

	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("standardedges").
		Body(stEdge).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *StandardEdgeClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("standardedges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.TODO())
}

func (c *StandardEdgeClient) Update(stEdge *v1.StandardEdge, opt metav1.UpdateOptions) (*v1.StandardEdge, error) {
	result := v1.StandardEdge{}

	err := c.restClient.
		Put().
		Namespace(c.namespace).
		Resource("standardedges").
		Body(stEdge).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}
