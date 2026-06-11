---
source: "https://productnow.ai/blogs/extracting-html-from-ai-prototyping-tools"
hn_url: "https://news.ycombinator.com/item?id=48492741"
title: "Which AI prototyping tools can export a self-contained HTML file?"
article_title: "How we extracted standalone HTML files from 7 / 13 AI prototyping tools | ProductNow"
author: "kadhirvelm"
captured_at: "2026-06-11T19:07:13Z"
capture_tool: "hn-digest"
hn_id: 48492741
score: 1
comments: 1
posted_at: "2026-06-11T16:42:37Z"
tags:
  - hacker-news
  - translated
---

# Which AI prototyping tools can export a self-contained HTML file?

- HN: [48492741](https://news.ycombinator.com/item?id=48492741)
- Source: [productnow.ai](https://productnow.ai/blogs/extracting-html-from-ai-prototyping-tools)
- Score: 1
- Comments: 1
- Posted: 2026-06-11T16:42:37Z

## Translation

タイトル: 自己完結型 HTML ファイルをエクスポートできる AI プロトタイピング ツールはどれですか?
記事のタイトル: 7/13 AI プロトタイピング ツールからスタンドアロン HTML ファイルを抽出した方法 |今すぐプロダクト
説明: 13 個の AI プロトタイピング ツールから自己完結型 HTML ファイルを抽出してみました。 7 つは機能しましたが、6 つは機能しませんでした。 ChatGPT のクリーンな 3.2 KB エクスポートから、まだオフラインで実行できなかった Canva の 3.5 MB バンドルまで、各プラットフォームについて私たちが学んだことは次のとおりです。

記事本文:
7/13 AI プロトタイピング ツールからスタンドアロン HTML ファイルを抽出した方法 | ProductNow PRODUCTNOW 当社のパートナー The Loop コンプライアンス 採用情報 ブログ ログイン 無料で始めましょう すべての投稿 7 / 13 AI プロトタイピング ツールからスタンドアロン HTML ファイルを抽出した方法 by Kadhir Mani (12.3 分) <section data-section-id='a8edf699-5e9f-48e9-a173-a42e03e36b30'><collapsible-card data-icon="file-text" data-title="TL;DR results"><p><span style="white-space: pre-wrap;">13 個のプロトタイピング プラットフォームをテストした結果、次のような全体像とファイル サイズの視覚化 (y 軸は kb) が表示されます。</span></p><p><bar-chart data-height="450"><bar-line value="3.2" color="#22C55E">ChatGPT</bar-line><bar-line value="205" color="#22C55E">クロード</bar-line><bar-line value="270" color="#22C55E">ジェミニ キャンバス</bar-line><bar-line value="288" color="#22C55E">ボルト</bar-line><bar-line value="293" color="#22C55E">愛らしい</bar-line><bar-line value="497" color="#22C55E">リプリット</bar-line><bar-line value="1334" color="#F59E0B">フィグマメイク</bar-line><bar-line value="74.3" color="#EF4444">V0</bar-line><bar-line value="487" color="#EF4444">ホスティンガー</bar-line><bar-line value="960" color="#EF4444">任意</bar-line><bar-line value="3500" color="#EF4444">Canva コード</bar-line><bar-line value="0" color="#EF4444">UX パイロット</bar-line><bar-line value="0" color="#EF4444">Base44</bar-line></bar-chart></p><p><b><strong style="white-space: pre-wrap;">緑</strong></b><span style="white-space: pre-wrap;"> = 自己完結型 HTML が達成、</span><b><strong style="white-space: pre-wrap;">琥珀</strong></b><span style="white-space: pre-wrap;"> = HTTP サーバーが必要、</span><b><strong style="white-space: pre-wrap;">赤</strong></b><span style="white-space: pre-wrap;"> = 達成不可</span></p></collapsible-card><p><br></p><

/セクション>
<section data-section-id='12e4a8f5-9680-4d8f-99a3-78b95d5d8fd4'><h2 id='12e4a8f5-9680-4d8f-99a3-78b95d5d8fd4'>背景</h2><p><span style="white-space: pre-wrap;">プロトタイプは私たちのお気に入りの方法の 1 つです
[切り捨てられた]
同僚が AI エージェントの場合、CRDT だけでは十分ではありません
CRDT は収束します。彼らは協調していない。この区別は、AI エージェントが編集者に加わった瞬間に重要になります。マルチプレイヤー コラボレーションにマージの上に調整レイヤーが必要な理由と、エージェントのマナーを保つために構築されたセクション レベルのリース プロトコル。
良いコード、間違った機能: ハンドオフの問題
誰かがあなたにチケットを送ります。「エクスポート機能を追加してください。何が必要かわかりますね。」実際にはそうではありません。 AI ネイティブ開発の本当の問題はコードではなく、発見から実装への引き継ぎにあるのはなぜでしょうか。
効果的に連携したチームが勝利します。
AI が実行能力を拡大するにつれて、次のボトルネックは調整になります。
hello@productnow.ai までお問い合わせください
© 2026 ProductNow.無断転載を禁じます。
Cookie の管理 hello@productnow.ai までお問い合わせください
© 2026 ProductNow.無断転載を禁じます。

## Original Extract

We tried extracting self-contained HTML files from 13 AI prototyping tools. 7 worked, 6 didn't. Here's what we learned about each platform — from ChatGPT's clean 3.2 KB export to Canva's 3.5 MB bundle that still couldn't run offline.

How we extracted standalone HTML files from 7 / 13 AI prototyping tools | ProductNow PRODUCTNOW Our partners The loop Compliance Careers Blog Log in Get started free All posts How we extracted standalone HTML files from 7 / 13 AI prototyping tools by Kadhir Mani (12.3 minutes) <section data-section-id='a8edf699-5e9f-48e9-a173-a42e03e36b30'><collapsible-card data-icon="file-text" data-title="TL;DR results"><p><span style="white-space: pre-wrap;">After testing thirteen prototyping platforms, here's the full picture we saw + a visualization on file size (y-axis in kb):</span></p><p><bar-chart data-height="450"><bar-line value="3.2" color="#22C55E">ChatGPT</bar-line><bar-line value="205" color="#22C55E">Claude</bar-line><bar-line value="270" color="#22C55E">Gemini Canvas</bar-line><bar-line value="288" color="#22C55E">Bolt</bar-line><bar-line value="293" color="#22C55E">Lovable</bar-line><bar-line value="497" color="#22C55E">Replit</bar-line><bar-line value="1334" color="#F59E0B">Figma Make</bar-line><bar-line value="74.3" color="#EF4444">V0</bar-line><bar-line value="487" color="#EF4444">Hostinger</bar-line><bar-line value="960" color="#EF4444">Anything</bar-line><bar-line value="3500" color="#EF4444">Canva Code</bar-line><bar-line value="0" color="#EF4444">UX Pilot</bar-line><bar-line value="0" color="#EF4444">Base44</bar-line></bar-chart></p><p><b><strong style="white-space: pre-wrap;">Green</strong></b><span style="white-space: pre-wrap;"> = self-contained HTML achieved, </span><b><strong style="white-space: pre-wrap;">Amber</strong></b><span style="white-space: pre-wrap;"> = requires HTTP server, </span><b><strong style="white-space: pre-wrap;">Red</strong></b><span style="white-space: pre-wrap;"> = not achievable</span></p></collapsible-card><p><br></p></section>
<section data-section-id='12e4a8f5-9680-4d8f-99a3-78b95d5d8fd4'><h2 id='12e4a8f5-9680-4d8f-99a3-78b95d5d8fd4'>Background</h2><p><span style="white-space: pre-wrap;">Prototypes are one of our favorite ways
[truncated]
CRDTs Are Not Enough When Your Coworker Is an AI Agent
CRDTs converge; they do not coordinate. That distinction becomes critical the moment an AI agent joins the editor. Why multiplayer collaboration needs a coordination layer above merge, and the section-level lease protocol we built to keep agents well-mannered.
Good code, Wrong feature: The handoff problem
Someone sends you a ticket: 'Add the export feature — you know what we need.' You don't, actually. Why the real breakdown in AI-native development isn't the code — it's the handoff from discovery to implementation.
Teams that coordinate effectively will win.
As AI expands execution capacity, the next bottleneck will be coordination..
Contact us at hello@productnow.ai
© 2026 ProductNow. All rights reserved.
Manage cookies Contact us at hello@productnow.ai
© 2026 ProductNow. All rights reserved.
