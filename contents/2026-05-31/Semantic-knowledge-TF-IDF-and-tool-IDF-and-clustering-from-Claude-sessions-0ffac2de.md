---
source: "https://github.com/chigichan24/crune"
hn_url: "https://news.ycombinator.com/item?id=48336809"
title: "Semantic knowledge (TF-IDF and tool-IDF and clustering) from Claude sessions"
article_title: "GitHub - chigichan24/crune: Decipher the traces etched in past sessions and resurrect them as reusable skills. crune is a static web dashboard that analyzes Claude Code session logs, providing session playback, analytics, a semantic knowledge graph, and skill synthesis. · GitHub"
author: "ankitg12"
captured_at: "2026-05-31T00:28:48Z"
capture_tool: "hn-digest"
hn_id: 48336809
score: 2
comments: 0
posted_at: "2026-05-30T14:44:43Z"
tags:
  - hacker-news
  - translated
---

# Semantic knowledge (TF-IDF and tool-IDF and clustering) from Claude sessions

- HN: [48336809](https://news.ycombinator.com/item?id=48336809)
- Source: [github.com](https://github.com/chigichan24/crune)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T14:44:43Z

## Translation

タイトル: Claude セッションからのセマンティック知識 (TF-IDF とツール IDF とクラスタリング)
記事タイトル: GitHub - chigichan24/crune: 過去のセッションで刻まれた痕跡を解読し、再利用可能なスキルとして復活させる。 crune は、Claude Code セッション ログを分析する静的 Web ダッシュボードで、セッションの再生、分析、セマンティック ナレッジ グラフ、およびスキル合成を提供します。 · GitHub
説明: 過去のセッションで刻まれた痕跡を解読し、再利用可能なスキルとして復活させます。 crune は、Claude Code セッション ログを分析する静的 Web ダッシュボードで、セッションの再生、分析、セマンティック ナレッジ グラフ、およびスキル合成を提供します。 - chigichan24/クルーン

記事本文:
GitHub - chigichan24/crune: 過去のセッションで刻まれた痕跡を解読し、再利用可能なスキルとして復活させます。 crune は、Claude Code セッション ログを分析する静的 Web ダッシュボードで、セッションの再生、分析、セマンティック ナレッジ グラフ、およびスキル合成を提供します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のアカウントに切り替えました

タブまたはウィンドウ。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ちぎちゃん24
/
クルーン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
193 コミット 193 コミット .github .github bin bin docs docs public public scripts scripts src src .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md THIRD_PARTY_NOTICES THIRD_PARTY_NOTICES eslint.config.js eslint.config.js Index.html Index.html package-lock.json package-lock.json package.json package.json tsconfig.app.json tsconfig.app.json tsconfig.cli.json tsconfig.cli.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード・コード・ルーン — 痕跡は残り、スキルは現れる
過去のセッションで刻まれた痕跡を解読し、再利用可能なスキルとして復活させます。 crune は、Claude Code セッション ログを分析する静的 Web ダッシュボードおよび CLI ツールで、セッションの再生、分析、セマンティック ナレッジ グラフ、およびスキル合成を提供します。
セッション再生 --- ミニマップ ナビゲーション、ツール呼び出し検査、サブエージェント ブランチ拡張を備えたターンバイターンの会話再生
概要ダッシュボード --- アクティビティのヒートマップ、プロジェクトの分布、ツールの使用傾向、期間の分布、モデルの使用状況、および上位に編集されたファイル
セマンティック ナレッジ グラフ --- TF-IDF + Tool-IDF + 構造特徴、Truncated SVD、凝集クラスタリング、Louvain コミュニティ検出、Brandes 中心性 (アルゴリズムの詳細)
暗黙知 --- 抽出されたワークフロー パターン、ツール シーケンス、問題点 (長時間のセッション、ホット ファイル)
セッションの要約 --- 自動セッションの要約と分類 (LLM は不要)、/insigh で強化

ts ファセット データ（利用可能な場合）
スキル合成 --- ナレッジ グラフから再利用可能なスキル定義を合成します (アルゴリズムの詳細)
セマンティックナレッジグラフ（ルーン）
再利用可能なクロード コード スキル定義をセッション ログから直接生成します。クローンは必要ありません。
npx @chigichan24/crune --dry-run # スキル候補のプレビュー
npx @chigichan24/crune --skip-Synthetic # ヒューリスティック スキルを生成します (LLM なし)
npx @chigichan24/crune --count 3 --model haiku # LLM 合成スキル (claude CLI が必要)
npx @chigichan24/crune --output-dir ~ /.claude/skills # スキルを直接インストールする
出力はクロード コードのスキル形式 ( <name>/SKILL.md ) に従い、 /skill-name コマンドとしてすぐに使用できます。
npmインストール
# クロードコードのセッションログを分析する (~/.claude/projects/)
npm run 分析セッション
# 開発サーバーを起動します
npm 実行開発
データパイプライン
npm runanalyze-sessions は、~/.claude/projects/ から JSONL セッション ファイルを読み取り、構造化された JSON を public/data/sessions/ に出力します。
~/.claude/projects/**/*.jsonl
-> ターンを解析して構築する
-> メタデータ、サブエージェント、リンクされたプランを抽出
-> セッションの要約 (中心性ベースの代表者プロンプト、workType 分類)
-> TF-IDF + Tool-IDF + 構造特徴 -> 切り捨てられた SVD -> 凝集クラスタリング
-> /insights ファセット統合 (狭いクラスターのマージ、目標ベースのラベル付け)
-> ルーヴァンコミュニティの検出
-> スキル合成 (再利用性スコア トップ N -> クロード -p、ファセットの洞察が充実)
-> 出力:
public/data/sessions/index.json (セッションリスト)
public/data/sessions/overview.json (セッション間分析 + ナレッジ グラフ)
public/data/sessions/detail/*.json (個々のセッション再生データ)
カスタムパス:
npm runanalyze-sessions -- --sessions-dir /path/to/sessions --output-dir /path/to/output
/insights の統合
分析前に分析セッションを自動化する

/insights を定期的に実行してファセット データを更新します。これにより、パイプラインが次のように強化されます。
TF-IDF キーワードの代わりに LLM で生成された Underground_goal を使用したトピック ラベルの改善
共有目標カテゴリに基づいた狭いクラスターの結合
セッションの結果と有用性を使用した、品質を意識した再利用性スコアリング
摩擦の詳細と成功率を示す、より豊富な合成プロンプト
ファセットの Brief_summary を使用したセッションの概要
--skip-facets でスキップするか、 --facets-dir <path> でファセット パスをカスタマイズします。
セッション リストには、自動生成された概要が表示されます。 /insights ファセット データが利用可能な場合は、LLM によって生成された Brief_summary が使用されます。それ以外の場合、要約は LLM なしでローカルで処理されます。
代表的なプロンプトの選択 : 位置の重み付けによる Jaccard 中心性を使用して、プラン モード ターンから最も代表的なプロンプトを選択します。
workType 分類 : 各セッションを次の 4 つのタイプのいずれかに自動的に分類します。
調査 --- 探索的で読み取りの多いセッション
実装 --- 編集/書き込みの多いコーディング セッション
デバッグ --- 編集を伴う Bash 負荷の高いセッション
計画 --- ツール呼び出しがほとんどない計画モード セッション
キーワード抽出 : セッションプロンプトから上位のキーワードを抽出します。
スコープの推定 : 編集されたファイルの共通ディレクトリプレフィックスからセッションスコープを推測します。
ナレッジグラフ分析から再利用可能なスキル定義を合成します。人類学/スキルの形式に従います。
事前合成 : 分析セッション中に claude -p を介して上位 5 つのスキル候補を自動的に合成します。
即時表示 : 合成されたスキルは UI ですぐに表示およびコピー可能です
オンデマンドの再合成: 再合成ボタンは、グラフのコンテキスト (関連するトピック、コミュニティ、中心性) を強化した完全なバージョンを生成します。
ローカルサーバー: npm run skill-server は合成 API サーバーを開始します (localhost:3456)
# 指定する

モデル (例: スピードに関する俳句)
npm run 分析セッション -- --Synthetic-model 俳句
# 合成する候補の数を変更します (デフォルト: 5)
npm run 分析セッション -- --合成カウント 10
# LLM 合成をスキップ (グラフ構築のみ)
npm run 分析セッション -- --skip-Synthetic
スクリプト
コマンド
説明
npm 実行開発
Vite開発サーバーを起動します
npm ビルドを実行する
型チェック + 本番ビルド
npm 実行プレビュー
本番ビルドのプレビュー
npmでlintを実行する
ESLint を実行する
npm run 分析セッション
データパイプラインを実行する
npm run スキルサーバー
スキル合成用ローカルサーバー(localhost:3456)
npm run dev:full
スキルサーバー + Vite 開発サーバーの組み合わせ
分析セッションのオプション
旗
説明
--合成モデル <モデル>
合成に使用するモデル (例: スピードの俳句)
--合成数 <n>
合成する候補の数 (デフォルト: 5)
--スキップ合成
LLM合成をスキップする
--facets-dir <パス>
カスタム ファセット ディレクトリ (デフォルト: ~/.claude/usage-data/facets )
--スキップファセット
/insights の更新とファセットの統合をスキップする
技術スタック
反応力グラフ-2d (d3-force)
プレーン CSS (CSS-in-JS、Tailwind なし)
ソース/
コンポーネント/
概要/ # ダッシュボード カード、セッション リスト、チャート
playback/ # セッションリプレイ、ツール呼び出しブロック、サブエージェントブランチ
知識/ # フォースグラフ、ノード詳細、暗黙知、スキル合成
hooks/ # データの取得 (useSessionIndex、useSessionDetail、useSessionOverview)
types/ # TypeScript の型定義
スクリプト/
cli.ts # npx CLI エントリ ポイント
session-parser.ts # JSONL 解析、ターン構築、メタデータ抽出
analyze-sessions.ts # ダッシュボード JSON パイプライン
session-summarizer.ts # セッションの要約 (ローカル NLP)
skill-synthesizer.ts # スキル合成 (claude -p)
skill-server.ts # 合成 HTTP サーバー
Knowledge-graph-builder.ts # セマンティック埋め込み + グラフ構築
ナレッジグラフ/
facets-reader.ts # /i

nsights ファセット データ リーダー + 正規化
公共/
data/sessions/ # 生成された JSON (gitignored)
前提条件
Claude Code セッション ログ (~/.claude/projects/)
過去のセッションで刻まれた痕跡を解読し、再利用可能なスキルとして復活させます。 crune は、Claude Code セッション ログを分析する静的 Web ダッシュボードで、セッションの再生、分析、セマンティック ナレッジ グラフ、およびスキル合成を提供します。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
5
v0.1.7
最新の
2026 年 4 月 27 日
+ 4 リリース
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Decipher the traces etched in past sessions and resurrect them as reusable skills. crune is a static web dashboard that analyzes Claude Code session logs, providing session playback, analytics, a semantic knowledge graph, and skill synthesis. - chigichan24/crune

GitHub - chigichan24/crune: Decipher the traces etched in past sessions and resurrect them as reusable skills. crune is a static web dashboard that analyzes Claude Code session logs, providing session playback, analytics, a semantic knowledge graph, and skill synthesis. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
chigichan24
/
crune
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
193 Commits 193 Commits .github .github bin bin docs docs public public scripts scripts src src .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md THIRD_PARTY_NOTICES THIRD_PARTY_NOTICES eslint.config.js eslint.config.js index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.app.json tsconfig.app.json tsconfig.cli.json tsconfig.cli.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Claude Code Rune — Traces linger, skills emerge
Decipher the traces etched in past sessions and resurrect them as reusable skills. crune is a static web dashboard and CLI tool that analyzes Claude Code session logs, providing session playback, analytics, a semantic knowledge graph, and skill synthesis.
Session Playback --- Turn-by-turn conversation replay with minimap navigation, tool call inspection, and subagent branch expansion
Overview Dashboard --- Activity heatmap, project distribution, tool usage trends, duration distribution, model usage, and top edited files
Semantic Knowledge Graph --- TF-IDF + Tool-IDF + structural features, Truncated SVD, agglomerative clustering, Louvain community detection, Brandes centrality ( algorithm details )
Tacit Knowledge --- Extracted workflow patterns, tool sequences, and pain points (long sessions, hot files)
Session Summarization --- Automatic session summary and classification (no LLM required), enriched with /insights facets data when available
Skill Synthesis --- Synthesize reusable skill definitions from the knowledge graph ( algorithm details )
Semantic Knowledge Graph (Rune)
Generate reusable Claude Code skill definitions directly from your session logs. No clone required.
npx @chigichan24/crune --dry-run # Preview skill candidates
npx @chigichan24/crune --skip-synthesis # Generate heuristic skills (no LLM)
npx @chigichan24/crune --count 3 --model haiku # LLM-synthesized skills (requires claude CLI)
npx @chigichan24/crune --output-dir ~ /.claude/skills # Install skills directly
Output follows the Claude Code skill format ( <name>/SKILL.md ), ready to use as /skill-name commands.
npm install
# Analyze Claude Code session logs (~/.claude/projects/)
npm run analyze-sessions
# Start dev server
npm run dev
Data Pipeline
npm run analyze-sessions reads JSONL session files from ~/.claude/projects/ and outputs structured JSON to public/data/sessions/ .
~/.claude/projects/**/*.jsonl
-> parse & build turns
-> extract metadata, subagents, linked plans
-> session summarization (centrality-based representative prompt, workType classification)
-> TF-IDF + Tool-IDF + structural features -> Truncated SVD -> agglomerative clustering
-> /insights facets integration (narrow cluster merging, goal-based labeling)
-> Louvain community detection
-> skill synthesis (reusability score top-N -> claude -p, enriched with facets insights)
-> output:
public/data/sessions/index.json (session list)
public/data/sessions/overview.json (cross-session analytics + knowledge graph)
public/data/sessions/detail/*.json (individual session playback data)
Custom paths:
npm run analyze-sessions -- --sessions-dir /path/to/sessions --output-dir /path/to/output
/insights Integration
Before analysis, analyze-sessions automatically runs /insights to refresh facets data. This enriches the pipeline with:
Better topic labels using LLM-generated underlying_goal instead of TF-IDF keywords
Narrow cluster merging based on shared goal categories
Quality-aware reusability scoring using session outcome and helpfulness
Richer synthesis prompts with friction details and success rates
Session summaries using brief_summary from facets
Skip with --skip-facets or customize the facets path with --facets-dir <path> .
The session list displays auto-generated summaries. When /insights facets data is available, the LLM-generated brief_summary is used. Otherwise, summaries are processed locally without LLM.
Representative prompt selection : Selects the most representative prompt from plan mode turns using Jaccard centrality with position weighting
workType classification : Automatically classifies each session into one of four types:
investigation --- Exploratory, read-heavy sessions
implementation --- Edit/write-heavy coding sessions
debugging --- Bash-heavy sessions with edits
planning --- Plan mode sessions with few tool calls
Keyword extraction : Extracts top keywords from session prompts
Scope estimation : Infers the session scope from the common directory prefix of edited files
Synthesizes reusable skill definitions from knowledge graph analysis. Follows the anthropics/skills format.
Pre-synthesis : Automatically synthesizes the top 5 skill candidates via claude -p during analyze-sessions
Instant display : Synthesized skills are immediately viewable and copyable in the UI
On-demand re-synthesis : The re-synthesis button generates a full version enriched with graph context (connected topics, community, centrality)
Local server : npm run skill-server starts the synthesis API server (localhost:3456)
# Specify model (e.g. haiku for speed)
npm run analyze-sessions -- --synthesis-model haiku
# Change number of candidates to synthesize (default: 5)
npm run analyze-sessions -- --synthesis-count 10
# Skip LLM synthesis (graph construction only)
npm run analyze-sessions -- --skip-synthesis
Scripts
Command
Description
npm run dev
Start Vite dev server
npm run build
Type-check + production build
npm run preview
Preview production build
npm run lint
Run ESLint
npm run analyze-sessions
Run data pipeline
npm run skill-server
Skill synthesis local server (localhost:3456)
npm run dev:full
skill-server + Vite dev server together
Options for analyze-sessions
Flag
Description
--synthesis-model <model>
Model to use for synthesis (e.g. haiku for speed)
--synthesis-count <n>
Number of candidates to synthesize (default: 5)
--skip-synthesis
Skip LLM synthesis
--facets-dir <path>
Custom facets directory (default: ~/.claude/usage-data/facets )
--skip-facets
Skip /insights refresh and facets integration
Tech Stack
react-force-graph-2d (d3-force)
Plain CSS (no CSS-in-JS, no Tailwind)
src/
components/
overview/ # Dashboard cards, session list, charts
playback/ # Session replay, tool call blocks, subagent branches
knowledge/ # Force graph, node detail, tacit knowledge, skill synthesis
hooks/ # Data fetching (useSessionIndex, useSessionDetail, useSessionOverview)
types/ # TypeScript type definitions
scripts/
cli.ts # npx CLI entry point
session-parser.ts # JSONL parsing, turn building, metadata extraction
analyze-sessions.ts # Dashboard JSON pipeline
session-summarizer.ts # Session summarization (local NLP)
skill-synthesizer.ts # Skill synthesis (claude -p)
skill-server.ts # Synthesis HTTP server
knowledge-graph-builder.ts # Semantic embedding + graph construction
knowledge-graph/
facets-reader.ts # /insights facets data reader + normalization
public/
data/sessions/ # Generated JSON (gitignored)
Prerequisites
Claude Code session logs at ~/.claude/projects/
Decipher the traces etched in past sessions and resurrect them as reusable skills. crune is a static web dashboard that analyzes Claude Code session logs, providing session playback, analytics, a semantic knowledge graph, and skill synthesis.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
5
v0.1.7
Latest
Apr 27, 2026
+ 4 releases
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
