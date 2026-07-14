---
source: "https://engineering.taktile.com/blog/adapting-offensive-security-for-the-ai-agent-age/"
hn_url: "https://news.ycombinator.com/item?id=48905690"
title: "Adapting offensive security for the AI agent age"
article_title: "Adapting offensive security for the AI agent age"
author: "mmoon2"
captured_at: "2026-07-14T13:09:06Z"
capture_tool: "hn-digest"
hn_id: 48905690
score: 4
comments: 0
posted_at: "2026-07-14T12:23:40Z"
tags:
  - hacker-news
  - translated
---

# Adapting offensive security for the AI agent age

- HN: [48905690](https://news.ycombinator.com/item?id=48905690)
- Source: [engineering.taktile.com](https://engineering.taktile.com/blog/adapting-offensive-security-for-the-ai-agent-age/)
- Score: 4
- Comments: 0
- Posted: 2026-07-14T12:23:40Z

## Translation

タイトル: AI エージェント時代に攻撃的なセキュリティを適応させる

記事本文:
AI エージェント時代に攻撃的なセキュリティを適応させる ブログ 採用情報 会社 Jul 14, 2026 · 3 min read AI エージェント時代に攻撃的なセキュリティを適応させる
AI エージェントがより賢く、より有能になるにつれて、セキュリティは最大の不安の原因の 1 つとなっています。 Anthropic のようなフロンティア研究所は、悪意のある攻撃者がそれらを極悪なサイバーセキュリティ活動に使用することを防ぐために、最も有能なモデルへのアクセスを条件付けすることに頼っています。
Taktile では、セキュリティ脆弱性の検出から修復まで、製品の安全性を高めるためにこれらの新機能を使用してきました。また、以前は決定論的スキャナーをベースとしていたセキュリティ製品が AI に移行し、AI を活用して影響を評価し、新たな脆弱性を発見するケースが増えていることにも気づきました。
ただし、これらの製品の中には、既存のノイズが増大する危険性があるものもあります。エンジニアリング チームはすでに、週に数百枚のチケットを生成する複数のセキュリティ スキャナーからの出力に圧倒されており、これらのツールの一部は、非常に古いテクノロジーである静的スキャナーの上に非決定性レイヤーを固定しています。
私の考えでは、これは AI を活用する上で間違ったアプローチです。
AI によって自動化されたペネトレーション テストは、何のガイダンスも提供されなければ、すぐにただの SAST になってしまう可能性がありますが、AI が最も優れているのはまさにノイズを低減することです。
フロンティア ラボが目標主導型のループを出荷する中で、キャプチャ ザ フラッグ イベントのようにセキュリティ評価を扱うことが本当の価値があることに気付きました。それは、SAST/DAST を利用したセキュリティ スキャンから侵入テストに移行することでした。
侵入テストとセキュリティ スキャン
たとえて言えば、SAST/DAST スキャンはトローリング、侵入テストはトローリングと考えることができます。広い網を使って海域であらゆる種類の魚を捕まえますが、その探求では生物圏に意図しないダメージを与えたり、ゴミを引きずったりすることにもなります。

乗ってます。
侵入テストの目的は、必要なあらゆる手段を使って、セキュリティ境界を突破する方法を見つけることです。個人情報を抜き出すことができないとしたら、残念です。 「強化された調査結果」も「多層防御」も、仮説上の調査結果もありません。必要なのは、外部の脅威から内部システムの侵害を経て、PII の侵害などの実際の影響に至るまでの現実的な経路だけです。
それを念頭に置いて、私たちは目標駆動型ループの力を、ソース コードとアーキテクチャの形で利用できるコンテキストと組み合わせます。
「重大度の高いセキュリティ脆弱性を 1 つ発見しました。」
この目標を達成するには、エージェントに「高重大度」の定義、つまりどのロールに低い特権と高い特権があるのか、どのリソースが重要なのかを提供する必要があり、テナント間の脆弱性への重点を強化します。
このアプローチの基本的な部分は、単一の脆弱性というリズムです。
セキュリティ分野の AI エージェントの素晴らしい点は、単一の脆弱性に労力を集中することで、最終的にセキュリティ エンジニアや開発者が SAST/DAST の数百のエントリにわたる影響を評価することから解放されることです。これを毎週実行することにしたのは、セキュリティ エンジニア、開発者、関係者がこの脆弱性を利用し、翌週に次の脆弱性が見つかる前に修正を配布できるためです。
生産コストは安いが消費コストが高い、無限の AI 出力によって疲労が生じる時代において、私たちはこれらのツールを純粋な信号生成器として使用することを選択しました。
私たちと一緒に働くことに興味がありますか?当社の採用ページをご覧ください。
© 2026 タクタイル.無断転載を禁じます。

## Original Extract

Adapting offensive security for the AI agent age Blog Careers Company Jul 14, 2026 · 3 min read Adapting offensive security for the AI agent age
Security has been one of the biggest sources of anxiety as AI agents get smarter and more capable. Frontier labs like Anthropic have resorted to conditioning access to their most capable models in an effort to prevent malicious actors from using them for nefarious cybersecurity activities.
At Taktile we’ve been using these new capabilities to make our product safer: from security vulnerability detection to remediation. We’ve also noticed more and more security products previously based on deterministic scanners make the jump to AI and leverage it to assess impact and find new vulnerabilities.
Some of these products risk adding to the existing noise, however. Engineering teams are already overwhelmed with output from multiple security scanners generating hundreds of tickets a week, and some of these tools are stapling a non-deterministic layer on top of a very old technology: static scanners.
This is, in my view, the wrong approach to leveraging AI.
An AI-automated penetration test can quickly turn into just another SAST if offered no guidance - yet it is precisely in reducing noise that AI is at its best.
With frontier labs shipping goal-driven loops , I realized that treating security assessments like a Capture the Flag event was where the real value lay: shifting from SAST/DAST-powered security scanning to penetration testing.
Pentesting versus Security Scanning
As an analogy, we can think of SAST/DAST scanning as trawling and pentesting as trolling . One will use a wide net to catch all sorts of fish in an area of ocean, but in that quest it will also cause unintended damage to the biosphere and drag trash aboard.
Penetration testing aims to find a way past your security perimeter, by any means necessary. If we can’t exfiltrate private information, then too bad. No “hardening findings”, no “defense in depth”, and no hypothetical findings. We just need a realistic path from an external threat, through internal system compromise, to real impact such as compromised PII.
With that in mind, we combine the power of goal-driven loops with the context available to us in the form of source code and architecture.
” find one single high-severity security vulnerability. ”
In order for it to achieve that goal, we need to supply the agent with a definition of “high severity”: what roles are low and high privileged, what resources are critical, and we tighten the focus towards cross-tenant vulnerabilities.
The fundamental part of this approach is the cadence: a single vulnerability.
The beautiful thing about AI agents in security is that they can finally free security engineers and developers from assessing impact across hundreds of entries from a SAST/DAST by focusing their effort on a single vulnerability. We chose to run this weekly as it allows this vulnerability to be consumed by security engineers, developers and stakeholders, and ship a fix before the next one is found the following week.
In an age of fatigue caused by endless AI output that is cheap to produce but expensive to consume, we choose to use these tools as pure signal generators.
Interested in working with us? Check out our careers page .
© 2026 Taktile. All rights reserved.
