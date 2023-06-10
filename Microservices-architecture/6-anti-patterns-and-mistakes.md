# When not to use a microservice

1. **small systems**: small systems with low complexity are better of being a monolith. Microservices add complexity.
A system is considered 'small' if the service mapping results in 2-3 services. In this case, we are better off with monolithic approach.
2. **intermingled functionality**: when functionalities cant be separated from each service(even after serious effort), its better not to go with microservice approach.
3. **performance-oriented systems**: microservices have inter-service network calls and message bus(events and their processing) which adds to the performance overhead.
4. **no planned updates**: when systems one-time development oriented. its better off using a monolith. usually ownership of MS is handled by the dev team. when no updates are planned, dev teams can build it and move to another project. Also since no updates, the salient features of short update cycles cant be leverged.

## Anti-patterns

1. **No well defined Services**: This results bloating, unwanted dependencies and uncontrolled growing of service and even its responsibilities.
2. **No well defined API**: API approach should be well thought of and standardised across the platform. Be it REST, gRPC or graphQL.
3. **Deprioritizing Cross Cutting Services**: services like user mgmt, Authz and AuthN, Logging should be implemented first. Although low on businnes priority and value. These should neber be pushed down.
4. **Expanding Service Boundaries**: Adding new functionalities outside the scope of the service is highly discouraged.
