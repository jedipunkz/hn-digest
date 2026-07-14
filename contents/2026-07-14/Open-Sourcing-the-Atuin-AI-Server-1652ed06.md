---
source: "https://blog.atuin.sh/atuin-ai-oss/"
hn_url: "https://news.ycombinator.com/item?id=48910998"
title: "Open Sourcing the Atuin AI Server"
article_title: "Open Sourcing the Atuin AI Server"
author: "johnbehnke"
captured_at: "2026-07-14T19:03:02Z"
capture_tool: "hn-digest"
hn_id: 48910998
score: 4
comments: 1
posted_at: "2026-07-14T18:21:23Z"
tags:
  - hacker-news
  - translated
---

# Open Sourcing the Atuin AI Server

- HN: [48910998](https://news.ycombinator.com/item?id=48910998)
- Source: [blog.atuin.sh](https://blog.atuin.sh/atuin-ai-oss/)
- Score: 4
- Comments: 1
- Posted: 2026-07-14T18:21:23Z

## Translation

タイトル: Atuin AI サーバーのオープンソース化
説明: Atuin AI 用に独自のサーバーをセルフホストします。端末内に常駐し、ユーザーの動作を認識するエージェントです。

記事本文:
サインイン
購読する
Atuin AI サーバーのオープンソース化
Atuin AI 用に独自のサーバーをセルフホストします。Atuin AI は、端末内に常駐し、ユーザーの動作を認識するエージェントです。
Atuin AI は、シェル内にある高速な端末中心のエージェントです。即座に起動し、端末での作業に必要なすべてのエージェント ツールが付属しています。
0:00
/ 0:30
1×
Atuin はすべてのユーザーに無料で寛大な使用予算を提供しており、AI での会話は保存されませんが、データ、特に端末の使用に起因するデータをより厳密に監視したいと考える人もいることは認識しています。
そのために、私たちは Atuin AI サーバーをオープンソース化できることを嬉しく思います。
Atuin AI サーバーは、GitHub の atuinsh/atuin-ai-server にあります。これは、実稼働 Atuin AI サーバーに使用しているものと同じライブラリである atuinsh/atuin-ai-core に基づいています。
Atuin AI サーバーは現在、OpenAI 互換のチャット完了スタイルのエンドポイントをサポートしています。ローカル モデルの場合、これには Ollama、vLLM、LM Studio、llama.cpp、LiteLLM などが含まれます。 OpenRouter などの OpenAI 互換 Web サービスを使用することもできます。
リポジトリのクローンを作成した後、サンプル構成ファイル config.example.toml を config.toml にコピーします。 Readme の構成セクションに従って、インスタンスをセットアップします。
Ollama ベースのセットアップの非常に基本的な例を次に示します。
ポート = 8080
endpoint = "http://localhost:11434/v1" # または host.docker.internal
api_key = "オラマ"
デフォルト_モデル = "ラマ31"
[リクエスト.ボディ]
stream_options = { include_usage = true }
[[モデル]]
別名 = "ラマ31"
名前 = "ラマ 3.1 70b"
説明 = "オラマ ラマ 3.1 70b"
モデル = "ラマ3.1:70b"
[[モデル]]
エイリアス = "gemma4"
名前 = "ジェマ 4 r4b"
description = "オラマ ジェマ 4 - 有効 4b"
model = "gemma4:e4b" サーバー側ツールの構成など、セットアップの詳細については、リポジトリの Readme を参照してください。

Web検索とWebコンテンツのスクレイピング。
完了したら、次の 2 つの方法のいずれかでサーバーを起動できます。
Erlang、Elixir、Gleam がインストールされている場合は、サーバーをネイティブに実行できます。
ミックスdeps.get
mix run --no-halt config.toml で環境変数を介して API キーを指定する場合は、サーバーの起動時に忘れずに API キーを設定してください。
docker を使用してサーバーを実行するには、次のコマンドを実行します。
ドッカーの実行 \
-v ./config.toml:/etc/atuin-ai/config.toml \
-p 8080:8080 \
ghcr.io/atuinsh/atuin-ai-server:latest Docker 経由で実行していて、Ollama などのホスト上で実行されているローカル LLM サービスに Atuin AI サーバーを接続したい場合は、localhost (コンテナー独自のループバック インターフェイスに解決される) の代わりに host.docker.internal をエンドポイントとして使用します。
サーバーが実行されたら、エンドポイント config を設定してサーバーに接続するように Atuin AI を構成できます。
[あい]
エンドポイント = "http://localhost:8080"
続きを読む
Atuin v18.13 – シェルの検索機能、PTY プロキシ、および AI の向上
新しいリリースが出ました！ v18.13 はおそらく、私たちがここしばらくリリースした中で最大の変更セットです。詳細については、以下をお読みください。
デーモンを使用すると、はるかに高速かつ優れた検索が可能になります
このデーモンは長い間存在しており、「実験的」とマークされています。
Atuin TUI のカスタム キーバインド
Atuin に対する最も長年の機能リクエストの 1 つである、検索 TUI の完全なカスタム キーバインドのサポートがついに実現しました。デフォルト設定にこだわる必要も、ctrl-d で何か他のことをしてほしかったという必要ももうありません。あなたのTUI、あなたのルール。
Atuin 開発ブログ: デスクトップ AI、UI + パフォーマンスの改善、SSH の修正など
Atuin Desktop の新機能を詳しく説明する月次投稿の第 1 回へようこそ。この投稿は、2025 年と 2026 年の最初の数週間の両方の多くの内容を取り上げているため、少しユニークです。
AIアシスタント
エージェントが変更したエンギの数

オタクが書く
新しい Runbook 実行エンジンの紹介
私たちは、Atuin Desktop に対する大幅なアーキテクチャの改善、つまり完全に再設計された Runbook 実行エンジンを発表できることを嬉しく思います。
これは大きな変更であり、Runbook を自動化の中核プリミティブにするための最初の大きな一歩です。
不安定なコンテキスト、消失する状態、または一貫性のない実行に遭遇したことがある場合は、このリリースで修正されます
魔法のシェル ツール Atuin に関する最新情報とニュース

## Original Extract

Self host your own server for Atuin AI - the agent that lives in your terminal, and knows how you work

Sign in
Subscribe
Open Sourcing the Atuin AI Server
Self host your own server for Atuin AI - the agent that lives in your terminal, and knows how you work
Atuin AI is a fast terminal-focused agent right in your shell. It starts instantly and comes with all the agentic tools needed to help with work in your terminal.
0:00
/ 0:30
1×
Atuin offers a generous usage budget for free for all users, and we don't store your AI conversations, but we recognize some people would prefer to keep tighter tabs on their data, especially data that originates from terminal usage.
To that end, we're happy to be open sourcing the Atuin AI server .
The Atuin AI server can be found on GitHub at atuinsh/atuin-ai-server ; it's based on atuinsh/atuin-ai-core , the same library we use for the production Atuin AI server.
The Atuin AI server currently supports any OpenAI-compatible, chat completions-style endpoint. For local models, this includes Ollama, vLLM, LM Studio, llama.cpp, and LiteLLM, among others. You can also use OpenAI-compatible web services, like OpenRouter.
After cloning the repository, copy the example config file, config.example.toml , to config.toml . Follow the configuration section of the readme to set up your instance.
Here's a very basic example of an Ollama-based setup:
port = 8080
endpoint = "http://localhost:11434/v1" # or host.docker.internal
api_key = "ollama"
default_model = "llama31"
[request.body]
stream_options = { include_usage = true }
[[models]]
alias = "llama31"
name = "Llama 3.1 70b"
description = "Ollama Llama 3.1 70b"
model = "llama3.1:70b"
[[models]]
alias = "gemma4"
name = "Gemma 4 r4b"
description = "Ollama Gemma 4 - Effective 4b"
model = "gemma4:e4b" See the repository readme for more setup details, including configuring server-side tools, like web search and web content scraping.
Once done, you can start the server one of two ways:
If you have Erlang, Elixir, and Gleam installed, you can run the server natively:
mix deps.get
mix run --no-halt If your config.toml specifies API keys via environment variables, remember to set them when you start the server.
To run the server with docker, run the following:
docker run \
-v ./config.toml:/etc/atuin-ai/config.toml \
-p 8080:8080 \
ghcr.io/atuinsh/atuin-ai-server:latest If you're running via Docker and want the Atuin AI server to connect with a local LLM service running on the host, like Ollama, use host.docker.internal as the endpoint instead of localhost (which would resolve to the container's own loopback interface).
Once your server is running, you can configure Atuin AI to connect to it by setting the endpoint config :
[ai]
endpoint = "http://localhost:8080"
Read more
Atuin v18.13 – better search, a PTY proxy, and AI for your shell
A new release is out! v18.13 is probably the biggest set of changes we have released in a good while, read on to find out more.
Much faster and better search with the daemon
The daemon has existed for a long time, and has been marked as "experimental&
Custom Keybindings for the Atuin TUI
One of the longest-standing feature requests for Atuin has finally landed: full custom keybinding support for the search TUI. No more being stuck with our defaults, no more wishing ctrl-d did something else. Your TUI, your rules.
Atuin Devlog: Desktop AI, UI + performance improvements, SSH fixes, and more
Welcome to the first installment of a monthly post detailing what's new in Atuin Desktop. This post is a little unique as we're covering a bunch of stuff from both 2025, and the first few weeks of 2026.
AI Assistant
Agents changed how many engineers write
Introducing the New Runbook Execution Engine
We're excited to announce a major architectural improvement to Atuin Desktop: a completely redesigned runbook execution engine.
This is a huge change, the first big step toward making runbooks a core automation primitive.
If you’ve ever hit flaky context, disappearing state, or inconsistent execution, this release fixes
Updates and news about Atuin, the magical shell tool
