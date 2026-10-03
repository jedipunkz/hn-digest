---
source: "https://thenewstack.io/kubecon-agent-harness-koordinator/"
hn_url: "https://news.ycombinator.com/item?id=49944432"
title: "What Kubernetes' \"monolith\" lesson means for AI agent harnesses"
article_title: "What Kubernetes’ \"monolith\" lesson means for AI agent harnesses - The New Stack"
image: "https://cdn.thenewstack.io/media/2026/10/de419750-growtika-qpkdga-kdik-unsplash-scaled.jpg"
author: "Brajeshwar"
captured_at: "2026-10-03T15:03:55Z"
capture_tool: "hn-digest"
hn_id: 49944432
score: 2
comments: 0
posted_at: "2026-10-03T14:12:07Z"
tags:
  - hacker-news
---

# What Kubernetes' "monolith" lesson means for AI agent harnesses

- HN: [49944432](https://news.ycombinator.com/item?id=49944432)
- Source: [thenewstack.io](https://thenewstack.io/kubecon-agent-harness-koordinator/)
- Score: 2
- Comments: 0
- Posted: 2026-10-03T14:12:07Z

## Translation

Title: What Kubernetes' "monolith" lesson means for AI agent harnesses
Article title: What Kubernetes’ "monolith" lesson means for AI agent harnesses - The New Stack
Description: In this edition of our Road to KubeCon series, Craig McLuckie wants agent harnesses off the laptop, and Koordinator pushes GPU allocation past 95%.

Article text:
What Kubernetes’ "monolith" lesson means for AI agent harnesses - The New Stack
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
K3s vs. K8s: When lightweight Kubernetes distros win (and when they don't)
Oct 2nd 2026 9:00am, by
Peter Smails
How buildpacks help enterprises finally operate container security controls at scale
[truncated]
Welcome to another edition of Road to KubeCon , your source for everything Kubernetes and cloud-native as we count down the days to KubeCon + CloudNativeCon NA , happening November 9-12 in Salt Lake City, Utah.
This week, we take a look at the state of agent harnesses and what needs to evolve. We also feature HPE’s thoughts on Kubernetes visibility, more sci-fi agentic DevOps, and an impressive scheduler that shows Kubernetes’ default scheduler who’s the boss of GPU utilization.
That plus news from CNCF: what to expect at this year’s ArgoCon, a great opportunity to improve your open source project’s security hygiene, Atlassian’s stack for slashing event-to-metric latency, and a special gift for TNS readers.
HPE shares tips for Kubernetes visibility on TNS
Getting Kubernetes running is a milestone. Knowing who owns the next upgrade, the access request, or failed recovery is an ongoing journey.
This week, we published “ A live Kubernetes cluster can still have an ownership gap ,” the second installment in Chris J. Preimesberger’s four-part, HPE-sponsored series. It examines how platform and application teams divide responsibilities after launch, from configuration drift and security policies to upgrade validation and recovery drills. A healthy cluster doesn’t necessarily mean a healthy application — and someone needs to own that gap.
Hewlett Packard Enterprise (HPE) is a presenting sponsor of Road to KubeCon. HPE Software helps IT organizations modernize infrastructure, streamline operations, and accelerate AI initiatives across hybrid, multi-vendor environments.
Missed the opener? Part one explores Kubernetes self-service : how developers can get approved environments without waiting through ticket queues, while platform teams retain responsibility for access, costs, and lifecycle controls. Together, the articles ask a practical question: How do you give developers more independence while making operational accountability clear?
Next, the series turns to diagnosing slow applications when Kubernetes looks healthy, then to measuring AI inference performance. Both will explore the visibility teams need as their workloads become more demanding on the road to KubeCon.
CNCF offers a 10% discount to TNS readers
This week, The New Stack readers get a special treat from Cloud Native Computing Foundation (CNCF). If you’re planning to attend KubeCon + CloudNativeCon NA in Salt Lake City, use the discount code KCNA26MED10 when you register at this link for 10% off your admission.
We’re only 38 days out til KubeCon (or four Road to KubeCon editions out if you count in columns), so better act soon to register and plan your trip.
Agent harnesses go cloud-native
The agent “harness” quickly became a catch-all for everything that surrounds an AI agent: the context, filesystem, subagents, permissions, and more. However, Craig McLuckie , founder and CEO of Stacklok , says it’s not enough.
To him, most harnesses are too local and built to serve a single developer at a laptop. The typical harness doesn’t scale well enough to serve hundreds of sessions. Sessions break, and you can’t easily move the experience between clients or devices.
This week, he takes to the CNCF blog to argue for a cloud-native agent harness. To him, that’s a distributed application that separates the agent loop from the infrastructure and services around it.
“Kubernetes taught this industry that a monolith in a container is still a monolith,” he writes. “The lesson applies to agents too.”
Koordinator boosts on-Kubernetes GPU allocation >95%
A case study published on Tuesday details how Zhuoyu Technology, a Chinese autonomous driving technology company, is dramatically improving Kubernetes utilization with Koordinator , a CNCF sandbox project for efficiently scheduling microservices, AI, and big data workloads.
Zhuoyu Technology runs autonomous driving workloads on Kubernetes-based environments but hit performance inefficiencies with the default Kubernetes scheduler, which capped allocation and utilization. By using Koordinator, the team pushed GPU allocation above 95% and overall GPU utilization above 55%.
The case study demonstrates how certain gaps in the default Kubernetes scheduler can lead to failed launches, low GPU utilization, stranded GPUs, and distributed-job scheduling problems. It also demonstrates how Koordinator is faring well in production environments.
CNCF and OpenSSF announce month-long challenge
Open Source Security Foundation (OpenSSF) and CNCF are teaming up to organize the Security Slam , a 30-day challenge that walks participants through using OpenSSF projects to improve their project’s security posture.
All open source projects are invited to participate. Write Eddie Knight and OpenSSF’s Stacey Potter , the Slam is “now taking advantage of new tools to greatly broaden the qualifications for participation.”
The challenge runs October 5 through November 6. Register here to get involved and follow the objectives as they’re announced. Complete the challenges, and you might just have a fancy award ready for you at the OpenSSF booth (#313) at KubeCon.
Atlassian’s cloud-native stack takes event-to-metric below 10 seconds
In incident detection and response, every second counts. On Wednesday, Deepak Biswas , senior engineering manager at Atlassian, shared a deep case study on the CNCF blog about Atlassian’s journey to show how far you can go to shave those seconds down.
The detection platform behind AutoHOT, its automated incident creation system, combines OpenTelemetry , Apache Kafka , and Apache Flink on Kubernetes. Operational telemetry tracks user actions across more than 10 cloud products serving millions of tenants, generating billions of events per day.
The headline says it all: They’ve reduced their event-to-metric metric (to be meta about it) from more than 40 seconds to under 10.
Yet, Biswas is honest: “It is not a success story with a bow on it.” They’re still working on fine-tuning. Recall, for instance, fell to 64% in August. Nevertheless, it’s a useful blueprint for others building automated incident detection and response workflows.
As Kubernetes evolves, so do the demands on the teams running it. Presenting sponsor HPE helps teams address that complexity with software spanning virtualization, cloud management, observability, and automation.
KubeCon NA will feature ArgoCon North America 2026 on the co-located day, Monday, Nov. 9. It’s a full-day, two-track event with practical ideas on improving software delivery, managing data and machine learning pipelines, and implementing progressive delivery.
In a post on the CNCF blog on Wednesday , ArgoCon co-chairs Dan Garfield, Christian Hernandez, and Katie Lamkin stress that the event comes as the community begins the visioning process for Argo CD 4.0.
They describe ArgoCon, whose schedule is live here , as “an opportunity to connect around what users are building today and where the projects are heading next.”
Cycle’s DevOps control plane gets sci-fi
“If you told me this existed just a couple of years ago, I would have thought ‘this is science fiction, and it shouldn’t be real,'” says head of engineering Alexander Mattoni in a feature announcement video this week, showing off a new remote MCP server for Cycle, the DevOps control plane.
The release essentially means Cycle users can provision, orchestrate, and manage workloads across multicloud and hybrid environments via natural language, using MCP-compatible AI assistants and coding tools.
It follows a string of agentic features being released in the cloud native industry that continue to abstract DevOps and put impressive capabilities into the prompt.

[truncated]

## Original Extract

In this edition of our Road to KubeCon series, Craig McLuckie wants agent harnesses off the laptop, and Koordinator pushes GPU allocation past 95%.

What Kubernetes’ "monolith" lesson means for AI agent harnesses - The New Stack
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
K3s vs. K8s: When lightweight Kubernetes distros win (and when they don't)
Oct 2nd 2026 9:00am, by
Peter Smails
How buildpacks help enterprises finally operate container security controls at scale
[truncated]
Welcome to another edition of Road to KubeCon , your source for everything Kubernetes and cloud-native as we count down the days to KubeCon + CloudNativeCon NA , happening November 9-12 in Salt Lake City, Utah.
This week, we take a look at the state of agent harnesses and what needs to evolve. We also feature HPE’s thoughts on Kubernetes visibility, more sci-fi agentic DevOps, and an impressive scheduler that shows Kubernetes’ default scheduler who’s the boss of GPU utilization.
That plus news from CNCF: what to expect at this year’s ArgoCon, a great opportunity to improve your open source project’s security hygiene, Atlassian’s stack for slashing event-to-metric latency, and a special gift for TNS readers.
HPE shares tips for Kubernetes visibility on TNS
Getting Kubernetes running is a milestone. Knowing who owns the next upgrade, the access request, or failed recovery is an ongoing journey.
This week, we published “ A live Kubernetes cluster can still have an ownership gap ,” the second installment in Chris J. Preimesberger’s four-part, HPE-sponsored series. It examines how platform and application teams divide responsibilities after launch, from configuration drift and security policies to upgrade validation and recovery drills. A healthy cluster doesn’t necessarily mean a healthy application — and someone needs to own that gap.
Hewlett Packard Enterprise (HPE) is a presenting sponsor of Road to KubeCon. HPE Software helps IT organizations modernize infrastructure, streamline operations, and accelerate AI initiatives across hybrid, multi-vendor environments.
Missed the opener? Part one explores Kubernetes self-service : how developers can get approved environments without waiting through ticket queues, while platform teams retain responsibility for access, costs, and lifecycle controls. Together, the articles ask a practical question: How do you give developers more independence while making operational accountability clear?
Next, the series turns to diagnosing slow applications when Kubernetes looks healthy, then to measuring AI inference performance. Both will explore the visibility teams need as their workloads become more demanding on the road to KubeCon.
CNCF offers a 10% discount to TNS readers
This week, The New Stack readers get a special treat from Cloud Native Computing Foundation (CNCF). If you’re planning to attend KubeCon + CloudNativeCon NA in Salt Lake City, use the discount code KCNA26MED10 when you register at this link for 10% off your admission.
We’re only 38 days out til KubeCon (or four Road to KubeCon editions out if you count in columns), so better act soon to register and plan your trip.
Agent harnesses go cloud-native
The agent “harness” quickly became a catch-all for everything that surrounds an AI agent: the context, filesystem, subagents, permissions, and more. However, Craig McLuckie , founder and CEO of Stacklok , says it’s not enough.
To him, most harnesses are too local and built to serve a single developer at a laptop. The typical harness doesn’t scale well enough to serve hundreds of sessions. Sessions break, and you can’t easily move the experience between clients or devices.
This week, he takes to the CNCF blog to argue for a cloud-native agent harness. To him, that’s a distributed application that separates the agent loop from the infrastructure and services around it.
“Kubernetes taught this industry that a monolith in a container is still a monolith,” he writes. “The lesson applies to agents too.”
Koordinator boosts on-Kubernetes GPU allocation >95%
A case study published on Tuesday details how Zhuoyu Technology, a Chinese autonomous driving technology company, is dramatically improving Kubernetes utilization with Koordinator , a CNCF sandbox project for efficiently scheduling microservices, AI, and big data workloads.
Zhuoyu Technology runs autonomous driving workloads on Kubernetes-based environments but hit performance inefficiencies with the default Kubernetes scheduler, which capped allocation and utilization. By using Koordinator, the team pushed GPU allocation above 95% and overall GPU utilization above 55%.
The case study demonstrates how certain gaps in the default Kubernetes scheduler can lead to failed launches, low GPU utilization, stranded GPUs, and distributed-job scheduling problems. It also demonstrates how Koordinator is faring well in production environments.
CNCF and OpenSSF announce month-long challenge
Open Source Security Foundation (OpenSSF) and CNCF are teaming up to organize the Security Slam , a 30-day challenge that walks participants through using OpenSSF projects to improve their project’s security posture.
All open source projects are invited to participate. Write Eddie Knight and OpenSSF’s Stacey Potter , the Slam is “now taking advantage of new tools to greatly broaden the qualifications for participation.”
The challenge runs October 5 through November 6. Register here to get involved and follow the objectives as they’re announced. Complete the challenges, and you might just have a fancy award ready for you at the OpenSSF booth (#313) at KubeCon.
Atlassian’s cloud-native stack takes event-to-metric below 10 seconds
In incident detection and response, every second counts. On Wednesday, Deepak Biswas , senior engineering manager at Atlassian, shared a deep case study on the CNCF blog about Atlassian’s journey to show how far you can go to shave those seconds down.
The detection platform behind AutoHOT, its automated incident creation system, combines OpenTelemetry , Apache Kafka , and Apache Flink on Kubernetes. Operational telemetry tracks user actions across more than 10 cloud products serving millions of tenants, generating billions of events per day.
The headline says it all: They’ve reduced their event-to-metric metric (to be meta about it) from more than 40 seconds to under 10.
Yet, Biswas is honest: “It is not a success story with a bow on it.” They’re still working on fine-tuning. Recall, for instance, fell to 64% in August. Nevertheless, it’s a useful blueprint for others building automated incident detection and response workflows.
As Kubernetes evolves, so do the demands on the teams running it. Presenting sponsor HPE helps teams address that complexity with software spanning virtualization, cloud management, observability, and automation.
KubeCon NA will feature ArgoCon North America 2026 on the co-located day, Monday, Nov. 9. It’s a full-day, two-track event with practical ideas on improving software delivery, managing data and machine learning pipelines, and implementing progressive delivery.
In a post on the CNCF blog on Wednesday , ArgoCon co-chairs Dan Garfield, Christian Hernandez, and Katie Lamkin stress that the event comes as the community begins the visioning process for Argo CD 4.0.
They describe ArgoCon, whose schedule is live here , as “an opportunity to connect around what users are building today and where the projects are heading next.”
Cycle’s DevOps control plane gets sci-fi
“If you told me this existed just a couple of years ago, I would have thought ‘this is science fiction, and it shouldn’t be real,'” says head of engineering Alexander Mattoni in a feature announcement video this week, showing off a new remote MCP server for Cycle, the DevOps control plane.
The release essentially means Cycle users can provision, orchestrate, and manage workloads across multicloud and hybrid environments via natural language, using MCP-compatible AI assistants and coding tools.
It follows a string of agentic features being released in the cloud native industry that continue to abstract DevOps and put impressive capabilities into the prompt.

[truncated]
