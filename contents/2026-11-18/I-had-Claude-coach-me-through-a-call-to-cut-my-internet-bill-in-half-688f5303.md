---
source: "https://patrick-reagan.com/work/earpiece"
hn_url: "https://news.ycombinator.com/item?id=50547347"
title: "I had Claude coach me through a call to cut my internet bill in half"
article_title: "Earpiece: A Live Call Coach | Patrick Reagan"
image: "https://patrick-reagan.com/assets/images/og/call-coach.jpg"
author: "reagent"
captured_at: "2026-11-18T16:14:19Z"
capture_tool: "hn-digest"
hn_id: 50547347
score: 2
comments: 1
posted_at: "2026-11-18T15:18:34Z"
tags:
  - hacker-news
---

# I had Claude coach me through a call to cut my internet bill in half

- HN: [50547347](https://news.ycombinator.com/item?id=50547347)
- Source: [patrick-reagan.com](https://patrick-reagan.com/work/earpiece)
- Score: 2
- Comments: 1
- Posted: 2026-11-18T15:18:34Z

## Translation

Title: I had Claude coach me through a call to cut my internet bill in half
Article title: Earpiece: A Live Call Coach | Patrick Reagan
Description: A live "coach" on an actual call that cut my internet bill in half

Article text:
pr Patrick Reagan software engineer Work Services Lab Articles & Talks Résumé Get in touch Work / Earpiece: A Live Call Coach
A live "coach" on an actual call that cut my internet bill in half
Getting Advice During the Call
I’ve been using Claude for a lot more than writing software. It helped me with a property tax appeal , with college planning and with lowering my cell phone bill. That got me wondering what other everyday chores it could take off my plate.
Internet service looked like a good candidate. Prices for the same plan vary a lot depending on who’s asking, and long-time customers tend to pay the most. I’d been with my provider for years at its top tier, and a competitor was offering the same speed at my address for half of what I paid. I wanted to see if AI could help me get that price without switching.
My first idea was to hand the most tedious part of the call to software. I set up a Twilio number and wrote a small app to dial the provider, work through the phone menu with keypad presses and wait on hold. When a person picked up, it would ring my cell and connect me so I could negotiate (a lo-fi version of FastCustomer ):
Bot dials the provider
Walks the menu keypad presses
Waits on hold for a person
Rings my cell and connects me
WHERE IT STALLED
The original plan: software waits on hold, and I only pick up once a person is on the line.
The menu took a lot of trial and error and didn’t recognize me as a customer, because the call came from the Twilio number. I lost interest in fiddling with it and set it aside, resolving to just make the call myself.
Before making the call, I thought to just ask Claude: “How can we win this negotiation?”
We worked through it in steps:
Pull my current bill and how long I’d been a customer.
Check what competitors offered at my address.
Read what other customers said worked on retention calls, from the provider’s own forums and a few blog posts.
From that, Claude wrote a pitch to lead with:
I'm not asking you to give me the new-customer promo. I'm asking what you can do to get me to [half] a month for gig service, because that's what I can get elsewhere.
We also reviewed a list of likely counteroffers with a response for each, and agreed not to ask for the new-customer promotion by name, since that was the most common reason other customers got turned down.
Getting Advice During the Call
I could have made the call with the pitch and a list of counters. But calls rarely follow the script, and I wanted advice on what the rep actually said while I was still on the phone.
So I flipped the setup around. Instead of a bot dialing and handing the call to me, I dialed from my cell through the Twilio number, and the app used Twilio’s live transcription to transcribe both sides as we talked. A Claude Code session in another project followed that live transcript and coached me in my terminal:
Me on my cell
Twilio number carries the call
The rep at the provider
The app logs each line
Claude Code reads the log live
My terminal advice in, questions out
BOTH SIDES TRANSCRIBED
LIVE TRANSCRIPT
THE CALL
How the call was set up. I talked to the rep through the Twilio number, the app logged both sides, and Claude Code read that log and coached me in my terminal.
Calling from my own phone meant the menu recognized me as a customer, so I skipped most of the menu steps my app had struggled with on the first attempt.
I made the call on October 26 th . It took about ten and a half minutes, and the rep offered the price I asked for about 75 seconds after my pitch. The new price started that day, and I accepted it from an emailed link before I hung up.
Rep Hi, thank you so much for calling [provider] .
Claude An agent picked up. Get their name, then give the opening line.
Me I'm a [long-time] customer, [top tier] . I'm paying [current price] a month for gig internet. A competitor has the same gig speed at my address for [half] , and your own site advertises [half] for new customers. I'd rather stay. What can you do to get me to [half] ?
Claude Good delivery. Now stay quiet and let them answer. I'll have counters ready.
Rep I see the 1 gig for [half] , with autopay and paperless active, and a 5-year price guarantee.
Claude [half] for 1 Gig with a 5-year price guarantee. That's your target. Accept it, but confirm these first: is it the total, is the gateway included, and when does it start.
Me Is [half] the total on the bill, or are taxes and fees extra?
Me, typed to Claude what is gateway? i have my own modem and router
Claude The gateway is the provider's rented modem and router in one box. Since you own yours, skip the equipment question.
Rep Your total is [half] flat after fees and taxes.
Me When does that start? Mid billing cycle, or at my next one?
Me Is there any contract or cancellation fee?
Rep No contract, no cancellation fees. The plan is unlimited, with no modem rental fees. If you want to cancel an hour after the call, you can.
Claude No contract, no cancellation fee, unlimited data, no modem rental. Everything checks out.
Rep I'm sending you the approval now. It'll be in your email in a couple of seconds.
Me, typed to Claude just saved a full webpage to downloads
Claude The terms match: [half] a month with taxes included, the 5-year guarantee, no contract. Click I agree.
Rep Your account is also eligible for a free mobile line for the first 12 months.
Claude That's an upsell. Decline politely: "I'm all set with my current carrier."
Me We'll look into it. Thanks for your help.
The next bill estimate, on October 28 th , showed the new price. The first full bill arrives in early December.
Working on something interesting?
I read every message and reply personally.

## Original Extract

A live "coach" on an actual call that cut my internet bill in half

pr Patrick Reagan software engineer Work Services Lab Articles & Talks Résumé Get in touch Work / Earpiece: A Live Call Coach
A live "coach" on an actual call that cut my internet bill in half
Getting Advice During the Call
I’ve been using Claude for a lot more than writing software. It helped me with a property tax appeal , with college planning and with lowering my cell phone bill. That got me wondering what other everyday chores it could take off my plate.
Internet service looked like a good candidate. Prices for the same plan vary a lot depending on who’s asking, and long-time customers tend to pay the most. I’d been with my provider for years at its top tier, and a competitor was offering the same speed at my address for half of what I paid. I wanted to see if AI could help me get that price without switching.
My first idea was to hand the most tedious part of the call to software. I set up a Twilio number and wrote a small app to dial the provider, work through the phone menu with keypad presses and wait on hold. When a person picked up, it would ring my cell and connect me so I could negotiate (a lo-fi version of FastCustomer ):
Bot dials the provider
Walks the menu keypad presses
Waits on hold for a person
Rings my cell and connects me
WHERE IT STALLED
The original plan: software waits on hold, and I only pick up once a person is on the line.
The menu took a lot of trial and error and didn’t recognize me as a customer, because the call came from the Twilio number. I lost interest in fiddling with it and set it aside, resolving to just make the call myself.
Before making the call, I thought to just ask Claude: “How can we win this negotiation?”
We worked through it in steps:
Pull my current bill and how long I’d been a customer.
Check what competitors offered at my address.
Read what other customers said worked on retention calls, from the provider’s own forums and a few blog posts.
From that, Claude wrote a pitch to lead with:
I'm not asking you to give me the new-customer promo. I'm asking what you can do to get me to [half] a month for gig service, because that's what I can get elsewhere.
We also reviewed a list of likely counteroffers with a response for each, and agreed not to ask for the new-customer promotion by name, since that was the most common reason other customers got turned down.
Getting Advice During the Call
I could have made the call with the pitch and a list of counters. But calls rarely follow the script, and I wanted advice on what the rep actually said while I was still on the phone.
So I flipped the setup around. Instead of a bot dialing and handing the call to me, I dialed from my cell through the Twilio number, and the app used Twilio’s live transcription to transcribe both sides as we talked. A Claude Code session in another project followed that live transcript and coached me in my terminal:
Me on my cell
Twilio number carries the call
The rep at the provider
The app logs each line
Claude Code reads the log live
My terminal advice in, questions out
BOTH SIDES TRANSCRIBED
LIVE TRANSCRIPT
THE CALL
How the call was set up. I talked to the rep through the Twilio number, the app logged both sides, and Claude Code read that log and coached me in my terminal.
Calling from my own phone meant the menu recognized me as a customer, so I skipped most of the menu steps my app had struggled with on the first attempt.
I made the call on October 26 th . It took about ten and a half minutes, and the rep offered the price I asked for about 75 seconds after my pitch. The new price started that day, and I accepted it from an emailed link before I hung up.
Rep Hi, thank you so much for calling [provider] .
Claude An agent picked up. Get their name, then give the opening line.
Me I'm a [long-time] customer, [top tier] . I'm paying [current price] a month for gig internet. A competitor has the same gig speed at my address for [half] , and your own site advertises [half] for new customers. I'd rather stay. What can you do to get me to [half] ?
Claude Good delivery. Now stay quiet and let them answer. I'll have counters ready.
Rep I see the 1 gig for [half] , with autopay and paperless active, and a 5-year price guarantee.
Claude [half] for 1 Gig with a 5-year price guarantee. That's your target. Accept it, but confirm these first: is it the total, is the gateway included, and when does it start.
Me Is [half] the total on the bill, or are taxes and fees extra?
Me, typed to Claude what is gateway? i have my own modem and router
Claude The gateway is the provider's rented modem and router in one box. Since you own yours, skip the equipment question.
Rep Your total is [half] flat after fees and taxes.
Me When does that start? Mid billing cycle, or at my next one?
Me Is there any contract or cancellation fee?
Rep No contract, no cancellation fees. The plan is unlimited, with no modem rental fees. If you want to cancel an hour after the call, you can.
Claude No contract, no cancellation fee, unlimited data, no modem rental. Everything checks out.
Rep I'm sending you the approval now. It'll be in your email in a couple of seconds.
Me, typed to Claude just saved a full webpage to downloads
Claude The terms match: [half] a month with taxes included, the 5-year guarantee, no contract. Click I agree.
Rep Your account is also eligible for a free mobile line for the first 12 months.
Claude That's an upsell. Decline politely: "I'm all set with my current carrier."
Me We'll look into it. Thanks for your help.
The next bill estimate, on October 28 th , showed the new price. The first full bill arrives in early December.
Working on something interesting?
I read every message and reply personally.
