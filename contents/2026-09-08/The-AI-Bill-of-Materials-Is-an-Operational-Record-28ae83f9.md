---
source: "https://jasondoyle.ie/whitepapers/the-ai-bill-of-materials-is-an-operational-record/"
hn_url: "https://news.ycombinator.com/item?id=49615861"
title: "The AI Bill of Materials Is an Operational Record"
article_title: "The AI Bill of Materials Is an Operational Record | Jason Doyle"
image: "https://jasondoyle.ie/assets/jason-doyle-og.png"
author: "jamesblakes"
captured_at: "2026-09-08T19:45:28Z"
capture_tool: "hn-digest"
hn_id: 49615861
score: 3
comments: 0
posted_at: "2026-09-08T19:42:32Z"
tags:
  - hacker-news
---

# The AI Bill of Materials Is an Operational Record

- HN: [49615861](https://news.ycombinator.com/item?id=49615861)
- Source: [jasondoyle.ie](https://jasondoyle.ie/whitepapers/the-ai-bill-of-materials-is-an-operational-record/)
- Score: 3
- Comments: 0
- Posted: 2026-09-08T19:42:32Z

## Translation

Title: The AI Bill of Materials Is an Operational Record
Article title: The AI Bill of Materials Is an Operational Record | Jason Doyle
Description: Why an AI Bill of Materials must connect models, data, software, configuration, evidence and deployments to support security, change and incident response.

Article text:
Skip to content
Jason Doyle
Open navigation
Home
Whitepapers
Contact
AI supply chain / Operational assurance
The AI Bill of Materials Is an Operational Record
What an inventory must show before it can support security, change and incident response.
2. What an AI Bill of Materials means
5. Different entries carry different evidence
Verifiable component facts
6. Bind the record to a deployed release
7. Four connected records
1. Component graph
8. What the record should support
Vulnerability response
Dataset withdrawal and correction
Procurement and supplier review
Regulation and technical documentation
9. Failure modes
Generated once at procurement
Accurate components, missing relationships
Stable name, changing implementation
Supplier claim presented as verification
Dynamic evidence copied into a static file
Inventory without response tooling
10. A practical operating model
Step 1: Define the questions
Step 2: Set the system boundary
Step 3: Start with the ordinary SBOM
Step 4: Add AI-specific components
Step 7: Bind every production release
Step 9: Connect operational systems
11. What to measure
Deployment binding coverage
12. Counterarguments
Model and system cards already provide this information
Hosted models make complete inventory impossible
The system changes too quickly
Automated generation solves the problem
A detailed BOM creates an attack map
Compliance programmes will define the required fields
13. What this paper does not claim
Appendix A: Illustrative CycloneDX JSON
Disclosure: These views are my own and do not represent my current or any former employers. This paper uses only public sources and does not describe non-public product information.
An AI system can change without its application code changing.
A model provider can retire a version or move a deployment to a replacement.
An embedding model can change. A retrieval index can be rebuilt. A prompt can be
edited. A fine-tuning dataset can be withdrawn. An agent can gain a new tool. A
model file can bring an unsafe deserialisation path into an otherwise familiar
software stack.
A conventional software bill of materials can identify libraries, packages,
containers, and some of their dependency relationships. It may say little about
the model, training and evaluation data, prompt configuration, retrieval
artifacts, hosted services, agent authority, or evidence used to approve the
deployed system.
The G7 Cybersecurity Working Group addressed this gap in 2026 with Software
Bill of Materials for AI - Minimum Elements . The guidance says that AI systems
remain software systems, so the AI-specific information supplements rather
than replaces an ordinary SBOM. It proposes 50 elements across seven clusters:
metadata, system-level properties, models, datasets, infrastructure, security
properties, and key performance indicators.[1][2]
The guidance is deliberately limited. The elements are non-mandatory,
non-exhaustive, and expected to evolve. The document also states that an SBOM
for AI is insufficient by itself. It has to connect to vulnerability
management, advisories, and other security tools before the inventory can
support protection of the supply chain.[1][2]
That is useful progress. It also exposes a category problem.
A model hash is a component identifier. A supplier's description of known
limitations is a declaration. An internal evaluation result is an assessment.
The model endpoint active for one production request is a runtime observation.
System uptime is a time-bound operational measure. These records have different
owners, evidence, freshness, and confidence. Placing them in one inventory
without preserving those differences can make weak claims look verified and
old observations look current.
This paper proposes an operational model for an AI Bill of Materials:
Record the AI system as a dependency graph rather than a flat component
list.
Bind that graph to a specific build, release, deployment, and operating
environment.
Classify each entry as a verifiable component fact, supplier declaration,
organisational assessment, deployment observation, runtime observation, or
explicit unknown.
Keep stable composition records separate from dynamic assurance and
operational evidence, while linking them through durable identifiers.
Connect the record to vulnerability response, model migration, evaluation,
incident investigation, and recovery workflows.
Protect sensitive versions of the record rather than assuming every detail
should be public.
An operational AI BOM should answer practical questions:
Which production systems use the affected model, dataset, library, or hosted
service?
Which prompts, retrieval indexes, tools, and policies were active in the
affected release?
Which supplier statements were independently tested?
Which outputs or actions were produced under that composition?
What must be reevaluated when a model, embedding service, or dataset changes?
Can the affected component be isolated or replaced?
The public record shows why these questions matter. A compromised
torchtriton package entered PyTorch nightly installations through dependency
confusion and attempted to exfiltrate system and user files.[16] A later
PyTorch advisory described remote code execution through a bypass of
torch.load(weights_only=True) .[17] Hugging Face created and independently
audited the safetensors format because pickle-based model loading can execute
malicious code.[18] LAION temporarily withdrew its datasets after researchers
reported links to suspected illegal material.[24][25] Major hosted model
providers publish retirement and migration policies because model versions and
service interfaces do not remain available indefinitely.[20][21][22][23]
These cases do not prove that an AI BOM would have prevented the underlying
event. An inventory does not establish that a component is safe, accurate,
fair, lawful, or suitable for a particular use. A hash proves equality with a
known digest, not trustworthiness. A signed declaration identifies the signer
and protects integrity, but it does not make the claim true.
The value of an AI BOM is narrower and operational. It reduces the time required
to identify what is present, where it came from, how it is connected, what
changed, which evidence applies, and which systems require action.
Treat the AI BOM as a map. Assurance comes from the evidence, controls, and
operating processes connected to that map.
Production AI is assembled from more than a model.
Consider a document assistant that answers questions and can create a support
case. Its behaviour may depend on:
the application and ordinary software packages;
the hosted or self-managed model;
model weights, architecture, tokeniser, and inference configuration;
safety and policy instructions;
source documents and their access controls;
a vector index and its build time;
tool definitions and argument schemas;
the identity used for retrieval and actions;
memory retained from earlier interactions;
output validation and post-processing;
infrastructure, hardware, and regional deployment choices.
Each component can change independently. Some are immutable files. Some are
database state. Some are configuration. Some are external services whose
internal implementation is unavailable to the deployer.
An ordinary SBOM remains necessary. The assistant still depends on an operating
system, language runtime, client libraries, parsers, web frameworks, container
images, and cryptographic packages. Those components create familiar
vulnerability and licensing obligations.
The ordinary SBOM cannot always identify:
which hosted model deployment served a request;
which model artifact a provider name referred to at that time;
which prompt or policy version governed the response;
which corpus snapshot and embedding model produced the active index;
which tools the agent could call;
which permission set the tool used;
which evaluation result justified promotion;
which supplier claims were accepted without independent verification.
This is why AI-specific inventory work cannot begin and end with the package
manager.
The missing information is also why "model inventory" is too narrow. A model
may be unchanged while the system behaves differently because retrieval,
prompting, tool access, or policy changed. A deployment can fail after an
embedding model retirement even when its generation model remains available.
A vulnerable model loader can matter even if the model weights themselves are
benign.
The operational unit is the assembled AI system.
2. What an AI Bill of Materials means
The terminology is not settled.
The G7 guidance uses "SBOM for AI". CISA describes it as supplemental
information for an AI system on top of the general SBOM minimum
elements.[1][3]
SPDX 3 defines an AIPackage with properties including model type, training
information, limitations, hyperparameters, metrics, safety risk assessment, and
use of sensitive personal information. Its DatasetPackage records properties
such as collection, preprocessing, availability, known bias, sensitivity, and
intended use.[5][6]
CycloneDX uses "AI/ML-BOM" and represents models, datasets, configurations,
training methods, provenance, and risk-related information within its broader
BOM model.[7][8]
Other records serve related purposes:
The model-card proposal described model cards as short records accompanying
released models, including intended uses, evaluation results, and
limitations.[9] SLSA defines an attestation as an authenticated statement
about one or more software artifacts.[10] Those records can be linked to a
BOM, but they should not be collapsed into the same concept.
This paper uses "AI Bill of Materials", or "AI BOM", to mean:
A machine-processable record of the components, relationships, evidence, and
deployment identity required to understand the composition of an AI system.
That definition is intentionally system-focused. It includes ordinary software
and AI-specific components. It also requires links to evidence and deployment
state because a generic list of possible ingredients cannot explain what was
actually running.
The G7 minimum-elements document is one of the more complete international
attempts so far to describe the additional information an AI supply-chain
record should contain.
It divides the record into seven clusters:
The count matters less than the shape of the proposal. It recognises that a
useful record needs:
a version and author of its own;
machine-processable relationships;
training and post-training information;
dataset provenance and sensitivity;
infrastructure beyond the model;
security and operational context.
The document also contains several important boundaries.
First, it states that the elements are non-mandatory and non-exhaustive. The
guidance does not create a standard, legislation, or a complete implementation
model.[1][2]
Second, it says AI systems are software systems. The AI-specific elements are
additional to the ordinary SBOM rather than a substitute for it.[1][2]
Third, it warns that the record is insufficient without operational security
tooling. The document calls for connections to vulnerability scanning,
vulnerability management, security advisories, and bulletins.[2]
Fourth, the discussion considered the decision-making or autonomy level of an
AI system but did not make it a separate element. The authors note that
autonomy may become more relevant as agentic systems develop.[2]
The baseline therefore establishes a useful starting point while leaving open
questions about prompts, retrieval indexes, memory, tool authority, approval
boundaries, and request-level runtime evidence. Some can fit inside
system-level properties or dependency relationships. They do not yet have
consistent first-class treatment across the main BOM formats.
Other guidance fills part of that gap. The UK National Cyber Security Centre
asks AI providers to document the creation, operation, and lifecycle management
of models, datasets, and meta- or system prompts. Its

[truncated]

## Original Extract

Why an AI Bill of Materials must connect models, data, software, configuration, evidence and deployments to support security, change and incident response.

Skip to content
Jason Doyle
Open navigation
Home
Whitepapers
Contact
AI supply chain / Operational assurance
The AI Bill of Materials Is an Operational Record
What an inventory must show before it can support security, change and incident response.
2. What an AI Bill of Materials means
5. Different entries carry different evidence
Verifiable component facts
6. Bind the record to a deployed release
7. Four connected records
1. Component graph
8. What the record should support
Vulnerability response
Dataset withdrawal and correction
Procurement and supplier review
Regulation and technical documentation
9. Failure modes
Generated once at procurement
Accurate components, missing relationships
Stable name, changing implementation
Supplier claim presented as verification
Dynamic evidence copied into a static file
Inventory without response tooling
10. A practical operating model
Step 1: Define the questions
Step 2: Set the system boundary
Step 3: Start with the ordinary SBOM
Step 4: Add AI-specific components
Step 7: Bind every production release
Step 9: Connect operational systems
11. What to measure
Deployment binding coverage
12. Counterarguments
Model and system cards already provide this information
Hosted models make complete inventory impossible
The system changes too quickly
Automated generation solves the problem
A detailed BOM creates an attack map
Compliance programmes will define the required fields
13. What this paper does not claim
Appendix A: Illustrative CycloneDX JSON
Disclosure: These views are my own and do not represent my current or any former employers. This paper uses only public sources and does not describe non-public product information.
An AI system can change without its application code changing.
A model provider can retire a version or move a deployment to a replacement.
An embedding model can change. A retrieval index can be rebuilt. A prompt can be
edited. A fine-tuning dataset can be withdrawn. An agent can gain a new tool. A
model file can bring an unsafe deserialisation path into an otherwise familiar
software stack.
A conventional software bill of materials can identify libraries, packages,
containers, and some of their dependency relationships. It may say little about
the model, training and evaluation data, prompt configuration, retrieval
artifacts, hosted services, agent authority, or evidence used to approve the
deployed system.
The G7 Cybersecurity Working Group addressed this gap in 2026 with Software
Bill of Materials for AI - Minimum Elements . The guidance says that AI systems
remain software systems, so the AI-specific information supplements rather
than replaces an ordinary SBOM. It proposes 50 elements across seven clusters:
metadata, system-level properties, models, datasets, infrastructure, security
properties, and key performance indicators.[1][2]
The guidance is deliberately limited. The elements are non-mandatory,
non-exhaustive, and expected to evolve. The document also states that an SBOM
for AI is insufficient by itself. It has to connect to vulnerability
management, advisories, and other security tools before the inventory can
support protection of the supply chain.[1][2]
That is useful progress. It also exposes a category problem.
A model hash is a component identifier. A supplier's description of known
limitations is a declaration. An internal evaluation result is an assessment.
The model endpoint active for one production request is a runtime observation.
System uptime is a time-bound operational measure. These records have different
owners, evidence, freshness, and confidence. Placing them in one inventory
without preserving those differences can make weak claims look verified and
old observations look current.
This paper proposes an operational model for an AI Bill of Materials:
Record the AI system as a dependency graph rather than a flat component
list.
Bind that graph to a specific build, release, deployment, and operating
environment.
Classify each entry as a verifiable component fact, supplier declaration,
organisational assessment, deployment observation, runtime observation, or
explicit unknown.
Keep stable composition records separate from dynamic assurance and
operational evidence, while linking them through durable identifiers.
Connect the record to vulnerability response, model migration, evaluation,
incident investigation, and recovery workflows.
Protect sensitive versions of the record rather than assuming every detail
should be public.
An operational AI BOM should answer practical questions:
Which production systems use the affected model, dataset, library, or hosted
service?
Which prompts, retrieval indexes, tools, and policies were active in the
affected release?
Which supplier statements were independently tested?
Which outputs or actions were produced under that composition?
What must be reevaluated when a model, embedding service, or dataset changes?
Can the affected component be isolated or replaced?
The public record shows why these questions matter. A compromised
torchtriton package entered PyTorch nightly installations through dependency
confusion and attempted to exfiltrate system and user files.[16] A later
PyTorch advisory described remote code execution through a bypass of
torch.load(weights_only=True) .[17] Hugging Face created and independently
audited the safetensors format because pickle-based model loading can execute
malicious code.[18] LAION temporarily withdrew its datasets after researchers
reported links to suspected illegal material.[24][25] Major hosted model
providers publish retirement and migration policies because model versions and
service interfaces do not remain available indefinitely.[20][21][22][23]
These cases do not prove that an AI BOM would have prevented the underlying
event. An inventory does not establish that a component is safe, accurate,
fair, lawful, or suitable for a particular use. A hash proves equality with a
known digest, not trustworthiness. A signed declaration identifies the signer
and protects integrity, but it does not make the claim true.
The value of an AI BOM is narrower and operational. It reduces the time required
to identify what is present, where it came from, how it is connected, what
changed, which evidence applies, and which systems require action.
Treat the AI BOM as a map. Assurance comes from the evidence, controls, and
operating processes connected to that map.
Production AI is assembled from more than a model.
Consider a document assistant that answers questions and can create a support
case. Its behaviour may depend on:
the application and ordinary software packages;
the hosted or self-managed model;
model weights, architecture, tokeniser, and inference configuration;
safety and policy instructions;
source documents and their access controls;
a vector index and its build time;
tool definitions and argument schemas;
the identity used for retrieval and actions;
memory retained from earlier interactions;
output validation and post-processing;
infrastructure, hardware, and regional deployment choices.
Each component can change independently. Some are immutable files. Some are
database state. Some are configuration. Some are external services whose
internal implementation is unavailable to the deployer.
An ordinary SBOM remains necessary. The assistant still depends on an operating
system, language runtime, client libraries, parsers, web frameworks, container
images, and cryptographic packages. Those components create familiar
vulnerability and licensing obligations.
The ordinary SBOM cannot always identify:
which hosted model deployment served a request;
which model artifact a provider name referred to at that time;
which prompt or policy version governed the response;
which corpus snapshot and embedding model produced the active index;
which tools the agent could call;
which permission set the tool used;
which evaluation result justified promotion;
which supplier claims were accepted without independent verification.
This is why AI-specific inventory work cannot begin and end with the package
manager.
The missing information is also why "model inventory" is too narrow. A model
may be unchanged while the system behaves differently because retrieval,
prompting, tool access, or policy changed. A deployment can fail after an
embedding model retirement even when its generation model remains available.
A vulnerable model loader can matter even if the model weights themselves are
benign.
The operational unit is the assembled AI system.
2. What an AI Bill of Materials means
The terminology is not settled.
The G7 guidance uses "SBOM for AI". CISA describes it as supplemental
information for an AI system on top of the general SBOM minimum
elements.[1][3]
SPDX 3 defines an AIPackage with properties including model type, training
information, limitations, hyperparameters, metrics, safety risk assessment, and
use of sensitive personal information. Its DatasetPackage records properties
such as collection, preprocessing, availability, known bias, sensitivity, and
intended use.[5][6]
CycloneDX uses "AI/ML-BOM" and represents models, datasets, configurations,
training methods, provenance, and risk-related information within its broader
BOM model.[7][8]
Other records serve related purposes:
The model-card proposal described model cards as short records accompanying
released models, including intended uses, evaluation results, and
limitations.[9] SLSA defines an attestation as an authenticated statement
about one or more software artifacts.[10] Those records can be linked to a
BOM, but they should not be collapsed into the same concept.
This paper uses "AI Bill of Materials", or "AI BOM", to mean:
A machine-processable record of the components, relationships, evidence, and
deployment identity required to understand the composition of an AI system.
That definition is intentionally system-focused. It includes ordinary software
and AI-specific components. It also requires links to evidence and deployment
state because a generic list of possible ingredients cannot explain what was
actually running.
The G7 minimum-elements document is one of the more complete international
attempts so far to describe the additional information an AI supply-chain
record should contain.
It divides the record into seven clusters:
The count matters less than the shape of the proposal. It recognises that a
useful record needs:
a version and author of its own;
machine-processable relationships;
training and post-training information;
dataset provenance and sensitivity;
infrastructure beyond the model;
security and operational context.
The document also contains several important boundaries.
First, it states that the elements are non-mandatory and non-exhaustive. The
guidance does not create a standard, legislation, or a complete implementation
model.[1][2]
Second, it says AI systems are software systems. The AI-specific elements are
additional to the ordinary SBOM rather than a substitute for it.[1][2]
Third, it warns that the record is insufficient without operational security
tooling. The document calls for connections to vulnerability scanning,
vulnerability management, security advisories, and bulletins.[2]
Fourth, the discussion considered the decision-making or autonomy level of an
AI system but did not make it a separate element. The authors note that
autonomy may become more relevant as agentic systems develop.[2]
The baseline therefore establishes a useful starting point while leaving open
questions about prompts, retrieval indexes, memory, tool authority, approval
boundaries, and request-level runtime evidence. Some can fit inside
system-level properties or dependency relationships. They do not yet have
consistent first-class treatment across the main BOM formats.
Other guidance fills part of that gap. The UK National Cyber Security Centre
asks AI providers to document the creation, operation, and lifecycle management
of models, datasets, and meta- or system prompts. Its

[truncated]
