---
source: "https://github.com/convictional/souls-only"
hn_url: "https://news.ycombinator.com/item?id=48445469"
title: "SoulsOnly.tff – A font for humans not AI and keyboard firmware to type in it"
article_title: "GitHub - convictional/souls-only: A font for humans not AI and keyboard firmware to type in it. · GitHub"
author: "billtarbell"
captured_at: "2026-06-08T16:31:18Z"
capture_tool: "hn-digest"
hn_id: 48445469
score: 28
comments: 13
posted_at: "2026-06-08T14:00:22Z"
tags:
  - hacker-news
  - translated
---

# SoulsOnly.tff – A font for humans not AI and keyboard firmware to type in it

- HN: [48445469](https://news.ycombinator.com/item?id=48445469)
- Source: [github.com](https://github.com/convictional/souls-only)
- Score: 28
- Comments: 13
- Posted: 2026-06-08T14:00:22Z

## Translation

タイトル: SoulsOnly.tff – AI ではなく人間が入力するためのフォントとキーボード ファームウェア
記事タイトル: GitHub - convictional/souls-only: AI ではなく人間が入力するためのフォントとキーボード ファームウェア。 · GitHub
説明: AI やキーボード ファームウェアではなく人間が入力するためのフォント。 - 確信犯/魂のみ

記事本文:
GitHub - 確信犯/魂のみ: AI ではなく人間が入力するためのフォントとキーボード ファームウェア。 · GitHub
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
有罪判決を受けた
/
魂のみ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
72 コミット 72 コミット 基本ベース 暗号 暗号デモ デモ dist dist fontbuild fontbuild メディア メディア スクリプト スクリプト テスト テスト ツール ツール .gitignore .gitignore LICENSE ライセンス README.md README.md font-cipher-brief.md font-cipher-brief.md 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルナビゲーション
Souls Only: 人間が読める、AI には非フレンドリーな入力可能な暗号フォント
レンダリングされたグリフが読み取り可能なテキストを綴るフォント。
文字ストリーム (コピーアンドペースト、HTML/PDF 抽出、スクレイパーが認識するもの)
騒音。フォントはデコーダであり、レンダリング レイヤでのみ適用されます。
暗号は通常のキーボードによって駆動されます。通常のキーを入力すると、キーボードが
ノイズ ストリームを発し、このフォントだけがそれをレンダリングして言葉に戻します。
これは技術とステートメントのプロジェクトであり、突破できないセキュリティを主張するものではありません。
font-cipher-brief.md の制限事項を参照してください。
フォントには、通常混同される 2 つのストリームがあります。文字ストリーム (保存された文字ストリーム)
バイト）とグリフ ストリーム（cmap と GSUB の実行後に描画されるもの）。これ
プロジェクトはそれらを切り離します。
すべての印刷可能な文字は 2 つの半分としてエンコードされ、それぞれの半分は次の時点で選択されます。
2 文字の ASCII コード (同音異義語) のプールからランダムに選択されます。それで、あるキャラクターは、
4 つの ASCII 記号として入力され、同じ文字が異なるバイトを生成します。
毎回。
フォントは、各 ASCII コード キャリアを cmap の空白のグリフにマップし、次に GSUB にマップします。
liga ルールでは、各 2 文字のコードが 1 つの不透明な半グリフに折りたたまれます。
2 つの半グリフが実際の文字にタイル表示されます。共有クラスは 1 つを再利用します
標準的な左半分なので、左の画像があいまいです: 小文字のボウル
a c de g o q 、小文字の語幹 m n r u 、および大文字のボウル
O C G Q はそれぞれ 1 つの左半分のグリフを共有します。半グリフ名は不透明です

えー、それで
フォント テーブル ダンプでは、意味のない半分の形状のみが明らかになり、
ハーフから文字へのマッピング。
4 つの ASCII 文字が 1 つのレンダリングされた文字に折りたたまれるため、保存される
バイト数とレンダリングされたグリフ数は意図的に異なります。
フォントに入力されたプレーン文字はデコードされません。ラテン文字のコードポイントは次のとおりです。
意図的に意味のない半グリフの断片にマッピングされているため、通常のテキストを貼り付ける
フォントを適用するとノイズが発生します。読める単語は暗号からのみ生まれます
これは、通常の書体ではなく、フォントがキーであることを強調します。
ビルドされたフォントは dist/ にコミットされます。
dist/SoulsOnly.ttf — 静的フォント。をレンダリングします
読み取り可能なテキストとしての暗号ストリーム。
dist/SoulsOnly.otf — CFF と同じ静的フォント
(PostScript) の概要 (OTF を優先するツール用)。
dist/SoulsOnly-VF.ttf — 可変フォント
(ファミリー「Souls Only VF」) REVL 散乱軸を使用。デフォルトは
散在 : テキストは REVL = 650 でのみ判読可能です (以下の「明らかに」を参照)。
TTF のみ: 散布軸は TrueType バリエーション データ内に存在します。
最も広くサポートされている可変フォント形式。
ダウンロードしてダブルクリックしてインストールします (macOS では Font Book を右クリック →
Windows にインストールする)、または Web 上で @font-face を使用します。フォントだけを覚えておいてください
暗号ストリームをデコードします - Souls Only でレンダリングされる通常のテキストはノイズです
設計上。以下のエンコーダーまたは暗号キーボードを使用してストリームを生成します
ファームウェア。フォントは OFL 1.1 に基づいてライセンスされています (ライセンスを参照)。
Souls Only は、印刷可能な US-QWERTY セット (小文字、大文字、数字) をすべてカバーしています。
0123456789 、および標準の記号。空白は全体的に編集可能なままです
文字: すべての文字は 4 バイトであるため、キーボードで削除および移動します。
4 つ単位 (バックスペースで 4 つ削除、矢印で 4 つ移動)、スペースで 1 つ出力
スペース文字、および Return e

本物の改行と 3 つの目に見えないパッドバイトを含みます
したがって、ストリームは 4 つ整列されたままになります。
Souls Only は、3 つの軸から構築されたカスタム REVL 軸を備えた可変フォントとして出荷されます。
マスター:
REVL = 0 (デフォルトなので安全な状態は判読不能) では、すべてのグリフは次のようになります。
ランダムな不均一変換とポイントごとによって認識から歪められました
ジッター、
REVL = 650 (軸の中央) では、すべての点が補間されて元の位置に戻ります。
真の位置とテキストが組み立てられます。
REVL = 1000 では、グリフは再び別の歪みに散乱します。
コントロールを一番上まで押してもテキストは表示されません。
デコードおよび表示メカニズム全体はフォント ( cmap 、 GSUB 、
ハーフグリフ タイリング、および fvar / gvar );ページは単一の REVL のみを提供します
1 つのコントロールを介して軸の値を変更できます。軸には名前が付けられておらず、判読可能な名前もありません。
したがって、公開値が自動リーダーに無料で渡されることはありません。
正直な制限 (仕様から再説明): REVL 値は 1 つの制限された数値です。
そのため、自動攻撃者は軸の値をスイープし、判読可能なフレームを OCR することができます。これ
レイヤーは、最も移植性が高く、自己完結型のリビールであり、攻撃に対して最も弱いものです。
自動化されたビジョン。これはステートメントデバイスであり、そのようにスコープされています。
既知の制限 (設計による): 共有される左半分は単一の侵害イメージです
クラス全体で再利用されます。ボウルクラスはきれいに共有されます。ステムクラス
m n r u はそうではありません。実際の文字から切り取られたステムは純粋なバーではないため、これらは
文字にはかすかなヘアラインの継ぎ目があります。延期された修正は手描きの合成です
共有ステムグリフ ( fontbuild/fragments.py の TODO を参照)。
ドキュメント/スーパーパワー/設計仕様と実装計画
cipher/charset.py 文字セット + ハーフスロット モデル: 唯一の信頼できる情報源
cipher/keyboard.py ASCII キャリアコード割り当て + オラクルのエンコード/デコード
cipher/qwerty.py US-QWERTY キーコード -> 文字マップ
f

ontbuild/fragments.py ハーフグリフのスライス (skia-pathops)
fontbuild/features.py FEA ファイルからの GSUB liga コンパイル
fontbuild/build_keyboard.py build dist/SoulsOnly.ttf (+ REVL 変数フォント)
fontbuild/reveal.py 3 つのマスターから REVL リビール フォントをビルドします
tools/make_qmk_table.py cipher/keyboard から QMK ファームウェア テーブルを生成します
tools/make_demo_assets.py ブラウザのデモ テーブルを生成し、VF をコピーします
tools/make_keys_preview.py 生成 dist/keys.html (REVL スライダー プレビュー)
デモ/物理暗号キーボードのデモ (QMK キーマップ、ランブック、ページ)
テスト/ pytest スイート (python -m pytest 経由で実行)
Base/Jost-REGULAR.ttf インスタンス化された OFL ベース フォント (グリフ アウトライン)
( dist/ および生成されたデモ アセットはビルド アーティファクトであり、gitigned されます。)
python3 -m venv .venv
./.venv/bin/pip install -rrequirements.txt
bash scripts/fetch_base_font.sh #base/Jost- Regular.ttf が見つからない場合
./.venv/bin/python -m fontbuild.build_keyboard # dist/SoulsOnly.ttf + SoulsOnly-VF.ttf
./.venv/bin/python tools/make_otf.py # dist/SoulsOnly.otf (CFF アウトライン)
./.venv/bin/python -m pytest # スイートを実行します
./.venv/bin/python tools/make_keys_preview.py # dist/keys.html (REVL スライダー)
# 次に: python -m http.server 8753 を実行し、dist/keys.html を開きます。
# 物理キーボードのデモ アセットを再生成します。
./.venv/bin/python tools/make_qmk_table.py #demo/qmk/cipher_table.h
./.venv/bin/python tools/make_demo_assets.py # デモ/cipher_table.js + デモ/SoulsOnly-VF.ttf
# 次に、demo/index.html を開きます (ハードウェア Runbook については、demo/BUILD.md を参照してください)
手動でエンコードおよびデコードする
./.venv/bin/python -c " 暗号インポートキーボードから k; print(k.encode('hello world')) "
./.venv/bin/python -c " 暗号インポートキーボードから k; print(k.decode(k.encode('hello world'))) "
ライセンス
コード (暗号、フォントビルド、ツール、ファームウェアグルー): MIT 。
フォントファイル

: グリフのアウトラインの由来
ジョスト 、著作権 2020 ジョスト
プロジェクト作成者、以下に基づいてライセンスを取得
SIL オープン フォント ライセンス 1.1 。献身的な
base/Jost-Regular.ttf (インスタンス化) およびビルドされた SoulsOnly*.ttf は次のとおりです。
派生フォント ソフトウェアであり、同じ OFL 1.1 の下で配布されます。
MITではありません。
AIやキーボードファームウェアではなく人間が入力するためのフォント。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A font for humans not AI and keyboard firmware to type in it. - convictional/souls-only

GitHub - convictional/souls-only: A font for humans not AI and keyboard firmware to type in it. · GitHub
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
convictional
/
souls-only
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
72 Commits 72 Commits base base cipher cipher demo demo dist dist fontbuild fontbuild media media scripts scripts tests tests tools tools .gitignore .gitignore LICENSE LICENSE README.md README.md font-cipher-brief.md font-cipher-brief.md requirements.txt requirements.txt View all files Repository files navigation
Souls Only: a human-readable, AI-unfriendly cipher font you can type
A font whose rendered glyphs spell readable text while the stored
character stream (what copy-paste, HTML/PDF extraction, and scrapers see) is
noise. The font is the decoder, applied only at the rendering layer, and the
cipher is driven by an ordinary keyboard: you type normal keys, the keyboard
emits the noise stream, and only this font renders it back into words.
This is a craft and statement project, not a claim of unbreakable security.
See Limitations in font-cipher-brief.md .
A font has two streams people usually conflate: the character stream (stored
bytes) and the glyph stream (what is drawn after cmap and GSUB run). This
project decouples them:
Every printable character is encoded as two halves, and each half is chosen at
random from a pool of 2-character ASCII codes (homophones). So one character is
typed as four ASCII symbols, and the same character produces different bytes
each time.
The font maps each ASCII code carrier to a blank glyph in cmap , then a GSUB
liga rule collapses each 2-character code into one opaque half-glyph.
The two half-glyphs tile into the real character. Shared classes reuse one
canonical left half so the left image is ambiguous: the lowercase bowl
a c d e g o q , the lowercase stem m n r u , and the uppercase bowl
O C G Q each share a single left half-glyph. Half-glyph names are opaque, so
a font-table dump reveals only meaningless half-shapes, never a
half-to-character mapping.
Because four ASCII characters collapse into one rendered character, the stored
byte count and the rendered glyph count deliberately diverge.
Plain letters typed in the font do NOT decode: the Latin letter codepoints are
deliberately mapped to meaningless half-glyph fragments, so pasting ordinary text
and applying the font yields noise. Readable words only ever come from the cipher
stream, which reinforces that the font is the key, not a normal typeface.
The built fonts are committed in dist/ :
dist/SoulsOnly.ttf — the static font. Renders a
cipher stream as readable text.
dist/SoulsOnly.otf — the same static font with CFF
(PostScript) outlines, for tools that prefer OTF.
dist/SoulsOnly-VF.ttf — the variable font
(family "Souls Only VF") with the REVL scatter axis. Defaults to
scattered : text is legible only at REVL = 650 (see "The reveal" below).
TTF only: the scatter axis lives in TrueType variation data, which is also
the most widely supported variable-font format.
Download and double-click to install (Font Book on macOS, right-click →
Install on Windows), or use @font-face on the web. Remember the font only
decodes the cipher stream — ordinary text rendered in Souls Only is noise
by design. Generate a stream with the encoder below or the cipher keyboard
firmware. The fonts are licensed under the OFL 1.1 (see Licensing).
Souls Only covers the full US-QWERTY printable set: lowercase, uppercase, digits
0123456789 , and the standard symbols. Whitespace stays editable in whole
characters: every character is four bytes, so the keyboard deletes and navigates
in units of four (Backspace removes four, the arrows move four), Space emits one
space character, and Return emits a real newline plus three invisible pad bytes
so the stream stays four-aligned.
Souls Only ships as a variable font with a custom REVL axis built from three
masters:
at REVL = 0 (the default, so the safe state is illegible) every glyph is
warped out of recognition by a random non-uniform transform plus per-point
jitter,
at REVL = 650 (the middle of the axis) every point interpolates back to its
true position and the text assembles,
at REVL = 1000 the glyphs scatter again into a different distortion, so
pushing the control all the way up does not reveal the text either.
The entire decode and reveal mechanism lives in the font ( cmap , GSUB ,
half-glyph tiling, and fvar / gvar ); a page contributes only the single REVL
axis value via one control. The axis is unnamed and there is no legible named
instance, so the reveal value is not handed to an automated reader for free.
Honest limit (restated from the spec): the REVL value is one bounded number,
so an automated attacker can sweep axis values and OCR the legible frame. This
layer is the most portable and self-contained reveal, and the weakest against
automated vision. It is a statement device, scoped as such.
Known limit (by design): the shared left half is a single compromise image
reused across a class. The bowl classes share cleanly. The stem class
m n r u does not: a stem clipped from a real letter is not a pure bar, so those
letters carry a faint hairline seam. The deferred fix is a hand-drawn synthetic
shared-stem glyph (see the TODO in fontbuild/fragments.py ).
docs/superpowers/ design specs and implementation plans
cipher/charset.py the charset + half-slot model: single source of truth
cipher/keyboard.py ASCII carrier-code allocation + encode/decode oracle
cipher/qwerty.py US-QWERTY keycode -> character map
fontbuild/fragments.py half-glyph slicing (skia-pathops)
fontbuild/features.py GSUB liga compilation from a FEA file
fontbuild/build_keyboard.py build dist/SoulsOnly.ttf (+ the REVL variable font)
fontbuild/reveal.py build the REVL reveal font from three masters
tools/make_qmk_table.py generate the QMK firmware table from cipher/keyboard
tools/make_demo_assets.py generate the browser demo table + copy the VF
tools/make_keys_preview.py generate dist/keys.html (the REVL slider preview)
demo/ the physical cipher keyboard demo (QMK keymap, runbook, page)
tests/ pytest suite (run via python -m pytest)
base/Jost-Regular.ttf instanced OFL base font (glyph outlines)
( dist/ and the generated demo assets are build artifacts and are gitignored.)
python3 -m venv .venv
./.venv/bin/pip install -r requirements.txt
bash scripts/fetch_base_font.sh # if base/Jost-Regular.ttf is missing
./.venv/bin/python -m fontbuild.build_keyboard # dist/SoulsOnly.ttf + SoulsOnly-VF.ttf
./.venv/bin/python tools/make_otf.py # dist/SoulsOnly.otf (CFF outlines)
./.venv/bin/python -m pytest # run the suite
./.venv/bin/python tools/make_keys_preview.py # dist/keys.html (the REVL slider)
# then: python -m http.server 8753 and open dist/keys.html
# regenerate the physical keyboard demo assets:
./.venv/bin/python tools/make_qmk_table.py # demo/qmk/cipher_table.h
./.venv/bin/python tools/make_demo_assets.py # demo/cipher_table.js + demo/SoulsOnly-VF.ttf
# then open demo/index.html (see demo/BUILD.md for the hardware runbook)
Encode and decode by hand
./.venv/bin/python -c " from cipher import keyboard as k; print(k.encode('hello world')) "
./.venv/bin/python -c " from cipher import keyboard as k; print(k.decode(k.encode('hello world'))) "
Licensing
Code (cipher, fontbuild, tools, firmware glue): MIT .
Font files : glyph outlines come from
Jost , Copyright 2020 The Jost
Project Authors, licensed under the
SIL Open Font License 1.1 . The committed
base/Jost-Regular.ttf (instanced) and any built SoulsOnly*.ttf are
derivative Font Software and are distributed under the same OFL 1.1 — they
are not MIT.
A font for humans not AI and keyboard firmware to type in it.
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
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
