package mesos

type MesosRegistrar struct {
	Master struct {
		Info struct {
			Hostname string `json:"hostname"`
			Id       string `json:"id"`
			Ip       int    `json:"ip"`
			Pid      string `json:"pid"`
			Port     int    `json:"port"`
		} `json:"info"`
	} `json:"master"`
	Slaves struct {
		Slaves []struct {
			Info struct {
				Checkpoint bool   `json:"checkpoint"`
				Hostname   string `json:"hostname"`
				Id         struct {
					Value string `json:"value"`
				} `json:"id"`
				Port      int `json:"port"`
				Resources []struct {
					Name   string `json:"name"`
					Role   string `json:"role"`
					Scalar struct {
						Value int `json:"value"`
					} `json:"scalar"`
					Type string `json:"type"`
				} `json:"resources"`
			} `json:"info"`
		} `json:"slaves"`
	} `json:"slaves"`
}

type MesosStats struct {
	AvgLoad15min  float64 `json:"avg_load_15min"`
	AvgLoad1min   float64 `json:"avg_load_1min"`
	AvgLoad5min   float64 `json:"avg_load_5min"`
	CpusTotal     int     `json:"cpus_total"`
	MemFreeBytes  int     `json:"mem_free_bytes"`
	MemTotalBytes int     `json:"mem_total_bytes"`
}

