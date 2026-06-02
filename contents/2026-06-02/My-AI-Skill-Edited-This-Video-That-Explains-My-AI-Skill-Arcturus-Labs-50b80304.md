---
source: "https://arcturus-labs.com/blog/2026/05/31/my-ai-skill-edited-this-video-that-explains-my-ai-skill/"
hn_url: "https://news.ycombinator.com/item?id=48353211"
title: "My AI Skill Edited This Video That Explains My AI Skill – Arcturus Labs"
article_title: "My AI Skill Edited This Video That Explains My AI Skill - Arcturus Labs"
author: "JohnBerryman"
captured_at: "2026-06-02T00:44:54Z"
capture_tool: "hn-digest"
hn_id: 48353211
score: 2
comments: 0
posted_at: "2026-06-01T06:15:22Z"
tags:
  - hacker-news
  - translated
---

# My AI Skill Edited This Video That Explains My AI Skill – Arcturus Labs

- HN: [48353211](https://news.ycombinator.com/item?id=48353211)
- Source: [arcturus-labs.com](https://arcturus-labs.com/blog/2026/05/31/my-ai-skill-edited-this-video-that-explains-my-ai-skill/)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T06:15:22Z

## Translation

タイトル: 私の AI スキルが私の AI スキルを説明するビデオを編集しました – Arcturus Labs
記事タイトル: 私の AI スキルを説明するこのビデオを編集しました - Arcturus Labs
説明: Cursor エージェントに AI ビデオ編集に関する YouTube チュートリアルを見てもらい、そのプロセスをエージェント スキルとしてパッケージ化して約 5 分で再現してもらいました。バグ修正が 1 つ行われ、ビデオが勝手に編集されるようになりました。これも含めて！

記事本文:
私の AI スキルが私の AI スキルを説明するビデオを編集しました - Arcturus Labs
コンテンツにスキップ
アルクトゥルス研究所
私の AI スキルが私の AI スキルを説明するビデオを編集しました
検索を初期化しています
アルクトゥルス研究所
ホーム
カテゴリー
カテゴリー
AI製品
インデックスに戻る
ジョン・ベリーマン
Arcturus Labsの創設者
メタデータ
2026 年 5 月 31 日
で
エージェントAI、
開発方法論
私の AI スキルが私の AI スキルを説明するビデオを編集しました
AI は素晴らしい時代です。自動ビデオ編集ツールを作成したところです。この投稿の最後にあるチュートリアルは、私が作成したツールによって編集されたものです。私がやった方法は次のとおりです。
それはビデオから始まりました: Hamel Hussein と Shaw Talebi と一緒にビデオを編集する AI エージェントを構築する。これらは AI 支援によるビデオ編集のプロセスを説明しており、私はそれを再現してみようと思うほどインスピレーションを受けましたが、非常にメタなプロセスを使用していました。私は、自分で何かを構築するのではなく、Cursor エージェントにビデオを見て、同じものをローカルで効果的に作成する方法を見つけるように言いました。
エージェントは、yt-dlp を使用してトランスクリプトを取得することから始めました。ちなみに、これは知っておくべき素晴らしいツールです。これを使用して、YouTube ビデオとそのトランスクリプトをダウンロードできます。トランスクリプトを取得したら、すべてを読み、プロセスを可能な限り再現するリポジトリをセットアップするように指示しました。また、スキル作成スキルを使用してすべてをエージェント スキルとしてパッケージ化するように指示しました。 ( 最近はすべてがメタ的です! ) 数分間ガクガクと音を立てて消え、戻ってきたら、あとは AssemblyAI API キーを設定することだけでした。
AssemblyAI は素晴らしい発見であることがわかりました。これは、重要な点として、あなたの間違い (「えー」「えー」) を保存し、すべての単語に対して正確な開始タイムスタンプと終了タイムスタンプを提供する音声テキスト変換サービスです。そのタイミングこそが​​自動カットを可能にするものなのです。それ

も簡単に試すことができます。サインアップするだけで、約 185 時間分の事前録音された文字起こしを無料で利用できます。ということで、もちろん登録させていただきました！
私はカメラに向かって話している自分のビデオをアップロードし、カメラに文字起こし、沈黙をカット、うーん、あのーを削除するように指示しました。そして、それはほとんどうまくいきました。問題は、カットポイントですべての単語の末尾が切り取られてしまうことでした。ということで、使えるビデオではありませんでしたが、実際にそれと同じくらい優れたビデオだったのには驚きました。
2 回目の試行では、同じタスクを与えましたが、クリッピングの問題にフラグを立てて、何が問題になっているのかを把握するように依頼しました。いくつかのアイデアが戻ってきて、私が何をしたいのか尋ねました。私は「あなたが最善だと思うことをしてください」と言いました。 🤣 (これは単なるサイド プロジェクトであるはずだったので、このプロジェクトではコードの品質は私の最大の関心事ではありません。) さらに 30 秒ほど実行し、次のビデオを実行しました。結果は悪くありませんでした。実際、合計投資約 5 分としてはかなり印象的でした。
それが私たちをここに連れてきました。私の 3 番目の試みはウォークスルーでした。OBS のカメラに映った私が編集スイートをどのように構築したかを説明しました。 (繰り返しますが、OBS がどのように機能するのかよくわかりませんが、AI にポインターを求めたところ、偽装するのに十分なほどうまくいきました。)
一番上のビデオは自動編集されたウォークスルーで、ツール自体によって編集され、ツールがどのように機能するかを説明するものです。下は編集前の同じ録音です。台本を読んだり、何度も立ち止まったりしていました。その違いは顕著です。
これを自分で試してみたい場合は、スキルがこのリポジトリにあります。厳重な警告: それは生です。デフォルトでは、スキルは Final Cut Pro にエ​​クスポートされますが、クリップを直接結合するように指示すると、代わりにそれが行われます。また、各ビデオプロジェクトのファイルを管理する方法は少し風変わりです。しかし、核となるアイデアは機能しており、それが重要なのです。つかんで、遊んで、

私がショー・タレビのアプローチを自分のものにしたのと同じように、コピーのコピーを自分のものにしてください。
ちなみに、この投稿は、上記の編集されたウォークスルーのトランスクリプトから半自動化されたものです。そのアプローチについても、再帰的なブログ記事をまた書こうと思います。

## Original Extract

I had my Cursor agent watch a YouTube tutorial on AI video editing and replicate the process – packaged as an agent skill – in about five minutes. One bug fix later, it was editing videos on its own. Including this one!

My AI Skill Edited This Video That Explains My AI Skill - Arcturus Labs
Skip to content
Arcturus Labs
My AI Skill Edited This Video That Explains My AI Skill
Initializing search
Arcturus Labs
Home
Categories
Categories
AI Product
Back to index
John Berryman
Founder of Arcturus Labs
Metadata
May 31, 2026
in
Agentic AI ,
Development Methodology
My AI Skill Edited This Video That Explains My AI Skill
These are amazing times in AI. I just created an automated video editing tool – and the walkthrough at the end of this post was edited by the tool I created. Here's how I did it.
It started with a video: Building an AI Agent to Edit Your Videos with Hamel Hussein and Shaw Talebi. They describe a process for AI-assisted video editing, and I was inspired enough to try to replicate it – but using a very meta process. Rather than building the thing myself, I told my Cursor agent to watch their video and figure out how to make effectively the same thing locally.
The agent started by pulling down the transcript using yt-dlp – a great tool to know about, by the way. You can use it to download YouTube videos and their transcripts. Once it had the transcript, I told it to read through and set up a repo that replicated the process as best it could. I also told it to package everything up as an agent skill using the create-skill skill. ( Everything is so meta these days! ) It chugged away for a few minutes, and when it came back, the only thing left on my plate was setting up an AssemblyAI API key.
AssemblyAI turns out to be a great find. It's a speech-to-text service that, crucially, preserves your mistakes – the ums, the uhs – and provides accurate start and stop timestamps for every single word. That timing is the thing that makes automated cutting possible. It's also easy to try: you get around 185 hours of free pre-recorded transcription just for signing up. So of course I signed up!
I uploaded a video of myself talking to my camera, told it to do the works – transcribe it, cut the silences, remove the ums and uhs – and it worked... mostly. The problem was that it was clipping the end of every word at the cut point. So, not usable video, but it was amazing to me that it was actually as good as it was.
For my second attempt I gave it the same task but flagged the clipping issue and asked it to figure out what was going wrong. It came back with some ideas and asked what I wanted to do. I said "do what you think is best." 🤣 (This was just supposed to be a side project, so code quality isn't my biggest concern on this one.) It ran for about 30 more seconds, I put the next video through, and the results were not bad – actually pretty impressive for about five minutes of total investment.
Which brings us here. My third attempt was the walkthrough – me on camera in OBS explaining how I built the editing suite. (Again, I don't really know how OBS works, but I asked an AI for pointers and it got me going well enough to fake it.)
The top video is the auto-edited walkthrough – the one where I explain how the tool works, edited by the tool itself. The bottom is the same recording before editing. I was reading from a script and pausing a lot. The difference is striking.
If you want to try this yourself, the skill is in this repo . Fair warning: it's raw. By default the skill exports to Final Cut Pro, though if you ask it to just stitch the clips together directly it'll do that instead. Also, the way it manages files for each video project is a bit quirky. But the core idea works, and that's the point. Grab it, play with it, make it your own – a copy of a copy – just like I made mine from Shaw Talebi's approach my own.
This post, by the way, was semi-automated from the transcript of the edited walkthrough above. I think I'll write another recursive blog post about that approach too!
