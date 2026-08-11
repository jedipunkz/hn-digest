---
source: "https://github.com/anthropics/claude-code/issues/65961"
hn_url: "https://news.ycombinator.com/item?id=49255222"
title: "Claude making verbose code comments – ignoring instructions to stop"
article_title: "[MODEL] Claude verbose code comments by default — ignores instructions to stop. · Issue #65961 · anthropics/claude-code · GitHub"
author: "nomilk"
captured_at: "2026-08-11T09:51:42Z"
capture_tool: "hn-digest"
hn_id: 49255222
score: 6
comments: 2
posted_at: "2026-08-11T09:04:55Z"
tags:
  - hacker-news
  - translated
---

# Claude making verbose code comments – ignoring instructions to stop

- HN: [49255222](https://news.ycombinator.com/item?id=49255222)
- Source: [github.com](https://github.com/anthropics/claude-code/issues/65961)
- Score: 6
- Comments: 2
- Posted: 2026-08-11T09:04:55Z

## Translation

タイトル: クロードが冗長なコード コメントを作成 – 停止の指示を無視
記事のタイトル: [モデル] クロードのデフォルトの冗長コード コメント — 停止命令を無視します。 · 問題 #65961 · anthropics/claude-code · GitHub
説明: プリフライト チェックリスト 既存の問題で同様の動作レポートを検索しました このレポートには機密情報 (API キー、パスワードなど) が含まれていません 動作の問題の種類 クロードがファイルを変更しました 変更するように依頼していません 何を...

記事本文:
[MODEL] デフォルトでクロードの冗長コード コメント — 停止命令を無視します。 · 問題 #65961 · anthropics/claude-code · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
人類学
/
クロードコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
[MODEL] デフォルトでクロードの冗長コード コメント — 停止命令を無視します。 #65961
リンクをコピーする 新規発行 リンクをコピーする 開く 開く [MODEL] デフォルトのクロード詳細コード コメント

— 停止の指示を無視します。 #65961 リンクをコピー ラベル エリア:モデルのバグ 何かが動作していません 何かが動作していません モデルの説明
発行本体アクションのプリフライト チェックリスト
既存の問題で同様の動作レポートを検索しました
このレポートには機密情報 (API キー、パスワードなど) は含まれません。
クロードは私が変更を依頼していないファイルを変更しました
Claude Code は、デフォルトで非常に多くのコード コメントを追加します。コメントはほとんどが冗長であり、隣接するコードですでに明らかになっている内容を再度述べています。それはすべての言語、すべてのモデルで発生します。
重要なのは、このデフォルトは、明示的に停止するように指示された場合でも維持されます。
CLAUDE.md の明確な必須ルールでは、これを確実に抑制できません。
記憶システムを介してルールを強化しても、それは止まりません。
中心的な問題は、詳細なコメントがすぐに使えるデフォルトであり、そのデフォルトが明示的なユーザー指示をオーバーライドするのに十分強力であることだと思います。ユーザーは、クリーンなコードを取得するためだけに、CLAUDE.md ルール + メモリ エントリ + 強制フックをスタックする必要はありません。
継続的な冗長なコードのコメント。
デフォルトとしての抑制されたコメント — 自明の「何を」ではなく、本当に非自明な制約または「なぜ」の決定のみをコメントします。また、プロジェクトの指示でコメントを最小限に抑えるように指示されている場合は、その指示を確実に尊重する必要があります。
別の指示があった場合でも、ほとんどのコード出力に再帰的な説明コメントが表示されます。定期的な手動クリーンアップが必要です。
編集の承認がオンになっていました (変更の自動承認)
はい、毎回同じプロンプトが表示されます
高 - 重大な望ましくない変更
リアクションは現在利用できません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Preflight Checklist I have searched existing issues for similar behavior reports This report does NOT contain sensitive information (API keys, passwords, etc.) Type of Behavior Issue Claude modified files I didn't ask it to modify What Y...

[MODEL] Claude verbose code comments by default — ignores instructions to stop. · Issue #65961 · anthropics/claude-code · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
anthropics
/
claude-code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
[MODEL] Claude verbose code comments by default — ignores instructions to stop. #65961
Copy link New issue Copy link Open Open [MODEL] Claude verbose code comments by default — ignores instructions to stop. #65961 Copy link Labels area:model bug Something isn't working Something isn't working model Description
Issue body actions Preflight Checklist
I have searched existing issues for similar behavior reports
This report does NOT contain sensitive information (API keys, passwords, etc.)
Claude modified files I didn't ask it to modify
Claude Code adds far too many code comments by default. The comments are mostly redundant, restating what the adjacent code already makes obvious. It happens on every language, every model.
Crucially, this default persists even when explicitly told to stop:
A clear, mandatory rule in CLAUDE.md does not reliably suppress it.
Reinforcing the rule via the memory system does not stop it either.
I suppose the core problem is that verbose commenting is the out-of-the-box default , and that default is strong enough to override explicit user instructions. Users shouldn't have to stack a CLAUDE.md rule + memory entries + enforcement hooks just to get clean code.
Contant verbose code commenting.
Restrained commenting as the default — comment only genuinely non-obvious constraints or "why" decisions, not self-evident "what." And when a project instruction says to minimize comments, that instruction should be reliably respected.
Reflexive explanatory comments on most code output, even when instructed otherwise. Requires constant manual cleanup .
Accept Edits was ON (auto-accepting changes)
Yes, every time with the same prompt
High - Significant unwanted changes
Reactions are currently unavailable Metadata
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
