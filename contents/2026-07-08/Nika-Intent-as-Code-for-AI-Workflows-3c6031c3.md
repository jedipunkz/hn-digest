---
source: "https://github.com/supernovae-st/nika"
hn_url: "https://news.ycombinator.com/item?id=48837692"
title: "Nika – Intent as Code for AI Workflows"
article_title: "GitHub - supernovae-st/nika: Intent as Code | the workflow language for AI. One file, 4 verbs, one Rust binary. Local-first, any model, AGPL-3.0. 🦋 · GitHub"
author: "ThibautMelen"
captured_at: "2026-07-08T22:01:22Z"
capture_tool: "hn-digest"
hn_id: 48837692
score: 1
comments: 0
posted_at: "2026-07-08T21:30:00Z"
tags:
  - hacker-news
  - translated
---

# Nika – Intent as Code for AI Workflows

- HN: [48837692](https://news.ycombinator.com/item?id=48837692)
- Source: [github.com](https://github.com/supernovae-st/nika)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T21:30:00Z

## Translation

タイトル: Nika – AI ワークフローのコードとしての意図
記事のタイトル: GitHub - supernovae-st/nika: コードとしての意図 | AI のワークフロー言語。 1 つのファイル、4 つの動詞、1 つの Rust バイナリ。ローカルファースト、あらゆるモデル、AGPL-3.0。 🦋 · GitHub
説明: コードとしての意図 | AI のワークフロー言語。 1 つのファイル、4 つの動詞、1 つの Rust バイナリ。ローカルファースト、あらゆるモデル、AGPL-3.0。 🦋 - 超新星-st/ニカ

記事本文:
GitHub - supernovae-st/nika: コードとしての意図 | AI のワークフロー言語。 1 つのファイル、4 つの動詞、1 つの Rust バイナリ。ローカルファースト、あらゆるモデル、AGPL-3.0。 🦋 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
超新星セント

/
ニカ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,257 コミット 1,257 コミット .agents/ plugins .agents/ plugins .cargo .cargo .claude-plugin .claude-plugin .claude .claude .github .github crates crates docs docs 例 例 ファズ ファズ メディア メディア スクリプト スクリプト .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CITATIONS.md CITATIONS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DIAMOND.md DIAMOND.md ライセンス ライセンス README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md crime.toml crime.toml Clippy.toml Clippy.toml Deny.toml Deny.toml lefthook.yml lefthook.yml llms.txt llms.txt Rust-toolchain.toml Rust-toolchain.toml typos.toml typos.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コードとしての意図。 AI のワークフロー言語: 1 つのファイル、4 つの動詞、
1 つのバイナリ。
有用な AI の仕事がチャットに消えてしまうべきではありません。ニカが反復可能なAIを変える
実行、レビュー、比較、共有できるファイルに変換します。あなたも同じことをしたら
AIタスクを2回行い、ワークフロー化します。
Nika ワークフローは、読み取り可能、移植可能、検証可能な単なるファイルです。走ります
選択した LLM 上でローカルに実行でき、クラウドは必要ありません。言語
これはオープンな Apache-2.0 仕様です。
このリポジトリはリファレンス エンジン、単一の Rust バイナリ (AGPL-3.0) です。の
SQL と PostgreSQL を組み合わせたり、Dockerfile と Docker を組み合わせたりする方法です。
brew install supernovae-st/tap/nika # または:curl -LsSf https://nika.sh/install.sh |しー
nika の例で実行 01-hello --model mock/echo # ゼロ セットアップ: キーなし、モデル サーバーなし
nika の例は 01-hello --model ollama/qwen3.5:4b を実行します # Ollama を取得しましたか?同じ実行、実数 + ローカル
ニカが監査する

単一のトークンが使用される前のワークフロー (プラン、コスト)
天井、シークレット フロー、タイプ、ツール引数) を指定して実行します。
$ nika チェック Brief.nika.yaml
✔ 計画 2 ウェーブ · 2 タスク · 最大並列処理 1
✔ 秘密 情報の流れが逃げない
✔ TYPES すべての深い出力参照は、宣言された形状に適合します
✔ ツールごとの nika: ツールは正規の組み込みに名前を付けます
✔ ARGS すべての呼び出し引数キーが宣言されている + 必要なすべての引数が存在する
✔ SCHEMA すべての作成されたスキーマ: 満足可能です
✔ クリーン — 単一のトークンが消費される前に監査される
$ nika run Brief.nika.yaml
🦋 ニカ · デイリーブリーフ · 2 つのタスク
✔ fetch_notes 実行者 · 猫
✔ 簡単な推論 · オラマ/ラマ3.2:3b
── 2/2 完了 · $0.000 · 経過 16.2 秒 ──────────
ワークフローはどのようなものですか
# review.nika.yaml: PR の差分を読み、そのリスクを判断し、リスクが高い場合にのみコメントします。
ニカ：v1
ワークフロー: 事前-リスク-レビュー
モデル : ollam/qwen3.5:9b # デフォルトではローカル。任意のプロバイダーに切り替える
タスク:
- id : diff # exec: 読み取り専用シェルコマンド
実行:
コマンド: " git diff 起点/メイン...HEAD "
キャプチャ: 構造化された
- id :assess #infer:構造化LLM判定
: { パッチ: ${{ task.diff.output.stdout }} }
推論：
プロンプト : " この差分をリスク評価します (秘密、重大な変更、不足しているテスト)。簡潔にお願いします。\n ${{ with.patch }} "
スキーマ:
タイプ: オブジェクト
必須: [リスク]
プロパティ:
リスク: { タイプ: 文字列、列挙: [低、中、高] }
- id : コメント # invoke: 唯一の書き込み、判定に基づいて制御される
とき: ${{ task.assess.output.risk == '高' }}
呼び出します:
ツール: " mcp:github/pr-comment "
引数: { 本体: ${{ タスク.評価.出力 }} }
実行前に確認してください
nika check は静的な監査です。壊れた参照、欠落している参照をキャッチします
モデルが呼び出される前の依存関係、スキーマ、および権限の問題
—そして、何

何かが間違っている場合、正確な修正を示しています。
このキャプチャの背後にある 2 つのフィクスチャは、
scripts/media/fixtures/ 、両方でゲートされます
scripts/media/validate-media.sh による指示: 壊れたものは保持する必要があります
nika チェックに失敗した場合、修正されたものはクリーンなままでなければなりません。
同じ監査には、ワークフローの宣言された爆発半径が保持されます。 A は次のことを許可します。
ブロックは、ファイル自体をセキュリティ境界 (ホスト、パス、プログラム、
ツールは、一度宣言されるとすべてデフォルトで拒否されます。それを超えた課題は
何かが実行される前に、マシンに適用可能な修正を加えて静的にキャッチされます。
また、障害処理はファイルの一部であり、運用ランブックではありません。タスクのとき
死亡、on_error: 回復: 宣言されたフォールバックに低下します — 実行
完了すると、出力に内容が表示されます。
nika:image_generate はすべてと同じ規律を通じてレンダリングします。
else — 宣言された許可: 保存ごとに境界ゲート、実行台帳
実際の支出はメートルであり、出所は構造的なものであり、約束ではありません。
ローカルファースト — プロバイダー: ローカルは OpenAI を話します - 画像ワイヤーは任意です
自己ホスト型サーバーが公開する (LocalAI · Ollama ·steady-diffusion.cpp ·
SGLang・vLLM-Omni）。ベース URL はエンジン構成であり、ワークフロー データではありません。
選択したクラウド: openai、gemini、xai — およびモック レンダリング
CI 用にオフラインでデコード可能な実際の PNG ファイル。
来歴は cp を存続します — 保存されたすべての PNG には nika tEXt が含まれます
チャンク (ツール、エンジン、プロバイダー、モデル、プロンプト、シード)、実践
ComfyUI と InvokeAI は標準化されており、他のワークフロー エンジンは付属していません。
サイドカー マニフェストには、sha256、解決されたリクエストとタイミングが追加されます。
正直さが強制されます - マジックバイトは宣言された MIME タイプに打ち勝ち、不可逆性を伴います
プロバイダー マッピングは、プロバイダーが返す画像の数が少ない場合に大声で警告します。
表示される count_shortfall: が尋ねられ、結果の URL は決して表示されません
フェッチされ、base64 はワークフロー出力 (アセット、

ブロブではありません。
台帳内の実際の支出 — レンダリングの正確なコスト (xAI の請求書は
ティック) がタスク ラインと実行合計に表示され、同じ正直な支出が発生します
チャネル推論: を使用します。
nika Inspection flow.nika.yaml # 解剖学、タスク、ウェーブ、コストフロア
nika check flow.nika.yaml # 監査 · 終了 0 クリーン · 2 所見
nika Explain NIKA-VAR-001 # 任意のコード · 原因 · カテゴリ · 修正形式
nika run flow.nika.yaml --var topic=rust # 起動入力 · 繰り返し可能
nika test flow.nika.yaml --update # ゴールデンを固定します · その後、`nika test` = オフライン CI
nika run flow.nika.yaml --task hero # 1 つのタスクとその上流のコーンを再生成します
nika run flow.nika.yaml --resume .nika/traces/ < run > .ndjson # ジャーナリングされた成功をスキップします
nika run flow.nika.yaml --resume <trace> --answerapprove=true # 一時停止したゲートを再装備します
nika trace show .nika/traces/ < run > .ndjson # 過去の実行を再レンダリングします
nika Doctor --ping # ローカルサーバーは実際にリッスンしていますか?
実行ごとに独自のジャーナルが .nika/traces/ に書き込まれます (実行ごとにオプトアウトするには
--no-trace-file 、NIKA_NO_TRACE_FILE でグローバル) — --resume および
ニカトレースはそれを直接読みます。一時停止された実行は 4 から終了します (ブロッキング
ニカ:プロンプトジャーナルの質問);再開時のキャッシュ ヒットは常に
表示 — 何も通知せずにスキップされることはありません。
バイナリには、実行可能なサンプルのバージョン管理されたパックが埋め込まれています。で閲覧する
nika の例のリスト、nika の例で 1 つを読む、 <slug> を表示、任意のプレビュー
--model ollam/qwen3.5:4b を使用してそれらを実行します (または --model Mouck/echo を使用してオフラインで使用します)。
完全なギャラリー (27 ワークフロー + 6 テンプレート) は次の場所にあります。
例/ : 基本パターン、ビジネス ショーケース、
スケルトン nika new --from <template> がインスタンス化されます。
共有ワークフローは nika-registry 上に存在します —
すべてのエントリはフルコミット + sha256 に固定され、CI によって再証明されます (
適合オラクル + このエンジンの静的証明書: exec · to

オルス ·
コスト、何かが実行される前に表示されます）。 nika add はロードマップにあります。今日
レジストリの get.py は、1 つのコマンドで解決→検証→監査を実行します。
動詞が 4 つあり、他には何もありません。任意に構成する小さなコア
現実世界のワークフロー。 Unix と SQL の規律「小さな表面、大きな表面」
構成。」
すべては、凍結されたバージョン管理された 1 つのエンベロープ nika: v1 の下に置かれます。
休憩する。すべてのワークフローにわたって 3 つのプロパティが保持されます。
プロバイダーに依存しない、ローカルファースト。ローカルの Ollama または LM Studio、または任意の API。
モデルが変化してもワークフローは変わりません。
安全な構造です。読み取り-XOR-書き込み機能モデル。という一歩
読み取りはサイレントに書き込みできません。すべてのエフェクトは明示的でゲートされています。
再現可能。ファイルとその実行トレースは監査可能です。
再実行可能なレコード。
フローチャート LR
F["workflow.nika.yaml<br/><i>ポータブル · 読み取り可能 · 検証可能</i>"] --> E["<b>ニカ</b><br/>単一のRustバイナリ"]
E -->|推測| L["LLM<br/>Ollama · LM Studio · 任意の API"]
E -->|実行| S["シェル"]
E -->|呼び出し| T[「ツール・MCP」]
E -->|エージェント| A[「自律ループ」]
読み込み中
依存関係により、すべてのワークフローがグラフ化されます。独立したタスクが実行されます。
並行して、エージェント ステップがファンアウトし、分岐ごとに結合が待機します。
計画全体が既知であり、実行が開始される前にコストが計算され、監査されます。
最も近い類似物は製品ではありません。それらは標準です。 SQL。の
ドッカーファイル。リファレンスエンジンを搭載したポータブル仕様。言語は
販売する製品ではなく、貢献です。
AI エージェントが現実世界に対して行動し始めると、AI エージェントが行動するインターフェースが
フリーテキスト (曖昧すぎる) や生のコード (リスクが高すぎる) は使用できません。それは、
検証可能なアクション言語: AI が作成し、人間がレビューして承認します。
そしてマシンは決定的に動作します。オープンかつ主権を保ち、ロックされない
あるベンダーのクラウド内で。
既存のワークフローがないもの

単一の Rust バイナリとポータブルなツールを組み合わせて提供します
宣言型 YAML · ローカルファースト · 読み取り-XOR-書き込み機能のセキュリティ · AGPL ·
クラウドは必要ありません。LLM を持ち込みます。
言語 (nika: v1 エンベロープとその 4 つの動詞) は安定しており、
壊れません。このエンジンは厳密なモジュール式の Rust ワークスペースです。最新の
タグ付きの公開リリースは v0.91.0 です。 main はすぐに次の項目に移動します
-dev のバージョンはリリースごとに変更されるため、ローカルのコントリビューター バイナリは使用できません。
Homebrew アセットと混同されます。 1.0.0 のリリースはリリースによって制限されたままです
日付ではなくチェックリスト。コード、
スペック、そして
ワークフローの例はすべて読みやすく、開発は次の場所で行われます。
オープンでのメイン。
nika: v1 言語エンベロープは永久に凍結されます。とは別軸です。
エンジンバージョン。すべてのリリースは、宣言されたスコープに対して完了しています。いいえ
半分の機能は将来のバージョンに保留されます。
# Homebrew (macOS · Linux): すぐに PATH に追加
brew install supernovae-st/tap/nika
ニカ --バージョン
# …または、Homebrew を使用しない場合: インストール スクリプト。検証済みリリースをダウンロードします
# binary を ~/.nika/bin に格納し、シェルに追加する単一の PATH 行を出力します
# profile (その後、ターミナルを再度開くか、「source」すると、「nika --version」が機能します)。
カール -LsSf https://nika.sh/install.sh |しー
ガイド付きページをご希望ですか?各インストール パスのステップバイステップ: nika.sh/install 。
完全手動/エアギャップ?ダウンロード

[切り捨てられた]

## Original Extract

Intent as Code | the workflow language for AI. One file, 4 verbs, one Rust binary. Local-first, any model, AGPL-3.0. 🦋 - supernovae-st/nika

GitHub - supernovae-st/nika: Intent as Code | the workflow language for AI. One file, 4 verbs, one Rust binary. Local-first, any model, AGPL-3.0. 🦋 · GitHub
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
supernovae-st
/
nika
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,257 Commits 1,257 Commits .agents/ plugins .agents/ plugins .cargo .cargo .claude-plugin .claude-plugin .claude .claude .github .github crates crates docs docs examples examples fuzz fuzz media media scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CITATIONS.md CITATIONS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DIAMOND.md DIAMOND.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md cliff.toml cliff.toml clippy.toml clippy.toml deny.toml deny.toml lefthook.yml lefthook.yml llms.txt llms.txt rust-toolchain.toml rust-toolchain.toml typos.toml typos.toml View all files Repository files navigation
Intent as Code. The workflow language for AI: one file, 4 verbs,
one binary.
Useful AI work shouldn't disappear into chats. Nika turns repeatable AI
work into files you can run, review, diff and share. If you do the same
AI task twice, make it a workflow.
A Nika workflow is just a file: readable, portable, verifiable. It runs
locally, on whichever LLM you choose, with no cloud required. The language
is an open Apache-2.0 spec ;
this repo is the reference engine, a single Rust binary (AGPL-3.0). The
way SQL pairs with PostgreSQL, or the Dockerfile with Docker.
brew install supernovae-st/tap/nika # or: curl -LsSf https://nika.sh/install.sh | sh
nika examples run 01-hello --model mock/echo # zero setup: no key, no model server
nika examples run 01-hello --model ollama/qwen3.5:4b # got Ollama? the same run, real + local
Nika audits a workflow before a single token is spent (plan, cost
ceiling, secret flows, types, tool args), then runs it:
$ nika check brief.nika.yaml
✔ PLAN 2 wave(s) · 2 task(s) · max parallelism 1
✔ SECRETS no information-flow escapes
✔ TYPES every deep output reference fits its declared shape
✔ TOOLS every nika: tool names a canonical builtin
✔ ARGS every invoke arg key is declared + every required arg is present
✔ SCHEMA every authored schema: is satisfiable
✔ clean — audited before a single token was spent
$ nika run brief.nika.yaml
🦋 nika · daily-brief · 2 tasks
✔ fetch_notes exec · cat
✔ brief infer · ollama/llama3.2:3b
── 2/2 done · $0.000 · elapsed 16.2s ───────────────────────────
What a workflow looks like
# review.nika.yaml: read a PR diff, judge its risk, comment only when it's high.
nika : v1
workflow : pr-risk-review
model : ollama/qwen3.5:9b # local by default. swap to any provider
tasks :
- id : diff # exec: a read-only shell command
exec :
command : " git diff origin/main...HEAD "
capture : structured
- id : assess # infer: structured LLM judgment
with : { patch: ${{ tasks.diff.output.stdout }} }
infer :
prompt : " Risk-assess this diff (secrets, breaking changes, missing tests). Be terse. \n ${{ with.patch }} "
schema :
type : object
required : [risk]
properties :
risk : { type: string, enum: [low, medium, high] }
- id : comment # invoke: the only write, gated on the verdict
when : ${{ tasks.assess.output.risk == 'high' }}
invoke :
tool : " mcp:github/pr-comment "
args : { body: ${{ tasks.assess.output }} }
Check before it runs
nika check is a static audit. It catches broken references, missing
dependencies, schema and permission problems before any model is called
— and when something is off, it points at the exact fix:
The two fixtures behind this capture live in
scripts/media/fixtures/ , gated in both
directions by scripts/media/validate-media.sh : the broken one must keep
failing nika check , the fixed one must stay clean.
The same audit holds the workflow's declared blast radius . A permits:
block makes the file itself the security boundary — hosts, paths, programs,
tools, all default-deny once declared. A task that reaches beyond it is
caught statically, with the machine-applicable fix, before anything runs:
And failure handling is part of the file, not an ops runbook. When a task
dies, on_error: recover: degrades to a declared fallback — the run
completes, and the output says what it is:
nika:image_generate renders through the same discipline as everything
else — a declared permits: boundary gates every save, the run ledger
meters real spend, and provenance is structural, not a promise:
Local-first — provider: local speaks the OpenAI-images wire any
self-hosted server exposes (LocalAI · Ollama · stable-diffusion.cpp ·
SGLang · vLLM-Omni). The base URL is engine config, never workflow data.
Clouds when you choose: openai · gemini · xai — and mock renders
real, decodable PNG files offline for CI.
Provenance survives cp — every saved PNG carries a nika tEXt
chunk (tool · engine · provider · model · prompt · seed), the practice
ComfyUI and InvokeAI standardized and no other workflow engine ships.
The sidecar manifest adds sha256, resolved request and timing.
Honesty is enforced — magic bytes beat declared MIME types, lossy
provider mappings warn loudly, a provider returning fewer images than
asked is a visible count_shortfall: , result URLs are never
fetched, and base64 never rides workflow outputs — assets, not blobs.
Real spend in the ledger — a render's exact cost (xAI bills in
ticks) lands on the task line and the run total, the same honest-spend
channel infer: uses.
nika inspect flow.nika.yaml # anatomy · tasks · waves · cost floor
nika check flow.nika.yaml # the audit · exit 0 clean · 2 findings
nika explain NIKA-VAR-001 # any code · cause · category · fix-form
nika run flow.nika.yaml --var topic=rust # launch inputs · repeatable
nika test flow.nika.yaml --update # pin the golden · then `nika test` = offline CI
nika run flow.nika.yaml --task hero # regenerate ONE task + its upstream cone
nika run flow.nika.yaml --resume .nika/traces/ < run > .ndjson # skip journaled successes
nika run flow.nika.yaml --resume < trace > --answer approve=true # re-arm a paused gate
nika trace show .nika/traces/ < run > .ndjson # re-render any past run
nika doctor --ping # are the local servers actually listening?
Every run writes its own journal to .nika/traces/ (opt out per run with
--no-trace-file , globally with NIKA_NO_TRACE_FILE ) — --resume and
nika trace read it directly. A paused run exits 4 (a blocking
nika:prompt journals its question); cache hits on resume are always
visible — nothing is skipped silently.
The binary embeds a versioned pack of runnable examples. Browse with
nika examples list , read one with nika examples show <slug> , preview any
of them with --model ollama/qwen3.5:4b (or offline with --model mock/echo ):
The full gallery (27 workflows + 6 templates) lives in
examples/ : foundation patterns, business showcases, and the
skeletons nika new --from <template> instantiates.
Shared workflows live on nika-registry —
every entry pinned to a full commit + sha256 and re-proven by CI (the
conformance oracle + this engine's static certificate: exec · tools ·
cost, visible before anything runs). nika add is on the roadmap; today
the registry's get.py does resolve → verify → audit in one command.
Four verbs, and nothing else. A small core that composes into arbitrary
real-world workflows. The Unix and SQL discipline of "small surface, large
composition."
Everything sits under one frozen, versioned envelope, nika: v1 , that won't
break. Three properties hold across every workflow:
Provider-agnostic, local-first. Local Ollama or LM Studio, or any API.
Your workflow doesn't change when the model does.
Safe by construction. A read-XOR-write capability model. A step that
reads cannot silently write; every effect is explicit and gated.
Reproducible. The file and its execution trace are an auditable,
re-runnable record.
flowchart LR
F["workflow.nika.yaml<br/><i>portable · readable · verifiable</i>"] --> E["<b>nika</b><br/>single Rust binary"]
E -->|infer| L["LLMs<br/>Ollama · LM Studio · any API"]
E -->|exec| S["shell"]
E -->|invoke| T["tools · MCP"]
E -->|agent| A["autonomous loop"]
Loading
Dependencies make every workflow a graph: independent tasks run in
parallel, an agent step fans out, joins wait for every branch — and the
whole plan is known, costed and audited before execution starts:
The closest analogues aren't products. They're standards . SQL. The
Dockerfile. A portable specification with a reference engine. The language is
the contribution, not a product to sell.
As AI agents start acting on the real world, the interface where they act
can't be free text (too vague) or raw code (too risky). It has to be a
verifiable action language : one an AI writes, a human reviews and approves,
and a machine runs deterministically. Kept open and sovereign, not locked
inside one vendor's cloud.
What no existing workflow tool offers together: a single Rust binary · portable
declarative YAML · local-first · read-XOR-write capability security · AGPL ·
no cloud required · bring-your-own-LLM.
The language (the nika: v1 envelope and its four verbs) is stable and
won't break. The engine is a strict, modular Rust workspace. The latest
tagged public release is v0.91.0 ; main moves immediately to the next
-dev version after each release so local contributor binaries cannot be
confused with Homebrew assets. The 1.0.0 launch remains gated by the release
checklist, not by a date. The code, the
spec , and the
example workflows are all readable, and development happens on
main in the open.
The nika: v1 language envelope is frozen forever. It is a separate axis from the
engine version. Every release is complete for its declared scope; no
half-features parked behind a future version.
# Homebrew (macOS · Linux): on your PATH immediately
brew install supernovae-st/tap/nika
nika --version
# …or, without Homebrew: the install script. It downloads the verified release
# binary into ~/.nika/bin and prints the single PATH line to add to your shell
# profile (then reopen the terminal, or `source` it, and `nika --version` works).
curl -LsSf https://nika.sh/install.sh | sh
Prefer a guided page? Every install path, step by step: nika.sh/install .
Fully manual / air-gapped? Download the

[truncated]
