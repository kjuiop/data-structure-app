function bubbleSort(arr) {

    const swap = (arr, idx1, idx2) => {
        [arr[idx1], arr[idx2]] = [arr[idx2], arr[idx1]]
    };

    for (let i=arr.length; i>0; i--) {
        for (let j=0; j<i-1; j++) {
            console.log(arr, arr[j], arr[j+1]);
            if (arr[j] > arr[j+1]) {
                swap(arr, j, j+1);
            }
        }
        console.log("ONE PASS COMPLETE");
    }

    return arr;
}


function bubbleSort2(arr) {

    for (let i=arr.length; i>0; i--) {
        for (let j=0; j<i-1; j++) {
            console.log(arr, arr[j], arr[j+1]);
            if (arr[j] > arr[j+1]) {
                let temp = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = temp;
            }
        }
        console.log("ONE PASS COMPLETE");
    }

    return arr;
}

function bubbleSort1(arr) {
    let noSwap;
    for (let i=0; i<arr.length; i++) {
        noSwap = true;
        for (let j=0; j<arr.length; j++) {
            console.log(arr, arr[j], arr[j+1]);
            if (arr[j] > arr[j+1]) {
                let temp = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = temp;
                noSwap = false;
            }
        }

        if (noSwap) {
            break;
        }
    }

    return arr;
}

console.log("result : ", bubbleSort1([37,45,29,8,22,22]));
