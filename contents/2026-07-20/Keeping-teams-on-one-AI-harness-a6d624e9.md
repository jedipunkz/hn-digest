---
source: "https://baselane.sh/blog/keeping-teams-on-one-harness/"
hn_url: "https://news.ycombinator.com/item?id=48976113"
title: "Keeping teams on one AI harness"
article_title: "Keeping teams on one AI harness — baselane"
author: "mohammad0omar"
captured_at: "2026-07-20T09:19:44Z"
capture_tool: "hn-digest"
hn_id: 48976113
score: 1
comments: 0
posted_at: "2026-07-20T09:03:34Z"
tags:
  - hacker-news
  - translated
---

# Keeping teams on one AI harness

- HN: [48976113](https://news.ycombinator.com/item?id=48976113)
- Source: [baselane.sh](https://baselane.sh/blog/keeping-teams-on-one-harness/)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T09:03:34Z

## Translation

タイトル: チームを 1 つの AI ハーネスに維持する
記事のタイトル: チームを 1 つの AI ハーネスに維持する — ベースレーン
説明: 私が話をするすべてのチームが同じ質問をします。全員が同じ AI セットアップを使用できるようにするにはどうすればよいですか? Wiki ページと同期スクリプトはこれで失敗します。保持されるのは、バージョン管理されたベースライン、チーム パック、およびすべてのリポジトリのマニフェストです。

記事本文:
b ベースレーン
製品
価格設定
ブログ
ドキュメント
GitHub
パック
ダウンロード
順番待ちリストに参加する
製品
チームを 1 つの AI ハーネスに維持します。
Baselane について私が話をするすべてのチームは、同じ質問をします。「全員を同じ AI セットアップに保つにはどうすればよいですか?」どのエージェントを使用するかを尋ねる人はもういません。みんな何個も使ってます。問題は、フルタイムのベビーシッターなしで、5 つのチームと数十のリポジトリがどのようにして 1 つの基準を維持できるかということです。これが私の答えです。
モハマド・オマル · 2026 年 7 月 18 日 · 6 分で読みました
通常、質問には裏話が含まれます。チーム リーダーは、フロントエンドとバックエンドが同じ AI ツールから明らかに異なるコードを生成していることに気づきました。誰かが 2 つのリポジトリの AGENTS.md ファイルを比較したところ、組織が数か月前に解決したと思われる内容について意見が一致していないことがわかりました。セキュリティ リードは、誰も承認した覚えのない権限でエージェントが実行されているのを発見しました。そして、それを持ち出した人は、すでに推奨設定を Wiki ページに書き留め、共有し、期待を寄せています。うまくいきませんでした。
私はその物語のミニチュア版を生きていたので、その物語を知っています。 1 人、23 リポジトリ、369 のスキルを含む ~/.claude フォルダー。私のリポジトリには同じ指示が 2 つもありませんでした。私が規律を守って自分のリポジトリを調整できなかったとしても、5 つのチームが調整することは決してありません。ドリフトは努力の失敗ではありません。これは、その下に機械がない場合に config が行うことです。
なぜチームは個人よりもドリフトが激しいのか
チームドリフトには、ソロドリフトにはない特性があります。それは、人々を通じてさらに複雑になります。すべての新しいリポジトリは、作成者が最後に触れたリポジトリから命令をフォークするため、規約はリリースではなく突然変異のように広がります。すべての開発者のラップトップには、自分にとって効果的であり、他の人には見えないスキル、ルール、権限のプライベート スタックが蓄積されています。誰かが規約を改善すると、その改善は 1 つのリポジトリに反映され、

トップス。組織の半数がスクロールして通過する Slack メッセージ以外に、「これはもっと良い、誰もがこれを持っているべきだ」というチャネルはありません。その結果、組織の実際の AI 設定は認識できなくなります。悪くないよ。不明。列挙することすらできないものを見直したり、確保したり、改善したりすることはできません。
うまくいかないこと (チームはすべて試しました)
恵まれた設定を説明する wiki ページは提案であり、システムではありません。分配もフィードバック ループもありません。共有テンプレート リポジトリは 0 日目を修正し、実際にドリフトが発生する 90 日目には何もしません。リポジトリごとのコンテンツは本物であり、開発者がそれを守るのは正しいため、各リポジトリのファイルを上書きする同期スクリプトは 1 週間以内に元に戻されます。すべてを 1 つの巨大な組織全体の指示ファイルに一元化すると、別の方向で失敗します。一般的すぎてデータ チームを支援できず、ノイズが多すぎてセキュリティを信頼できず、全員が所有しているため誰も所有できません。
機能: 標準のバージョンを作成し、リポジトリごとに宣言します
維持される形状は、エンジニアリング組織内の他のすべての共有事項がすでに正直に保たれている形状です。組織の共通基盤 (規律のコミット、ルールのレビュー、シークレットの処理) は、バージョン化された小規模なベースライン パックになります。各チームは、フロントエンドのコンポーネント パターン、セキュリティの強化された権限セットなど、独自のドメインのパックを所有しています。そして、すべてのリポジトリは、コミットされた harness.json で、正確なピンを使用して、使用するものを宣言します。
{
"$schema": "https://baselane.sh/harness.schema.json",
「バージョン」: 1、
"ターゲット": { "種類": "リポジトリ", "id": "チェックアウトウェブ" },
「パック」: {
"@acme/org-baseline": "3.0.0",
"@acme/frontend-taste": "1.2.0"
}
}
ツールは、各エージェントが実際に読み取るファイルにパックを展開し、マークされた管理領域内にのみ書き込むため、各リポジトリ自体のコンテンツが破壊されることはありません。マニフェストは、答えのない質問を検索に変えるものです

。このリポジトリの設定は何ですか?ファイルを読みます。現在のセキュリティ ベースラインに準拠しているのはどのチームですか?マニフェスト間でピンを比較します。それは調査ではなく台本です。規約の改善は、パックのリリースとピンバンプのプル リクエストとなり、各チームが独自のスケジュールでレビューしてマージします。リポジトリやラップトップには、何も黙って配置されるものはありません。チームがサイレント上書きを最初に発見するのは、その標準を信頼する最後の日です。
質問の背後にある質問
チームが「全員を同じセットアップに保つにはどうすればよいか」と尋ねるとき、彼らが本当に求めているのは、スタックの他のすべての部分がすでに持っているものです。つまり、何を実行しているか、誰が使用しているか、いつ変更したかに対する 1 つの答えです。 Git はコードとしてそれを提供しました。パッケージマネージャーは依存関係のためにそれを提供しました。 AI エージェントを操作する指示、スキル、権限は依然として民間伝承に基づいて実行されており、それが毎日リポジトリにコードを書き込む構成です。同じ機械を与えてください。バージョンを付け、固定し、プル リクエストによって配布し、アサートする代わりに採用を測定します。それがすべての答えであり、意図的に退屈なものです。
マニフェスト仕様は github.com/baselane-sh/harness.json で公開されています。これは、ツールを使用せずに役立ちます。つまり、人間またはエージェントが読み取れる 1 つの小さなファイルだけです。残りの部分 (監査、具体化、ドリフトチェック、すべてのチームにわたるフリート ビュー) を自動化するツールは、 Baselane です。チームがこの質問をしている場合、1 つのリポジトリで npx Baselane Audit を実行すると、実際の立ち位置を確認する 10 分間の方法になります。段階的なロールアウトは、関連記事「5 つのチームにわたって AI ハーネス構成を管理する方法」にあります。
AI エンジニアリングのコントロール プレーン。

## Original Extract

Every team I talk to asks the same question: how do we keep everyone on the same AI setup? Wiki pages and sync scripts fail at this. What holds is a versioned baseline, team packs, and a manifest in every repo.

b baselane
Product
Pricing
Blog
Docs
GitHub
Packs
Download
Join the waitlist
Product
Keeping teams on one AI harness .
Every team I talk to about baselane asks a version of the same question: how do we keep everyone on the same AI setup? Nobody asks which agent to use anymore. Everyone uses several. The question is how five teams and dozens of repos stay on one standard without a full-time babysitter. This is my answer.
Mohammad Omar · Jul 18, 2026 · 6 min read
The question usually arrives with a backstory. A team lead noticed that frontend and backend produce visibly different code from the same AI tools. Someone diffed two repos' AGENTS.md files and found they disagreed about things the org supposedly settled months ago. A security lead found an agent running with permissions nobody remembers approving. And whoever brings it up has already tried writing down the recommended setup on a wiki page, sharing it, and hoping. It didn't work.
I recognize the story because I lived a miniature version of it. One person, 23 repos, a ~/.claude folder with 369 skills. No two of my repos had the same instructions. If I couldn't keep my own repos aligned through discipline, five teams never will. Drift isn't a diligence failure. It's just what config does when there's no machinery underneath it.
Why teams drift harder than individuals
Team drift has a property solo drift doesn't: it compounds through people. Every new repo forks its instructions from whichever repo its creator touched last, so conventions spread like mutations rather than releases. Every developer's laptop accumulates a private stack of skills, rules and permissions that works great for them and is invisible to everyone else. When someone improves a convention, the improvement lands in one repo and stops. There is no channel for "this is better, everyone should have it" besides a Slack message that half the org scrolls past. The result is that the org's real AI setup is unknowable. Not bad. Unknowable. You can't review, secure, or improve what you can't even enumerate.
What doesn't work (teams have tried it all)
A wiki page describing the blessed setup is a suggestion, not a system. It has no distribution and no feedback loop. A shared template repo fixes day zero and does nothing for day ninety, which is when drift actually bites. A sync script that overwrites each repo's files gets reverted within a week, because per-repo content is real and developers are right to defend it. Centralizing everything into one giant org-wide instructions file fails in the other direction: too generic to help the data team, too noisy for security to trust, owned by everyone and therefore no one.
What works: version the standard, declare it per repo
The shape that holds is the one that already keeps every other shared thing in an engineering org honest. The org's common ground (commit discipline, review rules, secret handling) becomes a small versioned baseline pack . Each team owns a pack for its own domain: frontend's component patterns, security's hardened permission set. And every repo declares what it uses, with exact pins, in a committed harness.json :
{
"$schema": "https://baselane.sh/harness.schema.json",
"version": 1,
"target": { "kind": "repo", "id": "checkout-web" },
"packs": {
"@acme/org-baseline": "3.0.0",
"@acme/frontend-taste": "1.2.0"
}
}
Tooling expands the packs into the files each agent actually reads, and it writes only inside marked managed regions, so each repo's own content never gets clobbered. The manifest is what turns the unanswerable questions into lookups. What's this repo's setup? Read the file. Which teams are on the current security baseline? Compare pins across manifests. That's a script, not a survey. Improving a convention becomes a pack release plus pin-bump pull requests, reviewed and merged by each team on its own schedule. Nothing lands silently, on a repo or a laptop. The first silent overwrite a team discovers is the last day they trust the standard.
The question behind the question
When a team asks "how do we keep everyone on the same setup," what they're really asking for is the thing every other part of their stack already has: one answer to what are we running, who's on it, and when did it change . Git gave them that for code. Package managers gave them that for dependencies. The instructions, skills and permissions steering their AI agents still run on folklore, and that's configuration that writes code into their repos every day. Give it the same machinery. Version it, pin it, distribute it by pull request, and measure adoption instead of asserting it. That's the whole answer, and it's boring on purpose.
The manifest spec is open at github.com/baselane-sh/harness.json . It's useful with zero tooling: one small file any human or agent can read. The tooling that automates the rest (audit, materialize, drift-check, a fleet view across every team) is baselane . If your team is asking this question, npx baselane audit on one repo is the ten-minute way to see where you actually stand. The step-by-step rollout is in the companion post: How to manage AI-harness config across five teams .
The control plane for AI engineering.
