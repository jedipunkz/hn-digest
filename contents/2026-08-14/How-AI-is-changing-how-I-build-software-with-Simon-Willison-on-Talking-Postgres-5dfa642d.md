---
source: "https://talkingpostgres.com/episodes/how-ai-is-changing-software-development-with-simon-willison"
hn_url: "https://news.ycombinator.com/item?id=49303217"
title: "How AI is changing how I build software with Simon Willison, on Talking Postgres"
article_title: "Talking Postgres with Claire Giordano | How AI is changing software development with Simon Willison"
author: "clairegiordano"
captured_at: "2026-08-14T19:42:28Z"
capture_tool: "hn-digest"
hn_id: 49303217
score: 2
comments: 0
posted_at: "2026-08-14T19:04:31Z"
tags:
  - hacker-news
  - translated
---

# How AI is changing how I build software with Simon Willison, on Talking Postgres

- HN: [49303217](https://news.ycombinator.com/item?id=49303217)
- Source: [talkingpostgres.com](https://talkingpostgres.com/episodes/how-ai-is-changing-software-development-with-simon-willison)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T19:04:31Z

## Translation

タイトル: AI がソフトウェア構築方法をどのように変えるか、Simon Willison と Talking Postgres で語る
記事のタイトル: Claire Giordano と Postgres について話す | AI はソフトウェア開発をどのように変えるのか、サイモン・ウィリソン氏と語る
説明: AI で書かれたコードが本番環境に使えるかどうかを判断するための最も重要な基準は何ですか? Talking Postgres のエピソード 42 では、Datasette の作成者、Django の共同作成者、そして多作の作品を手がけた Simon Willison が語ります。

記事本文:
クレア・ジョルダーノと Postgres について語る | AI はソフトウェア開発をどのように変えるのか、サイモン・ウィリソン氏と語る
クレア・ジョルダーノと Postgres について話す
閉じる
ホーム
エピソード
人々
購読する
遊ぶ
一時停止する
AI はソフトウェア開発をどのように変えるのか、サイモン・ウィリソン氏と語る
MP3をダウンロード
ノートを表示 /
成績証明書
AI で書かれたコードが本番環境に使えるかどうかを判断するための最も重要な基準は何ですか? Talking Postgres のエピソード 42 では、Datasette の作成者、Django の共同作成者、多作のオープンソース開発者である Simon Willison がクレアに加わり、今日のソフトウェア構築方法が AI によってどのように変化しているかを共有します。 AI が生成したコードを出荷するための彼のテストが「これを他の人に説明できますか?」である理由、AI エージェントを管理する際にエンジニアリング管理スキルが驚くほど役立つ可能性があること、ボトルネックがコードを書くことではなくコードを理解することであることを掘り下げます。また、個人の信頼性、綿密な調査、ウィンチェスター ミステリー ハウス、ずさんなプロキシ、そして「機能は安い。だからと言って、すべてを構築する必要があるというわけではない」というサイモンの見解も含まれます。
以前の Talking Postgres: Talking Postgres ポッドキャスト Ep 30: Simon Willison によるデータ エンジニアのための AI
プロジェクト ページ: Datasette 、データ内のストーリーを見つけるための
ブログ投稿: alchemy-utils 0.1a0 のリリース (Postgres と DuckDB を搭載)
ブログ投稿: 自然言語テキストには可逆変換はありません、Sophie Alpert 著
ポッドキャスト エピソード: 「The Open Weight Revolution with Simon Willison」、「Oxide and Friends」
ポッドキャスト エピソード: AI 一般教書、レニーのポッドキャスト
ウィキペディア: ウィンチェスター ミステリー ハウス
神話上の人月のアイデア: 概念的な整合性
ブログ投稿: SQLite 圧縮テキスト履歴プロトタイプ、Simon Willison 著
ツール: Simon のその他のツール リポジトリ
リンクトイン
マストドン
ブルースカイ
GitHub
プロデューサー
アーロン・ウィスラング
オープンソース エンジン

Microsoft + Azure でのリング + 開発者関係 ☁️ | Go (golang)、クラウドネイティブ、Linux 🐧 🐍 🦀 ☕ 🍷📷 🎹 |トロント🇨🇦🌎 | 💨😷💉 | https://aaronw.dev/hello/
×
リンクトイン
マストドン
スレッド
ブルースカイ
GitHub
ウェブサイトへのリンク
ゲスト
サイモン・ウィリソン
独立系 AI 研究者、datasette.io および llm.datasette.io の作成者、データ ジャーナリズム用のオープン ソース ツールを構築し、https://simonwillison.net/ で多くのことについて執筆しています。
×
リンクトイン
マストドン
ブルースカイ
GitHub
ウェブサイトへのリンク
ヘッドフォン
どこでも聴ける
聞いてください
アップルのポッドキャスト
スポティファイ
聞いてください
スポティファイ
曇り
聞いてください
曇り
ポケットキャスト
聞いてください
ポケットキャスト
アマゾンミュージック
聞いてください
アマゾンミュージック
YouTube
聞いてください
YouTube
その他のオプション »
YouTube
ブルースカイ
不和
Postgres、PostgreSQL、および Slonik ロゴは、カナダ PostgreSQL コミュニティ協会の商標または登録商標であり、その許可を得て使用されています。

## Original Extract

What’s your gold standard for deciding whether AI-written code is ready for production? In Episode 42 of Talking Postgres, Simon Willison—creator of Datasette, co-creator of Django, and prolific op...

Talking Postgres with Claire Giordano | How AI is changing software development with Simon Willison
Talking Postgres with Claire Giordano
close
Home
Episodes
People
Subscribe
play
pause
How AI is changing software development with Simon Willison
Download MP3
Show Notes /
Transcript
What’s your gold standard for deciding whether AI-written code is ready for production? In Episode 42 of Talking Postgres, Simon Willison —creator of Datasette, co-creator of Django, and prolific open-source developer—joins Claire to share how AI is changing the way he builds software today. We dig into why his test for shipping AI-generated code is “Could I explain this to somebody else?”, how engineering management skills can be surprisingly useful when managing AI agents, and how the bottleneck is no longer writing code but understanding it. Also: personal credibility, deep research, the Winchester Mystery House, slop proxies, and Simon’s observation that “Features are cheap. That doesn’t mean you should build them all.”
Previously on Talking Postgres: Talking Postgres podcast Ep 30: AI for data engineers with Simon Willison
Project page: Datasette , for finding stories in data
Blog post: Release of alchemy-utils 0.1a0 (featuring Postgres & DuckDB)
Blog post: There are no lossless transformations of natural-language text , by Sophie Alpert
Podcast episode: The Open Weight Revolution with Simon Willison , on Oxide and Friends
Podcast episode: An AI state of the union , on Lenny’s Podcast
Wikipedia: Winchester Mystery House
Idea in Mythical Man-Month: Conceptual integrity
Blog post: SQLite compressed text-history prototypes , by Simon Willison
Tools: Simon’s miscellaneous tools repo
LinkedIn
Mastodon
Bluesky
GitHub
Producer
Aaron Wislang
Open Source Engineering + Developer Relations at Microsoft + Azure ☁️ | Go (golang), Cloud Native, Linux 🐧 🐍 🦀 ☕ 🍷📷 🎹 | Toronto 🇨🇦🌎 | 💨😷💉 | https://aaronw.dev/hello/
X
LinkedIn
Mastodon
Threads
Bluesky
GitHub
Website Link
Guest
Simon Willison
Independent AI researcher, creator of datasette.io and llm.datasette.io, building open source tools for data journalism, writing about a lot of stuff at https://simonwillison.net/
X
LinkedIn
Mastodon
Bluesky
GitHub
Website Link
headphones
Listen Anywhere
Listen On
Apple Podcasts
Spotify
Listen On
Spotify
Overcast
Listen On
Overcast
Pocket Casts
Listen On
Pocket Casts
Amazon Music
Listen On
Amazon Music
YouTube
Listen On
YouTube
More Options »
YouTube
Bluesky
Discord
Postgres, PostgreSQL and the Slonik Logo are trademarks or registered trademarks of the PostgreSQL Community Association of Canada, and used with their permission.
