---
source: "https://jamesoclaire.com/2026/07/15/how-to-stay-in-the-coding-flow-using-llms/"
hn_url: "https://news.ycombinator.com/item?id=48917138"
title: "How to Stay in the Coding Flow Using LLMs"
article_title: "How to stay in the coding flow using LLMs – James O'Claire"
author: "ddxv"
captured_at: "2026-07-15T07:07:20Z"
capture_tool: "hn-digest"
hn_id: 48917138
score: 2
comments: 1
posted_at: "2026-07-15T06:51:17Z"
tags:
  - hacker-news
  - translated
---

# How to Stay in the Coding Flow Using LLMs

- HN: [48917138](https://news.ycombinator.com/item?id=48917138)
- Source: [jamesoclaire.com](https://jamesoclaire.com/2026/07/15/how-to-stay-in-the-coding-flow-using-llms/)
- Score: 2
- Comments: 1
- Posted: 2026-07-15T06:51:17Z

## Translation

タイトル: LLM を使用してコーディング フローを維持する方法
記事のタイトル: LLM を使用してコーディング フローを維持する方法 – James O'Claire

記事本文:
LLM を使用してコーディング フローを維持する方法 – James O'Claire
コンテンツにスキップ
ジェームズ・オクレア
Appゴブリン – 無料のアプリマーケティングツール
Appゴブリン – アプリB2Bエコシステムダッシュ
LLM を使用してコーディング フローを維持する方法
LLM やエージェントに移行すると、コード ベースの一部、あるいは場合によってはすべてとの接触が失われるという感覚が生じることは誰もが知っています。これは、ビジネス ロジックから実装への変換の管理と処理に問題があるというだけではなく、骨が折れるという点でも問題です。
私は 12 時間続いたコーディング セッションを経験しましたが、その後はとても気分がよかったです。その間、私はLLMプロンプトを数時間実行しましたが、疲れ果てたか、自分が何をしたかわからないと感じました。
最近私はこのことを念頭に置き、フロー状態を維持して LLM を活用する方法をいくつか探しています。
1) ブラウザでチャットボットを使用する
2025 年当時、これがデフォルトの使用方法だったのを覚えていますか?実際、私は今でもこれが私の好みの方法だと思っています。プロジェクトでコード ハーネスとともに LLM を使用すると、不必要な情報が大量に注入されるため、単純な質問をすることは制御不能になります。
たとえば、ここではいくつかのデータを調査していて、簡単な正規表現が必要だったので、VSCode チャット ウィンドウに目を向け、それがエージェントであることを忘れて質問してしまいました。ファイルの確認、コードの実行などを開始します。すべてが私が必要とするものから外れています。
そこで次に、VSCode をエージェントではなく「Ask」に切り替えました。すると、再び LLM にプロジェクトに関するコンテキストが大量に入力され、気が散るようなトピックから外れたコードの提案が大量に出力され始めます。
解決？一部の質問をブラウザベースのチャットに切り替える
あなたが取り組んでいることについてのコンテキストがほとんど、またはまったくないブラウザのチャット ウィンドウに切り替えて、具体的な質問をすると、ブーッといくつかの短い言葉が吐き出されます。

私の Python リスト理解のための k 正規表現は、まさに私が必要としているものです。
これは悪いアドバイスですか?まあ、たぶん。しかし、これはあなたがすでにやっていることですか、間違いなく。ただし、ここでのポイントは、一度に複数のコーディングをマルチタスクで行うことです。この方が、ニュースを閲覧するよりもはるかに良いフロー状態を維持できることがわかりました。
したがって、エージェントから切り替える -> ソーシャル メディアを閲覧するのではなく、複数のプロジェクト間を切り替えます。これはコード/作業の構造によって異なりますが、スコープに応じて、同じプロジェクト内の複数のエージェント間で切り替えるか、複数のプロジェクトを同時に開くことを意味します。
ゾーンを維持するためのポジティブなマルチタスクの種類:
ファイルとプロジェクトのクリーンアップ。 LLM は多くの追加ファイルとコードを生成するため、自分で常に把握しておくことが最善です。余分なファイルを確認して削除します。何を削除するかについて LLM にアドバイスを求めてみてください。ただし、この考えには注意してください。
3) ビジネス/コアロジックがどこにあるのかを念頭に置く
Appゴブリンの無料ASOとモバイルアプリのエコシステムデータに取り組んでいる私にとって、何が起こっているのかを*私*が理解する必要がある特定の領域があります。そのような理由から、私はAIに定型コード以上のものを書かせないようにしています。これの最も明確な例は SQL です。SQL には、最も重要なリレーショナル ロジックが数多く存在します。
確かに、LLM に複雑な SQL を 1 回実行させると「機能」しますが、数週間 (または数か月!) 後に紛れ込んだ複雑なバグを見つけることになります。この状況で誰が正しかったか間違っていたかについては必ずしも関係なく、*私* はコードベースの特定の部分で何が起こっているのかを知る必要があるということです。 「良さそうだ」と思っていたものが、後でそれが自分の望んでいたものではなかったと感じると、ひどい気分になります。
4) LLM にコードを作成させる -> 手動でステップ実行する
この最後のものは、おそらく他のデータ処理に最も適しています。

