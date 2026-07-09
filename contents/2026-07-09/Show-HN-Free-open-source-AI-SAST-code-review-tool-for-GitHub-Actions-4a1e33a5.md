---
source: "https://hub.docker.com/r/kirill89/reviewcerberus"
hn_url: "https://news.ycombinator.com/item?id=48851684"
title: "Show HN: Free open source AI+SAST code review tool for GitHub Actions"
article_title: "kirill89/reviewcerberus - Docker Image"
author: "k1r111"
captured_at: "2026-07-09T20:42:24Z"
capture_tool: "hn-digest"
hn_id: 48851684
score: 1
comments: 0
posted_at: "2026-07-09T20:04:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Free open source AI+SAST code review tool for GitHub Actions

- HN: [48851684](https://news.ycombinator.com/item?id=48851684)
- Source: [hub.docker.com](https://hub.docker.com/r/kirill89/reviewcerberus)
- Score: 1
- Comments: 0
- Posted: 2026-07-09T20:04:41Z

## Translation

タイトル: Show HN: GitHub Actions 用の無料のオープンソース AI+SAST コード レビュー ツール
記事のタイトル: karill89/reviewcerberus - Docker Image

記事本文:
フォーラム ⁠ サポートへのお問い合わせ システムステータス ⁠ システムテーマ サインイン サインアップ 戻る
フォーラム ⁠ サポートへのお問い合わせ システムステータス ⁠ システムテーマ Docker Suite 探索
karill89 / reviewcerberus リポジトリの概要
git ブランチの違いを分析し、生成する AI を活用したコード レビュー ツール
構造化された出力を含む包括的なレビューレポート。
⁠ クイックスタート
docker run --rm -it -v $(pwd):/repo \
-e MODEL_PROVIDER=人類 \
-e ANTHROPIC_API_KEY=sk-ant-your-api-key \
キリル89/レビューケルベロス:最新\
--repo-path /repo --output /repo/review.md
コピー
それです！レビューは現在のファイルの review.md に保存されます。
ディレクトリ。
⁠ 主な機能
包括的なレビュー : ロジック、セキュリティ、パフォーマンス、
コードの品質
構造化された出力: 重大度別に整理された問題と要約表
マルチプロバイダー: AWS Bedrock、Anthropic API、Ollama、または Moonshot
スマートな分析: プロンプト キャッシュにより事前に提供されるコンテキスト
Git 統合 : あらゆるリポジトリで動作し、コミット ハッシュをサポートします
検証モード : 実験的
偽を減らすための検証チェーン ⁠
ポジティブ
⁠ 使用例
カスタムターゲットブランチ:
docker run --rm -it -v $(pwd):/repo \
-e MODEL_PROVIDER=人類 \
-e ANTHROPIC_API_KEY=sk-ant-your-api-key \
キリル89/レビューケルベロス:最新\
--repo-path /repo --target-branch 開発 --output /repo/review.md
コピー
カスタム レビュー ガイドラインを使用すると、次のようになります。
docker run --rm -it -v $(pwd):/repo \
-e MODEL_PROVIDER=人類 \
-e ANTHROPIC_API_KEY=sk-ant-your-api-key \
キリル89/レビューケルベロス:最新\
--repo-path /repo --instructions /repo/guidelines.md --output /repo/review.md
コピー
⁠ 含まれるもの
⁠ 包括的なコードレビュー
以下をカバーする詳細な分析:
論理と正確さ: バグ、エッジケース、エラー処理
セキュリティ: OWASP の問題、アクセス制御、入力検証
ペルフォ

ロマンス: N+1 クエリ、ボトルネック、スケーラビリティ
コードの品質: 重複、複雑さ、保守性
副作用: 他のシステム部分への影響
テスト: カバレッジギャップ、テストケースの欠落
ドキュメント: ドキュメントが不足している、または古い
⁠ 構造化された出力
すべてのレビューには次の内容が含まれます。
概要 : 変更とリスク領域の高レベルの概要
問題テーブル: 重大度インジケーター付きのすべての問題の一覧 (🔴 クリティカル、
🟠高、🟡中、🟢低)
詳細な問題: 各問題の説明、場所、および推奨される修正
⁠ 設定
⁠ 人間API
-e MODEL_PROVIDER=人類
-e ANTHROPIC_API_KEY=sk-ant-your-api-key
-e MODEL_NAME=claude-opus-4-5-20251101 # オプション
コピー
⁠ AWS Bedrock (デフォルト)
-e AWS_ACCESS_KEY_ID=your_key
-e AWS_SECRET_ACCESS_KEY=your_secret
-e AWS_REGION_NAME=us-east-1
-e MODEL_NAME=us.anthropic.claude-opus-4-5-20251101-v1:0 # オプション
コピー
⁠ オラマ (地元のモデル)
-e MODEL_PROVIDER=オラマ
-e OLLAMA_BASE_URL=http://host.docker.internal:11434
-e MODEL_NAME=deepseek-v3.1:671b-cloud # オプション
コピー
⁠ ムーンショット
-e MODEL_PROVIDER=ムーンショット
-e MOONSHOT_API_KEY=sk-あなたの API キー
-e MOONSHOT_API_BASE=https://api.moonshot.ai/v1 # オプション
-e MODEL_NAME=kimi-k2.5 # オプション
コピー
⁠ コマンドラインオプション
--target-branch : 比較するブランチ - デフォルト: main
--output : 出力ファイルのパスまたはディレクトリ
--repo-path : git リポジトリへのパス - デフォルト: /repo
--instructions : カスタム レビュー ガイドラインを含むマークダウン ファイルへのパス
--verify : 検証モードを有効にして誤検知を削減します (実験的)
--sast : OpenGrep SAST 事前スキャンを有効にする (実験的)
--json : レビューをマークダウンではなく JSON として出力します
⁠ 要件
/repo にマウントされた Git リポジトリ
Anthropic API キーまたは AWS Bedrock 認証情報のいずれか
出力ディレクトリは書き込み可能でなければなりません
⁠ リンク
ドキュメント: htt

ps://github.com/karill89/reviewcerberus ⁠
問題: https://github.com/karill89/reviewcerberus/issues ⁠
docker pull karill89/reviewcerberus:1.5.1 コピー理由 概要 コンテナ製品とは 製品概要 製品の内容 Docker デスクトップ Docker Hub の機能 コンテナ ランタイム 開発者ツール Docker アプリ Kubernetes 開発者 入門 Docker で遊ぶ コミュニティ オープンソース ドキュメント 会社概要 リソース ブログ お客様 パートナー ニュースルーム イベントとウェビナー 採用情報 お問い合わせ システム ステータス ⁠ © 2026 Docker, Inc. All Rights予約済み。 |利用規約 |サブスクリプションサービス契約 |プライバシー |法律上の

## Original Extract

Forums ⁠ Contact support System status ⁠ System theme Sign in Sign up Back
Forums ⁠ Contact support System status ⁠ System theme Docker Suite Explore
kirill89 / reviewcerberus repository overview
AI-powered code review tool that analyzes git branch differences and generates
comprehensive review reports with structured output.
⁠ Quick Start
docker run --rm -it -v $(pwd):/repo \
-e MODEL_PROVIDER=anthropic \
-e ANTHROPIC_API_KEY=sk-ant-your-api-key \
kirill89/reviewcerberus:latest \
--repo-path /repo --output /repo/review.md
Copy
That's it! The review will be saved to review.md in your current
directory.
⁠ Key Features
Comprehensive Reviews : Detailed analysis of logic, security, performance,
and code quality
Structured Output : Issues organized by severity with summary table
Multi-Provider : AWS Bedrock, Anthropic API, Ollama, or Moonshot
Smart Analysis : Context provided upfront with prompt caching
Git Integration : Works with any repository, supports commit hashes
Verification Mode : Experimental
Chain-of-Verification ⁠ to reduce false
positives
⁠ Usage Examples
Custom target branch:
docker run --rm -it -v $(pwd):/repo \
-e MODEL_PROVIDER=anthropic \
-e ANTHROPIC_API_KEY=sk-ant-your-api-key \
kirill89/reviewcerberus:latest \
--repo-path /repo --target-branch develop --output /repo/review.md
Copy
With custom review guidelines:
docker run --rm -it -v $(pwd):/repo \
-e MODEL_PROVIDER=anthropic \
-e ANTHROPIC_API_KEY=sk-ant-your-api-key \
kirill89/reviewcerberus:latest \
--repo-path /repo --instructions /repo/guidelines.md --output /repo/review.md
Copy
⁠ What's Included
⁠ Comprehensive Code Review
Detailed analysis covering:
Logic & Correctness: Bugs, edge cases, error handling
Security: OWASP issues, access control, input validation
Performance: N+1 queries, bottlenecks, scalability
Code Quality: Duplication, complexity, maintainability
Side Effects: Impact on other system parts
Testing: Coverage gaps, missing test cases
Documentation: Missing or outdated docs
⁠ Structured Output
Every review includes:
Summary : High-level overview of changes and risky areas
Issues Table : All issues at a glance with severity indicators (🔴 CRITICAL,
🟠 HIGH, 🟡 MEDIUM, 🟢 LOW)
Detailed Issues : Each issue with explanation, location, and suggested fix
⁠ Configuration
⁠ Anthropic API
-e MODEL_PROVIDER=anthropic
-e ANTHROPIC_API_KEY=sk-ant-your-api-key
-e MODEL_NAME=claude-opus-4-5-20251101 # optional
Copy
⁠ AWS Bedrock (default)
-e AWS_ACCESS_KEY_ID=your_key
-e AWS_SECRET_ACCESS_KEY=your_secret
-e AWS_REGION_NAME=us-east-1
-e MODEL_NAME=us.anthropic.claude-opus-4-5-20251101-v1:0 # optional
Copy
⁠ Ollama (local models)
-e MODEL_PROVIDER=ollama
-e OLLAMA_BASE_URL=http://host.docker.internal:11434
-e MODEL_NAME=deepseek-v3.1:671b-cloud # optional
Copy
⁠ Moonshot
-e MODEL_PROVIDER=moonshot
-e MOONSHOT_API_KEY=sk-your-api-key
-e MOONSHOT_API_BASE=https://api.moonshot.ai/v1 # optional
-e MODEL_NAME=kimi-k2.5 # optional
Copy
⁠ Command-Line Options
--target-branch : Branch to compare against - default: main
--output : Output file path or directory
--repo-path : Path to git repository - default: /repo
--instructions : Path to markdown file with custom review guidelines
--verify : Enable verification mode to reduce false positives (experimental)
--sast : Enable OpenGrep SAST pre-scan (experimental)
--json : Output review as JSON instead of markdown
⁠ Requirements
Git repository mounted to /repo
Either Anthropic API key or AWS Bedrock credentials
Output directory must be writable
⁠ Links
Documentation: https://github.com/kirill89/reviewcerberus ⁠
Issues: https://github.com/kirill89/reviewcerberus/issues ⁠
docker pull kirill89/reviewcerberus:1.5.1 Copy Why Overview What is a Container Products Product Overview Product Offerings Docker Desktop Docker Hub Features Container Runtime Developer Tools Docker App Kubernetes Developers Getting Started Play with Docker Community Open Source Documentation Company About Us Resources Blog Customers Partners Newsroom Events and Webinars Careers Contact Us System Status ⁠ © 2026 Docker, Inc. All rights reserved. | Terms of Service | Subscription Service Agreement | Privacy | Legal
