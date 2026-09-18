---
source: "https://github.com/unclebob/uml-viewer"
hn_url: "https://news.ycombinator.com/item?id=49759501"
title: "Uncle Bob Martin's UML Tool to Manage Grok AI Agents"
article_title: "GitHub - unclebob/uml-viewer: Live Quil UML class-diagram viewer driven by EDN · GitHub"
image: "https://opengraph.githubassets.com/42fcbb3ec6d976f27f43eda3683a0ce40700234ef6e8f2b98dbc54399b7b9dc7/unclebob/uml-viewer"
author: "rmason"
captured_at: "2026-09-18T20:30:22Z"
capture_tool: "hn-digest"
hn_id: 49759501
score: 1
comments: 0
posted_at: "2026-09-18T20:05:25Z"
tags:
  - hacker-news
---

# Uncle Bob Martin's UML Tool to Manage Grok AI Agents

- HN: [49759501](https://news.ycombinator.com/item?id=49759501)
- Source: [github.com](https://github.com/unclebob/uml-viewer)
- Score: 1
- Comments: 0
- Posted: 2026-09-18T20:05:25Z

## Translation

Title: Uncle Bob Martin's UML Tool to Manage Grok AI Agents
Article title: GitHub - unclebob/uml-viewer: Live Quil UML class-diagram viewer driven by EDN · GitHub
Description: Live Quil UML class-diagram viewer driven by EDN. Contribute to unclebob/uml-viewer development by creating an account on GitHub.

Article text:
GitHub - unclebob/uml-viewer: Live Quil UML class-diagram viewer driven by EDN · GitHub
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
unclebob
/
uml-viewer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
31 Commits 31 Commits Folders and files
.grok/ rules .grok/ rules .metrics .metrics examples examples spec/ uml_viewer spec/ uml_viewer src/ uml_viewer src/ uml_viewer .gitignore .gitignore README.md README.md deps.edn deps.edn View all files Repository files navigation
A live Quil app that lays out and draws UML from an EDN IR. A policy plus
a language-specific parser write the topology; this tool displays it, routes
the arrows, colors CRAP, and lets you click.
The IR is topology : the namespace tree, classes, and edges. Metrics
(CC, coverage, CRAP, killed/survived) come from .metrics/ snapshots produced
by crap4clj and
clj-mutate . The viewer overlays those
files at load, keyed by namespace + function name. Agents edit the policy, not
the IR. See Policy .
Needs Clojure CLI and Java 21+.
clj -M:ir # policy → examples/uml-viewer.edn
clj -M:run
clj -M:run examples/library.edn
clj -M:run examples/uml-viewer.edn
clj -M:run --help
A tmux session uml-viewer-grok starts interactive Grok in the
examined project's directory ( --yolo --trust --rules … plus a launch
prompt). On start it writes a hierarchical policy from that project's
namespaces and regenerates the IR. Type there; Esc is the real TUI interrupt.
Closing the diagram kills that tmux session (and the Terminal attach). That
instance — not every Grok in this repo — also runs clj -M:crap ,
clj -M:mutate , and IR generate after later changes. Project-wide rules live
in .grok/rules/uml-viewer.md .
The examined project (and this one) must expose two aliases:
Do not pass --restart (or use :uml-viewer-restart ) unless you are that
companion recycling the window after source changes. A stray --restart
skips spawning Grok and leaves a diagram with no agent. The companion
recycles the window by writing :quit-for-restart to
.uml-viewer/to-viewer.edn , waiting for the JVM to exit, then
clj -M:uml-viewer-restart . Do not SIGKILL. Closing the window still kills
Grok.
On a fresh start the canvas stays blank until the companion sends :display ,
with Waiting for agent to create diagram. R reloads the current EDN
immediately and does not wait. A missing or unreadable file prints
UML viewer: file not found: … in the inspector instead of throwing.
clj -M:spec
clj -M:cov
clj -M:ir # writes examples/uml-viewer.edn
clj -M:crap # writes .metrics/crap.edn
clj -M:mutate src/uml_viewer/engine/layout.clj
This project's :crap alias uses ../clojure/crap4clj . :mutate pins
clj-mutate by git SHA. Commit
.metrics/ so a clone has numbers without re-running those tools.
Rename or move of a function is a new form: overlay does not match old names.
Layer and component mean the same thing: a namespace grouping
(the first segment after the prefix, or a named proposal group).
First view: namespace components (layers). Dependencies between them
collapse to one arrow. Each component lists nested namespaces.
Double-click a component to open the next level. Esc or the ← label goes up.
Hover an arrow for a popup of every from -> to it bundles, in any
declutter mode. Violating pairs are red.
The inspector lists proposals . Click one to show it (marked as not
in the code). P returns to the ns tree. New Proposal adds a
timestamp-named proposal. Right-click to rename or delete. Declutter
cycles Declutter arrows / Declutter elements / Declutter classes /
Declutter none.
Double-click a leaf module for its class card .
The class card names the module ( :ns ). Click it to open that source
file at the top. Hover a member to highlight it; click it to open the same
file at the defn. See Source extractors .
Methods on the card are + public and - private. defn- is not drawn on
the class box.
Abstract classes show a white α in the upper-right; interfaces a white
I . Names of rectangles that are not classes (components/layers,
interfaces, enumerations, package banners) are italic. Foreign libraries
listed in policy are ovals outside the components.
Scroll to pan vertically; Shift-scroll (or left/right arrows) for
horizontal. Pan can follow arrows that bow past the origin.
Ctrl+ (or Ctrl+= ) zooms in 10%; Ctrl- zooms out 10%;
Ctrl+0 restores 100%. Zoom keeps the view center still.
Regen in the inspector asks the companion to rewrite policy and IR
(see Companion mailbox ).
R reloads the current EDN (the watcher also reloads on save). Overlay
re-reads .metrics/ on the next load.
Esc on the class card closes it. Closing the main window exits the app.
This project's diagram is generated . Do not edit examples/uml-viewer.edn .
Edit examples/uml-viewer.policy.edn , then run clj -M:ir (or press Regen).
The parser ( LanguageGraph ) reads source and emits facts: one class per
project namespace, :require / :use of another project ns as
:dependency , requiring-resolve of a quoted var as :dependency on that
var's namespace, defprotocol as :stereotype :interface , defrecord /
deftype of a protocol as :implements . External :require s and :import s
become foreign classes. Members are not authored — overlay fills them from
.metrics/ .
Do not invent layers (components)
The tree is the namespaces. After :prefix , every . is a nesting
level. uml-viewer.engine.layout is a child of engine .
uml-viewer.clojure-language.source-clojure is a child of clojure-language .
The policy does not assign nses to invented packages. If you want Domain /
Engine / Adapters boxes in the source tree , those segments must exist as
namespaces. To view a grouping that is not in the code, use :proposal
(see Proposed components ) — do not rewrite namespaces.
To write a policy for a project:
Set :src and :prefix to the project's source root and ns prefix
( src and foo for foo.bar.baz ).
Set :hierarchical true (or omit :packages and :diagrams ).
List top-level segments in :order — the first dotted part after
the prefix, in the order you want the boxes. Do not invent names.
List real libraries in :foreign if they should appear as ovals.
Optionally override a require with :edge-kinds {[:from :to] :association}
using the leaf ids ( clojure-language.source-clojure , not
clojure-language ).
Set :levels so the generator can mark dependency-rule violations
(see Dependency rule ).
Optionally set :proposal to name design components that are not namespaces
(see Proposed components ).
If foo.bar and foo.bar.baz both exist, the bar box lists bar (the
module) and baz (the child). Double-click the component to open that
level; double-click the bar module line for its class card.
:packages [{ :id :domain :nses [ir geom source]}
{ :id :engine :nses [layout route]}]
Right (the ns tree):
{ :title " UML viewer "
:src " src "
:prefix " uml-viewer "
:lang :clojure
:out " examples/uml-viewer.edn "
:hierarchical true
:foreign [quil]
:order [main adapters application engine source graph clojure-language domain]
:levels [[domain source graph clojure-language]
[engine]
[application]
[adapters]
[main]]
:edge-kinds {[ :engine.compose :engine.layout ] :association }}
Key
Role
:prefix
Strip this from each ns; remaining dots are the tree
:hierarchical
Namespace tree (default when :packages is omitted)
:order
Order of existing top-level ns segments, not new component names
:levels
Groups of those segments, inner (higher-level) first . Same group = same rank
:proposals
Named groupings of real segments; not instantiated in source. Inspector list; P returns to the ns tree
:edge-kinds
Override parser kind for [from to] (usually :association )
:omit-edges
Drop [from to]
:lang
Which LanguageGraph to use (default :clojure )
:foreign
External libs as ovals. A listed prefix collapses quil.core to quil .
Viewer Grok loop (passed with --rules to the companion session only)
On launch: from the examined directory, write or update the hierarchical
policy and regenerate the IR, then wait.
After every later source or policy change: clj -M:crap , clj -M:mutate
on the changed src/ files, then clj -M:ir . Uncovered mutants remaining are
coverage gaps; keep the snapshot and do not re-run the file or force a full
mutation because mutate exited non-zero.
Add/rename/delete a namespace: the tree updates on clj -M:ir . Put a new
top-level segment in :order if you care about box order.
Nested nses appear as contents of the parent component.
“This require is really an association”: one :edge-kinds entry.
Show a library like quil as an oval: add it to :foreign .
Do not add :packages to fake Clean Architecture components. Use
:proposals to view a grouping that is not in the code.
Preserve :proposals when rewriting policy. Do not invent them on launch.
If instructed, add a named proposal (default name is a timestamp).
Hand-written sample IRs (e.g. examples/library.edn ) are still valid; they
are not generated.
A :dependency edge is violating when it runs from a higher-level
(inner) component to a lower-level (outer) one. That is the Clean
Architecture dependency rule: source-code dependencies point inward.
Evaluation is deterministic given :levels :
Take the first dotted segment of each end ( engine.layout → engine ).
Look up that segment in :levels . Rank is the group's index; smaller
is inner / higher-level .
If both ends have a rank and from-rank < to-rank , the edge is
:violating true . Same rank is allowed. :implements and
:association are never violating. Foreign / unranked ends are not
compared.
Collapsed component arrows keep the flag if any bundled leaf dependency
was violating. Remapping a pair to :association clears it.
:order is visual box order, not rank. Nesting is not layering: you cannot
infer inner vs outer from the namespace tree alone, so :levels must group
segments that sit at the same architectural level (e.g. domain , source ,
and graph ). Omit :levels and nothing is marked. If :levels is omitted
and :proposals is set, rank follows the first proposal's component order.
:proposals is a list of named groupings of existing top-level segments.
Those names are not namespaces. Each item is {:id :name :layers [...]}
( :layers here are named components). The as-is diagram stays the ns tree.
The inspector lists the real diagram (the namespace tree) just above
Proposals ; click it to return to the tree. Click a proposal to show it
(canvas marked PROPOSAL — not instantiated in code ). New Proposal
adds an empty proposal named with a timestamp. Right-click a name to rename
or delete it. Double-click a ns box to drill the real tree.
The Declutter button cycles Declutter arrows (one arrow per
component pair per direction) → Declutter elements (al

[truncated]

## Original Extract

Live Quil UML class-diagram viewer driven by EDN. Contribute to unclebob/uml-viewer development by creating an account on GitHub.

GitHub - unclebob/uml-viewer: Live Quil UML class-diagram viewer driven by EDN · GitHub
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
unclebob
/
uml-viewer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
31 Commits 31 Commits Folders and files
.grok/ rules .grok/ rules .metrics .metrics examples examples spec/ uml_viewer spec/ uml_viewer src/ uml_viewer src/ uml_viewer .gitignore .gitignore README.md README.md deps.edn deps.edn View all files Repository files navigation
A live Quil app that lays out and draws UML from an EDN IR. A policy plus
a language-specific parser write the topology; this tool displays it, routes
the arrows, colors CRAP, and lets you click.
The IR is topology : the namespace tree, classes, and edges. Metrics
(CC, coverage, CRAP, killed/survived) come from .metrics/ snapshots produced
by crap4clj and
clj-mutate . The viewer overlays those
files at load, keyed by namespace + function name. Agents edit the policy, not
the IR. See Policy .
Needs Clojure CLI and Java 21+.
clj -M:ir # policy → examples/uml-viewer.edn
clj -M:run
clj -M:run examples/library.edn
clj -M:run examples/uml-viewer.edn
clj -M:run --help
A tmux session uml-viewer-grok starts interactive Grok in the
examined project's directory ( --yolo --trust --rules … plus a launch
prompt). On start it writes a hierarchical policy from that project's
namespaces and regenerates the IR. Type there; Esc is the real TUI interrupt.
Closing the diagram kills that tmux session (and the Terminal attach). That
instance — not every Grok in this repo — also runs clj -M:crap ,
clj -M:mutate , and IR generate after later changes. Project-wide rules live
in .grok/rules/uml-viewer.md .
The examined project (and this one) must expose two aliases:
Do not pass --restart (or use :uml-viewer-restart ) unless you are that
companion recycling the window after source changes. A stray --restart
skips spawning Grok and leaves a diagram with no agent. The companion
recycles the window by writing :quit-for-restart to
.uml-viewer/to-viewer.edn , waiting for the JVM to exit, then
clj -M:uml-viewer-restart . Do not SIGKILL. Closing the window still kills
Grok.
On a fresh start the canvas stays blank until the companion sends :display ,
with Waiting for agent to create diagram. R reloads the current EDN
immediately and does not wait. A missing or unreadable file prints
UML viewer: file not found: … in the inspector instead of throwing.
clj -M:spec
clj -M:cov
clj -M:ir # writes examples/uml-viewer.edn
clj -M:crap # writes .metrics/crap.edn
clj -M:mutate src/uml_viewer/engine/layout.clj
This project's :crap alias uses ../clojure/crap4clj . :mutate pins
clj-mutate by git SHA. Commit
.metrics/ so a clone has numbers without re-running those tools.
Rename or move of a function is a new form: overlay does not match old names.
Layer and component mean the same thing: a namespace grouping
(the first segment after the prefix, or a named proposal group).
First view: namespace components (layers). Dependencies between them
collapse to one arrow. Each component lists nested namespaces.
Double-click a component to open the next level. Esc or the ← label goes up.
Hover an arrow for a popup of every from -> to it bundles, in any
declutter mode. Violating pairs are red.
The inspector lists proposals . Click one to show it (marked as not
in the code). P returns to the ns tree. New Proposal adds a
timestamp-named proposal. Right-click to rename or delete. Declutter
cycles Declutter arrows / Declutter elements / Declutter classes /
Declutter none.
Double-click a leaf module for its class card .
The class card names the module ( :ns ). Click it to open that source
file at the top. Hover a member to highlight it; click it to open the same
file at the defn. See Source extractors .
Methods on the card are + public and - private. defn- is not drawn on
the class box.
Abstract classes show a white α in the upper-right; interfaces a white
I . Names of rectangles that are not classes (components/layers,
interfaces, enumerations, package banners) are italic. Foreign libraries
listed in policy are ovals outside the components.
Scroll to pan vertically; Shift-scroll (or left/right arrows) for
horizontal. Pan can follow arrows that bow past the origin.
Ctrl+ (or Ctrl+= ) zooms in 10%; Ctrl- zooms out 10%;
Ctrl+0 restores 100%. Zoom keeps the view center still.
Regen in the inspector asks the companion to rewrite policy and IR
(see Companion mailbox ).
R reloads the current EDN (the watcher also reloads on save). Overlay
re-reads .metrics/ on the next load.
Esc on the class card closes it. Closing the main window exits the app.
This project's diagram is generated . Do not edit examples/uml-viewer.edn .
Edit examples/uml-viewer.policy.edn , then run clj -M:ir (or press Regen).
The parser ( LanguageGraph ) reads source and emits facts: one class per
project namespace, :require / :use of another project ns as
:dependency , requiring-resolve of a quoted var as :dependency on that
var's namespace, defprotocol as :stereotype :interface , defrecord /
deftype of a protocol as :implements . External :require s and :import s
become foreign classes. Members are not authored — overlay fills them from
.metrics/ .
Do not invent layers (components)
The tree is the namespaces. After :prefix , every . is a nesting
level. uml-viewer.engine.layout is a child of engine .
uml-viewer.clojure-language.source-clojure is a child of clojure-language .
The policy does not assign nses to invented packages. If you want Domain /
Engine / Adapters boxes in the source tree , those segments must exist as
namespaces. To view a grouping that is not in the code, use :proposal
(see Proposed components ) — do not rewrite namespaces.
To write a policy for a project:
Set :src and :prefix to the project's source root and ns prefix
( src and foo for foo.bar.baz ).
Set :hierarchical true (or omit :packages and :diagrams ).
List top-level segments in :order — the first dotted part after
the prefix, in the order you want the boxes. Do not invent names.
List real libraries in :foreign if they should appear as ovals.
Optionally override a require with :edge-kinds {[:from :to] :association}
using the leaf ids ( clojure-language.source-clojure , not
clojure-language ).
Set :levels so the generator can mark dependency-rule violations
(see Dependency rule ).
Optionally set :proposal to name design components that are not namespaces
(see Proposed components ).
If foo.bar and foo.bar.baz both exist, the bar box lists bar (the
module) and baz (the child). Double-click the component to open that
level; double-click the bar module line for its class card.
:packages [{ :id :domain :nses [ir geom source]}
{ :id :engine :nses [layout route]}]
Right (the ns tree):
{ :title " UML viewer "
:src " src "
:prefix " uml-viewer "
:lang :clojure
:out " examples/uml-viewer.edn "
:hierarchical true
:foreign [quil]
:order [main adapters application engine source graph clojure-language domain]
:levels [[domain source graph clojure-language]
[engine]
[application]
[adapters]
[main]]
:edge-kinds {[ :engine.compose :engine.layout ] :association }}
Key
Role
:prefix
Strip this from each ns; remaining dots are the tree
:hierarchical
Namespace tree (default when :packages is omitted)
:order
Order of existing top-level ns segments, not new component names
:levels
Groups of those segments, inner (higher-level) first . Same group = same rank
:proposals
Named groupings of real segments; not instantiated in source. Inspector list; P returns to the ns tree
:edge-kinds
Override parser kind for [from to] (usually :association )
:omit-edges
Drop [from to]
:lang
Which LanguageGraph to use (default :clojure )
:foreign
External libs as ovals. A listed prefix collapses quil.core to quil .
Viewer Grok loop (passed with --rules to the companion session only)
On launch: from the examined directory, write or update the hierarchical
policy and regenerate the IR, then wait.
After every later source or policy change: clj -M:crap , clj -M:mutate
on the changed src/ files, then clj -M:ir . Uncovered mutants remaining are
coverage gaps; keep the snapshot and do not re-run the file or force a full
mutation because mutate exited non-zero.
Add/rename/delete a namespace: the tree updates on clj -M:ir . Put a new
top-level segment in :order if you care about box order.
Nested nses appear as contents of the parent component.
“This require is really an association”: one :edge-kinds entry.
Show a library like quil as an oval: add it to :foreign .
Do not add :packages to fake Clean Architecture components. Use
:proposals to view a grouping that is not in the code.
Preserve :proposals when rewriting policy. Do not invent them on launch.
If instructed, add a named proposal (default name is a timestamp).
Hand-written sample IRs (e.g. examples/library.edn ) are still valid; they
are not generated.
A :dependency edge is violating when it runs from a higher-level
(inner) component to a lower-level (outer) one. That is the Clean
Architecture dependency rule: source-code dependencies point inward.
Evaluation is deterministic given :levels :
Take the first dotted segment of each end ( engine.layout → engine ).
Look up that segment in :levels . Rank is the group's index; smaller
is inner / higher-level .
If both ends have a rank and from-rank < to-rank , the edge is
:violating true . Same rank is allowed. :implements and
:association are never violating. Foreign / unranked ends are not
compared.
Collapsed component arrows keep the flag if any bundled leaf dependency
was violating. Remapping a pair to :association clears it.
:order is visual box order, not rank. Nesting is not layering: you cannot
infer inner vs outer from the namespace tree alone, so :levels must group
segments that sit at the same architectural level (e.g. domain , source ,
and graph ). Omit :levels and nothing is marked. If :levels is omitted
and :proposals is set, rank follows the first proposal's component order.
:proposals is a list of named groupings of existing top-level segments.
Those names are not namespaces. Each item is {:id :name :layers [...]}
( :layers here are named components). The as-is diagram stays the ns tree.
The inspector lists the real diagram (the namespace tree) just above
Proposals ; click it to return to the tree. Click a proposal to show it
(canvas marked PROPOSAL — not instantiated in code ). New Proposal
adds an empty proposal named with a timestamp. Right-click a name to rename
or delete it. Double-click a ns box to drill the real tree.
The Declutter button cycles Declutter arrows (one arrow per
component pair per direction) → Declutter elements (al

[truncated]
