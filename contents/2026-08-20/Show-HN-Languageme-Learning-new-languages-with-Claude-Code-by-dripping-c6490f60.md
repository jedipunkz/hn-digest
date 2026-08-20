---
source: "https://github.com/magiclex/languageme"
hn_url: "https://news.ycombinator.com/item?id=49373850"
title: "Show HN: Languageme – Learning new languages with Claude Code by dripping"
article_title: "GitHub - MagicLex/languageme: Progressive language-immersion drip for Claude Code: preprompt a % of each reply in a target language, monitor the real ratio, ramp it up as you keep up. · GitHub"
image: "https://opengraph.githubassets.com/5ab1c73bc3ecf50bf4ecbd9654ec7fcfda9f3b62ee684545adc80a5c0520ddc6/MagicLex/languageme"
author: "LexSiga"
captured_at: "2026-08-20T13:39:26Z"
capture_tool: "hn-digest"
hn_id: 49373850
score: 1
comments: 0
posted_at: "2026-08-20T12:42:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Languageme – Learning new languages with Claude Code by dripping

- HN: [49373850](https://news.ycombinator.com/item?id=49373850)
- Source: [github.com](https://github.com/magiclex/languageme)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T12:42:34Z

## Translation

タイトル: Show HN: Languageme – クロード コードを滴下して新しい言語を学習する
記事のタイトル: GitHub - MagicLex/ languageme: クロード コードのためのプログレッシブ言語イマージョン ドリップ: ターゲット言語での各応答の % を事前にプロンプトし、実際の比率を監視し、追いつくにつれて増加させます。 · GitHub
説明: クロード コードのプログレッシブ言語イマージョン ドリップ: ターゲット言語での各応答の % を事前にプロンプトし、実際の比率を監視し、追いつくにつれて増加させます。 - MagicLex/言語ミー

記事本文:
GitHub - MagicLex/ languageme: クロード コードのプログレッシブ言語イマージョン ドリップ: ターゲット言語での各応答の % を事前にプロンプトし、実際の比率を監視し、追いつくにつれて増加させます。 · GitHub
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
マジックレックス
/
言語ミー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
19 コミット 19 コミット フォルダーとファイル
アセット アセット .gitignore .gitignore ライセンス ライセンス README.md README.md 言語ミー

languageme すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード・コードのための進歩的な言語イマージョンの点滴。事前にプロンプトを表示し、強制します
各返信のどれだけがターゲット言語で返されるかを監視し、増加します
カレンダーではなく、あなたが把握している証拠を共有します。
スタンドアロン: 1 つの Python ファイル、stdlib のみ、pip のインストールはありません。ブレンドドリップではなく、
ハードな言語切り替え。
4 つのレイヤー、3 つのクロード コード フックとステータスライン セグメントとして配線されています。
プリプロンプト (SessionStart フック)。 languageme フックが点滴を注入します
新しいセッションごと: 散文の ~N% をターゲットにブレンドし、残りはそのままにします
いつもの返信言語。 N を段落ごとの具体的な量子として記述します
(モデルは生のパーセンテージを自己調整できません)、それぞれを判断するように指示します。
自分の以前の返信に対してではなく、ターゲットに対して返信します。 Nさんに住んでいます
状態。ターゲットは散文チャットに限定されており、コード、コミット、ファイル名、
ツールの引数、または間違った単語が何かを壊す句。
ナッジ (UserPromptSubmit フック)。 languageme ナッジでループを閉じる
ワンショットのプリプロンプトはできません。各応答の前に、
前のターンの測定ブレンド: 最後の返信 80% 対 ~16% ターゲット、大幅に上回っており、段落ごとに 1 つの短いフレーズ。これにより、ターゲットが毎ターン再アサートされます（免疫
コンテキストの成長に合わせて)、オーバーシュートと完全なドロップアウトの両方を修正します (2 つの方法)
オープンループの点滴が漂います。
モニター（フック止め）。 languageme モニターは完成したトランスクリプトを読み取り、
あなたが生成した実際のターゲット言語のシェアを測定します (文字重み付け、
文内ブレンドを読み取るトークンごとの分類子、ML 依存関係なし)、
あなたが助けを求めたか、それとも自分で言語を書いたかを追跡し、N を移動します。
ランプポリシーに従って。ランプはセッションの累積ブレンドを読み取り、ゲートをオンにします。
バンド : ターゲットが多すぎると、マスではなく不遵守としてカウントされます

テリー、だからオーバーシュート
N を増加させることはありません。
ステータスラインセグメントが付いているので、ダイヤルがどこにあるかを常に確認できます。
読み取り値 12%↑ ~11 :
12% : 現在のブレンド ターゲット (プリプロンプトが目指すもの)
↑ : 最後のターンでゲージが上がったばかり ( ↓ 下降、何もない = 安定)
~11 : 最後のターンで測定されたターゲット言語の実際の割合 ( ~ =
おおよそ）。 No ~N はまだ回転が測定されていないことを意味します。
インストールは非破壊的に接続します。すでにステータスラインがある場合はラップされます。
それを実行してセグメントを追加します。追加しない場合はセグメントが追加されます。編集したり置き換えたりすることはありません
あなたのステータスライン。
実際に測定されたシェアが目標を維持し、あなたが
助けを求めているわけではありません。
./ languageme init sv # スウェーデン語を 10% シードします
./ languageme install # ワイヤーフック + ステータスライン (~/.claude/settings.json を編集)
./ languageme ステータス # ターゲット、ブレンド %、最近の測定値、ランプ状態
プレプロンプトは SessionStart でロードされるため、次のクロードでドリップが開始されます。
コードセッション。終了せずに現在のものでアクティブ化するには、/clear を実行します。
(SessionStart が再起動されます)。最初のターンが測定されるまでステータスライン
~N のないターゲットを示し、ゲージは置かれたままです。 ~N および任意のランプ
点滴の下でターンが実行されると表示されます。
必要に応じて、これを PATH に追加します: ln -s "$PWD/ languageme" ~/bin/ languageme 。
フックは絶対パスで呼び出すため、シンボリックリンクはオプションです。
いつでもターゲットを切り替えます: languageme lang sv|es|lv 。ゲージとランプが続きます
新しい言語。
ランプ ポリシー ( languageme ランプ <mode> )
マスタリー (デフォルト)。 3 つの異なるセッションの後に 1 ステップずつランプします。
累積ブレンドがターゲット バンド内に収まります (少なすぎず、多すぎず)。
助けを求められることはありませんでした。各セッションは、実行したターン数に関係なく 1 回としてカウントされます。執筆
言語自体にボーナスステップが追加されます。オーバーシュートまたはヘルプのリセットを要求する
筋。
カレン

ダル。目標までは監視セッションごとに 1 ステップ。ハンズオフ、ブラインド
あなたが追いついているかどうかに。
マニュアル。決して自動で移動することはありません。あなたはそれをバンプします: languageme バンプ +5 。
自然言語オーバーライドは常に両方向で優先されます。
言語を意識している。 「スウェーデン語が多すぎる」と言う（またはターゲット自身の言葉、「för mycket」
「svenska」、「demasiado español」、「trop de français」）チャットと次のモニターで
pass ステップでブレンドを下げます。 「もっとスウェーデン語」/「メル・スベンスカ」/「スペイン語」と言ってください
そしてそれはステップアップします。元に戻すか（「英語」、「フランス語」）、または
一言 (「X はどういう意味ですか」) を説明すると、助けが必要であるとみなされ、ランプが保持されます。
通常のコーディングの話 (「構成を変換する」、「クエリが遅くなった」など) はそうではありません。
ヘルプモード ( languageme help <mode> )
自動 (デフォルト)。初めて使用するときに新しい単語を単語 (翻訳) として光沢表示します。
再発するにつれて消えていきます。
の上。すべてのターゲットフレーズが光沢がありました。曖昧さがゼロになり、学習が遅くなります。
languageme の音声を追加して、短い発音のスペルを重ねます。
トリッキーな単語 (å/ä/ö 母音、sj/tj/rs クラスター、サイレント文字)、例:
sköldpadda (カメ、「HYULD-pad-da」) 。 IPA ではなく、単純な学習者のリスペリングのみ
スペルが音を隠しているところ。デフォルトではオフです。ヘルプがオフの場合は無視されます。
2 つのレベルがあるため、いつでも電源をオフにできます。
永続的な切り替え。 languageme off はドリップを一時停止します (プリプロンプトなし、なし
測定する。ステータスラインには sv off が表示されます)。 languageme がそれを再開します。の
状態はセッション間で持続します。
ハードキル (環境)。 export LANGUAGEME_OFF=1 は完全に無効にします
シェルまたはセッション、何に関係なく状態をオーバーライドします。フックは何も注入しません。
モニターは何もせず、ステータスラインは純粋なパススルーに戻ります。
(元の行、セグメントなし)。保証されたオフスイッチ。
init <lang> ターゲット言語 (sv、lv、de、es、it、...) をシードします。
ステータスフル状態の読み出し
ステータス

ne [--wrap] コンパクトセグメント。 --wrap はラップされたステータスラインも実行します
フック [SessionStart] プリプロンプトを出力します
モニター [--session P] [停止] ターンを測定し、% をランプします
[--full] [--quint] トランスクリプト全体を再測定 / レポートを沈黙させます
PCT|ゴール|フロア|ステップ<n>を設定し、数値を調整します
バンプ [+n] でブレンドを微調整します
ゴール <n> に向かって傾斜する上限
ヘルプ オン|オフ|自動インライングロス ポリシー
音声オン|オフ トリッキーな単語の発音のヒント (å/ä/ö...)
ランプマスタリー|カレンダー|マニュアル
lang [<code>] ターゲットの表示 / 切り替え
オフ |点滴を一時停止/再開します（継続）
インストール |ワイヤーのアンインストール / フックのアンワイヤー + ステータスライン
自己更新の更新 (git pull ツールのチェックアウト)
測定対象：sv、es、lv。他の人は点滴を受けていますが、まだ％はありません。
単一のシェル/セッションのハードキル (すべてをオーバーライド): LANGUAGEME_OFF=1
州
$LANGUAGEME_HOME (デフォルトは ~/. languageme/state.json ) にあります。コードは
ポータブル;状態だけが変化するのです。バックアップをインストールしてください
最初の編集前に設定を settings.json. languageme.bak に変更します。
測定可能なターゲット: sv、es、lv (さらに識別子として fr/en)。各船
ストップワード辞書とその辞書内の特徴的な文字 /
LANG_CHARS レジストリ。ターゲットは他のすべてのターゲットに対して得点されます。を追加する
language はこれら 2 つのエントリであり、他には何もありません。点滴とプリプロンプトの作業
すぐに使えるどの言語でも、測定の半分だけが辞書に縛られます (
未測定のターゲットは % を 0 に保持するだけです)。
install はフック エントリを追加し、ステータスラインをラップします (決して置き換えません)。
既存のフックとステータスラインはすべてそのまま残ります。アンインストールは削除します
独自のステータスラインのみを復元し、元のステータスラインを復元します。
スケールは平らです。状態は制限されています (履歴カーソルとセッション カーソルは制限されています)。
Stop は、バイト シークを介して新しいトランスクリプトの末尾のみを読み取ります (ファイル全体ではありません)。
そしてリードm

odify-write は flock でシリアル化されているため、クロード コードが並列化されています
セッションはお互いの測定値を妨害することはできません。
マサチューセッツ工科大学「ライセンス」を参照してください。使って、フォークして、出荷してください。
Claude Code のプログレッシブ言語イマージョン ドリップ: ターゲット言語での各応答の % を事前にプロンプ​​トし、実際の比率を監視し、追いつくにつれて増加させます。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Progressive language-immersion drip for Claude Code: preprompt a % of each reply in a target language, monitor the real ratio, ramp it up as you keep up. - MagicLex/languageme

GitHub - MagicLex/languageme: Progressive language-immersion drip for Claude Code: preprompt a % of each reply in a target language, monitor the real ratio, ramp it up as you keep up. · GitHub
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
MagicLex
/
languageme
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
19 Commits 19 Commits Folders and files
assets assets .gitignore .gitignore LICENSE LICENSE README.md README.md languageme languageme View all files Repository files navigation
A progressive language-immersion drip for Claude Code. It preprompts, forces
and monitors how much of each reply comes back in a target language, and ramps
that share up on evidence you're keeping up, not on a calendar.
Standalone: one Python file, stdlib only, no pip installs. A blend drip, not
a hard language switch.
Four layers, wired as three Claude Code hooks plus a statusline segment:
Preprompt (SessionStart hook). languageme hook injects the drip into
every new session: blend ~N% of your prose into the target, keep the rest in
your usual reply language. It states N as a concrete per-paragraph quantum
(a model cannot self-calibrate a raw percentage) and tells you to judge each
reply against the target, not against your own earlier replies. N lives in
state. The target is confined to chat prose, never code, commits, filenames,
tool arguments, or any clause where a wrong word breaks something.
Nudge (UserPromptSubmit hook). languageme nudge closes the loop the
one-shot preprompt cannot. Before each reply it injects one live line from the
previous turn's measured blend: last reply 80% vs ~16% target, well over, one short phrase per paragraph . This re-asserts the target every turn (immune
to context growth) and corrects both overshoot and full dropout, the two ways
an open-loop drip drifts.
Monitor (Stop hook). languageme monitor reads the finished transcript,
measures the real target-language share you produced (char-weighted,
per-token classifier that reads the intra-sentence blend, no ML dependency),
tracks whether you asked for help or wrote the language yourself, and moves N
per the ramp policy. The ramp reads the session-cumulative blend and gates on
a band : too much target counts as noncompliance, not mastery, so overshoot
never ramps N up.
A statusline segment rides along so you always see where the dial is.
Reading sv 12%↑ ~11 :
12% : current blend target (what the preprompt aims for)
↑ : the gauge just moved up on the last turn ( ↓ down, nothing = steady)
~11 : the real share of target language measured on the last turn ( ~ =
roughly). No ~N yet means no turn has been measured.
install wires it non-destructively: if you already have a statusline it wraps
it and appends the segment, if you don't it adds one. It never edits or replaces
your statusline.
The blend only ever grows when the real measured share holds the target and you
are not asking for help.
./languageme init sv # seed Swedish at 10%
./languageme install # wire hooks + statusline (edits ~/.claude/settings.json)
./languageme status # target, blend %, recent measurements, ramp state
The preprompt loads at SessionStart, so the drip starts on your next Claude
Code session. To activate it in the current one without leaving, run /clear
(it re-fires SessionStart). Until the first turn is measured the statusline
shows the target with no ~N , and the gauge stays put; the ~N and any ramp
appear once a turn has run under the drip.
Put it on your PATH if you like: ln -s "$PWD/languageme" ~/bin/languageme .
The hooks call it by absolute path, so a symlink is optional.
Switch target any time: languageme lang sv|es|lv . The gauge and ramp follow
the new language.
Ramp policies ( languageme ramp <mode> )
mastery (default). Ramps one step after 3 distinct sessions whose
cumulative blend lands in the target band (not too little, not too much) with
zero help asked. Each session counts once, however many turns it ran. Writing
the language yourself adds a bonus step. Overshoot or asking for help resets
the streak.
calendar. One step per monitored session until the goal. Hands-off, blind
to whether you're keeping up.
manual. Never auto-moves. You bump it: languageme bump +5 .
Natural-language overrides always win, both directions, and they are
language-aware. Say "too much swedish" (or the target's own words, "för mycket
svenska", "demasiado español", "trop de français") in chat and the next monitor
pass steps the blend down ; say "more swedish" / "mer svenska" / "más español"
and it steps up. Asking to switch back ("in english", "en français") or to
gloss a word ("what does X mean") counts as needing help and holds the ramp.
Ordinary coding talk ("translate the config", "the query got slower") does not.
Help modes ( languageme help <mode> )
auto (default). Glosses new words as word (translation) on first use,
fades as they recur.
on. Every target phrase glossed. zero ambiguity, slower learning.
Add languageme phonetic on to layer a short pronunciation respelling on the
tricky words (the å/ä/ö vowels, sj/tj/rs clusters, silent letters), e.g.
sköldpadda (turtle, "HYULD-pad-da") . Plain learner respelling, not IPA, only
where the spelling hides the sound. Off by default; ignored when help is off.
Two levels, so you can always shut it off:
Persistent toggle. languageme off pauses the drip (no preprompt, no
measuring; the statusline shows sv off ). languageme on resumes it. The
state persists across sessions.
Hard kill (env). export LANGUAGEME_OFF=1 fully disables it for that
shell or session, overriding state no matter what. The hook injects nothing,
the monitor does nothing, and the statusline falls back to pure passthrough
(your original line, no segment). The guaranteed off switch.
init <lang> seed a target language (sv, lv, de, es, it, ...)
status full state readout
statusline [--wrap] compact segment; --wrap runs the wrapped statusline too
hook [SessionStart] print the preprompt
monitor [--session P] [Stop] measure a turn, ramp the %
[--full] [--quiet] re-measure whole transcript / silence the report
set pct|goal|floor|step <n> tune the numbers
bump [+n] nudge the blend now
goal <n> ceiling to ramp toward
help on|off|auto inline-gloss policy
phonetic on|off pronunciation cues on tricky words (å/ä/ö...)
ramp mastery|calendar|manual
lang [<code>] show / switch target
off | on pause / resume the drip (persists)
install | uninstall wire / unwire hooks + statusline
update self-update (git pull the tool's checkout)
Measurable targets: sv, es, lv. Others get the drip but no % yet.
Hard kill for a single shell/session (overrides everything): LANGUAGEME_OFF=1
State
Lives in $LANGUAGEME_HOME (default ~/.languageme/state.json ). The code is
portable; the state is the only thing that mutates. install backs up your
settings to settings.json.languageme.bak before its first edit.
Measurable targets: sv, es, lv (plus fr/en as discriminators). Each ships
a stopword lexicon and its distinctive characters in the LEXICONS /
LANG_CHARS registries; the target is scored against all the others. Adding a
language is those two entries, nothing else. The drip and preprompt work for
any language out of the box, only the measurement half is lexicon-bound (an
unmeasured target just holds the % at 0).
install appends its hook entries and wraps (never replaces) your statusline;
it leaves every existing hook and statusline untouched. uninstall removes
only its own and restores the original statusline.
Scales flat. State is bounded (history and session cursors are capped), each
Stop reads only the new transcript tail via a byte seek (not the whole file),
and the read-modify-write is flock -serialized so parallel Claude Code
sessions can't clobber each other's measurements.
MIT. See LICENSE . Use it, fork it, ship it.
Progressive language-immersion drip for Claude Code: preprompt a % of each reply in a target language, monitor the real ratio, ramp it up as you keep up.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
