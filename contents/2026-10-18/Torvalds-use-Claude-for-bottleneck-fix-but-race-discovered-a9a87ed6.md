---
source: "https://thenewstack.io/torvalds-claude-code-kernel/"
hn_url: "https://news.ycombinator.com/item?id=50142323"
title: "Torvalds use Claude for bottleneck fix, but race discovered"
article_title: "Linus Torvalds let Claude Code fix the Linux page cache. Then reviewers found a race. - The New Stack"
image: "https://cdn.thenewstack.io/media/2026/10/39123855-dorothy-livelo-l1lklun853s-unsplash-scaled.jpg"
author: "PaulDavisThe1st"
captured_at: "2026-10-18T16:28:16Z"
capture_tool: "hn-digest"
hn_id: 50142323
score: 4
comments: 0
posted_at: "2026-10-18T16:08:48Z"
tags:
  - hacker-news
---

# Torvalds use Claude for bottleneck fix, but race discovered

- HN: [50142323](https://news.ycombinator.com/item?id=50142323)
- Source: [thenewstack.io](https://thenewstack.io/torvalds-claude-code-kernel/)
- Score: 4
- Comments: 0
- Posted: 2026-10-18T16:08:48Z

## Translation

Title: Torvalds use Claude for bottleneck fix, but race discovered
Article title: Linus Torvalds let Claude Code fix the Linux page cache. Then reviewers found a race. - The New Stack
Description: Linus Torvalds used Claude Code to build a Linux page cache fix that hit 23x read throughput, but kernel reviewers found a race before anything merged.

Article text:
Linus Torvalds let Claude Code fix the Linux page cache. Then reviewers found a race. - The New Stack
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
Fi
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
Linux creator and longtime lead maintainer Linus Torvalds turned to Anthropic’s Claude Code to tackle a kernel performance problem that developers have struggled with for more than two years. On Tuesday, he submitted a 10-patch series that delivered a 23-fold throughput improvement in a synthetic benchmark running on his Ryzen Threadripper.
Torvalds had spent 18 months running a broken version of his own fix for the page-cache reference-counting bottleneck, while other developers had encountered problems with their approaches. The new patches eliminate folio reference-count operations from a common buffered-read path, although Torvalds warned that the benchmark gains would be considerably smaller in real-world workloads.
Reviewers have since identified a race condition and raised concerns about memory reclamation. Torvalds has handed further development to Kiryl Shutsemau , a Linux kernel developer who previously worked on the same bottleneck.
Torvalds had spent 18 months running a broken version of his own fix for the page-cache reference-counting bottleneck, while other developers had encountered problems with their approaches.
On Torvalds’ Core i7-8565U laptop, folio reference-count contention accounted for 99.92% of cross-core load HITMs in eight-thread tests, according to Hardware Busters .
The bottleneck was flagged in 2024 by Matthew Wilcox, an Oracle engineer and longtime Linux kernel memory-management developer, after it surfaced in a customer workload.
Claude Code writes the patches
On Sunday, Phoronix’s Michael Larabel spotted an experimental branch named noref-filemap-read-with-llm in Torvalds’ tree, which rarely contains anything outside the main development branch.
The eight-patch series, based on Linux 7.3-rc6, came with a clear account of what Claude Code did. Torvalds said the assistant wrote the implementation, tests, and initial documentation, and reviewed the code under his direction. He ran the benchmarks himself, despite having previously criticized developers for claiming AI writes 99% of their code .
Earlier approaches copied data without holding a folio reference and checked its validity afterward, leaving races involving reclamation and replacement. The subsequent 10-patch rcu-free-v3 series instead uses a new PG_rcu_free flag and RCU-delayed reclamation to protect the folio through the userspace copy, provided it doesn’t fault.
Earlier approaches copied data without holding a folio reference and checked its validity afterward, leaving races involving reclamation and replacement.
Threadripper benchmark results
The original eight-patch branch nearly tripled throughput on Torvalds’ laptop, pushing eight-thread performance from 10 million to as many as 33 million reads per second while cutting kernel cycles per read from roughly 2,950 to 725. The revised series delivered a much larger gain under heavy contention on his 64-thread Threadripper, where throughput jumped from 13.3 million to 304 million reads per second and cycles per read fell from nearly 20,000 to 735.
The original eight-patch branch nearly tripled throughput on Torvalds’ laptop, pushing eight-thread performance from 10 million to as many as 33 million reads per second while cutting kernel cycles per read from roughly 2,950 to 725.
The performance gains came with trade-offs. In the first implementation, delayed reclamation forced an expedited grace period whenever a CPU queued more than 1,024 pages, according to Hardware Busters . Reads into unmapped user pages also triggered an additional fault before falling back to the existing path, pushing a 16-byte single-thread read from about 4,150 to 6,550 cycles. Whether that regression persists in rcu-free-v3 remains unclear.
Maintainers question the reclamation design
Wilcox questioned the new PG_rcu_free flag and proposed using a page type instead, along with reclaiming inactive-list folios after an RCU grace period.
Torvalds rejected both approaches, arguing that the flag must be set while a folio can still be mapped into userspace and that inactive-list reclamation wouldn’t cover truncate and hole-punch operations. Shutsemau subsequently suggested restricting the flag to 64-bit systems, an approach supported by Arm engineer and Linux memory-management maintainer David Hildenbrand .
Shutsemau also identified a race in filemap_read_rcu_batch() , where folio replacement followed by an extension of i_size could expose data that had never been written. Torvalds recommended moving the file-size read earlier and consolidating the folio checks, which Shutsemau incorporated into a revised patch.
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

Linus Torvalds used Claude Code to build a Linux page cache fix that hit 23x read throughput, but kernel reviewers found a race before anything merged.

Linus Torvalds let Claude Code fix the Linux page cache. Then reviewers found a race. - The New Stack
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
Fi
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
Linux creator and longtime lead maintainer Linus Torvalds turned to Anthropic’s Claude Code to tackle a kernel performance problem that developers have struggled with for more than two years. On Tuesday, he submitted a 10-patch series that delivered a 23-fold throughput improvement in a synthetic benchmark running on his Ryzen Threadripper.
Torvalds had spent 18 months running a broken version of his own fix for the page-cache reference-counting bottleneck, while other developers had encountered problems with their approaches. The new patches eliminate folio reference-count operations from a common buffered-read path, although Torvalds warned that the benchmark gains would be considerably smaller in real-world workloads.
Reviewers have since identified a race condition and raised concerns about memory reclamation. Torvalds has handed further development to Kiryl Shutsemau , a Linux kernel developer who previously worked on the same bottleneck.
Torvalds had spent 18 months running a broken version of his own fix for the page-cache reference-counting bottleneck, while other developers had encountered problems with their approaches.
On Torvalds’ Core i7-8565U laptop, folio reference-count contention accounted for 99.92% of cross-core load HITMs in eight-thread tests, according to Hardware Busters .
The bottleneck was flagged in 2024 by Matthew Wilcox, an Oracle engineer and longtime Linux kernel memory-management developer, after it surfaced in a customer workload.
Claude Code writes the patches
On Sunday, Phoronix’s Michael Larabel spotted an experimental branch named noref-filemap-read-with-llm in Torvalds’ tree, which rarely contains anything outside the main development branch.
The eight-patch series, based on Linux 7.3-rc6, came with a clear account of what Claude Code did. Torvalds said the assistant wrote the implementation, tests, and initial documentation, and reviewed the code under his direction. He ran the benchmarks himself, despite having previously criticized developers for claiming AI writes 99% of their code .
Earlier approaches copied data without holding a folio reference and checked its validity afterward, leaving races involving reclamation and replacement. The subsequent 10-patch rcu-free-v3 series instead uses a new PG_rcu_free flag and RCU-delayed reclamation to protect the folio through the userspace copy, provided it doesn’t fault.
Earlier approaches copied data without holding a folio reference and checked its validity afterward, leaving races involving reclamation and replacement.
Threadripper benchmark results
The original eight-patch branch nearly tripled throughput on Torvalds’ laptop, pushing eight-thread performance from 10 million to as many as 33 million reads per second while cutting kernel cycles per read from roughly 2,950 to 725. The revised series delivered a much larger gain under heavy contention on his 64-thread Threadripper, where throughput jumped from 13.3 million to 304 million reads per second and cycles per read fell from nearly 20,000 to 735.
The original eight-patch branch nearly tripled throughput on Torvalds’ laptop, pushing eight-thread performance from 10 million to as many as 33 million reads per second while cutting kernel cycles per read from roughly 2,950 to 725.
The performance gains came with trade-offs. In the first implementation, delayed reclamation forced an expedited grace period whenever a CPU queued more than 1,024 pages, according to Hardware Busters . Reads into unmapped user pages also triggered an additional fault before falling back to the existing path, pushing a 16-byte single-thread read from about 4,150 to 6,550 cycles. Whether that regression persists in rcu-free-v3 remains unclear.
Maintainers question the reclamation design
Wilcox questioned the new PG_rcu_free flag and proposed using a page type instead, along with reclaiming inactive-list folios after an RCU grace period.
Torvalds rejected both approaches, arguing that the flag must be set while a folio can still be mapped into userspace and that inactive-list reclamation wouldn’t cover truncate and hole-punch operations. Shutsemau subsequently suggested restricting the flag to 64-bit systems, an approach supported by Arm engineer and Linux memory-management maintainer David Hildenbrand .
Shutsemau also identified a race in filemap_read_rcu_batch() , where folio replacement followed by an extension of i_size could expose data that had never been written. Torvalds recommended moving the file-size read earlier and consolidating the folio checks, which Shutsemau incorporated into a revised patch.
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