そこは、私がゾーンに留まるための素晴らしいスイートスポットを見つける場所です。
コードを書く私のお気に入りの方法は、エディターでコードを書いて REPL に行を送信することです。これは多かれ少なかれ、SQL の作成方法でもあり、SQL エディターでゆっくりとデータに変更を加え、値や仮定をチェックしてクエリを構築し、最終的に最終的な SQL クエリに到達します。
LLM では、最近このフローを使用していることに気づきました。
データを処理するための新しいコードを作成するように LLM に指示します。
コードを 1 行ずつステップ実行し、前提条件やトリッキーなデータが存在する可能性があるホットスポットをチェックします。
以前とほぼ同じですが、書く量が大幅に減り、難しい概念をより長く保持できるようになります。
楽しんでいただけましたら、お気軽にシェアしてください。
モバイルマーケティングと広告
LLM を使用してコーディング フローを維持する方法
妊娠と健康アプリは2026年になってもデータ漏洩
オープンウェイトモデルの耐えられない安さ
Appゴブリンを使用して、iOS または Android アプリの SDK と API 呼び出しを無料でスキャンします。ログインは必要ありません。
ブラウザでのアトリビューション: Google と Meta の新しいプライバシー基準から本当に恩恵を受けるのは誰ですか

## Original Extract

How to stay in the coding flow using LLMs – James O'Claire
Skip to content
James O'Claire
AppGoblin – Free App Marketing Tools
AppGoblin – App B2B Ecosystem Dash
How to stay in the coding flow using LLMs
We all know that moving to LLMs and agents has caused the feeling of losing touch with parts, or maybe even all, of a code base. This isn’t just something that is problematic for managing and handling the translation from business logic to implementation it is a problem because it feels exhausting .
I’ve had coding sessions that lasted 12 hours and afterwards felt great. Meanwhile I’ve done LLM prompting for a few hours and felt exhausted or unsure of what I did.
Lately I’ve been keeping this in mind and have been looking for a few ways in which I can maintain a flow state and take advantage of LLMs.
1) Use chatbots in the browser
Remember back in 2025 when this was the default way of using them? I actually still find this to be my preferred way. Using LLMs with code harnesses in projects injects so much unnecessary information that asking simple questions gets out control.
For example, here I’m exploring some data, and I wanted a quick regex, I turned over to my VSCode chat window, and forgot that it was an agent, and asked it the question. It proceeds to start looking at the files, wanting to run code etc. All off target of what I need .
So next I switched VSCode to “Ask” instead of agent, again the LLM is flooded with context about my project and proceeds to output a massive amount of distracting and off topic code suggestions.
Solution? Switch some questions to browser based chat
Switch to a browser chat window which has little to no context about what you’re working on and ask it my specific question, boom it spits out a few quick regexes for my Python list comprehension that are exactly what I need .
Is this bad advice? Well, maybe. But was this what you’re already doing, definitely. But the point here is to multitask coding on more than one thing at a time. I’ve found that this keeps me in the flow state much better than if I let myself browse the news.
So instead of switching from your agent -> browse social media switch between multiple projects. This depends on how your code / work is structured, but depending on the scope this means either switching between several agents in the same project or having several projects open at once.
Types of positive multitasking to stay in the zone:
File and project cleanup. LLMs generate many extra files and code and it’s best to stay on top of that yourself. Go through and delete extra files. Try asking LLMs for advice on what to remove, but do be careful with this idea.
3) Keep in mind where your business / core logic is
For me, working on AppGoblin’s free ASO and mobile app ecosystem data , I have certain areas that *I* need to understand what is happening, for those reasons I do not let AI write anything more than boiler plate code. The clearest example of this I can give is SQL, where a lot of my most important relational logic exists.
Sure, I can let an LLM one shot a complicated SQL and it will “work” but come weeks (or months!) later and I’ll find a complicated bug that slipped in. It’s not even necessarily about who was right/wrong in this situation, it’s that *I* need to know what’s going on in certain parts of the codebase. Something that ‘looks fine’ is a terrible feeling that later it was not what I wanted.
4) Have the LLM Write Code -> Step through manually
This last one is probably best suited for other data crunchers out there, but it’s where I find a great sweet spot for staying in the zone.
My favorite way to write code has always been to write code in an editor and send line to a REPL. This is also more or less how SQL gets written as well where you build queries in your SQL editor by slowly making changes to the data, checking values / assumptions and eventually getting to your final SQL query.
With the LLMs, I find myself using this flow lately:
Tell LLM to write new code for processing data
Step through the code my self line by line, checking the hotspots where I know assumptions / tricky data might be
It’s more or less the same as I did before, just a lot less writing and let’s me hold onto the difficult concepts longer.
If you enjoyed this feel free to share.
Mobile Marketing and Advertising
How to stay in the coding flow using LLMs
The Pregnancy and Health Apps Still Leaking Data in 2026
The Unbearable Cheapness of Open Weight Models
Scan any iOS or Android App for SDKs and API Calls for Free with AppGoblin, no login
Attribution in the Browser: Who Really Benefits from Google and Meta’s New Privacy Standard
