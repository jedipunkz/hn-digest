---
source: "https://github.com/surya17495/centri"
hn_url: "https://news.ycombinator.com/item?id=48564976"
title: "Show HN: Building a Stateful AI Agent"
article_title: "GitHub - surya17495/centri: Memory-native OpenCode fork with Centri core, event-spine memory, verbatim recall, and Hermes integration. · GitHub"
author: "suryaankata"
captured_at: "2026-06-17T04:35:59Z"
capture_tool: "hn-digest"
hn_id: 48564976
score: 2
comments: 0
posted_at: "2026-06-17T02:18:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Building a Stateful AI Agent

- HN: [48564976](https://news.ycombinator.com/item?id=48564976)
- Source: [github.com](https://github.com/surya17495/centri)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T02:18:11Z

## Translation

タイトル: HN の表示: ステートフル AI エージェントの構築
記事のタイトル: GitHub - surya17495/centri: Centri コア、イベントスパイン メモリ、逐語的呼び出し、Hermes 統合を備えたメモリネイティブ OpenCode フォーク。 · GitHub
説明: Centri コア、イベントスパイン メモリ、逐語的呼び出し、Hermes 統合を備えたメモリネイティブ OpenCode フォーク。 - surya17495/centri
HN テキスト: 数か月前に構築の旅を始めたとき、AI エージェント/コーディング ツールがステートフルではないことが私にとって大きな苦痛でした。そこで、自律メモリ管理が組み込まれた Opencode フォーク (ビルド用の現在の毎日のドライバー) を出荷しました。メモリ コアは、プロンプトをほとんど表示せずに、Hermes (同梱) やその他のデスクトップ/コーディング エージェントに接続することもできます。まだ開発中ですが、試してみてフィードバックをお聞かせください。目標は、サンドボックスまたは VM でサーバーが実行される音声およびモバイルファーストのクライアントを構築することであり、Opencode と同様にすべてのモデルで動作し、Oracle VM で実行されている現在の hermes+opencode セットアップを置き換えることです。うーん、これでコラボしたいならどうぞ:)

記事本文:
GitHub - surya17495/centri: Centri コア、イベントスパイン メモリ、逐語的再現、Hermes 統合を備えたメモリネイティブ OpenCode フォーク。 · GitHub
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
スーリヤ17495
/
中心
公共
通知
通知設定を変更するにはサインインする必要があります
あ

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14,178 コミット 14,178 コミット .github .github .husky .husky .vscode .vscode .zed .zed core core デプロイ デプロイ docs docs github github infra infra nix nix パッケージ パッケージ パッチ パッチ perf perf スクリプト スクリプト スクリプト スクリプト sdks/ vscode sdks/ vscode シェル シェル スペック アップストリーム/ opencode アップストリーム/ opencode .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .gitleaksignore .gitleaksignore .oxlintrc.json .oxlintrc.json .prettierignore .prettierignore AGENTS.md AGENTS.md CONTEXT.md CONTEXT.md Contributing.md Contributing.md FORK-NOTES.md FORK-NOTES.md HANDOFF.md HANDOFF.md ライセンス ライセンス ライセンス-OPENCODE ライセンス-OPENCODE README.md README.md SECURITY.md SECURITY.md STATE-AND-HANDOFF.md STATE-AND-HANDOFF.md VERIFY.md VERIFY.md bun.lock bun.lock bunfig.toml bunfig.toml docker-compose.yml docker-compose.yml flake.lock flake.lock flake.nix flake.nix インストール インストール package.json package.json スクリーンショット-uk.png スクリーンショット-uk.png sst-env.d.ts sst-env.d.ts sst.config.ts sst.config.ts tsconfig.json tsconfig.json Turbo.json Turbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
メモリファーストのコーディング エージェント。耐久性のある 1 つの記憶背骨がすべてを記憶します
そしてあなたのツールもそうなったので、新しいターンは冷たいのではなく暖かい状態で始まります。
エージェントは忘れてしまいます。コンテキスト ウィンドウがいっぱいになり、セッションが終了し、新しいチャットが開始されます
ゼロから – すでに下した決定を再説明し、エージェントの様子を観察します。
既に拒否したアプローチを再提案します。 Centri はこれを逆転します。追加のみです。
イベントスパインは真実の源であり、記憶は派生インデックスであり、
破棄と再導出、およびターンごとのすべてのコンテキスト

によって新鮮に組み立てられます
すべての行にレシートを添付する確定的なキュレーション機能。の
コンテキスト ウィンドウはストレージではなくキャッシュです。
3 つの部分が 1 つの耐久性のあるメモリを共有します。
イベント スパインは真実の源です。メモリグラフは派生したもので、
それに対する再導出可能なインデックス。 Centri は MIT ライセンスのフォークです。
OpenCode プロジェクト — アップストリーム アトリビューション
LICENSE-OPENCODE に保存されます。
上流/opencode/ 、 FORK-NOTES.md 、および
docs/centri-app.md 。
追加専用イベント スパイン - すべてのツール呼び出し、ファイル編集、決定、および結果
書き込み時にシークレット編集を行って永続的にログに記録されます。背骨はシステムの
真実の源。
バイタイムスーパーセッションを使用した型付きメモリ グラフ — 意思決定、事実、および
オープンループ。新しい真実は古い真実を無効にしますが、歴史は保持されます: 古くなります
事実が要約で再浮上することはありませんが、完全なタイムラインは監査可能のままです。
レシートによる確定的なキュレーション — 各ターンごとのブリーフは、
1 つの純粋な curate() パス。すべての行には、後方を指すsource_event_idが含まれています
元のレジャーイベントに。同じ（グラフ、手がかり、予算、方針）
バイト同一のブリーフが生成されます。読み取り時に LLM は実行されません。
FTS5 を逐語的に思い出す — スパイン上の SQLite FTS5 インデックスにより、簡単にリフトアップできます。
正確な以前のトークン (ファイル名、エラー文字列、識別子) をコンテキストに取り込む
入力されたグラフの横にあります。
LLM 統合 - オフライン作業者が生の背骨を周囲環境に折り畳む
レイヤー: アイデンティティ、アクティブなプロジェクト、トップオープンループ、最近の短い物語、
エージェントがあなたを繰り返し見てきた好みや慣例のユーザー プロフィール。
時間的回想 — 「昨日から何が変わったのか」または「どこから出発したのか」を尋ねます
オフ」と選択すると、推測ではなく、台帳のタイムラインに基づいた答えが得られます。
ACP を介したコーディング委任 — 作業はコーディング エージェント (エージェント クライアント) に渡されます。
プロトコル、標準入出力上の JSON-RPC)

