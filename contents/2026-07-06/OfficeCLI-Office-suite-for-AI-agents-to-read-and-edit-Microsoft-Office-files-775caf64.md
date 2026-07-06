---
source: "https://github.com/iOfficeAI/OfficeCLI"
hn_url: "https://news.ycombinator.com/item?id=48807225"
title: "OfficeCLI: Office suite for AI agents to read and edit Microsoft Office files"
article_title: "GitHub - iOfficeAI/OfficeCLI: OfficeCLI is the first and best Office suite purpose-built for AI agents to read, edit, and automate Word, Excel, and PowerPoint files. Free, open-source, single binary, no Office installation required. · GitHub"
author: "maxloh"
captured_at: "2026-07-06T17:48:27Z"
capture_tool: "hn-digest"
hn_id: 48807225
score: 12
comments: 1
posted_at: "2026-07-06T16:47:44Z"
tags:
  - hacker-news
  - translated
---

# OfficeCLI: Office suite for AI agents to read and edit Microsoft Office files

- HN: [48807225](https://news.ycombinator.com/item?id=48807225)
- Source: [github.com](https://github.com/iOfficeAI/OfficeCLI)
- Score: 12
- Comments: 1
- Posted: 2026-07-06T16:47:44Z

## Translation

タイトル: OfficeCLI: AI エージェントが Microsoft Office ファイルを読み取り、編集するための Office スイート
記事のタイトル: GitHub - iOfficeAI/OfficeCLI: OfficeCLI は、AI エージェントが Word、Excel、PowerPoint のファイルを読み取り、編集、自動化することを目的として構築された最初で最高の Office スイートです。無料、オープンソース、単一バイナリで、Office のインストールは不要です。 · GitHub
説明: OfficeCLI は、AI エージェントが Word、Excel、PowerPoint のファイルを読み取り、編集、自動化することを目的として構築された最初で最高の Office スイートです。無料、オープンソース、単一バイナリで、Office のインストールは不要です。 - iOfficeAI/OfficeCLI

記事本文:
GitHub - iOfficeAI/OfficeCLI: OfficeCLI は、AI エージェントが Word、Excel、PowerPoint のファイルを読み取り、編集、自動化することを目的として構築された最初で最高の Office スイートです。無料、オープンソース、単一バイナリで、Office のインストールは不要です。 · GitHub
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
{{ メッセージ }

}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
iOfficeAI
/
OfficeCLI
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5,527 コミット 5,527 コミット .github/ workflows .github/ workflows アセット アセット ビルド サンプル サンプル npm npm プラグイン プラグイン スキーマ スキーマ sdk sdk スキル スキル src/ officecli src/ officecli COTRIBUTING.md COTRIBUTING.md COTRIBUTING.zh.md COTRIBUTING.zh.md ライセンス ライセンス通知 通知 README.md README.md README_ja.md README_ja.md README_ko.md README_ko.md README_zh.md README_zh.md SECURITY.md SECURITY.md SKILL.md SKILL.md THIRD-PARTY-NOTICES.txt THIRD-PARTY-NOTICES.txt build.sh build.sh dev-install.sh dev-install.sh install.ps1 install.ps1 install.sh install.sh officecli.slnx officecli.slnx すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OfficeCLI は、AI エージェント向けに設計された世界初かつ最高の Office スイートです。
1 行のコードで、AI エージェントに Word、Excel、PowerPoint を完全に制御させることができます。
オープンソース。単一のバイナリ。 Officeのインストールはありません。依存関係はありません。どこでも使えます。
OfficeCLI に組み込まれた HTML レンダリング エンジンは、ドキュメントを高い忠実度で再現します。これが AI の目を与えます。 .docx / .xlsx / .pptx を HTML または PNG にレンダリングし、レンダリング → 表示 → 修正のループを閉じます。
🌐 ウェブサイト: officecli.ai | 💬 コミュニティ: Discord
AionUi で OfficeCLI を使用した PPT 作成プロセス
上記のドキュメントはすべて、OfficeCLI を使用して AI エージェントによって完全に作成されました。テンプレートや手動編集はありません。
AI エージェント向け — 一行で始めましょう
これを AI エージェントのチャットに貼り付けます。スキル ファイルが読み込まれ、すべてが自動的にインストールされます。
カール -fsSL https://officecli.ai/SKILL.md
Th

それはそれです。スキル ファイルは、バイナリのインストール方法とすべてのコマンドの使用方法をエージェントに教えます。
オプション A — GUI: AionUi をインストールします。これは、内部で OfficeCLI を利用して、自然言語を通じて Office ドキュメントを作成および編集できるデスクトップ アプリです。必要なものを説明するだけで、残りは AionUi が処理します。
オプション B — CLI: GitHub Releases からプラットフォームのバイナリをダウンロードし、次を実行します。
officecli インストール
これにより、バイナリが PATH にコピーされ、検出されたすべての AI コーディング エージェント (Claude Code、Cursor、Windsurf、GitHub Copilot など) に officecli スキルがインストールされます。エージェントは、追加の構成を必要とせずに、ユーザーに代わって Office ドキュメントをすぐに作成、読み取り、編集できます。
開発者向け — 30 秒でライブでご覧いただけます
# 1. インストール (macOS / Linux)
カール -fsSL https://raw.githubusercontent.com/iOfficeAI/OfficeCLI/main/install.sh |バッシュ
# Windows (PowerShell): irm https://raw.githubusercontent.com/iOfficeAI/OfficeCLI/main/install.ps1 |アイエックス
# 2. 空の PowerPoint を作成する
officecli 作成デッキ.pptx
# 3. ライブ プレビューを開始します — ブラウザーで http://localhost:26315 を開きます
officecli ウォッチデッキ.pptx
# 4. 別のターミナルを開き、スライドを追加します - ブラウザーの更新を即座に確認します
officecli add Deck.pptx / --type slide --prop title= " Hello, World! "
それだけです。 add 、 set 、または delete コマンドを実行するたびに、プレビューがリアルタイムで更新されます。実験を続けてください。ブラウザはライブのフィードバック ループです。
# プレゼンテーションを作成してコンテンツを追加する
officecli 作成デッキ.pptx
officecli add Deck.pptx / --type slide --prop title= " Q4 レポート " --prop background=1A1A2E
officecli add Deck.pptx ' /slide[1] ' --type Shape \
--prop text= " 収益が 25% 増加 " --prop x=2cm --prop y=5cm \
--prop font=Arial --prop size=24 --prop color=FFFFFF
# アウトラインとして表示
officecli ビューデッキ.pptx o

概要
# → スライド 1: 第 4 四半期レポート
# → Shape 1 [TextBox]: 収益が 25% 増加
# HTML として表示 — ブラウザーでレンダリングされたプレビューが開きます。サーバーは必要ありません
officecli ビューデッキ.pptx html
# 任意の要素の構造化された JSON を取得する
officecli 取得デッキ.pptx ' /slide[1]/shape[1] ' --json
# 保存して閉じる — 常駐セッションをディスクにフラッシュします
officecli デッキを閉じる.pptx
{
"タグ" : "形状" ,
"パス" : " /slide[1]/shape[1] " ,
「属性」: {
"名前" : " テキストボックス 1 " ,
"text" : " 収益が 25% 増加しました " ,
"x" : " 720000 " 、
"y" : " 1800000 "
}
}
なぜ OfficeCLI なのか?
以前は 50 行の Python と 3 つの個別のライブラリが必要でした。
pptxインポートからプレゼンテーション
pptx から。 util import インチ、Pt
prs = プレゼンテーション ()
スライド = PRS 。スライド。 add_slide (prs . slide_layouts [ 0 ])
タイトル = スライド。形状。タイトル
タイトル。 text = "第 4 四半期レポート"
# ... あと 45 行 ...
広報。保存 ( 'deck.pptx' )
次のコマンドを 1 つ実行します。
officecli add Deck.pptx / --type slide --prop title= " 第 4 四半期レポート "
OfficeCLI でできること:
ドキュメントを最初から作成します -- 空白またはコンテンツ付き
テキスト、構造、スタイル、数式をプレーンテキストまたは構造化された JSON で読み取ります。
書式設定の問題、スタイルの不一致、構造上の問題を分析する
テキスト、フォント、色、レイアウト、数式、グラフ、画像などの要素を変更します
コンテンツを再編成 -- ドキュメント間で要素を追加、削除、移動、コピーする
Word — i18n および RTL の完全なサポート (スクリプトごとのフォント スロット、スクリプトごとの BCP-47 言語タグ lang.latin/ea/cs 、複合スクリプトの太字/斜体/サイズ、段落/実行/セクション/テーブル/スタイル/ヘッダー/フッター/docDefaults を介した方向 = rtl カスケード、rtlGutter + pgBorders 短縮表現、ロケール対応のページ番号付け)ヒンディー語/アラビア語/タイ語/CJK; create --locale ar-SA は RTL を自動有効にします)、段落 (framePr、タブの短縮表現、文字ベースのインデント)、ラン (underline.color、位置のハーフポイント)、テーブル (v)

仮想列操作の追加/削除/移動/コピーフロム、hMerge)、スタイル、テキストボックス/形状 (回転、verticalText eaVert / vert270 、グラデーション、シャドウ、不透明度)、ヘッダー/フッター、画像 (PNG/JPG/GIF/SVG)、方程式、コメント、脚注、透かし、ブックマーク、目次、チャート、ハイパーリンク、セクション、フォーム フィールド、コンテンツ コントロール (SDT) 、フィールド (22 のゼロパラメータ タイプ + MERGEFIELD / REF / PAGEREF / SEQ / STYLEREF / DOCPROPERTY / IF)、OLE オブジェクト 、リビジョン / 追跡された変更 (reviation.type=ins\|del\|format\|moveFrom\|moveTo +reviation.action=accept\|reject 、ターゲットごとの /revision[@author=Alice] セレクター、追跡された検索と置換)、ページの背景色、ドキュメントのプロパティ
Excel — セル（追加時にふりがな/ふりがな、Excel-UI --shift left\|up時に削除/shift=right\|down時に追加）、数式（自動評価機能付きの350以上の組み込み関数、_xlfnによる動的配列のスピル、自動プレフィックス、金融/債券および統計ファミリ、OFFSET/INDIRECT、解析時にインライン化される定義名式本体、行/列挿入時の数式参照書き換え）、シート(visible/hidden/veryHidden、印刷マージン、printTitleRows/Cols、RTL SheetView 、カスケード対応シート名変更、オープン時の空セル膨張フィルター)、ブール値および/またはセレクター ( row[Salary>5000 and Regional=EMEA] )、テーブル、並べ替え (シート/範囲、マルチキー、サイドカー対応)、条件付き書式設定、チャート (含む)ボックスひげ、自動並べ替え + 累積% 付きのパレート、対数軸)、ピボット テーブル (複数フィールド、日付グループ化、showDataAs、並べ替え、grandTotals、小計、コンパクト/アウトライン/表形式レイアウト、繰り返し項目ラベル、空白行、計算フィールド、永続 labelFilter / topN / fillDownLabels 、キャッシュ CoW + クロスピボット共有)、スライサー、名前付き範囲、データ検証、イメージ(二重表現フォールバックを備えた PNG/JPG/GIF/SVG)、スパークライン、コメント (RTL)、オートフィルター

