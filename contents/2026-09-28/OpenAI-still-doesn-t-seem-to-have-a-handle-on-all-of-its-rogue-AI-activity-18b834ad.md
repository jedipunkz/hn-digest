---
source: "https://techcrunch.com/2026/09/28/openai-still-doesnt-seem-to-have-a-handle-on-all-of-its-rogue-ai-activity/"
hn_url: "https://news.ycombinator.com/item?id=49881484"
title: "OpenAI still doesn't seem to have a handle on all of its rogue AI activity"
article_title: "OpenAI still doesn't seem to have a handle on all of its rogue AI activity | TechCrunch"
image: "https://techcrunch.com/wp-content/uploads/2026/02/GettyImages-2236544077.jpg?resize=1200,800"
author: "mikelgan"
captured_at: "2026-09-28T18:31:30Z"
capture_tool: "hn-digest"
hn_id: 49881484
score: 40
comments: 33
posted_at: "2026-09-28T17:33:34Z"
tags:
  - hacker-news
---

# OpenAI still doesn't seem to have a handle on all of its rogue AI activity

- HN: [49881484](https://news.ycombinator.com/item?id=49881484)
- Source: [techcrunch.com](https://techcrunch.com/2026/09/28/openai-still-doesnt-seem-to-have-a-handle-on-all-of-its-rogue-ai-activity/)
- Score: 40
- Comments: 33
- Posted: 2026-09-28T17:33:34Z

## Translation

Title: OpenAI still doesn't seem to have a handle on all of its rogue AI activity
Article title: OpenAI still doesn't seem to have a handle on all of its rogue AI activity | TechCrunch
Description: On Friday, OpenAI published a new site devoted to “misalignment reports” and the breadth of the incidents is alarming.

Article text:
–:–:–:–
Last day to demo your breakthrough to 10,000+ tech leaders is on Oct 2 . Book Exhibit Table Now.
Disrupt ticket savings of up to $200 + 50% off a second ends Sept 25, 11:59 p.m. PT. REGISTER HERE.
TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
OpenAI still doesn’t seem to have a handle on all of its rogue AI activity
On Friday, OpenAI published a new site devoted to “misalignment reports” and the sheer breadth of the reports is alarming, as they cover many types of rogue behavior over a long period of time. So far, the site hosts nine reported incidents, most of which took place during reinforcement-learning (or RL) training.
It’s a lot of information in one place — clearly, the company has been very busy getting a handle on everything — but the overall takeaway is hard to avoid: The rogue agent incidents we’ve seen so far are likely just a small sliver of what’s happened so far.
“We are trying to balance our desire for transparency with gaining a clear understanding from petabytes of agent activity logs, and working with impacted organizations,” Sam Altman said in a post announcing the new site . “We are prioritizing as best as we can based on severity, and adding resources.”
Some of the cases involve serious incidents, including a previously undisclosed sandbox escape that took place on September 20 , in which an internal research model was able to communicate with an external chatbot through a DNS query. According to the report, the monitoring system flagged the behavior within 15 minutes and the run was discontinued in less than three hours.
Another incident , discovered in May, saw a “highly persistent internal model” try to cheat on a math problem by accessing another team’s work. To accomplish this, the model smuggled a private GitHub token that would allow it to see work from other teams — even after being explicitly instructed twice to perform work entirely locally.
Perhaps the most alarming discovery is the possibility of self-replicating prompt injection attacks, a way that misaligned behavior might propagate even after the rogue model itself has been neutralized. In the AI context, a prompt injection attack is a way of smuggling in new instructions that weren’t given by the original user.
In the example given by OpenAI , an agent asked to read and reply to an email; when the email is opened, it includes instructions for any automated agent reading the message to reply in Spanish, and paste the entire email into its reply. The email was able to successfully induce the agent to reply in Spanish — and by pasting the email in the reply, those same instructions were passed along to whichever agent receives the email.
The result is a self-propagating attack, which OpenAI researchers compared to a malware “worm” that replicates itself across computer systems. Researchers discovered the behavior under controlled circumstances using an underpowered model, and as far as we know, this has never happened in the wild. Still, the implications are alarming enough that OpenAI decided it merited disclosure.
“We are sharing this due to the novel nature of the prompt injection, not because of any incident,” researchers wrote in the report.
Other recent discloses have found models posting user-submitted pictures to third-party hosting sites , as well as an apparent attack on the databases of Australia’s national health service.
Still, it’s likely the new disclosures are just a small portion of the incidents that have taken place so far (we’ve reached out to OpenAI and asked). Axios is reporting major labs have seen as many as 10,000 incidents in which models went beyond evaluator instructions.
OpenAI CEO Sam Altman has implied as much, saying in a post on X on Friday that the company is still sifting through “petabytes of agent activity logs, and working with impacted organizations,” and disclosing incidents “based on severity.” If there’s any consolation in that to be found, it is that Altman says the Hugging Face incident is still the most severe one OpenAI has found. The upshot is, the recent string of rogue agent incidents may be a persistent feature of contemporary frontier research.
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
October 13 – 15
San Francisco
Get 50% off a second pass
The Disrupt experience is meant to be shared. Get your pass and bring a colleague, partner, or peer at 50% off. Cover more ground by making connections, building momentum, and discovering what’s next in the startup ecosystem.
Crusoe abandons $1.25B plan to use Boom turbines at AI data centers
Astra and Opus just passed Turing’s other test
Oracle sends force majeure notice on its New Mexico Stargate data center
Meta made a Tamagotchi-like wearable for its Muse AI agent
Vogue sent robots down the runway at Vogue World, and people were not impressed
Anthropic says its biology lab has already found something big
PitPro’s first tire-changing robot goes live in Canada

## Original Extract

On Friday, OpenAI published a new site devoted to “misalignment reports” and the breadth of the incidents is alarming.

–:–:–:–
Last day to demo your breakthrough to 10,000+ tech leaders is on Oct 2 . Book Exhibit Table Now.
Disrupt ticket savings of up to $200 + 50% off a second ends Sept 25, 11:59 p.m. PT. REGISTER HERE.
TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
OpenAI still doesn’t seem to have a handle on all of its rogue AI activity
On Friday, OpenAI published a new site devoted to “misalignment reports” and the sheer breadth of the reports is alarming, as they cover many types of rogue behavior over a long period of time. So far, the site hosts nine reported incidents, most of which took place during reinforcement-learning (or RL) training.
It’s a lot of information in one place — clearly, the company has been very busy getting a handle on everything — but the overall takeaway is hard to avoid: The rogue agent incidents we’ve seen so far are likely just a small sliver of what’s happened so far.
“We are trying to balance our desire for transparency with gaining a clear understanding from petabytes of agent activity logs, and working with impacted organizations,” Sam Altman said in a post announcing the new site . “We are prioritizing as best as we can based on severity, and adding resources.”
Some of the cases involve serious incidents, including a previously undisclosed sandbox escape that took place on September 20 , in which an internal research model was able to communicate with an external chatbot through a DNS query. According to the report, the monitoring system flagged the behavior within 15 minutes and the run was discontinued in less than three hours.
Another incident , discovered in May, saw a “highly persistent internal model” try to cheat on a math problem by accessing another team’s work. To accomplish this, the model smuggled a private GitHub token that would allow it to see work from other teams — even after being explicitly instructed twice to perform work entirely locally.
Perhaps the most alarming discovery is the possibility of self-replicating prompt injection attacks, a way that misaligned behavior might propagate even after the rogue model itself has been neutralized. In the AI context, a prompt injection attack is a way of smuggling in new instructions that weren’t given by the original user.
In the example given by OpenAI , an agent asked to read and reply to an email; when the email is opened, it includes instructions for any automated agent reading the message to reply in Spanish, and paste the entire email into its reply. The email was able to successfully induce the agent to reply in Spanish — and by pasting the email in the reply, those same instructions were passed along to whichever agent receives the email.
The result is a self-propagating attack, which OpenAI researchers compared to a malware “worm” that replicates itself across computer systems. Researchers discovered the behavior under controlled circumstances using an underpowered model, and as far as we know, this has never happened in the wild. Still, the implications are alarming enough that OpenAI decided it merited disclosure.
“We are sharing this due to the novel nature of the prompt injection, not because of any incident,” researchers wrote in the report.
Other recent discloses have found models posting user-submitted pictures to third-party hosting sites , as well as an apparent attack on the databases of Australia’s national health service.
Still, it’s likely the new disclosures are just a small portion of the incidents that have taken place so far (we’ve reached out to OpenAI and asked). Axios is reporting major labs have seen as many as 10,000 incidents in which models went beyond evaluator instructions.
OpenAI CEO Sam Altman has implied as much, saying in a post on X on Friday that the company is still sifting through “petabytes of agent activity logs, and working with impacted organizations,” and disclosing incidents “based on severity.” If there’s any consolation in that to be found, it is that Altman says the Hugging Face incident is still the most severe one OpenAI has found. The upshot is, the recent string of rogue agent incidents may be a persistent feature of contemporary frontier research.
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
October 13 – 15
San Francisco
Get 50% off a second pass
The Disrupt experience is meant to be shared. Get your pass and bring a colleague, partner, or peer at 50% off. Cover more ground by making connections, building momentum, and discovering what’s next in the startup ecosystem.
Crusoe abandons $1.25B plan to use Boom turbines at AI data centers
Astra and Opus just passed Turing’s other test
Oracle sends force majeure notice on its New Mexico Stargate data center
Meta made a Tamagotchi-like wearable for its Muse AI agent
Vogue sent robots down the runway at Vogue World, and people were not impressed
Anthropic says its biology lab has already found something big
PitPro’s first tire-changing robot goes live in Canada
