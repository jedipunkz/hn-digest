---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48557937"
title: "Show HN: Memento – Self-hosted agentic search and LLM wiki over your email"
article_title: ""
author: "georgeck"
captured_at: "2026-06-16T19:22:05Z"
capture_tool: "hn-digest"
hn_id: 48557937
score: 6
comments: 4
posted_at: "2026-06-16T16:36:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Memento – Self-hosted agentic search and LLM wiki over your email

- HN: [48557937](https://news.ycombinator.com/item?id=48557937)
- Score: 6
- Comments: 4
- Posted: 2026-06-16T16:36:07Z

## Translation

タイトル: HN を表示: Memento – メールを介した自己ホスト型エージェント検索と LLM Wiki
HN テキスト: 私たちの電子メールの受信箱には、数十年分のメッセージ (10 万から 50 万) が保存されています。これは、あなたの人生で起こったすべての重要な出来事、あなたが行ったプロジェクト、あなたが関わった人々の良い代理です。受信トレイ内のメッセージを時系列に表示すると、これらの詳細は表示されなくなります。このアーカイブを、検索して管理できる個人的な Wiki にできたらどうでしょうか?それがメメントです。このような Wiki の情報アーキテクチャについて、Memento は、人 (連絡先の CRM ビューなど)、プロジェクト (開始日と終了日によって制限されるライフ イベント)、コンセプト (常緑のトピック)、ニュースレターという 4 つの高レベルの次元を作成するという独自の見解を採用しています。受信トレイではメッセージが送信者アドレスごとにグループ化されるため、職場、個人、エイリアスのアドレスにわたって同じ人が何度も表示されます。人の次元では、Memento は決定論的アルゴリズムを使用してこれらすべてを 1 人の正規の人物に解決し、グラフ アルゴリズムを通じてその人物に関連する人々をもたらします。これら 2 つだけで、あなたのライフヒストリーからすでに入力された CRM を取得できます。 LLM は関与しません。ここから、個人的なメモを追加して各人の Wiki ページをさらに充実させ、LLM を使用して一貫した物語を作成できます。他の次元にも同じことが当てはまります。この Wiki にクエリを実行するにはどうすればよいですか?ここで、Memento はこの厳選されたデータセットに対してエージェント検索を使用します。メールは SQLite DB に保存され、FTS とベクター埋め込みを使用してインデックスが付けられ、Msgvault [1] と呼ばれる別のオープンソース プロジェクトを使用して最新の状態に保たれます。 Memento は、この DB を独自のテーブルで拡張し、正規の人物検出、接続のクラスターを見つけるためのグラフ アルゴリズムなど、さまざまなアルゴリズムの出力を保存します。

エージェント検索を効果的に行うために、Memento は基礎となる FTS、ベクトル、グラフ データを構造化された方法でツールとしてエージェントに公開します。エージェントは、get-message-details、message-cluster などの追加ツールを使用して検索をさらに絞り込むことができます。私たちが発見したのは、結果として得られる検索は、Google が公開している典型的な「Ask Gmail」検索よりもはるかに強力であるということです。 Memento によって公開されたすべての事実上の主張は、送信元の実際の電子メールにまで遡ることができます。 Wiki に追加されたメモはすべて、次世代に組み込まれます。したがって、Memento は時間の経過とともにより豊かになり、より個人的なものになります。これは私たちの個人的な電子メール アーカイブには非常にうまく機能し、Memento が明らかにできた事柄には嬉しい驚きを感じました。個人情報を公開せずにこの機能を大規模に実証するために、私たちは Memento を数百のメールボックスを含むパブリック Enron データセットに接続しました。 SQLite ストア、Go バックエンド、Next.js UI を備えた Memento は、この 5 GB データセットを簡単に処理します。これで、このアーカイブにクエリを実行し、エージェント検索を実行して、エンロン スキャンダルを自分で再発見できるようになりました。ここでデモ セットを見ることができます [2]。このアプリは、ローカルホスト上で提供され、アーカイブを読み取り専用として扱う単一のバイナリであり、ローカル モデルであってもクラウド モデルであっても、OpenAI-API 互換の LLM を指すことができます。オープンソースなので、データがどのように処理されるかを検査できます。アーカイブに接続せずに、今すぐ試してみることができます。ホストされたデモ [2] を使用するか、GitHub リリースをダウンロードして、合成ローカル アーカイブ `./memento app --demo` [1] https://www.msgvault.io [2] ホストされたデモ (エンロン データ): https://memento-demo.latentsignal.org/home デモ ビデオ: https://www.youtube.com/watch?v=Ms1KeAYCN2A プロジェクト ホーム: https://latentsignal.org/projects/memento GitHub: https

://github.com/latentsignal-org/memento 私たちは Memento の作成者、George と Ann です。

## Original Extract

Our email inboxes carry multiple decades of messages (100K-500K). This is a good proxy for all the important things that happened in your life, the projects you have done and the people that you have connected with. With the chronological view of messages in the inbox, these details remain hidden. What if we could turn this archive into a personal wiki that you can search and curate? That is Memento. For the information architecture of such a wiki, Memento takes an opinionated view of creating four high level dimensions - People (like a CRM view of your contacts), Projects (life events that are bounded by some start and end dates), Concepts (evergreen topics) and Newsletters. Inboxes group messages by sender address, so the same person shows up many times across work, personal, and alias addresses. In the People dimension, Memento resolves all of that into one canonical person using deterministic algorithms and brings the people who are related to that person through graph algorithms. With just these two, you get an already populated CRM from your life history. No LLMs involved. From here, you can further enrich each person's wiki page by adding additional personal notes and create a cohesive narrative using LLM. The same applies to other dimensions as well. How can we query this wiki? This is where Memento uses an agentic search over this curated dataset. Your emails are stored in SQLite DB, indexed using FTS and vector embeddings and kept up-to-date, using another open-source project called Msgvault [1]. Memento extends this DB with its own tables to store the output of various algorithms - canonical people discovery, graph algorithms to find clusters of connections etc. For the agentic search to be effective, Memento exposes the underlying FTS, vector and graph data in a structured way to the agent as tools. The agent can further refine the search with additional tools like get-message-details, message-cluster etc. What we discovered is that the resulting search is much more powerful that the typical ‘Ask Gmail’ search exposed by Google. Every factual claim exposed by Memento can be traced back to the real email it came from. Any additional notes added to the wiki are incorporated in the next generation. So Memento becomes richer and more personal to you over time. This worked really well for our personal email archive and we were pleasantly surprised by the things that Memento was able to uncover. In order to demonstrate this capability at scale without exposing our private info, we connected Memento to the public Enron dataset that contains hundreds of mailboxes. With the SQLite store, Go backend and Next.js UI, Memento handles this 5 GB dataset with ease. Now you can query this archive and run agentic searches to re-discover the Enron scandal yourself - you can see the demo set here [2]. The app is a single binary that serves on localhost and treats your archive as read-only, and you can point it at any OpenAI-API-compatible LLM, whether that's a local model or a cloud one. It's open source, so you can inspect how your data is handled. You can try it out today without connecting your archive - use the hosted demo [2] or download the GitHub release and run it with a synthetic local archive `./memento app --demo` [1] https://www.msgvault.io [2] Hosted demo (Enron data): https://memento-demo.latentsignal.org/home Demo video: https://www.youtube.com/watch?v=Ms1KeAYCN2A Project home: https://latentsignal.org/projects/memento GitHub: https://github.com/latentsignal-org/memento We are George and Ann, creators of Memento.

