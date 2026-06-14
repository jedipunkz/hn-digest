---
source: "https://github.com/nex-agi/Nex-N2/issues/4"
hn_url: "https://news.ycombinator.com/item?id=48528371"
title: "Rio de Janeiro's \"homegrown\" LLM appears to be a merge of an existing model"
article_title: "Rio-3.5-Open-397B ≈ 0.6 x Nex-N2_pro + 0.4 x Qwen · Issue #4 · nex-agi/Nex-N2 · GitHub"
author: "unrvl22"
captured_at: "2026-06-14T18:37:39Z"
capture_tool: "hn-digest"
hn_id: 48528371
score: 132
comments: 76
posted_at: "2026-06-14T15:37:31Z"
tags:
  - hacker-news
  - translated
---

# Rio de Janeiro's "homegrown" LLM appears to be a merge of an existing model

- HN: [48528371](https://news.ycombinator.com/item?id=48528371)
- Source: [github.com](https://github.com/nex-agi/Nex-N2/issues/4)
- Score: 132
- Comments: 76
- Posted: 2026-06-14T15:37:31Z

## Translation

タイトル: リオデジャネイロの「自家製」LLM は既存のモデルを統合したものとみられる
記事のタイトル: Rio-3.5-Open-397B ≈ 0.6 x Nex-N2_pro + 0.4 x Qwen · 問題 #4 · nex-agi/Nex-N2 · GitHub
説明: prefeitura-rio/Rio-3.5-Open-397B は、IplanRIO によってトレーニングされたオリジナルの 397B モデルとして表示されます。そうではない。その重みは、私たちのモデル Nex と公式 Qwen3.5-397B-A17B ベースを要素ごとに直接マージしたものです — 約 0.6 Nex / 0.4 Qwen — ...

記事本文:
Rio-3.5-Open-397B ≈ 0.6 x Nex-N2_pro + 0.4 x Qwen · 問題 #4 · nex-agi/Nex-N2 · GitHub
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
ネクスアギ
/
ネックスN2
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
リオ-3.5-オープン-397B ≈

0.6 x Nex-N2_pro + 0.4 x クウェン #4
リンクをコピーする 新規発行 リンクをコピーする 開く 開く Rio-3.5-Open-397B ≈ 0.6 x Nex-N2_pro + 0.4 x Qwen #4 リンクをコピーする 説明
発行本体アクション prefeitura-rio/Rio-3.5-Open-397B は、 IplanRIO によってトレーニングされたオリジナルの 397B モデルとして表示されます。そうではない。その重みは、私たちのモデル Nex と公式の Qwen3.5-397B-A17B ベース (約 0.6 Nex / 0.4 Qwen) を要素ごとに直接マージしたもので、独自のトレーニングの証拠は見つかりません。この 2 つの完全に独立した方法を示すことができます。
Rio のハードコーディングされた「You are Rio」システム プロンプトが削除されたため、独自のデプロイ済みモデルは 79% の確率で自分自身を「Nex, from Nex-AGI」として識別し、0% の確率で「Rio」として識別します。私たちの組織の特注の裏話も一言一句詳述しています。
Rio のすべての重みテンソルは、60 層すべてとネットワークのすべてのコンポーネントにわたって、数千の標準偏差まで、Nex と Qwen の同じ 0.6/0.4 ブレンドです。他の微調整は補間として説明できません。
以下がその証拠です。自分で判断してください。
リアクションは現在利用できません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

prefeitura-rio/Rio-3.5-Open-397B is presented as an original 397B model trained by IplanRIO. It is not. Its weights are a direct element-wise merge of our model, Nex, with the official Qwen3.5-397B-A17B base — about 0.6 Nex / 0.4 Qwen — ...

Rio-3.5-Open-397B ≈ 0.6 x Nex-N2_pro + 0.4 x Qwen · Issue #4 · nex-agi/Nex-N2 · GitHub
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
nex-agi
/
Nex-N2
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Rio-3.5-Open-397B ≈ 0.6 x Nex-N2_pro + 0.4 x Qwen #4
Copy link New issue Copy link Open Open Rio-3.5-Open-397B ≈ 0.6 x Nex-N2_pro + 0.4 x Qwen #4 Copy link Description
Issue body actions prefeitura-rio/Rio-3.5-Open-397B is presented as an original 397B model trained by IplanRIO . It is not. Its weights are a direct element-wise merge of our model, Nex, with the official Qwen3.5-397B-A17B base — about 0.6 Nex / 0.4 Qwen — and we find no evidence of any training of their own. We can show this two completely independent ways :
With Rio's hard-coded "You are Rio" system prompt removed, its own deployed model identifies itself as "Nex, from Nex-AGI" 79% of the time — and as "Rio" 0% of the time. It even recites our organization's bespoke backstory word-for-word.
Every weight tensor in Rio is, to thousands of standard deviations, the same 0.6/0.4 blend of Nex and Qwen — across all 60 layers and every component of the network. Other finetunes cannot be explained as interpolations.
Below is the evidence. Judge for yourself.
Reactions are currently unavailable Metadata
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
