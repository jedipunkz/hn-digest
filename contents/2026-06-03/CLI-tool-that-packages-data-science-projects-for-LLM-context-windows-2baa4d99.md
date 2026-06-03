---
source: "https://github.com/arianmokhtariha/data2prompt"
hn_url: "https://news.ycombinator.com/item?id=48373045"
title: "CLI tool that packages data science projects for LLM context windows"
article_title: "GitHub - arianmokhtariha/data2prompt: A custom made repository-to-prompt packager made specifically for data analyst/scientist projects. · GitHub"
author: "ArianM"
captured_at: "2026-06-03T00:40:06Z"
capture_tool: "hn-digest"
hn_id: 48373045
score: 14
comments: 0
posted_at: "2026-06-02T17:08:25Z"
tags:
  - hacker-news
  - translated
---

# CLI tool that packages data science projects for LLM context windows

- HN: [48373045](https://news.ycombinator.com/item?id=48373045)
- Source: [github.com](https://github.com/arianmokhtariha/data2prompt)
- Score: 14
- Comments: 0
- Posted: 2026-06-02T17:08:25Z

## Translation

タイトル: LLM コンテキスト ウィンドウ用のデータ サイエンス プロジェクトをパッケージ化する CLI ツール
記事のタイトル: GitHub - arianmokhtariha/data2prompt: データ アナリスト/サイエンティスト プロジェクト専用に作られた、カスタムメイドのリポジトリからプロンプトへのパッケージャーです。 · GitHub
説明: データ アナリスト/サイエンティスト プロジェクト向けに特別に作成された、カスタムメイドのリポジトリからプロンプトへのパッケージャです。 - アリアンモクタリハ/data2prompt

記事本文:
GitHub - arianmokhtariha/data2prompt: データ アナリスト/科学者のプロジェクト向けに特別に作成された、カスタムメイドのリポジトリからプロンプトへのパッケージャです。 · GitHub
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
アリアンモクタリハ
/
データ2プロンプト
公共
通知
通知を変更するにはサインインする必要があります

n設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
58 コミット 58 コミット .github/ workflows .github/ workflows 資産 アセット ドキュメント ドキュメント プラン プラン src/ data2prompt src/ data2prompt テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
データ サイエンス ワークフローおよびデータ量の多いプロジェクト向けの、高パフォーマンスのコードベースからプロンプトへのオーケストレーション。
data2prompt は、ローカルのデータ量の多いプロジェクトとラージ言語モデル (LLM) コンテキスト ウィンドウの間のギャップを埋めるために設計された CLI ツールです。汎用のコード パッケージャーとは異なり、LLM アテンション メカニズムのインテリジェントで最適化された出力、プロジェクトの構造とコンテンツのトークン対応表現を提供します。
Data2prompt は、大規模な純粋なコード リポジトリではなく、データ量の多いプロジェクト ( .csv 、 .sql 、 .xlsx 、 .ipynb ) 専用に構築されています。データ ファイルをインテリジェントにサンプリングして切り詰め、セマンティック構造を維持しながらコンテキスト ウィンドウの爆発を防ぎます。
一般的なコードからプロンプトへのツールは、データ ファイルを完全にスキップするか、コンテキスト ウィンドウの 90% を無駄にする生の CSV をダンプして、データ ファイルを停止させます。 Data2Prompt は、データ サイエンス ワークフロー向けに特別に設計されたインテリジェントなサンプリング、スキーマ抽出、LLM に最適化された書式設定によってこの問題を解決します。
スマート Jupyter 解析 : 重い Base64 画像と生の HTML を削除してコンテキストを保持しながら、.ipynb ファイルからコード、マークダウン、およびテキスト出力をインテリジェントに抽出します。
マルチフォーマット サンプリング : CSV、SQL、Excel ファイルの高度なサンプリング戦略により、スキーマとデータ コンテキストを保持し、llm に必要なコンテキストを抽出しながらデータ サイズを大幅に削減します。
積極的な切り詰め: コンテキストを維持するために、長い行は切り詰められ、行の挿入を無効にします。

表形式のデータがサンプリング後にまだ大きすぎる場合は、一定量まで切り詰められます。また、ハンドルされていないタイプの生のテキスト ファイルが大きすぎる場合も、一定量まで切り詰められます。
防御処理: 自動バイナリ検出 (Null バイト チェック)。最初の 1024 バイトで Null バイトを検索することによって、ファイルがバイナリかどうかをチェックします。
最適化された LLM への注目 : デフォルトの出力形式は、適切に構造化されたスキーマを使用したマークダウンであり、別のオプションは、複雑な分析と大きなコンテキスト ウィンドウ向けに LLM アンカーを強化するための XML スタイル タグを使用した XML 出力です。
トークン対応出力: tiktoken ( o200k_base ) を使用したリアルタイムのトークン推定により、プロンプトがターゲット LLM (Claude 3.5、GPT-4o、Gemini 1.5) に適合することを確認し、 regex による高度なオフライン トークン カウントを実行します。
Professional TUI : Rich で構築された忠実度の高いターミナル インターフェイスで、Windows 上でマトリックス スタイルの起動アニメーションと対話型でスクロール可能なレポートを備えています。
動的マークダウン ラッピング : インテリジェントなバックティックの深さを使用して、最終出力におけるコード ブロックの堅牢なネストを保証します。
Gitignore 対応 : デフォルトで .gitignore ルールを尊重し、必要に応じて cli 引数 (--no-gitignore) を使用してこの機能をオフにできます。
🏗️ アーキテクチャとエンジニアリングの標準
このプロジェクトは、上級レベルのエンジニアリングの成熟度を反映した、モジュラー機能オーケストレーション (MFO) パターンのポートフォリオ グレードの実装です。
レジストリと戦略パターン : 拡張可能なファイル処理には ParserRegistry を使用し、複数の形式 (Markdown、XML) には OutputGenerator 戦略を使用します。
一元化された構成: すべてのコア ロジック、マジック ナンバー、およびデフォルトの無視リストは src/data2prompt/constants.py に存在します。
厳密な型ヒント: すべてのモジュールにわたる完全に型指定された関数シグネチャ (PEP 484)。
UI カプセル化:Al

l ターミナルフィードバックは専用の UIHandler によって処理され、ロジックとプレゼンテーションが明確に分離されます。
システム設計の詳細については、「アーキテクチャのドキュメント」を参照してください。
Python 3.10 以降がインストールされていることを確認してください。
推奨 - pipx を使用します (グローバル CLI ツールとしてインストールします)。
pipxをお持ちではありませんか？まずインストールしてください:
pip インストール pipx
pipx確保パス
次に、data2prompt をインストールします。
pipx インストール data2prompt
代替案 - pip を使用します (アクティブな仮想環境が必要です):
pip install data2prompt
最新バージョンに更新します。
# pipx 付き
pipx アップグレード data2prompt
# ピップ付き
pip install --upgrade data2prompt
ソースからインストールする
# リポジトリのクローンを作成します
git クローン https://github.com/arianmokhtariha/data2prompt.git
cd データ 2 プロンプト
# 通常通りインストール
pip インストール 。
# または編集可能モードでインストールします
pip install -e 。
使用法
プロジェクト ルートで data2prompt を実行して、構造化プロンプトを生成します。
# 基本的な使い方 (デフォルトはマークダウン出力)
データ2プロンプト
# XML 形式と特定のサンプリングを使用したカスタム出力
data2prompt --output my_analysis --format xml --csv-sample-size 50 --ignore-folders venv .pytest_cache
CLI の引数
引数
説明
デフォルト
-o 、 --output
生成されたファイルのベース名
プロンプト
-f 、 --format
出力形式 (xml または markdown)
値下げ
-s 、 --csv サンプルサイズ
CSV からサンプリングするランダムな行の数
15
--最大行数
ノートブックのセルあたりのテキスト出力の最大行数
40
--最大ファイルサイズ
全体を読み取るための最大ファイル サイズ (KB 単位)
70
引数の完全なリストについては、「CLI リファレンス」を参照してください。
詳細については、詳細なドキュメントを参照してください。
アーキテクチャ: MFO パターンとモジュール フロー。
CLI リファレンス : 引数の詳細な説明と使用法。
パーサー : さまざまなファイルタイプがどのように処理されるか。
出力形式: Markdown と XML 生成の詳細。
ユーザーインターフェース：high-tの特徴

ec TUI。
インストール : 包括的なセットアップ ガイド。
pip install -e .[dev]
pytest
🌟 サポートを示してください
Data2Prompt によってトークンのコストが節約されたり、ワークフローが高速化されたりする場合は、次のことを検討してください。
🐛 問題の報告または機能の提案
🔀 新しいファイルタイプに貢献するパーサー
最新の AI 支援開発ワークフローに合わせて正確に構築されています。
データ アナリスト/サイエンティスト プロジェクト向けに特別に作成された、カスタムメイドのリポジトリからプロンプトへのパッケージャです。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0 - 初期リリース
最新の
2026 年 4 月 23 日
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A custom made repository-to-prompt packager made specifically for data analyst/scientist projects. - arianmokhtariha/data2prompt

GitHub - arianmokhtariha/data2prompt: A custom made repository-to-prompt packager made specifically for data analyst/scientist projects. · GitHub
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
arianmokhtariha
/
data2prompt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
58 Commits 58 Commits .github/ workflows .github/ workflows assets assets docs docs plans plans src/ data2prompt src/ data2prompt tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
High-performance codebase-to-prompt orchestration for Data Science workflows and data-heavy projects.
data2prompt is a CLI tool designed to bridge the gap between local data-heavy projects and Large Language Model (LLM) context windows. Unlike generic code-packagers, it provides an intelligent,optimized output for LLM attention mechanism, token-aware representation of a project's structure and content.
Data2prompt is purpose-built for data-heavy projects ( .csv , .sql , .xlsx , .ipynb ), not large pure-code repositories. It intelligently samples and truncates data files to prevent context window explosion while preserving semantic structure.
Generic code-to-prompt tools choke on data files—they either skip them entirely or dump raw CSVs that waste 90% of your context window. Data2Prompt solves this with intelligent sampling, schema extraction, and LLM-optimized formatting specifically designed for data science workflows.
Smart Jupyter Parsing : Intelligently extracts code, markdown, and text outputs from .ipynb files while stripping heavy Base64 images and raw HTML to preserve context.
Multi-Format Sampling : Advanced sampling strategies for CSV, SQL, and Excel files to preserve schema and data context which reduces the data size significantly while extracting the needed context for llm.
Aggressive truncations : To preserve context, long lines are truncated to neutralize line injections and avoid exploding the context windows, if a tabular data was still to large after sampling it will get truncated to a certain amount, also if a raw text file of unhandled type was too large it will get truncated to a certain amount.
Defensive Processing : Automatic binary detection (Null-byte checks), Checks if a file is binary by looking for a Null byte in the first 1024 bytes.
Optimized LLM attention : The default output format is markdown with well structured schema and another option is xml output with xml style tags to enhance LLM anchoring for complex analysis and large context windows
Token-Aware Output : Real-time token estimation using tiktoken ( o200k_base ) to ensure prompts fit target LLMs (Claude 3.5, GPT-4o, Gemini 1.5) and advanced offline token counting via regex .
Professional TUI : A high-fidelity terminal interface built with Rich , featuring a Matrix-style startup animation and interactive, scrollable reports on Windows.
Dynamic Markdown Wrapping : Uses intelligent backtick depth to ensure robust nesting of code blocks in the final output.
Gitignore aware : Respects the .gitignore rules by default and you can turn this feature off with cli argument(--no-gitignore) if needed.
🏗️ Architecture & Engineering Standards
This project is a portfolio-grade implementation of the Modular Functional Orchestration (MFO) pattern, reflecting senior-level engineering maturity:
Registry & Strategy Patterns : Uses a ParserRegistry for extensible file handling and an OutputGenerator strategy for multiple formats (Markdown, XML).
Centralized Configuration : All core logic, magic numbers, and default ignore lists reside in src/data2prompt/constants.py .
Strict Type Hinting : Fully typed function signatures (PEP 484) across all modules.
UI Encapsulation : All terminal feedback is handled by a dedicated UIHandler , ensuring a clean separation between logic and presentation.
For a deep dive into the system design, see the Architecture Documentation .
Ensure you have Python 3.10+ installed.
Recommended — using pipx (installs as a global CLI tool):
Don't have pipx? Install it first:
pip install pipx
pipx ensurepath
Then install data2prompt:
pipx install data2prompt
Alternative — using pip (requires an active virtual environment):
pip install data2prompt
Update to the latest version:
# with pipx
pipx upgrade data2prompt
# with pip
pip install --upgrade data2prompt
Install from the source
# Clone the repository
git clone https://github.com/arianmokhtariha/data2prompt.git
cd data2prompt
# Install normally
pip install .
# Or Install in editable mode
pip install -e .
Usage
Run data2prompt in your project root to generate a structured prompt:
# Basic usage (defaults to markdown output)
data2prompt
# Custom output with xml format and specific sampling
data2prompt --output my_analysis --format xml --csv-sample-size 50 --ignore-folders venv .pytest_cache
CLI Arguments
Argument
Description
Default
-o , --output
Base name of the generated file
PROMPT
-f , --format
Output format ( xml or markdown )
markdown
-s , --csv-sample-size
Number of random rows to sample from CSVs
15
--max-lines
Max lines of text output per notebook cell
40
--max-file-size
Max file size in KB to read entirely
70
See the CLI Reference for a full list of arguments.
Explore the detailed documentation for more information:
Architecture : MFO pattern and module flow.
CLI Reference : Detailed argument descriptions and usage.
Parsers : How different file types are handled.
Output Formats : Details on Markdown and XML generation.
User Interface : Features of the high-tech TUI.
Installation : Comprehensive setup guide.
pip install -e .[dev]
pytest
🌟 Show Your Support
If Data2Prompt saves you token costs or speeds up your workflow, consider:
🐛 Reporting issues or suggesting features
🔀 Contributing parsers for new file types
Built with precision for the modern AI-assisted development workflow.
A custom made repository-to-prompt packager made specifically for data analyst/scientist projects.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.0 - Initial Release
Latest
Apr 23, 2026
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
