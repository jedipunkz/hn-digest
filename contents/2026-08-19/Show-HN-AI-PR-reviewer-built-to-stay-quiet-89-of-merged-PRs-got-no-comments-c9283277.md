---
source: "https://github.com/Kyeom1997/pr-sage"
hn_url: "https://news.ycombinator.com/item?id=49358609"
title: "Show HN: AI PR reviewer built to stay quiet – 89% of merged PRs got no comments"
article_title: "GitHub - Kyeom1997/pr-sage: AI-powered GitHub PR reviewer — inline comments + summary via Claude, OpenAI, or Gemini · GitHub"
image: "https://opengraph.githubassets.com/13bdfb468b15f2e0354f9329d0c3b816bc2335a8f933b415a722baf492f9dfb7/Kyeom1997/pr-sage"
author: "gudrua1543"
captured_at: "2026-08-19T08:24:44Z"
capture_tool: "hn-digest"
hn_id: 49358609
score: 1
comments: 0
posted_at: "2026-08-19T08:20:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI PR reviewer built to stay quiet – 89% of merged PRs got no comments

- HN: [49358609](https://news.ycombinator.com/item?id=49358609)
- Source: [github.com](https://github.com/Kyeom1997/pr-sage)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T08:20:28Z

## Translation

タイトル: Show HN: AI PR レビュアーは沈黙を保つように構築されています – マージされた PR の 89% にはコメントがありません
記事のタイトル: GitHub - Kyeom1997/pr-sage: AI を活用した GitHub PR レビューアー — クロード、OpenAI、または Gemini 経由のインライン コメント + 概要 · GitHub
説明: AI を活用した GitHub PR レビューアー — クロード、OpenAI、または Gemini によるインライン コメント + 概要 - Kyeom1997/pr-sage

記事本文:
GitHub - Kyeom1997/pr-sage: AI を活用した GitHub PR レビューアー — インライン コメント + Claude、OpenAI、または Gemini による概要 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
キョム1997
/
賢人向け
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のワークフローに追加するか、新しいワークフローを作成します マーケットプレイスで表示する メイン ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
12 コミット 12 コミット フォルダーとファイル
.github/ ワークフロー .github/ w

orkflows ベンチ結果 ベンチ結果 dist dist docs docs scripts scripts src src test test .gitignore .gitignore ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md action.yml action.yml package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI PR レビューアーは、レビューのノイズを増やすのではなく、排除するように構築されています。
ほとんどの AI レビュー担当者は、プッシュするたびに PR 全体を再レビューし、チームがミュートするまで繰り返します。 pr-sage は、逆の目標を中心に設計されています。つまり、すべてのことを 1 回ずつ言い、チームのルールに従い、新しいことが何もないときは沈黙することです。
🔇 重複コメントはありません。調査結果にはコンテンツの指紋が含まれます。行を移動しても同じコメントが 2 回表示されることはなく、何も変更されていない場合は再実行しても何も投稿されません。
✅ ライフサイクルを見つける。フォローアップレビューでは、どの調査結果が未解決のままで、どの調査結果が修正されたかが報告されます。
⏩ デフォルトでは増分です。最初のレビューの後は、それ以降にプッシュしたコミットのみがレビューされます。ノイズが減り、トークンも減ります。
📏 一般的なアドバイスではなく、あなたのルール。 .pr-sage.json 命令と自動 CLAUDE.md / CONTRIBUTING.md インジェクションにより、チームの規則に従ってレビューが行われます。
🚦 単なる解説ではなく、質の高いゲートです。 --fail-on クリティカル ブロックのマージ。 --event はクリーンな PR を自動的に承認し、実際の問題に関する変更を要求します。
🖥️ PR が存在する前のレビュー。 pr-sage local は、git diff の事前プッシュをレビューします。サーバー、PR、GitHub トークンはありません。
🔐 あなたの鍵、あなたのデータパス。サーバーも何も保存されません。コードは、選択したプロバイダー (Claude、OpenAI、または Gemini) にのみ送信されます。あるいは、自己ホスト型の OpenAI 互換エンドポイント (Ollama、vLLM、LM Studio) を使用してマシンをまったく離れることはありません。 SECURITY.md を参照してください。
CLI、GitHub Action、および

TypeScript ライブラリ。
最速のパス — 設定と GitHub Action ワークフローを書き込み、どのシークレットを登録するかを正確に指示する対話型ウィザードです。
npx pr-sage 初期化
npx pr-sage ドクター
または手動 (CLI):
エクスポート GITHUB_TOKEN=ghp_...
エクスポート ANTHROPIC_API_KEY=sk-ant-...
npx pr-sage レビュー --repo 所有者/名前 --pr 123
何も投稿せずにプレビューする:
npx pr-sage review --repo 所有者/名前 --pr 123 --dry-run
プッシュする前にローカルの変更を確認します (PR や GitHub トークンは必要ありません)。
npx pr-sage local --base main # diff vs main
npx pr-sage local --staged --fail-on クリティカル # ゲートステージング変更
別のプロバイダーで韓国語でレビューします:
エクスポート OPENAI_API_KEY=sk-...
npx pr-sage review --repo 所有者/名前 --pr 123 --provider openai --locale 韓国語
オプション
オプション
デフォルト
説明
-p、--pr <数値>
(必須)
プルリクエスト番号
-r、--repo <所有者/名前>
$GITHUB_REPOSITORY
ターゲットリポジトリ
--プロバイダー<名前>
人間的な
人類 |オープンナイ |ジェミニ
-m、--model <id>
プロバイダのデフォルト
モデル ID ( claude-opus-4-8 、 gpt-5 、 gemini-flash-latest )
--locale <言語>
英語
レビュー出力の言語。 PRのタイトル/本文から自動検出
--paths <グロブ>
—
これらのカンマ区切りのグロブに一致するファイルのみをレビューします (モノリポジトリのスコープ)
--max-tokens <n>
—
コストガード: これだけのトークンが消費されたら、新しいバッチの起動を停止します。
--force
—
ドラフト、WIP タイトル、またはスキップレビューのラベルが付いた PR もレビューします (デフォルトではスキップされます)。
--<パターン> を除外します
—
スキップするコンマ区切りのグロブまたはサブ文字列 (デフォルトに追加: lockfiles、dist/ 、 build/ 、…)
--min-severity <sev>
—
検出結果をこの重大度未満に削除します (例: 提案は要点を隠します)
--fail-on <sev>
—
検出結果がこの重大度以上の場合は 1 を終了します。CI 品質ゲートとして使用します。
--context <モード>
パッチ
full は完全なファイル内容を送信します

精度を高めるためにモデルに追加します (より多くのトークン)
--event <モード>
コメント
クリーンな PR を自動承認し、重要な結果の変更をリクエストします (フォールバックして自分の PR にコメントします)
--検証する
オフ
未確認の所見を拒否する 2 番目のモデル パス
--verify-provider <名前>
同じプロバイダー
検証には別のプロバイダーを使用する
--verify-model <id>
プロバイダのデフォルト
別の検証モデルを使用する
--verify-failure <モード>
中止する
検証が失敗した場合に abort 、 keep 、またはdrop
--output <フォーマット>
テキスト
機械可読な結果の場合は json または sarif
--不完全時のフェイル
オフ
フィルタリング、パッチの欠落、またはトークンバジェットにより変更の一部が未レビューのままになっている場合に失敗します。
--チェック実行
オフ
結果を GitHub Check Run アノテーションとして公開する
--重複排除なし
—
以前の pr-sage レビューによってすでにコメントされた調査結果を再投稿します (重複除去はデフォルトでオンになっています)
--増分なし
—
前回の pr-sage レビュー以降のコミットのみではなく、常に完全な PR diff をレビューします。
--batch-chars <n>
80000
モデルリクエストごとの最大差分文字数
[切り捨てられた]
繰り返しの実行 (例: PR にプッシュされた新しいコミット) では、pr-sage は前回のレビュー以降にプッシュされたコミットのみをレビューし (増分モード)、すでにコメントされている結果をスキップし、新しいものがない場合は何も投稿しません。つまり、重複コメントスパムや無駄なトークンはありません。リポジトリに CLAUDE.md または CONTRIBUTING.md がある場合、レビュー コンテキストとして自動的に挿入されます ( "repoContext": false で無効にします)。 GitHub Enterprise は、$GITHUB_API_URL または githubApiUrl 設定フィールドを介してすぐに使用できます。
実行ごとにトークンの使用量が標準エラー出力に出力されるため (LLM 使用量: N 呼び出し、X 入力 / Y 出力トークン)、コストが表示されたままになります。
各要約ではレビュー範囲も報告されます。部分的なレビューでは PR が自動承認されることはありません。使用する
--fail-on-incomplete は、不完全なカバレッジが CI 品質ゲートで不合格になる必要がある場合に使用します。
OpenAI をポイントする

OpenAI 互換サーバーのプロバイダーとプライベート コードがマシンから離れることはありません。API キーは必要ありません。
オラマ プル qwen2.5-coder:14b
OPENAI_BASE_URL=http://localhost:11434/v1 \
npx pr-sage review --repo owner/repo --pr 123 --provider openai --model qwen2.5-coder:14b
vLLM、LM Studio、または OpenAI チャット完了 API を使用するゲートウェイでも同様に機能します。
GitHub アクションの場合、init --provider self-hosted は、
自己ホスト型ランナー。 localhost は、GitHub でホストされている VM ではなく、そのランナーを参照する必要があります。
実行元のディレクトリに .pr-sage.json を配置します (CLI フラグによりオーバーライドされます)。
{
"プロバイダ" : " 人間 " ,
"ロケール" : " 自動 " 、
"除外" : [ " src/generated/** " , " **/*.snap " ],
"パス" : [ "packages/web/** " ],
"minSeverity" : "提案" ,
"failOn" : " クリティカル " 、
"コンテキスト" : "完全な" ,
"maxTokensPerRun" : 200000 、
"failOnIncomplete" : true 、
"skipLabels" : [ "スキップレビュー " ],
"検証" : true 、
"verifyProvider" : " gemini " ,
"verifyModel" : " gemini-flash-latest " ,
"verifyFailure" : " 中止 " 、
"checkRun" : true 、
"パスルール" : [
{
"パス" : [ "packages/api/** " ],
"instructions" : " パブリック API の下位互換性を確認します。 " ,
"minSeverity" : "提案" ,
"failOn" : " 警告 "
}
]、
"instructions" : " エラー処理には Result<T, E> を使用します。つまり、ドメイン コードでスローされた例外にフラグを立てます。ネストされた条件文よりも早期のリターンを優先します。"
}
指示はレビュープロンプトに挿入されます。レビュー担当者が強制する必要があるチームの規則としてそれを使用します。
# .github/workflows/pr-sage.yml
名前：AIレビュー
に:
プルリクエスト:
タイプ: [オープン、同期]
権限:
内容：読む
プルリクエスト : 書き込み
チェック：書き込み
同時実行性:
グループ: pr-sage-${{ github.event.pull_request.number }}
キャンセル中 : true
ジョブ:
レビュー：
実行: ubuntu-lates

t
手順:
# 信頼できる基本コードから .pr-sage.json を読み取ります。PR 制御されたコードではありません。
- 使用:actions/checkout@v4
付き:
参照: ${{ github.event.pull_request.base.sha }}
永続資格情報: false
- 使用: Kyeom1997/pr-sage@v1
付き:
プロバイダー: anthropic
anthropic-API-key : ${{ Secrets.ANTHROPIC_API_KEY }}
ロケール : 韓国語
フェイルオン : クリティカル # オプション: クリティカルな検出結果のブロック マージ
不完全な場合の失敗: " true "
チェック実行: " true "
SARIF をアップロードするには、 security-events: write を追加し、 sarif: "true" を設定します。アクション
ファーストクラスの path 、 max-tokens 、 verify-provider も公開します。
verify-model 、 verify-failure 、および openai-base-url 入力。
基本 SHA チェックアウトは意図的です: 構成とリポジトリのガイドライン
信頼できる基本コードから取得する必要があります。 PR diff 自体は常に次の方法で取得されます。
GitHub API を使用し、pr-sage は投稿する直前にヘッド SHA を再チェックします。
遅いレビューでは、置き換えられたコミットについてコメントすることはできません。
scripts/bench.mjs は、パブリック リポジトリの最近マージされた PR に対して pr-sage (ドライ、何も投稿されない) を実行し、結果、重大度の組み合わせ、レイテンシ、トークンの使用状況を記録し、さらに有効レビュー率を計算するためのラベル シートを記録します。
ノードスクリプト/bench.mjs --repos fastify/fastify --per-repo 5 --provider gemini
--mode recall はもう 1 つの軸を測定します。人間によるレビュー コメントを受け取ったマージされた PR を選択し、各 PR の最初のコミット (人間がレビューした状態) をレビューし、人間がフラグを立てた場所の数を pr-sage も報告します。手動検証用に横に並べたシートを使用します。
品質ベンチマーク ワークフローでは、いずれかのベンチマークを手動で実行してアップロードできます。
ワークフロー成果物として生成された JSON とラベル付けシート。
fastify、hono、GitHub CLI、Vite で最近マージされた 28 個の PR ( gemini-flash-lite 、ゼロラン失敗):
25/28 (89%) はコメントを生成しませんでした - コード上では静かです

すでに人間による審査を通過していました。その静寂さがポイントです。クリーンなデフにはノイズがありません。
他の 3 人の PR は 7 件の発見を得ました。実際の差分に対する各クレームの検証: 全体で 3/7 が有効、警告の重大​​度が 2/3 (提案層に集中したノイズ)。
これらの PR を --verify で再実行すると、2 つの diff-confirmed-valid 結果 (デプロイメント ワークフローにおける実際の G​​PG 署名回帰質問) が正確に保持され、無効な結果はすべて拒否されました。
PR あたり中央値 2.3 秒、入力トークン約 3.5,000 (Flash-lite のリスト価格では、28 の PR すべてで約 0.01 ドル)。
生の結果と結果ごとの検証ノートは Bench-result/ にあります。注意事項: サンプルが少なく、マージされた PR サンプリングではリコールではなくノイズが測定されます。検出ベンチマーク (後で人間による修正が行われた PR のレビュー前コミットのレビュー) は今後の作業です。
PR メタデータとファイルごとのパッチを GitHub API から取得します。
diff の両側に注釈が付けられます (追加/コンテキスト行には新しいファイル番号が、削除された行には古いファイル番号が付けられます)。そのため、検出結果を削除されたコードにもアンカーできます (例: 「この削除された検証はまだ必要でした」)。ロックファイル/ビルドアーティファクトはフィルターで除外されます。
LLM に構造化レビュー (JSON スキーマ - 解析ヒューリスティックなし) を要求します。概要 + path 、 line 、 severity を含む調査結果、およびオプションの単一行の提案。
検証

[切り捨てられた]

## Original Extract

AI-powered GitHub PR reviewer — inline comments + summary via Claude, OpenAI, or Gemini - Kyeom1997/pr-sage

GitHub - Kyeom1997/pr-sage: AI-powered GitHub PR reviewer — inline comments + summary via Claude, OpenAI, or Gemini · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Kyeom1997
/
pr-sage
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
.github/ workflows .github/ workflows bench-results bench-results dist dist docs docs scripts scripts src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts View all files Repository files navigation
An AI PR reviewer built to eliminate review noise — not add to it.
Most AI reviewers re-review the whole PR on every push and repeat themselves until the team mutes them. pr-sage is designed around the opposite goal: say each thing once, follow your team's rules, and stay silent when there is nothing new to say.
🔇 Zero duplicate comments. Findings carry content fingerprints — a line shift won't make the same comment appear twice, and re-runs post nothing when nothing changed.
✅ Finding lifecycle. Follow-up reviews report which findings remain unresolved and which were fixed.
⏩ Incremental by default. After the first review, only the commits you pushed since get reviewed. Less noise, fewer tokens.
📏 Your rules, not generic advice. .pr-sage.json instructions plus automatic CLAUDE.md / CONTRIBUTING.md injection make reviews follow team conventions.
🚦 A quality gate, not just commentary. --fail-on critical blocks merges; --event auto approves clean PRs and requests changes on real problems.
🖥️ Reviews before the PR exists. pr-sage local reviews your git diff pre-push — no server, no PR, no GitHub token.
🔐 Your keys, your data path. No server, nothing stored; code goes only to the provider you choose — Claude, OpenAI, or Gemini — or never leaves your machine at all with a self-hosted OpenAI-compatible endpoint (Ollama, vLLM, LM Studio). See SECURITY.md .
Ships as a CLI , a GitHub Action , and a TypeScript library .
The fastest path — an interactive wizard that writes your config and the GitHub Action workflow, and tells you exactly which secret to register:
npx pr-sage init
npx pr-sage doctor
Or by hand (CLI):
export GITHUB_TOKEN=ghp_...
export ANTHROPIC_API_KEY=sk-ant-...
npx pr-sage review --repo owner/name --pr 123
Preview without posting anything:
npx pr-sage review --repo owner/name --pr 123 --dry-run
Review your local changes before pushing (no PR, no GitHub token needed):
npx pr-sage local --base main # diff vs main
npx pr-sage local --staged --fail-on critical # gate staged changes
Review in Korean with a different provider:
export OPENAI_API_KEY=sk-...
npx pr-sage review --repo owner/name --pr 123 --provider openai --locale Korean
Options
Option
Default
Description
-p, --pr <number>
(required)
Pull request number
-r, --repo <owner/name>
$GITHUB_REPOSITORY
Target repository
--provider <name>
anthropic
anthropic | openai | gemini
-m, --model <id>
provider default
Model id ( claude-opus-4-8 , gpt-5 , gemini-flash-latest )
--locale <lang>
English
Language for the review output; auto detects it from the PR title/body
--paths <globs>
—
Only review files matching these comma-separated globs (monorepo scoping)
--max-tokens <n>
—
Cost guard: stop launching new batches once this many tokens are spent
--force
—
Review even draft, WIP-titled, or skip-review -labeled PRs (skipped by default)
--exclude <patterns>
—
Comma-separated globs or substrings to skip (added to defaults: lockfiles, dist/ , build/ , …)
--min-severity <sev>
—
Drop findings below this severity (e.g. suggestion hides nitpicks)
--fail-on <sev>
—
Exit 1 if any finding is at or above this severity — use as a CI quality gate
--context <mode>
patch
full sends complete file contents to the model for better accuracy (more tokens)
--event <mode>
comment
auto approves clean PRs and requests changes on critical findings (falls back to comment on your own PRs)
--verify
off
Second model pass that rejects unconfirmed findings
--verify-provider <name>
same provider
Use a separate provider for verification
--verify-model <id>
provider default
Use a separate verification model
--verify-failure <mode>
abort
abort , keep , or drop when verification fails
--output <format>
text
json or sarif for machine-readable results
--fail-on-incomplete
off
Fail when filtering, missing patches, or the token budget leaves part of the change unreviewed
--check-run
off
Publish findings as GitHub Check Run annotations
--no-dedupe
—
Repost findings already commented by a previous pr-sage review (dedup is on by default)
--no-incremental
—
Always review the full PR diff instead of only commits since the last pr-sage review
--batch-chars <n>
80000
Max diff characters per model reques
[truncated]
On repeat runs (e.g. new commits pushed to the PR), pr-sage reviews only the commits pushed since its last review (incremental mode), skips findings it has already commented, and posts nothing when there is nothing new — no duplicate-comment spam, no wasted tokens. If your repo has a CLAUDE.md or CONTRIBUTING.md , it is automatically injected as review context (disable with "repoContext": false ). GitHub Enterprise works out of the box via $GITHUB_API_URL or the githubApiUrl config field.
Each run prints its token usage to stderr ( LLM usage: N call(s), X input / Y output tokens ) so cost stays visible.
Every summary also reports review coverage. Partial reviews never auto-approve a PR. Use
--fail-on-incomplete when incomplete coverage must fail the CI quality gate.
Point the OpenAI provider at any OpenAI-compatible server and private code never leaves your machine — no API key required:
ollama pull qwen2.5-coder:14b
OPENAI_BASE_URL=http://localhost:11434/v1 \
npx pr-sage review --repo owner/repo --pr 123 --provider openai --model qwen2.5-coder:14b
Works the same with vLLM, LM Studio, or any gateway that speaks the OpenAI chat completions API.
For GitHub Actions, init --provider self-hosted generates a workflow for a
self-hosted runner; localhost must refer to that runner, not a GitHub-hosted VM.
Put a .pr-sage.json in the directory you run from (CLI flags override it):
{
"provider" : " anthropic " ,
"locale" : " auto " ,
"exclude" : [ " src/generated/** " , " **/*.snap " ],
"paths" : [ " packages/web/** " ],
"minSeverity" : " suggestion " ,
"failOn" : " critical " ,
"context" : " full " ,
"maxTokensPerRun" : 200000 ,
"failOnIncomplete" : true ,
"skipLabels" : [ " skip-review " ],
"verify" : true ,
"verifyProvider" : " gemini " ,
"verifyModel" : " gemini-flash-latest " ,
"verifyFailure" : " abort " ,
"checkRun" : true ,
"pathRules" : [
{
"paths" : [ " packages/api/** " ],
"instructions" : " Check public API backward compatibility. " ,
"minSeverity" : " suggestion " ,
"failOn" : " warning "
}
],
"instructions" : " We use Result<T, E> for error handling — flag thrown exceptions in domain code. Prefer early returns over nested conditionals. "
}
instructions is injected into the review prompt — use it for team conventions the reviewer should enforce.
# .github/workflows/pr-sage.yml
name : AI Review
on :
pull_request :
types : [opened, synchronize]
permissions :
contents : read
pull-requests : write
checks : write
concurrency :
group : pr-sage-${{ github.event.pull_request.number }}
cancel-in-progress : true
jobs :
review :
runs-on : ubuntu-latest
steps :
# Read .pr-sage.json from trusted base code, never PR-controlled code.
- uses : actions/checkout@v4
with :
ref : ${{ github.event.pull_request.base.sha }}
persist-credentials : false
- uses : Kyeom1997/pr-sage@v1
with :
provider : anthropic
anthropic-api-key : ${{ secrets.ANTHROPIC_API_KEY }}
locale : Korean
fail-on : critical # optional: block merge on critical findings
fail-on-incomplete : " true "
check-run : " true "
To upload SARIF, add security-events: write and set sarif: "true" . The Action
also exposes first-class paths , max-tokens , verify-provider ,
verify-model , verify-failure , and openai-base-url inputs.
The base-SHA checkout is deliberate: configuration and repository guidelines
must come from trusted base code. The PR diff itself is always fetched through
the GitHub API, and pr-sage rechecks the head SHA immediately before posting so
a slow review cannot comment on a superseded commit.
scripts/bench.mjs runs pr-sage (dry, nothing posted) over recent merged PRs of any public repo and records findings, severity mix, latency, and token usage, plus a labeling sheet for computing the valid-review rate:
node scripts/bench.mjs --repos fastify/fastify --per-repo 5 --provider gemini
--mode recall measures the other axis: it picks merged PRs that received human review comments, reviews each PR's first commit (the state humans reviewed), and reports how many human-flagged locations pr-sage also flags — with a side-by-side sheet for manual verification.
The Quality Benchmark workflow can run either benchmark manually and uploads
the generated JSON and labeling sheet as workflow artifacts.
28 recently merged PRs across fastify, hono, GitHub CLI, and Vite ( gemini-flash-lite , zero run failures):
25/28 (89%) produced zero comments — quiet on code that had already passed human review. That silence is the point: no noise on clean diffs.
The other 3 PRs got 7 findings. Verifying each claim against the actual diff: 3/7 valid overall, 2/3 for warning severity — the noise concentrated in the suggestion tier.
Re-running those PRs with --verify kept exactly the 2 diff-confirmed-valid findings (a real GPG-signing regression question in a deployment workflow) and rejected every invalid one.
Median 2.3 s and ~3.5 k input tokens per PR (≈ $0.01 for all 28 PRs at flash-lite list pricing).
Raw results and the per-finding verification notes live in bench-results/ . Caveats: small sample, and merged-PR sampling measures noise , not recall — a detection benchmark (reviewing pre-review commits of PRs that later got human fixes) is future work.
Fetches the PR metadata and per-file patches from the GitHub API.
Annotates both sides of the diff — added/context lines with new-file numbers, deleted lines with old-file numbers — so findings can anchor to removed code too (e.g. "this deleted validation was still needed"). Lockfiles/build artifacts are filtered out.
Asks the LLM for a structured review (JSON schema — no parsing heuristics): summary + findings with path , line , severity , and an optional single-line suggestion.
Validat

[truncated]
