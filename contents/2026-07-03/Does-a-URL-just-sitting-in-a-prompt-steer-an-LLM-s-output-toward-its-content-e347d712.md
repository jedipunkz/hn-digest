---
source: "https://aifoc.us/influencing-model-output-with-urls/"
hn_url: "https://news.ycombinator.com/item?id=48779535"
title: "Does a URL just sitting in a prompt steer an LLM's output toward its content?"
article_title: "does a url in a prompt steer an llm's output toward its content? | AI Focus"
author: "kinlan"
captured_at: "2026-07-03T20:53:30Z"
capture_tool: "hn-digest"
hn_id: 48779535
score: 2
comments: 0
posted_at: "2026-07-03T20:23:02Z"
tags:
  - hacker-news
  - translated
---

# Does a URL just sitting in a prompt steer an LLM's output toward its content?

- HN: [48779535](https://news.ycombinator.com/item?id=48779535)
- Source: [aifoc.us](https://aifoc.us/influencing-model-output-with-urls/)
- Score: 2
- Comments: 0
- Posted: 2026-07-03T20:23:02Z

## Translation

タイトル: URL がプロンプト内に存在するだけで、LLM の出力がそのコンテンツに向けられますか?
記事のタイトル: プロンプト内の URL は llm の出力をそのコンテンツに誘導しますか? | AI フォーカス
説明: 不透明な URL が記憶されたコンテンツへのポインターとして機能するかどうかをテストし、JavaScript でレンダリングされたサイトがモデルのトレーニングから完全に欠落している可能性があることを発見します。

記事本文:
プロンプト内の URL は、llm の出力をそのコンテンツに誘導しますか?
#035 AIフォーカス
プロンプト内の URL は、llm の出力をそのコンテンツに誘導しますか?
発行者
ポール・キンラン
日付: 2026 年 7 月 3 日。読書
時間:
22
分
最初はとても簡単な投稿でしたが、その後、いくつかのことに気づきました。たくさんのものを作りました。大量のトークンを費やしました…そして、これは私がこれまでに行った投稿の中で最も難しく、長く、最も費用がかかる投稿の 1 つになりました (API コストは小さくありませんでした)。
このことは長年頭の中にありましたが、プロンプトにテクノロジー名が存在するだけで出力がそのテクノロジーに偏ってしまうのではないかと考えたときに始まりました。
たとえば、Agentic ツールのシステム プロンプトをいくつか調べてみたところ、(React など) のようなテキストが含まれており、これらのツールは React コードを出力するのに対し、React について言及していない同様のプロンプトは React コードを出力するように感じました。
私はここ数週間、このかゆみを抑えるための実験を行ってきました。しかし、行き着く前に、助けてほしいことがあります。私は研究者ではありません。ここにあるものは説得力のある情報 (または少なくとも何かを教えてくれた) だと思いますが、多くの間違いを犯したり、出力に偏りのある仮定を立てたりした可能性があります。何かアドバイスがございましたら、ぜひお聞きしたいです。メールでご連絡ください。
私が抱いた質問は、プロンプト内に URL が存在すると、その URL のコンテンツまたは URL 自体のリテラル テキストに基づく LLM の出力に影響するでしょうか。
そうであれば、プロンプトに多くのコンテキストを埋め込む必要がなくなる可能性があります。たとえば、モデルの重みに深く統合されているスキル ファイルがあり、「知っていることを使用して、https://skills.sh/super-security-reviewer で詳細な分析を行う」と言うと、モデルの潜在空間内の情報によって出力がバイアスされる可能性があります。

その URL でエンコードされたコンテンツに送信されます。
プロンプト内の URL は出力に影響しますが、それはその URL とそのコンテンツがモデルのトレーニング データに組み込まれた場合に限られます。
LLM プロバイダーがトレーニングの対象となるデータをどのように収集しているかは非常に不明瞭であり、私たちに教えてもらうべきだと思います。
モデルにないデータが山ほどある
サイトが JavaScript に依存してコンテンツを読み込んでいる場合、そのコンテンツはモデルに含まれていない可能性が高くなります (これは機能であると考えることもできます)。私が検証できたトレーニング クローラー (ClaudeBot、GPTBot) はページのアセットを取得しますが、JavaScript は実行しません。私が JavaScript を実行していることを確認した唯一の検証済みボットは、OpenAI の検索クローラーである OAI-SearchBot です。
以下は私が行った旅です。
最初のステップは、さまざまなモデルにわたってさまざまな URL を分析し、仮説のテストに役立つ判定として LLM を使用できるシステムを構築することでした。私の計画は次のとおりでした。
各モデルの既知の「ナレッジ カットオフ日」を確認する
次に、その両側のコンテンツを見つけて、モデル内で知られているはずだと思われるデータをモデルが再現できるかどうかをテストします。
人気がありそうなコンテンツから、おそらく難解なものまで、さまざまなコンテンツを見つけてください。
カットオフの後にあることがわかっているコンテンツは、幻覚を抑えるのに役立ちます。私の最初の仮説が正しければ、そのコンテンツについてモデルは自信を持って何かをでっち上げるのではなく、拒否するか、分からないと言うべきです。
データを取得したら、モデルがどのように機能するかを理解するためにさまざまなテストを作成しました。テストは次のように分類されます。
description : 言葉で説明されたタスク、URL なし (ベースライン)
opaque-url : 不透明な URL 文字列のみであり、ページはフェッチされません。
mdn-url-only / spec-url-only / bcd-key-only : オプションの識別子プローブ。メインの比較の一部ではありません。
url+description : 不透明な URL p

説明されているタスク以外
フルコンテンツ / コンテンツのみ : タスクの詳細が記載されている場合とない場合に貼り付けられた実際のページ (天井)
fake-structural-url / random-url : コントロール (同じ形状の存在しない URL、および無関係な実際の URL)
opaque-url は、LLM がリテラル URL 文字列からコンテンツを推測できないことを確認するための実際のテストでした。たとえば、chromestatus.com (Chrome 機能の公開ダッシュボード) の URL をいくつか使用しました。 https://chromestatus.com/feature/5157805733183488 のような URL があり、LLM にとってそれらが Web 関連であることは明らかだと思いますが、それが CSS Gap Decorations に関するものであると推測することはできません。
次に、説明的な URL (たとえば、MDN は非常に説明的で、Web にとって非常に優れた UX) などの他のテストを実行して、リテラル URL が出力に影響を与えるかどうか、また追加のコンテキストを追加したときに何が起こるかを検証しました。
ここにレポートがあり、すべてのデータがここにあります (iframe も)。一見の価値があると思います。非常に明確な全体像と私の質問に対する答えがあります。
私の最初の予感は、URL はマジック コンテキストではなく、ChromeStatus の数値がそれを裏付けているように思えました。 ChromeStatus 機能の URL は、ドメインがモデルにそのページが Web 関連であることを伝えるため、不透明なテストとしては優れていますが、数値の機能 ID ではその背後にあるものについては何もわかりません。また、ほとんどのモデルは、その数値だけから適切な API を復元できませんでした。プロンプトに裸の不透明な URL を追加しても、平均してほとんど何も起こらず、多くの不透明な URL ではまったく何も回復されませんでした。
しかし、他にも非常に記憶に優れた URL がたくさんありましたが、そうでない不透明な ID も多数ありました。 StackOverflow の 1 つは混在していましたが、robots.txt を見てみると、ほとんどすべてが否定されていました。ふーむ。 ChromeStatus とは何ですか?確認しました

robots.txt で問題ないようでした。おそらく ChromeStatus URL が何らかの理由でモデルに含まれていないだけかもしれません。たとえば、Chrome の最も人気のある機能の 1 つである Service Worker は、URL から呼び出すことができませんでした… それはただ奇妙でした。
モデルがデータの取り込みに何を使用しているかを探しに行きましたが、クロール データの正確なコーパスを見つけるのは少々難しいですが、少し前のポッドキャストで、大量のデータのソースとして使用されている Common Crawl について説明していたことは覚えていました。そこで、Chromestatus が共通クロールに含まれているかどうかを確認してみました。そうです。これらのページは、ほぼ完全にデコードされた arXiv 論文と同じくらい頻繁に Common Crawl に表示されます。しかし、実際にクロールされたバイトを取り出してみると、そこにはコンテンツがありませんでした。
ChromeStatus は JavaScript アプリ (最初は Polymer で構築されたことを覚えています) で、クローラーは空のシェルをキャプチャしました。 CSS ギャップ デコレーション用に保存されたページは、「Chrome プラットフォーム ステータス」という 22 文字の表示テキストを含む約 3 KB の HTML であり、機能については一言も書かれていません (これは実際の Common Crawl キャプチャです)。 4 つの機能をチェックしましたが、それらはすべて同一の空の殻でした。対照的に、arXiv ページはサーバーでレンダリングされるため、クロールでは完全なタイトルと要約 (そのキャプチャ) が保持されます。
Common Crawl がデータ ソースである場合、ユーザーにデータを取得するために JS を必要とする SPA は、モデルのトレーニング データに含まれていない可能性が非常に高いと断言します (一部の人々にとっては、それが機能である可能性があります - へー)。私の証拠は、裸の ChromeStatus ID ですべてのモデルのフラットラインを監視し、実際のページに渡された機能を、ここにあるテストごとのビューで復元できることです。
私は振り払うのがさらに難しい2番目のケースを見つけました、そしてそれは私の「制御された」ビフォーアフターの役割を果たします。 「適応型 COVID-19 治療試験」はその好例です。

Clinicaltrials.gov で。数年前、サイトサーバーがそのページをレンダリングし、Common Crawl の 2022 年の治験のキャプチャがすべてです。「適応型 COVID-19 治療試験 (ACTT) - フルテキストビュー」というタイトルの 47,000 文字の表示テキストで、その中にすべて新型コロナウイルス、レムデシビル、プラセボが含まれています (古いキャプチャ)。その後、clinicaltrials.gov は JavaScript のシングルページ アプリに移行したようです。 Common Crawl によるまったく同じ治験の 2026 年のキャプチャは、「ClinicalTrials.gov 用語集を表示 用語を検索…」という 175 文字の表示テキストを含む 94 KB の HTML であり、新型コロナウイルスやレムデシビルについては 1 つも言及されていません (新しいキャプチャ)。
パンデミックに関する最も文書化された試験の 1 つは、完全に進行中の状態から事実上白紙状態になりました。いずれにせよ、モデルは依然として裸の URL から半分 (モデル全体で約 47%) を呼び出しており、その理由が重要です。 NCT ID は remdesivir の文献全体で引用されており、ページは移行直前までサーバーでレンダリングされ、クロール可能であったため、古いコンテンツはほぼ確実にすでにウェイトに組み込まれています。移民が壊すのは未来だ。 Clinicaltrials.gov が今後公開するものはすべて JavaScript でのみレンダリングされるため、おそらくクロールされることはありません。したがって、Common Crawl から欠落していることは、モデルから欠落していることと同じではありません。これはむしろスライドスケールです。NIST でサーバーでレンダリングされ、広く引用されている CVE は約 92% の確率で裸の URL から返されますが、このトライアル (現在はシェルですが、何年もクロールされ、依然としてどこでも引用されています) は約 47%、ChromeStatus 機能 (ブラウザーでレンダリングされ、どこにも引用されていません) はフラット ゼロです。
この空間全体が濁っていて、それを最も濁らせるのはレンダリングです。すべてのテスト URL の内容が生の HTML 内に存在するか、JavaScript が実行されたときにのみ表示されるかによってラベルを付けました。

裸の URL からのリコールを調べました。クライアントがレンダリングした 31 個の項目 (主に ChromeStatus 機能) の再現率は平均 6% で、そのうち 25 個はフラット ゼロです。これらは、あいまいな機能でもありません (ビュー遷移、ポップオーバー、アンカー位置、Temporal API)。サーバーでレンダリングされた 60 のソース (arXiv、CVE、RFC、Wikipedia) の平均は 55% です。名声はほぼ一定に保たれ、HTML にすでに存在していたコンテンツは、ブラウザが組み立てる必要があるコンテンツよりも約 9 倍よく思い出されます。
「もしかしたらクロールされていないだけかもしれない」という疑念を完全に払拭したかったので、モデル内でコンテンツが疑問の余地のないケースを試してみました。すべての Wikipedia 記事には、直接アドレス指定できる内部数値 ID があります。en.wikipedia.org/?curid=24544 は光合成です。コンテンツはサーバーでレンダリングされ、間違いなくどのモデルでも使用されます。しかし、URL の ?curid= 形式は私が調べたどのクロール インデックスにもありませんでしたが、正規の en.wikipedia.org/wiki/PhotoSynthetic URL はすべてのインデックス (200、全文) に含まれています。これは、Wikipedia が正規のタイトル URL で curid ページを指しており、クローラーはそれを尊重するためです。 5つの記事をチェックしました。すべての /wiki/ が存在し、すべての ?curid= が存在しません。名前を尋ねると、モデルは完璧なスコアを獲得し、記事を貼り付けると、モデルは完璧なスコアを獲得し、裸の数字の ID を伝えると、わーわー、まったくだめです。 5 つすべて同じ形状: 光合成 、 トランスフォーマー 、 ​​ミトコンドリア 、 HTTP 404 、 ビットコイン 。
したがって、裸の不透明な URL はほとんど何も行いません。しかし、URL が明らかにその重要性を発揮するケースが 2 つあり、どちらも ChromeStatus のストーリーと矛盾しません。
説明的な URL は出力に影響します。 URL に React 、 fetch 、 text-justify などの単語が含まれている場合、それらの単語は単なる通常のプロンプト テキストであり、モデルはそれらを他のトークンと同様に使用します。
一部の有名な不透明な識別子は実際にデコードします。ランドマー

k arXiv ID、古典的な RFC 、およびよく知られた CVE は、裸の識別子だけから驚くほどうまくコンテンツを復元します。他のヒントなしで arxiv.org/abs/1706.03762 だけから、モデルは「Attending Is All You Need」とトランスフォーマー (その裸の ID 上のすべてのモデル) を再構築します。これは、「URL がライブ コンテンツをポイントしている」というよりは、「この識別子とそのコンテンツが、記憶されるほど十分な頻度でトレーニング データに一緒に出現している」ということのように見えます。そして、それはスイッチではなくグラデーションです。デコードは有名な識別子に対しては強力ですが、内容がより曖昧になるにつれて着実に弱まり、ロングテールではほぼゼロになります。 GitHub コミットを使用して、その勾配を直接監視できます。有名な最初の Linux 、 Git 、および Bitcoin へのコミットは裸の SHA からデコードされますが、同じ種類のリポジトリからの通常のルーチン コミットは何も返しません。知識の遮断も同様に影響を及ぼします。たとえそれが有名な情報源であっても、消滅後に出版されたものはすべて。
これは、React に関する私の最初の質問に戻る部分です。上記のすべては、コマンドに応じて URL をデコードするようにモデルに要求します。しかし、私が本当に知りたかったのは、システム プロンプトで React に言及すると、コードが React に傾くように見えるように、プロンプト内に存在する URL が出力を傾けるかどうかでした。そこで私は 2 番目の実験を実行しました。

[切り捨てられた]

## Original Extract

Testing whether opaque URLs act as pointers to memorized content, and finding that JavaScript-rendered sites may be missing from model training entirely

does a url in a prompt steer an llm's output toward its content?
#035 AI Focus
does a url in a prompt steer an llm's output toward its content?
Published by
Paul Kinlan
on: July 3, 2026; Reading
time:
22
minutes
At first, this was a really easy post to write, but then I discovered some things. Built a lot of things. Spent a lot of tokens… And it became one of the hardest, longest, and most expensive posts I’ve done (the API costs were not small).
I’ve had this thing on my mind for ages and it started when I was thinking about how the mere presence of a technology name in a prompt seemed to bias the output to that technology.
For example, I looked through a number of system prompts for Agentic tooling and they would include text like (e.g. React) and then it felt like these tools would output React code vs a similar prompt that didn’t mention React.
I’ve spent the last few weeks running experiments to scratch this itch. But before I get too far, I have a request for help. I’m not a researcher. I think what I have here is compelling information (or at least it taught me something), but I might have made a lot of mistakes or made assumptions that have biased the output. If you have any advice I would LOVE to hear from you. Email me .
The question I had was: would the presence of a URL in a prompt influence the output of the LLM, based on the content at that URL or the literal text of the URL itself?
If yes, then this could lead to us not having to embed lots of context into the prompt. For example, you might have a Skills file that is deeply integrated into the model’s weights and by saying “use what you know about: https://skills.sh/super-security-reviewer do a deep analysis” then information in the model’s latent space would bias the output towards the content encoded at that URL.
A URL in the prompt does influence the output, but only when that URL and its content made it into the model’s training data
It’s really unclear how LLM providers gather the data they train on, and I think they should tell us.
There’s heaps of data that is not in the models
If your site relies on JavaScript to load its content, that content is very likely not in a model (you might consider that a feature). The training crawlers I could verify (ClaudeBot, GPTBot) fetch a page’s assets but never execute the JavaScript; the only verified bot I’ve caught running JavaScript is OpenAI’s search crawler, OAI-SearchBot.
What follows is the journey I took.
The first step was to build a system that can analyse a range of URLs across a range of models and use an LLM-as-a-judge to help me test the hypothesis. My plan was:
to find each model’s known “Knowledge Cut-off date”
then find content on either side of that to test if the model could recall the data that I believe should be known in the model.
find ranges of content ranging from content that I believe would be popular all the way to likely esoteric.
Content known to be after a cutoff would help me control against hallucination. If my original hypothesis was correct, then for that content the model should decline, or say it doesn’t know, rather than confidently make something up.
Once I had the data I created a range of tests to help me understand how the models work. The tests are classified as:
described : the task described in words, no URL (the baseline)
opaque-url : ONLY the opaque URL string, and the page is never fetched
mdn-url-only / spec-url-only / bcd-key-only : optional identifier probes, not part of the main comparison
url+described : the opaque URL plus the task described
full-content / content-only : the real page pasted in, with and without the task spelled out (the ceiling)
fake-structural-url / random-url : controls (a nonexistent URL of the same shape, and an unrelated real URL)
opaque-url was my real test, to try to ensure that the LLM couldn’t infer the contents from the literal URL string. So for example I used some URLs from chromestatus.com (which is our public dashboard of Chrome features) because it has URLs like https://chromestatus.com/feature/5157805733183488 , and while I believe it’s pretty clear to the LLM that they are web-related, you can’t infer that it’s about CSS Gap Decorations.
I then had other tests, like descriptive URLs (MDN for example is very descriptive, which is very good UX for the web) to validate whether the literal URL influenced the output, as well as what happens when we add in extra context.
I have a report here and all the data is here (iframed too). I think it’s worth looking at, and there’s a pretty clear picture and answer to my question.
My first hunch was that URLs are not magic context , and the ChromeStatus numbers seemed to back it up. ChromeStatus feature URLs are a good opaque test because the domain tells the model the page is web-related, but the numeric feature ID tells it nothing about what is behind it, and most models failed to recover the right API from that number alone. Adding a bare opaque URL to a prompt did almost nothing on average, and plenty of opaque URLs recovered nothing at all.
But then I had a lot of other URLs that had really good recall, and a lot of other opaque IDs that didn’t. StackOverflow, for one was mixed, and then I looked at their robots.txt and it’s pretty much deny everything. Hmm. What’s ChromeStatus’s? I checked its robots.txt and it looked fine… maybe ChromeStatus URLs are just not in the model for some other reason. For example, one of Chrome’s most popular features, Service Worker , couldn’t be recalled from the URL… It was just odd.
I went to look for what the models use to ingest data, and it’s kinda hard to find the exact corpus of crawl data, but I did remember a podcast from a little while ago that discussed Common Crawl being used as a source of a lot of data. So I went to check if Chromestatus was in the common crawl. It is. The pages show up in Common Crawl about as often as the arXiv papers that decode almost perfectly. But when I pulled the actual crawled bytes, there was no content in them!!!
ChromeStatus is a JavaScript app (I remember it first being built with Polymer) and the crawler captured an empty shell. The saved page for CSS Gap Decorations is about 3KB of HTML with 22 characters of visible text, “Chrome Platform Status”, and not one word about the feature ( here is the actual Common Crawl capture ). I checked four features and they were all identical empty shells. The arXiv page, by contrast, is server-rendered, so the crawl holds the full title and abstract ( its capture ).
If Common Crawl is a source of data, then I’m going to flat out say that SPAs that require JS to get data to the user are very likely to not be in the models training data (that might be a feature for some folks - heh.) My evidence is that you can watch every model flatline on the bare ChromeStatus id, then recover the feature once handed the actual page, in the per-test view here .
I found a second case that is even harder to wave away, and it doubles as my “controlled” before-and-after. “The Adaptive COVID-19 Treatment Trial” is a good example because it is on clinicaltrials.gov. A couple of years ago the site server-rendered its pages, and Common Crawl’s 2022 capture of the trial is the whole thing: 47,000 characters of visible text, titled “Adaptive COVID-19 Treatment Trial (ACTT) - Full Text View”, with COVID, remdesivir, and placebo all through it ( the old capture ). Then it appears that clinicaltrials.gov migrated to a JavaScript single-page app. Common Crawl’s 2026 capture of the very same trial is 94KB of HTML carrying 175 characters of visible text, “ClinicalTrials.gov Show glossary Search for terms…”, and not one mention of COVID or remdesivir ( the new capture ).
One of the most documented trials of the pandemic went from fully present in the crawl to effectively blank. The models still half-recall it from the bare URL anyway, around 47% across models , and the reason matters. The NCT id is cited all through the remdesivir literature, and the page was server-rendered and crawlable right up until the migration, so the old content is almost certainly already baked into the weights. What the migration breaks is the future. Anything clinicaltrials.gov publishes from here on renders only in JavaScript and will probably never make it into the crawl. So being missing from Common Crawl is not the same as being missing from the model. It’s more of a sliding scale: a server-rendered, widely-cited CVE over at NIST comes back from the bare URL about 92% of the time, this trial (a shell now, but crawled for years and still cited everywhere) about 47%, and a ChromeStatus feature (rendered in the browser and cited nowhere) a flat zero.
This whole space is murky, and rendering is what muddies it most. I labelled every test URL by whether its content sits in the raw HTML or only shows up once JavaScript runs, then looked at recall from the bare URL. The 31 client-rendered items, mostly ChromeStatus features, average 6% recall, and 25 of them are a flat zero. These are not obscure features either (view-transitions, popover, anchor positioning, the Temporal API). The 60 server-rendered sources (arXiv, CVEs, RFCs, Wikipedia) average 55%. Hold fame roughly constant, and content that was already in the HTML recalls about nine times better than content a browser has to assemble.
I really wanted to kill the “maybe it just wasn’t crawled” doubt entirely, so I tried a case where the content is beyond question in the model. Every Wikipedia article has an internal numeric id you can address directly: en.wikipedia.org/?curid=24544 is Photosynthesis. The content is server-rendered and unquestionably in every model. But the ?curid= form of the URL is in none of the crawl indexes I looked at, while the canonical en.wikipedia.org/wiki/Photosynthesis URL is in all of them (200, full text), because Wikipedia points the curid page at the canonical title URL and the crawler respects that. I checked five articles; every /wiki/ present, every ?curid= absent. Ask by name and the models score perfectly, paste the article in and they score perfectly, give the bare numeric id and wah wah, a fat nope. Same shape on all five: Photosynthesis , the Transformer , Mitochondrion , HTTP 404 , Bitcoin .
So the bare opaque URL mostly does nothing. But there are two cases where a URL clearly does pull its weight, and neither of them contradicts the ChromeStatus story.
Descriptive URLs influence output. If the URL contains words like React , fetch , or text-justify , those words are just normal prompt text, and the model uses them like any other token.
Some famous opaque identifiers really do decode. Landmark arXiv IDs, classic RFCs , and well-known CVEs recover their content surprisingly well from the bare identifier alone. From just arxiv.org/abs/1706.03762 , with no other hint, the models reconstruct “Attention Is All You Need” and the transformer ( every model on that bare id ). That looks less like “the URL points to live content” and more like “this identifier and its content appeared together often enough in the training data to be memorized”. And it’s a gradient, not a switch: the decoding is strong for famous identifiers and fades steadily as the content gets more obscure, down to roughly nothing for the long tail. You can watch that gradient directly with GitHub commits. The famous first commits to Linux , Git , and Bitcoin decode from the bare SHA, while ordinary routine commits from the same kinds of repos return nothing at all. The knowledge cutoff bites the same way. Anything published after it is gone, even for otherwise well-known sources.
This is the part that gets back to my original question about React. Everything above asks the model to decode the URL on command. But the thing I really wanted to know was whether a URL just sitting there in the prompt tilts the output, the way mentioning React in a system prompt seems to tilt code towards React. So I ran a second experiment where I

[truncated]