r 、シェイプ 、OLE オブジェクト 、CSV/TSV インポート、$Sheet:A1 セル アドレス指定
PowerPoint - スライド (ヘッダー/フッター/日付/スライド番号の切り替え、非表示)、図形 (パターン塗りつぶし、ぼかし効果、ハイパーリンク ツールヒント + スライド ジャンプ リンク、実行時のハイライト カラー、slideMaster / slideLayout 型の追加/設定/削除、矢印エイリアス、effect.X + active.X.src)、画像 (PNG/JPG/GIF/SVG、塗りつぶしモード: ストレッチ/含む/カバー/タイル、明るさ/コントラスト/グロー/シャドウ、回転、リンク + ツールチップ)、テーブル (組み込みの PowerPoint スタイル カタログ、仮想 /col[C] get + swap/copyFrom、行/列 Move/CopyFrom、塗りつぶし/背景エイリアス)、チャート (pieOfPie、barOfPie、属性ごとの axisLine/グリッドライン セッター、テーマ パレットによる系列の追加/削除、アンカー=x、y、w、h の短縮形)、アニメーション (15)強調 + 16 の終了テンプレートベースのプリセット、マルチエフェクト チェーン、モーション パス プリセット、リピート/リスタート/オートリバース、チャート アニメーション + chartBuild )、トランジション (モーフ + p14 + 12 p15 PowerPoint 2013+ プリセット)、3D モデル (.glb) (結合された回転=ax,ay,az )、スライド ズーム 、方程式 、テーマ 、コネクタ ( @name= セレクターfrom/to)、ビデオ/オーディオ (ループ、autoStart)、グループ (リンク + ツールヒント、取得/クエリ/追加/削除はすべてグループに分類されます)、メモ (RTL、lang)、コメント (RTL、レガシー + 最新の p188 スレッドの往復)、SmartArt (追加部分 + raw セットによる往復)、OLE オブジェクト、プレースホルダー (phType による追加/設定)
データベースまたは API からのレポート生成を自動化する
ドキュメントのバッチ処理 (一括検索/置換、スタイル更新)
CI/CD 環境でドキュメント パイプラインを構築する (テスト結果からドキュメントを生成)
Docker/コンテナ化環境でのヘッドレス オフィス オートメーション
ユーザー プロンプトからプレゼンテーションを生成します (上記の例を参照)
ドキュメントから構造化データを抽出して JSON に変換する
納品前に文書の品質を検証およびチェックする
ドキュメントテンプレートを複製してデータを入力する
オートム

