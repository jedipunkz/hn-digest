---
source: "https://github.com/kwadwoadu/modelfit"
hn_url: "https://news.ycombinator.com/item?id=48662127"
title: "ModelFit – find the cheapest LLM that can back up your main coding model"
article_title: "GitHub - kwadwoadu/modelfit: Benchmark LLMs on your own codebase. Repo-specific probes, blind rubric-based judging, and correctness-first rankings. · GitHub"
author: "kwadwoadu"
captured_at: "2026-06-24T16:42:27Z"
capture_tool: "hn-digest"
hn_id: 48662127
score: 1
comments: 1
posted_at: "2026-06-24T16:18:31Z"
tags:
  - hacker-news
  - translated
---

# ModelFit – find the cheapest LLM that can back up your main coding model

- HN: [48662127](https://news.ycombinator.com/item?id=48662127)
- Source: [github.com](https://github.com/kwadwoadu/modelfit)
- Score: 1
- Comments: 1
- Posted: 2026-06-24T16:18:31Z

## Translation

タイトル: ModelFit – メインのコーディング モデルをバックアップできる最も安価な LLM を見つけます
記事のタイトル: GitHub - kwadwoadu/modelfit: 独自のコードベースで LLM をベンチマークします。リポジトリ固有のプローブ、ブラインド ルーブリック ベースの判定、正確性優先のランキング。 · GitHub
説明: 独自のコードベースで LLM をベンチマークします。リポジトリ固有のプローブ、ブラインド ルーブリック ベースの判定、正確性優先のランキング。 - クワドゥ/モデルフィット

記事本文:
GitHub - kwadwoadu/modelfit: 独自のコードベースで LLM をベンチマークします。リポジトリ固有のプローブ、ブラインド ルーブリック ベースの判定、正確性優先のランキング。 · GitHub
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
クワドゥ
/
モデルフィット
公共
通知
通知を変更するにはサインインする必要があります

イオン設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット .claude/ コマンド .claude/ コマンド .github .github アセット アセット bin bin config config docs docs 例 例 プローブ プローブ プロンプト プロンプト テスト テスト .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md results.example.csv results.example.csv すべてのファイルを表示 リポジトリ ファイルのナビゲーション
他人のベンチマークではなく、自分のコードベースに最適な LLM を見つけてください。
ModelFit は、候補モデル全体でリポジトリ固有のコーディング プローブを実行し、明示的なルーブリックに照らして答えを盲目的に評価し、コストや遅延の前に正確さをランク付けします。公開ベンチマークは平均的なコードを測定します。 ModelFit は、安価なモデルまたはセカンダリ モデルが SwiftUI、Drizzle 移行、Cloudflare Worker、および障害モードに対応できるかどうかを尋ねます。
ターゲット リポジトリ ──▶ プローブ (PROMPT + RUBRIC) ──▶ run.sh ──▶ 回答候補
│
試行.csv + 評決.csv ◀── 裁判官.sh ◀───────┘
│
report.sh ──▶ カバレッジを意識したリーダーボード
なぜ違うのか
汎用スイートではなく、独自のワークフロー。プローブは、明示的に指定したターゲット リポジトリから生成されます。
対応機種なら何でも。 OpenAI 互換の /chat/completions および Anthropic 互換の /v1/messages エンドポイント。
ブラインドルーブリック採点。審査員は候補モデル名ではなく、課題、ルーブリック、解答を参照します。
まずは正しさ。コストとレイテンシによって正確性の損失が救われることはありません。
監査可能な実行。すべての実行は、不変の実行 ID、サンプルごとの出力、試行台帳および判定台帳を取得します。
ModelFit は、シークレットと実行出力が Git から除外されるように設計されています。

デフォルトですが、機密データが漏洩しないことを保証できるローカル ツールはありません。
config/models.json には、キーを保持する環境変数名のみが保存されます。実際のキーはシェルまたは .env 内に存在し、gitignored されます。
.env 、 config/models.json 、runs/、および results.csv は無視されます。
bin/scan-secrets.sh は、公開する前に、追跡されたファイルに共通のシークレット形式の文字列がないかチェックします。
生成されたプローブには、独自のコード、顧客データ、認証情報、または個人データが含まれる場合があります。プローブを実行する前に確認してください。
プローブ プロンプトは、構成されている各候補プロバイダーに送信されます。課題、ルーブリック、および回答候補が審査員プロバイダーに送信されます。
git clone https://github.com/kwadwoadu/modelfit.git
CDモデルフィット
brew install jq shellcheck # ローカル lint 用の shellcheck はオプションです
./bin/selftest.sh # API 消費ゼロ。モックプロバイダーテストが含まれます
cp config/models.example.json config/models.json # モデルの編集 + 判断
cp .env.example .env # キーを貼り付けます。決してコミットしないでください
./bin/modelfit Doctor --repo ../your-app
ModelFit リポジトリからの Claude コードを使用してプローブを生成します。
/modelfit --repo ../your-app
次に、スイート全体の前に 1 つのプローブ/モデルをスモークテストします。
./bin/modelfit 実行 example-chunk fake-model-key --samples 1
./bin/modelfit ジャッジ example-chunk fake-model-key
./bin/modelfit レポート
フルラン:
プローブの p の場合 / * .md ;する
n= $( ベース名 " $p " .md )
./bin/modelfit run " $n " all --samples 1
./bin/modelfit ジャッジ " $n " すべて
完了しました
./bin/modelfit レポート
1 つのモデルが失敗した場合、バッチは可能な限り続行されますが、ゼロ以外で終了し、レポートには不完全なカバレッジが表示されます。
エージェント生成のプローブ。 /modelfit --repo ../your-app を実行します。このコマンドはターゲット リポジトリを検査し、6 ～ 10 個のプローブをプローブ/ に書き込み、機密性のない来歴を記録します。
手動プローブ。コピープローブ/example-*.md : 各モデルに送信される # プロンプトと #

審査員が採点するルーブリック。
優れたプローブには、決定的な識別子が 1 つあります。それは、弱いモデルが間違える微妙な点です。
run.sh は、各プローブを候補に送信し、マークダウン フェンスを取り除き、空の/切り捨てられた応答をトークンの上限まで再試行し、すべての試行を run/<run-id>/attempts.csv に記録します。
judge.sh は、タスク + ルーブリック + 信頼できない候補者の回答を審査員に送信し、厳密な JSON 判定を検証して、 run/<run-id>/verdicts.csv を書き込みます。
report.sh は、合格率、品質、候補者のコストによってランク付けされ、判定された回数、試行回数、不完全な試行回数、および実際に記録された総コストを表示します。 --by-task を追加すると、プローブごとの候補コストの内訳 (どのモデルでどの種類のタスクが高価になるか) が表示されます。
候補コスト、判定コスト、再試行コストは、プロバイダー トークンの使用状況 (利用可能な場合) から追跡されます。欠落している使用量は NA であり、ゼロではありません。
LLM ジャッジは有益ですが、客観的ではありません。ブラインドラベルはモデル同一性のバイアスを軽減します。スタイルの偏りや即時注入のリスクは除去されません。
ジャッジ専用プローブは候補コードを実行しません。コンパイルが決定的な場合は、将来のプローブに実行可能ゲートを追加します。
config/models.example.json 内の価格はプレースホルダーです。コストの比較を信頼する前に、プロバイダーの価格を確認してください。
1 つのサンプルは統計的に信頼できるものではありません。実行間の差異が重要な場合は、 --samples N を使用します。
プロバイダーの「互換性」は異なります。大規模な実行の前に、./bin/modelfit Doctor とスモーク プローブを使用します。
モデルフィット/
§─ bin/modelfit run.sh judge.sh report.sh Doctor.sh selftest.sh scan-secrets.sh
§─ bin/lib/common.sh
§─ config/models.example.json
§─ プローブ/ example-honesty.md example-chunk.md
§─ プロンプト/generate-probes.md judge-system.md
§─ テスト/モックプロバイダーの信頼性テスト
§─ .claude/commands/modelfit.md
§─ 例/ results.example.csv .env.example .gitignore LIC

エンセ
MITライセンス取得済み。クワドウォ・アドゥによって建てられました。
独自のコードベースで LLM をベンチマークします。リポジトリ固有のプローブ、ブラインド ルーブリック ベースの判定、正確性優先のランキング。
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

Benchmark LLMs on your own codebase. Repo-specific probes, blind rubric-based judging, and correctness-first rankings. - kwadwoadu/modelfit

GitHub - kwadwoadu/modelfit: Benchmark LLMs on your own codebase. Repo-specific probes, blind rubric-based judging, and correctness-first rankings. · GitHub
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
kwadwoadu
/
modelfit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits .claude/ commands .claude/ commands .github .github assets assets bin bin config config docs docs examples examples probes probes prompts prompts tests tests .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md results.example.csv results.example.csv View all files Repository files navigation
Find the best LLM for your codebase—not someone else’s benchmark.
ModelFit runs repo-specific coding probes across candidate models, grades their answers blindly against explicit rubrics, and ranks correctness before cost and latency. Public benchmarks measure average code; ModelFit asks whether a cheaper or secondary model can handle your SwiftUI, your Drizzle migrations, your Cloudflare Worker, and your failure modes.
target repo ──▶ probes (PROMPT + RUBRIC) ──▶ run.sh ──▶ candidate answers
│
attempts.csv + verdicts.csv ◀── judge.sh ◀────────────┘
│
report.sh ──▶ coverage-aware leaderboard
Why it is different
Your workflow, not a generic suite. Probes are generated from a target repo you name explicitly.
Any compatible model. OpenAI-compatible /chat/completions and Anthropic-compatible /v1/messages endpoints.
Blind rubric grading. The judge sees the task, rubric and answer, not the candidate model name.
Correctness first. Cost and latency never rescue a correctness loss.
Auditable runs. Every run gets an immutable run ID, per-sample outputs, attempt ledger and verdict ledger.
ModelFit is designed so secrets and run outputs are excluded from Git by default, but no local tool can guarantee you will never leak sensitive data.
config/models.json stores only the environment variable names that hold keys. The real keys live in your shell or .env , which is gitignored.
.env , config/models.json , runs/ and results.csv are ignored.
bin/scan-secrets.sh checks tracked files for common secret-shaped strings before publishing.
Generated probes may contain proprietary code, customer data, credentials or personal data. Review probes before running them.
Probe prompts are sent to each configured candidate provider. Task, rubric and candidate answer are sent to the judge provider.
git clone https://github.com/kwadwoadu/modelfit.git
cd modelfit
brew install jq shellcheck # shellcheck optional, for local linting
./bin/selftest.sh # zero API spend; includes mock-provider tests
cp config/models.example.json config/models.json # edit models + judge
cp .env.example .env # paste keys; never commit
./bin/modelfit doctor --repo ../your-app
Generate probes with Claude Code from the ModelFit repo:
/modelfit --repo ../your-app
Then smoke-test one probe/model before the full suite:
./bin/modelfit run example-chunk fake-model-key --samples 1
./bin/modelfit judge example-chunk fake-model-key
./bin/modelfit report
Full run:
for p in probes/ * .md ; do
n= $( basename " $p " .md )
./bin/modelfit run " $n " all --samples 1
./bin/modelfit judge " $n " all
done
./bin/modelfit report
If one model fails, the batch continues where possible but exits non-zero and the report shows incomplete coverage.
Agent-generated probes. Run /modelfit --repo ../your-app . The command inspects the target repository, writes 6–10 probes into probes/ , and records non-sensitive provenance.
Manual probes. Copy probes/example-*.md : a # PROMPT sent to each model and a # RUBRIC the judge grades against.
A good probe has one decisive discriminator: the subtle thing a weaker model gets wrong.
run.sh sends each probe to candidates, strips markdown fences, retries empty/truncated replies up to the token ceiling, and records every attempt in runs/<run-id>/attempts.csv .
judge.sh sends task + rubric + untrusted candidate answer to the judge, validates strict JSON verdicts, and writes runs/<run-id>/verdicts.csv .
report.sh ranks by pass percentage, quality and candidate cost, while showing judged count, attempts, incomplete attempts and actual recorded total cost. Add --by-task for a per-probe candidate-cost breakdown (which kinds of task are expensive on which model).
Candidate cost, judge cost and retry cost are tracked from provider token usage when available. Missing usage is NA , not zero.
LLM judges are useful but not objective. Blind labels reduce model-identity bias; they do not remove style bias or prompt-injection risk.
Judge-only probes do not execute candidate code. If compilation is decisive, add an executable gate in a future probe.
Prices in config/models.example.json are placeholders. Verify provider pricing before trusting cost comparisons.
One sample is not statistical confidence. Use --samples N when run-to-run variance matters.
Provider “compatibility” varies. Use ./bin/modelfit doctor and a smoke probe before a large run.
modelfit/
├─ bin/ modelfit run.sh judge.sh report.sh doctor.sh selftest.sh scan-secrets.sh
├─ bin/lib/common.sh
├─ config/ models.example.json
├─ probes/ example-honesty.md example-chunk.md
├─ prompts/ generate-probes.md judge-system.md
├─ tests/ mock-provider reliability tests
├─ .claude/commands/modelfit.md
├─ examples/ results.example.csv .env.example .gitignore LICENSE
MIT licensed. Built by Kwadwo Adu.
Benchmark LLMs on your own codebase. Repo-specific probes, blind rubric-based judging, and correctness-first rankings.
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
