---
source: "https://github.com/thekiraproject/kira-project-site/blob/main/posts/2026-07-06-prefill-tax.md"
hn_url: "https://news.ycombinator.com/item?id=48822323"
title: "Local AI is re-reading its own prompt"
article_title: "kira-project-site/posts/2026-07-06-prefill-tax.md at main · thekiraproject/kira-project-site · GitHub"
author: "thekiraproject"
captured_at: "2026-07-07T19:43:29Z"
capture_tool: "hn-digest"
hn_id: 48822323
score: 1
comments: 0
posted_at: "2026-07-07T19:15:23Z"
tags:
  - hacker-news
  - translated
---

# Local AI is re-reading its own prompt

- HN: [48822323](https://news.ycombinator.com/item?id=48822323)
- Source: [github.com](https://github.com/thekiraproject/kira-project-site/blob/main/posts/2026-07-06-prefill-tax.md)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T19:15:23Z

## Translation

タイトル: ローカル AI が自身のプロンプトを再読み取りしています
記事のタイトル: kira-project-site/posts/2026-07-06-prefill-tax.md at main · thekiraproject/kira-project-site · GitHub
説明: Kira プロジェクト — ローカルファーストの複合 AI システム。測定したものを公開します。 - kira-project-site/posts/2026-07-06-prefill-tax.md (メイン) · thekiraproject/kira-project-site

記事本文:
kira-project-site/posts/2026-07-06-prefill-tax.md at main · thekiraproject/kira-project-site · GitHub
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
キラプロジェクト
/
キラプロジェクトサイト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

ns
コード
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 101 行 (58 loc) · 10.9 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルをコピー raw ファイルをダウンロード アウトライン編集と raw アクション ローカル AI が独自のプロンプトを再読み取りしています
ローカル AI パイプラインでプレフィル税を測定し、排除します。キラ プロジェクト、2026 年 7 月。特許出願中 (米国暫定番号 64/105,533)。
全員が同じ方法でローカル LLM のベンチマークを行います。つまり、1 秒あたりのトークンをデコードします。他の番号は誰も見ていません。私たちのシステムでは、もう 1 つの数字がモデル時間全体のほぼ半分を占めていました。そして、この時間は、モデルがすでに読んだテキストを、目に見えない形で 1 日に何千回も再読するのに費やされていました。
簡単に定義してからストーリーを説明します。 LLM は 2 つのフェーズで応答します。プレフィル読み取り — 最初の出力トークンが存在する前に、プロンプト全体がトークンごとにモデルの動作状態に処理されます (これは、応答が開始される前の一時停止です) — およびデコード書き込み (その状態に対して一度に 1 トークンずつ)。 「1 秒あたりのトークン数」はデコードのみを測定します。プレフィルのコストはプロンプトの長さに応じて変化します。マルチステージ パイプラインでは、プロンプトはほとんどがモデルが以前に 1 万回読み込んだ定型文です。証言録取の各質問に答える前に事件ファイル全体を読み直す弁護士を想像してみてください。
これは、その税金を発見し、そのコストを正確に測定し、キャッシュ自体よりも重要であると私たちが考える規律に従ってそれを削除する物語です。以下のすべては、公的証拠台帳に記録された領収書を使用して、実稼働システム上の実際のハードウェアで測定されます。
Kira はローカルの複合 AI アシスタントです。1 つの 90 億パラメータ モデル (8 ビット量子化、ハイブリッド アテンション) が、24 GB のメモリを搭載した Apple M4 Pro 上であらゆるパイプラインの役割 (ルーター、スペシャリスト、シンセサイザー、エージェント プランナー) を果たします。

Apple の MLX フレームワークを通じてメモリを強化します。クラウド呼び出しは一切ありません。クエリは、分類子、約 50 個のライブ ツールを使用した推論ループ、およびシンセサイザーを通過します。各ステージは、異なるプロンプトを身に着けた同じモデルです。
この「異なるプロンプト」の詳細に、ストーリーが隠されています。
誰も再開しなかった回帰
このプロジェクトの歴史は毎月、この記事の執筆時点でそのうち 5 件が建築決定記録に記録されており、そのうちの 1 件にはひっそりと逆転する発見が含まれています。私たちは早い段階で、KV キャッシュ プレフィックスの再利用をベンチマークし、無料で利用できると結論付けました。当時のランタイム (llama.cpp 上の Ollama) は、前のリクエストとトークン プレフィックスを共有するリクエストのキー/値の状態を自動的に再利用しました。静的ステージ プロンプト + ピン留めされたモデル = 最初に共有プレフィックスをスキップした後のすべての呼び出し。同じ調査では、MLX にはクロスリクエスト プレフィックス キャッシュが組み込まれていないことが指摘されています。その日の発見は両方とも正しかった。
次に、ランタイムを移行しました。Ollama を使用せず、インプロセス MLX を使用しました。これには、独立した優れた理由がありました (クエリあたり 93 秒対 180 秒、ファントム モデルの読み込みなし)。移行により自動プレフィックス再利用はサイレントに削除され、誰も古い検出結果を再度開くことはありませんでした。さらに悪いことに、新しいクライアントは呼び出しごとに、prompt_eval_count: 0 を報告しました。そのコストは単に注意を払っていないだけではなく、文字通り私たち自身の機器には見えませんでした。
教訓 1: ランタイムを変更するときは、「X を無料で入手します」をすべて再監査してください。無料のものは離脱を発表しません。
何かを構築する前に、推論エンジン独自の呼び出しごとの統計 (プロンプト トークン数とプリフィル期間) をテレメトリに表示しました。メカニズムやキャッシュはなく、目だけです。彼らがライブブローカーで見たもの:
そして重要な付随事実: 移行後、エンドツーエンドの所要時間 ≈ モデル時間 (エンドツーエンドの 46.7 秒 vs 45.7 秒)

f スタンダード ティアのモデル時間）。最適化に必要なモデル以外のオーバーヘッドは残りませんでした。残りのレイテンシはモデル時間であり、ユーザーが感じる層ではそのほぼ半分は、すべてのクエリの呼び出しごとに同じステージ システム プロンプト、ツール スキーマ、および少数ショットのサンプルを再エンコードするモデルでした。
組み立てる前に賞品のサイズを測定する
2 回目の測定、まだメカニズムはありません。環境ゲート フラグの下でライブ プロンプトをキャプチャし、バイト安定プレフィックス シェア (同じステージの呼び出し全体で同一のプロンプト バイトの割合) を計算しました。答え: 63% 。この数字は達成可能な勝利の限界を示しており、ビルドを作成する前にそのビルドを実行する価値があることがわかりました。
問題点も露呈した。一部のステージでは、静的コンテンツの前に動的コンテンツ (ツールの結果、セッション コンテキスト) がインターリーブされ、安定したプレフィックスが縮小されました。私たちのスペシャリストステージは、書かれたままでは 26% しかキャッシュできませんでした。
仕組み（目新しくない部分）
mlx-lm の LRUPromptCache をデフォルトのオフ フラグの背後でクライアントに接続しました。
ピン留めされたステージのプレフィックス。パイプライン ステージは、リクエストの前に静的プレフィックスをオーケストレーターに登録します。登録されたプレフィックスは、ブローカーの存続期間中、保証された常駐を取得します。パイプラインは静的データがどこで終了するかについての情報を所有しており、キャッシュ層は推測しません。
クエリ内チェーン。各呼び出しの後、完全なトークン シーケンスが挿入されるため、エージェント ループのターン N+1 (ターン N のトランスクリプトを拡張) は、より短いプレフィックス ヒットを取得します。転写物の再プレフィルは線形コストにまで縮小します。
静的優先パーティション。ステージ プロンプトは、静的コンテンツ (システム プロンプト ∥ ツール スキーマ ∥ エグザンプラ) がすべての動的コンテンツよりも前に配置されるように再構築されました。スペシャリスト ステージのキャッシュ可能なシェアは 26% から 96% になりました。
プレフィックス キャッシュ自体には目新しさはありません — RadixAttend、vLLM の自動プレフィックス キャッシング、Marco

ni (ハイブリッド アテンション キャッシング) と KVFlow (ワークフロー認識キャッシング) にはその根拠があり、サービス提供エンジンを構築している場合はこれらを読む必要があります。この分野に欠けていると思われるのは次の部分です。
測定ゲートによるアクティベーション (つまりその部分)
静的優先パーティションの特徴は、モデルに表示されるプロンプトを変更することです。サービス層キャッシュは、「正確な再利用は出力を変更しない」と正直に主張できますが、キャッシュ可能なプレフィックスを広げるためにプロンプ​​トを再構築した瞬間、構造上、正確性は保証されなくなります。それを測定する必要があります。ほとんどのシステムはそうではありません。出荷時には最適化が有効になっていると想定されます。
私たちは次の分野に基づいてアクティベーションを制御しました。
パーティションは、キャッシュが有効な構成でもキャッシュが無効な構成でも無条件に適用されます。これが負荷のかかる動きです。どちらの構成も同一のプロンプト構造を共有するため、パリティ評価では 1 つの変数 (キャッシュ) が分離されます。
バイト等価パリティ ゲートは、アサーション バッテリー (1,100 以上のライブ アサーション - フィクスチャは製品仕様です) にわたる 2 つの構成間のゲート結果を比較します。結果: ゲート等価性は正確です。
その場合にのみ、明示的な演算子が反転します。人間が測定を確認するまではデフォルトでオフになります。
キャッシュ有効化フラグは、セマンティック応答キャッシュのキーとなるパイプライン バージョン ハッシュに組み込まれます。ある有効化状態で生成された応答は、別の有効化状態では決して提供されません。アダプター (LoRA) ID の変更により、プレフィックス キャッシュの状態がフラッシュされます。これは、ある重みセットの下で計算された KV が別の重みセットの下では無効になるためです。
活性化後にエンドツーエンドで測定: 標準 -34%、薬剤 -45%、研究 -24%。そして、ステップ 1 の機器はまだ動作しており、問題を発見したメーターは修正を変更せずに検証します。
正確に描く価値のあるもう 1 つの区別は、

文献では曖昧になっているため、これは出力パリティに対する測定ゲートの有効化であり、すでにアクティブなキャッシュのパラメータを調整するものではなく (ヒット率ブートストラップがそれを行います)、エントリごとのアドミッション予測 (投機的挿入がそれを行います) ではありません。私たちのゲートの答えは、両方に先立つものです。この最適化はそもそもこのシステムに存在すべきでしょうか?私たちのより広範なランタイムは、トレーニングされたアダプターに同じ規律を適用します。つまり、ゲートは 4 つの候補を拒否し、0 を有効にしました。 「はい」としか言えない門は門ではありません。
防御的開示: 永続化ライフサイクル
自然な後続 (再起動後も固定されたプレフィックスを保持する) は 2 層の分割として構築され、従来技術として特許取得ではなく意図的に公開されています。
レシピ レイヤー (ピン留めする内容のマニフェスト: ステージ ID、プレフィックス境界、コンテンツ ハッシュ) はインストーラーに同梱されており、すべて存続します。
結果レイヤー (実際の KV バイト) は、レシピからの最初の起動時にデバイス上で再計算され、コンテンツ ハッシュによってキー設定されるため、古いことを自己検出し、アダプターの変更が発生すると完全にパージされます。
レシピは移動し、結果はデバイスローカルで、ハッシュが調停します。ローカル AI インストーラーを出荷する場合は、この形状が必要になります。この投稿の日付の時点で開示されていると考えてください。
ローカルのマルチステージ LLM システムを実行する場合は、次の順序で実行します。
エンジンのプレフィル統計を表示します。クライアントがゼロを報告した場合、それはコストがゼロではなく、可視性のバグです。
何かを構築する前に、キャプチャされたライブ プロンプトでバイト安定プレフィックス シェアを測定します。それはあなたの勝利を制限します。
パーティションは静的優先を要求します。キャッシュされた構成だけでなく、無条件に実行するため、常にパリティを証明できます。
ゲートのアクティブ化は、再利用が正確であるという議論ではなく、測定された出力パリティに基づいて行われます。プロンプト構造に触れた瞬間

そうですね、そうではありません。
キャッシュの有効性を構成 ID に結び付けます。フラグはバージョン ハッシュに、アダプター ハッシュはフラッシュ トリガーに組み込まれます。複数の構成状態にわたって機能するキャッシュは、デモを待っている正確性のバグです。
プレフィル税は実際のもので、おそらくローカル パイプラインでこれまで測定したことのない最大のレイテンシ項目であり、正直に測定するとほぼ全額返金可能です。
Kira は、ローカルファーストの複合 AI システムです。1 つの 9B モデル、約 50 のツール、クラウド推論はゼロで、AI ペア エンジニアリングを使用して 1 人によって構築および運用されます。ここで説明するイネーブルメントとガバナンスの規律は、米国特許出願中です (仮出願 64/105,533、2026 年 7 月 6 日出願)。測定記録 (テレメトリ、パリティ ゲート、証拠台帳) は、プロジェクトのアーキテクチャ決定記録に維持されます。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The Kira Project — a local-first compound-AI system. We publish what we measure. - kira-project-site/posts/2026-07-06-prefill-tax.md at main · thekiraproject/kira-project-site

kira-project-site/posts/2026-07-06-prefill-tax.md at main · thekiraproject/kira-project-site · GitHub
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
thekiraproject
/
kira-project-site
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 101 lines (58 loc) · 10.9 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions Your local AI is re-reading its own prompt
Measuring — and eliminating — the prefill tax in a local AI pipeline. The Kira Project, July 2026. Patent pending (U.S. provisional 64/105,533).
Everyone benchmarks local LLMs the same way: decode tokens per second. Nobody watches the other number. On our system, the other number was nearly half of all model time — and it was being spent re-reading text the model had already read, thousands of times a day, invisibly.
Quick definitions, then the story. An LLM answers in two phases: prefill reads — the entire prompt is processed token-by-token into the model's working state before the first output token can exist (that's the pause before the answer starts) — and decode writes, one token at a time against that state. "Tokens per second" measures decode only. Prefill cost scales with prompt length, and in a multi-stage pipeline the prompt is mostly boilerplate the model has read ten thousand times before. Picture a lawyer who re-reads the entire case file before answering each deposition question.
This is the story of finding that tax, measuring exactly what it cost, and removing it with a discipline we think matters more than the cache itself. Everything below is measured on real hardware, on a production system, with the receipts in a public evidence ledger.
Kira is a local compound-AI assistant: one 9-billion-parameter model (8-bit quantized, hybrid attention) serving every pipeline role — router, specialist, synthesizer, agentic planner — on an Apple M4 Pro with 24 GB of unified memory, through Apple's MLX framework. No cloud calls, ever. A query flows through a classifier, a reasoning loop with ~50 live tools, and a synthesizer; each stage is the same model wearing a different prompt.
That "different prompt" detail is where the story hides.
The regression nobody reopened
Every month of this project's history — five of them, at the time of writing — is recorded in architecture decision records, and one of them contains a finding that quietly inverted. Early on, we benchmarked KV-cache prefix reuse and concluded we had it for free: our then-runtime (Ollama, on llama.cpp) automatically reused the key/value state for any request sharing a token prefix with the previous one. Static stage prompts + pinned models = every call after the first skipped its shared prefix. The same research noted that MLX had no built-in cross-request prefix caching. Both findings were correct that day.
Then we migrated runtimes — Ollama out, in-process MLX in, for independently excellent reasons (93 s vs 180 s per query, no phantom model loads). The migration silently deleted the automatic prefix reuse , and nobody reopened the old finding. Worse: the new client reported prompt_eval_count: 0 on every call. The cost wasn't just unpaid attention — it was literally invisible to our own instrumentation .
Lesson one: when you change runtimes, re-audit every "we get X for free." The free things don't announce their departure.
Before building anything, we surfaced the inference engine's own per-call statistics — prompt token counts and prefill durations — into our telemetry. No mechanism, no cache, just eyes. What they saw, on the live broker:
And a critical companion fact: after the migration, end-to-end wall time ≈ model time (46.7 s end-to-end vs 45.7 s of model time on the standard tier). There was no non-model overhead left to optimize. The remaining latency was model time — and nearly half of it, on the tiers users feel, was the model re-encoding the same stage system prompts, tool schemas, and few-shot examples on every call of every query.
Measure the prize before building
Second measurement, still no mechanism: we captured live prompts under an environment-gated flag and computed the byte-stable prefix share — the fraction of prompt bytes identical across calls of the same stage. Answer: 63% . That number bounds the achievable win and told us the build was worth doing before we wrote it .
It also exposed a problem. Some stages interleaved dynamic content (tool results, session context) before static content, shrinking the stable prefix. Our specialist stage was only 26% cacheable as-written.
The mechanism (the part that isn't novel)
We wired mlx-lm's LRUPromptCache into our client behind a default-off flag:
Pinned stage prefixes. Pipeline stages register their static prefixes with the orchestrator before any request ; registered prefixes get guaranteed residency for the broker's lifetime. The pipeline owns the knowledge of where static ends — the cache layer doesn't guess.
Within-query chaining. After each call, the full token sequence is inserted, so an agentic loop's turn N+1 — which extends turn N's transcript — gets a shorter-prefix hit. Transcript re-prefill collapses to linear cost.
The static-first partition. Stage prompts were restructured so static content (system prompt ∥ tool schemas ∥ exemplars) precedes all dynamic content. The specialist stage's cacheable share went from 26% to 96% .
We claim no novelty for prefix caching itself — RadixAttention, vLLM's automatic prefix caching, Marconi (hybrid-attention caching), and KVFlow (workflow-aware caching) own that ground, and if you're building a serving engine you should read them. What we think the field is missing is the next part.
Measurement-gated activation (the part that is)
Here's the thing about that static-first partition: it changes the prompts the model sees. Serving-layer caches can honestly claim "exact reuse doesn't change output" — but the moment you restructure prompts to widen the cacheable prefix, correctness is no longer guaranteed by construction. You have to measure it. Most systems don't; they ship the optimization enabled and assume.
We gated activation on a discipline:
The partition is applied unconditionally — in both cache-enabled and cache-disabled configurations. This is the load-bearing move. Both configurations share identical prompt structure, so the parity evaluation isolates exactly one variable: the cache.
A byte-equivalence parity gate compares gate outcomes between the two configurations across our assertion battery (1,100+ live assertions — the fixture is our product spec). Result: gate equivalence exact .
Only then, an explicit operator flip. Default-off until a human reviews the measurements.
The cache-enablement flag is composed into the pipeline-version hash that keys our semantic response cache — responses generated under one enablement state are never served under another. A change of adapter (LoRA) identity flushes prefix-cache state, because KV computed under one set of weights is invalid under another.
Measured after activation, end-to-end: standard −34%, agentic −45%, research −24%. And the instrument from step one is still running — the meter that found the problem verifies the fix, unchanged.
One more distinction worth drawing precisely, because the literature blurs it: this is measurement gating enablement on output parity — not tuning a parameter of an already-active cache (hit-rate bootstraps do that), and not per-entry admission forecasting (speculative insertion does that). The question our gate answers is prior to both: should this optimization exist in this system at all? Our broader runtime applies the same discipline to trained adapters — where, for what it's worth, the gate has refused four candidates and enabled zero. A gate that can only say yes isn't a gate.
Defensive disclosure: the persistence lifecycle
The natural follow-on — persisting pinned prefixes across restarts — we built as a two-layer split and are deliberately publishing rather than patenting, as prior art:
The recipe layer (a manifest of what to pin: stage identities, prefix boundaries, content hashes) ships with the installer and survives everything.
The result layer (the actual KV bytes) is recomputed on the device at first boot from the recipe, keyed by content hash so staleness self-detects, and purged wholesale on any adapter mutation.
Recipe travels, results are device-local, hashes arbitrate. Anyone shipping local-AI installers will need this shape; consider it disclosed as of this post's date.
If you run a local multi-stage LLM system, in order:
Surface your engine's prefill stats. If your client reports zero, that's a bug in your visibility, not a zero cost.
Measure your byte-stable prefix share on captured live prompts before building anything. It bounds your win.
Partition prompts static-first — and do it unconditionally, not only in the cached configuration, so you can ever prove parity.
Gate activation on measured output parity , not on the argument that reuse is exact. The moment you touched prompt structure, it isn't.
Couple cache validity to configuration identity — flags into your version hash, adapter hashes into your flush triggers. A cache that can serve across config states is a correctness bug waiting for its demo.
The prefill tax is real, it is probably the largest latency line-item in your local pipeline you have never measured, and — measured honestly — it is almost entirely refundable.
Kira is a local-first compound-AI system: one 9B model, ~50 tools, zero cloud inference, built and operated by a single person with AI pair-engineering. The enablement-governance discipline described here is U.S. patent pending (provisional application 64/105,533, filed July 6, 2026). The measurement record — telemetry, parity gates, and the evidence ledger — is maintained in the project's architecture decision records.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
