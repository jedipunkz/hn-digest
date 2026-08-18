---
source: "https://getunblocked.com/blog/moving-agent-loops-from-anthropic-to-glm/"
hn_url: "https://news.ycombinator.com/item?id=49345796"
title: "What We Learned Moving Our Agent Loops from Anthropic to GLM"
article_title: "What We Learned Moving Our Agent Loops to GLM"
image: "https://cdn.sanity.io/images/31mw1ch6/production/c1dfef663fc9adef305d6150e5ab04d6bb4aa732-1600x836.png"
author: "dennispi"
captured_at: "2026-08-18T14:23:59Z"
capture_tool: "hn-digest"
hn_id: 49345796
score: 5
comments: 0
posted_at: "2026-08-18T14:02:55Z"
tags:
  - hacker-news
  - translated
---

# What We Learned Moving Our Agent Loops from Anthropic to GLM

- HN: [49345796](https://news.ycombinator.com/item?id=49345796)
- Source: [getunblocked.com](https://getunblocked.com/blog/moving-agent-loops-from-anthropic-to-glm/)
- Score: 5
- Comments: 0
- Posted: 2026-08-18T14:02:55Z

## Translation

タイトル: エージェント ループを人為的から GLM に移行して学んだこと
記事のタイトル: エージェント ループを GLM に移行して学んだこと
説明: トークンごとの価格設定により 95% の節約が約束されます。生産は68%を達成しました。エージェント ループを Claude Opus から GLM 5.2 に移行した方法と、その途中で何が壊れたか。

記事本文:
エージェント ループを GLM 製品に移行して学んだこと 成果 リソース お客様 セキュリティ ドキュメント 価格 もっと見る ログイン 始める デモを予約する 製品 コンテキスト エンジン コーディング エージェントと MCP AI コード レビュー 開発者 Q&A の結果 マージ可能なコード 保存トークン 手戻り削減 リソース ブログ オープンソース ビデオ AI 導入評価 お客様 セキュリティ ドキュメント 価格 ログイン 始める デモを予約する すべての記事 エージェント ループの移行で学んだこと人間から GLM へ
Unblocked のエージェント トラフィックのほとんどを Claude Opus から GLM 5.2 に移行した理由、ブラインド A/B と台帳が実際に示したもの、そして途中で何が壊れたのか。
TL;DR: Unblocked のエージェント トラフィックの大部分を Claude Opus から GLM 5.2 に移動しました。
「トークンごと」の計算では 95% の節約が約束されました。タスクごとの生産現実は 68% 達成されました。そこに到達するには、実際のコード レビューでのブラインド A/B 評価、回線が切断されたマルチプロバイダーのサービング プール、そして数週間にわたる「OpenAI 互換」の驚きが必要でした。
背景: Unblocked により、エージェントは優秀なエンジニアを組織的に理解できるようになります。そのコンテキスト エンジンは、コード、会話、問題、ドキュメント、運用システム全体にわたる知識を結び付けて調整し、作業が行われる場所に関係なくその理解を利用できるようにします。エージェントは MCP/CLI/API を通じてコン​​テキストにアクセスし、Code Review や Unblocked Code などの Unblocked 製品は、そのコンテキストをコードのレビューと作成に直接適用します。その結果、より少ないトークン、より少ない反復、より少ないやり取りで生成される、より優れたソフトウェアが得られます。
過去 2 年間のほとんどにおいて、Unblocked は Anthropic モデルで実行されました。当社の製品は、エンジニアリングの質問に答え、コードをレビューし、組織の組織的知識を検索する一連のエージェント ループです。クロード・オーパスはそのほぼすべての外側のループに座っていました。
この夏、私たちはほとんどの場所を引っ越しました

サードパーティの推論プロバイダーによって提供されるオープンウェイト モデルである GLM へのトラフィック。この投稿では、なぜそれを行ったのか、何を測定したのか、何が問題になったのか、そしてどこに到達したのかについて説明します。システムの内部構造については説明しませんが、調査結果はエージェント ループを大規模に実行しているすべての人に伝わるはずです。
AI 支出を監査したとき、1 つの数字が際立っていました。その 80% 以上が、メイン エージェント ループでのフロンティア モデルの使用によるもので、ほぼすべてが Opus でした。
利用可能なノブをリストしました。
モデル。コード レビューの出力品質を左右する最大の要因は、Q&A ではそれほどではありません。
ターンします。キャッシュが長いループ用に最適化されると、予想よりも重要ではなくなります。
ツール。変動性があり、機能の出荷に応じて成長します。
キャッシング。すでに最適化されています。積極的なプロンプト キャッシュにより、キャッシュ ヒット率の回転率が 95% 以上に保たれるため、そこに再利用する意味のあるものは何も残されていませんでした。
プロンプトが表示されます。ほとんどが最適化されており、モデルの世代が進むごとに問題は少なくなります。
キャッシュとプロンプトは利用されませんでした。機能の削除は議題に上らなかった。それによりモデル自体が残され、フロンティア価格設定とオープンウェイト価格設定の間のギャップは無視できなくなりました。 GLM 5.2 の請求額はフロンティア料金の 3 分の 1 未満であり、サービス提供プロバイダー全体のトークンごとの生の価格に基づいて、Opus と比較して 95% 近くの節約になるように見えました。
問題は、実際のユーザーによって判断され、実際のワークロードで品質が本番環境に耐えられるかどうかでした。ベンチマークはそれに答えるつもりはありませんでした。
トークンごとの価格設定は自分に言い聞かせる嘘 #
最初の修正はすぐに来ました。当社のエンジニアの 1 人が Fireworks ダッシュボードから実際の使用量を引き出しました。23 億 5,000 万のトークンが、Opus での同じ量のコストの約 5% に相当し、キャッシュが両側で考慮されました。紙の上では95%のコスト削減。
次に、同じコミット、同じ PR、同じ PR で同じコード レビューを実行しました。

ompt、1 回は Opus で、もう 1 回は GLM で。実際の節約率は 95% ではなく 68% でした。
算術演算は乗算であり、料金表には 2 つの項のうち 1 つだけが印刷されます。
テキストのコピー 有効コスト = トークン価格 × タスクごとに必要なトークン 価格だけを見ると、GLM は 20 倍の改善のように見えます。請求書では約3.1倍でした。
ギャップは行動的なものです。 GLM は、Opus よりもタスクごとにはるかに多くのトークンを使用します。
より多くのツール呼び出しが行われます。より積極的に探求しますが、それは良いこともあれば、無駄なこともあります。
それはより多くの間違いを犯し、間違いはターンを犠牲にします。不正な形式の検索または空の結果は、ループを再度往復することを意味します。
より大きなツールの結果を取り込み、それらについて推論するためにより多くのトークンを費やします。
全体的にはもっとおしゃべりです。
これは少数のサンプルから得た逸話ではありません。私たちの検索フローでは、通話あたりの平均トークンは移行期間中に約 18,500 から約 24,700 (3 分の 1 増加) に増加し、GLM はそれらの通話の約 45% のみを処理していました。探査のオーバーヘッドはフリート全体の台帳で確認できます。
したがって、重要な指標は、トークンあたりのコストではなく、タスクあたりのコストです。モデル スワップを評価するチームは、同一のワークロードをエンドツーエンドで実行し、レート表ではなく請求書を比較する必要があります。 68% の削減は依然として大きな成果です。それは料金表が約束した勝利ではありません。
私たちは社内の雰囲気や公的ベンチマークを信頼していませんでした。私たちは比較を製品に組み込みました。
コードレビューのために、まったく同じ PR、同じコミット、同じプロンプトで 2 つのモデルを実行するデュアル パイプラインをセットアップし、両方のコメント セットを公開します。エンド ユーザーは、どのモデルがどのコメントを生成したかを知ることができません。すでにレビューコメントの反応センチメントを記録しており、それをパイプラインごとに帰属させることができるため、すべての反応がブラインド投票になりました。
センチメントは精度を測定します: 反応があったコメントの数

どの部分がマイナスでしたか?何千もの製品レビューの中で、3 つのモデルがきれいに分類されました。
精度に関しては、Opus が最高のランクにランクされ、GLM がそれに続き、GPT ベースの代替品は、レビューごとに Opus のほぼ 3 倍のコメント量を投稿しているにもかかわらず、はるかに遅れています。 GLM は単純に、もっと積極的に話そうとしたのです。コード レビューでは、これは危険です。レビュー担当者が不十分な発見を投稿すると、開発者はそれを無視するように訓練されます。これはツールにとって最悪の形です。 GLMの精度は価格を考えれば許容範囲内でした。 3番目のモデルはそうではありませんでした。
ただし、これが何を意味するかについては正直に言ってください。反応を引き出すコメントはほんの一部であり、良いコメントよりも悪いコメントの方が反応する動機が大きいため、感情の絶対的な数値には偏りがあり、信頼できるのは順序だけです。レビューごとのコメントもリコールされません。モデルは、実際のバグをさらに見つけることなく、より多くのコメントを書くことができます。私たちはこれを、同等の能力の証明としてではなく、ワークロードのデプロイメントシグナルとして使用しました。 「このモデルはバグを見つけることができますか?」質問ではなくなりました。疑問は、どのくらいの頻度で発見が実用的になるか、それに伴う誤検知の数はどれくらいか、モデルはコメントすべきではない時期を知っているか、その出力は開発者の信頼を築くのか、それとも損なうのか、というものでした。
私たちは Q&A に対しても同じ操作を実行しました。まず自分の組織に対して GLM を有効にし、全員に悪い回答にフラグを立てるように指示し、1 週間監視しました。その後、無料枠を最初に、大規模な顧客が最後に、段階的に展開します。
「OpenAI 互換」はスペクトルであり、標準ではありません。
実際にエンジニアリング時間のほとんどがここに費やされました。すべてのオープンウェイト サービス プロバイダーは、OpenAI 互換 API を宣伝しています。実際には、互換性は興味深い機能が始まる時点で終わります。
すべて本番中またはその前の週にヒットしたもの:
Reasoning には標準のワイヤ形式がありません。あなたを思って花火が戻ってくる

reasoning_content フィールドに入力すると、Together はそれをreasoning で返しますが、ファーストパーティの OpenAI はそれをチャット補完にまったく入れません。マルチターンのツール呼び出しループを実行する場合は、推論コンテンツを以前のアシスタント ターンにリプレイする必要もあり、各ホストはそれについて独自の期待を持っています。
推論労力マッピングにより、請求額が静かに最大値に達する可能性があります。 GLM のチャット テンプレートは、リテラル文字列 high 以外の努力値を最大努力値として扱いました。私たちは「最小限」を意味すると考えた値を送信しており、それを捕捉するまですべての呼び出しで最大限の努力を払って推論するために料金を支払っていました。
プロンプト キャッシュのキーはどこでも異なります。ファーストパーティ OpenAI は、専用のキャッシュ キー フィールドを使用します。 Fireworks と Together はそのフィールドを完全に無視し、ユーザー フィールドからキャッシュをキーオフします。これを発見するまで、マルチターン セッションでは、これらのホスト上のすべてのリクエストでキャッシュが欠落していました。その失敗がどのようなものであるかに注目してください。成功した応答、正しい出力、どこにもエラーはありません。プロバイダーごとのキャッシュ ヒット メトリックがなければ、通常のコスト差異として読み取られます。
使用量のアカウンティングに一貫性がありません。一部の OpenAI 互換 API には、 prompt_tokens 内にキャッシュされたトークンが含まれています。コスト パイプラインがばらばらのバケットを想定している場合、請求が二重になり、下流のすべてのダッシュボードが間違っています。
構造化出力のサポートはホストごとに異なり、監視している場合にのみ大声で失敗します。 Baseten は、response_format: json_schema を完全に拒否しました。その結果、7 日間で 336 件の構造化呼び出しのうち成功したのは 0 件で、そのトラフィックは完全に Fireworks にサイレントに固定されました。関連: 構造化出力リクエストと一緒にツール定義を送信すると、一部のプロバイダーでハード 400 が発生し、他のプロバイダーでは解析不能なツール呼び出し応答が発生しました。
ストリーミングの使用状況の統計は形状によって異なります。 vLLM 上で機能する Baseten は、すべての st の累積トークン数を報告します。

単一の最終合計ではなく、リーミングされたチャンク。単純に集計すると、記録された使用量にチャンク数が乗算されます。そのオプションを要求した場合でも、Fireworks は 400 を返します。
プロンプトはそのままでは転送されません。クロード用に調整されたプロンプトは、思考連鎖の足場、人間味のある構造化出力命令、人間型のツール定義を蓄積します。推論が第一級の API ノブであるモデルでは、「段階的に考える」足場は過剰に指定され、逆効果になります。プロンプトを通過させるのではなく、ターゲット モデル ファミリごとにプロンプ​​トを再構築する決定論的なアダプター層を構築しました。 「決定的」という言葉が適切です。そのパスにはモデルによって生成された変換が存在しないため、特定の入力、ターゲット モデル、およびプロバイダーは常に同じリクエストを生成します。
これらはどれも珍しいものではありません。教訓は、あるプロバイダーが API を正しく実装し、他のプロバイダーが正しく実装しなかったということではありません。エッジでは共有される動作はありません。実際に必要なのは、依存するすべての組み合わせ (モデル × プロバイダー × ストリーミング モード × 推論モード × ツールの使用 × 構造化された出力) に対する適合性スイートです。なぜなら、これらの失敗のそれぞれが、両方を測定していないと、品質またはユニット エコノミクスのいずれかを何週間も静かに損なうからです。
モデルを購入するのではなく、サービング スタックを購入するのです #
Anthropic では、モデルとサービス提供インフラストラクチャが 1 つの会社から提供されます。オープンウェイト モデルでは、ウェイトが商品であり、サービング スタックが製品です。 Fireworks、Togetter、および Baseten で同じ GLM 5.2 を実行しましたが、レイテンシ、安定性、機能サポート、価格の点でそれぞれで動作が異なりました。移行期間中、Fireworks のワークロードでは、Opus サービス パスよりも平均して約 30% 高いレイテンシが発生しました。これは、ユーザーに表示されるパフォーマンスに十分な差があります。

応答時間。
そこで、GLM をプロバイダーとして考えるのをやめ、プールとして扱い始めました。
1 つのルーティング層の背後にある複数の独立したサービスプロバイダーであり、それぞれがサーキットブレーカーでラップされています。プロバイダーの機能が低下すると、ブレーカーが開き、トラフィックが正常なレッグに移行します。ブレーカーの状態がフリート全体に伝播するため、すべてのノードが一緒に反応します。
私たちは Fireworks と Together の均等なラウンドロビンから開始し、コンプライアンスに配慮した静的インフラストラクチャの 3 番目の機能として Baseten を追加し、最終的に Together を削除しました。その後、偶数ローテーションが間違っていることに気付きました。Baseten はレート表のすべての行で Fireworks を下回り (約 20% 安い)、ベースラインでの実行速度が速かったため、均等分割では信頼性の向上がゼロで、より高価で遅いリクエストを購入していました。優先ルーティングに移行しました。Baseten をプライマリ、Fireworks をフェイルオーバーとして使用します。フェイルオーバーは意図的にスティッキーではないため、呼び出し元全体を移行するのではなく、一時的なブリップによって 1 つのリクエストが迂回されます。
厳選された優先順位は、依然として動的な質問に対する静的な答えです。料金表は変更され、今月最速のプロバイダーが来月にはならない可能性があります。これが、私たちが最終的に適応ルーティングに移行した理由です。本番環境で各プロバイダーのコスト、信頼性、速度を継続的にサンプリングします。

[切り捨てられた]

## Original Extract

Per-token pricing promised 95% savings; production delivered 68%. How we moved our agent loops from Claude Opus to GLM 5.2, and what broke on the way.

What We Learned Moving Our Agent Loops to GLM Products Outcomes Resources Customers Security Docs Pricing More Log In Get Started Book a Demo Products Context Engine Coding Agents and MCP AI Code Review Developer Q&A Outcomes Mergeable Code Save Tokens Reduce Rework Resources Blog Open Source Videos AI Adoption Assessment Customers Security Docs Pricing Log In Get Started Book a Demo All Articles What We Learned Moving Our Agent Loops from Anthropic to GLM
Why we moved most of Unblocked's agent traffic from Claude Opus to GLM 5.2, what the blind A/Bs and the ledger actually showed, and what broke on the way.
TL;DR: We moved most of Unblocked's agent traffic from Claude Opus to GLM 5.2.
“Per-token” math promised 95% savings; per-task production reality delivered 68%. Getting there took blind A/B evaluation on real code reviews, a circuit-broken multi-provider serving pool, and weeks of "OpenAI-compatible" surprises.
Background: Unblocked gives agents the organizational understanding of your best engineers. Its context engine connects and reconciles knowledge across your code, conversations, issues, documentation, and production systems, then makes that understanding available wherever work happens. Agents access context through MCP/CLI/APIs and Unblocked products like Code Review and Unblocked Code apply that context directly to reviewing and writing code. The result is better software, produced with fewer tokens, fewer iterations, and less back-and-forth.
For most of the last two years, Unblocked ran on Anthropic models. Our product is a set of agentic loops: answering engineering questions, reviewing code, searching across an organization's institutional knowledge. Claude Opus sat in the outer loop of nearly all of it.
This summer we moved most of that traffic to GLM, an open-weight model served by third-party inference providers. This post covers why we did it, what we measured, what broke, and where we've landed. We're not going to describe our system internals, but the findings should transfer to anyone running agent loops at scale.
When we audited our AI spend, one number stood out: over 80% of it came from frontier model usage in our main agent loops, almost entirely Opus.
We listed the knobs available to us:
Model. The single biggest factor in output quality for code review, less so for Q&A.
Turns. Less important than expected once caching is optimized for long loops.
Tools. Variable, and growing as features ship.
Caching. Already optimized. Aggressive prompt caching holds us above a 95% cache hit rate turn-over-turn, so there was nothing meaningful left to reclaim there.
Prompts. Mostly optimized, and mattering less with every model generation.
Caching and prompts were tapped out. Deleting features wasn't on the table. That left the model itself, and the gap between frontier pricing and open-weight pricing had become impossible to ignore. GLM 5.2 was billed at less than a third of frontier rates, and on raw per-token pricing across serving providers the savings looked closer to 95% against Opus.
The question was whether the quality held up in production, on real workloads, judged by real users. Benchmarks weren't going to answer that.
Per-token pricing is a lie you tell yourself #
The first correction came fast. One of our engineers pulled real usage off the Fireworks dashboard: 2.35 billion tokens served for roughly 5% of what the same volume would have cost on Opus, with caching accounted for on both sides. A 95% cost reduction, on paper.
Then we ran the same code review on the same commit, same PR, same prompt, once with Opus and once with GLM. The real-world saving was 68%, not 95%.
The arithmetic is multiplicative, and only one of the two terms is printed on the rate card:
Copy text effective cost = token price × tokens required per task On price alone, GLM looked like a 20× improvement. On the bill it was about 3.1×.
The gap is behavioral. GLM uses far more tokens per task than Opus does:
It makes more tool calls. It explores more aggressively, which is sometimes good and sometimes waste.
It makes more mistakes, and mistakes cost turns. A malformed search or an empty result means another round trip through the loop.
It ingests larger tool results and spends more tokens reasoning about them.
It's more talkative in general.
This isn't an anecdote from a handful of samples. In our search flow, average tokens per call rose from about 18,500 to about 24,700 (up a third) over the migration window, and GLM was only serving around 45% of those calls. The exploration overhead is visible in the fleet-wide ledger.
So the metric that matters is cost per task, not cost per token. Any team evaluating a model swap should run identical workloads end to end and compare invoices, not rate cards. A 68% reduction is still an enormous win. It's just not the win the rate card promised.
We didn't trust internal vibes, and we didn't trust public benchmarks. We built the comparison into the product.
For code review, we set up dual pipelines that run two models on the exact same PR, same commit, same prompt, and publish both sets of comments. End users can't tell which model produced which comment. We already recorded reaction sentiment on review comments and could attribute it per pipeline, so every reaction became a blind vote.
Sentiment measures precision: of the comments that got a reaction, what fraction were negative? Across thousands of production reviews, three models sorted cleanly:
On precision, Opus ranked best, GLM close behind, and the GPT-based alternative far behind despite posting nearly triple Opus's comment volume per review. GLM was simply more willing to speak. In code review that's dangerous: a reviewer that posts weak findings trains developers to ignore it, which is the worst possible shape for the tool. GLM's precision was within a tolerance we could accept for the price. The third model wasn't.
Be honest about what this measures, though. Only a fraction of comments ever draw a reaction, and a bad comment is more motivating to react to than a good one, so the absolute sentiment numbers are biased and only the ordering is trustworthy. Comments per review is not recall either — a model can write more comments without finding more real bugs. We used this as a deployment signal for our workload, not as proof of equal capability. "Can this model find bugs?" stopped being the question. The questions were: how often is a finding actionable, how many false positives ride along with it, does the model know when not to comment, and does its output build or erode developer trust?
We ran the same play for Q&A: enable GLM for our own organization first, tell everyone to flag bad answers, watch for a week. Then roll out gradually, free tier first, largest customers last.
"OpenAI-compatible" is a spectrum, not a standard #
This is where most of the engineering time actually went. Every open-weight serving provider advertises an OpenAI-compatible API. In practice, compatibility ends where the interesting features begin.
Things we hit, all in production or in the week before it:
Reasoning has no standard wire format. Fireworks returns thinking output in a reasoning_content field, Together returns it in reasoning , and first-party OpenAI doesn't put it in chat completions at all. If you run multi-turn tool-calling loops, you also need to replay reasoning content back onto prior assistant turns, and each host has its own expectations about that.
Reasoning effort mappings can silently max out your bill. GLM's chat template treated any effort value other than the literal string high as maximum effort. We had been sending a value we thought meant "minimal" and were paying for maximum-effort reasoning on every call until we caught it.
Prompt caching is keyed differently everywhere. First-party OpenAI uses a dedicated cache-key field. Fireworks and Together ignore that field entirely and key the cache off the user field. Until we discovered this, multi-turn sessions were missing cache on every request on those hosts. Note what that failure looks like: successful responses, correct output, no errors anywhere. Without per-provider cache-hit metrics it reads as ordinary cost variance.
Usage accounting is inconsistent. Some OpenAI-compatible APIs include cached tokens inside prompt_tokens . If your cost pipeline assumes disjoint buckets, you double-bill yourself and every dashboard downstream is wrong.
Structured output support varies per host, and fails loudly only if you're watching. Baseten rejected response_format: json_schema outright. The result: zero of 336 structured calls succeeded over seven days, silently pinning that traffic entirely to Fireworks. Related: sending tool definitions alongside a structured-output request caused hard 400s on some providers and unparseable tool-call responses on others.
Streaming usage stats differ in shape. Baseten, which serves on vLLM, reports cumulative token counts on every streamed chunk rather than a single final total. Aggregate naively and you multiply your recorded usage by the chunk count. Fireworks returns 400 if you even request that option.
Prompts don't transfer verbatim. Prompts tuned for Claude accumulate chain-of-thought scaffolding, Anthropic-flavored structured-output instructions, and Anthropic-shaped tool definitions. On models where reasoning is a first-class API knob, "think step by step" scaffolding is over-specified and counterproductive. We built a deterministic adapter layer that restructures prompts per target model family instead of passing them through. Deterministic is the operative word: there is no model-generated translation in that path, so a given input, target model, and provider always produce the same request.
None of these are exotic, and the lesson isn't that one provider implemented the API correctly and the others didn't. There is no shared behavior at the edges. What you actually need is a conformance suite over every combination you depend on — model × provider × streaming mode × reasoning mode × tool use × structured output — because each of these failures quietly corrupts either your quality or your unit economics for weeks if you aren't measuring both.
You're not buying a model, you're buying a serving stack #
With Anthropic, the model and the serving infrastructure come from one company. With open-weight models, the weights are a commodity and the serving stack is the product. We ran the same GLM 5.2 across Fireworks, Together, and Baseten, and it behaved differently on each in latency, stability, feature support, and price. Over our migration window, Fireworks averaged roughly 30% higher latency on our workloads than our Opus serving paths — enough of a gap to show up in user-visible response times.
So we stopped thinking of GLM as a provider and started treating it as a pool:
Multiple independent serving providers behind one routing layer, each wrapped in a circuit breaker. When a provider degrades, its breaker opens and traffic shifts to healthy legs; breaker state propagates across our fleet so every node reacts together.
We started with an even round-robin across Fireworks and Together, added Baseten as a third leg for its compliance-friendly static infrastructure, and eventually dropped Together. Then we realized even rotation was wrong: Baseten undercut Fireworks on every line of the rate card (roughly 20% cheaper) and ran faster at baseline, so an even split was buying more expensive, slower requests for zero reliability gain. We moved to priority routing: Baseten as primary, Fireworks as failover. Failover is deliberately not sticky, so a transient blip diverts one request instead of migrating the whole caller.
Hand-picked priority is still a static answer to a moving question — rate cards change, and a provider that is fastest this month may not be next month. Which is why we eventually moved to adaptive routing: continuously sampling each provider's cost, reliability, and speed in prod

[truncated]
