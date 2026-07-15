---
source: "https://github.com/morrisalp/ConlangCrafter"
hn_url: "https://news.ycombinator.com/item?id=48918401"
title: "ConlangCrafter: Constructing Languages with a Multi-Hop LLM Pipeline"
article_title: "GitHub - morrisalp/ConlangCrafter: Constructing languages with LLMs, based on the ACL 2026 (Oral) paper: \"ConlangCrafter: Constructing Languages with a Multi-Hop LLM Pipeline\" · GitHub"
author: "jbotz"
captured_at: "2026-07-15T09:50:39Z"
capture_tool: "hn-digest"
hn_id: 48918401
score: 1
comments: 1
posted_at: "2026-07-15T09:45:23Z"
tags:
  - hacker-news
  - translated
---

# ConlangCrafter: Constructing Languages with a Multi-Hop LLM Pipeline

- HN: [48918401](https://news.ycombinator.com/item?id=48918401)
- Source: [github.com](https://github.com/morrisalp/ConlangCrafter)
- Score: 1
- Comments: 1
- Posted: 2026-07-15T09:45:23Z

## Translation

タイトル: ConlangCrafter: マルチホップ LLM パイプラインを使用した言語の構築
記事のタイトル: GitHub - morrisalp/ConlangCrafter: ACL 2026 (口頭) 論文に基づく LLM による言語の構築: 「ConlangCrafter: マルチホップ LLM パイプラインによる言語の構築」 · GitHub
説明: ACL 2026 (口頭) 論文に基づく、LLM を使用した言語の構築: 「ConlangCrafter: マルチホップ LLM パイプラインを使用した言語の構築」 - morrisalp/ConlangCrafter

記事本文:
GitHub - morrisalp/ConlangCrafter: ACL 2026 (口頭) 論文に基づく、LLM を使用した言語の構築: 「ConlangCrafter: マルチホップ LLM パイプラインを使用した言語の構築」 · GitHub
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
モリサルプ
/
コンランクラフター
公共
通知
君はね

サインインして通知設定を変更できない
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット configs configs プロンプト プロンプト src src .env.example .env.example .gitignore .gitignore .python-version .python-version LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml要件.txt要件.txt uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ConlangCrafter: マルチホップ LLM パイプラインを使用した言語の構築 (ACL 2026 口頭)
プロジェクトページ: conlangcrafter.github.io
論文: arxiv.org/abs/2508.06094
データセット：huggingface.co/datasets/malper/ConlangCrafter — 64 の生成言語
大規模な言語モデルを使用して言語 (conlang) を構築するための完全に自動化されたシステムを紹介します。当社の多段階パイプラインは、独自の音韻論、文法、語彙、翻訳機能を備えた一貫した多様な人工言語を作成します。
pip install -r 要件.txt
# または: UV を使用する場合は UV 同期
API キーを設定します。.env.example を .env にコピーし、使用する API のキーを追加します。
Google Gemini : GOOGLE_API_KEY — Google AI スタジオ
OpenAI : OPENAI_API_KEY — OpenAI API キー
DeepSeek (Together 経由) : TOGETHER_API_KEY — Together AI
言語スケッチを生成します (デフォルトのモデル: gemini-2.5-pro )。
Python src/run_pipeline.py
# または: uv run src/run_pipeline.py
すべてのオプションを表示するには、python src/run_pipeline.py --help を実行します。主要なフラグ:
Python src/run_pipeline.py \
--モデル ジェミニ-2.5-プロ \
--custom-constraints " この言語には母音が 3 つしかありません" \
--温度 0.8 \
--qa-disabled # QA 自己洗練ループはデフォルトでオンになっています。これを使用してオフにします
前回の実行を再開するには (例: 音韻論が完了した後に文法から開始する):
python src/run_pipeline.py -- language-id < id > --steps gram

3月、辞書
サポートされているモデルは次のとおりです。
Google Gemini (例: gemini-2.5-pro 、 gemini-1.5-flash )
OpenAI モデル (例: o4-mini 、 gpt-4o 、 gpt-5 )
Together AI 経由の DeepSeek (例: deepseek-ai/DeepSeek-R1 )
事前生成された言語スケッチ
次のスクリプトを使用して、事前生成された言語スケッチをデータセットからこのパイプラインの形式でロードできます。
Python src/load_hf_langages.py
翻訳
デフォルトでは翻訳は実行されません。生成された言語に翻訳するには、翻訳ステップを個別に実行します。デフォルトでは、 configs/sentences_default.txt 内の 10 文が翻訳されます。
python src/run_pipeline.py -- language-id < id > --steps 翻訳
代わりに単一のカスタム文を翻訳するには:
python src/run_pipeline.py -- language-id < id > --steps translation --translation-sentence " Hello, world! "
--translation-sketch-update を渡すと、翻訳中に導入された新しい語彙と文法規則が後続の各文のスケッチに戻され、翻訳が進むにつれて言語が拡張されます (建設的翻訳)。
この実装には、論文の結果に使用されるシステムへの小さな改良が含まれています。
QA ループ : 縮退出力 (テキストではなく JSON など) は、ポストホック拒否サンプリングではなく、インラインで検出およびスキップされます。
QA 修正プロンプト : システムとの一貫性を保つために、プロンプトの文言がわずかに調整されています。
@article { conlangcrafter2025 、
title = { ConlangCrafter: マルチホップ LLM パイプラインを使用した言語の構築 } ,
著者 = { モリス アルパー、モラン ヤヌカ、ラジャ ギリス、ガ{\v{s}}ペル ベグ{\v{s}} } ,
年 = { 2025 } 、
eprint = { 2508.06094 } 、
archivePrefix = { arXiv } 、
プライマリクラス = { cs.CL } 、
URL = { https://arxiv.org/abs/2508.06094 }
}
ライセンス
このプロジェクトは MIT ライセンスに基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
言語を構築する

th LLM、ACL 2026 (口頭) 論文「ConlangCrafter: Constructing Languages with a Multi-Hop LLM Pipeline」に基づく
conlangcrafter.github.io
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
3
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Constructing languages with LLMs, based on the ACL 2026 (Oral) paper: "ConlangCrafter: Constructing Languages with a Multi-Hop LLM Pipeline" - morrisalp/ConlangCrafter

GitHub - morrisalp/ConlangCrafter: Constructing languages with LLMs, based on the ACL 2026 (Oral) paper: "ConlangCrafter: Constructing Languages with a Multi-Hop LLM Pipeline" · GitHub
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
morrisalp
/
ConlangCrafter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits configs configs prompts prompts src src .env.example .env.example .gitignore .gitignore .python-version .python-version LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt uv.lock uv.lock View all files Repository files navigation
ConlangCrafter: Constructing Languages with a Multi-Hop LLM Pipeline (ACL 2026 Oral)
Project Page: conlangcrafter.github.io
Paper: arxiv.org/abs/2508.06094
Dataset: huggingface.co/datasets/malper/ConlangCrafter — 64 generated languages
We introduce a fully automated system for constructing languages (conlangs) using large language models. Our multi-stage pipeline creates coherent, diverse artificial languages with their own phonology, grammar, lexicon, and translation capabilities.
pip install -r requirements.txt
# or: uv sync if using uv
Set up API keys — copy .env.example to .env and add keys for whichever APIs you will use:
Google Gemini : GOOGLE_API_KEY — Google AI Studio
OpenAI : OPENAI_API_KEY — OpenAI API Keys
DeepSeek (via Together) : TOGETHER_API_KEY — Together AI
Generate a language sketch (default model: gemini-2.5-pro ):
python src/run_pipeline.py
# or: uv run src/run_pipeline.py
Run python src/run_pipeline.py --help to see all options. Key flags:
python src/run_pipeline.py \
--model gemini-2.5-pro \
--custom-constraints " The language has only 3 vowels " \
--temperature 0.8 \
--qa-disabled # QA self-refinement loops are on by default; use this to turn it off
To resume a previous run (e.g. starting from grammar after phonology completed):
python src/run_pipeline.py --language-id < id > --steps grammar,lexicon
Supported models are:
Google Gemini (e.g., gemini-2.5-pro , gemini-1.5-flash )
OpenAI models (e.g., o4-mini , gpt-4o , gpt-5 )
DeepSeek via Together AI (e.g., deepseek-ai/DeepSeek-R1 )
Pregenerated language sketches
You can load pregenerated language sketches from our dataset in this pipeline's format with this script:
python src/load_hf_languages.py
Translation
Translation is not run by default. To translate into a generated language, run the translation step separately. By default it translates the 10 sentences in configs/sentences_default.txt :
python src/run_pipeline.py --language-id < id > --steps translation
To translate a single custom sentence instead:
python src/run_pipeline.py --language-id < id > --steps translation --translation-sentence " Hello, world! "
Pass --translation-sketch-update to feed new vocabulary and grammar rules introduced during translation back into the sketch for each subsequent sentence, expanding the language as translation proceeds (constructive translation).
This implementation includes minor improvements to the system used for results from our paper:
QA loop : Degenerate outputs (e.g. JSON instead of text) are detected and skipped inline, rather than post-hoc rejection sampling.
QA amend prompt : Prompt wording is slightly adjusted for consistency with our system.
@article { conlangcrafter2025 ,
title = { ConlangCrafter: Constructing Languages with a Multi-Hop LLM Pipeline } ,
author = { Morris Alper and Moran Yanuka and Raja Giryes and Ga{\v{s}}per Begu{\v{s}} } ,
year = { 2025 } ,
eprint = { 2508.06094 } ,
archivePrefix = { arXiv } ,
primaryClass = { cs.CL } ,
url = { https://arxiv.org/abs/2508.06094 }
}
License
This project is licensed under the MIT License — see the LICENSE file for details.
Constructing languages with LLMs, based on the ACL 2026 (Oral) paper: "ConlangCrafter: Constructing Languages with a Multi-Hop LLM Pipeline"
conlangcrafter.github.io
Topics
There was an error while loading. Please reload this page .
3
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
