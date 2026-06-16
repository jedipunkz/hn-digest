---
source: "https://vulnfeed.novadyne.ai/"
hn_url: "https://news.ycombinator.com/item?id=48558710"
title: "Show HN: VulnFeed – 9 security tools your AI agent can call (MCP server)"
article_title: "VulnFeed — Dependency Vulnerability Monitoring for Claude Code"
author: "ngburke"
captured_at: "2026-06-16T19:20:26Z"
capture_tool: "hn-digest"
hn_id: 48558710
score: 1
comments: 1
posted_at: "2026-06-16T17:25:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: VulnFeed – 9 security tools your AI agent can call (MCP server)

- HN: [48558710](https://news.ycombinator.com/item?id=48558710)
- Source: [vulnfeed.novadyne.ai](https://vulnfeed.novadyne.ai/)
- Score: 1
- Comments: 1
- Posted: 2026-06-16T17:25:50Z

## Translation

タイトル: HN を表示: VulnFeed – AI エージェントが呼び出すことができる 9 つのセキュリティ ツール (MCP サーバー)
記事のタイトル: VulnFeed — クロード コードの依存関係の脆弱性の監視
説明: プロジェクトを監視する MCP サーバー

記事本文:
Novadyne による Vulnfeed
仕組み
比較する
セットアップ
ブログ
CVEデータベース
依存関係がいつ脆弱になるかを把握します。
MCP サーバーは、ロックファイルを読み取り、NVD + GitHub アドバイザリをチェックし、実際に重要なことを、実際のエクスプロイト確率に基づいて優先順位を付けて、正確な修正バージョンとともに通知します。
無料枠 — 1 日あたり 10 スキャン、サインアップなし。月額 14 ドルで無制限。
クロードに確認してもらったらどうでしょうか？
package-lock.json 、requirements.txt 、または go.sum を読み取り、実際の依存関係ツリーにヒットする CVE のみにフィルターをかけます。使用しないパッケージからノイズが発生することはありません。
ほとんどの CVE はノイズです。 EPSS (エクスプロイト予測スコアリング システム) は、現実世界のエクスプロイト可能性によってそれぞれをスコア付けします。 VulnFeed は、実際の攻撃で使用される可能性のあるものを明らかにします。
「脆弱です」だけではなく、Express 4.17.1 → 4.21.0 にアップグレードしてください。問題を修正する正確なバージョンについては、npm、PyPI、および Go レジストリを相互参照します。
プロジェクトを一度登録してください。新しい脆弱性がないかいつでもチェックしてください。新しい CVE は午前 3 時に公開されますか?午前のセッションの午前 3 時 15 分までにインデックスに登録されます。
ロックファイルのスキャン、パッケージのチェック、CVE の検索、プロジェクトの監視、アラートのチェック、deps の更新、プロジェクトのリスト。セキュリティ ワークフローに必要なものがすべて揃っています。
データ ソースは NVD、GitHub Advisory DB、EPSS であり、すべて無料のパブリック API です。ベンダーのロックインやデータ ブローカーの仲介者はいません。 14 ドルはデータ アクセスではなく、インテリジェンス レイヤーに支払われます。
無料利用枠 — サインアップや API キーは不要
1 日あたり 10 スキャン、1 つの監視プロジェクト。これを MCP 構成に追加するだけです。
{
"mcpサーバー": {
"ヴァルンフィード": {
"コマンド": "uvx",
"args": ["vulnfeed-mcp"]
}
}
}
Claude Code、Claude Desktop、Cursor、VS Code、および Windsurf で動作します。
カーソルに追加
VS コードに追加
ウィンドサーフィンに追加
クライアントを再起動します。私のプロジェクトの脆弱性をスキャンするように依頼します。それでおしまい。
無制限のスキャン、無制限の監視対象プロ

ジェクト。ライセンス キーを追加します。
{
"mcpサーバー": {
"ヴァルンフィード": {
"コマンド": "uvx",
"args": ["vulnfeed-mcp"],
"環境": {
"VULNFEED_API_KEY": "YOUR_LICENSE_KEY_HERE"
}
}
}
}
ライセンス キーを取得します。定額料金、シートごと、リポジトリごとではありません。
Pay-per-Scan — x402 マイクロペイメント
AI エージェントは、Base 上の USDC を使用してリクエストごとに支払うことができます。アカウント、API キー、サブスクリプションは必要ありません。エージェントは 402 応答を受け取り、0.01 ドルを支払い、結果を受け取ります。あらゆる x402 互換クライアントで動作します。
# エージェントがリクエストを送信し、支払いの詳細を含む HTTP 402 を取得します
# x402 クライアント ライブラリは支払いを自動的に処理します
# スキャンあたり 0.01 ドル · CVE ルックアップあたり 0.002 ドル · モニターあたり 0.05 ドル
# 検出エンドポイント:
カール https://vulnfeed-api.novadyne.ai/.well-known/x402
x402 プロトコルを使用します — Coinbase ファシリテーター経由の Base 上の USDC。仲介業者なし、即時決済。料金とエンドポイントを表示します。
依存関係の監視を開始します。
Novadyne による VulnFeed 。・お問い合わせ
GitHub · PyPI · データ: NVD、GitHub Advisory DB、EPSS。すべて無料の公開ソース。
また、Novadyne による: Agents — あなたのビジネスのための AI エージェント。 · 元帳 — 複式簿記の会計 API。 · パワーパック — クロードコードスキルキット。
© 2026 ノバダイン。無断転載を禁じます。

## Original Extract

An MCP server that monitors your project

vulnfeed by Novadyne
How it works
Compare
Setup
Blog
CVE Database
Know when your dependencies are vulnerable.
An MCP server that reads your lockfile, checks NVD + GitHub Advisories, and tells you what actually matters — prioritized by real-world exploit probability, with exact fix versions.
Free tier — 10 scans/day, no signup. $14/mo for unlimited.
Why not just ask Claude to check?
Reads your package-lock.json , requirements.txt , or go.sum and filters to only the CVEs that hit your actual dependency tree. No noise from packages you don't use.
Most CVEs are noise. EPSS (Exploit Prediction Scoring System) scores each one by real-world exploitability. VulnFeed surfaces the ones likely to be used in real attacks.
Not just "you're vulnerable" but upgrade express 4.17.1 → 4.21.0 . Cross-references npm, PyPI, and Go registries for the exact version that fixes the issue.
Register your project once. Check back any time for new vulnerabilities. New CVE published at 3am? It's in the index by 3:15am for your morning session.
Scan a lockfile, check a package, look up a CVE, monitor a project, check alerts, update deps, list projects. Everything a security workflow needs.
Data sources are NVD, GitHub Advisory DB, and EPSS — all free, public APIs. No vendor lock-in, no data broker middlemen. Your $14 pays for the intelligence layer, not data access.
Free tier — no signup, no API key
10 scans/day, 1 monitored project. Just add this to your MCP config:
{
"mcpServers": {
"vulnfeed": {
"command": "uvx",
"args": ["vulnfeed-mcp"]
}
}
}
Works in Claude Code, Claude Desktop, Cursor, VS Code, and Windsurf.
Add to Cursor
Add to VS Code
Add to Windsurf
Restart your client. Ask it to scan my project for vulnerabilities . That's it.
Unlimited scans, unlimited monitored projects. Add your license key:
{
"mcpServers": {
"vulnfeed": {
"command": "uvx",
"args": ["vulnfeed-mcp"],
"env": {
"VULNFEED_API_KEY": "YOUR_LICENSE_KEY_HERE"
}
}
}
}
Get your license key — flat rate, not per-seat, not per-repo.
Pay-per-scan — x402 micropayments
AI agents can pay per request with USDC on Base — no account, no API key, no subscription. Your agent gets a 402 response, pays $0.01, and gets results. Works with any x402-compatible client.
# Agent sends request, gets HTTP 402 with payment details
# x402 client library handles payment automatically
# $0.01 per scan · $0.002 per CVE lookup · $0.05 per monitor
# Discovery endpoint:
curl https://vulnfeed-api.novadyne.ai/.well-known/x402
Uses the x402 protocol — USDC on Base via Coinbase facilitator. No middleman, instant settlement. View pricing & endpoints .
Start monitoring your dependencies.
VulnFeed by Novadyne . · Contact
GitHub · PyPI · Data: NVD, GitHub Advisory DB, EPSS. All free public sources.
Also by Novadyne: Agents — AI agents for your business. · Ledger — double-entry accounting API. · Power Pack — Claude Code skills kit.
© 2026 Novadyne. All rights reserved.
