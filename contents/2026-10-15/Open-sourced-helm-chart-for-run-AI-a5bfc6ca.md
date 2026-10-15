---
source: "https://github.com/logicalclocks/run-ai-helm"
hn_url: "https://news.ycombinator.com/item?id=50102529"
title: "Open-sourced helm chart for run:AI"
article_title: "GitHub - logicalclocks/run-ai-helm: Helm chart that installs NVIDIA Run:ai (control plane + cluster) and KAI Scheduler · GitHub"
image: "https://opengraph.githubassets.com/b96d580dec2568f82faf0c35350f0e3e9d5d473982897d3d9a9e4d9308d69c81/logicalclocks/run-ai-helm"
author: "jamesblonde"
captured_at: "2026-10-15T13:07:24Z"
capture_tool: "hn-digest"
hn_id: 50102529
score: 1
comments: 0
posted_at: "2026-10-15T12:11:30Z"
tags:
  - hacker-news
---

# Open-sourced helm chart for run:AI

- HN: [50102529](https://news.ycombinator.com/item?id=50102529)
- Source: [github.com](https://github.com/logicalclocks/run-ai-helm)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T12:11:30Z

## Translation

Title: Open-sourced helm chart for run:AI
Article title: GitHub - logicalclocks/run-ai-helm: Helm chart that installs NVIDIA Run:ai (control plane + cluster) and KAI Scheduler · GitHub
Description: Helm chart that installs NVIDIA Run:ai (control plane + cluster) and KAI Scheduler - logicalclocks/run-ai-helm

Article text:
GitHub - logicalclocks/run-ai-helm: Helm chart that installs NVIDIA Run:ai (control plane + cluster) and KAI Scheduler · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
logicalclocks
run-ai-helm Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Helm chart that installs NVIDIA Run:ai (control plane + cluster) and KAI Scheduler
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.github .github charts charts scripts scripts templates templates .gitignore .gitignore .helmignore .helmignore CLAUDE.md CLAUDE.md Chart.lock Chart.lock Chart.yaml Chart.yaml LICENSE LICENSE README.md README.md README.md.gotmpl README.md.gotmpl helm_dependencies.sh helm_dependencies.sh values.cluster.yaml values.cluster.yaml values.control-plane.yaml values.control-plane.yaml values.dev.yaml values.dev.yaml values.kai-scheduler.yaml values.kai-scheduler.yaml values.openshift.yaml values.openshift.yaml values.prerequisites.yaml values.prerequisites.yaml values.schema.json values.schema.json values.yaml values.yaml View all files Repository files navigation
Installs NVIDIA Run:ai (control plane and cluster) and the KAI Scheduler it schedules with
The chart wraps the published upstream charts and adds the glue between them: release
profiles, validation, the Secrets the upstream charts expect to exist, and the cluster
registration step.
The chart is installed once per namespace. Each release enables its components with a profile
file:
The control plane and the cluster cannot share a release. The cluster's pre-install hooks call
the control-plane API with a cluster UUID and a bootstrap token, and the control plane issues
both only once it is running. scripts/register_cluster.sh registers the cluster and writes them
to a values file for release 3.
KAI Scheduler comes with release 3: runai-cluster depends on kai-resource-management , which
depends on kai-scheduler . The Run:ai operator writes KAI's Config and the
kai-resource-management node-pool controller writes one SchedulingShard per node pool. The
standalone kai-scheduler profile is for clusters that run KAI without Run:ai,
and the chart refuses to render it together with the cluster profile.
Kubernetes >= 1.30, Helm 3.x or 4.x, kubectl , curl , jq .
An ingress controller. The control plane creates an Ingress with class nginx by default
( global.ingress.ingressClass , nginx or haproxy ). OpenShift uses Routes.
A DNS name for the control plane ( global.domain ) that resolves to the ingress controller, both
from where scripts/install.sh runs and from inside the cluster.
A TLS certificate for global.domain : PEM in controlplane.tls , a cert-manager issuer in
controlplane.tls.certManager , or a kubernetes.io/tls Secret runai-backend-tls created in
namespace runai-backend before the install.
A default StorageClass for the PostgreSQL, Thanos (100Gi) and NATS volumes.
A Prometheus Operator and an NVIDIA GPU stack (DCGM exporter). Release 1 installs both; skip it
with --skip-prerequisites when the cluster already has them.
./helm_dependencies.sh
scripts/install.sh --domain runai.example.com -f my-values.yaml
my-values.yaml is passed to every release. A minimal one with a certificate for the domain:
controlplane :
tls :
certificate : |
-----BEGIN CERTIFICATE-----
...
key : |
-----BEGIN PRIVATE KEY-----
...
When the certificate is signed by a private CA, also set the CA so the cluster components trust
it:
global :
customCA :
enabled : true
_runai :
customCA :
certificate : |
-----BEGIN CERTIFICATE-----
...
and pass --cacert ca.pem to scripts/install.sh for the registration calls.
scripts/install.sh --help lists the options: --skip-prerequisites , --skip-control-plane ,
--skip-cluster , --dev (fake GPUs), --cluster-name , --admin-user , --cacert ,
--insecure , --timeout .
The UI is at https://<global.domain> . The admin user is test@run.ai
( controlplane.control-plane.tenantsManager.config.adminUsername ) and its password is generated:
kubectl get secret runai-backend-gen-tenants-manager -n runai-backend \
-o jsonpath= ' {.data.ADMIN_PASSWORD} ' | base64 -d
Step by step
helm upgrade --install runai-prerequisites . -n runai-prerequisites --create-namespace \
-f values.prerequisites.yaml --wait
helm upgrade --install runai-backend . -n runai-backend --create-namespace \
-f values.control-plane.yaml -f my-values.yaml --set global.domain=runai.example.com \
--wait --wait-for-jobs --timeout 20m
scripts/register_cluster.sh --domain runai.example.com --output cluster-registration.yaml
helm upgrade --install runai-cluster . -n runai --create-namespace \
-f values.cluster.yaml -f my-values.yaml -f cluster-registration.yaml \
--wait --wait-for-jobs --timeout 20m
cluster-registration.yaml holds a bootstrap token. The control plane mints a new one on every
call, so regenerate the file before each install or upgrade of runai-cluster .
scripts/install.sh --domain runai.example.com --dev -f my-values.yaml
kubectl label node < node > run.ai/simulated-gpu-node-pool=default
values.dev.yaml replaces the NVIDIA GPU Operator with the fake GPU operator, which reports
8 Tesla-A10 GPUs on each labelled node.
helm upgrade --install kai-scheduler . -n kai-scheduler --create-namespace \
-f values.kai-scheduler.yaml --wait
KAI creates the queues default-parent-queue and default-queue on first install. Pods opt in
with schedulerName: kai-scheduler and the label kai.scheduler/queue: default-queue .
Layer values.openshift.yaml on every release. Upstream detects OpenShift with lookup , which
returns nothing under helm template , so the file states it: it sets
global.config.kubernetesDistribution: openshift (Routes and SCCs in the control plane) and
forces OpenShift mode in runai-cluster , kai-resource-management and kai-scheduler . Leave
global._runai.monitoring off: the Run:ai operator uses openshift-monitoring .
ArgoCD renders with helm template , so:
Set global._runai.mode: install (or upgrade ): lookup returns nothing, and the chart then
renders the runai-reservation Namespace unconditionally.
Create one Application per release, synced in the order of the table above.
The cluster Application needs the registration values. Run scripts/register_cluster.sh and
store its output in a Secret-backed values source, since it holds a bootstrap token.
The Run:ai cluster takes its routing settings through clusterConfig , a merge patch into its
RunaiConfig . To use HAProxy instead of nginx:
cluster :
runai-cluster :
clusterConfig :
global :
ingress :
ingressClass : haproxy
Commercial charts
NVIDIA also publishes the commercial Run:ai line (2.x) as control-plane and runai-cluster in
https://helm.ngc.nvidia.com/nvidia/runai , which needs an NGC API key. Those charts take the
same install values ( controlPlane.url , controlPlane.clientSecret , cluster.uid ,
cluster.url , clusterConfig ). Their runai-cluster 2.26.50 schedules with the proprietary
runai-engine (v3.41.19, from nvcr.io/artifactory ) instead of KAI Scheduler. This chart wraps the
open-source 3.x line from nvidia/runai-oss .
Each subchart passes its upstream chart's values one level down, for example
controlplane.control-plane.keycloakx , cluster.runai-cluster.kai-resource-management or
monitoring.kube-prometheus-stack.grafana . The upstream charts also read these globals:
Key
Type
Default
Description
cluster
object
{}
Cluster subchart: the upstream values under cluster.runai-cluster, including the registration values and KAI under cluster.runai-cluster.kai-resource-management.kai-scheduler. See charts/cluster/values.yaml.
controlplane
object
{}
Control-plane subchart: TLS for global.domain, and the upstream values under controlplane.control-plane. See charts/controlplane/values.yaml.
global._runai.cluster.enabled
bool
false
Run:ai cluster, including KAI Scheduler. Release namespace must be runai. Needs a running control plane and the values scripts/register_cluster.sh writes.
global._runai.controlPlane.enabled
bool
false
Run:ai control plane. Release namespace must be runai-backend.
global._runai.customCA.certificate
string
""
PEM CA bundle rendered into the runai-ca-cert Secret of the release namespace when global.customCA.enabled is true. Leave empty to create runai-ca-cert yourself.
global._runai.gpu.enabled
bool
false
NVIDIA GPU Operator, or the fake GPU operator (gpu.fake-gpu-operator.enabled).
global._runai.kaiScheduler.enabled
bool
false
Standalone KAI Scheduler, without Run:ai. Upstream documents namespace kai-scheduler. Never together with cluster, which bundles KAI.
global._runai.mode
string
"auto"
"auto" lets templates call lookup. Set "install" or "upgrade" when rendering with helm template (ArgoCD), where lookup returns nothing.
global._runai.monitoring.enabled
bool
false
kube-prometheus-stack. The Run:ai cluster requires a Prometheus Operator.
global._runai.registry.password
string
""
Registry password or token.
global._runai.registry.server
string
""
Registry host for the control-plane pull secret runai-reg-creds. Nothing is rendered unless server, username and password are all set. The images on nvcr.io/nvidia/runai-oss are public.
global._runai.registry.username
string
""
Registry user.
global.customCA.enabled
bool
false
Trust the CA in Secret runai-ca-cert (key runai-ca.pem). Needed when the control-plane certificat
[truncated]
Helm chart that installs NVIDIA Run:ai (control plane + cluster) and KAI Scheduler
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Helm chart that installs NVIDIA Run:ai (control plane + cluster) and KAI Scheduler - logicalclocks/run-ai-helm

GitHub - logicalclocks/run-ai-helm: Helm chart that installs NVIDIA Run:ai (control plane + cluster) and KAI Scheduler · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
logicalclocks
run-ai-helm Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Helm chart that installs NVIDIA Run:ai (control plane + cluster) and KAI Scheduler
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.github .github charts charts scripts scripts templates templates .gitignore .gitignore .helmignore .helmignore CLAUDE.md CLAUDE.md Chart.lock Chart.lock Chart.yaml Chart.yaml LICENSE LICENSE README.md README.md README.md.gotmpl README.md.gotmpl helm_dependencies.sh helm_dependencies.sh values.cluster.yaml values.cluster.yaml values.control-plane.yaml values.control-plane.yaml values.dev.yaml values.dev.yaml values.kai-scheduler.yaml values.kai-scheduler.yaml values.openshift.yaml values.openshift.yaml values.prerequisites.yaml values.prerequisites.yaml values.schema.json values.schema.json values.yaml values.yaml View all files Repository files navigation
Installs NVIDIA Run:ai (control plane and cluster) and the KAI Scheduler it schedules with
The chart wraps the published upstream charts and adds the glue between them: release
profiles, validation, the Secrets the upstream charts expect to exist, and the cluster
registration step.
The chart is installed once per namespace. Each release enables its components with a profile
file:
The control plane and the cluster cannot share a release. The cluster's pre-install hooks call
the control-plane API with a cluster UUID and a bootstrap token, and the control plane issues
both only once it is running. scripts/register_cluster.sh registers the cluster and writes them
to a values file for release 3.
KAI Scheduler comes with release 3: runai-cluster depends on kai-resource-management , which
depends on kai-scheduler . The Run:ai operator writes KAI's Config and the
kai-resource-management node-pool controller writes one SchedulingShard per node pool. The
standalone kai-scheduler profile is for clusters that run KAI without Run:ai,
and the chart refuses to render it together with the cluster profile.
Kubernetes >= 1.30, Helm 3.x or 4.x, kubectl , curl , jq .
An ingress controller. The control plane creates an Ingress with class nginx by default
( global.ingress.ingressClass , nginx or haproxy ). OpenShift uses Routes.
A DNS name for the control plane ( global.domain ) that resolves to the ingress controller, both
from where scripts/install.sh runs and from inside the cluster.
A TLS certificate for global.domain : PEM in controlplane.tls , a cert-manager issuer in
controlplane.tls.certManager , or a kubernetes.io/tls Secret runai-backend-tls created in
namespace runai-backend before the install.
A default StorageClass for the PostgreSQL, Thanos (100Gi) and NATS volumes.
A Prometheus Operator and an NVIDIA GPU stack (DCGM exporter). Release 1 installs both; skip it
with --skip-prerequisites when the cluster already has them.
./helm_dependencies.sh
scripts/install.sh --domain runai.example.com -f my-values.yaml
my-values.yaml is passed to every release. A minimal one with a certificate for the domain:
controlplane :
tls :
certificate : |
-----BEGIN CERTIFICATE-----
...
key : |
-----BEGIN PRIVATE KEY-----
...
When the certificate is signed by a private CA, also set the CA so the cluster components trust
it:
global :
customCA :
enabled : true
_runai :
customCA :
certificate : |
-----BEGIN CERTIFICATE-----
...
and pass --cacert ca.pem to scripts/install.sh for the registration calls.
scripts/install.sh --help lists the options: --skip-prerequisites , --skip-control-plane ,
--skip-cluster , --dev (fake GPUs), --cluster-name , --admin-user , --cacert ,
--insecure , --timeout .
The UI is at https://<global.domain> . The admin user is test@run.ai
( controlplane.control-plane.tenantsManager.config.adminUsername ) and its password is generated:
kubectl get secret runai-backend-gen-tenants-manager -n runai-backend \
-o jsonpath= ' {.data.ADMIN_PASSWORD} ' | base64 -d
Step by step
helm upgrade --install runai-prerequisites . -n runai-prerequisites --create-namespace \
-f values.prerequisites.yaml --wait
helm upgrade --install runai-backend . -n runai-backend --create-namespace \
-f values.control-plane.yaml -f my-values.yaml --set global.domain=runai.example.com \
--wait --wait-for-jobs --timeout 20m
scripts/register_cluster.sh --domain runai.example.com --output cluster-registration.yaml
helm upgrade --install runai-cluster . -n runai --create-namespace \
-f values.cluster.yaml -f my-values.yaml -f cluster-registration.yaml \
--wait --wait-for-jobs --timeout 20m
cluster-registration.yaml holds a bootstrap token. The control plane mints a new one on every
call, so regenerate the file before each install or upgrade of runai-cluster .
scripts/install.sh --domain runai.example.com --dev -f my-values.yaml
kubectl label node < node > run.ai/simulated-gpu-node-pool=default
values.dev.yaml replaces the NVIDIA GPU Operator with the fake GPU operator, which reports
8 Tesla-A10 GPUs on each labelled node.
helm upgrade --install kai-scheduler . -n kai-scheduler --create-namespace \
-f values.kai-scheduler.yaml --wait
KAI creates the queues default-parent-queue and default-queue on first install. Pods opt in
with schedulerName: kai-scheduler and the label kai.scheduler/queue: default-queue .
Layer values.openshift.yaml on every release. Upstream detects OpenShift with lookup , which
returns nothing under helm template , so the file states it: it sets
global.config.kubernetesDistribution: openshift (Routes and SCCs in the control plane) and
forces OpenShift mode in runai-cluster , kai-resource-management and kai-scheduler . Leave
global._runai.monitoring off: the Run:ai operator uses openshift-monitoring .
ArgoCD renders with helm template , so:
Set global._runai.mode: install (or upgrade ): lookup returns nothing, and the chart then
renders the runai-reservation Namespace unconditionally.
Create one Application per release, synced in the order of the table above.
The cluster Application needs the registration values. Run scripts/register_cluster.sh and
store its output in a Secret-backed values source, since it holds a bootstrap token.
The Run:ai cluster takes its routing settings through clusterConfig , a merge patch into its
RunaiConfig . To use HAProxy instead of nginx:
cluster :
runai-cluster :
clusterConfig :
global :
ingress :
ingressClass : haproxy
Commercial charts
NVIDIA also publishes the commercial Run:ai line (2.x) as control-plane and runai-cluster in
https://helm.ngc.nvidia.com/nvidia/runai , which needs an NGC API key. Those charts take the
same install values ( controlPlane.url , controlPlane.clientSecret , cluster.uid ,
cluster.url , clusterConfig ). Their runai-cluster 2.26.50 schedules with the proprietary
runai-engine (v3.41.19, from nvcr.io/artifactory ) instead of KAI Scheduler. This chart wraps the
open-source 3.x line from nvidia/runai-oss .
Each subchart passes its upstream chart's values one level down, for example
controlplane.control-plane.keycloakx , cluster.runai-cluster.kai-resource-management or
monitoring.kube-prometheus-stack.grafana . The upstream charts also read these globals:
Key
Type
Default
Description
cluster
object
{}
Cluster subchart: the upstream values under cluster.runai-cluster, including the registration values and KAI under cluster.runai-cluster.kai-resource-management.kai-scheduler. See charts/cluster/values.yaml.
controlplane
object
{}
Control-plane subchart: TLS for global.domain, and the upstream values under controlplane.control-plane. See charts/controlplane/values.yaml.
global._runai.cluster.enabled
bool
false
Run:ai cluster, including KAI Scheduler. Release namespace must be runai. Needs a running control plane and the values scripts/register_cluster.sh writes.
global._runai.controlPlane.enabled
bool
false
Run:ai control plane. Release namespace must be runai-backend.
global._runai.customCA.certificate
string
""
PEM CA bundle rendered into the runai-ca-cert Secret of the release namespace when global.customCA.enabled is true. Leave empty to create runai-ca-cert yourself.
global._runai.gpu.enabled
bool
false
NVIDIA GPU Operator, or the fake GPU operator (gpu.fake-gpu-operator.enabled).
global._runai.kaiScheduler.enabled
bool
false
Standalone KAI Scheduler, without Run:ai. Upstream documents namespace kai-scheduler. Never together with cluster, which bundles KAI.
global._runai.mode
string
"auto"
"auto" lets templates call lookup. Set "install" or "upgrade" when rendering with helm template (ArgoCD), where lookup returns nothing.
global._runai.monitoring.enabled
bool
false
kube-prometheus-stack. The Run:ai cluster requires a Prometheus Operator.
global._runai.registry.password
string
""
Registry password or token.
global._runai.registry.server
string
""
Registry host for the control-plane pull secret runai-reg-creds. Nothing is rendered unless server, username and password are all set. The images on nvcr.io/nvidia/runai-oss are public.
global._runai.registry.username
string
""
Registry user.
global.customCA.enabled
bool
false
Trust the CA in Secret runai-ca-cert (key runai-ca.pem). Needed when the control-plane certificat
[truncated]
Helm chart that installs NVIDIA Run:ai (control plane + cluster) and KAI Scheduler
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
