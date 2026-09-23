---
source: "https://infertrail.com/blog/ruble-markup-ai-tokens/"
hn_url: "https://news.ycombinator.com/item?id=49809698"
title: "The ruble markup on AI tokens"
article_title: "The ruble markup on AI tokens | InferTrail"
image: ""
author: "whitef0x"
captured_at: "2026-09-23T00:02:48Z"
capture_tool: "hn-digest"
hn_id: 49809698
score: 1
comments: 0
posted_at: "2026-09-22T23:30:11Z"
tags:
  - hacker-news
---

# The ruble markup on AI tokens

- HN: [49809698](https://news.ycombinator.com/item?id=49809698)
- Source: [infertrail.com](https://infertrail.com/blog/ruble-markup-ai-tokens/)
- Score: 1
- Comments: 0
- Posted: 2026-09-22T23:30:11Z

## Translation

Title: The ruble markup on AI tokens
Article title: The ruble markup on AI tokens | InferTrail
Description: Why the price of an AI token in Russia can include more than model inference.

Article text:
InferTrail
Blog Research Security Analyze my traffic
Market structure
The ruble markup on AI tokens
How payment and region restrictions turned access to foreign AI models into a product of its own.
Arielle Jaffee · September 2026 · 8 minute read
On 23 September 2026, the Bank of Russia's official exchange rate was ₽84.0657 to the dollar. Anthropic lists Claude Sonnet 5 at $2 per million input tokens and $10 per million output tokens. Converted at that rate, that is roughly ₽168 and ₽841.
That is the provider list price converted into rubles. It is not the same as the retail price offered to a customer paying locally in Russia.
AITUNNEL's public page lists Sonnet 5 at ₽400 per million input tokens and ₽2,000 per million output tokens. That is roughly 2.38 times the converted provider list price on both sides of the meter.
Prices are a point-in-time comparison, not a margin calculation. Model prices change; resellers can change their tariffs; taxes, payment processing, foreign-exchange spread, support, and infrastructure all sit between the two numbers.
It is an access premium, not just a token markup
Russia's payment problem is often flattened into “Russia was removed from SWIFT.” That is incomplete. SWIFT is a financial messaging service, not a settlement system. It disconnected designated Russian entities in 2022 to comply with EU measures; the EU now describes restrictions affecting more than 100 Russian banks and related financial institutions. Neither statement means every Russian bank was disconnected from SWIFT.
The practical purchasing problem is broader. Visa said in 2022 that Russian-issued cards would no longer work outside Russia and foreign-issued Visa cards would no longer work within Russia. Mastercard suspended its Russian network services. Russia does not appear in Anthropic's current list of regions where it offers commercial API access.
For a developer or business trying to use a foreign model from Russia, the question is therefore not just: what does one million tokens cost? It is: how do I pay, establish an account that the provider supports, and keep a working route to the service?
A local-facing gateway can package some of that complexity into one account, one balance, and one compatible API endpoint. The scarce thing is not the token. It is the completed path from domestic payment to usable model access.
Public Russian-language services commonly advertise a ruble balance, domestic payment methods, invoices or closing documents, and access to multiple model providers behind a compatible interface. GPTunneL, for example, publicly advertises payment by SBP, bank cards, and invoices, and publishes a token price table.
The observable commercial chain looks like this:
customer
→ local payment, invoice, or stablecoin
→ local-facing gateway
→ foreign supplier, aggregator, or cloud account
→ model provider
→ compatible endpoint
Each arrow can introduce a real cost: acquiring payments, foreign exchange, tax, support, infrastructure, account administration, and the risk that an upstream provider changes its terms or route. A gateway may also provide a service the upstream provider's API price does not include: local language support, local paperwork, and one integration across multiple models.
That means the spread above is evidence of a local access premium. It is not evidence that any particular operator earns the full spread as profit.
One word, several different products
“Token” makes these offers look more comparable than they are. A formal gateway may publish separate input, output, cache-read, cache-write, and batch prices. Another seller may offer internal credits, nominal API balance, a subscription-derived allowance, or “unlimited” access subject to its own limits.
Compare prices only when the model version, direction of token use, cache treatment, billing unit, tax treatment, and service conditions match. A headline price by itself is not enough.
This matters particularly for unusually low offers. A public low price can be consistent with a number of arrangements: prepaid or subscription capacity, volume commitments, promotional credits, a different meter, model substitution, or a supplier relationship the buyer cannot see. It does not, by itself, establish credential theft, fraud, or sanctions evasion.
The same compatible API design that makes a gateway convenient also hides the provenance of a request. The downstream customer sees a model name, a token counter, a base URL, and an API key. They may not see the immediate supplier, billing currency, account relationship, actual routing decision, or whether the route changed after an outage.
That is a legitimate operational abstraction. It is also where investigation becomes hard.
A model provider may see a foreign account, cloud customer, or aggregator. The local customer may see a domestic invoice and a compatible key. The gateway owns the mapping between those two worlds. Without stronger identity, delegation, or reconciliation evidence, neither side can easily determine the provenance of every downstream request.
This is the security connection to InferTrail. Usage metadata can surface a credential or gateway whose behavior changes in ways worth investigating. It cannot establish authorization or provenance on its own. The right first output is verification_required , not an accusation.
Russian-facing gateways do not merely repackage model inference. They can package settlement, access, compatibility, support, and local commercial handling into a unit priced like a token.
That is why a converted provider price is not a local replacement price. When the direct payment and access path is missing or unreliable, the API becomes an intermediary product.
When payment rails break, an API can become a financial intermediary.
This article uses public pages and a point-in-time price check on 23 September 2026. It does not allege wrongdoing by named services or infer their upstream procurement from retail prices. Prices, policies, and supported regions can change without notice.
Anthropic: Claude Sonnet 5 price
AITUNNEL: published Sonnet 5 price
SWIFT: 2022 disconnection of designated Russian entities
Council of the EU: Russia sanctions overview
Visa: suspension of Russia operations
Mastercard: suspension of Russian operations

## Original Extract

Why the price of an AI token in Russia can include more than model inference.

InferTrail
Blog Research Security Analyze my traffic
Market structure
The ruble markup on AI tokens
How payment and region restrictions turned access to foreign AI models into a product of its own.
Arielle Jaffee · September 2026 · 8 minute read
On 23 September 2026, the Bank of Russia's official exchange rate was ₽84.0657 to the dollar. Anthropic lists Claude Sonnet 5 at $2 per million input tokens and $10 per million output tokens. Converted at that rate, that is roughly ₽168 and ₽841.
That is the provider list price converted into rubles. It is not the same as the retail price offered to a customer paying locally in Russia.
AITUNNEL's public page lists Sonnet 5 at ₽400 per million input tokens and ₽2,000 per million output tokens. That is roughly 2.38 times the converted provider list price on both sides of the meter.
Prices are a point-in-time comparison, not a margin calculation. Model prices change; resellers can change their tariffs; taxes, payment processing, foreign-exchange spread, support, and infrastructure all sit between the two numbers.
It is an access premium, not just a token markup
Russia's payment problem is often flattened into “Russia was removed from SWIFT.” That is incomplete. SWIFT is a financial messaging service, not a settlement system. It disconnected designated Russian entities in 2022 to comply with EU measures; the EU now describes restrictions affecting more than 100 Russian banks and related financial institutions. Neither statement means every Russian bank was disconnected from SWIFT.
The practical purchasing problem is broader. Visa said in 2022 that Russian-issued cards would no longer work outside Russia and foreign-issued Visa cards would no longer work within Russia. Mastercard suspended its Russian network services. Russia does not appear in Anthropic's current list of regions where it offers commercial API access.
For a developer or business trying to use a foreign model from Russia, the question is therefore not just: what does one million tokens cost? It is: how do I pay, establish an account that the provider supports, and keep a working route to the service?
A local-facing gateway can package some of that complexity into one account, one balance, and one compatible API endpoint. The scarce thing is not the token. It is the completed path from domestic payment to usable model access.
Public Russian-language services commonly advertise a ruble balance, domestic payment methods, invoices or closing documents, and access to multiple model providers behind a compatible interface. GPTunneL, for example, publicly advertises payment by SBP, bank cards, and invoices, and publishes a token price table.
The observable commercial chain looks like this:
customer
→ local payment, invoice, or stablecoin
→ local-facing gateway
→ foreign supplier, aggregator, or cloud account
→ model provider
→ compatible endpoint
Each arrow can introduce a real cost: acquiring payments, foreign exchange, tax, support, infrastructure, account administration, and the risk that an upstream provider changes its terms or route. A gateway may also provide a service the upstream provider's API price does not include: local language support, local paperwork, and one integration across multiple models.
That means the spread above is evidence of a local access premium. It is not evidence that any particular operator earns the full spread as profit.
One word, several different products
“Token” makes these offers look more comparable than they are. A formal gateway may publish separate input, output, cache-read, cache-write, and batch prices. Another seller may offer internal credits, nominal API balance, a subscription-derived allowance, or “unlimited” access subject to its own limits.
Compare prices only when the model version, direction of token use, cache treatment, billing unit, tax treatment, and service conditions match. A headline price by itself is not enough.
This matters particularly for unusually low offers. A public low price can be consistent with a number of arrangements: prepaid or subscription capacity, volume commitments, promotional credits, a different meter, model substitution, or a supplier relationship the buyer cannot see. It does not, by itself, establish credential theft, fraud, or sanctions evasion.
The same compatible API design that makes a gateway convenient also hides the provenance of a request. The downstream customer sees a model name, a token counter, a base URL, and an API key. They may not see the immediate supplier, billing currency, account relationship, actual routing decision, or whether the route changed after an outage.
That is a legitimate operational abstraction. It is also where investigation becomes hard.
A model provider may see a foreign account, cloud customer, or aggregator. The local customer may see a domestic invoice and a compatible key. The gateway owns the mapping between those two worlds. Without stronger identity, delegation, or reconciliation evidence, neither side can easily determine the provenance of every downstream request.
This is the security connection to InferTrail. Usage metadata can surface a credential or gateway whose behavior changes in ways worth investigating. It cannot establish authorization or provenance on its own. The right first output is verification_required , not an accusation.
Russian-facing gateways do not merely repackage model inference. They can package settlement, access, compatibility, support, and local commercial handling into a unit priced like a token.
That is why a converted provider price is not a local replacement price. When the direct payment and access path is missing or unreliable, the API becomes an intermediary product.
When payment rails break, an API can become a financial intermediary.
This article uses public pages and a point-in-time price check on 23 September 2026. It does not allege wrongdoing by named services or infer their upstream procurement from retail prices. Prices, policies, and supported regions can change without notice.
Anthropic: Claude Sonnet 5 price
AITUNNEL: published Sonnet 5 price
SWIFT: 2022 disconnection of designated Russian entities
Council of the EU: Russia sanctions overview
Visa: suspension of Russia operations
Mastercard: suspension of Russian operations
