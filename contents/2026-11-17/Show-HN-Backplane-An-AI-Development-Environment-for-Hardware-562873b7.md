---
source: "https://backplaneworks.com/"
hn_url: "https://news.ycombinator.com/item?id=50528196"
title: "Show HN: Backplane – An AI Development Environment for Hardware"
article_title: "Backplane"
image: "https://backplaneworks.com/static/og/backplane-preview-v2.jpg"
author: "i2cjak"
captured_at: "2026-11-17T06:51:36Z"
capture_tool: "hn-digest"
hn_id: 50528196
score: 1
comments: 1
posted_at: "2026-11-17T06:41:02Z"
tags:
  - hacker-news
---

# Show HN: Backplane – An AI Development Environment for Hardware

- HN: [50528196](https://news.ycombinator.com/item?id=50528196)
- Source: [backplaneworks.com](https://backplaneworks.com/)
- Score: 1
- Comments: 1
- Posted: 2026-11-17T06:41:02Z

## Translation

Title: Show HN: Backplane – An AI Development Environment for Hardware
Article title: Backplane
Description: A TOTALLY open-source workspace for agentic hardware development.
HN text: We made Backplane to actually use Fable, Astra, etc. as real electrical engineers! It incorporates KiCad + your favorite harness (Claude Code, Codex, OpenCode, Grok) to accelerate hardware development. You can use it "autonomously" or work alongside the agent, making changes or comments to redirect the agent as it works. Also has some slick skills for schematic editing, BOM management. Hardware is in its "tab to accept changes"/copilot era rather than its agentic era imo.

Article text:
Source on GitHub ★ 1.1k
curl -fsSL https://backplaneworks.com/install.sh | sh Copy
backplane runs on YOUR compute (mobile apps coming soon)
Backplane uses AGENTS to design HARDWARE.
Backplane is FREE AND OPEN SOURCE FOREVER.
Backplane is created + maintained by real electrical engineers.
PROJECTS! THREADS! PERSISTENT BOTS (they're cats)
4 bots · 2 rooms · 6 screens · space_set bot_send room_post memory_save routine_set remind_person
Very nice tools that interact DIRECTLY with KiCad (no file rewriting, still human editable!)
KiStack built in! Automatic updates as skills improve!
skills/bom/SKILL.md · 78 lines · 6.2 KB
description: Resolve and review exact PCB parts, packages, electrical ratings, availability, and manufacturer BOM fields for KiCad projects.
skills/export/SKILL.md · 120 lines · 4.6 KB
description: Generate deterministic KiCad ERC, DRC, fabrication, assembly, documentation, and 3D export outputs with kicad-cli.
references/kicad_cli_command_fp_export_svg.md
references/kicad_cli_command_fp_upgrade.md
references/kicad_cli_command_introduction.md
references/kicad_cli_command_jobset_run.md
references/kicad_cli_command_pcb_drc.md
references/kicad_cli_command_pcb_export_3dpdf.md
references/kicad_cli_command_pcb_export_brep.md
references/kicad_cli_command_pcb_export_drill.md
references/kicad_cli_command_pcb_export_dxf.md
references/kicad_cli_command_pcb_export_gencad.md
references/kicad_cli_command_pcb_export_gerbers.md
references/kicad_cli_command_pcb_export_glb.md
references/kicad_cli_command_pcb_export_hpgl.md
references/kicad_cli_command_pcb_export_ipc2581.md
references/kicad_cli_command_pcb_export_ipcd356.md
references/kicad_cli_command_pcb_export_odb.md
references/kicad_cli_command_pcb_export_pdf.md
references/kicad_cli_command_pcb_export_ply.md
references/kicad_cli_command_pcb_export_pos.md
references/kicad_cli_command_pcb_export_ps.md
references/kicad_cli_command_pcb_export_stats.md
references/kicad_cli_command_pcb_export_step.md
references/kicad_cli_command_pcb_export_stl.md
references/kicad_cli_command_pcb_export_stpz.md
references/kicad_cli_command_pcb_export_svg.md
references/kicad_cli_command_pcb_export_u3d.md
references/kicad_cli_command_pcb_export_vrml.md
references/kicad_cli_command_pcb_export_xao.md
references/kicad_cli_command_pcb_import.md
references/kicad_cli_command_pcb_render.md
references/kicad_cli_command_pcb_upgrade.md
references/kicad_cli_command_sch_erc.md
references/kicad_cli_command_sch_export_bom.md
references/kicad_cli_command_sch_export_dxf.md
references/kicad_cli_command_sch_export_hpgl.md
references/kicad_cli_command_sch_export_netlist.md
references/kicad_cli_command_sch_export_pdf.md
references/kicad_cli_command_sch_export_ps.md
references/kicad_cli_command_sch_export_python_bom.md
references/kicad_cli_command_sch_export_svg.md
references/kicad_cli_command_sch_upgrade.md
references/kicad_cli_command_sym_export_svg.md
references/kicad_cli_command_sym_upgrade.md
references/kicad_cli_command_version.md
** ERC messages: 0 Errors 0 Warnings 0 tome_nano_100-drc.rpt ** Found 0 DRC violations **
** Found 0 unconnected pads **
** Found 0 Footprint errors ** positions_raw.csv Ref,Val,Package,PosX,PosY,Rot,Side
"C10","1uF","C_0402_1005Metric_L",150.913501,-131.187500,0.000000,bottom
"C11","1uF","C_0402_1005Metric_L",151.401001,-136.212500,0.000000,bottom positions.csv (scripts/convert_position.py) Designator,Mid X,Mid Y,Layer,Rotation
C10,150.913501,-131.187500,B,0.000000
C11,151.401001,-136.212500,B,0.000000 top.png (kicad-cli pcb render --side top) 22 files in gerbs/
skills/footprint/SKILL.md · 25 lines · 2.6 KB
description: Generate compliant KiCad footprints from official KiCad Library Tools
skills/gerbers/SKILL.md · 18 lines · 0.9 KB
description: Generate, render, and visually review KiCad Gerber layers as a fabrication-output inspection step.
skills/layout/SKILL.md · 23 lines · 4.3 KB
description: Place components and route KiCad PCB layouts with footprint checks, schematic pin swaps, and signal integrity best practices.
## Editing the active board in place
skills/layout-placement/SKILL.md · 344 lines · 21.5 KB
description: Place components on a KiCad board so each part sits where its circuit needs it, placement DRC passes, and every pad stays reachable, with IPC scripts that move parts and measure distances and routability. Covers reading the circuit, fixed parts, IC orientation, and decoupling, regulator, clock, RF and filter parts. Use for component placement; not for routing or schematic design.
## Save progress to the configured board
## 1. Read the circuit before moving anything
## 2. Fix the immovable, then orient the ICs
## 3. Place in order of how little distance each connection tolerates
## 4. Keep every pad reachable
230 legal candidates for C48 at U1.D1 (current gap 2.925 mm)
blocks gap_mm partner_mm x y rot faces blocked pins
0 2.867 9.000 97.600 95.050 90 True []
0 2.868 9.000 97.850 95.050 90 True []
0 2.884 9.000 98.100 95.050 90 True []
0 3.067 9.000 98.100 95.300 0 True []
0 3.069 9.000 98.350 95.300 0 True [] skills/panelize/SKILL.md · 90 lines · 3.4 KB
description: Use KiKit to panelize a given PCB
skills/pcb/SKILL.md · 32 lines · 2.5 KB
description: Review KiCad PCB layouts with DRC, layer renders, and 3D-model interference inspection.
## Editing the active board in place
skills/pcb-product-render/SKILL.md · 53 lines · 5.0 KB
description: Create realistic Blender product renders from KiCad PCB projects, including refreshed 3D exports, physically plausible PCB materials, camera/stage setup, and visual validation.
skills/schematic/SKILL.md · 102 lines · 11.1 KB
description: Create, modify, and review KiCad schematics from selected parts, datasheets, reference designs, renders, and netlists.
skills/symbol/SKILL.md · 19 lines · 1.9 KB
description: Generate compliant KiCad symbols
11 skills · 904 lines of SKILL.md · 70 files · 300.2 KB · 47 commits · rev 641ade9
Bundled with the KiCad library REAL electrical engineers use!
199 footprints · 5 libraries · 42 STEP models
F.Cu F.Mask F.SilkS F.Fab F.CrtYd

## Original Extract

A TOTALLY open-source workspace for agentic hardware development.

We made Backplane to actually use Fable, Astra, etc. as real electrical engineers! It incorporates KiCad + your favorite harness (Claude Code, Codex, OpenCode, Grok) to accelerate hardware development. You can use it "autonomously" or work alongside the agent, making changes or comments to redirect the agent as it works. Also has some slick skills for schematic editing, BOM management. Hardware is in its "tab to accept changes"/copilot era rather than its agentic era imo.

Source on GitHub ★ 1.1k
curl -fsSL https://backplaneworks.com/install.sh | sh Copy
backplane runs on YOUR compute (mobile apps coming soon)
Backplane uses AGENTS to design HARDWARE.
Backplane is FREE AND OPEN SOURCE FOREVER.
Backplane is created + maintained by real electrical engineers.
PROJECTS! THREADS! PERSISTENT BOTS (they're cats)
4 bots · 2 rooms · 6 screens · space_set bot_send room_post memory_save routine_set remind_person
Very nice tools that interact DIRECTLY with KiCad (no file rewriting, still human editable!)
KiStack built in! Automatic updates as skills improve!
skills/bom/SKILL.md · 78 lines · 6.2 KB
description: Resolve and review exact PCB parts, packages, electrical ratings, availability, and manufacturer BOM fields for KiCad projects.
skills/export/SKILL.md · 120 lines · 4.6 KB
description: Generate deterministic KiCad ERC, DRC, fabrication, assembly, documentation, and 3D export outputs with kicad-cli.
references/kicad_cli_command_fp_export_svg.md
references/kicad_cli_command_fp_upgrade.md
references/kicad_cli_command_introduction.md
references/kicad_cli_command_jobset_run.md
references/kicad_cli_command_pcb_drc.md
references/kicad_cli_command_pcb_export_3dpdf.md
references/kicad_cli_command_pcb_export_brep.md
references/kicad_cli_command_pcb_export_drill.md
references/kicad_cli_command_pcb_export_dxf.md
references/kicad_cli_command_pcb_export_gencad.md
references/kicad_cli_command_pcb_export_gerbers.md
references/kicad_cli_command_pcb_export_glb.md
references/kicad_cli_command_pcb_export_hpgl.md
references/kicad_cli_command_pcb_export_ipc2581.md
references/kicad_cli_command_pcb_export_ipcd356.md
references/kicad_cli_command_pcb_export_odb.md
references/kicad_cli_command_pcb_export_pdf.md
references/kicad_cli_command_pcb_export_ply.md
references/kicad_cli_command_pcb_export_pos.md
references/kicad_cli_command_pcb_export_ps.md
references/kicad_cli_command_pcb_export_stats.md
references/kicad_cli_command_pcb_export_step.md
references/kicad_cli_command_pcb_export_stl.md
references/kicad_cli_command_pcb_export_stpz.md
references/kicad_cli_command_pcb_export_svg.md
references/kicad_cli_command_pcb_export_u3d.md
references/kicad_cli_command_pcb_export_vrml.md
references/kicad_cli_command_pcb_export_xao.md
references/kicad_cli_command_pcb_import.md
references/kicad_cli_command_pcb_render.md
references/kicad_cli_command_pcb_upgrade.md
references/kicad_cli_command_sch_erc.md
references/kicad_cli_command_sch_export_bom.md
references/kicad_cli_command_sch_export_dxf.md
references/kicad_cli_command_sch_export_hpgl.md
references/kicad_cli_command_sch_export_netlist.md
references/kicad_cli_command_sch_export_pdf.md
references/kicad_cli_command_sch_export_ps.md
references/kicad_cli_command_sch_export_python_bom.md
references/kicad_cli_command_sch_export_svg.md
references/kicad_cli_command_sch_upgrade.md
references/kicad_cli_command_sym_export_svg.md
references/kicad_cli_command_sym_upgrade.md
references/kicad_cli_command_version.md
** ERC messages: 0 Errors 0 Warnings 0 tome_nano_100-drc.rpt ** Found 0 DRC violations **
** Found 0 unconnected pads **
** Found 0 Footprint errors ** positions_raw.csv Ref,Val,Package,PosX,PosY,Rot,Side
"C10","1uF","C_0402_1005Metric_L",150.913501,-131.187500,0.000000,bottom
"C11","1uF","C_0402_1005Metric_L",151.401001,-136.212500,0.000000,bottom positions.csv (scripts/convert_position.py) Designator,Mid X,Mid Y,Layer,Rotation
C10,150.913501,-131.187500,B,0.000000
C11,151.401001,-136.212500,B,0.000000 top.png (kicad-cli pcb render --side top) 22 files in gerbs/
skills/footprint/SKILL.md · 25 lines · 2.6 KB
description: Generate compliant KiCad footprints from official KiCad Library Tools
skills/gerbers/SKILL.md · 18 lines · 0.9 KB
description: Generate, render, and visually review KiCad Gerber layers as a fabrication-output inspection step.
skills/layout/SKILL.md · 23 lines · 4.3 KB
description: Place components and route KiCad PCB layouts with footprint checks, schematic pin swaps, and signal integrity best practices.
## Editing the active board in place
skills/layout-placement/SKILL.md · 344 lines · 21.5 KB
description: Place components on a KiCad board so each part sits where its circuit needs it, placement DRC passes, and every pad stays reachable, with IPC scripts that move parts and measure distances and routability. Covers reading the circuit, fixed parts, IC orientation, and decoupling, regulator, clock, RF and filter parts. Use for component placement; not for routing or schematic design.
## Save progress to the configured board
## 1. Read the circuit before moving anything
## 2. Fix the immovable, then orient the ICs
## 3. Place in order of how little distance each connection tolerates
## 4. Keep every pad reachable
230 legal candidates for C48 at U1.D1 (current gap 2.925 mm)
blocks gap_mm partner_mm x y rot faces blocked pins
0 2.867 9.000 97.600 95.050 90 True []
0 2.868 9.000 97.850 95.050 90 True []
0 2.884 9.000 98.100 95.050 90 True []
0 3.067 9.000 98.100 95.300 0 True []
0 3.069 9.000 98.350 95.300 0 True [] skills/panelize/SKILL.md · 90 lines · 3.4 KB
description: Use KiKit to panelize a given PCB
skills/pcb/SKILL.md · 32 lines · 2.5 KB
description: Review KiCad PCB layouts with DRC, layer renders, and 3D-model interference inspection.
## Editing the active board in place
skills/pcb-product-render/SKILL.md · 53 lines · 5.0 KB
description: Create realistic Blender product renders from KiCad PCB projects, including refreshed 3D exports, physically plausible PCB materials, camera/stage setup, and visual validation.
skills/schematic/SKILL.md · 102 lines · 11.1 KB
description: Create, modify, and review KiCad schematics from selected parts, datasheets, reference designs, renders, and netlists.
skills/symbol/SKILL.md · 19 lines · 1.9 KB
description: Generate compliant KiCad symbols
11 skills · 904 lines of SKILL.md · 70 files · 300.2 KB · 47 commits · rev 641ade9
Bundled with the KiCad library REAL electrical engineers use!
199 footprints · 5 libraries · 42 STEP models
F.Cu F.Mask F.SilkS F.Fab F.CrtYd
