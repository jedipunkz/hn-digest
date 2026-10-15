---
source: "https://www.stackone.com/blog/tool-policies-launch/"
hn_url: "https://news.ycombinator.com/item?id=50107215"
title: "Tool Policies: one policy engine for every tool your AI agents call"
article_title: "Tool Policies: AI Agent Permissions Down to the Field | StackOne"
image: "https://www.stackone.com/og/blog/tool-policies-launch.png"
author: "shawneechasefos"
captured_at: "2026-10-15T18:02:38Z"
capture_tool: "hn-digest"
hn_id: 50107215
score: 1
comments: 1
posted_at: "2026-10-15T17:38:23Z"
tags:
  - hacker-news
---

# Tool Policies: one policy engine for every tool your AI agents call

- HN: [50107215](https://news.ycombinator.com/item?id=50107215)
- Source: [www.stackone.com](https://www.stackone.com/blog/tool-policies-launch/)
- Score: 1
- Comments: 1
- Posted: 2026-10-15T17:38:23Z

## Translation

Title: Tool Policies: one policy engine for every tool your AI agents call
Article title: Tool Policies: AI Agent Permissions Down to the Field | StackOne
Description: StackOne Tool Policies control what AI agents can do in your business systems, down to the field, with PII redaction and prompt injection defense built in.

Article text:
Tool Policies: AI Agent Permissions Down to the Field | StackOne -->
Skip to main content
Announcing Tool Gateway MCP: the universal MCP Read the announcement Connectors
Platform
Connect Pre-Built Connectors 550+ Connectors and 34,000+ Actions
Build/Extend your connector via code in minutes
OAuth, API keys, and token refresh – all handled
One universal MCP for every AI assistant you run
Trigger agentic workflows from app events
Connect to everything your agents need, without writing the infrastructure to get there.
Block prompt injections before they reach your agent
Remove personal data from tool calls and logs
Govern how teams use AI agents
Field-level control over reads and writes
Scope actions and share account access
SSO and directory sync for team access
Action logs, changes and sign-in activity
Regional cloud, dedicated or self-hosted
Enterprise-grade controls to ensure your agent always meets security & compliance standards.
Find the right action, skip the schema bloat
Filter and aggregate synced data for faster answers
Actions your agents get right the first time
Maximise agent efficiency – faster responses, increased accuracy, fewer tokens used.
Insights on agentic infrastructure
Talks, meetups, livestreams & launches
What’s new in agentic infrastructure
Compliance, certifications & trust
Connect to everything your agents need, without writing the infrastructure to get there.
Enterprise-grade controls to ensure your agent always meets security & compliance standards.
Maximise agent efficiency – faster responses, increased accuracy, fewer tokens used.
Tool Policies: one policy engine for every tool your AI agents call
Why do AI agent rollouts stall at the security review?
What are StackOne Tool Policies?
How are Tool Policies different from the controls built into AI platforms?
Semantic rules: one question instead of a hundred field names
How do you control which finance actions and fields AI agents can access?
Can AI agents analyse patient data under HIPAA?
How do you hide salary and personal data from AI agents in HR systems?
Why is StackOne the most complete way to secure AI agent tool calls?
Today we are launching StackOne Tool Policies , one policy engine for every tool your AI agents call. Tool Policies set what each person’s agents can do, see and send, down to the field, across every connector, MCP client and agent platform.
Why do AI agent rollouts stall at the security review?
A finance team asks IT to let its agent pay supplier bills. Security looks at what the payments connector can do, which includes paying any amount to any supplier and deleting the record afterwards, and the approval is denied. In the IT and security teams we talk to, this is where most agent rollouts stop. Security teams block what they cannot govern, and it is one of the three reasons AI rollouts stall that I wrote about last month.
The risk is real. Gartner predicts that by 2028, one in four enterprise breaches will be traced back to AI agents misused by outside attackers or by people inside the company. In McKinsey’s 2026 survey of about 500 organizations, nearly two-thirds named security and risk as the top barrier to scaling agentic AI.
Buyers tell us the same thing in their own words. An IT Director leading an internal AI rollout said they were “only tackling read only use cases” because “we’re not sure what guardrails we need.” An executive who already reads NetSuite through Claude will not let it post a journal entry until there are “more checks and balances, and security”.
The controls that ship with AI platforms do not get them there. Claude, ChatGPT, Microsoft Copilot, Gemini Enterprise and Glean all let an admin decide which connectors and tools an agent can use, and most can limit that by user or group. Those controls stop at the tool. As Guillaume, our CTO, wrote in August , the app switch “has no setting for ‘everything except salary’”, and none of the platforms can let an agent pay a $250 bill while refusing a $4,800 one.
So where do you go from here? Gartner’s advice is to build these controls in before a breach, and Tool Policies let you do that today.
What are StackOne Tool Policies?
Tool Policies address the answers security needs before it signs off: who can do what, to what data, through which agent.
Start with one rule, such as keeping customer PII out of the systems it should not reach. Scope the next one to a group, like contractors or everyone outside Finance. Then go down to a single field, or to a plain-language question like “Is this compensation?”
Tool Policies sit on top of the permissions in your systems: for an employee on their own login and for a background agent on a service account alike. Nobody changes a role in the HR, CRM or finance system.
Connectors open up the use cases. Doing it securely, with the right guardrails and governance in place, is what makes it possible to roll agents out across the whole business.
How are Tool Policies different from the controls built into AI platforms?
So do you still need the controls built into your AI platform? Yes. Keep them switched on: they decide which systems and tools an agent can reach. Tool Policies standardise the governance and guardrails across any platform you deploy, and decide how agents behave once they’re connected. Some platforms are adding argument checks in beta, but each one governs only its own agents, so a company running Copilot, Claude and Glean side by side writes the same rule three times. With Tool Policies, it writes the rule once.
“IT and functional leaders couldn’t safely enable more access, because of what ungoverned access would imply. Tool Policies is the exhaustive answer to that problem.”
Guillaume Lebedel, CTO and Co-founder, StackOne
Who it applies to: everyone, or users and groups synced from your identity provider, with named exemptions.
Where it applies: the whole project, a connector category, one connector or one linked account. A category rule also covers connectors you add later, and a name pattern like *delete* catches every matching tool.
What it does: five rule types. PII redaction strips personal data or blocks the call. Tool rules block tools. Value rules limit what a tool accepts, such as payments between $25 and $500. Field rules hide or protect fields by path. Semantic rules ask a plain-language question.
StackOne checks the input before it calls the business system, and the response before the agent sees it. Rules compile to Cedar, the open-source policy language built by AWS, where if two rules disagree the deny always wins. PII redaction runs on StackOne’s own state of the art model, which finds email addresses, phone numbers, SSNs and card numbers anywhere in a payload.
Build a policy in the no-code builder, have the StackOne Assistant draft it, or write it as Cedar or JSON. Every rule can run in Monitor first, which records who it would affect without blocking anything, and every decision is logged. Guillaume describes the rollout sequence we recommend in Monitor first, then enforce .
Semantic rules: one question instead of a hundred field names
Say you want to keep pay away from every agent. The obvious way is a field rule: you name the field to hide, such as salary. But every system names it differently. One HR system calls it salary, another base_pay, a third pay_rate, so you’d need a rule for each one. And a Slack message about someone’s raise has no field to name at all.
A semantic rule asks a question instead, such as “Is this compensation?” or “Does this message contain unreleased financial results?”, and a decision model answers it on every call. One rule covers the HR systems you connect today and the ones you add next year. When the model cannot decide, the field stays masked.
How do you control which finance actions and fields AI agents can access?
If a finance team wants one agent for analytics and forecasting and another to pay supplier bills, the security review is going to turn up four problems. The payments agent could pay any amount. Anyone with access to supplier records could change where the money goes. Contractors could delete records. And the analytics agent would see every supplier’s bank details.
Tool Policies fix each one with a rule for the group it applies to, using the groups already in your identity provider:
Accounts Payable can pay through an agent, between $25 and $500 per payment, and can prepare invoices but not authorize them. A $250 bill goes through. A $4,800 payment is stopped before the finance system is called.
Procurement can read supplier records, but a change to a bank account number is refused. A request to change a supplier’s bank details in SAP is blocked and logged.
Contractors cannot run any tool matching *delete* , on any connector.
Data teams can read accounting records with bank details masked.
See it end to end in our vendor payment scheduling and AP invoice processing use cases.
Can AI agents analyse patient data under HIPAA?
Take a healthcare provider with several clinics. Its operations team wants an agent to read patient messages and support tickets and report why appointments get rescheduled, or which clinics get the most billing questions. The analysis needs what each message says, not who sent it. HIPAA’s minimum necessary standard asks for exactly that: use only the patient data a task needs.
So the identifiers stay out of the agent’s context window and out of your logs:
A PII redaction rule on the patient support and ticketing connectors replaces email addresses, phone numbers, SSNs and card numbers with a tag that names the type, such as <POLICY_REDACTED:EMAIL> , wherever they appear in the text.
Field rules mask the fields that name the patient directly, such as name, date of birth and home address.
Advanced Logs redaction removes personal data from stored request and response bodies.
The agent reads hundreds of messages and reports the top reasons for rescheduling without seeing who wrote any of them. Every redaction is recorded in the action log for your compliance team. A Business Associate Agreement is available on Enterprise plans where PHI is in scope. I covered the wider picture in the reality of connecting AI agents to healthcare systems .
How do you hide salary and personal data from AI agents in HR systems?
People Ops coordinators use agents for onboarding, leave and employee questions. They need job titles, start dates and managers, and they do not need salary, bank details or Social Security numbers. For the People Operations group:
A semantic rule asks “Is this compensation?” on every HR read and write tool, so pay is hidden whatever each HR system calls it, and a raise attempted through the agent is refused.
Field rules mask Social Security number, personal email, home address and bank details on employee records.
PII redaction catches an SSN or phone number typed into a case note, where no field rule can reach.
The Total Rewards team is exempt, so their agents still see pay.
A coordinator asks for a new starter’s details and gets the title, start date and manager, with pay and identifiers removed before the agent sees the record. The coordinator may see more in the HR system itself: the policy governs the agent separately from the person, the gap Guillaume described in say yes to the Copilot, not to the salary field .
See how this fits an onboarding agent in our employee onboarding use case .
Why is StackOne the most complete way to secure AI agent tool calls?
Every action an agent takes in your business is a tool call. Updating a record, sending a message, paying a bill: each one goes through a tool, whether the agent runs in Claude , ChatGPT , Copilot , Gemini Enterprise , Glean or your own code. That makes the tool call the one point every agent passes through, and the place where security does the most good.
Securing a tool call down to the field means owning the tool. A gateway that passes calls through to other people’s MCP servers can allow or block a whole serve

[truncated]

## Original Extract

StackOne Tool Policies control what AI agents can do in your business systems, down to the field, with PII redaction and prompt injection defense built in.

Tool Policies: AI Agent Permissions Down to the Field | StackOne -->
Skip to main content
Announcing Tool Gateway MCP: the universal MCP Read the announcement Connectors
Platform
Connect Pre-Built Connectors 550+ Connectors and 34,000+ Actions
Build/Extend your connector via code in minutes
OAuth, API keys, and token refresh – all handled
One universal MCP for every AI assistant you run
Trigger agentic workflows from app events
Connect to everything your agents need, without writing the infrastructure to get there.
Block prompt injections before they reach your agent
Remove personal data from tool calls and logs
Govern how teams use AI agents
Field-level control over reads and writes
Scope actions and share account access
SSO and directory sync for team access
Action logs, changes and sign-in activity
Regional cloud, dedicated or self-hosted
Enterprise-grade controls to ensure your agent always meets security & compliance standards.
Find the right action, skip the schema bloat
Filter and aggregate synced data for faster answers
Actions your agents get right the first time
Maximise agent efficiency – faster responses, increased accuracy, fewer tokens used.
Insights on agentic infrastructure
Talks, meetups, livestreams & launches
What’s new in agentic infrastructure
Compliance, certifications & trust
Connect to everything your agents need, without writing the infrastructure to get there.
Enterprise-grade controls to ensure your agent always meets security & compliance standards.
Maximise agent efficiency – faster responses, increased accuracy, fewer tokens used.
Tool Policies: one policy engine for every tool your AI agents call
Why do AI agent rollouts stall at the security review?
What are StackOne Tool Policies?
How are Tool Policies different from the controls built into AI platforms?
Semantic rules: one question instead of a hundred field names
How do you control which finance actions and fields AI agents can access?
Can AI agents analyse patient data under HIPAA?
How do you hide salary and personal data from AI agents in HR systems?
Why is StackOne the most complete way to secure AI agent tool calls?
Today we are launching StackOne Tool Policies , one policy engine for every tool your AI agents call. Tool Policies set what each person’s agents can do, see and send, down to the field, across every connector, MCP client and agent platform.
Why do AI agent rollouts stall at the security review?
A finance team asks IT to let its agent pay supplier bills. Security looks at what the payments connector can do, which includes paying any amount to any supplier and deleting the record afterwards, and the approval is denied. In the IT and security teams we talk to, this is where most agent rollouts stop. Security teams block what they cannot govern, and it is one of the three reasons AI rollouts stall that I wrote about last month.
The risk is real. Gartner predicts that by 2028, one in four enterprise breaches will be traced back to AI agents misused by outside attackers or by people inside the company. In McKinsey’s 2026 survey of about 500 organizations, nearly two-thirds named security and risk as the top barrier to scaling agentic AI.
Buyers tell us the same thing in their own words. An IT Director leading an internal AI rollout said they were “only tackling read only use cases” because “we’re not sure what guardrails we need.” An executive who already reads NetSuite through Claude will not let it post a journal entry until there are “more checks and balances, and security”.
The controls that ship with AI platforms do not get them there. Claude, ChatGPT, Microsoft Copilot, Gemini Enterprise and Glean all let an admin decide which connectors and tools an agent can use, and most can limit that by user or group. Those controls stop at the tool. As Guillaume, our CTO, wrote in August , the app switch “has no setting for ‘everything except salary’”, and none of the platforms can let an agent pay a $250 bill while refusing a $4,800 one.
So where do you go from here? Gartner’s advice is to build these controls in before a breach, and Tool Policies let you do that today.
What are StackOne Tool Policies?
Tool Policies address the answers security needs before it signs off: who can do what, to what data, through which agent.
Start with one rule, such as keeping customer PII out of the systems it should not reach. Scope the next one to a group, like contractors or everyone outside Finance. Then go down to a single field, or to a plain-language question like “Is this compensation?”
Tool Policies sit on top of the permissions in your systems: for an employee on their own login and for a background agent on a service account alike. Nobody changes a role in the HR, CRM or finance system.
Connectors open up the use cases. Doing it securely, with the right guardrails and governance in place, is what makes it possible to roll agents out across the whole business.
How are Tool Policies different from the controls built into AI platforms?
So do you still need the controls built into your AI platform? Yes. Keep them switched on: they decide which systems and tools an agent can reach. Tool Policies standardise the governance and guardrails across any platform you deploy, and decide how agents behave once they’re connected. Some platforms are adding argument checks in beta, but each one governs only its own agents, so a company running Copilot, Claude and Glean side by side writes the same rule three times. With Tool Policies, it writes the rule once.
“IT and functional leaders couldn’t safely enable more access, because of what ungoverned access would imply. Tool Policies is the exhaustive answer to that problem.”
Guillaume Lebedel, CTO and Co-founder, StackOne
Who it applies to: everyone, or users and groups synced from your identity provider, with named exemptions.
Where it applies: the whole project, a connector category, one connector or one linked account. A category rule also covers connectors you add later, and a name pattern like *delete* catches every matching tool.
What it does: five rule types. PII redaction strips personal data or blocks the call. Tool rules block tools. Value rules limit what a tool accepts, such as payments between $25 and $500. Field rules hide or protect fields by path. Semantic rules ask a plain-language question.
StackOne checks the input before it calls the business system, and the response before the agent sees it. Rules compile to Cedar, the open-source policy language built by AWS, where if two rules disagree the deny always wins. PII redaction runs on StackOne’s own state of the art model, which finds email addresses, phone numbers, SSNs and card numbers anywhere in a payload.
Build a policy in the no-code builder, have the StackOne Assistant draft it, or write it as Cedar or JSON. Every rule can run in Monitor first, which records who it would affect without blocking anything, and every decision is logged. Guillaume describes the rollout sequence we recommend in Monitor first, then enforce .
Semantic rules: one question instead of a hundred field names
Say you want to keep pay away from every agent. The obvious way is a field rule: you name the field to hide, such as salary. But every system names it differently. One HR system calls it salary, another base_pay, a third pay_rate, so you’d need a rule for each one. And a Slack message about someone’s raise has no field to name at all.
A semantic rule asks a question instead, such as “Is this compensation?” or “Does this message contain unreleased financial results?”, and a decision model answers it on every call. One rule covers the HR systems you connect today and the ones you add next year. When the model cannot decide, the field stays masked.
How do you control which finance actions and fields AI agents can access?
If a finance team wants one agent for analytics and forecasting and another to pay supplier bills, the security review is going to turn up four problems. The payments agent could pay any amount. Anyone with access to supplier records could change where the money goes. Contractors could delete records. And the analytics agent would see every supplier’s bank details.
Tool Policies fix each one with a rule for the group it applies to, using the groups already in your identity provider:
Accounts Payable can pay through an agent, between $25 and $500 per payment, and can prepare invoices but not authorize them. A $250 bill goes through. A $4,800 payment is stopped before the finance system is called.
Procurement can read supplier records, but a change to a bank account number is refused. A request to change a supplier’s bank details in SAP is blocked and logged.
Contractors cannot run any tool matching *delete* , on any connector.
Data teams can read accounting records with bank details masked.
See it end to end in our vendor payment scheduling and AP invoice processing use cases.
Can AI agents analyse patient data under HIPAA?
Take a healthcare provider with several clinics. Its operations team wants an agent to read patient messages and support tickets and report why appointments get rescheduled, or which clinics get the most billing questions. The analysis needs what each message says, not who sent it. HIPAA’s minimum necessary standard asks for exactly that: use only the patient data a task needs.
So the identifiers stay out of the agent’s context window and out of your logs:
A PII redaction rule on the patient support and ticketing connectors replaces email addresses, phone numbers, SSNs and card numbers with a tag that names the type, such as <POLICY_REDACTED:EMAIL> , wherever they appear in the text.
Field rules mask the fields that name the patient directly, such as name, date of birth and home address.
Advanced Logs redaction removes personal data from stored request and response bodies.
The agent reads hundreds of messages and reports the top reasons for rescheduling without seeing who wrote any of them. Every redaction is recorded in the action log for your compliance team. A Business Associate Agreement is available on Enterprise plans where PHI is in scope. I covered the wider picture in the reality of connecting AI agents to healthcare systems .
How do you hide salary and personal data from AI agents in HR systems?
People Ops coordinators use agents for onboarding, leave and employee questions. They need job titles, start dates and managers, and they do not need salary, bank details or Social Security numbers. For the People Operations group:
A semantic rule asks “Is this compensation?” on every HR read and write tool, so pay is hidden whatever each HR system calls it, and a raise attempted through the agent is refused.
Field rules mask Social Security number, personal email, home address and bank details on employee records.
PII redaction catches an SSN or phone number typed into a case note, where no field rule can reach.
The Total Rewards team is exempt, so their agents still see pay.
A coordinator asks for a new starter’s details and gets the title, start date and manager, with pay and identifiers removed before the agent sees the record. The coordinator may see more in the HR system itself: the policy governs the agent separately from the person, the gap Guillaume described in say yes to the Copilot, not to the salary field .
See how this fits an onboarding agent in our employee onboarding use case .
Why is StackOne the most complete way to secure AI agent tool calls?
Every action an agent takes in your business is a tool call. Updating a record, sending a message, paying a bill: each one goes through a tool, whether the agent runs in Claude , ChatGPT , Copilot , Gemini Enterprise , Glean or your own code. That makes the tool call the one point every agent passes through, and the place where security does the most good.
Securing a tool call down to the field means owning the tool. A gateway that passes calls through to other people’s MCP servers can allow or block a whole serve

[truncated]
