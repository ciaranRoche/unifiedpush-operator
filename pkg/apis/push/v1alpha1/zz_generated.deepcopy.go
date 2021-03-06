// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServer) DeepCopyInto(out *UnifiedPushServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServer.
func (in *UnifiedPushServer) DeepCopy() *UnifiedPushServer {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnifiedPushServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServerBackup) DeepCopyInto(out *UnifiedPushServerBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServerBackup.
func (in *UnifiedPushServerBackup) DeepCopy() *UnifiedPushServerBackup {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServerBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServerDatabase) DeepCopyInto(out *UnifiedPushServerDatabase) {
	*out = *in
	out.Port = in.Port
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServerDatabase.
func (in *UnifiedPushServerDatabase) DeepCopy() *UnifiedPushServerDatabase {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServerDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServerList) DeepCopyInto(out *UnifiedPushServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UnifiedPushServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServerList.
func (in *UnifiedPushServerList) DeepCopy() *UnifiedPushServerList {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnifiedPushServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServerSpec) DeepCopyInto(out *UnifiedPushServerSpec) {
	*out = *in
	out.Database = in.Database
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]UnifiedPushServerBackup, len(*in))
		copy(*out, *in)
	}
	in.UnifiedPushResourceRequirements.DeepCopyInto(&out.UnifiedPushResourceRequirements)
	in.OAuthResourceRequirements.DeepCopyInto(&out.OAuthResourceRequirements)
	in.PostgresResourceRequirements.DeepCopyInto(&out.PostgresResourceRequirements)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServerSpec.
func (in *UnifiedPushServerSpec) DeepCopy() *UnifiedPushServerSpec {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServerStatus) DeepCopyInto(out *UnifiedPushServerStatus) {
	*out = *in
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = new(bool)
		**out = **in
	}
	if in.SecondaryResources != nil {
		in, out := &in.SecondaryResources, &out.SecondaryResources
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServerStatus.
func (in *UnifiedPushServerStatus) DeepCopy() *UnifiedPushServerStatus {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServerStatus)
	in.DeepCopyInto(out)
	return out
}
