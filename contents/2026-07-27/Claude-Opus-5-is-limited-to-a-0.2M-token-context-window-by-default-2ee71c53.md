---
source: "https://wadetregaskis.com/claude-opus-5-is-limited-to-a-0-2m-token-context-window-by-default/"
hn_url: "https://news.ycombinator.com/item?id=49068610"
title: "Claude Opus 5 is limited to a 0.2M token context window by default"
article_title: "Claude Opus 5 is limited to a 0.2M token context window by default – Wade Tregaskis"
author: "colejohnson66"
captured_at: "2026-07-27T12:51:21Z"
capture_tool: "hn-digest"
hn_id: 49068610
score: 2
comments: 0
posted_at: "2026-07-27T12:21:19Z"
tags:
  - hacker-news
  - translated
---

# Claude Opus 5 is limited to a 0.2M token context window by default

- HN: [49068610](https://news.ycombinator.com/item?id=49068610)
- Source: [wadetregaskis.com](https://wadetregaskis.com/claude-opus-5-is-limited-to-a-0-2m-token-context-window-by-default/)
- Score: 2
- Comments: 0
- Posted: 2026-07-27T12:21:19Z

## Translation

タイトル: Claude Opus 5 はデフォルトで 0.2M トークンのコンテキスト ウィンドウに制限されています
記事のタイトル: Claude Opus 5 はデフォルトで 0.2M トークンのコンテキスト ウィンドウに制限されています – Wade Tregaskis

記事本文:
Claude Opus 5 はデフォルトで 0.2M トークンのコンテキスト ウィンドウに制限されています
Opus 4.8 以前では、(少なくとも有料プランでは) 1M トークンのコンテキスト ウィンドウが使用されます。したがって、Anthropic が Opus 5 でこれを密かに 80% 削減したのは驚くべきことです。
ただし、使用クレジットを有効にするだけで、1M トークンのコンテキスト ウィンドウを取り戻すことができます。使用量クレジットを支払う必要も、自動リロードをオンにする必要も、ゼロ以外の制限も必要ありません。文字通り、使用クレジットの切り替えを有効にするだけです。 🤷‍♂️
コンテキスト ウィンドウのサイズが重要なのはなぜですか?特にコーディングの場合、単一のタスクが長時間実行され、非常に多くのトークンを使用する場合や、ソース ファイルの読み取り、複雑なロジックと制御フローの精神的な評価などで、モデルの効率を最大限に高めるために、すべてが 1 つのコンテキスト ウィンドウに収まることが重要です。モデルがタスクの途中で圧縮を実行する必要がある場合、(a) 重要な詳細を「忘れる」ため、より多くの間違いが発生し、(b) 再読み取りまたは再計算が必要になり、トークン (つまり、貴重な使用クォータ) が無駄になるという重大なリスクがあります。
0.2M は現代の標準からするとかなり小さいですが、私は実際にトークンとコンテキスト ウィンドウの使用状況を多少なりとも追跡しているので、何と言っても啓発のために – 日常のコーディング タスクではそれをはるかに超える量を定期的に使用していることを知っています。実際、100 万トークンであっても、時折タスクを実行するには十分なウィンドウではありません (ただし、Claude モデルの最近のバージョンではサブエージェントをより積極的に使用する傾向があり、これによりウィンドウ制限に達するのを回避できますが、全体的なトークン効率は低下します)。
Anthropic がなぜこれを行ったのかについては…Opus 5 はコンパクション全体でコンテキストを運ぶのが非常に上手なので、それを回避できると彼らが信じているのではないかと思います。テストしていません。それはそうだろう

本当であれば、これは素晴らしいことです。コンテキスト ウィンドウが小さいほど、計算コストが低くなります (セッション内で頻繁に会話を行っている場合は、計算コストも大幅に低くなります)。
カテゴリ コーディング , ハウツー タグ 人間 , クロード , クロード オーパス 4.8 , クロード オーパス 5 , コンテキスト ウィンドウ , トークン , 使用クレジット
Apple の notarytool は ZIP64 ファイルを処理できません
コメントを残す 返信をキャンセル
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
フォローアップコメントを電子メールで通知します。
新しい投稿をメールで通知します。
マストドン: @everything@wadetregaskis.com
このウェブサイトをフォローするその他の方法

## Original Extract

Claude Opus 5 is limited to a 0.2M token context window by default
Opus 4.8 and earlier use a 1M token context window (on any paid plan, at least). So it’s surprising that Anthropic quietly reduced this by 80% with Opus 5.
You can get the 1M token context window back, though – you just have to enable Usage Credits. You don’t have to pay for any usage credits, nor turn on auto-reload, nor have a non-zero limit. You literally just have to have the Usage Credits toggle enabled. 🤷‍♂️
Why does the context window size matter? Especially for coding – where single tasks can be long-running and use very large numbers of tokens, for reading source files, mentally evaluating complex logic & control flow, and so forth – it’s important that everything fit into a single context window, for maximum efficacy of the model. If the model has to run compaction mid-way through a task, there’s a significant risk it’ll (a) ‘forget’ important details, and therefore make more mistakes, and (b) have to re-read or re-compute things, wasting tokens (i.e. your precious usage quota).
0.2M is pretty tiny by modern standards, and – because I do actually track token & context window usage somewhat, for edification if nothing else – I know that I regularly use far more than that in everyday coding tasks. In fact even 1M tokens isn’t a big enough window for the occasional task (though recent versions of Claude models tend to use subagents more aggressively, which helps avoid hitting the window limit – albeit at the expense of less token efficiency overall).
As to why Anthropic did this… I suppose it’s conceivable that they believe Opus 5 is so good at carrying context across compaction that it can get away with it? I haven’t tested it. That’d be wonderful, if true – a smaller context window is less computationally costly (and much less costly if you’re conversing back-and-forth a lot within a session).
Categories Coding , Howto Tags Anthropic , Claude , Claude Opus 4.8 , Claude Opus 5 , context window , tokens , Usage Credits
Apple’s notarytool can’t handle ZIP64 files
Leave a Comment Cancel reply
Save my name, email, and website in this browser for the next time I comment.
Notify me of follow-up comments by email.
Notify me of new posts by email.
Mastodon: @everything@wadetregaskis.com
Other ways to follow this website
