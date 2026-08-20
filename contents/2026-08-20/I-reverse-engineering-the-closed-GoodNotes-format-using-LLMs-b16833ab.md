---
source: "https://github.com/Kaih1825/parser-for-goodnotes"
hn_url: "https://news.ycombinator.com/item?id=49377664"
title: "I reverse-engineering the closed GoodNotes format using LLMs"
article_title: "GitHub - Kaih1825/parser-for-goodnotes: The document parser for goodnotes · GitHub"
image: "https://opengraph.githubassets.com/884e9c3cc62eade1c741c76affc4e07bc6e0b2e9aa904232e90b84ff22e6d9fe/Kaih1825/parser-for-goodnotes"
author: "Kaih1825"
captured_at: "2026-08-20T18:24:02Z"
capture_tool: "hn-digest"
hn_id: 49377664
score: 2
comments: 1
posted_at: "2026-08-20T17:33:47Z"
tags:
  - hacker-news
  - translated
---

# I reverse-engineering the closed GoodNotes format using LLMs

- HN: [49377664](https://news.ycombinator.com/item?id=49377664)
- Source: [github.com](https://github.com/Kaih1825/parser-for-goodnotes)
- Score: 2
- Comments: 1
- Posted: 2026-08-20T17:33:47Z

## Translation

タイトル: LLM を使用してクローズド GoodNotes フォーマットをリバースエンジニアリングしてみました
記事のタイトル: GitHub - Kaih1825/parser-for-goodnotes: Goodnotes 用のドキュメント パーサー · GitHub
説明: Goodnotes 用のドキュメント パーサー。 GitHub でアカウントを作成して、Kaih1825/parser-for-goodnotes の開発に貢献してください。

記事本文:
GitHub - Kaih1825/parser-for-goodnotes: Goodnotes 用のドキュメント パーサー · GitHub
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
カイ1825
/
Goodnote 用のパーサー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
55 コミット 55 コミット フォルダーとファイル
.github .github アセット アセット スクリプト スクリプト src/goodnotes_re src/goodnotes_re テスト テスト Web Web ウィキ Wiki .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md LEGAL-NOTICE.md

LEGAL-NOTICE.md ライセンス ライセンス README.md README.md cli.md cli.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Vibe コーディングの免責事項: このプロジェクト全体は、Vibe コーディング (AI 支援の高速ペア プログラミングおよび探索的開発) を通じて開発されました。パーサーはテスト コーパスに対して検証されていますが、コードとアーキテクチャの選択には実験的な AI 主導の反復スタイルが反映されています。ご自身の判断でご使用ください！
ユーザーが提供した GoodNotes 5 および 6 の .goodnotes アーカイブを検査および解析するための、独立したオープンソースの完全に型指定された Python ツールキット。 protobuf ワイヤ形式を直接デコードし、Apple LZ4 フレーム化ストリームを解析し、Troy Hanson TPL メモリ画像をデコードし、観察された RGBA ストローク データを抽出し、ドキュメントを JSON および SVG にエクスポートします。
このプロジェクトは、Goodnotes Limited と提携、承認、後援、または正式に関係しているものではありません。リリースおよび使用上の注意については、LEGAL-NOTICE.md を参照してください。
ヒューリスティックな float スキャンを意図的に使用しません。
ソースアーカイブ
GoodNotes オリジナル ( .jpg )
このプロジェクトの SVG エクスポート ( .svg )
例 1: 手書きの数式と画像
( ex1.goodnotes )
例 2: ブラシ スタイルとストロークのバリエーション
( ex2.goodnotes )
例 3: 多層画像と中国語テキスト
( ex3.goodnotes )
モチベーションと内部構造
GoodNotes は素晴らしいメモアプリですが、その閉鎖的なエコシステムは常に悩みの種でした。ベクター ストロークを編集可能な状態に保ちながらノートブックをエクスポートしたい場合は、かなり運が悪いです。独自の .goodnotes ファイル形式を使用するか、編集機能を失ったフラット化された PDF にエクスポートする必要があります。これを解決するために、.goodnotes ファイルをデコードできるオープンソースのパーサーを構築しました。
私にはリバース エンジニアリングの知識がまったくなく、この形式をデコードするのは簡単ではありませんでした (

これに取り組んで多くの成功したプロジェクトがあります)。そこで、LLM (主に Gemini 3.1 Pro、Gemini 3.6 Flash、および Claude Sonnet 5) を使用した「バイブ コーディング」によってこれを完全に構築しました。
内部で解析がどのように行われるかは次のとおりです。
ZIP と Protobuf: いくつかの分析の結果、.goodnotes は本質的に ZIP アーカイブであることがわかりました。メインのストローク データは、シリアル化された Protobuf ファイルとして、notes/ ディレクトリにページごとに保存されます。公式の .proto スキーマを持っていなかったため、プロジェクトは基盤となる Wire Format を介して Protobuf を盲目的に解析し、抽象構文ツリー (AST) を構築します。
Apple LZ4: 次に、一部のデータ フィールドが bv41 または bv4- で始まり、 bv4$ で終わることがわかりました。これは、Apple 独自の framed LZ4 圧縮の署名です。 64KB のスライディング履歴ウィンドウを維持し、ビットごとの操作を使用して LZ4 トークンを処理することで、これを正常に解凍しました。
Troy Hanson の TPL: 解凍された平文はマジック バイト tpl\0 で始まり、Troy Hanson の TPL 形式 (C シリアル化ライブラリ) であることがわかります。フォーマット文字列を推測することで、圧力値を含む離散点で構成される生のストローク データを抽出することができました。
SVG リボン: 最後に、自然な可変幅でストロークをレンダリングするために、パーサーは隣接する点間の法線ベクトルを計算し、平滑化のためにスライド平均を適用します。次に、圧力値に基づいてエッジを外側に押し出し、離散点を閉じた SVG ポリゴン リボンに縫い合わせます。
現在、プロジェクトは .goodnotes アーカイブ内のバイナリ ファイルを解析してストローク、テキスト、その他の要素を抽出し、それらを .svg または .pdf に直接エクスポートできます。 GoodNotes アプリ自体からの標準エクスポートとは異なり、このパーサーは生データのロックを完全に解除します。最終的な目標は、他のオープンなファイルへの変換を可能にすることです。

ベクトル形式 (InkML など) を使用できるため、ユーザーは他のアプリに移行でき、労力が単一のベンダーに固定されるのを防ぐことができます。
解析原則の詳細については、GitHub Wiki を参照してください。
Protobuf Wire Decoder : 未知のフィールドを保持しながら、フレーム化されたおよびフレーム化されていない protobuf メッセージをロスレスで解析します。
Apple LZ4 & Troy Hanson TPL Decoder : bv41 LZ4 ストリームを解凍し、埋め込まれた TPL 形式の文字列 ( vuA(v)A(S(uu))... ) を構造化されたストローク ポイントと可変幅のリボンにデコードします。
ストロークとカラー解析: bv4$ の後に protobuf トレーラーを直接デコードして、正確な RGBA カラーとハイライターの透明度を抽出します。
ストローク サポート : テスト コーパス内で観察される点、線、曲線、ペン ツール、蛍光ペン、消しゴム、図形、および移動/コピーされた要素の増加するセットを解析します。
ページと背景の解像度 : PDF の背景/MediaBox の寸法の自動検出 (A4、レター、横向きと縦向き)、ページの順序付け、テキストの断片、付箋 (便條紙) のコンテンツ抽出。
CLI ツール: gn-inspect 、 gn-dump 、 gn-diff 、 gn-export-json 、 gn-export-svg 、 gn-export-pdf 。
✅ 動作/完全サポート
万年筆、ボールペン、筆ペン、蛍光ペン：カスタムカラーと線の太さを完全サポート。
付箋 : 付箋の内容と書式。
オートシェイプとシェイプ : ベクターシェイプのレンダリング。
消しゴム : 消しゴムのストロークとラインカット。
なげなわツール : 要素を変換、コピー、または移動します。
PDF-backed Notebooks : 寸法付きの PDF 背景ページ ( /MediaBox )。
Pencil : 出力は表示されますが、圧力感度が正しくなく、傾き感度がありません。
矢印 : レンダリング出力は現在非常に不安定です。
ステッカー / 要素 : 特定のベクター ライン/ストロークはエクスポートに失敗する場合があります。
音声録音 : まだ実装されていません。
UV同期
または標準の pip インストール:
pip install -e 。

以下を使用して単体テスト スイートを実行します。
UV で pytest を実行
注:samples/ は意図的に無視されており、公開ソース ツリーの一部ではありません。再配布の許可がない限り、.goodnotes ファイルまたは抽出されたアセットを公開しないでください。
# アーカイブ インベントリと sha256 チェックサムを検査します
gn-inspect サンプル.goodnotes
# 任意の protobuf メンバーの JSON へのロスレス ダンプ
gn-dump サンプル.goodnotes インデックス.notes.pb
# 2 つの .goodnotes アーカイブを比較します
gn-diff before.goodnotes after.goodnotes
# ドキュメント全体、メタデータ、ページ、ストローク、生のワイヤー ツリーを JSON にエクスポート
gn-export-json サンプル.goodnotes -o document.json
# 正確なストロークのリボン、色、寸法を含むベクター SVG ページをエクスポート
gn-export-svg サンプル.goodnotes -o Pages-svg
# ベクター SVG をエクスポートし、すべてのページを PDF にパッケージ化する
gn-export-svg サンプル.goodnotes -o Pages-svg --pdf
# 複数ページの PDF ドキュメントを直接エクスポートする
gn-export-pdf サンプル.goodnotes -o document.pdf
またはモジュール呼び出し経由:
PYTHONPATH=src python3 -m Goodnotes_re.cli export-svg Sample.goodnotes -o Pages-svg
完全な CLI リファレンス (すべてのフラグとバッチ エクスポートの例を含む) については、 cli.md を参照してください。
GoodNotes_re から GoodNotesDocument をインポート
GoodNotesDocument を使用して。 ( "sample.goodnotes" ) を doc として開きます。
# インベントリ
メンバー = ドキュメント 。在庫（）
# ストローク、寸法、テキストを含む文書ページ
ページ = ドキュメント 。ページ ()
ページ内のページの場合:
print (f"ページ {ページ .インデックス + 1} : {ページ .寸法 .幅 } x {ページ .寸法 .高さ } pt" )
print ( f"ストローク: { len ( page . ストローク ) } " )
ページ内のストロークの場合。ストローク:
print (f" Stroke { ストローク . uuid } : color { ストローク . color_hex } , alpha { ストローク . アルファ } , ポイント { len (ストローク . ポイント ) } " )
# テキストと付箋のコンテンツ
フラグメント = ドキュメント 。テキストフラグメント ()
フラグメント内のフラグメントの場合:
print ( f"[ { frag .source_path } ] {

フラグ。フォーマット } : { フラグ .テキスト } " )
# フォーマット分析のための構造的なページ要素の要約
ページ内のページの場合:
ページ内の要素の場合。要素:
print (要素 . kind 、要素 . uuid 、要素 .attachment_uuid 、要素 . RELATED_UUIDS )
ドキュメント
文書
説明
cli.md
完全な CLI リファレンス
GitHub ウィキ
詳細な技術文書 (アーキテクチャ、フォーマット、レンダリング)
法的通知.md
法的、商標、プライバシー、および再配布に関する通知
「貢献ガイド」も参照してください。
法的および商標に関する通知 : 「Goodnotes」および関連する名前、ロゴ、およびマークは、Goodnotes Limited の所有物です。 Document Parser for GoodNotes は、コミュニティによって開発された独立したプロジェクトであり、Goodnotes Limited と提携、承認、後援、または正式に関係しているわけではありません。法的、商標、プライバシー、および再配布の詳細については、 LEGAL-NOTICE.md をお読みください。
Vibecoding免責聲明 ：本案全面採用 Vibecoding（AI助快速結對程式設計與探索式開發）進行構建。驅動的實性疊代風格。請自行評估並謹慎使用！
一つのセキュリティ、オープンソース、完全型に特化した Python ツール群、使用者が提供する GoodNotes 5 と GoodNotes 6 .goodnotes は保存されています。筆跡資料，並將文件匯出為 JSON 與 SVG。
この計画案 Goodnotes Limited には、あらゆる提携関係、バックブック、ヘルプまたは官公庁関連の活動が含まれています。
原始封存
GoodNotes 原版渲染 ( .jpg )
本專案 SVG 匯出 ( .svg )
例 1：手寫公式與插圖
( ex1.goodnotes )
例2

：多彩なブラシとカラー手書き
(例2.goodnotes)
例 3: マルチレイヤーグラフィックスとテキストオーバーレイ、および中国語の手書き
(例3.goodnotes)
オープンエンジンと分析原理
GoodNotes は優れたメモ作成ソフトウェアですが、その閉鎖的な性質により、ベクター手書きを保存するには、閉じた .goodnotes 専用ファイル、または編集機能を失う PDF にエクスポートする以外にほとんどオプションがありません。そこで、私は .goodnotes ファイルを解析する解析ツールを開発しました。
私にはこのタイプのプロジェクトの経験がなく、この .goodnotes を解析するのは簡単ではないため (これまで解析に成功したプロジェクトはそれほど多くありません)、このプロジェクトを完了するためにバイブ コーディング (主に Gemini 3.1 Pro、Gemini 3.6 Flash、および Claude Sonnet 5) を使用しました。
圧縮と Protobuf のブラインド解釈: 分析の結果、.goodnotes は本質的に圧縮ファイルであり、主な手書き情報はページ単位で Notes/ に保存され、各ページの手書き情報は Protobuf シリアル化ファイルとして保存されることがわかりました。このプロジェクトでは、Wire Format メソッドを使用して抽象構文ツリーを解析および構築します。
Apple LZ4 リバース エンジニアリング: 次に、一部のフィールドが bv41 または bv4- で始まり bv4$ で終わっていることがわかりました。これは、これが Apple LZ4 によって圧縮されたデータであることを意味します。私たちは 64KB の履歴スライディング ウィンドウを維持し、ビット操作によって LZ4 トークンを処理して正常に解凍しました。
TPL メモリ イメージ: 解凍された平文は tpl\0 で始まるため、これが Troy Hanson の TPL 形式であることがわかり、彼の手書き情報 (筆跡感度を持つ離散点) が形式文字列を通じて推定されます。
ベクトル ジオメトリの再構築: 次に、隣接する 2 つの点間の法線ベクトルを計算し、圧力感度値の外側への移動と組み合わせてスライド平均平滑化を実行することにより、離散点を閉じた多角形にステッチすることができ、最終出力は太さが自然に変化した SVG ベクトル手書きになります。
現在、プロジェクトは .goodnotes 内のバイナリ ファイルを分析して手書き文字、テキスト、その他の情報を取得し、それを .svg に出力できます。

または .pdf 形式。 GoodNotes アプリから直接エクスポートする場合とは異なり、このプロジェクトは .goodnotes 形式を解析するため、解析されたデータは将来的に他のオープン ソース形式 (InkML など) に変換できるため、ユーザーの努力が単一のメーカーによってハイジャックされることはなくなります。
Protobuf Wire Decoder: フレーム化された protobuf メッセージとフレーム化されていない protobuf メッセージをロスレスに解析し、未知のフィールドを保持します。
Apple LZ4 および Troy Hanson TPL デコーダ: bv41 LZ4 ストリームを解凍し、埋め込まれた TPL 形式の文字列を構造化された手書きポイントと可変幅の手書きストリップにデコードします。
手書きと色の解析: bv4$ post-protobuf トレーラーを解析して、正確な RGBA カラーとハイライターの透明度をキャプチャします。
手書きおよび幾何学的コンポーネントのサポート: テスト サンプルで観察された単一の点、線、曲線、ペン/ボール ペン ツール、蛍光ペン、消しゴムの切り抜き、グラフィック、および移動/コピー要素を解析します。
ページと背景の分析: PDF/MediaBox サイズ (A4、レター、横向き、縦向き)、ページ順序、テキストの断片、メモ用紙の内容を自動的に検出します。
CLI ツール: gn-inspect、gn-dump、gn-diff、gn-export-json、gn-export-svg、gn-export-pdf。
✅ 現在テストは正常に出力できています
ペン、ボールペン、絵筆、蛍光ペン: さまざまな色や太さに対応できます。
UV で pytest を実行
注:samples/ は意図的に無視されており、パブリック ソース ツリーには属しません。再公開許可がない限り、.goodnotes ファイルまたは抽出されたアセットを公開しないでください。
# アーカイブ ディレクトリのリストを表示し、sha256 コードをチェックします
gn-inspect サンプル.goodnotes
# protobuf メンバーをロスレスで JSON に出力します
gn-dump サンプル.goodnotes インデックス.notes.pb
# 2 つの .goodnotes アーカイブの違いを比較する
gn-diff before.goodnotes after.goodnotes
# ドキュメント全体、メタデータ、ページ、手書き文字、およびオリジナルのワイヤをエクスポートします

[切り捨てられた]

## Original Extract

The document parser for goodnotes. Contribute to Kaih1825/parser-for-goodnotes development by creating an account on GitHub.

GitHub - Kaih1825/parser-for-goodnotes: The document parser for goodnotes · GitHub
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
Kaih1825
/
parser-for-goodnotes
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
55 Commits 55 Commits Folders and files
.github .github assets assets scripts scripts src/ goodnotes_re src/ goodnotes_re tests tests web web wiki wiki .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LEGAL-NOTICE.md LEGAL-NOTICE.md LICENSE LICENSE README.md README.md cli.md cli.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Vibe Coding Disclaimer : This entire project was developed through Vibe Coding (AI-assisted rapid pair-programming and exploratory development). While the parser has been verified against test corpora, code and architecture choices reflect an experimental AI-driven iteration style. Use at your own discretion!
An independent, open-source, fully typed Python toolkit for inspecting and parsing user-supplied GoodNotes 5 and 6 .goodnotes archives. It decodes protobuf wire format directly, parses Apple LZ4 framed streams, decodes Troy Hanson TPL memory images, extracts observed RGBA stroke data, and exports documents to JSON and SVG.
This project is not affiliated with, endorsed by, sponsored by, or officially connected to Goodnotes Limited. See LEGAL-NOTICE.md for release and usage notes.
It deliberately does NOT use heuristic float scanning .
Source Archive
GoodNotes Original ( .jpg )
This Project SVG Export ( .svg )
Example 1: Handwritten Formulas & Images
( ex1.goodnotes )
Example 2: Brush Styles & Stroke Variations
( ex2.goodnotes )
Example 3: Multi-Layer Images & Chinese Text
( ex3.goodnotes )
Motivation & Under the Hood
GoodNotes is an amazing note-taking app, but its closed ecosystem has always been a pain point. If you want to export your notebooks while keeping the vector strokes editable, you're pretty much out of luck—you either have to stick with the proprietary .goodnotes file format or export to a flattened PDF that loses editability. To solve this, I built an open-source parser that can decode .goodnotes files.
I have zero background in reverse engineering, and decoding this format wasn't easy (I haven't seen many successful projects tackling this). So, I built this entirely through "vibe coding" using LLMs—primarily Gemini 3.1 Pro, Gemini 3.6 Flash, and Claude Sonnet 5.
Here is how the parsing works under the hood:
ZIP & Protobuf: After some analysis, it turns out that .goodnotes is essentially a ZIP archive. The main stroke data is stored page by page in the notes/ directory as serialized Protobuf files. Since I didn't have the official .proto schemas, the project blindly parses the Protobuf via the underlying Wire Format to construct an Abstract Syntax Tree (AST).
Apple LZ4: Next, we discovered that some data fields start with bv41 or bv4- and end with bv4$ . This is a signature for Apple's proprietary Framed LZ4 compression. We successfully decompressed this by maintaining a 64KB sliding history window and using bitwise operations to handle the LZ4 tokens.
Troy Hanson's TPL: The decompressed plaintext starts with the magic bytes tpl\0 , revealing it as Troy Hanson's TPL format (a C serialization library). By inferring the format strings, we were able to extract the raw stroke data—which consists of discrete points containing pressure values.
SVG Ribbons: Finally, to render the strokes with natural variable widths, the parser calculates the normal vectors between adjacent points and applies a sliding average for smoothing. It then pushes the edges outward based on the pressure values, stitching the discrete points into a closed SVG polygon ribbon.
Currently, the project can parse the binary files inside a .goodnotes archive to extract strokes, text, and other elements, exporting them directly to .svg or .pdf . Unlike standard exports from the GoodNotes app itself, this parser completely unlocks the raw data. The ultimate goal is to allow conversions to other open vector formats (like InkML) so users can migrate to other apps and prevent their hard work from being locked in by a single vendor.
For more detailed parsing principles, please refer to the GitHub Wiki .
Protobuf Wire Decoder : Lossless parsing of framed and unframed protobuf messages, preserving unknown fields.
Apple LZ4 & Troy Hanson TPL Decoder : Decompresses bv41 LZ4 streams and decodes embedded TPL format strings ( vuA(v)A(S(uu))... ) into structured stroke points and variable-width ribbons.
Stroke & Color Parsing : Direct protobuf trailer decoding after bv4$ to extract exact RGBA colors and highlighter transparency.
Stroke Support : Parses a growing set of dots, lines, curves, pen tools, highlighters, erasers, shapes, and moved/copied elements observed in the test corpus.
Page & Background Resolution : Automatic PDF background /MediaBox dimension detection (A4, Letter, landscape vs. portrait), page ordering, text fragments, and sticky notes (便條紙) content extraction.
CLI Tools : gn-inspect , gn-dump , gn-diff , gn-export-json , gn-export-svg , gn-export-pdf .
✅ Working / Fully Supported
Fountain Pen, Ballpoint Pen, Brush Pen, Highlighter : Full support with custom colors and line thicknesses.
Sticky Notes : Sticky note contents and formatting.
Auto-Shapes & Shapes : Vector shape rendering.
Eraser : Eraser strokes and line cuts.
Lasso Tool : Transformed, copied, or moved elements.
PDF-backed Notebooks : PDF background pages with dimensions ( /MediaBox ).
Pencil : Visible output, but pressure sensitivity is incorrect and tilt sensitivity is missing.
Arrows : Render output is currently highly unstable.
Stickers / Elements : Certain vector lines/strokes may fail to export.
Audio Recordings : Not yet implemented.
uv sync
Or standard pip install:
pip install -e .
Run the unit test suite with:
uv run pytest
Note: samples/ is intentionally ignored and is not part of the public source tree. Do not publish .goodnotes files or extracted assets unless you have permission to redistribute them.
# Inspect archive inventory and sha256 checksums
gn-inspect sample.goodnotes
# Lossless dump of any protobuf member to JSON
gn-dump sample.goodnotes index.notes.pb
# Diff two .goodnotes archives
gn-diff before.goodnotes after.goodnotes
# Export entire document, metadata, pages, strokes, and raw wire trees to JSON
gn-export-json sample.goodnotes -o document.json
# Export vector SVG pages with exact stroke ribbons, colors, and dimensions
gn-export-svg sample.goodnotes -o pages-svg
# Export vector SVGs and package all pages into a PDF
gn-export-svg sample.goodnotes -o pages-svg --pdf
# Directly export multi-page PDF document
gn-export-pdf sample.goodnotes -o document.pdf
Or via module invocation:
PYTHONPATH=src python3 -m goodnotes_re.cli export-svg sample.goodnotes -o pages-svg
For the full CLI reference (including all flags and batch export examples), see cli.md .
from goodnotes_re import GoodNotesDocument
with GoodNotesDocument . open ( "sample.goodnotes" ) as doc :
# Inventory
members = doc . inventory ()
# Document pages with strokes, dimensions, and text
pages = doc . pages ()
for page in pages :
print ( f"Page { page . index + 1 } : { page . dimensions . width } x { page . dimensions . height } pt" )
print ( f"Strokes: { len ( page . strokes ) } " )
for stroke in page . strokes :
print ( f" Stroke { stroke . uuid } : color { stroke . color_hex } , alpha { stroke . alpha } , points { len ( stroke . points ) } " )
# Text and Sticky Notes content
fragments = doc . text_fragments ()
for frag in fragments :
print ( f"[ { frag . source_path } ] { frag . format } : { frag . text } " )
# Structural page-element summaries for format analysis
for page in pages :
for element in page . elements :
print ( element . kind , element . uuid , element . attachment_uuid , element . related_uuids )
Documentation
Document
Description
cli.md
Full CLI reference
GitHub Wiki
Deep-dive technical documentation (architecture, formats, rendering)
LEGAL-NOTICE.md
Legal, trademark, privacy, and redistribution notice
See also the Contributing Guide .
Legal & Trademark Notice : "Goodnotes" and related names, logos, and marks are the property of Goodnotes Limited. Document Parser for GoodNotes is an independent, community-developed project and is not affiliated with, endorsed by, sponsored by, or officially connected to Goodnotes Limited . For full legal, trademark, privacy, and redistribution details, please read LEGAL-NOTICE.md .
Vibe Coding 免責聲明 ：本專案完全採用 Vibe Coding （AI 輔助快速結對程式設計與探索式開發）進行構建。雖然解析器已通過測試樣本驗證，但程式碼結構與架構選擇體現了 AI 驅動的實驗性疊代風格。請自行評估並謹慎使用！
一套獨立、開源且完整型別化的 Python 工具組，用於檢視與解析使用者提供的 GoodNotes 5 與 GoodNotes 6 .goodnotes 封存檔。它可直接解碼 protobuf wire format 、解析 Apple LZ4 框架串流、解碼 Troy Hanson TPL 記憶體映像、擷取觀察到的 RGBA 筆跡資料，並將文件匯出為 JSON 與 SVG。
本專案與 Goodnotes Limited 沒有任何關聯、背書、贊助或官方合作關係。 發布與使用注意事項請參閱 LEGAL-NOTICE.md 。
原始封存檔
GoodNotes 原版渲染 ( .jpg )
本專案 SVG 匯出 ( .svg )
範例 1：手寫公式與插圖
( ex1.goodnotes )
範例 2：多款筆刷與色彩筆跡
( ex2.goodnotes )
範例 3：多層圖文疊加與中文手寫
( ex3.goodnotes )
開發動機與解析原理
GoodNotes 是一個很棒的筆記軟體，但由於其封閉性，除了匯出成封閉的 .goodnotes 專屬檔案，或是會失去編輯能力的 PDF 之外，幾乎沒有其他保留向量筆跡的選擇，因此我開發了一個解析工具來解析 .goodnotes 檔案。
因為我對這類工程沒有經驗，且解析此 .goodnotes 不是那麼容易（目前沒有看到幾個有成功解析出來的專案），因此我使用 vibe coding (主要是 Gemini 3.1 Pro, Gemini 3.6 Flash 及 Claude Sonnet 5) 來完成這個專案。
壓縮與 Protobuf 盲解： 經過分析，得知 .goodnotes 本質上是一個壓縮檔，而主要的筆跡資訊以頁為單位存在 notes/ 下，各頁面的筆跡資訊存為 Protobuf 序列化檔案。在這個專案中，使用了 Wire Format 的方式來解析，將其構建成抽象語法樹。
Apple LZ4 逆向： 接著，我們發現了部分欄位以 bv41 或 bv4- 開頭，並且以 bv4$ 結尾，代表這是經過 Apple LZ4 壓縮的資料，我們利用維護 64KB 的歷史滑動視窗，靠位元運算處理 LZ4 的 token 來將其成功解壓。
TPL 記憶體映像： 解壓後的明文以 tpl\0 開頭，因此得知他是 Troy Hanson's TPL 格式，並透過格式字串推導出他的筆跡資訊（帶有壓感的離散點）。
向量幾何重建： 接著，透過計算相鄰兩點間的法向量並進行滑動平均平滑化，再結合壓感值向外推移，就能將離散點縫合成封閉的多邊形，最終輸出為帶有自然粗細變化的 SVG 向量筆跡。
目前，專案可以分析 .goodnotes 中的二進制檔案來獲取筆跡、文字等資訊，並輸出成 .svg 或 .pdf 格式。與直接從 GoodNotes app 匯出不同，本專案因為解析了 .goodnotes 格式，因此未來完全可以將解析出的資料，轉換為其他開源格式（如 InkML），讓使用者的心血不再被單一廠商綁架。
Protobuf Wire Decoder ：無損解析有框架與無框架的 protobuf 訊息，並保留未知欄位。
Apple LZ4 與 Troy Hanson TPL Decoder ：解壓縮 bv41 LZ4 串流，並將嵌入的 TPL 格式字串解碼為結構化筆跡點與可變寬度筆跡帶。
筆跡與色彩解析 ：解析 bv4$ 後的 protobuf trailer，以擷取精確 RGBA 顏色與螢光筆透明度。
筆跡與幾何元件支援 ：解析測試樣本中觀察到的單點、直線、曲線、鋼筆/原子筆工具、螢光筆、橡皮擦切口、圖形及移動/複製元素。
頁面與背景解析 ：自動偵測 PDF /MediaBox 尺寸（A4、Letter、橫向與直向）、頁面順序、文字片段與便條紙內容。
CLI 工具 ： gn-inspect 、 gn-dump 、 gn-diff 、 gn-export-json 、 gn-export-svg 、 gn-export-pdf 。
✅ 目前測試可正常輸出
鋼筆、原子筆、畫筆、螢光筆 ：可處理不同顏色及粗細。
uv run pytest
注意： samples/ 刻意被忽略，且不屬於公開原始碼樹。除非您擁有重新發布的權限，否則請勿發布 .goodnotes 檔案或擷取出的資產。
# 檢視封存檔目錄清單與 sha256 校驗碼
gn-inspect sample.goodnotes
# 無損印出任何 protobuf 成員至 JSON
gn-dump sample.goodnotes index.notes.pb
# 比對兩個 .goodnotes 封存檔差異
gn-diff before.goodnotes after.goodnotes
# 匯出整份文件、元資料、頁面、筆跡與原始 wir

[truncated]
