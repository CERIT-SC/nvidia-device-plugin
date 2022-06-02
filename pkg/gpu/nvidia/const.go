package nvidia

import (
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
)

// MemoryUnit describes GPU Memory, now only supports Gi, Mi
type MemoryUnit string

const (
	resourceName  = "cerit.io/gpu-mem"
	resourceCount = "cerit.io/gpu-count"
	resourceNvidia = "nvidia.com/gpu"
	serverSock    = pluginapi.DevicePluginPath + "ceritioshare.sock"

	OptimisticLockErrorMsg = "the object has been modified; please apply your changes to the latest version and try again"

	allHealthChecks             = "xids"
	containerTypeLabelKey       = "io.kubernetes.docker.type"
	containerTypeLabelSandbox   = "podsandbox"
	containerTypeLabelContainer = "container"
	containerLogPathLabelKey    = "io.kubernetes.container.logpath"
	sandboxIDLabelKey           = "io.kubernetes.sandbox.id"

	envNVGPU                   = "NVIDIA_VISIBLE_DEVICES"
	EnvResourceIndex           = "CERIT_IO_GPU_MEM_IDX"
	EnvResourceByPod           = "CERIT_IO_GPU_MEM_POD"
	EnvResourceByContainer     = "CERIT_IO_GPU_MEM_CONTAINER"
	EnvResourceByDev           = "CERIT_IO_GPU_MEM_DEV"
	EnvAssignedFlag            = "CERIT_IO_GPU_MEM_ASSIGNED"
	EnvResourceAssumeTime      = "CERIT_IO_GPU_MEM_ASSUME_TIME"
	EnvResourceAssignTime      = "CERIT_IO_GPU_MEM_ASSIGN_TIME"
	EnvNodeLabelForDisableCGPU = "cgpu.disable.isolation"

	DeviceListStrategyEnvvar       = "envvar"
	DeviceListStrategyVolumeMounts = "volume-mounts"

	DeviceListHostPath      = "/dev/null"
	DeviceListContainerPath = "/var/run/nvidia-container-devices"

	GiBPrefix = MemoryUnit("GiB")
	MiBPrefix = MemoryUnit("MiB")
)
