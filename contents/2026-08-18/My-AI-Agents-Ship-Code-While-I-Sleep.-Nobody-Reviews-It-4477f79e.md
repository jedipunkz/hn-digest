---
source: "https://goatsquadstudios.com/blog/how-i-work-with-ai-agents-autonomously"
hn_url: "https://news.ycombinator.com/item?id=49354107"
title: "My AI Agents Ship Code While I Sleep. Nobody Reviews It"
article_title: "My AI Agents Ship Code While I Sleep. Nobody Reviews It. — Goat Squad Studios"
image: "https://goatsquadstudios.com/blog_images/how-i-work-with-ai-agents-autonomously/header.png"
author: "csgod"
captured_at: "2026-08-18T23:13:12Z"
capture_tool: "hn-digest"
hn_id: 49354107
score: 2
comments: 0
posted_at: "2026-08-18T23:06:19Z"
tags:
  - hacker-news
  - translated
---

# My AI Agents Ship Code While I Sleep. Nobody Reviews It

- HN: [49354107](https://news.ycombinator.com/item?id=49354107)
- Source: [goatsquadstudios.com](https://goatsquadstudios.com/blog/how-i-work-with-ai-agents-autonomously)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T23:06:19Z

## Translation

タイトル: 私の AI エージェントは私が寝ている間にコードを送信します。誰もレビューしません
記事のタイトル: 私の AI エージェントは私が寝ている間にコードを送信します。誰もレビューしません。 — ゴート スクワッド スタジオ
説明: 私の AI エージェントはコードを一晩でレビューせずに出荷します: 自律性ゲート、偽造できない検証、およびグラフ エンジニアリングの構成要素はプレーンなマークダウンで実行されます。

記事本文:
AI エージェントは私が寝ている間にコードを送信します。誰もレビューしません。 — Goat Squad Studios Goat Squad Studios サービス 仕事のブログ エクスペリエンス JP ES プロジェクトの開始 → サービス →
AI エージェントは私が寝ている間にコードを送信します。誰もレビューしません。
私の AI エージェントは、自律性ゲート、偽造できない検証、およびグラフ エンジニアリングの構成要素をプレーンなマークダウンで実行し、レビューを行わずに一晩でコードを出荷します。
私は日中に計画を立て、エージェントは夜間に構築し、朝起きてからレビューします。これは、私なしで出荷するものを決定するシステムです。
私の朝はボードと昨晩の差分から始まります。私が寝ている間に、AI エージェントがチケット SET-175 をキューから取り出しました。 4 つの cron ジョブが同じコピー＆ペーストされたデータベース スキャンを実行していました。エージェントはそれを 1 つの共有ヘルパーに取り込み、6 つのテストを作成し、2,282 のテスト スイート全体を実行して開発にデプロイし、チケットに完了のマークを付けました。同じ重複を記述した 2 番目のチケットにも気づき、それを閉じました。
発売前に誰もレビューしませんでした。私は翌朝のレビュー、つまり差分、アクティビティログ、ボードです。
同じエージェントが別のチケットを未完了のままにしておきました。 SET-184 は、プロパティ データ API を使用して保留中の取引の評価ギャップ リスクにフラグを立てます。コードは構築、テスト、デプロイされており、API が失敗すると安全に自動的に停止します。それはまだ稼働していません。なぜなら、これをオンにするということはデータ購読料を支払うことを意味し、それにお金を払う価値があるかどうかをテストで判断できないからです。その電話は私を待っていました。
自律性と呼ばれるチケットフィールドが、私なしでどれが完了できるかを決定しました。この記事の残りの部分は、そのフィールドとその周りのシステムについて説明します。また、自治に関する議論のほとんどは間違ったことについてのものだと思います。
一日中計画を立て、一晩かけて構築し、朝にレビューする
昼間は計画中です。アイデアを仕様に変換し、承認基準を作成し、エージェントに電話をかけます。

必要になる前に食事をしましょう。
夜はループです。バックグラウンド エージェントは、最上位の構築可能なチケットを選択し、それをエンドツーエンドで構築し、終了するか、質問を付けて保留します。
コードは毎日同じように動きます。エージェントはブランチし、dev にマージし、dev は dev Amplify サイトとバックエンドにデプロイします。私は 1 日に 1 回、開発から本番への PR をレビューします。毎日の PR が、宣伝につながる唯一の方法です。
エージェントはすべてを開発者に伝えます。 Prod は 1 日に 1 つの PR を実行し、私がレビュー担当者です。
エージェントはそれぞれのビルドを独自のブランチでブランチします → 開発ループにマージします ここに到達し、それ以上進むことはありません → 開発デプロイ 開発サイト + バックエンドを増幅します → 毎日の PR 開発 → 運用。自分でレビューします → チェック後に製品を発送します
今朝のチケットはこんな感じです。
---
id : SET-175
title : 共有 scanActiveDeals(status) の抽出 cron ヘルパー (API)
レーン：完了
サイズ：S
自律性 : 自動
コミット: [ 0117e92 ]
---
内容: 4 EventBridge crons (chaseScan、weeklyDigest、deadlineReminder、
モーニングブリーフィング通知) 同一のアクティブ取引スキャンをハンドロールします。
ループ ステータス -> GSI1 の QueryCommand のページ分割 -> MAX_DEALS_PER_RUN の上限。
scanActiveDeals(statuses, opts) を package/api/utils に抽出します。各クロン
それを呼び出し、その正確なステータスセットを保存します（chaseScan には「prospect」が含まれます）。
WeeklyDigest は除きます）。
ACCEPTANCE (動作保持、検証可能): 新しい scanActiveDeals 単体テスト
(ページ間のページネーション、MAX キャップの適用、ステータス パラメータにより GSI1 が制御されます)
キー);既存のハンドラー テストは緑色のままです。 APIユニット+TSC。
IDEA-40 から昇格 (用務員の技術的負債、注意台帳ごとに自動昇格)
AL-3 — 製品決定なし、build-next によって検証可能)。
すべての呼び出しは、エージェントが仕様に触れる前に仕様内で行われます。これは、なぜ auto に適しているのかさえ述べています。製品の決定はなく、すべての要件はテストによってチェックできます。
夜の計画