CI/CD パイプラインでの文書検証
単一の自己完結型バイナリとして出荷されます。 .NET ランタイムが組み込まれているため、インストールする必要も、ランタイムを管理する必要もありません。
# macOS / Linux
カール -fsSL https://raw.githubusercontent.com/iOfficeAI/OfficeCLI/main/install.sh |バッシュ
# Windows (PowerShell)
IRM https://raw.githubusercontent.com/iOfficeAI/OfficeCLI/main/install.ps1 |アイエックス
または、GitHub リリースから手動でダウンロードします。
インストールを確認します: officecli --version
または、ダウンロードしたバイナリから自己インストールします (または、ベア officecli を実行して自動インストールします)。
officecli インストール # 明示的
officecli # ベア呼び出しでもインストールがトリガーされます
更新はバックグラウンドで自動的にチェックされます。 officecli config autoUpdate false で無効にするか、 OFFICECLI_SKIP_UPDATE=1 で呼び出しごとをスキップします。構成は ~/.officecli/config.json の下にあります。
内蔵エンジンと生成プリミティブ
OfficeCLI は自己完結型です。以下の機能はバイナリ内に同梱されており、Office は必要ありません。
OfficeCLI の要石: AI エージェントが DOM から推測するのではなく、レンダリングされたドキュメントを確認できるようにする、最初から作成した忠実度の高い HTML レンダリング エンジン。形状、チャート (トレンドライン、エラーバー、ウォーターフォール、ローソク足、スパークライン)、方程式 (OMML → MathJax 互換)、3 つの 3D .glb モデルをカバーします。