type MesosMetrics struct {
	MasterCpusPercent                         float64 `json:"master__cpus_percent"`
	MasterCpusTotal                           float64 `json:"master__cpus_total"`
	MasterCpusUsed                            float64 `json:"master__cpus_used"`
	MasterDiskPercent                         float64 `json:"master__disk_percent"`
	MasterDiskTotal                           float64 `json:"master__disk_total"`
	MasterDiskUsed                            float64 `json:"master__disk_used"`
	MasterDroppedMessages                     float64 `json:"master__dropped_messages"`
	MasterElected                             float64 `json:"master__elected"`
	MasterEventQueueDispatches                float64 `json:"master__event_queue_dispatches"`
	MasterEventQueueHttpRequests              float64 `json:"master__event_queue_http_requests"`
	MasterEventQueueMessages                  float64 `json:"master__event_queue_messages"`
	MasterFrameworksActive                    float64 `json:"master__frameworks_active"`
	MasterFrameworksConnected                 float64 `json:"master__frameworks_connected"`
	MasterFrameworksDisconnected              float64 `json:"master__frameworks_disconnected"`
	MasterFrameworksInactive                  float64 `json:"master__frameworks_inactive"`
	MasterInvalidFrameworkToExecutorMessages  float64 `json:"master__invalid_framework_to_executor_messages"`
	MasterInvalidStatusUpdateAcknowledgements float64 `json:"master__invalid_status_update_acknowledgements"`
	MasterInvalidStatusUpdates                float64 `json:"master__invalid_status_updates"`
	MasterMemPercent                          float64 `json:"master__mem_percent"`
	MasterMemTotal                            float64 `json:"master__mem_total"`
	MasterMemUsed                             float64 `json:"master__mem_used"`
	MasterMessagesAuthenticate                float64 `json:"master__messages_authenticate"`
	MasterMessagesDeactivateFramework         float64 `json:"master__messages_deactivate_framework"`
	MasterMessagesDeclineOffers               float64 `json:"master__messages_decline_offers"`
	MasterMessagesExitedExecutor              float64 `json:"master__messages_exited_executor"`
	MasterMessagesFrameworkToExecutor         float64 `json:"master__messages_framework_to_executor"`
	MasterMessagesKillTask                    float64 `json:"master__messages_kill_task"`
	MasterMessagesLaunchTasks                 float64 `json:"master__messages_launch_tasks"`
	MasterMessagesReconcileTasks              float64 `json:"master__messages_reconcile_tasks"`
	MasterMessagesRegisterFramework           float64 `json:"master__messages_register_framework"`
	MasterMessagesRegisterSlave               float64 `json:"master__messages_register_slave"`
	MasterMessagesReregisterFramework         float64 `json:"master__messages_reregister_framework"`
	MasterMessagesReregisterSlave             float64 `json:"master__messages_reregister_slave"`
	MasterMessagesResourceRequest             float64 `json:"master__messages_resource_request"`
	MasterMessagesReviveOffers                float64 `json:"master__messages_revive_offers"`
	MasterMessagesStatusUpdate                float64 `json:"master__messages_status_update"`
	MasterMessagesStatusUpdateAcknowledgement float64 `json:"master__messages_status_update_acknowledgement"`
	MasterMessagesUnregisterFramework         float64 `json:"master__messages_unregister_framework"`
	MasterMessagesUnregisterSlave             float64 `json:"master__messages_unregister_slave"`
	MasterOutstandingOffers                   float64 `json:"master__outstanding_offers"`
	MasterRecoverySlaveRemovals               float64 `json:"master__recovery_slave_removals"`
	MasterSlaveRegistrations                  float64 `json:"master__slave_registrations"`
	MasterSlaveRemovals                       float64 `json:"master__slave_removals"`
	MasterSlaveReregistrations                float64 `json:"master__slave_reregistrations"`
	MasterSlavesActive                        float64 `json:"master__slaves_active"`
	MasterSlavesConnected                     float64 `json:"master__slaves_connected"`
	MasterSlavesDisconnected                  float64 `json:"master__slaves_disconnected"`
	MasterSlavesInactive                      float64 `json:"master__slaves_inactive"`
	MasterTasksFailed                         float64 `json:"master__tasks_failed"`
	MasterTasksFinished                       float64 `json:"master__tasks_finished"`
	MasterTasksKilled                         float64 `json:"master__tasks_killed"`
	MasterTasksLost                           float64 `json:"master__tasks_lost"`
	MasterTasksRunning                        float64 `json:"master__tasks_running"`
	MasterTasksStaging                        float64 `json:"master__tasks_staging"`
	MasterTasksStarting                       float64 `json:"master__tasks_starting"`
	MasterUptimeSecs                          float64 `json:"master__uptime_secs"`
	MasterValidFrameworkToExecutorMessages    float64 `json:"master__valid_framework_to_executor_messages"`
	MasterValidStatusUpdateAcknowledgements   float64 `json:"master__valid_status_update_acknowledgements"`
	MasterValidStatusUpdates                  float64 `json:"master__valid_status_updates"`
	RegistrarQueuedOperations                 float64 `json:"registrar__queued_operations"`
	RegistrarRegistrySizeBytes                float64 `json:"registrar__registry_size_bytes"`
	RegistrarStateFetchMs                     float64 `json:"registrar__state_fetch_ms"`
	RegistrarStateStoreMs                     float64 `json:"registrar__state_store_ms"`
	RegistrarStateStoreMsCount                float64 `json:"registrar__state_store_ms__count"`
	RegistrarStateStoreMsMax                  float64 `json:"registrar__state_store_ms__max"`
	RegistrarStateStoreMsMin                  float64 `json:"registrar__state_store_ms__min"`
	RegistrarStateStoreMsP50                  float64 `json:"registrar__state_store_ms__p50"`
	RegistrarStateStoreMsP90                  float64 `json:"registrar__state_store_ms__p90"`
	RegistrarStateStoreMsP95                  float64 `json:"registrar__state_store_ms__p95"`
	RegistrarStateStoreMsP99                  float64 `json:"registrar__state_store_ms__p99"`
	RegistrarStateStoreMsP999                 float64 `json:"registrar__state_store_ms__p999"`
	RegistrarStateStoreMsP9999                float64 `json:"registrar__state_store_ms__p9999"`
	SystemCpusTotal                           float64 `json:"system__cpus_total"`
	SystemLoad15min                           float64 `json:"system__load_15min"`
	SystemLoad1min                            float64 `json:"system__load_1min"`
	SystemLoad5min                            float64 `json:"system__load_5min"`
	SystemMemFreeBytes                        float64 `json:"system__mem_free_bytes"`
	SystemMemTotalBytes                       float64 `json:"system__mem_total_bytes"`
}

