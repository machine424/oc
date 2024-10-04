// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// MachineConfigNodeSpecPinnedImageSetApplyConfiguration represents a declarative configuration of the MachineConfigNodeSpecPinnedImageSet type for use
// with apply.
type MachineConfigNodeSpecPinnedImageSetApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// MachineConfigNodeSpecPinnedImageSetApplyConfiguration constructs a declarative configuration of the MachineConfigNodeSpecPinnedImageSet type for use with
// apply.
func MachineConfigNodeSpecPinnedImageSet() *MachineConfigNodeSpecPinnedImageSetApplyConfiguration {
	return &MachineConfigNodeSpecPinnedImageSetApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *MachineConfigNodeSpecPinnedImageSetApplyConfiguration) WithName(value string) *MachineConfigNodeSpecPinnedImageSetApplyConfiguration {
	b.Name = &value
	return b
}
