---
source: "https://blog.kilo.ai/p/sota-models-should-learn-to-kiss"
hn_url: "https://news.ycombinator.com/item?id=48497811"
title: "SOTA Models Should Learn to KISS (Claude Fable 5 Experiment)"
article_title: "SOTA Models Should Learn to KISS (Fable 5 Experiment)"
author: "justiceforsaas"
captured_at: "2026-06-12T01:02:27Z"
capture_tool: "hn-digest"
hn_id: 48497811
score: 1
comments: 0
posted_at: "2026-06-11T23:25:13Z"
tags:
  - hacker-news
  - translated
---

# SOTA Models Should Learn to KISS (Claude Fable 5 Experiment)

- HN: [48497811](https://news.ycombinator.com/item?id=48497811)
- Source: [blog.kilo.ai](https://blog.kilo.ai/p/sota-models-should-learn-to-kiss)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T23:25:13Z

## Translation

タイトル: SOTA モデルは KISS を学ぶべき (Claude Fable 5 実験)
記事のタイトル: SOTA モデルは KISS を学ぶべき (寓話 5 の実験)
説明: 最新のコーディング モデルで私がいつも気づいていることの 1 つは、その間違いがますます… 興味深いものになっているということです。

記事本文:
SOTA モデルは KISS を学ぶべきです (寓話 5 の実験)
キロのブログ
SOTA モデルは KISS を学ぶ必要がある (寓話 5 の実験)
Darko 2026 年 6 月 11 日 シェア 最新のコーディング モデルで私がいつも気づいていることの 1 つは、その間違いがますます… 興味深いものになっているということです。
古い AI モデルでは、間違いを見つけるのは非常に簡単でした。モデルは API を幻覚させたり、フレームワークを誤解したり、コンパイルされないコードを書いたり、自信を持ってアプリの間違った部分を操作したりしました。問題は非常に明らかなので、その提案をすぐに拒否できます。
Claude Fable 5 のような新しいモデルでは、故障モードはより微妙です。
Fable 5 がワンショット 3D ゲーム、クールなグラフィック、アニメーションなどをどのように実現できるかについて、誇張された主張を何度も目にしました。しかし、多くの人がコーディングにクロード モデルを使用しているため、コーディング タスクを与えてベンチマークを行うことにしました。
まず始めに、製品をリスト/追加できる基本的な Next.js アプリを作成しました。
アプリには、/products と /admin という 2 つの主要なルートがありました。また、/api/products に内部 API ルートもありました (これはほとんどの場合、実稼働アプリでは必要ないことはわかっていますが、モデルでは物事をより複雑にしたかったのです)。
/admin ページには、製品を追加するためのフォームが含まれていました。 (つまり、モデルに関して) 物事をより複雑にするために、フォームを /api/products ルート ハンドラーにポストするようにしました。ここで製品はメモリ内のデータ ストアに追加され、/products に戻って新しく追加されたアイテムが表示されると期待されるたびに、その変更を反映する必要がありました。バグの読み取り側は次のようになります。
const res = await fetch(`${baseUrl}/api/products?${params}`, {
明確にしておきますが、バグはキャッシュされたフェッチではなく (このコードだけでは何も問題ありません)、/products が読み取りを 60 秒間キャッシュし、

管理書き込みパスは、そのキャッシュされた結果を無効にすることはありませんでした。その結果、流れが壊れてしまいました。 /admin に製品を追加し、すぐに /products に戻ったとします。そのリストがすぐに表示されるという保証はありません。
Fable 5 がこの「バグ」をどのように処理するかをテストするのが好きでした。理由は次のとおりです。
ほぼ確実に問題を見つけて解決してくれることはわかっていましたが、より興味深い点は次のとおりです。
…どのようなアプローチが取られるのか、また何をすべきかをどのように決定するのかについて。この問題を解決するには、いくつかの部分を接続し、レンダリング プロセスを大幅に変更しない方法で適切なアプローチを選択する必要があります。
ちなみに、IMO が単独で意味をなしたアーキテクチャ上の選択肢のリストは次のとおりです。
カタログ ページのサーバー レンダリング
リスト用の URL ベースのフィルターもありました
製品データの 60 秒キャッシュは合理的なパフォーマンスの選択であると主張することもできます。
変更後に router.refresh() も呼び出しました
バグはこれらの決定の間のギャップに生息していました。誰かが /products にアクセスし、/admin から製品を追加し、60 秒のウィンドウが終了する前に /products に戻った場合、ページには古い製品リストが表示されたままになる可能性があります。
Next.js アプリを持っていますが、問題は、/admin から製品を追加してから /products に戻ると、新しい項目がすぐに表示されないことです。
何が問題ですか?これを解決できますか?
Fable 5 ではバグがすぐに修正されました。
ファイルを読み取ったところ、主な問題は製品作成後のキャッシュの無効化であることが認識されました。フィルターの変更、クライアント側コンポーネントへのリファクタリング、キャッシュの完全な無効化など、明らかに間違ったことは行っていません。重要な部分は理解していました。つまり、製品は保存されており、/products には書き込み前からキャッシュされた製品リストがまだ表示されている可能性があるということです。
ここまでは順調ですね。しかし、また期待されているのは、

死ぬほど宣伝されたSOTAモデル。
さて、修正自体です。その思考プロセスから得られたことは次のとおりです。
中心的な問題は、製品の作成後にキャッシュを無効にする必要があることです。最もクリーンなアプローチは、取得に next: { tags: [”products”] } のタグを付けてから、POST ハンドラーで revalidateTag(”products”) を呼び出すことです。
次へ: { 再検証: 60、タグ: [”製品”] }
import { revalidateTag } から “next/cache”;
const product = addProduct(...);
return NextResponse.json({ 製品 }, { ステータス: 201 });
これにより、通常の閲覧のために読み取りリクエストがキャッシュされたままになり、書き込み後に製品データが無効になり、サーバーでレンダリングされた製品ページが保存されました。全体として、ソリューションと答えは悪くなく、フロンティア モデルに期待されるものでした。
オーバーエンジニアリングは微妙だった
この特定のアプリの場合、最も簡単な修正は次のようになることを期待していました。
POST が成功したら、製品リストを表示するルートを無効にします。これでほぼ完了です。
タグベースのソリューションはよりスケーラブルですが、この特定のアプリが本当に必要とするものに合わせて過剰に設計されています。製品データがさまざまな場所 (ホームページ、カテゴリ ページ、関連製品セクション、サイトマップ、管理ビュー、場合によっては複数のレイアウト) に表示される場合、これが求められます。その世界では、すべてのパスを覚えるよりも、製品のコンセプトを無効にする方がクリーンです。
しかし、これは非常にシンプルなアプリで、数か月以上 Next.js を学んだ人であれば誰でも作成できました。この状況では、タグ自体が間違っているわけではありません。問題が必要とするものよりも少し多くのアーキテクチャが導入されています。
私は、ソーシャルメディア上で同じようなことをコメントする人々のパターンを繰り返し見てきました。 SOTA モデルはオーバーエンジニアリングであるということです。モデルはアーキテクチャ パターンを選択しますが、このアプリにそのパターンが必要な理由はまだ十分に正当化されていません。
私にとって、

最良の答えは次のように言うでしょう。
この小さなアプリの場合は、revalidatePath(”/products”) を使用します。
商品データが複数のルートで再利用される場合は、タグと revalidateTag(”products”) を使用します。
重要な更新: 明確にしておきますが、モデルの修正は有効でした。私の批判は、それが間違っているということではなく、要件に一致する最小の修正を検討する前に、スケーラブルな抽象化に到達したということです。
単純な例ですが、繰り返されるパターン
これは上級ソフトウェア エンジニアにとって非常に単純な例であることはわかっています。この例だけを基にして、Fable 5 のような SOTA モデルがオーバーエンジニアリングであると過度に一般化するのは、単純に言えば、愚かなことです。
しかし、このような意見を抱いているのは私だけではありません。同じパターンに気づいた人をさらに何人か紹介します。
Anthropic 氏はまた、「Claude Opus 4.5 と Claude Opus 4.6 は、余分なファイルを作成したり、不必要な抽象化を追加したり、要求されていない柔軟性を組み込んだりして、オーバーエンジニアリングする傾向がある」とも述べています。それをプロンプトで調整することを推奨していますが、プロンプトだけで十分なのか疑問です。
ほとんどのモデルの評価が特定のパターンに従っていることも興味深いです。
タスク -> タスクを評価する -> テストと比較する
モデルが向上するにつれて、ベンチマーク タスクはより複雑になります。新しいモデルは、この「ループ」を意識せずに最適化しているのだろうか。より複雑なタスクにはより複雑な解決策が含まれるため、通常はより複雑な問題に対するパターンが必要になります。その結果、モデルが最初から必要のないパターンを試行するため、単純なアプリでは問題が発生します。
共有 この投稿に関するディスカッション コメント 再スタック トップ 最新のディスカッション 投稿はありません

## Original Extract

One thing I keep noticing with the latest coding models is that their mistakes are getting more… interesting.

SOTA Models Should Learn to KISS (Fable 5 Experiment)
Kilo Blog
Subscribe Sign in SOTA Models Should Learn to KISS (Fable 5 Experiment)
Darko Jun 11, 2026 Share One thing I keep noticing with the latest coding models is that their mistakes are getting more… interesting.
With older AI models, it was pretty easy to spot a mistake. The model hallucinated an API, misunderstood the framework, wrote code that did not compile, or confidently worked on the wrong part of the app. You could reject its suggestions quickly because the problem was pretty obvious to spot.
With newer models like Claude Fable 5, the failure modes are more subtle.
I kept seeing overhyped claims about how Fable 5 can one-shot 3D games, cool graphics , animations, etc. A lot of people use Claude models for coding though, so I decided to benchmark it by giving it, well, a coding task.
To get things started, I created basic a Next.js app where you could list/add products:
The app had 2 main routes: /products and /admin. It also had an internal API route at /api/products (I know this isn’t necessary for production apps most of the time, but I wanted to make things more complicated for the model).
The /admin page contained a form for adding products. To make things more complex (for the model, that is), I made the form post to the /api/products route handler, where the product was added to an in-memory data store, and that change needed to reflect whenever you returned to /products and expected to see the newly added item. The read side of the bug looked like this:
const res = await fetch(`${baseUrl}/api/products?${params}`, {
Just to clarify, the bug was not the cached fetch (nothing’s wrong with this code in isolation) but the fact that /products cached its read for 60 seconds, while the admin write path never invalidated that cached result. This resulted in a broken flow; say you added a product under /admin and immediately went back to /products, there was no guarantee you’d immediately see that listing.
I liked testing how Fable 5 would handle this “bug” because:
I knew it would almost certainly find it and solve it, but the more interesting part was:
…about the approach it would take and how it’ll decide what to do. To solve this problem, you need to connect several pieces together, pick a suitable approach in a way you don’t drastically alter the rendering process.
By the way, here’s a list of the architectural choices, which IMO made sense in isolation:
Server-rendering a catalog page
I also had URL-based filters for the listing
You could argue that a 60-second cache on product data is a reasonable performance choice
I also called router.refresh() after mutating
The bug lived in the gap between those decisions. If someone visited /products, then added a product from /admin, then returned to /products before the 60-second window expired, the page could still show the old product list.
I have a Next.js App and the problem is that when I add a product from /admin, then go back to /products, the new item sometimes doesn’t appear immediately.
What’s the issue and can you fix this?
Fable 5 did fix the bug pretty quickly.
It read the files and recognized that the core issue was cache invalidation after product creation. It did not do obviously wrong things like modify the filters, refactor to a client-side component, disable caching entirely, etc. It understood the important part: the product was being saved, and that /products could still be showing a cached product list from before the write.
So far, so good. But also expected for a SOTA model that was hyped to death.
Now, the fix itself. This is what we got from its thinking process:
The core issue is that I need to invalidate the cache after creating a product. The cleanest approach is to tag the fetch with next: { tags: [”products”] } and then call revalidateTag(”products”) in the POST handler.
next: { revalidate: 60, tags: [”products”] }
import { revalidateTag } from “next/cache”;
const product = addProduct(...);
return NextResponse.json({ product }, { status: 201 });
This kept the read requests cached for normal browsing, invalidated product data after writes, and preserved the server-rendered products page. Overall, the solution and the answer weren’t bad and it was something you’d expect from a frontier model.
The overengineering was subtle
For this specific app, I would have expected the simplest fix to be:
After the POST succeeds, invalidate the route that shows the product list…and that’s pretty much it.
A tag-based solution is more scalable, but it is also over-engineered for what this specific app really needs. It is what you reach for when product data appears in many places: a homepage, category pages, related-product sections, sitemaps, admin views, maybe multiple layouts. In that world, invalidating the concept of products is cleaner than remembering every path.
But this was a very simple app that anyone who learned Next.js for more than a few months could write. In this situation, tags are not wrong per se; they introduce a bit more architecture than the problem needs.
I’ve seen a recurring pattern of people commenting a similar thing across social media; that SOTA models are overengineering. The model chooses an architectural pattern, but does not quite justify why this app needs that pattern yet.
To me, the best answer would have said:
For this small app, use revalidatePath(”/products”).
If product data is reused across multiple routes, use tags and revalidateTag(”products”).
Important update: To be clear, the model’s fix was valid. My criticism is not that it’s wrong, but that it reached for a scalable abstraction before considering the smallest fix that matched the requirements.
A simple example, but a recurring pattern
I know that this is a pretty simple example for a senior software engineer. Over-generalizing that SOTA models like Fable 5 are overengineering based on this example alone would be…well, stupid, to put it plain and simple.
However, I’m not alone in this observation. Here are a few more people who noticed the same pattern.
Anthropic also said that “Claude Opus 4.5 and Claude Opus 4.6 have a tendency to overengineer by creating extra files, adding unnecessary abstractions, or building in flexibility that wasn’t requested.” They recommend adjusting that with prompting, but I wonder if prompting is enough.
It’s also interesting how most model evals follow a specific pattern:
Task -> Evaluate a task -> Compare it against tests
As models get better, those benchmark tasks get more complicated. I wonder if new models are optimizing for this “loop” without being aware of it; more complex tasks involve more complex solutions, you usually need patterns for more complicated problems; as a result, simpler apps suffer because the model is trying patterns when there’s no need for it in the first place.
Share Discussion about this post Comments Restacks Top Latest Discussions No posts
