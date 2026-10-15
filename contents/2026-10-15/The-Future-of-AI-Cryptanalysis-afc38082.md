---
source: "https://www.schneier.com/blog/archives/2026/10/the-future-of-ai-cryptanalysis.html"
hn_url: "https://news.ycombinator.com/item?id=50110342"
title: "The Future of AI Cryptanalysis"
article_title: "The Future of AI Cryptanalysis - Schneier on Security"
image: ""
author: "wondersky"
captured_at: "2026-10-15T21:58:29Z"
capture_tool: "hn-digest"
hn_id: 50110342
score: 1
comments: 0
posted_at: "2026-10-15T21:14:14Z"
tags:
  - hacker-news
---

# The Future of AI Cryptanalysis

- HN: [50110342](https://news.ycombinator.com/item?id=50110342)
- Source: [www.schneier.com](https://www.schneier.com/blog/archives/2026/10/the-future-of-ai-cryptanalysis.html)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T21:14:14Z

## Translation

Title: The Future of AI Cryptanalysis
Article title: The Future of AI Cryptanalysis - Schneier on Security
Description: Matthew Green has been writing a rather pessimistic series of posts about the potential for AI systems to vastly outperform humans in cryptanalysis of both symmetric and public-key ciphers. I am more optimistic than he is. That is, I have stronger faith in the over-engineering cryptographers tend to
[truncated]

Article text:
The Future of AI Cryptanalysis
Matthew Green has been writing a rather pessimistic series of posts about the potential for AI systems to vastly outperform humans in cryptanalysis of both symmetric and public-key ciphers.
I am more optimistic than he is. That is, I have stronger faith in the over-engineering cryptographers tend to do. I am trying to write this up in more detail.
Tags: AI , cryptanalysis , cryptography
Posted on October 14, 2026 at 10:24 AM •
11 Comments
Ferentarius •
October 14, 2026 10:40 AM
Re: “I’m not going to get a prize for being right, and everyone will make fun of me for being wrong“
To be right is a barren victory, and to be wrong is a public crucifixion. Life is a theater where applause and ridicule are equally hollow; the only authentic moment is the one in which we recoil from both.
Dave Sanford •
October 14, 2026 11:38 AM
If AIs can solve the Navier-Stokes existence and smoothness problem, they can probably do this. I’m concerned that AI-driven offensive and defensive cryptanalysis could spark a new arms race.
Study4Pass •
October 14, 2026 12:46 PM
Your optimism about over-engineering is reassuring, but it raises a question for me: does that margin protect us mainly against incremental improvements, or also against a genuinely novel attack idea? Human cryptanalysis has historically advanced in jumps — differential, linear, side channels — not just faster brute force. If AI’s edge is in combining techniques and searching spaces humans would never explore by hand, the worry isn’t that it out-computes us, but that it finds a structural shortcut we assumed was impractical to discover. I’m curious whether your longer write-up will address how much of that safety margin is mathematical versus merely ‘nobody has found the trick yet.’ As someone studying cryptography fundamentals through Study4Pass, this distinction between a cipher being strong and being merely unbroken so far is the part I find hardest to internalize.
Clive Robinson •
October 14, 2026 1:10 PM
“I am more optimistic than [Matthew Green] is”
Actually I’m not I’m actually more pessimistic than Matthew Green is…
Cryptographic systems are way way more than algorithms.
We saw this with AES, the implementations designed by cryptographers and put up on the NIST Competition Web Site were not exactly what you might hope from the theoretical strength of the algorithms.
The AES implimentation was designed for speed and in the process opened up a series of side channels.
Which unfortunately leaked “secrets to the line”…
As others noted the NSA knew all about this sort of “implementation problem” but never said anything to NIST or anyone else…
It was in fact so bad that not only were “secrets leaked to line” so strong was the leakage it passed through upstream routers into the cloud.
Where users generally can not see but the NSA can easily do so…
There are still copies of that “bad code” in libraries that get used in things like “Smart Meters” and other “Internet of Things”(IoT) devices.
Is it “accidental or deliberate?”
It is hard to know but as this sort of thing goes on all the time, I suspect even “Post Quantum Cryptography”(PQC) practical implementations will have “accidental side channels”… Effectively “built in” in low cost devices like Smart Meters, Medical Electronics, and similar with half century installation being part of the design spec…
So it’s not the National Security or other similar systems I worry about, but the ordinary cheap and mundane devices.
We are already seeing Anthropic worrying about IoT devices as the next threat to worry about,
Anthropic conjures another threat that doesn’t exist yet to stoke fears of AI apocalypse
Self-replicating AI botnets could become feasible on IoT hardware within months, claims the Claude maker, if smaller models get smarter
“ Nothing sells AI products as well as fear. Using new research, Anthropic claims that self-replicating AI worms could be just months away from spreading across vulnerable IoT devices, which would be a problem if it came true.
The frontier lab published a research report Monday suggesting that AI models could eventually self-replicate without relying on data center hardware, with a 9-billion-parameter model successfully setting up inference on eight of the nine hardware types Anthropic tested. “
https://www.theregister.com/ai-and-ml/2026/10/13/anthropic-conjures-another-threat-that-doesnt-exist-yet-to-stoke-fears-of-ai-apocalypse/5302866
Now depending on what you believe this is a threat or it is not.
I err on the side of caution when assessing such things.
But… What I assess is often not the original idea (they are often to prosaic).
So consider how many IoT devices out there, that could be taken over by a suitably small model.
Just how many would be needed?
A lot but is it an issue with IoT being fairly ubiquitous…
Many years ago this idea was proposed as the “Chinese TV lottery” idea (If memory serves correctly this was before our host @Bruce published his first “Applied Cryptography” book in the early 1990’s (blue cover).
Kempton •
October 14, 2026 5:15 PM
Bruce, I hope your “stronger faith in the over-engineering cryptographers tend to do” is founded. The idea of the non-existence of “a weaker/weakest link(s)” fascinates me. And of course, the idea of #OverEngineering in many subsystems leading to a stronger and more resilient system as a whole seems, at least in theory (maybe also in practice, as your faith indicates), possible. Time will tell. And, fortunately, or unfortunately, with the speed of grow of #AIPower, we won’t need to wait too long if and when beep hits the fan above us all, and the shower of beep won’t be too pretty.
Clive Robinson •
October 15, 2026 12:34 AM
“The idea of the non-existence of “a weaker/weakest link(s)” fascinates me.”
All Systems not just software of any useful complexity will have multiple different links.
Simple logic dictates that if you make a chain of different links then there will always be a weakest link and a strongest link.
But from a more practical point systems “process” that is they have through put from input to output from subsystem to subsystem often what goes through a system is altered differently in each subsystem.
This means that the links within the chain of the system even if identical have different loads. Thus the strength or weakness of a link varies with how what is being processed changes.
“the idea of #OverEngineering in many subsystems leading to a stronger and more resilient system as a whole seems, at least in theory … possible.”
The bad news is “OverEngineering” always adds increased complexity, and complexity is where the harder to resolve vulnerabilities are usually found…
It’s why two of the fundamental design rules of secure design are,
1, Strong segregation by functional isolation with the minimum of needed coupling.
2, Reduce complexity in each functional block to the minimum needed to do a required function.
One “gotcher” is the assumption that things go from left to right and top down. The sad truth is that for reliable function either a designed feedback path or accidental feedback path exists at all times, and this harms the idea of segregation.
For instance there is over confidence in “data diodes” for segregation. Unfortunately to handle “errors or omissions” it necessitates some form of feedback signalling.
Such feedback moves in the reverse direction backwards through the data diode from output back through the input and potentially all the way back through an entire system. As such it can be used for “Fault Injection Attacks” and the creation of “covert channels” by timing attacks.
It’s knowing such things and that humans were effectively designed by evolution to err, that causes my pessimism to increase way faster than others (and for good reason).
So whilst “in theory” we can design tools to develope the parts of a secure system. In practice,
1, The tools will always be imperfect / incomplete.
2, New attack methods not covered by the tools will be discovered on a frequent basis.
Ismar •
October 15, 2026 1:49 AM
It all comes down to the old adage that it is much easier to break things than to make them (work as intended)
Robin •
October 15, 2026 4:35 AM
@Clive, you’re absolutely right about not just cryptographers over-engineering things.
A good counter-example to your design principles comes to mind: about 7 or 8 years ago Lidl were selling a multicooker that used a customised 7″ Android tablet for the control panel. It ran Android 6.0 (Marshmallow) with a MediaTek MT6580 quad-core processor. Don’t know how much RAM it had. Out of the box it was loosely locked down but I remember a project that hacked into it for laughs – maybe activating the microphone? It was obviously cheaper and more flexible to repurpose a commodity tablet that design a whole new bit of hardware. OK, this is an old story and maybe the new versions are better, but I can’t believe that this is the only example of using a hugely over-capable, and easily hackable device for a banal IoT product. Get used to talking nicely to your toaster.
Winter •
October 15, 2026 4:42 AM
It all comes down to the old adage that it is much easier to break things than to make them (work as intended)
That’s the second law of thermodynamics in a nutshell:
Things go wrong all by themselves. To make things go right takes work[1]
[1] Work both in the thermodynamics sense and the social sense.
Clive Robinson •
October 15, 2026 10:04 AM
“Out of the box it was loosely locked down but I remember a project that hacked into it for laughs”
It was not the first and it won’t be the last…
There are quite a number of “Rocket Competitions” for Uni Students to participate in…
The UK “Civil Aviation Authority”(CAA) has overall authority on non passive “toy” rockets. With the “Communications Agency / OfCom having control on EM radiated “communications” from “DC to Daylight and beyond”.
Because of the perceived “terrorism / hostile nation state” risk the CAA do not allow “flight course control” or indirectly the use of the likes of Android mobile phones.
Which from previous experience before the CAA lockdown used to be fun as modern phones have just about everything you need including “visual ground mapping to target to check / avoid GNS systems are not being jammed etc. And For non rocket drones it gives various operator feed back… And the building of a USB to Fiber network cable has likewise been done…
As most of them are made in China and shipped world wide putting a ban on the technology is beyond any National Government (even North Korea has illicit mobile phone issues close to it’s national boarders).
And with what is going on at the East of Europe, and now in the Middle East might not appear surprising.
But… It’s given the Defence Industry Market which used to suck in a lot of money into the UK what is in effect an existential crisis… So big money pressure is being brought on Politicians and the likes of MI6 are being left “Wrong Footed” and “flapping in the breeze” (as has recently been seen)
But that is all fairly modern and about to change drastically…
As some know I’ve gad a fairly eclectic work life in electronics, communications and other engineering …
I used to work for a while in the “Broadcast Industry” where a friend had well regarded product manufacturing that he and I had designed product. However he had issues arise in that customers wanted “fancy controls” and displays even on Transmitters in racks up the top of mountains… And more importantly remote control via computer networking.
Designing our own system was going to be ludicrously expensive and fraught with issues (CE / FCC compliance being the most problematic but most time consuming).
So I suggested using the guts of a two screen very popular hand held computer game console, that had already got a significant “hacker following”.
Thus it was not much more than a simple mechanical task taking it out of it’s plastic “consumer case” an

[truncated]

## Original Extract

Matthew Green has been writing a rather pessimistic series of posts about the potential for AI systems to vastly outperform humans in cryptanalysis of both symmetric and public-key ciphers. I am more optimistic than he is. That is, I have stronger faith in the over-engineering cryptographers tend to
[truncated]

The Future of AI Cryptanalysis
Matthew Green has been writing a rather pessimistic series of posts about the potential for AI systems to vastly outperform humans in cryptanalysis of both symmetric and public-key ciphers.
I am more optimistic than he is. That is, I have stronger faith in the over-engineering cryptographers tend to do. I am trying to write this up in more detail.
Tags: AI , cryptanalysis , cryptography
Posted on October 14, 2026 at 10:24 AM •
11 Comments
Ferentarius •
October 14, 2026 10:40 AM
Re: “I’m not going to get a prize for being right, and everyone will make fun of me for being wrong“
To be right is a barren victory, and to be wrong is a public crucifixion. Life is a theater where applause and ridicule are equally hollow; the only authentic moment is the one in which we recoil from both.
Dave Sanford •
October 14, 2026 11:38 AM
If AIs can solve the Navier-Stokes existence and smoothness problem, they can probably do this. I’m concerned that AI-driven offensive and defensive cryptanalysis could spark a new arms race.
Study4Pass •
October 14, 2026 12:46 PM
Your optimism about over-engineering is reassuring, but it raises a question for me: does that margin protect us mainly against incremental improvements, or also against a genuinely novel attack idea? Human cryptanalysis has historically advanced in jumps — differential, linear, side channels — not just faster brute force. If AI’s edge is in combining techniques and searching spaces humans would never explore by hand, the worry isn’t that it out-computes us, but that it finds a structural shortcut we assumed was impractical to discover. I’m curious whether your longer write-up will address how much of that safety margin is mathematical versus merely ‘nobody has found the trick yet.’ As someone studying cryptography fundamentals through Study4Pass, this distinction between a cipher being strong and being merely unbroken so far is the part I find hardest to internalize.
Clive Robinson •
October 14, 2026 1:10 PM
“I am more optimistic than [Matthew Green] is”
Actually I’m not I’m actually more pessimistic than Matthew Green is…
Cryptographic systems are way way more than algorithms.
We saw this with AES, the implementations designed by cryptographers and put up on the NIST Competition Web Site were not exactly what you might hope from the theoretical strength of the algorithms.
The AES implimentation was designed for speed and in the process opened up a series of side channels.
Which unfortunately leaked “secrets to the line”…
As others noted the NSA knew all about this sort of “implementation problem” but never said anything to NIST or anyone else…
It was in fact so bad that not only were “secrets leaked to line” so strong was the leakage it passed through upstream routers into the cloud.
Where users generally can not see but the NSA can easily do so…
There are still copies of that “bad code” in libraries that get used in things like “Smart Meters” and other “Internet of Things”(IoT) devices.
Is it “accidental or deliberate?”
It is hard to know but as this sort of thing goes on all the time, I suspect even “Post Quantum Cryptography”(PQC) practical implementations will have “accidental side channels”… Effectively “built in” in low cost devices like Smart Meters, Medical Electronics, and similar with half century installation being part of the design spec…
So it’s not the National Security or other similar systems I worry about, but the ordinary cheap and mundane devices.
We are already seeing Anthropic worrying about IoT devices as the next threat to worry about,
Anthropic conjures another threat that doesn’t exist yet to stoke fears of AI apocalypse
Self-replicating AI botnets could become feasible on IoT hardware within months, claims the Claude maker, if smaller models get smarter
“ Nothing sells AI products as well as fear. Using new research, Anthropic claims that self-replicating AI worms could be just months away from spreading across vulnerable IoT devices, which would be a problem if it came true.
The frontier lab published a research report Monday suggesting that AI models could eventually self-replicate without relying on data center hardware, with a 9-billion-parameter model successfully setting up inference on eight of the nine hardware types Anthropic tested. “
https://www.theregister.com/ai-and-ml/2026/10/13/anthropic-conjures-another-threat-that-doesnt-exist-yet-to-stoke-fears-of-ai-apocalypse/5302866
Now depending on what you believe this is a threat or it is not.
I err on the side of caution when assessing such things.
But… What I assess is often not the original idea (they are often to prosaic).
So consider how many IoT devices out there, that could be taken over by a suitably small model.
Just how many would be needed?
A lot but is it an issue with IoT being fairly ubiquitous…
Many years ago this idea was proposed as the “Chinese TV lottery” idea (If memory serves correctly this was before our host @Bruce published his first “Applied Cryptography” book in the early 1990’s (blue cover).
Kempton •
October 14, 2026 5:15 PM
Bruce, I hope your “stronger faith in the over-engineering cryptographers tend to do” is founded. The idea of the non-existence of “a weaker/weakest link(s)” fascinates me. And of course, the idea of #OverEngineering in many subsystems leading to a stronger and more resilient system as a whole seems, at least in theory (maybe also in practice, as your faith indicates), possible. Time will tell. And, fortunately, or unfortunately, with the speed of grow of #AIPower, we won’t need to wait too long if and when beep hits the fan above us all, and the shower of beep won’t be too pretty.
Clive Robinson •
October 15, 2026 12:34 AM
“The idea of the non-existence of “a weaker/weakest link(s)” fascinates me.”
All Systems not just software of any useful complexity will have multiple different links.
Simple logic dictates that if you make a chain of different links then there will always be a weakest link and a strongest link.
But from a more practical point systems “process” that is they have through put from input to output from subsystem to subsystem often what goes through a system is altered differently in each subsystem.
This means that the links within the chain of the system even if identical have different loads. Thus the strength or weakness of a link varies with how what is being processed changes.
“the idea of #OverEngineering in many subsystems leading to a stronger and more resilient system as a whole seems, at least in theory … possible.”
The bad news is “OverEngineering” always adds increased complexity, and complexity is where the harder to resolve vulnerabilities are usually found…
It’s why two of the fundamental design rules of secure design are,
1, Strong segregation by functional isolation with the minimum of needed coupling.
2, Reduce complexity in each functional block to the minimum needed to do a required function.
One “gotcher” is the assumption that things go from left to right and top down. The sad truth is that for reliable function either a designed feedback path or accidental feedback path exists at all times, and this harms the idea of segregation.
For instance there is over confidence in “data diodes” for segregation. Unfortunately to handle “errors or omissions” it necessitates some form of feedback signalling.
Such feedback moves in the reverse direction backwards through the data diode from output back through the input and potentially all the way back through an entire system. As such it can be used for “Fault Injection Attacks” and the creation of “covert channels” by timing attacks.
It’s knowing such things and that humans were effectively designed by evolution to err, that causes my pessimism to increase way faster than others (and for good reason).
So whilst “in theory” we can design tools to develope the parts of a secure system. In practice,
1, The tools will always be imperfect / incomplete.
2, New attack methods not covered by the tools will be discovered on a frequent basis.
Ismar •
October 15, 2026 1:49 AM
It all comes down to the old adage that it is much easier to break things than to make them (work as intended)
Robin •
October 15, 2026 4:35 AM
@Clive, you’re absolutely right about not just cryptographers over-engineering things.
A good counter-example to your design principles comes to mind: about 7 or 8 years ago Lidl were selling a multicooker that used a customised 7″ Android tablet for the control panel. It ran Android 6.0 (Marshmallow) with a MediaTek MT6580 quad-core processor. Don’t know how much RAM it had. Out of the box it was loosely locked down but I remember a project that hacked into it for laughs – maybe activating the microphone? It was obviously cheaper and more flexible to repurpose a commodity tablet that design a whole new bit of hardware. OK, this is an old story and maybe the new versions are better, but I can’t believe that this is the only example of using a hugely over-capable, and easily hackable device for a banal IoT product. Get used to talking nicely to your toaster.
Winter •
October 15, 2026 4:42 AM
It all comes down to the old adage that it is much easier to break things than to make them (work as intended)
That’s the second law of thermodynamics in a nutshell:
Things go wrong all by themselves. To make things go right takes work[1]
[1] Work both in the thermodynamics sense and the social sense.
Clive Robinson •
October 15, 2026 10:04 AM
“Out of the box it was loosely locked down but I remember a project that hacked into it for laughs”
It was not the first and it won’t be the last…
There are quite a number of “Rocket Competitions” for Uni Students to participate in…
The UK “Civil Aviation Authority”(CAA) has overall authority on non passive “toy” rockets. With the “Communications Agency / OfCom having control on EM radiated “communications” from “DC to Daylight and beyond”.
Because of the perceived “terrorism / hostile nation state” risk the CAA do not allow “flight course control” or indirectly the use of the likes of Android mobile phones.
Which from previous experience before the CAA lockdown used to be fun as modern phones have just about everything you need including “visual ground mapping to target to check / avoid GNS systems are not being jammed etc. And For non rocket drones it gives various operator feed back… And the building of a USB to Fiber network cable has likewise been done…
As most of them are made in China and shipped world wide putting a ban on the technology is beyond any National Government (even North Korea has illicit mobile phone issues close to it’s national boarders).
And with what is going on at the East of Europe, and now in the Middle East might not appear surprising.
But… It’s given the Defence Industry Market which used to suck in a lot of money into the UK what is in effect an existential crisis… So big money pressure is being brought on Politicians and the likes of MI6 are being left “Wrong Footed” and “flapping in the breeze” (as has recently been seen)
But that is all fairly modern and about to change drastically…
As some know I’ve gad a fairly eclectic work life in electronics, communications and other engineering …
I used to work for a while in the “Broadcast Industry” where a friend had well regarded product manufacturing that he and I had designed product. However he had issues arise in that customers wanted “fancy controls” and displays even on Transmitters in racks up the top of mountains… And more importantly remote control via computer networking.
Designing our own system was going to be ludicrously expensive and fraught with issues (CE / FCC compliance being the most problematic but most time consuming).
So I suggested using the guts of a two screen very popular hand held computer game console, that had already got a significant “hacker following”.
Thus it was not much more than a simple mechanical task taking it out of it’s plastic “consumer case” an

[truncated]
