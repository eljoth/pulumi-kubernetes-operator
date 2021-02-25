// *** WARNING: this file was generated by crd2pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Stack is the Schema for the stacks API
type Stack struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// StackSpec defines the desired state of Pulumi Stack being managed by this operator.
	Spec StackSpecPtrOutput `pulumi:"spec"`
	// StackStatus defines the observed state of Stack
	Status StackStatusPtrOutput `pulumi:"status"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOption) (*Stack, error) {
	if args == nil {
		args = &StackArgs{}
	}
	args.ApiVersion = pulumi.StringPtr("pulumi.com/v1alpha1")
	args.Kind = pulumi.StringPtr("Stack")
	var resource Stack
	err := ctx.RegisterResource("kubernetes:pulumi.com/v1alpha1:Stack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStack gets an existing Stack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackState, opts ...pulumi.ResourceOption) (*Stack, error) {
	var resource Stack
	err := ctx.ReadResource("kubernetes:pulumi.com/v1alpha1:Stack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stack resources.
type stackState struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// StackSpec defines the desired state of Pulumi Stack being managed by this operator.
	Spec *StackSpec `pulumi:"spec"`
	// StackStatus defines the observed state of Stack
	Status *StackStatus `pulumi:"status"`
}

type StackState struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// StackSpec defines the desired state of Pulumi Stack being managed by this operator.
	Spec StackSpecPtrInput
	// StackStatus defines the observed state of Stack
	Status StackStatusPtrInput
}

func (StackState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackState)(nil)).Elem()
}

type stackArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// StackSpec defines the desired state of Pulumi Stack being managed by this operator.
	Spec *StackSpec `pulumi:"spec"`
	// StackStatus defines the observed state of Stack
	Status *StackStatus `pulumi:"status"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// StackSpec defines the desired state of Pulumi Stack being managed by this operator.
	Spec StackSpecPtrInput
	// StackStatus defines the observed state of Stack
	Status StackStatusPtrInput
}

func (StackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackArgs)(nil)).Elem()
}