type UpgradeStrategyStruct struct {
	MinimumHealthCapacity float64 `json:"minimumHealthCapacity,omitempty"`
}

type PortMappingsStruct struct {
	ContainerPort int    `json:"containerPort,omitempty"`
	HostPort      int    `json:"hostPort,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	ServicePort   int    `json:"servicePort,omitempty"`
}

type DockerStruct struct {
	Image        string               `json:"image,omitempty"`
	Network      string               `json:"network,omitempty"`
	Parameters   map[string]string    `json:"parameters,omitempty"`
	PortMappings []PortMappingsStruct `json:"portMappings,omitempty"`
	Privileged   bool                 `json:"privileged,omitempty"`
}

type VolumesStruct struct {
	ContainerPath string `json:"containerPath,omitempty"`
	HostPath      string `json:"hostPath,omitempty"`
	Mode          string `json:"mode,omitempty"`
}

type ContainerStruct struct {
	Docker  DockerStruct    `json:"docker,omitempty"`
	Type    string          `json:"type,omitempty"`
	Volumes []VolumesStruct `json:"volumes,omitempty"`
}

type DeploymentsStruct struct {
	Id string `json:"id,omitempty"`
}

type CommandStruct struct {
	Value string `json:"value,omitempty"`
}

type HealthChecksStruct struct {
	Command                CommandStruct `json:"command,omitempty"`
	GracePeriodSeconds     int           `json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds        int           `json:"intervalSeconds,omitempty"`
	MaxConsecutiveFailures int           `json:"maxConsecutiveFailures,omitempty"`
	Path                   string        `json:"path,omitempty"`
	PortIndex              int           `json:"portIndex,omitempty"`
	Protocol               string        `json:"protocol,omitempty"`
	TimeoutSeconds         int           `json:"timeoutSeconds,omitempty"`
}

type MesosphereContainerStruct struct {
	Args            []string              `json:"args,omitempty"`
	BackoffFactor   int                   `json:"backoffFactor,omitempty"`
	BackoffSeconds  int                   `json:"backoffSeconds,omitempty"`
	Cmd             string                `json:"cmd,omitempty"`
	Constraints     []interface{}         `json:"constraints,omitempty"`
	Container       ContainerStruct       `json:"container,omitempty"`
	Cpus            float64               `json:"cpus,omitempty"`
	Dependencies    []string              `json:"dependencies,omitempty"`
	Deployments     []DeploymentsStruct   `json:"deployments,omitempty"`
	Env             map[string]string     `json:"env,omitempty"`
	Executor        string                `json:"executor,omitempty"`
	HealthChecks    []HealthChecksStruct  `json:"healthChecks,omitempty"`
	Id              string                `json:"id,omitempty"`
	Instances       int                   `json:"instances,omitempty"`
	Mem             float64               `json:"mem,omitempty"`
	Ports           []int                 `json:"ports,omitempty"`
	TasksRunning    int                   `json:"tasksRunning,omitempty"`
	TasksStaged     int                   `json:"tasksStaged,omitempty"`
	UpgradeStrategy UpgradeStrategyStruct `json:"upgradeStrategy,omitempty"`
	Uris            []string              `json:"uris,omitempty"`
	Version         string                `json:"version,omitempty"`
}

type TasksStruct struct {
	AppId        string `json:"appId"`
	Host         string `json:"host"`
	Id           string `json:"id"`
	Ports        []int  `json:"ports"`
	ServicePorts []int  `json:"servicePorts"`
	StagedAt     string `json:"stagedAt"`
	StartedAt    string `json:"startedAt"`
	Version      string `json:"version"`
}

type MesosAppTaskStruct struct {
	Tasks []TasksStruct `json:"tasks"`
}
