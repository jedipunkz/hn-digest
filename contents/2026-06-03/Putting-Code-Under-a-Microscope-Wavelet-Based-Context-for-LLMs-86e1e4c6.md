---
source: "https://yogthos.net/posts/2026-06-02-wavescope.html"
hn_url: "https://news.ycombinator.com/item?id=48373249"
title: "Putting Code Under a Microscope: Wavelet-Based Context for LLMs"
article_title: "(iterate think thoughts): Putting Code Under a Microscope: Wavelet-Based Context for LLMs"
author: "yogthos"
captured_at: "2026-06-03T00:39:42Z"
capture_tool: "hn-digest"
hn_id: 48373249
score: 3
comments: 0
posted_at: "2026-06-02T17:24:58Z"
tags:
  - hacker-news
  - translated
---

# Putting Code Under a Microscope: Wavelet-Based Context for LLMs

- HN: [48373249](https://news.ycombinator.com/item?id=48373249)
- Source: [yogthos.net](https://yogthos.net/posts/2026-06-02-wavescope.html)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T17:24:58Z

## Translation

タイトル: コードを顕微鏡で見る: LLM のウェーブレットベースのコンテキスト
記事のタイトル: (反復思考): コードを顕微鏡で見る: LLM のウェーブレット ベースのコンテキスト
説明: プログラミング、Clojure、ソフトウェア開発に関する Dmitri のブログ

記事本文:
(思考の反復): コードを顕微鏡で見る: LLM のウェーブレット ベースのコンテキスト
(反復
考える
思い）
テーマ
2026 年 6 月 2 日
コードを詳しく見る: LLM のウェーブレットベースのコンテキスト
AI コーディング ツールを試したことのあるすべての開発者は、モデルがコードベースを手探りして編集する関連セクションを見つけるのを観察するという問題に精通しています。大規模なプロジェクトのコンテキストにコードベース全体を読み込むことはできないため、いくつかのファイルを grep してコンテキストを与え、次に何をすべきかを推測します。ただし、コードはレイヤーと境界による階層構造を持っています。関数はクラス内にあります。クラスはファイル内に存在します。ファイルはモジュールを構成します。 1 つの 400 行ファイルには、それぞれ独自の目的を持つ 6 つの異なる概念領域を含めることができます。
人間の開発者がコードを読んで理解しようとするとき、その構造を活用します。ファイルとクラスがどのように編成されているかを調査し、それに基づいて関連するロジックを見つけようとします。あなたと同じように、エージェントがコードをズームインおよびズームアウトして全体像を確認し、ファイル全体を開かなくても必要な機能に直接ジャンプできたら素晴らしいと思いませんか。
それがWaveScopeの機能です。これは、ウェーブレット変換を使用して LLM にコードベースの多重解像度ビューを提供する MCP サーバーです。プログレッシブ画像読み込みのようなものを想像してみてください。ただし、ソース ファイルが対象です。新しいモデルは紙の上では大規模なコンテキストを処理できますが、実際には焦点が合わなくなり始めます。エージェントが持つコンテキストが増えるほど、実際に何に取り組む必要があるのか​​、何を優先すべきなのかを把握することが難しくなり、コンテキストの腐敗につながります。
現在、この問題に対処するには主に 2 つの方法があります。 Grep ベースの検索では完全一致が見つかりますが、パターンが一致するため構造は無視されます。

個々の行の ng では、クラス境界がどこにあるのか、またはエラー処理領域がどこから始まるのかがわかりません。埋め込みベースの RAG は、意味論的な意味は理解しますが、位置と構造を失うもう 1 つのアプローチです。どちらもモデルにコードのアーキテクチャを実際に認識させることはできません。
Clojure、TypeScript、Rust、Go などの言語でファイルを開くと、コード全体で共通の構造が繰り返されていることがわかります。輸入品がトップに位置している。クラスと関数の定義は定期的にポップアップ表示されます。インデントにも、すべての言語に独自の明確なパターンがあります。コメント ブロックと空白行は間に一時停止を挟みます。実際に言語の構文を知らなくても、これらのパターンを抽出して AST に似た構造を作成する方法があったとしたらどうでしょうか。
私たちにとって幸運なことに、ウェーブレットは、まさにこの種の信号を一度に複数のスケールに分解することで処理するために作成されました。変換により、大規模なパターンと同時にすべての細かいディテールが得られます。このアルゴリズム ファミリは非常に多用途であり、さまざまな分野で使用されています。地震学者はこれらを使用して地震を発見し、医師はこれを使用して MRI スキャンを鮮明にし、オーディオ エンジニアはベースラインをボーカルから分離します。コード構造も、同様の方法で分解できる別の種類の信号であることがわかりました。
では、それは実際にどのように機能するのでしょうか?各行にスコアが付けられると、ファイルは、構造の密度に応じて増減する一連の数値を表す 1D 信号として扱うことができます。リッカー ウェーブレットは、両側にくぼみのあるバンプのような形をした小さなテンプレートです。信号上で一度に 1 つの位置をスライドさせ、すべての位置で、その下の信号がその形状とどの程度一致しているかを測定します。強い一致は、間に高い領域があることを意味します。

構造的な境界を示唆する 2 つの静かな領域。出力は、境界のように見える度合いをスコア化した各ラインの係数です。
さまざまな解像度をエンコードするコツは、テンプレートの幅にあります。同じ図形をさまざまな幅に引き伸ばしてスライドさせることができ、それぞれが異なるスケールを表します。狭いウェーブレットは、隆起領域の幅が 1 ～ 2 ラインの場合にのみ一致するため、小さくて鋭いフィーチャで起動されます。広いものはラインレベルのノイズを無視し、周囲よりも高い大きな領域に反応し、大きな構造物を攻撃します。したがって、次のトリックは、一度に複数の幅を実行して、1 つだけではなく、幅のバンド全体で点灯する境界を確認することです。さまざまなスケールにわたって一貫したフィーチャは、私たちが重視する真の構造エッジである傾向があります。
コード構造は異なるサイズでネストされているため、各サイズはそれに適合する幅で表示されます。 WaveScope を独自の src/context.ts 上で実行することで、これがどのように機能するかの具体例を確認できます。単独の import ステートメントや、export class FileContext { 宣言などの 1 行は、静かな隣の行よりも目立ち、細かく細いウェーブレットがロックする鋭い 1 行のスパイクになっています。その係数はスケール 1 または 2 でピークに達し、より広いスケールになるほど減少します。もう一方の極端な例は、 inferLabel メソッド内のキーワードが密な長いカスケードです。これには、連続する if (tokens.includes("class")) 、interface 、 enum 、および struct 分岐が幅広く続きます。その領域全体は 1 つの広く隆起した台地として読み取られ、ウェーブレットが広がるにつれてその係数は着実に上昇し、スケール 16 で約 0.5、スケール 32 で 1.0、スケール 64 で 1.3、スケール 128 で約 2.3 となり、ファイル全体で最も強い構造応答としてピークに達します。最大の構造は最も強力な粗スケール信号を生成し、

これは、ズームアウトしたときに最初に表示したいスポットです。
これらすべてを生み出すラインスコアは同じパスから生まれています。この例では、クラス キーワードの重みが 1.0 でエクスポートが 0.6 であるため、export class FileContext { のスコアは 1.6、1 行の get lineCount() アクセサーのスコアは約 0.58、その下の読み取り専用フィールド宣言のスコアはそれぞれ約 0.08、クラスの周囲のコメントと空行のスコアは一律 0.0 です。クラス宣言はそれ自身の本体の上にそびえ立ち、本体はその周囲の空白の上にそびえ立ち、ウェーブレットはあらゆる幅でそれらの相対的な高さを読み取ります。
上記の直感は、パイプラインの前半部分で、各ラインを 1D 信号にスコア付けし、その後、1、2、4、8、16、32、64、128 ラインの 8 つのスケールでリッカー ウェーブレットをスライドさせ、各ラインと幅で係数を与えます。次に、生の係数配列をモデルで使用できるものに変換するには、さらにいくつかの手順が必要です。
1 つ目は、アレイの極大値をスキャンし、大きさによってランク付けするマルチスケール ピーク検出を実行することです。最も強い境界は、クラスの始まりやインポートとコード間の移行などの機能を表します。上で見たように、実際の境界は複数の隣接するスケールで表示されるため、同じスポットの重複でランキングが溢れかえるのを避けるために、これらの繰り返しを安全に 1 つのピークに折りたたむことができます。
2 番目のステップはバンドの組み立てで、ピークが 3 つの広範なズーム バンドに分割されます。スケール 1 ～ 2 の細かいバンドは、クエリ センターの周囲の閉じたウィンドウ内の生のソース行を示しています。スケール 4 ～ 16 の中程度のバンドは、関数とクラスのシグネチャを、その周囲のコンテキストとともに追跡します。最後に、スケール 32 ～ 128 の粗いバンドは、半径全体をセクション レベルのサマリーに圧縮します。
その処理はすべてハンドルです

d は MCP サーバーによって生成され、モデルはウェーブレットの計算を気にすることなく、バンドとピーク位置を含む構造化された JSON を単純に参照します。
エージェントが、認証ロジックを含む 500 行の TypeScript ファイル内の 150 行目にある query_wavelet_context を呼び出すとします。この場合、細かいバンドには検査される実際のコード行が含まれます。中間のバンドには、上部のインポートや下部のテスト ヘルパーなどのピークによってガイドされる、行 0 ～ 400 の概要が表示されます。
このモデルは、細かい帯域に注意を払うことで、updateUser が何を行うかについての情報を取得しますが、粗い帯域から認証コンテキストについても認識します。ファイルの 500 行すべてのテキストを確認しなくても、ウェーブレット ピーク内のクラスと関数の境界を認識することで、関連するコードにジャンプできます。
プロジェクト全体で動作する get_ important_positions というユーティリティもあります。すべてのソース ファイルを調べて、ウェーブレットのピークを平滑化し、コード内の最も重要な場所のランク付けされたリストを提供します。
構造境界を特定するだけでなく、サーバーは、前回の投稿で説明したハール離散ウェーブレット変換とビットコスト推定によって駆動される一対のエントロピー分析ツールを使用して、それらの構造がどの程度複雑または不規則であるかを測定することもできます。リッカー係数は、get_entropy_bands を使用して各スケールで量子化できます。これは、その解像度での構造的不規則性が高いことを示す高いコストでビットプレーン カウントを計算します。元のラインごとの構造信号は、get_complexity_heatmap を使用して複数の Haar レベルを通じて分解し、エントロピー コストをラインごとの不規則性スコアに投影し直すこともできます。モデルはこれらのスコアを一種のテクスチャ チャネルとして使用して、コードの危険な部分がどこに存在するかを理解できます。アン

y 定型領域は最終的に低いスコアになるため、安全に要約またはスキップできますが、高エントロピー領域には、特別な注意が必要な高密度のロジックまたは異常なパターンが含まれる可能性があります。これらのツールは、コードを高レベルでトリアージするためのデータ駆動型の方法をモデルに提供し、複雑なロジックを持つセクションをモデルが簡単に見つけて分割できるリファクタリング タスクに最適です。
ウェーブレットの主な利点は、正規表現などのツールを使用してのみ推論できるコード全体の構造に焦点を当てていることです。より正確な AST 解析を使用すると、同じ種類の分析を多数取得できることに言及する価値があります。ただし、AST ツールでは言語ごとにパーサーが必要ですが、ウェーブレットは適用されるテキストのセマンティクスをまったく意識しません。これらは単に統計的規則性にフラグを立てるだけであり、そもそもそれが非常に広範囲に適用可能なツールである理由です。
Wavescope のアプローチは、変換があらゆる 1D 信号に対して機能するため、grepping と本格的な AST 解析の間の適切な中間点に位置します。言語認識は、各言語のコア セマンティクスを記述するための 10 行の設定だけを必要とする完全なパーサーではなく、単純なキーワード重み付けレイヤーから得られます。また、全体の実行コストが低く、ファイル全体を数ミリ秒で処理して、階層的でマルチスケールの出力を作成できます。スケールは、LLM にナビゲートするためのマップを提供するコード構造の自然な表現です。
ウェーブレットはすべてのレベルで同時にエッジ位置を特定するため、高度なパーサーでも検出するのが困難な構造変化を簡単に特定します。たとえば、if / else ブロックの長いシーケンスは、多くの短いメソッドを持つクラスとは構造的に異なって見えますが、

文書の領域は山間の谷のように見えます。ウェーブレットは、構造的に異なることを認識するために、これらが実際に何であるかを知る必要はありません。そしてこれは、モデルが何をすべきかを理解するために必要なものであることがよくあります。
概念を説明するために、5,000 行弱の TypeScript を含む WaveScope 独自の 14 ファイル コードベースに対して 3 つの現実的な開発タスクを実行し、トークン コストを従来の方法と比較しました。ここでの「従来の」とは、LLM コーディング エージェントが実際に行うことを意味します。つまり、ランドマークを grep し、ターゲットのコード チャンクを読み取り、ファイル ヘッダーを読み取って構造がどのようなものであるかのサンプルを取得します。
一般的なタスクの 1 つは、大きなファイルの構造を理解することです。そこで、エージェントに 854 行の Index.ts を分析してもらい、インポートがどこで終了するかを確認し、ツール登録がクラスター化されている場所を見つけて、起動コードを特定しました。これらは典型的なタスクであり、アプローチとしては、エクスポート ランドマークとセクション コメントを grep してから、インポート ブロックを読み取り、登録例を検索し、起動末尾を認識します。単なるヒューリスティックである斑点のある画像を生成するには、約 2,000 トークンのコストがかかります。 WaveScope の粗いバンドと上位 15 の重要な位置から、sa が得られます。

[切り捨てられた]

## Original Extract

Dmitri's blog about programming, Clojure, and software development

(iterate think thoughts): Putting Code Under a Microscope: Wavelet-Based Context for LLMs
( iterate
think
thoughts )
Theme
June 2, 2026
Putting Code Under a Microscope: Wavelet-Based Context for LLMs
Every developer who has tried an AI coding tool is familiar with the problem of watching the model fumble with the codebase to find relevant sections to edit. Since it's not possible to load an entire codebase into the context for large projects, it greps through a few files to give it some context, and guesses what to do next. But code has a hierarchical structure with layers and boundaries. Functions sit inside classes. Classes live in files. Files make up modules. One 400-line file can contain six different conceptual areas, each with its own distinct purpose.
When a human developer reads the code, we leverage the structure when we try to understand it. We examine the way files and classes are organized, and try to find relevant logic based on that. Wouldn't it be nice if the agent could, like you, zoom in and out of the code so that it could look at the big picture, then jump directly to the exact function it needs without having to open the entire file.
That’s what WaveScope does. It’s an MCP server that uses wavelet transforms to give LLMs a multi-resolution view of the codebase. Imagine something like progressive image loading, but for source files. Even though newer models can handle large contexts on paper, what happens in practice is that they start losing focus. The more context the agent has the harder it is to figure out what it actually needs to work on, and what to prioritize leading to context rot.
Currently, there are two main ways of dealing with the problem. Grep-based search finds exact matches but ignores structure since pattern matching on individual lines can’t tell you where a class boundary is or where an error handling region starts. Embedding-based RAG is another approach which understands semantic meaning but loses position and structure. Neither gives the model a real sense of the architecture of the code.
Open a file in any language like Clojure, TypeScript, Rust, or Go and you’ll see repeating common structures throughout the code. Imports sit at the top. Class and function definitions pop up at regular intervals. Indentation also has its own distinct pattern in every language. Comment blocks and blank lines are pauses in between. What if there was some way to extract these patterns, and create a structure similar to an AST without actually having to know the syntax of the language.
Luckily for us, wavelets were made for processing exactly this sort of signal by decomposing it into multiple scales at once. The transform gives you all the fine details along with the large-scale patterns at the same time. This family of algorithms is very versatile and has been used in many different fields. Seismologists use them to spot earthquakes, doctors sharpen MRI scans with them, and audio engineers separate basslines from vocals. Code structure, it turns out, is just another kind of signal that can be decomposed in a similar fashion.
So how does that actually work? Once each line has a score, the file can be treated as a 1D signal representing a sequence of numbers that rises and falls with the density of the structure. The Ricker wavelet is a little template shaped like a bump with a dip on each side. You slide it across the signal one position at a time and, at every position, measure how well the signal underneath matches that shape. A strong match means there's an elevated region sitting between two quieter regions which suggests a structural boundary. The output is a coefficient at every line scored on how much it looks like a boundary.
The trick for encoding different resolutions lies in the width of the template. You can slide the same shape stretched to many widths, each representing a different scale. A narrow wavelet only matches when the elevated region is a line or two wide, so it fires on small, sharp features. A wide one ignores line-level noise and responds to larger regions elevated relative to their surroundings, firing for big structures. So, the next trick is to run multiple widths at once to see boundaries which light up across a band of widths rather than just one. Features consistent across different scales tend to be genuine structural edges that we care about.
Since code structure nests at different sizes, each size shows up at the width that fits it. We can see a concrete example of how this works by running WaveScope on its own src/context.ts . A single line such as a lone import statement, or export class FileContext { declaration stands out over its quieter neighbors making it a sharp one-line spike that the fine, narrow wavelets lock onto. Its coefficient peaks at scale 1 or 2 and fades away at wider scales. At the other extreme is the long keyword-dense cascade inside the inferLabel method, which has a wide run of consecutive if (tokens.includes("class")) , interface , enum , and struct branches. That whole region reads as one broad elevated plateau, and its coefficient climbs steadily as the wavelet widens to roughly 0.5 at scale 16, 1.0 at scale 32, 1.3 at scale 64, and about 2.3 at scale 128, where it peaks as the strongest structural response in the entire file. The biggest structure produces the strongest coarse-scale signal, which is exactly the spot you want surfaced first when you zoom out.
The line scores that feed all of this come from the same pass. In our example, export class FileContext { scores 1.6 because the class keyword weight of 1.0 and export at 0.6, the one-line get lineCount() accessor scores about 0.58, the readonly field declarations beneath it score roughly 0.08 each, while comments and blank lines around the class score a flat 0.0. The class declaration towers over its own body, the body towers over the whitespace around it, and the wavelet reads those relative heights at every width.
The intuition above is the front half of the pipeline where you score every line into a 1D signal, then slide the Ricker wavelet across it at eight scales which are 1, 2, 4, 8, 16, 32, 64, and 128 lines giving us a coefficient at every line and width. Next, we need a couple more steps to turn raw coefficient arrays into something the model can use.
First is to do multi-scale peak detection where the arrays are scanned for local maxima and ranked by magnitude. Strongest boundaries represent features such as the beginning of a class or the transition between imports and code. Because a real boundary shows up at several adjacent scales, as we saw above, these repeats can be safely collapsed into a single peak to avoid flooding the ranking with duplicates of the same spot.
The second step is the band assembly, where the peaks are separated into three broad zoom bands. The fine band at scales 1–2 shows raw source lines in a close window around the query center. The medium band at scales 4–16 tracks function and class signatures with some context around them. Finally, the coarse band at scales 32–128 compresses the whole radius into a section-level summary.
All of that processing is handled by the MCP server, and the model simply sees structured JSON with bands and peak positions without having to worry about any of the wavelet math.
Let's say an agent calls query_wavelet_context which is centered on line 150 inside a 500 line TypeScript file that has some authentication logic. In this case, the fine band will have the actual lines of code being inspected. The medium band will provide a summary of lines 0–400, guided by peaks such as the imports at top and test helpers at bottom.
The model derives its knowledge of what updateUser does by paying attention to the fine band, but it also knows about authentication context from the coarse band. It's able to jump to related code by recognising class and function boundaries in the wavelet peaks without needing to see all 500 lines of text from the file.
There is also a utility called get_important_positions , which operates on the whole project. It goes through every source file, smooths out the wavelet peaks, and gives you a ranked list of the most important places in the code.
Beyond locating structural boundaries, the server can also measure how complex or irregular those structures are, using a pair of entropy analysis tools driven by a Haar discrete wavelet transform and bit-cost estimation I discussed in my last post . Ricker coefficients can be quantized at each scale using get_entropy_bands which computes their bit-plane counts with higher cost indicating more structural irregularity at that resolution. The original per-line structural signal can also be decomposed through multiple Haar levels using get_complexity_heatmap to project the entropy cost back onto a per-line irregularity score. The model can use these scores as a sort of texture channel to understand where gnarly parts of the code live. Any boilerplate regions will end up with a low score, and so they can be safely summarized or skipped, while high-entropy regions will likely contain dense logic or unusual patterns that warrant extra attention. These tools give the model a data-driven way to triage code at high level and works great for refactoring tasks where the model can easily find sections with tangled logic and break it up.
The key advantage of wavelets is that they focus on the overall structure of the code which is something that can only be inferred using tools like regex. It's worth mentioning that it's possible to get a lot of the same type of analysis using AST parsing which is even more exact. However, AST tools necessitate having a parser for each language while wavelets are completely agnostic of the semantics of the text they're applied to. They simply flag statistical regularities, and that's what makes them such a broadly applicable tool in the first place.
Wavescope’s approach strikes a happy medium between grepping and full blown AST analysis since the transform works on any 1D signal. Language awareness comes from a simple keyword-weighting layer rather than a full parser with each language just needing a 10-line config to describe its core semantics. And the whole thing is cheap to run, able to process an entire file in milliseconds to create hierarchical and multi-scale outputs. The scales happen to be a natural representation of code structure which provides the LLM with a map to navigate it.
Since the wavelet locates edge positions at all levels simultaneously, it trivially locates structural changes that even sophisticated parsers might struggle to detect. For example, a long sequence of if / else blocks looks structurally different from a class with many short methods while regions of documentation appear as valleys between peaks. The wavelet doesn't need to know what these things actually are to perceive that they are structurally different. And this often happens to be just what the model needs to figure out what to do.
To illustrate the concept, I ran three realistic development tasks against WaveScope's own 14 file codebase containing just under five thousand lines of TypeScript to compare the token cost with the traditional way. "Traditional" here means what LLM coding agents actually do: grep for landmarks, read targeted chunks of code, and skim file headers to get a sample of what the structure looks like.
One common task is to understand the structure of a large file. So, I had the agent analyze index.ts weighing in at 854 lines to see where imports end, to find where the tool registrations cluster, and identify the startup code. These are your typical tasks where the approach would be to grep for export landmarks and section comments, then read the import block, find a registration example, and recognize the startup tail. That costs about 2,000 tokens producing a patchy picture which is just a heuristic. WaveScope's coarse band plus the top 15 important positions give the sa

[truncated]
