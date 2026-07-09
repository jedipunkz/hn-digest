---
source: "https://github.com/Vadale/R-ebirth"
hn_url: "https://news.ycombinator.com/item?id=48843409"
title: "Relm – local LLMs as base-R objects, with interpretability"
article_title: "GitHub - Vadale/R-ebirth: Make R a first-class environment for AI and data-science research — the rebirth package (Rust core embedding a patched llama.cpp). · GitHub"
author: "grauk"
captured_at: "2026-07-09T10:46:53Z"
capture_tool: "hn-digest"
hn_id: 48843409
score: 1
comments: 0
posted_at: "2026-07-09T09:58:28Z"
tags:
  - hacker-news
  - translated
---

# Relm – local LLMs as base-R objects, with interpretability

- HN: [48843409](https://news.ycombinator.com/item?id=48843409)
- Source: [github.com](https://github.com/Vadale/R-ebirth)
- Score: 1
- Comments: 0
- Posted: 2026-07-09T09:58:28Z

## Translation

タイトル: Relm – 解釈可能性を備えた、base-R オブジェクトとしてのローカル LLM
記事のタイトル: GitHub - Vadale/R-ebirth: R を AI およびデータサイエンス研究のためのファーストクラスの環境にする — rebirth パッケージ (パッチを当てた llama.cpp を埋め込む Rust コア)。 · GitHub
説明: R を AI およびデータサイエンス研究のための一流の環境にします — rebirth パッケージ (パッチを当てた llama.cpp を埋め込んだ Rust コア)。 - バダーレ/R-ebirth

記事本文:
GitHub - Vadale/R-ebirth: R を AI およびデータサイエンス研究のための一流の環境にする — rebirth パッケージ (パッチを適用した llama.cpp を埋め込む Rust コア)。 · GitHub
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
バダーレ
/
R-ebirth
公共
通知
nを変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
227 コミット 227 コミット .claude .claude .github/ workflows .github/ workflows docs docsgraphify-outgraphify-out rebirth rebirth テスト テスト ベンダー ベンダー .gitignore .gitignore AGENTS.md AGENTS.md API-GRAMMAR.md API-GRAMMAR.md ARCHITECTURE.md ARCHITECTURE.md CLAUDE.md CLAUDE.md DECISIONS.md DECISIONS.md HANDOFF.md HANDOFF.md LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT LICENSE.md LICENSE.md 通知通知 R-vs-Python-La-Rivoluzione-Definitiva.md R-vs-Python-La-Rivoluzione-Definitiva.md README.md README.md ROADMAP.md ROADMAP.md SOLO-PHASE-PLAN.md SOLO-PHASE-PLAN.md THESIS-PLAN.md THESIS-PLAN.md TRADEMARK.md TRADEMARK.md Rust-toolchain.toml Rust-toolchain.toml ビューすべてのファイル リポジトリ ファイルのナビゲーション
R-ebirth は、R を科学研究のための一流の環境にすることを目指しています。
データと AI — 機械的解釈可能性 (「AI 神経科学」)、機械学習
トピックモデリングやライフサイエンスを含む - シンプルさを保ちながら
研究者たち。
これは relm として提供されます。Rust ネイティブ コアを備えた R パッケージです。
パッチを適用した llama.cpp を埋め込み、ローカル LLM (ロード、生成、
埋め込み、アクティベーション トレース、ステアリング、およびアブレーション) を基底 R イディオム関数として使用
プレーンな data.frame と行列を返します。
Python を使用しないトピック モデリング: llm_embed() → UMAP → HDBSCAN → モデル名
各クラスター。 2 つの実行可能なデモのうちの 1 つ — パッケージ README を参照してください。
パッケージを使用しますか?パッケージの README から始めます
(クイックスタート、例、2 つのデモ) および
docs/getting-started.md (インストール オプション - バイナリ)
またはソースから - 最初の実行とトラブルシューティング)。このページは、
リポジトリ/開発者の概要。
初公開はこちら

。 relm はローカル GGUF モデルをロードして公開します。
Base-R オブジェクトとして:
llm() モデルの読み込み、llm_tokens() トークン化。
llm_generate() テキスト生成、llm_logits() 次のトークンの配布。
llm_trace() アクティベーション トレース、llm_steer() ステアリング、llm_ablate()
アブレーション — メカニズムの解釈の核心。
llm_download() チェックサム検証済みの固定モデルのフェッチ。
すべての数値特徴は、独立したデータに対して価値対価値が検証されます。
参考（ハーネスB）。 Vision (画像入力) は次のリリース (v0.2.0) です。 v0.1.0
テキストのみです。完全な計画は ROADMAP.md にあります。
rebirth/ R パッケージ (R/、src/ + src/rust/ extendr crate、tests/、vignettes/)
Rust/Cargo ワークスペース: rebirth-ffi (R <-> Rust 境界)、rebirth-llm (エンジン)
rebirth/src/llama.cpp/ 固定され、パッチが適用された llama.cpp (ベンダー済み。VENDORING.md を参照)
testing/llm-golden/ ハーネス B 数値ゴールデン
testing/demos/ 2 つのリファレンス デモ (解剖学ラボ、Python を使用しないトピック)
計画文書 (唯一の信頼できる情報源)
CLAUDE.md 、 SOLO-PHASE-PLAN.md 、 ROADMAP.md 、 API-GRAMMAR.md 、
ARCHITECTURE.md 、 DECISIONS.md 、および THESIS-PLAN.md 。他に何かあれば
これらのファイルに同意しない場合、ファイルが優先されます。
ソースからの構築 (開発者)
エンド ユーザーは、r-universe から事前に構築されたバイナリをインストールします (ツールチェーンは必要ありません)。
ソースからのビルドには R (>= 4.5)、C ツールチェーン、Rust ツールチェーンが必要です
(rusup ; ピン留めされたチャネルは Rust-toolchain.toml にあります)、および CMake (>= 3.28)
ベンダーエンジン用。
# ネイティブワークスペース
cd 錆び && カーゴテスト && カーゴクリッピー --all-targets -- -D 警告
#Rパッケージ
R CMD ビルド再誕生 && R CMD チェック relm_0.1.0.tar.gz
ライセンス
デュアル ライセンスの MIT または Apache-2.0 — LICENSE.md を参照してください。販売されている
llama.cpp は MIT です (「通知」を参照)。名前は保護されています: 修正された再配布
名前を変更する必要があります( TRADEMARK.md を参照)。
R をファーストクラス環境にする

AI およびデータ サイエンス研究のためのアイロンメント — rebirth パッケージ (パッチを当てた llama.cpp を埋め込んだ Rust コア)。
不明なライセンスと他に 2 つのライセンスが見つかりました
ライセンスが見つかりました
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
relm 0.1.0 — 初公開リリース
最新の
2026 年 7 月 8 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Make R a first-class environment for AI and data-science research — the rebirth package (Rust core embedding a patched llama.cpp). - Vadale/R-ebirth

GitHub - Vadale/R-ebirth: Make R a first-class environment for AI and data-science research — the rebirth package (Rust core embedding a patched llama.cpp). · GitHub
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
Vadale
/
R-ebirth
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
227 Commits 227 Commits .claude .claude .github/ workflows .github/ workflows docs docs graphify-out graphify-out rebirth rebirth tests tests vendor vendor .gitignore .gitignore AGENTS.md AGENTS.md API-GRAMMAR.md API-GRAMMAR.md ARCHITECTURE.md ARCHITECTURE.md CLAUDE.md CLAUDE.md DECISIONS.md DECISIONS.md HANDOFF.md HANDOFF.md LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT LICENSE.md LICENSE.md NOTICE NOTICE R-vs-Python-La-Rivoluzione-Definitiva.md R-vs-Python-La-Rivoluzione-Definitiva.md README.md README.md ROADMAP.md ROADMAP.md SOLO-PHASE-PLAN.md SOLO-PHASE-PLAN.md THESIS-PLAN.md THESIS-PLAN.md TRADEMARK.md TRADEMARK.md rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
R-ebirth aims to make R a first-class environment for scientific research on
data and AI — mechanistic interpretability ("AI neuroscience"), machine learning
including topic modelling, and the life sciences — while staying simple for
researchers.
It is delivered as relm : an R package with a Rust native core that
embeds a patched llama.cpp , exposing local LLMs (loading, generation,
embeddings, activation tracing, steering, and ablation) as base-R-idiom functions
returning plain data.frame s and matrix es.
Topic modelling with no Python: llm_embed() → UMAP → HDBSCAN → the model names
each cluster. One of two runnable demos — see the package README .
Using the package? Start with the package README
(quickstart, examples, the two demos) and
docs/getting-started.md (install options — binaries
or from source — a first run, and troubleshooting). This page is the
repository/developer overview.
The first public release is here. relm loads local GGUF models and exposes,
as base-R objects:
llm() model loading, llm_tokens() tokenization;
llm_generate() text generation, llm_logits() next-token distributions;
llm_trace() activation tracing, llm_steer() steering, llm_ablate()
ablation — the mechanistic-interpretability core;
llm_download() checksum-verified fetch of pinned models.
Every numerical feature is validated value-for-value against an independent
reference (harness B). Vision (image inputs) is the next release (v0.2.0); v0.1.0
is text-only. The full plan is in ROADMAP.md .
rebirth/ the R package (R/, src/ + src/rust/ extendr crate, tests/, vignettes/)
rust/ Cargo workspace: rebirth-ffi (R <-> Rust boundary), rebirth-llm (engine)
rebirth/src/llama.cpp/ pinned, patched llama.cpp (vendored; see its VENDORING.md)
tests/llm-golden/ Harness B numerical goldens
tests/demos/ the two reference demos (anatomy lab; topics without Python)
Planning documents (the single source of truth)
CLAUDE.md , SOLO-PHASE-PLAN.md , ROADMAP.md , API-GRAMMAR.md ,
ARCHITECTURE.md , DECISIONS.md , and THESIS-PLAN.md . If anything else
disagrees with these files, the files win.
Building from source (developers)
End users install prebuilt binaries from r-universe (no toolchain required).
Building from source requires R (>= 4.5), a C toolchain, a Rust toolchain
( rustup ; the pinned channel is in rust-toolchain.toml ), and CMake (>= 3.28)
for the vendored engine.
# native workspace
cd rust && cargo test && cargo clippy --all-targets -- -D warnings
# R package
R CMD build rebirth && R CMD check relm_0.1.0.tar.gz
License
Dual-licensed MIT OR Apache-2.0 — see LICENSE.md . The vendored
llama.cpp is MIT (see NOTICE ). The name is protected: modified redistributions
must rename (see TRADEMARK.md ).
Make R a first-class environment for AI and data-science research — the rebirth package (Rust core embedding a patched llama.cpp).
Unknown and 2 other licenses found
Licenses found
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
relm 0.1.0 — first public release
Latest
Jul 8, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
