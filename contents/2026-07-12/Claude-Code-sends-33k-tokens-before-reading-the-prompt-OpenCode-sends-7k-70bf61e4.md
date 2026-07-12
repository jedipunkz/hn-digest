---
source: "https://systima.ai/blog/claude-code-vs-opencode-token-overhead"
hn_url: "https://news.ycombinator.com/item?id=48883275"
title: "Claude Code sends 33k tokens before reading the prompt; OpenCode sends 7k"
article_title: "Claude Code Sends 4.7x More Tokens Than OpenCode Before Reading Your Prompt | Systima Blog"
author: "systima"
captured_at: "2026-07-12T18:51:43Z"
capture_tool: "hn-digest"
hn_id: 48883275
score: 9
comments: 0
posted_at: "2026-07-12T18:25:51Z"
tags:
  - hacker-news
  - translated
---

# Claude Code sends 33k tokens before reading the prompt; OpenCode sends 7k

- HN: [48883275](https://news.ycombinator.com/item?id=48883275)
- Source: [systima.ai](https://systima.ai/blog/claude-code-vs-opencode-token-overhead)
- Score: 9
- Comments: 0
- Posted: 2026-07-12T18:25:51Z

## Translation

タイトル: クロード コードは、プロンプトを読む前に 33,000 個のトークンを送信します。 OpenCode が 7k を送信
記事のタイトル: Claude Code は、プロンプトを読む前に OpenCode よりも 4.7 倍多くのトークンを送信します |システィマのブログ
説明: API 境界で測定されたクロード コードと OpenCode トークンのオーバーヘッド。すぐに使えるベースライン、命令ファイルの重み、MCP スキーマ税、サブエージェント乗数、およびキャッシュ書き込み動作。
HN text: これは予感に基づいて始まりました。私たちは通常 OpenCode を使用していますが、Meridian の問題により、しばらくの間 Claude Code を使用することを「強制」されました。その間、OpenCode を使用したときよりもはるかに速く、使用量メーターが上昇するのがわかりました。これは最初の事例証拠でしたが、私たちは経験的データを収集するためにこの小規模な調査に取り組みました。エージェント コーディング ツール (Claude Code と OpenCode) と Anthropic のエンドポイントの間にログを追加し、すべてのリクエスト (および返された使用ブロック) をキャプチャしました。注意点が 1 つあります (投稿の終わりの方) が、キャッシュ戦略とハーネス トークンの使用に関して、Claude Code は OpenCode よりもはるかに非効率であることが明確にわかりました。

記事本文:
Claude Code は、プロンプトを読む前に OpenCode よりも 4.7 倍多くのトークンを送信します。 Systima ブログ Systima サービス Agentic AI プラットフォーム 本番グレードのマルチエージェント システム フラクショナル AI ヘッド オンデマンドの AI 上級リーダーシップ AI ガバナンスとコンプライアンス EU AI 法と規制の枠組み AI 戦略とロードマップ ビジョンから実行計画まで LLM コストの最適化 LLM 支出の削減 AI デューデリジェンス 投資家のための技術的な AI 評価 量子ガバナンス 量子に関する管理された取締役会の立場
ソリューション LegalTech 向け AI 法務ワークフロー向け AI システム SaaS および B2B 向け AI 製品への AI の組み込み
Agentic AI プラットフォーム AI のフラクショナル ヘッド AI ガバナンスとコンプライアンス AI 戦略とロードマッピング LLM コスト最適化 AI デュー デリジェンス 量子ガバナンス ソリューション リーガルテック向け AI SaaS および B2B 向け AI 会社概要 ブログ 電話の予約
オープンコード
オープンソースのエージェントコーディング
クロード・コード
ANTHROPIC のフロンティア製品
VS
Applied Research LLM Engineering Agentic AI Claude コードは OpenCode よりもはるかにトークンを大量に消費します。正確にどれくらいかを測定しました
OK を言うことによる固定オーバーヘッド
ギャップが縮まる複数ステップのタスク
乗算器 1. 命令ファイル
乗算器 3. フレームワーク テンプレート
乗数 5. 拡張された思考
キャッシュの安定性、決定的な証拠
ドッグフーディング、または監査ログとしてのベンチマーク データセット
新しい投稿が受信箱に配信されます。
スパムはありません。いつでも購読を解除してください。
LinkedIn でフォローしてください。エージェント AI、EU AI 法のコンプライアンス、規制産業における自動化に関する洞察を常に入手してください。
Claude Code と OpenCode を同じモデル、同じマシン、同じタスクに配置し、送受信されたものをすべて調べました。
両方のハーネスに 1 行の応答を要求したところ、クロード コードは、プロンプトが到着する前に、システム プロンプト、ツール スキーマ、および注入されたスキャフォールディングの約 33,000 トークンを使用しました。オペ

enCodeは約7,000を使用しました。
クロード コードはキャッシュ効率がはるかに低いです。
OpenCode のリクエスト プレフィックスは、キャプチャしたすべての実行でバイト同一でした。ペイロードをセッションごとに 1 回キャッシュし、それを 2 セントで読み戻すのに費用がかかりました。
一方、Claude Code は、セッション中に数万のプロンプト キャッシュ トークンを何度も書き直し、同じタスクで OpenCode の最大 54 倍のキャッシュ トークンを書き込みました。
もちろん、キャッシュ書き込みには割増料金が請求されます。これは、Claude Code を使用するときにダッシュボードの使用量が増加することを考慮したものです。
Config を使用すると、プロンプトがさらに肥大化します。
実稼働リポジトリの 72 KB 命令 (AGENTS.md または CLAUDE.md) ファイルは、すべてのリクエストにさらに (平均) 20,000 トークンを追加します。 5 つの控えめな MCP サーバーでは、さらに 5,000 ～ 7,000 が追加されます。実際に動作するセットアップが最初のリクエストを送信するまでに、ユーザーが単語を入力するまでの深さは 75,000 ～ 85,000 トークンになります。
各サブエージェントには独自のブートストラップ コストがあり、親がそのトランスクリプトを消費するため、直接実行するのに 121,000 トークンのコストがかかる小規模なタスクは、2 つのサブエージェントにファンアウトすると 513,000 トークンのコストがかかります。
クロード コードを支持する結果が 1 つ見つかりました。
マルチステップタスクでは、Claude Code のタスク全体の合計は OpenCode よりも低くなりました。これは、OpenCode がより少ないベースラインをターンごとに再支払いする一方で、ツール呼び出しをより少ないリクエストにバッチ処理するためです。メーターは高く始まります。セッションの展開によって、誰がより多くの費用を費やすかが決まります。
この投稿の残りの部分では、API 境界でこれらすべてをどのように測定したか、トークンがどこに送信されるか、プロンプト キャッシュによって何が節約されるのか、何が節約されないのかを示します。
トークンのオーバーヘッドは、コスト、レイテンシー、およびコンテキストのバジェットです。ハーネス ペイロードのすべてのトークンは、コードに費やすことができない作業コンテキストのトークンであり、ベースラインは毎回再送信されるか、キャッシュから再読み取りされます。
本番環境でエージェント AI を運用する場合、特に

EU AI 法では、第 12 条でシステムの動作を記録して理解することが求められています。「エージェントが実際に何を送信するか」という質問は、民間伝承ではなくデータで答えられるべきです。
各ハーネスとモデル エンドポイントの間にログ プロキシを接続しました。
ハーネス (クロードコード / OpenCode)
→ ロギングプロキシ (リクエストペイロード + レスポンスの使用状況をキャプチャ)
→ モデルエンドポイント プロキシはリクエストごとに 2 つの内容を記録します。 1 つ目は、ハーネスが出力した正確な JSON ペイロード、つまりシステム ブロック、ツール スキーマ、メッセージです。 2 つ目は、API が返した使用ブロックで、入力トークン、キャッシュ書き込み、キャッシュ読み取り、出力トークンをカバーします。
ペイロード キャプチャは、ハーネスが送信するものの正確な情報です。使用量ブロックは、計測された内容の正確な情報です。
このような条件でテストを行いました。
ハーネス。 Claude Code 2.1.207 と OpenCode 1.17.18、どちらも claude-sonnet-4-5 に固定 (2026 年 7 月)。
ベースライン分離。 MCP サーバー、ユーザー設定、メモリのない新しい構成ディレクトリ。指示ファイルのない空のワークスペース。権限がバイパスされました。次に、乗算レーンは一度に 1 つの変数を追加します。
タスク。 T1 は「正確に応答してください: OK」と表示し、固定オーバーヘッド (ハーネスごとに 3 回の実行) を分離します。 T2 はシードされたファイルを読み取り、それを要約します。 T3 は、FizzBu​​zz に対する書き込み、実行、テスト、修正のループとチェッカー スクリプトです。
ゼロツールのバリエーション。 --tools "" を使用した Claude Code と "tools": {"*": false} を使用した OpenCode で、システム プロンプトをツール スキーマの重みから分離します。
数字の前に正直に言っておきます。私たちのトラフィックは、リクエストを独自のエンベロープにラップするローカル LLM ゲートウェイを通過します。この定数は、ベア キャリブレーション リクエストで約 6,200 トークンで測定され、以下のすべての計測数値から差し引かれています。ペイロード レベルの数値は、ゲートワがキャプチャしたリクエスト ボディから取得されます。

y は影響を与えることができず、正確です。
コンポーネント推定の文字からトークンへの変換では、一般的なヒューリスティックではなく、従量制書き込みが全ペイロードに等しいコールド キャッシュ アンカーから導出された、各ハーネスが独自に測定したトークンあたり 4.1 ～ 4.4 文字の比率が使用されます。
OK を言うことによる固定オーバーヘッド
課題は22文字でした。各ハーネスが最初のリクエストで送信した内容は次のとおりです。
7,997 文字の <system-reminder> ブロック
最初のターンのペイロード (校正済み)
OpenCode の要求は最小限に近いです。 「あなたは OpenCode、地球上で最高のコーディング エージェントです」で開く 1 つのシステム ブロック、10 の古典的なコーディング ツール、および唯一のユーザー コンテンツとしてのプロンプトがあります。
Claude Code のリクエストはプラットフォームのブートストラップです。 27 のツールには、コーディング コアに加え、CronCreate や Monitor からタスク ファミリ、ワークツリー管理、プッシュ通知に至るまでのバックグラウンド エージェントおよびオーケストレーション スイート全体が含まれます。
プロンプトの前に、最初のユーザー メッセージには 3 つのリマインダー ブロックが挿入されています。委任用のエージェント タイプのカタログ、利用可能なスキルのカタログ、およびユーザー コンテキスト。
ツール スキーマは、両方の主要な用語です。 Claude Code の約 33,000 トークンのうち約 24,000 がツール定義であるのに対し、OpenCode の約 6,900 トークンのうち約 4,800 がツール定義です。
ツールを削除すると、システム プロンプト自体が分離されます。クロード コードは 26,891 文字、約 6.5,000 トークンになります。 OpenCode は 8,811 文字、約 2.0k トークンです。
ツールが無効になっている場合、両方のハーネスでプロンプトがわずかにトリミングされます。ツールをまったく使用しない場合でも、Claude Code の命令セットのサイズは OpenCode の 3 倍を超えます。残りは行動原則であり、トーン ルール、安全指導、タスク管理の指示、環境の説明を意味します。
T2 は、各ハーネスにファイルを読み取って要約するように依頼しました。どちらも正しい要約を生成しました。
クロード

コードには 6 つの HTTP リクエストと約 199,000 の累積計測入力トークンが必要でした。 OpenCode は 4 件のリクエストと約 41,000 件のリクエストに加え、セッション タイトルを求める Haiku 側のコールを 1 件受け取りました。
これらのトークンのほとんどは、入力価格の 10 分の 1 で請求されるキャッシュ読み取りです。ペイロードに関係なく、3 つのことがスケーリングされます。最初のターンのキャッシュ書き込み、ターンごとの読み取り、およびコンテキスト ウィンドウの消費量 (キャッシュ割引なしでは減少しません)。
33,000 トークンのベースラインとは、コードが会話に入る前に、毎ターンが 200,000 ウィンドウの 6 分の 1 の時点で開始されることを意味します。
ギャップが縮まる複数ステップのタスク
T3 (書き込み、実行、テスト、修正のループ) は、ベースラインによって設定された期待を逆転させました。
1往復の並列バッチ
Claude Code は、ジョブ全体 (2 つのファイル書き込みと 2 つのスクリプト実行) を 1 つの並列ツールの往復にバッチ処理しました。 OpenCode は 1 ターンあたり 1 回のツール呼び出しを行い、9 回かかりました。
ベースラインはリクエストごとに再送信されるため、リクエスト数はベースラインを乗算します。 OpenCode は約 7,000 のベースラインを 9 回支払い、Claude Code は約 33,000 のベースラインを 3 回支払い、合計は収束しました。
タスク全体の入力は、ベースラインとリクエスト数に会話の増加を加えたものにほぼ等しくなります。積極的にバッチ処理する大規模なベースライン ハーネスとシリアル化する小規模なベースライン ハーネスは、同じ場所に着地する可能性があります。
ペイロードから 2 つの構造の詳細が明らかになりました。 Claude Code は会話の進行に応じて追加の <system-reminder> ブロックを挿入します。最初のターンで 3 つ、最初のツールの往復までに 4 つなので、その足場はターン数とともに増加します。 OpenCode のターンあたりの限界ペイロード (ターンあたり約 400 ～ 2,200 文字) は、純粋な会話コンテンツです。
フロアは、無駄なく始まり、短く続くセッションについて説明します。実際のセッションではどちらも行われません。実際の使用量が上に重なる各層を測定しました。
乗算器 1. 命令ファイル
私たちは落とします

実際の 72KB AGENTS.md を運用リポジトリからワークスペースにダウンロードし、T1 を再実行します。
効果は左右対称で大きい。どちらのハーネスも、リクエストごとに 20,000 トークンをわずかに超えるトークンを取得しました。 OpenCode の従量制の合計は 13,152 から 33,336 になりました。クロード・コードは39,005から59,243になりました。
非対称性はメカニズムにあり、実験中にそれが私たちを悩ませました。 Claude Code 2.1.207 は AGENTS.md を完全に無視し、 CLAUDE.md の名前が変更されたときにのみファイルを取り込み、最初のユーザー メッセージに挿入しました。 OpenCode はどちらかのファイル名を読み取り、システム プロンプトに挿入します。
2 つの実際的な結果が続きます。無視された命令ファイルはサイレントであるため、ハーネスが実際にどのファイル名を尊重するかを確認してください。そして、重い命令ファイルは無駄のないハーネスのベースラインをほぼ 4 倍にすることを知ってください。そのリポジトリ内のすべてのセッションのすべてのリクエストに応じます。
認証情報不要のパブリック MCP サーバーを 1 サーバー構成と 5 サーバー構成で接続しました。
スキーマはハーネス全体で同一であるため、税金もほぼ同一になります。小規模サーバーあたり、リクエストごとにおよそ 1,000 ～ 1,400 トークン。 5 つのサーバーでは、ペイロードごとに 4,900 トークンが Claude Code に追加され、OpenCode には 6,967 トークンが追加され、ツール数は 27 から 69、10 から 52 に増加しました。
小規模なパブリック サーバーは穏やかなケースです。豊富な API を備えた運用サーバーでは、数倍の大きさのスキーマが出荷されます。これは、以下のすべての測定結果が示すとおりです。
操作上の脚注が 1 つあります。 Claude Code は、明示的な --mcp-config フラグが渡されるまで、印刷モードでプロジェクト スコープの .mcp.json を無視しました。サーバーが接続されていると仮定した場合は、境界でそれを確認してください。
乗算器 3. フレームワーク テンプレート
BMAD などのストーリー主導のワークフロー フレームワークは、スラッシュ コマンドをペルソナ、プロトコル、チェックリストの大きなプロンプト テンプレートに拡張します。
8,405 文字の表現を実行しました

両方のハーネスの同じ T3 ストーリーのプロンプトとしてネイティブ テンプレートを使用します。テンプレート自体はわずか約 2,100 トークンですが、会話履歴に追加され、セッション内の後続のリクエストごとに再伝送されます。 9 リクエストのセッションでは、リクエストが 9 回再送信されます。
フレームワークの税は、テンプレートのサイズとリクエスト数の積であり、上記のすべてに加えて加算されます。
各ハーネスに、2 つの並列サブエージェントに作業を展開するように依頼しました。ここで合計が爆発します。
Claude Code は、3 つの異なるリクエスト クラスにわたる 9 つのモデル リクエストでタスクを完了しました。約 33,000 の完全なベースラインを持つメイン セッションと、それぞれ 3,554 文字のエージェント システム プロンプトと 27 のツールのうち 24 の独自のブートストラップを実行する 5 つのサブエージェント呼び出しがありました。
累積の計測入力は 513,000 トークンに達しました。これに対し、直接行われた同じ作業の場合は 121,000 トークンでした。これは、1 つの控えめなファンアウトに対して 4.2 倍の乗数になります。これは、すべてのサブエージェントが独自のブートストラップを支払い、そのトランスクリプトが親によって取り込まれるためです。
ここでの OpenCode の設計は、著しく無駄がありません。そのサブエージェント要求には、1,379 文字のシステム プロンプトと 5 つのツールの縮小プロファイルが含まれます。そのサブエージェント レーンはゲートウェイを介して完全に完了しなかったため、キャプチャとの設計の違いを報告します。

[切り捨てられた]

## Original Extract

Claude Code vs OpenCode token overhead measured at the API boundary. Out-of-the-box baselines, instruction file weight, MCP schema tax, subagent multipliers, and cache-write behaviour.

This started based off of a hunch. We usually use OpenCode, but were 'forced' to use Claude Code for a while due to issues with Meridian. In that time, we saw the usage meter rise much, much more quickly than when using OpenCode. This was the initial anecdotal evidence, but we undertook this small study to collect empirical data: We added logging between the agentic coding tool (Claude Code and OpenCode) and Anthropic's endpoint, and captured all requests (and the returned usage blocks). With one caveat (toward the end of the post) we found unambiguously that Claude Code was far more inefficient in terms of its cache strategy and its harness token usage than OpenCode.

Claude Code Sends 4.7x More Tokens Than OpenCode Before Reading Your Prompt | Systima Blog Systima Services Agentic AI Platforms Production-grade multi-agent systems Fractional Head of AI Senior AI leadership on demand AI Governance & Compliance EU AI Act and regulatory frameworks AI Strategy & Roadmapping From vision to execution plan LLM Cost Optimisation Cut LLM spend AI Due Diligence Technical AI assessment for investors Quantum Governance A governed board position on quantum
Solutions AI for LegalTech AI systems for legal workflows AI for SaaS & B2B Embed AI into your product
Agentic AI Platforms Fractional Head of AI AI Governance & Compliance AI Strategy & Roadmapping LLM Cost Optimisation AI Due Diligence Quantum Governance Solutions AI for LegalTech AI for SaaS & B2B About Blog Book a Call
OpenCode
OPEN-SOURCE AGENTIC CODING
Claude Code
ANTHROPIC'S FRONTIER OFFERING
VS
Applied Research LLM Engineering Agentic AI Claude Code Is Way More Token-Hungry Than OpenCode. We Measured Exactly How Much
The fixed overhead of saying OK
A multi-step task, where the gap closes
Multiplier 1. The instruction file
Multiplier 3. Framework templates
Multiplier 5. Extended thinking
Cache stability, the smoking gun
Dogfooding, or the benchmark dataset as an audit log
Get new posts delivered to your inbox.
No spam. Unsubscribe at any time.
Follow us on LinkedIn Stay ahead with insights on agentic AI, EU AI Act compliance, and automation in regulated industries.
We put Claude Code and OpenCode on the same model, the same machine, and the same tasks, then examined everything sent and received.
When we asked both harnesses for a one-line reply, Claude Code used roughly 33,000 tokens of system prompt, tool schemas, and injected scaffolding before the prompt even arrived. OpenCode used about 7,000.
Claude Code is far more cache inefficient:
OpenCode's request prefix was byte-identical in every run we captured; it paid to cache its payload once per session and read it back for pennies.
Claude Code on the other hand re-wrote tens of thousands of prompt-cache tokens mid-session, run after run, and on the same task wrote up to 54x more cache tokens than OpenCode .
Cache writes of course are billed at a premium, which accounted for the usage dashboard climbing when using Claude Code.
Config further bloats the prompt:
A production repository's 72KB instruction (AGENTS.md or CLAUDE.md) file adds another (avg) 20,000 tokens to every single request. Five modest MCP servers add 5,000 to 7,000 more. By the time a real working setup sends its first request, it is 75,000 to 85,000 tokens deep before the user has typed a word.
A small task that cost 121,000 tokens done directly cost 513,000 tokens when fanned out to two subagents , because every subagent has its own bootstrap cost, and the parent then consumes its transcript.
We found one result in favour of Claude Code:
On a multi-step task Claude Code's whole-task total came out lower than OpenCode's, because it batches tool calls into fewer requests while OpenCode re-pays its smaller baseline turn after turn. The meter starts higher; how the session unfolds decides who spends more.
The rest of this post shows how we measured all of this at the API boundary, where the tokens go, and what prompt caching does and does not save you.
Token overhead is cost, latency, and context budget. Every token of harness payload is a token of working context you cannot spend on code, and the baseline is re-sent, or re-read from cache, on every single turn.
If you operate agentic AI in production, particularly under the EU AI Act where Article 12 expects you to log and understand your system's behaviour, "what does my agent actually send" is a question you should be able to answer with data rather than folklore.
We spliced a logging proxy between each harness and the model endpoint.
harness (Claude Code / OpenCode)
→ logging proxy (captures request payloads + response usage)
→ model endpoint The proxy records two things per request. The first is the exact JSON payload the harness emitted, meaning the system blocks, tool schemas, and messages. The second is the usage block the API returned, covering input tokens, cache writes, cache reads, and output tokens.
The payload capture is ground truth for what the harness sends. The usage block is ground truth for what was metered.
We tested under these conditions.
Harnesses. Claude Code 2.1.207 and OpenCode 1.17.18, both pinned to claude-sonnet-4-5 , July 2026.
Baseline isolation. Fresh config directories with no MCP servers, no user settings, and no memory; an empty workspace with no instruction files; permissions bypassed. Multiplier lanes then add one variable at a time.
Tasks. T1 says "Reply with exactly: OK" and isolates fixed overhead (three runs per harness). T2 reads a seeded file and summarises it. T3 is a write-run-test-fix loop against FizzBuzz plus a checker script.
Zero-tools variant. Claude Code with --tools "" and OpenCode with "tools": {"*": false} , separating system prompt from tool schema weight.
One honesty note before the numbers. Our traffic passes through a local LLM gateway that wraps requests in its own envelope, a constant we measured at roughly 6,200 tokens with bare calibration requests and subtracted from every metered figure below. Payload-level figures come from the captured request bodies, which the gateway cannot affect, and are exact.
Character-to-token conversion for component estimates uses each harness's own measured ratio of 4.1 to 4.4 characters per token, derived from cold-cache anchors where the metered write equals the full payload, rather than a generic heuristic.
The fixed overhead of saying OK
The task was 22 characters. Here is what each harness sent with it on its first request.
7,997 chars of <system-reminder> blocks
First-turn payload (calibrated)
OpenCode's request is close to minimal. There is one system block that opens with "You are OpenCode, the best coding agent on the planet", ten classic coding tools, and your prompt as the only user content.
Claude Code's request is a platform bootstrap. The 27 tools include the coding core plus an entire background-agent and orchestration suite, from CronCreate and Monitor to the Task family, worktree management, and push notifications.
Before your prompt, its first user message carries three injected reminder blocks; a catalogue of agent types for delegation, a catalogue of available skills, and user context.
Tool schemas are the dominant term for both. Roughly 24,000 of Claude Code's ~33,000 tokens are tool definitions, versus roughly 4,800 of OpenCode's ~6,900.
Stripping the tools isolates the system prompt itself. Claude Code's weighs in at 26,891 chars, about 6.5k tokens. OpenCode's is 8,811 chars, about 2.0k tokens.
Both harnesses trim their prompt slightly when tools are disabled. Even with no tools at all, Claude Code's instruction set is over three times the size of OpenCode's; the residual is behavioural doctrine, meaning tone rules, safety guidance, task-management instructions, and environment description.
T2 asked each harness to read a file and summarise it. Both produced correct summaries.
Claude Code took 6 HTTP requests and roughly 199,000 cumulative metered input tokens. OpenCode took 4 requests and roughly 41,000, plus one Haiku side call for session titling.
Most of those tokens are cache reads billed at a tenth of the input price. Three things scale with payload regardless; the first-turn cache write, the per-turn read, and context-window consumption, which no cache discount reduces.
A 33k-token baseline means every turn starts a sixth of the way into a 200k window before any code enters the conversation.
A multi-step task, where the gap closes
T3, the write-run-test-fix loop, inverted the expectation set by the baselines.
parallel batch in one round trip
Claude Code batched the entire job, two file writes and two script executions, into a single parallel tool round trip. OpenCode made exactly one tool call per turn and took nine.
Because the baseline is re-sent on every request, request count multiplies baseline. OpenCode paid its ~7k baseline nine times, Claude Code paid its ~33k three times, and the totals converged.
Whole-task input roughly equals baseline times request count, plus conversation growth. A large-baseline harness that batches aggressively and a small-baseline harness that serialises can land in the same place.
Two structural details emerged from the payloads. Claude Code injects an additional <system-reminder> block as the conversation progresses, three on the first turn and four by the first tool round trip, so its scaffolding grows with turn count. OpenCode's per-turn marginal payload, roughly 400 to 2,200 chars per turn, is pure conversation content.
The floor explains a session that starts lean and stays short. Real sessions do neither. We measured each layer that real usage stacks on top.
Multiplier 1. The instruction file
We dropped a real 72KB AGENTS.md from a production repository into the workspace and re-ran T1.
The effect is symmetrical and large. Both harnesses gained just over 20,000 tokens per request. OpenCode's metered total went from 13,152 to 33,336. Claude Code's went from 39,005 to 59,243.
The asymmetry is in the mechanics, and it bit us during the experiment. Claude Code 2.1.207 ignored AGENTS.md entirely and only ingested the file when renamed CLAUDE.md , injecting it into the first user message. OpenCode reads either filename and injects it into the system prompt.
Two practical consequences follow. Check which filename your harness actually honours, because an ignored instruction file is silent. And know that a heavy instruction file nearly quadruples a lean harness's baseline; it rides on every request of every session in that repository.
We attached public, credential-free MCP servers in one-server and five-server configurations.
The schemas are identical across harnesses, so the tax is nearly identical too; roughly 1,000 to 1,400 tokens per small server, per request . Five servers added 4,900 tokens to Claude Code by payload and 6,967 metered to OpenCode, growing the tool counts from 27 to 69 and from 10 to 52.
Small public servers are the gentle case. Production servers with rich APIs ship schemas several times larger, which is exactly what the everything measurement below shows.
One operational footnote. Claude Code silently ignored a project-scoped .mcp.json in print mode until passed an explicit --mcp-config flag. If you assume a server is attached, verify it at the boundary.
Multiplier 3. Framework templates
Story-driven workflow frameworks such as BMAD expand a slash command into a large prompt template of personas, protocols, and checklists.
We ran an 8,405-char representative template as the prompt for the same T3 story in both harnesses. The template itself is only about 2,100 tokens, but it enters the conversation history and is re-carried by every subsequent request in the session . A 9-request session re-sends it nine times.
Framework tax is template size times request count, and it stacks on top of everything above.
We asked each harness to fan the work out to two parallel subagents. This is where totals detonate.
Claude Code completed the task with 9 model requests across three distinct request classes. There was the main session with its full ~33k baseline, and five subagent calls each carrying their own bootstrap of a 3,554-char agent system prompt plus 24 of the 27 tools.
Cumulative metered input reached 513,000 tokens, against 121,000 for the same work done directly. That is a 4.2x multiplier for one modest fan-out, because every subagent pays its own bootstrap and its transcript is then ingested by the parent.
OpenCode's design here is notably leaner. Its subagent requests carry a reduced profile of a 1,379-char system prompt and 5 tools. Its subagent lane did not complete cleanly through our gateway, so we report the design difference from the capture

[truncated]
