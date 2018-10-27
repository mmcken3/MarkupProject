MarkupProject
=============

Go project that will read HTML content input and score and give an arbitrary score based on a set of rules. Each score is given a unique id that is built using the html file name. Changes to the content can be re-run over time to determine improvement or regression of the score. Each unique run should be stored with the date and time it was run along with the score received for the content. The project will read from the data direcotry, parse the html in there, score it based on the system laid out below, and then store it in MySQL. If you want to see how the project does this using concurency look at src/README.md.

This was created in a couple of days for an initial interview project.

How to Start the Program
------------------------
Download and move code to $GOPATH, or:
```
go get github.com/mmcken3/MarkupProject
```
Naviage to that directory:
```
cd $GOATH/src/github.com/mmcken3/MarkupProject
```
Start a mysql server:
```
mysql.server start
```
Connect to mysql and run import for schema:
```
mysql -u root < ./schema/import.sql
```
Build cmd binary with:
```
go build ./src/cmd/markup
```
Start up the program with:
```
./markup
```

Working with Markup
-------------------
The program is interactive and can be stopped by sending the `quit` command. Help can be found anytime by sending `help`. This will display:
```
Score:      asks for a unique id, scores this html content, and then stores in the database
IdScore:    asks for a unique id, then prints the scores of this id
DateScores: asks for two dates, then reports all scores in those dates
HighId:     this reports the highest scored unique id
LowId:      this reports the lowest scored unique id
Avg:        this reports the average of all scores in the database
```
All of the commands are case insensitive and will need no parameters. Markup will request any information that it needs from you.

#### /data

* Contains the HTML content data to parse, format: (keyname_yyyy_mm_dd)

ie:
* dougs_2012_02_04.html
* dougs_2012_04_01.html
* dougs_2012_07_01.html

Example Use
-----------
```
$ ./markup
Welcome to MarkupScores, type help for instructions and quit to end!
Enter command: score
Enter Unique ID: john_2013_01_05

Scoring Complete!

Enter command: highid

Highest Scored ID: john_2013_01_05

Enter command: idscore
Enter Unique ID: john_2013_01_05

ID: john_2013_01_05
Score: 19

Enter command: quit
```

Scoring Rules
-------------
Each starting tag should below has been assigned a score. Each tag in the content should be added to or subtracted from the total score.

(We will assume for this project our html code creator created valid html)

| TagName | Score Modifier | TagName | Score Modifier |
| ------- | :------------: | ------- | -------------- |
| div     | 3              | font    | -1             |
| p       | 1              | center  | -2             |
| h1      | 3              | big     | -2             |
| h2      | 2              | strike  | -1             |
| html    | 5              | tt      | -2             |
| body    | 5              | frameset| -5             |
| header  | 10             | frame   | -5             |
| footer  | 10             |

example:

````
<html>
    <body>
      <p>foo</p>
      <p>bar</p>
      <div text-align='center'>
        <big>hello world</big>
      </div>
    </body>
</html>
````

2 p tags = 2 x 1 <br>
1 body tag = 1 x 5 <br>
1 html tag = 1 x 5 <br>
1 div tag = 1 x 3 <br>
1 big tag = 1 x -2 <br>
**Total Score: 13**

## Languages and tools used:

* go1.9.1 darwin/amd64
* mysql 14.14 Distrib 5.7.19
* macOS High Sierra 10.13
