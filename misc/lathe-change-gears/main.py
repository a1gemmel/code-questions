import itertools
from typing import Tuple, Dict

SPINDLE_GEAR = 48
LEADSCREW_GEAR = 62
LEADSCREW_TPI = 8

HAVE_GEARS = [108, 102, 60]

MIN_TOOTH_COUNT = 30
MAX_TOOTH_COUNT = 134

ACCEPTABLE_ERROR_PERCENT = 0.6

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
    return round(LEADSCREW_GEAR / SPINDLE_GEAR * change1 / change2 * LEADSCREW_TPI, 3)

def relative_error(expected : float, actual : float) -> float:
    return round(abs(1 - actual / expected) * 100, 3)

def mm(change1 : int, change2 : int) -> float:
    return 1 / tpi(change1, change2) * 25.4 


def combinatorial_find_best_for_n_gears(num_gears : int):
    best_possible_threads = {}
    best_gears = []

    for option in itertools.combinations(range(MIN_TOOTH_COUNT, MAX_TOOTH_COUNT + 1, 2), num_gears):
        all_gears = [*HAVE_GEARS, *option]
        possible_threads : Dict[float, Thread]= {}
        #print(f"Testing {all_gears}")

        permutations = itertools.permutations(all_gears, 2)

        for (gear1, gear2) in permutations:
            thread = tpi(gear1, gear2)

            for desired in DESIRED_TPIS:
                err = relative_error(desired, thread)
                if err <= ACCEPTABLE_ERROR_PERCENT:
                    if desired not in possible_threads or possible_threads[desired].error > err:
                        possible_threads[desired] = Thread(thread, err, gear1, gear2)

        if len(possible_threads) > len(best_possible_threads):
            best_possible_threads = possible_threads
            best_gears = all_gears
    print("================================")
    print(f"Best possibility is {len(best_possible_threads)}/{len(DESIRED_TPIS)} ({DESIRED_TPIS})")
    print("Threads:", best_possible_threads)
    print("Gears: ", best_gears)







if __name__ == "__main__":
    combinatorial_find_best_for_n_gears(4)