今では朝に復習するのが一般的です。 1 多くの人がこのスケジュールを実行し、朝はだらだらと起きています。違いは、エージェントが完了できる内容です。
私はコーディングエージェントとしてクロードコードをメインにしています。この投稿のすべてのループはセッションの 1 つです。ループ自体はスキルです。リポジトリ内の命令ファイルはコマンドのように実行されます。問題が発生するたびに微調整します。
私のスキル: リポジトリ内の命令ファイルをコマンドのように呼び出します。私の 2 つのリポジトリからの実名です。
すべての AI コミットをコードレビューしない理由
現在議論されているのは、エージェントをどこまで信頼するかということだ。極端な例の 1 つは YOLO モードです。権限はオフ、ガードレールはオフ、エージェントは無料で実行されます。 2 エージェントは安いので、ピッチには脚があります。迅速に発送し、壊れたものは元に戻します。しかし、リバートしても不良データがアン出荷されたり、クライアントの信頼を取り戻したりするわけではありません。もう 1 つの極端な方法では、すべての差分をレビューします。これは、エージェントが読み取れる以上のものを生成するまで機能します。
ほとんどの人は慎重な端に座っています。エージェントの完全な自律性に満足している人はわずか 8% です。 3 ほとんどの仕事で AI を使用しているエンジニアは、AI を完全に引き継ぐことができるものはほとんどないと述べています。 4 このツールは、安全と思われるアクションを自動承認し、残りのアクションにフラグを付ける分類子を使用して違いを分割します。
どちらの陣営も間違った質問をしていると思います。正しい質問は、テストでどのような成果が証明できるのかということです。
私の計画はフォルダー内にあります。チケットごとに 1 つのマークダウン ファイルです。エージェントはファイルを直接編集します。小さなローカルサーバーがそれらをカンバンボードとしてレンダリングします。
すべてのチケットには、すべてを決定するフィールドが 1 つあります。ニーズ入力 |ブロックされました。欠落しているフィールドは need-input としてカウントされるため、デフォルトでは何も構築されません。 auto は、エージェントが人間を介さずにチケットを最後まで処理できることを意味します。 「needs-input」は、テストがどれほどグリーンであっても、チケットが私を待っていることを意味します。
質問は簡単です: テストで証明できるでしょうか?

