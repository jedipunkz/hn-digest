---
source: "https://github.com/agentnameservice"
hn_url: "https://news.ycombinator.com/item?id=49374727"
title: "A registry and transparency log for discovering and verifying AI agents by name"
article_title: "agentnameservice · GitHub"
image: "https://avatars.githubusercontent.com/u/282946558?s=280&v=4"
author: "mooreds"
captured_at: "2026-08-20T14:26:29Z"
capture_tool: "hn-digest"
hn_id: 49374727
score: 2
comments: 0
posted_at: "2026-08-20T14:03:42Z"
tags:
  - hacker-news
  - translated
---

# A registry and transparency log for discovering and verifying AI agents by name

- HN: [49374727](https://news.ycombinator.com/item?id=49374727)
- Source: [github.com](https://github.com/agentnameservice)
- Score: 2
- Comments: 0
- Posted: 2026-08-20T14:03:42Z

## Translation

タイトル: AI エージェントを名前で検出および検証するためのレジストリおよび透明性ログ
記事のタイトル: エージェント名サービス · GitHub
説明: Agentnameservice には 9 つの使用可能なリポジトリがあります。 GitHub でコードをフォローしてください。

記事本文:
エージェント名サービス · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
AI エージェントを名前で検出して検証するためのレジストリと透明性のログ。
DNS はドメインをアドレスに解決します。 ANS はエージェント名を検証可能でバージョン管理された名前に解決します。
ID — 証明書と追加専用の透過性ログによって裏付けられます。
参考実装 ·
ゴーSDK
自律エージェントが組織を超えて相互に電話をかけ始めると、2 つの疑問が生じます。

ダメだ
今日の答え:
いったいこのエージェントは誰なのでしょうか？誰でも support-agent.acme.com であると主張できます。ないよ
それを証明する標準的な方法。
適切なものを見つけるにはどうすればよいですか?名前 + 機能 + バージョンをレジストリにマッピングするものはありません。
信頼できるエンドポイント。
ANS は両方に答えます。すべてのエージェントには、次のようなバージョン管理された DNS スタイルの名前が付けられます。
ans://v1.0.0.my-agent.example.com 、DNS/ACME ドメインの後にのみ発行される ID 証明書
検証と公開の透明性ログ記録 - そのため、あらゆる関係者が独立して検証できます。
仲介者を信頼することなく、エージェントの身元、能力、履歴を確認できます。
ピース
役割
レジストリ機関
エージェントを登録し、ドメイン所有権 (DNS + ACME) を検証し、ID とサーバー証明書を発行し、検索/解決を公開します。
透明性ログ
ある時点でのエージェントの状態を証明する COSE / SCITT レシート (RFC 9162、RFC 6962) を含む、すべてのエージェント イベントの耐久性のある追加専用のマークル ツリー ログ。
身分証明書
プライベート CA 署名証明書により、相互 TLS およびバッジ検証済みのエージェント間通話が可能になります。
オフライン検証者
ans-verify — ライブサービスに依存せずにエージェントの記録と証明を暗号的に検証します。
これは、エージェント向けに構築された DNS + 証明書の透明性と考えてください。
公開
IETF ドラフトに基づくエージェント ネーム サービス (ANS) のオープンソース リファレンス実装 (レジストリ、透過性ログ、Go のオフライン ベリファイア)。
読み込み中
種類
タイプを選択してください
すべて
公共
情報源
フォーク
アーカイブ済み
鏡
テンプレート
言語
言語を選択してください
すべて
行く
ジャワ
ルビー
さび
並べ替え
注文を選択してください
最終更新日
名前
星
9 件のリポジトリ中 9 件を表示
ans-sdk-rust
公共
Rust で書かれたエージェント名サービス SDK。
読み込み中にエラーが発生しました。このページをリロードしてください。
6
マサチューセッツ工科大学
0
2
7
2026 年 8 月 20 日更新
答え
公共
に基づくエージェント ネーム サービス (ANS) のオープンソース リファレンス実装

e IETF ドラフト — Go のレジストリ、透明性ログ、オフライン ベリファイア。
読み込み中にエラーが発生しました。このページをリロードしてください。
35
マサチューセッツ工科大学
5
6
4
2026 年 8 月 20 日更新
アンレジストリ
公共
https://datatracker.ietf.org/doc/html/draft-narajala-ans-00 に基づくエージェント ネーム サービス (ANS) レジストリ
読み込み中にエラーが発生しました。このページをリロードしてください。
29
13
5
5
2026 年 8 月 20 日更新
ans-sdk-java
公共
Javaで書かれたエージェントネームサービスSDK。
読み込み中にエラーが発生しました。このページをリロードしてください。
3
マサチューセッツ工科大学
1
1
4
2026 年 8 月 18 日更新
エージェントの信頼性の発見
公共
Agent Trust Discovery の Go リファレンス実装 — 登録された AI エージェントのインデックスを作成し、本番互換の検索 API を通じて説明可能な 5 次元の信頼スコア (完全性、アイデンティティ、支払い能力、動作、安全性) を提供します。
読み込み中にエラーが発生しました。このページをリロードしてください。
2
マサチューセッツ工科大学
3
1
2
2026 年 8 月 13 日更新
ans-sdk-go
公共
go (golang) で書かれたエージェント名サービス SDK。
読み込み中にエラーが発生しました。このページをリロードしてください。
5
マサチューセッツ工科大学
0
7
6
2026 年 8 月 13 日更新
読み込み中にエラーが発生しました。このページをリロードしてください。
スコップアン
公共
答えを得るにはスクープタップ
読み込み中にエラーが発生しました。このページをリロードしてください。
自作アンズ
公共
ansツールのHomebrewタブ
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

agentnameservice has 9 repositories available. Follow their code on GitHub.

agentnameservice · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
A registry and transparency log for discovering and verifying AI agents by name.
DNS resolves a domain to an address. ANS resolves an agent name to a verifiable, versioned
identity — backed by certificates and an append-only transparency log.
Reference implementation ·
Go SDK
As autonomous agents begin calling each other across organizations, two questions have no good
answer today:
Who is this agent, really? Anyone can claim to be support-agent.acme.com . There's no
standard way to prove it.
How do I find the right one? No registry maps a name + capability + version to a
trustworthy endpoint.
ANS answers both. Every agent gets a versioned, DNS-style name like
ans://v1.0.0.my-agent.example.com , an identity certificate issued only after DNS/ACME domain
validation, and a public transparency-log record — so any party can independently verify an
agent's identity, capabilities, and history without trusting a middleman.
Piece
Role
Registry Authority
Registers agents, validates domain ownership (DNS + ACME), issues identity & server certificates, and exposes search / resolution.
Transparency Log
A durable, append-only Merkle-tree log of every agent event, with COSE / SCITT receipts (RFC 9162, RFC 6962) proving an agent's state at a point in time.
Identity certificates
Private-CA-signed certs enabling mutual TLS and badge-verified agent-to-agent calls.
Offline verifier
ans-verify — cryptographically validates an agent's record and proofs with no live service dependency.
Think of it as DNS + Certificate Transparency, built for agents.
ans ans Public
Open-source reference implementation of the Agent Name Service (ANS) based on the IETF draft — registry, transparency log, and offline verifier in Go.
Loading
Type
Select type
All
Public
Sources
Forks
Archived
Mirrors
Templates
Language
Select language
All
Go
Java
Ruby
Rust
Sort
Select order
Last updated
Name
Stars
Showing 9 of 9 repositories
ans-sdk-rust
Public
Agent Name Service SDK written in rust.
There was an error while loading. Please reload this page .
6
MIT
0
2
7
Updated Aug 20, 2026
ans
Public
Open-source reference implementation of the Agent Name Service (ANS) based on the IETF draft — registry, transparency log, and offline verifier in Go.
There was an error while loading. Please reload this page .
35
MIT
5
6
4
Updated Aug 20, 2026
ans-registry
Public
An Agent Name Service (ANS) registry, based on https://datatracker.ietf.org/doc/html/draft-narajala-ans-00
There was an error while loading. Please reload this page .
29
13
5
5
Updated Aug 20, 2026
ans-sdk-java
Public
Agent Name Service SDK written in java.
There was an error while loading. Please reload this page .
3
MIT
1
1
4
Updated Aug 18, 2026
agent-trust-discovery
Public
Go reference implementation of the Agent Trust Discovery — indexes registered AI agents and serves explainable, five-dimension trust scores (integrity, identity, solvency, behavior, safety) through a production-compatible search API.
There was an error while loading. Please reload this page .
2
MIT
3
1
2
Updated Aug 13, 2026
ans-sdk-go
Public
Agent Name Service SDK written in go (golang).
There was an error while loading. Please reload this page .
5
MIT
0
7
6
Updated Aug 13, 2026
There was an error while loading. Please reload this page .
scoop-ans
Public
Scoop tap for ans
There was an error while loading. Please reload this page .
homebrew-ans
Public
Homebrew tab for ans tools
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
