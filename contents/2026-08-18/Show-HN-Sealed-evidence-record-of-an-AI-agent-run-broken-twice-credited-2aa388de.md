---
source: "https://github.com/cjchanh/aaap-challenge"
hn_url: "https://news.ycombinator.com/item?id=49350964"
title: "Show HN: Sealed evidence record of an AI agent run – broken twice, credited"
article_title: "GitHub - cjchanh/aaap-challenge · GitHub"
image: "https://opengraph.githubassets.com/8fbb000fcfc06da017e8fe610686fa8d50accbd4399bba6b9a51a1cf91b865c5/cjchanh/aaap-challenge"
author: "cjarchivist"
captured_at: "2026-08-18T19:21:33Z"
capture_tool: "hn-digest"
hn_id: 49350964
score: 3
comments: 0
posted_at: "2026-08-18T19:07:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sealed evidence record of an AI agent run – broken twice, credited

- HN: [49350964](https://news.ycombinator.com/item?id=49350964)
- Source: [github.com](https://github.com/cjchanh/aaap-challenge)
- Score: 3
- Comments: 0
- Posted: 2026-08-18T19:07:02Z

## Translation

タイトル: Show HN: AI エージェント実行の封印された証拠記録 - 2 回破られ、クレジットされている
記事タイトル: GitHub - cjchanh/aaap-challenge · GitHub
説明: GitHub でアカウントを作成して、cjchanh/aaap-challenge の開発に貢献します。

記事本文:
GitHub - cjchanh/aaap-challenge · GitHub
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
チチャン
/
aaap チャレンジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
11 コミット 11 コミット フォルダーとファイル
証拠 証拠 ドキュメント ドキュメント パケット パケット セキュリティ セキュリティ テスト テスト AGENTS.md AGENTS.md CHALLENGE.md CHALLENGE.md DAYBREAK_CREDIT.md DAYBREAK_CREDIT.md DAYBREAK_II_CREDIT.md DAYBREAK_II_CREDIT.md ライセンス ライセンス README.md READM

E.md SECURITY.md SECURITY.md ウォークスルー.md ウォークスルー.md アンカー.json アンカー.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AAAP v0.5 — 強化されたリリース後の課題
コミット時の AAAP v0.3
e0fca96651b30c76d0a29bf867cdd14cdc38db00 は不変として保存されます。
再現性の高い脆弱性のある過去のリリース。このリポジトリは、
頑固な後継者。
v0.4 パケットは、v0.3 のchain.jsonl 、 artifacts/ 、および
attestation.json バイトごと。人間/検証層を再生成します
そして、同じアンカー実稼働 ID を持つマニフェスト スキーマ v2 に署名します。の
したがって、チェーン ヘッドと公開キーは残ります。
頭: 14d14281170d89f6f8b918daf6541f81e7e13549dd4c66f5b085304ff6a61724
公開キー: e7fb4aad8b0e0246eb6569f49d301ad88a0b54333e3de8bb57e02e118fd3716c
CHALLENGE.md から始めます。完全な v0.3 再現ツールと
マシンの結果はセキュリティ下にあります/ ;リリース後の所見台帳は
docs/AAAP_POST_RELEASE_SECURITY.md 。
要件: Python 3.10 以降、および Python 暗号化パッケージまたは
Ed25519 pkeyutl -rawin サポートを備えた OpenSSL ビルド。
CDパケット
python3 verify_packet.py 。 \
--expected-head 14d14281170d89f6f8b918daf6541f81e7e13549dd4c66f5b085304ff6a61724 \
--expected-pubkey e7fb4aad8b0e0246eb6569f49d301ad88a0b54333e3de8bb57e02e118fd3716c
レジストリ ポリシーを検証するには、両方の正確なレジストリを個別に取得します。
スナップショットを作成し、 --packet-name Demon/packet を使用します。両方のリポジトリが制御されます
by github.com/cjchanh : これらは 2 つの Git 履歴であり、2 つの管理または履歴ではありません。
保管信頼の起源。
PASS は、チェックされたバイト、正規のエンベロープ セマンティクス、順序付け、
マニフェスト ポリシーと、提供されたアンカーの下の署名。証明されない
世界の真実、署名時間、キーチェーンの保管場所、オペレーターの所有権、検証者/ホスト
慈善、レジストリの新鮮さ、または独立性

起源。
強化ソース: オペレーターのプライベートのローカル作業コミット
ワークスペース;公開されている系統は e0fca96 -> 8f5ba31f -> c06b163 -> です。
939b1e7 -> 9655a65 -> 06a3982 (現在)。
Readme Apache-2.0 ライセンス セキュリティ ポリシー
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to cjchanh/aaap-challenge development by creating an account on GitHub.

GitHub - cjchanh/aaap-challenge · GitHub
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
cjchanh
/
aaap-challenge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
11 Commits 11 Commits Folders and files
EVIDENCE EVIDENCE docs docs packet packet security security tests tests AGENTS.md AGENTS.md CHALLENGE.md CHALLENGE.md DAYBREAK_CREDIT.md DAYBREAK_CREDIT.md DAYBREAK_II_CREDIT.md DAYBREAK_II_CREDIT.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md WALKTHROUGH.md WALKTHROUGH.md anchors.json anchors.json View all files Repository files navigation
AAAP v0.5 — Hardened Post-Release Challenge
AAAP v0.3 at commit
e0fca96651b30c76d0a29bf867cdd14cdc38db00 is preserved as an immutable,
reproducibly vulnerable historical release. This repository is its
hardened successor.
The v0.4 packet preserves the v0.3 chain.jsonl , artifacts/ , and
attestation.json byte-for-byte. It regenerates the human/verification layer
and signs manifest schema v2 with the same anchored production identity. The
chain head and public key therefore remain:
head: 14d14281170d89f6f8b918daf6541f81e7e13549dd4c66f5b085304ff6a61724
pubkey: e7fb4aad8b0e0246eb6569f49d301ad88a0b54333e3de8bb57e02e118fd3716c
Start with CHALLENGE.md . The complete v0.3 reproducer and
machine results are under security/ ; the post-release finding ledger is
docs/AAAP_POST_RELEASE_SECURITY.md .
Requirements: Python 3.10+ and either the Python cryptography package or an
OpenSSL build with Ed25519 pkeyutl -rawin support.
cd packet
python3 verify_packet.py . \
--expected-head 14d14281170d89f6f8b918daf6541f81e7e13549dd4c66f5b085304ff6a61724 \
--expected-pubkey e7fb4aad8b0e0246eb6569f49d301ad88a0b54333e3de8bb57e02e118fd3716c
For registry-policy verification, independently acquire both exact registry
snapshots and use --packet-name demo/packet . Both repositories are controlled
by github.com/cjchanh : they are two Git histories, not two administrative or
custodial trust origins.
A PASS authenticates the checked bytes, canonical envelope semantics, ordering,
manifest policy, and signatures under the supplied anchor. It does not prove
world truth, signing time, Keychain custody, operator ownership, verifier/host
benevolence, registry freshness, or independent origin.
Hardening source: local working commit in the operator's private
workspace; the published lineage is e0fca96 -> 8f5ba31f -> c06b163 ->
939b1e7 -> 9655a65 -> 06a3982 (current).
Readme Apache-2.0 license Security policy
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
