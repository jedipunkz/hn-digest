---
source: "https://github.com/Mojo0869/ABSL"
hn_url: "https://news.ycombinator.com/item?id=49129422"
title: "Show HN: Integer-only AI complete XOR with 99.7% accuracy all from scratch"
article_title: "GitHub - Mojo0869/ABSL: ABSL or Adaptive Bitshift learning is a lerning method only using Integers for AI i made. The goal is to make a good Integer Neuronr neuron · GitHub"
author: "Mojo_0869"
captured_at: "2026-07-31T22:55:39Z"
capture_tool: "hn-digest"
hn_id: 49129422
score: 1
comments: 0
posted_at: "2026-07-31T22:51:20Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Integer-only AI complete XOR with 99.7% accuracy all from scratch

- HN: [49129422](https://news.ycombinator.com/item?id=49129422)
- Source: [github.com](https://github.com/Mojo0869/ABSL)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T22:51:20Z

## Translation

タイトル: Show HN: 99.7% の精度を持つ整数のみの AI 完全な XOR をすべてゼロから作成
記事タイトル: GitHub - Mojo0869/ABSL: ABSL または Adaptive Bitshift 学習は、私が作成した AI の整数のみを使用した学習方法です。目標は、優れた Integer Neuronr ニューロンを作成することです · GitHub
説明: ABSL または Adaptive Bitshift 学習は、私が作成した AI の整数のみを使用した学習方法です。目標は、優れた Integer Neuronr ニューロンを作成することです - GitHub - Mojo0869/ABSL: ABSL または Adaptive Bitshift 学習は、私が作成した AI の整数のみを使用する学習方法です。目標は、優れた整数ニューロンを作成することです。
[切り捨てられた]

記事本文:
GitHub - Mojo0869/ABSL: ABSL または Adaptive Bitshift 学習は、私が作成した AI の整数のみを使用した学習方法です。目標は、優れた Integer Neuronr ニューロンを作成することです · GitHub
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
モジョ0869
/
ABSL
公共
通知

オン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
87 コミット 87 コミット結果 results src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
##ABSL - 適応型 bitshft 学習 (v.1.0.0 XOR)
ABSL は、Rust で最初から書かれた、実験的なニューラル ネットワーク用の 100% 整数のみの学習アルゴリズムです。
浮動小数点演算を完全に避けることにより、ABSL は FPU を必要としません。そのため、低電力組み込みシステム、8 ビット/16 ビット マイクロコントローラー、ニューロモーフィック ハードウェアにとって興味深いものになります。
ABSL は、固定学習率の代わりに、整数誤差の大きさに基づいた適応ビットシフトを使用して重み更新をスケーリングします。
XOR は通常、連続勾配を使用して解決されます。私は、純粋に整数ベースのアプローチでもそこに到達できるかどうかを確認したかったのです。
最小限の 2-2-1 の代わりに 2-4-1 アーキテクチャ (2 入力、4 つの隠れニューロン、1 出力) を使用すると、整数量子化から生じる丸め誤差を回避するのに十分な冗長性がネットワークに与えられます。
メトリック---------------------- v5
完璧な実行 (4/4)---------- 98.7%
グラオバルの精度----------- 99.7%
平均正しいケース/実行---- 3.99
合計不合格 (</=2 正解) 0.0%
(結果を見てください/最後のディレクトリを見てください)
私は少し怠け者なので、コードについては src/learning_rule/absl.rs を参照してください。
シフト量はエラーと重みのビット長 (ilog2) から計算され、速度を維持しながら重みが爆発するのを防ぎます。
可能な限り最高の実行を試みます: 99% 以上完了 (v5 では全体的な精度 99.7%)
私はモジョ 15 歳、北ドイツに住んでいます。現在 10 年生で、コンピューターと AI が大好きです。
私の目標は、土地について勉強して、良い仕事をして、暮らすことです

日本で
ABSLまたはAdaptive Bitshift学習は、私が作ったAIの整数のみを使った学習手法です。目標は、優れた Integer Neuronr ニューロンを作成することです
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

ABSL or Adaptive Bitshift learning is a lerning method only using Integers for AI i made. The goal is to make a good Integer Neuronr neuron - GitHub - Mojo0869/ABSL: ABSL or Adaptive Bitshift learning is a lerning method only using Integers for AI i made. The goal is to make a good Integer Neuronr n
[truncated]

GitHub - Mojo0869/ABSL: ABSL or Adaptive Bitshift learning is a lerning method only using Integers for AI i made. The goal is to make a good Integer Neuronr neuron · GitHub
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
Mojo0869
/
ABSL
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
87 Commits 87 Commits results results src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
##ABSL - Adaptive bitshft learning ( v.1.0.0 XOR)
ABSL is an experimental, 100% integer-only learning algorithm for neural networks, written from scratch in Rust.
By avoiding floating-point math entirely, ABSL doesn't need an FPU. That makes it interesting for low-power embedded systems, 8-bit/16-bit microcontrollers, and neuromorphic hardware.
Instead of a fixed learning rate, ABSL scales weight updates using an adaptive bit-shift based on the integer error magnitude.
XOR is normally solved with continuous gradients. I wanted to see if a purely integer-based approach could get there too.
Using a 2-4-1 architecture (2 inputs, 4 hidden neurons, 1 output) instead of the minimal 2-2-1 gives the network enough redundancy to work around the rounding errors that come from integer quantization.
Metric---------------------- v5
Perfekt runs (4/4)---------- 98.7%
Glaobal Accuracy------------ 99.7%
Avg. correct cases / run---- 3.99
total failure (</=2 correct) 0.0%
(look in results/ at the last dierctorys)
I'm a bit lazy so look in src/learning_rule/absl.rs for the Code
The shift amount is calculated from the bit-length (ilog2) of the error and weight, which keeps weights from exploding while staying fast.
Try to get the best run possible: 99%+ done (99.7% global accuracy with v5)
I am Mojo 15 years old and live in north germany im currently in 10th grade and love Computers an AI
My Goal is to study an land a good job and live in Japan
ABSL or Adaptive Bitshift learning is a lerning method only using Integers for AI i made. The goal is to make a good Integer Neuronr neuron
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
