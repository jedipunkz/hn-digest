---
source: "https://coolhand-labs.github.io/squishling/"
hn_url: "https://news.ycombinator.com/item?id=50295504"
title: "Squishling 0.5.0 – hybrid LLM/Ruby code classes"
article_title: "Squishling: write Ruby only where it's cheaper than the tokens"
image: "https://coolhand-labs.github.io/squishling/img/og-cover.png"
author: "mikecarroll"
captured_at: "2026-10-29T21:34:24Z"
capture_tool: "hn-digest"
hn_id: 50295504
score: 2
comments: 0
posted_at: "2026-10-29T20:35:23Z"
tags:
  - hacker-news
---

# Squishling 0.5.0 – hybrid LLM/Ruby code classes

- HN: [50295504](https://news.ycombinator.com/item?id=50295504)
- Source: [coolhand-labs.github.io](https://coolhand-labs.github.io/squishling/)
- Score: 2
- Comments: 0
- Posted: 2026-10-29T20:35:23Z

## Translation

Title: Squishling 0.5.0 – hybrid LLM/Ruby code classes
Article title: Squishling: write Ruby only where it's cheaper than the tokens
Description: Don

Article text:
Squishling · a Ruby gem for elastic software
Don't just use AI to write code.
Use AI to not need code.
Squishling lets a class run as ordinary Ruby or hand the work to an LLM, and return the same checked,
typed answer either way.
Write Ruby only where it's cheaper than the tokens.
Join the discussion on Hacker News .
Latest release · v0.5.0 · October 28, 2026
Squishling turns any Ruby class into a hybrid of code and LLM inference working together. v0.5.0 is about making sure
it doesn't get too magical in production.
💰 Token budget caps: limit what one call can spend across every retry, escalation and judge
✅ Squishling.verify : checks every squished class (schemas, models, params) in CI without a single LLM request
🔍 Clearer errors: tell when the model was cut off by its token limit versus actually answering wrong
🧪 A scripted test adapter, so your unit tests never hit a provider
🔌 RubyLLM is now an optional adapter: add gem "ruby_llm" to your Gemfile, or use any OpenAI-compatible endpoint or your own
With Squishling, every Ruby class can land somewhere between "an LLM handles it" and "code handles it."
Now you have an easy way to define, and adjust, that balance based off of the needs of a class, your velocity, or
how much tech debt your team is able to carry.
Drag the slider to see which tradeoff makes the most sense for the class you are building.
← Slide left to find the optimal spot
Fully deterministic, cheap to scale. Logical complexity and failure states often balloon with features.
Tech debt is leverage, if you can price it
Debt isn't always bad. Used well, it gives you leverage, and tech debt works the same way: it can be an asset or a
liability, depending on what you do with it.
The catch is that nobody can price it. It's hard to quantify the advantage tech debt gives you if you can't
quantify what it costs . Here's a common place where it quietly gets expensive.
Some requirements are too expensive to maintain in code
Say you're in health tech. You sign a clinical client and need their patient data in your system. There's a
standard for that, FHIR, so it looks easy.
def call_parser(data)
fhir_parser.call(data)
end
A hundred clients later
def call_parser(client_name, data)
case client_name
when client_a then client_a_parser.call(data)
when client_b then client_b_parser.call(data)
when client_c then client_c_parser.call(data)
when client_d then client_d_parser.call(data)
when client_e then client_e_parser.call(data)
when client_f then client_f_parser.call(data)
when client_g then client_g_parser.call(data)
when client_h then client_h_parser.call(data)
# ... ad nauseam
No matter whether people or AI wrote it, this is terrible architecture:
Any change to a data format can break it, and every breakage is a scramble.
It adds cost to every new contract. The more clients you sell, the worse it gets.
The worst kind of debt: uncontrollable, unavoidable and unmanageable ... and "better" code can't fix it.
Same job. Different implementation.
class Parser
include Squishling
purpose "Please parse this data into a patient record."
output_schema PatientRecord
# def call(data)
# fhir_parser.call(data)
# end
end
Let the model take the first version, and look at what you get. Flexible by default, and it bends instead of
breaking.
Nothing to break when a format shifts.
A new client works on day one.
It copes with external factors you can't change.
The objection is cost: every request spends tokens. The answer:
Refactor the high-volume bits to be deterministic.
Now your tech debt has a quantity you can measure .
Define the goal and the shape of the answer
Make a class squish. Give it a purpose and an output schema, and Squishling does the rest. This class has no
implementation at all, so every call is squished.
Squishling.configure do |config|
config.default_model = "claude-haiku-4-5-20251001"
# falls back to the adapter's default when nil
end
class Parser
include Squishling
purpose "Map the inbound record onto the schema..."
output_schema do
string :patient_name
string :patient_gender, enum: %w[male female trans unknown]
array :diagnosis_codes, of: :string
array :symptoms, of: :string
end
# no call method: every invocation is squished.
end
Run against 100 synthetic patient records (generated with Synthea):
Every diagnosis code was right. The misses clustered in three places: names written HL7-style, gender codes, and
symptoms the model invented.
class Parser
include Squishling
purpose <<~PURPOSE
Map the record onto the schema below.
Name: output "First Last" order, even if the source writes
it as HL7's LAST^FIRST, often in all caps.
Gender: translate to "male", "female", "trans", or "unknown".
"unknown" for any code you don't recognize, including HL7's
bare O / U / A / N letter codes. Never guess.
Symptoms: only from an explicit symptom field. A diagnosis is
not a symptom. Never infer one from a diagnosis name.
PURPOSE
...
end
95 / 100 pass
$0.51 per 100 calls
The two systematic failures that remain are HL7 letter codes and unlabeled SNOMED gender codes that need a
translation table, plus one dictation transcript with spoken digits.
Best practice: describe each field in the schema
This run kept the per-field rules in the purpose. A better home for them is a description: on
the field itself, as in this equivalent shape of the same class. The description travels with the field in the
schema the model fills in, so the guidance sits right where the value goes, and the purpose stays short and about
the goal. It also lives next to the type and enum it explains, so the two don't drift apart.
purpose "Map the record onto the schema."
output_schema do
string :patient_name,
description: "First Last order, even if the source writes it as HL7's LAST^FIRST, often in all caps."
string :patient_gender, enum: %w[male female trans unknown],
description: "\"unknown\" for any code you don't recognize, including HL7's bare O / U / A / N. Never guess."
array :symptoms, of: :string,
description: "Only from an explicit symptom field. A diagnosis is not a symptom."
end
The 95 / 100 above was measured with the rules in the purpose; this description version wasn't benchmarked.
Descriptions are guidance to the model, not enforcement: for rules that must hold, use an enum or squish_validate .
Use a harness to pay for higher accuracy
Squishsum compares two Squishling results. It asks the model twice, at the same time, and accepts
the answer only when both agree. That costs about twice the requests, so reach for it when a confidently wrong
answer costs more than a second one.
class Parser
include Squishling
...
squishling harness: :squishsum
end
When the two results match, you get the first one. When they differ, Squishling raises
Squishling::DisagreementError , and squish_fallback decides what happens next. By default the
results must match exactly; pass compare: to say which fields have to agree.
Add a judge with judged_squishsum
When the two results differ, a judge looks at both and picks one, or rejects both. The judge can be a bigger chat
model, as here, or a fast decision model such as Jev.
class Parser
include Squishling
...
squishling model: "claude-haiku-4-5",
harness: { type: :judged_squishsum, judge: "claude-opus-5-5" }
# two Haiku samples; Opus judges when they differ
end
Run against the same 100 records:
Four of the five misses are now right, for double the cost of the tuned purpose. The one that remains is a
reminder that a second opinion helps most with random mistakes, less with systematic misreadings.
Validation checks the shape of an answer, not that it's true.
Register a harness under a name and use it like a built-in. It decides how to use the escalation, and gets
samples, judges, squawk reports and squish_fallback for free. This one asks two models, and
calls in a third to break a tie:
class Tiebreaker < Squishling::Harness::Base
option :tiebreaker, :step, required: true
def call(run, tiebreaker:)
# Ask the first two models at the same time.
first, second = run.samples(run.stages.first(2))
# You decide what "agree" means.
return first if run.agree?(first, second, ->(a, b, **) { a.priority == b.priority })
# They disagree: ask the strongest model.
run.escalate(tiebreaker)
end
end
Squishling.register_harness(:tiebreaker, Tiebreaker)
class TicketTriager
include Squishling
squishling escalation: [
{ model: "claude-haiku-4-5", priority: 10 },
{ model: "gemma-4b", priority: 10 }
],
harness: { type: :tiebreaker, tiebreaker: "claude-opus-5-5" }
purpose "Assign a priority and a team."
output_schema do
string :priority, enum: %w[low medium high]
string :team
end
end
Every primitive returns the same validated, typed result, so you can compare fields directly.
Options are checked when you declare the harness, before any request is made.
Use squish_validate for rules the schema can't express. A failed check rejects that sample or attempt.
A harness that raises DisagreementError or InvalidOutputError goes to the method's
squish_fallback , like the built-ins.
The built-in options, and the rest of the plugin point, are covered in the
harness docs .
How · 4. Harden the common shapes
Code the standard shapes. Squish the rest.
A quick look at the data shows 63% of the requests fit a standard format . So route those to the
parser you already trust, and squish the rest:
class Parser
include Squishling
...
squishling model: "claude-haiku-4-5",
harness: { type: :judged_squishsum, judge: "claude-opus-5-5" }
def call(data)
return fhir_parser.call(data) if fhir_message?(data)
squish!
end
end
95 / 100 pass
$0.12 per 100 calls
Same 95 out of 100, the same five misses, and 76% cheaper than the tuned-purpose run that squishes every request. The router moved this integration from fully
flexible to mostly efficient without a rewrite, and without betting up front.
Results from the talk's demo, using 100 synthetic records and Anthropic Claude models. They show the
shape of the trade, not a benchmark of your data.
Day-one integrations at a fraction of the cost
Total LLM cost well below maintenance cost.
And you can keep revisiting this balance as the business evolves.
Drowning under a thousand legacy parsers? Find the ones at the tail of your usage curve. Squish them instead.
Then delete the code. Stop living in fear that a tiny client finds the breaking change you shipped two quarters ago.
The steps above show what a squishling does. These three principles decide where one belongs in your system.
One small job, one clear goal, one output shape, like the
Parser .
Start with the cheapest model that reliably does the job, and tune before reaching for more.
A tuned purpose on Haiku matched a second sample and a judge at half the cost.
Code where code is cheap, squishlings where it isn't, as in
code the 63%, squish the rest .
A reliable system isn't one with no squish in it. It's one where every squish is small, atomic and
easy to replace .
When a service call fails, squish it
Not parsing health data? Anything that talks to another service has the same shape. Keep your Ruby parser for well-formed payloads, and hand the rest to the
model. Save yourself, and the engineers who come after you, uncountable headaches.
class WebhookHandler
include Squishling
...
def call(service:, body:)
VendorWebhooks.fetch(service).parse(body)
rescue VendorWebhooks::ParseError => e
squish!(context: { parse_error: e })
end
end
With squish! , the model gets the call's inputs plus whatever context you pass, such as the error, and
still has to return the same validated result. If it can't, squish_fallback decides what happens next.
Even a bug in your own code can squish
Any rescued exception can be the trigger, including a plain old bug:
class SomeHelperMethod
def call(payload)
payload.dig("user", "email")
end
end
payload = { "user" => "alice@example.com" }
# oops: the caller sent a string, not a hash
SomeHelperMethod.new.call(payload)
# => TypeError: String does not have #dig method
And after:
class SomeHelperMethod
include Squishling
purpose "Return

[truncated]

## Original Extract

Don

Squishling · a Ruby gem for elastic software
Don't just use AI to write code.
Use AI to not need code.
Squishling lets a class run as ordinary Ruby or hand the work to an LLM, and return the same checked,
typed answer either way.
Write Ruby only where it's cheaper than the tokens.
Join the discussion on Hacker News .
Latest release · v0.5.0 · October 28, 2026
Squishling turns any Ruby class into a hybrid of code and LLM inference working together. v0.5.0 is about making sure
it doesn't get too magical in production.
💰 Token budget caps: limit what one call can spend across every retry, escalation and judge
✅ Squishling.verify : checks every squished class (schemas, models, params) in CI without a single LLM request
🔍 Clearer errors: tell when the model was cut off by its token limit versus actually answering wrong
🧪 A scripted test adapter, so your unit tests never hit a provider
🔌 RubyLLM is now an optional adapter: add gem "ruby_llm" to your Gemfile, or use any OpenAI-compatible endpoint or your own
With Squishling, every Ruby class can land somewhere between "an LLM handles it" and "code handles it."
Now you have an easy way to define, and adjust, that balance based off of the needs of a class, your velocity, or
how much tech debt your team is able to carry.
Drag the slider to see which tradeoff makes the most sense for the class you are building.
← Slide left to find the optimal spot
Fully deterministic, cheap to scale. Logical complexity and failure states often balloon with features.
Tech debt is leverage, if you can price it
Debt isn't always bad. Used well, it gives you leverage, and tech debt works the same way: it can be an asset or a
liability, depending on what you do with it.
The catch is that nobody can price it. It's hard to quantify the advantage tech debt gives you if you can't
quantify what it costs . Here's a common place where it quietly gets expensive.
Some requirements are too expensive to maintain in code
Say you're in health tech. You sign a clinical client and need their patient data in your system. There's a
standard for that, FHIR, so it looks easy.
def call_parser(data)
fhir_parser.call(data)
end
A hundred clients later
def call_parser(client_name, data)
case client_name
when client_a then client_a_parser.call(data)
when client_b then client_b_parser.call(data)
when client_c then client_c_parser.call(data)
when client_d then client_d_parser.call(data)
when client_e then client_e_parser.call(data)
when client_f then client_f_parser.call(data)
when client_g then client_g_parser.call(data)
when client_h then client_h_parser.call(data)
# ... ad nauseam
No matter whether people or AI wrote it, this is terrible architecture:
Any change to a data format can break it, and every breakage is a scramble.
It adds cost to every new contract. The more clients you sell, the worse it gets.
The worst kind of debt: uncontrollable, unavoidable and unmanageable ... and "better" code can't fix it.
Same job. Different implementation.
class Parser
include Squishling
purpose "Please parse this data into a patient record."
output_schema PatientRecord
# def call(data)
# fhir_parser.call(data)
# end
end
Let the model take the first version, and look at what you get. Flexible by default, and it bends instead of
breaking.
Nothing to break when a format shifts.
A new client works on day one.
It copes with external factors you can't change.
The objection is cost: every request spends tokens. The answer:
Refactor the high-volume bits to be deterministic.
Now your tech debt has a quantity you can measure .
Define the goal and the shape of the answer
Make a class squish. Give it a purpose and an output schema, and Squishling does the rest. This class has no
implementation at all, so every call is squished.
Squishling.configure do |config|
config.default_model = "claude-haiku-4-5-20251001"
# falls back to the adapter's default when nil
end
class Parser
include Squishling
purpose "Map the inbound record onto the schema..."
output_schema do
string :patient_name
string :patient_gender, enum: %w[male female trans unknown]
array :diagnosis_codes, of: :string
array :symptoms, of: :string
end
# no call method: every invocation is squished.
end
Run against 100 synthetic patient records (generated with Synthea):
Every diagnosis code was right. The misses clustered in three places: names written HL7-style, gender codes, and
symptoms the model invented.
class Parser
include Squishling
purpose <<~PURPOSE
Map the record onto the schema below.
Name: output "First Last" order, even if the source writes
it as HL7's LAST^FIRST, often in all caps.
Gender: translate to "male", "female", "trans", or "unknown".
"unknown" for any code you don't recognize, including HL7's
bare O / U / A / N letter codes. Never guess.
Symptoms: only from an explicit symptom field. A diagnosis is
not a symptom. Never infer one from a diagnosis name.
PURPOSE
...
end
95 / 100 pass
$0.51 per 100 calls
The two systematic failures that remain are HL7 letter codes and unlabeled SNOMED gender codes that need a
translation table, plus one dictation transcript with spoken digits.
Best practice: describe each field in the schema
This run kept the per-field rules in the purpose. A better home for them is a description: on
the field itself, as in this equivalent shape of the same class. The description travels with the field in the
schema the model fills in, so the guidance sits right where the value goes, and the purpose stays short and about
the goal. It also lives next to the type and enum it explains, so the two don't drift apart.
purpose "Map the record onto the schema."
output_schema do
string :patient_name,
description: "First Last order, even if the source writes it as HL7's LAST^FIRST, often in all caps."
string :patient_gender, enum: %w[male female trans unknown],
description: "\"unknown\" for any code you don't recognize, including HL7's bare O / U / A / N. Never guess."
array :symptoms, of: :string,
description: "Only from an explicit symptom field. A diagnosis is not a symptom."
end
The 95 / 100 above was measured with the rules in the purpose; this description version wasn't benchmarked.
Descriptions are guidance to the model, not enforcement: for rules that must hold, use an enum or squish_validate .
Use a harness to pay for higher accuracy
Squishsum compares two Squishling results. It asks the model twice, at the same time, and accepts
the answer only when both agree. That costs about twice the requests, so reach for it when a confidently wrong
answer costs more than a second one.
class Parser
include Squishling
...
squishling harness: :squishsum
end
When the two results match, you get the first one. When they differ, Squishling raises
Squishling::DisagreementError , and squish_fallback decides what happens next. By default the
results must match exactly; pass compare: to say which fields have to agree.
Add a judge with judged_squishsum
When the two results differ, a judge looks at both and picks one, or rejects both. The judge can be a bigger chat
model, as here, or a fast decision model such as Jev.
class Parser
include Squishling
...
squishling model: "claude-haiku-4-5",
harness: { type: :judged_squishsum, judge: "claude-opus-5-5" }
# two Haiku samples; Opus judges when they differ
end
Run against the same 100 records:
Four of the five misses are now right, for double the cost of the tuned purpose. The one that remains is a
reminder that a second opinion helps most with random mistakes, less with systematic misreadings.
Validation checks the shape of an answer, not that it's true.
Register a harness under a name and use it like a built-in. It decides how to use the escalation, and gets
samples, judges, squawk reports and squish_fallback for free. This one asks two models, and
calls in a third to break a tie:
class Tiebreaker < Squishling::Harness::Base
option :tiebreaker, :step, required: true
def call(run, tiebreaker:)
# Ask the first two models at the same time.
first, second = run.samples(run.stages.first(2))
# You decide what "agree" means.
return first if run.agree?(first, second, ->(a, b, **) { a.priority == b.priority })
# They disagree: ask the strongest model.
run.escalate(tiebreaker)
end
end
Squishling.register_harness(:tiebreaker, Tiebreaker)
class TicketTriager
include Squishling
squishling escalation: [
{ model: "claude-haiku-4-5", priority: 10 },
{ model: "gemma-4b", priority: 10 }
],
harness: { type: :tiebreaker, tiebreaker: "claude-opus-5-5" }
purpose "Assign a priority and a team."
output_schema do
string :priority, enum: %w[low medium high]
string :team
end
end
Every primitive returns the same validated, typed result, so you can compare fields directly.
Options are checked when you declare the harness, before any request is made.
Use squish_validate for rules the schema can't express. A failed check rejects that sample or attempt.
A harness that raises DisagreementError or InvalidOutputError goes to the method's
squish_fallback , like the built-ins.
The built-in options, and the rest of the plugin point, are covered in the
harness docs .
How · 4. Harden the common shapes
Code the standard shapes. Squish the rest.
A quick look at the data shows 63% of the requests fit a standard format . So route those to the
parser you already trust, and squish the rest:
class Parser
include Squishling
...
squishling model: "claude-haiku-4-5",
harness: { type: :judged_squishsum, judge: "claude-opus-5-5" }
def call(data)
return fhir_parser.call(data) if fhir_message?(data)
squish!
end
end
95 / 100 pass
$0.12 per 100 calls
Same 95 out of 100, the same five misses, and 76% cheaper than the tuned-purpose run that squishes every request. The router moved this integration from fully
flexible to mostly efficient without a rewrite, and without betting up front.
Results from the talk's demo, using 100 synthetic records and Anthropic Claude models. They show the
shape of the trade, not a benchmark of your data.
Day-one integrations at a fraction of the cost
Total LLM cost well below maintenance cost.
And you can keep revisiting this balance as the business evolves.
Drowning under a thousand legacy parsers? Find the ones at the tail of your usage curve. Squish them instead.
Then delete the code. Stop living in fear that a tiny client finds the breaking change you shipped two quarters ago.
The steps above show what a squishling does. These three principles decide where one belongs in your system.
One small job, one clear goal, one output shape, like the
Parser .
Start with the cheapest model that reliably does the job, and tune before reaching for more.
A tuned purpose on Haiku matched a second sample and a judge at half the cost.
Code where code is cheap, squishlings where it isn't, as in
code the 63%, squish the rest .
A reliable system isn't one with no squish in it. It's one where every squish is small, atomic and
easy to replace .
When a service call fails, squish it
Not parsing health data? Anything that talks to another service has the same shape. Keep your Ruby parser for well-formed payloads, and hand the rest to the
model. Save yourself, and the engineers who come after you, uncountable headaches.
class WebhookHandler
include Squishling
...
def call(service:, body:)
VendorWebhooks.fetch(service).parse(body)
rescue VendorWebhooks::ParseError => e
squish!(context: { parse_error: e })
end
end
With squish! , the model gets the call's inputs plus whatever context you pass, such as the error, and
still has to return the same validated result. If it can't, squish_fallback decides what happens next.
Even a bug in your own code can squish
Any rescued exception can be the trigger, including a plain old bug:
class SomeHelperMethod
def call(payload)
payload.dig("user", "email")
end
end
payload = { "user" => "alice@example.com" }
# oops: the caller sent a string, not a hash
SomeHelperMethod.new.call(payload)
# => TypeError: String does not have #dig method
And after:
class SomeHelperMethod
include Squishling
purpose "Return

[truncated]
