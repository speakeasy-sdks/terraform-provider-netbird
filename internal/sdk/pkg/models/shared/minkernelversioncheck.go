// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MinKernelVersionCheck - Posture check with the kernel version
type MinKernelVersionCheck struct {
	// Minimum acceptable version
	MinKernelVersion string `json:"min_kernel_version"`
}

func (o *MinKernelVersionCheck) GetMinKernelVersion() string {
	if o == nil {
		return ""
	}
	return o.MinKernelVersion
}
