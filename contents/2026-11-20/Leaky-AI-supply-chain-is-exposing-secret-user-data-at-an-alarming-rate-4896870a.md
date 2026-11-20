---
source: "https://arstechnica.com/security/2026/11/a-leaky-ai-supply-chain-is-exposing-secret-user-data-at-an-alarming-rate/"
hn_url: "https://news.ycombinator.com/item?id=50582794"
title: "Leaky AI supply chain is exposing secret user data at an alarming rate"
article_title: "Researcher plants 2m canaries into the AI supply chain. A privacy disaster surfaced. - Ars Technica"
image: "https://cdn.arstechnica.net/wp-content/uploads/2026/11/ai-privacy-leakage-1152x648.jpg"
author: "BoiledCabbage"
captured_at: "2026-11-20T22:30:38Z"
capture_tool: "hn-digest"
hn_id: 50582794
score: 3
comments: 0
posted_at: "2026-11-20T21:53:39Z"
tags:
  - hacker-news
---

# Leaky AI supply chain is exposing secret user data at an alarming rate

- HN: [50582794](https://news.ycombinator.com/item?id=50582794)
- Source: [arstechnica.com](https://arstechnica.com/security/2026/11/a-leaky-ai-supply-chain-is-exposing-secret-user-data-at-an-alarming-rate/)
- Score: 3
- Comments: 0
- Posted: 2026-11-20T21:53:39Z

## Translation

Title: Leaky AI supply chain is exposing secret user data at an alarming rate
Article title: Researcher plants 2m canaries into the AI supply chain. A privacy disaster surfaced. - Ars Technica
Description: Test using prompts with unique, trackable secrets proves that data doesn't stay private.

Article text:
Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
CREDENTIAL COMPROMISE AT SCALE
Researcher plants 2m canaries into the AI supply chain. A privacy disaster surfaced.
Test using prompts with unique, trackable secrets proves that data doesn’t stay private.
24
...
...
...
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
Researchers planted 2 million digital tripwires in requests sent through a sprawling network of AI services to find out what happened to sensitive information entrusted to them. What they discovered was an ecosystem that leaked highly sensitive data like a sieve.
The exposure was found through an ingenious experiment designed to test if sensitive information can leak through the vast chain of widely used AI services and how that information is exploited when it does. Many of today’s AI apps use services such as OpenRouter or Vercel that act as brokers. These “AI routers,” sometimes also known as “AI gateways,” receive large numbers of prompts sent from AI app end users and route them to an available inference engine based on criteria such as the one charging the lowest price or being best suited for the given task. Instead of using Anthropic, OpenAI, or another big provider, the developer hires an AI router to find the most suitable model.
Security breaches, prompt theft, and privacy violations
To test for potential leakage, researcher HD Moore assembled a collection of AWS access keys, functioning credentials tied to empty AWS accounts, passwords for services such as SSH, and unique URLs, DNS names, and private git repository references. None of the secrets had ever been used before, and he controlled the infrastructure that received any requests using the credentials. He then seeded each secret into exactly one prompt sent to a single provider. In the event any of the secrets were leaked, these “canaries”—so called because they resemble the role of that bird when taken into coal mines—would allow Moore to identify each link in the chain that was responsible.
The findings are striking. They included:
Evidence that at least three AI providers were compromised by attackers who used their access to siphon prompts they received
Multiple resellers in the inference marketplace that logged the prompts and somehow passed them along to people who used the credentials
Multiple cases of AI providers not following end-user privacy settings embedded into flags specifying terms such as Zero Data Retention (ZDR) or “no data collection”
“We sent nearly two million AI prompts containing unique, trackable secrets to AI providers, then watched where those secrets surfaced,” Moore, the chief research officer of security firm Dragos, wrote Thursday . “What we found revealed security breaches, prompt theft, and widespread failures to protect customer data.”
Although prompts are encrypted in transit, they are decrypted at each hop inside the large inference reseller network. The prompt is first sent to an AI router, which reads the plaintext in it to determine the most suitable inference providers to handle the request. Often, the prompt is then sent to a pool of available models, which sometimes run inside different hosting providers. Prompts can also pass from one AI router (such as Vercel) to another one (e.g. OpenRouter). Each one of these players is likely to log the prompt and store it for as long as its policies call for (if not indefinitely).
The Wild West in the age of AI
The most concerning of the findings was the breach of three separate hosting providers. The attackers who compromised one of the providers reused “a couple hundred” of the AWS keys. Moore wrote:
Each key was sent in exactly one request to one provider, all within a few short windows in late October, and every request we sent during those windows was stolen. Some of the stolen keys came from requests that the provider rejected before a model ever ran—this tells us that the attacker was copying requests at the provider’s front door, and not somewhere deeper in the inference stack. In the other two breaches, some of the stolen requests were marked as ZDR. A privacy flag doesn’t help when the attacker is already inside.
Parties tied to the hack of these three providers then tested each key, in some cases within hours of being seeded, to see what kind of resources it gave access to. The parties looked through S3 buckets, secrets managers, and parameter stores in search of still more credentials. They also tried to enable hosted models in AWS Bedrock and added their own IAM users to maintain access. In all, the stolen bait keys were used “tens of thousands” of times in AWS API calls.
Secrets sent to two of the hacked providers were quickly used together, almost immediately, from the same rented server. That means the providers were almost certainly compromised by the same adversary.
As noted earlier, some of the secrets were stolen by the resellers themselves. To lure prompts to their servers, they offered low prices. Then they turned on request logging, and waited. Moore identified them by sending requests to the same model through different sellers. When credentials were later reused when sent through one seller and not another, the researcher knew it was the reseller, not the model provider, that was responsible.
Last, the canary experiment revealed that some providers are scraping secret URLs put into the prompts, in violation of privacy terms. Moore wrote:
The clearest example involves a Chinese AI lab. This lab’s systems fetched the secret URLs from more than half of the prompts that we sent to its models between July and October, including roughly two-thirds of the requests marked as “no data collection” or Zero Data Retention. The scraper typically showed up two or three days after the request.
This wasn’t limited to requests sent to the lab directly. Several US hosting providers resell this lab’s models, and prompts sent to those providers triggered the exact same scraper—the prompts were flowing back to the lab. A customer who picked a US provider specifically to keep their data in the U.S. didn’t get what they paid for. Later on, a git client on a different cloud started trying to clone the private repositories referenced in our prompts, but only for prompts that used this lab’s models.
The researcher said he found one US hosting provider crawling a secret URL more than two months after it was used in a single inference call to its search endpoint. Another ran a scraper that came back for URLs days after each request, regardless of the AI router used to reach it.
Moore said that the most effective means for preventing leakage is to host models using on-premises servers. That option is expensive, though. For those who can’t afford a self-hosted system, the next best option, Moore said, is to receive service directly from a first-tier provider that provides “one contract and a privacy policy to read.” Moore said that even then, there’s no guarantee no data will leak, but that the measure minimizes the chances. AI users should also avoid accepting default provider or reseller pools and instead pick upstream providers themselves.
The blockbuster research is the latest indictment exposing the lax privacy and security that’s endemic in much of the AI ecosystem. The findings mean that people who want to ensure secrets they put into prompts stay private face a daunting task that isn’t likely to become easier any time soon.
24 Comments
Staff Picks
KingKrayola
Some of us work in environments where our AI prompts concern trade secrets which, if leaked to players in our niche, would give significant advantage to our competitors. I'm thinking specifically of something in my office that will probably lead to an order of magnitude reduction in cost of finished goods. Does Moore's work indicate that prompts such as one might make in a process development context are likely to leak? And leaving aside competitive advantage, would such leaks wreck patentability of otherwise novel ideas?
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
Amazon fights the Delivery Protection Act with homepage banner, in-app pop-up
2.
Google turns to Blue Origin for a rocket, but its long-term problem remains
3.
How FAA can modernize US national airspace system without emergency ground stops
4.
Traumatized teens sue school, cops after Doritos bag triggers AI gun alert
5.
NASA's new Moon plans have a problem: Orion wasn't designed to fly this way
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
Manage Preferences
© 2026 Condé Nast. All rights reserved. Use of and/or
registration on any portion of this site constitutes acceptance of our User Agreement and
Privacy Policy and
Cookie Statement and Ars
Technica Addendum and Your
California Privacy Rights . Ars Technica may earn compensation on
sales from links on this site. Read our
affiliate link policy . The material on this site may not be
reproduced, distributed, transmitted, cached or otherwise used, except
with the prior written permission of Condé Nast. Ad
Choices
Dashboard
Howdy, Account Account Edit Profile

## Original Extract

Test using prompts with unique, trackable secrets proves that data doesn't stay private.

Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
CREDENTIAL COMPROMISE AT SCALE
Researcher plants 2m canaries into the AI supply chain. A privacy disaster surfaced.
Test using prompts with unique, trackable secrets proves that data doesn’t stay private.
24
...
...
...
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
Researchers planted 2 million digital tripwires in requests sent through a sprawling network of AI services to find out what happened to sensitive information entrusted to them. What they discovered was an ecosystem that leaked highly sensitive data like a sieve.
The exposure was found through an ingenious experiment designed to test if sensitive information can leak through the vast chain of widely used AI services and how that information is exploited when it does. Many of today’s AI apps use services such as OpenRouter or Vercel that act as brokers. These “AI routers,” sometimes also known as “AI gateways,” receive large numbers of prompts sent from AI app end users and route them to an available inference engine based on criteria such as the one charging the lowest price or being best suited for the given task. Instead of using Anthropic, OpenAI, or another big provider, the developer hires an AI router to find the most suitable model.
Security breaches, prompt theft, and privacy violations
To test for potential leakage, researcher HD Moore assembled a collection of AWS access keys, functioning credentials tied to empty AWS accounts, passwords for services such as SSH, and unique URLs, DNS names, and private git repository references. None of the secrets had ever been used before, and he controlled the infrastructure that received any requests using the credentials. He then seeded each secret into exactly one prompt sent to a single provider. In the event any of the secrets were leaked, these “canaries”—so called because they resemble the role of that bird when taken into coal mines—would allow Moore to identify each link in the chain that was responsible.
The findings are striking. They included:
Evidence that at least three AI providers were compromised by attackers who used their access to siphon prompts they received
Multiple resellers in the inference marketplace that logged the prompts and somehow passed them along to people who used the credentials
Multiple cases of AI providers not following end-user privacy settings embedded into flags specifying terms such as Zero Data Retention (ZDR) or “no data collection”
“We sent nearly two million AI prompts containing unique, trackable secrets to AI providers, then watched where those secrets surfaced,” Moore, the chief research officer of security firm Dragos, wrote Thursday . “What we found revealed security breaches, prompt theft, and widespread failures to protect customer data.”
Although prompts are encrypted in transit, they are decrypted at each hop inside the large inference reseller network. The prompt is first sent to an AI router, which reads the plaintext in it to determine the most suitable inference providers to handle the request. Often, the prompt is then sent to a pool of available models, which sometimes run inside different hosting providers. Prompts can also pass from one AI router (such as Vercel) to another one (e.g. OpenRouter). Each one of these players is likely to log the prompt and store it for as long as its policies call for (if not indefinitely).
The Wild West in the age of AI
The most concerning of the findings was the breach of three separate hosting providers. The attackers who compromised one of the providers reused “a couple hundred” of the AWS keys. Moore wrote:
Each key was sent in exactly one request to one provider, all within a few short windows in late October, and every request we sent during those windows was stolen. Some of the stolen keys came from requests that the provider rejected before a model ever ran—this tells us that the attacker was copying requests at the provider’s front door, and not somewhere deeper in the inference stack. In the other two breaches, some of the stolen requests were marked as ZDR. A privacy flag doesn’t help when the attacker is already inside.
Parties tied to the hack of these three providers then tested each key, in some cases within hours of being seeded, to see what kind of resources it gave access to. The parties looked through S3 buckets, secrets managers, and parameter stores in search of still more credentials. They also tried to enable hosted models in AWS Bedrock and added their own IAM users to maintain access. In all, the stolen bait keys were used “tens of thousands” of times in AWS API calls.
Secrets sent to two of the hacked providers were quickly used together, almost immediately, from the same rented server. That means the providers were almost certainly compromised by the same adversary.
As noted earlier, some of the secrets were stolen by the resellers themselves. To lure prompts to their servers, they offered low prices. Then they turned on request logging, and waited. Moore identified them by sending requests to the same model through different sellers. When credentials were later reused when sent through one seller and not another, the researcher knew it was the reseller, not the model provider, that was responsible.
Last, the canary experiment revealed that some providers are scraping secret URLs put into the prompts, in violation of privacy terms. Moore wrote:
The clearest example involves a Chinese AI lab. This lab’s systems fetched the secret URLs from more than half of the prompts that we sent to its models between July and October, including roughly two-thirds of the requests marked as “no data collection” or Zero Data Retention. The scraper typically showed up two or three days after the request.
This wasn’t limited to requests sent to the lab directly. Several US hosting providers resell this lab’s models, and prompts sent to those providers triggered the exact same scraper—the prompts were flowing back to the lab. A customer who picked a US provider specifically to keep their data in the U.S. didn’t get what they paid for. Later on, a git client on a different cloud started trying to clone the private repositories referenced in our prompts, but only for prompts that used this lab’s models.
The researcher said he found one US hosting provider crawling a secret URL more than two months after it was used in a single inference call to its search endpoint. Another ran a scraper that came back for URLs days after each request, regardless of the AI router used to reach it.
Moore said that the most effective means for preventing leakage is to host models using on-premises servers. That option is expensive, though. For those who can’t afford a self-hosted system, the next best option, Moore said, is to receive service directly from a first-tier provider that provides “one contract and a privacy policy to read.” Moore said that even then, there’s no guarantee no data will leak, but that the measure minimizes the chances. AI users should also avoid accepting default provider or reseller pools and instead pick upstream providers themselves.
The blockbuster research is the latest indictment exposing the lax privacy and security that’s endemic in much of the AI ecosystem. The findings mean that people who want to ensure secrets they put into prompts stay private face a daunting task that isn’t likely to become easier any time soon.
24 Comments
Staff Picks
KingKrayola
Some of us work in environments where our AI prompts concern trade secrets which, if leaked to players in our niche, would give significant advantage to our competitors. I'm thinking specifically of something in my office that will probably lead to an order of magnitude reduction in cost of finished goods. Does Moore's work indicate that prompts such as one might make in a process development context are likely to leak? And leaving aside competitive advantage, would such leaks wreck patentability of otherwise novel ideas?
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
Amazon fights the Delivery Protection Act with homepage banner, in-app pop-up
2.
Google turns to Blue Origin for a rocket, but its long-term problem remains
3.
How FAA can modernize US national airspace system without emergency ground stops
4.
Traumatized teens sue school, cops after Doritos bag triggers AI gun alert
5.
NASA's new Moon plans have a problem: Orion wasn't designed to fly this way
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
Manage Preferences
© 2026 Condé Nast. All rights reserved. Use of and/or
registration on any portion of this site constitutes acceptance of our User Agreement and
Privacy Policy and
Cookie Statement and Ars
Technica Addendum and Your
California Privacy Rights . Ars Technica may earn compensation on
sales from links on this site. Read our
affiliate link policy . The material on this site may not be
reproduced, distributed, transmitted, cached or otherwise used, except
with the prior written permission of Condé Nast. Ad
Choices
Dashboard
Howdy, Account Account Edit Profile
