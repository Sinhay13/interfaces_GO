Greetings in Multiple Languages

This is a simple Go program that demonstrates the use of interfaces to implement different types of bots that can generate greetings in different languages. The program defines two types of bots: englishBot and spanishBot, both of which implement the bot interface. The bot interface defines a single method getGreeting() that returns a string representing a greeting in a specific language.

The program then defines a printGreeting() function that takes a bot interface as an argument and prints the greeting returned by the getGreeting() method of the bot. Finally, the program creates instances of both englishBot and spanishBot and calls the printGreeting() function with each of them.

To run the program, simply execute the main.go file using the Go compiler. The program will output the following greetings:

$ Hi There!
$ Hola!

That's it! This program is a simple example of how to use interfaces in Go to implement different types of objects that share a common behavior.