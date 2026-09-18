---
source: "https://swobu.com/blog/https-routes/"
hn_url: "https://news.ycombinator.com/item?id=49756982"
title: "Share an LLM endpoint over HTTPS while TLS still terminates on your machine"
article_title: "Share an LLM endpoint over HTTPS while TLS still terminates on your machine — Swobu Journal"
image: "https://swobu.com/og-image.png"
author: "metrofun"
captured_at: "2026-09-18T17:32:03Z"
capture_tool: "hn-digest"
hn_id: 49756982
score: 1
comments: 0
posted_at: "2026-09-18T16:48:59Z"
tags:
  - hacker-news
---

# Share an LLM endpoint over HTTPS while TLS still terminates on your machine

- HN: [49756982](https://news.ycombinator.com/item?id=49756982)
- Source: [swobu.com](https://swobu.com/blog/https-routes/)
- Score: 1
- Comments: 0
- Posted: 2026-09-18T16:48:59Z

## Translation

Title: Share an LLM endpoint over HTTPS while TLS still terminates on your machine
Article title: Share an LLM endpoint over HTTPS while TLS still terminates on your machine — Swobu Journal
Description: Share a stable, bearer-protected LLM endpoint without moving provider credentials. Application TLS still terminates on the machine running Swobu.

Article text:
Share an LLM endpoint over HTTPS while TLS still terminates on your machine — Swobu Journal SWOBU / JOURNAL What it does Docs Journal GitHub ↗ PRODUCT · HTTPS ROUTES 14 SEPT 2026 · UPDATED 16 SEPT 2026 · Swobu COPY LINK Share an LLM endpoint over HTTPS while TLS still terminates on your machine
Share a stable, bearer-protected LLM endpoint without moving provider credentials. Application TLS still terminates on the machine running Swobu.
I wanted to use the LLM setup on my workstation from a phone and another machine without copying provider credentials onto either of them.
I also did not want the relay that made the machine reachable to become the application TLS endpoint.
Terminal window swobu share dev/coding
That command creates a bearer for the coding route and returns a public HTTPS endpoint. A remote client can use OpenAI-compatible or Anthropic-compatible ingress for the protocols Swobu already supports. It does not need Swobu installed.
The remote machine gets the endpoint and the scoped Share bearer. Provider credentials, routing configuration, and the application TLS private key remain on the machine running Swobu.
remote client │ │ HTTPS + Share bearer ▼ Swobu Relay │ │ encrypted application TLS ▼ your Swobu │ ├─ route + fallback ├─ protocol translation └─ provider credentials
The Relay makes the public hostname reachable, but it does not terminate application TLS. It forwards the encrypted connection to the Swobu process on your machine. That process terminates application TLS and handles the HTTP request.
The phone in the demo has no Swobu installation and no provider credential. It uses the generated endpoint and Share bearer.
This is TLS passthrough for application traffic, not anonymity.
The Relay reads enough of the client hello to route by SNI. It can see the endpoint identity, public hostname, source-network and connection metadata, timing, traffic volume, and encrypted bytes. It also handles certificate control messages, including the certificate signing request and issued chain.
It does not receive the endpoint private key. It cannot read the Share bearer or application HTTP because both remain inside application TLS until the connection reaches the Owner Swobu process.
Provider credentials stay on the Owner machine and are used for its outbound provider request. They are not sent to the remote client or the Relay.
There is a separate TLS connection between the Owner and the Relay for the reverse transport. The Relay terminates that transport TLS. The application TLS connection shown above remains encrypted inside it and terminates on the Owner.
The Owner machine therefore has to stay online. The Relay provides reachability; the Swobu process on your machine remains the application endpoint.
TLS passthrough is standard. A generic TCP relay can also forward TLS without terminating it. The difference here is what the reachable endpoint represents.
The endpoint keeps routing behind it
The public hostname belongs to the Swobu daemon, not to a provider socket or one route. The Share bearer selects an exact route or workspace grant behind that stable endpoint.
Suppose ChatterUI is configured with:
https://<share>/v1 /v1">
The coding route might currently use a local model. Later I can change the provider, account, region, model, target order, or fallback sequence behind that same route identity. The remote client keeps the same URL, bearer, and route name.
A tunnel solves reachability to one local socket. A Swobu Share keeps a route or workspace boundary behind the reachable URL.
ChatterUI │ │ same HTTPS endpoint ▼ route: coding │ ├─ local model ├─ Azure └─ Bedrock
Changing targets or routing configuration behind coding does not require redistributing remote configuration. Renaming or deleting the route is different: it revokes the grant bound to that route identity. Recreating the old name does not revive its bearer.
There is no compatibility preflight. Swobu sends the real request to the configured target, translating it for that target as part of the attempt. If the attempt fails, routing can continue to the next configured target.
That distinction matters because providers with superficially similar APIs still differ in available models, tool behavior, reasoning state, limits, and protocol details. Swobu translates where the requested semantics are representable. It does not make providers equivalent.
Swobu also does not predict which target has spare quota, is cheapest, or will be fastest. Route order expresses the operator’s policy.
Share one route or a workspace
A Route Share exposes one route:
swobu share dev/coding
Its bearer can use dev/coding , but not sibling routes.
A Workspace Share exposes the workspace’s live client-visible route catalog:
swobu share dev
Its bearer follows route additions, removals, and default changes in dev . It cannot access another workspace, provider credentials, target configuration, workspace mutation, or operator APIs. Workspace and Route Shares are separate grants, even when they refer to the same workspace.
Shares expire after one day by default. The other duration choices are 7d , 30d , and never ; longer-lived sharing remains Preview functionality during the Shared Routes public beta.
Revoke the exact scope you shared:
swobu share revoke dev/coding
DEMO 03 Revoke the Route Share; the next request from the same remote client loses access.
Revoking a Route Share does not revoke a Workspace Share for dev , or vice versa. Revocation denies subsequent authorization with that bearer; it does not promise to cancel provider work that is already in flight. Renaming or deleting a workspace revokes the grants rooted in that workspace identity.
The Owner machine is part of the serving path. If it is asleep, offline, or Swobu is stopped, the remote API is unavailable.
The Relay is still an infrastructure dependency. It cannot read application plaintext, but it can observe connection metadata and interrupt availability. Shared Routes is a public beta, not a GA-scale reliability promise.
The Share bearer is a credential. Anyone who obtains it can use its exact scope until it expires or is revoked. It limits exposure compared with copying a provider credential, but it still needs to be handled as a secret.
Keeping one remote URL does not make Bedrock, Azure, OpenAI-compatible servers, and local models behave identically. The real workload still has to run on each configured path.
For my use case, those constraints are preferable to copying provider credentials onto every machine from which I want to use the setup. The remote client gets an ordinary API endpoint; routing policy and provider credentials stay on the Owner machine.

## Original Extract

Share a stable, bearer-protected LLM endpoint without moving provider credentials. Application TLS still terminates on the machine running Swobu.

Share an LLM endpoint over HTTPS while TLS still terminates on your machine — Swobu Journal SWOBU / JOURNAL What it does Docs Journal GitHub ↗ PRODUCT · HTTPS ROUTES 14 SEPT 2026 · UPDATED 16 SEPT 2026 · Swobu COPY LINK Share an LLM endpoint over HTTPS while TLS still terminates on your machine
Share a stable, bearer-protected LLM endpoint without moving provider credentials. Application TLS still terminates on the machine running Swobu.
I wanted to use the LLM setup on my workstation from a phone and another machine without copying provider credentials onto either of them.
I also did not want the relay that made the machine reachable to become the application TLS endpoint.
Terminal window swobu share dev/coding
That command creates a bearer for the coding route and returns a public HTTPS endpoint. A remote client can use OpenAI-compatible or Anthropic-compatible ingress for the protocols Swobu already supports. It does not need Swobu installed.
The remote machine gets the endpoint and the scoped Share bearer. Provider credentials, routing configuration, and the application TLS private key remain on the machine running Swobu.
remote client │ │ HTTPS + Share bearer ▼ Swobu Relay │ │ encrypted application TLS ▼ your Swobu │ ├─ route + fallback ├─ protocol translation └─ provider credentials
The Relay makes the public hostname reachable, but it does not terminate application TLS. It forwards the encrypted connection to the Swobu process on your machine. That process terminates application TLS and handles the HTTP request.
The phone in the demo has no Swobu installation and no provider credential. It uses the generated endpoint and Share bearer.
This is TLS passthrough for application traffic, not anonymity.
The Relay reads enough of the client hello to route by SNI. It can see the endpoint identity, public hostname, source-network and connection metadata, timing, traffic volume, and encrypted bytes. It also handles certificate control messages, including the certificate signing request and issued chain.
It does not receive the endpoint private key. It cannot read the Share bearer or application HTTP because both remain inside application TLS until the connection reaches the Owner Swobu process.
Provider credentials stay on the Owner machine and are used for its outbound provider request. They are not sent to the remote client or the Relay.
There is a separate TLS connection between the Owner and the Relay for the reverse transport. The Relay terminates that transport TLS. The application TLS connection shown above remains encrypted inside it and terminates on the Owner.
The Owner machine therefore has to stay online. The Relay provides reachability; the Swobu process on your machine remains the application endpoint.
TLS passthrough is standard. A generic TCP relay can also forward TLS without terminating it. The difference here is what the reachable endpoint represents.
The endpoint keeps routing behind it
The public hostname belongs to the Swobu daemon, not to a provider socket or one route. The Share bearer selects an exact route or workspace grant behind that stable endpoint.
Suppose ChatterUI is configured with:
https://<share>/v1 /v1">
The coding route might currently use a local model. Later I can change the provider, account, region, model, target order, or fallback sequence behind that same route identity. The remote client keeps the same URL, bearer, and route name.
A tunnel solves reachability to one local socket. A Swobu Share keeps a route or workspace boundary behind the reachable URL.
ChatterUI │ │ same HTTPS endpoint ▼ route: coding │ ├─ local model ├─ Azure └─ Bedrock
Changing targets or routing configuration behind coding does not require redistributing remote configuration. Renaming or deleting the route is different: it revokes the grant bound to that route identity. Recreating the old name does not revive its bearer.
There is no compatibility preflight. Swobu sends the real request to the configured target, translating it for that target as part of the attempt. If the attempt fails, routing can continue to the next configured target.
That distinction matters because providers with superficially similar APIs still differ in available models, tool behavior, reasoning state, limits, and protocol details. Swobu translates where the requested semantics are representable. It does not make providers equivalent.
Swobu also does not predict which target has spare quota, is cheapest, or will be fastest. Route order expresses the operator’s policy.
Share one route or a workspace
A Route Share exposes one route:
swobu share dev/coding
Its bearer can use dev/coding , but not sibling routes.
A Workspace Share exposes the workspace’s live client-visible route catalog:
swobu share dev
Its bearer follows route additions, removals, and default changes in dev . It cannot access another workspace, provider credentials, target configuration, workspace mutation, or operator APIs. Workspace and Route Shares are separate grants, even when they refer to the same workspace.
Shares expire after one day by default. The other duration choices are 7d , 30d , and never ; longer-lived sharing remains Preview functionality during the Shared Routes public beta.
Revoke the exact scope you shared:
swobu share revoke dev/coding
DEMO 03 Revoke the Route Share; the next request from the same remote client loses access.
Revoking a Route Share does not revoke a Workspace Share for dev , or vice versa. Revocation denies subsequent authorization with that bearer; it does not promise to cancel provider work that is already in flight. Renaming or deleting a workspace revokes the grants rooted in that workspace identity.
The Owner machine is part of the serving path. If it is asleep, offline, or Swobu is stopped, the remote API is unavailable.
The Relay is still an infrastructure dependency. It cannot read application plaintext, but it can observe connection metadata and interrupt availability. Shared Routes is a public beta, not a GA-scale reliability promise.
The Share bearer is a credential. Anyone who obtains it can use its exact scope until it expires or is revoked. It limits exposure compared with copying a provider credential, but it still needs to be handled as a secret.
Keeping one remote URL does not make Bedrock, Azure, OpenAI-compatible servers, and local models behave identically. The real workload still has to run on each configured path.
For my use case, those constraints are preferable to copying provider credentials onto every machine from which I want to use the setup. The remote client gets an ordinary API endpoint; routing policy and provider credentials stay on the Owner machine.
