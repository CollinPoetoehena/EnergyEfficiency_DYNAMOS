# Caching in Kubernetes
This guide explains how to set up caching in Kubernetes.

## Installing Redis
Follow these steps to install Redis.
```sh
# Go to the scripts path
cd energy-efficiency/scripts
# Make the script executable (needs to be done once)
chmod +x caching.sh
# Execute the script:
./caching.sh
```
Afterwards you should be able to see the caching namespace and the redis service. You can view the logs in the pod in Kubernetes Dashboard or k9s and it should say something like:
```sh
2025-01-14T15:18:29.759Z | 1:C 14 Jan 2025 15:18:29.759 * oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
2025-01-14T15:18:29.759Z | 1:C 14 Jan 2025 15:18:29.759 * Redis version=7.4.2, bits=64, commit=00000000, modified=0, pid=1, just started
2025-01-14T15:18:29.759Z | 1:C 14 Jan 2025 15:18:29.759 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
2025-01-14T15:18:29.759Z | 1:M 14 Jan 2025 15:18:29.759 * monotonic clock: POSIX clock_gettime
2025-01-14T15:18:29.759Z | 1:M 14 Jan 2025 15:18:29.759 * Running mode=standalone, port=6379.
2025-01-14T15:18:29.760Z | 1:M 14 Jan 2025 15:18:29.760 * Server initialized
2025-01-14T15:18:29.760Z | 1:M 14 Jan 2025 15:18:29.760 * Ready to accept connections tcp  
```

Further steps:
```sh
# Install redis tools so that you can use the redis-cli in WSL
sudo apt update
sudo apt install redis-tools

# Port-forward redis service in Kubernetes (when redis pod and service in kubernetes are running):
kubectl port-forward svc/redis -n caching 6379:6379
# Test connectivity with redis (it is exposed through a node port in redis.yaml):
redis-cli -h localhost -p 30006 ping
# Expected result: PONG
```
Also, test with Go code for example, see go/agent/main.go for an example of how this is done.