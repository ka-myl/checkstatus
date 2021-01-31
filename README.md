# Checks Status
`checkstatus` is a command line tool that allows you to check the status of your (or not yours, who cares) websites. Yes, I know it's not exactly the most innovative idea, this is just a practice project. 

## Installation
1. Clone the repository
2. Run `go build` to create executable
3. Add the folder that contains your executable to your `PATH` variable

## Usage

- `checkstatus add <someurl>` Adds someurl to the list of your sites. Keep in mind you have to use full url (i.e. including protocol)
- `checkstatus remove <someurl>` Removes someurl from the list of your sites. You don't have to be very precise here, you can just use parts of the url
- `checkstatus list` Lists all the sites you are currently tracking
- `checkstatus run` Runs the check on all the sites and prints their status to the terminal

## Notes

- If you try to use `checkstatus run` with out adding any sites before, it will just add a `http://google.com` to the list of your tracking sites. Because why not.
