---
source: "https://github.com/bryanyzhu/agentic-ai-system-course"
hn_url: "https://news.ycombinator.com/item?id=48616863"
title: "Agentic Systems Course: Learn AI Agents with an AI Coding Agent"
article_title: "GitHub - bryanyzhu/agentic-ai-system-course: Use agent to learn agent - A skeleton course on how to design, build, and operate production AI agents · GitHub"
author: "lum1104"
captured_at: "2026-06-21T10:18:25Z"
capture_tool: "hn-digest"
hn_id: 48616863
score: 1
comments: 0
posted_at: "2026-06-21T08:34:15Z"
tags:
  - hacker-news
  - translated
---

# Agentic Systems Course: Learn AI Agents with an AI Coding Agent

- HN: [48616863](https://news.ycombinator.com/item?id=48616863)
- Source: [github.com](https://github.com/bryanyzhu/agentic-ai-system-course)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T08:34:15Z

## Translation

タイトル: エージェント システム コース: AI コーディング エージェントで AI エージェントを学ぶ
記事のタイトル: GitHub - bryanyzhu/agentic-ai-system-course: エージェントを使用してエージェントを学習する - 本番 AI エージェントの設計、構築、運用方法に関するスケルトン コース · GitHub
説明: エージェントを使用してエージェントを学習する - 実稼働 AI エージェントを設計、構築、および操作する方法に関するスケルトン コース - GitHub - bryanyzhu/agentic-ai-system-course: エージェントを使用してエージェントを学習する - 実稼働 AI エージェントを設計、構築、および操作する方法に関するスケルトン コース

記事本文:
GitHub - bryanyzhu/agentic-ai-system-course: エージェントを使用してエージェントを学習する - 実稼働 AI エージェントを設計、構築、運用する方法に関するスケルトン コース · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ブライアンジュ
/
エージェント-ai-システム-コース
公共
通知
にサインインする必要があります

通知設定を変更する
追加のナビゲーション オプション
コード
bryanyzhu/agentic-ai-system-course
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .claude/ skill/ Agentic-system-reviewer .claude/ skill/ Agentic-system-reviewer course course .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md setup.sh setup.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェント システム コース - エージェントを使用してエージェントを学習する
一緒に学び、構築したい場合は、Discord チャンネルに参加してください。
これは、実稼働 AI エージェントの設計、構築、運用方法に関する 22 章のスケルトン コースであり、自分の AI パートナーと一緒に読むように書かれています。エージェント システムは、単一のプロンプトにのみ応答するのではなく、計画、意思決定、ツールの使用、フィードバックに基づく適応、記憶などによって自律的に目標を追求できる AI システムです。 LLM-wiki にある Andrej Karpathy のアイデア ファイルと同様に、このコースでは骨格を提供し、エージェントがそれに筋肉をつけるのを手伝います。
スケルトン — トレードオフのある、負荷のかかるトピック、パターン、意思決定。
ゆっくりと年齢を重ねるように書かれています。フレームワークの仕様はすぐに腐ります。アーキテクチャパターンではそうではありません。
人間が読むのと同じくらい AI が利用できるように設計されたファイル ペア (コース + AGENTS.md)。
ステップバイステップのチュートリアル。ウォークスルー プロジェクトはありません。
1 つのスタックに結び付けられます。このコースでは、「LangChain を使用する」または「Pydantic AI を使用する」とは決して言いません。 AI パートナーがプロジェクトに適したスタックを提案します。
リファレンスマニュアルです。正確な API 署名が必要な場合は、AI に尋ねるか、ドキュメントを読んでください。
リポジトリのクローンを作成し、通常の IDE で開いてコースのコンテンツを表示します。同時に、AI エージェント (クロード コード/コーデックス) をプロジェクト ルートにポイントし、次のプロンプトのいずれかを試してください。

章を勉強してください:
「これが重要な点について、実際の例を 3 つ挙げてください。」
「あなたが私にインタビューしていると仮定して、このトピックについて、簡単なものから難しいものまで 5 つの追加質問をしてください。」
「私が尋ねるべきで、していない質問は何ですか？」
「[パターン X] について読んだところです。[あなたのプロジェクト] を構築中です。パターンを、どの言語やツールでも機能する最小バージョンに翻訳し、書きながら各部分を説明してください。」
「ちょっと私のプロジェクトのことは忘れてください。OpenCode (または Hermes Agent、またはその他の主要なコーディング エージェント) がこれをどのように処理するのか、そしてそこから何を借りるべきなのかを教えてください。」
また、エージェントに Ch.22 のデザイン キャンバスを指示し、特定のプロジェクトを念頭に置いてそれを実行することもできます。これが、「アイデアがある」から「仕様がある」への最速の道です。
コースに対する PRD、設計ドキュメント、実装計画、またはエージェント コードをレビューします。このスキルを任意のエージェント システムで実行すると、コースに基づいたフィードバックを取得できます。これには、自分のプロジェクト、研究中のオープンソース エージェント、コードが存在する前の PRD、またはセカンドオピニオンが必要な同僚のリポジトリなどがあります。このスキルは、最初にスコープ (趣味/チームツール/顧客対応) を調整し、アーキタイプにとって重要な章を選択して読み、重大度、証拠、コースの引用、具体的な修正を含む結果優先のレポートを作成します。一般的な「見栄えが良い」または「安全性を追加する」レビューではありません。
Claude Code では、「コースに照らしてこれを確認してください」、「このエージェントの設計は適切ですか?」など、必要なことを説明するだけです。 、「何章が抜けているの？」 — エージェントがターゲットのリポジトリまたはドキュメントを指している間。スキルの説明が意図と一致すると、スキルが自動ロードされます。
Codex ユーザー: Codex の公式スキル作成スキルを使用して、このスキルを移植します。
章
テーマ
Ch.00
AI パートナーとこのコースを使用する方法
Ch.01–04
基礎: 1 ～

ol 呼び出し → ループ → コントラクトとしてのツール → プロンプトとキャッシュ
Ch.05–08
記憶と状態: 短期→長期→執筆とキュレーション→持続性
Ch.09–11
調整: 計画 → マルチエージェントの委任 → ハーネス
第12章～第14章
外部表面: 人間参加者 → コネクタ/MCP → スキル/MCP/サブエージェント
第 15 ～ 17 章
本番規模: バックエンド → 可観測性 → コスト、レイテンシー、モデル戦略
第18-19章
品質と運用: 安全性/敵対的入力 → 運用と前方展開
第20–21章
エージェンシー: 積極的なエージェント → 自己進化するエージェント
第22章
デザイン キャンバス: 独自のエージェントをデザインする
各章は順序付けされているため、各章は前に起こったことのみを前提としています。明確なプロジェクトがある場合は、まだ適用されていない章をざっと読んで、適用されたら戻ってください。プロジェクトの形態別の推奨学習パス (コーディング エージェント、パーソナル アシスタント、マルチテナント ツール、リサーチ エージェント、単なる探索) は CLAUDE.md にあります。
目標はコースを完走することではありません。目標は、とにかく出荷したいと思ったものを出荷し、そのすべての行を理解することです。
このコースでは、根拠のある例として 4 つのオープンソース システムを時折取り上げます。
OpenCode — コーディング エージェント (ターミナルファースト、型付きツール、セッション、圧縮)
Hermes Agent — パーソナル アシスタント (メモリ、スキル、cron、チャネル)
OpenClaw — セルフホスト型パーソナル アシスタント ゲートウェイ (チャネル アダプター)
Paperclip — ワークフロー コントロール プレーン (マルチエージェント オーケストレーション、耐久性のある Postgres 状態)
コースから価値を得るためにこれらを複製する必要はありません。これら 4 つは健全性チェックです。根拠のある答えが必要な場合 (「X は実際にコード内で Y をどのように行うのですか?」)、AI パートナーはオンデマンドで関連するリポジトリのクローンを作成することを提案してくれます。あるいは、事前にセットアップを実行することもできます。
./setup.sh
これにより、4 つの参照リポジトリすべてが References/ に複製されます。べき等 (sa

fe を再実行します）。フォークまたはピン留めされたコミットをポイントしたい場合は、Env-var-overridable。詳細については、スクリプトを参照してください。
他のソースについて質問したい場合は、お気軽に参照/に入力して、それについてチャットしてください。
course/ — 22 章 (設計により読み取り専用 — これがスケルトンです)
CLAUDE.md — AI パートナーのための行動ガイド (Claude Code によって読み取られます)
AGENTS.md — CLAUDE.md と同じ内容で、コーデックス / カーソル / 他のエージェントにちなんで名付けられました
README.md — このファイル
setup.sh — オプション: 4 つの参照リポジトリのクローンを作成します。
.claude/skills/ — AI パートナーが呼び出せる組み込みスキル (レビュー担当者、今後も追加予定)
docs/ — AI パートナーがリクエストに応じて作成するメモ（オンデマンドで作成）
質問/ — AI パートナーが自動的に書き込む毎日の Q&A ログ (オンデマンドで作成)
workspace/ — コード、演習、およびプロジェクトのビルド (オンデマンドで作成)
References/ — あなた (またはあなたのエージェント) がクローンを作成する場合の 4 つの参照リポジトリ
CLAUDE.md と AGENTS.md は重複しています。デフォルトでは、異なるエージェントが異なるファイルを検索するため、2 つのファイル名で同じ内容です。どちらもコースと同じように構造化されており、AI アシスタントが 1 回のパスで読み取って有用な動作に変えることができる耐久性のある独自のファイルです。
始める前に最後に一つ
開始するのに許可は必要ありません。エージェントを開く前にすべての章を読む必要はありません。最終的にどのフレームワークを使用することになるのかを知る必要はありません。最初に構築するものは、それらの決定が問題にならないほど十分に小さいものになります。そして、それが目の前にあり、機能するようになると、次の決定が明らかになります。
現在、最も興味深いエージェント システムを出荷している人々は、最も経験豊富な人々ではありません。彼らは、AI パートナーと最初に緊密なループに入り、その状態に最も長く留まった人たちです。このコースは、いつでも

あなたはそのループの中にいるので、何を尋ねるべきか知っています。
「ライセンス」を参照してください。コースのコンテンツは教育用途に公開されています。参照システムは元のライセンスを保持します。
エージェントを使用してエージェントを学習する - 実稼働 AI エージェントの設計、構築、運用方法に関するスケルトン コース
読み込み中にエラーが発生しました。このページをリロードしてください。
57
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Use agent to learn agent - A skeleton course on how to design, build, and operate production AI agents - GitHub - bryanyzhu/agentic-ai-system-course: Use agent to learn agent - A skeleton course on how to design, build, and operate production AI agents

GitHub - bryanyzhu/agentic-ai-system-course: Use agent to learn agent - A skeleton course on how to design, build, and operate production AI agents · GitHub
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
bryanyzhu
/
agentic-ai-system-course
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
bryanyzhu/agentic-ai-system-course
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .claude/ skills/ agentic-system-reviewer .claude/ skills/ agentic-system-reviewer course course .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md setup.sh setup.sh View all files Repository files navigation
Agentic System Course - Use Agent to Learn Agent
Join the discord channel if you want to learn and build together!
This is a 22-chapter skeleton course on how to design, build, and operate production AI agents — written to be read with your own AI partner at your side. An agentic system is an AI system that can autonomously pursue goals by planning, making decisions, using tools, adapting based on feedback, having memory, etc — instead of only responding to a single prompt. Similar to Andrej Karpathy's idea file on LLM-wiki , this course is giving you the skeleton and your agent will help you put the muscles on it .
A skeleton — load-bearing topics, patterns, and decisions, with trade-offs.
Written to age slowly. Framework specifics rot fast; architectural patterns do not.
A file pair (course + AGENTS.md) designed for AI consumption as much as human reading.
A step-by-step tutorial. There is no walked-through project.
Tied to one stack. The course never says "use LangChain" or "use Pydantic AI." Your AI partner suggests the stack that fits your project.
A reference manual. When you need an exact API signature, ask your AI or read the docs.
Clone the repo, open it in your usual IDE to view course content. At the same time, point your AI agent (Claude code/Codex) at the project root, and try one of these prompts when you study a chapter:
"Give me three real-world examples of where this matters."
"Suppose you are interviewing me, quiz me on this topic with five follow-up questions, easy to hard."
"What's a question I should be asking that I haven't?"
"I just read about [pattern X]. I am building [your project]. Translate the pattern into the smallest version that works in whatever language and tools fit, and explain each piece as you write it."
"Forget my project for a moment — show me how OpenCode (or Hermes Agent, or any leading coding agent) handles this, and what we should borrow from it."
You can also just point your agent at Ch.22's design canvas and walk through it with your specific project in mind — that's the fastest path from "I have an idea" to "I have a spec."
Reviews PRDs, design docs, implementation plans, or agent code against the course. You can run this skill on any agentic system to get course-grounded feedback — your own project, an open-source agent you're studying, a PRD before any code exists, or a coworker's repo you want a second opinion on. The skill calibrates scope first (hobby / team tool / customer-facing), picks the chapters that matter for your archetype, reads them, and produces a findings-first report with severity, evidence, course citations, and concrete fixes — not a generic "looks good" or "add safety" review.
In Claude Code, just describe what you want — "review this against the course" , "is this agent design good?" , "what chapters does this miss?" — while the agent is pointed at the target repo or doc. The skill auto-loads when its description matches your intent.
Codex users: use Codex's official skill-creator skill to port this skill over.
Chapters
Theme
Ch.00
How to use this course with your AI partner
Ch.01–04
Foundations: one tool call → the loop → tools as contract → prompts & cache
Ch.05–08
Memory and state: short-term → long-term → writing & curation → persistence
Ch.09–11
Coordination: planning → multi-agent delegation → the harness
Ch.12–14
External surface: human-in-the-loop → connectors/MCP → skills/MCP/subagents
Ch.15–17
Production scale: backend → observability → cost, latency, model strategy
Ch.18–19
Quality and ops: safety/adversarial inputs → operations and forward-deployed
Ch.20–21
Agency: proactive agents → self-evolving agents
Ch.22
Design canvas: designing your own agent
The chapters are ordered so each one only assumes what came before. If you have a clear project, skim chapters that do not apply yet and come back when they do. Suggested learning paths by project shape (coding agent, personal assistant, multi-tenant tool, research agent, just exploring) are in CLAUDE.md .
The goal is not to finish the course. The goal is to ship something you wanted to ship anyway, and to understand every line of it.
The course occasionally points at four open-source systems for grounded examples:
OpenCode — coding agent (terminal-first, typed tools, sessions, compaction)
Hermes Agent — personal assistant (memory, skills, cron, channels)
OpenClaw — self-hosted personal-assistant gateway (channel adapters)
Paperclip — workflow control plane (multi-agent orchestration, durable Postgres state)
You do not need to clone any of these to get value from the course, these four are sanity checks. If you find yourself wanting grounded answers ( "how does X actually do Y in code?" ), your AI partner will offer to clone the relevant repo on demand, or you can run setup beforehand:
./setup.sh
This clones all four reference repos to references/ . Idempotent (safe to re-run). Env-var-overridable if you want to point at forks or pinned commits. See the script for details.
If you want to ask questions about other sources, feel free to put them under references/ and chat about it.
course/ — 22 chapters (read-only by design — this is the skeleton)
CLAUDE.md — Behavioral guide for your AI partner (read by Claude Code)
AGENTS.md — Same content as CLAUDE.md, named for Codex / Cursor / other agents
README.md — This file
setup.sh — Optional: clone the four reference repos
.claude/skills/ — Built-in skills your AI partner can invoke (reviewer; more coming)
docs/ — Written notes your AI partner produces on request (created on demand)
questions/ — Daily Q&A log your AI partner writes automatically (created on demand)
workspace/ — Code, exercises, and project builds (created on demand)
references/ — The four reference repos if you (or your agent) clone them
CLAUDE.md and AGENTS.md are duplicates — same content under two filenames because different agents look for different files by default. Both are structured the same way the course is: a durable, opinionated file any AI assistant can read in one pass and turn into useful behavior.
One last thing before you start
You do not need permission to start. You do not need to read every chapter before opening your agent. You do not need to know which framework you will end up using. The first thing you build will be small enough that none of those decisions matter — and once it's in front of you, working, the next decision becomes obvious.
The people shipping the most interesting agentic systems today are not the ones with the most experience. They are the ones who got into a tight loop with their AI partner first and stayed in it longest. This course exists so that when you are in that loop, you know what to ask .
See LICENSE . The course content is open for educational use. Reference systems retain their original licenses.
Use agent to learn agent - A skeleton course on how to design, build, and operate production AI agents
There was an error while loading. Please reload this page .
57
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
