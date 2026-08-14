---
source: "https://fn2.ai/claude"
hn_url: "https://news.ycombinator.com/item?id=49305190"
title: "Show HN: I built a Claude Code plugin to query 10.6M earnings-call embeddings"
article_title: "FN2 for Claude Code | Research agents from your terminal"
author: "oceanplexian"
captured_at: "2026-08-14T22:12:35Z"
capture_tool: "hn-digest"
hn_id: 49305190
score: 1
comments: 0
posted_at: "2026-08-14T22:09:36Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I built a Claude Code plugin to query 10.6M earnings-call embeddings

- HN: [49305190](https://news.ycombinator.com/item?id=49305190)
- Source: [fn2.ai](https://fn2.ai/claude)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T22:09:36Z

## Translation

タイトル: HN を表示: 1,060 万件の収益コールの埋め込みをクエリするクロード コード プラグインを構築しました
記事のタイトル: クロード コードの FN2 |端末からエージェントを調査する
説明: クロード コードの FN2 を使用します。 142,758 件の通話と 8,186 件のティッカー、リアルタイム価格、および FRED マクロ データからの 1,060 万件の埋め込まれた決算報告のセマンティック検索。引用された調査エージェントをスケジュールに組み込むことができます。または、FN2 MCP コネクタを任意の MCP クライアントに追加します。
HN テキスト: こんにちは、HN。私は 2025 年以降、収益記録のデータベースに対する 1 回限りの市場質問に LLM を使用してきました。テクノロジー業界では長い時間がかかるため、Solo 開発者としてそれを中心にプラットフォームを構築し、その他のさまざまな金融ツールや統合を追加してきました。バックエンドインフラストラクチャは非常に退屈です! PostgreSQL と ElasticSearch。データは 2 つのリージョン間で適切にシャーディングおよびレプリケートされます。リクエストの例としては、「NVIDIA は最後の電話でデータセンターの需要について何と言いましたか?」 「各市場が開く前に NVDA、AMD、AVGO について説明してください」「AAPL が 1 日に 5% 以上動いたら通知してください」このプラグインは依存関係のない Python であり、MIT ライセンスを取得しています。コーパス、ベクトル検索、および定期エージェント サービスはホストされているため、これらを使用するには無料の FN2 アカウントが必要です。私は、スケジュールされたリサーチの部分が人々に役立つかどうかに特に興味がありますが、Claude Code、Hermes Agent、OpenClaw などの統合により、基本機能は無料です。ありがとうございます!

記事本文:
クロードコードのFN2 |端末からエージェントを調査する
FN2.AI
サインアップ
よくある質問
ブログ
収益
ドキュメント
価格設定
外観
ログイン
クロード コードから FN2 を使用する
1,060 万件の埋め込み決算報告文、リアルタイム価格、SEC 提出書類、および FRED マクロ データを直接クロード コードにセマンティック検索します。端末を離れることなく、FN2 リサーチ エージェントを開始、スケジュール設定し、読み取ります。
クロード コード v2.1.224
Opus 5 (1M コンテキスト) xhigh 努力 · Claude Max
/ユーザー/アンドレアス
> NVDA 経営陣は、過去数回の決算発表で Blackwell の供給について実際に何と言ってきましたか?
● fn2 スキルを使用して、収益コール コーパスに対してセマンティック検索を実行します。
● Bash (Python3 "$FN2" 認証ステータス)
⎿ サインイン · plan pro · 認証情報 fn2d_…fGlh (デバイストークン)
● Bash (python3 "$FN2" --json start "最近の決算発表における Blackwell の供給に関する NVDA 経営陣のコメント、引用および引用")
⎿ { "id": "cce9d5fa"、"ステータス": "保留中"、"繰り返し": false }
● Bash (python3 "$FN2" 結果 cce9d5fa dedde0a2)
⎿ 実行は 35 秒で完了 · モデル z-ai/glm-5.2 · 118,982 トークン
142,758 回の呼び出しで 10,602,684 個の埋め込みパッセージを検索 · Qwen3-Embedding-8B (4096d)
… +48 行 (Ctrl+O で展開)
● 通話の 4 分の 4 は 1 つのストーリーを物語っています。ブラックウェルは制約からランプに移行しました。 1 年前、この言語は「需要が供給を大幅に上回っている」という希少性があり、クラウド GPU は完売したと言われていました。
● 最新の発表（2027 年度第 1 四半期）によれば、これは「会社史上最速の製品増加」となっています。データセンターは 750 億ドルを売り上げ、前年比 92% 増となり、ネットワークはほぼ 3 倍になりました。
● 来年もまだ厳しい：次の四半期の誘導額は 910 億ドル、さらに 800 億ドルの自社株買いが見込まれ、中国は見通しをゼロに据え置いた。完全な引用とコール引用は実行レポートに記載されています。
✱ 1分5秒調理
> 大好きです。これまでにこのように簡潔に説明してください

平日のオープン前にエージェントが何をしているのか見せてください
●メール配信でスケジュールを設定します。 「市場が開く前」は Exchange カレンダーのサーバー側に対して解決されるため、週末と休日は自動的にスキップされます。
● Bash (python3 "$FN2" start "NVDA に関する毎日の市場前ブリーフ" -- "市場が開く前" の場合 -- 電子メール)
⎿ ✓ エージェントが作成されました: 5835e096 · アクティブ · メール送信先
リズム: 各市場が開く 30 分前 (アメリカ/ニューヨーク)
次の実行: 8 月 10 日月曜日、東部標準時間 09:00
● Bash (Python3 "$FN2" エージェント)
⎿ ID 名 ステータス スケジュールが実行されます
5835e096 NVIDIA アクティブ マーケットに関する毎日の市場前ブリーフ_open 0
5a415c62 市場オープンセンチメント ac
[切り捨てられた]
FN2 は 2 つの正面玄関を通ってクロード コードに到達します。スキルを使用して端末からエージェントを駆動するか、MCP コネクタを追加して FN2 研究データを任意の MCP クライアントに公開します。
ターミナルからエージェントを操作する Claude Code プラグイン。 API キーを管理することなく、自然言語で作成、スケジュール、実行、一時停止、結果の読み取りを行うことができます。
FN2 MCP サーバーを MCP 互換クライアントに追加し、FN2 調査データをモデルに直接公開します。142,758 件の決算報告と 8,186 件のティッカー、さらに価格、申告書、マクロ シリーズにわたるトランスクリプト検索が可能です。読み取り専用で、エージェント管理はありません。
pip インストール、ビルド手順、貼り付けるシークレットはありません。このスキルは、依存関係のない 1 つの Python 3 ファイル ツリーです。
マーケットプレイスを追加してインストールする
fn2-skill リポジトリを Claude Code マーケットプレイスとして追加し、プラグインをインストールします。
$ claude プラグイン マーケットプレイスに fn2ai/fn2-skill を追加
$ クロードプラグインインストール fn2@fn2
2
アカウントを接続する
ログインコマンドを実行します。端末はショートコードを出力します。 fn2.ai/start を開いてサインインし、コードを入力して承認します。 CLI は、キーチェーンに保存されているスコープ指定された取り消し可能なデバイス認証情報を受け取ります。秘密を見たりコピーしたりすることは決してありません。
$ fn2 認証ログイン
# https://fn2 を開く

.ai/start とコード XXXX-XXXX を入力します。
$ fn2 whoami # プラン + クォータを確認します
3
アシスタントに聞いてください
クロード・コードと自然言語で会話します。これは、エージェントを開始し、結果を確認し、エージェントをスケジュールに組み込むためのスキルを呼び出します。
あなたは › 通常とは異なるオプション フローを追跡するエージェントを開始します
私のウォッチリストに毎時
↳ ✓ エージェントが作成され、1 時間ごとに実行され、結果が電子メールで送信されます
毎週の研究費
すべてのプランには週ごとのリサーチ コールの割り当てがあり、月曜日の 00:00 UTC にリセットされます。調査通話ではモデル トークンが消費されないため、これは毎月のチャット割り当てとは別のものになります。
割り当ては月曜日の 00:00 UTC にリセットされます。エージェントの実行とチャット ターンは、毎月のトークン クォータから引き出されます。研究ツールの呼び出しは行われません。
FN2 と Claude Code で構築を開始する
無料のアカウントを作成し、スキルをインストールすると、数分で最初のスケジュールされたエージェントが実行されます。クレジットカードは必要ありません。
Claude Code は Anthropic、PBC の製品です。 FN2 は独立した製品であり、Anthropic と提携、承認、後援されていません。 MCP はモデルに依存しないプロトコルです。
提供される結果は教育のみを目的としており、財務上のアドバイスとして考慮されるべきではありません。
利用規約 |
プライバシーポリシー |
|
ステータス |
FN2合同会社

## Original Extract

Use FN2 from Claude Code. Semantic search across 10.6 million embedded earnings-call passages from 142,758 calls and 8,186 tickers, real-time prices, and FRED macro data, with cited research agents you can put on a schedule. Or add the FN2 MCP connector to any MCP client.

Hi HN, I've been using LLMs for one-off market questions against a database of earnings transcripts since ~2025. In tech years that's a long time so as a Solo dev I built a platform around it and have been adding miscellaneous other financial tools and integrations. Backend infrastructure is super boring! PostgreSQL & ElasticSearch, with the data appropriately sharded and replicated across two regions. Some example requests: "What did NVIDIA say about data-center demand on its last call?" "Brief me on NVDA, AMD and AVGO before each market open" "Notify me when AAPL moves more than 5% in a day." The plugin is dependency-free Python and MIT licensed. The corpus, vector search and scheduled-agent service are hosted, so using them requires a free FN2 account. I’m particularly interested in whether people find the scheduled-research part useful, but the basic features are free with integrations for Claude Code, Hermes Agent, OpenClaw, etc. Thanks!

FN2 for Claude Code | Research agents from your terminal
FN2 .AI
Sign Up
FAQ
Blog
Earnings
Docs
Pricing
Appearance
Log in
Use FN2 from Claude Code
Semantic search across 10.6 million embedded earnings-call passages, real-time prices, SEC filings, and FRED macro data, straight into Claude Code. Start, schedule, and read FN2 research agents without leaving your terminal.
Claude Code v2.1.224
Opus 5 (1M context) with xhigh effort · Claude Max
/Users/andreas
> what has NVDA management actually said about Blackwell supply on the last few earnings calls?
● I'll run a semantic search over the earnings-call corpus with the fn2 skill.
● Bash (python3 "$FN2" auth status)
⎿ signed in · plan pro · credential fn2d_…fGlh (device token)
● Bash (python3 "$FN2" --json start "NVDA management commentary on Blackwell supply across recent earnings calls, quoted and cited")
⎿ { "id": "cce9d5fa", "status": "pending", "recurring": false }
● Bash (python3 "$FN2" result cce9d5fa dedde0a2)
⎿ run completed in 35s · model z-ai/glm-5.2 · 118,982 tokens
searched 10,602,684 embedded passages across 142,758 calls · Qwen3-Embedding-8B (4096d)
… +48 lines (ctrl+o to expand)
● Four quarters of calls tell one story: Blackwell went from constraint to ramp. A year ago the language was scarcity, with "demand well ahead of supply" and cloud GPUs described as sold out.
● By the latest call (FY27 Q1) it's the "fastest product ramp in company history" . The numbers back the quote: Data Center did $75B , up 92% YoY, with networking nearly tripling.
● Still tight into next year: $91B guided next quarter plus an $80B buyback , and China held at zero in the outlook. Full quotes and call citations are in the run report.
✱ Cooked for 1m 5s
> love it. brief me like this every weekday before the open, and show me what my agents are up to
● Scheduling it with email delivery. "before market open" resolves against the exchange calendar server-side, so it skips weekends and holidays automatically.
● Bash (python3 "$FN2" start "Daily pre-market brief on NVDA" --when "before market open" --email)
⎿ ✓ agent created: 5835e096 · active · email on
cadence: 30 min before each market open (America/New_York)
next run: Mon Aug 10, 09:00 ET
● Bash (python3 "$FN2" agents)
⎿ ID NAME STATUS SCHEDULE RUNS
5835e096 Daily pre-market brief on NVIDIA active market_open 0
5a415c62 Market Open Sentiment ac
[truncated]
FN2 reaches Claude Code through two front doors. Use the skill to drive agents from the terminal, or add the MCP connector to expose FN2 research data to any MCP client.
A Claude Code plugin that drives your agents from the terminal. Create, schedule, run, pause, and read results in natural language, with no API key to manage.
Add the FN2 MCP server to any MCP-compatible client and expose FN2 research data directly to the model: transcript search across 142,758 earnings calls and 8,186 tickers, plus prices, filings, and macro series. Read-only, no agent management.
No pip install, no build step, no secret to paste. The skill is one dependency-free Python 3 file tree.
Add the marketplace and install
Add the fn2-skill repository as a Claude Code marketplace, then install the plugin.
$ claude plugin marketplace add fn2ai/fn2-skill
$ claude plugin install fn2@fn2
2
Connect your account
Run the login command. The terminal prints a short code; open fn2.ai/start, sign in, type the code, and approve. The CLI receives a scoped, revocable device credential stored in your keychain. You never see or copy a secret.
$ fn2 auth login
# Open https://fn2.ai/start and enter code: XXXX-XXXX
$ fn2 whoami # confirms plan + quota
3
Ask your assistant
Talk to Claude Code in natural language. It calls the skill to start agents, check results, and put them on a schedule.
you › start an agent that tracks unusual options flow
in my watchlist hourly
↳ ✓ agent created, runs hourly, results emailed
Weekly research allowance
Every plan gets a weekly research-call allowance that resets Monday 00:00 UTC. Research calls spend no model tokens, so this is separate from your monthly chat quota.
Allowance resets Monday 00:00 UTC. Agent runs and chat turns draw from your monthly token quota; research tool calls do not.
Start building with FN2 and Claude Code
Create a free account, install the skill, and have your first scheduled agent running in minutes. No credit card required.
Claude Code is a product of Anthropic, PBC. FN2 is an independent product and is not affiliated with, endorsed by, or sponsored by Anthropic. MCP is a model-agnostic protocol.
Results provided are for educational purposes only and should not be considered as financial advice.
Terms of Service |
Privacy Policy |
|
Status |
FN2 LLC
