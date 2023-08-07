# Service Mesh

Service mesh manages service-service communication. This is usually platform agnostic.

Service-mesh comprises of a set of coponents that sit near the service and manage all service-service communications. It provides all communication services and frees services from dealing with it. The service should interact with service-mesh only.

## When should you use Service Mesh

1. When we have a lot of services
2. when there is a lot of inter-service communication(sync-communication specifically)
3. When there is complex comm network, or various comm protocols(brittle network)

## Problems solved by service-mesh

Microservices communicate betweeen them a lot. The communincation can pose a lot of challenges and problems(eg: req timeouts, sercurity, retries, monitoring).

- protocol conversion
- comm security
- authentication
- reliability(timeouts, retries, health check, **circuit breaking**)
- monitoring
- servcie discovery
- Load balancing

**circuit breaker**: just like the name suggests, this helps in preventing cascading of failures when a service fails.

using service mesh helps offload devs from handling commuication aspects and focus on business logic.

## Service Mesh Architecture

When 2 services want to communicate, we place 2 service-mesh components between each of them which can handle communication. These service-mesh components are called **data-planes**. Data-planes do all the heavy-lifiting of service communication(i.e. communication, auth, retries, timeout, monitoring etc...).

![img](https://cdn.ttgtmedia.com/rms/onlineimages/how_a_service_mesh_works-f.png)

Types of service mesh(data-planes):

1. In-Process: These exist inside the service process itself and the service intercation with mesh is done by simply calling methods within code.
2. Sidecar: These are not part of service process. Here, the service first makes a network call to the mesh component which in turn makes another call to the other mesh component, which relays the call to the recipient service.

| In-Process| Sidecar |
| --------- | ------- |
|Performance| platform-agnostic|
|coupled with code| code-agnostic|

**NOTE**: Usually when referring to mesh components or data-planes we refer them as sidecars/ sidecar-proxies, because they are the more popular choice and also the recommended ones in most cases.

A few service mesh implementations are:

1. [Istio](https://istio.io/)
2. [LinkerD](https://linkerd.io/)
3. [maesh](https://traefik.io/traefik-mesh/)
