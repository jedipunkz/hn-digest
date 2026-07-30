---
source: "https://github.com/grafana/ai-sdk"
hn_url: "https://news.ycombinator.com/item?id=49108778"
title: "Go LLM SDK for streaming, tool-calling AI backends (plus frontend React lib)"
article_title: "GitHub - grafana/ai-sdk: Grafana AI SDK for Go — streaming, tool-calling AI backends that speak fluent @ai-sdk/react · GitHub"
author: "matryer"
captured_at: "2026-07-30T12:23:58Z"
capture_tool: "hn-digest"
hn_id: 49108778
score: 9
comments: 0
posted_at: "2026-07-30T11:55:39Z"
tags:
  - hacker-news
  - translated
---

# Go LLM SDK for streaming, tool-calling AI backends (plus frontend React lib)

- HN: [49108778](https://news.ycombinator.com/item?id=49108778)
- Source: [github.com](https://github.com/grafana/ai-sdk)
- Score: 9
- Comments: 0
- Posted: 2026-07-30T11:55:39Z

## Translation

タイトル: ストリーミング、ツール呼び出し AI バックエンド用の Go LLM SDK (およびフロントエンド React lib)
記事タイトル: GitHub - grafana/ai-sdk: Grafana AI SDK for Go — 流暢に話すストリーミング、ツール呼び出し AI バックエンド @ai-sdk/react · GitHub
説明: Go 用 Grafana AI SDK — 流暢に話すストリーミング、ツール呼び出し AI バックエンド @ai-sdk/react - grafana/ai-sdk

記事本文:
GitHub - grafana/ai-sdk: Go 用 Grafana AI SDK — 流暢に話すストリーミング、ツール呼び出し AI バックエンド @ai-sdk/react · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
グラファナ
/
ある

i-SDK
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .agents/ スキル .agents/ スキル .claude/ スキル .claude/ スキル .github .github ドキュメント ドキュメントの例 例 フォールバック フォールバック ゲートウェイ ゲートウェイ 内部 内部ミドルウェア ミドルウェア openspec openspec 出力 出力プロバイダ プロバイダ プロバイダ プロバイダ レジストリ レジストリ スキーマ スキーマ スクリプト スクリプト テスト test .gitignore .gitignore .golangci.yml .golangci.yml .markdownlint-cli2.jsonc .markdownlint-cli2.jsonc AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md エージェント.ゴー エージェント.ゴー エージェント_テスト.ゴー エージェント_テスト.ゴー チャンク.ゴー チャンク.ゴー チャンク_テスト.ゴー チャンク_テスト.ゴー 変換.ゴー 変換.GO Convert_test.GO Convert_test.GO doc.go doc.go extract_content.go extract_content.go extract_content_test.go extract_content_test.gogeneratetext.go Generatetext.go go.mod go.mod go.sum go.sum go.work go.work go.work.sum go.work.sum http.go http.go http_test.go http_test.go id.go id.go id_test.go id_test.go integration_test.go integration_test.go message.go message.go message_json.go message_json.go message_json_test.go message_json_test.go misse.toml misse.toml missing_tool_results_error.go missing_tool_results_error.go options.go options.go options_test.go options_test.go Output.go Output.go renovate.json renovate.json retry.go retry.go retry_error.go retry_error.go retry_integration_test.go retry_integration_test.go retry_test.go retry_test.go stream.go stream.go stream_test.go stream_test.go streamtext.go streamtext.go streamtext_output_test.go streamtext_output_test.go streamtext_test.go streamtext_test.go te

xt.go text.go textstream.go textstream.go to_response_messages.go to_response_messages.go to_response_messages_test.go to_response_messages_test.go tools.go tools.go tool_approval_errors.go tool_approval_errors.go tool_approval_signature.go tool_approval_signature.go totool_approval_signature_test.goツール_承認_署名_テスト. ツール_フィンガープリント. ツール_フィンガープリント. ツール_フィンガープリント_テスト。
[切り捨てられた]
言語モデルの呼び出し、応答のストリーミング、ツールの実行、AI を活用したサービスの提供
Go のエンドポイント。 SDK を単独で使用するか、AI SDK React と組み合わせて使用します
フロントエンド。
クイック スタート · ドキュメント · 例 · API リファレンス
SDK は、モデル呼び出し、ストリーミング、ツールなどのための 1 つの API を Go アプリケーションに提供します。
構造化された出力と、サポートされているプロバイダーにわたるマルチステップ エージェント。続いて
Vercel の AI SDK の設計と有線互換性の維持
TypeScript フロントエンド フックを使用します。 Go エンドポイントはサーバー送信イベントをストリーミングできます
(SSE) useChat などのフックに直接接続します。
Go バックエンド React フロントエンド
────────────
aisdk.StreamText(...) ── SSE ──▶ useChat({ トランスポート })
aisdk.WriteUIMessageStream(w, …) // 同じプロトコル
世代に対してリクエストがどのように実行されるかを参照してください。
ツールとストリーミング フロー。既存の AI SDK React フロントエンドを再利用するか、
プロトコル アダプターを追加せずに Go を使用した TypeScript バックエンド。
StreamText / GenerateText — 応答をストリーミングするか、完了を待ちます
結果、再試行と複数ステップのツール実行あり
React の互換性 — useChat 、 useCompletion 、および useObject を提供します
コンポーズ可能なツール — モデルからプレーンな Go 関数を呼び出し、
結果的なアクションの承認
構造化された出力 — スキーマ検証されたオブジェクト、配列、選択肢を生成します
複数のプロバイダー - Anthropic、Amazon Bedrock、OpenAI、
オープンA

I 互換 API、および内部サービスからの Grafana のホスト型エンドポイント
運用管理 - タイムアウト、フォールバック、ロギング、Prometheus を構成します。
メトリクスとエージェントの可観測性
Go プロジェクトを作成し、コア モジュールと 1 つのプロバイダーをインストールします。
mkdir ai-sdk-クイックスタート
cd ai-sdk-クイックスタート
go mod init example.com/ai-sdk-quickstart
github.com/grafana/ai-sdk を入手してください
github.com/grafana/ai-sdk/providers/anthropic を取得してください
「Amazon Bedrock、OpenAI、
OpenAI 互換 API、および内部的にプロビジョニングされた Grafana ホスト型エンドポイント。
この完全なプログラムを main.go として保存します。 1 つのモデル呼び出しを実行し、
応答:
パッケージメイン
インポート(
「コンテキスト」
「fmt」
「ログ」
「オス」
aisdk "github.com/grafana/ai-sdk"
「github.com/grafana/ai-sdk/provider」
「github.com/grafana/ai-sdk/providers/anthropic」
）
関数メイン () {
apiKey := os 。 Getenv ( "ANTHROPIC_API_KEY" )
if apiKey == "" {
ログ。致命的 (「ANTHROPIC_API_KEY が必要です」)
}
モデル:= 人間。新規 ( apiKey 、 "claude-sonnet-5" )
結果、エラー:= aisdk。 GenerateText ( context . Background ()、model 、
aisdk 。 WithModelMessages ( Provider . UserText ( "ゴルーチンを一文で説明します。" ))、
）
エラーの場合 != nil {
ログ。致命的 (エラー)
}
fmt 。 Println (結果.テキスト)
}
Anthropic API キーを使用して実行します。
ANTHROPIC_API_KEY=sk-... 走ってください。
プロジェクトの初期化と資格情報のガイダンスについては、次のとおりです。
インストール。この応答をストリーミングするには
React クライアントの場合は、「フルスタック チャットの構築」に進みます。
目標
ここから始めましょう
Go からモデル呼び出しを行う
Goからテキストを生成する
React チャットを構築する
フルスタックチャット
型指定されたデータを返す
構造化された出力
モデルに Go コードを呼び出してもらう
ツール
再利用可能なエージェントを構築する
エージェントループ
モデルプロバイダーを選択してください
プロバイダーの概要
ロギングまたは可観測性を追加する
ミドルウェアの概要
生産の準備をする
製作チェックリスト
完全なインデックス

: ドキュメント · 実行可能なコード: 例 · 正確な API: pkg.go.dev
貢献は大歓迎です。 CONTRIBUTING.md では、
開発セットアップ、このリポジトリを珍しいものにする 2 つの規則 —
Vercel AI SDK とのアップストリームの同等性、および仕様主導型開発
OpenSpec — およびプルリクエストのチェックリスト。参加者全員がフォローしています
行動規範。
Apache ライセンス 2.0 。この SDK は、次の設計に従っています。
Vercel の AI SDK も Apache-2.0 ライセンスを取得しています。
帰属は NOTICE に記録されます。
Go 用 Grafana AI SDK — 流暢に話すストリーミング、ツール呼び出し AI バックエンド @ai-sdk/react
Readme Apache-2.0 ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Grafana AI SDK for Go — streaming, tool-calling AI backends that speak fluent @ai-sdk/react - grafana/ai-sdk

GitHub - grafana/ai-sdk: Grafana AI SDK for Go — streaming, tool-calling AI backends that speak fluent @ai-sdk/react · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Uh oh!
There was an error while loading. Please reload this page .
grafana
/
ai-sdk
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .agents/ skills .agents/ skills .claude/ skills .claude/ skills .github .github docs docs examples examples fallback fallback gateway gateway internal internal middleware middleware openspec openspec output output provider provider providers providers registry registry schema schema scripts scripts test test .gitignore .gitignore .golangci.yml .golangci.yml .markdownlint-cli2.jsonc .markdownlint-cli2.jsonc AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md agent.go agent.go agent_test.go agent_test.go chunk.go chunk.go chunk_test.go chunk_test.go convert.go convert.go convert_test.go convert_test.go doc.go doc.go extract_content.go extract_content.go extract_content_test.go extract_content_test.go generatetext.go generatetext.go go.mod go.mod go.sum go.sum go.work go.work go.work.sum go.work.sum http.go http.go http_test.go http_test.go id.go id.go id_test.go id_test.go integration_test.go integration_test.go message.go message.go message_json.go message_json.go message_json_test.go message_json_test.go mise.toml mise.toml missing_tool_results_error.go missing_tool_results_error.go options.go options.go options_test.go options_test.go output.go output.go renovate.json renovate.json retry.go retry.go retry_error.go retry_error.go retry_integration_test.go retry_integration_test.go retry_test.go retry_test.go stream.go stream.go stream_test.go stream_test.go streamtext.go streamtext.go streamtext_output_test.go streamtext_output_test.go streamtext_test.go streamtext_test.go text.go text.go textstream.go textstream.go to_response_messages.go to_response_messages.go to_response_messages_test.go to_response_messages_test.go tool.go tool.go tool_approval_errors.go tool_approval_errors.go tool_approval_signature.go tool_approval_signature.go tool_approval_signature_test.go tool_approval_signature_test.go tool_fingerprint.go tool_fingerprint.go tool_fingerprint_test.
[truncated]
Call language models, stream responses, execute tools, and serve AI-powered
endpoints from Go. Use the SDK on its own or pair it with an AI SDK React
frontend.
Quick start · Documentation · Examples · API reference
The SDK gives Go applications one API for model calls, streaming, tools,
structured output, and multi-step agents across supported providers. It follows
the design of Vercel's AI SDK and stays wire-compatible
with its TypeScript frontend hooks. A Go endpoint can stream Server-Sent Events
(SSE) directly to hooks such as useChat .
Go backend React frontend
────────── ──────────────
aisdk.StreamText(...) ── SSE ──▶ useChat({ transport })
aisdk.WriteUIMessageStream(w, …) // same protocol
See How a request runs for the generation,
tool, and streaming flow. Reuse an existing AI SDK React frontend or replace a
TypeScript backend with Go without adding a protocol adapter.
StreamText / GenerateText — stream a response or wait for the complete
result, with retries and multi-step tool execution
React compatibility — serve useChat , useCompletion , and useObject
Composable tools — call plain Go functions from a model and require
approval for consequential actions
Structured output — generate schema-validated objects, arrays, and choices
Multiple providers — call Anthropic, Amazon Bedrock, OpenAI,
OpenAI-compatible APIs, and Grafana's hosted endpoint from internal services
Production controls — configure timeouts, fallback, logging, Prometheus
metrics, and Agent Observability
Create a Go project and install the core module and one provider:
mkdir ai-sdk-quickstart
cd ai-sdk-quickstart
go mod init example.com/ai-sdk-quickstart
go get github.com/grafana/ai-sdk
go get github.com/grafana/ai-sdk/providers/anthropic
See Choose a provider for Amazon Bedrock, OpenAI,
OpenAI-compatible APIs, and the internally provisioned Grafana hosted endpoint.
Save this complete program as main.go . It makes one model call and prints the
response:
package main
import (
"context"
"fmt"
"log"
"os"
aisdk "github.com/grafana/ai-sdk"
"github.com/grafana/ai-sdk/provider"
"github.com/grafana/ai-sdk/providers/anthropic"
)
func main () {
apiKey := os . Getenv ( "ANTHROPIC_API_KEY" )
if apiKey == "" {
log . Fatal ( "ANTHROPIC_API_KEY is required" )
}
model := anthropic . New ( apiKey , "claude-sonnet-5" )
result , err := aisdk . GenerateText ( context . Background (), model ,
aisdk . WithModelMessages ( provider . UserText ( "Explain goroutines in one sentence." )),
)
if err != nil {
log . Fatal ( err )
}
fmt . Println ( result . Text )
}
Run it with an Anthropic API key:
ANTHROPIC_API_KEY=sk-... go run .
For project initialization and credential guidance, follow
Installation . To stream this response to
a React client, continue with Build a full-stack chat .
Goal
Start here
Make model calls from Go
Generate text from Go
Build a React chat
Full-stack chat
Return typed data
Structured output
Let a model call Go code
Tools
Build a reusable agent
Agent loops
Choose a model provider
Provider overview
Add logging or observability
Middleware overview
Prepare for production
Production checklist
Full index: Documentation · Runnable code: Examples · Exact APIs: pkg.go.dev
Contributions are welcome. CONTRIBUTING.md covers the
development setup, the two conventions that make this repository unusual —
upstream parity with the Vercel AI SDK, and spec-driven development with
OpenSpec — and the pull request checklist. All participants follow our
Code of Conduct .
Apache License 2.0 . This SDK follows the design of
Vercel's AI SDK , also Apache-2.0 licensed;
attribution is recorded in NOTICE .
Grafana AI SDK for Go — streaming, tool-calling AI backends that speak fluent @ai-sdk/react
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
