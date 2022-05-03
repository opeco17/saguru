const sleep = (msec: number): Promise<void> => {
    return new Promise(function (resolve) {
        setTimeout(function () { resolve() }, msec);
    })
}

export default sleep;