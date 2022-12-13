package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Jump(t *testing.T) {
	tests := map[string]struct {
		arr    []int
		result int
	}{
		"case-1": {
			arr:    []int{2, 3, 1, 1, 4},
			result: 2,
		},
		"case-2": {
			arr:    []int{2, 3, 0, 1, 4},
			result: 2,
		},
		"degen": {
			arr:    []int{0},
			result: 0,
		},
		"options": {
			arr:    []int{2, 4, 1, 1, 0},
			result: 2,
		},
		"long": {
			arr:    []int{885, 387, 898, 946, 537, 938, 713, 194, 241, 104, 531, 980, 876, 427, 697, 51, 639, 669, 921, 150, 619, 583, 5, 540, 84, 75, 387, 51, 59, 248, 602, 953, 725, 959, 343, 183, 575, 867, 330, 596, 316, 670, 624, 292, 769, 21, 138, 816, 5, 358, 680, 605, 555, 553, 463, 208, 537, 410, 756, 469, 143, 779, 642, 504, 617, 682, 238, 366, 877, 699, 612, 556, 972, 6, 781, 216, 355, 170, 381, 234, 320, 434, 826, 382, 156, 121, 484, 512, 623, 723, 847, 948, 478, 355, 943, 740, 273, 835, 889, 613, 68, 446, 446, 480, 173, 943, 545, 522, 481, 887, 594, 274, 638, 532, 684, 539, 927, 916, 293, 92, 214, 690, 197, 331, 478, 101, 496, 454, 485, 954, 735, 718, 902, 644, 107, 25, 412, 596, 363, 794, 795, 768, 628, 621, 679, 696, 137, 962, 117, 454, 192, 801, 145, 471, 801, 54, 785, 548, 281, 69, 507, 886, 586, 695, 317, 530, 801, 786, 79, 180, 275, 675, 844, 963, 65, 627, 637, 129, 644, 837, 640, 499, 174, 138, 632, 745, 926, 913, 677, 662, 333, 188, 446, 287, 707, 257, 834, 101, 795, 24, 752, 618, 321, 866, 622, 932, 488, 867, 257, 844, 463, 161, 449, 406, 933, 996, 624, 625, 316, 622, 106, 486, 49, 396, 663, 34, 765, 972, 398, 853, 699, 441, 879, 479, 946, 816, 222, 67, 398, 682, 518, 174, 370, 227, 444, 764, 105, 378, 517, 130, 588, 628, 302, 630, 697, 715, 34, 163, 912, 688, 699, 301, 978, 791, 937, 96, 942, 946, 641, 539, 266, 350, 101, 178, 447, 261, 16, 487, 25, 39, 763, 775, 396, 302, 609, 568, 529, 558, 598, 203, 247, 469, 410, 465, 395, 486, 393, 84, 133, 389, 758, 317, 288, 666, 960, 900, 534, 268, 806, 348, 465, 131, 433, 810, 96, 470, 400, 411, 659, 850, 371, 6, 869, 646, 89, 728, 712, 898, 989, 369, 829, 255, 435, 902, 438, 426, 301, 67, 356, 272, 342, 90, 857, 331, 29, 864, 925, 765, 638, 839, 815, 862, 438, 482, 223, 648, 605, 290, 401, 915, 542, 608, 406, 672, 877, 59, 410, 423, 136, 967, 873, 519, 240, 653, 375, 179, 252, 392, 184, 989, 961, 184, 612, 395, 720, 856, 423, 830, 352, 366, 702, 574, 190, 367, 276, 10, 67, 610, 510, 903, 166, 102, 855, 855, 932, 689, 942, 772, 975, 300, 910, 901, 148, 743, 7, 852, 284, 990, 208, 623, 34, 524, 891, 226, 251, 839, 791, 225, 878, 895, 10, 164, 808, 974, 36, 166, 281, 448, 363, 438, 775, 289, 206, 69, 988, 531, 158, 972, 410, 484, 20, 983, 776, 252, 572, 221, 571, 52, 933, 246, 952, 501, 909, 750, 913, 660, 848, 166, 578, 356, 962, 773, 750, 133, 643, 555, 59, 614, 147, 744, 237, 696, 873, 517, 21, 561, 855, 56, 541, 905, 266, 679, 68, 87, 553, 325, 173, 982, 528, 315, 866, 312, 877, 521, 494, 486, 350, 784, 880, 125, 111, 223, 272, 747, 916, 848, 67, 211, 307, 783, 99, 372, 85, 691, 776, 965, 597, 693, 811, 570, 450, 131, 376, 345, 611, 606, 283, 513, 593, 588, 66, 278, 364, 226, 802, 181, 963, 389, 700, 565, 290, 277, 191, 254, 400, 419, 644, 941, 979, 253, 791, 593, 540, 995, 369, 591, 598, 836, 1000, 912, 728, 887, 146, 567, 971, 178, 822, 627, 460, 829, 249, 933, 619, 750, 530, 109, 232, 444, 986, 860, 289, 301, 257, 740, 435, 21, 264, 911, 996, 963, 249, 977, 633, 209, 992, 645, 400, 206, 303, 437, 849, 605, 618, 155, 781, 820, 11, 452, 94, 38, 434, 400, 863, 167, 911, 161, 664, 495, 761, 957, 380, 212, 189, 939, 303, 643, 754, 563, 769, 918, 468, 708, 28, 33, 573, 957, 8, 320, 173, 959, 289, 167, 643, 694, 165, 428, 930, 247, 248, 196, 678, 988, 259, 390, 895, 449, 349, 486, 630, 428, 813, 153, 355, 589, 469, 376, 969, 508, 439, 496, 748, 854, 481, 357, 967, 454, 806, 283, 615, 34, 970, 244, 846, 554, 619, 513, 57, 57, 447, 489, 992, 289, 916, 831, 740, 620, 686, 661, 825, 567, 47, 608, 977, 574, 163, 68, 597, 628, 905, 656, 120, 439, 228, 906, 867, 745, 925, 867, 686, 515, 339, 505, 879, 683, 146, 394, 133, 467, 606, 432, 894, 953, 126, 566, 296, 90, 108, 442, 499, 112, 338, 784, 780, 683, 625, 150, 771, 437, 589, 269, 256, 151, 413, 75, 890, 126, 636, 137, 139, 963, 96, 721, 990, 879, 391, 284, 249, 145, 848, 892, 279, 495, 424, 135, 753, 326, 608, 273, 803, 616, 699, 97, 961, 84, 724, 763, 744, 567, 182, 217, 464, 790, 521, 932, 630, 843, 836, 179, 467, 316, 273, 360, 804, 566, 808, 904, 898, 186, 789, 901, 925, 885, 861, 814, 241, 77, 913, 907, 260, 2, 642, 567, 394, 892, 279, 791, 787, 593, 898, 988, 778, 908, 855, 591, 673, 138, 527, 685, 239, 820, 300, 768, 986, 849, 503, 656, 96, 737, 359, 301, 252, 886, 208, 537, 272, 737, 532, 973, 323, 147, 184, 917, 631, 547, 690, 367, 518, 902, 948, 893, 611, 751, 327, 233, 836, 277, 472, 242, 653, 824, 865, 428, 601, 54, 600, 485, 303, 564, 18, 185, 80, 166, 600, 717, 645, 405, 91, 641, 942, 348, 720, 43, 432, 606, 172, 756, 332, 630, 747, 781, 65, 377, 992, 283, 232, 192, 86, 5, 573, 980, 692, 764, 458, 957, 631, 152, 719, 182, 781, 376, 169, 226, 473, 865, 159, 298, 34, 677, 332, 860, 235, 225, 792, 588, 573, 796, 781, 570, 339, 769, 85, 68, 646, 644, 851, 224, 667, 866, 408, 273, 464, 839, 510, 353, 142, 169, 1000, 523, 94, 522, 909, 906, 628, 549, 160, 475, 172, 712, 392, 725, 922, 54, 621, 651, 310, 90, 935, 453, 621, 858, 997, 892, 43, 928, 569, 262, 30, 704, 407, 977, 256, 155, 737, 920, 452, 852, 471, 79, 274, 113, 104, 268, 933, 213, 444, 442, 928, 885, 408, 337, 995, 360, 482, 883, 431, 459, 367, 276, 583, 830, 878, 77, 121, 374, 595, 598, 332, 838, 339, 677, 11, 137, 537, 231, 659, 16, 590, 576, 997, 562, 128, 131, 376, 357, 60, 257, 648, 930, 264, 217, 498, 765, 160, 598, 650, 283, 566, 789, 469, 77, 716, 651, 238, 830, 375, 439, 52, 254, 809, 804, 33, 227, 90, 274, 470, 885, 452, 50, 282, 52, 567, 63, 759, 185, 695, 574, 360, 554, 170, 214, 701, 299, 988, 435, 455, 891, 737, 961, 296, 849, 469, 715, 281, 447, 653, 29, 658, 567, 426, 336, 428, 620, 343, 568, 399, 936, 666, 263, 217, 190, 149, 866, 407, 944, 153, 270, 71, 533, 248, 92, 940, 634, 772, 619, 179, 376, 296, 507, 422, 864, 468, 731, 40, 583, 396, 258, 51, 900, 970, 923, 227, 557, 687, 360, 530, 344, 42, 479, 660, 12, 587, 5, 681, 743, 304, 857, 429, 951, 125, 639, 779, 632, 20, 960, 867, 349, 29, 397, 752, 633, 224, 495, 151, 690, 852, 234, 779, 115, 74, 508, 561, 380, 791, 864, 514, 512, 631, 998, 624, 453, 549, 346, 499, 944, 222, 523, 819, 222, 497, 38, 518, 800, 957, 346, 842, 762, 59, 873, 934, 580, 315, 880, 3, 769, 305, 521, 961, 349, 682, 438, 922, 813, 201, 325, 517, 451, 148, 666, 878, 983, 12, 919, 32, 40, 113, 2, 842, 45, 408, 557, 826, 144, 551, 479, 68, 632, 635, 893, 599, 128, 321, 762, 193, 596, 635, 48, 698, 335, 495, 158, 411, 815, 272, 894, 598, 201, 140, 639, 279, 607, 690, 12, 824, 435, 780, 361, 221, 712, 183, 260, 394, 417, 572, 310, 786, 84, 208, 673, 713, 923, 837, 239, 786, 618, 961, 756, 534, 409, 385, 607, 166, 491, 461, 549, 326, 465, 705, 557, 940, 484, 194, 951, 245, 216, 106, 89, 909, 895, 491, 60, 663, 326, 599, 703, 849, 865, 227, 643, 800, 646, 703, 62, 690, 587, 393, 497, 54, 709, 334, 747, 36, 800, 510, 574, 686, 500, 751, 368, 806, 510, 317, 697, 901, 690, 607, 221, 741, 686, 37, 151, 848, 787, 777, 18, 548, 539, 97, 287, 797, 651, 106, 267, 321, 369, 22, 199, 17, 64, 831, 527, 918, 735, 466, 711, 508, 385, 531, 745, 866, 111, 969, 618, 931, 884, 44, 878, 943, 403, 802, 376, 879, 807, 551, 371, 893, 244, 60, 499, 31, 753, 250, 458, 800, 45, 722, 905, 274, 99, 274, 907, 885, 866, 602, 380, 194, 282, 94, 502, 193, 6, 681, 885, 115, 930, 858, 613, 762, 34, 747, 173, 604, 713, 570, 627, 99, 928, 550, 169, 863, 284, 191, 940, 238, 88, 244, 851, 665, 75, 822, 705, 18, 360, 766, 931, 723, 919, 269, 374, 647, 551, 835, 672, 789, 739, 246, 568, 159, 980, 79, 294, 473, 339, 450, 554, 483, 527, 295, 362, 5, 191, 377, 355, 468, 533, 32, 674, 51, 640, 561, 526, 627, 390, 296, 570, 338, 329, 961, 1, 162, 298, 471, 730, 553, 906, 178, 214, 304, 736, 453, 944, 972, 546, 807, 892, 912, 582, 769, 267, 78, 8, 165, 236, 239, 995, 642, 783, 274, 516, 79, 901, 805, 485, 59, 630, 520, 41, 480, 132, 444, 126, 718, 155, 796, 558, 945, 799, 878, 864, 173, 911, 19, 448, 522, 761, 763, 603, 483, 653, 832, 552, 997, 843, 988, 365, 150, 830, 348, 657, 774, 396, 479, 727, 795, 409, 305, 841, 796, 747, 172, 372, 976, 891, 328, 9, 441, 561, 379, 187, 164, 446, 372, 151, 999, 951, 357, 545, 796, 919, 337, 257, 111, 194, 228, 982, 241, 725, 65, 681, 872, 329, 623, 538, 974, 272, 918, 86, 315, 556, 364, 942, 514, 905, 771, 113, 139, 746, 276, 324, 552, 927, 213, 514, 281, 66, 679, 72, 884, 988, 489, 8, 732, 880, 710, 120, 605, 482, 421, 584, 353, 709, 257, 989, 117, 471, 981, 430, 356, 302, 992, 920, 881, 504, 43, 210, 35, 906, 658, 946, 117, 457, 855, 221, 976},
			result: 2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			fmt.Println(name)
			assert.Equal(t, test.result, jump(test.arr))
		})
	}
}
