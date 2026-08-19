---
source: "https://lukasvonkunhardt.com/obsidian-claude-code-setup/"
hn_url: "https://news.ycombinator.com/item?id=49366812"
title: "Obsidian + Claude Code: The Setup I Use Every Day"
article_title: "Obsidian + Claude Code: The Setup I Use Every Day — Lukas von Kunhardt"
image: "https://lukasvonkunhardt.com/og/obsidian-claude-code-setup.png"
author: "lukasvk"
captured_at: "2026-08-19T21:17:46Z"
capture_tool: "hn-digest"
hn_id: 49366812
score: 2
comments: 0
posted_at: "2026-08-19T20:30:04Z"
tags:
  - hacker-news
  - translated
---

# Obsidian + Claude Code: The Setup I Use Every Day

- HN: [49366812](https://news.ycombinator.com/item?id=49366812)
- Source: [lukasvonkunhardt.com](https://lukasvonkunhardt.com/obsidian-claude-code-setup/)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T20:30:04Z

## Translation

タイトル: Obsidian + Claude コード: 私が毎日使用するセットアップ
記事のタイトル: Obsidian + Claude Code: 私が毎日使用するセットアップ — Lukas von Kunhardt
説明: Obsidian をクロード コードの隣のマークダウン エディターおよびプロジェクト オーガナイザーとして使用する方法。プラグインも MCP サーバーもなし、プロジェクトごとに 1 つのフォルダー。

記事本文:
Lukas von Kunhardt Obsidian + Claude Code について: 私が毎日使用するセットアップ
実際のセットアップはどのようなものですか?
クロード コードは私のメモをどのように認識しますか?
Obsidian Claude Code プラグインまたは MCP サーバーが必要ですか?
なぜ Notion や Word ではなく Markdown を使うのでしょうか?
通常のクロード チャットと比べて実際の利点は何ですか?
これではクロードに閉じ込められませんか？
データベースはマークダウン ファイルよりも優れているのではないでしょうか?
クロードコードをどのフォルダーで開きますか?
ボールトに CLAUDE.md が必要ですか?
あなたとクロードがお互いを上書きせずに同じノートを編集するにはどうすればよいですか?
実際にどの Obsidian プラグインを使用すればよいですか?
これは電話ではどう見えますか?
開発者である必要がありますか?
チャット ウィンドウでクロードとワークアウトした内容は、チャットの中に埋もれてしまいます。何も溜まらない。代わりに、Obsidian の保管庫にクロード コードを指定すると、同じ作業が自分のコンピュータ上のプレーン ファイルに保存され、そこに留まり、複合化することができます。
これが機能するのは、関連情報をファイルとして外部化し、両者が簡単にアクセスできるためです。あなたもそれを読むことができますし、クロードも読むことができます。蓄積されるのは共有記憶です。どちらも間に抽象化レイヤーを必要としません。これは思っている以上に重要です (これについては後で詳しく説明します)。
画面上でどのように見えるか、Vault をどのように整理するか、どこでうまく機能し、どこでうまくいかないかなど、その仕組みを説明します。特別なプラグインや MCP サーバーは必要なく、開発者である必要もありません。 Claude または他の言語モデル プロバイダーへのサブスクリプション以外に料金を支払う必要はありません。
実際のセットアップはどのようなものですか?
私の日常のセットアップは次のようになります。Obsidian を開いて、Claude Code を使用しています。十分な大きさの画面がある場合は、それらを同時に表示したいと思います。左側にクロードを開いてあります。私もディクテーションをよく使います

クロードと話すか (詳しく知りたい場合は、私の Mac 用の最高のディクテーション ツールの記事を読んでください)、または単に wisprflow などを使用してください。そして右側には金庫があります。
私は、Claude が書いているものの出力を表示するために Obsidian を使用しています。このアイデアは、Obsidian がマークダウン ビューアおよびエディタであるということです。これにより、クロードと一種のコラボレーションが可能になり、Google ドキュメントとほぼ同じように作業できるようになります。
仕事、個人、何でも、執筆など、あらゆる種類のプロジェクトにこれを使用します。プロジェクト自体はテキストである必要はありません。それについてのあなたの考えもそうですし、クロードとの会話も同様です。したがって、習慣は単純です。私がプロジェクトについて何かを解決するたびに、それはそのプロジェクトのフォルダー内のメモに保存され、クロードが保存する価値のある何かを作成するたびに、それは同じフォルダーに保存されます。それが私と同じプロジェクトのイメージをクロードに与えているのです。そして、マークダウン メモのフォルダーは、私たちのどちらよりも情報を保持するのに優れています。私は 3 月に決めたことを忘れてしまい、クロードは最後に話した内容を何も覚えていません。
クロード コードは私のメモをどのように認識しますか?
まず、クロード コードの概念と、クロード チャットとの違いを理解する必要があります。 Claude Code では、フォルダーを開くことから始めます。クロード コードをどこで使用する場合でも、ターミナル バージョンでもアプリ バージョンでも、ファイルの場所を指定するように求められます。これが開始点です。
これまでこれを使用したことがなく、最初から開始する場合は、まず Obsidian を開いてから、Vault を作成します。どこかの場所を指定するか、デフォルトの場所を選択するように求められ、フォルダーが作成されるだけです。 Obsidian にはサイドバーがあり、サイドバーにはネストされたフォルダーが表示され、フォルダー内には

メモ。これは、コンピュータ上に存在するファイル構造を正確に反映しています。 Obsidian でフォルダーとして表示されたり、うまく書かれてレンダリングされたメモは、コンピューター上の単なるフォルダーであり、その中にマークダウン ファイルが含まれています。これが魔法の核心です。クロード コードをコンピューター上のフォルダーに指定すると、表示されているのとまったく同じファイルにアクセスして表示できるようになります。
何かプロジェクトを始めたいと思ったとします。それは、研究プロジェクトや仕事のための新しいプロジェクトなど、より複雑なものになるでしょう。 Vault に新しいフォルダーを作成し、このフォルダーに、このプロジェクトに関連するメモを収集します。そして嬉しいのは、Claude Code に切り替えると、同じフォルダーの場所を指定できるようになり、書いている内容がすべて表示されることです。また、自分の文章を編集して変更を提案することもできます。
クロード コードで選択したモードに応じて、各変更を手動で承認して、メモに対するクロードの動作を完全に制御することも、より手動で行う場合はクロードに完全な制御を与えることもできます。
Obsidian Claude Code プラグインまたは MCP サーバーが必要ですか?
あなたがコンピューターで作業していて、そこにクロードがいて、そこにオブシディアンがいる場合は、他に何もする必要はありません。 Obsidian はマシン上のファイルを参照するだけであり、Claude Code は設計上、ネイティブにマシン上のファイルを操作することになっているためです。まさに天国のような試合だ。これらは互いに直接的に補完的です。
提供されているものは 3 つのグループに分類されており、どれを検討しているのかを知るのに役立ちます。
最初のグループは、Obsidian サイドバー内でクロード コードを実行します。要点がわかりません。 2 つのウィンドウを開くには特別なアプリは必要ありません。
2 番目のグループはエディター コンテキスト ブリです。

dges は、どのノートを開いているのか、どのテキストを選択しているのかをクロードに伝えます。これは少なくとも実際のギャップですが、クロードにどのファイルのことを言っているのかを伝えるだけで十分です。1 つのプロジェクト フォルダーで作業するという慣例に従えば、それはあまり重要ではなくなります。必要なファイルがコンピュータ全体に分散している場合は、より価値があります。
3 番目のグループは、コンテナーを他のアプリに公開する MCP サーバーであり、これについては明確にしておく価値があります。クロード コードはファイルを読み取ることができます。チャット モードの Claude デスクトップ アプリはできないため、MCP サーバーを Obsidian に接続してこれを修正しています。しかし、それはさらに悪い方向への遠回りです。自分のファイル上にある Obsidian で Claude Code を使用するだけです。他の方法では、このセットアップ全体が構築されている利点のほとんどを放棄することになります。
実際に使用するケースが1つあります。サーバー上でクロード コードを実行し、Web サイトのパフォーマンスに関する更新をプロジェクトの 1 つのメモに投稿したため、ローカル REST API プラグインを有効にしました。どこかのサーバー上にサイトがあり、そのプロジェクトのフォルダーに週次レポートを配置したい場合、それは本物です。コンピューターの外部にある何かが保管庫に書き込まれる必要があります。携帯電話からボールトを編集する場合も、同じ問題が発生します。この記事で取り上げている問題とは別の問題です。
なぜ Notion や Word ではなく Markdown を使うのでしょうか?
比較してみましょう。 Notion MCP サーバーで Notion を使用する場合、少なくともユーザー側では、同様のセットアップが行われるでしょう。 Notion の操作方法とそこでの記述方法はすでに知っているので、対話の半分はそれほど変わりません。しかし、クロードの場合、状況は大きく変わります。クロードは Notion ドキュメントをフォルダー内のファイルとして認識しないからです。代わりに、Notion と対話するには MCP サーバーまたは API を経由する必要があります。

これは、自分のデバイスに読みやすい形式でメモを保存しないメモ作成ツールに当てはまります。
それにはいくつかの欠点があります。このモデルは、はるかに多くのトークンを費やし、ファイルに到達する方法について考える必要がありますが、これは不必要に困難です。抽象化レイヤーを導入するたびに、作業が難しくなります。
モデルがコンピュータ上のどこかでテキストを検索したい場合、「grep」のようなものを使用できます。 grep は、約 50 年間、すべての Mac および Linux マシンに搭載されている検索ツールです。名前だけでなくファイルの内部も調べ、1 秒以内に何千ものファイルを調べます。
それはとても道が開けます。 Claude は、クライアント フォルダ内で 1 つのサプライヤーに言及するすべてのメモを、何百ものサブフォルダにまたがって見つけることができます。どのメモが入っていたかを覚えていなくても、3 月以降に触っていないすべてのメモを一覧表示できます。一度に 2 つのことについて言及しているメモを見つけることができます。これは、検索ボックスでは決して答えることができない質問です。一定の長さを超えるすべてのファイルを要求したり、規則に合わせて 50 個の音符の名前を変更したりすることができます。
プラグイン、インストール、または特別なツールは必要ありません。これらは grep と同じ世代の基本的なコマンド ライン コマンドであり、モデルではトレーニング データで他のほとんどのコマンドよりも多く使用されています。英語が流暢であるのと同じように、彼らも流暢です。それに比べて、Notion MCP サーバーは 2 年前には存在せず、使い方が複雑なので、モデルは実際の作業と同時にツールを実行する方法を考え出す必要があります。ファイルは、Claude コードが実行されるのと同じフォルダーにあります。
これは、これが私だけの好みではなく、モデル自身の好みであるという証拠です。クロード・コードに何かを書き留めてもらい、後で使用できるように保存し、参照文書として保管してください。

どのような種類のファイルが作成されると思いますか?マークダウン。やり方を全く変える必要はありません。
Word を使って作業することもできます。 Word は、多くの人がよく知っている一般的なドキュメント エディターです。ただし、Word ファイルは独自の形式であり、Word ファイルにとってのマークダウン ファイルは、スプレッドシートにとってのカンマ区切り値ファイルと同じものです。 CSV は、言語モデルにとって Excel ファイルよりもはるかに読みやすいです。クロードやその他の言語モデルが、これらのより複雑なファイル形式を処理する場合、出力品質が遅れていることで悪名高いのはそのためです。クロードに Word ファイルを編集させようとしている場合は、マークダウン ファイルで直接作業する場合と比べてパフォーマンスが劣ることに気づくでしょう。
Obsidian は本当に楽しいので、あなたにとっても非常に便利です。見た目も良く、必要な機能はほとんど備わっています。 Notion より少し劣り、Word よりも少し劣りますが、良い意味での最小公倍数です。どちらにとっても非常に使いやすいです。カスタマイズ性が高く、非常に肉付けされたメモ アプリです。何もインストールする必要がなく、すぐに使える PDF が表示されます。他のほとんどのファイル タイプにはプラグインがあります。
カンマ区切りテーブル用の「CSVエディター」
スプレッドシート用の「Sheet Plus」または Excel
PDF をただ読むだけでなく強調表示したい場合は、「Annotator」を使用します。
通常のクロード チャットと比べて実際の利点は何ですか?
あなたの労働の成果、つまりクロードとの対話の結果は、実際にはどこに保管されるのでしょうか?基本的にクロードを、特異な質問に答えるための Google の改良版として使用している場合、得られる成果の半分も得られていません。これがクロードとオブシディアンの真の利点です。長期的なプロジェクトに取り組み、共有知識を構築できるのです。

あなたとクロードが対話できるベース。これはあらゆる複雑なタスクに必要になります。
こちらが本物です。昨年の 11 月、私はほとんど知識のなかった分野の大規模なエンタープライズ プロジェクトに参加しました。最初の数週間で膨大な量の資料が送られてきましたし、覚えなければならない新しい人もたくさんいました。会議では、説明された内容のほとんどが理解できませんでした。彼らは私がまだ知らない 15 人の人の名前を出しましたが、時には 2 時間前の会議のときと同じ人もいたのですが、私はまだその人のことを覚えていませんでした。 9 か月後、その 1 つのフォルダーには 100 件のトランスクリプトと約 300 件の Wiki ページが保存されました。
クロードはそのフォルダー内のすべてを Wiki に整理し、それがずっと早く意味を理解し始めました。いつも愚かな質問をする男ではなく、そこで質問することができました。私は「これらすべてのファイルが何なのか、そして私がよく聞くこの用語が何を意味するのか説明してください」に 1 日費やすことができました。すると、その答えが領収書とともに戻ってきました。あなたはその会議でそれについて話し合ったし、5 日前に彼らが送った文書の 90 ページにそれが説明されていました。彼らは私に、誰もが入り込むのが難しいと知っていた分野で私が物事を素早く理解し始めたことは驚くべきであると感じたというフィードバックをくれました。
書き写す習慣もつきました

[切り捨てられた]

## Original Extract

How I use Obsidian as the markdown editor and project organizer next to Claude Code. No plugin, no MCP server, one folder per project.

Lukas von Kunhardt About Obsidian + Claude Code: The Setup I Use Every Day
What does the setup actually look like?
How does Claude Code see my notes?
Do you need the Obsidian Claude Code plugin or an MCP server?
Why markdown, and not Notion or Word?
What is the actual benefit over a normal Claude chat?
Doesn’t this lock you into Claude?
Wouldn’t a database be better than markdown files?
Which folder do you open Claude Code in?
Do you need a CLAUDE.md in your vault?
How do you and Claude edit the same note without overwriting each other?
Which Obsidian plugins do I actually use?
What does this look like on the phone?
Do you have to be a developer?
What you work out with Claude in a chat window gets buried in that chat. Nothing accumulates. Point Claude Code at your Obsidian vault instead, and the same work lands in plain files on your own computer, where it can stay and compound.
This works because relevant infromation can be externalized to files which both of you can easily access. You can read it, and so can Claude. What accumulates is a shared memory. Neither of you needs an abstraction layer in between, and that matters more than it sounds (more on this later).
I’m going to walk you through how it works: what it looks like on screen, how I organize the vault, and where it works well and where it doesn’t. You do not need special plugins, no MCP server, and you don’t have to be a developer. You don’t need to pay for anything other than a subscription to Claude, or to any other language model provider.
What does the setup actually look like?
My day to day setup looks like this: I have Obsidian open, and Claude Code. If I have a big enough screen, I like to view them at the same time. I have Claude open on the left. I mostly use a dictation tool to talk to Claude ( read my best dictation tool for Mac article if you want to know more), or just use wisprflow or something like that. And on the right side I have my vault.
I use Obsidian for viewing the outputs of whatever Claude is writing. The idea is that Obsidian is a markdown viewer and editor. It allows me to have a sort of collaboration with Claude, where I can almost work in it like in Google Docs.
I use this for all kinds of projects: work, personal, anything, writing. The project itself doesn’t have to be text. Your thinking about it is, and so is your conversation with Claude. So the habit is simple: every time I work something out about a project, it goes into some note in that project’s folder, and every time Claude works something out worth keeping, it goes into the same folder. That is what gives Claude the same picture of the project that I have. And the folder of markdown notes is better at retaining information than either of us. I forget what I decided in March, and Claude remembers nothing from the last chat we had.
How does Claude Code see my notes?
First of all, you have to understand the concept of Claude Code and how it’s different from the Claude chat. In Claude Code, you start by opening a folder. No matter where you use Claude Code, whether it’s the terminal version or the app version, it’s going to ask you to point it to some file location, and this is the starting point.
If you were starting from scratch and had never used this before, you would open Obsidian first, and then you would create a vault. It’s going to ask you to point it to some location or pick a default location, and it’s just going to create a folder. In Obsidian you have a sidebar, and in the sidebar you see nested folders, and in the folders you see your notes. And this mirrors exactly the file structure that exists on your computer. What you see in Obsidian as folders and nicely written and rendered notes is just folders on your computer with markdown files in them. This is the core of the magic. If you point Claude Code at some folder on your computer, it’s going to have access to, and see, the exact same files you’re seeing.
Say you wanted to start some project, and it’s going to be a more complicated one: some research project, some new project for work, whatever. You would create a new folder in your vault, and in this folder you might collect notes for anything relevant to this project. And the nice thing is: if you now switch to Claude Code, you can point it at the same folder location, and then it’s going to see everything you’re writing. And it can also edit and suggest changes to your own writing.
Depending on what mode you select in Claude Code, you can either approve each change manually, so you are in full control of what Claude does to your notes, or you can give it full control if it’s a more hands-off thing.
Do you need the Obsidian Claude Code plugin or an MCP server?
If you work on your computer and you have Claude there and you have Obsidian there, then there’s nothing else you need to do. Because Obsidian just looks at files on your machine, and Claude Code natively, by design, is supposed to work with files on your machine. They are a match made in heaven. They are directly complementary to each other.
What is on offer falls into three groups, and it helps to know which one you are looking at.
The first group runs Claude Code inside an Obsidian sidebar. I don’t see the point. Having two windows open is not something you need a special app for.
The second group are editor context bridges, which tell Claude which note you have open and what text you have selected. That is at least a real gap, but it is easy enough to just tell Claude which file you mean, and if you follow the convention of working in one project folder it stops mattering much. If the files you need are scattered across your whole computer, it is worth more.
The third group are MCP servers that expose your vault to other apps, and this is the one worth being clear about. Claude Code can read your files. The Claude desktop app in chat mode cannot, so people connect an MCP server to Obsidian to fix that. But that is the long way around to somewhere worse. Just use Claude Code on an Obsidian that sits on your own files. The other way you give up most of the advantage this whole setup is built on.
There is one case where I do use one. I have the Local REST API plugin enabled because I had Claude Code running on a server, posting updates into a note in one of my projects about how a website was performing. If you have a site on a server somewhere and you want a weekly report to land in that project’s folder, that is genuine: something outside your computer needs to write into your vault. Editing your vault from your phone is the same shape of problem. It is a different problem from the one this article is about.
Why markdown, and not Notion or Word?
Take a comparison. If you were to use Notion with a Notion MCP server, you would have a similar looking setup, at least on your side. You already know how to navigate Notion and how to write there, so your half of the interaction would not change much. But for Claude it changes a lot, because Claude does not see Notion documents as files in a folder. Instead it must go through an MCP server or an API to interact with Notion, and the same is true of any note taking tool that does not store your notes in an easily legible format on your own device.
That has several disadvantages. The model spends far more tokens and thinking on the question of how to reach the files at all, which is unnecessarily difficult. Every abstraction layer you introduce makes it harder.
If a model wants to find text somewhere on your computer, it can use something like “grep”. grep is a search tool that has been sitting on every Mac and Linux machine for about fifty years. It looks inside your files, not just at their names, and it goes through thousands of them in well under a second.
That opens up a lot. Claude can find every note in a client folder that mentions one supplier, across hundreds of subfolders, without you remembering which note it was in. It can list every note you have not touched since March. It can find the notes that mention two things at once, which is the question you can never answer with a search box. It can ask for every file longer than a certain length, or rename fifty notes to match a convention.
None of that needs a plugin, an installation, or any special tool. These are basic command line commands, from the same generation as grep, and the model has seen them used more times in its training data than almost anything else. It is fluent in them the way it is fluent in English. The Notion MCP server, by comparison, did not exist two years ago, and it is complicated to use, so the model has to work out how to drive the tool at the same time as it does your actual work. The files are just there, in the same folder Claude Code runs in.
Here is the proof that this is the model’s own preference and not just mine. Ask Claude Code to write something down, store it for later, keep it as a reference document. What kind of file do you think it creates? Markdown. You don’t have to change its ways at all.
You could also work with Word. Word is a common document editor that a lot of people are familiar with. But Word files are in a proprietary format, and the markdown file is to the Word file what the comma separated values file is to the spreadsheet. A CSV is much easier for a language model to read than an Excel file. That is why Claude and other language models have been notoriously lagging behind in output quality when they work with these more complicated file formats. If you have been trying to get Claude to edit your Word files, you will have noticed it underperforms compared to working directly in a markdown file.
And for you it is also quite convenient, because Obsidian is genuinely fun to use. It looks nice and it has most of the functionality you need. A little less than Notion, a little less than Word, but it is the lowest common denominator in the good sense: for both of you it is just really good to work in. It is a very fleshed out notes app with high customizability. PDFs it displays out of the box, with nothing to install. For most other file types there is a plugin:
“CSV Editor” for comma separated tables
“Sheet Plus” or Excel for spreadsheets
“Annotator” if you want to highlight a PDF rather than just read it.
What is the actual benefit over a normal Claude chat?
Where do the fruits of your labor, the results of the interaction with Claude, actually get stored? If you’re using Claude as basically a pimped-up version of Google to answer singular questions, you aren’t getting half of what you can out of it. That’s really the advantage that Claude plus Obsidian has: it allows you to work on long-term projects and to build up a shared knowledge base that you and Claude can interact with, which for any complicated task is going to be necessary.
Here is a real one. Last November I came onto a large enterprise project in a field I knew very little about. In the first weeks I was sent an enormous amount of material, and there were so many new people to memorize. In the meetings I did not get most of what was being explained to me. They would mention fifteen people I did not know yet, sometimes the same ones from a meeting two hours earlier, and I still did not remember them. Nine months in, that one folder holds a hundred transcripts and about three hundred wiki pages.
Claude organized all of it into a wiki inside that folder, and it started to make sense so much sooner. I could ask my questions there instead of being the guy who asks stupid questions all the time. I could spend a day on “explain back to me what all these files were, and what did they mean by this term I keep hearing”, and the answer came back with its receipts: you discussed it in that meeting, and the document they sent five days ago describes it on page 90. They gave me the feedback that they found it amazing how quickly I started to get things, in a field everyone knew was hard to get into.
I also got into the habit of transcribing the

[truncated]
