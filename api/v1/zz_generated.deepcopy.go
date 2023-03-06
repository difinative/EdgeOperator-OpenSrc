//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
func (in *Edge) DeepCopyInto(out *Edge) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Edge.
func (in *Edge) DeepCopy() *Edge {
	if in == nil {
		return nil
	}
	out := new(Edge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Edge) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeList) DeepCopyInto(out *EdgeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Edge, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeList.
func (in *EdgeList) DeepCopy() *EdgeList {
	if in == nil {
		return nil
	}
	out := new(EdgeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeSpec) DeepCopyInto(out *EdgeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeSpec.
func (in *EdgeSpec) DeepCopy() *EdgeSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeStatus) DeepCopyInto(out *EdgeStatus) {
	*out = *in
	in.HealthVitals.DeepCopyInto(&out.HealthVitals)
	if in.Usecase_Vitals != nil {
		in, out := &in.Usecase_Vitals, &out.Usecase_Vitals
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeStatus.
func (in *EdgeStatus) DeepCopy() *EdgeStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthVitals) DeepCopyInto(out *HealthVitals) {
	*out = *in
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthVitals.
func (in *HealthVitals) DeepCopy() *HealthVitals {
	if in == nil {
		return nil
	}
	out := new(HealthVitals)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthVitalsStatus) DeepCopyInto(out *HealthVitalsStatus) {
	*out = *in
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.FreeMemory = in.FreeMemory
	out.TeleportStatus = in.TeleportStatus
	out.Temperature = in.Temperature
	out.WifiStrength = in.WifiStrength
	out.NetworkLatency = in.NetworkLatency
	out.RamUtilization = in.RamUtilization
	out.CpuUtilization = in.CpuUtilization
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthVitalsStatus.
func (in *HealthVitalsStatus) DeepCopy() *HealthVitalsStatus {
	if in == nil {
		return nil
	}
	out := new(HealthVitalsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatsInt) DeepCopyInto(out *StatsInt) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatsInt.
func (in *StatsInt) DeepCopy() *StatsInt {
	if in == nil {
		return nil
	}
	out := new(StatsInt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatsString) DeepCopyInto(out *StatsString) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatsString.
func (in *StatsString) DeepCopy() *StatsString {
	if in == nil {
		return nil
	}
	out := new(StatsString)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsecaseVitals) DeepCopyInto(out *UsecaseVitals) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsecaseVitals.
func (in *UsecaseVitals) DeepCopy() *UsecaseVitals {
	if in == nil {
		return nil
	}
	out := new(UsecaseVitals)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UsecaseVitals) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsecaseVitalsList) DeepCopyInto(out *UsecaseVitalsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UsecaseVitals, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsecaseVitalsList.
func (in *UsecaseVitalsList) DeepCopy() *UsecaseVitalsList {
	if in == nil {
		return nil
	}
	out := new(UsecaseVitalsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UsecaseVitalsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsecaseVitalsSpec) DeepCopyInto(out *UsecaseVitalsSpec) {
	*out = *in
	if in.Vitals != nil {
		in, out := &in.Vitals, &out.Vitals
		*out = make(map[string][]VitalsToCheck, len(*in))
		for key, val := range *in {
			var outVal []VitalsToCheck
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]VitalsToCheck, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsecaseVitalsSpec.
func (in *UsecaseVitalsSpec) DeepCopy() *UsecaseVitalsSpec {
	if in == nil {
		return nil
	}
	out := new(UsecaseVitalsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsecaseVitalsStatus) DeepCopyInto(out *UsecaseVitalsStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsecaseVitalsStatus.
func (in *UsecaseVitalsStatus) DeepCopy() *UsecaseVitalsStatus {
	if in == nil {
		return nil
	}
	out := new(UsecaseVitalsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Usecases) DeepCopyInto(out *Usecases) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Usecases.
func (in *Usecases) DeepCopy() *Usecases {
	if in == nil {
		return nil
	}
	out := new(Usecases)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Usecases) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsecasesList) DeepCopyInto(out *UsecasesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Usecases, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsecasesList.
func (in *UsecasesList) DeepCopy() *UsecasesList {
	if in == nil {
		return nil
	}
	out := new(UsecasesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UsecasesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsecasesSpec) DeepCopyInto(out *UsecasesSpec) {
	*out = *in
	if in.Usecases != nil {
		in, out := &in.Usecases, &out.Usecases
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsecasesSpec.
func (in *UsecasesSpec) DeepCopy() *UsecasesSpec {
	if in == nil {
		return nil
	}
	out := new(UsecasesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsecasesStatus) DeepCopyInto(out *UsecasesStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsecasesStatus.
func (in *UsecasesStatus) DeepCopy() *UsecasesStatus {
	if in == nil {
		return nil
	}
	out := new(UsecasesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VitalsToCheck) DeepCopyInto(out *VitalsToCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VitalsToCheck.
func (in *VitalsToCheck) DeepCopy() *VitalsToCheck {
	if in == nil {
		return nil
	}
	out := new(VitalsToCheck)
	in.DeepCopyInto(out)
	return out
}
