---
source: "https://runware.ai/blog/the-closed-source-llm-premium-has-collapsed"
hn_url: "https://news.ycombinator.com/item?id=48685911"
title: "The closed-source LLM premium has collapsed"
article_title: "The closed-source LLM premium has collapsed | Runware"
author: "theally"
captured_at: "2026-06-26T12:47:23Z"
capture_tool: "hn-digest"
hn_id: 48685911
score: 3
comments: 0
posted_at: "2026-06-26T12:34:32Z"
tags:
  - hacker-news
  - translated
---

# The closed-source LLM premium has collapsed

- HN: [48685911](https://news.ycombinator.com/item?id=48685911)
- Source: [runware.ai](https://runware.ai/blog/the-closed-source-llm-premium-has-collapsed)
- Score: 3
- Comments: 0
- Posted: 2026-06-26T12:34:32Z

## Translation

タイトル: クローズドソースのLLMプレミアムは崩壊した
記事のタイトル: クローズドソースの LLM プレミアムが崩壊 |ランウェア
説明: オープンソース モデルは、87% 低いコストでクローズド モデルのベンチマークに匹敵するようになりました。これが、オープンソース推論と OpenAI のどちらを選択する開発者にとって何を意味するのかを次に示します。

記事本文:
クローズドソースのLLMプレミアムは崩壊した |ランウェアモデル
記事ショーケース 12 分で読むクローズドソースの LLM プレミアムが崩壊
オープンソース モデルは、87% 低いコストでクローズド モデルのベンチマークに匹敵するようになりました。これが、オープンソース推論と OpenAI のどちらを選択する開発者にとって何を意味するのかを次に示します。
LLM への最初の電話は何でしたか?ほぼ間違いなく、次のようなものです。
「openai」から OpenAI をインポートします。
const client = new OpenAI ();
const response = client.responses を待ちます。作成({
モデル: "gpt-5.5" 、
input: 「ユニコーンについての短い就寝前の話を書いてください。」 、
});
コンソール。ログ (response.output_text);
LLM へのエントリーポイントはフロンティアを通過することです。ここは誰もが始める場所ですが、あまりにも多くの人がそこに留まる場所でもあります。彼らは、この呼び出しのバリエーションに基づいてアプリ全体、ワークフロー、ハーネスを構築し、月末に来る API の請求を見てひるみます。
独自のモデルが非常に支配的だった時代には、それは理にかなっていました。製品を優れたものにするには、基礎となるモデルの品質が高くなければなりません。そして、高品質は常に独自性を意味します。そのため、開発者は卓越性を得るためにコストを費やしました。
それは今でも当てはまりますか？あまり。確かに、プロプライエタリは依然としてフロンティアですが、18 か月前にフロンティアだった場所は、現在では明確にマッピングされた領域になり、オープンソース モデルが数分の 1 のコストでそのギャップを埋めつつあります。
なぜ？そして開発者は、この新たに発見されたオープンソースの領域をどう活用すべきでしょうか?
オープンソースをベンチに置き続けた理由
はっきりさせておくべきです。オープン モデルがプロプライエタリ モデルに取って代わるまでには至っていません。まだまだ普及には遅れがございます。オープンソース モデルは、同等のインテリジェンスであれば 87% 安価ですが、それでもトークン シェアの 25 ～ 30% しか保持していません。
しかし、そのギャップはもはや品質の問題ではありません。長年にわたり、購入の決定は二者択一でした。独自の価格を支払うか、より悪いモデルを受け入れるかです。マジで

支払い済みです。この習慣により、クローズド モデルが閉じられることを正当化する品質格差がなくなってから長い間、今日でもクローズド モデルが優勢であり続けています。
GPT-4 を例として考えてみましょう。 2023 年に発売されたとき、GPT-4 の価格は 30 ドル/M トークンでしたが、実際に機能する唯一のモデルでもありました。
(出典: GPT-4 テクニカル レポートおよび Llama2 論文)
Llama 2 は大きく後れを取っていたわけではありませんが、製品レベルのものにとってリスクとなるには十分な遅れをとっていたのです。 HumanEval は厳しいものでした。37 ポイントの差は、Llama 2 がコードに隣接するものにとって現実的な選択肢ではないことを意味しました。
これをすべてのモデルに当てはめると、傾向は明らかです。クローズド モデルの方が「優れている」ということです。
図 9: クローズド ソース モデルとオープンソース モデルの即時価格とインテリジェンスの関係。出典: NBER Working Paper 34608 。
高価でもあります。これは OS モデルの利点ですが、落とし穴がありました。トークンに節約したものはすべて、GPU 配管で返済しました。 fp16 での Llama 2 70B には約 140GB の VRAM が必要だったので、少なくとも 2 つの H100、または独自の品質のトレードオフを伴うより困難な量子化セットアップが必要でした。これに、サービス スタックを接続するためのエンジニアリング リソースが追加されると、コストが均等化し始めます。
しかしおそらく、クローズドモデルが勝った最大の理由は、それが単なるモデルではなかったことでしょう。閉鎖されたラボでは、分銅とともに安定したドラムビートのような製品が出荷されました。関数の呼び出し、構造化された出力、ビジョン、ファイルのアップロード、バッチ処理、微調整はすべて連携され、バージョン管理されます。
以下は、18 か月にわたる OpenAI リリースのタイムラインです。
ビジョン (GPT-4V)、2023 年 11 月
構造化された成果物、2024 年 8 月
オープンソースでは重みが得られます。モデルを製品に変えるものはすべて自分で構築します。
しかし、AI が私たちに教えてくれたことが 1 つあるとすれば、それは、今日は明日ではないということです。 1 年前、エージェントを使ってコーディングをする人は誰もいませんでした。さて、私は誰も

それらを使用せずにコーディングします。 AI は常に変化しており、このフロンティアの進歩の裏で経済は漂い続けています。 Inference は年間約 10 倍安くなりました。能力の遅れは 18 か月から数か月に短縮されました。すると、次の 3 つのことが同時に起こりました。
オープンソースは最先端のベンチマークに到達し始めました。
フロンティアラボは値上げを始めた。
パフォーマンスマージンがなくなった
Moonshot AI は、2025 年 7 月に キミ K2 を出荷し、その後、2026 年 2 月にキミ K2.5、そして 4 月にキミ K2.6 を出荷しました。現在のバージョンは、32B がアクティブな 1 兆パラメータの MoE です。ベンチマークは、6 か月前には Opus のみだった領域に到達しました。
体験の背後にあるモデル
Kimi K2.6 は、最先端のコーディング、長期的な実行、エージェント スウォーム機能を備えた Moonshot AI の最新のオープンソース モデルです。
人類最後の試験 (フル) ツール付き
ターミナルベンチ 2.0 (ターミナル 2)
キミはトップモデルをリードしている、あるいはトップモデルと並んでいる。これは、フロンティアが処理するのと同じタスクのために、ダウンロードして独自のハードウェアで実行し、運用環境にデプロイできるモデルです。
そして追い上げはキミに限ったことではない。 DeepSeek V4 Pro、GLM-5.1、Qwen 3.6、および Mistral Medium 3.5 はすべて、2026 年第 1 四半期にフロンティア層のベンチマークを出荷しました。なぜですか?
オープンラボはスリップストリームに乗っています。フロンティアクローズドモデルは、高価な探索を行います。オープン ラボでは、軌跡を抽出し、合成データから学習し、フロンティアですでに証明されているパターンに対して事後トレーニングを行います。最初に問題を解決したモデルが全額を支払います。 2 番目のモデルの料金はほんのわずかです。
アーキテクチャ プレイブックが公開されました。 MoE ルーティング、ロングコンテキスト トリック、テスト時の推論、エージェントのハーネス。 3 年前、これらは研究室の秘密でした。現在、それらは論文、ブログ投稿、Hugging Face のリファレンス実装になっています。技術が公開されると、それを実装するまでの期間は数四半期ではなく、数週間になります。
コンプ

ute はもはやボトルネックではありません。 2023 年にフロンティアクラスのモデルをトレーニングするには、OpenAI サイズのクラスターが必要でした。 2026 年には、数千台の H100 を備えた資金豊富な研究所が、競争力のあるモデルを 1 四半期で出荷できるようになります。ディープシークがやってくれました。ムーンショットがやってくれました。志布がやったよ。障壁は十分に低くなり、「フロンティアクラス」はもはや一企業の成果ではなくなりました。
これらのモデルを出荷するラボは、初日からホストする準備ができている推論プロバイダーのロングテールにも依存しています。蒸留パイプライン、ベンチマークの検証、開発者の導入はすべて、そのレイヤーを介して実行されます。
まだギャップがあります。クローズド モデルは、最も困難な推論と、長年にわたる RLHF とレッド チームから得られる洗練をリードします。しかし、その遅れはわずかであり、オープンラボ自身がそう言っています。 DeepSeek 自体は、自社の V4 モデルを最先端のフロンティアより 3 ～ 6 か月遅れさせており、現行モデルには及ばないものの、前世代のフラッグシップモデルを上回っています。ほとんどの本番作業では、最も困難な問題に数か月遅れても、決定は変わりません。
独自の価格設定は現在混乱しており、さらに混乱しています。フロンティアラボは価格をどこに設定すればよいのか実際にはわかっておらず、過去6か月間は公の場でそれを検討してきた。
トリガーはエージェントの使用です。チャット通話では数百のトークンが消費されます。エージェントは簡単に数百万ドルを使い果たすことができます。 Claude Code Max ユーザーは、月額 200 ドルのプランから約 5,000 ドルの使用量を抽出していました。補助金付きの定額サブスクリプションであっても、このようなデルタでは生き残ることはできません。
価格設定はどのように変化しているのでしょうか?ユーザーには以下が表示されます:
明示的なハイキング。 GPT-5.5 は、前世代のトークンあたりのコストの約 2 倍でリリースされました。
ステルスハイキング。 Opus 4.7 は同じステッカー価格を維持しましたが、新しいトークナイザーは同じプロンプトに対して最大 35% 多くのトークンを生成します。同じ料金でも請求額は高くなります。
アクセスの変更。コーデックスはトークごとに移行

n 請求。 Anthropic は OpenClaw をクロードのサブスクリプションから切り離しました。 Google は Gemini API に支出上限を追加しました。
上記の「製品」モデルを通じて築かれた好意は、開発者が突然アクセスがロックされたり制限に達したりするたびに、少しずつ元に戻ります。一方、オープンソース価格の下限は下がり続けています。 Runware 上の MiniMax M2.7 は、100 万入出力トークンあたり 0.30 ドルまたは 1.20 ドルで実行されます。比較のために、Opus 4.7 は $5/$25 、GPT-5.5 は $5/$30 です。エージェントはアウトプットを重視しており、ギャップが最も大きいのはアウトプットです。重要なコーディング タスクは、数十万の出力トークンを簡単に実行できます。 GPT-5.5 ($30/M out) では、各タスクに数ドルのコストがかかります。 M2.7 ($1.20/M アウト) では、ペニーです。
トークンごとの課金はエージェント ワークロードに適したモデルであり、ラボはそれを認識しています。正しく実行すると、実際に実行した推論の秒数に対して料金を支払うことになり、最低額、コミットメント、切り上げはありません。フロンティアラボはまだそこにはいません。
しかし、彼らの顧客ベースは「定額料金、無制限」に固定されており、それをきれいに元に戻すことはできません。そのため、彼らは価格を上げずに価格を上げ、アクセスを制限せずにアクセスを制限し、誰も累積効果を計算しないことを望んでいます。
モデルよりもハーネスの重要性が高まっている
単一モデルへの固定観念は解消されつつあります。エージェントは、当面のタスクについてモデルやモダリティ全体をルーティングする必要があります。
これは Runware の基本的な前提です。当社はモダリティにわたる 400,000 以上のモデルを単一の宛先としており、モデルの選択を簡単な構成決定で行うことができます。これには 2 つの理由があります。
まず、フロンティア エージェントは複数のモデルを使用してタスクを実行できます。
Claude Code はメイン モデルを使用して重労働を実行し、小規模なモデルがセッションの 1 行の概要などの安価なバックグラウンド ジョブを処理し、別のモデルがサイド Q に応答します。

質問機能を使用すると、タスクの途中で簡単な質問をしてもメイン モデルが中断されなくなります。
ChatGPT へのメッセージは 1 つのモデルにヒットしません。小さなルーターが最初にリクエストを読み取り、それがどこに送られるかを決定します。簡単な事実に関する質問は高速モデルに、厳密な推論またはコーディング タスクはより深いモデルに送られます。ルーター自体は何も推論したり生成したりしません。発送します。簡単な作業は安価なモデルに、難しい作業は高価なモデルに移されるため、些細なリクエストに対してフロンティア料金を支払う必要がなくなります。
ディスパッチャーは安価でありながら、作業を実行するために高価なモデルにルーティングすることができます。モデルの選択が単なる構成になったら、同じものをオープンソースで構築できますが、クローズド製品にはない追加機能が 1 つあります。サイズだけでなく、そのモデルの得意分野によってモデルを選択できます。
推論と計画。リクエストをステップに分割し、いつ完了するかを決定するには、利用可能な最も強力な推論が必要です。それは現在、DeepSeek V4 や Kim K2.6 のようなオープンソース モデルになる可能性があり、どちらも閉鎖されたフロンティアから数か月以内に実現できます。
コード生成。明確な仕様を作業ファイルに変換することは、コーディング調整されたモデルに適しています。 Qwen ファミリーの何かが、推論モデルのコストの数分の 1 でそれを処理します。
会話層。エンドユーザーが何を話す場合でも、Gemma などのより会話型のモデルの恩恵を受けます。
これはコストの点ではるかに有利です。会話モデルに到達するチャット ターンにはフロンティア料金は決して支払われません。また、コーディング モデルに到達するコード生成ステップも同様です。各モデルが独自のジョブで機能する限り、リクエストごとの実質的なコストは 1 つの小規模から中規模のモデルであり、すべてのジョブをそれ自体で実行する 1 つの大規模モデルよりもはるかに低くなります。
ここから 2 番目の理由がわかります。ワークフローにはマルチモーダリティが必要です。最も有用なエージェントはテキスト以上のものを扱います。書面による提案を行うエージェントを構築しているとします。

概要を 60 秒の起動ビデオにまとめます。パイプラインは次のようにヒットします。
ナレーション用の TTS モデル。
背景ビジュアル用の画像モデル。
再びキャプション用のテキスト モデル。
統合 API がないと、エージェントは 5 つの個別のサービスに関連付けられ、それぞれが独自の SDK、認証フロー、レート制限、請求パイプライン、エラー セマンティクスを持ちます。これは、チームが一度作成すれば永久に維持できる統合レイヤーです。モダリティ全体で統一された API により、その層が削除されます。
これの長期的な形は、エージェントが LLM 市場経済におけるタスク リストにどのモデルを使用するかを決定することです。開発者は初期仕様を作成します。その後、エージェントはプロバイダー間のスプレッドをリアルタイムで購入します。モデルの能力が商品になります。
オープンソースはより優れたインフラストラクチャを構築します
モデルの問題が解決したら、次の問題はレイテンシと配置です。
エージェントは、チャットでは決して起こらなかった方法で遅延を増大させます。サポート チャットを開いているユーザーは、1 回 500 ミリ秒に気づきます。エージェントが計画、ルーティング、実行のために 50 回の連続呼び出しを行うと、その 500 ミリ秒が重なり、ユーザーが実際に待機するのは 25 秒の遅延となります。実際のエージェント ループにおけるコールド スタート、取得、およびツール呼び出しのホップが乗算されると、往復遅延がユーザー エクスペリエンスを支配し始めます。
垂直

[切り捨てられた]

## Original Extract

Open-source models now match closed-model benchmarks at 87% lower cost. Here's what that means for developers choosing between open source inference and OpenAI.

The closed-source LLM premium has collapsed | Runware Models
Articles Showcase 12 min read The closed-source LLM premium has collapsed
Open-source models now match closed-model benchmarks at 87% lower cost. Here's what that means for developers choosing between open source inference and OpenAI.
What was your first call to an LLM? Almost definitely, something like this:
import OpenAI from "openai" ;
const client = new OpenAI ();
const response = await client.responses. create ({
model: "gpt-5.5" ,
input: "Write a short bedtime story about a unicorn." ,
});
console. log (response.output_text);
The entry point to LLMs is through the frontier. This is where everyone starts, but it’s also where too many stay. They build entire apps, workflows, and harnesses around variations on this call, then wince at their API bill come end-of-month.
That used to make sense when the proprietary models were so dominant. If you want your product to excel, the quality of the underlying model must be high. And high quality has always meant proprietary. So developers ate the cost for excellence.
Does that still hold? Not really. Yes, proprietary is still the frontier, but what was the frontier 18 months ago is now well-mapped territory, and open-source models are closing the gap at a fraction of the cost.
Why? And what should developers make of, and with, these newfound lands of open source?
What kept open source on the bench
We should be clear. Open models aren’t close to taking over from proprietary. There is still a lag in uptake. Open source models are 87% cheaper at equal intelligence but still hold only 25-30% of the token share.
But that gap is no longer about quality. For years, the buying decision was binary: pay proprietary prices, or accept a worse model. Serious teams paid up. That habit is what keeps closed models dominant today, long after the quality gap that justified it closed.
Let’s take GPT-4 as our example. When it launched in 2023, GPT-4 was at $30/M tokens , but it was also the only model that could actually do the work.
(Sources: GPT-4 Technical Report and the Llama2 paper )
Llama 2 wasn’t wildly behind the curve, but it was far enough behind to be a risk for anything production-grade. HumanEval was the rough one: a 37-point gap meant Llama 2 wasn't a real option for anything code-adjacent.
When you extrapolate this to all models, the trend is clear: Closed models are “better”.
Figure 9: Prompt price vs. intelligence for closed- and open-source models. Source: NBER Working Paper 34608 .
They are also more expensive. This was the in for OS models, but it came with a catch. Whatever you saved on tokens, you paid back in GPU plumbing. Llama 2 70B at fp16 needed roughly 140GB of VRAM , so two H100s at a minimum, or a more painful quantized setup with its own quality tradeoffs. Add to that the engineering resources to wire together your serving stack, and the costs start to equalize.
But perhaps the biggest reason closed models won was that they weren’t just models. Closed labs shipped a steady drumbeat of product alongside the weights. Function calling, structured output, vision, file uploads, batch processing, fine-tuning, all wired together and versioned.
Here's a timeline of OpenAI releases over 18 months:
Vision (GPT-4V), November 2023
Structured outputs, August 2024
Open source gives you weights. Everything that turns a model into a product, you build yourself.
But if there is one thing AI has taught us, it is that today isn’t tomorrow. A year ago, no one was coding with agents significantly. Now, no one is coding without them. AI is in constant flux, and underneath all this frontier progress, the economics keep drifting. Inference got roughly 10x cheaper per year. The capability lag compressed from 18 months to a few. Then three things hit at once:
Open source started landing on frontier benchmarks.
Frontier labs started raising prices.
The performance margin has disappeared
Moonshot AI shipped Kimi K2 in July 2025, then Kimi K2.5 in February 2026, then Kimi K2.6 in April. The current version is a 1-trillion-parameter MoE with 32B active. The benchmarks land in territory that was Opus-only six months ago.
The model behind the experience
Kimi K2.6 is Moonshot AI's latest open-source model, with state-of-the-art coding, long-horizon execution, and agent swarm capabilities.
Humanity's Last Exam (Full) w/ tools
Terminal-Bench 2.0 (Terminus-2)
Kimi is leading or keeping pace with top models. It's a model you can download, run on your own hardware, and deploy to production for the same tasks the frontier handles.
And the catch-up isn't unique to Kimi. DeepSeek V4 Pro, GLM-5.1, Qwen 3.6, and Mistral Medium 3.5 all shipped frontier-tier benchmarks in Q1 2026. Why?
Open labs are riding the slipstream. Frontier closed models do the expensive exploration. Open labs distill trajectories, learn from synthetic data, and post-train against patterns the frontier has already proven out. The first model to solve a problem pays the full cost. The second model pays a fraction.
The architecture playbook is now public. MoE routing, long-context tricks, test-time reasoning, agent harnesses. Three years ago, these were lab secrets. Now they're papers, blog posts, and reference implementations on Hugging Face. Once a technique is in the open, the gap to implement it is weeks, not quarters.
Compute is no longer the bottleneck it was. Training a frontier-class model in 2023 took an OpenAI-sized cluster. In 2026, a well-funded lab with a few thousand H100s can ship a competitive model in a single quarter. DeepSeek did it. Moonshot did it. Zhipu did it. The barrier dropped enough that "frontier-class" is no longer a one-company achievement.
The labs that ship these models also depend on a long tail of inference providers ready to host them on day one. Distillation pipelines, benchmark validation, and developer adoption all run through that layer.
There are still gaps. Closed models lead on the hardest reasoning and on the polish that comes from years of RLHF and red-teaming. But the lag is small, and the open labs say so themselves. DeepSeek themselves put their own V4 models 3 to 6 months behind the state-of-the-art frontier, beating last generation's flagships while trailing the current ones. For most production work, a few months of lag on the hardest problems doesn't change the decision.
Proprietary pricing is messy right now, and getting messier. The frontier labs don't really know where to set prices, and the last six months have been them figuring it out in public.
The trigger is agentic usage. A chat call burns a few hundred tokens. An agent can easily burn through millions. Claude Code Max users were extracting around $5,000 in usage from $200 monthly plans. Even subsidized, flat-rate subscriptions don't survive a delta like that.
How is pricing shaking out? Users are seeing:
Explicit hikes . GPT-5.5 launched at roughly 2x the per-token cost of its predecessor.
Stealth hikes . Opus 4.7 kept the same sticker price, but its new tokenizer generates up to 35% more tokens for the same prompts. Same rate, higher bill.
Access changes . Codex shifted to per-token billing . Anthropic cut off OpenClaw from Claude subscriptions. Google added spend caps on the Gemini API.
The goodwill built through the “product” model above is undone bit by bit every time a developer suddenly finds their access locked or their limits hit. Meanwhile, the floor under open-source pricing continues to drop. MiniMax M2.7 on Runware runs at $0.30/$1.20 per million input/output tokens. For comparison, Opus 4.7 is $5/$25 , and GPT-5.5 is $5/$30 . Agents are output-heavy, and output is where the gap is widest. A non-trivial coding task can easily run through hundreds of thousands of output tokens. At GPT-5.5 ($30/M out), each task costs several dollars. At M2.7 ($1.20/M out), it's pennies.
Per-token billing is the right model for agentic workloads, and the labs know it. Done right, it means paying for the seconds of inference you actually run, no minimums, no commitments, no rounding up. The frontier labs aren't there yet.
But their customer base anchored on "flat rate, unlimited," and they can't walk that back cleanly. So they're raising prices without raising prices, restricting access without restricting access, and hoping no one tallies the cumulative effect.
Harnesses are becoming more important than models
Lock-in to a single model is dissipating. Agents must route across models and modalities for any task at hand.
This is a fundamental premise of Runware. We are a single destination for 400k+ models across modalities, making model choice a simple config decision. This needs to happen for two reasons.
First, frontier agents can use multiple models to perform tasks.
Claude Code uses a main model to do the heavy lifting, a small model handles cheap background jobs like the one-line summaries of your sessions, and a separate model answers the side-questions feature so a quick question mid-task doesn't interrupt the main model.
A message to ChatGPT doesn't hit one model. A small router reads the request first and decides where it goes: a quick factual question to the fast model, a hard reasoning or coding task to the deeper one. The router doesn't reason or generate anything itself. It dispatches. Easy work goes to the cheap model, hard work to the expensive one, so you stop paying frontier rates for trivial requests.
The dispatcher can be cheap and route to expensive models to do the work. Once model choice is just a configuration, you can build the same thing on open source, with one addition the closed products don't offer: you can pick models by what they're good at, not just by size.
Reasoning and planning . Breaking a request into steps and deciding when it's done wants the strongest reasoning available. That can now be an open-source model like DeepSeek V4 or Kimi K2.6, both within a few months of the closed frontier.
Code generation . Turning a clear spec into a working file suits a coding-tuned model. Something in the Qwen family handles it at a fraction of the reasoning model's cost.
The conversational layer . Whatever the end user talks to benefits from a more conversational model, such as Gemma.
This works much better for cost. A chat turn that lands on a conversational model never pays frontier rates, and neither does a code generation step that lands on a coding model. As long as each model is capable at its own job, your effective per-request cost is one small-to-medium model, well below a single large model doing every job itself.
Which brings us to the second reason: workflows necessitate multimodality . Most useful agents touch more than text. Say you're building an agent that turns a written product brief into a 60-second launch video. The pipeline hits:
A TTS model for the voiceover.
An image model for background visuals.
A text model again for captions.
Without a unified API, the agent is wired into five separate services, each with its own SDK, auth flow, rate limit, billing pipeline, and error semantics. That's a layer of integration your team writes once and maintains forever. A unified API across modalities removes that layer.
The long-term shape of this is agents deciding which models to use for their task lists in an LLM market economy . The developer writes the initial spec. After that, the agent shops the spread across providers in real time. Model capability becomes a commodity.
Open source will build better infrastructure
Once the model question is settled, the next question is latency and placement.
Agents amplify latency in a way that chat never did. A user opening a support chat notices 500ms once. With an agent making 50 sequential calls to plan, route, and execute, that 500ms compounds into a 25-second delay the user actually waits through. Multiply by the cold-start, retrieval, and tool-call hops in a real agent loop, and round-trip latency starts to dominate the user experience.
Vertical

[truncated]
