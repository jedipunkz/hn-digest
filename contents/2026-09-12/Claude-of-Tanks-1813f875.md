---
source: "https://github.com/Kevin-Liu-01/Claude-of-Tanks"
hn_url: "https://news.ycombinator.com/item?id=49677183"
title: "Claude of Tanks"
article_title: "GitHub - Kevin-Liu-01/Claude-of-Tanks: A World of Tanks-style, Vite-powered, engine-free pure Three.js armored combat simulator resolving plate-level armor, ballistics, modules, spotting, and physics, with 121 tanks and 20 destructible battlefields. Playable entirely in the browser on desktop and mo\n[truncated]"
image: "https://opengraph.githubassets.com/307c3221f14c6310682fca034d26da8ee6a3eda2e9b7377ba0e6f1df3848e2e9/Kevin-Liu-01/Claude-of-Tanks"
author: "tiredpanda"
captured_at: "2026-09-12T21:21:58Z"
capture_tool: "hn-digest"
hn_id: 49677183
score: 1
comments: 0
posted_at: "2026-09-12T21:03:49Z"
tags:
  - hacker-news
---

# Claude of Tanks

- HN: [49677183](https://news.ycombinator.com/item?id=49677183)
- Source: [github.com](https://github.com/Kevin-Liu-01/Claude-of-Tanks)
- Score: 1
- Comments: 0
- Posted: 2026-09-12T21:03:49Z

## Translation

Title: Claude of Tanks
Article title: GitHub - Kevin-Liu-01/Claude-of-Tanks: A World of Tanks-style, Vite-powered, engine-free pure Three.js armored combat simulator resolving plate-level armor, ballistics, modules, spotting, and physics, with 121 tanks and 20 destructible battlefields. Playable entirely in the browser on desktop and mo
[truncated]
Description: A World of Tanks-style, Vite-powered, engine-free pure Three.js armored combat simulator resolving plate-level armor, ballistics, modules, spotting, and physics, with 121 tanks and 20 destructible battlefields. Playable entirely in the browser on desktop and mobile devices. Built end-to-end by a mul
[truncated]

Article text:
GitHub - Kevin-Liu-01/Claude-of-Tanks: A World of Tanks-style, Vite-powered, engine-free pure Three.js armored combat simulator resolving plate-level armor, ballistics, modules, spotting, and physics, with 121 tanks and 20 destructible battlefields. Playable entirely in the browser on desktop and mobile devices. Built end-to-end by a multi-agent Claude/Codex pipeline. · GitHub
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
Kevin-Liu-01
/
Claude-of-Tanks
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3,760 Commits 3,760 Commits Folders and files
.agent-docs .agent-docs .agents/ skills/ improve-threejs .agents/ skills/ improve-threejs .github/ workflows .github/ workflows LICENSES LICENSES api api cloudflare/ signaling cloudflare/ signaling deploy deploy docs docs public public scripts scripts server server shots shots src src tools tools .dockerignore .dockerignore .env.example .env.example .env.multiplayer.example .env.multiplayer.example .gitignore .gitignore .vercelignore .vercelignore 404.html 404.html AGENTS.md AGENTS.md Dockerfile.multiplayer Dockerfile.multiplayer Dockerfile.selfhost Dockerfile.selfhost LICENSE LICENSE LICENSE-POLICY.md LICENSE-POLICY.md NOTICE.md NOTICE.md README.md README.md SKILL.md SKILL.md compose.multiplayer.yaml compose.multiplayer.yaml compose.selfhost.yaml compose.selfhost.yaml docs-ai.html docs-ai.html docs-audio.html docs-audio.html docs-build.html docs-build.html docs-interface.html docs-interface.html docs-models.html docs-models.html docs-multiplayer.html docs-multiplayer.html docs-performance.html docs-performance.html docs-rendering.html docs-rendering.html docs-simulation.html docs-simulation.html docs-studio.html docs-studio.html docs-topic.html docs-topic.html docs-vehicles.html docs-vehicles.html docs-worlds.html docs-worlds.html docs.html docs.html gallery.html gallery.html gt-lock.json gt-lock.json gt.config.json gt.config.json home.html home.html index.html index.html middleware.ts middleware.ts package-lock.json package-lock.json package.json package.json skills-lock.json skills-lock.json tsconfig.json tsconfig.json vercel.json vercel.json vite.config.ts vite.config.ts View all files Repository files navigation
Free browser-native armored combat built with Three.js . Take 171 production-visible first-party procedural vehicles
across 30 battlefields with physical gunnery, plate-level armor, internal damage, guided missiles,
magazine autoloaders, terrain-following suspension, X-ray killcams, multiplayer rooms, and Scene Studio.
These are handmade, deterministic scenes captured with the current game renderer—not concept art. The
88-frame showcase archive contains 13 owner-selected scenes, 30 action frames,
30 close foreground compositions, five directed Studio frames, and ten interface states. The
landing selection records the six hero frames, feature reel, five camera-rail
films, 24-shot mosaic, and directed Strv 122 versus Leclerc sequence now published on the site.
Fight: enter Standard Battle, Capture the Flag, Zone Control, armed super-speed Turbo Ball, or cooperative
Endless Horde in solo or multiplayer, with physical shell travel, armor geometry, component damage, spotting,
terrain, collision, destructible structures, persistent wrecks, mode-specific respawns, and authority-owned results.
Inspect: open any vehicle in Tank Gallery, articulate the live rig, isolate armor or internal anatomy, and export an
exact-surface review packet from the same specification used in combat.
Direct: stage any roster vehicle on any battlefield in Scene Studio, animate actors and cameras, schedule game FX,
and capture stills or video from reproducible scene data.
Current runtime
Fleet
171 production-visible and 208 keyed local-development procedural vehicles across 210 saved roster records; 0 GLB-sourced playables
Worlds
30 authored battlefields with shared structures, wrecks, utility networks, loose props, placement, collision, and destruction
Authority
Fixed 60 Hz movement, ballistics, armor, damage, spotting, bots, destructibles, and result
Presentation
Direct Three.js/WebGL renderer with a measured 120 FPS test path, adaptive quality, stable shadows, SMAA/FSR, and GPU recovery
Play
Five battle rules, solo bots, browser-hosted private rooms, LAN rooms, room chat, spectators, respawns, and rematches
Platforms
Mouse/keyboard and complete touch controls with safe-area layout and device-adaptive rendering
Languages
English and Simplified Chinese across the game, Garage, Studio, Gallery, and public docs; local reviewed catalogs managed with General Translation
Tools
Scene Studio, Tank Gallery, exact-surface review, deterministic capture, vehicle anatomy, and release gates
Language selection lives under Settings → Graphics → Language . See the
localization guide for coverage, intentional exclusions,
the General Translation workflow, and release gates.
The provenance gate currently reports 181 first-party procedural battle playables, 0 GLB-sourced playables, and 7
tracked external comparison models with exact source records . Comparison inputs are never a playable loading path and are stripped from public builds.
Each frame below opens a looping WebM recorded from the game. The camera rails use live vehicles, maps, ballistics,
recoil, impacts, sparks, smoke, debris, and destruction effects.
OPEN THE FULL LANDING SHOWCASE
The fleet does not reduce every vehicle to the same gun and movement model. Specifications can add their own loading,
guidance, suspension, ammunition, anatomy, and control behavior while remaining inside the fixed-step authority.
Magazine autoloaders track ready rounds, intra-magazine cycle time, full replenishment, manual magazine reloads,
shell changes, and damage to the loading mechanism. The Leclerc, Type 90, PL-01, and four-round PL-01 105 each carry
their own magazine timings.
Guided missiles travel at visible weapon-specific speeds and steer toward the authority-owned aim point. Launcher
damage, rack damage, ammunition count, reload state, guidance, impact, and presentation all remain synchronized.
Hydropneumatic aim uses the same canonical pose for armor, muzzle, collision, wheels, tracks, remote snapshots, and
the camera; there is no separate visual-only tank tilt.
Free look keeps the sight and turret on target while the camera moves. Hold Shift on keyboard, RB on a standard
controller, or assign another binding in settings.
Plate-level armor resolves the actual plate, slope, impact angle, normalization, ricochet, overmatch, spaced armor,
composites, ERA, and separate kinetic/chemical protection.
Five ammunition families model muzzle velocity, gravity, penetration loss, ricochet, and damage differently.
Internal anatomy tracks crew, ammunition racks, engine, fuel, gun, turret ring, optics, radio, and tracks.
Tank-specific mobility combines drivetrain, terrain resistance, per-wheel support, suspension-damped hull attitude,
flexible terrain-following tracks, collision, ramming, and crushable cover.
Real battlefield knowledge combines view range, concealment, movement/firing bloom, foliage, radio sharing, and the
15 m bush rule. Multiplayer authority filters hidden enemies before serializing a snapshot.
Every playable hull, turret, gun, fitting, suspension, road wheel, and track run is assembled by the repository's
first-party vehicle pipeline. Vehicle changes pass combat-anatomy receipts, generated technical diagrams, geometry
checks, visual fingerprints, and a targeted release gate.
The Garage is its own authored presentation system. Verdant retains the enclosed Motor Pool; the other nine locations
use lightweight scene packs derived from a real battlefield's terrain, structures, PBR materials, detailed trees,
approach route, horizon, sky, and atmosphere without loading that battlefield. One immutable hero pose keeps the tank
and camera comparable across the complete set, while connected maintenance stations, heavy-lift equipment, real fleet
exhibits, and full-detail tank parts establish a distinct service story around every platform.
The release gate reviews every location from four orbit angles and responsive layouts, rejects unsupported or floating
parts, certifies geometry-derived structure collision, and requires a score of at least 90/100 with bounded draw calls,
triangles, transitions, cache residency, and repeated-cycle memory. See the
Garage environments guide for the complete visual, lifecycle, and verification contract.
Renderer, drivers, and performance
The renderer treats quality as a device contract instead of a single desktop preset. It selects a GPU/driver-aware
profile, caps pixel density, prewarms shader paths, adapts costly effects, and can recover from a black frame or WebGL
context loss. The current presentation path combines:
four quality-scaled, stable texel-anchored shadow cascades with articulation-aware tank shadow hulls;
fused output grading, anti-aliasing, adaptive render scale, fog/atmosphere, and bounded transparent depth work;
instance/batch paths for repeated world objects, pooled particles, and explicit GPU resource disposal;
reusable hot-loop scratch state, fixed-step simulation, render interpolation, and high-refresh presentation;
an in-game diagnostics surface for FPS, ping, frame timing, draw calls, triangles, memory, network telemetry, quality,
and renderer/driver identity.
The renderer reached 120 FPS on the certified test hardware . Actual performance depends on refresh rate, browser,
thermal limits, GPU and driver, resolution, and quality level. Combat rules remain fixed at 60 Hz at every render rate.
LAN and browser-hosted private rooms use the shared renderer-free movement and combat rules. No database or dedicated game server is required; Internet room codes need a lightweight signaling endpoint, and restricted networks may need TURN relay. Ranked tooling remains internal rather than a player-facing mode. See multiplayer hosting .
Clients send intent, never trusted hits or damage. Snapshot filtering, local prediction/reconciliation, bounded remote
interpolation, reliable fire edges, reconnectable room state, and separate control/chat delivery keep a moving and firing
7v7 battle responsive without giving the client authority.
Two screens, opposing sights: paired live 1v1 captures preserve the real battle HUD, authoritative simulation, filtered snapshots, and each commander's view.
Action
Keyboard and mouse
Standard controller
Drive and steer
WASD or arrow keys
Left stick
Aim and fire
Mouse + LMB
Right stick + RT
Precision sight
Hold RMB by default; wheel in also enters
LT
Free look without moving the turret
Hold Shift or Left Alt
RB
Select ammuniti

[truncated]

## Original Extract

A World of Tanks-style, Vite-powered, engine-free pure Three.js armored combat simulator resolving plate-level armor, ballistics, modules, spotting, and physics, with 121 tanks and 20 destructible battlefields. Playable entirely in the browser on desktop and mobile devices. Built end-to-end by a mul
[truncated]

GitHub - Kevin-Liu-01/Claude-of-Tanks: A World of Tanks-style, Vite-powered, engine-free pure Three.js armored combat simulator resolving plate-level armor, ballistics, modules, spotting, and physics, with 121 tanks and 20 destructible battlefields. Playable entirely in the browser on desktop and mobile devices. Built end-to-end by a multi-agent Claude/Codex pipeline. · GitHub
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
Kevin-Liu-01
/
Claude-of-Tanks
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3,760 Commits 3,760 Commits Folders and files
.agent-docs .agent-docs .agents/ skills/ improve-threejs .agents/ skills/ improve-threejs .github/ workflows .github/ workflows LICENSES LICENSES api api cloudflare/ signaling cloudflare/ signaling deploy deploy docs docs public public scripts scripts server server shots shots src src tools tools .dockerignore .dockerignore .env.example .env.example .env.multiplayer.example .env.multiplayer.example .gitignore .gitignore .vercelignore .vercelignore 404.html 404.html AGENTS.md AGENTS.md Dockerfile.multiplayer Dockerfile.multiplayer Dockerfile.selfhost Dockerfile.selfhost LICENSE LICENSE LICENSE-POLICY.md LICENSE-POLICY.md NOTICE.md NOTICE.md README.md README.md SKILL.md SKILL.md compose.multiplayer.yaml compose.multiplayer.yaml compose.selfhost.yaml compose.selfhost.yaml docs-ai.html docs-ai.html docs-audio.html docs-audio.html docs-build.html docs-build.html docs-interface.html docs-interface.html docs-models.html docs-models.html docs-multiplayer.html docs-multiplayer.html docs-performance.html docs-performance.html docs-rendering.html docs-rendering.html docs-simulation.html docs-simulation.html docs-studio.html docs-studio.html docs-topic.html docs-topic.html docs-vehicles.html docs-vehicles.html docs-worlds.html docs-worlds.html docs.html docs.html gallery.html gallery.html gt-lock.json gt-lock.json gt.config.json gt.config.json home.html home.html index.html index.html middleware.ts middleware.ts package-lock.json package-lock.json package.json package.json skills-lock.json skills-lock.json tsconfig.json tsconfig.json vercel.json vercel.json vite.config.ts vite.config.ts View all files Repository files navigation
Free browser-native armored combat built with Three.js . Take 171 production-visible first-party procedural vehicles
across 30 battlefields with physical gunnery, plate-level armor, internal damage, guided missiles,
magazine autoloaders, terrain-following suspension, X-ray killcams, multiplayer rooms, and Scene Studio.
These are handmade, deterministic scenes captured with the current game renderer—not concept art. The
88-frame showcase archive contains 13 owner-selected scenes, 30 action frames,
30 close foreground compositions, five directed Studio frames, and ten interface states. The
landing selection records the six hero frames, feature reel, five camera-rail
films, 24-shot mosaic, and directed Strv 122 versus Leclerc sequence now published on the site.
Fight: enter Standard Battle, Capture the Flag, Zone Control, armed super-speed Turbo Ball, or cooperative
Endless Horde in solo or multiplayer, with physical shell travel, armor geometry, component damage, spotting,
terrain, collision, destructible structures, persistent wrecks, mode-specific respawns, and authority-owned results.
Inspect: open any vehicle in Tank Gallery, articulate the live rig, isolate armor or internal anatomy, and export an
exact-surface review packet from the same specification used in combat.
Direct: stage any roster vehicle on any battlefield in Scene Studio, animate actors and cameras, schedule game FX,
and capture stills or video from reproducible scene data.
Current runtime
Fleet
171 production-visible and 208 keyed local-development procedural vehicles across 210 saved roster records; 0 GLB-sourced playables
Worlds
30 authored battlefields with shared structures, wrecks, utility networks, loose props, placement, collision, and destruction
Authority
Fixed 60 Hz movement, ballistics, armor, damage, spotting, bots, destructibles, and result
Presentation
Direct Three.js/WebGL renderer with a measured 120 FPS test path, adaptive quality, stable shadows, SMAA/FSR, and GPU recovery
Play
Five battle rules, solo bots, browser-hosted private rooms, LAN rooms, room chat, spectators, respawns, and rematches
Platforms
Mouse/keyboard and complete touch controls with safe-area layout and device-adaptive rendering
Languages
English and Simplified Chinese across the game, Garage, Studio, Gallery, and public docs; local reviewed catalogs managed with General Translation
Tools
Scene Studio, Tank Gallery, exact-surface review, deterministic capture, vehicle anatomy, and release gates
Language selection lives under Settings → Graphics → Language . See the
localization guide for coverage, intentional exclusions,
the General Translation workflow, and release gates.
The provenance gate currently reports 181 first-party procedural battle playables, 0 GLB-sourced playables, and 7
tracked external comparison models with exact source records . Comparison inputs are never a playable loading path and are stripped from public builds.
Each frame below opens a looping WebM recorded from the game. The camera rails use live vehicles, maps, ballistics,
recoil, impacts, sparks, smoke, debris, and destruction effects.
OPEN THE FULL LANDING SHOWCASE
The fleet does not reduce every vehicle to the same gun and movement model. Specifications can add their own loading,
guidance, suspension, ammunition, anatomy, and control behavior while remaining inside the fixed-step authority.
Magazine autoloaders track ready rounds, intra-magazine cycle time, full replenishment, manual magazine reloads,
shell changes, and damage to the loading mechanism. The Leclerc, Type 90, PL-01, and four-round PL-01 105 each carry
their own magazine timings.
Guided missiles travel at visible weapon-specific speeds and steer toward the authority-owned aim point. Launcher
damage, rack damage, ammunition count, reload state, guidance, impact, and presentation all remain synchronized.
Hydropneumatic aim uses the same canonical pose for armor, muzzle, collision, wheels, tracks, remote snapshots, and
the camera; there is no separate visual-only tank tilt.
Free look keeps the sight and turret on target while the camera moves. Hold Shift on keyboard, RB on a standard
controller, or assign another binding in settings.
Plate-level armor resolves the actual plate, slope, impact angle, normalization, ricochet, overmatch, spaced armor,
composites, ERA, and separate kinetic/chemical protection.
Five ammunition families model muzzle velocity, gravity, penetration loss, ricochet, and damage differently.
Internal anatomy tracks crew, ammunition racks, engine, fuel, gun, turret ring, optics, radio, and tracks.
Tank-specific mobility combines drivetrain, terrain resistance, per-wheel support, suspension-damped hull attitude,
flexible terrain-following tracks, collision, ramming, and crushable cover.
Real battlefield knowledge combines view range, concealment, movement/firing bloom, foliage, radio sharing, and the
15 m bush rule. Multiplayer authority filters hidden enemies before serializing a snapshot.
Every playable hull, turret, gun, fitting, suspension, road wheel, and track run is assembled by the repository's
first-party vehicle pipeline. Vehicle changes pass combat-anatomy receipts, generated technical diagrams, geometry
checks, visual fingerprints, and a targeted release gate.
The Garage is its own authored presentation system. Verdant retains the enclosed Motor Pool; the other nine locations
use lightweight scene packs derived from a real battlefield's terrain, structures, PBR materials, detailed trees,
approach route, horizon, sky, and atmosphere without loading that battlefield. One immutable hero pose keeps the tank
and camera comparable across the complete set, while connected maintenance stations, heavy-lift equipment, real fleet
exhibits, and full-detail tank parts establish a distinct service story around every platform.
The release gate reviews every location from four orbit angles and responsive layouts, rejects unsupported or floating
parts, certifies geometry-derived structure collision, and requires a score of at least 90/100 with bounded draw calls,
triangles, transitions, cache residency, and repeated-cycle memory. See the
Garage environments guide for the complete visual, lifecycle, and verification contract.
Renderer, drivers, and performance
The renderer treats quality as a device contract instead of a single desktop preset. It selects a GPU/driver-aware
profile, caps pixel density, prewarms shader paths, adapts costly effects, and can recover from a black frame or WebGL
context loss. The current presentation path combines:
four quality-scaled, stable texel-anchored shadow cascades with articulation-aware tank shadow hulls;
fused output grading, anti-aliasing, adaptive render scale, fog/atmosphere, and bounded transparent depth work;
instance/batch paths for repeated world objects, pooled particles, and explicit GPU resource disposal;
reusable hot-loop scratch state, fixed-step simulation, render interpolation, and high-refresh presentation;
an in-game diagnostics surface for FPS, ping, frame timing, draw calls, triangles, memory, network telemetry, quality,
and renderer/driver identity.
The renderer reached 120 FPS on the certified test hardware . Actual performance depends on refresh rate, browser,
thermal limits, GPU and driver, resolution, and quality level. Combat rules remain fixed at 60 Hz at every render rate.
LAN and browser-hosted private rooms use the shared renderer-free movement and combat rules. No database or dedicated game server is required; Internet room codes need a lightweight signaling endpoint, and restricted networks may need TURN relay. Ranked tooling remains internal rather than a player-facing mode. See multiplayer hosting .
Clients send intent, never trusted hits or damage. Snapshot filtering, local prediction/reconciliation, bounded remote
interpolation, reliable fire edges, reconnectable room state, and separate control/chat delivery keep a moving and firing
7v7 battle responsive without giving the client authority.
Two screens, opposing sights: paired live 1v1 captures preserve the real battle HUD, authoritative simulation, filtered snapshots, and each commander's view.
Action
Keyboard and mouse
Standard controller
Drive and steer
WASD or arrow keys
Left stick
Aim and fire
Mouse + LMB
Right stick + RT
Precision sight
Hold RMB by default; wheel in also enters
LT
Free look without moving the turret
Hold Shift or Left Alt
RB
Select ammuniti

[truncated]
