---
source: "https://ploy.ai/blog/migrating-a-production-ai-agent-to-gpt-5-6"
hn_url: "https://news.ycombinator.com/item?id=48882716"
title: "Migrating a production AI agent to GPT-5.6: 2.2x faster, 27% cheaper"
article_title: "Migrating a production AI agent to GPT-5.6 | Ploy"
author: "brryant"
captured_at: "2026-07-12T17:52:00Z"
capture_tool: "hn-digest"
hn_id: 48882716
score: 3
comments: 0
posted_at: "2026-07-12T17:13:07Z"
tags:
  - hacker-news
  - translated
---

# Migrating a production AI agent to GPT-5.6: 2.2x faster, 27% cheaper

- HN: [48882716](https://news.ycombinator.com/item?id=48882716)
- Source: [ploy.ai](https://ploy.ai/blog/migrating-a-production-ai-agent-to-gpt-5-6)
- Score: 3
- Comments: 0
- Posted: 2026-07-12T17:13:07Z

## Translation

タイトル: 実稼働 AI エージェントの GPT-5.6 への移行: 2.2 倍高速、27% 安価
記事のタイトル: 実稼働 AI エージェントを GPT-5.6 に移行する |策略
説明: 私たちはフロンティアモデルに高いハードルを課しましたが、4 か月間クロード・オーパスに勝るものはありませんでした。 GPT-5.6はそうでした。これがあればよかったと思う移行ガイドです。

記事本文:
実稼働 AI エージェントを GPT-5.6 に移行する |製品を策略する
製品リソース ソリューション 会社の価格設定 展開効果 ログイン 無料で始める ブログに戻る 運用 AI エージェントの GPT-5.6 への移行
私たちはフロンティアモデルに高いハードルを課し、4か月間クロード・オーパスに勝るものはありませんでした。 GPT-5.6はそうでした。これがあればよかったと思う移行ガイドです。
今日の時点で、Ploy のエージェントは、今朝リリースされた OpenAI モデル ファミリのフラッグシップ層である GPT-5.6 Sol 上で実行されています。信じられないほど高い品質基準を考慮すると、何ヶ月もの間、クロード・オーパスに匹敵するモデルを見つけることができませんでした。それは GPT 5.6 Sol で変わりました。 Claude Opus との直接対決を実行した後、GPT 5.6 Sol をすべての Ploy ワークスペースを強化するデフォルト モデルにしました。
それは思っているよりも大きなスイッチです。 Ploy のエージェントは、実際のマーケティング Web サイトを構築および編集します。ページを計画し、コードベースを読み取り、コンポーネントを書き込み、画像を生成し、独自の作業のスクリーンショットを撮り、いつ完了したかを決定します。この職務内容はモデルに対して非常に高いハードルを設定しており、私たちはすべてのフロンティア リリースをそれに対してテストします。 4 か月間、Opus はデフォルトのスロット (最初は Opus 4.7、次に 4.8) を保持していましたが、テストしたものでこれを超えるものはありませんでした。 GPT-5.6はそれを実現した最初のモデルです。
最初の評価実行が完璧だったわけではありません。実際の故障モードがありましたので、それをお見せします。しかし、それは非常にうまくいき、約束は即座かつ具体的でした。ビルドは実時間の半分以下で完了し、コストは 27% 低く、完了した作業に関しては既存のスコア以上のスコアを獲得しました。このような数字は、実際の移行作業のモデルを買うのに役立ちます。
ユニバーサル LLM SDK である Vercel の AI SDK を使用しているにもかかわらず、Claude Opus 4.8 から GPT 5.6 Sol に切り替えるには、一度に 1 つの評価失敗で、私たちが「モデル」と考えているものがプロバイダー固有の動作であることを発見する必要がありました。

ack は、ツールの引数をどのように埋めるか、プロンプト キャッシュがどのように機能するか、ターン間で独自の推論をどのように再生するかという点に密かに特化しています。
必要なことは次のとおりです。評価ハーネスを修正し、次にツール スキーマを修正し、次にキャッシュを修正し、次に推論再生を修正します。
ステップ 0: 単一の番号を信頼する前にハーネスを修正する
私たちの評価スイートは、実際のフィクスチャ ワークスペースに対して実際のエージェントを実行します。 「ホームページを最初から構築する」から「このクローン リクエストは安全に実行できるか」に至るまで、何百ものケースが含まれます。ビルド ケースは、リファレンス デザインに対してバイナリ チェックを実行する視覚的な審査員によって採点されます。「ヒーローはフルブリードの写真シーンである」または「主要な CTA は錠剤ではなく丸い長方形である」などの 10 個のはい/いいえの質問に加えて、コンテンツ チェック、ツール軌道チェック、およびファイル アサーションが行われます。失敗したすべてのケースは、スコアだけでなく、実際のツール呼び出しとモデル テキストの完全なトレースに基づいてトリアージされます。
2 つのモデル ファミリにわたってそのスイートを実行したことは、個々の結果よりも私たちを驚かせました。
あなたのハーネスは現在のモデルに合わせて調整されていますが、あなたはそれを知りません。私たちのツールコールの予算は、Opus のシーケンシャル スタイルに合わせて設定されました。 GPT-5.6 は並列呼び出しをファンアウトし、正しく解決していたケースではそれらを吹き飛ばしました。私たちの eval executor はバッチ ファイル読み取りをサポートしていませんでした。これは Opus ではほとんど使用されず、GPT-5.6 では常に使用されていました。最初のクロスモデル実行における生の障害の約 3 分の 1 は、モデルの動作ではなくハーネスの仮定に遡り、それらはモデル間で均等に分散されていませんでした。既存モデルに対してチャレンジャー モデルを評価している場合は、合格率を信頼する前にトレースを優先順位付けしてください。そうしないと、古いモデルをどれだけ模倣しているかによって新しいモデルを評価することになります。
eval ではモデルを公平に評価していることを確認してください。 minScore しきい値を省略したデータセットは、暗黙的に d を継承しました。

デフォルトは 1.0 なので、GPT-5.6 はスコア 0.98 のヒーローを「不合格」にし、Opus はすべての個別のチェックをパスしながらケースを「不合格」にしました。 2 つの防御可能な設計方向。目に見えない閾値が 1 つあります。
第一印象：すぐに期待できる
ハーネスがクリーンアップされたので、エージェントがリファレンス デザインに基づいてブランドのホームページを再構築する再デザイン スイートのサンプルを次に示します。
これは約束の形です。完成したページまでの速度は 2.2 倍、コストは 27%、出力トークンは約半分になります。 GPT-5.6 は無駄のないコードを記述します。 1 つの一致したペアで、Opus は 174 個の CSS 変数 (フルカラー ランプ、ほとんど未使用) を含む 17,957 文字の globals.css を生成しました。GPT-5.6 では、同等の (場合によってはより優れた) レンダリングされたページ用に 2,508 文字と 45 個の変数が書き込まれました。
デザイン: シャープ、クリーン、しかし少し均一
GPT-5.6 のデザイン作業についての全体的な見解は、クリーンでモダン、しっかりとグリッドが配置されたレイアウトが非常に優れていますが、うまく操作しない限り、その外観に収束する傾向があります。 Opus 4.8 用に設計された古いハーネスを使用する GPT 5.6 Sol は、既存の設計システムを無視する傾向があり、代わりにシャープで抑制された、目に見えて一般的な出力を生成します。
これをどのように修正したかの詳細については、別のブログ投稿に記載する価値があります。当社の設計チームとエンジニアリング チームの専門知識により、私たちは、そのままでは得られない世界クラスのブランドの支持を達成するようにモデルを導くことができます。
これは、私たちが発見する前に静かに結果を破壊していたものです。
エージェントのコード ツールには 25 のトップレベル パラメーターがあり、1 つは必須 ( action )、残りはオプションです。クロードは使用している 2 つまたは 3 つを送信し、残りを省略します。 GPT-5.6 は、毎回 25 個すべてを送信し、必要のない値に対して妥当な値を作成します: offset: 0 、 timeout: 120000 、 siteId: "00000000-0000-0000-0000-0000000000000" 。
3日間の本番

アクション トレース、すべてのプロパティを運ぶ code(read) 呼び出し:
問題は冗長さではありません。それは、発明された価値は意図された価値と区別できないということです。 offset: 0 は実際の引数のように見えます。私たちのファイル読み取り実装はそれを 1 つとして扱い、そのために GPT-5.6 のファイル読み取りの 52% ～ 64% が空に返されていました。ツールは両方の方法で success: true を返したので、モデルは空のファイルを読み取っていることを知る方法がありませんでした。通話が増えて仕事が悪化するだけでした。
プロンプトを表示してもこれは解決しません。 「未使用のパラメータを省略する」というツール説明ディレクティブ: 依然として 25/25。プロパティごとの「オプション、未使用の場合は省略」ヒント: 依然として 25/25。 OpenAI の厳密モード: 同一の動作 (私たちが測定した)。これを採用すると、すべてのスキーマから pattern 、 format 、および配列バインドされた検証を削除する必要がありました。これは、モデルが関数呼び出しを発行する方法に組み込まれています。あなたはそれを遠ざけるように指示するのではありません。あなたはそれを中心にデザインします。
有効な修正は、プロバイダー境界でのスキーマ変換です。 OpenAI ファミリ モデルの場合のみ、 anyOf: [T, null] を使用して、すべてのオプションのプロパティを必須であるが nullable に書き換えます。これにより、モデルに「これを使用しない」という明示的な方法が与えられます。次に、すべてのツール呼び出しが通過する単一の継ぎ目で、検証前に null が取り除かれるため、ツールの実装はまったく変更されません。往復: モデルは誠実さが表現可能なスキーマを認識します。ツールは常に同じ入力を参照します。
// 前: 25 個のキー、それぞれに発明された値が含まれる
{ "アクション" : "読み取り" 、 "ファイルパス" : [ ... ]、 "オフセット" : 0 、 "タイムアウト" : 120000 , ... }
// 後: 25 個のキー、4 個の実数値、21 個の明示的な null (ツールの実行前に削除)
{ "アクション" : "読み取り" 、 "ファイルパス" : [ ... ]、 "オフセット" : null 、 "タイムアウト" : null 、 ... }
結果: 空のファイルの読み取りが 52% から増加しました。

エージェントが同じ作業に必要なツール呼び出しが約 30% 減りました。これは、空白に戻ったファイルを再度読み取る必要がなくなったためです。
ステップ 2: プロンプト キャッシュを再構築する
これはエンジニアリング上の最も有益な違いでした。なぜなら、表面的には両方のプロバイダーが「即時キャッシュ」を提供しており、その言葉には 2 つのまったく異なる設計が隠されているからです。慎重に移行する場合は、次のようにしてください。移行する前は、GPT-5.6 は Opus よりも約 50% 高価に見えました。それはモデルの価格ではありませんでした。それはキャッシュ構成でした。
エージェントのプロンプトは、すべての会話で同一の約 29,000 トークン (ツール スキーマとコア システム プロンプト) の静的プレフィックスで始まります。 Claude では、cache_control でキャッシュ ブレークポイントをマークし、その接頭辞が組織全体のキャッシュに適用されます。あらゆる会話、あらゆるワークスペース、1 つの共有エントリで、スループット バジェットを考慮する必要はありません。キャッシュ ヒット率は 92% ～ 96% であり、キャッシュはバックグラウンドに消えていきます。
GPT-5.6 は、OpenAI のキャッシュ モデルを私たちの管理下から変更しました。以前の GPT モデルは、部分的なプレフィックス一致で暗黙的にキャッシュされていたため、無料でまともなヒット率が得られました。 GPT-5.6 では部分プレフィックス マッチングが削除されました。暗黙的キャッシュでは、最新のメッセージをキーとしたプロンプト エントリ全体のみが作成されるようになりました。 29K の静的プレフィックスを共有する新しい会話は、その 0% をキャッシュしました。すべての会話では、キャッシュされていないレートで完全なプレフィックスが再請求され、GPT-5.6 では、キャッシュを使用するかどうかに関係なく、キャッシュされていないすべてのプロンプトにも 1.25 倍のキャッシュ書き込み追加料金が支払われます。
意図されたメカニズムは明示的です。プロンプトキャッシュ_ブレークポイントマーカーと必須のプロンプトキャッシュ_キーです。そして重要なのは、設計が実際にどこで分岐するかということです。なぜなら、それはキャッシュのアイデンティティの一部だからです。同じプロンプト、異なるキー: キャッシュ ヒットはゼロ。各キーは、1 分あたり約 15 リクエストをサポートするキャッシュ ノードにマップされます。

以前は、OpenAI は独立したコールド キャッシュを使用してトラフィックを他のノードに送っていました。
これにより、「キャッシュを有効にする」ということが実際の設計上の決定に変わります。つまり、キーのスコープをどのエンティティに設定するかということになります。
会話ごとのキーは、新しい会話が共有プレフィックスにヒットしないことを意味します。ファーストコールのヒット率: 0%。 (この間違いを測定しました。高価です。)
1 つのグローバル キーは、すべてのリクエストが 1 つのキャッシュ ノードにハッシュされることを意味し、運用トラフィックにより 15 rpm バジェットが消滅します。リクエストがコールドノードに溢れてしまい、また失敗してしまいます。
ワークスペースごとのキーがスイートスポットです。顧客ワークスペース内のすべての会話はエントリを共有します。キーごとのトラフィックは低いままです。
ワークスペース スコープのキーを出荷し、システム プロンプトをブレークポイント付きのレイヤーに分割し、Anthropic で既に使用した構造をミラーリングします。
リクエスト ──► ハッシュ(プロンプト ヘッド + プロンプト キャッシュ キー) ──► キャッシュ ノード (キーあたり最大 15 リクエスト/分)
│
┌───────────────────────┴─────────┐
│ ノード上のエントリ。すべてキー ws:{workspaceId} によって名前空間が設定されています │
│ │
│ [ツール + 静的プレフィックス].................................................................................................... セッションごと │
│ [ ツール + 静的プレフィックス + ワークスペース コンテキスト ]... B 同じコンテキスト │
│ [ · · · · · · · · · + ターン 1 + … + 最新 ] C このセッション │
━━━━━━━━━━━━━━━━━━━━━━━┘
エントリ A は、セッションの最初のコールを安くするものです。エントリ B は自己修復します。ワークスペースのメモリが変化すると、リクエストは B をミスしますが、A にはヒットし、新しい B を書き込みます。

完全な 29K の再請求の代わりに、コンテキストサイズの書き込みが行われます。エントリ C は OpenAI の暗黙的なプロンプト全体のチェーンであり、プロンプトは厳密に追加専用であるため、セッション内で正常に機能します。
結果の 1 つは回避策がありません。OpenAI では静的プレフィックスをワークスペース間で共有することが構造的に不可能です。 Anthropic は、そのキャッシュがキー分割なしで組織スコープであるため、それを共有できます。 GPT-5.6 では、すべてのワークスペースは、アイドル ウィンドウごとに 1 回の 29K コールド書き込み (約 0.18 ドル) を支払います。実際のコストですが、制限があり、予測可能です。
変更後の結果: 初回呼び出しのキャッシュ ヒットは約 0% から 83.7% に減少し、キャッシュされていない入力トークンの合計は 28% 減少し、GPT-5.6 のスイートごとのコストは Opus のコストを下回りました。私たちが注目していたギャップは、モデルの価格設定ではなく、キャッシュの構成ミスでした。モデルのコストを比較していて、そのうちの 1 つにコールド キャッシュがある場合は、モデルではなく構成を比較していることになります。
ステップ 3: 推論リプレイを自己完結型にする
短くなりましたが、実際の会話は中断されました。 GPT-5.6 の Response API は、デフォルトで前のターンの推論をサーバー側のアイテム参照として再生します。私たちのものは、会話の途中で断続的に失敗し始め、アイテム「rs_...」が見つかりませんでした。修正は store: false です。これにより、SDK は暗号化された推論コンテンツを要求し、サーバー状態へのポインターの代わりに自己完結型 BLOB を再生します。その結果、午後のデバッグに時間がかかりました。サーバー側の推論状態がループ内にあると、送信するバイトが追加専用である場合でも、効果的なプロンプトが上流で変更される可能性があります。
あなたも試してみてください。 GPT-5.6 は本日リリースされ、すでに Ploy 上で稼働しています。今すぐ無料で試すことができます。Web サイトを構築して、4 分以内に構築される様子を確認してください。 ploy.ai で無料で始めましょう。
プロイはマーケティングを行っています

[切り捨てられた]

## Original Extract

We hold frontier models to a high bar, and for four months nothing beat Claude Opus. GPT-5.6 did. Here's the migration guide we wish we'd had.

Migrating a production AI agent to GPT-5.6 | Ploy Product
Product Resources Solutions Company Pricing T h e P l o y E f f e c t Log in Start Free Back to blog Migrating a production AI agent to GPT-5.6
We hold frontier models to a high bar, and for four months nothing beat Claude Opus. GPT-5.6 did. Here's the migration guide we wish we'd had.
As of today, Ploy’s agent runs on GPT-5.6 Sol, the flagship tier of the model family OpenAI released this morning. For months, we couldn’t find a model that challenges Claude Opus given our incredibly high bar for quality. That changed with GPT 5.6 Sol. After running it head-to-head against Claude Opus, we’ve made GPT 5.6 Sol the default model powering every Ploy workspace.
That’s a bigger switch than it sounds. Ploy’s agent builds and edits real marketing websites. It plans a page, reads the codebase, writes components, generates imagery, screenshots its own work, and decides when it’s done. That job description sets a very high bar for a model, and we test every frontier release against it. For the four months Opus held the default slot (first Opus 4.7, then 4.8), nothing we tested beat it. GPT-5.6 is the first model that did.
Not that the first eval run was perfect. It had real failure modes, which we’ll show you. But it did extremely well, and the promise was immediate and specific: builds finishing in less than half the wall-clock time, at 27% lower cost, scoring at or above our incumbent on completed work. Numbers like that buy a model a real migration effort.
Despite using Vercel’s AI SDK , a universal LLM SDK, switching from Claude Opus 4.8 to GPT 5.6 Sol required discovering, one eval failure at a time, that the things we think of as “the model” are provider-specific behaviors our whole stack has quietly specialized around: how it fills in tool arguments, how its prompt cache works, how it replays its own reasoning between turns.
Here’s what it took: fix the eval harness, then the tool schemas, then caching, then reasoning replay.
Step 0: Fix your harness before you trust a single number
Our eval suite runs the real agent against real fixture workspaces. Hundreds of cases, from “build a homepage from scratch” to “is this clone request safe to execute.” Build cases are scored by a visual judge running binary checks against a reference design, ten yes/no questions like “the hero is a full-bleed photographic scene” or “primary CTAs are rounded rectangles, not pills” , plus content checks, tool-trajectory checks, and file assertions. Every failed case gets triaged against its full trace: the actual tool calls and model text, not just the score.
Running that suite across two model families surprised us more than any individual result:
Your harness is tuned to your incumbent model, and you don’t know it. Our tool-call budgets were sized for Opus’s sequential style; GPT-5.6 fans out parallel calls and blew through them on cases it was solving correctly. Our eval executor didn’t support batched file reads, which Opus rarely used and GPT-5.6 uses constantly. Roughly a third of the raw failures in the first cross-model run traced back to harness assumptions, not model behavior, and they were not evenly distributed between the models. If you’re evaluating a challenger model against an incumbent, triage the traces before you trust the pass rate . Otherwise you’re grading the new model on how well it imitates the old one.
Make sure you’re grading models fairly in evals. A dataset that omitted its minScore threshold silently inherited a default of 1.0, so GPT-5.6 “failed” a hero it scored 0.98 on, and Opus “failed” a case while passing every individual check. Two defensible design directions; one invisible threshold.
First impression: immediately promising
With the harness cleaned up, here’s a sample from our redesign suite, where the agent rebuilds a brand’s homepage against a reference design:
This is the shape of the promise: 2.2× faster to a finished page, 27% cheaper, and about half the output tokens. GPT-5.6 writes lean code. On one matched pair, Opus produced a 17,957-character globals.css with 174 CSS variables (full color ramps, mostly unused) where GPT-5.6 wrote 2,508 characters and 45 variables for a comparable (and sometimes better) rendered page.
Design: sharp, clean, but a little bit uniform
Our overall read on GPT-5.6’s design work: it is very good at clean, modern, tightly-gridded layouts, but it tends to converge towards that look unless you steer it well. With our older harness designed for Opus 4.8, GPT 5.6 Sol tends to ignore existing design systems and instead produces sharp, restrained, and visibly generic output.
The details of how we fixed this are worth a separate blog post of its own. With the expertise of our design and engineering teams, we are able to steer models to achieve world-class brand adherence that you can’t get out of the box.
Here’s the one that was silently corrupting results before we caught it.
Our agent’s code tool has 25 top-level parameters, one required ( action ) and the rest optional. Claude sends the two or three it’s using and omits the rest. GPT-5.6 sends all 25, every time , inventing plausible values for the ones it doesn’t need: offset: 0 , timeout: 120000 , siteId: "00000000-0000-0000-0000-000000000000" .
Three days of production traces, code(read) calls carrying every property:
The problem isn’t verbosity. It’s that an invented value is indistinguishable from an intended one . offset: 0 looks like a real argument. Our file-read implementation treated it as one, and 52% to 64% of GPT-5.6’s file reads were coming back empty because of it. The tool returned success: true both ways, so the model had no way to know it was reading blank files. It just did the work worse, with more calls.
Prompting doesn’t fix this. A tool-description directive to “omit unused parameters”: still 25/25. Per-property “OPTIONAL, omit if unused” hints: still 25/25. OpenAI’s strict mode: identical behavior (we measured), and adopting it would have forced us to strip pattern , format , and array-bound validation from every schema. This is baked into how the model emits function calls . You don’t instruct it away; you design around it.
The fix that worked is a schema transform at the provider boundary. For OpenAI-family models only, we rewrite every optional property to be required but nullable , using anyOf: [T, null] , which gives the model an explicit way to say “not using this.” Then, at the single seam every tool invocation passes through, we strip the nulls back out before validation, so no tool implementation changes at all. Round trip: the model sees a schema where honesty is expressible; the tools see the same inputs they always did.
// Before: 25 keys, every one carrying an invented value
{ "action" : "read" , "file_paths" : [ ... ], "offset" : 0 , "timeout" : 120000 , ... }
// After: 25 keys, 4 real values, 21 explicit nulls (stripped before the tool runs)
{ "action" : "read" , "file_paths" : [ ... ], "offset" : null , "timeout" : null , ... }
Results: empty file reads went from 52% to 0%, and the agent needed roughly 30% fewer tool calls for the same work, because it was no longer re-reading files that came back blank.
Step 2: Rebuild prompt caching
This was the most instructive engineering difference, because on the surface both providers offer “prompt caching” and the words hide two entirely different designs. If you migrate one thing carefully, make it this: before we did, GPT-5.6 looked about 50% more expensive than Opus. It wasn’t the model’s pricing; it was our cache configuration.
Our agent’s prompt opens with a static prefix of roughly 29K tokens (tool schemas plus the core system prompt) that’s identical for every conversation. On Claude, we mark cache breakpoints with cache_control and that prefix caches across the whole organization : any conversation, any workspace, one shared entry, no throughput budget to think about. Cache hit rates run 92% to 96% and caching fades into the background.
GPT-5.6 changed OpenAI’s caching model out from under us. Earlier GPT models cached implicitly on partial prefix matches, which gave decent hit rates for free. GPT-5.6 dropped partial-prefix matching : implicit caching now only creates whole-prompt entries keyed on the latest message. A new conversation sharing our 29K static prefix cached 0% of it. Every conversation re-billed the full prefix at the uncached rate, and on GPT-5.6 every uncached prompt also pays a 1.25× cache- write surcharge, whether or not you use caching.
The intended mechanism is explicit: prompt_cache_breakpoint markers plus a mandatory prompt_cache_key . And the key is where the design really diverges, because it’s part of cache identity. Identical prompt, different key: zero cache hits. Each key maps to a cache node that sustains roughly 15 requests per minute before OpenAI fans traffic to other nodes with independent, cold caches.
That turns “enable caching” into an actual design decision: what entity do you scope the key to?
Per-conversation key means a new conversation never hits the shared prefix. First-call hit rate: 0%. (We measured this mistake. It’s expensive.)
One global key means every request hashes to one cache node, and production traffic obliterates the 15 rpm budget; requests spill to cold nodes and you’re back to misses.
Per-workspace key is the sweet spot. All conversations in a customer workspace share entries; per-key traffic stays low.
We ship the workspace-scoped key and split the system prompt into breakpointed layers, mirroring the structure we already used for Anthropic:
request ──► hash(prompt head + prompt_cache_key) ──► cache node (~15 req/min per key)
│
┌──────────────────────────────────────────────────────┴───────────────┐
│ entries on the node, all namespaced by key ws:{workspaceId} │
│ │
│ [ tools + static prefix ]······················ A every session │
│ [ tools + static prefix + workspace context ]·· B same context │
│ [ ····················· + turn 1 + … + latest ] C this session │
└──────────────────────────────────────────────────────────────────────┘
Entry A is what makes a session’s first call cheap. Entry B self-heals: when workspace memory changes, the request misses B but still hits A, then writes a fresh B. One context-sized write instead of a full 29K re-bill. Entry C is OpenAI’s implicit whole-prompt chain, which works fine within a session because our prompts are strictly append-only.
One consequence has no workaround: cross-workspace sharing of the static prefix is structurally impossible on OpenAI. Anthropic can share it because its cache is org-scoped without key partitioning. On GPT-5.6, every workspace pays one 29K cold write per idle window, about $0.18. A real cost, but bounded and predictable.
Results after the change: first-call cache hits went from roughly 0% to 83.7% , total uncached input tokens dropped 28%, and GPT-5.6’s per-suite cost landed below Opus’s. Every dollar of the gap we’d been staring at was cache misconfiguration, not model pricing. If you’re cost-comparing models and one of them has a cold cache, you are comparing your config, not the models.
Step 3: Make reasoning replay self-contained
Shorter, but it broke real conversations. GPT-5.6’s Responses API replays prior-turn reasoning as server-side item references by default; ours started intermittently failing mid-conversation with Item 'rs_...' not found . The fix is store: false , which makes the SDK request encrypted reasoning content and replay self-contained blobs instead of pointers to server state. A corollary that cost us a debugging afternoon: with server-side reasoning state in the loop, the effective prompt can change upstream of you even when the bytes you send are append-only.
Try it yourself. GPT-5.6 launched today, and it’s already live on Ploy. You can try it for free right now: give it a website to build and see what a sub-four-minute build looks like. Start free at ploy.ai .
Ploy is marketing

[truncated]
