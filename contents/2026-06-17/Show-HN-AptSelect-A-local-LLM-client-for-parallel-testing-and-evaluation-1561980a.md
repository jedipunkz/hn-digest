---
source: "https://aptselect.com"
hn_url: "https://news.ycombinator.com/item?id=48570680"
title: "Show HN: AptSelect – A local LLM client for parallel testing and evaluation"
article_title: "Test LLMs Side-by-Side"
author: "dhavalt"
captured_at: "2026-06-17T16:23:31Z"
capture_tool: "hn-digest"
hn_id: 48570680
score: 3
comments: 0
posted_at: "2026-06-17T14:00:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AptSelect – A local LLM client for parallel testing and evaluation

- HN: [48570680](https://news.ycombinator.com/item?id=48570680)
- Source: [aptselect.com](https://aptselect.com)
- Score: 3
- Comments: 0
- Posted: 2026-06-17T14:00:13Z

## Translation

タイトル: Show HN: AptSelect – 並行テストと評価のためのローカル LLM クライアント
記事のタイトル: LLM を並べてテストする
説明: LLM をテストするためのローカルファーストのデスクトップ ワークベンチ。 GPT、Claude、Mistral、Gemini、Grok を並べて実行します。
HN テキスト: さまざまな LLM が特定の命令をどのように処理し、エッジ ケースをプロンプトするかをテストする必要があるたびに、使い捨てスクリプトを作成するのをやめるため、AptSelect を構築しました。機能: 並列実行: 単一のプロンプトを OpenAI、Anthropic、Mistral、および Gemini に同時に送信します。出力、レイテンシー、正確なトークン使用量を並べて比較します。バッチ評価: CSV データセットをアップロードして、複数のモデルにわたる一括テストを一度に実行します。手動診断: 出力を手動で評価 (合格/不合格) し、診断タグ (幻覚、フォーマット エラーなど) を割り当てて、人間が検証したパフォーマンス リーダーボードを構築します。ローカルファースト: API キーは OS キーリングで暗号化されます。ローカル SQLite DB に保存される履歴。テレメトリーはありません。技術的なフィードバックを求めています。現在の LLM テスト/評価ツールで最も間違っている点は何だと思いますか?

記事本文:
_ ">
aptselect
変更履歴
ロードマップ
マニュアル
購入する
🌓-->
-->
繰り返します。比較してください。ベンチマーク。
主要な LLM 全体でプロンプトをテスト、評価、ベンチマークするように設計されたローカルファーストのデスクトップ クライアント。推測をやめてください
モデルがどのように動作し、データセットに対してそれを証明するか。
私たちは、迅速なエンジニアリングという厄介な現実に対処するために、ローカルファーストの環境を構築しています。魔法のボタンはありませんが、
テスト、反復、ベンチマークのための基本的なツールにすぎません。
単一のプロンプト テンプレートを GPT-4、Claude 3、および Gemini に同時に送信します。生のデータを瞬時に比較
複数のデータを管理することなく、JSON 出力、レイテンシ メトリクス、正確なトークン消費量を並べて表示
ブラウザのタブ。
API キーとプロンプト履歴は、ローカルの SQLite データベースに保存されます。私たちには何も触れません
サーバー。
自動プロンプトチェックポイント
すべての反復はローカル データベースに自動的に保存されます。プロンプトをフォークして新しい変数をテストします。
出力を改善した正確な変更を追跡し、過去の構成に簡単に戻すことができます。
テスト データをプロンプト テンプレートに挿入して、ベースラインを確立します。新しい LLM が下落したときのベンチマーク
運用環境で信頼する前に、履歴データと照合してください。
複数のモデルにわたる完全なテスト データセットに対してプロンプトを一度に実行します。バッチ出力を確認する
並べて合格/不合格のグレードを割り当てて、どのモデルがエッジ ケースに対応しているかを正確に確認します。
プロンプトのバージョン管理。
反復の履歴をきれいに保管してください。プロンプトをフォークして新しい変数をテストし、変更を追跡し、
過去のバージョンに簡単に戻すこともできます。
チャット インターフェイスでは詳細が隠されます。生の API レスポンス、レイテンシー統計、正確なトークン使用状況を検査します。
あらゆるリクエスト。
複数のモデルにわたる完全なテスト データセットに対してプロンプトを一度に実行します。バッチ出力を確認する
並べて合格/不合格のグレードを割り当てて、どのモデルが適切かを正確に確認します

エッジケースに対応します。
資格情報をマシン上に保管してください。キーは OS キーリングを介して暗号化され、
ローカルデータベースに保存され、プロバイダーに厳密に送信されます。私たちは何も追跡しません。
API の接続、アクティブなモデルの切り替え、リクエスト パラメーターの調整をすべて 1 つの中央インターフェイスから行います。
鍵はご自身でご用意ください。接続する
OpenAI、Anthropic、Mistral、Gemini、XAI を数秒で実行できます。モデルのオン/オフを切り替えて、
作業スペースがきれい。
温度、top_p、周波数ペナルティを調整して、さまざまな制約がどのような影響を与えるかを観察します。
すぐに結果が出ます。
クロスプラットフォームのサポートのために Electron を選択しましたが、スタックはそのままにしました。
できるだけシンプルに。
大きなフレームワークのオーバーヘッドはありません。標準の HTML、CSS、バニラを使用してインターフェイスを構築しました
JavaScript。
データはディスク上の標準 SQLite ファイルに保存されます。バックアップ、バージョン管理、または削除
いつでも好きなときに。
{
"ランタイム" : "電子" ,
"セキュリティ" : "コンテキスト分離" ,
"フロントエンド" : "バニラ JS + Web コンポーネント" ,
"データベース" : "SQLite3 (ローカルのみ)"
}
ライブデモ
同時モデルテスト。
Anthropic、Mistral、OpenAI、xAI を並べて評価します。ブラウザーのタブや手動のコピー＆ペーストは必要ありません。
AI 機能を確実に出荷します。
本番環境に移行する前に、データセットをバッチテストしてモデルの信頼性を証明します。
Mac 用ダウンロード
Windows 用のダウンロード
Linux 用のダウンロード
永久ライセンスと早期導入価格については、こちらをご覧ください。
ライセンスの詳細を表示 →
このサイトでは、 aptselect の成長を支援するために分析を使用しています。私たちは内部であなたのデータを決して追跡しません
デスクトップアプリ。続きを読む 。

## Original Extract

A local-first desktop workbench to test LLMs. Run GPT, Claude, Mistral, Gemini, and Grok side-by-side.

I built AptSelect to stop writing throwaway scripts every time I needed to test how different LLMs handle specific instructions and prompt edge cases. What it does: Parallel Execution: Send a single prompt to OpenAI, Anthropic, Mistral, and Gemini simultaneously. Compare the outputs, latency, and exact token usage side-by-side. Batch Evaluations: Upload a CSV dataset to run bulk tests across multiple models at once. Manual Diagnostics: Grade outputs manually (Pass/Fail) and assign diagnostic tags (e.g., Hallucination, Format Error) to build a human-verified performance leaderboard. Local-first: API keys encrypted with your OS keyring; history stored in a local SQLite DB; no telemetry. I’m looking for technical feedback. What do you think current LLM testing/evaluation tools get most wrong?

_ ">
aptselect
Changelog
Roadmap
Manual
Buy
🌓-->
-->
Iterate. Compare. Benchmark.
A local-first desktop client designed to test, grade, and benchmark prompts across major LLMs. Stop guessing
how a model will perform and prove it against your datasets.
We're building a local-first environment for the messy reality of prompt engineering. No magic buttons,
just the foundational tools to test, iterate, and benchmark.
Send a single prompt template to GPT-4, Claude 3, and Gemini simultaneously. Instantly compare raw
JSON outputs, latency metrics, and exact token consumption side-by-side without managing multiple
browser tabs.
Your API keys and prompt history are stored in a local SQLite database. Nothing touches our
servers.
Automatic Prompt Checkpointing
Every iteration is automatically saved to your local database. Fork a prompt to test a new variable,
track the exact changes that improved the output, and easily revert to past configurations.
Inject test data into your prompt templates to establish a baseline. When a new LLM drops, benchmark
it against your historical data before trusting it in production.
Run your prompt against a full test dataset across multiple models at once. Review the batch outputs
side-by-side and assign pass/fail grades to see exactly which model handles your edge cases.
Version Control for Your Prompts.
Keep a clean history of your iterations. Fork a prompt to test a new variable, track the changes,
and easily switch back to past versions.
Chat interfaces hide the details. Inspect raw API responses, latency stats, and exact token usage for
every single request.
Run your prompt against a full test dataset across multiple models at once. Review the batch outputs
side-by-side and assign pass/fail grades to see exactly which model handles your edge cases.
Keep your credentials on your machine. Your keys are encrypted via your OS keyring, saved to your
local database, and sent strictly to the providers. We track nothing.
Connect your APIs, toggle active models, and adjust request parameters all from one central interface.
Bring your own keys. Connect
OpenAI, Anthropic, Mistral, Gemini and XAI in seconds. Toggle models on/off to keep your
workspace clean.
Adjust temperature, top_p, and frequency penalties to observe how different constraints impact
your prompt results.
We chose Electron for cross-platform support, but kept the stack as
simple as possible.
No heavy frameworks overhead. We built the interface using standard HTML, CSS, and vanilla
JavaScript.
Your data lives in a standard SQLite file on your disk. Backup, version control, or delete it
whenever you want.
{
"runtime" : "Electron" ,
"security" : "Context Isolated" ,
"frontend" : "Vanilla JS + Web Components" ,
"database" : "SQLite3 (Local-only)"
}
LIVE DEMO
Concurrent Model Testing.
Evaluating Anthropic, Mistral, OpenAI, and xAI side-by-side. No browser tabs, no manual copy-pasting.
Ship AI Features With Certainty.
Batch-test your datasets and prove model reliability before hitting production.
Download for Mac
Download for Windows
Download for Linux
Learn about our permanent licensing and early-adopter pricing:
View License Details →
We use analytics on this site to help us grow aptselect . We never track your data inside the
desktop app. Read more .
