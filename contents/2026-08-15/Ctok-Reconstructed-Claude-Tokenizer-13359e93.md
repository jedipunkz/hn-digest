---
source: "https://github.com/sanderland/ctok"
hn_url: "https://news.ycombinator.com/item?id=49313865"
title: "Ctok: Reconstructed Claude Tokenizer"
article_title: "GitHub - sanderland/ctok: Claude's Tokenizer, Offline, Kinda · GitHub"
author: "krackers"
captured_at: "2026-08-15T20:11:32Z"
capture_tool: "hn-digest"
hn_id: 49313865
score: 1
comments: 0
posted_at: "2026-08-15T20:10:23Z"
tags:
  - hacker-news
  - translated
---

# Ctok: Reconstructed Claude Tokenizer

- HN: [49313865](https://news.ycombinator.com/item?id=49313865)
- Source: [github.com](https://github.com/sanderland/ctok)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T20:10:23Z

## Translation

タイトル: Ctok: 再構築されたクロード トークナイザー
記事のタイトル: GitHub - Sanderland/ctok: Claude's Tokenizer、オフライン、一種 · GitHub
説明: クロードのトークナイザー、オフラインのようなもの。 GitHub でアカウントを作成して、sanderland/ctok の開発に貢献してください。

記事本文:
GitHub - Sanderland/ctok: Claude の Tokenizer、オフライン、一種 · GitHub
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
サンダーランド
/
クトク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
18 コミット 18 コミット .github/ workflows .github/ workflows ctok ctok testing testing .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイル ナビゲーション
ctok再構築

クロード・トークンは、API 呼び出し、ネットワーク・アクセス、またはランタイムを必要とせずに、オフラインでカウントされます。
依存関係。非公式であり、Anthropic とは関係ありません。
再構築のターゲットは重要です。クロードはトークン境界を公開しないため、tokenize() は戻り値を返します。
Anthropic の正確なセグメンテーションに関する主張ではなく、1 つの有効な最小コスト タイリングです。背後にある研究
このモデルは、「On the biology of Claude's tokenizer」で説明されています。
from ctok import token_count 、トークン化
token_count ( "hello, world" ) # 10、v3 ファミリを使用
token_count ( "hello, world" , "4.7" ) # 15
token_count ( "hello, world" , "5.0" ) # 10
tokens = tokenize ( "NASA はトークナイザーが好き" )
assert len ( tokens ) == token_count ( "NASA はトークナイザーが好き" )
コマンドライン インターフェイスは、マークされたストリームとそのタイリングを出力します。
ctok「ハロー、ワールド」
サポートされている家族
version は常に文字列です。例: "4.7" 、コンポーネントごとに比較 — "4.10" はその後にソート
「4.9」。「4.2」未満ではありません。 float ではその区別ができません (Python ではリテラルが折りたたまれます)
ここのコードがそれを認識する前に 4.10 から 4.1 まで)、そのため、非 str バージョンでは TypeError が発生します。
v5 は、固定オーバーヘッドが若干異なる v4.7 です。
NFC やファミリー固有の引用符の折りたたみなど、テキストを正規化します。
単語、大文字と小文字、バイトのマーカーを使用してストリームに書き換えます。
測定された語彙と UTF-8 バイト フォールバックに対する最小コストのタイリングを見つけます。
測定されたメッセージ フレームを追加します。
token_count(text) は len(tokenize(text)) です。出力表記により、内部構造が表示されます。
これらの結果は、ctok と記録された count_tokens 応答を比較します。
v5 は同じ語彙を使用しているため、v4.7 と同じコンテンツ結果になります。
保存された測定セットにはアンダーカウントは含まれません: 1,664,940 個の v3 テキストのうち 0 個と 1,722,961 個のうち 0 個
v4.7のテキスト。これは経験的な結果であり、任意の入力を保証するものではありません。ゴル

dfish、ロゼッタ、
(最後の 6 つの部分が二等分して選択されたため) UDHR が候補を選択する場合があります。マルチPL-E
決してそうなることはなく、このテーブルに残っている 1 つの保留されたコーパスです。
UV で pytest を実行
uv 実行 python テスト/gates.py --markdown
語彙の証拠
2 つの語彙ファイルには、48,645 個の v3 部分と 15,240 個の v4.7 部分が含まれています。すべてのエントリには固定値があります
メンバーシップ証人、またはテストスイートによってチェックされた構造マーカー原子の 1 つです。
ctokからの輸入品、証人
レン ( 個 ( "4.7" )) # 15240
証人 ( "⟨bow⟩the⟨eow⟩" 、 "4.7" )
# {'プローブ': 'the', 'raw': 12, 'kind': 'raw'}
目撃者によると、マーク付きの部品 1 個につき、校正済みプローブのトークン 1 個の費用がかかるそうです。を証明するものではありません
エンコーダは、等コスト タイル間の関係を書き換えるか、解決します。 testing/test_witness.py はすべてのチェックを行います
公開された証人であり、完全な目撃取材または特別取材が必要です。
クロードのトークナイザー、オフライン、ちょっと
Readme MIT ライセンス アクティビティ スター
3 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude's Tokenizer, Offline, Kinda. Contribute to sanderland/ctok development by creating an account on GitHub.

GitHub - sanderland/ctok: Claude's Tokenizer, Offline, Kinda · GitHub
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
sanderland
/
ctok
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
18 Commits 18 Commits .github/ workflows .github/ workflows ctok ctok tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
ctok reconstructs Claude token counts offline, with no API call, network access, or runtime
dependencies. It is unofficial and is not affiliated with Anthropic.
The reconstruction targets counts. Claude does not expose token boundaries, so tokenize() returns
one valid minimum-cost tiling, not a claim about Anthropic's exact segmentation. The research behind
the model is described in On the biology of Claude's tokenizer .
from ctok import token_count , tokenize
token_count ( "hello, world" ) # 10, using the v3 family
token_count ( "hello, world" , "4.7" ) # 15
token_count ( "hello, world" , "5.0" ) # 10
tokens = tokenize ( "NASA likes tokenizers" )
assert len ( tokens ) == token_count ( "NASA likes tokenizers" )
The command-line interface prints the marked stream and its tiling:
ctok " hello, world "
Supported families
version is always a string, e.g. "4.7" , compared component by component — "4.10" sorts after
"4.9" , not below "4.2" . A float can't make that distinction (Python collapses the literal
4.10 to 4.1 before any code here sees it), so a non- str version raises TypeError .
v5 is v4.7 with a slightly different fixed overhead.
normalizes the text, including NFC and family-specific quote folding;
rewrites it into a stream with word, case, and byte markers;
finds a minimum-cost tiling over the measured vocabulary and UTF-8 byte fallback;
adds the measured message frame.
token_count(text) is len(tokenize(text)) . The output notation makes internal structure visible:
These results compare ctok with recorded count_tokens responses:
v5 has the same content result as v4.7 because it uses the same vocabulary.
The stored measurement sets contain no under-counts: 0 of 1,664,940 v3 texts and 0 of 1,722,961
v4.7 texts. This is an empirical result, not a guarantee for arbitrary input. Goldfish, Rosetta and
(since its final six pieces were selected by bisecting it) UDHR may select candidates; MultiPL-E
never does, and is the one remaining held-out corpus in this table.
uv run pytest
uv run python tests/gates.py --markdown
Vocabulary evidence
The two vocabulary files contain 48,645 v3 pieces and 15,240 v4.7 pieces. Every entry has a fixed
membership witness or is one of the structural marker atoms checked by the test suite.
from ctok import pieces , witness
len ( pieces ( "4.7" )) # 15240
witness ( "⟨bow⟩the⟨eow⟩" , "4.7" )
# {'probe': 'the', 'raw': 12, 'kind': 'raw'}
A witness says that one marked piece costs one token in a calibrated probe. It does not prove the
encoder rewrite or resolve ties between equal-cost tilings. tests/test_witness.py checks every
published witness and requires complete witnessed-or-special coverage.
Claude's Tokenizer, Offline, Kinda
Readme MIT license Activity Stars
3 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
