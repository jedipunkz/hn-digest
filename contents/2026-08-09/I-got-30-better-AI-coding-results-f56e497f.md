---
source: "https://github.com/akasula09/CodeSlimmer"
hn_url: "https://news.ycombinator.com/item?id=49236555"
title: "I got 30% better AI coding results"
article_title: "GitHub - akasula09/CodeSlimmer: Package codebases into AI-optimized XML context. · GitHub"
author: "kushthedevmmm"
captured_at: "2026-08-09T22:20:06Z"
capture_tool: "hn-digest"
hn_id: 49236555
score: 1
comments: 0
posted_at: "2026-08-09T22:01:47Z"
tags:
  - hacker-news
  - translated
---

# I got 30% better AI coding results

- HN: [49236555](https://news.ycombinator.com/item?id=49236555)
- Source: [github.com](https://github.com/akasula09/CodeSlimmer)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T22:01:47Z

## Translation

タイトル: AI コーディングの結果が 30% 向上しました
記事のタイトル: GitHub - akasula09/CodeSlimmer: コードベースを AI に最適化された XML コンテキストにパッケージ化します。 · GitHub
説明: コードベースを AI に最適化された XML コンテキストにパッケージ化します。 GitHub でアカウントを作成して、akasula09/CodeSlimmer の開発に貢献してください。

記事本文:
GitHub - akasula09/CodeSlimmer: コードベースを AI に最適化された XML コンテキストにパッケージ化します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アカスラ09
/
コードスリマー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット .github/ workflows .github/ workflows ライセンス ライセンス README.md README.md codelimmer.py codelimmer.py pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
安全に会話する

トークン推定とディレクトリ マッピングを使用して、ローカル コードベースまたはパブリック GitHub リポジトリをクリーンな LLM 対応コンテキストに組み込みます。
CodeSlimmer は、ローカル マシンに保存されているか、GitHub でホストされているかに関係なく、あらゆるコード リポジトリを、Claude 3.5 、Gemini 3.1 Pro/3.6 Flash 、GPT-4o 、DeepSeek R1 などの AI モデル専用に構造化された単一のテキスト ファイルにパックします。
🌐 リモート GitHub URL サポート: パブリック GitHub URL ( [https://github.com/user/repo](https://github.com/user/repo) ) を渡して、リポジトリを直接フェッチ、抽出、スリム化します。Git は必要ありません。
📋 クロスプラットフォームのクリップボード統合: 出力をシステムのクリップボード (Windows、macOS、または Linux) に即座に直接コピーします。
🛡️ 上書きの安全性: 署名ヘッダーを検査して、ローカル ファイルが誤って上書きされないようにします。
🌳 ASCII ディレクトリ ツリー: コンテキスト出力の最上部に視覚的なプロジェクト構造マップを構築します。
🏷️ LLM フレンドリーな XML 出力: プロンプトの精度を最大限に高めるために、コードを明確な <file path="..."> タグで囲みます。
📊 トークン推定ツール: コンテキスト サイズ (~4 文字/トークン) を推定し、主要なモデルと比較して容量を確認します。
📏 スマート ファイル サイズ: 物理サイズ ( B 、 KB 、 MB 、 GB ) を計算してフォーマットします。
オプション A: PyPI 経由でインストールする (推奨)
pip を使用して CodeSlimmer をインストールします。
pip インストール コードスリム
オプション B: ソースから実行
リポジトリのクローンを作成し、スクリプトを直接実行します。
git clone [https://github.com/akasula09/codeslimmer.git](https://github.com/akasula09/codeslimmer.git)
CDコードスリマー
🚀 使用方法
コードスリマー
ソースから直接実行する場合:
Pythonコードスリマー.py
2. インタラクティブなワークフロー
CodeSlimmer は 2 つの異なる入力モードをサポートしています。
ローカル フォルダーの相対パスまたは絶対パスを入力します。 CodeSlimmer はそのフォルダー内に ai_context.txt を作成し、オプションでクリップボードのコピーを提供します。例を以下に示します。
--- コード S

リマー v2.2 ---
ローカル フォルダー パスまたは GitHub URL を入力します: my_project_folder
📄 .gitignore からロードされたルール
---成功! ---
12 個のコード ファイルがパックされています。
最終的なコンテキスト サイズ: 18.45 KB
💾 「ai_context.txt」をフォルダーに直接追加しました!
📊 推定トークン数: ~4,612 トークン
--------------------------------------------------
どの AI プロバイダーをターゲットにしていますか? (ジェミニ / クロード / chatgpt / ディープシーク): ジェミニ
--------------------------------------------------
🎯 Gemini モデルに対して評価:
✅ Gemini 3.1 Pro: フィットします! (4,612 / 2,000,000 トークン — 0.2306% の容量)
✅ Gemini 3.6 フラッシュ: 適合します! (4,612 / 1,000,000 トークン — 0.4612% の容量)
✅ Gemini 3.5 Flash-Lite: 適合します! (4,612 / 1,000,000 トークン — 0.4612% の容量)
出力テキストもクリップボードに直接コピーしますか?
はいの場合は「y」を入力するか、Enter キーを押してスキップします: y
📋 完全なコンテキストをクリップボードに直接コピーしました! Ctrl+V を AI に入力する準備ができました。
モード 2: リモート GitHub リポジトリ URL
パブリック GitHub URL を貼り付けます。 CodeSlimmer は、一時ワークスペース内のリポジトリをフェッチしてスリム化し、トークンを推定し、乱雑なファイルを残さずに結果をクリップボードに直接コピーするように求めるプロンプトを表示します。
--- コードスリマー v2.2 ---
ローカル フォルダー パスまたは GitHub URL を入力します: [https://github.com/akasula09/codeslimmer](https://github.com/akasula09/codeslimmer)
🌐 GitHub からリポジトリ「akasula09/codeslimmer」を取得しています...
---成功! ---
5つのコードファイルがパックされています。
最終的なコンテキスト サイズ: 12.30 KB
📊 推定トークン数: ~3,075 トークン
--------------------------------------------------
どの AI プロバイダーをターゲットにしていますか? (ジェミニ / クロード / chatgpt / ディープシーク): クロード
--------------------------------------------------
🎯 クロードモデルに対して評価:
✅ クロード 3.5 ソネット / 俳句: ぴったりです! (3,075 / 200,000 トークン — 1.5375

%容量)
✅ クロード 3 オーパス: ぴったりです! (3,075 / 200,000 トークン — 1.5375% の容量)
出力テキストをクリップボードに直接コピーしますか?
はいの場合は「y」を入力するか、Enter キーを押してスキップします: y
📋 完全なコンテキストをクリップボードに直接コピーしました! Ctrl+V を AI に入力する準備ができました。
ℹ️ プライベート リポジトリ: プライベート リポジトリまたは存在しないリポジトリへの認証されていないリクエストは、次の出力を出力します。
❌ リポジトリが存在しないか、プライベートです。
CodeSlimmer は、次のクリーンな LLM に最適化された XML 形式を使用してプロンプト コンテキストを構築します。
<!-- CODESLIMMER v2.2 によって生成 - 手動で編集しないでください -->
<プロジェクト構造ツリー>
私のプロジェクトフォルダー/
§── src/
│ §──index.js
│ └── utils.js
§── .gitignore
━── package.json
</プロジェクト_構造_ツリー>
<ファイルパス="src/index.js">
console.log("Hello World");
</ファイル>
<ファイルパス="src/utils.js">
エクスポート関数 add(a, b) { return a + b; }
</ファイル>
⚖️ライセンス
PolyForm 非営利ライセンス 1.0.0 に基づいて配布されます。個人、教育、研究、企業内での使用は無料です。詳細については、「ライセンス」を参照してください。
コードベースを AI に最適化された XML コンテキストにパッケージ化します。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Package codebases into AI-optimized XML context. Contribute to akasula09/CodeSlimmer development by creating an account on GitHub.

GitHub - akasula09/CodeSlimmer: Package codebases into AI-optimized XML context. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
akasula09
/
CodeSlimmer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits .github/ workflows .github/ workflows LICENSE LICENSE README.md README.md codeslimmer.py codeslimmer.py pyproject.toml pyproject.toml View all files Repository files navigation
Safely convert local codebases or public GitHub repos into clean, LLM-ready context with token estimation and directory mapping.
CodeSlimmer packs any code repository—whether stored on your local machine or hosted on GitHub—into a single text file structured specifically for AI models like Claude 3.5 , Gemini 3.1 Pro/3.6 Flash , GPT-4o , and DeepSeek R1 .
🌐 Remote GitHub URL Support: Pass any public GitHub URL ( [https://github.com/user/repo](https://github.com/user/repo) ) to fetch, extract, and slim repositories directly—zero git required.
📋 Cross-Platform Clipboard Integration: Instantly copies output directly to your system clipboard (Windows, macOS, or Linux).
🛡️ Overwrite Safety: Inspects signature headers to ensure local files are never overwritten accidentally.
🌳 ASCII Directory Tree: Builds a visual project structure map at the very top of your context output.
🏷️ LLM-Friendly XML Output: Encloses code in clear <file path="..."> tags for max prompt accuracy.
📊 Token Estimator: Estimates context size (~4 chars/token) and checks capacity against major models.
📏 Smart File Sizes: Calculates and formats physical sizes ( B , KB , MB , GB ).
Option A: Install via PyPI (Recommended)
Install CodeSlimmer using pip :
pip install code-slimmer
Option B: Run from Source
Clone the repository and run the script directly:
git clone [https://github.com/akasula09/codeslimmer.git](https://github.com/akasula09/codeslimmer.git)
cd codeslimmer
🚀 How to Use
codeslimmer
If running directly from source:
python codeslimmer.py
2. Interactive Workflows
CodeSlimmer supports two distinct input modes:
Enter a relative or absolute local folder path. CodeSlimmer creates ai_context.txt inside that folder and offers an optional clipboard copy. An example is below.
--- CODE SLIMMER v2.2 ---
Enter a local folder path OR a GitHub URL: my_project_folder
📄 Loaded rules from .gitignore
--- SUCCESS! ---
Packed 12 code files.
Final Context Size: 18.45 KB
💾 Added 'ai_context.txt' directly into your folder!
📊 Estimated Token Count: ~4,612 tokens
--------------------------------------------------
Which AI provider are you targeting? (gemini / claude / chatgpt / deepseek): gemini
--------------------------------------------------
🎯 Evaluated against Gemini Models:
✅ Gemini 3.1 Pro: Fits! (4,612 / 2,000,000 tokens — 0.2306% capacity)
✅ Gemini 3.6 Flash: Fits! (4,612 / 1,000,000 tokens — 0.4612% capacity)
✅ Gemini 3.5 Flash-Lite: Fits! (4,612 / 1,000,000 tokens — 0.4612% capacity)
Do you want to copy the output text directly to your clipboard as well?
Type 'y' for Yes, or press Enter to skip: y
📋 Copied full context directly to your clipboard! Ready to Ctrl+V into AI.
Mode 2: Remote GitHub Repository URL
Paste any public GitHub URL. CodeSlimmer fetches and slims the repository in a temporary workspace, estimates tokens, and prompts you to copy the result straight to your clipboard without leaving messy files behind.
--- CODE SLIMMER v2.2 ---
Enter a local folder path OR a GitHub URL: [https://github.com/akasula09/codeslimmer](https://github.com/akasula09/codeslimmer)
🌐 Fetching repository 'akasula09/codeslimmer' from GitHub...
--- SUCCESS! ---
Packed 5 code files.
Final Context Size: 12.30 KB
📊 Estimated Token Count: ~3,075 tokens
--------------------------------------------------
Which AI provider are you targeting? (gemini / claude / chatgpt / deepseek): claude
--------------------------------------------------
🎯 Evaluated against Claude Models:
✅ Claude 3.5 Sonnet / Haiku: Fits! (3,075 / 200,000 tokens — 1.5375% capacity)
✅ Claude 3 Opus: Fits! (3,075 / 200,000 tokens — 1.5375% capacity)
Do you want to copy the output text directly to your clipboard?
Type 'y' for Yes, or press Enter to skip: y
📋 Copied full context directly to your clipboard! Ready to Ctrl+V into AI.
ℹ️ Private Repositories: Unauthenticated requests to private or non-existent repositories will output:
❌ Repository does not exist or is private.
CodeSlimmer structures the prompt context using this clean, LLM-optimized XML format:
<!-- GENERATED BY CODESLIMMER v2.2 - DO NOT EDIT MANUALLY -->
<project_structure_tree>
my_project_folder/
├── src/
│ ├── index.js
│ └── utils.js
├── .gitignore
└── package.json
</project_structure_tree>
<file path="src/index.js">
console.log("Hello World");
</file>
<file path="src/utils.js">
export function add(a, b) { return a + b; }
</file>
⚖️ License
Distributed under the PolyForm Noncommercial License 1.0.0 . Free for personal, educational, research, and internal business use. See LICENSE for details.
Package codebases into AI-optimized XML context.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
