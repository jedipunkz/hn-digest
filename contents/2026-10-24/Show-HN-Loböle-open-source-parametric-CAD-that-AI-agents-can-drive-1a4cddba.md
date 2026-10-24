---
source: "https://github.com/eliasaronson/lobole"
hn_url: "https://news.ycombinator.com/item?id=50223070"
title: "Show HN: Loböle, open-source parametric CAD that AI agents can drive"
article_title: "GitHub - eliasaronson/lobole: Parametric 3D CAD with assemblies, drawings, FEM and CAM, scriptable from Python and agents over MCP. · GitHub"
image: "https://repository-images.githubusercontent.com/1424814351/b5dad83b-cbd9-4b57-af1e-8ff730a632b8"
author: "eliasaronson"
captured_at: "2026-10-24T17:53:14Z"
capture_tool: "hn-digest"
hn_id: 50223070
score: 1
comments: 0
posted_at: "2026-10-24T16:58:21Z"
tags:
  - hacker-news
---

# Show HN: Loböle, open-source parametric CAD that AI agents can drive

- HN: [50223070](https://news.ycombinator.com/item?id=50223070)
- Source: [github.com](https://github.com/eliasaronson/lobole)
- Score: 1
- Comments: 0
- Posted: 2026-10-24T16:58:21Z

## Translation

Title: Show HN: Loböle, open-source parametric CAD that AI agents can drive
Article title: GitHub - eliasaronson/lobole: Parametric 3D CAD with assemblies, drawings, FEM and CAM, scriptable from Python and agents over MCP. · GitHub
Description: Parametric 3D CAD with assemblies, drawings, FEM and CAM, scriptable from Python and agents over MCP. - eliasaronson/lobole
HN text: I could never get used to FreeCAD, so I built my own. It turns out Claude is really good at using it. Loböle is a parametric B-Rep CAD built on OpenCASCADE and Vulkan, with sketches, assemblies, drawings, FEM and CAM. The GUI, Python scripts and AI agents all go through the same commands, so an agent can do anything you can in the GUI. The engine, watch and bicycle in the README were each built by Claude writing a script against it.

Article text:
GitHub - eliasaronson/lobole: Parametric 3D CAD with assemblies, drawings, FEM and CAM, scriptable from Python and agents over MCP. · GitHub
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
eliasaronson
lobole Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Parametric 3D CAD with assemblies, drawings, FEM and CAM, scriptable from Python and agents over MCP.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2,523 Commits 2,523 Commits Folders and files
.githooks .githooks .github .github LICENSES LICENSES api api assets assets benchmarks benchmarks cmake cmake deps deps docs docs examples/ catalogs examples/ catalogs generators generators nix nix packaging packaging python/ lobole python/ lobole src src tests tests tools tools .clang-format .clang-format .clang-tidy .clang-tidy .clangd .clangd .gitignore .gitignore .gitmodules .gitmodules .mcp.json .mcp.json ADDITIONAL_PERMISSION.md ADDITIONAL_PERMISSION.md AGENTS.md AGENTS.md CLA.md CLA.md CLAUDE.md CLAUDE.md CMakeLists.txt CMakeLists.txt CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY.md THIRD_PARTY.md TRADEMARKS.md TRADEMARKS.md VERSION VERSION flake.lock flake.lock flake.nix flake.nix pyproject.toml pyproject.toml shell.nix shell.nix View all files Repository files navigation
Open-source parametric CAD: B-Rep modeling on OpenCASCADE, a Vulkan
viewport, and one command layer that the GUI, Python scripts and AI agents
all drive. Linux first; Windows builds exist, macOS is untested.
Version 0.1: so expect some rough edges. Issues are welcome.
Parametric modeling: a constraint sketcher, features that survive
upstream edits (stable topological naming), variables and expressions,
configurations, sheet metal, bent tube.
Assemblies: mates, sub-assemblies, motion studies with clash checks,
exploded views, a BOM, standard fasteners.
Drawings: sheets with projected views, dimensions and tables, out to
SVG and PDF.
Simulation: static and modal FEM, and topology optimization, on the
part itself.
Manufacturing: 2.5D, 3-axis and 4th-axis milling posted for Mach3,
LinuxCNC or GRBL, each file verified by simulating it; checks for 3D
printing.
Versioning: history, named versions, branches and merges, backed by
git.
Scripting and agents: JSON-RPC, a Python SDK and an MCP server reach
every command the GUI has; a terminal panel runs an agent inside the
window.
Interchange: STEP, STL and BREP in; those plus glTF, OBJ, PLY, DXF,
SVG and PDF out.
CAM has cut real parts on one machine only : a 6090 CNC router on
Mach3. Simulate and dry-run every file before cutting.
Every model here is built by a generator, an ordinary Python program driving
the API; generators/README.md lists them all.
Linux: lobole-<version>-x86_64.AppImage , the full GUI on glibc
2.39 or newer (Ubuntu 24.04, Debian 13, Fedora 40 and later). chmod +x
it and run it; linked onto your PATH as lobole , the MCP server can
start it by itself.
Ubuntu 26.04: sudo apt install ./lobole_<version>-1_amd64.deb .
Windows: a zip or an installer. Tested on Windows by hand and
under Wine on every release, unsigned (SmartScreen warns), no FEM.
Linux, headless: an AppImage with the engine and its JSON-RPC
interface, no window, for servers ( caveats ).
Nix: nix run github:eliasaronson/lobole , the full GUI. Outside
NixOS the GUI needs nixGL .
Container: ghcr.io/eliasaronson/lobole-headless
( below ).
Python: pip install lobole , the client, SDK and MCP server
( docs/scripting ).
Other distributions: build from source .
Ubuntu 24.04: CI installs exactly this list on every push:
sudo apt install git cmake ninja-build build-essential python3 \
libvulkan-dev glslc libglfw3-dev libglm-dev \
libspdlog-dev libgtk-3-dev nlohmann-json3-dev libeigen3-dev libboost-dev \
libocct-data-exchange-dev libocct-modeling-algorithms-dev libocct-modeling-data-dev \
libocct-foundation-dev libocct-ocaf-dev libocct-visualization-dev \
libtbb-dev libvterm-dev
Ubuntu 26.04: the same list, plus libnglib-dev netgen python3-netgen libspectra-dev for FEM meshing and modal analysis; CI installs that too. The last two are needed only because
Ubuntu's Netgen CMake config checks for them.
Nix / NixOS: nix-shell in the repo root pins everything. The Vulkan
driver comes from your system ( hardware.graphics.enable = true ).
Flatpak: tools/flatpak-build.sh builds for any distribution
( packaging/flatpak/ ).
Other distributions: the Ubuntu list says what is needed; OpenCASCADE
must be 7.6 or newer, CMake 3.14+, and the compiler C++23.
Windows cross-compiles from Linux ( packaging/windows/ ). macOS is
untested.
Clone with --recurse-submodules (or run git submodule update --init --recursive ) for the vendored dependencies.
cmake prints which optional features it compiled in. CAM and the terminal
panel can be switched off ( -DENABLE_CAM=OFF , -DENABLE_TERMINAL=OFF ), and
then they do not exist at all. Without Netgen there is no FEM meshing,
and without Spectra no modal analysis; Ubuntu has both from 26.04 on,
and every release download except Windows carries them.
On OpenCASCADE 7.6 (Ubuntu 24.04), PLY export is missing and some booleans
come out wrong ( details ).
cmake -S . -B build -G Ninja
cmake --build build
cmake --install build --prefix "$HOME/.local" # optional: app, assets, desktop entry
Without Dear ImGui Test Engine
The test engine is the one dependency that is not open source: a company
over 2M USD turnover that builds or ships Loböle needs a paid licence from
DISCO HELLO. -DENABLE_TEST_ENGINE=OFF removes it, along with the GUI test
suite and the remote ui_action tier.
build/lobole [-f path/to/file.step]
Container (headless)
The engine and its JSON-RPC interface, no GPU or display, with the full
feature set (OpenCASCADE 7.9, Netgen, Spectra):
docker run --rm -p 8437:8437 ghcr.io/eliasaronson/lobole-headless
lobole-remote 8437 get_state # from pip install lobole
nix build .#dockerHeadless builds the same image.
The JSON-RPC interface is unauthenticated and can read and write
files . The image listens on all interfaces so the published port works;
never publish it on a network you do not trust.
tools/ctest -LE gui -j$(nproc) # headless: no GPU, no display
./tools/gui-tests-xvfb.sh # the GUI suite on a virtual display
More in CONTRIBUTING.md .
With an AI agent: install the MCP server and register it, e.g. with
Claude Code. It attaches to a running Loböle, or starts one without a
window if none is running:
uv tool install lobole # or: pip install lobole
claude mcp add lobole -- lobole-mcp
Scripts drive a live session over JSON-RPC (localhost only, opt-in):
./lobole --remote 8437 # or --headless --remote 8437, with no GPU
export PYTHONPATH=$PWD/python
python3 -m lobole.client 8437 get_state
python3 -m lobole.client 8437 add_feature \
'{"feature": {"type": "Box", "name": "Box 1", "on": "Top",
"width": 2, "depth": 2, "height": 2}}'
Every GUI command is reachable, plus UI actions that click and type in the
real window. See the Python SDK and the
MCP server .
Copyright © Elias Högbom Aronsson
Loböle is licensed under the GNU Affero General Public License, version 3
or any later version ( LICENSE ). A distributed or
network-hosted derivative must be released under the AGPL with its source.
Files you make with Loböle (models, drawings, exports) are yours.
It comes without any warranty . Check simulation results before relying
on them for anything that carries load, and dry-run every G-code file.
Linking Dear ImGui Test Engine
An additional permission under section 7 of the AGPL allows linking Loböle
with Dear ImGui Test Engine and conveying the result; the text is in
ADDITIONAL_PERMISSION.md .
Other products named here are trademarks of their owners; Loböle is not
affiliated with any of them. The name and logo of Loböle are not licensed
under the AGPL : a fork or modified build takes its own name
( TRADEMARKS.md ).
Listed with their licences, including the dynamically linked LGPL
libraries, in THIRD_PARTY.md .
Parametric 3D CAD with assemblies, drawings, FEM and CAM, scriptable from Python and agents over MCP.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Parametric 3D CAD with assemblies, drawings, FEM and CAM, scriptable from Python and agents over MCP. - eliasaronson/lobole

