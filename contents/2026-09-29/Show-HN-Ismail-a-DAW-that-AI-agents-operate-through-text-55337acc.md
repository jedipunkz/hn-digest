---
source: "https://github.com/newsbubbles/ismail"
hn_url: "https://news.ycombinator.com/item?id=49891000"
title: "Show HN: Ismail, a DAW that AI agents operate through text"
article_title: "GitHub - newsbubbles/ismail: A DAW for AI agents: write music as text, read the audio back as text. MCP server, CLI and an agent skill. · GitHub"
image: "https://opengraph.githubassets.com/ac0ffb517ee9a37639a1261ada1747df3f0b75b880449e382ebacc8f47085521/newsbubbles/ismail"
author: "natecodes"
captured_at: "2026-09-29T11:10:15Z"
capture_tool: "hn-digest"
hn_id: 49891000
score: 1
comments: 1
posted_at: "2026-09-29T10:44:42Z"
tags:
  - hacker-news
---

# Show HN: Ismail, a DAW that AI agents operate through text

- HN: [49891000](https://news.ycombinator.com/item?id=49891000)
- Source: [github.com](https://github.com/newsbubbles/ismail)
- Score: 1
- Comments: 1
- Posted: 2026-09-29T10:44:42Z

## Translation

Title: Show HN: Ismail, a DAW that AI agents operate through text
Article title: GitHub - newsbubbles/ismail: A DAW for AI agents: write music as text, read the audio back as text. MCP server, CLI and an agent skill. · GitHub
Description: A DAW for AI agents: write music as text, read the audio back as text. MCP server, CLI and an agent skill. - newsbubbles/ismail

Article text:
GitHub - newsbubbles/ismail: A DAW for AI agents: write music as text, read the audio back as text. MCP server, CLI and an agent skill. · GitHub
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
newsbubbles
/
ismail
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.cursor .cursor .github/ workflows .github/ workflows ismail ismail skills/ ismail skills/ ismail tests tests .gitattributes .gitattributes .gitignore .gitignore .mcp.json .mcp.json LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
A DAW for AI agents. It can't hear, so it reads.
Listen to songs an agent made with it , each shown with the text the agent read while making it. The playhead runs across that text as the song plays.
Luigi Manson, made with ismail (fan remix of the Luigi's Mansion theme).
ismail is not a model that turns a prompt into audio, like Suno. It is a set of tools your own agent uses to write the song as notes, sounds and code, render it, read back what came out, and edit it. That changes what you get:
Iterative, precise edits. Change one note, one patch, one bar or one fader and re-render; nothing else moves.
Songs are code. A project file and a build script: git history, diffs, branches and code review work on a track.
Any sound. Synths, drum synths, samplers, voices written in Python and speech; any sound can become an instrument.
Local and open. MIT licensed, runs on your machine, no content filter and no music subscription (you bring the agent).
It improves with your model. The music is the agent's own work, so a stronger model with the same prompt should write a better song.
Suno is still better at realistic sung vocals, a polished song from one sentence in under a minute, and genre sound learned from recorded music. And why not Ableton or FL Studio? They were built for a person with ears and a mouse; an agent can press their buttons through bridges but still can't hear what it did. ismail puts everything an agent needs to write and to perceive into compact text, and if something is missing, your agent can add it. More on the showcase page .
A DAW built to be operated by an AI agent. Everything goes in as text (notes, instrument patches, effect chains, automation) and everything comes back as text: levels, spectra, drum patterns, piano rolls, chords, vowels, song structure, and structured comparisons against a reference track. The agent never needs ears or images to work (a spectrogram PNG is there if you want one).
One set of operations, three ways in:
MCP server for Claude Code, Cursor or any MCP client: python -m ismail.mcp_server (stdio, about 70 tools)
CLI : python -m ismail -p <project> <op> [args]
Python : from ismail import api
git clone https://github.com/newsbubbles/ismail
cd ismail
pip install -e . # engine, analysis, CLI, MCP server
pip install -e " .[perceptual] " # optional: CLAP perceptual metric (torch + transformers, model about 600 MB)
pip install -e " .[separate] " # optional: demucs stem separation for reference tracks
If demucs fights your torch install, use pip install --no-deps demucs and then pip install dora-search einops julius lameenc openunmix .
MP3 previews need ffmpeg on your PATH (or set ISMAIL_FFMPEG to the binary).
Runs on Windows, macOS and Linux; CI tests all three on every push. The one OS-specific op is sound_speak (text to speech for vocal samples), which uses the engine the OS already has:
Tools. Open Claude Code in this folder and the bundled .mcp.json registers the server; the tools show up as mcp__ismail__* . To use ismail from any folder instead:
claude mcp add -s user ismail -- python -m ismail.mcp_server
Skill (recommended). skills/ismail teaches the agent how to compose with ismail: plan a Session Sheet before writing notes, write a Listening Report after every render, and judge reference matches with the comparison tools instead of by feel. Link it into your skills folder:
# macOS / Linux
ln -s " $( pwd ) /skills/ismail " ~ /.claude/skills/ismail
# Windows
New-Item - ItemType Junction - Path " $ env: USERPROFILE \.claude\skills\ismail " - Target " $PWD \skills\ismail "
Ask for music. For example: "make a 16 bar deep house loop in F minor in songs/demo and render an mp3". The agent calls guide once for the conventions (it is a tool and a CLI op), then works through the tools.
Tools. Opening this folder in Cursor picks up .cursor/mcp.json . To use ismail in other projects, add the same entry to ~/.cursor/mcp.json :
{ "mcpServers" : { "ismail" : { "command" : " python " , "args" : [ " -m " , " ismail.mcp_server " ]}}}
Skill. .cursor/rules/ismail.mdc is an agent-requested rule that points Cursor's agent at skills/ismail/SKILL.md . Copy that rule (and the skills/ismail folder) into another project to use it there.
Any other MCP client works the same way: run python -m ismail.mcp_server over stdio.
Every tool is also a CLI op. Arguments are key=value pairs (values parsed as JSON when they can be) or one JSON object.
python -m ismail guide # read first: workflow and conventions
python -m ismail ops # list operations
python -m ismail help notes_write # one op's arguments and docs
python -m ismail -p songs/demo project_new bpm=124 length_bars=8
python -m ismail -p songs/demo track_add name=bass instrument= ' "preset:acid_bass" '
python -m ismail -p songs/demo notes_write ' {"track": "bass", "bar": 1, "notes": "0 E2 0.5 110; 0.5 E3 0.25", "repeat": 8} '
python -m ismail -p songs/demo render stems=true out=v1 mp3=also
python -m ismail -p songs/demo analyze_melody source=track:bass bars=[1,2]
render writes renders/latest.wav (every analysis tool reads it), plus renders/<out>.wav when you name the render. mp3='also' adds renders/<out>.mp3 for listening; mp3='only' writes the named render as mp3 only.
Keep your projects under songs/ (git-ignored) or anywhere else; a project is just a folder.
Project : a folder with project.json (tempo, grid offset, tracks, buses, master, sound bank, reference) plus sounds/ , renders/ , cache/ , history/ (undo snapshots) and comparisons/ .
Time : bars are 1-indexed; note times are beats relative to the bar you write at. offset_sec is the time of bar 1, so a project can sit exactly on a reference recording's grid.
Notes : '<beat> <pitch> <dur> [vel]' , one per line or ; -separated. Drum and step patterns: pattern_write with strings like X...x...X...x... (X 127, x 100, o 70, - 45, _ ties).
Instruments : synth (saw, square, pulse, triangle, sine, additive, wavetable and noise oscillators, unison, FM, drive, SVF and ladder filters, envelopes, LFOs, mono glide), sampler , drum synths ( kick , snare , hat , clap , tom , noise_hit ), kit (pitch to instrument map) and code (a Python voice function for anything else). presets_list has starting points.
Effects : eq, filter, distortion, bitcrush, compressor (with sidechain), duck, gate, delay, reverb, chorus, flanger, phaser, tremolo/autopan, width, limiter, vocoder, formant. Tracks, buses and the master fader can be automated.
Voices : engineered instruments kept as Python modules, so a project stores a name instead of code (see below).
Sound bank : sounds made from any instrument and effect chain ( sound_make ), speech ( sound_speak ), imported files, and averaged events cut from a recording ( sound_extract ). Bank sounds work as sampler sources, wavetables, vocoder modulators and audio clips.
Undo and batch : every edit snapshots the project ( undo ); batch applies a list of ops atomically.
Some instruments are easier to write than to patch: a measured grand piano, a dubstep bass whose note velocity picks the articulation, a set of sound effects. These live as voice modules, Python files that define voice(freq, t, vel, gate, sr) and return a mono (n,) or stereo (2, n) array. freq is in Hz, t is an array of seconds from the note start that covers the held time plus the instrument's tail , vel is 0 to 1, gate is how long the note is held in seconds, and sr is the sample rate.
Use one with instrument={"type": "code", "voice": "grand_piano", "tail": 4.0} or "preset:grand_piano" . voices_list shows what is available and voice_help(name) explains a voice's velocity mapping, functions and parameters.
Voices are looked up in this order:
<project>/voices/<name>.py : the song's own. Same name as a built-in overrides it; a song voice can also extend one ( from ismail.voices.growl import * , then add words or articulations).
Each folder in $ISMAIL_VOICES (a path list): your personal library, outside any repo.
ismail/voices/ : the built-ins.
A voice function may take extra keyword arguments: bpm is passed automatically, and the track's "params" dict is passed as keywords ( {"type": "code", "voice": "mine", "params": {"brightness": 0.3}} ). A module-level INFO dict documents it for voice_help ; every key is optional: summary (one line for voices_list ), range , velocity (what velocity does), functions (name to description), params (name to description) and tail (recommended tail). Data files sit next to the module ( grand_piano.json ) and are found through __file__ . Editing a voice file invalidates the render cache for the tracks that use it.
To add a voice to the library, move it from a song's voices/ folder into ismail/voices/ , give it an INFO dict, and add a line to the test that renders every built-in.
Question
Tool
Tempo and where bar 1 is
analyze_grid , align
Song form, what plays where
analyze_structure (arrangement map, sections, loop length, root per bar)
Levels, bands and chords per bar
analyze_bars , analyze_chords , analyze_key
Drum pattern
analyze_drums (step strings you can paste into pattern_write )
Notes
analyze_pitches (per beat), analyze_roll (piano roll), analyze_melody , analyze_notes
Rhythm of level (pumping, gating)
analyze_envelope
What a sound is
analyze_timbre , analyze_spectrum , sound_compare
Vowels of a voice
analyze_formants
A picture, if you really need one
spectrogram (PNG)
Sources are render , track:<name> (after render(stems=True) ), ref , ref:<stem> , sound:<name> or a file path.
Bring your own reference audio ( project_new(..., reference=<file>) ); none is included here.
analyze_grid(source='ref') , then align a rendered drum track against ref:drums and correct offset_sec .
separate(source='ref') (demucs) and read analyze_structure(source='ref') .
Transcribe with notes_from_audio_loop . It keeps only notes that recur across repetitions of the loop, because raw transcription copies echoes, leakage and distortion partials as hard notes. If the song alternates versions of its loop, transcribe each from its own repetitions and pass base_ba

[truncated]

## Original Extract

A DAW for AI agents: write music as text, read the audio back as text. MCP server, CLI and an agent skill. - newsbubbles/ismail

GitHub - newsbubbles/ismail: A DAW for AI agents: write music as text, read the audio back as text. MCP server, CLI and an agent skill. · GitHub
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
newsbubbles
/
ismail
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.cursor .cursor .github/ workflows .github/ workflows ismail ismail skills/ ismail skills/ ismail tests tests .gitattributes .gitattributes .gitignore .gitignore .mcp.json .mcp.json LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
A DAW for AI agents. It can't hear, so it reads.
Listen to songs an agent made with it , each shown with the text the agent read while making it. The playhead runs across that text as the song plays.
Luigi Manson, made with ismail (fan remix of the Luigi's Mansion theme).
ismail is not a model that turns a prompt into audio, like Suno. It is a set of tools your own agent uses to write the song as notes, sounds and code, render it, read back what came out, and edit it. That changes what you get:
Iterative, precise edits. Change one note, one patch, one bar or one fader and re-render; nothing else moves.
Songs are code. A project file and a build script: git history, diffs, branches and code review work on a track.
Any sound. Synths, drum synths, samplers, voices written in Python and speech; any sound can become an instrument.
Local and open. MIT licensed, runs on your machine, no content filter and no music subscription (you bring the agent).
It improves with your model. The music is the agent's own work, so a stronger model with the same prompt should write a better song.
Suno is still better at realistic sung vocals, a polished song from one sentence in under a minute, and genre sound learned from recorded music. And why not Ableton or FL Studio? They were built for a person with ears and a mouse; an agent can press their buttons through bridges but still can't hear what it did. ismail puts everything an agent needs to write and to perceive into compact text, and if something is missing, your agent can add it. More on the showcase page .
A DAW built to be operated by an AI agent. Everything goes in as text (notes, instrument patches, effect chains, automation) and everything comes back as text: levels, spectra, drum patterns, piano rolls, chords, vowels, song structure, and structured comparisons against a reference track. The agent never needs ears or images to work (a spectrogram PNG is there if you want one).
One set of operations, three ways in:
MCP server for Claude Code, Cursor or any MCP client: python -m ismail.mcp_server (stdio, about 70 tools)
CLI : python -m ismail -p <project> <op> [args]
Python : from ismail import api
git clone https://github.com/newsbubbles/ismail
cd ismail
pip install -e . # engine, analysis, CLI, MCP server
pip install -e " .[perceptual] " # optional: CLAP perceptual metric (torch + transformers, model about 600 MB)
pip install -e " .[separate] " # optional: demucs stem separation for reference tracks
If demucs fights your torch install, use pip install --no-deps demucs and then pip install dora-search einops julius lameenc openunmix .
MP3 previews need ffmpeg on your PATH (or set ISMAIL_FFMPEG to the binary).
Runs on Windows, macOS and Linux; CI tests all three on every push. The one OS-specific op is sound_speak (text to speech for vocal samples), which uses the engine the OS already has:
Tools. Open Claude Code in this folder and the bundled .mcp.json registers the server; the tools show up as mcp__ismail__* . To use ismail from any folder instead:
claude mcp add -s user ismail -- python -m ismail.mcp_server
Skill (recommended). skills/ismail teaches the agent how to compose with ismail: plan a Session Sheet before writing notes, write a Listening Report after every render, and judge reference matches with the comparison tools instead of by feel. Link it into your skills folder:
# macOS / Linux
ln -s " $( pwd ) /skills/ismail " ~ /.claude/skills/ismail
# Windows
New-Item - ItemType Junction - Path " $ env: USERPROFILE \.claude\skills\ismail " - Target " $PWD \skills\ismail "
Ask for music. For example: "make a 16 bar deep house loop in F minor in songs/demo and render an mp3". The agent calls guide once for the conventions (it is a tool and a CLI op), then works through the tools.
Tools. Opening this folder in Cursor picks up .cursor/mcp.json . To use ismail in other projects, add the same entry to ~/.cursor/mcp.json :
{ "mcpServers" : { "ismail" : { "command" : " python " , "args" : [ " -m " , " ismail.mcp_server " ]}}}
Skill. .cursor/rules/ismail.mdc is an agent-requested rule that points Cursor's agent at skills/ismail/SKILL.md . Copy that rule (and the skills/ismail folder) into another project to use it there.
Any other MCP client works the same way: run python -m ismail.mcp_server over stdio.
Every tool is also a CLI op. Arguments are key=value pairs (values parsed as JSON when they can be) or one JSON object.
python -m ismail guide # read first: workflow and conventions
python -m ismail ops # list operations
python -m ismail help notes_write # one op's arguments and docs
python -m ismail -p songs/demo project_new bpm=124 length_bars=8
python -m ismail -p songs/demo track_add name=bass instrument= ' "preset:acid_bass" '
python -m ismail -p songs/demo notes_write ' {"track": "bass", "bar": 1, "notes": "0 E2 0.5 110; 0.5 E3 0.25", "repeat": 8} '
python -m ismail -p songs/demo render stems=true out=v1 mp3=also
python -m ismail -p songs/demo analyze_melody source=track:bass bars=[1,2]
render writes renders/latest.wav (every analysis tool reads it), plus renders/<out>.wav when you name the render. mp3='also' adds renders/<out>.mp3 for listening; mp3='only' writes the named render as mp3 only.
Keep your projects under songs/ (git-ignored) or anywhere else; a project is just a folder.
Project : a folder with project.json (tempo, grid offset, tracks, buses, master, sound bank, reference) plus sounds/ , renders/ , cache/ , history/ (undo snapshots) and comparisons/ .
Time : bars are 1-indexed; note times are beats relative to the bar you write at. offset_sec is the time of bar 1, so a project can sit exactly on a reference recording's grid.
Notes : '<beat> <pitch> <dur> [vel]' , one per line or ; -separated. Drum and step patterns: pattern_write with strings like X...x...X...x... (X 127, x 100, o 70, - 45, _ ties).
Instruments : synth (saw, square, pulse, triangle, sine, additive, wavetable and noise oscillators, unison, FM, drive, SVF and ladder filters, envelopes, LFOs, mono glide), sampler , drum synths ( kick , snare , hat , clap , tom , noise_hit ), kit (pitch to instrument map) and code (a Python voice function for anything else). presets_list has starting points.
Effects : eq, filter, distortion, bitcrush, compressor (with sidechain), duck, gate, delay, reverb, chorus, flanger, phaser, tremolo/autopan, width, limiter, vocoder, formant. Tracks, buses and the master fader can be automated.
Voices : engineered instruments kept as Python modules, so a project stores a name instead of code (see below).
Sound bank : sounds made from any instrument and effect chain ( sound_make ), speech ( sound_speak ), imported files, and averaged events cut from a recording ( sound_extract ). Bank sounds work as sampler sources, wavetables, vocoder modulators and audio clips.
Undo and batch : every edit snapshots the project ( undo ); batch applies a list of ops atomically.
Some instruments are easier to write than to patch: a measured grand piano, a dubstep bass whose note velocity picks the articulation, a set of sound effects. These live as voice modules, Python files that define voice(freq, t, vel, gate, sr) and return a mono (n,) or stereo (2, n) array. freq is in Hz, t is an array of seconds from the note start that covers the held time plus the instrument's tail , vel is 0 to 1, gate is how long the note is held in seconds, and sr is the sample rate.
Use one with instrument={"type": "code", "voice": "grand_piano", "tail": 4.0} or "preset:grand_piano" . voices_list shows what is available and voice_help(name) explains a voice's velocity mapping, functions and parameters.
Voices are looked up in this order:
<project>/voices/<name>.py : the song's own. Same name as a built-in overrides it; a song voice can also extend one ( from ismail.voices.growl import * , then add words or articulations).
Each folder in $ISMAIL_VOICES (a path list): your personal library, outside any repo.
ismail/voices/ : the built-ins.
A voice function may take extra keyword arguments: bpm is passed automatically, and the track's "params" dict is passed as keywords ( {"type": "code", "voice": "mine", "params": {"brightness": 0.3}} ). A module-level INFO dict documents it for voice_help ; every key is optional: summary (one line for voices_list ), range , velocity (what velocity does), functions (name to description), params (name to description) and tail (recommended tail). Data files sit next to the module ( grand_piano.json ) and are found through __file__ . Editing a voice file invalidates the render cache for the tracks that use it.
To add a voice to the library, move it from a song's voices/ folder into ismail/voices/ , give it an INFO dict, and add a line to the test that renders every built-in.
Question
Tool
Tempo and where bar 1 is
analyze_grid , align
Song form, what plays where
analyze_structure (arrangement map, sections, loop length, root per bar)
Levels, bands and chords per bar
analyze_bars , analyze_chords , analyze_key
Drum pattern
analyze_drums (step strings you can paste into pattern_write )
Notes
analyze_pitches (per beat), analyze_roll (piano roll), analyze_melody , analyze_notes
Rhythm of level (pumping, gating)
analyze_envelope
What a sound is
analyze_timbre , analyze_spectrum , sound_compare
Vowels of a voice
analyze_formants
A picture, if you really need one
spectrogram (PNG)
Sources are render , track:<name> (after render(stems=True) ), ref , ref:<stem> , sound:<name> or a file path.
Bring your own reference audio ( project_new(..., reference=<file>) ); none is included here.
analyze_grid(source='ref') , then align a rendered drum track against ref:drums and correct offset_sec .
separate(source='ref') (demucs) and read analyze_structure(source='ref') .
Transcribe with notes_from_audio_loop . It keeps only notes that recur across repetitions of the loop, because raw transcription copies echoes, leakage and distortion partials as hard notes. If the song alternates versions of its loop, transcribe each from its own repetitions and pass base_ba

[truncated]
