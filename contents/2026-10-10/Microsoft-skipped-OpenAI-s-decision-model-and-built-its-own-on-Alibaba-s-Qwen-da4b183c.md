---
source: "https://thenewstack.io/microsoft-decision-model-foundry/"
hn_url: "https://news.ycombinator.com/item?id=50034279"
title: "Microsoft skipped OpenAI's decision model and built its own on Alibaba's Qwen"
article_title: "Microsoft skipped OpenAI’s decision model and built its own on Alibaba’s Qwen - The New Stack"
image: "https://cdn.thenewstack.io/media/2026/10/57c94021-logan-voss-5kzn1icmroq-unsplash-1-scaled.jpg"
author: "Brajeshwar"
captured_at: "2026-10-10T16:10:19Z"
capture_tool: "hn-digest"
hn_id: 50034279
score: 1
comments: 0
posted_at: "2026-10-10T16:08:31Z"
tags:
  - hacker-news
---

# Microsoft skipped OpenAI's decision model and built its own on Alibaba's Qwen

- HN: [50034279](https://news.ycombinator.com/item?id=50034279)
- Source: [thenewstack.io](https://thenewstack.io/microsoft-decision-model-foundry/)
- Score: 1
- Comments: 0
- Posted: 2026-10-10T16:08:31Z

## Translation

Title: Microsoft skipped OpenAI's decision model and built its own on Alibaba's Qwen
Article title: Microsoft skipped OpenAI’s decision model and built its own on Alibaba’s Qwen - The New Stack
Description: Microsoft’s Decision-1 matches Jev’s price and skips OpenAI’s model. Here’s who it’s built for and what it still leaves out for developers building agents.

Article text:
Microsoft skipped OpenAI’s decision model and built its own on Alibaba’s Qwen - The New Stack
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
Finla
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
What Kubernetes’ "monolith" lesson means for AI agent harnesses
Oct 2nd 2026 12:31pm, by
Bill Doerrfeld
OpenTelemetry and Prometheus are getting along. What’s still missing?
Sep 25th 2026 12:56pm, by
Bill Doerrfeld
Kubernetes v1.37 brings 67 enhancements. Which matter for operators?
Sep 11th 2026 1:40pm, by
Bill Doerrfeld
Greptile, Cursor, and Devin agree that agents should run their code. What they run it against matters.
Jun 27th 2026 11:00am, by
Arjun Iyer
Agentic development hinges on verification. For cloud-native software, that is a runtime problem.
Jun 11th 2026 10:00am, by
Arjun Iyer
Amazon ECS now auto-repairs failing GPUs and instances. Here's why it matters for SREs.
Oct 9th 2026 9:00am, by
Anirudh Aithal
K3s vs. K8s: When lightweight Kubernetes distros win (and when they d
[truncated]
Microsoft on Friday launched Microsoft-Decision-1 in Foundry, three days after OpenAI opened its Decisions API to every developer in public beta. Microsoft post-trained the decision model on Alibaba’s Qwen3.5-9B rather than anything its partner OpenAI built, though it plans to rebase it soon on its own MAI models as well as OpenAI’s.
Decision-1 costs $0.042 per million input tokens with free output, exactly what TypeSafe charges for Jev , and it arrived the same day TypeSafe announced an $870 million Series A at a $7.5 billion valuation led by a16z.
Microsoft is about three and a half weeks behind the startup that started this category, and in that time OpenAI, Upstage, Perplexity, Cloudflare, and AWS have all shipped decision models of their own .
TypeSafe says nearly 30% of the Fortune 500 has tried Jev, though it hasn’t named any of them, and TypeSafe CEO Diogo Almeida posted on X that “29.4% of the Fortune 500 showed up” in the three weeks since launch.
launched a model
accidentally served trillions of tokens a day
29.4% of the fortune 500 showed up
raised a really big series A from @a16z
it's been 3 weeks
we would like to sleep now pic.twitter.com/skpaLPd2Bz
Those are the same companies Microsoft sells Azure to, which makes TypeSafe’s early traction hard for Microsoft to ignore.
Microsoft’s first customer is itself
Microsoft Chairman and CEO Satya Nadella announced the model on Friday on X, writing, “We’re already testing it across Microsoft,” and four internal teams back him up.
Introducing Microsoft-Decision-1, our new model for fast decision-making.
It delivers top performance on structured decision tasks, outperforming both LLMs and other decision models in latency and quality.
We’re already testing it across Microsoft for everything from incident… pic.twitter.com/M3fAOD90r6
Xbox Research used it to sort more than 10,000 pieces of player feedback, the Copilot team graded chat and agent responses with it, on-call engineers used it to pull context during live incidents, and Microsoft Discovery used it to score experiments before an agent replans.
Microsoft’s numbers show it’s more than 14 times as fast as GPT-6 Sol at a fraction of the cost for Xbox, and 46 times more consistent in Discovery.
But the model may have a bigger job in mind. On Wednesday, the company said GitHub Copilot will soon decide when to run a task on-device and when to send it to cloud-scale models , though it hasn’t disclosed what Copilot sends to the cloud.
Model routing is among the use cases Microsoft lists for Decision-1, but the company hasn’t said whether the model will make those calls for Copilot. With routing decisions to make across Copilot, GitHub, and Xbox, Microsoft has an incentive to handle them in-house.
With routing decisions to make across Copilot, GitHub, and Xbox, Microsoft has an incentive to handle them in-house.
Cognition’s vice president of engineering, Jared Palmer , spent about $95 in Modal H100 time porting his open-source Kev models to Qwen3.5, and Cloudflare built Clef-flash on the same Qwen3.5-9B base Microsoft picked.
Jev’s price is becoming the going rate: Palmer lists Kev-4B on OpenRouter at $0.042 per million input tokens, Perplexity charges $0.02, and OpenAI charges more than double Jev’s rate. At those prices, revenue from individual decision calls is minimal, and Microsoft’s bigger opportunity is keeping agent traffic, including the generative calls around each decision, running through Foundry. Plus, the OpenRouter listing could bring developers who aren’t using Azure into that ecosystem.
Achint Srivastava , vice president of software engineering in Microsoft’s Office of the CTO, introduced Decision-1 as a way to add decision-making to existing applications, agents and workflows “in a secure, trusted environment.”
This shows where Decision-1 falls short of its rivals. Its Foundry listing says it accepts up to 32,768 tokens of text and returns JSON, but it doesn’t support images, unlike OpenAI’s Decisions API and Cloudflare’s Clef, which uses a vision encoder.
While Cloudflare released Clef under Apache 2.0 , Microsoft has yet to announce open weights. AWS, Upstage, and Ollama have also adopted TypeSafe’s System One API , now a common interface across the category, but Microsoft hasn’t said whether Decision-1 is fully compatible, though its Foundry sample code calls a /systemone endpoint.
Calibration under adversarial pressure
The company says Decision-1’s probabilities are calibrated, meaning a 90% prediction should be right about nine times in 10 on representative cases, but research on Jev shows how far a confident score can drift when the input is written to mislead. Microsoft’s own Foundry documentation advises customers to validate calibration on their own data.
The company says Decision-1’s probabilities are calibrated, meaning a 90% prediction should be right about nine times in 10 on representative cases.
In the JevOut preprint , USC computer science researcher Zixiang Xu and his co-authors found that short, natural-sounding additions to the context flipped 312 of Jev’s 508 initially correct decisions. In 229 cases, Jev assigned at least 70% probability to the wrong answer. Three other scoring systems showed flip rates between 64.9% and 73.2% under the same testing approach.
Microsoft tested Decision-1 with eight kinds of perturbations, including reordered options and paraphrased descriptions, which changed 1.3% of its answers on average. JevOut didn’t test Decision-1, so it’s unclear how the model would hold up against similar attacks. Microsoft also hasn’t confirmed whether Decision-1 fully supports the System One API competitors adopted, leaving questions about interoperability and how much trust to place in its confidence scores.
Amanda Caswell is an AI journalist, certified prompt engineer, and technology commentator whose work and expertise have been featured on Fox News and CBS News. She covers artificial intelligence, developer tools, foundation models, and emerging technologies, with a particular focus...
Read more from Amanda Caswell
SHARE THIS STORY
-->
TRENDING STORIES
TNS owner Insight Partners is an investor in: OpenAI.
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

Microsoft’s Decision-1 matches Jev’s price and skips OpenAI’s model. Here’s who it’s built for and what it still leaves out for developers building agents.

Microsoft skipped OpenAI’s decision model and built its own on Alibaba’s Qwen - The New Stack
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
Finla
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
What Kubernetes’ "monolith" lesson means for AI agent harnesses
Oct 2nd 2026 12:31pm, by
Bill Doerrfeld
OpenTelemetry and Prometheus are getting along. What’s still missing?
Sep 25th 2026 12:56pm, by
Bill Doerrfeld
Kubernetes v1.37 brings 67 enhancements. Which matter for operators?
Sep 11th 2026 1:40pm, by
Bill Doerrfeld
Greptile, Cursor, and Devin agree that agents should run their code. What they run it against matters.
Jun 27th 2026 11:00am, by
Arjun Iyer
Agentic development hinges on verification. For cloud-native software, that is a runtime problem.
Jun 11th 2026 10:00am, by
Arjun Iyer
Amazon ECS now auto-repairs failing GPUs and instances. Here's why it matters for SREs.
Oct 9th 2026 9:00am, by
Anirudh Aithal
K3s vs. K8s: When lightweight Kubernetes distros win (and when they d
[truncated]
Microsoft on Friday launched Microsoft-Decision-1 in Foundry, three days after OpenAI opened its Decisions API to every developer in public beta. Microsoft post-trained the decision model on Alibaba’s Qwen3.5-9B rather than anything its partner OpenAI built, though it plans to rebase it soon on its own MAI models as well as OpenAI’s.
Decision-1 costs $0.042 per million input tokens with free output, exactly what TypeSafe charges for Jev , and it arrived the same day TypeSafe announced an $870 million Series A at a $7.5 billion valuation led by a16z.
Microsoft is about three and a half weeks behind the startup that started this category, and in that time OpenAI, Upstage, Perplexity, Cloudflare, and AWS have all shipped decision models of their own .
TypeSafe says nearly 30% of the Fortune 500 has tried Jev, though it hasn’t named any of them, and TypeSafe CEO Diogo Almeida posted on X that “29.4% of the Fortune 500 showed up” in the three weeks since launch.
launched a model
accidentally served trillions of tokens a day
29.4% of the fortune 500 showed up
raised a really big series A from @a16z
it's been 3 weeks
we would like to sleep now pic.twitter.com/skpaLPd2Bz
Those are the same companies Microsoft sells Azure to, which makes TypeSafe’s early traction hard for Microsoft to ignore.
Microsoft’s first customer is itself
Microsoft Chairman and CEO Satya Nadella announced the model on Friday on X, writing, “We’re already testing it across Microsoft,” and four internal teams back him up.
Introducing Microsoft-Decision-1, our new model for fast decision-making.
It delivers top performance on structured decision tasks, outperforming both LLMs and other decision models in latency and quality.
We’re already testing it across Microsoft for everything from incident… pic.twitter.com/M3fAOD90r6
Xbox Research used it to sort more than 10,000 pieces of player feedback, the Copilot team graded chat and agent responses with it, on-call engineers used it to pull context during live incidents, and Microsoft Discovery used it to score experiments before an agent replans.
Microsoft’s numbers show it’s more than 14 times as fast as GPT-6 Sol at a fraction of the cost for Xbox, and 46 times more consistent in Discovery.
But the model may have a bigger job in mind. On Wednesday, the company said GitHub Copilot will soon decide when to run a task on-device and when to send it to cloud-scale models , though it hasn’t disclosed what Copilot sends to the cloud.
Model routing is among the use cases Microsoft lists for Decision-1, but the company hasn’t said whether the model will make those calls for Copilot. With routing decisions to make across Copilot, GitHub, and Xbox, Microsoft has an incentive to handle them in-house.
With routing decisions to make across Copilot, GitHub, and Xbox, Microsoft has an incentive to handle them in-house.
Cognition’s vice president of engineering, Jared Palmer , spent about $95 in Modal H100 time porting his open-source Kev models to Qwen3.5, and Cloudflare built Clef-flash on the same Qwen3.5-9B base Microsoft picked.
Jev’s price is becoming the going rate: Palmer lists Kev-4B on OpenRouter at $0.042 per million input tokens, Perplexity charges $0.02, and OpenAI charges more than double Jev’s rate. At those prices, revenue from individual decision calls is minimal, and Microsoft’s bigger opportunity is keeping agent traffic, including the generative calls around each decision, running through Foundry. Plus, the OpenRouter listing could bring developers who aren’t using Azure into that ecosystem.
Achint Srivastava , vice president of software engineering in Microsoft’s Office of the CTO, introduced Decision-1 as a way to add decision-making to existing applications, agents and workflows “in a secure, trusted environment.”
This shows where Decision-1 falls short of its rivals. Its Foundry listing says it accepts up to 32,768 tokens of text and returns JSON, but it doesn’t support images, unlike OpenAI’s Decisions API and Cloudflare’s Clef, which uses a vision encoder.
While Cloudflare released Clef under Apache 2.0 , Microsoft has yet to announce open weights. AWS, Upstage, and Ollama have also adopted TypeSafe’s System One API , now a common interface across the category, but Microsoft hasn’t said whether Decision-1 is fully compatible, though its Foundry sample code calls a /systemone endpoint.
Calibration under adversarial pressure
The company says Decision-1’s probabilities are calibrated, meaning a 90% prediction should be right about nine times in 10 on representative cases, but research on Jev shows how far a confident score can drift when the input is written to mislead. Microsoft’s own Foundry documentation advises customers to validate calibration on their own data.
The company says Decision-1’s probabilities are calibrated, meaning a 90% prediction should be right about nine times in 10 on representative cases.
In the JevOut preprint , USC computer science researcher Zixiang Xu and his co-authors found that short, natural-sounding additions to the context flipped 312 of Jev’s 508 initially correct decisions. In 229 cases, Jev assigned at least 70% probability to the wrong answer. Three other scoring systems showed flip rates between 64.9% and 73.2% under the same testing approach.
Microsoft tested Decision-1 with eight kinds of perturbations, including reordered options and paraphrased descriptions, which changed 1.3% of its answers on average. JevOut didn’t test Decision-1, so it’s unclear how the model would hold up against similar attacks. Microsoft also hasn’t confirmed whether Decision-1 fully supports the System One API competitors adopted, leaving questions about interoperability and how much trust to place in its confidence scores.
Amanda Caswell is an AI journalist, certified prompt engineer, and technology commentator whose work and expertise have been featured on Fox News and CBS News. She covers artificial intelligence, developer tools, foundation models, and emerging technologies, with a particular focus...
Read more from Amanda Caswell
SHARE THIS STORY
-->
TRENDING STORIES
TNS owner Insight Partners is an investor in: OpenAI.
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
