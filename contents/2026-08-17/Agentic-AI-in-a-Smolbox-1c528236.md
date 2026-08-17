---
source: "https://remyhax.xyz/posts/smolbox/"
hn_url: "https://news.ycombinator.com/item?id=49338338"
title: "Agentic AI in a Smolbox"
article_title: "Agentic AI in a Smolbox |\nREMY HAX"
image: "https://remyhax.xyz/image/feature-agentic-ai-in-a-smolbox.png"
author: "jerrythegerbil"
captured_at: "2026-08-17T22:14:12Z"
capture_tool: "hn-digest"
hn_id: 49338338
score: 2
comments: 0
posted_at: "2026-08-17T22:07:38Z"
tags:
  - hacker-news
  - translated
---

# Agentic AI in a Smolbox

- HN: [49338338](https://news.ycombinator.com/item?id=49338338)
- Source: [remyhax.xyz](https://remyhax.xyz/posts/smolbox/)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T22:07:38Z

## Translation

タイトル: Smolbox のエージェント AI
記事のタイトル: Smolbox 内のエージェント AI |
レミー・ハックス
説明: Smolbox は完全にブラウザ タブ内で実行されます。これは、WebAssembly (WASM) で実行される Alpine Linux ベースの完全な x86_64 仮想マシン (VM) サンドボックスです。また、WebGPU を介して実行される LLM モデルの形式の「AI」人工知能も備えており、VM 内のツール呼び出しとして公開されます。オプションでできます
[切り捨てられた]

記事本文:
Smolbox は完全にブラウザ タブ内で実行されます。これは、WebAssembly (WASM) で実行される Alpine Linux ベースの完全な x86_64 仮想マシン (VM) サンドボックスです。また、WebGPU を介して実行される LLM モデルの形式の「AI」人工知能も備えており、VM 内のツール呼び出しとして公開されます。必要に応じて、Web ブラウザーのファイル ピッカーを使用して、VM に追加するファイルを含むフォルダーをマウントできます。これらのフォルダーは、 /mnt/host の下に読み取り専用でマウントされます。
結合された結果は、サーバー側の処理がないことを除けば、Claude Code または Codex の結果と非常によく似ています。 Linux VM サンドボックス、LLM、ツール呼び出し、およびファイル アクセス制御は、ユーザーのブラウザ タブにのみ存在します。
このブログでは、最終結果について、あたかも私が自分の手柄を主張できるかのように詩的に語るつもりはありません。なぜなら、私にはそれができないし、そうすべきではないからです。私は自分が何を望んでいるのか、それをどのように行うのか、何を使用するのか、そして成功を検証する方法を正確に知っていました。 AI は、曖昧さゼロで、私が尋ねたとおりに、smolbox 用にすべての要素をまとめてくれました。
Smolbox は、「これはいつでも可能だった」ことを指摘し、AI セキュリティの現状が次世代のユーザーに与えている不利益について語るための概念実証です。
私は30代前半です。昨日はゴミ出しで腰を抜かしたので、座ってこのブログを書いています。私はクレジット カードを持っており、コストを気にすることなく 24 時間年中無休でエージェント ワークフローを実行する 2 つの異なる AI サブスクリプションの料金を支払っています。私はコンピューターをたくさん持っているので、何ヶ月も気づかずに紛失してしまいます。
しかし、2000 年代のいつか、私は携帯電話 (LG Dare) を持っていて、10 代でした。私は無制限のテキストと写真 (SMS/MMS) メッセージを持っていましたが、「データ」は依然として法外に高価なものだったので、携帯電話で誤って間違ったボタンをタップすると、月々の請求書に表示されてしまいます。スパムが届きました

ある日、電子メール アドレスから届いたテキスト メッセージ。 my-ten-digit-cell-number@vzwpix.com にメッセージを送信して少し遊んだ後、Verizon が管理する MMS<->電子メール ブリッジがあることがわかりました。そこで私は、件名を URL として電子メールを送信し、Web ページの内容を含む電子メールを受信できる電子メール サービスを運営している進取的な PHP 開発者をオンラインで見つけることができました。彼らに丁寧に尋ねた後、彼らは私の電子メールアドレスを確認し、無制限のテキストメッセージを通じてインターネットを「閲覧」することができました。関連コストの増加なしで情報に無料でアクセスできるため、携帯電話自体の開発者マニュアルと SDK を読むようになりました (小さなテキスト メッセージが期待されるコンテンツの Web ページ全体を処理しようとしたときに、携帯電話が一度に数分間ロックアップしてフリーズしました)。最終的に、着信音が MicroSD カード スロット経由で読み込まれるときにパス トラバーサルの脆弱性があることに気付き、MacroMedia Flash Lite で書かれた組み込みの世界時計「アプリ」をミニゴルフ ゲームで上書きできることに気づきました。
私は何度も連絡を取ったあのランダムな開発者のことを思い出します。サービスは遅かったです。ウェブページのフォーマットは最悪でした。そのせいで携帯電話がフリーズしてしまいました。それでも、それは私がこれまで望んでいたすべてでした。なぜなら、それは何かを可能にしてくれたからです。
おそらく彼らはそれについてあまり考えていないか、自分たちが構築した小さなサービスを誰かが使ってくれているのを見て喜んでいただけだったのでしょうが、彼らは水門を開け、最終的に私のキャリアパス全体を形作ることになる道を私に示してくれました。彼らがやったことは、何かを可能にすることだけでした。
しかし、十代の若者は十分にエージェント的でしょうか?
もし私が今 10 代だったら、インターネットにアクセスする方法を理解していたことは間違いありません。子供たちはそのように機知に富んでいます。学校支給のiPadでも、近所のコンピュータでも

図書館、または不要になった中古のおさがりノートパソコン。それは簡単です。
しかし、ロックダウンされた学校支給の iPad にクロード コードをインストールできるでしょうか?図書館の共用コンピュータに任意のプログラムをインストールできるのでしょうか？ AI サブスクリプションを購入して、これらの仕組みを学ぶことはできますか?彼らはそのサブスクリプションを購入するためのクレジット カードを持っていますか?
はい、はい、できます！ロックダウンされたシステムのガードレールを回避するための抜け穴を知っている人は常に存在しますが、アクセスが制限されたり、捕まった場合に問題が発生したりするという注意が伴います。通貨問題も同様だ。物々交換できるカードを持っている少し年上の友人が必ずいます。あるいは、取引可能な資産やゲーム内通貨を備えた主要なオンライン ゲームをランダムに選択すると、実際の有形通貨がすぐに使用できる形でシステムから出てくる、マネーロンダリングや詐欺行為と全く異なるものではない方法で運営されている経済がすぐに見つかります。そして彼らはおそらく、ランダムな Telegram メッセージング チャネルで、夜間のベンダーを通じて安価な「未使用」 (ハッキングされたアカウント) のトークン容量を購入することで、その資金を可能な限り拡張するでしょう。
もし私が今 10 代だったら、ほぼ間違いなくエージェント AI で遊びたいと思うでしょう。結局のところ、それは未来ですよね？もし私が今 10 代だったら、ほぼ確実にそれを可能にする道を選ぶでしょう。
コンピュータ セキュリティにおいて、サンドボックスは実行中のプログラムを分離するためのセキュリティ メカニズムであり、通常はシステム障害やソフトウェアの脆弱性の拡散を軽減する目的で使用されます。サンドボックスの比喩は、子供用サンドボックスの概念に由来しています。これは、子供たちが現実世界に損害を与えることなく、建設、破壊、実験ができる遊び場です。 - https://en.wikipedia.org

/wiki/サンドボックス_(コンピューター_セキュリティ)
Smolbox はサンドボックスです。小さいもの。現実世界に損害を与えることのない安全な遊び場であり、特に特別な許可は必要ありません。学校支給の iPad、図書館のコンピュータ、お下がりの中古ラップトップで無料で動作します。サンドボックスの機能を拡張するのは簡単です。箱にもっとたくさんの物を入れて遊ぶだけです。
スモルボックスは遅いです。それは時々厄介です。確かに、プライベートであることとデータをまったく収集しないこと以外は特に優れているわけではありません。
しかし、これにより、大多数の人が安全に作業を行うことが可能になり、無作為の開発者にとっても簡単に行うことができます。それが誰かにとって何を意味するのか、私は正確に知っています。
それは可能です。あとは誰が能力と意欲を兼ね備えているかという問題だけですが、今は腰が痛くて少し不機嫌です。

## Original Extract

Smolbox runs entirely in a browser tab. It is a full x86_64 virtual machine (VM) sandbox based on Alpine linux that runs under WebAssembly (WASM). It also has “AI” artificial intelligence in the form of an LLM model that runs via WebGPU which is exposed as tool calls inside the VM. You can optionall
[truncated]

Smolbox runs entirely in a browser tab. It is a full x86_64 virtual machine (VM) sandbox based on Alpine linux that runs under WebAssembly (WASM) . It also has “AI” artificial intelligence in the form of an LLM model that runs via WebGPU which is exposed as tool calls inside the VM. You can optionally use the web browser file-picker to mount a folder containing files you want added to the VM, which are mounted read-only under /mnt/host .
The combined result being very similar to that of Claude Code or Codex, except that there’s no server side processing. The linux VM sandbox, LLM, tool calls, and file access controls only exist in the user’s browser tab.
This blog will not wax poetic about the end result as if it’s something I feel like I can claim any credit for, because I can’t, and shouldn’t. I knew exactly what I wanted, how to do it, what to use, and how to validate success. AI put all the pieces together for smolbox exactly as I asked with zero ambiguity.
Smolbox is a proof-of-concept to point out “this was always possible” and talk about the disservice the state of AI security is doing for the next generation of their users.
I’m in my early 30’s. I’m sat down writing this blog because I pulled my back taking out the trash yesterday. I have a credit card and pay for two different AI subscriptions which run 24/7 agentic workflows without flinching about the cost. I have so many computers I lose them for months without noticing.
But way back sometime in the 2000’s I had a cell phone (LG Dare) and was a teenager. I had unlimited text and pics (SMS/MMS) messages, but “data” was still such a prohibitively expensive thing that accidentally tapping the wrong button on my phone would be visible on the monthly bill. I got a spam text message one day that was from an email address . After playing around a bit with sending messages to my-ten-digit-cell-number@vzwpix.com I’d figured out that there was an MMS<->Email bridge managed by Verizon. This led me to find an enterprising PHP developer online who ran an email service that allowed you to send an email with a subject line as a URL and receive an email back with the contents of a web page. After asking them nicely, they confirmed my email address and I was able to “browse” the internet via my unlimited text messages. This free access to information without associated cost increase lead me to read the developer manuals and SDK for the phone itself (with the phone locking up and freezing for minutes at a time as it tried to handle a full web page of content where a small text message was expected), eventually realizing that there was a path traversal vulnerability when ringtones were loaded via the MicroSD card slot, and being able to overwrite the builtin World Clock “app” written in MacroMedia Flash Lite with a mini-golf game.
I think about that random developer I reached out to a lot. The service was slow. The formatting of the web pages sucked. It made my phone freeze up. And yet it was everything I had ever wanted, because it made something possible .
They probably didn’t think much of it, or were just happy to see someone using their small service they’d built, but they opened the floodgates and set me down a path for what would ultimately shape my entire career path. All they did was make something possible .
But are teenagers agentic enough?
If I were a teen today, I have no doubts I’d have figured out how to get access the internet. Kids are resourceful like that. Whether it be a school-issued iPad, computer at the local library, or a used hand-me-down laptop that was no longer needed. That’s the easy part.
But can they install Claude code on those locked down school-issued iPads? Are they able to install arbitrary programs on the public computers at the library? Can they purchase an AI subscription to learn how these things work? Do they even have a credit card to purchase that subscription?
Yes, yes they can! There’s always someone who knows the loopholes to getting around the guardrails of those locked down systems, but that comes with caveats of limited access and getting in trouble when you’re caught. It’s the same with the currency problem. There’s always a slightly older friend who has a card who can be bartered with. Alternatively, pick a random major online game with tradeable assets or an in-game currency and you’ll soon find an economy that operates in such a way that it’s not entirely unlike money laundering and fraud whereby real tangible currency exits the system in a readily usable form. Then they’ll probably stretch that money as far as they can by buying cheap “unused” (hacked accounts) token capacity through fly-by-night vendors on some random Telegram messaging channel.
If I were a teen today, I’d almost certainly want to be playing with agentic AI. After all, it’s the future, right? If I were a teen today, I’d almost certainly take the path that made it possible .
In computer security, a sandbox is a security mechanism for separating running programs, usually in an effort to mitigate system failures and/or software vulnerabilities from spreading. The sandbox metaphor derives from the concept of a child’s sandbox—a play area where children can build, destroy, and experiment without causing any real-world damage. - https://en.wikipedia.org/wiki/Sandbox_(computer_security)
Smolbox is a sandbox; a small one. A safe play area that cannot cause any real-world damage, requiring no particularly special permissions. It works on school issued iPads, library computers, and hand-me-down used laptops for free. Sandboxes are trivial to extend the features of. You just put more things in the box to play with.
Smolbox is slow. It’s kludgy at times. It’s admittedly not particularly great at anything other than being private and not collecting any data whatsoever.
But it makes things possible for a large population of people to do safely, and it’s easy for a random developer to do. I know exactly what that can mean to someone.
It’s possible. What remains to be done is only a question of who is both capable and willing, and right now my back hurts and I’m a little grumpy.
