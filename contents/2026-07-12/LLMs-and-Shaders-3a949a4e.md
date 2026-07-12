---
source: "https://amitp.blogspot.com/2026/07/llms-and-shaders.html"
hn_url: "https://news.ycombinator.com/item?id=48885572"
title: "LLMs and Shaders"
article_title: "Amit's Thoughts: LLMs and shaders"
author: "bobbiechen"
captured_at: "2026-07-12T22:38:18Z"
capture_tool: "hn-digest"
hn_id: 48885572
score: 1
comments: 0
posted_at: "2026-07-12T22:35:02Z"
tags:
  - hacker-news
  - translated
---

# LLMs and Shaders

- HN: [48885572](https://news.ycombinator.com/item?id=48885572)
- Source: [amitp.blogspot.com](https://amitp.blogspot.com/2026/07/llms-and-shaders.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T22:35:02Z

## Translation

タイトル: LLM とシェーダー
記事のタイトル: Amit の考え: LLM とシェーダー
説明: 前回の投稿では、LLM へのチャット インターフェイスが LLM の機能について限定的なビューを提供する方法について説明しました。この投稿で私が入りたかったのは...

記事本文:
アミットの考え: LLM とシェーダー
LLM とシェーダー
アミットの想いから
▼
2026年
(1)
▼
7月
(1)
LLM とシェーダー
前回の投稿では、LLM へのチャット インターフェイスが、LLM ができることについて限定的なビューを提供していることについて説明しました。この記事では、これをグラフィックス パイプラインと比較したいと思いました。
グラフィックス パイプラインは以前は「固定機能」であり、グラフィックス システム (シリコン グラフィックス ハードウェアなど) がグラフィックス機能を提供していました。
常に行列の乗算 (移動、拡大縮小、回転、遠近法など) が得られます。
あなたはいつもグーローシェーディングを持っています
常に標準の照明モデル (アンビエント、拡散、鏡面、発光、光沢) を利用できます。
glLight() を使用して、いくつかのスタイルを使用して最大 8 つのライトを選択する必要があります
常にオブジェクトごとにカラーを取得していましたが、後にテクスチャ ルックアップが提供されました
glFog() で霧のレベルを制御する必要があります (濃度、色など)。
しかし、フォン シェーディングのような別のモデルが必要な場合はどうすればよいでしょうか?残念な！霧以外の効果が必要な場合はどうすればよいでしょうか?残念な！ 8 つ以上のライトが必要な場合はどうすればよいでしょうか?残念な！ 2 つのテクスチャが必要な場合はどうすればよいでしょうか?残念な！少なくとも、2 番目のテクスチャ ルックアップを可能にする機能が追加されるまでは。それ以外のものが必要な場合は、独自のグラフィックス パイプラインを実装し、ハードウェア アクセラレーションを諦める必要がありました。
時間が経つにつれて、これらの固定機能は「シェーダー」に置き換えられました。頂点シェーダーは、マトリックスの乗算、クリッピング、フォグ、カラーを使用して全体の構造を処理します。それを制御できるようになると、人々は布、草、水、地形、アニメーション、手続き型ジオメトリ、その他多くのものを思いつきました。フラグメント シェーダは、テクスチャ ルックアップとライティングを使用して細かいディテールを処理します。それを制御できるようになると、人々はカスタムの照明、有機的なテクスチャ、アニメーション、プロシージャル アート、その他多くのものを思いつきました。
今日の雲場

sed LLM というと、固定関数グラフィックス パイプラインを思い出します。彼らが通過するステップの固定パイプラインがあります。
プロンプトの前に「システム プロンプト」を追加します
トランスクリプトを LLM にフィードします
LLM に、次に考えられるすべての単語 (トークン) について予測させます。
次に来るトークンを選択します (例: 最も高い確率)
そのトークンをトランスクリプトに追加し、最初のステップに戻ります（トランスクリプト全体を再度 LLM にフィードします）
ここには少し柔軟性があります。たとえば、「ツールの使用」では、次のトークンをトランスクリプトに追加する代わりに、次のトークンが特別なマーカーである場合、外部プログラムを実行してその出力をトランスクリプトに挿入する必要があると指示します。これは、グラフィックスにおけるシェーダーの動作に似ています。常に特定の方法で処理するのではなく、パイプライン全体を再利用できるようにしたいのですが、その一部を置き換えます。しかし、私たちにはまだ多くのことはできません。おそらく、考えられるトークンを 3 つ選択し、これを 3 つのトランスクリプトに分割して、それらがどのように進行するかを確認したいと思います。もしかしたらそれらは収束するかもしれない。たぶん、それらは分岐します。既存のトランスクリプトに戻って編集し、もう一度実行したいと思うかもしれません。おそらく、トランスクリプト内で人間とコンピュータの役割を入れ替えたいと考えています。今のところ、そういったことはできません。トランスクリプトはグラフであり、「思考の木」を探索したり、単一のパスだけを確認するのではなく、グラフ全体を分析したりしたいと考えています。
彼らが提供するもう 1 つの柔軟性は、「構造化された出力」を使用してトークンが選択される方法を制限することです。しかし、私たちにはまだ多くのことはできません。時々、珍しい単語を選択したり (これがどのように機能するかについてのベイジー伯爵の視覚的な説明を参照)、文字 e を含まない単語のみを選択したり (アーネスト ライトの物語「ギャズビー」は文字 'e' なしで書かれたように)、または を表す単語のみを選択したりすることもできます。

数字を送信するか、キキまたはブーバの言葉だけを選ぶか、韻を踏む言葉を選びます。ローカル モデルでは「ロジッツ プロセッサ」を使用できますが、クラウド ベースの LLM には同等のものはないようです。
現時点では、標準機能以外のことを実行したい場合は、独自のパイプラインを作成し、ローカル モデルを使用する必要があります。私はクラウドベースの LLM が提供するハードウェアと利便性をあきらめています。
そこで私は、LLM 以外のグラフィックス シェーダに相当するものを進化させることができるかどうか疑問に思っています。大まかに言うと、トランスクリプト全体の構造 (単一の追加専用パスではなく可変グラフ/ツリー) をより詳細に制御し、個々のトークンがどのように選択されるか (完全にプログラム可能) をより詳細に制御したいと考えています。これらは、頂点シェーダーとフラグメント シェーダーに相当する場合があります。今日のクラウド LLM は、グラフィックス API が固定ライティングとフォグのみをサポートしていた頃と似ていると思います。プロバイダーは、私たちがやりたいことをすべて想像して、それらの特定のことだけをプログラムしようとします。将来的には、コードをプラグインしてプロセスをカスタマイズできるようになるかもしれません。
グラフィックス パイプラインにこの柔軟性を追加したことの副作用として、GPU が科学コンピューティング、暗号通貨、ニューラル ネットワークなど、3D グラフィックス以外の多くの用途に使用されるようになりました。 LLM がチャットのトランスクリプトに限定されないとしたら、何に使用されるでしょうか?
ラベル:
推測、
構造

## Original Extract

In the last post I described how chat interfaces to LLMs are giving a limited view of what the LLMs could do. In this post I wanted to com...

Amit's Thoughts: LLMs and shaders
LLMs and shaders
From Amit’s Thoughts
▼
2026
(1)
▼
July
(1)
LLMs and shaders
In the last post I described how chat interfaces to LLMs are giving a limited view of what the LLMs could do. In this post I wanted to compare this to graphics pipelines.
Graphics pipelines used to be "fixed function" , where the graphics system (such as Silicon Graphics hardware) would offer you the graphics features:
you always got a matrix multiply (translate, scale, rotate, perspective, etc.)
you always got gouraud shading
you always got their standard lighting model (ambient, diffuse, specular, emission, shininess)
you got to choose up to 8 lights with glLight() with a handful of styles
you always got a color per object, but later they offered a texture lookup
you got to control the fog level with glFog() (density, color, etc.)
But what if you wanted a different model like phong shading? Too bad! What if you wanted any effect other than fog? Too bad! What if you wanted more than 8 lights? Too bad! What if you wanted 2 textures? Too bad! At least until they added a feature to enable a second texture lookup. If you wanted anything else, you needed to implement your own graphics pipeline and give up on hardware acceleration.
Over time, these fixed features were replaced by " shaders ". The vertex shader would handle the overall structure with matrix multiply, clipping, fog, colors. Once we could control it, people came up with cloth, grass, water, terrain, animation, procedural geometry, and lots of other things. The fragment shader would handle fine details with texture lookup and lighting. Once we could control it, people came up with custom lighting, organic textures, animation, procedural art, and lots of other things.
Today's cloud-based LLMs remind me of the fixed function graphics pipeline. There's a fixed pipeline of steps they go through:
prepend a "system prompt" before your prompt
feed the transcript to the LLM
have the LLM make predictions about all the possible next words (tokens)
choose which token comes next (e.g. the highest probability)
append that token to the transcript, and go back to first step (feeding the entire transcript again to the LLM)
There's a bit of flexibility here. For example instead of appending the next token to the transcript, "tool use" says if the next token is a special marker, then we should run an outside program and insert its output into the transcript. That's kind of like what a shader does in graphics — instead of always doing things a certain way, we want to be able to reuse the pipeline as a whole, but replace one piece of it. But we still can't do much. Maybe I want to pick three possible tokens, and then split this into three transcripts, and see how they progress. Maybe they converge. Maybe they diverge. Maybe I want to go back an edit the existing transcript and let it run through again. Maybe I want to switch the roles of the human and computer in the transcript. Right now I can't do those things. The transcript is a graph and I would like to explore a "tree of thoughts" or analyze the whole graph instead of only seeing a single path through it.
Another bit of flexibility they offer is to constrain the how the tokens are chosen with " structured outputs ". But we still can't do much. We might want to pick an unusual word occasionally (see Count Bayesie's visual explanation of how this can work ), or only pick words that don't have the letter e (like Ernest Wright's story "Gadsby" that was written without the letter 'e'), or only pick words that represent a number, or only pick kiki or bouba words , or pick words that rhyme. With local models we can use a "logits processor", but there seems to be no equivalent for cloud based LLMs.
Right now, if I want to do anything other than the standard functions, I have to write my own pipeline and use local models. I give up on the hardware and convenience that the cloud-based LLMs offer.
So I've been wondering if we'll evolve the equivalent of graphics shaders but for LLMs. Broadly speaking, I think I'd like to have more control over how the overall transcript structure (mutable graph/tree instead of a single append-only path), and more control over how the individual tokens are chosen (fully programmable). These might be the equivalent of vertex and fragment shaders. I think today's cloud LLMs are similar to when graphics APIs only supported fixed lighting and fog. The provider is trying to imagine all the things we might want to do, and then program only those specific things. The future may allow people to plug in code to customize the process.
A side effect of adding this flexibility to the graphics pipeline is that GPUs got used for lots of things other than 3D graphics, such as scientific computing, cryptocurrency, and neural networks. What would LLMs be used for if they weren't limited to chat transcripts?
Labels:
speculation ,
structure
