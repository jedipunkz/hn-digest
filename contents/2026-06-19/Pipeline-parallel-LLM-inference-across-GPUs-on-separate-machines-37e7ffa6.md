---
source: "https://github.com/leyten/shard"
hn_url: "https://news.ycombinator.com/item?id=48602121"
title: "Pipeline-parallel LLM inference across GPUs on separate machines"
article_title: "GitHub - leyten/shard: Pipeline-parallel LLM inference across GPUs on separate machines. · GitHub"
author: "ngaut"
captured_at: "2026-06-19T21:30:54Z"
capture_tool: "hn-digest"
hn_id: 48602121
score: 1
comments: 0
posted_at: "2026-06-19T19:14:34Z"
tags:
  - hacker-news
  - translated
---

# Pipeline-parallel LLM inference across GPUs on separate machines

- HN: [48602121](https://news.ycombinator.com/item?id=48602121)
- Source: [github.com](https://github.com/leyten/shard)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T19:14:34Z

## Translation

タイトル: 別々のマシン上の GPU にわたるパイプライン並列 LLM 推論
記事のタイトル: GitHub - leyten/shard: 別々のマシン上の GPU にわたるパイプライン並列 LLM 推論。 · GitHub
説明: 別々のマシン上の GPU にわたるパイプライン並列 LLM 推論。 - ライテン/シャード

記事本文:
GitHub - leyten/shard: 別々のマシン上の GPU にわたるパイプライン並列 LLM 推論。 · GitHub
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
ライテン
/
シャード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ T

ags ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
63 コミット 63 コミット docs docs Phase0 Phase0 Research Research シャード シャード サイドカー サイドカー .gitignore .gitignore ライセンス ライセンス通知 NOTICE README.md README.md STATE.md STATE.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
別々のマシン上の GPU にわたるパイプライン並列 LLM 推論。モデルもね
単一カードの場合は大きいが、レイヤーの連続したブロックに分割されます - 1 つのシャード
GPU ごと — リクエストは、シャードを介したストリーミング アクティベーションによって処理されます。
注文。データセンター、単一のホスト、ノードがモデル全体を保持することはありません。
Shard は c0mpute の推論エンジンです。
GLM-5.2 (744B)、オープン インターネット経由で 7 つの分散したプロシューマ GPU にまたがる
7,440 億パラメータのフロンティア モデル、7 か所で最大 30 トーク/秒でサービス
米国 6 つの州のプロシューマー向け Blackwell GPU - WAN 経由、貪欲、決定的。
GLM-5.2 (NVFP4、78 レイヤー) は、6 台の RTX PRO 6000 にわたってノードごとに 13 レイヤーに分割されています。
それを保持できるのは 1 枚のカードではなく、6 枚のカードです。各ノードは、自身のブロックのみをロードします。コーディネーター
モデル層は保持せず、トークンの埋め込み/ヘッドと小さな CUDA グラフのみを保持します。
トークンを提案する GLM-4-9B ドラフト。分散型 744B が検証します。
実行ごとに、個別の GPU UUID / パブリック IP / リージョンなどの検証可能なレシートが発行されます。
測定された WAN エッジ RTT (22 ～ 75 ミリ秒)、出力トークン ハッシュ、およびロスレス最適化
チェックしてください。この実行のレシート: docs/receipts/glm52-nvfp4-wan-20260618.json
(懐疑論者がどのようにチェックするかについては、docs/PROOF.md を参照してください)。
これが、この 1 行での論文全体です。フロンティアサイズのモデルは、誰にとっても大きすぎるのです。
単一カード、異なるネットワーク上のマシン間でサービス - 複数のアクティベーション
実際に使用可能な速度で、横断するたびに国を訪問します。
WAN 経由のプレーン パイプライン デコードは遅延の影響を受けやすい

d: トークンごとに 1 往復、~1 ～ 2
tok/s、使用不可。 30 へのパスは、それぞれがコミットされた一連の測定ステップでした。
重要な洞察: WAN 経由では、コンピューティングではなく往復が希少なリソースです。
したがって、データセンターでは限界に達する投機的なデコードがゲーム全体になります。小さな
草案では K 個のトークンを提案しています。分散型 744B は単一のパイプラインでそれらを検証します。
横断。貪欲な受け入れは、検証されたプレフィックスをコミットします。次に、2 つの複利が得られます。
リング上の非同期パイプライン。リングは直帰なので複数回
チャンクが一度に飛行できることを確認します。コーディネーターは継続的なストリームの草案を作成します
そして、重複するチャンクを待機せずにパイプラインに送り込みます。そのため、ループは次の時点で実行されます。
パイプラインのレイテンシではなく、スループットです。以前はすべてを支配していた WAN
試行すると、ループの最大 5% に低下します。
CUDA グラフ化されたドラフト。 WAN が非表示になると、GLM-4-9B ドラフト (単一トークン)
デコード、起動オーバーヘッドバウンド) がループの 94% になります。 CUDAグラフとしてキャプチャする
3.8 倍 (49.7→13.1 ms/tok) に短縮されます。難しい部分は、静的 KV キャッシュを優先させることでした
グラフ キャプチャでの投機的ロールバック — 書き込みスロットを
静的アドレス位置テンソル。結果は、eager path とバイト同一であるため、
最適化はロスレスであることが証明されています。 ( Research/glm_swarm_nvfp4_cg.py ,
Research/glm_swarm_nvfp4_cg_diff.py 。)
トランスは層の積み重ねです。シャードはスタックを連続したブロックに分割します。
GPUごとに1ブロック。トークンは、次のブロックを介してアクティベーションを渡すことによって生成されます。
注文;各ノードは、独自のレイヤーの KV キャッシュを保持します。
コーディネーター (WA) ── GLM-4-9B ドラフト (CUDA グラフ化) + embed / lm_head
│
§─► stage0 ─► stage1 ─► stage2 ─► stage3 ─► stage4 ─► stage5 ─┐ (チャンクの検証、パイプライン化)
│ NV TX (・) MN MOUT │
│ 0–

12 13–25 26–38 39–51 52–64 65–77 │
━───── ダイレクトリターン（テール→コーディネーター、1ホップ） ────┘
コーディネーター (エントリ ノード) は 744B レイヤーを保持せず、ドラフトとシン レイヤーのみを保持します。
運転手。各ラウンド: ドラフトでは K 個のトークンが提案されます。コーディネーターが船 [cur, d₁..dₖ]
それらを埋め込むステージ 0 に。チェーンは 1 回の前方走査ですべての K+1 を検証します。
テールは argmaxes をコーディネーターに直接返します (中継されずに 1 ホップ)。
コーディネーターは、一致する最長のプレフィックスを貪欲に受け入れます。そのようなチャンクがたくさんあります
フライト (パイプライン) を一度に実行し、ドラフトはキャプチャされた CUDA グラフを
静的 KV キャッシュ。
同じ場所にある GPU 間でモデルを分割することはよく理解されています。マシン間で実行する
オープンなインターネットでは、使用できるほどの速度はありませんが、それがシャードの部分です
所有しています。
待ち時間。すべてのトークンはパイプライン全体を通過します。投機的なデコードは償却されます
多くのコミットされたトークンを 1 往復します。パイプライン処理はトラバーサルと重なるため、
WAN はフロアではなくなります。 CUDA でグラフ化されたドラフトにより、残った部分は安価に保たれます。
輸送。アクティベーション テンソルは、ステップごとにパブリック インターネットを通過します。シャード
このレイヤーを所有しています — フェイルファストで再接続する監視対象エッジ、エッジごとの健全性
ログ記録、不透明な「壊れたパイプ」はありません。ワイヤは認証され、暗号化されます。
pickle-free フレーミング (phase0/wire.py ; 共有 SHARD_PSK 下の ChaCha20-Poly1305)、
したがって、受動的な観察者は何も学習せず、偽造されたフレームはコードではなく解析エラーです。
実行。 (NAT ホールパンチ + ホーム ルーターのリレー フォールバックが残ります)
フェーズ 1 の作業。現在は直接オープンポートが機能しています。)
シャードは計算インフラストラクチャであり、次の 3 つの保証が適用されます。
無修正。エンジンはモデルをそのまま実行します。推論パスにコンテンツ フィルターがありません。
12月

