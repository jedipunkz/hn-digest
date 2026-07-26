---
source: "https://github.com/De-Cri/dont-kick-me-out"
hn_url: "https://news.ycombinator.com/item?id=49060721"
title: "Show HN: I made Claude quiz me on my own code before showing it"
article_title: "GitHub - De-Cri/dont-kick-me-out: Copilot, not autopilot. A Claude skill that keeps you thinking while AI writes your code. · GitHub"
author: "Ssamu3909"
captured_at: "2026-07-26T18:57:30Z"
capture_tool: "hn-digest"
hn_id: 49060721
score: 2
comments: 0
posted_at: "2026-07-26T18:11:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I made Claude quiz me on my own code before showing it

- HN: [49060721](https://news.ycombinator.com/item?id=49060721)
- Source: [github.com](https://github.com/De-Cri/dont-kick-me-out)
- Score: 2
- Comments: 0
- Posted: 2026-07-26T18:11:00Z

## Translation

タイトル: HN を表示: 表示する前に、クロードに自分のコードについてクイズを出しました。
記事のタイトル: GitHub - De-Cri/dont-kick-me-out: 自動操縦ではなく副操縦。 AI がコードを書いている間、考え続けるクロード スキル。 · GitHub
説明: 自動操縦ではなく、副操縦士です。 AI がコードを書いている間、考え続けるクロード スキル。 - デクリ/キックミーアウトしないでください

記事本文:
GitHub - De-Cri/dont-kick-me-out: 自動操縦ではなく副操縦。 AI がコードを書いている間、考え続けるクロード スキル。 · GitHub
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
デクリ
/
追い出さないで
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .claude-plugin .claude-plugin skill/ dont-kick-me-out skill/ dont-kick-me-out .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自動操縦ではなく副操縦士。 AI がコードを書いている間、考え続けるクロード スキル。
AI が機能するコードを生成すると、「このコードのすべての選択肢を理解できているか?」というレビューをやめてしまいがちです。そして「それがもっともらしく見えるかどうか」を検討し始めます。 。そのギャップは理解の負債です。今日では目に見えませんが、何かが壊れて理由がわからない日、またはコードレビューで設計上の決定を擁護できない日には残酷です。
これは「クロードがもっと説明する」ではありません。良い説明を読むことはまだ受動的です。このスキルは 2 つのアイデアに基づいて構築されています。
理解は読むことではなく、生み出すことから生まれます。重要な瞬間に、クロードは方向を変えます。解決策を見つける前に、動作や障害モードを予測します。失敗した予測は、最も価値の高い学習の瞬間です。
自動操縦は段階的に行われます。ステップ 1 で離脱する人はいません。ステップ 6 で離脱し、疲れて「ok」と入力する速度がますます速くなります。このスキルは、チェックポイントを機械的に起動するのではなく、エンゲージメントのフェードアウトを検出して適応します。
最初に番号が付けられた計画 — 後のすべてのインタラクションはそれに固定されます
コードの後ではなく、コードの前のトレードオフ — このアプローチがなぜ行われるのか、代わりに何が実行可能で、何が放棄されるのか
アンカーされたチェックポイント — 「特にポイント 2 について何か疑問はありますか?」という質問ではなく、一般的な「見た目は良いですか?」という質問は決してありません。
予測の瞬間 — 2 ～ 3 ステップごとに、クロードが示す前に行動を推測します。
疲労検出 — 3 回連続の「OK」とクロードが名付けたもので、儀式的な遵守ではなく正直なダウンシフトを提供します
終わりの要約 b

あなた — 最後に、重要な決定事項を要約します。クロードが隙間を埋める
あなた: /login にレート制限を追加します
クロード: (コードの壁) 完了!他に何か必要な場合はお知らせください。
あなた: (流し読み、マージ、忘れる)
クロード: 計画: 1) 制限戦略と状態の所在、2) ミドルウェア、3) /login への接続、4) 拒否された要求への応答、5) 境界テスト。
ポイント 1 では、実際の決定は、Redis のスライディング ウィンドウとインプロセス トークン バケットのどちらを選択するかです。 Redis はネットワーク ホップを追加しますが、インプロセス リミッターはロード バランサーの背後にある 2 つのインスタンスで静かに壊れます。ログインの不正使用がまさに問題となります。単一インスタンスを永遠に?私なら選択をひっくり返します。あなたにとってはどれが当てはまりますか？
…
クロード: ステップ 4 の前に、簡単な予測をしてみます。ウィンドウが実装されている場合、クライアントは 6 回目の試行で何を取得しますか。また、カウンターは一度にリセットされるのでしょうか、それとも段階的にリセットされるのでしょうか?
プライベート プロジェクトの実際のセッションからの 2 つの交換。ドメイン名は変更され、私の返信はイタリア語から翻訳されています。それ以外の点では、やり取り自体は、私が間違っている部分も含めて編集されていません。
クロードは、モジュールレベルの関数のセットをクラスに移動していました。テストでは、モジュールレベルの名前にモンキーパッチを適用して、どのブランチが選択されるかをキャプチャします。
クロード: make_chart をクラスに移動すると、本能的に self._renderer.render(data) を呼び出します。レンダラはすでにオブジェクトとして持っているのに、なぜモジュール レベルの関数を使用するのでしょうか?
書く前に: テストは失敗するのか、合格しても嘘をつくのか、それとも何も変わらないのか?
私：何も変わらないよ。以前はデータをモジュールに渡していたのですが、今ではデータをオブジェクトに渡しているので、self.render で十分です。
クロード: 間違っています — そして、それは 3 つの中で最も厄介です。それは通り過ぎますが、あなたに嘘をつきます。
patch.object(mod, "render_chart", cap) は関数を置き換えません。モジュールの namesp 内の render_chart という名前の属性を置き換えます。

