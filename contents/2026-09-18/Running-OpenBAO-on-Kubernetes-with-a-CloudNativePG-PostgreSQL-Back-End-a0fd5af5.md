---
source: "https://www.cncf.io/blog/2026/09/16/running-openbao-on-kubernetes-with-a-cloudnativepg-postgresql-backend/"
hn_url: "https://news.ycombinator.com/item?id=49749784"
title: "Running OpenBAO on Kubernetes with a CloudNativePG PostgreSQL Back End"
article_title: "Running OpenBao on Kubernetes with a CloudNativePG PostgreSQL backend | CNCF"
image: "https://www.cncf.io/wp-content/uploads/2026/09/Running-OpenBao-on-K8s-1.jpg"
author: "geoffbp"
captured_at: "2026-09-18T03:29:45Z"
capture_tool: "hn-digest"
hn_id: 49749784
score: 1
comments: 0
posted_at: "2026-09-18T03:07:44Z"
tags:
  - hacker-news
---

# Running OpenBAO on Kubernetes with a CloudNativePG PostgreSQL Back End

- HN: [49749784](https://news.ycombinator.com/item?id=49749784)
- Source: [www.cncf.io](https://www.cncf.io/blog/2026/09/16/running-openbao-on-kubernetes-with-a-cloudnativepg-postgresql-backend/)
- Score: 1
- Comments: 0
- Posted: 2026-09-18T03:07:44Z

## Translation

Title: Running OpenBAO on Kubernetes with a CloudNativePG PostgreSQL Back End
Article title: Running OpenBao on Kubernetes with a CloudNativePG PostgreSQL backend | CNCF
Description: Managing infrastructure secrets on Kubernetes needs a backend that is self-healing and free of vendor lock-in, and that is exactly what OpenBao (the Linux…

Article text:
Skip to content
Accessibility
help
KubeCon + CloudNativeCon NA 2026 · Nov 9 - 12 · Salt Lake City · REGISTER NOW
Membership Hub – For Current Members Learn about all the benefits of being a CNCF Member
Technical Oversight Committee The TOC defines CNCF’s technical vision and provides experienced technical leadership to the cloud native community
Governing Board The GB is responsible for marketing, business oversight, and budget decisions for CNCF
End User Technical Advisory Board The End User TAB serves as the voice of the end users in CNCF community decisions
Ambassadors Meet our Ambassadors—experienced practitioners passionate about helping others learn about cloud native technologies
Graduated Projects considered stable, widely adopted, and production ready, attracting thousands of contributors
Incubating Projects used successfully in production by a small number users with a healthy pool of contributors
Sandbox Experimental projects not yet widely tested in production on the bleeding edge of technology
Archived Projects that have reached the end of their lifecycle and have become inactive
Project Metrics View metrics of CNCF projects moving through maturity levels
Contribute Join the 150K+ folx in #TeamCloudNative who’ve contributed their expertise to CNCF hosted projects
Services for CNCF Projects CNCF services for our open source projects – from marketing to legal services
Cloud Native Landscape A comprehensive categorical overview of projects and product offerings in the cloud native space
Project Journey Reports Showing how CNCF has impacted the progress and growth of various graduated projects
Project Tools Quick links to tools and resources for your CNCF project
Latest Project Journey Reports
Training Overview Find your path to cloud native success with training and certificates from the pioneer of cloud-native technology
Certifications Get certified by the only authoritative source for cloud-native certification accepted by industry
Courses Learn the basics of cloud native or train for a certification with training courses built in collaboration with the Linux Foundation
Kubestronaut Program Uplevel your Kubernetes skills
Kubernetes Training Partners Find a qualified KCNTP to prepare for your next certification
Certified Kubernetes Software conformance ensures your versions of CNCF projects support the required APIs
Cloud Native Network Function Certification (Beta) CNF Certification ensures applications demonstrate cloud native best practices
Kubernetes Certified Service Provider KCSPs have deep experience helping enterprises successfully adopt cloud native technologies
Enroll your company as a CNCF End User and save more than $10K in training and conference costs
End User Community Join our vendor-neutral community using cloud native technologies to build products and services
Events Meet #TeamCloudNative and CNCF staff at events around the world
Case Studies Read real-world case studies about the impact cloud native projects are having on organizations around the world
Humans of Cloud Native Read stories of amazing individuals and their contributions
The Cloud Native Heroes Challenge Be a cloud native hero! Help us defeat patent trolls to earn swag and prizes
Online Programs Watch our free online programs for the latest insights into cloud native technologies and projects
Community Groups Join #TeamCloudNative at events and meetups near you
Phippy & Friends Phippy explains core cloud native concepts in simple terms through stories perfect for all ages
Cloud Native Glossary Explore cloud native concepts in clear and simple language – no technical knowledge required!
Blog Catch up on the latest happenings and technical insights from #TeamCloudNative
Announcements Media releases and official CNCF announcements
News CNCF projects and #TeamCloudNative in the media
Reports Read transparent, in-depth reports on our organization, events, and projects
Search
CNCF
Blog
/
Ambassador Post
Running OpenBao on Kubernetes with a CloudNativePG PostgreSQL backend
Posted on September 16, 2026
by Gabriele Bartolini (EnterpriseDB) and CNCF Ambassador, Rob Kenefeck (ControlPlane)
CNCF projects highlighted in this post
Managing infrastructure secrets on Kubernetes needs a backend that is self-healing and free of vendor lock-in, and that is exactly what OpenBao (the Linux Foundation’s open-source fork of HashiCorp Vault) and CloudNativePG give you: an entirely open-source stack built on two CNCF projects, Kubernetes , long since graduated, and CloudNativePG, a CNCF Sandbox project currently under evaluation for Incubation by the CNCF Technical Oversight Committee. OpenBao’s postgresql storage backend turns any PostgreSQL cluster into its encrypted key-value store, and CloudNativePG turns that cluster into a self-healing, synchronously replicated, certificate-authenticated Postgres instance with no cloud database dependency underneath it.
This recipe deploys a three-instance CNPG cluster as OpenBao’s storage backend and removes every password from the connection: the schema-owning role and the application role OpenBao itself uses both authenticate with a DatabaseRole-issued TLS client certificate, enforced by explicit pg_hba rules rather than by the absence of a password. pg_hba.conf is PostgreSQL’s client-authentication file, the thing that actually decides, per connection, whether a role needs a certificate, a password, or nothing at all.
Setting up a local test environment with cnpg-playground
Nothing about this recipe is specific to any one Kubernetes distribution: any conformant cluster with enough worker capacity will do. To follow along locally, though, the official cnpg-playground repository is the fastest path to one, since it is pre-configured with the CloudNativePG operator already. It is designed primarily around CNPG’s own demos, so it is worth knowing what it actually gives you: a single Kind cluster with six nodes, a control plane node, one node labelled for infrastructure workloads, one labelled for application workloads, and three carrying a node-role.kubernetes.io/postgres taint. That taint is exactly what our Cluster ‘s tolerations in Step 1 target, and it is also what leaves OpenBao itself with only the two general-purpose nodes to schedule onto, which matters once pod anti-affinity enters the picture in Step 3. setup.sh provisions one Kind cluster per argument it is given, normally used to model separate regions; passing it a single, arbitrary label gives you one local cluster and skips the two-region disaster recovery demo entirely.
Prerequisites: Docker , Kind , Helm and kubectl .
# Clone the CNPG Playground repository
git clone https://github.com/cloudnative-pg/cnpg-playground.git
cd cnpg-playground
# 1. Provision a single local cluster labelled "openbao"
./scripts/setup.sh openbao
# 2. Deploy CloudNativePG, cert-manager, the Barman Cloud plugin and a
# ClusterImageCatalog only, skipping the demo databases
REQUIREMENTS_ONLY=true ./demo/setup.sh
Architecture blueprint
Storage engine: OpenBao’s native postgresql storage backend, with ha_enabled = "true" for its HA lock table.
Database cluster: a 3-instance CNPG cluster with quorum-based synchronous replication ( method: any, number: 1, the default dataDurability: required ) for zero-data-loss failover.
Workload isolation: node selectors, tolerations and required zonal pod anti-affinity keep PostgreSQL on dedicated nodes across separate failure domains, following CNPG’s scheduling guidance .
Authentication: passwordless mTLS via the DatabaseRole CRD’s clientCertificate block, for both the schema owner and the application role, enforced by explicit pg_hba rules.
Step 1: deploy the CNPG cluster, roles and database
The Cluster below points imageCatalogRef at the postgresql-minimal-trixie ClusterImageCatalog that REQUIREMENTS_ONLY=true ./demo/setup.sh already deployed in the previous step, rather than pinning an image tag directly: CNPG resolves it to the latest minimal PostgreSQL 18 image in that catalog, so a kubectl apply against the same manifest keeps picking up new patch releases as the catalog is updated, no Cluster edit required. It also declares synchronous replication, workload isolation, and the two pg_hba rules that force certificate authentication for both roles OpenBao will use. Two DatabaseRole objects follow: role-openbao , the schema owner used once to run DDL, and role-openbao-rw , the restricted role OpenBao itself connects as at runtime. Both get a clientCertificate , because a one-shot DDL job is no more entitled to a password lying around than the application is.
{{apiVersion: postgresql.cnpg.io/v1
kind: Cluster
metadata:
name: openbao-db
namespace: openbao
spec:
instances: 3
# Tracks the latest minimal PostgreSQL 18 image via the ClusterImageCatalog
# the playground's REQUIREMENTS_ONLY step already deploys.
# See https://cloudnative-pg.io/docs/current/image_catalog
imageCatalogRef:
apiGroup: postgresql.cnpg.io
kind: ClusterImageCatalog
name: postgresql-minimal-trixie
major: 18
# See https://cloudnative-pg.io/docs/current/scheduling
affinity:
nodeSelector:
node-role.kubernetes.io/postgres: ""
tolerations:
- key: node-role.kubernetes.io/postgres
operator: Exists
effect: NoSchedule
enablePodAntiAffinity: true
topologyKey: topology.kubernetes.io/zone
podAntiAffinityType: required
postgresql:
# Synchronous replication: dataDurability defaults to "required", giving
# RPO=0 at the cost of pausing writes if no standby is available.
# See https://cloudnative-pg.io/docs/current/replication
synchronous:
method: any
number: 1
# The operator does not add cert rules for DatabaseRole client
# certificates automatically: without these, "openbao" and "openbao-rw"
# would fall through to the default scram-sha-256 rule, and since
# neither role has a passwordSecret, every connection would simply fail.
pg_hba:
- hostssl openbao openbao all cert
- hostssl openbao openbao-rw all cert
- hostnossl openbao openbao all reject
- hostnossl openbao openbao-rw all reject
# See https://cloudnative-pg.io/docs/current/postgresql_conf
parameters:
max_connections: '100'
log_checkpoints: 'on'
log_lock_waits: 'on'
hot_standby_feedback: 'on'
shared_memory_type: 'sysv'
dynamic_shared_memory_type: 'sysv'
storage:
size: 10Gi
---
apiVersion: postgresql.cnpg.io/v1
kind: DatabaseRole
metadata:
name: role-openbao
namespace: openbao
spec:
cluster:
name: openbao-db
name: openbao
login: true
clientCertificate:
enabled: true
databaseRoleReclaimPolicy: retain
---
apiVersion: postgresql.cnpg.io/v1
kind: DatabaseRole
metadata:
name: role-openbao-rw
namespace: openbao
spec:
c
[truncated]
Both standbys show up as Standby (sync) with a Sync State of quorum at the same time, which is exactly the dynamic behaviour method: any is meant to give: with number: 1 , either standby satisfies durability, and CNPG does not pin a fixed “the” synchronous standby.
Once reconciled, the operator has created two client certificate secrets, role-openbao-client-cert and role-openbao-rw-client-cert , following its <databaserole-name>-client-cert naming convention. openbao , as the database owner, already has CREATE on the public schema by default (PostgreSQL grants that to the owner even though it revoked it from PUBLIC in v15), so no extra schema grant is needed before the DDL step.
Every manifest that mounts one of these secrets sets defaultMode: 0640 on the volume. Kubernetes mounts Secret volumes at 0644 by default, which libpq refuses outright: it rejects a private key file that is group-or-world-readable, whether owned by root ( 0640 or less) or by the connecting user ( 0600 or less). Since the mounted files stay root-owned and only their group matches the pod’s fsGroup, 0640 is the setting that satisfies libpq here, and it applies to every pod in this recipe that reads a client certificate, the schema-init Job and the OpenBao pods alike.
Step 2: initialise the schema and grant table privileges
DatabaseRole does not yet manage table-level grants: the permissi

[truncated]

## Original Extract

Managing infrastructure secrets on Kubernetes needs a backend that is self-healing and free of vendor lock-in, and that is exactly what OpenBao (the Linux…

Skip to content
Accessibility
help
KubeCon + CloudNativeCon NA 2026 · Nov 9 - 12 · Salt Lake City · REGISTER NOW
Membership Hub – For Current Members Learn about all the benefits of being a CNCF Member
Technical Oversight Committee The TOC defines CNCF’s technical vision and provides experienced technical leadership to the cloud native community
Governing Board The GB is responsible for marketing, business oversight, and budget decisions for CNCF
End User Technical Advisory Board The End User TAB serves as the voice of the end users in CNCF community decisions
Ambassadors Meet our Ambassadors—experienced practitioners passionate about helping others learn about cloud native technologies
Graduated Projects considered stable, widely adopted, and production ready, attracting thousands of contributors
Incubating Projects used successfully in production by a small number users with a healthy pool of contributors
Sandbox Experimental projects not yet widely tested in production on the bleeding edge of technology
Archived Projects that have reached the end of their lifecycle and have become inactive
Project Metrics View metrics of CNCF projects moving through maturity levels
Contribute Join the 150K+ folx in #TeamCloudNative who’ve contributed their expertise to CNCF hosted projects
Services for CNCF Projects CNCF services for our open source projects – from marketing to legal services
Cloud Native Landscape A comprehensive categorical overview of projects and product offerings in the cloud native space
Project Journey Reports Showing how CNCF has impacted the progress and growth of various graduated projects
Project Tools Quick links to tools and resources for your CNCF project
Latest Project Journey Reports
Training Overview Find your path to cloud native success with training and certificates from the pioneer of cloud-native technology
Certifications Get certified by the only authoritative source for cloud-native certification accepted by industry
Courses Learn the basics of cloud native or train for a certification with training courses built in collaboration with the Linux Foundation
Kubestronaut Program Uplevel your Kubernetes skills
Kubernetes Training Partners Find a qualified KCNTP to prepare for your next certification
Certified Kubernetes Software conformance ensures your versions of CNCF projects support the required APIs
Cloud Native Network Function Certification (Beta) CNF Certification ensures applications demonstrate cloud native best practices
Kubernetes Certified Service Provider KCSPs have deep experience helping enterprises successfully adopt cloud native technologies
Enroll your company as a CNCF End User and save more than $10K in training and conference costs
End User Community Join our vendor-neutral community using cloud native technologies to build products and services
Events Meet #TeamCloudNative and CNCF staff at events around the world
Case Studies Read real-world case studies about the impact cloud native projects are having on organizations around the world
Humans of Cloud Native Read stories of amazing individuals and their contributions
The Cloud Native Heroes Challenge Be a cloud native hero! Help us defeat patent trolls to earn swag and prizes
Online Programs Watch our free online programs for the latest insights into cloud native technologies and projects
Community Groups Join #TeamCloudNative at events and meetups near you
Phippy & Friends Phippy explains core cloud native concepts in simple terms through stories perfect for all ages
Cloud Native Glossary Explore cloud native concepts in clear and simple language – no technical knowledge required!
Blog Catch up on the latest happenings and technical insights from #TeamCloudNative
Announcements Media releases and official CNCF announcements
News CNCF projects and #TeamCloudNative in the media
Reports Read transparent, in-depth reports on our organization, events, and projects
Search
CNCF
Blog
/
Ambassador Post
Running OpenBao on Kubernetes with a CloudNativePG PostgreSQL backend
Posted on September 16, 2026
by Gabriele Bartolini (EnterpriseDB) and CNCF Ambassador, Rob Kenefeck (ControlPlane)
CNCF projects highlighted in this post
Managing infrastructure secrets on Kubernetes needs a backend that is self-healing and free of vendor lock-in, and that is exactly what OpenBao (the Linux Foundation’s open-source fork of HashiCorp Vault) and CloudNativePG give you: an entirely open-source stack built on two CNCF projects, Kubernetes , long since graduated, and CloudNativePG, a CNCF Sandbox project currently under evaluation for Incubation by the CNCF Technical Oversight Committee. OpenBao’s postgresql storage backend turns any PostgreSQL cluster into its encrypted key-value store, and CloudNativePG turns that cluster into a self-healing, synchronously replicated, certificate-authenticated Postgres instance with no cloud database dependency underneath it.
This recipe deploys a three-instance CNPG cluster as OpenBao’s storage backend and removes every password from the connection: the schema-owning role and the application role OpenBao itself uses both authenticate with a DatabaseRole-issued TLS client certificate, enforced by explicit pg_hba rules rather than by the absence of a password. pg_hba.conf is PostgreSQL’s client-authentication file, the thing that actually decides, per connection, whether a role needs a certificate, a password, or nothing at all.
Setting up a local test environment with cnpg-playground
Nothing about this recipe is specific to any one Kubernetes distribution: any conformant cluster with enough worker capacity will do. To follow along locally, though, the official cnpg-playground repository is the fastest path to one, since it is pre-configured with the CloudNativePG operator already. It is designed primarily around CNPG’s own demos, so it is worth knowing what it actually gives you: a single Kind cluster with six nodes, a control plane node, one node labelled for infrastructure workloads, one labelled for application workloads, and three carrying a node-role.kubernetes.io/postgres taint. That taint is exactly what our Cluster ‘s tolerations in Step 1 target, and it is also what leaves OpenBao itself with only the two general-purpose nodes to schedule onto, which matters once pod anti-affinity enters the picture in Step 3. setup.sh provisions one Kind cluster per argument it is given, normally used to model separate regions; passing it a single, arbitrary label gives you one local cluster and skips the two-region disaster recovery demo entirely.
Prerequisites: Docker , Kind , Helm and kubectl .
# Clone the CNPG Playground repository
git clone https://github.com/cloudnative-pg/cnpg-playground.git
cd cnpg-playground
# 1. Provision a single local cluster labelled "openbao"
./scripts/setup.sh openbao
# 2. Deploy CloudNativePG, cert-manager, the Barman Cloud plugin and a
# ClusterImageCatalog only, skipping the demo databases
REQUIREMENTS_ONLY=true ./demo/setup.sh
Architecture blueprint
Storage engine: OpenBao’s native postgresql storage backend, with ha_enabled = "true" for its HA lock table.
Database cluster: a 3-instance CNPG cluster with quorum-based synchronous replication ( method: any, number: 1, the default dataDurability: required ) for zero-data-loss failover.
Workload isolation: node selectors, tolerations and required zonal pod anti-affinity keep PostgreSQL on dedicated nodes across separate failure domains, following CNPG’s scheduling guidance .
Authentication: passwordless mTLS via the DatabaseRole CRD’s clientCertificate block, for both the schema owner and the application role, enforced by explicit pg_hba rules.
Step 1: deploy the CNPG cluster, roles and database
The Cluster below points imageCatalogRef at the postgresql-minimal-trixie ClusterImageCatalog that REQUIREMENTS_ONLY=true ./demo/setup.sh already deployed in the previous step, rather than pinning an image tag directly: CNPG resolves it to the latest minimal PostgreSQL 18 image in that catalog, so a kubectl apply against the same manifest keeps picking up new patch releases as the catalog is updated, no Cluster edit required. It also declares synchronous replication, workload isolation, and the two pg_hba rules that force certificate authentication for both roles OpenBao will use. Two DatabaseRole objects follow: role-openbao , the schema owner used once to run DDL, and role-openbao-rw , the restricted role OpenBao itself connects as at runtime. Both get a clientCertificate , because a one-shot DDL job is no more entitled to a password lying around than the application is.
{{apiVersion: postgresql.cnpg.io/v1
kind: Cluster
metadata:
name: openbao-db
namespace: openbao
spec:
instances: 3
# Tracks the latest minimal PostgreSQL 18 image via the ClusterImageCatalog
# the playground's REQUIREMENTS_ONLY step already deploys.
# See https://cloudnative-pg.io/docs/current/image_catalog
imageCatalogRef:
apiGroup: postgresql.cnpg.io
kind: ClusterImageCatalog
name: postgresql-minimal-trixie
major: 18
# See https://cloudnative-pg.io/docs/current/scheduling
affinity:
nodeSelector:
node-role.kubernetes.io/postgres: ""
tolerations:
- key: node-role.kubernetes.io/postgres
operator: Exists
effect: NoSchedule
enablePodAntiAffinity: true
topologyKey: topology.kubernetes.io/zone
podAntiAffinityType: required
postgresql:
# Synchronous replication: dataDurability defaults to "required", giving
# RPO=0 at the cost of pausing writes if no standby is available.
# See https://cloudnative-pg.io/docs/current/replication
synchronous:
method: any
number: 1
# The operator does not add cert rules for DatabaseRole client
# certificates automatically: without these, "openbao" and "openbao-rw"
# would fall through to the default scram-sha-256 rule, and since
# neither role has a passwordSecret, every connection would simply fail.
pg_hba:
- hostssl openbao openbao all cert
- hostssl openbao openbao-rw all cert
- hostnossl openbao openbao all reject
- hostnossl openbao openbao-rw all reject
# See https://cloudnative-pg.io/docs/current/postgresql_conf
parameters:
max_connections: '100'
log_checkpoints: 'on'
log_lock_waits: 'on'
hot_standby_feedback: 'on'
shared_memory_type: 'sysv'
dynamic_shared_memory_type: 'sysv'
storage:
size: 10Gi
---
apiVersion: postgresql.cnpg.io/v1
kind: DatabaseRole
metadata:
name: role-openbao
namespace: openbao
spec:
cluster:
name: openbao-db
name: openbao
login: true
clientCertificate:
enabled: true
databaseRoleReclaimPolicy: retain
---
apiVersion: postgresql.cnpg.io/v1
kind: DatabaseRole
metadata:
name: role-openbao-rw
namespace: openbao
spec:
c
[truncated]
Both standbys show up as Standby (sync) with a Sync State of quorum at the same time, which is exactly the dynamic behaviour method: any is meant to give: with number: 1 , either standby satisfies durability, and CNPG does not pin a fixed “the” synchronous standby.
Once reconciled, the operator has created two client certificate secrets, role-openbao-client-cert and role-openbao-rw-client-cert , following its <databaserole-name>-client-cert naming convention. openbao , as the database owner, already has CREATE on the public schema by default (PostgreSQL grants that to the owner even though it revoked it from PUBLIC in v15), so no extra schema grant is needed before the DDL step.
Every manifest that mounts one of these secrets sets defaultMode: 0640 on the volume. Kubernetes mounts Secret volumes at 0644 by default, which libpq refuses outright: it rejects a private key file that is group-or-world-readable, whether owned by root ( 0640 or less) or by the connecting user ( 0600 or less). Since the mounted files stay root-owned and only their group matches the pod’s fsGroup, 0640 is the setting that satisfies libpq here, and it applies to every pod in this recipe that reads a client certificate, the schema-init Job and the OpenBao pods alike.
Step 2: initialise the schema and grant table privileges
DatabaseRole does not yet manage table-level grants: the permissi

[truncated]
