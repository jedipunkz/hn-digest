---
source: "https://yapern.com/blog/aivs-the-ai-video-studio/"
hn_url: "https://news.ycombinator.com/item?id=50241243"
title: "Aivs – The AI Video Studio"
article_title: "aivs - the ai video studio · Yapern"
image: "https://yapern.com/images/blog/aivs-the-ai-video-studio/directed-shot-frame.webp"
author: "andrewfromx"
captured_at: "2026-10-26T09:47:48Z"
capture_tool: "hn-digest"
hn_id: 50241243
score: 1
comments: 0
posted_at: "2026-10-26T09:32:12Z"
tags:
  - hacker-news
---

# Aivs – The AI Video Studio

- HN: [50241243](https://news.ycombinator.com/item?id=50241243)
- Source: [yapern.com](https://yapern.com/blog/aivs-the-ai-video-studio/)
- Score: 1
- Comments: 0
- Posted: 2026-10-26T09:32:12Z

## Translation

Title: Aivs – The AI Video Studio
Article title: aivs - the ai video studio · Yapern
Description: AIVS turns a cast screenplay into Codex-directed shots, renders them locally with LTX-2.5, and experiments with fixed voices and reusable character references.

Article text:
aivs - the ai video studio · Yapern
Skip to content
yapern .
Search Yapern
Create
Create message
Create sub-yap
This site
Create message
Create sub-yap
this site
Notifications
Login
yapern .
Create a message
B
I
S
x 2
H
“”
</>
•••
Replace image
Remove image
October 25, 2026 · Updated October 26, 2026 · Andrew Arrow
Sue and Joe are stranded beside a station wagon on a desert road. Joe has looked under the hood. “Good news. It’s definitely the car.” Sue has a postcard from her dad and a questionable promise that the ocean is somewhere nearby. It’s a tiny scene, five lines of dialogue, but it already creates a useful problem for an AI video tool: I want to choose who plays these people without rewriting who they are.
AIVS is the AI video studio I’m building around that workflow. It’s a native Rust app with a screenplay, a character gallery, and video generation on Apple Silicon. It now asks Codex to turn the cast screenplay into a shooting plan, then renders and assembles the short takes locally. The GUI still starts with the built-in Last Light scene; the command line can accept other scripts. Changing the camera shouldn’t quietly give me a different actor.
Sue is a role. Ava can play her.
The script originally called its two people Ava and Eli, the same names as their saved character profiles. The rename to Sue and Joe made the distinction explicit. Ava belongs in the library of performers; Sue belongs in the story. Recasting Sue shouldn’t require hunting through the script to replace every name.
Now you can drag Ava’s portrait or name onto any of Sue’s dialogue blocks. Every Sue cue gets the same assignment. Drop Eli onto Joe, and both parts are cast. Clicking a dialogue block also opens a picker, including an option to clear the assignment. Selecting a portrait in the gallery just browses it; casting requires choosing a role.
The implementation stores one assignment per role, rather than one per line. The screenplay preview and generation request come from the same scene data. Both roles must be cast before Generate video becomes available.
The old request embedded performer references in a long prompt. Rust now sends the screenplay text and structured cast bindings to Python. The cast portion looks like this:
"cast": [
{ "role": "JOE", "character": "eli" },
{ "role": "SUE", "character": "ava" }
]
Python loads those saved profiles explicitly. A role called Ava can’t accidentally cast itself just by sharing a library name. The scene’s dialogue keeps its role names, and the renderer adds appearance and voice descriptions under those names. Assignments still last only for the session, and the GUI screenplay remains a preview rather than a script editor.
The screenplay needs a shooting script
The first version estimated 14 seconds for the five spoken lines. Adding time for the action brought that estimate to 32 seconds, but a longer clip still left the model to figure out the entire scene. Somebody has to decide where the camera starts, who holds the postcard, and when there’s room for the final walk.
The shooting-script work puts a planning step between the screenplay and video generation. Python parses the script into ordered action and dialogue beats. Codex proposes takes lasting four to twelve seconds, with a prompt, opening state, ending state, and a decision about whether to continue the previous composition. One saved early plan for Last Light totals 27 seconds:
Codex chooses staging, but Python checks the proposed plan before filming it. Every source beat must appear exactly once, in order, and a take can’t cross scene boundaries. Dialogue enters the template through placeholders such as {{line_3}} ; Python fills in the original speaker, delivery cue, and words. It rejects duplicated placeholders, raw quoted dialogue, insufficient duration, and prompts over 1,024 Gemma tokens. A validation failure gets one corrective Codex attempt before the job fails.
That protects the source words and the plan’s structure. It doesn’t prove that an action described in the prompt happens on screen. The renderer films the takes in sequence, optionally uses the preceding take’s last frame as the next starting image, and joins them with FFmpeg. Assembly supplies the final fade, so the video model doesn’t have to interpret that instruction.
The next fixes came from a render report describing Joe emerging from the car without opening a door, plus gibberish at the end of the first shot. A valid shot list hadn’t made the performance physically sensible. Here’s a frame six seconds into the first take of that 27-second run. Even the requested red-rock desert has become a tree-lined road.
The director instructions now make Codex account for the opening frame and each movement that follows. Joe is already on the ground beside the raised hood, with the cabin doors shut. He steps around the fender into view. If a scene really requires an exit through a closed door, the plan has to include opening it. The postcard needs the same care: Sue gives it to Joe, but she pockets it later, so somebody has to give it back.
The soundtrack also needs an ending. Each dialogue take now asks for the speaker’s lips to settle after the last word, with the remaining action accompanied by environmental sound. Silent takes get explicit direction too. These instructions live in a director-instructions file , where they can improve staging across scripts without adding Python rules about particular cars or postcards.
One easy mistake is to put the important detail only in start_state or end_state . Those fields are continuity notes for review. The model receives the assembled prompt, so its opening positions, prop ownership, and quiet ending have to be written there. A carried last frame can help preserve a composition; it can’t supply a missing physical cause.
Planning is visible now, too. The GUI shows elapsed time and streams Codex activity while the shooting script is being prepared. The run keeps the requests, responses, and streamed diagnostics, including partial output when planning stops or times out. A long planning turn no longer has to look like a frozen Generate button.
Keeping the performer when the location changes
Better staging still leaves a casting problem. The normal screenplay path supplies appearance and voice descriptions, not each performer’s saved portrait and recorded voice. Continuation can carry a face into the next shot, but a new angle or location needs a way to establish it again. Text describing auburn hair and freckles doesn’t uniquely identify Ava.
The latest character-consistency work adds a separate command-line experiment: Lena plays Sue and Owen plays Joe, first on a desert road, then in a café. Each shot has one visible speaker. The prepared starting images use the original character portraits as references. These two images are inputs for Lena’s takes, with their generation prompts and source references saved alongside them:
These images condition complete starting frames. That’s why a solo gallery portrait is awkward for an ensemble scene: it also tells the model to fill the frame with one face. The experiment prepares a composition for each take and changes the camera separately from the performer’s original identity reference.
It also makes the dialogue before the video. A CPU speech worker uses Lena’s saved Kokoro voice ID, af_sarah , and Owen’s am_michael . Each source line is synthesized once, measured, and reused in both locations. The shot soundtrack adds a short opening pause and silence after the line. This gives the video a fixed vocal performance to follow instead of asking it to invent a new voice for every take.
The audio-driven backend accepts that soundtrack and the starting image together. Final encoding selects the supplied soundtrack rather than the model’s generated audio. The preparation records image and audio hashes, voice and model versions, measured durations, and speaker intervals. Saved-plan freshness checks now include the character reference-asset hashes as well, so replacing an asset can invalidate a cached plan even if its description hasn’t changed.
This is still an opt-in experiment, separate from GUI screenplay generation. The four-shot test uses 768 × 448 at 24 fps. The inspected run has one completed four-second take, marked rendered-needs-review ; the comparison across both locations remains incomplete. Faces, voice identity, spoken words, and lip synchronization still need review. A supplied voice track doesn’t by itself tell a model which of two visible faces should be speaking.
A playable file still needs a director
The normal GUI path uses LTX-2.5 to generate picture and speech together at 1280 × 704 and 24 fps. Video inference runs locally from cached models; shooting-script preparation uses the signed-in Codex client. Initial model setup still downloads about 43 GB and requires 50 GiB free. Rust owns process lifetime, including Stop, and completed films play inside the native window through macOS AVKit.
Runs now save the human screenplay, structured request, readable and JSON shooting plans, cast snapshots, and individual shot reports. A later failure leaves earlier takes available to inspect. Before publishing video.mp4 , media validation checks decoded frames, dimensions, frame rate, and audio/video timing. Those checks can catch a broken export while missing an impossible entrance or a swapped speaker.
I can now fix how Joe enters the shot without changing his dialogue, and test a camera change while keeping the prepared voice track. Character consistency is still unfinished. The next test is whether Lena and Owen survive all four takes with their faces, voices, and lines intact. Casting only earns its keep if I can carry those performers into another scene.

## Original Extract

AIVS turns a cast screenplay into Codex-directed shots, renders them locally with LTX-2.5, and experiments with fixed voices and reusable character references.

aivs - the ai video studio · Yapern
Skip to content
yapern .
Search Yapern
Create
Create message
Create sub-yap
This site
Create message
Create sub-yap
this site
Notifications
Login
yapern .
Create a message
B
I
S
x 2
H
“”
</>
•••
Replace image
Remove image
October 25, 2026 · Updated October 26, 2026 · Andrew Arrow
Sue and Joe are stranded beside a station wagon on a desert road. Joe has looked under the hood. “Good news. It’s definitely the car.” Sue has a postcard from her dad and a questionable promise that the ocean is somewhere nearby. It’s a tiny scene, five lines of dialogue, but it already creates a useful problem for an AI video tool: I want to choose who plays these people without rewriting who they are.
AIVS is the AI video studio I’m building around that workflow. It’s a native Rust app with a screenplay, a character gallery, and video generation on Apple Silicon. It now asks Codex to turn the cast screenplay into a shooting plan, then renders and assembles the short takes locally. The GUI still starts with the built-in Last Light scene; the command line can accept other scripts. Changing the camera shouldn’t quietly give me a different actor.
Sue is a role. Ava can play her.
The script originally called its two people Ava and Eli, the same names as their saved character profiles. The rename to Sue and Joe made the distinction explicit. Ava belongs in the library of performers; Sue belongs in the story. Recasting Sue shouldn’t require hunting through the script to replace every name.
Now you can drag Ava’s portrait or name onto any of Sue’s dialogue blocks. Every Sue cue gets the same assignment. Drop Eli onto Joe, and both parts are cast. Clicking a dialogue block also opens a picker, including an option to clear the assignment. Selecting a portrait in the gallery just browses it; casting requires choosing a role.
The implementation stores one assignment per role, rather than one per line. The screenplay preview and generation request come from the same scene data. Both roles must be cast before Generate video becomes available.
The old request embedded performer references in a long prompt. Rust now sends the screenplay text and structured cast bindings to Python. The cast portion looks like this:
"cast": [
{ "role": "JOE", "character": "eli" },
{ "role": "SUE", "character": "ava" }
]
Python loads those saved profiles explicitly. A role called Ava can’t accidentally cast itself just by sharing a library name. The scene’s dialogue keeps its role names, and the renderer adds appearance and voice descriptions under those names. Assignments still last only for the session, and the GUI screenplay remains a preview rather than a script editor.
The screenplay needs a shooting script
The first version estimated 14 seconds for the five spoken lines. Adding time for the action brought that estimate to 32 seconds, but a longer clip still left the model to figure out the entire scene. Somebody has to decide where the camera starts, who holds the postcard, and when there’s room for the final walk.
The shooting-script work puts a planning step between the screenplay and video generation. Python parses the script into ordered action and dialogue beats. Codex proposes takes lasting four to twelve seconds, with a prompt, opening state, ending state, and a decision about whether to continue the previous composition. One saved early plan for Last Light totals 27 seconds:
Codex chooses staging, but Python checks the proposed plan before filming it. Every source beat must appear exactly once, in order, and a take can’t cross scene boundaries. Dialogue enters the template through placeholders such as {{line_3}} ; Python fills in the original speaker, delivery cue, and words. It rejects duplicated placeholders, raw quoted dialogue, insufficient duration, and prompts over 1,024 Gemma tokens. A validation failure gets one corrective Codex attempt before the job fails.
That protects the source words and the plan’s structure. It doesn’t prove that an action described in the prompt happens on screen. The renderer films the takes in sequence, optionally uses the preceding take’s last frame as the next starting image, and joins them with FFmpeg. Assembly supplies the final fade, so the video model doesn’t have to interpret that instruction.
The next fixes came from a render report describing Joe emerging from the car without opening a door, plus gibberish at the end of the first shot. A valid shot list hadn’t made the performance physically sensible. Here’s a frame six seconds into the first take of that 27-second run. Even the requested red-rock desert has become a tree-lined road.
The director instructions now make Codex account for the opening frame and each movement that follows. Joe is already on the ground beside the raised hood, with the cabin doors shut. He steps around the fender into view. If a scene really requires an exit through a closed door, the plan has to include opening it. The postcard needs the same care: Sue gives it to Joe, but she pockets it later, so somebody has to give it back.
The soundtrack also needs an ending. Each dialogue take now asks for the speaker’s lips to settle after the last word, with the remaining action accompanied by environmental sound. Silent takes get explicit direction too. These instructions live in a director-instructions file , where they can improve staging across scripts without adding Python rules about particular cars or postcards.
One easy mistake is to put the important detail only in start_state or end_state . Those fields are continuity notes for review. The model receives the assembled prompt, so its opening positions, prop ownership, and quiet ending have to be written there. A carried last frame can help preserve a composition; it can’t supply a missing physical cause.
Planning is visible now, too. The GUI shows elapsed time and streams Codex activity while the shooting script is being prepared. The run keeps the requests, responses, and streamed diagnostics, including partial output when planning stops or times out. A long planning turn no longer has to look like a frozen Generate button.
Keeping the performer when the location changes
Better staging still leaves a casting problem. The normal screenplay path supplies appearance and voice descriptions, not each performer’s saved portrait and recorded voice. Continuation can carry a face into the next shot, but a new angle or location needs a way to establish it again. Text describing auburn hair and freckles doesn’t uniquely identify Ava.
The latest character-consistency work adds a separate command-line experiment: Lena plays Sue and Owen plays Joe, first on a desert road, then in a café. Each shot has one visible speaker. The prepared starting images use the original character portraits as references. These two images are inputs for Lena’s takes, with their generation prompts and source references saved alongside them:
These images condition complete starting frames. That’s why a solo gallery portrait is awkward for an ensemble scene: it also tells the model to fill the frame with one face. The experiment prepares a composition for each take and changes the camera separately from the performer’s original identity reference.
It also makes the dialogue before the video. A CPU speech worker uses Lena’s saved Kokoro voice ID, af_sarah , and Owen’s am_michael . Each source line is synthesized once, measured, and reused in both locations. The shot soundtrack adds a short opening pause and silence after the line. This gives the video a fixed vocal performance to follow instead of asking it to invent a new voice for every take.
The audio-driven backend accepts that soundtrack and the starting image together. Final encoding selects the supplied soundtrack rather than the model’s generated audio. The preparation records image and audio hashes, voice and model versions, measured durations, and speaker intervals. Saved-plan freshness checks now include the character reference-asset hashes as well, so replacing an asset can invalidate a cached plan even if its description hasn’t changed.
This is still an opt-in experiment, separate from GUI screenplay generation. The four-shot test uses 768 × 448 at 24 fps. The inspected run has one completed four-second take, marked rendered-needs-review ; the comparison across both locations remains incomplete. Faces, voice identity, spoken words, and lip synchronization still need review. A supplied voice track doesn’t by itself tell a model which of two visible faces should be speaking.
A playable file still needs a director
The normal GUI path uses LTX-2.5 to generate picture and speech together at 1280 × 704 and 24 fps. Video inference runs locally from cached models; shooting-script preparation uses the signed-in Codex client. Initial model setup still downloads about 43 GB and requires 50 GiB free. Rust owns process lifetime, including Stop, and completed films play inside the native window through macOS AVKit.
Runs now save the human screenplay, structured request, readable and JSON shooting plans, cast snapshots, and individual shot reports. A later failure leaves earlier takes available to inspect. Before publishing video.mp4 , media validation checks decoded frames, dimensions, frame rate, and audio/video timing. Those checks can catch a broken export while missing an impossible entrance or a swapped speaker.
I can now fix how Joe enters the shot without changing his dialogue, and test a camera change while keeping the prepared voice track. Character consistency is still unfinished. The next test is whether Lena and Owen survive all four takes with their faces, voices, and lines intact. Casting only earns its keep if I can carry those performers into another scene.
