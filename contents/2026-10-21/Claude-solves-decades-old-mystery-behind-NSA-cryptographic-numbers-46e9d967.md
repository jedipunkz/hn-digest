---
source: "https://cyberinsider.com/claude-solves-decades-old-mystery-behind-nsa-cryptographic-numbers/"
hn_url: "https://news.ycombinator.com/item?id=50174641"
title: "Claude solves decades-old mystery behind NSA cryptographic numbers"
article_title: "Claude solves decades-old mystery behind NSA cryptographic numbers | CyberInsider"
image: "https://cyberinsider.com/wp-content/uploads/2026/10/Claude-solves-decades-old-mystery-behind-NSA-cryptographic-numbers.jpeg"
author: "devonnull"
captured_at: "2026-10-21T01:20:02Z"
capture_tool: "hn-digest"
hn_id: 50174641
score: 4
comments: 1
posted_at: "2026-10-21T00:40:48Z"
tags:
  - hacker-news
---

# Claude solves decades-old mystery behind NSA cryptographic numbers

- HN: [50174641](https://news.ycombinator.com/item?id=50174641)
- Source: [cyberinsider.com](https://cyberinsider.com/claude-solves-decades-old-mystery-behind-nsa-cryptographic-numbers/)
- Score: 4
- Comments: 1
- Posted: 2026-10-21T00:40:48Z

## Translation

Title: Claude solves decades-old mystery behind NSA cryptographic numbers
Article title: Claude solves decades-old mystery behind NSA cryptographic numbers | CyberInsider
Description: Claude AI has reconstructed the origins of unexplained numbers used in widely adopted security standards, and supplied by the NSA.

Article text:
Skip to after header navigation
Reliable cybersecurity news and resources
Search site Submit search
General
Password
Best Password Managers
Best Windows Password Manager
Comparisons
Bitwarden vs Proton Pass
Identity
Best Identity Theft Protection Services
VPN
VPN Reviews
NordVPN Review
Best VPN Services
Best VPN for iPhone
VPN Comparisons
NordVPN vs Surfshark
Data Removal
Best Data Removal Services
Comparisons
Incogni vs DeleteMe
Signal fixes critical flaw that could keep audio active after hangup
WhatsApp told its AI chatbot not to mention privacy regulators
EFF brings Privacy Badger tracker blocker to Safari on macOS
Yamato says 5.3 million records stolen in payment service breach
CyberInsider covers the latest news in the cybersecurity and data privacy world. In addition to news, we also publish in-depth guides and resources.
See our Mission >
Password
Best Password Managers
Best Windows Password Manager
Comparisons
Bitwarden vs Proton Pass
Identity
Best Identity Theft Protection Services
VPN
VPN Reviews
NordVPN Review
Best VPN Services
Best VPN for iPhone
VPN Comparisons
NordVPN vs Surfshark
Data Removal
Best Data Removal Services
Comparisons
Incogni vs DeleteMe
Claude solves decades-old mystery behind NSA cryptographic numbers
X LinkedIn Reddit Facebook Share
Anthropic says its Claude AI has reconstructed the origins of unexplained numbers used in widely adopted security standards, tracing them to a conventional random-number generator initialized with dates and times from the 1990s.
The discovery helps answer a question that has lingered for decades: whether the U.S. National Security Agency (NSA) deliberately chose those numbers to introduce hidden weaknesses.
In a report published October 16, Anthropic said an internal Claude research model identified how the seeds for five standardized mathematical curves were generated in under half an hour. Earlier investigations had followed an incorrect lead suggesting the numbers came from humorous phrases.
The curves underpin implementations of the Elliptic Curve Digital Signature Algorithm (ECDSA), which helps computers verify authenticity and detect tampering. The U.S. National Institute of Standards and Technology (NIST) published them in its digital signature standard in 2000.
Although the NSA supplied the seeds used to derive the curves, their origins were unexplained. Researchers could check whether a seed produced particular mathematical parameters, but could not explain why that starting number was chosen. That gap fueled speculation about possible backdoors.
According to Anthropic, Claude approached the mystery by considering how software from that era would have generated the numbers.
It tested recognizable patterns left by older random-number generators, then used mathematical techniques to reconstruct their internal state. The results pointed to lrand48, a generator available in C programming libraries, initialized with Unix timestamps, numbers representing dates and times.
For the five NIST curves, the recovered timestamps fall between May 1997 and April 1998. For example, the P-256 seed was reproduced using August 16, 1997, at 01:52:34 UTC and the 1,842nd candidate generated. Anthropic published code allowing others to reproduce the results.
Cryptographer Filippo Valsorda, who organized a bounty for solving the mystery, independently reimplemented the generation algorithm in Go and verified the finding. He confirmed that the Anthropic team claimed the prize and chose a charitable payout of $36,864 to the Internet Security Research Group, which operates Let’s Encrypt. His verification supports the mathematical result; the account of Claude’s discovery process and timing comes from Anthropic.
Public numbers, not secret keys
A predictable generator would be dangerous if used to create private keys. Here, however, it produced starting material for mathematical parameters that were always public.
Reconstructing their origins does not expose users’ private keys, enable forged signatures, or unlock encrypted communications. Anthropic says the findings instead strengthen confidence that the curves were selected without hidden weaknesses.
[fact-checking!] NSA’s darkest cryptographic secret of the 1990s: it seeded its random numbers like a first-year student. Claude tested the decades-old theory of some cryptographers that the NSA secretly planted weak NIST elliptic curves behind unexplained seeds. It was a… pic.twitter.com/WbhdfGK5jr
Some details remain unresolved. The generation process skipped candidate curves that appear to meet the relevant requirements, and the researchers do not know why. They also could not explain one seed from a separate family of curves or the origins of several standardized starting points used in signature calculations.
The discovery substantially narrows the uncertainty surrounding these standards, leaving only a few gaps in their history.
If you liked this article, be sure to follow us on X/Twitter and also LinkedIn for more exclusive content.
X LinkedIn Facebook Reddit Share
More from CyberInsider
Signal fixes critical flaw that could keep audio active after hangup
WhatsApp told its AI chatbot not to mention privacy regulators
EFF brings Privacy Badger tracker blocker to Safari on macOS
Yamato says 5.3 million records stolen in payment service breach
AOMEI update servers hijacked in brief bursts to spread Windows malware
ETS data breach exposes California students’ personal information
Bill specializes in explaining complex technical topics to a non-technical audience. In his 30+ year career, he has covered many of the technological advances that shape our lives. Today, Bill uses those skills to help people protect their privacy and security against the ever-growing assaults on both.
In the health-field, in the not-too-distant history, one could still use the term “lay” audience. UK had moved on to “plain” language audience.
Your email address will not be published. Required fields are marked *
Signal fixes critical flaw that could keep audio active after hangup
WhatsApp told its AI chatbot not to mention privacy regulators
EFF brings Privacy Badger tracker blocker to Safari on macOS
Yamato says 5.3 million records stolen in payment service breach
AOMEI update servers hijacked in brief bursts to spread Windows malware
ETS data breach exposes California students’ personal information
Claude solves decades-old mystery behind NSA cryptographic numbers
Court filing reveals the user data Telegram is handing over to US authorities
Burger King Japan data breach exposes 11.18 million app member records
Brave adds end-to-end encryption to chats with AI assistant Leo
Copyright © 2026 · CyberInsider.com · All Rights Reserved · Privacy Policy · Terms of Use · Contact

## Original Extract

Claude AI has reconstructed the origins of unexplained numbers used in widely adopted security standards, and supplied by the NSA.

Skip to after header navigation
Reliable cybersecurity news and resources
Search site Submit search
General
Password
Best Password Managers
Best Windows Password Manager
Comparisons
Bitwarden vs Proton Pass
Identity
Best Identity Theft Protection Services
VPN
VPN Reviews
NordVPN Review
Best VPN Services
Best VPN for iPhone
VPN Comparisons
NordVPN vs Surfshark
Data Removal
Best Data Removal Services
Comparisons
Incogni vs DeleteMe
Signal fixes critical flaw that could keep audio active after hangup
WhatsApp told its AI chatbot not to mention privacy regulators
EFF brings Privacy Badger tracker blocker to Safari on macOS
Yamato says 5.3 million records stolen in payment service breach
CyberInsider covers the latest news in the cybersecurity and data privacy world. In addition to news, we also publish in-depth guides and resources.
See our Mission >
Password
Best Password Managers
Best Windows Password Manager
Comparisons
Bitwarden vs Proton Pass
Identity
Best Identity Theft Protection Services
VPN
VPN Reviews
NordVPN Review
Best VPN Services
Best VPN for iPhone
VPN Comparisons
NordVPN vs Surfshark
Data Removal
Best Data Removal Services
Comparisons
Incogni vs DeleteMe
Claude solves decades-old mystery behind NSA cryptographic numbers
X LinkedIn Reddit Facebook Share
Anthropic says its Claude AI has reconstructed the origins of unexplained numbers used in widely adopted security standards, tracing them to a conventional random-number generator initialized with dates and times from the 1990s.
The discovery helps answer a question that has lingered for decades: whether the U.S. National Security Agency (NSA) deliberately chose those numbers to introduce hidden weaknesses.
In a report published October 16, Anthropic said an internal Claude research model identified how the seeds for five standardized mathematical curves were generated in under half an hour. Earlier investigations had followed an incorrect lead suggesting the numbers came from humorous phrases.
The curves underpin implementations of the Elliptic Curve Digital Signature Algorithm (ECDSA), which helps computers verify authenticity and detect tampering. The U.S. National Institute of Standards and Technology (NIST) published them in its digital signature standard in 2000.
Although the NSA supplied the seeds used to derive the curves, their origins were unexplained. Researchers could check whether a seed produced particular mathematical parameters, but could not explain why that starting number was chosen. That gap fueled speculation about possible backdoors.
According to Anthropic, Claude approached the mystery by considering how software from that era would have generated the numbers.
It tested recognizable patterns left by older random-number generators, then used mathematical techniques to reconstruct their internal state. The results pointed to lrand48, a generator available in C programming libraries, initialized with Unix timestamps, numbers representing dates and times.
For the five NIST curves, the recovered timestamps fall between May 1997 and April 1998. For example, the P-256 seed was reproduced using August 16, 1997, at 01:52:34 UTC and the 1,842nd candidate generated. Anthropic published code allowing others to reproduce the results.
Cryptographer Filippo Valsorda, who organized a bounty for solving the mystery, independently reimplemented the generation algorithm in Go and verified the finding. He confirmed that the Anthropic team claimed the prize and chose a charitable payout of $36,864 to the Internet Security Research Group, which operates Let’s Encrypt. His verification supports the mathematical result; the account of Claude’s discovery process and timing comes from Anthropic.
Public numbers, not secret keys
A predictable generator would be dangerous if used to create private keys. Here, however, it produced starting material for mathematical parameters that were always public.
Reconstructing their origins does not expose users’ private keys, enable forged signatures, or unlock encrypted communications. Anthropic says the findings instead strengthen confidence that the curves were selected without hidden weaknesses.
[fact-checking!] NSA’s darkest cryptographic secret of the 1990s: it seeded its random numbers like a first-year student. Claude tested the decades-old theory of some cryptographers that the NSA secretly planted weak NIST elliptic curves behind unexplained seeds. It was a… pic.twitter.com/WbhdfGK5jr
Some details remain unresolved. The generation process skipped candidate curves that appear to meet the relevant requirements, and the researchers do not know why. They also could not explain one seed from a separate family of curves or the origins of several standardized starting points used in signature calculations.
The discovery substantially narrows the uncertainty surrounding these standards, leaving only a few gaps in their history.
If you liked this article, be sure to follow us on X/Twitter and also LinkedIn for more exclusive content.
X LinkedIn Facebook Reddit Share
More from CyberInsider
Signal fixes critical flaw that could keep audio active after hangup
WhatsApp told its AI chatbot not to mention privacy regulators
EFF brings Privacy Badger tracker blocker to Safari on macOS
Yamato says 5.3 million records stolen in payment service breach
AOMEI update servers hijacked in brief bursts to spread Windows malware
ETS data breach exposes California students’ personal information
Bill specializes in explaining complex technical topics to a non-technical audience. In his 30+ year career, he has covered many of the technological advances that shape our lives. Today, Bill uses those skills to help people protect their privacy and security against the ever-growing assaults on both.
In the health-field, in the not-too-distant history, one could still use the term “lay” audience. UK had moved on to “plain” language audience.
Your email address will not be published. Required fields are marked *
Signal fixes critical flaw that could keep audio active after hangup
WhatsApp told its AI chatbot not to mention privacy regulators
EFF brings Privacy Badger tracker blocker to Safari on macOS
Yamato says 5.3 million records stolen in payment service breach
AOMEI update servers hijacked in brief bursts to spread Windows malware
ETS data breach exposes California students’ personal information
Claude solves decades-old mystery behind NSA cryptographic numbers
Court filing reveals the user data Telegram is handing over to US authorities
Burger King Japan data breach exposes 11.18 million app member records
Brave adds end-to-end encryption to chats with AI assistant Leo
Copyright © 2026 · CyberInsider.com · All Rights Reserved · Privacy Policy · Terms of Use · Contact
