---
source: "https://worklifenotes.com/2026/06/12/the-last-line-of-defense-must-not-be-ai/"
hn_url: "https://news.ycombinator.com/item?id=48503245"
title: "The last line of defense must not be AI"
article_title: "The last line of defense must not be AI - Work & Life Notes"
author: "taleodor"
captured_at: "2026-06-12T13:17:14Z"
capture_tool: "hn-digest"
hn_id: 48503245
score: 2
comments: 2
posted_at: "2026-06-12T12:24:58Z"
tags:
  - hacker-news
  - translated
---

# The last line of defense must not be AI

- HN: [48503245](https://news.ycombinator.com/item?id=48503245)
- Source: [worklifenotes.com](https://worklifenotes.com/2026/06/12/the-last-line-of-defense-must-not-be-ai/)
- Score: 2
- Comments: 2
- Posted: 2026-06-12T12:24:58Z

## Translation

タイトル: 最終防衛ラインは AI であってはなりません
記事タイトル: 最終防衛ラインは AI であってはなりません - Work & Life Notes
説明: 大規模な作業を行う AI をどのように管理するかという質問に対して、頻繁に広まっている答えは、「AI は完全にダウンする」です。 AIが増えるということは

記事本文:
コンテンツにスキップ
仕事と生活のメモ
Pavel Shukhman のブログ: テクノロジー、旅行、人間関係、生活
最終防衛線はAIであってはなりません
大規模な作業を行う AI をどのように管理するかという質問に対してよく出回っている答えは、「AI は完全にダウンする」です。つまり、下流側の AI が増えれば、上流側で発生したあらゆる問題を解決できるようになります。私たちの世界が世界亀の上に成り立っていないのと同じように、これは間違いであることが今でははっきりとわかると思います。
問題 I – 根本的に壊れた孤立
前の層を管理する各 AI 層は、その前の層のデータを読み取る必要があります。そのデータは、予測不可能な形で管理層に影響を与える可能性があり、最も単純なケースでは、防御をうまく回避するために使用される可能性があります。
問題 II – DDoS AI インフラストラクチャ
問題 I に対する確実な解決策が存在することは知られていませんが、十分な数のレイヤーを追加すると、攻撃が成功する可能性を大幅に下げることができると主張できます。
ただし、基盤となる防御インフラストラクチャを飽和させる攻撃手法が存在し、すべてをブロックするか、監視されていないリクエストを通過させるかの選択を迫られます。
AI インフラストラクチャと DDoS 緩和は無料ではありません。制御税に関する調査では、制御手段を AI パイプラインに統合するための運用コストと財務コストが現実のものであり、安全性保証のレベルに応じて増加することが示されています。これにより、新たな攻撃対象領域が開かれる可能性がさらに高まります。
巧妙に作られた単一の入力により、防御側に多くのレイヤーを強制的に通過させることができ、防御側に不利なコストの非対称性が生じます。わかりやすい経済学がそれをさらに複雑にします。安価な AI の時代は終わり、多層の LLM ベースの防御フレームワークには、それ自体が悪用される可能性のある値札が付いています。
ソリューションはどのようなものであるべきか
最終的な権限は、決定論的でバイパスできないものの背後になければなりません。

ゲート。 AI は、破壊的で元に戻せないアクション (運用データベースの削除、資金の移動、運用環境へのプッシュ) に対する直接の権限を決して保持してはなりません。したがって、最後の防御線は常に人間による監視か、AI による回避策のない決定論的なスクリプトのいずれかでなければなりません。
これは、多層防御の意味を再構成するものでもあります。 LLM をさらに積み重ねても、障害は相関しているため、独立したレイヤーは追加されません。古典的な非 LLM 機械学習技術 (つまり、さまざまな分類子) は、LLM と同様に分離リーク問題の影響を受けにくいため、代わりにここで使用できます。したがって、決定論的ゲート、分類子、スクリプトおよび/または人間の直接監視によって、LLM のスタックだけでは解決できない相関関係が解消されます。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
ReARM: ガバナンス AI コーディング エージェントのデモ
最終防衛線はAIであってはなりません
ReARM: ガバナンス AI コーディング エージェントのデモ
新しい ReARM ピッチ – エージェント コーディングのガバナンス
パラダイムが変わるとき: AI エージェントのゼロトラスト モデル
OnPod で QA とサイバー セキュリティについて話しました
パラダイム シフトのとき: AI エージェントのゼロトラスト モデル - Git Push の恐怖を取り除くための仕事と生活のメモ
OnPod で QA とサイバー セキュリティについて話しました - 現在のテクノロジー時代を生き抜きたいという仕事と生活のメモ - 優れた QA になることを学びましょう
NTIA 準拠の SBOM の実践ガイド - 単一の SBOM では十分ではない理由に関する仕事と生活のメモ
開発マシンを信頼できないものとして扱い始めるべき時 - Linux および Chromebook の SSH 用 YubiKey に関する仕事と生活のメモ
開発マシンを信頼できないものとして扱い始める時期 - Windows 上の SSH 用 YubiKey に関する仕事と生活のメモ: 完全なチュートリアル

## Original Extract

The frequently circulating answer to the question of how we govern AI doing the work at scale is "AI turtles all the way down". Meaning that more AI

Skip to content
Work & Life Notes
Pavel Shukhman's blog: Tech, Travels, Relationships, Life
The last line of defense must not be AI
The frequently circulating answer to the question of how we govern AI doing the work at scale is “AI turtles all the way down”. Meaning that more AI downstream can solve any problems originating upstream. I believe we can now clearly see it’s a fallacy – the same way our world does not rest on a World Turtle.
Problem I – Fundamentally broken isolation
Each AI layer governing previous layer has to read data of that previous layer. That data may affect the governing layer in the unpredictable way, and in the most straightforward case may be used to successfully evade defenses .
Problem II – DDoS AI infrastructure
While no robust solution to Problem I is known to exist, one can argue that adding a sufficient number of layers can significantly reduce the likelihood of a successful attack.
However, there are attack techniques to saturate the underlying defensive infrastructure forcing a choice to either block everything or let unmonitored requests go through.
AI infrastructure and DDoS mitigation is not free: research on the control tax shows that the operational and financial cost of integrating control measures into an AI pipeline is real and scales with the level of safety assurance. This has further potential to open a new attack surface.
A single crafted input can force the defense through many layers creating cost asymmetry not in favor of the defender. The plain economics compound it: the era of cheap AI is over, and a multi-layered LLM-based defensive framework carries a price tag that may itself be exploited.
What should the solution look like
The final authority must sit behind a deterministic, non-bypassable gate. AI must never hold direct permissions for destructive, irreversible actions ( deleting a production database , moving funds, pushing to prod). So the last line of defense must always be either human oversight or a deterministic script with no AI workarounds.
This also reframes what defense-in-depth should mean. Stacking more LLMs doesn’t add independent layers as their failures are correlated. The classic non-LLM machine learning techniques (i.e. various classifiers) may be used here instead as they are not susceptible to isolation leak problems the same way as LLMs. So a deterministic gate, plus a classifier, plus a script and/or direct human oversight breaks the correlation that a stack of LLMs cannot resolve on its own.
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
ReARM: Governing AI Coding Agents Demo
The last line of defense must not be AI
ReARM: Governing AI Coding Agents Demo
New ReARM Pitch – Governance for Agentic Coding
When the Paradigm Shifts: A Zero-Trust Model for AI Agents
Talked about QA and Cyber Security at OnPod
When the Paradigm Shifts: A Zero-Trust Model for AI Agents - Work & Life Notes on Take The Fear Out Of Git Push
Talked about QA and Cyber Security at OnPod - Work & Life Notes on Want to Survive Current Tech Era – Learn to Be a Good QA
Practical Guide to NTIA Compliant SBOM - Work & Life Notes on Why a Single SBOM is Never Enough
Time to Start Treating Dev Machines as Untrusted - Work & Life Notes on YubiKey for SSH on Linux and Chromebook
Time to Start Treating Dev Machines as Untrusted - Work & Life Notes on YubiKey for SSH on Windows: Complete Walkthrough
