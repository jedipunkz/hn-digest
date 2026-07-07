---
source: "https://github.com/qualcomm/GenieX"
hn_url: "https://news.ycombinator.com/item?id=48825032"
title: "Qualcomm acquires Nexa AI, open-sources GenAI runtime for Hexagon NPUs"
article_title: "GitHub - qualcomm/GenieX: Run frontier LLMs and VLMs locally on Qualcomm devices across NPU, GPU, and CPU with a few lines of code · GitHub"
author: "BUFU"
captured_at: "2026-07-07T22:51:46Z"
capture_tool: "hn-digest"
hn_id: 48825032
score: 2
comments: 1
posted_at: "2026-07-07T22:44:11Z"
tags:
  - hacker-news
  - translated
---

# Qualcomm acquires Nexa AI, open-sources GenAI runtime for Hexagon NPUs

- HN: [48825032](https://news.ycombinator.com/item?id=48825032)
- Source: [github.com](https://github.com/qualcomm/GenieX)
- Score: 2
- Comments: 1
- Posted: 2026-07-07T22:44:11Z

## Translation

タイトル: クアルコムが Nexa AI を買収、Hexagon NPU 用のオープンソース GenAI ランタイム
記事のタイトル: GitHub - qualcomm/GenieX: 数行のコードで NPU、GPU、CPU 全体で Qualcomm デバイス上でフロンティア LLM と VLM をローカルに実行する · GitHub
説明: 数行のコードでフロンティア LLM と VLM を Qualcomm デバイス上で NPU、GPU、CPU 全体でローカルに実行します - qualcomm/GenieX

記事本文:
GitHub - qualcomm/GenieX: 数行のコードで NPU、GPU、CPU 全体で Qualcomm デバイス上でローカルにフロンティア LLM と VLM を実行 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
クアルコム
/
ジーニーX
公共
通知

ション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,729 コミット 1,729 コミット .claude .claude .github .github バインディング バインディング cli cli docs docs 例/ Python 例/ Python ノート ノート スクリプト スクリプト sdk sdk テスト テスト サードパーティ サードパーティ .bazelignore .bazelignore .bazelrc .bazelrc .bazelversion .bazelversion .clang-format .clang-format .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules .mailmap .mailmap BUILD.bazel BUILD.bazel CLAUDE.md CLAUDE.md CODE-OF-CONDUCT.md CODE-OF-CONDUCT.md CODEOWNERS CODEOWNERS CTRIBUTING.md CONTRIBUTING.md GenieX-Logo-Hor-1-Black.png GenieX-Logo-Hor-1-Black.png GenieX-Logo-Hor-1-White.png GenieX-Logo-Hor-1-White.png ライセンス ライセンス MODULE.bazel MODULE.bazel MODULE.bazel.lock MODULE.bazel.lock 通知 通知 README.md README.md README_zh-CN.md README_zh-CN.md SECURITY.md SECURITY.md repolint.json repolint.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
フロンティア LLM および VLM を Qualcomm デバイス上でローカルに実行する最も簡単な方法
ドキュメント · クイックスタート · モデル · コミュニティ
GenieX は、Qualcomm デバイス用のオンデバイス Gen AI 推論ランタイムです。 Hugging Face からほぼすべての GGUF モデル、または Qualcomm AI Hub からコンパイル済みのバンドルを取得し、数行のコードで Hexagon NPU、Adreno GPU、または CPU 上でローカルに実行します。その下に 1 つの C SDK があり、CLI、Python、Kotlin/Java、Docker、OpenAI 互換サーバーを通じて公開されます。 Qualcomm GENIE のコミュニティ版です。
GenieX は Qualcomm Snapdragon 上でのみ実行されます。プラットフォームを見つけて、使用するインターフェイスに直接ジャンプします。
手元にデバイスがありませんか? Qualcomm Device Cloud でリモート セッションを起動します。
以下からインターフェースを選択してください。それぞれ

この例では、同じ 3 つの手順 ( Install 、 Run 、 Docs ) に従い、両方のランタイムを示しています。 Hugging Face ( llama_cpp ) の GGUF モデルと、Qualcomm AI Hub ( qairt 、 NPU) のコンパイル済みバンドルです。
Windows ARM64 — インストーラーをダウンロードして実行し、新しいターミナルを開きます。
Linux ARM64 — 1 行、 sudo なし:
カール -fsSL https://qaihub-public-assets.s3.us-west-2.amazonaws.com/qai-hub-geniex/install.sh |しー
実行 — 任意のモデルと 1 行でチャットします (VLM の場合は画像をドラッグします)。
# ハグフェイスからGGUF → llama.cpp (NPU / GPU / CPU)
geniex 推論 google/gemma-4-E4B-it-qat-q4_0-gguf
# Qualcomm AI Hub → Qualcomm AI Engine Direct (NPU) からのプリコンパイル済みバンドル
geniex infer ai-hub-models/Qwen2.5-VL-7B-Instruct
📖 ドキュメント — インストール · クイックスタート · コマンドリファレンス
pip インストール geniex
実行 — ハグフェイストランスフォーマーをミラーリングします ( from_pretrained() → .generate() ):
# ハグフェイスのGGUF → llama.cpp
geniex から AutoModelForCausalLM をインポート
モデル = AutoModelForCausalLM 。 from_pretrained ( "unsloth/Qwen3.5-2B-GGUF" 、精度 = "Q4_0" )
messages = [{ "role" : "user" , "content" : "2+2 とは何ですか?" }]
プロンプト = モデル 。トークナイザー 。 apply_chat_template (メッセージ、add_generation_prompt = True)
モデル内のチャンク用。生成 (プロンプト、max_new_tokens = 256、ストリーム = True):
print (チャンク、エンド = ""、フラッシュ = True)
モデル。閉じる（）
# Qualcomm AI Hub → Qualcomm AI Engine Direct (NPU) からのプリコンパイル済みバンドル
geniex から AutoModelForCausalLM をインポート
モデル = AutoModelForCausalLM 。 from_pretrained ( "ai-hub-models/Qwen3-4B" )
messages = [{ "role" : "user" , "content" : "2+2 とは何ですか?" }]
プロンプト = モデル 。トークナイザー 。 apply_chat_template (メッセージ、add_generation_prompt = True)
モデル内のチャンク用。生成 (プロンプト、max_new_tokens = 256、ストリーム = True):
print ( チャンク , 終了 = ""

、フラッシュ = True )
モデル。閉じる（）
📖 ドキュメント — インストール · クイックスタート · API リファレンス
インストール — CLI に同梱されています (上記のインストール)。
実行 — 任意のモデル (GGUF または Qualcomm AI Hub バンドル) をプルし、OpenAI 互換 API を提供します。
geniex プル ai-hub-models/Qwen3-4B-Instruct-2507
geniex サーブ # サーブ http://127.0.0.1:18181/v1
カール http://127.0.0.1:18181/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {
"モデル": "ai-hub-models/Qwen3-4B-Instruct-2507",
"メッセージ": [{"役割": "ユーザー", "コンテンツ": "こんにちは!"}]
} '
OpenAI クライアントを http://127.0.0.1:18181/v1 に指定します。コードは変更されません。
インストール — SDK をアプリ モジュールの build.gradle.kts に追加します。
依存関係 {
実装( " com.qualcomm.qti:geniex-android:0.3.1 " )
}
実行 — 最速のパスはサンプル アプリです (チャット UI、GGUF のモデル ピッカー + Qualcomm AI Hub バンドル、VLM サポート)。
Android デモ アプリは qualcomm/ai-hub-apps にあります。クローンを作成し、Android Studio でサンプル アプリを開き、 Run をクリックします。
📖 ドキュメント — インストール · クイックスタート · API リファレンス
docker pull docker.io/qualcomm/geniex:latest
実行 — コンテナーが CLI をラップするため、geniex infer … は上記とまったく同じように機能します。
インストール — 単一の C ヘッダー sdk/include/geniex.h に対するリンク。他のすべてのインターフェイスは、その上の薄いラッパーです。
📖 ドキュメント — sdk/README.md · Notes/build.md
GenieX には 2 つのランタイムがあるため、1 つのスタックで幅広いモデルをカバーし、最高の Snapdragon パフォーマンスを得ることができます。 LLM と VLM の両方がサポートされています。
llama.cpp の場合は、プロンプトが表示されたら Q4_0 精度を選択します。これは最高の Hexagon NPU サポートを備えています。完全なリスト、精度、ローカル モデルの実行方法については、→「モデル ガイド」を参照してください。
貢献は大歓迎です！ PR を開く前に、ブランチの名前付け、コミット / PR タイトルの形式、コミット前のチェック、およびブランチの FFI 更新ルールについて CONTRIBUTING.md を読んでください。

パブリック SDK ヘッダー。
質問、アイデア、または自分が構築したものを披露したいことがありますか?挨拶に来てください。
💬 Slack — リアルタイムで質問したり、コミュニティとチャットしたりできます。
🐛 GitHub の問題 — バグを報告するか、機能をリクエストしてください。
🔗 LinkedIn — Qualcomm AI Hub をフォローしてニュースや最新情報を入手してください。
GenieX を構築している皆さんに感謝します 💙
BSD 3 条項 — 「ライセンス」と「通知」を参照してください。
このプロジェクトの使用には、クアルコムの利用規約も適用されます。
数行のコードでフロンティア LLM と VLM を NPU、GPU、CPU 全体で Qualcomm デバイス上でローカルに実行します
geniex.aihub.qualcomm.com/en/get-started/what-is-geniex
トピックス
BSD-3 条項ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1k
フォーク
レポートリポジトリ
リリース
60
v0.3.14
最新の
2026 年 7 月 2 日
+ 59 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Run frontier LLMs and VLMs locally on Qualcomm devices across NPU, GPU, and CPU with a few lines of code - qualcomm/GenieX

GitHub - qualcomm/GenieX: Run frontier LLMs and VLMs locally on Qualcomm devices across NPU, GPU, and CPU with a few lines of code · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
qualcomm
/
GenieX
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,729 Commits 1,729 Commits .claude .claude .github .github bindings bindings cli cli docs docs examples/ python examples/ python notes notes scripts scripts sdk sdk tests tests third-party third-party .bazelignore .bazelignore .bazelrc .bazelrc .bazelversion .bazelversion .clang-format .clang-format .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules .mailmap .mailmap BUILD.bazel BUILD.bazel CLAUDE.md CLAUDE.md CODE-OF-CONDUCT.md CODE-OF-CONDUCT.md CODEOWNERS CODEOWNERS CONTRIBUTING.md CONTRIBUTING.md GenieX-Logo-Hor-1-Black.png GenieX-Logo-Hor-1-Black.png GenieX-Logo-Hor-1-White.png GenieX-Logo-Hor-1-White.png LICENSE LICENSE MODULE.bazel MODULE.bazel MODULE.bazel.lock MODULE.bazel.lock NOTICE NOTICE README.md README.md README_zh-CN.md README_zh-CN.md SECURITY.md SECURITY.md repolint.json repolint.json View all files Repository files navigation
The easiest way to run frontier LLMs & VLMs locally on Qualcomm devices
Documentation · Quickstart · Models · Community
GenieX is an on-device Gen AI inference runtime for Qualcomm devices . Bring almost any GGUF model from Hugging Face — or a pre-compiled bundle from Qualcomm AI Hub — and run it locally on the Hexagon NPU, Adreno GPU, or CPU in a few lines of code. One C SDK underneath, exposed through a CLI, Python, Kotlin/Java, Docker, and an OpenAI-compatible server. It is the community version of Qualcomm GENIE.
GenieX runs only on Qualcomm Snapdragon . Find your platform, then jump straight to the interface you want to use.
No device on hand? Spin up a remote session on Qualcomm Device Cloud .
Pick your interface below. Each one follows the same three steps — Install , Run , and Docs — and shows both runtimes: a GGUF model from Hugging Face ( llama_cpp ) and a pre-compiled bundle from Qualcomm AI Hub ( qairt , NPU).
Windows ARM64 — download the installer , run it, then open a new terminal.
Linux ARM64 — one line, no sudo :
curl -fsSL https://qaihub-public-assets.s3.us-west-2.amazonaws.com/qai-hub-geniex/install.sh | sh
Run — chat with any model in one line (drag in an image for VLMs):
# GGUF from Hugging Face → llama.cpp (NPU / GPU / CPU)
geniex infer google/gemma-4-E4B-it-qat-q4_0-gguf
# Pre-compiled bundle from Qualcomm AI Hub → Qualcomm AI Engine Direct (NPU)
geniex infer ai-hub-models/Qwen2.5-VL-7B-Instruct
📖 Docs — Install · Quickstart · Command reference
pip install geniex
Run — mirrors Hugging Face transformers ( from_pretrained() → .generate() ):
# GGUF from Hugging Face → llama.cpp
from geniex import AutoModelForCausalLM
model = AutoModelForCausalLM . from_pretrained ( "unsloth/Qwen3.5-2B-GGUF" , precision = "Q4_0" )
messages = [{ "role" : "user" , "content" : "What is 2+2?" }]
prompt = model . tokenizer . apply_chat_template ( messages , add_generation_prompt = True )
for chunk in model . generate ( prompt , max_new_tokens = 256 , stream = True ):
print ( chunk , end = "" , flush = True )
model . close ()
# Pre-compiled bundle from Qualcomm AI Hub → Qualcomm AI Engine Direct (NPU)
from geniex import AutoModelForCausalLM
model = AutoModelForCausalLM . from_pretrained ( "ai-hub-models/Qwen3-4B" )
messages = [{ "role" : "user" , "content" : "What is 2+2?" }]
prompt = model . tokenizer . apply_chat_template ( messages , add_generation_prompt = True )
for chunk in model . generate ( prompt , max_new_tokens = 256 , stream = True ):
print ( chunk , end = "" , flush = True )
model . close ()
📖 Docs — Install · Quickstart · API reference
Install — ships with the CLI ( install above ).
Run — pull any model (GGUF or Qualcomm AI Hub bundle), then serve an OpenAI-compatible API:
geniex pull ai-hub-models/Qwen3-4B-Instruct-2507
geniex serve # serves http://127.0.0.1:18181/v1
curl http://127.0.0.1:18181/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {
"model": "ai-hub-models/Qwen3-4B-Instruct-2507",
"messages": [{"role": "user", "content": "Hello!"}]
} '
Point any OpenAI client at http://127.0.0.1:18181/v1 — no code changes.
Install — add the SDK to your app module's build.gradle.kts :
dependencies {
implementation( " com.qualcomm.qti:geniex-android:0.3.1 " )
}
Run — fastest path is the sample app (chat UI, model picker for GGUF + Qualcomm AI Hub bundles, VLM support):
The Android demo app lives in qualcomm/ai-hub-apps . Clone it, open the sample app in Android Studio, and hit Run .
📖 Docs — Install · Quickstart · API reference
docker pull docker.io/qualcomm/geniex:latest
Run — the container wraps the CLI, so geniex infer … works exactly as above.
Install — link against the single C header sdk/include/geniex.h ; every other interface is a thin wrapper over it.
📖 Docs — sdk/README.md · notes/build.md
GenieX has two runtimes so you get broad model coverage and peak Snapdragon performance in one stack. Both LLMs and VLMs are supported.
For llama.cpp, pick the Q4_0 precision when prompted — it has the best Hexagon NPU support. See the Models guide → for the full list, precisions, and how to run a local model.
Contributions are welcome! Before opening a PR, please read CONTRIBUTING.md for branch naming, commit / PR title format, pre-commit checks, and the FFI-update rule for public SDK headers.
Questions, ideas, or want to show off what you built? Come say hi.
💬 Slack — ask questions and chat with the community in real time.
🐛 GitHub Issues — report a bug or request a feature.
🔗 LinkedIn — follow Qualcomm AI Hub for news and updates.
Thanks to everyone building GenieX 💙
BSD 3-Clause — see LICENSE and NOTICE .
Use of this project is also subject to Qualcomm's Terms of Use .
Run frontier LLMs and VLMs locally on Qualcomm devices across NPU, GPU, and CPU with a few lines of code
geniex.aihub.qualcomm.com/en/get-started/what-is-geniex
Topics
BSD-3-Clause license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1k
forks
Report repository
Releases
60
v0.3.14
Latest
Jul 2, 2026
+ 59 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
