---
source: "https://blog.kukiel.dev/posts/cfml-ai-images-with-cloudflare.html"
hn_url: "https://news.ycombinator.com/item?id=48546320"
title: "AI Image Generation from CFML in One Function Call"
article_title: "AI Image Generation from CFML in One Function Call"
author: "rmason"
captured_at: "2026-06-15T21:56:02Z"
capture_tool: "hn-digest"
hn_id: 48546320
score: 2
comments: 0
posted_at: "2026-06-15T20:03:09Z"
tags:
  - hacker-news
  - translated
---

# AI Image Generation from CFML in One Function Call

- HN: [48546320](https://news.ycombinator.com/item?id=48546320)
- Source: [blog.kukiel.dev](https://blog.kukiel.dev/posts/cfml-ai-images-with-cloudflare.html)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T20:03:09Z

## Translation

タイトル: 1 回の関数呼び出しで CFML から AI 画像を生成
説明: CFML から画像を生成したいと考えていました。重い Python サービスを呼び出したり、GPU を立ち上げたりする必要はありません。Lucee にプロンプ​​トを書くだけで、画像が返されます。そこで、Cloudflare の Workers AI イメージ モデルをラップする小さなライブラリを構築しました。全体は次のようになります。

記事本文:
1 つの関数呼び出しで CFML から AI 画像を生成 |ポールのストーリーライン
ポール@ホーム:~$
興味深い技術的な問題が解決され、説明されました
1 回の関数呼び出しで CFML から AI 画像を生成
CFMLから画像を生成したいと考えていました。重い Python サービスを呼び出したり、GPU を立ち上げたりする必要はありません。Lucee にプロンプ​​トを書くだけで、画像が返されます。そこで、Cloudflare の Workers AI イメージ モデルをラップする小さなライブラリを構築しました。全体は次のようになります。
cf = 新しいcloudflareimages.CloudflareImages(); // env からの認証情報
img = cf.generate (prompt = "夕暮れ時の岩だらけの海岸にある灯台、絵画風" );
img.toFile (expandPath ( "./lighthouse.jpg" ) );
この呼び出しにより、次のような結果が得られました。
私は公開デモを作成しました — プロンプトを入力し、生成を押し、画像を取得します。
cfml-image-with-cloudflare.kukiel.dev
1 日に数枚の画像に制限されているため無料のままですが、実際の AI モデルを呼び出す本物のライブ CFML アプリです。人々 (つまり私) がそれを使って作ったものをいくつか紹介します。
Cloudflareはテキストから画像への変換モデルを多数提供していますが、その答えについては意見が一致していません。 Stable Diffusion ファミリのほとんどは、生の画像バイトを渡します。 Flux は、base64 でエンコードされた画像を含む JSON を内部に渡します。 API を直接呼び出すと、2 つのコード パスと大量のコンテンツ タイプ スニッフィングを記述することになります。
図書館はそれを消してくれます。どのモデルを選択しても、同じ GenerationResult が返されます。
img.getBinary(); // 生のバイト
img.toBase64 (); // Base64 (データ: URI に便利)
img.toFile (パス); // ディスクに書き込みます
醜い部分はまさに 1 か所 (応答ノーマライザー) に存在するため、それについて考える必要はありません。 @cf/black-forest-labs/flux-1-schnell を @cf/stabilityai/stable-diffusion-xl-base-1.0 に置き換えてもコードは変わりません。
画像から画像への変換や入力も行います

使用可能なモデルをリストできます。
// 既存の画像を再考する
img = cf.imageToImage (prompt = "冬にしましょう" , image = ExpandPath ( "./Summer.png" ) );
// マスクされた領域に何かをペイントします
img = cf.inpaint (プロンプト = "帽子" 、画像 = 写真、マスク = マスクバイト );
// 何を使えばいいのでしょうか？
モデル = cf.listModels ();
何か問題が発生すると、型指定された例外がスローされるため、トークンの欠落 ( ConfigError )、Cloudflare のノー通知 (実際のメッセージとコードを含む APIError )、ネットワーク ブリップ ( TransportError ) など、失敗したものに実際に対応できます。このデモでは、安全フィルターが作動したときに生のスタック トレースの代わりに、これを使用して「そのプロンプトはブロックされました」というわかりやすいメッセージを表示します。
私が最も満足している点は、テストスイートが Cloudflare を呼び出さないことです。ネットワークに接続する 1 つの場所 (小さなトランスポート コンポーネント) は交換可能であるため、テストでは定型応答 (SDXL の場合は生のバイト、Flux の場合は Base64 JSON) を返す偽のものを挿入します。スイート全体がミリ秒単位でオフラインで実行され、費用はかからず、SDXL 対 Flux 正規化が実際に機能することが証明されています。本物の認証情報が存在する場合にのみ起動するライブスモークテストが 1 つあります。
フローチャート: コード → ファサード → (検証) → トランスポート (唯一の cfhttp)
→ Cloudflare → ノーマライザー → GenerationResult → あなたに戻る
各ピースは 1 つのジョブを備えた 1 つの小さな CFC です。これがドロップインの理由です。フォルダーをコピーし、Cloudflare アカウントを指定して、写真を要求します。
無料の Cloudflare アカウントと Workers AI API トークン (ダッシュボード → AI → Workers AI → REST API) が必要です。無料枠では、1 日あたり 10,000 個の「ニューロン」が無料で提供されます。これは、サイズに応じて数百枚の画像に十分な量です。 512×512 の画像は安価です。 1024×1024 では約 4 倍のコストがかかるため、デモは無料許容範囲を広げるために 512 で実行されます。

セ。
それはすべてオープンソースです - GitHub の Cloudflareimages です。クローンを作成し、cloudflareimages/ フォルダーをアプリにコピーすれば、準備は完了です。 README には、完全な API、ピースがどのように適合するかを示すマーメイド図、および TestBox スイートが含まれています。それを使用し、フォークし、PR を送信し、周囲に共有してください。CFML の世界ではこれをもっと活用できるでしょう。
デモを試してみたら、何ができたのか教えてください。

## Original Extract

I wanted to generate images from CFML. Not call out to some heavyweight Python service, not stand up a GPU — just write a prompt in Lucee and get a picture back. So I built a small library that wraps Cloudflare’s Workers AI image models, and the whole thing comes down to this:

AI Image Generation from CFML in One Function Call | Pauls Storyline
paul@home:~$
Interesting and Technical problems solved and explained
AI Image Generation from CFML in One Function Call
I wanted to generate images from CFML. Not call out to some heavyweight Python service, not stand up a GPU — just write a prompt in Lucee and get a picture back. So I built a small library that wraps Cloudflare’s Workers AI image models, and the whole thing comes down to this:
cf = new cloudflareimages.CloudflareImages (); // creds from env
img = cf.generate ( prompt = "a lighthouse on a rocky coast at sunset, painterly" );
img.toFile ( expandPath ( "./lighthouse.jpg" ) );
That call produced this:
I put a public demo up — type a prompt, hit generate, get an image:
cfml-image-with-cloudflare.kukiel.dev
It’s capped at a handful of images a day so it stays free, but it’s a real, live CFML app calling real AI models. Here are a few things people (well, me) have made with it:
Cloudflare gives you a bunch of text-to-image models, and they don’t agree on how to answer. Most of the Stable Diffusion family hand you raw image bytes . Flux hands you JSON with the image base64-encoded inside it . If you call the API directly you end up writing two code paths and a pile of content-type sniffing.
The library makes that disappear. Whatever model you pick, you get back the same GenerationResult :
img.getBinary (); // the raw bytes
img.toBase64 (); // base64 (handy for a data: URI)
img.toFile ( path ); // write it to disk
The ugly part lives in exactly one place — a response normalizer — so you never think about it. Swap @cf/black-forest-labs/flux-1-schnell for @cf/stabilityai/stable-diffusion-xl-base-1.0 and your code doesn’t change.
It also does image-to-image and inpainting, and it can list the available models:
// reimagine an existing picture
img = cf.imageToImage ( prompt = "make it winter" , image = expandPath ( "./summer.png" ) );
// paint something into a masked region
img = cf.inpaint ( prompt = "a hat" , image = photo , mask = maskBytes );
// what can I use?
models = cf.listModels ();
When something goes wrong it throws typed exceptions, so you can actually react to what failed — a missing token ( ConfigError ), Cloudflare saying no ( APIError , with its real message and code), a network blip ( TransportError ). The demo uses that to show a friendly “that prompt was blocked” message when the safety filter trips, instead of a raw stack trace.
The bit I’m happiest with: the test suite never calls Cloudflare. The one place that touches the network — a tiny Transport component — is swappable, so the tests inject a fake that returns canned responses (raw bytes for SDXL, base64 JSON for Flux). The whole suite runs offline, in milliseconds, costs nothing, and still proves the SDXL-vs-Flux normalization actually works. There’s one live smoke test that only wakes up if real credentials are present.
flowchart: your code → facade → (validate) → Transport (the only cfhttp)
→ Cloudflare → normalizer → GenerationResult → back to you
Each piece is one small CFC with one job. That’s what makes it a drop-in: copy the folder, point it at a Cloudflare account, ask for a picture.
You need a free Cloudflare account and a Workers AI API token (dashboard → AI → Workers AI → REST API). The free tier gives you 10,000 “Neurons” a day at no charge — enough for a few hundred images depending on size. A 512×512 image is cheap; 1024×1024 costs about 4× more, so the demo runs at 512 to stretch the free allowance.
It’s all open source — cloudflareimages on GitHub . Clone it, copy the cloudflareimages/ folder into your app, and you’re away. The README has the full API, a Mermaid diagram of how the pieces fit, and the TestBox suite. Use it, fork it, send a PR, and please share it around — the CFML world could use more of this.
If you give the demo a spin, let me know what you make.
