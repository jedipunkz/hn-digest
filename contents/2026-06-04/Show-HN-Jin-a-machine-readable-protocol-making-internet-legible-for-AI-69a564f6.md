---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48397992"
title: "Show HN: Jin – a machine-readable protocol making internet legible for AI"
article_title: ""
author: "ankushvishnu"
captured_at: "2026-06-04T13:13:28Z"
capture_tool: "hn-digest"
hn_id: 48397992
score: 1
comments: 0
posted_at: "2026-06-04T12:55:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Jin – a machine-readable protocol making internet legible for AI

- HN: [48397992](https://news.ycombinator.com/item?id=48397992)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T12:55:03Z

## Translation

タイトル: 番組 HN: Jin – AI がインターネットを判読できるようにする機械可読プロトコル
HN テキスト: こんにちは、HN!私は、AI エージェントが日常のインターネットを読み取れるようにする、エージェント ビルダーとウェブマスター向けのオープンソース プロトコルである Jin を構築しています。私はコーディングと研究に AI エージェントを使用していますが、私が観察したことの 1 つは、エージェントにスクレイピング権限を与えるのに多くの時間、お金、あるいはその両方を費やしているということです。そこで私は、AI エージェントが汗をかかずにインターネットから欲しいものを届けてくれるシステムがあったらどうなるだろうかと考えました。そして、それが私にインテントレイヤーのアイデアを与えました。ジンの仕組みは次のとおりです -
1. AI エージェント側: skill.json または skill.xml 形式でエージェントに指示を渡すだけです。それが最初のステップです。このスキルは、Web サイト上で jin.json を探し、インテント マップを使用して移動し、CRUD などのアクションを実行するようにエージェントに指示します。このスキルは app/.well-known/jin.json に存在します。 2. Web マスター側: Web マスターはプロジェクトのルートで jin-cli を実行でき、cli はプロジェクトをマップし、jin.json ファイルにインテント マップを生成し、アプリが存在する場所に保存します。たとえば、モノリシック構造では、jin.json がルートに存在できます。モノリポジトリ構造の場合、app/web または app/marketing などに置くことができます。このファイルは必要に応じて完全に編集できます。インテント マップの生成とは別に、簡単に編集する方法、何を配置できるか、何を入れるべきか、何を入れてはいけないのかなどを示します。 jin.json には To Do リストがあります。これに加えて、最新バージョンでは、不正なスクレーパーから Web サイトを保護するための Jin シールドを追加しました。このシールドは、使用しているファイアウォールの隣にあります。これを防御の第一線のように考えてください。誰が入ってきたのかを密告します。meetjin.com にエージェントを無料で登録し、生成されたキーを使用して、Jin シールドを使用する Web サイトにアクセスできます。エージェントの車

暗号化 RS256 JWT パスポートと、Jin シールドが、キャッシュされた JWKS 公開キーを使用してローカルでこれを検証します。仕様はオープン (CC0) で、ツールは Apache 2.0 です。使って、壊して、焼いて褒めてください。 cli にアクセスします: npm i @papercargo/jin-cli
キーにアクセスします: https://meetjin.com
Github: https://github.com/meetjin/jin 使用例やフィードバックをお聞かせいただければ幸いです。乾杯！

## Original Extract

Hi HN! I'm building Jin - an open-source protocol for agent builders and webmasters to make the everyday internet legible for AI agents. I use AI agents for coding & research and one thing I've observed is that we spend to much time, money or both on giving scraping powers to our agents. That got me thinking, what if there was a system for my AI agent to deliver what I want from the internet without having to break a sweat? And that gave me the intent layer idea. Here's how Jin works -
1. AI agent side: you only pass the instructions to your agent in a skill.json or skill.xml format. That's the first step. The skill will tell the agent to look for jin.json on the website and use the intent map to navigate, take action like CRUD, etc. It lives under app/.well-known/jin.json 2. Webmaster side: the webmasters can run the jin-cli at the root of their project, the cli will map the project, generate intent map in a jin.json file and save it wherever the app lives. For example, in a monolithic structure, the jin.json can live at the root. For a monorepo structure, it can live under app/web or app/marketing, etc. This file can be totally edited the way you want. Apart from generating intent maps, it tells you how you can easily edit, what you can put, what should go & what not, etc. There's a to-do list in jin.json Along with this, in the latest version, I've added a Jin shield to protect websites from unauthorised scrapers. This shiled lives next to whatever firewall you are using, consider this like a first line of defence. It will snitch who came in. You can register your agent on meetjin.com for free and use the generated key to access any website that uses Jin shield. The agents carry a cryptographic RS256 JWT passport and the Jin shield verifies this locally using cached JWKS public keys. The specification is open (CC0), the tooling is Apache 2.0. Use it, break it, roast it praise it. Access the cli: npm i @papercargo/jin-cli
Access the key: https://meetjin.com
Github: https://github.com/meetjin/jin I would really appreciate to know your use case and feedback. Cheers!

