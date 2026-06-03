---
source: "https://thomasdhughes.com/brontosaurus/"
hn_url: "https://news.ycombinator.com/item?id=48377982"
title: "Show HN: Brontosaurus, a voice-driven generative AI canvas"
article_title: "Brontosaurus: A Voice-Driven Generative AI Canvas | Thomas Hughes"
author: "thomasdhughes2"
captured_at: "2026-06-03T00:34:02Z"
capture_tool: "hn-digest"
hn_id: 48377982
score: 1
comments: 1
posted_at: "2026-06-03T00:12:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Brontosaurus, a voice-driven generative AI canvas

- HN: [48377982](https://news.ycombinator.com/item?id=48377982)
- Source: [thomasdhughes.com](https://thomasdhughes.com/brontosaurus/)
- Score: 1
- Comments: 1
- Posted: 2026-06-03T00:12:46Z

## Translation

タイトル: 番組 HN: 音声駆動の生成 AI キャンバス、ブロントサウルス
記事のタイトル: ブロントサウルス: 音声駆動の生成 AI キャンバス |トーマス・ヒューズ
説明: 人間と AI のコラボレーションに関する追加の考え
HN テキスト: こんにちは、HN!これを作成したのは、音声と超高速推論によって、自分のマシン上で何かを存在させているように感じられるかどうかを確認するためです。数時間遊んでみて、魔法のようなものを感じました。技術的な詳細も投稿にあります。良い一日！

記事本文:
ブロントサウルス: 音声主導の生成 AI キャンバス |トーマス・ヒューズ トーマス・ヒューズ
プロジェクトと執筆
/
メモ
/
以前
/
プロジェクトと執筆に戻る Brontosaurus: A Voice-Driven Generative AI Canvas 人間と AI のコラボレーションについての補足説明 2026 年 6 月 2 日
Brontosaurus は Web ベースの生成キャンバスで、見たいものを声に出して話すと、Bronto が 1 秒以内にそのウィジェットを構築します。基盤となるエージェントは OpenAI の gpt-oss-120b 上で実行され、Cerebras によって毎秒 3,000 トークンという驚異的な速度でサービスが提供され、全体が魔法のように感じられます :)
Brontosaurus は現在開発中です。アイデアや見たいものがある場合は、私に送ってください。ぜひ聞きたいです。 ( hello@thomasdhughes.com )
このプロジェクトは 2 つの驚異的なブログ投稿からインスピレーションを受けました。
1 つ目は、Thinking Machines によるインタラクション モデルのリリースでした。
技術的な面では、会話の流れを中断することなくリクエストを静かに実行する、より強力な準バックグラウンド エージェントに接続されたマルチモーダル (音声 + ビジョン + テキスト) モデルのアーキテクチャが気に入りました。現在、これを行っている研究室は確かに他にもありますが、Thinking Machines は独自のアプローチを採用しています。Sean Goedecke は、自身のブログ投稿でその斬新さをうまく説明しています。
哲学的な側面として、Thinking Machines は、AI エージェントに関する現在の議論は、人間と AI のコラボレーション (エージェントが人間と協力してタスクに取り組む能力) ではなく、エージェントの自律性 (エージェントがタスクを受け取り、その後何時間も中断されずに作業できる能力) を誤って中心に置いていると主張します。彼らは、これは間違いであり、このようにモデルを設計すると、「仕事に人間が必要ないからではなく、インターフェイスに人間の余地がないために人間がますます追い出されていく」ことにつながると主張しています。あなたが入れた

あなたのプロンプトで、邪魔にならないようにしてください。逆に言えば、この新しいインタラクション モデル ファミリは、チームメイトと協力するような方法で動作します。話したり、タイプしたり、物を指さしたり、新しいアイデアで中断したりすることができます。コラボレーションすることができます。
このアプローチの技術的な部分と哲学的な部分の両方が私に訴えかけ、それらが合わさって同じ精神のものを構築したいと思うようになりました。ブロントでは、それは何よりも思考のスピードでの創造を優先するという形で現れています。それは、能力、知性、長期にわたるタスクの自律性以上に、1 秒以内に何かを存在させる能力があれば、何でも可能であると感じさせるという主張です。
2 番目のブログ投稿は、Ink & Switch からのものでした。Ink & Switch は、柔軟なソフトウェアに焦点を当てた独立した研究機関であり、LLM で生成された Javascript ブックマークを使用した Web サイトの変更で私が検討した概念です。この投稿は「chitter chatter」と呼ばれ、生成キャンバスのビジョンを概説しています。誰かの日記のように読めて、提案されたソフトウェアはフレンドリーで温かみのあるものに聞こえますが、コードは冷たく数字で表現されることが多く、それができる人にはいつも感心します。私は誇りを持っているので、その表現について謝罪するつもりはありませんが、私はその言葉を支持しません。今後は言及しないでおくつもりです。
私はこの作品を (Thinking Machines と一緒に) 読んで気に入って、同じようなものを作りたいと思っていました。ブロントサウルスが誕生しました。
内部では、マルチエージェントのオーケストレーションが進行しています。
Conductor と Builder という 2 つのエージェント タイプが使用されます。どちらも OpenAI の gpt-oss-120b 上で実行され、Cerebras によって 3,000 トークン/秒で提供されます。これを大局的に見ると、ブラウザーの ChatGPT は 1 秒あたり約 50 トークンで応答します。
スペースバーをタップすると、Web アプリがリスニングを開始し、もう一度タップすると、Chrome の機能を使用して音声テキスト変換を実行します。

ilt-in Web 音声 API。
発言内容のテキストは、現在キャンバス上にあるウィジェットを記述する JSON 配列とともに Conductor エージェントに渡されます。
ID (ツール呼び出しの一意の識別子)、
タイトル (各ウィジェットの上部に表示されるもの)、
説明 (ウィジェットの内容についての内部向けの説明)、および
四角形 (キャンバス上のウィジェットのサイズと位置)。
その後、Conductor エージェントがツール呼び出しを行います。それはできます
配置 - 内容を変更せずに、 id によってウィジェットを移動またはサイズ変更します (これは、そのウィジェットのrect値を更新することによって行われます)。
delete - id によってウィジェットを削除します。
クリア - すべてのウィジェットを一度に削除します。
最初の 3 つのツール呼び出しは決定的に処理されます。最後の 2 つ (作成と編集) は、Builder エージェントに指示を送信します。
作成時に、Builder エージェントは要求されたウィジェットの説明のみを受け取ります。編集時には、現在のウィジェットの完全な HTML と変更指示を受け取ります。
次に、Builder エージェントは完全な自己完結型 HTML ドキュメントを返し、その後クリーンアップされて iframe でレンダリングされます。
魔法をさらに加える追加のデザインの選択肢:
Conductor エージェントは 1 つの命令で複数のツール呼び出しを行うことができます。これにより、「ピアノを削除し、計算機を作成し、ピアノがあった場所に置きます」という命令がすべて一度に実行されます。
アレンジ呼び出しでは、Builder エージェントの構築が完了するまで待つ必要はありません。作成が実行されるとすぐに ID が存在するため、まだ設定中のウィジェットを移動できます。
ビルダー エージェントは並行して実行されるため、複数のウィジェットを一度に作成できます。
ここには改善の余地がたくさんあります。
たとえば、gpt-oss-120b は 9 か月前のもので、パラメータはわずか 120B です。これは、非常に安いことを意味します。私はブロントサウルスを 1 時間以上ノンストップで使用しました。

1 ドル未満ではなく、成果物の品質の上限ははるかに高いということです。 Z.ai (Cerebras も提供) の GLM 4.7 のようなモデルを使用した場合、コストは 4 倍、速度は 1/3 になりますが、パラメータは 3 倍になるため、より複雑なウィジェットを構築できる可能性があります。問題は、速度を犠牲にする価値があるかどうかです。
これに関して、私は当初、Bronto が天気やライブ株価などを取得できるように Exa AI を介してライブ検索機能を追加しましたが、私の評価では最大 0.9 秒の遅延が追加されました。これは、第 2 世代未満の速度に慣れると苦痛になります。
最後に、主要なものは仮想ファイル システムです。現時点では、Bronto が作成するすべてのウィジェットは iframe 内の一時的な使い捨て HTML ファイルですが、VFS を使用すると、既存のドキュメントや以前に構築されたウィジェットを表示して反復処理できるほか、Bronto がウィジェットのコンテンツを選択的にコンテキスト ウィンドウに取り込むことができるため、「成分リストから既に持っているものにチェックを入れました。それらを削除してください」のようなコマンドが機能します。
上記はすべて技術的な変更です。しかし、現在のアーキテクチャだけで実現できる興味深いことがたくさんあることはわかっています。デモの最後にある8列ステップシーケンサー（ビートメーカー）は、最初は私が「音楽を作りたい」と言って作ったもので、Brontoが3つのウィジェットを制作し、そのうちの1つでした。完全に衝撃を受けました。これが、私が一番上に書いた理由です。アイデアがある場合は、ぜひ送ってください。 ( hello@thomasdhughes.com ) 試してみて、ビデオを送り返すことを約束します。
あなたの仕事と執筆で私にインスピレーションを与えてくれた Thinking Machines と Ink & Switch に感謝します。
決して叶えられないことを願うリクエストとともに、ブロントサウルスの初期バージョンのプレッシャーテストをしてくれたトリスタンとスティーブンに感謝します。

## Original Extract

With additional thoughts on human-AI collaboration

Hi HN! I made this to see if voice + super fast inference could make it feel like you were speaking things into existence on your machine. In my several hours of playing with it, it feels something like magic. Technical details in the post as well. Good day!

Brontosaurus: A Voice-Driven Generative AI Canvas | Thomas Hughes Thomas Hughes
projects & writing
/
notes
/
previously
/
Back to projects & writing Brontosaurus: A Voice-Driven Generative AI Canvas With additional thoughts on human-AI collaboration Jun 2, 2026
Brontosaurus is a web-based generative canvas, where you speak aloud what you want to see, and Bronto builds a widget of it in under a second. The underlying agents run on OpenAI’s gpt-oss-120b , served by Cerebras at a blistering 3,000 tokens/second, which makes the whole thing feel like magic :)
Brontosaurus is very much in development - if you’ve got ideas or things you’d want to see, send them my way, I’d love to hear them! ( hello@thomasdhughes.com )
Two phenomenal blog posts inspired this project.
The first was Thinking Machines’ release of an Interaction Model .
On the technical side, I liked the architecture of a multi-modal (voice+vision+text) model hooked into a more powerful quasi-background agent which quietly executes requests without interrupting the flow of conversation. There are certainly other labs doing this today, but Thinking Machines takes a unique approach - Sean Goedecke does a great job of breaking down that novelty in his own blog post .
On the philosophical side, Thinking Machines argues that current discussions around AI agents mistakenly center agentic autonomy (the ability for an agent to receive a task, then work for hours uninterrupted) as opposed to human-AI collaboration (the ability for an agent to work on a task in tandem with a person). They assert that this is a mistake - that designing models this way leads to “humans increasingly get[ing] pushed out not because the work doesn’t need them, but because the interface has no room for them.” You put in your prompt and get out of the way. This new family of Interaction Models, on the flip side, operates more in the way you’d work with a teammate - you can talk, type, point at things, interrupt with new ideas. You can collaborate .
Both the technical and philosophical pieces of this approach spoke to me, and they together made me want to build something of the same ethos. In Bronto, that manifests as prioritizing creation at the speed of thought above all else. It is an argument that more than capability, more than intelligence, more than long-running task autonomy, the ability to speak something into existence in under a second makes you feel like anything is possible.
The second blog post was from Ink & Switch, an independent research lab with a focus on malleable software, a concept I explored in Modifying Websites with LLM-Generated Javascript Bookmarks . The post is called “chitter chatter”, and it outlines a vision for a generative canvas. It reads like somebody’s diary, and makes the proposed software sound friendly and warm, which I always admire when people can do - code is too often cold and numbers. I will not apologize for that phrasing because I am prideful, but I do not stand by it let us not mention it going forward.
I read this piece (along with Thinking Machines’), loved it, and wanted to build something like it. Brontosaurus was born.
Under the hood, there’s some multi-agent orchestration going on.
There are two agent types at play: Conductor and Builder. Both run on OpenAI’s gpt-oss-120b , served by Cerebras at 3,000 tokens/second. To put this in perspective, ChatGPT in the browser responds at ~50 tokens/second.
When you tap the space bar, the web app starts listening, and when you tap again, it does speech-to-text with Chrome’s built-in Web Speech API.
The text of what you said gets passed to the Conductor agent, along with a JSON array describing the widgets currently on the canvas, each of which have
an id (unique identifier for tool calls),
a title (what you see at the top of each widget),
a description (internal-facing explanation of what the widget contains), and
a rect (the widget’s size and position on the canvas).
The Conductor agent then makes tool calls. It can
arrange - move or resize a widget by id , without changing the contents (this is done by updating the rect value for that widget),
delete - remove a widget by id ,
clear - remove all widgets at once,
The first three tool calls are handled deterministically. The final two - create and edit - send instructions to a Builder agent.
When create ing, the Builder agent receives just the requested widget’s description . When edit ing, it receives the full HTML of the current widget along with the change instructions.
The Builder agent then returns a complete, self-contained HTML document, which is subsequently cleaned and rendered in an iframe .
Additional design choices which add to the magic:
The Conductor agent can make multiple tool calls with a single instruction - this makes it possible to say “delete the piano, make a calculator, put it where the piano was” and it all happens at once.
arrange calls don’t need to wait for the Builder agent to finish building - as soon as the create is run, an id exists, so widgets which are still populating can be moved around.
Builder agents run in parallel so multiple widgets can be made at once.
There’s lots of room for improvement here.
For one, gpt-oss-120b is 9 months old and just 120B parameters. This means it’s dirt cheap - I used Brontosaurus nonstop for over an hour and spent less than a dollar - and that the ceiling for quality of output is way higher. If we used a model like GLM 4.7 from Z.ai (also served by Cerebras), it’d be 4x the cost and 1/3rd the speed, but 3x the parameters so could likely build far more complex widgets. The question is if the speed tradeoff would be worth it.
On this note, I initially added live search capabilities via Exa AI so that Bronto could pull things like weather and live stock price, but in my evals it added a delay of ~0.9s, which stings once you’ve gotten used to the sub-second generation speed.
Finally, the major one - a virtual file system! At the moment, all the widgets Bronto makes are ephemeral, single-use HTML files in iframe s, but a VFS would allow for surfacing pre-existing documents and previously-built widgets to iterate on, as well as make it possible for Bronto to selectively pull the contents of widgets into its context window, so commands like “I checked off what I already have from the ingredients list, please remove those” would work.
The above are all technical changes. But I know there is a lot of interesting things that can be done just with the current architecture. The 8 row step sequencer (beat maker) at the end of the demo initially came from me saying “I want to make some music” and Bronto produced three widgets, that being one of them. It totally blew my mind. Which is why I put at the top: if you’ve got ideas, please send them my way! ( hello@thomasdhughes.com ) I’ll try them out and send back a video I promise.
Thank you Thinking Machines and Ink & Switch for inspiring me with your work and writing.
Thank you Tristan and Steven for pressure testing early versions of Brontosaurus with requests I hope it never fulfills.
