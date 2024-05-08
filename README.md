# Issues

This is a program that queries a project's Github issue tracker, via the GitHub web-service interface, and lists all the issues related to that project/repository.

The program displays essential information such as issue number, creator's GitHub username, and issue title in an organized tabular format

You can tweak it to display more details.

## How to Use
- Make sure you have the Go compiler installed on your computer.
- Clone this repository.

        git clone git@github.com:Samueelx/gh-issues.git


- Compile the program:

        go build github.com/Samueelx/issues

- Run the program, the command-line arguments specify your search params(such as the repo, status of the issues you are interested in, etc)

        ./issues repo:<repo-you-are-interested-in> is:<open|closed>

### Examples


        ./issues repo:sqlc-dev/sqlc is:open

        ./issues repo:sqlc-dev/sqlc is:open UPDATE