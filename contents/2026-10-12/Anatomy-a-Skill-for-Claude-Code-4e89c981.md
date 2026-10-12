---
source: "https://skills.wheresryan.sh/anatomy"
hn_url: "https://news.ycombinator.com/item?id=50053620"
title: "Anatomy · a Skill for Claude Code"
article_title: "Anatomy · A skill for Claude Code"
image: "https://skills.wheresryan.sh/anatomy/og.png"
author: "handfuloflight"
captured_at: "2026-10-12T10:20:05Z"
capture_tool: "hn-digest"
hn_id: 50053620
score: 1
comments: 0
posted_at: "2026-10-12T09:20:21Z"
tags:
  - hacker-news
---

# Anatomy · a Skill for Claude Code

- HN: [50053620](https://news.ycombinator.com/item?id=50053620)
- Source: [skills.wheresryan.sh](https://skills.wheresryan.sh/anatomy)
- Score: 1
- Comments: 0
- Posted: 2026-10-12T09:20:21Z

## Translation

Title: Anatomy · a Skill for Claude Code
Article title: Anatomy · A skill for Claude Code
Description: A Claude skill that draws an idea as the real machine that does it: a live isometric figure you can turn, press and read.

Article text:
Anatomy · A skill for Claude Code Anatomy
Anatomy · a skill for Claude Code
Name an idea in one line, and Claude draws the machine that does it.
Most explanations of how something works are drawn as boxes and arrows. They tell you the parts exist, but not what the parts do, and the reader has to imagine the motion the arrows stand for. I wanted Claude to draw the thing itself : a real object whose moving parts behave the way the idea does.
Anatomy is a skill for Claude Code that does this. You name an idea in one line, Claude invents a machine that does it, and draws it as a live isometric figure you can turn, press and read. Every figure is made of dozens to hundreds of small solid parts, and every number it shows comes from the same model that moves them .
This is a mechanical watch movement. It came from one line, /anatomy explain how a mechanical watch keeps time, interactive , and just under two hours of Claude working on its own. 1 1 One hour and fifty-four minutes of active time, built in one pass and checked by the same agent.
Drag the crown upward to wind it: the ratchet turns, the click ticks and the power reserve climbs. Press S to slow time to a twentieth, and again to a hundredth, and you can watch the pallet fork let the fifteen-tooth escape wheel through half a tooth at a time . Press E to lift the bridges apart and see the train underneath, its wheels cut to true ratios: the fourth wheel turns once a minute and the centre wheel once an hour.
The skill rests on one rule, written at the top of its instructions, above everything else:
The object is the explanation. SKILL.md, “The one idea that matters”
No boxes, no arrows, no labels floating in space. Claude has to invent a real, well-made thing whose working is the concept, and say in one sentence what moves and what that motion means.
A figure fails in two ways, and I watched it fail both ways many times before the rule existed. Too abstract: stacked plates, floating slabs and labelled boxes that don’t read as anything. Too flat: one big slab with a few lines on it, with no volume, no parts and nothing to discover. The cure for both is the same, a recognisable machine built from many small parts that each do something. 2 2 The machine is never copied from a reference picture. A reference can give the finish, the materials and the mood; the object is invented.
Some ideas map almost by themselves. A queue is a marble run: two hoppers feed one sloping lane, and the marbles wait in order behind an escapement until a timer drum takes them. A cache that forgets its oldest entry is a bookcase with four places on top, most recently used at the left. Others take more invention. A bloom filter is a fruit machine whose three reels are the three hash functions, and it only pays out when all three of a word’s flags are already up.
None of these were drawn by hand. Each one started from a single line naming the idea. There are 33 so far, all live in the collection .
The prompt is the idea , not the drawing. I almost never say what to draw. “Explain how an LRU cache evicts entries” is enough, and deciding that it should be a bookcase is the part Claude is asked to do.
Three words change what you get, and Claude only reaches for them when you write them in the line :
If you already know the object, name it, with whatever true numbers matter. And if the figure is going somewhere in particular, say where. The skill writes framework-free TypeScript that outputs SVG and HTML, or React components for a Next.js site, in a dark or a light theme.
Name the idea /anatomy explain how an LRU cache evicts entries, interactive
Add shaders /anatomy explain how a jet engine afterburner works, interactive , WebGL
Make it turn /anatomy explain why the Earth has seasons, interactive , 3D
Or name the object /anatomy a mechanical metronome on a small wooden plinth, swinging at 120 bpm, with a sliding weight you can drag, interactive
Say where it goes /anatomy explain how a CDN edge cache serves a request, interactive , as React components for our Next.js docs, dark theme
You don’t always need the command. The skill’s description covers any request for an explanatory diagram, a technical illustration or an interactive explainer, even one that never says “isometric”, so Claude reaches for it on its own. 3 3 Until a request matches, Claude sees only the skill’s name and its description. It reads the rest only when it needs it.
Between your line and the finished figure, Claude works through a fixed order. Most of the quality is decided in the first four steps , before any geometry exists.
Size it. Say up front whether this is a Figure, a Hero or an Epic, and roughly how long it will take.
Truth first. Collect the real numbers and behaviour, and write them as a small model that the motion and the readout both read.
Invent the object. A machine whose mechanism maps one to one onto the idea.
List the parts. At least forty, grouped into base, structure, mechanism, controls and the one lit subject. Each sits on another.
Plan the world. World units, plans as rounded rectangles and circles, a camera, then a fit to the card.
Build the geometry once. Every static path computed ahead of time. Moving parts are built at rest and moved.
Paint back to front. Sorted by depth, wrapping parts split in two, and the order proved by an audit once it gets dense.
Frame it. A card with a title, the interaction hint and a live readout. Words stay out of the drawing.
Make it live. Calm, time-based motion, an idle tour, pointer and keyboard, and stillness under reduced motion.
Verify like a critic. Screenshots, contact sheets, crops at four times and two automatic checks, until nothing is left to fix.
Step two looks like this in practice. The test rig, one of the skill’s examples, explains how a glass button swells and leans under a finger, and its whole behaviour lives in a small spring model. The parts, the readout, the screen-reader text and the caption all read from it, so none of them can disagree.
function spring(held: Spring, want: number, dt: number, [stiff, drag]: Tuning, calm: boolean) {
const damping = calm ? 2 * Math.sqrt(stiff) : drag;
const steps = Math. max (Math.ceil(dt * SUB_RATE), 1);
const slice = dt / steps;
for ( let step = 0; step < steps; step++) {
held.rate += (want - held.at) * stiff * slice - held.rate * damping * slice;
held.at += held.rate * slice;
}
if (Math.abs(want - held.at) < QUIET && Math.abs(held.rate) < QUIET) {
held.at = want;
held.rate = 0;
}
} And step six. The dial gauge, the smallest example, starts with its base: four turned feet in depth order, a plate with a bevel, a groove inset five units from the edge, and a screw with a ring at each corner. Every one of those is a solid from the kit, not a line drawn by hand.
function plate() {
const out = [ `<path class="iso-halo" d="${k.haloOf(PLATE, -PLATE_H - FOOT_H, PLATE_H, P)}"/>` ];
for ( const [x, y] of byDepth(k.corners(PLATE, 22))) {
out.push(S(k.cylinder(x, y, 7, -PLATE_H - FOOT_H, 1.6, P, 28), { tone: "lo" }));
out.push(S(k.cylinder(x, y, 5.4, -PLATE_H - FOOT_H + 1.6, FOOT_H - 1.6, P, 28), { tone: "mid" }));
}
out.push(S(k.slabOf(PLATE, -PLATE_H, PLATE_H, P, 8, 1.6), { tone: "mid" }));
out.push(L(k.planOutline(k.insetPlan(PLATE, 5), 0, P), { tone: "lo" }));
const screws = k.corners(PLATE, 10);
out.push(L(screws.map(([x, y]) => k.ring(x, y, 2.2, 0, P, 16)).join( "" ), { tone: "lo" }));
out.push(D(screws.map(([x, y]) => at([x, y, 0])), { size: 0.5 }));
return out.join( "" );
} The first step sets the budget. Every request lands in one of three sizes:
A Figure is built without asking. A Hero is announced first, then built. An Epic is never started without you : Claude gives the estimate, proposes a Hero that keeps the core of the idea, and lists what the full version would add. 4 4 Whatever the size, there is a working page in the browser within the first hour. The rest of the time goes into deepening it.
Every figure follows the same drawing rules. They are what make it read as a solid, made thing rather than a sketch.
A figure is also a small program . It plays an idle tour when nobody touches it, pauses when you hover or focus it, answers the pointer and the keyboard, stops when it scrolls off screen and goes still under reduced motion. Its readout is announced to screen readers as it changes.
On a page of its own, a figure sits in a card with its number, title, hint and live readout in the four corners. On this site I moved the hint and the readout into the margin beside each figure, so the drawing stands on the page by itself.
The collection reads best one machine at a time. Each of these names the idea, then the object Claude chose for it; open one to play with it.
A hash map files each key in a bucket, and keys that land in the same bucket pile up. A post-office sorting desk. A coding machine prints each letter's bucket, the hash of its name mod 8, and files it in that pigeonhole, on top of any letters already there.
A bloom filter can say “definitely not” for certain, but only “maybe” for yes. A fruit machine. Three reels are the three hash functions and a row of sixteen flags is the bit array; it pays out only when all three of a word's flags are already up.
Two writers reading the same value at once will both act on a stale number. A water tank filled by two pumps. Without the interlock key both controllers read the same level and overfill it; with it, only the one holding the key may read and run.
Before any data moves, both ends have to agree that the other one is listening. A mail ropeway between two stations. Canisters flying SYN, SYN with an ACK pennant, then ACK ride the rope across, and each station's semaphore shows its connection state.
Temperature reshapes the odds of the next token without changing which tokens are possible. A roulette wheel whose sixteen pockets are the likeliest next tokens, each as wide as its probability. A thermostat knob sets the temperature and moves the frets.
Debounce waits for quiet; throttle lets one call through per window, however many arrive. A music-box recorder with a turning pin disc. Every click of the knob is a call; the throttle and debounce dashpots each set their own pins, so you can read when each function actually ran.
The pattern is the same every time. The object has to do the idea, not stand for it, so the thing you watch move is the thing being explained.
Small ideas make small machines. The big ones are where the skill earns its keep: hundreds of parts, real engineering numbers, and shaders that live inside the drawing rather than on top of it.
A Rolls-Royce Trent 970 , the engine of the A380, hung from a ground-run thrust frame with its near upper quarter cut away. Drag the lever and the fan, intermediate and high-pressure spools wind up at their own speeds. The fan swallows 1.2 tonnes of air a second, most of it leaves as a cold bypass jet while the rest burns in the combustor, and the load cell reads the thrust climbing from 14 to 334 kN . Press Space for a take-off run. 5 5 The audit measured 6,320 overlapping pairs across 541 solids and found every one drawn in the right order.
Raptor 2 , SpaceX’s full-flow staged-combustion engine, hanging in a test stand. It is one of the skill’s own examples and its reference for dense pipework. Hover it and it comes apart into eleven assemblies, from the gimbal to the nozzle; press I to fire it. Every drop of oxygen passes through an oxygen-rich preburner that spins the oxygen pump, every drop of methane through a fuel-rich one that spins the methane pump, and the two hot gases meet in a main chamber at 300 bar. 6 6 The plume, shock, splash and steam are WebGL, drawn between a back and a front SVG, so the stand passes in front of the flame.
The first working version took 38 minutes. Making it right took ten more hours, and an audit now proves that no pipe passes through anything, at rest, apart and in every frame between.
Most figures are pure SVG. When an idea needs fire,

[truncated]

## Original Extract

A Claude skill that draws an idea as the real machine that does it: a live isometric figure you can turn, press and read.

Anatomy · A skill for Claude Code Anatomy
Anatomy · a skill for Claude Code
Name an idea in one line, and Claude draws the machine that does it.
Most explanations of how something works are drawn as boxes and arrows. They tell you the parts exist, but not what the parts do, and the reader has to imagine the motion the arrows stand for. I wanted Claude to draw the thing itself : a real object whose moving parts behave the way the idea does.
Anatomy is a skill for Claude Code that does this. You name an idea in one line, Claude invents a machine that does it, and draws it as a live isometric figure you can turn, press and read. Every figure is made of dozens to hundreds of small solid parts, and every number it shows comes from the same model that moves them .
This is a mechanical watch movement. It came from one line, /anatomy explain how a mechanical watch keeps time, interactive , and just under two hours of Claude working on its own. 1 1 One hour and fifty-four minutes of active time, built in one pass and checked by the same agent.
Drag the crown upward to wind it: the ratchet turns, the click ticks and the power reserve climbs. Press S to slow time to a twentieth, and again to a hundredth, and you can watch the pallet fork let the fifteen-tooth escape wheel through half a tooth at a time . Press E to lift the bridges apart and see the train underneath, its wheels cut to true ratios: the fourth wheel turns once a minute and the centre wheel once an hour.
The skill rests on one rule, written at the top of its instructions, above everything else:
The object is the explanation. SKILL.md, “The one idea that matters”
No boxes, no arrows, no labels floating in space. Claude has to invent a real, well-made thing whose working is the concept, and say in one sentence what moves and what that motion means.
A figure fails in two ways, and I watched it fail both ways many times before the rule existed. Too abstract: stacked plates, floating slabs and labelled boxes that don’t read as anything. Too flat: one big slab with a few lines on it, with no volume, no parts and nothing to discover. The cure for both is the same, a recognisable machine built from many small parts that each do something. 2 2 The machine is never copied from a reference picture. A reference can give the finish, the materials and the mood; the object is invented.
Some ideas map almost by themselves. A queue is a marble run: two hoppers feed one sloping lane, and the marbles wait in order behind an escapement until a timer drum takes them. A cache that forgets its oldest entry is a bookcase with four places on top, most recently used at the left. Others take more invention. A bloom filter is a fruit machine whose three reels are the three hash functions, and it only pays out when all three of a word’s flags are already up.
None of these were drawn by hand. Each one started from a single line naming the idea. There are 33 so far, all live in the collection .
The prompt is the idea , not the drawing. I almost never say what to draw. “Explain how an LRU cache evicts entries” is enough, and deciding that it should be a bookcase is the part Claude is asked to do.
Three words change what you get, and Claude only reaches for them when you write them in the line :
If you already know the object, name it, with whatever true numbers matter. And if the figure is going somewhere in particular, say where. The skill writes framework-free TypeScript that outputs SVG and HTML, or React components for a Next.js site, in a dark or a light theme.
Name the idea /anatomy explain how an LRU cache evicts entries, interactive
Add shaders /anatomy explain how a jet engine afterburner works, interactive , WebGL
Make it turn /anatomy explain why the Earth has seasons, interactive , 3D
Or name the object /anatomy a mechanical metronome on a small wooden plinth, swinging at 120 bpm, with a sliding weight you can drag, interactive
Say where it goes /anatomy explain how a CDN edge cache serves a request, interactive , as React components for our Next.js docs, dark theme
You don’t always need the command. The skill’s description covers any request for an explanatory diagram, a technical illustration or an interactive explainer, even one that never says “isometric”, so Claude reaches for it on its own. 3 3 Until a request matches, Claude sees only the skill’s name and its description. It reads the rest only when it needs it.
Between your line and the finished figure, Claude works through a fixed order. Most of the quality is decided in the first four steps , before any geometry exists.
Size it. Say up front whether this is a Figure, a Hero or an Epic, and roughly how long it will take.
Truth first. Collect the real numbers and behaviour, and write them as a small model that the motion and the readout both read.
Invent the object. A machine whose mechanism maps one to one onto the idea.
List the parts. At least forty, grouped into base, structure, mechanism, controls and the one lit subject. Each sits on another.
Plan the world. World units, plans as rounded rectangles and circles, a camera, then a fit to the card.
Build the geometry once. Every static path computed ahead of time. Moving parts are built at rest and moved.
Paint back to front. Sorted by depth, wrapping parts split in two, and the order proved by an audit once it gets dense.
Frame it. A card with a title, the interaction hint and a live readout. Words stay out of the drawing.
Make it live. Calm, time-based motion, an idle tour, pointer and keyboard, and stillness under reduced motion.
Verify like a critic. Screenshots, contact sheets, crops at four times and two automatic checks, until nothing is left to fix.
Step two looks like this in practice. The test rig, one of the skill’s examples, explains how a glass button swells and leans under a finger, and its whole behaviour lives in a small spring model. The parts, the readout, the screen-reader text and the caption all read from it, so none of them can disagree.
function spring(held: Spring, want: number, dt: number, [stiff, drag]: Tuning, calm: boolean) {
const damping = calm ? 2 * Math.sqrt(stiff) : drag;
const steps = Math. max (Math.ceil(dt * SUB_RATE), 1);
const slice = dt / steps;
for ( let step = 0; step < steps; step++) {
held.rate += (want - held.at) * stiff * slice - held.rate * damping * slice;
held.at += held.rate * slice;
}
if (Math.abs(want - held.at) < QUIET && Math.abs(held.rate) < QUIET) {
held.at = want;
held.rate = 0;
}
} And step six. The dial gauge, the smallest example, starts with its base: four turned feet in depth order, a plate with a bevel, a groove inset five units from the edge, and a screw with a ring at each corner. Every one of those is a solid from the kit, not a line drawn by hand.
function plate() {
const out = [ `<path class="iso-halo" d="${k.haloOf(PLATE, -PLATE_H - FOOT_H, PLATE_H, P)}"/>` ];
for ( const [x, y] of byDepth(k.corners(PLATE, 22))) {
out.push(S(k.cylinder(x, y, 7, -PLATE_H - FOOT_H, 1.6, P, 28), { tone: "lo" }));
out.push(S(k.cylinder(x, y, 5.4, -PLATE_H - FOOT_H + 1.6, FOOT_H - 1.6, P, 28), { tone: "mid" }));
}
out.push(S(k.slabOf(PLATE, -PLATE_H, PLATE_H, P, 8, 1.6), { tone: "mid" }));
out.push(L(k.planOutline(k.insetPlan(PLATE, 5), 0, P), { tone: "lo" }));
const screws = k.corners(PLATE, 10);
out.push(L(screws.map(([x, y]) => k.ring(x, y, 2.2, 0, P, 16)).join( "" ), { tone: "lo" }));
out.push(D(screws.map(([x, y]) => at([x, y, 0])), { size: 0.5 }));
return out.join( "" );
} The first step sets the budget. Every request lands in one of three sizes:
A Figure is built without asking. A Hero is announced first, then built. An Epic is never started without you : Claude gives the estimate, proposes a Hero that keeps the core of the idea, and lists what the full version would add. 4 4 Whatever the size, there is a working page in the browser within the first hour. The rest of the time goes into deepening it.
Every figure follows the same drawing rules. They are what make it read as a solid, made thing rather than a sketch.
A figure is also a small program . It plays an idle tour when nobody touches it, pauses when you hover or focus it, answers the pointer and the keyboard, stops when it scrolls off screen and goes still under reduced motion. Its readout is announced to screen readers as it changes.
On a page of its own, a figure sits in a card with its number, title, hint and live readout in the four corners. On this site I moved the hint and the readout into the margin beside each figure, so the drawing stands on the page by itself.
The collection reads best one machine at a time. Each of these names the idea, then the object Claude chose for it; open one to play with it.
A hash map files each key in a bucket, and keys that land in the same bucket pile up. A post-office sorting desk. A coding machine prints each letter's bucket, the hash of its name mod 8, and files it in that pigeonhole, on top of any letters already there.
A bloom filter can say “definitely not” for certain, but only “maybe” for yes. A fruit machine. Three reels are the three hash functions and a row of sixteen flags is the bit array; it pays out only when all three of a word's flags are already up.
Two writers reading the same value at once will both act on a stale number. A water tank filled by two pumps. Without the interlock key both controllers read the same level and overfill it; with it, only the one holding the key may read and run.
Before any data moves, both ends have to agree that the other one is listening. A mail ropeway between two stations. Canisters flying SYN, SYN with an ACK pennant, then ACK ride the rope across, and each station's semaphore shows its connection state.
Temperature reshapes the odds of the next token without changing which tokens are possible. A roulette wheel whose sixteen pockets are the likeliest next tokens, each as wide as its probability. A thermostat knob sets the temperature and moves the frets.
Debounce waits for quiet; throttle lets one call through per window, however many arrive. A music-box recorder with a turning pin disc. Every click of the knob is a call; the throttle and debounce dashpots each set their own pins, so you can read when each function actually ran.
The pattern is the same every time. The object has to do the idea, not stand for it, so the thing you watch move is the thing being explained.
Small ideas make small machines. The big ones are where the skill earns its keep: hundreds of parts, real engineering numbers, and shaders that live inside the drawing rather than on top of it.
A Rolls-Royce Trent 970 , the engine of the A380, hung from a ground-run thrust frame with its near upper quarter cut away. Drag the lever and the fan, intermediate and high-pressure spools wind up at their own speeds. The fan swallows 1.2 tonnes of air a second, most of it leaves as a cold bypass jet while the rest burns in the combustor, and the load cell reads the thrust climbing from 14 to 334 kN . Press Space for a take-off run. 5 5 The audit measured 6,320 overlapping pairs across 541 solids and found every one drawn in the right order.
Raptor 2 , SpaceX’s full-flow staged-combustion engine, hanging in a test stand. It is one of the skill’s own examples and its reference for dense pipework. Hover it and it comes apart into eleven assemblies, from the gimbal to the nozzle; press I to fire it. Every drop of oxygen passes through an oxygen-rich preburner that spins the oxygen pump, every drop of methane through a fuel-rich one that spins the methane pump, and the two hot gases meet in a main chamber at 300 bar. 6 6 The plume, shock, splash and steam are WebGL, drawn between a back and a front SVG, so the stand passes in front of the flame.
The first working version took 38 minutes. Making it right took ten more hours, and an audit now proves that no pipe passes through anything, at rest, apart and in every frame between.
Most figures are pure SVG. When an idea needs fire,

[truncated]
