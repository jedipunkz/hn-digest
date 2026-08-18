---
source: "https://code.claude.com/docs/en/changelog"
hn_url: "https://news.ycombinator.com/item?id=49352511"
title: "Claude Code 2.1.234 continues your session automatically once usage limit resets"
article_title: "Claude Code changelog - Claude Code Docs"
image: "https://claude-code.mintlify.app/_next/image?url=%2F_mintlify%2Fapi%2Fog%3Fdivision%3DGetting%2Bstarted%26appearance%3Dsystem%26title%3DClaude%2BCode%2Bchangelog%26description%3DRelease%2Bnotes%2Bfor%2BClaude%2BCode%252C%2Bincluding%2Bnew%2Bfeatures%252C%2Bimprovements%252C%2Band%2Bbug%2Bfixes%2Bby%2Bversion.%26logoLight%3Dhttps%253A%252F%252Fmintcdn.com%252Fclaude-code%252Fc5r9_6tjPMzFdDDT%252Flogo%252Flight.svg%253Ffit%253Dmax%2526auto%253Dformat%2526n%253Dc5r9_6tjPMzFdDDT%2526q%253D85%2526s%253D78fd01ff4f4340295a4f66e2ea54903c%26logoDark%3Dhttps%253A%252F%252Fmintcdn.com%252Fclaude-code%252Fc5r9_6tjPMzFdDDT%252Flogo%252Fdark.svg%253Ffit%253Dmax%2526auto%253Dformat%2526n%253Dc5r9_6tjPMzFdDDT%2526q%253D85%2526s%253D1298a0c3b3a1da603b190d0de0e31712%26primaryColor%3D%25230E0E0E%26lightColor%3D%2523D4A27F%26backgroundLight%3D%2523FDFDF7%26backgroundDark%3D%252309090B&w=1200&q=100"
author: "janpeuker"
captured_at: "2026-08-18T21:14:20Z"
capture_tool: "hn-digest"
hn_id: 49352511
score: 2
comments: 1
posted_at: "2026-08-18T20:50:37Z"
tags:
  - hacker-news
  - translated
---

# Claude Code 2.1.234 continues your session automatically once usage limit resets

