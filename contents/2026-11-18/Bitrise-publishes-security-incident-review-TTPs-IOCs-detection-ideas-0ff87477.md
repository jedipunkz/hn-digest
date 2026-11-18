---
source: "https://bitrise.io/blog/post/the-bitrise-security-incident-preliminary-review-and-recommended-steps"
hn_url: "https://news.ycombinator.com/item?id=50547869"
title: "Bitrise publishes security incident review: TTPs, IOCs, detection ideas"
article_title: "The Bitrise security incident: preliminary review and recommended steps - Bitrise Blog"
image: "https://cdn.prod.website-files.com/5db35de024bb988f5fb4e168/6afd799252067fce5198c4c4_bitrise-security-update-cover-im.jpg"
author: "flreln"
captured_at: "2026-11-18T16:13:47Z"
capture_tool: "hn-digest"
hn_id: 50547869
score: 3
comments: 0
posted_at: "2026-11-18T15:56:41Z"
tags:
  - hacker-news
---

# Bitrise publishes security incident review: TTPs, IOCs, detection ideas

- HN: [50547869](https://news.ycombinator.com/item?id=50547869)
- Source: [bitrise.io](https://bitrise.io/blog/post/the-bitrise-security-incident-preliminary-review-and-recommended-steps)
- Score: 3
- Comments: 0
- Posted: 2026-11-18T15:56:41Z

## Translation

Title: Bitrise publishes security incident review: TTPs, IOCs, detection ideas
Article title: The Bitrise security incident: preliminary review and recommended steps - Bitrise Blog
Description: After an extensive investigation over the last couple days, we want to share more context on the incident and the related impact.

Article text:
The Bitrise security incident: preliminary review and recommended steps - Bitrise Blog
Platform Products CI Automate builds, tests, and deployment
Build Cache Speed up your builds on any platform with remote caching
Release Management Distribute your builds internally and externally
CodePush Ship React Native updates in minutes
Build Hub Build infrastructure for GitHub Actions
Insights Identify exactly what's slowing your team down
Bitrise AI Get intelligent fixes from context-aware tools
Remote Dev Environments Code, build, and run AI agents in the cloud
Support Learn how to speed up your builds in ten minutes Watch demo
Solutions Company stage Enterprise Mobile DevOps for industry leaders
Scaling Flexible automation to fuel growth
Startup Fast mobile CI/CD for lean teams
Finance and Banking Secure, compliant DevOps solutions
Ecommerce and retail Fast, reliable mobile experiences
Nonprofits and charities Get your first year free
Build Hub vs GitHub-Hosted Runners
Support Learn how to speed up your builds in ten minutes Watch demo
Developers Learn & connect Documentation
Support Learn how to speed up your builds in ten minutes Watch demo
Shopify Sped up builds and reduced toolstack by 50%
Buzzfeed Shipping for mobile 85% faster with Bitrise
N26 60% reduction in regression testing time
Maven Clinic Cut testing time by 80% and 3x release speed
Substack 75% reduction in CI tooling
Back Market 15 minute build-to-release pipeline
Bitcoin.com Stable CI with x2 AI-driven velocity
Rightmove 83% faster builds across 5,000 releases yearly
React Native OTA updates: the complete guide for 2026 Documentation Support Learn how to speed up your builds in ten minutes Watch demo
Support Learn how to speed up your builds in ten minutes Watch demo Platform Products CI Automate builds, tests, and deployment
Build Cache Speed up your builds on any platform with remote caching
Release Management Distribute your builds internally and externally
CodePush Ship React Native updates in minutes
Build Hub Build infrastructure for GitHub Actions
Insights Identify exactly what's slowing your team down
Bitrise AI Get intelligent fixes from context-aware tools
Remote Dev Environments Code, build, and run AI agents in the cloud
Support Learn how to speed up your builds in ten minutes Watch demo
Solutions Company stage Enterprise Mobile DevOps for industry leaders
Scaling Flexible automation to fuel growth
Startup Fast mobile CI/CD for lean teams
Finance and Banking Secure, compliant DevOps solutions
Ecommerce and retail Fast, reliable mobile experiences
Nonprofits and charities Get your first year free
Build Hub vs GitHub-Hosted Runners
Support Learn how to speed up your builds in ten minutes Watch demo
Developers Learn & connect Documentation
Support Learn how to speed up your builds in ten minutes Watch demo
Shopify Sped up builds and reduced toolstack by 50%
Buzzfeed Shipping for mobile 85% faster with Bitrise
N26 60% reduction in regression testing time
Maven Clinic Cut testing time by 80% and 3x release speed
Substack 75% reduction in CI tooling
Back Market 15 minute build-to-release pipeline
Bitcoin.com Stable CI with x2 AI-driven velocity
Rightmove 83% faster builds across 5,000 releases yearly
React Native OTA updates: the complete guide for 2026 Documentation Support Learn how to speed up your builds in ten minutes Watch demo
Support Learn how to speed up your builds in ten minutes Watch demo
← Back to the blog Mobile DevOps The Bitrise security incident: preliminary review and recommended steps
Bitrise was started in 2014 to make building and releasing mobile apps easier and faster. More than 8,500 mobile teams trust Bitrise, and we want to keep that trust. That is why we are sharing a full and transparent review of the security incident.
On November 13, we detected suspicious activities with regards to a bot account that had access to our internal system running CI builds for build steps. Four hours later, we turned off the whole platform.
We have addressed the security risks, and Bitrise is safe to use. Most services are back to normal.
From Friday on, most of the company worked on nothing but the incident. Engineers from every product and infrastructure team worked with security, support and customer teams in shifts across Europe and the US, with a short call every few hours so everyone shared a single unified operational context. Support updated the status page almost every hour, around the clock.
We’re also publishing a technical advisory with what we observed during the attack: their IP addresses, user agents and sign-up patterns, how they tried to use decoy. The attacker’s indicators are also included as CSV and STIX files for threat response tools for import. Other teams can use them to check their own systems for the same attacker or patterns.
Before we go into the details, here’s what we recommend now. While we have no indication customer secrets were accessed in any way, out of an abundance of caution we suggest rotating the secrets you store in Bitrise. The attackers took some of Bitrise’s own secrets, not related to customers, from our CI workflows. Based on the log analysis of the attack, we have no indication of customer secret store access, nor attempts to access customer secret store.
Unfortunately, with the advance of AI accelerated attacks, we expect more such incidents to happen across the board. What we can do is to be prepared for fast response and rapid secret rotation. We are here to help our customers with that.
Rotate secrets, tokens and certificates stored in Bitrise, and the Git credentials connected to it.
You can find your secrets: Workflow Editor > Secrets and under Workspace settings > Shared resources.
To help identify your secret values to rotate, we released the Protected Secret Masked Review , which doesn't rotate anything for you. Click the eye icon in the field to reveal the last 4 characters. As you see below.
Authorize the Bitrise GitHub App again if your projects use it. On November 13 we replaced the private key of the Bitrise GitHub App and revoked all user authorizations for it. If enabling GitHub Checks or linking additional repositories fails with a 422 error, authorize the Bitrise GitHub App again on your Bitrise account.
Reconnect your GitHub OAuth connection in Bitrise. On November 13 we revoked every access token issued by the Bitrise GitHub OAuth app and rotated its client secret. This disconnected every Bitrise account that was connected to GitHub through OAuth. If your projects use the OAuth connection rather than the Bitrise GitHub App, reconnect your GitHub account in Bitrise to restore the connection .
Re-add your SSH key if your project uses SSH and a build can’t clone your repository. This only affects projects that use an SSH key to access the repository. When the OAuth tokens were revoked, GitHub automatically removed any SSH deploy keys that Bitrise had added to your repositories with them. Keys you added to GitHub yourself were not affected. If the Bitrise key is missing, configure it from your repository’s Settings > Deploy keys in GitHub, re-add the project’s SSH key in Bitrise after reconnecting your GitHub account.
Not related to the incident: As a best practice, review your public install pages. They allow download without signing in, by design. Turn off any you no longer need.
What was accessed and what wasn't
We know the first thing you want to know is whether your workspace was affected, so here’s what we found.
Most of what the attacker accessed was Bitrise's own: internal credentials, some of our repositories, and operational records like infrastructure configuration and build machine logs. Those records mention customers by name or ID in places, but contain no customer confidential data.
Files in 15 customer workspaces were accessed without authorization.The attacker downloaded:
1 code-signing certificate file, but could not access the password for it that is stored in the secret manager to open it,
13 diagnostic build logs from May 2026 containing short snippets of source or test code, a few lines at a time. One log contained a credential, and another one an email address.
All 15 customers were contacted immediately.
For every other workspace, our file-level in depth analysis found no evidence of the customer's secrets, code-signing files, build artifacts or source code in what the attacker accessed.
We are waiting for audit logs from GitHub to correlate with ours before we can move to close the investigation and enter the final report phase.
Timeline and how the attacker got in
Our investigation has linked the attacker to scanning and probing between October 25 and 27, 2026. During that period, they scanned more than 20 Bitrise hostnames and registered test accounts using disposable email addresses. This is a typical pattern we see all the time. Some of their tools described themselves as authorized penetration tests or security research, but those labels were supplied by the attacker without authorization. Practically, the attacker during their discovery blended in to business as usual customer initiated legitimate scanning activity.
During their discovery the attackers accessed one of our customers’ public applications that they use for their own penetration testing workflows, falsely establishing connection between the adversarial scan and legitimate pentest activity. We found this a sophisticated way of trying to decoy actual adversarial scanning and discovery activity.
On November 12 at 14:00 UTC, the attacker created a new account to run builds containing their own scripts against a public placeholder repository, running the same decoy patterns described above. Those scripts probed the build environment for credentials, and used said credentials to test those for cross-account listing vulnerabilities. All those requests were refused. This still seemed like a normal pentest pattern. The activity on this day happened between 2026-11-12 12:52 and 19:19. Then pause. Up until now, no out of the norm suspicious activity happened.
The day of the attack – The attack moved fast
The next day, on November 13 at 16:32 UTC, we detected suspicious activity.
The entry point was an internal Bitrise testing bot user account ( [email protected] ) used within a workspace designated for building and testing build steps. The attacker gained access by supplying the valid username and password combination; how those credentials were obtained remains under investigation. Originally configured to validate authentication with the Apple Developer Portal via AppleID, the bot account lacked two-factor authentication (2FA), a critical oversight that allowed initial entry. Once inside, the threat actor extracted additional credentials from the account to expand their access, executing an enumeration and privilege escalation scheme they had prepared over previous weeks. We are going to share more details in the final report, but wanted to share here early in the preliminary that we identified how the attackers got in.
During the early hours of the incident, we disabled all bot accounts, reviewed their credentials and 2FA settings so that this does not happen again. The initial attack vector was closed at that time.
We declared a SEV-1, our highest incident level at 17:42. As we did the live investigation initially, the attack seemed to be contained, but new results came in and facts were changing by the minute. We were correlating these realtime, but felt we were falling behind on the response curve to contain. When we saw evidence of unauthorized access in our GitHub repository, I made the call in the incident meeting to initiate “shutdown protocol”, turning off everything to contain the attack. At 19:26 UTC, we started shutting down our backend systems, database, secret access, all data centers, and GitHub access to stop any further possible exfiltration. We switched off the platform at 20:34 UTC.
Making that call in the middle of a fast-movin

[truncated]

## Original Extract

After an extensive investigation over the last couple days, we want to share more context on the incident and the related impact.

The Bitrise security incident: preliminary review and recommended steps - Bitrise Blog
Platform Products CI Automate builds, tests, and deployment
Build Cache Speed up your builds on any platform with remote caching
Release Management Distribute your builds internally and externally
CodePush Ship React Native updates in minutes
Build Hub Build infrastructure for GitHub Actions
Insights Identify exactly what's slowing your team down
Bitrise AI Get intelligent fixes from context-aware tools
Remote Dev Environments Code, build, and run AI agents in the cloud
Support Learn how to speed up your builds in ten minutes Watch demo
Solutions Company stage Enterprise Mobile DevOps for industry leaders
Scaling Flexible automation to fuel growth
Startup Fast mobile CI/CD for lean teams
Finance and Banking Secure, compliant DevOps solutions
Ecommerce and retail Fast, reliable mobile experiences
Nonprofits and charities Get your first year free
Build Hub vs GitHub-Hosted Runners
Support Learn how to speed up your builds in ten minutes Watch demo
Developers Learn & connect Documentation
Support Learn how to speed up your builds in ten minutes Watch demo
Shopify Sped up builds and reduced toolstack by 50%
Buzzfeed Shipping for mobile 85% faster with Bitrise
N26 60% reduction in regression testing time
Maven Clinic Cut testing time by 80% and 3x release speed
Substack 75% reduction in CI tooling
Back Market 15 minute build-to-release pipeline
Bitcoin.com Stable CI with x2 AI-driven velocity
Rightmove 83% faster builds across 5,000 releases yearly
React Native OTA updates: the complete guide for 2026 Documentation Support Learn how to speed up your builds in ten minutes Watch demo
Support Learn how to speed up your builds in ten minutes Watch demo Platform Products CI Automate builds, tests, and deployment
Build Cache Speed up your builds on any platform with remote caching
Release Management Distribute your builds internally and externally
CodePush Ship React Native updates in minutes
Build Hub Build infrastructure for GitHub Actions
Insights Identify exactly what's slowing your team down
Bitrise AI Get intelligent fixes from context-aware tools
Remote Dev Environments Code, build, and run AI agents in the cloud
Support Learn how to speed up your builds in ten minutes Watch demo
Solutions Company stage Enterprise Mobile DevOps for industry leaders
Scaling Flexible automation to fuel growth
Startup Fast mobile CI/CD for lean teams
Finance and Banking Secure, compliant DevOps solutions
Ecommerce and retail Fast, reliable mobile experiences
Nonprofits and charities Get your first year free
Build Hub vs GitHub-Hosted Runners
Support Learn how to speed up your builds in ten minutes Watch demo
Developers Learn & connect Documentation
Support Learn how to speed up your builds in ten minutes Watch demo
Shopify Sped up builds and reduced toolstack by 50%
Buzzfeed Shipping for mobile 85% faster with Bitrise
N26 60% reduction in regression testing time
Maven Clinic Cut testing time by 80% and 3x release speed
Substack 75% reduction in CI tooling
Back Market 15 minute build-to-release pipeline
Bitcoin.com Stable CI with x2 AI-driven velocity
Rightmove 83% faster builds across 5,000 releases yearly
React Native OTA updates: the complete guide for 2026 Documentation Support Learn how to speed up your builds in ten minutes Watch demo
Support Learn how to speed up your builds in ten minutes Watch demo
← Back to the blog Mobile DevOps The Bitrise security incident: preliminary review and recommended steps
Bitrise was started in 2014 to make building and releasing mobile apps easier and faster. More than 8,500 mobile teams trust Bitrise, and we want to keep that trust. That is why we are sharing a full and transparent review of the security incident.
On November 13, we detected suspicious activities with regards to a bot account that had access to our internal system running CI builds for build steps. Four hours later, we turned off the whole platform.
We have addressed the security risks, and Bitrise is safe to use. Most services are back to normal.
From Friday on, most of the company worked on nothing but the incident. Engineers from every product and infrastructure team worked with security, support and customer teams in shifts across Europe and the US, with a short call every few hours so everyone shared a single unified operational context. Support updated the status page almost every hour, around the clock.
We’re also publishing a technical advisory with what we observed during the attack: their IP addresses, user agents and sign-up patterns, how they tried to use decoy. The attacker’s indicators are also included as CSV and STIX files for threat response tools for import. Other teams can use them to check their own systems for the same attacker or patterns.
Before we go into the details, here’s what we recommend now. While we have no indication customer secrets were accessed in any way, out of an abundance of caution we suggest rotating the secrets you store in Bitrise. The attackers took some of Bitrise’s own secrets, not related to customers, from our CI workflows. Based on the log analysis of the attack, we have no indication of customer secret store access, nor attempts to access customer secret store.
Unfortunately, with the advance of AI accelerated attacks, we expect more such incidents to happen across the board. What we can do is to be prepared for fast response and rapid secret rotation. We are here to help our customers with that.
Rotate secrets, tokens and certificates stored in Bitrise, and the Git credentials connected to it.
You can find your secrets: Workflow Editor > Secrets and under Workspace settings > Shared resources.
To help identify your secret values to rotate, we released the Protected Secret Masked Review , which doesn't rotate anything for you. Click the eye icon in the field to reveal the last 4 characters. As you see below.
Authorize the Bitrise GitHub App again if your projects use it. On November 13 we replaced the private key of the Bitrise GitHub App and revoked all user authorizations for it. If enabling GitHub Checks or linking additional repositories fails with a 422 error, authorize the Bitrise GitHub App again on your Bitrise account.
Reconnect your GitHub OAuth connection in Bitrise. On November 13 we revoked every access token issued by the Bitrise GitHub OAuth app and rotated its client secret. This disconnected every Bitrise account that was connected to GitHub through OAuth. If your projects use the OAuth connection rather than the Bitrise GitHub App, reconnect your GitHub account in Bitrise to restore the connection .
Re-add your SSH key if your project uses SSH and a build can’t clone your repository. This only affects projects that use an SSH key to access the repository. When the OAuth tokens were revoked, GitHub automatically removed any SSH deploy keys that Bitrise had added to your repositories with them. Keys you added to GitHub yourself were not affected. If the Bitrise key is missing, configure it from your repository’s Settings > Deploy keys in GitHub, re-add the project’s SSH key in Bitrise after reconnecting your GitHub account.
Not related to the incident: As a best practice, review your public install pages. They allow download without signing in, by design. Turn off any you no longer need.
What was accessed and what wasn't
We know the first thing you want to know is whether your workspace was affected, so here’s what we found.
Most of what the attacker accessed was Bitrise's own: internal credentials, some of our repositories, and operational records like infrastructure configuration and build machine logs. Those records mention customers by name or ID in places, but contain no customer confidential data.
Files in 15 customer workspaces were accessed without authorization.The attacker downloaded:
1 code-signing certificate file, but could not access the password for it that is stored in the secret manager to open it,
13 diagnostic build logs from May 2026 containing short snippets of source or test code, a few lines at a time. One log contained a credential, and another one an email address.
All 15 customers were contacted immediately.
For every other workspace, our file-level in depth analysis found no evidence of the customer's secrets, code-signing files, build artifacts or source code in what the attacker accessed.
We are waiting for audit logs from GitHub to correlate with ours before we can move to close the investigation and enter the final report phase.
Timeline and how the attacker got in
Our investigation has linked the attacker to scanning and probing between October 25 and 27, 2026. During that period, they scanned more than 20 Bitrise hostnames and registered test accounts using disposable email addresses. This is a typical pattern we see all the time. Some of their tools described themselves as authorized penetration tests or security research, but those labels were supplied by the attacker without authorization. Practically, the attacker during their discovery blended in to business as usual customer initiated legitimate scanning activity.
During their discovery the attackers accessed one of our customers’ public applications that they use for their own penetration testing workflows, falsely establishing connection between the adversarial scan and legitimate pentest activity. We found this a sophisticated way of trying to decoy actual adversarial scanning and discovery activity.
On November 12 at 14:00 UTC, the attacker created a new account to run builds containing their own scripts against a public placeholder repository, running the same decoy patterns described above. Those scripts probed the build environment for credentials, and used said credentials to test those for cross-account listing vulnerabilities. All those requests were refused. This still seemed like a normal pentest pattern. The activity on this day happened between 2026-11-12 12:52 and 19:19. Then pause. Up until now, no out of the norm suspicious activity happened.
The day of the attack – The attack moved fast
The next day, on November 13 at 16:32 UTC, we detected suspicious activity.
The entry point was an internal Bitrise testing bot user account ( [email protected] ) used within a workspace designated for building and testing build steps. The attacker gained access by supplying the valid username and password combination; how those credentials were obtained remains under investigation. Originally configured to validate authentication with the Apple Developer Portal via AppleID, the bot account lacked two-factor authentication (2FA), a critical oversight that allowed initial entry. Once inside, the threat actor extracted additional credentials from the account to expand their access, executing an enumeration and privilege escalation scheme they had prepared over previous weeks. We are going to share more details in the final report, but wanted to share here early in the preliminary that we identified how the attackers got in.
During the early hours of the incident, we disabled all bot accounts, reviewed their credentials and 2FA settings so that this does not happen again. The initial attack vector was closed at that time.
We declared a SEV-1, our highest incident level at 17:42. As we did the live investigation initially, the attack seemed to be contained, but new results came in and facts were changing by the minute. We were correlating these realtime, but felt we were falling behind on the response curve to contain. When we saw evidence of unauthorized access in our GitHub repository, I made the call in the incident meeting to initiate “shutdown protocol”, turning off everything to contain the attack. At 19:26 UTC, we started shutting down our backend systems, database, secret access, all data centers, and GitHub access to stop any further possible exfiltration. We switched off the platform at 20:34 UTC.
Making that call in the middle of a fast-movin

[truncated]