それとも人間の判断が必要ですか? SET-175 は、動作が変わるとテストが失敗する純粋なリファクタリングでした。証明可能なので、自動的に出荷されました。 SET-184 のコードも同様に証明可能であり、ビルドされました。お金の呼び出しはなかったので、待ちました。
両方のチケットが構築され、検証されました。自律分野は、どれを自分自身で出荷できるかを決定しました。
共有 scanActiveDeals(status) cron ヘルパーの抽出
純粋なリファクタリング。行動変化ゼロ。スキャン動作が変化すると失敗する 6 つの新しいテスト。
すべての保留中の取引について AVM + 報酬を使用して評価ギャップ リスクにフラグを立てる
構築、テスト、展開。プロパティデータ API が失敗した場合は、自動的に安全にシャットオフされます。
それが全体的な考え方です。製品の方向性、価格設定、ブランドの声、法的文章、セキュリティ ポリシー、顧客データ、取り消しできないものなど、テストでは解決できない決定を下します。
明確にするために、 unreviewed は dev に対してレビューされていないことを意味します。エージェントは、開発データに基づいて開発環境にデプロイされます。 prod ブランチは彼らの手の届かないところにあります。自分でカットを実行した場合にのみ出荷されます。
顧客データは、価格設定やコピーと同じ門の後ろにあります。不正な夜間ビルドの影響範囲は、壊れた開発 URL です。
完全なシステムをマップとして示します。
エージェントと 1 人の人間がノードとして。エッジとしてのレーン、プロモーション、およびエスカレーション。紫色のエッジは証明可能であり、単独で実行されます。琥珀色の縁は人を待ちます。
raw →洗練されたプロモート ゲート: 製品コールなし + 証明可能なゲート: 自律性 = 自動欠落フィールド = ニーズ入力検証ラダー すべての段 緑 エリックのみ製品コール エスカレート 2 つの質問ゲートの回答 ルールがクラスをリタイアすると、プロダクト マネージャー エージェントがボード/アイデア/IDEA-n (生) を書き込む アイデアリファイナー エージェントのプレッシャー テスト · サイズ · 強制終了可能、自動キュー ボード/チケットを受け入れられない · レーン: 次のビルド エージェント TDD → E2E → 開発アンカー テストのデプロイ· 実際のブラウザと開発 URL · 完了したログ (開発) 出荷された (本番) 一人の人間であるエリック

