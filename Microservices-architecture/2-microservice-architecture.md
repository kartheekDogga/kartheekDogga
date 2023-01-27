# Microservice Architecture

[Reference:Martin-Fowler-Article](https://martinfowler.com/articles/microservices.html)

The microservice architectural style [1] is an approach to developing a single application as a suite of small services, each running in its own process and communicating with lightweight mechanisms, often an HTTP resource API. These services are built around business capabilities and independently deployable by fully automated deployment machinery. There is a bare minimum of centralized management of these services, which may be written in different programming languages and use different data storage technologies.

## Charachterstics of Microservices 

- [Microservice Architecture](#microservice-architecture)
  - [Charachterstics of Microservices](#charachterstics-of-microservices)
    - [__Componentization via Services__](#componentization-via-services)
    - [__Organized around business capabilities__](#organized-around-business-capabilities)
    - [__Products not Projects__](#products-not-projects)
    - [__Smart endpoints and dumb pipes__](#smart-endpoints-and-dumb-pipes)
    - [__Decentralized Governance__](#decentralized-governance)
    - [__Decentralized Data Management__](#decentralized-data-management)
    - [__Infrastructure Automation__](#infrastructure-automation)
    - [__Design For Failure__](#design-for-failure)
    - [__Evolutionary Design__](#evolutionary-design)

### __Componentization via Services__

Microservice architectures will use libraries, but their primary way of componentizing their own software is by breaking down into services. We define libraries as components that are linked into a program and called using in-memory function calls, while services are out-of-process components who communicate with a mechanism such as a web service request, or remote procedure call.

One main reason for using services as components (rather than libraries) is that services are independently deployable. If you have an application [4] that consists of a multiple libraries in a single process, a change to any single component results in having to redeploy the entire application. But if that application is decomposed into multiple services, you can expect many single service changes to only require that service to be redeployed.

Using services like this does have downsides. Remote calls are more expensive than in-process calls, and thus remote APIs need to be coarser-grained, which is often more awkward to use. If you need to change the allocation of responsibilities between components, such movements of behavior are harder to do when you're crossing process boundaries.

### __Organized around business capabilities__

When looking to split a large application into parts, often management focuses on the technology layer, leading to UI teams, server-side logic teams, and database teams. This is slow and cumbersome due to inter-group communication and varying terminologies across teams. Each team could have different goals and often in-coherent.

The microservice approach to division is different, splitting up into services organized around business capability. Such services take a broad-stack implementation of software for that business area, including UI, database and external interactions. Consequently, the teams are cross-functional, including the full range of skills required for the development.

__TLDR__: Instead of splitting across tech boundaries like DB,UI, backend, microservice teams include handling of all domains within itself. The teams are cross-functional with each team having full range of skills.

### __Products not Projects__

Most application development efforts that we see use a project model: where the aim is to deliver some piece of software which is then considered to be completed. On completion the software is handed over to a maintenance organization and the project team that built it is disbanded.

Microservice approach tend to avoid this model, preferring instead the idea that a team should own a product over its full lifetime. A development team takes full responsibility for the software in production. This brings developers into day-to-day contact with how their software behaves in production and increases contact with their users, as they have to take on at least some of the support burden.

The product mentality, ties in with the linkage to business capabilities. Rather than looking at the software as a set of functionality to be completed, there is an on-going relationship where the question is how can software assist its users to enhance the business capability.

There's no reason why this same approach can't be taken with monolithic applications, but the smaller granularity of services can make it easier to create the personal relationships between service developers and their users.

### __Smart endpoints and dumb pipes__

Applications built from microservices aim to be as decoupled and as cohesive as possible - they own their own domain logic. receiving a request, applying logic as appropriate and producing a response. These are achieved using simple RESTish(alternatives: gRPC, graphQL) protocols rather than complex protocols, such as WS-Choreography or BPEL or orchestration by a central tool.

The two protocols used most commonly are HTTP request-response with resource API's and lightweight messaging.

The second approach in common use is messaging over a lightweight message bus. The infrastructure chosen is typically dumb (dumb as in acts as a message router only) - simple implementations such as RabbitMQ or ZeroMQ don't do much more than provide a reliable asynchronous messaging means - the smart-logic still live in the end points that are producing and consuming messages i.e. in the services.

### __Decentralized Governance__

One of the consequences of centralised governance is the tendency to standardise on single technology platforms. Monolithic applications are usually build on a single language\tech-stack where the choice this stack for every problem might not be ideal/best-possible.

Splitting the monolith's components out into services we have a choice when building each of them. You want to use Node.js to standup a simple reports page? Go for it. C++ for a particularly gnarly near-real-time component? Fine. You want to swap in a different flavour of database that better suits the read behaviour of one component? We have the technology to rebuild him.

### __Decentralized Data Management__

Decentralization of data management presents in a number of different ways. At the most abstract level, it means that the conceptual model of the world will differ between systems. This is a common issue when integrating across a large enterprise, the sales view of a customer will differ from the support view. Some things that are called customers in the sales view may not appear at all in the support view. Those that do may have different attributes and (worse) common attributes with subtly different properties.

A useful way of thinking about this is the Domain-Driven Design notion of Bounded Context. DDD divides a complex domain up into multiple bounded contexts and maps out the relationships between them. This process is useful for both monolithic and microservice architectures, but there is a natural correlation between service and context boundaries that helps clarify, and as we described above, business capabilities reinforce the separations.

While monolithic applications prefer a single logical database for persistant data, enterprises often prefer a single database across a range of applications - many of these decisions driven through vendor's commercial models around licensing. Microservices prefer letting each service manage its own database, either different instances of the same database technology, or entirely different database systems - an approach called [Polyglot Persistence](https://martinfowler.com/bliki/PolyglotPersistence.html). You can use polyglot persistence in a monolith, but it appears more frequently with microservices.

### __Infrastructure Automation__

Infrastructure automation techniques have evolved enormously over the last few years - the evolution of the cloud and AWS in particular has reduced the operational complexity of building, deploying and operating microservices.

[sample](https://martinfowler.com/articles/microservices/images/basic-pipeline.png)

### __Design For Failure__

Since services can fail at any time, it's important to be able to detect the failures quickly and, if possible, automatically restore service. Microservice applications put a lot of emphasis on real-time monitoring of the application, checking both architectural elements (how many requests per second is the database getting) and business relevant metrics (such as how many orders per minute are received). Semantic monitoring can provide an early warning system of something going wrong that triggers development teams to follow up and investigate.

This is particularly important to a microservices architecture because the microservice preference towards choreography and event collaboration leads to emergent behavior. While many pundits praise the value of serendipitous emergence, the truth is that emergent behavior can sometimes be a bad thing. Monitoring is vital to spot bad emergent behavior quickly so it can be fixed.

Monoliths can be built to be as transparent as a microservice - in fact, they should be. The difference is that you absolutely need to know when services running in different processes are disconnected. With libraries within the same process this kind of transparency is less likely to be useful.

### __Evolutionary Design__

Microservice practitioners, usually have come from an evolutionary design background and see service decomposition as a further tool to enable application developers to control changes in their application without slowing down change. Change control doesn't necessarily mean change reduction - with the right attitudes and tools you can make frequent, fast, and well-controlled changes to software.

Whenever you try to break a software system into components, you're faced with the decision of how to divide up the pieces - what are the principles on which we decide to slice up our application? The key property of a component is the notion of independent replacement and upgradeability[12] - which implies we look for points where we can imagine rewriting a component without affecting its collaborators.
