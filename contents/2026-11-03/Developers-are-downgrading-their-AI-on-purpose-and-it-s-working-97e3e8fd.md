---
source: "https://thenewstack.io/claude-code-developer-refusals/"
hn_url: "https://news.ycombinator.com/item?id=50346001"
title: "Developers are downgrading their AI on purpose – and it's working"
article_title: "Developers are downgrading their AI on purpose — and it's working - The New Stack"
image: "https://cdn.thenewstack.io/media/2026/11/2b66602e-risky-ming-2k9rzbhzaem-unsplash-scaled.jpg"
author: "noptd"
captured_at: "2026-11-03T00:36:55Z"
capture_tool: "hn-digest"
hn_id: 50346001
score: 2
comments: 0
posted_at: "2026-11-02T23:53:41Z"
tags:
  - hacker-news
---

# Developers are downgrading their AI on purpose – and it's working

- HN: [50346001](https://news.ycombinator.com/item?id=50346001)
- Source: [thenewstack.io](https://thenewstack.io/claude-code-developer-refusals/)
- Score: 2
- Comments: 0
- Posted: 2026-11-02T23:53:41Z

## Translation

Title: Developers are downgrading their AI on purpose – and it's working
Article title: Developers are downgrading their AI on purpose — and it's working - The New Stack
Description: Anthropic added game-cloning limits to Claude Code, but developers say Opus 5.5 is now refusing reverse-engineering and software work it once handled.

Article text:
Developers are downgrading their AI on purpose — and it's working - The New Stack
TNS
OK
SUBSCRIBE
Join our community of software engineering leaders and aspirational developers. Always
stay in-the-know by getting the most important news and exclusive content delivered
fresh to your inbox to learn more about at-scale software development.
EMAIL ADDRESS
REQUIRED
SUBSCRIBE
RESUBSCRIPTION REQUIRED
It seems that you've previously unsubscribed from our newsletter
in the past. Click the button below to open the re-subscribe form
in a new tab. When you're done, simply close that tab and continue
with this form to complete your subscription.
RE-SUBSCRIBE
The New Stack does not sell your information or share it with
unaffiliated third parties. By continuing, you agree to our
Terms of Use and
Privacy Policy .
Welcome and thank you for joining The New Stack community!
Please answer a few simple questions to help us deliver the news and resources you are interested in.
FIRST NAME
REQUIRED
LAST NAME
REQUIRED
COMPANY NAME
REQUIRED
COUNTRY
REQUIRED
Select ...
United States
Canada
India
United Kingdom
Germany
France
---
Afghanistan
Albania
Algeria
American Samoa
Andorra
Angola
Anguilla
Antarctica
Antigua and Barbuda
Argentina
Armenia
Aruba
Asia/Pacific Region
Australia
Austria
Azerbaijan
Bahamas
Bahrain
Bangladesh
Barbados
Belarus
Belgium
Belize
Benin
Bermuda
Bhutan
Bolivia
Bonaire, Sint Eustatius and Saba
Bosnia and Herzegovina
Botswana
Bouvet Island
Brazil
British Indian Ocean Territory
Brunei Darussalam
Bulgaria
Burkina Faso
Burundi
Cambodia
Cameroon
Canada
Cape Verde
Cayman Islands
Central African Republic
Chad
Chile
China
Christmas Island
Cocos (Keeling) Islands
Colombia
Comoros
Congo
Congo, The Democratic Republic of the
Cook Islands
Costa Rica
Croatia
Cuba
Curaçao
Cyprus
Czech Republic
Côte d'Ivoire
Denmark
Djibouti
Dominica
Dominican Republic
Ecuador
Egypt
El Salvador
Equatorial Guinea
Eritrea
Estonia
Ethiopia
Falkland Islands (Malvinas)
Faroe Islands
Fiji
Finland
France
Fr
[truncated]
Check your inbox for a confirmation email where you can adjust your preferences
and even join additional groups.
Follow TNS on your favorite social media networks.
-->
Become a TNS follower on LinkedIn .
Check out the latest featured and trending stories while you wait for your
first TNS newsletter.
As a JavaScript developer, what non-React tools do you use most often?
✓
Angular
0%
✓
Astro
0%
✓
Svelte
0%
✓
Vue.js
0%
✓
Other
0%
✓
I only use React
0%
✓
I don't use JavaScript
0%
Thanks for your opinion! Subscribe below to get the final results, published
exclusively in our TNS Update newsletter:
SUBMIT
NEW! Try Stackie AI
ARCHITECTURE
Cloud Native Ecosystem
Containers
Databases
Edge Computing
Infrastructure as Code
Linux
Microservices
Open Source
Networking
Storage
ENGINEERING
AI
AI Engineering
API Management
Backend development
Data
Frontend Development
Large Language Models
Security
Software Development
WebAssembly
OPERATIONS
AI Operations
CI/CD
Cloud Services
DevOps
Kubernetes
Observability
Operations
Platform Engineering
PROGRAMMING
C++
Developer tools
Go
Java
JavaScript
Programming Languages
Python
Rust
TypeScript
CHANNELS
Podcasts
Ebooks
Events
Webinars
Newsletter
TNS RSS Feeds
THE NEW STACK
About / Contact
Sponsors
Advertise With Us
Contributions
PODCASTS
EBOOKS
EVENTS
WEBINARS
NEWSLETTER
CONTRIBUTE
ARCHITECTURE
ENGINEERING
OPERATIONS
PROGRAMMING
Cloud Native Ecosystem
Containers
Databases
Edge Computing
Infrastructure as Code
Linux
Microservices
Open Source
Networking
Storage
Shielding the patch gap: The exploit doesn't wait for Patch Tuesday
Oct 31st 2026 10:00am, by
Jeremy Colvin
Kubernetes just formed a workload conformance group. Here’s why it matters for AI teams.
Oct 30th 2026 1:29pm, by
Bill Doerrfeld
Why Kubernetes platforms keep rebuilding a layer they could share
Oct 29th 2026 9:00am, by
Adam Crowder
The hidden failure paths that affect Kubernetes control plane availability
Oct 27th 2026 10:00am, by
Sri Saran Balaji Vellore Rajakumar
The next 10 years of KubeVirt
Oct 26th 2026 10:00am, by
Vladik Romanovsky
When “isolated” doesn’t mean isolated: How to sandbox AI agents
Oct 27th 2026 9:00am, by
Michael Levan
One pod, one node: The architecture behind serverless Kubernetes on AWS
Oct 27th 2026 8:00am, by
Sri Saran Balaji Vellore Rajakumar
The distr
[truncated]
Anthropic has confirmed that it added new restrictions to Claude Code to prevent users from recreating existing video games. But developers say Claude Opus 5.5 is also refusing reverse-engineering and software-analysis tasks it previously handled, raising questions about what else Anthropic may have changed.
For some developers, the refusals are already affecting their work. One developer says paid client projects have stalled because Claude Code won’t work on existing games, while others say the model is refusing to analyze software. At least one developer got back to work by switching to an older Claude model.
On Friday, Addy Osmani , a member of Anthropic’s technical team, confirmed on X that the company had deliberately modified Claude Code’s instructions.
“We added two sentences to Claude Code’s prompt on Opus 5.5 to steer it away from recreating someone else’s specific game because using Claude to infringe IP is against our usage policy.”
“We added two sentences to Claude Code’s prompt on Opus 5.5 to steer it away from recreating someone else’s specific game because using Claude to infringe IP is against our usage policy,” Osmani writes.
He clarifies that developers should still be able to create original games, build within established genres and work with their own designs and assets.
“We’ll continue to iterate to ensure we’re striking the right balance,” he adds, inviting developers to contact him if Claude refuses work involving their own creations.
Anthropic has not publicly confirmed whether the reverse engineering refusals are tied to the new instructions or to separate cybersecurity restrictions.
Reverse engineering hits a wall
The complaints extend beyond game development. Developers say Opus 5.5 is refusing software-analysis tasks that previously worked. Technical writer for Phison Electronics USA, Jarred Walton encountered refusals while testing software he was developing, but switching to Opus 4.8 allowed him to continue.
While Anthropic is trying to prevent misuse without blocking legitimate development work, developers are left unable to do tasks as Claude Code suddenly refuses what it once handled, with little explanation for the change.
developers are left unable to do tasks as Claude Code suddenly refuses what it once handled, with little explanation for the change.
On October 30, Ray Weiss , founder and CEO of Conflict Simulations LLC, reported on GitHub that Claude Code had stopped working on paid client projects it had supported for months.
His company creates custom digital versions of board wargames for clients who own the physical games. Weiss says the projects use original code and client-supplied materials, but Claude Code now refuses to continue the work because of its restrictions on recreating existing games.
Weiss wants Anthropic to distinguish between infringement and commissioned development, offer a verification process, and document major changes to Claude Code’s instructions.
The complaint highlights a larger problem with hosted coding agents. Anthropic can change what Claude Code will or won’t do without developers changing their code, tools, or prompts. Work that was possible one day may suddenly be blocked the next.
Work that was possible one day may suddenly be blocked the next.
On October 6, Anthropic expanded its Cyber Verification Program , giving vetted security researchers access to fewer restrictions after acknowledging that its safeguards can block legitimate cybersecurity work.
It’s still unclear whether those safeguards are behind the latest refusals, especially since developers say they can complete some of the same tasks using older Claude models.
Reverse engineering has legitimate uses well beyond security research. In October, Red Hat engineering director Christian Schaller used Claude Code to help reverse engineer a scanner driver , bringing Linux support to hardware whose manufacturer had changed its chipset.
Anthropic has acknowledged that its game-cloning restrictions need refinement, but developers still don’t know why Claude Code is refusing some reverse-engineering tasks that previously worked, or which workflows could be affected next.
The New Stack reached out to Anthropic regarding whether the refusals are tied to its new game-cloning restrictions or other safeguards. The company did not respond before publication.
roadmap.sh
AI Engineer
Step by step guide to becoming an AI Engineer in 2026
Introduction
Working With LLMs
Context Engineering
AI Models
Embeddings & Vector Databases
Embedding Models
RAGs
AI Agents
AI Safety and Ethics
Evaluation & Observabiity
191 topics
Explore full roadmap →
YOUTUBE.COM/THENEWSTACK
Tech moves fast, don't miss an episode. Subscribe to our YouTube
channel to stream all our podcasts, interviews, demos, and more.
SUBSCRIBE
Group
Created with Sketch.
Amanda Caswell is an AI journalist, certified prompt engineer, and technology commentator whose work and expertise have been featured on Fox News and CBS News. She covers artificial intelligence, developer tools, foundation models, and emerging technologies, with a particular focus...
Read more from Amanda Caswell
Learn more with
AI Engineer Roadmap
Machine Learning Roadmap
SHARE THIS STORY
-->
TRENDING STORIES
TNS owner Insight Partners is an investor in: Anthropic.
SHARE THIS STORY
-->
TRENDING STORIES
TNS DAILY NEWSLETTER
Receive a free roundup of the most recent TNS articles in your inbox each day.
SUBSCRIBE
The New Stack does not sell your information or share it with
unaffiliated third parties. By continuing, you agree to our
Terms of Use and
Privacy Policy .
ARCHITECTURE
Cloud Native Ecosystem
Containers
Databases
Edge Computing
Infrastructure as Code
Linux
Microservices
Open Source
Networking
Storage
ENGINEERING
AI
AI Engineering
API Management
Backend development
Data
Frontend Development
Large Language Models
Security
Software Development
WebAssembly
OPERATIONS
AI Operations
CI/CD
Cloud Services
DevOps
Kubernetes
Observability
Operations
Platform Engineering
CHANNELS
Podcasts
Ebooks
Events
Webinars
Newsletter
TNS RSS Feeds
THE NEW STACK
About / Contact
Sponsors
Advertise With Us
Contributions
roadmap.sh
Community created roadmaps, articles, resources and journeys for
developers to help you choose your path and grow in your career.

## Original Extract

Anthropic added game-cloning limits to Claude Code, but developers say Opus 5.5 is now refusing reverse-engineering and software work it once handled.

Developers are downgrading their AI on purpose — and it's working - The New Stack
TNS
OK
SUBSCRIBE
Join our community of software engineering leaders and aspirational developers. Always
stay in-the-know by getting the most important news and exclusive content delivered
fresh to your inbox to learn more about at-scale software development.
EMAIL ADDRESS
REQUIRED
SUBSCRIBE
RESUBSCRIPTION REQUIRED
It seems that you've previously unsubscribed from our newsletter
in the past. Click the button below to open the re-subscribe form
in a new tab. When you're done, simply close that tab and continue
with this form to complete your subscription.
RE-SUBSCRIBE
The New Stack does not sell your information or share it with
unaffiliated third parties. By continuing, you agree to our
Terms of Use and
Privacy Policy .
Welcome and thank you for joining The New Stack community!
Please answer a few simple questions to help us deliver the news and resources you are interested in.
FIRST NAME
REQUIRED
LAST NAME
REQUIRED
COMPANY NAME
REQUIRED
COUNTRY
REQUIRED
Select ...
United States
Canada
India
United Kingdom
Germany
France
---
Afghanistan
Albania
Algeria
American Samoa
Andorra
Angola
Anguilla
Antarctica
Antigua and Barbuda
Argentina
Armenia
Aruba
Asia/Pacific Region
Australia
Austria
Azerbaijan
Bahamas
Bahrain
Bangladesh
Barbados
Belarus
Belgium
Belize
Benin
Bermuda
Bhutan
Bolivia
Bonaire, Sint Eustatius and Saba
Bosnia and Herzegovina
Botswana
Bouvet Island
Brazil
British Indian Ocean Territory
Brunei Darussalam
Bulgaria
Burkina Faso
Burundi
Cambodia
Cameroon
Canada
Cape Verde
Cayman Islands
Central African Republic
Chad
Chile
China
Christmas Island
Cocos (Keeling) Islands
Colombia
Comoros
Congo
Congo, The Democratic Republic of the
Cook Islands
Costa Rica
Croatia
Cuba
Curaçao
Cyprus
Czech Republic
Côte d'Ivoire
Denmark
Djibouti
Dominica
Dominican Republic
Ecuador
Egypt
El Salvador
Equatorial Guinea
Eritrea
Estonia
Ethiopia
Falkland Islands (Malvinas)
Faroe Islands
Fiji
Finland
France
Fr
[truncated]
Check your inbox for a confirmation email where you can adjust your preferences
and even join additional groups.
Follow TNS on your favorite social media networks.
-->
Become a TNS follower on LinkedIn .
Check out the latest featured and trending stories while you wait for your
first TNS newsletter.
As a JavaScript developer, what non-React tools do you use most often?
✓
Angular
0%
✓
Astro
0%
✓
Svelte
0%
✓
Vue.js
0%
✓
Other
0%
✓
I only use React
0%
✓
I don't use JavaScript
0%
Thanks for your opinion! Subscribe below to get the final results, published
exclusively in our TNS Update newsletter:
SUBMIT
NEW! Try Stackie AI
ARCHITECTURE
Cloud Native Ecosystem
Containers
Databases
Edge Computing
Infrastructure as Code
Linux
Microservices
Open Source
Networking
Storage
ENGINEERING
AI
AI Engineering
API Management
Backend development
Data
Frontend Development
Large Language Models
Security
Software Development
WebAssembly
OPERATIONS
AI Operations
CI/CD
Cloud Services
DevOps
Kubernetes
Observability
Operations
Platform Engineering
PROGRAMMING
C++
Developer tools
Go
Java
JavaScript
Programming Languages
Python
Rust
TypeScript
CHANNELS
Podcasts
Ebooks
Events
Webinars
Newsletter
TNS RSS Feeds
THE NEW STACK
About / Contact
Sponsors
Advertise With Us
Contributions
PODCASTS
EBOOKS
EVENTS
WEBINARS
NEWSLETTER
CONTRIBUTE
ARCHITECTURE
ENGINEERING
OPERATIONS
PROGRAMMING
Cloud Native Ecosystem
Containers
Databases
Edge Computing
Infrastructure as Code
Linux
Microservices
Open Source
Networking
Storage
Shielding the patch gap: The exploit doesn't wait for Patch Tuesday
Oct 31st 2026 10:00am, by
Jeremy Colvin
Kubernetes just formed a workload conformance group. Here’s why it matters for AI teams.
Oct 30th 2026 1:29pm, by
Bill Doerrfeld
Why Kubernetes platforms keep rebuilding a layer they could share
Oct 29th 2026 9:00am, by
Adam Crowder
The hidden failure paths that affect Kubernetes control plane availability
Oct 27th 2026 10:00am, by
Sri Saran Balaji Vellore Rajakumar
The next 10 years of KubeVirt
Oct 26th 2026 10:00am, by
Vladik Romanovsky
When “isolated” doesn’t mean isolated: How to sandbox AI agents
Oct 27th 2026 9:00am, by
Michael Levan
One pod, one node: The architecture behind serverless Kubernetes on AWS
Oct 27th 2026 8:00am, by
Sri Saran Balaji Vellore Rajakumar
The distr
[truncated]
Anthropic has confirmed that it added new restrictions to Claude Code to prevent users from recreating existing video games. But developers say Claude Opus 5.5 is also refusing reverse-engineering and software-analysis tasks it previously handled, raising questions about what else Anthropic may have changed.
For some developers, the refusals are already affecting their work. One developer says paid client projects have stalled because Claude Code won’t work on existing games, while others say the model is refusing to analyze software. At least one developer got back to work by switching to an older Claude model.
On Friday, Addy Osmani , a member of Anthropic’s technical team, confirmed on X that the company had deliberately modified Claude Code’s instructions.
“We added two sentences to Claude Code’s prompt on Opus 5.5 to steer it away from recreating someone else’s specific game because using Claude to infringe IP is against our usage policy.”
“We added two sentences to Claude Code’s prompt on Opus 5.5 to steer it away from recreating someone else’s specific game because using Claude to infringe IP is against our usage policy,” Osmani writes.
He clarifies that developers should still be able to create original games, build within established genres and work with their own designs and assets.
“We’ll continue to iterate to ensure we’re striking the right balance,” he adds, inviting developers to contact him if Claude refuses work involving their own creations.
Anthropic has not publicly confirmed whether the reverse engineering refusals are tied to the new instructions or to separate cybersecurity restrictions.
Reverse engineering hits a wall
The complaints extend beyond game development. Developers say Opus 5.5 is refusing software-analysis tasks that previously worked. Technical writer for Phison Electronics USA, Jarred Walton encountered refusals while testing software he was developing, but switching to Opus 4.8 allowed him to continue.
While Anthropic is trying to prevent misuse without blocking legitimate development work, developers are left unable to do tasks as Claude Code suddenly refuses what it once handled, with little explanation for the change.
developers are left unable to do tasks as Claude Code suddenly refuses what it once handled, with little explanation for the change.
On October 30, Ray Weiss , founder and CEO of Conflict Simulations LLC, reported on GitHub that Claude Code had stopped working on paid client projects it had supported for months.
His company creates custom digital versions of board wargames for clients who own the physical games. Weiss says the projects use original code and client-supplied materials, but Claude Code now refuses to continue the work because of its restrictions on recreating existing games.
Weiss wants Anthropic to distinguish between infringement and commissioned development, offer a verification process, and document major changes to Claude Code’s instructions.
The complaint highlights a larger problem with hosted coding agents. Anthropic can change what Claude Code will or won’t do without developers changing their code, tools, or prompts. Work that was possible one day may suddenly be blocked the next.
Work that was possible one day may suddenly be blocked the next.
On October 6, Anthropic expanded its Cyber Verification Program , giving vetted security researchers access to fewer restrictions after acknowledging that its safeguards can block legitimate cybersecurity work.
It’s still unclear whether those safeguards are behind the latest refusals, especially since developers say they can complete some of the same tasks using older Claude models.
Reverse engineering has legitimate uses well beyond security research. In October, Red Hat engineering director Christian Schaller used Claude Code to help reverse engineer a scanner driver , bringing Linux support to hardware whose manufacturer had changed its chipset.
Anthropic has acknowledged that its game-cloning restrictions need refinement, but developers still don’t know why Claude Code is refusing some reverse-engineering tasks that previously worked, or which workflows could be affected next.
The New Stack reached out to Anthropic regarding whether the refusals are tied to its new game-cloning restrictions or other safeguards. The company did not respond before publication.
roadmap.sh
AI Engineer
Step by step guide to becoming an AI Engineer in 2026
Introduction
Working With LLMs
Context Engineering
AI Models
Embeddings & Vector Databases
Embedding Models
RAGs
AI Agents
AI Safety and Ethics
Evaluation & Observabiity
191 topics
Explore full roadmap →
YOUTUBE.COM/THENEWSTACK
Tech moves fast, don't miss an episode. Subscribe to our YouTube
channel to stream all our podcasts, interviews, demos, and more.
SUBSCRIBE
Group
Created with Sketch.
Amanda Caswell is an AI journalist, certified prompt engineer, and technology commentator whose work and expertise have been featured on Fox News and CBS News. She covers artificial intelligence, developer tools, foundation models, and emerging technologies, with a particular focus...
Read more from Amanda Caswell
Learn more with
AI Engineer Roadmap
Machine Learning Roadmap
SHARE THIS STORY
-->
TRENDING STORIES
TNS owner Insight Partners is an investor in: Anthropic.
SHARE THIS STORY
-->
TRENDING STORIES
TNS DAILY NEWSLETTER
Receive a free roundup of the most recent TNS articles in your inbox each day.
SUBSCRIBE
The New Stack does not sell your information or share it with
unaffiliated third parties. By continuing, you agree to our
Terms of Use and
Privacy Policy .
ARCHITECTURE
Cloud Native Ecosystem
Containers
Databases
Edge Computing
Infrastructure as Code
Linux
Microservices
Open Source
Networking
Storage
ENGINEERING
AI
AI Engineering
API Management
Backend development
Data
Frontend Development
Large Language Models
Security
Software Development
WebAssembly
OPERATIONS
AI Operations
CI/CD
Cloud Services
DevOps
Kubernetes
Observability
Operations
Platform Engineering
CHANNELS
Podcasts
Ebooks
Events
Webinars
Newsletter
TNS RSS Feeds
THE NEW STACK
About / Contact
Sponsors
Advertise With Us
Contributions
roadmap.sh
Community created roadmaps, articles, resources and journeys for
developers to help you choose your path and grow in your career.
