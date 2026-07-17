---
source: "https://www.olafalders.com/2026/07/17/claude-code-anatomy-of-a-misfeature/"
hn_url: "https://news.ycombinator.com/item?id=48947776"
title: "Claude Code: Anatomy of a Misfeature"
article_title: "Claude Code: Anatomy of a Misfeature · Olaf Alders"
author: "oalders"
captured_at: "2026-07-17T15:06:41Z"
capture_tool: "hn-digest"
hn_id: 48947776
score: 10
comments: 2
posted_at: "2026-07-17T14:26:58Z"
tags:
  - hacker-news
  - translated
---

# Claude Code: Anatomy of a Misfeature

- HN: [48947776](https://news.ycombinator.com/item?id=48947776)
- Source: [www.olafalders.com](https://www.olafalders.com/2026/07/17/claude-code-anatomy-of-a-misfeature/)
- Score: 10
- Comments: 2
- Posted: 2026-07-17T14:26:58Z

## Translation

タイトル: クロード コード: 誤機能の解剖学
記事のタイトル: Claude Code: Anatomy of a Misfeature · Olaf Alders
説明: 2026 年のカナダデー (7 月 1 日) に、Anthropic は驚くべき製品を出荷しました。

記事本文:
コンテンツへスキップ ~ / オラファルダーズ My Mind is Racing
～についてのトピックを投稿します
検索
/ ~/投稿
/
claude-code-anatomy-of-a-misfeature #AI #security #cli #automation クロード コード: 誤機能の解剖学
どのリリースがそれを出荷したか、そしてメモには何が書かれているか
ドキュメントが追いつきました。その後。
なぜ追加されたのでしょうか? (確認ではなく証拠)
差分もできますか？それでは実際に差分をしてみましょう
これが焼かれたトークン以上のものである理由
自動更新をオフにする 落とし穴: これにより、プラグインの更新も停止します
Hustvett の「メカニカル エッグ タイマー」は CC BY-SA 3.0 に基づいてライセンスされています。幅広のフレームにパッドが入っています。この適応も同様に CC BY-SA 3.0 でライセンスされています。
2026 年のカナダデー (7 月 1 日) に、アンスロピックは驚くべき「イースターエッグ」を送りました。
クロード コードのユーザー: 2.1.198 には効率バイパスが含まれており、
エージェントは人間からの指示に妨げられることなく作業を続行できます。あなた
基本的に、Claude Code が入力を求めた後、60 秒のタイマーを取得します。見逃したら
ウィンドウでは、Claude Code が最善と思われることを実行し、続行します。
その方法。次のようになります。
●クロードはこう尋ねた。
⎿ …
● 60 秒を超えると応答なし — 応答なしで継続
● ユーザーが立ち去った。最善の判断で進めさせていただきます。私の計画:
注: 上記は、私自身のクロード セッションの 1 つからそのまま抜粋したもので、質問は削除されています。
この行動に驚くと感じたとしても、あなたは一人ではありません。考えられる結果を考えてみましょう。
サンドイッチを作るとき、ラップトップをキッチンに持っていく必要がありますか?この期間中に afk をするとどうなりますか?
一度にいくつのエージェントを実行していますか?それらすべてを同時に観察できるでしょうか?同じ 60 秒間に 2 人以上のエージェントが入力を求めてきた場合はどうなりますか?
エージェントが間違った選択をした場合はどうなりますか?トークンは何個ありますか

その間に火傷したの？
導入にエージェントを使用している場合はどうなりますか? (はい、わかっていますが、もし だったらどうなるでしょう)
これらは、この機能をリリースするときに考慮すべき合理的な事項であり、その理由を変更ログに文書化することもできます。しかし、変更ログで新しいデフォルトについてまったく言及しなかったらどうなるでしょうか?それはさらに驚くべきことではないでしょうか？ （ネタバレ：そうだった！）
物語には（ある種の）ハッピーエンドがあります。速く動いて物事を壊すことは、必ずしも速く動いて物事を解決することを妨げるものではありません。数日以内に修正版が発送されましたが、この製品に対するユーザーの信頼はどうなるのでしょうか?
Claude Code の驚くべき機能は、理論的には (そして実際には) Anthropic によって毎日のペースで出荷されます。
すべての機能が変更ログに表示されるとは限りません
デフォルトであってはいけないものにはオフスイッチが文書化されていない可能性があります
Claude Code の自動更新機能は、私たちが当初予想していたよりも YOLO モードのように感じられます。
私たちが学んだかどうかわからないことがいくつかあります。
人間はこの方程式にどのように当てはめられるのでしょうか?
この機能は人間が思いついたのでしょうか?
この機能は人間が書いた（またはエージェントに書かせた）のでしょうか？
人間がその機能をレビューしましたか?
人間がその機能を承認しましたか?
人間がその機能をマージしたのでしょうか?
人間は機能の文書化や変更ログへの追加をオプトアウトしましたか?
人間のリリース マネージャーがリリースと以前のリリースを比較し、リリース前に承認の印を押しましたか?
個人的には、人間がこれらすべてをゲートしていたとは信じがたいです
「これは良いアイデアですか?」と尋ねることなく、手順を実行します。クロード・コードって言ってくれたら
実際に機能を構築し、出荷し、承認してからそれをみなした
文書化する価値がない、私はそれを信じたいのですが、
ただ分からないだけです。もしかしたら何かの組み合わせかもしれない

その2つのことのうち。たぶん
多くのことがうまくいかなかったが、これが決してあってはならないことは明らかだと思う
起こった。少なくとも 1 回は業績評価を受けた者としてこれを言います
そこで私のマネージャーは、「そうですね、本番環境に重大なバグを入れてしまったのですね」と言いました。
これがどのようにして起こったのか、そして死後検査とはどのようなものなのかについて少し疑問に思いました
公的な記録で入手可能です。そこで、クロード・コードに調査を依頼しました
それ自体。クロードの名誉のために言っておきますが、これを防ぐフィルターはないようです。
このコードに関しては反省しています。したがって、完全に開示すると、以下は次のとおりです
ほとんどがクロードの作品なので、それを価値あるものとして受け止めてください。
主要な前提条件を個別に再現する価値があります。
クロードの研究はここから始まります。
2026-06-29 — 2.1.196 がリリースされました。記者の「最後に動作するバージョン（推測）」
2026-06-30 — 2.1.197 がリリースされました。変更ログの 1 行、Sonnet 5 のリリース
2026-07-01 — 2.1.198 リリース — 記者が回帰を指摘したバージョン。変更を示すパブリック コミットはありません。このリリースの唯一の公開された痕跡は、そのノート ( 75709ea ) を公開するボット コミットです。これは CHANGELOG.md と feed.xml 以外には何も触れていません。
2026-07-02 02:54 UTC — 問題 #73125 は Aleksey Nogin によって提出されました
2026-07-02 03:45 UTC — コメント投稿者が脱出ハッチを浮上させます: CLAUDE_AFK_TIMEOUT_MS 。スレッド内でピアツーピアで取引されており、どのリリース ノートからも指摘されていません
2026-07-02 — 2.1.199 は、問題が公開されている間に 24 のエントリを出荷します。まだ言及はありません。
2026-07-03 — 2.1.200 では動作が逆転します。繰り返しますが、唯一の公開トレースはメモのコミット ( 1322e9b ) です。
2026-07-04 18:04 UTC — 問題は終了しました
この問題に対する反応/規模: 384 👍 、コメント 143 件 — ニッチな苦情ではありません
記者の環境: 2.1.198 、「最後に動作していた 2.1.196 (推測)」、Opus、AWS Bedrock、VS Cod

e端子
git clone https://github.com/anthropics/claude-code.git
cd クロードコード
# 修正はいつ、どのバージョンに適用されましたか?
git log -1 --format = '%h %ai' -S 'デフォルトでは自動続行されなくなりました' -- CHANGELOG.md
#1322e9ba 2026-07-03 16:52:26 +0000
# バグが出荷されたバージョンはいつ発生しましたか?
git log -1 --format = '%h %ai' -S '## 2.1.198' -- CHANGELOG.md
#75709eac 2026-07-01 20:45:29 +0000
# 不具合
AskUserQuestion は、クロード コードがタスクの途中で人間を止めて質問するために使用するツールです。
新しい動作: 60 秒間の非アクティブ状態が続くと、ツールはブロックする代わりに「とにかく続行」の結果を自動的に返します。
モデルに返されるメッセージ — これはテンプレートであり、ここではデフォルトの 60 秒でレンダリングされます。
// v2.1.198、そのまま。 `Thl` はミニファイアの名前です。散文はバイナリ独自のものです。
// ファイル内の「60s」はリテラルではなく補間されることに注意してください:
function Thl ( e ){ return `${ Math の後に応答がありません。 round ( e / 1000 ) } s — ユーザーは
キーボードから離れてください。これまでの状況に基づいて最善の判断を行ってください。
まだ関連性がある場合は、後でこの質問を再質問できます。` }
「後で再質問する」というエスケープ ハッチは循環しています。再質問された質問は同じタイムアウトに達します。 Aleksey Nogin 氏は、提出から数分以内に問題スレッドでこの点を指摘しました。
トランスクリプト行には 2 つのバリアントがあります。バイナリは、回答を開始したかどうかをそれらの間で選択します。
// v2.1.198、そのまま。 `a` は「いくつかの答えが存在する」ことを表すミニファイアの名前です
// (`s=Object.entries(r)` を答えの上に、`a=s.length>0`);両方の文字列値
// バイナリ自体のものです。
d = a としますか? 「これまでに選択された回答で続行」 : 「回答なしで続行」
したがって、半分答えられたダイアログは部分的な入力を破棄せず、それを送信します。 3 つの質問のうち 1 つに答えて、その場から離れてください。

imeout は、あなたの 1 つの回答に加えて、モデルが他の 2 つの回答として選択したものをコミットします。
どちらの文字列も 2.1.197 には存在しませんが、2.1.198 には存在します。
公平に言えば、画面上では沈黙はありませんでした。ダイアログにはライブカウントダウンが表示され、キーを押すとタイマーが再スタートしました。 1 つの文字列として保存されるのではなく、実行時に組み立てられるため、これはレンダリングされた形式であり、リテラルの grep ヒットではありません。
// v2.1.198、逐語的に - 部分。 「s」は残り秒の値です。
Children : [ "auto- continue in " , s , " s \xB7 任意のキーを入力してください" ]
// 次のようにレンダリングされます: 12 秒で自動続行 · 任意のキーを維持
見た目よりも遠くまで切れません。カウントダウンは画面を見ている人にのみ届きます。この機能の前提は、あなたには届かないことです。内部名は AFK です。 「ユーザーがキーボードから離れている可能性があります」というメッセージが表示される
複数のエージェントを同時に実行している場合、「画面の監視」は 1 か所ではなく、必要なカウントダウンは別のタブで行われていました。
そしてカウントダウンが遅い。しきい値のデフォルトは 20 秒 ( CLAUDE_AFK_COUNTDOWN_MS ) で、経過時間ではなく残り時間を基準とします。そのため、最初の 40 秒間は、ダイアログは通常のブロック質問のように見えます。画面上には表示されていますが、タイマーが実行中であることは何も表示されません。
// v2.1.198 はこれを縮小して出荷します — ローカルは破壊されていますが、プロパティは
// 名前は存続するため、`showCountdown` と `remainingSeconds` は独自の単語になります。
u = i * 1000 <= n ; return { 残り秒 : i 、 showCountdown : u 、 timeoutMs : t }
// ...n がデフォルトの 20000 の場合、次のようになります。
showCountdown = 残り秒 * 1000 <= 20000
警告は最後の 3 分の 1 で届きます
影響しない点: タイムアウトは AskUserQuestion にのみ適用されます。 Anthropic のツールのリファレンスには、「プランの承認を含む許可プロンプトは、アイドル状態では決して自動解決されない」と記載されており、ドキュメントの他の主張とは異なります。

この投稿では、これは修正後に書かれたページではなく、出荷された 2.1.198 コードと照合して確認できます。カウントダウン コンポーネント ( q0m ) にはバンドル全体で呼び出しサイトが 1 つだけあり、そのタイマー フック ( _Rc ) には呼び出し元が 1 つだけあり、それは q0m 自体です。タイマーは 1 つのコンポーネント内に存在します
そのコンポーネントのプロパティは、引数を超えてそれを識別します: jsx(dRc,{question:V,questions:s,currentQuestionIndex:$,answers:R,questionStates:O,onAnswer:be,onSubmit:N, …}) 。そのタイムアウト ハンドラーは tengu_ask_user_question_afk_auto_advance を起動するものです
許可プロンプト (続行しますか、許可しますか…) はタイマーが接続されていない別のコンポーネントです
ただし、免除によって保護されるのは、許可を求めるプロンプトが表示された場合のみです。 2.1.198 には、 bypassPermissions 、 acceptEdits 、 allowedTools 、 --dangerously-skip-permissions 、および PreToolUse フックが同梱されています。デプロイメントに対してエージェントを実行している人は、おそらくデプロイ コマンドを許可リストに登録するか、プロンプトをオフにしています。これが自動化の意味です。彼らにとって、許可レイヤーは決して起動しないため、タイマーに対する耐性は何も意味しません。
より狭い主張で、より噛みつく主張は、タイマーは許可を与えることはできませんが、選択はできるということです。 AskUserQuestion は許可を求めません。「ステージングか本番か?」「どの構成か?」という決定をユーザーに求めます。タイムアウトになると、モデルは最善の判断を使用して続行するように指示され、部分回答パスは「これまでに選択された回答を使用して」続行されます。ホワイトリストまたはバイパスによって許可がすでに付与されている場合、その選択肢が残された唯一のゲートでした
ツール スキーマにはタイムアウト パラメーターが存在しません。モデルはそれを設定したり制御したりすることはできません。バイナリ ( U_f=…H.strictObject({questions:…}) ) で検証されており、問題スレッドから取得されたものではありません。入力パラメータは次のとおりです。
質問、

回答、注釈、メタデータ
したがって、何もスキップしなかったというモデルは真実です。モデルではなくハーネスが答えを返しました。
2.1.200 では自動継続がデフォルトでオフになりました
アイドル タイムアウトは /config 経由でオプトインされるようになり、全員に強制されるわけではありません
変更ログの動詞「デフォルトでは自動継続は行われなくなりました」に注目してください。何も削除されず、デフォルトが反転されました。
メモで示されている /config 設定は、修正前には存在しませんでした。 askUserQuestionTimeout を 2.1.198 で Grep すると、ヒット数はゼロです。これが出荷されたとき、それをエスケープする唯一の方法は環境変数であり、リリース ノートではその名前が付けられていませんでした。
残りは現在のバイナリ (2.1.211、11 リリース後) で確認できます。機構は完全に無傷です。設定は askUserQuestionTimeout で、 /config に「質問の自動続行タイムアウト」として表示されます。
受け入れられる値は 60s 、 5m 、 10m 、 Never です。 unset は Never に解決されます。これによりオプトインされます。
フックのデフォルトのタイムアウト定数は依然として 60000 です。カウントダウンのしきい値は 20000 です
2 つの環境変数 (CLAUDE_AFK_TIMEOUT_MS および CLAUDE_AFK_COUNTDOWN_MS) が引き続き設定をオーバーライドします。
したがって、修正はゲートの 1 行の変更であり、削除ではありません。
// v2.1.211、そのまま — 空白のみが追加されました。 `Upf` はミニファイアの名前です。
// ここにある他のすべてはバイナリ独自のものであり、文字列値も含まれます。
関数 Upf ( e ){ スイッチ ( e ){
ケース "60s" : 60000 を返します。ケース "5m" : 300000 を返します。

[切り捨てられた]

## Original Extract

On Canada Day (July 1), 2026, Anthropic shipped a surprising

skip to content ~ / olafalders My Mind is Racing
posts topics about
search
/ ~/posts
/
claude-code-anatomy-of-a-misfeature #AI #security #cli #automation Claude Code: Anatomy of a Misfeature
Which release shipped it — and what the notes said
The docs caught up. Afterwards.
Why was it added? (evidence, not confirmation)
Can you even diff it? So let’s actually diff it
Why this is more than burned tokens
Turning auto-update off The catch: this also stops updating your plugins
" Mechanical egg timer " by Hustvedt is licensed under CC BY-SA 3.0 . Padded to a wider frame; this adaptation is likewise licensed CC BY-SA 3.0.
On Canada Day (July 1), 2026, Anthropic shipped a surprising “easter egg” to
users of Claude Code: 2.1.198 includes an efficiency bypass which allows
agents to continue on without being blocked on direction from a human. You
essentially get a 60 second timer after Claude Code asks for input. If you miss
the window, Claude Code helpfully does what it thinks is best and continues on
its way. It looks like this:
● Claude asked:
⎿ …
● No response after 60s — continued without an answer
● The user stepped away. I'll proceed with best judgment. My plan:
Note: the above is taken verbatim from one of my own claude sessions, with the questions having been trimmed.
If you find this behaviour surprising, you’re not alone. Let’s consider the possible consequences:
Do you have to take your laptop to the kitchen with you when you’re making a sandwich? What happens if you are afk during this window?
How many agents are you running at once? Can you possibly observe them all at the same time? What if two or more agents ask for your input during the same 60 second window?
What if the agent makes the wrong choice? How many tokens have been burned in the meantime?
What if you are using agents for deployments? (Yes, I know, but what if )
These are reasonable things you might consider when shipping this feature and maybe you’d document your reasoning in the changelog. But what if you never mentioned the new defaults in the changelog at all? Wouldn’t that be even more surprising? (Spoiler: it was!)
The story has a (sort of) happy ending. Move fast and break things does not necessarily preclude move fast and fix things. Within a couple of days a fix was shipped, but where does that leave user trust in this product?
Surprising features in Claude Code can in theory (and in practice) be shipped by Anthropic on a daily cadence
Not every feature will necessarily appear in the changelog
Things that should not be defaults may not have a documented off switch
Claude Code’s auto update feature feels more like YOLO mode than we might have suspected early on
There are a few things I don’t know if we have learned:
How do humans fit into this equation?
Did a human dream up the feature?
Did a human write (or have an agent write) this feature?
Did a human review the feature?
Did a human sign off on the feature?
Did a human merge the feature?
Did a human opt out of documenting the feature or adding it to the changelog?
Did a human release manager diff the release with the previous release and give it their seal of approval before it went out the door?
Personally I would find it hard to believe that a human was gating all of these
steps without asking “is this a good idea?”. If you told me that Claude Code
actually built the feature, shipped it, signed off on it and then deemed it
unworthy of documentation, that’s something I’m more inclined to believe, but I
just don’t know. Maybe it’s some combination of those two things. Maybe a
number of things went wrong, but I think it’s clear that this never should have
happened. And I say this as someone who has had at least one performance review
where my manager said “well, you did put a serious bug into production”.
I’ve wondered a bit about how this happened and what kind of post-mortem is
available in the public record. So, I’ve asked Claude Code to investigate
itself. To Claude’s credit, it seems to have no filter that prevents
self-reflection regarding this code. So, full disclosure, what follows is
mostly Claude’s work, so take that for what it’s worth and if you rely on any
key assumptions, it’s worth reproducing them in isolation.
Claude’s research begins here.
2026-06-29 — 2.1.196 released; the reporter’s “last working version (I am guessing)”
2026-06-30 — 2.1.197 released; one changelog line, the Sonnet 5 launch
2026-07-01 — 2.1.198 released — the version the reporter pins the regression to. No public commit shows the change; the only public trace of this release is the bot commit publishing its notes ( 75709ea ), which touches nothing but CHANGELOG.md and feed.xml
2026-07-02 02:54 UTC — issue #73125 filed by Aleksey Nogin
2026-07-02 03:45 UTC — a commenter surfaces the escape hatch : CLAUDE_AFK_TIMEOUT_MS . Traded peer-to-peer in the thread, not pointed to from any release note
2026-07-02 — 2.1.199 ships 24 entries while the issue is open. Still no mention.
2026-07-03 — 2.1.200 reverses the behaviour; again the only public trace is the notes commit ( 1322e9b )
2026-07-04 18:04 UTC — issue closed
Reaction/scale on the issue: 384 👍 , 143 comments — not a niche complaint
Reporter’s environment: 2.1.198 , “last working 2.1.196 (guessing)”, Opus, AWS Bedrock, VS Code terminal
git clone https://github.com/anthropics/claude-code.git
cd claude-code
# When did the fix land, and in which version?
git log -1 --format = '%h %ai' -S 'no longer auto-continue by default' -- CHANGELOG.md
# 1322e9ba 2026-07-03 16:52:26 +0000
# When did the version that shipped the bug land?
git log -1 --format = '%h %ai' -S '## 2.1.198' -- CHANGELOG.md
# 75709eac 2026-07-01 20:45:29 +0000
# The misfeature
AskUserQuestion is the tool Claude Code uses to stop and ask the human a question mid-task
New behaviour: after 60 seconds of inactivity , the tool auto-returns a “proceed anyway” result instead of blocking
The message handed back to the model — this is the template, rendered here at the 60-second default:
// v2.1.198, verbatim. `Thl` is the minifier's name; the prose is the binary's own.
// Note "60s" is interpolated, not a literal in the file:
function Thl ( e ){ return `No response after ${ Math . round ( e / 1000 ) } s — the user may be
away from keyboard. Proceed using your best judgment based on the context so far;
you can re-ask this question later if it's still relevant.` }
The “re-ask later” escape hatch is circular: the re-asked question hits the same timeout. Aleksey Nogin made this point in the issue thread within minutes of filing it
The transcript line has two variants — the binary picks between them on whether you’d started answering:
// v2.1.198, verbatim. `a` is the minifier's name for "some answers exist"
// (`s=Object.entries(r)` over the answers, `a=s.length>0`); both string values
// are the binary's own.
let d = a ? "continued with the answers selected so far" : "continued without an answer"
So a half-answered dialog does not discard the partial input — it submits it . Answer question one of three, step away, and the timeout commits your one answer plus whatever the model picks for the other two
Both strings are absent from 2.1.197 and present in 2.1.198
In fairness: it was not silent on screen. The dialog rendered a live countdown, and a keypress restarted the timer. Assembled at runtime rather than stored as one string, so this is the rendered form, not a literal grep hit:
// v2.1.198, verbatim — the pieces. `s` is the remaining-seconds value.
children : [ "auto-continue in " , s , "s \xB7 any key to stay" ]
// renders as: auto-continue in 12s · any key to stay
Which cuts less far than it looks. The countdown only reaches someone watching the screen, and the premise of the feature is that you aren’t: the internal name is AFK ; the message says “the user may be away from keyboard”
running several agents at once, “watching the screen” is not one place — the countdown you needed was on another tab
And the countdown is late. The threshold defaults to 20 seconds ( CLAUDE_AFK_COUNTDOWN_MS ), and it gates on remaining time, not elapsed — so for the first 40 seconds the dialog looks like an ordinary blocking question. It is on screen, but nothing on it says a timer is running:
// v2.1.198 ships this minified — the locals are mangled, but the property
// names survive, so `showCountdown` and `remainingSeconds` are its own words:
let u = i * 1000 <= n ; return { remainingSeconds : i , showCountdown : u , timeoutMs : t }
// ...which reads, with n at its default of 20000:
showCountdown = remainingSeconds * 1000 <= 20000
The warning arrives in the last third
What it does not touch: the timeout applies only to AskUserQuestion . Anthropic’s tools reference says “permission prompts, including plan approval, never auto-resolve on idle” — and unlike the docs claims elsewhere in this post, this one is checkable against the shipped 2.1.198 code rather than a page written after the fix: the countdown component ( q0m ) has exactly one call site in the entire bundle, and its timer hook ( _Rc ) has exactly one caller — q0m itself. The timer exists in one component
that component’s props identify it past argument: jsx(dRc,{question:V, questions:s, currentQuestionIndex:$, answers:R, questionStates:O, onAnswer:be, onSubmit:N, …}) . Its timeout handler is what fires tengu_ask_user_question_afk_auto_advance
the permission prompt ( Do you want to proceed , Do you want to allow … ) is a separate component with no timer attached
But the exemption only protects you if a permission prompt appears at all. 2.1.198 ships bypassPermissions , acceptEdits , allowedTools , --dangerously-skip-permissions and PreToolUse hooks. Anyone running agents against deployments has plausibly allowlisted the deploy command or turned prompts off — that is what automating it means . For them the permission layer was never going to fire, so its immunity to the timer buys nothing
The narrower claim, and the one that bites: the timer could not grant permission, but it could make the choice. AskUserQuestion doesn’t ask permission, it asks you to decide — “staging or production?”, “which config?” On timeout the model is told to proceed using your best judgment , and the partial-answers path continues “with the answers selected so far”. If permission was already granted by allowlist or bypass, the choice was the only gate left
No timeout parameter exists in the tool schema — the model can neither set nor control it. Verified in the binary ( U_f=…H.strictObject({questions:…}) ), not taken from the issue thread. The input params are only:
questions, answers, annotations, metadata
So the model is truthful when it says it didn’t skip anything — the harness returned the answer, not the model
2.1.200 made auto-continue off by default
Idle timeout is now opt-in via /config , not forced on everyone
Note the changelog’s verb: “no longer auto-continue by default ” . Nothing was removed — the default flipped.
The /config setting the note points you to did not exist before the fix . Grep 2.1.198 for askUserQuestionTimeout and you get zero hits: when this shipped, the only way to escape it was an env var — and the release notes never named it.
You can confirm the rest in the current binary ( 2.1.211 , eleven releases later). The machinery is entirely intact: the setting is askUserQuestionTimeout , surfaced in /config as “Question auto-continue timeout”
accepted values are 60s , 5m , 10m , never — and unset resolves to never , which is what makes it opt-in
the hook’s default timeout constant is still 60000 ; the countdown threshold is 20000
two env vars still override the setting — CLAUDE_AFK_TIMEOUT_MS and CLAUDE_AFK_COUNTDOWN_MS
So the fix is a one-line change in a gate, not a deletion:
// v2.1.211, verbatim — only whitespace added. `Upf` is the minifier's name for it;
// everything else here is the binary's own, string values included.
function Upf ( e ){ switch ( e ){
case "60s" : return 60000 ; case "5m" : return 300000 ;

[truncated]
