---
source: "https://github.com/chenxiachan/thoughtdag"
hn_url: "https://news.ycombinator.com/item?id=49057101"
title: "I built a canvas where LLM chats become an editable map"
article_title: "GitHub - chenxiachan/thoughtdag: Your thinking deserves a map: an infinite canvas where LLM conversations grow into an editable thought graph. Wires are the context. · GitHub"
author: "chatchan"
captured_at: "2026-07-26T11:55:02Z"
capture_tool: "hn-digest"
hn_id: 49057101
score: 1
comments: 0
posted_at: "2026-07-26T11:42:22Z"
tags:
  - hacker-news
  - translated
---

# I built a canvas where LLM chats become an editable map

- HN: [49057101](https://news.ycombinator.com/item?id=49057101)
- Source: [github.com](https://github.com/chenxiachan/thoughtdag)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T11:42:22Z

## Translation

タイトル: LLM チャットが編集可能なマップになるキャンバスを構築しました
記事のタイトル: GitHub - chenxiachan/thoughtdag: あなたの思考には地図が必要です: LLM の会話が編集可能な思考グラフに成長する無限のキャンバス。ワイヤーはコンテキストです。 · GitHub
説明: あなたの思考にはマップが必要です。LLM の会話が編集可能な思考グラフに成長する無限のキャンバスです。ワイヤーはコンテキストです。 - チェンシアチャン/thoughtdag

記事本文:
GitHub - chenxiachan/thoughtdag: あなたの思考には地図が必要です。LLM の会話が編集可能な思考グラフに成長する無限のキャンバスです。ワイヤーはコンテキストです。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
チェンシアチャン
/
思考ダグ
プ

ブリック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
292 コミット 292 コミット docs docs 関数/ API 関数/ API public public scripts scripts src src video video .env.example .env.example .gitignore .gitignore CITATION.cff CITATION.cff DEPLOYING.md DEPLOYING.md ライセンス ライセンス README.md README.md README_ZH.md README_ZH.md eslint.config.js eslint.config.js Index.html Index.html mcp.config.example.json mcp.config.example.json package-lock.json package-lock.json package.json package.json server.mjs server.mjs tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts worker.js worker.js wrangler.jsonc wrangler.jsonc すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あなたの考えには地図が必要です。 LLM の会話が編集可能な思考グラフに成長する無限のキャンバス。
インストールもサインアップも必要ありません。サンプルのキャンバスにはキーは必要ありません
中文 · クイックスタート · その他の機能 · モデル · コストとプライバシー
ワイヤーはコンテキストです。モデルが認識しているものは、まさにノードに接続されているものです。グラフを編集すると、モデルのメモリが編集されます。
すべてのジェスチャーの背後にある 1 つの原則: 人間はループの中にあり、モデルはワイヤー上にあります。自律エージェントがグラフを再描画することはありません。
✂️ 1 つのエッジを削除すると、別の答えが得られます
モデルは配線内容のみを認識します。ノイズ エッジを削除して再度質問すると、同じプロンプトで明確な回答が返されます。サンプルキャンバスの③章で再現します。
パッセージを選択して、その場で質問してください。答えはページ番号とともにキャンバスに表示され、p.N チップはそのページに戻ります。紙を書き終えると、地図が描かれます。
💎 思考は手に凝縮
ノードを 1 つの高さにマージする

えー結論；ハイライトを要約に織り込みます。グラフは広がるのではなく、内側に折りたたまれます。人間はループの中で磨きをかけます。
🖍️ あなたがマークした箇所は、引用された散文に織り込まれています
ハイライトはモデルの判断ではなく、あなたの判断です。サブセットをチェックし、すべての文を遡って 1 つのパッセージを織ります。
🗺️ ズームアウト: 思考が地図になる
フルカード、持ち帰り用のプラーク、アイコンのスケルトン: 3 つのセマンティック層、各ステップにバッジが付けられます ✕ ⚖ ↩ ?。迂回路は地図の一部です。
# オンライン: app.thoughtdag.workers.dev (サンプル キャンバスにはキーは必要ありません)
# ローカル:
npmインストール
npm 実行サーバー # LLM プロキシ :3001
npm run dev # → localhost:5173
# .env はありませんか?アプリ内の OpenAI 互換エンドポイントに接続します
最初の起動では、シードされたサンプル キャンバスが開きます。日常の 1 つの質問 (保存された記事が読まれない理由) を中心とした 4 つの章で、実際の PDF が埋め込まれた読書ループが含まれます。環境変数、フリーキー、設定の詳細 → docs/setup.md
能力
何をするのか
📤 読み取り専用共有
1 つのリンクでグラフ全体が伝達されます。アカウントもサーバー ストレージもありません。
🧭 古いものとリプレイ
上流の編集では、無効となる回答にマークが付けられます。依存関係の順序で再生し、トークンを最初に推定します
🧪 パラダイム
ヒューマンマシンワークフローはファイルとして保存されます。入力を変更して実験を再現する
🔌 どのモデルでも
線に続くノードごとのピン。画像リクエストはビジョンモデルに自動的に再ルーティングされます
🔒 地元第一主義
自動フォルダー バックアップは実際のファイルを書き込みます。クロスデバイスの同期フォルダーをポイントします
全機能リスト (60 以上、領域ごとにグループ化) → docs/features.md
Zhipu、Qwen、OpenAI、Anthropic、Google、DeepSeek、Kimi、OpenRouter、Ollama、または任意の OpenAI 互換エンドポイント。画像を含むリクエストは自動的にビジョン モデルに再ルーティングされます。環境変数とデフォルトモデル → docs/setup.md
フランス

ee モデル層はすべての機能をカバーします。ローカルの Ollama は完全にオフラインで実行されます
ホストされたデモでは、モデルのトラフィックはブラウザーで直接実行されます。キーがサーバーに接触することはありません。
PDF がマシンから離れることはありません。質問すると、抽出されたテキストのみが送信されます
バックアップ形式には下位互換性が維持されます。マークダウン エクスポートは永続的な避難口です
グラフは非周期的です。あなたはそのループです。
MIT © 2026 Xia Chen · ロードマップ · フィードバック · 引用
あなたの思考には地図が必要です。LLM の会話が編集可能な思考グラフに成長する無限のキャンバスです。ワイヤーはコンテキストです。
app.thoughtdag.workers.dev トピック
Readme MIT ライセンス このリポジトリを引用する アクティビティのスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Your thinking deserves a map: an infinite canvas where LLM conversations grow into an editable thought graph. Wires are the context. - chenxiachan/thoughtdag

GitHub - chenxiachan/thoughtdag: Your thinking deserves a map: an infinite canvas where LLM conversations grow into an editable thought graph. Wires are the context. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
chenxiachan
/
thoughtdag
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
292 Commits 292 Commits docs docs functions/ api functions/ api public public scripts scripts src src video video .env.example .env.example .gitignore .gitignore CITATION.cff CITATION.cff DEPLOYING.md DEPLOYING.md LICENSE LICENSE README.md README.md README_ZH.md README_ZH.md eslint.config.js eslint.config.js index.html index.html mcp.config.example.json mcp.config.example.json package-lock.json package-lock.json package.json package.json server.mjs server.mjs tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts worker.js worker.js wrangler.jsonc wrangler.jsonc View all files Repository files navigation
Your thinking deserves a map. An infinite canvas where LLM conversations grow into an editable thought graph.
no install, no signup; the example canvas needs no key
中文 · Quick start · More capabilities · Models · Cost & privacy
Wires are the context. What the model sees is exactly what wires into the node. Editing the graph edits the model's memory.
One principle behind every gesture: the human in the loop, the model on the wires . No autonomous agent redraws your graph.
✂️ Delete one edge, get a different answer
The model sees only what wires in. Delete the noise edge, ask again, and the same prompt returns a clean answer. Reproduce it in chapter ③ of the example canvas.
Select a passage, ask right there. The answer lands on the canvas with its page number, and the p.N chip jumps back to the page. Finish the paper, and the map is drawn.
💎 Thinking condenses in your hands
Merge nodes into one higher conclusion; weave highlights into a summary. The graph folds inward instead of sprawling. The human refines in the loop.
🖍️ The passages you marked, woven into cited prose
Highlights are your judgment, not the model's. Check any subset and weave one passage where every sentence traces back.
🗺️ Zoom out: thinking becomes a map
Full cards, takeaway plaques, an icon skeleton: three semantic tiers, every step badged ✕ ⚖ ↩ ?. The detours are part of the map.
# Online: app.thoughtdag.workers.dev (example canvas needs no key)
# Local:
npm install
npm run server # LLM proxy :3001
npm run dev # → localhost:5173
# No .env? Connect any OpenAI-compatible endpoint inside the app
The first launch opens a seeded example canvas: four chapters around one everyday question (why saved articles stay unread), including a reading loop with a real embedded PDF. Environment variables, free keys and configuration details → docs/setup.md
Capability
What it does
📤 Read-only share
One link carries the whole graph: no account, no server storage
🧭 Staleness & replay
Upstream edits mark the answers they invalidate; replay in dependency order, token estimate first
🧪 Paradigms
Human-machine workflows saved as files; change the input, replay the experiment
🔌 Any model
Per-node pins that follow the line; image requests reroute to vision models automatically
🔒 Local-first
Automatic folder backup writes real files; point it at a synced folder for cross-device
Full feature list (60+, grouped by area) → docs/features.md
Zhipu · Qwen · OpenAI · Anthropic · Google · DeepSeek · Kimi · OpenRouter · Ollama, or any OpenAI-compatible endpoint. Requests with images reroute to vision models automatically. Environment variables and default models → docs/setup.md
The free model tier covers every feature ; a local Ollama runs fully offline
On the hosted demo, model traffic runs browser-direct : keys never touch the server
PDFs never leave your machine ; only extracted text travels when you ask
The backup format stays backward compatible ; Markdown export is the permanent escape hatch
The graph is acyclic. You are the loop.
MIT © 2026 Xia Chen · Roadmap · Feedback · Cite
Your thinking deserves a map: an infinite canvas where LLM conversations grow into an editable thought graph. Wires are the context.
app.thoughtdag.workers.dev Topics
Readme MIT license Cite this repository Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
