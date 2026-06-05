---
source: "https://github.com/Infinite-Labs-OS/infinite-skills"
hn_url: "https://news.ycombinator.com/item?id=48417952"
title: "OpenAI wrote a guide on how to use /goal mode. I made that guide into a skill"
article_title: "GitHub - Infinite-Labs-OS/infinite-skills: Codex skills from Infinite Labs OS. · GitHub"
author: "RiverXR"
captured_at: "2026-06-05T21:34:41Z"
capture_tool: "hn-digest"
hn_id: 48417952
score: 1
comments: 0
posted_at: "2026-06-05T20:41:35Z"
tags:
  - hacker-news
  - translated
---

# OpenAI wrote a guide on how to use /goal mode. I made that guide into a skill

- HN: [48417952](https://news.ycombinator.com/item?id=48417952)
- Source: [github.com](https://github.com/Infinite-Labs-OS/infinite-skills)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T20:41:35Z

## Translation

タイトル: OpenAI は、/goal モードの使用方法に関するガイドを作成しました。そのガイドをスキルにしました
記事のタイトル: GitHub - Infinite-Labs-OS/infinite-skills: Infinite Labs OS のコーデックス スキル。 · GitHub
説明: Infinite Labs OS の Codex スキル。 GitHub でアカウントを作成して、Infinite-Labs-OS/infinite-skills の開発に貢献してください。

記事本文:
GitHub - Infinite-Labs-OS/infinite-skills: Infinite Labs OS の Codex スキル。 · GitHub
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
Infinite-Labs-OS
/
無限スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
無限

e-Labs-OS/無限スキル
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .codex .codex スキル/目標 スキル/目標 ライセンス ライセンス README.md README.md package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
OpenAI は最近、エージェントが結果に向けて作業できる機能である /goal を開始しました。
彼らのチームは最近、新機能を最大限に活用する方法を説明する記事を公開しました。
https://x.com/dkundel/status/2062650378089594955
記事内で共有されているベストプラクティスをスキル化してみました。
Infinite Skills は、Codex 用のエージェント スキルの小さなコレクションであり、次のように構造化されています。
スーパーパワー: 各スキルはフラット Skills/<skill-name>/SKILL.md ディレクトリに存在します。
コーデックス/ゴールモードのプリフライトインタビュー。大まかな目標を具体的に明確にし、
完全に合成された目標を使用してゴール モードを呼び出す前に、検証可能なコントラクトを作成します。
git clone https://github.com/Infinite-Labs-OS/infinite-skills.git ~ /.codex/infinite-skills
目標スキルをインストールします。
mkdir -p ~ /.codex/skills
ln -s ~ /.codex/infinite-skills/skills/goal ~ /.codex/skills/goal
インストール後に Codex を再起動すると、新しいスキルが検出されます。
インストールに関する詳細なメモは .codex/INSTALL.md にあります。
スキル/
目標/
スキル.md
ライセンス
Infinite Labs OS の Codex スキル。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Codex skills from Infinite Labs OS. Contribute to Infinite-Labs-OS/infinite-skills development by creating an account on GitHub.

GitHub - Infinite-Labs-OS/infinite-skills: Codex skills from Infinite Labs OS. · GitHub
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
Infinite-Labs-OS
/
infinite-skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Infinite-Labs-OS/infinite-skills
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .codex .codex skills/ goal skills/ goal LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
OpenAI recently launched /goal, a function that lets agents work towards an outcomt.
Their team recently published an article explaining how to get the most out of the new feature.
https://x.com/dkundel/status/2062650378089594955
I took the best practises shared article in the article and made it into a skill.
Infinite Skills is a small collection of agent skills for Codex, structured like
Superpowers: each skill lives in a flat skills/<skill-name>/SKILL.md directory.
Preflight interview for Codex /goal mode. It clarifies a rough goal into a concrete,
verifiable contract before invoking goal mode with the full synthesized objective.
git clone https://github.com/Infinite-Labs-OS/infinite-skills.git ~ /.codex/infinite-skills
Install the goal skill:
mkdir -p ~ /.codex/skills
ln -s ~ /.codex/infinite-skills/skills/goal ~ /.codex/skills/goal
Restart Codex after installing so the new skill is discovered.
Detailed install notes are in .codex/INSTALL.md .
skills/
goal/
SKILL.md
License
Codex skills from Infinite Labs OS.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
