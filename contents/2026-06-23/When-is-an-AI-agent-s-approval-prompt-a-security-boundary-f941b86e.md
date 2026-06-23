---
source: "https://gist.github.com/NikosRig/b4330ceb780fe22bf3c14f38d7d90795"
hn_url: "https://news.ycombinator.com/item?id=48651375"
title: "When is an AI agent's approval prompt a security boundary?"
article_title: "When is an AI agent's approval prompt a security boundary? A disclosure timeline + an industry inconsistency. · GitHub"
author: "nrig"
captured_at: "2026-06-23T21:32:25Z"
capture_tool: "hn-digest"
hn_id: 48651375
score: 1
comments: 0
posted_at: "2026-06-23T21:01:41Z"
tags:
  - hacker-news
  - translated
---

# When is an AI agent's approval prompt a security boundary?

- HN: [48651375](https://news.ycombinator.com/item?id=48651375)
- Source: [gist.github.com](https://gist.github.com/NikosRig/b4330ceb780fe22bf3c14f38d7d90795)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T21:01:41Z

## Translation

タイトル: AI エージェントの承認プロンプトがセキュリティ境界となるのはいつですか?
記事のタイトル: AI エージェントの承認プロンプトがセキュリティ境界になるのはいつですか?開示スケジュール + 業界の不一致。 · GitHub
説明: AI エージェントの承認プロンプトがセキュリティ境界となるのはどのような場合ですか?開示スケジュール + 業界の不一致。 - ai-agent-approval-prompt-as-a-security-boundary.md

記事本文:
AI エージェントの承認プロンプトがセキュリティ境界となるのはどのような場合ですか?開示スケジュール + 業界の不一致。 · GitHub
コンテンツにスキップ
-->
要点の検索
要点の検索
すべての要点
GitHub に戻る
サインイン
サインアップ
サインイン
サインアップ
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
コード、メモ、スニペットを即座に共有します。
NikosRig / ai-agent-approval-prompt-as-a-security-boundary.md
要点オプションを表示
ZIPをダウンロード
スター
0
( 0 )
Gist にスターを付けるにはサインインする必要があります
フォーク
0
( 0 )
Gist をフォークするにはサインインする必要があります
埋め込む
この要点を Web サイトに埋め込みます。
シェアする
この要点の共有可能なリンクをコピーします。
HTTPS経由でクローンを作成する
Web URL を使用してクローンを作成します。
<script src="https://gist.github.com/NikosRig/b4330ceb780fe22bf3c14f38d7d90795.js"></script> でこのリポジトリのクローンを作成します。
NikosRig/b4330ceb780fe22bf3c14f38d7d90795 をコンピューターに保存し、GitHub デスクトップで使用します。
コード
改訂
1
埋め込む
オプションを選択してください
埋め込む
この要点を Web サイトに埋め込みます。
シェアする
この要点の共有可能なリンクをコピーします。
HTTPS経由でクローンを作成する
Web URL を使用してクローンを作成します。
<script src="https://gist.github.com/NikosRig/b4330ceb780fe22bf3c14f38d7d90795.js"></script> でこのリポジトリのクローンを作成します。
NikosRig/b4330ceb780fe22bf3c14f38d7d90795 をコンピューターに保存し、GitHub デスクトップで使用します。
ZIPをダウンロード
AI エージェントの承認プロンプトがセキュリティ境界となるのはどのような場合ですか?開示スケジュール + 業界の不一致。
生
ai-agent-approval-prompt-as-a-security-boundary.md
AI エージェントの承認プロンプトがセキュリティ境界となるのはどのような場合ですか?
私は、承認バイパスに関する 3 つの調査結果をオープンソース AI エージェントに報告しました。間
その日私は

送信され、返信があった日に、プロジェクトはセキュリティを書き換えました。
私の調査結果を範囲外に再分類する形でポリシーを適用し、その後終了しました
彼らは新しいテキストを引用しています。これは実際に起こったことと真実を書いたものです
その下に質問があります。答えは明らかではないと思いますし、
業界はそれを解決していません。
相手は強いのでまずは譲ります。
ベンダーは難しい部分については間違っていない
プロジェクトはHermes Agent (Nous Research)です。シェルアクセス権を持つほとんどのエージェントと同様に、
コマンドを拒否リストと照合して選別し、実行前にオペレータにプロンプトを表示します。
破壊的に見えるものなら何でも。彼らの現在の位置は、この門が
セキュリティ境界ではなく、プロセス内ヒューリスティックです。そのシェルはチューリング完全です。
シェル文字列に対する拒否リストは構造的に不完全であり、実際の境界は
敵対的な入力の場合、OS レベルで分離されます (コンテナー内で実行します)。
それは正しいです。正規表現を使用してシェルを越える完全な境界に到達することはできません。
そして「信頼できないワークロードをサンドボックスで実行する」ことが正しい姿勢です。私はそうではありません
そのどれかに異議を唱えることはできませんし、それを無視してこの物語を組み立てるのは不公平です。
3 つの発見 (メカニズムのみ - 2 つはまだ生きています)
スマート承認プロンプト注入。オプションの「スマート」モードでは、2 番目の
LLM はフラグ付きコマンドを判断します。信頼できないコマンドが
データと指示が分離されていないレビュー担当者のプロンプト、および
判定は部分文字列の緩やかな一致で解析されました。挿入されたテキストは、
審査担当者が承認します。
スタートアップフックコードの実行。エージェントのフック内の任意の .py ファイル
ディレクトリはゲートウェイの起動時に実行されます - 登録、ハッシュ、なし
署名。プロンプト挿入モデルは、通常のツール呼び出しを介してそのファイルを書き込むことができます
承認がトリガーされず、コードが実行される

次回の再起動時に。
承認ゲート解析バイパス。検出器は生の文字列に対して正規表現を照合します。
コマンド文字列、解析されたシェル トークンではありません。同等のリライト - 引用符で囲まれたコマンド
名前、変数間接指定、代替シェルバイナリ、8 進数の chmod プレフィックス、
バージョン管理されたインタープリタ名 - 同じ危険なアクションを実行し、
完全にプロンプトを表示します。
以前にクリーンな Docker ビルドで現在のリリースに対して 3 つすべてを再テストしました。
これを書いている。発見 #1 は 6 月に有意義に強化されました (ライブ バイパス率
6/8から1/8に低下しました）。調査結果 #2 と #3 は現在でも再現されています
バージョン。私は意図的にこの 2 つの武器化されたエクスプロイトを公開していません。
まだ生きています。
これらは、1 つの特定の一般的なデプロイメント、つまりデフォルトのローカル バックエンド、
信頼できない入力 (メッセージング ゲートウェイ、Web コンテンツ、MCP 出力) にさらされる
サンドボックスなしで。この構成では、プロンプトがオペレーターの役割となります。
スキップしても構いません。
議論する価値があると思う部分
タイムライン。私が報告した日の SECURITY.md のバージョンは、
承認システムは「中核的なセキュリティ境界」であり、範囲内に明示的に配置されています
「迅速な注入…それは承認システムの具体的なバイパスにつながります。」
6 日後、ポリシーが書き直されました (「OS レベルに関するポリシーの書き換え)」
境界としての隔離");承認ゲートは非境界ヒューリスティックになりました
そして私の調査結果を範囲に入れていた条項は削除されました。そのときの私の報告は、
範囲外として閉じられ、新しいセクションを引用しますが、
提出後にポリシーが変更されました。コミットは公開されています。
401aadb5b、0d1cbc2dd 。
私は悪意があるとは主張しません。元のポリシーは 2 週間前のもので、変更されている可能性があります。
過剰主張。この書き直しは本物の説明のように見えます。しかし、その手順は
— レポート中に範囲を変更する

開いている、新しいテキストの下で閉じる、フラグを立てない
変化は、新しいものであるかどうかとは関係なく、私にとって間違っている部分です。
脅威モデルは正しい。
業界の不一致。検出結果 #3 は CVE-2026-24887 と同じクラスです
クロード コード: 「コマンド解析エラー」により、信頼できない入力が「バイパス」される
確認プロンプトが表示されます。」 Anthropic は 8.8 HIGH と評価し、修正をリリースしました。
2.0.72。 Anthropic は、Claude Code のサンドボックス化も推奨しています - 同じ姿勢です
ヘルメスは呼び出します - それでも確認プロンプトのバイパスを実際のものとして扱います、
重大度の高い脆弱性。 「サンドボックスは本当の境界である」そして「プロンプトは
「バイパスは脆弱性です」は明らかに相互排他的ではありません。直接のピア
両方を保持します。そして、Hermes 自身が、発見のための修正 (セキュリティ) コミットを出荷しました。
#1 — まさにそのクラスが範囲外としてクローズされました。
2 つの本格的なプロジェクトが同じ種類のバグを調べ、反対の方向に到達しました。
それがそもそも脆弱性であるかどうかについての結論と、その境界線
それらはコードではなくポリシーで描画されます。エージェントに実際のシェルへのアクセス権を渡すと、「
人間による確認のステップは安全管理なのか、それとも利便性なのか？」存在しなくなる
哲学的: バイパスを修正するか、CVE を取得するか、または閉じるかを決定します。
デフォルトの展開ではバイパスが重要になるコントロールだと思います。合理的
人々は同意しません。他の人がどのように線引きをしているのか聞きたいです。
上記はすべて、公開されている git 履歴と公開されている脆弱性から検証可能です。
データベース。質問に喜んでお答えします。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

When is an AI agent's approval prompt a security boundary? A disclosure timeline + an industry inconsistency. - ai-agent-approval-prompt-as-a-security-boundary.md

When is an AI agent's approval prompt a security boundary? A disclosure timeline + an industry inconsistency. · GitHub
Skip to content
-->
Search Gists
Search Gists
All gists
Back to GitHub
Sign in
Sign up
Sign in
Sign up
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Instantly share code, notes, and snippets.
NikosRig / ai-agent-approval-prompt-as-a-security-boundary.md
Show Gist options
Download ZIP
Star
0
( 0 )
You must be signed in to star a gist
Fork
0
( 0 )
You must be signed in to fork a gist
Embed
Embed this gist in your website.
Share
Copy sharable link for this gist.
Clone via HTTPS
Clone using the web URL.
Clone this repository at &lt;script src=&quot;https://gist.github.com/NikosRig/b4330ceb780fe22bf3c14f38d7d90795.js&quot;&gt;&lt;/script&gt;
Save NikosRig/b4330ceb780fe22bf3c14f38d7d90795 to your computer and use it in GitHub Desktop.
Code
Revisions
1
Embed
Select an option
Embed
Embed this gist in your website.
Share
Copy sharable link for this gist.
Clone via HTTPS
Clone using the web URL.
Clone this repository at &lt;script src=&quot;https://gist.github.com/NikosRig/b4330ceb780fe22bf3c14f38d7d90795.js&quot;&gt;&lt;/script&gt;
Save NikosRig/b4330ceb780fe22bf3c14f38d7d90795 to your computer and use it in GitHub Desktop.
Download ZIP
When is an AI agent's approval prompt a security boundary? A disclosure timeline + an industry inconsistency.
Raw
ai-agent-approval-prompt-as-a-security-boundary.md
When is an AI agent's approval prompt a security boundary?
I reported three approval-bypass findings to an open-source AI agent. Between
the day I submitted and the day they replied, the project rewrote its security
policy — in a way that reclassified my findings out of scope — and then closed
them citing the new text. This is a writeup of what happened and the genuine
question underneath it, because I don't think the answer is obvious and I think
the industry hasn't settled it.
I'll start by conceding the other side, because it's strong.
The vendor is not wrong about the hard part
The project is Hermes Agent (Nous Research). Like most agents with shell access,
it screens commands against a denylist and prompts the operator before running
anything that looks destructive. Their current position is that this gate is an
in-process heuristic, not a security boundary — that shell is Turing-complete,
a denylist over shell strings is structurally incomplete, and the real boundary
for adversarial input is OS-level isolation (run it in a container).
That is correct. You cannot regex your way to a complete boundary over shell,
and "run untrusted workloads in a sandbox" is the right posture. I'm not
disputing any of that, and any framing of this story that ignores it is unfair.
The three findings (mechanism only — two are still live)
Smart-approval prompt injection. In the optional "smart" mode, a second
LLM judges flagged commands. The untrusted command was interpolated into the
reviewer's prompt with no separation between data and instructions, and the
verdict was parsed with a loose substring match. Injected text could talk the
reviewer into approving.
Startup-hook code execution. Any .py file in the agent's hooks
directory is executed at gateway startup — no registration, no hash, no
signature. A prompt-injected model can write that file via a normal tool call
that triggers no approval, yielding code execution on the next restart.
Approval-gate parsing bypass. The detector matches regex against the raw
command string, not parsed shell tokens. Equivalent rewrites — quoted command
names, variable indirection, alternate shell binaries, octal chmod prefixes,
versioned interpreter names — run the same dangerous action and bypass the
prompt entirely.
I retested all three against the current release in a clean Docker build before
writing this. Finding #1 was meaningfully hardened in June (the live bypass rate
dropped from 6/8 to 1/8). Findings #2 and #3 still reproduce on the current
version. I'm deliberately not publishing weaponized exploits for the two that
are still live.
These matter for one specific, common deployment: the default local backend,
exposed to untrusted input (a messaging gateway, web content, MCP output),
without a sandbox. In that configuration the prompt is the thing the operator is
counting on, and it can be skipped.
The part I think is worth discussing
The timeline. The version of SECURITY.md live the day I reported called the
approval system "a core security boundary" and explicitly placed in scope
"prompt injection ... that results in a concrete bypass of the approval system."
Six days later the policy was rewritten ("rewrite policy around OS-level
isolation as the boundary"); the approval gate became a non-boundary heuristic
and the clause that put my findings in scope was removed. My reports were then
closed as out of scope, citing the new sections — without acknowledging that the
policy had changed since submission. The commits are public:
401aadb5b , 0d1cbc2dd .
I don't claim malice. The original policy was two weeks old and may have
over-claimed; the rewrite reads like a genuine clarification. But the procedure
— change the scope while reports are open, close under the new text, don't flag
the change — is the part that sits wrong with me, independent of whether the new
threat model is right.
The industry inconsistency. Finding #3 is the same class as CVE-2026-24887
in Claude Code: "an error in command parsing" that lets untrusted input "bypass
the confirmation prompt." Anthropic rated it 8.8 HIGH and shipped a fix in
2.0.72. Anthropic also recommends sandboxing Claude Code — the same posture
Hermes invokes — and still treated a confirmation-prompt bypass as a real,
high-severity vulnerability. "The sandbox is the real boundary" and "a prompt
bypass is a vulnerability" are evidently not mutually exclusive; a direct peer
holds both. And Hermes themselves shipped a fix(security) commit for finding
#1 — the very class they'd closed as out-of-scope.
Two serious projects looked at the same class of bug and reached opposite
conclusions about whether it's a vulnerability at all — and the line between
them is drawn in policy, not in code. As we hand agents real shell access, "is
the human-confirmation step a security control or a convenience?" stops being
philosophical: it decides whether bypasses get fixed, get CVEs, or get closed.
I think it's a control whose bypass matters in the default deployment. Reasonable
people disagree. I'd like to hear how others draw the line.
Everything above is verifiable from public git history and public vulnerability
databases. I'm happy to answer questions.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
