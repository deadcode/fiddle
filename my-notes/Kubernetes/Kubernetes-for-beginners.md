# Course Options
1. Kubernetes for Beginners
2. Kubernetes for Admistrators -> Certified Adminstrators exam
3. Kubernetes for (Application) Developers -> Certified Developer exam
   
# Container Orchestration
Multiple orchestration technology options:
* Docker Swarm: simple to use but fewer features
* Mesos: from Apache, fully featured but complex setup
* Kubernetes or K8s: from Google

**What is Kubernetes?**: It is a container orchestration technology used to deploy containers in a distributed + clustered environment.

# Kubernetes Architecture
## **Components**
* Nodes (or minions): Physical nodes 
* Cluster: A set of nodes deployed together
* Master: Another node in the cluster that manages other nodes and orchestrate containers on nodes

## Parts of kubernetes deployment
* API server: Gateway for user interaction
* Scheduler: schedules containers on nodes
* Controller: decides to bring new containers and nodes up/down
* Container runtime: underlying runtime to run the containers (eg. Docker)
* Kubelet (service): Agent that runs on all nodes, make sure containers are running.
* etcd (service) - distributed key-value pair DB
* cubectl (cube-ctl / cube command line tool)

Container runtime options:
* Docker
* RKT (rocket)
* CRI-O (cryo)

Worker nodes run: Kubelets + Container runtime
Master nodes run: Apiserver + etcd + controller + scheduler

### CRI
- container runtime interface
- Kubernetes was designed to only work with Docker. CRI was introduced to add support for other container runtimes eg. Rkt

### OCI
- Open container initiative
- Contains
  - Imagespec : how the container image is built
  - runtimespec: how to build the container runtime
- dockershim - a shim layer built to support docker runtime in kubernetes without the use of CRI

### containerd
- dockers container runtime env
- is seperate project and part of CNCF (Cloud Native Computing Foundation)
- ctr - CLI for debugging containerd
- nerdctl (nerd control) - docker-like CLI for managing containerd
- crictl (cry control) - CLI used to manage the CRI, mostly for debugging.


| Tool -> | CTR | nerdctl | crictl |
| ------- | --- | ------- | -------|
| Purpose | Debugging | General Purpose | Debugging |
| Communnity | containerd | containerd | Kubernetes |
| Works with | containerd| containerd | All CRI compatible runtimes |
| ----- | ----- | ----- | ----- |

# Kubernetes Concepts
What is a pod?
- A pod is an instance of an application (on a node)
- A pod is the smallest object managed by kubernetes
- A pod is usually a 1 container 
  - a pod <-> container relation is usually 1-to-1
- A pod COULD contain more than 1 container but they are usually not the same kind
- containers within the pod share the same network space (talk via localhost) and share the same storage as well