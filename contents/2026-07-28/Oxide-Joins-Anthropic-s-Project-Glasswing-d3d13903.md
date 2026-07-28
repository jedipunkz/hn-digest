---
source: "https://oxide.computer/blog/oxide-anthropic-project-glasswing"
hn_url: "https://news.ycombinator.com/item?id=49082926"
title: "Oxide Joins Anthropic's Project Glasswing"
article_title: "Oxide Joins Anthropic's Project Glasswing | Oxide Computer Company"
author: "ErenayDev"
captured_at: "2026-07-28T12:48:14Z"
capture_tool: "hn-digest"
hn_id: 49082926
score: 4
comments: 0
posted_at: "2026-07-28T12:39:48Z"
tags:
  - hacker-news
  - translated
---

# Oxide Joins Anthropic's Project Glasswing

- HN: [49082926](https://news.ycombinator.com/item?id=49082926)
- Source: [oxide.computer](https://oxide.computer/blog/oxide-anthropic-project-glasswing)
- Score: 4
- Comments: 0
- Posted: 2026-07-28T12:39:48Z

## Translation

タイトル: Oxide が Anthropic のプロジェクト Glasswing に参加
記事のタイトル: Oxide が Anthropic のプロジェクト Glasswing に参加 |酸化物コンピューター株式会社
説明: Oxide は、Anthropic の Project Glasswing に参加し、Claude Mythos 5 を適用して、オープンソース スタック、ファームウェアからネットワーク全体にわたる脆弱性を見つけてパッチを適用します。

記事本文:
Oxide が Anthropic のプロジェクト Glasswing に参加 |オキサイドコンピューター社製品
オキサイドがAnthropicのプロジェクトGlasswingに参加
世界で最も重要なソフトウェアを保護するための共同作業に参加する
独自のコンピューティング インフラストラクチャを運用したことがある人は誰でも、ファームウェア、マイクロコード、組み込みコントローラー、オペレーティング システムを監査するという課題を経験したことがあります。これらのソフトウェアはすべて、悪意のある者の主な標的となります。この経験により、すべての重要なソフトウェアは検査可能であるべきだという私たちの信念が形成され、私たちが作成するすべてのものをオープンソースにする理由もここにあります。本日、私たちは世界で最も重要なソフトウェアを保護するための共同作業である Anthropic の Project Glasswing に参加します。このコラボレーションの一環として、Oxide は Claude Mythos 5 を適用して、ファームウェアからオペレーティング システム、コントロール プレーン、ネットワーク スタックに至るまで、自社のコードベースの潜在的な脆弱性を積極的に特定してパッチを適用しています。
Oxide におけるエンジニアリングの核となる信念は、ハードウェアとソフトウェアは共同設計されるべきであるということです。私たちは信頼のルートからファームウェア、オペレーティング システム、コントロール プレーンを経てネットワークに至るまで、あらゆる層を制御します。これが、Mythos が私たちにとって非常に役立つ理由です。私たちは第一原理から完全なスタックを設計したので、すべてのレイヤーに Mythos を指定し、見つかったものに基づいて迅速に行動することができます。
Mythos は、責任と厳格さを犠牲にすることなくサイバーセキュリティ防御を加速します。私たちは、NXP LPC55S69 の TrustZone におけるギャップの以前の発見を含め、セキュリティ作業をオープンに共有する責任を受け入れています。 RFD 576 では、大規模な言語モデルを使用して、困難な問題に対してエンジニアリングの厳密さをさらに高め、私たち自身の思考を研ぎ澄ます方法について説明します。
私たちは、不透明で脆弱性がたくさんあるインフラストラクチャの運用に何年も費やして Oxide を構築しました。

隠れる。透明で検証可能なインフラストラクチャを設計することは、私たちにとって基礎でした。その哲学は、Mythos を使用して独自のスタックの耐圧テストを行う方法にも適用されます。この作業をオープンに行うことは、重要なソフトウェアに依存するすべての組織に利益をもたらします。
Oxide の詳細については、https://oxide.computer をご覧ください。
Oxide への参加を申請するには、https://oxide.computer/careers にアクセスしてください。

## Original Extract

Oxide joins Anthropic's Project Glasswing, applying Claude Mythos 5 to find and patch vulnerabilities across its open-source stack, firmware to network.

Oxide Joins Anthropic's Project Glasswing | Oxide Computer Company Product
Oxide joins Anthropic's Project Glasswing
Joining a collaborative effort to secure the world’s most critical software
Everyone who has operated their own compute infrastructure has experienced the challenge of auditing firmware, microcode, embedded controllers, and operating systems — all software that are prime targets for bad actors. That experience shaped our belief that all critical software should be inspectable, and is why we open source everything we write. Today we’re joining Anthropic’s Project Glasswing, a collaborative effort to secure the world’s most important software. As part of this collaboration, Oxide is applying Claude Mythos 5 to proactively identify and patch potential vulnerabilities in our own codebase — from firmware through the operating system, control plane, and network stack.
Our core engineering belief at Oxide is that hardware and software should be co-designed — we control every layer, from the root of trust through firmware, operating system, and control plane, out to the network. This is what makes Mythos so useful to us: because we’ve designed the full stack from first principles, we can point it at every layer and act quickly on what it finds.
Mythos accelerates our cybersecurity defenses without sacrificing responsibility and rigor. We embrace the responsibility to share our security work in the open, including our earlier discovery of a gap in the NXP LPC55S69’s TrustZone . In RFD 576 , we describe how we use large language models to bring more engineering rigor to difficult problems and sharpen our own thinking.
We built Oxide having spent years operating infrastructure that was opaque and full of places for vulnerabilities to hide. Designing transparent and verifiable infrastructure was foundational for us. That philosophy extends to how we use Mythos to pressure-test our own stack. Doing this work in the open will benefit every organization that depends on critical software.
To learn more about Oxide, visit https://oxide.computer
To apply to join Oxide, visit https://oxide.computer/careers
