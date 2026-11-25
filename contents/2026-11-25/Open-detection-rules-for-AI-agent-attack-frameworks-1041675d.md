---
source: "https://github.com/ThreatDefence/agentic-attack-detections"
hn_url: "https://news.ycombinator.com/item?id=50637987"
title: "Open detection rules for AI agent attack frameworks"
article_title: "GitHub - ThreatDefence/agentic-attack-detections: Detection rules and IOCs for AI agent attack frameworks: ARTEX, Cairn, PentAGI, HexStrike, CyberStrikeAI, Strix, Hermes Agent, Agent Zero, RedAmon, Xalgorix, Pentest Swarm AI. KQL hunts for Microsoft Sentinel and Defender XDR, Sigma, Suricata, Zeek,\n[truncated]"
image: "https://repository-images.githubusercontent.com/1464795219/d1137601-3acd-4e93-a98c-f5d3bb45ef7a"
author: "AntonGuzhevskiy"
captured_at: "2026-11-25T03:29:31Z"
capture_tool: "hn-digest"
hn_id: 50637987
score: 1
comments: 0
posted_at: "2026-11-25T03:18:00Z"
tags:
  - hacker-news
---

# Open detection rules for AI agent attack frameworks

- HN: [50637987](https://news.ycombinator.com/item?id=50637987)
- Source: [github.com](https://github.com/ThreatDefence/agentic-attack-detections)
- Score: 1
- Comments: 0
- Posted: 2026-11-25T03:18:00Z

## Translation

Title: Open detection rules for AI agent attack frameworks
Article title: GitHub - ThreatDefence/agentic-attack-detections: Detection rules and IOCs for AI agent attack frameworks: ARTEX, Cairn, PentAGI, HexStrike, CyberStrikeAI, Strix, Hermes Agent, Agent Zero, RedAmon, Xalgorix, Pentest Swarm AI. KQL hunts for Microsoft Sentinel and Defender XDR, Sigma, Suricata, Zeek,
[truncated]
Description: Detection rules and IOCs for AI agent attack frameworks: ARTEX, Cairn, PentAGI, HexStrike, CyberStrikeAI, Strix, Hermes Agent, Agent Zero, RedAmon, Xalgorix, Pentest Swarm AI. KQL hunts for Microsoft Sentinel and Defender XDR, Sigma, Suricata, Zeek, YARA, Splunk, Elastic, STIX 2.1. - ThreatDefence/a
[truncated]

Article text:
GitHub - ThreatDefence/agentic-attack-detections: Detection rules and IOCs for AI agent attack frameworks: ARTEX, Cairn, PentAGI, HexStrike, CyberStrikeAI, Strix, Hermes Agent, Agent Zero, RedAmon, Xalgorix, Pentest Swarm AI. KQL hunts for Microsoft Sentinel and Defender XDR, Sigma, Suricata, Zeek, YARA, Splunk, Elastic, STIX 2.1. · GitHub
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
ThreatDefence
agentic-attack-detections Public Notifications You must be signed in to change notification settings
Star 632 ( 632 ) You must be signed in to star a repository
Detection rules and IOCs for AI agent attack frameworks: ARTEX, Cairn, PentAGI, HexStrike, CyberStrikeAI, Strix, Hermes Agent, Agent Zero, RedAmon, Xalgorix, Pentest Swarm AI. KQL hunts for Microsoft Sentinel and Defender XDR, Sigma, Suricata, Zeek, YARA, Splunk, Elastic, STIX 2.1.
threatdefence.com/article/artex-agentic-attack-framework-detection-hunting Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
132 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
78 Commits 78 Commits Folders and files
.github .github behaviors behaviors deploy deploy docs docs exposure exposure feeds feeds frameworks frameworks generated generated iocs iocs lab lab tests tests tools tools waf waf .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md zkg.meta zkg.meta View all files Repository files navigation
Open-source detection rules, hunting queries and indicators for security teams. They cover agentic attack frameworks: tools that hand an intrusion to AI agents, which plan, run real attack tools and keep going until they reach a goal.
Free to use under the MIT license. Fixes, false-positive reports and new indicators are welcome.
Framework folders. ARTEX and its AntiFox fork, Cairn, PentAGI, HexStrike, CyberStrikeAI, Strix, Hermes Agent, Agent Zero, RedAmon, Xalgorix and Pentest Swarm AI. Each folder is built from a review of the framework's source code and the public reporting on its use.
Shared behaviors. Rules for what these frameworks have in common, whatever the tool or the model behind it.
Indicators and feeds. One CSV of 580 sourced indicators, and the feeds generated from it.
Exposed consoles. A dated list of addresses where one of these consoles answers on the internet, in exposure/ , with a count per framework . It is a record of what is exposed, not a detection feed.
Formats: KQL for Microsoft Sentinel and Defender XDR, Sigma, Suricata, Zeek, YARA and STIX 2.1, plus WAF rules for Cloudflare, AWS WAF and Azure WAF. Splunk has one app for every framework. Elastic has one import with a rule for every framework.
With Suricata installed, run the ARTEX rules over the sample capture in this repository:
git clone https://github.com/ThreatDefence/agentic-attack-detections
cd agentic-attack-detections
mkdir -p out
suricata -r frameworks/artex/tests/artex-sample.pcap -S frameworks/artex/suricata/artex.rules -l out -k none
cat out/fast.log
You should see alerts such as:
[1:3800001:1] TD ARTEX user agent artex-<role>/1.0
[1:3800020:1] TD ARTEX placeholder probe path
[1:3800023:2] TD Unsigned alg none JWT in URI
The capture is test traffic written from the framework source code and published reports. It is not a recorded intrusion.
You have
Do this
Microsoft Sentinel
One deployment creates every analytics rule (disabled until you enable them), every hunting query and the AGENTIC_IOCs watchlist. Give it the workspace name. deploy/sentinel/ has one template per framework too. Run the hunts in each folder's hunting-order.md order
Defender XDR
Run the hunts from frameworks/*/sentinel/hunting/ in advanced hunting. They work where the tables exist
Suricata
Add the .rules files from frameworks/*/suricata/ and behaviors/suricata/ . HTTP rules need cleartext or decrypted traffic. DNS and TLS SNI rules work on encrypted traffic
Zeek
Install the behaviors package with zkg install https://github.com/ThreatDefence/agentic-attack-detections (or load behaviors/zeek/agentic-behaviors/ by path), set server_nets to your server ranges ( setup ), load the framework scripts frameworks/artex/zeek/artex/ , frameworks/cyberstrikeai/zeek/cyberstrikeai/ and frameworks/hexstrike/zeek/hexstrike/ , and feed frameworks/*/zeek/intel/*.intel to the Intel framework
YARA
Scan container layers, home folders, /tmp and web roots with frameworks/*/yara/
Any SIEM with Sigma
Convert frameworks/*/sigma/ and behaviors/sigma/ with sigma convert , or take the conversions in generated/
Splunk
Install agentic-attack-detections-splunk.tgz from the latest release (Apps > Install app from file). One app has the searches for every framework, the indicator lookups and a dashboard with a card per framework. Point the aad_indexes macro at your indexes. Setup
Elastic
Import agentic-attack-detections.ndjson through Security > Rules > Import, or post it to /api/detection_engine/rules/_import . One file carries every framework and the shared behaviors, all disabled: the Sigma rules as query rules, the correlations as ES|QL rules, the hand-written ARTEX rules, and an indicator match rule. For that one, load agentic-attack-indicators.ndjson with POST logs-ti_agentic_attacks-default/_bulk . Point the index patterns at your own data before you enable th
[truncated]
KQL hunts
Sentinel analytics
Sigma
Suricata
YARA
Zeek
Indicators
ARTEX
31
7
36
43
4
script, intel
247
Cairn
6
2
3
26
6
intel
73
PentAGI
5
2
3
10
3
intel
39
HexStrike
4
2
3
7
3
script, intel
22
CyberStrikeAI
5
2
7
11
4
script, intel
74
Strix
3
1
3
5
3
intel
49
Hermes Agent
4
2
3
7
4
intel
79
Agent Zero
2
1
3
1
intel
8
RedAmon
3
3
10
1
intel
19
Xalgorix
3
3
10
1
intel
24
Pentest Swarm AI
3
3
11
1
intel
15
Shared behaviors
9
2
6
18
package
Suricata counts are hand-written rules; each framework also has address rules generated from its indicators. Some indicators belong to more than one framework. Every framework is in the Splunk app . ARTEX also has Elastic rules and ES|QL queries, and an osquery pack.
Framework
What it is
Seen in
ARTEX
Go planner and worker agents with a web console. AntiFox is the same code renamed
Storm-0940, Storm-3358, VAPOR PANDA and criminal groups, against Microsoft 365, Okta, AWS, Azure, web applications and Active Directory
Cairn
Python engine that drives Claude Code, Codex and Pi coding agents in Kali containers
Online retailers, payment sites, Japanese web applications and WSUS servers
PentAGI
Go multi-agent platform that runs Kali tools in a container per engagement
Intrusions GTIG and Mandiant reported since August 2026, and a PaperCut exploitation operator's infrastructure
HexStrike
Python MCP server that gives any AI assistant about 150 security tools over an unauthenticated API
NetScaler exploitation discussion in 2025, and intrusions GTIG and Mandiant reported
CyberStrikeAI
Go operator console whose agents run about 90 tools, with its own implants, C2 listeners and webshell manager
FortiGate mass-exploitation infrastructure, an AWS intrusion, and education organizations in Africa and Latin America
Strix
Python AI pentest agent that runs teams of agents in a Kali sandbox, on any model
The retail card-theft campaign Gambit reported, where it found the flaws Cairn exploited
Hermes Agent
A legitimate general-purpose agent turned into an attack console: approvals off, a bundled jailbreak skill, red-team personas
The Gambit retail operator, the Carbonato Docker botnet, and Thailand's Ministry of Finance
Agent Zero
General-purpose agent in a Kali container with a built-in hacker profile and an open web UI
Threat-actor experimentation GTIG reported, with no published indicators
RedAmon
Self-hosted penetration testing platform: a web console, a recon pipeline and an agent that exploits from a Kali sandbox with Metasploit
No public report of use in an attack. Covered from its source code
Xalgorix
Go binary with a web dashboard whose agent tests web applications and APIs without approval steps and re-exploits each finding
No public report o
[truncated]
Hosts on the internet that serve the web console of each framework, found in internet-wide scan data with the fingerprints the rules here use: favicons, page titles, page text that survives a rename, and certificates.
An active host is one seen in the seven days to November 23, 2026. Hosts the scan data flags as honeypots are left out of both counts.
An exposed console is not proof of an attack. Researchers, students and security teams run these tools, and some hosts are decoys.
More than half of the pages titled CyberStrikeAI are flagged as honeypots, which is why that count is so much lower than a title search suggests. The flagged hosts stay in the list with honeypot set to yes .
One network that is not flagged holds more than a third of the remaining CyberStrikeAI hosts and of the PentAGI hosts. That looks like one mass deployment or a set of decoys, not that many operators ( #21 ).
Hermes Agent and Agent Zero are general-purpose agents, so their numbers count exposed agents, not attack consoles. Hermes is counted on its default port only.
Consoles served by host name only, or on a port that has not been scanned, are not counted.
The addresses, with the port and the fingerprint that found each one, are in exposure/consoles.csv . The counts of every snapshot are kept in exposure/history.csv , so the numbers can be followed over time.
Behaviors shared by agentic frameworks
behaviors/ holds rules that do not depend on any one framework.
The ARTEX folder also carries behavior hunts that apply to every framework: user agent family bursts (hunt 17), web and database services starting shells (hunt 11), servers calling webhook and tunnel services (hunt 12), scripts written into web folders (hunt 13), and Kerberoasting, DCSync and WinRM fan-out (hunt 16).
Path
Contents
iocs/indicators.csv
580 indicators. Each row has its source, confidence, a review date for addresses and the framework it belongs to. This file is the single source for every feed
feeds/
Generated from the CSV: a STIX 2.1 bundle, a Sentinel watchlist, and plain lists of addresses, domains, URLs, user agents and hashes. The plain lists carry medium and high confidence rows only. Low confidence rows, which include legitimate services such as AI model APIs, need their context, so they stay in the CSV, the watchlist and the STIX bundle. feeds/by-framework/ has the lists per framework
generated/
Every Sigma rule converted to Splunk SPL, Sentinel and Defender XDR KQL, and Elastic ES|QL and Lucene, plus elastic_ndjson/ where all of them are one Kibana import. SUMMARY.md shows which rule

[truncated]

## Original Extract

Detection rules and IOCs for AI agent attack frameworks: ARTEX, Cairn, PentAGI, HexStrike, CyberStrikeAI, Strix, Hermes Agent, Agent Zero, RedAmon, Xalgorix, Pentest Swarm AI. KQL hunts for Microsoft Sentinel and Defender XDR, Sigma, Suricata, Zeek, YARA, Splunk, Elastic, STIX 2.1. - ThreatDefence/a
[truncated]

GitHub - ThreatDefence/agentic-attack-detections: Detection rules and IOCs for AI agent attack frameworks: ARTEX, Cairn, PentAGI, HexStrike, CyberStrikeAI, Strix, Hermes Agent, Agent Zero, RedAmon, Xalgorix, Pentest Swarm AI. KQL hunts for Microsoft Sentinel and Defender XDR, Sigma, Suricata, Zeek, YARA, Splunk, Elastic, STIX 2.1. · GitHub
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
ThreatDefence
agentic-attack-detections Public Notifications You must be signed in to change notification settings
Star 632 ( 632 ) You must be signed in to star a repository
Detection rules and IOCs for AI agent attack frameworks: ARTEX, Cairn, PentAGI, HexStrike, CyberStrikeAI, Strix, Hermes Agent, Agent Zero, RedAmon, Xalgorix, Pentest Swarm AI. KQL hunts for Microsoft Sentinel and Defender XDR, Sigma, Suricata, Zeek, YARA, Splunk, Elastic, STIX 2.1.
threatdefence.com/article/artex-agentic-attack-framework-detection-hunting Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
132 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
78 Commits 78 Commits Folders and files
.github .github behaviors behaviors deploy deploy docs docs exposure exposure feeds feeds frameworks frameworks generated generated iocs iocs lab lab tests tests tools tools waf waf .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md zkg.meta zkg.meta View all files Repository files navigation
Open-source detection rules, hunting queries and indicators for security teams. They cover agentic attack frameworks: tools that hand an intrusion to AI agents, which plan, run real attack tools and keep going until they reach a goal.
Free to use under the MIT license. Fixes, false-positive reports and new indicators are welcome.
Framework folders. ARTEX and its AntiFox fork, Cairn, PentAGI, HexStrike, CyberStrikeAI, Strix, Hermes Agent, Agent Zero, RedAmon, Xalgorix and Pentest Swarm AI. Each folder is built from a review of the framework's source code and the public reporting on its use.
Shared behaviors. Rules for what these frameworks have in common, whatever the tool or the model behind it.
Indicators and feeds. One CSV of 580 sourced indicators, and the feeds generated from it.
Exposed consoles. A dated list of addresses where one of these consoles answers on the internet, in exposure/ , with a count per framework . It is a record of what is exposed, not a detection feed.
Formats: KQL for Microsoft Sentinel and Defender XDR, Sigma, Suricata, Zeek, YARA and STIX 2.1, plus WAF rules for Cloudflare, AWS WAF and Azure WAF. Splunk has one app for every framework. Elastic has one import with a rule for every framework.
With Suricata installed, run the ARTEX rules over the sample capture in this repository:
git clone https://github.com/ThreatDefence/agentic-attack-detections
cd agentic-attack-detections
mkdir -p out
suricata -r frameworks/artex/tests/artex-sample.pcap -S frameworks/artex/suricata/artex.rules -l out -k none
cat out/fast.log
You should see alerts such as:
[1:3800001:1] TD ARTEX user agent artex-<role>/1.0
[1:3800020:1] TD ARTEX placeholder probe path
[1:3800023:2] TD Unsigned alg none JWT in URI
The capture is test traffic written from the framework source code and published reports. It is not a recorded intrusion.
You have
Do this
Microsoft Sentinel
One deployment creates every analytics rule (disabled until you enable them), every hunting query and the AGENTIC_IOCs watchlist. Give it the workspace name. deploy/sentinel/ has one template per framework too. Run the hunts in each folder's hunting-order.md order
Defender XDR
Run the hunts from frameworks/*/sentinel/hunting/ in advanced hunting. They work where the tables exist
Suricata
Add the .rules files from frameworks/*/suricata/ and behaviors/suricata/ . HTTP rules need cleartext or decrypted traffic. DNS and TLS SNI rules work on encrypted traffic
Zeek
Install the behaviors package with zkg install https://github.com/ThreatDefence/agentic-attack-detections (or load behaviors/zeek/agentic-behaviors/ by path), set server_nets to your server ranges ( setup ), load the framework scripts frameworks/artex/zeek/artex/ , frameworks/cyberstrikeai/zeek/cyberstrikeai/ and frameworks/hexstrike/zeek/hexstrike/ , and feed frameworks/*/zeek/intel/*.intel to the Intel framework
YARA
Scan container layers, home folders, /tmp and web roots with frameworks/*/yara/
Any SIEM with Sigma
Convert frameworks/*/sigma/ and behaviors/sigma/ with sigma convert , or take the conversions in generated/
Splunk
Install agentic-attack-detections-splunk.tgz from the latest release (Apps > Install app from file). One app has the searches for every framework, the indicator lookups and a dashboard with a card per framework. Point the aad_indexes macro at your indexes. Setup
Elastic
Import agentic-attack-detections.ndjson through Security > Rules > Import, or post it to /api/detection_engine/rules/_import . One file carries every framework and the shared behaviors, all disabled: the Sigma rules as query rules, the correlations as ES|QL rules, the hand-written ARTEX rules, and an indicator match rule. For that one, load agentic-attack-indicators.ndjson with POST logs-ti_agentic_attacks-default/_bulk . Point the index patterns at your own data before you enable th
[truncated]
KQL hunts
Sentinel analytics
Sigma
Suricata
YARA
Zeek
Indicators
ARTEX
31
7
36
43
4
script, intel
247
Cairn
6
2
3
26
6
intel
73
PentAGI
5
2
3
10
3
intel
39
HexStrike
4
2
3
7
3
script, intel
22
CyberStrikeAI
5
2
7
11
4
script, intel
74
Strix
3
1
3
5
3
intel
49
Hermes Agent
4
2
3
7
4
intel
79
Agent Zero
2
1
3
1
intel
8
RedAmon
3
3
10
1
intel
19
Xalgorix
3
3
10
1
intel
24
Pentest Swarm AI
3
3
11
1
intel
15
Shared behaviors
9
2
6
18
package
Suricata counts are hand-written rules; each framework also has address rules generated from its indicators. Some indicators belong to more than one framework. Every framework is in the Splunk app . ARTEX also has Elastic rules and ES|QL queries, and an osquery pack.
Framework
What it is
Seen in
ARTEX
Go planner and worker agents with a web console. AntiFox is the same code renamed
Storm-0940, Storm-3358, VAPOR PANDA and criminal groups, against Microsoft 365, Okta, AWS, Azure, web applications and Active Directory
Cairn
Python engine that drives Claude Code, Codex and Pi coding agents in Kali containers
Online retailers, payment sites, Japanese web applications and WSUS servers
PentAGI
Go multi-agent platform that runs Kali tools in a container per engagement
Intrusions GTIG and Mandiant reported since August 2026, and a PaperCut exploitation operator's infrastructure
HexStrike
Python MCP server that gives any AI assistant about 150 security tools over an unauthenticated API
NetScaler exploitation discussion in 2025, and intrusions GTIG and Mandiant reported
CyberStrikeAI
Go operator console whose agents run about 90 tools, with its own implants, C2 listeners and webshell manager
FortiGate mass-exploitation infrastructure, an AWS intrusion, and education organizations in Africa and Latin America
Strix
Python AI pentest agent that runs teams of agents in a Kali sandbox, on any model
The retail card-theft campaign Gambit reported, where it found the flaws Cairn exploited
Hermes Agent
A legitimate general-purpose agent turned into an attack console: approvals off, a bundled jailbreak skill, red-team personas
The Gambit retail operator, the Carbonato Docker botnet, and Thailand's Ministry of Finance
Agent Zero
General-purpose agent in a Kali container with a built-in hacker profile and an open web UI
Threat-actor experimentation GTIG reported, with no published indicators
RedAmon
Self-hosted penetration testing platform: a web console, a recon pipeline and an agent that exploits from a Kali sandbox with Metasploit
No public report of use in an attack. Covered from its source code
Xalgorix
Go binary with a web dashboard whose agent tests web applications and APIs without approval steps and re-exploits each finding
No public report o
[truncated]
Hosts on the internet that serve the web console of each framework, found in internet-wide scan data with the fingerprints the rules here use: favicons, page titles, page text that survives a rename, and certificates.
An active host is one seen in the seven days to November 23, 2026. Hosts the scan data flags as honeypots are left out of both counts.
An exposed console is not proof of an attack. Researchers, students and security teams run these tools, and some hosts are decoys.
More than half of the pages titled CyberStrikeAI are flagged as honeypots, which is why that count is so much lower than a title search suggests. The flagged hosts stay in the list with honeypot set to yes .
One network that is not flagged holds more than a third of the remaining CyberStrikeAI hosts and of the PentAGI hosts. That looks like one mass deployment or a set of decoys, not that many operators ( #21 ).
Hermes Agent and Agent Zero are general-purpose agents, so their numbers count exposed agents, not attack consoles. Hermes is counted on its default port only.
Consoles served by host name only, or on a port that has not been scanned, are not counted.
The addresses, with the port and the fingerprint that found each one, are in exposure/consoles.csv . The counts of every snapshot are kept in exposure/history.csv , so the numbers can be followed over time.
Behaviors shared by agentic frameworks
behaviors/ holds rules that do not depend on any one framework.
The ARTEX folder also carries behavior hunts that apply to every framework: user agent family bursts (hunt 17), web and database services starting shells (hunt 11), servers calling webhook and tunnel services (hunt 12), scripts written into web folders (hunt 13), and Kerberoasting, DCSync and WinRM fan-out (hunt 16).
Path
Contents
iocs/indicators.csv
580 indicators. Each row has its source, confidence, a review date for addresses and the framework it belongs to. This file is the single source for every feed
feeds/
Generated from the CSV: a STIX 2.1 bundle, a Sentinel watchlist, and plain lists of addresses, domains, URLs, user agents and hashes. The plain lists carry medium and high confidence rows only. Low confidence rows, which include legitimate services such as AI model APIs, need their context, so they stay in the CSV, the watchlist and the STIX bundle. feeds/by-framework/ has the lists per framework
generated/
Every Sigma rule converted to Splunk SPL, Sentinel and Defender XDR KQL, and Elastic ES|QL and Lucene, plus elastic_ndjson/ where all of them are one Kibana import. SUMMARY.md shows which rule

[truncated]
