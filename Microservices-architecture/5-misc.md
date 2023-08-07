# Logging and Monitoring

When flow goes through multiple processes jumping across services, its crucial to get a wholistic view of whats happening in the entire flow and within each microservices.

| Logging   | Monitoring |
| --------- | -------    |
| Recording the system’s activity| Based on system’s metrics |
| Audit-trace| Alerting when needed|
| Documenting errors | |

## Logging approach in microservices

Essential Logging parameters

- timestamp
- user
- Severity
- Service
- Message
- Stack trace
- correlation ID( like a datadog trace-id)
