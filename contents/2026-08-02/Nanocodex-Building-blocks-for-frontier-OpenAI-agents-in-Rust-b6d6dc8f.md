---
source: "https://github.com/gakonst/nanocodex"
hn_url: "https://news.ycombinator.com/item?id=49146991"
title: "Nanocodex: Building blocks for frontier OpenAI agents in Rust"
article_title: "GitHub - gakonst/nanocodex: Building blocks for frontier OpenAI agents in Rust. Nanocodex empowers you with Codex-level performance anywhere. · GitHub"
author: "sygma"
captured_at: "2026-08-02T18:55:18Z"
capture_tool: "hn-digest"
hn_id: 49146991
score: 2
comments: 0
posted_at: "2026-08-02T18:25:19Z"
tags:
  - hacker-news
  - translated
---

# Nanocodex: Building blocks for frontier OpenAI agents in Rust

- HN: [49146991](https://news.ycombinator.com/item?id=49146991)
- Source: [github.com](https://github.com/gakonst/nanocodex)
- Score: 2
- Comments: 0
- Posted: 2026-08-02T18:25:19Z

## Translation

タイトル: Nanocodex: Rust のフロンティア OpenAI エージェントのビルディング ブロック
記事のタイトル: GitHub - gakonst/nanocodex: Rust のフロンティア OpenAI エージェントのビルディング ブロック。 Nanocodex は、どこでも Codex レベルのパフォーマンスを実現します。 · GitHub
説明: Rust のフロンティア OpenAI エージェントのビルディング ブロック。 Nanocodex は、どこでも Codex レベルのパフォーマンスを実現します。 - ガコンスト/ナノコーデックス

記事本文:
GitHub - gakonst/nanocodex: Rust のフロンティア OpenAI エージェントのビルディング ブロック。 Nanocodex は、どこでも Codex レベルのパフォーマンスを実現します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ガコンスト
/
ナノコーデックス
公共
通知
あなたはきっとsiでしょう

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
654 コミット 654 コミット .cargo .cargo .github .github ベンチマーク ベンチマーク bin bin crates cratesdeploy デプロイ docs docs evals evals 例 例 Harbor_adapter Harbor_adapter js js py py スクリプト スクリプト Web Web .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .python-version .python-version AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Justfile Justfile LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT PLAN.md PLAN.md README.md README.md REFACTOR.md REFACTOR.md a.sh a.sh崖.toml崖.tomldeny.tomldeny.toml docker-compose.otel.yml docker-compose.otel.yml install install jaeger-ui-config.json jaeger-ui-config.json nanocodex-vm.entitlements nanocodex-vm.entitlements pyproject.toml pyproject.toml release.toml release.toml typos.toml typos.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
フロンティア OpenAI エージェントのビルディング ブロック。
インストール · エージェント API ·
論文 · コンポーネント ·
VM ベースのツール · ドキュメント
Nanocodex CLI を macOS または Linux にインストールします。
カール -fsSL https://nanocodex.paradigm.xyz |バッシュ
または、Rust SDK をアプリケーションに追加します。
カーゴ追加ナノコーデックス
インストールされている CLI をビルド間で切り替えます。
nanocodex アップデート # 最新の安定版リリース
nanocodex アップデート 0.2.0 # ダウングレードを含む正確なリリース
nanocodex 更新 --nightly # 最新の夜間更新
nanocodex update --pr 50 # 検証済みのオンデマンド PR アーティファクト
nanocodex update --path ./nanocodex # 信頼できるローカル バイナリ
ダウンロードしたビルドは ~/.nanocodex/versions に保存されます。ランニング
nanocodex アップデート 0.2.0 は、別のバイナリを使用せずにキャッシュされたバイナリに再び切り替わります。
ダウンロード。安定した打ち上げ

彼女は、
古いバイナリがアクティブであり、~/.nanocodex/current は選択されたバイナリを指します。
バージョン。
PR アーティファクトには、認証された gh CLI とすでに完了した
その PR のオンデマンド アーティファクト ワークフロー。
nanocodex を使用します:: { Nanocodex 、 OpenAi } ;
let openai = OpenAi :: new ( std :: env :: var ( "OPENAI_API_KEY" ) ? ) ? ;
let ( エージェント , mut イベント ) = Nanocodex :: builder ( openai )
。指示 (
「あなたは Rust コーディング エージェントです。焦点を絞った変更を加え、無関係な作業を保存します。\
終了する前に関連するテストを実行します。" ,
)
。ワークスペース ( std :: env :: current_dir ( ) ? )
。建てる （ ） ？ ;
letevent_task = tokio :: spawn ( async move {
while let Some (event) = events 。受信() 。待ってください{
エプリントリン！ ( "イベント {}: {:?}" 、イベント . seq 、イベント . kind ) ;
イベントの場合種類もis_terminal ( ) {
休憩;
}
}
} ) ;
// 代替案: このターンの応答が到着したときにストリーミングします。
// futures_util::StreamExt を使用します。
// nanocodex::agent::events::{AgentEventData, AssistantEvent} を使用します。
// mut を回しましょう = Agent.prompt("失敗したパーサー テストを見つけて修正します。").await?;
// while let Some(event) =turn.next().await {
// AgentEventData::Assistant(AssistantEvent::Delta(delta)) = events.data() とすると? {
// print!("{}", delta.text);
// }
// }
結果 = エージェントにする
。プロンプト (「失敗したパーサー テストを見つけて修正します。」)
。待ってますか？
。待ってますか？ ;
イベントタスク 。待ってますか？ ;
プリントイン！ ( "{}" , result .final_message() ) ;
最初の await はプロンプトを受け入れて命令します。 2 番目は入力されるのを待ちます
TurnResult 。後続のプロンプトでは、エージェントが保持しているプロンプトが自動的に再利用されます。
履歴、WebSocket、ツール、シェル セッション、およびプロンプト キャッシュ ID。
Agent.clone() は、同じセッションへの安価なハンドルです。独立して
返された AgentEvents ストリームは、セッション全体のイベント Firehose です。
Nanocodex は gpt をサポートします

-5.6-sol (デフォルト) および gpt-5.6-luna 。を選択します。
エージェントを作成するときに .model(Model::Luna) を使用してモデルを作成します。モデルは固定されています
そのスレッド: 後で切り替えるとプロバイダーのチェックポイントが無効になり、
保持されている完全なコンテキストを非効率的に再生する必要があります。
非 TUI デスクトップの例は、デフォルトのマイクとスピーカーを直接所有します。
Rust では、運用 TUI と同じ VoiceSessionBuilder を使用します。
nanocodex 認証ログイン # 1 回; ~/.codex/auth.json を Codex と共有します
カーゴ ラン -p nanocodex-examples --bin voice
下位アダプターは、デバイスとメディアの所有権を Nanocodex の外に残します。こう書かれています
標準入力からの 24 kHz モノラル PCM16 リトルエンディアン オーディオ、同じフォーマットを書き込みます
stdout に保存し、トランスクリプトとエージェント イベントを stderr に保存します。
カーゴ ラン -p nanocodex-examples --bin realtime-pipe < マイク.pcm > スピーカー.pcm
# 同様に、ライブ キャプチャ/デコーダと再生/エンコーダを作成します。
キャプチャ-s16le |カーゴ実行 --quiet -p nanocodex-examples --bin realtime-pipe |プレイ-s16le
どちらも 1 つのコーディング エージェント セッションを保持します。音声による要求は、アイドル状態の間に作業を開始します。
その作業中に受け取ったフォローアップは、アクティブなターンをアトミックにその時点で操作します。
次の安全なモデル境界。どちらも、共有 Codex/ChatGPT サブスクリプション認証を使用します。
API キー。 NANOCODEX_AUTH_FILE を設定して、通常の Codex 資格情報をオーバーライドします。
パス。
小さくて優れた構成要素
エージェント インフラストラクチャは、各部分に
シャープな所有者と独自の便利な API。 OpenAI クライアントは、
エージェントループ。ツールは CLI なしで動作する必要があります。上級エージェントは次のことを行う必要があります。
それらの別の実装を隠すのではなく、それらの部分を構成します。
Nanocodex は、Rust、Tower、typed など、少数の意図的な選択を行います。
プロトコル、所有ライフサイクル状態、ビルダー API - 境界を維持します
退屈。
モ

デルとハーネスは共同設計されています
私たちは、フロンティアモデルやコーデックスがすでに行っている行動を出し抜こうとはしません。
明示的な。コンテキスト管理、AGENTS.md、圧縮、キャッシュ ID、ツール
シェイプ、継続、再接続、リプレイ、キャンセル、プロセスのクリーンアップは、
モデル向け契約の一部。
Nanocodex は、これらの不変条件をより小さなライブラリファースト API に組み込みます。
アプリケーション ポリシーを呼び出し元に残します。
代表的なカーゴベンチのワークロード、OpenTelemetry トレース、差分
テストとエンドツーエンドの評価により、ハーネスが正直に保たれます。目標はシンプル、「普通」
エージェントのターンは、トークンの使用量と、モデルとネットワークの遅延に依存する必要があります。
結果と同じ型付き境界で表示される推定 USD コスト。
nanocodex 合金スタイルのファサードとプレリュード
§── エージェント nanocodex-agent
│ §── oai nanocodex-oai-api
│ └── ツール nanocodex-tools
│ └── マクロ nanocodex-tools-macros
§── oai nanocodex-oai-api
§── ツール nanocodex-tools
└── 可観測性 nanocodex-可観測性 (オプション)
ファサードは、正規の共通インポートを提供します。下部の各クレートも、
上位のオーケストレーションをインポートせずに、直接役立つように設計されています。
層。
薄いファサードは、クレート ルートでゴールデン エージェント パスを再エクスポートし、
nanocodex::agent 、 nanocodex::oai 、および
nanocodex::tools 。そのプレリュードには、構築に必要な一般的なタイプのみが含まれています
エージェント。
ファサードガイド・
APIドキュメント
バッテリーを含むライフサイクル: 所有する専用ドライバー、安価なクローン可能ドライバー
Nanocodex ハンドル、型指定された Turn 値と TurnResult 値、およびオプションのイベント
ストリーム。プロンプト注文、ツールループ、AGENTS.md ディスカバリ、
圧縮のタイミング、キャンセル、スナップショット、 spawn による分岐、
fork 、および fork_from 。
発信者が以前のメッセージを渡すことはありません。

応答 ID、またはツールの結果を返します
エージェント。
エージェントガイド・
APIドキュメント
完全な OpenAI 境界: API キーと ChatGPT 認証、型付き
応答プロトコル値、永続的な WebSocket トランスポート、クライアント所有
コンテキスト、継続と再生、自動価格設定、および汎用タワー
クライアント。
スタンドアロンの OpenAi -> セッション -> ResponseTurn -> Response パスは、
エージェントのポリシーを採用せずに会話を管理できます。カスタムタワーレイヤーと
サービスは具体的で名前を付けることができるため、ボックス化クライアントやグローバル クライアントは必要ありません。
OpenAI API ガイド ·
APIドキュメント
モデル対応ツール ランタイム: ツール コントラクト、異種ツール
レジストリ、標準ワークスペース ツール、シェルとプロセスのライフサイクル、コード モード、
deferred tool_search 、リモート ディスパッチ、および MCP。 MCP は常に次の場所で利用できます。
ネイティブターゲット。
アプリケーションはツールを直接実装するか、再エクスポートされた #[tool] を使用できます。
マクロ。別個の nanocodex-tools-macros パッケージは Rust 用にのみ存在します。
手続き型マクロの境界。
ツールガイド・
APIドキュメント
すでに流れているデータに対するアプリケーション所有のトレースと OpenTelemetry セットアップ
エージェントを通して。構造化されたライフサイクル、モデル、ツール、使用法、コスト、
コア ランタイム パスを変更せずに、キャッシュとレイテンシ テレメトリを実行できます。
ファサードの可観測性機能を有効にするか、コンポーネントに依存します
直接的に。
可観測性ガイド ·
APIドキュメント
公的契約がまだ満期に達しているコンポーネントは、以下の条件下で存続します。
クレート/実験用/ :
CLI はこれらのクレートのコンシューマです。音声および VM ベースのツールは依然として機能が薄いままですが、
安定したライブラリ契約を介したオプトイン アダプター。
CLI/TUI、Python パッケージ、ノード/ブラウザ パッケージ、React バインディング、および例
これらは、同じ所有セッション API のシン コンシューマです。秒を定義するものではありません
エージェントプロトコル。
例 · JavaScript ·
パイソン・W

eb
通常の TUI およびワンショット セッションでは、デフォルトでホスト ワークスペース ツールが保持されます。彼らはできる
代わりに、 exec_command 、 write_stdin 、 apply_patch 、および view_image をルートします。
1 つの保持された VM 経由:
build-vm-guest だけ
ナノコーデックス\
--vm .nanocodex/vm/session-rootfs.ext4 \
--vm-guest-runtime target/aarch64-unknown-linux-musl/debug/nanocodex-vm-guest \
--vm-workspace /app
nanocodex を実行します「リポジトリを検査します」\
--vm .nanocodex/vm/session-rootfs.ext4 \
--vm-guest-runtime target/aarch64-unknown-linux-musl/debug/nanocodex-vm-guest \
--vm-workspace /app
ディレクトリルートには代わりに次のものが含まれる場合があります
/usr/local/bin/nanocodex-vm-guest 。を参照してください。
イメージの準備、ライフサイクル、エグレス、Linux に関する VM ガイド
要件と macOS の署名。
ベンチマークと保持された測定値
Rust のフロンティア OpenAI エージェントのビルディング ブロック。 Nanocodex は、どこでも Codex レベルのパフォーマンスを実現します。
Readme Apache-2.0、MIT ライセンスが見つかりました。
32 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Building blocks for frontier OpenAI agents in Rust. Nanocodex empowers you with Codex-level performance anywhere. - gakonst/nanocodex

GitHub - gakonst/nanocodex: Building blocks for frontier OpenAI agents in Rust. Nanocodex empowers you with Codex-level performance anywhere. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
gakonst
/
nanocodex
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
654 Commits 654 Commits .cargo .cargo .github .github benchmarks benchmarks bin bin crates crates deploy deploy docs docs evals evals examples examples harbor_adapter harbor_adapter js js py py scripts scripts web web .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .python-version .python-version AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Justfile Justfile LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT PLAN.md PLAN.md README.md README.md REFACTOR.md REFACTOR.md a.sh a.sh cliff.toml cliff.toml deny.toml deny.toml docker-compose.otel.yml docker-compose.otel.yml install install jaeger-ui-config.json jaeger-ui-config.json nanocodex-vm.entitlements nanocodex-vm.entitlements pyproject.toml pyproject.toml release.toml release.toml typos.toml typos.toml uv.lock uv.lock View all files Repository files navigation
Building blocks for frontier OpenAI agents.
Install · Agent API ·
Thesis · Components ·
VM-backed tools · Documentation
Install the Nanocodex CLI on macOS or Linux:
curl -fsSL https://nanocodex.paradigm.xyz | bash
Or add the Rust SDK to an application:
cargo add nanocodex
Switch the installed CLI between builds:
nanocodex update # latest stable release
nanocodex update 0.2.0 # exact release, including downgrades
nanocodex update --nightly # latest nightly
nanocodex update --pr 50 # verified on-demand PR artifact
nanocodex update --path ./nanocodex # trusted local binary
Downloaded builds are retained under ~/.nanocodex/versions . Running
nanocodex update 0.2.0 again switches to the cached binary without another
download. A stable launcher keeps nanocodex update available even while an
older binary is active, and ~/.nanocodex/current points to the selected
version.
PR artifacts require an authenticated gh CLI and an already completed
on-demand artifact workflow for that PR.
use nanocodex :: { Nanocodex , OpenAi } ;
let openai = OpenAi :: new ( std :: env :: var ( "OPENAI_API_KEY" ) ? ) ? ;
let ( agent , mut events ) = Nanocodex :: builder ( openai )
. instructions (
"You are a Rust coding agent. Make focused changes, preserve unrelated work, \
and run relevant tests before finishing." ,
)
. workspace ( std :: env :: current_dir ( ) ? )
. build ( ) ? ;
let event_task = tokio :: spawn ( async move {
while let Some ( event ) = events . recv ( ) . await {
eprintln ! ( "event {}: {:?}" , event . seq , event . kind ) ;
if event . kind . is_terminal ( ) {
break ;
}
}
} ) ;
// Alternative: stream this turn's response as it arrives:
// use futures_util::StreamExt;
// use nanocodex::agent::events::{AgentEventData, AssistantEvent};
// let mut turn = agent.prompt("Find and fix the failing parser test.").await?;
// while let Some(event) = turn.next().await {
// if let AgentEventData::Assistant(AssistantEvent::Delta(delta)) = event.data()? {
// print!("{}", delta.text);
// }
// }
let result = agent
. prompt ( "Find and fix the failing parser test." )
. await ?
. await ? ;
event_task . await ? ;
println ! ( "{}" , result . final_message ( ) ) ;
The first await accepts and orders the prompt. The second waits for its typed
TurnResult . Follow-on prompts automatically reuse the agent's retained
history, WebSocket, tools, shell sessions, and prompt-cache identity.
agent.clone() is a cheap handle to that same session; the independently
returned AgentEvents stream is the session-wide event firehose.
Nanocodex supports gpt-5.6-sol (the default) and gpt-5.6-luna . Select the
model with .model(Model::Luna) when creating an agent. The model is fixed for
that thread: switching later would invalidate the provider checkpoint and
require an inefficient replay of the complete retained context.
The non-TUI desktop example owns the default microphone and speaker directly
in Rust, using the same VoiceSessionBuilder as the production TUI:
nanocodex auth login # once; shares ~/.codex/auth.json with Codex
cargo run -p nanocodex-examples --bin voice
The lower adapter leaves device and media ownership outside Nanocodex. It reads
24 kHz mono PCM16 little-endian audio from stdin, writes the same format to
stdout, and keeps transcripts and agent events on stderr:
cargo run -p nanocodex-examples --bin realtime-pipe < microphone.pcm > speaker.pcm
# Equivalently, compose any live capture/decoder and playback/encoder:
capture-s16le | cargo run --quiet -p nanocodex-examples --bin realtime-pipe | play-s16le
Both retain one coding-agent session. A spoken request starts work while idle;
a follow-up received during that work atomically steers the active turn at its
next safe model boundary. Both use shared Codex/ChatGPT subscription auth, not
an API key. Set NANOCODEX_AUTH_FILE to override the normal Codex credential
path.
Small, excellent building blocks
Agent infrastructure is easier to understand and reuse when each piece has a
sharp owner and a useful API of its own. An OpenAI client should work without
an agent loop. Tools should work without a CLI. The high-level agent should
compose those pieces rather than hide another implementation of them.
Nanocodex makes a small number of deliberate choices—Rust, Tower, typed
protocols, owned lifecycle state, and builder APIs—then keeps the boundaries
boring.
The model and harness are co-designed
We do not try to outsmart behavior that frontier models and Codex already make
explicit. Context management, AGENTS.md , compaction, cache identity, tool
shapes, continuation, reconnect replay, cancellation, and process cleanup are
parts of the model-facing contract.
Nanocodex carries those invariants into a smaller, library-first API while
leaving application policy with the caller.
Representative cargo bench workloads, OpenTelemetry traces, differential
tests, and end-to-end evals keep the harness honest. The goal is simple: normal
agent turns should be model- and network-latency bound, with token usage and
estimated USD cost visible at the same typed boundary as the result.
nanocodex Alloy-style facade and prelude
├── agent nanocodex-agent
│ ├── oai nanocodex-oai-api
│ └── tools nanocodex-tools
│ └── macros nanocodex-tools-macros
├── oai nanocodex-oai-api
├── tools nanocodex-tools
└── observability nanocodex-observability (optional)
The facade provides the canonical common imports. Each lower crate is also
designed to be useful directly, without importing the higher orchestration
layer.
The thin facade reexports the golden agent path at the crate root and keeps
detailed APIs under nanocodex::agent , nanocodex::oai , and
nanocodex::tools . Its prelude contains only the common types needed to build
an agent.
Facade guide ·
API documentation
The batteries-included lifecycle: an owned private driver, a cheap cloneable
Nanocodex handle, typed Turn and TurnResult values, and an optional event
stream. It owns prompt ordering, the tool loop, AGENTS.md discovery,
compaction timing, cancellation, snapshots, and branching through spawn ,
fork , and fork_from .
Callers never pass previous messages, response IDs, or tool results back into
the agent.
Agent guide ·
API documentation
The complete OpenAI boundary: API-key and ChatGPT authentication, typed
Responses protocol values, a persistent WebSocket transport, client-owned
context, continuation and replay, automatic pricing, and a generic Tower
client.
Its standalone OpenAi -> Session -> ResponseTurn -> Response path provides a
managed conversation without taking on agent policy. Custom Tower layers and
services remain concrete and nameable—no boxing or global client is required.
OpenAI API guide ·
API documentation
The model-facing tool runtime: the Tool contract, heterogeneous Tools
registry, standard workspace tools, shell and process lifecycle, Code Mode,
deferred tool_search , remote dispatch, and MCP. MCP is always available on
native targets.
Applications can implement Tool directly or use the reexported #[tool]
macro. The separate nanocodex-tools-macros package exists only for Rust's
procedural-macro boundary.
Tools guide ·
API documentation
Application-owned tracing and OpenTelemetry setup for the data already flowing
through the agent. It provides structured lifecycle, model, tool, usage, cost,
cache, and latency telemetry without changing the core runtime path.
Enable the facade's observability feature or depend on the component
directly.
Observability guide ·
API documentation
Components whose public contracts are still maturing live under
crates/experimental/ :
The CLI is a consumer of these crates. Voice and VM-backed tools remain thin,
opt-in adapters over the stable library contracts.
The CLI/TUI, Python package, Node/browser package, React bindings, and examples
are thin consumers of the same owned session API. They do not define a second
agent protocol.
Examples · JavaScript ·
Python · Web
Normal TUI and one-shot sessions keep host workspace tools by default. They can
instead route exec_command , write_stdin , apply_patch , and view_image
through one retained VM:
just build-vm-guest
nanocodex \
--vm .nanocodex/vm/session-rootfs.ext4 \
--vm-guest-runtime target/aarch64-unknown-linux-musl/debug/nanocodex-vm-guest \
--vm-workspace /app
nanocodex run " inspect the repository " \
--vm .nanocodex/vm/session-rootfs.ext4 \
--vm-guest-runtime target/aarch64-unknown-linux-musl/debug/nanocodex-vm-guest \
--vm-workspace /app
Directory roots may instead contain
/usr/local/bin/nanocodex-vm-guest . See the
VM guide for image preparation, lifecycle, egress, Linux
requirements, and macOS signing.
Benchmarks and retained measurements
Building blocks for frontier OpenAI agents in Rust. Nanocodex empowers you with Codex-level performance anywhere.
Readme Apache-2.0, MIT licenses found Activity Stars
32 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
