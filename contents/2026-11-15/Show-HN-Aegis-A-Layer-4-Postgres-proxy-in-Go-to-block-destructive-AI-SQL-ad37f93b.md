---
source: "https://github.com/raymondhani/aegis-proxy"
hn_url: "https://news.ycombinator.com/item?id=50509721"
title: "Show HN: Aegis – A Layer-4 Postgres proxy in Go to block destructive AI SQL"
article_title: "GitHub - raymondhani/aegis-proxy: An open-source, wire-transparent security proxy for PostgreSQL that intercepts and blocks destructive SQL queries from AI agents before they hit your database. · GitHub"
image: "https://opengraph.githubassets.com/b6f82a918b75ab4f60c03224750b6af3c47cad9ec5096612f7a55fac9b0d939e/raymondhani/aegis-proxy"
author: "raymondartin"
captured_at: "2026-11-15T19:34:34Z"
capture_tool: "hn-digest"
hn_id: 50509721
score: 1
comments: 0
posted_at: "2026-11-15T19:11:40Z"
tags:
  - hacker-news
---

# Show HN: Aegis – A Layer-4 Postgres proxy in Go to block destructive AI SQL

- HN: [50509721](https://news.ycombinator.com/item?id=50509721)
- Source: [github.com](https://github.com/raymondhani/aegis-proxy)
- Score: 1
- Comments: 0
- Posted: 2026-11-15T19:11:40Z

## Translation

Title: Show HN: Aegis – A Layer-4 Postgres proxy in Go to block destructive AI SQL
Article title: GitHub - raymondhani/aegis-proxy: An open-source, wire-transparent security proxy for PostgreSQL that intercepts and blocks destructive SQL queries from AI agents before they hit your database. · GitHub
Description: An open-source, wire-transparent security proxy for PostgreSQL that intercepts and blocks destructive SQL queries from AI agents before they hit your database. - raymondhani/aegis-proxy

Article text:
GitHub - raymondhani/aegis-proxy: An open-source, wire-transparent security proxy for PostgreSQL that intercepts and blocks destructive SQL queries from AI agents before they hit your database. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
raymondhani
aegis-proxy Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
An open-source, wire-transparent security proxy for PostgreSQL that intercepts and blocks destructive SQL queries from AI agents before they hit your database.
github.com/raymondhani/aegis-proxy Topics
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
81 Commits 81 Commits Folders and files
.github/ workflows .github/ workflows proxy-engine proxy-engine python-sdk python-sdk scripts scripts .env.sample .env.sample .gitignore .gitignore AGENTS.md AGENTS.md Dockerfile.oss Dockerfile.oss LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml go.work go.work go.work.sum go.work.sum View all files Repository files navigation
Aegis sits between your application and your Postgres database and blocks destructive queries before they run, without you changing a line of your database code.
👋 Note: If you are cloning this to test the Wasm AST parser locally, please drop a ⭐ on the repo! It helps me know people are actually using the open-source tier.
sequenceDiagram
participant App as Your Application<br/>(psycopg2, or any client)
participant Proxy as Aegis Proxy<br/>:5433 TCP intercept / :5434 admin+metrics
participant DB as Your Database<br/>(e.g. Neon Postgres)
App->>Proxy: SELECT ... (benign query)
Proxy->>DB: forwarded over TLS
DB-->>Proxy: rows
Proxy-->>App: rows
App->>Proxy: DROP TABLE ... (destructive query)
Note over Proxy: parsed, blocked, connection terminated<br/>(the "Guillotine")
Proxy-->>App: fatal error
Note over DB: never reached
Loading
Your application never talks to the database directly — every query passes through the proxy first. A benign query is forwarded and the rows come straight back. A destructive one ( DROP TABLE , an unfiltered DELETE , etc.) is intercepted and the connection is terminated before it reaches your database.
Your Application Aegis Proxy Your Database
(psycopg2, or any client) (e.g. Neon Postgres)
+--------------------+ Postgres wire +-----------------------------+ TLS +--------------------+
| | ------------------> | :5433 TCP intercept | -------> | |
| SELECT ... | | - parses every query | | real database |
| DROP TABLE ... | <------------------ | - enforces rate limits | <------- | |
| | rows, or a | - blocks destructive | | |
+--------------------+ fatal error | statements (the | +--------------------+
if blocked | "Guillotine") |
| :5434 admin + metrics API |
+-----------------------------+
Quickstart
Three steps, one command to start the stack, and a running proxy.
Step 1 — Get Aegis and your credentials
The primary way to install Aegis is to pull the published image straight from Docker Hub:
docker pull raymondartin2/aegis-proxy
Expected output (tag/digest will vary by release):
latest: Pulling from raymondartin2/aegis-proxy
Digest: sha256:...
Status: Downloaded newer image for raymondartin2/aegis-proxy:latest
docker.io/raymondartin2/aegis-proxy:latest
You still need this repo's docker-compose.yml and .env.sample to run the stack, so grab those too:
git clone https://github.com/raymondhani/aegis-proxy.git
cd aegis-proxy
cp .env.sample .env
Building the image yourself instead of pulling it is also supported — see Dockerfile in this repo — but pulling the published image is the recommended, primary path.
Now open .env and replace the four placeholders with your own values:
None of these are optional: the proxy exits immediately at startup if AEGIS_JWT_SECRET is missing, and every step below assumes all four are set.
docker compose up -d
Expected output (container name may vary slightly depending on your folder name):
Container aegis-proxy Started
This single command starts the container Aegis needs: the proxy itself, listening on 5433 (the database connection) and 5434 (an admin/metrics API). The OSS proxy runs on in-memory state and has no hard dependency on Redis — Redis is used only by the Enterprise tier for distributed/HA state.
curl -s http://localhost:5434/metrics
Expected output:
{ "queries_processed" : 0 , "queries_blocked" : 0 , "active_connections" : 0 , "sessions_jailed" : 0 , "anomalies_detected" : 0 }
If you see that, the proxy is up and ready to accept connections.
Try it: block a destructive query
Install the SDK and the Postgres driver, then run this against the stack you just started — nothing here needs modification beyond the credentials you already set in .env .
pip install aegis-proxy-sdk psycopg2-binary
import psycopg2
from aegis_sdk import safe_db_run
@ safe_db_run ()
def run_demo ( db_url : str ):
conn = psycopg2 . connect ( db_url )
conn . autocommit = True
cursor = conn . cursor ()
print ( "Running an ordinary query..." )
cursor . execute ( "SELECT 1;" )
print ( f"Result: { cursor . fetchone () } " )
print ( "Running a destructive query: DROP TABLE users;" )
cursor . execute ( "DROP TABLE users;" )
if __name__ == "__main__" :
try :
run_demo ()
except psycopg2 . OperationalError as e :
print ( f" \n [Aegis] Connection severed by the proxy: { e } " )
The @safe_db_run() decorator does the setup for you: it provisions a throwaway Neon database branch, mints a session token, and points db_url at the local proxy instead of the real database — then deletes the branch again when the function returns. Your function just receives a working db_url and uses it like any other Postgres connection string.
Running an ordinary query...
Result: (1,)
Running a destructive query: DROP TABLE users;
[Aegis] Connection severed by the proxy: server closed the connection unexpectedly
This probably means the server terminated abnormally
before or while processing the request.
The SELECT 1 succeeds normally. The DROP TABLE never reaches the database — the proxy detects it, terminates the connection, and psycopg2 surfaces that as OperationalError . Catching that exception around the specific call you expect to be blocked is the intended pattern, not a workaround.
@safe_db_run also accepts a few optional keyword arguments for less common setups:
Drop-in Guardrails for LangChain & AI Agents
Because Aegis operates completely transparently at Layer-4 (the network layer), developers do not need to write custom application code, install framework-specific plugins, or change their AI logic. To secure an AI agent against hallucinations, simply spin up Aegis and change your standard database connection port from 5432 to 5433.
# The Old, Dangerous Way (Direct to DB):
db = SQLDatabase . from_uri ( "postgresql://user:pass@localhost:5432/my_db" )
# The Secure, Aegis Way (Routed through the Layer-4 Proxy):
db = SQLDatabase . from_uri ( "postgresql://user:pass@localhost:5433/my_db" )
What each tier gives you
Aegis ships as three products built on the same core. Every capability below is enforced by the proxy itself, not just documented — the automated benchmark suite in this workspace asserts each one directly against a running instance of every tier.
Symptom
Cause
Fix
Wrong port — your app hangs or times out connecting on 5432
Aegis is not a Postgres server on the standard port. It listens on 5433 specifically, so it can run alongside a real Postgres instance on the same host without colliding.
Point your connection string at 5433 — the docker-compose.yml in this repo already maps it.
Missing credentials — the container exits right after starting ( docker compose ps shows Exited ), or your script raises ValueError: AEGIS_JWT_SECRET environment variable is missing.
.env was never created, or one of the four values is still the placeholder from .env.sample .
cp .env.sample .env , fill in real values ( AEGIS_JWT_SECRET needs to be 32+ characters), then docker compose up -d again.
Container not started — docker ps shows nothing for aegis-proxy
docker compose up -d was never run in this directory, or it failed earlier without you noticing.
Run docker compose up -d from the repo root and read its output — each container should say Started .
Connection refused — docker compose ps shows aegis-proxy as Up , but your client still can't connect
Almost always a host mismatch. A client running inside its own Docker container can't reach the proxy via localhost — it needs the Docker-internal address.
From another container, connect to host.docker.internal:5433 (Mac/Windows) or join the aegis-net network and use the service name aegis-proxy . Running your script directly on the host, as in the quickstart above, needs no special host.
Unexpectedly severed query — psycopg2.OperationalError: server closed the connection unexpectedly right after running a query
This is almost always the Guillotine working as intended — the proxy detected a destructive statement and terminated the connection to block it. It is not a network problem.
Catch psycopg2.OperationalError around the specific call, exactly as the example above does. If a query you didn't expect to be destructive gets blocked, that's a policy question, not a connectivity bug
[truncated]
Aegis is designed to support a Multi-Agent Workflow where several autonomous AI agents can operate simultaneously without cross-contamination.
Sandboxed Execution: Every agent interaction is securely sandboxed in a dedicated Copy-on-Write (CoW) ephemeral branch provisioned via the Neon API.
Z-Score Anomaly Detection: Agents' queries are mathematically analyzed in real-time using native PostgreSQL AST parsing via wasilibs/go-pgquery .
Cryptographic Agent Identity: Agents are identified using signed JWTs for secure, tamper-proof authentication.
Advanced Threat Telemetry: Zero-latency monitoring is powered by OpenTelemetry.
🌉 Go/Python Bridging Mechanism
Aegis seamlessly bridges high-performance systems-level intercept logic with accessible data-science workflows.
The Go Proxy Engine (Layer 4):
Written in Go using Clean Architecture. It binds to local ports ( 5433 for TCP, 5434 for the HTTP admin/metrics API) and intercepts PostgreSQL wire-protocol traffic.
The Python SDK (Application Layer):
Using the @safe_db_run decorator, the SDK automatically provisions an ephemeral branch, configures the agent's connection string to point to the local Go proxy instead of directly to the database, and manages the lifecycle of the run.
Aegis uses a hardened CI/CD pipeline for distribution:
Continuous Integration: Every commit to main and every pull request triggers a suite of unit tests.
Docker Publishing: The pipeline builds multi-architecture Docker images (ARM64 & AMD64) and pushes them to Docker Hub.
While @sa

[truncated]

## Original Extract

An open-source, wire-transparent security proxy for PostgreSQL that intercepts and blocks destructive SQL queries from AI agents before they hit your database. - raymondhani/aegis-proxy

GitHub - raymondhani/aegis-proxy: An open-source, wire-transparent security proxy for PostgreSQL that intercepts and blocks destructive SQL queries from AI agents before they hit your database. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
raymondhani
aegis-proxy Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
An open-source, wire-transparent security proxy for PostgreSQL that intercepts and blocks destructive SQL queries from AI agents before they hit your database.
github.com/raymondhani/aegis-proxy Topics
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
81 Commits 81 Commits Folders and files
.github/ workflows .github/ workflows proxy-engine proxy-engine python-sdk python-sdk scripts scripts .env.sample .env.sample .gitignore .gitignore AGENTS.md AGENTS.md Dockerfile.oss Dockerfile.oss LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml go.work go.work go.work.sum go.work.sum View all files Repository files navigation
Aegis sits between your application and your Postgres database and blocks destructive queries before they run, without you changing a line of your database code.
👋 Note: If you are cloning this to test the Wasm AST parser locally, please drop a ⭐ on the repo! It helps me know people are actually using the open-source tier.
sequenceDiagram
participant App as Your Application<br/>(psycopg2, or any client)
participant Proxy as Aegis Proxy<br/>:5433 TCP intercept / :5434 admin+metrics
participant DB as Your Database<br/>(e.g. Neon Postgres)
App->>Proxy: SELECT ... (benign query)
Proxy->>DB: forwarded over TLS
DB-->>Proxy: rows
Proxy-->>App: rows
App->>Proxy: DROP TABLE ... (destructive query)
Note over Proxy: parsed, blocked, connection terminated<br/>(the "Guillotine")
Proxy-->>App: fatal error
Note over DB: never reached
Loading
Your application never talks to the database directly — every query passes through the proxy first. A benign query is forwarded and the rows come straight back. A destructive one ( DROP TABLE , an unfiltered DELETE , etc.) is intercepted and the connection is terminated before it reaches your database.
Your Application Aegis Proxy Your Database
(psycopg2, or any client) (e.g. Neon Postgres)
+--------------------+ Postgres wire +-----------------------------+ TLS +--------------------+
| | ------------------> | :5433 TCP intercept | -------> | |
| SELECT ... | | - parses every query | | real database |
| DROP TABLE ... | <------------------ | - enforces rate limits | <------- | |
| | rows, or a | - blocks destructive | | |
+--------------------+ fatal error | statements (the | +--------------------+
if blocked | "Guillotine") |
| :5434 admin + metrics API |
+-----------------------------+
Quickstart
Three steps, one command to start the stack, and a running proxy.
Step 1 — Get Aegis and your credentials
The primary way to install Aegis is to pull the published image straight from Docker Hub:
docker pull raymondartin2/aegis-proxy
Expected output (tag/digest will vary by release):
latest: Pulling from raymondartin2/aegis-proxy
Digest: sha256:...
Status: Downloaded newer image for raymondartin2/aegis-proxy:latest
docker.io/raymondartin2/aegis-proxy:latest
You still need this repo's docker-compose.yml and .env.sample to run the stack, so grab those too:
git clone https://github.com/raymondhani/aegis-proxy.git
cd aegis-proxy
cp .env.sample .env
Building the image yourself instead of pulling it is also supported — see Dockerfile in this repo — but pulling the published image is the recommended, primary path.
Now open .env and replace the four placeholders with your own values:
None of these are optional: the proxy exits immediately at startup if AEGIS_JWT_SECRET is missing, and every step below assumes all four are set.
docker compose up -d
Expected output (container name may vary slightly depending on your folder name):
Container aegis-proxy Started
This single command starts the container Aegis needs: the proxy itself, listening on 5433 (the database connection) and 5434 (an admin/metrics API). The OSS proxy runs on in-memory state and has no hard dependency on Redis — Redis is used only by the Enterprise tier for distributed/HA state.
curl -s http://localhost:5434/metrics
Expected output:
{ "queries_processed" : 0 , "queries_blocked" : 0 , "active_connections" : 0 , "sessions_jailed" : 0 , "anomalies_detected" : 0 }
If you see that, the proxy is up and ready to accept connections.
Try it: block a destructive query
Install the SDK and the Postgres driver, then run this against the stack you just started — nothing here needs modification beyond the credentials you already set in .env .
pip install aegis-proxy-sdk psycopg2-binary
import psycopg2
from aegis_sdk import safe_db_run
@ safe_db_run ()
def run_demo ( db_url : str ):
conn = psycopg2 . connect ( db_url )
conn . autocommit = True
cursor = conn . cursor ()
print ( "Running an ordinary query..." )
cursor . execute ( "SELECT 1;" )
print ( f"Result: { cursor . fetchone () } " )
print ( "Running a destructive query: DROP TABLE users;" )
cursor . execute ( "DROP TABLE users;" )
if __name__ == "__main__" :
try :
run_demo ()
except psycopg2 . OperationalError as e :
print ( f" \n [Aegis] Connection severed by the proxy: { e } " )
The @safe_db_run() decorator does the setup for you: it provisions a throwaway Neon database branch, mints a session token, and points db_url at the local proxy instead of the real database — then deletes the branch again when the function returns. Your function just receives a working db_url and uses it like any other Postgres connection string.
Running an ordinary query...
Result: (1,)
Running a destructive query: DROP TABLE users;
[Aegis] Connection severed by the proxy: server closed the connection unexpectedly
This probably means the server terminated abnormally
before or while processing the request.
The SELECT 1 succeeds normally. The DROP TABLE never reaches the database — the proxy detects it, terminates the connection, and psycopg2 surfaces that as OperationalError . Catching that exception around the specific call you expect to be blocked is the intended pattern, not a workaround.
@safe_db_run also accepts a few optional keyword arguments for less common setups:
Drop-in Guardrails for LangChain & AI Agents
Because Aegis operates completely transparently at Layer-4 (the network layer), developers do not need to write custom application code, install framework-specific plugins, or change their AI logic. To secure an AI agent against hallucinations, simply spin up Aegis and change your standard database connection port from 5432 to 5433.
# The Old, Dangerous Way (Direct to DB):
db = SQLDatabase . from_uri ( "postgresql://user:pass@localhost:5432/my_db" )
# The Secure, Aegis Way (Routed through the Layer-4 Proxy):
db = SQLDatabase . from_uri ( "postgresql://user:pass@localhost:5433/my_db" )
What each tier gives you
Aegis ships as three products built on the same core. Every capability below is enforced by the proxy itself, not just documented — the automated benchmark suite in this workspace asserts each one directly against a running instance of every tier.
Symptom
Cause
Fix
Wrong port — your app hangs or times out connecting on 5432
Aegis is not a Postgres server on the standard port. It listens on 5433 specifically, so it can run alongside a real Postgres instance on the same host without colliding.
Point your connection string at 5433 — the docker-compose.yml in this repo already maps it.
Missing credentials — the container exits right after starting ( docker compose ps shows Exited ), or your script raises ValueError: AEGIS_JWT_SECRET environment variable is missing.
.env was never created, or one of the four values is still the placeholder from .env.sample .
cp .env.sample .env , fill in real values ( AEGIS_JWT_SECRET needs to be 32+ characters), then docker compose up -d again.
Container not started — docker ps shows nothing for aegis-proxy
docker compose up -d was never run in this directory, or it failed earlier without you noticing.
Run docker compose up -d from the repo root and read its output — each container should say Started .
Connection refused — docker compose ps shows aegis-proxy as Up , but your client still can't connect
Almost always a host mismatch. A client running inside its own Docker container can't reach the proxy via localhost — it needs the Docker-internal address.
From another container, connect to host.docker.internal:5433 (Mac/Windows) or join the aegis-net network and use the service name aegis-proxy . Running your script directly on the host, as in the quickstart above, needs no special host.
Unexpectedly severed query — psycopg2.OperationalError: server closed the connection unexpectedly right after running a query
This is almost always the Guillotine working as intended — the proxy detected a destructive statement and terminated the connection to block it. It is not a network problem.
Catch psycopg2.OperationalError around the specific call, exactly as the example above does. If a query you didn't expect to be destructive gets blocked, that's a policy question, not a connectivity bug
[truncated]
Aegis is designed to support a Multi-Agent Workflow where several autonomous AI agents can operate simultaneously without cross-contamination.
Sandboxed Execution: Every agent interaction is securely sandboxed in a dedicated Copy-on-Write (CoW) ephemeral branch provisioned via the Neon API.
Z-Score Anomaly Detection: Agents' queries are mathematically analyzed in real-time using native PostgreSQL AST parsing via wasilibs/go-pgquery .
Cryptographic Agent Identity: Agents are identified using signed JWTs for secure, tamper-proof authentication.
Advanced Threat Telemetry: Zero-latency monitoring is powered by OpenTelemetry.
🌉 Go/Python Bridging Mechanism
Aegis seamlessly bridges high-performance systems-level intercept logic with accessible data-science workflows.
The Go Proxy Engine (Layer 4):
Written in Go using Clean Architecture. It binds to local ports ( 5433 for TCP, 5434 for the HTTP admin/metrics API) and intercepts PostgreSQL wire-protocol traffic.
The Python SDK (Application Layer):
Using the @safe_db_run decorator, the SDK automatically provisions an ephemeral branch, configures the agent's connection string to point to the local Go proxy instead of directly to the database, and manages the lifecycle of the run.
Aegis uses a hardened CI/CD pipeline for distribution:
Continuous Integration: Every commit to main and every pull request triggers a suite of unit tests.
Docker Publishing: The pipeline builds multi-architecture Docker images (ARM64 & AMD64) and pushes them to Docker Hub.
While @sa

[truncated]
