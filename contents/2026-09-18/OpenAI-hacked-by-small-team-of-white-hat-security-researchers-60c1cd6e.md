---
source: "https://venturebeat.com/security/openai-hacked-by-small-team-of-white-hat-security-researchers-using-anthropics-claude-opus-5"
hn_url: "https://news.ycombinator.com/item?id=49761145"
title: "OpenAI hacked by small team of white hat security researchers"
article_title: "OpenAI hacked by small team of white hat security researchers using Anthropic's Claude Opus 5 | VentureBeat"
image: "https://images.ctfassets.net/jdtwqhzvc2n1/013tvomW1y7WYPpCWQjDIb/67342f7455afc072ce063588d8ab874c/ChatGPT_Image_Sep_18__2026__12_15_50_AM.png?w=800&q=75"
author: "s3p"
captured_at: "2026-09-18T23:03:28Z"
capture_tool: "hn-digest"
hn_id: 49761145
score: 2
comments: 0
posted_at: "2026-09-18T22:32:43Z"
tags:
  - hacker-news
---

# OpenAI hacked by small team of white hat security researchers

- HN: [49761145](https://news.ycombinator.com/item?id=49761145)
- Source: [venturebeat.com](https://venturebeat.com/security/openai-hacked-by-small-team-of-white-hat-security-researchers-using-anthropics-claude-opus-5)
- Score: 2
- Comments: 0
- Posted: 2026-09-18T22:32:43Z

## Translation

Title: OpenAI hacked by small team of white hat security researchers
Article title: OpenAI hacked by small team of white hat security researchers using Anthropic's Claude Opus 5 | VentureBeat
Description: For enterprise teams, the larger lesson extends beyond one vulnerable image decoder. AI agents are increasingly being given credentials and connectors that span source code, communications and corporate data.

Article text:
OpenAI hacked by small team of white hat security researchers using Anthropic's Claude Opus 5 | VentureBeat The Independent
Voice on AI Orchestration
OpenAI hacked by small team of white hat security researchers using Anthropic's Claude Opus 5
Credit: VentureBeat made with OpenAI ChatGPT-Images-2.5
Security researchers say they used Anthropic’s newly released Claude Opus 5 to help turn an image-processing vulnerability into an exploit chain that compromised an OpenAI employee’s ChatGPT account and reached the company’s internal GitHub environment — an incident that illustrates how AI coding agents are changing the economics of sophisticated vulnerability exploitation.
The researchers, part of security startup Hacktron AI , disclosed the operation this week after reporting it to OpenAI and Discourse in July. The Wall Street Journal independently reported that the team gained access to an OpenAI employee’s ChatGPT account and had a path to read and propose changes to private OpenAI software.
Hacktron says the researchers stopped short of examining sensitive source code. Instead, they used the compromised employee account’s access to Codex to create a harmless pull request in OpenAI’s internal monorepo, demonstrating that the account compromise could extend beyond ChatGPT itself into connected developer infrastructure. The company promoted a video released tonight by YouTuber @LiveOverflow discussing their approach:
The incident matters for enterprises because it combines three increasingly important security boundaries: vulnerable third-party infrastructure, federated identity and AI agents connected to business systems.
Hacktron has framed the OpenAI incident as part of a much broader libheif research campaign, saying the same line of work extended to Slack, Meta, Zoom, Shopify, GitHub Enterprise and other widely used platforms. In the accompanying video, the researchers go further, alleging that the team also “hacked Slack, Meta, and many more.”
But its blog post and the supporting communications provided so far are focused overwhelmingly on OpenAI: they devote the detailed timeline, exploit chain, account-takeover mechanism, internal GitHub proof of access, disclosure process and bounty discussion to the OpenAI case, while the other companies are mentioned mainly as part of the wider HEIF Heist campaign rather than documented with the same level of technical detail or vendor confirmation.
From an image upload to an OpenAI account
The initial entry point was OpenAI’s community forum, community.openai.com, which runs on Discourse.
Hacktron found that HEIC and HEIF images uploaded to Discourse could be passed through ImageMagick and ultimately decoded using libheif . According to the researchers, the version of libheif present in the relevant Discourse environment contained a heap buffer overflow that could be developed into remote code execution.
Discourse has independently confirmed the image-processing vulnerability. Its July 28 security advisory says an upstream libheif vulnerability allowed remote code execution through image uploads and assigns the issue a CVSS score of 8.8. Discourse patched affected releases and added additional sandboxing around image processing.
Remote code execution on the forum was only the first stage.
Hacktron says it discovered a separate flaw in OpenAI’s single sign-on implementation that allowed the researchers to turn control of the forum environment into access to ChatGPT and Codex accounts belonging to users who had authenticated through the service. The researchers say those accounts included OpenAI employees.
That distinction is significant. Hacktron says the account escalation was not a Discourse vulnerability but an OpenAI identity issue, meaning the compromised forum acted as the foothold rather than the ultimate security boundary.
Connectors increased the blast radius
Once inside an affected ChatGPT or Codex account, the potential impact depended on what that account had connected.
Hacktron says affected accounts could have access to services including GitHub, Slack, Outlook, Gmail and Google Drive. In the case used for its proof of concept, an OpenAI employee’s Codex environment was connected to OpenAI’s GitHub organization.
Instead of inspecting the repository, the researchers instructed Codex to make a benign change and prepare a pull request in OpenAI’s internal openai/openai monorepo. Hacktron says the demonstration was intended to prove the level of access while minimizing exposure to proprietary information.
For enterprises deploying AI agents, that architecture creates a different risk model from a conventional chatbot. An AI account connected to source repositories, email, document stores and collaboration tools becomes an identity and authorization hub. Compromise of the AI account can therefore inherit the permissions of those downstream systems.
Claude helped operationalize the memory-corruption bug
The other notable part of the incident is how Hacktron says it developed the exploit.
The researchers initially used Claude Opus 4.8 to investigate the vulnerable libheif package and develop an exploit, but say the model struggled to make the attack reliable with address-space layout randomization enabled.
Anthropic released Claude Opus 5 on July 24. The company described the model as a major improvement for long-running agents and coding work.
Hacktron says that after switching to Opus 5, the model produced a working ARM64 exploit within hours and was subsequently used to adapt it to the x86-64 and jemalloc environment used by Discourse. The researchers say the full path from discovery to access to OpenAI’s repository environment took less than 72 hours.
That claim is particularly relevant for security teams because memory-corruption exploitation has traditionally required specialized expertise and substantial manual work. Hacktron argues that frontier coding agents can increasingly perform much of that incremental exploit-development work under human direction.
The broader libheif ecosystem reinforces the maintenance problem. The project says it published dozens of security advisories during 2026, while its September 6 v1.23.4 release included additional high-severity security fixes and urged users to upgrade.
Hacktron says it reported the findings through OpenAI’s Bugcrowd program on July 25 and that OpenAI confirmed later that day that its side of the vulnerability had been fixed.
OpenAI reportedly ultimately paid the researchers $6,500, according to Hacktron, while noting that testing against the Discourse-hosted community site itself was outside OpenAI’s bounty scope.
OpenAI has not, as far as could be verified, published its own detailed account of this particular incident. In an emailed statement to VentureBeat, an OpenAI spokesperson said: “ We thank the researchers for contacting us and sharing their findings. We narrowed the permissions on Community sign-in tokens and revoked affected tokens and sessions.”
The Wall Street Journal independently reported the breach and the researchers’ access to an OpenAI employee account, while Discourse has publicly confirmed and patched the underlying image-upload RCE.
For enterprise teams, the larger lesson extends beyond one vulnerable image decoder. AI agents are increasingly being given credentials and connectors that span source code, communications and corporate data. At the same time, the same class of models is making technically difficult exploit development faster and less expensive.
That combination puts more pressure on organizations to isolate untrusted file-processing pipelines, keep low-level dependencies aggressively patched, constrain federated identity trust and treat AI-agent credentials with the same scrutiny as privileged human accounts.
Do Not Sell or Share My Personal Information
Limit the Use Of My Sensitive Personal Information
© 2026 VentureBeat. All rights reserved.

## Original Extract

For enterprise teams, the larger lesson extends beyond one vulnerable image decoder. AI agents are increasingly being given credentials and connectors that span source code, communications and corporate data.

OpenAI hacked by small team of white hat security researchers using Anthropic's Claude Opus 5 | VentureBeat The Independent
Voice on AI Orchestration
OpenAI hacked by small team of white hat security researchers using Anthropic's Claude Opus 5
Credit: VentureBeat made with OpenAI ChatGPT-Images-2.5
Security researchers say they used Anthropic’s newly released Claude Opus 5 to help turn an image-processing vulnerability into an exploit chain that compromised an OpenAI employee’s ChatGPT account and reached the company’s internal GitHub environment — an incident that illustrates how AI coding agents are changing the economics of sophisticated vulnerability exploitation.
The researchers, part of security startup Hacktron AI , disclosed the operation this week after reporting it to OpenAI and Discourse in July. The Wall Street Journal independently reported that the team gained access to an OpenAI employee’s ChatGPT account and had a path to read and propose changes to private OpenAI software.
Hacktron says the researchers stopped short of examining sensitive source code. Instead, they used the compromised employee account’s access to Codex to create a harmless pull request in OpenAI’s internal monorepo, demonstrating that the account compromise could extend beyond ChatGPT itself into connected developer infrastructure. The company promoted a video released tonight by YouTuber @LiveOverflow discussing their approach:
The incident matters for enterprises because it combines three increasingly important security boundaries: vulnerable third-party infrastructure, federated identity and AI agents connected to business systems.
Hacktron has framed the OpenAI incident as part of a much broader libheif research campaign, saying the same line of work extended to Slack, Meta, Zoom, Shopify, GitHub Enterprise and other widely used platforms. In the accompanying video, the researchers go further, alleging that the team also “hacked Slack, Meta, and many more.”
But its blog post and the supporting communications provided so far are focused overwhelmingly on OpenAI: they devote the detailed timeline, exploit chain, account-takeover mechanism, internal GitHub proof of access, disclosure process and bounty discussion to the OpenAI case, while the other companies are mentioned mainly as part of the wider HEIF Heist campaign rather than documented with the same level of technical detail or vendor confirmation.
From an image upload to an OpenAI account
The initial entry point was OpenAI’s community forum, community.openai.com, which runs on Discourse.
Hacktron found that HEIC and HEIF images uploaded to Discourse could be passed through ImageMagick and ultimately decoded using libheif . According to the researchers, the version of libheif present in the relevant Discourse environment contained a heap buffer overflow that could be developed into remote code execution.
Discourse has independently confirmed the image-processing vulnerability. Its July 28 security advisory says an upstream libheif vulnerability allowed remote code execution through image uploads and assigns the issue a CVSS score of 8.8. Discourse patched affected releases and added additional sandboxing around image processing.
Remote code execution on the forum was only the first stage.
Hacktron says it discovered a separate flaw in OpenAI’s single sign-on implementation that allowed the researchers to turn control of the forum environment into access to ChatGPT and Codex accounts belonging to users who had authenticated through the service. The researchers say those accounts included OpenAI employees.
That distinction is significant. Hacktron says the account escalation was not a Discourse vulnerability but an OpenAI identity issue, meaning the compromised forum acted as the foothold rather than the ultimate security boundary.
Connectors increased the blast radius
Once inside an affected ChatGPT or Codex account, the potential impact depended on what that account had connected.
Hacktron says affected accounts could have access to services including GitHub, Slack, Outlook, Gmail and Google Drive. In the case used for its proof of concept, an OpenAI employee’s Codex environment was connected to OpenAI’s GitHub organization.
Instead of inspecting the repository, the researchers instructed Codex to make a benign change and prepare a pull request in OpenAI’s internal openai/openai monorepo. Hacktron says the demonstration was intended to prove the level of access while minimizing exposure to proprietary information.
For enterprises deploying AI agents, that architecture creates a different risk model from a conventional chatbot. An AI account connected to source repositories, email, document stores and collaboration tools becomes an identity and authorization hub. Compromise of the AI account can therefore inherit the permissions of those downstream systems.
Claude helped operationalize the memory-corruption bug
The other notable part of the incident is how Hacktron says it developed the exploit.
The researchers initially used Claude Opus 4.8 to investigate the vulnerable libheif package and develop an exploit, but say the model struggled to make the attack reliable with address-space layout randomization enabled.
Anthropic released Claude Opus 5 on July 24. The company described the model as a major improvement for long-running agents and coding work.
Hacktron says that after switching to Opus 5, the model produced a working ARM64 exploit within hours and was subsequently used to adapt it to the x86-64 and jemalloc environment used by Discourse. The researchers say the full path from discovery to access to OpenAI’s repository environment took less than 72 hours.
That claim is particularly relevant for security teams because memory-corruption exploitation has traditionally required specialized expertise and substantial manual work. Hacktron argues that frontier coding agents can increasingly perform much of that incremental exploit-development work under human direction.
The broader libheif ecosystem reinforces the maintenance problem. The project says it published dozens of security advisories during 2026, while its September 6 v1.23.4 release included additional high-severity security fixes and urged users to upgrade.
Hacktron says it reported the findings through OpenAI’s Bugcrowd program on July 25 and that OpenAI confirmed later that day that its side of the vulnerability had been fixed.
OpenAI reportedly ultimately paid the researchers $6,500, according to Hacktron, while noting that testing against the Discourse-hosted community site itself was outside OpenAI’s bounty scope.
OpenAI has not, as far as could be verified, published its own detailed account of this particular incident. In an emailed statement to VentureBeat, an OpenAI spokesperson said: “ We thank the researchers for contacting us and sharing their findings. We narrowed the permissions on Community sign-in tokens and revoked affected tokens and sessions.”
The Wall Street Journal independently reported the breach and the researchers’ access to an OpenAI employee account, while Discourse has publicly confirmed and patched the underlying image-upload RCE.
For enterprise teams, the larger lesson extends beyond one vulnerable image decoder. AI agents are increasingly being given credentials and connectors that span source code, communications and corporate data. At the same time, the same class of models is making technically difficult exploit development faster and less expensive.
That combination puts more pressure on organizations to isolate untrusted file-processing pipelines, keep low-level dependencies aggressively patched, constrain federated identity trust and treat AI-agent credentials with the same scrutiny as privileged human accounts.
Do Not Sell or Share My Personal Information
Limit the Use Of My Sensitive Personal Information
© 2026 VentureBeat. All rights reserved.
