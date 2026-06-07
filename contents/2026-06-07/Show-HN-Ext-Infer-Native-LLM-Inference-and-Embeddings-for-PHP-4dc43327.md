---
source: "https://infer.displace.tech"
hn_url: "https://news.ycombinator.com/item?id=48431237"
title: "Show HN: Ext-Infer – Native LLM Inference and Embeddings for PHP"
article_title: "Introduction - ext-infer"
author: "eamann"
captured_at: "2026-06-07T04:34:39Z"
capture_tool: "hn-digest"
hn_id: 48431237
score: 1
comments: 0
posted_at: "2026-06-07T02:39:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ext-Infer – Native LLM Inference and Embeddings for PHP

- HN: [48431237](https://news.ycombinator.com/item?id=48431237)
- Source: [infer.displace.tech](https://infer.displace.tech)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T02:39:24Z

## Translation

タイトル: Show HN: Ext-Infer – PHP のネイティブ LLM 推論と埋め込み
記事のタイトル: はじめに - ext-infer
説明: PHP のローカル LLM 推論、プロセス中。 ext-infer 拡張機能のドキュメント。

記事本文:
← または → を押して章間を移動します
S または / を押して書籍内を検索します
ext-infer は、GGUF をロードする PHP 8.3 以降の拡張機能です。
モデルを作成し、PHP プロセス内で LLM 推論を実行します。
ラマ.cpp 。 PHPネイティブのセマンティクス
検索、RAG パイプライン、および CLI / ワーカー推論はシェルなしで実行されます
Python に出力するか、リモート API にアクセスします。
Rustで書かれています
ext-php-rs と
ラマ-cpp-2バインディング。一般の人々
PHP サーフェスは、ネイティブに感じられるように設計されており、流暢で役割を認識したプロンプトを提供します。
ビルダー;推論と答えを分ける応答。埋め込み
それ自体を正規化し、コサイン類似度を計算する方法を知っています。あなた
<|im_start|> トークンについて考える必要があるのは、たとえあったとしてもほとんどありません。
ローカル推論が PHP の隣ではなく PHP に属する 3 つの理由:
待ち時間。サブプロセス フォークまたは HTTP ラウンドトリップは少なくとも
ミリ秒、多くの場合は数十。インプロセス呼び出しは次によってのみ制限されます。
デコード時間。
操作面。パッケージ化する Python サイドカーやデーモンは不要
監視するため、FPM と並行して拡張する推論サーバーがありません。 PHP
プロセスは推論サーバーです。
API の人間工学。ローカル LLM の呼び出しは、PHP と同様に自然であるべきです
intl または pdo の呼び出しとして。拡張機能 API は次のような形状になっています。
それ — プロンプトと
チャットの完了。
このガイドは 5 つのレイヤーに分かれており、サイドバーから移動できます。
ext-infer はプレリリースです。クラス サーフェスは安定していますが、
最初にタグ付けされたリリース ( v0.1.0 ) はまだ開発中です。参照
リリース.md
カット・ア・リリース・フローと PLAN.md 用
次に何が起こるかについて。
コード ブロックは、1 つの例外を除いて、書かれたとおりに実行できます: PHP コード
拡張機能がロードされていることを前提としています。システム全体にインストールするか、
php コマンドの先頭に -d extension=… を追加します。参照
インストール。
名前空間プレフィックスのないモデルは、 Displace\Infer\Model を意味します。
Prompt 、 Response 、 Embedding についても同様です。本物

コードを使用する必要があります
ファイルの先頭にあるステートメント。
CLI スニペットは、POSIX シェル (bash / zsh) 用に作成されます。調整する
必要に応じて魚/PowerShell の場合。通常、相違点は引用のみです。

## Original Extract

Local LLM inference for PHP, in-process. Documentation for the ext-infer extension.

Press ← or → to navigate between chapters
Press S or / to search in the book
ext-infer is a PHP 8.3+ extension that loads a GGUF
model and runs LLM inference inside the PHP process via
llama.cpp . PHP-native semantic
search, RAG pipelines, and CLI / worker inference run without shelling
out to Python or hitting a remote API.
It is written in Rust on top of
ext-php-rs and the
llama-cpp-2 bindings. The public
PHP surface is designed to feel native: a fluent, role-aware Prompt
builder; a Response that splits reasoning from answer; an Embedding
that knows how to normalize itself and compute cosine similarity. You
should rarely, if ever, need to think about <|im_start|> tokens.
Three reasons local inference belongs in PHP rather than next to it:
Latency. A subprocess fork or HTTP roundtrip is at least
milliseconds, often tens. An in-process call is bounded only by
decode time.
Operational surface. No Python sidecar to package, no daemon to
supervise, no inference server to scale alongside FPM. The PHP
process is the inference server.
API ergonomics. Calling a local LLM should be as natural in PHP
as calling intl or pdo . The extension API is shaped to match
that — see Prompts and
Chat completions .
This guide is split into five layers, navigable from the sidebar:
ext-infer is pre-release — the class surface is stable but the
first tagged release ( v0.1.0 ) is still in flight. See
RELEASE.md
for the cut-a-release flow and PLAN.md
for what’s coming next.
Code blocks are runnable as written, with one exception: PHP code
assumes the extension is loaded. Either install it system-wide or
prepend -d extension=… to your php command. See
Installation .
Model without a namespace prefix means Displace\Infer\Model ;
same for Prompt , Response , Embedding . Real code needs the use
statement at the top of the file.
CLI snippets are written for a POSIX shell (bash / zsh). Adjust
for fish / PowerShell as needed; differences are usually only quoting.
