set continue 1
set contentToWrite ""

while continue
    input str "Type a string or just press 'Enter' to stop: "
    length str strLength
    if strLength
        concat contentToWrite "$(str)\n"
    else
        set continue 0
    endif
endwhile

input fileDir "Enter the directory to save the file: "
input fileName "Enter the file name: "
set filePath ""

concat filePath fileDir
concat filePath "/"
concat filePath fileName

writefile filePath contentToWrite