ライブストリーミングの進行状況、承認ゲート
破壊的なアクション、およびフォールバック ハンドへのフェイルオーバーの場合。
ファーストクラスのツール契約 — ツール (Composio、例: Tavily 検索) が隣にあります
コーディングの手。すべての呼び出しはイベント記録され、副作用のあるツールが渡されます。
承認ゲートを通過すると、読み取り専用の結果がメモリにフォールバックされます。
履歴のインポート — ワンショットのブートストラップと OpenCode の継続的なテール、
クロード コードとカーソル履歴のため、新規インストールは完全な状態で開始されます。
白紙の状態ではなく記憶を。
ヘルメスの構造化された取り込み - ヘルメスのチャットは入力された重複排除可能なものとして取り込まれます
エンベロープ ( hermes.user.message 、 hermes.assistant.message 、
hermes.tool.result 、 hermes.memory.write )、フラット化されたテキストではありません。
出来事は真実の源です。メモリは、それらに対する導出された再導出可能なインデックスです。
完全なインデックスについては docs/README.md を参照してください。または
docs/architecture.md 、
docs/memory-architecture.md 、および
詳細については、docs/event-contract.md を参照してください。
cp .env.example .env # モデルのゲートウェイ キーを入力します (BYOK)
docker compose up -d # コアは :8760、Web シェルは :8761
マニュアル
CDコア
Python -m venv .venv
。 .venv/bin/activate
pip install -e " .[dev] "
cp ../.env.example ../.env # キーを入力します (BYOK)
python -m pytest testing/ -v # スイートを実行します
centri # サーバーを起動します (または: python -m centri.cli)
API はデフォルトで 127.0.0.1:8760 をリッスンします。ヘルスチェック:
カールローカルホスト:8760/health
カールローカルホスト:8760/ステータス
構成
すべての構成は環境主導型です。.env.example を次の場所にコピーします。
.env に移動し、キーを入力します。何もコミットされていません。 .env と *.db は
無視されました。 Centri は BYOK (Bring Your Own Key) であり、モデルに依存しません。
モデルゲートウェイ - 任意のLITELLM_BASE_URL / LITELLM_API_KEYポイント
OpenAI互換プロバイダー。ロールモデルはタスクごとに設定されます ( MODEL_INTENT 、
MODEL_REASONING , …)。
認証 — CENTRI_AUTH

