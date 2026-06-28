---
source: "https://github.com/Adirdabush1/cerberus"
hn_url: "https://news.ycombinator.com/item?id=48704458"
title: "Cerberus – a local firewall for AI agents' tool calls"
article_title: "GitHub - Adirdabush1/cerberus: 🐺 Cerberus — local-first security gateway for AI coding agents. Intercept, risk-score & human-approve every tool call (Claude Code, Codex, Cursor, Cline). · GitHub"
author: "cerberussec"
captured_at: "2026-06-28T05:28:50Z"
capture_tool: "hn-digest"
hn_id: 48704458
score: 2
comments: 0
posted_at: "2026-06-28T04:49:30Z"
tags:
  - hacker-news
  - translated
---

# Cerberus – a local firewall for AI agents' tool calls

- HN: [48704458](https://news.ycombinator.com/item?id=48704458)
- Source: [github.com](https://github.com/Adirdabush1/cerberus)
- Score: 2
- Comments: 0
- Posted: 2026-06-28T04:49:30Z

## Translation

タイトル: Cerberus – AI エージェントのツール呼び出し用のローカル ファイアウォール
記事のタイトル: GitHub - Adirdabush1/cerberus: 🐺 Cerberus — AI コーディング エージェント用のローカル ファースト セキュリティ ゲートウェイ。すべてのツール呼び出し (クロード コード、コーデックス、カーソル、クライン) をインターセプトし、リスク スコアを付け、人間が承認します。 · GitHub
説明: 🐺 Cerberus — AI コーディング エージェント用のローカル ファースト セキュリティ ゲートウェイ。すべてのツール呼び出し (クロード コード、コーデックス、カーソル、クライン) をインターセプトし、リスク スコアを付け、人間が承認します。 - アディルダブッシュ1/ケルベロス

記事本文:
GitHub - Adirdabush1/cerberus: 🐺 Cerberus — AI コーディング エージェント用のローカル ファースト セキュリティ ゲートウェイ。すべてのツール呼び出し (クロード コード、コーデックス、カーソル、クライン) をインターセプトし、リスク スコアを付け、人間が承認します。 · GitHub
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
アディルダブッシュ1
/
ケルベロス
公共
通知

オン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
32 コミット 32 コミット .github/ workflows .github/ workflows bin bin Brainstorms Brainstorms ダッシュボード ダッシュボードの例 例 パッケージ/ インジェクションモデル パッケージ/ インジェクションモデル ルール ルール スクリプト スクリプト スパイク スパイク src src .gitignore .gitignore LICENSE LICENSE PLAN.md PLAN.md README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md cerberus-demo.gif cerberus-demo.gif package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自律型 AI コーディング エージェント向けのローカル ファースト セキュリティ ゲートウェイ。ケルベロスはエージェントの間に座る
(Claude Code、Codex、Cursor、Cline) とマシンは、すべてのツール呼び出しを実行前にインターセプトします。
4 つのシグナルにわたってリスク スコアリングを行い、許可、監査、人間の承認を求める、ブロックのいずれかを行います。
それはすべてあなたのマシン上にあり、外部 API や何も外に出す必要はありません。
自律コーディング エージェントは、ユーザーに代わってシェル コマンドを実行し、ファイルを編集し、ネットワーク呼び出しを行います。
マシンの速度が低下し、無人になることがよくあります。 1 つの悪いステップ (rm -rf、不要な git プッシュ、漏洩した .env、
エージェントを騙して機密情報を漏洩させる毒入り README)、そしてそのループに人間が関与することはありません。
やめてください。 Cerberus は、エージェントが実際に動作するツール境界にそのチェックポイントを置きます。
PreToolUse ─▶ インターセプト ─▶ ポリシー + 動作 + コンテンツ + インジェクション ─▶ リスク エンジン ─▶ 許可 · 監査 · HITL · ブロック
PostToolUse ─▶ 検査 ─▶ シークレット + インジェクション検出 ─▶ セッション汚染状態
4 つの決定論的シグナルが 1 つの重み付きリスク スコアに集約され、ハードフロアが設定されます。

絶対的な
禁止事項を無効にすることはできません。
🟢 シークレットの抽出 — コンテキストに読み込まれたシークレットを検出し、アウトバウンドのコンテンツと一致します。
payload : 実際にキー (生または Base64/hex/URL エンコード) を運ぶ呼び出しを、来歴とともに保持します。
(source: .env:4 · sha256:… · 97% ) シークレット自体をログに記録することはありません。
🟢 過剰な権限 — すべての通話がゲートされます。未知のツールはフェールクローズされました。機密パス ( ~/.ssh 、
~/.aws 、認証情報、 /etc/passwd ) が保持されます。破壊的なコマンド ( rm -rf 、 Remove-Item -Recurse 、
chmod 777 、 kill -9 ) ブロックまたは保留されました。
🟢 危険な出口 — 宛先ポリシー: 信頼できるホスト (レジストリ、GitHub、OpenAI/Anthropic)
自動的に許可されます。ペースト サイト / Webhook キャッチャー / Raw-IP 宛先が保持されます。
🟡 ツールの不正使用 — 暴走ループとツールの呼び出し率/繰り返しの検出。
🟡 プロンプトインジェクション — ツールの結果でインジェクションを検出し、次の出力をゲートします (ヒューリスティック)
分類子;オプションのローカル DeBERTa モデル)。 LLM プロンプトではなくツール呼び出しを認識するため、
インジェクション自体ではなく、インジェクション (出口) の悪用。
端末優先の承認 - 保留中の通話がエージェントのネイティブ許可プロンプトに表示されます (Claude Code /
カーソル)、または cerberus 経由で <id> / localhost ダッシュボードを承認します。
フォレンジック ダッシュボード — セッションごとのタイムライン、リスク要因の内訳、ステップを実行するリプレイ プレーヤー
セッションのリスクがどのように蓄積されたかを通じて。
マルチエージェント — 1 つのアダプター層がクロード コード、コーデックス、カーソル、およびクラインに対応します。
データとしてのポリシー — ルールとリスク重みは、コードではなく編集可能な YAML です。
ローカルファースト — 127.0.0.1 にバインドします。外部 API やテレメトリはありません。シークレット値はディスクやログに触れることはありません。
npm i -g @cerberussec/core # または次のようにアドホックを実行します: npx @cerberussec/core <cmd>
# Cerberus をエージェントに接続します (エージェントの構成にマージします — バックアップ、idem)

潜在的):
cerberus init # クロード コード、プロジェクト レベル (--agent codex|cursor|cline、--global、--print)
# ゲートウェイ + ダッシュボードを開始します (1 つのプロセス):
cerberus エンジン # 次に http://127.0.0.1:9000/ を開きます
通常どおりエージェントを使用します。ツール呼び出しは Cerberus を介してルーティングされるようになりました。デフォルトでは、保留（HITL）コールは次のようになります。
ターミナル内で直接承認されました: Cerberus が ask を返すため、Claude コードはそのネイティブを示します
Cerberus の理由を含む許可プロンプト — セッションを離れることなく承認/拒否します。
ダッシュボード ( http://127.0.0.1:9000/ ) には、ライブ タブ (アクション センター + ストリーム) と
[セッション] タブ — リスク要因の内訳とリプレイを含むセッションごとのフォレンジック タイムライン
プレーヤーは、セッションのリスクがどのように蓄積されたかを段階的に確認できます。
Cerberus はエージェントの実行ループ内で実行されるため、ターミナルがリアルタイムの意思決定ポイントとなります。
そしてダッシュボードはさらに奥深いものです。重大度ごと (デフォルトは AG_APPROVAL_SURFACE=terminal ):
代わりに中央の Web キューを使用したいですか? AG_APPROVAL_SURFACE=dashboard を設定します - 通話を保留してから一時停止します
エンジンの同期保留と、ダッシュボード (または帯域外の端末) から承認/拒否を行います。
cerberus pending # レビューのために保留されている呼び出しをリストします (ID 付き)
cerberus accept < id > # 保留中の通話を解放します …
ケルベロスは < id > # … または拒否します
追加の端末アラートは制御端末 ( /dev/tty 、標準エラー出力にフォールバック) に書き込まれるため、
クロード コードへのプロトコル チャネルはクリーンなままです。 env 経由で調整します:
エンジン + シグナル + リスク + ダッシュボードはエージェントに依存しません。シン アダプタのみ (エージェントの解析)
フックイベント → 正規化 → 判定形状の出力) はエージェントごとに行われます。 cerberus init --agent <name> を使用して接続します。
codex/cursor/cline アダプターは、公開されているフック仕様に従います。インストールされているバージョンと照合して確認する
( cerberus init --agent <name> --print は正確な構成を示します)。 Roo コードはサポートされていません

