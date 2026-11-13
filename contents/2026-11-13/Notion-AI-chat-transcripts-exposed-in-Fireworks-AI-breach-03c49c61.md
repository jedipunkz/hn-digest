---
source: "https://www.notion.com/en-gb/security/advisories/security-incident-involving-a-third-party-ai-vendor"
hn_url: "https://news.ycombinator.com/item?id=50477479"
title: "Notion AI chat transcripts exposed in Fireworks AI breach"
article_title: "Fireworks AI incident"
image: "https://www.notion.com/front-static/meta/global-evergreen-03.png"
author: "fourfire"
captured_at: "2026-11-13T03:10:36Z"
capture_tool: "hn-digest"
hn_id: 50477479
score: 2
comments: 0
posted_at: "2026-11-13T02:49:47Z"
tags:
  - hacker-news
---

# Notion AI chat transcripts exposed in Fireworks AI breach

- HN: [50477479](https://news.ycombinator.com/item?id=50477479)
- Source: [www.notion.com](https://www.notion.com/en-gb/security/advisories/security-incident-involving-a-third-party-ai-vendor)
- Score: 2
- Comments: 0
- Posted: 2026-11-13T02:49:47Z

## Translation

Title: Notion AI chat transcripts exposed in Fireworks AI breach
Article title: Fireworks AI incident
Description: Updated version: 10 November 2026. This post informs you about the recent Fireworks AI incident and how it’s impacting <1% of Notion workspaces. For the latest on the Fireworks AI incident, please see their official updates here. We have contacted our customers who were affected by the Fireworks inc
[truncated]

Article text:
Bring everything into one system of record.
Get answers, instantly – with citations.
Keep work moving 24/7 with agents.
. See what’s new → Download the Notion App → Resources Notion for
← All security advisories Updates
Investigating In addition to previously notifying the impacted customers, on November 12 we reported exposed third-party API keys found in affected user messages to the vendors that issued them.
Investigating On 11 November, we revoked exposed Notion API keys that were found in affected user messages during our investigation.
Updated version: 10 November 2026.
This post informs you about the recent Fireworks AI incident and how it’s impacting <1% of Notion workspaces. For the latest on the Fireworks AI incident, please see their official updates here .
We have contacted our customers who were affected by the Fireworks incident. If you haven’t heard from us, we haven’t identified your workspace as affected, and no action is needed right now. Our investigation is ongoing. If we determine that your workspace was impacted, we will contact you directly.
What happened and how we addressed it :
Upon analysis of the dataset Fireworks AI gave us on 8 November, we discovered that it included Notion AI chat transcripts used for the Insights feature, which for the vast majority of impacted workspaces were unattributable. “Unattributable” means the chats weren’t labelled with who wrote them or which workspace they came from, and the transcripts did not contain tool call results. Beyond what’s in the chat transcript itself, there’s no information tying the chat back to workspaces, users or companies.
No other Notion features were affected, and we continue to monitor this incident.
What is Fireworks AI and how we used them: Fireworks is one of the most used global AI infrastructure platforms powering tens of thousands of customers and used by hundreds of thousands of developers. We used Fireworks for classification jobs to help customers understand their own AI usage through a feature in Notion called Insights, and we used them as a general inference provider to serve open-weight models inside our product. The Insights feature has been disabled as we investigate this incident, and we are no longer serving open-weight models through Fireworks.
What we recommend for affected customers: Affected customers are able to retrieve and review the impacted data securely inside their workspace, with instructions sent via email. We also recommend that affected customers treat this data as if it were exposed. If those with affected workspaces have pasted sensitive information, passwords, API keys or other credentials into Notion AI during the timeframe of the incident, rotate them.
We will update this page as we learn more.
We take your privacy and the security of your data seriously. Even though this was not a breach of our systems, keeping your data safe is fundamental to using Notion for your life’s work. It is our mission to make sure this never happens again. We will raise the standard for how your data is protected.
On 10 November, we updated this post for style and clarity in response to customer questions and feedback. No changes have been made to previously reported facts. View the previous version of this article .
Originally published 12 November 2026
We shape our tools, and thereafter our tools shape us.

## Original Extract

Updated version: 10 November 2026. This post informs you about the recent Fireworks AI incident and how it’s impacting <1% of Notion workspaces. For the latest on the Fireworks AI incident, please see their official updates here. We have contacted our customers who were affected by the Fireworks inc
[truncated]

Bring everything into one system of record.
Get answers, instantly – with citations.
Keep work moving 24/7 with agents.
. See what’s new → Download the Notion App → Resources Notion for
← All security advisories Updates
Investigating In addition to previously notifying the impacted customers, on November 12 we reported exposed third-party API keys found in affected user messages to the vendors that issued them.
Investigating On 11 November, we revoked exposed Notion API keys that were found in affected user messages during our investigation.
Updated version: 10 November 2026.
This post informs you about the recent Fireworks AI incident and how it’s impacting <1% of Notion workspaces. For the latest on the Fireworks AI incident, please see their official updates here .
We have contacted our customers who were affected by the Fireworks incident. If you haven’t heard from us, we haven’t identified your workspace as affected, and no action is needed right now. Our investigation is ongoing. If we determine that your workspace was impacted, we will contact you directly.
What happened and how we addressed it :
Upon analysis of the dataset Fireworks AI gave us on 8 November, we discovered that it included Notion AI chat transcripts used for the Insights feature, which for the vast majority of impacted workspaces were unattributable. “Unattributable” means the chats weren’t labelled with who wrote them or which workspace they came from, and the transcripts did not contain tool call results. Beyond what’s in the chat transcript itself, there’s no information tying the chat back to workspaces, users or companies.
No other Notion features were affected, and we continue to monitor this incident.
What is Fireworks AI and how we used them: Fireworks is one of the most used global AI infrastructure platforms powering tens of thousands of customers and used by hundreds of thousands of developers. We used Fireworks for classification jobs to help customers understand their own AI usage through a feature in Notion called Insights, and we used them as a general inference provider to serve open-weight models inside our product. The Insights feature has been disabled as we investigate this incident, and we are no longer serving open-weight models through Fireworks.
What we recommend for affected customers: Affected customers are able to retrieve and review the impacted data securely inside their workspace, with instructions sent via email. We also recommend that affected customers treat this data as if it were exposed. If those with affected workspaces have pasted sensitive information, passwords, API keys or other credentials into Notion AI during the timeframe of the incident, rotate them.
We will update this page as we learn more.
We take your privacy and the security of your data seriously. Even though this was not a breach of our systems, keeping your data safe is fundamental to using Notion for your life’s work. It is our mission to make sure this never happens again. We will raise the standard for how your data is protected.
On 10 November, we updated this post for style and clarity in response to customer questions and feedback. No changes have been made to previously reported facts. View the previous version of this article .
Originally published 12 November 2026
We shape our tools, and thereafter our tools shape us.
