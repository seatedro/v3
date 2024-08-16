---
author:
  name: Rohit Prakash
date: 2021-09-19
summary: K8s explained.
tags:
  - cloud
title: What is K8s? ☸️
---

![K8s](/static/images/k8s.webp)
> “Kubernetes is an open-source container-orchestration system for automating computer application deployment, scaling, and management.”

When I first read that explanation, I could not understand a word. If you can’t either, then don’t worry, you’re not alone.

If you want some advice, read this
> "You don't need k8s for your 0 user hobby project"

If you just came here to learn, keep reading <3.

Kubernetes (or k8s) is a platform that automatically manages your application which is usually split into micro-services in the form of containers. Kubernetes provides the necessary resources needed to run these containers with the help of Load Balancing, Heartbeat Monitors, Horizontal and Vertical Scaling, Automated Rollout, and several other features.

Usually, you would have to deploy your application on a server, and when you eventually have more people using your service, it would be a hassle to scale your application since you would have to do everything manually. With k8s, most of the difficult manual aspects are now automated.
## Simplifying the important k8s jargon
### Cluster

Clusters consist of the entire k8s architecture, essentially it is a series of nodes connected together.
### Pod

A pod holds one or more container(s). Pods are the simplest unit that exists within Kubernetes.
### Node

Nodes are the hardware components. They can be Physical Servers or Virtual Machines. Each node will automatically assign pods to itself based on resource availability.
### Deployment

A deployment is a set of rules that state how the container must be run in the cluster. Some of the important things specified in deployments are :

- **Replicas**: Defines the number of replicas to be made for the container.
- **Image**: Specifies the container image to be used, this is updated in future deployments.
- **ImagePullPolicy**: Specifies how often a container image is pulled from the container registry.

### Service

This is an object that specifies how a deployment or another resource can be accessed.

There are 4 different types of Services, (ServiceTypes):

- **ClusterIP**: Allocates a cluster IP (internal IP) to a resource and doesn’t allow external access. This is the default configuration.
- **NodePort**: Allocates a static port for the service on all the Nodes. You can either specify this or let k8s pick one for you. The range of ports is 30000–32767.
- **LoadBalancer**: Creates an External Load Balancer on the cloud platform that k8s is running ( GCP, AWS, etc. ), and this allows external access to the service.
    ExternalName: Instead of using k8s selectors, it uses domain names instead.

More info on these types can be found here.
## Benefits of k8s

- **Automatic Load Balancing**: Kubernetes automatically balances the incoming load to the micro-service by spreading it out across nodes so that a single node doesn’t handle all the requests.
- **Self-Healing**: When a container fails, k8s automatically procures a new pod to provide failover with minimum possible downtime.
- **Automated Rollout and Rollbacks**: When a new release needs to be deployed, k8s can rollout this release while keeping the service running without downtime. Similarly, it can provide support for rolling back to a previous release.
- **Secrets and ConfigMaps**: These are additional configurations where you can store sensitive data like API keys, OAuth keys, passwords, etc. These can then be made available to the containers. When any update needs to be done to this data, you can do so without having to rebuild the containers.

## Drawbacks

- Quite unnecessary for small applications: When starting with a new app that doesn’t necessarily have that much traffic, Kubernetes is an extra hassle for the developers.
- Migrating to k8s is painstaking: Existing software will have to be modified to accommodate the k8s architecture.
- **Expensive**: Due to all the automated features that it provides, the costs of running a k8s cluster can add up quickly.

# Conclusion

With all that being said, you do not need Kubernetes in your zero user side project. Please stop.