エース。現在、これが機能するのは、make_chart が呼び出し時にそのグローバルを解決するためです。 self._renderer.render(...) を経由しても、その検索は実行されません。パッチは、もう誰も読まない名前に置かれています。
モジュール経由 -> result='PNG(pie)' Captured={'chart_type': 'pie'}
オブジェクト経由 -> 結果='PNG(パイ)' キャプチャ=何もありません
どちらの方法でも同じ出力になります。実際のテストでは、すべてのケースでキャプチャされたタイプが None に設定されて PASS が報告されることを意味します。テストが検証するために存在する唯一のものは、サイレントに消えます。
私が議論した後に却下された推奨事項
高価な部分を再実行せずに、保存された結果を再スコアリングするスクリプトを保持するかどうかを決定します。
クロード: オプション A — 削除 — は、今後修正するたびに完全なライブ再実行が必要になることを意味します。
私: でも、オプション A は完全な再実行が必要だとおっしゃっていますが、変更されていないケースをスキップするフラグを既に追加しました。
クロード: 本当です、そしてあなたが私を訂正するのは正しいことです - 「完全な再実行」は誇張でした。ただし、キャッシュ署名が実際に何をカバーしているかを見てください。キャッシュ署名はケースのコンテンツのみに対して計算されます。
ケースを変更すると、署名が変更され、そのケースのみが再実行されます。スキップはあなたの言うとおりに機能します。
スコアリング コードを変更しても、サインは 1 ミリも動きません。すべてのケースはスキップされ、壊れたスコアラーで計算された古いスコアは、何も起こらなかったかのようにファイルに残ります。
クロードの勧めはそれを保管することでした。とにかく、私はそれを削除しました。何を交換するのかを正確に知っていました。それが決定を下すこととそれを受け入れることの違いです。
/プラグインマーケットプレイス追加De-Cri/dont-kick-me-out
/プラグインのインストール dont-kick-me-out@dont-kick-me-out
その後の修正は、/plugin マーケットプレイスのアップデートで提供されます。再インストールは必要ありません。
mkdir -p ~ /.claude/skills/dont-kick-me-out
カール -sL https://raw.githubusercontent.com/De-Cri/dont-kick-me-out/main/skills/dont-kick-me-out/SK

ILL.md \
-o ~ /.claude/skills/dont-kick-me-out/SKILL.md
クロード・アイ
リリースから dont-kick-me-out.zip をダウンロードします
設定→機能→スキル→スキルのアップロード
コーディングに関する会話を開始します - 重要なタスクが発生すると自動的にトリガーされます
名前の変更、タイプミスの修正、フォーマット、機械的な移行 — チェックポイントやエッセイはありません。ループは、PR 内で査読者が異議を唱える可能性のある選択肢が少なくとも 1 つある場合にのみ適用されます。強度は、コードの長さではなく、意思決定の賭け金を追跡します。
自動操縦ではなく副操縦士。 AI がコードを書いている間、考え続けるクロード スキル。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Copilot, not autopilot. A Claude skill that keeps you thinking while AI writes your code. - De-Cri/dont-kick-me-out

