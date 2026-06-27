---
source: "https://github.com/infiniteregrets/kv-psi"
hn_url: "https://news.ycombinator.com/item?id=48702538"
title: "Show HN: KV-psi, using Linux PSI to to trim an LLM KV cache"
article_title: "GitHub - infiniteregrets/kv-psi · GitHub"
author: "infiniteregrets"
captured_at: "2026-06-27T23:21:10Z"
capture_tool: "hn-digest"
hn_id: 48702538
score: 1
comments: 0
posted_at: "2026-06-27T22:50:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: KV-psi, using Linux PSI to to trim an LLM KV cache

- HN: [48702538](https://news.ycombinator.com/item?id=48702538)
- Source: [github.com](https://github.com/infiniteregrets/kv-psi)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T22:50:54Z

## Translation

タイトル: HN を表示: KV-psi、Linux PSI を使用して LLM KV キャッシュをトリミングする
記事タイトル: GitHub -finityregrets/kv-psi · GitHub
説明: GitHub でアカウントを作成して、infiniteregrets/kv-psi の開発に貢献します。
HN テキスト: LLM ランタイムに Linux PSI (Pressure Stall Information) を使用して KV キャッシュをトリミングするのは興味深いだろうと思いました。これは主に、統合メモリを備えた Jetson Orin スーパー ナノ キットのようなエッジ デバイスに役立ちます。私はあまりベンチを試していませんが、時間をかけてベンチを増やし、ローカル LLM を実行する際に実際に活用できるかどうかを確認する予定です。それが意味があるかどうか教えてください:P (もちろん、私はこのアイデアに共感しました)

記事本文:
GitHub - 無限後悔/kv-psi · GitHub
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
無限の後悔
/
kv-psi
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダー

とファイル
1 コミット 1 コミット ベンチマーク ベンチマークの例 例 統合/ llama.cpp 統合/ llama.cpp スクリプト スクリプト src/ psi_kv_governor src/ psi_kv_governor テスト テスト .gitignore .gitignore README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
PSI KV ガバナーは、Linux Pressure を使用するための小規模なリファレンス実装です。
システムがメモリ不足の場合に LLM KV キャッシュをトリミングするためのストール情報
プレッシャー。
PSI が有効になっている Linux: cgroupmemory.pressure または /proc/pressure/memory
llama.cpp はランナーの依存関係を構築します
GGUF モデル (例: models/SmolLM2-135M-Instruct-Q2_K.gguf)
猫/proc/圧力/メモリ
PYTHONPATH=src python benchmarks/pressure_bench.py --preflight-only
基本的な使い方
PYTHONPATH=src python -m psi_kv_governor.cli シミュレート
llama.cpp ランナーをビルドします。
scripts/build_llama_runner.sh
必要に応じて、小さなベンチマーク モデルをダウンロードします。
Python スクリプト/download_demo_model.py
PSI ベンチマーク
両方のバリアント注文を実行します。 PSI avg10 、キャッシュ、および zram/swap が重要であるため、これは重要です。
この状態は、最初の圧力実行から 2 番目の圧力実行まで引き継がれる可能性があります。
PYTHONPATH=src python benchmarks/pressure_bench.py \
-c 2048 \
-n 1536 \
-- 64 \ を維持します
--尾 256 \
--min-プルーン 64 \
--pressure-mib 6000 \
--圧力ステップ-mib 1024 \
--圧力ウォームアップ-s 10 \
--variant-cooldown-s 45 \
--out-dir data/bench-pressure/fixed-first
PYTHONPATH=src python benchmarks/pressure_bench.py \
--variant-order psi-first \
-c 2048 \
-n 1536 \
-- 64 \ を維持します
--尾 256 \
--min-プルーン 64 \
--pressure-mib 6000 \
--圧力ステップ-mib 1024 \
--圧力ウォームアップ-s 10 \
--variant-cooldown-s 45 \
--out-dir データ/ベンチ圧力/psi-first
最近の Jetson の結果:
データ/ベンチ圧力/実際の psi-6000m-1536tok-cooldown
データ/ベンチ圧力/real-psi-6000m-1536tok-cooldown-psi-first
エラーが発生しました

ロード中。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to infiniteregrets/kv-psi development by creating an account on GitHub.

I thought it'd be interesting to use Linux PSI (Pressure Stall Information) for an LLM runtime to trim the KV cache. This is mainly useful imo for edge devices like the Jetson Orin super nano kit which have unified memory. I haven't benched much, but plan to do so more over time and see if I can make a real use of it as I run local LLMs. Let me know if it makes sense :P (I of course vibed this idea)

GitHub - infiniteregrets/kv-psi · GitHub
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
infiniteregrets
/
kv-psi
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit benchmarks benchmarks examples examples integrations/ llama.cpp integrations/ llama.cpp scripts scripts src/ psi_kv_governor src/ psi_kv_governor tests tests .gitignore .gitignore README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
PSI KV Governor is a small reference implementation for using Linux Pressure
Stall Information to trim an LLM KV cache when the system is under memory
pressure.
Linux with PSI enabled: cgroup memory.pressure or /proc/pressure/memory
llama.cpp build dependencies for the runner
a GGUF model, for example models/SmolLM2-135M-Instruct-Q2_K.gguf
cat /proc/pressure/memory
PYTHONPATH=src python benchmarks/pressure_bench.py --preflight-only
Basic Usage
PYTHONPATH=src python -m psi_kv_governor.cli simulate
Build the llama.cpp runner:
scripts/build_llama_runner.sh
Download the small benchmark model if needed:
python scripts/download_demo_model.py
PSI Benchmark
Run both variant orders. This matters because PSI avg10 , cache, and zram/swap
state can carry over from the first pressure run into the second.
PYTHONPATH=src python benchmarks/pressure_bench.py \
-c 2048 \
-n 1536 \
--keep 64 \
--tail 256 \
--min-prune 64 \
--pressure-mib 6000 \
--pressure-step-mib 1024 \
--pressure-warmup-s 10 \
--variant-cooldown-s 45 \
--out-dir data/bench-pressure/fixed-first
PYTHONPATH=src python benchmarks/pressure_bench.py \
--variant-order psi-first \
-c 2048 \
-n 1536 \
--keep 64 \
--tail 256 \
--min-prune 64 \
--pressure-mib 6000 \
--pressure-step-mib 1024 \
--pressure-warmup-s 10 \
--variant-cooldown-s 45 \
--out-dir data/bench-pressure/psi-first
Recent Jetson result:
data/bench-pressure/real-psi-6000m-1536tok-cooldown
data/bench-pressure/real-psi-6000m-1536tok-cooldown-psi-first
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
