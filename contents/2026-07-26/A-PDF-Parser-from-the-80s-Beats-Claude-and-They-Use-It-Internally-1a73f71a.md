---
source: "https://fagnerbrack.com/a-pdf-parser-from-the-80s-beats-claude-and-they-use-it-internally-8ee45a533e80"
hn_url: "https://news.ycombinator.com/item?id=49058324"
title: "A PDF Parser from the 80s Beats Claude (and They Use It Internally)"
article_title: "Medium"
author: "fagnerbrack"
captured_at: "2026-07-26T14:28:07Z"
capture_tool: "hn-digest"
hn_id: 49058324
score: 1
comments: 0
posted_at: "2026-07-26T14:00:21Z"
tags:
  - hacker-news
  - translated
---

# A PDF Parser from the 80s Beats Claude (and They Use It Internally)

- HN: [49058324](https://news.ycombinator.com/item?id=49058324)
- Source: [fagnerbrack.com](https://fagnerbrack.com/a-pdf-parser-from-the-80s-beats-claude-and-they-use-it-internally-8ee45a533e80)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T14:00:21Z

## Translation

タイトル: 80 年代の PDF パーサーがクロードを超える (内部で使用されている)
記事タイトル: 中
説明: クロードと一緒に CIA PDF を解析しようとしましたが、1980 年代のパーサーに負けました。適切なツールを選択することは依然として重要です。

記事本文:
Medium 80 年代の PDF パーサーがクロードを超える (内部で使用されている) |フェイナー・ブラック著 |中 サイトマップ アプリで開く サインアップ
80 年代の PDF パーサーがクロードを超える (内部で使用されている)
CIA PDF を解析するという私の旅は、その仕事に適したツールを選ぶということを発見するまでには至りませんでした
1980 年代のスキャナーがページを送り、きれいな文字の線を発している様子を示す温かみのあるトーンのイラスト。スイッチがオフになったクロム球が隅に小さく薄暗く鎮座している。後で戻ってきますか?これを readplace.com に保存します。
80 年代の PDF パーサーがクロードを超える (内部で使用されている) |リーダービュー
Tesseract は 1985 年の OCR ツールで、PDF 解析では最新の AI モデルよりも優れたパフォーマンスを発揮し、Claude によって内部的に使用されています。
AI に PDF を渡すとき、モデルがそれを読んでいる様子を想像します。ただし、特定の種類の PDF の場合、実際のモデルはテキストをほとんど認識せず、実際のソフトウェア エンジニアが以前に指示したツールを呼び出すだけです。
私はスキャンが不十分な CIA PDF を解析するのに何日も費やしました (理由は聞かないでください。さもなければあなたを殺さなければなりません)。私は最も強力な Anthropic モデルと Google モデルを使用してみましたが、最終的には Tesseract (1985 年のツール) の方が優れた仕事をし、10 倍も高速であるという結論に達しました。
それだけではなく、Claude が内部で PDF パーサー スキルで PyTesseract を使用していることがわかりました。彼らは、ヒューリスティックを実行し、最も困難なスキャンされたページをラスタライズし、トークンをストリーミングバックするときに幻覚単語で欠落している文を完成させることによって、それを選択します。
Claude Code AI モデルは他のモデルより優れているわけではなく、熟練したエンジニアに適切なツールを提供してもらうためにお金を払っているだけです。
Claude にはスキル、タスクの処理方法を指示する小さな Markdown 命令ファイルが付属しています。 Anthropic は、フォーマットとこれらのスキルのセットをオープンに公開したため、

それらを食べてください。読書戦略の部分は (どうやら) オープンソースではなく、疲れたエンジニアが午前 2 時に書いたランブックのようなものです。
### 読書戦略の選択
**テキストの多いドキュメント** (レポート、記事、書籍):
→ テキスト抽出が主です。特定の図形のみをラスタライズするか、
レイアウトが重要なページ。
**スキャンされたドキュメント** (「pdffonts」にはフォントが表示されません):
→ `pdftotext` は何も返しません。実行しないでください。ページをラスタライズする
150 DPI で視覚的に読み取ります。一括テキスト抽出には OCR を使用します
(ページを画像に変換した後の pytesseract - 詳細については REFERENCE.md を参照してください)
完全な例）。
**スライドデッキ PDF** (エクスポートされたプレゼンテーション):
→ どのページも主に視覚的なものです。オンデマンドで個々のページをラスタライズします。
テキスト抽出では箇条書きテキストが得られますが、レイアウトはすべて失われます。
**フォームの多いドキュメント**:
→ 最初にプログラムでフォームフィールド値を抽出します (以下を参照)。ラスタライズ
必要に応じて、視覚的なコンテキスト用のフォーム ページ。
**データ量の多いドキュメント** (表、グラフ、図):
→ テーブルには pdfplumber を使用します。チャート/図を含むページをラスタライズします。
周囲の物語のテキストを抽出します。テキストと画像の両方を考慮する
精度が重要な場合は、同じページに対して。 PDF の場合、最初の動きはモデルではありません。診断です。このスキルは pdffont を実行し、ファイルにテキスト レイヤーが含まれているかどうかを確認します。 2 つのケースが続き、それぞれ異なるツールが使用されます。
ボーンデジタル PDF はその中にテキストを保持します。 pdftotext は、モデルを関与させずに、ミリ秒単位でテキストを取り出します。登場人物たちはずっとそこに座っていました。
スキャンされた PDF はピクセルです。コピー機は言葉を平坦にして画像にしたので、取り出すものは何もありません。ここでスキルは OCR に到達し、そのツールの名前は pytesseract です。
pytesseract は、Tesseract への Python バインディングです。 HP Labs は 1985 年に Tesseract を開始しました。HP は 2005 年にそれをオープンソース化しました。

そしてGoogleはその後もそれを続けた。このエンジンは、現在 AI 製品を出荷しているほとんどの人々よりも古いものです。
# すべてのページを 300 DPI で PNG にレンダリングします (私は 150 dpi を試しましたが、もっとひどい結果になりました)
pdftoppm -r 300 -png input.pdf ページ
# OCR 1 ページ BAM!
tesseract --psm 1 --oem 1 -l eng page-01.png - --psm 1 は、認識前に方向とスクリプトの検出を実行します。 --oem 1 は LSTM エンジンを選択します。 API キーは必要なく、ネットワーク呼び出しも行わないため、ウォールド ガーデンにアクセスするためにネゴシエートするレート制限はありません。
これが重要な部分です。このスキルは、モデルに各ページを見て読み上げるように指示できた可能性があります。モデルならそれができるのです。スキャンされたページは画像であり、モデルは画像を記述することができます。
ページからきれいな文字を読み取ることは、古くから解決済みの問題です。 Tesseract は数十年前のバージョンを処理していました。高速で、ローカルで実行され、通話料金はかかりません。実行するたびに同じテキストが返されますが、場合によってはゴミが含まれます。
同じページを指すモデルは速度が遅く、トークン ($$) を消費し、ネットワークのラウンドトリップが必要で、存在しない単語を静かに発明することができます。したがって、慎重なエンジニアは狭いツールに手を伸ばします。スキルも同じ呼び出しを行います。
クロードはあなたのスキャンを読んでいるわけではありません。 1980年代のキャラクターエンジンです。
レイアウトやチャートが重要な場合、モデルはページを指すことがあります。スキルはそのページを画像にラスタライズし、モデルを表示させます。それでも分割は成立します。このモデルは文字の単純な読み取りではなく、視覚的な判断を処理します。
スキル ファイルは基本的にエンジニアの反射神経を書き留めたものです。
私はこれをゆっくりとした方法で学びました。そこで私は、保存された PDF をクリーン テキストに変換するリーダー アプリ ( readplace.com ) を作成しました。早い段階で、私はスキャンしたページに Google Vision モデルを向けて、単語を尋ねました。簡単なスキャンでは機能しましたが、高密度のスキャンでは機能しませんでした。

各ページが独自のラムダを並行して取得し、212 ページのイベントを処理するには最大 15 分かかります (AWS Lambda の最大タイムアウト設定)。多くのページが自動的にタイムアウトになり、半分空のページ (または今日の気分に応じて半分いっぱい) で戻ってくるページもありました。
実際のお金と忍耐がかかりました。
私はそのパスを削除して、「F$$$ IT」と言いました。代わりに、ローカルの同じコンテナ内で Tesseract を実行してください。同じドキュメントがほんの少しの時間で完了し、費用もかからず、空のページも残りませんでした。
Vision LLM と Tesseract を単独で使用した場合のコストを示す 31 ページの PDF の例。レンダリング DPI を 150 から 300 に上げ、LSTM エンジンを固定すると、同じファイル上で認識される単語が 21,335 から 23,719 に増加しました。この品質を実現するために、壁時計は 317 秒から 48 秒になりました。
Fayner Brack のストーリーを受信箱で入手してください
無料で Medium に参加して、このライターから最新情報を入手してください。
より速くサインインするために私を覚えておいてください
多言語化には、言語検出や言語ごとのモデルは必要ありませんでした。 Tesseract は、スクリプト ファミリごとに 1 つのモデルのスクリプト バンドルを出荷します。
# 1 回の呼び出しで、存在するスクリプトをすべて読み取ります
tesseract --psm 1 -l 'script/Latin+script/Arabic+script/HanS+script/Cyrillic' page.png - --psm 1 はページごとに OSD を実行し、リージョンごとに適切なバンドルを選択します。約 35 のスクリプト パックで 100 以上の言語をカバーします。アプリコードには検出ステップがなく、言語ごとの分岐もありません。
残念ながら、多言語にすると処理時間が大幅に長くなったので、無効にしました。問題はまた別の日にします。
それから私は逆の間違いを犯しました。 Tesseract が機能しているので、モデルを完全に切り取りました。リーダーは正しい単語で満たされていますが、形はなく、すべての見出し、リスト、列が 1 つの灰色のブロックに平坦化されていました。テッセラクトは手紙を読みます。文から見出しがわかりません。その判断は曖昧であり、人工知能の働きも曖昧である

ence モデルは得意です ;)。
そこでモデルは、認められた後、仕事のために戻ってきました。残存するエラーは確率的な種類のものでした。ページをまたぐハイフネーション (Veposi-thentry) や、V↔D や m↔rn などの置換は、誤読が依然として実際の単語であったため生き残りました。それはモデルが得意とする形状です。ピクセルからの文字認識はできません。
作業は 3 回の通話に分割され、それぞれの範囲は厳密でした。
ページごとに 1 つのチャット完了。プロンプトは、約 90% の信頼度を超える単語のみを変更し、数字と固有名詞はそのまま残します。 Lambda ガードレールは長さのデルタを 30% に制限し、数字のマルチセットを保持します。拒否された場合は、元の Tesseract テキストが変更されずに通過します。
ステージ 2 - ドキュメントの差分レビュー
ドキュメントごとに 1 回の呼び出し。ページごとの単語レベルの差分と完全にクリーンアップされたテキストを参照し、変更ごとに APPROVE、REJECT、MODIFY、または NEW を発行します。 12 ページ中 1 ページに到達した修正は、他の 11 ページで元の単語が正しいと示されると拒否されます。
ステージ 3 — Readability.js のセマンティック HTML
1 ページにつき 1 回の呼び出し。これは、 h2 、 h3 、 ul 、 ol 、 pre 、 code 、 blockquote 、および table のサニタイズされたフラグメントを出力し、ビジョン モデルが使用するであろう視覚的な手がかりの代わりとなるテキスト パターン ルールを使用します。ガードレールでは 70% の可視テキスト保持が必要です。そうでない場合、ページはステージ 2 テキストのプレーンな <p> 段落に戻ります。
インフラストラクチャは 4 つの Lambda であり、それぞれの段階に応じたサイズになっています。
Lambda セットアップの例では、ページごとの Lambda のタイムアウトが 5 分に設定され、すべてのページで動作するオーケストレーターの最大許容値が 15 分に設定されています。オーケストレーターは、最大 300 ページまで、ページごとに 1 つの Tesseract 呼び出しと 1 つのクリーンアップ呼び出しをファンアウトします。アカウントの同時実行数は 1000 なので、最悪のケースでは約 30% が使用されます。 Lambda クライアントの maxSockets はデフォルトを大きく上回る 400 に設定されています

そうでない場合は、SDK レイヤーで呼び出しをキューに入れることになる 50 個。モデルは生の読み取り値には影響しません。 Tesseractはまだそれを所有しています。
このシステムは完璧には程遠いですが (決して完璧ではありません)、重要なのはそこではありません。
このパターンはクロードにとって特別なものではありません。最強の AI 製品は、1 つのモデルがすべての仕事を行うわけではありません。これらは、単純で予測可能なツールの列の隣に配線されたモデルであり、それぞれが最も得意とする部分を処理します。実際に自分たちがやっていることを理解しているエンジニアによってコード化されています。
トークンを支払うと、クロードがいくつかの CLI ツールを順番に呼び出すことができます。これは、熟練したエンジニアを大規模に雇用する場合と同様です。
入力を確認し、それに適合するツールを選択し、モデルだけが処理できる部分についてはモデルを元に戻します。あらゆる問題にモデルを投入するのは簡単な習慣ですが、本物のエンジニアは、解決しようとしている問題に合わせて構築されたツールに手を伸ばします。
OK、すべてが順調で素晴らしいですが…エンジニアリングは完全に死んでいるので、これはどれも実際には重要ではありません…そうですよね? 😢
少し考えてみましょう。クロードは、壁に囲まれた庭園の中で、企業が順番に呼び出すツールの選択に対して使用料を請求するために静かに働いています。
これが気に入った場合は、まさにこの種の読書用に構築された readplace.com も気に入るかもしれません。
読んでいただきありがとうございます。フィードバックがある場合は、 LinkedIn 、 Reddit 、または Github までご連絡ください。
読者の PDF パイプラインのより詳細な技術的再構築、メトリクスと完全なアーキテクチャ (永久 WIP): https://readplace.com/blog/pdf-ocr-pipeline-tesseract-llm-hybrid?utm_source=fagnerbrack.com&utm_content=pdf-ocr-pipeline-tesseract-llm-hybrid
Tesseractの歴史と起源の日付。 https://github.com/tesseract-ocr/tesseract
Anthropic Agent Skills、パブリック リポジトリ。 https://github.com/anthropics/skills
pytesseract、Tesseract の Python ラッパー。 https://github.com/madmaze/pytesseract
AIソフトウェアエン

ソフトウェア開発コーディングのエンジニアリング --
私は知識はオープンかつ自由であるべきだと信じています。 2015 年以来、AI が教えてくれない難しいことを共有しています。私の読書システム: https://readplace.com?utm_source=m

## Original Extract

I tried to parse a CIA PDF with Claude and lost to a parser from the 1980s. Picking the right tool still matters.

Medium A PDF Parser from the 80s Beats Claude (And They Use It Internally) | by Fayner Brack | Medium Sitemap Open in app Sign up
A PDF Parser from the 80s Beats Claude (And They Use It Internally)
My journey to parse a CIA PDF , only to discover picking the right tool for the job is still a thing
A warm-toned illustration of an 1980s scanner feeding a page through and emitting clean lines of characters, with a switched-off chrome sphere sitting small and dim in the corner. Want to come back later? Save this to readplace.com .
A PDF Parser from the 80s Beats Claude (And They Use It Internally) | Reader View
Tesseract, an OCR tool from 1985, outperforms modern AI models for PDF parsing and is used internally by Claude.
When you give a PDF to an AI, you picture the model reading it. However, for certain kinds of PDFs, the actual model barely sees the text, they only call tools previously dictated by their actual software engineers.
I spent days trying to parse a badly scanned CIA PDF (don't ask me why, otherwise I'll have to kill you). I tried using the most powerful Anthropic and Google models but I ended up concluding Tesseract (a 1985 tool), does a better job and much faster by 10x.
Not just that, but I found out that Claude uses PyTesseract in their PDF parser skill under the hood. They pick that by going through a heuristic, rasterising the most difficult scanned pages and completing missing sentences with hallucinated words when streaming their tokens back.
Claude Code AI model is not better than anyone else, they just pay skilled engineers to give them the right tools to call.
Claude ships with skills, small Markdown instruction files that tell it how to handle a task. Anthropic published the format and a set of these skills in the open , so you can read them. The reading strategy part is (apparently) not open source and reads like a runbook that a tired engineer wrote at 2am.
### Choosing your reading strategy
**Text-heavy documents** (reports, articles, books):
→ Text extraction is primary. Rasterize only for specific figures or
pages where layout matters.
**Scanned documents** (`pdffonts` shows no fonts):
→ `pdftotext` will return nothing - don't run it. Rasterize pages at
150 DPI and Read them visually. For bulk text extraction, use OCR
(pytesseract after converting pages to images - see REFERENCE.md for
a complete example).
**Slide-deck PDFs** (exported presentations):
→ Every page is primarily visual. Rasterize individual pages on demand.
Text extraction gives you bullet-point text but loses all layout.
**Form-heavy documents**:
→ Extract form field values programmatically first (see below). Rasterize
the form page for visual context if needed.
**Data-heavy documents** (tables, charts, figures):
→ Use pdfplumber for tables. Rasterize pages with charts/figures.
Extract text for surrounding narrative. Consider both text AND image
for the same page when precision matters. For a PDF, the first move is not the model. It is a diagnostic. The skill runs pdffonts and checks whether the file carries a text layer. Two cases follow, and they get different tools.
A born-digital PDF holds its text inside it. pdftotext pulls that text out in milliseconds, with no model involved. The characters were sitting there the whole time.
A scanned PDF is pixels. A photocopier flattened the words into an image, and there is nothing to pull out. Here the skill reaches for OCR, and the tool it names is pytesseract .
pytesseract is the Python binding to Tesseract. HP Labs started Tesseract in 1985. HP open-sourced it in 2005, and Google kept it going after that. The engine is older than most of the people shipping AI products today.
# render every page to a PNG at 300 DPI (I tried 150dpi, shittier results)
pdftoppm -r 300 -png input.pdf page
# OCR one page BAM!
tesseract --psm 1 --oem 1 -l eng page-01.png - --psm 1 runs orientation and script detection before recognition. --oem 1 selects the LSTM engine. It needs no API key and makes no network call, so there is no rate limit to negotiate to access a walled garden.
This is the part that matters. The skill could have told the model to look at each page and read it off. The model can do that. A scanned page is an image, and the model can describe an image.
Reading clean letters off a page is an old, settled problem. Tesseract handled a version of it decades ago. It is fast, it runs locally, and it costs nothing to call. It returns the same text on every run but some garbage here and then.
A model pointed at the same page is slower, burns tokens ($$), needs a network round trip, and can quietly invent words that were not there. So a careful engineer reaches for the narrow tool. The skill makes the same call.
Claude is not the thing reading your scan. A character engine from the 1980s is.
The model does get pointed at a page sometimes, when layout or a chart matters. The skill rasterizes that page to an image and lets the model look. Even then the split holds. The model handles the visual judgment, not the plain reading of characters.
The skill file is basically an engineer’s reflex, written down.
I learned this the slow way. So I built this reader app ( readplace.com ) that turns saved PDFs into clean text. Early on, I pointed a Google Vision model at scanned pages and asked it for the words. It worked on easy scans and fell apart on a dense one, taking up to 15 minutes (AWS Lambda max timeout setting) to process 212 pages event with each page getting its own lambda in parallel. Many pages timed out on their own and some came back half empty (or half full depending on how you're feeling today).
It cost real money and patience.
I tore that path out and said "F$$$ IT", just run Tesseract instead, locally, in the same container. The same document finished in a fraction of the time, cost nothing, and left no page empty.
Example of a PDF with 31 pages showing the costs using Vision LLM vs Tesseract alone. Bumping render DPI from 150 to 300 and pinning the LSTM engine lifted recognized words from 21,335 to 23,719 on the same file. Wall clock went from 317s to 48s for that quality.
Get Fayner Brack ’s stories in your inbox
Join Medium for free to get updates from this writer.
Remember me for faster sign in
Going multilingual needed no language detection and no per-language model. Tesseract ships script bundles, one model per script family:
# one invocation reads whatever scripts are present
tesseract --psm 1 -l 'script/Latin+script/Arabic+script/HanS+script/Cyrillic' page.png - --psm 1 runs OSD per page and picks the right bundle per region. About 35 script packs cover 100+ languages. There is no detection step in the app code and no per-language branch.
Unfortunately going multi-language increased the time to process significantly, so I disabled it.. a problem for another day.
Then I made the opposite mistake. With Tesseract working, I cut the model out completely. The reader filled with correct words and no shape, every heading, list, and column flattened into one gray block. Tesseract reads letters. It does not know a heading from a sentence. That judgment is fuzzy, and fuzzy is the work an Artificial Intelligence model is good at ;).
So the model came back, for the job AFTER recognition. The residual errors were the probabilistic kind: cross-page hyphenations ( Veposi- then tory ), and substitutions like V↔D and m↔rn that survived because the misread was still a real word. Those are the shape a model is good at. Letter recognition from pixels is not.
The work split into three calls, each scoped tight.
One chat completion per page. The prompt changes a word only above about 90% confidence, and leaves digits and proper nouns alone. A Lambda guardrail caps length delta at 30% and preserves the digit multiset, and on any rejection the original Tesseract text passes through unchanged.
Stage 2 - Document Diff Review
One call per document. It sees the per-page word-level diff plus the full cleaned text, and emits APPROVE, REJECT, MODIFY, or NEW per change. A fix that landed on one page out of twelve gets rejected when eleven other pages show the original word correct.
Stage 3 — semantic HTML for Readability.js
One call per page. It emits a sanitized fragment of h2 , h3 , ul , ol , pre , code , blockquote , and table , with text-pattern rules standing in for the visual cues a vision model would have used. A guardrail requires 70% visible-text retention, or the page falls back to plain <p> paragraphs of the Stage 2 text.
The infrastructure is four Lambdas, each sized to its stage:
Example of the Lambda setup, per-page lambdas are set timeout to 5 minutes while the orchestrators that work with all pages are set to the maximum allowed of 15 minutes. The orchestrator fans out one Tesseract invocation and one cleanup call per page, up to 300 pages. Account concurrency is 1000, so the worst case uses about 30%. The Lambda client’s maxSockets is set to 400, well above the default of 50 that would otherwise queue invocations at the SDK layer. The model does not touch the raw reading. Tesseract still owns that.
The system is far from perfect (it will never be), but that's not the point:
This pattern is not special to Claude. The strongest AI products are not one model doing the whole job. They are a model wired next to a row of plain, predictable tools, each handling the part it does best. Coded by engineers that actually know the fuck they're doing.
You pay tokens to have Claude calling some CLI tools in sequence, like hiring a skilled engineer at scale.
Check the input, pick the tool that fits it, and keep the model back for the part only a model can handle. Throwing a model at every problem is an easy habit but a real engineer reaches for the tool built for the exact problem it's trying to solve.
Ok, all is good and great but… engineering is soo dead, so none of this really matters.. right? 😢
Think for a moment: Claude is quietly working inside a walled garden to charge usage premium for companies for a selection of tools to call in sequence.
If you liked this, you might like readplace.com , built for exactly this kind of reading.
Thanks for reading. If you have some feedback, reach out to me on LinkedIn , Reddit or Github .
A more detailed technical rebuild of my reader’s PDF pipeline, with the metrics and the full architecture (forever WIP): https://readplace.com/blog/pdf-ocr-pipeline-tesseract-llm-hybrid?utm_source=fagnerbrack.com&utm_content=pdf-ocr-pipeline-tesseract-llm-hybrid
Tesseract history and origin dates. https://github.com/tesseract-ocr/tesseract
Anthropic Agent Skills, public repository. https://github.com/anthropics/skills
pytesseract, the Python wrapper for Tesseract. https://github.com/madmaze/pytesseract
AI Software Engineering Software Development Coding --
I believe knowledge should be open and free(dom). Since 2015, sharing challenging stuff AI won't tell you. My reading system: https://readplace.com?utm_source=m
