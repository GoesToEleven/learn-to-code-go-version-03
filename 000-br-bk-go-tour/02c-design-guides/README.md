# "A coach is someone who will tell you what you don't want to hear, and show you what you don't want to see, so that you can be who you want to be."

# Video
[video on youtube](https://youtu.be/WkQFrctSDsc)

### This document is to capture **design guidelines** that can be used for writing better Go code. These guidelines can provide guidance for code reviews. You are welcome to contribute to this document.

# Code review check

### Clarity, readability, and abstraction
- Eliminate needless abstraction. Focus on readability. Write code that can be understood by average developers.
    - "I've seen go code where the author has you going down a rabbit hole of functions calling other functions, those functions are in different files, and each of those functions are only five to ten lines, and are only used once... If it's only used once - why create a new function? Very hard to follow that code. It seems like they think there is a size limit, that a function that has more than 10 lines in it is bad." ~ Richard M.

### Clarity, readability, and nesting
- "One of my largest concerns for design is nesting. Nesting with for loops, if else if, switch statements. When I see at the end of a chunk of code five levels of curly braces -- I need to simplify that. Another is nested functions." ~ Richard M. 