テッド (2026 年アーカイブ)。
PreToolUse フック → /intercept は単一のハード強制ポイント (許可/拒否/要求、または HITL が保持します)
決定するまでソケットは開いたままにしておきます)。
PostToolUse フック → /inspect は観察専用です。セッションの汚染状態を更新します。
次のアクションは完全なコンテキストに基づいて判断されます。ツールの結果が変更されることはありません。
このエンジンは本質的にエージェントに依存しません。エージェントごとのアダプター ( --agent ) のみが異なります。
PreToolUse ─▶ /intercept ─▶ ポリシー + 動作 + コンテンツ/インジェクション ─▶ RiskEngine ─▶ ALLOW/AUDIT/HITL/BLOCK
PostToolUse ─▶ /inspect ─▶ シークレット検出 + インジェクション分類器 ─▶ セッション汚染状態
(監査ログ + WebSocket → ダッシュボード)
単一ノード + TypeScript パッケージ。ダッシュボードは、エンジンによって提供される Vite/React アプリです。ルールと
リスク ウェイトは編集可能な YAML データであり、コード ( rules/ ) ではありません。
Cerberus は、ツール境界上のランタイム ゲートウェイです。秘密漏洩では最強です
防止および許可チョークポイントとして。 (LLM プロンプトではなく) ツール呼び出しを認識するため、キャッチします。
インジェクション自体ではなく、プロンプト インジェクションの悪用はカバーされません。
データ パイプライン/RAG ポイズニング。 exfil 一致は信頼性が高いですが、密閉性がありません (新しいシークレット形式、
スプリットアクロスコールエンコーディング）。偽りの保証よりも誠実な債務不履行。
外部 API や API キーはなく、マシンからは何も出ません。オプションのインジェクションモデル
( @cerberussec/injection-model 、ProtectAI DeBERTa、Apache-2.0) のアップグレード
組み込みのヒューリスティック分類子。必要な場合にのみインストールしてください。コアは OSS クリーンです
(Apache/MIT互換のdeps); Meta Prompt-Guard は意図的にコア (Llama ライセンス) から除外されます。
# クローンから: インストール (ルート + ダッシュボードは別の npm プロジェクト) とビルド
npm インストール && npm --prefix ダッシュボード インストール
n

pm run build # エンジン (tsc → dist) + ダッシュボード (vite → ダッシュボード/dist) をコンパイルします
npm run エンジン # tsx 経由でソースから実行 (dev)
npm タイプチェックを実行する
npm 実行テスト:動作 && npm 実行テスト:コンテンツ && npm 実行テスト:インジェクション && npm 実行テスト:リスク \
&& npm 実行テスト:init && npm 実行テスト:プロジェクター && npm 実行テスト:監査 && npm 実行テスト:通知 \
&& npm 実行テスト:セキュリティ && npm 実行テスト:ポリシー && npm 実行テスト:アダプター
npm run e2e:behavioral && npm run e2e:content && npm run e2e:injection && npm run e2e:risk
マイルストーンとブレインストーミングについては、PLAN.md を参照してください。各決定の背後にある設計記録については、PLAN.md を参照してください。
🐺 Cerberus — AI コーディング エージェント向けのローカルファーストのセキュリティ ゲートウェイ。すべてのツール呼び出し (クロード コード、コーデックス、カーソル、クライン) をインターセプトし、リスク スコアを付け、人間が承認します。
Apache-2.0ライセンス
セキュリティポリシー
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
ケルベロス v0.1.0 🐺
最新の
2026 年 6 月 12 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

🐺 Cerberus — local-first security gateway for AI coding agents. Intercept, risk-score & human-approve every tool call (Claude Code, Codex, Cursor, Cline). - Adirdabush1/cerberus

GitHub - Adirdabush1/cerberus: 🐺 Cerberus — local-first security gateway for AI coding agents. Intercept, risk-score & human-approve every tool call (Claude Code, Codex, Cursor, Cline). · GitHub
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
Adirdabush1
/
cerberus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
32 Commits 32 Commits .github/ workflows .github/ workflows bin bin brainstorms brainstorms dashboard dashboard examples examples packages/ injection-model packages/ injection-model rules rules scripts scripts spike spike src src .gitignore .gitignore LICENSE LICENSE PLAN.md PLAN.md README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md cerberus-demo.gif cerberus-demo.gif package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A local-first security gateway for autonomous AI coding agents. Cerberus sits between the agent
(Claude Code, Codex, Cursor, Cline) and your machine, intercepts every tool call before it runs,
risk-scores it across four signals, and either allows, audits, asks for human approval, or blocks
it — all on your machine, with no external API and nothing leaving the box.
Autonomous coding agents run shell commands, edit files, and make network calls on your behalf — at
machine speed, often unattended. One bad step ( rm -rf , an unwanted git push , a leaked .env , a
poisoned README that tricks the agent into exfiltrating secrets) and there's no human in the loop to
stop it. Cerberus puts that checkpoint on the tool boundary , where the agent actually acts.
PreToolUse ─▶ intercept ─▶ Policy + Behavioral + Content + Injection ─▶ Risk Engine ─▶ ALLOW · AUDIT · HITL · BLOCK
PostToolUse ─▶ inspect ─▶ secret + injection detection ─▶ session contamination state
Four deterministic signals aggregated into one weighted risk score, with a hard floor that absolute
prohibitions can never override.
🟢 Secret exfiltration — detects secrets loaded into context, then content-matches the outbound
payload : holds the call that actually carries the key (raw or base64/hex/url-encoded), with provenance
( source: .env:4 · sha256:… · 97% ) and never logging the secret itself.
🟢 Excessive permissions — every call gated; unknown tools fail-closed; sensitive paths ( ~/.ssh ,
~/.aws , credentials, /etc/passwd ) held; destructive commands ( rm -rf , Remove-Item -Recurse ,
chmod 777 , kill -9 ) blocked or held.
🟢 Dangerous egress — destination policy: trusted hosts (registries, GitHub, OpenAI/Anthropic)
auto-allowed; paste sites / webhook catchers / raw-IP destinations held.
🟡 Tool abuse — runaway-loop and tool-call-rate/repetition detection.
🟡 Prompt injection — detects injection in tool results and gates the next egress (heuristic
classifier; optional local DeBERTa model). It sees tool calls, not the LLM prompt — so it catches the
exploitation of an injection (the egress), not the injection itself.
Terminal-first approval — held calls surface in the agent's native permission prompt (Claude Code /
Cursor), or via cerberus approve <id> / a localhost dashboard.
Forensic dashboard — per-session timeline, risk-factor breakdown, and a Replay player that steps
through how a session's risk built up.
Multi-agent — one adapter layer serves Claude Code, Codex, Cursor, and Cline.
Policy as data — rules and risk weights are editable YAML, not code.
Local-first — binds to 127.0.0.1 , no external API, no telemetry; secret values never touch disk or logs.
npm i -g @cerberussec/core # or run ad-hoc with: npx @cerberussec/core <cmd>
# wire Cerberus into your agent (merges into the agent's config — backed up, idempotent):
cerberus init # Claude Code, project-level (--agent codex|cursor|cline, --global, --print)
# start the gateway + dashboard (one process):
cerberus engine # then open http://127.0.0.1:9000/
Use your agent as usual — tool calls now route through Cerberus. By default a held (HITL) call is
approved right in the terminal : Cerberus returns ask , so Claude Code shows its native
permission prompt with Cerberus's reason — approve/deny without leaving your session.
The dashboard ( http://127.0.0.1:9000/ ) has a Live tab (Action Center + stream) and a
Sessions tab — a forensic timeline per session with a risk-factor breakdown and a Replay
player to step through how a session's risk built up.
Cerberus runs inside the agent's execution loop, so the terminal is the realtime decision point
and the dashboard is the deep dive. Per severity (default AG_APPROVAL_SURFACE=terminal ):
Prefer a central web queue instead? Set AG_APPROVAL_SURFACE=dashboard — held calls then pause on
the engine's synchronous hold and you Approve/Deny from the dashboard (or the terminal, out-of-band):
cerberus pending # list calls held for review (with their ids)
cerberus approve < id > # release a held call …
cerberus deny < id > # … or deny it
Extra terminal alerts write to the controlling terminal ( /dev/tty , falling back to stderr) so the
protocol channel to Claude Code stays clean. Tune via env:
The engine + signals + risk + dashboard are agent-agnostic; only a thin adapter (parse the agent's
hook event → normalize → emit its verdict shape) is per-agent. Wire one with cerberus init --agent <name> :
codex / cursor / cline adapters follow the published hook specs; verify against your installed version
( cerberus init --agent <name> --print shows the exact config). Roo Code is unsupported (archived 2026).
PreToolUse hook → /intercept is the single hard enforcement point (allow/deny/ask; or HITL holds the
socket open until you decide).
PostToolUse hook → /inspect is observe-only: it updates the session's contamination state so
the next action is judged with full context. It never modifies a tool result.
The engine is agent-agnostic at its core; per-agent adapters ( --agent ) are the only thing that differs.
PreToolUse ─▶ /intercept ─▶ Policy + Behavioral + Content/Injection ─▶ RiskEngine ─▶ ALLOW/AUDIT/HITL/BLOCK
PostToolUse ─▶ /inspect ─▶ secret detection + injection classifier ─▶ session contamination state
(audit log + WebSocket → dashboard)
Single Node + TypeScript package; the dashboard is a Vite/React app served by the engine. Rules and
risk weights are editable YAML data , not code ( rules/ ).
Cerberus is a runtime gateway on the tool boundary . It's strongest at secret-exfiltration
prevention and as a permission chokepoint. Because it sees tool calls (not the LLM prompt), it catches
the exploitation of a prompt injection — not the injection itself — and it does not cover
data-pipeline / RAG poisoning. The exfil match is high-confidence but not airtight (novel secret formats,
split-across-calls encoding). Honest defaults over false guarantees.
No external API, no API key, nothing leaves the machine. The optional injection model
( @cerberussec/injection-model , ProtectAI DeBERTa, Apache-2.0) upgrades
the built-in heuristic classifier; install it only if you want it. The core is OSS-clean
(Apache/MIT-compatible deps); Meta Prompt-Guard is deliberately kept out of core (Llama license).
# from a clone: install (root + dashboard are separate npm projects) and build
npm install && npm --prefix dashboard install
npm run build # compile the engine (tsc → dist) + dashboard (vite → dashboard/dist)
npm run engine # run from source via tsx (dev)
npm run typecheck
npm run test:behavioral && npm run test:content && npm run test:injection && npm run test:risk \
&& npm run test:init && npm run test:projector && npm run test:audit && npm run test:notify \
&& npm run test:security && npm run test:policy && npm run test:adapters
npm run e2e:behavioral && npm run e2e:content && npm run e2e:injection && npm run e2e:risk
See PLAN.md for milestones and brainstorms/ for the design records behind each decision.
🐺 Cerberus — local-first security gateway for AI coding agents. Intercept, risk-score & human-approve every tool call (Claude Code, Codex, Cursor, Cline).
Apache-2.0 license
Security policy
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Cerberus v0.1.0 🐺
Latest
Jun 12, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