集中した。誰でも 1 つのコマンドで GPU に参加し、ブロックを割り当てることができます。
層。中央の推論サーバーはありません。
プライベート。モデル全体を保持するノードはありません。ストーリー全体ではなく、本当のスタートです。の
ワイヤは密封されているため（認証暗号化、ピクルスフリー）、漏洩はありません。
パス;ただし、参加ノードはその層を実行するために復号化する必要があるため、
処理するアクティベーション。中間のアクティベーションでも、まだ一部がリークする可能性があります。
ユーザーのトークンを悪意のあるノードに送信します。計画 - 漏洩境界層を信頼できるものに固定する
ノード、リクエストごとの信頼できるルーティング、決してオーバークレームしない - です
docs/ARCHITECTURE.md 。これは最大の未解決の問題であり、
一つとして扱われます。
gpt-oss-120B (WAN 経由で約 40 tok/s) — パーミッションレス ビルド ターゲット
米国のさまざまな州に点在する 3 つの RTX 4090 で 120B (MXFP4、36 レイヤー) +
コーディネーター、~40 tok/s (ピーク ~42)、貪欲、正確。これは許可のないリグです
作業 (フェーズ 3+) は、プレーンな 24 GB の消費者向けカード、実際のボランティアが実行するハードウェアに基づいて構築されています。
この実行の検証可能な受信 (個別の GPU UUID/IP/状態、WAN エッジ RTT、出力)
ハッシュ、同期対パイプライン トークンの一致): docs/receipts/gpt-oss-120b-wan-20260619.json 。
レイテンシー限界 ~18 tok/s からの上昇を各ステップで測定:
最後のステップは誰も探していないステップです。システム内で最も安価なノードです。
レイヤーレスコーディネーター — 群れから 1 大陸離れた場所に座っており、2 つの長い料金を支払っていました
すべてのトークンの往復。それをステージの隣、同じ分散ノード上に置きます。
無料で遅延が最大 40% 削減されました。全記録:
docs/research/wan-speculative-decoding.md 。
GLM-5.2 (上) は依然としてフロンティアサイズのフラッグシップであり、パラメータは 744B の 6 倍です。
gpt-oss-120B は、ネットワークがブートストラップされる、より高速なコンシューマ カード ビルド ターゲットです。
フェーズ0/トランスポート + デプロイ: Wire.py (シールドされたフレーム)

