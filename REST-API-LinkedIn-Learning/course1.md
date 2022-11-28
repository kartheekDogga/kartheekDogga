# Learning REST APIs

## URI vs URL vs URN

`` TODO:
    there's confusion whether to use 'universal' or 'uniform' in the abbreviation of the acronym. This needs to be cleared.
``

- __URI__(Uniform Resource Identifier): It identifies a resource either by location, or a name, or both.<br /> <br /> 
eg: http://www.example.org/Icons/home <br />
In the above case, home icon is identified by it's location, which is pretty unique, however no details about the type of resource its pointing is provided.(.jpg or .png or .GIF).

- __URN__(Uniform Resource Name): It identifies a resource by name in a given namespace. A namespace refers a group of names or identifiers. A simple real-life example would be the usage of our last names.<br />
URN does not define how the resource may be obtained.<br /> <br /> 
eg: "John Smith" is a very common name, it’s not globally unique and would not be appropriate as a URN. However, "John Marhshton Smith Jr." could potentially point to a unique person in a given namespace.

- __URL__(Uniform Resource Locator): it is a specialization of URI that defines the network location of a specific resource. Unlike a URN, the URL defines how the resource can be obtained.<br />
All URLs are URIs, but not all URIs are URLs. A URL is a specialization of URI that defines the location of a specific representation for a given resource.<br /><br />
eg: From the __URI__ example, there could be multiple URLs possible,
http://www.example.org/Icons/home.gif
http://www.example.org/Icons/home.png<br />


    ![Img](https://prateekvjoshi.files.wordpress.com/2014/02/uri-vs-url-vs-urn.jpg?w=300&h=99)

## Constraints of REST-API

[REST API Architectural Constraints](https://www.geeksforgeeks.org/rest-api-architectural-constraints/)

1. __Stateless__: The server should not store/care about the state of the client. The entire necessary details should be present in the request itself in form of query params, or headers, or
URI itself. Statelessness enables greater availability since the server does not have to maintain, update or communicate that session state.<br /><br /> There is a drawback when the client need to send too much data to the server so it reduces the scope of network optimization and requires more bandwidth.

2. __Cacheable__: Every response should include whether the response is cacheable or not and for how much duration responses can be cached at the client side. Client will return the data from its cache for any subsequent request and there would be no need to send the request again to the server. A well-managed caching partially or completely eliminates some client–server interactions, further improving availability and performance.

3. __Client-Server__: REST application should have a client-server architecture. A Client is someone who is requesting resources and are not concerned with data storage, which remains internal to each server, and server is someone who holds the resources and are not concerned with the user interface or user state. They can evolve independently. Client doesn’t need to know anything about business logic and server doesn’t need to know anything about frontend UI.<br />
A client can be anything, for eg: a browser, mobile application, a CLI,some system application like postman etc.

4. __Layered system__: An application architecture needs to be composed of multiple layers. Each layer doesn’t know any thing about any layer other than that of immediate layer and there can be lot of intermediate servers between client and the end server. Intermediary servers may improve system availability by enabling load-balancing and by providing shared caches.<br />
Between the client who requests a representation of a resource’s state, and the server who sends the response back, there might be a number of servers in the middle. These servers might provide a security layer, a caching layer, a load-balancing layer, or other functionality. Those layers should not affect the request or the response. The client is agnostic as to how many layers, if any, there are between the client and the actual server responding to the request. Layering can also be used for caching.<br /><br />
Important properties:

    1. _simplifies clients_: similar to the client-server constraint, this constraint improves simplicity by separating concerns;<br />
    2. _shared caching_: placing shared caches at boundaries of organizational domain can result in significant benefits.<br />
    3. _scalability_: intermediaries can be used to improve system scalability by enabling load balancing<br />
    4. _lecagy encapsulation_ : can be used to encapsulate legacy services or protect new services from legacy clients.<br />
    
    Important trade-off:

 	1. _latency_: adds overhead and latency and can reduce user perceived performance;

![Img](https://restapilinks.com/wp-content/uploads/2021/03/layered.png)

5. __Uniform Interface__: It is a key constraint that differentiate between a REST API and Non-REST API. It suggests that there should be an uniform way of interacting with a given server irrespective of device or type of application (website, mobile app).<br /><br />
There are four guidelines principle of Uniform Interface are:<br /><br />
    1. __Resource-Based__: Individual resources are identified in requests. For example: API/users.
    2. __Manipulation of Resources Through Representations__: Client has representation of resource and it contains enough information to modify or delete the resource on the server, provided it has permission to do so. Example: Usually user get a user id when user request for a list of users and then use that id to delete or modify that particular user.
    3. __Self-descriptive Messages__: Each message includes enough information to describe how to process the message so that server can easily analyses the request.
    4. __Hypermedia as the Engine of Application State (HATEOAS)__: It need to include links for each response so that client can discover other resources easily.


## HTTP METHODS

[Anatomy of REST Request REFERENCE 1](https://medium.com/swlh/an-overview-of-restful-apis-2df8ffa8c803#:~:text=one%20we%20made.-,Anatomy%20of%20REST%3A,-now%20we%20know)

[Anatomy of REST Request REFERENCE 2](https://www.numpyninja.com/post/anatomy-of-a-rest-request)

| Method        | Description   |
| ------------- | ------------- |
|GET	        | Retrieve information about the REST API resource |
|POST	        | Create a REST API resource |
|PUT	        | Update a REST API resource |
|DELETE	        | Delete a REST API resource or related component |