// +build !ignore_autogenerated

/*


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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOMKillProtector) DeepCopyInto(out *OOMKillProtector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOMKillProtector.
func (in *OOMKillProtector) DeepCopy() *OOMKillProtector {
	if in == nil {
		return nil
	}
	out := new(OOMKillProtector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OOMKillProtector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOMKillProtectorList) DeepCopyInto(out *OOMKillProtectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OOMKillProtector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOMKillProtectorList.
func (in *OOMKillProtectorList) DeepCopy() *OOMKillProtectorList {
	if in == nil {
		return nil
	}
	out := new(OOMKillProtectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OOMKillProtectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOMKillProtectorSpec) DeepCopyInto(out *OOMKillProtectorSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOMKillProtectorSpec.
func (in *OOMKillProtectorSpec) DeepCopy() *OOMKillProtectorSpec {
	if in == nil {
		return nil
	}
	out := new(OOMKillProtectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOMKillProtectorStatus) DeepCopyInto(out *OOMKillProtectorStatus) {
	*out = *in
	if in.ProtectorStatus != nil {
		in, out := &in.ProtectorStatus, &out.ProtectorStatus
		*out = new(ProtectorStatus)
		**out = **in
	}
	if in.StatusUpdateTime != nil {
		in, out := &in.StatusUpdateTime, &out.StatusUpdateTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOMKillProtectorStatus.
func (in *OOMKillProtectorStatus) DeepCopy() *OOMKillProtectorStatus {
	if in == nil {
		return nil
	}
	out := new(OOMKillProtectorStatus)
	in.DeepCopyInto(out)
	return out
}
