// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	networkv1 "github.com/openshift/api/network/v1"
	internal "github.com/openshift/client-go/network/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// NetNamespaceApplyConfiguration represents a declarative configuration of the NetNamespace type for use
// with apply.
type NetNamespaceApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	NetName                          *string                          `json:"netname,omitempty"`
	NetID                            *uint32                          `json:"netid,omitempty"`
	EgressIPs                        []networkv1.NetNamespaceEgressIP `json:"egressIPs,omitempty"`
}

// NetNamespace constructs a declarative configuration of the NetNamespace type for use with
// apply.
func NetNamespace(name string) *NetNamespaceApplyConfiguration {
	b := &NetNamespaceApplyConfiguration{}
	b.WithName(name)
	b.WithKind("NetNamespace")
	b.WithAPIVersion("network.openshift.io/v1")
	return b
}

// ExtractNetNamespace extracts the applied configuration owned by fieldManager from
// netNamespace. If no managedFields are found in netNamespace for fieldManager, a
// NetNamespaceApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// netNamespace must be a unmodified NetNamespace API object that was retrieved from the Kubernetes API.
// ExtractNetNamespace provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractNetNamespace(netNamespace *networkv1.NetNamespace, fieldManager string) (*NetNamespaceApplyConfiguration, error) {
	return extractNetNamespace(netNamespace, fieldManager, "")
}

// ExtractNetNamespaceStatus is the same as ExtractNetNamespace except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractNetNamespaceStatus(netNamespace *networkv1.NetNamespace, fieldManager string) (*NetNamespaceApplyConfiguration, error) {
	return extractNetNamespace(netNamespace, fieldManager, "status")
}

func extractNetNamespace(netNamespace *networkv1.NetNamespace, fieldManager string, subresource string) (*NetNamespaceApplyConfiguration, error) {
	b := &NetNamespaceApplyConfiguration{}
	err := managedfields.ExtractInto(netNamespace, internal.Parser().Type("com.github.openshift.api.network.v1.NetNamespace"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(netNamespace.Name)

	b.WithKind("NetNamespace")
	b.WithAPIVersion("network.openshift.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithKind(value string) *NetNamespaceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithAPIVersion(value string) *NetNamespaceApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithName(value string) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithGenerateName(value string) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithNamespace(value string) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithUID(value types.UID) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithResourceVersion(value string) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithGeneration(value int64) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithCreationTimestamp(value metav1.Time) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *NetNamespaceApplyConfiguration) WithLabels(entries map[string]string) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *NetNamespaceApplyConfiguration) WithAnnotations(entries map[string]string) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *NetNamespaceApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *NetNamespaceApplyConfiguration) WithFinalizers(values ...string) *NetNamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *NetNamespaceApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithNetName sets the NetName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetName field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithNetName(value string) *NetNamespaceApplyConfiguration {
	b.NetName = &value
	return b
}

// WithNetID sets the NetID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetID field is set to the value of the last call.
func (b *NetNamespaceApplyConfiguration) WithNetID(value uint32) *NetNamespaceApplyConfiguration {
	b.NetID = &value
	return b
}

// WithEgressIPs adds the given value to the EgressIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EgressIPs field.
func (b *NetNamespaceApplyConfiguration) WithEgressIPs(values ...networkv1.NetNamespaceEgressIP) *NetNamespaceApplyConfiguration {
	for i := range values {
		b.EgressIPs = append(b.EgressIPs, values[i])
	}
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *NetNamespaceApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
