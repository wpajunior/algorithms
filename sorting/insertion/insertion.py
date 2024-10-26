def insertionSort(arr):
    """
    Sorts an array using the insertion sort algorithm.
    
    Parameters:
    arr (list): List of integers to be sorted.
    """
    n = len(arr)
    for i in range(1, n):
        key = arr[i]
        j = i - 1
        while j >= 0 and arr[j] > key:
            arr[j + 1] = arr[j]
            j = j - 1

        arr[j + 1] = key


def main():
    print("Insertion Sort Algorithm")
    n = int(input("Enter the number of items: "))

    print("Enter the items separated by space")
    arr = [int(x) for x in input().split()]


    insertionSort(arr)

    print(arr)

if __name__ == "__main__":
    main()

