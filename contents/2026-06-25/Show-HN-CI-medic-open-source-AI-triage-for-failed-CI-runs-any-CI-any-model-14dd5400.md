---
source: "https://github.com/alitariq4589/ci-medic"
hn_url: "https://news.ycombinator.com/item?id=48675935"
title: "Show HN: CI-medic – open-source AI triage for failed CI runs (any CI, any model)"
article_title: "GitHub - alitariq4589/ci-medic: AI triage for failed CI runs. Privacy-first, vendor-agnostic. · GitHub"
author: "alitariq4589"
captured_at: "2026-06-25T16:45:14Z"
capture_tool: "hn-digest"
hn_id: 48675935
score: 1
comments: 0
posted_at: "2026-06-25T16:37:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CI-medic – open-source AI triage for failed CI runs (any CI, any model)

- HN: [48675935](https://news.ycombinator.com/item?id=48675935)
- Source: [github.com](https://github.com/alitariq4589/ci-medic)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T16:37:33Z

## Translation

タイトル: Show HN: CI-medic – 失敗した CI 実行に対するオープンソース AI トリアージ (任意の CI、任意のモデル)
記事のタイトル: GitHub - alitariq4589/ci-medic: 失敗した CI 実行に対する AI トリアージ。プライバシーを最優先し、ベンダーに依存しません。 · GitHub
説明: 失敗した CI 実行に対する AI トリアージ。プライバシーを最優先し、ベンダーに依存しません。 - alitariq4589/ci-medic

記事本文:
GitHub - alitariq4589/ci-medic: 失敗した CI 実行に対する AI トリアージ。プライバシーを最優先し、ベンダーに依存しません。 · GitHub
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
アリタリク4589
/
シメディック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
まい

n ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
28 コミット 28 コミット 資産 アセット ドキュメント ドキュメント サンプル サンプル src/ ci_medic src/ ci_medic テスト テスト .gitignore .gitignore Dockerfile Dockerfile ライセンス ライセンス README.md README.md action.yml action.yml pull.sh pull.sh pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
失敗した CI 実行に対する AI トリアージ。パイプラインがレッドになると、ci-medic はログを読み取り、ノイズの中から実際のエラーを見つけ、AI ベースの判定 ( code 、 flake 、 infra 、 dependency 、または config ) を、すでに見ている場所にポストします。
シークレットは、モデルが参照する前に編集されます。ローカル モデルを指定すると、ログがネットワークから離れることはありません。
「このエラーを説明する」ためにクリックするボタンではありません。 ci-medic は障害時に自動的に実行され、数千のノイズの多いログ行を抽出して実際のエラーを抽出し、機密情報を編集して、単一のホスト型 AI アシスタントでは到達できない CI システムに構造化された判定をポストします。
GitHub アクション: スティッキー編集可能な PR コメント、セットアップ不要
Jenkins : ビルドの説明 (エージェントでの 1 回限りのセットアップが必要です。ロードマップ上のネイティブ プラグイン)
現在、他の CI はログ ファイルに対して CLI を使用できます。
失敗した場合にのみ実行される優先順位付けジョブを追加します。
# .github/workflows/ci.yml
トリアージ :
if : 失敗()
need : [build, test] # 優先順位を付けるジョブ
実行: ubuntu-最新
権限:
アクション: 読み取り # 失敗したジョブログの読み取り
プルリクエスト : 書き込み # コメントを投稿
手順:
- 使用: alitariq4589/ci-medic@v0.1.0
付き:
github トークン: ${{ Secrets.GITHUB_TOKEN }}
api-key : ${{ Secrets.CI_MEDIC_API_KEY }} # オプション (モデルプロバイダーを参照)
github-token は、パイプライン内の組み込みトークンを自動的に使用します。 api-key を使用しない場合でも、ci-medic は AI の判定なしで、抽出されたエラー ウィンドウを投稿します。コメントがベタベタして再録される

適切な場所に更新しないでください。
スクリーンショット付きの完全なチュートリアル: docs/github-actions.md
Jenkins にはツールをパイプラインにプロビジョニングするマーケットプレイスがないため、セットアップには 2 つの役割があります。管理者が ci-medic をエージェントで一度利用できるようにし、その後パイプライン作成者が小さなブロックを追加します。 (管理ステップを削除するネイティブ Jenkins プラグインがロードマップにあります。)
1. 管理者のセットアップ: エージェントで ci-medic を利用できるようにします (1 回限り)。次のいずれか:
エージェント環境にインストールします: pip install git+https://github.com/alitariq4589/ci-medic
または、Docker ベースのエージェントを実行する場合は、それをエージェント イメージに追加します。
モデル キーを CI_MEDIC_API_KEY 環境変数 (またはステージに挿入された Jenkins 資格情報) としてエージェントに提供します。
2. パイプライン作成者のセットアップ: 優先順位付けブロックを追加します。ビルド出力をファイルに保存し、失敗した場合に優先順位を付けます。
ポスト {
失敗 {
スクリプト {
デフォルトの評決 = sh(
スクリプト: ' ci-medic jenkins --file ci-medic-console.log --job "$JOB_NAME" ' 、
returnStdout : true
）。トリム()
currentBuild 。説明 = 評決
}
}
}
ビルドステージがログを書き込む場所:
sh ''' #!/bin/bash
set -o パイプ失敗
{ ./あなたのビルドとテスト; 2>&1 | ci-medic-console.log を参照してください。
「」
set -o Pipefail は、エラー時にステージが失敗したままになります (したがって、失敗ブロックが起動します)。評決はビルド ページの上部に表示されます。 ci-medic はファイルを読み取るだけなので、Jenkins API トークンやスクリプト承認は必要ありません。
→ スクリーンショット付きの完全なチュートリアル: docs/jenkins.md
任意のログ ファイルに対して ci-medic を実行します (CI 統合は必要ありません)。
pip install git+https://github.com/alitariq4589/ci-medic
# 蒸留のみ: キーがないため、マシンから何も出ません:
ci-medic 分析 --file failed.log --no-llm
# 完全な AI ベースのトリアージ:
import CI_MEDIC_API_KEY= " your-key " # openrouter、anthropic、openai API トークンを使用できます
ci-medic 分析 --file failed.log
モデルプロバイダー
ci-medicの作品

OpenAI 互換 API または Anthropic API を使用します。環境変数または .ci-medic.yml を使用してプロバイダーを設定します。
OpenRouter : 1 つのキー、多数のモデル (無料モデルも含まれるため、これがデフォルトです。デフォルトの openrouter 選択モデルは src/ci_medic/config.py にあります) - テスト済み
Ollama : ローカル、ゼロ出力 - 未テスト
llama.cpp ( llama-server ): ローカル、ゼロ出力 - 未テスト
その他の OpenAI 互換エンドポイント
import CI_MEDIC_API_KEY= " your-key "
import CI_MEDIC_BASE_URL= " https://openrouter.ai/api/v1 " # プロバイダー エンドポイント
import CI_MEDIC_MODEL= " your-model " # プロバイダーのモデル ID
モデルを選択するのはあなたです。 ci-medic は、あなたを 1 つに縛ることはありません。モデルが指定されていない場合、優先チェーンを試行し、レート制限またはクレジットなしで失敗し、どちらを使用したかを記録します。
ci-medic のコードはオープン ソース (Apache-2.0) です。GitHub ではパブリック アクションを参照し、自由にフォークしたりベンダーしたりできます。 「プライベート」に制御されるのはデータ、つまりログの保存場所です。
ci-medic をローカル モデル サーバーにポイントすると、ログ コンテンツがネットワークから流出することはありません。
import CI_MEDIC_BASE_URL= " http://localhost:11434/v1 " # Ollama (または llama.cpp の llama-server)
import CI_MEDIC_MODEL= " ローカルモデル "
# クラウドキーを設定しない
蒸留と秘密編集の手順は、プロバイダーに関係なく常にローカルで実行されます。抽出され編集されたウィンドウのみがモデルに送信され、ローカル モデルの場合は、それさえもマシン上に残ります。
モデル呼び出しの前に、ci-medic は既知の秘密形式を編集し、未知の形式の高ランダム文字列に対してエントロピー フィルターを実行します。 (カバーされている正確なパターンについては、docs/redaction.md を参照してください。)ゼロエグレスの場合は、ローカル モデルを使用し、クラウド キーを設定しません。
リポジトリ ルート内のオプションの .ci-medic.yml (examples/.ci-medic.yml を参照):
プロバイダー: openai-compat # または: anthropic
Base_url : https://openrouter.ai/ap

i/v1
char_budget : 12000
ignore_jobs : ["lint", "codeql"] # これらのジョブをスキップします
extra_signals : ["MY_CUSTOM_ERROR_TOKEN"] # エラーとして扱う追加の文字列
環境変数はファイルをオーバーライドします。
GitHub Actions と Jenkins (CLI ベース) は v0.1 で出荷されます。
Jenkins プラグイン (アップデート センターからインストール、エージェントのセットアップなし)、
実行全体にわたる慢性的なフレークを追跡する、フレークテストメモリ
失敗した CI 実行に対する AI トリアージ。プライバシーを最優先し、ベンダーに依存しません。
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

AI triage for failed CI runs. Privacy-first, vendor-agnostic. - alitariq4589/ci-medic

GitHub - alitariq4589/ci-medic: AI triage for failed CI runs. Privacy-first, vendor-agnostic. · GitHub
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
alitariq4589
/
ci-medic
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
28 Commits 28 Commits assets assets docs docs examples examples src/ ci_medic src/ ci_medic tests tests .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md action.yml action.yml pull.sh pull.sh pyproject.toml pyproject.toml View all files Repository files navigation
AI triage for failed CI runs. When your pipeline goes red, ci-medic reads the logs, finds the real error in the noise, and posts an AI-based verdict — code , flake , infra , dependency , or config — right where you already look.
Secrets are redacted before any model sees them. Point it at a local model and your logs never leave your network.
Not a button you click to "explain this error". ci-medic runs automatically on failure, distills thousands of noisy log lines to the real error, redacts secrets, and posts a structured verdict on CI systems where single hosted AI assistant doesn't reach.
GitHub Actions : sticky editable PR comment, zero setup
Jenkins : build description (requires one-time setup on your agent; native plugin on the roadmap)
Any other CI can use the CLI against a log file today.
Add a triage job that runs only on failure:
# .github/workflows/ci.yml
triage :
if : failure()
needs : [build, test] # the jobs to triage
runs-on : ubuntu-latest
permissions :
actions : read # read failed job logs
pull-requests : write # post the comment
steps :
- uses : alitariq4589/ci-medic@v0.1.0
with :
github-token : ${{ secrets.GITHUB_TOKEN }}
api-key : ${{ secrets.CI_MEDIC_API_KEY }} # optional (see Model providers)
github-token uses the built-in token in the pipeline automatically. Without api-key , ci-medic still posts the extracted error window, just without the AI verdict. The comment is sticky and re-runs update it in place.
Full walkthrough with screenshots: docs/github-actions.md
Jenkins has no marketplace that provisions tools into pipelines, so setup has two roles: an admin makes ci-medic available on the agent once, then pipeline authors add a small block. (A native Jenkins plugin that removes the admin step is on the roadmap.)
1. Admin setup: make ci-medic available on your agent (one-time). Any one of:
Install it on the agent environment: pip install git+https://github.com/alitariq4589/ci-medic
Or, if you run Docker-based agents, add it to your agent image.
Provide the model key to the agent as the CI_MEDIC_API_KEY environment variable (or a Jenkins credential injected into the stage).
2. Pipeline author setup: add the triage block. Tee build output to a file, then triage it on failure:
post {
failure {
script {
def verdict = sh(
script : ' ci-medic jenkins --file ci-medic-console.log --job "$JOB_NAME" ' ,
returnStdout : true
) . trim()
currentBuild . description = verdict
}
}
}
where your build stage writes the log:
sh ''' #!/bin/bash
set -o pipefail
{ ./your-build-and-test; } 2>&1 | tee ci-medic-console.log
'''
set -o pipefail keeps the stage failing on error (so the failure block fires); the verdict lands at the top of the build page. No Jenkins API token and no script-approval needed because ci-medic only reads a file.
→ Full walkthrough with screenshots: docs/jenkins.md
Run ci-medic on any log file (no CI integration required):
pip install git+https://github.com/alitariq4589/ci-medic
# Distill only: No key, nothing leaves your machine:
ci-medic analyze --file failed.log --no-llm
# Full AI-based triage:
export CI_MEDIC_API_KEY= " your-key " # can be openrouter, anthropic, openai API token
ci-medic analyze --file failed.log
Model providers
ci-medic works with any OpenAI-compatible API or the Anthropic API. Set the provider via environment variables or .ci-medic.yml :
OpenRouter : one key, many models (this is default as it includes free models too. The default openrouter selected models are in src/ci_medic/config.py ) - Tested
Ollama : local, zero egress - Untested
llama.cpp ( llama-server ): local, zero egress - Untested
Any other OpenAI-compatible endpoint
export CI_MEDIC_API_KEY= " your-key "
export CI_MEDIC_BASE_URL= " https://openrouter.ai/api/v1 " # provider endpoint
export CI_MEDIC_MODEL= " your-model " # provider's model id
You choose the model; ci-medic doesn't lock you to one. With no model specified, it tries a priority chain and falls through on rate-limit or no-credit, recording which it used.
ci-medic's code is open source (Apache-2.0) — on GitHub you reference the public action, and you're free to fork or vendor it. What "private" controls is your data : where your logs go.
Point ci-medic at a local model server and no log content ever leaves your network :
export CI_MEDIC_BASE_URL= " http://localhost:11434/v1 " # Ollama (or llama.cpp's llama-server)
export CI_MEDIC_MODEL= " your-local-model "
# set no cloud key
The distillation and secret-redaction steps always run locally regardless of provider; only the distilled, redacted window is ever sent to a model, and with a local model, even that stays on your machine.
Before any model call, ci-medic redacts known secret formats and runs an entropy filter for high-randomness strings in unknown formats. (See docs/redaction.md for the exact patterns covered.) For zero egress, use a local model and set no cloud key.
Optional .ci-medic.yml in your repo root (see examples/.ci-medic.yml ):
provider : openai-compat # or: anthropic
base_url : https://openrouter.ai/api/v1
char_budget : 12000
ignore_jobs : ["lint", "codeql"] # skip these jobs
extra_signals : ["MY_CUSTOM_ERROR_TOKEN"] # extra strings to treat as errors
Environment variables override the file.
GitHub Actions and Jenkins (CLI-based) ship in v0.1.
Jenkins plugin (install from the Update Center, no agent setup),
flaky-test memory that tracks chronic flakes across runs
AI triage for failed CI runs. Privacy-first, vendor-agnostic.
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
