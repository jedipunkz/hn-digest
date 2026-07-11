---
source: "https://open-sandbox.ai/"
hn_url: "https://news.ycombinator.com/item?id=48874053"
title: "OpenSandbox Universal Sandbox Infrastructure for AI Applications"
article_title: "OpenSandbox"
author: "hek2sch"
captured_at: "2026-07-11T18:45:10Z"
capture_tool: "hn-digest"
hn_id: 48874053
score: 1
comments: 0
posted_at: "2026-07-11T17:45:49Z"
tags:
  - hacker-news
  - translated
---

# OpenSandbox Universal Sandbox Infrastructure for AI Applications

- HN: [48874053](https://news.ycombinator.com/item?id=48874053)
- Source: [open-sandbox.ai](https://open-sandbox.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T17:45:49Z

## Translation

タイトル: OpenSandbox AI アプリケーション向けユニバーサル サンドボックス インフラストラクチャ
記事タイトル: OpenSandbox
説明: AI アプリケーション用のユニバーサル サンドボックス インフラストラクチャ

記事本文:
コンテンツへスキップ OpenSandbox 検索 K メイン ナビゲーション スタートガイド リファレンス SDK API 仕様 CLI コンポーネント Kubernetes 移行ガイド 例 コミュニティの外観
OpenSandbox AI アプリケーション用のユニバーサル サンドボックス インフラストラクチャ
多言語 SDK を使用して、分離された環境でコマンド、コード インタープリター、ブラウザ、開発者ツールを安全に実行します。
Docker および Kubernetes ランタイムを使用して、サンドボックス インスタンスをプロビジョニング、監視、更新、一時停止/再開、終了します。
標準化されたライフサイクルおよび実行プロトコルに基づいて、Python、Java/Kotlin、JavaScript/TypeScript、C#/.NET、Go SDK を使用して構築します。
シェル コマンドの実行、ファイルの管理、多言語コード インタープリターの実行、ポートの公開、ログとメトリックのストリーミングを行います。
コーディング エージェント、ブラウザ自動化、リモート開発、コード実行、強化学習シナリオをサポートします。
OpenSandbox は CNCF Landscape にリストされています。
Claude Code、Gemini CLI、Codex、およびその他のコーディング エージェントを分離されたサンドボックスで実行します。
制御されたランタイム、ファイルシステム、ネットワークを使用して Chrome および Playwright ワークロードを実行します。
安全なクラウド開発ワークフローのための VS Code Web およびデスクトップのような環境をホストします。
モデルで生成されたコードを安全に実行し、出力をストリーミングし、再現可能な環境で迅速に反復します。
管理されたサンドボックスのライフサイクルとリソース制御を使用して強化学習タスクを開始します。
例 ですべての例を確認してください。
pip install opensandbox bash npm install @alibaba-group/opensandbox kotlin 依存関係 {
実装 ( "com.alibaba.opensandbox:sandbox:{latest_version}" )
} bash go get github.com/alibaba/OpenSandbox/sdks/sandbox/go bash dotnet add package Alibaba.OpenSandbox 完全なセットアップ ガイドについては、「はじめに」を参照してください。
Apache 2.0 ライセンスに基づいてリリースされています。
著作権 © 2024 年現在 OpenSandbox 貢献者

## Original Extract

Universal Sandbox Infrastructure for AI Applications

Skip to content OpenSandbox Search K Main Navigation Getting Started Guides Reference SDKs API Specs CLI Components Kubernetes Migration Guides Examples Community Appearance
OpenSandbox Universal Sandbox Infrastructure for AI Applications
Securely run commands, code interpreters, browsers, and developer tools in isolated environments with multi-language SDKs.
Provision, monitor, renew, pause/resume, and terminate sandbox instances with Docker and Kubernetes runtimes.
Build with Python, Java/Kotlin, JavaScript/TypeScript, C#/.NET, and Go SDKs on top of standardized lifecycle and execution protocols.
Execute shell commands, manage files, run multi-language code interpreters, expose ports, and stream logs and metrics.
Supports coding agents, browser automation, remote development, code execution, and reinforcement learning scenarios.
OpenSandbox is listed in the CNCF Landscape .
Run Claude Code, Gemini CLI, Codex, and other coding agents in isolated sandboxes.
Execute Chrome and Playwright workloads with controlled runtime, filesystem, and networking.
Host VS Code Web and desktop-like environments for secure cloud development workflows.
Run model-generated code safely, stream outputs, and iterate quickly with reproducible environments.
Launch reinforcement learning tasks with managed sandbox lifecycle and resource controls.
Explore all examples in Examples .
pip install opensandbox bash npm install @alibaba-group/opensandbox kotlin dependencies {
implementation ( "com.alibaba.opensandbox:sandbox:{latest_version}" )
} bash go get github.com/alibaba/OpenSandbox/sdks/sandbox/go bash dotnet add package Alibaba.OpenSandbox See Getting Started for the full setup guide.
Released under the Apache 2.0 License.
Copyright © 2024-present OpenSandbox Contributors
