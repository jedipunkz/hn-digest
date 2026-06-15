---
source: "https://mlsysbook.ai/"
hn_url: "https://news.ycombinator.com/item?id=48545049"
title: "Machine Learning Systems"
article_title: "Machine Learning Systems"
author: "ibobev"
captured_at: "2026-06-15T19:25:32Z"
capture_tool: "hn-digest"
hn_id: 48545049
score: 3
comments: 0
posted_at: "2026-06-15T18:14:15Z"
tags:
  - hacker-news
  - translated
---

# Machine Learning Systems

- HN: [48545049](https://news.ycombinator.com/item?id=48545049)
- Source: [mlsysbook.ai](https://mlsysbook.ai/)
- Score: 3
- Comments: 0
- Posted: 2026-06-15T18:14:15Z

## Translation

タイトル: 機械学習システム

記事本文:
/`
そして、内部に存在する navbar.href 値を相対化します。
サイト URL のドメイン。の共有`href: https://mlsysbook.ai/`
したがって、navbar-common.yml は `./index.html` として出力されます。
vol1/vol2/tinytorch/など。ではなく *そのサブサイト* ルートに解決されます
エコシステムのランディング ページ。その後、JS でブランド リンクの href をオーバーライドします。
Quarto のレンダリング: これは 1 行のユーザー エクスペリエンス修正です。
クアルトの URL 機構と戦います。
-->
機械学習システム
ホーム
ランディングページ
📖 ML Systems — インテリジェント システムのエンジニアリングに関するオープンアクセスの教科書。 Vol I: 基礎 → · Vol II: 大規模 →
🛠️ 本と並行して: TinyTorch (ビルド) · ハードウェア キット (デプロイ) · MLSys·im (モデル) · Labs (探索) · StaffML (実践)
📬 ニュースレター: ML Systems の洞察と最新情報 — 購読 →
AI エンジニアリングの物理学。
単一マシンからフリート規模のインフラストラクチャに至るまで、ML システムがどのように構築、最適化、デプロイされるかについての厳密で原則第一の扱い。
ハーバード大学 · MIT 出版物 2026
積極的にメンテナンス
·
最終更新日: 2026 年 4 月
·
リリースノート
機械学習システムの概要
ボリューム I のダウンロード:
HTML
PDF
EPUB
大規模な機械学習システム
ボリューム II のダウンロード:
HTML
PDF
EPUB
AIエンジニアリングのための充実したカリキュラム。
パスを選択します。書籍を読む、ラボでトレードオフを検討する、TinyTorch で内部を構築する、MLSys·im で制約をモデル化する、実際のハードウェアにデプロイする、StaffML で練習する、またはブループリントでフルコースを採用するなどです。
インタラクティブなマリモノート。パラメータを変更し、何が壊れるかを確認し、直感を構築します。
20 のプログレッシブ モジュールにわたって独自の ML フレームワークをゼロから構築します。魔力ゼロ。
第一原理パフォーマンスモデリング。 1 つのコマンドであらゆるボトルネックを実現。
メモリバウンド
コンピューティングバウンド
b=1
b=32
b=128
あ

算術強度
フロップ/秒
導入
ML を Arduino、Seeed、Grove、Raspberry Pi にデプロイします。実際のメモリ制限、実際の電力バジェット。
ML システムの役割に対する物理学に基づいた面接の質問。跳馬、訓練、模擬面接。
AI エンジニアリング ブループリント: 2 学期のシラバス、教育ガイド、ルーブリック、および TA ハンドブック。
ブループリント — コースのアーキテクチャ
ML Systems · 2 学期制カリキュラム
学期 1: 基礎
16 週間 · Vol I · 8 つの課題
学期 2: 大規模に
16 週間 · Vol II · キャップストーン
評価
ルーブリック · ピアレビュー · 採点
教職員
教育学・TAハンドブック
準備完了
教える
スピーカー ノートと 266 のオリジナル SVG 図を含む 35 の Beamer デッキ。立ち寄って教えてください。
カリキュラム、新しい章、コミュニティが構築しているものに関する最新情報。
12,000 人以上の登録者に参加しましょう
ミッションをサポートする
↓
私たちの使命
AI教育はこうあるべきだ
無料で誰でも利用できます。
誰もが AI を新しい電力と呼んでいますが、送電網を構築できるエンジニアがいないと電力は役に立ちません。 AI が効率的で信頼性が高く、安全であるためには、AI の構築方法を理解するエンジニアが世界に必要です。
その知識は、学びたい人なら誰でもアクセスできるものでなければなりません。このカリキュラムは、それを実現するための私たちの取り組みです。
ライブ読者数 — 180 か国以上
23,000 個以上のスター · 243,000 人以上の読者 · 180 か国以上
私たちの目標: 2030 年までに AI エンジニアを 100 万人にする
次のマイルストーン: 100,000 ★ — 現在 23,000 以上です。
すべてのスターは、他の人がこのリソースを発見するのに役立ちます。
GitHub でスターを付ける
スポンサーとパートナーのご紹介 →
© 2024-2026 ハーバード大学。 CC-BY-NC-SA 4.0 に基づいてライセンス供与されています
ボリューム I、ボリューム II、概要、コミュニティ、ニュースレター

## Original Extract

/`
and then relativizes any navbar.href value that lives inside the
site-url's domain. The shared `href: https://mlsysbook.ai/` in
navbar-common.yml therefore gets emitted as `./index.html`, which on
vol1/vol2/tinytorch/etc. resolves to *that subsite's* root rather than
the ecosystem landing page. Override the brand link href in JS after
Quarto's render: this is a one-line user-experience fix that does not
fight Quarto's URL machinery.
-->
Machine Learning Systems
Home
Landing Page
📖 ML Systems — an open-access textbook on the engineering of intelligent systems. Vol I: Foundations → · Vol II: At Scale →
🛠️ Alongside the book: TinyTorch (build) · Hardware Kits (deploy) · MLSys·im (model) · Labs (explore) · StaffML (practice)
📬 Newsletter: ML Systems insights & updates — Subscribe →
The physics of AI engineering.
A rigorous, principles-first treatment of how ML systems are built, optimized, and deployed — from a single machine to fleet-scale infrastructure.
Harvard University · MIT Press 2026
Actively maintained
·
Last updated April 2026
·
Release notes
Introduction to Machine Learning Systems
Volume I downloads:
HTML
PDF
EPUB
Machine Learning Systems at Scale
Volume II downloads:
HTML
PDF
EPUB
A complete curriculum for AI engineering.
Choose a path: read the books, explore trade-offs in labs, build the internals with TinyTorch, model constraints with MLSys·im, deploy on real hardware, practice with StaffML, or adopt the full course with the Blueprint.
Interactive Marimo notebooks. Change a parameter, see what breaks, build intuition.
Build your own ML framework from scratch across 20 progressive modules. Zero magic.
First-principles performance modeling. One command, every bottleneck.
mem-bound
compute-bound
b=1
b=32
b=128
Arithmetic Intensity
FLOP/s
DEPLOY
Deploy ML to Arduino, Seeed, Grove, and Raspberry Pi. Real memory limits, real power budgets.
Physics-grounded interview questions for ML systems roles. Vault, drills, and mock interviews.
The AI Engineering Blueprint: two-semester syllabi, pedagogy guide, rubrics, and TA handbook.
The Blueprint — Course Architecture
ML Systems · Two-Semester Curriculum
Semester 1: Foundations
16 wks · Vol I · 8 assignments
Semester 2: At Scale
16 wks · Vol II · capstone
Assessment
Rubrics · Peer review · Grading
Teaching Staff
Pedagogy · TA handbook
READY
TEACH
35 Beamer decks with speaker notes and 266 original SVG diagrams. Drop in and teach.
Updates on the curriculum, new chapters, and what the community is building.
Join 12,000+ subscribers
Support the Mission
↓
OUR MISSION
AI education should be
free and open to everyone.
Everyone calls AI the new electricity — but electricity is useless without engineers who can build the grid. For AI to be efficient, reliable, and safe, the world needs engineers who understand how to build it.
That knowledge should be accessible to anyone willing to learn. This curriculum is our commitment to making it so.
Live readership — 180+ countries
23,000+ stars · 243,000+ readers · 180+ countries
Our goal: 1,000,000 AI engineers by 2030
Next milestone: 100,000 ★ — we're at 23,000+ .
Every star helps others discover this resource.
Star on GitHub
Meet our sponsors & partners →
© 2024-2026 Harvard University. Licensed under CC-BY-NC-SA 4.0
Volume I · Volume II · About · Community · Newsletter
