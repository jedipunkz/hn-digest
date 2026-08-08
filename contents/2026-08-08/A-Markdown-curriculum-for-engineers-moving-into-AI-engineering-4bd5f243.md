---
source: "https://github.com/bimal1023/AI-native-engineer"
hn_url: "https://news.ycombinator.com/item?id=49219402"
title: "A Markdown curriculum for engineers moving into AI engineering"
article_title: "GitHub - bimal1023/AI-native-engineer: Tracking my journey into AI engineering, one module at a time. Fundamentals to agents, with hands-on projects. Open to all. · GitHub"
author: "bkumal"
captured_at: "2026-08-08T07:41:02Z"
capture_tool: "hn-digest"
hn_id: 49219402
score: 1
comments: 1
posted_at: "2026-08-08T06:40:44Z"
tags:
  - hacker-news
  - translated
---

# A Markdown curriculum for engineers moving into AI engineering

- HN: [49219402](https://news.ycombinator.com/item?id=49219402)
- Source: [github.com](https://github.com/bimal1023/AI-native-engineer)
- Score: 1
- Comments: 1
- Posted: 2026-08-08T06:40:44Z

## Translation

タイトル: AI エンジニアリングに移行するエンジニアのための Markdown カリキュラム
記事のタイトル: GitHub - bimal1023/AI-native-engineer: AI エンジニアリングへの私の道のりを一度に 1 モジュールずつ追跡します。エージェント向けの基礎と実践的なプロジェクト。すべての人に開かれています。 · GitHub
説明: AI エンジニアリングへの私の道のりを、一度に 1 モジュールずつ追跡します。エージェント向けの基礎と実践的なプロジェクト。すべての人に開かれています。 - bimal1023/AI ネイティブ エンジニア

記事本文:
GitHub - bimal1023/AI-native-engineer: AI エンジニアリングへの私の道のりを一度に 1 モジュールずつ追跡しています。エージェント向けの基礎と実践的なプロジェクト。すべての人に開かれています。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ビマル1023
/
AIネイティブエンジニア
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
16 コミット 16 コミット 01-llm-fundamentals 01-llm-fundamentals 02-prompting-and-context-engineering 02-prompting-and-contex

tエンジニアリング 03-取得とラグ 03-取得とラグ 04-エージェントとツールの使用 04-エージェントとツールの使用 05-評価と観測可能性 05-評価と観測可能性 06-展開とAIインフラ 06-展開とAIインフラ07-emerging-topics 07-emerging-topics アセット アセット .gitignore .gitignore ライセンス ライセンス README.md README.md capstone.md capstone.md hot-topics.md hot-topics.md Soft-skills.md Soft-skills.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI/ML エンジニアリングに移行するソフトウェア エンジニアのための、構造化された独自のカリキュラム。LLM API を呼び出すことができる人と、本番環境に耐える AI システムを設計、評価、出荷、運用できる人を分ける分野です。これは単純なマークダウンとして構築されています。基礎からフロンティアまで 7 つのモジュールがあり、それぞれに正規のリソース、業界で実際に使用されているツール、実践的なプロジェクト、そして最初に人々を苦しめる落とし穴が含まれています。私は学びながら（間違いも含めて公開していますが）書いており、同じ移行をしようとしている人、つまり独学の開発者、自分の範囲に AI を追加するバックエンド/フルスタック エンジニア、そして「OpenAI SDK を一度使った」以上の深みを求める新卒者など、誰にでも公開されています。
最終レビュー日: 2026 年 8 月 · 何が急速に進んでいるのか、何が解決されているのかについては、hot-topics.md を参照してください。
#
モジュール
何ができるようになるのか
01
LLM の基礎
モデルが実際に何をしているのかを説明し、証拠に基づいてモデルを選択してください
02
プロンプトとコンテキスト エンジニアリング
信頼性の高い構造化された出力を取得し、予算としてコンテキスト ウィンドウを管理します
03
取得とRAG
適切なものを測定可能な形で見つける検索を構築する
04
エージェントとツールの使用
終了、回復、予算内にとどまるエージェント ループを設計する
05
評価と可観測性
オフラインでも実稼働でも、雰囲気を数字に置き換えます
06
導入と AI インフラストラクチャ
既知のコスト、レイテンシー、障害でモデルを提供する

プロフィール
07
新しいトピック
デモを追わずに新機能を評価
横断的なファイル
hot-topics.md — 過去 6 ～ 12 か月での変化と安定したファンダメンタルズ
Soft-skills.md — AI 生成コードのレビュー、コストとレイテンシーのトレードオフ、セキュリティ意識、制限事項の伝達
capstone.md — すべてのモジュールを結合した最終プロジェクト
モデルを呼び出しているのは、中央の小さなボックスです。このリポジトリは、その周りのすべてについて説明しています。
リポジトリをフォークし、作業に応じてこれらにチェックを入れます。アイデアを他のエンジニアに説明でき、それを使用するコードを出荷できた時点で、ボックスは「完成」したことになります。
トランスフォーマーと注意メカニズム
トークン化とコンテキスト ウィンドウ
トレーニングのライフサイクル: 事前トレーニング → SFT → 好みの調整
推論とデコード: サンプリング パラメータ、KV キャッシュ、ストリーミング
モデルの選択、スケーリング則、および推論モデル
プロジェクト: モデル ベイクオフ ハーネス
02 · プロンプトとコンテキスト エンジニアリング
迅速な解剖とその後の指示
少数のショットの例と推論の引き出し
構造化された出力とスキーマの適用
コンテキスト エンジニアリング: アセンブリ、圧縮、キャッシュ
迅速なバージョン管理、テスト、最適化
プロジェクト: 構造化抽出パイプライン
チャンク化、解析、インデックスの設計
クエリの理解: 書き換え、ルーティング、フィルタリング
RAG の評価と故障モード
プロジェクト: 実際のコーパス上で引用された回答
ツール呼び出しとエージェントループ
エージェントのアーキテクチャ: ワークフローとエージェント
環境とプロトコル (MCP、サンドボックス、コンピューターの使用)
信頼性: 予算、再試行、人間参加型
プロジェクト: 永続的な状態を持つ制限付きリサーチ エージェント
05・評価と可観測性
評価の基礎とエラー分析
LLM-as-judge およびジャッジキャリブレーション
トレース、メトリクス、および OTel GenAI 規則
実稼働フィードバック ループと CI 回帰ゲート
エージェントと RAG 固有

フィクションの評価
プロジェクト: CI に接続された評価ハーネス
サービス提供と推論の最適化
API 層アーキテクチャ: ゲートウェイ、ストリーミング、フォールバック
微調整 vs. RAG vs. プロンプト
プロジェクト: ゲートウェイの背後にあるセルフホスト型モデル、負荷テスト済み
推論モデルとテスト時の計算
プロトコルと相互運用性
マルチモーダル、リアルタイム、コンピューターの使用
小規模モデル、蒸留、オンデバイス
プロジェクト: 1 つの新興分野に関する機能メモ
ソフトスキル — 全 5 つのセクション
注目のトピック — レビューおよび更新
Capstone — 出荷され、書き上げられます
順番に進めますが、金メッキはしないでください。モジュール 01 ～ 03 は、他のすべての前提条件です。 04～06は一度入手すればどの順番でも取得可能です。
ボックスにチェックを入れる前にプロジェクトをビルドします。すべてのモジュールは 1 つの成果物で終了します。再ランキングについて読んでも何もわかりません。 remember@10 が 0.61 から 0.88 にジャンプするのを見ると、モジュール全体がわかります。
何が壊れたかを書きます。バグ、間違った仮定、数値を記録したメモ/フォルダーを保管してください。そのログは実際のポートフォリオの成果物です。
読書をタイムボックス化します。各モジュールには、サブトピックごとに 2 ～ 3 つの正規リソースがリストされます。一次ソースを一度読み、残りをざっと読んでからビルドに進みます。
hot-topics.md を四半期ごとに読み返します。この材料の約 3 分の 1 には保存期限があります。このファイルには、どの 3 番目であるかがわかります。
推奨ペース: パートタイムでは 1 ～ 2 週間ごとに 1 モジュール、その後はキャップストーンで 3 ～ 4 週間。すでに職場で AI 機能をリリースしている場合は、より高速でも問題ありません。
これは公開学習リソースであり、注目が集まるほど改善されます。役立つ貢献:
壊れたリンクまたは古いリンク – 最も迅速で価値のある修正。
正規のリソースの改善 — 論文や投稿でサブトピックがより明確に説明されている場合は、それを置き換えてその理由を述べます。
工具の修正 — 工具リストが古くなります。実際の p に基づいてツールを追加、削除、または再ランク付けする PR

特にローテーションでのご利用は大歓迎です。
新たな落とし穴 — 1 日を犠牲にするものがある場合、それは「よくある落とし穴」セクションに属します。
プロジェクトの記事をリンクします。具体的な例は散文に勝ります。
構造的なもの (新しいモジュール、並べ替え) については、作成する前に議論するためにイシューを開いてください。小さな修正については、PR を開いてください。自社のスタイルを維持します。タイトな散文、実際のツール名、アグリゲーターを介した一次ソースへのリンク、マーケティング言語の禁止。
求めていないもの: SEO フィラー、ゲートされたコンテンツへのリンク、「トップ 10 AI ツール」のリスト、ベンダーの宣伝。
MIT — それをフォークし、そこから教え、同意できない部分を取り除きます。帰属表示は歓迎されますが、ライセンスでは著作権表示を保持することのみが求められます。
AI エンジニアリングへの私の道のりを、一度に 1 モジュールずつ追跡します。エージェント向けの基礎と実践的なプロジェクト。すべての人に開かれています。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Tracking my journey into AI engineering, one module at a time. Fundamentals to agents, with hands-on projects. Open to all. - bimal1023/AI-native-engineer

GitHub - bimal1023/AI-native-engineer: Tracking my journey into AI engineering, one module at a time. Fundamentals to agents, with hands-on projects. Open to all. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
bimal1023
/
AI-native-engineer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits 01-llm-fundamentals 01-llm-fundamentals 02-prompting-and-context-engineering 02-prompting-and-context-engineering 03-retrieval-and-rag 03-retrieval-and-rag 04-agents-and-tool-use 04-agents-and-tool-use 05-evaluation-and-observability 05-evaluation-and-observability 06-deployment-and-ai-infra 06-deployment-and-ai-infra 07-emerging-topics 07-emerging-topics assets assets .gitignore .gitignore LICENSE LICENSE README.md README.md capstone.md capstone.md hot-topics.md hot-topics.md soft-skills.md soft-skills.md View all files Repository files navigation
A structured, opinionated curriculum for software engineers moving into AI/ML engineering — the discipline that separates someone who can call an LLM API from someone who can design, evaluate, ship, and operate AI systems that hold up in production. It's built as plain markdown: seven modules from foundations to frontier, each with canonical resources, the tools actually used in industry, a hands-on project, and the pitfalls that bite people first. I'm writing it as I learn (publicly, mistakes included), and it's open for anyone making the same transition — self-taught devs, backend/full-stack engineers adding AI to their scope, and new grads who want depth beyond "I used the OpenAI SDK once."
Last reviewed: August 2026 · See hot-topics.md for what's moving fast vs. what's settled.
#
Module
What you'll be able to do
01
LLM Fundamentals
Explain what the model is actually doing, and pick one on evidence
02
Prompting & Context Engineering
Get reliable, structured output and manage the context window as a budget
03
Retrieval & RAG
Build retrieval that measurably finds the right thing
04
Agents & Tool Use
Design agent loops that terminate, recover, and stay in budget
05
Evaluation & Observability
Replace vibes with numbers, offline and in production
06
Deployment & AI Infra
Serve models at a known cost, latency, and failure profile
07
Emerging Topics
Evaluate new capabilities without chasing demos
Cross-cutting files
hot-topics.md — what changed in the last 6–12 months vs. stable fundamentals
soft-skills.md — reviewing AI-generated code, cost/latency tradeoffs, security awareness, communicating limitations
capstone.md — the final project that combines every module
Calling the model is the small box in the middle. This repo is about everything around it.
Fork the repo and tick these off as you go. A box is "done" when you can explain the idea to another engineer and you've shipped code that uses it.
Transformers and the attention mechanism
Tokenization and the context window
The training lifecycle: pretraining → SFT → preference tuning
Inference and decoding: sampling params, KV cache, streaming
Model selection, scaling laws, and reasoning models
Project: model bake-off harness
02 · Prompting & Context Engineering
Prompt anatomy and instruction following
Few-shot examples and reasoning elicitation
Structured output and schema enforcement
Context engineering: assembly, compaction, and caching
Prompt versioning, testing, and optimization
Project: structured extraction pipeline
Chunking, parsing, and index design
Query understanding: rewriting, routing, filtering
RAG evaluation and failure modes
Project: cited answers over a real corpus
Tool calling and the agent loop
Agent architectures: workflows vs. agents
Environments and protocols (MCP, sandboxes, computer use)
Reliability: budgets, retries, human-in-the-loop
Project: bounded research agent with durable state
05 · Evaluation & Observability
Eval fundamentals and error analysis
LLM-as-judge and judge calibration
Tracing, metrics, and OTel GenAI conventions
Production feedback loops and CI regression gates
Agent and RAG-specific evaluation
Project: eval harness wired into CI
Serving and inference optimization
API-layer architecture: gateways, streaming, fallbacks
Fine-tuning vs. RAG vs. prompting
Project: self-hosted model behind a gateway, load-tested
Reasoning models and test-time compute
Protocols and interoperability
Multimodal, realtime, and computer use
Small models, distillation, and on-device
Project: capability memo on one emerging area
Soft skills — all five sections
Hot topics — reviewed and re-dated
Capstone — shipped and written up
Go in order, but don't gold-plate. Modules 01–03 are prerequisites for everything else. 04–06 can be taken in any order once you have them.
Build the project before ticking the boxes. Every module ends with one deliverable. Reading about reranking teaches you nothing; watching recall@10 jump from 0.61 to 0.88 teaches you the whole module.
Write down what broke. Keep a notes/ folder with the bugs, the wrong assumptions, and the numbers. That log is the actual portfolio artifact.
Timebox the reading. Each module lists 2–3 canonical resources per subtopic. Read the primary source once, skim the rest, then go build.
Re-read hot-topics.md quarterly. Roughly a third of this material has a shelf life. That file tells you which third.
Suggested pace: one module every 1–2 weeks part-time, then 3–4 weeks for the capstone . Faster is fine if you're already shipping AI features at work.
This is a public learning resource and it gets better with more eyes on it. Useful contributions:
Broken or outdated links — the fastest, most valuable fix.
Better canonical resources — if a paper or post explains a subtopic more clearly, swap it in and say why.
Tool corrections — tooling lists go stale. PRs that add, remove, or re-rank tools based on real production use are especially welcome.
New pitfalls — if something cost you a day, it belongs in a Common Pitfalls section.
Your project write-ups — link them; concrete examples beat prose.
Open an issue to discuss anything structural (new modules, reordering) before writing it. For small fixes, just open a PR. Keep the house style: tight prose, real tool names, links to primary sources over aggregators, no marketing language.
Not looking for: SEO filler, links to gated content, "top 10 AI tools" listicles, or vendor pitches.
MIT — fork it, teach from it, rip out the parts you disagree with. Attribution appreciated but the license only asks you to keep the copyright notice.
Tracking my journey into AI engineering, one module at a time. Fundamentals to agents, with hands-on projects. Open to all.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
