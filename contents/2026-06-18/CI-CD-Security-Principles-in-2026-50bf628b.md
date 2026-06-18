---
source: "https://worklifenotes.com/2026/06/18/ci-cd-security-principles-in-2026/"
hn_url: "https://news.ycombinator.com/item?id=48584680"
title: "CI/CD Security Principles in 2026"
article_title: "CI/CD Security Principles in 2026 - Work & Life Notes"
author: "taleodor"
captured_at: "2026-06-18T13:16:54Z"
capture_tool: "hn-digest"
hn_id: 48584680
score: 1
comments: 0
posted_at: "2026-06-18T13:03:56Z"
tags:
  - hacker-news
  - translated
---

# CI/CD Security Principles in 2026

- HN: [48584680](https://news.ycombinator.com/item?id=48584680)
- Source: [worklifenotes.com](https://worklifenotes.com/2026/06/18/ci-cd-security-principles-in-2026/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T13:03:56Z

## Translation

タイトル: 2026 年の CI/CD セキュリティ原則
記事のタイトル: 2026 年の CI/CD セキュリティ原則 - 仕事と生活のメモ
説明: これは、私の古い投稿「最新の CI/CD の 7 つのベスト プラクティス」の続編です。そこに概説されている点は依然として当てはまりますが、いくつかの重要な点が欠けています。

記事本文:
コンテンツにスキップ
仕事と生活のメモ
Pavel Shukhman のブログ: テクノロジー、旅行、人間関係、生活
2026 年の CI/CD セキュリティ原則
これは、私の古い投稿「最新の CI/CD の 7 つのベスト プラクティス」の続編です。そこで説明されている点は今でも当てはまりますが、セキュリティに関する重要な考慮事項がいくつか抜けています。 2026 年の今日、CI/CD パイプラインは主要なサプライ チェーン攻撃ベクトルの 1 つとなっています (たとえば、最近の Trivy 侵害を参照)。それは更新を保証するものです。
1. 冗長性: 侵害に成功するには、少なくとも 2 つの独立したシステムに障害が発生する必要があります
2. 異なるパイプラインで認証情報を共有してはなりません
4. 安全でない入力または悪意のある入力を想定する
5. CI によって使用されるすべての依存関係をピン留めする
1. 冗長性: 侵害に成功するには、少なくとも 2 つの独立したシステムに障害が発生する必要があります
これは基本原則ですが、抽象的でもあります。以下のものはより実用的です。ここでは、すでにどこかで妥協が行われている可能性があると仮定します。したがって、各システムは、特定のシステムが侵害されただけでは攻撃者が目的を達成できないように設計する必要があります。この原則の実際の実装については、ReARM ブログを参照してください。
2. 異なるパイプラインで認証情報を共有してはなりません
ライブラリの従来の GitHub Actions パブリッシング パイプラインを考えてみましょう。すべてのワークフロー YAML ファイルは 1 つのディレクトリにあります。そのうちの 1 つは初期トリアージを実行し、別の 1 つは実際のビルドを実行し、さらに別の 1 つは最終アーティファクトを公開します。
GitHub Actions の仕組みを考えると、すべてのシークレットはこれらすべてのファイル間で共有されます。そのため、トリアージ パイプラインで重大な侵害が発生すると、発行者の資格情報が公開され、攻撃者が実際の発行パイプラインに一切触れずに悪意のあるペイロードを発行できる可能性があります。
これを軽減するにはどうすればよいでしょうか?初代校長の精神を貫く

ファイル: 異なるパイプラインが同じシークレット プールを共有しないようにします。これは、GitHub Actions の場合、パイプラインが異なるリポジトリに存在する必要があることを意味します。
これも最初の原則を拡張したものです。具体的には、ビルド パイプラインとパブリッシュ パイプラインの間に、成果物を評価できるポイントが必要です。アーティファクトをワンショットで構築してプッシュする単一のパイプラインがあるということは、攻撃者がそのようなパイプラインを侵害することに成功した場合、追加の保護策なしで何かを公開することもできることを意味します。
4. 安全でない入力または悪意のある入力を想定する
これはパブリック リポジトリにとって特に重要です。 GitHub Actions の pull_request_target (決して使用すべきではない) については多くのことが言われていますが、攻撃者は正当なアクションが受け入れる入力を侵害しようとする可能性もあります。ここで使用できる、または使用すべきリンターがあります。 GitHub アクションの Zizmor。
5. CI によって使用されるすべての依存関係をピン留めする
あらゆる種類のサプライチェーン攻撃が増加しており、CI は実験の場ではありません。 CI 自体の依存関係やビルドの依存関係を含むすべての依存関係は、可能な限りダイジェストによって固定される必要があります。
これには、ハッシュによる CI パイプラインの参照、可能な場合はロックファイルのみに基づいてインストール、ダイジェストによる Docker ベース イメージの固定が含まれます。
CI の成果物と成果物には署名する必要があります。 CI の下流段階では、続行する前に署名を検証する必要があります。 CD は展開前にすべてを検証する必要があります。それが基本原則です。
現在、多くの情報源が、このプロセスの中核要素として署名ではなく証明書について議論しています。本質的に、認証とは何かについてのステートメントです。このステートメントは、コンテキストに基づいて暗黙的または明示的になります。
暗黙的な構成証明の例は、GitHub Actions からのパブリック Sigstore を使用してビルドに署名することです。そういったことの検証

署名は、ビルドが特定の GitHub アクションによって生成されたことを意味します。明示的な構成証明には、署名される特定のステートメントも含まれます (たとえば、この SBOM はこのビルドに属します)。多くの場合、実際には、サプライ チェーンの検証には暗黙的な認証で十分です。ただし、in-toto などの明示的な認証システムは、特定のシナリオでは追加の利点をもたらす可能性があります。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
最終防衛線はAIであってはなりません
2026 年の CI/CD セキュリティ原則
最終防衛線はAIであってはなりません
ReARM: ガバナンス AI コーディング エージェントのデモ
新しい ReARM ピッチ – エージェント コーディングのガバナンス
パラダイムが変わるとき: AI エージェントのゼロトラスト モデル
2026 年の CI/CD セキュリティ原則 - 最新の CI/CD の 7 つのベスト プラクティスに関する仕事と生活のメモ
パラダイム シフトのとき: AI エージェントのゼロトラスト モデル - Git Push の恐怖を取り除くための仕事と生活のメモ
OnPod で QA とサイバー セキュリティについて話しました - 現在のテクノロジー時代を生き抜きたいという仕事と生活のメモ - 優れた QA になることを学びましょう
NTIA 準拠の SBOM の実践ガイド - 単一の SBOM では十分ではない理由に関する仕事と生活のメモ
開発マシンを信頼できないものとして扱い始めるべき時 - Linux および Chromebook の SSH 用 YubiKey に関する仕事と生活のメモ

## Original Extract

This is a follow up on my older post "7 Best Practices of Modern CI/CD". Points outlined there still hold true, but they are missing several important

Skip to content
Work & Life Notes
Pavel Shukhman's blog: Tech, Travels, Relationships, Life
CI/CD Security Principles in 2026
This is a follow up on my older post “ 7 Best Practices of Modern CI/CD “. Points outlined there still hold true, but they are missing several important security considerations. Today, in 2026, CI/CD pipelines have become one of the key supply chain attack vectors (refer, for example, to the recent Trivy compromise). That warrants an update.
1. Redundancy: At Least 2 Independent Systems Need to Fail for a Successful Compromise
2. Different Pipelines Must Not Share Credentials
4. Assume Unsafe or Malicious Inputs
5. Pin All Dependencies Consumed by CI
1. Redundancy: At Least 2 Independent Systems Need to Fail for a Successful Compromise
This is the core principle, but it is also abstract. The following ones will be more actionable. Here we assume that compromise may already have happened somewhere. Therefore, each system must be designed in such a way that any particular compromised system should not be enough for an attacker to reach their goal. See a practical implementation of this principle in the ReARM Blog .
2. Different Pipelines Must Not Share Credentials
Consider a traditional GitHub Actions publishing pipeline for a library. All workflow YAML files are located in one directory. One of them does initial triage, another one does the actual build, yet another one publishes the final artifact.
Given how GitHub Actions works, all secrets are shared between all these files. So a significant compromise in the triage pipeline may also expose publisher credentials and allow an attacker to publish a malicious payload without ever touching the actual publishing pipeline.
How to mitigate this? Follow the spirit of the first principle: make sure that different pipelines don’t share the same secret pools – which in the case of GitHub Actions means they must live in different repositories.
This again expands on the first principle. Specifically, there must be a point between build and publish pipelines at which deliverables can be assessed. Having a single pipeline that builds and pushes artifact in one shot means that an attacker who manages to compromise such a pipeline can also publish something without further safeguards.
4. Assume Unsafe or Malicious Inputs
This is especially important for public repositories. A lot has been said about pull_request_target on GitHub Actions (that should never be used), but an attacker may also try to compromise any input that a legitimate action accepts. There are linters that can and should be used here, e.g. Zizmor for GitHub Actions.
5. Pin All Dependencies Consumed by CI
Supply chain attacks of all kinds are on the rise, and CI is not a place for experiments. All dependencies, including CI’s own dependencies and any build dependencies, should be pinned by digests to the maximum extent possible.
This includes referencing CI pipelines by hash, installing only based on lockfiles where possible, pinning any Docker base images by digests.
CI deliverables and artifacts must be signed. Downstream stages of CI must verify signatures before proceeding. CD must verify everything before deploying. That’s the core principle.
A lot of sources now discuss attestations rather than signatures as the core element in this process. Essentially, attestation is a statement about something; that statement can be implicit or explicit based on the context.
An example of implicit attestation is signing a build with public Sigstore from GitHub Actions. Verification of such a signature means that the build was produced by the specific GitHub Action. An explicit attestation would also contain a specific statement being signed (e.g., this SBOM belongs to this build ). In many cases in practice an implicit attestation is enough for supply chain verification purposes; however explicit attestation systems such as in-toto can bring additional benefits in certain scenarios.
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
The last line of defense must not be AI
CI/CD Security Principles in 2026
The last line of defense must not be AI
ReARM: Governing AI Coding Agents Demo
New ReARM Pitch – Governance for Agentic Coding
When the Paradigm Shifts: A Zero-Trust Model for AI Agents
CI/CD Security Principles in 2026 - Work & Life Notes on 7 Best Practices of Modern CI/CD
When the Paradigm Shifts: A Zero-Trust Model for AI Agents - Work & Life Notes on Take The Fear Out Of Git Push
Talked about QA and Cyber Security at OnPod - Work & Life Notes on Want to Survive Current Tech Era – Learn to Be a Good QA
Practical Guide to NTIA Compliant SBOM - Work & Life Notes on Why a Single SBOM is Never Enough
Time to Start Treating Dev Machines as Untrusted - Work & Life Notes on YubiKey for SSH on Linux and Chromebook
