// +build !ignore_autogenerated

/*
Copyright 2020 the original author or authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpringBootApplication) DeepCopyInto(out *SpringBootApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpringBootApplication.
func (in *SpringBootApplication) DeepCopy() *SpringBootApplication {
	if in == nil {
		return nil
	}
	out := new(SpringBootApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpringBootApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpringBootApplicationList) DeepCopyInto(out *SpringBootApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpringBootApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpringBootApplicationList.
func (in *SpringBootApplicationList) DeepCopy() *SpringBootApplicationList {
	if in == nil {
		return nil
	}
	out := new(SpringBootApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpringBootApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpringBootApplicationSpec) DeepCopyInto(out *SpringBootApplicationSpec) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ApplicationProperties != nil {
		in, out := &in.ApplicationProperties, &out.ApplicationProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpringBootApplicationSpec.
func (in *SpringBootApplicationSpec) DeepCopy() *SpringBootApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(SpringBootApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpringBootApplicationStatus) DeepCopyInto(out *SpringBootApplicationStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.AppliedOpinions != nil {
		in, out := &in.AppliedOpinions, &out.AppliedOpinions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ApplicationPropertiesRef != nil {
		in, out := &in.ApplicationPropertiesRef, &out.ApplicationPropertiesRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpringBootApplicationStatus.
func (in *SpringBootApplicationStatus) DeepCopy() *SpringBootApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(SpringBootApplicationStatus)
	in.DeepCopyInto(out)
	return out
}
