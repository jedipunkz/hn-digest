---
source: "https://www.csoonline.com/article/4234854/ai-agents-arent-going-rogue-frontier-labs-are-skipping-security-basics.html"
hn_url: "https://news.ycombinator.com/item?id=50091788"
title: "AI agents aren't 'going rogue.' Frontier labs are skipping security basics"
article_title: "AI agents aren’t ‘going rogue.’ Frontier labs are skipping security basics. | CSO Online"
image: "https://www.csoonline.com/wp-content/uploads/2026/10/4234854-0-10823900-1791966468-shutterstock_2472843149.jpg?quality=50&strip=all&w=1024"
author: "latexr"
captured_at: "2026-10-14T18:20:28Z"
capture_tool: "hn-digest"
hn_id: 50091788
score: 97
comments: 27
posted_at: "2026-10-14T17:45:23Z"
tags:
  - hacker-news
---

# AI agents aren't 'going rogue.' Frontier labs are skipping security basics

- HN: [50091788](https://news.ycombinator.com/item?id=50091788)
- Source: [www.csoonline.com](https://www.csoonline.com/article/4234854/ai-agents-arent-going-rogue-frontier-labs-are-skipping-security-basics.html)
- Score: 97
- Comments: 27
- Posted: 2026-10-14T17:45:23Z

## Translation

Title: AI agents aren't 'going rogue.' Frontier labs are skipping security basics
Article title: AI agents aren’t ‘going rogue.’ Frontier labs are skipping security basics. | CSO Online
Description: Cyber experts say AI labs’ containment and monitoring failures — not agents with minds of their own — are to blame for recent alarming hacking reports. CISOs should vet their AI vendors and tighten security basics as safeguards.

Article text:
AI agents aren’t ‘going rogue.’ Frontier labs are skipping security basics. | CSO Online
Topics
Identity and Access Management
Buyer’s Guides & Enterprise Spotlights
Identity and Access Management
AI agents aren’t ‘going rogue.’ Frontier labs are skipping security basics.
Cyber experts say AI labs’ containment and monitoring failures — not agents with minds of their own — are to blame for recent alarming hacking reports. CISOs should vet their AI vendors and tighten security basics as safeguards.
Since July, US frontier AI labs and other AI leaders have revealed a string of incidents in which their models accessed external computer systems without authorization, events popularly described as AI agents “going rogue.”
The most well-known and studied of these incidents is OpenAI’s hack into Hugging Face, which OpenAI initially characterized as “an unprecedented cyber incident, involving state-of-the-art cyber capabilities,” and later said was due to “misaligned behavior,” or behavior its creators did not intend.
Following news of the Hugging Face incident, Anthropic revealed that one of its models gained unauthorized access to the production infrastructure of three organizations; Meta said one of its AI models hacked another company ​during cybersecurity testing; and Google confessed that its Gemini model carried out three “undirected computer hacks.”
Although press reports and the model providers themselves use language such as rogue agents or misalignment, cybersecurity experts who spoke with CSO argue that language obscures a simple truth: These incidents were caused by inexcusable failures by the frontier labs to implement cybersecurity practices that have been well-known for decades.
Compounding the problem, according to a half dozen cybersecurity experts CSO interviewed at three industry conferences in September and in subsequent conversations, is that the frontier labs, particularly OpenAI and Anthropic, have rejected their advice on how to improve their cybersecurity risk management practices to avoid these kinds of surreptitious hacks, often lumping cybersecurity practices under “AI safety,” in which they believe they have superior expertise.
“They are not as good as they think they are, honestly,” Gary McGraw, co-founder of the Berryville Institute of Machine Learning, tells CSO. “It would be better if they had more respect for expertise that is hard-earned from people in all sorts of subsets of security, like forensics or security engineering, building sandboxes, setting up or monitoring network security situations.”
Neither Anthropic nor OpenAI responded to CSO’s requests for comment.
Inexcusable failures on cybersecurity basics, particularly containment
According to the cybersecurity professionals CSO spoke with, one of the biggest cybersecurity failures by the frontier AI labs is the lack of implementation of tried-and-true containment practices that stretch back to the 1990s, such as setting up secure sandboxes that would effectively bar agents from escaping their test environments.
“These are things that we’ve been dealing with for years,” Rob T. Lee, chief AI officer and chief of research at SANS Institute, tells CSO. “There’s no reason why every single time this happens, it’s, ‘Oops, we don’t know how they got out.’”
“You can tell that their security teams are fairly nascent,” Lee says. “Folks on the cybersecurity containment side do argue that the majority of reasons that we’re seeing containment breached is that it’s not set up correctly.”
Lee adds, “The Hugging Face incident definitively showed that it wasn’t the fact that the agents were super smart. It’s that there literally was a hole that was exploited because it was not configured correctly on the containment side.”
Dan Wallach, expert technical resident at RAND’s Center on AI, Security, and Technology, thinks the security failures of the AI labs boil down to the fact that they don’t prioritize cybersecurity.
“I would say they are not prioritizing cybersecurity,” he tells CSO. “I don’t think I can say that for all of the labs,” Wallach says, pointing to French AI lab Mistral’s launch of its Large 4 AI model, which the company said tried to operate beyond its testing environment before Mistral cut it off.
And, he adds, “It’s not rocket science to isolate one of these AI things. Security people have been talking about how to do either air-gapping, so there’s literally no electrical way for data to get from here to here, or absent that, to have really strong egress control to make sure that nothing pushes through.”
On top of that, the frontier labs appear to be conducting insufficient monitoring of their agents’ behavior. The Hugging Face incident “went on for days, and nobody noticed,” Wallach says. “People started these experiments and weren’t watching them. That’s an absence of control. And then they’re surprised by what happened.”
Divorcing safety from security
One barrier to persuading frontier labs to up their security game is the intense focus they have on what they call “safety,” which they don’t extend to traditional cybersecurity practices.
“I can see how security behavior looks like AI safety,” RAND’s Wallach says. “But if the way that you’re evaluating this AI tool is by unleashing it on the internet and seeing what happens, then this is what classical security tooling is for.”
SANS Institute’s Lee says, “If these frontier labs really wanted to test their stuff, they need to beef up their cybersecurity component. But their gut reaction in the response to their agents getting out is to say, ‘It’s not a cybersecurity problem.’ It has literally become an AI safety problem, which is divorced from the cybersecurity side.”
Many AI experts suggest that if the AI companies were to embrace advice from cybersecurity specialists, it could help put an end to the kind of unrestrained AI agent behavior causing concerns. “AI safety focuses on the rules of the road,” Lee says. “Everything is about guardrail rules.”
But “I do believe cybersecurity can contain AI,” he says. “We just need to be in the room where it happens and have a seat at the table. It’s important for our field to be as close to the AI safety folks as possible.”
Bad cybersecurity could increase liability
If frontier labs are overlooking cybersecurity fundamentals that could reduce or eliminate the types of incidents that have created widespread concerns, that strengthens the case that they should be held liable for their “rogue” agents. The insurance industry is already lining up on the side of holding AI industry executives liable for these actions.
Democratic Senators have introduced legislation that would hold AI agent operators liable for hacking incidents. Even the AI-friendly White House has seemingly concluded that the companies should be held liable for their errant agents, while the US Federal Trade Commission is conducting a probe into Anthropic, OpenAI, and other AI labs into the dangers their technologies pose to consumers.
However, in recent legislative and council hearings, lab representatives have held fast to the notion that they are not responsible for AI agents run amok.
Experts say this position is untenable, particularly in light of the labs’ cybersecurity failures.
BIML’s McGraw characterizes the labs’ position as: “Don’t send the lawyers to us, send it to our program. It’s our program’s fault.” His assessment of that argument: “That’s just high BS.”
He also questions whether the labs’ operational failures undermine confidence in their broader security capabilities. “They live inside their own bubble,” he says. “The fact that they cannot do clear, easy operations properly does not bode well for their capabilities in security engineering, which is much harder to do.”
SANS Lee says, “That’s where the traceability for both liability and insurance comes into play, because if one of these agents is detected, and they’re able to trace it back to the person that initiated it, liability’s going to kick in and say, ‘Hey, you, you broke the law, ergo, we need to prosecute.’”
What security leaders should do now
Even as evidence mounts that frontier labs are not implementing sound cybersecurity practices, and the debate over liability remains unsolved , CISOs and other security leaders can take steps to protect themselves from AI model misfires.
First, when it comes to assessing tools based on frontier lab models, “I would always get an external entity to help with that,” McGraw says. “AI is extremely powerful and extremely useful, but it comes with very serious grown-up risks that do impact the enterprise. Identify those risks, put in the right controls, make sure that AI adoption is not just arbitrary and willy-nilly and all over the enterprise, but has some centralized authority that is looking at the problem, and then get ready for the next big thing, which is swarms of agents.”
For Lee, the best protection is to ask the right questions. “It starts with asking the right questions and then finding the tools and capabilities for those [AI labs] that are able to demonstrate a higher level of cybersecurity competency,” he says. “You’re really looking for the people on the teams that are dealing with really hard problems.”
But fundamentally, the best preparation to deal with AI model cyber weaknesses is to make sure all the long-standing cybersecurity risk management practices are in place across the organization. “This is the least exciting part of the discussion,” Wallach says. “It’s the cyber version of ‘eat your vegetables.’”
“We need to re-engineer our software so it has better security,” he adds. “We need to have multi-factor authentication on all of our accounts. We need to eat our vegetables from a defensive cyber perspective. And we know how to do it.”
Cynthia Brumfield is a veteran communications and technology analyst who is currently focused on cybersecurity. She runs a cybersecurity news destination site, Metacurity.com , consults with companies through her firm DCT-Associates, and is the author of the book published by Wiley, Cybersecurity Risk Management: Mastering the Fundamentals Using the NIST Cybersecurity Framework .
Cynthia holds a Master of Planning Degree from the University of Virginia and a Bachelor’s degree from The George Washington University. She has won multiple AZBEE awards for her work on CSO, including two in 2025.
feature AI is turning offensive security into a continuous necessity
feature CISOs can no longer ignore the nation-state threat
feature AI made software development unrecognizable. Is cybersecurity next?
news OpenAI targets small utilities with $1 billion cyber defense initiative
news analysis How China industrialized the infrastructure behind state hacking
news analysis Critical infrastructure’s long, undefended tail exposed by UK energy attack
feature The AI harness is the new attack surface
How Claude Code’s 256KB blind spot put developers’ secrets in the preview pane
Hackers abuse the Ivanti Connect Secure bug to gain access without accounts
Unplug your smart fridges, French supermarkets warned after cyberattack
Identity Has Become the New Security Perimeter
Why Email Still Works for Attackers
The Evolving CISO: AI, Leadership and the Future of Cybersecurity
Identity Has Become the New Security Perimeter
Why Email Still Works for Attackers
The Evolving CISO: AI, Leadership and the Future of Cybersecurity
Add CSO as a Preferred Source in Google Search
Your California Privacy Rights
Buyer’s Guides & Enterprise Spotlights

## Original Extract

Cyber experts say AI labs’ containment and monitoring failures — not agents with minds of their own — are to blame for recent alarming hacking reports. CISOs should vet their AI vendors and tighten security basics as safeguards.

AI agents aren’t ‘going rogue.’ Frontier labs are skipping security basics. | CSO Online
Topics
Identity and Access Management
Buyer’s Guides & Enterprise Spotlights
Identity and Access Management
AI agents aren’t ‘going rogue.’ Frontier labs are skipping security basics.
Cyber experts say AI labs’ containment and monitoring failures — not agents with minds of their own — are to blame for recent alarming hacking reports. CISOs should vet their AI vendors and tighten security basics as safeguards.
Since July, US frontier AI labs and other AI leaders have revealed a string of incidents in which their models accessed external computer systems without authorization, events popularly described as AI agents “going rogue.”
The most well-known and studied of these incidents is OpenAI’s hack into Hugging Face, which OpenAI initially characterized as “an unprecedented cyber incident, involving state-of-the-art cyber capabilities,” and later said was due to “misaligned behavior,” or behavior its creators did not intend.
Following news of the Hugging Face incident, Anthropic revealed that one of its models gained unauthorized access to the production infrastructure of three organizations; Meta said one of its AI models hacked another company ​during cybersecurity testing; and Google confessed that its Gemini model carried out three “undirected computer hacks.”
Although press reports and the model providers themselves use language such as rogue agents or misalignment, cybersecurity experts who spoke with CSO argue that language obscures a simple truth: These incidents were caused by inexcusable failures by the frontier labs to implement cybersecurity practices that have been well-known for decades.
Compounding the problem, according to a half dozen cybersecurity experts CSO interviewed at three industry conferences in September and in subsequent conversations, is that the frontier labs, particularly OpenAI and Anthropic, have rejected their advice on how to improve their cybersecurity risk management practices to avoid these kinds of surreptitious hacks, often lumping cybersecurity practices under “AI safety,” in which they believe they have superior expertise.
“They are not as good as they think they are, honestly,” Gary McGraw, co-founder of the Berryville Institute of Machine Learning, tells CSO. “It would be better if they had more respect for expertise that is hard-earned from people in all sorts of subsets of security, like forensics or security engineering, building sandboxes, setting up or monitoring network security situations.”
Neither Anthropic nor OpenAI responded to CSO’s requests for comment.
Inexcusable failures on cybersecurity basics, particularly containment
According to the cybersecurity professionals CSO spoke with, one of the biggest cybersecurity failures by the frontier AI labs is the lack of implementation of tried-and-true containment practices that stretch back to the 1990s, such as setting up secure sandboxes that would effectively bar agents from escaping their test environments.
“These are things that we’ve been dealing with for years,” Rob T. Lee, chief AI officer and chief of research at SANS Institute, tells CSO. “There’s no reason why every single time this happens, it’s, ‘Oops, we don’t know how they got out.’”
“You can tell that their security teams are fairly nascent,” Lee says. “Folks on the cybersecurity containment side do argue that the majority of reasons that we’re seeing containment breached is that it’s not set up correctly.”
Lee adds, “The Hugging Face incident definitively showed that it wasn’t the fact that the agents were super smart. It’s that there literally was a hole that was exploited because it was not configured correctly on the containment side.”
Dan Wallach, expert technical resident at RAND’s Center on AI, Security, and Technology, thinks the security failures of the AI labs boil down to the fact that they don’t prioritize cybersecurity.
“I would say they are not prioritizing cybersecurity,” he tells CSO. “I don’t think I can say that for all of the labs,” Wallach says, pointing to French AI lab Mistral’s launch of its Large 4 AI model, which the company said tried to operate beyond its testing environment before Mistral cut it off.
And, he adds, “It’s not rocket science to isolate one of these AI things. Security people have been talking about how to do either air-gapping, so there’s literally no electrical way for data to get from here to here, or absent that, to have really strong egress control to make sure that nothing pushes through.”
On top of that, the frontier labs appear to be conducting insufficient monitoring of their agents’ behavior. The Hugging Face incident “went on for days, and nobody noticed,” Wallach says. “People started these experiments and weren’t watching them. That’s an absence of control. And then they’re surprised by what happened.”
Divorcing safety from security
One barrier to persuading frontier labs to up their security game is the intense focus they have on what they call “safety,” which they don’t extend to traditional cybersecurity practices.
“I can see how security behavior looks like AI safety,” RAND’s Wallach says. “But if the way that you’re evaluating this AI tool is by unleashing it on the internet and seeing what happens, then this is what classical security tooling is for.”
SANS Institute’s Lee says, “If these frontier labs really wanted to test their stuff, they need to beef up their cybersecurity component. But their gut reaction in the response to their agents getting out is to say, ‘It’s not a cybersecurity problem.’ It has literally become an AI safety problem, which is divorced from the cybersecurity side.”
Many AI experts suggest that if the AI companies were to embrace advice from cybersecurity specialists, it could help put an end to the kind of unrestrained AI agent behavior causing concerns. “AI safety focuses on the rules of the road,” Lee says. “Everything is about guardrail rules.”
But “I do believe cybersecurity can contain AI,” he says. “We just need to be in the room where it happens and have a seat at the table. It’s important for our field to be as close to the AI safety folks as possible.”
Bad cybersecurity could increase liability
If frontier labs are overlooking cybersecurity fundamentals that could reduce or eliminate the types of incidents that have created widespread concerns, that strengthens the case that they should be held liable for their “rogue” agents. The insurance industry is already lining up on the side of holding AI industry executives liable for these actions.
Democratic Senators have introduced legislation that would hold AI agent operators liable for hacking incidents. Even the AI-friendly White House has seemingly concluded that the companies should be held liable for their errant agents, while the US Federal Trade Commission is conducting a probe into Anthropic, OpenAI, and other AI labs into the dangers their technologies pose to consumers.
However, in recent legislative and council hearings, lab representatives have held fast to the notion that they are not responsible for AI agents run amok.
Experts say this position is untenable, particularly in light of the labs’ cybersecurity failures.
BIML’s McGraw characterizes the labs’ position as: “Don’t send the lawyers to us, send it to our program. It’s our program’s fault.” His assessment of that argument: “That’s just high BS.”
He also questions whether the labs’ operational failures undermine confidence in their broader security capabilities. “They live inside their own bubble,” he says. “The fact that they cannot do clear, easy operations properly does not bode well for their capabilities in security engineering, which is much harder to do.”
SANS Lee says, “That’s where the traceability for both liability and insurance comes into play, because if one of these agents is detected, and they’re able to trace it back to the person that initiated it, liability’s going to kick in and say, ‘Hey, you, you broke the law, ergo, we need to prosecute.’”
What security leaders should do now
Even as evidence mounts that frontier labs are not implementing sound cybersecurity practices, and the debate over liability remains unsolved , CISOs and other security leaders can take steps to protect themselves from AI model misfires.
First, when it comes to assessing tools based on frontier lab models, “I would always get an external entity to help with that,” McGraw says. “AI is extremely powerful and extremely useful, but it comes with very serious grown-up risks that do impact the enterprise. Identify those risks, put in the right controls, make sure that AI adoption is not just arbitrary and willy-nilly and all over the enterprise, but has some centralized authority that is looking at the problem, and then get ready for the next big thing, which is swarms of agents.”
For Lee, the best protection is to ask the right questions. “It starts with asking the right questions and then finding the tools and capabilities for those [AI labs] that are able to demonstrate a higher level of cybersecurity competency,” he says. “You’re really looking for the people on the teams that are dealing with really hard problems.”
But fundamentally, the best preparation to deal with AI model cyber weaknesses is to make sure all the long-standing cybersecurity risk management practices are in place across the organization. “This is the least exciting part of the discussion,” Wallach says. “It’s the cyber version of ‘eat your vegetables.’”
“We need to re-engineer our software so it has better security,” he adds. “We need to have multi-factor authentication on all of our accounts. We need to eat our vegetables from a defensive cyber perspective. And we know how to do it.”
Cynthia Brumfield is a veteran communications and technology analyst who is currently focused on cybersecurity. She runs a cybersecurity news destination site, Metacurity.com , consults with companies through her firm DCT-Associates, and is the author of the book published by Wiley, Cybersecurity Risk Management: Mastering the Fundamentals Using the NIST Cybersecurity Framework .
Cynthia holds a Master of Planning Degree from the University of Virginia and a Bachelor’s degree from The George Washington University. She has won multiple AZBEE awards for her work on CSO, including two in 2025.
feature AI is turning offensive security into a continuous necessity
feature CISOs can no longer ignore the nation-state threat
feature AI made software development unrecognizable. Is cybersecurity next?
news OpenAI targets small utilities with $1 billion cyber defense initiative
news analysis How China industrialized the infrastructure behind state hacking
news analysis Critical infrastructure’s long, undefended tail exposed by UK energy attack
feature The AI harness is the new attack surface
How Claude Code’s 256KB blind spot put developers’ secrets in the preview pane
Hackers abuse the Ivanti Connect Secure bug to gain access without accounts
Unplug your smart fridges, French supermarkets warned after cyberattack
Identity Has Become the New Security Perimeter
Why Email Still Works for Attackers
The Evolving CISO: AI, Leadership and the Future of Cybersecurity
Identity Has Become the New Security Perimeter
Why Email Still Works for Attackers
The Evolving CISO: AI, Leadership and the Future of Cybersecurity
Add CSO as a Preferred Source in Google Search
Your California Privacy Rights
Buyer’s Guides & Enterprise Spotlights
