---
source: "https://github.com/jaketothepast/codetutor"
hn_url: "https://news.ycombinator.com/item?id=48449430"
title: "Show HN: AI Pair Programmer for Emacs"
article_title: "GitHub - jaketothepast/codetutor: An AI Pair Programmer, that teaches you to code as you write, for Emacs · GitHub"
author: "jakewindle47"
captured_at: "2026-06-08T18:56:00Z"
capture_tool: "hn-digest"
hn_id: 48449430
score: 1
comments: 0
posted_at: "2026-06-08T18:34:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Pair Programmer for Emacs

- HN: [48449430](https://news.ycombinator.com/item?id=48449430)
- Source: [github.com](https://github.com/jaketothepast/codetutor)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T18:34:21Z

## Translation

タイトル: Show HN: Emacs 用 AI ペア プログラマー
記事のタイトル: GitHub - jaketothepast/codetutor: Emacs 用に、書きながらコードを書くことを教える AI ペア プログラマー · GitHub
説明: Emacs 用の、書きながらコーディングを教える AI ペア プログラマー - jaketothepast/codetutor
HN テキスト: AI を使用してペア プログラマを構築しましたが、現在、このパッケージを使用して新しい言語とパラダイムをゆっくりと進めようとしています。自分のプログラミングスキルの低下に気づき、なぜペアプログラマーとしてAIを使えないのか疑問に思っていました。なぜ私が行っていることを肩越しに監視し、変更を提案できるのでしょうか?そうすれば、問題を解決しながらも苦労しながら、実際に何かを学ぶことができます。これまでのところ結果は良好です:)

記事本文:
GitHub - jaketothepast/codetutor: Emacs 用に、書きながらコードを教える AI ペア プログラマー · GitHub
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
ジェイケト過去
/
コードチューター
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション操作

ション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット spec spec test test .gitignore .gitignore PROJECT.md PROJECT.md README.md README.md codetutor-pkg.el codetutor-pkg.el codetutor.el codetutor.el すべてのファイルを表示 リポジトリ ファイルのナビゲーション
CodeTutor は、コードを書きながら学習できる Emacs パッケージです。右側の講師パネルを開き、ファイルの保存を監視し、プロジェクトのコンテキストを収集し、ローカルの AI アシスタントにシニア/スタッフ エンジニアのペア プログラミング講師のように応答するよう依頼します。
重要な境界: CodeTutor はコードの作成を支援します。プロジェクト ファイルに自動的に書き込まれるわけではありません。
保存後に何が変更されたかを確認する
フィードバックの背後にある概念を説明する
コンパクトなコードサンプルを示す
最近の会話ターンを使用してフォローアップの質問に答える
耐久性のあるアーキテクチャに関するメモを .codetutor/ARCHITECTURE.md に保存します。
ファイル全体の置換を生成する
正確なタスクにすぐに貼り付けられる完全な実装を提供します
これは初期のローカル パッケージです。これは、ローカル コーデックスと pi バックエンドを備えたストック Emacs と Doom Emacs 用に設計されています。
オプションですが推奨: ツリーシッターサポートが組み込まれた Emacs 29+
ツリーシッターが利用できない場合、可能な限り imenu を使用することで、CodeTutor は正常に機能を低下させます。
(リストに追加 'load-path " /Users/jacobwindle/Projects/codetutor " )
( 'codetutor が必要です)
( setq codetutor-backend 'auto )
(コードチューターモード 1 )
ドゥーム Emacs
; ; ~/.config/doom/packages.el
(パッケージ! codetutor
:recipe ( :local-repo " ~/Projects/codetutor " ))
設定します:
; ; ~/.config/doom/config.el
(パッケージを使用してください! codetutor
:コマンド (codetutor-open
コードチューター-次は何ですか
コードチューターの質問
コードチューターのフォローアップ
codetutor-リフレッシュ-アーキテクチャ-メモリ)
:初期化
( setq codetutor-backend 'auto
codetutor-open-on-enable nil
codetutor-start-session-on-open t
コードチューター再

保存時に表示 )
:config
(コードチューターモード 1 ))
次に、次を実行します。
~ /.config/emacs/bin/doom sync
同期後に Emacs を再起動します。
コマンド
キーバインド
何をするのか
コードチューターモード
なし
グローバル CodeTutor フックを有効または無効にします。
コードチューターオープン
C-c to o
サイドパネルを開いてプロジェクトの評価を開始します。
コードチューター-次は何ですか
C-ctn
単一の最善の次のステップを尋ねます。
コードチューターの質問
C-CTA
現在のファイル/プロジェクトのコンテキストをミニバッファーからプロンプト表示します。
コードチューターのフォローアップ
C-ctf
最近のターンを使用して、前の回答についてのフォローアップを尋ねます。
codetutor-リフレッシュ-アーキテクチャ-メモリ
C-Ctm
家庭教師に耐久性のある建築ノートを更新するよう依頼します。
仕組み
CodeTutor には、起動評価、レビューの保存、手動プロンプト、およびフォローアップの 4 つのコア ループがあります。
M-x codetutor-open を実行すると、CodeTutor は次のようになります。
読み取り専用バックエンドリクエストを開始します。
パネルを Status: Thinking に置き換えます。
そのステータスを講師の最終的な回答に置き換えます。
スタートアップの答えは、どこから始めるべきか、最初に何を学ぶべきか、コードを書く前にエンジニアリング上のどのような判断が重要であるかを示している必要があります。
codetutor-mode と codetutor-review-on-save が有効になっている場合、CodeTutor は Emacs にフックして保存します。
before-save-hook は、現在のディスク上のファイルを読み取ります。
after-save-hook は、以前のディスク上の内容と保存されたバッファ テキストを比較します。
CodeTutor は統合された diff を構築します。
CodeTutor は、差分とプロジェクト コンテキストをバックエンドに送信します。
サイドパネルには「ステータス: 思考中」が表示されます。
サイドパネルは最終的なティーチング応答に置き換えられます。
保存レビューは差分に比例します。講師はコンセプト、リスク、アーキテクチャ、テスト、そして次の動きに焦点を当てる必要があります。
M-x codetutor-ask プロンプトをミニバッファーから取得します。これには次のものが含まれます。
バックエンドは、サポートされている場合、読み取り専用ツールを通じて他のプロジェクト ファイルを検査/検索できます。
M-x codetutor-follow-up は、前の回答について質問します。

えーっと。サイドパネルにプロンプ​​トは表示されません。
フォローアップには最近の個人的な会話が含まれるため、次のような質問をすることができます。
もう少し小さな例を見せていただけますか?
または:
テストの形状はどのようなものになるでしょうか?
フォローアップには現在のファイルとプロジェクトのコンテキストが含まれているため、講師は以前の回答を現在の状況に結び付けることができます。
M-x codetutor-what-next は、利用可能なコンテキストを検査し、最適な次のステップを 1 つ推奨するように講師に依頼します。
これは、実装スライスの間にいて、次に何をすべきかについて上級エンジニアの判断が必要な場合に便利です。
家庭教師は、耐久性のある建築の観察をフェンスで囲まれたブロックに含めるように求められます。
``` コードチューターメモリ
- 境界: エディタ統合はコンテキスト収集を所有します。バックエンドは個別指導を所有します。
「」
CodeTutor はこれらの行を抽出し、次の行に新しい行を追加します。
.codetutor/ARCHITECTURE.md
このメモリ ファイルは、CodeTutor が自動的に書き込む唯一のプロジェクト ファイルです。
CodeTutor は、いくつかのローカル ソースからプロンプトを構築します。
オープンバッファコンテキストには上限が設定されているため、大規模な Emacs セッションがバックエンドを圧倒することはありません。
Codex バックエンドは、codex exec を非対話的に使用します。これは次のように構成されています。
--output-last-message パネルに最終回答のみが表示されるようにする
コマンドは大まかに次のように構築されます。
コーデックス\
--サンドボックス読み取り専用 \
--ask-for-approval 決して\
--検索\
実行\
-C " $PROJECT_ROOT " \
--skip-git-repo-check \
--色は決して\
--一時的な\
--output-last-message " $TEMP_FILE " \
-
円周率
pi バックエンドは、読み取り専用ツールのみを備えた非対話型印刷モードを使用します。
pi --print --tools 読み取り、grep、検索、ls
ファイルを検査することはできますが、ファイルの書き込みや編集は行わないでください。
サイドパネルは意図的にチャットの記録ではありません。
家庭教師は、目に見える答えを以下で囲むように求められます。
< codetutor-answer >
...
</ codetutor-answer >
CodeTutor は、そのブロックが存在する場合は抽出し、ar を削除します

可視パネル出力からのアーキテクチャ メモリ ブロック。
組み込みのプロンプトは講師に次のことを指示します。
ユーザーを解決策に導く
アイデアを明確にする場合は、コンパクトなコード サンプルを含めます。
最終的なコードを引き渡す代わりに例を適応させる方法を説明する
パッチ、ファイル全体の置き換え、すぐに貼り付けられる完全な実装を回避します。
CodeTutor の優れた出力は、オートコンプリート エンジンのようなものではなく、上級エンジニアがあなたとペアを組んでいるように感じる必要があります。
( setq codetutor-backend 'auto ) ; ; 'auto、'codex、または 'pi
( setq codetutor-model nil ) ; ; nil はバックエンドのデフォルトを使用します
( setq codetutor-window-width 84 )
( setq codetutor-review-on-save t )
( setq codetutor-enable-web-search t )
( setq codetutor-include-open-buffers-on-save t )
コンテキストの制限:
( setq codetutor-max-project-context-bytes 80000 )
( setq codetutor-max-current-file-bytes 50000 )
( setq codetutor-max-diff-bytes 60000 )
( setq codetutor-max-open-buffers 12 )
( setq codetutor-max-open-buffer-bytes 20000 )
( setq codetutor-max-open-buffer-context-bytes 80000 )
( setq codetutor-max-conversation-turns 8 )
( setq codetutor-max-conversation-bytes 30000 )
即時のカスタマイズ:
(setq codetutor-system-prompt
「あなたは私の読み取り専用スタッフエンジニアの家庭教師です...」)
安全境界線
CodeTutor には 2 つの保護層があります。
プロンプトレベルの境界は、講師にファイルの編集やパッチの作成を行わないように指示します。
バックエンド コマンド境界では、使用可能な場合は読み取り専用モード/ツールが使用されます。
Codex の場合、パッケージは読み取り専用のサンドボックスを使用します。 pi の場合、読み取り/検索/リスト ツールのみが有効になります。
CodeTutor 自体が実行する自動書き込みは .codetutor/ARCHITECTURE.md のみです。
パネルにバックエンドエラーが表示される
コーデックスドクター
codex exec --help
pi --ヘルプ
インストールされているバージョンのコマンドライン フラグが CodeTutor が使用するものと一致していることを確認します。
答えだけではなくプロンプト/コンテキストが表示されます

CodeTutor は Codex の --output-last-message ファイルを優先し、存在する場合は <codetutor-answer> ブロックを抽出します。バックエンドが依然としてコンテキストをエコーする場合は、Codex バックエンドを使用するか、codetutor-system-prompt を調整して出力コントラクトを強化します。
( setq codetutor-max-diff-bytes 30000 )
( setq codetutor-max-open-buffers 6 )
( setq codetutor-max-open-buffer-context-bytes 40000 )
レビューを自動保存したくない
( setq codetutor-review-on-save nil )
codetutor-ask 、 codetutor-follow-up 、および codetutor-what-next は引き続き使用できます。
emacs --batch -L 。 -l test/codetutor-test.el -fert-run-tests-batch-and-exit
バイトコンパイル:
emacs --batch -L 。 -f バッチバイトコンパイル codetutor.el
コミットする前に、生成されたバイトコードを削除します。
codetutor.el パッケージの実装
codetutor-pkg.el パッケージのメタデータ
PROJECT.md CodeTutor によって使用されるプロジェクトの指示
spec/initial-behavior.md 初期動作仕様
test/codetutor-test.el ERT 回帰テスト
README.md ユーザーおよび開発者ガイド
について
Emacs 向けに、書きながらコーディングを教える AI ペア プログラマー
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

An AI Pair Programmer, that teaches you to code as you write, for Emacs - jaketothepast/codetutor

I built a pair programmer, using AI, but now am trying to go slow with new languages and paradigms using this package. I have noticed a decline in my programming skills, and wondered why I couldn't use AI as a pair programmer. Why can it watch what I'm doing over my shoulder and suggest changes? That way I can struggle through problems still and actually learn something. Results have been good so far :)

GitHub - jaketothepast/codetutor: An AI Pair Programmer, that teaches you to code as you write, for Emacs · GitHub
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
jaketothepast
/
codetutor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits spec spec test test .gitignore .gitignore PROJECT.md PROJECT.md README.md README.md codetutor-pkg.el codetutor-pkg.el codetutor.el codetutor.el View all files Repository files navigation
CodeTutor is an Emacs package for learning while you code. It opens a right-side tutor panel, watches file saves, gathers project context, and asks a local AI assistant to respond like a senior/staff engineer pair-programming tutor.
The important boundary: CodeTutor helps you write the code. It does not write into your project files for you.
review what changed after a save
explain the concept behind feedback
show compact illustrative code samples
answer follow-up questions using recent conversation turns
keep durable architecture notes in .codetutor/ARCHITECTURE.md
produce full-file replacements
hand you a complete ready-to-paste implementation for the exact task
This is an early local package. It is designed for stock Emacs and Doom Emacs, with local codex and pi backends.
Optional but recommended: Emacs 29+ with built-in tree-sitter support
CodeTutor degrades gracefully when tree-sitter is not available by using imenu where possible.
( add-to-list 'load-path " /Users/jacobwindle/Projects/codetutor " )
( require 'codetutor )
( setq codetutor-backend 'auto )
(codetutor-mode 1 )
Doom Emacs
; ; ~/.config/doom/packages.el
(package! codetutor
:recipe ( :local-repo " ~/Projects/codetutor " ))
Configure it:
; ; ~/.config/doom/config.el
(use-package! codetutor
:commands (codetutor-open
codetutor-what-next
codetutor-ask
codetutor-follow-up
codetutor-refresh-architecture-memory)
:init
( setq codetutor-backend 'auto
codetutor-open-on-enable nil
codetutor-start-session-on-open t
codetutor-review-on-save t )
:config
(codetutor-mode 1 ))
Then run:
~ /.config/emacs/bin/doom sync
Restart Emacs after syncing.
Command
Keybinding
What It Does
codetutor-mode
none
Enables/disables global CodeTutor hooks.
codetutor-open
C-c t o
Opens the side panel and starts a project assessment.
codetutor-what-next
C-c t n
Asks for the single best next step.
codetutor-ask
C-c t a
Prompts from the minibuffer with current file/project context.
codetutor-follow-up
C-c t f
Asks a follow-up about the previous answer using recent turns.
codetutor-refresh-architecture-memory
C-c t m
Asks the tutor to refresh durable architecture notes.
How It Works
CodeTutor has four core loops: startup assessment, save review, manual prompt, and follow-up.
When you run M-x codetutor-open , CodeTutor:
Starts a read-only backend request.
Replaces the panel with Status: thinking .
Replaces that status with the final tutor answer.
The startup answer should tell you where to begin, what to learn first, and what engineering judgment matters before writing code.
When codetutor-mode and codetutor-review-on-save are enabled, CodeTutor hooks into Emacs saves:
before-save-hook reads the current on-disk file.
after-save-hook compares the previous on-disk contents with the saved buffer text.
CodeTutor builds a unified diff.
CodeTutor sends the diff plus project context to the backend.
The side panel shows Status: thinking .
The side panel is replaced with the final teaching response.
Save reviews are proportional to the diff. The tutor should focus on concept, risk, architecture, tests, and one next move.
M-x codetutor-ask prompts from the minibuffer. It includes:
The backend may inspect/search other project files through read-only tools when supported.
M-x codetutor-follow-up asks a question about the previous answer. It does not display the prompt in the side panel.
Follow-ups include recent private conversation turns, so you can ask things like:
Can you show me a smaller example?
or:
What would the test shape look like?
The follow-up still includes current file and project context, so the tutor can connect the prior answer to where you are now.
M-x codetutor-what-next asks the tutor to inspect available context and recommend one best next step.
This is useful when you are between implementation slices and want a senior engineer's judgment on what to do next.
The tutor is asked to include durable architecture observations in a fenced block:
``` codetutor-memory
- Boundary: The editor integration owns context gathering; the backend owns tutoring.
```
CodeTutor extracts those lines and appends new ones to:
.codetutor/ARCHITECTURE.md
That memory file is the only project file CodeTutor writes automatically.
CodeTutor builds a prompt from several local sources.
Open-buffer context is capped so a large Emacs session does not overwhelm the backend.
The Codex backend uses codex exec non-interactively. It is configured with:
--output-last-message so the panel displays only the final answer
The command is built roughly like this:
codex \
--sandbox read-only \
--ask-for-approval never \
--search \
exec \
-C " $PROJECT_ROOT " \
--skip-git-repo-check \
--color never \
--ephemeral \
--output-last-message " $TEMP_FILE " \
-
pi
The pi backend uses non-interactive print mode with only read-only tools:
pi --print --tools read,grep,find,ls
It can inspect files, but should not write or edit them.
The side panel is intentionally not a chat transcript.
The tutor is asked to wrap visible answers in:
< codetutor-answer >
...
</ codetutor-answer >
CodeTutor extracts that block when present and strips architecture memory blocks from the visible panel output.
The built-in prompt tells the tutor to:
guide the user toward the solution
include compact code samples when they clarify the idea
explain how to adapt examples instead of handing over final code
avoid patches, full-file replacements, and complete ready-to-paste implementations
Good CodeTutor output should feel like a senior engineer pairing with you, not like an autocomplete engine.
( setq codetutor-backend 'auto ) ; ; 'auto, 'codex, or 'pi
( setq codetutor-model nil ) ; ; nil uses backend default
( setq codetutor-window-width 84 )
( setq codetutor-review-on-save t )
( setq codetutor-enable-web-search t )
( setq codetutor-include-open-buffers-on-save t )
Context limits:
( setq codetutor-max-project-context-bytes 80000 )
( setq codetutor-max-current-file-bytes 50000 )
( setq codetutor-max-diff-bytes 60000 )
( setq codetutor-max-open-buffers 12 )
( setq codetutor-max-open-buffer-bytes 20000 )
( setq codetutor-max-open-buffer-context-bytes 80000 )
( setq codetutor-max-conversation-turns 8 )
( setq codetutor-max-conversation-bytes 30000 )
Prompt customization:
( setq codetutor-system-prompt
" You are my read-only staff engineer tutor... " )
Safety Boundaries
CodeTutor has two layers of protection:
Prompt-level boundaries tell the tutor not to edit files or produce patches.
Backend command boundaries use read-only modes/tools where available.
For Codex, the package uses a read-only sandbox. For pi, it only enables read/search/list tools.
The only automatic write CodeTutor performs itself is .codetutor/ARCHITECTURE.md .
The panel shows backend errors
codex doctor
codex exec --help
pi --help
Confirm the command-line flags in your installed version match what CodeTutor uses.
I see prompt/context instead of just the answer
CodeTutor prefers Codex's --output-last-message file and extracts <codetutor-answer> blocks when present. If a backend still echoes context, use the Codex backend or adjust codetutor-system-prompt to reinforce the output contract.
( setq codetutor-max-diff-bytes 30000 )
( setq codetutor-max-open-buffers 6 )
( setq codetutor-max-open-buffer-context-bytes 40000 )
I do not want automatic save reviews
( setq codetutor-review-on-save nil )
You can still use codetutor-ask , codetutor-follow-up , and codetutor-what-next .
emacs --batch -L . -l test/codetutor-test.el -f ert-run-tests-batch-and-exit
Byte compile:
emacs --batch -L . -f batch-byte-compile codetutor.el
Remove generated bytecode before committing:
codetutor.el Package implementation
codetutor-pkg.el Package metadata
PROJECT.md Project direction consumed by CodeTutor
spec/initial-behavior.md Initial behavior spec
test/codetutor-test.el ERT regression tests
README.md User and developer guide
About
An AI Pair Programmer, that teaches you to code as you write, for Emacs
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
