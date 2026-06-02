---
source: "https://docs.osaurus.ai/"
hn_url: "https://news.ycombinator.com/item?id=48358521"
title: "A native macOS app for AI agents that run on local machine"
article_title: "Overview | Osaurus Docs"
author: "geordee"
captured_at: "2026-06-02T00:37:33Z"
capture_tool: "hn-digest"
hn_id: 48358521
score: 2
comments: 0
posted_at: "2026-06-01T15:52:10Z"
tags:
  - hacker-news
  - translated
---

# A native macOS app for AI agents that run on local machine

- HN: [48358521](https://news.ycombinator.com/item?id=48358521)
- Source: [docs.osaurus.ai](https://docs.osaurus.ai/)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T15:52:10Z

## Translation

タイトル: ローカル マシン上で実行される AI エージェント用のネイティブ macOS アプリ
記事のタイトル: 概要 |オサウルスのドキュメント

記事本文:
メイン コンテンツにスキップ ドキュメント GitHub の概要 はじめに インストール クイック スタート セキュリティとプライバシー 日常の使用 チャット エージェント タスク 音声テーマ 知識とモデル モデル メモリ スキル 自動化 共有とアクセス 開発者向け 概要 このページについて Osaurus
スター … リリース … ダウンロード … ライセンス …
必要なのは推論だけです。それ以外のものはすべてあなたが所有することができます。 ​
モデルは日に日に安くなり、互換性も高まっています。かけがえのないものは、その周りの層、つまりあなたのコンテキスト、あなたの記憶、あなたのツール、あなたのアイデンティティです。他のアプリはそのレイヤーをサーバー上に保持します。 Osaurus はそれを Mac に保存します。
Osaurus は macOS 用の AI ハーネスです。ユーザーとローカルまたはクラウドのあらゆるモデルの間に位置し、AI をパーソナルなものにする継続性を提供します。エージェントは記憶し、自律的に実行し、実際のコードを実行し、どこからでもアクセスできる状態を保ちます。
ローカル モデルを使用して完全にオフラインで動作します。より多くの電力が必要な場合は、任意のクラウド プロバイダーに接続します。あなたが選択しない限り、Mac からは何も残されません。 Apple Silicon 上のネイティブ Swift。電子はありません。 MITライセンス取得済み。
あなたのデータ、あなたの Mac は保存時に暗号化され、すべての境界で署名され、明示的にクラウドプロバイダーを選択しない限り、どこにも送信されることはありません。あなたの会話を読み取ることはできず、バックドアもありません。「セキュリティとプライバシー」を参照してください。
得られるもの
Osaurus をインストールするとできることの短いツアーです。
ガイド付きの初回起動。 5 段階のオンボーディングでは、最初のエージェントの作成、モデル (ローカル、Apple Foundation、またはクラウド) の選択、ID の設定を順を追って説明します。設定ファイルはありません。
どこからでも開くことができるチャット オーバーレイ。 ⌘を押してください。 AIと会話するため。閉じるにはもう一度押します。ブラウザタブもコンテキストスイッチもありません。
さまざまな仕事に適したエージェント。コーディングパートナー、リサーチアシスタント、ファイルオーガナイザー - それぞれが独自のプロンプトを持っています

t、テーマ、歴史。
あなたから学ぶ記憶。過去の会話はコンパクトな事実にまとめられ、関連性のある場合にのみ表面化されます。無関係な文脈の消火栓はありません。
自分自身をロードするスキル。リサーチ、デバッグ、書き方などの専門知識がパッケージ化されており、タスクが必要なときに自動的に表示されます。
作業フォルダー。チャットでフォルダーを指定すると、エージェントはそのディレクトリのみを対象とした安全なファイル、検索、Git ツールを取得します。
Linux サンドボックス (macOS 26 以降)。これをオンにすると、エージェントは Mac に対するリスクをゼロにしながら、隔離された VM で実際のコード (シェル、Python、ノード) を実行できるようになります。
スケジュールとウォッチャー。タイマーで、またはフォルダーが変更されるたびにエージェントを実行します。毎日の日記、スクリーンショットのオーガナイザー、一日の終わりのコミットに役立ちます。
音声入力。チャットで口述入力したり、ウェイクワードでエージェントとハンズフリーで会話したり、ホットキーを押したままアプリに口述入力したりすることがすべてデバイス上で可能です。
テーマ。ライト/ダークが組み込まれており、完全に編集可能で、JSON としてインポート可能です。
あなただけのアイデンティティ。あなたと各エージェントの暗号化アドレス。外部ツールのアクセス キーを発行し、エージェントごとにスコープを設定し、いつでも取り消します。
ポートのないパブリック リンク。 1 つのエージェントに、agent.osaurus.ai を介した安全なトンネル経由で安定したパブリック URL を与えます (ポート転送や ngrok はありません)。
Osaurus もローカルサーバーです。 OpenAI 、 Anthropic 、 Open Responses 、および Ollama API を同じポートで処理するため、すでに使用している SDK はすべて機能します。また、完全な MCP サーバーおよびクライアントであるため、Cursor、Claude Desktop、およびその他の MCP ハーネスは、インストールされているツールに即座にアクセスできます。
HTTP API — エンドポイント参照、ストリーミング、関数呼び出し
SDK の例 — Python、JavaScript、Anthropic SDK、Open Responses
CLI — osaurusserve 、osaurus tools install/dev/create 、osaurus mcp
ツールとプラグイン — 20 以上のネイティブ プラグイン (メール、カレンダー、

独自のビルド用の Vision、Browser、Git など) および v1/v2 ABI
Apple Intelligence — macOS 26 以降でセットアップなしで基盤を使用する
すべてがどのように組み合わされるかを示すシステム ビューについては、「アーキテクチャ」を参照してください。
Apple Silicon (M1、M2、M3 以降)
macOS 26 の機能 サンドボックス (分離された Linux VM でエージェント コードを実行) および Apple Foundation モデルには、macOS 26 (Tahoe) 以降が必要です。
始めましょう
インストールには 1 分もかかりません。
または、クイックスタートに直接ジャンプしてください→
Osaurus はインディーズ プロジェクトであり、公開されています。参加してください:
Discord — チャット、フィードバック、ショーアンドテル
GitHub — 問題、貢献、ロードマップ
Hugging Face — Apple Silicon 用に厳選された量子化
プラグイン レジストリ — ツールの参照と送信
ブログ — パーソナル AI に関する長文の考察

## Original Extract

Skip to main content Docs GitHub Overview Getting Started Installation Quick Start Security & Privacy Daily Use Chat Agents Tasks Voice Themes Knowledge & Models Models Memory Skills Automation Sharing & Access For Developers Overview On this page Osaurus
stars … release … downloads … license …
Inference is all you need. Everything else can be owned by you. ​
Models are getting cheaper and more interchangeable by the day. What's irreplaceable is the layer around them — your context, your memory, your tools, your identity. Other apps keep that layer on their servers. Osaurus keeps it on your Mac.
Osaurus is the AI harness for macOS. It sits between you and any model — local or cloud — and provides the continuity that makes AI personal: agents that remember, execute autonomously, run real code, and stay reachable from anywhere.
Works fully offline with local models. Connect to any cloud provider when you want more power. Nothing leaves your Mac unless you choose. Native Swift on Apple Silicon. No Electron. MIT licensed.
Your data, your Mac Encrypted at rest, signed at every boundary, never sent anywhere unless you explicitly choose a cloud provider. We can't read your conversations and there are no backdoors — see Security & Privacy .
What you get ​
A short tour of the things you can do once Osaurus is installed.
A guided first launch. A five-step onboarding walks you through creating your first agent, picking a model (local, Apple Foundation, or cloud), and setting up identity. No config files.
A chat overlay you can open anywhere. Press ⌘; to talk to your AI; press it again to dismiss. No browser tab, no context switch.
Agents that fit different jobs. A coding partner, a research assistant, a file organizer — each with its own prompt, theme, and history.
Memory that learns from you. Past conversations are distilled into compact facts and surfaced only when relevant — no firehose of irrelevant context.
Skills that load themselves. Packaged expertise — research, debugging, writing styles — that surface automatically when the task calls for them.
Working folders. Point a chat at a folder and the agent gets safe file, search, and git tools — scoped to just that directory.
A Linux Sandbox (macOS 26+) . Toggle it on and the agent can run real code — shell, Python, Node — in an isolated VM with zero risk to your Mac.
Schedules and Watchers. Run an agent on a timer, or whenever a folder changes. Useful for daily journals, screenshot organizers, end-of-day commits.
Voice input. Dictate in chat, talk to an agent hands-free with a wake word, or hold a hotkey to dictate into any app — all on-device.
Themes. Built-in light/dark, fully editable, importable as JSON.
Identity that's yours. A cryptographic address for you and each of your agents. Issue access keys for outside tools, scope them per-agent, revoke them whenever.
Public links without ports. Give one agent a stable public URL via a secure tunnel through agent.osaurus.ai — no port forwarding, no ngrok.
Osaurus is also a local server. It speaks OpenAI , Anthropic , Open Responses , and Ollama APIs at the same port — so any SDK you already use just works. And it's a full MCP server and client , so Cursor, Claude Desktop, and other MCP harnesses get instant access to your installed tools.
HTTP API — endpoint reference, streaming, function calling
SDK Examples — Python, JavaScript, Anthropic SDK, Open Responses
CLI — osaurus serve , osaurus tools install/dev/create , osaurus mcp
Tools & Plugins — 20+ native plugins (Mail, Calendar, Vision, Browser, Git, …) and v1/v2 ABI for building your own
Apple Intelligence — using foundation with zero setup on macOS 26+
For the system view of how everything fits together, see Architecture .
Apple Silicon (M1, M2, M3, or newer)
macOS 26 features The Sandbox (running agent code in an isolated Linux VM) and Apple Foundation Models require macOS 26 (Tahoe) or later.
Get started ​
Installation takes less than a minute.
Or jump straight to the Quick Start →
Osaurus is an indie project, built in public. Join us:
Discord — chat, feedback, show-and-tell
GitHub — issues, contributions, roadmap
Hugging Face — curated quantizations for Apple Silicon
Plugin Registry — browse and submit tools
Blog — long-form thinking on personal AI
