---
source: "https://www.agentvoyagerproject.com/captains-log/2"
hn_url: "https://news.ycombinator.com/item?id=48359075"
title: "We gave an AI agent eyes. It didn't even use them"
article_title: "We gave an AI agent eyes. It didn't even use them · Captain's Log #2 · Agent Voyager Project"
author: "pmkelly4444"
captured_at: "2026-06-02T00:36:45Z"
capture_tool: "hn-digest"
hn_id: 48359075
score: 2
comments: 0
posted_at: "2026-06-01T16:29:47Z"
tags:
  - hacker-news
  - translated
---

# We gave an AI agent eyes. It didn't even use them

- HN: [48359075](https://news.ycombinator.com/item?id=48359075)
- Source: [www.agentvoyagerproject.com](https://www.agentvoyagerproject.com/captains-log/2)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T16:29:47Z

## Translation

タイトル: AIエージェントに目を与えました。それはそれらを使用さえしませんでした
記事タイトル: AIエージェントに目を与えました。それらも使用しませんでした · キャプテンズ ログ #2 · エージェント ボイジャー プロジェクト
説明: 私たちは、見つけられた最も厄介な PDF テーブルの 1 つに安っぽい Haiku を指定し、AVP ですべてのステップを記録しました。エージェントの「目」は決して機能しませんでした。とにかく根性でテーブルを釘付けにした。

記事本文:
AIエージェントに目を与えました。それらも使用しませんでした · キャプテンズ ログ #2 · エージェント ボイジャー プロジェクト avp | \ | _\ | __\ | ___\ | ____\ | _____\ \______________/ \____________/ / キャプテンのログ # 2 powered by Port of Context AVP キャプテンのログとベンチマークについて フルスペック ↗ AI エージェントに目を与えました。それらも使用しませんでした。
Opus 4.8 の価格を確認したとき、模型通路の一番下の棚がどのようなものかを見てみることにしました。その結果、ある種の不況耐性ベンチマークが得られました。つまり、固体エージェント ハーネス ( Goose ) で包まれている場合に、安価なモデルでどれだけのハードワークを達成できるかということです。
そこで私たちは Claude Haiku 4.5 に手を出し、抽出するのに非常に面倒な PDF ページ (arXiv の論文からそのまま引用した ParseBench のページ) を提供しました。私たちは 2 つのエージェント構成をテストしました。1 つは pdf-vision MCP サーバーを介して表示できるもので、もう 1 つは Goose の組み込み pdf_tool を介してテキストの読み取りのみが可能でした。
読み進める前に、私たちが発見したことについての軽いネタバレをいくつか紹介します。
優れたハーネス (Goose) は、適切なツールにアクセスできれば、古くて安価なモデルでも本当に困難なタスクを実行できます。
エージェントにビジョンという贈り物を与えても、それを実際に活用するわけではありません。この場合、エージェントは視覚をまったく使用しませんでした。
以下のすべてのステップは、エージェントの動作をキャプチャするための無料でオープンな、プラットフォームに依存しない標準である Agent Voyager Project (AVP) を使用して記録されます。数字と引用符は、 claude-haiku-4-5 の軌跡からそのまま引用しています。
1 枚のシートに 4 つのテーブルが詰め込まれている
これは、ParseBench から抜粋した 2012 年の計量経済学の論文の 7 ページです。 4つの独立したテーブルが詰め込まれています。重要なのは表 7 です。2 つの 6 行 6 列の相関行列が互いに積み重ねられており、三角形でセルの半分が空白で、値は 0.47 [0.49] のような値です。

数値は 1 つのセルを共有します。
私たちが Goose に与えたタスクは、簡単に言うと、ページをダウンロードし、HTML テーブルとして再構築し、誤解しないでください。
まず、明らかな動き。 Goose の組み込み PDF リーダー ( pdf_tool 、 pdfplumber ラッパー) は、ページからテキストを取り出します。返ってきたものはこちらです。
… March FB 4 - 7.309 O 9 - 1.513 69.312 1531.360 7.270 表 7. 6 つの C P I タイムシリーズの相互相関係数とその最初の差異。オリジナルのシリーズには、1 24 回の読みと、最初の差分 123 回の読みが含まれます。 F FB SEF V O R PR R SH O F 1 FB 0.99998 1 SE F V 0.99714 0.99671 1 O R PR 0.98356 0.98295 0.98702 1 R SH 0.97533 0.97478 0.97736 0.99698 1 O 0.97752 0.97661 0.98664 0.95629 0.93924 1 d F d FB d SE F V d O R PR d R SH d O d F 1 d FB 0.994 1 d SE F V 0.47 [ 0. 49] 0.4 8 [ 0. 49] 1 … ページ上のすべてのテーブルに注がれます1 つのランオン ストリーム。 8月ORP R .時間はかかります。行も列もなく、あるテーブルがどこで終わり、次のテーブルが始まるのかを知る方法もありません。 Goose は構造用の画像を取得しようとしましたが、「PDF に画像が見つかりません」という結果が返されたため、テキストでも機能しました。マトリックスを再構築し、再読み込みして確認し、勝利を宣言しました。
スコアは 53% で、明らかな失敗でしたが、ラン中に何も問題はありませんでした。5 ターン 5 セントで、ボード上の最も難しいページで記録した最速のランとなりました。スコア自体は「失敗」と表示されているだけですが、軌跡はさらに悪いものを示しています。これはエージェントが自信を持って間違っており、何も知らなかったということです。
、モデル claude-haiku-4-5 ツール シェル、書き込み、編集、computercontroller__web_scrape、computercontroller__pdf_tool 航海全体が記録されました: 素早く、スムーズで、運命の 5 ターン。再生ボタンを押してから、「AVP JSON の全文を表示」を開いて、各ステップでエージェント自身の言葉を読んでください。
ダウンロード

{input} にある PDF ページを読み取って、単一の HTML <table> として再構築し、元のページを再度読み込み、戻る前にテーブルが一致していることを確認します。 HTMLのみを出力します。試行 2 Goose + pdf-vision 24 ターン $0.33 100 % 合格 同じモデル、同じページ。今回は、内蔵リーダーを pdf-vision に交換しました。pdf-vision は、モデルがそれを見ることができるようにページを画像にレンダリングする見出しトリックを備えた MCP サーバーです。私たちはグースに、プロンプトの中で、テキストよりも画像を信頼するように指示しました。
目が機能しませんでした。画像は空のまま返され続けるため、Goose は別の方法でページを見ようと試み、新しいアプローチが失敗するたびに試み続けました。
「では、テーブル構造を確認するためにページを画像としてレンダリングしてみましょう。」
「表示できるように、ファイル出力付きのページ画像を取得させてください。」
「では、コンピューターのコントローラーを使用して画像を見てみましょう。」
「コンピューターのコントローラーを使用して画像を表示させてください。」
「PDF をさらに注意深く調べて、どのページに表が含まれているかを特定させてください。」
「各 Y 位置は行を表します。完全なテーブルを作成しましょう。」
「まったく異なるアプローチを取らせてください。」
「完璧です！マークダウン形式で表がわかりやすく表示されます。」
これは、画像を見ようとする試みが 8 回ありましたが、読み込まれることはありませんでした。そのうちの 1 回は、手動で読み取ろうと、端末の Base64 を介して PNG をパイプしました。目が失敗すると、Goose は同じツールの別のトリックであるレイアウト認識テキストに戻り、フラット リーダーが破壊した構造を使用してページをマークダウンとしてエクスポートしました。
表 7. 6 つの CPI 時系列の相互相関係数
そして彼らの最初の違い。
||F|FB|SEFV|ORPR|RSH|O|
|---|---|---|---|---|---|---|
|F|1||||||
|FB|0.99998|1|||||
|SEFV|0.99714|0.99671|1||||
|ORPR|0.98356|0.98295|0.98702|1|||
|RSH|0.97533|0.97478|0.97736|0.99698|1||
|O|0.97752|0。

97661|0.98664|0.95629|0.93924|1|
||dF|dFB|dSEFV|dORPR|dRSH|dO|
|---|---|---|---|---|---|---|
|dF|1||||||
|dFB|0.994|1|||||
|dSEFV|0.47 [0.49]|0.48 [0.49]|1||||
|dORPR|0.12 [0.26]|0.12 [0.26]|0.31 [0.35]|1|||
|dRSH|0.13 [0.30]|0.12 [0.28]|0.10 [0.29]|0.31 [0.37]|1||
|dO|-0.18 [0.30]|-0.18 [0.28]|0.06 [0.29]|0.002 [-0.21]|0.04 [-0.29]|1|このエクスポートには実際の行と列が含まれており、2 つの行列は分離されており、括弧で囲まれたセルはそのままです。そこからは簡単で、グースがテーブルを完成させました。
プレーンテキストでは触れられないボード上の 1 ページで、グースは 100% 得点して合格しました。
、モデル claude-haiku-4-5 ツール シェル、書き込み、編集 mcp pdf-vision 同じ実行の 24 ターン、エンドツーエンドで記録されました。今回の航海が前回の航海よりどれほど長くなるか見てみましょう。
PDF ページから表を単一の HTML <table> として再現します。
1. ダウンロードします:curl -sL '{input}' -o page.pdf
2. pdf-vision get_page_image ツールを使用してページを画像にレンダリングし、それを確認します。画像は 2D レイアウトの正確な情報です。つまり、列がいくつあるか、どのセルがマージまたは複数の行にまたがっているか、セクション ヘッダー行、およびどの値が 1 つのセルを視覚的に共有しているか (たとえば、1 つのセル内の複数のホルダー) です。 get_page_text は、正確なテキストをコピーする場合にのみ使用してください。画像の構造を信頼してください。
3. 表示内容と一致する 1 つの HTML <table> として再構築します。ビジュアル行ごとに 1 つの <tr>、ヘッダー セルの <th>、データの <td>、結合されたセルの Colspan/rowspan です。列の順序と正確なセルのテキストを維持します。値を余分な列に分割したり、視覚的に別々の行を結合したりしないでください。
4. 画像と照らし合わせて確認します。ヘッダーの列数と各行のセル数がページと一致する必要があります。そうでない場合は修正してやり直します。明らかに間違っているとわかる表は提出しないでください。
5. 最終的な HTML <table> のみを出力します。私たちが実際に学んだこと 数字がもたらす 3 つのこと

教えてください。
モデルはそのページを一度も見たことがありません。勝利は完全にテキストによるもので、特に構造を損なわないように保つツールによるものでした。 pdf-vision は、実際にこの実行を救ったものの誤解を招く名前です。
安価なモデルが根性の難しいページをクリアした。
Haiku は 2 回の実行の間に賢くなることはありませんでした。変わったのは、2回目には頑固に扱う価値のあるツールと、止まらないハーネスを持っていたことだ。 24 ターン、8 つの行き止まり、1 つの静かなピボット、完璧な答えです。それは、弱いモデルをフィニッシュまでずっと苦しめるハーネスです。
これは記録されたものだからこそ分かることです。
金額に換算すると、ビジョンのコストは約 7 倍になります。スコアは視覚が勝利したことを示しています。どちらも、エージェントが一度も目を使っていなかったとか、本当のヒーローがターン 22 に到達したマークダウン エクスポートだったということを示すものではありません。AVP は、すべてのステップ、すべてのツール呼び出し、エージェント自身の推論のすべての行を 1 つのオープン形式でキャプチャします。 「ビジョンが勝った」と「粘り強さが勝った、ここが正確なターンだ」との間のギャップこそが、私たちが軌跡を記録する理由のすべてです。
ハーネスが優れていて、それが何をしているのかを確認できる場合、安価なモデルでも、その価格が示すよりも難しいことを行うことができます。私たちはそれを突き続けます。研究室からの詳細は近日中にお知らせします。

## Original Extract

We pointed cheap Haiku at one of the nastiest PDF tables we could find and recorded every step with AVP. The agent's 'eyes' never worked. It nailed the table anyway, on grit.

We gave an AI agent eyes. It didn't even use them · Captain's Log #2 · Agent Voyager Project avp | \ | _\ | __\ | ___\ | ____\ | _____\ \______________/ \____________/ / Captain's Log # 2 powered by Port of Context About AVP Captain's Log & Benchmarks Full Spec ↗ We gave an AI agent eyes. It didn't even use them.
When we saw how much Opus 4.8 cost, we decided to take a look at what the bottom shelf of the model aisle looked like. What resulted is a sort of recession-proof benchmark: how much hard work can a cheaper model accomplish, provided it's wrapped by a solid agent harness ( Goose )?
So we reached for Claude Haiku 4.5 and gave it an extremely annoying PDF page to extract (a page from ParseBench , lifted straight from an arXiv paper). We tested two agent configurations: one that could see , via the pdf-vision MCP server, and one that could only read text, via Goose's built-in pdf_tool .
A few gentle spoilers on what we found, before you read on:
A good harness (Goose) can deliver on a genuinely tough task with an older, cheaper model, provided it has access to the right tools.
Blessing an agent with the gift of vision doesn't actually mean it is going to use it. In this case, the agent didn't use its eyesight at all.
Every step below is recorded with the Agent Voyager Project (AVP) , a free, open, platform-agnostic standard for capturing what an agent does. Numbers and quotes are verbatim from the trajectories, on claude-haiku-4-5 .
Four tables crammed onto one sheet
This is page 7 of a 2012 econometrics paper, pulled from ParseBench . Four separate tables are crammed onto it. The one that matters is Table 7 : two six-by-six correlation matrices stacked on top of each other, triangular, half the cells blank, and values like 0.47 [0.49] where two numbers share one cell.
The task we gave Goose was easy to state: download the page, rebuild it as an HTML table, do not get it wrong.
First, the obvious move. Goose's built-in PDF reader ( pdf_tool , a pdfplumber wrapper) pulls the text off the page. Here is what it handed back.
… Mar c h FB 4 - 7.309 O 9 - 1.513 69.312 1531.360 7.270 Ta ble 7. C ross c or r e latio n c oe f fic ien ts fo r six C P I ti me se rie s a nd their fir s t diff e r e nc e s. Or i g inal se rie s include 1 24 r e a din g s, and th e ir f irst di ff e re n c e s 123 r e a din g s. F FB SE F V O R PR R SH O F 1 FB 0.99998 1 SE F V 0.99714 0.99671 1 O R PR 0.98356 0.98295 0.98702 1 R SH 0. 97533 0.97478 0.97736 0.99698 1 O 0.97752 0.97661 0.98664 0.95629 0.93924 1 d F d FB d SE F V d O R PR d R SH d O d F 1 d FB 0.994 1 d SE F V 0.47 [ 0. 49] 0.4 8 [ 0. 49] 1 … Every table on the page, poured into one run-on stream. Aug ust . ORP R . ti me se rie s . No rows, no columns, no way to tell where one table ends and the next begins. Goose even tried to pull images for the structure and got back “No images found in PDF,” so it worked with the text. It rebuilt the matrices, re-read to check, and declared victory.
It scored 53% , a clear fail, and nothing in the run flagged it: five turns and five cents, which made it the fastest run we recorded on the hardest page on the board. The score on its own just says “fail,” but the trajectory shows something worse, which is an agent that was confidently wrong and had no idea.
, model claude-haiku-4-5 tools shell, write, edit, computercontroller__web_scrape, computercontroller__pdf_tool The whole voyage, recorded: a quick, smooth, doomed five turns. Press play, then open “View full AVP JSON” to read the agent's own words at each step.
Download the PDF page at {input}, read it, rebuild it as a single HTML <table>, then re-read the original and verify your table matches before returning. Output only the HTML. Attempt 2 Goose + pdf-vision 24 turns $0.33 100 % passed Same model, same page. This time we swapped the built-in reader for pdf-vision , an MCP server whose headline trick is rendering the page to an image so the model can look at it. We even told Goose, in the prompt, to trust the picture over the text.
The eyes did not work. The image kept coming back empty, so Goose tried to see the page another way, and then kept trying as each new approach failed.
“ Now let me render the page as an image to see the table structure. ”
“ Let me get the page image with file output so I can view it. ”
“ Now let me view the image using the computer controller. ”
“ Let me use computer controller to view the image. ”
“ Let me examine the PDF more carefully to identify which page contains the table. ”
“ Each Y position represents a row. Let me build the complete tables. ”
“ Let me take a completely different approach. ”
“ Perfect! The markdown format shows the tables clearly. ”
That was eight attempts to look at a picture that was never going to load, one of which piped a PNG through base64 in the terminal in the hope of reading it by hand. When the eyes failed, Goose fell back to the same tool's other trick, layout-aware text, exporting the page as markdown with the structure the flat reader had destroyed.
Table 7. Cross correlation coefficients for six CPI time series
and their first differences.
||F|FB|SEFV|ORPR|RSH|O|
|---|---|---|---|---|---|---|
|F|1||||||
|FB|0.99998|1|||||
|SEFV|0.99714|0.99671|1||||
|ORPR|0.98356|0.98295|0.98702|1|||
|RSH|0.97533|0.97478|0.97736|0.99698|1||
|O|0.97752|0.97661|0.98664|0.95629|0.93924|1|
||dF|dFB|dSEFV|dORPR|dRSH|dO|
|---|---|---|---|---|---|---|
|dF|1||||||
|dFB|0.994|1|||||
|dSEFV|0.47 [0.49]|0.48 [0.49]|1||||
|dORPR|0.12 [0.26]|0.12 [0.26]|0.31 [0.35]|1|||
|dRSH|0.13 [0.30]|0.12 [0.28]|0.10 [0.29]|0.31 [0.37]|1||
|dO|-0.18 [0.30]|-0.18 [0.28]|0.06 [0.29]|0.002 [-0.21]|0.04 [-0.29]|1| That export had real rows and columns, the two matrices kept separate, and the bracketed cells intact. From there it was straightforward, and Goose finished the table.
Goose scored 100% and passed, on the one page on the board that plain text couldn't touch.
, model claude-haiku-4-5 tools shell, write, edit mcp pdf-vision Twenty-four turns of the same run, recorded end to end. Watch how much longer this voyage is than the last one.
Reproduce the table from a PDF page as a single HTML <table>.
1. Download it: curl -sL '{input}' -o page.pdf
2. Render the page to an image with the pdf-vision get_page_image tool and LOOK at it. The image is the ground truth for the 2D layout: how many columns there are, which cells are merged or span rows, section-header rows, and which values visually share one cell (e.g. multiple holders inside a single cell). Use get_page_text only to copy exact text; trust the image for structure.
3. Rebuild as ONE HTML <table> that matches what you see: one <tr> per visual row, <th> for header cells, <td> for data, colspan/rowspan for merged cells. Keep column order and exact cell text. Do NOT split a value into extra columns or merge rows that are visually separate.
4. Verify against the image: the header column count and each row's cell count must match the page. Fix and redo if not. Do not submit a table you can see is wrong.
5. Output only the final HTML <table>. What we actually learned Three things the numbers won't tell you.
The model never once saw the page. The win came entirely from text, and specifically from the tool that kept the structure intact. pdf-vision is a misleading name for what actually rescued this run.
A cheap model cleared a hard page on grit.
Haiku did not get smarter between the two runs. What changed is that the second time it had a tool worth being stubborn with, and a harness that refused to stop. Twenty-four turns, eight dead ends, one quiet pivot, a perfect answer. That is the harness wrangling a weaker model all the way to the finish.
You only know this because it was recorded.
The dollar figure says vision cost about 7x more. The score says vision won. Neither tells you the agent never used its eyes, or that the real hero was a markdown export it reached for on turn 22. AVP captures every step, every tool call, every line of the agent's own reasoning, in one open format. The gap between “vision won” and “persistence won, here is the exact turn” is the whole reason we record trajectories.
Cheaper models can do harder things than their price tag suggests, if the harness is good and you can see what it is doing. We are going to keep poking at that. More from the lab soon.