ng)、mesh.py (エッジ RTT)、
proof_receipt.py (run-receipt build/verify)、起動 + ベンチ ツール
Research/ swarm ドライバー — glm_swarm_nvfp4_kv.py (NVFP4 KV キャッシュ ステージ)、
glm_swarm_nvfp4_pipe.py (パイプライン仕様デコード)、glm_swarm_nvfp4_cg.py
(CUDA グラフ化されたドラフト)、*_cg_diff.py / *_fwdcmp.py (正確性診断)
docs/ ARCHITECTURE、ROADMAP、PROOF.md、領収書/、および調査記録
シャード/エンジン モジュールの足場 (ノード、トランスポート、specdec、トポロジ)
ロードマップ
フェーズ 0 — 輸送、実証済み。多段スプリットによる確実なサーブ。
フェーズ 1 — WAN。 NAT の背後にあるさまざまなネットワーク: ホールパンチ、リレー フォールバック、
活性化量子化、エッジ監視。
フェーズ 2 — 投機的なデコード。 swarm 上でのドラフトと検証 — で行われます
GLM-5.2 744B スケール、WAN 経由で ~30 tok/s グリーディ (上記の ~18 ～ 25 の gpt-oss-120B)。
フェーズ 3 — 許可のない群れ。ワンコマンド結合、動的レイヤー割り当て
異種 GPU 全体、トークンごとの支払い、フォールト トレランス。
詳細、合否基準、リスクについては、 docs/ROADMAP.md を参照してください。
Apache ライセンス 2.0 © 2026 leyten
別々のマシン上の GPU にわたるパイプライン並列 LLM 推論。
読み込み中にエラーが発生しました。このページをリロードしてください。
17
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Pipeline-parallel LLM inference across GPUs on separate machines. - leyten/shard