_TOKEN は、/health (および
/events/stream WebSocket (?token= 経由)。空は認証がオフであることを意味します。前に設定してください
ポートを公開します。
ツール — CENTRI_COMPOSIO_API_KEY は Composio を有効にします。 CENTRI_COMPOSIO_TOOLS
カンマ区切りの許可リストです (デフォルトは TAVILY_SEARCH )。鍵が無いと、
プロバイダーは、何らかの理由で利用できないことを報告し、ネットワークに接続することはありません。
埋め込み (オプション) — デフォルトではオフです (字句 + FTS5 リコールのみ)。ターン
CENTRI_EMBEDDING_* によるセマンティック・リコールについて。 .env.example を参照してください。
取り込みパス — CENTRI_INGEST_OPENCODE_PATHS 、
CENTRI_INGEST_CLAUDE_CODE_PATHS および CENTRI_INGEST_CURSOR_PATHS は
異常な場所の履歴をプラットフォームごとに調査します。
CENTRI_INGEST_DISABLED_AGENTS はエージェントをオプトアウトします。
リファレンス デプロイでは、1 つのメモリ DB を共有する 2 つのサービスが実行されます。
( ~/.centri/state.db ):
sudo systemctl Enable --now centri-core opencode
systemctl is-active centri-core opencode
カール -fsS http://127.0.0.1:8760/health # -> {"ステータス":"ok",...}
完全な展開ガイド (systemd、Caddy/TLS、ポート): docs/DEPLOY.md 。
opencode.service ログの xdg-open ENOENT 行は無害です — OpenCode
ヘッドレスボックス上でブラウザを自動的に開こうとし、生成後も実行し続けます。
失敗します。
Centri は、Hermes メモリ.プロバイダとしても出荷されます。薄いプラグイン
( CentriMemoryProvider ) は、Hermes メモリ呼び出しをコアの HTTP API に変換します。
プリフェッチ → POST /memory/recall 、 sync_turn / on_memory_write → バッチ処理
POST /events/import 。デプロイ可能なプラグインは次の場所にあります。
デプロイ/hermes-plugin/centri/ 。
# ~/.hermes/config.yaml
メモリ:
プロバイダー：セントリ
中心:
api_base : http://127.0.0.1:8760
auth_token : <CENTRI_AUTH_TOKEN> # コア + フォークの CENTRI_TOKEN と同じ値
プラグインを変更した後は、Hermes を再起動します。完全なガイド:
docs/エルメス-統合.md 。
新規インストールでは y がインポートされます

既存のコーディング エージェントの履歴があるため、記憶は
初日から完了。利用可能なものを発見し、一度ブートストラップします (冪等)
— インポートを再実行しても新しいものは何もありません):
curl localhost:8760/ingest/discover # "N 個の OpenCode メッセージ、M 個のカーソル セッションが見つかりました"
curl -X POST localhost:8760/ingest/bootstrap \
-H ' content-type: application/json ' -d ' {} ' # 1 回限りの完全インポート
ブートストラップの後、アンビエント テールはスケジューラがティックするたびに新しい履歴を取得し続けます。
セントリベンチは反証可能な直接対決です。各エンジンはターンごとに組み立てられます。
同じシード履歴からのブリーフ、ブリーフの完全性、再提案でスコア付け
レート、次のステップの正確性、および古い事実の処理。ネイティブ Centri が実行される
実際の Letta サーバー (v0.16.8、pgvector アーカイブ) に対して — 両方とも
決定論的ルーブリックとLLM裁判官はセントリの勝利に同意、その差は完全に
古い事実の置き換え。
Letta はベンチマーク比較のみであり、実行時の依存関係ではありません — Centri は実行します
それなしで。自分で実行してください:
CDコア
python -m centri.bench.run # 人間が読めるレポート
python -m centri.bench.run --json # 機械可読スコア
方法論: docs/centri-bench.md ;生の結果:
docs/ベンチ結果/ 。
コア、メモリ グラフ、キュレーション、ACP コーディング ループ、ツール コントラクト、および履歴
取り込みは 385 テスト スイート ( cd core && python -m pytest testing/ ) でカバーされています。
ユニットスイートはグリーンオフラインで実行されます。ライブコアを必要とする統合テスト、BYOK
モデルキー、オープンコードバイナリ、または実際のLettaサーバーは環境ゲートされています
それらが存在しない場合はきれいにスキップします。
OpenCode フォーク Web アプリはビルドと型チェックを行います。ロードマップは
docs/ROADMAP.md ;完全なドキュメントインデックスは次のとおりです
docs/README.md 。
Centri は MIT ライセンスに基づいてライセンスされています。「ライセンス」を参照してください。の
package/opencode/ の OpenCode アプリ シェルは、
MIT ライセンスの OpenCode プロジェクト。その
上流l

