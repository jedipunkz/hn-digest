---
source: "https://github.com/Vector-Compute-Engine/api-quickstart"
hn_url: "https://news.ycombinator.com/item?id=49327240"
title: "OpenAI-Compatible API Quickstart in Python, JavaScript, and PowerShell"
article_title: "GitHub - Vector-Compute-Engine/api-quickstart: Quickstart examples for the official OpenAI-compatible API · GitHub"
image: "https://opengraph.githubassets.com/9c094448c52f8569ad2b70149fdf6a435c8e14479afcd4f2e8149b3ac386ba5a/Vector-Compute-Engine/api-quickstart"
author: "VectorEngine"
captured_at: "2026-08-17T07:44:08Z"
capture_tool: "hn-digest"
hn_id: 49327240
score: 2
comments: 0
posted_at: "2026-08-17T06:46:21Z"
tags:
  - hacker-news
  - translated
---

# OpenAI-Compatible API Quickstart in Python, JavaScript, and PowerShell

- HN: [49327240](https://news.ycombinator.com/item?id=49327240)
- Source: [github.com](https://github.com/Vector-Compute-Engine/api-quickstart)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T06:46:21Z

## Translation

タイトル: Python、JavaScript、PowerShell での OpenAI 互換 API クイックスタート
記事タイトル: GitHub - Vector-Compute-Engine/api-quickstart: 公式 OpenAI 互換 API のクイックスタート例 · GitHub
説明: 公式 OpenAI 互換 API のクイックスタート例 - Vector-Compute-Engine/api-quickstart

記事本文:
GitHub - Vector-Compute-Engine/api-quickstart: 公式 OpenAI 互換 API のクイックスタート サンプル · GitHub
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
ベクトル計算エンジン
/
APIクイックスタート
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット docs docs 例 例 .env.example .env.example .gitignore .gitignore CHANGELOG.md CHA

NGELOG.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
このリポジトリでは、VectorNode を通じて最初の OpenAI 互換 API リクエストを行う方法を示します。
https://www.vectronode.com/register
VectorNode コンソールにサインインし、API キーを作成します。
完全な API キーを GitHub、スクリーンショット、記事、または Discord で公開しないでください。
プロジェクト フォルダーで PowerShell を開き、次を実行します。
py -m venv .venv
.\.venv\Scripts\ python.exe - m pip install - rrequirements.txt
コピー項目 .env.example .env
4. APIキーを設定する
.env ファイルを開き、YOUR_API_KEY を独自の API キーに置き換えます。
API_KEY=あなたの_API_KEY
BASE_URL=https://www.vectronode.com/v1
MODEL=gpt-4o
VectorNode アカウントで gpt-4o が利用できない場合は、gpt-4o を置き換えてください。現在のモデル リストまたはコンソールに表示されているモデル名を使用します。
実際の API キーを .env.example に入れないでください。
.\.venv\Scripts\ python.exe example\python\chat.py
リクエストが成功すると、PowerShell でモデル応答が出力されます。
401 : API キーが正しく、有効であるかどうかを確認します。
model_not_found : アカウントで現在利用可能なモデルを選択します。
429 : 待ってから再試行するか、アカウント残高とレート制限を確認してください。
タイムアウト: ネットワークを確認し、後でもう一度試してください。
https://www.vectronode.com/customersupport
パスワード、確認コード、支払い詳細、完全な API キーを決して共有しないでください。
公式の OpenAI 互換 API のクイックスタート例
Readme MIT ライセンス セキュリティ ポリシー
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Quickstart examples for the official OpenAI-compatible API - Vector-Compute-Engine/api-quickstart

GitHub - Vector-Compute-Engine/api-quickstart: Quickstart examples for the official OpenAI-compatible API · GitHub
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
Vector-Compute-Engine
/
api-quickstart
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits docs docs examples examples .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md requirements.txt requirements.txt View all files Repository files navigation
This repository shows how to make your first OpenAI-compatible API request through VectorNode.
https://www.vectronode.com/register
Sign in to the VectorNode console and create an API key.
Never publish your complete API key in GitHub, screenshots, articles, or Discord.
Open PowerShell in the project folder and run:
py - m venv .venv
.\.venv\Scripts\ python.exe - m pip install - r requirements.txt
Copy-Item .env.example .env
4. Configure the API key
Open the .env file and replace YOUR_API_KEY with your own API key.
API_KEY=YOUR_API_KEY
BASE_URL=https://www.vectronode.com/v1
MODEL=gpt-4o
Replace gpt-4o if it is not available in your VectorNode account. Use a model name shown in the current model list or console.
Do not put your real API key in .env.example .
.\.venv\Scripts\ python.exe examples\python\chat.py
A successful request prints a model response in PowerShell.
401 : Check whether the API key is correct and active.
model_not_found : Select a model currently available in your account.
429 : Wait and try again, or check the account balance and rate limits.
Timeout: Check the network and try again later.
https://www.vectronode.com/customersupport
Never share your password, verification code, payment details, or complete API key.
Quickstart examples for the official OpenAI-compatible API
Readme MIT license Security policy
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
