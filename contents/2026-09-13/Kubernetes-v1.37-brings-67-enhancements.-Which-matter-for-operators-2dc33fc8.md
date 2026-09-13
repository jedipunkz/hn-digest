---
source: "https://thenewstack.io/kubecon-kubernetes-updates-security/"
hn_url: "https://news.ycombinator.com/item?id=49679663"
title: "Kubernetes v1.37 brings 67 enhancements. Which matter for operators?"
article_title: "Kubernetes v1.37 brings 67 enhancements. Which matter for operators? - The New Stack"
image: "https://cdn.thenewstack.io/media/2026/09/cf894579-growtika-6isceqbipmo-unsplash-scaled.jpg"
author: "geoffbp"
captured_at: "2026-09-13T03:30:40Z"
capture_tool: "hn-digest"
hn_id: 49679663
score: 1
comments: 0
posted_at: "2026-09-13T03:21:17Z"
tags:
  - hacker-news
---

# Kubernetes v1.37 brings 67 enhancements. Which matter for operators?

- HN: [49679663](https://news.ycombinator.com/item?id=49679663)
- Source: [thenewstack.io](https://thenewstack.io/kubecon-kubernetes-updates-security/)
- Score: 1
- Comments: 0
- Posted: 2026-09-13T03:21:17Z

## Translation

Title: Kubernetes v1.37 brings 67 enhancements. Which matter for operators?
Article title: Kubernetes v1.37 brings 67 enhancements. Which matter for operators? - The New Stack
Description: Kubernetes is gaining AI and operations tools. Road to KubeCon explores what’s changing and the access-control gap teams still need to address.

Article text:
Kubernetes v1.37 brings 67 enhancements. Which matter for operators? - The New Stack
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
How Jaeger hit 8.6× compression on 10 million spans with ClickHouse
May 24th 2026 11:00am, by
Mahad Zaryab
Your container runs. Everything around it shouldn't be your problem.
Aug 29th 2026 11:00am, by
Satej Sawant
Your container images are unsigned. In the AI era, that's a ticking
[truncated]
Welcome to the first edition of Road to KubeCon , where we’ll track the world of Kubernetes as we approach KubeCon + CloudNativeCon North America, November 9-12 in Salt Lake City.
This week, we’re catching up on recent developments across the Kubernetes universe, including Kubernetes v1.37 Garhwal, CNCF project graduations, HPE, AKS, and VMware updates, and why access control deserves more attention.
HPE talks Morpheus and Terraform updates
In a recent HPE Developer Community Meetup session , technologists Colin Taylor , Don Wake , and Eamonn O’Toole from HPE Hybrid Cloud dove deep into updates to HPE Morpheus , the platform for operating infrastructure as code for hybrid clouds.
Hewlett Packard Enterprise (HPE) is a presenting sponsor of Road to KubeCon. HPE Software helps IT organizations modernize infrastructure, streamline operations, and accelerate AI initiatives across hybrid, multi-vendor environments.
The major news is around the Morpheus Terraform Provider, whose functionality has now been converged into the HPE Terraform provider. HPE also released tfmigrator , a tool that automates migration from the standalone Morpheus provider to the unified HPE provider.
The session explored how HPE Morpheus and Terraform support infrastructure management across hybrid environments, including changes to the HPE Terraform provider and tools for migrating existing configurations.
If you’re using Morpheus and want to get into the weeds of the latest platform updates, or are just curious if someone named Morpheus will offer you a red or blue pill, definitely check out the latest community chat .
CNCF graduates Kubeflow, Karmada, Cloud Native Buildpacks
Cloud Native Computing Foundation (CNCF), the arm of the Linux Foundation that shepherds Kubernetes and countless other cloud-native open source projects, all replete with Kube-this and Kube-that branding and cuddly mascots (228 projects at the time of writing), announced a few major graduations in recent weeks.
For those unaware, “graduation” status means the project is highly mature, has completed security reviews, and has a vendor-neutral governance model in place to sustain it. That’s a good sign it’ll stick around for a while. A rare blessing for open-source .
Probably the most noteworthy recent graduation is Kubeflow , the platform for AI and ML training on Kubernetes, which has had 260 million PyPI downloads to date. “Graduation marks a critical milestone, cementing Kubeflow as a mature option for enterprise AI workloads on Kubernetes,” says CNCF CTO Chris Aniszczyk in the graduation announcement .
Karmada , another graduated project , is a multicluster, multi-cloud Kubernetes orchestration project. Its graduation is a win for those building cloud-agnostic, multi-cloud Kubernetes. Its latest release, v1.19, advances multi-component scheduling for distributed AI training jobs.
Lastly, the other big graduation announcement was for Cloud Native Buildpacks . The project, which can transform application code into OCI-compliant container images, joined CNCF as a sandbox project in 2018.
Kubernetes reaches new peaks with v1.37 Garhwal
The latest minor Kubernetes release, v1.37 , is here. It’s nicknamed Garhwal, as an homage to the snow-capped peaks of the Garhwal Himalaya mountain range.
v1.37 includes 67 enhancements: 16 stable, 23 beta, 27 alpha, and one deprecation. Notable features include completing resilient watch cache initialization, which can improve resilience for large clusters and help avoid control plane outages.
One interesting update: KYAML has now reached stable status. It’s billed as a solution to headaches with YAML , including whitespace sensitivity and the dreaded “ Norway Problem .” (I had no idea something as fundamental as YAML had so many issues, but I guess it does.)
KYAML should be able to help. Every KYAML file is still valid YAML, so don’t worry about rewriting anything for backward compatibility. Will KYAML become a more common way to write Kubernetes configuration? Time will tell.
Other notable updates include HorizontalPodAutoscaler scale to zero graduating to beta and being enabled by default. For workloads using object or external metrics, this enables pods to scale down to zero when idle. Other key updates include beta support for manifest-based admission control , and alpha support for pod-level checkpoint and restore.
As Kubernetes evolves, so do the demands on the teams running it. Presenting sponsor HPE helps teams address that complexity with software spanning virtualization, cloud management, observability and automation.
KubeCon travel-scholarship applications close soon: apply now
The schedule for KubeCon + CloudNativeCon North America 2026 is announced . As if the four-day agenda wasn’t jam-packed and mouth-watering enough, this year we’re getting a new AI inference and agentic track.
Thankfully, not everyone has to miss out on the fun. KubeCon offers a scholarship program intended to help fund travel and registration for those in underrepresented groups, or those without the means to do so otherwise.
The deadline to submit a travel funding request is this Sunday . Be sure to submit your request by Sunday, September 13, 11:59 p.m. Mountain Daylight Time (MDT). Registration applications don’t close until Sunday, October 4, 11:59 p.m. MDT.
Access control for Kubernetes finally makes the list
Kolawole Olowoporoku , CNCF Ambassador and senior platform engineer at Armada , is on the CNCF blog this week spotlighting an area that doesn’t always get much attention: identity and access control. He starts with a potent message: “Access control belongs on the same day-zero checklist as networking and storage. On most on-prem clusters, it never makes the list.”
Self-hosted Kubernetes includes authentication and authorization mechanisms, but teams must configure integration with an external identity provider. Without that integration, operators may rely on static client certificates or long-lived tokens.
Such credentials can create security risks when they remain valid longer than intended. Olowoporoku recommends authenticating through an OpenID Connect identity provider using a public client with PKCE. After login, kubectl sends the resulting ID token to the Kubernetes API server, which validates it and applies the configured access permissions.
VMware AI-ifies private cloud visibility
More news on the private cloud front: VMware Cloud Foundation (VCF) 9.1.1 adds new capabilities that help operators gain visibility into their environments.
One addition is enhanced observability into real-time Kubernetes operations, reducing standard five-minute polling intervals to two-second metric streaming. This can help operators detect short-lived pods, memory spikes, and transient performance bottlenecks that might otherwise go unnoticed.
The next major addition is a new AI Assistant for VCF. The conversational interface can help with troubleshooting and diagnostics, check the health of VCF environments, pinpoint root causes, and more. It’s one of many recent moves to add generative AI capabilities to Kubernetes and private cloud operations.
In the latest 2026-09-04 release notes , the Azure Kubernetes Service (AKS) team notes that the latest Kubernetes v1.37 preview is rolling out, with patches for previous versions now available.
Autoscaling for virtual machine node pools has reached general availability. New preview capabilities also give operators more flexibility in managing node pools throughout their lifecycle.
The world surrounding Kubernetes never sleeps. Here are some quick and interesting tidbits in other areas:
CNCF project owners should check out the lat

[truncated]

## Original Extract

Kubernetes is gaining AI and operations tools. Road to KubeCon explores what’s changing and the access-control gap teams still need to address.

Kubernetes v1.37 brings 67 enhancements. Which matter for operators? - The New Stack
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
How Jaeger hit 8.6× compression on 10 million spans with ClickHouse
May 24th 2026 11:00am, by
Mahad Zaryab
Your container runs. Everything around it shouldn't be your problem.
Aug 29th 2026 11:00am, by
Satej Sawant
Your container images are unsigned. In the AI era, that's a ticking
[truncated]
Welcome to the first edition of Road to KubeCon , where we’ll track the world of Kubernetes as we approach KubeCon + CloudNativeCon North America, November 9-12 in Salt Lake City.
This week, we’re catching up on recent developments across the Kubernetes universe, including Kubernetes v1.37 Garhwal, CNCF project graduations, HPE, AKS, and VMware updates, and why access control deserves more attention.
HPE talks Morpheus and Terraform updates
In a recent HPE Developer Community Meetup session , technologists Colin Taylor , Don Wake , and Eamonn O’Toole from HPE Hybrid Cloud dove deep into updates to HPE Morpheus , the platform for operating infrastructure as code for hybrid clouds.
Hewlett Packard Enterprise (HPE) is a presenting sponsor of Road to KubeCon. HPE Software helps IT organizations modernize infrastructure, streamline operations, and accelerate AI initiatives across hybrid, multi-vendor environments.
The major news is around the Morpheus Terraform Provider, whose functionality has now been converged into the HPE Terraform provider. HPE also released tfmigrator , a tool that automates migration from the standalone Morpheus provider to the unified HPE provider.
The session explored how HPE Morpheus and Terraform support infrastructure management across hybrid environments, including changes to the HPE Terraform provider and tools for migrating existing configurations.
If you’re using Morpheus and want to get into the weeds of the latest platform updates, or are just curious if someone named Morpheus will offer you a red or blue pill, definitely check out the latest community chat .
CNCF graduates Kubeflow, Karmada, Cloud Native Buildpacks
Cloud Native Computing Foundation (CNCF), the arm of the Linux Foundation that shepherds Kubernetes and countless other cloud-native open source projects, all replete with Kube-this and Kube-that branding and cuddly mascots (228 projects at the time of writing), announced a few major graduations in recent weeks.
For those unaware, “graduation” status means the project is highly mature, has completed security reviews, and has a vendor-neutral governance model in place to sustain it. That’s a good sign it’ll stick around for a while. A rare blessing for open-source .
Probably the most noteworthy recent graduation is Kubeflow , the platform for AI and ML training on Kubernetes, which has had 260 million PyPI downloads to date. “Graduation marks a critical milestone, cementing Kubeflow as a mature option for enterprise AI workloads on Kubernetes,” says CNCF CTO Chris Aniszczyk in the graduation announcement .
Karmada , another graduated project , is a multicluster, multi-cloud Kubernetes orchestration project. Its graduation is a win for those building cloud-agnostic, multi-cloud Kubernetes. Its latest release, v1.19, advances multi-component scheduling for distributed AI training jobs.
Lastly, the other big graduation announcement was for Cloud Native Buildpacks . The project, which can transform application code into OCI-compliant container images, joined CNCF as a sandbox project in 2018.
Kubernetes reaches new peaks with v1.37 Garhwal
The latest minor Kubernetes release, v1.37 , is here. It’s nicknamed Garhwal, as an homage to the snow-capped peaks of the Garhwal Himalaya mountain range.
v1.37 includes 67 enhancements: 16 stable, 23 beta, 27 alpha, and one deprecation. Notable features include completing resilient watch cache initialization, which can improve resilience for large clusters and help avoid control plane outages.
One interesting update: KYAML has now reached stable status. It’s billed as a solution to headaches with YAML , including whitespace sensitivity and the dreaded “ Norway Problem .” (I had no idea something as fundamental as YAML had so many issues, but I guess it does.)
KYAML should be able to help. Every KYAML file is still valid YAML, so don’t worry about rewriting anything for backward compatibility. Will KYAML become a more common way to write Kubernetes configuration? Time will tell.
Other notable updates include HorizontalPodAutoscaler scale to zero graduating to beta and being enabled by default. For workloads using object or external metrics, this enables pods to scale down to zero when idle. Other key updates include beta support for manifest-based admission control , and alpha support for pod-level checkpoint and restore.
As Kubernetes evolves, so do the demands on the teams running it. Presenting sponsor HPE helps teams address that complexity with software spanning virtualization, cloud management, observability and automation.
KubeCon travel-scholarship applications close soon: apply now
The schedule for KubeCon + CloudNativeCon North America 2026 is announced . As if the four-day agenda wasn’t jam-packed and mouth-watering enough, this year we’re getting a new AI inference and agentic track.
Thankfully, not everyone has to miss out on the fun. KubeCon offers a scholarship program intended to help fund travel and registration for those in underrepresented groups, or those without the means to do so otherwise.
The deadline to submit a travel funding request is this Sunday . Be sure to submit your request by Sunday, September 13, 11:59 p.m. Mountain Daylight Time (MDT). Registration applications don’t close until Sunday, October 4, 11:59 p.m. MDT.
Access control for Kubernetes finally makes the list
Kolawole Olowoporoku , CNCF Ambassador and senior platform engineer at Armada , is on the CNCF blog this week spotlighting an area that doesn’t always get much attention: identity and access control. He starts with a potent message: “Access control belongs on the same day-zero checklist as networking and storage. On most on-prem clusters, it never makes the list.”
Self-hosted Kubernetes includes authentication and authorization mechanisms, but teams must configure integration with an external identity provider. Without that integration, operators may rely on static client certificates or long-lived tokens.
Such credentials can create security risks when they remain valid longer than intended. Olowoporoku recommends authenticating through an OpenID Connect identity provider using a public client with PKCE. After login, kubectl sends the resulting ID token to the Kubernetes API server, which validates it and applies the configured access permissions.
VMware AI-ifies private cloud visibility
More news on the private cloud front: VMware Cloud Foundation (VCF) 9.1.1 adds new capabilities that help operators gain visibility into their environments.
One addition is enhanced observability into real-time Kubernetes operations, reducing standard five-minute polling intervals to two-second metric streaming. This can help operators detect short-lived pods, memory spikes, and transient performance bottlenecks that might otherwise go unnoticed.
The next major addition is a new AI Assistant for VCF. The conversational interface can help with troubleshooting and diagnostics, check the health of VCF environments, pinpoint root causes, and more. It’s one of many recent moves to add generative AI capabilities to Kubernetes and private cloud operations.
In the latest 2026-09-04 release notes , the Azure Kubernetes Service (AKS) team notes that the latest Kubernetes v1.37 preview is rolling out, with patches for previous versions now available.
Autoscaling for virtual machine node pools has reached general availability. New preview capabilities also give operators more flexibility in managing node pools throughout their lifecycle.
The world surrounding Kubernetes never sleeps. Here are some quick and interesting tidbits in other areas:
CNCF project owners should check out the lat

[truncated]
