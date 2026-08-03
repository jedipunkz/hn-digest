---
source: "https://github.com/patchy631/time-to-first-token"
hn_url: "https://news.ycombinator.com/item?id=49158330"
title: "Time-to-first-token: A 10-week, 30-minutes-a-day roadmap for LLM inference"
article_title: "GitHub - patchy631/time-to-first-token: A 10-week, 30-minutes-a-day roadmap for LLM inference serving and optimization. vLLM, SGLang, quantization, speculative decoding, benchmarking. · GitHub"
author: "simonpure"
captured_at: "2026-08-03T17:48:04Z"
capture_tool: "hn-digest"
hn_id: 49158330
score: 3
comments: 0
posted_at: "2026-08-03T16:54:10Z"
tags:
  - hacker-news
  - translated
---

# Time-to-first-token: A 10-week, 30-minutes-a-day roadmap for LLM inference

- HN: [49158330](https://news.ycombinator.com/item?id=49158330)
- Source: [github.com](https://github.com/patchy631/time-to-first-token)
- Score: 3
- Comments: 0
- Posted: 2026-08-03T16:54:10Z

## Translation

タイトル: 最初のトークンまでの時間: LLM 推論の 10 週間、1 日 30 分のロードマップ
記事のタイトル: GitHub - patchy631/time-to-first-token: LLM 推論の提供と最適化のための 10 週間、1 日 30 分のロードマップ。 vLLM、SGLang、量子化、投機的デコード、ベンチマーク。 · GitHub
説明: LLM 推論の提供と最適化のための 10 週間、1 日 30 分のロードマップ。 vLLM、SGLang、量子化、投機的デコード、ベンチマーク。 - patchy631/最初のトークンまでの時間

記事本文:
GitHub - patchy631/time-to-first-token: LLM 推論の提供と最適化のための 10 週間、1 日 30 分のロードマップ。 vLLM、SGLang、量子化、投機的デコード、ベンチマーク。 · GitHub
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
パッチー631
/
t

imeから最初のトークンまで
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット ライセンス ライセンス README.md README.md Index.html Index.html progress.md progress.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
1 つのサービスを出荷することで LLM 推論サービスを学ぶ
LLM 推論について読むだけでなく、運用環境で実際に LLM 推論を実行したいエンジニア向けの 10 週間、1 日 30 分のロードマップです。
50 セッション。それらのそれぞれが単一のアーティファクトを提供します。OpenAI 互換の推論サービスは、レンタルした GPU、計測器にデプロイし、1000 を超える同時リクエストの負荷テストを行い、量子化と投機的デコードによる最適化を行い、コストを意識したルーターを前に置き、再現可能なベンチマークとして公開します。
別のアプローチでは、17 の接続されていない実験を実行し、セットアップにほとんどの時間を費やします。成長を続ける 1 つのサービスは、同じ内容をカバーし、表示できるものを残します。
Python、アーキテクチャ レベルのトランスフォーマー、およびコマンド ラインに慣れている必要があります。以前のサービス提供、Kubernetes、または CUDA の経験は必要ありません。
トピックをすでに知っている場合は、「スキム」とマークされたセッションが圧縮の対象になります。ビルド セッションを圧縮しないでください。それが重要です。
自分で構成、計測、調整したサービススタック
TTFT、トークン間のレイテンシー、スループット、キューの深さ、リクエストあたりのコストを表示する Grafana ダッシュボード
1000 を超える同時リクエストを再現可能に駆動する負荷テスト ハーネス
FP16、FP8、INT4、投機的デコード、および KV エビクションにわたるベンチマークされたバリアント
リクエストごとのトークンバジェットを備えたコスト/遅延/品質のルーター
固定されたバージョンと再現可能なコミュニケーションを含む公開されたベンチマークの記事

そして
セッションの長さ
30分
週あたりのセッション数
5 日プラス 2 緩衝日
合計
10 週間、50 セッション、25 時間
GPUコスト
30 分までにレンタルされる 24 GB カード約 1 枚と、H100 セッション 2 つ
火曜日を欠席しても計画が崩れないように緩衝日が存在します。それらは追いつくためのものであり、新しい素材のためのものではありません。
ここでの順序は意図的なものであり、ほとんどの人がこのリストを書き留める方法とは異なります。 5 つの決断がそれを推進します。
ルーフラインモデルが最初です。その後のすべての最適化は、同じプロット上での移動となります。量子化は、メモリ帯域幅に制限されたデコードを攻撃します。継続的なバッチ処理により、演算強度がコンピューティング ルーフに向けて向上します。投機的デコードでは、メモリーに制約された状況では安価な FLOP を消費して、シーケンシャルなメモリー負荷を削減します。プリフィルはコンピューティングに依存し、デコードは帯域幅に依存し、1 つの GPU をめぐって競合するため、分離が存在します。このモデルがなければ、残りはトリックの詰め合わせです。
最適化の前に測定が行われます。計測器は第 3 週に到着し、負荷テストは第 5 週で、チューニングノブのかなり前に行われます。それらがなければ、それ以降の何も検証できません。また、公開ベンチマークは、ほとんどの場合、モデルが接続された信頼できる測定ハーネスです。
ページング アテンションと連続バッチ処理は構築演習ではありません。 vLLM は両方を実装しており、チャンク プレフィルが vLLM と SGLang のデフォルトのスケジューリング戦略です。これらを再実装すると、メモリの無駄が 4% 未満に低下する理由と、リクエスト間で GPU がアイドリングを停止する理由をコードから説明できるようになるまで、ブロック マネージャーとスケジューラを読むだけで十分です。
ルーターはユニットエコノミクスの隣にあります。コスト、レイテンシ、品質によってバックエンドを選択するルーターは、コードに金額とレイテンシの予算を強制する 1 つのコンポーネントです。同じ週に経済学を勉強すると、その読書がブロではなくルーティングポリシーに変わります

g 投稿に同意しました。
エッジ展開はオプションであり、最後です。 ONNX ランタイム、TensorRT-LLM、および WebLLM は、クライアント側の組み込みランタイムです。他の 9 週間で構築されるデータセンター スタックと運用面をほとんど共有しません。
7 ～ 8B モデルには、およそ 1 つの 24GB GPU が必要です。 30 分までにレンタルし、セッション間ではインスタンスをシャットダウンします。
第 1 週のすべて、第 3 週のほとんど、第 6 週の講義日、および第 9 週のすべての読書には GPU はまったく必要ありません。構築日の前後にレンタルをまとめてください。 1 つの H100 は、正確に 2 つのセッションに必要です。実際に実行する場合は、第 5 週の 1000 同時テストと第 8 週の分解です。
pip インストール vllm
pip インストールガイドllm
pip インストール sglang
さらに、Prometheus および Grafana スタック用の Docker、第 8 週用の種類または小規模のマネージド クラスター。
凡例: 読み取りセッションはコンテキストを読み込み、構築セッションは出力を生成します。スキミングは、概念的にはすでに知っているかもしれない内容をマークします。その価値は実践的な部分にあります。
集中。他のすべてが依存する 1 つのモデルをインストールします。ルーフライン、演算強度、およびプリフィルがコンピューティングで待機するのにデコードがメモリで待機する理由。
中くらい。ビデオファースト。ルーフライン推論は、ライブ導出と図から恩恵を受けます。
成果物。提供するモデルの手動で導出された算術強度の数値と、どのフェーズが帯域幅制限を受けているか、およびその理由を判断する機能。
第 2 週: vLLM: デプロイして内部を読み取る
集中してください。サービスを起動し、コードからページングの設計が明らかになるまで、ブロック マネージャーとスケジューラーに進みます。
中くらい。 PagedAttendance ではテキストファーストであり、文書がより正確です。 V1 アーキテクチャのビデオファースト。
成果物。 7 ～ 8B モデルを提供する OpenAI 互換エンドポイント、およびソースからの PagedAttendance と V1 スケジューラを説明する独自のメモ。
第 3 週: 測定インフラストラクチャ
集中してください。を構築します

最適化前のレンズ。今週以降のすべては、今週だからこそ読み取れるのです。
中くらい。テキストファースト。ベンチマーク手法に関する記述は、入手可能などのビデオよりも強力です。
成果物。独自のサービスの TTFT、トークン間レイテンシ、スループット、キューの深さを表示するライブ Grafana ダッシュボード。
第 4 週: SGLang、RadixAttendant、バッチ処理の内部処理
集中してください。設計空間を教える対比: 固定ブロック ページングとツリー構造のプレフィックスの再利用。
中くらい。混合。 RadixAttend の講演はプロジェクト リーダーによるものです。チャンクされたプリフィルはテキストのみです。
成果物。同じサービスの SGLang バリアントと、vLLM に対する最初のプレフィックス再利用の比較。
第 5 週: 同時 1000 までの負荷テスト
集中してください。再現可能なハーネスを構築し、防御すべきものがなくなる前に、何がベンチマークを防御可能にするのかを学びましょう。
中くらい。テキストファースト。ドキュメントと実践のみ。
成果物。スクリプト化された同時実行数は、p50、p95、および p99 の出力を使用して 1000 を超える同時リクエストをスイープし、2 つの独立したツールによってクロスチェックされます。
第 6 週: 量子化のトレードオフ
集中してください。最初のチューニングノブ、そして最初の週にハーネスが維持されます。
中くらい。ビデオファースト。講義は入り口であり、論文は詳細な部分です。
成果物。 FP8 および INT4/AWQ バリアントは、スループットと高品質プロキシの両方で FP16 ベースラインに対してベンチマークされました。
第 7 週: 投機的デコードと KV エビクション
集中してください。概念は広く知られていますが、実装の詳細は知られていない 2 つのノブ。
中くらい。 vLLM の実装を作成した人によって教えられた、投機的デコードのためのビデオファースト。 KV エビクションのテキストファースト。
成果物。投機的なデコードのバリアントとロングコンテキストのエビクション構成。どちらもベンチマークされ、否定的な結果も含まれます。
第 8 週: 細分化されたサービスと Kubernetes
集中してください。プレフィル

そして、第 1 週から非対称性をデコードすることがアーキテクチャになり、次にそれを実行する運用層になります。
中くらい。細分化のためのビデオファースト。 Kubernetes ではテキストファーストです。変化が速すぎるため、ビデオは最新の状態を維持できません。
成果物。 CPU ではなくキューの深さによって自動スケーリングが行われる Kubernetes 上のサービス。
第 9 週: ユニットエコノミクスとルーター
集中してください。コストの読み取りをコードに変換します。ここで、ロードマップはレイテンシーに関するものではなくなり、マージンに関するものになり始めます。
成果物。コスト、レイテンシー、品質に基づいてバックエンドを選択するルーター。さらに、ダッシュボード上のコストを使用してリクエストごとのトークンの予算を設定します。
第 10 週: 出版して読書習慣を築く
集中してください。アーティファクトを出荷し、それを最新の状態に保つ入力をセットアップします。
成果物。公開されたリポジトリと、ピン留めされたバージョンと再現可能なコマンドを含む記事、さらにカレンダー上の毎週のリサーチ セッション。
必要な設定が少ない順に 3 つのオプションがあります。
1. Web ページ。 Index.html または GitHub Pages バージョンを開き、終了するセッションにチェックを入れます。進行状況はブラウザのローカル ストレージに保存されるため、タブを閉じたりマシンを再起動しても進行状況は失われません。これはブラウザごと、デバイスごとであり、アカウントやサーバーは関係ありません。
進行状況を別のデバイスに移動するには、状態を URL 内の 12 文字のコードにパックする [同期リンクのコピー] を使用するか、後で再インポートできる JSON ファイルをダウンロードする [エクスポート] を使用します。どちらもブラウザのデータをクリアした場合のバックアップにもなります。
プライベート ブラウジングまたはサンドボックス フレーム内でページを開くと、ローカル ストレージがブロックされる可能性があります。ページはこれを検出して警告し、その訪問に対してティックは引き続き機能します。
2. GitHub の問題。リポジトリをフォークし、進行状況テンプレートから新しい問題を開き、そこにあるチェックボックスをオンにします。 GitHub はタスク リストの状態をアカウントに関連付けられた独自のサーバーに保存するため、

サインインするすべてのデバイスで機能し、毎週終了した日付のコメント履歴が表示されます。これは、ブラウザのプロファイルを超えて進行状況を維持したい場合に最適なオプションです。
3. README 自体。リポジトリをフォークし、コピー内のチェックリストにチェックを入れます。コミット履歴が記録になります。
SQLite であろうとなかろうと、サーバー側ストアはここでは間違ったツールであり、その理由を説明する価値があります。それには、ホスティング、サーバーが誰の進行状況を認識するためのアカウント、セッション処理、現在保持しているデータのプライバシー ポリシー、およびリンクのリスト全体が値であるプロジェクトの継続的なメンテナンスが必要です。特に SQLite は、デプロイされた Web アプリにはあまり適していません。これは、ほとんどの安価なホスティングには、再デプロイ時にデータベース ファイルが破棄される一時的なファイル システムが含まれているためです。
ローカル ストレージでは、それが何もない単一のリーダーでも同じ結果が得られます。唯一提供できないのはクロスデバイス同期であり、同期リンクと JSON エクスポートはバックエンドなしでこのケースをカバーします。
本当にホストされたバージョンが必要な場合、たとえばコホートが共有リーダーボードとともにこれを通過する場合、現実的なスタックは、ID 用の GitHub OAuth とマネージド Postgres または Turso インスタンスであり、Web 上の SQLite ファイルではありません。

[切り捨てられた]

## Original Extract

A 10-week, 30-minutes-a-day roadmap for LLM inference serving and optimization. vLLM, SGLang, quantization, speculative decoding, benchmarking. - patchy631/time-to-first-token

GitHub - patchy631/time-to-first-token: A 10-week, 30-minutes-a-day roadmap for LLM inference serving and optimization. vLLM, SGLang, quantization, speculative decoding, benchmarking. · GitHub
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
patchy631
/
time-to-first-token
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits LICENSE LICENSE README.md README.md index.html index.html progress.md progress.md View all files Repository files navigation
Learn LLM Inference Serving by Shipping One Service
A 10-week, 30-minutes-a-day roadmap for engineers who want to actually run LLM inference in production, not just read about it.
Fifty sessions. Every one of them feeds a single artifact: an OpenAI-compatible inference service that you deploy on a rented GPU, instrument, load test past 1000 concurrent requests, optimize with quantization and speculative decoding, put a cost-aware router in front of, and publish as a reproducible benchmark.
The alternative approach, running seventeen disconnected experiments, spends most of its time on setup. One service that keeps growing gets you the same coverage and leaves you with something to show.
You should be comfortable with Python, transformers at the architecture level, and the command line. You do not need prior serving, Kubernetes, or CUDA experience.
If you already know a topic, the sessions marked skim are the ones to compress. Do not compress the build sessions, they are the point.
A serving stack you configured, instrumented, and tuned yourself
Grafana dashboards showing TTFT, inter-token latency, throughput, queue depth, and cost per request
A load test harness that reproducibly drives 1000+ concurrent requests
Benchmarked variants across FP16, FP8, INT4, speculative decoding, and KV eviction
A cost/latency/quality router with per-request token budgeting
A published benchmark writeup with pinned versions and reproducible commands
Session length
30 minutes
Sessions per week
5, plus 2 buffer days
Total
10 weeks, 50 sessions, 25 hours
GPU cost
Roughly one 24GB card rented by the half hour, plus two H100 sessions
Buffer days exist so that missing a Tuesday does not collapse the plan. They are for catching up, not for new material.
The order here is deliberate and differs from how most people write this list down. Five decisions drive it.
The roofline model comes first. Every optimization later is a move on the same plot. Quantization attacks memory-bandwidth-bound decode. Continuous batching raises arithmetic intensity toward the compute roof. Speculative decoding spends FLOPs, which are cheap in the memory-bound regime, to cut sequential memory loads. Disaggregation exists because prefill is compute-bound and decode is bandwidth-bound and they fight over one GPU. Without this model, the rest is a bag of tricks.
Measurement comes before optimization. Instrumentation lands in week 3 and load testing in week 5, well before the tuning knobs. Nothing after them is verifiable without them, and a public benchmark is mostly a credible measurement harness with a model attached.
Paged attention and continuous batching are not build exercises. vLLM implements both, and chunked prefill is the default scheduling strategy in vLLM and SGLang. Reimplementing them teaches less than reading the block manager and the scheduler until you can explain from the code why memory waste drops under 4 percent and why the GPU stops idling between requests.
The router sits next to unit economics. A router that picks a backend by cost, latency, and quality is the one component that forces a dollar figure and a latency budget into code. Studying economics the same week turns the reading into a routing policy instead of a blog post you agreed with.
Edge deployment is optional and last. ONNX Runtime, TensorRT-LLM, and WebLLM are client-side and embedded runtimes. They share almost no operational surface with the datacenter stack the other nine weeks build.
A 7-8B model needs roughly one 24GB GPU. Rent by the half hour and shut the instance down between sessions.
All of week 1, most of week 3, the week 6 lecture days, and all of week 9 reading need no GPU at all. Batch your rentals around the build days. One H100 is needed for exactly two sessions: the 1000-concurrent test in week 5, and disaggregation in week 8 if you run it for real.
pip install vllm
pip install guidellm
pip install sglang
Plus Docker for the Prometheus and Grafana stack, and kind or a small managed cluster for week 8.
Legend: read sessions load context, build sessions produce output. skim marks material you may already know conceptually, where the value is in the hands-on part.
Focus. Install the one model everything else hangs on. Roofline, arithmetic intensity, and why decode waits on memory while prefill waits on compute.
Medium. Video-first. Roofline reasoning benefits from live derivation and diagrams.
Deliverable. A hand-derived arithmetic intensity figure for the model you will serve, and the ability to say which phase is bandwidth-bound and why.
Week 2: vLLM: deploy it, then read its internals
Focus. Stand the service up, then go into the block manager and scheduler until the paging design is obvious from the code.
Medium. Text-first for PagedAttention, where the paper is more precise. Video-first for the V1 architecture.
Deliverable. An OpenAI-compatible endpoint serving a 7-8B model, plus your own notes explaining PagedAttention and the V1 scheduler from source.
Week 3: Measurement infrastructure
Focus. Build the lens before the optimizations. Everything after this week is only legible because of this week.
Medium. Text-first. The writing on benchmarking methodology is stronger than any available video.
Deliverable. A live Grafana dashboard showing TTFT, inter-token latency, throughput, and queue depth on your own service.
Week 4: SGLang, RadixAttention, and batching internals
Focus. The contrast that teaches the design space: fixed-block paging versus tree-structured prefix reuse.
Medium. Mixed. The RadixAttention talk is from the project lead. Chunked prefill is text-only.
Deliverable. An SGLang variant of the same service, and a first prefix-reuse comparison against vLLM.
Week 5: Load testing to 1000 concurrent
Focus. Build the reproducible harness and learn what makes a benchmark defensible before you have anything to defend.
Medium. Text-first. Documentation and hands-on only.
Deliverable. A scripted concurrency sweep past 1000 concurrent requests with p50, p95, and p99 output, cross-checked by two independent tools.
Week 6: Quantization tradeoffs
Focus. The first tuning knob, and the first week your harness earns its keep.
Medium. Video-first. The lectures are the entry point, the papers are the deep dive.
Deliverable. FP8 and INT4/AWQ variants benchmarked against the FP16 baseline on both throughput and a quality proxy.
Week 7: Speculative decoding and KV eviction
Focus. Two knobs where the concept is widely known but the implementation detail is not.
Medium. Video-first for speculative decoding, taught by the person who wrote vLLM's implementation. Text-first for KV eviction.
Deliverable. A speculative decoding variant and a long-context eviction configuration, both benchmarked, including negative results.
Week 8: Disaggregated serving and Kubernetes
Focus. The prefill and decode asymmetry from week 1 becomes an architecture, then the operations layer that runs it.
Medium. Video-first for disaggregation. Text-first for Kubernetes, which changes too fast for video to stay current.
Deliverable. Your service on Kubernetes with autoscaling driven by queue depth rather than CPU.
Week 9: Unit economics and the router
Focus. Turn cost reading into code. This is where the roadmap stops being about latency and starts being about margin.
Deliverable. A router that picks a backend by cost, latency, and quality, plus per-request token budgeting with cost on the dashboard.
Week 10: Publish, then build the reading habit
Focus. Ship the artifact and set up the inputs that keep it current.
Deliverable. A published repo and writeup with pinned versions and reproducible commands, plus a weekly research session on the calendar.
Three options, in order of how little setup they need.
1. The web page. Open index.html , or the GitHub Pages version, and tick sessions as you finish them. Progress saves to your browser's local storage, so closing the tab or restarting your machine does not lose it. It is per browser and per device, with no account and no server involved.
To move progress to another device, use Copy sync link , which packs your state into a 12-character code in the URL, or Export , which downloads a JSON file you can re-import later. Both are also your backup if you clear browser data.
If you open the page in private browsing or inside a sandboxed frame, local storage may be blocked. The page detects this and warns you, and ticks will still work for that visit.
2. A GitHub issue. Fork the repo, open a new issue from the progress template, and tick the checkboxes there. GitHub stores task-list state on its own servers, tied to your account, so it works across every device you sign in from and gives you a dated comment history of when you finished each week. This is the best option if you want your progress to outlive a browser profile.
3. The README itself. Fork the repo and tick the checklists in your copy. Your commit history becomes the record.
A server-side store, SQLite or otherwise, is the wrong tool here and worth explaining why. It would require hosting, accounts so the server knows whose progress is whose, session handling, a privacy policy for the data you are now holding, and ongoing maintenance for a project whose entire value is a list of links. SQLite in particular is a poor fit for a deployed web app, because most cheap hosting has an ephemeral filesystem that discards the database file on redeploy.
Local storage gives the same outcome for a single reader with none of that. The one thing it does not give is cross-device sync, and the sync link and JSON export cover that case without a backend.
If you genuinely want a hosted version, for example a cohort going through this together with a shared leaderboard, the realistic stack is GitHub OAuth for identity plus a managed Postgres or Turso instance, not a SQLite file on a w

[truncated]
