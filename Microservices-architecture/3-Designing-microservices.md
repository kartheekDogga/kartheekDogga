# Designing-microservices

## Mapping the Components

- The most important step in design process
- This determines how the system looks like in the long run
- this once set, it can be very diffivult to change

### What is it?

Components = Services

When talking about components in microservices, we refer to services as opposed to libraries like in monoliths.

Mapping should be based on:

- [Designing-microservices](#designing-microservices)
  - [Mapping the Components](#mapping-the-components)
    - [What is it?](#what-is-it)
    - [Business Requirements](#business-requirements)
    - [Functional Autonomy](#functional-autonomy)
    - [Data Entities](#data-entities)
    - [Data Autonomy](#data-autonomy)
  - [cross-cutting services](#cross-cutting-services)
  - [Designing Communication Patterns](#designing-communication-patterns)

### Business Requirements

The collection of requirements around a specific business capability. capability is the frame of the component and the requirements is the actions it can do.

For e.g. Consider, Orders management: It should include capability to create/update/remove orders. These are the features it should provide from a business perspective. observe that the thechnical complexity is not considered here. This can vary heavily for each different feature/component.

### Functional Autonomy

The maximum functionality that does not involve other business requirements.

For e.g. In orders management, consider these 2 cases:

1. Retrieve the orders made in the last week.
2. Get all orders made by users aged 34-45

while the first case aligns well with the capability of order mgmt, the second doesnt. In  the second case, there is a requirement of user-data(user-ages) which lies outside the service boundaries.

While mapping components, in reality, we can always encounter gray areas where there could be such overlapping. A good design aims to have as less such gray areas as possible.

### Data Entities

 Services should be defined around well-defined data-entites. For example: orders, users, items etc.

 Data can be related to other entites by just ID.

 For e.g. Orders store the customerID which is a mapping to users entity. Note that we are just storing just a refernence to the users entity and not the whole users data.

### Data Autonomy

- underlying data is an atomic unit
- service should not depend on data from other services to function properly

for e.g. if an **employees service** depends on **address service** to return employee data, this means, this always depends on an external service to execute a funtionality. This is a design flaw and needs to be redesigned properly.

If we were to choose between data-replication tight dependency of microservices, its always better to choose a little bit of data replication. Because tight dependency like above mentioned example can cause many problems

## cross-cutting services

there could be some services at times which might not have a business capability but these services can be used by all other services across the system.

ex: logging, caching, user mgmt and auth services are some examples.
such services must be a part of the design and need to be developes before all other services.

## Designing Communication Patterns

Efficient communication patterns is crucial. An inefficient comm. pattern can impact system performance and result in poor error handling.

Some of the main communication patterns in microservices are:

1. **1-1 Synchronours communication**: used when a service needs a response from ext source to continue the process.
It has to be noted that this is diferent from tight-coupling microservices, which is a design flaw. Many a times, there could be a significant portion of funtionality that falls into business capability of other service, but the client(caller-service) needs a response(success/failure/or some data) to take next logical action. \
one good example of this scenario is user/order service awaiting payment service to respond to payment-request with a response.
2. **1-1 Asynchronours communication**: This process can be thought of like fire-and-forget.
3. **pub-sub/Event-driven**: unlike async communication, the producer service has no idea who the consumer/receiver of the event is. often these events are of one-to-many or many-to-many relationship.
