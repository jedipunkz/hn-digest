---
source: "https://github.com/arterialist/kube32"
hn_url: "https://news.ycombinator.com/item?id=50291812"
title: "I rebuilt Kubernetes for ESP32 microcontrollers"
article_title: "GitHub - arterialist/kube32: Kubernetes-compatible control plane, Kubelet agent, and WebAssembly runtime for ESP32 microcontrollers · GitHub"
image: "https://opengraph.githubassets.com/357df73f89a2f3323b38f1cf1d92a68b61470d2e4c0e12003581c7f8dd90bdaa/arterialist/kube32"
author: "Arterialist"
captured_at: "2026-10-29T16:39:25Z"
capture_tool: "hn-digest"
hn_id: 50291812
score: 2
comments: 1
posted_at: "2026-10-29T16:35:14Z"
tags:
  - hacker-news
---

# I rebuilt Kubernetes for ESP32 microcontrollers

- HN: [50291812](https://news.ycombinator.com/item?id=50291812)
- Source: [github.com](https://github.com/arterialist/kube32)
- Score: 2
- Comments: 1
- Posted: 2026-10-29T16:35:14Z

## Translation

Title: I rebuilt Kubernetes for ESP32 microcontrollers
Article title: GitHub - arterialist/kube32: Kubernetes-compatible control plane, Kubelet agent, and WebAssembly runtime for ESP32 microcontrollers · GitHub
Description: Kubernetes-compatible control plane, Kubelet agent, and WebAssembly runtime for ESP32 microcontrollers - arterialist/kube32

Article text:
GitHub - arterialist/kube32: Kubernetes-compatible control plane, Kubelet agent, and WebAssembly runtime for ESP32 microcontrollers · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
arterialist
kube32 Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Kubernetes-compatible control plane, Kubelet agent, and WebAssembly runtime for ESP32 microcontrollers
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
115 Commits 115 Commits Folders and files
.github .github apps apps cmake cmake docs docs examples examples federation federation images images include include scripts scripts sdk/ rust sdk/ rust src src tests tests third_party/ wasm3 third_party/ wasm3 tools tools .gitignore .gitignore CHANGELOG.md CHANGELOG.md CMakeLists.txt CMakeLists.txt CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml partitions.csv partitions.csv platformio.ini platformio.ini sdkconfig.defaults sdkconfig.defaults sdkconfig.esp32s3 sdkconfig.esp32s3 View all files Repository files navigation
kube32 is a Kubernetes-compatible control plane and WebAssembly pod runtime for ESP32-S3 microcontrollers.
kubectl and k9s manage a cluster of boards over Wi-Fi or USB serial. Pods are WebAssembly modules executed by the wasm3 interpreter. One firmware image runs on every board. A board with PSRAM acts as the control plane and runs a local kubelet, while boards without PSRAM join as workers. The same sources build kube32-sim , a host simulator for macOS and Linux.
Kubernetes API serving Pod, Node, Namespace, Service, Endpoints, EndpointSlice, Ingress, Deployment, DaemonSet, Job, ConfigMap, Secret, Lease, Event, and metrics for kubectl and k9s, including watch streams, Table output, OpenAPI v3 validation, and log tailing.
WebAssembly pods in wasm3 with SDKs and templates for Rust, C, and AssemblyScript, using either a request-handler ABI with k32_handle or an entry-point ABI with app_main , plus configurable memory limits and per-call gas budgets.
Service routing, HTTP Ingress with host and prefix rules, Server-Sent Events streams, embedded static asset bundles via k32_resp_asset and StaticBundle , and synchronous pod-to-pod HTTP calls via k32_http_call .
Persistent object store backed by a CRC-checked journal on a raw flash partition or the microSD card, plus persistentVolumeClaim pod volumes via k32_fs_* and rotated container logs for kubectl logs --previous .
Peripheral access gated per pod by a capability firewall covering GPIO pins, ADC channels, I2C and UART buses, ST7789 display, key matrix, I2S speaker, and microSD storage.
Container lifecycle management with restartPolicy set to Always , OnFailure , or Never , exponential CrashLoopBackOff, and distinct exit codes for Wasm traps, gas exhaustion, and OOMKilled .
Optional bearer-token authentication and automatic self-signed ECDSA P-256 TLS on port 6443.
Pub/sub events via k32_event_emit and kube32.io/event-topics , bridged across nodes over ESP-NOW and UDP port 7400.
Multi-node host simulator kube32-sim and kube32-cli for building, checking, and deploying pods over the network or to node storage.
Role
Requirements
Reference board
Control plane
ESP32-S3, 8 MB flash, 2 MB or more quad SPI PSRAM
Generic ESP32-S3 devkit with 2 MB PSRAM
Worker
Any ESP32-S3 with 8 MB flash
M5Stack Cardputer with ST7789 display, keypad, speaker, and microSD
See docs/hardware-spec.md for pin assignments and memory budgets.
Build and start a two-node cluster on macOS or Linux:
cmake -S . -B build && cmake --build build -j8
./build/kube32-sim --ingress-port 8080 --pod-port 17000
In a second terminal, configure kubectl and deploy the echo example:
export KUBECONFIG= $PWD /build/kubeconfig
kubectl config set-cluster kube32-sim --server=http://127.0.0.1:6443
kubectl config set-context kube32-sim --cluster=kube32-sim
kubectl config use-context kube32-sim
kubectl get nodes
kubectl apply -f examples/echo.yaml
Boards
Build and flash the firmware to each board with PlatformIO:
pio run -e esp32s3
bash scripts/flash.sh < serial-port > 921600
Open the serial console at 115200 baud and provision Wi-Fi and the board profile:
K32_CMD:BOARD:cardputer # Cardputer worker only
K32_CMD:WIFI:<ssid>:<password> # both boards
K32_CMD:AP:off # control plane
K32_CMD:API:<control-plane-ip> # worker
Point kubectl at the control plane IP and apply a manifest:
kubectl config set-cluster kube32 --server=http:// < control-plane-ip > :6443
kubectl config set-context kube32 --cluster=kube32
kubectl config use-context kube32
kubectl get nodes
kubectl apply -f examples/echo.yaml
See docs/quickstart.md for full setup instructions and TLS configuration.
Tessera is an issue tracker made of four WebAssembly pods routed by an Ingress. tessera-web and tessera-api run on the control plane to serve the web UI and in-memory search index. tessera-db and tessera-hud run on the Cardputer worker to persist records to an append-only log on the microSD card and drive the LCD and keypad ( tessera-hud is silent by default; sound is opt-in via HUD_SOUND=1 and the speaker capability).
Generate the sample dataset, upload it to the worker's microSD card, and deploy:
make -C apps/tessera dataset
tools/kube32-cli push apps/tessera/build/card/kube32/volumes/tessera-db --node < worker-node > --path volumes/tessera-db
kubectl apply -f examples/tessera.yaml
Open http://<control-plane-ip>/ or http://127.0.0.1:8080/ on the simulator, and sign in as mara with password demo .
Full documentation is published at https://arterialist.github.io/kube32/ .
Flashing and the serial console
Cluster access and authentication
Services, Ingress and pod-to-pod calls
Contributions are welcome. Read CONTRIBUTING.md for build instructions, test suites, and coding guidelines, or open an issue on GitHub .
See SECURITY.md for the deployment security model and instructions on reporting vulnerabilities privately.
Apache License 2.0. See LICENSE . third_party/wasm3 is licensed under the MIT License in third_party/wasm3/LICENSE .
Kubernetes-compatible control plane, Kubelet agent, and WebAssembly runtime for ESP32 microcontrollers
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Kubernetes-compatible control plane, Kubelet agent, and WebAssembly runtime for ESP32 microcontrollers - arterialist/kube32

GitHub - arterialist/kube32: Kubernetes-compatible control plane, Kubelet agent, and WebAssembly runtime for ESP32 microcontrollers · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
arterialist
kube32 Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Kubernetes-compatible control plane, Kubelet agent, and WebAssembly runtime for ESP32 microcontrollers
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
115 Commits 115 Commits Folders and files
.github .github apps apps cmake cmake docs docs examples examples federation federation images images include include scripts scripts sdk/ rust sdk/ rust src src tests tests third_party/ wasm3 third_party/ wasm3 tools tools .gitignore .gitignore CHANGELOG.md CHANGELOG.md CMakeLists.txt CMakeLists.txt CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml partitions.csv partitions.csv platformio.ini platformio.ini sdkconfig.defaults sdkconfig.defaults sdkconfig.esp32s3 sdkconfig.esp32s3 View all files Repository files navigation
kube32 is a Kubernetes-compatible control plane and WebAssembly pod runtime for ESP32-S3 microcontrollers.
kubectl and k9s manage a cluster of boards over Wi-Fi or USB serial. Pods are WebAssembly modules executed by the wasm3 interpreter. One firmware image runs on every board. A board with PSRAM acts as the control plane and runs a local kubelet, while boards without PSRAM join as workers. The same sources build kube32-sim , a host simulator for macOS and Linux.
Kubernetes API serving Pod, Node, Namespace, Service, Endpoints, EndpointSlice, Ingress, Deployment, DaemonSet, Job, ConfigMap, Secret, Lease, Event, and metrics for kubectl and k9s, including watch streams, Table output, OpenAPI v3 validation, and log tailing.
WebAssembly pods in wasm3 with SDKs and templates for Rust, C, and AssemblyScript, using either a request-handler ABI with k32_handle or an entry-point ABI with app_main , plus configurable memory limits and per-call gas budgets.
Service routing, HTTP Ingress with host and prefix rules, Server-Sent Events streams, embedded static asset bundles via k32_resp_asset and StaticBundle , and synchronous pod-to-pod HTTP calls via k32_http_call .
Persistent object store backed by a CRC-checked journal on a raw flash partition or the microSD card, plus persistentVolumeClaim pod volumes via k32_fs_* and rotated container logs for kubectl logs --previous .
Peripheral access gated per pod by a capability firewall covering GPIO pins, ADC channels, I2C and UART buses, ST7789 display, key matrix, I2S speaker, and microSD storage.
Container lifecycle management with restartPolicy set to Always , OnFailure , or Never , exponential CrashLoopBackOff, and distinct exit codes for Wasm traps, gas exhaustion, and OOMKilled .
Optional bearer-token authentication and automatic self-signed ECDSA P-256 TLS on port 6443.
Pub/sub events via k32_event_emit and kube32.io/event-topics , bridged across nodes over ESP-NOW and UDP port 7400.
Multi-node host simulator kube32-sim and kube32-cli for building, checking, and deploying pods over the network or to node storage.
Role
Requirements
Reference board
Control plane
ESP32-S3, 8 MB flash, 2 MB or more quad SPI PSRAM
Generic ESP32-S3 devkit with 2 MB PSRAM
Worker
Any ESP32-S3 with 8 MB flash
M5Stack Cardputer with ST7789 display, keypad, speaker, and microSD
See docs/hardware-spec.md for pin assignments and memory budgets.
Build and start a two-node cluster on macOS or Linux:
cmake -S . -B build && cmake --build build -j8
./build/kube32-sim --ingress-port 8080 --pod-port 17000
In a second terminal, configure kubectl and deploy the echo example:
export KUBECONFIG= $PWD /build/kubeconfig
kubectl config set-cluster kube32-sim --server=http://127.0.0.1:6443
kubectl config set-context kube32-sim --cluster=kube32-sim
kubectl config use-context kube32-sim
kubectl get nodes
kubectl apply -f examples/echo.yaml
Boards
Build and flash the firmware to each board with PlatformIO:
pio run -e esp32s3
bash scripts/flash.sh < serial-port > 921600
Open the serial console at 115200 baud and provision Wi-Fi and the board profile:
K32_CMD:BOARD:cardputer # Cardputer worker only
K32_CMD:WIFI:<ssid>:<password> # both boards
K32_CMD:AP:off # control plane
K32_CMD:API:<control-plane-ip> # worker
Point kubectl at the control plane IP and apply a manifest:
kubectl config set-cluster kube32 --server=http:// < control-plane-ip > :6443
kubectl config set-context kube32 --cluster=kube32
kubectl config use-context kube32
kubectl get nodes
kubectl apply -f examples/echo.yaml
See docs/quickstart.md for full setup instructions and TLS configuration.
Tessera is an issue tracker made of four WebAssembly pods routed by an Ingress. tessera-web and tessera-api run on the control plane to serve the web UI and in-memory search index. tessera-db and tessera-hud run on the Cardputer worker to persist records to an append-only log on the microSD card and drive the LCD and keypad ( tessera-hud is silent by default; sound is opt-in via HUD_SOUND=1 and the speaker capability).
Generate the sample dataset, upload it to the worker's microSD card, and deploy:
make -C apps/tessera dataset
tools/kube32-cli push apps/tessera/build/card/kube32/volumes/tessera-db --node < worker-node > --path volumes/tessera-db
kubectl apply -f examples/tessera.yaml
Open http://<control-plane-ip>/ or http://127.0.0.1:8080/ on the simulator, and sign in as mara with password demo .
Full documentation is published at https://arterialist.github.io/kube32/ .
Flashing and the serial console
Cluster access and authentication
Services, Ingress and pod-to-pod calls
Contributions are welcome. Read CONTRIBUTING.md for build instructions, test suites, and coding guidelines, or open an issue on GitHub .
See SECURITY.md for the deployment security model and instructions on reporting vulnerabilities privately.
Apache License 2.0. See LICENSE . third_party/wasm3 is licensed under the MIT License in third_party/wasm3/LICENSE .
Kubernetes-compatible control plane, Kubelet agent, and WebAssembly runtime for ESP32 microcontrollers
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
