---
source: "https://github.com/ZenKhalil/session-porter"
hn_url: "https://news.ycombinator.com/item?id=49247202"
title: "Move your Claude Code sessions between computers (Mac/Windows, any direction)"
article_title: "GitHub - ZenKhalil/session-porter: Move your Claude Code conversations between computers. Mac and Windows, any direction. One file, no dependencies. · GitHub"
author: "feattzen"
captured_at: "2026-08-10T18:43:00Z"
capture_tool: "hn-digest"
hn_id: 49247202
score: 1
comments: 0
posted_at: "2026-08-10T17:50:33Z"
tags:
  - hacker-news
  - translated
---

# Move your Claude Code sessions between computers (Mac/Windows, any direction)

- HN: [49247202](https://news.ycombinator.com/item?id=49247202)
- Source: [github.com](https://github.com/ZenKhalil/session-porter)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T17:50:33Z

## Translation

タイトル: クロード コード セッションをコンピュータ間で移動 (Mac/Windows、任意の方向)
記事のタイトル: GitHub - ZenKhalil/session-porter: クロード コードの会話をコンピューター間で移動します。 Mac と Windows、どの方向でも。ファイルは 1 つで、依存関係はありません。 · GitHub
説明: クロード コードの会話をコンピューター間で移動します。 Mac と Windows、どの方向でも。ファイルは 1 つで、依存関係はありません。 - ZenKhalil/セッションポーター

記事本文:
GitHub - ZenKhalil/session-porter: クロード コードの会話をコンピューター間で移動します。 Mac と Windows、どの方向でも。ファイルは 1 つで、依存関係はありません。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ゼンカリル
/
セッションポーター
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .gitignore .gitignore ライセンス ライセンス README.md README.md セッション Porter.bat セッション Porter.bat セッション Porter.command S

ession Porter.command session_porter.py session_porter.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コードの会話をコンピュータ間で移動します: Mac から Mac、Mac から Windows、
Windows から Mac、または Windows から Windows。ファイル 1 つ、インストール不要、何も必要ありません
マシンを残します。
Mac: Session Porter.command をダブルクリックします (または python3 session_porter.py を実行します)。
Windows: Session Porter.bat をダブルクリックします (または py session_porter.py を実行します)。
[このコンピュータを梱包する] をクリックします。デスクトップにバンドル zip が 1 つ作成され、次のコピーが含まれます。
その隣に保存されたアプリ。
USB スティック、AirDrop、ネットワーク ドライブ、
または他の何か。
新しいコンピュータで同じ方法でアプリを実行し、[バンドルの解凍] をクリックします。
Porter は、デスクトップ上、ダウンロード内、または接続されたドライブ上でバンドルを見つけます。
それ自体、またはパスを貼り付けることもできます。概要を確認し、「ここで解凍」をクリックしてから、
Claude アプリを完全に終了し、再度開きます。あなたのチャットがそこにあります。
Windows には、python.org の Python 3.9 以降が必要です (「PATH に追加」にチェックを入れます)
インストール中)。 Macにはすでにそれがあります。
.claude フォルダーのプレーン zip だけではセッションを移動するのに十分ではありません。ポーターカバー
ギャップ:
アプリのチャット リストは .claude の外にあります (Mac:
~/ライブラリ/Application Support/Claude/claude-code-sessions 、Windows:
%APPDATA%\Claude\claude-code-sessions )。これがなければ、インポートされたセッションは次の場所に存在します。
ディスクにはありますが、アプリのどこにも表示されません。ポーターはそれを梱包し、元の状態から再構築します。
トランスクリプトが見つからない場合は、作成されたプレーンな .claude zip を取り込むこともできます。
ポーターなしで。
すべてのパスは、4 つの JSON エスケープすべてで新しいホーム フォルダーに書き換えられます。
深さは、両側に境界がある古いホームプレフィックスに固定されています。ユーザー名
これは通常の単語でもあり、テキストが破損することはありません。また、長いユーザー名でも構いません
同じプレフィックスを共有しても、自宅と間違われることはありません。
広報

オブジェクト フォルダー名はエンコードされたパス ( C--Users-Name-... 、
-Users-name-... ) となり、新しいホームに一致するように名前が変更されます。
マージのみを行い、決して上書きしないでください。新しいファイルが追加されます。続けたチャット
ここの古いコピーは
長い方のクリーンなプレフィックスがそこにあります。同じチャットが両方で変更された場合
コンピュータが独立している場合、ローカル バージョンが常に優先され、他のバージョンが優先されます。
は .claude/porter/conflicts-... に保存されているため、何も失われません。を実行する
同じインポートを 2 回実行しても何も変わりません。インポートごとにバックアップを含むログが書き込まれます。
python3 session_porter.py undo はすべてを元に戻します。
オプションの原点メモ。 「新しいチャットの送信元をマークする」にチェックを入れます。
解凍すると、このコンピュータにとって初めてのチャットはすべてメモに記録されます。
タイトル (ログイン フローを修正する (サムから) など)。チャットは別の場所で継続されただけです
自分の名前を保持します。
断片化された会話が再び結合されます。自動圧縮とフォークによる分割
多くのトランスクリプト ファイル間でチャットします。ポーターは記録されたリンクをたどるので、アプリは
会話ごとに 1 つのチャットを表示します。
資格情報は決して転送されません。新しいコンピュータに通常どおりサインインすると、
アカウント ID フィールドはバンドルから削除されます。
Claude アプリがサインインしたことのないコンピューターで解凍した後、
アプリを 1 回実行し、Porter で [セットアップを完了] をクリックします。
Windows に移動されたバンドルでは、スラッシュ パス ( C:/Users/Name/... ) が使用されます。これらは
Windows では有効であり、同一のエンコードされたフォルダー名が生成されますが、これらを使用することはできません。
ネストされた JSON エスケープによって破損します。
すべてはコマンドラインからも機能します:export DIR 、inspect ZIP 、
import ZIP 、 register 、 undo 、および selftest (組み込みのテストスイート、
あらゆる方向をカバーします）。
Session Porter は独立したツールであり、Anthropic とは提携していません。
クロード コードの会話をコンピューター間で移動します。マック

および Windows、任意の方向。ファイルは 1 つで、依存関係はありません。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Move your Claude Code conversations between computers. Mac and Windows, any direction. One file, no dependencies. - ZenKhalil/session-porter

GitHub - ZenKhalil/session-porter: Move your Claude Code conversations between computers. Mac and Windows, any direction. One file, no dependencies. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
ZenKhalil
/
session-porter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .gitignore .gitignore LICENSE LICENSE README.md README.md Session Porter.bat Session Porter.bat Session Porter.command Session Porter.command session_porter.py session_porter.py View all files Repository files navigation
Moves your Claude Code conversations between computers: Mac to Mac, Mac to Windows,
Windows to Mac, or Windows to Windows. One file, no installation, and nothing ever
leaves your machines.
Mac: double-click Session Porter.command (or run python3 session_porter.py )
Windows: double-click Session Porter.bat (or run py session_porter.py )
Click Pack this computer . You get one bundle zip on the desktop, with a copy of
the app saved next to it.
Copy both files to the new computer using a USB stick, AirDrop, a network drive,
or anything else.
On the new computer , run the app the same way and click Unpack a bundle .
Porter finds the bundle on the desktop, in downloads, or on an attached drive by
itself, or you can paste the path. Check the summary, click Unpack here , then
quit the Claude app completely and reopen it. Your chats are there.
Windows needs Python 3.9+ from python.org (tick "Add to PATH"
during install). Macs already have it.
A plain zip of the .claude folder is not enough to move sessions. Porter covers
the gaps:
The app's chat list lives outside .claude (Mac:
~/Library/Application Support/Claude/claude-code-sessions , Windows:
%APPDATA%\Claude\claude-code-sessions ). Without it, imported sessions exist on
disk but appear nowhere in the app. Porter packs it, and rebuilds it from the
transcripts when it is missing, so it can also ingest a plain .claude zip made
without Porter.
Every path is rewritten for the new home folder, at all four JSON escape
depths, anchored on the old home prefix with boundaries on both sides. A username
that is also an ordinary word never corrupts your text, and a longer username
sharing the same prefix is never mistaken for the home.
Project folder names are encoded paths ( C--Users-Name-... ,
-Users-name-... ) and are renamed to match the new home.
Merge only, never overwrite. New files are added. A chat you continued on
the other computer has its new turns picked up, because the old copy here is
a clean prefix of the longer one there. If the same chat changed on both
computers independently, your local version always wins and the other version
is parked in .claude/porter/conflicts-... so nothing is lost. Running the
same import twice changes nothing. Every import writes a log with backups,
and python3 session_porter.py undo puts everything back.
Optional origin notes. Tick "Mark where new chats came from" when
unpacking and every chat that is new to this computer gets a note in its
title, like Fix login flow (from sam) . Chats merely continued elsewhere
keep their own names.
Fragmented conversations are reunited. Auto-compaction and forks split one
chat across many transcript files. Porter follows the recorded links so the app
shows one chat per conversation.
Credentials never travel. You sign in normally on the new computer, and
account identity fields are stripped from the bundle.
After unpacking on a computer where the Claude app has never signed in, open the
app once, then click Finish setup in Porter.
Bundles moved to Windows use forward-slash paths ( C:/Users/Name/... ). These are
valid on Windows and produce identical encoded folder names, and they cannot be
corrupted by nested JSON escaping.
Everything also works from the command line: export DIR , inspect ZIP ,
import ZIP , register , undo , and selftest (the built-in test suite,
covering every direction).
Session Porter is an independent tool and is not affiliated with Anthropic.
Move your Claude Code conversations between computers. Mac and Windows, any direction. One file, no dependencies.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
