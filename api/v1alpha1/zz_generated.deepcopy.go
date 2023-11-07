//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backstage) DeepCopyInto(out *Backstage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backstage.
func (in *Backstage) DeepCopy() *Backstage {
	if in == nil {
		return nil
	}
	out := new(Backstage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backstage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackstageList) DeepCopyInto(out *BackstageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backstage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackstageList.
func (in *BackstageList) DeepCopy() *BackstageList {
	if in == nil {
		return nil
	}
	out := new(BackstageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackstageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackstageSpec) DeepCopyInto(out *BackstageSpec) {
	*out = *in
	in.LocalDb.DeepCopyInto(&out.LocalDb)
	in.Deployment.DeepCopyInto(&out.Deployment)
	in.Service.DeepCopyInto(&out.Service)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackstageSpec.
func (in *BackstageSpec) DeepCopy() *BackstageSpec {
	if in == nil {
		return nil
	}
	out := new(BackstageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackstageStatus) DeepCopyInto(out *BackstageStatus) {
	*out = *in
	out.LocalDb = in.LocalDb
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackstageStatus.
func (in *BackstageStatus) DeepCopy() *BackstageStatus {
	if in == nil {
		return nil
	}
	out := new(BackstageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDbParameters) DeepCopyInto(out *LocalDbParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDbParameters.
func (in *LocalDbParameters) DeepCopy() *LocalDbParameters {
	if in == nil {
		return nil
	}
	out := new(LocalDbParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDbPersistentVolume) DeepCopyInto(out *LocalDbPersistentVolume) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDbPersistentVolume.
func (in *LocalDbPersistentVolume) DeepCopy() *LocalDbPersistentVolume {
	if in == nil {
		return nil
	}
	out := new(LocalDbPersistentVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDbSpec) DeepCopyInto(out *LocalDbSpec) {
	*out = *in
	out.Parameters = in.Parameters
	in.PersistentVolume.DeepCopyInto(&out.PersistentVolume)
	in.PersistentVolumeClaim.DeepCopyInto(&out.PersistentVolumeClaim)
	in.Deployment.DeepCopyInto(&out.Deployment)
	in.Service.DeepCopyInto(&out.Service)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDbSpec.
func (in *LocalDbSpec) DeepCopy() *LocalDbSpec {
	if in == nil {
		return nil
	}
	out := new(LocalDbSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDbStatus) DeepCopyInto(out *LocalDbStatus) {
	*out = *in
	out.PersistentVolume = in.PersistentVolume
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDbStatus.
func (in *LocalDbStatus) DeepCopy() *LocalDbStatus {
	if in == nil {
		return nil
	}
	out := new(LocalDbStatus)
	in.DeepCopyInto(out)
	return out
}
