---
source: "https://www.collabora.com/news-and-blog/blog/2026/11/02/sunscreen-a-stylus-and-16-hours-of-pair-programming-with-an-ai/"
hn_url: "https://news.ycombinator.com/item?id=50340453"
title: "Sunscreen, a stylus, and 16 hours of pair programming with an AI"
article_title: "Sunscreen, a stylus, and 16 hours of pair programming with an AI"
image: "https://www.collabora.com/assets/images/blog/Collabora-Hand-Recognition-AI.jpg"
author: "losgehts"
captured_at: "2026-11-02T17:53:13Z"
capture_tool: "hn-digest"
hn_id: 50340453
score: 3
comments: 0
posted_at: "2026-11-02T17:38:25Z"
tags:
  - hacker-news
---

# Sunscreen, a stylus, and 16 hours of pair programming with an AI

- HN: [50340453](https://news.ycombinator.com/item?id=50340453)
- Source: [www.collabora.com](https://www.collabora.com/news-and-blog/blog/2026/11/02/sunscreen-a-stylus-and-16-hours-of-pair-programming-with-an-ai/)
- Score: 3
- Comments: 0
- Posted: 2026-11-02T17:38:25Z

## Translation

Title: Sunscreen, a stylus, and 16 hours of pair programming with an AI
Description: A prototype of handwriting recognition for KDE’s on-screen keyboard, built as an AI-assisted side project.

Article text:
Sunscreen, a stylus, and 16 hours of pair programming with an AI
Sunscreen, a stylus, and 16 hours of pair programming with an AI
Posted on 02/11/2026 by Kernel Bob
Posted on 02/11/2026 by Kernel Bob
My old laptop died on a trip, killed by a bottle of sunscreen that decided my luggage was a good place to empty itself. Annoying, but these things happen. The replacement was a Lenovo Yoga 7a 2-in-1 convertible with a stylus, chosen after the kind of research you do when you write kernel code for a living and have no intention of fighting your own hardware. Everything in it is supported out of the box. Suspend works. The stylus works. The screen rotates and the touchpad knows to stop listening. It is, honestly, a lovely machine to run Linux on in 2026.
And then I folded it into tablet mode, picked up the stylus, and discovered that KDE has no handwriting recognition.
Not "has a rough implementation". Not "has one that needs work". The previous attempts had been shelved years ago, and the current state of the art is that you fold your expensive convertible into tablet mode and then poke at an on-screen keyboard with one finger like it's 2009.
The hole where handwriting should be
KDE's new on-screen keyboard, plasma-keyboard , is built on Qt Virtual Keyboard, and Qt Virtual Keyboard does have a handwriting mode. There is a HandwritingModeKey . There are handwriting.qml layouts sitting right there in the source tree. It looks, at first glance, like a feature someone forgot to switch on.
It isn't. Availability is gated on QVirtualKeyboardFeatures::Handwriting , which is a compile-time constant . It is zero unless Qt was built against one of the proprietary engines — Cerence, or MyScript — or the example backend that emits random characters to prove the plumbing works. No open source build of Qt 6 ships with any of those, and the flag cannot be flipped from outside Qt. The layouts are dormant by design.
So the honest summary is: the feature is absent because the free software world does not have a handwriting engine that Qt is willing to ship. That is a much bigger hole than a missing checkbox, and it explains why the old attempts got shelved.
What can a developer do in that situation but roll up their sleeves?
The part where I admit I'm not fluent in KDE development
Here is my problem. I am a kernel developer. I am comfortable in C, in subsystem maintainer email threads, in the parts of the stack where things are measured in interrupts. I am not fluent in Qt Quick, QML, Wayland input-method protocols, or the internals of KWin. I know what those things are, but it's been over 4 years since I last looked at any QML or Qt internals. It was clear that while I knew the concepts, getting familiar with the code paths that tell a Wayland client that a key is held down, and how to produce that input after analyzing stylus inputs would take me a long time.
Ordinarily that is where a side project dies. Not because the idea is bad, but because the ratio of "things I'd have to learn" to "hours I actually have" is unfavorable, and there is always something else to do.
So I tried something different: I sat down with Claude and had a go at this whole "Vibe Coding" I'd been hearing so much about.
This post is about what that was actually like. Not the marketing version.
Step one: try to do it properly, and fail
My first instinct was the ambitious one. If the missing piece is an open source handwriting engine, then build an open source handwriting engine. Modern online handwriting recognition is well-trodden ground — bidirectional LSTM over stroke features, CTC loss; the architecture is not a mystery. The hard part is data.
I spent the first evening on a training spike. Generate synthetic handwriting from single-stroke vector fonts, extract cubic Bézier features from the stroke paths, train a BiLSTM with CTC, and see whether the thing learns (don't worry if some of those terms seem esoteric; I was aware of most of it from my university days far too many years ago now, but the AI had to educate me on a couple of more recent state-of-the-art from the intervening 20 years). I have a desktop with a Radeon RX 9070, so we set up ROCm and let it run overnight.
It worked, in the narrow sense. The model trained. The loss came down. On held-out samples from the fonts it had seen, it read them back.
Then we ran the experiment that mattered: we held out an entire font family and saw how it did on handwriting it had never seen the shape of. 82.93% character error rate. Which is to say: useless. And there are only four usable single-stroke font families in existence, so there was no path to fixing it with more data of the same kind.
Claude reported that clearly, without hedging, and recommended abandoning the approach. I want to flag that because it's the opposite of the failure mode I expected: it spent hours of compute on that path, then told me the experiment had failed and the result wasn't worth having. I got it to document the negative result up in the training repository and moved on.
That evening cost me maybe five hours and produced nothing shippable. It was still the right call — it's the difference between "we chose WritePad" and "we chose WritePad because the alternative demonstrably doesn't work", and I'd have wondered otherwise.
Step two: stand on someone else's shoulders
PhatWare's WritePad engine is GPL-3 and on GitHub. It is a mature cursive recognizer with an actual dictionary. I got it built, wrote an evaluation harness, and measured it on real cursive: 6.75% character error rate, 24.69% word error rate .
That is not perfect. It is enormously better than 82.93%, and — crucially — it is good enough that the guidelines and the size slider do more for accuracy than any model tweaking would have.
The license is a genuine consideration for upstream: linking GPL-3 code into a GPL-2-or-later project makes the resulting binary effectively GPL-3-only. I kept it behind a build option that defaults to off, and isolated it in its own commit, so upstream can take the handwriting mode and refuse the engine.
Four commits on a branch, about 5,000 lines across 32 files: A handwriting mode. A toggle key swaps the keyboard for a writing surface ruled like lined paper. You write on the lines, not between them, with room below each for descenders — so the tail of a "g" is still read as belonging to the line it was written on. Each row is recognized separately and joined with newlines, so a paragraph arrives in your document as a paragraph.
Pen input taken from the pen. More on this below; it's the best story in the project.
Editing keys and modifiers. Escape, Tab, Backspace, Delete, Return, cursor keys, and sticky Ctrl/Shift/Alt/AltGr around the writing area. Tap a modifier and it applies to the next thing you write; double-tap and it holds. Write a "c" with Ctrl armed and it sends Ctrl+C. A lot of these are still not working and I am not convinced that they need to stay.
Scribble to correct. Scribble through a word and it's deleted, leaving a gap; write the replacement in the gap and it goes back in the right place in the text. Scribble through a single letter and only that letter goes, so one misrecognized character doesn't cost you the whole word.
79 tests , covering the parts that are pure logic: gesture detection, the row geometry, the text assembly, the modifier state machine, the gap lifecycle.
Where the AI was genuinely, obviously better than me
Volume. 5,000 lines in two days including tests and commit messages that actually explain reasoning. I don't write that fast.
Reading unfamiliar code fast. At one point we needed to know exactly what KWin does with a modifier mask sent by an input method. Claude fetched the KWin source for the exact version I'm running, found keysymReceived , and showed me the lines that mattered:
cpp
seat()->notifyKeyboardModifiers(keyCode->modifiers, 0, 0, layout);
seat()->notifyKeyboardKey(keyCode->keyCode, Pressed, ...);
Every keysym an input method sends resets the modifier state to whatever that keysym alone needs, microseconds before the key goes out . Our Ctrl was being set and then wiped. The fix was to use the protocol's separate modifiers request with raw keycodes instead. I would have got there eventually. It would have taken me an evening, not ten minutes.
Instrumentation. When the stylus was drawing while hovering, I got it to add an event log and read 8,000 pointer events out of it. The pressure profile showed writing at 0.3–0.7 and hover sitting at a flat, latched 0.018 that decayed towards zero. That's not something you diagnose by staring at code. Having the idea, then having an AI that can trawl through the logs turns a potential "I give up, it's not worth the effort to fix that little bit" into a pleasurable exercise where I enjoy the problem solving and don't get frustrated with the donkey work required.
Being told "no" . Twice I asked for something and got a reasoned refusal with evidence, which saved me from myself. When I asked whether the Meta key could open the launcher, the answer was: no, and here's why — every key an input method sends goes to the focused window, and KWin's global shortcuts only see hardware input. It offered the D-Bus workaround and flagged that it would cost the Flatpak its first D-Bus permission. I decided to remove the button instead.
Where a human engineer was load-bearing
This is the part I think matters most for anyone evaluating this way of working.
Catching a wrong assumption stated as fact . At one point Claude set a flag called bSort on the recognition call and told me it "sorts strokes into reading order". Recognition quality immediately got worse and scribble deletion started failing on long words but not short ones. I said: I think maybe that bsort doesn't do what you think it does, please confirm your understanding .
It went and read the WritePad source. What it found was better than either of our guesses: the sort is skipped entirely below four strokes , which is exactly why short words worked and long ones didn't. That wasn't randomness; it was a threshold — and nobody would have found it by reasoning about the design.
That challenge took me one sentence. I could only make it because I'd noticed the shape of the failure was wrong: bugs that depend on word length are telling you something.
Noticing when it was hiding evidence from itself . Early on, a log file kept starting with (null): No such file or directory . Claude had been filtering that line out of its greps as noise. I pointed at it. It traced to a missing amdgpu.ids file, and — more importantly — the investigation exposed that a data-prep script was pulling in torch transitively, which mattered. Filtering noise is a reasonable instinct that occasionally throws away the signal.
Product judgment . Should the Meta key be a modifier or a key? Should scribbling a letter delete the whole word? Should the panel grow upward or stay fixed? Does the preview need reordering, or is it already right? Every one of those was a call I made in a sentence, and several of them stopped work that was about to be done unnecessarily. At one point I said, "I'm not sure you need to do the word ordering," and saved us an entire piece of machinery.
Precise bug reports . The best debugging I did all weekend was: "most times it didn't delete the letter at all, but instead seemed to add new letters in place of the scribble." That one sentence located the bug — if the scribble is appearing as letters , then the gesture detector returned false, which meant the problem was in detection, not deletion. Knowing what to notice is a skill that doesn't transfer to the machine yet.
Insisting on evidence over plausibility . This is the real lesson. Claude's characteristic failure mode, over two days, was reaching for a plausible fix instead of measuring . On the modifier problem, it made four reasonable-sounding attempts in a row before I pushed back and it went and read KWin's source. On the scrib

[truncated]

## Original Extract

A prototype of handwriting recognition for KDE’s on-screen keyboard, built as an AI-assisted side project.

Sunscreen, a stylus, and 16 hours of pair programming with an AI
Sunscreen, a stylus, and 16 hours of pair programming with an AI
Posted on 02/11/2026 by Kernel Bob
Posted on 02/11/2026 by Kernel Bob
My old laptop died on a trip, killed by a bottle of sunscreen that decided my luggage was a good place to empty itself. Annoying, but these things happen. The replacement was a Lenovo Yoga 7a 2-in-1 convertible with a stylus, chosen after the kind of research you do when you write kernel code for a living and have no intention of fighting your own hardware. Everything in it is supported out of the box. Suspend works. The stylus works. The screen rotates and the touchpad knows to stop listening. It is, honestly, a lovely machine to run Linux on in 2026.
And then I folded it into tablet mode, picked up the stylus, and discovered that KDE has no handwriting recognition.
Not "has a rough implementation". Not "has one that needs work". The previous attempts had been shelved years ago, and the current state of the art is that you fold your expensive convertible into tablet mode and then poke at an on-screen keyboard with one finger like it's 2009.
The hole where handwriting should be
KDE's new on-screen keyboard, plasma-keyboard , is built on Qt Virtual Keyboard, and Qt Virtual Keyboard does have a handwriting mode. There is a HandwritingModeKey . There are handwriting.qml layouts sitting right there in the source tree. It looks, at first glance, like a feature someone forgot to switch on.
It isn't. Availability is gated on QVirtualKeyboardFeatures::Handwriting , which is a compile-time constant . It is zero unless Qt was built against one of the proprietary engines — Cerence, or MyScript — or the example backend that emits random characters to prove the plumbing works. No open source build of Qt 6 ships with any of those, and the flag cannot be flipped from outside Qt. The layouts are dormant by design.
So the honest summary is: the feature is absent because the free software world does not have a handwriting engine that Qt is willing to ship. That is a much bigger hole than a missing checkbox, and it explains why the old attempts got shelved.
What can a developer do in that situation but roll up their sleeves?
The part where I admit I'm not fluent in KDE development
Here is my problem. I am a kernel developer. I am comfortable in C, in subsystem maintainer email threads, in the parts of the stack where things are measured in interrupts. I am not fluent in Qt Quick, QML, Wayland input-method protocols, or the internals of KWin. I know what those things are, but it's been over 4 years since I last looked at any QML or Qt internals. It was clear that while I knew the concepts, getting familiar with the code paths that tell a Wayland client that a key is held down, and how to produce that input after analyzing stylus inputs would take me a long time.
Ordinarily that is where a side project dies. Not because the idea is bad, but because the ratio of "things I'd have to learn" to "hours I actually have" is unfavorable, and there is always something else to do.
So I tried something different: I sat down with Claude and had a go at this whole "Vibe Coding" I'd been hearing so much about.
This post is about what that was actually like. Not the marketing version.
Step one: try to do it properly, and fail
My first instinct was the ambitious one. If the missing piece is an open source handwriting engine, then build an open source handwriting engine. Modern online handwriting recognition is well-trodden ground — bidirectional LSTM over stroke features, CTC loss; the architecture is not a mystery. The hard part is data.
I spent the first evening on a training spike. Generate synthetic handwriting from single-stroke vector fonts, extract cubic Bézier features from the stroke paths, train a BiLSTM with CTC, and see whether the thing learns (don't worry if some of those terms seem esoteric; I was aware of most of it from my university days far too many years ago now, but the AI had to educate me on a couple of more recent state-of-the-art from the intervening 20 years). I have a desktop with a Radeon RX 9070, so we set up ROCm and let it run overnight.
It worked, in the narrow sense. The model trained. The loss came down. On held-out samples from the fonts it had seen, it read them back.
Then we ran the experiment that mattered: we held out an entire font family and saw how it did on handwriting it had never seen the shape of. 82.93% character error rate. Which is to say: useless. And there are only four usable single-stroke font families in existence, so there was no path to fixing it with more data of the same kind.
Claude reported that clearly, without hedging, and recommended abandoning the approach. I want to flag that because it's the opposite of the failure mode I expected: it spent hours of compute on that path, then told me the experiment had failed and the result wasn't worth having. I got it to document the negative result up in the training repository and moved on.
That evening cost me maybe five hours and produced nothing shippable. It was still the right call — it's the difference between "we chose WritePad" and "we chose WritePad because the alternative demonstrably doesn't work", and I'd have wondered otherwise.
Step two: stand on someone else's shoulders
PhatWare's WritePad engine is GPL-3 and on GitHub. It is a mature cursive recognizer with an actual dictionary. I got it built, wrote an evaluation harness, and measured it on real cursive: 6.75% character error rate, 24.69% word error rate .
That is not perfect. It is enormously better than 82.93%, and — crucially — it is good enough that the guidelines and the size slider do more for accuracy than any model tweaking would have.
The license is a genuine consideration for upstream: linking GPL-3 code into a GPL-2-or-later project makes the resulting binary effectively GPL-3-only. I kept it behind a build option that defaults to off, and isolated it in its own commit, so upstream can take the handwriting mode and refuse the engine.
Four commits on a branch, about 5,000 lines across 32 files: A handwriting mode. A toggle key swaps the keyboard for a writing surface ruled like lined paper. You write on the lines, not between them, with room below each for descenders — so the tail of a "g" is still read as belonging to the line it was written on. Each row is recognized separately and joined with newlines, so a paragraph arrives in your document as a paragraph.
Pen input taken from the pen. More on this below; it's the best story in the project.
Editing keys and modifiers. Escape, Tab, Backspace, Delete, Return, cursor keys, and sticky Ctrl/Shift/Alt/AltGr around the writing area. Tap a modifier and it applies to the next thing you write; double-tap and it holds. Write a "c" with Ctrl armed and it sends Ctrl+C. A lot of these are still not working and I am not convinced that they need to stay.
Scribble to correct. Scribble through a word and it's deleted, leaving a gap; write the replacement in the gap and it goes back in the right place in the text. Scribble through a single letter and only that letter goes, so one misrecognized character doesn't cost you the whole word.
79 tests , covering the parts that are pure logic: gesture detection, the row geometry, the text assembly, the modifier state machine, the gap lifecycle.
Where the AI was genuinely, obviously better than me
Volume. 5,000 lines in two days including tests and commit messages that actually explain reasoning. I don't write that fast.
Reading unfamiliar code fast. At one point we needed to know exactly what KWin does with a modifier mask sent by an input method. Claude fetched the KWin source for the exact version I'm running, found keysymReceived , and showed me the lines that mattered:
cpp
seat()->notifyKeyboardModifiers(keyCode->modifiers, 0, 0, layout);
seat()->notifyKeyboardKey(keyCode->keyCode, Pressed, ...);
Every keysym an input method sends resets the modifier state to whatever that keysym alone needs, microseconds before the key goes out . Our Ctrl was being set and then wiped. The fix was to use the protocol's separate modifiers request with raw keycodes instead. I would have got there eventually. It would have taken me an evening, not ten minutes.
Instrumentation. When the stylus was drawing while hovering, I got it to add an event log and read 8,000 pointer events out of it. The pressure profile showed writing at 0.3–0.7 and hover sitting at a flat, latched 0.018 that decayed towards zero. That's not something you diagnose by staring at code. Having the idea, then having an AI that can trawl through the logs turns a potential "I give up, it's not worth the effort to fix that little bit" into a pleasurable exercise where I enjoy the problem solving and don't get frustrated with the donkey work required.
Being told "no" . Twice I asked for something and got a reasoned refusal with evidence, which saved me from myself. When I asked whether the Meta key could open the launcher, the answer was: no, and here's why — every key an input method sends goes to the focused window, and KWin's global shortcuts only see hardware input. It offered the D-Bus workaround and flagged that it would cost the Flatpak its first D-Bus permission. I decided to remove the button instead.
Where a human engineer was load-bearing
This is the part I think matters most for anyone evaluating this way of working.
Catching a wrong assumption stated as fact . At one point Claude set a flag called bSort on the recognition call and told me it "sorts strokes into reading order". Recognition quality immediately got worse and scribble deletion started failing on long words but not short ones. I said: I think maybe that bsort doesn't do what you think it does, please confirm your understanding .
It went and read the WritePad source. What it found was better than either of our guesses: the sort is skipped entirely below four strokes , which is exactly why short words worked and long ones didn't. That wasn't randomness; it was a threshold — and nobody would have found it by reasoning about the design.
That challenge took me one sentence. I could only make it because I'd noticed the shape of the failure was wrong: bugs that depend on word length are telling you something.
Noticing when it was hiding evidence from itself . Early on, a log file kept starting with (null): No such file or directory . Claude had been filtering that line out of its greps as noise. I pointed at it. It traced to a missing amdgpu.ids file, and — more importantly — the investigation exposed that a data-prep script was pulling in torch transitively, which mattered. Filtering noise is a reasonable instinct that occasionally throws away the signal.
Product judgment . Should the Meta key be a modifier or a key? Should scribbling a letter delete the whole word? Should the panel grow upward or stay fixed? Does the preview need reordering, or is it already right? Every one of those was a call I made in a sentence, and several of them stopped work that was about to be done unnecessarily. At one point I said, "I'm not sure you need to do the word ordering," and saved us an entire piece of machinery.
Precise bug reports . The best debugging I did all weekend was: "most times it didn't delete the letter at all, but instead seemed to add new letters in place of the scribble." That one sentence located the bug — if the scribble is appearing as letters , then the gesture detector returned false, which meant the problem was in detection, not deletion. Knowing what to notice is a skill that doesn't transfer to the machine yet.
Insisting on evidence over plausibility . This is the real lesson. Claude's characteristic failure mode, over two days, was reaching for a plausible fix instead of measuring . On the modifier problem, it made four reasonable-sounding attempts in a row before I pushed back and it went and read KWin's source. On the scrib

[truncated]
