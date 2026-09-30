---
source: "https://buildfreely.com/"
hn_url: "https://news.ycombinator.com/item?id=49902647"
title: "Build Freely: Non-AI Shed Plan Generator"
article_title: "Build Freely Non-AI Shed Plan Generator"
image: ""
author: "darkstar999"
captured_at: "2026-09-30T01:11:08Z"
capture_tool: "hn-digest"
hn_id: 49902647
score: 1
comments: 0
posted_at: "2026-09-30T00:11:39Z"
tags:
  - hacker-news
---

# Build Freely: Non-AI Shed Plan Generator

- HN: [49902647](https://news.ycombinator.com/item?id=49902647)
- Source: [buildfreely.com](https://buildfreely.com/)
- Score: 1
- Comments: 0
- Posted: 2026-09-30T00:11:39Z

## Translation

Title: Build Freely: Non-AI Shed Plan Generator
Article title: Build Freely Non-AI Shed Plan Generator

Article text:
It does not use AI. All code was written manually.
Currently, it can generate only single story, rectangular structures with either a shed, gable, or hip roof.
It is possible to generate designs which are unsafe and do not adhere to proper building codes. Review design with a engineer before building.
Video: Talk @PTWD
For feedback, email
←Dan at this domain.
BOM part plans sometimes display poorly rotated parts causing
difficultly in measuring lumber and often causing associated lumber
material to be "Unknown".
Openings above openings not fully supported -- framing intersects other window.
An opening with impossible header may prevent successful rendition.
Improves part edges in 3D client.
Adds hip roof support (still a work in progress).
Concrete foundations can now be generated with central girders to support long floor joists.
Gable roof design now generates lookouts.
Rafter tail max plumb height now working for gable roofs.
Fascia and fly rafter max plumb heights can now be specified.
Improved console messages when rendering.
Windows can now be stacked above each other.
Fixes bug in common stud placement and crippling when overlapping opening.
Add support for ridge beam posts in gable wall.
Improved BOM part grouping, so parts of similar shape are more likely to appear together in same part cut plan.
Shed roof now support walls with flat top plate and birdsmouth cut in rafters (instead of top plates sloped to match rafter angle).
Gable roofs now include optional collar ties. Ceiling joists renamed to rafter ties.
Sheathing now covers gable walls of gable roof.
Bug fixed regarding loading new design over current one, was resulting in messed up openings.
Opening X values now describe position of center of opening (previously was left edge of opening for non-percentage values).
Opening Z values can now be negative to express distance from top of wall to top of opening.
Maximum height above grade of structure is now reported (see top left corner of screen).
Enclosure dimensions now describe dimensions of exterior of framing (previously included sheathing).
Sheathing no longer laps sheathing of adjacent wall. (might add this back as an option later, would be nice especially for thicking sheating that includes insulation layer).
Better layout of sheathing and roof decking to ensure edges of wall and roof begin with large panel for greater strength.
Adds fancy incremental reveal of parts for initial load of 3D model.
Simplifies gable end wall, no longer employs top and bottom plates in triangular portion of wall.
New code for sheath staggering
Adds "skid" foundation type (aka girder only).
Fixes pier foundation bugs regarding cantilever, and now allows negative cantilever.
Fixes some lookout and subFascia bugs.
Fixes interference of top plates of adjacent walls.
Fixes errors in shed roof top plates where plates were generated too thin.
Adds foundation mudsill material max length specification so long mudsill boards will be split to smaller size.
Reduces mouse control zoom and rotate rate for more precise control.
Better BOM scaling, and dimension label positioning.
Adds concrete foundation type.
X Positioning of opening supports mutliples and percents.

## Original Extract

It does not use AI. All code was written manually.
Currently, it can generate only single story, rectangular structures with either a shed, gable, or hip roof.
It is possible to generate designs which are unsafe and do not adhere to proper building codes. Review design with a engineer before building.
Video: Talk @PTWD
For feedback, email
←Dan at this domain.
BOM part plans sometimes display poorly rotated parts causing
difficultly in measuring lumber and often causing associated lumber
material to be "Unknown".
Openings above openings not fully supported -- framing intersects other window.
An opening with impossible header may prevent successful rendition.
Improves part edges in 3D client.
Adds hip roof support (still a work in progress).
Concrete foundations can now be generated with central girders to support long floor joists.
Gable roof design now generates lookouts.
Rafter tail max plumb height now working for gable roofs.
Fascia and fly rafter max plumb heights can now be specified.
Improved console messages when rendering.
Windows can now be stacked above each other.
Fixes bug in common stud placement and crippling when overlapping opening.
Add support for ridge beam posts in gable wall.
Improved BOM part grouping, so parts of similar shape are more likely to appear together in same part cut plan.
Shed roof now support walls with flat top plate and birdsmouth cut in rafters (instead of top plates sloped to match rafter angle).
Gable roofs now include optional collar ties. Ceiling joists renamed to rafter ties.
Sheathing now covers gable walls of gable roof.
Bug fixed regarding loading new design over current one, was resulting in messed up openings.
Opening X values now describe position of center of opening (previously was left edge of opening for non-percentage values).
Opening Z values can now be negative to express distance from top of wall to top of opening.
Maximum height above grade of structure is now reported (see top left corner of screen).
Enclosure dimensions now describe dimensions of exterior of framing (previously included sheathing).
Sheathing no longer laps sheathing of adjacent wall. (might add this back as an option later, would be nice especially for thicking sheating that includes insulation layer).
Better layout of sheathing and roof decking to ensure edges of wall and roof begin with large panel for greater strength.
Adds fancy incremental reveal of parts for initial load of 3D model.
Simplifies gable end wall, no longer employs top and bottom plates in triangular portion of wall.
New code for sheath staggering
Adds "skid" foundation type (aka girder only).
Fixes pier foundation bugs regarding cantilever, and now allows negative cantilever.
Fixes some lookout and subFascia bugs.
Fixes interference of top plates of adjacent walls.
Fixes errors in shed roof top plates where plates were generated too thin.
Adds foundation mudsill material max length specification so long mudsill boards will be split to smaller size.
Reduces mouse control zoom and rotate rate for more precise control.
Better BOM scaling, and dimension label positioning.
Adds concrete foundation type.
X Positioning of opening supports mutliples and percents.
