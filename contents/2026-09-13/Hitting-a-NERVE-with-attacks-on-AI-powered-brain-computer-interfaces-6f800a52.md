---
source: "https://anil.recoil.org/notes/nerve-attacks"
hn_url: "https://news.ycombinator.com/item?id=49688047"
title: "Hitting a NERVE with attacks on AI-powered brain-computer interfaces"
article_title: "Hitting a NERVE with attacks on AI-powered brain-computer interfaces | Anil Madhavapeddy"
image: "https://anil.recoil.org/images/papers/2026-nerve-attacks.640.webp"
author: "matt_d"
captured_at: "2026-09-13T20:28:14Z"
capture_tool: "hn-digest"
hn_id: 49688047
score: 3
comments: 1
posted_at: "2026-09-13T19:59:52Z"
tags:
  - hacker-news
---

# Hitting a NERVE with attacks on AI-powered brain-computer interfaces

- HN: [49688047](https://news.ycombinator.com/item?id=49688047)
- Source: [anil.recoil.org](https://anil.recoil.org/notes/nerve-attacks)
- Score: 3
- Comments: 1
- Posted: 2026-09-13T19:59:52Z

## Translation

Title: Hitting a NERVE with attacks on AI-powered brain-computer interfaces
Article title: Hitting a NERVE with attacks on AI-powered brain-computer interfaces | Anil Madhavapeddy
Description: New preprint defining attack dimensions for Brain Computer Interfaces and an analysis framework called EEGle.

Article text:
>_ Anil Madhavapeddy @avsm / About
Subscribe Atom full JSON Feed full Atom perma JSON Feed perma >_ Anil Madhavapeddy About Projects Ideas Papers Notes Talks Network Links Hitting a NERVE with attacks on AI-powered brain-computer interfaces
New preprint defining attack dimensions for Brain Computer Interfaces and an analysis framework called EEGle.
Human brain interfaces are advancing at a startling pace; Ray Kurzweil has just
joined a startup that dispenses with skull surgery
in favour of you snorting charged nanoparticles that are then driven by
magnetic coils in a cap.
This all sounds very scifi, but the software stack on these BCI devices
have an enormous and unguarded attack surface. Zahra Tarkhani has just uploaded a preprnt NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces that we worked on with
two brilliant interns who visited last summer, Georgios Akkogiounoglou (KTH) and Isabel Tscherniak (TUM),
and Lorena Qendro at Nokia Bell Labs.
Back in 2022, our first pass at this found over
300 vulnerabilities across the stacks of commodity BCI headsets.
The new paper categorises this surface into what we dub the NERVE attacks , which
between them cover the entire Brain-Computer stack from the radio interface to the underlying trained model:
N euro-mimetic Forgery synthesises a physiologically plausible brain
signal that the classifier accepts without any data from the user.
E vasion via Desynchronization nudges the timing of the signal by a
fraction of a second so discriminative features fall outside the model's
receptive window.
R eplay-based Hijacking records a real epoch and replays it over the
wireless link to reissue commands without the user knowing.
V ein Tapping the unencrypted radio links, finding missing authentication,
exporting world-readable model files, and (of course) memory-unsafe SDKs that
expose everything else to the sort of buffer overflow
I was chasing over twenty years ago.
E mbedded Backdoors via trojaned models that behave normally until they see
a trigger disguised as an eye blink, a muscle twitch or even mains interference.
In order to help find these vulnerability points, we built 'EEGle' (no relation to my Energy and Environment Group !). EEGle's an extensible framework for BCI security analysis that exposed 17 novel neuro-specific attacks.
And like other areas recently LLMs remove the need for any expertise in neuroscience and
make it super easy to supply the physiological parameters for a forged motor-imagery signal.
The same models are just as happy scanning for the flaws as finding ways through them.
The end-to-end demonstration in the paper against the CYBATHLON 2024 BCI game
is pretty cool. This game has three tasks: cursor control, wheelchair navigation, and operating an ice machine with a robotic arm.
All three were hijacked via raw replay, synthesised epochs and augmented replay in about 8 seconds.
The defences aren't in good shape unfortunately. There's no (published) protection against some of these such as the evasion/desynchronisation signal attacks,
and the information flow control approach we proposed back in 2022 only covers the systems half of the problem.
The reason it's all so worrying is that chained attacks are really
easy to execute these days, several of which are shown in the paper and often take <10ms.
In the ice machine task, a misclassification tips the robotic arm over, showing the physical stakes involved given
the same mechanisms are intended to be used for wheelchairs and prosthetics .
So there's certainly some brainy work to be done here! The preprint is on arXiv and comments are very welcome.
Read more about NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces .
Tarkhani et al (2026). NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces. arXiv. 10.48550/arXiv.2609.08971 [2] Madhavapeddy (2026). Just a rumour of a bug is enough to find a security exploit these days. 10.59350/tngsm-6rx23 [3] Tarkhani et al (2022). Enhancing the Security & Privacy of Wearable Brain-Computer Interfaces. arXiv. 10.48550/arXiv.2201.07711 [4] Madhavapeddy (2025). EEG internships for the summer of 2025. 10.59350/tf22g-p1822 [5] Madhavapeddy (2026). The Internet needs an antibotty immune system, stat. 10.59350/snnnf-asc02 Related
NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces Sep 2026 Zahra Tarkhani, Georgios Akkogiounoglou et al. Security scanning my own code with Scrutineer and local coding models Aug 2026 Deploying Alpha-Omega's Scrutineer on a VM, driving its scans with a local GLM 5.3 on my Mac Studio, and thinking through practical workflows Just a rumour of a bug is enough to find a security exploit these days Aug 2026 Thinking through how the conventional OSS security embargoes no longer buy us time, and what open source maintainers might do instead to respond The Internet needs an antibotty immune system, stat Apr 2026 Anthropic's Mythos makes autonomous vulnerability chaining across devices a sudden reality, so I've been thinking about how digital 'antibotty' inoculation networks may be needed far sooner than I expected. EEG internships for the summer of 2025 Jun 2025 Coordination note for summer 2025 undergraduate and graduate internships covering projects from evidence databases to remote sensing and embedded systems. Enhancing the Security & Privacy of Wearable Brain-Computer Interfaces Jan 2022 Zahra Tarkhani, Lorena Qendro et al. Security analysis of brain-computing interfaces Jan 2021 Completed Information Flow for Trusted Execution Jan 2020 >_ nerve-attacks New preprint defining attack dimensions for Brain Computer Interfaces and an analysis framework called EEGle.
NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces Sep 2026 Security scanning my own code with Scrutineer and local coding models Aug 2026 Just a rumour of a bug is enough to find a security exploit these days Aug 2026 The Internet needs an antibotty immune system, stat Apr 2026 EEG internships for the summer of 2025 Jun 2025 Enhancing the Security & Privacy of Wearable Brain-Computer Interfaces Jan 2022 Security analysis of brain-computing interfaces Jan 2021 Information Flow for Trusted Execution Jan 2020 Enhancing the Security & Privacy of Wearable Brain-Computer Interfaces · Tarkhani et al (2022) · DOI Zahra Tarkhani NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces · Tarkhani et al (2026) · DOI Georgios Akkogiounoglou Isabel Tscherniak Lorena Qendro Security analysis of brain-computing interfaces · Completed · Any EEG internships for the summer of 2025 · 2025 · 1359w Just a rumour of a bug is enough to find a security exploit these days · 2026 · 1868w Security scanning my own code with Scrutineer and local coding models · 2026 · 1541w Information Flow for Trusted Execution · 2020–present The Internet needs an antibotty immune system, stat · 2026 · 1450w © 1998–2026 Anil Madhavapeddy.

## Original Extract

New preprint defining attack dimensions for Brain Computer Interfaces and an analysis framework called EEGle.

>_ Anil Madhavapeddy @avsm / About
Subscribe Atom full JSON Feed full Atom perma JSON Feed perma >_ Anil Madhavapeddy About Projects Ideas Papers Notes Talks Network Links Hitting a NERVE with attacks on AI-powered brain-computer interfaces
New preprint defining attack dimensions for Brain Computer Interfaces and an analysis framework called EEGle.
Human brain interfaces are advancing at a startling pace; Ray Kurzweil has just
joined a startup that dispenses with skull surgery
in favour of you snorting charged nanoparticles that are then driven by
magnetic coils in a cap.
This all sounds very scifi, but the software stack on these BCI devices
have an enormous and unguarded attack surface. Zahra Tarkhani has just uploaded a preprnt NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces that we worked on with
two brilliant interns who visited last summer, Georgios Akkogiounoglou (KTH) and Isabel Tscherniak (TUM),
and Lorena Qendro at Nokia Bell Labs.
Back in 2022, our first pass at this found over
300 vulnerabilities across the stacks of commodity BCI headsets.
The new paper categorises this surface into what we dub the NERVE attacks , which
between them cover the entire Brain-Computer stack from the radio interface to the underlying trained model:
N euro-mimetic Forgery synthesises a physiologically plausible brain
signal that the classifier accepts without any data from the user.
E vasion via Desynchronization nudges the timing of the signal by a
fraction of a second so discriminative features fall outside the model's
receptive window.
R eplay-based Hijacking records a real epoch and replays it over the
wireless link to reissue commands without the user knowing.
V ein Tapping the unencrypted radio links, finding missing authentication,
exporting world-readable model files, and (of course) memory-unsafe SDKs that
expose everything else to the sort of buffer overflow
I was chasing over twenty years ago.
E mbedded Backdoors via trojaned models that behave normally until they see
a trigger disguised as an eye blink, a muscle twitch or even mains interference.
In order to help find these vulnerability points, we built 'EEGle' (no relation to my Energy and Environment Group !). EEGle's an extensible framework for BCI security analysis that exposed 17 novel neuro-specific attacks.
And like other areas recently LLMs remove the need for any expertise in neuroscience and
make it super easy to supply the physiological parameters for a forged motor-imagery signal.
The same models are just as happy scanning for the flaws as finding ways through them.
The end-to-end demonstration in the paper against the CYBATHLON 2024 BCI game
is pretty cool. This game has three tasks: cursor control, wheelchair navigation, and operating an ice machine with a robotic arm.
All three were hijacked via raw replay, synthesised epochs and augmented replay in about 8 seconds.
The defences aren't in good shape unfortunately. There's no (published) protection against some of these such as the evasion/desynchronisation signal attacks,
and the information flow control approach we proposed back in 2022 only covers the systems half of the problem.
The reason it's all so worrying is that chained attacks are really
easy to execute these days, several of which are shown in the paper and often take <10ms.
In the ice machine task, a misclassification tips the robotic arm over, showing the physical stakes involved given
the same mechanisms are intended to be used for wheelchairs and prosthetics .
So there's certainly some brainy work to be done here! The preprint is on arXiv and comments are very welcome.
Read more about NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces .
Tarkhani et al (2026). NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces. arXiv. 10.48550/arXiv.2609.08971 [2] Madhavapeddy (2026). Just a rumour of a bug is enough to find a security exploit these days. 10.59350/tngsm-6rx23 [3] Tarkhani et al (2022). Enhancing the Security & Privacy of Wearable Brain-Computer Interfaces. arXiv. 10.48550/arXiv.2201.07711 [4] Madhavapeddy (2025). EEG internships for the summer of 2025. 10.59350/tf22g-p1822 [5] Madhavapeddy (2026). The Internet needs an antibotty immune system, stat. 10.59350/snnnf-asc02 Related
NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces Sep 2026 Zahra Tarkhani, Georgios Akkogiounoglou et al. Security scanning my own code with Scrutineer and local coding models Aug 2026 Deploying Alpha-Omega's Scrutineer on a VM, driving its scans with a local GLM 5.3 on my Mac Studio, and thinking through practical workflows Just a rumour of a bug is enough to find a security exploit these days Aug 2026 Thinking through how the conventional OSS security embargoes no longer buy us time, and what open source maintainers might do instead to respond The Internet needs an antibotty immune system, stat Apr 2026 Anthropic's Mythos makes autonomous vulnerability chaining across devices a sudden reality, so I've been thinking about how digital 'antibotty' inoculation networks may be needed far sooner than I expected. EEG internships for the summer of 2025 Jun 2025 Coordination note for summer 2025 undergraduate and graduate internships covering projects from evidence databases to remote sensing and embedded systems. Enhancing the Security & Privacy of Wearable Brain-Computer Interfaces Jan 2022 Zahra Tarkhani, Lorena Qendro et al. Security analysis of brain-computing interfaces Jan 2021 Completed Information Flow for Trusted Execution Jan 2020 >_ nerve-attacks New preprint defining attack dimensions for Brain Computer Interfaces and an analysis framework called EEGle.
NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces Sep 2026 Security scanning my own code with Scrutineer and local coding models Aug 2026 Just a rumour of a bug is enough to find a security exploit these days Aug 2026 The Internet needs an antibotty immune system, stat Apr 2026 EEG internships for the summer of 2025 Jun 2025 Enhancing the Security & Privacy of Wearable Brain-Computer Interfaces Jan 2022 Security analysis of brain-computing interfaces Jan 2021 Information Flow for Trusted Execution Jan 2020 Enhancing the Security & Privacy of Wearable Brain-Computer Interfaces · Tarkhani et al (2022) · DOI Zahra Tarkhani NERVE Attacks: Breaking AI-Powered Brain-Computer Interfaces · Tarkhani et al (2026) · DOI Georgios Akkogiounoglou Isabel Tscherniak Lorena Qendro Security analysis of brain-computing interfaces · Completed · Any EEG internships for the summer of 2025 · 2025 · 1359w Just a rumour of a bug is enough to find a security exploit these days · 2026 · 1868w Security scanning my own code with Scrutineer and local coding models · 2026 · 1541w Information Flow for Trusted Execution · 2020–present The Internet needs an antibotty immune system, stat · 2026 · 1450w © 1998–2026 Anil Madhavapeddy.
