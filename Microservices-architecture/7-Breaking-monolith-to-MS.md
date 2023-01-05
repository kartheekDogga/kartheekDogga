# breaking a monolith to MS

- when MS needs to improve current system
- needs to be thouroghly planned
- prone to high rate of failures

## motivation to breaking down

- Shorten Update cycle
- modularize

basically all the problems that MS solve arising from a monolith

There are 3 approaches to splitting monoliths to MS

1. __New Modules as Services__: we do not modify existing code, we add new modules as services.
   | pros | cons |
   | ---- | ---- |
   | Minimum code changes| Takes time |
   |Easy to implement | End result is not pure Microservices architecture|
2. __Separate Existing Modules as Services__:
   | pros | cons |
   | ---- | ---- |
   |End result is pure Microservices architecture| Takes time |
   || A lot of code changes and regression testing required |
3. __Complete Rewrite__: When system is too old or outdated, a rewrite gives us an oppourtunity to modernize the technology. This results in a new modern system in MS architecture, but this takes time and lot of testing.
