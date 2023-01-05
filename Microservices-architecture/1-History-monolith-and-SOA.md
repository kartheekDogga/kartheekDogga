# History of Microservices

Microservices are a result of problems faced with Monolith and SOA(Service-Oriented Architecture)

## Monolith Application

### features

- All components are executed in a single process
- No Distribution of any kind
- strong coupling between all classes/features

### Pros

- Easier to design
- No network interaction, no messaging, no cross-process debugging
- No serialization/de-serialization processes, monoliths if designed correctly can be quite performant.

### Cons

- Monolithic systems lack the agility and flexibility that many modern businesses require. Because all functions and services are locked into each other, it’s hard to optimize any one function without taking the entire application apart. This makes them hard to scale.
- Less reusability
- Large code base; tough for developers and QA to understand the code and business knowledge
- Less Scalable
- Does not follow SRP (Single Responsibility Principle)
- More deployment and restart times
  
## SOA(Service-Oriented Architecture)

SOA-based applications are distributed multi-tier applications that have presentation, business logic, and persistence layers. Services are the building blocks of SOA applications. While any functionality can be made into a service, the challenge is to define a service interface that is at the right level of abstraction. Services should provide coarse-grained functionality.

- SOA-based computing packages functionalities into a set of interoperable services, which can be integrated into different software systems belonging to separate business domains.
- SOA allows users to combine a large number of facilities from existing services to form applications.
- SOA encompasses a set of design principles that structure system development and provide means for integrating components into a coherent and decentralized system.

There are two major roles within Service-oriented Architecture:

__Service provider__: The service provider is the maintainer of the service and the organization that makes available one or more services for others to use. To advertise services, the provider can publish them in a registry, together with a service contract that specifies the nature of the service, how to use it, the requirements for the service, and the fees charged.

__Service consumer__: The service consumer can locate the service metadata in the registry and develop the required client components to bind and use the service.

![SampleDesign](https://miro.medium.com/max/710/1*AmQGD2rxKvKBknYw9dJDCw.png)

[__MORE ABOUT SOA__](https://www.geeksforgeeks.org/service-oriented-architecture/)

### pros

- Maintenance is Easy
- Quality of Code Improved
- Platform Independence
- Independent of Other Services
  
### cons

In SOA, ESBs(Enterprise Service Bus) are the main pain-point in its architecture, They handled, validation, routing and authentication all in one place. This caused them to be difficult to maintain, and required high resources(both manual and technical).

- __Extra overload__: In SOA, all inputs are validated before it is sent to the service. If you are using multiple services then it will overload your system with extra computation.
- SOA is costly in terms of human resource, development, and technology.
- __High bandwidth server__: As some web service sends and receives messages and information frequently so it easily reaches a million requests per day. So it involves a high-speed server with a lot of data bandwidth to run a web service.
