---
source: "https://pewdiepie-archdaemon.github.io/odysseus/"
hn_url: "https://news.ycombinator.com/item?id=48347265"
title: "PewDiePie's AI Workspace"
article_title: "Odysseus — A Self-Hosted AI Workspace"
author: "r0xsh"
captured_at: "2026-06-01T00:27:45Z"
capture_tool: "hn-digest"
hn_id: 48347265
score: 6
comments: 0
posted_at: "2026-05-31T16:52:04Z"
tags:
  - hacker-news
  - translated
---

# PewDiePie's AI Workspace

- HN: [48347265](https://news.ycombinator.com/item?id=48347265)
- Source: [pewdiepie-archdaemon.github.io](https://pewdiepie-archdaemon.github.io/odysseus/)
- Score: 6
- Comments: 0
- Posted: 2026-05-31T16:52:04Z

## Translation

タイトル: PewDiePie の AI ワークスペース
記事のタイトル: Odysseus — セルフホスト型 AI ワークスペース
説明: Odysseus — 自己ホスト型 AI ワークスペース: チャット、エージェント、ツール、モデル提供、電子メール、リサーチなど。モデル、ハードウェア、データ。

記事本文:
オデュッセウス
特長
お客様の声
始まりの経緯
始めましょう
GitHub
オデュッセウス
航海をよろしくお願いします。
あなた自身の AI ワークスペース、
ハードウェア上で実行されます。
Odysseus は、言語モデルと対話するための自己ホスト型インターフェイスです (チャット、
自律エージェント、ツール、モデル提供、電子メール、調査など。地元第一主義、
プライバシー最優先で、テレメトリはありません。あなたとあなたのモデルだけです。
(素晴らしい API も追加したい場合は、私は人生の生き方を教えるためにここにいるわけではありません…)
1 つのアプリで多くの機能を実現
AIチャットとしてスタート。ワークスペースになりました。各ピースはローカルで実行されます
どのエンドポイントを指していても。
マルチターン チャットと、計画を立て、ツールを呼び出し、タスクを実行する自律エージェント。
組み込みツール (bash、ファイル、Web、メモリ) と接続する MCP サーバー。ツールごとに切り替えます。
ハードウェアを認識したモデルの推奨事項と、270 を超えるカタログ化されたモデルをワンクリックで提供します。
AI 概要、スタイルが一致した返信の下書き、自動タグ付け、および IMAP/SMTP 経由のスパムトリアージ。
情報源を収集、読み取り、統合してレポートを作成する、複数段階の調査を実行します。
1 つのプロンプトを複数のモデルに一度に送信し、その回答を並べて比較します。
アシスタントが構築し、すべての会話を通じて思い出す永続的な記憶。
アシスタントは自身のスキルを作成、洗練、再利用し、時間の経過とともにより能力を高めます。
独自のエンドポイントに対してマシン上で実行されます。テレメトリはありません。オプションの外部統合を選択すると利用できます。
企業に愛される
お客様の声
「オデュッセウスは、船を輸送しながら、より多くの船を出荷するのに役立ちました。まさにクラス最高の輸送です。」
「私は実在の人物です。これは本物の証言です。本物の女性によるものです。」
「ああああああああああああああ」
「とにかく、私が言ったように、クラス最高です。」
オデュッセウスは、慎重に作成されたワンショット AI プロンプトによって作成されました。
> わかりません

何かを思いつく、AI チャットを作る、でもそれを良いものにして、見栄えをよくする
✕ ターミナルを再開します
実際の動作を確認してください
カーソルを合わせるかタップして詳しく見てください
各パネルにマウスを置くかタップすると、各パネルが展開され、プレビューが再生されます。モバイルではスワイプしてそれらの間を移動します。
[ チャットとエージェント ]
チャットとエージェント ローカル モデルと会話するか、ツールを与えてエージェントを実行させます。
[ 料理本 ]
クックブック マシン全体でローカル モデルをダウンロード、提供、管理します。
[ 深い研究 ]
Deep Research Ask Once: 検索し、ソースを読み、引用されたレポートを書き戻します。
[ 比較 ]
比較 1 つのプロンプトを一度に多くのモデルに送信し、それらの応答を並べて観察します。
【資料】
ドキュメント ユーザーを第一に考えたドキュメント エディタです。必要なときに AI の助けを借りて、必要な作業に取り組むことができます。
[ メモとタスク ]
メモとタスク メモや To-Do を記録したり、スケジュールされたエージェントに作業を任せて後で説明したりできます。
【画像ギャラリー】
イメージ ギャラリー 独自のギャラリーで背景を生成、編集、削除し、修復します。
【テーマ】
テーマのスタイルを変更して自分のものにします。自分で編集するか、エージェントに作成を依頼します。
実際に始まった経緯
妥協のないローカル LLM エクスペリエンス。
私が Odysseus プロジェクトに取り組み始めたのは、ローカル AI を実行するのが楽しくて強力だと感じたからです。
しかし、LLM と関わるという当時の選択肢は、一歩後退するように感じられました。あなたがそう思うという考えは、
AI を自己ホストするだけでサブスクリプション料金を支払わないということは存在しませんでした。すべてのツールと機能
それはすべての魔法が欠けていたことを意味します。
そこで、私はオデュッセウスを少しずつ構築し始めました。そして、取り組めば取り組むほど、
それは私にとって良かったです。モデルがあなたについて知れば知るほど、モデルはより便利になることがわかりました。
これが自己ホストするもう 1 つの理由です。プライベートを渡さずにコンテキストをすべて取得できるからです。
データを他人のクラウドに転送します。
始めましょう
オデュッセウスはあなたのものです。
オープンソースです

そして無料。営業チームもデモリクエストもトロイの木馬もありません。

## Original Extract

Odysseus — a self-hosted AI workspace: chat, agents, tools, model serving, email, research, and more. Your models, your hardware, your data.

Odysseus
Features
Testimonials
How it started
Get started
GitHub
Odysseus
Yours for the voyage.
Your own AI workspace ,
running on your hardware.
Odysseus is a self-hosted interface for talking to language models — chat,
autonomous agents, tools, model serving, email, research, and more. Local-first,
privacy-first, and no telemetry. Just you and your models.
(if you want to add an API that's cool too — I'm not here to tell you how to live your life…)
One app, a lot of capabilities
Started as an AI chat. Became a workspace. Each piece runs locally against
whatever endpoints you point it at.
Multi-turn chat plus autonomous agents that plan, call tools, and work through tasks.
Built-in tools (bash, files, web, memory) plus any MCP server you connect. Toggle per tool.
Hardware-aware model recommendations and one-click serving across 270+ catalogued models.
AI summaries, style-matched draft replies, auto-tagging and spam triage over IMAP/SMTP.
Multi-step research runs that gather, read, and synthesize sources into a written report.
Send one prompt to several models at once and compare their answers side-by-side.
Persistent memory the assistant builds up and recalls across all your conversations.
The assistant writes, refines, and reuses its own skills — getting more capable over time.
Runs on your machine against your own endpoints. No telemetry, with optional external integrations when you choose them.
Loved by enterprises
What our customers are saying
"Odysseus helped us ship more ships while shipping ships. Truly best-in-class shipping."
"I'm a real person. This is a real testimonial. By a real woman."
"AHHHHHHHHHHHHHHHHHHHHHHHHHHHHH"
"Anyway, as I was saying — best-in-class."
Odysseus was created by a carefully crafted one-shot AI prompt:
> idk what to make come up with something oh make an AI chat but make it good and make it look nice
✕ reopen terminal
See it in action
Hover or tap to take a closer look
Each panel expands and plays its preview when you hover or tap it. Swipe on mobile to move through them.
[ Chat & Agents ]
Chat & Agents Talk to any local model, or give it tools and let the agent run.
[ Cookbook ]
Cookbook Download, serve, and manage local models across your machines.
[ Deep Research ]
Deep Research Ask once: it searches, reads sources, and writes back a cited report.
[ Compare ]
Compare Send one prompt to many models at once and watch them answer side by side.
[ Documents ]
Documents A document editor that puts you first — work on what you want, with AI help when you want it.
[ Notes & Tasks ]
Notes & Tasks Capture notes and to-dos, or let scheduled agents work and brief you after.
[ Image Gallery ]
Image Gallery Generate, edit, remove backgrounds, and inpaint in your own gallery.
[ Themes ]
Themes Restyle and make it yours — edit your own, or ask the agent to make one.
How it actually started
Uncompromised local LLM experience.
I started working on the Odysseus project because running local AI felt fun and powerful.
But the options at the time to engage with LLMs felt like taking steps back. The idea that you
could just self-host AI and not pay for a subscription wasn't there. All the tools and functions
that make it all magic were missing.
So I started building Odysseus bit by bit — and the more I gave it to work with, the
better it served me. Turns out the more your model knows about you, the more useful it gets.
Which is the other reason to self-host: you get all that context without handing your private
data to someone else's cloud.
Get started
Odysseus is yours.
It's open source and free. No sales team, no demo request, no Trojan horse.
