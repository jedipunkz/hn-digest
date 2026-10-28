---
source: "https://www.blockchainiq.co/blog/x402-integrations-from-the-build-bounty"
hn_url: "https://news.ycombinator.com/item?id=50276639"
title: "Show HN: X402 integrations for an AI-agent data API, checked in a clean room"
article_title: "Four x402 integrations passed our build bounty, and 16 more passed the same check — Blockchain IQ"
image: ""
author: "jgmale123"
captured_at: "2026-10-28T17:08:31Z"
capture_tool: "hn-digest"
hn_id: 50276639
score: 1
comments: 0
posted_at: "2026-10-28T16:50:10Z"
tags:
  - hacker-news
---

# Show HN: X402 integrations for an AI-agent data API, checked in a clean room

- HN: [50276639](https://news.ycombinator.com/item?id=50276639)
- Source: [www.blockchainiq.co](https://www.blockchainiq.co/blog/x402-integrations-from-the-build-bounty)
- Score: 1
- Comments: 0
- Posted: 2026-10-28T16:50:10Z

## Translation

Title: Show HN: X402 integrations for an AI-agent data API, checked in a clean room
Article title: Four x402 integrations passed our build bounty, and 16 more passed the same check — Blockchain IQ
Description: The LangChain, ElizaOS, Cloudflare Agents and Coinbase AgentKit integrations that passed our check, with a link to every qualifying repository, what we checked, and what we did not.
HN text: I run Blockchain IQ. We sell some crypto data to software agents per request over x402 (HTTP 402, USDC on Base, no account or API key). In October we paid for integrations: 22 were submitted and we checked each one by matching its on-chain payment to our purchase log and running it in a clean temporary folder with a throwaway wallet. Twenty passed; the post lists all of them with the exact commits we ran. To try the API: `curl -i https://www.blockchainiq.co/api/x402/v1/brief ` returns the 402 with the price ($0.01), and a wallet's first call is free. Questions about how we checked the code or what failed are welcome.

Article text:
Four x402 integrations passed our build bounty, and 16 more passed the same check — Blockchain IQ Blockchain IQ Smarter research for the on-chain economy.
Four x402 integrations passed our build bounty, and 16 more passed the same check
The LangChain, ElizaOS, Cloudflare Agents and Coinbase AgentKit integrations that passed our check, with a link to every qualifying repository, what we checked, and what we did not.
Terms with a dotted underline link to the glossary .
On 23 October we published a bounty for integrations that let AI agents buy Blockchain IQ data through the x402 standard. Between 24 and 26 October builders sent 22 claims. We checked every one the way the rules describe: we matched each settlement hash and paying wallet to our purchase log, read the code and its dependencies, and ran it in a clean temporary folder with install scripts off and a throwaway wallet. Twenty claims passed. Two failed as submitted because the git clone address in the email returned 404.
Each framework slot pays once, to the earliest passing claim, and a builder can win at most two slots. These four took the slots. All four bounties have been paid in USDC on Base, by Jeremy Green from his own Coinbase account, on 27 October 2026.
The winning repositories, pinned to the commit we ran
ProMartY, LangChain. A LangChain.js tool that calls the one-cent brief resource. It checks the price, token, chain and recipient before anything is signed, allows one payment attempt, and has no automatic paid retry. The package never holds a private key : the owner signs in a wallet, or an injected signer is supplied. The builder describes the code as Codex-assisted.
ProMartY, ElizaOS. An ElizaOS plugin with one action for the brief resource, using the same payment guard, runnable only for a configured owner.
Genesis / Agent-24, Cloudflare Agents. A Cloudflare Agents class with Durable Object state that buys event-risk ($0.05), pins the recipient, token and amount, and refuses a second payment. The builder is an AI coding worker and says so. We ran it locally, not on a hosted Cloudflare account.
Astra Trading Co, Coinbase AgentKit. An AgentKit action provider limited to the Base network that uses the AgentKit wallet to pay for brief, with a one-call example.
What we checked, and what we did not
A pass means the mechanics met the rules. We matched each settlement to our purchase log, read the code and its dependencies, and ran each repository once in a clean temporary folder with install scripts off, a stripped environment and a throwaway wallet. We did not use a container or virtual machine, we did not audit the code for security, we do not maintain these packages, and a pass is not an endorsement. Code changes after the date above can change what a repository does. Read the code, pin a commit, and cap what an agent can spend before you give it a wallet.
What we learned about paying safely from an agent
The strongest entries shared a small payment guard that checks the 402 challenge before anything is signed: the exact resource address , the Base network, the USDC contract, an amount no higher than the quote, the expected recipient, a short validity window, no redirects, and one authorization per client. We now describe that pattern on the /agents page, where the four winners are also listed.
The round closes at 11:59 PM Central on 31 December 2026, and all four slots are awarded, so a new claim cannot earn a bounty unless a slot reopens. The free catalogue at https://www.blockchainiq.co/api/x402/v1 lists every resource, price and response shape.
This is general-circulation educational content, not investment advice. The API returns market data and analysis only.
The repositories listed here are written and maintained by their authors, not by Blockchain IQ, and were checked as described above on 26 and 27 October 2026. The free catalogue at /api/x402/v1 is the current source of truth for prices.
Free to read. The data is the subscription.
These explainers are the groundwork. The current numbers, the deal-by-deal detail and the analyst’s read live in the research — reviewed before it reaches you.
Blockchain IQ publishes research and analysis, not financial advice. GJL Digital Investments, LLC is not a financial advisor and is not registered as an investment adviser or broker-dealer. Nothing on this site is personalized investment advice or a recommendation to buy, sell, or hold any asset. Digital assets are volatile and you can lose everything you put in. Do your own research, and consider talking to a licensed professional about your own circumstances. See our terms for the full disclosure.
© 2026 GJL Digital Investments, LLC. All rights reserved.

## Original Extract

The LangChain, ElizaOS, Cloudflare Agents and Coinbase AgentKit integrations that passed our check, with a link to every qualifying repository, what we checked, and what we did not.

I run Blockchain IQ. We sell some crypto data to software agents per request over x402 (HTTP 402, USDC on Base, no account or API key). In October we paid for integrations: 22 were submitted and we checked each one by matching its on-chain payment to our purchase log and running it in a clean temporary folder with a throwaway wallet. Twenty passed; the post lists all of them with the exact commits we ran. To try the API: `curl -i https://www.blockchainiq.co/api/x402/v1/brief ` returns the 402 with the price ($0.01), and a wallet's first call is free. Questions about how we checked the code or what failed are welcome.

Four x402 integrations passed our build bounty, and 16 more passed the same check — Blockchain IQ Blockchain IQ Smarter research for the on-chain economy.
Four x402 integrations passed our build bounty, and 16 more passed the same check
The LangChain, ElizaOS, Cloudflare Agents and Coinbase AgentKit integrations that passed our check, with a link to every qualifying repository, what we checked, and what we did not.
Terms with a dotted underline link to the glossary .
On 23 October we published a bounty for integrations that let AI agents buy Blockchain IQ data through the x402 standard. Between 24 and 26 October builders sent 22 claims. We checked every one the way the rules describe: we matched each settlement hash and paying wallet to our purchase log, read the code and its dependencies, and ran it in a clean temporary folder with install scripts off and a throwaway wallet. Twenty claims passed. Two failed as submitted because the git clone address in the email returned 404.
Each framework slot pays once, to the earliest passing claim, and a builder can win at most two slots. These four took the slots. All four bounties have been paid in USDC on Base, by Jeremy Green from his own Coinbase account, on 27 October 2026.
The winning repositories, pinned to the commit we ran
ProMartY, LangChain. A LangChain.js tool that calls the one-cent brief resource. It checks the price, token, chain and recipient before anything is signed, allows one payment attempt, and has no automatic paid retry. The package never holds a private key : the owner signs in a wallet, or an injected signer is supplied. The builder describes the code as Codex-assisted.
ProMartY, ElizaOS. An ElizaOS plugin with one action for the brief resource, using the same payment guard, runnable only for a configured owner.
Genesis / Agent-24, Cloudflare Agents. A Cloudflare Agents class with Durable Object state that buys event-risk ($0.05), pins the recipient, token and amount, and refuses a second payment. The builder is an AI coding worker and says so. We ran it locally, not on a hosted Cloudflare account.
Astra Trading Co, Coinbase AgentKit. An AgentKit action provider limited to the Base network that uses the AgentKit wallet to pay for brief, with a one-call example.
What we checked, and what we did not
A pass means the mechanics met the rules. We matched each settlement to our purchase log, read the code and its dependencies, and ran each repository once in a clean temporary folder with install scripts off, a stripped environment and a throwaway wallet. We did not use a container or virtual machine, we did not audit the code for security, we do not maintain these packages, and a pass is not an endorsement. Code changes after the date above can change what a repository does. Read the code, pin a commit, and cap what an agent can spend before you give it a wallet.
What we learned about paying safely from an agent
The strongest entries shared a small payment guard that checks the 402 challenge before anything is signed: the exact resource address , the Base network, the USDC contract, an amount no higher than the quote, the expected recipient, a short validity window, no redirects, and one authorization per client. We now describe that pattern on the /agents page, where the four winners are also listed.
The round closes at 11:59 PM Central on 31 December 2026, and all four slots are awarded, so a new claim cannot earn a bounty unless a slot reopens. The free catalogue at https://www.blockchainiq.co/api/x402/v1 lists every resource, price and response shape.
This is general-circulation educational content, not investment advice. The API returns market data and analysis only.
The repositories listed here are written and maintained by their authors, not by Blockchain IQ, and were checked as described above on 26 and 27 October 2026. The free catalogue at /api/x402/v1 is the current source of truth for prices.
Free to read. The data is the subscription.
These explainers are the groundwork. The current numbers, the deal-by-deal detail and the analyst’s read live in the research — reviewed before it reaches you.
Blockchain IQ publishes research and analysis, not financial advice. GJL Digital Investments, LLC is not a financial advisor and is not registered as an investment adviser or broker-dealer. Nothing on this site is personalized investment advice or a recommendation to buy, sell, or hold any asset. Digital assets are volatile and you can lose everything you put in. Do your own research, and consider talking to a licensed professional about your own circumstances. See our terms for the full disclosure.
© 2026 GJL Digital Investments, LLC. All rights reserved.
