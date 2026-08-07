---
source: "https://github.com/tacyan/zaivern-code"
hn_url: "https://news.ycombinator.com/item?id=49215129"
title: "Show HN: Zaivern Code – a Rust cockpit for parallel AI coding agents"
article_title: "GitHub - tacyan/zaivern-code: Zaivern Code — Rust-native AI cockpit editor · GitHub"
author: "tacyan"
captured_at: "2026-08-07T19:44:14Z"
capture_tool: "hn-digest"
hn_id: 49215129
score: 1
comments: 0
posted_at: "2026-08-07T19:23:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Zaivern Code – a Rust cockpit for parallel AI coding agents

- HN: [49215129](https://news.ycombinator.com/item?id=49215129)
- Source: [github.com](https://github.com/tacyan/zaivern-code)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T19:23:25Z

## Translation

タイトル: Show HN: Zaivern Code – 並列 AI コーディング エージェント用の Rust コックピット
記事タイトル: GitHub - tayan/zaivern-code: Zaivern コード — Rust ネイティブ AI コックピット エディター · GitHub
説明: Zaivern コード — Rust ネイティブの AI コックピット エディター。 GitHub でアカウントを作成して、tachan/zaivern コードの開発に貢献してください。

記事本文:
GitHub - tayan/zaivern-code: Zaivern コード — Rust ネイティブ AI コックピット エディター · GitHub
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
タシアン
/
ザイバーンコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
348 コミット 348 コミット .config .config .gemini .gemini .github/ workflows .github/ workflows Assets Assets docs docs src src tools/ licgen tools/ licgen ベンダー/ vt100 ベンダー/ vt100 .gitignore .gitignore CLAUDE.md CLA

UDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.en.md README.en.md README.md README.md build.rs build.rs install.ps1 install.ps1 install.sh install.sh View all files Repository files navigation
Claude Code・Codex・Gemini CLIなど、複数のAIコーディングツールをひとつの画面で動かす。
macOS・Windows・Linuxで使える、Rust製のAI開発コックピットです。
🌐 公式サイト ・ ⬇️ ダウンロード ・ 🗒️ リリース履歴
Zaivern CodeはAIそのものではなく、 複数のAIコーディングツールをまとめて操作するアプリ です。まずはClaude Code・Codex・Gemini CLIなど、使いたいAIツールを1つインストールしてログインしてください。3つすべてを用意する必要はありません。
Claude Codeが実装し、Codexがテストし、Gemini CLIがドキュメントを書く。Zaivern Codeは、散らばったターミナルを ひとつの操縦席 にまとめます。
curl -fsSL https://raw.githubusercontent.com/tacyan/zaivern-code/main/install.sh | sh
zai .
Windows PowerShell
irm https: // raw.githubusercontent.com / tacyan / zaivern - code / main / install.ps1 | iex
zai .
インストーラはお使いのOSに合うアプリを自動で取得します。同じコマンドをもう一度実行すれば最新版へ更新できます。
まずは1体だけで試し、慣れてから2体、3体と増やすのがおすすめです。
複数のAIツールをタイル状に並べ、動いているか、止まっているかをひと目で確認できます。Claude Code、Codex、Gemini CLIを含む29種類の起動設定を収録しています。
ひとつの入力欄から、動いている全AIへ同じ指示をまとめて送れます。も

ちろん、1体だけを選んで指示することもできます。
「この操作を許可しますか？」という確認待ちや、停止・異常終了を検知して通知します。安全のため、自動承認は 初期状態ではオフ です。
各AIが考え中・編集中・実行中・確認中のどこにいるかを一覧表示します。慣れてきたら、同じ課題を複数のAIへ同時に解かせて結果を比較することもできます。
スマホから進捗確認、指示、承認、ファイル編集ができます。まずは同じWi-Fi内で簡単に試せます。
コードを読んだり、AIが変更した箇所を確認したりできるエディタを内蔵しています。画像・PDF・CSV・Markdownなどもアプリ内で開けます。
TypeScript・Swift・Kotlin・Dart・Zig・TOML・Dockerfile・Terraformなどの構文ハイライト
6枚以上のCockpitタイルを読みやすく保つスクロール表示
品質: 2,565テスト成功、 cargo fmt --check 差分なし、clippy警告0。
項目
対応内容
OS
macOS arm64/x86_64・Linux x86_64/arm64・Windows x86_64
AI CLI
Claude Code・Codex・Gemini CLIほか、29種類のプリセット
Rust
ソースビルド時のみ1.88以上
ライセンス
Apache-2.0
安全設計
SSHトンネル使用時はリモートサーバーを 127.0.0.1 のみにバインド
セッション破棄・終了時に子プロセスを停止し、孤児プロセスを残さない
いいえ。Claude Code・Codex・Gemini CLIなど、使いたいAIツールを別途インストールしてログインしてください。
いいえ。1種類だけでも使えます。最初は普段使っているAIツール1つで試すのがおすすめです。
Zaivern CodeはApache-2.0ライセンスの

無料オープンソースソフトウェアです。各AIサービスの利用料金や契約は別途必要です。
初期状態では承認が必要です。自動承認を使う場合も、自分で明示的にオンへ切り替えます。
git clone https://github.com/tacyan/zaivern-code.git
cd zaivern-code
rustup update stable
cargo run --release -- .
テスト
cargo fmt --all --check
cargo nextest run --profile ci
プラグイン開発については プラグインガイド と 仕様書 を参照してください。
不具合報告・機能提案・Pull Requestを歓迎します。まず Issues で既存の報告を確認してください。
エージェントは、もう十分に速い。次に速くなるのは、指揮するあなたです。
Zaivern Code — Rust-native AI cockpit editor
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Zaivern Code — Rust-native AI cockpit editor . Contribute to tacyan/zaivern-code development by creating an account on GitHub.

GitHub - tacyan/zaivern-code: Zaivern Code — Rust-native AI cockpit editor · GitHub
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
tacyan
/
zaivern-code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
348 Commits 348 Commits .config .config .gemini .gemini .github/ workflows .github/ workflows assets assets docs docs src src tools/ licgen tools/ licgen vendor/ vt100 vendor/ vt100 .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.en.md README.en.md README.md README.md build.rs build.rs install.ps1 install.ps1 install.sh install.sh View all files Repository files navigation
Claude Code・Codex・Gemini CLIなど、複数のAIコーディングツールをひとつの画面で動かす。
macOS・Windows・Linuxで使える、Rust製のAI開発コックピットです。
🌐 公式サイト ・ ⬇️ ダウンロード ・ 🗒️ リリース履歴
Zaivern CodeはAIそのものではなく、 複数のAIコーディングツールをまとめて操作するアプリ です。まずはClaude Code・Codex・Gemini CLIなど、使いたいAIツールを1つインストールしてログインしてください。3つすべてを用意する必要はありません。
Claude Codeが実装し、Codexがテストし、Gemini CLIがドキュメントを書く。Zaivern Codeは、散らばったターミナルを ひとつの操縦席 にまとめます。
curl -fsSL https://raw.githubusercontent.com/tacyan/zaivern-code/main/install.sh | sh
zai .
Windows PowerShell
irm https: // raw.githubusercontent.com / tacyan / zaivern - code / main / install.ps1 | iex
zai .
インストーラはお使いのOSに合うアプリを自動で取得します。同じコマンドをもう一度実行すれば最新版へ更新できます。
まずは1体だけで試し、慣れてから2体、3体と増やすのがおすすめです。
複数のAIツールをタイル状に並べ、動いているか、止まっているかをひと目で確認できます。Claude Code、Codex、Gemini CLIを含む29種類の起動設定を収録しています。
ひとつの入力欄から、動いている全AIへ同じ指示をまとめて送れます。もちろん、1体だけを選んで指示することもできます。
「この操作を許可しますか？」という確認待ちや、停止・異常終了を検知して通知します。安全のため、自動承認は 初期状態ではオフ です。
各AIが考え中・編集中・実行中・確認中のどこにいるかを一覧表示します。慣れてきたら、同じ課題を複数のAIへ同時に解かせて結果を比較することもできます。
スマホから進捗確認、指示、承認、ファイル編集ができます。まずは同じWi-Fi内で簡単に試せます。
コードを読んだり、AIが変更した箇所を確認したりできるエディタを内蔵しています。画像・PDF・CSV・Markdownなどもアプリ内で開けます。
TypeScript・Swift・Kotlin・Dart・Zig・TOML・Dockerfile・Terraformなどの構文ハイライト
6枚以上のCockpitタイルを読みやすく保つスクロール表示
品質: 2,565テスト成功、 cargo fmt --check 差分なし、clippy警告0。
項目
対応内容
OS
macOS arm64/x86_64・Linux x86_64/arm64・Windows x86_64
AI CLI
Claude Code・Codex・Gemini CLIほか、29種類のプリセット
Rust
ソースビルド時のみ1.88以上
ライセンス
Apache-2.0
安全設計
SSHトンネル使用時はリモートサーバーを 127.0.0.1 のみにバインド
セッション破棄・終了時に子プロセスを停止し、孤児プロセスを残さない
いいえ。Claude Code・Codex・Gemini CLIなど、使いたいAIツールを別途インストールしてログインしてください。
いいえ。1種類だけでも使えます。最初は普段使っているAIツール1つで試すのがおすすめです。
Zaivern CodeはApache-2.0ライセンスの無料オープンソースソフトウェアです。各AIサービスの利用料金や契約は別途必要です。
初期状態では承認が必要です。自動承認を使う場合も、自分で明示的にオンへ切り替えます。
git clone https://github.com/tacyan/zaivern-code.git
cd zaivern-code
rustup update stable
cargo run --release -- .
テスト
cargo fmt --all --check
cargo nextest run --profile ci
プラグイン開発については プラグインガイド と 仕様書 を参照してください。
不具合報告・機能提案・Pull Requestを歓迎します。まず Issues で既存の報告を確認してください。
エージェントは、もう十分に速い。次に速くなるのは、指揮するあなたです。
Zaivern Code — Rust-native AI cockpit editor
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
