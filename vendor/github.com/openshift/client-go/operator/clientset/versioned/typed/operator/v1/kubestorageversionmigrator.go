// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// KubeStorageVersionMigratorsGetter has a method to return a KubeStorageVersionMigratorInterface.
// A group's client should implement this interface.
type KubeStorageVersionMigratorsGetter interface {
	KubeStorageVersionMigrators() KubeStorageVersionMigratorInterface
}

// KubeStorageVersionMigratorInterface has methods to work with KubeStorageVersionMigrator resources.
type KubeStorageVersionMigratorInterface interface {
	Create(ctx context.Context, kubeStorageVersionMigrator *v1.KubeStorageVersionMigrator, opts metav1.CreateOptions) (*v1.KubeStorageVersionMigrator, error)
	Update(ctx context.Context, kubeStorageVersionMigrator *v1.KubeStorageVersionMigrator, opts metav1.UpdateOptions) (*v1.KubeStorageVersionMigrator, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kubeStorageVersionMigrator *v1.KubeStorageVersionMigrator, opts metav1.UpdateOptions) (*v1.KubeStorageVersionMigrator, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.KubeStorageVersionMigrator, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.KubeStorageVersionMigratorList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.KubeStorageVersionMigrator, err error)
	Apply(ctx context.Context, kubeStorageVersionMigrator *operatorv1.KubeStorageVersionMigratorApplyConfiguration, opts metav1.ApplyOptions) (result *v1.KubeStorageVersionMigrator, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, kubeStorageVersionMigrator *operatorv1.KubeStorageVersionMigratorApplyConfiguration, opts metav1.ApplyOptions) (result *v1.KubeStorageVersionMigrator, err error)
	KubeStorageVersionMigratorExpansion
}

// kubeStorageVersionMigrators implements KubeStorageVersionMigratorInterface
type kubeStorageVersionMigrators struct {
	*gentype.ClientWithListAndApply[*v1.KubeStorageVersionMigrator, *v1.KubeStorageVersionMigratorList, *operatorv1.KubeStorageVersionMigratorApplyConfiguration]
}

// newKubeStorageVersionMigrators returns a KubeStorageVersionMigrators
func newKubeStorageVersionMigrators(c *OperatorV1Client) *kubeStorageVersionMigrators {
	return &kubeStorageVersionMigrators{
		gentype.NewClientWithListAndApply[*v1.KubeStorageVersionMigrator, *v1.KubeStorageVersionMigratorList, *operatorv1.KubeStorageVersionMigratorApplyConfiguration](
			"kubestorageversionmigrators",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.KubeStorageVersionMigrator { return &v1.KubeStorageVersionMigrator{} },
			func() *v1.KubeStorageVersionMigratorList { return &v1.KubeStorageVersionMigratorList{} }),
	}
}
