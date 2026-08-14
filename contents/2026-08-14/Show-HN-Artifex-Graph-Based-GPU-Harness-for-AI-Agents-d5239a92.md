---
source: "https://gatewai.studio/artifex"
hn_url: "https://news.ycombinator.com/item?id=49297723"
title: "Show HN: Artifex - Graph Based GPU Harness for AI Agents"
article_title: "Loading..."
author: "oknaslnkn"
captured_at: "2026-08-14T12:41:53Z"
capture_tool: "hn-digest"
hn_id: 49297723
score: 2
comments: 0
posted_at: "2026-08-14T12:15:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Artifex - Graph Based GPU Harness for AI Agents

- HN: [49297723](https://news.ycombinator.com/item?id=49297723)
- Source: [gatewai.studio](https://gatewai.studio/artifex)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T12:15:13Z

## Translation

タイトル: Show HN: Artifex - AI エージェント用のグラフベースの GPU ハーネス
記事タイトル: 読み込み中...
HN テキスト: Artifex は、自律コーディング エージェントがローカルでメディア ノード グラフを作成、検証、レンダリングできるように構築された、マシンファーストのヘッドレス CLI ランタイムです。昨日リリースされた Deepseek ハーネスに似ています。 Artifex は、モジュラー実行環境でモデルを構築するため、ノードをプラグインとして備えた DAG エンジンを使用してツールとパイプラインを調整するための構造化されたマシン コントラクトをエージェントに提供します。各ノードには、グラフ処理、WebGPU レンダリング、オーディオ処理、および独自の SKILL.md ファイルにロジックを挿入する機能があります。ノードは、その React コンポーネント (CLI では利用できません) を注入することもできます。これは、デスクトップ アプリで利用可能になります。実行はトポロジカルであり、状態フラグ (--state および --from-state) を介したチェックポイント キャッシュをサポートしているため、エージェントは高価で長いアップストリーム生成呼び出し (Fal AI を使用したビデオ生成や OpenRouter を使用したエージェント ノードなど) を再計算することなく、レイアウト座標やダウンストリーム フィルターを調整できます。アニメーション、ブレンディング、メディア カット、クロップ、カラー LUT、フィルター グラフ、オーディオ エフェクト、その他のメディア編集ノードを含むマルチトラック ビジュアル コンポジション タイムラインは、ローカル GPU ランタイム上で直接コンパイルおよびレンダリングされ、パイプラインの各ステップでの確定的なオフライン実行と中間アーティファクトを保証します。 CLI[0] はまだ完全にはオープンソースではありませんが、そのリポジトリにはビルド バンドルが含まれています。計画では、Web アプリケーションを強化するモノレポから切り離した後、オープンソースにする予定です [1]。私たちのロードマップには、エージェント/プログラマがカスタム ユースケース用に独自のファイル システムでローカルにプラグインを構築するための SDK も含まれています。 [0]: https://github.com/gatewai-dev/artifex
[1]: https://gatewai.studio

記事本文:
読み込み中...

## Original Extract

Artifex is a machine-first, headless CLI runtime built for autonomous coding agents to author, validate, and render media node graphs locally. Similar to Deepseek harness released yesterday; which ground models in modular execution environments, Artifex gives agents a structured machine contract to orchestrate tools and pipelines using DAG engine with nodes as plugins. Each node has capability to inject logic into graph processing, WebGPU rendering, audio processing and their own SKILL.md file. Nodes can also inject their react components (not available with CLI) - which will be available with the desktop app. Execution is topological and supports checkpoint caching via state flags (--state and --from-state), allowing agents to adjust layout coordinates or downstream filters without recomputing expensive and long upstream generation calls (such as Video Generation using Fal AI or an Agent Node using OpenRouter). Multi-track visual composition timeline with animations, blending, media cut, crop, color LUTs, filter graphs, audio effects and other media editing node's are compiled and rendered directly on the local GPU runtime, ensuring deterministic offline execution and intermediate artifacts at each step of the pipeline. The CLI[0] is not fully open source yet, but its repository contains the build bundle. The plan is to make it open source after we decouple it from our monorepo that powers our web application [1]. Our roadmap also contains a SDK for agents/programmers to build plugins locally in their own filesystem for custom use-cases. [0]: https://github.com/gatewai-dev/artifex
[1]: https://gatewai.studio

Loading...
