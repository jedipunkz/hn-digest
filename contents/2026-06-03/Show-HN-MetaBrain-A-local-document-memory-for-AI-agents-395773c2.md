---
source: "https://metabrain.eu"
hn_url: "https://news.ycombinator.com/item?id=48372976"
title: "Show HN: MetaBrain – A local document memory for AI agents"
article_title: "metaBrain - open-source local memory for AI agents"
author: "acoye"
captured_at: "2026-06-03T00:40:21Z"
capture_tool: "hn-digest"
hn_id: 48372976
score: 6
comments: 1
posted_at: "2026-06-02T17:03:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: MetaBrain – A local document memory for AI agents

- HN: [48372976](https://news.ycombinator.com/item?id=48372976)
- Source: [metabrain.eu](https://metabrain.eu)
- Score: 6
- Comments: 1
- Posted: 2026-06-02T17:03:09Z

## Translation

タイトル: Show HN: MetaBrain – AI エージェントのためのローカル ドキュメント メモリ
記事のタイトル: metaBrain - AI エージェント用のオープンソース ローカル メモリ
説明: metaBrain は、検索可能なメモ、メタデータ、タグ、リンク、バージョン履歴を備えた、AI ツール、コーディング エージェント、および人間のためのオープンソースのローカル ドキュメント メモリです。
HN テキスト: こんにちは、HN 私は最近エージェント コーディングを実験し、プロジェクトごとにより多くのコンテキスト データを追跡する必要性を感じました。
また、1D チャットを超えてエージェントとコミュニケーションできる必要性も感じました。そこで、エージェント自身が検出できるローカル ドキュメント メモリを作成しました。
CLI は、エージェントが簡単に使用できるように設計されています。
人間もストア内のドキュメントを読み取り、検索、編集することで共同作業することができます。現在、Mac ネイティブ GUI をレビュー中です。近いうちに App Store に表示されることを願っています。簡単に試すことができます。手順はこちら: https://metabrain.eu/
GitHub はこちらです https://github.com/OpenCow42/metaBrain このプロジェクトは、私にとって、真のクロスプラットフォーム (Mac / Linux / Windows) で迅速なプロジェクトを構築するための実験でもあります。
このプロジェクトを実行するために Swift でラップした LevelDB と同じライセンスでオープンソース化されています。エージェント (および人間) は、検索によってコンテンツを迅速に取得できるため、エージェントの作業中に特定のコンテキストで特定の知識を再注入できます。
面白いことに、私は「推論ルールベース」を、古い関数エキスパート システムの廃れたアイデアのようなものだと考えていました。
エージェントと協力し始めた今、そのようなベースで以前に機能していたソリューションを動的に選択しに行く必要性をますます感じています。フィードバックをいただければ幸いです。
製品のフィット感については、これはあなたにとって役に立ちますか、それとも満足しているのは私だけでしょうか?最後に、ドキュメントの圧縮を楽しみました。圧縮できない場合は、ZSTD をすばやく試します。

データを 10 パーセント以上圧縮せずに保存するか、データに対して ZSTD レベル 9 圧縮を実行します。私はこのトリックを OpenZFS から学びました。ありがとう

記事本文:
コンテンツにスキップ
MB
メタブレイン
インストール
エージェント
特長
オープンソース
GitHub
オープンソース。まずは地元。エージェントの準備ができました。
AI ツール、コーディング エージェント、人間のためのローカル ドキュメント メモリ
彼らと一緒に働く人たち。
すべてのエージェントにメモやソースを保存できる耐久性のある検索可能な場所を 1 つ提供する
スニペット、タスク コンテキスト、メタデータ、タグ、リンク、バージョン履歴。
これをインストールして、エージェントに共有メモリを与えます。
metaBrain は、`mb` CLI、オプションの `mbd` ローカル デーモンとして出荷されます。
そして埋め込み可能な`MetaBrainCore` Swiftライブラリ。
brew タップ OpenCow42/tap && brew install mb
AI とエージェントの両方のために作られています
エージェントが実際に使用できる記憶層。
エージェントは、すべてのタスクに対してプライベート ファイル規則を必要としません。
metaBrain は、予測可能なパスを備えた安定したローカル ストアを提供します。
検索可能なコンテンツ、タグ、メタデータ、参照、および保持
バージョン。
概要、決定事項、タスクの状態を次のエージェントが見つけられる場所に保存します。
mb put /tasks/release-checklist \
「初公開の準備をします。」 \
--tag release --meta status=active
02
ワークスペースメモリの検索
スクラッチ ファイルを検索する代わりに、単語、パス、タグ、またはメタデータによってコンテキストを取得します。
mb 検索「公開リリース」 -- タグリリース
03
書き換えなしでパッチを適用する
保存されているドキュメントに集中した統合差分を適用し、古いバージョンを利用可能な状態に保ちます。
mb patch /tasks/release-checklist --patch-filechange.diff
なぜ試してみるのか
小規模な CLI。便利なメモリモデル。
metaBrain は、シェル スクリプトに適したプレーンなインターフェイスを維持し、
長時間実行されるエージェント ワークフローに十分な機能が豊富です。
デフォルトのストアは `.metabrain/store.leveldb` にあるため、ツールはセットアップの儀式なしでそれを検出できます。
現在のドキュメント チャンクには、タグ、メタデータ、およびパス フィルターを使用した字句検索用のインデックスが付けられます。
ドキュメントには、メタデータ、タグ、参照、保持されたバージョン、およびファイルシステムのようなパスが含まれています。
更新ではスナップショットが保持され、`

patch` は、保存されているドキュメント本文に統合された差分を適用できます。
「mbd」は 1 つのデーモンを通じて複数のローカル metaBrain データベースにサービスを提供できるため、エージェントは別々のストアを同時に操作できます。
「MetaBrainCore」は、将来のツールやインターフェースのために共有動作を Swift ライブラリに保持します。
クリーンなワークスペースで CLI を試してください。
ストアを作成し、メモを書き、ツリーを参照し、検索し、読み取ります
体を後ろに。ツールの形状を感じるには十分です。
MB初期化
mb put /notes/today 「語彙ストアを思い出してください。」 --タグプランニング --metasource=agent
mb リスト /notes --recursive --dates
mb ツリー --最大深さ 2
mb 検索「語彙ストア」 -- タグ企画
mb 取得 /notes/today
オープンソース
ローカルファーストのエージェント作業のためにオープンに構築されています。
このプロジェクトは BSD 3 条項ライセンスを取得しており、GitHub でホストされています。
現在、macOS、Linux、および Windows をサポートしています。
metaBrain: AI エージェントと人間の協力者のためのローカル メモリ。

## Original Extract

metaBrain is open-source local document memory for AI tools, coding agents, and humans, with searchable notes, metadata, tags, links, and version history.

Hello there HN I experimented with agentic coding recently and I felt the need to track more contextual data by project.
Also I felt the need to be able to go beyond the 1D chat to communicate with agents. So I created a local document memory, that is discoverable by agents themselves.
The CLI is designed to be easy to pick up by agents.
It allows humans to collaborate too by reading / searching / editing documents in the store. I have a Mac native GUI in the review process, I hope it will show up in the App Store soon. You can try it easily, instructions here: https://metabrain.eu/
Here is the GitHub https://github.com/OpenCow42/metaBrain The project is also an experiment for me to build some swift project truly cross platform (Mac / Linux / Windows)
It is open-sourced with the same license as LevelDB that I wrapped in swift to do this project. The agents (and humans) can retrieve content quickly with a search, allowing to re-injecting specific knowledge in a specific context during agentic work.
It’s funny, I’ve thought of "inference rule base" as something of a derelict idea of the old functional expert systems.
Now that I start working with agents I feel more and more the need to go pick previously working solutions dynamically in such a base. I’d be happy to get feedback.
Product fit wise, would this be useful to you or is this just me who is happy with it ? Finally I had fun with the compression of documents, it tries ZSTD quick, if it does not compress the data by more than 10 percent it stores data uncompressed, else it does a ZSTD level 9 compression on the data. I picked up this trick form OpenZFS. Thanks

Skip to content
mb
metaBrain
Install
Agents
Features
Open source
GitHub
Open source. Local first. Agent ready.
Local document memory for AI tools, coding agents, and the humans
who work with them.
Give every agent one durable, searchable place for notes, source
snippets, task context, metadata, tags, links, and version history.
Install it, then give agents a shared memory.
metaBrain ships as the `mb` CLI, the optional `mbd` local daemon,
and the embeddable `MetaBrainCore` Swift library.
brew tap OpenCow42/tap && brew install mb
Made for AI and agents alike
A memory layer agents can actually use.
Agents do not need a private file convention for every task.
metaBrain gives them a stable local store with predictable paths,
searchable content, tags, metadata, references, and retained
versions.
Store summaries, decisions, and task state where the next agent can find them.
mb put /tasks/release-checklist \
"Prepare first public release." \
--tag release --meta status=active
02
Search workspace memory
Retrieve context by words, paths, tags, or metadata instead of hunting through scratch files.
mb search "public release" --tag release
03
Patch without rewriting
Apply focused unified diffs to stored documents and keep the old versions available.
mb patch /tasks/release-checklist --patch-file change.diff
Why try it
Small CLI. Useful memory model.
metaBrain keeps the interface plain enough for shell scripts and
rich enough for long-running agent workflows.
The default store lives at `.metabrain/store.leveldb`, so tools can discover it without setup ceremony.
Current document chunks are indexed for lexical search with tag, metadata, and path filters.
Documents carry metadata, tags, references, retained versions, and filesystem-like paths.
Updates keep snapshots, and `patch` can apply unified diffs to stored document bodies.
`mbd` can serve multiple local metaBrain databases through one daemon, so agents can work with separate stores concurrently.
`MetaBrainCore` keeps shared behavior in a Swift library for future tools and interfaces.
Try the CLI in a clean workspace.
Create a store, write a note, browse the tree, search it, and read
the body back. That is enough to feel the shape of the tool.
mb init
mb put /notes/today "Remember the lexical store." --tag planning --meta source=agent
mb list /notes --recursive --dates
mb tree --max-depth 2
mb search "lexical store" --tag planning
mb get /notes/today
Open source
Built in the open for local-first agent work.
The project is BSD 3-Clause licensed, hosted on GitHub, and
currently supports macOS, Linux, and Windows.
metaBrain: local memory for AI agents and human collaborators.
