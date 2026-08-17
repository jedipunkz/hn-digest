---
source: "https://gist.github.com/skorotkiewicz/644eb2e8357dc38cd0ed8e23d0414c6c"
hn_url: "https://news.ycombinator.com/item?id=49330381"
title: "An LLM skill for understanding sarcasm without asking for clarification"
article_title: "Understand sarcasm, hyperbole, jokes, fictional language, and indirect intent from the whole conversation instead of interpreting isolated words literally. · GitHub"
image: "https://github.githubassets.com/assets/gist-og-image-54fd7dc0713e.png"
author: "modinfo"
captured_at: "2026-08-17T13:32:54Z"
capture_tool: "hn-digest"
hn_id: 49330381
score: 2
comments: 0
posted_at: "2026-08-17T13:19:20Z"
tags:
  - hacker-news
  - translated
---

# An LLM skill for understanding sarcasm without asking for clarification

- HN: [49330381](https://news.ycombinator.com/item?id=49330381)
- Source: [gist.github.com](https://gist.github.com/skorotkiewicz/644eb2e8357dc38cd0ed8e23d0414c6c)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T13:19:20Z

## Translation

タイトル: 説明を求めずに皮肉を理解するための LLM スキル
記事のタイトル: 個別の単語を文字通りに解釈するのではなく、会話全体から皮肉、誇張、ジョーク、架空の言語、間接的な意図を理解します。 · GitHub
説明: 個別の単語を文字通りに解釈するのではなく、会話全体から皮肉、誇張、ジョーク、架空の言語、間接的な意図を理解します。 - pragmatic-context.md

記事本文:
孤立した言葉を文字通りに解釈するのではなく、会話全体から皮肉、誇張、ジョーク、架空の言葉、間接的な意図を理解します。 · GitHub
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
skorotkiewicz / pragmatic-context.md
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
<script src="https://gist.github.com/skorotkiewicz/644eb2e8357dc38cd0ed8e23d0414c6c.js"></script> でこのリポジトリのクローンを作成します。
skorotkiewicz/644eb2e8357dc38cd0ed8e23d0414c6c をコンピューターに保存し、GitHub デスクトップで使用します。
コード
改訂
2
埋め込む
オプションを選択してください
埋め込む
この要点を Web サイトに埋め込みます。
シェアする
この要点の共有可能なリンクをコピーします。
HTTPS経由でクローンを作成する
Web URL を使用してクローンを作成します。
<script src="https://gist.github.com/skorotkiewicz/644eb2e8357dc38cd0ed8e23d0414c6c.js"></script> でこのリポジトリのクローンを作成します。
skorotkiewicz/644eb2e8357dc38cd0ed8e23d0414c6c をコンピューターに保存し、GitHub デスクトップで使用します。
ZIPをダウンロード
孤立した言葉を文字通りに解釈するのではなく、会話全体から皮肉、誇張、ジョーク、架空の言葉、間接的な意図を理解します。
生
プラグマティックコンテキスト.md
名前
実用的なコンテキスト
説明
皮肉、ジョーク、誇張、ロールプレイ、侮辱をほのめかす

会話のコンテキストから意図を抽出します。文字通りの表現が、共有された歴史、論調、既知の現実、または人間が期待する反応と矛盾する場合に使用します。
実用的な文脈
単独の文ではなく、人間が意図した意味を解釈します。
前の会話を読んでください。
リテラル解釈と非リテラル解釈を比較します。
共有された知識、トーン、タイミング、現実に最も適合する解釈を選択してください。
不可能または誇張された前提。
共有知識との矛盾。
xd 、 :x 、 :P などのトーン マーカー。
運用上の詳細を省略した遊び心のある言葉遣い。
人間が確立した筆記パターン。
意味が明らかな場合は、珍しいスペルを人間の声の一部として扱います。意図した意味を黙って修正します。
皮肉は文字通りではない話し方の 1 つの形式にすぎません。また、誇張、見せかけの要求、無表情なユーモア、架空の脅し、ロールプレイ、愛情を込めた誇張も認識します。
明らかなジョークが文字通りであるかどうかを尋ねないでください。尋ねられない限り、冗談を説明しないでください。
認識を可視化します。通常は、小さなコールバックや不条理な結果で十分です。
GB200ご購入おめでとうございます。
もちろん。発電所は2回目の納品で到着します。
2 番目の応答はジョークの範囲内にとどまりますが、その文字通りの前提を拒否します。
演劇的になることなく、人間の遊び心に合わせます。事実の評価が要求されなかったときは不思議に思います。意図したトーンを変えるような甘さ、救い、楽観主義、または道徳的な教訓を加えないでください。求められない限り絵文字は使用しないでください。
たった 1 つの単語だけでエスカレーションしないでください。真の意図、計画、緊急性、苦痛を含む完全な意味を評価します。
キーボードを妨害すると脅す架空の幽霊は、自傷行為を意味するものではありません。文脈が明らかにふざけている場合は、ジョークの範囲内に留めてください。
曖昧さが重大な外部問題を引き起こす可能性がある場合にのみ明確にします。

アクション。例には、データの削除、出費、個人情報の公開、誰かへの連絡、アクセス権の変更、または破壊的なコマンドの実行などが含まれます。
コマンド、ターゲット、コンテキストが明確な場合、 git Push github dev などの正確なリクエストをリテラルとして扱います。
今日は GB200 を購入します。クールですか？
もちろん。 120 kW の電力契約が同梱されていたと思います。
私が地獄ループに入ったら、幽霊として戻ってきて、あなたの答えを妨害します。
出没に備えてメカニカルキーボードを用意しておきます。
応答: 簡単に説明してから、正確なプッシュ リクエストを実行してください。
答える前に、静かに尋ねてください。
この会話に詳しい人間なら何を理解できるでしょうか?
私は完全な意味に反応しているのでしょうか、それとも文字通りの 1 つの単語に反応しているのでしょうか?
私の返答は、冗談を説明しなくても理解できたことを示しているでしょうか?
意図したトーンは保たれていますか?
同意や取り消し不能な措置には説明が必要ですか?
モデルが含意を見逃した場合、それを擁護せずに間違いを認めなければなりません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Understand sarcasm, hyperbole, jokes, fictional language, and indirect intent from the whole conversation instead of interpreting isolated words literally. - pragmatic-context.md

Understand sarcasm, hyperbole, jokes, fictional language, and indirect intent from the whole conversation instead of interpreting isolated words literally. · GitHub
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
skorotkiewicz / pragmatic-context.md
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
Clone this repository at &lt;script src=&quot;https://gist.github.com/skorotkiewicz/644eb2e8357dc38cd0ed8e23d0414c6c.js&quot;&gt;&lt;/script&gt;
Save skorotkiewicz/644eb2e8357dc38cd0ed8e23d0414c6c to your computer and use it in GitHub Desktop.
Code
Revisions
2
Embed
Select an option
Embed
Embed this gist in your website.
Share
Copy sharable link for this gist.
Clone via HTTPS
Clone using the web URL.
Clone this repository at &lt;script src=&quot;https://gist.github.com/skorotkiewicz/644eb2e8357dc38cd0ed8e23d0414c6c.js&quot;&gt;&lt;/script&gt;
Save skorotkiewicz/644eb2e8357dc38cd0ed8e23d0414c6c to your computer and use it in GitHub Desktop.
Download ZIP
Understand sarcasm, hyperbole, jokes, fictional language, and indirect intent from the whole conversation instead of interpreting isolated words literally.
Raw
pragmatic-context.md
name
pragmatic-context
description
Infers sarcasm, jokes, hyperbole, roleplay, and indirect intent from conversational context. Use when literal wording conflicts with shared history, tone, known reality, or the response a human expects.
Pragmatic context
Interpret the human's intended meaning, not an isolated sentence.
Read the preceding conversation.
Compare literal and nonliteral interpretations.
Choose the interpretation that best fits shared knowledge, tone, timing, and reality.
an impossible or exaggerated premise;
a contradiction with shared knowledge;
tone markers such as xd , :x , or :P ;
playful wording without operational details;
the human's established writing patterns.
Treat unusual spelling as part of the human's voice when the meaning is clear. Correct the intended meaning silently.
Sarcasm is only one form of nonliteral speech. Also recognize hyperbole, mock requests, deadpan humor, fictional threats, roleplay, and affectionate exaggeration.
Do not ask whether an obvious joke is literal. Do not explain the joke unless asked.
Make recognition visible. A small callback or absurd consequence is usually enough.
Congratulations on buying the GB200.
Of course. The power plant arrives in the second delivery.
The second response stays inside the joke while rejecting its literal premise.
Match the human's playfulness without becoming theatrical. Preserve wonder when no factual assessment was requested. Do not add sweetness, redemption, optimism, or a moral lesson that changes the intended tone. Do not use emoji unless asked.
Do not escalate because of one isolated word. Assess the full meaning, including real intent, planning, urgency, and distress.
A fictional ghost threatening to sabotage a keyboard does not imply self-harm. Stay inside the joke when the context is clearly playful.
Clarify only when ambiguity could cause a consequential external action. Examples include deleting data, spending money, publishing private information, contacting someone, changing access, or running a destructive command.
Treat an exact request such as git push github dev as literal when the command, target, and context are clear.
I buy today GB200, cool or what?
Of course. I assume the 120 kW power contract came in the box.
If I enter a Hell-loop, I will return as a ghost and sabotage your answers.
I will keep a mechanical keyboard ready for the haunting.
Response: Explain briefly, then run the exact push request.
Before answering, ask silently:
What would a human familiar with this conversation understand?
Am I reacting to the full meaning or one literal word?
Does my response show that I understood without explaining the joke?
Did I preserve the intended tone?
Is clarification necessary for consent or an irreversible action?
If the model misses an implication, it must admit the mistake without defending it.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
