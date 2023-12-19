/*
Copyright 2023 IONOS Cloud.

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

package v1alpha1

import (
	"github.com/ionos-cloud/cluster-api-provider-proxmox/api/v1alpha2"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this ProxmoxMachineTemplate to the Hub version (v1alpha2).
func (src *ProxmoxMachineTemplate) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*v1alpha2.ProxmoxMachineTemplate)
	if err := Convert_v1alpha1_ProxmoxMachineTemplate_To_v1alpha2_ProxmoxMachineTemplate(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &v1alpha2.ProxmoxMachineTemplate{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha1) to this version.
func (dst *ProxmoxMachineTemplate) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*v1alpha2.ProxmoxMachineTemplate)
	if err := Convert_v1alpha2_ProxmoxMachineTemplate_To_v1alpha1_ProxmoxMachineTemplate(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}
	return nil
}

// ConvertTo converts this ProxmoxMachineTemplateList to the Hub version (v1alpha2).
func (src *ProxmoxMachineTemplateList) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*v1alpha2.ProxmoxMachineTemplateList)
	return Convert_v1alpha1_ProxmoxMachineTemplateList_To_v1alpha2_ProxmoxMachineTemplateList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1alpha1) to this version.
func (dst *ProxmoxMachineTemplateList) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*v1alpha2.ProxmoxMachineTemplateList)
	return Convert_v1alpha2_ProxmoxMachineTemplateList_To_v1alpha1_ProxmoxMachineTemplateList(src, dst, nil)
}