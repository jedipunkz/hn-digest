---
source: "https://lemire.me/blog/2026/07/11/x-just-gave-us-an-interface-that-ai-agents-can-use-i-pointed-it-at-my-own-posts/"
hn_url: "https://news.ycombinator.com/item?id=48904900"
title: "X just gave us an interface that AI agents can use. I pointed it at my own posts"
article_title: "X just gave us an interface that AI agents can use. I pointed it at my own posts. – Daniel Lemire's blog"
author: "ibobev"
captured_at: "2026-07-14T11:18:12Z"
capture_tool: "hn-digest"
hn_id: 48904900
score: 1
comments: 0
posted_at: "2026-07-14T11:04:56Z"
tags:
  - hacker-news
  - translated
---

# X just gave us an interface that AI agents can use. I pointed it at my own posts

- HN: [48904900](https://news.ycombinator.com/item?id=48904900)
- Source: [lemire.me](https://lemire.me/blog/2026/07/11/x-just-gave-us-an-interface-that-ai-agents-can-use-i-pointed-it-at-my-own-posts/)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T11:04:56Z

## Translation

タイトル: X は、AI エージェントが使用できるインターフェイスを提供しました。自分の投稿に指摘しました
記事のタイトル: X は、AI エージェントが使用できるインターフェイスを提供しました。自分の投稿にそれを指摘しました。 – ダニエル・レミアのブログ
説明: 私は長い間 X を使用しています。定期的に投稿するほとんどの人たちと同じように、私も人々が興味を持ちそうなものには直感があります。朝投稿します。長い投稿のほうが効果が高いようです。しかし、直感は測定値ではありません。そして最近まで、自分の投稿データを掘り下げるには、あちこちをクリックするか、
[切り捨てられた]

記事本文:
X は、AI エージェントが使用できるインターフェイスを提供しました。自分の投稿にそれを指摘しました。 – ダニエル・レミアのブログ
コンテンツにスキップ
ダニエル・レミアのブログ
Daniel Lemire はソフトウェア パフォーマンスの専門家です。彼は世界の科学者の上位 2% にランクされ (スタンフォード/エルゼビア 2025)、GitHub で最もフォローされている開発者のトップ 1000 の 1 人です。
12,500 人を超える電子メール購読者に加わりましょう:
X は、AI エージェントが使用できるインターフェイスを提供しました。自分の投稿にそれを指摘しました。
AI とチャットしてもトップ プログラマーになれるわけではない
C++26 静的リフレクションを使用したコンパイル時の JSON の解析
amd64 マイクロアーキテクチャ レベルは Go にどの程度役立ちますか?
サイモン・シュレーダー氏、AI とチャットしてもトップ プログラマーにはなれない
Kevin on あなたの CPU はいくつの分岐を予測できますか?
マーク・ハーン氏「AI とチャットしてもトップ プログラマーにはなれない」
ダニエル・レミアが語る、AI とチャットしてもトップ プログラマーにはなれない
C++26 静的リフレクションを使用したコンパイル時の JSON 解析の名前
X は、AI エージェントが使用できるインターフェイスを提供しました。自分の投稿にそれを指摘しました。
私は長い間Xに参加しています。定期的に投稿するほとんどの人たちと同じように、私も人々が興味を持ちそうなものには直感があります。朝投稿します。長い投稿のほうが効果が高いようです。
しかし、直感は測定値ではありません。そして最近まで、自分の投稿データを詳しく調べるには、Web UI をクリックするか、カスタム スクリプトを作成する必要がありました。 AI アシスタントにその場限りの質問をしたい場合には、どちらもあまりフレンドリーではありません。
X は最近、AI ツールが接続できる公式エンドポイントであるホスト型 MCP サーバーを立ち上げました。 MCP は、言語モデルにツールをプラグインするためのプロトコルです。モデルは、投稿の検索、ブックマークの管理、トレンドの取得などを行うことができます。実際に、AI コーディング エージェントを X MCP サーバーに接続し、単に自分のアカウントに関する質問を開始しました。

私はセッションで約 2 か月間、自分自身の活動を探求しました。ここで私が興味深かったことを紹介します。
約 60 日間 (2026 年 5 月中旬から 7 月中旬まで) に、私は他の人の純粋なリツイートではない 435 件の投稿を公開しました。そのほとんどは、元の投稿、返信、およびいくつかの X 記事の組み合わせでした。エージェントは MCP ツールを通じてそれらを取得し、公開指標 (いいね!、ビュー、再投稿) を保持し、簡単な分析を実行しました。
すべての投稿を現地の時間帯 (アメリカ/トロント、東部時間) ごとに、時間ごとに投稿数と最小 / 中央値 / 最大表示数を指定してビンに入れるように依頼しました。
私の投稿は朝に大きく偏っています:
午前 9 時は私にとって最も忙しい時間帯であり、中央値によると、忙しい時間帯の中でも最も忙しい時間帯です。全時間の全体の中央値はわずか約 188 ビューであったため、私が書いたもののほとんどは静かなものです。分布はヘビーテールです。いくつかの投稿が数万、数十万のインプレッションを獲得します。残りはバックグラウンドノイズです。
次に、投稿を 25 文字単位の文字長ごとにビン分けしました (短い t.co URL を含む、API から返されたテキストを使用)。
私の文章の大部分は短く、多くの場合は数十文字の返信です。
約 175 文字未満では、中央値は 0 または 1 件のままです。フルレングスの投稿 (およそ従来の 280 文字体制とそれを少し超えるもの) では、エンゲージメントが急増します。 300 ～ 325 文字の範囲には、私の「真剣な」投稿の大部分が含まれており、そこにあるいいねの中央値は、短い返信よりも桁違いに高くなります。
また、AI に「いいね」が最も多かった投稿を特定するよう依頼しました。次のタイプの投稿が「いいね」されました。
AI 対「専門家」、モデルは人間の知能には遠く及ばないと主張
SIMD スタイルのデータ並列処理を標準ライブラリに追加する
SIMD アクセラレーションによるデータ処理に関する説明とライブラリに関するメモ (JSON、

トリング→整数マップ、脆弱性レポート疲労)
Nvidia ハードウェア、大学の AI 不正行為、C++ 契約
興味深いのはワークフローです。 CSV を手動でエクスポートして Excel を開いたわけではありません。私は、X の MCP サーバーに接続しているエージェントに尋ねました。 AI エージェントが 1 つのアカウントのメトリクスに対してこれを実行できるのであれば、バグ トラッカー、ログ、紙のドラフト、コードベースに対しても実行できます。
Daniel Lemire、「X は AI エージェントが使用できるインターフェイスを提供しました。私は自分の投稿にそれを指定しました。」、Daniel Lemire のブログ、2026 年 7 月 11 日、https://lemire.me/blog/2026/07/11/x-just-gave-us-an-interface-that-ai-agents-can-use-i-pointed-it-at-my-own-posts/ 。
https://lemire.me/blog/2026/07/11/x-just-gave-us-an-interface-that-ai-agents-can-use-i-pointed-it-at-my-own-posts/ '> [BibTeX]
ケベック大学 (TELUQ) のコンピューターサイエンス教授。
ダニエル・レミアの投稿をすべて表示
あなたのメールアドレスは公開されません。
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
このブログを電子メールで購読することもできます (非営利、広告なし、毎週の電子メール)
コードを投稿する場合は、 tohtml などのツールを使用してコードをフォーマットすることを検討してください。

## Original Extract

I have been on X for a long time. Like most people who post regularly, I have a gut feeling for what might interest people. I post in the morning. Longer posts seem to do better. But gut feelings are not measurements. And until recently, digging into your own posting data meant either clicking aroun
[truncated]

X just gave us an interface that AI agents can use. I pointed it at my own posts. – Daniel Lemire's blog
Skip to content
Daniel Lemire's blog
Daniel Lemire is a software performance expert. He ranks among the top 2% of scientists globally (Stanford/Elsevier 2025) and is one of GitHub's top 1000 most followed developers.
Join over 12,500 email subscribers:
X just gave us an interface that AI agents can use. I pointed it at my own posts.
Chatting with an AI Won’t Make You a Top Programmer
Parsing JSON at compile time with C++26 static reflection
How much do amd64 microarchitecture levels help in Go?
Simon Schröder on Chatting with an AI Won’t Make You a Top Programmer
Kevin on How many branches can your CPU predict?
Mark Hahn on Chatting with an AI Won’t Make You a Top Programmer
Daniel Lemire on Chatting with an AI Won’t Make You a Top Programmer
name on Parsing JSON at compile time with C++26 static reflection
X just gave us an interface that AI agents can use. I pointed it at my own posts.
I have been on X for a long time. Like most people who post regularly, I have a gut feeling for what might interest people. I post in the morning. Longer posts seem to do better.
But gut feelings are not measurements. And until recently, digging into your own posting data meant either clicking around the web UI or writing custom scripts. Neither is particularly friendly when you want to ask ad hoc questions with an AI assistant.
X recently launched hosted MCP servers: official endpoints that AI tools can connect to. MCP is a protocol for plugging tools into language models: the model can search posts, manage bookmarks, fetch trends, and so on. In practice, I connected an AI coding agent to the X MCP server and simply started asking questions about my account.
I spent a session exploring about two months of my own activity. Here is what I found interesting.
Over roughly sixty days (mid-May through mid-July 2026), I published on the order of 435 posts that were not pure retweets of other people—mostly a mix of original posts, replies, and a few X Articles. The agent pulled them through the MCP tools, kept the public metrics (likes, views, reposts), and ran simple analyses.
I asked for every post to be binned by local hour of day (America/Toronto, Eastern time), and for each hour: how many posts, and the min / median / max view count.
My posting is heavily skewed toward the morning:
The 9 a.m. hour is both my busiest and, among busy hours, my strongest by median views. The overall median across all hours was only about 188 views, so most of what I write is quiet. The distribution is heavy-tailed: a few posts get tens or hundreds of thousands of impressions; the rest are background noise.
I then binned posts by character length in steps of 25 characters (using the text as returned by the API, including short t.co URLs).
The bulk of my writing is short, often a reply of a few dozen characters:
Under about 175 characters, the median stays at zero or one like. Around full-length posts (roughly the old 280-character regime and a bit beyond), engagement jumps. The 300–325 character band is where a large fraction of my “serious” posts live, and the median likes there are an order of magnitude higher than for short replies.
I also asked the AI to identify the posts that had the most likes, the following types of posts were liked:
AI vs. “experts” claiming models are nowhere near human intelligence
Go adding SIMD-style data-parallelism to the standard library
SIMD-accelerated data processing talks and library notes (JSON, string→integer maps, vulnerability-report fatigue)
Nvidia hardware, university AI-cheating, C++ contracts
The interesting part is the workflow. I did not export a CSV by hand and open Excel. I asked an agent, connected to X’s MCP server. If AI agents can do this for one account’s metrics, they can do it for bug trackers, logs, paper drafts, and codebases.
Daniel Lemire, "X just gave us an interface that AI agents can use. I pointed it at my own posts.," in Daniel Lemire's blog , July 11, 2026, https://lemire.me/blog/2026/07/11/x-just-gave-us-an-interface-that-ai-agents-can-use-i-pointed-it-at-my-own-posts/ .
https://lemire.me/blog/2026/07/11/x-just-gave-us-an-interface-that-ai-agents-can-use-i-pointed-it-at-my-own-posts/ '> [BibTeX]
A computer science professor at the University of Quebec (TELUQ).
View all posts by Daniel Lemire
Your email address will not be published.
Save my name, email, and website in this browser for the next time I comment.
You can also subscribe by email to this blog (non commercial, no ads, weekly email)
If you want to post code, consider formatting it with a tool like tohtml .
