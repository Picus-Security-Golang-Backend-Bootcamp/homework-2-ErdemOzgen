package banner

import "fmt"

var asciiBanner string = `
#............^!JB@@#: .. J@@G5P55PB&@&PJ!^:....:~?5#@@@@@#^^^^~&@@@@@@@@@#~^^^~#@@@#5?~:.....:~7YB&@@
#               ~B@?    :&@#55PB&@G7:             .~5&@@B    .&@@@@@@@@@#.   .#@#!.              :G@
#    P#BBBBP!    .B&.    J@@B#@@#~    .!YGBBBG57:    :P@B    :&@@@@@@@@@#.   .#@~    !5GBBBG5?^ ^P@@
#    B@@@@@@@~    J@J    .#@@@@B.    ?#@@@@@@@@@&Y!?5G#@B    :&@@@@@@@@@#.   .#@:    J&@@@@@@@@#&@@@
#    B@@@@@&P.    P@&:    ?@@@@^    J@@@@@@@@@@@@@@@@@@@B    :&@@@@@@@@@#.   .#@5.     :^~7?Y5G#@@@@
#    ~!!!!~:     J@@@Y    .#@@#.    #@@@@@@@@@@@@@@@@@@@B    :&@@@@@@@@@#.   .#@@#Y!:.         .^?B@
#            .~JB@@@@&:    ?@@@^    Y@@@@@@@@@@@@@@@@@@@B    .&@@@@@@@@@#.   .#@@@@@&#GPYJ?!~.    .5
#    JP555PGB&@@@@@@@@Y    .#@@G.    ?&@@@@@@@@@&5!J5G#@&:    5@@@@@@@@@Y    ^@@@@P#@@@@@@@@@@J    :
#    B@@@@@@@@@@@@@@@@&:    7@@@B~    .!5GB##B57:    :P@@G.    !PB#&#B5!    .G@@G^ .~?5GB###B5^    !
#    B@@@@@@@@@@@@@@@@@5     B@@@@P!.      .      .^Y&@@@@#7.     ...     .7#@@&?:       ...     .?&
JJJJJ#@@@@@@@@@@@@@@@@@&J????G@@@@@@#P?~:.   .:^75B@@@@@@@@@#57~:.   .:~?5#@@@@@@&GY7~:.    .:^7Y#@@
			      PATIKA DEV BOOTCAMP ERDEM OZGEN
				   Version: 2.0.0
				  Date: 2022-03-04
`

var helpMenu string = `
Usage:
  homework-2-ErdemOzgen [command] # if you installed with go install
  go run main.go [command] # if you did not

Available Commands:
  buy         For buy books with id takes two arguments id and amount
  delete  	  For Delete books with id takes one argument id
  help        Print this help menu
  list        List all books in the dataset
  search	  Search for books with author,ISBN,ID takes one argument keyword
  banner 	  Print the ASCII banner that shows the name of the author,version,date and the command line
`

func PrintBanner() {
	fmt.Println(asciiBanner)
}

func PrintHelp() {
	PrintBanner()
	fmt.Println(helpMenu)
}
