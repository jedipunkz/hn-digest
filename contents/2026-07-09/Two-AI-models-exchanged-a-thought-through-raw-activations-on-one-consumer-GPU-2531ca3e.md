---
source: "https://github.com/VitaAI-SCG/one-gpu-lab"
hn_url: "https://news.ycombinator.com/item?id=48852688"
title: "Two AI models exchanged a thought through raw activations on one consumer GPU"
article_title: "GitHub - VitaAI-SCG/one-gpu-lab: ai experiments reproducible on a single consumer GPU · GitHub"
author: "HASHIRAMA1337"
captured_at: "2026-07-09T22:14:17Z"
capture_tool: "hn-digest"
hn_id: 48852688
score: 2
comments: 0
posted_at: "2026-07-09T21:37:02Z"
tags:
  - hacker-news
  - translated
---

# Two AI models exchanged a thought through raw activations on one consumer GPU

- HN: [48852688](https://news.ycombinator.com/item?id=48852688)
- Source: [github.com](https://github.com/VitaAI-SCG/one-gpu-lab)
- Score: 2
- Comments: 0
- Posted: 2026-07-09T21:37:02Z

## Translation

タイトル: 2 つの AI モデルが、1 つのコンシューマ GPU での生のアクティベーションを通じて意見を交換しました
記事タイトル: GitHub - VitaAI-SCG/one-gpu-lab: 単一のコンシューマ GPU で再現可能な AI 実験 · GitHub
説明: 単一のコンシューマ GPU で再現可能な AI 実験 - VitaAI-SCG/one-gpu-lab

記事本文:
GitHub - VitaAI-SCG/one-gpu-lab: 単一のコンシューマ GPU で再現可能な AI 実験 · GitHub
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
VitaAI-SCG
/
ワンGPUラボ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラン

hes タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット エピソード エピソード .gitignore .gitignore ライセンス ライセンス README.md README.md 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
単一のコンシューマ GPU で再現できるフロンティア AI 実験。
ここは人間のオペレーターとAI（クロード）が空き部屋で運営する研究所で、
16 GB グラフィックス カード 1 枚に搭載。私たちは機械の心の内部を研究しています - そして時々
ビットコインの内部 — ガラスボックス手法による: アクティベーションを読み取り、
回路を作成し、コントロールを公開し、障害を保持します。
クラウドクラスターはありません。 API キーはありません。ペイウォールはありません。 1 つのコンシューマー GPU で実行できない場合は、
発送されません。
#
エピソード
TL;DR
02
形態形成: 1 つのボクセルからの体
最大 10,000 のパラメータ ネットワークにより、1 つのシード セルから 3D ボディが成長し、切断された手足が再成長し、コマンドに応じて他のボディに再形成されます。私たちは、1 つのネットワークが多くの物体、つまり 12 チャネルに重ね合わされた N 個の相互に互換性のない座標フレームを保持し、独自に発明した放射状モルフォゲン勾配によって組織化される仕組みを詳しく分析します。 GIF付き。
01
マシンテレパシー
2 つの異なる言語モデルが、生の神経活動を通じて思考を交換しました。トークンやブリッジはなく、精度は 93% でした。残差ストリームのきれいな 2 チャネル法則と、正直なマイナス点です。モデル サイズを越えると、私たちが投げかけたすべての線形ブリッジが無効になります。
さらに多くのエピソードが登場します — バックカタログには、記憶形成を撮影したレイヤーが含まれています
レイヤー、反復的な深さの下での潜在的な推論、およびビットコインマネーグラフ考古学。
GPU が 1 つ。すべての実験は 1 枚の消費者向けカード (当社のカード: RTX 5060 Ti、16 GB) で再現されます。
ガラスの箱。私たちはプロンプトや雰囲気だけでなく、内部情報を読んで操作します。
お祝い前のコントロール。シャッフル コントロール、新鮮なチューニング セット、初期の評価、事前に宣言されたキル ゲート。
ホーン

st ネガティブは結果です。失敗した実験は、成功した実験と同じように慎重に公開されます。
主権者であり、自由です。ローカルのオープンウェイト モデル、無料のデータ ソース、すべてが最初のダウンロード後にオフラインで実行されます。
人間 (オペレーター: 質問を選択し、引き金を引き、評決を所有します)
クロード (共同科学者: 実験の草案を作成し、コードを書き、制御について議論します)。
すべてのエピソードは、そのコラボレーションの実際の作業記録です。台帳は次のとおりです。
実験が行われ、殺害されたことなどすべてが書かれています。
マサチューセッツ工科大学役立つ場合は引用またはリンクしてください。
AI 実験は単一のコンシューマ GPU で再現可能
読み込み中にエラーが発生しました。このページをリロードしてください。
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

ai experiments reproducible on a single consumer GPU - VitaAI-SCG/one-gpu-lab

GitHub - VitaAI-SCG/one-gpu-lab: ai experiments reproducible on a single consumer GPU · GitHub
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
VitaAI-SCG
/
one-gpu-lab
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits episodes episodes .gitignore .gitignore LICENSE LICENSE README.md README.md requirements.txt requirements.txt View all files Repository files navigation
Frontier AI experiments you can reproduce on a single consumer GPU.
This is a research lab run by a human operator and an AI (Claude) out of a spare room,
on one 16 GB graphics card. We study the insides of machine minds — and occasionally
the insides of Bitcoin — with glass-box methods: read the activations, steer the
circuits, publish the controls, keep the failures.
No cloud cluster. No API keys. No paywall. If it doesn't run on one consumer GPU,
it doesn't ship.
#
Episode
TL;DR
02
Morphogenesis: a body from one voxel
A ~10k-parameter network grows a 3D body from one seed cell, regrows severed limbs, and reshapes into other bodies on command. We dissect how one network holds many bodies : N mutually-incompatible coordinate frames superposed in 12 channels, organized by a self-invented radial morphogen gradient. With GIFs.
01
Machine Telepathy
Two different language models exchanged a thought through raw neural activations — no tokens, no bridge, 93% accuracy. A clean two-channel law of the residual stream, and an honest negative: crossing model sizes defeats every linear bridge we threw at it.
More episodes coming — the back-catalog includes memory formation filmed layer by
layer, latent reasoning under recurrent depth, and Bitcoin money-graph archaeology.
One GPU. Every experiment reproduces on a single consumer card (ours: RTX 5060 Ti, 16 GB).
Glass box. We read and manipulate internals, not just prompt and vibe.
Controls before celebration. Shuffle controls, fresh tuning sets, pristine evals, pre-declared kill gates.
Honest negatives are results. Failed experiments get published with the same care as the wins.
Sovereign and free. Local open-weight models, free data sources, everything runs offline after first download.
A human (the operator: picks the questions, pulls the trigger, owns the verdicts) and
Claude (the co-scientist: drafts the experiments, writes the code, argues about controls).
Every episode is the actual working record of that collaboration — the ledgers are
written as the experiments happen, kills and all.
MIT . Cite or link if it's useful to you.
ai experiments reproducible on a single consumer GPU
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