GitHub - leyten/shard: Pipeline-parallel LLM inference across GPUs on separate machines. · GitHub
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
leyten
/
shard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
63 Commits 63 Commits docs docs phase0 phase0 research research shard shard sidecar sidecar .gitignore .gitignore LICENSE LICENSE NOTICE NOTICE README.md README.md STATE.md STATE.md pyproject.toml pyproject.toml View all files Repository files navigation
Pipeline-parallel LLM inference across GPUs on separate machines. A model too
large for any single card is split into contiguous blocks of layers — one shard
per GPU — and a request is served by streaming activations through the shards in
order. No datacenter, no single host, and no node ever holds the whole model.
Shard is the inference engine for c0mpute .
GLM-5.2 (744B) across seven scattered prosumer GPUs, over the open internet
A 744-billion-parameter frontier model, served at ~30 tok/s across seven
prosumer Blackwell GPUs in six US states — over WAN, greedy, deterministic.
GLM-5.2 (NVFP4, 78 layers) is split 13 layers per node across 6× RTX PRO 6000;
no single card holds it, six do. Each node loads only its own block . A coordinator
holds no model layers — just the token embedding/head and a small CUDA-graphed
GLM-4-9B draft that proposes tokens, which the distributed 744B verifies.
Every run emits a verifiable receipt — distinct GPU UUIDs / public IPs / regions,
measured WAN edge RTTs (22–75 ms), the output token hash, and a lossless-optimization
check. This run's receipt: docs/receipts/glm52-nvfp4-wan-20260618.json
(see docs/PROOF.md for how a skeptic checks it).
That is the whole thesis in one line: a frontier-size model, far too big for any
single card, served across machines on different networks — activations crossing
the country on every traversal — at a speed that is actually usable.
Plain pipeline decode over WAN is latency-bound: one round-trip per token, ~1–2
tok/s, unusable. The path to 30 was a sequence of measured steps, each committed:
The key insight: over WAN the round-trip is the scarce resource, not compute —
so speculative decoding, marginal in a datacenter, becomes the whole game. A small
draft proposes K tokens; the distributed 744B verifies them in a single pipeline
traversal; greedy acceptance commits the verified prefix. Then two compounding wins:
Async pipelining over the ring. Because the ring is direct-return, multiple
verify chunks can be in flight at once. The coordinator drafts a continuous stream
and pumps overlapping chunks into the pipeline without waiting — so the loop runs at
the pipeline's throughput , not its latency . The WAN, which dominated every prior
attempt, drops to ~5% of the loop.
CUDA-graphed draft. Once the WAN is hidden, the GLM-4-9B draft (single-token
decode, launch-overhead-bound) becomes 94% of the loop. Capturing it as a CUDA graph
cuts it 3.8× (49.7→13.1 ms/tok). The hard part was making the static KV cache honor
speculative rollback under graph capture — solved by driving the write slot through a
static-address position tensor; the result is byte-identical to the eager path , so
the optimization is provably lossless. ( research/glm_swarm_nvfp4_cg.py ,
research/glm_swarm_nvfp4_cg_diff.py .)
A transformer is a stack of layers. Shard splits the stack into contiguous blocks,
one block per GPU. A token is produced by passing activations through the blocks in
order; each node keeps a KV-cache for its own layers.
coordinator (WA) ── GLM-4-9B draft (CUDA-graphed) + embed / lm_head
│
├─► stage0 ─► stage1 ─► stage2 ─► stage3 ─► stage4 ─► stage5 ─┐ (verify chunks, pipelined)
│ NV TX (·) MN MO UT │
│ 0–12 13–25 26–38 39–51 52–64 65–77 │
└──────────────── direct return (tail → coordinator, 1 hop) ────┘
The coordinator (entry node) holds no 744B layers — only the draft and a thin
driver. Each round: the draft proposes K tokens; the coordinator ships [cur, d₁..dₖ]
into stage 0, which embeds them; the chain verifies all K+1 in one forward traversal;
the tail returns the argmaxes straight to the coordinator (one hop, not relayed back);
the coordinator greedy-accepts the longest matching prefix. Many such chunks are in
flight at once (the pipeline), and the draft replays a captured CUDA graph against a
static KV cache.
Splitting a model across co-located GPUs is well understood. Doing it across machines
on the open internet, fast enough to be usable, is not — and that is the part Shard
owns.
Latency. Every token traverses the whole pipeline. Speculative decoding amortizes
one round-trip over many committed tokens; pipelining overlaps the traversals so the
WAN stops being the floor; the CUDA-graphed draft keeps what's left cheap.
Transport. The activation tensor crosses the public internet on every step. Shard
owns this layer — supervised edges that fail fast and reconnect, per-edge health
logging, no opaque "broken pipe." The wire is authenticated and encrypted with
pickle-free framing ( phase0/wire.py ; ChaCha20-Poly1305 under a shared SHARD_PSK ),
so a passive observer learns nothing and a forged frame is a parse error, not code
execution. (NAT hole-punching + relay fallback for home routers is the remaining
Phase 1 work; a direct open port stands in today.)
Shard is c0mpute infrastructure, held to its three guarantees:
Uncensored. The engine runs models as-is. No content filter in the inference path.
Decentralized. Anyone can join a GPU with one command and be assigned a block of
layers. No central inference server.
Private. No node holds the whole model — a real start, not the whole story. The
wire is sealed (authenticated encryption, pickle-free), so the leak is not on the
path; but a participating node must decrypt to run its layer, so it sees the
activations it processes. Intermediate activations can still leak a fraction of a
user's tokens to a malicious node. The plan — pin leaky boundary layers to trusted
nodes, per-request trusted routing, never overclaim — is in
docs/ARCHITECTURE.md . It is the number-one open problem and is
treated as one.
gpt-oss-120B at ~40 tok/s over WAN — the permissionless build target
120B (MXFP4, 36 layers) across 3 scattered RTX 4090s in different US states + a
coordinator, ~40 tok/s (peak ~42), greedy, exact . This is the rig the permissionless
work (Phase 3+) is built on — plain 24GB consumer cards, the hardware a real volunteer runs.
This run's verifiable receipt (distinct GPU UUIDs / IPs / states, WAN edge RTTs, output
hash, sync-vs-pipelined token match): docs/receipts/gpt-oss-120b-wan-20260619.json .
The climb from a latency-bound ~18 tok/s, each step measured:
The last step is the one nobody looks for: the cheapest node in the system — the
layer-less coordinator — was sitting a continent away from the swarm, paying two long
round-trips on every token. Putting it next to the stages, on the same scattered nodes,
was a ~40% latency cut for free. Full record:
docs/research/wan-speculative-decoding.md .
GLM-5.2 (above) remains the frontier-size flagship — 6× the parameters at 744B;
gpt-oss-120B is the faster, consumer-card build target the network is bootstrapped on.
phase0/ transport + deploy: wire.py (sealed framing), mesh.py (edge RTTs),
proof_receipt.py (run-receipt build/verify), launch + bench tooling
research/ the swarm drivers — glm_swarm_nvfp4_kv.py (NVFP4 KV-cached stages),
glm_swarm_nvfp4_pipe.py (pipelined spec-decode), glm_swarm_nvfp4_cg.py
(CUDA-graphed draft), *_cg_diff.py / *_fwdcmp.py (correctness diagnostics)
docs/ ARCHITECTURE, ROADMAP, PROOF.md, receipts/, and the research records
shard/ engine module scaffolding (node, transport, specdec, topology)
Roadmap
Phase 0 — Transport, proven. Reliable serving through a multi-stage split.
Phase 1 — WAN. Different networks behind NAT: hole-punching, relay fallback,
activation quantization, edge supervision.
Phase 2 — Speculative decoding. Draft-and-verify over the swarm — done at
GLM-5.2 744B scale, ~30 tok/s greedy over WAN (and gpt-oss-120B at ~18–25, above).
Phase 3 — Permissionless swarm. One-command join, dynamic layer allocation
across heterogeneous GPUs, per-token payouts, fault tolerance.
Full detail, pass/fail criteria, and risks: docs/ROADMAP.md .
Apache License 2.0 © 2026 leyten
Pipeline-parallel LLM inference across GPUs on separate machines.
There was an error while loading. Please reload this page .
17
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
