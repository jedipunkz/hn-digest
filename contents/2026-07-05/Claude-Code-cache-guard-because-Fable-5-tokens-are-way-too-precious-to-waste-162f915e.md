---
source: "https://github.com/yuliang615/claude-cache-guard"
hn_url: "https://news.ycombinator.com/item?id=48796616"
title: "Claude Code cache guard because Fable 5 tokens are way too precious to waste"
article_title: "GitHub - yuliang615/claude-cache-guard: Local-only Claude Code statusLine bridge for usage-aware session handoff. · GitHub"
author: "yuliang615"
captured_at: "2026-07-05T19:07:12Z"
capture_tool: "hn-digest"
hn_id: 48796616
score: 2
comments: 0
posted_at: "2026-07-05T18:28:20Z"
tags:
  - hacker-news
  - translated
---

# Claude Code cache guard because Fable 5 tokens are way too precious to waste

- HN: [48796616](https://news.ycombinator.com/item?id=48796616)
- Source: [github.com](https://github.com/yuliang615/claude-cache-guard)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T18:28:20Z

## Translation

タイトル: Fable 5 トークンが貴重すぎて無駄にできないため、クロード コード キャッシュ ガード
記事のタイトル: GitHub - yuliang615/claude-cache-guard: 使用状況を認識したセッション ハンドオフのためのローカル専用クロード コード statusLine ブリッジ。 · GitHub
説明: 使用状況を認識したセッションハンドオフのためのローカル専用のクロードコード statusLine ブリッジ。 - yuliang615/クロード-キャッシュ-ガード

記事本文:
GitHub - yuliang615/claude-cache-guard: 使用状況を認識したセッションハンドオフのためのローカル専用クロードコード statusLine ブリッジ。 · GitHub
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
ユリアン615
/
クロードキャッシュガード
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github/ workflows .github/ workflows bin bin docs docs scripts scripts src src test test .gitignore .gitignore ライセンス ライセンス README.md README.md README.zh-CN.md README.zh-CN.md README.zh-TW.md README.zh-TW.md package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
claude-cache-guard (CLI: ccg ) は、クロード コードのローカル専用の使用法ガードです。 5 時間の使用期間が終了する前に、Claude にコンパクトなハンドオフ ファイル ( next_session.md ) を作成させるため、新しい小さなセッションで作業を再開できます。キャッシュされていないトークンとして古い会話を再読み込みしたり、 /clear スタイルの詳細が失われることはありません。
理由: Claude Code は会話全体を毎ターン再送信し、通常はプロンプト キャッシュによってそれが吸収されますが、使用済みのウィンドウが終了するまで待機している間、キャッシュはコールドになります。その後、大規模なセッションを再開すると、実際の作業が開始される前に、会話全体がキャッシュされていない入力トークンとして再読み取りされ、使用量に対して全額請求されます。そして、ウィンドウは最初はあまり保持されませんでした。
ステップ 1 — インストール (マシンごとに 1 回)
npm install -g クロードキャッシュガード
CCGのインストール
インストーラは 1 つの質問をします。
クロードは 5 時間の使用率が何パーセントになったら next_session.md の準備を開始する必要がありますか? [90]:
Enter を押してデフォルト ( 90 ) を受け入れるか、1 ～ 99 の数字を入力します。後でプロジェクトごとに変更できます。
インストール後、Claude Code を再起動します。終了して、claude を再度実行します (または、アプリで新しいセッションを開きます)。これにより /ccg* スラッシュ コマンドがロードされ、すべてのプロジェクトで機能します。 /clear だけでは十分ではありません。会話はクリアされますが、コマンドはリロードされません。
git clone https://github.com/yuliang615/claude-cache-guard.git && cd claude-cache-guard && ./scripts/install.sh
ステップ2

— プロジェクトに対して有効にする
プロジェクトのディレクトリで、claude を起動し、次のコマンドを実行します。
/ccgenable
以上です。このプロジェクトではハンドオフ ワークフローがアクティブになりました。クロード コード v2.1.169 以降では、すぐに有効になります。古いバージョンでは、完全に起動するために新しいセッションを開始します (コマンドはこれに関するメモを出力します)。ターミナルの方がいいですか？プロジェクト ディレクトリ内の ccg Enable も同じことを行います。実際に何を設定するかについては、以下の「仕組み」で説明します。
5 時間の使用量がしきい値に達すると、CCG はクロードにコンパクトな next_session.md を作成するように依頼し、目標を停止します。ここでは何もする必要はありません。以下のプロジェクトごとの設定で、しきい値と警告の動作をカスタマイズします。
ステップ 4 — 作業のバックアップを選択する
ウィンドウがリセットされたら、新しいクロード コード セッションを開き (または現在のセッションを /clear)、次のコマンドを実行します。
/ccgresume
クロードはハンドオフを読み取り、作業内容を自動的に選択してバックアップします。古いセッションを再読み取りする必要はありません。
新しいセッションで /ccgresume を実行します。古い大規模なセッションに戻ると、会話全体がキャッシュされていないトークンとして再読み取りされます。まさに、ハンドオフによって節約されたコストと同じです。
プロジェクトを有効にした後、クロード コード内の /ccgconfig スラッシュ コマンドを使用して、ハンドオフしきい値と警告動作をカスタマイズします。引数なしで実行すると、現在の値が表示され、両方の設定 (90 / 95 / 97 % の 5 時間しきい値の選択 (「その他」は 1 ～ 99 のいずれかを受け入れます)) と警告モードのネイティブ選択メニューがポップアップ表示されます。選択すると、それらが適用されます。
/ccgconfig
直接またはスクリプトによる変更の場合は、代わりにフラグを渡し、メニューをスキップします。 (ccg 設定ターミナル コマンドは同じフラグを受け取ります。! ccg 設定 ... クロード コード内では、モデルが関与せずに即座に実行されます。)
デフォルトのしきい値は 90 % です。メニューを使用せずに単一プロジェクトの値を下げる (または上げる) には:
/ccgconfig --five-hour-warni

ng70
5 時間の使用量がこの割合に達すると、ハンドオフが開始されます。
--on-warning は、使用量がしきい値に達した場合の動作を制御します。次の 2 つのモードがあります。
auto (デフォルト) — CCG は自動的にクロードにハンドオフ プロンプトを送信します。クロードは next_session.md を書き込み、ファイルが書き込まれると、ccg は現在の目標を停止します。ユーザーによる操作は必要ありません。次のセッションはハンドオフ ファイルから取得されます。
ask — CCG は、ハンドオフを作成して停止するか、続行するかを尋ねます。勝手に止まることはありません。ウィンドウごとに 1 回ではなく、しきい値を超えるごとに 1 回質問します。使用量がしきい値を下回って再びしきい値を超えると、再度尋ねることができます。予算がまだ残っており、すぐに決めたい場合に便利です。
/ccgconfig --on-warning ask # ask モードに切り替える
/ccgconfig --on-warning auto # 自動モードに戻す
両方のオプションを 1 回の呼び出しで組み合わせることができます。
/ccgconfig --five-hour-warning 70 --on-warning ask
オーバーライドはプロジェクトの .claude-cache-guard.json に保存され、すぐに有効になります。
カスタムハンドオフプロンプト (オプション)
これは何のためのものですか: 使用量がしきい値に近づくと、CCG はクロードにハンドオフ ファイル next_session.md の状況を書き留めるように要求します。ハンドオフを書き込むたびに Claude に処理してもらいたいことがある場合 (たとえば、常に最初にテストを実行する)、それらのリマインダーをプロジェクト内の .claude/ccg-handoff.md に置きます。それ以来、彼らは引き継ぎのたびに同行します。
最も簡単な方法は、次のように言うことです。ハンドオフ時に何を監視したいかをクロードに伝えます。そうすれば、.claude/ccg-handoff.md が編集されます。ファイル自体が、要求に応じてそれを実行できることをクロードに伝えます。手動で行う方が良いですか? ccg Enable はすでに .claude/ccg-handoff.md を作成しています。上部にはそれ自体を説明するコメント ブロックがありますが、無視されます。これはリマインダーとしてカウントされません。

その下に独自の行を 1 行につき 1 つのリマインダーとして追加します。
ファイルを壊してしまった場合でも、クリーン コピーがファイルのすぐ隣にある .claude/ccg-handoff.md.bak に保存されています。それをコピーして元に戻すと、最初からやり直すことができます。
それが有効になったことを確認するには、 /ccgstatus を実行し、ハンドオフ ガイダンス: active を探します。空のファイル (またはコメント ブロックのみのファイル) は無視されます。リマインダーは CCG の標準的な引き継ぎ指示の後に追加されるため、リマインダーが置き換えられることはありません。そのため、ここに何を書いても再開フローが中断されることはありません。
ccg install は、すべてのプロジェクトで使用できるスラッシュ コマンドのセットを ~/.claude/commands/ に配置します。それぞれの名前は ccg と、ハイフンなしでラップされる CLI サブコマンドです。読み取り専用コマンドは、コマンドが展開された瞬間に ccg 呼び出しを事前に実行し、出力を軽量モデルに渡してフォーマットするため、応答が速く、トークンをほとんど使用しません。
各コマンドとフラグについてはリファレンスを参照してください。
グローバルデフォルトは ~/.claude/cache-guard/config.json にあります。プロジェクトごとのオーバーライドは、各プロジェクトの .claude-cache-guard.json にあります。優先順位は、組み込みデフォルト < グローバル設定 < プロジェクト オーバーライド < CLI フラグです。
プロジェクトの有効な設定 (5 時間のしきい値、その発生元、および現在の警告モード) を確認するには、 /ccgstatus を実行します。それらを変更するには、上記のプロジェクトごとの設定で説明したように、 /ccgconfig を使用します。 (ターミナルに相当するもの: ccg status および ccg settings ; ccg config show は生のグローバル設定 JSON を出力します。) 完全な設定スキーマとマージ ルールについては、リファレンスを参照してください。
ガードは、source 、 updated_at 、 model 、 context_window 、 Five_hour 、および seven_day のみを保持する、許可リストに登録された単一のファイル ~/.claude/usage-state.json を書き込みます。トークン、認証、トランスクリプト、プロンプト、またはツール入力は決して保存されず、マシンから何も残されません。正確なスキーマについてはリファレンスを参照してください。
Claude Co で /ccgusage を実行します

de (またはターミナルでの CCG の使用法):
更新日: 2026-06-13T12:00:00.000Z
モデル：オーパス4.6
コンテキスト: 42%
5 時間: 75% リセット 2026-06-13T17:00:00Z
7d: 31% リセット 2026-06-18T09:00:00Z
クロード コードに使用可能な rate_limits が含まれていない場合、使用状況と医師は、フィールドを use-state.json に追加せずにそのように指示します。
何か問題が発生した場合は、クロード コードで /ccgdebug を実行します。読み取り専用の診断が実行され、検出された問題があれば説明されます。 (同等のターミナル: CCG ドクターおよび CCG デバッグ。)
出力は状態とタイムスタンプのみであり、statusLine コマンド テキスト、設定値、資格情報は含まれません。ドクター チェックは、statusLine がこのガードを指しているかどうか、~/.claude/usage-state.json が更新されているかどうか、有効なプロジェクトのフックがインストールされているかどうか、フック エラーが記録されているかどうかなど、一般的なセットアップの問題にフラグを立てます。完全なコマンド リストについては、ccg help を実行してください。
CCGのアンインストール
ccg アンインストールにより、マシンは ccg インストール前の状態に戻ります。statusLine が復元され、グローバル設定と使用状況がクリアされ、グローバル /ccg* スラッシュ コマンドが削除されます (同じ名前で自分で作成したファイルはそのまま残されます)。ハンドオフ ファイルは、 ~/.claude/next-session/ の下の正しい場所に残ります。これは、インストール状態ではなく、ユーザーの作業であるためです。
--remove は、以前のものを復元する代わりに statusLine を削除し、 --restore-backup <path> は代わりに正確なバックアップを復元します。 --rmconfig は引き続き受け入れられますが、必要なくなりました。構成はデフォルトですでに削除されています。両方のフラグについてはリファレンスを参照してください。
ccg install は、グローバルなものを 1 つだけ変更します。 ~/.claude/settings.json をバックアップし、 statusLine フィールドのみにパッチを適用し、 /ccg* スラッシュ コマンドを ~/.claude/commands/ にコピーします。 (ステータス行自体は次回の更新時に有効になります。インストール後の再起動によってスラッシュ コマンドが読み込まれます。)

g enable は厳密にプロジェクト ローカルです。 Stop フックと PostToolBatch フックを .claude/settings.local.json にインストールし、プロジェクト構成を書き込み、スターター next_session.md と .claude/ccg-handoff.md リマインダー ファイルを作成します (上記のカスタム ハンドオフ プロンプトを参照)。
statusLine が更新されるたびに、ガードは標準入力からクロード コードの JSON を読み取り、許可リストに登録された使用法フィールドのみを抽出して、それらを ~/.claude/usage-state.json にアトミックに書き込み、以前のステータス ライン (またはコンパクトなデフォルト) をレンダリングします。プロジェクト フックはそのファイルを読み取り、5 時間の使用量がしきい値を超えるとハンドオフをトリガーします。 statusLine の接続と完全なフックの動作についてはリファレンスを参照してください。また、非対話型インストールの場合は ccg install のフラグを参照してください (しきい値はデフォルトで 90 に設定されています。 --five-hour-warning でオーバーライドされます)。
📖 完全なリファレンス: docs/REFERENCE.md — すべてのコマンドとフラグ、設定とデータのスキーマ、ハンドオフ形式、および完全なフック動作。完全なコマンドとフラグの概要については、ccg help を実行してください。
クロード コード プロンプトのキャッシュと圧縮に関するドキュメント (圧縮では、キャッシュされたプレフィックスが再利用されます)
使用状況を認識したセッションハンドオフのためのローカル専用のクロードコードステータスラインブリッジ。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
途中でエラーが発生しました

[切り捨てられた]

## Original Extract

Local-only Claude Code statusLine bridge for usage-aware session handoff. - yuliang615/claude-cache-guard

GitHub - yuliang615/claude-cache-guard: Local-only Claude Code statusLine bridge for usage-aware session handoff. · GitHub
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
yuliang615
/
claude-cache-guard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github/ workflows .github/ workflows bin bin docs docs scripts scripts src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md README.zh-TW.md README.zh-TW.md package.json package.json View all files Repository files navigation
claude-cache-guard (CLI: ccg ) is a local-only usage guard for Claude Code. Before your 5-hour usage window runs out, it has Claude write a compact handoff file ( next_session.md ) so you pick the work back up in a fresh, small session — no re-reading the old conversation as uncached tokens, no /clear -style loss of detail .
Why: Claude Code re-sends the whole conversation every turn, and prompt caching normally absorbs that — but the cache goes cold while you wait out a spent window. Resume a big session then, and the entire conversation is re-read as uncached input tokens, charged in full against your usage, before any real work starts — and the window didn't hold much to begin with.
Step 1 — Install (once per machine)
npm install -g claude-cache-guard
ccg install
The installer asks one question:
At what 5-hour usage percentage should Claude start preparing next_session.md? [90]:
Press Enter to accept the default ( 90 ), or type a number from 1 to 99 — you can change it per project later.
After installing, restart Claude Code — exit and run claude again (or open a new session in the app). That's what loads the /ccg* slash commands, which then work in every project. /clear is not enough: it clears the conversation but does not reload commands.
git clone https://github.com/yuliang615/claude-cache-guard.git && cd claude-cache-guard && ./scripts/install.sh
Step 2 — Enable for the project
In the project's directory, start claude and run:
/ccgenable
That's it — the handoff workflow is now active for this project. On Claude Code v2.1.169+ it takes effect right away; on older versions, start a new session for it to fully kick in (the command prints a note about this). Prefer the terminal? ccg enable in the project directory does the same thing. What it actually sets up is covered in How it works below.
When 5-hour usage reaches the threshold, ccg asks Claude to write a compact next_session.md and stops the goal — nothing for you to do here. Customize the threshold and warning behavior in Per-project settings below.
Step 4 — Pick the work back up
After the window resets, open a new Claude Code session (or /clear the current one) and run:
/ccgresume
Claude reads the handoff and picks the work back up automatically — no re-reading the old session.
Run /ccgresume in a fresh session. Going back to the old, big session re-reads the whole conversation as uncached tokens — exactly the cost the handoff just saved you.
After enabling a project, customize its handoff threshold and warning behavior with the /ccgconfig slash command inside Claude Code. Run it with no arguments and it shows the current values, then pops up a native selection menu for both settings — 5-hour threshold choices of 90 / 95 / 97 % ("Other" accepts any 1–99), and the warning mode. Pick, and they are applied for you:
/ccgconfig
For a direct or scriptable change, pass flags instead and skip the menu. (The ccg setting terminal command takes the same flags; ! ccg setting ... inside Claude Code runs it instantly with no model involvement.)
The default threshold is 90 %. To lower (or raise) it for a single project without the menu:
/ccgconfig --five-hour-warning 70
When 5-hour usage reaches this percentage, the handoff kicks in.
--on-warning controls what happens when usage hits the threshold. There are two modes:
auto (the default) — ccg sends Claude the handoff prompt automatically. Claude writes next_session.md , and once the file is written ccg stops the current goal. No user interaction is needed; the next session picks up from the handoff file.
ask — ccg asks you whether to write the handoff and stop, or keep going. It never stops on its own. It asks once per threshold crossing, not once per window: once usage falls back below the threshold and crosses it again, it can ask again. This is useful when you still have budget left and want to decide in the moment.
/ccgconfig --on-warning ask # switch to ask mode
/ccgconfig --on-warning auto # switch back to auto mode
Both options can be combined in one call:
/ccgconfig --five-hour-warning 70 --on-warning ask
Overrides are stored in the project's .claude-cache-guard.json and take effect immediately.
Custom handoff prompt (optional)
What this is for: when usage nears the threshold, CCG asks Claude to write down where things stand in the handoff file next_session.md . If there are things you want Claude to take care of every time it writes that handoff — say, always run the tests first — put those reminders in .claude/ccg-handoff.md inside the project. They ride along on every handoff from then on.
The easiest way is to just say it : tell Claude what you want watched at handoff time, and it edits .claude/ccg-handoff.md for you — the file itself tells Claude it may do that on your request. Prefer to do it by hand? ccg enable already created .claude/ccg-handoff.md , with a comment block at the top that explains itself and is ignored — it doesn't count as a reminder. Add your own lines below it, one reminder per line.
If you end up breaking the file, a clean copy is sitting right next to it at .claude/ccg-handoff.md.bak — copy it back over and you're starting fresh.
To check it took effect: run /ccgstatus and look for handoff guidance: active . An empty file (or one with just the comment block) is ignored. Your reminders are appended after CCG's standard handoff instructions — they never replace them — so nothing you write here can break the resume flow.
ccg install puts a set of slash commands in ~/.claude/commands/ , available in every project — each named ccg + the CLI subcommand it wraps, with no hyphen. Read-only commands pre-run their ccg call the moment the command expands and hand the output to a lightweight model to format, so they respond fast and use almost no tokens:
See the reference for every command and flag.
Global defaults live in ~/.claude/cache-guard/config.json ; per-project overrides live in each project's .claude-cache-guard.json . Priority is built-in defaults < global config < project overrides < CLI flags.
To see a project's effective settings — the 5-hour threshold, where it comes from, and the current warning mode — run /ccgstatus . To change them, use /ccgconfig , as covered in Per-project settings above. (Terminal equivalents: ccg status and ccg setting ; ccg config show prints the raw global config JSON.) See the reference for the full config schema and merge rules.
The guard writes a single allowlisted file, ~/.claude/usage-state.json , holding only source , updated_at , model , context_window , five_hour , and seven_day . No tokens, auth, transcripts, prompts, or tool inputs are ever stored, and nothing leaves your machine. See the reference for the exact schema.
Run /ccgusage in Claude Code (or ccg usage in a terminal):
updated: 2026-06-13T12:00:00.000Z
model: Opus 4.6
context: 42%
5h: 75% reset 2026-06-13T17:00:00Z
7d: 31% reset 2026-06-18T09:00:00Z
If Claude Code does not include usable rate_limits , usage and doctor say so without adding fields to usage-state.json .
When something looks off, run /ccgdebug in Claude Code — it runs both read-only diagnostics and explains any problem it finds. (Terminal equivalents: ccg doctor and ccg debug .)
The output is just state and timestamps — no statusLine command text, config values, or credentials. The doctor check flags the common setup problems: whether statusLine points to this guard, whether ~/.claude/usage-state.json is updating, whether the enabled project's hooks are installed, and whether any hook errors were recorded. Run ccg help for the full command list.
ccg uninstall
ccg uninstall puts your machine back the way it was before ccg install : the statusLine gets restored, the global config and usage state get cleared out, and the global /ccg* slash commands are removed (any file you wrote yourself under the same name is left alone). Handoff files stay right where they are, under ~/.claude/next-session/ , since that's your work, not install state.
--remove drops the statusLine instead of restoring the previous one, and --restore-backup <path> restores an exact backup instead. --rmconfig is still accepted but no longer needed — the config is already gone by default. See the reference for both flags.
ccg install changes exactly one global thing: it backs up ~/.claude/settings.json , patches only its statusLine field, and copies the /ccg* slash commands into ~/.claude/commands/ . (The status line itself goes live on the next refresh — the post-install restart is what loads the slash commands.) ccg enable is strictly project-local: it installs the Stop and PostToolBatch hooks in .claude/settings.local.json , writes the project config, and creates a starter next_session.md plus the .claude/ccg-handoff.md reminder file (see Custom handoff prompt above).
On every statusLine refresh the guard reads Claude Code's JSON from stdin, extracts only the allowlisted usage fields, writes them atomically to ~/.claude/usage-state.json , and then renders your previous status line (or a compact default). The project hooks read that file and trigger the handoff when 5-hour usage crosses the threshold. See the reference for the statusLine wiring and the full hook behavior, and ccg install 's flags for non-interactive installs (they default the threshold to 90 ; override with --five-hour-warning ).
📖 Full reference: docs/REFERENCE.md — every command and flag, the configuration and data schemas, the handoff format, and the complete hook behavior. Run ccg help for the full command and flag summary.
Claude Code prompt caching and compaction docs (compaction reuses the cached prefix)
Local-only Claude Code statusLine bridge for usage-aware session handoff.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error whil

[truncated]
