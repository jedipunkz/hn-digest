---
source: "https://aws.amazon.com/blogs/aws/introducing-amazon-cloudwatch-omni-collaborative-ai-powered-observability-for-your-applications/"
hn_url: "https://news.ycombinator.com/item?id=49858866"
title: "AWS CloudWatch Omni AI"
article_title: "Introducing Amazon CloudWatch Omni: collaborative AI-powered observability for your applications | AWS News Blog"
image: "https://d2908q01vomqb2.cloudfront.net/da4b9237bacccdf19c0760cab7aec4a8359010b0/2026/09/17/2026-amazon-cloudwatch-omni-for-app.jpg"
author: "based2"
captured_at: "2026-09-26T18:22:56Z"
capture_tool: "hn-digest"
hn_id: 49858866
score: 1
comments: 0
posted_at: "2026-09-26T17:50:05Z"
tags:
  - hacker-news
---

# AWS CloudWatch Omni AI

- HN: [49858866](https://news.ycombinator.com/item?id=49858866)
- Source: [aws.amazon.com](https://aws.amazon.com/blogs/aws/introducing-amazon-cloudwatch-omni-collaborative-ai-powered-observability-for-your-applications/)
- Score: 1
- Comments: 0
- Posted: 2026-09-26T17:50:05Z

## Translation

Title: AWS CloudWatch Omni AI
Article title: Introducing Amazon CloudWatch Omni: collaborative AI-powered observability for your applications | AWS News Blog
Description: Amazon CloudWatch Omni is the next evolution of CloudWatch — unified observability that brings your applications and AI agents into one reimagined experience, with auto-discovered topology, natural language queries, and AI-guided investigation powered by AWS DevOps Agent.

Article text:
Data Migration
On-prem databases weren't built for agentic AI. Amazon RDS is
Independent Software Vendors
AI for software & tech: build, scale, and monetize
AI for Small Businesses
Solutions designed for your business. Delivered by AWS Partners
Artificial Intelligence (AI)
Accelerate AI from experimentation to production with AWS
AI agents
AgentCore: One platform to build, connect and optimize agents
News & announcements
AWS blog
About AWS
AWS is the world's most comprehensive cloud, enabling organizations to accelerate innovation, reduce costs, and scale more efficiently
Get started with one of these featured services or browse all
Transform
Eliminate tech debt with agentic AI to modernize legacy systems and code
Aurora
Serverless relational database service for PostgreSQL, MySQL, and DSQL
Amazon Bedrock
The end-to-end platform for building generative AI applications and agents
Amazon Connect Customer
AI-native solution for delivering exceptional experiences across customer interactions
EC2
Secure and resizable compute capacity for virtually any workload
Nova
Foundation models delivering frontier intelligence and top price performance
OpenSearch Service
Real-time search, observability, and security analytics
S3
Virtually unlimited secure object storage for AI, analytics, and archives
What’s new
Explore new capabilities and the latest technologies
Customer success stories
Learn how customers around the world accelerate their cloud
By industry
Transform your industry with AWS cloud solutions. Access proven architectures, compliance guides, and success stories of customers using AWS products tailored to your sector. Explore AWS solutions for your industry.
Aerospace and Satellite
Cloud solutions to help customers build satellites, conduct space and launch operations, and reimagine space exploration
Automotive
Build smarter vehicles and transform mobility with cloud solutions
Education
Solutions to help facilitate teaching, learning, student engagement, better learning outcomes and modernize IT operations
Energy and Utilities
Revamp legacy operations and accelerate the development of innovative renewable energy business models
Financial Services
Develop innovative and secure solutions across banking, capital markets, insurance, and payments
Games
Enable game development for every genre and platform, from AAA titles to indie studios
Government
Solutions designed to help government agencies modernize, meet mandates, reduce costs, and deliver mission outcomes
Healthcare and Life Sciences
Accelerate innovation and improve patient care with healthcare data management and security
Industrial
Services and solutions for customers across Manufacturing, Automotive, Energy, Power & Utilities, Transportation & Logistics
Manufacturing
Optimize production and speed time-to-market
Media and Entertainment
Transform media & entertainment with the most purpose-built capabilities and partner solutions of any cloud
Nonprofit
Cloud solutions to help nonprofits and NGOs maximize their mission impact and donor relationships
Retail and Consumer Goods
Empower brands to drive market differentiation and growth with cloud solutions
Semiconductor
Design and deliver next-generation chip solutions with cloud technology
Sports
Fuel innovative fan, broadcast, and athlete experiences
Sustainability
Tools and resources to help organizations achieve their sustainability goals with AWS
Telecommunications
Accelerate innovation, scale with confidence, and add agility with cloud-based telecom solutions
Travel and Hospitality
Elevate travel and hospitality with solutions that boost customer experiences and operational efficiency
What's new
Explore new capabil
[truncated]
Introducing Amazon CloudWatch Omni: collaborative AI-powered observability for your applications
Amazon CloudWatch now offers CloudWatch Omni, an AI-powered observability experience for the applications and AI agents you run together. You reach Omni through a dedicated URL for your organization and sign in with the identities you already manage, so working in Omni does not require access to the AWS Management Console. Omni is built on OpenTelemetry: the telemetry you already send to CloudWatch appears in Omni with nothing to reconfigure, and any other workload you instrument with OpenTelemetry sends its telemetry to an OpenTelemetry Protocol (OTLP) endpoint.
CloudWatch Omni offers both agent observability and application observability in a single experience. In our companion post , we introduced the agent observability capabilities of Omni for generative AI and agentic workloads. In this post, we present the application observability experience.
Engineering teams spend a significant portion of their observability time maintaining dashboards, tuning thresholds, and switching between tools to piece together what happened during an incident. When an issue crosses team boundaries, context gets lost in Slack threads and screenshots rather than flowing naturally to the next engineer. CloudWatch Omni changes this by organizing observability around your applications rather than individual signals, and bringing your whole team into the same workspace.
What CloudWatch Omni brings
CloudWatch Omni addresses three problems that engineering teams told us they face today.
One collaborative experience for your whole team. Every engineer accesses CloudWatch Omni through a single URL with enterprise SSO (via IAM Identity Center , supporting Okta, EntraID, and other providers). No AWS Console access is required. SREs, developers, database engineers, and managers share the same data and investigation context. When an investigation escalates, the next person joins the same session with full context already in front of them.
The system adapts as your applications evolve. CloudWatch Omni discovers your services, maps dependencies, and adjusts alarms automatically. Instead of manually curating dashboards and tuning thresholds, you declare what matters (availability targets, latency budgets, error rate thresholds) and Omni adapts as your system changes. When you deploy new services, Omni updates the application topology automatically.
AI-powered investigation with Amazon DevOps Agent. Amazon DevOps Agent participates alongside your team in investigation sessions, correlating signals and suggesting next steps. The agent works from the same telemetry your engineers see, so its suggestions are grounded in the actual state of your application. It identifies correlated events across services, traces root cause paths through your dependency graph, and maintains investigation history for post-incident review.
How an investigation works
When something breaks, CloudWatch Omni opens an investigation session pre-loaded with context. Here is a typical incident workflow:
An alarm fires on elevated error rates in your checkout service. Omni opens a session showing the service topology, correlated signals (a deployment 10 minutes earlier, increased latency from a downstream payment API), and DevOps Agent’s initial analysis.
Your on-call SRE confirms the deployment correlation, pulls in the trace view to identify failing endpoints, and checks if the payment API latency correlates with a capacity limit.
The SRE escalates to the payments team. The payments engineer joins the same session and sees everything found so far, plus DevOps Agent’s correlation with a configuration change in the payment provider’s API gateway. They identify the root cause and roll back.
The entire investigation history is captured automatically. No separate incident report needed.
Walkthrough: setting up your first Space
To set up CloudWatch Omni for your team, open the CloudWatch console and click “Try CloudWatch Omni.”
Next, connect your identity provider through IAM Identity Center (supporting Okta, Azure AD, and other SAML 2.0 providers). Once connected, your team members access Omni directly at your dedicated URL without needing AWS Console credentials.
Create a Space for your team. A Space groups the applications your team owns and the telemetry associated with them.
Once created, Omni discovers your services automatically and maps the dependencies between them. You see your application topology immediately.
You can ask CloudWatch Omni any question about your applications in plain English, and Omni will analyze your telemetry data and surface insights.
You can also set up service health alerts, configure what matters to your team, and trigger an AWS DevOps agent investigation to identify the root cause and develop a mitigation plan.
Application-centric organization
CloudWatch Omni organizes telemetry by application rather than by infrastructure component. The system automatically discovers services from the telemetry data and AWS Config resource discovery, maps dependencies, and lets you see your application as a connected system rather than a collection of isolated resources.
Each team gets a Space that contains the applications they own. A Space points at existing CloudWatch data (logs, metrics, traces, and alarms) with no additional data movement required. Dynamic views replace the maintenance burden of static dashboards, providing ongoing visibility into SLOs and application health.
Getting started
Getting started takes minutes and doesn’t require reconfiguration of your existing CloudWatch setup.
If you’re an existing CloudWatch customer: Click “Try CloudWatch Omni” in the CloudWatch console. All your existing telemetry (logs, metrics, traces, and alarms) is immediately available. Workloads are discovered automatically, and you can start an investigation or browse your application topology right away.
For organization-wide deployment: An administrator configures a domain, connects your identity provider via IAM Identity Center, defines Spaces for teams and environments, and invites users. Each Space points at existing CloudWatch data with no additional data movement required.
For applications in other environments: CloudWatch Omni provides connectors that make it easy to bring in telemetry from additional environments. All ingested telemetry appears alongside your AWS data in the same Spaces and investigation sessions.
For generative AI and agentic workloads: The same CloudWatch Omni experience delivers purpose-built observability for AI agents, including trace exploration, evaluation frameworks, and real-time monitoring. In our companion post, we introduced the agent observability capabilities of Omni; for that walkthrough, see Introducing Amazon CloudWatch Omni: AI-powered observability for generative AI and agentic workloads .
CloudWatch Omni extends CloudWatch. Existing alarms, dashboards, APIs, and console workflows continue unchanged.
Access is through a dedicated web application with enterprise SSO. Engineers don’t need AWS Console access to use it.
Once you setup, DevOps Agent is enabled by default in every Omni investigation session.
Pricing and availability
Amazon CloudWatch Omni is now available. Existing CloudWatch customers can try it directly from the CloudWatch console. For pricing details, visit the Amazon CloudWatch pricing page .
To get started, visit Amazon CloudWatch Omni or click “Try CloudWatch Omni” in the Amazon CloudWatch console .
If you want to call APIs, search documentation, find regional availability, and check troubleshooting about this feature, try using the AWS MCP Server and plugins with your preferred AI tool. Share your feedback on AWS re:Post or reach out through your usual AWS Support contacts.
9/24/2026 – Editor’s note: Azure AD has been renamed to Microsoft Entra ID.
Daniel Abib is a senior specialist solutions architect at AWS, focused on generative AI and Amazon Bedrock — and passionate about serverless. He helps startups and enterprises build AI-powered applications and modernize with cloud-native architectures. A three-time Ironman finisher and four-time re:Invent speaker, he brings the same endurance mindset to building cloud solutions.
English
Back to top
Am

[truncated]

## Original Extract

Amazon CloudWatch Omni is the next evolution of CloudWatch — unified observability that brings your applications and AI agents into one reimagined experience, with auto-discovered topology, natural language queries, and AI-guided investigation powered by AWS DevOps Agent.

Data Migration
On-prem databases weren't built for agentic AI. Amazon RDS is
Independent Software Vendors
AI for software & tech: build, scale, and monetize
AI for Small Businesses
Solutions designed for your business. Delivered by AWS Partners
Artificial Intelligence (AI)
Accelerate AI from experimentation to production with AWS
AI agents
AgentCore: One platform to build, connect and optimize agents
News & announcements
AWS blog
About AWS
AWS is the world's most comprehensive cloud, enabling organizations to accelerate innovation, reduce costs, and scale more efficiently
Get started with one of these featured services or browse all
Transform
Eliminate tech debt with agentic AI to modernize legacy systems and code
Aurora
Serverless relational database service for PostgreSQL, MySQL, and DSQL
Amazon Bedrock
The end-to-end platform for building generative AI applications and agents
Amazon Connect Customer
AI-native solution for delivering exceptional experiences across customer interactions
EC2
Secure and resizable compute capacity for virtually any workload
Nova
Foundation models delivering frontier intelligence and top price performance
OpenSearch Service
Real-time search, observability, and security analytics
S3
Virtually unlimited secure object storage for AI, analytics, and archives
What’s new
Explore new capabilities and the latest technologies
Customer success stories
Learn how customers around the world accelerate their cloud
By industry
Transform your industry with AWS cloud solutions. Access proven architectures, compliance guides, and success stories of customers using AWS products tailored to your sector. Explore AWS solutions for your industry.
Aerospace and Satellite
Cloud solutions to help customers build satellites, conduct space and launch operations, and reimagine space exploration
Automotive
Build smarter vehicles and transform mobility with cloud solutions
Education
Solutions to help facilitate teaching, learning, student engagement, better learning outcomes and modernize IT operations
Energy and Utilities
Revamp legacy operations and accelerate the development of innovative renewable energy business models
Financial Services
Develop innovative and secure solutions across banking, capital markets, insurance, and payments
Games
Enable game development for every genre and platform, from AAA titles to indie studios
Government
Solutions designed to help government agencies modernize, meet mandates, reduce costs, and deliver mission outcomes
Healthcare and Life Sciences
Accelerate innovation and improve patient care with healthcare data management and security
Industrial
Services and solutions for customers across Manufacturing, Automotive, Energy, Power & Utilities, Transportation & Logistics
Manufacturing
Optimize production and speed time-to-market
Media and Entertainment
Transform media & entertainment with the most purpose-built capabilities and partner solutions of any cloud
Nonprofit
Cloud solutions to help nonprofits and NGOs maximize their mission impact and donor relationships
Retail and Consumer Goods
Empower brands to drive market differentiation and growth with cloud solutions
Semiconductor
Design and deliver next-generation chip solutions with cloud technology
Sports
Fuel innovative fan, broadcast, and athlete experiences
Sustainability
Tools and resources to help organizations achieve their sustainability goals with AWS
Telecommunications
Accelerate innovation, scale with confidence, and add agility with cloud-based telecom solutions
Travel and Hospitality
Elevate travel and hospitality with solutions that boost customer experiences and operational efficiency
What's new
Explore new capabil
[truncated]
Introducing Amazon CloudWatch Omni: collaborative AI-powered observability for your applications
Amazon CloudWatch now offers CloudWatch Omni, an AI-powered observability experience for the applications and AI agents you run together. You reach Omni through a dedicated URL for your organization and sign in with the identities you already manage, so working in Omni does not require access to the AWS Management Console. Omni is built on OpenTelemetry: the telemetry you already send to CloudWatch appears in Omni with nothing to reconfigure, and any other workload you instrument with OpenTelemetry sends its telemetry to an OpenTelemetry Protocol (OTLP) endpoint.
CloudWatch Omni offers both agent observability and application observability in a single experience. In our companion post , we introduced the agent observability capabilities of Omni for generative AI and agentic workloads. In this post, we present the application observability experience.
Engineering teams spend a significant portion of their observability time maintaining dashboards, tuning thresholds, and switching between tools to piece together what happened during an incident. When an issue crosses team boundaries, context gets lost in Slack threads and screenshots rather than flowing naturally to the next engineer. CloudWatch Omni changes this by organizing observability around your applications rather than individual signals, and bringing your whole team into the same workspace.
What CloudWatch Omni brings
CloudWatch Omni addresses three problems that engineering teams told us they face today.
One collaborative experience for your whole team. Every engineer accesses CloudWatch Omni through a single URL with enterprise SSO (via IAM Identity Center , supporting Okta, EntraID, and other providers). No AWS Console access is required. SREs, developers, database engineers, and managers share the same data and investigation context. When an investigation escalates, the next person joins the same session with full context already in front of them.
The system adapts as your applications evolve. CloudWatch Omni discovers your services, maps dependencies, and adjusts alarms automatically. Instead of manually curating dashboards and tuning thresholds, you declare what matters (availability targets, latency budgets, error rate thresholds) and Omni adapts as your system changes. When you deploy new services, Omni updates the application topology automatically.
AI-powered investigation with Amazon DevOps Agent. Amazon DevOps Agent participates alongside your team in investigation sessions, correlating signals and suggesting next steps. The agent works from the same telemetry your engineers see, so its suggestions are grounded in the actual state of your application. It identifies correlated events across services, traces root cause paths through your dependency graph, and maintains investigation history for post-incident review.
How an investigation works
When something breaks, CloudWatch Omni opens an investigation session pre-loaded with context. Here is a typical incident workflow:
An alarm fires on elevated error rates in your checkout service. Omni opens a session showing the service topology, correlated signals (a deployment 10 minutes earlier, increased latency from a downstream payment API), and DevOps Agent’s initial analysis.
Your on-call SRE confirms the deployment correlation, pulls in the trace view to identify failing endpoints, and checks if the payment API latency correlates with a capacity limit.
The SRE escalates to the payments team. The payments engineer joins the same session and sees everything found so far, plus DevOps Agent’s correlation with a configuration change in the payment provider’s API gateway. They identify the root cause and roll back.
The entire investigation history is captured automatically. No separate incident report needed.
Walkthrough: setting up your first Space
To set up CloudWatch Omni for your team, open the CloudWatch console and click “Try CloudWatch Omni.”
Next, connect your identity provider through IAM Identity Center (supporting Okta, Azure AD, and other SAML 2.0 providers). Once connected, your team members access Omni directly at your dedicated URL without needing AWS Console credentials.
Create a Space for your team. A Space groups the applications your team owns and the telemetry associated with them.
Once created, Omni discovers your services automatically and maps the dependencies between them. You see your application topology immediately.
You can ask CloudWatch Omni any question about your applications in plain English, and Omni will analyze your telemetry data and surface insights.
You can also set up service health alerts, configure what matters to your team, and trigger an AWS DevOps agent investigation to identify the root cause and develop a mitigation plan.
Application-centric organization
CloudWatch Omni organizes telemetry by application rather than by infrastructure component. The system automatically discovers services from the telemetry data and AWS Config resource discovery, maps dependencies, and lets you see your application as a connected system rather than a collection of isolated resources.
Each team gets a Space that contains the applications they own. A Space points at existing CloudWatch data (logs, metrics, traces, and alarms) with no additional data movement required. Dynamic views replace the maintenance burden of static dashboards, providing ongoing visibility into SLOs and application health.
Getting started
Getting started takes minutes and doesn’t require reconfiguration of your existing CloudWatch setup.
If you’re an existing CloudWatch customer: Click “Try CloudWatch Omni” in the CloudWatch console. All your existing telemetry (logs, metrics, traces, and alarms) is immediately available. Workloads are discovered automatically, and you can start an investigation or browse your application topology right away.
For organization-wide deployment: An administrator configures a domain, connects your identity provider via IAM Identity Center, defines Spaces for teams and environments, and invites users. Each Space points at existing CloudWatch data with no additional data movement required.
For applications in other environments: CloudWatch Omni provides connectors that make it easy to bring in telemetry from additional environments. All ingested telemetry appears alongside your AWS data in the same Spaces and investigation sessions.
For generative AI and agentic workloads: The same CloudWatch Omni experience delivers purpose-built observability for AI agents, including trace exploration, evaluation frameworks, and real-time monitoring. In our companion post, we introduced the agent observability capabilities of Omni; for that walkthrough, see Introducing Amazon CloudWatch Omni: AI-powered observability for generative AI and agentic workloads .
CloudWatch Omni extends CloudWatch. Existing alarms, dashboards, APIs, and console workflows continue unchanged.
Access is through a dedicated web application with enterprise SSO. Engineers don’t need AWS Console access to use it.
Once you setup, DevOps Agent is enabled by default in every Omni investigation session.
Pricing and availability
Amazon CloudWatch Omni is now available. Existing CloudWatch customers can try it directly from the CloudWatch console. For pricing details, visit the Amazon CloudWatch pricing page .
To get started, visit Amazon CloudWatch Omni or click “Try CloudWatch Omni” in the Amazon CloudWatch console .
If you want to call APIs, search documentation, find regional availability, and check troubleshooting about this feature, try using the AWS MCP Server and plugins with your preferred AI tool. Share your feedback on AWS re:Post or reach out through your usual AWS Support contacts.
9/24/2026 – Editor’s note: Azure AD has been renamed to Microsoft Entra ID.
Daniel Abib is a senior specialist solutions architect at AWS, focused on generative AI and Amazon Bedrock — and passionate about serverless. He helps startups and enterprises build AI-powered applications and modernize with cloud-native architectures. A three-time Ironman finisher and four-time re:Invent speaker, he brings the same endurance mindset to building cloud solutions.
English
Back to top
Am

[truncated]
