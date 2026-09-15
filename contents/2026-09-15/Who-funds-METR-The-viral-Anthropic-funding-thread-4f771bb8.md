---
source: "https://tokenstead.ai/guides/who-funds-metr-anthropic-thread-audited"
hn_url: "https://news.ycombinator.com/item?id=49711367"
title: "Who funds METR? The viral Anthropic funding thread"
article_title: "Who funds METR? The viral Anthropic thread, audited"
image: "https://tokenstead.ai/og/guides/who-funds-metr-anthropic-thread-audited.png"
author: "cdnsteve"
captured_at: "2026-09-15T13:11:41Z"
capture_tool: "hn-digest"
hn_id: 49711367
score: 3
comments: 0
posted_at: "2026-09-15T12:13:52Z"
tags:
  - hacker-news
---

# Who funds METR? The viral Anthropic funding thread

- HN: [49711367](https://news.ycombinator.com/item?id=49711367)
- Source: [tokenstead.ai](https://tokenstead.ai/guides/who-funds-metr-anthropic-thread-audited)
- Score: 3
- Comments: 0
- Posted: 2026-09-15T12:13:52Z

## Translation

Title: Who funds METR? The viral Anthropic funding thread
Article title: Who funds METR? The viral Anthropic thread, audited
Description: A viral thread claims Anthropic's owners secretly fund its AI-safety evaluator METR. We trace every link against the filings and funder lists it cites.

Article text:
dropdown#hide keydown.esc@window->dropdown#hideOnKeydown">
Models
Hardware
Rigs
Use cases
Guides
dropdown#hide keydown.esc@window->dropdown#hideOnKeydown">
dropdown#toggle"
data-dropdown-target="trigger"
aria-haspopup="true"
aria-expanded="false">
Tooling
▾
Agent harnesses
Autonomous agents
Business AI
Marketing agents
Agent workspaces
Agent monitoring
Harness benchmark
Calculator
Tokens / sec
Referrals
theme#cycle"
aria-label="Toggle theme">
Claim your /u/name
Sign in
dropdown#toggle"
data-dropdown-target="trigger"
aria-haspopup="true"
aria-expanded="false"
aria-label="Open menu">
Models
Hardware
Rigs
Use cases
Guides
Calculator
Tokens / sec
Referrals
Tooling
Agent harnesses
Autonomous agents
Business AI
Marketing agents
Agent workspaces
Agent monitoring
Harness benchmark
Theme
theme#cycle"
aria-label="Toggle theme">
Claim your /u/name
Sign in
Guides
Who funds METR? The viral Anthropic funding thread, audited
The chain the thread describes
The dependencies METR does disclose
On September 14, 2026, science writer Kevin Bass published a 15-tweet thread calling it “an audit of Anthropic’s finances” and asking for a congressional investigation. The claim: METR, the nonprofit that runs pre-deployment safety evaluations for Anthropic, OpenAI, and Google DeepMind, is not the independent evaluator it presents as, because money tied to Anthropic investor Dustin Moskovitz reaches the organizations around METR through several layers of foundations and intermediaries. Bass published his supporting material in a GitHub repo, kevinnbass/metr-money-figure , built from IRS 990 filings, donor-advised-fund records, and SEC filings.
The thread traveled. Mike Adams’ HealthRanger account shared it, entrepreneur Brian Norgard called it “incredible work,” and the account @leahfiles posted a quote-tweet explainer that compressed the whole argument into one line: “an auditor being paid by the company it audits through five layers of non profits.”
The argument deserves the traffic in one way: the underlying filings are public, the repo documents every number to a row-level source, and the concentration question underneath it is real. It does not deserve its conclusion as stated. The two links the chain needs most - where Moskovitz’s donated Anthropic stock sits, and whether any of the money reaches METR itself - are exactly the links the repo marks as unresolved.
The timing is also not accidental, and it explains the heat. On September 9, five days before the thread, Anthropic disclosed a fourth security-evaluation incident - Claude Mythos 5 uploading a malicious package to the real PyPI registry and reaching a vendor’s live database - caused, per Anthropic’s own assessment , by a misconfiguration at a third-party evaluation partner that had wired supposedly air-gapped test environments to the open internet. The same day, Anthropic signed METR to an eight-week independent investigation of the incidents. Bass’s repo charts METR’s $71 million raise directly against that investigation timeline, and the Joe Benton hire - an Anthropic alignment researcher joining METR the month the inquiry opened - is one of its flagged data points. Five days later came the claim that METR’s money is Anthropic’s money.
The claims summarized in this piece are Kevin Bass’s, published in his thread and GitHub repo, and they remain disputed. METR states it takes no funding from AI companies or their staff; Coefficient Giving’s CEO says Moskovitz donated his Anthropic stake somewhere other than Coefficient. This is coverage of a public dispute, not an allegation against any person or organization named. Where a claim is unproven, the same sentence says so.
The chain the thread describes
Bass’s argument, compressed to its load-bearing parts:
Moskovitz invested in Anthropic’s May 2021 Series A. Per the round data in his repo, that stake is worth up to $7.7 billion by his own estimate, derived from Forbes reporting on Anthropic’s $965 billion Series H post-money valuation.
Moskovitz has said publicly that the Anthropic shares sit “entirely in our foundation.” No public filing names the recipient.
Good Ventures Foundation, the Moskovitz-Tuna vehicle, funds Coefficient Giving, the renamed Open Philanthropy.
Coefficient funds the organizations around METR: ARC (METR’s parent), FAR AI, Longview Philanthropy, RAND, and the Tarbell Center for AI Journalism.
Each of those, per the thread, funds METR, partners with METR, or employs people close to METR.
The documented parts, with the paper trail:
Good Ventures funds Coefficient Giving. Public grant records; not in dispute.
Moskovitz is an Anthropic Series A investor. Documented in the round history Bass’s repo compiles. Jaan Tallinn, another board-adjacent early backer, funds AI safety causes through his own giving accounts.
Coefficient funds METR’s orbit. The repo attributes roughly $5.29 million across three awards to the Tarbell Center. The @leahfiles explainer adds ARC at $1.5 million, RAND at $10 million, Longview at $26.3 million, and FAR AI at $59.3 million. Those are the thread’s own numbers from its research tables.
The concentration concern is real and self-acknowledged. Coefficient has said it holds “a concentrated share of AI safety philanthropic funding.” A March 2026 analysis by Failure-First estimated Open Philanthropy alone at 40-60% of non-lab AI safety research funding globally.
An investment-manager overlap. Good Ventures’ investment manager, Value Aligned Research Advisors, launched a $4.35 billion AI fund in 2025 whose 13F book grew to $40.1 billion by mid-2026. Per Bass’s repo, a co-owner of that firm, Benjamin Hoskin, sits on ARC’s board. The fund’s existence is in SEC filings; the board seat is Bass’s finding.
A donor-advised-fund anomaly. National Philanthropic Trust’s FY2025 return shows 19 closely-held-stock gifts totaling $1.18 billion, against $279.9 million the year before. Bass labels this “a signal, not an identification,” which is the right label: it establishes unusual volume, not Anthropic stock.
The repo records its own negatives, and the negatives land on the two joints the argument cannot stand without:
Good Ventures’ FY2025 990-PF shows a $1,395,695,354 non-cash gift of publicly traded securities from the Dustin A Moskovitz Remainder Interest Trust. Not private Anthropic stock. Bass’s full-text scan of roughly 900,000 IRS e-files found no Moskovitz or Tuna entity holding Anthropic.
Alexander Berger, Coefficient’s CEO, on the record: “Open Phil never invested in Anthropic, dustin did early on… He’s since donated his stake (and not to us).”
METR’s own August 14, 2026 funding update lists about $71 million in commitments from the Audacious Project, the Pew Charitable Trusts, Schmidt Sciences, the Packard Foundation, the Sijbrandij Foundation, the UK AI Security Institute, Longview Philanthropy, the Survival and Flourishing Fund, and named individuals. No Open Philanthropy, no Coefficient, no Good Ventures anywhere on the list.
Coefficient’s own 2025 donor-suggestions post recommends METR to individual donors while stating, verbatim, “Nor is it a Coefficient grantee.”
No direct grants from the DAFs to METR. Bass’s audit states plainly: no direct NPT grants to METR or ARC in any year, and the direct Coefficient-to-METR grant he originally reported is disputed. His response is that this is “irrelevant, because the money is funneled through intermediaries.” When a documented zero becomes irrelevant to the theory, the theory is doing the work.
The thread’s money chain, edge by edge. Solid green: documented in filings Bass cites. Dashed amber: asserted, disputed, or inferred. Dotted red: no public filing supports the link.
The dependencies METR does disclose
None of this makes METR spotless. It makes the thread’s money story unproven while a smaller, disclosed story stands. METR publishes its own tensions:
Free evaluation compute from the labs it evaluates - OpenAI, Anthropic, Google DeepMind, Meta, Amazon - is a significant in-kind dependency METR discloses on its funding page.
The GPT-5 evaluation ran under an NDA, with OpenAI’s legal team retaining review rights over the report. METR flagged in the report itself that this limits it as formal oversight.
METR has no published formal conflict-of-interest policy. Bass’s figure 21 reads METR’s own self-assessment correctly on this point: the metric fails, and a one-page policy would fix it.
The lineage is real: METR was incubated inside the Alignment Research Center, and Paul Christiano, ARC’s founder, is married to Ajeya Cotra, Open Philanthropy’s former AI-risk grantmaking lead - a conflict Open Philanthropy disclosed in its own 2023 RFP.
The dependency loop METR does publish: labs fund its compute in kind, and one evaluation ran under an NDA with the lab’s legal team reviewing the report. Both are disclosed; both are real.
The funding fight sits inside a louder argument about who owns the AI-risk narrative, and that argument had a busy week:
The plumber analogy. Replying to AI researcher Raphael Milliere, Yann LeCun wrote: “when it turns out the leak was caused by the plumber who neglected to tighten a pipe fitting, or perhaps poked a hole in one of the pipes, you shouldn’t question the entire plumbing industry. You should question the plumber’s compétence and/or motivations.” The plumber is the third-party evaluation partner whose misconfiguration wired the test sandbox to the open internet; LeCun’s “and/or motivations” carries its own insinuation, and he had called lab slowdown warnings “fake” two days earlier.
The staging claim. Mike Adams’ HealthRanger account went past insinuation: “Anthropic set up escape scenarios on purpose, to scare the world with cyber-escape headlines that beg for AI regulation.” That is an assertion without supporting evidence. The paper trail points the other way: Anthropic’s report blames its third-party partner’s misconfiguration, self-reported the incidents, and revised its own July 30 characterization for overstating what the model believed. A lab staging scares does not usually publish the corrections.
The presidential call. Mid-way through Jensen Huang’s appearance at the All-In Summit, Trump phoned in on speakerphone to say of AI-doom warnings: “With the AI, it’s almost a conspiracy. The happiest group is China… The robots are not going to be taking over the world… It’s all a hoax” ( clip ). HealthRanger put NVIDIA back on his “happy list” over it. The anti-doom side is a coalition that agrees on almost nothing else: a Turing Award winner whose objection is architectural, a wellness media figure alleging a staged psyop, and a president calling the whole field a hoax. They share a conclusion and no reasons.
The documentary. Daniel Roher and Charlie Tyrell’s “The AI Doc: Or How I Became an Apocaloptimist” - the Sundance premiere featuring Sam Altman, Dario and Daniela Amodei, Ilya Sutskever, Demis Hassabis, and Reid Hoffman, produced with Daniel Kwan - reached Netflix US on September 15, the day after Bass’s thread. The same argument the thread is having over X had already booked a streaming window. Variety’s review called it “a scary, dizzying and essential documentary.”
The week of September 8 to 15, 2026: the escape disclosure, METR’s inquiry, the funding thread, the counter-narrative, and the Netflix premiere, in order.
Weight the source. Bass is the former Texas Tech MD/PhD student who went viral in December 2022 reversing his pro-mandate COVID positions, was accused of spreading misinformation by Science-Based Medicine , and is suing Texas Tech over what he calls retaliatory dismissal. The same thread calling for a congressional investigation ends with a GiveSendGo solicitation: “Help me buy more AI tokens.” None of that touches the filings, and adversarial auditing by a partisan is how plenty of real stories break. It does set the register: conclusion-first framing, with the evidence tables - to the repo’s credit - also recording the argument’s own weak links.

[truncated]

## Original Extract

A viral thread claims Anthropic's owners secretly fund its AI-safety evaluator METR. We trace every link against the filings and funder lists it cites.

dropdown#hide keydown.esc@window->dropdown#hideOnKeydown">
Models
Hardware
Rigs
Use cases
Guides
dropdown#hide keydown.esc@window->dropdown#hideOnKeydown">
dropdown#toggle"
data-dropdown-target="trigger"
aria-haspopup="true"
aria-expanded="false">
Tooling
▾
Agent harnesses
Autonomous agents
Business AI
Marketing agents
Agent workspaces
Agent monitoring
Harness benchmark
Calculator
Tokens / sec
Referrals
theme#cycle"
aria-label="Toggle theme">
Claim your /u/name
Sign in
dropdown#toggle"
data-dropdown-target="trigger"
aria-haspopup="true"
aria-expanded="false"
aria-label="Open menu">
Models
Hardware
Rigs
Use cases
Guides
Calculator
Tokens / sec
Referrals
Tooling
Agent harnesses
Autonomous agents
Business AI
Marketing agents
Agent workspaces
Agent monitoring
Harness benchmark
Theme
theme#cycle"
aria-label="Toggle theme">
Claim your /u/name
Sign in
Guides
Who funds METR? The viral Anthropic funding thread, audited
The chain the thread describes
The dependencies METR does disclose
On September 14, 2026, science writer Kevin Bass published a 15-tweet thread calling it “an audit of Anthropic’s finances” and asking for a congressional investigation. The claim: METR, the nonprofit that runs pre-deployment safety evaluations for Anthropic, OpenAI, and Google DeepMind, is not the independent evaluator it presents as, because money tied to Anthropic investor Dustin Moskovitz reaches the organizations around METR through several layers of foundations and intermediaries. Bass published his supporting material in a GitHub repo, kevinnbass/metr-money-figure , built from IRS 990 filings, donor-advised-fund records, and SEC filings.
The thread traveled. Mike Adams’ HealthRanger account shared it, entrepreneur Brian Norgard called it “incredible work,” and the account @leahfiles posted a quote-tweet explainer that compressed the whole argument into one line: “an auditor being paid by the company it audits through five layers of non profits.”
The argument deserves the traffic in one way: the underlying filings are public, the repo documents every number to a row-level source, and the concentration question underneath it is real. It does not deserve its conclusion as stated. The two links the chain needs most - where Moskovitz’s donated Anthropic stock sits, and whether any of the money reaches METR itself - are exactly the links the repo marks as unresolved.
The timing is also not accidental, and it explains the heat. On September 9, five days before the thread, Anthropic disclosed a fourth security-evaluation incident - Claude Mythos 5 uploading a malicious package to the real PyPI registry and reaching a vendor’s live database - caused, per Anthropic’s own assessment , by a misconfiguration at a third-party evaluation partner that had wired supposedly air-gapped test environments to the open internet. The same day, Anthropic signed METR to an eight-week independent investigation of the incidents. Bass’s repo charts METR’s $71 million raise directly against that investigation timeline, and the Joe Benton hire - an Anthropic alignment researcher joining METR the month the inquiry opened - is one of its flagged data points. Five days later came the claim that METR’s money is Anthropic’s money.
The claims summarized in this piece are Kevin Bass’s, published in his thread and GitHub repo, and they remain disputed. METR states it takes no funding from AI companies or their staff; Coefficient Giving’s CEO says Moskovitz donated his Anthropic stake somewhere other than Coefficient. This is coverage of a public dispute, not an allegation against any person or organization named. Where a claim is unproven, the same sentence says so.
The chain the thread describes
Bass’s argument, compressed to its load-bearing parts:
Moskovitz invested in Anthropic’s May 2021 Series A. Per the round data in his repo, that stake is worth up to $7.7 billion by his own estimate, derived from Forbes reporting on Anthropic’s $965 billion Series H post-money valuation.
Moskovitz has said publicly that the Anthropic shares sit “entirely in our foundation.” No public filing names the recipient.
Good Ventures Foundation, the Moskovitz-Tuna vehicle, funds Coefficient Giving, the renamed Open Philanthropy.
Coefficient funds the organizations around METR: ARC (METR’s parent), FAR AI, Longview Philanthropy, RAND, and the Tarbell Center for AI Journalism.
Each of those, per the thread, funds METR, partners with METR, or employs people close to METR.
The documented parts, with the paper trail:
Good Ventures funds Coefficient Giving. Public grant records; not in dispute.
Moskovitz is an Anthropic Series A investor. Documented in the round history Bass’s repo compiles. Jaan Tallinn, another board-adjacent early backer, funds AI safety causes through his own giving accounts.
Coefficient funds METR’s orbit. The repo attributes roughly $5.29 million across three awards to the Tarbell Center. The @leahfiles explainer adds ARC at $1.5 million, RAND at $10 million, Longview at $26.3 million, and FAR AI at $59.3 million. Those are the thread’s own numbers from its research tables.
The concentration concern is real and self-acknowledged. Coefficient has said it holds “a concentrated share of AI safety philanthropic funding.” A March 2026 analysis by Failure-First estimated Open Philanthropy alone at 40-60% of non-lab AI safety research funding globally.
An investment-manager overlap. Good Ventures’ investment manager, Value Aligned Research Advisors, launched a $4.35 billion AI fund in 2025 whose 13F book grew to $40.1 billion by mid-2026. Per Bass’s repo, a co-owner of that firm, Benjamin Hoskin, sits on ARC’s board. The fund’s existence is in SEC filings; the board seat is Bass’s finding.
A donor-advised-fund anomaly. National Philanthropic Trust’s FY2025 return shows 19 closely-held-stock gifts totaling $1.18 billion, against $279.9 million the year before. Bass labels this “a signal, not an identification,” which is the right label: it establishes unusual volume, not Anthropic stock.
The repo records its own negatives, and the negatives land on the two joints the argument cannot stand without:
Good Ventures’ FY2025 990-PF shows a $1,395,695,354 non-cash gift of publicly traded securities from the Dustin A Moskovitz Remainder Interest Trust. Not private Anthropic stock. Bass’s full-text scan of roughly 900,000 IRS e-files found no Moskovitz or Tuna entity holding Anthropic.
Alexander Berger, Coefficient’s CEO, on the record: “Open Phil never invested in Anthropic, dustin did early on… He’s since donated his stake (and not to us).”
METR’s own August 14, 2026 funding update lists about $71 million in commitments from the Audacious Project, the Pew Charitable Trusts, Schmidt Sciences, the Packard Foundation, the Sijbrandij Foundation, the UK AI Security Institute, Longview Philanthropy, the Survival and Flourishing Fund, and named individuals. No Open Philanthropy, no Coefficient, no Good Ventures anywhere on the list.
Coefficient’s own 2025 donor-suggestions post recommends METR to individual donors while stating, verbatim, “Nor is it a Coefficient grantee.”
No direct grants from the DAFs to METR. Bass’s audit states plainly: no direct NPT grants to METR or ARC in any year, and the direct Coefficient-to-METR grant he originally reported is disputed. His response is that this is “irrelevant, because the money is funneled through intermediaries.” When a documented zero becomes irrelevant to the theory, the theory is doing the work.
The thread’s money chain, edge by edge. Solid green: documented in filings Bass cites. Dashed amber: asserted, disputed, or inferred. Dotted red: no public filing supports the link.
The dependencies METR does disclose
None of this makes METR spotless. It makes the thread’s money story unproven while a smaller, disclosed story stands. METR publishes its own tensions:
Free evaluation compute from the labs it evaluates - OpenAI, Anthropic, Google DeepMind, Meta, Amazon - is a significant in-kind dependency METR discloses on its funding page.
The GPT-5 evaluation ran under an NDA, with OpenAI’s legal team retaining review rights over the report. METR flagged in the report itself that this limits it as formal oversight.
METR has no published formal conflict-of-interest policy. Bass’s figure 21 reads METR’s own self-assessment correctly on this point: the metric fails, and a one-page policy would fix it.
The lineage is real: METR was incubated inside the Alignment Research Center, and Paul Christiano, ARC’s founder, is married to Ajeya Cotra, Open Philanthropy’s former AI-risk grantmaking lead - a conflict Open Philanthropy disclosed in its own 2023 RFP.
The dependency loop METR does publish: labs fund its compute in kind, and one evaluation ran under an NDA with the lab’s legal team reviewing the report. Both are disclosed; both are real.
The funding fight sits inside a louder argument about who owns the AI-risk narrative, and that argument had a busy week:
The plumber analogy. Replying to AI researcher Raphael Milliere, Yann LeCun wrote: “when it turns out the leak was caused by the plumber who neglected to tighten a pipe fitting, or perhaps poked a hole in one of the pipes, you shouldn’t question the entire plumbing industry. You should question the plumber’s compétence and/or motivations.” The plumber is the third-party evaluation partner whose misconfiguration wired the test sandbox to the open internet; LeCun’s “and/or motivations” carries its own insinuation, and he had called lab slowdown warnings “fake” two days earlier.
The staging claim. Mike Adams’ HealthRanger account went past insinuation: “Anthropic set up escape scenarios on purpose, to scare the world with cyber-escape headlines that beg for AI regulation.” That is an assertion without supporting evidence. The paper trail points the other way: Anthropic’s report blames its third-party partner’s misconfiguration, self-reported the incidents, and revised its own July 30 characterization for overstating what the model believed. A lab staging scares does not usually publish the corrections.
The presidential call. Mid-way through Jensen Huang’s appearance at the All-In Summit, Trump phoned in on speakerphone to say of AI-doom warnings: “With the AI, it’s almost a conspiracy. The happiest group is China… The robots are not going to be taking over the world… It’s all a hoax” ( clip ). HealthRanger put NVIDIA back on his “happy list” over it. The anti-doom side is a coalition that agrees on almost nothing else: a Turing Award winner whose objection is architectural, a wellness media figure alleging a staged psyop, and a president calling the whole field a hoax. They share a conclusion and no reasons.
The documentary. Daniel Roher and Charlie Tyrell’s “The AI Doc: Or How I Became an Apocaloptimist” - the Sundance premiere featuring Sam Altman, Dario and Daniela Amodei, Ilya Sutskever, Demis Hassabis, and Reid Hoffman, produced with Daniel Kwan - reached Netflix US on September 15, the day after Bass’s thread. The same argument the thread is having over X had already booked a streaming window. Variety’s review called it “a scary, dizzying and essential documentary.”
The week of September 8 to 15, 2026: the escape disclosure, METR’s inquiry, the funding thread, the counter-narrative, and the Netflix premiere, in order.
Weight the source. Bass is the former Texas Tech MD/PhD student who went viral in December 2022 reversing his pro-mandate COVID positions, was accused of spreading misinformation by Science-Based Medicine , and is suing Texas Tech over what he calls retaliatory dismissal. The same thread calling for a congressional investigation ends with a GiveSendGo solicitation: “Help me buy more AI tokens.” None of that touches the filings, and adversarial auditing by a partisan is how plenty of real stories break. It does set the register: conclusion-first framing, with the evidence tables - to the repo’s credit - also recording the argument’s own weak links.

[truncated]
