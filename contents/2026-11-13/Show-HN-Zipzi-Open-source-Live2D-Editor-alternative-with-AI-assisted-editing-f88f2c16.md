---
source: "https://github.com/getzipzi/zipzi-2d"
hn_url: "https://news.ycombinator.com/item?id=50477525"
title: "Show HN: Zipzi – Open-source Live2D Editor alternative with AI-assisted editing"
article_title: "GitHub - getzipzi/zipzi-2d: An AI Agent-Ready 2D Modeling and Animation Tool. · GitHub"
image: "https://opengraph.githubassets.com/95744218201b2b0bc27624f46f8178fbed39969ac31067facb4a1c6fb9191d58/getzipzi/zipzi-2d"
author: "pakholeung37"
captured_at: "2026-11-13T03:10:31Z"
capture_tool: "hn-digest"
hn_id: 50477525
score: 1
comments: 0
posted_at: "2026-11-13T02:54:37Z"
tags:
  - hacker-news
---

# Show HN: Zipzi – Open-source Live2D Editor alternative with AI-assisted editing

- HN: [50477525](https://news.ycombinator.com/item?id=50477525)
- Source: [github.com](https://github.com/getzipzi/zipzi-2d)
- Score: 1
- Comments: 0
- Posted: 2026-11-13T02:54:37Z

## Translation

Title: Show HN: Zipzi – Open-source Live2D Editor alternative with AI-assisted editing
Article title: GitHub - getzipzi/zipzi-2d: An AI Agent-Ready 2D Modeling and Animation Tool. · GitHub
Description: An AI Agent-Ready 2D Modeling and Animation Tool. Contribute to getzipzi/zipzi-2d development by creating an account on GitHub.

Article text:
GitHub - getzipzi/zipzi-2d: An AI Agent-Ready 2D Modeling and Animation Tool. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
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
getzipzi
zipzi-2d Public Notifications You must be signed in to change notification settings
Star 8 ( 8 ) You must be signed in to star a repository
An AI Agent-Ready 2D Modeling and Animation Tool.
Readme GPL-3.0, MIT licenses found Activity Custom properties Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
775 Commits 775 Commits Folders and files
.github/ workflows .github/ workflows .vscode .vscode benchmarks benchmarks crates crates docs docs fixtures fixtures models models native/ project-preview native/ project-preview third_party third_party tools tools vendor vendor .clang-format .clang-format .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules .python-version .python-version AGENTS.md AGENTS.md CMakeLists.txt CMakeLists.txt Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-GPL-3.0 LICENSE-GPL-3.0 LICENSE-MIT LICENSE-MIT LICENSING.md LICENSING.md MONOREPO.md MONOREPO.md README.ja.md README.ja.md README.md README.md README.zh-CN.md README.zh-CN.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
An open-source 2D modeling and animation tool for creating alongside AI.
You shape the artwork. Your AI agent helps with the modeling work. Zipzi connects
a visual desktop editor with a Python API, so you and your agent can inspect and
edit the same model, share undo history, and build on each other’s changes.
Zipzi 2D is an open-source alternative to Live2D Cubism Editor, built with Rust
and GPUI. Its name comes from Cantonese 摺紙 ( zip3 zi2 ), meaning “origami”.
Download Alpha · Python SDK · Report an issue · Contribute
Early development: There is no stable release yet. APIs and file formats
may change. The editor supports macOS and Windows.
Open source and commercial friendly — Create, refine, and use your models
commercially with an open-source editor and an MIT-licensed SDK.
Native and high performance — Built with Rust and GPUI for low memory
usage and responsive editing, even with complex models and dense meshes.
Agent-ready — Create from scratch, collaborate with AI, or import an
existing model for further refinement.
zipzi-2d-demo.mp4
Your editor. Your AI agent. One model.
Describe what you want to your AI coding agent. It uses the Python SDK to inspect
models, view rendered images and make edits you can review and undo in the editor.
You and your agent work on the same document, with shared undo history.
Bring your own Python-capable agent and point it to the
SDK guide .
First code: let AI give every mesh a meaningful name
You describe the task; your AI agent writes and runs the code. The examples
below were generated by an AI coding agent and tested against a running Editor.
They show how the agent works through the Python SDK—you do not need to write
these scripts yourself. Your role is to explain what you want and review the
results in the editor.
Ask your agent: “Look at every mesh and give it a meaningful name.”
Install the Python client
(CPython 3.14), start a compatible Editor with GPU observation support, and open
one model. An image-capable coding agent can run the following two steps.
1. The agent observes every mesh. Its generated script exports one image per mesh, with its
wireframe highlighted over the textured model, plus a JSON naming worksheet.
observe supplies visual evidence; the agent supplies the semantic interpretation.
import json
from pathlib import Path
from uuid import uuid4
import zipzi
from zipzi import inspection , observation , spatial
output = Path ( "mesh-review-" + uuid4 (). hex [: 8 ]). resolve ()
output . mkdir ()
with zipzi . connect () as runtime :
sessions = runtime . sessions . list ()
if len ( sessions ) != 1 :
raise RuntimeError ( "Open exactly one model for this example." )
session = runtime . sessions . attach ( sessions [ 0 ]. id )
# Follow every page: a single list() call may not contain all meshes.
meshes , cursor , source = [], None , None
while True :
page = session . meshes . list ( cursor = cursor )
if source is None :
source = page . source
if page . source != source :
raise RuntimeError ( "Model changed; restart the inspection." )
meshes . extend ( page . items )
cursor = page . next_cursor
if cursor is None :
break
rows = []
for index , mesh in enumerate ( meshes ):
# Use this mesh's bounds to frame it, including hidden geometry.
# {} evaluates default parameter values, not the editor's current pose.
query = spatial . Bounds (
targets = ( spatial . MeshTarget ( value = mesh . id ),), include_hidden = True ,
)
bounds = session . modeling . spatial (
({},), ( query ,), expected = source ,
). wait (). value . frames [ 0 ]. results [ 0 ]. value . model_bounds
row = dict ( id = mesh . id , old_name = mesh . name , new_name = mesh . name , image = None )
if bounds is not None :
x0 , y0 , x1 , y1 = bounds
margin = max ( x1 - x0 , y1 - y0 , 1.0 ) * 0.15
view = observation . View (
width = 768 , height = 768 ,
roi = ( x0 - margin , y0 - margin , x1 + margin , y1 + margin ),
topology = observation . Overlay (
targets = ( observation . MeshTarget ( value = mesh . id ),),
wireframe = True , include_hidden = True ,
),
)
# Freeze all observations to the inventory's source. If the model
# changes, the SDK refuses stale reads; rerun instead of mixing them.
images = session . modeling . observe (({},), view , expected = source ). wait ()
report = inspection . save ( i
[truncated]
2. The agent applies the semantic names together. After you review the
proposed names, the agent writes and runs the following script with the path to
names.json . You do not need to copy the code or run a terminal command. Keep
the same model open and unchanged between inspection and application.
import json
import sys
from pathlib import Path
import zipzi
review = json . loads ( Path ( sys . argv [ 1 ]). read_text ( encoding = "utf-8" ))
rows = review [ "meshes" ]
if len ({ row [ "id" ] for row in rows }) != len ( rows ):
raise ValueError ( "Duplicate mesh IDs in the worksheet." )
if any ( not isinstance ( row [ "new_name" ], str ) or not row [ "new_name" ]. strip ()
for row in rows ):
raise ValueError ( "Every mesh needs a nonempty name." )
with zipzi . connect () as runtime :
# Rebuild the typed handle; a restarted Editor requires a new inspection.
session = runtime . sessions . attach ( zipzi . SessionId ( ** review [ "session_id" ]))
# All renames commit together as one undo step. A missing mesh or changed
# original name aborts the edit instead of leaving a partly renamed model.
with session . edit ( "AI: semantic mesh names" ) as edit :
for row in rows :
current = edit . meshes . get ( row [ "id" ]). value
if current . name != row [ "old_name" ]:
raise RuntimeError ( "Mesh changed; inspect the model again." )
if row [ "new_name" ] != row [ "old_name" ]:
edit . meshes . rename (
row [ "id" ], row [ "new_name" ], expected_name = row [ "old_name" ],
)
print ( "Review the names in the editor. Undo to revert, or save to keep them." )
No project save happens automatically. The name check detects renamed or removed
meshes, not every possible geometry or texture change; repeat the inspection if
the model has changed. For more workflows, see the
capability guide .
Built for hands-on creation, too
Work with meshes, parameters, deformers and BlendShapes, then bring them together
in the animation workspace. Generate mesh geometry or refine it by hand. Mesh
Edit lets you work on the posed model or switch to the source image, while
resampling stored shapes as the geometry changes.
Python scripts can also run independently with a local Runtime for automation
outside the editor. Explore the Python SDK
for modeling, animation authoring, previews and inspection.
Use the editor: Download an Alpha build from
GitHub Releases . Check the
release notes for the available platform builds and requirements.
Automate with Python: Follow the
installation and connection guide .
The SDK currently requires CPython 3.14.
Build your own tools: Explore the Rust SDK
for programmatic authoring.
Area
What you can work with
Modeling
Meshes, parameters, keyforms, parts, deformers, BlendShapes and Glue
Animation
Motion, expression, physics and pose authoring; animation previews
Project I/O
PSD import and Cubism-related formats, including Model3/MOC3, projects and runtime sidecars; support and loss rules vary by format and direction
Automation
Python sessions, batch edits, shared undo/redo, project operations and model inspection
Platforms
macOS and Windows
See the SDK capability guide for details.
GPU inspection requires a compatible host and usable GPU. Format support does
not imply lossless round trips for every model.
Try an Alpha build and tell us what gets in your way .
For bug reports, include the build version, platform and steps to reproduce;
share model assets only when you have permission to do so.
For code contributions, start with the repository map ,
editor guide or
Python SDK guide . The
documentation index covers deeper references.
Development commands and focused test selection are below.
Use Rust 1.95 or newer, uv 0.10.7 or newer, and CPython 3.14.
The macOS Editor requires Xcode Command Line Tools and a Metal-capable GPU.
The Windows Editor requires
MSVC, the Windows SDK and a usable GPU.
Run these commands from the repository root:
uv sync --locked
python3 tools/workspace/cargo_domain.py --check
# Focused SDK development
cargo test --locked -p zipzi-sdk --test integration physics_editing::
# Native Editor (macOS)
cargo run --locked -p zipzi-editor
./tools/test-editor.sh fast timeline::
# Repository tooling tests; no native application build
uv run --locked python tools/test_tools.py
For the Python SDK, ./tools/test-sdk.sh python refreshes its native adapter
and runs Local/Connected behavior and static types. python-connected installs
the pure wheel and validates real subprocess host discovery and session binding;
add --editor-host /absolute/path/to/zipzi-editor for the current desktop host.
Root pyproject.toml and uv.lock manage development tools.
Focused test commands and coverage
Choose the smallest gate covering the change, then broaden when needed. Run
commands from the repository root; inspect available modes with
./tools/test-sdk.sh --help , ./tools/test-editor.sh --help or
uv run --locked python tools/test_tools.py --help .
Filters and filenames are placeholders; omit them to run the selected gate.
For a crate's integration targets, inspect its README and tests/ , or list cases
with cargo test --locked -p <crate> --test <target> -- --list .
Editor shell gates clear ZIPZI_DEV_SESSION_DIR ; clear i

[truncated]

## Original Extract

An AI Agent-Ready 2D Modeling and Animation Tool. Contribute to getzipzi/zipzi-2d development by creating an account on GitHub.

GitHub - getzipzi/zipzi-2d: An AI Agent-Ready 2D Modeling and Animation Tool. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
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
getzipzi
zipzi-2d Public Notifications You must be signed in to change notification settings
Star 8 ( 8 ) You must be signed in to star a repository
An AI Agent-Ready 2D Modeling and Animation Tool.
Readme GPL-3.0, MIT licenses found Activity Custom properties Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
775 Commits 775 Commits Folders and files
.github/ workflows .github/ workflows .vscode .vscode benchmarks benchmarks crates crates docs docs fixtures fixtures models models native/ project-preview native/ project-preview third_party third_party tools tools vendor vendor .clang-format .clang-format .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules .python-version .python-version AGENTS.md AGENTS.md CMakeLists.txt CMakeLists.txt Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-GPL-3.0 LICENSE-GPL-3.0 LICENSE-MIT LICENSE-MIT LICENSING.md LICENSING.md MONOREPO.md MONOREPO.md README.ja.md README.ja.md README.md README.md README.zh-CN.md README.zh-CN.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
An open-source 2D modeling and animation tool for creating alongside AI.
You shape the artwork. Your AI agent helps with the modeling work. Zipzi connects
a visual desktop editor with a Python API, so you and your agent can inspect and
edit the same model, share undo history, and build on each other’s changes.
Zipzi 2D is an open-source alternative to Live2D Cubism Editor, built with Rust
and GPUI. Its name comes from Cantonese 摺紙 ( zip3 zi2 ), meaning “origami”.
Download Alpha · Python SDK · Report an issue · Contribute
Early development: There is no stable release yet. APIs and file formats
may change. The editor supports macOS and Windows.
Open source and commercial friendly — Create, refine, and use your models
commercially with an open-source editor and an MIT-licensed SDK.
Native and high performance — Built with Rust and GPUI for low memory
usage and responsive editing, even with complex models and dense meshes.
Agent-ready — Create from scratch, collaborate with AI, or import an
existing model for further refinement.
zipzi-2d-demo.mp4
Your editor. Your AI agent. One model.
Describe what you want to your AI coding agent. It uses the Python SDK to inspect
models, view rendered images and make edits you can review and undo in the editor.
You and your agent work on the same document, with shared undo history.
Bring your own Python-capable agent and point it to the
SDK guide .
First code: let AI give every mesh a meaningful name
You describe the task; your AI agent writes and runs the code. The examples
below were generated by an AI coding agent and tested against a running Editor.
They show how the agent works through the Python SDK—you do not need to write
these scripts yourself. Your role is to explain what you want and review the
results in the editor.
Ask your agent: “Look at every mesh and give it a meaningful name.”
Install the Python client
(CPython 3.14), start a compatible Editor with GPU observation support, and open
one model. An image-capable coding agent can run the following two steps.
1. The agent observes every mesh. Its generated script exports one image per mesh, with its
wireframe highlighted over the textured model, plus a JSON naming worksheet.
observe supplies visual evidence; the agent supplies the semantic interpretation.
import json
from pathlib import Path
from uuid import uuid4
import zipzi
from zipzi import inspection , observation , spatial
output = Path ( "mesh-review-" + uuid4 (). hex [: 8 ]). resolve ()
output . mkdir ()
with zipzi . connect () as runtime :
sessions = runtime . sessions . list ()
if len ( sessions ) != 1 :
raise RuntimeError ( "Open exactly one model for this example." )
session = runtime . sessions . attach ( sessions [ 0 ]. id )
# Follow every page: a single list() call may not contain all meshes.
meshes , cursor , source = [], None , None
while True :
page = session . meshes . list ( cursor = cursor )
if source is None :
source = page . source
if page . source != source :
raise RuntimeError ( "Model changed; restart the inspection." )
meshes . extend ( page . items )
cursor = page . next_cursor
if cursor is None :
break
rows = []
for index , mesh in enumerate ( meshes ):
# Use this mesh's bounds to frame it, including hidden geometry.
# {} evaluates default parameter values, not the editor's current pose.
query = spatial . Bounds (
targets = ( spatial . MeshTarget ( value = mesh . id ),), include_hidden = True ,
)
bounds = session . modeling . spatial (
({},), ( query ,), expected = source ,
). wait (). value . frames [ 0 ]. results [ 0 ]. value . model_bounds
row = dict ( id = mesh . id , old_name = mesh . name , new_name = mesh . name , image = None )
if bounds is not None :
x0 , y0 , x1 , y1 = bounds
margin = max ( x1 - x0 , y1 - y0 , 1.0 ) * 0.15
view = observation . View (
width = 768 , height = 768 ,
roi = ( x0 - margin , y0 - margin , x1 + margin , y1 + margin ),
topology = observation . Overlay (
targets = ( observation . MeshTarget ( value = mesh . id ),),
wireframe = True , include_hidden = True ,
),
)
# Freeze all observations to the inventory's source. If the model
# changes, the SDK refuses stale reads; rerun instead of mixing them.
images = session . modeling . observe (({},), view , expected = source ). wait ()
report = inspection . save ( i
[truncated]
2. The agent applies the semantic names together. After you review the
proposed names, the agent writes and runs the following script with the path to
names.json . You do not need to copy the code or run a terminal command. Keep
the same model open and unchanged between inspection and application.
import json
import sys
from pathlib import Path
import zipzi
review = json . loads ( Path ( sys . argv [ 1 ]). read_text ( encoding = "utf-8" ))
rows = review [ "meshes" ]
if len ({ row [ "id" ] for row in rows }) != len ( rows ):
raise ValueError ( "Duplicate mesh IDs in the worksheet." )
if any ( not isinstance ( row [ "new_name" ], str ) or not row [ "new_name" ]. strip ()
for row in rows ):
raise ValueError ( "Every mesh needs a nonempty name." )
with zipzi . connect () as runtime :
# Rebuild the typed handle; a restarted Editor requires a new inspection.
session = runtime . sessions . attach ( zipzi . SessionId ( ** review [ "session_id" ]))
# All renames commit together as one undo step. A missing mesh or changed
# original name aborts the edit instead of leaving a partly renamed model.
with session . edit ( "AI: semantic mesh names" ) as edit :
for row in rows :
current = edit . meshes . get ( row [ "id" ]). value
if current . name != row [ "old_name" ]:
raise RuntimeError ( "Mesh changed; inspect the model again." )
if row [ "new_name" ] != row [ "old_name" ]:
edit . meshes . rename (
row [ "id" ], row [ "new_name" ], expected_name = row [ "old_name" ],
)
print ( "Review the names in the editor. Undo to revert, or save to keep them." )
No project save happens automatically. The name check detects renamed or removed
meshes, not every possible geometry or texture change; repeat the inspection if
the model has changed. For more workflows, see the
capability guide .
Built for hands-on creation, too
Work with meshes, parameters, deformers and BlendShapes, then bring them together
in the animation workspace. Generate mesh geometry or refine it by hand. Mesh
Edit lets you work on the posed model or switch to the source image, while
resampling stored shapes as the geometry changes.
Python scripts can also run independently with a local Runtime for automation
outside the editor. Explore the Python SDK
for modeling, animation authoring, previews and inspection.
Use the editor: Download an Alpha build from
GitHub Releases . Check the
release notes for the available platform builds and requirements.
Automate with Python: Follow the
installation and connection guide .
The SDK currently requires CPython 3.14.
Build your own tools: Explore the Rust SDK
for programmatic authoring.
Area
What you can work with
Modeling
Meshes, parameters, keyforms, parts, deformers, BlendShapes and Glue
Animation
Motion, expression, physics and pose authoring; animation previews
Project I/O
PSD import and Cubism-related formats, including Model3/MOC3, projects and runtime sidecars; support and loss rules vary by format and direction
Automation
Python sessions, batch edits, shared undo/redo, project operations and model inspection
Platforms
macOS and Windows
See the SDK capability guide for details.
GPU inspection requires a compatible host and usable GPU. Format support does
not imply lossless round trips for every model.
Try an Alpha build and tell us what gets in your way .
For bug reports, include the build version, platform and steps to reproduce;
share model assets only when you have permission to do so.
For code contributions, start with the repository map ,
editor guide or
Python SDK guide . The
documentation index covers deeper references.
Development commands and focused test selection are below.
Use Rust 1.95 or newer, uv 0.10.7 or newer, and CPython 3.14.
The macOS Editor requires Xcode Command Line Tools and a Metal-capable GPU.
The Windows Editor requires
MSVC, the Windows SDK and a usable GPU.
Run these commands from the repository root:
uv sync --locked
python3 tools/workspace/cargo_domain.py --check
# Focused SDK development
cargo test --locked -p zipzi-sdk --test integration physics_editing::
# Native Editor (macOS)
cargo run --locked -p zipzi-editor
./tools/test-editor.sh fast timeline::
# Repository tooling tests; no native application build
uv run --locked python tools/test_tools.py
For the Python SDK, ./tools/test-sdk.sh python refreshes its native adapter
and runs Local/Connected behavior and static types. python-connected installs
the pure wheel and validates real subprocess host discovery and session binding;
add --editor-host /absolute/path/to/zipzi-editor for the current desktop host.
Root pyproject.toml and uv.lock manage development tools.
Focused test commands and coverage
Choose the smallest gate covering the change, then broaden when needed. Run
commands from the repository root; inspect available modes with
./tools/test-sdk.sh --help , ./tools/test-editor.sh --help or
uv run --locked python tools/test_tools.py --help .
Filters and filenames are placeholders; omit them to run the selected gate.
For a crate's integration targets, inspect its README and tests/ , or list cases
with cargo test --locked -p <crate> --test <target> -- --list .
Editor shell gates clear ZIPZI_DEV_SESSION_DIR ; clear i

[truncated]
