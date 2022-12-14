// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	externaldnsv1 "github.com/nginxinc/kubernetes-ingress/pkg/apis/externaldns/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSEndpoints implements DNSEndpointInterface
type FakeDNSEndpoints struct {
	Fake *FakeExternaldnsV1
	ns   string
}

var dnsendpointsResource = schema.GroupVersionResource{Group: "externaldns.nginx.org", Version: "v1", Resource: "dnsendpoints"}

var dnsendpointsKind = schema.GroupVersionKind{Group: "externaldns.nginx.org", Version: "v1", Kind: "DNSEndpoint"}

// Get takes name of the dNSEndpoint, and returns the corresponding dNSEndpoint object, and an error if there is any.
func (c *FakeDNSEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *externaldnsv1.DNSEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsendpointsResource, c.ns, name), &externaldnsv1.DNSEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*externaldnsv1.DNSEndpoint), err
}

// List takes label and field selectors, and returns the list of DNSEndpoints that match those selectors.
func (c *FakeDNSEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *externaldnsv1.DNSEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsendpointsResource, dnsendpointsKind, c.ns, opts), &externaldnsv1.DNSEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &externaldnsv1.DNSEndpointList{ListMeta: obj.(*externaldnsv1.DNSEndpointList).ListMeta}
	for _, item := range obj.(*externaldnsv1.DNSEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSEndpoints.
func (c *FakeDNSEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsendpointsResource, c.ns, opts))

}

// Create takes the representation of a dNSEndpoint and creates it.  Returns the server's representation of the dNSEndpoint, and an error, if there is any.
func (c *FakeDNSEndpoints) Create(ctx context.Context, dNSEndpoint *externaldnsv1.DNSEndpoint, opts v1.CreateOptions) (result *externaldnsv1.DNSEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsendpointsResource, c.ns, dNSEndpoint), &externaldnsv1.DNSEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*externaldnsv1.DNSEndpoint), err
}

// Update takes the representation of a dNSEndpoint and updates it. Returns the server's representation of the dNSEndpoint, and an error, if there is any.
func (c *FakeDNSEndpoints) Update(ctx context.Context, dNSEndpoint *externaldnsv1.DNSEndpoint, opts v1.UpdateOptions) (result *externaldnsv1.DNSEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsendpointsResource, c.ns, dNSEndpoint), &externaldnsv1.DNSEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*externaldnsv1.DNSEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDNSEndpoints) UpdateStatus(ctx context.Context, dNSEndpoint *externaldnsv1.DNSEndpoint, opts v1.UpdateOptions) (*externaldnsv1.DNSEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnsendpointsResource, "status", c.ns, dNSEndpoint), &externaldnsv1.DNSEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*externaldnsv1.DNSEndpoint), err
}

// Delete takes name of the dNSEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeDNSEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dnsendpointsResource, c.ns, name, opts), &externaldnsv1.DNSEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &externaldnsv1.DNSEndpointList{})
	return err
}

// Patch applies the patch and returns the patched dNSEndpoint.
func (c *FakeDNSEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *externaldnsv1.DNSEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsendpointsResource, c.ns, name, pt, data, subresources...), &externaldnsv1.DNSEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*externaldnsv1.DNSEndpoint), err
}
