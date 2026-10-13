---
source: "https://encore.dev"
hn_url: "https://news.ycombinator.com/item?id=50068998"
title: "Encore – a simpler way to build on AWS and GCP"
article_title: "Automated Infrastructure for Humans and Agents — Encore"
image: "https://encore.dev/assets/branding/social_card_2026.png"
author: "dohguy"
captured_at: "2026-10-13T10:10:10Z"
capture_tool: "hn-digest"
hn_id: 50068998
score: 3
comments: 0
posted_at: "2026-10-13T09:32:40Z"
tags:
  - hacker-news
---

# Encore – a simpler way to build on AWS and GCP

- HN: [50068998](https://news.ycombinator.com/item?id=50068998)
- Source: [encore.dev](https://encore.dev)
- Score: 3
- Comments: 0
- Posted: 2026-10-13T09:32:40Z

## Translation

Title: Encore – a simpler way to build on AWS and GCP
Article title: Automated Infrastructure for Humans and Agents — Encore
Description: Encore automates infrastructure from development to production, so developers and agents ship without waiting on Terraform PRs or platform tickets.

Article text:
Automated Infrastructure for Humans and Agents — Encore Encore Product Solutions Developers Customers Docs Pricing ★12K+ Sign in Get started ★12k+ Automated infrastructure for humans and agents
Encore automates infrastructure from development to production, so developers and agents can build, validate, and ship without waiting on Terraform PRs or platform tickets.
Curious how Encore could work for your team?
Copy Run in your terminal to get started locally.
Watch video 0:33 “What used to take days or weeks of back-and-forth between developers and infra teams is now automated and completed in minutes.”
Skip the Terraform handoffs.
Go from code to production in minutes.
Encore replaces Terraform files, infrastructure PRs, and DevOps handoffs with an automated workflow that provisions and validates infrastructure from development through production.
Build with infrastructure
in the feedback loop
Encore lets developers and agents define the infrastructure a feature needs directly in application code. Databases, queues, buckets and more run locally with the app, so you can build, run and validate the complete feature without needing to deploy.
Run background agents
in cloud sandboxes Coming soon
Each agent gets an isolated environment with your app and infrastructure, running independently of your laptop. Spin up as many as you need and let them work in parallel.
agent/refactor-checkout Reading service docs Live Add refunds API Codex agent/refunds-api Reading the orders schema Live Migrate orders schema Claude Code agent/migrate-schema Reading existing migrations Live Fix webhook retries Codex agent/webhook-retries Branch merged · ran 6m 30s · sandbox destroyed Inactive Add search index Claude Code agent/search-index Branch closed · ran 3m 44s · sandbox destroyed Inactive VALIDATE VALIDATE Test every feature end-to-end
in its own preview environment
Push a branch and Encore sets up a dedicated temporary environment. Integrate with external systems and validate complete workflows without relying on a shared staging environment where someone else’s changes can break your tests.
Infrastructure is provisioned
in your cloud on deploy
Encore automatically provisions the databases, queues, buckets, and other infrastructure your application needs in your AWS or GCP account. You keep your cloud account, VPC, and compliance boundary.
Building and testing your app 0s Search log Overview Raw Build & test — Compile 5 steps waiting Test 5 steps waiting Provision infrastructure 4 changes ( 4 to create ) — Storage 1 Bucket ("receipts") Amazon S3 Security 3 IAM role ("orders") Task role for orders IAM policy ("receipts-writer") s3:PutObject on receipts Policy attachment ("receipts-writer") Attached to role orders Deploy release 1 change ( 1 to modify ) — Compute 1 Service ("orders") Amazon ECS on Fargate Your app is live Deployment to production in your AWS account is complete. https://api.acme.com Go to Dashboard Open App OPERATE OPERATE Let developers self-serve without losing control
Environment-specific configuration is managed centrally, separate from the development loop. Developers and agents move independently within approved policies, without requiring a platform review for every change.
Settings Infra config Secrets Cloud account 1 AWS RDS Instance (1 database) orders-prod-aws-rds-4f80bff us-east-1 CONFIGURATION EDIT Instance class db.t3.medium Max connections 500 Allocated storage 20 GiB Postgres version 17.6 DATABASES (1) orders No pending changes Audit trail Adopt Encore without a platform rewrite
Start with one service, team, or new project.
Encore works alongside your existing stack and deploys to your VPC, so you can prove it in production before gradually expanding service by service.
Loved by teams who refuse to slow down
" Encore is our foundation for all new development. Since adopting it, we've seen a 2-3x increase in development speed and 90% shorter project lead times. "
" Encore lets us go from idea to shipped product at least 3x faster and saves us hours of devops work every week. "
" We've reduced time spent on DevOps by 95%, and we're now on track to save $60K annually in cloud costs. "
" Encore is an unfair advantage. At Pave we've moved 2x faster than we did at Monzo. "
" We're saving $200k+ annually compared to additional hiring. The ROI is outstanding, easily 10x. "
" You define infrastructure in your code and it's automatically spun up for you. You can go from zero to a deployed service on GCP in an hour. I don't know how else to do it. "
" Encore's inferred infrastructure has allowed us to move at unprecedented speed. Combined with database branching in Neon, we're able to spin ephemeral environments up and down with incredible ease. Stakeholders love preview app URLs. "
View all stories See how teams ship faster → Join the growing community building with Encore
Discord Hang out with other Encore developers
GitHub Star and contribute to the project
YouTube Watch tutorials and talks
No. You write application code and Encore sets up the cloud resources it needs, including networking, IAM and secrets. Everything is created in your own account, so you can inspect it in your cloud console at any time.
Instead of writing separate config files, you declare infrastructure directly in your application code. Encore provisions everything automatically in each environment, so there are no Terraform configs or Pulumi stacks to maintain.
No. Encore works alongside what you have. Teams usually start with new services and leave existing infrastructure, and the Terraform that manages it, where it is. Encore can also deploy into an existing Kubernetes cluster.
Encore supports TypeScript and Go, both with the full feature set. Python and Rust support is coming soon. You deploy to AWS or GCP in your own account. On the frontend, use any framework and generate a typed API client for your backend.
Yes. Agents define infrastructure in the same code they already write, and run and validate the full app locally or in a preview environment. Encore's MCP server gives them context about your services, APIs and traces.
Yes. Adopt Encore incrementally by converting individual services to the Encore SDK, one by one. You can also connect existing databases, buckets and topics instead of recreating them.
Use it directly. Your application runs in your own cloud account, so you can call any AWS or GCP service from your code alongside the resources Encore manages. Nothing is walled off.
Yes. Encore provisions standard managed services like RDS, Cloud SQL, ECS on Fargate, Cloud Run and GKE in your own cloud account. Your application runs there, with no runtime dependency on Encore's platform, so production traffic never passes through Encore and there's nothing to migrate as you grow. Teams like Groupon and Coinbase run production workloads on Encore.
Agents never get cloud credentials. They declare infrastructure in code, and Encore provisions it deterministically, following best practices for security and reliability. Environment-specific configuration, like instance sizes and secrets, is applied in Encore's control plane, separate from the code agents write. Every pull request gets its own isolated preview environment, and changes are recorded in the audit trail.
An IAM role (AWS) or service account (GCP) that you grant in your own account and can revoke at any time. Encore uses it to create and update the resources your services declare, and never deletes infrastructure without your approval. On the Enterprise plan, we support just-in-time permissions that are granted only when you approve a specific change.
The Starter plan is free for local development and small projects. The Pro plan is $49 per team member per month, plus per-environment and per-resource fees. Your cloud resources are billed by AWS or GCP directly in your own account.
Yes. Your infrastructure runs in your own cloud account, so your cloud bill, and any credits on it, work as usual. Credits cover what you pay AWS or GCP, not the Encore plan fee.
Encore is designed to avoid vendor lock-in. You deploy to your own cloud account (AWS or GCP), so you own and control your infrastructure.
If you decide to stop using the Encore platform for infrastructure automation, you can keep using your existing infrastructure and set up your own CI/CD pipeline using Encore's Open Source tools.
From idea to deployed in minutes
Install Encore, then use your AI agent to build, run, and deploy an app with real infrastructure in minutes.
Copy Run in your terminal to get started locally.
Product updates and engineering deep-dives.
Discord GitHub YouTube © 2026 Encore

## Original Extract

Encore automates infrastructure from development to production, so developers and agents ship without waiting on Terraform PRs or platform tickets.

Automated Infrastructure for Humans and Agents — Encore Encore Product Solutions Developers Customers Docs Pricing ★12K+ Sign in Get started ★12k+ Automated infrastructure for humans and agents
Encore automates infrastructure from development to production, so developers and agents can build, validate, and ship without waiting on Terraform PRs or platform tickets.
Curious how Encore could work for your team?
Copy Run in your terminal to get started locally.
Watch video 0:33 “What used to take days or weeks of back-and-forth between developers and infra teams is now automated and completed in minutes.”
Skip the Terraform handoffs.
Go from code to production in minutes.
Encore replaces Terraform files, infrastructure PRs, and DevOps handoffs with an automated workflow that provisions and validates infrastructure from development through production.
Build with infrastructure
in the feedback loop
Encore lets developers and agents define the infrastructure a feature needs directly in application code. Databases, queues, buckets and more run locally with the app, so you can build, run and validate the complete feature without needing to deploy.
Run background agents
in cloud sandboxes Coming soon
Each agent gets an isolated environment with your app and infrastructure, running independently of your laptop. Spin up as many as you need and let them work in parallel.
agent/refactor-checkout Reading service docs Live Add refunds API Codex agent/refunds-api Reading the orders schema Live Migrate orders schema Claude Code agent/migrate-schema Reading existing migrations Live Fix webhook retries Codex agent/webhook-retries Branch merged · ran 6m 30s · sandbox destroyed Inactive Add search index Claude Code agent/search-index Branch closed · ran 3m 44s · sandbox destroyed Inactive VALIDATE VALIDATE Test every feature end-to-end
in its own preview environment
Push a branch and Encore sets up a dedicated temporary environment. Integrate with external systems and validate complete workflows without relying on a shared staging environment where someone else’s changes can break your tests.
Infrastructure is provisioned
in your cloud on deploy
Encore automatically provisions the databases, queues, buckets, and other infrastructure your application needs in your AWS or GCP account. You keep your cloud account, VPC, and compliance boundary.
Building and testing your app 0s Search log Overview Raw Build & test — Compile 5 steps waiting Test 5 steps waiting Provision infrastructure 4 changes ( 4 to create ) — Storage 1 Bucket ("receipts") Amazon S3 Security 3 IAM role ("orders") Task role for orders IAM policy ("receipts-writer") s3:PutObject on receipts Policy attachment ("receipts-writer") Attached to role orders Deploy release 1 change ( 1 to modify ) — Compute 1 Service ("orders") Amazon ECS on Fargate Your app is live Deployment to production in your AWS account is complete. https://api.acme.com Go to Dashboard Open App OPERATE OPERATE Let developers self-serve without losing control
Environment-specific configuration is managed centrally, separate from the development loop. Developers and agents move independently within approved policies, without requiring a platform review for every change.
Settings Infra config Secrets Cloud account 1 AWS RDS Instance (1 database) orders-prod-aws-rds-4f80bff us-east-1 CONFIGURATION EDIT Instance class db.t3.medium Max connections 500 Allocated storage 20 GiB Postgres version 17.6 DATABASES (1) orders No pending changes Audit trail Adopt Encore without a platform rewrite
Start with one service, team, or new project.
Encore works alongside your existing stack and deploys to your VPC, so you can prove it in production before gradually expanding service by service.
Loved by teams who refuse to slow down
" Encore is our foundation for all new development. Since adopting it, we've seen a 2-3x increase in development speed and 90% shorter project lead times. "
" Encore lets us go from idea to shipped product at least 3x faster and saves us hours of devops work every week. "
" We've reduced time spent on DevOps by 95%, and we're now on track to save $60K annually in cloud costs. "
" Encore is an unfair advantage. At Pave we've moved 2x faster than we did at Monzo. "
" We're saving $200k+ annually compared to additional hiring. The ROI is outstanding, easily 10x. "
" You define infrastructure in your code and it's automatically spun up for you. You can go from zero to a deployed service on GCP in an hour. I don't know how else to do it. "
" Encore's inferred infrastructure has allowed us to move at unprecedented speed. Combined with database branching in Neon, we're able to spin ephemeral environments up and down with incredible ease. Stakeholders love preview app URLs. "
View all stories See how teams ship faster → Join the growing community building with Encore
Discord Hang out with other Encore developers
GitHub Star and contribute to the project
YouTube Watch tutorials and talks
No. You write application code and Encore sets up the cloud resources it needs, including networking, IAM and secrets. Everything is created in your own account, so you can inspect it in your cloud console at any time.
Instead of writing separate config files, you declare infrastructure directly in your application code. Encore provisions everything automatically in each environment, so there are no Terraform configs or Pulumi stacks to maintain.
No. Encore works alongside what you have. Teams usually start with new services and leave existing infrastructure, and the Terraform that manages it, where it is. Encore can also deploy into an existing Kubernetes cluster.
Encore supports TypeScript and Go, both with the full feature set. Python and Rust support is coming soon. You deploy to AWS or GCP in your own account. On the frontend, use any framework and generate a typed API client for your backend.
Yes. Agents define infrastructure in the same code they already write, and run and validate the full app locally or in a preview environment. Encore's MCP server gives them context about your services, APIs and traces.
Yes. Adopt Encore incrementally by converting individual services to the Encore SDK, one by one. You can also connect existing databases, buckets and topics instead of recreating them.
Use it directly. Your application runs in your own cloud account, so you can call any AWS or GCP service from your code alongside the resources Encore manages. Nothing is walled off.
Yes. Encore provisions standard managed services like RDS, Cloud SQL, ECS on Fargate, Cloud Run and GKE in your own cloud account. Your application runs there, with no runtime dependency on Encore's platform, so production traffic never passes through Encore and there's nothing to migrate as you grow. Teams like Groupon and Coinbase run production workloads on Encore.
Agents never get cloud credentials. They declare infrastructure in code, and Encore provisions it deterministically, following best practices for security and reliability. Environment-specific configuration, like instance sizes and secrets, is applied in Encore's control plane, separate from the code agents write. Every pull request gets its own isolated preview environment, and changes are recorded in the audit trail.
An IAM role (AWS) or service account (GCP) that you grant in your own account and can revoke at any time. Encore uses it to create and update the resources your services declare, and never deletes infrastructure without your approval. On the Enterprise plan, we support just-in-time permissions that are granted only when you approve a specific change.
The Starter plan is free for local development and small projects. The Pro plan is $49 per team member per month, plus per-environment and per-resource fees. Your cloud resources are billed by AWS or GCP directly in your own account.
Yes. Your infrastructure runs in your own cloud account, so your cloud bill, and any credits on it, work as usual. Credits cover what you pay AWS or GCP, not the Encore plan fee.
Encore is designed to avoid vendor lock-in. You deploy to your own cloud account (AWS or GCP), so you own and control your infrastructure.
If you decide to stop using the Encore platform for infrastructure automation, you can keep using your existing infrastructure and set up your own CI/CD pipeline using Encore's Open Source tools.
From idea to deployed in minutes
Install Encore, then use your AI agent to build, run, and deploy an app with real infrastructure in minutes.
Copy Run in your terminal to get started locally.
Product updates and engineering deep-dives.
Discord GitHub YouTube © 2026 Encore