GitHub - De-Cri/dont-kick-me-out: Copilot, not autopilot. A Claude skill that keeps you thinking while AI writes your code. · GitHub
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
De-Cri
/
dont-kick-me-out
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .claude-plugin .claude-plugin skills/ dont-kick-me-out skills/ dont-kick-me-out .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Copilot, not autopilot. A Claude skill that keeps you thinking while AI writes your code.
When AI generates working code, it's easy to stop reviewing for "do I understand every choice in this?" and start reviewing for "does it look plausible?" . That gap is comprehension debt — invisible today, brutal the day something breaks and you don't know why, or the day you can't defend a design decision in code review.
This is not "Claude explains more." Reading a good explanation is still passive. This skill is built on two ideas:
Understanding comes from producing, not reading. At key moments, Claude reverses direction: you predict the behavior or failure mode before seeing the solution. A failed prediction is the highest-value learning moment there is.
Autopilot is gradual. Nobody disengages at step 1 — they disengage at step 6, tired, typing "ok" faster and faster. The skill detects fading engagement and adapts instead of firing checkpoints mechanically.
Numbered plan first — every later interaction anchors to it
Tradeoffs before code, not after — why this approach, what was viable instead, what's being given up
Anchored checkpoints — "any doubts on point 2 specifically?", never a generic "look good?"
Prediction moments — every 2–3 steps, you guess the behavior before Claude shows it
Fatigue detection — three bare "ok"s in a row and Claude names it, offering an honest downshift instead of ritual compliance
Closing recap by you — at the end, you summarize the key decisions; Claude fills the gaps
You: add rate limiting to /login
Claude: (wall of code) Done! Let me know if you need anything else.
You: (skims, merges, forgets)
Claude: Plan: 1) limiting strategy & where state lives, 2) middleware, 3) wire to /login, 4) rejected-request response, 5) boundary tests.
On point 1, the real decision: sliding-window in Redis vs in-process token bucket. Redis adds a network hop — but an in-process limiter silently breaks with two instances behind a load balancer, and login abuse is exactly where that matters. Single-instance forever? I'd flip the choice. Which holds for you?
…
Claude: Before step 4, quick prediction: with the window as implemented, what does a client get on their 6th attempt — and does the counter reset all at once or gradually?
Two exchanges from real sessions on a private project. Domain names are changed and my replies are translated from Italian — the exchanges themselves are otherwise unedited, including the part where I get it wrong.
Claude was moving a set of module-level functions into a class. A test monkeypatches a module-level name to capture which branch gets taken.
Claude: When I move make_chart into the class, the instinct is to call self._renderer.render(data) — I already have the renderer as an object, so why go through the module-level function?
Before I write it: does the test blow up, pass but lie to you, or change nothing?
Me: nothing changes. Before we were handing data to a module, now we hand it to an object, so self.render is enough.
Claude: Wrong — and it's the nastiest of the three: it passes but lies to you.
patch.object(mod, "render_chart", cap) doesn't replace the function. It replaces the attribute named render_chart in the module's namespace. That works today only because make_chart resolves that global at call time. Going through self._renderer.render(...) never performs that lookup — the patch is sitting on a name nobody reads anymore.
via module -> result='PNG(pie)' captured={'chart_type': 'pie'}
via object -> result='PNG(pie)' captured=NOTHING
Same output either way. On the real test that means every case reports PASS with the captured type set to None — the one thing the test exists to verify, gone, silently.
A recommendation I argued with, then overruled
Deciding whether to keep a script that re-scores saved results without re-running the expensive part.
Claude: Option A — delete it — means every future correction costs you a full live re-run.
Me: but you say option A costs a full re-run, and I already added a flag that skips unchanged cases.
Claude: True, and you're right to correct me — "full re-run" was an overstatement. But look at what the cache signature actually covers: it's computed over the case content only.
Change a case, and the signature changes: only that case re-runs. Your skip works exactly as you say.
Change the scoring code , and the signature doesn't move a millimetre. Every case is skipped, and the old scores — computed with the broken scorer — stay in the file as if nothing had happened.
Claude's recommendation was to keep it. I deleted it anyway, knowing precisely what I was trading away. That is the difference between making a decision and accepting one.
/plugin marketplace add De-Cri/dont-kick-me-out
/plugin install dont-kick-me-out@dont-kick-me-out
Later fixes reach you with /plugin marketplace update — no reinstall.
mkdir -p ~ /.claude/skills/dont-kick-me-out
curl -sL https://raw.githubusercontent.com/De-Cri/dont-kick-me-out/main/skills/dont-kick-me-out/SKILL.md \
-o ~ /.claude/skills/dont-kick-me-out/SKILL.md
Claude.ai
Download dont-kick-me-out.zip from Releases
Settings → Capabilities → Skills → Upload skill
Start any coding conversation — it triggers automatically on non-trivial tasks
Renames, typo fixes, formatting, mechanical migrations — no checkpoints, no essays. The loop applies only when there's at least one choice a reviewer could contest in a PR. Intensity tracks the stakes of the decision, not the length of the code.
Copilot, not autopilot. A Claude skill that keeps you thinking while AI writes your code.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
