---
source: "https://github.com/kobblestoneio/kobblestone"
hn_url: "https://news.ycombinator.com/item?id=49765874"
title: "Show HN: Kobblestone – Minecraft on Kubernetes"
article_title: "GitHub - kobblestoneio/kobblestone: Minecraft Server Infrastructure on Kubernetes · GitHub"
image: "https://opengraph.githubassets.com/05680bf89ab14a23e684bfe3a63527bb7ac241b15693d2f92e21554bfb803dfc/kobblestoneio/kobblestone"
author: "devmojo"
captured_at: "2026-09-19T12:55:20Z"
capture_tool: "hn-digest"
hn_id: 49765874
score: 2
comments: 0
posted_at: "2026-09-19T12:07:15Z"
tags:
  - hacker-news
---

# Show HN: Kobblestone – Minecraft on Kubernetes

- HN: [49765874](https://news.ycombinator.com/item?id=49765874)
- Source: [github.com](https://github.com/kobblestoneio/kobblestone)
- Score: 2
- Comments: 0
- Posted: 2026-09-19T12:07:15Z

## Translation

Title: Show HN: Kobblestone – Minecraft on Kubernetes
Article title: GitHub - kobblestoneio/kobblestone: Minecraft Server Infrastructure on Kubernetes · GitHub
Description: Minecraft Server Infrastructure on Kubernetes. Contribute to kobblestoneio/kobblestone development by creating an account on GitHub.
HN text: Hi folks, just released a personal project of mine, called Kobblestone. Essentially, it's an operator for Kubernetes, allowing you to manage Minecraft infrastructure on Kubernetes. It's not just a small wrapper for deploying servers, it aims to integrate the whole Minecraft ecosystem on Kubernetes (routers, networks, commands, backups, etc.). I'd highly appreciate your feedback and I am looking forward to potential contributions! Thanks.

Article text:
GitHub - kobblestoneio/kobblestone: Minecraft Server Infrastructure on Kubernetes · GitHub
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
kobblestoneio
/
kobblestone
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
examples examples kustomize kustomize src src .dockerignore .dockerignore .editorconfig .editorconfig .gitignore .gitignore Directory.Build.props Directory.Build.props Directory.Packages.props Directory.Packages.props Dockerfile Dockerfile LICENSE LICENSE README.md README.md dotnet-tools.json dotnet-tools.json global.json global.json kobblestone.yaml kobblestone.yaml View all files Repository files navigation
Kobblestone allows you to manage Minecraft server infrastructure on Kubernetes through custom resource types, turning
Minecraft into a first-class citizen of Kubernetes.
Supports Minecraft servers of many different types and versions (Vanilla, PaperMC, Purpur, Fabric, etc.).
Hostname based routing across multiple servers.
Combine multiple servers to build a network behind a Velocity proxy.
Live and offline on-demand server backups with support for multiple storage backends (S3, PVC).
Native Gateway API integration for exposing
servers, networks and routers (also supports LoadBalancer Services).
Execute server commands declaratively via Kubernetes resources.
Automatic server scale-down/-up based on player inactivity.
Bot programming and deployment based on Mineflayer .
Manage Microsoft accounts and device-flow authorization for authenticated bots.
Server : Manages a Minecraft server instance.
Router : Defines a router that is able to route Minecraft traffic based on hostnames.
Route : Defines a routable hostname and target.
Network : Manages and configures a Velocity proxy instance.
BackupRepo : Defines a storage backend for backups. Supports S3 buckets and PVC.
Backup : Creates a backup of a server and stores it inside the referenced repo. Supports live backups coordinated via
RCON.
Restore : Restores a backup on a server.
Command : Execute a command on a server via RCON and capture the response.
CommandGroup : Creates and executes multiple commands in sequential order.
Account : Manages the authorization state for a Microsoft account.
BotBehavior : Defines the behavior/logic of a bot.
Bot : Manages a Minecraft bot instance.
BotGroup : Manages multiple instances of the same bot against different target servers.
Kobblestone can be installed via a simple kubectl apply command:
kubectl apply -f https://get.kobblestone.io/v0.1.0/kobblestone.yaml
Deploying a simple Vanilla server
Define your server in a server.yaml file...
apiVersion : kobblestone.io/v1alpha1
kind : Server
metadata :
name : my-server
spec :
eula : true
type : Vanilla
version : " 26.2 "
storage :
size : 1Gi
resources :
limits :
memory : 2Gi
jvm :
memory :
# You may want to decrease this for lower memory limits
heapPercentage : 75
... and apply.
kubectl apply -f server.yaml \
&& kubectl wait --for=condition=Running=True server/my-server --timeout=5m
The command will wait for the server to be up and running. This may take some time if launched for the first time.
Exposing your server outside the cluster
By default, servers are only available inside the cluster. You have multiple options for exposing your server, which are
more or less applicable depending on your specific environment.
Option A: LoadBalancer Service (not recommended)
The simplest and most straight forward option is to make the servers Service of type LoadBalancer . If supported by
your environment, it will give your server an external IP where it is reachable.
Warning : Beware that if you are in the cloud, this will probably allocate resources that cost money. It is also
recommended to disable RCON if the external IP is routable directly from the internet. Even though RCON is
password-protected, it is still unencrypted.
apiVersion : kobblestone.io/v1alpha1
kind : Server
metadata :
name : my-server
spec :
# ...
service :
type : LoadBalancer
annotations : { } # Configure environment specifics of LB here (like MetalLB ip-pool)
rcon :
disabled : true
Option B: Gateway API
If you have Gateway API >= v1.6.0 installed on your
cluster, this might be the most optimal option for you.
You can easily attach servers (and other resources) to your gateways. Kobblestone will manage the TCPRoute for you.
apiVersion : kobblestone.io/v1alpha1
kind : Server
metadata :
name : my-server
spec :
# ...
tcpRoute :
parentRefs :
- name : my-gateway
sectionName : minecraft # TCP listener on your gateway
Your server will then be accessible through the gateways TCP listener.
You can also expose your server indirectly through a central Router instance. The advantage here is that we only have
to expose one component (the router) and all the servers stay cluster internal. In addition to that, we also have
hostname based routing.
For that, we need to deploy and expose a Router first.
apiVersion : kobblestone.io/v1alpha1
kind : Router
metadata :
name : my-router
spec :
replicas : 1 # You can scale routers horizontally
# You can use a LoadBalancer Service...
service :
type : LoadBalancer
# ... or the Gateway API
tcpRoute :
parentRefs :
- name : my-gateway
sectionName : minecraft # TCP listener on your gateway
Secondly, we can configure the default route on the server.
apiVersion : kobblestone.io/v1alpha1
kind : Server
metadata :
name : my-server
spec :
# ...
route :
parentRef :
name : my-router
hostname : server.example.com # Some hostname that resolves to the routers external endpoint
Special Thanks
itzg for
creating docker-minecraft-server , mc-router , docker-mc-backup
and docker-mc-proxy . These awesome projects provide core functionality for
Kobblestone. Check them out!
PrismarineJs for creating mineflayer .
Kobblestone uses it to implement bots.
Minecraft Server Infrastructure on Kubernetes
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Minecraft Server Infrastructure on Kubernetes. Contribute to kobblestoneio/kobblestone development by creating an account on GitHub.

Hi folks, just released a personal project of mine, called Kobblestone. Essentially, it's an operator for Kubernetes, allowing you to manage Minecraft infrastructure on Kubernetes. It's not just a small wrapper for deploying servers, it aims to integrate the whole Minecraft ecosystem on Kubernetes (routers, networks, commands, backups, etc.). I'd highly appreciate your feedback and I am looking forward to potential contributions! Thanks.

GitHub - kobblestoneio/kobblestone: Minecraft Server Infrastructure on Kubernetes · GitHub
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
kobblestoneio
/
kobblestone
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
examples examples kustomize kustomize src src .dockerignore .dockerignore .editorconfig .editorconfig .gitignore .gitignore Directory.Build.props Directory.Build.props Directory.Packages.props Directory.Packages.props Dockerfile Dockerfile LICENSE LICENSE README.md README.md dotnet-tools.json dotnet-tools.json global.json global.json kobblestone.yaml kobblestone.yaml View all files Repository files navigation
Kobblestone allows you to manage Minecraft server infrastructure on Kubernetes through custom resource types, turning
Minecraft into a first-class citizen of Kubernetes.
Supports Minecraft servers of many different types and versions (Vanilla, PaperMC, Purpur, Fabric, etc.).
Hostname based routing across multiple servers.
Combine multiple servers to build a network behind a Velocity proxy.
Live and offline on-demand server backups with support for multiple storage backends (S3, PVC).
Native Gateway API integration for exposing
servers, networks and routers (also supports LoadBalancer Services).
Execute server commands declaratively via Kubernetes resources.
Automatic server scale-down/-up based on player inactivity.
Bot programming and deployment based on Mineflayer .
Manage Microsoft accounts and device-flow authorization for authenticated bots.
Server : Manages a Minecraft server instance.
Router : Defines a router that is able to route Minecraft traffic based on hostnames.
Route : Defines a routable hostname and target.
Network : Manages and configures a Velocity proxy instance.
BackupRepo : Defines a storage backend for backups. Supports S3 buckets and PVC.
Backup : Creates a backup of a server and stores it inside the referenced repo. Supports live backups coordinated via
RCON.
Restore : Restores a backup on a server.
Command : Execute a command on a server via RCON and capture the response.
CommandGroup : Creates and executes multiple commands in sequential order.
Account : Manages the authorization state for a Microsoft account.
BotBehavior : Defines the behavior/logic of a bot.
Bot : Manages a Minecraft bot instance.
BotGroup : Manages multiple instances of the same bot against different target servers.
Kobblestone can be installed via a simple kubectl apply command:
kubectl apply -f https://get.kobblestone.io/v0.1.0/kobblestone.yaml
Deploying a simple Vanilla server
Define your server in a server.yaml file...
apiVersion : kobblestone.io/v1alpha1
kind : Server
metadata :
name : my-server
spec :
eula : true
type : Vanilla
version : " 26.2 "
storage :
size : 1Gi
resources :
limits :
memory : 2Gi
jvm :
memory :
# You may want to decrease this for lower memory limits
heapPercentage : 75
... and apply.
kubectl apply -f server.yaml \
&& kubectl wait --for=condition=Running=True server/my-server --timeout=5m
The command will wait for the server to be up and running. This may take some time if launched for the first time.
Exposing your server outside the cluster
By default, servers are only available inside the cluster. You have multiple options for exposing your server, which are
more or less applicable depending on your specific environment.
Option A: LoadBalancer Service (not recommended)
The simplest and most straight forward option is to make the servers Service of type LoadBalancer . If supported by
your environment, it will give your server an external IP where it is reachable.
Warning : Beware that if you are in the cloud, this will probably allocate resources that cost money. It is also
recommended to disable RCON if the external IP is routable directly from the internet. Even though RCON is
password-protected, it is still unencrypted.
apiVersion : kobblestone.io/v1alpha1
kind : Server
metadata :
name : my-server
spec :
# ...
service :
type : LoadBalancer
annotations : { } # Configure environment specifics of LB here (like MetalLB ip-pool)
rcon :
disabled : true
Option B: Gateway API
If you have Gateway API >= v1.6.0 installed on your
cluster, this might be the most optimal option for you.
You can easily attach servers (and other resources) to your gateways. Kobblestone will manage the TCPRoute for you.
apiVersion : kobblestone.io/v1alpha1
kind : Server
metadata :
name : my-server
spec :
# ...
tcpRoute :
parentRefs :
- name : my-gateway
sectionName : minecraft # TCP listener on your gateway
Your server will then be accessible through the gateways TCP listener.
You can also expose your server indirectly through a central Router instance. The advantage here is that we only have
to expose one component (the router) and all the servers stay cluster internal. In addition to that, we also have
hostname based routing.
For that, we need to deploy and expose a Router first.
apiVersion : kobblestone.io/v1alpha1
kind : Router
metadata :
name : my-router
spec :
replicas : 1 # You can scale routers horizontally
# You can use a LoadBalancer Service...
service :
type : LoadBalancer
# ... or the Gateway API
tcpRoute :
parentRefs :
- name : my-gateway
sectionName : minecraft # TCP listener on your gateway
Secondly, we can configure the default route on the server.
apiVersion : kobblestone.io/v1alpha1
kind : Server
metadata :
name : my-server
spec :
# ...
route :
parentRef :
name : my-router
hostname : server.example.com # Some hostname that resolves to the routers external endpoint
Special Thanks
itzg for
creating docker-minecraft-server , mc-router , docker-mc-backup
and docker-mc-proxy . These awesome projects provide core functionality for
Kobblestone. Check them out!
PrismarineJs for creating mineflayer .
Kobblestone uses it to implement bots.
Minecraft Server Infrastructure on Kubernetes
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
