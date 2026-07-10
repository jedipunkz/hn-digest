---
source: "https://corvin-labs.com/"
hn_url: "https://news.ycombinator.com/item?id=48857754"
title: "CorvinOS – self-hosted OS for AI agents, compliance built into the runtime"
article_title: "CorvinOS — The Operating System for AI Organizations"
author: "shumway"
captured_at: "2026-07-10T10:14:46Z"
capture_tool: "hn-digest"
hn_id: 48857754
score: 1
comments: 0
posted_at: "2026-07-10T09:33:19Z"
tags:
  - hacker-news
  - translated
---

# CorvinOS – self-hosted OS for AI agents, compliance built into the runtime

- HN: [48857754](https://news.ycombinator.com/item?id=48857754)
- Source: [corvin-labs.com](https://corvin-labs.com/)
- Score: 1
- Comments: 0
- Posted: 2026-07-10T09:33:19Z

## Translation

タイトル: CorvinOS – AI エージェント用のセルフホスト OS、ランタイムに組み込まれたコンプライアンス
記事のタイトル: CorvinOS — AI 組織のためのオペレーティング システム
説明: CorvinOS は、オープンソースの AI オペレーティング システムです。すべてのチャネルにエージェントをデプロイし、パイプラインを実行し、A2A 経由で組織を接続し、EU AI 法に準拠し、完全に自己ホストされます。

記事本文:
CorvinOS — AI 組織のためのオペレーティング システム
CorvinOS
プラットフォーム
コンプライアンス
コンソール
コンピューティング
ワークフロー
価格設定
クイックスタート
スター
Apache-2.0 · セルフホスト型エージェント OS
エージェントOS
あなたが実際に所有しています。
設計上EUに準拠。
AI エージェント用の自己ホスト型オペレーティング システム。エンジン、ペルソナ、ワークフロー、ランタイムで偽造されたツールを実行し、すべてのメッセンジャーでユーザーに対応し、実際のブラウザを駆動し、すべてのハッシュ チェーン監査を維持します。コンプライアンスは、ボルトオンではなくアーキテクチャに組み込まれています。
GitHub でスターを付ける
クイックスタート→
0
メッセージングブリッジ
∞
エンジンが検出されました
0
コンプライアンスの枠組み
0%
自己ホスト型
コマンドセンター
チャットは OS 全体を実行します
チャットはデモ ウィンドウではなく、シェルです。ペルソナ、エンジン、ワークフロー、Forge、RAG、エージェント コンピューティング、および A2A メッシュはすべて 1 つのプロンプトからアクセスできます。また、音声でエンドツーエンドで操作することもできます。リクエストを話すと、CorvinOS がそれを文字に起こし、ルーティングし、作業を行って、ユーザーの聞き方に合わせて答えを読み上げます。また、コンソールを開く必要さえありません。すべてのメッセージング ブリッジは完全なコントロール サーフェスであるため、Telegram、WhatsApp、または Signal スレッドから OS 全体を直接実行できます。
マイクを叩いて話します。音声はデフォルトでローカルに文字起こしされます。音声は返された瞬間に削除され、メタデータのみが監査チェーンに影響します。
1 つのプロンプトがすべてのレイヤーに到達します
すべてのコンプライアンス メカニズムは構造設計上の制約です。誤ってオフにしてしまう「コンプライアンス モード」はありません。
EU AI 法第 2 条50
ボットの開示
ユーザーごとに 1 回限りの AI の性質の開示。構造的にロックされています。構成によって無効にすることはできません。
トランスクリプト共有に対するデフォルトの同意を拒否します。ユーザーごと、TTL 制限があり、消費時に再検証されます。
GDPR アート。 30＆32
ハッシュチェーン監査
すべてのイベントは SHA に追加されます

256連チェーン。改ざんすると無効になります。オフライン検証可能。
EU AI 法第 2 条14
データの所在地
コンプライアンス ゾーン ルーティングと出力ロックダウンにより、データが許可されたゾーンから出ることが構造的に防止されます。
適応するコンピューティング
あらゆる問題に
それが Agentic Compute です。ワーカーがループを所有し、モデルがフレームを所有します。エージェントは何を最適化し、いつ停止するかを指示します。サンドボックス化されたワーカーは問題に適合する戦略を選択し、簡単な調整から 8 GB データセットのベイジアン検索まで、コンテキスト ウィンドウ内の生のバイトを決して使用せずに方向転換します。
ワーカー内で N 回の反復が実行されます。モデルは結果が返されるまでアイドル状態になります。
グリッド内のすべての組み合わせを評価します。送信時に事前生成 — 決定的で完全に再現可能。
連続範囲からのバッチごとの均一なサンプル。実行ごとに 1 つのシードが、再現性のために保存されます。
ガウス プロセス サロゲートは、最も高い期待改善ポイントを選択し、最小限の実行で良好な領域を見つけます。
十分な答えが得られるまで動作し続けます
マネージャー エンジンはタスクを計画し、並列ワーカーの実行に部分を委任し、すべての結果をスコアリングして、損失目標に向かって反復します。再帰のポイントは適応です。機能が欠落している場合、ワーカーは実行時に必要なツールやスキルを鍛えます。そのため、同じフローでモデルを調整したり、バックテストを実行したり、プロンプトから直接 8 GB データセットを分析したりできます。
ITER1
列に並んだ
損失 0.82
W-A ★ サブマネージャ ↳ スポーン SW-B1 · hermes SW-B2 · hermes W-C ⚒ forge.create_tool()
ITER2
列に並んだ
損失 0.51
W-D W-E W-F ✦ skill.create()
ITER3
列に並んだ
損失 0.12
W-G W-H W-Iは鍛造ツール＋スキルを使用
★サブマネージャー・再帰
⚒ 鍛造ツール · 実行中、共有
✦ スキル · 次のターンに注入
マネージャーのソネット · 労働者の俳句 · 機密のエルメス
Forge と SkillForge はループ内で実行されます。労働者

forge.create_tool() を実行すると、後のすべてのワーカーがツールを即座に利用できるようになります。別の関数は skill.create() を呼び出し、スキルは次のターンに挿入されます。そのランタイム生成が適応エンジンです。つまり、平易な言葉での質問が、コンテキスト ウィンドウには大きすぎるデータに対する実用的な分析にどのように変換されるのかを示します。
管理者があなたのデータに触れることはありません。すべてのスポーンはデータ ゲートを通過するため、縮小する予算の範囲内でクリーンで編集された概要のみがツリーに逆流されます。
データセットを一度登録します。同じファイルが 2 つの方法で提供されます。編集されたスナップショットがモデルに、実際のバイトがサンドボックスに提供されます。
→ 22 文字の不透明な data_handle
→ スキーマ + 編集されたサンプル、上限は 4,000 トークン
→ HyperLogLog および P5/P95 による統計 – 決して生の極端な値ではありません
サンドボックスツールへ
→ bwrap に読み取り専用でバインドされた実際のファイル パス
→ すべての行が読み取り可能 — pandas.read_csv(path) は正常に機能します
→ DSI プロトコルを介した構造化クエリと SQL ソース
すでに実行しているデータベースを指します。モデルは DSI プロトコルを介して名前でデータベースをクエリします。生の行は移動しません。
結果を説明してください。全体を 1 つのファイルとして送信します。
ボックスや矢印は決して描画しません。結果を平易な言葉で説明します。マネージャーが実行を委任し、実行グラフがそれ自体を発見します。そのグラフを反復可能なワークフローとして保存し、必要なものすべて (ロジック、ツール、スキル、ペルソナ、データ バインディング) を単一のポータブルな .awpkg バンドルにエクスポートします。
送信元
支払う
crm
ミスター
セグ
参加する
チャーン
fcst
タイル
#フィン
csv
ACSワークフローグラフ・分岐→3出力
›
3
ワークフローとして保存
毎週のミスターブリーフ
cron・月曜9:00
1 フェッチストライプ MRR
2 チャーン要因を分析する
3フォーマットまとめ
4 届ける → #ファイナンス
繰り返し可能、スケジュール可能、リプレイ可能
DAG リプレイ
実行で検出された正確なグラフをフリーズします — 同じノード、サム

注文 — 忠実な再実行の場合。
グラフをパラメーターを含む再利用可能なブループリントに一般化します。新しい入力を指定して、再度実行します。
エージェント コンピューティング委任ツリー全体 (サブマネージャー、ワーカー、およびその反復ループ) が 1 つの .awpkg としてキャプチャされ、スイープを再実行できるようになります。
ワークフローに必要なものすべてが 1 つのアーカイブに
.awpkg は署名された ZIP であり、レシピだけでなく機能全体が含まれます。 1 つのコマンドで任意のテナントにインストールすると、ツール、スキル、ペルソナ、委任ロジックがすべて同時に表示されます。宣言します。ランタイムはインストールされます。スクリプトやバイナリはなく、1 バイトが抽出される前に 8 回のフェールクローズ チェックが行われます。
完全な AWP トポロジ (エージェント ステップ、並列ファンアウト/ファンイン、再帰的 delegation_loop ノード) に加えて、トリガーと配信ターゲット。
実行によって構築されたすべてのカスタム ツールは、マニフェストが許可しない限り、ネットワーク: 拒否を使用してサンドボックスに固定されます。
再利用可能なスキルボディ。インストール前にリントされ、ワー​​カーの将来のターンに挿入されます。
ワーカーが引き受ける役割 - そのため、パッケージ化されたチームはどのマシンでも同じように動作します。
RAG プロバイダーとデータソースの参照は移動します。シークレット名は宣言されますが、その値はボールト内に残り、パッケージ内には決して含まれません。
あなたのデータセンター。組織全体が尋ねるだけです。
データセンター内の 1 台のマシンで CorvinOS を実行します。これはメッシュ内の別のノードにすぎず、機密データに対して高負荷のエージェント コンピューティングをたまたま実行するノードです。すべてのチームは A2A 経由で他のチームとペアになります。ハブもセンターもありません。誰もが平易な言語でタスクを送信し、回答が戻ってきますが、生データが建物の外に出ることはありません。
あなた社員・アシスタント
エンジニアリング acme-eng · spawn ✓
Analytics acme-data · spawn ✓
acme-ops のサポート · 受信箱
マーケティング acme-mktg · spawn ✓
リーガルアクメ - リーガル・オブザーバー
◆ データセンター
CorvinOS · エージェント コー

切断する
GPU-1
GPU-2
バッチ
機密 · エルメス · ゼロエグレス
⚿
外部請負業者の範囲内 · 10 rpm · 内部キャップ · スポーン ✗
● でのリクエスト
● 答えを出す
● チーム ↔ チームメッシュ
⚿ HMAC ペアリング · 同意ゲート型
ピアツーピア · 中央サーバーなし
すべての接続はフレンドシップ トークン、つまり帯域外で交換される共有 HMAC キーです。真ん中にプラットフォームはありません。各ノードが他のノードを認証し、独自のハッシュ チェーン上のすべてのエンベロープを監査します。
CorvinOS は、準拠する固定の製品ではありません。ユーザーの好みの作業方法を学習し、実行中に不足している機能を構築し、プラグインする必要があるものをすべてきれいに開きます。
Voice & Profile は、語彙、専門用語のレベル、深さなど、各リスナーの話し方を調整し、ペルソナはあらゆる場面で適切な音声、ツール、エンジンにルーティングします。オーバーライドは常に勝ちます。
ギャップに遭遇すると、ワーカーは実行時またはタスクの途中でツールを作成し、スキルを作成します。それぞれは永続的で、その後のターンですぐに利用できるため、システムは使えば使うほど機能が向上していきます。
カタログから外部 MCP サーバーをインストールしたり、拡張 API を介してベンダー レイヤーを追加したり、カスタム コンピューティング エンジンをドロップしたりできます。 corvin.* コアは暗号的にロックされたままです。拡張機能はガードレールを追加できますが、ガードレールを弱めることはできません。
本番環境の AI エージェントのためのすべて
CorvinOS は、マルチチャネル メッセージングからランタイム ツールの生成まで、すぐに使える運用グレードのインフラストラクチャを出荷します。
Discord、Telegram、WhatsApp、Slack、電子メール、Teams、Signal のネイティブ デーモン。ホットリロード設定、チャットごとのプロファイル、レート制限。
Forge は、4 つのスコープのワークスペース階層を持つ MCP を介して実行時にスキーマ バインドされたサンドボックス ツールを生成します。
SkillForge は、自動採点、昇進ゲート、インジェクション リンターを使用して、将来のターンにマークダウン スキルを注入します。
スワップベット

Claude Code、Codex、OpenCode、ローカルの Hermes/Ollama、GitHub Copilot を使用 — ブリッジの変更はありません。アダプティブ Haiku/Sonnet モデル選択機能が組み込まれています。
PII 編集されたインデックス作成と GDPR Art を使用した FTS5 SQLite リコール。 17 /forget による消去。
DSI プロトコルを通じて SQL ストアとベクター ストアを登録します。モデルはそれらを名前でクエリします。生の行がそのコンテキストに入ることはありません。
ライブを視聴し、機密性の高いアクションを承認しながら、実際の Chromium を操作して、ナビゲート、入力、クリック、読み取りを行います。ピクセルではなく要素インデックスによって機能します。シークレットはボールトから取得され、モデルや監査ログに到達することはありません。
カタログから外部 MCP サーバーをインストールするか、プロジェクト、テナント、またはユーザーごとにスコープ設定された拡張 API を介してベンダー レイヤーを配布します。 corvin.* コアは暗号的にロックされたままになります。拡張機能は追加しますが、決して弱めることはありません。
すべてのアクション (モデルの呼び出し、ツールの実行、メッセージ、ブラウザーのクリック) はハッシュチェーンされたログに記録され、毎日検証されます。メタデータのみ: テキスト、入力された秘密、または生の行を決して転写しないでください。
同じランタイム (ブリッジ、音声、RAG、ワークフロー、エージェント コンピューティング、データ ファイアウォール、A2A メッシュ) は、誰が質問するかによってまったく異なって表示されます。ラップトップを使用する 1 人のユーザーから組織全体まで。
エージェントは 1 人、メッセンジャーは 1 人
チャットに常駐するパーソナル アシスタント
Telegram、WhatsApp、Signal、または Web からテキストまたは音声で 1 人のエージェントと会話します。すべてのチャネルにわたってあなたのペルソナとプロフィールを伝えるので、常にあなたのアシスタントのように聞こえます。
プライベートRAG
自分のメモやファイルを尋ねる
文書を検索して、情報源を伴う根拠のある回答を得ることができます。すべてが自分のマシン上で自己ホストされます。サードパーティの受信箱や SaaS にデータが渡されることはありません。
音声ファースト
メモを送信して応答を聞く
音声メッセージを発すると、好みに合わせて音声による応答が得られます。スピーチは文字起こしです

デフォルトではローカルに保存され、音声は返されると同時に削除されます。
ビジネスデータ・平易な言語
倉庫に質問して問い合わせます
すでに実行している Postgres または Snowflake を接続します。チームは平易な言葉で質問し、サンドボックスは実際のバイトで動作し、編集され監査された回答のみが返されます。生の行は決して残されません。
ワークフローとパッケージ化
繰り返しの作業を自動化する
ガイド付き自然言語で複数ステップの AWP パイプラインを設計し、オンデマンドまたはスケジュールに従って実行し、全体を 1 つの .awpkg としてエクスポートして共有または再実行します。
漏れないサポート
接地され記録された顧客に応答する
エージェントは、お客様自身のドキュメントに基づいてサポート チャネルを処理します。編集は常にオンになっており、すべてのアクションは実際に検証できるハッシュ チェーン監査ログに書き込まれます。
組織メッシュ・A2A
すべてのチームが 1 つのファブリックで実行されます
法務部門、サポート部門、およびデータ部門はそれぞれ独自のエージェントを実行し、HMAC 署名付きエンベロープを介して相互に委任します。ハブも中央所有者もありません。ピアはスコープ付きのオブザーバーまたはエグゼキューターです。
社内エージェントコンピューティング
組織全体がデータセンターに問い合わせるだけです
データセンター ノードは 100 回の反復スイープを実行し、機密データに対してチューニングとバックテストを実行します。全員がリクエストを送信すると、応答が返されます。従業員がクランクを回している間、データは保存され、モデルはアイドル状態のままです。
コンプライアンスと常駐
構造によって管理され、設計によって拡張可能
EU-AI-Act と GDPR 管理は構造的なローカル専用エンジンです

[切り捨てられた]

## Original Extract

CorvinOS is the open-source AI operating system. Deploy agents on every channel, run pipelines, connect orgs via A2A, EU AI Act compliant — fully self-hosted.

CorvinOS — The Operating System for AI Organizations
CorvinOS
Platform
Compliance
Console
Compute
Workflows
Pricing
Quick Start
Star
Apache-2.0 · The self-hosted agentic OS
The agentic OS
you actually own.
EU-compliant by design.
A self-hosted operating system for AI agents: it runs the engines, personas, workflows and runtime-forged tools, meets your users on every messenger, drives a real browser, and keeps a hash-chained audit of it all — with compliance built into the architecture, not bolted on.
Star on GitHub
Quick Start →
0
Messaging bridges
∞
Engines detected
0
Compliance frameworks
0 %
Self-hosted
The command center
The chat runs the whole OS
The chat isn't a demo window — it's the shell. Personas, engines, workflows, the Forge, RAG, agentic compute and the A2A mesh are all reachable from one prompt. And you can drive it end‑to‑end by voice: speak a request, CorvinOS transcribes it, routes it, does the work, and speaks the answer back — tuned to how you listen. And you don't even need the console open: every messaging bridge is a full control surface, so you can run the whole OS straight from a Telegram, WhatsApp or Signal thread.
Hit the mic and talk. Speech is transcribed locally by default — the audio is deleted the moment it returns, and only metadata ever touches the audit chain.
One prompt reaches every layer
Every compliance mechanism is a structural design constraint. There is no "compliance mode" you can accidentally leave off.
EU AI Act Art. 50
Bot Disclosure
One-time AI-nature disclosure per user, structurally locked. Cannot be disabled via configuration.
Deny-by-default consent for transcript sharing. Per-user, TTL-capped, re-validated at consume time.
GDPR Art. 30 & 32
Hash-Chained Audit
Every event appends to a SHA-256-linked chain. Tampering invalidates it. Offline-verifiable.
EU AI Act Art. 14
Data Residency
Compliance-zone routing and egress lockdown structurally prevent data from leaving the permitted zone.
Compute that adapts
to any problem
That's Agentic Compute : the worker owns the loop, the model owns the framing. The agent says what to optimise and when to stop; a sandboxed worker picks the strategy that fits the problem and turns the crank — from a quick tune to a Bayesian search over an 8 GB dataset, and never a raw byte in the context window.
N iterations run in the worker — the model is idle until the result returns.
Evaluates every combination in the grid. Pre-generated at submit — deterministic and fully reproducible.
Uniform samples per batch from continuous ranges. One seed per run, stored for reproducibility.
A Gaussian-Process surrogate picks the highest Expected-Improvement points — finds a good region in the fewest runs.
It keeps working until the answer is good enough
A manager engine plans the task, delegates pieces to parallel worker runs, scores every result, and iterates toward a loss target. The point of the recursion is adaptation : where a capability is missing, a worker forges the tool or skill it needs at runtime — so the same flow can tune a model, run a backtest, or analyse an 8 GB dataset straight from a prompt.
ITER 1
queued
loss 0.82
W-A ★ sub-mgr ↳ spawns SW-B1 · hermes SW-B2 · hermes W-C ⚒ forge.create_tool()
ITER 2
queued
loss 0.51
W-D W-E W-F ✦ skill.create()
ITER 3
queued
loss 0.12
W-G W-H W-I uses forged tool + skill
★ sub-manager · recursion
⚒ forge tool · mid-run, shared
✦ skill · injected next turn
manager sonnet · workers haiku · confidential hermes
Forge and SkillForge run inside the loop . A worker calls forge.create_tool() and the tool is instantly available to every later worker; another calls skill.create() and the skill is injected on the next turn. That runtime generation is the adaptation engine — how a plain-language question turns into a working analysis over data far too big for any context window.
The manager never touches your data. Every spawn passes the data gate, so only clean, redacted summaries flow back up the tree — bounded by a shrinking budget.
Register a dataset once. The same file is presented two ways — a redacted snapshot to the model, the real bytes to the sandbox.
→ A 22-char opaque data_handle
→ Schema + redacted sample, capped at 4,000 tokens
→ Stats via HyperLogLog & P5/P95 — never raw extremes
To the sandbox tool
→ The real file path, read-only-bound into bwrap
→ Every row readable — pandas.read_csv(path) just works
→ Structured queries & SQL sources via the DSI protocol
Point it at the databases you already run — the model queries them by name over the DSI protocol, and the raw rows never move.
Describe the outcome. Ship the whole thing as one file.
You never draw the boxes and arrows. You describe an outcome in plain language — the manager delegates the run, and the execution graph discovers itself . Save that graph as a repeatable workflow, then export everything it needs — logic, tools, skills, personas and data bindings — into a single portable .awpkg bundle.
src
pay
crm
mrr
seg
join
churn
fcst
tile
#fin
csv
ACS workflow graph · branch → 3 outputs
›
3
Saved as a workflow
weekly-mrr-brief
cron · Mon 9:00
1 fetch Stripe MRR
2 analyse churn drivers
3 format summary
4 deliver → #finance
repeatable · schedulable · replayable
DAG Replay
Freezes the exact graph the run discovered — same nodes, same order — for a faithful re-run.
Generalises the graph into a reusable blueprint with parameters — point it at new inputs and run again.
A whole agentic-compute delegation tree — sub-managers, workers and their iteration loops — captures as one .awpkg , ready to re-run the sweep.
Everything the workflow needs, in one archive
An .awpkg is a signed ZIP — the whole capability, not just the recipe. Install it into any tenant with one command and the tools, skills, personas and delegation logic all light up together. It declares ; the runtime installs — no scripts, no binaries, eight fail-closed checks before a single byte is extracted.
The full AWP topology — agent steps, parallel fan-out/fan-in and recursive delegation_loop nodes — plus triggers and delivery targets.
Every custom tool the run built, each pinned to a sandbox with network: deny unless the manifest grants it.
Reusable skill bodies, linted before install and injected into the workers' future turns.
The roles the workers assume — so the packaged team behaves the same on any machine.
RAG-provider and datasource references travel; secret names are declared but their values stay in the vault — never in the package.
Your datacenter. The whole org just asks.
Run CorvinOS on one machine in your datacenter — it's just another node in the mesh, the one that happens to do the heavy agentic compute over your confidential data. Every team pairs with every other over A2A: no hub, no center. Anyone sends plain-language tasks, answers flow back, but the raw data never leaves the building.
You employee · assistant
Engineering acme-eng · spawn ✓
Analytics acme-data · spawn ✓
Support acme-ops · inbox
Marketing acme-mktg · spawn ✓
Legal acme-legal · observer
◆ DATACENTER
CorvinOS · agentic compute
gpu-1
gpu-2
batch
CONFIDENTIAL · hermes · zero egress
⚿
External contractor scoped · 10 rpm · INTERNAL cap · spawn ✗
● request in
● answer out
● team ↔ team mesh
⚿ HMAC-paired · consent-gated
peer-to-peer · no central server
Every connection is a Friendship Token — a shared HMAC key exchanged out-of-band. No platform sits in the middle; each node authenticates the other and audits every envelope on its own hash chain.
CorvinOS isn't a fixed product you conform to. It learns how you like to work, it builds the capabilities it's missing while it runs, and it opens cleanly to whatever else you need to plug in.
Voice & Profile tunes how each listener is spoken to — vocabulary, jargon level, depth — and personas route every turn to the right voice, tools and engine. Your overrides always win.
Hit a gap and a worker forges the tool and writes the skill — at runtime, mid-task. Each one persists and is instantly available to every later turn, so the system keeps getting more capable the more you use it.
Install external MCP servers from the catalog, add vendor layers through the Extension API, or drop in custom compute engines. The corvin.* core stays cryptographically locked — an extension can add a guardrail, never weaken one.
Everything for production AI agents
From multi-channel messaging to runtime tool generation, CorvinOS ships production-grade infrastructure out of the box.
Native daemons for Discord, Telegram, WhatsApp, Slack, Email, Teams and Signal. Hot-reload settings, per-chat profiles, rate limiting.
Forge generates schema-bound, sandboxed tools at runtime via MCP with a four-scope workspace hierarchy.
SkillForge injects markdown skills into future turns, with automatic grading, promotion gates and an injection linter.
Swap between Claude Code, Codex, OpenCode, local Hermes/Ollama and GitHub Copilot — no bridge changes. Adaptive Haiku/Sonnet model selection built in.
FTS5 SQLite recall with PII-redacted indexing and GDPR Art. 17 erasure via /forget .
Register SQL and vector stores through the DSI protocol. The model queries them by name — raw rows never enter its context.
Drive a real Chromium — navigate, fill, click, read — while you watch live and approve sensitive actions. It acts by element index, not pixels; secrets come from the vault and never reach the model or the audit log.
Install external MCP servers from the catalog, or ship vendor layers through the Extension API — scoped per project, tenant or user. The corvin.* core stays cryptographically locked; extensions add, never weaken.
Every action — model call, tool run, message, browser click — lands on a hash-chained log that's verified daily. Metadata only: never transcript text, typed secrets or raw rows.
The same runtime — bridges, voice, RAG, workflows, agentic compute, the data firewall and the A2A mesh — shows up very differently depending on who is asking. From one person on a laptop to a whole organization.
One agent, every messenger
A personal assistant that lives in your chats
Talk to a single agent from Telegram, WhatsApp, Signal or the web — by text or voice. It carries your persona and profile across every channel, so it always sounds like your assistant.
Private RAG
Ask your own notes and files
Point retrieval at your documents and get grounded answers with sources. Everything is self-hosted on your own machine — no third-party inbox, no data handed to a SaaS.
Voice-first
Send a note, hear an answer
Fire off a voice message and get a spoken reply tuned to how you like to listen. Speech is transcribed locally by default and the audio is deleted the instant it returns.
Business data · plain language
Query your warehouse by asking
Connect the Postgres or Snowflake you already run. The team asks in plain language, the sandbox works on the real bytes, and only redacted, audited answers come back — raw rows never leave.
Workflows & packaging
Automate the recurring work
Design a multi-step AWP pipeline in guided natural language, run it on demand or on a schedule, and export the whole thing as one .awpkg to share or re-run.
Support that never leaks
Answer customers, grounded and logged
An agent handles your support channels, grounded in your own docs, with redaction always on and every action written to a hash-chained audit log you can actually verify.
Organization mesh · A2A
Every team run on one fabric
Legal, support and data each run their own agent and delegate to one another over HMAC-signed envelopes — no hub, no central owner. Peers are scoped Observer or Executor.
In-house agentic compute
The whole org just asks the datacenter
A datacenter node runs 100-iteration sweeps, tuning and backtests over confidential data. Everyone sends a request and gets an answer back — the data stays put, the models stay idle while workers turn the crank.
Compliance & residency
Governed by construction, extensible by design
EU-AI-Act and GDPR controls are structural, local-only engines

[truncated]
