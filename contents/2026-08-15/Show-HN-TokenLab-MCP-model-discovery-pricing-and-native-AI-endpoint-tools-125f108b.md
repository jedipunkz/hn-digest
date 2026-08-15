---
source: "https://tokenlab.sh/zh/mcp"
hn_url: "https://news.ycombinator.com/item?id=49314483"
title: "Show HN: TokenLab MCP, model discovery, pricing, and native AI endpoint tools"
article_title: "TokenLab MCP｜让 AI 智能体调用模型与多模态工具"
author: "tokenlabai"
captured_at: "2026-08-15T22:10:49Z"
capture_tool: "hn-digest"
hn_id: 49314483
score: 2
comments: 0
posted_at: "2026-08-15T21:26:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: TokenLab MCP, model discovery, pricing, and native AI endpoint tools

- HN: [49314483](https://news.ycombinator.com/item?id=49314483)
- Source: [tokenlab.sh](https://tokenlab.sh/zh/mcp)
- Score: 2
- Comments: 0
- Posted: 2026-08-15T21:26:34Z

## Translation

タイトル: HN を表示: TokenLab MCP、モデル検出、価格設定、およびネイティブ AI エンドポイント ツール
記事タイトル：TokenLab MCP｜AIエージェントにモデルやマルチモーダルツールを呼び出してもらう
説明: TokenLab を Claude、Codex、Cursor などの MCP クライアントに接続します。モデルを確認し、価格を比較し、テキスト、画像、ビデオ、オーディオなどの 31 または 80 のツールにアクセスします。

記事本文:
TokenLab MCP｜AI エージェントにモデルとマルチモーダル ツールを呼び出しさせる ホーム / MCP TOKENLAB MCP
Claude、Codex、Cursor などの MCP クライアントに接続して、モデルを確認し、価格を比較し、テキスト、画像、ビデオ、音楽、3D、オーディオ機能を呼び出します。よく使用される 31 個のツールがデフォルトで有効になっており、必要に応じて 80 個まで拡張できます。
MCP クイック構成は直接インストールできます。 MCP クライアントに TokenLab を追加する
クライアント設定をコピーし、API キーを置き換えます。モデルと価格のみを問い合わせる場合は、キーを入力せずに使用できます。
API キーを取得 {
"mcpサーバー": {
"トークンラボ": {
"コマンド": "npx",
"引数": [
「-y」、
「@tokenlabai/mcp-server」
]、
"環境": {
"TOKENLAB_API_KEY": "<YOUR_TOKENLAB_API_KEY>",
"TOKENLAB_MCP_TOOL_PROFILE": "コア"
}
}
}
構成をコピーし、コアを完全に変更して、80 個のツールすべてを有効にします。
よく使用される 31 のツールでほとんどの通話をカバーします。バッチ処理、応答管理、またはメディア資産管理が必要な場合は、フル モードを有効にします。
チャットの完了、応答、メッセージ、Gemini に加え、画像、ビデオ、音楽、3D、オーディオをサポートします。
レスポンス、バッチリクエスト、ワールド、メディア資産管理用のツールを 49 個追加します。
オンライン版はモデルと価格の確認のみに適しています。ローカル バージョンでは TokenLab の全機能を呼び出すことができます。
stdio · 31 / 80 ツール TokenLab の全機能を呼び出す
API キーを追加すると、テキスト、マルチモーダル コンテンツ、ファイル、タスク、埋め込み、翻訳を操作するための 31 または 80 のツールを使用できます。
Claude Desktop、Codex、Cline、Cursor、およびその他の標準入出力クライアントの完全な構成例を参照してください。
すべての AI モデルに 1 つの API - 統合された API を通じて数百の AI モデルにアクセスします。
SSL 暗号化 GDPR 準拠 Claude で構築 | GPT Copyright © 2026 TOKENLAB AI INC. 全著作権所有。
TokenLab は、独立したサードパーティ サービス プロバイダーです。当社は、OpenAI、Anthrop を含むがこれらに限定されない、AI モデル プロバイダーと提携、承認、または正式に接続されていません。

ic、Google、Microsoft、Meta、Mistral、Cohere、DeepSeek、xAI、Stability AI、またはそれらの子会社または関連会社。すべての商標はそれぞれの所有者の財産です。
すべての記事をブログに掲載する チュートリアル、製品の説明、比較分析、エンジニアリングの最新情報、および調査レポートを読んでください。分類
🌐 簡体字中国語ログイン メニューを開く TokenLab ログイン モデル モデル データ トレンド 研究ランキング MCP ドキュメント ブログ チュートリアル 15 記事 機能 12 記事 ニュース 3 記事 比較 30 記事 エンジニアリング 6 記事 インサイト 8 記事 研究 4 記事 設定

## Original Extract

把 TokenLab 接入 Claude、Codex、Cursor 等 MCP 客户端。查模型、比价格，并调用文本、图像、视频、音频等 31 或 80 个工具。

TokenLab MCP｜让 AI 智能体调用模型与多模态工具 首页 / MCP TOKENLAB MCP
接入 Claude、Codex、Cursor 等 MCP 客户端，查模型、比价格，调用文本、图像、视频、音乐、3D 与音频能力。默认启用 31 个常用工具，需要时可扩展到 80 个。
MCP 快速配置 可直接安装 添加 TokenLab 到你的 MCP 客户端
复制客户端配置并替换 API 密钥。只查询模型和价格时，不填密钥也能使用。
获取 API 密钥 {
"mcpServers": {
"tokenlab": {
"command": "npx",
"args": [
"-y",
"@tokenlabai/mcp-server"
],
"env": {
"TOKENLAB_API_KEY": "<YOUR_TOKENLAB_API_KEY>",
"TOKENLAB_MCP_TOOL_PROFILE": "core"
}
}
}
} 复制配置 将 core 改为 full，即可启用全部 80 个工具。
31 个常用工具覆盖大多数调用；需要批处理、Responses 管理或媒体素材管理时，再启用完整模式。
支持 Chat Completions、Responses、Messages、Gemini，以及图像、视频、音乐、3D 和音频。
再增加 49 个用于 Responses、批量请求、World 和媒体素材管理的工具。
在线版适合只查模型和价格；本地版可调用 TokenLab 的完整能力。
stdio · 31 / 80 个工具 调用 TokenLab 全部能力
添加 API 密钥以使用 31 或 80 工具来处理文本、多模态内容、文件、任务、嵌入和翻译。
查看 Claude Desktop、Codex、Cline、Cursor 和其他 stdio 客户端的完整配置示例。
一个 API 适配所有 AI 模型 - 通过统一的 API 访问数百个 AI 模型。
SSL 加密 符合 GDPR 标准 使用 Claude | GPT 构建 Copyright © 2026 TOKENLAB AI INC . 版权所有。
TokenLab 是一家独立的第三方服务提供商。我们与任何 AI 模型提供商均无隶属、认可或官方联系，包括但不限于 OpenAI、Anthropic、Google、Microsoft、Meta、Mistral、Cohere、DeepSeek、xAI、Stability AI 或其任何子公司或关联公司。所有商标均为其各自所有者的财产。
博客 所有文章 阅读教程、产品说明、对比分析、工程更新和研究报告。 分类
🌐 简体中文 登录 打开菜单 TokenLab 登录 模型 模型数据 趋势 研究 排行榜 MCP 文档 博客 教程 15 篇 功能 12 篇 新闻 3 篇 对比 30 篇 工程 6 篇 洞察 8 篇 研究 4 篇 设置
