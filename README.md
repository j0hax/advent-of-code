# Advent of Code

> [!NOTE]  
> **All of my solutions here are devised without the assistance of LLMs.** While I have nothing against AI (I believe it can be a legitimate productivity tool), the purpose of *Advent of Code* is, in my opinion, to hone one's skills and enjoy the challenge... much in the same way that people run in a marathon, although driving the same distance is faster and less effort.

## 2024

<p align="center">
  <img src="fourteen/step7572.png">
</p>

Implementing in Go with the [Helix](https://helix-editor.com/) editor. The code is purposefully more verbose than many other solutions I see in discussion threads because I try to mimic more complex projects. As such, I often declare custom types and take a Post-OOP approach of defining methods on these.

My goal is to get **all stars** in a timely manner this year.[^stars]

**Why Go?**

I like this language: Although it's not perfect (is any language?), it combines speed, simplicity, and idiomatic/opinionated patterns.

**Practice makes perfect:** I would like to work as a software developer focusing on Go 😉

[^stars]: The holidays are a busy time for everyone; but I'm doing fine so far.

### Lessons Learned

- Memoization makes a **huge** difference in large recursion problems. I always took the approach of "solve first, optimize later," but realized the power of dynamic programming on day eleven, when simple but deep recursion algorithm can be reduced from many hours to milliseconds with only a hashmap.
- Go Syntax: I somehow missed the part on [named return values](https://go.dev/tour/basics/7). It's a neat way to shorten code while improving legibility, I guess.
