---
source: "https://github.com/timothygao8710/minWhisper"
hn_url: "https://news.ycombinator.com/item?id=48423017"
title: "OpenAI Whisper in 150 lines of NumPy"
article_title: "GitHub - timothygao8710/minWhisper · GitHub"
author: "timothygao"
captured_at: "2026-06-06T09:45:04Z"
capture_tool: "hn-digest"
hn_id: 48423017
score: 1
comments: 0
posted_at: "2026-06-06T09:21:25Z"
tags:
  - hacker-news
  - translated
---

# OpenAI Whisper in 150 lines of NumPy

- HN: [48423017](https://news.ycombinator.com/item?id=48423017)
- Source: [github.com](https://github.com/timothygao8710/minWhisper)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T09:21:25Z

## Translation

タイトル: 150 行の NumPy で OpenAI Whisper
記事タイトル: GitHub - timothygao8710/minWhisper · GitHub
説明: GitHub でアカウントを作成して、timothygao8710/minWhisper の開発に貢献します。

記事本文:
GitHub - timothygao8710/minWhisper · GitHub
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
ティモシーガオ8710
/
最小ウィスパー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く

フォルダーとファイル
7 コミット 7 コミット .gitignore .gitignore README.md README.md example.mp3 example.mp3 example.wav example.wav main.py main.py main_kv.py main_kv.py postprocess.py postprocess.py preprocess.py preprocess.py pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
このリポジトリは、Einsum / Einops を使用して、OpenAI Whisper のすべてのフォワード パスを 150 行未満の Numpy で実装します。
圧縮_デモ.mov
KV キャッシュは main.py の上に 7 行あります (O(seq_len ^ 3) -> O(seq_len ^ 2))
Whisper ファミリのあらゆるモデル サイズ、バッチ推論、さまざまなオーディオ形式をサポート
簡潔さを優先するため、layernorm や近似ゲルなどの詳細は、huggingface の実装とはわずかに異なります。
main.py と比較すると、main_kv.py の主な変更点は次のとおりです。
+ kv_cache = {}
+ 名前が kv_cache にない場合:
kv_cache[名前] = np.array([kv_x @ W_k.T, kv_x @ W_v.T + B_v]) # プレフィル
エリフは_カジュアル:
kv_cache[名前], _ = Pack([kv_cache[名前], np.array([kv_x @ W_k.T, kv_x @ W_v.T + B_v])], 'm b * c') # デコード
is_casual = False # 何気ない注意がクロス注意に減少します
クロス アテンション (プレフィルのみ) 用の KV キャッシュを実装し、1 つのクエリでクロス アテンションとして表示することでマスクされたアテンションをデコードします
そしてもちろん
tokens_input, _ = Pack([tokens_input, x[:, -1:]], 'b *') --> tokens_input = x[:, -1:]
新しいトークンごとに発生する「新しい」作業のみを実行することで、実際に複雑さが軽減されるのです
クイックスタート
任意のモデル チェックポイントをダウンロードします。
カール -L -o tiny.pt https://openaipublic.azureedge.net/main/whisper/models/d3dd57d32accea0b295c96e26691aa14d8822fac7d9d27d5dc00b4ca2826dd03/tiny.en.pt
カール -L -o small.pt https://openaipublic.azureedge.net/main/whisper/models/f953ad0fd29cacd07d5a9eda5624af0f6bcf2258be67c92b79389873d91e0872/small.en.pt
カール -L -o med.pt

https://openaipublic.azureedge.net/main/whisper/models/d7440d1dc186f76616474e0ff0b3b6b879abc9d1a4926b7adfa41db2d497ab4f/medium.en.pt
詳細は https://github.com/openai/whisper/blob/main/whisper/__init__.py で入手できます。多言語バージョンでは異なるトークン化が必要であることに注意してください。
preprocess.pyを実行します。これは、正しい入力形式であるメルスペクトグラムとトークン化 (通常はホスト上で行われます) への変換を処理します。また、テンプレート テキスト トークン スキャフォールディングを含む numpy ファイルも生成されます。
ポストプロセスを実行して、モデルの出力トークンを人間が読める形式に非トークン化します (通常はホスト上で行われます)。
MacBook Pro M2 Pro、2023 で実行
tiny.py の Example.wav 出力: <|startoftranscript|><|notimestamps|> 彼らが語る小さな物語は嘘です。ドアにはかんぬきがかけられ、鍵がかけられ、ボルトも締められていました。右の梨は女王様の食卓にぴったりです。丸いカーペットに大きな濡れたシミがあった。凧は沈み、揺れましたが、空中に留まりました。楽しい時間はあっという間に過ぎてしまいます。部屋は人でいっぱいだった
https://cdn.openai.com/papers/whisper.pdf より
https://github.com/huggingface/transformers/blob/main/src/transformers/models/whisper/modeling_whisper.py
https://cdn.openai.com/papers/whisper.pdf
https://jessicastringham.net/2018/01/01/einsum/
https://einops.rocks/1-einops-basics/
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

Contribute to timothygao8710/minWhisper development by creating an account on GitHub.

GitHub - timothygao8710/minWhisper · GitHub
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
timothygao8710
/
minWhisper
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .gitignore .gitignore README.md README.md example.mp3 example.mp3 example.wav example.wav main.py main.py main_kv.py main_kv.py postprocess.py postprocess.py preprocess.py preprocess.py pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
This repo implements all of OpenAI Whisper's forward pass in under 150 lines of Numpy using Einsum / Einops.
compressed_demo.mov
KV cache is 7 lines on top of main.py (O(seq_len ^ 3) -> O(seq_len ^ 2))
Supports any model size in the Whisper family, batched inference, and different audio formats
Details like layernorm and approximate gelu differ slightly from huggingface's implementation to prefer conciseness
Compare to main.py, the key changes in main_kv.py are
+ kv_cache = {}
+ if name not in kv_cache:
kv_cache[name] = np.array([kv_x @ W_k.T, kv_x @ W_v.T + B_v]) # prefill
elif is_casual:
kv_cache[name], _ = pack([kv_cache[name], np.array([kv_x @ W_k.T, kv_x @ W_v.T + B_v])], 'm b * c') # decode
is_casual = False # casual attention reduces to cross attention
Implements KV cache for cross attention (prefill-only), and decode in masked attention by viewing it as cross attention with one query
And of course
tokens_input, _ = pack([tokens_input, x[:, -1:]], 'b *') --> tokens_input = x[:, -1:]
Is what actually buys us the reduction in complexity, by only doing the "new" work incurred for each new token
Quickstart
Download any choice of model checkpoint:
curl -L -o tiny.pt https://openaipublic.azureedge.net/main/whisper/models/d3dd57d32accea0b295c96e26691aa14d8822fac7d9d27d5dc00b4ca2826dd03/tiny.en.pt
curl -L -o small.pt https://openaipublic.azureedge.net/main/whisper/models/f953ad0fd29cacd07d5a9eda5624af0f6bcf2258be67c92b79389873d91e0872/small.en.pt
curl -L -o med.pt https://openaipublic.azureedge.net/main/whisper/models/d7440d1dc186f76616474e0ff0b3b6b879abc9d1a4926b7adfa41db2d497ab4f/medium.en.pt
More are avaliable at: https://github.com/openai/whisper/blob/main/whisper/__init__.py . Note multilingal versions require different tokenization.
Run preprocess.py. This handles converting to mel-spectogram and tokenization (usually done on-host), the correct input format. It will also generate a numpy file containing template text token scaffolding.
Run post-process to detokenize the model's output tokens into human-readable form (usually done on-host)
Ran on MacBook Pro M2 Pro, 2023
Example.wav output on tiny.py: <|startoftranscript|><|notimestamps|> The little tales they tell are false. The door was barred, locked and bolted as well. Right pears are fit for a queen's table. A big wet stain was on the round carpet. The kite dipped and swayed but stayed aloft. The pleasant hours fly by much too soon. The room was crowded with a
From https://cdn.openai.com/papers/whisper.pdf
https://github.com/huggingface/transformers/blob/main/src/transformers/models/whisper/modeling_whisper.py
https://cdn.openai.com/papers/whisper.pdf
https://jessicastringham.net/2018/01/01/einsum/
https://einops.rocks/1-einops-basics/
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
