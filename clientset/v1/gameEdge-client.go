package v1

import (
	"context"

	v1 "github.com/difinative/Edge-Operator/api/v1"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type GameEdgeInterface interface {
	List(opt metav1.ListOptions) (*v1.GameEdgeList, error)
	Get(name string, opt metav1.GetOptions) (*v1.GameEdge, error)
	Create(stEdge *v1.GameEdge) (*v1.GameEdge, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type GameEdgeClient struct {
	restClient rest.Interface
	namespace  string
}

func (c *GameEdgeClient) List(opts metav1.ListOptions) (*v1.GameEdgeList, error) {
	result := v1.GameEdgeList{}

	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("gameedges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *GameEdgeClient) Get(name string, opts metav1.GetOptions) (*v1.GameEdge, error) {
	result := v1.GameEdge{}

	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("gameedges").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *GameEdgeClient) Create(gameEdge *v1.GameEdge) (*v1.GameEdge, error) {
	result := v1.GameEdge{}

	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("gameedges").
		Body(gameEdge).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *GameEdgeClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("gameedges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.TODO())
}
