// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new clusters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for clusters API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddCluster creates cluster

This operation is used to create a new Kubernetes cluster for an
organization. The desired configuration can be specified using the
__cluster definition format__ (see
[external documentation](https://github.com/giantswarm/api-spec/blob/master/details/CLUSTER_DEFINITION.md)
for details).

The cluster definition format allows to set a number of optional
configuration details, like memory size and number of CPU cores.
However, one attribute is __mandatory__ upon creation: The `owner`
attribute must carry the name of the organization the cluster will
belong to. Note that the acting user must be a member of that
organization in order to create a cluster.

It is *recommended* to also specify the `name` attribute to give the
cluster a friendly name, like e. g. "Development Cluster".

Additional definition attributes can be used. Where attributes are
omitted, default configuration values will be applied. For example, if
no `release_version` is specified, the most recent version is used.

The number of `availability_zones` affects the total number of nodes
that can be created in the cluster. The number of availability zones
splits the IP range that can be used for the cluster in multiple smaller
IP ranges. The [getInfo](#operation/getInfo) endpoint provides more
details about the cluster IP range.

IP range example:

If a cluster gets a `/22` range (1022 hosts) and the cluster should be
spawned across 3 availability zones, the range will then be split up
into four `/24` (254 hosts) that can be assigned to four different
availability zones. One range will stay unused because network
addresses must be powers of two. See [CIDR addressing](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).
Each of the `/24` will then be split into two `/25` (126 hosts) for
public and private subnets. The private subnet is used for nodes and
internal loadbalancer (only if you create them within Kubernetes). The
public subnet will be used by the public loadbalancers. Tenant cluster
come with two public loadbalancers by default. One for the Kubernetes API
and one for Ingress.

__Note:__ AWS ELBs can take up to 8 IP addresses due to the way how
they scale. In addition to this, every AWS subnet has four first
addresses (.1-.4) reserved for internal use.

The `workers` attribute, if present, must contain an array of node
definition objects. The number of objects given determines the number
of workers created.

For example, requesting three worker nodes with default configuration
can be achieved by submitting an array of three empty objects:

```"workers": [{}, {}, {}]```

For clusters on AWS, note that all worker nodes must use the same instance type.

*/
func (a *Client) AddCluster(params *AddClusterParams, authInfo runtime.ClientAuthInfoWriter) (*AddClusterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addCluster",
		Method:             "POST",
		PathPattern:        "/v4/clusters/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddClusterCreated), nil

}

/*
DeleteCluster deletes cluster

This operation allows to delete a cluster.

__Caution:__ Deleting a cluster causes the termination of all workloads running on the cluster. Data stored on the worker nodes will be lost. There is no way to undo this operation.

The response is sent as soon as the request is validated.
At that point, workloads might still be running on the cluster and may be accessible for a little wile, until the cluster is actually deleted.

*/
func (a *Client) DeleteCluster(params *DeleteClusterParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteClusterAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCluster",
		Method:             "DELETE",
		PathPattern:        "/v4/clusters/{cluster_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterAccepted), nil

}

/*
GetCluster gets cluster details

This operation allows to obtain all available details on a particular cluster.

*/
func (a *Client) GetCluster(params *GetClusterParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCluster",
		Method:             "GET",
		PathPattern:        "/v4/clusters/{cluster_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterOK), nil

}

/*
GetClusters gets clusters

This operation fetches a list of clusters.

The result depends on the permissions of the user.
A normal user will get all the clusters the user has access
to, via organization membership.
A user with admin permission will receive a list of all existing
clusters.

The result array items are sparse representations of the cluster objects.
To fetch more details on a cluster, use the [getCluster](#operation/getCluster)
operation.

*/
func (a *Client) GetClusters(params *GetClustersParams, authInfo runtime.ClientAuthInfoWriter) (*GetClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusters",
		Method:             "GET",
		PathPattern:        "/v4/clusters/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClustersOK), nil

}

/*
ModifyCluster modifies cluster

This operation allows to modify an existing cluster.

A cluster modification is performed by submitting a `PATCH` request
to the cluster resource (as described in the
[addCluster](#operation/addCluster) and [getCluster](#operation/getCluster))
in form of a [JSON Patch Merge
(RFC 7386)](https://tools.ietf.org/html/rfc7386). This means, only the
attributes to be modified have to be contained in the request body.

The following attributes can be modified:

- `name`: Rename the cluster to something more fitting.

- `owner`: Changing the owner organization name means to change cluster
ownership from one organization to another. The user performing the
request has to be a member of both organizations.

- `release_version`: By changing this attribute you can upgrade a
cluster to a newer
[release](https://docs.giantswarm.io/api/#tag/releases).

- `workers`: By modifying the array of workers, nodes can be added to
increase the cluster's capacity. See details below.

### Adding and Removing Worker Nodes (Scaling)

Adding worker nodes to a cluster or removing worker nodes from a cluster
works by submitting the `workers` attribute, which contains a (sparse)
array of worker node defintions.

_Sparse_ here means that all configuration details are optional. In the
case that worker nodes are added to a cluster, wherever a configuration
detail is missing, defaults will be applied. See
[Creating a cluster](#operation/addCluster) for details.

When modifying the cluster resource, you describe the desired state.
For scaling, this means that the worker node array submitted must
contain as many elements as the cluster should have worker nodes.
If your cluster currently has five nodes and you submit a workers
array with four elements, this means that one worker node will be removed.
If your submitted workers array has six elements, this means one will
be added.

As an example, this request body could be used to scale a cluster to
three worker nodes:

```json
{
  "workers": [{}, {}, {}]
}
```

If the scaled cluster had four worker nodes before, one would be removed.
If it had two worker nodes before, one with default settings would be
added.

### Limitations

- As of now, existing worker nodes cannot be modified.
- The number of availability zones cannot be modified afterwards.
- When removing nodes (scaling down), it is not possible to determine
which nodes will be removed.
- On AWS based clusters, all worker nodes must use the same EC2 instance
type (`instance_type` node attribute). By not setting an `instance_type`
when submitting a PATCH request, you ensure that the right instance type
is used automatically.

*/
func (a *Client) ModifyCluster(params *ModifyClusterParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "modifyCluster",
		Method:             "PATCH",
		PathPattern:        "/v4/clusters/{cluster_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyClusterOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