[切り捨てられた]

## Original Extract

OfficeCLI is the first and best Office suite purpose-built for AI agents to read, edit, and automate Word, Excel, and PowerPoint files. Free, open-source, single binary, no Office installation required. - iOfficeAI/OfficeCLI

GitHub - iOfficeAI/OfficeCLI: OfficeCLI is the first and best Office suite purpose-built for AI agents to read, edit, and automate Word, Excel, and PowerPoint files. Free, open-source, single binary, no Office installation required. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
iOfficeAI
/
OfficeCLI
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5,527 Commits 5,527 Commits .github/ workflows .github/ workflows assets assets build build examples examples npm npm plugins plugins schemas schemas sdk sdk skills skills src/ officecli src/ officecli CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTING.zh.md CONTRIBUTING.zh.md LICENSE LICENSE NOTICE NOTICE README.md README.md README_ja.md README_ja.md README_ko.md README_ko.md README_zh.md README_zh.md SECURITY.md SECURITY.md SKILL.md SKILL.md THIRD-PARTY-NOTICES.txt THIRD-PARTY-NOTICES.txt build.sh build.sh dev-install.sh dev-install.sh install.ps1 install.ps1 install.sh install.sh officecli.slnx officecli.slnx View all files Repository files navigation
OfficeCLI is the world's first and the best Office suite designed for AI agents.
Give any AI agent full control over Word, Excel, and PowerPoint — in one line of code.
Open-source. Single binary. No Office installation. No dependencies. Works everywhere.
OfficeCLI's built-in HTML rendering engine reproduces documents with high fidelity — and that's what gives AI eyes. It renders .docx / .xlsx / .pptx to HTML or PNG, closing the render → look → fix loop.
🌐 Website: officecli.ai | 💬 Community: Discord
PPT creation process using OfficeCLI on AionUi
All documents above were created entirely by AI agents using OfficeCLI — no templates, no manual editing.
For AI Agents — Get Started in One Line
Paste this into your AI agent's chat — it will read the skill file and install everything automatically:
curl -fsSL https://officecli.ai/SKILL.md
That's it. The skill file teaches the agent how to install the binary and use all commands.
Option A — GUI: Install AionUi — a desktop app that lets you create and edit Office documents through natural language, powered by OfficeCLI under the hood. Just describe what you want, and AionUi handles the rest.
Option B — CLI: Download the binary for your platform from GitHub Releases , then run:
officecli install
This copies the binary to your PATH and installs the officecli skill into every AI coding agent it detects — Claude Code, Cursor, Windsurf, GitHub Copilot, and more. Your agent can immediately create, read, and edit Office documents on your behalf, no extra configuration needed.
For Developers — See It Live in 30 Seconds
# 1. Install (macOS / Linux)
curl -fsSL https://raw.githubusercontent.com/iOfficeAI/OfficeCLI/main/install.sh | bash
# Windows (PowerShell): irm https://raw.githubusercontent.com/iOfficeAI/OfficeCLI/main/install.ps1 | iex
# 2. Create a blank PowerPoint
officecli create deck.pptx
# 3. Start live preview — opens http://localhost:26315 in your browser
officecli watch deck.pptx
# 4. Open another terminal, add a slide — watch the browser update instantly
officecli add deck.pptx / --type slide --prop title= " Hello, World! "
That's it. Every add , set , or remove command you run will refresh the preview in real time. Keep experimenting — the browser is your live feedback loop.
# Create a presentation and add content
officecli create deck.pptx
officecli add deck.pptx / --type slide --prop title= " Q4 Report " --prop background=1A1A2E
officecli add deck.pptx ' /slide[1] ' --type shape \
--prop text= " Revenue grew 25% " --prop x=2cm --prop y=5cm \
--prop font=Arial --prop size=24 --prop color=FFFFFF
# View as outline
officecli view deck.pptx outline
# → Slide 1: Q4 Report
# → Shape 1 [TextBox]: Revenue grew 25%
# View as HTML — opens a rendered preview in your browser, no server needed
officecli view deck.pptx html
# Get structured JSON for any element
officecli get deck.pptx ' /slide[1]/shape[1] ' --json
# Save and close — flushes the resident session to disk
officecli close deck.pptx
{
"tag" : " shape " ,
"path" : " /slide[1]/shape[1] " ,
"attributes" : {
"name" : " TextBox 1 " ,
"text" : " Revenue grew 25% " ,
"x" : " 720000 " ,
"y" : " 1800000 "
}
}
Why OfficeCLI?
What used to take 50 lines of Python and 3 separate libraries:
from pptx import Presentation
from pptx . util import Inches , Pt
prs = Presentation ()
slide = prs . slides . add_slide ( prs . slide_layouts [ 0 ])
title = slide . shapes . title
title . text = "Q4 Report"
# ... 45 more lines ...
prs . save ( 'deck.pptx' )
Now takes one command:
officecli add deck.pptx / --type slide --prop title= " Q4 Report "
What OfficeCLI can do:
Create documents from scratch -- blank or with content
Read text, structure, styles, formulas -- in plain text or structured JSON
Analyze formatting issues, style inconsistencies, and structural problems
Modify any element -- text, fonts, colors, layout, formulas, charts, images
Reorganize content -- add, remove, move, copy elements across documents
Word — full i18n & RTL support (per-script font slots, per-script BCP-47 lang tags lang.latin/ea/cs , complex-script bold/italic/size, direction=rtl cascading through paragraph/run/section/table/style/header/footer/docDefaults, rtlGutter + pgBorders shorthand, locale-aware page numbering for Hindi/Arabic/Thai/CJK; create --locale ar-SA auto-enables RTL), paragraphs (framePr, tabs shorthand, char-based indents), runs (underline.color, position half-pts), tables (virtual column ops add/remove/move/copyfrom, hMerge), styles , textbox / shape (rotation, verticalText eaVert / vert270 , gradient, shadow, opacity), headers/footers , images (PNG/JPG/GIF/SVG), equations , comments , footnotes , watermarks , bookmarks , TOC , charts , hyperlinks , sections , form fields , content controls (SDT) , fields (22 zero-param types + MERGEFIELD / REF / PAGEREF / SEQ / STYLEREF / DOCPROPERTY / IF), OLE objects , revisions / tracked changes ( revision.type=ins\|del\|format\|moveFrom\|moveTo + revision.action=accept\|reject , per-target /revision[@author=Alice] selector, tracked Find&Replace), page background color, document properties
Excel — cells (phonetic guide / furigana on add, Excel-UI --shift left\|up on remove / shift=right\|down on add), formulas (350+ built-in functions with auto-evaluation, spilling dynamic arrays with _xlfn. auto-prefix, financial / bond and statistical families, OFFSET/INDIRECT, defined-name formula bodies inlined at parse, formula-ref rewrite on row/col insert), sheets (visible/hidden/veryHidden, print margins, printTitleRows/Cols, RTL sheetView , cascade-aware sheet rename, empty-cell bloat filter on open), boolean and / or selectors ( row[Salary>5000 and Region=EMEA] ), tables , sort (sheet / range, multi-key, sidecar-aware), conditional formatting , charts (including box-whisker, pareto with auto-sort + cumulative-%, log axis), pivot tables (multi-field, date grouping, showDataAs, sort, grandTotals, subtotals, compact/outline/tabular layout, repeat item labels, blank rows, calculated fields, persistent labelFilter / topN / fillDownLabels , cache CoW + cross-pivot sharing), slicers , named ranges , data validation , images (PNG/JPG/GIF/SVG with dual-representation fallback), sparklines , comments (RTL), autofilter , shapes , OLE objects , CSV/TSV import, $Sheet:A1 cell addressing
PowerPoint — slides (header/footer/date/slidenum toggles, hidden), shapes (pattern fill, blur effect, hyperlink tooltip + slide-jump links, highlight color on runs, slideMaster / slideLayout typed add/set/remove, arrow alias, effective.X + effective.X.src), images (PNG/JPG/GIF/SVG, fill modes: stretch/contain/cover/tile, brightness/contrast/glow/shadow, rotation, link + tooltip), tables (built-in PowerPoint style catalogue, virtual /col[C] get + swap/copyFrom, row/col Move/CopyFrom, fill/background alias), charts (pieOfPie, barOfPie, per-attr axisLine/gridline setters, series add/remove with theme palette, anchor=x,y,w,h shorthand), animations (15 emphasis + 16 exit template-backed presets, multi-effect chains, motion-path presets, repeat/restart/autoReverse, chart animations + chartBuild ), transitions (morph + p14 + 12 p15 PowerPoint 2013+ presets), 3D models (.glb) (combined rotation=ax,ay,az ), slide zoom , equations , themes , connectors ( @name= selector for from/to), video/audio (loop, autoStart), groups (link + tooltip; Get/Query/Add/Remove all descend into groups), notes (RTL, lang), comments (RTL, legacy + modern p188 threaded round-trip), SmartArt (round-trip via add-part + raw-set), OLE objects , placeholders (add/set by phType)
Automate report generation from databases or APIs
Batch-process documents (bulk find/replace, style updates)
Build document pipelines in CI/CD environments (generate docs from test results)
Headless Office automation in Docker/containerized environments
Generate presentations from user prompts (see examples above)
Extract structured data from documents to JSON
Validate and check document quality before delivery
Clone document templates and populate with data
Automated document validation in CI/CD pipelines
Ships as a single self-contained binary. The .NET runtime is embedded -- nothing to install, no runtime to manage.
# macOS / Linux
curl -fsSL https://raw.githubusercontent.com/iOfficeAI/OfficeCLI/main/install.sh | bash
# Windows (PowerShell)
irm https://raw.githubusercontent.com/iOfficeAI/OfficeCLI/main/install.ps1 | iex
Or download manually from GitHub Releases :
Verify installation: officecli --version
Or self-install from a downloaded binary (or run bare officecli to auto-install):
officecli install # explicit
officecli # bare invocation also triggers install
Updates are checked automatically in the background. Disable with officecli config autoUpdate false or skip per-invocation with OFFICECLI_SKIP_UPDATE=1 . Configuration lives under ~/.officecli/config.json .
Built-in Engines & Generation Primitives
OfficeCLI is self-contained. The capabilities below ship inside the binary — no Office required .
OfficeCLI's keystone: a from-scratch, high-fidelity HTML rendering engine that lets an AI agent see the rendered document instead of guessing from the DOM. It covers shapes, charts (trendlines, error bars, waterfall, candlestick, sparklines), equations (OMML → MathJax-compatible), 3D .glb models via Three

[truncated]
