---
source: "https://wayflow.build/"
hn_url: "https://news.ycombinator.com/item?id=48820001"
title: "Show HN: Wayflow – an embeddable AI workflow builder (open source)"
article_title: "Wayflow"
author: "tahazsh"
captured_at: "2026-07-07T17:06:53Z"
capture_tool: "hn-digest"
hn_id: 48820001
score: 2
comments: 0
posted_at: "2026-07-07T16:23:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Wayflow – an embeddable AI workflow builder (open source)

- HN: [48820001](https://news.ycombinator.com/item?id=48820001)
- Source: [wayflow.build](https://wayflow.build/)
- Score: 2
- Comments: 0
- Posted: 2026-07-07T16:23:22Z

## Translation

タイトル: Show HN: Wayflow – 埋め込み可能な AI ワークフロー ビルダー (オープン ソース)
記事タイトル: ウェイフロー
説明: Web 用の埋め込み可能なビジュアル ワークフロー エディター。必要なときは AI を活用し、必要なときはシンプルなロジックを活用
HN テキスト: こんにちは!タハです。ワークフローをサポートする多くのエージェント製品 (私が取り組んだ製品も含む) では、ノードベースのエディターをサポートしていないか、React Flow を使用し、それを製品に統合して実行し、既存のロジックと連携するという困難な作業を行っていることに気づきました。そこで私は、エディターとランタイムの間のギャップを埋めることでこれを支援できるツールを作成することを考えました。それが私が Wayflow を作成した理由です。基本的なアーキテクチャはシンプルです。必要なのは、ランタイムが実行方法を認識するグラフ (JSON オブジェクト) を作成することだけです。ランタイムはグラフがどこから来たのかを気にせず、適切なスキーマだけを必要とします。エディターを使用すると、グラフを作成してエクスポートしたり、バックエンドのデータベースに直接保存したりできます。そして、それを実行したいときは、それをランタイムに渡すだけです。ランタイムは、実行をストリーミングすることも (これはエディターにとって便利です)、最終結果を提供することもできます。グラフをどのように実行するかは、API エンドポイント、cron ジョブ、またはその他の任意の方法を使用するなど、ユーザー次第です。 Wayflow には、入力、出力、LLM、画像生成、配列操作、条件付きなどの多数の組み込みノードが付属しており、アプリが必要とする一般的なケースのほとんどをそれらがカバーしていることを確認するために多くの時間を費やしました。ただし、特定のアプリで機能するカスタム ノードを簡単に追加できます。主な課題の 1 つ (そして楽しい部分の 1 つ) は、プロバイダー中立にすることでした。したがって、特定の LLM プロバイダーにロックされません。私は、ほぼすべてのフロンティア モデルである OpenAI 互換 API をターゲットにしてこれを行いました。

のサポート。フレームワークとライブラリについても同じです。依存関係はなく、フレームワークに依存しません。ツールの呼び出し、完全なワークフローのツールへの変換、人間参加型など、ここで取り上げていないことがたくさんあります。これについてはすべてドキュメントで読むことができます。言及したい重要な点の 1 つは、LLM は作成するワークフローの大部分を占めていますが、完全にオプションであるということです。文字通り、AI を使用せずに (より多くのカスタム ノードを利用して) あらゆるワークフローを構築し、完全な決定論的な実行を実現できます。 https://wayflow.build でブラウザで試すことができます (サインアップなし、独自のキーを持参してください)。これは npm (「npm install wayflow」) にもあり、コードは GitHub ( https://github.com/TahaSh/wayflow ) にあります。これがあなたのプロジェクトでの使用に役立つことを願っています。コーディングを楽しんでください!

記事本文:
Wayflow コンテンツへスキップ Wayflow Docs ドキュメントを検索… ⌘K Wayflow Docs GitHub Theme v0.3.0 · MIT ライセンス取得済み ビジュアル ワークフロー エディター
リビルドではなく埋め込みます。
完全なワークフロー エディター (キャンバス、ノード、および
実際にそれらを実行するランタイム。必要なときに AI を利用できます。
そうでないときのロジック。
GitHub でスターを付ける
インストール
使用法
$ npm install wayflow import { createWorkflowEditor } を「wayflow」からコピーします
const editor = createWorkflowEditor (document.getElementById ( 'editor' ))
タップして対話する
キーを使用する 全画面表示
編集者は一人。任意のフレームワーク、または何もフレームワークなし。
それは単純な TypeScript と DOM です。 1 回の呼び出しでどこにでもマウントできます。
すべてがすでに箱に入っています。
createWorkflowEditor() は、キャンバス、ノード パレット、構成パネル、および実行コントロールをマウントします。これは、自分で完成させた空のキャンバスではなく、完全なワークスペースです。
実際の実行エンジンが同梱されています。バックエンドを使用せずにブラウザーでプロトタイプを作成し、本番環境のサーバーで同じグラフを実行します。
必要な場合は AI、不要な場合はロジック
組み込みの LLM、ツール呼び出し、分岐、マップオーバーリスト ノードは、独自の決定論的なステップと自由に組み合わせることができます。業務に必要なだけ AI を利用することも、まったく利用しないこともできます。
人間のために一時停止し、後で再開する
承認または決定のために実行を一時停止し、すぐに再開します。人間参加型機能は、ボルトで取り付けられるものではなく、組み込まれています。
プロバイダーに依存しない LLM およびイメージ アダプター。キーと選択したベンダーを接続すれば、1 つにロックされることはありません。
アクセントを設定すると、すべてのサーフェス、ホバー、フォーカス リングが再計算されます。明暗は常に同期しているため、製品内でネイティブに見えます。
簡単なヘルパーから完全なエージェントまで。
チームが到達する形状のいくつか — キャンバスは一般的なものなので、構築します
製品に必要なものは何でも。
人間によるチケットのトリアージ、ルーティング、返信の下書き

重要なタイミングを確認してください。
接続されたいくつかのステップにわたってコピーを生成、変換、再利用します。
モデルに計画を立てさせ、ツールを呼び出し、独自に動作させます。
リストをマッピングし、各項目を強化または分類して、結果をマージします。
ウェイフロー
Web 用の埋め込み可能なビジュアル ワークフロー エディター。

## Original Extract

An embeddable visual workflow editor for the web. Powered by AI when you want it, plain logic when you don

Hi! I’m Taha. In many agentic products that support workflows (including one I worked on), I noticed they either don’t support node-based editors, or use React Flow and go through the difficult work of integrating it into their product to run it and work with their existing logic. So I thought about creating a tool that could help with this by closing the gap between the editor and the runtime. That’s why I created Wayflow. The basic architecture is simple: you just need to create a graph (which is a JSON object) that the runtime knows how to run. The runtime doesn’t care where that graph is coming from, it just needs the right schema. And with the help of the editor, you can create the graph, and then export it or directly save it on your backend in your database. And then when you want to execute it, you just hand it to the runtime. The runtime can either stream the execution (which is useful for the editor), or give you the final result. How you execute the graph is up to you: through an API endpoint, a cron job, or any other way you want. Wayflow comes with a number of built-in nodes like Input, Output, LLM, Image Generation, Array Operations, Conditional, ... and I spent a lot of time making sure they cover most of the common cases your app needs. But you can easily add your custom nodes that work for your specific app. One of the main challenges (and one of the fun parts) was making it provider-neutral. So it’s not locked to a specific LLM provider. I did this by targeting the OpenAI-compatible API, which almost all frontier models support. The same thing for frameworks and libraries. It has zero dependencies and it’s framework-agnostic. There are many things I didn’t cover here, like tool calling, converting full workflows into tools, human-in-the-loop, etc. You can read all about this in the docs. One important thing I’d like to mention is that even though LLMs are a huge part of the workflows you create, they are completely optional. You can literally build any workflow (with the help of more custom nodes) with zero AI, for full deterministic execution. You can try it in your browser (no signup, bring your own key) at https://wayflow.build . It’s also on npm (`npm install wayflow`), and the code is on GitHub ( https://github.com/TahaSh/wayflow ). I hope this is something you’ll find helpful to use in your projects. Happy coding!

Wayflow Skip to content Wayflow Docs Search docs… ⌘K Wayflow Docs GitHub Theme v0.3.0 · MIT licensed The visual workflow editor
you embed , not rebuild.
Drop a complete workflow editor into your product — canvas, nodes, and a
runtime that actually runs them. Powered by AI when you want it, plain
logic when you don’t.
Star on GitHub
Install
Usage
Copy $ npm install wayflow import { createWorkflowEditor } from 'wayflow'
const editor = createWorkflowEditor (document. getElementById ( 'editor' ))
Tap to interact
Use your key Fullscreen
One editor. Any framework — or none.
It’s plain TypeScript and the DOM. One call mounts it anywhere.
Everything’s already in the box.
createWorkflowEditor() mounts the canvas, node palette, config panel, and run controls — a complete workspace, not a blank canvas you finish yourself.
A real execution engine ships in the box. Prototype in the browser with no backend, then run the same graph on your server in production.
AI when you want it, logic when you don’t
Built-in LLM, tool-calling, branching, and map-over-list nodes mix freely with your own deterministic steps. Reach for as much AI as the job needs — or none.
Pause for a human, resume later
Suspend a run for an approval or a decision, then pick it right back up. Human-in-the-loop is built in, not bolted on.
Provider-neutral LLM and image adapters. Plug in your keys and your vendor of choice — you’re never locked to one.
Set the accent and every surface, hover, and focus ring recomputes. Light and dark stay in sync, so it looks native in your product.
From a quick helper to a full agent.
A few of the shapes teams reach for — the canvas is general, so build
whatever your product needs.
Triage tickets, route them, and draft replies — with a human check when it counts.
Generate, translate, and repurpose copy across a few connected steps.
Let a model plan, call your tools, and act on its own.
Map over a list, enrich or classify each item, and merge the results.
Wayflow
An embeddable visual workflow editor for the web.
