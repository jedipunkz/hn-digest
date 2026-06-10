---
source: "https://github.com/prakhar1114/ai_mime"
hn_url: "https://news.ycombinator.com/item?id=48480492"
title: "Show HN: AI Mime - Use a screen recording for context instead of prompting"
article_title: "GitHub - prakhar1114/ai_mime: Automation harness that compiles a screen recording into a rerunnable script, with LLMs, only at decision points and self-healing reruns. · GitHub"
author: "prakharjain"
captured_at: "2026-06-10T18:59:44Z"
capture_tool: "hn-digest"
hn_id: 48480492
score: 1
comments: 0
posted_at: "2026-06-10T18:21:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Mime - Use a screen recording for context instead of prompting

- HN: [48480492](https://news.ycombinator.com/item?id=48480492)
- Source: [github.com](https://github.com/prakhar1114/ai_mime)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T18:21:10Z

## Translation

タイトル: HN を表示: AI Mime - プロンプトの代わりにコンテキストとして画面記録を使用します。
記事のタイトル: GitHub - prakhar1114/ai_mime: 決定点と自己修復再実行時にのみ、LLM を使用して画面記録を再実行可能なスクリプトにコンパイルする自動化ハーネス。 · GitHub
説明: 決定ポイントと自己修復再実行時にのみ、LLM を使用して画面記録を再実行可能なスクリプトにコンパイルするオートメーション ハーネス。 - prakhar1114/ai_mime

記事本文:
GitHub - prakhar1114/ai_mime: 決定点と自己修復再実行時にのみ、LLM を使用して画面記録を再実行可能なスクリプトにコンパイルする自動化ハーネス。 · GitHub
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
プラカール1114
/
ai_mime
公共
通知
署名する必要があります

n 通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
173 コミット 173 コミット AppIcon.appiconset AppIcon.appiconset docs docs harness harness フック フック パッケージ/ llm-resolver パッケージ/ llm-resolver スクリプト スクリプト src/ ai_mime src/ ai_mime テスト テスト .gitignore .gitignore .gitmodules .gitmodules .python-version .python-version AppIcon.icns AppIcon.icns ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml user_config.yml user_config.yml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
タスクを一度記録します。永遠に実行してください。
AI Mime は自動化ハーネスです。繰り返し実行している自分自身を画面記録します
タスク。エージェントはそれを監視し、エンドツーエンドで実行する方法を学習し、それをコンパイルして
高速で決定的なスクリプト — 反復可能な部分のプレーンコード、LLM のみ
本物の決定点ではコンピュータを使用できない表面にのみ使用します。
他の方法で自動化できます。タスクのコンテキストはポータブルなクロード スキルにキャプチャされます。
そのため、今後の実行はすべて高速かつ低コストで再現可能です。
AI Mime はあなたのスキルも実行します。新しい入力を入力して実行します。環境が
シフトして実行が中断されると、エージェントはスクリプトを新しい状態に自己修復します。
再実行が失敗したままになるのではなく、環境を改善します。
実際の動作を確認する · 例 · インストール · クイック デモ フロー · 仕組み · 開発者ガイド
1. 記録
2. 構築する
3. 走る
タスクを 1 回実行すると、クリック、入力、画面が自動的にキャプチャされます。
エージェントはそれを高速で決定論的なスキル、つまり脊椎用のコード、決定点でのみ LLM にコンパイルします。
新しい入力で再生します。環境が変化した場合、失敗するのではなく自己修復します。
はじめに
AI Mime をローカルでクローン作成、構成、実行するための高速パス。
OS : macOS 13+ (Windows su

ポートはまだ利用できません）
ランタイム : Python >= 3.12、< 3.13
システム権限 : アクセシビリティ、画面録画、入力監視
エージェント ランタイム: Anthropic / Claude Code またはオンボーディング ウィザードによる OpenAI / Codex。カスタムプロバイダーは user_config.yml を通じて構成できます。
git clone --recurse-submodules https://github.com/prakhar1114/ai_mime
cd ai_mime
uv venv .venv
ソース .venv/bin/activate
uv pip install -e 。
# ブラウザ自動化スキルが必要です。
uv ツールのインストール --python .venv/bin/python --with-editablepackages/llm-resolver harness/browser-harness
ソースからビルドしたくないですか?事前に構築されたデスクトップ アプリを入手します。
通常のセットアップ パスの .env または user_config.yml を手動で作成する必要はありません。
初めて起動すると、ネイティブ macOS オンボーディング ウィザードの指示に従って次の手順が実行されます。
アクセシビリティと画面録画の権限を付与します。
Anthropic / Claude Code または OpenAI / Codex を選択します。
API キーの保存またはローカル CLI ログインの検出。
アプリ管理のブラウザー ハーネスをインストールします。
後でダッシュボードからプロバイダー設定を更新できます。
ソース .venv/bin/activate
start_app
次に、メニュー バー アプリまたはダッシュボードを使用します。
録画 : [録画開始] をクリックし、繰り返しタスクを 1 回実行します。
Reflect : AI Mime は、キャプチャしたトレースを再利用可能なセマンティック ワークフローにコンパイルします。
ビルド スキル: ビルド エージェントは入力と出力を確認し、実行計画を最適化し、 workflows/<id>/skills/<slug>/ にポータブル スキルを作成します。
Run : スキルを開き、新しい入力を指定して実行します。
Inspect : 生成されたアーティファクトを読み取り、 workflows/ ディレクトリから実行履歴を読み取ります。
パス
含まれるもの
レコーディング/<id>/manifest.jsonl
キャプチャされた生のイベント ストリーム: スクリーンショット、クリック、入力、ホットキー、抽出、メモ。
ワークフロー/<id>/schema.json
レコードから生成される座標フリーのセマンティック ワークフロー

NG。
ワークフロー/<id>/optimized_plan.json
script 、browser_harness 、または ui_agent ステップを選択する Executor プラン。
ワークフロー/<id>/スキル/<スラッグ>/
run.sh 、 scripts/run.py 、入力、およびフォールバック参照を備えた Claude Skill 互換のポータブル パッケージ。
ワークフロー/<id>/runs/
実行ごとのログ、出力、コピーされたアセット、および再生概要。
デスクトップアプリ
macOS では、ソース ビルド ステップのないパッケージ化されたデスクトップ アプリを利用できます。
ダウンロード後、アプリをアプリケーションにドラッグし、起動時に要求された権限を付与します。
AI Mime は、説明するよりも見せる方が簡単な作業に適しています。
内部ツール、スプレッドシート、Web ポータルにわたるデータ入力。
API が弱いか欠落しているシステムからレポートを取得する。
通常の自動化ツールではきれいにカバーできないブラウザ + ネイティブ macOS ワークフロー。
人間がワークフローを一度定義し、必要に応じて出力をレビューする必要がある反復タスク。
これは、オープンエンドの研究、創造的な生成、完全に自律的な意思決定、またはスケジュールされた cron スタイルのジョブにはまだ適していません。
デモンストレーションによる学習 : エージェントは録画を見て、タスクをエンドツーエンドで学習します。トリガー、セレクター、フィールド マップを接続する必要はありません。
決定的優先コンパイル: 反復可能なスパインは、プレーンな実行可能なコードになります。 LLM は真の決定点でのみ呼び出されます。コンピューターの使用は、他のものが到達できない表面のために予約されています。
ポータブルなクロード スキル: 完成したすべてのタスクは自己完結型のスキル (読み取り可能なコード、JSON 入力コントラクト、キャプチャされたコンテキスト) であり、再実行のたびに高速かつ低コストで実行されます。
自己修復再実行 : 環境が変化して実行が中断されると、エージェントは失敗せずにジョブを終了し、スクリプトを新しい環境にパッチします。
グラフLR
A["1. 画面 + アクションを記録する"] --> B["2. セマンティック ワークフローを反映する"]
B --> C["3. 最適化<br/>実行

「r計画」]
C --> D["4.ビルドスキル<br/>ポータブルパッケージ」]
D --> E["5.実行/修復<br/>実行、回復、パッチ"]
読み込み中
Record は、クリック、キーストローク、スクリーンショット、抽出、およびユーザーメモを Recordings/<id>/ にキャプチャします。
Reflect は、トレースを、タスク パラメーター、サブタスク、座標フリー ステップを含む再利用可能な schema.json に変換します。
Optimize は、決定論的なスクリプト ステップを優先して、optimized_plan.json を書き込み、次に、browser_harness 、次に真の GUI のみの作業用の ui_agent を書き込みます。
Build Skill は、 run.sh 、 scripts/run.py 、入力テンプレート、およびreferences/fallback_plan.md を使用してポータブル スキル パッケージを作成します。
Run / Heal は、新しい入力を使用してスキルを実行します。変更された環境で障害が発生した場合、エージェントはログを検査し、実行を完了し、スクリプトを自己修復して、次回の実行が再び決定的になるようにします。
生成されたパッケージは実行可能で読み取り可能なため、スキル ディレクトリを Claude Code または Codex に公開して、ターミナル エージェントにオートメーションを直接呼び出すこともできます。
画面上で行うすべての反復タスクは、手作業で行うプログラミングです。既存の自動化ツールでは、通常、別の文法の使用を強制されます。
Zapier / Make にはトリガーとフィールド マップが必要で、API が存在する場合にのみ機能します。
ノードベースのビルダーは強力ですが、インターフェイスが作業になります。
RPA は多くの場合、脆弱なセレクターと特殊な実装作業に依存しています。
コンピュータを使用するエージェントはより多くの表面に到達できますが、実行のたびにタスクを再解決します。
AI Mime は、決定論優先、エージェントオンフォールバックのハイブリッドを使用します。反復可能なスパイン用のスクリプト、限定された判断を求める LLM、Web 自動化用のブラウザハーネス、タスクが本当に必要とする場合にのみネイティブ コンピューターを使用します。
今日できることとできないこと
反復的で実証可能なタスク。
Web ポータル、スプレッドシート、ファイル、従来のデスクトップ アプリに関連するタスク。
ワークフ

共通パスが高速に実行される必要があるが、回復が重要な場合は低くなる。
次のステップに進む前に出力をレビューする必要がある人間参加型のワークフロー。
任意の質問に回答したり、レポートや画像を最初から作成したりできます。
オープンエンドの判断または一か八かの決断。
ほぼ同じ方法で 2 回デモンストレーションすることはできません。
経験則: 冗長性を入力し、クリエイティブを出力します。毎回同じ方法で行うのであれば、それがここに当てはまります。判断力はあなたにあります。
開発者ガイド: セットアップ、パッケージ レイアウト、コマンド、ランタイム環境、および手動スキルの実行。
アーキテクチャ : 現在の記録 -> 反映 -> 最適化 -> スキルの構築 -> パイプラインの実行/修復。
AI Mime は、チーム向けのより大規模な実行可能プレイブック レイヤーのシングル オペレーター コアです。
各スキルの視覚的なフローチャートビュー
スケジューリング、Webhook、自然言語トリガー
取り消し不能なアクションに対する人間参加型ゲート
AI Mime でクリック数が節約できる場合は、⭐を付けてください
スターは他の人がプロジェクトを見つけるのを助け、それが構築する価値があることを私たちに知らせます。
壊れたランを見つけましたか?ログと再現の詳細について問題を提起します。
意思決定ポイントと自己修復再実行時にのみ、LLM を使用して画面記録を再実行可能なスクリプトにコンパイルする自動化ハーネス。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
3
v2.0.1
最新の
2026 年 6 月 10 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Automation harness that compiles a screen recording into a rerunnable script, with LLMs, only at decision points and self-healing reruns. - prakhar1114/ai_mime

GitHub - prakhar1114/ai_mime: Automation harness that compiles a screen recording into a rerunnable script, with LLMs, only at decision points and self-healing reruns. · GitHub
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
prakhar1114
/
ai_mime
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
173 Commits 173 Commits AppIcon.appiconset AppIcon.appiconset docs docs harness harness hooks hooks packages/ llm-resolver packages/ llm-resolver scripts scripts src/ ai_mime src/ ai_mime tests tests .gitignore .gitignore .gitmodules .gitmodules .python-version .python-version AppIcon.icns AppIcon.icns LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml user_config.yml user_config.yml uv.lock uv.lock View all files Repository files navigation
Record a task once. Run it forever.
AI Mime is an automation harness. Screen record yourself doing any repetitive
task. An agent watches, learns to do it end-to-end, then compiles it into a
fast, deterministic script — plain code for the repeatable parts, an LLM only
at the genuine decision points, and computer-use only for surfaces that can't be
automated any other way. The task context is captured into a portable Claude Skill
so every future run is fast, cheap, and repeatable.
AI Mime also runs your skills. Provide new inputs and run. When the environment
shifts and a run breaks, the agent self-heals the script to the new
environment instead of leaving you with a failed rerun.
See it in action · Examples · Install · Quick Demo Flow · How it works · Developer guide
1. Record
2. Build
3. Run
Do the task once — clicks, typing, and screens are captured automatically.
An agent compiles it into a fast, deterministic skill — code for the spine, an LLM only at decision points.
Replay with new inputs. If the environment changed, it self-heals instead of failing.
Getting Started
The fast path for cloning, configuring, and running AI Mime locally.
OS : macOS 13+ (Windows support not yet available)
Runtime : Python >= 3.12, < 3.13
System permissions : Accessibility, Screen Recording, and Input Monitoring
Agent runtime : Anthropic / Claude Code or OpenAI / Codex through the onboarding wizard. Custom providers can be configured through user_config.yml .
git clone --recurse-submodules https://github.com/prakhar1114/ai_mime
cd ai_mime
uv venv .venv
source .venv/bin/activate
uv pip install -e .
# Required for browser automation skills.
uv tool install --python .venv/bin/python --with-editable packages/llm-resolver harness/browser-harness
Prefer not to build from source? Grab the prebuilt Desktop App .
You do not need to manually create .env or user_config.yml for the normal setup path.
On first launch, the native macOS onboarding wizard guides you through:
Granting Accessibility and Screen Recording permissions.
Selecting Anthropic / Claude Code or OpenAI / Codex.
Saving API keys or detecting a local CLI login.
Installing the app-managed browser harness.
You can update provider settings later from the dashboard.
source .venv/bin/activate
start_app
Then use the menu bar app or dashboard:
Record : click Start Recording and perform a repetitive task once.
Reflect : AI Mime compiles the captured trace into a reusable semantic workflow.
Build Skill : the build agent confirms inputs and outputs, optimizes the execution plan, and creates a portable skill under workflows/<id>/skills/<slug>/ .
Run : open the skill, provide new inputs, and run it.
Inspect : read the generated artifacts and run history from the workflows/ directory.
Path
What it contains
recordings/<id>/manifest.jsonl
Raw captured event stream: screenshots, clicks, typing, hotkeys, extracts, and notes.
workflows/<id>/schema.json
Coordinate-free semantic workflow generated from the recording.
workflows/<id>/optimized_plan.json
Executor plan that chooses script , browser_harness , or ui_agent steps.
workflows/<id>/skills/<slug>/
Claude Skill-compatible portable package with run.sh , scripts/run.py , inputs, and fallback references.
workflows/<id>/runs/
Per-run logs, outputs, copied assets, and replay summaries.
Desktop App
A packaged desktop app with no source build step is available for macOS.
After downloading, drag the app to Applications and grant the requested permissions on launch.
AI Mime is for work that is easier to show than to describe:
Data entry across internal tools, spreadsheets, and web portals.
Pulling reports from systems with weak or missing APIs.
Browser + native macOS workflows that normal automation tools cannot cover cleanly.
Repetitive tasks where a human should define the workflow once, then review outputs as needed.
It is not the right fit yet for open-ended research, creative generation, fully autonomous decisions, or scheduled cron-style jobs.
Learn by demonstration : an agent watches your recording and learns the task end-to-end — no triggers, selectors, or field maps to wire up.
Deterministic-first compilation : the repeatable spine becomes plain runnable code; an LLM is invoked only at genuine decision points; computer-use is reserved for surfaces nothing else can reach.
Portable Claude Skills : every finished task is a self-contained skill — readable code, a JSON input contract, and captured context — that runs fast and cheap on every rerun.
Self-healing reruns : when the environment shifts and a run breaks, the agent finishes the job and patches the script to the new environment instead of failing.
graph LR
A["1. Record<br/>screens + actions"] --> B["2. Reflect<br/>semantic workflow"]
B --> C["3. Optimize<br/>executor plan"]
C --> D["4. Build Skill<br/>portable package"]
D --> E["5. Run / Heal<br/>execute, recover, patch"]
Loading
Record captures clicks, keystrokes, screenshots, extracts, and user notes into recordings/<id>/ .
Reflect converts the trace into a reusable schema.json with task parameters, subtasks, and coordinate-free steps.
Optimize writes optimized_plan.json , preferring deterministic script steps, then browser_harness , then ui_agent for true GUI-only work.
Build Skill creates a portable skill package with run.sh , scripts/run.py , input templates, and references/fallback_plan.md .
Run / Heal executes the skill with new inputs. If it breaks on a changed environment, the agent inspects logs, completes the run, and self-heals the script so the next run is deterministic again.
Because the generated package is executable and readable, you can also expose the skill directory to Claude Code or Codex and let terminal agents call the automation directly.
Every repetitive task you do on a screen is programming you are doing by hand. Existing automation tools usually force you into a different grammar:
Zapier / Make need triggers and field maps, and only work where APIs exist.
Node-based builders can be powerful, but the interface becomes the work.
RPA often depends on brittle selectors and specialized implementation effort.
Computer-use agents can reach more surfaces, but they re-solve the task every run.
AI Mime uses a deterministic-first, agent-on-fallback hybrid: scripts for the repeatable spine, LLM calls for bounded judgment, browser harnesses for web automation, and native computer-use only where the task genuinely needs it.
What It Can And Cannot Do Today
Repetitive, demonstrable tasks.
Tasks that touch web portals, spreadsheets, files, and legacy desktop apps.
Workflows where the common path should run fast but recovery matters.
Human-in-the-loop workflows where outputs should be reviewed before the next step.
Answering arbitrary questions or generating reports/images from scratch.
Open-ended judgment or high-stakes decisions.
Anything you cannot demonstrate roughly the same way twice.
Rule of thumb: redundant in, creative out. If you would do it the same way every time, it belongs here. Judgment stays with you.
Developer guide : setup, package layout, commands, runtime environment, and manual skill execution.
Architecture : current Record -> Reflect -> Optimize -> Build Skill -> Run/Heal pipeline.
AI Mime is the single-operator core of a larger executable playbook layer for teams.
Visual flowchart view of each skill
Scheduling, webhooks, natural-language triggers
Human-in-the-loop gates for irreversible actions
If AI Mime saves you clicks, give it a ⭐
Stars help others find the project - and let us know it's worth building.
Found a broken run? Open an issue with logs and reproduction details.
Automation harness that compiles a screen recording into a rerunnable script, with LLMs, only at decision points and self-healing reruns.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
3
v2.0.1
Latest
Jun 10, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
