---
source: "https://favtutor.com/claude-opus-5-5-real-examples/"
hn_url: "https://news.ycombinator.com/item?id=49823070"
title: "Claude Opus 5.5: Things People Created"
article_title: "Claude Opus 5.5: 10 Amazing Things People Created"
image: "https://favtutor.com/images/og-claude-opus-5-5-real-examples-v1.jpg"
author: "wslh"
captured_at: "2026-09-23T22:44:11Z"
capture_tool: "hn-digest"
hn_id: 49823070
score: 1
comments: 0
posted_at: "2026-09-23T21:50:24Z"
tags:
  - hacker-news
---

# Claude Opus 5.5: Things People Created

- HN: [49823070](https://news.ycombinator.com/item?id=49823070)
- Source: [favtutor.com](https://favtutor.com/claude-opus-5-5-real-examples/)
- Score: 1
- Comments: 0
- Posted: 2026-09-23T21:50:24Z

## Translation

Title: Claude Opus 5.5: Things People Created
Article title: Claude Opus 5.5: 10 Amazing Things People Created
Description: See 10 real projects built with Claude Opus 5.5 by real users to test it, from playable games to 3D animations and interactive tools.

Article text:
AI News Tutorials Tools Claude Watermark Remover AI Code Generator AI Code Debugger AI Data Analysis AI Code Converter
Sign in
Home › Opus 5.5 Examples
AI News
Claude Opus 5.5: 10 Amazing Things People Created
Published September 23, 2026
Updated 5:44 PM IST · 8 min read
The gist
Anthropic released Claude Opus 5.5 on 22 September. It matches Claude Fable 5.1 on most work and costs 40% less to run than Opus 5.
It generates output more than 30% faster , and a Fast Mode in Claude Code goes up to 2.5x that speed.
People are already using it to build playable games, 3D animations, browser films and travel planners, all from a single prompt.
Anthropic released Claude Opus 5.5 on 22 September, and people started building with it the same day. The results are the point of this piece, so most of it is examples.
The short version of the model itself is that it does more for less money. It performs at the level of Claude Fable 5.1 on most work while costing 40% less to run than Opus 5.
What Makes Claude Opus 5.5 Special?
It is the first model in the Claude 5.5 family, and the cheapest way Anthropic has offered to get this level of work done.
The pricing is where the difference shows up fastest. Input and output tokens are $4 and $20 per million, 20% below Opus 5. Cache reads, which Anthropic says make up most of the cost of agentic and coding work, fell from $0.50 per million to $0.20 . It also uses fewer tokens per task, and those two things together are where the 40% comes from.
Speed moved as well. Opus 5.5 generates output more than 30% faster than Opus 5. A Fast Mode in Claude Code and the Claude Platform pushes that to 2.5x , at $8 and $40 per million tokens.
The coding numbers Anthropic published are the ones worth holding onto. One early tester finished a 680,000-line code migration in under a day. Another audited and fixed a 200,000-line codebase in under three hours, where Opus 5 had taken over 20 hours and burned 2.5 times as many tokens. On Terminal-Bench 4.0 it scores 66.4% , against 55.8% for Fable 5.1 and 57.9% for GPT-6 Astra.
Anthropic also says the model writes more clearly. It puts the important information first and uses less jargon, which was a common complaint about Opus 5. One tester told them "it writes the way I do".
There is a bigger piece of context here. This is Anthropic's first model release since Dario Amodei publicly called for pacing the frontier , the same argument Nvidia's Jensen Huang rejected days later . A faster and cheaper flagship arriving weeks later is a fair thing to raise an eyebrow at. Anthropic says Opus 5.5 scored better than any model it has tested on its automated behavioural audit, and that external evaluators including METR and Frontier Design tested it before release. Because its biology and cybersecurity abilities are close to Claude Mythos 5.1, it ships with the same safeguards as Fable 5.1.
It is available on Claude Pro, Max, Team and Enterprise plans and through the API. Sonnet 5.5 and Haiku 5.5 are due in the coming weeks.
10 Beautiful Things People Built With Opus 5.5
Everything below was made in the first two days, mostly from a single prompt.
1. A pencil sketch turned into a 3D physics simulation. Ben Poole, formerly a researcher at Google Brain and DeepMind, drew a trebuchet on paper and fed the photo to Opus 5.5.
What came back was a working launcher simulator with real ballistics and interactive controls. You can change the projectile weight and the angle, fire at a stack of virtual blocks, and replay the physics. That used to need 3D modelling software and a physics engine.
2. A playable Minecraft clone in the browser. AI creator Noah Wachnik asked for a playable version of Minecraft with advanced shaders, as real and beautiful as possible.
He got a game called Lumen Vale in 1 hour and 37 minutes at max effort. You walk around in first person, breaking blocks and placing new ones. The water ripples and catches the light, the lighting shifts with the time of day, and the generated terrain runs from beaches up into mountains. All of it plays in a browser tab with nothing installed.
3. A travel planner with a 3D globe. Nate Herk ran Opus 5.5 against GPT-6 Sol across 10 use cases on YouTube. One was a 30-day trip from Chicago mixing tech events, nature and international travel, with booking links and an interactive globe.
Opus 5.5 built the 3D interface, so you can see every destination at once and click through to book. Its research also surfaced real events such as San Francisco Tech Week that Sol missed. Sol produced a working itinerary with links, but it never delivered the 3D part that was asked for.
4. A 4D view of an animal moving through time. Bilawal Sidhu, a former Google product manager, built a visualisation that shows a creature's whole path through a forest as one glowing structure.
The idea comes from Kurt Vonnegut's Slaughterhouse-Five and its image of humans as great millipedes seen across time. Technically it takes 2D video, turns each frame into 3D Gaussians, and pins them in space-time to produce the 4D render.
5. A holiday sketchbook in acrylic marker style. AI educator Ann Nong handed over her New Zealand trip photos and asked for them as drawings.
The result is an interactive digital sketchbook she can flip through. Every page is captioned and carries the original photo as a thumbnail, so you can see the sheep paddock or the harbour it was drawn from.
6. Two minutes of US history in sand animation. Michael Guo asked for a lively, tasteful sand animation covering 250 years, with background music and sound design.
It produced a full 2D animated film with music and sound effects to match. It opens on the 1776 Declaration with the Liberty Bell and closes on fireworks over Washington in 2026, hitting the moon landing and the first flight along the way. None of it involved Blender or Three.js.
7. A watch brand site with 3D models. AI consultant Bijan Bowen asked for a cinematic website with rendered watches and an exploded view of the movement.
The first attempt missed the special edition watch and picked a font that was hard to read. Bijan gave follow-up instructions and Opus 5.5 fixed both on the second pass. The finished site has realistic reflections, stitching detail on the leather straps, and an exploded view you can rotate to inspect parts like the balance wheel.
8. An animated pixel art wizard. Entrepreneur Majid wrote an unusually precise prompt. He asked for a self-contained HTML file at 128x96 resolution with a fixed 24-colour palette. It needed a particle system for the spell effects, plus a state machine running the wizard from idle through charge and cast to recovery.
He got exactly that. The wizard bobs when idle, raises its staff and throws sparks while charging, and fires a burst with a pixel-perfect screen shake when casting. Particles snap to the pixel grid and shift through the palette as they die. It runs at 60 frames per second with no object allocations in the animation loop.
9. An 80-second film made of glass tiles. Chris Riley asked for a square animated film as a single HTML file using WebGL2 and plain JavaScript, with no libraries and no image, font or audio files.
The brief described a glass and gold leaf mosaic whose tiles were never glued down. They lift off the wall and click back into place, and every figure is a flock of tiles. A fish swims because its tiles swim. The finished piece has the fish, cranes and doves formed by flocking tiles, a day to night transition with stars becoming a constellation, and a final spiral where everything disperses.
10. A fantasy castle rendered in Blender. 3D AI researcher Stefan Vaskevich asked Opus 5.5 and GPT-6 Astra to procedurally build and render a 10-second Blender animation from one prompt, with no manual editing. The scene called for a detailed castle reflected in a lake, with fireworks going off above the mountains behind it.
Opus finished in 35 minutes on 199,600 output tokens for roughly $13.30 in API costs. Astra finished quicker at 28 minutes and used fewer tokens, but still cost $14.50 . Stefan's read was that Opus 5.5 held more complexity at once and did more of the creative work itself.
The thread running through all 10 is that the specialist software is missing. Nobody opened Blender or a game engine to make any of this. Someone who cannot open Blender got a rendered castle out of it, and someone who cannot write WebGL got an 80-second film.
The cost is the part I would watch. Animation, game development and 3D work are expensive mostly because the tools are specialised and the people who can drive them are scarce. Stefan's castle cost $13.30. That does not put studios out of business, but it does change what one person with an idea and an afternoon can reasonably attempt.
Worth keeping the frame honest though. These are demos, chosen and posted by people who were impressed enough to post them, and nobody shares the four attempts that came out broken. Bijan's watch site needed a second pass before it was right, and that is the more realistic picture of what building with it actually feels like.
I’m Kaustubh Saini, founder of FavTutor. I love breaking down complex AI concepts, trends, and news, writing about them until an AGI agent takes over my job. When I’m not writing, I’m building AI-powered tools to make learning more accessible and engaging at FavTutor.
What Makes Claude Opus 5.5 Special?
10 Beautiful Things People Built With Opus 5.5
AI news, practical guides, tools and coding resources for developers, students and builders.

## Original Extract

See 10 real projects built with Claude Opus 5.5 by real users to test it, from playable games to 3D animations and interactive tools.

AI News Tutorials Tools Claude Watermark Remover AI Code Generator AI Code Debugger AI Data Analysis AI Code Converter
Sign in
Home › Opus 5.5 Examples
AI News
Claude Opus 5.5: 10 Amazing Things People Created
Published September 23, 2026
Updated 5:44 PM IST · 8 min read
The gist
Anthropic released Claude Opus 5.5 on 22 September. It matches Claude Fable 5.1 on most work and costs 40% less to run than Opus 5.
It generates output more than 30% faster , and a Fast Mode in Claude Code goes up to 2.5x that speed.
People are already using it to build playable games, 3D animations, browser films and travel planners, all from a single prompt.
Anthropic released Claude Opus 5.5 on 22 September, and people started building with it the same day. The results are the point of this piece, so most of it is examples.
The short version of the model itself is that it does more for less money. It performs at the level of Claude Fable 5.1 on most work while costing 40% less to run than Opus 5.
What Makes Claude Opus 5.5 Special?
It is the first model in the Claude 5.5 family, and the cheapest way Anthropic has offered to get this level of work done.
The pricing is where the difference shows up fastest. Input and output tokens are $4 and $20 per million, 20% below Opus 5. Cache reads, which Anthropic says make up most of the cost of agentic and coding work, fell from $0.50 per million to $0.20 . It also uses fewer tokens per task, and those two things together are where the 40% comes from.
Speed moved as well. Opus 5.5 generates output more than 30% faster than Opus 5. A Fast Mode in Claude Code and the Claude Platform pushes that to 2.5x , at $8 and $40 per million tokens.
The coding numbers Anthropic published are the ones worth holding onto. One early tester finished a 680,000-line code migration in under a day. Another audited and fixed a 200,000-line codebase in under three hours, where Opus 5 had taken over 20 hours and burned 2.5 times as many tokens. On Terminal-Bench 4.0 it scores 66.4% , against 55.8% for Fable 5.1 and 57.9% for GPT-6 Astra.
Anthropic also says the model writes more clearly. It puts the important information first and uses less jargon, which was a common complaint about Opus 5. One tester told them "it writes the way I do".
There is a bigger piece of context here. This is Anthropic's first model release since Dario Amodei publicly called for pacing the frontier , the same argument Nvidia's Jensen Huang rejected days later . A faster and cheaper flagship arriving weeks later is a fair thing to raise an eyebrow at. Anthropic says Opus 5.5 scored better than any model it has tested on its automated behavioural audit, and that external evaluators including METR and Frontier Design tested it before release. Because its biology and cybersecurity abilities are close to Claude Mythos 5.1, it ships with the same safeguards as Fable 5.1.
It is available on Claude Pro, Max, Team and Enterprise plans and through the API. Sonnet 5.5 and Haiku 5.5 are due in the coming weeks.
10 Beautiful Things People Built With Opus 5.5
Everything below was made in the first two days, mostly from a single prompt.
1. A pencil sketch turned into a 3D physics simulation. Ben Poole, formerly a researcher at Google Brain and DeepMind, drew a trebuchet on paper and fed the photo to Opus 5.5.
What came back was a working launcher simulator with real ballistics and interactive controls. You can change the projectile weight and the angle, fire at a stack of virtual blocks, and replay the physics. That used to need 3D modelling software and a physics engine.
2. A playable Minecraft clone in the browser. AI creator Noah Wachnik asked for a playable version of Minecraft with advanced shaders, as real and beautiful as possible.
He got a game called Lumen Vale in 1 hour and 37 minutes at max effort. You walk around in first person, breaking blocks and placing new ones. The water ripples and catches the light, the lighting shifts with the time of day, and the generated terrain runs from beaches up into mountains. All of it plays in a browser tab with nothing installed.
3. A travel planner with a 3D globe. Nate Herk ran Opus 5.5 against GPT-6 Sol across 10 use cases on YouTube. One was a 30-day trip from Chicago mixing tech events, nature and international travel, with booking links and an interactive globe.
Opus 5.5 built the 3D interface, so you can see every destination at once and click through to book. Its research also surfaced real events such as San Francisco Tech Week that Sol missed. Sol produced a working itinerary with links, but it never delivered the 3D part that was asked for.
4. A 4D view of an animal moving through time. Bilawal Sidhu, a former Google product manager, built a visualisation that shows a creature's whole path through a forest as one glowing structure.
The idea comes from Kurt Vonnegut's Slaughterhouse-Five and its image of humans as great millipedes seen across time. Technically it takes 2D video, turns each frame into 3D Gaussians, and pins them in space-time to produce the 4D render.
5. A holiday sketchbook in acrylic marker style. AI educator Ann Nong handed over her New Zealand trip photos and asked for them as drawings.
The result is an interactive digital sketchbook she can flip through. Every page is captioned and carries the original photo as a thumbnail, so you can see the sheep paddock or the harbour it was drawn from.
6. Two minutes of US history in sand animation. Michael Guo asked for a lively, tasteful sand animation covering 250 years, with background music and sound design.
It produced a full 2D animated film with music and sound effects to match. It opens on the 1776 Declaration with the Liberty Bell and closes on fireworks over Washington in 2026, hitting the moon landing and the first flight along the way. None of it involved Blender or Three.js.
7. A watch brand site with 3D models. AI consultant Bijan Bowen asked for a cinematic website with rendered watches and an exploded view of the movement.
The first attempt missed the special edition watch and picked a font that was hard to read. Bijan gave follow-up instructions and Opus 5.5 fixed both on the second pass. The finished site has realistic reflections, stitching detail on the leather straps, and an exploded view you can rotate to inspect parts like the balance wheel.
8. An animated pixel art wizard. Entrepreneur Majid wrote an unusually precise prompt. He asked for a self-contained HTML file at 128x96 resolution with a fixed 24-colour palette. It needed a particle system for the spell effects, plus a state machine running the wizard from idle through charge and cast to recovery.
He got exactly that. The wizard bobs when idle, raises its staff and throws sparks while charging, and fires a burst with a pixel-perfect screen shake when casting. Particles snap to the pixel grid and shift through the palette as they die. It runs at 60 frames per second with no object allocations in the animation loop.
9. An 80-second film made of glass tiles. Chris Riley asked for a square animated film as a single HTML file using WebGL2 and plain JavaScript, with no libraries and no image, font or audio files.
The brief described a glass and gold leaf mosaic whose tiles were never glued down. They lift off the wall and click back into place, and every figure is a flock of tiles. A fish swims because its tiles swim. The finished piece has the fish, cranes and doves formed by flocking tiles, a day to night transition with stars becoming a constellation, and a final spiral where everything disperses.
10. A fantasy castle rendered in Blender. 3D AI researcher Stefan Vaskevich asked Opus 5.5 and GPT-6 Astra to procedurally build and render a 10-second Blender animation from one prompt, with no manual editing. The scene called for a detailed castle reflected in a lake, with fireworks going off above the mountains behind it.
Opus finished in 35 minutes on 199,600 output tokens for roughly $13.30 in API costs. Astra finished quicker at 28 minutes and used fewer tokens, but still cost $14.50 . Stefan's read was that Opus 5.5 held more complexity at once and did more of the creative work itself.
The thread running through all 10 is that the specialist software is missing. Nobody opened Blender or a game engine to make any of this. Someone who cannot open Blender got a rendered castle out of it, and someone who cannot write WebGL got an 80-second film.
The cost is the part I would watch. Animation, game development and 3D work are expensive mostly because the tools are specialised and the people who can drive them are scarce. Stefan's castle cost $13.30. That does not put studios out of business, but it does change what one person with an idea and an afternoon can reasonably attempt.
Worth keeping the frame honest though. These are demos, chosen and posted by people who were impressed enough to post them, and nobody shares the four attempts that came out broken. Bijan's watch site needed a second pass before it was right, and that is the more realistic picture of what building with it actually feels like.
I’m Kaustubh Saini, founder of FavTutor. I love breaking down complex AI concepts, trends, and news, writing about them until an AGI agent takes over my job. When I’m not writing, I’m building AI-powered tools to make learning more accessible and engaging at FavTutor.
What Makes Claude Opus 5.5 Special?
10 Beautiful Things People Built With Opus 5.5
AI news, practical guides, tools and coding resources for developers, students and builders.
