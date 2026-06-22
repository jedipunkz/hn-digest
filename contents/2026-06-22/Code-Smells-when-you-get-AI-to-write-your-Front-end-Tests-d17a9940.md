---
source: "https://howtotestfrontend.com/resources/frontend-ai-generated-test-code-smells"
hn_url: "https://news.ycombinator.com/item?id=48636394"
title: "Code Smells when you get AI to write your Front end Tests"
article_title: "Code Smells when you get AI to write your Frontend Tests | How To Test Frontend"
author: "howToTestFE"
captured_at: "2026-06-22T21:41:56Z"
capture_tool: "hn-digest"
hn_id: 48636394
score: 2
comments: 0
posted_at: "2026-06-22T21:22:35Z"
tags:
  - hacker-news
  - translated
---

# Code Smells when you get AI to write your Front end Tests

- HN: [48636394](https://news.ycombinator.com/item?id=48636394)
- Source: [howtotestfrontend.com](https://howtotestfrontend.com/resources/frontend-ai-generated-test-code-smells)
- Score: 2
- Comments: 0
- Posted: 2026-06-22T21:22:35Z

## Translation

タイトル: AI にフロントエンド テストを作成させるとコードの匂いがする
記事のタイトル: AI にフロントエンド テストを作成させるとコードの匂いがする |フロントエンドをテストする方法
説明: AI は非常に強力ですが、フロントエンド アプリケーションのコードを生成するときに、多くの微妙なコード臭が発生する可能性があります。注意すべき主なものは次のとおりです。

記事本文:
AI にフロントエンド テストを作成させるとコードの匂いがする |フロントエンドをテストする方法 HowToTest FrontEnd
ホーム コースを検索 ブログについて 最近のブログ投稿
FE AI が生成したテストでコードの匂いがする
テストで Faker を使用する必要がある理由
これらのツールを使用して CI/CD 時間を 50% 削減します
優れたフロントエンド テストとは何でしょうか?
Vitest ブラウザ モードのガイド - およびその使用を検討する必要がある理由
フロントエンド アプリケーションでのアクセシビリティのテスト
Vitest での非同期 RSC コンポーネントのテスト
Vitest テストに型チェックを追加する
React テストの act() について知っておくべきことすべて
Vitest/Jest で React Server Components (RSC) をテストする方法
これらのツールを使用して CI/CD 時間を 50% 削減します
FE テストの作成 - 35 以上のヒントとコツ
Vitest テストに型チェックを追加する
React テスト コース (React テスト ライブラリ)
React テスト ライブラリ DOM マッチャー
Vitest/Jest テストの課題
AI にフロントエンド テストを作成させるとコードの匂いがする
テストの追加を忘れているか、変更のごく一部のみの単体テストを行っている
AIレビューではテスト不足が指摘されないことが多い
AIに実装を書いた後にテストを書くように依頼する
フィクスチャを使用しない（引数/小道具として渡すために準備されたモックデータ）
既存のヘルパー関数がわからない
AI はレンダリングされた要素を照合するときに正規表現を好みます
正確な DOM 構造のテスト
追加のタイムアウトの追加 (特に e2e テスト)
実際の実装ではなく、テストをテストする
AI による beforeAll() または beforeEach() の使用
ここ 1 ～ 2 年で、ソフトウェア業界は大きく変わりました。 Claude Code などの AI ツールのおかげで、1 年前と比べて、私が手動で入力するコードの割合はほんのわずかになりました。
私の日々の仕事は React を使用し、Vitest、Playwright などを使用してテストしています。おそらくほとんどの人が使用している非常に標準的なセットアップです。もちろんテストは重要な部分です

ソフトウェアを配信しています。
そして AI は確実にコードを記述し、そのコードにテストを追加できます。
しかし、私は AI が常に導入する多くのパターンに気づきました。これらはコードの臭いだと考え、通常は避けようとします。
私は、AI によってもたらされるテスト コードの匂いをいくつか文書化したいと考えていました。
これらは多くの場合、テストを迅速に通過させるのに役立ちますが、数か月後に初めて明らかになるメンテナンスの問題を引き起こす可能性もあります。
これを読んでいるあなたは、これらの点のいくつかに同意できないでしょう。
特に、規約をカバーする適切な SKILL.md ファイルを使用して、適切な AGENTS.md (または CLAUDE.md ) のセットアップにすでに時間を費やしている場合は特にそうです。
私がこれを書いているのは、AI を正しい方向に導く手綱がない場合に、AI がテストでどのような問題を引き起こす可能性があるかを指摘するためです。
本当の解決策は、 AGENTS.md 、 CLAUDE.md 、および SKILL.md ファイルを改良することです。
今後、クロード コードを活用して、より優れた保守可能なテストを作成するために使用する具体的なスキルや手順をブログ投稿する予定です。
これらのコードの多くは互いに矛盾しています (例: 過剰なテストと過小なテスト!)。これらの一般的なコードの匂いは、すべて同時に現れるわけではありません。
これらのベスト プラクティスの多くは、一般に単なるベスト プラクティスであるため、このサイトの他の場所で説明したものと似ています。
テストの追加を忘れているか、変更のごく一部のみの単体テストを行っている
コードの最大の匂いは、特に指示しない限り、AI がテストを書くことを完全に忘れてしまうということです。
AI エージェントにテストを書くように指示したとしても、何をテストするか、どのようにテストするかについて非常に具体的にしない限り、コードの一部しかカバーされないことがよくあることに私は繰り返し気づきました。
以下を追加する変更がある場合:
新しいメイン (トップレベル) React コンポーネント
メインコンポーネントが使用するその他の小さなコンポーネント
セベラ

l カスタムフック: コンポーネントで使用されます。
フックやコンポーネントが使用するいくつかの単純な純粋関数 (ビジネス ロジック、数学など)
通常、このようなケースでは、少なくとも一般的な幸せなパスと一部の主要な不幸なパスについては、メインコンポーネントをテストする統合テストから最大限の価値が得られます。次に、より小さなユニット (純粋な関数、またはカスタム フック) で複雑なロジックやエッジ ケースのロジックをテストします。
AI はメイン コンポーネントのテストを忘れて、純粋な関数かカスタム フックだけを単体テストすることが多いことがわかりました。
したがって、500 行の実装と 1000 行のテストがあり、一見すると十分にテストされていると思われるでしょう。
エッジ ケースやエラーをまだ処理していないコンポーネントのテストを追加するように AI に依頼した場合、AI がそれらのエッジ ケースやエラー パスの動作を確認するためのテストをわざわざ作成することはほとんどありません。
多くの場合、特に複雑なフロントエンド UI テストでは、AI は単に「ハッピー パス」 (成功状態) をテストし、エッジ ケースを回避します。
エッジケースや不幸なパスを処理する場合、実際には動作にバグがある場合でも、テストをパスさせることがよくあります。
AIレビューではテスト不足が指摘されないことが多い
AI を使用してプル リクエストのレビューを自動化する方法はたくさんあります。そして正直に言うと、彼らはかなり素晴らしいです。
ただし、AI のプルリクエスト レビューでは、機能が追加または変更される理由が理解できない場合がありますが、タイプミスやその他の間違い、パフォーマンスの問題などを見つけるのには非常に優れています。
しかし、すべての変更に対して良好なテストカバレッジを期待するという具体的な指示を与えない限り、何かがテストされていないことを認識して指摘することは非常にまれであることがわかりました。
また、不適切に書かれたテストを取り上げてコメントするのも見たことがありません。

(このページにリストされている問題など)。ほとんどの場合、大量の不適切に記述されたテストやテストが欠落している PR が承認されます。
AI に機能のテストを作成するよう依頼すると、AI は場合によってはやりすぎて、入力、プロパティ、状態の考えられるすべての組み合わせをテストしようとすることがあります。
私はテストが大好きで、コードを文書化したり、リファクタリングに役立ちます。しかし、過剰なテストが大きな問題であるという有力なケースがあります。
実際には大きな価値が付加されていないにもかかわらず、すべての組み合わせを過剰にテストしている場合は、テストをクリーンアップして、AI が書き込むジャンクを削除した方がよいでしょう。
私は、AI にテストを書かせて、その値が実際にどこから来るのかを判断し、あまり価値を追加していないテストを削除するように依頼することがよくありました。複雑なテストの数を大幅に削減できます。
オーバーテストを行う際に注意すべき点がいくつかあります。
不可能な状態のテスト (これらには価値がないため、回避する必要があります。テストを作成する人間は、これらの状態が不可能であるため、テストするのが無意味であることがよくわかりますが、AI は知りません)
実際のビジネス ロジックが関与しないゲッターやセッターなどの些細なものをテストする
各テストは、グループ化するのではなく、一度に 1 つの小さな事柄を主張します。
場合によっては、AI を使用して既存のコードを更新し、コードベースの他の場所で (間接的または直接的に) テストすることもあります。
AI が既にテストされていることに気付かず、既存の関数を更新してテストするように要求すると、新しい変更をテストするだけでなく、テストされていないと思われる既存のロジックに対してテストを再度 (新しいテスト ファイルに) 書き込むのを見てきました。
そして、私の経験では、これが起こったとき、それらのテストの品質は毎回非常に低く（単なる「AIのスロップ」）、機能が何であるのか、なぜ存在するのかを理解せず、100％のカバレッジを得ようとしているだけです。
(宣伝と少し重複しますが、

オーバーテストの危険点）。
前述の過剰テストの点に関連して、AI がすべての組み合わせをテストすることを決定した場合、1 つの入力と 1 つの期待されるアサーションを変更するだけで、テストのほぼ 90% をコピー/ペーストします。
.each() (Jest と Vitest の両方がサポートしています) を使用する方がはるかにクリーンです。
複数の非常に類似したテストのようなものを、テストに渡すデータの配列を含む 1 つのテストに変えることができます。
たとえば、 each() を使用して 1 つの関数でテストされる複数のケースを次に示します。
テスト。それぞれの`
ユーザー年齢 |アルコールを購入できます |期待されるメッセージ
${ 21 } | ${ true } | ${ '購入は承認されました' }
${ 18 } | ${ false } | ${ '年齢確認が必要です' }
${ 65 } | ${ true } | ${ 'シニア割引適用' }
${ 17 } | ${ false } | ${ '18 歳以上である必要があります' }
` (
'ユーザーが $userAge 歳の場合、「$expectedMessage」を表示します' ,
( { userAge , canPurchaseAlcohol , ExpectedMessage } ) => {
render (< CheckoutForm userAge = { userAge } /> ) ;
const submitButton = 画面 。 getByRole ( 'ボタン' , { name : '購入を完了' } ) ;
if (アルコールを購入できる) {
Expect (submitButton) 。 toBeEnabled() ;
} それ以外の場合は {
Expect (submitButton) 。 toBeDisabled() ;
}
Expect(screen .getByText(expectedMessage)) . toBeInTheDocument() ;
}
) ;
AIに実装を書いた後にテストを書くように依頼する
これについてはすでに触れましたが、AI に機能のテストの作成を依頼すると、実装が正しくテストされるようにデフォルトでテストが作成されます。
これは非常に危険であり、テストによってアプリケーションが動作するという自信が得られると考えてしまいますが、それは誤った期待です。
AI が作成したテストには注意が必要です。ロジックにバグがある場合、テストは正しく実装されているかのようにそれをアサートするだけです。
代わりに、実装と書き込みテストが行​​われます

それは合格です。
(これらすべての問題と同様、SKILL.md を適切に使用するよう注意深く促しれば、これも回避できます)。
注: タイプミスや間違ったコード構文を見つけるのに優れています。しかし、機能が存在する理由を理解できず、ビジネス ロジックのバグがある場合には機能が失われることがよくあります。
AI が作成したテストの最大の問題の 1 つは、ほとんどのテストが単体テストとして扱われ、単にモック ( jest.mock() または vi.mock() を使用) されることです。
これは AI がテストに合格するのに非常に役立ちますが、実際の動作ではなくモックされた実装をテストするだけになるため、AI の価値の多くが失われることがよくあります。
依存関係注入 (DI) システムをセットアップしている場合は、モック オブジェクトを渡すためにその DI システムを使用する必要があります。 JavaScript/TypeScript アプリの AI は通常、これを無視し、単に jest.mock を使用します。
長期的なメンテナンスに関してこれに関して私が抱えている主な問題は、mock() を使用すると、モジュール全体 (ファイル/インポート) をモックすることになるということです。したがって、エクスポートされた 1 つの関数の一部のデータを偽装する必要がある場合は、モジュール全体をモックする必要があります。また、AI はテストに合格できない場合にただモックをするだけであるため、多くの場合、システムの設計が間違っていたことを意味し、モックは脱出用のハッチであり、AI が最初に試みるべきではないこともわかりました。
前述のオーバーモックの点に関連して、AI が既存のコード (作成したばかりのコード) のテストを作成する場合、通常はテストが合格するようにテストを試みます。
前のセクションで述べたように、多くの場合、問題を引き起こすものをすべてモックすることが近道になります。
私は AI に既存のコード (テストのないコード) のテストを書くように依頼してきましたが、AI がバグを見つけて修正したことはまだ見たことがありません。テストのない既存のコードのテストを作成した場合、バグ (負の数などのエッジ ケースのバグ) が見つかり始めるのはよくあることです。

s、空の配列など）。
AI は文字通り実装を見てコードが正しいと想定し、テストではその実装をテストする必要があります。実際には、関数が何をしようとしているのかを理解するよう AI に依頼し、テストを作成する必要があります。
上記と同様に、AI が何かのモックを提供する必要があると判断した場合、モジュール全体をモックすることがわかりました。
より良い方法は、 spyOn() を使用することです。
これにより、より詳細な制御が可能になり、モジュール全体 (ファイル全体) をモックすることはなくなります。一般に、長期的には spyOn を維持する方が、mock を過剰に使用するよりもはるかに簡単であることがわかります。
// ❌ 悪い例: モジュール全体をモック化する
vi 。モック ( '../services/logger' ) ;
test ( 'API が失敗した場合のエラーをログに記録' , async ( ) => {
// ...
} ) ;
// ✅ 良い点: spyOn を使用して必要なものだけをモックする
import * '../services/logger' からロガーとしてインポートします。
test ( 'API が失敗した場合のエラーをログに記録' , async ( ) => {
const logErrorSpy = vi 。 spyOn ( ロガー 、 'logError' ) 。モック実装 ( ( ) => { } ) ;
render (< UserProfile userId = " 123 " /> ) ;
await waitFor ( ( ) => {
期待 (logErrorSpy) 。 toHaveBeenCalledWith ( 'ユーザーの取得に失敗しました' ) ;
} ) ;
} ) ;
フィクスチャを使用しない（引数/小道具として渡すために準備されたモックデータ）
フィクスト

[切り捨てられた]

## Original Extract

AI is very powerful, but it can introduce lots of subtle code smells when it generates code for your frontend applications. Here are the main ones to look out for.

Code Smells when you get AI to write your Frontend Tests | How To Test Frontend HowToTest FrontEnd
Search Home Courses About Blog Recent blog posts
Code Smells in your FE AI Generated Tests
Why you should use Faker in your tests
Reduce CI/CD times by 50% with these tools
What makes a good frontend test?
Guide to Vitest Browser Mode - and why you should think about using it
Testing for accessibility in your Frontend application
Testing async RSC components in Vitest
Add type checks to your Vitest tests
Everything you need to know about act() in React tests
How to test React Server Components (RSC) in Vitest/Jest
Reduce CI/CD times by 50% with these tools
Writing FE tests - 35+ tips and tricks
Add type checks to your Vitest tests
React Testing Course (React Testing Library)
React Testing Library DOM Matchers
Vitest/Jest Testing Challenges
Code Smells when you get AI to write your Frontend Tests
Forgetting to add any tests, or only unit testing tiny portions of its changes
AI reviews often don't point out lack of tests
Asking AI to write tests after it wrote the implementation
Not using fixtures (prepared mock data to pass in as args/props)
Not knowing your existing helper functions
AI loves Regex when matching rendered elements
Testing for exact DOM structure
Adding extra timeouts (especially to e2e tests)
Testing your test, not your actual implementation
AI's use of beforeAll() or beforeEach()
In the last year or two, the software industry has changed so much. Thanks to AI tools such as Claude Code, I now manually type only a tiny percentage of the code I write compared to a year ago.
My day to day work is with React, testing with Vitest, Playwright etc. Pretty standard setup that most of you are probably also using. Tests are of course an important part of delivering software.
And AI can definitely write code and add tests for that code .
But I've noticed many patterns that AI constantly introduces which I would consider code smells and generally try to avoid.
I wanted to document some of the testing code smells I keep seeing that AI introduces.
They often help get tests passing quickly, but they can also create maintenance problems that only become apparent months later.
I bet as you read this, you will disagree with some of these points.
Especially if you have already spent the time setting up good AGENTS.md (or CLAUDE.md ), with good SKILL.md files covering your conventions.
I'm writing this up to point out the sort of issues that AI can introduce in tests when you don't have any reins to nudge it into the right direction.
The real solution is to refine your AGENTS.md , CLAUDE.md , and SKILL.md files.
I will have an upcoming blog post with the specific skills/instructions I use to nudge Claude Code into writing much better and maintainable tests.
A lot of these code smells contradict each other (e.g. over-testing and then under-testing!). These common code smells don't all appear at the same time!
A lot of these best practices are similar to what I've covered elsewhere on this site, as they're generally just best practices.
Forgetting to add any tests, or only unit testing tiny portions of its changes
The biggest code smell is that AI will completely forget to write tests, unless you specifically tell it.
Even if you tell AI agents to write tests, I've repeatedly found that unless you are very specific about what to test and how to test it , it will often cover only part of the code.
If you have a change that adds:
a new main (top level) React component
some other smaller components that the main component uses
several custom hooks , which are used in your components
and a few simple pure functions (such as some business logic, maths, etc) that your hooks and/or components use
Typically in those cases you will get most value out of an integration test, testing the main component - at least for the general happy path and some main unhappy paths. Then smaller unit tests for complex or edge case logic in the smaller units (the pure functions or maybe the custom hooks).
I've found that AI will often forget to test the main component , and just unit test the hell out of the pure functions or maybe the custom hooks.
So you might have 500 lines of implementation, and 1000 lines of tests and at first glance you would assume it is well tested.
If you ask AI to add tests for a component that doesn't handle edge cases or errors already, then it is very rare to see AI bother to write tests to confirm behaviour of those edge case or error paths.
Often AI, especially on complex frontend UI tests, will just test the 'happy path' (the success state), and avoid edge cases.
If it does handle edge cases and unhappy paths, it will often make the tests pass even when the behaviour is actually buggy.
AI reviews often don't point out lack of tests
There are lots of ways to automate pull request reviews with AI. And honestly they are pretty amazing.
But AI pull-request reviews can sometimes lack an understanding of why a feature is being added or modified - but they are really great at spotting typos and other mistakes, performance issues, and so on.
But I have found that unless you give it specific instructions that you expect good test coverage for all changes, it is quite rare for it to pick up and point out that something is not tested.
I've also never seen it pick up and comment on badly written tests (such as the issues listed on this page). It will almost always approve a PR with tons of badly written tests or missing tests.
Ask AI to write tests for a feature, and it will sometimes go overboard , attempting to test every possible combination of inputs, props, and states.
I love testing, it documents code, it helps with refactors. But there is a strong case that over testing is a huge problem .
If you are over testing every single combination, when really they are not adding much value, it is better to clean up tests and remove the junk that AI writes.
I've often gotten AI to write tests and then ask it to work out where the value actually comes from, and to delete tests that are not adding much value. It can significantly reduce the number of complex tests.
Some things to look for when over testing :
testing for impossible states (they have no value and should be avoided. Humans writing tests will often know these states are impossible and therefore pointless to test for, but AI doesn't)
testing trivial things like getters or setters, where no real business logic is involved
Each test asserting one minor thing at a time, instead of grouping them together
Sometimes you will use AI to update existing code, which may be tested (indirectly or directly) elsewhere in the codebase.
I've seen AI not notice that it is already tested, and when updating an existing function and asking it to test it, it will test not only your new changes but write tests again (in a new test file) for the existing logic that it believes is not tested .
And in my experience when this does happen, the quality of those tests is very low (just 'AI slop') every single time, not understanding what or why the function exists and it is just trying to get 100% coverage.
(This overlaps a bit with a previous point of over testing).
Related to the previous point of over testing, when AI decides to test every combination it will almost copy/paste 90% of a test to just change one input and one expected assertion .
It is much cleaner to get it to use .each() (which Jest and Vitest both support)
It can turn something like multiple very similar tests into just one test, with an array of data to pass into the test.
For example, here are multiple cases that are tested in one function, using each() :
test . each `
userAge | canPurchaseAlcohol | expectedMessage
${ 21 } | ${ true } | ${ 'Purchase approved' }
${ 18 } | ${ false } | ${ 'Age verification required' }
${ 65 } | ${ true } | ${ 'Senior discount applied' }
${ 17 } | ${ false } | ${ 'Must be 18 or older' }
` (
'displays "$expectedMessage" when user is $userAge years old' ,
( { userAge , canPurchaseAlcohol , expectedMessage } ) => {
render ( < CheckoutForm userAge = { userAge } /> ) ;
const submitButton = screen . getByRole ( 'button' , { name : 'Complete Purchase' } ) ;
if ( canPurchaseAlcohol ) {
expect ( submitButton ) . toBeEnabled ( ) ;
} else {
expect ( submitButton ) . toBeDisabled ( ) ;
}
expect ( screen . getByText ( expectedMessage ) ) . toBeInTheDocument ( ) ;
}
) ;
Asking AI to write tests after it wrote the implementation
I've touched on this already but whenever you ask AI to write tests for a feature, it will by default write tests so your implementation is correctly tested .
This can be quite dangerous and leave you thinking your tests give you confidence in a working application, but it is false hope.
You have to be careful with the tests that AI wrote, as if there is a bug in the logic the test will just assert it as if it was correctly implemented.
Instead it will take the implementation and write tests that pass.
(Like any and all of these issues, with some careful prompting of good usage of SKILL.md then this can be avoided).
Note: It is great at finding typos or incorrect code syntax. But it fails to understand why a feature exists and often missing when there is a business logic bug.
One of the biggest annoyances of AI written tests is that it will treat most tests as unit tests, and just mock ( with jest.mock() or vi.mock() ).
This really helps AI get tests passing, but it often removes much of their value because they end up just testing mocked implementations instead of real behaviour.
If you have a dependency injection (DI) system set up, then it is the DI system that should be used to pass in mock objects. AI on JavaScript/TypeScript apps will generally ignore this and just go for jest.mock .
The main issue I have with this for long term maintenance is that when you use mock() you are mocking an entire module (file/import). So if you need to fake some data for one exported function, the entire module has to be mocked. I also find that because AI will just mock when it cannot get a test to pass, it often means there was some badly designed system and the mock is an escape hatch that should not be the first thing it tries.
Related to the previous point of over mocking, when AI writes tests for existing code (code that it may have just written) it will generally try to test so the tests pass.
As said in the previous section often the shortcut will be to mock anything that causes issues.
I have often asked AI to write tests for existing code (that had no tests), and I have yet to see it ever find and fix any bugs. I know that when I've written tests for existing code that had no tests, it is very common to start to find bugs (edge case bugs, like negative numbers, empty arrays, etc).
AI will literally just look at the implementation and assume that code is correct, and the tests must test that implementation. Really we need to be asking AI to try to understand what the function is intending on doing, and write those tests.
Similar to above - if AI decides it has to provide a mock for something, then I find that it will mock an entire module.
The better way is often to use spyOn() .
This will give you much more control, and doesn't mock the entire module (entire file). Generally I find that maintaining spyOn is much easier long term than over use of mock .
// ❌ Bad: Mocking the entire module
vi . mock ( '../services/logger' ) ;
test ( 'logs error when API fails' , async ( ) => {
// ...
} ) ;
// ✅ Good: Using spyOn to mock just what you need
import * as logger from '../services/logger' ;
test ( 'logs error when API fails' , async ( ) => {
const logErrorSpy = vi . spyOn ( logger , 'logError' ) . mockImplementation ( ( ) => { } ) ;
render ( < UserProfile userId = " 123 " /> ) ;
await waitFor ( ( ) => {
expect ( logErrorSpy ) . toHaveBeenCalledWith ( 'Failed to fetch user' ) ;
} ) ;
} ) ;
Not using fixtures (prepared mock data to pass in as args/props)
Fixt

[truncated]
