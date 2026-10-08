---
source: "https://github.com/p10node/k10s"
hn_url: "https://news.ycombinator.com/item?id=50009904"
title: "Show HN: K10s – A Clickable Kubernetes TUI (Go, Bubble Tea)"
article_title: "GitHub - p10node/k10s: Kubernetes TUI you can click. Instant search, logs/exec/port-forward, 7 themes, context-aware AI. Single Go binary. · GitHub"
image: "https://opengraph.githubassets.com/e380b8f53a4c2a58b0c5ca17e2dc54859ad4da67485788ce915dedd431ed935f/p10node/k10s"
author: "pierreneter"
captured_at: "2026-10-08T18:30:24Z"
capture_tool: "hn-digest"
hn_id: 50009904
score: 1
comments: 0
posted_at: "2026-10-08T18:29:06Z"
tags:
  - hacker-news
---

# Show HN: K10s – A Clickable Kubernetes TUI (Go, Bubble Tea)

- HN: [50009904](https://news.ycombinator.com/item?id=50009904)
- Source: [github.com](https://github.com/p10node/k10s)
- Score: 1
- Comments: 0
- Posted: 2026-10-08T18:29:06Z

## Translation

Title: Show HN: K10s – A Clickable Kubernetes TUI (Go, Bubble Tea)
Article title: GitHub - p10node/k10s: Kubernetes TUI you can click. Instant search, logs/exec/port-forward, 7 themes, context-aware AI. Single Go binary. · GitHub
Description: Kubernetes TUI you can click. Instant search, logs/exec/port-forward, 7 themes, context-aware AI. Single Go binary. - p10node/k10s

Article text:
GitHub - p10node/k10s: Kubernetes TUI you can click. Instant search, logs/exec/port-forward, 7 themes, context-aware AI. Single Go binary. · GitHub
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
p10node
k10s Public Notifications You must be signed in to change notification settings
Star 101 ( 101 ) You must be signed in to star a repository
Kubernetes TUI you can click. Instant search, logs/exec/port-forward, 7 themes, context-aware AI. Single Go binary.
Readme Apache-2.0 license Activity Custom properties Stars
7 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
28 Commits 28 Commits Folders and files
.github/ workflows .github/ workflows assets assets cmd/ shot cmd/ shot docs docs examples examples internal internal .gitignore .gitignore Justfile Justfile LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh main.go main.go main_test.go main_test.go uninstall.sh uninstall.sh View all files Repository files navigation
The Kubernetes terminal UI you can click .
k9s taught us to live in the terminal. k10s makes that terminal
point-and-click, instantly searchable, themeable - and gives it an AI that
already knows your cluster, namespace and selected object.
Install · Try it with no cluster ·
Features · k9s → k10s · Docs ·
Build log
A real terminal capture of k10s demo running - just screenshot , on the offline demo backend. Every frame in these docs comes out of the running binary, never hand-drawn.
A cluster dashboard is something you open twenty times a day, usually while
something is on fire. The two things that matter are how fast it opens
and how little you have to remember .
Nothing is hidden behind memorised keys. The actions that apply to the
thing you selected are listed, right there, in their own pane. Click one,
or press the letter next to it.
Your mouse works. Click a row, click a pane, click [ zoom ] , click
ns default ▾ , scroll the table. Every border button is real.
One search box for the whole cluster. ctrl+p searches resource kinds
and objects together. No prefix language to learn.
It opens instantly. Startup registers zero watches and waits for
nothing - informers start lazily, per kind, the first time you look at
one. Opening Pods watches pods, not every Secret and Event you own.
( how, and the regression guards )
AI that can see the screen. ctrl+a , ask in plain English. The current
context, namespace, kind and selected object are injected into the prompt,
so "why is this pod unhealthy?" means this pod.
It updates itself. /update installs the newest release over the
running binary - checksum-verified, atomic, offers to restart into it.
Your k9s-style command plugins fit. Put scoped shortcuts in
~/.k10s/plugins.yaml ; they appear beside built-in actions and receive the
selected object, namespace, context and column values.
Try it in 30 seconds (no cluster required)
k10s ships an offline demo backend: a realistic cluster to click around in,
including a CrashLoopBackOff to poke at. It is opt-in , because sample
data should never be mistaken for your machine.
git clone https://github.com/p10node/k10s && cd k10s
go run . demo # the sample cluster - fake data, clearly labelled
go run . # your real cluster, or "No cluster" if there isn't one
The demo is a context , not a mode. k10s demo opens on it, /demo
switches to it from anywhere, and :ctx always lists it (labelled
k10s demo · sample data , with a legend under the list). To leave it,
pick any other context - there is no separate exit. While it is up the
header carries a DEMO marker, so no frame of it can be mistaken for a real
cluster.
Plain k10s reads the same kubeconfig kubectl does ( $KUBECONFIG , else
~/.kube/config ) and shows what that context can reach - and nothing else.
With no kubeconfig, or a context whose API server does not answer, the main
panel says No cluster and points at the way in: r retries, :ctx picks
another context, /setup has the kubectl and kubeconfig links.
docs/cluster-setup.md is the same guide, longer.
curl -fsSL https://p10node.com/k10s/install.sh | sh
macOS and Linux, amd64 and arm64 . It picks the right prebuilt binary,
verifies its sha256 against the release manifest, and installs it into
/usr/local/bin (or ~/.local/bin when that needs a password it cannot
ask for). --dir , --version and --no-sudo are in docs/install.md ,
along with how to read it before you run it and the matching
uninstall.sh .
go install github.com/p10node/k10s@latest
Or from a clone, which also stamps the version into the binary:
git clone https://github.com/p10node/k10s && cd k10s
just install # → $GOBIN/k10s, version-stamped
Prebuilt static binaries for darwin/amd64, darwin/arm64, linux/amd64,
linux/arm64 and windows/amd64 are published on every tag — Windows is the
one platform the installer script sends to the
release page instead. Every later
upgrade is just:
k10s # your current kubeconfig context
k10s demo # the built-in sample cluster, no cluster needed
k10s --readonly # look, never touch: nothing that changes the cluster
k10s --version # which build is this
What you get
Every resource, grouped the way you think about them
30 kinds across Workloads · Network · Config · Storage · RBAC · Cluster ·
Custom Resources , each with a live row count for the current namespace.
Your CRDs are discovered automatically - no configuration.
Groups fold: space , left or a click on the header, remembered across
restarts, with Config/Storage/RBAC folded to begin with. A folded group also
asks your cluster for nothing - the sidebar is what decides which kinds get
counted, and counting is one limit=1 request per visible kind, six at a
time, with refusals remembered.
Pods · Deployments · ReplicaSets · StatefulSets · DaemonSets · Jobs ·
CronJobs · HPAs · Services · Endpoints · Ingresses · NetworkPolicies ·
ConfigMaps · Secrets · ResourceQuotas · LimitRanges · PDBs · PVCs · PVs ·
StorageClasses · ServiceAccounts · Roles · RoleBindings · ClusterRoles ·
ClusterRoleBindings · Nodes · Namespaces · Events · CRDs · Custom Resources
The whole day-2 toolkit, on single keys
d describe
real kubectl describe output, from the API
y YAML
the live object
l logs
follows ( -f ), newest at the bottom, scroll back 500 lines
s shell
a real interactive exec session - raw TTY, resize-aware, in-panel
p port-forward
real SPDY forward, start/stop from the pane
m top
pod/node metrics, per container, with requests vs limits
e edit · r restart · c scale
rollout restart, scale, $EDITOR
o / u
cordon-uncordon / drain - offered only when Nodes is selected
D delete
red confirm modal, because it should be scary
Add your own scoped actions with the core k9s plugins.yaml format. They can
run foreground or background commands, request confirmation, override a
built-in shortcut deliberately, and are clickable in the same pane. See the
plugin guide and installable example .
Logs that follow, in the pane you were already looking at ( z to zoom)
╭─ logs -f billing-worker-6f8d9c5b7-qq91x ────────────────────────────────── [ close ] [ restore ] ╮
│ 8 2026-08-25T08:12:18.331Z INFO http GET /v1/users/me 200 3.1ms │
│ trace=44b1e2f9 │
│ 7 2026-08-25T08:12:21.660Z INFO worker flushed batch size=250 dur=41ms │
│ 6 2026-08-25T08:12:25.019Z INFO http GET /healthz 200 0.3ms │
│ 5 2026-08-25T08:12:31.402Z INFO http DELETE /v1/sessions/9a1 204 5.7ms │
│ trace=7c0d19ba │
│ 4 2026-08-25T08:12:33.881Z WARN gc pause=18ms heap=412Mi │
│ 3 2026-08-25T08:12:40.117Z INFO http GET /v1/orders/88213 200 7.4ms │
│ trace=e21f8b05 │
│ 2 2026-08-25T08:12:44.590Z INFO metrics scrape ok series=1842 │
│ 1 2026-08-25T08:12:51.008Z INFO http GET /healthz 200 0.3ms │
│ ● following newest at bottom 500 loaded · ↑ for older │
╰──────────────────────────────────────────────────────────────────────────────────────────────────╯
An AI prompt that already has the context
╭─ Prompt · plain text → AI · /commands still work · esc close ─ [ grow ] [ AI · claude-sonnet-5 ] ╮
│ ✦ ask about your cluster… · /settings to change provider/model │
╰──────────────────────────────────────────────────────────────────────────────────────────────────╯
● AI mode — plain text goes to claude-sonnet-5 tab panes · enter open · ctrl+p search · f find…
ctrl+a toggles it. Bring your own key - OpenAI-compatible (so also
Groq, Together, OpenRouter, vLLM, Ollama, LM Studio, any local gateway) or
Anthropic . Answers open as a normal text view you can scroll, zoom and
close. No cluster data leaves your machine unless you press enter in AI
mode. The only other network call k10s makes on its own is the once-a-day
update check, which asks GitHub for a version number and nothing else.
Eight built-in themes, plus your own — previewed live
tokyo-night (default) · catppuccin-mocha · dracula · nord ·
gruvbox-dark · solarized-dark · solarized-light · matrix
T cycles, /theme opens a picker that applies each theme as you move
through the list - you judge it on the real UI, not on a name - and esc
puts back whatever you had. Drop a YAML palette into ~/.k10s/themes , restart,
and it appears in the same picker — no rebuild needed. See the
custom-theme guide and installable demo .
Details that only show up after a long day
Copy mode ( ctrl+s ) releases the mouse so your terminal can drag-select
and copy - the one thing every mouse-capturing TUI breaks.
An honest loading state instead of "no resources found" while a watch
is still syncing.
Namespace and context pickers that never ask you to type a name.
No setup screen. First run opens the cluster; every setting has a
working default and /settings is one keystroke away when you want it.
A grow-able prompt ( ctrl+z ) - because a long kubectl line or an AI
question does not fit in a one-row field that scrolls sideways.
A terminal-too-small notice instead of a garbled layout.
k10s exists because of k9s . Credit where
it is due: k9s is mature, enormous in scope, plugin-extensible, and it is the
reason a whole generation of us stopped typing kubectl get pods all day.
k10s is younger and deliberately narrower. The difference is philosophy:
Use k9s if you want the biggest feature surface, its full plugin ecosystem,
and years of production mileage. Use k10s if you want something you can hand
to a teammate who has never opened a TUI, and have them find logs on their
own in ten seconds.
The short version - the full reference is here .
Two prefixes, and each one only ever shows its own set. / is k10s itself
(theme, settings, updates). : is the cluster, in the k9s vocabulary - a
resource view, the namespace, the context, or something acting on what is on
screen. enter runs the highlighted suggestion immediately - no second trip
through the prompt.
/theme /settings /mouse /update /version /demo /setup /help k10s itself
:po :deploy :

[truncated]

## Original Extract

Kubernetes TUI you can click. Instant search, logs/exec/port-forward, 7 themes, context-aware AI. Single Go binary. - p10node/k10s

GitHub - p10node/k10s: Kubernetes TUI you can click. Instant search, logs/exec/port-forward, 7 themes, context-aware AI. Single Go binary. · GitHub
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
p10node
k10s Public Notifications You must be signed in to change notification settings
Star 101 ( 101 ) You must be signed in to star a repository
Kubernetes TUI you can click. Instant search, logs/exec/port-forward, 7 themes, context-aware AI. Single Go binary.
Readme Apache-2.0 license Activity Custom properties Stars
7 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
28 Commits 28 Commits Folders and files
.github/ workflows .github/ workflows assets assets cmd/ shot cmd/ shot docs docs examples examples internal internal .gitignore .gitignore Justfile Justfile LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh main.go main.go main_test.go main_test.go uninstall.sh uninstall.sh View all files Repository files navigation
The Kubernetes terminal UI you can click .
k9s taught us to live in the terminal. k10s makes that terminal
point-and-click, instantly searchable, themeable - and gives it an AI that
already knows your cluster, namespace and selected object.
Install · Try it with no cluster ·
Features · k9s → k10s · Docs ·
Build log
A real terminal capture of k10s demo running - just screenshot , on the offline demo backend. Every frame in these docs comes out of the running binary, never hand-drawn.
A cluster dashboard is something you open twenty times a day, usually while
something is on fire. The two things that matter are how fast it opens
and how little you have to remember .
Nothing is hidden behind memorised keys. The actions that apply to the
thing you selected are listed, right there, in their own pane. Click one,
or press the letter next to it.
Your mouse works. Click a row, click a pane, click [ zoom ] , click
ns default ▾ , scroll the table. Every border button is real.
One search box for the whole cluster. ctrl+p searches resource kinds
and objects together. No prefix language to learn.
It opens instantly. Startup registers zero watches and waits for
nothing - informers start lazily, per kind, the first time you look at
one. Opening Pods watches pods, not every Secret and Event you own.
( how, and the regression guards )
AI that can see the screen. ctrl+a , ask in plain English. The current
context, namespace, kind and selected object are injected into the prompt,
so "why is this pod unhealthy?" means this pod.
It updates itself. /update installs the newest release over the
running binary - checksum-verified, atomic, offers to restart into it.
Your k9s-style command plugins fit. Put scoped shortcuts in
~/.k10s/plugins.yaml ; they appear beside built-in actions and receive the
selected object, namespace, context and column values.
Try it in 30 seconds (no cluster required)
k10s ships an offline demo backend: a realistic cluster to click around in,
including a CrashLoopBackOff to poke at. It is opt-in , because sample
data should never be mistaken for your machine.
git clone https://github.com/p10node/k10s && cd k10s
go run . demo # the sample cluster - fake data, clearly labelled
go run . # your real cluster, or "No cluster" if there isn't one
The demo is a context , not a mode. k10s demo opens on it, /demo
switches to it from anywhere, and :ctx always lists it (labelled
k10s demo · sample data , with a legend under the list). To leave it,
pick any other context - there is no separate exit. While it is up the
header carries a DEMO marker, so no frame of it can be mistaken for a real
cluster.
Plain k10s reads the same kubeconfig kubectl does ( $KUBECONFIG , else
~/.kube/config ) and shows what that context can reach - and nothing else.
With no kubeconfig, or a context whose API server does not answer, the main
panel says No cluster and points at the way in: r retries, :ctx picks
another context, /setup has the kubectl and kubeconfig links.
docs/cluster-setup.md is the same guide, longer.
curl -fsSL https://p10node.com/k10s/install.sh | sh
macOS and Linux, amd64 and arm64 . It picks the right prebuilt binary,
verifies its sha256 against the release manifest, and installs it into
/usr/local/bin (or ~/.local/bin when that needs a password it cannot
ask for). --dir , --version and --no-sudo are in docs/install.md ,
along with how to read it before you run it and the matching
uninstall.sh .
go install github.com/p10node/k10s@latest
Or from a clone, which also stamps the version into the binary:
git clone https://github.com/p10node/k10s && cd k10s
just install # → $GOBIN/k10s, version-stamped
Prebuilt static binaries for darwin/amd64, darwin/arm64, linux/amd64,
linux/arm64 and windows/amd64 are published on every tag — Windows is the
one platform the installer script sends to the
release page instead. Every later
upgrade is just:
k10s # your current kubeconfig context
k10s demo # the built-in sample cluster, no cluster needed
k10s --readonly # look, never touch: nothing that changes the cluster
k10s --version # which build is this
What you get
Every resource, grouped the way you think about them
30 kinds across Workloads · Network · Config · Storage · RBAC · Cluster ·
Custom Resources , each with a live row count for the current namespace.
Your CRDs are discovered automatically - no configuration.
Groups fold: space , left or a click on the header, remembered across
restarts, with Config/Storage/RBAC folded to begin with. A folded group also
asks your cluster for nothing - the sidebar is what decides which kinds get
counted, and counting is one limit=1 request per visible kind, six at a
time, with refusals remembered.
Pods · Deployments · ReplicaSets · StatefulSets · DaemonSets · Jobs ·
CronJobs · HPAs · Services · Endpoints · Ingresses · NetworkPolicies ·
ConfigMaps · Secrets · ResourceQuotas · LimitRanges · PDBs · PVCs · PVs ·
StorageClasses · ServiceAccounts · Roles · RoleBindings · ClusterRoles ·
ClusterRoleBindings · Nodes · Namespaces · Events · CRDs · Custom Resources
The whole day-2 toolkit, on single keys
d describe
real kubectl describe output, from the API
y YAML
the live object
l logs
follows ( -f ), newest at the bottom, scroll back 500 lines
s shell
a real interactive exec session - raw TTY, resize-aware, in-panel
p port-forward
real SPDY forward, start/stop from the pane
m top
pod/node metrics, per container, with requests vs limits
e edit · r restart · c scale
rollout restart, scale, $EDITOR
o / u
cordon-uncordon / drain - offered only when Nodes is selected
D delete
red confirm modal, because it should be scary
Add your own scoped actions with the core k9s plugins.yaml format. They can
run foreground or background commands, request confirmation, override a
built-in shortcut deliberately, and are clickable in the same pane. See the
plugin guide and installable example .
Logs that follow, in the pane you were already looking at ( z to zoom)
╭─ logs -f billing-worker-6f8d9c5b7-qq91x ────────────────────────────────── [ close ] [ restore ] ╮
│ 8 2026-08-25T08:12:18.331Z INFO http GET /v1/users/me 200 3.1ms │
│ trace=44b1e2f9 │
│ 7 2026-08-25T08:12:21.660Z INFO worker flushed batch size=250 dur=41ms │
│ 6 2026-08-25T08:12:25.019Z INFO http GET /healthz 200 0.3ms │
│ 5 2026-08-25T08:12:31.402Z INFO http DELETE /v1/sessions/9a1 204 5.7ms │
│ trace=7c0d19ba │
│ 4 2026-08-25T08:12:33.881Z WARN gc pause=18ms heap=412Mi │
│ 3 2026-08-25T08:12:40.117Z INFO http GET /v1/orders/88213 200 7.4ms │
│ trace=e21f8b05 │
│ 2 2026-08-25T08:12:44.590Z INFO metrics scrape ok series=1842 │
│ 1 2026-08-25T08:12:51.008Z INFO http GET /healthz 200 0.3ms │
│ ● following newest at bottom 500 loaded · ↑ for older │
╰──────────────────────────────────────────────────────────────────────────────────────────────────╯
An AI prompt that already has the context
╭─ Prompt · plain text → AI · /commands still work · esc close ─ [ grow ] [ AI · claude-sonnet-5 ] ╮
│ ✦ ask about your cluster… · /settings to change provider/model │
╰──────────────────────────────────────────────────────────────────────────────────────────────────╯
● AI mode — plain text goes to claude-sonnet-5 tab panes · enter open · ctrl+p search · f find…
ctrl+a toggles it. Bring your own key - OpenAI-compatible (so also
Groq, Together, OpenRouter, vLLM, Ollama, LM Studio, any local gateway) or
Anthropic . Answers open as a normal text view you can scroll, zoom and
close. No cluster data leaves your machine unless you press enter in AI
mode. The only other network call k10s makes on its own is the once-a-day
update check, which asks GitHub for a version number and nothing else.
Eight built-in themes, plus your own — previewed live
tokyo-night (default) · catppuccin-mocha · dracula · nord ·
gruvbox-dark · solarized-dark · solarized-light · matrix
T cycles, /theme opens a picker that applies each theme as you move
through the list - you judge it on the real UI, not on a name - and esc
puts back whatever you had. Drop a YAML palette into ~/.k10s/themes , restart,
and it appears in the same picker — no rebuild needed. See the
custom-theme guide and installable demo .
Details that only show up after a long day
Copy mode ( ctrl+s ) releases the mouse so your terminal can drag-select
and copy - the one thing every mouse-capturing TUI breaks.
An honest loading state instead of "no resources found" while a watch
is still syncing.
Namespace and context pickers that never ask you to type a name.
No setup screen. First run opens the cluster; every setting has a
working default and /settings is one keystroke away when you want it.
A grow-able prompt ( ctrl+z ) - because a long kubectl line or an AI
question does not fit in a one-row field that scrolls sideways.
A terminal-too-small notice instead of a garbled layout.
k10s exists because of k9s . Credit where
it is due: k9s is mature, enormous in scope, plugin-extensible, and it is the
reason a whole generation of us stopped typing kubectl get pods all day.
k10s is younger and deliberately narrower. The difference is philosophy:
Use k9s if you want the biggest feature surface, its full plugin ecosystem,
and years of production mileage. Use k10s if you want something you can hand
to a teammate who has never opened a TUI, and have them find logs on their
own in ten seconds.
The short version - the full reference is here .
Two prefixes, and each one only ever shows its own set. / is k10s itself
(theme, settings, updates). : is the cluster, in the k9s vocabulary - a
resource view, the namespace, the context, or something acting on what is on
screen. enter runs the highlighted suggestion immediately - no second trip
through the prompt.
/theme /settings /mouse /update /version /demo /setup /help k10s itself
:po :deploy :

[truncated]
