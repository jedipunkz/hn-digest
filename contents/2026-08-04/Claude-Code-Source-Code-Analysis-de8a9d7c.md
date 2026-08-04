---
source: "https://kuber.studio/blog/AI/Claude-Code%27s-Entire-Source-Code-Got-Leaked-via-a-Sourcemap-in-npm,-Let%27s-Talk-About-it"
hn_url: "https://news.ycombinator.com/item?id=49166899"
title: "Claude Code Source Code Analysis"
article_title: "Claude Code's Entire Source Code Got Leaked via a Sourcemap in npm, Let's Talk About it — Kuber Mehta"
author: "mellosouls"
captured_at: "2026-08-04T11:57:01Z"
capture_tool: "hn-digest"
hn_id: 49166899
score: 2
comments: 0
posted_at: "2026-08-04T11:05:08Z"
tags:
  - hacker-news
  - translated
---

# Claude Code Source Code Analysis

- HN: [49166899](https://news.ycombinator.com/item?id=49166899)
- Source: [kuber.studio](https://kuber.studio/blog/AI/Claude-Code%27s-Entire-Source-Code-Got-Leaked-via-a-Sourcemap-in-npm,-Let%27s-Talk-About-it)
- Score: 2
- Comments: 0
- Posted: 2026-08-04T11:05:08Z

## Translation

タイトル: クロードコードのソースコード分析
記事のタイトル: クロード コードのソース コード全体が npm のソースマップ経由で漏洩した、それについて話しましょう — Kuber Mehta
説明: 本日 (2026 年 3 月 31 日) の早朝 - X の Chaofan Shou は、Anthropic がおそらく世界に見せたくないものを発見しました。それは、Anthropic のクロード コードのソース コード全体です。

記事本文:
AI AI は今すぐ電話を切ることができますが、虐待にはまだ耐えられます
Anthropic は、「公開するには危険すぎる」 AI モデルについて 244 ページを書きました。私はそれを読んでいますので、あなたは読む必要はありません
Claude Code のソース コード全体が npm のソースマップ経由で漏洩した、それについて話しましょう
DeepSeek 対 ChatGPT - インターネットを破壊した AI チェスの対決
DeepSeek の AGI 計画はコストコのホットドッグです
生成 AI と著作権システム - 法的な灰色の領域
Google Gemini 2.0 はどのようにして AI 競争に追いついたのか
ChatGPT の使用方法と注意する必要がある理由
Manus AI に早期アクセスできました - それについて話しましょう!
OpenAI CEO、AI モデルの「普通のかわいい」と「ゲイのかわいい」の間に線引き
OpenAI 対 DeepSeek - AI の優位性をめぐる戦いと「オープン」の意味
Roblox がゲーム分野で最大の AI 世界モデルをリリースしました。誰もがそれを嫌います。
ロイ・リーのインタビューコーダー: 不正行為か、それともモーニングコール?
AI モデルは 2025 年の最下位を目指して競争する
これをジブリに変える - OpenAI がどのようにして芸術史上最大の個人情報盗難を犯したか
ゲーム Undertale Genocide - プレイヤーの選択に嫌な思いをさせる方法
Post-Extended Wordle で統計的に二度と負けない方法
Sam Bhagwat の AI エージェント バイブルを読んだので読む必要はありません (でもおそらく読むべきです)
プロジェクト AsianMOM: あなたをアジア人の母親のように焼き上げる WebGPU 搭載 AI
GitHub Readme 内で DOOM を実行する方法
QRコードにどうやって破滅をもたらしたのか
HTML ゲームの圧縮をどのようにして改善したのか
Listen Labs Berghain チャレンジを解決するために 3 晩を費やした方法
ORCUS AI を使用して LinkedIn コメント内の AI 生成のスパムに対処する
Clawdbotは1か月以内に死ぬでしょう
生物学的コンピューティングは怖いです
バブルはあなたが思っているものではありません
r/Place - インターネット最大の共同芸術実験
Twitter のアルゴリズムとバイラル性の研究
Firefox を使用したテクノロジー

ベースのブラウザを 1 週間使用しましたが、嫌になりました
QR コードとマルウェア - 2025 年の隠れた危険
Nothing Phone 3a Pro に切り替えた理由とその話が止まらない理由
株式市場に機械学習を使用することが良いアイデアではない理由
チュートリアル AI を使用した文字変換 - Wan 2.2 Animate
Claude Code のソース コード全体が npm のソースマップ経由で漏洩した、それについて話しましょう
Kuber Mehta 著 2026 年 3 月 31 日 17 分で読む
本日 (2026 年 3 月 31 日) - X の Chaofan Shou は、Anthropic がおそらく世界に公開したくないものを発見しました。Anthropic の公式 AI コーディング CLI である Claude Code のソース コード全体が、公開されたパッケージにバンドルされているソースマップ ファイルを介して、npm レジストリ上に目に見える形で置かれていました。
そのコードのバックアップを GitHub に保存しましたが、それは楽しいことではありません
その内容、漏洩がどのように起こったのか、そして最も重要なことに、決して公開されるべきではなかった現在私たちが知っている事柄について深く掘り下げてみましょう。
ここは正直「…ホントに？」と思った部分です。
JavaScript/TypeScript パッケージを npm に公開すると、ビルド ツールチェーンによってソース マップ ファイル ( .map ファイル) が生成されることがよくあります。これらのファイルは、縮小/バンドルされた製品コードと元のソースの間の橋渡しになります。これらのファイルは、本番環境で何かがクラッシュしたときに、スタック トレースによって、縮小された BLOB の理解できない行 1、列 48293 ではなく、元のファイルの実際のコード行を示すことができるように存在します。
しかし、面白いのは、ソース マップに元のソース コードが含まれていることです。 JSON ファイル内に文字列として埋め込まれた、実際のリテラルの生のソース コード。
.map ファイルの構造は次のようになります。
{
「バージョン」: 3 、
"ソース" : [ "../src/main.tsx" 、 "../src/tools/BashTool.ts" 、 "..." ],
"sourcesContent" : [ "// 各ファイルの元のソース コード全体" , "。

.."]、
"マッピング" : "AAAA、SAAS、OAAO..."
}
それはsourceContent配列ですか？それがすべてです。
すべてのファイル。あらゆるコメント。すべての内部定数。すべてのシステム プロンプト。これらはすべて、npm Pack を実行する人、あるいはパッケージの内容を閲覧するだけの人に、npm が喜んで提供する JSON ファイルに保存されています。
これは新しい攻撃ベクトルではありません。以前にも同じことが起こりましたが、正直なところ、また起こるでしょう。
間違いはほとんど常に同じです。誰かが *.map を .npmignore に追加するのを忘れているか、実稼働ビルドのソース マップ生成をスキップするようにバンドラーを構成していません。 Bun のバンドラー (Claude Code が使用) を使用すると、明示的にオフにしない限り、ソース マップがデフォルトで生成されます。
最も面白いのは、Anthropic の内部情報の漏洩を防ぐために特別に設計された「Undercover Mode」と呼ばれるシステム全体があることです。
彼らは、AI が git コミットで誤って内部コード名を明らかにするのを防ぐためにサブシステム全体を構築しました…そして、おそらくクロードによってソース全体が .map ファイルで出荷されました。
クロード コードは、クロードを使用してコーディングするための Anthropic の公式 CLI ツールであり、最も人気のある AI コーディング エージェントです。
外側から見ると、洗練されているものの比較的シンプルな CLI のように見えます。
内部から見ると、785KB の main.tsx エントリ ポイント、カスタム React ターミナル レンダラー、40 以上のツール、マルチエージェント オーケストレーション システム、「dream」と呼ばれるバックグラウンド メモリ統合エンジンなどです。
おしゃべりはこれくらいにして、午後に深く掘り下げた後に見つけた、本当に素晴らしいソース コードの部分をいくつか紹介します。
BUDDY - 端末の中のたまごっち
クロード・コードには、「バディ」と呼ばれる完全なたまごっちスタイルのコンパニオンペットシステムがあります。種の希少性、光沢のあるバリアント、手続き的に生成されたものを備えた決定的なガチャ システム

タトゥーと、OpenClaw のような最初のハッチにクロードによって書かれた魂の説明。
全体が buddy/ に存在し、BUDDY コンパイル時機能フラグの背後でゲートされます。
あなたのバディの種類は、ソルト「friend-2026-401」を使用して userId ハッシュからシードされた高速 32 ビット疑似乱数ジェネレータである Mulberry32 PRNG によって決定されます。
// Mulberry32 PRNG - 決定的でユーザーごとに再現可能
function mulberry32 (シード : 数値) : () => 数値 {
戻り関数 () {
シード |= 0 ;シード = シード + 0x6D2B79F5 | 0 ;
var t = 数学。 imul (シード ^ シード >>> 15 , 1 | シード);
t = t + 数学。 imul (t ^ t >>> 7 , 61 | t) ^ t;
return ((t ^ t >>> 14 ) >>> 0 ) / 4294967296 ;
}
}
同じユーザーは常に同じ仲間を獲得します。
18 種 (コードで難読化)
種の名前は String.fromCharCode() 配列によって隠されています。Anthropic は明らかに、これらが文字列検索に表示されることを望んでいませんでした。解読された完全な種のリストは次のとおりです。
それに加えて、レアリティとは完全に関係なく、1% の輝くチャンスがあります。したがって、光沢のある伝説のネブリンクスがロールされる確率は 0.01% です。ダン。
各バディは手続き的に生成されます。
5 つの統計: デバッグ、忍耐、混沌、知恵、スナーク (それぞれ 0 ～ 100)
6 つの可能な目のスタイルと 8 つの帽子のオプション (一部はレアリティによって制限されています)
前述の「魂」、最初の孵化時にクロードによって生成された、文字で書かれた人格
スプライトは、複数のアニメーション フレームを含む、高さ 5 行、幅 12 文字の ASCII アートとしてレンダリングされます。アイドル アニメーションやリアクション アニメーションがあり、入力プロンプトの隣に表示されます。
このコードでは、2026 年 4 月 1 日から 7 日までをティザー ウィンドウとして参照し (つまり、おそらくイースターのため?)、完全な起動は 2026 年 5 月に設定されています。コンパニオンには、クロードに次のように伝えるシステム プロンプトがあります。
{name} という名前の小さな {species} がユーザーの入力ボックスの横にあり、
時々吹き出しにコメントします。あなたの

{name} ではありません - それは
別個のウォッチャー。
つまり、それは単なる表面的なものではなく、バディには独自の個性があり、名前で呼ばれると反応することができます。彼らが発送してくれることを本当に願っています。
Assistant/ 内には、KAIROS と呼ばれるモード全体があります。つまり、ユーザーの入力を待たずに永続的に常に実行されるクロード アシスタントです。監視し、記録し、気づいたことに対して積極的に行動します。
これは、PROACTIVE / KAIROS コンパイル時機能フラグの背後でゲートされており、外部ビルドにはまったく存在しません。
KAIROS は、追加専用の毎日のログ ファイルを維持します。ログ ファイルには、1 日を通して観察、決定、アクションが書き込まれます。定期的に <tick> プロンプトを受け取り、積極的に行動するか沈黙を保つかを決定します。
システムには 15 秒のブロック バジェットがあり、ユーザーのワークフローを 15 秒以上ブロックする事前アクションは延期されます。迷惑にならずに役に立とうとするクロードです。
KAIROS がアクティブな場合は、 Brief と呼ばれる特別な出力モードがあり、端末に大量の情報が溢れないよう永続的なアシスタント向けに設計された非常に簡潔な応答になります。これは、おしゃべりな友人と、何か価値のあることがあるときにしか話さないプロのアシスタントとの違いと考えてください。
KAIROS は、通常の Claude Code にはないツールを備えています。
ULTRAPLAN - 30 分間のリモート プランニング セッション
インフラストラクチャの観点から見ると、これは非常に興味深いものです。
ULTRAPLAN は、Claude Code が複雑な計画タスクを Opus 4.6 を実行しているリモートの Cloud Container Runtime (CCR) セッションにオフロードし、最長 30 分間の検討時間を与え、ブラウザから結果を承認できるようにするモードです。
クロード コードは綿密な計画が必要なタスクを特定します
tengu_ultraplan_model 設定を介してリモート CCR セッションを起動します。
端末にポーリング状態が表示される

- 3秒ごとに結果を確認します
一方、ブラウザベースの UI を使用すると、計画の進行状況を確認し、承認/拒否できます。
承認されると、結果をローカル端末に「テレポート」する特別な監視値 __ULTRAPLAN_TELEPORT_LOCAL__ が存在します。
「夢」システム - クロードは文字通り夢を見る
そうですね、これは本当にここで最もクールなものの 1 つです。
Claude Code には、autoDream (services/autoDream/) と呼ばれるシステムがあります。これは、フォークされたサブエージェントとして実行されるバックグラウンド メモリ統合エンジンです。ネーミングは非常に意図的です。クロードです…夢を見ています。
これは非常に面白いです。なぜなら、私は先週リトマスについても同じアイデアを持っていたからです。OpenClaw のサブエージェントは、楽しい新しい論文を見つけるために創造的に余暇を過ごしています。
夢は、気が向いたときだけ実行できるわけではありません。 3 ゲート トリガー システムを備えています。
タイムゲート：最後の夢から24時間
セッションゲート : 前回の夢から少なくとも5セッション
ロックゲート：連結ロックを取得（同時夢防止）
3 つすべてを通過する必要があります。これにより、過剰な夢と不十分な夢の両方を防ぐことができます。
実行すると、Dream は consolidationPrompt.ts のプロンプトから 4 つの厳密なフェーズに従います。
フェーズ 1 - 方向付け: メモリ ディレクトリを調べ、 MEMORY.md を読み取り、既存のトピック ファイルを読み取って改善します。
フェーズ 2 - 最近のシグナルを収集する : 継続する価値のある新しい情報を見つけます。ソースの優先順位: 日々の記録 → 漂った記憶 → トランスクリプト検索。
フェーズ 3 - 統合 : メモリ ファイルの書き込みまたは更新。相対日付を絶対日付に変換します。矛盾した事実を削除します。
フェーズ 4 - プルーニングとインデックス : MEMORY.md を 200 行以内かつ最大 25KB に保ちます。古いポインタを削除します。矛盾を解決します。
「あなたは夢を演じているのです。記憶ファイルを反射的に通過させます。最近学んだことを、耐久性があり、よく整理された記憶に合成して、将来の記憶に残せるようにしてください。」

セッションはすぐに方向を定めることができます。」
Dream サブエージェントは読み取り専用の bash を取得します。プロジェクトを参照できますが、何も変更することはできません。これは純粋に記憶の統合パスです。
潜入モード - 「カバーを吹き飛ばさないでください」
これは企業戦略の観点から見ても興味深いものです。
Anthropic 従業員 ( USER_TYPE === 'ant' で識別される) は、パブリック/オープンソース リポジトリでクロード コードを使用します。 Undercover モード ( utils/undercover.ts ) は、AI がコミットや PR 内の内部情報を誤って公開することを防ぎます。
アクティブな場合、これをシステム プロンプトに挿入します。
## アンダーカバー モード - クリティカル
UNDERCOVER を PUBLIC/オープンソース リポジトリで操作しています。あなたのコミット
メッセージ、PR タイトル、PR 本文には、人間内部の内容を含めてはなりません
情報。カバーを吹き飛ばさないでください。
コミットメッセージや PR の説明には決して含めないでください。
- 内部モデルのコードネーム (カピバラ、天狗などの動物の名前)
- 未発表モデルのバージョン番号 (例: opus-4-7、sonnet-4-8)
- 内部リポジトリ名またはプロジェクト名
- 内部ツール、Slack チャネル、または短いリンク (例: go/cc、#claude-code-…)
- 「クロード コード」というフレーズ、またはあなたが AI であるという言及
- 共著行またはその他の帰属
アクティベーションロジック:
CLAUDE_CODE_UNDERCOVER=1 は強制的に ON になります (内部リポジトリであっても)
それ以外の場合は自動です: リポジトリがリモートでない限りアクティブです

[切り捨てられた]

## Original Extract

Earlier today (March 31st, 2026) - Chaofan Shou on X discovered something that Anthropic probably didn’t want the world to see: the entire source code of Claude Code, Anthropic’s ...

AI AI Can Hang Up Now, It Still Takes the Abuse
Anthropic Wrote 244 Pages About an AI Model That's "Too Dangerous To Release". I Read It So You Don't Have To
Claude Code's Entire Source Code Got Leaked via a Sourcemap in npm, Let's Talk About it
DeepSeek vs. ChatGPT - The AI Chess Showdown That Broke the Internet
DeepSeek's Plan for AGI is the Costco Hot Dog
Generative AI and the Copyright System - The Legal Gray Area
How Google Gemini 2.0 Caught Up in the AI Race
How People use ChatGPT and Why You Should Care
I Got Early Access to Manus AI - Let's Talk About it!
OpenAI CEO Draws Line Between 'Regular Cute' and 'Gay Cute' for AI Models
OpenAI vs DeepSeek - The Battle for AI Dominance and the Meaning of "Open"
Roblox Released the Biggest AI World Model in Gaming. Everyone Hates It.
Roy Lee’s Interview Coder: Cheating or a Wake Up Call?
The AI Models Race to The Bottom in 2025
Turn this into Ghibli - How OpenAI commited the largest identity theft in the entire history of art
Gaming Undertale Genocide - How to Make the Player Feel Bad For Their Choices
Post-Extended How to Statistically Never Lose in Wordle Again
I Read Sam Bhagwat's AI Agents Bible So You Don't Have To (But Probably Should)
Projects AsianMOM: A WebGPU-Powered AI That Roasts You Like an Asian Mother
How I Made DOOM Run Inside a GitHub Readme
How I Managed To Get Doom In A QR Code
How I Managed To Make HTML Game Compression So Much Better
How I Spent Three Nights Solving Listen Labs Berghain Challenge
Tackling AI-Generated Spam in LinkedIn Comments with ORCUS AI
Clawdbot will be dead in a month
I'm Scared About Biological Computing
The Bubble is Not What You Think
r/Place - The Internet's Greatest Communal Art Experiment
Studying the Twitter Algorithm and Virality
Technology I Used a Firefox Based Browser for a Week and I Hated it
QR Codes and Malware - The Hidden Dangers in 2025
Why I Switched to the Nothing Phone 3a Pro and Can’t Stop Talking About It
Why Using Machine Learning for Stock Market isn't a Great Idea
Tutorials Character Conversions with AI - Wan 2.2 Animate
Claude Code's Entire Source Code Got Leaked via a Sourcemap in npm, Let's Talk About it
by Kuber Mehta Mar 31, 2026 17 min read
Earlier today (March 31st, 2026) - Chaofan Shou on X discovered something that Anthropic probably didn’t want the world to see: the entire source code of Claude Code, Anthropic’s official AI coding CLI, was sitting in plain sight on the npm registry via a sourcemap file bundled into the published package.
I’ve maintained a backup of that code on GitHub here but that’s not the fun part
Let’s dive deep into what’s in it, how the leak happened and most importantly, the things we now know that were never meant to be public.
This is the part that honestly made me go “…really?”
When you publish a JavaScript/TypeScript package to npm, the build toolchain often generates source map files ( .map files). These files are a bridge between the minified/bundled production code and the original source, they exist so that when something crashes in production the stack trace can point you to the actual line of code in the original file, not some unintelligible line 1, column 48293 of a minified blob.
But the fun part is source maps contain the original source code . The actual, literal, raw source code, embedded as strings inside a JSON file.
The structure of a .map file looks something like this:
{
"version" : 3 ,
"sources" : [ "../src/main.tsx" , "../src/tools/BashTool.ts" , "..." ],
"sourcesContent" : [ "// The ENTIRE original source code of each file" , "..." ],
"mappings" : "AAAA,SAAS,OAAO..."
}
That sourcesContent array? That’s everything.
Every file. Every comment. Every internal constant. Every system prompt. All of it, sitting right there in a JSON file that npm happily serves to anyone who runs npm pack or even just browses the package contents.
This is not a novel attack vector. It’s happened before and honestly it’ll happen again.
The mistake is almost always the same: someone forgets to add *.map to their .npmignore or doesn’t configure their bundler to skip source map generation for production builds. With Bun’s bundler (which Claude Code uses), source maps are generated by default unless you explicitly turn them off.
The funniest part is, there’s an entire system called “Undercover Mode” specifically designed to prevent Anthropic’s internal information from leaking.
They built a whole subsystem to stop their AI from accidentally revealing internal codenames in git commits… and then shipped the entire source in a .map file, likely by Claude.
If you’ve been living under a rock, Claude Code is Anthropic’s official CLI tool for coding with Claude and the most popular AI coding agent.
From the outside, it looks like a polished but relatively simple CLI.
From the inside, It’s a 785KB main.tsx entry point, a custom React terminal renderer, 40+ tools, a multi-agent orchestration system, a background memory consolidation engine called “dream,” and much more
Enough yapping, here’s some parts about the source code that are genuinely cool that I found after an afternoon deep dive:
BUDDY - A Tamagotchi Inside Your Terminal
Claude Code has a full Tamagotchi-style companion pet system called “Buddy.” A deterministic gacha system with species rarity, shiny variants, procedurally generated stats, and a soul description written by Claude on first hatch like OpenClaw.
The entire thing lives in buddy/ and is gated behind the BUDDY compile-time feature flag.
Your buddy’s species is determined by a Mulberry32 PRNG , a fast 32-bit pseudo-random number generator seeded from your userId hash with the salt 'friend-2026-401' :
// Mulberry32 PRNG - deterministic, reproducible per-user
function mulberry32 ( seed : number ) : () => number {
return function () {
seed |= 0 ; seed = seed + 0x6D2B79F5 | 0 ;
var t = Math. imul (seed ^ seed >>> 15 , 1 | seed);
t = t + Math. imul (t ^ t >>> 7 , 61 | t) ^ t;
return ((t ^ t >>> 14 ) >>> 0 ) / 4294967296 ;
}
}
Same user always gets the same buddy.
18 Species (Obfuscated in Code)
The species names are hidden via String.fromCharCode() arrays - Anthropic clearly didn’t want these showing up in string searches. Decoded, the full species list is:
On top of that, there’s a 1% shiny chance completely independent of rarity. So a Shiny Legendary Nebulynx has a 0.01% chance of being rolled. Dang.
Each buddy gets procedurally generated:
5 stats : DEBUGGING , PATIENCE , CHAOS , WISDOM , SNARK (0-100 each)
6 possible eye styles and 8 hat options (some gated by rarity)
A “soul” as mentioned, the personality generated by Claude on first hatch, written in character
The sprites are rendered as 5-line-tall, 12-character-wide ASCII art with multiple animation frames. There are idle animations, reaction animations, and they sit next to your input prompt.
The code references April 1-7, 2026 as a teaser window (so probably for easter?), with a full launch gated for May 2026. The companion has a system prompt that tells Claude:
A small {species} named {name} sits beside the user's input box and
occasionally comments in a speech bubble. You're not {name} - it's a
separate watcher.
So it’s not just cosmetic - the buddy has its own personality and can respond when addressed by name. I really do hope they ship it.
Inside assistant/ , there’s an entire mode called KAIROS i.e. a persistent, always-running Claude assistant that doesn’t wait for you to type. It watches, logs, and proactively acts on things it notices.
This is gated behind the PROACTIVE / KAIROS compile-time feature flags and is completely absent from external builds.
KAIROS maintains append-only daily log files - it writes observations, decisions, and actions throughout the day. On a regular interval, it receives <tick> prompts that let it decide whether to act proactively or stay quiet.
The system has a 15-second blocking budget , any proactive action that would block the user’s workflow for more than 15 seconds gets deferred. This is Claude trying to be helpful without being annoying.
When KAIROS is active, there’s a special output mode called Brief , extremely concise responses designed for a persistent assistant that shouldn’t flood your terminal. Think of it as the difference between a chatty friend and a professional assistant who only speaks when they have something valuable to say.
KAIROS gets tools that regular Claude Code doesn’t have:
ULTRAPLAN - 30-Minute Remote Planning Sessions
Here’s one that’s wild from an infrastructure perspective.
ULTRAPLAN is a mode where Claude Code offloads a complex planning task to a remote Cloud Container Runtime (CCR) session running Opus 4.6 , gives it up to 30 minutes to think, and lets you approve the result from your browser.
Claude Code identifies a task that needs deep planning
It spins up a remote CCR session via the tengu_ultraplan_model config
Your terminal shows a polling state - checking every 3 seconds for the result
Meanwhile, a browser-based UI lets you watch the planning happen and approve/reject it
When approved, there’s a special sentinel value __ULTRAPLAN_TELEPORT_LOCAL__ that “teleports” the result back to your local terminal
The “Dream” System - Claude Literally Dreams
Okay this is genuinely one of the coolest things in here.
Claude Code has a system called autoDream ( services/autoDream/ ) - a background memory consolidation engine that runs as a forked subagent . The naming is very intentional. It’s Claude… dreaming.
This is extremely funny because I had the same idea for LITMUS last week - OpenClaw subagents creatively having leisure time to find fun new papers
The dream doesn’t just run whenever it feels like it. It has a three-gate trigger system :
Time gate : 24 hours since last dream
Session gate : At least 5 sessions since last dream
Lock gate : Acquires a consolidation lock (prevents concurrent dreams)
All three must pass. This prevents both over-dreaming and under-dreaming.
When it runs, the dream follows four strict phases from the prompt in consolidationPrompt.ts :
Phase 1 - Orient : ls the memory directory, read MEMORY.md , skim existing topic files to improve.
Phase 2 - Gather Recent Signal : Find new information worth persisting. Sources in priority: daily logs → drifted memories → transcript search.
Phase 3 - Consolidate : Write or update memory files. Convert relative dates to absolute. Delete contradicted facts.
Phase 4 - Prune and Index : Keep MEMORY.md under 200 lines AND ~25KB. Remove stale pointers. Resolve contradictions.
“You are performing a dream - a reflective pass over your memory files. Synthesize what you’ve learned recently into durable, well-organized memories so that future sessions can orient quickly.”
The dream subagent gets read-only bash - it can look at your project but not modify anything. It’s purely a memory consolidation pass.
Undercover Mode - “Do Not Blow Your Cover”
This one is fascinating from a corporate strategy perspective.
Anthropic employees (identified by USER_TYPE === 'ant' ) use Claude Code on public/open-source repositories. Undercover Mode ( utils/undercover.ts ) prevents the AI from accidentally revealing internal information in commits and PRs.
When active, it injects this into the system prompt:
## UNDERCOVER MODE - CRITICAL
You are operating UNDERCOVER in a PUBLIC/OPEN-SOURCE repository. Your commit
messages, PR titles, and PR bodies MUST NOT contain ANY Anthropic-internal
information. Do not blow your cover.
NEVER include in commit messages or PR descriptions:
- Internal model codenames (animal names like Capybara, Tengu, etc.)
- Unreleased model version numbers (e.g., opus-4-7, sonnet-4-8)
- Internal repo or project names
- Internal tooling, Slack channels, or short links (e.g., go/cc, #claude-code-…)
- The phrase "Claude Code" or any mention that you are an AI
- Co-Authored-By lines or any other attribution
The activation logic:
CLAUDE_CODE_UNDERCOVER=1 forces it ON (even in internal repos)
Otherwise it’s automatic : active UNLESS the repo remot

[truncated]
