# Prometheus server URL (make sure you port forwarded Prometheus to localhost:9090)
PROMETHEUS_URL = 'http://localhost:9090'

# The steps used for the metrics (in seconds)
METRIC_STEP = 5
# Maximum resolution of Prometheus (i.e. the maximum number of data points that can be returned in a single query)
MAX_RESOLUTION = 11_000
# How far back in time Prometheus should look to calculate the rate of change
# Set to two minutes to get more data points for the rate calculation (1m sometimes gives no data)
DURATION = "2m" 

# TODO: fix queries not working with by, needs to be something different

# Prometheus queries to get relevant energy metrics (same as study from Ivano and colleagues)
# Group by container_label_io_kubernetes_container_name because cadvisor uses this label to identify container names 
# (see TroubleShooting.md for explanation of container label in cadvisor)
QUERIES = {
    "cpu": f"sum(rate(container_cpu_usage_seconds_total[{DURATION}])) by (container_label_io_kubernetes_container_name)",
    "memory": f"sum(rate(container_memory_usage_bytes[{DURATION}])) by (container_label_io_kubernetes_container_name)",
    "memory_rss": f"sum(rate(container_memory_rss[{DURATION}])) by (container_label_io_kubernetes_container_name)",
    "memory_cache": f"sum(rate(container_memory_cache[{DURATION}])) by (container_label_io_kubernetes_container_name)",
    # Total bytes read by the container
    "disk": f"sum(rate(container_fs_reads_bytes_total[{DURATION}])) by (container_label_io_kubernetes_container_name)",
    # Power consumption using Kepler (https://sustainable-computing.io/design/metrics/)
    "power": 'sum by (pod_name, container_name, container_namespace, node)(irate(kepler_container_joules_total{}[{DURATION}]))',
}

# Relevant containers to monitor the energy consumption 
# # (use scripts/describeContainers.sh to print containers in the Kubernetes cluster)
# TODO: replace with containers used by DYNAMOS later (containers from orchestrator, vu, uva, etc.)
CONTAINERS = {
    "etcd",
    "rabbitmq",
    "loki",
    "nginx-ingress",
    "policy",
}


# TODO: remove later, not necessary anymore (now uses queries/metrics from study from Ivano and colleagues)
# # Relevant Prometheus energy metrics
# RELEVANT_PROMETHEUS_CONTAINER_METRICS = [
#     'container_blkio_device_usage_total',
#     'container_cpu_usage_seconds_total',
#     'container_fs_reads_bytes_total',
#     'container_fs_reads_total',
#     'container_fs_writes_bytes_total',
#     'container_fs_writes_total',
#     'container_last_seen',
#     'container_memory_cache',
#     'container_memory_failcnt',
#     'container_memory_failures_total',
#     'container_memory_kernel_usage',
#     'container_memory_max_usage_bytes',
#     'container_memory_rss',
#     'container_memory_usage_bytes',
#     'container_memory_working_set_bytes',
#     'container_oom_events_total',
#     'container_processes',
#     'container_scrape_error',
#     'container_sockets',
#     'container_start_time_seconds',
#     'container_threads',
# ]

# # Relevant Kepler metrics (https://sustainable-computing.io/design/metrics/)
# RELEVANT_KEPLER_CONTAINER_METRICS = [
#     'kepler_container_joules_total',
#     'kepler_container_core_joules_total',
#     'kepler_container_dram_joules_total',
#     'kepler_container_uncore_joules_total',
#     'kepler_container_package_joules_total',
#     'kepler_container_other_joules_total',
#     'kepler_container_gpu_joules_total',
#     'kepler_container_energy_stat',
#     'kepler_container_bpf_cpu_time_us_total',
#     'kepler_container_cpu_cycles_total',
#     'kepler_container_cpu_instructions_total',
#     'kepler_container_cache_miss_total',
#     'kepler_container_cgroupfs_cpu_usage_us_total',
#     'kepler_container_cgroupfs_memory_usage_bytes_total',
#     'kepler_container_cgroupfs_system_cpu_usage_us_total',
#     'kepler_container_cgroupfs_user_cpu_usage_us_total',
#     'kepler_container_bpf_net_tx_irq_total',
#     'kepler_container_bpf_net_rx_irq_total',
#     'kepler_container_bpf_block_irq_total',
#     'kepler_node_info',
#     'kepler_node_core_joules_total',
#     'kepler_node_uncore_joules_total',
#     'kepler_node_dram_joules_total',
#     'kepler_node_package_joules_total',
#     'kepler_node_other_joules_total',
#     'kepler_node_gpu_joules_total',
#     'kepler_node_platform_joules_total',
#     'kepler_node_energy_stat',
#     'kepler_node_accelerator_intel_qat',
# ]
