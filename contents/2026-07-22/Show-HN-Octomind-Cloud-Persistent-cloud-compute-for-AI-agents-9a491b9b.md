---
source: "https://octomind.run/blog/octomind-cloud-launch"
hn_url: "https://news.ycombinator.com/item?id=49007062"
title: "Show HN: Octomind Cloud – Persistent cloud compute for AI agents"
article_title: "Octomind Cloud Is Live: Real Machines for Your Agent, Every Model Through One Key | Octomind"
author: "brightside667"
captured_at: "2026-07-22T14:56:56Z"
capture_tool: "hn-digest"
hn_id: 49007062
score: 3
comments: 1
posted_at: "2026-07-22T14:06:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Octomind Cloud – Persistent cloud compute for AI agents

- HN: [49007062](https://news.ycombinator.com/item?id=49007062)
- Source: [octomind.run](https://octomind.run/blog/octomind-cloud-launch)
- Score: 3
- Comments: 1
- Posted: 2026-07-22T14:06:07Z

## Translation

タイトル: Show HN: Octomind Cloud – AI エージェント向けの永続的なクラウド コンピューティング
記事のタイトル: Octomind Cloud が稼働中: エージェント用の実マシン、1 つのキーですべてのモデル |オクトマインド
説明: コーディング エージェントはラップトップの蓋とともに機能しなくなり、API キーは 5 つのダッシュボードに存在し、AI の請求額は驚くべきものになります。私たちは次の 3 つすべてを修正するために Octomind Cloud を構築しました: 永続マシンは 1 時間あたり 0.05 ドルから – 本物の Docker を内蔵 – 公開されたトークンごとの価格で 1 つのキーの背後にある 21 のモデルのハブ、および 1 つのウォレット
[切り捨てられた]

記事本文:
Octomind Cloud が稼働中: エージェントの実マシン、すべてのモデルを 1 つのキーで利用可能 |オクトマインド
製品をタップ クラウド ハブ ブログ ビジョン ドキュメント インストール パネル GitHub ⭐ 87 ブログに戻る Octomind Cloud が稼働中: エージェント用の実マシン、1 つのキーですべてのモデル
コーディング エージェントはラップトップの蓋とともに機能しなくなり、API キーは 5 つのダッシュボードに保存され、AI の請求額は驚くべきものになります。私たちは、3 つすべてを修正するために Octomind Cloud を構築しました。永続マシンは 1 時間あたり 0.05 ドルから、本物の Docker が内蔵されています。公開されたトークンごとの価格で 1 つのキーの背後にある 21 モデルのハブと、上限が確認できる 1 つのウォレットです。
本日、Octomind Cloud がベータ版として公開されます。それが何であるかを説明する前に、私たちがそれを構築するに至った 3 つの問題について説明しましょう。おそらく、あなたは今、そのうちの少なくとも 2 つを抱えているはずだからです。
エージェントはラップトップとともに生き、そして死んでいきます。長いリファクタリングを開始すると、エージェントが開始して 40 分が経過します。その後、蓋を閉めたり、電車に乗ったり、Wi-Fi が中断したりする必要があります。セッションは失われ、コンテキストも失われ、生成された作品に対して支払ったトークンは決して目にすることはできません。エージェントは長時間実行されるプロセスです。ラップトップは長時間稼働するマシンではありません。
API キーは 5 つのダッシュボードに存在します。安価なモデルには OpenRouter、優れたモデルには Anthropic、ツールで必要なため OpenAI、埋め込みには Voyage、そしてパブリックになるまで 1 git add -A で済む .env ファイル。すべてのプロバイダーには独自の請求ページ、独自のレート制限、独自の「最低 5 ドルの追加」があり、エージェントが実際に費やした金額を単一のビューで確認することはできません。
AI の請求額は数字ではなく気分です。シート価格のツールを使用すると、レート制限ルーレットが提供されます。これは、他人のスケジュールに合わせてリセットされる不透明な許容量です。従量制のツールを使用すると、逆の恐怖が生じます。暴走エージェントが一晩で 80 ドルを消費するのを止めるものは何もありません。どちらもトークンごとの p を示しません。

ご飯を食べたり、メーターを見たりできます。
Octomind Cloud は、3 つすべてに対する私たちの答えです。 1 つのサブスクリプション、1 つのウォレット、1 つのキーを共有するのは、ハブ (モデル) とマシン (コンピューティング) の 2 つです。この投稿は、現在公開されているすべての情報をまとめた地図です。各セクションはそれぞれ詳細に説明する価値があり、今後も詳細に説明していきます。しかし、この 1 つの記事で、経済的なことも含めて、何が得られるのかを正確に知ることができます。理由がわからない価格設定は、ビジネス モデルではなくバグだからです。
ハブ: 一度ログインすれば、すべてのモデルが動作します
ハブは、厳選されたモデル名簿の前にある OpenAI 互換のゲートウェイです。パネル、CLI、OpenAI API を使用するあらゆるツールで、1 つのキーですべてのモデルが機能します。
カール https:///hub.octomind.run/v 1 /chat/completions \
-H "認可: ベアラー $OCTOMIND_KEY " \
-d '{"model": "deepseek-v4-flash", "messages": [{"role": "user", "content": "hi"}]}' CLI を使用している場合は、そのキーを処理する必要はまったくありません。 0.38.0 以降、octomind ログインはショート コードを表示し、既にサインインしているパネルを開き、キーをディスク自体に書き込みます。
$ オクトマインドログイン
╭ログイン・オクトマインドアカウント
│ コード TX8T-UC7U
│ URL https://octomind.run/app/login/cli
│
│ ブラウザでコードを確認してサインインを完了します。
│ 待っています… これは標準のデバイス フローです (GH 認証ログインで使用されるものと同じです)。そのため、パスワードが端末に入力されることはなく、各マシンは独自の取り消し可能なキーを取得し、/usage コマンドはパネルを開かずにキャップに対して費やした金額を表示します。
プロバイダーアカウントはありません。キーの広がりはありません。そして、私が最も気にしている部分は、公開価格と表に正確に一致するメーターです。
この名簿は、400 モデルのドロップダウンではなく、意図的に厳選されており、適切な位置を獲得した 21 モデルです。オープンシェルフ (DeepSeek V4 Flash および Pro、Qwen3.7 Plus および

Max、Kimi 2.6 および 2.7 Code、MiniMax M3、GLM-5.2、Gemma 4) はすべての有料プランに含まれており、プランの上限から引き出されます。プレミアム棚ではプリペイド クレジットからトークンごとに請求が行われ、それは本物です。クロード フェイブル 5、オーパス 4.8、ソネット 5。 GPT-5.5、5.5 Pro、高速小型の 5.4 mini および nano を備えた GPT-5.6 ファミリー全体 (Sol、Terra、Luna)。 Gemini 3.5 フラッシュおよび 3.1 Pro。選択する前にすべての価格を確認できます。そして、他のすべてが使い果たされても動作し続ける日当付きの無料モデルもあります。上限に達した瞬間に完全に暗転するエージェントは、遅いエージェントよりも悪いからです。
フェアユースはキーごとではなくアカウントごとです。Free は一度に 1 つのモデル呼び出しを実行し、Pro は 2 つを並行して実行し、最大 5 つ、チーム 15 はプールされます。 1,000 個のキーを作成しても予算は動かず、それを超えると、リクエストは失敗せずに一時的にキューに入れられます。これは、騒々しい隣人の保証を兼ねた悪用防止ルールです。誰も暴走したスクリプトによって完了を待たされることはありません。
内部的には、これは octohub です。オープンソース、Rust、1 つの静的バイナリです。ホストされたハブは、キーの問題が解決された同じコードです。独自のプロバイダー キーを使用して自分で実行したいですか?それは抜け穴ではありません。それを公開することが重要です。
マシン: ラップトップに存在するかどうかに関係なく存在する開発ボックス
マシンは、エージェント ツールチェーン全体が組み込まれた、クラウド内の永続的な Linux コンテナーです。ファイル、リポジトリ、インデックスは、独自のディスク上に永続化されます。これには、octomind (エージェント)、octocode (コード インデックス付け + セマンティック検索)、octobrain (メモリ)、Web ターミナル、git、Docker (詳細は後ほど説明します)、そして — デフォルトのイメージとして — Node、bun、ヘッドレス ブラウザ、ffmpeg、イメージ ツールが含まれています。パネルで作成すると、1 分以内に準備が完了し、あなたのものになります。
請求モデルは私が望む部分です

ほとんどの雲が不透明になったり、敵対的になったりする場所であるため、詳細に防御する必要があります。マシンは常に 3 つの状態のいずれかにあり、その状態によって支払い額が決まります。
秒ごとの計測、自動トランジションなど、何も設定する必要はありません。重要な結果: エージェントがモデルを待機している間 (エージェントの実時間のほとんどはこの時間です)、アクティブ料金の 5 分の 1 を支払うことになります。 5 分間のセッションの料金は約 0.001 セントで、セッション開始ごとに 0.001 ドルがかかります。すべてのマシンに 10 GB の作業ディスクが無料で付属しており、サイズは直線的に増加します。中規模 (2 vCPU/4 GB) は 0.10 ドル/時間、大規模 (4 vCPU/8 GB) は 0.20 ドル/時間です。マシンが 2 倍、価格もちょうど 2 倍になります。アップグレードは決して罰せられません。ディスクもボックスに合わせて同じように拡張できます。月額 0.10 ドル/GB で、Small で最大 50 GB、Medium で 80 GB、Large で 160 GB を予約できます。本当にさらに必要な場合は、ノーと言われる代わりに、2 倍のオーバーフローでバンドルを 2 倍に拡張します。
マシンを十分な期間 (3 日間停止) 保留すると、マシンはコールド アーカイブ ストレージに静かに移動します。マシンはその ID、ファイル、アカウント内の場所を保持し、ホスト リソースが解放され、次回の起動時に数秒ではなく数分で利用可能なホストにマシンを復元します。アーカイブされている間、無料の 10 GB を超える追加ディスクには一律 0.04 ドル/GB が請求されます (すべての主流のスナップショット価格に基づいて、復元料金や最低保持期間はありません)。デフォルトのディスク マシンは、サブスクライブしている限り、ちょうど 0 ドルで利用できます。
すべてのマシンは、独自のユーザー名前空間の sysbox ランタイムの下で実行されます。マシン内の root はホスト上の誰でもありません。エージェントは、物理的にクォータを超えて拡張できないディスク上で非 root ユーザーとして実行されます。シークレットは書き込み専用の暗号化された変数として入力されます。パネルには値が表示されることはなく、マスクされたヒントのみが表示されます。
はい、Docker が内部で実行されます

— すべてのマシン、すべてのプラン
すべてのマシンには本物の Docker エンジンが内蔵されており、無料プランも含まれています。何もチェックする必要はありません。何も設定する必要はありません。 docker build 、スタックの作成、テストコンテナーなどすべて、 --privileged はどこにもありません。 sysbox ランタイムは、各マシンに適切に分離された環境を提供し、通常の dockerd がホストを認識せずに実行できる環境を提供します。
私たちが最も苦労したのはアカウンティングです。ネストされた Docker は通常、クラウドがリークする場所であるためです。内部エンジンが書き込むすべてのバイト (イメージ、レイヤー、ボリューム) は、マシン自身のキャップされた作業ディスクに配置されます。マシン内の df、パネルのディスク バー、および請求書はすべて同じ番号を報告します。ディスクをいっぱいにすると、突然の請求書ではなく、正直なディスク満杯エラーでプルが失敗します。スペースを空けて続行します。イメージとボリュームは、マシン上の他のすべてのものと同様に、サスペンド後も保持されます。
これは私たちが作った私のお気に入りのもので、動作しているときは見えません。
ターンを開始してからタブを閉じます。移動してください。トンネル内で接続が切断されます。ターンはマシン上で実行され続け、リレーは終了するまですべてのステップ (ツール呼び出し、出力、コスト) を記録し続けます。 20分後に携帯電話からセッションを開くと、誰も見ていない間に起こったすべてのことを含む完全な記録がそこにあります。バイトを失わなくなるまでターンの途中でブラウザを強制終了することでこれをテストしました。
このスクリーンショットは、私がこの投稿を書いている間の実際のセッションです。エージェントは Bun サーバーをスキャフォールディングし、起動し、 /health をカールし、その後クリーンアップしました。コーナーの合計コスト: 0.008 ドル — この記事の他のスクリーンショットを撮るために、途中で 2 回移動しました。何も失われませんでした。
同じセッションには、開発者 API ( api.octo ) を介して、パネル、通常の Web 端末からアクセスできます。

min.run/api/v1 — セッションの作成、メッセージの送信、SSE 経由のテール出力、すべて CI からスクリプト化可能)、および Telegram 経由: ワンタップでボットに接続 (トークンを所有したい場合は独自のボットを持参) し、マシンがチャットで応答します。実際のフォーマット、テーブルが含まれており、生のマークダウンの壁はありません。エージェントはブラウザのタブではなくなり、いくつかのドアを持ったあなた自身のものになります。開発者 API と Telegram ワークフローは、それぞれ独自の完全な投稿に値します。それらは今後登場します。これが概要です。
セッションは自分の考え方も覚えています。すべてのセッションはデフォルトのモデルとロールで開始され、どちらもセッションごとに切り替えることができます。セッション ヘッダーのモデル ドロップダウンから選択するか、チャットで /model gpt-5.6-luna と入力します。同じコマンドが Telegram および API を通じて機能し、選択はセッションで持続します。厄介なリファクタリングにはプレミアム モデルを使用し、クリーンアップには安価なモデルに戻します。 /role 、 /effort 、 /info 、 /done 、 /plan 、 /schedule 、および /agents は、すべてのサーフェスから同じレールに乗ります。私たちが誇りに思っている詳細の 1 つは、プランで呼び出すことができないモデルへの切り替えは、切り替え時に明確なメッセージで拒否され、3 メッセージ後には不可解な 403 として発見されないことです。
そして、マシン上のエージェントは、ローカルで実行するのと同じ octomind です。つまり、同じ構成形式、同じサブエージェント委任、同じツールです。パネルの Config ページには、octomind / octocode / octobrain TOML が保持され、作成時にすべてのマシンに焼き付けられます。モデルのルーティングと API キーはクラウドによって挿入されるため、octohub:auto というだけの設定は、ドットファイルにキーを 1 つも持たずにどこでも機能します。ハブは呼び出しの目的ごとに実際のモデルを選択し (セッション、スーパーバイザーのチェック、圧縮はそれぞれ独自のものを取得します)、ユーザーは c に触れることなく、モデル ページからいずれかを再調整します。

図上。
機械はエージェントに自分の場所を伝えます
すべてのマシンに非常に大きな影響を与える小さなファイルがあります。それは、 AGENTS.md 標準に従って、開始コンテキストとしてすべてのセッションにロードされる AGENTS.md です。これは、モデルが推測できないことをエージェントが知る方法です。再起動しても /home/octo は生き残るが /tmp は生き残らないこと、マシンが自動サスペンドしてフォアグラウンドでスリープしたままになっているプロセスよりもスケジュールが優先されること、端末が見えずに電話で答えを読んでいる可能性があること、つまり、成果物は決して表示されないスクロールバックではなく、エージェントが指定したファイルに属します。
新しいマシンにはベースラインがシードされており、2 つのレベルで変更できます。 「構成」→「説明」でアカウントのベースラインを編集すると、その後作成するすべてのマシンがそのバージョンから開始されます。または、マシンの「命令」タブを開いて、ライブ ファイルを編集します。マシンは最初の起動時のコピーを所有し、編集内容は再起動のたびに残ります (シェルも機能します。これはホーム ディレクトリにある実際のファイルです)。そして、これは私たちが発明したものではなく標準であるため、独自の AGENTS.md を保持するリポジトリは正常に機能します。エージェントはリポジトリ内で作業している間はリポジトリの指示に従い、マシンは他の場所にあります。
埋め込みは、システム全体で最も大量のモデル呼び出しです。オクトコードは、インデックスを作成するすべてのリポジトリのすべてのチャンクを埋め込みます。

[切り捨てられた]

## Original Extract

Your coding agent dies with your laptop lid, your API keys live in five dashboards, and your AI bill is a surprise. We built Octomind Cloud to fix all three: persistent machines from $0.05/hr — with real Docker inside — a hub of 21 models behind one key at published per-token prices, and one wallet
[truncated]

Octomind Cloud Is Live: Real Machines for Your Agent, Every Model Through One Key | Octomind
Tap Products Cloud Hub Blog Vision Docs Install Panel GitHub ⭐ 87 Back to Blog Octomind Cloud Is Live: Real Machines for Your Agent, Every Model Through One Key
Your coding agent dies with your laptop lid, your API keys live in five dashboards, and your AI bill is a surprise. We built Octomind Cloud to fix all three: persistent machines from $0.05/hr — with real Docker inside — a hub of 21 models behind one key at published per-token prices, and one wallet with caps you can see.
Today Octomind Cloud goes live in beta. Before I tell you what it is, let me tell you the three problems that made us build it — because odds are you have at least two of them right now.
Your agent lives and dies with your laptop. You kick off a long refactor, the agent is forty minutes into it, and then you need to close the lid, catch a train, or your Wi-Fi hiccups. The session is gone, the context is gone, and the tokens you paid for produced work you'll never see. Agents are long-running processes; laptops are not long-running machines.
Your API keys live in five dashboards. OpenRouter for the cheap models, Anthropic for the good ones, OpenAI because some tool demands it, Voyage for embeddings, and a .env file that's one git add -A away from being public. Every provider has its own billing page, its own rate limits, its own "add $5 minimum" — and no single view of what your agent actually spends.
Your AI bill is a mood, not a number. Seat-priced tools give you rate-limit roulette — an opaque allowance that resets on someone else's schedule. Usage-priced tools give you the opposite fear: nothing stops a runaway agent from burning $80 overnight. Neither shows you a per-token price or lets you watch the meter.
Octomind Cloud is our answer to all three. It's two things that share one subscription, one wallet, and one key: the hub (models) and machines (compute). This post is the map of everything that's live today — each section deserves its own deep dive, and those are coming. But this one piece should tell you exactly what you're getting, economics included, because pricing you can't reason about is a bug, not a business model.
The hub: log in once, every model works
The hub is an OpenAI-compatible gateway in front of a curated model roster — one key, every model works, in the panel, in the CLI, in any tool that speaks the OpenAI API:
curl https:/ /hub.octomind.run/v 1 /chat/completions \
-H "Authorization: Bearer $OCTOMIND_KEY " \
-d '{"model": "deepseek-v4-flash", "messages": [{"role": "user", "content": "hi"}]}' If you're using the CLI you don't have to handle that key at all. As of 0.38.0 , octomind login shows you a short code, opens the panel where you're already signed in, and writes the key to disk itself:
$ octomind login
╭ login · octomind account
│ code TX8T-UC7U
│ url https://octomind.run/app/login/cli
│
│ Confirm the code in your browser to finish signing in.
│ waiting… It's the standard device flow — the same one gh auth login uses — so no password is ever typed into a terminal, each machine gets its own revocable key, and the /usage command tells you what you've spent against your caps without opening the panel.
No provider accounts. No key sprawl. And the part I care most about: published prices, and metering that matches the table exactly.
The roster is deliberately curated, not a 400-model dropdown — twenty-one models that earn their place. The open shelf (DeepSeek V4 Flash and Pro, Qwen3.7 Plus and Max, Kimi 2.6 and 2.7 Code, MiniMax M3, GLM-5.2, Gemma 4) is included in every paid plan and draws from your plan caps. The premium shelf bills per token from prepaid credits, and it's the real thing: Claude Fable 5, Opus 4.8 and Sonnet 5; the whole GPT-5.6 family (Sol, Terra, Luna) with GPT-5.5, 5.5 Pro and the fast little 5.4 mini and nano; Gemini 3.5 Flash and 3.1 Pro. You see every price before you pick. And there's a free model with a daily allowance that keeps working even when everything else is exhausted, because an agent that goes completely dark the moment you hit a cap is worse than a slow one.
Fair use is per account , not per key: Free runs one model call at a time, Pro two in parallel, Max five, Team fifteen pooled. Mint a thousand keys and the budget doesn't move — and past it, requests queue briefly instead of failing. It's the anti-abuse rule that doubles as a noisy-neighbor guarantee: nobody's runaway script makes your completion wait.
Under the hood this is octohub — open source, Rust, one static binary. The hosted hub is the same code with the keys problem solved. Want to run it yourself with your own provider keys? That's not a loophole — it's the point of publishing it.
Machines: a dev box that exists whether or not your laptop does
A machine is a persistent Linux container in our cloud with the whole agent toolchain baked in — your files, your repos, your indexes, persisted on its own disk. It's got: octomind (the agent), octocode (code indexing + semantic search), octobrain (memory), a web terminal, git, Docker (more on that in a minute), and — on the default image — Node, bun, a headless browser, ffmpeg and image tools. You create one in the panel, it's ready in under a minute, and it's yours.
The billing model is the part I want to defend in detail, because it's where most clouds get either opaque or hostile. A machine is always in one of three states, and the state decides what you pay:
Per-second metering, automatic transitions, you configure nothing. The consequence that matters: while your agent waits on the model — which is most of an agent's wall-clock life — you pay a fifth of the active rate. A five-minute session costs about half a cent, plus $0.001 per session start. 10 GB of working disk comes free with every machine, and sizes scale linearly — Medium (2 vCPU/4 GB) is $0.10/hr, Large (4 vCPU/8 GB) is $0.20/hr. Twice the machine, exactly twice the price; upgrading is never punished. Disk scales with the box the same way — reserve up to 50 GB on a Small, 80 on a Medium, 160 on a Large at $0.10/GB-month, and if you genuinely need more, stretch to double the bundle with the overflow at 2× instead of being told no.
Park a machine long enough — three days suspended — and it quietly moves to cold archive storage : the machine keeps its identity, its files and its place in your account, host resources are freed, and the next start restores it onto any available host in a few minutes instead of seconds. While archived, extra disk above the free 10 GB bills a flat $0.04/GB-month — under every mainstream snapshot price, with no restore fee and no minimum retention — and a default-disk machine parks at exactly $0 for as long as you're subscribed.
Every machine runs under the sysbox runtime in its own user namespace — root inside the machine is nobody on the host. The agent runs as a non-root user on a disk that physically cannot grow past its quota. Secrets go in as write-only encrypted variables: the panel can never show a value back, only a masked hint.
Yes, Docker runs inside — every machine, every plan
Every machine comes up with a real Docker Engine inside, free plan included — nothing to tick, nothing to configure: docker build , compose stacks, testcontainers, all of it, with no --privileged anywhere. The sysbox runtime gives each machine a properly isolated environment where an ordinary dockerd can live without ever seeing the host.
The part we sweated hardest is the accounting, because nested Docker is where clouds usually leak: every byte the inner engine writes — images, layers, volumes — lands on the machine's own capped working disk. df inside the machine, the disk bar in the panel, and the bill all report the same number. Fill the disk and a pull fails with an honest disk-full error, not a surprise invoice; free some space and carry on. Your images and volumes persist across suspend like everything else on the machine.
This is my favorite thing we built, and it's invisible when it works.
Start a turn, then close the tab. Navigate away. Lose your connection in a tunnel. The turn keeps running on the machine , and the relay keeps recording every step — tool calls, output, cost — until it finishes. Open the session from your phone twenty minutes later and the full transcript is there, including everything that happened while nobody was watching. We tested this by killing browsers mid-turn until we couldn't lose a byte.
That screenshot is a real session from while I wrote this post: the agent scaffolded a Bun server, started it, curled /health , and cleaned up after itself. Total cost in the corner: $0.008 — and I navigated away twice in the middle to take the other screenshots in this article. Nothing was lost.
The same session is reachable from the panel, from a plain web terminal, through the developer API ( api.octomind.run/api/v1 — create sessions, send messages, tail output over SSE, all scriptable from CI), and through Telegram: one tap connects our bot (or bring your own, if you'd rather own the token), and your machine answers you in chat — real formatting, tables included, not a wall of raw markdown. The agent stops being a browser tab and becomes a thing you own that happens to have several doors. The developer API and the Telegram workflow each deserve a full post of their own — they're coming; this is the overview.
Sessions also remember how they think. Every session starts on your default model and role, and both are switchable per session: pick from the model dropdown in the session header, or type /model gpt-5.6-luna in the chat — the same command works from Telegram and through the API, and the choice persists with the session. Take a premium model for the gnarly refactor, drop back to a cheap one for the cleanup; /role , /effort , /info , /done , /plan , /schedule and /agents ride the same rails from every surface. One detail we're proud of: switching to a model your plan can't call is rejected at switch time with a clear message — not discovered as a cryptic 403 three messages later.
And the agent on the machine is the same octomind you'd run locally — same config format, same sub-agent delegation, same tools. The panel's Config page holds your octomind / octocode / octobrain TOML and bakes it into every machine at creation; model routing and API keys are injected by the cloud, so a config that just says octohub:auto works everywhere without a single key in a dotfile — the hub picks the real model per call purpose (your sessions, the supervisor's checks, compression each get their own), and you retune any of them from the Models page without touching config.
The machine tells the agent where it is
There's a small file with an outsized effect on every machine: AGENTS.md , following the AGENTS.md standard , loaded into every session as opening context. It's how the agent knows the things no model can guess: that /home/octo survives restarts and /tmp doesn't, that the machine auto-suspends and a schedule beats a process left sleeping in the foreground, that you might be reading its answer on a phone with no terminal in sight — so deliverables belong in files it names, not in scrollback you'll never see.
New machines come seeded with our baseline, and it's yours to change at two levels. Edit your account's baseline under Configs → instructions and every machine you create afterwards starts from your version. Or open a machine's Instructions tab and edit the live file — the machine owns its copy from first boot, and your edits survive every restart (a shell works too; it's a real file in the home directory). And because it's the standard rather than something we invented, a repo that carries its own AGENTS.md just works: the agent follows the repo's instructions while working inside it, and the machine's everywhere else.
Embeddings are the highest-volume model call in the whole system — octocode embeds every chunk of every repo it indexes, a

[truncated]
