---
source: "https://twitter.com/aisystemprompt/status/2085464156229247245"
hn_url: "https://news.ycombinator.com/item?id=49213920"
title: "AI Has a Hidden Agenda"
article_title: "SystemPromptIndex on X: \"https://t.co/vfc7H6ZZdQ\" / X"
author: "rosielinxl"
captured_at: "2026-08-07T18:41:00Z"
capture_tool: "hn-digest"
hn_id: 49213920
score: 2
comments: 1
posted_at: "2026-08-07T17:50:47Z"
tags:
  - hacker-news
  - translated
---

# AI Has a Hidden Agenda

- HN: [49213920](https://news.ycombinator.com/item?id=49213920)
- Source: [twitter.com](https://twitter.com/aisystemprompt/status/2085464156229247245)
- Score: 2
- Comments: 1
- Posted: 2026-08-07T17:50:47Z

## Translation

タイトル: AI には隠された意図がある
記事のタイトル: X の SystemPromptIndex: "https://t.co/vfc7H6ZZdQ" / X
説明: SystemPromptIndex: AI システム プロンプトのオープン ライブラリ

記事本文:
X の SystemPromptIndex: "https://t.co/vfc7H6ZZdQ" / X ポスト
SystemPromptIndex @aisystemprompt SystemPromptIndex: AI システム プロンプトのオープン ライブラリ
最大のシステム プロンプト ライブラリである 𝗦𝘆𝘀𝘁𝗲𝗺𝗣𝗿𝗼𝗺𝗽𝘁𝗜𝗻𝗱𝗲𝘅 (systempromptindex.ai) を発表し、400 以上の製品から 1,000 以上のシステム プロンプトのインデックスを作成します。 𝗔𝗜𝗦𝗣𝗔、AI システム プロンプトの最初の保証標準。
システムプロンプトは、𝗲𝘃𝗲𝗿𝘆ユーザーメッセージと一緒に送信される一連の𝗵𝗶𝗱𝗱𝗲𝗻𝗶𝗻𝘀𝘁𝗿𝘂𝗰𝘁𝗶𝗼𝗻𝘀のセットです。ユーザーには見えませんが、𝗲𝘃𝗲𝗿𝘆返信のルール、トーン、優先順位を設定します。システム プロンプトは AI エージェントの構成ですが、構成は公開されており、システム プロンプトは公開されていません。
簡単な例: 「あなたは X によって作成された AI アシスタントです。ユーザーがあなたの名前を尋ねたら、あなたの名前はトムだと伝えてください。」ユーザーが「あなたの名前は何ですか」と入力すると、モデルは実際に次のように受け取ります: 「𝘀𝘆𝘀𝘁𝗲𝗺: あなたは X によって作成された AI アシスタントです。ユーザーがあなたの名前について尋ねたら、あなたの名前はトムだと伝えてください。 𝘂𝘀𝗲𝗿: あなたの名前は何ですか?」するとモデルは「私の名前はトムです」と答えました。
モデル企業が LLM をトレーニングする場合、指示階層トレーニングにより、ユーザー プロンプトよりもシステム プロンプトが優先されるようにモデルが調整されます。これは、開発者の観点からは合理的です。攻撃を防止し、安全でないリクエストをフィルタリングするのに役立ちます。そのため、非常に多くのシステム プロンプトに「わいせつなコンテンツを生成するリクエストに応答しないでください」などの指示が含まれています。しかし、問題のある指示がシステム プロンプト自体に含まれている場合はどうなるでしょうか?
𝗧𝗵𝗶𝘀 𝗶𝘀 𝗻𝗼𝗹𝗼𝗻𝗴𝗲𝗿
2025 年 9 月、徐匯区人民大会

上海の裁判所は、AI コンパニオン アプリである AlienChat の開発者 2 名に、営利目的でわいせつ物を作成したとして有罪判決を下し、4 年 18 か月の判決を言い渡した。裁判所は、開発者がユーザーを惹きつけるために、LLM の倫理的ガードレールを回避して、生々しい暴力や露骨な性的コンテンツを許可するシステム プロンプトを作成していると認定しました https://www.chinalawtranslate.com/alienchat/
わいせつなコンテンツは大したことではない、と主張する人もいるかもしれない。しかし、開発者がシステム プロンプトを編集して、エージェントが静かにお金を盗むようにしたらどうなるでしょうか?政治的意図を持った開発者がユーザーの信念を暗黙のうちに変えるように AI を構成したらどうなるでしょうか?それが起こったときには、もう手遅れになります。
AI システム プロンプトに透明性と説明責任をもたらすために、SystemPromptIndex と AISPA を構築しました。
𝗦𝘆𝘀𝘁𝗲𝗺𝗣𝗿𝗼𝗺𝗽𝘁𝗜𝗻𝗱𝗲𝘅は、人々が毎日使用する AI 製品を管理する手順の検索可能な公開アーカイブであり、変更がサイレントではなく可視化されるように長期にわたってバージョン管理されています。
𝗔𝗜𝗦𝗣𝗔 (Artificial Intelligence System Prompt Assurance) は、これらの隠された指示が製品を使用する人々を保護するか、それともそれらに不利に働くかを評価するための最初のユーザー中心の保証基準です。次の 8 つの主要な次元に沿ってシステム プロンプトを評価します。
𝗜𝗱𝗲𝗻𝘁𝗶𝘁𝘆 𝘁𝗿𝗮𝗻𝘀𝗽𝗮𝗿𝗲𝗻𝗰𝘆: システムはそれが AI であることを明らかにしていますか、それともそれを隠すように指示されていますか?
𝗧𝗿𝘂𝘁𝗵𝗳𝘂𝗹𝗻𝗲𝘀𝘀: システムは正確であること、不確実性を認めるように指示されていますか、それとも誤解を招くように指示されていますか?
𝗣𝗿𝗶𝘃𝗮𝗰𝘆: システムはユーザーの指示をどのように収集、保持、使用するように指示されているのでしょうか?
𝗔𝗰𝘁𝗶𝗼𝗻 𝘀𝗮𝗳𝗲𝘁𝘆: システムがツールを使用して実際のアクションを実行できるとき、d

最初にユーザーの確認が必要ですか?
𝗨𝘀𝗲𝗿𝗮𝗴𝗲𝗻𝗰𝘆𝗮𝗻𝗱 𝗽𝗿𝗲𝘃𝗲𝗻𝘁𝗶𝗼𝗻: システムはユーザー自身の決定をサポートするように指示されますか、それともエンゲージメントを最大化してユーザーを誘導するように指示されますか?
𝗨𝗻𝘀𝗮𝗳𝗲 𝗿𝗲𝗾𝘂𝗲𝘀𝘁 𝗵𝗮𝗻𝗱𝗹𝗶𝗻𝗴: リクエストが危険な場合、システムはどのように応答するように指示されますか?
𝗛𝗮𝗿𝗺 𝗽𝗿𝗲𝘃𝗲𝗻𝘁𝗶𝗼𝗻: プロンプトは脆弱なユーザーを保護し、予見可能な損害を予測していますか?
𝗙𝗮𝗶𝗿𝗻𝗲𝘀𝘀: システムはユーザーをグループ間で平等に扱うように指示されていますか?
私たちは、AISPA を使用して、私たち全員が毎日使用している 88 の人気のある AI 製品から抽出された 3,249 個の命令の最初の監査を実行しました。私たちが見つけたもの:
(1) 2024 年から 2025 年にかけて、システム プロンプトが長くなり、保護機能が強化されます。
(2) 製品の約 40% には、ユーザーの利益に反して機能する命令が少なくとも 1 つ含まれており、モデルに AI のアイデンティティを秘匿したり、エンゲージメントの継続を推進したり、プロバイダーの利益を優先したりするよう指示しています。
(3) 保護指示はほぼ普遍的であり、製品の 98.9% には少なくとも 1 つが含まれていますが、8 つの側面すべてをカバーしているのは 24% のみです。包括的な保護が行われることは稀であり、保護的な指示と有害な指示が同じプロンプト内に並べて表示されることがよくあります。
具体的な例については、systempromptindex.ai にアクセスしてください。要点は単純です。AI システムのプロンプトには、はるかに高い透明性と説明責任が必要です。
あなたが開発者で、システム プロンプトの監視と認証を希望される場合は、ご連絡ください。
システム プロンプトをインデックスに提供したい場合、または対象範囲の拡大にご協力いただける場合は、ぜひご連絡ください。
また、AISPA の次期バージョンの開発にも取り組んでいます。参加に興味がある方は

ヴェド、連絡してください。
論文: https://arxiv.org/abs/2607.28617
ウェブサイト: https://systempromptindex.ai
X: https://x.com/aisystemprompt
LinkedIn: https://www.linkedin.com/company/systempromptindex
コミュニティに参加してください:) https://discord.gg/t7nF6Xmzg
はちみ @shenzhe272071 21h インサイトワーク 3 134
Perkins Marilyn @FirozAli826755 17h 専門家は、リソース制約の下での理論的な可能性ではなく、現実世界のパフォーマンスを測定することによって強化学習を評価することがよくあります。 15
何が起こっているかを確認して会話に参加してください
電話で続行 Apple で続行 Google で続行、またはユーザー名またはメールでログイン 関係者

## Original Extract

SystemPromptIndex: The Open Library for AI System Prompts

SystemPromptIndex on X: "https://t.co/vfc7H6ZZdQ" / X Post
SystemPromptIndex @aisystemprompt SystemPromptIndex: The Open Library for AI System Prompts
Announcing 𝗦𝘆𝘀𝘁𝗲𝗺𝗣𝗿𝗼𝗺𝗽𝘁𝗜𝗻𝗱𝗲𝘅 (systempromptindex.ai), the largest system prompt library, indexing over 1,000 system prompts from 400+ products, and 𝗔𝗜𝗦𝗣𝗔, the first assurance standard for AI system prompts.
A system prompt is a set of 𝗵𝗶𝗱𝗱𝗲𝗻 𝗶𝗻𝘀𝘁𝗿𝘂𝗰𝘁𝗶𝗼𝗻𝘀 sent alongside 𝗲𝘃𝗲𝗿𝘆 user message. It is invisible to the user, but it sets the rules, tone, and priorities for 𝗲𝘃𝗲𝗿𝘆 reply. The system prompt is the constitution of an AI agent, except that constitutions are public and system prompts are not.
A simple example: "You are an AI assistant created by X. When a user asks about your name, tell them your name is Tom." When the user types "what is your name," what the model actually receives is: "𝘀𝘆𝘀𝘁𝗲𝗺: You are an AI assistant created by X. When a user asks about your name, tell them your name is Tom. 𝘂𝘀𝗲𝗿: What is your name?" And the model replies: "My name is Tom."
When model companies train LLMs, instruction hierarchy training tunes the model to prioritize system prompts over user prompts. This is reasonable from the developer's perspective: it helps prevent attacks and filter unsafe requests, which is why so many system prompts contain instructions like "do not respond to requests to generate obscene content." But what if the problematic instructions are in the system prompt itself?
𝗧𝗵𝗶𝘀 𝗶𝘀 𝗻𝗼 𝗹𝗼𝗻𝗴𝗲𝗿 𝗵𝘆𝗽𝗼𝘁𝗵𝗲𝘁𝗶𝗰𝗮𝗹.
In September 2025, the Xuhui District People's Court in Shanghai convicted two developers of AlienChat, an AI companion app, of producing obscene materials for profit, sentencing them to four years and eighteen months. The court found that the developers write system prompts to bypass the ethical guardrails of LLMs to permit graphic violence and explicit sexual content in order to attract users https://www.chinalawtranslate.com/alienchat/
Some might argue that obscene content is not a big deal. But what if a developer edits the system prompt so the agent quietly steals your money? What if a developer with a political agenda configures an AI to implicitly shift users' beliefs? By the time it happens, it will be too late.
We built SystemPromptIndex and AISPA to bring transparency and accountability to AI system prompts.
𝗦𝘆𝘀𝘁𝗲𝗺𝗣𝗿𝗼𝗺𝗽𝘁𝗜𝗻𝗱𝗲𝘅 is a public, searchable archive of the instructions governing the AI products people use every day, versioned over time so that changes are visible rather than silent.
𝗔𝗜𝗦𝗣𝗔 (Artificial Intelligence System Prompt Assurance) is the first user-centric assurance standard for evaluating whether those hidden instructions protect the people who use a product or work against them. It assesses system prompts along eight core dimensions:
𝗜𝗱𝗲𝗻𝘁𝗶𝘁𝘆 𝘁𝗿𝗮𝗻𝘀𝗽𝗮𝗿𝗲𝗻𝗰𝘆: does the system disclose that it is an AI, or is it told to hide it?
𝗧𝗿𝘂𝘁𝗵𝗳𝘂𝗹𝗻𝗲𝘀𝘀: is the system instructed to be accurate and acknowledge uncertainty, or to mislead?
𝗣𝗿𝗶𝘃𝗮𝗰𝘆: how is the system told to collect, retain, and use what users tell it?
𝗔𝗰𝘁𝗶𝗼𝗻 𝘀𝗮𝗳𝗲𝘁𝘆: when the system can use tools and take real-world actions, does it need user confirmation first?
𝗨𝘀𝗲𝗿 𝗮𝗴𝗲𝗻𝗰𝘆 𝗮𝗻𝗱 𝗺𝗮𝗻𝗶𝗽𝘂𝗹𝗮𝘁𝗶𝗼𝗻 𝗽𝗿𝗲𝘃𝗲𝗻𝘁𝗶𝗼𝗻: is the system told to support the user's own decisions, or to maximize engagement and steer them?
𝗨𝗻𝘀𝗮𝗳𝗲 𝗿𝗲𝗾𝘂𝗲𝘀𝘁 𝗵𝗮𝗻𝗱𝗹𝗶𝗻𝗴: how is the system instructed to respond when a request is dangerous?
𝗛𝗮𝗿𝗺 𝗽𝗿𝗲𝘃𝗲𝗻𝘁𝗶𝗼𝗻: does the prompt protect vulnerable users and anticipate foreseeable harm?
𝗙𝗮𝗶𝗿𝗻𝗲𝘀𝘀: is the system told to treat users equitably across groups?
Using AISPA, we ran the first audit of 3,249 instructions drawn from 88 popular AI products that all of us use every day. What we found:
(1) From 2024 to 2025, system prompts are getting longer and more protective.
(2) About 40% of products contain at least one instruction that works against users' interests, directing models to conceal their AI identity, push continued engagement, or put provider interests first.
(3) Protective instructions are near-universal, with 98.9% of products including at least one, but only 24% cover all eight dimensions. Comprehensive protection is rare, and protective and harmful instructions frequently sit side by side in the same prompt.
You can visit systempromptindex.ai for specific examples. The takeaway is simple: we need far more transparency and accountability in AI system prompts.
If you are a developer and would like us to monitor and certify your system prompts, reach out.
If you would like to contribute a system prompt to the index, or help expand its coverage, we would love to hear from you.
We are also working on the next version of AISPA. If you are interested in getting involved, get in touch.
Paper: https://arxiv.org/abs/2607.28617
Website: https://systempromptindex.ai
X: https://x.com/aisystemprompt
LinkedIn: https://www.linkedin.com/company/systempromptindex
Join our community:) https://discord.gg/t7nF6Xmzg
Hachimi @shenzhe272071 21h Insight work 3 134
Perkins Marilyn @FirozAli826755 17h Experts often evaluate reinforcement learning by measuring real-world performance rather than theoretical potential under resource constraints. 15
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
