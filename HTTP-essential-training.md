# HTTP vs HTTP2 vs HTTPS

|| HTTP 1.0        | HTTP 1.1 | HTTP 2.0 |
| -------------                                 | -------------                                 | -------------                                    |  -------------                            |
|**FEARURES**|For every TCP connection there is only one request and one response.|It supports connection reuse i.e. for every TCP connection there could be multiple requests and responses, and pipelining where the client can request several resources from the server at once. However, pipelining was hard to implement due to issues such as head-of-line blocking and was not a feasible solution.| Uses multiplexing, where over a single TCP connection resources to be delivered are interleaved and arrive at the client almost at the same time. It is done using streams which can be prioritized, can have dependencies and individual flow control.\\ It also provides a feature called server push that allows the server to send data that the client will need but has not yet requested.|
|**AUTHENTICATION**|Uses basic authentication scheme which is unsafe since username and passwords are transmitted in clear text or base64 encoded.|It is relatively secure since it uses digest authentication, [NTLM](https://learn.microsoft.com/en-us/windows/win32/secauthn/microsoft-ntlm) authentication.| Security concerns from previous versions will continue to be seen in HTTP/2. However, it is better equipped to deal with them due to new TLS features like connection error of type Inadequate_Security.|
|**CACHING**|Provides support for caching via the If-Modified-Since header.|Expands on the caching support by using additional headers like cache-control, conditional headers like If-Match and by using entity tags|HTTP/2 does not change much in terms of caching. With the server push feature if the client finds the resources are already present in the cache, it can cancel the pushed stream.|
|**WEB TRAFFIC**|HTTP/1.1 provides faster delivery of web pages and reduces web traffic as compared to HTTP/1.0. However, TCP starts slowly and with domain sharding (resources can be downloaded simultaneously by using multiple domains), connection reuse and pipelining, there is an increased risk of network congestion.||HTTP/2 utilizes multiplexing and server push to effectively reduce the page load time by a greater margin along with being less sensitive to network delays.|



Before knowing difference between HTTP and HTTPS, it is important to know about \
**TLS**([**What is TLS?**](https://www.cloudflare.com/en-in/learning/ssl/transport-layer-security-tls/)),\
[**TLS 1.3 vs 1.2**](https://www.cloudflare.com/en-in/learning/ssl/why-use-tls-1.3/),\
[**TLS 1.3 in detail**](https://blog.cloudflare.com/rfc-8446-aka-tls-1-3/).\

Above has been also explained in the [following video](https://www.youtube.com/watch?v=AlE5X1NlHgg)

[An article on HTTP vs HTTPS](https://www.cloudflare.com/en-in/learning/ssl/why-is-http-not-secure/)

## HTTP STATUS MESSAGES

This should help understand all available status codes  
<https://developer.mozilla.org/en-US/docs/Web/HTTP/Status>

## REQ-RES HEADERS (IN-DEPTH)

## COOKIES

## CACHING - IN-DEPTH(TODO)

## CORS
