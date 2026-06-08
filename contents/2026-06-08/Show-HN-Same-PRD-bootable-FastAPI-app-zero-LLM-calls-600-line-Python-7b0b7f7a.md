---
source: "https://github.com/Anioko/spec-driven-development"
hn_url: "https://news.ycombinator.com/item?id=48448063"
title: "Show HN: Same PRD → bootable FastAPI app, zero LLM calls (600-line Python)"
article_title: "GitHub - Anioko/spec-driven-development: Spec-driven development guide: maturity levels, Spec Kit comparison, PRD-to-app demo · GitHub"
author: "anioko1"
captured_at: "2026-06-08T18:58:14Z"
capture_tool: "hn-digest"
hn_id: 48448063
score: 1
comments: 0
posted_at: "2026-06-08T17:11:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Same PRD → bootable FastAPI app, zero LLM calls (600-line Python)

- HN: [48448063](https://news.ycombinator.com/item?id=48448063)
- Source: [github.com](https://github.com/Anioko/spec-driven-development)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T17:11:23Z

## Translation

タイトル: HN を表示: 同じ PRD → 起動可能な FastAPI アプリ、LLM 呼び出しなし (600 行の Python)
記事タイトル: GitHub - Anioko/spec-driven-development: 仕様駆動開発ガイド: 成熟度レベル、仕様キットの比較、PRD からアプリへのデモ · GitHub
説明: 仕様駆動開発ガイド: 成熟度レベル、仕様キットの比較、PRD からアプリへのデモ - Anioko/spec-driven-development

記事本文:
GitHub - Anioko/spec-driven-development: 仕様駆動開発ガイド: 成熟度レベル、仕様キットの比較、PRD からアプリへのデモ · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アニオコ
/
スペック駆動開発
公共
通知
通知設定を変更するにはサインインする必要があります

GS
追加のナビゲーション オプション
コード
あにこ/スペック駆動開発
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github .github docs docs 例 例 HN.md HN.md ライセンス ライセンス README.md README.md _config.yml _config.yml デモ.sh デモ.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
2026 年の仕様駆動開発: エージェント プロンプトからコンパイル済みアプリケーションまで
最終更新日: 2026 年 6 月
SDD の成熟度レベル、GitHub Spec Kit が状況にどのように適合するか、エージェント ワークフローの代わりに仕様コンパイラーが必要な場合についての実用的なガイドです。
トピック: 仕様駆動開発、仕様駆動、prd、codegen、archimate、モデル駆動、オープンソース
スペック駆動開発とは何ですか?
仕様駆動開発 (SDD) は、実装に関するすべての決定が、チャット ログや「エージェントが最後に書いたもの」、そして誰かの頭の中にある固有の知識ではなく、正式な仕様または構造化された仕様にまで遡る方法論です。
システムの動作 (エンティティ、ワークフロー、API)
制約 (認証、テナンシー、コンプライアンス) の下でどのように動作する必要があるか
重要な決定が行われた理由 (ADR、アーキテクチャーノート)
コードは仕様の出力であり、真実の情報源ではありません。
GitHub は、2025 年 9 月に Spec Kit と AI を使用した SDD に関するブログ投稿で現在の波を広めました。 Amazon の Kiro 氏と Andrej Karpathy 氏の「ポストバイブコーディング」の論文と、それに続く比較記事の洪水が続きました。この用語が Google、Perplexity、GitHub の検索でランク付けされているのは、チームが壁にぶつかっているためです。プロンプト専用 AI はデモでは高速ですが、本番環境では脆弱です。
このガイドでは、ほとんどの記事が統合されている SDD の 2 層について説明し、4 つの成熟度レベルをマップし、生成ステップで LLM を使用せずに 1 つのコマンドで PRD ファイルから起動可能なアプリケーションに移行する方法を示します。
2 層の SDD (混同しないでください)
階層
私は何

しません
こんな方に最適
例
エージェントのワークフロー
スペック ファイルは AI エージェントにリポジトリの編集を指示します
既存のコードベースの機能
GitHub 仕様キット、Kiro、BMad
仕様コンパイラ
仕様は新しいアプリケーションにコンパイルされます
グリーンフィールド製品、再現可能な足場
アーキエット-マイクロコードジェン、アーキエット
GitHub の行:「仕様は実行可能になります。」
このリポジトリの行: リポジトリ内のエージェント編集ではなく、出荷可能なアプリにコンパイルされる実行可能仕様。
どちらもスペック重視です。これらは交換可能ではありません。
→ 完全な比較: Archiet vs Spec Kit
レベル
名前
LLM はコードを書きますか?
例
1
スペックファースト
はい
アドホック PRD + カーソル
2
スペック固定
はい
スペックキット、Kiro、BMad
3
ソースとしてのスペック
はい（リジェネ）
テッスル、オープンスペック
4
仕様からアプリケーションまで
いいえ
アーキエ-マイクロコードジェン
レベル 4 は仕様を正式なモデル (ゲノム) として扱います。コードは決定論的なモデルからテキストへの (M2T) 変換によって生成されます。この変換は、MDA、AUTOSAR、SCADE が数十年にわたって使用してきたものと同じもので、現在は Web アプリに適用されています。
git clone https://github.com/Anioko/spec-driven-development.git
CD スペック駆動開発
chmod +x デモ.sh && ./demo.sh
または PyPI のみを使用する場合:
pip インストール archive-microcodegen
Archiet-microcodegen 例/sample-prd.md --out ./myapp/
cd myapp && docker 構成
# → http://localhost:8000/docs
入力:examples/sample-prd.md — 30 行のタスクマネージャー PRD。
出力: FastAPI + SQLAlchemy + Alembic + JWT 認証 + テスト + ARCHITECTURE.md + docker-compose.yml 。
同じ PRD → 同じファイル。毎回。 APIキーがありません。
Archiet-microcodegen は、Archiet のコア アルゴリズムを実装します (メインのモノリポジトリの scripts/microcodegen.py に文書化されています)。
PRD テキスト
│
▼ parse_prd() 正規表現 + ヒューリスティック → マニフェスト (エンティティ、ストーリー、認証)
マニフェスト
│
▼ ArchiMate 3.2 要素型指定を使用したmanifest_to_genome() 形式的 IR
ゲノム
│
▼ render_genome() 文字列テンプレート → {パス

:コンテンツ}
ファイル
│
▼ Pack() ZIP または --out ディレクトリ
起動可能なアプリ
LLM はアップストリームで許可されます (乱雑な散文をマニフェストに変換します)。 LLM は低下ステップでは禁止されています。ゲノムを取得すると、発光は決定的になります。この境界により、出力が監査可能になります。
マルチスタック: 同じ仕様、異なるエコシステム
FastAPI パッケージは 1 つのエントリ ポイントです。同じパイプラインがエコシステムごとに出荷されます。
同じ PRD を 3 つのスタックで生成します。構造を比較します。 1 つ選んでください。ベンダーの IDE は必要ありません。
Spec Kit と Archiet-microcodegen (正直)
GitHub 仕様キット
アーキエ-マイクロコードジェン
あなたは得ます
エージェントはプロジェクトにタスクを実装します
新しいアプリディレクトリ/ZIP
決定論
エージェント依存
同じスペック→同じ出力
オフライン
エージェント + ネットワークが必要
エアギャップで実行
ブラウンフィールド
✅
❌ (グリーンフィールド再生)
コンプライアンスパック
DIY
基本アーキテクチャ.md; Archiet のフルパック
Cursor/Copilot を使用し、増分機能を出荷する場合は、Spec Kit を使用してください。
要件ファイルから再現可能なグリーンフィールド アプリが必要な場合は、archiet-microcodegen を使用します。
多くのチームは両方を使用します。スキャフォールドをコンパイルしてから、反復のために仕様キットをコンパイルします。
完全な Archiet プラットフォームが必要な場合
お客様がレビューしたゲノムを使用した LLM PRD 抽出
15 以上のバックエンド スタック + Next.js + Expo を 1 つの仕様から
デリバリーゲート — 出荷性監査、合成ブートテスト、コンサルティンググレードの ARCHITECTURE.md
コンプライアンス アーティファクト — 同じモデルから派生した SOC 2、GDPR、HIPAA、PCI、DORA、NIS2
トレーサビリティ マトリックス — 要件 → エンティティ → ルート → テスト → ADR
オープンソース CLI はアルゴリズムです。プラットフォームは最上位の効率層です。
これは GitHub Spec Kit と同じですか?
いいえ。Spec Kit は、リポジトリ内のエージェント ワークフローを構造化します ( /speckit.specify 、 /speckit.plan など)。 Archiet-microcodegen は、codegen パスに LLM を含めずに仕様を新しいアプリケーションにコンパイルします。完了

初歩的なものであり、競合するものではありません。
スペック主導の開発には AI が必要ですか?
いいえ。レベル 4 SDD は決定的です。 AI は乱雑な PRD の解析に役立ちます (レベル 1 ～ 2)。生成ステップは純粋な M2T にすることができます。
規制された環境で使用できますか?
レベル 4 は、再現性とエアギャップ使用を考慮して設計されています。クリティカル パスに LLM API がありません。完全なコンプライアンスの説明には Archiet プラットフォームが必要です。
Vibe コーディングはデモまでの速度を最適化します。 SDD は、移行、認証パターン、アーキテクチャ ドキュメント、トレーサビリティなど、実稼働までの速度を最適化します。 Karpathy の 2026 年 2 月の投稿は文化的な変化を示しました。スペックキットはそれにツールを与えました。
ゲノムは ArchiMate 3.2 要素型指定 (ApplicationComponent、DataObject、ApplicationService など) を使用します。生成されたすべての ARCHITECTURE.md は、コード アーティファクトを正式なモデルにマッピングします。
これは JHipster / Nest CLI スキャフォールドとどう違うのですか?
足場は空の構造を放出します。このパイプラインは、PRD コンテンツからエンティティ固有の CRUD、認証、移行、テスト、OpenAPI を生成し、仕様が変更されたときに再生成します。
パス
目的
例/sample-prd.md
デモ用の PRD をコピーアンドペースト
デモ.sh
ワンコマンド PRD → アプリの実行
docs/what-is-spec-driven-development.md
短い標準的な説明
docs/archiet-vs-spec-kit.md
比較ページ
docs/成熟度レベル.md
レベル 1 ～ 4 のフレームワーク
HN.md
HN タイトル + 最初のコメントの下書きを表示
リンク
コンプライアンス ガイド: github.com/Anioko/compliance-from-architecture — 3 層 (Vanta/Drata と codegen)、EU AI 法
フルプラットフォーム:archiet.com/spec-driven-development
アルゴリズムソース: github.com/aniekanasuquookono-web/archiet ( scripts/microcodegen.py )
Dev.to ウォークスルー: IDE を使用しない仕様駆動開発
コミュニティ比較：cameronsjo/spec-compare
仕様キット: github.com/github/spec-kit
アニーカン・アスクオ・オコノによって管理されています。貢献を歓迎します: 比較修正、

新しいエコシステム パッケージ、成熟度レベルの例。
仕様主導型開発ガイド: 成熟度レベル、仕様キットの比較、PRD からアプリへのデモ
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Spec-driven development guide: maturity levels, Spec Kit comparison, PRD-to-app demo - Anioko/spec-driven-development

GitHub - Anioko/spec-driven-development: Spec-driven development guide: maturity levels, Spec Kit comparison, PRD-to-app demo · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Anioko
/
spec-driven-development
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Anioko/spec-driven-development
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github .github docs docs examples examples HN.md HN.md LICENSE LICENSE README.md README.md _config.yml _config.yml demo.sh demo.sh View all files Repository files navigation
Spec-Driven Development in 2026: From Agent Prompts to Compiled Applications
Last updated: June 2026
A practical guide to SDD maturity levels, how GitHub Spec Kit fits the landscape, and when you need a spec compiler instead of an agent workflow .
Topics: spec-driven-development · spec-driven · prd · codegen · archimate · model-driven · opensource
What is spec-driven development?
Spec-driven development (SDD) is a methodology where every implementation decision traces back to a formal or structured specification — not a chat log, not "whatever the agent wrote last," and not tribal knowledge in someone's head.
What the system does (entities, workflows, APIs)
How it must behave under constraints (auth, tenancy, compliance)
Why key decisions were made (ADRs, architecture notes)
Code is an output of the spec , not the source of truth.
GitHub popularised the current wave in September 2025 with Spec Kit and their blog post on SDD with AI . Amazon's Kiro, Andrej Karpathy's "post-vibe-coding" thesis, and a flood of comparison articles followed. The term is ranking on Google, Perplexity, and GitHub search because teams hit a wall: prompt-only AI is fast for demos and fragile for production.
This guide explains two tiers of SDD most articles conflate, maps four maturity levels , and shows how to go from a PRD file to a bootable application in one command — no LLM in the generation step.
Two tiers of SDD (don't conflate them)
Tier
What it does
Best for
Example
Agent workflow
Spec files guide an AI agent to edit your repo
Features in existing codebases
GitHub Spec Kit, Kiro, BMad
Spec compiler
Spec is compiled into a new application
Greenfield products, reproducible scaffolds
archiet-microcodegen, Archiet
GitHub's line: "Specifications become executable."
This repo's line: Executable specifications that compile to shippable apps — not agent edits in your repo.
Both are spec-driven. They are not interchangeable.
→ Full comparison: Archiet vs Spec Kit
Level
Name
LLM writes code?
Examples
1
Spec-First
Yes
Ad-hoc PRDs + Cursor
2
Spec-Anchored
Yes
Spec Kit , Kiro, BMad
3
Spec-as-Source
Yes (regen)
Tessl, OpenSpec
4
Spec-to-Application
No
archiet-microcodegen
Level 4 treats the spec as a formal model (genome). Code is emitted by deterministic model-to-text (M2T) transformation — the same discipline MDA, AUTOSAR, and SCADE have used for decades, now applied to web apps.
git clone https://github.com/Anioko/spec-driven-development.git
cd spec-driven-development
chmod +x demo.sh && ./demo.sh
Or with PyPI only:
pip install archiet-microcodegen
archiet-microcodegen examples/sample-prd.md --out ./myapp/
cd myapp && docker compose up
# → http://localhost:8000/docs
Input: examples/sample-prd.md — a 30-line task-manager PRD.
Output: FastAPI + SQLAlchemy + Alembic + JWT auth + tests + ARCHITECTURE.md + docker-compose.yml .
Same PRD → same files. Every time. No API key.
archiet-microcodegen implements Archiet's core algorithm (documented in scripts/microcodegen.py in the main monorepo):
PRD text
│
▼ parse_prd() regex + heuristics → manifest (entities, stories, auth)
manifest
│
▼ manifest_to_genome() formal IR with ArchiMate 3.2 element typing
genome
│
▼ render_genome() string templates → {path: content}
files
│
▼ pack() ZIP or --out directory
bootable app
LLMs are allowed upstream (turning messy prose into a manifest). LLMs are forbidden in the lowering step — once you have a genome, emission is deterministic. That boundary is what makes output auditable.
Multi-stack: same spec, different ecosystems
The FastAPI package is one entry point. The same pipeline ships per ecosystem:
Generate the same PRD in three stacks. Compare structure. Pick one. No vendor IDE required.
Spec Kit vs archiet-microcodegen (honest)
GitHub Spec Kit
archiet-microcodegen
You get
Agent implements tasks in your project
New app directory / ZIP
Determinism
Agent-dependent
Same spec → same output
Offline
Needs agent + network
Runs air-gapped
Brownfield
✅
❌ (greenfield regen)
Compliance pack
DIY
Basic ARCHITECTURE.md; full pack on Archiet
Use Spec Kit when you live in Cursor/Copilot and ship incremental features.
Use archiet-microcodegen when you need a reproducible greenfield app from a requirements file.
Many teams use both: compile the scaffold, then Spec Kit for iteration.
When you need the full Archiet platform
LLM PRD extraction with customer-reviewed genome
15+ backend stacks + Next.js + Expo from one spec
Delivery gates — shippability audit, synthetic boot test, consulting-grade ARCHITECTURE.md
Compliance artifacts — SOC 2, GDPR, HIPAA, PCI, DORA, NIS2 derived from the same model
Traceability matrix — requirement → entity → route → test → ADR
The open-source CLIs are the algorithm . The platform is the efficiency layer on top.
Is this the same as GitHub Spec Kit?
No. Spec Kit structures agent workflows in your repo ( /speckit.specify , /speckit.plan , …). archiet-microcodegen compiles a spec into a new application without an LLM in the codegen path. Complementary, not competing.
Does spec-driven development require AI?
No. Level 4 SDD is deterministic . AI helps parse messy PRDs (Level 1–2). The generation step can be pure M2T.
Can I use this in regulated environments?
Level 4 is designed for reproducibility and air-gapped use. No LLM API in the critical path. Full compliance narratives require the Archiet platform .
Vibe coding optimises for speed-to-demo. SDD optimises for speed-to-production — migrations, auth patterns, architecture docs, traceability. Karpathy's Feb 2026 post marked the cultural shift; Spec Kit gave it tooling.
The genome uses ArchiMate 3.2 element typing (ApplicationComponent, DataObject, ApplicationService, …). Every generated ARCHITECTURE.md maps code artifacts back to the formal model.
How is this different from JHipster / Nest CLI scaffolds?
Scaffolds emit empty structure. This pipeline emits entity-specific CRUD, auth, migrations, tests, and OpenAPI from your PRD content — and regenerates when the spec changes.
Path
Purpose
examples/sample-prd.md
Copy-paste PRD for demos
demo.sh
One-command PRD → running app
docs/what-is-spec-driven-development.md
Short canonical explainer
docs/archiet-vs-spec-kit.md
Comparison page
docs/maturity-levels.md
Levels 1–4 framework
HN.md
Show HN title + first-comment draft
Links
Compliance guide: github.com/Anioko/compliance-from-architecture — three layers (Vanta/Drata vs codegen), EU AI Act
Full platform: archiet.com/spec-driven-development
Algorithm source: github.com/aniekanasuquookono-web/archiet ( scripts/microcodegen.py )
Dev.to walkthrough: Spec-Driven Development Without an IDE
Community comparison: cameronsjo/spec-compare
Spec Kit: github.com/github/spec-kit
Maintained by Aniekan Asuquo Okono . Contributions welcome: comparison corrections, new ecosystem packages, maturity-level examples.
Spec-driven development guide: maturity levels, Spec Kit comparison, PRD-to-app demo
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
