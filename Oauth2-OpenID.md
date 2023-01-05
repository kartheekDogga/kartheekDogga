# OAUTH2 and OPENID connect

- JWT

- Token revocation

- Token introspection

  - Examines a token to describe its contents
  - useful for opaque tokens(instead of JWT)
  - desccribes if token is active or not
  - mandatory if you have token revocation (query auth server whether token is valid or revoked)\

- Dynamic Client Registration(RFC 7591)

- Authorization server metadata discovery(RFC 8414)

## Authorization Grant types

1. Authorization code
2. Implicit or authorization code with PKCE - mobile apps(with js backend)
3. client credential - service accounts or microservices
4. resource owner password - legacy apps(not used frequently)

## scope naming patterns

1. simple scopes[simple strings](user_read, public_repo, user:rw)
2. java style namespaces [com.app.resource, com.app.resource.read, com.app.resource.subresource.write ....etc]
3. URL style namespaces(quite similar to java style) [<https://api.company.com/resource>, <https://api.company.com/resource.read>, <https://api.company.com/resource/subresource.write>]
