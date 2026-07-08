---
source: "https://azuanz.com/posts/making-ai-code-review-measurable/"
hn_url: "https://news.ycombinator.com/item?id=48829851"
title: "Making AI Code Review Measurable"
article_title: "Making AI Code Review Measurable | azuan"
author: "azuanrb"
captured_at: "2026-07-08T10:07:46Z"
capture_tool: "hn-digest"
hn_id: 48829851
score: 2
comments: 0
posted_at: "2026-07-08T09:50:02Z"
tags:
  - hacker-news
  - translated
---

# Making AI Code Review Measurable

- HN: [48829851](https://news.ycombinator.com/item?id=48829851)
- Source: [azuanz.com](https://azuanz.com/posts/making-ai-code-review-measurable/)
- Score: 2
- Comments: 0
- Posted: 2026-07-08T09:50:02Z

## Translation

タイトル: AI コードレビューを測定可能にする
記事のタイトル: AI コードレビューを測定可能にする |あずあん
説明: 過去のプル リクエスト、コーディング エージェント、コスト追跡を使用して、AI コード レビューのための反復可能な評価ワークフローを構築した方法。

記事本文:
投稿
AI コードレビューを測定可能にする
ソフトウェア エンジニアにとって、コード レビューは常にボトルネックの 1 つでした。 loveholidays では、日々のワークフローに AI を多用しています。 PR が AI の助けを借りて書かれたかどうかにかかわらず、誰かがそれをレビューする必要があります。
次に、Intercom がコード レビューだけでなく承認にも AI をどのように使用しているかについてのスレッドを読みました。危険に思えますが、非常に興味深いものでもあります。
免責事項: この投稿は推奨するものではありません。これは、このスレッドと、そのようなシステムがどのように機能するかについての私自身の解釈に触発された実験にすぎません。これも、OpenAI が 3 か月の試用期間を提供してくれたからこそ可能になりました。その過程で、私は数十億のトークンを焼き尽くしました。素晴らしい学習体験でしたが、決して安くはありませんでした。
すでに多くのコードレビューツールが世に出ています。そのうちのいくつかは素晴らしいものであると確信していますし、中には私が構築したものよりも優れているものさえあるかもしれません。しかし、AI 以前から、私は常にコード レビュー ツールに対して愛憎の感情を抱いてきました。当たり外れが多いです。また、コストが高い場合、そのツールが実際にどれだけ役立つかを理解せずに、そのツールを広く推奨することは困難です。
また、既存のワークフロー内でローカルに実行でき、個人または企業を問わず、既存の ChatGPT Codex サブスクリプションを使用できるものが必要でした。本当に、もう購読料を払いたくありませんでした。
ほとんどすべてのコーディング エージェントには非対話モードがあります。 Codex を使用すると、 codex exec を通じてトリガーできます。私はこのアプローチであらゆる種類の自動化を実行しました。
codex exec を使用してコードレビューを実行することにしました。 git diff のみを使用すると、はるかに高速かつ安価になりますが、個人的にはそのアプローチは好きではありません。コンテキストは非常に重要です。
プル リクエストをレビューするとき、差分を個別に確認することはほとんどありません。関連ファイルを開く

、周囲のコードがどのように動作するかを確認し、同様のパターンを探し、変更によって既存の動作がどのように破壊されるかを考えます。 AI レビュー担当者にも同じことができるようにしたいと考えました。
私がこのアプローチを気に入ったもう 1 つの理由は、それが私の既存のワークフローに適合していることです。エージェントは、私がすでに日常的に使用しているものと同じスキル、MCP サーバー、コマンド、プロジェクト コンテキストを使用できます。そのため、レビューは一般的な外部ツールというよりは、私の既存の作業方法の延長に近いものになります。
基本的な流れは次のようになります。
プル リクエストのクローンを新しい git ワークツリーの場所に作成します。 git worktree を使用すると、複数の評価を並行して簡単に実行できます。
カスタム プロンプト、ルール、出力形式を使用して codex exec を実行します。リスクと制限を理解している限り、サンドボックス、自分のマシン、または使いやすい環境でこれを実行してください。
構造化された JSON 出力を返すようにエージェントに依頼します。
{
"評決" : "却下" ,
「レビュー」: [
{
"ファイル" : "src/file_1.ts" ,
「行」：42、
"comment" : "これにより機能 A が壊れます。" 、
"suggestion" : "このように書き直すことを検討してください。"
}
]
}
このアプローチにより、Codex はレビューを実行する前に関連ファイルを検査し、コンテキストを収集できます。出力は構造化されているため、システムを中心に単純な Web インターフェイスを構築することもできます。
次のステップは質問に答えることでした。コードレビューシステムを素早くテストするにはどうすればよいですか?
Twitter スレッドと同様に、過去のプル リクエストを真実の情報源として使用することにしました。これは明らかに単純化されていますが、プル リクエストには既知のライフサイクルがあります。通常、特定のコミットでは、期待される結果は明らかです。これは承認されるべきでしょうか、それとも作成者が変更を加えるまでブロックされるべきでしょうか?
それを念頭に置いて、模擬レビューの備品として以前のプル リクエスト データを使用しました。
PR#1 10x エンジニアがレビューをリクエストしたところ、すべてが LG でした

TM.
PR#2 著者は commit bbb でレビューをリクエストし、レビュー担当者はコメントを残しました。
作成者は commit ccc でプル リクエストを更新し、最終的に承認されました。
PR#3 このプル リクエストは当初承認されましたが、誤ってインシデントを引き起こしてしまいました。
模擬審査ではこのPRを却下してほしいです。このシナリオについて質問してくれたチーム リーダーに感謝します。
ここでも同僚からのフィードバックが非常に役に立ちました。 They pointed out scenarios that I had missed, especially around the kind of issues that should block a pull request versus comments that are merely nice to have.これにより、試合リストがより便利になり、評価がチームからの実際のレビューの期待をより適切に反映できるようになりました。
次の質問は、このタイプのプル リクエストに対して、どのプロンプトが最も信頼性の高いレビューを生成するかということでした。
コーディング エージェントと模擬レビューのセットアップを適切に配置すると、さまざまなプロンプトをテストして結果を比較できました。
あなたは一般的な上級ソフトウェア エンジニアです。この PR を確認して、間違いのないようにしてください。
あなたは予約後チームのシニア ソフトウェア エンジニアです。
ここで eval セットアップが非常に役に立ちます。あるプロンプトが別のプロンプトよりも優れているかどうかを推測する代わりに、同じ過去の例に対して両方を実行し、結果を比較することができました。
このステップがどれほど貴重であるかは、どれだけ強調してもしすぎることはありません。 eval を使用すると、変更が正しい方向に進んでいるかどうかをより決定的な方法で測定できます。
模擬レビュー表のタグに注目してください。すべてのチームとドメインは異なるという考えです。予約関連のプル リクエストには、支払い関連のプル リクエストとは異なるプロンプト、チェックリスト、または優先順位が必要になる場合があります。タグを使用すると、特定の領域を対象とした評価を実行し、それに応じてレビュー担当者を調整できます。
正しいプロンプトは何ですか?理想的には、最適なプロンプトを自分で探す必要はありません

f.したがって、代わりに、システムはそれ自体を反復処理して、独自のコード レビュー プロンプトを改善できる必要があります。完璧とは程遠いですが、一部のプロンプトでは当初の 50% から 88% の精度を得ることができました。
コーディング エージェントを十分に長く使用すると、モデルや推論レベルだけでなく、ハーネスも重要であることがわかるでしょう。
そのため、他のハーネスをサポートするためにコード レビュー システムを拡張しました。私のお気に入りの 1 つは pi です。これは、同様の方法で pi --print コマンドを通じて実行できます。
模擬レビュー システムをセットアップすると、さまざまな変数を比較するのがはるかに簡単になりました。
一般的なプロンプトとチーム固有のプロンプト
完全なリポジトリ コンテキストとより制限されたコンテキスト
クロスリポジトリ コンテキスト (バックエンド、フロントエンド、インフラストラクチャ)
これにより、システムは 1 人の特定のエージェントではなく、反復可能な評価ワークフローを重視したものになりました。また、短期間に何十億ものトークンを使い切ることが危険なほど簡単になりました。
コストの計測も重要です。コーディング エージェントの実行ごとに、セッション ID が生成されます。 ccusage と組み合わせると、コードレビューごとにトークンの使用量とコストを確認できます。
ccusage セッション --id 4aa0a585-b135-47e4-bdfb-c07bb958942d
pi が興味深い理由の 1 つは、高度にカスタマイズ可能であり、デフォルトで肥大化が少ないため、実行コストが安くなる可能性があることです。この設定を使用すると、その理論に真実があるかどうかをテストできます。
たとえば、あるハーネスが 30% 安くても許容精度範囲内にある場合、それを追求する価値があるかもしれません。コストはレビューの品質とともに測定する必要があります。
コードレビューは多くの場合双方向の会話です。コードを理解するのに助けが必要な場合があります。特に、コードがまだあまり経験のない領域に関わる場合です。
プル リクエストの作成者とペアになることもできますが、いくつかの質問しかない場合もあります

PR に関する質問はわかりません。 GitHub/GitLab のコード レビュー フローと同様に、Web UI でコードの任意の行をクリックして会話を開始できます。そして、コーディング エージェントは、関連するコンテキストを説明し、質問に答え、数秒または数分で変更を理解するのに役立ちます。
私にとって、システムがプル リクエストを自動承認することを完全には信頼していないとしても、これは依然として便利なツールです。これは、リスクのある領域を要約し、なじみのないコードを説明し、レビュー担当者がより適切な質問をするのに役立つレビューのコンパニオンとして機能します。
私が考えるもう 1 つの方法は、プル リクエストのレビューを他の人に依頼する前の最初のパスとしてです。明らかな問題を発見したり、危険な領域を早期に強調できれば、人間のレビュー担当者は実際に判断が必要な部分により多くの時間を費やすことができます。
これはスキル番号だった可能性があります
確かにそれは可能だ。派手な UI をすべて削除すると、このシステムの中核は、プロンプト、いくつかのルール、および構造化された出力形式だけになります。その形ではSKILLとして機能する可能性がある。
JSON 構造は重要な部分の 1 つです。これにより、レビュー担当者と AI エージェントの両方にコード レビュー プロセスの明確な形式が与えられます。
ここにはまだ多くのリスクがあります。
最大のものは誤った自信です。 AI は正しいように聞こえ、賢明なコメントを残しても、重要なことを見逃してしまう可能性があります。逆に、コードや製品のコンテキストを誤解したために、まったく問題のないプル リクエストをブロックすることもできます。
コスト、遅延、プライバシー、セキュリティ、説明責任などの現実的な問題もあります。 AI が承認したプル リクエストによってインシデントが発生した場合でも、誰かがその決定を所有する必要があります。
したがって、私はこれが人間の査読者の代わりになるとは考えていません。少なくとも盲目的ではありません。今のところ、私はレビューをより再現可能で測定可能にし、学習ツールとして有用にする方法だと考えています。
これはサイドプロジェクトとして始まりました

ct を使用し、主に gpt-5.5 ミディアムで Codex CLI を使用したバイブ コーディングを通じて構築されました。その結果には非常に感銘を受けました。
私にとって最大の気づきは、AI コードレビューは測定可能になって初めて面白くなるということです。たった 1 つの印象的なレビューにはあまり意味がありません。しかし、過去のプル リクエストに対して同じレビュアーを実行し、プロンプトを比較し、コストを測定し、どこで失敗したかを理解できれば、改善できるものになります。
私にとって重要な問題は、AI がコードをレビューできるかどうかではなく、そのレビュー プロセスを観察可能で再現可能で、エンジニアリング ワークフローに適合できるほど安全なものにできるかどうかです。
私は人間の査読者を置き換えようとしているわけではありません。より良いレビューワークフローを構築しようとしています。 AI が最初のパスを実行し、有用なコンテキストを表面化し、人間によるレビューをより焦点を絞ったものにすることができます。

## Original Extract

How I built a repeatable evaluation workflow for AI code review using historical pull requests, coding agents, and cost tracking.

Posts
Making AI Code Review Measurable
Code review has always been one of the bottlenecks for software engineers. At loveholidays , we have been using a lot of AI in our daily workflows. Whether a PR is written with the help of AI or not, someone still needs to review it.
Then I read a thread by Intercom on how they use AI not just for code review, but also to approve it. Sounds risky, but also very interesting.
Disclaimer: This post is not an endorsement. This was just an experiment inspired by the thread and my own interpretation of how a system like that might work. It was also only possible because OpenAI gave us a three-month trial. During the process, I burned through a few billion tokens. It was a great learning experience, but definitely not cheap.
There are already a lot of code review tools out there. I’m sure some of them are great, and some may even be better than what I built. But I have always had a love and hate relationship with code review tools, even before AI. They are often hit or miss. And when the cost is high, it is hard for me to recommend a tool broadly without first understanding how useful it actually is.
I also wanted something that could run locally inside my existing workflow and use my existing ChatGPT Codex subscription, whether personal or enterprise. I really, really did not want to pay another subscription.
Almost all coding agents have a non-interactive mode. With Codex, you can trigger it via codex exec . I ran all sorts of automation with this approach.
I decided to use codex exec to run the code review. Using only the git diff would be a lot faster and cheaper, but I personally do not like that approach. Context matters a lot.
When I review a pull request, I rarely look at the diff in isolation. I open related files, check how the surrounding code works, look for similar patterns, and think about how the change could break existing behaviour. I wanted the AI reviewer to be able to do the same.
Another reason I liked this approach is that it fits into my existing workflow. The agent can use the same skills, MCP servers, commands, and project context that I already use day to day. That makes the review less like a generic external tool and more like an extension of how I already work.
The basic flow looks like this:
Clone the pull request into a new git worktree location. git worktree makes it easier to run multiple evaluations in parallel.
Run codex exec with a custom prompt, rules, and output format. Run this in a sandbox, your own machine, or any environment you are comfortable with, as long as you understand the risks and limitations.
Ask the agent to return a structured JSON output.
{
"verdict" : "rejected" ,
"reviews" : [
{
"file" : "src/file_1.ts" ,
"line" : 42 ,
"comment" : "This will break feature A." ,
"suggestion" : "Consider rewriting it this way."
}
]
}
With this approach, Codex can inspect relevant files and gather context before performing the review. Because the output is structured, I can also build a simple web interface around the system.
The next step was to answer a question. How can I test the code review system quickly?
Similar to the Twitter thread, I decided to use past pull requests as the source of truth. This is obviously simplified, but a pull request has a known lifecycle. At a specific commit, the expected outcome is usually clear. Should this be approved, or should it be blocked until the author makes changes?
With that in mind, I used previous pull request data as fixtures for mock reviews.
PR#1 A 10x engineer requested a review and everything was LGTM.
PR#2 The author requested a review at commit bbb , and reviewers left comments.
The author updated the pull request at commit ccc , and it was finally approved.
PR#3 This pull request was originally approved, but it accidentally caused an incident.
In the mock review, I want this PR to be rejected. Thanks to my team lead, who asked about this scenario.
This is also where feedback from my colleagues helped a lot. They pointed out scenarios that I had missed, especially around the kind of issues that should block a pull request versus comments that are merely nice to have. That made the fixture list more useful, and made the evaluation better reflect real review expectations from the team.
My next question was what prompt would produce the most reliable review for this type of pull request.
With the coding agent and mock review setup in place, I could test different prompts and compare the results.
You are a generic senior software engineer. Review this PR and make no mistakes.
You are a senior software engineer for the post-booking team.
This is where the eval setup became really useful. Instead of guessing whether one prompt was better than another, I could run both against the same historical examples and compare the results.
I cannot stress enough how valuable this step is. With evals, I can measure whether the changes are moving in the right direction in a more deterministic way.
Notice the tags in the mock review table. The idea is that every team and domain is different. A booking-related pull request may need a different prompt, checklist, or priority from a payment-related pull request. With tags, I can run targeted evals for specific areas and tune the reviewer accordingly.
What is the right prompt? Ideally, I don’t want to search for the best prompt myself. So instead, the system should be able to iterate on itself and improve its own code review prompt. It is far from perfect, but some prompts managed to get 88% accuracy, up from the initial 50%.
If you use coding agents long enough, you will realise that the harness matters too, not just the model or reasoning level.
Because of that, I expanded the code review system to support other harnesses. One of my favourites is pi , which can be run through the pi --print command in a similar way.
Once the mock review system had been set up, it became much easier to compare different variables:
generic prompts vs team-specific prompts
full repository context vs more constrained context
cross-repository context (backend, frontend, infrastructure)
This made the system less about one specific agent and more about a repeatable evaluation workflow. It also made it dangerously easy to burn through billions of tokens in a short period of time.
Measuring cost is also important. For any coding agent execution, a session ID is generated. Combined with ccusage , I can check the token usage and cost for each code review.
ccusage session --id 4aa0a585-b135-47e4-bdfb-c07bb958942d
One reason pi is interesting is that it is highly customisable and less bloated by default, which can make it cheaper to run. With this setup, I can test whether there is any truth to that theory.
For example, if one harness is 30% cheaper but still stays within an acceptable accuracy range, it might be worth pursuing. Cost must be measured together with review quality.
Code review is often a two-way conversation. There are times when I need help understanding code, especially when it touches an area I do not have much experience with yet.
I could pair with the pull request author, but sometimes I just have a few quick questions about the PR. Similar to the GitHub/GitLab code review flow, I can click on any line of code in the web UI and start a conversation. And the coding agent can explain the relevant context, answer questions, and help me understand the change in seconds or minutes.
To me, even if I never fully trust the system to auto-approve pull requests, this is still a useful tool. It can act as a review companion to summarise risky areas, explain unfamiliar code, and help reviewers ask better questions.
Another way I think about it is as a first pass before asking another person to review the pull request. If it can catch obvious issues or highlight risky areas early, the human reviewer can spend more time on the parts that actually need judgement.
This could have been a SKILL #
It definitely could. If we remove all the fancy UI, the core of this system is just a prompt, a few rules, and a structured output format. In that form, it could work as a SKILL.
The JSON structure is one of the important parts. It gives both the reviewer and the AI agent a clearer shape for the code review process.
There are still plenty of risks here.
The biggest one is false confidence. The AI can sound right, leave sensible-looking comments, and still miss something important. It can also do the opposite and block a perfectly fine pull request because it misunderstood the code or the product context.
There are also practical problems like cost, latency, privacy, security, and accountability. If an AI-approved pull request causes an incident, someone still needs to own that decision.
So I do not see this as a replacement for human reviewers. Not blindly, at least. For now, I see it as a way to make reviews more repeatable, measurable, and useful as a learning tool.
This started as a side project and was built mostly through vibe coding using Codex CLI with gpt-5.5 medium and I was quite impressed with the result.
The biggest takeaway for me is that AI code review only becomes interesting once it is measurable. A single impressive review does not mean much. But if I can run the same reviewer against historical pull requests, compare prompts, measure cost, and understand where it fails, then it becomes something I can improve.
For me, the important question is not whether AI can review code, but whether we can make that review process observable, repeatable, and safe enough to fit into an engineering workflow.
I am not trying to replace human reviewers. I am trying to build a better review workflow. One where AI can take the first pass, surface useful context, and make the human review more focused.
