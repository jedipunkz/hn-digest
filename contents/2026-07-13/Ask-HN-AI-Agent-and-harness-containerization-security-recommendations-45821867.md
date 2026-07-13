---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48899674"
title: "Ask HN: AI Agent and harness containerization/security recommendations"
article_title: ""
author: "dv35z"
captured_at: "2026-07-13T22:45:03Z"
capture_tool: "hn-digest"
hn_id: 48899674
score: 2
comments: 0
posted_at: "2026-07-13T22:19:31Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: AI Agent and harness containerization/security recommendations

- HN: [48899674](https://news.ycombinator.com/item?id=48899674)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T22:19:31Z

## Translation

タイトル: HN に聞く: AI エージェントとコンテナ化/セキュリティの推奨事項の活用
HN テキスト: こんにちは、HN クルーです。AI エージェントがローカル プロジェクトやユーザー フォルダー、さらにはそれ以降 (オペレーティング システム ファイル) にアクセスすることを許可する傾向が見られます。私は次のような質問をしようと思いました。
システムとデータの整合性を保護し、生産的に使用できるほど簡単にするには、ワークフローでローカル システム コマンドやネットワーク コマンドが実行されることが多い AI ツールを安全に使用するにはどうすればよいでしょうか?コメント構造のアイデア: オペレーティング システム / エージェント ハーネス / エージェント / セキュリティ戦略とツール スタック / ワークフロー いくつかの例を挙げると、現在 Linux Mint Debian Edition (LMDE) で OpenCode (+DeepSeek) を使用していますが、エージェントとハーネスが /tmp/ にファイルを作成するように頻繁に要求していることに気付きました。通常はこれを禁止し、現在のフォルダー内のファイルのみを使用するように指示しています。私としては、ツールがシステム強制的な方法で特定のプロジェクト フォルダー (~/Projects/PROJECT-NAME) にのみアクセスできるようにしたいと考えています。ただし、エージェントは多くの場合、システム コマンド (ローカル開発サーバーをデバッグするための「ps」、ファイル変換のための「pandoc」または「ffmpeg」など) を実行したいことがわかります。これにより、境界線があいまいになり始めます。そこで、AI エージェントが役立つようにするには、分離されたオペレーティング システムへのアクセスを与えることを検討できるのではないかと考えています。その結果、仮想マシンや Docker コンテナーなどについて考えるようになりました。そのため、「システム/VM 間で共有フォルダーを使用する必要があるか?」などの複雑さが生じます。あなたにとって何がうまくいっているのか、「理想的な」状態、そしてそれを実際にどのように実装できるのかをぜひ聞きたいです。言っておきますが、私は初心者にテクノロジ (AI を含む) を頻繁に教えています。そのため、彼らが何かを簡単に効果的に実行できるようにしながら、デフォルトでセキュリティ/整合性の実践を提供できるバランスを見つけようとしています。

ひどい混乱に陥らないように（「AI が私のプロジェクト/コンピューターを削除した！」）、うまくいきます。ありがとう！

## Original Extract

Hello HN crew - I am seeing the tendency for people to allow AI agents to access local project & user folders, and beyond (operating system files). I thought to ask the question:
How can we best use AI tools safely - where the workflow often runs local system commands and network commands - to protect the integrity of our systems & data AND be easy enough to use productively? Comment structure idea: Operating system / Agent harness / Agent / Security strategy & tool stack / Workflow To give some examples: I am currently using OpenCode (+DeepSeek) on Linux Mint Debian Edition (LMDE), and I notice that the Agent & harness is frequently asking to create files in /tmp/ - I usually forbid this, and instruct it to only use files in the current folder. I would rather that the tool ONLY have access to a given project folder (~/Projects/PROJECT-NAME) in a system-enforced way. However, I see that the agent often wants to run system commands (like `ps` to debug a local dev server, `pandoc` or `ffmpeg` for file conversion, etc). This starts blurring the line - so I'm thinking that in order for the AI agent to be useful, I can consider giving it access to an isolated operating system - leading me to think about Virtual Machines, Docker containers, and so on. That introduces complexity like, "Should I be using shared folders between my system/VM?" etc. Would love to hear what's been working for you, the "ideal" state, and practically how we can implement it. I ought to mention that I'm frequently teaching technology (including AI) to somewhat newbies - so I am trying to find that balance, that lets them get easily something done effectively, but gives them security/integrity practices by default starting off, so they don't get into an awful jam ("The AI deleted my project/computer!"). Thanks!

