import itertools
from typing import Tuple, Dict, List

SPINDLE_GEAR = 48
LEADSCREW_GEAR = 62
LEADSCREW_TPI = 8

HAVE_GEARS = [108, 102, 60]

MIN_TOOTH_COUNT = 30
MAX_TOOTH_COUNT = 134

ACCEPTABLE_ERROR = 0.005

# From 1/4" to 1"
UNC_TPI = {20, 18, 16, 14, 13, 12, 11, 10, 9, 8}
UNF_TPI = {28, 24, 20, 18, 16, 14, 12}
MM = {1, 1.25, 1.5, 1.75,2,2.5}
MM_TPI = [25.4 / mm for mm in MM]

DESIRED_TPIS = {*UNC_TPI, *UNF_TPI, *MM_TPI}

class Thread:
    def __init__(self, thread, error, gear1, gear2):
        self.thread = thread
        self.error = error
        self.gear1 = gear1
        self.gear2 = gear2

    def __repr__(self):
        return f"{self.thread} tpi ({self.error}%) @ {self.gear1}&{self.gear2}"


def tpi(change1 : int, change2 : int) -> float:
    return LEADSCREW_GEAR / SPINDLE_GEAR * change1 / change2 * LEADSCREW_TPI

def relative_error(expected : float, actual : float) -> float:
    return abs(1 - actual / expected)

def mm(change1 : int, change2 : int) -> float:
    return 1 / tpi(change1, change2) * 25.4 


def find_best(have_gears : List[int]):
    num_threads =0
    while num_threads < len(DESIRED_TPIS):
        have_gears, num_threads = find_best_next_gear(have_gears)

    print(have_gears)


def find_best_next_gear(have_gears : List[int]) -> Tuple[List[int], int]:
    """
    Returns a list of gears and the number of threads they can generate
    """
    best_possible_threads = {}
    best_gears = []

    for option in range(MIN_TOOTH_COUNT, MAX_TOOTH_COUNT + 1):
        all_gears = [*have_gears, option]
        possible_threads : Dict[float, Thread]= {}
        permutations = itertools.permutations(all_gears, 2)

        for (gear1, gear2) in permutations:
            thread = tpi(gear1, gear2)

            for desired in DESIRED_TPIS:
                err = relative_error(desired, thread)
                if err <= ACCEPTABLE_ERROR:
                    if desired not in possible_threads or possible_threads[desired].error > err:
                        possible_threads[desired] = Thread(thread, err, gear1, gear2)

        if len(possible_threads) > len(best_possible_threads):
            best_possible_threads = possible_threads
            best_gears = all_gears
    
    print("===============================")
    print(f"{len(have_gears) + 1} gears")
    print(f"Best possibility is {len(best_possible_threads)}/{len(DESIRED_TPIS)} ({DESIRED_TPIS})")
    print("Threads:", best_possible_threads)
    print("Gears: ", best_gears)
    print("MSE: ", sum(t.error **2 for t in best_possible_threads.values()) / len(best_possible_threads))

    return (best_gears, len(best_possible_threads))


if __name__ == "__main__":
    find_best(HAVE_GEARS)


