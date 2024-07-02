#!/bin/bash

# Check if an argument was provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <chartsPath>"
    exit 1
fi

# First argument is the chartsPath (the path to core folder in DYNAMOS project)
chartsPath="$1"
monitoringChartsPath="$chartsPath/monitoring"
monitoringValues="$monitoringPath/monitoring-values.yaml"

# Create the namespace in the Kubernetes cluster (if not exists)
kubectl create namespace kepler

# Install and add Kepler
helm repo add kepler https://sustainable-computing-io.github.io/kepler-helm-chart
helm repo update
# Install Kepler
# This also creates a service monitor for the prometheus stack
helm upgrade -i kepler kepler/kepler --namespace kepler --set serviceMonitor.enabled=true --set serviceMonitor.labels.release=prometheus 

# After this final installation you should be able to view the Kepler namespace in minikube dashboard
# See EnerConMeasInDYNAMOS.md file for how to run Prometheus and see the metrics.

# Finally, apply/install the monitoring release (will use the monitoring charts,
# which includes the deamonset, service and sesrvicemonitor for cadvisor for example)
helm upgrade -i -f "$monitoringChartsPath" monitoring $monitoringPath