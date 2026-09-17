---
source: "https://suganthan.com/blog/x402-pay-per-crawl/"
hn_url: "https://news.ycombinator.com/item?id=49734392"
title: "My website charged AI agents a penny per page. I watched Claude pay it"
article_title: "I made my website charge AI agents a penny per page. Then I watched Claude pay it. — Suganthan"
image: "https://suganthan.com/_astro/heroImage.cMhCdLQ6.jpg"
author: "gmays"
captured_at: "2026-09-17T00:09:44Z"
capture_tool: "hn-digest"
hn_id: 49734392
score: 3
comments: 1
posted_at: "2026-09-16T23:22:56Z"
tags:
  - hacker-news
---

# My website charged AI agents a penny per page. I watched Claude pay it

- HN: [49734392](https://news.ycombinator.com/item?id=49734392)
- Source: [suganthan.com](https://suganthan.com/blog/x402-pay-per-crawl/)
- Score: 3
- Comments: 1
- Posted: 2026-09-16T23:22:56Z

## Translation

Title: My website charged AI agents a penny per page. I watched Claude pay it
Article title: I made my website charge AI agents a penny per page. Then I watched Claude pay it. — Suganthan
Description: Google is testing paying publishers when their pages feed its AI answers, on terms only Google can see. My site runs the other version over x402. It returns HTTP 402 to AI agents and unlocks for a cent in USDC, with a public receipt for every crawl, which is the loop those Cloudflare wallet handles
[truncated]

Article text:
I made my website charge AI agents a penny per page. Then I watched Claude pay it. — Suganthan
Skip to content Suganthan ® AI SEO Research
Writing
Articles Research and longer reads Notes Short takes and observations Newsletter New writing in your inbox Tools About Work with Snippet AI SEO Research Writing Articles Notes Newsletter Tools About Work with Snippet Home / Blog I made my website charge AI agents a penny per page. Then I watched Claude pay it.
Google is testing paying publishers when their pages feed its AI answers, on terms only Google can see. My site runs the other version over x402. It returns HTTP 402 to AI agents and unlocks for a cent in USDC, with a public receipt for every crawl, which is the loop those Cloudflare wallet handles will run.
What that Cloudflare wallet handle is for
Watching Claude pay during a task
Should you start charging for your pages?
Moving from test tokens to real money
Google’s AI contribution pilot gives participating publishers a monthly earnings figure without a detailed explanation of how it was calculated.
My site charges one cent per page through x402. Five testnet payments settled on 15 September, including one made through a Claude Code hook during a task.
Paying for content use may reflect its value more closely, but a website owner can independently verify a paid fetch.
This is a working experiment using my own agents and test tokens. Ordinary search crawlers aren’t currently arriving with wallets to pay for access.
Google is paying some publishers when their content helps generate its AI answers.
Barry Schwartz covered it at Search Engine Roundtable on 14 September , following Digiday’s report that morning. Google confirmed the programme to Digiday. It’s called the AI contribution pilot .
Publishers who accept the terms get an earnings figure in Search Console each month. What they don’t get is a breakdown of how Google calculated it.
Google pays when it decides a page contributed significantly to an answer while it was being generated in AI Overviews, AI Mode or the Gemini app. If the page only gets linked after the answer has been written, it doesn’t qualify.
I’m not in the pilot. According to Digiday, Google has approached at least dozens of publishers, with smaller publishers showing more interest than the larger ones. One executive familiar with it described it as a black box. Barry had actually found the help page back in April , but it wasn’t clear then what the feature did.
I’ve been testing a different approach on my own site since August.
My page sets a price before an agent reads it. If the agent wants access, it pays.
This morning, 15 September, five payments went through. Each one settled on a blockchain and has a public transaction hash. One came from Claude Code, running on my machine and using a wallet I’d set up for it.
To be clear, all the payments so far have come from my own agents. No search crawler is turning up and paying me. The payments also use testnet USDC , which means the tokens have no real monetary value.
But the payment process works. An agent requests a page, checks the price, pays and receives the content. Cloudflare is building products around that same idea, and I wanted to understand how it works before it becomes another setting in a dashboard.
What that Cloudflare wallet handle is for
On 4 August, during Agents Week, Cloudflare announced Wallets and opened handle reservations.
Claiming a handle is free. I suspect plenty of people grabbed theirs the way we grab usernames : quickly, before someone else gets there, with a vague intention of reading the details later.
If you did that, this is the kind of thing those wallets are being built for.
Cloudflare’s design has two parts:
Account Wallets belong to people. You add funds, set spending limits and allocate money to your agents.
Virtual Wallets are for the agents. They spend within the permissions and limits you’ve set.
Your handle gives that account a recognisable identity.
As of 15 September, though, reserving a handle gets you a page displaying the name and a notification when Wallets becomes available. Cloudflare’s documentation explicitly says a reserved handle can’t yet send, receive or hold funds.
The other part is the Monetization Gateway , which will let sites charge agents for access to resources. Cloudflare announced it on 1 July , and access is still through the waitlist linked from that announcement.
So the products are announced, but you can’t yet use a reserved wallet handle to run this whole process.
My demo shows what that process looks like today using x402. The site charges one cent per page, my agent pays, and spending limits control what it can buy. It’s the same general arrangement Cloudflare describes for Virtual Wallets.
The arrangement between search engines and websites used to be fairly straightforward. Search engines crawled your content, showed it in their results and sent visitors back.
Publishers got traffic. Search engines got something useful to show their users.
AI answers have changed that arrangement. An agent can read your page, use the information in its response and answer the user without sending them to your site.
I’ve spent months looking at what these systems actually fetch . There’s plenty of crawling. The traffic coming back often doesn’t reflect it.
Publishers have responded in different ways, partly depending on how much negotiating power they have.
Smaller sites often block AI crawlers. I understand why. If someone is taking your content and you’re getting little in return, blocking them is a reasonable response. But it comes with a trade-off: restricting access can also reduce your visibility in training datasets and AI answers .
Large publishers can negotiate licensing deals. News Corp, the Financial Times and Reddit all have agreements with AI companies. According to Press Gazette’s reporting on 4 September , People Inc’s chief executive told investors that blocking crawlers through Cloudflare had helped bring AI companies to the negotiating table.
Google’s pilot offers another approach: paying publishers according to the value their content contributes to AI answers.
In that same Press Gazette piece, Cloudflare said it was moving its own default approach from pay per crawl towards pay per use, with pilots involving Ceramic.ai and You.com. Its reasoning was that fetching a page doesn’t prove the content was actually used.
I think that distinction matters.
Using a page to answer a question is a stronger indication of value than simply downloading it. But it’s also something the website owner can’t independently observe. Once my page has left my server, I can’t see whether it influenced an answer.
With pay per use, I’m relying on the platform to tell me what happened and what it was worth. With pay per crawl, I can check the request and the payment myself.
Google’s eligibility rules make this particularly clear. As Search Engine Land explains , contributing during answer generation can earn a payment. Being linked afterwards doesn’t. The publisher sees the total, without the calculation behind it.
Each approach leaves something unresolved.
Blocking earns nothing directly. Individual licensing agreements don’t scale to millions of small websites. Revenue-sharing arrangements leave the platform deciding how much to pay.
Putting a price into the request gives a site a standard way to sell access. An agent can accept that price and receive the content, with a receipt for the transaction. There’s no need to negotiate a separate commercial agreement with every buyer.
That’s the pay-per-crawl model I wanted to test.
Cloudflare already has a product called Pay Per Crawl , but it’s still in closed beta. On the paying side, bot operators need verification through Web Bot Auth, Stripe onboarding and programme approval.
My personal agent can’t simply join that programme.
So I built a version using the open x402 protocol , which Cloudflare is also using for its upcoming Monetization Gateway. My demo makes actual x402 payments independently of Cloudflare’s closed Pay Per Crawl programme.
HTTP has included a status code called 402 Payment Required since 1997. The specification reserved it for future use.
Twenty-nine years later, x402 gives it a practical purpose.
In my setup, the exchange has four steps:
The agent requests a page. The server responds with 402 Payment Required , including a PAYMENT-REQUIRED header that describes the price and payment requirements.
The agent signs an authorisation to transfer the exact amount in USDC. It then retries the request with a PAYMENT-SIGNATURE header.
A payment facilitator verifies the signature and settles the transaction on the blockchain. My demo uses Coinbase’s facilitator at x402.org .
The server returns the content, along with a PAYMENT-RESPONSE header containing the settlement receipt.
Here’s the offer from my site, decoded from a live 402 response this morning:
{
"x402Version" : 2 ,
"accepts" : [{
"scheme" : "exact" ,
"network" : "eip155:84532" ,
"amount" : "10000" ,
"asset" : "0x036CbD53842c5426634e7929541eC2318f3dCF7e" ,
"payTo" : "0xEdF2444D0259BBB8aC5094216D0148938F8308ff" ,
"maxTimeoutSeconds" : 300 ,
"extra" : { "name" : "USDC" , "version" : "2" }
}]
}
That offer asks for one cent in testnet USDC on Base Sepolia , paid to the burner wallet I’m using for the demo.
USDC uses six decimal places, so 10000 represents 0.01 USDC .
A compatible client can read those payment requirements and pay without creating an account with my site or asking me for an API key.
That’s a useful difference from robots.txt . A rule in robots.txt depends on a crawler choosing to respect it. Here, the server can withhold the protected content until payment succeeds.
The buyer wallet also held zero ETH , yet every settlement went through. The facilitator submitted the transactions and covered the blockchain transaction fees, usually called gas. I checked the balances twice.
For this setup, the agent only needed the testnet USDC it was spending.
There’s broader industry support behind the protocol too. The x402 Foundation became operational under the Linux Foundation on 14 July , with 40 members, including Visa, Mastercard, Google, AWS, Stripe and Cloudflare.
That’s useful context, but the actual exchange is still those four steps.
The demo is at paid.suganthan.com .
It runs on a Cloudflare Worker, using roughly 300 lines of Hono code and the x402 middleware. It’s been running since 8 August, and the public earnings page records every paid crawl since then.
The landing page is free and explains the demo.
/research/agentic-seo/ charges $0.01 per crawl.
/earnings/ shows the running total, timestamps, payer addresses and transaction hashes.
Request the article with curl and you’ll get this:
$ curl -i https://paid.suganthan.com/research/agentic-seo/
HTTP/2 402
payment-required: eyJ4NDAyVmVyc2lvbiI6MiwiZXJyb3Ii...
Open the same URL in a browser and you get a page explaining the charge. It still returns HTTP status 402 , but a human visitor gets something readable.
There’s also a preview button. That opens the article for free with a banner, so you can inspect the content as well as the payment screen.
I built another mode that isn’t enabled in the demo. Changing one variable makes the Worker charge only requests that identify themselves as known AI crawlers, including GPTBot, ClaudeBot and PerplexityBot. Human visitors can browse for free.
I tested that mode using spoofed user-agent strings, and it behaved as expected.
Of course, those strings can be faked. Checking a user agent is enough to demonstrate the behaviour, but it doesn’t prove who made the request. Cloudflare’s approach uses cryptographic verification of crawler identity to address that problem.
The buyer is a small Node script. Before it pays, it checks whether the request fits its budget.
A $0.25 daily allowance , recorded in a spending ledger on disk.
The script decodes the payment offer and checks both limits b

[truncated]

## Original Extract

Google is testing paying publishers when their pages feed its AI answers, on terms only Google can see. My site runs the other version over x402. It returns HTTP 402 to AI agents and unlocks for a cent in USDC, with a public receipt for every crawl, which is the loop those Cloudflare wallet handles
[truncated]

I made my website charge AI agents a penny per page. Then I watched Claude pay it. — Suganthan
Skip to content Suganthan ® AI SEO Research
Writing
Articles Research and longer reads Notes Short takes and observations Newsletter New writing in your inbox Tools About Work with Snippet AI SEO Research Writing Articles Notes Newsletter Tools About Work with Snippet Home / Blog I made my website charge AI agents a penny per page. Then I watched Claude pay it.
Google is testing paying publishers when their pages feed its AI answers, on terms only Google can see. My site runs the other version over x402. It returns HTTP 402 to AI agents and unlocks for a cent in USDC, with a public receipt for every crawl, which is the loop those Cloudflare wallet handles will run.
What that Cloudflare wallet handle is for
Watching Claude pay during a task
Should you start charging for your pages?
Moving from test tokens to real money
Google’s AI contribution pilot gives participating publishers a monthly earnings figure without a detailed explanation of how it was calculated.
My site charges one cent per page through x402. Five testnet payments settled on 15 September, including one made through a Claude Code hook during a task.
Paying for content use may reflect its value more closely, but a website owner can independently verify a paid fetch.
This is a working experiment using my own agents and test tokens. Ordinary search crawlers aren’t currently arriving with wallets to pay for access.
Google is paying some publishers when their content helps generate its AI answers.
Barry Schwartz covered it at Search Engine Roundtable on 14 September , following Digiday’s report that morning. Google confirmed the programme to Digiday. It’s called the AI contribution pilot .
Publishers who accept the terms get an earnings figure in Search Console each month. What they don’t get is a breakdown of how Google calculated it.
Google pays when it decides a page contributed significantly to an answer while it was being generated in AI Overviews, AI Mode or the Gemini app. If the page only gets linked after the answer has been written, it doesn’t qualify.
I’m not in the pilot. According to Digiday, Google has approached at least dozens of publishers, with smaller publishers showing more interest than the larger ones. One executive familiar with it described it as a black box. Barry had actually found the help page back in April , but it wasn’t clear then what the feature did.
I’ve been testing a different approach on my own site since August.
My page sets a price before an agent reads it. If the agent wants access, it pays.
This morning, 15 September, five payments went through. Each one settled on a blockchain and has a public transaction hash. One came from Claude Code, running on my machine and using a wallet I’d set up for it.
To be clear, all the payments so far have come from my own agents. No search crawler is turning up and paying me. The payments also use testnet USDC , which means the tokens have no real monetary value.
But the payment process works. An agent requests a page, checks the price, pays and receives the content. Cloudflare is building products around that same idea, and I wanted to understand how it works before it becomes another setting in a dashboard.
What that Cloudflare wallet handle is for
On 4 August, during Agents Week, Cloudflare announced Wallets and opened handle reservations.
Claiming a handle is free. I suspect plenty of people grabbed theirs the way we grab usernames : quickly, before someone else gets there, with a vague intention of reading the details later.
If you did that, this is the kind of thing those wallets are being built for.
Cloudflare’s design has two parts:
Account Wallets belong to people. You add funds, set spending limits and allocate money to your agents.
Virtual Wallets are for the agents. They spend within the permissions and limits you’ve set.
Your handle gives that account a recognisable identity.
As of 15 September, though, reserving a handle gets you a page displaying the name and a notification when Wallets becomes available. Cloudflare’s documentation explicitly says a reserved handle can’t yet send, receive or hold funds.
The other part is the Monetization Gateway , which will let sites charge agents for access to resources. Cloudflare announced it on 1 July , and access is still through the waitlist linked from that announcement.
So the products are announced, but you can’t yet use a reserved wallet handle to run this whole process.
My demo shows what that process looks like today using x402. The site charges one cent per page, my agent pays, and spending limits control what it can buy. It’s the same general arrangement Cloudflare describes for Virtual Wallets.
The arrangement between search engines and websites used to be fairly straightforward. Search engines crawled your content, showed it in their results and sent visitors back.
Publishers got traffic. Search engines got something useful to show their users.
AI answers have changed that arrangement. An agent can read your page, use the information in its response and answer the user without sending them to your site.
I’ve spent months looking at what these systems actually fetch . There’s plenty of crawling. The traffic coming back often doesn’t reflect it.
Publishers have responded in different ways, partly depending on how much negotiating power they have.
Smaller sites often block AI crawlers. I understand why. If someone is taking your content and you’re getting little in return, blocking them is a reasonable response. But it comes with a trade-off: restricting access can also reduce your visibility in training datasets and AI answers .
Large publishers can negotiate licensing deals. News Corp, the Financial Times and Reddit all have agreements with AI companies. According to Press Gazette’s reporting on 4 September , People Inc’s chief executive told investors that blocking crawlers through Cloudflare had helped bring AI companies to the negotiating table.
Google’s pilot offers another approach: paying publishers according to the value their content contributes to AI answers.
In that same Press Gazette piece, Cloudflare said it was moving its own default approach from pay per crawl towards pay per use, with pilots involving Ceramic.ai and You.com. Its reasoning was that fetching a page doesn’t prove the content was actually used.
I think that distinction matters.
Using a page to answer a question is a stronger indication of value than simply downloading it. But it’s also something the website owner can’t independently observe. Once my page has left my server, I can’t see whether it influenced an answer.
With pay per use, I’m relying on the platform to tell me what happened and what it was worth. With pay per crawl, I can check the request and the payment myself.
Google’s eligibility rules make this particularly clear. As Search Engine Land explains , contributing during answer generation can earn a payment. Being linked afterwards doesn’t. The publisher sees the total, without the calculation behind it.
Each approach leaves something unresolved.
Blocking earns nothing directly. Individual licensing agreements don’t scale to millions of small websites. Revenue-sharing arrangements leave the platform deciding how much to pay.
Putting a price into the request gives a site a standard way to sell access. An agent can accept that price and receive the content, with a receipt for the transaction. There’s no need to negotiate a separate commercial agreement with every buyer.
That’s the pay-per-crawl model I wanted to test.
Cloudflare already has a product called Pay Per Crawl , but it’s still in closed beta. On the paying side, bot operators need verification through Web Bot Auth, Stripe onboarding and programme approval.
My personal agent can’t simply join that programme.
So I built a version using the open x402 protocol , which Cloudflare is also using for its upcoming Monetization Gateway. My demo makes actual x402 payments independently of Cloudflare’s closed Pay Per Crawl programme.
HTTP has included a status code called 402 Payment Required since 1997. The specification reserved it for future use.
Twenty-nine years later, x402 gives it a practical purpose.
In my setup, the exchange has four steps:
The agent requests a page. The server responds with 402 Payment Required , including a PAYMENT-REQUIRED header that describes the price and payment requirements.
The agent signs an authorisation to transfer the exact amount in USDC. It then retries the request with a PAYMENT-SIGNATURE header.
A payment facilitator verifies the signature and settles the transaction on the blockchain. My demo uses Coinbase’s facilitator at x402.org .
The server returns the content, along with a PAYMENT-RESPONSE header containing the settlement receipt.
Here’s the offer from my site, decoded from a live 402 response this morning:
{
"x402Version" : 2 ,
"accepts" : [{
"scheme" : "exact" ,
"network" : "eip155:84532" ,
"amount" : "10000" ,
"asset" : "0x036CbD53842c5426634e7929541eC2318f3dCF7e" ,
"payTo" : "0xEdF2444D0259BBB8aC5094216D0148938F8308ff" ,
"maxTimeoutSeconds" : 300 ,
"extra" : { "name" : "USDC" , "version" : "2" }
}]
}
That offer asks for one cent in testnet USDC on Base Sepolia , paid to the burner wallet I’m using for the demo.
USDC uses six decimal places, so 10000 represents 0.01 USDC .
A compatible client can read those payment requirements and pay without creating an account with my site or asking me for an API key.
That’s a useful difference from robots.txt . A rule in robots.txt depends on a crawler choosing to respect it. Here, the server can withhold the protected content until payment succeeds.
The buyer wallet also held zero ETH , yet every settlement went through. The facilitator submitted the transactions and covered the blockchain transaction fees, usually called gas. I checked the balances twice.
For this setup, the agent only needed the testnet USDC it was spending.
There’s broader industry support behind the protocol too. The x402 Foundation became operational under the Linux Foundation on 14 July , with 40 members, including Visa, Mastercard, Google, AWS, Stripe and Cloudflare.
That’s useful context, but the actual exchange is still those four steps.
The demo is at paid.suganthan.com .
It runs on a Cloudflare Worker, using roughly 300 lines of Hono code and the x402 middleware. It’s been running since 8 August, and the public earnings page records every paid crawl since then.
The landing page is free and explains the demo.
/research/agentic-seo/ charges $0.01 per crawl.
/earnings/ shows the running total, timestamps, payer addresses and transaction hashes.
Request the article with curl and you’ll get this:
$ curl -i https://paid.suganthan.com/research/agentic-seo/
HTTP/2 402
payment-required: eyJ4NDAyVmVyc2lvbiI6MiwiZXJyb3Ii...
Open the same URL in a browser and you get a page explaining the charge. It still returns HTTP status 402 , but a human visitor gets something readable.
There’s also a preview button. That opens the article for free with a banner, so you can inspect the content as well as the payment screen.
I built another mode that isn’t enabled in the demo. Changing one variable makes the Worker charge only requests that identify themselves as known AI crawlers, including GPTBot, ClaudeBot and PerplexityBot. Human visitors can browse for free.
I tested that mode using spoofed user-agent strings, and it behaved as expected.
Of course, those strings can be faked. Checking a user agent is enough to demonstrate the behaviour, but it doesn’t prove who made the request. Cloudflare’s approach uses cryptographic verification of crawler identity to address that problem.
The buyer is a small Node script. Before it pays, it checks whether the request fits its budget.
A $0.25 daily allowance , recorded in a spending ledger on disk.
The script decodes the payment offer and checks both limits b

[truncated]
