---
source: "https://www.orchesterra.in/"
hn_url: "https://news.ycombinator.com/item?id=50506664"
title: "Show HN: OrchesTerra – Generate Terraform from a repo, then catch drift"
article_title: "OrchesTerra — AI-Native Infrastructure Governance Control Plane"
image: "https://www.orchesterra.in/opengraph-image?8f57f4cdfb9b9453"
author: "akshaykilam"
captured_at: "2026-11-15T13:31:08Z"
capture_tool: "hn-digest"
hn_id: 50506664
score: 2
comments: 0
posted_at: "2026-11-15T13:06:47Z"
tags:
  - hacker-news
---

# Show HN: OrchesTerra – Generate Terraform from a repo, then catch drift

- HN: [50506664](https://news.ycombinator.com/item?id=50506664)
- Source: [www.orchesterra.in](https://www.orchesterra.in/)
- Score: 2
- Comments: 0
- Posted: 2026-11-15T13:06:47Z

## Translation

Title: Show HN: OrchesTerra – Generate Terraform from a repo, then catch drift
Article title: OrchesTerra — AI-Native Infrastructure Governance Control Plane
Description: Paste a repo. OrchesTerra infers topology, generates Terraform, detects drift, and orchestrates governed reconciliation across AWS, Azure, GCP, and OCI.

Article text:
OrchesTerra — AI-Native Infrastructure Governance Control Plane Product Use Cases Features Pricing Managed DevSecOps New Blog Resources Sign in Start trial AI Operating System for Cloud Infrastructure
Understand cloud architecture from any repository.
Paste a repository and OrchesTerra automatically discovers services, maps dependencies, generates architecture, estimates cloud costs, validates governance, and creates deployment-ready infrastructure.
Reduce architecture discovery from weeks to minutes.
The live demo opens a read-only sample workspace (Acme Retail). No signup.
Rather not run it yourself? A dedicated DevSecOps team can manage your infrastructure on OrchesTerra. Analysis output
· Multi-cloud topology generated
· Estimated deployment cost: $4.2k/month
· Governance validation complete
Starting cognition scan… Two ways to work with us
Run it yourself, or let our team run it.
Same platform, same approvals, same audit trail. You choose who operates it.
Your platform team uses OrchesTerra to generate, govern and reconcile infrastructure. From $99 a month, with a 14-day trial.
Repository to Terraform in minutes
Approval chains and policy packs
A dedicated team of our engineers runs your cloud on OrchesTerra: pipelines, patching, drift fixes, DR drills and cost reviews. Every change still waits for your approval.
Named engineers who know your estate
Landing zones on any of 4 clouds
Monthly cost and security report
Incident response on your hours
One control plane from repository to production.
Scans a repository and finds its services, dependencies, Docker Compose, Helm charts, Kubernetes manifests, Terragrunt and CI pipelines, then draws the topology.
Generation Terraform & OpenTofu
Writes reviewable IaC from that topology, provisions through CloudFormation change sets when you prefer, and imports the resources you already run.
Prices every plan against live AWS, Azure, GCP and OCI price lists, and flags savings in what is already running.
Scheduled and predictive scans compare state with the cloud; restore, accept or forget each drift through an approval.
One inbox for every change, with reviewer groups, delegation and maintenance windows. Nothing applies until the chain is complete.
OPA policy packs, a simulator to test them, Checkov config scanning and compliance packs, all enforced in the same pipeline.
Plan replication and standby regions, run failover, and check readiness before the day you need it.
An AI agent that reads your estate and your bill, answers in plain words, and asks before it changes anything.
MULTI-CLOUD STATE RECONCILIATION
Continuously reconcile declared and actual cloud state .
Detect infrastructure drift before it becomes an incident.
Compare Terraform and OpenTofu state against AWS, Azure, GCP, and OCI in real time.
Identify configuration drift and unmanaged resources before changes create operational risk.
Detect drift early, reduce noise, and maintain reliable cloud environments.
Terraform · OpenTofu vs live cloud state
Critical drift events detected
Workloads potentially impacted
Resources outside declared state
Continuously reconciling declared and actual cloud state
Differences between Terraform/OpenTofu state and live cloud resources.
Ingress 0.0.0.0/0:443 added outside Terraform
aws_lb_listener_rule.public_alb
Catch-all listener rule rerouted to public target group
azurerm_network_security_rule.aks_api
AKS API endpoint opened to broad source CIDR block
oci_load_balancer_listener.public
Public listener certificate policy downgraded outside IaC
4 drift events identified during state reconciliation.
From intent to infrastructure.
OrchesTerra understands your intent, asks the right questions, and builds the optimal infrastructure—ready for you to review and launch.
Describe what you need in natural language.
Orbit designs the best-fit architecture in real time.
Review the pre-filled configuration.
Generated from your conversation
Best practices built-in at every step.
Adapts to your inputs in real time.
Security, reliability and compliance first.
Short clips from a live sample workspace: from repository to governed infrastructure, with cost, drift and approvals along the way.
Reopen a scanned repository and land on the architecture OrchesTerra generated from it.
Watch the 2-minute tour Try the live demo Click around a read-only sample workspace for Acme Retail. No signup needed.
Ask questions in plain words. Rehearse the bad day.
Orbit answers questions about your estate and your bill, and drafts changes for approval. Disaster recovery keeps standby regions ready and drills on record.
Why did our AWS bill go up last month?
Cross-region replication ● Healthy
Infrastructure Approval Workflows
Detect infrastructure drift, route approvals, and enforce controlled execution before changes reach production.
Change Type Security group remediation Resource aws_security_group.app_sg Drift Summary Inbound rule 0.0.0.0/0:22 detected in actual state Triggered By Drift detection Approval Chain
All required approvals must complete before execution can proceed.
Your infrastructure, run by a dedicated team.
Our engineers operate your cloud on OrchesTerra, so your developers ship product instead of patching clusters. You keep the approvals and the audit log.
Landing zones on AWS, Azure, GCP and OCI
Terraform modules and state hygiene
CI/CD and GitOps with security gates
Vulnerability and patch management
Drift and policy findings fixed
Compliance evidence from the audit trail
Monthly cost reviews and rightsizing
Incident response on your hours
We scan your repositories and accounts, map drift, risk and spend, and agree the scope.
Existing resources are imported into Terraform state; pipelines and approval chains are set up.
Your named engineers run day-to-day changes, patching and incident response.
A monthly review of cost, security posture, drift and DR readiness.
You stay in control. Every change our team makes waits for your approvers and lands in your audit log.
Replace fragmented infrastructure workflows with a single orchestration layer for generation, approvals, deployment, and reconciliation.
Natural language infrastructure generation
Multi-cloud blueprint generation
Deployment-ready IAC generation
AI-assisted infrastructure reasoning
Managed DevSecOps team on request
Orches Terra is the only platform that unifies infrastructure creation, governance, and operations into a single AI-native orchestration layer .
See seats, cloud accounts, and governance limits on the pricing page .
Built for teams operating production infrastructure with governance, access control, auditability, and secure credential boundaries.
Enterprise patterns baked in, not bolted on.
Least-privilege access, credential isolation, and policy guardrails.
Complete traceability across every action and environment.
Reliable, consistent operations across teams, clouds, and regions.
Route infrastructure changes through configurable approval chains before deployment or reconciliation.
Control access to environments, projects, and infrastructure operations with precision.
Track deployments, approvals, reconciliations, and operator actions with immutable logs.
Integrate cloud credentials securely with scoped permissions and controlled access.
Every deployment, approval, and reconciliation action is fully traceable, auditable, and governed through a single control plane.
Priced per team, not per resource.
Unlimited users · Unlimited cloud accounts
AWS, Azure, Google Cloud and Oracle Cloud (OCI). Cost estimates, drift scans and policy checks run across all four from one workspace.
Generated code is Terraform or OpenTofu, your choice. On AWS you can also provision through CloudFormation change sets and import existing stacks.
No. OrchesTerra imports existing AWS and Azure resources into Terraform state, so drift detection and approvals cover what you run today.
A scoped role per cloud account. On AWS you create it with a one-click CloudFormation template that pins a per-tenant External ID, with separate read-only and deploy roles. Sessions are labelled for CloudTrail and can be locked to source IPs, credentials are encrypted and scoped to your workspace, plans and applies run on a separate sandboxed runner, and every apply waits for the approval chain you configure.
A dedicated team of our engineers runs your infrastructure on OrchesTerra: pipelines, patching, drift and policy fixes, DR drills and cost reviews. Every change they make goes through your approvals and lands in your audit log.
Pick Starter or Business and keep your workspace, or talk to us about Enterprise or Managed DevSecOps. Nothing is applied to your cloud without approval, during the trial or after it.
Launch your infrastructure intent
Describe your infrastructure intent, operational constraints, and deployment priorities — Orches Terra will generate the execution path.
What OrchesTerra will generate
Multi-cloud infrastructure blueprints
Private · Guided · Architecture-specific Infrastructure Intent Company Work email Primary cloud environment Select environment AWS Azure GCP OCI Multi-cloud Team size Select team size 1–10 10–50 50–200 200+ Describe your deployment, drift, or governance challenges 0 /1000 Start Orchestration Deployment-ready outputs. Governed before execution.
AI-native infrastructure orchestration for modern platform teams, and a DevSecOps team that can run it for you.
The AI orchestration engine powering infrastructure intent, governance, and execution.
© 2026 OrchesTerra · The Red Ruby

## Original Extract

Paste a repo. OrchesTerra infers topology, generates Terraform, detects drift, and orchestrates governed reconciliation across AWS, Azure, GCP, and OCI.

OrchesTerra — AI-Native Infrastructure Governance Control Plane Product Use Cases Features Pricing Managed DevSecOps New Blog Resources Sign in Start trial AI Operating System for Cloud Infrastructure
Understand cloud architecture from any repository.
Paste a repository and OrchesTerra automatically discovers services, maps dependencies, generates architecture, estimates cloud costs, validates governance, and creates deployment-ready infrastructure.
Reduce architecture discovery from weeks to minutes.
The live demo opens a read-only sample workspace (Acme Retail). No signup.
Rather not run it yourself? A dedicated DevSecOps team can manage your infrastructure on OrchesTerra. Analysis output
· Multi-cloud topology generated
· Estimated deployment cost: $4.2k/month
· Governance validation complete
Starting cognition scan… Two ways to work with us
Run it yourself, or let our team run it.
Same platform, same approvals, same audit trail. You choose who operates it.
Your platform team uses OrchesTerra to generate, govern and reconcile infrastructure. From $99 a month, with a 14-day trial.
Repository to Terraform in minutes
Approval chains and policy packs
A dedicated team of our engineers runs your cloud on OrchesTerra: pipelines, patching, drift fixes, DR drills and cost reviews. Every change still waits for your approval.
Named engineers who know your estate
Landing zones on any of 4 clouds
Monthly cost and security report
Incident response on your hours
One control plane from repository to production.
Scans a repository and finds its services, dependencies, Docker Compose, Helm charts, Kubernetes manifests, Terragrunt and CI pipelines, then draws the topology.
Generation Terraform & OpenTofu
Writes reviewable IaC from that topology, provisions through CloudFormation change sets when you prefer, and imports the resources you already run.
Prices every plan against live AWS, Azure, GCP and OCI price lists, and flags savings in what is already running.
Scheduled and predictive scans compare state with the cloud; restore, accept or forget each drift through an approval.
One inbox for every change, with reviewer groups, delegation and maintenance windows. Nothing applies until the chain is complete.
OPA policy packs, a simulator to test them, Checkov config scanning and compliance packs, all enforced in the same pipeline.
Plan replication and standby regions, run failover, and check readiness before the day you need it.
An AI agent that reads your estate and your bill, answers in plain words, and asks before it changes anything.
MULTI-CLOUD STATE RECONCILIATION
Continuously reconcile declared and actual cloud state .
Detect infrastructure drift before it becomes an incident.
Compare Terraform and OpenTofu state against AWS, Azure, GCP, and OCI in real time.
Identify configuration drift and unmanaged resources before changes create operational risk.
Detect drift early, reduce noise, and maintain reliable cloud environments.
Terraform · OpenTofu vs live cloud state
Critical drift events detected
Workloads potentially impacted
Resources outside declared state
Continuously reconciling declared and actual cloud state
Differences between Terraform/OpenTofu state and live cloud resources.
Ingress 0.0.0.0/0:443 added outside Terraform
aws_lb_listener_rule.public_alb
Catch-all listener rule rerouted to public target group
azurerm_network_security_rule.aks_api
AKS API endpoint opened to broad source CIDR block
oci_load_balancer_listener.public
Public listener certificate policy downgraded outside IaC
4 drift events identified during state reconciliation.
From intent to infrastructure.
OrchesTerra understands your intent, asks the right questions, and builds the optimal infrastructure—ready for you to review and launch.
Describe what you need in natural language.
Orbit designs the best-fit architecture in real time.
Review the pre-filled configuration.
Generated from your conversation
Best practices built-in at every step.
Adapts to your inputs in real time.
Security, reliability and compliance first.
Short clips from a live sample workspace: from repository to governed infrastructure, with cost, drift and approvals along the way.
Reopen a scanned repository and land on the architecture OrchesTerra generated from it.
Watch the 2-minute tour Try the live demo Click around a read-only sample workspace for Acme Retail. No signup needed.
Ask questions in plain words. Rehearse the bad day.
Orbit answers questions about your estate and your bill, and drafts changes for approval. Disaster recovery keeps standby regions ready and drills on record.
Why did our AWS bill go up last month?
Cross-region replication ● Healthy
Infrastructure Approval Workflows
Detect infrastructure drift, route approvals, and enforce controlled execution before changes reach production.
Change Type Security group remediation Resource aws_security_group.app_sg Drift Summary Inbound rule 0.0.0.0/0:22 detected in actual state Triggered By Drift detection Approval Chain
All required approvals must complete before execution can proceed.
Your infrastructure, run by a dedicated team.
Our engineers operate your cloud on OrchesTerra, so your developers ship product instead of patching clusters. You keep the approvals and the audit log.
Landing zones on AWS, Azure, GCP and OCI
Terraform modules and state hygiene
CI/CD and GitOps with security gates
Vulnerability and patch management
Drift and policy findings fixed
Compliance evidence from the audit trail
Monthly cost reviews and rightsizing
Incident response on your hours
We scan your repositories and accounts, map drift, risk and spend, and agree the scope.
Existing resources are imported into Terraform state; pipelines and approval chains are set up.
Your named engineers run day-to-day changes, patching and incident response.
A monthly review of cost, security posture, drift and DR readiness.
You stay in control. Every change our team makes waits for your approvers and lands in your audit log.
Replace fragmented infrastructure workflows with a single orchestration layer for generation, approvals, deployment, and reconciliation.
Natural language infrastructure generation
Multi-cloud blueprint generation
Deployment-ready IAC generation
AI-assisted infrastructure reasoning
Managed DevSecOps team on request
Orches Terra is the only platform that unifies infrastructure creation, governance, and operations into a single AI-native orchestration layer .
See seats, cloud accounts, and governance limits on the pricing page .
Built for teams operating production infrastructure with governance, access control, auditability, and secure credential boundaries.
Enterprise patterns baked in, not bolted on.
Least-privilege access, credential isolation, and policy guardrails.
Complete traceability across every action and environment.
Reliable, consistent operations across teams, clouds, and regions.
Route infrastructure changes through configurable approval chains before deployment or reconciliation.
Control access to environments, projects, and infrastructure operations with precision.
Track deployments, approvals, reconciliations, and operator actions with immutable logs.
Integrate cloud credentials securely with scoped permissions and controlled access.
Every deployment, approval, and reconciliation action is fully traceable, auditable, and governed through a single control plane.
Priced per team, not per resource.
Unlimited users · Unlimited cloud accounts
AWS, Azure, Google Cloud and Oracle Cloud (OCI). Cost estimates, drift scans and policy checks run across all four from one workspace.
Generated code is Terraform or OpenTofu, your choice. On AWS you can also provision through CloudFormation change sets and import existing stacks.
No. OrchesTerra imports existing AWS and Azure resources into Terraform state, so drift detection and approvals cover what you run today.
A scoped role per cloud account. On AWS you create it with a one-click CloudFormation template that pins a per-tenant External ID, with separate read-only and deploy roles. Sessions are labelled for CloudTrail and can be locked to source IPs, credentials are encrypted and scoped to your workspace, plans and applies run on a separate sandboxed runner, and every apply waits for the approval chain you configure.
A dedicated team of our engineers runs your infrastructure on OrchesTerra: pipelines, patching, drift and policy fixes, DR drills and cost reviews. Every change they make goes through your approvals and lands in your audit log.
Pick Starter or Business and keep your workspace, or talk to us about Enterprise or Managed DevSecOps. Nothing is applied to your cloud without approval, during the trial or after it.
Launch your infrastructure intent
Describe your infrastructure intent, operational constraints, and deployment priorities — Orches Terra will generate the execution path.
What OrchesTerra will generate
Multi-cloud infrastructure blueprints
Private · Guided · Architecture-specific Infrastructure Intent Company Work email Primary cloud environment Select environment AWS Azure GCP OCI Multi-cloud Team size Select team size 1–10 10–50 50–200 200+ Describe your deployment, drift, or governance challenges 0 /1000 Start Orchestration Deployment-ready outputs. Governed before execution.
AI-native infrastructure orchestration for modern platform teams, and a DevSecOps team that can run it for you.
The AI orchestration engine powering infrastructure intent, governance, and execution.
© 2026 OrchesTerra · The Red Ruby
