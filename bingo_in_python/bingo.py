"""
🖋️Written by GRISZ (!!-|GRISZ|-!!#2705) the week-end of 15/16 october 2022 (it was on an another repo but i delete it so its not the good number of days on github)
"""

import os
import random

base_min_number = -1
base_max_number = 101
not_found = True


def error_warn(reason="error", color="c", pause_message="|Press enter to continue..."):
    os.system('cls')
    os.system(f'color {color}')
    print(reason)
    input(pause_message)
    os.system('color f')
    os.system('cls')


while True:

    if not not_found:
        break

    os.system('cls')
    print("\n|---------- Play To Bingo ----------|\n")
    min_number = input("Grasp the minimum number > ")
    max_number = input("Grasp the maximum number > ")

    try:
        min_number = int(min_number)
        max_number = int(max_number)

        if base_min_number < min_number < max_number < base_max_number:

            to_find_number = random.randint(min_number, max_number)

            print("\n|You can now try to find the good number")

            while not_found:

                inputted_number = input("Enter a number > ")

                try:
                    inputted_number = int(inputted_number)

                    if inputted_number == to_find_number:
                        os.system('cls')
                        os.system('color a')
                        print("|Wow you are a winner here is my password : azerty")
                        input("Press enter to quit the program...")
                        os.system('color f')
                        os.system('cls')
                        not_found = False

                    else:
                        print("\n|Maybe you should try again ... Or not Loser")
                        
                except ValueError:
                    error_warn("|Please enter number(s) not character(s)")

        elif min_number > max_number:
            error_warn("|Please enter a maximum number greater that the minimum number")

        else:
            error_warn("|Please enter numbers between 0 and 100")

    except ValueError:
        error_warn("|Please enter number(s) not character(s)")
        
"""
🖋️Written by GRISZ
"""
