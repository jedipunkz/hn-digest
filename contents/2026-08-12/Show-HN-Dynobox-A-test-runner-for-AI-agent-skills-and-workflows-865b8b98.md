---
source: "https://github.com/dynobox/dynobox"
hn_url: "https://news.ycombinator.com/item?id=49274758"
title: "Show HN: Dynobox – A test runner for AI agent skills and workflows"
article_title: "GitHub - dynobox/dynobox: Cross-harness testing for multi-step agent flows · GitHub"
author: "bhkdotdev"
captured_at: "2026-08-12T16:46:58Z"
capture_tool: "hn-digest"
hn_id: 49274758
score: 1
comments: 0
posted_at: "2026-08-12T16:14:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Dynobox – A test runner for AI agent skills and workflows

- HN: [49274758](https://news.ycombinator.com/item?id=49274758)
- Source: [github.com](https://github.com/dynobox/dynobox)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T16:14:42Z

## Translation

タイトル: Show HN: Dynobox – AI エージェントのスキルとワークフローのテスト ランナー
記事のタイトル: GitHub - dynobox/dynobox: マルチステップ エージェント フローのクロスハーネス テスト · GitHub
説明: マルチステップのエージェント フローのクロスハーネス テスト。 GitHub でアカウントを作成して、dynobox/dynobox の開発に貢献してください。

記事本文:
GitHub - dynobox/dynobox: マルチステップ エージェント フローのクロスハーネス テスト · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ダイノボックス
/
ダイノボックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
218 コミット 218 コミット .agents/ スキル .agents/ スキル .github .github アプリ アセット アセット ドキュメント ドキュメント パッケージ パッケージ .gitignore .gitignore .prettierr

c.json .prettierrc.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md RELEASES.md RELEASES.md SECURITY.md SECURITY.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
決定的なエージェントの検証。
エージェントとスキルをコーディングするためのオープンソースのテスト ランナー: 実際のワークフローを実行し、
証拠をキャプチャし、ツール、コマンド、ファイル、および回答についてアサートします。
始めましょう ·
仕組み ·
ドキュメント ·
ダッシュボード ·
エージェントのスキル ·
例
ツール呼び出し、コマンド、ファイル、HTTP リクエスト、トランスクリプト、および最終的なメッセージに対するアサート
あるモデルが別のモデルを判断する必要なく答えが得られます。
Claude Code、OpenAI Codex、OpenCode を使用して同じシナリオを実行します (
さらに今後)、環境間で異なる動作を見つけることができます。
新しい一時作業ディレクトリで複数ステップのタスクをテストし、それを繰り返します。
不安定な行動を暴露し、何かが失敗したときの証拠を保管します。
裸の実行可能呼び出しを、シナリオ スコープの静的呼び出し、順次呼び出し、またはシナリオ スコープの呼び出しに置き換えます。
実験的な CLI モックを使用したハンドラー応答。
Node.js 22+ および少なくとも 1 つのサポートされているハーネス (Claude Code、Codex、または OpenCode)
インストールされ、認証され、 PATH 上で使用可能になります。
npx dynobox init # クロードコード (デフォルト)
npx dynobox init --harness codex # OpenAI Codex
npx dynobox init --harness opencode # OpenCode
npxダイノボックスの実行
インストールされ認証されたハーネスに対して init コマンドを 1 つ実行してから、
生成された dyno を実行します。
dynobox init は、 dynobox/example.dyno.mjs にスターター dyno を作成します。
dynobox を実行すると、現在のファイルの下に *.dyno.{mjs,js,ts,mts,yaml,yml} ファイルが検出されます
ディレクトリに移動して実行します

設定されたハーネスに対する継承シナリオ。
信頼できる Dyno のみを実行してください。 JavaScript および TypeScript 構成がインポートされ、
セットアップおよび検証コマンドがマシン上で実行されます。臨時の仕事
ディレクトリはジョブ ファイルを分離しますが、セキュリティ サンドボックスではありません。プロセス
権限に従ってホストにアクセスできます。
ダイノは、プロンプト、治具セットアップ、ハーネス、および受け入れ基準を組み合わせます。
1 つの TypeScript、JavaScript、または YAML ファイル:
# パッケージ.dyno.yaml
名前 : パッケージスクリプト
ハーネス:
- クロードコード
- コーデックス
- オープンコード
シナリオ:
- name : テストスクリプトを検出します
セットアップ：
- |
cat > package.json <<'JSON'
{"スクリプト":{"テスト":"vitest 実行"}}
JSON
プロンプト: >-
cat package.json を使用して、このプロジェクトにテスト スクリプトがあるかどうかを教えてください。
アサーション :
- タイプ: コマンド.コール
実行可能ファイル:猫
コマンド: {args: [package.json]}
- タイプ:tool.notCalled
ツール : edit_file
- タイプ: artifact.contains
パス : package.json
テキスト：vitest run
- タイプ:finalMessage.contains
テキスト: テスト
テストを実行します。
npx dynobox 実行 package.dyno.yaml
同じ形状が TypeScript と JavaScript で機能します。
@ダイノボックス/SDK
(defineDyno、ツール、コマンド、...)。参照
構成オーサリング。
Dynobox は、選択したハーネスを新しい一時作業ディレクトリで起動します。
観察可能な動作を記録し、キャプチャされた各アサーションを評価します。
証拠。
チェックが失敗した場合、出力には予想された内容と観察された内容が表示されます。
証拠
チェック例
答えは何ですか
ツール
ツール.コール済み、ツール.ノットコール済み
エージェントはツールを使用しましたか、それとも回避しましたか?
コマンド
command.called 、command.notCalled
期待されたシェルコマンドが実行されましたか?
ファイル
artifact.exists 、 artifact.contains 、 artifact.unchanged
正しいファイルを作成、変更、または保存しましたか?
スキル
スキル.参照
必要なスキルの指示を参照しましたか?
ネットワーク*

http.called 、 http.notCalled
プロキシ対応の子プロセスが予期されたエンドポイントを呼び出しましたか?
応答
トランスクリプト.contains 、finalMessage.contains
やり取りには必要な情報が含まれていましたか?
ロジック
sequence.inOrder 、 anyOf
観察された行動は受け入れられた経路または順序に従っていましたか?
検証
verify.コマンド
完成した作品はカスタム実行可能ファイルのチェックに合格しますか?
* HTTP キャプチャは、Dynobox のプロキシを尊重する子プロセスのトラフィックを観察し、
CAの設定。独立したネットワークでネイティブ Web ツールとクライアントを利用する
捕獲できない可能性があります。 HTTP キャプチャの方法を学ぶ
作品。
アサーションのリファレンスを参照してください
すべてのマッチャーとオーサリング オプションに対して。
コマンドラインから設定されたハーネスをオーバーライドします。
npx dynobox run --harness claude-code
npx dynobox run --harness codex
npx dynobox run --harness opencode
npx dynobox run --harness claude-code,codex,opencode
すべてのシナリオとハーネスのペアを繰り返して合格率を測定します。
npx dynobox 実行 \
--ハーネス クロード コード、コーデックス、オープンコード \
--反復 5
繰り返し実行すると、.FF.. などのコンパクトなパスレート行が表示されますが、
反復の失敗による診断の証拠。
Dynobox はローカルで実行され、アカウントなしで動作します。端末出力を使用して、
開発、自動化のための --reporter json、または認証
dynobox にログインし、 --save-run を追加して、コンパクトな実行概要を
Dynobox ダッシュボード。 DYNOBOX_UPLOAD_URL を投稿に設定します
Dynobox 認証情報を送信せずに、同じペイロードを独自のエンドポイントに送信します。
保存された実行データには長さの制限がありますが、編集されません。利用可能な Git が含まれています
ID とリビジョンのメタデータ。すべてのジョブには、作成されたアサーション データを含めることができます。
要求されたエンドポイント URL、ツール コマンド、
検証出力。失敗したジョブにはさらにコマンドまたはハーネスが含まれる場合があります
診断。これらの場合は --save-run を使用しないでください。

価値観は共有されるべきではありません。
Dynobox は、モデルによる評価評価よりも単体テストに近いものです。捕獲したものを評価する
動作 - ツール呼び出し、シェルコマンド、ファイル、および HTTP リクエスト - 明示的な
反復可能な主張。
1 回限りのスクリプトと比較して、Dynobox は正規化された新鮮なアサーションを提供します。
一時作業ディレクトリ、取得された失敗の証拠、およびイテレーションの合格率
不安定な動作を測定するため。
パッケージ
説明
ダイノボックス
dyno を検出して実行するための CLI
@ダイノボックス/SDK
dyno をオーサリングするためのタイプセーフなヘルパー
Dynobox は早期アクセス ソフトウェアです。 CLI、SDK、レポート形式は進化する可能性があります
1.0より前。
モノリポジトリをローカルで実行して提案するには、CONTRIBUTING.md を参照してください。
変化。
Dynobox は Apache-2.0 ライセンスを取得しています。
複数ステップのエージェントフローのクロスハーネステスト
Readme Apache-2.0 ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Cross-harness testing for multi-step agent flows. Contribute to dynobox/dynobox development by creating an account on GitHub.

GitHub - dynobox/dynobox: Cross-harness testing for multi-step agent flows · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
dynobox
/
dynobox
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
218 Commits 218 Commits .agents/ skills .agents/ skills .github .github apps apps assets assets docs docs packages packages .gitignore .gitignore .prettierrc.json .prettierrc.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASES.md RELEASES.md SECURITY.md SECURITY.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
Deterministic agent verification.
Open-source test runner for coding agents and skills: run real workflows,
capture evidence, assert on tools, commands, files, and answers.
Get started ·
How it works ·
Documentation ·
Dashboard ·
Agent skills ·
Examples
Assert on tool calls, commands, files, HTTP requests, transcripts, and final
answers without requiring one model to judge another.
Run the same scenario through Claude Code, OpenAI Codex, and OpenCode (with
more to come) to find behavior that varies between environments.
Test multi-step tasks in fresh temporary work directories, repeat them to
expose flaky behavior, and keep the evidence when something fails.
Replace bare executable calls with scenario-scoped static, sequential, or
handler responses using experimental CLI mocks .
Node.js 22+ and at least one supported harness (Claude Code, Codex, or OpenCode)
installed, authenticated, and available on PATH .
npx dynobox init # Claude Code (default)
npx dynobox init --harness codex # OpenAI Codex
npx dynobox init --harness opencode # OpenCode
npx dynobox run
Run one init command for a harness that is installed and authenticated, then
run the generated dyno.
dynobox init creates a starter dyno in dynobox/example.dyno.mjs .
dynobox run discovers *.dyno.{mjs,js,ts,mts,yaml,yml} files below the current
directory and runs their scenarios against the configured harnesses.
Only run dynos you trust. JavaScript and TypeScript configs are imported, and
setup and verification commands execute on your machine. Temporary work
directories separate job files, but they are not security sandboxes; processes
can access the host according to their permissions.
A dyno combines the prompt, fixture setup, harnesses, and acceptance criteria in
one TypeScript, JavaScript, or YAML file:
# package.dyno.yaml
name : package-script
harnesses :
- claude-code
- codex
- opencode
scenarios :
- name : detects the test script
setup :
- |
cat > package.json <<'JSON'
{"scripts":{"test":"vitest run"}}
JSON
prompt : >-
Use cat package.json and tell me whether this project has a test script.
assertions :
- type : command.called
executable : cat
command : {args: [package.json]}
- type : tool.notCalled
tool : edit_file
- type : artifact.contains
path : package.json
text : vitest run
- type : finalMessage.contains
text : test
Run the test:
npx dynobox run package.dyno.yaml
The same shape works in TypeScript and JavaScript via
@dynobox/sdk
( defineDyno , tool , command , ...). See
Config Authoring .
Dynobox launches the selected harnesses in fresh temporary work directories,
records the observable behavior, and evaluates each assertion against captured
evidence.
When a check fails, the output shows what was expected and what was observed.
Evidence
Example checks
What it answers
Tools
tool.called , tool.notCalled
Did the agent use or avoid a tool?
Commands
command.called , command.notCalled
Did it execute the expected shell command?
Files
artifact.exists , artifact.contains , artifact.unchanged
Did it create, change, or preserve the right files?
Skills
skill.referenced
Did it reference the required skill instructions?
Network*
http.called , http.notCalled
Did a proxy-aware child process call the expected endpoint?
Response
transcript.contains , finalMessage.contains
Did the interaction contain required information?
Logic
sequence.inOrder , anyOf
Did the observed behavior follow an accepted path or order?
Verification
verify.command
Does the completed work pass a custom executable check?
* HTTP capture observes child-process traffic that honors Dynobox's proxy and
CA settings. Harness-native web tools and clients with independent networking
may not be captured. Learn how HTTP capture
works .
See the assertion reference
for every matcher and authoring option.
Override the configured harnesses from the command line:
npx dynobox run --harness claude-code
npx dynobox run --harness codex
npx dynobox run --harness opencode
npx dynobox run --harness claude-code,codex,opencode
Repeat every scenario and harness pair to measure pass rates:
npx dynobox run \
--harness claude-code,codex,opencode \
--iterations 5
Repeated runs render compact pass-rate rows such as .FF.. while retaining the
failed iteration evidence for diagnosis.
Dynobox runs locally and works without an account. Use terminal output for
development, --reporter json for automation, or authenticate with
dynobox login and add --save-run to publish a compact run summary to the
Dynobox dashboard . Set DYNOBOX_UPLOAD_URL to post
the same payload to your own endpoint without sending Dynobox credentials.
Saved-run data is length-capped but not redacted. It includes available Git
identity and revision metadata. All jobs can include authored assertion data and
matched evidence such as requested endpoint URLs, tool commands, and
verification output. Failed jobs can additionally include command or harness
diagnostics. Do not use --save-run when those values should not be shared.
Dynobox is closer to unit testing than model-graded evals. It evaluates captured
behavior—tool calls, shell commands, files, and HTTP requests—using explicit,
repeatable assertions.
Compared with one-off scripts, Dynobox provides normalized assertions, fresh
temporary work directories, captured failure evidence, and iteration pass rates
for measuring flaky behavior.
Package
Description
dynobox
CLI for discovering and running dynos
@dynobox/sdk
Type-safe helpers for authoring dynos
Dynobox is early-access software. The CLI, SDK, and report formats may evolve
before 1.0.
See CONTRIBUTING.md to run the monorepo locally and propose
a change.
Dynobox is Apache-2.0 licensed .
Cross-harness testing for multi-step agent flows
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
