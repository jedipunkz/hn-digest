---
source: "https://www.agentation.com/"
hn_url: "https://news.ycombinator.com/item?id=48873337"
title: "Agentation – Visual UI Annotation for AI Coding Agents"
article_title: "Agentation"
author: "rekl"
captured_at: "2026-07-11T16:43:13Z"
capture_tool: "hn-digest"
hn_id: 48873337
score: 2
comments: 0
posted_at: "2026-07-11T16:16:51Z"
tags:
  - hacker-news
  - translated
---

# Agentation – Visual UI Annotation for AI Coding Agents

- HN: [48873337](https://news.ycombinator.com/item?id=48873337)
- Source: [www.agentation.com](https://www.agentation.com/)
- Score: 2
- Comments: 0
- Posted: 2026-07-11T16:16:51Z

## Translation

タイトル: Agentation – AI コーディング エージェントのためのビジュアル UI アノテーション
記事タイトル: エージェント
説明: エージェント用の視覚的なフィードバック ツール。

記事本文:
エージェント エージェントは現在デスクトップのみです。概要 インストール機能 出力スキーマ MCP API Webhook 変更ログ ブログ FAQ 概要 インストール機能 出力スキーマ AFS 1.1 ツール MCP API Webhook リソース 変更ログ ブログ FAQ · 視覚的なフィードバック。
エージェント向け。
エージェント化は、UI アノテーションを、AI コーディング エージェントが理解して動作できる構造化コンテキストに変換します。任意の要素をクリックしてメモを追加し、出力をクロード コード、コーデックス、または任意の AI ツールに貼り付けます。
1 2 3 4 5 ボタン.送信ボタン |キャンセル 追加 Benji のプロジェクト クロード コード v2.1.14 おかえり、Benji! Opus 4.5 · ~/Code/agentation █ 使用方法
右下隅のアイコンをクリックして有効にします
要素の上にマウスを置くと、その名前が強調表示されます。
任意の要素をクリックして注釈を追加します
フィードバックを書いて「追加」をクリックします
クリックして書式設定されたマークダウンをコピーします
注: MCP を使用すると、コピーと貼り付けの手順を完全にスキップできます。エージェントは、あなたが指している内容をすでに認識しています。 「フィードバックに対処する」または「注釈 3 を修正する」と言うだけです。
エージェントは、コードベース (クロード コード、カーソルなど) にアクセスできる AI ツールで最も効果的に機能します。出力を貼り付けると、エージェントは次の情報を取得します。
コードベースを grep するための CSS セレクター
右側の行に直接ジャンプするソースファイルのパス
階層を理解するための React コンポーネント ツリー
現在の外観を理解するために計算されたスタイル
意図と優先度を備えたフィードバック
エージェントがなければ、要素 (「サイドバーの青いボタン」) を説明し、エージェントが正しく推測することを期待する必要があります。 Agentation を使用すると、.sidebar > button.primary を指定すると、それを直接 grep できます。
このページではツールバーがアクティブになっています。これらのデモ要素に注釈を付けてみてください。
注釈を作成するには、このカードをクリックするか、このテキストを選択します。出力には要素のパスとフィードバックが含まれます。
ツールバーをクリックします

このアニメーションをフリーズするには:
MCP の統合と注釈フォーマット スキーマを使用すると、エージェントは注釈を読むだけでなく、それに応答することができます。
「どんな注釈があるの？」 — ページ全体のすべてのフィードバックをリストします
「これは 24 ピクセルにするべきですか、それとも 16 ピクセルにするべきですか?」 — エージェントが説明を求めます
「パディングを修正しました」 - エージェントが概要を表示して解決します
「すべての注釈を消去」 - すべてを一度に消去します
あなたのフィードバックは会話となり、虚空への片道切符ではなくなります。
具体的にする - 「ボタンのテキストが不明瞭」は「これを修正してください」よりも優れています
注釈ごとに 1 つの問題 - エージェントが個別に対処しやすくなります
コンテキストを含める — 期待したものと実際に見たものを比較して説明します
テキスト選択を使用する — タイプミスやコンテンツの問題がある場合は、正確なテキストを選択します
アニメーションを一時停止する — 特定のアニメーション フレームに注釈を付ける
エージェントは、個人および企業の内部使用であれば無料です。これを使用して、独自のプロジェクトに注釈を付けたり、独自のアプリをデバッグしたり、チーム内のフィードバックを合理化したりできます。 Agentation を販売する製品の一部として再配布する場合は、商用ライセンスについてお問い合わせください。
MCP とのリアルタイム同期を設定する →
アノテーションを外部サービスにプッシュ →
API との独自の統合を構築する →
Benji Taylor、Dennis Jin、Alex Vanderzon によって作成されました

## Original Extract

The visual feedback tool for agents.

Agentation Agentation is currently desktop only. Overview Install Features Output Schema MCP API Webhooks Changelog Blog FAQ Overview Install Features Output Schema AFS 1.1 Tools MCP API Webhooks Resources Changelog Blog FAQ · Visual feedback.
For agents.
Agentation turns UI annotations into structured context that AI coding agents can understand and act on. Click any element, add a note, and paste the output into Claude Code, Codex, or any AI tool.
1 2 3 4 5 button.submit-btn | Cancel Add Benji's Project Claude Code v2.1.14 Welcome back Benji! Opus 4.5 · ~/Code/agentation █ How you use it
Click the icon in the bottom-right corner to activate
Hover over elements to see their names highlighted
Click any element to add an annotation
Write your feedback and click Add
Click to copy formatted markdown
Note: With MCP , you can skip the copy-paste step entirely — your agent already sees what you're pointing at. Just say “address my feedback” or “fix annotation 3.”
Agentation works best with AI tools that have access to your codebase (Claude Code, Cursor, etc.). When you paste the output, agents get:
CSS selectors to grep your codebase
Source file paths to jump directly to the right line
React component tree to understand the hierarchy
Computed styles to understand current appearance
Your feedback with intent and priority
Without Agentation, you’d have to describe the element (“the blue button in the sidebar”) and hope the agent guesses right. With Agentation, you give it .sidebar > button.primary and it can grep for that directly.
The toolbar is active on this page. Try annotating these demo elements:
Click on this card or select this text to create an annotation. The output will include the element path and your feedback.
Click in the toolbar to freeze this animation:
With MCP integration and the Annotation Format Schema , agents don’t just read your annotations — they can respond to them:
“What annotations do I have?” — List all feedback across pages
“Should this be 24px or 16px?” — Agent asks for clarification
“Fixed the padding” — Agent resolves with a summary
“Clear all annotations” — Dismiss everything at once
Your feedback becomes a conversation, not a one-way ticket into the void.
Be specific — “Button text unclear” is better than “fix this”
One issue per annotation — easier for the agent to address individually
Include context — mention what you expected vs. what you see
Use text selection — for typos or content issues, select the exact text
Pause animations — to annotate a specific animation frame
Agentation is free for individuals and companies for internal use. Use it to annotate your own projects, debug your own apps, or streamline feedback within your team. Contact us for a commercial license if you're redistributing Agentation as part of a product you sell.
Set up real-time sync with MCP →
Push annotations to external services →
Build your own integration with the API →
Made by Benji Taylor , Dennis Jin , and Alex Vanderzon
