---
source: "https://github.com/Tura-AI/tura"
hn_url: "https://news.ycombinator.com/item?id=48922039"
title: "Agent runtime reduces LLM turns by 80% with a higher success rate in DeepSWE"
article_title: "GitHub - Tura-AI/tura: Across 348 long-horizon benchmark sessions, Tura used up to 83.1% fewer turns on the rewrite benchmark and improved the DeepSWE pass rate by up to 16.7 percentage points compared with Codex CLI. · GitHub"
author: "yohji1984"
captured_at: "2026-07-15T15:18:40Z"
capture_tool: "hn-digest"
hn_id: 48922039
score: 1
comments: 1
posted_at: "2026-07-15T15:10:14Z"
tags:
  - hacker-news
  - translated
---

# Agent runtime reduces LLM turns by 80% with a higher success rate in DeepSWE

- HN: [48922039](https://news.ycombinator.com/item?id=48922039)
- Source: [github.com](https://github.com/Tura-AI/tura)
- Score: 1
- Comments: 1
- Posted: 2026-07-15T15:10:14Z

## Translation

タイトル: エージェントのランタイムにより、DeepSWE の成功率が向上し、LLM ターンが 80% 削減されました
記事のタイトル: GitHub - Tura-AI/tura: 348 のロングホライズン ベンチマーク セッション全体で、Tura はリライト ベンチマークで使用するターン数が最大 83.1% 少なく、Codex CLI と比較して DeepSWE の合格率が最大 16.7 パーセント ポイント向上しました。 · GitHub
説明: 348 のロングホライズン ベンチマーク セッション全体で、Tura は、Codex CLI と比較して、書き換えベンチマークで使用するターン数を最大 83.1% 削減し、DeepSWE の合格率を最大 16.7 パーセント向上させました。 - Tura-AI/トゥーラ

記事本文:
GitHub - Tura-AI/tura: 348 のロングホライズン ベンチマーク セッション全体で、Tura は書き換えベンチマークで使用するターン数が最大 83.1% 少なく、Codex CLI と比較して DeepSWE の合格率が最大 16.7 パーセント ポイント向上しました。 · GitHub
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
そこには

読み込み中にエラーが発生しました。このページをリロードしてください。
トゥーラAI
/
トゥーラ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
610 コミット 610 コミット .github .github エージェント エージェント アプリ アプリ アセット アセット コマンド コマンド クレート クレート ドキュメント ドキュメント npm ペルソナ ペルソナ スクリプト スクリプト テスト テスト xtask xtask .gitattributes .gitattributes .gitignore .gitignore .npmignore .npmignore ARCHITECTURE.md ARCHITECTURE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md ROADMAP.md ROADMAP.md package.json package.json要件.txt要件.txtrust-toolchain.tomlrust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
トゥーラ: ターン数が 83.1% 減りました。成功率は 16.7 パーセント ポイント高かった。
Tura は、曖昧なスキルの主張、証拠のないトークン節約拡張機能、判断力のないエージェントがリポジトリを破壊することにうんざりしている開発者のための、ローカルのオープンソース コーディング エージェントです。
Tura は、20 の DeepSWE v1.1 タスクにわたって、GPT-5.6 SOL による高い推論労力で 60 セッションをテストし、繰り返されるコンテキストとモデルの往復を削減することで、トークン予算の大幅な利点を生み出しました。このメリットを 2 つの方法で活用できます。ダイレクトでは、そのほとんどが低コストに変換されます。公式の Codex CLI High 構成よりも総トークンが 83.5% 少なく、検証成功率は 65.0% 対 60.0% です。 Balanced では、節約された予算の多くを推論、調査、検証に戻します。成功率は 80.0% に達し、Codex CLI High より 20 パーセントポイント高く、それでも使用するトークンは 49.6% 少なくなりました。 1 2
長期的なタスクのベンチマークは、洗練された分離されたプロンプトを超えて、エージェントが実際の作業をどのように処理するかを確認する 1 つの方法です。公開された

比較では、アーカイブされたプロンプト、ラウンドごとのツール呼び出し、トークンの使用法、パッチ、検証結果を含むハーネスベースの開発タスクを使用します。
以下の主な比較では、モデルと推論ラベルが固定されています: Tura Balanced High、Tura Direct High、および 20 個の DeepSWE タスクと 5 個の書き換えタスクに対する公式 Codex CLI High 構成。証拠レコードには、Codex CLI Medium も別のセカンダリ構成として保持されます。ベンチマーク手法では、個別にレビューされた 2 つの設計タスクがハーネスのスコア付け対象外に保たれます。 1
公開された結果は、同等の品質やパフォーマンスを確立するものではありません。
構成されているすべてのプロバイダー。より広範な人間/クロード、Google/ジェミニ、
OpenAI 互換、ローカルプロバイダー、UI レイテンシ、ランタイム/セッション解析、および
クロス OS 測定は文書化されたものの一部として残ります
ロードマップと既知の証拠のギャップ。
マルチセッションの同時作業と HTML リッチ テキストのサポートを備えた GUI ページ。
マルチセッションの同時作業と HTML リッチ テキストのサポートを備えた TUI ページ。
以下の結果は、引用されていない集計ではなく、公開されているベンチマーク成果物から得られたものです。ほとんどの作業は 3 つのシステムで行われます。
ほとんどのコーディング エージェントは依然として、検査、待機、パッチ、待機、ビルド、待機、テスト、待機という繰り返しのツール呼び出しループに依存しています。
# ターン 1 — 環境を検査する
rg -n " TODO|command_run|handler " crates/
rg --files クレート/ランタイム/src クレート/ツール/src
# ターン 2 — パッチを適用します
*** パッチの開始
*** 更新ファイル: crates/tools/src/command_run/handler.rs
@@
- // 古いコマンド ハンドラー ロジック
+ // パッチされたコマンド ハンドラー ロジック
*** エンドパッチ
# ターン 3 — 構築
カーゴビルド -p ランタイム
# ターン 4 — テストの実行
カーゴテスト -p runtime --lib
# ターン 5 — lint 検証を実行する
カーゴクリッピー -p runtime --all-targets
トゥーラは別の道を歩みます。モデルに多数の小さなツールを公開する代わりに、1 つのマクロ ツールを公開します。

: command_run 。その後、エージェントは複数ステップの実行ツリーを構築し、LLM の 1 ターンで関連アクションを実行できます。
以下の例では、両方のエージェントが同じコマンドを実行します。通常のツール呼び出しエージェントには 5 つの LLM ターンが必要です。 Tura はシーケンスを 1 つの構造化されたマクロ ワークフローとして処理します。保存される作業は会話のオーバーヘッドであり、エンジニアリングの規律ではありません。
{
"名前" : " command_run " ,
"引数" : {
「コマンド」: [
{
「ステップ」 : 1 、
"command_type" : " シェルコマンド " ,
"command_line" : " rg -n \" TODO|command_run|handler \" crates/ "
}、
{
「ステップ」 : 1 、
"command_type" : " シェルコマンド " ,
"command_line" : " rg --files crates/runtime/src crates/tools/src "
}、
{
「ステップ」 : 2 、
"command_type" : " apply_patch " ,
"command_line" : " *** パッチ開始 \n *** ファイル更新: crates/tools/src/command_run/handler.rs \n @@ \n - // 古いコマンド ハンドラー ロジック \n + // パッチされたコマンド ハンドラー ロジック \n *** パッチ終了 "
}、
{
「ステップ」 : 3 、
"command_type" : " シェルコマンド " ,
"command_line" : " カーゴビルド -p ランタイム "
}、
{
「ステップ」 : 4 、
"command_type" : " シェルコマンド " ,
"command_line" : " カーゴテスト -p runtime --lib "
}、
{
「ステップ」 : 4 、
"command_type" : " シェルコマンド " ,
"command_line" : " Cargo Clippy -p runtime --all-targets "
}
】
}
}
command_run だけが Tura のターンとトークンの使用量を低下させることを証明するアブレーション テストはありません。ただし、Matched-High DeepSWE の比較では、Balanced は Codex CLI High よりも使用するモデル ラウンドが 66.8% 少なく、トークンが 49.6% 少ないのに対し、Direct は使用するラウンドが 84.0% 少なく、トークンが 83.5% 少ないです。 1 2
LLM がどれほど優れたものであっても、LLM の核心はテキスト トークンの確率に対する統計的帰納モデルであることに変わりはありません。
たとえば、LLM にジャンケンから選択するように依頼しても、均一でランダムな結果が得られるとは限りません。本当の 3 分の 1 の配分が重要である場合、

選択には、モデルの出力確率に関する引用されていない仮定ではなく、外部の乱数ソースが必要です。
コーディング作業では、これは致命的なことがよくあります。
エージェントは、統計的により一般的なコードとロジックを実行および生成する可能性が高くなります。しかし、共通のコードや共通のロジックは、多くの場合平凡で、十分に考慮されていません。
トゥーラは別の戦略を使用します。
推論中、共通エージェントは現在の状態からプロンプトの目標まで推論します。この場合、$s_1$ は現在の状態であり、$s_n$ はユーザー プロンプトによって指定された目標です。
$$
s_1 \rightarrow s_2 \rightarrow s_3 \rightarrow \cdots \rightarrow s_n
$$
代わりに、Tura は LLM が最初に $s_{n-1}$ を統計的に推定し、次に $s_{n-1}$ の状態から $s_{n-2}$ まで逆算するようにガイドします。
以下の例では、LLM はじゃんけんを正しくプレイするための最適な戦略を導き出すことができます。
> じゃんけんを公平かつ挑戦的に保つために、
> 私たちには公平なプレーが必要です。
> 各動きには実際に 3 分の 1 の確率がなければなりません。
> LLM はテキストの確率だけからそれを保証することはできません。
> 乱数生成スクリプトを使用して randint(1, 3) を生成する
> 次に、じゃんけん、紙、またははさみをその数字にマッピングします。
これは、プログラミング タスクにおいて、エージェントがフロントエンドのバグの修正などの目標を認識したときに、コードを記述する前に完全な実行パスを通じて推論し、障害状態を再構築し、根本原因を特定するように誘導されることを意味します。 Matched-High DeepSWE の比較では、Tura Balanced は Codex CLI High よりも 60 個のバイナリ タスク ベリファイアーのうち 12 個多く合格しました。
この対照的な両方の構成は、「高」推論ラベルが付いた GPT-5.6 SOL を使用しているため、「高」と「中」の努力の不一致は、合格率の 20 ポイントの差を説明しません。結果は依然としてシステムレベルの関連性であり、後方推論やその他の個別の特徴に対する因果関係の推定ではありません。

。 1 2
ランタイムコンテキストとプロンプトマネージャー
スキルは多くの場合、コンテキストに読み込まれた弱いプロンプトに過ぎません。
多くのエージェント フレームワークでは、存続期間の長いセッションにより、スキル ファイル、ツール出力、および古いタスク履歴が蓄積され続けます。コンテキストが大きくなりすぎると、エージェントは別の圧縮ターンに入りますが、その圧縮では通常、圧縮された概要のみが保存されます。重要な実行の詳細が曖昧になったり、失われたりする可能性があります。
Tura はコンテキストをランタイム ステート マシンの一部として扱います。
ユーザーに手動でセッションをリセットしたり、Markdown スキルを蓄積させたりするのではなく、Tura は task_status 、ランタイム プロンプト、および再帰的実行マニュアルを使用して、アクティブ コンテキストのスコープを現在のタスクに保ちます。
従来のスキルベースのエージェントは通常、ユーザーが別のセッションを開始するまで 1 つのセッションを実行し続け、広範な Markdown スキルをそのセッションにロードし、リセットまたは圧縮されるまでアクティブなままにしておきます。代わりに、Tura は実行時プロンプトを明示的なタスク状態に結び付けます。セッションは自動的に名前変更、更新、管理できます。タスク固有のマニュアルと CLI コマンドは、再帰的なタスク ツリーを通じてロードされます。無関係なコンテキストは CLI から削除、置換、または圧縮できます。チェックポイントは、大まかな概要だけではなく、コードの場所、パッチ、テスト、タスクのステータスを保持できます。実際には、これは、古いコンテキストが少なくなり、タスク範囲のトークンコストが低くなり、古いスキルや曖昧な概要が現在の仕事の方向性を決める可能性が少なくなることを意味します。
圧縮は CLI 操作であるため、Tura は正確な実行状態を task_status.compact_context に保存できます。公開されたベンチマーク セッションでは、Tura は読み取り専用インスペクションを超えて、圧縮後に平均 2.6 ラウンドで実行を再開しました。これに対し、Codex では推定 5.4 ラウンドでした。 3 4 5 6 7
Tura の 2.6 ラウンドの結果は、明示的な Compact_conte から計算されます。

アーカイブされたラウンド契約内の xt イベント。 Codex は同等の圧縮イベントを公開していないため、その 5.4 ラウンドの結果は、識別可能なメディア読み取り境界を除いて、入力トークンの使用量が急激に低下するポイントから推定されます。
npm インストール tura-ai
トゥーラ
Windows:
npm インストール - g tura - ai
トゥーラ
同じメイン パッケージが @tura-ai/tura として GitHub パッケージにも公開されます。
https://npm.pkg.github.com の @tura-ai スコープを構成し、認証します
read:packages を持つトークンを使用して、 @tura-ai/tura をインストールします。の
npm 上のスコープ外の tura-ai パッケージは、依然として最も単純なパブリック インストールです。
Tura はプロバイダーの資格情報をバンドルしません。初回起動時にLLMを構成します
プロバイダーを選択し、プロンプトを送信する前にそのモデルの 1 つを選択します。参照
プロバイダーのセットアップ
CLI、TUI、および GUI フロー。
git clone https://github.com / Tura - AI / tura.git
CDトゥーラ
.\scripts\install.ps1
トゥーラ
macOS または Linux シェル:
git clone https://github.com/Tura-AI/tura.git
CDトゥーラ
./scripts/install.sh
トゥーラ
ソース インストーラーは、完全な環境セットアップ、リリース ビルド、
およびユーザーPATHの登録フロー。 PowerShell で -EnvironmentOnly を渡すか、
意図的に依存関係が必要な場合のみ、macOS/Linux で --environment-only
構築せずにセットアップする

[切り捨てられた]

## Original Extract

Across 348 long-horizon benchmark sessions, Tura used up to 83.1% fewer turns on the rewrite benchmark and improved the DeepSWE pass rate by up to 16.7 percentage points compared with Codex CLI. - Tura-AI/tura

GitHub - Tura-AI/tura: Across 348 long-horizon benchmark sessions, Tura used up to 83.1% fewer turns on the rewrite benchmark and improved the DeepSWE pass rate by up to 16.7 percentage points compared with Codex CLI. · GitHub
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
Tura-AI
/
tura
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
610 Commits 610 Commits .github .github agents agents apps apps assets assets commands commands crates crates docs docs npm npm personas personas scripts scripts tests tests xtask xtask .gitattributes .gitattributes .gitignore .gitignore .npmignore .npmignore ARCHITECTURE.md ARCHITECTURE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md package.json package.json requirements.txt requirements.txt rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Tura: 83.1% fewer turns; 16.7 percentage points higher success.
Tura is a local, open-source coding agent for developers who are tired of vague skill claims, token-saving extensions with no evidence, and agents without judgment wreck their repos.
Across 20 DeepSWE v1.1 tasks, tested 60 sessions with GPT-5.6 SOL at High reasoning effort, Tura creates a substantial token-budget advantage by reducing repeated context and model round trips. You can spend that advantage in two ways. Direct turns most of it into lower cost: 83.5% fewer aggregate tokens than the official Codex CLI High configuration, with a verifier success rate of 65.0% versus 60.0%. Balanced puts more of the saved budget back into reasoning, investigation, and verification. It reached an 80.0% success rate—20 percentage points higher than Codex CLI High—while still using 49.6% fewer tokens. 1 2
Long-horizon task benchmarks are one way to look past a polished isolated prompt and see how an agent handles real work. The published comparison uses harness-based development tasks with archived prompts, per-round tool calls, token usage, patches, and verifier results.
The primary comparison below holds the model and reasoning label fixed: Tura Balanced High, Tura Direct High, and the official Codex CLI High configuration on 20 DeepSWE tasks and 5 rewrite tasks. The evidence record also retains Codex CLI Medium as a separate secondary configuration; the benchmark methodology keeps 2 separately reviewed design tasks outside the harness-scored population. 1
The published results do not establish equivalent quality or performance for
every configured provider. Broader Anthropic/Claude, Google/Gemini,
OpenAI-compatible, local-provider, UI-latency, runtime/session parsing, and
cross-OS measurements remain part of the documented
roadmap and known evidence gaps .
GUI page with multi-session concurrent work and HTML rich text support.
TUI page with multi-session concurrent work and HTML rich text support.
The results below come from published benchmark artifacts, not an uncited aggregate. Three systems do most of the work:
Most coding agents still depend on repetitive tool-calling loops: inspect, wait, patch, wait, build, wait, test, wait.
# Turn 1 — inspect environment
rg -n " TODO|command_run|handler " crates/
rg --files crates/runtime/src crates/tools/src
# Turn 2 — apply patch
*** Begin Patch
*** Update File: crates/tools/src/command_run/handler.rs
@@
- // old command handler logic
+ // patched command handler logic
*** End Patch
# Turn 3 — build
cargo build -p runtime
# Turn 4 — run tests
cargo test -p runtime --lib
# Turn 5 — run lint validation
cargo clippy -p runtime --all-targets
Tura takes a different route. Instead of exposing dozens of small tools to the model, it exposes one macro tool: command_run . The agent can then build a multi-step execution tree and run related actions in one LLM turn.
In the example below, both agents run the same commands. A normal tool-calling agent needs five LLM turns; Tura handles the sequence as one structured macro workflow. The saved work is conversational overhead, not engineering discipline.
{
"name" : " command_run " ,
"arguments" : {
"commands" : [
{
"step" : 1 ,
"command_type" : " shell_command " ,
"command_line" : " rg -n \" TODO|command_run|handler \" crates/ "
},
{
"step" : 1 ,
"command_type" : " shell_command " ,
"command_line" : " rg --files crates/runtime/src crates/tools/src "
},
{
"step" : 2 ,
"command_type" : " apply_patch " ,
"command_line" : " *** Begin Patch \n *** Update File: crates/tools/src/command_run/handler.rs \n @@ \n - // old command handler logic \n + // patched command handler logic \n *** End Patch "
},
{
"step" : 3 ,
"command_type" : " shell_command " ,
"command_line" : " cargo build -p runtime "
},
{
"step" : 4 ,
"command_type" : " shell_command " ,
"command_line" : " cargo test -p runtime --lib "
},
{
"step" : 4 ,
"command_type" : " shell_command " ,
"command_line" : " cargo clippy -p runtime --all-targets "
}
]
}
}
There is no ablation test proving that command_run alone causes Tura's lower turn and token usage. In the matched-High DeepSWE comparison, however, Balanced used 66.8% fewer model rounds and 49.6% fewer tokens than Codex CLI High, while Direct used 84.0% fewer rounds and 83.5% fewer tokens. 1 2
However impressive LLMs can be, an LLM is still, at its core, a statistical induction model over text-token probabilities.
For example, asking an LLM to choose among rock, paper, and scissors does not guarantee a uniform random result. If a true one-in-three distribution matters, the choice needs an external random-number source rather than an uncited assumption about model output probabilities.
In coding tasks, this is often fatal.
An agent is more likely to execute and generate code and logic that are statistically more common. But common code and common logic are often mediocre and under-considered.
Tura uses a different strategy.
During reasoning, a common agent reasons from the current state to the prompt goal. In that case, $s_1$ is the current state, and $s_n$ is the goal given by the user prompt.
$$
s_1 \rightarrow s_2 \rightarrow s_3 \rightarrow \cdots \rightarrow s_n
$$
Instead, Tura guides the LLM to statistically estimate $s_{n-1}$ first, then reason backward from the state of $s_{n-1}$ to $s_{n-2}$ .
In the example below, the LLM can derive the optimal strategy for playing rock-paper-scissors correctly.
> To keep rock-paper-scissors fair and challenging,
> We need unbiased play.
> Each move must have a true one-in-three chance.
> An LLM cannot guarantee that from text probabilities alone.
> Use a random-number generator script to generate randint(1, 3)
> Then map rock, paper, or scissors to the number.
In programming tasks, this means that when an agent sees a goal like fixing a frontend bug, it is guided to reason through the full execution path, reconstruct the failure state, and identify the root cause before writing code. In the matched-High DeepSWE comparison, Tura Balanced passed 12 more of 60 binary task verifiers than Codex CLI High.
Both configurations in that contrast use GPT-5.6 SOL with the High reasoning label, so a High-versus-Medium effort mismatch does not explain the 20-point pass-rate difference. The result is still a system-level association, not a causal estimate for backward reasoning or any other individual feature. 1 2
Runtime Context and Prompt Manager
Skills are often just weaker prompts loaded into context.
In many agent frameworks, a long-lived session keeps accumulating skill files, tool outputs, and stale task history. When the context becomes too large, the agent enters a separate compaction turn, but that compaction usually preserves only a compressed summary. Important execution details can become vague or lost.
Tura treats context as part of the runtime state machine.
Instead of relying on users to manually reset sessions or letting Markdown skills pile up, Tura uses task_status , runtime prompts, and recursive execution manuals to keep the active context scoped to the current task.
Traditional skill-based agents usually keep one session running until the user starts another, load broad Markdown skills into that session, and leave them active until a reset or compaction. Tura instead ties runtime prompts to explicit task state: sessions can be renamed, refreshed, and managed automatically; task-specific manuals and CLI commands are loaded through a recursive task tree; and irrelevant context can be removed, replaced, or compacted from the CLI. The checkpoint can retain code locations, patches, tests, and task status rather than only a loose summary. In practice, that means less stale context, lower task-scoped token cost, and fewer chances for an old skill or vague summary to steer the current job.
Because compaction is a CLI operation, Tura can preserve exact execution state in task_status.compact_context . In the published benchmark sessions, Tura moved beyond read-only inspection and resumed execution an average of 2.6 rounds after compaction, compared with an estimated 5.4 rounds for Codex. 3 4 5 6 7
Tura's 2.6-round result is calculated from explicit compact_context events in its archived round contracts. Codex does not expose equivalent compaction events, so its 5.4-round result is estimated from points where input-token usage drops sharply, excluding identifiable media-reading boundaries.
npm install tura-ai
tura
Windows:
npm install - g tura - ai
tura
The same main package is also published to GitHub Packages as @tura-ai/tura .
Configure the @tura-ai scope for https://npm.pkg.github.com , authenticate
with a token that has read:packages , then install @tura-ai/tura . The
unscoped tura-ai package on npm remains the simplest public installation.
Tura does not bundle provider credentials. On first launch, configure an LLM
provider and select one of its models before sending a prompt. See
Provider setup for
the CLI, TUI, and GUI flows.
git clone https: // github.com / Tura - AI / tura.git
cd tura
.\scripts\install.ps1
tura
macOS or Linux shell:
git clone https://github.com/Tura-AI/tura.git
cd tura
./scripts/install.sh
tura
The source installer performs the complete environment setup, release build,
and user PATH registration flow. Pass -EnvironmentOnly on PowerShell or
--environment-only on macOS/Linux only when you intentionally want dependency
setup without building o

[truncated]
