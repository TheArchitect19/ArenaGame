# Magical Arena

## Overview

The Magical Arena is a command-line game where players battle based on defined attributes (health, strength, and attack) until one player's health reaches zero. Players alternate turns attacking and defending using dice rolls to determine the outcome.

## Features

- Add new players to the arena with specified health, strength, and attack attributes.
- Display the list of current players in the arena.
- Simulate battles between two players until one player's health is depleted.
- Automatically remove defeated players from the arena.
- Comprehensive input validation for player attributes and commands.

## Requirements

- GoLang (version 1.22 or higher)

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/TheArchitect19/ArenaGame.git
    cd ArenaGame
    ```

2. Build the application:
    ```bash
    go build -o magical-arena
    ```

3. Run the application:
    ```bash
    ./magical-arena
    ```

## Usage

Upon running the application, you will encounter the following menu options:

1. **Add new player**: Allows you to add a new player to the arena. You will be prompted to enter the player's name, health, strength, and attack attributes.
2. **Battle**: Initiates a battle between two players. Enter the IDs of the two players who will fight.
3. **End game**: Terminates the game and exits the application.

## Example

Welcome to the arena!!!

Options:
1> Add new player
2> Battle
3> End game

Enter your choice (integer): 1
Enter the player's name: PlayerA
Enter PlayerA's health: 50
Enter PlayerA's attack: 10
Enter PlayerA's strength: 5

Options:
1> Add new player
2> Battle
3> End game

Enter your choice (integer): 1
Enter the player's name: PlayerB
Enter PlayerB's health: 100
Enter PlayerB's attack: 5
Enter PlayerB's strength: 10

Options:
1> Add new player
2> Battle
3> End game

Enter your choice (integer): 2
Enter the first player's id: 0
Enter the second player's id: 1

PlayerA v/S PlayerB

PlayerA hits PlayerB with power = 50
PlayerB defends with power = 20
PlayerB's health: 70

PlayerB hits PlayerA with power = 20
PlayerA defends with power = 15
PlayerA's health: 45

...

Bye Bye...


## Testing

To run unit tests for the application, execute the following command:

```bash
cd test
go test -v
```
## Running with Docker Compose

To start the application using Docker Compose, run the following command in your terminal:

```bash
docker-compose up


