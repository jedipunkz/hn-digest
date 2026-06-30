---
source: "https://cli.revos.dev/"
hn_url: "https://news.ycombinator.com/item?id=48733674"
title: "Show HN: Scaffold a BigQuery and dbt and Cube project an AI agent can operate"
article_title: "Getting Started | RevOS CLI"
author: "zubairov"
captured_at: "2026-06-30T15:50:18Z"
capture_tool: "hn-digest"
hn_id: 48733674
score: 6
comments: 2
posted_at: "2026-06-30T15:00:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Scaffold a BigQuery and dbt and Cube project an AI agent can operate

- HN: [48733674](https://news.ycombinator.com/item?id=48733674)
- Source: [cli.revos.dev](https://cli.revos.dev/)
- Score: 6
- Comments: 2
- Posted: 2026-06-30T15:00:04Z

## Translation

タイトル: HN の表示: AI エージェントが操作できる BigQuery、dbt、および Cube プロジェクトのスキャフォールディング
記事のタイトル: はじめに | RevOS CLI
説明: RevOS リソース (テーブル、スコア、セグメント、アクション、キューブ、組織など) を管理するためのコマンドライン インターフェイス。

記事本文:
メイン コンテンツにスキップ RevOS CLI ドキュメント プラットフォーム ドキュメント Web サイト はじめに
RevOS リソース (テーブル、スコア、セグメント、アクション、キューブ、組織など) を管理するためのコマンドライン インターフェイス。
# インストール
npm install -g @revos/cli
# ブラウザ経由でログイン
revos 認証ログイン
# 組織を確認してください
revos 組織リスト
# 新しいデータプロジェクトを初期化する
revos init my-project
あなたにできること
認証 — ブラウザベースの OAuth ログイン、CI/CD のトークンベースの認証
組織の管理 — 組織のリスト、切り替え、検査
リソースの管理 — CRUD テーブル、スコア、セグメント、アクション、AI 命令、サービス アカウント
IaC 経由でキューブとパイプラインを管理 — 接続とキューブを YAML として宣言します。 revos apply と調整し、 revos diff および revos status で検査します
足場プロジェクト — メダリオン レイアウト、開発コンテナ、dbt、および AI コンパニオン ファイルを使用してデータ エンジニアリング プロジェクトを生成します。
ゼロからセマンティック モデルへ — 実際のデータセットを使用したエンドツーエンドのウォークスルー
インストール - CLI をインストールして確認する
認証 — 資格情報を設定します
構成 - グローバル オプションと環境変数

## Original Extract

Command-line interface for managing RevOS resources — tables, scores, segments, actions, cubes, organizations, and more.

Skip to main content RevOS CLI Docs Platform Docs Website Getting Started
Command-line interface for managing RevOS resources — tables, scores, segments, actions, cubes, organizations, and more.
# Install
npm install -g @revos/cli
# Login via browser
revos auth login
# Check your organizations
revos org list
# Initialize a new data project
revos init my-project
What You Can Do ​
Authenticate — Browser-based OAuth login, token-based auth for CI/CD
Manage Organizations — List, switch, and inspect orgs
Manage Resources — CRUD tables, scores, segments, actions, AI instructions, and service accounts
Manage Cubes & Pipelines via IaC — Declare Connections and Cubes as YAML; reconcile with revos apply , inspect with revos diff and revos status
Scaffold Projects — Generate data engineering projects with medallion layout, Dev Containers, dbt, and AI companion files
From Zero to Semantic Model — End-to-end walkthrough using a real dataset
Installation — Install and verify the CLI
Authentication — Set up your credentials
Configuration — Global options and environment variables