ノード注目台帳 AL-n · 回答済みクラス
LangGraph を使用したことがある場合は、これに見覚えがあるでしょう。この手法はグラフ エンジニアリングと呼ばれます。ワイヤー エージェントがこのようなグラフにステップインします。 5 今年は、すべての本格的なフレームワークが、型付き状態、条件付きルーティング、チェックポイント、割り込みゲートなど、同じ構成要素に基づいて開発されました。 6 私のボードには、マークダウンとルールとして 4 つすべてがあります。 Frontmatter は型付けされた状態です。自律性はルーティングです。 need-input は割り込みです。アクティビティ ログがチェックポイントです。
私が放棄するのは執行です。グラフ ランタイムは不正な移動をブロックし、どのステップが失敗したかを正確に通知します。私のルールは、モデルがそれに従っており、朝にドリフトを捕まえるという理由だけで成立します。得られるもの: コミットによってルールを変更でき、私とモデルの間に何も存在しません。本当の違いは、エージェントかあなたか、誰がその道を選ぶかです。 5 グラフがルートを前面に固定します。エージェントに独自のルートを選択させ、代わりにチェックに労力を注ぎました。 Anthropic 独自のエージェント ガイドにも同じことが書かれています。フレームワークを介したシンプルで構成可能なパターンです。 7 グラフ チームでさえこの方向に進んでいます。2026 年の生産トレンドでは、危険なアクションの前に人間による承認ゲートが追加されています。 6
エージェントが偽造できないことの検証
未レビューの出荷は、モデルがそれ自体で完了したと宣言できない場合にのみ機能します。すべての自動チケットは、モデルの外にあるチェックをクリアする必要があります。
すべての自動チケットは 6 段すべてを登ります。星付きの 2 つの部分は、「AI が完成させた」という主張が通常は破綻する部分です。
失敗した vitest は、すべての受け入れ基準 (ハッピー、ヌル/空、無効、および認可境界) について最初にテストします。
Playwright の仕様は、模擬認証ロールの切り替えを通じて実際の UI を駆動します。仕様のないユーザー向けの変更は行われません。
反応ルーター typegen + tsc、タッチされたサーフェスとそれを消費するすべてのものをクリーンアップします。
押す

開発ブランチ。ローカル編集では、これが行われるまでライブ チェックで確認できる変更は何もありません。
デプロイジョブは成功し、開発 URL がレンダリングされ、そのライブ URL に対して Playwright が緑色になり、実行されたパスでエラー モニターがクリーンになりました。四人全員、毎回。
docs/ と CLAUDE.md は、動作やコントラクトが変更されるたびに更新されます。書類が完成するまで発送は完了しません。
エージェントのアクティビティメモがどれほど自信を持って聞こえたとしても、ラングが欠落している場合は、チケットが完了していないことを意味します。
以下は、エージェントが SET-308 を終了するところです。チケットのアクティビティ ログから直接抜粋したものです (切り取られています)。
2026-07-03 [claude] 完了 (自動操縦、開発者検証済み)。 ... 義務付けられたものを追加しました
class-retireingguardpackages/api/infra-guards.test.ts ... 検証
ラダークリア: インフラガード (5、ネガティブコントロールを含む) + フル API スイート
2718パス。 API+インフラタイプチェック緑; SettleStack-Dev (cdk 135s) をデプロイしました。
doc-intelligence E2E PASS とデプロイされた開発者との比較。 AWS ログのライブ確認 —
dev でドキュメントを処理し、selt-doc-process-dev が ZERO でクリーンに終了しました
ウィンドウで「ses:SendEmail AccessDenied」と表示されます。ドキュメントが同期されました。 2552b89 + 4b8e352 をコミットします。
ほとんどのセットアップでは、エージェントが localhost 以降を認識できないため、ライブ検証ステップがスキップされます。私の場合は缶です。デプロイが失敗するとビルド ログを読み取り、デプロイされたと判断する前にデプロイ ステータスをチェックします。 CloudWatch が接触した表面上の新しいエラーを追跡します。 Playwright はローカル サーバーではなく、デプロイされた開発 URL に対して実行されるため、Playwright がテストしたビルドはユーザーがヒットしたビルドになります。
午前 2 時にデプロイが 500 秒を超えると、エージェントはビルド ログを取得し、見つかった内容を修正します。チケットのアクティビティ行には、ギャップも含め、チェックされた内容が正確に記載されています。
エージェントが読み取りまたは運転できるすべての表面。検証行は、モデルが偽装できない部分です。
ハーネス、コード、計画
ドキュメント アトラスと ADR、同じコミット AWS で同期
Amplify Hosting の展開。ビルドログを読み取ります

