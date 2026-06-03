---
source: "https://github.com/Dvbxtreme/ai-sales-agent"
hn_url: "https://news.ycombinator.com/item?id=48369940"
title: "Open-source AI Sales Agent with Next.js 15 and Ollama – zero API costs"
article_title: "GitHub - Dvbxtreme/ai-sales-agent: AI Sales Agent - Next.js 15 chatbot, content generator, lead finder with Ollama · GitHub"
author: "sdev99"
captured_at: "2026-06-03T00:46:27Z"
capture_tool: "hn-digest"
hn_id: 48369940
score: 2
comments: 0
posted_at: "2026-06-02T13:21:04Z"
tags:
  - hacker-news
  - translated
---

# Open-source AI Sales Agent with Next.js 15 and Ollama – zero API costs

- HN: [48369940](https://news.ycombinator.com/item?id=48369940)
- Source: [github.com](https://github.com/Dvbxtreme/ai-sales-agent)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T13:21:04Z

## Translation

タイトル: Next.js 15 と Ollama を使用したオープンソース AI セールス エージェント – API コストゼロ
記事のタイトル: GitHub - Dvbxtreme/ai-sales-agent: AI セールス エージェント - Next.js 15 チャットボット、コンテンツ ジェネレーター、Ollama によるリード ファインダー · GitHub
説明: AI セールス エージェント - Next.js 15 チャットボット、コンテンツ ジェネレーター、Ollama を使用したリード ファインダー - Dvbxtreme/ai-sales-agent

記事本文:
GitHub - Dvbxtreme/ai-sales-agent: AI セールス エージェント - Next.js 15 チャットボット、コンテンツ ジェネレーター、Ollama によるリード ファインダー · GitHub
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
DVBエクストリーム
/
ai販売代理店
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビ

イゲーションオプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット public public src src .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Next.js 15 で構築されたオープンソースの AI を活用した販売エージェント。コンテンツの生成、リードの発見、アウトリーチの自動化はすべて Ollama を使用してローカルで実行されます (API コストはかかりません)。
AI チャットボット — Ollama を利用した会話型 AI アシスタント (ローカル、無料、API キーなし)
コンテンツ ジェネレーター — ブログ投稿、ツイート、LinkedIn 投稿、電子メールを自動生成します。
Lead Finder — 潜在的な顧客を Web からかき集めます
自動アウトリーチ — パーソナライズされたアウトリーチ電子メールを生成
管理者ダッシュボード — リード、コンテンツ、設定を管理する
ブログ — AI が生成したコンテンツを含む組み込みブログ
レイヤー
テクノロジー
フレームワーク
Next.js 15 (アプリルーター)
AI
Ollama (ローカル LLM、無料)
データベース
SQLite (better-sqlite3 経由)
認証
JWT (httpOnly Cookie)
スタイリング
追い風CSS
はじめに
# 依存関係をインストールする
npmインストール
# 環境をセットアップする
cp .env.example .env.local
# Ollama を実行する (AI 機能に必要)
# https://ollama.com からダウンロードして実行します。
オラマ プル ラマ3
# 開発サーバーを起動します
npm 実行開発
ブラウザで http://localhost:3001 を開きます。
npm ビルドを実行する
npmスタート
完全な SaaS キットが必要ですか?
この AI Sales Agent は、Stripe サブスクリプション、OpenAI ストリーミング、JWT 認証、管理ダッシュボード、Docker デプロイメントを備えた本番環境に対応した SaaS ボイラープレートである AI SaaS スターター キットと完全に連携します。
1回997ドル。商用ライセンス。無制限のプロジェクト。
AI セールス エージェント - Next.js 15 チャットボット、c

Ollama によるコンテンツ ジェネレーター、リード ファインダー
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI Sales Agent - Next.js 15 chatbot, content generator, lead finder with Ollama - Dvbxtreme/ai-sales-agent

GitHub - Dvbxtreme/ai-sales-agent: AI Sales Agent - Next.js 15 chatbot, content generator, lead finder with Ollama · GitHub
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
Dvbxtreme
/
ai-sales-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit public public src src .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json View all files Repository files navigation
An open-source, AI-powered sales agent built with Next.js 15 . Generates content, finds leads, and automates outreach — all running locally with Ollama (no API costs).
AI Chatbot — Conversational AI assistant powered by Ollama (local, free, no API key)
Content Generator — Auto-generate blog posts, tweets, LinkedIn posts, and emails
Lead Finder — Scrape potential customers from the web
Auto Outreach — Generate personalized outreach emails
Admin Dashboard — Manage leads, content, and settings
Blog — Built-in blog with AI-generated content
Layer
Technology
Framework
Next.js 15 (App Router)
AI
Ollama (local LLM, free)
Database
SQLite (via better-sqlite3)
Auth
JWT (httpOnly cookies)
Styling
Tailwind CSS
Getting Started
# Install dependencies
npm install
# Set up environment
cp .env.example .env.local
# Run Ollama (required for AI features)
# Download from https://ollama.com and run:
ollama pull llama3
# Start dev server
npm run dev
Open http://localhost:3001 in your browser.
npm run build
npm start
Need a Complete SaaS Kit?
This AI Sales Agent pairs perfectly with the AI SaaS Starter Kit — a production-ready SaaS boilerplate with Stripe subscriptions, OpenAI streaming, JWT auth, admin dashboard, and Docker deployment.
$997 one-time. Commercial license. Unlimited projects.
AI Sales Agent - Next.js 15 chatbot, content generator, lead finder with Ollama
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
