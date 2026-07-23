---
source: "https://www.highflame.com/blog/the-three-moments-your-ai-gateway-can-ruin/"
hn_url: "https://news.ycombinator.com/item?id=49017178"
title: "How good is your AI Gateway?"
article_title: "The Three Moments Your AI Gateway Can Ruin · Highflame"
author: "sharathr"
captured_at: "2026-07-23T05:44:45Z"
capture_tool: "hn-digest"
hn_id: 49017178
score: 1
comments: 2
posted_at: "2026-07-23T05:07:16Z"
tags:
  - hacker-news
  - translated
---

# How good is your AI Gateway?

- HN: [49017178](https://news.ycombinator.com/item?id=49017178)
- Source: [www.highflame.com](https://www.highflame.com/blog/the-three-moments-your-ai-gateway-can-ruin/)
- Score: 1
- Comments: 2
- Posted: 2026-07-23T05:07:16Z

## Translation

タイトル: あなたの AI ゲートウェイはどの程度優れていますか?
記事のタイトル: AI ゲートウェイが台無しになる 3 つの瞬間 · Highflame
説明: 3 つの AI ゲートウェイの背後に、実際のタイミングを備えたモデルを配置しました。 Highflame は最初のトークンに約 2 ミリ秒を追加し、733 MB で保持された 5,000 件の会話の 100% に応答し、200 ミリ秒の MCP ツール呼び出しに 17 ミリ秒を追加します。

記事本文:
Highflame Identity はオープン ソースになりました。オープン スタンダード上のエージェント ID。リリースを読む → プラットフォーム ソリューション ▾ エンジニアリング エージェント向け コントロールが組み込まれているため、より迅速に出荷できます。 セキュリティについて すべてのアクションを承認します。爆発範囲が含まれます。 IT およびプラットフォームの場合 他のアイデンティティと同様にガバナンス エージェントも対象となります。コンプライアンスの帰属と監査対応の証拠用。ドキュメントのリサーチ リソース デモを予約する プラットフォーム ソリューション エンジニアリング向け セキュリティ向け IT およびプラットフォーム コンプライアンス向け ドキュメントのリサーチ リソース デモを予約する ← すべての記事 AI ゲートウェイが台無しになる 3 つの瞬間
高速モデルを間違ったゲートウェイの背後に配置すると、ユーザーは空のチャット バブルを 1 秒半見つめることになります。モデルはすでに最初のトークンを送信しています。ゲートウェイがそれを保持しています。私たちはまさにその出来事を測定しました。
実稼働環境でモデルを実行すると、最終的にはゲートウェイが作成されます。優れたものは、単一の URL を通じてすべてのプロバイダーをルーティングし、API キーを開発者のラップトップに近づけず、チームごとの支出を測定し、トラフィックが離れる前にスキャンします。問題は、ユーザーが行うすべてのリクエストの中にそれが存在することです。したがって、どれかを選ぶ前に、それが実際に感じる瞬間にどのような影響を与えるのかを知りたいと思うでしょう。
パート 1 では、インスタント バックエンドに対する生のゲートウェイ コストを測定しました。誰もインスタント バックエンドを出荷していないため、今回は実際のモデルのように動作します。最初のトークンは約 300 ミリ秒、ターンは 2 秒で、トークンは一度に 1 つずつストリーミングされます。その前に、Highflame (Rust)、Bifrost (Go)、および LiteLLM (Python) があります。
エクスペリエンスを決定する 3 つの瞬間:
最初のトークン。アプリはまだインスタントだと感じますか?
ピーク時間帯。全員が同時に話すと電話はかかってきますか?
ツール呼び出し。エージェントはアクションごとにいくら支払いますか?
最初のトークンの前の沈黙の部分で、ユーザーはアプリがハングしたかどうかを判断します。モデルはその静寂のほぼすべてを所有しており、ゲートウェイはどこにもありません。

目に見えて、最初のトークンが生成されるまでに約 300 ミリ秒かかります。ゲートウェイにとって唯一の問題は、ゲートウェイに何を追加するかです。
Highflame では約 2 ms が追加されます。これは、ノイズが最も小さい 10 個のチャットで測定された料金全体であり、最初に 1 回支払われます。完了した回答は、実行したすべての同時実行で直接ストリームから 1 桁ミリ秒以内に到達するため、トークンごとにコストが蓄積されることはありません。 100 の同時チャットでは、Highflame を介した最初のトークンは 307 ミリ秒で到達します。これは、ゲートウェイがまったくないモデルに到達するのと区別できません。
仕様書には決して記載されないため、これらの失敗には名前を付ける価値があります。
ビフロストバッファ。標準設定では、ストリームを転送しますが、モデルが終了するまですべてのバイトを保持するため、最初のトークンが最後のトークンと一緒に表示されます。これは、毎ターン、ゲートウェイによって追加された待機時間が 1.3 秒あり、トークンがずっと流れていたにもかかわらず、UI がフリーズしたように表示されます。
LiteLLM 化合物。 10 チャットでは適切にストリーミングされますが、トークンごとの作業は負荷がかかると積み重なります。 50 の同時チャットでは、最初のトークンに約 350 ミリ秒が追加されます。 100 では、最初のトークンの中央値は 6 秒を超え、直接 1.6 秒かかるストリーミングされた回答には 7.7 秒かかります。
瞬間 2: 全員が同時に話す
実際の会話はつながりを維持します。各リクエストは順番が終了するまで約 2 秒待機するため、忙しい午後 (エージェント フリートが展開し、CI が実行され、会社全体が同じ 1 時間内に) は、一度に数千のソケットが開かれたままになることになります。
5,000 件の会話を 5 分間同時にオープンし、各会話は 2 秒間のモデルターンで待機しました。 Highflame は、提供された 572,563 件のリクエストすべてに真の 100% で応答し、p99 は 3.3 秒で、ラッシュ全体を 733 MB のメモリで実行しました。ビフロストの答え

すべてを実行しましたが、実行するには 2 倍のメモリが必要でした。ラッシュを保持することは、コンパイルされたゲートウェイが行うべきことです。マシンの半分に保持するかどうかが違います。
い草が畑の底を隔てています。 LiteLLM はリクエストの 30% に応答し、生存者を中央値 6.2 秒待たせ、p99 は 26.7 秒でした。失われた 10 件中 7 件の呼び出しはまったく戻ってきませんでした。運用環境では、ユーザーが既に溺れかけているシステムにプロンプ​​トを再送信することになります。
モデルの 2 秒間がゲートウェイの非表示を停止した瞬間に上部が分離されます。パート 1 では、同じ 5,000 の接続を即時回答で急いで実行しました。Highflame は 0.83 秒の p99 で、Bifrost の 0.99 秒で回答しました。どちらも 100% でした。 Highflame は 5 分間フルスロットルで動作し、1 秒あたり 14,331 件のリクエストを処理し、430 万件のリクエストにわたって失敗はゼロでした。これは Bifrost の速度の 2 倍強です。
購入者にとっての罠は、会話が 100 回行われた時点で、3 つのゲートウェイすべてが依然としてすべてに応答し、2 秒の応答で 100% 成功することです。 500 で、LiteLLM の中央値は、モデル自身の回転をすでに 2 秒上回っていますが、他の 2 つは横ばいです。違いは、成長しようとしている規模に応じて存在します。これは、まさに、選択する前に手動でテストすることのない規模です。
エージェントは通話ツールに時間を費やしますが、MCP はそれらの通話を伝達する手段です。ここでは、レースが始まる前にフィールドが薄くなります。Highflame は、MCP クライアントと MCP サーバーの間に透過的に配置でき、クライアント自身のセッションが直接通過できる 3 つのうちの唯一の 1 つです。 Bifrost の MCP サポートは、ツールをスタンドアロンの MCP クライアントではなく LLM 呼び出しに向けており、LiteLLM の MCP ゲートウェイでは同等の透過的な測定パススル​​ーが提供されません。
つまり、バーはツールそのものであり、重要なのは数値だけです。

プロキシがそれに上書きします。 Highflame は、10 個の同時セッションで 17 ミリ秒を追加します。 100 セッションでは 181 ミリ秒が追加され、そのほとんどのセッションはゲートウェイ作業ではなく同じツールのキューに入れられ、1 秒あたり 260 件のコールを失敗なしで配信します。
100 セッションという数字も天井には程遠いです。インスタント ツールに対して専用のロード リグでプッシュされると、4 コア上の単一の Highflame インスタンスは 1 秒あたり 8,200 のツール呼び出しでピークに達し、1,000 の同時エージェント セッションで約 6,800 を保持します。そして、キャパシティ プランナーが考慮すべき部分は、同じツール サーバーがゲートウェイをまったく使用せずに直接攻撃し、その 1,000 セッションのファンインで座屈したということです。
このヘッドルームが重要なのは、MCP プロキシが Highflame のガバナンス作業を行う場所であるためです。MCP プロキシは呼び出し元を認証し、Cedar のポリシー チェックと Shield のスキャナをオンにしたときに実行される適用ポイントでもあります。
Highflame は 25 MB で起動し、66 MB を携えて 5,000 会話のラッシュに入りました。ラッシュ自体は 733 MB でピークに達し、その後は落ち着きました。 LiteLLM は、すでに Highflame のワーキング セットの 7 倍である 486 MB を保持している同じラッシュに参加し、その 30% に応答しました。
このようなフットプリントにより、ゲートウェイを、すべてが到達する必要がある 1 つの中央ボックスとしてではなく、ワークロードの隣のサイドカーとして、またはエージェント サンドボックスごとに 1 つとして実行できるようになります。
まだモデルの前にゲートウェイを配置していない場合は、ユーザーが感じた順序で 3 つの質問が重要です。ハイフレイムの答えは次のとおりです。
最初のトークンはまだインスタントに感じられますか?はい。モデルには最大 300 ミリ秒かかります。 Highflame では約 2 ミリ秒が追加され、100 の同時チャットではこの 2 つは区別できません。
ピーク時に通話が切断されることはありますか?いいえ。733 MB のメモリ内で、5 分間に行われた 5,000 件の会話で、572,563 件のリクエストすべてに応答しました。
ツールの呼び出しには料金がかかりますか?ハイフラ

私は 200 ミリ秒のツール呼び出しに 17 ミリ秒を追加しましたが、どれも失敗しませんでした。
Highflame がそのホット パス上で実際に検査する内容は、プラットフォーム ページにあります。
リグは 2 つの AWS ホストで構成されます。ゲートウェイとシミュレートされたモデルは 4 コアの c7i.xlarge を共有し、ロード ジェネレーターは独自の 8 コア ボックスで実行され、すべてのリクエストは実際のネットワークを通過します。モデルの時間は実際のモデルと同様です (最初のトークンまで約 300 ミリ秒、ターンは 2 秒、トークンはストリーミングされます)。すべてのゲートウェイは、新規ブートからガードレールをオフにしたデフォルト設定でストック ビルドを実行しました（2026 年 7 月 22 日に測定）。 5 分の数字ごとに、途切れることのない 5 分が表示されます。 MCP の天井は、別の 3 ホスト リグからのものです。パート 1 の完全なデータセットはベンチマーク ページにあります。
Highflame では、エージェント ID、ランタイム ポリシー、自律型 AI の保護に関するエンジニアリングと研究を行っています。
エンジニアリング 2026 年 7 月 1 日 AI ゲートウェイは、全員が同時にアクセスするまでは問題ありません
エンジニアリング 2026 年 6 月 22 日 Anthropic には独自のコーディング エージェントがどのように含まれているか、およびその範囲をフリート全体に適用する方法
エンジニアリング 2026 年 6 月 3 日 3 つのゲートウェイ、1 つの意思決定ファブリック
今すぐエージェントの保護を始めましょう。
Highflame は、すべての AI エージェントに信頼できる ID を与え、すべてのアクションに承認の決定を与えます。社内で構築されたもの、ベンダーから購入されたもの、マーケットプレイスからダウンロードされたものなど、環境全体でエージェントを検出、管理、制御します。

## Original Extract

We put a model with real timing behind three AI gateways. Highflame adds about 2 ms to the first token, answers 100% of 5,000 held conversations in 733 MB, and adds 17 ms to a 200 ms MCP tool call.

Highflame Identity is now open source: agent identity on open standards. Read the launch → Platform Solutions ▾ For Engineering Ship agents faster, with controls built in. For Security Authorize every action; contain blast radius. For IT & Platform Govern agents like every other identity. For Compliance Attribution and audit-ready evidence. Docs Research Resources Book a Demo Platform Solutions For Engineering For Security For IT & Platform For Compliance Docs Research Resources Book a Demo ← All articles The Three Moments Your AI Gateway Can Ruin
Put a fast model behind the wrong gateway and your user stares at an empty chat bubble for a second and a half. The model already sent its first token; the gateway is holding it. We measured exactly that happening.
If you run models in production, you end up with a gateway eventually. A good one routes every provider through a single URL, keeps API keys off developer laptops, meters spend per team, and scans traffic before it leaves. The catch is that it sits inside every request your users make. So before you pick one, you want to know what it does to the moments they actually feel.
In Part 1 we measured raw gateway cost against an instant backend. Nobody ships an instant backend, so this time it behaves like a real model: first token around 300 ms, a two-second turn, tokens streamed one at a time. In front of it: Highflame (Rust), Bifrost (Go), and LiteLLM (Python).
Three moments decide the experience:
The first token. Does the app still feel instant?
Peak hour. When everyone talks at once, do calls come back?
The tool call. What does an agent pay per action?
The silence before the first token is where a user decides whether the app hung. The model owns almost all of that silence: with no gateway anywhere in sight, it takes about 300 ms to produce its first token. The only question for a gateway is what it adds on top of that.
Highflame adds about 2 ms . That is the whole toll, measured at ten chats where noise is smallest, and it is paid once at the front: the completed answer lands within single-digit milliseconds of the direct stream at every concurrency we ran, so no cost accumulates per token. At a hundred simultaneous chats the first token through Highflame lands in 307 ms , indistinguishable from hitting the model with no gateway at all.
The failures are worth naming, because you would never catch them on a spec sheet:
Bifrost buffers. In its stock configuration it forwards the stream but holds every byte until the model finishes, so the first token shows up together with the last one. That is 1.3 seconds of gateway-added wait on every turn, and the UI reads as frozen even though tokens were flowing the whole time.
LiteLLM compounds. It streams properly at ten chats, but its per-token work stacks up under load. At fifty concurrent chats it adds about 350 ms to the first token; at a hundred, the median first token slips past six seconds , and the streamed answer that takes 1.6 seconds direct takes 7.7 through it.
Moment two: everyone talks at once
Real conversations hold connections open. Each request waits about two seconds for its turn to finish, so a busy afternoon (agent fleets fanning out, a CI run, the whole company inside the same hour) turns into thousands of sockets held open at once.
We held 5,000 conversations open simultaneously for five minutes, every one waiting on a two-second model turn. Highflame answered all 572,563 requests it was offered, a true 100% , with a p99 of 3.3 seconds , and it carried the whole rush in 733 MB of memory. Bifrost answered everything too, but needed twice the memory to do it. Holding the rush is what a compiled gateway should do; holding it on half the machine is the difference.
The rush separates the bottom of the field. LiteLLM answered 30% of its requests and made the survivors wait a median of 6.2 seconds, with a p99 of 26.7 seconds. The missing seven in ten calls simply never came back, which in production is a user re-sending the prompt into a system that is already drowning.
The top separates the moment the model’s two seconds stop hiding the gateways. In Part 1 we rushed the same 5,000 connections at instant answers: Highflame answered at a p99 of 0.83 seconds to Bifrost’s 0.99, both at 100% . Held at full throttle for five minutes, Highflame moved 14,331 requests a second with zero failures across 4.3 million requests, a little over twice Bifrost’s rate.
The trap for a buyer is that at 100 held conversations all three gateways still answer everything: two-second responses, 100% success. At 500, LiteLLM’s median is already two full seconds over the model’s own turn while the other two sit flat. The differences live at the scale you are planning to grow into, which is exactly the scale you will not test by hand before picking one.
Agents spend their time calling tools, and MCP is how those calls travel. Here the field thins out before the race starts: Highflame is the only one of the three we could place transparently between an MCP client and an MCP server, with the client’s own session passing straight through. Bifrost’s MCP support fronts tools to LLM calls rather than to a standalone MCP client, and LiteLLM’s MCP gateway would not give us an equivalent transparent pass-through to measure.
So the bar is the tool itself, and the only number that matters is what the proxy adds over it. Highflame adds 17 ms at ten concurrent sessions. At a hundred sessions it adds 181 ms , most of that sessions queueing for the same tool rather than gateway work, while delivering 260 calls a second with zero failures .
A hundred sessions is also nowhere near the ceiling. Pushed with a dedicated load rig against an instant tool, a single Highflame instance on four cores peaks at 8,200 tool calls a second and holds around 6,800 with a thousand concurrent agent sessions on it. And the part a capacity planner should care about: the same tool server hit directly, with no gateway at all, buckled under that thousand-session fan-in.
That headroom matters because the MCP proxy is where Highflame does its governance work: it authenticates the caller, and it is the enforcement point where Cedar policy checks and Shield’s scanners run when you turn them on.
Highflame boots at 25 MB and walked into the 5,000-conversation rush carrying 66 MB ; the rush itself peaked at 733 MB and settled back down after. LiteLLM entered the same rush already holding 486 MB , seven times Highflame’s working set, and answered 30% of it.
A footprint like that is what lets you run the gateway as a sidecar next to the workload, or one per agent sandbox, instead of as one central box everything must reach.
If you have not put a gateway in front of your models yet, three questions matter, in the order your users feel them. Highflame’s answers:
Does the first token still feel instant? Yes. The model takes its ~300 ms; Highflame adds about 2 ms on top, and at a hundred concurrent chats the two are indistinguishable.
Does peak hour drop calls? No. All 572,563 requests answered across five minutes of 5,000 held conversations, in 733 MB of memory.
Do tool calls pay a toll? Highflame adds 17 ms to a 200 ms tool call, and none fail.
What Highflame actually inspects on that hot path is on the platform page .
The rig is two AWS hosts: the gateway and the simulated model share a 4-core c7i.xlarge, the load generator runs on its own 8-core box, and every request crosses a real network. The model times like a real one (~300 ms to first token, two-second turns, streamed tokens). Every gateway ran its stock build with its default config, guardrails off, from a fresh boot, measured on 2026-07-22; every five-minute number is five unbroken minutes. The MCP ceiling came from a separate three-host rig. Part 1’s full dataset lives on the benchmarks page .
Engineering and research on agent identity, runtime policy, and securing autonomous AI at Highflame.
Engineering Jul 01, 2026 Your AI Gateway Is Fine, Until Everyone Hits It at Once
Engineering Jun 22, 2026 How Anthropic Contains Its Own Coding Agents, and How to Get That Coverage Across Your Fleet
Engineering Jun 03, 2026 Three Gateways, One Decision Fabric
Start securing your agents today.
Highflame gives every AI agent a trusted identity and every action an authorization decision. Discover, govern, and control agents across your environment, whether built in-house, bought from vendors, or downloaded from marketplaces.