失敗した場合、CloudWatch のログとアラームが表面に表示される Lambda 30 ハンドラー、ログ ストリームによってデバッグされる Step Functions ドキュメント パイプライン、エンドツーエンドでトレースされる DynamoDB の単一テーブルの状態がデバッグ中に読み取られる S3 ドキュメント バケットと抽出されたテキスト 製品自体が呼び出すモデルの基盤 検証
デプロイされた開発 URL に対して Playwright が駆動する 認証フローの実際のテスト ユーザーと組織を担当する
Playwright を実行したことがあるなら、不安定なテストについて疑問に思うでしょう。ルール: 失敗した仕様を単独で再実行します。変更が失敗の原因となった場合は、修正するかビルドを中止してください。以前にすでに失敗していた場合は、キューをブロックする代わりにメンテナンス チケットに進みます。不安定かどうかにかかわらず、チェックは通過する必要があります。
これが、バイブコーディングされたプロトタイプが本番環境でばらばらになる理由です。localhost の外部でプロトタイプをテストしたことはありません。 8
両方のプロジェクトを合わせたこれまでの数字: チケット 410 件、完了 335 件、本番環境に出荷された 123 件。チケットが完了後に再度オープンされたことはありませんが、理事会は私が朝のレビューで発見した問題についてのみ知っています。失敗はもっと前に発生します。SET-212 は E2E チェックに失敗したため、エージェントは自身のコードを元に戻し、完了とマークする代わりに質問を残しました。 Settle の API スイートだけでも 2,700 を超えるテストがあります。
私が書いたかのように読める AI コード
セント

[切り捨てられた]

## Original Extract

My AI agents ship code overnight, unreviewed: autonomy gates, verification they can't fake, and graph engineering's building blocks done in plain markdown.

