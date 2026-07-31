---
source: "https://nanocollective.org/blog/prompt-scrub-v100-a-local-first-scrubber-for-prompts-and-their-responses-76"
hn_url: "https://news.ycombinator.com/item?id=49124405"
title: "Show HN: Prompt-scrub – local-first PII redaction for LLM prompts and responses"
article_title: ""
author: "mrspence"
captured_at: "2026-07-31T15:54:57Z"
capture_tool: "hn-digest"
hn_id: 49124405
score: 2
comments: 1
posted_at: "2026-07-31T15:32:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Prompt-scrub – local-first PII redaction for LLM prompts and responses

- HN: [49124405](https://news.ycombinator.com/item?id=49124405)
- Source: [nanocollective.org](https://nanocollective.org/blog/prompt-scrub-v100-a-local-first-scrubber-for-prompts-and-their-responses-76)
- Score: 2
- Comments: 1
- Posted: 2026-07-31T15:32:39Z

## Translation

タイトル: HN の表示: プロンプト スクラブ – LLM プロンプトと応答に対するローカル ファーストの PII 編集
説明: [Nano Collective](https://nanocollective.org) によって構築されました。これは、営利を目的とするのではなく、コミュニティのために AI ツールを構築するコミュニティ集団です。
これは、「prompt-scrub」の最初の公開リリースです。

記事本文:
Nano Collective ビルド ブログ ドキュメント コントリビューター スポンサー ビルド ブログ ドキュメント コントリビューター スポンサー < ブログに戻る [ パッケージ ] [ 新しいコンセプト ] [ リリース ] Prompt-scrub v1.0.0: プロンプトとその応答のローカルファースト スクラバー
2026 年 7 月 15 日 | @LottieOxford | 0 コメント Nano Collective によって構築されました。Nano Collective は、営利目的ではなくコミュニティのために AI ツールを構築するコミュニティ集団です。
これは、マシン上で完全に実行される小規模な Node.js ユーティリティである Prompt-scrub の最初の公開リリースです。プロンプト内の識別コンテンツ (電子メール、パス、シークレット、電話番号、URL、住所、およびいくつかのオプトイン カテゴリ) を検出し、各検出結果を Email_1 や Path_2 などの安定したプレースホルダーに置き換え、モデルの応答が返された後にローカルで元の値にリハイドレートできるようにします。
動機は単純です。クラウド LLM への偶発的な識別子の漏洩のほとんどは、プロンプトのテキストとその応答のテキストに存在します。プロンプトがマシンを離れる前に、そこでそれを決定的に削除することは、プライバシー体制において有用なレイヤーであり、ネットワークのラウンドトリップ、新しいアカウント、またはホストされたサービスの動作を必要としないレイヤーです。
このパッケージは 2 つの関数と CLI を公開します。
scrub() は、プレーン文字列または { role, content } メッセージの配列を受け取り、テキストに対して構成された検出器を実行し、各検出結果をカテゴリ名前空間のプレースホルダーに置き換えて、スクラブされたコンテンツとセッション ID を返します。
import { スクラブ、リハイドレート } から '@nanocollective/prompt-scrub';
const プロンプト = "私のキーは sk-12345 で、電子メールは [email protected] ";
const { scrubbedContent, sessionId } = scrub({ content: プロンプト });
// scrubbedContent: "私のキーは Secret_1 で、電子メールは Email_1 です。"
すでに使用している LLM プロバイダーに scrubedContent を送信します。いつ

応答が返され、 rewind() がテキストを調べ、セッション マップ内の各プレースホルダーを検索して、それらを元に戻します。不明なプレースホルダー (モデルが幻覚を起こしたプレースホルダー テキスト、または前のセッションのプレースホルダー) は変更されずに渡され、警告として表示されるため、それらを信頼するかどうかを決定できます。
const response = "あなたのメール Email_1 は正しいようで、キー Secret_1 も問題ありません。";
const { コンテンツ, 警告 } = re水和物({ コンテンツ: 応答, セッション ID });
// content: "あなたの電子メール [email protected] は正しいようで、キー sk-12345 は問題ありません。"
セッション ID は 2 つのステップ間のリンクです。これは、プレースホルダーと元のマッピングを保持する、OS config ディレクトリの下にある小さな JSON ファイルを指します。ファイルは制限されたアクセス許可でアトミックに書き込まれ、破損ファイルの隔離パスが含まれているため、書きかけのマップによってリハイドレーションが自動的に無効になることはありません。この場所は、PROMPT_SCRUB_CONFIG_DIR 環境変数を介して上書きできます。
8 個の検出器が同梱されています。
パス (Unix スタイルおよび Windows スタイル)
Secret (共通の API キーと資格情報の形状に合わせて調整されています。資格情報が欠落していることは、名前が欠落していることよりも悪いためです)
名前 (より厳密な許可リスト モードを使用した固有名詞検出器)
code-tell (内部プロジェクトのコード名など、ユーザーが列挙したプライベート識別子)
URL 検出では、サブドメインが一致する信頼できるホストのホワイトリストも受け入れられるため、必要に応じて、制御する内部サービスをプレースホルダーなしで通過させることができます。重複する検出結果 (URL フラグメントのように見える電子メールなど) は、文書化された優先順位に従って解決され、同点の場合はスパンが長い方が優先されます。決定論は意図的です。同じ入力とセッションは常に同じスクラブされた出力を生成し、プロバイダーのプロンプト キャッシュ プレフィックスをバイト安定に保ちます。
小さな命令――

同じロジックの行ラッパー。推奨されるワークフローは、最初に検査し、次にスクラブします。
# 何も書かずに何が変わるかを正確に確認する
echo "私の電子メールは [電子メールで保護されています] " |プロンプトスクラブ検査
# stdin (またはファイル) をスクラブし、セッション ID を stderr に出力します
echo "私の電子メールは [電子メールで保護されています] " |プロンプトスクラブ スクラブ
# --session-id を使用して応答をリハイドレートする
echo "詳細については、Email_1 にお問い合わせください。" |プロンプト スクラブ リハイドレート --session-id <id>
# セッションマップを検査する
プロンプトスクラブセッションリスト
プロンプト スクラブ セッションで <id> が表示される
# ルールパックの追加を含む、アクティブな検出器セットを表示します
プロンプト スクラブ ルールのリスト
検査は実際に使ってもらいたい部分です。セッション ファイルは書き込まれず、スクラブされた出力の SHA-256 ハッシュが出力されるため、実行全体でバイト安定キャッシュ プレフィックスを検証できます。
検出されたエンティティ:
[電子メール] [電子メール保護] → 電子メール_1 (文字 12 ～ 29)
セッションが書き込まれませんでした。
ハッシュ: 41beda4af0b83488fdf6eea9347775450a1c7c887a6ef377212340f36c445132
拡張性
カスタム ディテクタは、ライブラリ API ( ScrubOptions の CustomDetectors ) 経由で渡すことができます。それぞれが組み込みと同じ形式で一致を返すため、同じ優先順位/スパン ロジックに組み込まれます。
ルール パックは、追加の検出機能を提供する個別の npm パッケージです。これらは config または package.json で宣言され、アクティブ セットにマージされ、 rules list に表示されます。これは、パッケージをフォークせずにプロジェクト間でディテクタを共有するためのパスです。
これは何で、何がそうではないのか
プロンプトスクラブは、コンテンツ層での ID 漏洩を軽減します。これは匿名性ではなく部分防御であり、区別は重要であるため、README と脅威モデルでは区別について明示されています。
プロンプトでの偶発的な秘密漏洩。秘密検出器は共通 AP 上で高精度になるように調整されています

I キー、トークン、および資格情報の形状。
1 回限りのプロンプトでの偶発的な識別子の漏洩。電子メール、電話番号、住所、パス、および URL は保守的なデフォルトで捕捉され、元の値はセッション マップの下のディスクに残ります。
クラウド LLM プロバイダーが識別コンテンツを読み取ります。スクラブ後、プロンプトにはアドレスの代わりに Email_1 が表示されます。これは、識別性が大幅に低下しますが、ゼロではありません。また、特定のプロバイダーは、独自のログにアクセスできる場合、セッション全体で相関関係を維持できます。
プロンプトコンテンツから長期的なプロフィールを構築します。安定したセッション マッピングにより、単一セッション内での識別子レベルの相関関係が防止されます。文体的な特徴には触れておらず、物事の表現方法は変更されません。
ツール呼び出しの結果、エージェント設定が行われます。スクラバーは、発信元に関係なくすべてのメッセージに対して実行されるため、 ls 、 git log 、 cat 、および grep の出力は、次の LLM ターンの前にスクラブされます。対象範囲は、構成された検出器に限定されます。
侵害されたローカル マシン。スクラバーはローカルで実行されます。環境が危険にさらされている場合、プロンプトも危険にさらされます。ディスク上のセッション マップは、v1 ではプレーンテキストの JSON であり、あなたのアカウントを持つ攻撃者がそれらを読み取ることができます。保存時の暗号化は v1.1 のフォローアップです。
セマンティック漏洩。本質的に個人を特定できる質問 (個人的なコードベース、あなただけが持つニッチなバグ、会計士だけが知っている番号) は、識別子を削除しても匿名にすることはできません。
文体的な指紋採取。 v1 のブルート フォース アプローチはスタイルを書き換えません。
ネットワークまたはキー層。 IP アドレス、リクエストのタイミング、ヘッダーはスクラバーの範囲外です。必要に応じて、独自に選択したネットワーク ツールと組み合わせてください。
このツールを使用すると匿名化できると信じているユーザーは、プロンプトを読まなくなってしまうため、まったく使用しなかったユーザーよりも悪い状況にあります。

デフォルトを使用してください。常に最初に検査を使用してください。脅威モデルの文書には、全体像が 1 か所で詳しく説明されています。
# グローバル CLI として
npm install -g @nanocollective/prompt-scrub
# Node.js の依存関係として
npm install @nanocollective/prompt-scrub
パッケージはプロジェクトのライセンスに基づいてオープンソースです。ソース、完全なドキュメント、および脅威モデルはリポジトリにあります。
https://github.com/Nano-Collective/prompt-scrubber
これが初の一般公開です。私たちが最も聞きたいこと:
検出器の範囲。私たちが欠けている一般的な図形を貼り付けますか?ワークフローに対してデフォルトでどのオプトイン カテゴリ (名前、コードテル) をオンにする必要がありますか?
ルールパックのパッケージ。拡張機能の形状は決まっていますが、実際のルール パックはまだ確認されておらず、パッケージが 0.x である間に形式は変更される可能性があります。
脅威モデル。私たちは意図的に部分防御のフレームを選択しました。何が保護され、何が保護されていないのかがドキュメントで不明瞭な場合は、その上にコードを追加して出荷する前にそれを知りたいと考えます。
UX を検査します。ハッシュ出力と予行演習の差分は、人々が実際に使用することを目的としています。実際にぎこちない場合は、すぐに修正できます。
問題や PR はリポジトリで歓迎されます。 Nano Collective Discord もあり、この集団が何を構築しているのか、そしてその理由についてより広範な会話が可能です。
この投稿についてご意見がありますか? GitHub ディスカッションにアクセスしてフィードバックを共有してください。
Nano Collective 開発者向けに構築されたオープンソース AI ツール。プライバシー第一、ローカル第一。
© 2026 ナノコレクティブ。無断転載を禁じます。

## Original Extract

Built by the [Nano Collective](https://nanocollective.org), a community collective building AI tooling not for profit, but for the community.
This is the first public release of `prompt-scrub`, a s...

Nano Collective Build Blog Docs Contributors Sponsor Build Blog Docs Contributors Sponsor < Back to Blogs [ Package ] [ New Concept ] [ Released ] prompt-scrub v1.0.0: a local-first scrubber for prompts and their responses
July 15, 2026 | @ LottieOxford | 0 comments Built by the Nano Collective , a community collective building AI tooling not for profit, but for the community.
This is the first public release of prompt-scrub , a small Node.js utility that runs entirely on your machine. It detects identifying content inside a prompt (emails, paths, secrets, phone numbers, URLs, postal addresses, and a couple of opt-in categories), replaces each finding with a stable placeholder like Email_1 or Path_2 , and lets you rehydrate the model's response back to the original values locally after it comes back.
The motivation is simple: most accidental identifier leakage to a cloud LLM lives in the text of the prompt and the text of its response. Stripping it there, deterministically, before the prompt leaves your machine is a useful layer in a privacy posture, and one that does not need a network round-trip, a new account, or a hosted service to work.
The package exposes two functions and a CLI.
scrub() takes either a plain string or an array of { role, content } messages, runs the configured detectors over the text, replaces each finding with a category-namespaced placeholder, and returns the scrubbed content plus a session id:
import { scrub, rehydrate } from '@nanocollective/prompt-scrub';
const prompt = "My key is sk-12345 and my email is [email protected] ";
const { scrubbedContent, sessionId } = scrub({ content: prompt });
// scrubbedContent: "My key is Secret_1 and my email is Email_1"
You send scrubbedContent to whatever LLM provider you already use. When the response comes back, rehydrate() walks the text, looks up each placeholder in the session map, and swaps them back. Unknown placeholders (placeholder text the model hallucinated, or placeholders from a previous session) are passed through unchanged and surfaced as warnings so you can decide whether to trust them:
const response = "Your email Email_1 looks correct and your key Secret_1 is fine.";
const { content, warnings } = rehydrate({ content: response, sessionId });
// content: "Your email [email protected] looks correct and your key sk-12345 is fine."
The session id is the link between the two steps. It points at a small JSON file under your OS config directory that holds the placeholder-to-original mapping. The file is written atomically, with restrictive permissions, and includes a corrupt-file quarantine path so a half-written map does not silently disable rehydration. The location is overridable via the PROMPT_SCRUB_CONFIG_DIR environment variable.
Eight detectors ship in the box.
path (Unix-style and Windows-style)
secret (tuned for common API key and credential shapes, since missing a credential is worse than missing a name)
name (proper-noun detector, with a stricter allowlist mode)
code-tell (user-enumerated private identifiers, for things like internal project codenames)
URL detection also accepts a trusted-host allowlist with subdomain matching, so internal services you control can pass through without a placeholder if that is what you want. Overlapping detector findings (an email that looks like a URL fragment, say) resolve by a documented priority order, with longer span winning on ties. Determinism is deliberate: the same input and session always produce the same scrubbed output, which keeps provider prompt-cache prefixes byte-stable.
A small command-line wrapper around the same logic. The recommended workflow is inspect first, then scrub :
# See exactly what would change without writing anything
echo "My email is [email protected] " | prompt-scrub inspect
# Scrub stdin (or a file), prints the session id to stderr
echo "My email is [email protected] " | prompt-scrub scrub
# Rehydrate a response using --session-id
echo "Contact Email_1 for details." | prompt-scrub rehydrate --session-id <id>
# Inspect a session map
prompt-scrub sessions list
prompt-scrub sessions show <id>
# See the active detector set, including any rule-pack additions
prompt-scrub rules list
inspect is the part we want people to actually use. It does not write a session file and it prints a SHA-256 hash of the scrubbed output so you can verify byte-stable cache prefixes across runs:
Detected entities:
[Email] [email protected] → Email_1 (chars 12-29)
No session written.
Hash: 41beda4af0b83488fdf6eea9347775450a1c7c887a6ef377212340f36c445132
Extensibility
Custom detectors can be passed in via the library API ( customDetectors in ScrubOptions ). Each one returns matches in the same shape the built-ins do, so they slot into the same priority / span logic.
Rule packs are separate npm packages that contribute additional detectors. They are declared in your config or in package.json , merged into the active set, and visible in rules list . This is the path for sharing detectors across projects without forking the package.
What this is, and what it is not
prompt-scrub reduces identity leakage at the content layer. It is partial defence, not anonymity, and the README and the Threat Model are explicit about the distinction because the distinction matters.
Accidental secret leakage in prompts. The secret detector is tuned for high precision on the common API key, token, and credential shapes.
Accidental identifier leakage in one-off prompts. Emails, phone numbers, postal addresses, paths, and URLs are caught with conservative defaults, and the original values stay on disk under the session map.
Cloud LLM providers reading identifying content. After scrubbing, the prompt contains Email_1 instead of your address. That is materially less identifying, but not zero, and a determined provider can still correlate across a session if it has access to its own logs.
Long-term profile building from prompt content. Stable session mappings prevent identifier-level correlation within a single session. They do not address stylistic fingerprinting, the way you phrase things goes out unchanged.
Tool call results in agentic settings. The scrubber runs on every message regardless of origin, so ls , git log , cat , and grep outputs are scrubbed before the next LLM turn. Coverage is limited to the configured detectors.
A compromised local machine. The scrubber runs locally. If your environment is compromised, your prompt is too. The session maps on disk are plaintext JSON in v1, and an attacker with your account can read them. Encryption at rest is a v1.1 follow-up.
Semantic leakage. A question that is inherently identifying (your private codebase, a niche bug only you have, a number only your accountant knows) cannot be made anonymous by stripping identifiers.
Stylistic fingerprinting. The v1 brute-force approach does not rewrite style.
The network or key layer. Your IP address, request timing, and headers are outside the scrubber's scope. Pair it with a network tool of your own choosing if you need that.
A user who believes this tool makes them anonymous is worse off than one who never used it, because they stop reading their prompts and trust the defaults. Always use inspect first. The Threat Model document spells out the full picture in one place.
# As a global CLI
npm install -g @nanocollective/prompt-scrub
# As a Node.js dependency
npm install @nanocollective/prompt-scrub
The package is open source under the project's licence. Source, full docs, and the Threat Model live at the repo:
https://github.com/Nano-Collective/prompt-scrubber
This is the first public release. The things we would most like to hear about:
Detector coverage. Which common shapes do you paste in that we are missing? Which opt-in category (name, code-tell) should be on by default for your workflow?
Rule pack packaging. The extension shape is in, but we have not yet seen real-world rule packs, and the format is open to change while the package is at 0.x.
The Threat Model. We chose partial-defence framing deliberately. If the documentation is unclear about what is and is not defended, we want to know before we ship more code on top of it.
The inspect UX. The hash output and the dry-run diff are intended to be the thing people actually use. If they are awkward in practice, that is a fast fix.
Issues and PRs are welcome at the repo. There is also a Nano Collective Discord for the wider conversation about what the collective is building and why.
Have thoughts on this post? Head over to our GitHub discussions to share your feedback.
Nano Collective Open-source AI tools built for developers. Privacy-first, local-first.
© 2026 Nano Collective. All rights reserved.
