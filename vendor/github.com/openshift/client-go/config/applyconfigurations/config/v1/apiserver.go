// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	internal "github.com/openshift/client-go/config/applyconfigurations/internal"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// APIServerApplyConfiguration represents a declarative configuration of the APIServer type for use
// with apply.
type APIServerApplyConfiguration struct {
	metav1.TypeMetaApplyConfiguration    `json:",inline"`
	*metav1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                                 *APIServerSpecApplyConfiguration `json:"spec,omitempty"`
	Status                               *configv1.APIServerStatus        `json:"status,omitempty"`
}

// APIServer constructs a declarative configuration of the APIServer type for use with
// apply.
func APIServer(name string) *APIServerApplyConfiguration {
	b := &APIServerApplyConfiguration{}
	b.WithName(name)
	b.WithKind("APIServer")
	b.WithAPIVersion("config.openshift.io/v1")
	return b
}

// ExtractAPIServer extracts the applied configuration owned by fieldManager from
// aPIServer. If no managedFields are found in aPIServer for fieldManager, a
// APIServerApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// aPIServer must be a unmodified APIServer API object that was retrieved from the Kubernetes API.
// ExtractAPIServer provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractAPIServer(aPIServer *configv1.APIServer, fieldManager string) (*APIServerApplyConfiguration, error) {
	return extractAPIServer(aPIServer, fieldManager, "")
}

// ExtractAPIServerStatus is the same as ExtractAPIServer except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractAPIServerStatus(aPIServer *configv1.APIServer, fieldManager string) (*APIServerApplyConfiguration, error) {
	return extractAPIServer(aPIServer, fieldManager, "status")
}

func extractAPIServer(aPIServer *configv1.APIServer, fieldManager string, subresource string) (*APIServerApplyConfiguration, error) {
	b := &APIServerApplyConfiguration{}
	err := managedfields.ExtractInto(aPIServer, internal.Parser().Type("com.github.openshift.api.config.v1.APIServer"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(aPIServer.Name)

	b.WithKind("APIServer")
	b.WithAPIVersion("config.openshift.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithKind(value string) *APIServerApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithAPIVersion(value string) *APIServerApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithName(value string) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithGenerateName(value string) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithNamespace(value string) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithUID(value types.UID) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithResourceVersion(value string) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithGeneration(value int64) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithCreationTimestamp(value apismetav1.Time) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithDeletionTimestamp(value apismetav1.Time) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *APIServerApplyConfiguration) WithLabels(entries map[string]string) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *APIServerApplyConfiguration) WithAnnotations(entries map[string]string) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *APIServerApplyConfiguration) WithOwnerReferences(values ...*metav1.OwnerReferenceApplyConfiguration) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *APIServerApplyConfiguration) WithFinalizers(values ...string) *APIServerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *APIServerApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &metav1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithSpec(value *APIServerSpecApplyConfiguration) *APIServerApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *APIServerApplyConfiguration) WithStatus(value configv1.APIServerStatus) *APIServerApplyConfiguration {
	b.Status = &value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *APIServerApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}
