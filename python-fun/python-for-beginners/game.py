command = ""
started = False
while True:
    command = input(">").lower()
    if command == "start":
        if started:
            print("Game already started")
        else:
            started = True
        print("Game started")
    elif command == "stop":
        if not started:
            print("Game already stopped")
        else:
            started = False
        print("Game stopped")
    elif command == "help":
        print("""
        start - to start the game
        stop - to stop the game
        quit - to exit the game
        help - to get help       
              """)
    elif command == "quit":
        print("Game exited")
        break
    else:
        print("I don't understand that")


# The game is a simple text-based game where the user can start, stop, or get help about the game.