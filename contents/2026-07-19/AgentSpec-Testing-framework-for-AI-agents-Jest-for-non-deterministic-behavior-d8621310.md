---
source: "https://github.com/Ozperium/agentspec"
hn_url: "https://news.ycombinator.com/item?id=48970483"
title: "AgentSpec: Testing framework for AI agents (Jest for non-deterministic behavior)"
article_title: "GitHub - Ozperium/agentspec: Testing framework for AI agents — Jest for non-deterministic AI behavior · GitHub"
author: "pawfromoz"
captured_at: "2026-07-19T18:52:05Z"
capture_tool: "hn-digest"
hn_id: 48970483
score: 1
comments: 1
posted_at: "2026-07-19T18:21:11Z"
tags:
  - hacker-news
  - translated
---

# AgentSpec: Testing framework for AI agents (Jest for non-deterministic behavior)

- HN: [48970483](https://news.ycombinator.com/item?id=48970483)
- Source: [github.com](https://github.com/Ozperium/agentspec)
- Score: 1
- Comments: 1
- Posted: 2026-07-19T18:21:11Z

## Translation

タイトル: AgentSpec: AI エージェントのテスト フレームワーク (非決定的動作の Jest)
記事のタイトル: GitHub - Ozperium/agentspec: AI エージェントのテスト フレームワーク — 非決定的な AI 動作のための Jest · GitHub
説明: AI エージェントのテスト フレームワーク — 非決定的な AI 動作用の Jest - Ozperium/agentspec

記事本文:
GitHub - Ozperium/agentspec: AI エージェントのテスト フレームワーク — 非決定的な AI 動作のための Jest · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
オズペリウム
/
エージェントスペック
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .agentspec .agentspec .github/ workflows .github/ workflows 例 例 ランディング ランディング src src テスト テスト .gitignore .gitignore COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SHOW-HN.md SHOW-HN.md action.yml action.yml devto-article.md devto-article.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントのテスト フレームワーク — 非決定的な AI 動作用の Jest
AI エージェントは非決定的です。プロンプトを変更したり、モデルを交換したり、ツールを更新したりすると、動作は予測できない形で変化します。 AgentSpec を使用すると、本番環境に到達する前にそれらの変化をキャッチするテストを作成できます。
npm install -g @ozperium/agentspec
エージェントスペックの初期化
これにより、agentspec.yaml が作成されます。
名前: " my-agent-tests "
テスト:
- 名前: 「挨拶を処理します」
入力:「こんにちは」
期待する:
含まれるもの:「ヘルプ」
テストを実行します。
エージェントスペックの実行
非決定的な出力のアサーション
AgentSpec は、文字列の等価性だけでなく、AI 出力用に設計されたアサーションを提供します。
テスト:
- name : "期限切れのトークンを処理します"
入力:「トークンの有効期限が切れました」
期待する:
含まれる内容:「リフレッシュトークン」
not_contains : " わかりません "
max_latency_ms : 5000
- 名前: 「ルート請求の質問」
入力:「返金してほしい」
期待する:
contains_any : [「請求」、「返金」、「サポート」]
呼び出されたツール:「請求先への転送」
- name : 「応答は意味的に主題に沿っています」
input : 「パスワードをリセットするにはどうすればよいですか?」
期待する:
semantically_similar : "パスワード認証情報をリセット"
min_confidence : 0.5
- name : " JSON 出力の構造は正しい "
入力:「ユーザープロファイルを取得」
期待する:
json_pat

h: " data.email "
json_value : " user@example.com "
- name : " エラー コード パターンと一致します "
入力:「エラーをシミュレート」
期待する:
正規表現: " ERR- \\ d{5} "
裁判官としてのLLM
ローカル モデルを使用して応答品質を評価します。API コストなし、完全なプライバシー:
# 地元のオラマを審査員として起用する
Agentspec run --judge-endpoint http://127.0.0.1:11434/v1/chat/completions --judge-model qwen2.5:7b
テスト:
- 名前: " 回答は役に立ちました "
input : 「パスワードをリセットするにはどうすればよいですか?」
期待する:
llm_judge : " 応答にはパスワードのリセット方法が説明されているはずです。"
動作差分レポート
AgentSpec は、テストごとに最後に合格した出力を保存します。動作が変化すると、何が変化したかが正確にわかります。
⚠️ REGRESSION サポート エージェント > 期限切れのトークンを処理します
⚠ 動作が REGRESSED — テストは合格していましたが、現在は失敗しています
+ 追加: あなたの、トークン、持っています、期限切れです、連絡してください、サポート
- 削除されました: あなた、必要、更新、トークン、設定、セキュリティ
実際のエージェントのテスト
--endpoint を使用して、HTTP アクセス可能なエージェントをテストします。
Agentspec run --endpoint https://my-agent.example.com/chat
AgentSpec はエンドポイントに {input: "..."} を POST し、応答で {output: "..."} を期待します。
# 失敗時の終了コード 1
エージェントスペック実行 --ci
# CI システム用の JUnit XML
Agentspec run --junit > results.xml
# JSON出力
エージェントスペック実行 --json
GitHub アクション
- 使用:agentspec/action@v1
付き:
テストディレクトリ: テスト
フォーマット : ジュニット
CLIコマンド
Agentspec run [--dir テスト] [--ci] [--json] [--junit] [--watch] [--test "パターン"]
エージェントスペックの初期化
エージェントスペック リスト [--dir テスト]
エージェントスペックのバージョン
AgentSpec を選ぶ理由
問題
エージェントスペック
モデルの更新によるエージェントの動作の変更
回帰テストで変化を捉える
「昨日はうまくいきました」デバッグ
何が変更されたかを示す差分レポート
各展開前の手動テスト
--ci 終了コードを使用した CI/CD ゲート
Jest/Vitest は非決定的な出力を処理しません
contains_any 、

regex 、 llm_judge 、 semantically_similar
AI 機能の CI 統合なし
GitHub アクション + JUnit XML
ロードマップ
コア アサーション (contains、not_contains、contains_any、contains_all、regex)
メタデータ アサーション (max_latency_ms、max_tokens、tool_called)
意味上の類似性（単語の重複）
run/init/list コマンドを使用した CLI
動作差分レポート — 実行間で何が変わったかを確認します
LLM-as-judge — ローカルモデル (Ollama) を使用して出力品質を評価します
HTTP エージェント アダプター — --endpoint 経由で実際のエージェントをテストする
会話履歴からの自動テスト生成
GitHub PR の統合 (PR に関するコメントと動作の差分)
AI エージェントのテスト フレームワーク — 非決定的な AI 動作用の Jest
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

Testing framework for AI agents — Jest for non-deterministic AI behavior - Ozperium/agentspec

GitHub - Ozperium/agentspec: Testing framework for AI agents — Jest for non-deterministic AI behavior · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Ozperium
/
agentspec
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .agentspec .agentspec .github/ workflows .github/ workflows examples examples landing landing src src tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SHOW-HN.md SHOW-HN.md action.yml action.yml devto-article.md devto-article.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Testing framework for AI agents — Jest for non-deterministic AI behavior
AI agents are non-deterministic. When you change a prompt, swap a model, or update a tool, behavior shifts in ways you can't predict. AgentSpec lets you write tests that catch those shifts before they reach production.
npm install -g @ozperium/agentspec
agentspec init
This creates agentspec.yaml :
name : " my-agent-tests "
tests :
- name : " handles greeting "
input : " hello "
expect :
contains : " help "
Run tests:
agentspec run
Assertions for non-deterministic output
AgentSpec provides assertions designed for AI output, not just string equality:
tests :
- name : " handles expired token "
input : " my token expired "
expect :
contains : " refresh token "
not_contains : " I don't know "
max_latency_ms : 5000
- name : " routes billing question "
input : " I want a refund "
expect :
contains_any : ["billing", "refund", "support"]
tool_called : " transfer_to_billing "
- name : " response is semantically on-topic "
input : " how do I reset my password? "
expect :
semantically_similar : " reset password credentials "
min_confidence : 0.5
- name : " JSON output has correct structure "
input : " get user profile "
expect :
json_path : " data.email "
json_value : " user@example.com "
- name : " matches error code pattern "
input : " simulate error "
expect :
regex : " ERR- \\ d{5} "
LLM-as-judge
Use a local model to evaluate response quality — no API costs, full privacy:
# Use local Ollama as judge
agentspec run --judge-endpoint http://127.0.0.1:11434/v1/chat/completions --judge-model qwen2.5:7b
tests :
- name : " response is helpful "
input : " How do I reset my password? "
expect :
llm_judge : " The response should explain how to reset a password "
Behavior diff reports
AgentSpec stores the last passing output per test. When behavior changes, you see exactly what shifted:
⚠️ REGRESSION support agent > handles expired token
⚠ Behavior REGRESSED — test was passing, now failing
+ added: your, token, has, expired, please, contact, support
- removed: you, need, refresh, the, token, settings, security
Testing real agents
Use --endpoint to test any HTTP-accessible agent:
agentspec run --endpoint https://my-agent.example.com/chat
AgentSpec POSTs {input: "..."} to your endpoint and expects {output: "..."} in the response.
# Exit code 1 on failure
agentspec run --ci
# JUnit XML for CI systems
agentspec run --junit > results.xml
# JSON output
agentspec run --json
GitHub Action
- uses : agentspec/action@v1
with :
test-dir : tests
format : junit
CLI commands
agentspec run [--dir tests] [--ci] [--json] [--junit] [--watch] [--test "pattern"]
agentspec init
agentspec list [--dir tests]
agentspec version
Why AgentSpec?
Problem
AgentSpec
Agent behavior changes with model updates
Regression tests catch the shift
"It worked yesterday" debugging
Diff reports show what changed
Manual testing before each deploy
CI/CD gates with --ci exit codes
Jest/Vitest don't handle non-deterministic output
contains_any , regex , llm_judge , semantically_similar
No CI integration for AI features
GitHub Action + JUnit XML
Roadmap
Core assertions (contains, not_contains, contains_any, contains_all, regex)
Metadata assertions (max_latency_ms, max_tokens, tool_called)
Semantic similarity (word overlap)
CLI with run/init/list commands
Behavior diff reports — see what changed between runs
LLM-as-judge — use a local model (Ollama) to evaluate output quality
HTTP agent adapter — test real agents via --endpoint
Auto-test generation from conversation history
GitHub PR integration (comment on PRs with behavior diffs)
Testing framework for AI agents — Jest for non-deterministic AI behavior
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