ライセンスは LICENSE-OPENCODE に保存され、
オリジナルのアップストリーム README は、upstream/opencode/ に保管されています。
ランタイムの依存関係には寛容なライセンスが適用されます (MIT/BSD/Apache-2.0)。
letta-client はベンチ専用ツールであり、製品に依存するものではありません。
Centri コア、イベントスパイン メモリ、逐語的呼び出し、Hermes 統合を備えたメモリネイティブの OpenCode フォーク。
MIT、MIT ライセンスが見つかりました
ライセンスが見つかりました
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

Memory-native OpenCode fork with Centri core, event-spine memory, verbatim recall, and Hermes integration. - surya17495/centri

As I started my building journey few months ago, AI agents/coding tools not being stateful has been a big pain for me. So I just shipped an Opencode fork (my current daily driver for building) with built-in autonomous memory management, the memory core can also be plugged into Hermes (included) and other desktop/coding agents with few prompts. It's WIP, but please try it and give me your feedback. Goal is to build a voice and mobile first client with server running in a sandbox or VM and works with all models like Opencode does, and replaces my current hermes+opencode setup running on a Oracle VM. hmu if you wanna collab on this :)

GitHub - surya17495/centri: Memory-native OpenCode fork with Centri core, event-spine memory, verbatim recall, and Hermes integration. · GitHub
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
surya17495
/
centri
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14,178 Commits 14,178 Commits .github .github .husky .husky .vscode .vscode .zed .zed core core deploy deploy docs docs github github infra infra nix nix packages packages patches patches perf perf script script scripts scripts sdks/ vscode sdks/ vscode shell shell specs specs upstream/ opencode upstream/ opencode .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .gitleaksignore .gitleaksignore .oxlintrc.json .oxlintrc.json .prettierignore .prettierignore AGENTS.md AGENTS.md CONTEXT.md CONTEXT.md CONTRIBUTING.md CONTRIBUTING.md FORK-NOTES.md FORK-NOTES.md HANDOFF.md HANDOFF.md LICENSE LICENSE LICENSE-OPENCODE LICENSE-OPENCODE README.md README.md SECURITY.md SECURITY.md STATE-AND-HANDOFF.md STATE-AND-HANDOFF.md VERIFY.md VERIFY.md bun.lock bun.lock bunfig.toml bunfig.toml docker-compose.yml docker-compose.yml flake.lock flake.lock flake.nix flake.nix install install package.json package.json screenshot-uk.png screenshot-uk.png sst-env.d.ts sst-env.d.ts sst.config.ts sst.config.ts tsconfig.json tsconfig.json turbo.json turbo.json View all files Repository files navigation
A memory-first coding agent. One durable memory spine remembers everything you
and your tools have done, so every new turn starts warm instead of cold.
Agents forget. Context windows fill, sessions end, and every new chat starts
from scratch — so you re-explain decisions you already made and watch the agent
re-propose approaches you already rejected. Centri inverts this: an append-only
event spine is the source of truth, memory is a derived index that can be
thrown away and re-derived, and every per-turn context is assembled fresh by a
deterministic curation function that attaches a receipt to every line. The
context window is a cache, not storage.
Three parts share one durable memory:
The event spine is the source of truth ; the memory graph is a derived,
re-derivable index over it. Centri is a fork of the MIT-licensed
OpenCode project — upstream attribution
is preserved in LICENSE-OPENCODE ,
upstream/opencode/ , FORK-NOTES.md , and
docs/centri-app.md .
Append-only event spine — every tool call, file edit, decision, and result
is durably logged with secret redaction on write. The spine is the system's
source of truth.
Typed memory graph with bi-temporal supersession — decisions, facts, and
open loops. New truth invalidates old truth, but history is retained: stale
facts never resurface in a brief, while the full timeline stays auditable.
Deterministic curation with receipts — each per-turn brief is rendered by
one pure curate() path; every line carries a source_event_id pointing back
to the ledger event it came from. The same (graph, cue, budget, policy)
yields a byte-identical brief. No LLM runs at read time.
FTS5 verbatim recall — a SQLite FTS5 index over the spine lets a brief lift
exact prior tokens (file names, error strings, identifiers) into context
alongside the typed graph.
LLM consolidation — an offline worker folds the raw spine into the ambient
layer: identity, active projects, top open loops, a short recent narrative, and
a user profile of preferences and conventions the agent has seen you repeat.
Temporal recall — ask "what changed since yesterday" or "where did we leave
off" and get an answer grounded in the ledger's timeline, not a guess.
Coding delegation over ACP — work is handed to a coding agent (Agent Client
Protocol, JSON-RPC over stdio) with live streaming progress, an approval gate
for destructive actions, and failover to a fallback hand.
First-class tool contract — tools (Composio, e.g. Tavily search) sit beside
the coding hands. Every invocation is event-ledgered, side-effectful tools pass
through the approval gate, read-only results fold back into memory.
History import — a one-shot bootstrap plus a continuous tail of OpenCode,
Claude Code, and Cursor histories, so a fresh install starts with complete
memory instead of a blank slate.
Hermes structured ingestion — Hermes chat is ingested as typed, dedupable
envelopes ( hermes.user.message , hermes.assistant.message ,
hermes.tool.result , hermes.memory.write ), not flattened text.
Events are the source of truth; memory is a derived, re-derivable index over them.
See docs/README.md for the full index, or
docs/architecture.md ,
docs/memory-architecture.md , and
docs/event-contract.md for detail.
cp .env.example .env # fill in your model gateway keys (BYOK)
docker compose up -d # core on :8760, web shell on :8761
Manual
cd core
python -m venv .venv
. .venv/bin/activate
pip install -e " .[dev] "
cp ../.env.example ../.env # fill in your keys (BYOK)
python -m pytest tests/ -v # run the suite
centri # start the server (or: python -m centri.cli)
The API listens on 127.0.0.1:8760 by default. Health checks:
curl localhost:8760/health
curl localhost:8760/status
Configuration
All configuration is environment-driven — copy .env.example to
.env and fill in your keys. Nothing is committed; .env and *.db are
gitignored. Centri is BYOK (bring your own keys) and model-agnostic:
Model gateway — LITELLM_BASE_URL / LITELLM_API_KEY point at any
OpenAI-compatible provider. Role models are set per task ( MODEL_INTENT ,
MODEL_REASONING , …).
Auth — CENTRI_AUTH_TOKEN gates every REST route except /health (and the
/events/stream WebSocket via ?token= ). Empty means auth off; set it before
exposing a port.
Tools — CENTRI_COMPOSIO_API_KEY enables Composio; CENTRI_COMPOSIO_TOOLS
is a comma-separated allowlist (default TAVILY_SEARCH ). With no key the
provider reports unavailable-with-reason and never touches the network.
Embeddings (optional) — off by default (lexical + FTS5 recall only). Turn
on semantic recall with CENTRI_EMBEDDING_* ; see .env.example .
Ingest paths — CENTRI_INGEST_OPENCODE_PATHS ,
CENTRI_INGEST_CLAUDE_CODE_PATHS , and CENTRI_INGEST_CURSOR_PATHS override the
per-platform probe for histories in unusual locations;
CENTRI_INGEST_DISABLED_AGENTS opts an agent out.
The reference deployment runs two services sharing one memory DB
( ~/.centri/state.db ):
sudo systemctl enable --now centri-core opencode
systemctl is-active centri-core opencode
curl -fsS http://127.0.0.1:8760/health # -> {"status":"ok",...}
Full deployment guide (systemd, Caddy/TLS, ports): docs/DEPLOY.md .
The xdg-open ENOENT line in the opencode.service logs is harmless — OpenCode
tries to auto-open a browser on a headless box and keeps running after the spawn
fails.
Centri also ships as a Hermes memory.provider . A thin plugin
( CentriMemoryProvider ) translates Hermes memory calls into the core's HTTP API:
prefetch → POST /memory/recall , sync_turn / on_memory_write → batched
POST /events/import . The deployable plugin lives at
deploy/hermes-plugin/centri/ .
# ~/.hermes/config.yaml
memory :
provider : centri
centri :
api_base : http://127.0.0.1:8760
auth_token : <CENTRI_AUTH_TOKEN> # same value as the core + the fork's CENTRI_TOKEN
Restart Hermes after any plugin change. Full guide:
docs/HERMES-INTEGRATION.md .
A fresh install imports your existing coding-agent histories so memory is
complete from day one. Discover what's available, then bootstrap once (idempotent
— re-running imports nothing new):
curl localhost:8760/ingest/discover # "found N OpenCode messages, M Cursor sessions"
curl -X POST localhost:8760/ingest/bootstrap \
-H ' content-type: application/json ' -d ' {} ' # one-time full import
After bootstrap, the ambient tail keeps pulling new history each scheduler tick.
centri-bench is a falsifiable head-to-head: each engine assembles a per-turn
brief from the same seeded history, scored on brief completeness, re-proposal
rate, next-step correctness, and stale-fact handling. Native Centri is run
against a real Letta server (v0.16.8, pgvector archival) — both a
deterministic rubric and an LLM judge agree Centri wins, the gap entirely on
stale-fact supersession.
Letta is a benchmark comparison only, not a runtime dependency — Centri runs
without it. Run it yourself:
cd core
python -m centri.bench.run # human-readable report
python -m centri.bench.run --json # machine-readable scores
Methodology: docs/centri-bench.md ; raw results:
docs/bench-results/ .
The core, memory graph, curation, ACP coding loop, tool contract, and history
ingest are covered by a 385-test suite ( cd core && python -m pytest tests/ ).
The unit suite runs green offline; integration tests that need a live core, BYOK
model keys, the opencode binary, or a real Letta server are environment-gated
and skip cleanly when those are absent.
The OpenCode fork web app builds and typechecks. The roadmap lives in
docs/ROADMAP.md ; the full docs index is
docs/README.md .
Centri is licensed under the MIT License — see LICENSE . The
OpenCode app shell in packages/opencode/ is a fork of the
MIT-licensed OpenCode project; its
upstream license is preserved at LICENSE-OPENCODE and the
original upstream README is kept at upstream/opencode/ .
Runtime dependencies are permissively licensed (MIT/BSD/Apache-2.0);
letta-client is bench-only tooling, not a product dependency.
Memory-native OpenCode fork with Centri core, event-spine memory, verbatim recall, and Hermes integration.
MIT, MIT licenses found
Licenses found
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
