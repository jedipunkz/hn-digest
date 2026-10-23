---
source: "https://nebusec.ai/research/pingcap-security-incident/"
hn_url: "https://news.ycombinator.com/item?id=50215283"
title: "Investigating the PingCAP Security Incident"
article_title: "Investigating the PingCAP Security Incident: Credential Exposure and Cloud Data Collection | NebuSec"
image: "https://nebusec.ai/_astro/social-preview.C-KyjYo7_Z1PWcoR.jpeg"
author: "jinqueeny"
captured_at: "2026-10-23T21:52:42Z"
capture_tool: "hn-digest"
hn_id: 50215283
score: 2
comments: 0
posted_at: "2026-10-23T21:02:27Z"
tags:
  - hacker-news
---

# Investigating the PingCAP Security Incident

- HN: [50215283](https://news.ycombinator.com/item?id=50215283)
- Source: [nebusec.ai](https://nebusec.ai/research/pingcap-security-incident/)
- Score: 2
- Comments: 0
- Posted: 2026-10-23T21:02:27Z

## Translation

Title: Investigating the PingCAP Security Incident
Article title: Investigating the PingCAP Security Incident: Credential Exposure and Cloud Data Collection | NebuSec
Description: Between September 9 and October 7, 2026, an unauthorized party used a compromised GitHub account and exposed cloud credentials to copy PingCAP source code and internal analytics, storage and support records. NebuSec led the independent investigation and found no evidence of command execution, persis
[truncated]

Article text:
Investigating the PingCAP Security Incident: Credential Exposure and Cloud Data Collection | NebuSec Skip to main content P r o d u c t P r o d u c t Code Security Continuous code security with verified findings and merge-ready fixes. Supply Chain Security Coming soon Secure the dependencies, build pipelines, and releases you rely on. Cloud Security Coming soon Continuous visibility and control across cloud infrastructure, identities, and configurations. S e c u r i t y A u d i t S e c u r i t y A u d i t R e s e a r c h R e s e a r c h Reports Deep dives into vulnerabilities, tools, techniques, and product updates. Vulnerability Index A live record of validated vulnerabilities found by NebuSec, responsibly disclosed or redacted when needed. A b o u t A b o u t Home Product Code Security Continuous code security with verified findings and merge-ready fixes.
Supply Chain Security Secure the dependencies, build pipelines, and releases you rely on. Coming soon
Cloud Security Continuous visibility and control across cloud infrastructure, identities, and configurations. Coming soon
Reports Deep dives into vulnerabilities, tools, techniques, and product updates.
Vulnerability Index A live record of validated vulnerabilities found by NebuSec, responsibly disclosed or redacted when needed.
Source-Code Review and Hardening
Indicators of Compromise (IOCs)
MITRE ATT&CK® and Artifact Mapping
Investigating the PingCAP Security Incident: Credential Exposure and Cloud Data Collection
Investigating the PingCAP Security Incident: Credential Exposure and Cloud Data Collection
Incident Response · October 2026
NebuSec October 21, 2026 7 min read
Source-Code Review and Hardening
Indicators of Compromise (IOCs)
MITRE ATT&CK® and Artifact Mapping
In October 2026, NebuSec responded to a security incident affecting PingCAP’s internal systems. Between September 9 and October 7, 2026, an unauthorized party used a compromised GitHub account and exposed cloud credentials to access PingCAP internal systems, obtain source code, and copy records from internal analytics, storage, and support environments. PingCAP engaged NebuSec to conduct an independent investigation, assess the scope of the exposed information, and support the containment and remediation efforts outlined in PingCAP’s October 19 security update .
Throughout our review of PingCAP’s production environments, NebuSec found no evidence of attacker command execution or persistence, and no evidence that data stored in customer TiDB databases was accessed. Instead, the intrusion was confined to internal source repositories and credential-based access to internal analytics, storage, and operational tools.
PingCAP began disabling the affected credentials on October 7, 2026—the same day it became aware of the activity. All critical credentials were invalidated that day. By October 14, all identified compromised credentials had been invalidated, broad cross-project permissions had been removed, and several vulnerabilities identified by NebuSec during a proactive review of the exposed source code had been patched and verified.
The threat actor gained initial entry to the PingCAP environment via a compromised GitHub account, exposing private repositories. Utilizing compromised keys and tokens harvested from those repositories and Terraform state, the attacker pivoted into Google Cloud, Jira, and the enterprise internal IM used for TiDB alerts.
Through Jira support tickets, the attacker harvested additional credentials previously shared by customers during routine support workflows. This included AWS keys, which ultimately allowed the intruder to reach several S3 buckets containing customers’ TiDB diagnostic data.
Over several weeks in September and early October, the attacker used stored keys for Google Cloud service accounts, eventually exporting tables from PingCAP’s BigQuery data warehouse on September 23–24 and downloading data from Cloud Storage. The data taken includes copies of the TiDB Cloud control-plane metadata, customer SQL editor history, CRM records and billing data, along with source code, Jira issues and the chat history of the enterprise internal IM.
Based on the evidence available, the attacker did not access live customer TiDB clusters.
Figure 1: Credential access and data collection (UTC).
NebuSec correlated cloud activity logs, the incident file inventory and reviewed data samples to identify the exposed data and sources summarized below.
Table 1: Exposed data mapped to its source systems.
The TiDB Cloud API-key values in the exported table were encrypted, and the encryption key was not exposed. Without that key, the encrypted values could not be decrypted to recover usable API credentials.
Customer-related information was copied from PingCAP’s internal analytics, storage, and support systems. That exposure is distinct from accessing or extracting rows directly from production customer TiDB databases.
NebuSec supported PingCAP in containing the intrusion, identifying exposed credentials and verifying remediation. The response combined access revocation with measures to keep legitimate business services running.
NebuSec confirmed PingCAP has disabled compromised service accounts, revoked exposed keys, replaced affected identities and tightened access permissions. All identified compromised credentials were invalidated and the identified unauthorized access paths were closed.
NebuSec also reviewed PingCAP’s exposed data and code repositories, scanned for additional exposed credentials, and worked with PingCAP on credential rotation and verification.
Source-Code Review and Hardening
Because the attacker obtained copies of internal source code, NebuSec conducted a focused security review of the affected services to look for flaws that an attacker with source-code access might try to exploit later. NebuSec reported four vulnerabilities with the potential to affect production data. PingCAP patched all four, and NebuSec retested each fix to verify remediation. A retrospective review of production logs found no evidence that any of these vulnerabilities had been exploited prior to being fixed.
Alongside these fixes, PingCAP expanded its logging and perimeter defenses, enabling additional Cloud Storage access logging in the affected project. PingCAP also confirmed that all 16 internet-facing services identified during the review were protected by a web application firewall (WAF). Broader credential rotation, cross-project IAM reviews and expanded secret scanning remain ongoing.
Our testing found that the TiDB Cloud API keys in the exported table are encrypted and could not be used directly. This does not apply to credentials recovered from support materials, including the AWS keys described above. We recommend that customers replace active TiDB Cloud management API keys now, following PingCAP’s published rotation guidance.
Customers who shared credentials with PingCAP support are advised to rotate them.
PingCAP is determining customer-specific impact and will contact affected customers directly. Customers should follow the published precautionary steps without waiting for individual notice.
For engineering and security teams, three practices can help reduce exposure and strengthen incident response:
Manage Credentials and Apply Least Privilege: Give service accounts, keys and tokens clear owners and only the access their workloads require. Regularly review permissions and credential lifetimes, and maintain reliable processes for rotation and revocation.
Review Code for Vulnerabilities and Exposed Secrets: Combine security reviews with automated scanning of source code, configuration and repository history. Address both software flaws and exposed credentials, and verify remediation to reduce opportunities for further access.
Improve Cloud Visibility: Maintain visibility into identity use, permission changes and data access across cloud services. Centralized audit records and alerts for unusual access or data movement help teams detect activity earlier and assess its scope.
Indicators of Compromise (IOCs)
The attacker used the following IP addresses to access PingCAP’s cloud environment between September 9 and October 7, 2026.
MITRE ATT&CK® and Artifact Mapping
The following techniques and cloud operations were observed during the investigation. Because these APIs and tools also have routine administrative uses, their investigative value comes from correlating them with anomalous source IPs, user agents, and access patterns.
All event times are UTC. Findings reflect the state of the investigation as of October 21, 2026. Internal resource identifiers and credential values have been omitted.
Previous post Beyond the Model: Benchmarking Codex Security, Claude Security, and Nebu svg]:px-3 rounded-lg group flex items-center justify-end size-full pointer-events-none opacity-50 cursor-not-allowed" aria-disabled="true"> Next post Newest post Building an AI defense system for everyone
Privacy Policy Terms of Service © 2026 NebuSec · All rights reserved

## Original Extract

Between September 9 and October 7, 2026, an unauthorized party used a compromised GitHub account and exposed cloud credentials to copy PingCAP source code and internal analytics, storage and support records. NebuSec led the independent investigation and found no evidence of command execution, persis
[truncated]

Investigating the PingCAP Security Incident: Credential Exposure and Cloud Data Collection | NebuSec Skip to main content P r o d u c t P r o d u c t Code Security Continuous code security with verified findings and merge-ready fixes. Supply Chain Security Coming soon Secure the dependencies, build pipelines, and releases you rely on. Cloud Security Coming soon Continuous visibility and control across cloud infrastructure, identities, and configurations. S e c u r i t y A u d i t S e c u r i t y A u d i t R e s e a r c h R e s e a r c h Reports Deep dives into vulnerabilities, tools, techniques, and product updates. Vulnerability Index A live record of validated vulnerabilities found by NebuSec, responsibly disclosed or redacted when needed. A b o u t A b o u t Home Product Code Security Continuous code security with verified findings and merge-ready fixes.
Supply Chain Security Secure the dependencies, build pipelines, and releases you rely on. Coming soon
Cloud Security Continuous visibility and control across cloud infrastructure, identities, and configurations. Coming soon
Reports Deep dives into vulnerabilities, tools, techniques, and product updates.
Vulnerability Index A live record of validated vulnerabilities found by NebuSec, responsibly disclosed or redacted when needed.
Source-Code Review and Hardening
Indicators of Compromise (IOCs)
MITRE ATT&CK® and Artifact Mapping
Investigating the PingCAP Security Incident: Credential Exposure and Cloud Data Collection
Investigating the PingCAP Security Incident: Credential Exposure and Cloud Data Collection
Incident Response · October 2026
NebuSec October 21, 2026 7 min read
Source-Code Review and Hardening
Indicators of Compromise (IOCs)
MITRE ATT&CK® and Artifact Mapping
In October 2026, NebuSec responded to a security incident affecting PingCAP’s internal systems. Between September 9 and October 7, 2026, an unauthorized party used a compromised GitHub account and exposed cloud credentials to access PingCAP internal systems, obtain source code, and copy records from internal analytics, storage, and support environments. PingCAP engaged NebuSec to conduct an independent investigation, assess the scope of the exposed information, and support the containment and remediation efforts outlined in PingCAP’s October 19 security update .
Throughout our review of PingCAP’s production environments, NebuSec found no evidence of attacker command execution or persistence, and no evidence that data stored in customer TiDB databases was accessed. Instead, the intrusion was confined to internal source repositories and credential-based access to internal analytics, storage, and operational tools.
PingCAP began disabling the affected credentials on October 7, 2026—the same day it became aware of the activity. All critical credentials were invalidated that day. By October 14, all identified compromised credentials had been invalidated, broad cross-project permissions had been removed, and several vulnerabilities identified by NebuSec during a proactive review of the exposed source code had been patched and verified.
The threat actor gained initial entry to the PingCAP environment via a compromised GitHub account, exposing private repositories. Utilizing compromised keys and tokens harvested from those repositories and Terraform state, the attacker pivoted into Google Cloud, Jira, and the enterprise internal IM used for TiDB alerts.
Through Jira support tickets, the attacker harvested additional credentials previously shared by customers during routine support workflows. This included AWS keys, which ultimately allowed the intruder to reach several S3 buckets containing customers’ TiDB diagnostic data.
Over several weeks in September and early October, the attacker used stored keys for Google Cloud service accounts, eventually exporting tables from PingCAP’s BigQuery data warehouse on September 23–24 and downloading data from Cloud Storage. The data taken includes copies of the TiDB Cloud control-plane metadata, customer SQL editor history, CRM records and billing data, along with source code, Jira issues and the chat history of the enterprise internal IM.
Based on the evidence available, the attacker did not access live customer TiDB clusters.
Figure 1: Credential access and data collection (UTC).
NebuSec correlated cloud activity logs, the incident file inventory and reviewed data samples to identify the exposed data and sources summarized below.
Table 1: Exposed data mapped to its source systems.
The TiDB Cloud API-key values in the exported table were encrypted, and the encryption key was not exposed. Without that key, the encrypted values could not be decrypted to recover usable API credentials.
Customer-related information was copied from PingCAP’s internal analytics, storage, and support systems. That exposure is distinct from accessing or extracting rows directly from production customer TiDB databases.
NebuSec supported PingCAP in containing the intrusion, identifying exposed credentials and verifying remediation. The response combined access revocation with measures to keep legitimate business services running.
NebuSec confirmed PingCAP has disabled compromised service accounts, revoked exposed keys, replaced affected identities and tightened access permissions. All identified compromised credentials were invalidated and the identified unauthorized access paths were closed.
NebuSec also reviewed PingCAP’s exposed data and code repositories, scanned for additional exposed credentials, and worked with PingCAP on credential rotation and verification.
Source-Code Review and Hardening
Because the attacker obtained copies of internal source code, NebuSec conducted a focused security review of the affected services to look for flaws that an attacker with source-code access might try to exploit later. NebuSec reported four vulnerabilities with the potential to affect production data. PingCAP patched all four, and NebuSec retested each fix to verify remediation. A retrospective review of production logs found no evidence that any of these vulnerabilities had been exploited prior to being fixed.
Alongside these fixes, PingCAP expanded its logging and perimeter defenses, enabling additional Cloud Storage access logging in the affected project. PingCAP also confirmed that all 16 internet-facing services identified during the review were protected by a web application firewall (WAF). Broader credential rotation, cross-project IAM reviews and expanded secret scanning remain ongoing.
Our testing found that the TiDB Cloud API keys in the exported table are encrypted and could not be used directly. This does not apply to credentials recovered from support materials, including the AWS keys described above. We recommend that customers replace active TiDB Cloud management API keys now, following PingCAP’s published rotation guidance.
Customers who shared credentials with PingCAP support are advised to rotate them.
PingCAP is determining customer-specific impact and will contact affected customers directly. Customers should follow the published precautionary steps without waiting for individual notice.
For engineering and security teams, three practices can help reduce exposure and strengthen incident response:
Manage Credentials and Apply Least Privilege: Give service accounts, keys and tokens clear owners and only the access their workloads require. Regularly review permissions and credential lifetimes, and maintain reliable processes for rotation and revocation.
Review Code for Vulnerabilities and Exposed Secrets: Combine security reviews with automated scanning of source code, configuration and repository history. Address both software flaws and exposed credentials, and verify remediation to reduce opportunities for further access.
Improve Cloud Visibility: Maintain visibility into identity use, permission changes and data access across cloud services. Centralized audit records and alerts for unusual access or data movement help teams detect activity earlier and assess its scope.
Indicators of Compromise (IOCs)
The attacker used the following IP addresses to access PingCAP’s cloud environment between September 9 and October 7, 2026.
MITRE ATT&CK® and Artifact Mapping
The following techniques and cloud operations were observed during the investigation. Because these APIs and tools also have routine administrative uses, their investigative value comes from correlating them with anomalous source IPs, user agents, and access patterns.
All event times are UTC. Findings reflect the state of the investigation as of October 21, 2026. Internal resource identifiers and credential values have been omitted.
Previous post Beyond the Model: Benchmarking Codex Security, Claude Security, and Nebu svg]:px-3 rounded-lg group flex items-center justify-end size-full pointer-events-none opacity-50 cursor-not-allowed" aria-disabled="true"> Next post Newest post Building an AI defense system for everyone
Privacy Policy Terms of Service © 2026 NebuSec · All rights reserved
