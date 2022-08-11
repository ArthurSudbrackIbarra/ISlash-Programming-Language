input word "Type a word: "

length word wordLength

rangearray wordLength indexes

foreach index indexes
    getchar word index char
    say "Char $(index): $(char)"
endforeach
