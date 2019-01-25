# quiz
This quiz needs a csv file within the directory of the program.  This csv file
should be formatted question, answer.  for example:

8 + 6, 14

The default name of the file is problems.csv, but can be changed on the
command line via the flag --csvFile=  .  For example:

go run main.go --csvFile=test.csv

This program has a timer you can set that determines how much time (in seconds)
are given to complete the quiz.  The default time limit is 30 seconds.  This
can also be changed via the command line flag of --time= . For example:

go run main.go --time=60

This quiz will tally the number of correct answers given.  Once the time limit
is up, the test will stop.  The number of correct answers will be given as
well as the total amount of questions on the quiz.  

If the student is in the middle of answering a question and time is up, the
test will stop, not allowing the student to finish the question.

Have fun formulating your own tests!
