---
source: "https://www.firezone.dev/blog/terraform-provider"
hn_url: "https://news.ycombinator.com/item?id=49823125"
title: "Manage Firezone as code with Terraform"
article_title: "Firezone Terraform Provider | Firezone Blog"
image: "https://www.firezone.dev/blog/opengraph-image?e9f6d643a533f056"
author: "jamilbk"
captured_at: "2026-09-23T22:44:09Z"
capture_tool: "hn-digest"
hn_id: 49823125
score: 1
comments: 0
posted_at: "2026-09-23T21:55:43Z"
tags:
  - hacker-news
---

# Manage Firezone as code with Terraform

- HN: [49823125](https://news.ycombinator.com/item?id=49823125)
- Source: [www.firezone.dev](https://www.firezone.dev/blog/terraform-provider)
- Score: 1
- Comments: 0
- Posted: 2026-09-23T21:55:43Z

## Translation

Title: Manage Firezone as code with Terraform
Article title: Firezone Terraform Provider | Firezone Blog
Description: Manage Firezone Portal configuration with Terraform.

Article text:
Firezone Terraform Provider | Firezone Blog Open main menu Product
Manage Firezone as code with Terraform
As we've developed Firezone, we've always wanted it to be useful to customers and teams of all sizes.
For a homelab or a small team, the Firezone Portal gives administrators a straightforward way to configure access. As a team adds more Resources, Groups, Users, and Policies, working in the Portal UI can start to have a time cost. Along with that, if there are multiple administrators for your Firezone account, it becomes difficult to know why someone added a Resource, deleted a Group, or changed a Policy.
To address these issues, we’re announcing the official Firezone Terraform Provider. With the provider, your team can see planned access changes before they reach Firezone, discuss them in a pull request, and apply the same configuration across environments.
Here's an example of how quick it is to apply a configuration:
The recording builds a two-site lab: two Docker networks, private NGINX workloads, and four Gateway containers. It also creates the matching Firezone Sites, Resources, Groups, Actors, memberships, and Policies.
Each firezone_gateway produces an enrollment token that the Docker configuration passes to its matching Gateway container and is brought up automatically. Add a Gateway or rotate its token in code, then let Terraform update the matching container without copying secrets through the Portal or coordinating the rollout by hand.
We're using the Docker provider to keep the demo easy to run locally. The same pattern works with AWS, Azure, Google Cloud, or any infrastructure platform that you manage through Terraform.
Create an API client token in the Firezone Portal under Settings , set it as FIREZONE_TOKEN , and add the provider to your Terraform configuration:
terraform {
required_providers {
firezone = {
source = "firezone/firezone"
}
}
}
provider "firezone" {
endpoint = "https://rest-api.firezone.dev"
}
Terraform reads the token from your environment. See the provider reference for the full configuration.
Sites, Resources, and Policies
The core Firezone access model maps directly onto Terraform resources. This example exposes a production database on its Postgres port to Engineering members who sign in through corporate Okta during business hours from a device with a valid X.509 certificate:
resource "firezone_site" "main" {
name = "us-east"
}
resource "firezone_resource" "database" {
site_id = firezone_site.main.id
name = "postgres-prod"
type = "ip"
address = "10.0.0.100"
address_description = "Production Postgres DB"
filters {
protocol = "tcp"
ports = [ "5432" ]
}
}
data "firezone_group" "engineering" {
name = "Engineering"
}
data "firezone_okta_auth_provider" "corp" {
name = "Okta SSO"
}
resource "firezone_policy" "engineering_db" {
group_id = data .firezone_group.engineering.id
resource_id = firezone_resource.database.id
description = "Engineering access to production Postgres"
condition {
property = "auth_provider_id"
operator = "is_in"
values = [ data .firezone_okta_auth_provider.corp.id]
}
condition {
property = "current_utc_datetime"
operator = "is_in_day_of_week_time_ranges"
values = [
"M/09:00-17:00/America/New_York" ,
"T/09:00-17:00/America/New_York" ,
"W/09:00-17:00/America/New_York" ,
"R/09:00-17:00/America/New_York" ,
"F/09:00-17:00/America/New_York" ,
]
}
condition {
property = "device_attested"
operator = "is"
values = [ "true" ]
}
}
firezone_resource supports CIDR, IP, and DNS Resources. You can add protocol and port filters, and choose an IP family ( ip_stack ) for DNS Resources. Policy conditions can match on:
the device's region (based on its remote IP)
the device's remote IP, by CIDR range
the auth provider the user signed in with
day-of-week time windows in any IANA timezone
whether the device has been verified by an admin
whether the device presented a valid X.509 certificate from one of
your account's trust anchors ( device_attested )
firezone_gateway creates a Gateway and returns its token as a sensitive
attribute. You can pass that token straight to whatever deploys the
Gateway host, such as our
AWS ,
Azure , or
Google Cloud Gateway modules:
resource "firezone_gateway" "gw" {
for_each = {
"a" = {
name = "gw-us-east-1a"
rotated = "2026-09-14"
}
"b" = {
name = "gw-us-east-1b"
rotated = "2026-09-14"
}
}
site_id = firezone_site.main.id
name = each .value.name
token_rotation_trigger = each .value.rotated
}
module "gateway" {
source = "firezone/gateway/aws"
firezone_tokens = [for gw in firezone_gateway.gw : gw.token]
# ... VPC, subnet, AMI, and security group inputs
}
token_rotation_trigger rotates the token whenever its value changes. Update a Gateway's rotated value when you want to rotate its token. The value is arbitrary, but a date makes the rotation history easy to read. The old token keeps working until the Gateway connects with the new one or a grace period runs out, so the Gateway host needs to pick up the new token within that window.
Terraform stores a Gateway token in state when it creates or rotates a Gateway, so protect your remote state backend. You can pass the sensitive token to AWS Secrets Manager or another secret store with an ordinary attribute reference. On Terraform 1.11 or newer, use a write-only argument such as secret_string_wo so the secret store's copy does not add another state entry.
Authentication providers and directories stay in the Portal because their setup can require an administrator to complete an interactive OAuth authorization flow. After setup, Terraform exposes those integrations as data sources that you can reference in Policies.
Devices remain read-only in Terraform. You can reference a Device and manage members of an imported static device pool, but Devices enroll themselves and management remains in the Portal.
Service accounts have a separate credential step. An administrator can create a service account with Terraform, then create the device token that lets a headless Device sign in through the Portal.
The Portal also remains the place to inspect a running Gateway. It shows whether the Gateway is online, which IP addresses it has, and which version it runs. Account settings, logs, external identities, the X.509 authentication provider, and device posture integrations remain outside the provider.
See the provider reference for the full resource and data-source list.
The provider is available now on the Terraform Registry . The full reference and examples for every resource and data source are in the docs and the source is on GitHub .
If one of the gaps above is blocking you, please open an issue and tell us what you're trying to automate. That feedback directly shapes what we build next.
Sign up with your email to receive roadmap updates, how-tos, and product announcements from the Firezone team.
By checking the box below, you agree to receive communications from Firezone. You can unsubscribe anytime.
To deliver your service, we need your permission to store and process your personal data. We care about your privacy. Learn how we handle your data in our Privacy Policy .
Give your organization the protection it deserves.
Get a personalized walkthrough.
WireGuard is a registered trademark of Jason A. Donenfeld.
Firezone is a registered trademark of Firezone, Inc.

## Original Extract

Manage Firezone Portal configuration with Terraform.

Firezone Terraform Provider | Firezone Blog Open main menu Product
Manage Firezone as code with Terraform
As we've developed Firezone, we've always wanted it to be useful to customers and teams of all sizes.
For a homelab or a small team, the Firezone Portal gives administrators a straightforward way to configure access. As a team adds more Resources, Groups, Users, and Policies, working in the Portal UI can start to have a time cost. Along with that, if there are multiple administrators for your Firezone account, it becomes difficult to know why someone added a Resource, deleted a Group, or changed a Policy.
To address these issues, we’re announcing the official Firezone Terraform Provider. With the provider, your team can see planned access changes before they reach Firezone, discuss them in a pull request, and apply the same configuration across environments.
Here's an example of how quick it is to apply a configuration:
The recording builds a two-site lab: two Docker networks, private NGINX workloads, and four Gateway containers. It also creates the matching Firezone Sites, Resources, Groups, Actors, memberships, and Policies.
Each firezone_gateway produces an enrollment token that the Docker configuration passes to its matching Gateway container and is brought up automatically. Add a Gateway or rotate its token in code, then let Terraform update the matching container without copying secrets through the Portal or coordinating the rollout by hand.
We're using the Docker provider to keep the demo easy to run locally. The same pattern works with AWS, Azure, Google Cloud, or any infrastructure platform that you manage through Terraform.
Create an API client token in the Firezone Portal under Settings , set it as FIREZONE_TOKEN , and add the provider to your Terraform configuration:
terraform {
required_providers {
firezone = {
source = "firezone/firezone"
}
}
}
provider "firezone" {
endpoint = "https://rest-api.firezone.dev"
}
Terraform reads the token from your environment. See the provider reference for the full configuration.
Sites, Resources, and Policies
The core Firezone access model maps directly onto Terraform resources. This example exposes a production database on its Postgres port to Engineering members who sign in through corporate Okta during business hours from a device with a valid X.509 certificate:
resource "firezone_site" "main" {
name = "us-east"
}
resource "firezone_resource" "database" {
site_id = firezone_site.main.id
name = "postgres-prod"
type = "ip"
address = "10.0.0.100"
address_description = "Production Postgres DB"
filters {
protocol = "tcp"
ports = [ "5432" ]
}
}
data "firezone_group" "engineering" {
name = "Engineering"
}
data "firezone_okta_auth_provider" "corp" {
name = "Okta SSO"
}
resource "firezone_policy" "engineering_db" {
group_id = data .firezone_group.engineering.id
resource_id = firezone_resource.database.id
description = "Engineering access to production Postgres"
condition {
property = "auth_provider_id"
operator = "is_in"
values = [ data .firezone_okta_auth_provider.corp.id]
}
condition {
property = "current_utc_datetime"
operator = "is_in_day_of_week_time_ranges"
values = [
"M/09:00-17:00/America/New_York" ,
"T/09:00-17:00/America/New_York" ,
"W/09:00-17:00/America/New_York" ,
"R/09:00-17:00/America/New_York" ,
"F/09:00-17:00/America/New_York" ,
]
}
condition {
property = "device_attested"
operator = "is"
values = [ "true" ]
}
}
firezone_resource supports CIDR, IP, and DNS Resources. You can add protocol and port filters, and choose an IP family ( ip_stack ) for DNS Resources. Policy conditions can match on:
the device's region (based on its remote IP)
the device's remote IP, by CIDR range
the auth provider the user signed in with
day-of-week time windows in any IANA timezone
whether the device has been verified by an admin
whether the device presented a valid X.509 certificate from one of
your account's trust anchors ( device_attested )
firezone_gateway creates a Gateway and returns its token as a sensitive
attribute. You can pass that token straight to whatever deploys the
Gateway host, such as our
AWS ,
Azure , or
Google Cloud Gateway modules:
resource "firezone_gateway" "gw" {
for_each = {
"a" = {
name = "gw-us-east-1a"
rotated = "2026-09-14"
}
"b" = {
name = "gw-us-east-1b"
rotated = "2026-09-14"
}
}
site_id = firezone_site.main.id
name = each .value.name
token_rotation_trigger = each .value.rotated
}
module "gateway" {
source = "firezone/gateway/aws"
firezone_tokens = [for gw in firezone_gateway.gw : gw.token]
# ... VPC, subnet, AMI, and security group inputs
}
token_rotation_trigger rotates the token whenever its value changes. Update a Gateway's rotated value when you want to rotate its token. The value is arbitrary, but a date makes the rotation history easy to read. The old token keeps working until the Gateway connects with the new one or a grace period runs out, so the Gateway host needs to pick up the new token within that window.
Terraform stores a Gateway token in state when it creates or rotates a Gateway, so protect your remote state backend. You can pass the sensitive token to AWS Secrets Manager or another secret store with an ordinary attribute reference. On Terraform 1.11 or newer, use a write-only argument such as secret_string_wo so the secret store's copy does not add another state entry.
Authentication providers and directories stay in the Portal because their setup can require an administrator to complete an interactive OAuth authorization flow. After setup, Terraform exposes those integrations as data sources that you can reference in Policies.
Devices remain read-only in Terraform. You can reference a Device and manage members of an imported static device pool, but Devices enroll themselves and management remains in the Portal.
Service accounts have a separate credential step. An administrator can create a service account with Terraform, then create the device token that lets a headless Device sign in through the Portal.
The Portal also remains the place to inspect a running Gateway. It shows whether the Gateway is online, which IP addresses it has, and which version it runs. Account settings, logs, external identities, the X.509 authentication provider, and device posture integrations remain outside the provider.
See the provider reference for the full resource and data-source list.
The provider is available now on the Terraform Registry . The full reference and examples for every resource and data source are in the docs and the source is on GitHub .
If one of the gaps above is blocking you, please open an issue and tell us what you're trying to automate. That feedback directly shapes what we build next.
Sign up with your email to receive roadmap updates, how-tos, and product announcements from the Firezone team.
By checking the box below, you agree to receive communications from Firezone. You can unsubscribe anytime.
To deliver your service, we need your permission to store and process your personal data. We care about your privacy. Learn how we handle your data in our Privacy Policy .
Give your organization the protection it deserves.
Get a personalized walkthrough.
WireGuard is a registered trademark of Jason A. Donenfeld.
Firezone is a registered trademark of Firezone, Inc.
