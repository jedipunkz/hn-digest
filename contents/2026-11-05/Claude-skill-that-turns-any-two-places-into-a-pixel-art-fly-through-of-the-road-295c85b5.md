---
source: "https://github.com/jonwheatley/pixel-path"
hn_url: "https://news.ycombinator.com/item?id=50382882"
title: "Claude skill that turns any two places into a pixel-art fly-through of the road"
article_title: "GitHub - jonwheatley/pixel-path: Turn any two places into a pixel-art fly-through of the real road. A Claude Code skill. · GitHub"
image: "https://repository-images.githubusercontent.com/1415902477/811fd8c2-ef97-4650-971d-56562d872424"
author: "tfinch"
captured_at: "2026-11-05T19:36:02Z"
capture_tool: "hn-digest"
hn_id: 50382882
score: 2
comments: 0
posted_at: "2026-11-05T18:54:05Z"
tags:
  - hacker-news
---

# Claude skill that turns any two places into a pixel-art fly-through of the road

- HN: [50382882](https://news.ycombinator.com/item?id=50382882)
- Source: [github.com](https://github.com/jonwheatley/pixel-path)
- Score: 2
- Comments: 0
- Posted: 2026-11-05T18:54:05Z

## Translation

Title: Claude skill that turns any two places into a pixel-art fly-through of the road
Article title: GitHub - jonwheatley/pixel-path: Turn any two places into a pixel-art fly-through of the real road. A Claude Code skill. · GitHub
Description: Turn any two places into a pixel-art fly-through of the real road. A Claude Code skill. - jonwheatley/pixel-path

Article text:
GitHub - jonwheatley/pixel-path: Turn any two places into a pixel-art fly-through of the real road. A Claude Code skill. · GitHub
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
jonwheatley
pixel-path Public Notifications You must be signed in to change notification settings
Star 302 ( 302 ) You must be signed in to star a repository
Turn any two places into a pixel-art fly-through of the real road. A Claude Code skill.
22 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.claude-plugin .claude-plugin docs docs scripts scripts templates templates LICENSE LICENSE README.md README.md SKILL.md SKILL.md View all files Repository files navigation
A Claude Code skill that turns any two places into a pixel-art fly-through of the real road. Ask Claude for "a pixel
path from Bath to Bristol" and it:
bakes the real buildings, land cover, landmarks, hills and horizon along it from OpenStreetMap and elevation data
generates scenery art for the region with RetroDiffusion (affiliate
link)
checks its own work with screenshots
hands you a page where you zoom along the route. Scroll to go faster, slower or backwards.
Bath, England: real streets from OpenStreetMap, art generated with the moderate kit.
A first-person, pixel-art view of any real road, built from the real map. Useful wherever you want a real place or route
on screen:
Games: a real road as a level for a driving, racing, idle or endless-runner game, instead of a made-up one.
Showing a journey: a commute, a delivery round, a marathon or cycling stage, a road trip you took, a move between cities.
Content: trip recaps, travel posts, video intros, stream backgrounds and lo-fi loops.
Places and history: Route 66, a pilgrimage route, the road in a news story, a preview of a neighbourhood.
Learning: a commented pseudo-3D renderer and an OpenStreetMap bake pipeline to read or build on.
A folder with an index.html . Open it through any static web server and you are moving down the real road:
Buildings on their real footprints. Every building that faces the road is a block on its real line, with the
kind of picture its type and height call for. A long frontage becomes a row of different buildings, and cracks and
short lots are closed.
More of the real place. A second row of buildings shows across squares and car parks, tall buildings make a
skyline, and landmarks stand at their real places.
The real ground. Fields, forest, water, parking lots and sidewalks are drawn per pixel from the map. The hills
are real, and so is the horizon, marched over a 61 × 61 km height grid with the Earth's curvature.
Signs drawn by code, so the names are exact: town names, distances, street names, brown landmark signs and
railway crossings.
No interface apart from a small credits line:
Scroll (or drag on a touch screen) to speed up from a crawl to very fast. Keep going the other way to slow down
and reverse.
The arrow keys do the same, and Space pauses.
I shows where you are and M shows a small map.
Before anything slow happens, the first step prints an estimate for your route, and the skill asks you to choose.
Only the art costs money: $0.18 a picture at RetroDiffusion's rd_pro__default (October 2026). Cost depends on the
number of regions (distinct looks along the route), not on distance.
The levels are nested, so moving up a level only pays for the new pictures. The exact quote is free and comes after
the bake. Allow about 25% on top for re-rolls: in testing, 10% (Bath) to 34% (Manhattan) of pictures came back cut off
or on a background. The cleanup dropped them, and the self-check re-rolled them within the budget.
Bath to Bristol at moderate (two regions, Georgian Bath, then Victorian Keynsham and Bristol, 32 pictures, $5.76):
Kyoto Station to Kinkaku-ji, Japan, at sparse (one region, four pictures, $0.72):
Both runs passed the full self-check, including the sweep: 627 and 345 moving frames, with no sky through or between
buildings. The estimate printed after routing matched the exact quote in every test run.
/plugin marketplace add jonwheatley/pixel-path
/plugin install pixel-path@pixel-path
By hand. Copy this folder into your skills folder, so that the skill is at
~/.claude/skills/pixel-path/SKILL.md . For a single project, use .claude/skills/pixel-path/ inside it.
Then ask Claude: "make me a pixel path from Paris to Versailles".
Node 18 or newer. No npm packages are needed to build a route.
A RetroDiffusion account with credit, for real art: https://www.retrodiffusion.ai?ref=YBQJM823 . This is an
affiliate link (the author of this skill gets a referral).
Create an API key and set it in your own terminal with export RETRODIFFUSION_API_KEY=... .
On a Mac you can store it in the Keychain instead: security add-generic-password -s pixel-path -a retrodiffusion -w .
Never paste the key into a chat. The skill never asks for it and never prints it.
Without an account, placeholders cost nothing.
For the self-check's screenshots: Node 20 or newer with Playwright and Chromium ( npm i -D playwright && npx playwright install chromium ). Without them, the layout and art checks still run.
Network access while baking to Nominatim, OSRM, Overpass, AWS Terrain Tiles and Wikidata. Baking takes about 20
to 60 seconds per 10 km.
The skill doesn't call a route done when the code runs. It calls it done when the pictures pass:
Every building is a solid. Each building is laid along its real line and always gets a shaded end wall, so it
never reads as a flat card. A neighbour that joins it paints over the part it covers.
In a dense street, a building mapped far back is pulled up to the sidewalk.
A backing building stands behind every gap in a built-up row.
A second row of real buildings shows across squares and car parks.
Cleanup. Each picture is cleaned and its backdrop flood-filled away. Holes in facades are filled, including a
roof-to-ground gap between two houses drawn in one picture. Pictures that still carry a background, painted sky, a
cut-off edge or a sliver are dropped. A contact sheet of every picture is written for Claude to look at.
Layout check. The whole route is laid out exactly as the page will draw it, checking for gaps in street walls
and for a building with the same picture as its neighbour.
The sweep. Every 75 m of built-up road, the camera moves and looks straight ahead and 30° to each side. The sky
is painted magenta and every building is drawn into a mask, and it counts:
slits of sky between two parts of the same street wall that reach down to the ground
The 20 worst frames are written side by side with the problem marked.
Review. Claude is told to look at the worst frames and the samples and keep fixing until the sweep is clean.
The page shows a credits line that must stay visible wherever a fly-through is shown.
Map data © OpenStreetMap contributors , under the
Open Database Licence . A route's data/ folder is a derived database: if
you share it publicly, it stays under the ODbL, with credit.
Routing: OSRM , from OpenStreetMap data. The public servers are for light use.
Geocoding: Nominatim , under the OSM Foundation's usage policy.
Elevation: Terrain Tiles (Mapzen/Tilezen on AWS Open Data), from
SRTM, GMTED2010, ETOPO1, NED and others
( attribution ).
Art: generated by you with RetroDiffusion (affiliate link), under
RetroDiffusion's terms of service. The pictures in this README were generated with this skill.
Code: MIT (see LICENSE ), by Jon Wheatley. The road projection is the classic arcade technique (see Lou's Pseudo
3D Page and Jake Gordon's Javascript Racer).
OpenStreetMap coverage decides the street. Where buildings aren't mapped, plausible ones are filled in by land
use (houses on lawns, shops on paved land, a row behind in dense towns). Heights are often missing and are guessed
from building type.
Open spaces look open. Squares, junctions, car parks and parks read as wide paving or grass with buildings across
them, because that is what is there. Some frames in a city centre show more paving than you would expect.
The art is the model's. The cleanup catches backgrounds, painted sky, cut-off pictures and slivers, but not
everything: a building drawn at an angle, for example. Claude is told to look at every picture and re-roll bad ones,
at the cost of a picture each.
One look per region. Every town in a region uses the same kit, and long streets show the same pictures again
further along. Split a long route into more regions for more variety, at more cost.
Landmarks are pictures, not models. A landmark faces the road and does not turn as you pass. You need to
describe what it looks like (the model does not know buildings by name), and a landmark just behind the start point
is not shown.
Overpass is the slow, flaky step. The public servers time out under load. The bake retries across mirrors and
resumes, but a long route can take an hour on a bad day.
Driving routes only, in one afternoon light. Routes come from OSRM's car profile. There is no traffic, no people
and no weather.
Turn any two places into a pixel-art fly-through of the real road. A Claude Code skill.
22 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Turn any two places into a pixel-art fly-through of the real road. A Claude Code skill. - jonwheatley/pixel-path

GitHub - jonwheatley/pixel-path: Turn any two places into a pixel-art fly-through of the real road. A Claude Code skill. · GitHub
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
jonwheatley
pixel-path Public Notifications You must be signed in to change notification settings
Star 302 ( 302 ) You must be signed in to star a repository
Turn any two places into a pixel-art fly-through of the real road. A Claude Code skill.
22 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.claude-plugin .claude-plugin docs docs scripts scripts templates templates LICENSE LICENSE README.md README.md SKILL.md SKILL.md View all files Repository files navigation
A Claude Code skill that turns any two places into a pixel-art fly-through of the real road. Ask Claude for "a pixel
path from Bath to Bristol" and it:
bakes the real buildings, land cover, landmarks, hills and horizon along it from OpenStreetMap and elevation data
generates scenery art for the region with RetroDiffusion (affiliate
link)
checks its own work with screenshots
hands you a page where you zoom along the route. Scroll to go faster, slower or backwards.
Bath, England: real streets from OpenStreetMap, art generated with the moderate kit.
A first-person, pixel-art view of any real road, built from the real map. Useful wherever you want a real place or route
on screen:
Games: a real road as a level for a driving, racing, idle or endless-runner game, instead of a made-up one.
Showing a journey: a commute, a delivery round, a marathon or cycling stage, a road trip you took, a move between cities.
Content: trip recaps, travel posts, video intros, stream backgrounds and lo-fi loops.
Places and history: Route 66, a pilgrimage route, the road in a news story, a preview of a neighbourhood.
Learning: a commented pseudo-3D renderer and an OpenStreetMap bake pipeline to read or build on.
A folder with an index.html . Open it through any static web server and you are moving down the real road:
Buildings on their real footprints. Every building that faces the road is a block on its real line, with the
kind of picture its type and height call for. A long frontage becomes a row of different buildings, and cracks and
short lots are closed.
More of the real place. A second row of buildings shows across squares and car parks, tall buildings make a
skyline, and landmarks stand at their real places.
The real ground. Fields, forest, water, parking lots and sidewalks are drawn per pixel from the map. The hills
are real, and so is the horizon, marched over a 61 × 61 km height grid with the Earth's curvature.
Signs drawn by code, so the names are exact: town names, distances, street names, brown landmark signs and
railway crossings.
No interface apart from a small credits line:
Scroll (or drag on a touch screen) to speed up from a crawl to very fast. Keep going the other way to slow down
and reverse.
The arrow keys do the same, and Space pauses.
I shows where you are and M shows a small map.
Before anything slow happens, the first step prints an estimate for your route, and the skill asks you to choose.
Only the art costs money: $0.18 a picture at RetroDiffusion's rd_pro__default (October 2026). Cost depends on the
number of regions (distinct looks along the route), not on distance.
The levels are nested, so moving up a level only pays for the new pictures. The exact quote is free and comes after
the bake. Allow about 25% on top for re-rolls: in testing, 10% (Bath) to 34% (Manhattan) of pictures came back cut off
or on a background. The cleanup dropped them, and the self-check re-rolled them within the budget.
Bath to Bristol at moderate (two regions, Georgian Bath, then Victorian Keynsham and Bristol, 32 pictures, $5.76):
Kyoto Station to Kinkaku-ji, Japan, at sparse (one region, four pictures, $0.72):
Both runs passed the full self-check, including the sweep: 627 and 345 moving frames, with no sky through or between
buildings. The estimate printed after routing matched the exact quote in every test run.
/plugin marketplace add jonwheatley/pixel-path
/plugin install pixel-path@pixel-path
By hand. Copy this folder into your skills folder, so that the skill is at
~/.claude/skills/pixel-path/SKILL.md . For a single project, use .claude/skills/pixel-path/ inside it.
Then ask Claude: "make me a pixel path from Paris to Versailles".
Node 18 or newer. No npm packages are needed to build a route.
A RetroDiffusion account with credit, for real art: https://www.retrodiffusion.ai?ref=YBQJM823 . This is an
affiliate link (the author of this skill gets a referral).
Create an API key and set it in your own terminal with export RETRODIFFUSION_API_KEY=... .
On a Mac you can store it in the Keychain instead: security add-generic-password -s pixel-path -a retrodiffusion -w .
Never paste the key into a chat. The skill never asks for it and never prints it.
Without an account, placeholders cost nothing.
For the self-check's screenshots: Node 20 or newer with Playwright and Chromium ( npm i -D playwright && npx playwright install chromium ). Without them, the layout and art checks still run.
Network access while baking to Nominatim, OSRM, Overpass, AWS Terrain Tiles and Wikidata. Baking takes about 20
to 60 seconds per 10 km.
The skill doesn't call a route done when the code runs. It calls it done when the pictures pass:
Every building is a solid. Each building is laid along its real line and always gets a shaded end wall, so it
never reads as a flat card. A neighbour that joins it paints over the part it covers.
In a dense street, a building mapped far back is pulled up to the sidewalk.
A backing building stands behind every gap in a built-up row.
A second row of real buildings shows across squares and car parks.
Cleanup. Each picture is cleaned and its backdrop flood-filled away. Holes in facades are filled, including a
roof-to-ground gap between two houses drawn in one picture. Pictures that still carry a background, painted sky, a
cut-off edge or a sliver are dropped. A contact sheet of every picture is written for Claude to look at.
Layout check. The whole route is laid out exactly as the page will draw it, checking for gaps in street walls
and for a building with the same picture as its neighbour.
The sweep. Every 75 m of built-up road, the camera moves and looks straight ahead and 30° to each side. The sky
is painted magenta and every building is drawn into a mask, and it counts:
slits of sky between two parts of the same street wall that reach down to the ground
The 20 worst frames are written side by side with the problem marked.
Review. Claude is told to look at the worst frames and the samples and keep fixing until the sweep is clean.
The page shows a credits line that must stay visible wherever a fly-through is shown.
Map data © OpenStreetMap contributors , under the
Open Database Licence . A route's data/ folder is a derived database: if
you share it publicly, it stays under the ODbL, with credit.
Routing: OSRM , from OpenStreetMap data. The public servers are for light use.
Geocoding: Nominatim , under the OSM Foundation's usage policy.
Elevation: Terrain Tiles (Mapzen/Tilezen on AWS Open Data), from
SRTM, GMTED2010, ETOPO1, NED and others
( attribution ).
Art: generated by you with RetroDiffusion (affiliate link), under
RetroDiffusion's terms of service. The pictures in this README were generated with this skill.
Code: MIT (see LICENSE ), by Jon Wheatley. The road projection is the classic arcade technique (see Lou's Pseudo
3D Page and Jake Gordon's Javascript Racer).
OpenStreetMap coverage decides the street. Where buildings aren't mapped, plausible ones are filled in by land
use (houses on lawns, shops on paved land, a row behind in dense towns). Heights are often missing and are guessed
from building type.
Open spaces look open. Squares, junctions, car parks and parks read as wide paving or grass with buildings across
them, because that is what is there. Some frames in a city centre show more paving than you would expect.
The art is the model's. The cleanup catches backgrounds, painted sky, cut-off pictures and slivers, but not
everything: a building drawn at an angle, for example. Claude is told to look at every picture and re-roll bad ones,
at the cost of a picture each.
One look per region. Every town in a region uses the same kit, and long streets show the same pictures again
further along. Split a long route into more regions for more variety, at more cost.
Landmarks are pictures, not models. A landmark faces the road and does not turn as you pass. You need to
describe what it looks like (the model does not know buildings by name), and a landmark just behind the start point
is not shown.
Overpass is the slow, flaky step. The public servers time out under load. The bake retries across mirrors and
resumes, but a long route can take an hour on a bad day.
Driving routes only, in one afternoon light. Routes come from OSRM's car profile. There is no traffic, no people
and no weather.
Turn any two places into a pixel-art fly-through of the real road. A Claude Code skill.
22 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
