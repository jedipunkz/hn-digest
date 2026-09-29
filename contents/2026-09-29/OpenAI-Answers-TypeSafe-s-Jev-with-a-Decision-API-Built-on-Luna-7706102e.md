---
source: "https://thenewstack.io/openai-decision-api-luna/"
hn_url: "https://news.ycombinator.com/item?id=49896979"
title: "OpenAI Answers TypeSafe's Jev with a Decision API Built on Luna"
article_title: "OpenAI answers TypeSafe's Jev with a Decision API built on Luna - The New Stack"
image: "https://cdn.thenewstack.io/media/2026/09/177ab814-decisions-api.png"
author: "yawnxyz"
captured_at: "2026-09-29T17:45:50Z"
capture_tool: "hn-digest"
hn_id: 49896979
score: 2
comments: 0
posted_at: "2026-09-29T17:27:00Z"
tags:
  - hacker-news
---

# OpenAI Answers TypeSafe's Jev with a Decision API Built on Luna

- HN: [49896979](https://news.ycombinator.com/item?id=49896979)
- Source: [thenewstack.io](https://thenewstack.io/openai-decision-api-luna/)
- Score: 2
- Comments: 0
- Posted: 2026-09-29T17:27:00Z

## Translation

Title: OpenAI Answers TypeSafe's Jev with a Decision API Built on Luna
Article title: OpenAI answers TypeSafe's Jev with a Decision API built on Luna - The New Stack
Description: OpenAI's Decision API, built on its small Luna model, returns predefined answers with confidence scores in 150 milliseconds. Pricing is still unknown.

Article text:
OpenAI answers TypeSafe's Jev with a Decision API built on Luna - The New Stack
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
Fren
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
Vendor neutrality isn’t magic: A hard look at the OpenTelemetry ecosystem
May 29th 2026 10:00am, by
Adriana Villela and Josh Lee
How buildpacks help enterprises finally operate container security controls at scale
Sep 18th 2026 10:00am, by
Catherine Edelveis
Your container runs. Everything around it
[truncated]
With the sudden rise of TypeSafe’s Jev , decision models have become incredibly popular, so it’s not surprising that OpenAI is also announcing its take on this model type at its annual DevDay conference on Tuesday.
The company’s new Decisions API is based on the company’s Luna model — the smallest and most affordable model in its current lineup.
The Decisions API is likely a reaction to TypeSafe and Jev, and OpenAI probably rushed the announcement ahead of its DevDay, so for now, this is all OpenAI is sharing about the Decisions API.
An OpenAI spokesperson tells The New Stack that the company plans to share more “at broad rollout.”
For now, the new API is available in limited preview, and the broad release is planned for the coming days.
Predefined answers, real confidence scores
The core idea behind these decision models is that they are explicitly not chat models; instead, they return a set of predefined answers with confidence scores. Regular LLMs are not always very good at this. Their confidence scores, after all, are often a rough guess — and they burn quite a few tokens to get there.
This makes this kind of model ideal for classifying content, routing requests, or choosing an agent’s next action from a limited set of choices.
Decision models, however, can provide more realistic confidence scores, and they tend to return them extremely fast. OpenAI says its model returns results in 150 milliseconds, compared to GPT-6 Luna, which would take 1.6 seconds.
All the developer has to do is provide the questions, answers, and context.
Today, most teams handle this today with a regular chat model and a carefully worded prompt, asking it to pick from a list and, if they’re lucky, reading the token probabilities to get something that resembles a confidence score.
The alternative is to train a small classifier. That is fast and cheap but needs labeled data and a retraining run every time the label set changes.
A decision model basically sits between those two options. It takes new labels in the prompt but returns a score a developer can work with.
This isn’t a completely new concept for OpenAI. Its Moderation API has long returned per-category scores instead of prose, too, though in this API, the categories are pre-set by OpenAI, not the developer.
What’s still unclear is what the Decision API costs per call, how many candidate answers a single request can handle, and whether developers can tune it on their own data. Those details will decide whether this becomes a standard building block in agent frameworks or stays a niche tool next to the chat models.
Before joining The New Stack as its senior editor for AI, Frederic was the enterprise editor at TechCrunch, where he covered everything from the rise of the cloud and the earliest days of Kubernetes to the advent of quantum computing....
Read more from Frederic Lardinois
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

OpenAI's Decision API, built on its small Luna model, returns predefined answers with confidence scores in 150 milliseconds. Pricing is still unknown.

OpenAI answers TypeSafe's Jev with a Decision API built on Luna - The New Stack
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
Fren
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
Vendor neutrality isn’t magic: A hard look at the OpenTelemetry ecosystem
May 29th 2026 10:00am, by
Adriana Villela and Josh Lee
How buildpacks help enterprises finally operate container security controls at scale
Sep 18th 2026 10:00am, by
Catherine Edelveis
Your container runs. Everything around it
[truncated]
With the sudden rise of TypeSafe’s Jev , decision models have become incredibly popular, so it’s not surprising that OpenAI is also announcing its take on this model type at its annual DevDay conference on Tuesday.
The company’s new Decisions API is based on the company’s Luna model — the smallest and most affordable model in its current lineup.
The Decisions API is likely a reaction to TypeSafe and Jev, and OpenAI probably rushed the announcement ahead of its DevDay, so for now, this is all OpenAI is sharing about the Decisions API.
An OpenAI spokesperson tells The New Stack that the company plans to share more “at broad rollout.”
For now, the new API is available in limited preview, and the broad release is planned for the coming days.
Predefined answers, real confidence scores
The core idea behind these decision models is that they are explicitly not chat models; instead, they return a set of predefined answers with confidence scores. Regular LLMs are not always very good at this. Their confidence scores, after all, are often a rough guess — and they burn quite a few tokens to get there.
This makes this kind of model ideal for classifying content, routing requests, or choosing an agent’s next action from a limited set of choices.
Decision models, however, can provide more realistic confidence scores, and they tend to return them extremely fast. OpenAI says its model returns results in 150 milliseconds, compared to GPT-6 Luna, which would take 1.6 seconds.
All the developer has to do is provide the questions, answers, and context.
Today, most teams handle this today with a regular chat model and a carefully worded prompt, asking it to pick from a list and, if they’re lucky, reading the token probabilities to get something that resembles a confidence score.
The alternative is to train a small classifier. That is fast and cheap but needs labeled data and a retraining run every time the label set changes.
A decision model basically sits between those two options. It takes new labels in the prompt but returns a score a developer can work with.
This isn’t a completely new concept for OpenAI. Its Moderation API has long returned per-category scores instead of prose, too, though in this API, the categories are pre-set by OpenAI, not the developer.
What’s still unclear is what the Decision API costs per call, how many candidate answers a single request can handle, and whether developers can tune it on their own data. Those details will decide whether this becomes a standard building block in agent frameworks or stays a niche tool next to the chat models.
Before joining The New Stack as its senior editor for AI, Frederic was the enterprise editor at TechCrunch, where he covered everything from the rise of the cloud and the earliest days of Kubernetes to the advent of quantum computing....
Read more from Frederic Lardinois
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
