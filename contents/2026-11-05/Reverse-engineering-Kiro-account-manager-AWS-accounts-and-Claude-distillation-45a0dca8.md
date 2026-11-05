---
source: "https://hansdekker.dev/writing/fraudulent-aws-accounts-claude-distillation/"
hn_url: "https://news.ycombinator.com/item?id=50383281"
title: "Reverse-engineering Kiro-account-manager: AWS accounts and Claude distillation"
article_title: "How fraudulent AWS accounts feed a Claude distillation market"
image: "https://hansdekker.dev/og/kiro-v2.png"
author: "ioshaan"
captured_at: "2026-11-05T19:35:44Z"
capture_tool: "hn-digest"
hn_id: 50383281
score: 1
comments: 1
posted_at: "2026-11-05T19:27:16Z"
tags:
  - hacker-news
---

# Reverse-engineering Kiro-account-manager: AWS accounts and Claude distillation

- HN: [50383281](https://news.ycombinator.com/item?id=50383281)
- Source: [hansdekker.dev](https://hansdekker.dev/writing/fraudulent-aws-accounts-claude-distillation/)
- Score: 1
- Comments: 1
- Posted: 2026-11-05T19:27:16Z

## Translation

Title: Reverse-engineering Kiro-account-manager: AWS accounts and Claude distillation
Article title: How fraudulent AWS accounts feed a Claude distillation market
Description: An open-source tool mass-creates fraudulent AWS accounts by forging Amazon

Article text:
← Hans Dekker How fraudulent AWS accounts feed a Claude distillation market: reverse-engineering Kiro-account-manager
A seller in a Chinese-language Telegram channel offers Claude access at six cents on the dollar. Behind prices like that is a supply chain of fraudulently created accounts, and in those channels I found one example of how those accounts can be sourced: an open-source tool that mass-creates free AWS Builder ID accounts by forging Amazon’s anti-bot fingerprint. I investigate how it works and what defenders can learn from it.
I reverse-engineer ‘Kiro-account-manager’, an open-source desktop tool that mass-creates free AWS Builder ID accounts that give access to Claude through Kiro IDE.
Account creation is automated without a browser: it forges the browser fingerprint Amazon’s bot checks expect.
It bundles the accounts behind one local proxy, creating high-volume free Claude API access.
The account type this tool produces is sold in Chinese-language Telegram channels at a fraction of Anthropic’s rate, openly marketed as distillation-capable.
I show how it works, how its randomized fingerprints can help defenders catch it, and the wider pattern it belongs to.
Why this matters to detection teams
For AI products like Kiro, an account’s free tier is backed by a frontier model: Kiro is Amazon’s IDE, and a model behind it is Anthropic’s Claude. This gives attackers a strong economic incentive to extract that value. Mass-created free accounts enable multiple types of abuse:
Bypassing usage limits: Pooling many accounts bypasses the per-account limits, for cheap, high-volume access.
Relays and discounted resale: That access is resold as an API in an underground market of relays , at a fraction of the official price.
Model distillation: High-volume access lets a buyer harvest a model’s answers and the reasoning it shows to train a model of their own.
Distillation lets a competitor acquire an AI lab’s capabilities at a fraction of the time and cost. Anthropic has reported distillation campaigns against Claude that ran through fraudulent accounts and resold access, and the US government has started treating it as a national-security problem .
Kiro-account-manager is an open source Electron app published under the GitHub account ‘chaogei’. Despite the name, it isn’t only a manager: you can run a full registration pipeline from the UI.
Register: reimplementing AWS’s signup as raw JSON
Creating an AWS Builder ID is meant to be done from a browser. Kiro-account-manager does it without one: it reimplements the whole signup flow as a scripted sequence of more than a dozen HTTP requests to AWS’s own endpoints.
The first request registers an OAuth client under the name of Amazon’s official CLI, Amazon Q Developer for command line, so its traffic looks like it came from that tool. The rest goes through the signup similarly: submit the email, create the identity, complete the SSO flow, exchange the token. The one step that happens outside AWS is the email OTP, which the tool reads back from a mailbox it controls (see next section).
// registrar.ts - step 1: register the OIDC client
const payload = {
clientName: 'Amazon Q Developer for command line', clientType: 'public', scopes: [ /* codewhisperer:* */]
}
await this.doPost(this.cfg.oidcBase + '/client/register', payload, headers)
Verify: scraping the OTP
Signup needs a verified email to receive the OTP, and the tool gets it two ways. It can make throwaway addresses on domains the operator owns, with a catch-all rule that forwards every address to a single inbox on a disposable-mail service (GPTmail or TempMail.Plus). It builds each of those addresses from fixed lists of first and last names (e.g. john.michael.smith ). Or it can use real mailboxes the operator supplies: Outlook accounts in bulk, or dot-variants of one Proton address.
The rest is automatic: set how many accounts to make and how many at once, then start. The tool runs the signup, reads the OTP, retries failures, and saves the ones that work.
To prevent automated account creation, Amazon loads a piece of client-side anti-bot code on the sign-in page called ‘FWCIM’. This code collects fingerprints of the device and behavior into a blob, encrypts it (XXTEA), and reports it to AWS.
The tool fabricates the fingerprint. It fills the same schema FWCIM collects, hardcoded in fingerprint.ts, by randomizing the behavioral values per account and drawing device values from fixed lists:
// fingerprint.ts - the forged blob
result.set('gpu', { vendor: identity.gpuVendor, model: identity.gpuModel, extensions: identity.webGLExts })
result.set('math', { tan: identity.mathTan, sin: identity.mathSin, cos: identity.mathCos })
result.set('canvas', { hash: canvasHash, emailHash: null, histogramBins: [...histogram] })
result.set('webDriver', false)
result.set('timeZone', 8) // hardcoded UTC+8
result.set('userAgent', identity.ua)
// ...
The tool packages this fingerprint the same way FWCIM does. To do so, it reads the encryption key, identifier, and version out of AWS’s app.js at runtime, with a hardcoded fallback. The key is stored as four numbers in a JavaScript array in a shuffled order, and the tool reorders them to get the real key:
// xxtea.ts (simplified) - read the key out of AWS's app.js and reorder its array
const js = await (await fetchFn('https://us-east-1.signin.aws/assets/js/app.js', { headers })).text()
// key stored as [a, "identifier", b, c, d]
const keyMatch = js.match(/var\s+\w+\s*=\s*\[(\d+),\s*"([A-Za-z0-9]+)",\s*(\d+),\s*(\d+),\s*(\d+)\]/)
const nums = [keyMatch[1], keyMatch[3], keyMatch[4], keyMatch[5]].map(Number)
const key = [nums[2], nums[0], nums[3], nums[1]] // real XXTEA key: those four numbers, reordered
const identifier = keyMatch[2]
const version = js.match(/FWCIM_VERSION\s*=\s*"(\d+\.\d+\.\d+)"/)[1]
Because the tool can read the key out of AWS’s own public script, the encryption only proves the blob arrived intact, not that a real browser made it. Anyone with the key can encrypt a fabricated one, so the fingerprint is untrusted input.
A fabricated fingerprint gets one account past the check. A thousand accounts made from one IP is suspicious too, so the tool supports proxies to avoid IP-based detection.
The operator configures a residential proxy, which is the address AWS sees. Because the residential proxies the tool supports (like bestproxy) reject mainland-China source IPs, and the tool is targeted at Chinese-language users, the operator can also configure an upstream proxy, which acts as a hop before connecting to the residential proxy.
A single free account has capped usage limits, so the value comes from pooling them at scale.
The tool runs a local reverse proxy that exposes the standard OpenAI/Anthropic API. Behind it sits the pool of farmed accounts. For each incoming request it picks an account, dresses the request up as the Kiro IDE, and forwards it to AWS’s CodeWhisperer endpoint.
Any tool built for the OpenAI or Anthropic API works against the relay too, so the pooled free accounts behave like one paid API key.
The accounts this tool makes have a name in these channels: kiro, after the IDE they sign into. Buyers’ messages show these are the same free AWS Builder ID accounts: they ask how to log into a kiro account “with the Builder ID”.
Across these channels, mentions of kiro climbed from a few hundred a month in spring 2026 to tens of thousands a month by midsummer.
Kiro moves through a layered market:
Farmers mass-create the accounts
Relay operators source them and pool them into shared key groups
These relays sell discounted Claude API access
Buyers downstream are promised distillation-capable API access, or even bulk-usage data
Relay sellers quote prices as a multiple of Anthropic’s API rate: 1× is the official price and kiro runs at 0.06× (Figure 1).
Distillation is a named product tier. Relay operators advertise their pools as distillation-capable and pitch them to enterprise buyers (Figure 1). Some run distillation-only pools with reserved capacity and a minimum spend, others advertise the reverse: “no distillation.”
Downstream, vendors sell Claude and GPT “interaction data” at 300 gigabytes a day (Figure 2).
Kiro is one of several AI products farmed this way. Farmers go after other third-party clients that offer free-tier access to frontier models in the same way. I also see indications of paid accounts being farmed with stolen credit card details.
Inconsistencies in the forgery
To evade detection the tool picks each fingerprint value at random from a pool, resulting in hundreds of millions of possible combinations. Each account claims to be Chrome on Windows in its User-Agent, with a version picked at random from 137 to 146. But some of those values are inconsistent with the real Chrome it claims to be, each in its own way: some are impossible on their own, some are only wrong in context, and some add randomization where a real Chrome does not.
One of the fields it randomizes is a math calculation: the sine, cosine and tangent of a single fixed number. I pulled FWCIM’s own collection out of AWS’s live script. It takes sin, cos and tan of one constant:
// AWS app.js (minified, deobfuscated) - FWCIM's real math collection
math: { sin: ""+Math.sin(CONSTANT), cos: ""+Math.cos(CONSTANT), tan: ""+Math.tan(CONSTANT) } // CONSTANT = -1e300
The tool picks this result randomly from a list of ten options:
// browser-identity.ts - ten "random" options for sin/cos/tan(-1e300)
const MATH_POOL = [
{ sin: '0.8178819121159085', cos: '-0.5753861119575491', tan: '-1.4214488238747245' }, // what Chrome (V8) returns for sin/cos/tan(-1e300), verified
{ sin: '0.8178819121159085', cos: '-0.5765775004286854', tan: '-1.4214488238747245' }, // impossible: sin²+cos² = 1.00137
{ sin: '0.8178819121159083', cos: '-0.5753861119575489', tan: '-1.4214488238747243' }, // last digits changed, not what the V8 build I tested returns
{ sin: '0.8178819121159084', cos: '-0.5765775004286855', tan: '-1.4214488238747244' }, // impossible: sin²+cos² = 1.00137
// ...10 rows total: 1 real, 4 arithmetically impossible, 5 last-digit perturbations
]
Chrome computes this in its own bundled math library rather than the operating system’s, so a given Chrome version should return the same triple on every machine. I ran -1e300 through V8 and only the first row matched.
Four of the rows are arithmetically impossible. A real sine and cosine of the same angle satisfy sin² + cos² = 1 , but in these rows the squares sum to 1.0013724361669656 .
Another field it picks randomly is the graphics card, which the browser reports as a string of text. On recent Chrome on Windows, that string includes the card’s hardware ID , a short hex code like (0x00002206) . The tool picks its card from a pool of twenty-two hardcoded strings, and twenty-one of them have that code missing. The Chrome it claims to be reports an RTX 3080 as ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 (0x00002206) Direct3D11 ...) , while the tool reports the same line without the code. So here again, nearly all account creation requests have a value the Windows Chrome they claim to be would not report.
// browser-identity.ts - GPU_CONFIGS (excerpt)
{ model: 'ANGLE (Intel, Intel(R) Iris(R) Xe Graphics (0x000046A6) Direct3D11 ...)' }, // 1/22 has the device ID
{ model: 'ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 Direct3D11 ...)' }, // no (0x....)
{ model: 'ANGLE (AMD, AMD Radeon RX 580 Direct3D11 ...)' }, // no (0x....)
The tool randomizes one more field, and this one should not vary at all. Chrome froze navigator.plugins years ago to a fixed list of five PDF entries in a fixed order, so a default Chrome lists the plugins in the same order on every machine. The tool has that exact list but shuffles it per account:
// browser-identity.ts - Chrome's frozen 5-entry plugin list, shuffled per account
const plugins = shuffle([...PLUGINS_POOL])
Fabricated behavior
The tool also fabricates behavior. FWCIM records mouse and keyboard activity, so the tool inv

[truncated]

## Original Extract

An open-source tool mass-creates fraudulent AWS accounts by forging Amazon

← Hans Dekker How fraudulent AWS accounts feed a Claude distillation market: reverse-engineering Kiro-account-manager
A seller in a Chinese-language Telegram channel offers Claude access at six cents on the dollar. Behind prices like that is a supply chain of fraudulently created accounts, and in those channels I found one example of how those accounts can be sourced: an open-source tool that mass-creates free AWS Builder ID accounts by forging Amazon’s anti-bot fingerprint. I investigate how it works and what defenders can learn from it.
I reverse-engineer ‘Kiro-account-manager’, an open-source desktop tool that mass-creates free AWS Builder ID accounts that give access to Claude through Kiro IDE.
Account creation is automated without a browser: it forges the browser fingerprint Amazon’s bot checks expect.
It bundles the accounts behind one local proxy, creating high-volume free Claude API access.
The account type this tool produces is sold in Chinese-language Telegram channels at a fraction of Anthropic’s rate, openly marketed as distillation-capable.
I show how it works, how its randomized fingerprints can help defenders catch it, and the wider pattern it belongs to.
Why this matters to detection teams
For AI products like Kiro, an account’s free tier is backed by a frontier model: Kiro is Amazon’s IDE, and a model behind it is Anthropic’s Claude. This gives attackers a strong economic incentive to extract that value. Mass-created free accounts enable multiple types of abuse:
Bypassing usage limits: Pooling many accounts bypasses the per-account limits, for cheap, high-volume access.
Relays and discounted resale: That access is resold as an API in an underground market of relays , at a fraction of the official price.
Model distillation: High-volume access lets a buyer harvest a model’s answers and the reasoning it shows to train a model of their own.
Distillation lets a competitor acquire an AI lab’s capabilities at a fraction of the time and cost. Anthropic has reported distillation campaigns against Claude that ran through fraudulent accounts and resold access, and the US government has started treating it as a national-security problem .
Kiro-account-manager is an open source Electron app published under the GitHub account ‘chaogei’. Despite the name, it isn’t only a manager: you can run a full registration pipeline from the UI.
Register: reimplementing AWS’s signup as raw JSON
Creating an AWS Builder ID is meant to be done from a browser. Kiro-account-manager does it without one: it reimplements the whole signup flow as a scripted sequence of more than a dozen HTTP requests to AWS’s own endpoints.
The first request registers an OAuth client under the name of Amazon’s official CLI, Amazon Q Developer for command line, so its traffic looks like it came from that tool. The rest goes through the signup similarly: submit the email, create the identity, complete the SSO flow, exchange the token. The one step that happens outside AWS is the email OTP, which the tool reads back from a mailbox it controls (see next section).
// registrar.ts - step 1: register the OIDC client
const payload = {
clientName: 'Amazon Q Developer for command line', clientType: 'public', scopes: [ /* codewhisperer:* */]
}
await this.doPost(this.cfg.oidcBase + '/client/register', payload, headers)
Verify: scraping the OTP
Signup needs a verified email to receive the OTP, and the tool gets it two ways. It can make throwaway addresses on domains the operator owns, with a catch-all rule that forwards every address to a single inbox on a disposable-mail service (GPTmail or TempMail.Plus). It builds each of those addresses from fixed lists of first and last names (e.g. john.michael.smith ). Or it can use real mailboxes the operator supplies: Outlook accounts in bulk, or dot-variants of one Proton address.
The rest is automatic: set how many accounts to make and how many at once, then start. The tool runs the signup, reads the OTP, retries failures, and saves the ones that work.
To prevent automated account creation, Amazon loads a piece of client-side anti-bot code on the sign-in page called ‘FWCIM’. This code collects fingerprints of the device and behavior into a blob, encrypts it (XXTEA), and reports it to AWS.
The tool fabricates the fingerprint. It fills the same schema FWCIM collects, hardcoded in fingerprint.ts, by randomizing the behavioral values per account and drawing device values from fixed lists:
// fingerprint.ts - the forged blob
result.set('gpu', { vendor: identity.gpuVendor, model: identity.gpuModel, extensions: identity.webGLExts })
result.set('math', { tan: identity.mathTan, sin: identity.mathSin, cos: identity.mathCos })
result.set('canvas', { hash: canvasHash, emailHash: null, histogramBins: [...histogram] })
result.set('webDriver', false)
result.set('timeZone', 8) // hardcoded UTC+8
result.set('userAgent', identity.ua)
// ...
The tool packages this fingerprint the same way FWCIM does. To do so, it reads the encryption key, identifier, and version out of AWS’s app.js at runtime, with a hardcoded fallback. The key is stored as four numbers in a JavaScript array in a shuffled order, and the tool reorders them to get the real key:
// xxtea.ts (simplified) - read the key out of AWS's app.js and reorder its array
const js = await (await fetchFn('https://us-east-1.signin.aws/assets/js/app.js', { headers })).text()
// key stored as [a, "identifier", b, c, d]
const keyMatch = js.match(/var\s+\w+\s*=\s*\[(\d+),\s*"([A-Za-z0-9]+)",\s*(\d+),\s*(\d+),\s*(\d+)\]/)
const nums = [keyMatch[1], keyMatch[3], keyMatch[4], keyMatch[5]].map(Number)
const key = [nums[2], nums[0], nums[3], nums[1]] // real XXTEA key: those four numbers, reordered
const identifier = keyMatch[2]
const version = js.match(/FWCIM_VERSION\s*=\s*"(\d+\.\d+\.\d+)"/)[1]
Because the tool can read the key out of AWS’s own public script, the encryption only proves the blob arrived intact, not that a real browser made it. Anyone with the key can encrypt a fabricated one, so the fingerprint is untrusted input.
A fabricated fingerprint gets one account past the check. A thousand accounts made from one IP is suspicious too, so the tool supports proxies to avoid IP-based detection.
The operator configures a residential proxy, which is the address AWS sees. Because the residential proxies the tool supports (like bestproxy) reject mainland-China source IPs, and the tool is targeted at Chinese-language users, the operator can also configure an upstream proxy, which acts as a hop before connecting to the residential proxy.
A single free account has capped usage limits, so the value comes from pooling them at scale.
The tool runs a local reverse proxy that exposes the standard OpenAI/Anthropic API. Behind it sits the pool of farmed accounts. For each incoming request it picks an account, dresses the request up as the Kiro IDE, and forwards it to AWS’s CodeWhisperer endpoint.
Any tool built for the OpenAI or Anthropic API works against the relay too, so the pooled free accounts behave like one paid API key.
The accounts this tool makes have a name in these channels: kiro, after the IDE they sign into. Buyers’ messages show these are the same free AWS Builder ID accounts: they ask how to log into a kiro account “with the Builder ID”.
Across these channels, mentions of kiro climbed from a few hundred a month in spring 2026 to tens of thousands a month by midsummer.
Kiro moves through a layered market:
Farmers mass-create the accounts
Relay operators source them and pool them into shared key groups
These relays sell discounted Claude API access
Buyers downstream are promised distillation-capable API access, or even bulk-usage data
Relay sellers quote prices as a multiple of Anthropic’s API rate: 1× is the official price and kiro runs at 0.06× (Figure 1).
Distillation is a named product tier. Relay operators advertise their pools as distillation-capable and pitch them to enterprise buyers (Figure 1). Some run distillation-only pools with reserved capacity and a minimum spend, others advertise the reverse: “no distillation.”
Downstream, vendors sell Claude and GPT “interaction data” at 300 gigabytes a day (Figure 2).
Kiro is one of several AI products farmed this way. Farmers go after other third-party clients that offer free-tier access to frontier models in the same way. I also see indications of paid accounts being farmed with stolen credit card details.
Inconsistencies in the forgery
To evade detection the tool picks each fingerprint value at random from a pool, resulting in hundreds of millions of possible combinations. Each account claims to be Chrome on Windows in its User-Agent, with a version picked at random from 137 to 146. But some of those values are inconsistent with the real Chrome it claims to be, each in its own way: some are impossible on their own, some are only wrong in context, and some add randomization where a real Chrome does not.
One of the fields it randomizes is a math calculation: the sine, cosine and tangent of a single fixed number. I pulled FWCIM’s own collection out of AWS’s live script. It takes sin, cos and tan of one constant:
// AWS app.js (minified, deobfuscated) - FWCIM's real math collection
math: { sin: ""+Math.sin(CONSTANT), cos: ""+Math.cos(CONSTANT), tan: ""+Math.tan(CONSTANT) } // CONSTANT = -1e300
The tool picks this result randomly from a list of ten options:
// browser-identity.ts - ten "random" options for sin/cos/tan(-1e300)
const MATH_POOL = [
{ sin: '0.8178819121159085', cos: '-0.5753861119575491', tan: '-1.4214488238747245' }, // what Chrome (V8) returns for sin/cos/tan(-1e300), verified
{ sin: '0.8178819121159085', cos: '-0.5765775004286854', tan: '-1.4214488238747245' }, // impossible: sin²+cos² = 1.00137
{ sin: '0.8178819121159083', cos: '-0.5753861119575489', tan: '-1.4214488238747243' }, // last digits changed, not what the V8 build I tested returns
{ sin: '0.8178819121159084', cos: '-0.5765775004286855', tan: '-1.4214488238747244' }, // impossible: sin²+cos² = 1.00137
// ...10 rows total: 1 real, 4 arithmetically impossible, 5 last-digit perturbations
]
Chrome computes this in its own bundled math library rather than the operating system’s, so a given Chrome version should return the same triple on every machine. I ran -1e300 through V8 and only the first row matched.
Four of the rows are arithmetically impossible. A real sine and cosine of the same angle satisfy sin² + cos² = 1 , but in these rows the squares sum to 1.0013724361669656 .
Another field it picks randomly is the graphics card, which the browser reports as a string of text. On recent Chrome on Windows, that string includes the card’s hardware ID , a short hex code like (0x00002206) . The tool picks its card from a pool of twenty-two hardcoded strings, and twenty-one of them have that code missing. The Chrome it claims to be reports an RTX 3080 as ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 (0x00002206) Direct3D11 ...) , while the tool reports the same line without the code. So here again, nearly all account creation requests have a value the Windows Chrome they claim to be would not report.
// browser-identity.ts - GPU_CONFIGS (excerpt)
{ model: 'ANGLE (Intel, Intel(R) Iris(R) Xe Graphics (0x000046A6) Direct3D11 ...)' }, // 1/22 has the device ID
{ model: 'ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 Direct3D11 ...)' }, // no (0x....)
{ model: 'ANGLE (AMD, AMD Radeon RX 580 Direct3D11 ...)' }, // no (0x....)
The tool randomizes one more field, and this one should not vary at all. Chrome froze navigator.plugins years ago to a fixed list of five PDF entries in a fixed order, so a default Chrome lists the plugins in the same order on every machine. The tool has that exact list but shuffles it per account:
// browser-identity.ts - Chrome's frozen 5-entry plugin list, shuffled per account
const plugins = shuffle([...PLUGINS_POOL])
Fabricated behavior
The tool also fabricates behavior. FWCIM records mouse and keyboard activity, so the tool inv

[truncated]