My AI Agents Ship Code While I Sleep. Nobody Reviews It. — Goat Squad Studios Goat Squad Studios Services Work Blog Experiences EN ES Start a project → Services →
My AI Agents Ship Code While I Sleep. Nobody Reviews It.
My AI agents ship code overnight, unreviewed: autonomy gates, verification they can't fake, and graph engineering's building blocks done in plain markdown.
I plan during the day, agents build overnight, and I review when I wake up. This is the system that decides what ships without me.
My morning starts with the board and last night's diffs. While I slept, an AI agent picked ticket SET-175 off the queue. Four cron jobs were running the same copy-pasted database scan. The agent pulled it into one shared helper, wrote six tests for it, ran the full 2,282-test suite, deployed to dev, and marked the ticket done . It even noticed a second ticket describing the same duplication and closed it.
Nobody reviewed any of that before it went out. I'm the review, the next morning: the diff, the activity log, the board.
The same agent left a different ticket unfinished. SET-184 flags appraisal-gap risk on pending deals using a property-data API. The code was built, tested, and deployed, and it shuts itself off safely if the API fails. It still didn't go live, because turning it on means paying for a data subscription, and no test can tell you if that's worth the money. That call waited for me.
A ticket field called autonomy decided which one could finish without me. The rest of this post is that field and the system around it. I also think most of the autonomy debate is about the wrong thing.
Plan all day, build overnight, review in the morning
Daytime is planning. Turning ideas into specs, writing acceptance criteria, making the calls agents will need before they need them.
Nights are loops. A background agent picks the top buildable ticket, builds it end to end, and either finishes or parks it with a question.
Code moves the same way every day. Agents branch, merge to dev, and dev deploys to the dev Amplify site plus the backend. Once a day I review the PR from dev into prod. That daily PR is the only way anything reaches prod.
Agents get everything up to dev. Prod goes through one PR a day, and I'm the reviewer.
agent branches each build on its own branch → merge to dev loops land here, never further → dev deploy Amplify dev site + backend → the daily PR dev → prod. I review it myself → prod ships after my check
Here's the ticket from this morning:
---
id : SET-175
title : Extract shared scanActiveDeals(statuses) cron helper (api)
lane : done
size : S
autonomy : auto
commits : [ 0117e92 ]
---
WHAT: 4 EventBridge crons (chaseScan, weeklyDigest, deadlineReminder,
morningBriefingNotification) hand-roll the identical active-deal scan:
loop statuses -> paginated QueryCommand on GSI1 -> MAX_DEALS_PER_RUN cap.
EXTRACT scanActiveDeals(statuses, opts) into packages/api/utils; each cron
calls it and PRESERVES its EXACT status set (chaseScan includes 'prospect';
weeklyDigest excludes it).
ACCEPTANCE (behavior-preserving; verifiable): new scanActiveDeals unit test
(pagination across pages, MAX cap enforced, status param drives the GSI1
keys); existing handler tests stay green. api unit + tsc.
Promoted from IDEA-40 (janitor tech-debt; auto-promoted per attention-ledger
AL-3 — no product decision, verifiable by build-next).
Every call is made in the spec before an agent touches it. This one even says why it qualified for auto : no product decision, and every requirement can be checked by a test.
Planning at night and reviewing in the morning is common practice now. 1 Plenty of people run this schedule and wake up to slop. The difference is what the agent is allowed to finish.
I main Claude Code as my coding agent. Every loop in this post is one of its sessions. The loops themselves are skills: instruction files in the repo, run like commands. I tweak them every time one goes wrong.
My skills: instruction files in the repo, invoked like commands. Real names from my two repos.
Why I don't code-review every AI commit
The debate right now is about how much to trust the agent. One extreme is YOLO mode: permissions off, guardrails off, agent runs free. 2 Agents are cheap, so the pitch has legs; ship fast, revert whatever breaks. But a revert doesn't un-ship bad data or win back a client's trust. The other extreme reviews every diff, which works until the agent produces more than you can read.
Most people sit at the careful end. Only 8% are comfortable with full agent autonomy. 3 Engineers who use AI on most of their work say they can fully hand off almost none of it. 4 The tools split the difference with classifiers that auto-approve safe-looking actions and flag the rest.
I think both camps are asking the wrong question. The right question is: what work can a test prove?
My planning lives in a folder: one markdown file per ticket. Agents edit the files directly. A small local server renders them as a Kanban board for me.
Every ticket has one field that decides everything: autonomy: auto | needs-input | blocked . A missing field counts as needs-input , so nothing builds by default . auto means an agent can take the ticket all the way to done with no human involved. needs-input means the ticket waits for me no matter how green its tests are.
The question is simple: can a test prove this correct, or does it need human judgment? SET-175 was a pure refactor with tests that fail if the behavior changes. Provable, so it shipped itself. SET-184 's code was just as provable and got built; the money call wasn't, so it waited.
Both tickets got built and verified; the autonomy field decided which one could ship itself.
Extract shared scanActiveDeals(statuses) cron helper
Pure refactor. Zero behavior change. Six new tests that fail if the scan behavior changes.
Flag appraisal-gap risk with AVM + comps on every pending deal
Built, tested, deployed. Shuts itself off safely if the property-data API fails.
That's the whole idea. I make the decisions a test can't settle : product direction, pricing, brand voice, legal text, security policy, client data, anything irreversible.
To be clear: unreviewed means unreviewed to dev . Agents deploy to a dev environment on dev data. The prod branch is out of their reach. shipped only moves when I run the cut myself.
Client data sits behind the same gates as pricing and copy. The blast radius of a bad overnight build is a broken dev URL.
Here's the full system as a map:
Agents and one human as nodes; lanes, promotion, and escalation as edges. Purple edges are provable and run alone; amber edges wait for a person.
raw → refined promote gate: no product call + provable gate: autonomy = auto missing field = needs-input verification ladder all rungs green Erik only product calls escalate two-question gate answer once rule retires the class product-manager agent writes board/ideas/IDEA-n (raw) idea-refiner agent pressure-tests · sizes · can kill, can't accept auto queue board/tickets · lane: next build agent TDD → E2E → deploy dev anchors tests · real browser vs dev URL · logs done (dev) shipped (prod) Erik the one human node attention ledger AL-n · answered classes
If you've used LangGraph, this looks familiar. The practice is called graph engineering: wire agent steps into a graph like this one. 5 Every serious framework landed on the same building blocks this year: typed state, conditional routing, checkpoints, interrupt gates. 6 My board has all four as markdown and rules. Frontmatter is the typed state. autonomy is the routing. needs-input is the interrupt. The activity log is the checkpoint.
What I give up is enforcement. A graph runtime can block an illegal move and tell you exactly which step failed. My rules only hold because the model follows them and I catch drift in the morning. What I get back: I can change a rule with a commit, and nothing sits between me and the model. The real difference is who picks the path, the agent or you. 5 A graph locks the route in up front. I let the agent pick its own route and put the effort into checks instead. Anthropic's own agent guide says the same thing: simple, composable patterns over frameworks. 7 Even the graph teams are moving this way: the 2026 trend in production is adding human approval gates before risky actions. 6
Verification the agent can't fake
Unreviewed shipping only works if the model can't declare done by itself. Every auto ticket has to clear checks that live outside the model:
Every auto ticket climbs all six rungs. The two starred ones are where 'the AI finished it' claims usually fall apart.
Failing vitest tests first, for every acceptance criterion: happy, null/empty, invalid, and the authorization boundary.
A Playwright spec drives the real UI via mock-auth role switching. A user-facing change with no spec is not done.
react-router typegen + tsc, clean on the touched surface and everything that consumes it.
Push the dev branch. Local edits change nothing a live check can see until this happens.
Deploy job SUCCEED, dev URL renders, Playwright green against that live URL, error monitor clean on the exercised path. All four, every time.
docs/ and CLAUDE.md updated wherever behavior or a contract changed. Shipping isn't done until the docs are.
A missing rung means the ticket isn't done, no matter how confident the agent's activity note sounds.
Here's an agent finishing SET-308 , straight from the ticket's activity log (trimmed):
2026-07-03 [claude] DONE (autopilot, dev-verified). ... ADDED the mandated
class-retiring guard packages/api/infra-guards.test.ts ... Verification
ladder cleared: infra-guards (5, incl. negative control) + full api suite
2718 pass; api+infra typecheck green; deployed SettleStack-Dev (cdk 135s);
doc-intelligence E2E PASS vs deployed dev; AWS-log live confirmation —
processed a doc on dev, settle-doc-process-dev finalized clean with ZERO
ses:SendEmail AccessDenied in the window. Docs synced. Commits 2552b89 + 4b8e352.
Most setups skip the live-verify step because their agent can't see past localhost . Mine can. It reads the build log when a deploy fails, and checks deploy status before saying it deployed. It tails CloudWatch for new errors on the surface it touched. Playwright runs against the deployed dev URL, not a local server, so the build it tested is the build users hit.
When a deploy 500s at 2am, the agent pulls the build log and fixes what it finds. The ticket's activity line says exactly what was checked, including the gaps.
Every surface my agents can read or drive. The verification row is the part the model can't fake.
The harness, code, and planning
docs atlas and ADRs, synced in the same commit AWS
Amplify Hosting deploys; reads the build log when one fails CloudWatch logs and alarms on the surface it touched Lambda 30 handlers, debugged by log stream Step Functions the doc pipeline, traced end to end DynamoDB single-table state reads while debugging S3 document buckets and extracted text Bedrock the models the product itself calls Verification
Playwright driven against the deployed dev URL Clerk real test users and orgs for auth flows
If you've run Playwright, you're wondering about flaky tests. The rule: re-run the failing spec by itself. If the change caused the failure, fix it or abort the build. If it was already failing before, it goes on a maintenance ticket instead of blocking the queue. Flaky or not, the check still has to pass.
This is why vibe-coded prototypes fall apart in production: nothing ever tested them outside localhost. 8
The numbers so far, across both projects: 410 tickets, 335 done, 123 shipped to prod. No ticket has ever been reopened after done , though the board only knows about the problems my morning reviews caught. Failures happen earlier: SET-212 failed its E2E check, so the agent reverted its own code and left a question instead of marking it done. Settle's API suite alone is over 2,700 tests.
AI code that reads like I wrote it
The st

[truncated]
