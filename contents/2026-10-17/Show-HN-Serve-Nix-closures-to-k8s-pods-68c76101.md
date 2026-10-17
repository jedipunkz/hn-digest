---
source: "https://github.com/Kranzes/nix-store-csi"
hn_url: "https://news.ycombinator.com/item?id=50136139"
title: "Show HN: Serve Nix closures to k8s pods"
article_title: "GitHub - Kranzes/nix-store-csi: Serves Nix store closures from binary caches to Kubernetes pods as CSI volumes · GitHub"
image: "https://opengraph.githubassets.com/d0354290ac07e9542d759939ce41612863fd34588702c2c63457b1c7d3865af8/Kranzes/nix-store-csi"
author: "kranzes"
captured_at: "2026-10-17T23:47:42Z"
capture_tool: "hn-digest"
hn_id: 50136139
score: 1
comments: 0
posted_at: "2026-10-17T23:24:52Z"
tags:
  - hacker-news
---

# Show HN: Serve Nix closures to k8s pods

- HN: [50136139](https://news.ycombinator.com/item?id=50136139)
- Source: [github.com](https://github.com/Kranzes/nix-store-csi)
- Score: 1
- Comments: 0
- Posted: 2026-10-17T23:24:52Z

## Translation

Title: Show HN: Serve Nix closures to k8s pods
Article title: GitHub - Kranzes/nix-store-csi: Serves Nix store closures from binary caches to Kubernetes pods as CSI volumes · GitHub
Description: Serves Nix store closures from binary caches to Kubernetes pods as CSI volumes - Kranzes/nix-store-csi
HN text: I’ve been working on a Nix CSI for kubernetes since NixCon when I realized that everyone seems to working towards killing OCI layers (and images) in Kubernetes deployments. I’m hoping to get people to give it a spin and report back to me bugs, feature requests or general feedback. I don’t have a lot of data on how this scales and that makes it hard to decide whether to justify adding a feature or not (e.g. cache locality scheduling). An LLM was used in the development of this project.

Article text:
GitHub - Kranzes/nix-store-csi: Serves Nix store closures from binary caches to Kubernetes pods as CSI volumes · GitHub
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
Kranzes
nix-store-csi Public Notifications You must be signed in to change notification settings
Star 11 ( 11 ) You must be signed in to star a repository
Serves Nix store closures from binary caches to Kubernetes pods as CSI volumes
Readme MIT license Activity Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
57 Commits 57 Commits Folders and files
.github .github deploy deploy docs docs nix nix rust rust .envrc .envrc .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md flake.lock flake.lock flake.nix flake.nix View all files Repository files navigation
nix-store-csi mounts Nix store paths in Kubernetes pods. A pod lists store
paths in an inline volume. The plugin fetches their closure from binary caches
into a store shared by the node, then mounts a view holding only that closure
at the pod's /nix/store . Pods can run in the runner image, which holds tini,
or in any other image.
nix-store-csi unpack fills a store the way the plugin does, so you can try
it with podman:
exe= $( nix eval --raw nixpkgs#pkgs --apply ' pkgs: pkgs.lib.getExe pkgs.hello ' )
store= $( nix run github:Kranzes/nix-store-csi/v0.1.6 -- --state-dir /tmp/nix-store-csi unpack " $exe " )
podman run --rm -v " $store " :/nix/store:ro ghcr.io/kranzes/nix-store-csi-runner:0.1.6 " $exe "
Why a CSI driver
nix-snapshotter also runs
Nix store paths in Kubernetes, as a containerd snapshotter that mounts them as
image layers. nix-store-csi mounts them as a volume instead:
No images or registry. A new version is a new store path in the pod
spec, and pods run in a fixed image like the runner image. nix-snapshotter
needs a registry, or kubelet pointed at its image service and Nix on the
node to resolve images by store path.
Its own storage model. Lazy and chunked snapshotters exist, like
stargz and Nydus, but they need their own image formats served from a
registry. Here the plugin picks the store's format alone, so it could later
be lazy via FUSE, or a castore with FastCDC chunks in place of NARs.
Nothing to install on nodes. The plugin is a DaemonSet that fetches NARs
itself, so nodes need no containerd plugin and no Nix. That suits managed
node pools, as long as their kernel meets the
requirements . nix-snapshotter runs on the host
next to containerd and uses the host's Nix store.
Any CRI runtime. Kubelet mounts the volume, so the plugin doesn't depend
on the runtime. nix-snapshotter needs containerd.
Image scanners, signature checks and registry allowlists don't see
storePaths . Scan an SBOM of the closure instead, e.g. from
bombon . The plugin fetches only from
substituters and checks narinfo signatures against trustedPublicKeys .
Any pod that may use the driver can make a node fetch whatever the trusted
caches serve. maxClosureSize and admissionPolicy.allowedNamespaces can
limit that, see Security .
Serves Nix store closures from binary caches to Kubernetes pods as CSI volumes
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Serves Nix store closures from binary caches to Kubernetes pods as CSI volumes - Kranzes/nix-store-csi

I’ve been working on a Nix CSI for kubernetes since NixCon when I realized that everyone seems to working towards killing OCI layers (and images) in Kubernetes deployments. I’m hoping to get people to give it a spin and report back to me bugs, feature requests or general feedback. I don’t have a lot of data on how this scales and that makes it hard to decide whether to justify adding a feature or not (e.g. cache locality scheduling). An LLM was used in the development of this project.

GitHub - Kranzes/nix-store-csi: Serves Nix store closures from binary caches to Kubernetes pods as CSI volumes · GitHub
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
Kranzes
nix-store-csi Public Notifications You must be signed in to change notification settings
Star 11 ( 11 ) You must be signed in to star a repository
Serves Nix store closures from binary caches to Kubernetes pods as CSI volumes
Readme MIT license Activity Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
57 Commits 57 Commits Folders and files
.github .github deploy deploy docs docs nix nix rust rust .envrc .envrc .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md flake.lock flake.lock flake.nix flake.nix View all files Repository files navigation
nix-store-csi mounts Nix store paths in Kubernetes pods. A pod lists store
paths in an inline volume. The plugin fetches their closure from binary caches
into a store shared by the node, then mounts a view holding only that closure
at the pod's /nix/store . Pods can run in the runner image, which holds tini,
or in any other image.
nix-store-csi unpack fills a store the way the plugin does, so you can try
it with podman:
exe= $( nix eval --raw nixpkgs#pkgs --apply ' pkgs: pkgs.lib.getExe pkgs.hello ' )
store= $( nix run github:Kranzes/nix-store-csi/v0.1.6 -- --state-dir /tmp/nix-store-csi unpack " $exe " )
podman run --rm -v " $store " :/nix/store:ro ghcr.io/kranzes/nix-store-csi-runner:0.1.6 " $exe "
Why a CSI driver
nix-snapshotter also runs
Nix store paths in Kubernetes, as a containerd snapshotter that mounts them as
image layers. nix-store-csi mounts them as a volume instead:
No images or registry. A new version is a new store path in the pod
spec, and pods run in a fixed image like the runner image. nix-snapshotter
needs a registry, or kubelet pointed at its image service and Nix on the
node to resolve images by store path.
Its own storage model. Lazy and chunked snapshotters exist, like
stargz and Nydus, but they need their own image formats served from a
registry. Here the plugin picks the store's format alone, so it could later
be lazy via FUSE, or a castore with FastCDC chunks in place of NARs.
Nothing to install on nodes. The plugin is a DaemonSet that fetches NARs
itself, so nodes need no containerd plugin and no Nix. That suits managed
node pools, as long as their kernel meets the
requirements . nix-snapshotter runs on the host
next to containerd and uses the host's Nix store.
Any CRI runtime. Kubelet mounts the volume, so the plugin doesn't depend
on the runtime. nix-snapshotter needs containerd.
Image scanners, signature checks and registry allowlists don't see
storePaths . Scan an SBOM of the closure instead, e.g. from
bombon . The plugin fetches only from
substituters and checks narinfo signatures against trustedPublicKeys .
Any pod that may use the driver can make a node fetch whatever the trusted
caches serve. maxClosureSize and admissionPolicy.allowedNamespaces can
limit that, see Security .
Serves Nix store closures from binary caches to Kubernetes pods as CSI volumes
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