I could never get used to FreeCAD, so I built my own. It turns out Claude is really good at using it. Loböle is a parametric B-Rep CAD built on OpenCASCADE and Vulkan, with sketches, assemblies, drawings, FEM and CAM. The GUI, Python scripts and AI agents all go through the same commands, so an agent can do anything you can in the GUI. The engine, watch and bicycle in the README were each built by Claude writing a script against it.

GitHub - eliasaronson/lobole: Parametric 3D CAD with assemblies, drawings, FEM and CAM, scriptable from Python and agents over MCP. · GitHub
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
eliasaronson
lobole Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Parametric 3D CAD with assemblies, drawings, FEM and CAM, scriptable from Python and agents over MCP.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2,523 Commits 2,523 Commits Folders and files
.githooks .githooks .github .github LICENSES LICENSES api api assets assets benchmarks benchmarks cmake cmake deps deps docs docs examples/ catalogs examples/ catalogs generators generators nix nix packaging packaging python/ lobole python/ lobole src src tests tests tools tools .clang-format .clang-format .clang-tidy .clang-tidy .clangd .clangd .gitignore .gitignore .gitmodules .gitmodules .mcp.json .mcp.json ADDITIONAL_PERMISSION.md ADDITIONAL_PERMISSION.md AGENTS.md AGENTS.md CLA.md CLA.md CLAUDE.md CLAUDE.md CMakeLists.txt CMakeLists.txt CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY.md THIRD_PARTY.md TRADEMARKS.md TRADEMARKS.md VERSION VERSION flake.lock flake.lock flake.nix flake.nix pyproject.toml pyproject.toml shell.nix shell.nix View all files Repository files navigation
Open-source parametric CAD: B-Rep modeling on OpenCASCADE, a Vulkan
viewport, and one command layer that the GUI, Python scripts and AI agents
all drive. Linux first; Windows builds exist, macOS is untested.
Version 0.1: so expect some rough edges. Issues are welcome.
Parametric modeling: a constraint sketcher, features that survive
upstream edits (stable topological naming), variables and expressions,
configurations, sheet metal, bent tube.
Assemblies: mates, sub-assemblies, motion studies with clash checks,
exploded views, a BOM, standard fasteners.
Drawings: sheets with projected views, dimensions and tables, out to
SVG and PDF.
Simulation: static and modal FEM, and topology optimization, on the
part itself.
Manufacturing: 2.5D, 3-axis and 4th-axis milling posted for Mach3,
LinuxCNC or GRBL, each file verified by simulating it; checks for 3D
printing.
Versioning: history, named versions, branches and merges, backed by
git.
Scripting and agents: JSON-RPC, a Python SDK and an MCP server reach
every command the GUI has; a terminal panel runs an agent inside the
window.
Interchange: STEP, STL and BREP in; those plus glTF, OBJ, PLY, DXF,
SVG and PDF out.
CAM has cut real parts on one machine only : a 6090 CNC router on
Mach3. Simulate and dry-run every file before cutting.
Every model here is built by a generator, an ordinary Python program driving
the API; generators/README.md lists them all.
Linux: lobole-<version>-x86_64.AppImage , the full GUI on glibc
2.39 or newer (Ubuntu 24.04, Debian 13, Fedora 40 and later). chmod +x
it and run it; linked onto your PATH as lobole , the MCP server can
start it by itself.
Ubuntu 26.04: sudo apt install ./lobole_<version>-1_amd64.deb .
Windows: a zip or an installer. Tested on Windows by hand and
under Wine on every release, unsigned (SmartScreen warns), no FEM.
Linux, headless: an AppImage with the engine and its JSON-RPC
interface, no window, for servers ( caveats ).
Nix: nix run github:eliasaronson/lobole , the full GUI. Outside
NixOS the GUI needs nixGL .
Container: ghcr.io/eliasaronson/lobole-headless
( below ).
Python: pip install lobole , the client, SDK and MCP server
( docs/scripting ).
Other distributions: build from source .
Ubuntu 24.04: CI installs exactly this list on every push:
sudo apt install git cmake ninja-build build-essential python3 \
libvulkan-dev glslc libglfw3-dev libglm-dev \
libspdlog-dev libgtk-3-dev nlohmann-json3-dev libeigen3-dev libboost-dev \
libocct-data-exchange-dev libocct-modeling-algorithms-dev libocct-modeling-data-dev \
libocct-foundation-dev libocct-ocaf-dev libocct-visualization-dev \
libtbb-dev libvterm-dev
Ubuntu 26.04: the same list, plus libnglib-dev netgen python3-netgen libspectra-dev for FEM meshing and modal analysis; CI installs that too. The last two are needed only because
Ubuntu's Netgen CMake config checks for them.
Nix / NixOS: nix-shell in the repo root pins everything. The Vulkan
driver comes from your system ( hardware.graphics.enable = true ).
Flatpak: tools/flatpak-build.sh builds for any distribution
( packaging/flatpak/ ).
Other distributions: the Ubuntu list says what is needed; OpenCASCADE
must be 7.6 or newer, CMake 3.14+, and the compiler C++23.
Windows cross-compiles from Linux ( packaging/windows/ ). macOS is
untested.
Clone with --recurse-submodules (or run git submodule update --init --recursive ) for the vendored dependencies.
cmake prints which optional features it compiled in. CAM and the terminal
panel can be switched off ( -DENABLE_CAM=OFF , -DENABLE_TERMINAL=OFF ), and
then they do not exist at all. Without Netgen there is no FEM meshing,
and without Spectra no modal analysis; Ubuntu has both from 26.04 on,
and every release download except Windows carries them.
On OpenCASCADE 7.6 (Ubuntu 24.04), PLY export is missing and some booleans
come out wrong ( details ).
cmake -S . -B build -G Ninja
cmake --build build
cmake --install build --prefix "$HOME/.local" # optional: app, assets, desktop entry
Without Dear ImGui Test Engine
The test engine is the one dependency that is not open source: a company
over 2M USD turnover that builds or ships Loböle needs a paid licence from
DISCO HELLO. -DENABLE_TEST_ENGINE=OFF removes it, along with the GUI test
suite and the remote ui_action tier.
build/lobole [-f path/to/file.step]
Container (headless)
The engine and its JSON-RPC interface, no GPU or display, with the full
feature set (OpenCASCADE 7.9, Netgen, Spectra):
docker run --rm -p 8437:8437 ghcr.io/eliasaronson/lobole-headless
lobole-remote 8437 get_state # from pip install lobole
nix build .#dockerHeadless builds the same image.
The JSON-RPC interface is unauthenticated and can read and write
files . The image listens on all interfaces so the published port works;
never publish it on a network you do not trust.
tools/ctest -LE gui -j$(nproc) # headless: no GPU, no display
./tools/gui-tests-xvfb.sh # the GUI suite on a virtual display
More in CONTRIBUTING.md .
With an AI agent: install the MCP server and register it, e.g. with
Claude Code. It attaches to a running Loböle, or starts one without a
window if none is running:
uv tool install lobole # or: pip install lobole
claude mcp add lobole -- lobole-mcp
Scripts drive a live session over JSON-RPC (localhost only, opt-in):
./lobole --remote 8437 # or --headless --remote 8437, with no GPU
export PYTHONPATH=$PWD/python
python3 -m lobole.client 8437 get_state
python3 -m lobole.client 8437 add_feature \
'{"feature": {"type": "Box", "name": "Box 1", "on": "Top",
"width": 2, "depth": 2, "height": 2}}'
Every GUI command is reachable, plus UI actions that click and type in the
real window. See the Python SDK and the
MCP server .
Copyright © Elias Högbom Aronsson
Loböle is licensed under the GNU Affero General Public License, version 3
or any later version ( LICENSE ). A distributed or
network-hosted derivative must be released under the AGPL with its source.
Files you make with Loböle (models, drawings, exports) are yours.
It comes without any warranty . Check simulation results before relying
on them for anything that carries load, and dry-run every G-code file.
Linking Dear ImGui Test Engine
An additional permission under section 7 of the AGPL allows linking Loböle
with Dear ImGui Test Engine and conveying the result; the text is in
ADDITIONAL_PERMISSION.md .
Other products named here are trademarks of their owners; Loböle is not
affiliated with any of them. The name and logo of Loböle are not licensed
under the AGPL : a fork or modified build takes its own name
( TRADEMARKS.md ).
Listed with their licences, including the dynamically linked LGPL
libraries, in THIRD_PARTY.md .
Parametric 3D CAD with assemblies, drawings, FEM and CAM, scriptable from Python and agents over MCP.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
