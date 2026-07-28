---
source: "https://github.com/scouttyg/writekin"
hn_url: "https://news.ycombinator.com/item?id=49088436"
title: "Show HN: Writekin – fine-tune a local LLM on your own writing, on your Mac"
article_title: "GitHub - scouttyg/writekin: Writing that's kin to yours · GitHub"
author: "eggbrain"
captured_at: "2026-07-28T19:07:56Z"
capture_tool: "hn-digest"
hn_id: 49088436
score: 1
comments: 0
posted_at: "2026-07-28T19:05:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Writekin – fine-tune a local LLM on your own writing, on your Mac

- HN: [49088436](https://news.ycombinator.com/item?id=49088436)
- Source: [github.com](https://github.com/scouttyg/writekin)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T19:05:12Z

## Translation

タイトル: Show HN: Writekin – Mac 上で自分の書き込みに合わせてローカル LLM を微調整する
記事のタイトル: GitHub - scootyg/writekin: あなたに近いものを書く · GitHub
説明: あなたに近いものを書いています。 GitHub でアカウントを作成して、scouttyg/writekin の開発に貢献してください。
HN テキスト: こんにちは、ハッカー ニュース!機能しない AI ライティングにうんざりしたため、先週 Writekin を構築しました。
たとえ私がそれを大々的に書いたわけではなく、AIを使ってクリーンアップしただけだったとしても、私と同じように聞こえます。私がオンラインで見つけたこの問題に対する通常の修正は次のとおりです。 - ある種の SKILL.md、または - 一般的な AI 指示を取り除くためのルールが満載のシステム プロンプト (例: エムダッシュなし、ストック フレーズなし、文の長さの変更など)。これらは表面を少しきれいにしましたが、Pangram は依然として ~100% AI で書かれた状態で戻ってきました。これも主に私の雑なコピーを取り込んで微調整していたため、これにはイライラしました。そこで、Writekin を構築するときは、別の方法を採用しました。Writekin は、ユーザー自身の書き込みに基づいてローカル モデルを微調整します。あなたがすでに書いたもの（Apple Mail、iMessage、
ローカル ドキュメント、チャット エクスポートなど）をトレーニング コーパスにキュレーションし、
Apple の MLX を介してデバイス上で QLoRA 微調整を実行します。作成画面
その声で下書きとリライトを行います。すべてが Mac 上で実行されます。摂取、トレーニング、生成はすべて
地元の。唯一のネットワーク呼び出しは次のとおりです: (1) Hugging Face からモデル ウェイトをダウンロードするとき、および (2) Sparkle 更新チェック。独自のメール/メッセージでのトレーニングは、エンド ユーザーがそれを確認できる場合にのみ配布しても問題ないと考えられるため、ソースは公開されており、その内容を正確に読むことができます。クイック腸チェック: v0.9 ですが、出力が不均一です。正直なところ、時々それはあなたを釘付けにします
声が聞こえなくなったり、まったく聞こえなくなったりすることもあります。これはむしろ「これは可能であり、
完成品よりも「作品の種類」。するだろう

本当にフィードバックが大好きです!

記事本文:
GitHub - scootyg/writekin: 自分に近いものを書く · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
スカウト
/
ライトキン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店

es タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット ベンダー ベンダー Writekin Writekin WritekinTests WritekinTests スクリプト スクリプト サイト site .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md COTRIBUTING.md COTRIBUTING.md LICENSE.md LICENSE.md NOTICE.md NOTICE.md README.md README.md RELEASING.md RELEASING.md THIRD_PARTY_LICENSES.md THIRD_PARTY_LICENSES.md project.yml project.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あなたに近いものを書きます。 Writekin はあなた自身の文章を取り込みます —
メール、メッセージ、ドキュメント、チャットのエクスポート — トレーニングにまとめます
コーパス、ローカル言語モデルを微調整します (Apple Silicon 上の QLoRA
MLX )、次にドラフトと
あなたのように聞こえるテキストを書き換えます。すべては Mac 上で行われます。
概要 — コーパスの概要。
作成 — 自分の声で下書きして書き直します。
Train — ライブ損失チャートを使用したローカル QLoRA 実行。
モデル — デバイス上のモデル ライブラリ。
マシンからは何も残りません。摂取、トレーニング、生成
すべて地元のものです。テレメトリ、分析、アカウントはありません。
自動更新チェック (Sparkle) は 1 つのファイル (リリース フィード) を要求します。
システム プロファイルは送信されません。それがアプリの唯一のネットワーク トラフィックです
モデル/ツールのダウンロード以外にも、明示的に開始します。
モデルの重みはオンデマンドで Hugging Face からダウンロードできます。オプションの
メッセージ サポート ツールは、独自の公式 GitHub リリースからダウンロードされます。
チェックサム検証済み。どちらもクリックしたときにのみ発生します。
このリポジトリは存在するため、上記のいずれかを実行する必要はありません
信仰：情報源は証拠です。
Apple Silicon Mac (M シリーズ)。最小 16 GB のユニファイド メモリ。 32GB以上
より大きなモデルのトレーニングに推奨されます。
メールとメッセージを取り込みたい場合は、フル ディスク アクセス (ユーザー許可)。
最新の DMG を次からダウンロードします。
リリース 、ドラッグして
アプリケーション、開きます。 T

このアプリはソース検出をガイドします。完全版
初回起動時のディスク アクセスとモデルのダウンロード。
メッセージのサポートはオープンソースを使用します
imessage-エクスポーター
(GPL-3.0)、Writekin がオンデマンドでダウンロードするか、既存の
brew install imessage-exporter 。
醸造インストールxcodegen
xcodegen 生成
open Writekin.xcodeproj # Writekin スキームを構築する
完全なテスト スイート ( xcodebuild test -scheme Writekin ) に合格する必要があります。
変更を送信する前に — CONTRIBUTING.md を参照してください (貢献内容に注意してください)
ライセンス条項が記載されています)。
コードベースは初めてですか? ARCHITECTURE.md から始めます — パイプライン、
フォルダー マップと、それを正直に保つ不変条件です。
Writekin はソースが入手可能で、個人 (非営利) には無料です
PolyForm Noncommercial 1.0.0 ライセンスに基づいて使用します。
商用利用には別のライセンスが必要です。詳細については NOTICE.md を参照してください。
連絡先。サードパーティ製コンポーネントのリストは次のとおりです。
THIRD_PARTY_LICENSES.md 。
scouttyg.github.io/writekin/ リソース
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Writing that's kin to yours. Contribute to scouttyg/writekin development by creating an account on GitHub.

Hey Hacker News! I built Writekin over the past week because I was tired of AI writing that didn't
sound like me, even though I had just used AI to clean it up, rather than wholesale write it. The usual fixes I found online for this were: - Some sort of SKILL.md, or - A system prompt full of rules to strip the generic AI tells (e.g. no em-dashes, none of the stock phrases, varying the sentence length, etc). While those cleaned up the surface a bit, Pangram still came back as ~100% AI written, which was frustrating, as again it was mainly taking my sloppy copy and tweaking it. So when building Writekin I took a different route: Writekin fine-tunes a local model on your own writing. It reads what you've already written (Apple Mail, iMessage,
local documents, chat exports), curates it into a training corpus, and
runs a QLoRA fine-tuning on-device via Apple's MLX. A Compose screen then
drafts and rewrites in that voice. Everything runs on your Mac. Ingestion, training, and generation are all
local. The only network calls are: (1) When you download the model weights from Hugging Face and (2) The Sparkle update check. Training on your own Mail/Messages only felt okay to ship if the end user could verify that, so the source is public — so you can read exactly what it does! Quick gut check: It's v0.9 and the output is uneven. Honestly, sometimes it nails your
voice, and sometimes it's just completely off. This is more a "this is possible and
kind of works" than a finished product. Would genuinely love feedback!

