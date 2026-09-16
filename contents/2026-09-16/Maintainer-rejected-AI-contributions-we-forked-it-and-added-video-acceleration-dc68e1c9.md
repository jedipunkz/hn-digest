---
source: "https://twitter.com/SiliconForested/status/2100247223368024109"
hn_url: "https://news.ycombinator.com/item?id=49730521"
title: "Maintainer rejected AI contributions, we forked it and added video acceleration"
article_title: "Chris Kearney on X: \"🔥Omarchy Mac Users 🔥\nThe original maintainer didn’t want AI-generated contributions. Fair enough—we respected that boundary, forked the project, and got to work.\nThe result after roughly 21 hours:\n• 29 PRs merged across two repos\n• 115 new commits beyond the inherited base… / X"
image: "https://pbs.twimg.com/card_img/2099925015068299264/PRXvSzmr?format=webp&name=medium"
author: "mococa"
captured_at: "2026-09-16T18:42:22Z"
capture_tool: "hn-digest"
hn_id: 49730521
score: 14
comments: 1
posted_at: "2026-09-16T17:49:45Z"
tags:
  - hacker-news
---

# Maintainer rejected AI contributions, we forked it and added video acceleration

- HN: [49730521](https://news.ycombinator.com/item?id=49730521)
- Source: [twitter.com](https://twitter.com/SiliconForested/status/2100247223368024109)
- Score: 14
- Comments: 1
- Posted: 2026-09-16T17:49:45Z

## Translation

Title: Maintainer rejected AI contributions, we forked it and added video acceleration
Article title: Chris Kearney on X: "🔥Omarchy Mac Users 🔥
The original maintainer didn’t want AI-generated contributions. Fair enough—we respected that boundary, forked the project, and got to work.
The result after roughly 21 hours:
• 29 PRs merged across two repos
• 115 new commits beyond the inherited base… / X
Description: 🔥Omarchy Mac Users 🔥
The original maintainer didn’t want AI-generated contributions. Fair enough—we respected that boundary, forked the project, and got to work.
The result after roughly 21 hours:
• 29 PRs merged across two repos
• 115 new commits beyond the inherited baseline
• 7 PR authors c…

Article text:
Chris Kearney on X: "🔥Omarchy Mac Users 🔥
The original maintainer didn’t want AI-generated contributions. Fair enough—we respected that boundary, forked the project, and got to work.
The result after roughly 21 hours:
• 29 PRs merged across two repos
• 115 new commits beyond the inherited baseline
• 7 PR authors contributing in parallel
We now have working Apple Video Decoder acceleration for H.264, HEVC and VP9 on M-series Macs running Asahi/Omarchy, plus opt-in H.264 High 10.
The work includes:
• Chrome’s solid-green-video bug fixed
• 10-bit HEVC blank-frame bug fixed
• HEVC tile and wavefront hangs fixed
• H.264 multi-slice firmware hangs fixed
• VP9 state, color-range and submission handling fixed
• Decoder lifecycle races and cross-context reference aliasing fixed
• Multiple memory-safety, error-handling and malformed-stream bugs fixed
• 15 AVD kernel patches
• Automatic installation, recovery and kernel-update rebuilding
• Sanitizer CI, guarded hardware testing and reproducible conformance infrastructure
On real M1 hardware we’re currently bit-exact on:
• HEVC: 144/147 streams
• AVC: 73/135
• VP9: 216/305
• 864 additional generated hardware-frame comparisons
Huge credit to the contributors: davefano, grudev, laihenyi, stefanakerwall, thebytorsnowdog and service-cyber—as well as the upstream and carried work from megi, sofus13, Igor Ryzhkov, Ante042, Aaron/aquarat and the Asahi Linux contributors.
This is what agent-native open-source development can look like: humans set direction and validate reality; agents attack a shared roadmap in parallel; every claim is tied to tests and hardware evidence.
We didn’t argue about whether AI belongs in open source.
We forked it and demonstrated what it can do.
https://t.co/xzyqrvPoaf
https://t.co/c2cluKorX3"
@SiliconForested 🔥Omarchy Mac Users 🔥
The original maintainer didn’t want AI-generated contributions. Fair enough—we respected that boundary, forked the project, and got to work.
The result after roughly 21 hours:
• 29 PRs merged across two repos
• 115 new commits beyond the inherited baseline
• 7 PR authors contributing in parallel
We now have working Apple Video Decoder acceleration for H.264, HEVC and VP9 on M-series Macs running Asahi/Omarchy, plus opt-in H.264 High 10.
The work includes:
• Chrome’s solid-green-video bug fixed
• 10-bit HEVC blank-frame bug fixed
• HEVC tile and wavefront hangs fixed
• H.264 multi-slice firmware hangs fixed
• VP9 state, color-range and submission handling fixed
• Decoder lifecycle races and cross-context reference aliasing fixed
• Multiple memory-safety, error-handling and malformed-stream bugs fixed
• 15 AVD kernel patches
• Automatic installation, recovery and kernel-update rebuilding
• Sanitizer CI, guarded hardware testing and reproducible conformance infrastructure
On real M1 hardware we’re currently bit-exact on:
• HEVC: 144/147 streams
• AVC: 73/135
• VP9: 216/305
• 864 additional generated hardware-frame comparisons
Huge credit to the contributors: davefano, grudev, laihenyi, stefanakerwall, thebytorsnowdog and service-cyber—as well as the upstream and carried work from megi, sofus13, Igor Ryzhkov, Ante042, Aaron/aquarat and the Asahi Linux contributors.
This is what agent-native open-source development can look like: humans set direction and validate reality; agents attack a shared roadmap in parallel; every claim is tied to tests and hardware evidence.
We didn’t argue about whether AI belongs in open source.
We forked it and demonstrated what it can do.
github.com/iconidentify/l…
github.com/iconidentify/o… Chris Kearney @SiliconForested 17h 💪 OMARCHY MAC USERS!
Got some tokens to spare? Put them toward making Omarchy on M-series Macs better!
I’ve been working with agents to investigate bugs, missing features, and the st
[truncated]
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
© 2026 X Corp. Chris Kearney @SiliconForested 🔥Omarchy Mac Users 🔥
The original maintainer didn’t want AI-generated contributions. Fair enough—we respected that boundary, forked the project, and got to work.
The result after roughly 21 hours:
• 29 PRs merged across two repos
• 115 new commits beyond the inherited baseline
• 7 PR authors contributing in parallel
We now have working Apple Video Decoder acceleration for H.264, HEVC and VP9 on M-series Macs running Asahi/Omarchy, plus opt-in H.264 High 10.
The work includes:
• Chrome’s solid-green-video bug fixed
• 10-bit HEVC blank-frame bug fixed
• HEVC tile and wavefront hangs fixed
• H.264 multi-slice firmware hangs fixed
• VP9 state, color-range and submission handling fixed
• Decoder lifecycle races and cross-context reference aliasing fixed
• Multiple memory-safety, error-handling and malformed-stream bugs fixed
• 15 AVD kernel patches
• Automatic installation, recovery and kernel-update rebuilding
• Sanitizer CI, guarded hardware testing and reproducible conformance infrastructure
On real M1 hardware we’re currently bit-exact on:
• HEVC: 144/147 streams
• AVC: 73/135
• VP9: 216/305
• 864 additional generated hardware-frame comparisons
Huge credit to the contributors: davefano, grudev, laihenyi, stefanakerwall, thebytorsnowdog and service-cyber—as well as the upstream and carried work from megi, sofus13, Igor Ryzhkov, Ante042, Aaron/aquarat and the Asahi Linux contributors.
This is what agent-native open-source development can look like: humans set direction and validate reality; agents attack a shared roadmap in parallel; every claim is tied to tests and hardware evidence.
We didn’t argue about whether AI belongs in open source.
We forked it and demonstrated what it can do.
github.com/iconidentify/l…
github.com/iconidentify/o… Chris Kearney @SiliconForested 17h 💪 OMARCHY MAC USERS!
Got some tokens to spare? Put them toward making Omarchy on M-series Macs better!
I’ve been working with agents to investigate bugs,
[truncated]
Jeff Davis, MD @jeffdavismd 1h Sayin you don’t want AI generated contributions in 2026 is like saying we’ll only accept PRs on punch cards. 1 1 15 255
Jeremy @ocdjeremy 3h Fork it, we'll do it live!
Marvelous. Simply marvelous. 1 13 374
Omarchy Mac @OmarchyMac 1h GIF 10 223

## Original Extract

🔥Omarchy Mac Users 🔥
The original maintainer didn’t want AI-generated contributions. Fair enough—we respected that boundary, forked the project, and got to work.
The result after roughly 21 hours:
• 29 PRs merged across two repos
• 115 new commits beyond the inherited baseline
• 7 PR authors c…

Chris Kearney on X: "🔥Omarchy Mac Users 🔥
The original maintainer didn’t want AI-generated contributions. Fair enough—we respected that boundary, forked the project, and got to work.
The result after roughly 21 hours:
• 29 PRs merged across two repos
• 115 new commits beyond the inherited baseline
• 7 PR authors contributing in parallel
We now have working Apple Video Decoder acceleration for H.264, HEVC and VP9 on M-series Macs running Asahi/Omarchy, plus opt-in H.264 High 10.
The work includes:
• Chrome’s solid-green-video bug fixed
• 10-bit HEVC blank-frame bug fixed
• HEVC tile and wavefront hangs fixed
• H.264 multi-slice firmware hangs fixed
• VP9 state, color-range and submission handling fixed
• Decoder lifecycle races and cross-context reference aliasing fixed
• Multiple memory-safety, error-handling and malformed-stream bugs fixed
• 15 AVD kernel patches
• Automatic installation, recovery and kernel-update rebuilding
• Sanitizer CI, guarded hardware testing and reproducible conformance infrastructure
On real M1 hardware we’re currently bit-exact on:
• HEVC: 144/147 streams
• AVC: 73/135
• VP9: 216/305
• 864 additional generated hardware-frame comparisons
Huge credit to the contributors: davefano, grudev, laihenyi, stefanakerwall, thebytorsnowdog and service-cyber—as well as the upstream and carried work from megi, sofus13, Igor Ryzhkov, Ante042, Aaron/aquarat and the Asahi Linux contributors.
This is what agent-native open-source development can look like: humans set direction and validate reality; agents attack a shared roadmap in parallel; every claim is tied to tests and hardware evidence.
We didn’t argue about whether AI belongs in open source.
We forked it and demonstrated what it can do.
https://t.co/xzyqrvPoaf
https://t.co/c2cluKorX3"
@SiliconForested 🔥Omarchy Mac Users 🔥
The original maintainer didn’t want AI-generated contributions. Fair enough—we respected that boundary, forked the project, and got to work.
The result after roughly 21 hours:
• 29 PRs merged across two repos
• 115 new commits beyond the inherited baseline
• 7 PR authors contributing in parallel
We now have working Apple Video Decoder acceleration for H.264, HEVC and VP9 on M-series Macs running Asahi/Omarchy, plus opt-in H.264 High 10.
The work includes:
• Chrome’s solid-green-video bug fixed
• 10-bit HEVC blank-frame bug fixed
• HEVC tile and wavefront hangs fixed
• H.264 multi-slice firmware hangs fixed
• VP9 state, color-range and submission handling fixed
• Decoder lifecycle races and cross-context reference aliasing fixed
• Multiple memory-safety, error-handling and malformed-stream bugs fixed
• 15 AVD kernel patches
• Automatic installation, recovery and kernel-update rebuilding
• Sanitizer CI, guarded hardware testing and reproducible conformance infrastructure
On real M1 hardware we’re currently bit-exact on:
• HEVC: 144/147 streams
• AVC: 73/135
• VP9: 216/305
• 864 additional generated hardware-frame comparisons
Huge credit to the contributors: davefano, grudev, laihenyi, stefanakerwall, thebytorsnowdog and service-cyber—as well as the upstream and carried work from megi, sofus13, Igor Ryzhkov, Ante042, Aaron/aquarat and the Asahi Linux contributors.
This is what agent-native open-source development can look like: humans set direction and validate reality; agents attack a shared roadmap in parallel; every claim is tied to tests and hardware evidence.
We didn’t argue about whether AI belongs in open source.
We forked it and demonstrated what it can do.
github.com/iconidentify/l…
github.com/iconidentify/o… Chris Kearney @SiliconForested 17h 💪 OMARCHY MAC USERS!
Got some tokens to spare? Put them toward making Omarchy on M-series Macs better!
I’ve been working with agents to investigate bugs, missing features, and the st
[truncated]
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
© 2026 X Corp. Chris Kearney @SiliconForested 🔥Omarchy Mac Users 🔥
The original maintainer didn’t want AI-generated contributions. Fair enough—we respected that boundary, forked the project, and got to work.
The result after roughly 21 hours:
• 29 PRs merged across two repos
• 115 new commits beyond the inherited baseline
• 7 PR authors contributing in parallel
We now have working Apple Video Decoder acceleration for H.264, HEVC and VP9 on M-series Macs running Asahi/Omarchy, plus opt-in H.264 High 10.
The work includes:
• Chrome’s solid-green-video bug fixed
• 10-bit HEVC blank-frame bug fixed
• HEVC tile and wavefront hangs fixed
• H.264 multi-slice firmware hangs fixed
• VP9 state, color-range and submission handling fixed
• Decoder lifecycle races and cross-context reference aliasing fixed
• Multiple memory-safety, error-handling and malformed-stream bugs fixed
• 15 AVD kernel patches
• Automatic installation, recovery and kernel-update rebuilding
• Sanitizer CI, guarded hardware testing and reproducible conformance infrastructure
On real M1 hardware we’re currently bit-exact on:
• HEVC: 144/147 streams
• AVC: 73/135
• VP9: 216/305
• 864 additional generated hardware-frame comparisons
Huge credit to the contributors: davefano, grudev, laihenyi, stefanakerwall, thebytorsnowdog and service-cyber—as well as the upstream and carried work from megi, sofus13, Igor Ryzhkov, Ante042, Aaron/aquarat and the Asahi Linux contributors.
This is what agent-native open-source development can look like: humans set direction and validate reality; agents attack a shared roadmap in parallel; every claim is tied to tests and hardware evidence.
We didn’t argue about whether AI belongs in open source.
We forked it and demonstrated what it can do.
github.com/iconidentify/l…
github.com/iconidentify/o… Chris Kearney @SiliconForested 17h 💪 OMARCHY MAC USERS!
Got some tokens to spare? Put them toward making Omarchy on M-series Macs better!
I’ve been working with agents to investigate bugs,
[truncated]
Jeff Davis, MD @jeffdavismd 1h Sayin you don’t want AI generated contributions in 2026 is like saying we’ll only accept PRs on punch cards. 1 1 15 255
Jeremy @ocdjeremy 3h Fork it, we'll do it live!
Marvelous. Simply marvelous. 1 13 374
Omarchy Mac @OmarchyMac 1h GIF 10 223
