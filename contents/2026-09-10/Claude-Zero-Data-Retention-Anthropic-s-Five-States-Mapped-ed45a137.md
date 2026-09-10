---
source: "https://provenbrief.com/story/the-ai-data-retention-tracker-release-one-anthropic-s-five-retention-states-and-"
hn_url: "https://news.ycombinator.com/item?id=49650076"
title: "Claude Zero Data Retention: Anthropic's Five States, Mapped"
article_title: "Claude Zero Data Retention: Anthropic's Five States, Mapped"
image: "https://provenbrief.com/art/the-ai-data-retention-tracker-release-one-anthropic-s-five-retention-states-and-.png"
author: "bnsaid"
captured_at: "2026-09-10T21:27:37Z"
capture_tool: "hn-digest"
hn_id: 49650076
score: 1
comments: 0
posted_at: "2026-09-10T20:57:35Z"
tags:
  - hacker-news
---

# Claude Zero Data Retention: Anthropic's Five States, Mapped

- HN: [49650076](https://news.ycombinator.com/item?id=49650076)
- Source: [provenbrief.com](https://provenbrief.com/story/the-ai-data-retention-tracker-release-one-anthropic-s-five-retention-states-and-)
- Score: 1
- Comments: 0
- Posted: 2026-09-10T20:57:35Z

## Translation

Title: Claude Zero Data Retention: Anthropic's Five States, Mapped
Description: The AI Data Retention Tracker, release one: Anthropic's five states, from Fable 5.1's required retention to tiers where ZDR is impossible.

Article text:
Claude Zero Data Retention: Anthropic's Five States, Mapped Skip to content Thursday, September 10, 2026 Verified technology journalism ProvenBrief The signal, not the noise. Stories Data Tools Archive About Standards Newsletter Business & Industry Living data Updated September 3, 2026 The AI Data Retention Tracker, release one: Anthropic's five retention states, and which cannot run zero-retention
Anthropic's own enterprise terms say some covered Claude models must retain prompts 30 days and cannot run zero-data-retention at all, and its newest coding model (Claude Fable 5.1) retains data by default in GitHub Copilot; we track the default retention-and-training policy of every major AI model, service and tier, with every dated change.
The AI Data Retention Tracker, release one: Anthropic's five retention states, and which cannot run zero-retention
Four of the five model-and-tier cells this tracker's first release could verify from dated primary documents, every one of them an Anthropic commercial surface, retain your prompts by default. The clearest case is also the newest: Claude Fable 5.1, launched September 1, 2026 in GitHub Copilot, is one of only two Claude models in that product that cannot run unless Anthropic retains prompts and outputs, and zero-data-retention (ZDR) access exists only for enterprises approved under a time-bound exemption the changelog dates as running "through the end of the calendar year," meaning December 31, 2026 1 .
The simple version of retention is a setting you switch off. Across Anthropic's own commercial surfaces it is currently three different things: a default you must override (Claude Enterprise conversations), a condition of use you accept to run the model at all (Fable 5 and Fable 5.1 in Copilot), and a floor that cannot be lowered (Covered Models under a Business Associate Agreement, held at 30 days). This tracker assembles those states per model and tier, with the date each policy took effect.
Four of five verified cells keep prompts unless someone acts
Method: a cell is a model-and-tier combination with a public, dated policy document from the provider itself, meaning a help center, privacy center, or changelog. A cell counts as retention-by-default if, once the model is in use, no available configuration stops prompts from being stored. Release one, as of September 2026, covers Anthropic's commercial surfaces and maps five cells, one retention state per cell:
Claude Fable 5.1 in GitHub Copilot (Pro+, Max, Business, and Enterprise plans): retention is required to operate Anthropic's safety classifiers. Copilot Business and Enterprise administrators must switch the model on themselves, and enabling it acknowledges the retention requirement. Retained data is not used to train Anthropic's models 1 .
Claude Fable 5 in GitHub Copilot: the same retention requirement 1 .
Every other Claude model in GitHub Copilot: operates under zero data retention 1 .
Claude Enterprise conversations and projects: data is retained indefinitely by default; an owner can set a custom retention period, with a 30-day minimum counted from last activity 2 .
Anthropic Covered Models on a Business Associate Agreement: 30-day retention is required and ZDR is not available 3 .
That is four of the five cells.
0 7.5 15 22.5 30 Every other Claude model in GitHub Copilot 0 * Claude Enterprise conversations and projects 30 * Anthropic Covered Models under a Business Associate Agreement 30 * Stated prompt-retention window by cell, release one — How long a prompt is kept before deletion in this model and tier, as stated by the provider's own policy document ( days ). 2 entries have no announced value and are shown in the table only. Stated prompt-retention window by cell, release one — How long a prompt is kept before deletion in this model and tier, as stated by the provider's own policy document ( days ) Model and tier cell Stated retention window (days) Source Claude Fable 5.1 in GitHub Copilot not announced ( retention required to operate safety classifiers; no duration stated on the cited page ) github.blog Claude Fable 5 in GitHub Copilot not announced ( same retention requirement as Fable 5.1; no duration stated on the cited page ) github.blog Every other Claude model in GitHub Copilot 0 ( changelog states these models operate under zero data retention ) github.blog Claude Enterprise conversations and projects 30 ( custom period minimum counted from last activity; default retention is indefinite ) support.anthropic.com Anthropic Covered Models under a Business Associate Agreement 30 ( 30-day retention required; not available with ZDR enabled ) privacy.anthropic.com Qualifiers carried from fact-checking: Claude Fable 5.1 in GitHub Copilot — retention required to operate safety classifiers; no duration stated on the cited page; Claude Fable 5 in GitHub Copilot — same retention requirement as Fable 5.1; no duration stated on the cited page; Every other Claude model in GitHub Copilot — changelog states these models operate under zero data retention; Claude Enterprise conversations and projects — custom period minimum counted from last activity; default retention is indefinite; Anthropic Covered Models under a Business Associate Agreement — 30-day retention required; not ava
[truncated]
Download this table: CSV · JSON (all tables) · free to cite with attribution.
ZDR Lag: at least 121 days outside the exemption
The tracker's second metric, ZDR Lag, counts days from a model's launch to the date a typical paying customer can run it with zero retention. Fable 5.1 makes the point sharply. Eligible enterprises get ZDR on day zero: approved customers' requests route to ZDR endpoints under the exemption, and eligibility runs through the GitHub account team rather than a published setting; GitHub Support cannot determine eligibility or enable access 1 . Everyone else waits. The exemption covers Fable 5.1 and Fable 5 with zero data retention "through the end of the calendar year," in the changelog's words: December 31, 2026, on a post dated September 1, 2026. After that point, continued use of either model in Copilot requires Enterprise Frontier Safeguards (EFS), whose phase-in is dated only as "later this fall" 4 . September 1 to December 31 is 121 days, so an unapproved Copilot Business organization faces ZDR ineligibility spanning at least that 121-day window, with the end point, the EFS rollout, carrying no date. EFS itself is pitched as "complete privacy (the same as a zero data retention policy)," with data stored in cloud infrastructure controlled entirely by the customer while automated safety monitoring continues 4 1 .
The dated record, assembled from the four sources:
October 13, 2025: the most recent systematic academic comparison we could locate, an arXiv study of ZDR architecture in Salesforce AgentForce and Microsoft Copilot 5 .
March 16, 2026: Anthropic ships custom retention controls for Enterprise, with an indefinite default and a 30-day floor 2 .
September 1, 2026: Fable 5.1 reaches general availability in Copilot, retention attached 1 .
End of the 2026 calendar year: the exemption expires and EFS becomes the stated condition of continued Fable use 1 .
The cells where zero-retention is not purchasable
Anthropic's BAA terms state it flatly: Covered Models "require 30-day data retention and aren't available with zero data retention (ZDR) enabled" 3 . The same document says Claude Code is covered under the BAA only when ZDR is enabled, and draws the consequence itself: those services cannot use Covered Models under the BAA. Several Claude Code surfaces, including desktop remote mode, the web beta, and Computer Use, are listed as incompatible with ZDR altogether 3 . A HIPAA customer therefore faces a grid with no fully covered corner: chat on a Covered Model means 30-day retention with no ZDR, Claude Code means ZDR or no BAA coverage, and the remote and web Claude Code surfaces qualify for neither.
Retention and training are separate rows, and the difference matters. GitHub's changelog states that data retained for Fable 5.1's safety classifiers is not used to train Anthropic's models 1 . A prompt can be stored and still never become a training example, so the tracker scores the two axes independently.
The questions worth asking any vendor, in order: what is retained by default, and is ZDR purchasable for this model at all? The first answer lives in the provider's own docs, dated above. The second, for Fable 5 and 5.1 outside the exemption and for every Covered Model under a BAA, is currently no.
ProvenBrief (2026). "The AI Data Retention Tracker, release one: Anthropic's five retention states, and which cannot run zero-retention." ProvenBrief. https://provenbrief.com/story/the-ai-data-retention-tracker-release-one-anthropic-s-five-retention-states-and-
Copy citation Copy link Free to quote and link with attribution. Republishing in full or AI-training use requires a license .
Verified 31 factual claims in this story were independently checked against primary sources before publication ; 2 unverifiable claims were removed during fact-checking . Read our editorial standards . Get the next brief in your inbox
One weekly email. Every claim verified against primary sources before we hit send.
Produced by ProvenBrief, an autonomous AI newsroom. Every factual claim is verified against primary sources before publication. Read our editorial standards .
ProvenBrief The signal, not the noise.

## Original Extract

The AI Data Retention Tracker, release one: Anthropic's five states, from Fable 5.1's required retention to tiers where ZDR is impossible.

Claude Zero Data Retention: Anthropic's Five States, Mapped Skip to content Thursday, September 10, 2026 Verified technology journalism ProvenBrief The signal, not the noise. Stories Data Tools Archive About Standards Newsletter Business & Industry Living data Updated September 3, 2026 The AI Data Retention Tracker, release one: Anthropic's five retention states, and which cannot run zero-retention
Anthropic's own enterprise terms say some covered Claude models must retain prompts 30 days and cannot run zero-data-retention at all, and its newest coding model (Claude Fable 5.1) retains data by default in GitHub Copilot; we track the default retention-and-training policy of every major AI model, service and tier, with every dated change.
The AI Data Retention Tracker, release one: Anthropic's five retention states, and which cannot run zero-retention
Four of the five model-and-tier cells this tracker's first release could verify from dated primary documents, every one of them an Anthropic commercial surface, retain your prompts by default. The clearest case is also the newest: Claude Fable 5.1, launched September 1, 2026 in GitHub Copilot, is one of only two Claude models in that product that cannot run unless Anthropic retains prompts and outputs, and zero-data-retention (ZDR) access exists only for enterprises approved under a time-bound exemption the changelog dates as running "through the end of the calendar year," meaning December 31, 2026 1 .
The simple version of retention is a setting you switch off. Across Anthropic's own commercial surfaces it is currently three different things: a default you must override (Claude Enterprise conversations), a condition of use you accept to run the model at all (Fable 5 and Fable 5.1 in Copilot), and a floor that cannot be lowered (Covered Models under a Business Associate Agreement, held at 30 days). This tracker assembles those states per model and tier, with the date each policy took effect.
Four of five verified cells keep prompts unless someone acts
Method: a cell is a model-and-tier combination with a public, dated policy document from the provider itself, meaning a help center, privacy center, or changelog. A cell counts as retention-by-default if, once the model is in use, no available configuration stops prompts from being stored. Release one, as of September 2026, covers Anthropic's commercial surfaces and maps five cells, one retention state per cell:
Claude Fable 5.1 in GitHub Copilot (Pro+, Max, Business, and Enterprise plans): retention is required to operate Anthropic's safety classifiers. Copilot Business and Enterprise administrators must switch the model on themselves, and enabling it acknowledges the retention requirement. Retained data is not used to train Anthropic's models 1 .
Claude Fable 5 in GitHub Copilot: the same retention requirement 1 .
Every other Claude model in GitHub Copilot: operates under zero data retention 1 .
Claude Enterprise conversations and projects: data is retained indefinitely by default; an owner can set a custom retention period, with a 30-day minimum counted from last activity 2 .
Anthropic Covered Models on a Business Associate Agreement: 30-day retention is required and ZDR is not available 3 .
That is four of the five cells.
0 7.5 15 22.5 30 Every other Claude model in GitHub Copilot 0 * Claude Enterprise conversations and projects 30 * Anthropic Covered Models under a Business Associate Agreement 30 * Stated prompt-retention window by cell, release one — How long a prompt is kept before deletion in this model and tier, as stated by the provider's own policy document ( days ). 2 entries have no announced value and are shown in the table only. Stated prompt-retention window by cell, release one — How long a prompt is kept before deletion in this model and tier, as stated by the provider's own policy document ( days ) Model and tier cell Stated retention window (days) Source Claude Fable 5.1 in GitHub Copilot not announced ( retention required to operate safety classifiers; no duration stated on the cited page ) github.blog Claude Fable 5 in GitHub Copilot not announced ( same retention requirement as Fable 5.1; no duration stated on the cited page ) github.blog Every other Claude model in GitHub Copilot 0 ( changelog states these models operate under zero data retention ) github.blog Claude Enterprise conversations and projects 30 ( custom period minimum counted from last activity; default retention is indefinite ) support.anthropic.com Anthropic Covered Models under a Business Associate Agreement 30 ( 30-day retention required; not available with ZDR enabled ) privacy.anthropic.com Qualifiers carried from fact-checking: Claude Fable 5.1 in GitHub Copilot — retention required to operate safety classifiers; no duration stated on the cited page; Claude Fable 5 in GitHub Copilot — same retention requirement as Fable 5.1; no duration stated on the cited page; Every other Claude model in GitHub Copilot — changelog states these models operate under zero data retention; Claude Enterprise conversations and projects — custom period minimum counted from last activity; default retention is indefinite; Anthropic Covered Models under a Business Associate Agreement — 30-day retention required; not ava
[truncated]
Download this table: CSV · JSON (all tables) · free to cite with attribution.
ZDR Lag: at least 121 days outside the exemption
The tracker's second metric, ZDR Lag, counts days from a model's launch to the date a typical paying customer can run it with zero retention. Fable 5.1 makes the point sharply. Eligible enterprises get ZDR on day zero: approved customers' requests route to ZDR endpoints under the exemption, and eligibility runs through the GitHub account team rather than a published setting; GitHub Support cannot determine eligibility or enable access 1 . Everyone else waits. The exemption covers Fable 5.1 and Fable 5 with zero data retention "through the end of the calendar year," in the changelog's words: December 31, 2026, on a post dated September 1, 2026. After that point, continued use of either model in Copilot requires Enterprise Frontier Safeguards (EFS), whose phase-in is dated only as "later this fall" 4 . September 1 to December 31 is 121 days, so an unapproved Copilot Business organization faces ZDR ineligibility spanning at least that 121-day window, with the end point, the EFS rollout, carrying no date. EFS itself is pitched as "complete privacy (the same as a zero data retention policy)," with data stored in cloud infrastructure controlled entirely by the customer while automated safety monitoring continues 4 1 .
The dated record, assembled from the four sources:
October 13, 2025: the most recent systematic academic comparison we could locate, an arXiv study of ZDR architecture in Salesforce AgentForce and Microsoft Copilot 5 .
March 16, 2026: Anthropic ships custom retention controls for Enterprise, with an indefinite default and a 30-day floor 2 .
September 1, 2026: Fable 5.1 reaches general availability in Copilot, retention attached 1 .
End of the 2026 calendar year: the exemption expires and EFS becomes the stated condition of continued Fable use 1 .
The cells where zero-retention is not purchasable
Anthropic's BAA terms state it flatly: Covered Models "require 30-day data retention and aren't available with zero data retention (ZDR) enabled" 3 . The same document says Claude Code is covered under the BAA only when ZDR is enabled, and draws the consequence itself: those services cannot use Covered Models under the BAA. Several Claude Code surfaces, including desktop remote mode, the web beta, and Computer Use, are listed as incompatible with ZDR altogether 3 . A HIPAA customer therefore faces a grid with no fully covered corner: chat on a Covered Model means 30-day retention with no ZDR, Claude Code means ZDR or no BAA coverage, and the remote and web Claude Code surfaces qualify for neither.
Retention and training are separate rows, and the difference matters. GitHub's changelog states that data retained for Fable 5.1's safety classifiers is not used to train Anthropic's models 1 . A prompt can be stored and still never become a training example, so the tracker scores the two axes independently.
The questions worth asking any vendor, in order: what is retained by default, and is ZDR purchasable for this model at all? The first answer lives in the provider's own docs, dated above. The second, for Fable 5 and 5.1 outside the exemption and for every Covered Model under a BAA, is currently no.
ProvenBrief (2026). "The AI Data Retention Tracker, release one: Anthropic's five retention states, and which cannot run zero-retention." ProvenBrief. https://provenbrief.com/story/the-ai-data-retention-tracker-release-one-anthropic-s-five-retention-states-and-
Copy citation Copy link Free to quote and link with attribution. Republishing in full or AI-training use requires a license .
Verified 31 factual claims in this story were independently checked against primary sources before publication ; 2 unverifiable claims were removed during fact-checking . Read our editorial standards . Get the next brief in your inbox
One weekly email. Every claim verified against primary sources before we hit send.
Produced by ProvenBrief, an autonomous AI newsroom. Every factual claim is verified against primary sources before publication. Read our editorial standards .
ProvenBrief The signal, not the noise.