- HN: [49352511](https://news.ycombinator.com/item?id=49352511)
- Source: [code.claude.com](https://code.claude.com/docs/en/changelog)
- Score: 2
- Comments: 1
- Posted: 2026-08-18T20:50:37Z

## Translation

タイトル: Claude Code 2.1.234 は、使用制限がリセットされるとセッションを自動的に継続します
記事のタイトル: Claude Code 変更履歴 - Claude Code Docs
説明: Claude Code のリリース ノート。バージョンごとの新機能、改善点、バグ修正が含まれます。

記事本文:
Claude Code 変更ログ - Claude Code Docs Documentation Index
/docs/llms.txt で完全なドキュメントのインデックスを取得します。
さらに探索する前に、このファイルを使用して利用可能なすべてのページを検出します。
メイン コンテンツにスキップ Claude Code Docs ホーム ページ 英語 検索... ⌘ K Ask Assistant Claude 開発者プラットフォーム
検索... ナビゲーション はじめに クロード コードの変更履歴 はじめに クロード コードによるビルド 管理 構成リファレンス エージェント SDK 新機能 リソース はじめに
指示と記憶を保存する
ページをコピー ページをコピー バージョンごとの新機能、改善点、バグ修正を含む、Claude Code のリリース ノート。
ページのコピー ページのコピー このページは、GitHub の CHANGELOG.md から生成されます。
clude --version を実行して、インストールされているバージョンを確認します。
2.1.234 2026 年 8 月 17 日
オプションの CLAUDE_CODE_PROJECT_DIR_NAME 環境変数を追加しました。各セッションに独自の構成ディレクトリを与えるホストは、プロジェクトごとのトランスクリプト ディレクトリの短い名前を選択できます。
selection:clear キーバインド アクションを追加しました。これにより、キーをバインドしてアプリ内テキストの選択をクリアできるようになりました。エージェントビューでも機能します
フッターとステータスラインに GitLab マージ リクエスト バッジを追加しました。GitLab リモートと認証された glab CLI を使用したリポジトリは、ドラフト/保留中/緑色の状態で MR !N を表示します。
Claude Code は、claude.ai の使用制限がリセットされたときにセッションを自動的に継続するようになりました。 /config でオフにします (「使用量制限で自動的に続行」)
クロードは現在、アカウントの電子メールはあなたを識別する目的のみに使用し、あなたが要求しない限り無関係のサービスに送信しないように指示されています。
セキュリティ: リモート ファイルの読み取り、セッションの復元、CLAUDE.md インクルード、ワークフロー スクリプトとファイルのアップロードで Windows NT 名前空間 ( \??\ ) パスが拒否されるようになり、残りの事前承認ファイル アクセスが NTLM 資格情報に対して強化されます。

l-リークベクトル
会話が圧縮された後、非常に長いセッションで自動モードが繰り返し再チェックされ、サンドボックス化されたコマンドのネットワーク アクセスが拒否される問題を修正しました。
バックグラウンド サブエージェント ツールの権限プロンプトに応答するときに、セッション スコープの権限の回答 (拒否を含む) がドロップされる問題を修正しました。
非ストリーミング フォールバック パス (通常はサードパーティのゲートウェイ経由) 上の API 応答に、思考フィールドが欠落している思考ブロック、またはテキスト フィールドが欠落しているテキスト ブロックが含まれている場合のクラッシュを修正しました。
異常な Unicode シーケンスを含む一部のメッセージでマークダウンのレンダリングが非常に遅くなる問題を修正しました。
セッション名が 200 文字以内または絵文字が多い場合に、ListAgents からコピーされた受信者を拒否する SendMessage を修正しました。
リポジトリ検出が異常なユーザー情報を含む git リモートのホストを誤って読み取り、間違ったホストに対してリンクとリポジトリ固有の動作を生成する問題を修正しました。
解決されたシークレットを出力する MCP 診断を修正しました。スコープ競合の警告には設定済みの ${VAR} フォームが表示され、接続失敗の詳細にはサーバーのオリジンのみが表示されるようになりました。
git が実際に接続するホストとは異なる SCP スタイルの git マーケットプレイス ソースを受け入れる strictKnownMarketplaces ホワイトリストを修正しました
全画面でコピーすると文字が失われる /login OAuth URL などのモーダル テキストを修正しました
レンダリングされたマークダウンの後の行に流れ込む --- 水平ルールを修正しました
Todo/タスクの更新が複数の「Ran 1 シェル コマンド」行にインターリーブされた場合に、連続するシェル コマンドが複数の「Ran 1 シェル コマンド」行に分割される問題を修正しました。
! キーを押している間に開く /permissions などのダイアログを修正しました。シェルコマンドが実行中だったが、コマンドが終了すると終了した
キューを修正しました!上矢印を押してキューに入れられた入力を編集した後、シェル コマンドがプレーン テキストとしてモデルに送信されます
キューに入れられたメッセージがプロンプト履歴に再表示される問題を修正しました。

e はまだキューにあり、キューにあるメッセージを選択中に Esc を押してもターンは中断されなくなり、!モードはターン途中のサブミット後に固定されなくなりました
「新しいフルスクリーン レンダラーを試しますか?」の受け入れを修正しました。許可モード (例: --dangerously-skip-permissions )、ツールの許可/拒否ルール、モデルまたはエフォート フラグを使用せずにセッションを再起動するプロンプト
/tui が再起動時に launch --allowed-tools / --disallowed-tools ルールを削除する問題を修正しました。セッションに制限がある場合、再起動が引き継がれないという理由で、切り替えが拒否されるようになりました。
リポジトリが存在する前にディレクトリが最初に表示されたときに、リポジトリ全体のスコープの警告が省略される信頼プロンプトを修正しました。
権限の再プロンプト中に IDE の差分タブを閉じると、以前の入力で新しいプロンプトに応答できる場合がある問題を修正しました。
修正: Claude Code Desktop または VS Code によってホストされるリモート コントロール セッション中にユーザーに送信されたファイルがアップロードされるようになり、空のカードが表示されるのではなく、電話や Web で開かれるようになりました。
修正: CLAUDE_CODE_OAUTH_TOKEN が設定されているときに /login を実行した後、古いトークンのリマインダーがクロードの自動的に再開されたターンに漏れなくなりました。これはあなたにのみ表示されるようになりました。
修正: 権限プレビューは、受信トラスト ゲートによって許可されたチャネル サーバーにのみ中継されるようになり、サーバーの明示的な権限機能のオプトアウトが尊重されるようになりました。
修正: リレーされた権限プレビューの認証情報マスキングで、コマンド、パス、宛先を承認者から隠すことができなくなりました。特大の秘密キー ブロックが完全な強度のリダクションでリダクションされるようになりました
修正: 権限プレビューでマスクするプロバイダー API トークンは、直後にシェル区切り文字が続く場合でもマスクされるようになりました。
Claude Desktop のセッション間メッセージが、セッション間メッセージングが無効になっていると読み取られたときに受信者セッションによってサイレントにドロップされ、送信者のクエリが「考えられる」ままになる問題を修正しました。

「ng」を何分間も言い続ける
リモート コントロール: このコンピューターを別の claude.ai アカウントまたは組織にサインインすると、実行中のセッションが数秒以内に停止され、誤解を招く HTTP 404 時間後のメッセージではなく、その理由が表示されるようになりました。
Claude Code Desktop または VS Code から開始されたリモート コントロール セッションでは、電話機と claude.ai/code がセッションの許可モード (およびモデルの claude.ai/code) に応じて更新されるようになりました。
リモート コントロール: 電話または claude.ai/code で行われたエフォート ピックがターミナルおよびデスクトップ/VS Code でホストされるセッションに適用されるようになり、セッションはそのエフォート レベルを接続されたクライアントに公開します。
SendMessage と ListAgents は、アカウントのセッション リストが長すぎて完全に確認できない場合に、未表示のセッションを不在として扱うのではなく、通知するようになりました。
claude.ai ログインが優先される場合、期限切れの Anthropic プロファイル認証情報が /login を指すようになりました。
トランスクリプトの改善: 独自のプロンプトで、返信と同じ方法でマークダウン (強調表示されたコード ブロック、インライン コード、リスト) がレンダリングされるようになりました。
「API が空または不正な応答を返しました」エラーを改善し、返された内容 (コンテンツ タイプ、本文の種類、サイズ、リクエスト ID) と元のストリーミング リクエストが失敗した理由を示すようにしました。
自動生成されるセッション タイトルが改善され、リクエストを言い換えた文章 (例: 「モバイルのログイン ボタンを修正する」) ではなく、短く具体的な名前 (例: 「ログイン ボタンのバグ」) として読み取れるようになりました。
オンデマンドでリファレンス ドキュメントをロードすることにより、組み込みの claude-api スキルをロードするコンテキスト コストを最大 200,000 トークン以上から最大 25,000 トークンに削減しました。
クロードの作業中に /permissions を開くことができるようになりました — ルールの変更は現在のターンの残りの部分に適用されます
クロードの作業中に /add-dir <path> を使用できるようになりました。 /add-dir 、 /autocompact 、 /theme 、 /help 、 /config および /advisor ダイアログがフルスクリーン TUI でターンの途中で開きます
/goal は no で自動的にクリアされるようになりました

回復不能なエラー（認証の取り消し、クレジット残高の枯渇、コンテキストのオーバーフローなど）により、武装したままではなくターンが終了した場合
/goal : バックグラウンド タスクによってゴールが 30 分以上待たされる場合、クロードは無期限に待機するのではなく、バックグラウンド タスクにチェックインするようになりました (オプトアウトするには CLAUDE_CODE_GOAL_CHECKIN_MINUTES=0 を設定します)
クロードセットアップトークンは、予期しない追加の引数を黙って無視するのではなく、拒否するようになりました。
フルスクリーン モードで Esc キーを変更して、マウス テキストの選択をクリアしなくなりました。通常どおり中断または消去され、選択内容は強調表示されたままになります。
自動モードがエージェント ツール呼び出しごとに表示する冗長な「自動モード分類子によって許可」行を削除しました。
/config から「デフォルトのチームメイト モデル」設定を削除しました。エージェントチームのチームメイトは、スポーンが名前を付けない限り、リーダーのモデルを使用するようになりました
実行中のツールのヘッダーの経過時間カウンターを淡色表示にして、太字のカウントと競合しないようにしました。
ターン間に配信されるバックグラウンド タスク通知は、ターン途中の配信と一致して、<system-reminder> タグ内のモデルに送信されるようになりました。
Mantle: メインループ モデルが既に選択されている場合、起動時に管理ピンの可用性プローブをスキップします。
Windows: ~/.claude.json が読み取り専用の場合に名前変更の再試行を繰り返しても起動が停止しなくなりました
2.1.233 2026 年 8 月 14 日
GitLab マージ リクエスト URL サポートを --worktree フラグとクロード エージェント ビューに追加しました (MR は !N として表示されます)。
サインインしているユーザーの ID をヘッダーとして送信する、Anthropic アップストリームにオプトインの forward_user_identity アプリ ゲートウェイ設定を追加しました。これにより、ゲートウェイの背後にあるプロキシがユーザーごとの支出を属性にできるようになります。
Linux 上の Bash ツール コマンド ( CLAUDE_CODE_TOOL_MEMORY_LIMIT ) にオプトイン メモリ cgroup サポートを追加したため、暴走ビルドによってセッションが停止することはありません
CLAUDE_CODE_WEBFETCH_CACHE_TTL_MS 環境変数をに追加しました

WebFetch セッション URL キャッシュ TTL を構成します (デフォルトは変更なし: 15 分)
クロードが許可プロンプトを待っている間に環境がシャットダウンすると、クラウド セッションが失われたとマークされることがある問題を修正しました。
固定タイムアウトで長期保持ストリームを終了するサーバー (サーバーレス ホストなど) に対してサブスクリプション/リッスン ストリームを際限なく再オープンする MCP v2 接続を修正しました。
Claude Desktop または VS Code で実行しているときに権限プロンプトに対して通知フックが起動しない問題を修正しました。
サンドボックスが有効になっている場合、Linux 上のアイドル セッションが 1 つの CPU コアを 100% に維持することがある問題を修正しました。
-p モードまたはプラグイン/MCP がロードされている場合、ユーザーまたはプロジェクトのスキルがバンドルされたスキルをシャドウするときに、/checkup や /review などのバンドルされたスキルのエイリアスが「不明なコマンド」を報告する問題を修正しました。
引数の値がテンプレート マーカーとして再展開されないように、スキル/コマンド引数の置換を修正しました。
NT \??\ デバイス プレフィックスで綴られた Windows パスが UNC パス検証をバイパスし、NTLM 資格情報漏洩ベクトルを閉じる問題を修正しました。
クロード セルフホスト ランナー セッションの開始時間の改善: セッション ブランチは作業ツリーを書き換えることなく作成されるようになり、サーバーの 2 つのラウンド トリップによってエージェントの起動がブロックされなくなりました。
アプリゲートウェイエラー転送の改善: AWS アップストリーム上の Vertex、Foundry、および Claude Platform からの 400/413 エラーには、アップストリーム独自のメッセージが含まれるようになりました。アプリゲートウェイの自動圧縮に関するバグを修正しました
裸の .claude/skills ディレクトリをチェックするための claude プラグインの検証が改善され、フロントマターが解析に失敗した SKILL.md ファイルが報告されるようになりました。
改善されたスクリーン リーダー モード: /effort セレクターは、入力された番号プロンプトを含む番号付きリストとして表示され、ヒントとダイアログ テキストが切り取られなくなりました。
印刷モード診断の改善: モデルに対するリクエストが送信されると、[claude-code:unrecognized_model] 行が stderr に書き込まれます。

ID クロード コードは認識しません。それをmodelOverridesでサイレンスにマップします
GitHub アプリのセットアップ ヒントが、元のリモートが gitlab.com または bitbucket.org にあるリポジトリに表示されなくなるように変更されました。エンタープライズ マーケットプレイスのヒントでは、GitHub 以外の内部 Git ホストもカバーされるようになりました
Todo/タスク追跡ツール (TaskCreate/Get/Update/List、TodoWrite) は、Opus 4.8、Sonnet 5、Fable 5、Mythos 5、およびそれ以降のモデルでは利用できなくなりました。 CLAUDE_CODE_ENABLE_TODO_TOOLS=1 を設定して元に戻します
Windows: 通常の cd <dir> && <command> > file Bash コマンドで手動承認のために自動モードが繰り返し停止する問題を修正しました (2.1.232 のリグレッション)
Windows 上の Cygwin スタイルのシンボリックリンクと入力リダイレクト (< file ) に対する 2.1.232 Bash 権限の変更を元に戻しました。より狭いバージョンは後のリリースで戻る予定です
2.1.232 2026 年 8 月 13 日
サブエージェントのフォークがデフォルトでオンになりました: subagent_type: "fork" サブエージェントは完全な会話とプロンプト キャッシュを継承し、対話型セッションで非チームメイト エージェントが生成され、デフォルトでバックグラウンドで実行されるようになりました。
プロンプトに @ を入力して、別のクロード セッションを名前で言及します。その後、クロードは SendMessage を使用してそのセッションに直接アクセスします
SendMessage は、最初に参照による確認を求める代わりに、1 つのライブ セッションに正確に一致する裸の名前を配信するようになりました。
インタラクティブな設定

[切り捨てられた]

## Original Extract

Release notes for Claude Code, including new features, improvements, and bug fixes by version.

Claude Code changelog - Claude Code Docs Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Claude Code Docs home page English Search... ⌘ K Ask Assistant Claude Developer Platform
Search... Navigation Getting started Claude Code changelog Getting started Build with Claude Code Administration Configuration Reference Agent SDK What's New Resources Getting started
Store instructions and memories
Copy page Copy page Release notes for Claude Code, including new features, improvements, and bug fixes by version.
Copy page Copy page This page is generated from the CHANGELOG.md on GitHub .
Run claude --version to check your installed version.
​ 2.1.234 August 17, 2026
Added the optional CLAUDE_CODE_PROJECT_DIR_NAME environment variable: hosts that give each session its own config directory can choose a short name for the per-project transcript directory
Added the selection:clear keybinding action, so a key can be bound to clear an in-app text selection; also works in the agents view
Added a GitLab merge request badge to the footer and statusline: repos with a GitLab remote and an authenticated glab CLI show MR !N with draft/pending/green states
Claude Code now continues your session automatically when a claude.ai usage limit resets; turn it off in /config (“Continue automatically at usage limit”)
Claude is now told to use your account email only to identify you, and not to send it to unrelated services unless you ask
Security: remote file reads, session restore, CLAUDE.md includes, workflow scripts and file uploads now reject Windows NT-namespace ( \??\ ) paths, hardening the remaining pre-approval file accesses against the NTLM credential-leak vector
Fixed auto mode in very long sessions repeatedly re-checking and denying sandboxed commands’ network access after the conversation had been compacted
Fixed session-scoped permission answers (including denies) being dropped when answering background subagent tool permission prompts
Fixed a crash when an API response on the non-streaming fallback path (typically via third-party gateways) contained a thinking block missing its thinking field or a text block missing its text field
Fixed markdown rendering becoming extremely slow for some messages containing unusual Unicode sequences
Fixed SendMessage rejecting a recipient copied from ListAgents when the session name is at the 200-character cap or emoji-heavy
Fixed repository detection mis-reading the host of git remotes with unusual userinfo, producing links and repo-specific behavior for the wrong host
Fixed MCP diagnostics printing resolved secrets: scope-conflict warnings now show the configured ${VAR} form, and connection-failure details show only the server origin
Fixed strictKnownMarketplaces allowlists accepting SCP-style git marketplace sources whose host differs from the one git would actually connect to
Fixed modal text such as the /login OAuth URL losing characters when copied in fullscreen
Fixed a --- horizontal rule in rendered markdown running into the line after it
Fixed consecutive shell commands splitting into multiple “Ran 1 shell command” rows when todo/task updates were interleaved between them
Fixed dialogs like /permissions opened while a ! shell command was running being dismissed when the command finished
Fixed a queued ! shell command being sent to the model as plain text after pressing up-arrow to edit the queued input
Fixed queued messages reappearing in the prompt history while still queued, Esc while selecting a queued message no longer interrupts the turn, and ! mode no longer sticks after a mid-turn submit
Fixed accepting the “Try the new fullscreen renderer?” prompt restarting the session without its permission mode (e.g. --dangerously-skip-permissions ), tool allow/deny rules, model or effort flags
Fixed /tui dropping launch --allowed-tools / --disallowed-tools rules when it restarts; it now declines to switch, with the reason, when the session has restrictions a restart can’t carry over
Fixed trust prompts omitting the repository-wide scope warning when the directory was first seen before the repository existed there
Fixed a case where an IDE diff tab closing during a permission re-prompt could answer the new prompt with the previous input
Fixed: files sent to the user during Remote Control sessions hosted by Claude Code Desktop or VS Code now upload, so they open on phone and web instead of showing an empty card
Fixed: after /login while CLAUDE_CODE_OAUTH_TOKEN is set, the stale-token reminder no longer leaks into Claude’s automatically resumed turn — it now appears only to you
Fixed: permission previews now relay only to channel servers admitted by the inbound trust gate, and a server’s explicit permission-capability opt-out is honored
Fixed: credential masking on relayed permission previews can no longer hide commands, paths, or destinations from the approver; oversized private-key blocks now redact under full-strength redaction
Fixed: provider API tokens that mask on permission previews now mask even when directly followed by shell delimiters
Fixed Claude Desktop inter-session messages being silently dropped by the recipient session when cross-session messaging read as disabled, which left the sender’s query “thinking” for many minutes
Remote Control: signing this computer in to a different claude.ai account or organization now stops the running session within seconds and says why, instead of a misleading HTTP 404 hours later
Remote Control sessions started from Claude Code Desktop or VS Code now keep phones and claude.ai/code updated on the session’s permission mode (and claude.ai/code on the model) as they change
Remote Control: effort picks made on a phone or on claude.ai/code now apply to terminal- and Desktop/VS Code-hosted sessions, and the session publishes its effort level to connected clients
SendMessage and ListAgents now say when your account’s session list was too long to check completely, instead of treating unseen sessions as absent
Expired Anthropic profile credential now points you at /login when a claude.ai login would take precedence
Improved the transcript: your own prompts now render markdown (highlighted code blocks, inline code, lists) the same way replies do
Improved the “API returned an empty or malformed response” error to say what came back (content type, body kind, size, request ID) and why the original streaming request failed
Improved auto-generated session titles to read as short, specific names (e.g. “Login button bug”) rather than sentences restating your request (e.g. “Fix the login button on mobile”)
Reduced the context cost of loading the built-in claude-api skill from ~200k+ tokens to ~25k by loading reference docs on demand
/permissions can now be opened while Claude is working — rule changes apply to the rest of the current turn
/add-dir <path> can now be used while Claude is working; /add-dir , /autocompact , /theme , /help , /config and /advisor dialogs open mid-turn in the fullscreen TUI
/goal now clears itself with a notice when a turn dies on an unrecoverable error (e.g. revoked auth, an exhausted credit balance, or a context overflow) instead of staying armed
/goal : when background tasks keep a goal waiting for 30+ minutes, Claude now checks in on them instead of waiting indefinitely (set CLAUDE_CODE_GOAL_CHECKIN_MINUTES=0 to opt out)
claude setup-token now rejects unexpected extra arguments instead of silently ignoring them
Changed Esc in fullscreen mode to no longer clear a mouse text selection: it interrupts or dismisses as usual and the selection stays highlighted
Removed the redundant “Allowed by auto mode classifier” line that auto mode showed under every Agent tool call
Removed the “Default teammate model” setting from /config ; agent-team teammates now use the leader’s model unless the spawn names one
Dimmed the elapsed-time counter on the running tool header so it no longer competes with the bold counts
Background task notifications delivered between turns are now sent to the model inside <system-reminder> tags, matching mid-turn delivery
Mantle: skip the admin-pin availability probe at startup when a main-loop model is already picked
Windows: startup no longer stalls on repeated rename retries when ~/.claude.json is read-only
2.1.233 August 14, 2026
Added GitLab merge request URL support to the --worktree flag and the claude agents view (where MRs display as !N )
Added an opt-in forward_user_identity apps gateway setting on Anthropic upstreams that sends the signed-in user’s identity as headers, so a proxy behind the gateway can attribute spend per user
Added opt-in memory cgroup support for Bash tool commands on Linux ( CLAUDE_CODE_TOOL_MEMORY_LIMIT ) so a runaway build can’t stall the session
Added CLAUDE_CODE_WEBFETCH_CACHE_TTL_MS environment variable to configure the WebFetch session URL cache TTL (default unchanged: 15 minutes)
Fixed cloud sessions occasionally being marked as lost when the environment shut down while Claude was waiting on a permission prompt
Fixed MCP v2 connections endlessly reopening the subscriptions/listen stream against servers that terminate long-held streams on a fixed timeout (e.g. serverless hosts)
Fixed Notification hooks not firing for permission prompts when running under Claude Desktop or VS Code
Fixed idle sessions on Linux sometimes keeping one CPU core at 100% when sandboxing is enabled
Fixed bundled skill aliases like /checkup and /review reporting “Unknown command” in -p mode or with plugins/MCP loaded when a user or project skill shadows the bundled skill
Fixed skill/command argument substitution to prevent argument values from being re-expanded as template markers
Fixed Windows paths spelled with the NT \??\ device prefix bypassing UNC path validation, closing an NTLM credential-leak vector
Improved claude self-hosted-runner session start time: the session branch is now created without rewriting the working tree, and two server round trips no longer block the agent’s launch
Improved apps gateway error forwarding: 400/413 errors from Vertex, Foundry, and Claude Platform on AWS upstreams now carry the upstream’s own message; fixes a bug with auto-compact on apps gateway
Improved claude plugin validate to check a bare .claude/skills directory, reporting SKILL.md files whose frontmatter fails to parse
Improved screen reader mode: the /effort selector renders as a numbered list with a typed-number prompt, and hint and dialog text is no longer clipped
Improved print mode diagnostics: a [claude-code:unrecognized_model] line is written to stderr when a request goes out for a model ID Claude Code doesn’t recognize; map it with modelOverrides to silence
Changed the GitHub app setup tip to no longer appear in repositories whose origin remote is on gitlab.com or bitbucket.org; the enterprise marketplace tip now covers non-GitHub internal git hosts
Todo/task-tracking tools (TaskCreate/Get/Update/List, TodoWrite) are no longer available on Opus 4.8, Sonnet 5, Fable 5, Mythos 5, and newer models; set CLAUDE_CODE_ENABLE_TODO_TOOLS=1 to bring them back
Windows: fixed auto mode repeatedly stopping for manual approval on ordinary cd <dir> && <command> > file Bash commands (a 2.1.232 regression)
Reverted the 2.1.232 Bash permission changes for Cygwin-style symlinks on Windows and for input redirections ( < file ); a narrower version will return in a later release
2.1.232 August 13, 2026
Subagent forking is now on by default: a subagent_type: "fork" subagent inherits the full conversation and prompt cache, and non-teammate agent spawns in interactive sessions now run in the background by default
Type @ in the prompt to mention another Claude session by name; Claude then uses SendMessage to reach that session directly
SendMessage now delivers to a bare name that exactly matches one live session, instead of asking to confirm with a ref first
Interactive se

[truncated]
