---
source: "https://thenewstack.io/coding-agents-leaked-screenshots/"
hn_url: "https://news.ycombinator.com/item?id=49921821"
title: "AI coding agents leaked 13,000 screenshots, and nobody hacked them"
article_title: "AI coding agents leaked 13,000 screenshots, and nobody hacked them. - The New Stack"
image: "https://cdn.thenewstack.io/media/2026/09/669492b7-mariola-grobelska-mm0ratz7ogg-unsplash-scaled.jpg"
author: "Brajeshwar"
captured_at: "2026-10-01T14:15:42Z"
capture_tool: "hn-digest"
hn_id: 49921821
score: 1
comments: 0
posted_at: "2026-10-01T14:00:21Z"
tags:
  - hacker-news
---

# AI coding agents leaked 13,000 screenshots, and nobody hacked them

- HN: [49921821](https://news.ycombinator.com/item?id=49921821)
- Source: [thenewstack.io](https://thenewstack.io/coding-agents-leaked-screenshots/)
- Score: 1
- Comments: 0
- Posted: 2026-10-01T14:00:21Z

## Translation

Title: AI coding agents leaked 13,000 screenshots, and nobody hacked them
Article title: AI coding agents leaked 13,000 screenshots, and nobody hacked them. - The New Stack
Description: AI coding agents leaked 13,000 internal screenshots from 300+ organizations by creating public GitHub repos. Here's how it happened and how to stop it.

Article text:
AI coding agents leaked 13,000 screenshots, and nobody hacked them. - The New Stack
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
AI coding agents trying to work around a limitation in GitHub’s command-line tool ended up publishing more than 13,000 internal images to public repositories, according to an incident report Glow Labs released this week. The company calls the incident PixelLeak.
Developers at more than 300 organizations were affected, including one of the world’s largest tech companies, a frontier AI lab, a major enterprise software vendor, and a Fortune 500 travel company, along with teams in cloud, healthcare, fintech, and government, and the exposed images are spread across more than 900 repositories.
None of it came from an attack, Glow Labs says the agents were completing tasks their developers had assigned them.
None of it came from an attack, the agents were completing tasks their developers had assigned them.
The trigger was one of the most routine steps in frontend work. A developer finishes a UI change and asks the agent to attach before-and-after screenshots to the pull request, but GitHub’s image attachment feature was built for people working in the web interface, while coding agents operate through the text-based CLI, and GitHub’s CLI didn’t get an image attachment option until version 2.99.0 on September 1. When the agent couldn’t attach the images that way, it looked for another route to make them visible to reviewers.
Glow reproduced the behavior in its lab by asking an agent running Anthropic’s Claude Opus 5 in Claude Code to change the header color on a private Minesweeper project. The agent created a new public repository and pinned the screenshots to a commit there so reviewers could see them from the private pull request.
“GitHub cannot render images from a private repo in a PR description — its image proxy fetches anonymously, so anything committed here (branch, release asset, whatever) shows up broken for reviewers,” the agent reasoned.
“The only way to satisfy both ‘reviewers see the images’ and ‘nothing but index.html in the repo’ was to host the PNGs elsewhere, so I created a new public repo.” Glow says that reasoning was representative of what it found at many of the affected organizations.
“The only way to satisfy both ‘reviewers see the images’ and ‘nothing but index.html in the repo’ was to host the PNGs elsewhere, so I created a new public repo.”
The exposed material went well beyond interface tweaks. At a manufacturer with more than 100,000 employees, an agent working on an internal billing screen published screenshots to a public repository under the developer’s personal GitHub account. Those images included billing records from a utility company involved in the fix.
Because the repository lived under the employee’s personal GitHub account rather than the company’s organization, the security team never spotted it, and the images were still public when Glow made contact.
Because the repository lived under the employee’s personal GitHub account rather than the company’s organization, the security team never spotted it, and the images were still public when Glow made contact.
That case reflects a broader pattern, with Glow finding that 93% of the images were stored in repositories under employees’ personal usernames, putting them outside the reach of scans focused on company GitHub organizations.
Even where the images were visible, the tools teams rely on to catch leaks, such as secret scanners and static analysis, analyze code and text rather than image content, allowing screenshots of internal consoles to pass through undetected.
Roughly a third of affected organizations had developers using gitshot, an unvetted open-source tool for publishing screenshots during code review, and at several large companies, agents discovered the tool and used it on their own.
Glow found more than 100 public accounts exposing internal work through _gitshot tags, including development work from a frontier AI lab and, at one financial services firm, an internal treasury and settlement console, a withdrawal screen naming an institutional client and two screen recordings of its money-movement console.
Coding agents can install tools and packages without anyone on the team vetting them first , and gitshot shows what can happen when one of those tools gives an agent a path around existing controls.
The incident became systemic once the workaround became a reusable instruction. At one software vendor, agents serving multiple engineers began publishing review screenshots publicly in early July, and within a week more than a dozen of them had encoded the approach as a skill applied to every development ticket.
Running that skill, the agents uploaded more than a thousand screenshots and screen recordings of the company’s product, along with written summaries of features that were weeks or months from release.
Agent skills have already become a supply chain risk in their own right , and this case shows that a skill doesn’t need to be malicious to spread risk across an engineering team when a mistaken one propagates just as efficiently. Glow began notifying affected organizations on September 9, 2026, and believes others are also affected.
For teams running coding agents, triage comes first. Glow recommends starting with everyone who commits to your private repositories, including former employees, and reviewing their personal accounts; then checking releases and gists in addition to file listings, since images attached to a release can make the file view look empty. Remove anything that turns up everywhere it exists, and rotate any credentials or other secrets legible in the images.
Remove anything that turns up everywhere it exists, and rotate any credentials or other secrets legible in the images.
After triage, the focus shifts to the agents and tools developers are actually using. Glow recommends removing tools like gitshot that haven’t gone through a security review, keeping git tooling up to date and requiring approval before an agent takes potentially risky actions.
Teams should also review the shared rules and instruction files their agents load, since that’s how a one-off workaround can become something agents repeat automatically.
Runtime controls for coding agents
The control Glow says actually stops this behavior operates at runtime. Glow, which actually sells endpoint runtime protection, recommends a pre-execution hook that blocks or holds for approval any attempt to create a new public repository, push to a personal account instead of the company’s organization, push to a gist, or switch a repository from private to public.
These gates work best when they sit outside the agent and evaluate the action itself , regardless of whether it came from careful reasoning or an injected prompt. An agent that cannot create a public repository has no way to improvise this particular escape.
Amanda Caswell is an AI journalist, certified prompt engineer, and technology commentator whose work and expertise have been featured on Fox News and CBS News. She covers artificial intelligence, developer tools, foundation models, and emerging technologies, with a particular focus...
Read more from Amanda Caswell
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
ENGINEE

[truncated]

## Original Extract

AI coding agents leaked 13,000 internal screenshots from 300+ organizations by creating public GitHub repos. Here's how it happened and how to stop it.

AI coding agents leaked 13,000 screenshots, and nobody hacked them. - The New Stack
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
AI coding agents trying to work around a limitation in GitHub’s command-line tool ended up publishing more than 13,000 internal images to public repositories, according to an incident report Glow Labs released this week. The company calls the incident PixelLeak.
Developers at more than 300 organizations were affected, including one of the world’s largest tech companies, a frontier AI lab, a major enterprise software vendor, and a Fortune 500 travel company, along with teams in cloud, healthcare, fintech, and government, and the exposed images are spread across more than 900 repositories.
None of it came from an attack, Glow Labs says the agents were completing tasks their developers had assigned them.
None of it came from an attack, the agents were completing tasks their developers had assigned them.
The trigger was one of the most routine steps in frontend work. A developer finishes a UI change and asks the agent to attach before-and-after screenshots to the pull request, but GitHub’s image attachment feature was built for people working in the web interface, while coding agents operate through the text-based CLI, and GitHub’s CLI didn’t get an image attachment option until version 2.99.0 on September 1. When the agent couldn’t attach the images that way, it looked for another route to make them visible to reviewers.
Glow reproduced the behavior in its lab by asking an agent running Anthropic’s Claude Opus 5 in Claude Code to change the header color on a private Minesweeper project. The agent created a new public repository and pinned the screenshots to a commit there so reviewers could see them from the private pull request.
“GitHub cannot render images from a private repo in a PR description — its image proxy fetches anonymously, so anything committed here (branch, release asset, whatever) shows up broken for reviewers,” the agent reasoned.
“The only way to satisfy both ‘reviewers see the images’ and ‘nothing but index.html in the repo’ was to host the PNGs elsewhere, so I created a new public repo.” Glow says that reasoning was representative of what it found at many of the affected organizations.
“The only way to satisfy both ‘reviewers see the images’ and ‘nothing but index.html in the repo’ was to host the PNGs elsewhere, so I created a new public repo.”
The exposed material went well beyond interface tweaks. At a manufacturer with more than 100,000 employees, an agent working on an internal billing screen published screenshots to a public repository under the developer’s personal GitHub account. Those images included billing records from a utility company involved in the fix.
Because the repository lived under the employee’s personal GitHub account rather than the company’s organization, the security team never spotted it, and the images were still public when Glow made contact.
Because the repository lived under the employee’s personal GitHub account rather than the company’s organization, the security team never spotted it, and the images were still public when Glow made contact.
That case reflects a broader pattern, with Glow finding that 93% of the images were stored in repositories under employees’ personal usernames, putting them outside the reach of scans focused on company GitHub organizations.
Even where the images were visible, the tools teams rely on to catch leaks, such as secret scanners and static analysis, analyze code and text rather than image content, allowing screenshots of internal consoles to pass through undetected.
Roughly a third of affected organizations had developers using gitshot, an unvetted open-source tool for publishing screenshots during code review, and at several large companies, agents discovered the tool and used it on their own.
Glow found more than 100 public accounts exposing internal work through _gitshot tags, including development work from a frontier AI lab and, at one financial services firm, an internal treasury and settlement console, a withdrawal screen naming an institutional client and two screen recordings of its money-movement console.
Coding agents can install tools and packages without anyone on the team vetting them first , and gitshot shows what can happen when one of those tools gives an agent a path around existing controls.
The incident became systemic once the workaround became a reusable instruction. At one software vendor, agents serving multiple engineers began publishing review screenshots publicly in early July, and within a week more than a dozen of them had encoded the approach as a skill applied to every development ticket.
Running that skill, the agents uploaded more than a thousand screenshots and screen recordings of the company’s product, along with written summaries of features that were weeks or months from release.
Agent skills have already become a supply chain risk in their own right , and this case shows that a skill doesn’t need to be malicious to spread risk across an engineering team when a mistaken one propagates just as efficiently. Glow began notifying affected organizations on September 9, 2026, and believes others are also affected.
For teams running coding agents, triage comes first. Glow recommends starting with everyone who commits to your private repositories, including former employees, and reviewing their personal accounts; then checking releases and gists in addition to file listings, since images attached to a release can make the file view look empty. Remove anything that turns up everywhere it exists, and rotate any credentials or other secrets legible in the images.
Remove anything that turns up everywhere it exists, and rotate any credentials or other secrets legible in the images.
After triage, the focus shifts to the agents and tools developers are actually using. Glow recommends removing tools like gitshot that haven’t gone through a security review, keeping git tooling up to date and requiring approval before an agent takes potentially risky actions.
Teams should also review the shared rules and instruction files their agents load, since that’s how a one-off workaround can become something agents repeat automatically.
Runtime controls for coding agents
The control Glow says actually stops this behavior operates at runtime. Glow, which actually sells endpoint runtime protection, recommends a pre-execution hook that blocks or holds for approval any attempt to create a new public repository, push to a personal account instead of the company’s organization, push to a gist, or switch a repository from private to public.
These gates work best when they sit outside the agent and evaluate the action itself , regardless of whether it came from careful reasoning or an injected prompt. An agent that cannot create a public repository has no way to improvise this particular escape.
Amanda Caswell is an AI journalist, certified prompt engineer, and technology commentator whose work and expertise have been featured on Fox News and CBS News. She covers artificial intelligence, developer tools, foundation models, and emerging technologies, with a particular focus...
Read more from Amanda Caswell
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
ENGINEE

[truncated]
