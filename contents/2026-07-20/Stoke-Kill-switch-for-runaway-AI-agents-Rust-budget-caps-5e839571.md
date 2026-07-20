---
source: "https://stokegate.com"
hn_url: "https://news.ycombinator.com/item?id=48983316"
title: "Stoke – Kill switch for runaway AI agents (Rust, budget caps)"
article_title: "Stoke — your agents, your machines, your rules"
author: "pawfromoz"
captured_at: "2026-07-20T19:25:37Z"
capture_tool: "hn-digest"
hn_id: 48983316
score: 1
comments: 0
posted_at: "2026-07-20T18:59:32Z"
tags:
  - hacker-news
  - translated
---

# Stoke – Kill switch for runaway AI agents (Rust, budget caps)

- HN: [48983316](https://news.ycombinator.com/item?id=48983316)
- Source: [stokegate.com](https://stokegate.com)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T18:59:32Z

## Translation

タイトル: Stoke – 暴走 AI エージェントのキル スイッチ (Rust、予算上限)
記事のタイトル: Stoke — エージェント、マシン、ルール
説明: AI エージェントを独自のハードウェアまたはクラウドにルーティングし、予算を超えるものは拒否し、暴走ループを排除します。プロバイダーが呼び出される前に実行されます。 Rust バイナリ 1 つ。

記事本文:
ストーク — エージェント、マシン、ルール
⬡ ストーク
執行
ルーティング
比較する
インストール
次は何ですか
ガイド
GitHub
ただ観察するのではなく、強制しましょう。
あなたのエージェント。あなたのマシン。
あなたのルール。
独自のハードウェア、または必要に応じてクラウドにルーティングします。予算を超えるものはお断りします。ループを数秒でキルします。お金が使われた後に報告するダッシュボードではなく、お金が使われる前に決定する 1 つの ~5.5 MB Rust バイナリ。
任意の OpenAI 互換エージェントで動作します。base_url を localhost:8787/v1 に指定します。
Claude Code は ANTHROPIC_BASE_URL 経由で接続します。プレリリース、MIT、毎日ドッグフーディングされています。
午前 2 時にループを開始するエージェントは、誰かが起きるまで支出を続けます。請求アラートは受信通知であり、制御ではありません。損害が発生した後に通知されます。ダッシュボードには何が起こったかがグラフ化されます。ストークは何が起こるかを決定します。
再試行エージェントは、従量制キーを使用して、失敗した同じリクエストを一晩中喜んで再実行します。
支出通知には実際の使用量が記録されます。着火する頃には予算はすでになくなっている。
ダッシュボードが観察します。ストークは拒否する。
予算超過またはループしているリクエストはプロバイダーに到達しません。ゲートウェイで拒否され、チャートにフラグが立てられません。
プロバイダーがリクエストを確認する前にリクエストを拒否する制限であり、朝に人間が見つけるためのアラートではありません。
キーごとの使用上限 (USD)。 [[keys]] 構成テーブルで設定され (各キーは STOKE_API_KEYS にもリストされている必要があります)、起動時にバジェット ガードに適用されます。プロバイダーへの呼び出し前にチェックされます。上限を超えている場合、リクエストは 429 予算超過で拒否されます。このエラーはエージェントが処理でき、後のログ行ではありません。 /v1/budget でのキーごとのライブ費用。リクエストごとのコストは、非ストリーミング応答の stoke_cost フィールドに反映されます。ストリーミングされた応答にも費用が発生します。Stoke は、バイトが通過する際にプロバイダーが報告する使用量を読み取り、課金します。

ストリームが終了するか、クライアントがハングアップしたとき。何も報告しない従量制のプロバイダーは見積もりに基づいて請求され、1 としてフラグが立てられ、0 ドルになることはありません。
2 層サーキット ブレーカー: 正確なプロンプト ハッシュ マッチングとオプトイン セマンティック類似性。そのため、それ自体を言い換えるループもループとしてカウントされます。 60 秒以内に同様のリクエストが 5 回発生すると、キーが 120 秒間ブロックされます。現在、しきい値はグローバルな定数となっており、キーごとの調整がロードマップに含まれています。
API キーごとにスライドする 60 秒のウィンドウ。同じ [[keys]] テーブルで rate_limit_rpm として設定され、起動時に適用されます。不正行為を行ったエージェントの 1 人が拒否されました (429)。他の人はみんな働き続けます。
API キーが設定されておらず、開発フラグも設定されていない場合、すべてのリクエストが拒否されます。 /health を除くすべてのエンドポイントには認証が必要です。モデル インベントリとノード負荷は偵察データであるため、ステータス エンドポイントも含まれます。フェデレーション ポーリングでも認証されます。構成が間違っているということは、オープンではなくクローズされていることを意味します。
独自のモデルを実行します。あなたが所有するすべてのマシンで。
Stoke は、実際の MacBook + Mac Studio + Ollama Cloud セットアップで毎日ドッグフードを行っています。それを Ollama ノードに向けると、ノードをポーリングするため、ノードを認識しているかのようにルーティングされます。
バックグラウンド レジストリは、各ノードの /api/tags (プルされたモデル) と /api/ps (RAM にロードされたモデル) をポーリングします。配置ではウォーム ノードが優先され、ライブ インフライト カウントによってマシン間で同じモデルのバランスがとられ、レイテンシ EWMA との関係が解消され、自動的にフェイルオーバーされます。到達不能なノードは、戻ってくるまで除外されます。
すべてのライブ ストリームがそれを教えてくれます。最初のトークンまでの時間とトークン/秒はノードごとにモデルごとに測定され、コンテキスト ウィンドウとツール呼び出しサポートは /api/show から提供されます。つまり、予測は推測ではなくハードウェアから得られます。
連邦: ストークの後ろのストーク
プロバイダー タイプ「stoke」では、1 つのゲートウェイが別のゲートウェイの前面に配置されます。ピアのモデル、ウォームステート、ライブロード、および測定された速度

/v1/node を通じてインポートされます。生の Ollama がレポートできるよりも豊富な状態です。転送されたリクエストは x-stoke-hop を伝送します。深さは正確に 1 であり、意図的な A-to-B-to-A サイクルはループできません。出荷されたテスト スクリプトがそれを証明しています。
実質的な勝利: Ollama は 127.0.0.1 にバインドされたままになります。 Stoke は唯一のネットワークに面したドアであり、認証されます。
推測ではなく最適化する自動ルーティング
エージェントに自動車モデルを指示すると、Stoke が推定コスト、予測レイテンシ、および指定された優先順位に基づいて対象となるすべての (モデル、ノード) ペアをスコアリングし、選択します。自動では低コストのウェイトを使用し、自動では高速のウェイトを使用します。モデルは構成と検出されたノードからのみ取得されます。 Stoke はモデル名なしで出荷されます。
雰囲気ではなく、厳密な除外: モデルのコンテキスト ウィンドウに適合しないプロンプトと、構造化されたツール呼び出しを発行できないモデルを対象としたツール呼び出しリクエストは、レシートに理由が記載されて除外されます。
オプトイン ( hedge = true ): 2 台のマシン上でモデルが温存されている小さなプロンプトは、両方のマシンで起動されます。最初にプレフィルを通過した方が勝ち、敗者はドロップされます。アイドル状態のハードウェアがあなたのために競争します。重複したローカル コンピューティングにより、単一ノードでは実現できないテール レイテンシが発生します。
マージンゼロのノードのみ - ヘッジによってクラウドの請求額が 2 倍になることはありません。
あらゆる決断にはその理由が示されており、選ばれなかった道もある
自動ルーティングされた応答には、選択された候補と推定コストと予測レイテンシ、すべての選択肢とその敗因、そして正直な反事実、つまり構成されたクラウド モデルが請求するであろう定価の完全な領収書が含まれます。 GET /v1/budget は、それらを実行中の台帳に集約します。見積もりを明確にラベル付け。決して品質を主張するものではありません。
すべての呼び出しはパイプラインを通過します。
これは、1 つのリクエストに対する実際の操作の順序です。プロバイダーの呼び出し前のすべての段階で拒否することができ、そのうち 4 つは拒否される可能性があります。

書くのはあなたです。
あなたのステージでは、任意の言語の Webhook を使用できます。プレーンな HTTP エンドポイントはプラグインです。 JS/TS プラグインは、コンパイル時フラグの背後にある組み込み V8 上で実行されます。ブラック ボックスではありません。プロバイダー呼び出しより上位のすべての段階で no と言うことができ、プロンプトを変換する段階は、プロンプトがマシンを離れる前に実行されます。
従量制の API キーに対してコーディング エージェント (Claude Code、OpenCode、aider、Codex) を実行すると、1 回の悪いループで朝の 4 桁が発生します。 Stoke は、エージェントとキーの間にループ キル スイッチ、レート制限、ハード USD キャップを設けます。ループのキル スイッチとレート制限は、ストリーミングを含むすべてのリクエストに適用されるため、ループがどのように通信するかに関係なく、暴走ループは停止されます。現在、OpenCode はエンドツーエンドで機能します。 OpenAI 互換エージェントは、base_url を介して接続します。
すべてのハードウェアに 1 つのエンドポイント。
ここには MacBook、あちらには Mac Studio、おそらく GPU を備えた Linux ボックスがあります。 Stoke は、ウォームアウェアな配置、マシン間での同じモデルのライブ ロード バランシング、自動フェイルオーバー、ゲートウェイ間のフェデレーションなど、すべてを実現します。一方、Ollama は LAN に公開されることはありません。
AI 製品のコントロール プレーン。
AI 製品を出荷しており、使用権の適用が必要です。請求チームが信頼できる、テナントごとの厳しい上限がリクエスト パスで適用されます。これは、すべての顧客の LLM 支出に対するコントロール プレーンです。ストークはここに向かっていますが、まだ出荷されていません。デザインパートナーを募集しています。
正直な範囲: サブスクリプション プランは割り当ての世界です。
エージェントがサブスクリプション プラン (Claude Max または Pro、ChatGPT Plus) で実行されている場合、リクエストごとのドル価格はないため、ドルの上限はありません。 Stoke は引き続きそのトラフィックをレート制限し、その中のループを強制終了できますが、ドルで制限することはできません。また、そうでないふりをするつもりもありません。 USD のハードキャップは従量制 API キーに適用され、すべてのリクエストに価格が設定されます。
あなたのプロム

ポイントはインフラストラクチャから離れることはありません。
データの保存場所は契約によるものではなく、アーキテクチャによるものです。 EU ベースのプロジェクトで、データ処理協定を読んでもそれを必要としない人々のために構築されました。
リクエストは所有するマシンにルーティングされます。明示的に構成しない限り、クラウド プロバイダーには何も送信されません。一度構成すると、自動フェイルオーバーでそれが使用される可能性があるため、プロンプトがマシンから決して出てはいけない場合は、単純にクラウド層を構成しないでください。
Ollama は 127.0.0.1 にバインドされたままになります。 Stoke は唯一のネットワークに面したサーフェスであり、/health を除くすべてのエンドポイントには認証が必要です。
モデル インベントリとノード負荷は偵察データであるため、ステータス エンドポイントも認証されます。キーが設定されていないということは、サービスがないこと、つまりサービスが開いていないことを意味します。
ストークが勝つところ。そしてそうでないところ。
LiteLLM、Portkey、Helicone は優れたツールです。これらははるかに多くのプロバイダーをカバーしており、可観測性についてさらに深く掘り下げています。 Stoke は、執行という別の仕事のために作られています。
この表には意図的にベンチマークの数値がありません。パフォーマンスに関する主張を公開する場合、最初に再現可能なハーネスが同梱されて出荷されます。
1 行でエージェントに指示します。
macOS と Linux の静的バイナリは、メインのすべてのコミットからビルドされ、夜間のローリング ビルドに公開されます。Rust ツールチェーンは必要ありません。正直なところ: そのローリング ビルドが現在存在するものです。最初のタグ付き安定版リリースはまだ先であり、インストーラーは、出荷されていないプラットフォーム上のソースからのコンパイルにフォールバックします。
自分で構築したいですか? git clone と Cargo build --release — プラットフォームに一致するバイナリがない場合、インストーラーはまさにそれを行います。
最小限の stoke.toml — 1 つのローカル Ollama ノード、デフォルト ポート 8787:
次に、エージェントをプロバイダーではなく Stoke に向けます。OpenAI 互換ツールと Claude Code は両方とも機能します。
フェールクローズのデフォルトに注意してください: STOKE_API_KEYS (または STOKE_ なし)

DEV=1（ローカル実験の場合）、Stoke はすべてのリクエストを拒否します。それが特徴です。
実際にヒットしたものに基づいて構築されています。
ストークは、ロードマップに基づいて約束されたものではなく、最初にドッグフード化され、その地位を獲得したときに機能を出荷します。次に何が構築されるかは、実際のユーザーが遭遇するものによって決まります。 Stoke があなたのために機能するまであと 1 つの能力である場合、それが私たちに教えていただける最も有益な情報です。
私たちが検討している未解決の質問 — あなたの質問があれば、検討してください:
現在、フェイルオーバーは別のマシンで同じモデルを再試行します。ロールが連鎖的に落ちた場合、クラウド モデルが 429 または停止に達し、ベンダー間でもトラフィックがローカル モデルに低下しますか?どのように動作させたいかを教えてください。
AI 製品のテナントごとの上限
リクエスト パスでの、顧客ごとのハード使用権の強制。価格階層に歯止めが必要な場合は、スレッドを開始してください。
施行ベンチマーク、OpenTelemetry エクスポート、人間的↔ローカル変換 - すべてがテーブルにあります。どれがあなたのブロックを解除し、上に上がるかを教えてください。
このページの言葉を鵜呑みにしないでください。
リポジトリには、再現可能なテスト ハーネスが同梱されています。自分のマシンで実行して、主張が成立するか成立しないかを観察してください。
2 つのモック Ollama ノードに対して 1 つのゲートウェイ — Ollama のインストールは必要ありません: モデルの検出、ウォーム プレースメント、ストリーミング インフライト アカウンティング、フェイルオーバー、ヘルス。
2 つのゲートウェイが意図的なサイクルに配線されています: インベントリのインポート、クロスゲートウェイ ルーティング、ホップ ガード、サイクルの安全性。
実行中に、GET /v1/nodes (すべてのルーティングの決定が行われるライブ レジストリ) をポーリングします。
施行ベンチマーク (オーバーヘッド、トリップまでの時間、誤検知率) は、公開されている再現可能なハーネスに同梱されています。自分で実行できる場合にのみ数字を表示します。
エージェントと請求書の間にファイアウォールを置きます。
オープンソース、MIT ライセンス、サインアップやロックインはありません。プレリリースでオープンに組み込まれています。
データ プレーンは MIT ライセンスを取得しており、そのまま残ります。

MITよ永遠に。それが変わった場合は、フォークしてください。それがライセンスのポイントです。

## Original Extract

Route AI agents to your own hardware or the cloud, refuse anything over budget, kill runaway loops — enforced before any provider is called. One Rust binary.

Stoke — your agents, your machines, your rules
⬡ Stoke
Enforcement
Routing
Compare
Install
What's next
Guides
GitHub
Enforce, don't just observe.
Your agents. Your machines.
Your rules.
Route to your own hardware — or the cloud, when you choose. Refuse anything over budget. Kill loops in seconds. One ~5.5 MB Rust binary that decides before the money is spent, not a dashboard that reports it after.
Works with any OpenAI-compatible agent — point base_url at localhost:8787/v1 .
Claude Code connects via ANTHROPIC_BASE_URL . Pre-release · MIT · dogfooded daily.
An agent that starts looping at 2 a.m. keeps spending until someone wakes up. The billing alert is a receipt, not a control — it fires after the damage. Dashboards chart what happened. Stoke decides what is allowed to happen.
A retrying agent will happily replay the same failing request all night, on your metered key.
Spend notifications trail actual usage. By the time one fires, the budget is already gone.
Dashboards observe. Stoke refuses.
An over-budget or looping request never reaches the provider. Refused at the gateway, not flagged in a chart.
Limits that reject requests before a provider ever sees them — not alerts for a human to find in the morning.
Per-key spend caps in USD, set in a [[keys]] config table (each key must also be listed in STOKE_API_KEYS ) and applied to the budget guard at startup. Checked before any provider call: over the cap, the request is refused with 429 Budget exceeded — an error your agent can handle, not a log line for later. Live spend per key at /v1/budget ; per-request cost rides along in a stoke_cost field on non-streaming responses. Streamed responses accrue spend too — Stoke reads the usage the provider reports as the bytes pass through, and charges when the stream ends or the client hangs up. A metered provider that reports nothing is billed from an estimate, flagged as one, never as $0.
A two-layer circuit breaker: exact prompt-hash matching plus opt-in semantic similarity, so a loop that rephrases itself still counts as a loop. Five similar requests within 60 seconds blocks the key for 120 seconds. Thresholds are global constants today — per-key tuning is on the roadmap.
A sliding 60-second window per API key, set as rate_limit_rpm in the same [[keys]] table and applied at startup. One misbehaving agent gets refused (429); everyone else keeps working.
No API keys configured and no dev flag: every request is rejected. Every endpoint except /health requires auth — status endpoints included, because model inventory and node load are reconnaissance data. Federation polling authenticates too. Misconfigured means closed, not open.
Run your own models. On every machine you own.
Stoke is dogfooded daily on a real MacBook + Mac Studio + Ollama Cloud setup. Point it at your Ollama nodes and it routes like it knows them — because it polls them.
A background registry polls each node's /api/tags (models pulled) and /api/ps (models loaded in RAM). Placement prefers warm nodes , balances the same model across machines by live in-flight count, breaks ties with a latency EWMA, and fails over automatically. Unreachable nodes are excluded until they come back.
Every live stream teaches it: time-to-first-token and tokens/sec are measured per model per node, and context windows plus tool-calling support come from /api/show — so predictions come from your hardware, not guesses.
Federation: Stoke behind Stoke
Provider type "stoke" lets one gateway front another. The peer's models, warm state, live load, and measured speeds are imported through its /v1/nodes — richer state than raw Ollama can report. Forwarded requests carry x-stoke-hop ; depth is exactly one, and a deliberate A-to-B-to-A cycle cannot loop — a shipped test script proves it.
The practical win: Ollama stays bound to 127.0.0.1 . Stoke is the only network-facing door, and it authenticates.
Auto-routing that optimizes, not guesses
Point your agent at the auto model and Stoke scores every eligible (model, node) pair on estimated cost, predicted latency, and your stated preference order — then picks. auto-cheap weights spend, auto-fast weights speed. Models come only from your config and your discovered nodes; Stoke ships with zero model names.
Hard exclusions, not vibes: prompts that don't fit a model's context window and tool-calling requests aimed at models that can't emit structured tool_calls are ruled out with the reason on the receipt.
Opt-in ( hedge = true ): small prompts whose model sits warm on two of your machines are fired at both — first past prefill wins, the loser is dropped. Your idle hardware races for you; duplicate local compute buys tail latency that no single node can.
Zero-marginal nodes only — hedging never doubles a cloud bill.
Every decision shows its reasoning — and the road not taken
Auto-routed responses carry the full receipt: the chosen candidate with estimated cost and predicted latency, every alternative and why it lost, and an honest counterfactual — the list price your configured cloud model would have charged. GET /v1/budget aggregates them into a running ledger. Estimates, clearly labeled; never a quality claim.
Every call passes through your pipeline.
This is the actual order of operations for one request. Every stage before the provider call can refuse it — and four of them are yours to write.
The yours stages take webhooks in any language — a plain HTTP endpoint is a plugin. JS/TS plugins run on an embedded V8 behind a compile-time flag. Nothing is a black box: every stage above the provider call can say no , and the ones that transform your prompt run before it leaves the machine.
You run coding agents — Claude Code, OpenCode, aider, Codex — against metered API keys, and one bad loop is a four-figure morning. Stoke puts a loop kill switch, a rate limit, and a hard USD cap between the agent and the key. The loop kill switch and rate limit apply to every request, streaming included — so a runaway loop is stopped no matter how it talks. OpenCode works end-to-end today; any OpenAI-compatible agent connects via base_url .
One endpoint for all your hardware.
A MacBook here, a Mac Studio there, maybe a Linux box with a GPU. Stoke fronts them all: warm-aware placement, live load balancing of the same model across machines, automatic failover, and federation between gateways — while Ollama never gets exposed to the LAN.
A control plane for your AI product.
You ship an AI product and need usage entitlement enforcement: hard per-tenant caps your billing team can trust, enforced at the request path — a control plane over every customer's LLM spend. This is where Stoke is headed, not shipped yet; we are looking for design partners.
Honest scope: subscription plans are quota-world.
If your agents run on a subscription plan — Claude Max or Pro, ChatGPT Plus — there is no per-request dollar price, so there is nothing to dollar-cap. Stoke can still rate-limit that traffic and kill loops in it, but it cannot cap it in dollars, and we won't pretend otherwise. Hard USD caps apply to metered API keys, where every request has a price.
Your prompts never leave your infrastructure.
Data residency by architecture, not by contract. An EU-based project, built for people who read data-processing agreements and would rather not need them.
Requests route to machines you own. Nothing goes to a cloud provider unless you explicitly configure one. Once configured, automatic failover may use it — so if prompts must never leave your machines, simply don't configure a cloud tier.
Ollama stays bound to 127.0.0.1 . Stoke is the only network-facing surface, and every endpoint except /health requires auth.
Model inventory and node load are reconnaissance data, so even the status endpoints authenticate. No keys configured means no service — not open service.
Where Stoke wins. And where it doesn't.
LiteLLM, Portkey and Helicone are good tools. They cover far more providers and go much deeper on observability. Stoke is built for a different job: enforcement.
No benchmark numbers in this table, on purpose. When we publish performance claims, they will ship with a reproducible harness first.
One line, then point your agent at it.
Static binaries for macOS and Linux are built from every commit on main and published to a rolling nightly build — no Rust toolchain needed. Honest status: that rolling build is what exists today; the first tagged stable release is still to come, and the installer falls back to compiling from source on any platform we don't ship.
Prefer to build it yourself? git clone and cargo build --release — the installer does exactly that when no binary matches your platform.
A minimal stoke.toml — one local Ollama node, default port 8787:
Then point any agent at Stoke instead of the provider — OpenAI-compatible tools and Claude Code both work:
Note the fail-closed default: without STOKE_API_KEYS (or STOKE_DEV=1 for local experiments), Stoke rejects every request. That is a feature.
Built on what you actually hit.
Stoke ships features when they earn their place — dogfooded first, not promised against a roadmap. What gets built next is driven by what real users run into. If Stoke is one capability away from working for you, that's the most useful thing you can tell us.
Open questions we're weighing — weigh in if one is yours:
Today failover retries the same model on another machine. Should a role fall through a chain — cloud model hits a 429 or an outage, traffic degrades to your local model, even across vendors? Tell us how you'd want it to behave.
Per-tenant caps for AI products
Hard usage-entitlement enforcement per customer, at the request path. If your pricing tiers need teeth, start a thread .
Enforcement benchmarks, OpenTelemetry export, Anthropic↔local translation — all on the table. Tell us which one unblocks you and it moves up.
Don't take this page's word for it.
The repo ships reproducible test harnesses. Run them on your own machine and watch the claims hold or break.
One gateway against two mock Ollama nodes — no Ollama install needed: model discovery, warm placement, streaming in-flight accounting, failover, health.
Two gateways wired into a deliberate cycle: inventory import, cross-gateway routing, hop guard, cycle safety.
While they run, poll GET /v1/nodes — the live registry every routing decision is made from.
Enforcement benchmarks — overhead, time-to-trip, false-positive rate — ship with a public reproducible harness. Numbers only when you can run them yourself.
Put a firewall between your agents and your bill.
Open source, MIT licensed, no signup, no lock-in. Pre-release and built in the open.
The data plane is MIT-licensed and stays MIT forever. If that ever changes, fork it — that's the point of the license.
