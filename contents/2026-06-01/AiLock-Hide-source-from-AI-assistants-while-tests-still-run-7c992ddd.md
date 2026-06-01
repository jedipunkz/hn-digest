---
source: "https://github.com/lo2589/AILOCK"
hn_url: "https://news.ycombinator.com/item?id=48345582"
title: "AiLock: Hide source from AI assistants while tests still run"
article_title: "GitHub - lo2589/AILOCK · GitHub"
author: "yoliliya"
captured_at: "2026-06-01T00:30:59Z"
capture_tool: "hn-digest"
hn_id: 48345582
score: 2
comments: 0
posted_at: "2026-05-31T13:37:53Z"
tags:
  - hacker-news
  - translated
---

# AiLock: Hide source from AI assistants while tests still run

- HN: [48345582](https://news.ycombinator.com/item?id=48345582)
- Source: [github.com](https://github.com/lo2589/AILOCK)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T13:37:53Z

## Translation

タイトル: AiLock: テストの実行中に AI アシスタントからソースを非表示にする
記事タイトル: GitHub - lo2589/AILOCK · GitHub
説明: GitHub でアカウントを作成して、lo2589/AILOCK の開発に貢献します。

記事本文:
GitHub - lo2589/AILOCK · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
lo2589
/
アイロック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード もっと開く ac

オプションメニュー フォルダーとファイル
6 コミット 6 コミット .ailock .ailock aloc aloc testrun testrun .gitignore .gitignore README.md README.md README_zh.md README_zh.md pyproject.toml pyproject.toml test_file.txt test_file.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コードをディスク上で暗号化したままにし、メモリ内でのみ復号化し、通常どおり実行します。
AiLock はファイルをその場で暗号化するため、ファイルシステムレベルの AI アクセス ( read_file 、 grep 、
cat 、コードベースのインデックス作成) はバイナリ暗号文のみを参照します。同時に、
開発者は、暗号化された Python コードを実行し、暗号化されたモジュールをインポートし、暗号化されたモジュールを読み取ることができます。
データ ファイルを管理し、制御されたプレーンテキスト ビューを通じてロックされたファイルを編集します。中央
アイデアはメモリのみの復号化です。平文は AiLock 内で実体化されます。
ランタイム プロセス。作業ツリーには書き戻されません。
ディスク: AI および通常のファイル リーダー用の暗号文。
ランタイム: 制御された実行プロセス内でのみプレーンテキスト。
ほとんどの暗号化ツールは保存中のファイルを保護しますが、暗号化が完了するまでコードは使用できなくなります。
復号化されてディスクに戻されます。 AiLock は、さまざまなワークフロー向けに構築されています。
AI 不透明性 : 作業ツリーを読み取るコーディング アシスタントは暗号文を参照します。
メモリのみの実行: ailock run は内部の暗号化された Python ファイルを復号化します
プロセスを削除し、ディスク上に平文を復元せずに実行します。
透過的なインポート: 暗号化されたモジュールは相互にインポートできます。
透過的なファイル I/O : open() 、 Path.read_text() 、および
Path.read_bytes() はランタイム内でプレーンテキストを返すことができます。
GUI プレーンテキスト ビューポート: ailock open により、開発者は検査および編集できるようになります
作業ツリーに平文を残さずにファイルを保存します。
回復パス: 暗号化されたバックアップとオプションの回復キーが回復に役立ちます
破損したファイルまたはパスワードを忘れたファイル。
pyproject.toml から自動的にインストールされるランタイム パッケージ:
argon2-cfi 、暗号化、および pyzipper
きん

エアロックオープン用のター。多くの Python インストールにバンドルされていますが、
一部の Linux ディストリビューションでは、python3-tk として個別にパッケージ化されています。
git clone https://github.com/lo2589/AILOCK.git
cd アイロック
pip インストール 。
編集可能な開発インストールの場合:
git clone https://github.com/lo2589/AILOCK.git
cd アイロック
pip install -e 。
コマンドを確認してください:
aillock --ヘルプ
コマンドが PATH 上にない場合は、モジュールのエントリ ポイントを使用します。
python -m aloc --help
クイックスタート
# ファイルをその場で暗号化します。
aillock ロックの秘密.py
# AI/ファイル ツールについては、暗号文を参照してください。
猫の秘密.py
grep "パスワード" 。
# コードは引き続き使用できます。
aillock show Secret.py
aillock run secret.py
エアロックが開いています。
# 必要に応じてディスク上の平文を復元します。
AILOCK ロック解除 Secret.py
重要なアイデア:
ailock lock app.py # app.py はディスク上の暗号文になります
ailock run app.py # app.py がメモリ内で復号化されて実行される
メモリのみの実行
aillock run がコア機能です。メモリ上のエントリファイルを復号化して実行します。
Python プロセス内のプレーンテキストを保存し、作業ツリー ファイルをそのまま残します。
暗号文。暗号化されたファイルの隣には平文のコピーは書き込まれません。
aillock は main.py を実行します
aillock run -m mypackage
aillock run app.py -- --port 8080
プログラムの実行中に、AiLock はフックをインストールするので、アプリケーション コードは
ファイルがプレーンであるかのように動作します。
ディスク上の暗号化された.py -> メモリ内で復号化 -> Python 内で実行/インポート
暗号化されたデータファイル -> メモリ内で復号化 -> open()/Path.read_text()
プログラム内では、AiLock 固有のコードは必要ありません。
jsonをインポートする
Secret_module インポート アルゴリズムから
open ( "config.json" ) を f として使用:
設定 = json 。負荷 ( f )
print (アルゴリズム(config))
Secret_module.py または config.json がロックされている場合、AiLock はそれを復号化します。
ファイルシステムに暗号文がまだ含まれている間、実行時に実行されます。
ファイルまたはディレクトリをその場で暗号化します。
エアロックロックの秘密

ぴー
ailock ロック src/
ailock ロック Secret.py --recovery
注:
ディレクトリは再帰的に処理されます。
すでにロックされているファイルはスキップされます。
プレーンテキストのバックアップは、暗号化された ZIP バックアップとして .ailock/backups/ に保存されます。
デフォルトでは。
--recovery は回復キーを出力します。別途保存してください。再度表示されることはありません。
平文をディスクに書き戻さずに、暗号化された Python コードを実行します。
aillock は main.py を実行します
aillock run -m mypackage
aillock run app.py -- --port 8080
ランタイムインターセプトレイヤー:
暗号化された Python モジュールのインポート フック
パッチを適用した pathlib.Path.read_text および pathlib.Path.read_bytes
ディレクトリの GUI プレーンテキスト ビューポート/エディタを開きます。
エアロックが開いています。
エアロックオープンsrc/
ロックされたファイルは表示のために復号化されます。保存すると暗号化されたコンテンツが書き戻されます
ディスク。
ファイルを変更せずに、復号化されたコンテンツを標準出力に出力します。
aillock show Secret.py
aillock は、secret.py を表示します。頭
aillock ロック解除 <パス>
ファイルまたはディレクトリを復号化し、ディスク上の平文に戻します。
AILOCK ロック解除 Secret.py
ailock ロック解除 src/ --backup
aillock リカバリ <ファイル>
--recovery で生成された回復キーを使用して、ロックされたファイルを回復します。
aillock 回復 Secret.py
アイロック フリーロック [パス]
制御されたプレーンテキスト アクセスのために stdin/stdout JSON-RPC ワークスペース サーバーを起動します。
エイロック フリーロック 。
リクエストの例:
{ "メソッド" : " list_files " 、 "params" : {}、 "id" : 1 }
{ "メソッド" : " read_file " 、 "params" : { "パス" : " main.py " }、 "id" : 2 }
{ "メソッド" : " grep " , "params" : { "パターン" : " TODO " }, "id" : 3 }
{ "メソッド" : " write_file " 、 "params" : { "パス" : " main.py " 、 "コンテンツ" : " ... " }、 "id" : 4 }
{ "メソッド" : "フラッシュ" 、 "パラメータ" : {}、 "id" : 5 }
その他のコマンド
aillockステータスファイル.py
エアロック忘れてください
aillock 忘れます --all
エアロック設定
ailock config バックアップ ディレクトリ /path/to/backups
ailock init --as aa
ailock init --as <name> は、ローカル ランチャーを

カスタムコマンド名。
これは、ロック解除コマンドを展開固有のものにしたい場合に便利です。
AiLock は、ファイルシステム レベルの AI アクセスをターゲットとしています。コーディングアシスタント向けに設計されています
通常の読み取りを通じてファイルを検査するインデクサー。そのモデルでは、ロックされています
ファイルからは暗号文のみが明らかになります。
AiLock は、十分な情報を得て実行できる地元の敵を阻止するとは主張していません。
任意のコマンド、プロセスメモリのキャプチャ、またはユーザーをだまして復号化させる
ファイル。より強力な分離を実現するには、AiLock とオペレーティング システムの実行を組み合わせます
ポリシー、プロセスの分離、および慎重なシークレットの取り扱い。
パスワード派生キーの Argon2id
認証暗号化用の ChaCha20-Poly1305
ファイルキーのパスワードラッピング
オプションの回復キーのラッピング
緊急時の回復のための暗号化された ZIP バックアップ
アロック/
cli.py コマンドライン インターフェイス
runner.py インメモリ実行エンジン
workspace.py 復号化されたワークスペース API と JSON-RPC ハンドラー
gui.py tkinter GUI エディター
crypto.py Argon2id および ChaCha20-Poly1305 ヘルパー
format.py ロックされたファイル形式パーサー/エンコーダー
fileops.py アトミック書き込みとバックアップ ヘルパー
Cache.py sudo スタイルのパスワード キャッシュ
manifest.py .ailock マニフェストとバックアップの管理
Recovery.py 回復キーのサポート
install.py カスタム コマンド名ランチャー
依存関係の概要
多くの Python インストールで提供される GUI 用の tkinter
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

Contribute to lo2589/AILOCK development by creating an account on GitHub.

GitHub - lo2589/AILOCK · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
lo2589
/
AILOCK
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .ailock .ailock aloc aloc testrun testrun .gitignore .gitignore README.md README.md README_zh.md README_zh.md pyproject.toml pyproject.toml test_file.txt test_file.txt View all files Repository files navigation
Keep code encrypted on disk, decrypt it only in memory, and still run it normally.
AiLock encrypts files in place so filesystem-level AI access ( read_file , grep ,
cat , codebase indexing) sees only binary ciphertext. At the same time,
developers can run encrypted Python code, import encrypted modules, read encrypted
data files, and edit locked files through controlled plaintext views. The central
idea is memory-only decryption : plaintext is materialized inside the AiLock
runtime process, not written back to the working tree.
Disk: ciphertext for AI and ordinary file readers.
Runtime: plaintext only inside the controlled execution process.
Most encryption tools protect files at rest, but make the code unusable until it
is decrypted back onto disk. AiLock is built for a different workflow:
AI-opacity : coding assistants that read the working tree see ciphertext.
Memory-only execution : ailock run decrypts encrypted Python files inside
the process and executes them without restoring plaintext on disk.
Transparent imports : encrypted modules can import each other.
Transparent file I/O : open() , Path.read_text() , and
Path.read_bytes() can return plaintext inside the runtime.
GUI plaintext viewport : ailock open lets the developer inspect and edit
files without leaving plaintext in the working tree.
Recovery path : encrypted backups and optional recovery keys help recover
damaged or forgotten-password files.
Runtime packages installed automatically from pyproject.toml :
argon2-cffi , cryptography , and pyzipper
tkinter for ailock open ; it is bundled with many Python installations, but
some Linux distributions package it separately as python3-tk
git clone https://github.com/lo2589/AILOCK.git
cd AILOCK
pip install .
For editable development installs:
git clone https://github.com/lo2589/AILOCK.git
cd AILOCK
pip install -e .
Check the command:
ailock --help
If the command is not on your PATH , use the module entry point:
python -m aloc --help
Quick Start
# Encrypt a file in place.
ailock lock secret.py
# AI/file tools see ciphertext.
cat secret.py
grep " password " .
# You can still use the code.
ailock show secret.py
ailock run secret.py
ailock open .
# Restore plaintext on disk when needed.
ailock unlock secret.py
The key idea:
ailock lock app.py # app.py becomes ciphertext on disk
ailock run app.py # app.py is decrypted in memory and executed
Memory-only Execution
ailock run is the core feature. It decrypts the entry file in memory, executes
the plaintext inside the Python process, and leaves the working-tree file as
ciphertext. No plaintext copy is written next to the encrypted file.
ailock run main.py
ailock run -m mypackage
ailock run app.py -- --port 8080
While the program is running, AiLock installs hooks so application code can
behave as if the files were plain:
encrypted .py on disk -> decrypt in memory -> exec/import inside Python
encrypted data file -> decrypt in memory -> open()/Path.read_text()
Inside your program, no AiLock-specific code is required:
import json
from secret_module import algorithm
with open ( "config.json" ) as f :
config = json . load ( f )
print ( algorithm ( config ))
If secret_module.py or config.json is locked, AiLock decrypts it for the
runtime while the filesystem still contains ciphertext.
Encrypt a file or directory in place.
ailock lock secret.py
ailock lock src/
ailock lock secret.py --recovery
Notes:
Directories are processed recursively.
Already locked files are skipped.
Plaintext backups are stored as encrypted ZIP backups under .ailock/backups/
by default.
--recovery prints a recovery key. Save it separately; it is not shown again.
Run encrypted Python code without writing plaintext back to disk.
ailock run main.py
ailock run -m mypackage
ailock run app.py -- --port 8080
Runtime interception layers:
import hook for encrypted Python modules
patched pathlib.Path.read_text and pathlib.Path.read_bytes
Open a GUI plaintext viewport/editor for a directory.
ailock open .
ailock open src/
Locked files are decrypted for display. Saving writes encrypted content back to
disk.
Print decrypted content to stdout without modifying the file.
ailock show secret.py
ailock show secret.py | head
ailock unlock <path>
Decrypt a file or directory back to plaintext on disk.
ailock unlock secret.py
ailock unlock src/ --backup
ailock recover <file>
Recover a locked file using a recovery key generated by --recovery .
ailock recover secret.py
ailock freelock [path]
Start a stdin/stdout JSON-RPC workspace server for controlled plaintext access.
ailock freelock .
Example requests:
{ "method" : " list_files " , "params" : {}, "id" : 1 }
{ "method" : " read_file " , "params" : { "path" : " main.py " }, "id" : 2 }
{ "method" : " grep " , "params" : { "pattern" : " TODO " }, "id" : 3 }
{ "method" : " write_file " , "params" : { "path" : " main.py " , "content" : " ... " }, "id" : 4 }
{ "method" : " flush " , "params" : {}, "id" : 5 }
Other Commands
ailock status file.py
ailock forget
ailock forget --all
ailock config
ailock config backup-dir /path/to/backups
ailock init --as aa
ailock init --as <name> installs a local launcher under a custom command name.
This is useful when you want the unlock command to be deployment-specific.
AiLock targets filesystem-level AI access. It is designed for coding assistants
and indexers that inspect files through ordinary reads. In that model, locked
files reveal only ciphertext.
AiLock does not claim to stop a fully informed local adversary who can run
arbitrary commands, capture process memory, or trick the user into decrypting
files. For stronger isolation, combine AiLock with operating-system execution
policy, process isolation, and careful secret handling.
Argon2id for password-derived keys
ChaCha20-Poly1305 for authenticated encryption
password wrapping for file keys
optional recovery-key wrapping
encrypted ZIP backups for emergency recovery
aloc/
cli.py command-line interface
runner.py in-memory execution engine
workspace.py decrypted workspace API and JSON-RPC handler
gui.py tkinter GUI editor
crypto.py Argon2id and ChaCha20-Poly1305 helpers
format.py locked-file format parser/encoder
fileops.py atomic writes and backup helpers
cache.py sudo-style password cache
manifest.py .ailock manifest and backup management
recovery.py recovery key support
install.py custom command-name launcher
Dependency Summary
tkinter for the GUI, provided by many Python installations
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