GitHub - scouttyg/writekin: Writing that's kin to yours · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
scouttyg
/
writekin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit Vendor Vendor Writekin Writekin WritekinTests WritekinTests scripts scripts site site .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE.md LICENSE.md NOTICE.md NOTICE.md README.md README.md RELEASING.md RELEASING.md THIRD_PARTY_LICENSES.md THIRD_PARTY_LICENSES.md project.yml project.yml View all files Repository files navigation
Writing that's kin to yours. Writekin ingests your own writing —
Mail, Messages, documents, chat exports — curates it into a training
corpus, fine-tunes a local language model (QLoRA on Apple Silicon via
MLX ), and then drafts and
rewrites text that sounds like you . Everything happens on your Mac.
Overview — your corpus at a glance.
Compose — draft and rewrite in your voice.
Train — a local QLoRA run with a live loss chart.
Models — the on-device model library.
Nothing leaves your machine. Ingestion, training, and generation
are all local. There is no telemetry, no analytics, and no account.
The auto-update check (Sparkle) requests one file — the release feed —
and sends no system profile. That is the app's only network traffic
besides model/tool downloads you explicitly start.
Model weights download from Hugging Face on demand; the optional
Messages-support tool downloads from its own official GitHub release,
checksum-verified. Both happen only when you click.
This repository exists so you don't have to take any of the above on
faith: the source is the proof.
Apple Silicon Mac (M-series). 16 GB unified memory minimum; 32 GB+
recommended for training larger models.
Full Disk Access (user-granted) if you want Mail and Messages ingested.
Download the latest DMG from
Releases , drag to
Applications, open. The app guides you through source detection, Full
Disk Access, and model download on first launch.
Messages support uses the open-source
imessage-exporter
(GPL-3.0), which Writekin downloads on demand — or use an existing
brew install imessage-exporter .
brew install xcodegen
xcodegen generate
open Writekin.xcodeproj # build the Writekin scheme
The full test suite ( xcodebuild test -scheme Writekin ) must pass
before submitting changes — see CONTRIBUTING.md (note the contribution
licensing terms there).
New to the codebase? Start with ARCHITECTURE.md — the pipeline, the
folder map, and the invariants that keep it honest.
Writekin is source-available, free for personal (noncommercial)
use under the PolyForm Noncommercial 1.0.0 license.
Commercial use requires a separate license — see NOTICE.md for
contact. Third-party components are listed in
THIRD_PARTY_LICENSES.md .
scouttyg.github.io/writekin/ Resources
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
