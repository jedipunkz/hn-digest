---
source: "https://infisical.com/blog/ai-agent-security-hazmat-suits"
hn_url: "https://news.ycombinator.com/item?id=50198506"
title: "Securing AI Agents Requires Hazmat Suits"
article_title: "AI Agent Security Requires Hazmat Suits, Not Self-Defense"
image: "https://images.ctfassets.net/rzezkvk1rm65/4cuz60OBtpXbOhutyFSvRi/50d6f5700e5af4b8e5264511cea76493/ai-agent-security-hazmat-suits.png"
author: "FinnLobsien"
captured_at: "2026-10-22T18:21:01Z"
capture_tool: "hn-digest"
hn_id: 50198506
score: 2
comments: 0
posted_at: "2026-10-22T17:33:07Z"
tags:
  - hacker-news
---

# Securing AI Agents Requires Hazmat Suits

- HN: [50198506](https://news.ycombinator.com/item?id=50198506)
- Source: [infisical.com](https://infisical.com/blog/ai-agent-security-hazmat-suits)
- Score: 2
- Comments: 0
- Posted: 2026-10-22T17:33:07Z

## Translation

Title: Securing AI Agents Requires Hazmat Suits
Article title: AI Agent Security Requires Hazmat Suits, Not Self-Defense
Description: AI agents can't see the risks they face. Agent security needs architectural safeguards like sandboxes, credential brokers, and gateways, not behavioral rules.

Article text:
Platform Secrets Management Certificate Management Privileged Access Agent Vault Pricing Developers Docs Community Resources Blog Videos Webinars Case Studies Compare Infisical vs Hashicorp Vault Careers 27k Talk to an expert Get started for free Platform Pricing Developers Resources Careers 27k stars Talk to an expert Get started for free ↩ Blog AI Agent Security Requires Hazmat Suits, Not Self-Defense
In securing AI agents, most organizations operate in one of two failure modes:
Developers give AI agents access using .env files (or other plaintext credentials) and trust them to not wreak havoc.
The organization declares AI agents a risk and makes deploying/using them nearly impossible.
That’s because AI agent security is poorly understood. There’s no go-to recipe for it yet. We even lack a good mental model for how to secure agents (and what they are in the first place).
From self-defense classes to hazmat suits
If you planned travel to a dangerous place, you might prepare by researching how you could get robbed or scammed and take self-defense classes, buy protective clothing, fill pockets with decoy wallets, etc.
A lot of cybersecurity infrastructure works this way: we know what risks exist and mitigate them with tools and policies.
For instance, attackers often access sensitive data with a chain of compromised credentials (e.g. a compromised GitHub session token exposes a hardcoded AWS IAM token, which lets attackers mint production database credentials).
Cybersecurity measures directly counter this pattern:
We limit who can access production and only grant it via privileged access solutions.
We teach employees about keeping credentials safe and sensitize them to phishing attempts.
We adopt tools like secrets managers to know exactly who has what access and frequently rotate credentials.
These are like self-defense classes: they help us prevent or deal with attacks. But even a Navy SEAL with a bulletproof vest, a black belt in Jiu Jitsu, and a deadly arsenal can die from a deadly infection because humans can’t perceive these threats.
This is the security challenge we face with agents: running agents securely requires protecting them from risks they can’t see.
Humans solve this with hazmat suits: an area is contaminated with invisible dangers, so the only effective safety mechanism is to physically separate the body from the viruses, bacteria, or chemicals.
Even a Nobel prize-winning virologist can’t see a deadly virus, and AI agents can’t reliably recognize risky instructions and actions. That’s not a flaw, but a corollary of their architecture: agents are powered by probabilistic LLMs. This lets them iterate and accomplish tasks autonomously, but makes security assurances impossible without security infrastructure that constrains them.
Why LLM architecture makes agents inherently dangerous
Agents are basically a constant cycle of prompts and outputs. If you tell an agent “change my website’s button color to blue”, it might make LLM calls, produce reasoning tokens, and run a grep command. The agent’s runtime executes the command and the output becomes LLM context, which makes it part of the prompt. That prompt might result in commands to open files, change CSS properties, etc.
This becomes dangerous if an agent may take harmful actions to accomplish its task. There may be an issue in changing the website’s button color. Perhaps the agent troubleshoots by reading a GitHub issue on the Next.js repo, where a comment links to a skill the agent installs. Ostensibly, the skill teaches agents how to work with Next.js, but installing it also sends every credential to an attacker’s webhook.
Any human would deny permission to call that webhook, but most agents run unsupervised. People trust their Claude Code/Codex to skip permissions, grant it full disk access, and let it read any credential. After all, nothing bad has happened. Yet.
Vendors harden their products against attacks, but can’t guarantee security because attacks evolve and even one mistake can be catastrophic. The only safe assumption is that agents can and will take harmful actions because their architecture makes them impossible to rule out.
Behavioral vs. architectural safeguards
A behavioral safeguard is anything you can teach. If you instruct your engineering team to never share API keys in Slack or emails, they’ll get the point and won’t post them in Asana either, even if they have a .env they could theoretically publish on social media. This is the mental model of self-defense classes and protective clothing.
Behavioral safeguards don’t reliably work with AI agents. You might include “never open .env files” in the agent’s system prompt, but the agent may obtain plaintext credentials by running grep DATABASE_URL .env , so it technically never opened .env.
Agents need architectural safeguards. These are measures that make harmful actions impossible. An agent that has no credentials on its machine can’t find a workaround to access them. This is a hazmat suit.
Architectural safeguards aren’t new. They power many conventional security architecture principles:
Zero-trust: Never assume something can be trusted because it’s in your network, but constantly authenticate for access.
Least privilege: Give each identity only the exact access it needs to limit blast radius in case of compromise.
Isolation: Isolate workloads so they can’t cause damage beyond their own reach.
Strict architectural safeguards create friction. Least privilege can block someone’s work when they lack a credential they only need occasionally. Organizations struggle to admit it, but architectural safeguards are tradeoffs between efficient operations and security.
A paranoid organization could ensure security by issuing air-gapped company laptops that factory-reset at 12am and receive a daily batch of strictly necessary files. This would make human work impossible.
But AI agents won’t lambast you in workplace happiness surveys and, as we’ve seen, require architectural safeguards because a compromised AI agent is an infinitely patient attacker.
This is why most AI agent security infrastructure needs to provide architectural guardrails. The default stack isn’t settled, but a few pieces are emerging as consensus. All of the pieces below enable AI agents to fulfill their task while offering deterministic security assurances.
Sandboxes are isolated environments (VMs, gVisor, containers, etc.) agents run in. Organizations can configure security measures at the machine level. A behavioral guardrail forbidding an agent from reading certain files is vulnerable to workarounds, but a sandbox isn’t because the files never exist on the machine.
Organizations may have multiple types of sandboxes. A developer’s coding agent may launch each session in a somewhat permissive sandbox while a customer support agent may get a persistent sandbox that can only use the internet to call the organization’s documentation MCP.
Sandboxes aren’t novel AI technology. They’ve existed for a long time to isolate workloads, but have a natural fit with AI agents.
To realize their power, agents typically need access to HTTP services like APIs, CLIs, MCPs, and so on, which is often secured by a credential broker.
Credential brokers like Agent Vault secure how agents use third-party services by providing access to HTTP services without giving them credentials. Anthropic , Vercel , Cloudflare , and Google , all implement the concept: A proxy with a credential store intercepts the agent’s traffic and attaches credentials on egress, then forwards the response. The agent’s own environment contains no secret values, which pre-empts multiple AI agent risks:
Credential harvesting attacks only yield placeholder key-value pairs.
Agents can’t send credentials to their LLM provider (or anywhere else).
Credential brokering is increasingly popular because it’s nearly universal: it works anywhere that honors HTTPS_PROXY, which includes the vast majority of environments agents run in. It also requires no code changes to the agent because it makes requests to the same endpoints it would if no proxy is running (e.g. api.github.com), which makes it compatible with nearly any agent harness, tool, or homegrown agent infrastructure.
Beyond separating agents from secrets, agents also authenticate to the proxy, which makes a credential brokering solution a convenient place to define agent identity, which is missing in many organizations. As a proxy that intercepts all traffic, credential brokers also provide agent observability and audit logs.
Credential brokers share some overlap with gateways, which make permissions more granular, but change which endpoints an agent connects to.
Agent gateways let organizations configure detailed permissions. An API key may grant read/write access, but rarely offers more detailed permissions. If you want to authorize a customer support agent to process refunds of up to $100 up to 5 times a day, you can’t configure a Stripe API key for that.
These granular permissions exist within gateways. The agent authenticates to the gateway instead of the raw API (e.g. https://gateway.example.com/github/repos/acme/backend/ ) which allows or rejects the request.
Gateways are often limited to certain protocols (often MCP) and need to support tools to enable granular permissions, meaning they’ll never cover every service an organization uses.
The AI agent security stack is new, but its principles aren’t
Every organization’s security posture is (and has always been) enforced by a mix of behavioral and architectural safeguards, which map roughly to humans and machines:
Behavioral: Humans need to accomplish disparate tasks with different tools and approaches, so organizations need to equip them with the tools to do so. This means many humans could cause catastrophic damage, but we train and trust them not to.
Humans may make sloppy, catastrophic errors like hardcoding IAM admin credentials, but they could do worse if they were actively hostile.
Architectural: Traditional machine workloads have fixed execution paths. Every behavior they can perform is known, so behavioral guardrails don’t apply. Machine workloads are usually secured by architectural guardrails: Zero-trust architecture requires that every request an EC2 instance makes to S3 is signed and complies with IAM policy.
This Venn diagram has lots of overlap: Humans face architectural guardrails like role-based access control (RBAC) that deterministically denies what they shouldn’t access.
The point is that AI agents are machine workloads that can adapt their tools and approaches like humans, but can’t be trusted not to find ways to cause catastrophic damage.
This requires new types of security infrastructure like agent sandboxes, credential brokers, and gateways. These (among others) let organizations deterministically prevent AI agents’ harmful actions without stripping agents of their power.
This is why, at Infisical, we launched Agent Vault , our dedicated AI agent security product. It brokers credentials to any AI agent (using any off-the-shelf or homegrown harness) on any environment that honors HTTPS_PROXY, and to any HTTP service.
Start a free trial today or talk to an expert about your use case.
Technical Content Marketer, Infisical
From self-defense classes to hazmat suits
Why LLM architecture makes agents inherently dangerous
Behavioral vs. architectural safeguards Sandboxes
The AI agent security stack is new, but its principles aren’t
Secure humans, machines, and agents with Infisical today. Start a free trial or talk to an expert.
Book a demo READ NEXT October 20, 2026 Multi-Cloud Certificate Management: One Lifecycle Across AWS, Azure, GCP, and Windows
October 16, 2026 How Infisical Uses Greptile for AI Code Review
October 12, 2026 IBM deprecated the HashiCorp Boundary Community Edition. What’s next?
Starting with Infisical is simple, fast, and free.
Request a demo PRODUCT Secrets Management

## Original Extract

AI agents can't see the risks they face. Agent security needs architectural safeguards like sandboxes, credential brokers, and gateways, not behavioral rules.

Platform Secrets Management Certificate Management Privileged Access Agent Vault Pricing Developers Docs Community Resources Blog Videos Webinars Case Studies Compare Infisical vs Hashicorp Vault Careers 27k Talk to an expert Get started for free Platform Pricing Developers Resources Careers 27k stars Talk to an expert Get started for free ↩ Blog AI Agent Security Requires Hazmat Suits, Not Self-Defense
In securing AI agents, most organizations operate in one of two failure modes:
Developers give AI agents access using .env files (or other plaintext credentials) and trust them to not wreak havoc.
The organization declares AI agents a risk and makes deploying/using them nearly impossible.
That’s because AI agent security is poorly understood. There’s no go-to recipe for it yet. We even lack a good mental model for how to secure agents (and what they are in the first place).
From self-defense classes to hazmat suits
If you planned travel to a dangerous place, you might prepare by researching how you could get robbed or scammed and take self-defense classes, buy protective clothing, fill pockets with decoy wallets, etc.
A lot of cybersecurity infrastructure works this way: we know what risks exist and mitigate them with tools and policies.
For instance, attackers often access sensitive data with a chain of compromised credentials (e.g. a compromised GitHub session token exposes a hardcoded AWS IAM token, which lets attackers mint production database credentials).
Cybersecurity measures directly counter this pattern:
We limit who can access production and only grant it via privileged access solutions.
We teach employees about keeping credentials safe and sensitize them to phishing attempts.
We adopt tools like secrets managers to know exactly who has what access and frequently rotate credentials.
These are like self-defense classes: they help us prevent or deal with attacks. But even a Navy SEAL with a bulletproof vest, a black belt in Jiu Jitsu, and a deadly arsenal can die from a deadly infection because humans can’t perceive these threats.
This is the security challenge we face with agents: running agents securely requires protecting them from risks they can’t see.
Humans solve this with hazmat suits: an area is contaminated with invisible dangers, so the only effective safety mechanism is to physically separate the body from the viruses, bacteria, or chemicals.
Even a Nobel prize-winning virologist can’t see a deadly virus, and AI agents can’t reliably recognize risky instructions and actions. That’s not a flaw, but a corollary of their architecture: agents are powered by probabilistic LLMs. This lets them iterate and accomplish tasks autonomously, but makes security assurances impossible without security infrastructure that constrains them.
Why LLM architecture makes agents inherently dangerous
Agents are basically a constant cycle of prompts and outputs. If you tell an agent “change my website’s button color to blue”, it might make LLM calls, produce reasoning tokens, and run a grep command. The agent’s runtime executes the command and the output becomes LLM context, which makes it part of the prompt. That prompt might result in commands to open files, change CSS properties, etc.
This becomes dangerous if an agent may take harmful actions to accomplish its task. There may be an issue in changing the website’s button color. Perhaps the agent troubleshoots by reading a GitHub issue on the Next.js repo, where a comment links to a skill the agent installs. Ostensibly, the skill teaches agents how to work with Next.js, but installing it also sends every credential to an attacker’s webhook.
Any human would deny permission to call that webhook, but most agents run unsupervised. People trust their Claude Code/Codex to skip permissions, grant it full disk access, and let it read any credential. After all, nothing bad has happened. Yet.
Vendors harden their products against attacks, but can’t guarantee security because attacks evolve and even one mistake can be catastrophic. The only safe assumption is that agents can and will take harmful actions because their architecture makes them impossible to rule out.
Behavioral vs. architectural safeguards
A behavioral safeguard is anything you can teach. If you instruct your engineering team to never share API keys in Slack or emails, they’ll get the point and won’t post them in Asana either, even if they have a .env they could theoretically publish on social media. This is the mental model of self-defense classes and protective clothing.
Behavioral safeguards don’t reliably work with AI agents. You might include “never open .env files” in the agent’s system prompt, but the agent may obtain plaintext credentials by running grep DATABASE_URL .env , so it technically never opened .env.
Agents need architectural safeguards. These are measures that make harmful actions impossible. An agent that has no credentials on its machine can’t find a workaround to access them. This is a hazmat suit.
Architectural safeguards aren’t new. They power many conventional security architecture principles:
Zero-trust: Never assume something can be trusted because it’s in your network, but constantly authenticate for access.
Least privilege: Give each identity only the exact access it needs to limit blast radius in case of compromise.
Isolation: Isolate workloads so they can’t cause damage beyond their own reach.
Strict architectural safeguards create friction. Least privilege can block someone’s work when they lack a credential they only need occasionally. Organizations struggle to admit it, but architectural safeguards are tradeoffs between efficient operations and security.
A paranoid organization could ensure security by issuing air-gapped company laptops that factory-reset at 12am and receive a daily batch of strictly necessary files. This would make human work impossible.
But AI agents won’t lambast you in workplace happiness surveys and, as we’ve seen, require architectural safeguards because a compromised AI agent is an infinitely patient attacker.
This is why most AI agent security infrastructure needs to provide architectural guardrails. The default stack isn’t settled, but a few pieces are emerging as consensus. All of the pieces below enable AI agents to fulfill their task while offering deterministic security assurances.
Sandboxes are isolated environments (VMs, gVisor, containers, etc.) agents run in. Organizations can configure security measures at the machine level. A behavioral guardrail forbidding an agent from reading certain files is vulnerable to workarounds, but a sandbox isn’t because the files never exist on the machine.
Organizations may have multiple types of sandboxes. A developer’s coding agent may launch each session in a somewhat permissive sandbox while a customer support agent may get a persistent sandbox that can only use the internet to call the organization’s documentation MCP.
Sandboxes aren’t novel AI technology. They’ve existed for a long time to isolate workloads, but have a natural fit with AI agents.
To realize their power, agents typically need access to HTTP services like APIs, CLIs, MCPs, and so on, which is often secured by a credential broker.
Credential brokers like Agent Vault secure how agents use third-party services by providing access to HTTP services without giving them credentials. Anthropic , Vercel , Cloudflare , and Google , all implement the concept: A proxy with a credential store intercepts the agent’s traffic and attaches credentials on egress, then forwards the response. The agent’s own environment contains no secret values, which pre-empts multiple AI agent risks:
Credential harvesting attacks only yield placeholder key-value pairs.
Agents can’t send credentials to their LLM provider (or anywhere else).
Credential brokering is increasingly popular because it’s nearly universal: it works anywhere that honors HTTPS_PROXY, which includes the vast majority of environments agents run in. It also requires no code changes to the agent because it makes requests to the same endpoints it would if no proxy is running (e.g. api.github.com), which makes it compatible with nearly any agent harness, tool, or homegrown agent infrastructure.
Beyond separating agents from secrets, agents also authenticate to the proxy, which makes a credential brokering solution a convenient place to define agent identity, which is missing in many organizations. As a proxy that intercepts all traffic, credential brokers also provide agent observability and audit logs.
Credential brokers share some overlap with gateways, which make permissions more granular, but change which endpoints an agent connects to.
Agent gateways let organizations configure detailed permissions. An API key may grant read/write access, but rarely offers more detailed permissions. If you want to authorize a customer support agent to process refunds of up to $100 up to 5 times a day, you can’t configure a Stripe API key for that.
These granular permissions exist within gateways. The agent authenticates to the gateway instead of the raw API (e.g. https://gateway.example.com/github/repos/acme/backend/ ) which allows or rejects the request.
Gateways are often limited to certain protocols (often MCP) and need to support tools to enable granular permissions, meaning they’ll never cover every service an organization uses.
The AI agent security stack is new, but its principles aren’t
Every organization’s security posture is (and has always been) enforced by a mix of behavioral and architectural safeguards, which map roughly to humans and machines:
Behavioral: Humans need to accomplish disparate tasks with different tools and approaches, so organizations need to equip them with the tools to do so. This means many humans could cause catastrophic damage, but we train and trust them not to.
Humans may make sloppy, catastrophic errors like hardcoding IAM admin credentials, but they could do worse if they were actively hostile.
Architectural: Traditional machine workloads have fixed execution paths. Every behavior they can perform is known, so behavioral guardrails don’t apply. Machine workloads are usually secured by architectural guardrails: Zero-trust architecture requires that every request an EC2 instance makes to S3 is signed and complies with IAM policy.
This Venn diagram has lots of overlap: Humans face architectural guardrails like role-based access control (RBAC) that deterministically denies what they shouldn’t access.
The point is that AI agents are machine workloads that can adapt their tools and approaches like humans, but can’t be trusted not to find ways to cause catastrophic damage.
This requires new types of security infrastructure like agent sandboxes, credential brokers, and gateways. These (among others) let organizations deterministically prevent AI agents’ harmful actions without stripping agents of their power.
This is why, at Infisical, we launched Agent Vault , our dedicated AI agent security product. It brokers credentials to any AI agent (using any off-the-shelf or homegrown harness) on any environment that honors HTTPS_PROXY, and to any HTTP service.
Start a free trial today or talk to an expert about your use case.
Technical Content Marketer, Infisical
From self-defense classes to hazmat suits
Why LLM architecture makes agents inherently dangerous
Behavioral vs. architectural safeguards Sandboxes
The AI agent security stack is new, but its principles aren’t
Secure humans, machines, and agents with Infisical today. Start a free trial or talk to an expert.
Book a demo READ NEXT October 20, 2026 Multi-Cloud Certificate Management: One Lifecycle Across AWS, Azure, GCP, and Windows
October 16, 2026 How Infisical Uses Greptile for AI Code Review
October 12, 2026 IBM deprecated the HashiCorp Boundary Community Edition. What’s next?
Starting with Infisical is simple, fast, and free.
Request a demo PRODUCT Secrets Management
