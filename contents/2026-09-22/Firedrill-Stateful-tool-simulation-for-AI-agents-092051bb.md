---
source: "https://github.com/firedrill-tools/firedrill"
hn_url: "https://news.ycombinator.com/item?id=49801073"
title: "Firedrill: Stateful tool simulation for AI agents"
article_title: "GitHub - firedrill-tools/firedrill: Stateful testing and simulation framework for action-taking AI agents · GitHub"
image: "https://opengraph.githubassets.com/557688cc84ecfa100f93cdf88b1c9f41d154c9b3b24afd18f490d751b9c562c7/firedrill-tools/firedrill"
author: "newton_reload"
captured_at: "2026-09-22T14:23:55Z"
capture_tool: "hn-digest"
hn_id: 49801073
score: 2
comments: 0
posted_at: "2026-09-22T13:35:34Z"
tags:
  - hacker-news
---

# Firedrill: Stateful tool simulation for AI agents

- HN: [49801073](https://news.ycombinator.com/item?id=49801073)
- Source: [github.com](https://github.com/firedrill-tools/firedrill)
- Score: 2
- Comments: 0
- Posted: 2026-09-22T13:35:34Z

## Translation

Title: Firedrill: Stateful tool simulation for AI agents
Article title: GitHub - firedrill-tools/firedrill: Stateful testing and simulation framework for action-taking AI agents · GitHub
Description: Stateful testing and simulation framework for action-taking AI agents - firedrill-tools/firedrill

Article text:
GitHub - firedrill-tools/firedrill: Stateful testing and simulation framework for action-taking AI agents · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
firedrill-tools
/
firedrill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
130 Commits 130 Commits Folders and files
.github .github ci ci docs docs examples/ quickstart examples/ quickstart packages packages python python registry registry release release skills/ firedrill skills/ firedrill templates templates tool-packs tool-packs tooling tooling .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md biome.json biome.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
Firedrill is a simulation and testing framework for AI agents. Define synthetic
tools and data, run your agent against them, and assert on tool calls, state
changes, and events.
Stateful tools with HTTP, MCP, CLI, and function bindings.
Scenario-based tests with faults, response overrides, and virtual time.
Isolated world state, seeded data, snapshots, and resets.
HTML, JSON, and JUnit reports with timelines and optional browser captures.
Repository-defined tools, including independently distributed packages.
Quickstart · Python · TypeScript ·
Documentation · Neutral example ·
Gmail Agent example
Install the current release from PyPI
into your virtual environment with Python 3.10 or later:
python -m pip install " firedrill-run[pytest] "
firedrill --help
Import the local SDK with from firedrill import World, run_drills . The wheel
includes the runtime, CLI, inspector, and report engine. No separate Node.js or
npm installation is required. Choose Tools with firedrill init ; they install
on demand and use the same packages as TypeScript. See the
Python guide for pytest, async agents, mocks, and browser tests.
Requires Node.js 20.19 or later.
Install the current release in your project:
npm install --save-dev @firedrill-run/cli @firedrill-run/sdk
npx firedrill --help
Or run the CLI without adding a project dependency:
npx @firedrill-run/cli init
For a source checkout, use pnpm 9.15–10:
pnpm install --frozen-lockfile
pnpm build
# Use the built CLI in this terminal.
export FIREDRILL_CLI= " $PWD /packages/cli/dist/bin.js "
firedrill () { node " $FIREDRILL_CLI " " $@ " ; }
The examples below use npx firedrill inside a project that has
@firedrill-run/cli installed. Never run bare npx firedrill elsewhere: outside
such a project npm resolves an unrelated package with that name. Python
installations expose firedrill directly, so drop the npx prefix. From source,
use the shell function above or invoke the CLI directly with
node /path/to/firedrill/packages/cli/dist/bin.js .
The programmatic API is @firedrill-run/sdk . To prepare installable archives of the
CLI, SDK, and other packages from this checkout, run
pnpm pack:artifacts -- --output /absolute/path/to/an/empty/directory .
The output includes a package manifest.
Create a project with a synthetic record store:
mkdir firedrill-example
cd firedrill-example
npm init -y
npm install --save-dev @firedrill-run/cli
npx firedrill init --custom records
npx firedrill serve
init creates a Tool declaration, a behavior module, and starting data.
serve starts the backend and opens the inspector.
Open Tools to inspect the implementation or call an operation.
State & activity shows records and calls; Connect agent provides the
connection settings. Tools with a bundled UI also have an Open app action.
Browser actions and API calls use the same state.
The server listens on loopback using available ports. Keep the terminal open;
Ctrl+C stops it. Use --no-open to skip opening the inspector automatically.
Run firedrill init in an existing project for guided setup, or select one of
the published packages in the Tool catalog , for example
npx firedrill init --tool gmail --install or
npx firedrill tool add gmail --install . Both add the catalog's exact
@firedrill-tools/<id> npm release as an ordinary dev dependency and refuse an
archive whose integrity differs from the catalog. The CLI shows the exact
package and version before it asks for installation consent.
Tools can run independently of tests. To check an agent's behavior, add a drill.
For a model-backed project, see the
Gmail Agent example :
an existing Claude Agent SDK assistant runs three drills against a pinned
stateful Gmail Tool, with its synthetic mailbox and reports kept in the project.
A drill defines an agent task, starting conditions, and assertions about the
result. To try one, stop the server, return to the parent directory ( cd .. ),
and create the example test project:
mkdir firedrill-first-drill
cd firedrill-first-drill
npm init -y
npm install --save-dev @firedrill-run/cli
npx firedrill init --path template
npx firedrill validate
npx firedrill plan
npx firedrill run changes-resource
npx firedrill inspect
The template contains a deterministic example agent that writes 7 to a record.
Its drill checks that the write succeeded once and that the final value is 7 .
Replace the example target with your agent when adding your own tests.
To inspect a failing result, edit
firedrill/drills/changes-resource.drill.yaml : change the value-changed
assertion's expected value to 8 , keeping task.input.value at 7 .
Rerun the drill. The report shows expected 8 and actual 7 ; the process exits
with code 1 . Restore the expectation afterwards.
Run the same repository-owned drill command locally and in CI. A pull request is
one useful trigger, not a requirement: pushes, scheduled suites, manual jobs, and
other CI providers use the same command and reports.
firedrill ci describe
firedrill ci init github
The initializer detects the project's package manager or Python setup, preserves
an existing test:firedrill / drills script when present, and writes
.github/workflows/firedrill.yml . It uploads self-contained HTML, JSON, and JUnit
reports even when a drill fails. Use --run for a caller-owned SDK harness and
repeat --trigger to select exact events. See the
continuous-integration guide .
Add a focused safety suite to pull requests without moving the agent into
Firedrill or changing production agent code:
firedrill ci init github \
--run " npm run test:firedrill " \
--trigger pull-request \
--trigger manual
Commit the generated workflow. GitHub runs the repository's existing command at
the exact pull-request revision and shows its exit status as a normal check.
Firedrill retains .firedrill/reports/ as an Actions artifact even when a drill
fails, so a reviewer can open the HTML result, JUnit output, changed state, and
ordered evidence behind the verdict. Requiring that check in branch protection
is optional.
The hosted workflow adds managed worlds, a base-versus-head behavior comparison,
a durable evidence link, and a Firedrill check and pull-request summary. Local CI
remains complete and account-free.
Term
Meaning
Tool
A synthetic dependency with callable operations, input/output schemas, and an implementation
World
Tools, starting data, identities, permissions, and a clock
Scenario
A variation of the starting data, permissions, faults, or scheduled events
Target
Configuration for invoking the agent under test
Drill
A task and its assertions
Run
A recorded execution result, including checks, calls, and state changes
Actors identify who is calling a tool and which operations they may use.
Personas provide descriptions of those identities. See
people and permissions .
Configure the agent's dependencies in test setup:
Bindings use existing configuration or test-side adapters, leaving production
agent logic unchanged. Hardcoded dependencies need an interceptable boundary or
an explicit adapter. See binding recipes
and test-side mocking .
Targets can invoke a module, start a command, call an HTTP endpoint, or use an
external callback supplied by a test harness. External targets run through the
SDK; module, command, and HTTP targets can also run through the CLI.
Model credentials belong to the agent process. For command targets, pass model
credentials and other required host variables through environmentFromHost .
Set target timeouts for the complete model/tool loop.
Use runDrills from an existing test runner:
import { runDrills } from "@firedrill-run/sdk" ;
const result = await runDrills ( {
root : process . cwd ( ) ,
drill : "my-drill" ,
agent : ( { task , binding , signal } ) =>
runMyAgent ( { task , environment : binding . environment , signal } ) ,
} ) ;
expect ( result . verdict ) . toBe ( "passed" ) ;
This example assumes a declared my-drill with an external target.
runMyAgent is your test adapter; it applies the supplied connection values to
your agent. expect comes from your test runner.
runDrills({ setup }) supports per-test data, fault, and Tool overrides.
createLocalWorld() provides direct control over calls, state, time, and resets.
See the SDK reference for lifecycle hooks, concurrency,
capture, and report APIs.
The example template uses the following layout:
your-project/
firedrill.json # source location and selected packages
firedrill/
world.yaml # starting data, identities, access, time
tools/resource-store/
resource-store.tool.yaml # operations and state schemas
behavior.mjs # operation implementations
scenarios/baseline.scenario.yaml # starting conditions
targets/starter-agent.target.yaml # agent invocation
drills/changes-resource.drill.yaml # task and assertions
suites/resource-store-conformance.suite.yaml
firedrill-example/agent.mjs # example agent
.firedrill-tools/ # vendored Tool dependencies
.firedrill/ # generated state, builds, and reports
Definitions support JSON or YAML; use either consistently or mix them.
Resource suffixes identify file types, such as .tool.json and .drill.yaml .
References use IDs inside the files, so you can organize folders as needed.
Tool implementations are JavaScript or TypeScript.
Commit definitions, behavior modules, test code, package manifests, lockfiles,
and referenced .firedrill-tools/ archives. Generated files under .firedrill/
are ignored by init . Keep credentials, .env files, and node_modules/ ignored
as well.
Tool state lives in SQLite. Runtime writes leave the source definitions unchanged.
Source edits apply to the next build; restart serve to load them.
Each trial or retry uses an isolated world. A full reset restores its baseline;
a scoped reset restores selected Tools. Reset affects

[truncated]

## Original Extract

Stateful testing and simulation framework for action-taking AI agents - firedrill-tools/firedrill

GitHub - firedrill-tools/firedrill: Stateful testing and simulation framework for action-taking AI agents · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
firedrill-tools
/
firedrill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
130 Commits 130 Commits Folders and files
.github .github ci ci docs docs examples/ quickstart examples/ quickstart packages packages python python registry registry release release skills/ firedrill skills/ firedrill templates templates tool-packs tool-packs tooling tooling .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md biome.json biome.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
Firedrill is a simulation and testing framework for AI agents. Define synthetic
tools and data, run your agent against them, and assert on tool calls, state
changes, and events.
Stateful tools with HTTP, MCP, CLI, and function bindings.
Scenario-based tests with faults, response overrides, and virtual time.
Isolated world state, seeded data, snapshots, and resets.
HTML, JSON, and JUnit reports with timelines and optional browser captures.
Repository-defined tools, including independently distributed packages.
Quickstart · Python · TypeScript ·
Documentation · Neutral example ·
Gmail Agent example
Install the current release from PyPI
into your virtual environment with Python 3.10 or later:
python -m pip install " firedrill-run[pytest] "
firedrill --help
Import the local SDK with from firedrill import World, run_drills . The wheel
includes the runtime, CLI, inspector, and report engine. No separate Node.js or
npm installation is required. Choose Tools with firedrill init ; they install
on demand and use the same packages as TypeScript. See the
Python guide for pytest, async agents, mocks, and browser tests.
Requires Node.js 20.19 or later.
Install the current release in your project:
npm install --save-dev @firedrill-run/cli @firedrill-run/sdk
npx firedrill --help
Or run the CLI without adding a project dependency:
npx @firedrill-run/cli init
For a source checkout, use pnpm 9.15–10:
pnpm install --frozen-lockfile
pnpm build
# Use the built CLI in this terminal.
export FIREDRILL_CLI= " $PWD /packages/cli/dist/bin.js "
firedrill () { node " $FIREDRILL_CLI " " $@ " ; }
The examples below use npx firedrill inside a project that has
@firedrill-run/cli installed. Never run bare npx firedrill elsewhere: outside
such a project npm resolves an unrelated package with that name. Python
installations expose firedrill directly, so drop the npx prefix. From source,
use the shell function above or invoke the CLI directly with
node /path/to/firedrill/packages/cli/dist/bin.js .
The programmatic API is @firedrill-run/sdk . To prepare installable archives of the
CLI, SDK, and other packages from this checkout, run
pnpm pack:artifacts -- --output /absolute/path/to/an/empty/directory .
The output includes a package manifest.
Create a project with a synthetic record store:
mkdir firedrill-example
cd firedrill-example
npm init -y
npm install --save-dev @firedrill-run/cli
npx firedrill init --custom records
npx firedrill serve
init creates a Tool declaration, a behavior module, and starting data.
serve starts the backend and opens the inspector.
Open Tools to inspect the implementation or call an operation.
State & activity shows records and calls; Connect agent provides the
connection settings. Tools with a bundled UI also have an Open app action.
Browser actions and API calls use the same state.
The server listens on loopback using available ports. Keep the terminal open;
Ctrl+C stops it. Use --no-open to skip opening the inspector automatically.
Run firedrill init in an existing project for guided setup, or select one of
the published packages in the Tool catalog , for example
npx firedrill init --tool gmail --install or
npx firedrill tool add gmail --install . Both add the catalog's exact
@firedrill-tools/<id> npm release as an ordinary dev dependency and refuse an
archive whose integrity differs from the catalog. The CLI shows the exact
package and version before it asks for installation consent.
Tools can run independently of tests. To check an agent's behavior, add a drill.
For a model-backed project, see the
Gmail Agent example :
an existing Claude Agent SDK assistant runs three drills against a pinned
stateful Gmail Tool, with its synthetic mailbox and reports kept in the project.
A drill defines an agent task, starting conditions, and assertions about the
result. To try one, stop the server, return to the parent directory ( cd .. ),
and create the example test project:
mkdir firedrill-first-drill
cd firedrill-first-drill
npm init -y
npm install --save-dev @firedrill-run/cli
npx firedrill init --path template
npx firedrill validate
npx firedrill plan
npx firedrill run changes-resource
npx firedrill inspect
The template contains a deterministic example agent that writes 7 to a record.
Its drill checks that the write succeeded once and that the final value is 7 .
Replace the example target with your agent when adding your own tests.
To inspect a failing result, edit
firedrill/drills/changes-resource.drill.yaml : change the value-changed
assertion's expected value to 8 , keeping task.input.value at 7 .
Rerun the drill. The report shows expected 8 and actual 7 ; the process exits
with code 1 . Restore the expectation afterwards.
Run the same repository-owned drill command locally and in CI. A pull request is
one useful trigger, not a requirement: pushes, scheduled suites, manual jobs, and
other CI providers use the same command and reports.
firedrill ci describe
firedrill ci init github
The initializer detects the project's package manager or Python setup, preserves
an existing test:firedrill / drills script when present, and writes
.github/workflows/firedrill.yml . It uploads self-contained HTML, JSON, and JUnit
reports even when a drill fails. Use --run for a caller-owned SDK harness and
repeat --trigger to select exact events. See the
continuous-integration guide .
Add a focused safety suite to pull requests without moving the agent into
Firedrill or changing production agent code:
firedrill ci init github \
--run " npm run test:firedrill " \
--trigger pull-request \
--trigger manual
Commit the generated workflow. GitHub runs the repository's existing command at
the exact pull-request revision and shows its exit status as a normal check.
Firedrill retains .firedrill/reports/ as an Actions artifact even when a drill
fails, so a reviewer can open the HTML result, JUnit output, changed state, and
ordered evidence behind the verdict. Requiring that check in branch protection
is optional.
The hosted workflow adds managed worlds, a base-versus-head behavior comparison,
a durable evidence link, and a Firedrill check and pull-request summary. Local CI
remains complete and account-free.
Term
Meaning
Tool
A synthetic dependency with callable operations, input/output schemas, and an implementation
World
Tools, starting data, identities, permissions, and a clock
Scenario
A variation of the starting data, permissions, faults, or scheduled events
Target
Configuration for invoking the agent under test
Drill
A task and its assertions
Run
A recorded execution result, including checks, calls, and state changes
Actors identify who is calling a tool and which operations they may use.
Personas provide descriptions of those identities. See
people and permissions .
Configure the agent's dependencies in test setup:
Bindings use existing configuration or test-side adapters, leaving production
agent logic unchanged. Hardcoded dependencies need an interceptable boundary or
an explicit adapter. See binding recipes
and test-side mocking .
Targets can invoke a module, start a command, call an HTTP endpoint, or use an
external callback supplied by a test harness. External targets run through the
SDK; module, command, and HTTP targets can also run through the CLI.
Model credentials belong to the agent process. For command targets, pass model
credentials and other required host variables through environmentFromHost .
Set target timeouts for the complete model/tool loop.
Use runDrills from an existing test runner:
import { runDrills } from "@firedrill-run/sdk" ;
const result = await runDrills ( {
root : process . cwd ( ) ,
drill : "my-drill" ,
agent : ( { task , binding , signal } ) =>
runMyAgent ( { task , environment : binding . environment , signal } ) ,
} ) ;
expect ( result . verdict ) . toBe ( "passed" ) ;
This example assumes a declared my-drill with an external target.
runMyAgent is your test adapter; it applies the supplied connection values to
your agent. expect comes from your test runner.
runDrills({ setup }) supports per-test data, fault, and Tool overrides.
createLocalWorld() provides direct control over calls, state, time, and resets.
See the SDK reference for lifecycle hooks, concurrency,
capture, and report APIs.
The example template uses the following layout:
your-project/
firedrill.json # source location and selected packages
firedrill/
world.yaml # starting data, identities, access, time
tools/resource-store/
resource-store.tool.yaml # operations and state schemas
behavior.mjs # operation implementations
scenarios/baseline.scenario.yaml # starting conditions
targets/starter-agent.target.yaml # agent invocation
drills/changes-resource.drill.yaml # task and assertions
suites/resource-store-conformance.suite.yaml
firedrill-example/agent.mjs # example agent
.firedrill-tools/ # vendored Tool dependencies
.firedrill/ # generated state, builds, and reports
Definitions support JSON or YAML; use either consistently or mix them.
Resource suffixes identify file types, such as .tool.json and .drill.yaml .
References use IDs inside the files, so you can organize folders as needed.
Tool implementations are JavaScript or TypeScript.
Commit definitions, behavior modules, test code, package manifests, lockfiles,
and referenced .firedrill-tools/ archives. Generated files under .firedrill/
are ignored by init . Keep credentials, .env files, and node_modules/ ignored
as well.
Tool state lives in SQLite. Runtime writes leave the source definitions unchanged.
Source edits apply to the next build; restart serve to load them.
Each trial or retry uses an isolated world. A full reset restores its baseline;
a scoped reset restores selected Tools. Reset affects

[truncated]
