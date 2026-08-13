---
source: "https://blog.lysk.tech/markleft-ai-markdown-review/"
hn_url: "https://news.ycombinator.com/item?id=49284329"
title: "Show HN: Markleft – how I review Claude's Markdown plans"
article_title: "Markleft: How I review claude plans in Markdown | Martin Lysk"
author: "mlysk"
captured_at: "2026-08-13T11:39:25Z"
capture_tool: "hn-digest"
hn_id: 49284329
score: 4
comments: 0
posted_at: "2026-08-13T11:19:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Markleft – how I review Claude's Markdown plans

- HN: [49284329](https://news.ycombinator.com/item?id=49284329)
- Source: [blog.lysk.tech](https://blog.lysk.tech/markleft-ai-markdown-review/)
- Score: 4
- Comments: 0
- Posted: 2026-08-13T11:19:01Z

## Translation

タイトル: Show HN: Markleft – クロードのマークダウン計画をどのようにレビューするか
記事のタイトル: Markleft: Markdown でクロード計画をレビューする方法 |マーティン・リスク
説明: AI 出力が最初のプロンプトの意図を満たしていることはほとんどないため、反復が必要です。現在、反復とは通常、フォローアップ プロンプトでフィードバックを提供することを意味し、ドキュメント全体の書き直しにつながります。改訂の背後にある意図を説明するプロンプトはチャットに埋もれたままです。

記事本文:
メイン コンテンツにスキップ Martin Lysk ブログ GitHub すべての投稿 2026
Markleft: Markdown でクロード計画をレビューする方法
Git 上の SQLite、パート II: Zlib のあまり知られていない機能のロックを解除する
Git 上の SQLite、パート I: .git フォルダー - ウサギの穴に落ちる
Git 上の SQLite、プロローグ: なぜ git でランダム アクセスが必要なのか
Markleft: Markdown でクロード計画をレビューする方法
AI の出力が最初のプロンプトの意図を満たしていることはほとんどないため、反復が必要です。現在、反復とは通常、フォローアップ プロンプトでフィードバックを提供することを意味し、ドキュメント全体の書き直しにつながります。改訂の背後にある意図を説明するプロンプトはチャットに埋もれたままです。
同僚に 5 ページの文書を送ることを想像してください。彼らはそれを読んで、いくつかの懸念を抱いています。ワークフローは次のようになりません。
あなたが嫌いなものすべてについて散文的な説明をメールで送ってください
完全に書き直した文書をお送りします
次に、2 つのファイルを比較して、私があなたの言うことを理解したかどうかを判断します
それは不合理です。このワークフローは、Word のコメントと変更の追跡、または Google ドキュメントの提案モードによって何十年も解決されてきました。
実際に動作しているところを見てみたいですか?スポイラー
Markdown での AI の操作とその問題点
ドキュメントの作成プロセスを試してみましょう。一度繰り返してみましょう。
Markdown で自転車に乗ったペリカンに関するちょっとしたストーリーを書き、ペリカンの小さな SVG イラストを追加してください。
### かわいそうなペリカン
ペリカンのペリーは明るいものを見つけました
桟橋に立てかけられた赤い自転車、そして、
好奇心旺盛に一瞥した後、
それはまさに冒険の類だと判断した
朝が必要です。
彼はよろよろと漁船の前を通り過ぎた、
くちばしで小さな鐘を鳴らし、
そしてビーチに転がり出して、ちょうど朝食の時間に間に合いました。
![ペリカン](./pelican-on-a-bycicle-v1.svg)
レンダリング済み
かわいそうなペリカン
ペリカンのペリーは、桟橋に立てかけられた真っ赤な自転車を見つけました。

好奇心をそそられて一瞥した後、これはまさに朝に必要な種類の冒険であると判断しました。彼はよろよろと漁船の前を通り過ぎ、くちばしで小さな鐘を鳴らし、ビーチに転がり落ちました。ちょうど朝食の時間でした。
ステップ 1: 批評を収集して定式化する
フィードバックを有益な指示に変えることは、すでに精度の問題です。
見出しは間違った記事を約束しています。ダッシュを削除し、2 番目の文を短くし、ペリカンの頭を修正し、左足をペダルに乗せて、鳥をよりスポーティに見せます。
それは具体的ですね。そうではない。
「見出し」が機能するのは、このおもちゃの文書に明らかな見出しが 1 つあるからです。長い記事では、変更したい見出しを説明する必要があります。
「 2 番目の文」はより具体的ですが、レビュー担当者と AI の両方が同じ方法で文をカウントすることを期待して文をカウントします。
「左足」はさらに悪いです。散文では、幻想を打ち破った SVG 内の正確な形状を示すことができません。 Figma などのツールを使用すると、デザイン画面に直接コメントを追加できるのはこのためです。
次に、繰り返しの見出し、ネストされた表、図、コードを含む複数ページのドキュメントに対して同じ種類のフィードバックを提供してみます。
ネットワーク アーキテクチャを表示する 3 番目の段落の 2 番目の画像を更新してください。バックアップ サーバーを読み取り専用モードで、右上の 2 番目のデータベース サーバーの隣に移動します。
より大きなドキュメントでは、最初にフィードバックを収集するという別の問題が発生します。パラグラフ 1 を読んで「意味がわからない」と思う人もいるかもしれませんが、パラグラフ 3 でそれが説明されているかもしれません。そのような考えが浮かんだとき、私は読むのをやめてプロンプトを作成したくありません。足跡を残して続きを読み、後で戻ってきたいと思います。
ステップ 2: 変化を理解する
いくつかのアプローチが役立ちますが、共通する点は次のとおりです。

akness: 作成者が提案された変更をコンテキスト内でレビューする前に、AI が新しいドキュメントを作成します。
私たちはまだ最初のイテレーションの途中です。 LLM によって更新が実行されたからといって、フィードバックが考慮されたことや意図が満たされたことを意味するものではありません。残るタスクは 2 つです。
変化を見つけてください。文書全体で「間違い探し」を 1 回行います。 Git は差分を表示できますが、レンダリングされた Mermaid ダイアグラム、テーブル、SVG の背後に隠された変更を把握するのははるかに困難です。
意図を回復します。何が変更されたのかがわかったら、その変更が元のプロンプトのどの部分に対処することを意図していたかを思い出し、その結果が実際にフィードバックを満たすかどうかを判断する必要があります。
より具体的なプロンプトを思い出してください。
AIによる手直しで最初の一文が抜けてしまったらどうなるでしょうか？このプロンプトをインテントにマッピングすることは、次のことを意味します。
変更前の文 2 は何でしたか?
これは短くするという意図を反映しているのでしょうか？
その文が短くなったかどうかだけを気にする必要があります。
アイデア — Markdown の提案モード
必要とされているのは、文書の各部分の位置を曖昧に説明するのではなく、直接注釈を付ける方法と、それらの注釈に対して提案された変更をレビューする方法です。
これはまさに Markleft が提供するものです。これは、次の 3 つの主要コンポーネントに基づいています。
人間が基礎となる形式を考慮せずに コメント を作成および編集し、変更を提案し、提案を適用できる WYSIWYG Markdown エディター。
AI によって利用可能な Markdown 互換の注釈仕様。これにより、ユーザーはテキスト、コード、表、人魚図、画像、SVG にコメントしたり、提案を提案したりできます。
コメントへの対処方法を AI に指示し、 Markleft にのみ提案を追加するように指示するプロンプト。
AI はブックマークをキックするとマークダウンし、発言がある場合はマークアップとしてレビューします。

- メモを残してマークダウンを保存すると、マークが残っています。それを AI に渡し、提案が返されたら、マークが正しく取得されたかどうかを確認できます。
Markdown はドキュメント形式のままです。
この仕様では、通常の Markdown 構造を使用して提案モードを有効にします。
Markleft エディタはローカルの Markdown ファイルで動作します。 Chrome ではブックマークレットとして実行されるため、Markdown ファイルを開いてブックマークレットをアクティブ化し、コメントや提案を含むフル機能のエディターを取得できます。
Markdown ファイルを読み取り、脚注を解析して Markleft の注釈、コメント、提案を作成し、レンダリングします。
ユーザーが選択範囲にコメントしたり、画像内にマーカーを配置したりすると、エディターは対応する Markleft 脚注を挿入します。ユーザーがファイルを保存すると、エディターは AI に対して Markleft を説明し、コメントに対処する提案を作成する方法を説明するプロンプトをコンパイルします。
AI は提案を作成し、Markdown ファイルに追加します。エディターはこれらの変更を検出し、対応するコメントへのリンクを含む新しい提案を表示します。テキストは、対応するレンダリングされた要素内で差分されます。リスト項目とテーブルのセルはテキストが比較される前にペアにされるため、構造全体で無意味な差分が 1 つ発生することがなくなります。画像のみの置換は前後のスライダーになります。
注釈は、コメントと同様に、コメントの内容を識別するアンカーとコメント自体の 2 つの部分で構成されます。 Markdown には、ほとんどの Markdown レンダラーがサポートする脚注の概念があります。アンカーは [^id-of-the-footnote] の形式を使用しますが、その定義は [^id-of-the-footnote]: body の形式で別の行に表示されます。選択した単語や画像内の X/Y 座標などの追加情報をエンコードするには、脚注 ID 自体のスキーマを使用します。

この文にはそれほど儀式は必要ありません。[^range-prev-12-chars-14824-a1b2]
[ ^range-prev-12-chars-14824-a1b2 ] : これをより直接的にします。
range-prev-12-chars は、注釈が前の 12 個の表示される空白以外の文字をカバーすることを示します。残りのコンポーネントは ID とコンテンツのフィンガープリントを提供するため、Markleft はアンカーが古くなったときを検出できます。
他の ID は他の種類のアンカーをエンコードします。
image-X-Y-* には、正規化された画像座標が格納されます。
code-line-L-col-C-len-N-* はコード範囲を識別します。
block-* は、それを含むブロックをアドレス指定します。
comment-* は、別のコメントへの返信を表します。
コメントは脚注であるため、Markdown ツールは Markleft を理解できない場合でもコメントを保持します。 GitHub などの通常のレンダラにとって、これは単なる脚注です。 AI と Markleft エディターにとって、画像上の点、またはテキスト ブロック内の強調表示された文が識別されます。
提案が確実にブロックをターゲットにできるようにするには、安定した識別子が必要です。 Markleft は、実際のドキュメント ブロックの直前に HTML コメントを挿入します。
< !- - マークレフト: ブロック ID = " babf825b " -- >
提案は参照されておらず、追加のみの脚注です
提案は、予約された ID を持つ脚注定義であり、元の本文に意図的にインライン脚注アンカーを含めないことです。
[ ^suggestion-s2-update-block-babf825b ] : マークダウンの置換
ID には次のように記載されています。
ターゲットはブロック babf825b です。
前に挿入、後に挿入、および削除の操作は同じパターンを使用します。提案本文の最後の行には、提案で取り上げられるコメントの脚注アンカーが含まれます。その行はメタデータであり、提案されたコンテンツの一部ではありません。
これは重要な追加専用プロパティです。AI は、レビューしているドキュメントを変更する許可を取得せずに提案を追加できます。
ブックマーク バーを表示します (macOS では ⌘⇧B、Windows/Linu では Ctrl+Shift+B)

x) クロムの場合
このリンク「Open in M←」を Chrome のブックマーク バーにドラッグします。
Chrome でローカルの Markdown ファイルを開き、ブックマークをクリックします
ソース コードと完全なインストール手順はここで見つけてください: https://github.com/martin-lysk/markleft
Markdown での AI の使用とその問題点
ステップ 1: 批評を収集して定式化する
ステップ 2: 変化を理解する
アイデア - Markdown の提案モード
Markleft Markleft - 編集者

## Original Extract

AI output seldom meets the intent on the first prompt, so iteration is necessary. Today, iteration usually means providing feedback in a follow-up prompt, which leads to a rewrite of the whole document. The prompt that explains the intent behind the revision remains buried in the chat.

Skip to main content Martin Lysk Blog GitHub All posts 2026
Markleft: How I review claude plans in Markdown
SQLite on Git, Part II: Unlocking Zlib's less known Feature
SQLite on Git, Part I: The .git folder - Falling down the Rabbithole
SQLite on Git, Prologue: Why do we need random access in git
Markleft: How I review claude plans in Markdown
AI output seldom meets the intent on the first prompt, so iteration is necessary. Today, iteration usually means providing feedback in a follow-up prompt, which leads to a rewrite of the whole document. The prompt that explains the intent behind the revision remains buried in the chat.
Imagine sending a colleague a five page document. They read it and have some concerns. The workflow would not be:
Email me a prose description of everything you dislike
I send you a completely rewritten document
Then compare the two files and figure out whether I understood you
That is absurd—this workflow has been solved for decades by comments and tracked changes in Word, or by Suggesting mode in Google Docs.
Want to see it in action? spoiler
Working with AI in Markdown and where it breaks ​
Let's play through the process of creating a document—and iterate once.
Please write a little story about a pelican on a bike in Markdown, and add a small SVG illustration of the pelican.
### A Poor Pelican
Perry the pelican found a bright
red bicycle leaning against the pier—and,
after one curious glance,
decided it was exactly the sort of adventure
the morning required.
He wobbled past the fishing boats,
rang the tiny bell with his beak,
and rolled onto the beach—just in time for breakfast.
![pelican](./pelican-on-a-bycicle-v1.svg)
Rendered
A Poor Pelican ​
Perry the pelican found a bright red bicycle leaning against the pier—and, after one curious glance, decided it was exactly the sort of adventure the morning required. He wobbled past the fishing boats, rang the tiny bell with his beak, and rolled onto the beach—just in time for breakfast.
Step one: Collecting and formulating the critique ​
Turning your feedback into useful instructions is already a precision problem.
The headline promises the wrong story. Remove the em dashes, shorten the second sentence , fix the pelican's head, put the left foot on a pedal, and make the bird look sportier.
That sounds specific. It is not.
“ The headline ” works only because this toy document has one obvious headline. In a longer article you need to describe the headline you want to change.
“ The second sentence ” is way more specific, but it makes both the reviewer and the AI count sentences, hoping both count them the same way.
“The left foot” is worse: prose cannot point to the exact shape in the SVG that broke the illusion. That is why tools such as Figma allow comments directly on the design surface.
Now try giving the same kind of feedback on a multi-page document with repeated headings, nested tables, diagrams, and code.
Please update the second image in the third paragraph that displays the network architecture. Move the backup server to the top right, next to the second database server, in read-only mode.
In larger documents, another problem emerges: collecting the feedback in the first place. While reading paragraph 1, one may think, “This doesn't make sense,” but maybe paragraph 3 explains it. When such a thought comes up, I don't want to stop reading and formulate a prompt. I want to leave a mark, continue reading, and come back later.
Step two: Understand the change ​
Several approaches can help, but they share the same weakness: the AI produces a new document before the author can review the proposed changes in context.
We are still in the first iteration. An update executed by an LLM does not mean that our feedback was taken into account or that our intent was met. We are left with two tasks:
Find the change. One round of “spot the difference” across the whole document. Git can show diffs, but changes hidden behind rendered Mermaid diagrams, tables, and SVGs are much harder to grasp.
Recover the intent. Once we know what changed, we still have to remember which part of our original prompt the change was meant to address and decide whether the result actually satisfies the feedback.
Remember that more specific prompt:
What if the first sentence is dropped in the rework by the AI? Mapping this prompt to the intent means:
What was sentence 2 before the change?
Does this reflect the intent of being shorter?
We should only care whether that sentence is now shorter.
The Idea—Suggestion Mode for Markdown ​
What is needed is a way to annotate parts of the document directly instead of describing their location vaguely, along with a way to review proposed changes against those annotations.
This is exactly what Markleft provides. It is based on three main components:
A WYSIWYG Markdown editor for humans to create and edit comments , propose changes, and apply suggestions without thinking about the underlying format.
A Markdown-compatible annotation spec consumable by AI, allowing users to comment on text, code, tables, Mermaid diagrams, images, and SVGs, and to propose suggestions .
A prompt that tells the AI how to address the comments and instructs it to append suggestions only in Markleft .
AI gives you a Mark down you kick of a bookmark , review it as a Mark up , when you have remarks - you leave a note and save the Markdown than now has a Mark left , you pass it to AI and when it comes back with Suggestions you can check if it got your Mark righ.
Markdown remains the document format.
The spec uses ordinary Markdown constructs to enable suggestion mode.
The Markleft editor operates on local Markdown files. It runs as a bookmarklet in Chrome, so you can open a Markdown file, activate the bookmarklet, and get a full-featured editor with comments and suggestions.
It reads the Markdown file, parses its footnotes into Markleft annotations, comments, and suggestions, and renders them.
If a user comments on a selection or places a marker inside an image, the editor injects a corresponding Markleft footnote. When the user saves the file, the editor compiles a prompt that explains Markleft to the AI and describes how to compose suggestions that address the comments.
The AI composes suggestions and appends them to the Markdown file. The editor detects those changes and renders the new suggestions with links to the comments they address. Text is diffed within corresponding rendered elements. List items and table cells are paired before their text is compared, avoiding one meaningless diff across an entire structure. Image-only replacements become a before-and-after slider.
An annotation—like a comment—consists of two parts: an anchor that identifies what it comments on and the comment itself. Markdown has a concept of footnotes that most Markdown renderers support. An anchor uses the format [^id-of-the-footnote] , while its definition appears on a separate line in the format [^id-of-the-footnote]: body of the footnote . To encode additional information—such as selected words or x/y coordinates inside an image—we use a schema in the footnote ID itself.
This sentence needs less ceremony.[^range-prev-12-chars-14824-a1b2]
[ ^range-prev-12-chars-14824-a1b2 ] : Make this more direct.
range-prev-12-chars says that the annotation covers the previous twelve visible, non-whitespace characters. The remaining components provide identity and a content fingerprint so Markleft can detect when an anchor has become stale.
Other IDs encode other kinds of anchors:
image-X-Y-* stores normalized image coordinates.
code-line-L-col-C-len-N-* identifies a code range.
block-* addresses the containing block.
comment-* represents a reply to another comment.
Because the comment is a footnote, Markdown tools preserve it even when they do not understand Markleft. To a normal renderer such as GitHub, it is just a footnote. To the AI and the Markleft editor, it identifies a point on an image or a highlighted sentence inside a text block.
To allow suggestions to target blocks reliably, we need stable identifiers. Markleft injects an HTML comment immediately before each real document block:
< !- - markleft: block id = " babf825b " -- >
Suggestions are unreferenced, append-only footnotes ​
A suggestion is a footnote definition with a reserved ID and intentionally no inline footnote anchor in the original body.
[ ^suggestion-s2-update-block-babf825b ] : replacement Markdown
The ID says:
the target is block babf825b .
Insert-before, insert-after, and delete operations use the same pattern. The last line of a suggestion body contains footnote anchors for the comments the suggestion addresses; that line is metadata, not part of the proposed content.
This is the crucial append-only property: an AI can add a proposal without receiving permission to alter the document it is reviewing.
Show the bookmarks bar (⌘⇧B on macOS or Ctrl+Shift+B on Windows/Linux) in chrome
Drag this link: Open in M← into chromes bookmark bar.
Open a local Markdown file in chrome and click the bookmark
Find the source code and a full installation instruction here: https://github.com/martin-lysk/markleft
Working with AI in Markdown and where it breaks
Step one: Collecting and formulating the critique
Step two: Understand the change
The Idea—Suggestion Mode for Markdown
Markleft Markleft - the editor
