---
source: "https://johnstawinski.com/2026/06/18/repo-jacking-anthropics-claude-community-plugins-and-the-shas-that-saved-them/"
hn_url: "https://news.ycombinator.com/item?id=48600557"
title: "Repo-Jacking Anthropic's Claude Community Plugins (and the SHAs That Saved Them)"
article_title: "Repo-jacking Anthropic’s Claude Community Plugins (And the SHAs That Saved Them) – John Stawinski IV"
author: "cyberbender"
captured_at: "2026-06-19T18:43:43Z"
capture_tool: "hn-digest"
hn_id: 48600557
score: 2
comments: 0
posted_at: "2026-06-19T16:55:11Z"
tags:
  - hacker-news
  - translated
---

# Repo-Jacking Anthropic's Claude Community Plugins (and the SHAs That Saved Them)

- HN: [48600557](https://news.ycombinator.com/item?id=48600557)
- Source: [johnstawinski.com](https://johnstawinski.com/2026/06/18/repo-jacking-anthropics-claude-community-plugins-and-the-shas-that-saved-them/)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T16:55:11Z

## Translation

タイトル: Repo-Jacking Anthropic のクロード コミュニティ プラグイン (およびそれらを保存した SHA)
記事のタイトル: Anthropic のクロード コミュニティ プラグインのリポジャッキング (そしてそれらを救った SHA) – John Stawinski IV
説明: いくつかの Claude Community Plugin はリポジャッキングに対して脆弱でした。コードの直接インストール パスは SHA チェックによって軽減されましたが、Claude Code の「プラグイン UI の表示」機能はユーザーをリポジトリジャックされたリポジトリにリダイレクトし、信頼できるコミュニティ プラグインを活用するソーシャル エンジニアリングのベクトルを開きます。
[切り捨てられた]

記事本文:
Anthropic の Claude コミュニティ プラグインのリポジャッキング (およびそれらを保存した SHA)
いくつかの Claude Community Plugin はリポジャッキングに対して脆弱でした。コードの直接インストール パスは SHA チェックによって軽減されましたが、Claude Code の「プラグイン UI の表示」機能はユーザーをリポジトリジャックされたリポジトリにリダイレクトし、信頼できるコミュニティ プラグインを活用するソーシャル エンジニアリングのベクトルを開きます。
私の経験に基づくと、サプライ チェーンとソーシャル エンジニアリングは、組織を侵害する最も簡単な方法です。
2025 年と 2026 年には、TeamPCP などがオープンソース ソフトウェアで蔓延する中、脅威アクターはサプライ チェーンの角度に注目しました。
AI ツールはサプライチェーンの問題を悪化させています。以前は、人間が自分のマシンにランダムなものをインストールできました。今、私たちは彼らに代わってこれを行うことができるエージェントを持っています（そして、実際に「-dangerously-skip-permissions」なしでクロードコードを使用する人は誰ですか？）
私たちは今後数年間にわたって、AI ツールに関するサプライ チェーンの多くの教訓を学び直すことになるでしょう。
数週間前、私はエージェント プラットフォームのサプライ チェーンの脆弱性を探すことで、脆弱性の研究に再び足を踏み入れることにしました。私が最初に発見した脆弱性の 1 つは、Claude Code Community Plugins にありました。
クロード コード コミュニティ プラグインとは何ですか?
Claude Code コミュニティ プラグインは、Claude Code を拡張するために個人や企業が構築するアドオンです。それぞれにカスタム コマンド、AI エージェント、MCP サーバー、自動フックなどがバンドルされており、すべて 1 ステップでインストールでき、チームメイトと簡単に共有できます。 Anthropic は公式ディレクトリを運営していますが、Anthropic のレビューに合格した後は、誰でも独自のプラグインをパブリック コミュニティ マーケットプレイスに送信できます。これを使用するには、コマンドを使用してマーケットプレイスを Claude Code に追加し、そこから必要なプラグインをインストールします。
クロード

e Code コミュニティ プラグイン マーケットプレイスは、Anthropic によって管理されている単なる GitHub リポジトリであり、その中心となるのは、プラグインをリストし、それらをホストする GitHub リポジトリを示す単一のファイルです。 Claude Code を通じてプラグインをインストールすると、marketplace.json が指すリモート リポジトリからコードが自動的に取得されます。
「deep-research」プラグインのソース コードは https://github.com/oduffy-delphi/deep-research-claude に保存されています。
他の人のリポジトリへのポインタは、時間の経過とともに静かに朽ちていく一種の信頼です。人々は自分の GitHub アカウントの名前を変更します。人々はそれらを削除します。マーケットプレイスへの参加者は気づきません。そこからが始まります。
レポジャッキングは腐敗を糧とするサプライチェーン攻撃です。
GitHub アカウントの名前が変更または削除されても、古い所有者/リポジトリのパスは単に消えるわけではありません。 GitHub はリダイレクトを維持するため、既存のリンクは機能し続けますが、これは誰も古いユーザー名を再登録しない限り機能し続けます。誰かがその放棄された名前を主張し、その下にリポジトリを再作成した瞬間、リダイレクトはなくなり、その人が名前空間を所有します。まだ古いパスにリンクしている人は、攻撃者にリンクしていることになります。
GitHub には、プロジェクトの移動または削除後にその名前空間を予約するなど、リポジャッキング攻撃に対するいくつかの保護機能がありますが、これはスター数とトラフィックが十分に高い人気のあるプロジェクトに限られます。
リポジャッキングはよく知られた手法であり、多くの大規模プロジェクトに影響を与えてきました。興味深いのは、トリックそのものではありません。信頼できる公式情報源がまだ注目されている名前を示している場所を見つけることです。
プラグイン マーケットプレイスは非常に良い場所です。
私 (エージェント) は、marketplace.json を調べて、所有者がなくなったリポジトリを参照しているエントリを探しました。数人が戻ってきました。以下のプラグインエントリはすべて次のことを指しています

所有者が名前変更または削除され、名前空間が要求可能なままになっている所有者/リポジトリのパス:
チャーリー・グリーンマン/ghostlty-dynamic-主題
Chipkorvyn/戦略コンサルタント
oduffy-delphi/ディープリサーチ-クロード
それぞれは、Anthropic の公式マーケットプレイス内にぶら下がっている参照です。それぞれは、名前を登録し、リポジトリを再作成し、エコシステムが依然として正当なものとして扱う場所から必要なものを提供するためのオープンな招待状です。
概念実証: deep-research-claude
deep-research-claude プラグインは github.com/oduffy-delphi/deep-research-claude を指していました。 oduffy-delphi アカウントはもう存在しませんでした。実際のプロジェクトは dbc-oduffy/deep-research-claude に移動していましたが、マーケットプレイスにはそのメモが届きませんでした。
組織を登録し、その下に deep-research-claude を再作成し、正規のリポジトリのクローンをシードして、プラグインが引き続きユーザーの期待どおりに機能するようにしました。私が変更したのは、README の先頭にある「PoC — John Stawinski」という行だけです。実際の攻撃者はさらに大幅に変化するでしょう。
ユーザーの席からは、この攻撃は何の変哲もないように見えます。
マーケットプレイスを追加します: クロード プラグイン マーケットプレイス add anthropics/claude-plugins-community
プラグインをインストールします: /plugin install deep-research@claude-community
私のリポジャックされたリポジトリ、PoC バナー、その他すべてにアクセスします。
警告も摩擦もありません。マーケットプレイスが正式になり、ホームページがオープンします。信頼は下までずっと流れていきます。
これは純粋なフィッシングの原始的手法です。ホームページをリポジャックした攻撃者は、Anthropic のエコシステムの暗黙の承認の下でレンダリングされる攻撃者制御のページを所有しており、ユーザーがすでに「このプラグインをセットアップ中」モードになっているまさにその瞬間に、インストール フローから 1 回のクリックでアクセスできます。 「実はここが核心だ」と言うのに最適な場所です

これをインストールする方法はありません」と攻撃したり、誰かに毒された設定を渡したり、ターゲットがすでに自分がどこか正当な場所にいると信じている場合に機能する十数のことのいずれかを実行することもできます。
ユーザーがホームページを開いて、そこに仕掛けられた攻撃者による悪意のある指示に従うと、ホストが侵害されてしまいます。
ただし、ユーザーがホームページを開かず、代わりに Claude Code によるプラグインの自動インストールを許可した場合、ユーザーは輝く鎧を着た SHA によって保存されます。
マーケットプレイスのエントリを詳しく見ると、リポジトリがコミット SHA に固定されていることがわかります。
プラグインのインストール パスは、特定のコミット SHA に固定されます。 Claude Code がプラグイン コードをフェッチするとき、所有者/リポジトリを信頼するだけでなく、期待する正確なコンテンツのハッシュを信頼します。再作成されたリポジトリには同じ名前と同じファイルを含めることができますが、内容が変更されたピン留めされたコミットを再現することはできません。整合性チェックは失敗し、悪意のあるインストールは行われません。
これは正しい設計であり、Claude Community Plugin のリポジャックが、プラグインをインストールする将来のすべてのユーザーに対してすぐに任意のコード実行にならない理由です。プラグイン エコシステムの防御について 1 つだけ覚えている場合は、コミット SHA などの不変の参照にソースを固定してください。
この脆弱性を報告した後、私はプラグインのエコシステムについて考え始めました。プラグインのメンテナがコードを更新するのはごく一般的なことでしょう。では、Anthropic はいつ、どのように SHA を更新するのでしょうか?
結局、Anthropic のエンジニアは、このような大量のプル リクエストを送信し、ボットが最新の SHA を調べて追跡し、プラグインを検証して、新しい SHA をマーケットプレイスに提供することが判明しました。
ボットが新しい SHA ポインティングを取り込むようなプラグイン検証プロセスをだますのは難しいとは思えません。

リポジャックされたリポジトリの悪意のあるコミットに。この理論を検証する時間がありませんでしたが、SHA は「輝く鎧を着た騎士」というよりも、「今のところはあなたを守ってくれるジェイミー・ラニスター」のような存在ではないかと思います。
これを書く前に、実際の攻撃者がそれらを悪用できないように、影響を受ける各プレフィックスを GitHub 組織として登録しました。 PoC に必要な oduffy-delphi/deep-research-claude を除いて、特定のリポジトリは作成しないままにしたため、GitHub のリダイレクト ループは正規のプロジェクトが現在存在する場所に解決されます。
攻撃者の手に渡らないように名前を不法占拠することは、参照がクリーンアップされる間の安価な保険です。
私はこれらすべてを、バグ報奨金プログラムを通じて Anthropic に報告しました。
2026年5月25日 報告書を提出しました。
5月25日 事前審査通過。
5 月 25 日 参考情報としてクローズ、対象外。
Anthropic が問題を解決した理由は、依存関係のハイジャックは明示的に対象外であり、SHA ピン留めによってインストール パスが保護され、残りの懸念事項である、所有権が変更されたページへのホームページ リンクをたどるユーザーはソーシャル エンジニアリングに該当するというものでした。
コミュニティ プラグイン マーケットプレイスを信頼すべきですか?
組織は、Claude プラグイン マーケットプレイスを、たとえ Anthropic によって精査されたものであっても、他のサードパーティのコードを信頼できないものとして扱うのと同じように扱う必要があります。どこかの公式リストに載っているからといって、それが安全であるというわけではありません (そこにあるすべての悪意のある「検証済み」VSCode プラグインを見てください)。
プラグイン エコシステムは、パッケージ レジストリの最も古い問題を継承しています。私たちは今後数年間にわたって、AI ツールに関するサプライ チェーンの多くの教訓を学び直すことになるでしょう。
これはその中でも安いものの一つでした。
X で共有 (新しいウィンドウで開きます)
×
Facebook で共有 (新しいウィンドウで開きます)
顔

本
コメントを書く...
メールアドレス (必須)
お名前 (必須)
ウェブサイト
購読する
購読済み
ジョン・スタウィンスキー4世
すでに WordPress.com アカウントをお持ちですか?今すぐログインしてください。

## Original Extract

Several Claude Community Plugins were vulnerable to repo-jacking. The direct code installation path was mitigated by SHA checks, but Claude Code’s “view plugin UI” feature would redirect users to the repo-jacked repository, opening up a social engineering vector leveraging trusted community plugins.
[truncated]

Repo-jacking Anthropic’s Claude Community Plugins (And the SHAs That Saved Them)
Several Claude Community Plugins were vulnerable to repo-jacking. The direct code installation path was mitigated by SHA checks, but Claude Code’s “view plugin UI” feature would redirect users to the repo-jacked repository, opening up a social engineering vector leveraging trusted community plugins.
Based on my experience, supply chain and social engineering are the easiest ways to compromise an organization.
In 2025 and 2026, threat actors caught on to the supply chain angle, as TeamPCP and others ran rampant on open-source-software.
AI tooling has exacerbated the supply chain problem. Before, we had humans that could go install random stuff on their machines. Now, we have agents that can do this for them (and who really uses Claude Code without “–dangerously-skip-permissions”)?
We’re going to be relearning a lot of supply chain lessons in AI tooling over the next few years.
Several weeks ago, I decided to dip my toes back into vulnerability research by looking for supply chain vulnerabilities in agentic platforms. One of the first vulnerabilities I picked up was in the Claude Code Community Plugins.
What Are Claude Code Community Plugins?
Claude Code community plugins are add-ons that people and companies build to extend Claude Code. Each one bundles together things like custom commands, AI agents, MCP servers, and automated hooks, all installable in one step and easy to share with teammates. Anthropic runs an official directory, but anyone can submit their own plugin to a public community marketplace after it passes an Anthropic review. To use it, you add the marketplace to Claude Code with a command, then install whatever plugins you want from it.
The Claude Code community plugin marketplace is just a GitHub repository maintained by Anthropic , and the heart of it is a single file that lists plugins and points at the GitHub repositories that host them. Installing a plugin through Claude Code will automatically retrieve the code from the remote repository pointed to by marketplace.json.
Source code for the “deep-research” plugin was stored at https://github.com/oduffy-delphi/deep-research-claude .
Pointers to other people’s repositories are the kind of trust that quietly rots over time. People rename their GitHub accounts. People delete them. The marketplace entry doesn’t notice. That’s where this starts.
Repo-jacking is a supply chain attack that lives off of rot.
When a GitHub account is renamed or deleted, the old owner/repo path doesn’t just vanish. GitHub keeps a redirect alive so existing links keep working, but only as long as nobody re-registers the old username. The moment someone claims that abandoned name and recreates a repo under it, the redirect is gone and they own the namespace. Anyone still linking to the old path is now linking to the attacker.
GitHub has some protections against repo-jacking attacks, like reserving namespaces of projects after they are moved or deleted, but only for popular project with a high enough star count and traffic.
Repo-jacking is a well known technique, and it has bitten plenty of large projects. The interesting part is never the trick itself– it’s finding a place where a trusted, official source still points at a name that’s up for grabs.
A plugin marketplace is a very good place to look.
I (my agents) looked through marketplace.json for entries referencing repositories that no longer had an owner. A handful came back. The plugin entries below all pointed at owner/repo paths where the owner had been renamed or deleted, leaving the namespace claimable:
CharlieGreenman/ghostlty-dynamic-themes
Chipkorvyn/Strategy-consultant
oduffy-delphi/deep-research-claude
Each one is a dangling reference inside Anthropic’s official marketplace. Each one is an open invitation to register the name, recreate the repo, and serve whatever you want from a location that the ecosystem still treats as legitimate.
Proof of concept: deep-research-claude
The deep-research-claude plugin pointed at github.com/oduffy-delphi/deep-research-claude. The oduffy-delphi account no longer existed. The real project had moved to dbc-oduffy/deep-research-claude, but the marketplace never got the memo.
I registered the org, recreated deep-research-claude under it, and seeded it with a clone of the legitimate repository so the plugin would still function exactly as a user expected. The only thing I changed was a line at the top of the README: “PoC — John Stawinski”. A real attacker would change considerably more.
From a user’s seat, the attack looks like nothing at all:
Add the marketplace: claude plugin marketplace add anthropics/claude-plugins-community
Install the plugin: /plugin install deep-research@claude-community
Land on my repo-jacked repository, PoC banner and all.
No warning, no friction. The marketplace is official, and the homepage opens. Trust flows the entire way down.
That’s a clean phishing primitive. An attacker who has repo-jacked the homepage now has an attacker-controlled page that renders under the implicit endorsement of Anthropic’s ecosystem, reachable in one click from the install flow, at the exact moment a user is already in “I’m setting up this plugin” mode. It’s the perfect place to say “actually, here’s the correct way to install this,” or to hand someone a poisoned config, or to do any of the dozen things that work when a target already believes they’re somewhere legitimate.
If a user opens the homepage, follows malicious instructions an attacker planted there, then their host will be compromised.
However, if the user doesn’t open the homepage, and instead allows Claude Code to automatically install the plugin, they will be saved by a SHA in shining armor.
Looking more closely at marketplace entries, we can see that repos are pinned to a commit SHA.
The plugin install path is pinned to a specific commit SHA. When Claude Code fetches the plugin code, it isn’t just trusting owner/repo, it’s trusting a hash of the exact contents it expects. A recreated repository can carry the same name and the same files, but it cannot reproduce the pinned commit with any modified content. The integrity check fails and the malicious install doesn’t happen.
This is the right design, and it’s the reason a repo-jack of a Claude Community Plugin doesn’t immediately turn into arbitrary code execution on every future user who installs the plugin. If you only remember one thing about defending a plugin ecosystem, pin your sources to immutable references, like commit SHAs.
After submitting this vulnerability, I started thinking about the plugin ecosystem. It must be pretty common for plugin maintainers to update their code, so when and how does Anthropic update the SHAs?
It turns out, Anthropic engineers submit massive pull requests, like this one , where a bot will go through and track down the latest SHA, validate the plugin, and then provide the new SHA for the marketplace.
I doubt that it would be hard to fool the plugin validation process such that the bot ingests a new SHA pointing to a malicious commit of a repo-jacked repository. I didn’t have time to test this theory, but I suspect the SHAs are less of a “knight in shining armor” and more of a “Jamie Lannister who will protect you…..for now.”
Before writing any of this up, I registered each of the affected prefixes as a GitHub organization so an actual attacker couldn’t exploit them. With the exception of oduffy-delphi/deep-research-claude, which I needed for the PoC, I left the specific repositories uncreated so GitHub’s redirect loop still resolves to wherever the legitimate projects live now.
Squatting on names to keep them out of attacker hands is cheap insurance while the references get cleaned up.
I reported all of this to Anthropic through their bug bounty program.
May 25, 2026 Report submitted.
May 25 Passed preliminary review.
May 25 Closed as Informative, out of scope.
Anthropic’s reasoning for closing the issue was that dependency hijacking is explicitly out of scope, the SHA pinning protects the install path, and the remaining concern, a user following a homepage link to a page that has changed ownership, falls under social engineering.
Should You Trust Community Plugin Marketplaces?
Organizations should treat Claude plugin marketplaces, even ones vetted by Anthropic, just like they would treat any other third party code– as untrusted. Just because it’s in an official list somewhere, does not mean it’s secure (just look at all the malicious “verified” VSCode plugins out there).
Plugin ecosystems inherit the package registry’s oldest problems. We’re going to be relearning a lot of supply chain lessons in AI tooling over the next few years.
This was one of the cheaper ones.
Share on X (Opens in new window)
X
Share on Facebook (Opens in new window)
Facebook
Write a Comment...
Email (Required)
Name (Required)
Website
Subscribe
Subscribed
John Stawinski IV
Already have a WordPress.com account? Log in now.
