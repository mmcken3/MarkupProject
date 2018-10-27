# Source
This folder contains the cmd folder where the markup main.go file is located. It also contains four interfaces that are used for the MarkupScores project.

### Parser
This interface will parse html data to find tokens using its only method, `Parse(*os.String, chan html.Token)`. It will find each complete token and then send it out on a channel. 

### Scorer
This interface will recieve html tokens on a channel and score them according to the scoring guide. Once the token is scored it will send out the scores on an output channel. All of this is done through its method `Score(<-chan html.Token) <-chan int`.

### Calculator
This interface is fairly simple and will just add up any scores it recieves to a total and return the score through its method `Calc(<-chan int) int`.

### Mysql
The Mysql interface will be used for communication of the gocode to the mysql database. It contains multiple functions for connecting to and working with the database. These are:
```
StoreScore(string, int)             Stores a score and id into the db
ScoresForID(string)                 Prints the scores associated with an id.
FirstAndLastID(int)                 Prints the highest or lowest score based on the param.
ScoresInRange(string, string)       Prints all scores in the date range.
AvgScore()                          Prints the average of all scores in the db.
```