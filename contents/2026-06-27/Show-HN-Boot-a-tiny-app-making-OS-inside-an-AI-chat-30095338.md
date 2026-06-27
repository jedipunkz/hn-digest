---
source: "https://hollingsworthd.github.io/RAu/"
hn_url: "https://news.ycombinator.com/item?id=48693789"
title: "Show HN: Boot a tiny app-making OS inside an AI chat"
article_title: "RFC BBOS-RAU-0001 — BlueBookOS: Thee GPT Microkernel™"
author: "logn"
captured_at: "2026-06-27T00:30:18Z"
capture_tool: "hn-digest"
hn_id: 48693789
score: 1
comments: 0
posted_at: "2026-06-27T00:27:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Boot a tiny app-making OS inside an AI chat

- HN: [48693789](https://news.ycombinator.com/item?id=48693789)
- Source: [hollingsworthd.github.io](https://hollingsworthd.github.io/RAu/)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T00:27:01Z

## Translation

タイトル: Show HN: AI チャット内で小さなアプリ作成 OS を起動する
記事のタイトル: RFC BBOS-RAU-0001 — BlueBookOS: GPT Microkernel™
説明: 読みやすいインストール ページ、ソースファースト RAu ブート パケット、および BlueBookOS 用の 1 つのスタンドアロン HTML5 デモ。

記事本文:
これをコピーしてください。貼り付けてください。 Enterを押します。
これにより、AI チャット内で小さなアプリ作成 OS が起動します。最初に RAu や BlueBookOS について知る必要はありません。ボックス全体を新しいチャットに貼り付け、Enter キーを押すと、何を構築するかを尋ねられます。
OSブートパケットをロード中…
AIを使ったアプリを作るための小さなOS
BlueBookOS
チャットでブーツを使用します。
BlueBookOS は、別の AI チャットに貼り付ける小さなルールとソース コードのセットです。チャットでアプリを慎重に構築するのに役立ちます。最初に RAu コントラクト、次に作業ファイルです。 RAu は、その中にある小さな命令言語です。どちらも知らなくても始められます。
このドキュメントでは、BlueBookOS の実用的なアーティファクト パッケージ パターン、つまり RAu を搭載した GPT Microkernel™ を指定します。
ここでは、RAu がソースファーストのアーティファクト コントラクトとして使用されます。 HTML5 ファイルは、RAu ソースによって記述された動作をレンダリングするホスト ポートです。このアプリは意図的に操作が簡単です。アプリを開いて、デモを開き、最終的な RAu を検査し、最終的な HTML を検査し、レンダリングされたページを適切な場所で実行します。
製品の主な主張はシンプルです。優れた AI 成果物は、見た目が印象的であるだけではありません。再利用可能、レビュー可能、コピー可能で、次のプロジェクトに移植できる必要があります。
これは、要点を儀式に埋没させることなく、RFC スタイルで書かれた、読みやすい製品向けのドラフトです。
キーワード MUST 、 SHOULD 、および MAY は、日常の標準文書の意味で使用されます。 MUST は、このショーケース パターンに必須であることを意味し、 SHOULD はほとんどのプロジェクトに推奨されることを意味し、 MAY はオプションだがサポートされることを意味します。
このファイルは自己完結型です。サーバー、パッケージ マネージャー、ビルド パイプライン、リモート スタイルシート、CDN、またはネットワーク呼び出しは必要ありません。
これらの用語は、ソースファースト モデルを維持しながら、仕様を読みやすく保ちます。
アーティファクトのセマンティック ソース コントラクトとして使用されるプログラミング言語とランタイム。

出荷またはリミックスを目的とした完全なアプリ、ゲーム、エディター、リーダー、ビジュアライザー、またはワークフロー ユニット。
実行可能なターゲットの実装。このパッケージでは、すべてのホスト ポートがスタンドアロン HTML5 です。
最終的な HTML ポートを埋め込みフレームにロードすることによって生成されるライブ ブラウザ ビュー。
RAu コントラクトは、ホスト言語の変更が行われる前の永続的な製品定義として扱われます。
人間が完全な成果物、明確な目標、および承認基準を備えたフロンティア モデルを操作する高速ビルダー ワークフロー。
4. アーティファクトの梱包仕様
ショーケースでは、最終的なサンプル、最終的な RAu、最終的な HTML、ライブ レンダリングされた出力など、ビルダーが必要とする部分だけを公開します。
RAu は引き続き耐久製品契約です。 HTML5 は、実行可能なブラウザ配信形式です。
リファレンス フローは意図的に軽量になっています。ゲーム、エディタ、リーダー、キャンバス ツール、ダッシュボード、シミュレーション、その他の UI を重視するアーティファクトで機能します。
アーティファクト、ユーザー コントロール、データ モデル、レンダー ターゲット、成功条件に名前を付けます。
動作、ガード、イベント、レンダリングの製品コントラクトとして RAu を使用します。
完全な入力およびレンダリング動作を備えたスタンドアロンのブラウザ ポートを生成します。
ページを実行し、ソースをコピーし、コントラクトを変更して、パターンを再利用します。
アーティファクトを選択して、その最終的な RAu、最終的な HTML、および最終的にレンダリングされたページを検査します。
このページの上部にある大きなプロンプトを使用して、新しいソースファースト RAu アーティファクトをリクエストします。
推奨されるワークフロー: 上部の大きなプロンプトをコピーし、括弧で囲まれた目標を必要なアプリに置き換え、スタンドアロン HTML5 ファイルの前に RAu コントラクトを返すようにモデルに依頼します。ソースファースト パターンがどのように見えるかに関する参考として、以下の例を使用してください。
スクリーンショットに頼らないでください。便利なパターンは、ソースファースト契約の後に、動作するスタンドアロン ホスト ポートが続くことです。
例: 「作る」

カレンダー アプリ」、「パズル ゲームの派生」、「これを ToDo ツールに移植」、または「5 番目の例を追加」。
ソースファーストのルールを維持します。ホスト ポートを変更する前に、RAu コントラクトを更新または派生します。
明示的に意図されていない限り、ビルド ステップ、CDN、およびランタイム ネットワーク依存性のない単一の HTML5 ファイルが必要です。
モデルに、コントロール、予想される動作、エッジケース、および元の例からの変更点をリストするように依頼します。
ラピッドプロトタイプ、クライアントデモ、内部ツール、インタラクティブ仕様、ゲームメカニクス、UI実験、製品の売り込み。
完全な HTML ファイルと明確な指示: RAu/HTML パターンを抽出し、次のアーティファクトを生成します。
改訂された RAu 契約と完全なスタンドアロン HTML5 アプリ。ローカル ファイルとして保存して実行する準備ができています。
コピーした HTML アプリでこれらのプロンプトを使用して、フロンティア モデルを有用で出荷可能な出力に導きます。
プロンプト B · デモをリファクタリングまたは拡張する
プロンプトのコピー
9. 適合性チェックリスト
新しい BlueBookOS / RAu ショーケース エントリは、これらのチェックを満たした場合にこの仕様に準拠します。
このアーティファクトは、明らかに BlueBookOS: Thee GPT Microkernel™ および Powered by RAu としてブランド化されています。
最終的な RAu ソースは、ビルド手順なしで検査およびコピーできます。
最終的なスタンドアロン ホスト ポートは、検査、コピー、ダウンロード、および開くことができます。
プレビューはローカルの埋め込みソースを使用して実行され、リモート アセットには依存しません。
パッケージには、人間またはモデルがパターンから別のアーティファクトを導き出すのに十分なコンテキストが含まれています。

## Original Extract

A readable install page, source-first RAu boot packet, and one standalone HTML5 demo for BlueBookOS.

Copy this. Paste it. Press enter.
This boots a tiny app-making OS inside an AI chat. You do not need to know RAu or BlueBookOS first. Paste the whole box into a fresh chat, press enter, and it will ask what you want to build.
Loading OS boot packet…
A tiny OS for making apps with AI
BlueBookOS
boots in your chat.
BlueBookOS is a small set of rules and source code you paste into another AI chat. It helps that chat build apps carefully: first the RAu contract, then the working file. RAu is the little instruction language inside it. You can start without knowing either one.
This document specifies a practical artifact packaging pattern for BlueBookOS: Thee GPT Microkernel™, Powered by RAu.
RAu is used here as a source-first artifact contract. The HTML5 files are host ports that render the behavior described by the RAu source. This app is intentionally plain to operate: open it, open the demo, inspect the final RAu, inspect the final HTML, and run the rendered page in place.
The main product claim is simple: a good AI artifact should not only be impressive to view; it should be reusable, reviewable, copyable, and portable into the next project.
This is a readable product-facing draft, written in an RFC style without burying the point in ceremony.
The key words MUST , SHOULD , and MAY are used in their everyday standards-document sense: MUST means required for this showcase pattern, SHOULD means recommended for most projects, and MAY means optional but supported.
This file is self-contained. It does not require a server, package manager, build pipeline, remote stylesheet, CDN, or network call.
These terms keep the spec readable while preserving the source-first model.
The programming language and runtime used as the semantic source contract for an artifact.
A complete app, game, editor, reader, visualizer, or workflow unit intended to be shipped or remixed.
The executable target implementation. In this package, every host port is standalone HTML5.
The live browser view generated by loading the final HTML port into an embedded frame.
The RAu contract is treated as the durable product definition before host-language changes are made.
A fast builder workflow where a human steers a frontier model with complete artifacts, clear goals, and acceptance criteria.
4. Artifact Packaging Specification
The showcase exposes just the pieces a builder needs: final examples, final RAu, final HTML, and live rendered output.
RAu remains the durable product contract; HTML5 is the executable browser delivery format.
The reference flow is intentionally lightweight. It works for games, editors, readers, canvas tools, dashboards, simulations, and other UI-heavy artifacts.
Name the artifact, user controls, data model, render targets, and success conditions.
Use RAu as the product contract for behavior, guards, events, and rendering.
Generate a standalone browser port with complete input and render behavior.
Run the page, copy the source, modify the contract, and reuse the pattern.
Select an artifact to inspect its final RAu, final HTML, and final rendered page.
Use the large prompt at the top of this page to request a new source-first RAu artifact.
Recommended workflow: copy the large prompt at the top, replace the bracketed goal with the app you want, and ask the model to return the RAu contract before the standalone HTML5 file. Use the examples below as references for how the source-first pattern should look.
Do not rely on screenshots. The useful pattern is the source-first contract followed by a working standalone host port.
Examples: “make a calendar app,” “derive a puzzle game,” “port this into a todo tool,” or “add a fifth example.”
Preserve the source-first rule: update or derive the RAu contract before changing the host port.
Require a single HTML5 file with no build step, no CDN, and no runtime network dependency unless explicitly intended.
Ask the model to list controls, expected behaviors, edge cases, and what changed from the original examples.
Rapid prototypes, client demos, internal tools, interactive specs, game mechanics, UI experiments, and product pitches.
The full HTML file plus a clear instruction: extract the RAu/HTML pattern and generate the next artifact.
A revised RAu contract and a complete standalone HTML5 app, ready to save as a local file and run.
Use these prompts with the copied HTML app to guide a frontier model toward useful, shippable output.
Prompt B · Refactor or extend the demo
Copy prompt
9. Conformance Checklist
A new BlueBookOS / RAu showcase entry conforms to this spec when it satisfies these checks.
The artifact is clearly branded as BlueBookOS: Thee GPT Microkernel™ and Powered by RAu.
The final RAu source can be inspected and copied without a build step.
The final standalone host port can be inspected, copied, downloaded, and opened.
The preview runs with local embedded source and does not depend on remote assets.
The package includes enough context for a human or model to derive another artifact from the pattern.
