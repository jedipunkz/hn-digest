---
source: "https://github.com/skidvis/lcars-k8s"
hn_url: "https://news.ycombinator.com/item?id=49650108"
title: "Lcars-k8 – Star Trek LCARS inspired Kubernetes dashboard"
article_title: "GitHub - skidvis/lcars-k8s · GitHub"
image: "https://opengraph.githubassets.com/cfa3064ee724495809576f17e29ecfdbb152701690dfc160f3985d206afb1db5/skidvis/lcars-k8s"
author: "SkidVis"
captured_at: "2026-09-10T21:27:33Z"
capture_tool: "hn-digest"
hn_id: 49650108
score: 1
comments: 0
posted_at: "2026-09-10T20:59:49Z"
tags:
  - hacker-news
---

# Lcars-k8 – Star Trek LCARS inspired Kubernetes dashboard

- HN: [49650108](https://news.ycombinator.com/item?id=49650108)
- Source: [github.com](https://github.com/skidvis/lcars-k8s)
- Score: 1
- Comments: 0
- Posted: 2026-09-10T20:59:49Z

## Translation

Title: Lcars-k8 – Star Trek LCARS inspired Kubernetes dashboard
Article title: GitHub - skidvis/lcars-k8s · GitHub
Description: Contribute to skidvis/lcars-k8s development by creating an account on GitHub.

Article text:
GitHub - skidvis/lcars-k8s · GitHub
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
skidvis
/
lcars-k8s
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
9 Commits 9 Commits Folders and files
dev dev lcarsk8s lcarsk8s Agent.md Agent.md DESIGN.md DESIGN.md LICENSE LICENSE PRODUCT.md PRODUCT.md README.md README.md install.sh install.sh kmscon-lcars.conf kmscon-lcars.conf lcars-k8s-animation-demo.mp4 lcars-k8s-animation-demo.mp4 preview-graphics.gif preview-graphics.gif preview-graphics.png preview-graphics.png pyproject.toml pyproject.toml rbac.yaml rbac.yaml View all files Repository files navigation
lcars-k8s is a Kubernetes operations dashboard with two interfaces:
A pixel-rendered LCARS dashboard for a local display
A responsive Textual terminal interface for SSH, Kmscon, FbTerm, and ordinary terminals
It displays cluster CPU and memory history, node utilization, pods, deployments,
events, container logs, pod details, and Kubernetes manifests. The dashboard is
strictly read-only.
The installer targets Ubuntu Server 22.04 and 24.04. The application requires:
Network access to Ubuntu package repositories and PyPI during installation
sudo access when a required Ubuntu package is missing
A usable kubeconfig for a real cluster, or the included demo data source
An active local kernel VT for fullscreen graphical mode
The terminal interface also works over SSH. A desktop environment is not
required for fullscreen graphical mode.
Dependencies installed automatically
Running ./install.sh installs missing Ubuntu system packages:
It then creates a private virtual environment and installs these Python runtime
dependencies from pyproject.toml :
Kmscon and FbTerm are optional terminal hosts. The installer does not install
them.
1. Enter the project directory
If you downloaded an archive, extract it and enter the directory containing
install.sh :
cd lcars-k8s
If necessary, make the installer executable:
chmod +x install.sh
Do not run the entire installer with sudo . It uses sudo only when it needs
to install a missing Ubuntu package.
./install.sh
The installer performs these steps:
Installs missing Python and Xorg packages with apt-get .
Verifies that Python 3.10 or newer is available.
Creates ~/.local/share/lcars-k8s as a private virtual environment.
Installs lcars-k8s and its Python dependencies into that environment.
Creates ~/.local/bin/lcars-k8s as a symbolic link to the command.
The installer replaces an existing lcars-k8s virtual environment when rerun.
It does not install packages into the system Python environment.
3. Ensure the command is on PATH
Ubuntu normally includes ~/.local/bin after a new login. For the current
shell, use:
export PATH= " $PATH : $HOME /.local/bin "
To add it permanently when your shell does not already include it:
printf ' \nexport PATH="$PATH:$HOME/.local/bin"\n ' >> " $HOME /.bashrc "
source " $HOME /.bashrc "
4. Verify the installation
lcars-k8s --version
lcars-k8s --demo
The demo uses a synthetic cluster and does not need Kubernetes credentials.
Press q to exit.
If Python, pipx, and the required Xorg packages are already installed:
pipx install .
A pipx installation does not install the Xorg system packages. Install them
separately before using fullscreen graphical mode:
sudo apt-get update
sudo apt-get install -y xinit xserver-xorg-core
Connecting to Kubernetes
Use the current kubeconfig context
The default command uses the current context from the standard kubeconfig:
kubectl config current-context
kubectl cluster-info
lcars-k8s
The Kubernetes Python client searches the normal kubeconfig location. You can
also set KUBECONFIG :
export KUBECONFIG= " $HOME /.kube/config "
lcars-k8s
Select a kubeconfig and context explicitly
lcars-k8s --kubeconfig /path/to/config --context production
Start in one namespace when desired:
lcars-k8s --namespace kube-system
The n key cycles through namespaces after startup. The a key returns to all
namespaces.
Required Kubernetes permissions
The active Kubernetes identity needs these permissions:
Check the important permissions for the current identity:
kubectl auth can-i list nodes
kubectl auth can-i list pods --all-namespaces
kubectl auth can-i get pods/log --all-namespaces
kubectl auth can-i list events --all-namespaces
kubectl auth can-i list deployments.apps --all-namespaces
kubectl auth can-i list nodes.metrics.k8s.io
rbac.yaml creates a service account, ClusterRole, and ClusterRoleBinding with
the required permissions:
kubectl apply -f rbac.yaml
When lcars-k8s runs inside a Kubernetes pod, it automatically uses the pod's
service account if no usable kubeconfig is available. Set the pod spec's
serviceAccountName to lcars-k8s after applying rbac.yaml .
Live usage requires Metrics Server. Verify it independently:
kubectl top nodes
kubectl top pods --all-namespaces
If Metrics Server is unavailable, lcars-k8s continues using summed resource
requests. The status area indicates that allocation data is being shown instead
of live usage.
Running the terminal interface
Run against the current Kubernetes context:
lcars-k8s --demo
lcars-k8s --context production
lcars-k8s --namespace kube-system
lcars-k8s --interval 5
lcars-k8s --timeout 20
lcars-k8s --view nodes
The terminal layout adapts down to 80 columns by 24 rows. A terminal with at
least 120 columns by 40 rows is recommended. For the full palette and braille
graphs, use a terminal with truecolor and a font that contains Unicode braille
and block glyphs.
Over SSH, screen , or tmux , a suitable environment is:
export TERM=xterm-256color
lcars-k8s
Use COLORTERM=truecolor only when the terminal actually supports 24-bit
color.
The Linux virtual console supports only a limited palette and lacks many of the
required glyphs. lcars-k8s detects TERM=linux and selects its console palette
and solid graphs automatically.
You can select that mode explicitly:
lcars-k8s --colors console --glyphs solid
Available glyph profiles are:
Color and glyph settings are independent. For example:
lcars-k8s --colors console --glyphs braille
Running fullscreen graphical mode
The graphical renderer draws a 1920x1080 LCARS composition and scales it to the
active display. Normal fullscreen mode starts a private Xorg server with
xinit ; it does not require a desktop environment.
Log in directly on the machine using a local kernel VT.
Stop or leave any Kmscon, desktop compositor, or X server that owns that VT.
Confirm that xinit and Xorg are installed.
Run the command in the foreground.
command -v xinit
command -v Xorg
Launch with demo data first:
lcars-k8s --graphics --demo
Then connect to the real cluster:
lcars-k8s --graphics
The launcher clears inherited DISPLAY and WAYLAND_DISPLAY values, selects
SDL's X11 backend, starts Xorg on display :1 , and binds Xorg to the active VT
when the terminal can be identified.
Do not start fullscreen graphical mode as a background SSH job. Xorg needs an
active local session and control of a local VT.
Under an existing X11 or Wayland desktop session, use windowed mode:
lcars-k8s --graphics --windowed --demo
Specify the initial window size when needed:
lcars-k8s --graphics --windowed --resolution 1600x900 --demo
The resolution must use the WIDTHxHEIGHT form.
Raw SDL KMSDRM remains available for hardware diagnostics:
lcars-k8s --graphics --direct-kms --demo
This is not the recommended normal mode. Display and keyboard support depends
on the SDL build, DRM device, active session, and GPU driver. A black screen or
missing keyboard input can leave the process running.
Recover from another shell or SSH session with:
pgrep -af ' Xorg|xinit|lcarsk8s|lcars-k8s '
pkill -TERM -f ' /lcars-k8s/bin/python3 -m lcarsk8s '
Inspect the process list before terminating anything so that an unrelated Xorg
session is not stopped.
The graphical renderer writes the first completed design frame to:
/tmp/lcars-k8s-live-frame.png
Capture startup output for troubleshooting:
rm -f /tmp/lcars-graphics.log /tmp/lcars-k8s-live-frame.png
lcars-k8s --graphics --demo > /tmp/lcars-graphics.log 2>&1
After a failed Xorg launch, inspect:
cat /tmp/lcars-graphics.log
find " $HOME /.local/share/xorg " /var/log -maxdepth 2 -name ' Xorg*.log ' -type f 2> /dev/null
The newest Xorg log normally identifies VT ownership, GPU selection, device
permission, display lock, or driver problems.
Kmscon provides a local terminal with 256 colors, Pango font rendering, UTF-8,
and KMS/DRM output. Install it separately if desired:
sudo apt-get update
sudo apt-get install -y kmscon fonts-dejavu-core
The supplied configuration selects DejaVu Sans Mono and an LCARS-oriented
256-color palette. Review an existing configuration before replacing it:
sudo mkdir -p /etc/kmscon
sudo cp kmscon-lcars.conf /etc/kmscon/kmscon.conf
Start lcars-k8s from a shell hosted by Kmscon:
lcars-k8s --kmscon --demo
lcars-k8s --kmscon
The --kmscon option selects the application's Kmscon display profile. It does
not launch Kmscon itself. Do not set COLORTERM=truecolor ; Kmscon advertises a
256-color terminal.
Kmscon must release the active display before fullscreen graphical mode starts.
FbTerm provides another 256-color local terminal with Freetype font rendering.
Install and start FbTerm separately, then run:
lcars-k8s --fbterm --demo
lcars-k8s --fbterm
The --fbterm option sets the correct application profile and TERM=fbterm .
It does not launch FbTerm. A useful ~/.fbtermrc font configuration is:
font-names =DejaVu Sans Mono
font-size =16
Do not set COLORTERM=truecolor for FbTerm.
Key
Action
1
show or hide cluster graphs
2
open pods
3
open nodes
4
open events
5
open deployments
Tab
cycle through views
n
select the next namespace
a
show all namespaces
/ or f
filter the current view
Esc
clear or close the current filter or modal
< or >
change the pod sort column
r
reverse pod sort order
arrow keys
move the selected row or scroll a modal
Page Up or Page Down
move through graphical rows in larger steps
Home or End
select the first or last graphical row
l
open the selected pod's container log
d
open pod details and manifest
m
toggle the manifest in the terminal detail modal
c
select the next container in the terminal log modal
p
include previous container logs in the terminal log modal
Space
hold or resume scanning
+
increase the interval and scan less often
-
decrease the interval and scan more often
Ctrl+R
scan immediately
F5
scan immediately in graphical mode
?
open help
q
quit or close the active modal
Command line reference
lcars-k8s [options]
Option
Description
--demo
use synthetic cluster data
--kubeconfig PATH
use a specific kubeconfig file
--context NAME
use a specific kubeconfig context
-n NAME , --namespace NAME
start in o

[truncated]

## Original Extract

Contribute to skidvis/lcars-k8s development by creating an account on GitHub.

GitHub - skidvis/lcars-k8s · GitHub
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
skidvis
/
lcars-k8s
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
9 Commits 9 Commits Folders and files
dev dev lcarsk8s lcarsk8s Agent.md Agent.md DESIGN.md DESIGN.md LICENSE LICENSE PRODUCT.md PRODUCT.md README.md README.md install.sh install.sh kmscon-lcars.conf kmscon-lcars.conf lcars-k8s-animation-demo.mp4 lcars-k8s-animation-demo.mp4 preview-graphics.gif preview-graphics.gif preview-graphics.png preview-graphics.png pyproject.toml pyproject.toml rbac.yaml rbac.yaml View all files Repository files navigation
lcars-k8s is a Kubernetes operations dashboard with two interfaces:
A pixel-rendered LCARS dashboard for a local display
A responsive Textual terminal interface for SSH, Kmscon, FbTerm, and ordinary terminals
It displays cluster CPU and memory history, node utilization, pods, deployments,
events, container logs, pod details, and Kubernetes manifests. The dashboard is
strictly read-only.
The installer targets Ubuntu Server 22.04 and 24.04. The application requires:
Network access to Ubuntu package repositories and PyPI during installation
sudo access when a required Ubuntu package is missing
A usable kubeconfig for a real cluster, or the included demo data source
An active local kernel VT for fullscreen graphical mode
The terminal interface also works over SSH. A desktop environment is not
required for fullscreen graphical mode.
Dependencies installed automatically
Running ./install.sh installs missing Ubuntu system packages:
It then creates a private virtual environment and installs these Python runtime
dependencies from pyproject.toml :
Kmscon and FbTerm are optional terminal hosts. The installer does not install
them.
1. Enter the project directory
If you downloaded an archive, extract it and enter the directory containing
install.sh :
cd lcars-k8s
If necessary, make the installer executable:
chmod +x install.sh
Do not run the entire installer with sudo . It uses sudo only when it needs
to install a missing Ubuntu package.
./install.sh
The installer performs these steps:
Installs missing Python and Xorg packages with apt-get .
Verifies that Python 3.10 or newer is available.
Creates ~/.local/share/lcars-k8s as a private virtual environment.
Installs lcars-k8s and its Python dependencies into that environment.
Creates ~/.local/bin/lcars-k8s as a symbolic link to the command.
The installer replaces an existing lcars-k8s virtual environment when rerun.
It does not install packages into the system Python environment.
3. Ensure the command is on PATH
Ubuntu normally includes ~/.local/bin after a new login. For the current
shell, use:
export PATH= " $PATH : $HOME /.local/bin "
To add it permanently when your shell does not already include it:
printf ' \nexport PATH="$PATH:$HOME/.local/bin"\n ' >> " $HOME /.bashrc "
source " $HOME /.bashrc "
4. Verify the installation
lcars-k8s --version
lcars-k8s --demo
The demo uses a synthetic cluster and does not need Kubernetes credentials.
Press q to exit.
If Python, pipx, and the required Xorg packages are already installed:
pipx install .
A pipx installation does not install the Xorg system packages. Install them
separately before using fullscreen graphical mode:
sudo apt-get update
sudo apt-get install -y xinit xserver-xorg-core
Connecting to Kubernetes
Use the current kubeconfig context
The default command uses the current context from the standard kubeconfig:
kubectl config current-context
kubectl cluster-info
lcars-k8s
The Kubernetes Python client searches the normal kubeconfig location. You can
also set KUBECONFIG :
export KUBECONFIG= " $HOME /.kube/config "
lcars-k8s
Select a kubeconfig and context explicitly
lcars-k8s --kubeconfig /path/to/config --context production
Start in one namespace when desired:
lcars-k8s --namespace kube-system
The n key cycles through namespaces after startup. The a key returns to all
namespaces.
Required Kubernetes permissions
The active Kubernetes identity needs these permissions:
Check the important permissions for the current identity:
kubectl auth can-i list nodes
kubectl auth can-i list pods --all-namespaces
kubectl auth can-i get pods/log --all-namespaces
kubectl auth can-i list events --all-namespaces
kubectl auth can-i list deployments.apps --all-namespaces
kubectl auth can-i list nodes.metrics.k8s.io
rbac.yaml creates a service account, ClusterRole, and ClusterRoleBinding with
the required permissions:
kubectl apply -f rbac.yaml
When lcars-k8s runs inside a Kubernetes pod, it automatically uses the pod's
service account if no usable kubeconfig is available. Set the pod spec's
serviceAccountName to lcars-k8s after applying rbac.yaml .
Live usage requires Metrics Server. Verify it independently:
kubectl top nodes
kubectl top pods --all-namespaces
If Metrics Server is unavailable, lcars-k8s continues using summed resource
requests. The status area indicates that allocation data is being shown instead
of live usage.
Running the terminal interface
Run against the current Kubernetes context:
lcars-k8s --demo
lcars-k8s --context production
lcars-k8s --namespace kube-system
lcars-k8s --interval 5
lcars-k8s --timeout 20
lcars-k8s --view nodes
The terminal layout adapts down to 80 columns by 24 rows. A terminal with at
least 120 columns by 40 rows is recommended. For the full palette and braille
graphs, use a terminal with truecolor and a font that contains Unicode braille
and block glyphs.
Over SSH, screen , or tmux , a suitable environment is:
export TERM=xterm-256color
lcars-k8s
Use COLORTERM=truecolor only when the terminal actually supports 24-bit
color.
The Linux virtual console supports only a limited palette and lacks many of the
required glyphs. lcars-k8s detects TERM=linux and selects its console palette
and solid graphs automatically.
You can select that mode explicitly:
lcars-k8s --colors console --glyphs solid
Available glyph profiles are:
Color and glyph settings are independent. For example:
lcars-k8s --colors console --glyphs braille
Running fullscreen graphical mode
The graphical renderer draws a 1920x1080 LCARS composition and scales it to the
active display. Normal fullscreen mode starts a private Xorg server with
xinit ; it does not require a desktop environment.
Log in directly on the machine using a local kernel VT.
Stop or leave any Kmscon, desktop compositor, or X server that owns that VT.
Confirm that xinit and Xorg are installed.
Run the command in the foreground.
command -v xinit
command -v Xorg
Launch with demo data first:
lcars-k8s --graphics --demo
Then connect to the real cluster:
lcars-k8s --graphics
The launcher clears inherited DISPLAY and WAYLAND_DISPLAY values, selects
SDL's X11 backend, starts Xorg on display :1 , and binds Xorg to the active VT
when the terminal can be identified.
Do not start fullscreen graphical mode as a background SSH job. Xorg needs an
active local session and control of a local VT.
Under an existing X11 or Wayland desktop session, use windowed mode:
lcars-k8s --graphics --windowed --demo
Specify the initial window size when needed:
lcars-k8s --graphics --windowed --resolution 1600x900 --demo
The resolution must use the WIDTHxHEIGHT form.
Raw SDL KMSDRM remains available for hardware diagnostics:
lcars-k8s --graphics --direct-kms --demo
This is not the recommended normal mode. Display and keyboard support depends
on the SDL build, DRM device, active session, and GPU driver. A black screen or
missing keyboard input can leave the process running.
Recover from another shell or SSH session with:
pgrep -af ' Xorg|xinit|lcarsk8s|lcars-k8s '
pkill -TERM -f ' /lcars-k8s/bin/python3 -m lcarsk8s '
Inspect the process list before terminating anything so that an unrelated Xorg
session is not stopped.
The graphical renderer writes the first completed design frame to:
/tmp/lcars-k8s-live-frame.png
Capture startup output for troubleshooting:
rm -f /tmp/lcars-graphics.log /tmp/lcars-k8s-live-frame.png
lcars-k8s --graphics --demo > /tmp/lcars-graphics.log 2>&1
After a failed Xorg launch, inspect:
cat /tmp/lcars-graphics.log
find " $HOME /.local/share/xorg " /var/log -maxdepth 2 -name ' Xorg*.log ' -type f 2> /dev/null
The newest Xorg log normally identifies VT ownership, GPU selection, device
permission, display lock, or driver problems.
Kmscon provides a local terminal with 256 colors, Pango font rendering, UTF-8,
and KMS/DRM output. Install it separately if desired:
sudo apt-get update
sudo apt-get install -y kmscon fonts-dejavu-core
The supplied configuration selects DejaVu Sans Mono and an LCARS-oriented
256-color palette. Review an existing configuration before replacing it:
sudo mkdir -p /etc/kmscon
sudo cp kmscon-lcars.conf /etc/kmscon/kmscon.conf
Start lcars-k8s from a shell hosted by Kmscon:
lcars-k8s --kmscon --demo
lcars-k8s --kmscon
The --kmscon option selects the application's Kmscon display profile. It does
not launch Kmscon itself. Do not set COLORTERM=truecolor ; Kmscon advertises a
256-color terminal.
Kmscon must release the active display before fullscreen graphical mode starts.
FbTerm provides another 256-color local terminal with Freetype font rendering.
Install and start FbTerm separately, then run:
lcars-k8s --fbterm --demo
lcars-k8s --fbterm
The --fbterm option sets the correct application profile and TERM=fbterm .
It does not launch FbTerm. A useful ~/.fbtermrc font configuration is:
font-names =DejaVu Sans Mono
font-size =16
Do not set COLORTERM=truecolor for FbTerm.
Key
Action
1
show or hide cluster graphs
2
open pods
3
open nodes
4
open events
5
open deployments
Tab
cycle through views
n
select the next namespace
a
show all namespaces
/ or f
filter the current view
Esc
clear or close the current filter or modal
< or >
change the pod sort column
r
reverse pod sort order
arrow keys
move the selected row or scroll a modal
Page Up or Page Down
move through graphical rows in larger steps
Home or End
select the first or last graphical row
l
open the selected pod's container log
d
open pod details and manifest
m
toggle the manifest in the terminal detail modal
c
select the next container in the terminal log modal
p
include previous container logs in the terminal log modal
Space
hold or resume scanning
+
increase the interval and scan less often
-
decrease the interval and scan more often
Ctrl+R
scan immediately
F5
scan immediately in graphical mode
?
open help
q
quit or close the active modal
Command line reference
lcars-k8s [options]
Option
Description
--demo
use synthetic cluster data
--kubeconfig PATH
use a specific kubeconfig file
--context NAME
use a specific kubeconfig context
-n NAME , --namespace NAME
start in o

[truncated]
